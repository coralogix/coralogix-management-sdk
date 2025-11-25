// Copyright 2024 Coralogix Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package examples

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	rulegroups "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/rule_groups_service"
)

func TestRuleGroups(t *testing.T) {
	region, _ := cxsdk.URLFromRegion(cxsdk.RegionFromEnv())
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		region,
		cxsdk.APIKeyFromEnv(),
		true,
	)

	client := cxsdk.NewRuleGroupsClient(cpc)
	ctx := context.Background()

	groupName := "block-rule"
	createReq := rulegroups.RuleGroupsServiceCreateRuleGroupRequest{
		Name:          &groupName,
		Description:   rulegroups.PtrString("rule-group from sdk"),
		RuleSubgroups: getSubgroups(),
	}

	createResp, httpResp, err := client.
		RuleGroupsServiceCreateRuleGroup(ctx).
		RuleGroupsServiceCreateRuleGroupRequest(createReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotNil(t, createResp)
	require.NotEmpty(t, createResp.RuleGroup)

	groupID := createResp.RuleGroup.Id

	listResp, httpResp, err := client.
		RuleGroupsServiceListRuleGroups(ctx).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotNil(t, listResp)
	found := false
	for _, g := range listResp.RuleGroups {
		if g.Id != nil && *g.Id == *groupID {
			found = true
			break
		}
	}
	require.True(t, found, "created group not found in list")

	getResp, httpResp, err := client.
		RuleGroupsServiceGetRuleGroup(ctx, *groupID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotNil(t, getResp)
	require.Equal(t, groupName, *getResp.RuleGroup.Name)

	updatedDesc := "rule-group updated via OpenAPI test"
	updateReq := rulegroups.RuleGroupsServiceCreateRuleGroupRequest{
		Name:          &groupName,
		Description:   &updatedDesc,
		RuleSubgroups: getSubgroups(),
	}

	updateResp, httpResp, err := client.
		RuleGroupsServiceUpdateRuleGroup(ctx, *groupID).
		RuleGroupsServiceCreateRuleGroupRequest(updateReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotNil(t, updateResp)
	require.Equal(t, groupID, updateResp.RuleGroup.Id)
	require.Equal(t, updatedDesc, *updateResp.RuleGroup.Description)

	_, httpResp, err = client.
		RuleGroupsServiceDeleteRuleGroup(ctx, *groupID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	_, httpResp, err = client.
		RuleGroupsServiceGetRuleGroup(ctx, *groupID).
		Execute()
	require.Error(t, err, "expected error after deleting rule group")
}

func getSubgroups() []rulegroups.CreateRuleGroupRequestCreateRuleSubgroup {
	return []rulegroups.CreateRuleGroupRequestCreateRuleSubgroup{
		{
			Rules: []rulegroups.CreateRuleGroupRequestCreateRuleSubgroupCreateRule{
				{
					Name:        rulegroups.PtrString("Block 28000"),
					Description: rulegroups.PtrString("Block 2800 pg error"),
					SourceField: rulegroups.PtrString("text"),
					Order:       rulegroups.PtrInt64(int64(1)),
					Parameters: &rulegroups.RuleParameters{
						RuleParametersBlockParameters: &rulegroups.RuleParametersBlockParameters{
							BlockParameters: &rulegroups.BlockParameters{
								KeepBlockedLogs: rulegroups.PtrBool(false),
								Rule:            rulegroups.PtrString(`sql_error_code\s*=\s*28000`),
							},
						},
					},
				},
			},
			Order: rulegroups.PtrInt64(int64(1)),
		},
	}
}
