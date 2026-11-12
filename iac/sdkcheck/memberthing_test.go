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

package sdkcheck

import (
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/coralogix/coralogix-management-sdk/iac/model"
)

// memberPolicy describes the golden members contract with every key.
func memberPolicy() model.Policy {
	return model.Policy{
		Existing: true, RequestWrapper: "template", UpdateIDInBody: true, UpdateIDOutsideWrapper: true, ServerSetID: true,
		IDProperty: "memberId", ReadOnly: []string{"Member.memberId"},
		BatchWrite:   true,
		CreateResult: &model.CreateResultLine{Resource: "member", Status: "status", Created: []string{"CREATE_MEMBER_STATUS_CREATED"}, Message: "message"},
		CreateValues: map[string]string{"onboardingMode": "ONBOARDING_MODE_NO_INVITE"},
		GetFromList:  true, ListFilters: map[string]string{"login": "login"}, GetQuery: map[string]string{"status": "MEMBER_STATUS_ACTIVE"},
		DeleteIDs: "memberIds", DeleteValues: map[string]string{"status": "MEMBER_STATUS_INACTIVE"},
		Operations: model.OperationIDs{Create: "MembersService_CreateMembers", Get: "MembersService_SearchMembers", Update: "MembersService_UpdateMembers", Delete: "MembersService_UpdateMembersStatuses"},
	}
}

func TestMemberThingSDKNames(t *testing.T) {
	data, err := os.ReadFile("../model/testdata/memberthing.yaml")
	if err != nil {
		t.Fatal(err)
	}
	doc, err := model.Load(data)
	if err != nil {
		t.Fatal(err)
	}
	r, err := model.BuildWithPolicy(doc, "Member", model.OperationIDs{}, memberPolicy())
	if err != nil {
		t.Fatal(err)
	}
	refs, err := Names(r, "Members Service", "example/sdk", nil)
	if err != nil {
		t.Fatal(err)
	}
	want := map[string][]Ref{
		"create.body": {
			{Kind: KindType, Name: "CreateMemberRequest"},
			{Kind: KindMethod, Owner: "ApiMembersServiceCreateMembersRequest", Name: "CreateMemberRequest", Want: "func(createMemberRequest []CreateMemberRequest) ApiMembersServiceCreateMembersRequest"},
		},
		"create.body.template":                                 {{Kind: KindField, Owner: "CreateMemberRequest", Name: "Template", Want: "*MemberTemplate"}},
		"create.body.onboardingMode":                           {{Kind: KindField, Owner: "CreateMemberRequest", Name: "OnboardingMode", Want: "*OnboardingMode"}, {Kind: KindType, Name: "OnboardingMode"}},
		"create.body.onboardingMode.ONBOARDING_MODE_NO_INVITE": {{Kind: KindConst, Name: "ONBOARDINGMODE_ONBOARDING_MODE_NO_INVITE"}},
		"create.response.results":                              {{Kind: KindField, Owner: "CreateMembersResponse", Name: "Results", Want: "[]CreateMemberResult"}},
		"create.result":                                        {{Kind: KindType, Name: "CreateMemberResult"}},
		"create.result.member":                                 {{Kind: KindField, Owner: "CreateMemberResult", Name: "Member", Want: "*Member"}},
		"create.result.status":                                 {{Kind: KindField, Owner: "CreateMemberResult", Name: "Status", Want: "*CreateMemberStatus"}, {Kind: KindType, Name: "CreateMemberStatus"}},
		"create.result.status.CREATE_MEMBER_STATUS_CREATED":    {{Kind: KindConst, Name: "CREATEMEMBERSTATUS_CREATE_MEMBER_STATUS_CREATED"}},
		"create.result.message":                                {{Kind: KindField, Owner: "CreateMemberResult", Name: "Message", Want: "*string"}},
		"get": {
			{Kind: KindMethod, Owner: "MembersServiceAPIService", Name: "MembersServiceSearchMembers", Want: "func(ctx context.Context) ApiMembersServiceSearchMembersRequest"},
			{Kind: KindType, Name: "ApiMembersServiceSearchMembersRequest"},
			{Kind: KindMethod, Owner: "ApiMembersServiceSearchMembersRequest", Name: "Execute", Want: "func() (*SearchMembersResponse, *http.Response, error)"},
		},
		"get.page.page_size":                    {{Kind: KindMethod, Owner: "ApiMembersServiceSearchMembersRequest", Name: "PageSize", Want: "func(pageSize int64) ApiMembersServiceSearchMembersRequest"}},
		"get.page.page_token":                   {{Kind: KindMethod, Owner: "ApiMembersServiceSearchMembersRequest", Name: "PageToken", Want: "func(pageToken int64) ApiMembersServiceSearchMembersRequest"}},
		"get.filter.login":                      {{Kind: KindMethod, Owner: "ApiMembersServiceSearchMembersRequest", Name: "Login", Want: "func(login string) ApiMembersServiceSearchMembersRequest"}},
		"get.query.status":                      {{Kind: KindMethod, Owner: "ApiMembersServiceSearchMembersRequest", Name: "Status", Want: "func(status MemberStatus) ApiMembersServiceSearchMembersRequest"}, {Kind: KindType, Name: "MemberStatus"}},
		"get.query.status.MEMBER_STATUS_ACTIVE": {{Kind: KindConst, Name: "MEMBERSTATUS_MEMBER_STATUS_ACTIVE"}},
		"get.response.members":                  {{Kind: KindField, Owner: "SearchMembersResponse", Name: "Members", Want: "[]Member"}},
		"get.response.nextPageToken":            {{Kind: KindField, Owner: "SearchMembersResponse", Name: "NextPageToken", Want: "*int64"}},
		"update.body": {
			{Kind: KindType, Name: "UpdateMemberRequest"},
			{Kind: KindMethod, Owner: "ApiMembersServiceUpdateMembersRequest", Name: "UpdateMemberRequest", Want: "func(updateMemberRequest []UpdateMemberRequest) ApiMembersServiceUpdateMembersRequest"},
		},
		"update.body.memberId":    {{Kind: KindField, Owner: "UpdateMemberRequest", Name: "MemberId", Want: "*string"}},
		"update.response.members": {{Kind: KindField, Owner: "UpdateMembersResponse", Name: "Members", Want: "[]Member"}},
		"delete": {
			{Kind: KindMethod, Owner: "MembersServiceAPIService", Name: "MembersServiceUpdateMembersStatuses", Want: "func(ctx context.Context) ApiMembersServiceUpdateMembersStatusesRequest"},
			{Kind: KindType, Name: "ApiMembersServiceUpdateMembersStatusesRequest"},
			{Kind: KindMethod, Owner: "ApiMembersServiceUpdateMembersStatusesRequest", Name: "Execute", Want: "func() (*UpdateMembersStatusesResponse, *http.Response, error)"},
		},
		"delete.body": {
			{Kind: KindType, Name: "UpdateMemberStatusRequest"},
			{Kind: KindMethod, Owner: "ApiMembersServiceUpdateMembersStatusesRequest", Name: "UpdateMemberStatusRequest", Want: "func(updateMemberStatusRequest UpdateMemberStatusRequest) ApiMembersServiceUpdateMembersStatusesRequest"},
		},
		"delete.body.memberIds":                     {{Kind: KindField, Owner: "UpdateMemberStatusRequest", Name: "MemberIds", Want: "[]string"}},
		"delete.body.status":                        {{Kind: KindField, Owner: "UpdateMemberStatusRequest", Name: "Status", Want: "*MemberStatus"}, {Kind: KindType, Name: "MemberStatus"}},
		"delete.body.status.MEMBER_STATUS_INACTIVE": {{Kind: KindConst, Name: "MEMBERSTATUS_MEMBER_STATUS_INACTIVE"}},
	}
	got := map[string][]Ref{}
	for _, ref := range refs {
		got[ref.Path] = append(got[ref.Path], Ref{Kind: ref.Kind, Owner: ref.Owner, Name: ref.Name, Want: ref.Want})
	}
	for path, refs := range want {
		for _, ref := range refs {
			if !containsRef(got[path], ref) {
				t.Errorf("%s: no %+v in %+v", path, ref, got[path])
			}
		}
	}
	// Get is the List call: no List of its own, and no id in a call.
	for path, refs := range got {
		if strings.HasPrefix(path, "list") {
			t.Errorf("a List ref beside Get: %s %+v", path, refs)
		}
		for _, ref := range refs {
			if ref.Kind == KindMethod && strings.Contains(ref.Want, "memberId string") {
				t.Errorf("%s takes the id: %+v", path, ref)
			}
		}
	}
}

// TestMemberThingSDKNamesWithoutList shows that the reads of Get do not depend on Resource.List, which
// a tool clears when it renders no List of its own.
func TestMemberThingSDKNamesWithoutList(t *testing.T) {
	data, err := os.ReadFile("../model/testdata/memberthing.yaml")
	if err != nil {
		t.Fatal(err)
	}
	doc, err := model.Load(data)
	if err != nil {
		t.Fatal(err)
	}
	r, err := model.BuildWithPolicy(doc, "Member", model.OperationIDs{}, memberPolicy())
	if err != nil {
		t.Fatal(err)
	}
	r.List = nil
	refs, err := Names(r, "Members Service", "example/sdk", nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, path := range []string{"get.page.page_token", "get.filter.login", "get.query.status", "get.response.nextPageToken", "get.response.members"} {
		if !slices.ContainsFunc(refs, func(ref Ref) bool { return ref.Path == path }) {
			t.Errorf("no ref %s without the List", path)
		}
	}
}

func containsRef(refs []Ref, want Ref) bool {
	for _, ref := range refs {
		if ref == want {
			return true
		}
	}
	return false
}
