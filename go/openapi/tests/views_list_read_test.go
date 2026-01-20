package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
)

func TestViewsListRead(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewViewsClient(cfg)

	listResp, httpResp, err := client.
		ViewsServiceListViews(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	if len(listResp.Views) == 0 {
		t.Skip("no resources to read")
	}

	viewID := listResp.Views[0].GetId()
	if viewID == "" {
		t.Skip("no resources to read")
	}

	getResp, httpResp, err := client.
		ViewsServiceGetView(context.Background(), viewID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, viewID, getResp.GetId())
}
