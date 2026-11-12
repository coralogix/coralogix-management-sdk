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

// memberCreateResult is the createResult line of the members contract.
func memberCreateResult() *CreateResultLine {
	return &CreateResultLine{Resource: "member", Status: "status", Created: []string{"CREATE_MEMBER_STATUS_CREATED"}, Message: "message"}
}

func TestMemberCreateResult(t *testing.T) {
	p := memberBatchPolicy()
	p.CreateResult = memberCreateResult()
	r := buildMember(t, memberSource(t, memberWithItemPaths), p)
	got := r.Create.Response
	if got.Schema != "CreateMembersResponse" || got.Field != "" || got.Items || got.Result == nil {
		t.Fatalf("create response: %+v", got)
	}
	result := *got.Result
	if result.Items != "results" || result.Component != "CreateMemberResult" || result.Resource != "member" || result.Status != "status" || result.Message != "message" {
		t.Fatalf("create result: %+v", result)
	}
	if result.StatusType.Kind != Enum || result.StatusType.Schema != "CreateMemberStatus" || !slices.Equal(result.Created, []string{"CREATE_MEMBER_STATUS_CREATED"}) {
		t.Fatalf("create status: %+v, %v", result.StatusType, result.Created)
	}
	p.CreateResult.Message = ""
	if r := buildMember(t, memberSource(t, memberWithItemPaths), p); r.Create.Response.Result.Message != "" {
		t.Fatalf("message without a line: %+v", r.Create.Response.Result)
	}
}

func TestMemberCreateResultIssues(t *testing.T) {
	const results = "        results:\n          type: array\n          minItems: 0\n          items: {$ref: '#/components/schemas/CreateMemberResult'}\n"
	line := func(edit func(*CreateResultLine)) Policy {
		p := memberBatchPolicy()
		p.CreateResult = memberCreateResult()
		edit(p.CreateResult)
		return p
	}
	same := func(*CreateResultLine) {}
	withoutBatch := line(same)
	withoutBatch.BatchWrite = false
	withOnlyID := line(same)
	withOnlyID.CreateResponseOnlyIDProperty = "memberId"
	tests := map[string]struct {
		edits         []string
		policy        Policy
		code, message string
	}{
		"no results array": {
			edits: []string{results, "        results: {$ref: '#/components/schemas/CreateMemberResult'}\n"}, policy: line(same),
			code: "CREATE_RESULT_ITEMS_NOT_FOUND", message: "has no array of result components",
		},
		"two results arrays": {
			edits: []string{results, results + "        failures:\n          type: array\n          items: {$ref: '#/components/schemas/CreateMemberResult'}\n"}, policy: line(same),
			code: "CREATE_RESULT_ITEMS_AMBIGUOUS", message: "has more than one array of components",
		},
		"resource property": {
			policy: line(func(l *CreateResultLine) { l.Resource = "login" }),
			code:   "CREATE_RESULT_RESOURCE_NOT_FOUND", message: `no property "login" that holds Member`,
		},
		"status not enum": {
			policy: line(func(l *CreateResultLine) { l.Status = "login" }),
			code:   "CREATE_RESULT_STATUS_NOT_ENUM", message: `no enum property "login"`,
		},
		"unknown status": {
			policy: line(func(l *CreateResultLine) { l.Created = []string{"CREATE_MEMBER_STATUS_DONE"} }),
			code:   "CREATE_RESULT_STATUS_VALUE_INVALID", message: "is not a value of its enum",
		},
		"zero status": {
			policy: line(func(l *CreateResultLine) { l.Created = []string{"CREATE_MEMBER_STATUS_UNSPECIFIED"} }),
			code:   "CREATE_RESULT_STATUS_VALUE_INVALID", message: "is the zero value",
		},
		"no status": {
			policy: line(func(l *CreateResultLine) { l.Created = nil }),
			code:   "CREATE_RESULT_STATUS_VALUE_INVALID", message: "names no status value",
		},
		"message not string": {
			policy: line(func(l *CreateResultLine) { l.Message = "member" }),
			code:   "CREATE_RESULT_MESSAGE_NOT_STRING", message: `no string property "member"`,
		},
		"without batch write": {
			policy: withoutBatch, code: "RESOURCE_LIFECYCLE_INCOMPLETE", message: "api.batchWrite.createResult needs api.batchWrite",
		},
		"with only-id response": {
			policy: withOnlyID, code: "RESOURCE_LIFECYCLE_INCOMPLETE", message: "api.batchWrite.createResult and api.createResponseOnlyIDProperty both say what the Create response holds",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			memberIssue(t, memberSource(t, memberWithItemPaths, tt.edits), tt.policy, tt.code, tt.message)
		})
	}
}

func TestMemberCreateValues(t *testing.T) {
	p := memberBatchPolicy()
	p.CreateResult = memberCreateResult()
	p.CreateValues = map[string]string{"onboardingMode": "ONBOARDING_MODE_NO_INVITE"}
	r := buildMember(t, memberSource(t, memberWithItemPaths), p)
	if len(r.CreateValues) != 1 {
		t.Fatalf("create values %+v", r.CreateValues)
	}
	v := r.CreateValues[0]
	if v.Name != "onboardingMode" || v.Value != "ONBOARDING_MODE_NO_INVITE" || v.Type.Kind != Enum || v.Type.Schema != "OnboardingMode" {
		t.Fatalf("create value %+v, type %+v", v, v.Type)
	}
	if slices.ContainsFunc(r.Fields, func(f *ResourceField) bool { return f.Name == "onboardingMode" }) {
		t.Fatal("a create value is a resource field")
	}
}

func TestMemberCreateValueIssues(t *testing.T) {
	values := func(name, value string) Policy {
		p := memberBatchPolicy()
		p.CreateResult = memberCreateResult()
		p.CreateValues = map[string]string{name: value}
		return p
	}
	notExisting := values("onboardingMode", "ONBOARDING_MODE_NO_INVITE")
	notExisting.Existing, notExisting.BatchWrite, notExisting.CreateResult = false, false, nil
	tests := map[string]struct {
		policy        Policy
		code, message string
	}{
		"unknown property":   {values("teamId", "1"), "CREATE_VALUE_UNKNOWN", `no property "teamId" beside "template"`},
		"inside the wrapper": {values("login", "admin"), "CREATE_VALUE_INSIDE_WRAPPER", `"login" is the request wrapper "template" or a property inside it`},
		"the wrapper":        {values("template", "x"), "CREATE_VALUE_INSIDE_WRAPPER", `"template" is the request wrapper`},
		"unknown value":      {values("onboardingMode", "ONBOARDING_MODE_LATER"), "CREATE_VALUE_INVALID", "is not a value of its enum"},
		"zero value":         {values("onboardingMode", "ONBOARDING_MODE_UNSPECIFIED"), "CREATE_VALUE_INVALID", "is the zero value"},
		"new resource":       {notExisting, "RESOURCE_LIFECYCLE_INCOMPLETE", "api.createValues is only for existing resources"},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			memberIssue(t, memberSource(t, memberWithItemPaths), tt.policy, tt.code, tt.message)
		})
	}
}

// memberWithNumberOnboarding replaces the enum onboarding mode of the Create body with a bounded
// number, so the tests can pin the schema validation of a fixed numeric value.
var memberWithNumberOnboarding = []string{
	"        onboardingMode:\n          allOf:\n            - $ref: '#/components/schemas/OnboardingMode'\n",
	"        onboardingMode: {type: number, format: double, minimum: 0, maximum: 1}\n",
}

func TestMemberCreateValueNumberIssues(t *testing.T) {
	values := func(value string) Policy {
		p := memberBatchPolicy()
		p.CreateResult = memberCreateResult()
		p.CreateValues = map[string]string{"onboardingMode": value}
		return p
	}
	tests := map[string]struct {
		value, message string
	}{
		"not a number":  {"not-a-number", "is not a number"},
		"above maximum": {"2", "is above the maximum 1"},
		"below minimum": {"-0.5", "is below the minimum 0"},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			memberIssue(t, memberSource(t, memberWithItemPaths, memberWithNumberOnboarding), values(tt.value), "CREATE_VALUE_INVALID", tt.message)
		})
	}
	t.Run("value in range", func(t *testing.T) {
		r := buildMember(t, memberSource(t, memberWithItemPaths, memberWithNumberOnboarding), values("0.5"))
		if len(r.CreateValues) != 1 || r.CreateValues[0].Name != "onboardingMode" || r.CreateValues[0].Value != "0.5" || r.CreateValues[0].Type.Kind != Number {
			t.Fatalf("create values %+v", r.CreateValues)
		}
	})
}

// memberWithDeletePath gives the golden members contract a Delete of one member by its id path parameter,
// so a test can describe the List read apart from the batch Delete.
var memberWithDeletePath = []string{
	"  /members/v1:updateStatuses:\n", "  /members/v1/{memberId}:\n" + memberItemPaths[strings.Index(memberItemPaths, "    delete:"):] + "  /members/v1:updateStatuses:\n",
	"components:\n  schemas:\n", "components:\n  schemas:\n" + memberItemSchemas,
}

// memberListPolicy describes the members contract read from its List, and removed by id.
func memberListPolicy() Policy {
	p := memberBatchPolicy()
	p.CreateResult = memberCreateResult()
	p.CreateValues = map[string]string{"onboardingMode": "ONBOARDING_MODE_NO_INVITE"}
	p.GetFromList, p.IDProperty, p.ListFilters = true, "memberId", map[string]string{"login": "login"}
	p.GetQuery = map[string]string{"status": "MEMBER_STATUS_ACTIVE"}
	p.Operations.Get = "MembersService_SearchMembers"
	return p
}

func TestMemberGetFromList(t *testing.T) {
	r := buildMember(t, memberSource(t, memberWithDeletePath), memberListPolicy())
	if !r.GetFromList || r.Singleton || r.ID() != "memberId" || r.IDParam != "memberId" || r.IDType.Kind != String {
		t.Fatalf("identity: %+v", r)
	}
	if got := r.Get.Response; got.Schema != "SearchMembersResponse" || got.Field != "members" || !got.Items {
		t.Fatalf("get response: %+v", got)
	}
	l := r.List
	if l == nil || l.OperationID != r.Get.OperationID || l.ItemsField != "members" || !l.Full || l.IDField != "memberId" || !l.Paged {
		t.Fatalf("list: %+v", l)
	}
	var paging []string
	for _, p := range r.PageParameters {
		paging = append(paging, p.Name+":"+string(p.Type.Kind)+":"+p.Type.Format)
	}
	if !slices.Equal(paging, []string{"page_size:integer:int64", "page_token:integer:int64"}) || r.NextPage == nil || r.NextPage.Name != "nextPageToken" || r.NextPage.Type.Format != "int64" {
		t.Fatalf("paging %v, next page %+v", paging, r.NextPage)
	}
	if len(r.ListFilters) != 1 || r.ListFilters[0].Param != "login" || r.ListFilters[0].Field != "login" || r.ListFilters[0].Type.Kind != String {
		t.Fatalf("filters: %+v", r.ListFilters)
	}
	// Get always asks the List for active members: a removed member is inactive, and still listed.
	if len(r.GetQuery) != 1 || r.GetQuery[0].Name != "status" || r.GetQuery[0].Value != "MEMBER_STATUS_ACTIVE" || r.GetQuery[0].Type.Kind != Enum || r.GetQuery[0].Type.Schema != "MemberStatus" {
		t.Fatalf("get query: %+v", r.GetQuery)
	}
	if r.Delete.Path != "/members/v1/{memberId}" || r.Update.Path != r.Create.Path {
		t.Fatalf("paths: update %s, delete %s", r.Update.Path, r.Delete.Path)
	}
	// Without filters and a query, Get reads the whole List.
	p := memberListPolicy()
	p.ListFilters, p.GetQuery = nil, nil
	if r := buildMember(t, memberSource(t, memberWithDeletePath), p); len(r.ListFilters) != 0 || len(r.GetQuery) != 0 {
		t.Fatalf("filters and query without lines: %+v, %+v", r.ListFilters, r.GetQuery)
	}
	// The List can be named, if it is Get.
	p.Operations.List = p.Operations.Get
	buildMember(t, memberSource(t, memberWithDeletePath), p)
}

func TestMemberGetFromListOperationPagingOverridesPath(t *testing.T) {
	source := memberSource(t, memberWithDeletePath, []string{
		"  /members/v1:\n    get:\n", "  /members/v1:\n    parameters:\n      - {name: page_token, in: query, required: false, schema: {type: string}}\n    get:\n",
	})
	r := buildMember(t, source, memberListPolicy())
	var paging []string
	for _, p := range r.PageParameters {
		paging = append(paging, p.Name+":"+string(p.Type.Kind)+":"+p.Type.Format)
	}
	if !slices.Equal(paging, []string{"page_size:integer:int64", "page_token:integer:int64"}) {
		t.Fatalf("paging parameters: %v", paging)
	}
}

func TestMemberGetFromListIssues(t *testing.T) {
	const listItems = "        members:\n          type: array\n          minItems: 0\n          items: {$ref: '#/components/schemas/Member'}\n    UpdateMembersStatusesResponse:"
	listPolicy := func(edit func(*Policy)) Policy {
		p := memberListPolicy()
		edit(&p)
		return p
	}
	tests := map[string]struct {
		paths, edits  []string // paths are the operations by id; nil is the Delete by id
		policy        Policy
		code, message string
	}{
		"not a list": {
			edits:  []string{listItems, "        members: {$ref: '#/components/schemas/Member'}\n    UpdateMembersStatusesResponse:"},
			policy: memberListPolicy(), code: "GET_FROM_LIST_NOT_LIST", message: "is not a List that the model can describe: The response SearchMembersResponse has no array",
		},
		"summary items": {
			edits: []string{
				listItems, "        members:\n          type: array\n          items: {$ref: '#/components/schemas/MemberSummary'}\n    MemberSummary:\n      type: object\n      properties:\n        memberId: {type: string, minLength: 0}\n    UpdateMembersStatusesResponse:",
			},
			policy: memberListPolicy(), code: "GET_FROM_LIST_NOT_FULL", message: "The List item MemberSummary is not Member",
		},
		"no id property": {
			policy: listPolicy(func(p *Policy) { p.IDProperty = "" }),
			code:   "GET_FROM_LIST_ID_PROPERTY_REQUIRED", message: "must name the id property",
		},
		"new resource": {
			policy: listPolicy(func(p *Policy) { p.Existing, p.CreateValues = false, nil }),
			code:   "RESOURCE_LIFECYCLE_INCOMPLETE", message: "api.getFromList is only for existing resources",
		},
		"update id in path": {
			policy: listPolicy(func(p *Policy) { p.UpdateIDInBody, p.UpdateIDOutsideWrapper = false, false }),
			code:   "RESOURCE_LIFECYCLE_INCOMPLETE", message: "api.getFromList needs api.updateIDInBody",
		},
		"with getByIDs": {
			policy: listPolicy(func(p *Policy) { p.GetByIDs = "login" }),
			code:   "RESOURCE_LIFECYCLE_INCOMPLETE", message: "api.getFromList and api.getByIDs both say how Get reads the resource",
		},
		"another list": {
			policy: listPolicy(func(p *Policy) { p.Operations.List = "MembersService_UpdateMembersStatuses" }),
			code:   "RESOURCE_LIFECYCLE_INCOMPLETE", message: "api.getFromList reads the List with Get MembersService_SearchMembers",
		},
		"unknown filter": {
			policy: listPolicy(func(p *Policy) { p.ListFilters = map[string]string{"nickname": "login"} }),
			code:   "LIST_FILTER_UNKNOWN", message: `no query parameter "nickname"`,
		},
		"paging filter": {
			policy: listPolicy(func(p *Policy) { p.ListFilters = map[string]string{"page_size": "memberAccountId"} }),
			code:   "LIST_FILTER_UNKNOWN", message: `no query parameter "page_size"`,
		},
		"unknown field": {
			policy: listPolicy(func(p *Policy) { p.ListFilters = map[string]string{"login": "nickname"} }),
			code:   "LIST_FILTER_FIELD_UNKNOWN", message: `no field "nickname"`,
		},
		"other type": {
			policy: listPolicy(func(p *Policy) { p.ListFilters = map[string]string{"team_id": "login"} }),
			code:   "LIST_FILTER_TYPE_INCOMPATIBLE", message: `The filter "team_id" does not have the scalar type of the field "login"`,
		},
		"filters without the list read": {
			paths: memberWithItemPaths,
			policy: listPolicy(func(p *Policy) {
				p.GetFromList, p.IDProperty, p.GetQuery, p.Operations.Get = false, "", nil, "MembersService_GetMember"
			}),
			code: "RESOURCE_LIFECYCLE_INCOMPLETE", message: "api.getFromList.filters needs api.getFromList",
		},
		"unknown query": {
			policy: listPolicy(func(p *Policy) { p.GetQuery = map[string]string{"state": "MEMBER_STATUS_ACTIVE"} }),
			code:   "GET_QUERY_UNKNOWN", message: `no query parameter "state"`,
		},
		"paging query": {
			policy: listPolicy(func(p *Policy) { p.GetQuery = map[string]string{"page_size": "100"} }),
			code:   "GET_QUERY_UNKNOWN", message: `The query parameter "page_size" pages the List`,
		},
		"query value": {
			policy: listPolicy(func(p *Policy) { p.GetQuery = map[string]string{"status": "MEMBER_STATUS_GONE"} }),
			code:   "GET_QUERY_VALUE_INVALID", message: "is not a value of its enum",
		},
		"query type": {
			edits:  []string{"        - {name: page_size,", "        - {name: statuses, in: query, required: false, schema: {type: array, items: {$ref: '#/components/schemas/MemberStatus'}}}\n        - {name: page_size,"},
			policy: listPolicy(func(p *Policy) { p.GetQuery = map[string]string{"statuses": "MEMBER_STATUS_ACTIVE"} }),
			code:   "GET_QUERY_VALUE_INVALID", message: `The query parameter "statuses" is a list, not a scalar or an enum`,
		},
		"query and filter": {
			policy: listPolicy(func(p *Policy) { p.GetQuery = map[string]string{"login": "admin"} }),
			code:   "RESOURCE_LIFECYCLE_INCOMPLETE", message: `api.getFromList.query and api.getFromList.filters both name "login"`,
		},
		"query without the list read": {
			paths: memberWithItemPaths,
			policy: listPolicy(func(p *Policy) {
				p.GetFromList, p.IDProperty, p.ListFilters, p.Operations.Get = false, "", nil, "MembersService_GetMember"
			}),
			code: "RESOURCE_LIFECYCLE_INCOMPLETE", message: "api.getFromList.query needs api.getFromList",
		},
		"query on a new resource": {
			paths: memberWithItemPaths,
			policy: listPolicy(func(p *Policy) {
				p.Existing, p.BatchWrite, p.CreateResult, p.CreateValues = false, false, nil, nil
				p.GetFromList, p.IDProperty, p.ListFilters, p.Operations.Get = false, "", nil, "MembersService_GetMember"
			}),
			code: "RESOURCE_LIFECYCLE_INCOMPLETE", message: "api.getFromList.query needs api.getFromList",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			paths := tt.paths
			if paths == nil {
				paths = memberWithDeletePath
			}
			memberIssue(t, memberSource(t, paths, tt.edits), tt.policy, tt.code, tt.message)
		})
	}
}

// memberPolicy describes the golden members contract with every key.
func memberPolicy() Policy {
	p := memberListPolicy()
	p.Operations.Delete = "MembersService_UpdateMembersStatuses"
	p.DeleteIDs, p.DeleteValues = "memberIds", map[string]string{"status": "MEMBER_STATUS_INACTIVE"}
	return p
}

func TestMemberThingBuilds(t *testing.T) {
	r := buildMember(t, memberSource(t), memberPolicy())
	if !r.GetFromList || r.IDParam != "" || r.ID() != "memberId" || r.Singleton {
		t.Fatalf("identity: %+v", r)
	}
	if r.Delete.Method != "PATCH" || r.Delete.Body != "inline" || r.Delete.BodyTitle != "UpdateMemberStatusRequest" || r.Delete.BodyTitleIsComponent {
		t.Fatalf("delete: %+v", r.Delete)
	}
	if len(r.DeleteValues) != 1 || r.DeleteValues[0].Name != "status" || r.DeleteValues[0].Value != "MEMBER_STATUS_INACTIVE" || r.DeleteValues[0].Type.Schema != "MemberStatus" {
		t.Fatalf("delete values: %+v", r.DeleteValues)
	}
	if !r.Create.BodyItems || !r.Update.BodyItems || r.Create.Response.Result == nil || !r.Update.Response.Items || !r.Get.Response.Items || len(r.CreateValues) != 1 || len(r.ListFilters) != 1 || len(r.GetQuery) != 1 {
		t.Fatalf("resource: %+v", r)
	}
	// The resource of the golden contract is new nowhere: without Existing, no key applies.
	p := memberPolicy()
	p.Existing = false
	doc, err := Load([]byte(memberSource(t)))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := BuildWithPolicy(doc, "Member", OperationIDs{}, p); err == nil {
		t.Fatal("a new resource builds with the keys of a released one")
	}
}

func TestMemberDeleteIDsIssues(t *testing.T) {
	const memberIDs = "                memberIds:\n                  type: array\n                  minItems: 1\n"
	deletes := func(edit func(*Policy)) Policy {
		p := memberPolicy()
		edit(&p)
		return p
	}
	tests := map[string]struct {
		paths, edits  []string
		policy        Policy
		code, message string
	}{
		"no ids property": {
			policy: deletes(func(p *Policy) { p.DeleteIDs = "ids" }),
			code:   "DELETE_IDS_NOT_FOUND", message: `no property "ids" for the ids`,
		},
		"numeric ids": {
			policy: deletes(func(p *Policy) { p.DeleteIDs = "memberAccountIds" }),
			code:   "DELETE_IDS_NOT_STRING_ARRAY", message: `"memberAccountIds" is not an array of string ids`,
		},
		"scalar ids": {
			policy: deletes(func(p *Policy) { p.DeleteIDs, p.DeleteValues = "teamId", nil }),
			code:   "DELETE_IDS_NOT_STRING_ARRAY", message: `"teamId" is not an array of string ids`,
		},
		"two ids required": {
			edits:  []string{memberIDs, "                memberIds:\n                  type: array\n                  minItems: 2\n"},
			policy: memberPolicy(), code: "DELETE_IDS_ONE_ITEM_REFUSED", message: `"memberIds" does not accept one id`,
		},
		"unknown value": {
			policy: deletes(func(p *Policy) { p.DeleteValues = map[string]string{"reason": "LEFT"} }),
			code:   "DELETE_VALUE_UNKNOWN", message: `no property "reason" beside the ids`,
		},
		"ids as a value": {
			policy: deletes(func(p *Policy) { p.DeleteValues = map[string]string{"memberIds": "x"} }),
			code:   "DELETE_VALUE_UNKNOWN", message: `no property "memberIds" beside the ids`,
		},
		"invalid value": {
			policy: deletes(func(p *Policy) { p.DeleteValues = map[string]string{"status": "MEMBER_STATUS_GONE"} }),
			code:   "DELETE_VALUE_INVALID", message: "is not a value of its enum",
		},
		"zero value": {
			policy: deletes(func(p *Policy) { p.DeleteValues = map[string]string{"status": "MEMBER_STATUS_UNSPECIFIED"} }),
			code:   "DELETE_VALUE_INVALID", message: "is the zero value",
		},
		"required property": {
			edits:  []string{"              title: UpdateMemberStatusRequest\n              type: object\n", "              title: UpdateMemberStatusRequest\n              type: object\n              required: [teamId]\n"},
			policy: memberPolicy(), code: "DELETE_BODY_REQUIRED_PROPERTY", message: `requires "teamId", which a remove does not send`,
		},
		"required value": {
			edits:  []string{"              title: UpdateMemberStatusRequest\n              type: object\n", "              title: UpdateMemberStatusRequest\n              type: object\n              required: [memberIds, status]\n"},
			policy: memberPolicy(),
		},
		"delete by id": {
			paths: memberWithItemPaths,
			policy: deletes(func(p *Policy) {
				p.GetFromList, p.ListFilters, p.GetQuery, p.Operations.Get = false, nil, nil, "MembersService_GetMember"
			}),
			code: "RESOURCE_LIFECYCLE_INCOMPLETE", message: "api.delete.ids is for a resource that is read from its List",
		},
		"delete method": {
			policy: deletes(func(p *Policy) { p.DeleteIDs = "" }),
			code:   "OPERATION_METHOD_INCOMPATIBLE", message: "MembersService_UpdateMembersStatuses is PATCH, want DELETE",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			source := memberSource(t, tt.paths, tt.edits)
			if tt.code == "" {
				buildMember(t, source, tt.policy)
				return
			}
			memberIssue(t, source, tt.policy, tt.code, tt.message)
		})
	}
}
