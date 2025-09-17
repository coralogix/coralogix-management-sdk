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
	"github.com/stretchr/testify/assert"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	views "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/views_service"
)

func TestViews(t *testing.T) {
	t.Skip("Skipping until response bug is fixed")
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewViewsClient(cpc)

	createReq := views.ViewFolder1{
		Name: fmt.Sprintf("TestView-%v", uuid.NewString()),
		SearchQuery: &views.SearchQuery{
			Query: "source logs | filter $l.applicationname == 'default'",
		},
		TimeSelection: views.ViewTimeSelection{
			ViewTimeSelectionOneOf: &views.ViewTimeSelectionOneOf{
				QuickSelection: &views.ViewTimeSelectionOneOfQuickSelection{
					Seconds: 86400,
				},
			},
		},
	}

	created, _, err := client.
		ViewsServiceCreateView(context.Background()).
		ViewFolder1(createReq).
		Execute()
	assertNilAndPrintError(t, err)
	assert.NotNil(t, created.Id)

	updated := views.View1{
		Name:          fmt.Sprintf("TestViewUpdated-%v", uuid.NewString()),
		SearchQuery:   created.SearchQuery,
		TimeSelection: created.TimeSelection,
		Filters:       created.Filters,
	}

	_, _, err = client.
		ViewsServiceReplaceView(context.Background(), created.Id).
		View1(updated).
		Execute()
	assertNilAndPrintError(t, err)

	got, _, err := client.
		ViewsServiceGetView(context.Background(), created.Id).
		Execute()
	assertNilAndPrintError(t, err)
	assert.Equal(t, "GoTestViewUpdated", got.Name)

	_, _, err = client.
		ViewsServiceDeleteView(context.Background(), created.Id).
		Execute()
	assertNilAndPrintError(t, err)
}

//func TestViewFolders(t *testing.T) {
//	t.Skip("Skipping until response bug is fixed")
//	client := cxsdk.NewClientBuilder().WithRegionFromEnv().WithAPIKeyFromEnv().Build()
//	ctx := context.Background()
//
//	before, _, err := client.FoldersForViewsServiceAPI.
//		ViewsFoldersServiceListViewFolders(ctx).
//		Execute()
//	assertNilAndPrintError(t, err)
//	initialCount := len(before.Folders)
//
//	createReq := cxsdk.CreateViewFolderRequest{
//		Name: cxsdk.PtrString(fmt.Sprintf("GoTestViewFolder-%s", uuid.NewString())),
//	}
//	created, _, err := client.FoldersForViewsServiceAPI.
//		ViewsFoldersServiceCreateViewFolder(ctx).
//		CreateViewFolderRequest(createReq).
//		Execute()
//	assertNilAndPrintError(t, err)
//	assert.NotNil(t, created.Folder)
//	assert.NotNil(t, created.Folder.Id)
//
//	updatedName := fmt.Sprintf("GoTestViewFolderUpdated-%s", uuid.NewString())
//	updateReq := cxsdk.ViewFolder{
//		Name: updatedName,
//	}
//	_, _, err = client.FoldersForViewsServiceAPI.
//		ViewsFoldersServiceReplaceViewFolder(ctx, *created.Folder.Id).
//		ViewFolder(updateReq).
//		Execute()
//	assertNilAndPrintError(t, err)
//
//	got, _, err := client.FoldersForViewsServiceAPI.
//		ViewsFoldersServiceGetViewFolder(ctx, *created.Folder.Id).
//		Execute()
//	assertNilAndPrintError(t, err)
//	assert.Equal(t, updatedName, got.Folder.Name)
//
//	after, _, err := client.FoldersForViewsServiceAPI.
//		ViewsFoldersServiceListViewFolders(ctx).
//		Execute()
//	assertNilAndPrintError(t, err)
//	assert.Equal(t, initialCount+1, len(after.Folders))
//
//	_, _, err = client.FoldersForViewsServiceAPI.
//		ViewsFoldersServiceDeleteViewFolder(ctx, *created.Folder.Id).
//		Execute()
//	assertNilAndPrintError(t, err)
//
//	final, _, err := client.FoldersForViewsServiceAPI.
//		ViewsFoldersServiceListViewFolders(ctx).
//		Execute()
//	assertNilAndPrintError(t, err)
//	assert.Equal(t, initialCount, len(final.Folders))
//}
