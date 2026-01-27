package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
)

func TestConnectorsListRead(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewConnectorsClient(cfg)

	listResp, httpResp, err := client.
		ConnectorsServiceListConnectors(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	if len(listResp.Connectors) == 0 {
		t.Skip("no resources to read")
	}

	connectorID := listResp.Connectors[0].GetId()
	if connectorID == "" {
		t.Skip("no resources to read")
	}

	getResp, httpResp, err := client.
		ConnectorsServiceGetConnector(context.Background(), connectorID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, connectorID, *getResp.GetConnector().Id)
}
