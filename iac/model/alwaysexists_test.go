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
	"regexp"
	"strings"
	"testing"

	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
)

func loadAlwaysThing(t *testing.T, edit func(string) string) *v3.Document {
	t.Helper()
	raw, err := os.ReadFile("testdata/alwaysthing.yaml")
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

// A Get and an Update on one path, with no Create and no Delete, is a singleton that always exists.
func TestAlwaysExistsSingleton(t *testing.T) {
	doc := loadAlwaysThing(t, nil)
	if report := Validate(doc, "AlwaysThing", OperationIDs{}); len(report) != 0 {
		t.Fatalf("Validate: %v", report)
	}
	r, err := Build(doc, "AlwaysThing")
	if err != nil {
		t.Fatal(err)
	}
	if !r.Singleton || !r.AlwaysExists {
		t.Fatalf("Singleton=%v AlwaysExists=%v, want both true", r.Singleton, r.AlwaysExists)
	}
	if r.Create.OperationID != r.Update.OperationID || r.Create.Method != "PATCH" {
		t.Errorf("Create = %s %s, want the Update call", r.Create.Method, r.Create.OperationID)
	}
	if r.Delete.OperationID != "" {
		t.Errorf("Delete = %q, want no operation", r.Delete.OperationID)
	}
	if r.Update.Response.Schema != "AlwaysThing" || r.Create.Response.Schema != "AlwaysThing" {
		t.Errorf("Create response %q, Update response %q, want AlwaysThing", r.Create.Response.Schema, r.Update.Response.Schema)
	}
}

// A resource that has a Create or a Delete is no always-exists singleton, so a missing step is still
// an error.
func TestAlwaysExistsNeedsBothStepsMissing(t *testing.T) {
	post := func(s string) string {
		return strings.Replace(s, "    patch:", "    post:\n      operationId: AlwaysThingsService_CreateAlwaysThing\n      tags: [Always Things Service]\n      requestBody: {content: {application/json: {schema: {type: object, properties: {enabled: {type: boolean}}}}}}\n      responses:\n        '200': {description: ok, content: {application/json: {schema: {$ref: '#/components/schemas/AlwaysThing'}}}}\n    patch:", 1)
	}
	doc := loadAlwaysThing(t, post)
	if _, err := Build(doc, "AlwaysThing"); err == nil || !strings.Contains(err.Error(), "Delete") {
		t.Fatalf("Build with a Create and no Delete: %v, want a Delete error", err)
	}
	if report := Validate(doc, "AlwaysThing", OperationIDs{}); len(report) == 0 {
		t.Fatal("Validate with a Create and no Delete found no issue")
	}
}

// A Get that has a path parameter is no singleton, so it keeps the four-step rule.
func TestAlwaysExistsNeedsASingleton(t *testing.T) {
	doc := loadAlwaysThing(t, func(s string) string {
		return strings.Replace(s, "  /v1/always-things:\n", "  /v1/always-things/{id}:\n    parameters: [{name: id, in: path, required: true, schema: {type: string}}]\n", 1)
	})
	if _, err := Build(doc, "AlwaysThing"); err == nil {
		t.Fatal("Build of a Get with a path parameter and no Create found no error")
	}
}

// An explicit Create or Delete operation id that the contract does not have is a lookup error. It
// must not read as a contract with no such step.
func TestAlwaysExistsKeepsUnresolvedExplicitIDs(t *testing.T) {
	doc := loadAlwaysThing(t, nil)
	for _, ids := range []OperationIDs{{Delete: "DoesNotExist"}, {Create: "DoesNotExist"}} {
		if _, err := BuildWithOperationIDs(doc, "AlwaysThing", ids); err == nil {
			t.Errorf("Build with %+v found no error", ids)
		}
		if report := ValidateWithPolicy(doc, "AlwaysThing", ids, Policy{}); len(report) == 0 {
			t.Errorf("Validate with %+v found no issue", ids)
		}
	}
}

// A PUT replaces the resource, so an omitted field is cleared and not kept. A field that is optional
// in the request and required in Get still needs a declared default.
func TestAlwaysExistsPutKeepsTheDefaultRule(t *testing.T) {
	put := func(s string) string {
		s = strings.Replace(s, "    patch:", "    put:", 1)
		return regexp.MustCompile(`(?s)      parameters:\n        - name: update_mask.*?x-coralogix-presence: true\}\n`).ReplaceAllString(s, "")
	}
	patch := Validate(loadAlwaysThing(t, nil), "AlwaysThing", OperationIDs{})
	if len(patch) != 0 {
		t.Fatalf("PATCH: %v", patch)
	}
	report := Validate(loadAlwaysThing(t, put), "AlwaysThing", OperationIDs{})
	found := false
	for _, i := range report {
		found = found || i.Code == "FIELD_SERVER_DEFAULT_UNDECLARED"
	}
	if !found {
		t.Fatalf("PUT: want FIELD_SERVER_DEFAULT_UNDECLARED, got %v", report)
	}
}

// Omitting a top-level field of an always-exists singleton keeps the value that the API holds, so a
// field without presence gets no zero default. Zero is then a value that the API can hold.
func TestAlwaysExistsGetsNoZeroDefault(t *testing.T) {
	noPresence := func(s string) string {
		return strings.Replace(s, "enabled: {type: boolean, x-coralogix-presence: true}", "enabled: {type: boolean}", 1)
	}
	r, err := Build(loadAlwaysThing(t, noPresence), "AlwaysThing")
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range r.Fields {
		if f.Name != "enabled" {
			continue
		}
		if f.Create.Default != nil || f.Update.Default != nil || f.Get.Default != nil {
			t.Errorf("defaults create=%v update=%v get=%v, want none", f.Create.Default, f.Update.Default, f.Get.Default)
		}
		return
	}
	t.Fatal("no field enabled")
}
