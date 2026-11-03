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

// requestOnlyThing has request fields that the resource component Thing does not have: teamId in
// Create and Update, seed and secret in Create only, and access, an enum without its zero value,
// in Create and Update. mode is an enum without its zero value in every schema. Create requires
// secret.
const requestOnlyThing = `openapi: 3.1.0
info: {title: t, version: 1.0.0}
paths:
  /v1/things:
    post:
      operationId: Svc_CreateThing
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              required: [name, secret]
              properties:
                name: {type: string, minLength: 1}
                mode: {$ref: '#/components/schemas/Mode'}
                teamId: {$ref: '#/components/schemas/TeamId'}
                access: {$ref: '#/components/schemas/Access'}
                seed: {type: string}
                secret: {type: string, minLength: 1}
      responses: {'200': {description: ok, content: {application/json: {schema: {$ref: '#/components/schemas/Thing'}}}}}
  /v1/things/{id}:
    parameters: [{name: id, in: path, required: true, schema: {type: string}}]
    get:
      operationId: Svc_GetThing
      responses: {'200': {description: ok, content: {application/json: {schema: {$ref: '#/components/schemas/Thing'}}}}}
    put:
      operationId: Svc_UpdateThing
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              required: [name]
              properties:
                name: {type: string, minLength: 1}
                mode: {$ref: '#/components/schemas/Mode'}
                teamId: {$ref: '#/components/schemas/TeamId'}
                access: {$ref: '#/components/schemas/Access'}
      responses: {'200': {description: ok, content: {application/json: {schema: {$ref: '#/components/schemas/Thing'}}}}}
    delete:
      operationId: Svc_DeleteThing
      responses: {'200': {description: ok, content: {application/json: {schema: {type: object}}}}}
components:
  schemas:
    TeamId:
      type: object
      properties:
        id: {type: integer, format: int64}
    Mode:
      type: string
      enum: [MODE_FAST, MODE_SLOW]
    Access:
      type: string
      enum: [ACCESS_OPEN, ACCESS_CLOSED]
    Thing:
      type: object
      required: [id, name]
      properties:
        id: {type: string, readOnly: true}
        name: {type: string, minLength: 1}
        mode: {$ref: '#/components/schemas/Mode'}
`

// A skip line of the resource component also skips a field that only the Create and Update bodies
// have, or only the Create body. The skip applies before the lifecycle and enum checks, and the
// built resource does not have the field, so the generated requests never send it. A skipped field
// that a request requires is still an issue.
func TestSkipRequestOnlyField(t *testing.T) {
	doc, err := Load([]byte(requestOnlyThing))
	if err != nil {
		t.Fatal(err)
	}
	validate := func(skip ...string) map[string][]string {
		p := Policy{Existing: true, Skip: skip, Released: skip}
		found := map[string][]string{}
		for _, i := range ValidateWithPolicy(doc, "Thing", OperationIDs{}, p) {
			found[i.Location] = append(found[i.Location], i.Code)
		}
		return found
	}

	fields, err := RequestOnlyFields(doc, "Thing", OperationIDs{}, Policy{Existing: true})
	if err != nil {
		t.Fatal(err)
	}
	if want := []string{"access", "secret", "seed", "teamId"}; !slices.Equal(fields, want) {
		t.Fatalf("RequestOnlyFields = %v, want %v", fields, want)
	}

	unskipped := validate()
	for location, code := range map[string]string{
		"components.schemas.Thing.teamId": "FIELD_LIFECYCLE_UNSUPPORTED", // Create and Update
		"components.schemas.Thing.seed":   "FIELD_LIFECYCLE_UNSUPPORTED", // Create only
		"components.schemas.Thing.access": "FIELD_LIFECYCLE_UNSUPPORTED", // an enum without its zero value
		"Thing.mode":                      "ENUM_ZERO_INVALID",
	} {
		if !slices.Contains(unskipped[location], code) {
			t.Fatalf("without a skip line, %s has %v, want %s", location, unskipped[location], code)
		}
	}

	// secret has no skip line yet, so only its lifecycle is reported.
	skipped := validate("Thing.teamId", "Thing.seed", "Thing.access", "Thing.mode")
	delete(skipped, "Thing") // Build names the field that it cannot place: secret
	if len(skipped) != 1 || !slices.Equal(skipped["components.schemas.Thing.secret"], []string{"FIELD_LIFECYCLE_UNSUPPORTED"}) {
		t.Fatalf("with skip lines, the issues are %v, want only the lifecycle of the unskipped secret", skipped)
	}
	p := Policy{Existing: true, Skip: []string{"Thing.teamId", "Thing.seed", "Thing.access", "Thing.mode", "Thing.secret"}}
	p.Released = p.Skip
	r, err := BuildWithPolicy(doc, "Thing", OperationIDs{}, p)
	if err != nil {
		t.Fatal(err)
	}
	var built []string
	for _, f := range r.Fields {
		built = append(built, f.Name)
	}
	if want := []string{"id", "name"}; !slices.Equal(built, want) {
		t.Fatalf("built fields = %v, want %v: a skipped field is never sent", built, want)
	}

	// Create requires secret, so the generated request must send it.
	required := validate("Thing.teamId", "Thing.seed", "Thing.access", "Thing.mode", "Thing.secret")
	if len(required) != 1 || !slices.Equal(required["components.schemas.Thing.secret"], []string{"SKIP_REQUIRED_FIELD"}) {
		t.Fatalf("the issues are %v, want SKIP_REQUIRED_FIELD on secret only", required)
	}
}

// A skip line on a field that neither the resource component nor a request body has matches
// nothing. RequestOnlyFields does not return it, so the tool that checks the lines of its file
// reports the line as unused.
func TestSkipLineThatMatchesNothing(t *testing.T) {
	doc, err := Load([]byte(requestOnlyThing))
	if err != nil {
		t.Fatal(err)
	}
	p := Policy{Existing: true, Skip: []string{"Thing.missing"}, Released: []string{"Thing.missing"}}
	fields, err := RequestOnlyFields(doc, "Thing", OperationIDs{}, p)
	if err != nil {
		t.Fatal(err)
	}
	if slices.Contains(fields, "missing") || slices.Contains(fields, "name") || slices.Contains(fields, "mode") {
		t.Fatalf("RequestOnlyFields = %v: it has only fields that the requests have and Thing does not", fields)
	}
	if _, err := RequestOnlyFields(doc, "Other", OperationIDs{}, p); err == nil {
		t.Fatal("RequestOnlyFields of a resource without operations returned no error")
	}
}

// With the update mask in the body, the mask property of the Update body is no request-only field:
// Build reads it as the mask, so a skip line on it would change nothing. Without that policy the
// property is an ordinary field that only the Update body has.
func TestRequestOnlyFieldsLeaveOutABodyMask(t *testing.T) {
	spec := strings.Replace(requestOnlyThing, "                access: {$ref: '#/components/schemas/Access'}\n      responses: {'200': {description: ok, content: {application/json: {schema: {$ref: '#/components/schemas/Thing'}}}}}\n    delete:",
		"                access: {$ref: '#/components/schemas/Access'}\n                updateMask: {type: string}\n      responses: {'200': {description: ok, content: {application/json: {schema: {$ref: '#/components/schemas/Thing'}}}}}\n    delete:", 1)
	if spec == requestOnlyThing {
		t.Fatal("the fixture has no Update body to add the mask to")
	}
	doc, err := Load([]byte(spec))
	if err != nil {
		t.Fatal(err)
	}
	fields, err := RequestOnlyFields(doc, "Thing", OperationIDs{}, Policy{Existing: true, UpdateMaskInBody: true})
	if err != nil {
		t.Fatal(err)
	}
	if slices.Contains(fields, "updateMask") {
		t.Fatalf("RequestOnlyFields = %v, want no updateMask with the mask in the body", fields)
	}
	fields, err = RequestOnlyFields(doc, "Thing", OperationIDs{}, Policy{Existing: true})
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Contains(fields, "updateMask") {
		t.Fatalf("RequestOnlyFields = %v, want updateMask without the body-mask policy", fields)
	}
}
