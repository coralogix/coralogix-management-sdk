// Copyright 2024 Coralogix Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package examples

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	apikeys "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/api_keys_service"
	ipaccess "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/ip_access_service"
	customroles "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/role_management_service"
	scopes "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/scopes_service"
	groups "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/team_permissions_management_service"
)

func TestApiKeys(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewAPIKeysClient(cfg)

	teamIDEnv := os.Getenv("TEAM_ID")
	if teamIDEnv == "" {
		t.Fatal("TEAM_ID environment variable is not set")
	}
	teamID, err := strconv.ParseInt(teamIDEnv, 10, 64)
	require.NoError(t, err)

	name := "My APM KEY"
	createReq := apikeys.CreateApiKeyRequest{
		Name: &name,
		Owner: &apikeys.Owner{
			OwnerTeamId: &apikeys.OwnerTeamId{
				TeamId: &teamID,
			},
		},
		KeyPermissions: &apikeys.CreateApiKeyRequestKeyPermissions{
			Presets:     []string{"APM"},
			Permissions: []string{"ALERTS-MAP:READ"},
		},
		Hashed: apikeys.PtrBool(false),
	}

	created, httpResp, err := client.
		ApiKeysServiceCreateApiKey(context.Background()).
		CreateApiKeyRequest(createReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	newName := "new-name"
	updateReq := apikeys.UpdateApiKeyRequest{
		NewName: &newName,
	}
	_, httpResp, err = client.
		ApiKeysServiceUpdateApiKey(context.Background(), *created.KeyId).
		UpdateApiKeyRequest(updateReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	updated, httpResp, err := client.ApiKeysServiceGetApiKey(context.Background(), *created.KeyId).Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, newName, *updated.KeyInfo.Name)

	_, httpResp, err = client.ApiKeysServiceDeleteApiKey(context.Background(), *created.KeyId).Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
}

func TestIpAccess(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewIPAccessClient(cfg)

	enabled := false
	ipRange1 := "31.154.215.114/32"
	ipRange2 := "198.51.100.0/24"
	// replace without ID, so it will create a new settings in case it doesn't exist
	// if it exists, it will replace the existing settings
	createReq := ipaccess.ReplaceCompanyIPAccessSettingsRequest{
		IpAccess: []ipaccess.IpAccess{
			{
				Name:    ipaccess.PtrString("Office Network"),
				IpRange: &ipRange1,
				Enabled: &enabled,
			},
			{
				Name:    ipaccess.PtrString("VPN"),
				IpRange: &ipRange2,
				Enabled: &enabled,
			},
		},
		EnableCoralogixCustomerSupportAccess: ipaccess.CORALOGIXCUSTOMERSUPPORTACCESS_CORALOGIX_CUSTOMER_SUPPORT_ACCESS_DISABLED.Ptr(),
	}

	createRes, httpResp, err := client.
		IpAccessServiceReplaceCompanyIpAccessSettings(context.Background()).
		ReplaceCompanyIPAccessSettingsRequest(createReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	getRes, httpResp, err := client.
		IpAccessServiceGetCompanyIpAccessSettings(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, createRes.Settings.Id, getRes.Settings.Id)

	updateReq := ipaccess.ReplaceCompanyIPAccessSettingsRequest{
		IpAccess: []ipaccess.IpAccess{
			{
				Name:    ipaccess.PtrString("Office Network"),
				IpRange: ipaccess.PtrString("31.154.215.114/32"),
				Enabled: &enabled,
			},
			{
				Name:    ipaccess.PtrString("VPN"),
				IpRange: ipaccess.PtrString("198.51.100.0/18"), // modified
				Enabled: &enabled,
			},
		},
		EnableCoralogixCustomerSupportAccess: ipaccess.CORALOGIXCUSTOMERSUPPORTACCESS_CORALOGIX_CUSTOMER_SUPPORT_ACCESS_DISABLED.Ptr(),
	}

	replaceRes, httpResp, err := client.
		IpAccessServiceReplaceCompanyIpAccessSettings(context.Background()).
		ReplaceCompanyIPAccessSettingsRequest(updateReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	found := false
	for _, ipAccess := range *replaceRes.Settings.IpAccess {
		if ipAccess.Name != nil && *ipAccess.Name == "VPN" {
			found = true
			require.Equal(t, *ipAccess.IpRange, "198.51.100.0/18")
		}
	}
	if !found {
		t.Error("Expected 'VPN' entry not found in IpAccess")
	}

	_, httpResp, err = client.
		IpAccessServiceDeleteCompanyIpAccessSettings(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
}

func TestGroups(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewGroupsClient(cfg)

	ctx := context.Background()
	teamIDEnv := os.Getenv("TEAM_ID")
	if teamIDEnv == "" {
		t.Fatal("TEAM_ID environment variable is not set")
	}
	teamID, err := strconv.ParseInt(teamIDEnv, 10, 64)
	require.NoError(t, err)

	// groupsResp, httpResp, err := client.
	// 	TeamPermissionsMgmtServiceGetTeamGroups(ctx).
	// 	TeamId(groups.TeamPermissionsMgmtServiceGetTeamGroupsTeamIdParameter{Id: &teamID}).
	// 	Execute()
	// require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	// require.Greater(t, len(groupsResp.Groups), 0)

	groupName := fmt.Sprintf("Test Group %v", time.Now().UnixMilli())
	groupDesc := "A Test Group"
	createReq := groups.CreateTeamGroupRequest{
		Name:        &groupName,
		Description: &groupDesc,
		TeamId:      &groups.PermissionsV1TeamId{Id: &teamID},
		RoleIds:     []groups.RoleId{{Id: groups.PtrInt64(1)}},
		UserIds:     []groups.V1UserId{},
	}

	createdGroup, httpResp, err := client.
		TeamPermissionsMgmtServiceCreateTeamGroup(ctx).
		CreateTeamGroupRequest(createReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotNil(t, createdGroup.GroupId)
	groupID := createdGroup.GroupId.Id

	//gotGroup, httpResp, err := client.
	//	TeamPermissionsMgmtServiceGetTeamGroup(ctx, *groupID).
	//	Execute()
	//require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	//require.Equal(t, groupID, gotGroup.Group.GroupId.Id)

	newName := fmt.Sprintf("Updated Test Group %v", time.Now().UnixMilli())
	updateReq := groups.UpdateTeamGroupRequest{
		GroupId:     &groups.TeamGroupId{Id: groupID},
		Name:        &newName,
		Description: &groupDesc,
	}

	_, httpResp, err = client.
		TeamPermissionsMgmtServiceUpdateTeamGroup(ctx).
		UpdateTeamGroupRequest(updateReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	_, httpResp, err = client.
		TeamPermissionsMgmtServiceDeleteTeamGroup(ctx, *groupID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
}

func TestCustomRoles(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewCustomRolesClient(cfg)

	ctx := context.Background()
	teamIDEnv := os.Getenv("TEAM_ID")
	if teamIDEnv == "" {
		t.Fatal("TEAM_ID environment variable is not set")
	}
	teamID, err := strconv.ParseInt(teamIDEnv, 10, 64)
	require.NoError(t, err)

	createReq := customroles.RoleManagementServiceCreateRoleRequest{
		CreateRoleRequestParentRoleName: &customroles.CreateRoleRequestParentRoleName{
			Name:           customroles.PtrString("custom-role-sample"),
			Description:    customroles.PtrString("This is a sample custom role"),
			ParentRoleName: customroles.PtrString("Standard User"),
			Permissions: []string{
				"team-actions:UpdateConfig",
				"TEAM-CUSTOM-API-KEYS:READCONFIG",
			},
			TeamId: &teamID,
		},
	}

	created, httpResp, err := client.
		RoleManagementServiceCreateRole(ctx).
		RoleManagementServiceCreateRoleRequest(createReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	roleID := created.Id

	newDesc := "Updated description"
	updateReq := customroles.RoleManagementServiceUpdateRoleRequest{
		NewDescription: &newDesc,
	}

	_, httpResp, err = client.
		RoleManagementServiceUpdateRole(ctx, *roleID).
		RoleManagementServiceUpdateRoleRequest(updateReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	getResp, httpResp, err := client.
		RoleManagementServiceGetCustomRole(ctx, *roleID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, newDesc, getResp.Role.GetDescription())

	listResp, httpResp, err := client.RoleManagementServiceListCustomRoles(ctx).
		TeamId(teamID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.GreaterOrEqual(t, len(listResp.Roles), 1)

	_, httpResp, err = client.
		RoleManagementServiceDeleteRole(ctx, *roleID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
}

func TestScopes(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewScopesClient(cfg)

	ctx := context.Background()
	description := "Data Access Rule intended for testing"
	defaultExpr := "<v1>true"
	createReq := scopes.CreateScopeRequest{
		DisplayName:       "Test Data Access Rule",
		Description:       &description,
		DefaultExpression: &defaultExpr,
		Filters: []scopes.ScopesV1Filter{
			{
				EntityType: scopes.V1ENTITYTYPE_ENTITY_TYPE_LOGS.Ptr(),
				Expression: scopes.PtrString("<v1> foo == 'bar'"),
			},
		},
	}

	created, httpResp, err := client.
		ScopesServiceCreateScope(ctx).
		CreateScopeRequest(createReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	scopeID := created.Scope.Id

	newDisplayName := "Updated Test Data Access Rule"
	updateReq := scopes.UpdateScopeRequest{
		Id:                *scopeID,
		DisplayName:       newDisplayName,
		DefaultExpression: *created.Scope.DefaultExpression,
		Filters:           created.Scope.Filters,
	}

	updated, httpResp, err := client.
		ScopesServiceUpdateScope(ctx).
		UpdateScopeRequest(updateReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotNil(t, updated)

	getResp, httpResp, err := client.
		ScopesServiceGetTeamScopesByIds(ctx).
		Ids([]string{*scopeID}).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Greater(t, len(getResp.Scopes), 0)
	require.Equal(t, newDisplayName, *getResp.Scopes[0].DisplayName)

	listResp, httpResp, err := client.
		ScopesServiceGetTeamScopes(ctx).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.GreaterOrEqual(t, len(listResp.Scopes), 1)

	_, httpResp, err = client.
		ScopesServiceDeleteScope(ctx, *scopeID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	// Skipping because of required `scopes` fields not returned in the list response
	//listAfterDelete, httpResp, err := client.
	//	ScopesServiceGetTeamScopes(ctx).
	//	Execute()
	//require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	//
	//for _, s := range listAfterDelete.Scopes {
	//	require.NotEqual(t, scopeID, s.Id)
	//}
}
