package cxsdk

import (
	"context"
	slos "coralogix-management-sdk/go/internal/coralogixapis/apm/services/v1"
)

// SLOsClient is a client for the Coralogix SLOs API.
type SLOsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// CreateSLO creates a new SLO.
func (c SLOsClient) CreateSLO(ctx context.Context, req *slos.CreateServiceSloRequest) (*slos.CreateServiceSloResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := slos.NewServiceSloServiceClient(conn)

	return client.CreateServiceSlo(callProperties.Ctx, req, callProperties.CallOptions...)
}

// GetSLO gets the specified SLO.
func (c SLOsClient) GetSLO(ctx context.Context, req *slos.GetServiceSloRequest) (*slos.GetServiceSloResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := slos.NewServiceSloServiceClient(conn)

	return client.GetServiceSlo(callProperties.Ctx, req, callProperties.CallOptions...)
}

// UpdateSLO updates the specified SLO.
func (c SLOsClient) UpdateSLO(ctx context.Context, req *slos.ReplaceServiceSloRequest) (*slos.ReplaceServiceSloResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := slos.NewServiceSloServiceClient(conn)

	return client.ReplaceServiceSlo(callProperties.Ctx, req, callProperties.CallOptions...)
}

// DeleteSLO deletes the specified SLO.
func (c SLOsClient) DeleteSLO(ctx context.Context, req *slos.DeleteServiceSloRequest) (*slos.DeleteServiceSloResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := slos.NewServiceSloServiceClient(conn)

	return client.DeleteServiceSlo(callProperties.Ctx, req, callProperties.CallOptions...)
}

// NewSLOsClient creates a new SLOs client.
func NewSLOsClient(c *CallPropertiesCreator) *SLOsClient {
	return &SLOsClient{callPropertiesCreator: c}
}
