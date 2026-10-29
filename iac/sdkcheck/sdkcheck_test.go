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
