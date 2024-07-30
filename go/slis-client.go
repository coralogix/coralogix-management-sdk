package cxsdk

import (
	"context"

	slis "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/catalog/v1"
)

// SLIClient is a client for the Coralogix SLI API.
type SLIClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// CreateSLI creates a new SLI.
func (c SLIClient) CreateSLI(ctx context.Context, req *slis.CreateSliRequest) (*slis.CreateSliResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := slis.NewSliServiceClient(conn)

	return client.CreateSli(callProperties.Ctx, req, callProperties.CallOptions...)
}

// GetSLIs gets the specified SLIs.
func (c SLIClient) GetSLIs(ctx context.Context, req *slis.GetSlisRequest) (*slis.GetSlisResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := slis.NewSliServiceClient(conn)

	return client.GetSlis(callProperties.Ctx, req, callProperties.CallOptions...)
}

// UpdateSLI updates the specified SLI.
func (c SLIClient) UpdateSLI(ctx context.Context, req *slis.UpdateSliRequest) (*slis.UpdateSliResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := slis.NewSliServiceClient(conn)

	return client.UpdateSli(callProperties.Ctx, req, callProperties.CallOptions...)
}

// DeleteSLI deletes the specified SLI.
func (c SLIClient) DeleteSLI(ctx context.Context, req *slis.DeleteSliRequest) (*slis.DeleteSliResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := slis.NewSliServiceClient(conn)

	return client.DeleteSli(callProperties.Ctx, req, callProperties.CallOptions...)
}

// NewSLIsClient creates a new SLIs client.
func NewSLIsClient(c *CallPropertiesCreator) *SLIClient {
	return &SLIClient{callPropertiesCreator: c}
}
