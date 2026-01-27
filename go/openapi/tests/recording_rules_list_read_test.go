package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
)

func TestRecordingRulesListRead(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewRecordingRulesClient(cfg)

	listResp, httpResp, err := client.
		RuleGroupSetsList(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	if len(listResp.Sets) == 0 {
		t.Skip("no resources to read")
	}

	setID := listResp.Sets[0].GetId()
	if setID == "" {
		t.Skip("no resources to read")
	}

	getResp, httpResp, err := client.
		RuleGroupSetsFetch(context.Background(), setID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, setID, getResp.GetId())
}
