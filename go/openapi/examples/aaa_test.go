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

	"github.com/stretchr/testify/assert"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	apikeys "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/api_keys_service"
	ipaccess "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/ip_access_service"
	saml "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/saml_configuration_service"
	groups "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/team_permissions_management_service"
)

func TestApiKeys(t *testing.T) {
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewAPIKeysClient(cpc)

	teamID, err := strconv.ParseInt(os.Getenv("TEAM_ID"), 10, 32)
	assertNilAndPrintError(t, err)

	createReq := apikeys.CreateApiKeyRequest{
		Name: "My APM KEY",
		Owner: apikeys.Owner{
			OwnerOneOf1: &apikeys.OwnerOneOf1{
				TeamId: &teamID,
			},
		},
		KeyPermissions: apikeys.CreateApiKeyRequestKeyPermissions{
			Presets:     []string{"APM"},
			Permissions: []string{},
		},
		Hashed: false,
	}

	created, _, err := client.
		ApiKeysServiceCreateApiKey(context.Background()).
		CreateApiKeyRequest(createReq).
		Execute()
	assertNilAndPrintError(t, err)

	newName := "new-name"
	updateReq := apikeys.UpdateApiKeyRequest{
		NewName: &newName,
	}
	_, _, err = client.
		ApiKeysServiceUpdateApiKey(context.Background(), created.KeyId).
		UpdateApiKeyRequest(updateReq).
		Execute()
	assertNilAndPrintError(t, err)

	updated, _, err := client.ApiKeysServiceGetApiKey(context.Background(), created.KeyId).Execute()
	assertNilAndPrintError(t, err)
	assert.Equal(t, newName, updated.KeyInfo.Name)

	_, _, err = client.ApiKeysServiceDeleteApiKey(context.Background(), created.KeyId).Execute()
	assertNilAndPrintError(t, err)
}

func TestIpAccess(t *testing.T) {
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewIPAccessClient(cpc)

	// replace without ID, so it will create a new settings in case it doesn't exist
	// if it exists, it will replace the existing settings
	createReq := ipaccess.ReplaceCompanyIPAccessSettingsRequest{
		IpAccess: []ipaccess.IpAccess{
			{
				Name:    ipaccess.PtrString("Office Network"),
				IpRange: "31.154.215.114/32",
				Enabled: ipaccess.PtrBool(false),
			},
			{
				Name:    ipaccess.PtrString("VPN"),
				IpRange: "198.51.100.0/24",
				Enabled: ipaccess.PtrBool(false),
			},
		},
		EnableCoralogixCustomerSupportAccess: ipaccess.CORALOGIXCUSTOMERSUPPORTACCESS_CORALOGIX_CUSTOMER_SUPPORT_ACCESS_DISABLED.Ptr(),
	}

	createRes, _, err := client.
		IpAccessServiceReplaceCompanyIpAccessSettings(context.Background()).
		ReplaceCompanyIPAccessSettingsRequest(createReq).
		Execute()
	assertNilAndPrintError(t, err)

	getRes, _, err := client.
		IpAccessServiceGetCompanyIpAccessSettings(context.Background()).
		Execute()
	assertNilAndPrintError(t, err)
	assert.Equal(t, createRes.Settings.Id, getRes.Settings.Id)

	updateReq := ipaccess.ReplaceCompanyIPAccessSettingsRequest{
		IpAccess: []ipaccess.IpAccess{
			{
				Name:    ipaccess.PtrString("Office Network"),
				IpRange: "31.154.215.114/32",
				Enabled: ipaccess.PtrBool(false),
			},
			{
				Name:    ipaccess.PtrString("VPN"),
				IpRange: "198.51.100.0/18", // modified
				Enabled: ipaccess.PtrBool(false),
			},
		},
		EnableCoralogixCustomerSupportAccess: ipaccess.CORALOGIXCUSTOMERSUPPORTACCESS_CORALOGIX_CUSTOMER_SUPPORT_ACCESS_DISABLED.Ptr(),
	}

	replaceRes, _, err := client.
		IpAccessServiceReplaceCompanyIpAccessSettings(context.Background()).
		ReplaceCompanyIPAccessSettingsRequest(updateReq).
		Execute()
	assertNilAndPrintError(t, err)

	found := false
	for _, ipAccess := range *replaceRes.Settings.IpAccess {
		if ipAccess.Name != nil && *ipAccess.Name == "VPN" {
			found = true
			assert.Equal(t, ipAccess.IpRange, "198.51.100.0/18")
		}
	}
	if !found {
		t.Error("Expected 'VPN' entry not found in IpAccess")
	}

	_, _, err = client.
		IpAccessServiceDeleteCompanyIpAccessSettings(context.Background()).
		Execute()
	assertNilAndPrintError(t, err)
}

func TestGroups(t *testing.T) {
	t.Skip("Skipping Groups test due to query param issue")
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewGroupsClient(cpc)
	ctx := context.Background()

	teamID, err := strconv.ParseInt(os.Getenv("TEAM_ID"), 10, 64)
	assertNilAndPrintError(t, err)

	groupsResp, _, err := client.
		TeamPermissionsMgmtServiceGetTeamGroups(ctx).
		TeamId(groups.TeamPermissionsMgmtServiceGetTeamGroupsTeamIdParameter{Id: &teamID}).
		Execute()
	assertNilAndPrintError(t, err)
	assert.Greater(t, len(groupsResp.Groups), 0)

	groupName := fmt.Sprintf("Test Group %v", time.Now().UnixMilli())
	groupDesc := "A Test Group"
	createReq := groups.CreateTeamGroupRequest{
		Name:        &groupName,
		Description: &groupDesc,
		TeamId:      &groups.PermissionsV1TeamId{Id: &teamID},
		RoleIds:     []groups.RoleId{{Id: groups.PtrInt64(1)}},
		UserIds:     []groups.V1UserId{},
	}

	createdGroup, _, err := client.
		TeamPermissionsMgmtServiceCreateTeamGroup(ctx).
		CreateTeamGroupRequest(createReq).
		Execute()
	assertNilAndPrintError(t, err)
	assert.NotNil(t, createdGroup.GroupId)
	groupID := createdGroup.GroupId.Id

	gotGroup, _, err := client.
		TeamPermissionsMgmtServiceGetTeamGroup(ctx, *groupID).
		Execute()
	assertNilAndPrintError(t, err)
	assert.Equal(t, groupID, gotGroup.Group.GroupId.Id)

	newName := fmt.Sprintf("Updated Test Group %v", time.Now().UnixMilli())
	updateReq := groups.UpdateTeamGroupRequest{
		GroupId:     &groups.TeamGroupId{Id: groupID},
		Name:        &newName,
		Description: &groupDesc,
	}

	_, _, err = client.
		TeamPermissionsMgmtServiceUpdateTeamGroup(ctx).
		UpdateTeamGroupRequest(updateReq).
		Execute()
	assertNilAndPrintError(t, err)

	_, _, err = client.
		TeamPermissionsMgmtServiceDeleteTeamGroup(ctx, *groupID).
		Execute()
	assertNilAndPrintError(t, err)
}

func TestSamlConfigurationRetrieval(t *testing.T) {
	t.Skip("Skipping SAML Configuration test")
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)
	client := cxsdk.NewSAMLClient(cpc)

	teamIdStr := os.Getenv("TEAM_ID")
	teamId, err := strconv.ParseInt(teamIdStr, 10, 64)
	assertNilAndPrintError(t, err)

	ctx := context.Background()

	spParams, _, err := client.
		SamlConfigurationServiceGetSPParameters(ctx).
		TeamId(teamId).
		Execute()
	assertNilAndPrintError(t, err)
	assert.NotNil(t, spParams)

	config, _, err := client.
		SamlConfigurationServiceGetConfiguration(ctx).
		TeamId(teamId).
		Execute()
	assertNilAndPrintError(t, err)
	assert.NotNil(t, config)

	setReq := saml.SetActiveRequest{
		TeamId:   teamId,
		IsActive: false,
	}
	_, _, err = client.
		SamlConfigurationServiceSetActive(ctx).
		SetActiveRequest(setReq).
		Execute()
	assertNilAndPrintError(t, err)
}
