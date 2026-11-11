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
	"testing"
)

// methodPolicy is the policy of a released resource whose Create is a PUT and whose Update is a
// POST: Create returns only the integer id, and the oneOf of the parent has one arm left.
func methodPolicy() Policy {
	return Policy{
		Existing:                     true,
		CreateMethod:                 "PUT",
		UpdateMethod:                 "POST",
		IDProperty:                   "thingId",
		CreateResponseOnlyIDProperty: "id",
		Skip:                         []string{"MethodThing.parentId"},
		Unwrap:                       []string{"MethodThingPermissions"},
		// Update names three fields differently, and the released resource requires them.
		UpdateNames: map[string]string{
			"MethodThing.name": "newName", "MethodThing.description": "newDescription", "MethodThing.permissions": "newPermissions",
		},
		Released: []string{
			"MethodThing.thingId", "MethodThing.name", "MethodThing.description", "MethodThing.parentName", "MethodThing.permissions",
		},
		ClientRequired: []string{"MethodThing.name", "MethodThing.description", "MethodThing.parentName", "MethodThing.permissions"},
	}
}

var methodOps = OperationIDs{
	Create: "MethodThingsService_CreateMethodThing",
	Get:    "MethodThingsService_GetMethodThing",
	Update: "MethodThingsService_UpdateMethodThing",
	Delete: "MethodThingsService_DeleteMethodThing",
}

func loadMethodThing(t *testing.T) *DocumentAlias {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("testdata", "methodthing.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	doc, err := Load(data)
	if err != nil {
		t.Fatal(err)
	}
	return doc
}

func methodCodes(t *testing.T, p Policy) []string {
	t.Helper()
	var codes []string
	for _, issue := range ValidateWithPolicy(loadMethodThing(t), "MethodThing", methodOps, p) {
		codes = append(codes, issue.Code)
	}
	return codes
}

// A PUT Create, a POST Update, an integer id in the Create response, and a oneOf whose other arm
// is skipped pass the checks, and the POST Update replaces the resource.
func TestMethodPolicyBuilds(t *testing.T) {
	if codes := methodCodes(t, methodPolicy()); len(codes) != 0 {
		t.Fatalf("issues = %v, want none", codes)
	}
	r, err := BuildWithPolicy(loadMethodThing(t), "MethodThing", methodOps, methodPolicy())
	if err != nil {
		t.Fatal(err)
	}
	if !r.Replace {
		t.Error("a POST Update that the policy states must replace the resource")
	}
	if len(r.Groups) != 0 {
		t.Errorf("groups = %v, want none: the oneOf has one arm left", r.Groups)
	}
}

// Without the method lines the model stays strict: a PUT Create and a POST Update are refused.
func TestMethodsAreRefusedWithoutThePolicy(t *testing.T) {
	p := methodPolicy()
	p.CreateMethod, p.UpdateMethod = "", ""
	codes := methodCodes(t, p)
	if n := slices.Index(codes, "OPERATION_METHOD_INCOMPATIBLE"); n < 0 {
		t.Fatalf("codes = %v, want OPERATION_METHOD_INCOMPATIBLE", codes)
	}
}

// A oneOf of the Create request that Get and Update lack is inconsistent while both arms are
// managed, and no group when the policy skips one.
func TestRootOneOfWithASkippedArm(t *testing.T) {
	p := methodPolicy()
	p.Skip = nil
	if codes := methodCodes(t, p); !slices.Contains(codes, "ROOT_ONEOF_LIFECYCLE_INCONSISTENT") {
		t.Fatalf("codes = %v, want ROOT_ONEOF_LIFECYCLE_INCONSISTENT", codes)
	}
}

// A Create response id that is an integer is refused when it is not the type of the id field.
func TestIntegerCreateResponseIDMustMatchTheIDField(t *testing.T) {
	r, err := BuildWithPolicy(loadMethodThing(t), "MethodThing", methodOps, methodPolicy())
	if err != nil {
		t.Fatal(err)
	}
	if r.Create.Response.ID != "id" || r.IDType.Format != "int64" {
		t.Fatalf("response id = %q, id type = %+v, want id as int64", r.Create.Response.ID, r.IDType)
	}
}
