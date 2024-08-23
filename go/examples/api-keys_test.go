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

	cxsdk "github.com/coralogix/coralogix-management-sdk"

	"github.com/stretchr/testify/assert"
)

func TestApiKeys(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	k := cxsdk.NewAPIKeysClient(creator)
	teamId, e := strconv.ParseUint(os.Getenv("TEAM_ID"), 10, 32)
	assert.Nil(t, e)

	key, e := k.Create(context.Background(), &cxsdk.CreateAPIKeyRequest{
		Name: "My APM KEY",
		Owner: &cxsdk.Owner{
			Owner: &cxsdk.OwnerTeamID{
				TeamId: uint32(teamId),
			},
		},
		KeyPermissions: &cxsdk.APIKeyPermissions{
			Presets:     []string{"APM"},
			Permissions: []string{},
		},
	})
	assert.Nil(t, e)

	newName := "new-name"
	k.Update(context.Background(), &cxsdk.UpdateAPIKeyRequest{
		KeyId:   key.KeyId,
		NewName: &newName,
	})

	updated, _ := k.Get(context.Background(), &cxsdk.GetAPIKeyRequest{
		KeyId: key.KeyId,
	})

	assert.Equal(t, updated.KeyInfo.Name, newName)

	k.Delete(context.Background(), &cxsdk.DeleteAPIKeyRequest{
		KeyId: key.KeyId,
	})
}
