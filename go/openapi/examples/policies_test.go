// Copyright 2025 Coralogix Ltd.
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

	tcopolicies "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/policies_service"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestPolicies(t *testing.T) {
	ctx := context.Background()
	region, _ := cxsdk.URLFromRegion(cxsdk.RegionFromEnv())
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		region,
		cxsdk.APIKeyFromEnv(),
		true,
	)

	client := cxsdk.NewTCOPoliciesClient(cpc)

	policyName := "Example tco_policy from SDK" + uuid.NewString()
	createReq := tcopolicies.PoliciesServiceCreatePolicyRequest{
		CreatePolicyRequestLogRules: &tcopolicies.CreatePolicyRequestLogRules{
			Name:     policyName,
			Priority: tcopolicies.QUOTAV1PRIORITY_PRIORITY_TYPE_LOW,
			LogRules: &tcopolicies.LogRules{
				Severities: []tcopolicies.QuotaV1Severity{
					tcopolicies.QUOTAV1SEVERITY_SEVERITY_ERROR,
					tcopolicies.QUOTAV1SEVERITY_SEVERITY_CRITICAL,
				},
			},
			ApplicationRule: &tcopolicies.QuotaV1Rule{
				Name:       tcopolicies.PtrString("prod-app"),
				RuleTypeId: tcopolicies.RULETYPEID_RULE_TYPE_ID_IS.Ptr(),
			},
			SubsystemRule: &tcopolicies.QuotaV1Rule{
				Name:       tcopolicies.PtrString("mobile"),
				RuleTypeId: tcopolicies.RULETYPEID_RULE_TYPE_ID_IS_NOT.Ptr(),
			},
			ArchiveRetention: &tcopolicies.ArchiveRetention{
				Id: tcopolicies.PtrString("e1c980d0-c910-4c54-8326-67f3cf95645a"),
			},
		},
	}

	createResp, httpResp, err := client.
		PoliciesServiceCreatePolicy(ctx).
		PoliciesServiceCreatePolicyRequest(createReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	policyID := createResp.GetPolicy().PolicyLogRules.Id

	getResp, _, err := client.
		PoliciesServiceGetPolicy(ctx, policyID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, policyName, getResp.GetPolicy().PolicyLogRules.Name)

	updatedName := "Updated Example tco_policy from SDK" + uuid.NewString()
	updateReq := tcopolicies.PoliciesServiceUpdatePolicyRequest{
		UpdatePolicyRequestLogRules: &tcopolicies.UpdatePolicyRequestLogRules{
			Name:     tcopolicies.PtrString(updatedName),
			Id:       policyID,
			Priority: tcopolicies.QUOTAV1PRIORITY_PRIORITY_TYPE_LOW.Ptr(),
			LogRules: &tcopolicies.LogRules{
				Severities: []tcopolicies.QuotaV1Severity{
					tcopolicies.QUOTAV1SEVERITY_SEVERITY_ERROR,
					tcopolicies.QUOTAV1SEVERITY_SEVERITY_CRITICAL,
				},
			},
			ApplicationRule: &tcopolicies.QuotaV1Rule{
				Name:       tcopolicies.PtrString("prod-app"),
				RuleTypeId: tcopolicies.RULETYPEID_RULE_TYPE_ID_IS.Ptr(),
			},
			SubsystemRule: &tcopolicies.QuotaV1Rule{
				Name:       tcopolicies.PtrString("mobile"),
				RuleTypeId: tcopolicies.RULETYPEID_RULE_TYPE_ID_IS_NOT.Ptr(),
			},
			ArchiveRetention: &tcopolicies.ArchiveRetention{
				Id: tcopolicies.PtrString("e1c980d0-c910-4c54-8326-67f3cf95645a"),
			},
		},
	}
	updateResp, _, err := client.
		PoliciesServiceUpdatePolicy(ctx).
		PoliciesServiceUpdatePolicyRequest(updateReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, updatedName, updateResp.GetPolicy().PolicyLogRules.Name)

	listResp, _, err := client.
		PoliciesServiceGetCompanyPolicies(ctx).
		SourceType(tcopolicies.V1SOURCETYPE_SOURCE_TYPE_LOGS).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	found := false
	for _, policy := range listResp.GetPolicies() {
		if policy.PolicyLogRules != nil && policy.PolicyLogRules.Id == policyID {
			found = true
			break
		}
	}
	require.True(t, found, "created policy not found in the list")

	_, httpResp, err = client.
		PoliciesServiceDeletePolicy(ctx, policyID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
}
