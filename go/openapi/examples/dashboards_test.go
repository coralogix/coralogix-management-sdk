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
	"net/http"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	dashboardfolders "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/dashboard_folders_service"
	dashboards "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/dashboard_service"
)

func TestDashboards(t *testing.T) {
	cfg := newTestConfig()
	client := cxsdk.NewDashboardClient(cfg)

	name := "sdk-openapi-" + uuid.NewString()
	created := dashboardFromFixture(t, name)
	updated := dashboardFromFixture(t, name+" updated")

	exerciseDashboardLifecycle(t, client, created, updated)
}

func TestDashboardFolders(t *testing.T) {
	cfg := newTestConfig()
	client := cxsdk.NewDashboardFoldersClient(cfg)

	id := uuid.New().String()
	folderName := "sdk-openapi-folder-" + uuid.NewString()
	folder := dashboardfolders.DashboardFolder{
		Id:   dashboardfolders.PtrString(id),
		Name: dashboardfolders.PtrString(folderName),
	}
	createDashboardFolderRequest := dashboardfolders.CreateDashboardFolderRequestDataStructure{
		Folder: &folder,
	}
	_, httpResp, err := client.
		DashboardFoldersServiceCreateDashboardFolder(context.Background()).CreateDashboardFolderRequestDataStructure(createDashboardFolderRequest).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	t.Cleanup(func() {
		deleteDashboardFolder(t, client, id)
	})

	updatedFolderName := "updated " + folderName
	replaceReq := dashboardfolders.ReplaceDashboardFolderRequestDataStructure{
		Folder: &dashboardfolders.DashboardFolder{
			Id:   dashboardfolders.PtrString(id),
			Name: dashboardfolders.PtrString(updatedFolderName),
		},
	}
	_, httpResp, err = client.
		DashboardFoldersServiceReplaceDashboardFolder(context.Background()).
		ReplaceDashboardFolderRequestDataStructure(replaceReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	getResp, httpResp, err := client.
		DashboardFoldersServiceGetDashboardFolder(context.Background(), id).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, updatedFolderName, getResp.Folder.GetName())

	_, httpResp, err = client.
		DashboardFoldersServiceListDashboardFolders(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	dashboardClient := cxsdk.NewDashboardClient(cfg)
	t.Run("dashboard with folderId", func(t *testing.T) {
		name := "sdk-openapi-folder-id-" + uuid.NewString()
		created := dashboardFromFixture(t, name, withFolderID(id))
		updated := dashboardFromFixture(t, name+" updated", withFolderID(id))

		exerciseDashboardLifecycle(t, dashboardClient, created, updated)
	})

	t.Run("dashboard with folderPath", func(t *testing.T) {
		name := "sdk-openapi-folder-path-" + uuid.NewString()
		created := dashboardFromFixture(t, name, withFolderPath(updatedFolderName))
		updated := dashboardFromFixture(t, name+" updated", withFolderPath(updatedFolderName))

		exerciseDashboardLifecycle(t, dashboardClient, created, updated)
	})
}

func exerciseDashboardLifecycle(t *testing.T, client *dashboards.DashboardServiceAPIService, created, updated dashboards.Dashboard) string {
	t.Helper()

	req := dashboards.CreateDashboardRequestDataStructure{
		Dashboard: created,
		RequestId: uuid.NewString(),
	}
	createResp, httpResp, err := client.
		DashboardsServiceCreateDashboard(context.Background()).
		CreateDashboardRequestDataStructure(req).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotNil(t, createResp.DashboardId)

	dashboardID := *createResp.DashboardId
	t.Cleanup(func() {
		deleteDashboard(t, client, dashboardID)
	})

	getResp, httpResp, err := client.
		DashboardsServiceGetDashboard(context.Background(), dashboardID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotNil(t, getResp.Dashboard)

	replaceReq := dashboards.ReplaceDashboardRequestDataStructure{
		Dashboard: withDashboardID(t, updated, dashboardID),
		RequestId: uuid.NewString(),
	}
	_, httpResp, err = client.
		DashboardsServiceReplaceDashboard(context.Background()).
		ReplaceDashboardRequestDataStructure(replaceReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	getResp, httpResp, err = client.
		DashboardsServiceGetDashboard(context.Background(), dashboardID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotNil(t, getResp.Dashboard)

	_, httpResp, err = client.
		DashboardsServicePinDashboard(context.Background(), dashboardID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	_, httpResp, err = client.
		DashboardsServiceUnpinDashboard(context.Background(), dashboardID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	catalog, httpResp, err := client.
		DashboardCatalogServiceGetDashboardCatalog(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	found := false
	for _, v := range catalog.Items {
		if v.Id != nil && *v.Id == dashboardID {
			found = true
			break
		}
	}
	require.True(t, found)

	return dashboardID
}

type dashboardFixtureOption func(map[string]interface{})

func dashboardFromFixture(t *testing.T, name string, opts ...dashboardFixtureOption) dashboards.Dashboard {
	t.Helper()

	raw := dashboardFixtureMap(t)
	raw["name"] = name
	raw["description"] = "dashboards OpenAPI SDK validation"
	for _, opt := range opts {
		opt(raw)
	}

	data, err := json.Marshal(raw)
	require.NoError(t, err)

	var dashboard dashboards.Dashboard
	err = json.Unmarshal(data, &dashboard)
	require.NoError(t, err)
	require.Equal(t, name, dashboard.GetName())
	return dashboard
}

func dashboardFixtureMap(t *testing.T) map[string]interface{} {
	t.Helper()

	data, err := os.ReadFile("dashboard.json")
	require.NoError(t, err)

	var raw map[string]interface{}
	require.NoError(t, json.Unmarshal(data, &raw))
	return raw
}

func withDashboardID(t *testing.T, dashboard dashboards.Dashboard, id string) dashboards.Dashboard {
	t.Helper()

	data, err := json.Marshal(dashboard)
	require.NoError(t, err)

	var raw map[string]interface{}
	require.NoError(t, json.Unmarshal(data, &raw))
	raw["id"] = id

	data, err = json.Marshal(raw)
	require.NoError(t, err)

	err = json.Unmarshal(data, &dashboard)
	require.NoError(t, err)
	require.Equal(t, id, dashboard.GetId())
	return dashboard
}

func withFolderID(id string) dashboardFixtureOption {
	return func(raw map[string]interface{}) {
		raw["folderId"] = map[string]string{"value": id}
		delete(raw, "folderPath")
	}
}

func withFolderPath(name string) dashboardFixtureOption {
	return func(raw map[string]interface{}) {
		raw["folderPath"] = map[string][]string{"segments": []string{name}}
		delete(raw, "folderId")
	}
}

func deleteDashboard(t *testing.T, client *dashboards.DashboardServiceAPIService, dashboardID string) {
	t.Helper()

	_, httpResp, err := client.
		DashboardsServiceDeleteDashboard(context.Background(), dashboardID).
		Execute()
	if err != nil && (httpResp == nil || httpResp.StatusCode != http.StatusNotFound) {
		t.Errorf("failed to delete dashboard %s: %v", dashboardID, cxsdk.NewAPIError(httpResp, err))
	}
}

func deleteDashboardFolder(t *testing.T, client *dashboardfolders.DashboardFoldersServiceAPIService, folderID string) {
	t.Helper()

	_, httpResp, err := client.
		DashboardFoldersServiceDeleteDashboardFolder(context.Background(), folderID).
		Execute()
	if err != nil && (httpResp == nil || httpResp.StatusCode != http.StatusNotFound) {
		t.Errorf("failed to delete dashboard folder %s: %v", folderID, cxsdk.NewAPIError(httpResp, err))
	}
}
