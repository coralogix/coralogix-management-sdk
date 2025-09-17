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

	cxsdk "github.com/coralogix/coralogix-management-sdk/go/openapi"
)

func TestApiKeys(t *testing.T) {
	region := os.Getenv(cxsdk.EnvCoralogixRegion)
	if region == "" {
		t.Fatalf("env %s not set", cxsdk.EnvCoralogixRegion)
	}

	apiKey := os.Getenv(cxsdk.EnvCoralogixAPIKey)
	if apiKey == "" {
		t.Fatalf("env %s not set", cxsdk.EnvCoralogixAPIKey)
	}

	client := cxsdk.NewClientBuilder().
		WithRegion(region).
		WithAPIKey(apiKey).
		Build()

	teamID, err := strconv.ParseInt(os.Getenv("TEAM_ID"), 10, 32)
	assert.NoError(t, err)

	createReq := cxsdk.CreateApiKeyRequest{
		Name: "My APM KEY",
		Owner: cxsdk.ComCoralogixapisAaaApikeysV3Owner{
			ComCoralogixapisAaaApikeysV3OwnerOneOf1: &cxsdk.ComCoralogixapisAaaApikeysV3OwnerOneOf1{
				TeamId: &teamID,
			},
		},
		KeyPermissions: cxsdk.ComCoralogixapisAaaApikeysV3CreateApiKeyRequestKeyPermissions{
			Presets:     []string{"APM"},
			Permissions: []string{},
		},
		Hashed: false,
	}

	created, _, err := client.APIKeysServiceAPI.
		ApiKeysServiceCreateApiKey(context.Background()).
		CreateApiKeyRequest(createReq).
		Execute()
	assert.NoError(t, err)

	newName := "new-name"
	updateReq := cxsdk.UpdateApiKeyRequest{
		NewName: &newName,
	}
	_, _, err = client.APIKeysServiceAPI.
		ApiKeysServiceUpdateApiKey(context.Background(), created.KeyId).
		UpdateApiKeyRequest(updateReq).
		Execute()
	assert.NoError(t, err)

	updated, _, err := client.APIKeysServiceAPI.
		ApiKeysServiceGetApiKey(context.Background(), created.KeyId).
		Execute()
	assert.NoError(t, err)
	assert.Equal(t, newName, updated.KeyInfo.Name)

	_, _, err = client.APIKeysServiceAPI.
		ApiKeysServiceDeleteApiKey(context.Background(), created.KeyId).
		Execute()
	assert.NoError(t, err)
}
