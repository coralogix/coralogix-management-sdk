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

package cxsdk

import (
	"context"

	dashboards "github.com/coralogix/coralogix-management-sdk/internal/coralogixapis/dashboards/v1"
)

// DashboardsClient is a client for the Coralogix Dashboards API.
type DashboardsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// CreateDashboardRequest is a request to create a dashboard.
type CreateDashboardRequest = dashboards.CreateDashboardRequest

// ReplaceDashboardRequest is a request to replace a dashboard.
type ReplaceDashboardRequest = dashboards.ReplaceDashboardRequest

// DeleteDashboardRequest is a request to delete a dashboard.
type DeleteDashboardRequest = dashboards.DeleteDashboardRequest

// GetDashboardRequest is a request to get a dashboard.
type GetDashboardRequest = dashboards.GetDashboardRequest

// PinDashboardRequest is a request to pin a dashboard.
type PinDashboardRequest = dashboards.PinDashboardRequest

// UnpinDashboardRequest is a request to unpin a dashboard.
type UnpinDashboardRequest = dashboards.UnpinDashboardRequest

// DashboardFolderPath is a dashboard folder path.
type DashboardFolderPath = dashboards.Dashboard_FolderPath

// FolderPath is a dashboard folder path.
type FolderPath = dashboards.FolderPath

// Dashboard is a dashboard.
type Dashboard = dashboards.Dashboard

// Create Creates a new dashboard.
func (d DashboardsClient) Create(ctx context.Context, req *CreateDashboardRequest) (*dashboards.CreateDashboardResponse, error) {
	callProperties, err := d.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := dashboards.NewDashboardsServiceClient(conn)

	return client.CreateDashboard(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Get gets a dashboard.
func (d DashboardsClient) Get(ctx context.Context, req *GetDashboardRequest) (*dashboards.GetDashboardResponse, error) {
	callProperties, err := d.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := dashboards.NewDashboardsServiceClient(conn)

	return client.GetDashboard(callProperties.Ctx, req, callProperties.CallOptions...)
}

// List lists all dashboards.
func (d DashboardsClient) List(ctx context.Context) (*dashboards.GetDashboardCatalogResponse, error) {
	callProperties, err := d.callPropertiesCreator.GetUserLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := dashboards.NewDashboardCatalogServiceClient(conn)

	return client.GetDashboardCatalog(callProperties.Ctx, &dashboards.GetDashboardCatalogRequest{}, callProperties.CallOptions...)
}

// Replace replaces a dashboard.
func (d DashboardsClient) Replace(ctx context.Context, req *ReplaceDashboardRequest) (*dashboards.ReplaceDashboardResponse, error) {
	callProperties, err := d.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := dashboards.NewDashboardsServiceClient(conn)

	return client.ReplaceDashboard(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Delete deletes a dashboard.
func (d DashboardsClient) Delete(ctx context.Context, req *DeleteDashboardRequest) (*dashboards.DeleteDashboardResponse, error) {
	callProperties, err := d.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := dashboards.NewDashboardsServiceClient(conn)

	return client.DeleteDashboard(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Pin pins a dashboard.
func (d DashboardsClient) Pin(ctx context.Context, req *PinDashboardRequest) (*dashboards.PinDashboardResponse, error) {
	callProperties, err := d.callPropertiesCreator.GetUserLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := dashboards.NewDashboardsServiceClient(conn)

	return client.PinDashboard(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Unpin unpins a dashboard.
func (d DashboardsClient) Unpin(ctx context.Context, req *UnpinDashboardRequest) (*dashboards.UnpinDashboardResponse, error) {
	callProperties, err := d.callPropertiesCreator.GetUserLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := dashboards.NewDashboardsServiceClient(conn)

	return client.UnpinDashboard(callProperties.Ctx, req, callProperties.CallOptions...)
}

// AssignToFolder assigns a dashboard to a folder.
func (d DashboardsClient) AssignToFolder(ctx context.Context, req *dashboards.AssignDashboardFolderRequest) (*dashboards.AssignDashboardFolderResponse, error) {
	callProperties, err := d.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := dashboards.NewDashboardsServiceClient(conn)

	return client.AssignDashboardFolder(callProperties.Ctx, req, callProperties.CallOptions...)
}

// NewDashboardsClient creates a new DashboardsClient.
func NewDashboardsClient(c *CallPropertiesCreator) *DashboardsClient {
	return &DashboardsClient{callPropertiesCreator: c}
}
