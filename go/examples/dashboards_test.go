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
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	c := cxsdk.NewDashboardsClient(creator)
	dat, err := os.ReadFile("dashboard.json")
	assertNilAndPrintError(t, err)
	var d cxsdk.Dashboard
	assertNilAndPrintError(t, protojson.Unmarshal(dat, &d))
	d.Id = wrapperspb.String(uuid.New().String()[:21])
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
	assertNilAndPrintError(t, e)
	_, e = c.Pin(context.Background(), &cxsdk.PinDashboardRequest{
		DashboardId: d.Id,
	})
	assertNilAndPrintError(t, e)
	_, e = c.Unpin(context.Background(), &cxsdk.UnpinDashboardRequest{
		DashboardId: d.Id,
	})
	assertNilAndPrintError(t, e)

	all, e := c.List(context.Background())
	assertNilAndPrintError(t, e)
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
	assertNilAndPrintError(t, e)
}

func TestDashboardFolders(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
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

	assertNilAndPrintError(t, creationError)

	_, updateError := c.Replace(context.Background(), &cxsdk.ReplaceDashboardFolderRequest{
		Folder: &cxsdk.DashboardFolder{
			Id:   wrapperspb.String(id),
			Name: wrapperspb.String("updated " + id),
		},
	})

	assertNilAndPrintError(t, updateError)

	getDashboardFolderResponse, retrievalError := c.Get(context.Background(), &cxsdk.GetDashboardFolderRequest{
		FolderId: wrapperspb.String(id),
	})

	assertNilAndPrintError(t, retrievalError)
	assert.Equal(t, "updated "+id, getDashboardFolderResponse.Folder.Name.Value)

	_, listError := c.List(context.Background())

	assertNilAndPrintError(t, listError)

	_, deletionError := c.Delete(context.Background(), &cxsdk.DeleteDashboardFolderRequest{
		FolderId: wrapperspb.String(id),
	})

	assertNilAndPrintError(t, deletionError)

}
