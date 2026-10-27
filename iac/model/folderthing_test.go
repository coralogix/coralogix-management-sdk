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
)

// folderPolicy is the policy of a released resource like the dashboard folders: the request
// wraps the resource, the id is in the Update body, and Create returns only the id.
func folderPolicy() Policy {
	return Policy{
		Existing:                     true,
		RequestWrapper:               "thing",
		UpdateIDInBody:               true,
		IDProperty:                   "id",
		ServerSetID:                  true,
		CreateResponseOnlyIDProperty: "thingId",
		ReadOnly:                     []string{"FolderThing.id"},
	}
}

func buildFolderThing(t *testing.T, policy Policy) (*Resource, error) {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("testdata", "folderthing.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	doc, err := Load(data)
	if err != nil {
		t.Fatal(err)
	}
	return BuildWithPolicy(doc, "FolderThing", OperationIDs{}, policy)
}

// A Create response that holds only the id, and an Update response with no fields, are read as
// such: the tool reads the resource with Get.
func TestCreateResponseOnlyIDPropertyAndEmptyUpdate(t *testing.T) {
	r, err := buildFolderThing(t, folderPolicy())
	if err != nil {
		t.Fatal(err)
	}
	if got := r.Create.Response; got.ID != "thingId" || got.Schema != "CreateFolderThingResponse" || got.Direct || got.Field != "" {
		t.Fatalf("create response = %+v, want only the id thingId", got)
	}
	if got := r.Update.Response; !got.Empty {
		t.Fatalf("update response = %+v, want an empty response", got)
	}
	if got := r.Get.Response; got.Field != "thing" {
		t.Fatalf("get response = %+v, want the resource in thing", got)
	}
}

// The id of a Create response is a string, and the property must exist.
func TestCreateResponseOnlyIDPropertyIsChecked(t *testing.T) {
	tests := map[string]struct {
		key  string
		want string
	}{
		"unknown property": {"key", `has no property "key"`},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			policy := folderPolicy()
			policy.CreateResponseOnlyIDProperty = test.key
			if _, err := buildFolderThing(t, policy); err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("err = %v, want %q", err, test.want)
			}
		})
	}
	// Without the key, the Create response is not the resource, and the build says so.
	policy := folderPolicy()
	policy.CreateResponseOnlyIDProperty = ""
	if _, err := buildFolderThing(t, policy); err == nil {
		t.Fatal("create response with only an id and no createResponseOnlyIDProperty: want an error")
	}
	// A new resource has no Policy.Existing, so an empty Update response is not accepted.
	if _, err := buildFolderThing(t, Policy{CreateResponseOnlyIDProperty: "thingId"}); err == nil {
		t.Fatal("empty update response of a new resource: want an error")
	}
}

// An override that is stale, or that names a property of a response with no properties, is an
// error and does not panic.
func TestCreateResponseOnlyIDPropertyStaleAndEmpty(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("testdata", "folderthing.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	cases := map[string]struct{ old, new, want string }{
		"holds the resource": {"        thingId: {type: string, x-coralogix-presence: true}\n", "        thingId: {type: string, x-coralogix-presence: true}\n        thing:\n          allOf: [{$ref: '#/components/schemas/FolderThing'}]\n", "holds the resource"},
		"no properties":      {"    CreateFolderThingResponse:\n      type: object\n      properties:\n        thingId: {type: string, x-coralogix-presence: true}\n", "    CreateFolderThingResponse:\n      type: object\n", `has no property "thingId"`},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			changed := strings.Replace(string(data), c.old, c.new, 1)
			if changed == string(data) {
				t.Fatal("the test contract did not change: update the replaced text")
			}
			doc, err := Load([]byte(changed))
			if err != nil {
				t.Fatal(err)
			}
			if _, err := BuildWithPolicy(doc, "FolderThing", OperationIDs{}, folderPolicy()); err == nil || !strings.Contains(err.Error(), c.want) {
				t.Fatalf("err = %v, want %q", err, c.want)
			}
		})
	}
}
