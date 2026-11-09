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
	"strings"
	"testing"

	"github.com/coralogix/coralogix-management-sdk/iac/issue"
)

// wrappedIDPolicy is the policy of a released resource like the teams: the id is inside the
// one-property object thingId, which the policy unwraps.
func wrappedIDPolicy() Policy {
	return Policy{
		Existing:                     true,
		IDProperty:                   "thingId",
		CreateResponseOnlyIDProperty: "thingId",
		Unwrap:                       []string{"ThingId"},
	}
}

var wrappedIDOps = OperationIDs{
	Create: "WrappedThingsService_CreateThing",
	Get:    "WrappedThingsService_GetThing",
	Update: "WrappedThingsService_UpdateThing",
	Delete: "WrappedThingsService_DeleteThing",
}

func loadWrappedIDThing(t *testing.T) *DocumentAlias {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("testdata", "wrappedidthing.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	doc, err := Load(data)
	if err != nil {
		t.Fatal(err)
	}
	return doc
}

// An id inside a wrapper that the policy unwraps is the inner value: the id checks pass, and
// the id field records the wrapper, so a tool can read the id from it and send it as the path
// parameter.
func TestWrappedID(t *testing.T) {
	doc := loadWrappedIDThing(t)
	policy := wrappedIDPolicy()
	if report := ValidateWithPolicy(doc, "GetThingResponse", wrappedIDOps, policy); len(report) != 0 {
		t.Fatalf("validate: %v", report)
	}
	r, err := BuildWithPolicy(doc, "GetThingResponse", wrappedIDOps, policy)
	if err != nil {
		t.Fatal(err)
	}
	if r.ID() != "thingId" {
		t.Fatalf("id field = %q, want thingId", r.ID())
	}
	var id *ResourceField
	for _, f := range r.Fields {
		if f.Name == r.ID() {
			id = f
		}
	}
	if id == nil {
		t.Fatal("no id field")
	}
	if id.Type.Kind != Integer || id.Type.Format != "int64" || len(id.Type.Wrappers) != 1 || id.Type.Wrappers[0].Schema != "ThingId" || id.Type.Wrappers[0].Field != "id" {
		t.Fatalf("id type = %+v, want int64 inside ThingId.id", id.Type)
	}
	if got := r.Create.Response; got.ID != "thingId" || got.idType == nil || !sameIDWrapper(got.idType, id.Type) {
		t.Fatalf("create response = %+v, want the id in thingId.id", got)
	}
}

// BuildWithPolicy is exported, so a caller can skip Validate. A policy that skips the id field
// returns an error that names the id, not a panic.
func TestWrappedIDSkipped(t *testing.T) {
	doc := loadWrappedIDThing(t)
	policy := wrappedIDPolicy()
	policy.Skip = []string{"GetThingResponse.thingId"}
	if _, err := BuildWithPolicy(doc, "GetThingResponse", wrappedIDOps, policy); err == nil || !strings.Contains(err.Error(), `"thingId"`) {
		t.Fatalf("build err = %v, want the skipped id field rejected", err)
	}
}

// Without the unwrap line, the id is an object, which is not an id type. The check reports it
// and does not panic.
func TestWrappedIDNeedsUnwrap(t *testing.T) {
	doc := loadWrappedIDThing(t)
	policy := wrappedIDPolicy()
	policy.Unwrap = nil
	report := ValidateWithPolicy(doc, "GetThingResponse", wrappedIDOps, policy)
	if !reportHas(report, "RESOURCE_ID_TYPE_UNSUPPORTED") {
		t.Fatalf("report = %v, want RESOURCE_ID_TYPE_UNSUPPORTED", report)
	}
	if _, err := BuildWithPolicy(doc, "GetThingResponse", wrappedIDOps, policy); err == nil || !strings.Contains(err.Error(), "createResponseOnlyIDProperty") {
		t.Fatalf("build err = %v, want the create response id rejected", err)
	}
}

func reportHas(report issue.Report, code string) bool {
	for _, i := range report {
		if i.Code == code {
			return true
		}
	}
	return false
}
