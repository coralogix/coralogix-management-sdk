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

	dashboards "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/dashboards/v1"
)

// DashboardsFoldersClient is a client for the Coralogix Dashboards Folders API.
type DashboardsFoldersClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// ListDashboardFolderRequest is a request to get a dashboard folders.
type ListDashboardFolderRequest = dashboards.ListDashboardFoldersRequest

// CreateDashboardFolderRequest is a request to create a dashboards folder.
type CreateDashboardFolderRequest = dashboards.CreateDashboardFolderRequest

// DashboardFolder is a dashboard folder.
type DashboardFolder = dashboards.DashboardFolder

// GetDashboardFolderRequest is a request to get a dashboard folders.
type GetDashboardFolderRequest = dashboards.GetDashboardFolderRequest

// ReplaceDashboardFolderRequest is a request to replace a dashboard folder.
type ReplaceDashboardFolderRequest = dashboards.ReplaceDashboardFolderRequest

// DeleteDashboardFolderRequest is a request to delete a dashboard folder.
type DeleteDashboardFolderRequest = dashboards.DeleteDashboardFolderRequest

// Create creates a new dashboard folder.
func (c DashboardsFoldersClient) Create(ctx context.Context, req *dashboards.CreateDashboardFolderRequest) (*dashboards.CreateDashboardFolderResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := dashboards.NewDashboardFoldersServiceClient(conn)

	return client.CreateDashboardFolder(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Get dashboard folder details.
func (c DashboardsFoldersClient) Get(ctx context.Context, req *dashboards.GetDashboardFolderRequest) (*dashboards.GetDashboardFolderResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := dashboards.NewDashboardFoldersServiceClient(conn)

	return client.GetDashboardFolder(callProperties.Ctx, req, callProperties.CallOptions...)
}

// List gets all dashboard folders.
func (c DashboardsFoldersClient) List(ctx context.Context) (*dashboards.ListDashboardFoldersResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetUserLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := dashboards.NewDashboardFoldersServiceClient(conn)

	return client.ListDashboardFolders(callProperties.Ctx, &dashboards.ListDashboardFoldersRequest{}, callProperties.CallOptions...)
}

// Replace updates a dashboard folder.
func (c DashboardsFoldersClient) Replace(ctx context.Context, req *dashboards.ReplaceDashboardFolderRequest) (*dashboards.ReplaceDashboardFolderResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := dashboards.NewDashboardFoldersServiceClient(conn)

	return client.ReplaceDashboardFolder(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Delete deletes a dashboard folder.
func (c DashboardsFoldersClient) Delete(ctx context.Context, req *dashboards.DeleteDashboardFolderRequest) (*dashboards.DeleteDashboardFolderResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := dashboards.NewDashboardFoldersServiceClient(conn)

	return client.DeleteDashboardFolder(callProperties.Ctx, req, callProperties.CallOptions...)
}

// NewDashboardsFoldersClient Creates a new DashboardsFoldersClient.
func NewDashboardsFoldersClient(c *CallPropertiesCreator) *DashboardsFoldersClient {
	return &DashboardsFoldersClient{callPropertiesCreator: c}
}
