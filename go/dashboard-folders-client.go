package cxsdk

import (
	"context"
	dashboards "coralogix-management-sdk/go/internal/coralogixapis/dashboards/v1"
)

// DashboardsFoldersClient is a client for the Coralogix Dashboards Folders API.
type DashboardsFoldersClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// CreateDashboardsFolder creates a new dashboard folder.
func (c DashboardsFoldersClient) CreateDashboardsFolder(ctx context.Context, req *dashboards.CreateDashboardFolderRequest) (*dashboards.CreateDashboardFolderResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := dashboards.NewDashboardFoldersServiceClient(conn)

	return client.CreateDashboardFolder(callProperties.Ctx, req, callProperties.CallOptions...)
}

// GetDashboardsFolders gets all dashboard folders.
func (c DashboardsFoldersClient) GetDashboardsFolders(ctx context.Context, req *dashboards.ListDashboardFoldersRequest) (*dashboards.ListDashboardFoldersResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := dashboards.NewDashboardFoldersServiceClient(conn)

	return client.ListDashboardFolders(callProperties.Ctx, req, callProperties.CallOptions...)
}

// UpdateDashboardsFolder updates a dashboard folder.
func (c DashboardsFoldersClient) UpdateDashboardsFolder(ctx context.Context, req *dashboards.ReplaceDashboardFolderRequest) (*dashboards.ReplaceDashboardFolderResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := dashboards.NewDashboardFoldersServiceClient(conn)

	return client.ReplaceDashboardFolder(callProperties.Ctx, req, callProperties.CallOptions...)
}

// DeleteDashboardsFolder deletes a dashboard folder.
func (c DashboardsFoldersClient) DeleteDashboardsFolder(ctx context.Context, req *dashboards.DeleteDashboardFolderRequest) (*dashboards.DeleteDashboardFolderResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
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
