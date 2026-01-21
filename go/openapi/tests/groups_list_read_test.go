package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
)

func TestGroupsListRead(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewGroupsClient(cfg)

	listResp, httpResp, err := client.
		TeamPermissionsMgmtServiceGetTeamGroups(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	if len(listResp.Groups) == 0 {
		t.Skip("no resources to read")
	}

	groupID := listResp.Groups[0].GetGroupId().Id
	if groupID == nil {
		t.Skip("no resources to read")
	}

	getResp, httpResp, err := client.
		TeamPermissionsMgmtServiceGetTeamGroup(context.Background(), *groupID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, groupID, getResp.GetGroup().GroupId.Id)
}
