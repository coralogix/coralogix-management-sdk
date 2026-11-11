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
	"slices"
	"strings"
	"testing"
)

// memberItemPaths gives the golden members contract a Get and a Delete of one member by its id path
// parameter, so a test can describe the batch writes apart from the List read and the batch Delete.
const memberItemPaths = `  /members/v1/{memberId}:
    get:
      tags: [Members Service]
      operationId: MembersService_GetMember
      parameters:
        - {name: memberId, in: path, required: true, schema: {type: string}}
      responses:
        '200':
          description: ok
          content:
            application/json:
              schema: {$ref: '#/components/schemas/GetMemberResponse'}
    delete:
      tags: [Members Service]
      operationId: MembersService_DeleteMember
      parameters:
        - {name: memberId, in: path, required: true, schema: {type: string}}
      responses:
        '200':
          description: ok
          content:
            application/json:
              schema: {$ref: '#/components/schemas/DeleteMemberResponse'}
`

const memberItemSchemas = `    GetMemberResponse:
      type: object
      properties:
        member: {$ref: '#/components/schemas/Member'}
    DeleteMemberResponse:
      type: object
`

// Edits of the golden members contract.
var (
	memberWithItemPaths = []string{
		"  /members/v1:updateStatuses:\n", memberItemPaths + "  /members/v1:updateStatuses:\n",
		"components:\n  schemas:\n", "components:\n  schemas:\n" + memberItemSchemas,
	}
	// memberCreateReturnsMember answers Create with the member in a wrapper, as before createResult.
	memberCreateReturnsMember = []string{
		"schema: {$ref: '#/components/schemas/CreateMembersResponse'}", "schema: {$ref: '#/components/schemas/GetMemberResponse'}",
	}
)

// memberSource returns the golden members contract after the edits, pairs of old and new text. An
// edit that matches nothing fails the test, so the contract cannot drift away from a test.
func memberSource(t *testing.T, edits ...[]string) string {
	t.Helper()
	data, err := os.ReadFile("testdata/memberthing.yaml")
	if err != nil {
		t.Fatal(err)
	}
	source := string(data)
	for _, pairs := range edits {
		for i := 0; i+1 < len(pairs); i += 2 {
			if !strings.Contains(source, pairs[i]) {
				t.Fatalf("the members contract has no %q", pairs[i])
			}
			source = strings.Replace(source, pairs[i], pairs[i+1], 1)
		}
	}
	return source
}

// memberBatchPolicy describes the writes of the members contract read by id.
func memberBatchPolicy() Policy {
	return Policy{
		Existing: true, BatchWrite: true, RequestWrapper: "template", UpdateIDInBody: true, UpdateIDOutsideWrapper: true, ServerSetID: true,
		ReadOnly:   []string{"Member.memberId"},
		Operations: OperationIDs{Create: "MembersService_CreateMembers", Get: "MembersService_GetMember", Update: "MembersService_UpdateMembers", Delete: "MembersService_DeleteMember"},
	}
}

func buildMember(t *testing.T, source string, p Policy) *Resource {
	t.Helper()
	doc, err := Load([]byte(source))
	if err != nil {
		t.Fatal(err)
	}
	if issues := ValidateWithPolicy(doc, "Member", OperationIDs{}, p); len(issues) != 0 {
		t.Fatal(issues)
	}
	r, err := BuildWithPolicy(doc, "Member", OperationIDs{}, p)
	if err != nil {
		t.Fatal(err)
	}
	return r
}

// memberIssue checks that Validate reports code and Build fails with an error that has message.
func memberIssue(t *testing.T, source string, p Policy, code, message string) {
	t.Helper()
	doc, err := Load([]byte(source))
	if err != nil {
		t.Fatal(err)
	}
	if codes := reportCodes(ValidateWithPolicy(doc, "Member", OperationIDs{}, p)); !slices.Contains(codes, code) {
		t.Fatalf("codes %v do not contain %s", codes, code)
	}
	if _, err := BuildWithPolicy(doc, "Member", OperationIDs{}, p); err == nil || !strings.Contains(err.Error(), message) {
		t.Fatalf("Build error = %v, want %q", err, message)
	}
}

func TestMemberBatchWrites(t *testing.T) {
	r := buildMember(t, memberSource(t, memberWithItemPaths, memberCreateReturnsMember), memberBatchPolicy())
	if r.Create.Body != "CreateMemberRequest" || !r.Create.BodyItems || r.Update.Body != "UpdateMemberRequest" || !r.Update.BodyItems {
		t.Fatalf("batch bodies: create %+v, update %+v", r.Create, r.Update)
	}
	if r.Create.RequestComponent != "MemberTemplate" || r.Update.RequestComponent != "MemberTemplate" {
		t.Fatalf("request components: %q, %q", r.Create.RequestComponent, r.Update.RequestComponent)
	}
	// The Update response holds the written member in an array, beside the account ids.
	if got := r.Update.Response; !got.Items || got.Field != "members" || got.Schema != "UpdateMembersResponse" {
		t.Fatalf("update response: %+v", got)
	}
	if r.Create.Response.Items || r.Get.Response.Items || r.Delete.Body != "" {
		t.Fatalf("only the Update response is an array: %+v", r)
	}
	behaviors := map[string]Behavior{}
	for _, f := range r.Fields {
		behaviors[f.Name] = f.Behavior
	}
	want := map[string]Behavior{
		"displayName": Normal, "login": Normal, "loginModes": Normal, "status": Normal,
		"groupIds": Computed, "memberAccountId": Computed, "memberId": Computed,
	}
	for name, behavior := range want {
		if behaviors[name] != behavior {
			t.Errorf("%s: behavior %q, want %q", name, behaviors[name], behavior)
		}
	}
	if len(behaviors) != len(want) {
		t.Errorf("fields %v, want %v", behaviors, want)
	}
	// memberAccountId beside the wrapper of Update is no request field.
	doc, err := Load([]byte(memberSource(t, memberWithItemPaths, memberCreateReturnsMember)))
	if err != nil {
		t.Fatal(err)
	}
	if fields, err := RequestOnlyFields(doc, "Member", OperationIDs{}, memberBatchPolicy()); err != nil || len(fields) != 0 {
		t.Fatalf("request-only fields %v, %v", fields, err)
	}
}

func TestMemberBatchUpdateWrappedResponse(t *testing.T) {
	r := buildMember(t, memberSource(t, memberWithItemPaths, memberCreateReturnsMember, []string{
		"schema: {$ref: '#/components/schemas/UpdateMembersResponse'}", "schema: {$ref: '#/components/schemas/GetMemberResponse'}",
	}), memberBatchPolicy())
	if got := r.Update.Response; got.Items || got.Field != "member" {
		t.Fatalf("update response: %+v", got)
	}
}

func TestMemberBatchWriteIssues(t *testing.T) {
	const (
		createArray = "            schema:\n              type: array\n              minItems: 0\n              items: {$ref: '#/components/schemas/CreateMemberRequest'}\n"
		updateArray = "            schema:\n              type: array\n              minItems: 0\n              items: {$ref: '#/components/schemas/UpdateMemberRequest'}\n"
		createItem  = "            schema: {$ref: '#/components/schemas/CreateMemberRequest'}\n"
		updateItem  = "            schema: {$ref: '#/components/schemas/UpdateMemberRequest'}\n"
	)
	notExisting := memberBatchPolicy()
	notExisting.Existing = false
	tests := map[string]struct {
		edits         []string
		policy        Policy
		code, message string
	}{
		"object create body": {
			edits: []string{createArray, createItem}, policy: memberBatchPolicy(),
			code: "BATCH_WRITE_BODY_NOT_ARRAY", message: "The Create body is not an array.",
		},
		"object bodies": {
			edits: []string{createArray, createItem, updateArray, updateItem}, policy: memberBatchPolicy(),
			code: "OVERRIDE_UNUSED", message: "Neither the Create nor the Update body is an array.",
		},
		"inline items": {
			edits:  []string{"items: {$ref: '#/components/schemas/UpdateMemberRequest'}", "items: {type: object, properties: {memberId: {type: string}}}"},
			policy: memberBatchPolicy(), code: "BATCH_WRITE_BODY_NOT_ARRAY", message: "The items of the Update body are not a component.",
		},
		"two items required": {
			edits:  []string{"minItems: 0\n              items: {$ref: '#/components/schemas/UpdateMemberRequest'}", "minItems: 2\n              items: {$ref: '#/components/schemas/UpdateMemberRequest'}"},
			policy: memberBatchPolicy(), code: "BATCH_WRITE_ONE_ITEM_REFUSED", message: "The Update body does not accept an array of one item.",
		},
		"no item accepted": {
			edits:  []string{"minItems: 0\n              items: {$ref: '#/components/schemas/CreateMemberRequest'}", "maxItems: 0\n              items: {$ref: '#/components/schemas/CreateMemberRequest'}"},
			policy: memberBatchPolicy(), code: "BATCH_WRITE_ONE_ITEM_REFUSED", message: "The Create body does not accept an array of one item.",
		},
		"new resource": {
			policy: notExisting, code: "RESOURCE_LIFECYCLE_INCOMPLETE", message: "api.batchWrite is only for existing resources",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			memberIssue(t, memberSource(t, memberWithItemPaths, memberCreateReturnsMember, tt.edits), tt.policy, tt.code, tt.message)
		})
	}
}

// TestMemberArrayBodyNeedsBatchWrite shows that an array body stays outside the model without the key.
func TestMemberArrayBodyNeedsBatchWrite(t *testing.T) {
	p := memberBatchPolicy()
	p.BatchWrite = false
	doc, err := Load([]byte(memberSource(t, memberWithItemPaths, memberCreateReturnsMember)))
	if err != nil {
		t.Fatal(err)
	}
	if len(ValidateWithPolicy(doc, "Member", OperationIDs{}, p)) == 0 {
		t.Fatal("an array body validates without api.batchWrite")
	}
	if _, err := BuildWithPolicy(doc, "Member", OperationIDs{}, p); err == nil {
		t.Fatal("an array body builds without api.batchWrite")
	}
}
