package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
)

func TestDashboardsListRead(t *testing.T) {
	t.Skip("Dashboards cannot be deserialized right now")
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewDashboardClient(cfg)

	listResp, httpResp, err := client.
		DashboardCatalogServiceGetDashboardCatalog(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	if len(listResp.Items) == 0 {
		t.Skip("no resources to read")
	}

	dashboardID := listResp.Items[0].GetId()
	if dashboardID == "" {
		t.Skip("no resources to read")
	}

	getResp, httpResp, err := client.
		DashboardsServiceGetDashboard(context.Background(), dashboardID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	actual := getResp.GetDashboard().GetActualInstanceValue()
	if actual == nil {
		t.Skip("no resources to read")
	}
	idProvider, ok := actual.(interface{ GetId() string })
	require.True(t, ok)
	require.Equal(t, dashboardID, idProvider.GetId())
}
