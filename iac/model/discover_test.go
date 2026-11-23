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
	"strings"
	"testing"

	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
)

func loadSettingsThing(t *testing.T, edit func(string) string) *v3.Document {
	t.Helper()
	raw, err := os.ReadFile("testdata/settingsthing.yaml")
	if err != nil {
		t.Fatal(err)
	}
	if edit != nil {
		raw = []byte(edit(string(raw)))
	}
	doc, err := Load(raw)
	if err != nil {
		t.Fatal(err)
	}
	return doc
}

// extraGet adds a second GET of the same component, on another path, in the same service.
func extraGet(operationID string) func(string) string {
	return func(s string) string {
		return strings.Replace(s, "components:", "  /v1/settings/other:\n    get:\n      operationId: "+operationID+"\n      tags: [Settings Service]\n      responses:\n        '200': {description: ok, content: {application/json: {schema: {allOf: [{$ref: '#/components/schemas/SettingsThing'}], description: The thing.}}}}\ncomponents:", 1)
	}
}

// A service whose operations do not end in the component name is found with no explicit ids.
func TestServiceFallbackFindsOperations(t *testing.T) {
	doc := loadSettingsThing(t, nil)
	if report := Validate(doc, "SettingsThing", OperationIDs{}); len(report) != 0 {
		t.Fatalf("Validate: %v", report)
	}
	r, err := Build(doc, "SettingsThing")
	if err != nil {
		t.Fatal(err)
	}
	if r.Get.OperationID != "SettingsService_GetSettings" || r.Update.OperationID != "SettingsService_UpdateSettings" {
		t.Errorf("Get = %q, Update = %q", r.Get.OperationID, r.Update.OperationID)
	}
	if !r.Singleton || !r.AlwaysExists {
		t.Errorf("Singleton=%v AlwaysExists=%v, want both true", r.Singleton, r.AlwaysExists)
	}
}

// Two Get operations of one service that return the component are not a unique match.
func TestServiceFallbackAmbiguousStaysAnError(t *testing.T) {
	doc := loadSettingsThing(t, extraGet("SettingsService_GetOtherSettings"))
	report := Validate(doc, "SettingsThing", OperationIDs{})
	var found bool
	for _, i := range report {
		if i.Code == "OPERATION_AMBIGUOUS" && i.Location == "paths.get" {
			found = true
			for _, id := range []string{"SettingsService_GetSettings", "SettingsService_GetOtherSettings"} {
				if !strings.Contains(i.Message, id) {
					t.Errorf("message %q does not name %s", i.Message, id)
				}
			}
		}
	}
	if !found {
		t.Fatalf("Validate: %v, want OPERATION_AMBIGUOUS for Get", report)
	}
	if _, err := Build(doc, "SettingsThing"); err == nil {
		t.Fatal("Build found operations for an ambiguous service")
	}
}

// Two Delete operations that return the component do not leave an always-exists singleton.
func TestServiceFallbackAmbiguousDeleteIsNotImplied(t *testing.T) {
	doc := loadSettingsThing(t, func(s string) string {
		deleteOp := func(id, path string) string {
			return "  " + path + ":\n    delete:\n      operationId: " + id + "\n      tags: [Settings Service]\n      responses:\n        '200': {description: ok, content: {application/json: {schema: {allOf: [{$ref: '#/components/schemas/SettingsThing'}], description: The thing.}}}}\n"
		}
		return strings.Replace(s, "components:", deleteOp("SettingsService_DeleteSettings", "/v1/settings/a")+deleteOp("SettingsService_DeleteAllSettings", "/v1/settings/b")+"components:", 1)
	})
	var found bool
	for _, i := range Validate(doc, "SettingsThing", OperationIDs{}) {
		found = found || (i.Code == "OPERATION_AMBIGUOUS" && i.Location == "paths.delete")
	}
	if !found {
		t.Fatal("Validate found no OPERATION_AMBIGUOUS for Delete")
	}
	if _, err := Build(doc, "SettingsThing"); err == nil {
		t.Fatal("Build treated an ambiguous Delete as an always-exists singleton")
	}
}

// With no operation that returns the component, the error has the hint and the old message.
func TestServiceFallbackNoMatchHint(t *testing.T) {
	doc := loadSettingsThing(t, func(s string) string {
		return strings.ReplaceAll(s, "allOf: [{$ref: '#/components/schemas/SettingsThing'}], description: The thing.", "type: object, description: The thing.")
	})
	report := Validate(doc, "SettingsThing", OperationIDs{})
	for _, i := range report {
		if i.Code == "OPERATION_NOT_FOUND" && i.Location == "paths.get" {
			if !strings.HasPrefix(i.Remediation, "Fix the source API contract or select the exact operation ID.") || !strings.Contains(i.Remediation, "found none") {
				t.Errorf("Remediation = %q", i.Remediation)
			}
			return
		}
	}
	t.Fatalf("Validate: %v, want OPERATION_NOT_FOUND for Get", report)
}

// An explicit operationId wins over the fallback.
func TestServiceFallbackExplicitWins(t *testing.T) {
	doc := loadSettingsThing(t, extraGet("SettingsService_GetOtherSettings"))
	ids := OperationIDs{Get: "SettingsService_GetSettings"}
	if report := Validate(doc, "SettingsThing", ids); len(report) != 0 {
		t.Fatalf("Validate: %v", report)
	}
	r, err := BuildWithOperationIDs(doc, "SettingsThing", ids)
	if err != nil {
		t.Fatal(err)
	}
	if r.Get.OperationID != "SettingsService_GetSettings" {
		t.Errorf("Get = %q, want the explicit operation", r.Get.OperationID)
	}
}

// A resource that the suffix search resolves keeps its operations, even when its service has another
// Get that returns the component.
func TestServiceFallbackKeepsSuffixResult(t *testing.T) {
	doc := loadAlwaysThing(t, func(s string) string {
		return strings.Replace(s, "components:", "  /v1/always-things/other:\n    get:\n      operationId: AlwaysThingsService_GetOtherThing\n      tags: [Always Things Service]\n      responses:\n        '200': {description: ok, content: {application/json: {schema: {allOf: [{$ref: '#/components/schemas/AlwaysThing'}], description: The thing.}}}}\ncomponents:", 1)
	})
	if report := Validate(doc, "AlwaysThing", OperationIDs{}); len(report) != 0 {
		t.Fatalf("Validate: %v", report)
	}
	r, err := Build(doc, "AlwaysThing")
	if err != nil {
		t.Fatal(err)
	}
	if r.Get.OperationID != "AlwaysThingsService_GetAlwaysThing" || !r.AlwaysExists {
		t.Errorf("Get = %q AlwaysExists=%v, want the suffix match and an always-exists singleton", r.Get.OperationID, r.AlwaysExists)
	}
}
