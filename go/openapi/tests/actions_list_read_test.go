package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
)

func TestActionsListRead(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewActionsClient(cfg)

	listResp, httpResp, err := client.
		ActionsServiceListActions(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	if len(listResp.Actions) == 0 {
		t.Skip("no resources to read")
	}

	actionID := listResp.Actions[0].GetId()
	if actionID == "" {
		t.Skip("no resources to read")
	}

	getResp, httpResp, err := client.
		ActionsServiceGetAction(context.Background(), actionID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, actionID, getResp.GetAction().Id)
}
