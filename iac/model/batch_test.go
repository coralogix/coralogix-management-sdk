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
	"fmt"
	"os"
	"slices"
	"strings"
	"testing"
)

func TestBatchGet(t *testing.T) {
	data, err := os.ReadFile("testdata/batchthing.yaml")
	if err != nil {
		t.Fatal(err)
	}
	p := Policy{Existing: true, GetByIDs: "ids", UpdateIDInBody: true, ReadOnly: []string{"BatchThing.id"}}
	doc, err := Load(data)
	if err != nil {
		t.Fatal(err)
	}
	if issues := ValidateWithPolicy(doc, "BatchThing", OperationIDs{}, p); len(issues) != 0 {
		t.Fatal(issues)
	}
	r, err := BuildWithPolicy(doc, "BatchThing", OperationIDs{}, p)
	if err != nil {
		t.Fatal(err)
	}
	if r.Singleton || r.ID() != "id" || r.IDType.Kind != String || !r.Get.Response.Items || r.Get.Response.Field != "things" {
		t.Fatalf("wrong batch identity or response: %+v", r)
	}
	tests := map[string]struct {
		old, replacement string
		policy           Policy
		wantError        string
	}{
		"new resource": {
			old: "name: ids, in: query, required: true", replacement: "name: ids, in: query, required: false",
			policy: Policy{GetByIDs: "ids", UpdateIDInBody: true}, wantError: "api.getByIDs is only for existing resources",
		},
		"missing query": {
			old: "name: ids, in: query, required: true", replacement: "name: ids, in: query, required: false",
			policy: Policy{Existing: true, GetByIDs: "missing", UpdateIDInBody: true}, wantError: `api.getByIDs: no query parameter "missing"`,
		},
		"scalar query": {
			old: "schema: {type: array, items: {type: string}}", replacement: "schema: {type: string}",
			policy: p, wantError: "api.getByIDs needs an array query parameter",
		},
		"numeric query": {
			old: "schema: {type: array, items: {type: string}}", replacement: "schema: {type: array, items: {type: integer, format: int32}}",
			policy: p, wantError: "api.getByIDs needs string ids",
		},
		"single response": {
			old: "things: {type: array, items: {$ref: '#/components/schemas/BatchThing'}}", replacement: "things: {$ref: '#/components/schemas/BatchThing'}",
			policy: p, wantError: "Get: api.getByIDs needs a response array of BatchThing",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			source := strings.Replace(string(data), tt.old, tt.replacement, 1)
			doc, err := Load([]byte(source))
			if err != nil {
				t.Fatal(err)
			}
			issues := ValidateWithPolicy(doc, "BatchThing", OperationIDs{}, tt.policy)
			var diagnostics []string
			for _, issue := range issues {
				diagnostics = append(diagnostics, issue.Code+": "+issue.Message)
			}
			if !slices.Contains(diagnostics, "RESOURCE_LIFECYCLE_INCOMPLETE: "+tt.wantError) {
				t.Fatalf("want batch diagnostic %q, got %v", tt.wantError, issues)
			}
			if _, err := BuildWithPolicy(doc, "BatchThing", OperationIDs{}, tt.policy); err == nil || err.Error() != tt.wantError {
				t.Fatalf("Build error = %v, want %q", err, tt.wantError)
			}
		})
	}
}

func TestReleasedRequiredFieldCanBeOptionalOnCreate(t *testing.T) {
	data, err := os.ReadFile("testdata/batchthing.yaml")
	if err != nil {
		t.Fatal(err)
	}
	data = []byte(strings.Replace(string(data), "required: [name]", "required: []", 1))
	doc, err := Load(data)
	if err != nil {
		t.Fatal(err)
	}
	p := Policy{Existing: true, GetByIDs: "ids", UpdateIDInBody: true, ReadOnly: []string{"BatchThing.id"}}
	if codes := reportCodes(ValidateWithPolicy(doc, "BatchThing", OperationIDs{}, p)); !slices.Contains(codes, "FIELD_REQUIREDNESS_UNSUPPORTED") {
		t.Fatalf("codes %v do not contain FIELD_REQUIREDNESS_UNSUPPORTED", codes)
	}
	p.ClientRequired = []string{"BatchThing.name"}
	if issues := ValidateWithPolicy(doc, "BatchThing", OperationIDs{}, p); len(issues) != 0 {
		t.Fatal(issues)
	}
	if _, err := BuildWithPolicy(doc, "BatchThing", OperationIDs{}, p); err != nil {
		t.Fatal(err)
	}
}

func TestContractSchemasBatchResponse(t *testing.T) {
	data, err := os.ReadFile("testdata/batchthing.yaml")
	if err != nil {
		t.Fatal(err)
	}
	for _, array := range []bool{false, true} {
		for _, batch := range []bool{false, true} {
			t.Run(fmt.Sprintf("array=%t/batch=%t", array, batch), func(t *testing.T) {
				source := string(data)
				if !array {
					source = strings.Replace(source, "things: {type: array, items: {$ref: '#/components/schemas/BatchThing'}}", "things: {$ref: '#/components/schemas/BatchThing'}", 1)
				}
				doc, err := Load([]byte(source))
				if err != nil {
					t.Fatal(err)
				}
				p := Policy{Existing: true, GetByIDs: "ids", UpdateIDInBody: true, ReadOnly: []string{"BatchThing.id"}}
				ops, err := findOperations(doc, "BatchThing", OperationIDs{}, p)
				if err != nil {
					t.Fatal(err)
				}
				if !batch {
					p.GetByIDs = ""
				}
				_, _, get, issues, ok := contractSchemas(p, "BatchThing", ops)
				if len(issues) != 0 {
					t.Fatal(issues)
				}
				if want := !array || batch; ok != want {
					t.Fatalf("resource schema found = %t, want %t", ok, want)
				}
				if ok && propertyOf(get, "name") == nil {
					t.Fatal("resolved schema has no resource name field")
				}
			})
		}
	}
}
