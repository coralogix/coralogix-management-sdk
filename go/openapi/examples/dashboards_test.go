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
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"

	cxsdk "github.com/coralogix/coralogix-management-sdk/go/openapi"
)

func TestDashboards(t *testing.T) {
	//t.Skip("Skipping dashboard test")
	region := os.Getenv(cxsdk.EnvCoralogixRegion)
	if region == "" {
		t.Fatalf("env %s not set", cxsdk.EnvCoralogixRegion)
	}

	apiKey := os.Getenv(cxsdk.EnvCoralogixAPIKey)
	if apiKey == "" {
		t.Fatalf("env %s not set", cxsdk.EnvCoralogixAPIKey)
	}

	client := cxsdk.NewClientBuilder().
		WithRegion(region).
		WithAPIKey(apiKey).
		Build()

	data, err := os.ReadFile("dashboard.json")
	if err != nil {
		t.Fatalf("failed to read dashboard.json: %v", err)
	}

	var dashboard cxsdk.Dashboard
	err = json.Unmarshal(data, &dashboard)
	assert.NoError(t, err)

	req := cxsdk.CreateDashboardRequestDataStructure{
		Dashboard: cxsdk.DashboardAsComCoralogixapisDashboardsV1AstDashboard(&dashboard),
	}
	created, _, err := client.DashboardServiceAPI.
		DashboardsServiceCreateDashboard(context.Background()).
		CreateDashboardRequestDataStructure(req).
		Execute()
	assert.NoError(t, err)
	assert.NotNil(t, created.DashboardId)

	dashboardID := *created.DashboardId

	_, _, err = client.DashboardServiceAPI.
		DashboardsServicePinDashboard(context.Background(), dashboardID).
		PinDashboardRequestDataStructure(cxsdk.PinDashboardRequestDataStructure{}).
		Execute()
	assert.NoError(t, err)

	_, _, err = client.DashboardServiceAPI.
		DashboardsServiceUnpinDashboard(context.Background(), dashboardID).
		UnpinDashboardRequestDataStructure(cxsdk.UnpinDashboardRequestDataStructure{}).
		Execute()
	assert.NoError(t, err)

	catalog, _, err := client.DashboardServiceAPI.
		DashboardCatalogServiceGetDashboardCatalog(context.Background()).
		Execute()
	assert.NoError(t, err)

	found := false
	for _, v := range catalog.Items {
		if v.Id != nil && *v.Id == dashboardID {
			found = true
			break
		}
	}
	assert.True(t, found)

	_, _, err = client.DashboardServiceAPI.
		DashboardsServiceDeleteDashboard(context.Background(), dashboardID).
		Execute()
	assert.NoError(t, err)
}

func TestDashboardFolders(t *testing.T) {
	region := os.Getenv(cxsdk.EnvCoralogixRegion)
	if region == "" {
		t.Fatalf("env %s not set", cxsdk.EnvCoralogixRegion)
	}

	apiKey := os.Getenv(cxsdk.EnvCoralogixAPIKey)
	if apiKey == "" {
		t.Fatalf("env %s not set", cxsdk.EnvCoralogixAPIKey)
	}

	client := cxsdk.NewClientBuilder().
		WithRegion(region).
		WithAPIKey(apiKey).
		Build()

	id := uuid.New().String()

	createReq := cxsdk.CreateDashboardFolderRequestDataStructure{
		Folder: &cxsdk.ComCoralogixapisDashboardsV1CommonDashboardFolder{
			Id:   cxsdk.PtrString(id),
			Name: cxsdk.PtrString(id),
		},
	}
	_, _, err := client.DashboardFoldersServiceAPI.
		DashboardFoldersServiceCreateDashboardFolder(context.Background()).
		CreateDashboardFolderRequestDataStructure(createReq).
		Execute()
	assert.NoError(t, err)

	replaceReq := cxsdk.ReplaceDashboardFolderRequestDataStructure{
		Folder: &cxsdk.ComCoralogixapisDashboardsV1CommonDashboardFolder{
			Id:   cxsdk.PtrString(id),
			Name: cxsdk.PtrString("updated " + id),
		},
	}
	_, _, err = client.DashboardFoldersServiceAPI.
		DashboardFoldersServiceReplaceDashboardFolder(context.Background()).
		ReplaceDashboardFolderRequestDataStructure(replaceReq).
		Execute()
	assert.NoError(t, err)

	getResp, _, err := client.DashboardFoldersServiceAPI.
		DashboardFoldersServiceGetDashboardFolder(context.Background(), id).
		Execute()
	assert.NoError(t, err)
	assert.Equal(t, "updated "+id, getResp.Folder.GetName())

	_, _, err = client.DashboardFoldersServiceAPI.
		DashboardFoldersServiceListDashboardFolders(context.Background()).
		Execute()
	assert.NoError(t, err)

	_, _, err = client.DashboardFoldersServiceAPI.
		DashboardFoldersServiceDeleteDashboardFolder(context.Background(), id).
		Execute()
	assert.NoError(t, err)
}
