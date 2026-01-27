package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
)

func TestAlertsListRead(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewAlertsClient(cfg)

	listResp, httpResp, err := client.
		AlertDefsServiceListAlertDefs(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	if len(listResp.AlertDefs) == 0 {
		t.Skip("no resources to read")
	}

	alertID := listResp.AlertDefs[0].GetId()
	if alertID == "" {
		t.Skip("no resources to read")
	}

	getResp, httpResp, err := client.
		AlertDefsServiceGetAlertDef(context.Background(), alertID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, alertID, *getResp.GetAlertDef().Id)
}
