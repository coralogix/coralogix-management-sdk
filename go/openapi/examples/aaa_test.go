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
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	apikeys "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/api_keys_service"
	ipaccess "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/ip_access_service"
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
		Owner: apikeys.CreateApiKeyRequestOwner{
			CreateApiKeyRequestOwnerOneOf1: &apikeys.CreateApiKeyRequestOwnerOneOf1{
				TeamId: &teamID,
			},
		},
		KeyPermissions: apikeys.KeyPermissions{
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
	t.Skip("Skipping IP Access test")
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewIPAccessClient(cpc)

	// replace without ID, so it will create a new settings in case it doesn't exist
	// if it exists, it will replace the existing settings
	createReq := ipaccess.ReplaceCompanyIPAccessSettingsRequest{
		IpAccess: []ipaccess.IPAccess{
			{
				Name:    ipaccess.PtrString("Office Network"),
				IpRange: "31.154.215.114/32",
				Enabled: false,
			},
			{
				Name:    ipaccess.PtrString("VPN"),
				IpRange: "198.51.100.0/24",
				Enabled: false,
			},
		},
		EnableCoralogixCustomerSupportAccess: ipaccess.PtrString("DISABLED"),
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
		IpAccess: []ipaccess.IPAccess{
			{
				Name:    ipaccess.PtrString("Office Network"),
				IpRange: "31.154.215.114/32",
				Enabled: false,
			},
			{
				Name:    ipaccess.PtrString("VPN"),
				IpRange: "198.51.100.0/18", // modified
				Enabled: false,
			},
		},
		EnableCoralogixCustomerSupportAccess: ipaccess.PtrString("DISABLED"),
	}

	replaceRes, _, err := client.
		IpAccessServiceReplaceCompanyIpAccessSettings(context.Background()).
		ReplaceCompanyIPAccessSettingsRequest(updateReq).
		Execute()
	assertNilAndPrintError(t, err)

	found := false
	for _, ipAccess := range replaceRes.Settings.IpAccess {
		if ipAccess.Value.Name != nil && *ipAccess.Value.Name == "VPN" {
			found = true
			assert.Equal(t, "198.51.100.0/18", ipAccess.Value.IpRange)
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
