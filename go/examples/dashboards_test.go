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

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/wrapperspb"

	cxsdk "github.com/coralogix/coralogix-management-sdk/go"
)

func TestDashboards(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewSDKCallPropertiesCreator(region, authContext)
	c := cxsdk.NewDashboardsClient(creator)
	dat, err := os.ReadFile("dashboard.json")
	assertNilAndPrintError(t, err)
	var d cxsdk.Dashboard
	e := protojson.Unmarshal(dat, &d)
	if e != nil {
		log.Fatal(e.Error())
	}

	createRes, e := c.Create(context.Background(), &cxsdk.CreateDashboardRequest{
		Dashboard: &d,
	})
	if e != nil {
		log.Fatal(e.Error())
	}
	assertNilAndPrintError(t, e)
	_, e = c.Pin(context.Background(), &cxsdk.PinDashboardRequest{
		DashboardId: createRes.DashboardId,
	})
	assertNilAndPrintError(t, e)
	_, e = c.Unpin(context.Background(), &cxsdk.UnpinDashboardRequest{
		DashboardId: createRes.DashboardId,
	})
	assertNilAndPrintError(t, e)

	all, e := c.List(context.Background())
	assertNilAndPrintError(t, e)
	found := false
	for _, v := range all.Items {
		if v.Id.Value == createRes.DashboardId.Value {
			found = true
			break
		}
	}
	assert.True(t, found)
	_, e = c.Delete(context.Background(), &cxsdk.DeleteDashboardRequest{
		DashboardId: createRes.DashboardId,
	})
	assertNilAndPrintError(t, e)
}

func TestDashboardFolders(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewSDKCallPropertiesCreator(region, authContext)
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
