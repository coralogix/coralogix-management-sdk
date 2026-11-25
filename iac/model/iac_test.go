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
	"reflect"
	"slices"
	"strings"
	"testing"

	"github.com/coralogix/coralogix-management-sdk/iac/issue"
)

func iacSpec(t *testing.T, replacements ...string) []byte {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("testdata", "iacthing.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	spec := string(data)
	for i := 0; i < len(replacements); i += 2 {
		if !strings.Contains(spec, replacements[i]) {
			t.Fatalf("the spec has no %q", replacements[i])
		}
		spec = strings.Replace(spec, replacements[i], replacements[i+1], 1)
	}
	return []byte(spec)
}

func iacReport(t *testing.T, spec []byte) issue.Report {
	t.Helper()
	doc, err := Load(spec)
	if err != nil {
		t.Fatal(err)
	}
	return ValidateWithPolicy(doc, "IacThing", OperationIDs{}, Policy{})
}

func iacResource(t *testing.T) *Resource {
	t.Helper()
	doc, err := Load(iacSpec(t))
	if err != nil {
		t.Fatal(err)
	}
	r, err := Build(doc, "IacThing")
	if err != nil {
		t.Fatal(err)
	}
	return r
}

func iacField(t *testing.T, r *Resource, name string) *ResourceField {
	t.Helper()
	i := slices.IndexFunc(r.Fields, func(f *ResourceField) bool { return f.Name == name })
	if i < 0 {
		t.Fatalf("no field %s", name)
	}
	return r.Fields[i]
}

func TestIaCHintsValidSpecHasNoIssues(t *testing.T) {
	if report := iacReport(t, iacSpec(t)); len(report) != 0 {
		t.Fatalf("report = %v, want none", report)
	}
}

func TestIaCHintsAreRead(t *testing.T) {
	r := iacResource(t)

	name := iacField(t, r, "name").IaC()
	if name == nil || !name.RequiredByAPI || !reflect.DeepEqual(name.Examples, []any{"alpha", "beta"}) {
		t.Fatalf("name hints = %+v, want required and two examples", name)
	}
	// The hint next to an enum $ref is read, and the type stays the enum.
	mode := iacField(t, r, "mode")
	if mode.Type.Kind != Enum || mode.IaC() == nil || !reflect.DeepEqual(mode.IaC().Examples, []any{"IAC_MODE_FAST", "IAC_MODE_SAFE"}) {
		t.Fatalf("mode = %+v hints %+v, want the enum with two examples", mode.Type.Kind, mode.IaC())
	}
	priority := iacField(t, r, "priority").IaC()
	want := []IaCRequires{{Field: "mode", Equals: "IAC_MODE_FAST", HasEquals: true}, {Field: "name"}}
	if priority == nil || !reflect.DeepEqual(priority.ConflictsWith, []string{"note"}) || !reflect.DeepEqual(priority.Requires, want) {
		t.Fatalf("priority hints = %+v, want conflicts_with note and two requires", priority)
	}
	// A map example is an object that holds the keys.
	colors := iacField(t, r, "colors").IaC()
	if colors == nil || len(colors.Examples) != 2 {
		t.Fatalf("colors hints = %+v, want two examples", colors)
	}
	if first, ok := colors.Examples[0].(map[string]any); !ok || first["cpu"] == nil {
		t.Fatalf("colors example = %#v, want an object with the key cpu", colors.Examples[0])
	}
	// A nested field carries hints too.
	search := iacField(t, r, "search")
	text := fieldNamed(search.Type.Fields, "text").IaC()
	limit := fieldNamed(search.Type.Fields, "limit").IaC()
	if text == nil || len(text.Examples) != 2 || limit == nil || len(limit.Examples) != 2 {
		t.Fatalf("search hints = %+v and %+v, want two examples each", text, limit)
	}
	// arm_fixed is on an arm, and the oneOf still builds.
	source := iacField(t, r, "source")
	if arm := fieldNamed(source.Type.Fields, "logs").IaC(); arm == nil || !arm.ArmFixed {
		t.Fatalf("logs hints = %+v, want arm_fixed", arm)
	}
	if arm := fieldNamed(source.Type.Fields, "metrics").IaC(); arm != nil {
		t.Fatalf("metrics hints = %+v, want none", arm)
	}
	// A hint only in a request is found through the request attributes.
	note := iacField(t, r, "note")
	if note.Get.IaC != nil || note.Create.IaC == nil || note.IaC() == nil || len(note.IaC().Examples) != 2 {
		t.Fatalf("note: get %+v create %+v, want the hint in Create only", note.Get.IaC, note.Create.IaC)
	}
	// A field without a hint has none.
	if r.Fields[0].IaC() != nil {
		t.Fatalf("%s has hints %+v, want none", r.Fields[0].Name, r.Fields[0].IaC())
	}
}

// A hint changes no schema: the resource of a spec without hints is the resource with hints,
// except for the hints themselves.
func TestIaCHintsChangeNoType(t *testing.T) {
	with := iacResource(t)
	doc, err := Load([]byte(stripIaC(string(iacSpec(t)))))
	if err != nil {
		t.Fatal(err)
	}
	without, err := Build(doc, "IacThing")
	if err != nil {
		t.Fatal(err)
	}
	if len(with.Fields) != len(without.Fields) {
		t.Fatalf("fields %d and %d, want the same", len(with.Fields), len(without.Fields))
	}
	for i, f := range with.Fields {
		clearIaC(f.Type)
		if !reflect.DeepEqual(f.Type, without.Fields[i].Type) || f.Behavior != without.Fields[i].Behavior {
			t.Errorf("field %s differs with hints", f.Name)
		}
	}
}

// clearIaC removes the hints from the nested fields of t, so a type can be compared with one
// that was built without them.
func clearIaC(t *Type) {
	if t == nil {
		return
	}
	clearIaC(t.Elem)
	for _, f := range t.Fields {
		f.Attrs.IaC = nil
		if f.Create != nil {
			f.Create.IaC = nil
		}
		if f.Update != nil {
			f.Update.IaC = nil
		}
		clearIaC(f.Type)
	}
}

// stripIaC renames the extension, so the reader ignores it as a foreign x- key.
func stripIaC(spec string) string {
	return strings.ReplaceAll(spec, "x-coralogix-iac", "x-foreign-iac")
}

func TestIaCHintMistakesAreReported(t *testing.T) {
	tests := []struct {
		name     string
		old, new string
		want     string // a part of the message
		location string // a part of the location
	}{
		{"example is not an enum value", `examples: ["IAC_MODE_FAST", "IAC_MODE_SAFE"]`, `examples: ["IAC_MODE_FAST", "IAC_MODE_NOPE"]`, `"IAC_MODE_NOPE", which is not a value`, "IacThing.mode"},
		{"example is the enum zero value", `examples: ["IAC_MODE_FAST", "IAC_MODE_SAFE"]`, `examples: ["IAC_MODE_UNSPECIFIED"]`, "zero value", "IacThing.mode"},
		{"example is a number for a string", `examples: ["alpha", "beta"]`, `examples: [1, "beta"]`, "want a string", "IacThing.name"},
		{"example is too short", `examples: ["alpha", "beta"]`, `examples: ["", "beta"]`, "the minimum is 1", "IacThing.name"},
		{"integer example is a fraction", `examples: [10, 20]`, `examples: [1.5]`, "want an integer", "IacThing.search.limit"},
		{"integer example is over the maximum", `examples: [10, 20]`, `examples: [101]`, "the maximum is 100", "IacThing.search.limit"},
		{"map example is a list", `{"cpu": {"color": "#ff0000"}}`, `["cpu"]`, "want an object that maps", "IacThing.colors"},
		{"map value has an unknown key", `{"cpu": {"color": "#ff0000"}}`, `{"cpu": {"colour": "#ff0000"}}`, `the key "colour", which is not a field`, "IacThing.colors"},
		{"map value is not an object", `{"cpu": {"color": "#ff0000"}}`, `{"cpu": "red"}`, "want an object", "IacThing.colors"},
		{"map value misses a required field", `{"cpu": {"color": "#ff0000"}}`, `{"cpu": {}}`, `does not set the required field "color"`, "IacThing.colors"},
		{"conflicts_with names a stranger", `conflictsWith: [note]`, `conflictsWith: [nope]`, `"nope" is not a sibling field`, "IacThing.priority"},
		{"conflicts_with names the field", `conflictsWith: [note]`, `conflictsWith: [priority]`, "names the field itself", "IacThing.priority"},
		{"requires names a stranger", `- {field: name}`, `- {field: nope}`, `"nope" is not a sibling field`, "IacThing.priority"},
		{"requires equals is not an enum value", `equals: "IAC_MODE_FAST"`, `equals: "IAC_MODE_NOPE"`, "not a value of the enum", "IacThing.priority"},
		{"arm_fixed on a plain field", `conflictsWith: [note]`, "conflictsWith: [note]\n            armFixed: true", "not a oneOf arm", "IacThing.priority"},
		{"required_by_api is not a bool", `requiredByApi: true`, `requiredByApi: maybe`, `"maybe", want true or false`, "IacThing.name"},
		{"examples is not a list", `examples: ["alpha", "beta"]`, `examples: alpha`, "examples is not a list", "IacThing.name"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			report := iacReport(t, iacSpec(t, test.old, test.new))
			i := slices.IndexFunc(report, func(item issue.Issue) bool {
				return item.Code == "IAC_HINT_INVALID" && strings.Contains(item.Message, test.want) && strings.Contains(item.Location, test.location)
			})
			if i < 0 {
				t.Fatalf("report = %v, want IAC_HINT_INVALID at %s with %q", report, test.location, test.want)
			}
			if report[i].Remediation == "" {
				t.Errorf("issue %+v has no remediation", report[i])
			}
			for _, item := range report {
				if item.Code != "IAC_HINT_INVALID" {
					t.Errorf("report has %s (%s), want only IAC_HINT_INVALID", item.Code, item.Message)
				}
			}
		})
	}
}

func TestIaCHintsDifferBetweenResponseAndRequest(t *testing.T) {
	// The response has a hint on note, and the Create request has another one.
	spec := iacSpec(t, `        note: {type: string, x-coralogix-presence: true}
        colors:`, `        note:
          type: string
          x-coralogix-presence: true
          x-coralogix-iac: {examples: ["response note"]}
        colors:`)
	report := iacReport(t, spec)
	if i := slices.IndexFunc(report, func(item issue.Issue) bool {
		return item.Code == "IAC_HINT_INVALID" && strings.Contains(item.Message, "different iac knobs") && strings.HasSuffix(item.Location, "IacThing.note")
	}); i < 0 {
		t.Fatalf("report = %v, want IAC_HINT_INVALID for different knobs on note", report)
	}
}

func TestIaCArmFixedOnAnArmIsValid(t *testing.T) {
	if report := iacReport(t, iacSpec(t, `x-coralogix-iac: {armFixed: true}`, `x-coralogix-iac: {armFixed: true, examples: [{"query": "a"}]}`)); len(report) != 0 {
		t.Fatalf("report = %v, want none", report)
	}
}

func TestIaCUnknownKeyIsIgnored(t *testing.T) {
	// An older reader works with a newer fork: a key that it does not know is not an error.
	spec := iacSpec(t, `conflictsWith: [note]`, "conflictsWith: [note]\n            futureKnob: 7")
	if report := iacReport(t, spec); len(report) != 0 {
		t.Fatalf("report = %v, want none", report)
	}
}

func TestIaCSiblingExtensionsKeepOtherCoralogixKeysOutOfRef(t *testing.T) {
	// x-coralogix-iac next to a $ref is allowed. Another x-coralogix key next to it is not.
	spec := iacSpec(t, `x-coralogix-iac: {examples: ["IAC_MODE_FAST", "IAC_MODE_SAFE"]}`, `x-coralogix-iac: {examples: ["IAC_MODE_FAST"]}
          x-coralogix-collection: set`)
	if report := iacReport(t, spec); len(report) == 0 {
		t.Fatal("report is empty, want an issue for x-coralogix-collection next to a $ref")
	}
}
