package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
)

func TestIPAccessListRead(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewIPAccessClient(cfg)

	listResp, httpResp, err := client.
		IpAccessServiceGetCompanyIpAccessSettings(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	settingsID := listResp.GetSettings().GetId()
	if settingsID == "" {
		t.Skip("no resources to read")
	}

	getResp, httpResp, err := client.
		IpAccessServiceGetCompanyIpAccessSettings(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, settingsID, getResp.GetSettings().GetId())
}
