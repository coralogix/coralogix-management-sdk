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

// keyPolicy is the rule set of testdata/keything.yaml, as its behavior-overrides file states it.
func keyPolicy() Policy {
	return Policy{
		Existing:                     true,
		IDProperty:                   "id",
		CreateResponseOnlyIDProperty: "keyId",
		UpdateNames:                  map[string]string{"KeyThing.name": "newName", "KeyThing.enabled": "isEnabled"},
		Skip:                         []string{"KeyThing.presets", "KeyThing.permissions"},
		Operations: OperationIDs{
			Create: "KeyThingsService_CreateKeyThing", Get: "KeyThingsService_GetKeyThing",
			Update: "KeyThingsService_UpdateKeyThing", Delete: "KeyThingsService_DeleteKeyThing",
		},
	}
}

func keyDoc(t *testing.T) []byte {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("testdata", "keything.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	return data
}

// keyDocWithoutPermissions is the contract without the permission lists, whose shapes need more than updateAs.
func keyDocWithoutPermissions(t *testing.T) []byte {
	t.Helper()
	spec := string(keyDoc(t))
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
	return []byte(spec)
}

func keyCodes(t *testing.T, data []byte, p Policy) []string {
	t.Helper()
	doc, err := Load(data)
	if err != nil {
		t.Fatal(err)
	}
	return slices.Compact(reportCodes(ValidateWithPolicy(doc, "KeyThing", OperationIDs{}, p)))
}

func buildKey(t *testing.T, data []byte, p Policy) *Resource {
	t.Helper()
	doc, err := Load(data)
	if err != nil {
		t.Fatal(err)
	}
	r, err := BuildWithPolicy(doc, "KeyThing", OperationIDs{}, p)
	if err != nil {
		t.Fatal(err)
	}
	return r
}

// An updateAs line names the path of a value in the Update body only: Create keeps the field name,
// and Update sends the value under the other name. Without the line the field is Update-only for
// Get, so the API is refused.
func TestUpdateAsNamesTheUpdatePathOnly(t *testing.T) {
	if got := keyCodes(t, keyDocWithoutPermissions(t), keyPolicy()); len(got) != 0 {
		t.Fatalf("issues = %v, want none", got)
	}
	r := buildKey(t, keyDocWithoutPermissions(t), keyPolicy())
	name, enabled := renameField(t, r, "name"), renameField(t, r, "enabled")
	if name.RequestPath != "" || name.UpdateRequestPath != "newName" {
		t.Errorf("name paths = %q and %q, want the field name for Create and newName for Update", name.RequestPath, name.UpdateRequestPath)
	}
	if name.RequestPathOf(false) != "" || name.RequestPathOf(true) != "newName" {
		t.Errorf("name.RequestPathOf = %q (create), %q (update)", name.RequestPathOf(false), name.RequestPathOf(true))
	}
	if name.Behavior != Normal || name.Create == nil || name.Update == nil || !name.Create.Required || name.Update.Required {
		t.Errorf("name: behavior %s, create %+v, update %+v, want a managed field that Create requires and Update may leave out", name.Behavior, name.Create, name.Update)
	}
	if enabled.Create != nil || enabled.Update == nil || enabled.UpdateRequestPath != "isEnabled" {
		t.Errorf("enabled: create %+v, update %+v, path %q, want an Update-only field that Update sends as isEnabled", enabled.Create, enabled.Update, enabled.UpdateRequestPath)
	}
	p := keyPolicy()
	p.UpdateNames = nil
	if got := keyCodes(t, keyDocWithoutPermissions(t), p); !slices.Contains(got, "FIELD_LIFECYCLE_UNSUPPORTED") {
		t.Fatalf("issues = %v, want FIELD_LIFECYCLE_UNSUPPORTED without updateAs", got)
	}
}

// updateAs and requestAs can name different paths, and the line is refused when it names nothing.
func TestUpdateAsIssues(t *testing.T) {
	p := keyPolicy()
	p.UpdateNames = map[string]string{"KeyThing.name": "newName", "KeyThing.enabled": "isEnabled", "KeyThing.hashed": "hashed"}
	if got := keyCodes(t, keyDocWithoutPermissions(t), p); !slices.Contains(got, "OVERRIDE_UNUSED") {
		t.Errorf("issues = %v, want OVERRIDE_UNUSED for an updateAs that names the field", got)
	}
	p = keyPolicy()
	p.UpdateNames["KeyThing.name"] = "noSuchProperty"
	if got := keyCodes(t, keyDocWithoutPermissions(t), p); !slices.Contains(got, "REQUEST_NAME_NOT_FOUND") {
		t.Errorf("issues = %v, want REQUEST_NAME_NOT_FOUND for a path that the Update body does not have", got)
	}
	p = keyPolicy()
	p.UpdateNames["KeyThing.nothing"] = "x"
	if got := keyCodes(t, keyDocWithoutPermissions(t), p); !slices.Contains(got, "REQUEST_NAME_UNKNOWN_FIELD") {
		t.Errorf("issues = %v, want REQUEST_NAME_UNKNOWN_FIELD for a field that the resource does not have", got)
	}
}

// keyDocWithTopLevelLists is the contract with `presets` and `permissions` at the top level of
// Create and Get, as an inline line makes them. Get holds each preset as an object, Create holds
// the names, and Update holds each list in a one-property wrapper.
func keyDocWithTopLevelLists(t *testing.T) []byte {
	t.Helper()
	spec := string(keyDoc(t))
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
	return []byte(spec)
}

func keyWrapperPolicy() Policy {
	p := keyPolicy()
	p.Skip = nil
	p.Unwrap = []string{"KeyThingPresets", "KeyThingPermissions"}
	p.ItemNames = map[string]string{"KeyThing.presets": "name"}
	p.Sets = []string{"KeyThing.presets", "KeyThing.permissions"}
	// The released fields state what an omitted list does, so they need no declared server default.
	p.Released = []string{"KeyThing.presets", "KeyThing.permissions"}
	return p
}

// An unwrap line collapses a wrapper that only Update holds: Create and Get hold the list, Update
// holds it in a one-property object. The model records the wrapper on the type with the component
// of the Update request only, so the Create request has none.
func TestUnwrapCollapsesAWrapperThatOnlyUpdateHolds(t *testing.T) {
	data := keyDocWithTopLevelLists(t)
	if got := keyCodes(t, data, keyWrapperPolicy()); len(got) != 0 {
		t.Fatalf("issues = %v, want none", got)
	}
	r := buildKey(t, data, keyWrapperPolicy())
	f := renameField(t, r, "permissions")
	if f.Type.Kind != Set || f.Behavior != Normal || f.Create == nil || f.Update == nil {
		t.Fatalf("permissions: kind %s, behavior %s, create %+v, update %+v, want a managed set", f.Type.Kind, f.Behavior, f.Create, f.Update)
	}
	if got := f.Type.UpdateType().Kind; got != Set {
		t.Errorf("the Update request type is a %s, want the set of the collapsed wrapper", got)
	}
	if len(f.Type.Wrappers) != 0 {
		t.Fatalf("the response type has wrappers %+v, want none: only Update holds one", f.Type.Wrappers)
	}
	w := f.Type.RequestWrapper
	if w == nil || w.Schema != "" || w.CreateSchema != "" || w.UpdateSchema != "KeyThingPermissions" || w.Field != "permissions" || w.Update == nil || w.Create != nil {
		t.Fatalf("RequestWrapper = %+v, want the Update component KeyThingPermissions around the property permissions", w)
	}
	if got := f.Type.UpdateType().Wrappers; len(got) != 1 || got[0].Schema != "KeyThingPermissions" {
		t.Errorf("the Update request type has wrappers %+v, want KeyThingPermissions", got)
	}
	if got := f.Type.CreateType().Wrappers; len(got) != 0 {
		t.Errorf("the Create request type has wrappers %+v, want none", got)
	}
	if !r.UnwrapUsed["unwrap.KeyThingPermissions"] {
		t.Errorf("UnwrapUsed = %v, want the list entry to count as used", r.UnwrapUsed)
	}
}

// Without the unwrap line the Update type is an object and the response type a list.
func TestRequestOnlyWrapperNeedsTheUnwrapLine(t *testing.T) {
	p := keyWrapperPolicy()
	p.Unwrap = nil
	if got := keyCodes(t, keyDocWithTopLevelLists(t), p); !slices.Contains(got, "FIELD_TYPE_INCONSISTENT") {
		t.Fatalf("issues = %v, want FIELD_TYPE_INCONSISTENT without the unwrap line", got)
	}
}

// An itemName line shows a list of objects as the strings of one property: Get holds each preset as
// an object, and the requests send the names.
func TestItemNameShowsTheStringsOfAListOfObjects(t *testing.T) {
	data := keyDocWithTopLevelLists(t)
	if got := keyCodes(t, data, keyWrapperPolicy()); len(got) != 0 {
		t.Fatalf("issues = %v, want none", got)
	}
	r := buildKey(t, data, keyWrapperPolicy())
	f := renameField(t, r, "presets")
	if f.Type.Kind != Set || f.Type.Elem.Kind != String {
		t.Fatalf("presets type = %s of %s, want a set of strings", f.Type.Kind, f.Type.Elem.Kind)
	}
	if p := f.Type.Projected; p == nil || p.Schema != "PresetThing" || p.Property != "name" || !p.Required {
		t.Errorf("Projected = %+v, want the property name of PresetThing, required", p)
	}
	if f.Type.CreateType().Projected != nil || f.Type.UpdateType().Projected != nil {
		t.Error("a request type has a projection: the requests send the strings")
	}
	if w := f.Type.RequestWrapper; w == nil || w.UpdateSchema != "KeyThingPresets" {
		t.Errorf("RequestWrapper = %+v, want the Update wrapper KeyThingPresets next to the projection", w)
	}
}

// Without the line the request strings differ from the response objects.
func TestItemNameIsNeededForAListOfObjects(t *testing.T) {
	p := keyWrapperPolicy()
	p.ItemNames = nil
	if got := keyCodes(t, keyDocWithTopLevelLists(t), p); !slices.Contains(got, "FIELD_TYPE_INCONSISTENT") {
		t.Fatalf("issues = %v, want FIELD_TYPE_INCONSISTENT without itemName", got)
	}
}

func TestItemNameIssues(t *testing.T) {
	for name, tc := range map[string]struct {
		line, field string
	}{
		"no such property":         {"noSuch", "presets"},
		"not a list of objects":    {"name", "permissions"},
		"property is not a string": {"permissions", "presets"},
	} {
		t.Run(name, func(t *testing.T) {
			p := keyWrapperPolicy()
			p.ItemNames = map[string]string{"KeyThing." + tc.field: tc.line}
			if got := keyCodes(t, keyDocWithTopLevelLists(t), p); !slices.Contains(got, "ITEM_NAME_UNSUPPORTED") {
				t.Fatalf("issues = %v, want ITEM_NAME_UNSUPPORTED", got)
			}
		})
	}
}
