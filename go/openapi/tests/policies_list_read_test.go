package tests

import (
	"context"
	"testing"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	"github.com/stretchr/testify/require"
)

func TestPoliciesListRead(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewTCOPoliciesClient(cfg)

	listResp, httpResp, err := client.
		PoliciesServiceGetCompanyPolicies(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	if len(listResp.Policies) == 0 {
		t.Skip("no resources to read")
	}

	policyID := ""
	selected := listResp.Policies[0]
	span := false
	if selected.PolicyLogRules != nil {
		policyID = selected.PolicyLogRules.Id
	}
	if selected.PolicySpanRules != nil {
		policyID = selected.PolicySpanRules.Id
		span = true
	}

	if policyID == "" {
		t.Skip("no resources to read")
	}

	getResp, httpResp, err := client.
		PoliciesServiceGetPolicy(context.Background(), policyID).
		Execute()
	actualId := ""
	if span {
		actualId = getResp.GetPolicy().PolicySpanRules.Id
	} else {
		actualId = getResp.GetPolicy().PolicyLogRules.Id
	}

	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, policyID, actualId)
}
