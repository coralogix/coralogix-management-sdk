package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
)

func TestCustomRolesListRead(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewCustomRolesClient(cfg)

	listResp, httpResp, err := client.
		RoleManagementServiceListCustomRoles(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	if len(listResp.Roles) == 0 {
		t.Skip("no resources to read")
	}

	roleID := listResp.Roles[0].GetRoleId()
	if roleID == 0 {
		t.Skip("no resources to read")
	}

	getResp, httpResp, err := client.
		RoleManagementServiceGetCustomRole(context.Background(), roleID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, roleID, getResp.GetRole().GetRoleId())
}
