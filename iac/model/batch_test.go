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
	}{
		"new resource":    {policy: Policy{GetByIDs: "ids", UpdateIDInBody: true}},
		"missing query":   {policy: Policy{Existing: true, GetByIDs: "missing", UpdateIDInBody: true}},
		"scalar query":    {old: "schema: {type: array, items: {type: string}}", replacement: "schema: {type: string}", policy: p},
		"numeric query":   {old: "schema: {type: array, items: {type: string}}", replacement: "schema: {type: array, items: {type: integer, format: int32}}", policy: p},
		"single response": {old: "things: {type: array, items: {$ref: '#/components/schemas/BatchThing'}}", replacement: "things: {$ref: '#/components/schemas/BatchThing'}", policy: p},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			source := strings.Replace(string(data), tt.old, tt.replacement, 1)
			doc, err := Load([]byte(source))
			if err != nil {
				t.Fatal(err)
			}
			if issues := ValidateWithPolicy(doc, "BatchThing", OperationIDs{}, tt.policy); len(issues) == 0 {
				t.Fatal("invalid batch contract was eligible")
			}
			if _, err := BuildWithPolicy(doc, "BatchThing", OperationIDs{}, tt.policy); err == nil {
				t.Fatal("invalid batch contract built")
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
	if issues := ValidateWithPolicy(doc, "BatchThing", OperationIDs{}, p); len(issues) == 0 {
		t.Fatal("optional Create with required Update must be rejected")
	}
	p.Required = []string{"BatchThing.name"}
	if issues := ValidateWithPolicy(doc, "BatchThing", OperationIDs{}, p); len(issues) != 0 {
		t.Fatal(issues)
	}
	if _, err := BuildWithPolicy(doc, "BatchThing", OperationIDs{}, p); err != nil {
		t.Fatal(err)
	}
}
