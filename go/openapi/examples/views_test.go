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
			TimeSelectionOneOf: &views.TimeSelectionOneOf{
				QuickSelection: &views.QuickTimeSelection{
					Seconds: 86400,
				},
			},
		},
	}

	created, _, err := client.
		ViewsServiceCreateView(context.Background()).
		ViewFolder(createReq).
		Execute()
	assertNilAndPrintError(t, err)
	assert.NotNil(t, created.Id)

	updatedName := fmt.Sprintf("TestViewUpdated-%v", uuid.NewString())
	updated := views.View1{
		Name:          updatedName,
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
	assert.Equal(t, updatedName, got.Name)

	_, _, err = client.
		ViewsServiceDeleteView(context.Background(), created.Id).
		Execute()
	assertNilAndPrintError(t, err)
}

func TestViewsFolders(t *testing.T) {
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewViewsFoldersClient(cpc)

	before, _, err := client.
		ViewsFoldersServiceListViewFolders(context.Background()).
		Execute()
	assertNilAndPrintError(t, err)
	initialCount := len(before.Folders)

	createReq := viewsfolders.CreateViewFolderRequest{
		Name: viewsfolders.PtrString(fmt.Sprintf("GoTestViewFolder-%s", uuid.NewString())),
	}
	created, _, err := client.
		ViewsFoldersServiceCreateViewFolder(context.Background()).
		CreateViewFolderRequest(createReq).
		Execute()
	assertNilAndPrintError(t, err)
	assert.NotNil(t, created)
	assert.NotNil(t, created.Id)

	updatedName := fmt.Sprintf("GoTestViewFolderUpdated-%s", uuid.NewString())
	updateReq := viewsfolders.ViewFolder1{
		Name: updatedName,
	}
	_, _, err = client.
		ViewsFoldersServiceReplaceViewFolder(context.Background(), *created.Id).
		ViewFolder1(updateReq).
		Execute()
	assertNilAndPrintError(t, err)

	got, _, err := client.
		ViewsFoldersServiceGetViewFolder(context.Background(), *created.Id).
		Execute()
	assertNilAndPrintError(t, err)
	assert.Equal(t, updatedName, got.Name)

	after, _, err := client.
		ViewsFoldersServiceListViewFolders(context.Background()).
		Execute()
	assertNilAndPrintError(t, err)
	assert.Equal(t, initialCount+1, len(after.Folders))

	_, _, err = client.
		ViewsFoldersServiceDeleteViewFolder(context.Background(), *created.Id).
		Execute()
	assertNilAndPrintError(t, err)

	final, _, err := client.
		ViewsFoldersServiceListViewFolders(context.Background()).
		Execute()
	assertNilAndPrintError(t, err)
	assert.Equal(t, initialCount, len(final.Folders))
}
