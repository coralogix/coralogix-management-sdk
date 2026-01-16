package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
)

func TestAPIKeysListRead(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewAPIKeysClient(cfg)

	listResp, httpResp, err := client.
		ApiKeysServiceGetSendDataApiKeys(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	if len(listResp.Keys) == 0 {
		t.Skip("no resources to read")
	}

	keyID := listResp.Keys[0].GetId()
	if keyID == "" {
		t.Skip("no resources to read")
	}

	getResp, httpResp, err := client.
		ApiKeysServiceGetApiKey(context.Background(), keyID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, keyID, getResp.GetKeyInfo().GetId())
}
