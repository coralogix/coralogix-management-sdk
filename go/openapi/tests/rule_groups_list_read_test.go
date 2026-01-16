package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
)

func TestRuleGroupsListRead(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewRuleGroupsClient(cfg)

	listResp, httpResp, err := client.
		RuleGroupsServiceListRuleGroups(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	if len(listResp.RuleGroups) == 0 {
		t.Skip("no resources to read")
	}

	groupID := listResp.RuleGroups[0].GetId()
	if groupID == "" {
		t.Skip("no resources to read")
	}

	getResp, httpResp, err := client.
		RuleGroupsServiceGetRuleGroup(context.Background(), groupID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, groupID, getResp.GetRuleGroup().GetId())
}
