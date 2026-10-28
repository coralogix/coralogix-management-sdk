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
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

// renamePolicy is the rule set of testdata/renamething.yaml, as its behavior-overrides file states it.
func renamePolicy() Policy {
	return Policy{
		Existing:                     true,
		RequestWrapper:               "spec",
		UpdateIDInBody:               true,
		UpdateIDOutsideWrapper:       true,
		IDProperty:                   "id",
		CustomMethods:                true,
		CreateResponseOnlyIDProperty: "thingId",
		RequestNames:                 map[string]string{"kind": "specKind", "version": "specVersion", "parameters": "holder.items"},
		Operations: OperationIDs{
			Create: "RenameThingsService_SaveRenameThing", Get: "RenameThingsService_GetRenameThing",
			Update: "RenameThingsService_UpdateRenameThing", Delete: "RenameThingsService_DeleteRenameThing",
		},
	}
}

// renameSpec returns testdata/renamething.yaml with each pair of old and new text replaced.
func renameSpec(t *testing.T, replacements ...string) []byte {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("testdata", "renamething.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	spec := string(data)
	for i := 0; i < len(replacements); i += 2 {
		if !strings.Contains(spec, replacements[i]) {
			t.Fatalf("the spec has no %q", replacements[i])
		}
		spec = strings.ReplaceAll(spec, replacements[i], replacements[i+1])
	}
	return []byte(spec)
}

func renameCodes(t *testing.T, spec []byte, p Policy) []string {
	t.Helper()
	doc, err := Load(spec)
	if err != nil {
		t.Fatal(err)
	}
	return slices.Compact(reportCodes(ValidateWithPolicy(doc, "RenameThing", OperationIDs{}, p)))
}

func buildRename(t *testing.T, p Policy) *Resource {
	t.Helper()
	doc, err := Load(renameSpec(t))
	if err != nil {
		t.Fatal(err)
	}
	r, err := BuildWithPolicy(doc, "RenameThing", OperationIDs{}, p)
	if err != nil {
		t.Fatal(err)
	}
	return r
}

func renameField(t *testing.T, r *Resource, name string) *ResourceField {
	t.Helper()
	for _, f := range r.Fields {
		if f.Name == name {
			return f
		}
	}
	t.Fatalf("no field %q in %v", name, r.Fields)
	return nil
}

// Mapped names and a dotted path make the request properties the resource fields: the contract
// checks pass, and the model records the request path.
func TestRequestNamesMatchFieldsByPath(t *testing.T) {
	if got := renameCodes(t, renameSpec(t), renamePolicy()); len(got) != 0 {
		t.Fatalf("issues = %v, want none", got)
	}
	r := buildRename(t, renamePolicy())
	if !r.UpdateIDOutsideWrapper {
		t.Error("UpdateIDOutsideWrapper = false")
	}
	if r.IDParam != "thing_id" || r.ID() != "id" {
		t.Errorf("id param %q, id field %q, want thing_id and id", r.IDParam, r.ID())
	}
	var names []string
	for _, f := range r.Fields {
		names = append(names, f.Name)
	}
	if !slices.Equal(names, []string{"id", "kind", "version", "parameters"}) {
		t.Fatalf("fields = %v, want the resource fields only: no request-only field", names)
	}
	for field, want := range map[string]string{"id": "", "kind": "specKind", "version": "specVersion", "parameters": "holder.items"} {
		if got := renameField(t, r, field).RequestPath; got != want {
			t.Errorf("%s.RequestPath = %q, want %q", field, got, want)
		}
	}
	kind, version, params, id := renameField(t, r, "kind"), renameField(t, r, "version"), renameField(t, r, "parameters"), renameField(t, r, "id")
	if kind.Behavior != Normal || kind.Create == nil || kind.Update == nil || !kind.Create.Required || !kind.Update.Required {
		t.Errorf("kind: behavior %s, create %+v, update %+v, want a managed field that both requests require", kind.Behavior, kind.Create, kind.Update)
	}
	if version.Create == nil || version.Create.Required || !version.Create.Presence {
		t.Errorf("version create attributes = %+v, want optional with presence", version.Create)
	}
	if params.Create == nil || params.Update == nil || !params.Create.Required {
		t.Errorf("parameters: create %+v, update %+v, want a managed list that Create requires (holder and items are required)", params.Create, params.Update)
	}
	if id.Behavior != Computed || id.Create != nil || id.Update == nil || !id.Update.Required {
		t.Errorf("id: behavior %s, create %+v, update %+v, want computed, sent by Update only", id.Behavior, id.Create, id.Update)
	}
}

// The model names the component that the request wrapper holds, which is not the resource, and the
// objects that a dotted request path goes through.
func TestRequestComponentsAreRecorded(t *testing.T) {
	r := buildRename(t, renamePolicy())
	if r.Create.RequestComponent != "RenameSpec" || r.Update.RequestComponent != "RenameSpec" {
		t.Errorf("request components = %q and %q, want RenameSpec", r.Create.RequestComponent, r.Update.RequestComponent)
	}
	if got := renameField(t, r, "parameters").RequestObjects; !slices.Equal(got, []string{"RenameHolder"}) {
		t.Errorf("parameters.RequestObjects = %v, want [RenameHolder]", got)
	}
	if got := renameField(t, r, "kind").RequestObjects; got != nil {
		t.Errorf("kind.RequestObjects = %v, want none for a path of one segment", got)
	}
	if got := buildThing(t, Policy{}).Create.RequestComponent; got != "" {
		t.Errorf("a resource without a request wrapper has RequestComponent %q", got)
	}
}

// Without the keys the same API is refused: nothing matches the names by accident, and nothing changes
// for a resource that does not use them.
func TestRequestNamesAbsentChangesNothing(t *testing.T) {
	p := renamePolicy()
	p.RequestNames, p.UpdateIDOutsideWrapper = nil, false
	got := renameCodes(t, renameSpec(t), p)
	if !slices.Contains(got, "FIELD_LIFECYCLE_UNSUPPORTED") {
		t.Fatalf("issues = %v, want FIELD_LIFECYCLE_UNSUPPORTED without the keys", got)
	}
	doc, err := Load(renameSpec(t))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := BuildWithPolicy(doc, "RenameThing", OperationIDs{}, p); err == nil {
		t.Fatal("Build accepted the API without the keys")
	}
	// Legacy and wrapper fixtures keep their results, and their fields have no request path.
	legacy := legacyCodes(t, legacyPolicy())
	if len(legacy) != 0 {
		t.Fatalf("legacy issues = %v, want none", legacy)
	}
	doc, err = Load(legacyDoc(t))
	if err != nil {
		t.Fatal(err)
	}
	r, err := BuildWithPolicy(doc, "LegacyThing", OperationIDs{}, legacyPolicy())
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range r.Fields {
		if f.RequestPath != "" {
			t.Errorf("%s.RequestPath = %q, want empty", f.Name, f.RequestPath)
		}
	}
	if r.UpdateIDOutsideWrapper {
		t.Error("UpdateIDOutsideWrapper is set without the key")
	}
}

// A skip line still works next to the keys: the skipped field is not managed.
func TestRequestNamesWithSkip(t *testing.T) {
	spec := renameSpec(t, "        version: {type: string, x-coralogix-presence: true}\n        parameters:", "        version: {type: string, x-coralogix-presence: true}\n        extra: {type: string, enum: [FIRST, SECOND]}\n        parameters:")
	p := renamePolicy()
	if got := renameCodes(t, spec, p); !slices.Contains(got, "ENUM_ZERO_INVALID") && !slices.Contains(got, "FIELD_LIFECYCLE_UNSUPPORTED") {
		t.Fatalf("issues = %v, want a problem with the extra field", got)
	}
	p.Skip = []string{"RenameThing.extra"}
	if got := renameCodes(t, spec, p); len(got) != 0 {
		t.Fatalf("issues with skip = %v, want none", got)
	}
}

func TestRequestNamesIssues(t *testing.T) {
	tests := map[string]struct {
		names map[string]string
		spec  []string
		want  string
	}{
		"unknown resource field":   {map[string]string{"kind": "specKind", "gone": "specVersion"}, nil, "REQUEST_NAME_UNKNOWN_FIELD"},
		"path not in a body":       {map[string]string{"kind": "nothing"}, nil, "REQUEST_NAME_NOT_FOUND"},
		"dotted path not found":    {map[string]string{"parameters": "holder.nothing"}, nil, "REQUEST_NAME_NOT_FOUND"},
		"segment not an object":    {map[string]string{"parameters": "specKind.items"}, nil, "REQUEST_NAME_NOT_OBJECT"},
		"two keys one path":        {map[string]string{"kind": "specKind", "version": "specKind"}, nil, "REQUEST_NAME_DUPLICATE"},
		"path inside another":      {map[string]string{"parameters": "holder", "version": "holder.items"}, nil, "REQUEST_NAME_DUPLICATE"},
		"name of the field itself": {map[string]string{"kind": "kind"}, nil, "OVERRIDE_UNUSED"},
		"request property also matched by name": {
			map[string]string{"kind": "specKind"},
			[]string{"        specKind: {type: string, minLength: 1}\n", "        specKind: {type: string, minLength: 1}\n        kind: {type: string, minLength: 1}\n"},
			"REQUEST_NAME_CONFLICT",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			p := renamePolicy()
			p.RequestNames = tc.names
			got := renameCodes(t, renameSpec(t, tc.spec...), p)
			if !slices.Contains(got, tc.want) {
				t.Fatalf("issues = %v, want %s", got, tc.want)
			}
		})
	}
}

// A path that only one body has is not an error: the field is then not in the other request.
func TestRequestNameInOneBodyOnly(t *testing.T) {
	spec := renameSpec(t,
		"        spec: {$ref: '#/components/schemas/RenameSpec'}\n    UpdateRenameThingResponse",
		"        spec: {$ref: '#/components/schemas/RenameSpecUpdate'}\n    UpdateRenameThingResponse",
		"    RenameHolder:",
		"    RenameSpecUpdate:\n      type: object\n      required: [specKind, holder]\n      properties:\n        specKind: {type: string, minLength: 1}\n        holder: {$ref: '#/components/schemas/RenameHolder'}\n    RenameHolder:")
	if got := renameCodes(t, spec, renamePolicy()); len(got) != 0 {
		t.Fatalf("issues = %v, want none: only Create has specVersion", got)
	}
}

// The leaf type is still checked against the Get type.
func TestRequestNamesLeafTypeIsChecked(t *testing.T) {
	spec := renameSpec(t, "        specKind: {type: string, minLength: 1}\n", "        specKind: {type: integer, format: int32}\n")
	if got := renameCodes(t, spec, renamePolicy()); !slices.Contains(got, "FIELD_TYPE_INCONSISTENT") {
		t.Fatalf("issues = %v, want FIELD_TYPE_INCONSISTENT", got)
	}
}

func TestUpdateIDOutsideWrapperIssues(t *testing.T) {
	p := renamePolicy()
	p.RequestWrapper = ""
	if got := renameCodes(t, renameSpec(t), p); !slices.Contains(got, "UPDATE_ID_OUTSIDE_WRAPPER_UNSUPPORTED") {
		t.Fatalf("no wrapper: issues = %v, want UPDATE_ID_OUTSIDE_WRAPPER_UNSUPPORTED", got)
	}
	// The id is not beside the wrapper.
	spec := renameSpec(t, "        id: {type: string, minLength: 1}\n        spec: {$ref: '#/components/schemas/RenameSpec'}\n    UpdateRenameThingResponse", "        spec: {$ref: '#/components/schemas/RenameSpec'}\n    UpdateRenameThingResponse")
	if got := renameCodes(t, spec, renamePolicy()); !slices.Contains(got, "UPDATE_ID_NOT_IN_BODY") {
		t.Fatalf("no id: issues = %v, want UPDATE_ID_NOT_IN_BODY", got)
	}
	// The id is inside the wrapper: the key is wrong.
	spec = renameSpec(t, "        specKind: {type: string, minLength: 1}\n", "        specKind: {type: string, minLength: 1}\n        id: {type: string, minLength: 1}\n")
	if got := renameCodes(t, spec, renamePolicy()); !slices.Contains(got, "UPDATE_ID_INSIDE_WRAPPER") {
		t.Fatalf("id inside: issues = %v, want UPDATE_ID_INSIDE_WRAPPER", got)
	}
	// Without the key, the id beside the wrapper is not a field of the resource.
	p = renamePolicy()
	p.UpdateIDOutsideWrapper = false
	if got := renameCodes(t, renameSpec(t), p); !slices.Contains(got, "UPDATE_ID_NOT_IN_BODY") {
		t.Fatalf("no key: issues = %v, want UPDATE_ID_NOT_IN_BODY", got)
	}
}

// api.idProperty names the resource field when the path parameter has another name.
func TestIDPropertyForAnotherPathParameter(t *testing.T) {
	p := renamePolicy()
	p.IDProperty, p.UpdateIDOutsideWrapper = "", false
	doc, err := Load(renameSpec(t))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := BuildWithPolicy(doc, "RenameThing", OperationIDs{}, p); err == nil || !strings.Contains(err.Error(), `no field "thing_id"`) {
		t.Fatalf("Build without idProperty: %v, want no field thing_id", err)
	}
	if got := renameCodes(t, renameSpec(t), renamePolicy()); len(got) != 0 {
		t.Fatalf("with idProperty: issues = %v, want none", got)
	}
}

// pinnedOpenAPI is the openapi.yaml of the SDK version that the released integrations resource
// was designed against. The test skips when the module cache does not have it.
const pinnedOpenAPI = "go/pkg/mod/github.com/coralogix/coralogix-management-sdk@v1.9.4-0.20261016082651-37b0b2075c43/openapi.yaml"

// The deployed integration has a request shape that differs from its read shape. With the keys,
// only the Parameter union remains.
func TestDeployedIntegrationRequestShape(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Skip(err)
	}
	data, err := os.ReadFile(filepath.Join(home, pinnedOpenAPI))
	if err != nil {
		t.Skip("the pinned SDK openapi.yaml is not in the module cache: ", err)
	}
	doc, err := Load(data)
	if err != nil {
		t.Fatal(err)
	}
	p := Policy{
		Existing:                     true,
		RequestWrapper:               "metadata",
		UpdateIDInBody:               true,
		UpdateIDOutsideWrapper:       true,
		IDProperty:                   "id",
		ServerSetID:                  true,
		CustomMethods:                true,
		CreateResponseOnlyIDProperty: "integrationId",
		RequestNames: map[string]string{
			"definitionKey": "integrationKey", "definitionVersion": "version", "parameters": "integrationParameters.parameters",
		},
		Operations: OperationIDs{
			Create: "IntegrationService_SaveIntegration", Get: "IntegrationService_GetDeployedIntegration",
			Update: "IntegrationService_UpdateIntegration", Delete: "IntegrationService_DeleteIntegration",
		},
		Skip: []string{"DeployedIntegrationInformation.integrationStatus"},
	}
	report := ValidateWithPolicy(doc, "DeployedIntegrationInformation", OperationIDs{}, p)
	if len(report) == 0 {
		t.Fatal("no issues: the Parameter union is not supported, so the report must name it")
	}
	for _, i := range report {
		union := strings.Contains(i.Message, "parameters[]") && strings.Contains(i.Message, "oneOf arm")
		if !union {
			t.Errorf("%s at %s: %s, want only the Parameter union issues", i.Code, i.Location, i.Message)
		}
	}
}
