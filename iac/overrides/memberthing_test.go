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

package overrides

import (
	"os"
	"slices"
	"testing"

	"go.yaml.in/yaml/v4"

	"github.com/coralogix/coralogix-management-sdk/iac/model"
)

func exceptionsOf(t *testing.T, source string) Exceptions {
	t.Helper()
	var e Exceptions
	if err := yaml.Unmarshal([]byte(source), &e); err != nil {
		t.Fatal(err)
	}
	return e
}

func TestBatchWritePolicy(t *testing.T) {
	e := exceptionsOf(t, "api: {batchWrite: {}, requestWrapper: template}\n")
	if err := e.API.Check(); err != nil {
		t.Fatal(err)
	}
	if p := e.Policy(); !p.BatchWrite || !p.Existing {
		t.Fatalf("policy %+v, want an existing batch write", p)
	}
	if p := exceptionsOf(t, "api: {requestWrapper: template}\n").Policy(); p.BatchWrite {
		t.Fatal("batchWrite is set without its key")
	}
	overwrite := exceptionsOf(t, "api: {batchWrite: {}, overwrite: {}, operations: {create: W, get: G, update: W, delete: W}}\n")
	if err := overwrite.API.Check(); err == nil {
		t.Fatal("api.overwrite goes with api.batchWrite")
	}
}

func TestCreateResultPolicy(t *testing.T) {
	e := exceptionsOf(t, "api:\n  batchWrite:\n    createResult: {resource: member, status: status, created: [CREATED], message: message}\n")
	if err := e.API.Check(); err != nil {
		t.Fatal(err)
	}
	r := e.Policy().CreateResult
	if r == nil || r.Resource != "member" || r.Status != "status" || r.Message != "message" || !slices.Equal(r.Created, []string{"CREATED"}) {
		t.Fatalf("create result %+v", r)
	}
	for _, line := range []string{"{status: status, created: [CREATED]}", "{resource: member, created: [CREATED]}", "{resource: member, status: status}"} {
		if err := exceptionsOf(t, "api: {batchWrite: {createResult: "+line+"}}\n").API.Check(); err == nil {
			t.Errorf("createResult %s passes the check", line)
		}
	}
}

func TestCreateValuesPolicy(t *testing.T) {
	e := exceptionsOf(t, "api: {requestWrapper: template, createValues: {onboardingMode: ONBOARDING_MODE_NO_INVITE}}\n")
	if err := e.API.Check(); err != nil {
		t.Fatal(err)
	}
	if got := e.Policy().CreateValues; len(got) != 1 || got["onboardingMode"] != "ONBOARDING_MODE_NO_INVITE" {
		t.Fatalf("create values %v", got)
	}
	if err := exceptionsOf(t, "api: {createValues: {onboardingMode: ''}}\n").API.Check(); err == nil {
		t.Fatal("an empty create value passes the check")
	}
}

func TestGetFromListPolicy(t *testing.T) {
	e := exceptionsOf(t, "api: {idProperty: memberId, getFromList: {filters: {login: login}, query: {status: MEMBER_STATUS_ACTIVE}}}\n")
	if err := e.API.Check(); err != nil {
		t.Fatal(err)
	}
	if p := e.Policy(); !p.GetFromList || len(p.ListFilters) != 1 || p.ListFilters["login"] != "login" || len(p.GetQuery) != 1 || p.GetQuery["status"] != "MEMBER_STATUS_ACTIVE" {
		t.Fatalf("policy %+v", p)
	}
	if p := exceptionsOf(t, "api: {getFromList: {}}\n").Policy(); !p.GetFromList || len(p.ListFilters) != 0 || len(p.GetQuery) != 0 {
		t.Fatalf("policy without filters %+v", p)
	}
	for _, source := range []string{"api: {getByIDs: ids, getFromList: {}}\n", "api: {getFromList: {filters: {login: ''}}}\n", "api: {getFromList: {query: {status: ''}}}\n"} {
		if err := exceptionsOf(t, source).API.Check(); err == nil {
			t.Errorf("%q passes the check", source)
		}
	}
}

func TestDeleteIDsPolicy(t *testing.T) {
	for _, source := range []string{
		"api:\n  operations: {create: C, get: G, update: U, delete: D}\n  delete: {ids: memberIds, values: {status: INACTIVE}}\n",
		"api:\n  delete: {operation: D, ids: memberIds, values: {status: INACTIVE}}\n",
	} {
		e := exceptionsOf(t, source)
		if err := e.API.Check(); err != nil {
			t.Fatalf("%q: %v", source, err)
		}
		if p := e.Policy(); p.DeleteIDs != "memberIds" || len(p.DeleteValues) != 1 || p.DeleteValues["status"] != "INACTIVE" {
			t.Fatalf("policy %+v", p)
		}
	}
	for _, source := range []string{
		"api: {delete: {ids: memberIds}}\n",
		"api: {delete: {operation: D, values: {status: INACTIVE}}}\n",
		"api: {delete: {operation: D, ids: memberIds, values: {status: ''}}}\n",
		"api:\n  operations: {create: C, get: G, update: U, delete: D}\n  delete: {operation: D, ids: memberIds}\n",
	} {
		if err := exceptionsOf(t, source).API.Check(); err == nil {
			t.Errorf("%q passes the check", source)
		}
	}
}

// TestMemberThingPolicyBuildsResource reads the keys of the members contract from an overrides file.
func TestMemberThingPolicyBuildsResource(t *testing.T) {
	e := exceptionsOf(t, `api:
  requestWrapper: template
  updateIDInBody: true
  updateIDOutsideWrapper: true
  serverSetID: true
  idProperty: memberId
  batchWrite:
    createResult: {resource: member, status: status, created: [CREATE_MEMBER_STATUS_CREATED], message: message}
  createValues: {onboardingMode: ONBOARDING_MODE_NO_INVITE}
  getFromList:
    filters: {login: login}
    query: {status: MEMBER_STATUS_ACTIVE}
  operations:
    create: MembersService_CreateMembers
    get: MembersService_SearchMembers
    update: MembersService_UpdateMembers
    delete: MembersService_UpdateMembersStatuses
  delete:
    ids: memberIds
    values: {status: MEMBER_STATUS_INACTIVE}
fields: {Member: {memberId: {readOnly: true}}}
`)
	if err := e.API.Check(); err != nil {
		t.Fatal(err)
	}
	raw, err := os.ReadFile("../model/testdata/memberthing.yaml")
	if err != nil {
		t.Fatal(err)
	}
	doc, err := model.Load(raw)
	if err != nil {
		t.Fatal(err)
	}
	if issues := model.ValidateWithPolicy(doc, "Member", model.OperationIDs{}, e.Policy()); len(issues) != 0 {
		t.Fatal(issues)
	}
	r, err := model.BuildWithPolicy(doc, "Member", model.OperationIDs{}, e.Policy())
	if err != nil {
		t.Fatal(err)
	}
	if !r.GetFromList || r.ID() != "memberId" || r.Create.Response.Result == nil || len(r.CreateValues) != 1 || len(r.DeleteValues) != 1 || len(r.ListFilters) != 1 || len(r.GetQuery) != 1 {
		t.Fatalf("members resource: %+v", r)
	}
}
