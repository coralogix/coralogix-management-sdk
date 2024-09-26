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
	"log"
	"os"
	"testing"

	cxsdk "github.com/coralogix/coralogix-management-sdk/go"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestDashboards(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	c := cxsdk.NewDashboardsClient(creator)
	dat, err := os.ReadFile("dashboard.json")
	assert.Nil(t, err)
	var d cxsdk.Dashboard
	assert.Nil(t, protojson.Unmarshal(dat, &d))

	_, e := c.Get(context.Background(), &cxsdk.GetDashboardRequest{
		DashboardId: d.Id,
	})
	if e != nil {
		c.Delete(context.Background(), &cxsdk.DeleteDashboardRequest{
			DashboardId: d.Id,
		})
	}
	_, e = c.Create(context.Background(), &cxsdk.CreateDashboardRequest{
		Dashboard: &d,
	})
	if e != nil {
		log.Fatal(e.Error())
	}
	assert.Nil(t, e)
	_, e = c.Pin(context.Background(), &cxsdk.PinDashboardRequest{
		DashboardId: d.Id,
	})
	assert.Nil(t, e)
	_, e = c.Unpin(context.Background(), &cxsdk.UnpinDashboardRequest{
		DashboardId: d.Id,
	})
	assert.Nil(t, e)

	all, e := c.List(context.Background())
	assert.Nil(t, e)
	found := false
	for _, v := range all.Items {
		if v.Id.Value == d.Id.Value {
			found = true
			break
		}
	}
	assert.True(t, found)
	_, e = c.Delete(context.Background(), &cxsdk.DeleteDashboardRequest{
		DashboardId: d.Id,
	})
	assert.Nil(t, e)
}

func TestDashboardFolders(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	c := cxsdk.NewDashboardsFoldersClient(creator)

	id := uuid.New().String()

	dashboardFolder := cxsdk.DashboardFolder{
		Id:   wrapperspb.String(id),
		Name: wrapperspb.String(id),
	}

	_, creationError := c.Create(context.Background(), &cxsdk.CreateDashboardFolderRequest{
		Folder: &dashboardFolder,
	})

	assert.Nil(t, creationError)

	_, updateError := c.Replace(context.Background(), &cxsdk.ReplaceDashboardFolderRequest{
		Folder: &cxsdk.DashboardFolder{
			Id:   wrapperspb.String(id),
			Name: wrapperspb.String("updated " + id),
		},
	})

	assert.Nil(t, updateError)

	getDashboardFolderResponse, retrievalError := c.Get(context.Background(), &cxsdk.GetDashboardFolderRequest{
		FolderId: wrapperspb.String(id),
	})

	assert.Nil(t, retrievalError)
	assert.Equal(t, "updated "+id, getDashboardFolderResponse.Folder.Name.Value)

	_, listError := c.List(context.Background())

	assert.Nil(t, listError)

	_, deletionError := c.Delete(context.Background(), &cxsdk.DeleteDashboardFolderRequest{
		FolderId: wrapperspb.String(id),
	})

	assert.Nil(t, deletionError)

}
