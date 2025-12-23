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
	quota "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/quota_allocation_rule_set_service"
)

func TestQuotaAllocationRuleSet(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewQuotasClient(cfg)
	ctx := context.Background()

	createReq := quota.CreateQuotaAllocationRuleSetRequest{
		RuleSet: quota.QuotaAllocationEntityTypeRuleSet{
			Rules: []quota.QuotaAllocationEntityTypeRule{
				{
					EntityType:  "logs",
					Allocation:  70,
					CanOverflow: true,
					Enabled:     true,
				},
				{
					EntityType:  "metrics",
					Allocation:  30,
					CanOverflow: false,
					Enabled:     true,
				},
			},
		},
	}

	createResp, httpResp, err := client.
		QuotaAllocationRuleSetServiceCreateQuotaAllocationRuleSet(ctx).
		CreateQuotaAllocationRuleSetRequest(createReq).
		Execute()

	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotNil(t, createResp)
	require.NotNil(t, createResp.RuleSet.Id)

	ruleSetID := *createResp.RuleSet.Id
	getResp, httpResp, err := client.
		QuotaAllocationRuleSetServiceGetQuotaAllocationRuleSet(ctx).
		Id(ruleSetID).
		Execute()

	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotNil(t, getResp)
	require.Equal(t, ruleSetID, *getResp.RuleSet.Id)
	require.Len(t, getResp.RuleSet.Rules, 2)

	replaceReq := quota.ReplaceQuotaAllocationRuleSetRequest{
		RuleSet: quota.QuotaAllocationEntityTypeRuleSet{
			Id: &ruleSetID,
			Rules: []quota.QuotaAllocationEntityTypeRule{
				{
					EntityType:  "logs",
					Allocation:  60,
					CanOverflow: true,
					Enabled:     true,
				},
				{
					EntityType:  "metrics",
					Allocation:  40,
					CanOverflow: false,
					Enabled:     true,
				},
			},
		},
	}

	replaceResp, httpResp, err := client.
		QuotaAllocationRuleSetServiceReplaceQuotaAllocationRuleSet(ctx).
		ReplaceQuotaAllocationRuleSetRequest(replaceReq).
		Execute()

	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotNil(t, replaceResp)

	foundLogs := false
	for _, rule := range replaceResp.RuleSet.Rules {
		if rule.EntityType == "logs" {
			foundLogs = true
			require.Equal(t, float32(60), rule.Allocation)
		}
	}
	require.True(t, foundLogs, "expected updated logs rule to exist")

	_, httpResp, err = client.
		QuotaAllocationRuleSetServiceDeleteQuotaAllocationRuleSet(ctx).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
}
