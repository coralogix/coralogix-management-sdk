package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
)

func TestViewsFoldersListRead(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewViewsFoldersClient(cfg)

	listResp, httpResp, err := client.
		ViewsFoldersServiceListViewFolders(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	if len(listResp.Folders) == 0 {
		t.Skip("no resources to read")
	}

	folderID := listResp.Folders[0].GetId()
	if folderID == "" {
		t.Skip("no resources to read")
	}

	getResp, httpResp, err := client.
		ViewsFoldersServiceGetViewFolder(context.Background(), folderID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, folderID, getResp.GetFolder().GetId())
}
