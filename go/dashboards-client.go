package cxsdk

import (
	"context"
	dashboards "coralogix-management-sdk/go/internal/coralogixapis/dashboards/v1"
)

type DashboardsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

type CreateDashboardRequest = dashboards.CreateDashboardRequest
type ReplaceDashboardRequest = dashboards.ReplaceDashboardRequest
type DeleteDashboardRequest = dashboards.DeleteDashboardRequest
type GetDashboardRequest = dashboards.GetDashboardRequest
type PinDashboardRequest = dashboards.PinDashboardRequest
type UnpinDashboardRequest = dashboards.UnpinDashboardRequest
type Dashboard = dashboards.Dashboard

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

func NewDashboardsClient(c *CallPropertiesCreator) *DashboardsClient {
	return &DashboardsClient{callPropertiesCreator: c}
}
