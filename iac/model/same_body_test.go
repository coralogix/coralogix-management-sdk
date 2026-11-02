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
	"strings"
	"testing"
)

// sameBodyThing has Create and Update inline bodies without a title that are the same schema, as
// the parsing rule groups API has. openapi-generator makes one type of the two.
const sameBodyThing = `openapi: 3.1.0
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
              description: Request to create a thing.
              properties:
                name: {type: string, x-coralogix-presence: true}
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
              description: Request to create a thing.
              properties:
                name: {type: string, x-coralogix-presence: true}
      responses: {'200': {description: ok, content: {application/json: {schema: {$ref: '#/components/schemas/Thing'}}}}}
    delete:
      operationId: Svc_DeleteThing
      responses: {'200': {description: ok, content: {application/json: {schema: {type: object}}}}}
components:
  schemas:
    Thing:
      type: object
      required: [id]
      properties:
        id: {type: string, readOnly: true}
        name: {type: string, x-coralogix-presence: true}
`

func TestUpdateBodyAsCreateWhenTheInlineBodiesAreTheSame(t *testing.T) {
	r := buildSameBody(t, sameBodyThing)
	if !r.Update.BodyAsCreate {
		t.Fatal("Update.BodyAsCreate = false, want true for two equal inline bodies")
	}
	if r.Create.BodyAsCreate {
		t.Fatal("Create.BodyAsCreate = true, want false")
	}
}

func TestUpdateBodyIsItsOwnWhenTheInlineBodiesDiffer(t *testing.T) {
	spec := strings.Replace(sameBodyThing, "Request to create a thing.", "Request to update a thing.", 2)
	spec = strings.Replace(spec, "Request to update a thing.", "Request to create a thing.", 1)
	if r := buildSameBody(t, spec); r.Update.BodyAsCreate {
		t.Fatal("Update.BodyAsCreate = true for bodies with another description")
	}
	titled := strings.Replace(sameBodyThing, "              type: object\n              description: Request to create a thing.", "              title: ThingBody\n              type: object\n              description: Request to create a thing.", 2)
	if r := buildSameBody(t, titled); r.Update.BodyAsCreate {
		t.Fatal("Update.BodyAsCreate = true for bodies with a title, which names each one")
	}
}

func buildSameBody(t *testing.T, spec string) *Resource {
	t.Helper()
	doc, err := Load([]byte(spec))
	if err != nil {
		t.Fatal(err)
	}
	r, err := Build(doc, "Thing")
	if err != nil {
		t.Fatal(err)
	}
	return r
}
