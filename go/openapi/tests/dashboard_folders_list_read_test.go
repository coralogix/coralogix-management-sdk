package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
)

func TestDashboardFoldersListRead(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewDashboardFoldersClient(cfg)

	listResp, httpResp, err := client.
		DashboardFoldersServiceListDashboardFolders(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	if len(listResp.Folder) == 0 {
		t.Skip("no resources to read")
	}

	folderID := listResp.Folder[0].GetId()
	if folderID == "" {
		t.Skip("no resources to read")
	}

	getResp, httpResp, err := client.
		DashboardFoldersServiceGetDashboardFolder(context.Background(), folderID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, folderID, getResp.GetFolder().Id)
}
