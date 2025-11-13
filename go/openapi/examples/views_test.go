// Copyright 2024 Coralogix Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied. See the License for the specific language
// governing permissions and limitations under the License.

package examples

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	viewsfolders "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/folders_for_views_service"
	views "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/views_service"
)

func TestViews(t *testing.T) {
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewViewsClient(cpc)

	createReq := views.ViewFolder{
		Name: fmt.Sprintf("TestView-%v", uuid.NewString()),
		SearchQuery: &views.SearchQuery{
			Query: "source logs | filter $l.applicationname == 'default'",
		},
		TimeSelection: views.TimeSelection{
			TimeSelectionQuickSelection: &views.TimeSelectionQuickSelection{
				QuickSelection: &views.QuickTimeSelection{
					Seconds: 86400,
				},
			},
		},
	}

	created, httpResp, err := client.
		ViewsServiceCreateView(context.Background()).
		ViewFolder(createReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotNil(t, created.Id)

	updatedName := fmt.Sprintf("TestViewUpdated-%v", uuid.NewString())
	updated := views.View1{
		Name:          updatedName,
		SearchQuery:   created.SearchQuery,
		TimeSelection: created.TimeSelection,
		Filters:       created.Filters,
	}

	_, httpResp, err = client.
		ViewsServiceReplaceView(context.Background(), created.Id).
		View1(updated).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	got, httpResp, err := client.
		ViewsServiceGetView(context.Background(), created.Id).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, updatedName, got.Name)

	_, httpResp, err = client.
		ViewsServiceDeleteView(context.Background(), created.Id).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
}

func TestViewsFolders(t *testing.T) {
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewViewsFoldersClient(cpc)

	before, httpResp, err := client.
		ViewsFoldersServiceListViewFolders(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	initialCount := len(before.Folders)

	createReq := viewsfolders.CreateViewFolderRequest{
		Name: viewsfolders.PtrString(fmt.Sprintf("GoTestViewFolder-%s", uuid.NewString())),
	}
	created, httpResp, err := client.
		ViewsFoldersServiceCreateViewFolder(context.Background()).
		CreateViewFolderRequest(createReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotNil(t, created)
	require.NotNil(t, created.Id)

	updatedName := fmt.Sprintf("GoTestViewFolderUpdated-%s", uuid.NewString())
	updateReq := viewsfolders.ViewFolder1{
		Id:   created.Id,
		Name: updatedName,
	}
	_, httpResp, err = client.
		ViewsFoldersServiceReplaceViewFolder(context.Background()).
		ViewFolder1(updateReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	got, httpResp, err := client.
		ViewsFoldersServiceGetViewFolder(context.Background(), *created.Id).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, updatedName, got.Name)

	after, httpResp, err := client.
		ViewsFoldersServiceListViewFolders(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, initialCount+1, len(after.Folders))

	_, httpResp, err = client.
		ViewsFoldersServiceDeleteViewFolder(context.Background(), *created.Id).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	final, httpResp, err := client.
		ViewsFoldersServiceListViewFolders(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, initialCount, len(final.Folders))
}
