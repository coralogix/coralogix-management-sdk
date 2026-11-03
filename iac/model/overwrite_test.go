// Copyright 2026 Coralogix Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"errors"
	"os"
	"slices"
	"strings"
	"testing"

	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
)

const overwriteWrite = "ThingsService_OverwriteKindAThings"

// overwritePolicy is the policy of the kind A things of testdata/overwritething.yaml.
func overwritePolicy() Policy {
	return Policy{
		Existing:  true,
		Overwrite: true,
		GetQuery:  map[string]string{"kind": "THING_KIND_A"},
		Operations: OperationIDs{
			Create: overwriteWrite, Get: "ThingsService_ListThings", Update: overwriteWrite, Delete: overwriteWrite,
		},
		RequestWrappers: map[string]string{"Thing": "thing"},
		RequestNames:    map[string]string{"Thing.enabled": "disabled"},
		RequestNegated:  []string{"Thing.enabled"},
	}
}

func loadOverwriteThing(t *testing.T) *v3.Document {
	t.Helper()
	raw, err := os.ReadFile("testdata/overwritething.yaml")
	if err != nil {
		t.Fatal(err)
	}
	doc, err := Load(raw)
	if err != nil {
		t.Fatal(err)
	}
	return doc
}

// A collection resource: Get with a fixed query lists the items, one write replaces them, and each
// written item holds its fields in a wrapper under other names.
func TestOverwriteResource(t *testing.T) {
	doc := loadOverwriteThing(t)
	if report := ValidateWithPolicy(doc, "ThingList", OperationIDs{}, overwritePolicy()); len(report) != 0 {
		t.Fatalf("issues = %v, want none", report)
	}
	r, err := BuildWithPolicy(doc, "ThingList", OperationIDs{}, overwritePolicy())
	if err != nil {
		t.Fatal(err)
	}
	if !r.Overwrite || !r.Singleton {
		t.Fatalf("overwrite = %t, singleton = %t, want both", r.Overwrite, r.Singleton)
	}
	if r.Create.OperationID != overwriteWrite || r.Update.OperationID != overwriteWrite || r.Delete.OperationID != overwriteWrite {
		t.Fatalf("create %s, update %s, delete %s, want the write each time", r.Create.OperationID, r.Update.OperationID, r.Delete.OperationID)
	}
	// The write returns what it created in its own shape. It is not the resource.
	if r.Create.Response.Direct || r.Create.Response.Field != "" || r.Create.Response.Schema != "OverwriteResponse" {
		t.Fatalf("create response = %+v, want OverwriteResponse, not read", r.Create.Response)
	}
	if r.Delete.Body != "inline" || r.Delete.BodyTitle != "Overwrite Kind A Things Request" {
		t.Fatalf("delete body = %q, title %q, want the inline body of the write", r.Delete.Body, r.Delete.BodyTitle)
	}
	if len(r.GetQuery) != 1 || r.GetQuery[0].Name != "kind" || r.GetQuery[0].Value != "THING_KIND_A" || r.GetQuery[0].Type.Kind != Enum || r.GetQuery[0].Type.Schema != "ThingKind" {
		t.Fatalf("get query = %+v, want kind = THING_KIND_A of ThingKind", r.GetQuery)
	}
	things := r.Fields[0]
	if things.Name != "things" || things.Type.Kind != List || things.Create == nil || things.Update == nil {
		t.Fatalf("field = %+v, want the managed list things", things)
	}
	// The list takes the limits of the write: Get lists the things of every kind.
	if things.Type.MaxItems == nil || *things.Type.MaxItems != 50 {
		t.Fatalf("maxItems = %v, want 50 of the write", things.Type.MaxItems)
	}
	item := things.Type.Elem
	if item.CreateSchema != "KindAThingWrite" {
		t.Fatalf("item create schema = %q, want KindAThingWrite", item.CreateSchema)
	}
	for _, want := range []struct {
		name     string
		behavior Behavior
		path     []string
		objects  []string
		negated  bool
	}{
		{"name", Normal, []string{"thing", "name"}, []string{"ThingCore"}, false},
		{"description", Normal, []string{"thing", "description"}, []string{"ThingCore"}, false},
		{"enabled", Normal, []string{"thing", "disabled"}, []string{"ThingCore"}, true},
		{"aRules", Normal, nil, nil, false},
		// The write of kind A never sends the other arm, so only the server sets it.
		{"bRules", Computed, nil, nil, false},
		{"id", Computed, nil, nil, false},
		{"order", Computed, nil, nil, false},
	} {
		f := fieldNamed(item.Fields, want.name)
		if f == nil {
			t.Errorf("no item field %s", want.name)
			continue
		}
		if f.Behavior != want.behavior || !slices.Equal(f.RequestPath, want.path) || !slices.Equal(f.RequestObjects, want.objects) || f.Negated != want.negated {
			t.Errorf("%s: behavior %q, path %v, objects %v, negated %t; want %q, %v, %v, %t",
				want.name, f.Behavior, f.RequestPath, f.RequestObjects, f.Negated, want.behavior, want.path, want.objects, want.negated)
		}
	}
	// The request type keeps the path, so a tool builds the request from it.
	request := things.Type.CreateType().Elem
	if f := fieldNamed(request.Fields, "enabled"); f == nil || !slices.Equal(f.RequestPath, []string{"thing", "disabled"}) || !f.Negated {
		t.Fatalf("request field enabled = %+v, want thing.disabled, negated", f)
	}
}

func TestOverwriteResourceIssues(t *testing.T) {
	for name, test := range map[string]struct {
		edit func(p *Policy)
		spec func(s string) string
		want string
	}{
		"unknown query parameter": {edit: func(p *Policy) { p.GetQuery = map[string]string{"type": "A"} }, want: "GET_QUERY_UNKNOWN"},
		"enum value not accepted": {edit: func(p *Policy) { p.GetQuery = map[string]string{"kind": "THING_KIND_C"} }, want: "GET_QUERY_VALUE_INVALID"},
		"boolean value not accepted": {edit: func(p *Policy) {
			p.GetQuery = map[string]string{"kind": "THING_KIND_A", "enabled_only": "yes"}
		}, want: "GET_QUERY_VALUE_INVALID"},
		"another delete": {
			spec: func(s string) string {
				return strings.Replace(s, "  /v1/kind-a-things:\n", "  /v1/kind-a-things/remove:\n    post:\n      operationId: ThingsService_RemoveThings\n      tags: [Things Service]\n      responses: {'200': {description: ok, content: {application/json: {schema: {type: object}}}}}\n  /v1/kind-a-things:\n", 1)
			},
			edit: func(p *Policy) { p.Operations.Delete = "ThingsService_RemoveThings" },
			want: "OVERWRITE_SHAPE_UNSUPPORTED",
		},
		"a scalar the write sends": {
			spec: func(s string) string {
				s = strings.Replace(s, "              required: [things]\n              properties:\n", "              required: [things]\n              properties:\n                label: {type: string}\n", 1)
				return strings.Replace(s, "      properties:\n        things:\n          type: array\n          minItems: 0\n          maxItems: 1000\n", "      properties:\n        label: {type: string}\n        things:\n          type: array\n          minItems: 0\n          maxItems: 1000\n", 1)
			},
			want: "OVERWRITE_FIELD_NOT_COLLECTION",
		},
		"inline write response with fields": {
			spec: func(s string) string {
				return strings.Replace(s, "        '200': {description: ok, content: {application/json: {schema: {$ref: '#/components/schemas/OverwriteResponse'}}}}",
					"        '200': {description: ok, content: {application/json: {schema: {type: object, properties: {count: {type: integer}}}}}}", 1)
			},
			want: "OVERWRITE_SHAPE_UNSUPPORTED",
		},
		"query value above the maximum": {
			spec: func(s string) string {
				return strings.Replace(s, "        - {name: enabled_only,", "        - {name: limit, in: query, required: false, schema: {type: integer, format: int32, maximum: 10}}\n        - {name: enabled_only,", 1)
			},
			edit: func(p *Policy) { p.GetQuery = map[string]string{"kind": "THING_KIND_A", "limit": "99"} },
			want: "GET_QUERY_VALUE_INVALID",
		},
		"query value outside int32": {
			spec: func(s string) string {
				return strings.Replace(s, "        - {name: enabled_only,", "        - {name: limit, in: query, required: false, schema: {type: integer, format: int32}}\n        - {name: enabled_only,", 1)
			},
			edit: func(p *Policy) { p.GetQuery = map[string]string{"kind": "THING_KIND_A", "limit": "3000000000"} },
			want: "GET_QUERY_VALUE_INVALID",
		},
		"wrapper not in the request":  {edit: func(p *Policy) { p.RequestWrappers = map[string]string{"Thing": "core"} }, want: "REQUEST_WRAPPER_NOT_FOUND"},
		"nested path not found":       {edit: func(p *Policy) { p.RequestNames = map[string]string{"Thing.enabled": "off"} }, want: "REQUEST_NAME_NOT_FOUND"},
		"negated field is not a bool": {edit: func(p *Policy) { p.RequestNegated = []string{"Thing.enabled", "Thing.name"} }, want: "REQUEST_NEGATED_UNSUPPORTED"},
		"line on a component that the resource does not hold": {
			edit: func(p *Policy) { p.RequestWrappers = map[string]string{"Thing": "thing", "Other": "inner"} },
			want: "REQUEST_NAME_UNKNOWN_FIELD",
		},
		"negated list on a nested component": {
			edit: func(p *Policy) { p.RequestNegated = []string{"Thing.enabled", "Rules.values"} },
			want: "REQUEST_NEGATED_UNSUPPORTED",
		},
	} {
		t.Run(name, func(t *testing.T) {
			raw, err := os.ReadFile("testdata/overwritething.yaml")
			if err != nil {
				t.Fatal(err)
			}
			spec := string(raw)
			if test.spec != nil {
				spec = test.spec(spec)
			}
			doc, err := Load([]byte(spec))
			if err != nil {
				t.Fatal(err)
			}
			p := overwritePolicy()
			if test.edit != nil {
				test.edit(&p)
			}
			var codes []string
			for _, i := range ValidateWithPolicy(doc, "ThingList", OperationIDs{}, p) {
				codes = append(codes, i.Code)
			}
			if !slices.Contains(codes, test.want) {
				t.Fatalf("issues = %v, want %s", codes, test.want)
			}
			if _, err := BuildWithPolicy(doc, "ThingList", OperationIDs{}, p); err == nil {
				t.Fatalf("Build accepted the resource, want %s", test.want)
			}
		})
	}
}

// A write that is a POST replaces every item as a PUT does, so the resource is a replacement and
// needs no update mask.
func TestOverwritePostIsAReplace(t *testing.T) {
	raw, err := os.ReadFile("testdata/overwritething.yaml")
	if err != nil {
		t.Fatal(err)
	}
	doc, err := Load([]byte(strings.Replace(string(raw), "  /v1/kind-a-things:\n    put:\n", "  /v1/kind-a-things:\n    post:\n", 1)))
	if err != nil {
		t.Fatal(err)
	}
	r, err := BuildWithPolicy(doc, "ThingList", OperationIDs{}, overwritePolicy())
	if err != nil {
		t.Fatal(err)
	}
	if r.Update.Method != "POST" || !r.Replace {
		t.Fatalf("update method %s, replace %t; want POST and a replacement", r.Update.Method, r.Replace)
	}
}

// A fixed query value inside the range of its parameter is accepted.
func TestOverwriteQueryInRange(t *testing.T) {
	raw, err := os.ReadFile("testdata/overwritething.yaml")
	if err != nil {
		t.Fatal(err)
	}
	doc, err := Load([]byte(strings.Replace(string(raw), "        - {name: enabled_only,", "        - {name: limit, in: query, required: false, schema: {type: integer, format: int32, minimum: 1, maximum: 10}}\n        - {name: enabled_only,", 1)))
	if err != nil {
		t.Fatal(err)
	}
	p := overwritePolicy()
	p.GetQuery = map[string]string{"kind": "THING_KIND_A", "limit": "10"}
	if _, err := BuildWithPolicy(doc, "ThingList", OperationIDs{}, p); err != nil {
		t.Fatal(err)
	}
}

// Without the overwrite policy, the same contract has no lifecycle: the write is no Create.
func TestOverwriteNeedsThePolicy(t *testing.T) {
	doc := loadOverwriteThing(t)
	p := overwritePolicy()
	p.Overwrite = false
	if _, err := BuildWithPolicy(doc, "ThingList", OperationIDs{}, p); err == nil {
		t.Fatal("Build accepted the write as a Create without api.overwrite")
	}
}

// A keyedBy line keys a set of objects as it keys a list: a map has no order either.
func TestKeyedBySet(t *testing.T) {
	for _, kind := range []Kind{List, Set} {
		item := &Type{Kind: Object, Schema: "TagRule", Fields: []*Field{
			{Name: "tagName", Type: &Type{Kind: String}},
			{Name: "tagValue", Type: &Type{Kind: String}},
		}}
		typ := &Type{Kind: kind, Elem: item}
		p := &Policy{KeyedBy: map[string]KeyedLine{"SpanRules.tagRules": {Key: "tagName"}}}
		p.asSet("SpanRules", "tagRules", typ)
		if typ.Kind != Map || typ.Keyed == nil || typ.Keyed.Key != "tagName" || len(typ.Elem.Fields) != 1 {
			t.Fatalf("%s: type = %+v, want a map keyed by tagName", kind, typ)
		}
	}
}

// The collection lines of the response component apply to the request object that the item shape
// rewrites: a keyedBy line keys the list of the request item as it keys the response.
func TestItemShapeAppliesCollectionLines(t *testing.T) {
	item := &Type{Kind: Object, Schema: "TagRule", Fields: []*Field{{Name: "tagName", Type: &Type{Kind: String}}, {Name: "tagValue", Type: &Type{Kind: String}}}}
	get := &Type{Kind: Object, Schema: "Thing", Fields: []*Field{{Name: "tags", Type: &Type{Kind: Map}}}}
	req := &Type{Kind: Object, Schema: "ThingWrite", Fields: []*Field{{Name: "tags", Type: &Type{Kind: Set, Elem: item}}}}
	p := Policy{RequestWrappers: map[string]string{}, RequestNegated: []string{}, KeyedBy: map[string]KeyedLine{"Thing.tags": {Key: "tagName"}}}
	shaped, err := p.shapeObject("things", get, req, map[string]bool{})
	if err != nil {
		t.Fatal(err)
	}
	if f := fieldNamed(shaped.Fields, "tags"); f == nil || f.Type.Kind != Map {
		t.Fatalf("request tags = %+v, want a map keyed by tagName", f)
	}
	if req.Fields[0].Type.Kind != Set {
		t.Fatal("the shape changed the request type that it was given")
	}
}

// nestedShapes is a response component Rules and its request object, which holds enabled negated
// as details.off, next to a note that stays in details.
func nestedShapes() (get, req *Type, p Policy) {
	get = &Type{Kind: Object, Schema: "Rules", Fields: []*Field{
		{Name: "values", Type: &Type{Kind: String}},
		{Name: "enabled", Type: &Type{Kind: Bool}},
	}}
	details := &Type{Kind: Object, Schema: "Details", Fields: []*Field{
		{Name: "off", Type: &Type{Kind: Bool}, Attrs: Attrs{Required: true}},
		{Name: "note", Type: &Type{Kind: String}},
	}}
	req = &Type{Kind: Object, Schema: "RulesWrite", Fields: []*Field{
		{Name: "values", Type: &Type{Kind: String}},
		{Name: "details", Type: details, Attrs: Attrs{Required: true}},
	}}
	p = Policy{RequestNames: map[string]string{"Rules.enabled": "details.off"}, RequestNegated: []string{"Rules.enabled"}}
	return get, req, p
}

// A nested requestAs path through an object of the request moves the leaf beside the other
// fields, and records the objects on the way.
func TestNestedRequestPathThroughObject(t *testing.T) {
	get, req, p := nestedShapes()
	used := map[string]bool{}
	shaped, err := p.shapeObject("things", get, req, used)
	if err != nil {
		t.Fatal(err)
	}
	f := fieldNamed(shaped.Fields, "enabled")
	if f == nil || !slices.Equal(f.RequestPath, []string{"details", "off"}) || !slices.Equal(f.RequestObjects, []string{"Details"}) || !f.Negated || !f.Attrs.Required {
		t.Fatalf("enabled = %+v, want details.off through Details, negated and required", f)
	}
	details := fieldNamed(shaped.Fields, "details")
	if details == nil || len(details.Type.Fields) != 1 || details.Type.Fields[0].Name != "note" {
		t.Fatalf("details = %+v, want only note left", details)
	}
	if len(req.Fields[1].Type.Fields) != 2 {
		t.Fatal("the shape changed the request type that it was given")
	}
	if !used["Rules.enabled"] {
		t.Fatal("the line is not marked used")
	}
}

// A renamed nested field that only one request has (an immutable field) is fine; a path that no
// request has is REQUEST_NAME_NOT_FOUND.
func TestNestedRequestPathInOneRequest(t *testing.T) {
	get, req, p := nestedShapes()
	without := &Type{Kind: Object, Schema: "RulesUpdate", Fields: []*Field{{Name: "values", Type: &Type{Kind: String}}}}
	used := map[string]bool{}
	for _, r := range []*Type{req, without} {
		if _, err := p.shapeObject("things", get, r, used); err != nil {
			t.Fatal(err)
		}
	}
	if err := p.missingNestedNames("ThingList", used); err != nil {
		t.Fatalf("missing = %v, want none: Create has the path", err)
	}
	used = map[string]bool{}
	if _, err := p.shapeObject("things", get, without, used); err != nil {
		t.Fatal(err)
	}
	var ce *contractError
	if err := p.missingNestedNames("ThingList", used); !errors.As(err, &ce) || ce.code != "REQUEST_NAME_NOT_FOUND" {
		t.Fatalf("missing = %v, want REQUEST_NAME_NOT_FOUND", err)
	}
}
