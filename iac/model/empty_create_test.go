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

// emptyCreateSingleton is a singleton whose Create returns google.protobuf.Empty, an inline object
// with no fields.
const emptyCreateSingleton = `openapi: 3.1.0
info: {title: t, version: 1.0.0}
paths:
  /v1/setting:
    post:
      operationId: Svc_CreateSetting
      requestBody: {required: true, content: {application/json: {schema: {$ref: '#/components/schemas/Setting'}}}}
      responses: {'200': {description: ok, content: {application/json: {schema: {type: object}}}}}
    get:
      operationId: Svc_GetSetting
      responses: {'200': {description: ok, content: {application/json: {schema: {$ref: '#/components/schemas/Setting'}}}}}
    put:
      operationId: Svc_UpdateSetting
      requestBody: {required: true, content: {application/json: {schema: {$ref: '#/components/schemas/Setting'}}}}
      responses: {'200': {description: ok, content: {application/json: {schema: {$ref: '#/components/schemas/Setting'}}}}}
    delete:
      operationId: Svc_DeleteSetting
      responses: {'200': {description: ok, content: {application/json: {schema: {type: object}}}}}
components:
  schemas:
    Setting:
      type: object
      properties:
        id: {type: string, readOnly: true}
        name: {type: string}
`

// emptyCreateItem is the same API for an item: Get, Update, and Delete need the id, which the
// empty Create response does not have.
const emptyCreateItem = `openapi: 3.1.0
info: {title: t, version: 1.0.0}
paths:
  /v1/setting:
    post:
      operationId: Svc_CreateSetting
      requestBody: {required: true, content: {application/json: {schema: {$ref: '#/components/schemas/Setting'}}}}
      responses: {'200': {description: ok, content: {application/json: {schema: {type: object}}}}}
  /v1/setting/{id}:
    parameters: [{name: id, in: path, required: true, schema: {type: string}}]
    get:
      operationId: Svc_GetSetting
      responses: {'200': {description: ok, content: {application/json: {schema: {$ref: '#/components/schemas/Setting'}}}}}
    put:
      operationId: Svc_UpdateSetting
      requestBody: {required: true, content: {application/json: {schema: {$ref: '#/components/schemas/Setting'}}}}
      responses: {'200': {description: ok, content: {application/json: {schema: {$ref: '#/components/schemas/Setting'}}}}}
    delete:
      operationId: Svc_DeleteSetting
      responses: {'200': {description: ok, content: {application/json: {schema: {type: object}}}}}
components:
  schemas:
    Setting:
      type: object
      required: [id]
      properties:
        id: {type: string, readOnly: true}
        name: {type: string}
`

// An existing singleton whose Create returns an inline object with no fields reads the resource
// with Get after Create. A new resource, and an existing item resource whose Get needs the id that
// the empty response does not have, are rejected.
func TestEmptyInlineCreateResponse(t *testing.T) {
	build := func(item bool, policy Policy) (*Resource, error) {
		spec := emptyCreateSingleton
		if item {
			spec = emptyCreateItem
		}
		doc, err := Load([]byte(spec))
		if err != nil {
			t.Fatal(err)
		}
		return BuildWithPolicy(doc, "Setting", OperationIDs{}, policy)
	}
	r, err := build(false, Policy{Existing: true})
	if err != nil {
		t.Fatal(err)
	}
	if !r.Singleton || !r.Create.Response.Empty || r.Create.Response.Schema != "" {
		t.Fatalf("singleton = %t, create response = %+v, want an empty response", r.Singleton, r.Create.Response)
	}
	for name, test := range map[string]struct {
		item   bool
		policy Policy
	}{
		"new singleton": {false, Policy{}},
		"existing item": {true, Policy{Existing: true}},
		"new item":      {true, Policy{}},
	} {
		t.Run(name, func(t *testing.T) {
			if _, err := build(test.item, test.policy); err == nil || !strings.Contains(err.Error(), "200 response schema is inline, want a $ref") {
				t.Fatalf("err = %v, want the inline Create response rejected", err)
			}
		})
	}
	// An id property for a Create response that has no fields is a stale override: it is
	// reported, not ignored.
	if _, err := build(false, Policy{Existing: true, CreateResponseOnlyIDProperty: "settingId"}); err == nil || !strings.Contains(err.Error(), `has no property "settingId"`) {
		t.Fatalf("err = %v, want the stale api.createResponseOnlyIDProperty reported", err)
	}
}
