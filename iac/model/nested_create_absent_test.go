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
	"slices"
	"strings"
	"testing"
)

// foldedThingSpec is a resource whose Create sends the limit flat as limitValue. Update and Get hold
// it as quota.limit. The quota component also has a field that Create cannot send (extra), set by
// the placeholder EXTRA_FIELD. It is the shape of a createAs line.
const foldedThingSpec = `openapi: 3.1.0
info: {title: t, version: 1.0.0}
paths:
  /v1/things:
    post:
      operationId: Svc_CreateThing
      requestBody: {required: true, content: {application/json: {schema: {$ref: '#/components/schemas/CreateThingRequest'}}}}
      responses: {'200': {description: ok, content: {application/json: {schema: {$ref: '#/components/schemas/Thing'}}}}}
  /v1/things/{id}:
    parameters: [{name: id, in: path, required: true, schema: {type: string}}]
    get:
      operationId: Svc_GetThing
      responses: {'200': {description: ok, content: {application/json: {schema: {$ref: '#/components/schemas/Thing'}}}}}
    put:
      operationId: Svc_UpdateThing
      requestBody: {required: true, content: {application/json: {schema: {$ref: '#/components/schemas/UpdateThingRequest'}}}}
      responses: {'200': {description: ok, content: {application/json: {schema: {$ref: '#/components/schemas/Thing'}}}}}
    delete:
      operationId: Svc_DeleteThing
      responses: {'200': {description: ok, content: {application/json: {schema: {type: object}}}}}
components:
  schemas:
    CreateThingRequest:
      type: object
      properties:
        name: {type: string}
        limitValue: {type: integer, format: int32}
    UpdateThingRequest:
      type: object
      properties:
        name: {type: string}
        quota: {$ref: '#/components/schemas/Quota'}
    Thing:
      type: object
      required: [id]
      properties:
        id: {type: string, readOnly: true}
        name: {type: string}
        quota: {$ref: '#/components/schemas/Quota'}
    Quota:
      type: object
      properties:
        limit: {type: integer, format: int32}
EXTRA_FIELD`

func foldedThingPolicy() Policy {
	return Policy{Existing: true, CreateAs: map[string]map[string]string{"Thing.quota": {"limit": "limitValue"}}}
}

func foldedThing(t *testing.T, extra string) (*Resource, error) {
	t.Helper()
	doc, err := Load([]byte(strings.Replace(foldedThingSpec, "EXTRA_FIELD", extra, 1)))
	if err != nil {
		t.Fatal(err)
	}
	return BuildWithPolicy(doc, "Thing", OperationIDs{}, foldedThingPolicy())
}

func quotaBehaviors(t *testing.T, r *Resource) map[string]Behavior {
	t.Helper()
	for _, f := range r.Fields {
		if f.Name != "quota" {
			continue
		}
		out := map[string]Behavior{}
		for _, nested := range f.Type.Fields {
			out[nested.Name] = nested.Behavior
		}
		return out
	}
	t.Fatal("no quota field")
	return nil
}

// A new field of an object that a createAs line folds is computed: Create sends only the folded
// property, so the user cannot set the new field at create.
func TestFoldedObjectNewFieldIsComputed(t *testing.T) {
	r, err := foldedThing(t, "")
	if err != nil {
		t.Fatal(err)
	}
	if got := quotaBehaviors(t, r); !slices.Equal([]Behavior{got["limit"]}, []Behavior{Normal}) || len(got) != 1 {
		t.Fatalf("behaviors without the new field = %v, want only limit as normal", got)
	}
	r, err = foldedThing(t, "        added: {type: string}\n")
	if err != nil {
		t.Fatalf("Build with the new field = %v, want no error", err)
	}
	if got := quotaBehaviors(t, r); got["limit"] != Normal || got["added"] != Computed {
		t.Fatalf("behaviors = %v, want limit normal and added computed", got)
	}
}

// An object that the server sets inside a folded object is one computed value: its own fields have no
// lifecycle, as in an object that only Get has.
func TestFoldedObjectNewObjectIsComputed(t *testing.T) {
	r, err := foldedThing(t, "        detail: {$ref: '#/components/schemas/QuotaDetail'}\n    QuotaDetail:\n      type: object\n      properties:\n        note: {type: string}\n")
	if err != nil {
		t.Fatalf("Build with the new object = %v, want no error", err)
	}
	for _, f := range r.Fields {
		if f.Name != "quota" {
			continue
		}
		detail := fieldNamed(f.Type.Fields, "detail")
		if detail.Behavior != Computed || fieldNamed(detail.Type.Fields, "note").Behavior != "" {
			t.Fatalf("detail = %q, note = %q, want computed and no lifecycle", detail.Behavior, fieldNamed(detail.Type.Fields, "note").Behavior)
		}
	}
}

// Without a createAs line, a field that Update and Get have and Create does not is still an error.
func TestUnfoldedObjectNewFieldIsStillRefused(t *testing.T) {
	doc, err := Load([]byte(strings.Replace(foldedThingSpec, "EXTRA_FIELD", "        added: {type: string}\n", 1)))
	if err != nil {
		t.Fatal(err)
	}
	policy := foldedThingPolicy()
	policy.CreateAs = nil
	var codes []string
	for _, item := range ValidateWithPolicy(doc, "Thing", OperationIDs{}, policy) {
		codes = append(codes, item.Code)
	}
	if !slices.Contains(codes, "FIELD_LIFECYCLE_UNSUPPORTED") {
		t.Fatalf("issues = %v, want FIELD_LIFECYCLE_UNSUPPORTED without a createAs line", codes)
	}
}

// TestClassifyFieldUnchangedResults pins the results that were valid before the folded case, and the
// parents that Update does not send.
func TestClassifyFieldUnchangedResults(t *testing.T) {
	attrs := &Attrs{}
	tests := []struct {
		name                   string
		create, update, folded bool
		inCreate, inUpdate     bool
		want                   Behavior
		wantErr                bool
	}{
		{"normal", true, true, false, true, true, Normal, false},
		{"create only", true, true, false, true, false, Immutable, false},
		{"get only", true, true, false, false, false, Computed, false},
		{"update only is refused", true, true, false, false, true, "", true},
		{"update only, no createAs object", false, true, false, false, true, "", true},
		{"no update parent, create field", true, false, false, true, false, Normal, false},
		{"no update parent, get only", true, false, false, false, false, Computed, false},
		{"no parent at all", false, false, false, false, false, "", false},
		{"folded, folded property", true, true, true, true, true, Normal, false},
		{"folded, get only", true, true, true, false, false, Computed, false},
		{"folded, update only", true, true, true, false, true, Computed, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f := &Field{Name: "x", Type: &Type{Kind: String}}
			if test.inCreate {
				f.Create = attrs
			}
			if test.inUpdate {
				f.Update = attrs
			}
			err := classifyField("Resource.p.x", f, test.create, test.update, test.folded)
			if (err != nil) != test.wantErr {
				t.Fatalf("error = %v, want error %t", err, test.wantErr)
			}
			if f.Behavior != test.want {
				t.Errorf("behavior = %q, want %q", f.Behavior, test.want)
			}
		})
	}
}
