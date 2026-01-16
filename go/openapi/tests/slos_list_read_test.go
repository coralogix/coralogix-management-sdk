package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
)

func TestSLOsListRead(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewSLOsClient(cfg)

	listResp, httpResp, err := client.
		SlosServiceListSlos(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	if len(listResp.Slos) == 0 {
		t.Skip("no resources to read")
	}

	sloID := listResp.Slos[0].GetId()
	if sloID == "" {
		t.Skip("no resources to read")
	}

	getResp, httpResp, err := client.
		SlosServiceGetSlo(context.Background(), sloID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, sloID, getResp.GetSlo().GetId())
}
