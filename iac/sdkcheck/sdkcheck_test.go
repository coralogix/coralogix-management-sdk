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

package sdkcheck

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/coralogix/coralogix-management-sdk/iac/issue"
	"github.com/coralogix/coralogix-management-sdk/iac/model"
)

func thingNames(t *testing.T, zeroConstants func(string) bool) []Ref {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("..", "model", "testdata", "openapi.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	doc, err := model.Load(data)
	if err != nil {
		t.Fatal(err)
	}
	r, err := model.Build(doc, "Thing")
	if err != nil {
		t.Fatal(err)
	}
	tag, err := ResourceTag(doc, r)
	if err != nil {
		t.Fatal(err)
	}
	if tag != "Things Service" || PackageName(tag) != "things_service" || ClientType(tag) != "ThingsServiceAPIService" {
		t.Fatalf("tag %q, package %q, client %q", tag, PackageName(tag), ClientType(tag))
	}
	refs, err := Names(r, tag, "example.com/sdk", zeroConstants)
	if err != nil {
		t.Fatal(err)
	}
	return refs
}

func TestFieldTypeUsesOperationWireCollection(t *testing.T) {
	object := &model.Type{Kind: model.Object, WireKind: model.Object, Schema: "Item", Fields: []*model.Field{{Name: "name", Type: &model.Type{Kind: model.String, WireKind: model.String}}}}
	for _, test := range []struct {
		name, want string
		wire       model.Kind
	}{
		{name: "request list", wire: model.List, want: "[]Item"},
		{name: "response map", wire: model.Map, want: "map[string]Item"},
	} {
		t.Run(test.name, func(t *testing.T) {
			got, _, err := fieldType(&model.Type{Kind: model.Set, WireKind: test.wire, Elem: object})
			if err != nil || got != test.want {
				t.Fatalf("fieldType = %q, %v; want %q", got, err, test.want)
			}
		})
	}

	keyed := &model.Type{Kind: model.Object, WireKind: model.Object, Schema: "Item"}
	for _, test := range []struct {
		name, want string
		wire       model.Kind
	}{
		{name: "keyed Get map", wire: model.Map, want: "map[string]Item"},
		{name: "keyed write list", wire: model.List, want: "[]Item"},
	} {
		t.Run(test.name, func(t *testing.T) {
			typeValue := &model.Type{
				Kind: model.Map, WireKind: test.wire,
				Keyed: &model.Keyed{Key: "name", Item: keyed},
			}
			got, _, err := fieldType(typeValue)
			if err != nil || got != test.want {
				t.Fatalf("fieldType = %q, %v; want %q", got, err, test.want)
			}
		})
	}
}

// Names lists only names of the SDK module: the generated package of the tag and the cxsdk
// package. A tool appends the names of its own code.
func TestNamesListsOnlySDKNames(t *testing.T) {
	refs := thingNames(t, nil)
	gen, cxsdk := "example.com/sdk/go/openapi/gen/things_service", "example.com/sdk/go/openapi/cxsdk"
	for _, ref := range refs {
		if ref.Pkg != gen && ref.Pkg != cxsdk {
			t.Errorf("ref %s %s is in package %s", ref.Kind, ref.SDKName(), ref.Pkg)
		}
	}
	for _, want := range []string{"ThingsServiceAPIService.ThingsServiceGetThing", "Thing.Name", "IsNotFound"} {
		if !slices.ContainsFunc(refs, func(ref Ref) bool { return ref.SDKName() == want }) {
			t.Errorf("refs do not contain %s", want)
		}
	}
}

// The zero constant of an enum is in the list only when the caller uses it.
func TestNamesListsZeroConstantsOnRequest(t *testing.T) {
	zero := func(refs []Ref) bool {
		return slices.ContainsFunc(refs, func(ref Ref) bool { return ref.Name == "THINGKIND_THING_KIND_UNSPECIFIED" })
	}
	if zero(thingNames(t, nil)) {
		t.Fatal("nil zeroConstants lists the zero constant")
	}
	if !zero(thingNames(t, func(component string) bool { return component == "ThingKind" })) {
		t.Fatal("zeroConstants for ThingKind does not list its zero constant")
	}
}

func TestIssuesChecksNamesAndSources(t *testing.T) {
	const pkg = "example.com/fakesdk/widgets"
	refs := []Ref{
		{Path: "resource", Kind: KindPackage, Pkg: pkg, Name: "widgets"},
		{Path: "fields", Kind: KindType, Pkg: pkg, Name: "Widget"},
		{Path: "fields.name", Kind: KindField, Pkg: pkg, Owner: "Widget", Name: "Name", Want: "*string"},
		{Path: "fields.count", Kind: KindField, Pkg: pkg, Owner: "Widget", Name: "Count", Want: "*int32", WantValue: "int32"},
		{Path: "get", Kind: KindMethod, Pkg: pkg, Owner: "WidgetsAPIService", Name: "Widget", Want: "func() string", ByType: true},
		{Path: "fields.size", Kind: KindField, Pkg: pkg, Owner: "Widget", Name: "Size", Want: "*int64"},
	}
	loaded, err := Load(refs, filepath.Join("testdata", "fakesdk"))
	if err != nil {
		t.Fatal(err)
	}
	report := Issues(refs, loaded, func(Ref) Source { return Source{Module: "example.com/fakesdk"} })
	if codes := codesOf(report); !slices.Equal(codes, []string{"SDK_SYMBOL_MISSING"}) || report[0].Location != "fields.size" {
		t.Fatalf("report = %v", report)
	}
	// The SDK has the value type and another method name: the refs now name what the SDK has.
	if refs[3].Want != "int32" || refs[4].Name != "GetWidget" {
		t.Fatalf("refs = %+v, %+v", refs[3], refs[4])
	}

	report = Issues(refs[:1], loaded, func(Ref) Source { return Source{Module: "example.com/other", Remediation: "Load the other module."} })
	if codes := codesOf(report); !slices.Equal(codes, []string{"SDK_SOURCE_MISMATCH"}) || report[0].Remediation != "Load the other module." {
		t.Fatalf("report = %v", report)
	}
	if !strings.Contains(report[0].Message, "expected example.com/other") {
		t.Fatalf("message = %s", report[0].Message)
	}
}

func TestSDKLoadingIsOffline(t *testing.T) {
	t.Setenv("GOPROXY", "https://proxy.invalid")
	t.Setenv("GOSUMDB", "sum.invalid")
	env := offlineGoEnv()
	if got := environmentValue(env, "GOPROXY"); got != "off" {
		t.Fatalf("GOPROXY = %q, want off", got)
	}
	if got := environmentValue(env, "GOSUMDB"); got != "off" {
		t.Fatalf("GOSUMDB = %q, want off", got)
	}
}

func environmentValue(env []string, name string) string {
	prefix := name + "="
	for _, item := range env {
		if strings.HasPrefix(item, prefix) {
			return strings.TrimPrefix(item, prefix)
		}
	}
	return ""
}

func codesOf(report issue.Report) []string {
	var codes []string
	for _, item := range report {
		codes = append(codes, item.Code)
	}
	return codes
}

// openapi-generator puts "_" between a digit and the letter after it in the constant of an enum
// value, and keeps the type name as one word.
func TestEnumConstName(t *testing.T) {
	for _, test := range []struct{ typ, value, want string }{
		{"E2MType", "E2M_TYPE_LOGS2METRICS", "E2MTYPE_E2_M_TYPE_LOGS2_METRICS"},
		{"Logs2metricsV2Severity", "SEVERITY_DEBUG", "LOGS2METRICSV2SEVERITY_SEVERITY_DEBUG"},
		{"ThingKind", "THING_KIND_P5_OR_UNSPECIFIED", "THINGKIND_THING_KIND_P5_OR_UNSPECIFIED"},
	} {
		if got := enumConstName(test.typ, test.value); got != test.want {
			t.Errorf("enumConstName(%q, %q) = %q, want %q", test.typ, test.value, got, test.want)
		}
	}
}

// The List names are in the list only when the resource has a List. A summary item adds its
// type and the fields that a tool reads.
func TestNamesListsTheList(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("..", "model", "testdata", "lists.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	doc, err := model.Load(data)
	if err != nil {
		t.Fatal(err)
	}
	names := func(resource string) []string {
		r, err := model.Build(doc, resource)
		if err != nil {
			t.Fatal(err)
		}
		tag, err := ResourceTag(doc, r)
		if err != nil {
			t.Fatal(err)
		}
		refs, err := Names(r, tag, "example.com/sdk", nil)
		if err != nil {
			t.Fatal(err)
		}
		var out []string
		for _, ref := range refs {
			if strings.HasPrefix(ref.Path, "list") {
				out = append(out, string(ref.Kind)+" "+ref.SDKName()+" "+ref.Want)
			}
		}
		return out
	}
	want := []string{
		"method LabelsServiceAPIService.LabelsServiceListLabels func(ctx context.Context) ApiLabelsServiceListLabelsRequest",
		"type ApiLabelsServiceListLabelsRequest ",
		"type ListLabelsResponse ",
		"method ApiLabelsServiceListLabelsRequest.Execute func() (*ListLabelsResponse, *http.Response, error)",
		"field ListLabelsResponse.Labels []LabelSummary",
		"type LabelSummary ",
		"field LabelSummary.Id *string",
	}
	if got := names("Label"); !slices.Equal(got, want) {
		t.Fatalf("Label list names =\n%s\nwant\n%s", strings.Join(got, "\n"), strings.Join(want, "\n"))
	}
	// A full item is the resource, whose names the fields cover.
	if got := names("Connector"); len(got) != 5 || got[4] != "field ListConnectorsResponse.Connectors []Connector" {
		t.Fatalf("Connector list names = %v", got)
	}
	if got := names("Pair"); len(got) != 0 {
		t.Fatalf("Pair list names = %v, want none: the List is not described", got)
	}
	if got := thingNames(t, nil); slices.ContainsFunc(got, func(ref Ref) bool { return strings.HasPrefix(ref.Path, "list") }) {
		t.Fatal("a resource without a List lists List names")
	}
}

// inlineRenameSpec is model/testdata/renamething.yaml with inline Create and Update bodies, as the
// SDK generator gets them from a proto request.
func inlineRenameSpec(t *testing.T) []byte {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("..", "model", "testdata", "renamething.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	spec := string(data)
	for old, inline := range map[string]string{
		"{$ref: '#/components/schemas/SaveRenameThingRequest'}":   "{title: Save rename thing request, type: object, required: [spec], properties: {spec: {$ref: '#/components/schemas/RenameSpec'}}}",
		"{$ref: '#/components/schemas/UpdateRenameThingRequest'}": "{title: Update rename thing request, type: object, required: [id, spec], properties: {id: {type: string, minLength: 1}, spec: {$ref: '#/components/schemas/RenameSpec'}}}",
	} {
		if !strings.Contains(spec, old) {
			t.Fatalf("the spec has no %q", old)
		}
		spec = strings.Replace(spec, old, inline, 1)
	}
	return []byte(spec)
}

// A request wrapper holds the request component of the spec, not the resource. A renamed field is a
// field of that component, or of an object inside it, and the Update id beside the wrapper is a
// field of the body.
func TestNamesForARequestShapeThatDiffersFromTheResource(t *testing.T) {
	doc, err := model.Load(inlineRenameSpec(t))
	if err != nil {
		t.Fatal(err)
	}
	policy := model.Policy{
		Existing: true, RequestWrapper: "spec", UpdateIDInBody: true, UpdateIDOutsideWrapper: true, IDProperty: "id",
		CustomMethods: true, CreateResponseOnlyIDProperty: "thingId",
		RequestNames: map[string]string{"RenameThing.kind": "specKind", "RenameThing.version": "specVersion", "RenameThing.parameters": "holder.items"},
		Operations: model.OperationIDs{
			Create: "RenameThingsService_SaveRenameThing", Get: "RenameThingsService_GetRenameThing",
			Update: "RenameThingsService_UpdateRenameThing", Delete: "RenameThingsService_DeleteRenameThing",
		},
	}
	r, err := model.BuildWithPolicy(doc, "RenameThing", model.OperationIDs{}, policy)
	if err != nil {
		t.Fatal(err)
	}
	tag, err := ResourceTag(doc, r)
	if err != nil {
		t.Fatal(err)
	}
	refs, err := Names(r, tag, "example.com/sdk", nil)
	if err != nil {
		t.Fatal(err)
	}
	find := func(path string) *Ref {
		for i := range refs {
			if refs[i].Path == path && refs[i].Kind == KindField {
				return &refs[i]
			}
		}
		return nil
	}
	for path, want := range map[string]string{
		"create.body.spec":       "SaveRenameThingRequest.Spec *RenameSpec",
		"update.body.spec":       "UpdateRenameThingRequest.Spec *RenameSpec",
		"update.body.id":         "UpdateRenameThingRequest.Id *string",
		"create.body.kind":       "RenameSpec.SpecKind *string",
		"update.body.version":    "RenameSpec.SpecVersion *string",
		"create.body.parameters": "RenameHolder.Items []RenameParam",
	} {
		ref := find(path)
		if ref == nil {
			t.Errorf("no field ref at %s", path)
			continue
		}
		if got := ref.SDKName() + " " + ref.Want; got != want {
			t.Errorf("%s = %q, want %q", path, got, want)
		}
	}
	typePath := func(path string) *Ref {
		for i := range refs {
			if refs[i].Path == path && refs[i].Kind == KindType {
				return &refs[i]
			}
		}
		return nil
	}
	if ref := typePath("create.body.spec"); ref == nil || ref.Name != "RenameSpec" || ref.Schema != "RenameSpec" {
		t.Errorf("type ref at create.body.spec = %+v, want the request component RenameSpec", ref)
	}
	if ref := typePath("create.body"); ref == nil || ref.Name != "SaveRenameThingRequest" {
		t.Errorf("type ref at create.body = %+v, want the body type SaveRenameThingRequest", ref)
	}
	if ref := find("create.body.id"); ref != nil {
		t.Errorf("Create has a field ref for the id: %s", ref.SDKName())
	}
	if ref := find("update.body.holder"); ref != nil {
		t.Errorf("a ref with the path %s: the field is named parameters", ref.Path)
	}
	for _, want := range []string{"RenameSpec.Holder", "RenameThing.Kind", "SaveRenameThingResponse.ThingId"} {
		if !slices.ContainsFunc(refs, func(ref Ref) bool { return ref.SDKName() == want }) {
			t.Errorf("refs do not contain %s", want)
		}
	}
	// The wrapper is not the resource, so no ref expects the resource type in it, and the resource
	// has no field for the request names.
	for _, ref := range refs {
		if ref.Name == "Spec" && ref.Want != "*RenameSpec" {
			t.Errorf("wrapper ref %s has type %s", ref.SDKName(), ref.Want)
		}
		if ref.Owner == "RenameThing" && (ref.Name == "SpecKind" || ref.Name == "SpecVersion") {
			t.Errorf("ref %s names the request path on the resource", ref.SDKName())
		}
	}
}

// When the Update body is the same inline schema as the Create body, openapi-generator makes one
// type with the name of the Create body, and the Update builder takes it.
func TestNamesGiveTheUpdateBodyTheCreateBodyName(t *testing.T) {
	spec := `openapi: 3.1.0
info: {title: t, version: 1.0.0}
paths:
  /v1/things:
    post:
      tags: [Things Service]
      operationId: ThingsService_CreateThing
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              properties:
                name: {type: string, x-coralogix-presence: true}
      responses: {'200': {description: ok, content: {application/json: {schema: {$ref: '#/components/schemas/Thing'}}}}}
  /v1/things/{id}:
    parameters: [{name: id, in: path, required: true, schema: {type: string}}]
    get:
      tags: [Things Service]
      operationId: ThingsService_GetThing
      responses: {'200': {description: ok, content: {application/json: {schema: {$ref: '#/components/schemas/Thing'}}}}}
    put:
      tags: [Things Service]
      operationId: ThingsService_UpdateThing
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              properties:
                name: {type: string, x-coralogix-presence: true}
      responses: {'200': {description: ok, content: {application/json: {schema: {$ref: '#/components/schemas/Thing'}}}}}
    delete:
      tags: [Things Service]
      operationId: ThingsService_DeleteThing
      responses: {'200': {description: ok, content: {application/json: {schema: {type: object}}}}}
components:
  schemas:
    Thing:
      type: object
      required: [id]
      properties:
        id: {type: string, readOnly: true}
        name: {type: string, x-coralogix-presence: true}
`
	doc, err := model.Load([]byte(spec))
	if err != nil {
		t.Fatal(err)
	}
	r, err := model.Build(doc, "Thing")
	if err != nil {
		t.Fatal(err)
	}
	refs, err := Names(r, "Things Service", "example.com/sdk", nil)
	if err != nil {
		t.Fatal(err)
	}
	var setter, field bool
	for _, ref := range refs {
		switch ref.SDKName() {
		case "ApiThingsServiceUpdateThingRequest.ThingsServiceCreateThingRequest":
			setter = true
		case "ThingsServiceCreateThingRequest.Name":
			field = true
		case "ApiThingsServiceUpdateThingRequest.ThingsServiceUpdateThingRequest", "ThingsServiceUpdateThingRequest":
			t.Errorf("ref %s names an Update body type that the SDK does not have", ref.SDKName())
		}
	}
	if !setter || !field {
		t.Fatalf("refs do not set the Update body with the Create body type: setter %t, field %t", setter, field)
	}
}

// A ByType method whose name belongs to a method of another type is found by its type.
func TestIssuesFindsAByTypeMethodWhoseNameHasAnotherType(t *testing.T) {
	const pkg = "example.com/fakesdk/widgets"
	refs := []Ref{
		{Path: "resource", Kind: KindPackage, Pkg: pkg, Name: "widgets"},
		{Path: "client", Kind: KindMethod, Pkg: pkg, Owner: "WidgetsAPIService", Name: "Gadgets", Want: "func() bool", ByType: true},
	}
	loaded, err := Load(refs, filepath.Join("testdata", "fakesdk"))
	if err != nil {
		t.Fatal(err)
	}
	if report := Issues(refs, loaded, func(Ref) Source { return Source{Module: "example.com/fakesdk"} }); len(report) != 0 {
		t.Fatalf("report = %v", report)
	}
	if refs[1].Name != "RestGadgets" {
		t.Fatalf("ref name = %s, want RestGadgets", refs[1].Name)
	}
}

// A collection resource sets the Get query with a builder method, and a written item that holds its
// fields in a wrapper names them on the wrapper component.
func TestNamesForAnOverwriteResource(t *testing.T) {
	raw, err := os.ReadFile("../model/testdata/overwritething.yaml")
	if err != nil {
		t.Fatal(err)
	}
	doc, err := model.Load(raw)
	if err != nil {
		t.Fatal(err)
	}
	write := "ThingsService_OverwriteKindAThings"
	policy := model.Policy{
		Existing: true, Overwrite: true, GetQuery: map[string]string{"kind": "THING_KIND_A"},
		Operations:      model.OperationIDs{Create: write, Get: "ThingsService_ListThings", Update: write, Delete: write},
		RequestWrappers: map[string]string{"Thing": "thing"},
		RequestNames:    map[string]string{"Thing.enabled": "disabled"},
		RequestNegated:  []string{"Thing.enabled"},
	}
	r, err := model.BuildWithPolicy(doc, "ThingList", model.OperationIDs{}, policy)
	if err != nil {
		t.Fatal(err)
	}
	tag, err := ResourceTag(doc, r)
	if err != nil {
		t.Fatal(err)
	}
	refs, err := Names(r, tag, "example.com/sdk", nil)
	if err != nil {
		t.Fatal(err)
	}
	has := func(path string, kind Kind, want string) {
		t.Helper()
		for _, ref := range refs {
			if ref.Path == path && ref.Kind == kind && strings.TrimSpace(ref.SDKName()+" "+ref.Want) == want {
				return
			}
		}
		t.Errorf("no %s ref %q at %s", kind, want, path)
	}
	has("get.query.kind", KindMethod, "ApiThingsServiceListThingsRequest.Kind func(kind ThingKind) ApiThingsServiceListThingsRequest")
	has("get.query.kind.THING_KIND_A", KindConst, "THINGKIND_THING_KIND_A")
	has("delete.body", KindType, "OverwriteKindAThingsRequest")
	has("create.body.things.name.via.thing", KindField, "KindAThingWrite.Thing *ThingCore")
	has("create.body.things.name", KindField, "ThingCore.Name *string")
	has("create.body.things.enabled", KindField, "ThingCore.Disabled *bool")
	has("create.body.things.aRules", KindField, "KindAThingWrite.ARules *Rules")
	// The response keeps its own names.
	has("fields.things.enabled", KindField, "Thing.Enabled *bool")
	for _, ref := range refs {
		if strings.HasPrefix(ref.Path, "fields.") && strings.Contains(ref.Path, ".via.") {
			t.Errorf("response ref %s goes through a request wrapper", ref.Path)
		}
	}
}

// A write body without a wrapper can hold the list of the resource under another name. The SDK
// names of that list are those of the request path.
func TestNamesForARenamedFieldOfAnUnwrappedBody(t *testing.T) {
	raw, err := os.ReadFile("../model/testdata/overwritething.yaml")
	if err != nil {
		t.Fatal(err)
	}
	spec := strings.Replace(string(raw), "required: [things]\n              properties:\n                things:", "required: [thingWrites]\n              properties:\n                thingWrites:", 1)
	if spec == string(raw) {
		t.Fatal("the spec has no write body property things")
	}
	doc, err := model.Load([]byte(spec))
	if err != nil {
		t.Fatal(err)
	}
	write := "ThingsService_OverwriteKindAThings"
	policy := model.Policy{
		Existing: true, Overwrite: true, GetQuery: map[string]string{"kind": "THING_KIND_A"},
		Operations:      model.OperationIDs{Create: write, Get: "ThingsService_ListThings", Update: write, Delete: write},
		RequestWrappers: map[string]string{"Thing": "thing"},
		RequestNames:    map[string]string{"ThingList.things": "thingWrites", "Thing.enabled": "disabled"},
		RequestNegated:  []string{"Thing.enabled"},
	}
	r, err := model.BuildWithPolicy(doc, "ThingList", model.OperationIDs{}, policy)
	if err != nil {
		t.Fatal(err)
	}
	tag, err := ResourceTag(doc, r)
	if err != nil {
		t.Fatal(err)
	}
	refs, err := Names(r, tag, "example.com/sdk", nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, path := range []string{"create.body.things", "update.body.things"} {
		found := false
		for _, ref := range refs {
			if ref.Path != path || ref.Kind != KindField {
				continue
			}
			found = true
			if got := ref.SDKName(); got != "OverwriteKindAThingsRequest.ThingWrites" {
				t.Errorf("%s = %s, want OverwriteKindAThingsRequest.ThingWrites", path, got)
			}
		}
		if !found {
			t.Errorf("no field ref at %s", path)
		}
	}
}

// A write that returns the resource has the resource fields on its response type, which a tool
// copies into the resource.
func TestNamesForAWriteThatReturnsTheResource(t *testing.T) {
	raw, err := os.ReadFile("../model/testdata/overwritething.yaml")
	if err != nil {
		t.Fatal(err)
	}
	spec := strings.Replace(string(raw), "    OverwriteResponse:\n      type: object\n      properties:\n        created:", "    OverwriteResponse:\n      type: object\n      properties:\n        things:", 1)
	doc, err := model.Load([]byte(spec))
	if err != nil {
		t.Fatal(err)
	}
	write := "ThingsService_OverwriteKindAThings"
	policy := model.Policy{
		Existing: true, Overwrite: true, WriteReturnsResource: true, GetQuery: map[string]string{"kind": "THING_KIND_A"},
		Operations:      model.OperationIDs{Create: write, Get: "ThingsService_ListThings", Update: write, Delete: write},
		RequestWrappers: map[string]string{"Thing": "thing"},
		RequestNames:    map[string]string{"Thing.enabled": "disabled"},
		RequestNegated:  []string{"Thing.enabled"},
	}
	r, err := model.BuildWithPolicy(doc, "ThingList", model.OperationIDs{}, policy)
	if err != nil {
		t.Fatal(err)
	}
	tag, err := ResourceTag(doc, r)
	if err != nil {
		t.Fatal(err)
	}
	refs, err := Names(r, tag, "example.com/sdk", nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, path := range []string{"create.response.things", "update.response.things"} {
		if !slices.ContainsFunc(refs, func(ref Ref) bool {
			return ref.Path == path && ref.Kind == KindField && ref.SDKName()+" "+ref.Want == "OverwriteResponse.Things []Thing"
		}) {
			t.Errorf("no field ref OverwriteResponse.Things []Thing at %s", path)
		}
	}
	if slices.ContainsFunc(refs, func(ref Ref) bool { return ref.Path == "delete.response.things" }) {
		t.Error("Delete has a ref for the write response: its response is not read")
	}
}

// A Create response that holds the wrapper of the id, as the id field of the resource does, needs
// the SDK type of that wrapper, not a string.
func TestNamesWrappedCreateResponseID(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("..", "model", "testdata", "wrappedidthing.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	doc, err := model.Load(data)
	if err != nil {
		t.Fatal(err)
	}
	r, err := model.BuildWithPolicy(doc, "GetThingResponse", model.OperationIDs{
		Create: "WrappedThingsService_CreateThing",
		Get:    "WrappedThingsService_GetThing",
		Update: "WrappedThingsService_UpdateThing",
		Delete: "WrappedThingsService_DeleteThing",
	}, model.Policy{Existing: true, IDProperty: "thingId", CreateResponseOnlyIDProperty: "thingId", Unwrap: []string{"ThingId"}})
	if err != nil {
		t.Fatal(err)
	}
	refs, err := Names(r, "Wrapped Things Service", "example.com/sdk", nil)
	if err != nil {
		t.Fatal(err)
	}
	i := slices.IndexFunc(refs, func(ref Ref) bool { return ref.Path == "create.response.thingId" })
	if i < 0 {
		t.Fatal("no ref for create.response.thingId")
	}
	if got := refs[i]; got.Want != "*ThingId" || got.WantValue != "ThingId" {
		t.Fatalf("create.response.thingId wants %s / %s, want *ThingId / ThingId", got.Want, got.WantValue)
	}
}

// A field that Update holds under another name is a field of the Update body of that name, and the
// Create body keeps the field name.
func TestNamesForAnUpdatePathThatDiffersFromCreate(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("..", "model", "testdata", "keything.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	spec := string(data)
	for _, r := range [][2]string{
		{"required: [name, keyPermissions]", "required: [name]"},
		{"                keyPermissions: {$ref: '#/components/schemas/CreateKeyThingRequest.KeyPermissions'}\n", ""},
		{"required: [id, name, enabled, hashed, keyPermissions]", "required: [id, name, enabled, hashed]"},
		{"        keyPermissions: {$ref: '#/components/schemas/KeyThing.KeyPermissions'}\n", ""},
	} {
		if !strings.Contains(spec, r[0]) {
			t.Fatalf("the contract has no %q", r[0])
		}
		spec = strings.ReplaceAll(spec, r[0], r[1])
	}
	doc, err := model.Load([]byte(spec))
	if err != nil {
		t.Fatal(err)
	}
	policy := model.Policy{
		Existing: true, IDProperty: "id", CreateResponseOnlyIDProperty: "keyId",
		UpdateNames:    map[string]string{"KeyThing.name": "newName", "KeyThing.enabled": "isEnabled"},
		UpdateDefaults: map[string]string{"KeyThing.enabled": "true"},
		Skip:           []string{"KeyThing.presets", "KeyThing.permissions"},
		Operations: model.OperationIDs{
			Create: "KeyThingsService_CreateKeyThing", Get: "KeyThingsService_GetKeyThing",
			Update: "KeyThingsService_UpdateKeyThing", Delete: "KeyThingsService_DeleteKeyThing",
		},
	}
	r, err := model.BuildWithPolicy(doc, "KeyThing", model.OperationIDs{}, policy)
	if err != nil {
		t.Fatal(err)
	}
	tag, err := ResourceTag(doc, r)
	if err != nil {
		t.Fatal(err)
	}
	refs, err := Names(r, tag, "example.com/sdk", nil)
	if err != nil {
		t.Fatal(err)
	}
	got := map[string]string{}
	for _, ref := range refs {
		if ref.Kind == KindField && strings.HasSuffix(ref.Path, ".body.name") || ref.Path == "update.body.enabled" || ref.Path == "create.body.enabled" {
			got[ref.Path] = ref.SDKName() + " " + ref.Want
		}
	}
	for path, want := range map[string]string{
		"create.body.name":    "CreateKeyThingRequest.Name *string",
		"update.body.name":    "UpdateKeyThingRequest.NewName *string",
		"update.body.enabled": "UpdateKeyThingRequest.IsEnabled *bool",
	} {
		if got[path] != want {
			t.Errorf("%s = %q, want %q", path, got[path], want)
		}
	}
	if _, ok := got["create.body.enabled"]; ok {
		t.Error("the Create body has a ref for enabled, which only Update sends")
	}
}

// A wrapper that only Update holds is a type and a field of the Update body, and the Create body
// and the response hold the list itself.
func TestNamesForAWrapperThatOnlyUpdateHolds(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("..", "model", "testdata", "keything.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	spec := string(data)
	for _, r := range [][2]string{
		{"required: [name, keyPermissions]", "required: [name]"},
		{"                keyPermissions: {$ref: '#/components/schemas/CreateKeyThingRequest.KeyPermissions'}\n",
			"                permissions: {type: array, items: {type: string, minLength: 1}}\n"},
		{"required: [id, name, enabled, hashed, keyPermissions]", "required: [id, name, enabled, hashed, permissions]"},
		{"        keyPermissions: {$ref: '#/components/schemas/KeyThing.KeyPermissions'}\n",
			"        permissions: {type: array, items: {type: string, minLength: 1}}\n"},
	} {
		if !strings.Contains(spec, r[0]) {
			t.Fatalf("the contract has no %q", r[0])
		}
		spec = strings.ReplaceAll(spec, r[0], r[1])
	}
	doc, err := model.Load([]byte(spec))
	if err != nil {
		t.Fatal(err)
	}
	policy := model.Policy{
		Existing: true, IDProperty: "id", CreateResponseOnlyIDProperty: "keyId",
		UpdateNames:    map[string]string{"KeyThing.name": "newName", "KeyThing.enabled": "isEnabled"},
		UpdateDefaults: map[string]string{"KeyThing.enabled": "true"},
		Skip:           []string{"KeyThing.presets"}, Unwrap: []string{"KeyThingPermissions"}, Released: []string{"KeyThing.permissions"},
		Operations: model.OperationIDs{
			Create: "KeyThingsService_CreateKeyThing", Get: "KeyThingsService_GetKeyThing",
			Update: "KeyThingsService_UpdateKeyThing", Delete: "KeyThingsService_DeleteKeyThing",
		},
	}
	r, err := model.BuildWithPolicy(doc, "KeyThing", model.OperationIDs{}, policy)
	if err != nil {
		t.Fatal(err)
	}
	tag, err := ResourceTag(doc, r)
	if err != nil {
		t.Fatal(err)
	}
	refs, err := Names(r, tag, "example.com/sdk", nil)
	if err != nil {
		t.Fatal(err)
	}
	got := map[string]string{}
	for _, ref := range refs {
		got[ref.Path+" "+string(ref.Kind)] = ref.SDKName() + " " + ref.Want
	}
	for path, want := range map[string]string{
		"update.body.permissions field":             "UpdateKeyThingRequest.Permissions *KeyThingPermissions",
		"update.body.permissions.permissions field": "KeyThingPermissions.Permissions []string",
		"create.body.permissions field":             "CreateKeyThingRequest.Permissions []string",
	} {
		if got[path] != want {
			t.Errorf("%s = %q, want %q", path, got[path], want)
		}
	}
	for path, ref := range got {
		if strings.Contains(ref, "*.") || strings.HasPrefix(ref, ".") {
			t.Errorf("%s has a ref without a component: %q", path, ref)
		}
	}
}

// A list that the resource shows as the strings of one property of its items has the type of the
// items in the response, and the property is a field of that type. The requests hold the strings.
func TestNamesForAListShownAsItemNames(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("..", "model", "testdata", "keything.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	spec := string(data)
	for _, r := range [][2]string{
		{"required: [name, keyPermissions]", "required: [name]"},
		{"                keyPermissions: {$ref: '#/components/schemas/CreateKeyThingRequest.KeyPermissions'}\n",
			"                presets: {type: array, items: {type: string, minLength: 1}}\n                permissions: {type: array, items: {type: string, minLength: 1}}\n"},
		{"required: [id, name, enabled, hashed, keyPermissions]", "required: [id, name, enabled, hashed, presets, permissions]"},
		{"        keyPermissions: {$ref: '#/components/schemas/KeyThing.KeyPermissions'}\n",
			"        presets: {type: array, items: {$ref: '#/components/schemas/PresetThing'}}\n        permissions: {type: array, items: {type: string, minLength: 1}}\n"},
	} {
		if !strings.Contains(spec, r[0]) {
			t.Fatalf("the contract has no %q", r[0])
		}
		spec = strings.ReplaceAll(spec, r[0], r[1])
	}
	doc, err := model.Load([]byte(spec))
	if err != nil {
		t.Fatal(err)
	}
	policy := model.Policy{
		Existing: true, IDProperty: "id", CreateResponseOnlyIDProperty: "keyId",
		UpdateNames:    map[string]string{"KeyThing.name": "newName", "KeyThing.enabled": "isEnabled"},
		UpdateDefaults: map[string]string{"KeyThing.enabled": "true"},
		Unwrap:         []string{"KeyThingPresets", "KeyThingPermissions"}, ItemNames: map[string]string{"KeyThing.presets": "name"},
		Released: []string{"KeyThing.presets", "KeyThing.permissions"},
		Operations: model.OperationIDs{
			Create: "KeyThingsService_CreateKeyThing", Get: "KeyThingsService_GetKeyThing",
			Update: "KeyThingsService_UpdateKeyThing", Delete: "KeyThingsService_DeleteKeyThing",
		},
	}
	r, err := model.BuildWithPolicy(doc, "KeyThing", model.OperationIDs{}, policy)
	if err != nil {
		t.Fatal(err)
	}
	tag, err := ResourceTag(doc, r)
	if err != nil {
		t.Fatal(err)
	}
	refs, err := Names(r, tag, "example.com/sdk", nil)
	if err != nil {
		t.Fatal(err)
	}
	got := map[string]string{}
	for _, ref := range refs {
		got[ref.Path+" "+string(ref.Kind)] = ref.SDKName() + " " + ref.Want
	}
	for path, want := range map[string]string{
		"update.body.presets field":         "UpdateKeyThingRequest.Presets *KeyThingPresets",
		"update.body.presets.presets field": "KeyThingPresets.Presets []string",
		"create.body.presets field":         "CreateKeyThingRequest.Presets []string",
	} {
		if got[path] != want {
			t.Errorf("%s = %q, want %q", path, got[path], want)
		}
	}
	for path, want := range map[string]string{
		"fields.presets field":      "KeyThing.Presets []PresetThing",
		"fields.presets.name field": "PresetThing.Name *string",
	} {
		if got[path] != want {
			t.Errorf("%s = %q, want %q", path, got[path], want)
		}
	}
}

// An Update by POST has no update mask: the names hold no mask method for it, and building them
// must not read the empty mask name.
func TestNamesForAnUpdateByPOSTHaveNoMask(t *testing.T) {
	raw, err := os.ReadFile("../model/testdata/methodthing.yaml")
	if err != nil {
		t.Fatal(err)
	}
	doc, err := model.Load(raw)
	if err != nil {
		t.Fatal(err)
	}
	policy := model.Policy{
		Existing: true, CreateMethod: "PUT", UpdateMethod: "POST",
		IDProperty: "thingId", CreateResponseOnlyIDProperty: "id",
		Skip:   []string{"MethodThing.parentId"},
		Unwrap: []string{"MethodThingPermissions"},
		UpdateNames: map[string]string{
			"MethodThing.name": "newName", "MethodThing.description": "newDescription", "MethodThing.permissions": "newPermissions",
		},
		Operations: model.OperationIDs{
			Create: "MethodThingsService_CreateMethodThing", Get: "MethodThingsService_GetMethodThing",
			Update: "MethodThingsService_UpdateMethodThing", Delete: "MethodThingsService_DeleteMethodThing",
		},
	}
	r, err := model.BuildWithPolicy(doc, "MethodThing", model.OperationIDs{}, policy)
	if err != nil {
		t.Fatal(err)
	}
	if r.UpdateMethod != "POST" || r.UpdateMask != "" {
		t.Fatalf("update method = %q, update mask = %q, want POST and no mask", r.UpdateMethod, r.UpdateMask)
	}
	tag, err := ResourceTag(doc, r)
	if err != nil {
		t.Fatal(err)
	}
	refs, err := Names(r, tag, "example.com/sdk", nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, ref := range refs {
		if ref.Path == "update.mask" {
			t.Errorf("ref %s: an Update by POST has no mask method", ref.Path)
		}
	}
}
