package tests

import (
	"testing"
)

func TestPoliciesListRead(t *testing.T) {
	t.Skip("later")
	// cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	// client := cxsdk.NewTCOPoliciesClient(cfg)

	// listResp, httpResp, err := client.
	// 	PoliciesServiceGetCompanyPolicies(context.Background()).
	// 	Execute()
	// require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	// if len(listResp.Policies) == 0 {
	// 	t.Skip("no resources to read")
	// }

	// policyID := listResp.Policies[0].GetId()
	// if policyID == "" {
	// 	t.Skip("no resources to read")
	// }

	// getResp, httpResp, err := client.
	// 	PoliciesServiceGetPolicy(context.Background(), policyID).
	// 	Execute()
	// require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	// require.Equal(t, policyID, getResp.GetPolicy().GetId())
}
