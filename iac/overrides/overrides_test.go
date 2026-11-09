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

package overrides

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"go.yaml.in/yaml/v4"

	"github.com/coralogix/coralogix-management-sdk/iac/model"
)

func TestPolicyStatesTheAPIKeys(t *testing.T) {
	on, off := true, false
	e := Exceptions{
		API: API{
			RequestWrapper: "router", UpdateIDInBody: true, ClientSetID: true, CustomMethods: true, UpdateMaskInBody: true,
			Operations: &Operations{Create: "C", Get: "G", Update: "U", Delete: "D", List: "L"},
		},
		Unwrap: []string{"LuceneQuery"},
		Fields: map[string]map[string]Field{
			"Thing":  {"id": {ReadOnly: true}, "query": {Unwrap: &off}},
			"Target": {"id": {Skip: true}, "label": {Unwrap: &on}},
		},
	}
	want := model.Policy{
		Existing: true, RequestWrapper: "router", UpdateIDInBody: true, ClientSetID: true, CustomMethods: true, UpdateMaskInBody: true,
		Operations:   model.OperationIDs{Create: "C", Get: "G", Update: "U", Delete: "D", List: "L"},
		Unwrap:       []string{"LuceneQuery"},
		UnwrapFields: map[string]bool{"Target.label": true, "Thing.query": false},
		Skip:         []string{"Target.id"},
		ReadOnly:     []string{"Thing.id"},
		Released:     []string{"Target.id", "Thing.id"},
	}
	if got := e.Policy(); !reflect.DeepEqual(got, want) {
		t.Fatalf("Policy() =\n%+v\nwant\n%+v", got, want)
	}
	if got := (Exceptions{API: API{Delete: &Delete{Operation: "Archive"}}}).Policy(); got.DeleteOperation != "Archive" {
		t.Fatalf("DeleteOperation = %q", got.DeleteOperation)
	}
}

func TestPolicyStatesTheFieldShapeKeys(t *testing.T) {
	e := Exceptions{
		API: API{IDProperty: "uniqueIdentifier"},
		Fields: map[string]map[string]Field{
			"Thing": {
				"name":   {Optional: true},
				"tags":   {Collection: "set"},
				"labels": {KeyedBy: &KeyedBy{Key: "key", Value: "value"}},
				"limits": {CreateAs: map[string]string{"limit": "permutationsLimit"}},
			},
		},
	}
	want := model.Policy{
		Existing: true, IDProperty: "uniqueIdentifier",
		Optional: []string{"Thing.name"},
		Sets:     []string{"Thing.tags"},
		KeyedBy:  map[string]model.KeyedLine{"Thing.labels": {Key: "key", Value: "value"}},
		CreateAs: map[string]map[string]string{"Thing.limits": {"limit": "permutationsLimit"}},
	}
	if got := e.Policy(); !reflect.DeepEqual(got, want) {
		t.Fatalf("Policy() =\n%+v\nwant\n%+v", got, want)
	}
}

func TestFieldCheckRejectsBadShapeKeys(t *testing.T) {
	for name, l := range map[string]Field{
		"keyedBy without a key":        {KeyedBy: &KeyedBy{Value: "value"}},
		"keyedBy value is the key":     {KeyedBy: &KeyedBy{Key: "key", Value: "key"}},
		"keyedBy object with value":    {KeyedBy: &KeyedBy{Key: "key", Value: "value", Object: true}},
		"keyedBy arms without object":  {KeyedBy: &KeyedBy{Key: "key", Arms: map[string]string{"A": "a"}}},
		"keyedBy and collection":       {KeyedBy: &KeyedBy{Key: "key"}, Collection: "set"},
		"createAs with an empty name":  {CreateAs: map[string]string{"limit": ""}},
		"createAs names a field twice": {CreateAs: map[string]string{"a": "x", "b": "x"}},
		"collection other than set":    {Collection: "list"},
	} {
		if l.Check() == nil {
			t.Errorf("%s: Check() = nil", name)
		}
	}
	ok := Field{Collection: "set", CreateAs: map[string]string{"limit": "permutationsLimit"}}
	if err := ok.Check(); err != nil {
		t.Fatal(err)
	}
}

func TestCheckRejectsConflictingAPIKeys(t *testing.T) {
	for name, a := range map[string]API{
		"both id owners":        {ClientSetID: true, ServerSetID: true},
		"delete with no name":   {Delete: &Delete{}},
		"incomplete operations": {Operations: &Operations{Create: "C"}},
		"no operations":         {Operations: &Operations{}},
		"list and one more":     {Operations: &Operations{Get: "G", List: "L"}},
		"two delete operations": {Operations: &Operations{Create: "C", Get: "G", Update: "U", Delete: "D"}, Delete: &Delete{Operation: "Archive"}},
	} {
		if a.Check() == nil {
			t.Errorf("%s: Check() = nil", name)
		}
	}
	if err := (API{Delete: &Delete{Operation: "Archive"}}).Check(); err != nil {
		t.Fatal(err)
	}
	// The List can be the only operation that the file names: the lifecycle uses discovery.
	if err := (API{Operations: &Operations{List: "L"}, Delete: &Delete{Operation: "Archive"}}).Check(); err != nil {
		t.Fatal(err)
	}
	if CheckUnwrap([]string{"A", "A"}) == nil || CheckUnwrap([]string{""}) == nil || CheckUnwrap([]string{"A", "B"}) != nil {
		t.Fatal("CheckUnwrap does not reject an empty or a repeated name")
	}
}

// A tool decodes its whole file once. Its field line embeds Field inline, so the API keys and
// its own keys share one line, and a key that neither part knows is still an error.
func TestToolEmbedsFieldInItsLine(t *testing.T) {
	type toolField struct {
		Field       `yaml:",inline"`
		Description *string `yaml:"description"`
	}
	type toolFile struct {
		API   API                             `yaml:"api"`
		Types map[string]map[string]toolField `yaml:"types"`
	}
	decode := func(text string) (toolFile, error) {
		dec := yaml.NewDecoder(bytes.NewReader([]byte(text)))
		dec.KnownFields(true)
		var f toolFile
		return f, dec.Decode(&f)
	}
	f, err := decode("api: {requestWrapper: router}\ntypes:\n  Target:\n    id: {skip: true, description: The id.}\n")
	if err != nil {
		t.Fatal(err)
	}
	if line := f.Types["Target"]["id"]; !line.Skip || line.Description == nil || *line.Description != "The id." || f.API.RequestWrapper != "router" {
		t.Fatalf("decoded %+v", f)
	}
	if _, err := decode("types:\n  Target:\n    id: {skipp: true}\n"); err == nil || !strings.Contains(err.Error(), "skipp") {
		t.Fatalf("unknown key error = %v", err)
	}
}

func TestContractReadOnly(t *testing.T) {
	doc, err := model.Load([]byte(`openapi: 3.1.0
info: {title: t, version: "1"}
paths: {}
components:
  schemas:
    Thing:
      type: object
      properties:
        id: {type: string, readOnly: true}
        name: {type: string}
`))
	if err != nil {
		t.Fatal(err)
	}
	schema, err := doc.Components.Schemas.GetOrZero("Thing").BuildSchema()
	if err != nil {
		t.Fatal(err)
	}
	if !ContractReadOnly(schema, "id") || ContractReadOnly(schema, "name") {
		t.Fatal("ContractReadOnly does not match the readOnly keyword")
	}
}

// api.createResponseOnlyIDProperty is an API key of the shared overrides: it reaches the policy.
func TestCreateResponseOnlyIDPropertyReachesThePolicy(t *testing.T) {
	var e Exceptions
	e.API.CreateResponseOnlyIDProperty = "thingId"
	if got := e.Policy().CreateResponseOnlyIDProperty; got != "thingId" {
		t.Fatalf("policy CreateResponseOnlyIDProperty = %q, want thingId", got)
	}
}

func TestPolicyStatesTheRequestShapeKeys(t *testing.T) {
	e := Exceptions{
		API: API{RequestWrapper: "metadata", UpdateIDInBody: true, UpdateIDOutsideWrapper: true},
		Fields: map[string]map[string]Field{"Thing": {
			"definitionKey": {RequestAs: "integrationKey"},
			"parameters":    {RequestAs: "integrationParameters.parameters"},
			"id":            {},
		}},
	}
	got := e.Policy()
	want := map[string]string{"Thing.definitionKey": "integrationKey", "Thing.parameters": "integrationParameters.parameters"}
	if !got.UpdateIDOutsideWrapper || !reflect.DeepEqual(got.RequestNames, want) {
		t.Fatalf("Policy() = %+v, want updateIDOutsideWrapper and the request names %v", got, want)
	}
	if p := (Exceptions{}).Policy(); p.UpdateIDOutsideWrapper || p.RequestNames != nil {
		t.Fatalf("Policy() without the keys = %+v", p)
	}
}

func TestCheckRequestShapeKeys(t *testing.T) {
	tests := map[string]struct {
		api  API
		want string
	}{
		"outside wrapper needs a wrapper":  {API{UpdateIDInBody: true, UpdateIDOutsideWrapper: true}, "updateIDOutsideWrapper"},
		"outside wrapper needs id in body": {API{RequestWrapper: "metadata", UpdateIDOutsideWrapper: true}, "updateIDOutsideWrapper"},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if err := tc.api.Check(); err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("Check() = %v, want an error naming %q", err, tc.want)
			}
		})
	}
	ok := API{RequestWrapper: "metadata", UpdateIDInBody: true, UpdateIDOutsideWrapper: true}
	if err := ok.Check(); err != nil {
		t.Fatalf("Check() = %v", err)
	}
	for _, path := range []string{"b..c", ".b", "b."} {
		if err := (Field{RequestAs: path}).Check(); err == nil || !strings.Contains(err.Error(), "requestAs") {
			t.Fatalf("Field{RequestAs: %q}.Check() = %v, want a requestAs error", path, err)
		}
	}
	if err := (Field{RequestAs: "d.e"}).Check(); err != nil {
		t.Fatalf("Check() = %v", err)
	}
}

func TestPolicyStatesTheOverwriteKeys(t *testing.T) {
	e := Exceptions{
		API: API{
			Overwrite:  &Overwrite{GetQuery: map[string]string{"source_type": "SOURCE_TYPE_LOGS"}, WriteReturnsResource: true},
			Operations: &Operations{Create: "W", Get: "G", Update: "W", Delete: "W"},
		},
		Fields:          map[string]map[string]Field{"Policy": {"enabled": {RequestAs: "disabled", Negate: true}}},
		RequestWrappers: map[string]string{"Policy": "policy"},
	}
	if err := e.API.Check(); err != nil {
		t.Fatal(err)
	}
	p := e.Policy()
	if !p.Overwrite || !p.WriteReturnsResource || !reflect.DeepEqual(p.GetQuery, map[string]string{"source_type": "SOURCE_TYPE_LOGS"}) {
		t.Errorf("overwrite = %t, write returns resource = %t, get query = %v", p.Overwrite, p.WriteReturnsResource, p.GetQuery)
	}
	if !reflect.DeepEqual(p.RequestWrappers, map[string]string{"Policy": "policy"}) || !reflect.DeepEqual(p.RequestNegated, []string{"Policy.enabled"}) || p.RequestNames["Policy.enabled"] != "disabled" {
		t.Errorf("wrappers = %v, negated = %v, names = %v", p.RequestWrappers, p.RequestNegated, p.RequestNames)
	}
}

func TestCheckOverwriteKeys(t *testing.T) {
	ops := func(create, update, del string) *Operations {
		return &Operations{Create: create, Get: "G", Update: update, Delete: del}
	}
	for name, api := range map[string]API{
		"no operations":          {Overwrite: &Overwrite{}},
		"another update":         {Overwrite: &Overwrite{}, Operations: ops("W", "U", "W")},
		"another delete":         {Overwrite: &Overwrite{}, Operations: ops("W", "W", "D")},
		"with a request wrapper": {Overwrite: &Overwrite{}, Operations: ops("W", "W", "W"), RequestWrapper: "thing"},
		"with an id key":         {Overwrite: &Overwrite{}, Operations: ops("W", "W", "W"), ServerSetID: true},
		"empty query value":      {Overwrite: &Overwrite{GetQuery: map[string]string{"kind": ""}}, Operations: ops("W", "W", "W")},
	} {
		t.Run(name, func(t *testing.T) {
			if err := api.Check(); err == nil {
				t.Fatal("Check accepted the keys")
			}
		})
	}
}

func TestOverwriteKeysParse(t *testing.T) {
	var api API
	in := "overwrite:\n  getQuery: {source_type: SOURCE_TYPE_SPANS}\n  writeReturnsResource: true\noperations: {create: W, get: G, update: W, delete: W}\n"
	if err := yaml.Unmarshal([]byte(in), &api); err != nil {
		t.Fatal(err)
	}
	if api.Overwrite == nil || api.Overwrite.GetQuery["source_type"] != "SOURCE_TYPE_SPANS" || !api.Overwrite.WriteReturnsResource {
		t.Fatalf("overwrite = %+v", api.Overwrite)
	}
	var line Field
	if err := yaml.Unmarshal([]byte("{requestAs: disabled, negate: true}"), &line); err != nil {
		t.Fatal(err)
	}
	if !line.Negate || line.RequestAs != "disabled" {
		t.Fatalf("line = %+v", line)
	}
}
