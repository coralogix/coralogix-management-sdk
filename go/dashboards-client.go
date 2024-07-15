package cxsdk

import (
	"context"
	dashboards "coralogix-management-sdk/go/internal/coralogixapis/dashboards/v1"
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

// Dashboard is a dashboard.
type Dashboard = dashboards.Dashboard

// CreateDashboard Creates a new dashboard.
func (d DashboardsClient) CreateDashboard(ctx context.Context, req *CreateDashboardRequest) (*dashboards.CreateDashboardResponse, error) {
	callProperties, err := d.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := dashboards.NewDashboardsServiceClient(conn)

	return client.CreateDashboard(callProperties.Ctx, req, callProperties.CallOptions...)
}

// GetDashboard gets a dashboard.
func (d DashboardsClient) GetDashboard(ctx context.Context, req *GetDashboardRequest) (*dashboards.GetDashboardResponse, error) {
	callProperties, err := d.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := dashboards.NewDashboardsServiceClient(conn)

	return client.GetDashboard(callProperties.Ctx, req, callProperties.CallOptions...)
}

// UpdateDashboard updates a dashboard.
func (d DashboardsClient) UpdateDashboard(ctx context.Context, req *ReplaceDashboardRequest) (*dashboards.ReplaceDashboardResponse, error) {
	callProperties, err := d.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := dashboards.NewDashboardsServiceClient(conn)

	return client.ReplaceDashboard(callProperties.Ctx, req, callProperties.CallOptions...)
}

// DeleteDashboard deletes a dashboard.
func (d DashboardsClient) DeleteDashboard(ctx context.Context, req *DeleteDashboardRequest) (*dashboards.DeleteDashboardResponse, error) {
	callProperties, err := d.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := dashboards.NewDashboardsServiceClient(conn)

	return client.DeleteDashboard(callProperties.Ctx, req, callProperties.CallOptions...)
}

// PinDashboard pins a dashboard.
func (d DashboardsClient) PinDashboard(ctx context.Context, req *PinDashboardRequest) (*dashboards.PinDashboardResponse, error) {
	callProperties, err := d.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := dashboards.NewDashboardsServiceClient(conn)

	return client.PinDashboard(callProperties.Ctx, req, callProperties.CallOptions...)
}

// UnpinDashboard unpins a dashboard.
func (d DashboardsClient) UnpinDashboard(ctx context.Context, req *UnpinDashboardRequest) (*dashboards.UnpinDashboardResponse, error) {
	callProperties, err := d.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := dashboards.NewDashboardsServiceClient(conn)

	return client.UnpinDashboard(callProperties.Ctx, req, callProperties.CallOptions...)
}

// NewDashboardsClient creates a new DashboardsClient.
func NewDashboardsClient(c *CallPropertiesCreator) *DashboardsClient {
	return &DashboardsClient{callPropertiesCreator: c}
}
