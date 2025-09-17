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
	"encoding/json"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	dashboardfolders "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/dashboard_folders_service"
	dashboards "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/dashboard_service"
)

func TestDashboards(t *testing.T) {
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewDashboardClient(cpc)

	data, err := os.ReadFile("dashboard.json")
	if err != nil {
		t.Fatalf("failed to read dashboard.json: %v", err)
	}

	var dashboard dashboards.Dashboard
	err = json.Unmarshal(data, &dashboard)
	assertNilAndPrintError(t, err)

	req := dashboards.CreateDashboardRequestDataStructure{
		Dashboard: dashboards.GetDashboardByURLSlugResponseDataStructureDashboard{
			Dashboard: &dashboard,
		},
	}
	created, _, err := client.
		DashboardsServiceCreateDashboard(context.Background()).
		CreateDashboardRequestDataStructure(req).
		Execute()
	assertNilAndPrintError(t, err)
	assert.NotNil(t, created.DashboardId)

	dashboardID := *created.DashboardId

	_, _, err = client.
		DashboardsServicePinDashboard(context.Background(), dashboardID).
		PinDashboardRequestDataStructure(dashboards.PinDashboardRequestDataStructure{}).
		Execute()
	assertNilAndPrintError(t, err)

	_, _, err = client.
		DashboardsServiceUnpinDashboard(context.Background(), dashboardID).
		UnpinDashboardRequestDataStructure(dashboards.UnpinDashboardRequestDataStructure{}).
		Execute()
	assertNilAndPrintError(t, err)

	catalog, _, err := client.
		DashboardCatalogServiceGetDashboardCatalog(context.Background()).
		Execute()
	assertNilAndPrintError(t, err)

	found := false
	for _, v := range catalog.Items {
		if v.Id != nil && *v.Id == dashboardID {
			found = true
			break
		}
	}
	assert.True(t, found)

	_, _, err = client.
		DashboardsServiceDeleteDashboard(context.Background(), dashboardID).
		Execute()
	assertNilAndPrintError(t, err)
}

func TestDashboardFolders(t *testing.T) {
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewDashboardFoldersClient(cpc)

	id := uuid.New().String()

	createReq := dashboardfolders.CreateDashboardFolderRequestDataStructure{
		Folder: &dashboardfolders.ListDashboardFoldersResponseDataStructureFolderInner{
			Id:   dashboardfolders.PtrString(id),
			Name: dashboardfolders.PtrString(id),
		},
	}
	_, _, err := client.
		DashboardFoldersServiceCreateDashboardFolder(context.Background()).
		CreateDashboardFolderRequestDataStructure(createReq).
		Execute()
	assertNilAndPrintError(t, err)

	replaceReq := dashboardfolders.ReplaceDashboardFolderRequestDataStructure{
		Folder: &dashboardfolders.ListDashboardFoldersResponseDataStructureFolderInner{
			Id:   dashboardfolders.PtrString(id),
			Name: dashboardfolders.PtrString("updated " + id),
		},
	}
	_, _, err = client.
		DashboardFoldersServiceReplaceDashboardFolder(context.Background()).
		ReplaceDashboardFolderRequestDataStructure(replaceReq).
		Execute()
	assertNilAndPrintError(t, err)

	getResp, _, err := client.
		DashboardFoldersServiceGetDashboardFolder(context.Background(), id).
		Execute()
	assertNilAndPrintError(t, err)
	assert.Equal(t, "updated "+id, getResp.Folder.GetName())

	_, _, err = client.
		DashboardFoldersServiceListDashboardFolders(context.Background()).
		Execute()
	assertNilAndPrintError(t, err)

	_, _, err = client.
		DashboardFoldersServiceDeleteDashboardFolder(context.Background(), id).
		Execute()
	assertNilAndPrintError(t, err)
}
