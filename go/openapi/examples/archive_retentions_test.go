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
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	retentions "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/retentions_service"
)

func TestArchiveRetentions(t *testing.T) {
	region, _ := cxsdk.URLFromRegion(cxsdk.RegionFromEnv())
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		region,
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewArchiveRetentionsClient(cpc)
	ctx := context.Background()

	getResp, httpResp, err := client.
		RetentionsServiceGetRetentions(ctx).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotNil(t, getResp)
	require.NotEmpty(t, getResp.Retentions, "expected at least one retention")

	var updateElems []retentions.RetentionUpdateElement
	for _, r := range getResp.Retentions {
		elem := retentions.RetentionUpdateElement{}
		if r.Id != nil {
			elem.Id = r.Id
		}
		if r.Name != nil {
			elem.Name = r.Name
		}
		updateElems = append(updateElems, elem)
	}

	if len(updateElems) > 1 && updateElems[1].Name != nil {
		newName := *updateElems[1].Name + "_updated"
		updateElems[1].Name = &newName
	}

	updateReq := retentions.UpdateRetentionsRequest{
		RetentionUpdateElements: updateElems,
	}

	updateResp, httpResp, err := client.
		RetentionsServiceUpdateRetentions(ctx).
		UpdateRetentionsRequest(updateReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotNil(t, updateResp)
	require.NotEmpty(t, updateResp.Retentions)

	verifyResp, httpResp, err := client.
		RetentionsServiceGetRetentions(ctx).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotNil(t, verifyResp)

	if len(verifyResp.Retentions) > 1 {
		require.Contains(t, verifyResp.Retentions[1].GetName(), "_updated")
	}

	activateResp, httpResp, err := client.
		RetentionsServiceActivateRetentions(ctx).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotNil(t, activateResp)
}
