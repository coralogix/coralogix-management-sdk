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

package model_test

import (
	"os"
	"strings"
	"testing"

	"github.com/coralogix/coralogix-management-sdk/iac/issue"
	"github.com/coralogix/coralogix-management-sdk/iac/model"
	"github.com/coralogix/coralogix-management-sdk/iac/overrides"
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
)

// upsertSpec is a singleton that a GET reads and a POST on the same path sets. The request sends
// enabled at the top and a spec object; the response holds enabled in meta and spec as one arm of
// a oneOf.
func upsertSpec(t *testing.T) string {
	t.Helper()
	spec, err := os.ReadFile("testdata/upsert.yaml")
	if err != nil {
		t.Fatal(err)
	}
	return string(spec)
}

// upsertPolicy is the policy of the upsert: the API keys, plus the neutral fields that a tool's
// own keys fill (the upsert calls, the hoisted objects, and the fields the tool computes).
func upsertPolicy() model.Policy {
	p := overrides.Exceptions{
		Fields: map[string]map[string]overrides.Field{"Target": {"other": {Skip: true}, "token": {ReadOnly: true}}},
	}.Policy()
	p.Upsert = &model.Upsert{Get: "UpsertService_GetTarget", Set: "UpsertService_SetTarget"}
	p.Hoist = []string{"spec", "meta"}
	p.Computed = []string{"Target.enabled", "Target.zone"}
	p.Defaults = []string{"Target.enabled"}
	p.Released = append(p.Released, "Target.enabled")
	return p
}

func loadUpsert(t *testing.T, spec string) *v3.Document {
	t.Helper()
	doc, err := model.Load([]byte(spec))
	if err != nil {
		t.Fatal(err)
	}
	return doc
}

// An upsert has the Set call as Create and Update, a full replace, and no Delete. The fields of
// the hoisted objects are top-level fields that know the object that holds them in each place.
func TestUpsertModel(t *testing.T) {
	doc := loadUpsert(t, upsertSpec(t))
	if report := model.ValidateWithPolicy(doc, "Target", model.OperationIDs{}, upsertPolicy()); len(report) != 0 {
		t.Fatalf("report = %v, want none", report)
	}
	r, err := model.BuildWithPolicy(doc, "Target", model.OperationIDs{}, upsertPolicy())
	if err != nil {
		t.Fatal(err)
	}
	if !r.Upsert || !r.Replace || !r.Singleton || r.Create.OperationID != "UpsertService_SetTarget" || r.Update.OperationID != "UpsertService_SetTarget" || r.Delete.OperationID != "" {
		t.Fatalf("resource = %+v, want an upsert singleton with the Set call as Create and Update and no Delete", r)
	}
	fields := map[string]*model.ResourceField{}
	for _, f := range r.Fields {
		fields[f.Name] = f
	}
	if f := fields["enabled"]; f == nil || f.CreateParent != nil || f.GetParent == nil || f.GetParent.Field != "meta" || f.GetParent.Arm {
		t.Fatalf("enabled = %+v, want it at the top of the request and in meta in the response", f)
	}
	if f := fields["name"]; f == nil || f.CreateParent == nil || f.CreateParent.Field != "spec" || f.GetParent == nil || !f.GetParent.Arm {
		t.Fatalf("name = %+v, want it in spec in the request and in the spec arm in the response", f)
	}
	if f := fields["token"]; f == nil || f.Behavior != model.Computed || f.Create != nil {
		t.Fatalf("token = %+v, want a computed field that no request sends", f)
	}
	if _, ok := fields["other"]; ok {
		t.Fatal("the skipped arm is a field")
	}
	if parents := r.Parents(); len(parents) != 2 {
		t.Fatalf("parents = %+v, want meta and spec", parents)
	}
}

func TestUpsertAndHoistRejectIncompatibleContracts(t *testing.T) {
	tests := map[string]struct {
		spec           func(string) string
		policy         func(*model.Policy)
		ids            model.OperationIDs
		code, location string
	}{
		"other arm not skipped": {
			policy: func(p *model.Policy) { p.Skip = nil },
			code:   "HOIST_UNSUPPORTED", location: "components.schemas.Target",
		},
		"hoisted object optional in the request": {
			spec: func(s string) string {
				return strings.Replace(s, "required: [enabled, spec]", "required: [enabled]", 1)
			},
			code: "HOIST_UNSUPPORTED", location: "components.schemas.Target",
		},
		"set on another path": {
			spec: func(s string) string {
				return strings.Replace(s, "    post:\n", "  /upsert/v1/target/set:\n    post:\n", 1)
			},
			code: "UPSERT_SHAPE_UNSUPPORTED", location: "behavior-overrides.yaml:api.upsert.set",
		},
		"get with a path parameter": {
			spec: func(s string) string {
				s = strings.Replace(s, "  /upsert/v1/target:\n    get:\n", "  /upsert/v1/target/{id}:\n    get:\n      parameters:\n        - {name: id, in: path, required: true, schema: {type: string}}\n", 1)
				s = strings.Replace(s, "    post:\n", "  /upsert/v1/target:\n    post:\n", 1)
				return strings.Replace(s, "        meta: {$ref: '#/components/schemas/Meta'}\n", "        meta: {$ref: '#/components/schemas/Meta'}\n        id: {type: string}\n", 1)
			},
			code: "UPSERT_SHAPE_UNSUPPORTED", location: "behavior-overrides.yaml:api.upsert.get",
		},
		"operation flag": {
			ids:  model.OperationIDs{Get: "UpsertService_GetTarget"},
			code: "UPSERT_OPERATION_CONFLICT", location: "behavior-overrides.yaml:api.upsert",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			spec := upsertSpec(t)
			if test.spec != nil {
				changed := test.spec(spec)
				if changed == spec {
					t.Fatal("the test input did not change")
				}
				spec = changed
			}
			p := upsertPolicy()
			if test.policy != nil {
				test.policy(&p)
			}
			report := model.ValidateWithPolicy(loadUpsert(t, spec), "Target", test.ids, p)
			if !hasIssueAt(report, test.code, test.location) {
				t.Fatalf("report = %v, want %s at %s", report, test.code, test.location)
			}
		})
	}
}

// A hoisted field sits inside its object in the API, and an update mask names API fields, so
// hoist needs an Update that sends the whole resource.
func TestHoistNeedsAFullReplace(t *testing.T) {
	spec, err := os.ReadFile("testdata/openapi.yaml")
	if err != nil {
		t.Fatal(err)
	}
	p := model.Policy{Existing: true, Hoist: []string{"config"}}
	_, err = model.BuildWithPolicy(loadUpsert(t, string(spec)), "Thing", model.OperationIDs{}, p)
	if err == nil || !strings.Contains(err.Error(), "hoist needs an Update that sends the whole resource") {
		t.Fatalf("err = %v, want hoist to need a full replace", err)
	}
}

func hasIssueAt(report issue.Report, code, location string) bool {
	for _, i := range report {
		if i.Code == code && i.Location == location {
			return true
		}
	}
	return false
}
