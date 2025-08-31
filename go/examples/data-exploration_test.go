// Copyright 2025 Coralogix Ltd.
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

	cxsdk "github.com/coralogix/coralogix-management-sdk/go"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestViews(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewSDKCallPropertiesCreator(region, authContext)
	c := cxsdk.NewViewsClient(creator)

	view, e := c.Create(context.Background(), &cxsdk.CreateViewRequest{
		Name: &wrapperspb.StringValue{Value: "GoTestView"},
		SearchQuery: &cxsdk.SearchQuery{
			Query: &wrapperspb.StringValue{Value: "source logs | filter $l.applicationname == 'default'"},
		},
		TimeSelection: &cxsdk.TimeSelection{
			SelectionType: &cxsdk.ViewTimeSelectionQuick{
				QuickSelection: &cxsdk.QuickTimeSelection{
					Seconds: uint32(86400), // 24h
				},
			},
		},
		Filters: &cxsdk.SelectedFilters{},
	})
	assertNilAndPrintError(t, e)

	view.View.Name = &wrapperspb.StringValue{Value: "GoTestViewUpdated"}

	_, e = c.Replace(context.Background(), &cxsdk.ReplaceViewRequest{
		View: view.View,
	})

	assertNilAndPrintError(t, e)

	_, e = c.Get(context.Background(), &cxsdk.GetViewRequest{
		Id: view.View.Id,
	})

	assertNilAndPrintError(t, e)

	c.Delete(context.Background(), &cxsdk.DeleteViewRequest{
		Id: view.View.Id,
	})
}

func TestViewFolders(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewSDKCallPropertiesCreator(region, authContext)
	c := cxsdk.NewViewFoldersClient(creator)

	allFolders, e := c.List(context.Background(), &cxsdk.ListViewFoldersRequest{})
	assertNilAndPrintError(t, e)

	numOfFolders := len(allFolders.Folders)

	createResponse, e := c.Create(context.Background(), &cxsdk.CreateViewFolderRequest{
		Name: &wrapperspb.StringValue{Value: "GoTestViewFolder"},
	})
	assertNilAndPrintError(t, e)

	createResponse.Folder.Name = &wrapperspb.StringValue{Value: "GoTestViewFolderUpdated"}

	_, e = c.Replace(context.Background(), &cxsdk.ReplaceViewFolderRequest{
		Folder: createResponse.Folder,
	})
	assertNilAndPrintError(t, e)

	updatedFolder, e := c.Get(context.Background(), &cxsdk.GetViewFolderRequest{
		Id: createResponse.Folder.Id,
	})

	assertNilAndPrintError(t, e)
	assert.Equal(t, updatedFolder.Folder.Name.Value, "GoTestViewFolderUpdated")

	allFoldersWithNewFolder, e := c.List(context.Background(), &cxsdk.ListViewFoldersRequest{})

	assertNilAndPrintError(t, e)
	assert.Equal(t, numOfFolders+1, len(allFoldersWithNewFolder.Folders))

	c.Delete(context.Background(), &cxsdk.DeleteViewFolderRequest{
		Id: createResponse.Folder.Id,
	})

	assertNilAndPrintError(t, e)

	allFolders, e = c.List(context.Background(), &cxsdk.ListViewFoldersRequest{})

	assertNilAndPrintError(t, e)
	assert.Equal(t, numOfFolders, len(allFolders.Folders))

}
