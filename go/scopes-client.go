package cxsdk

import (
	"context"
	slos "coralogix-management-sdk/go/internal/coralogixapis/scopes/v1"
)

type CreateScopeRequest = slos.CreateScopeRequest
type GetTeamScopesByIdsRequest = slos.GetTeamScopesByIdsRequest
type GetTeamScopesRequest = slos.GetTeamScopesRequest
type UpdateScopeRequest = slos.UpdateScopeRequest
type DeleteScopeRequest = slos.DeleteScopeRequest

type Filter = slos.Filter

const (
	EntityType_UNSPECIFIED = slos.EntityType_UNSPECIFIED
	EntityType_LOGS        = slos.EntityType_LOGS
	EntityType_SPANS       = slos.EntityType_SPANS
)

type ScopesClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// Create a new scope
func (c ScopesClient) Create(ctx context.Context, req *CreateScopeRequest) (*slos.CreateScopeResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := slos.NewScopesServiceClient(conn)

	return client.CreateScope(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Get a scope by its ID
func (c ScopesClient) Get(ctx context.Context, req *GetTeamScopesByIdsRequest) (*slos.GetScopesResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := slos.NewScopesServiceClient(conn)

	return client.GetTeamScopesByIds(callProperties.Ctx, req, callProperties.CallOptions...)
}

// List all scopes for the current team
func (c ScopesClient) List(ctx context.Context, req *GetTeamScopesRequest) (*slos.GetScopesResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := slos.NewScopesServiceClient(conn)

	return client.GetTeamScopes(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Update a scope
func (c ScopesClient) Update(ctx context.Context, req *UpdateScopeRequest) (*slos.UpdateScopeResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := slos.NewScopesServiceClient(conn)

	return client.UpdateScope(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Delete a scope
func (c ScopesClient) Delete(ctx context.Context, req *DeleteScopeRequest) (*slos.DeleteScopeResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := slos.NewScopesServiceClient(conn)

	return client.DeleteScope(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Create a new ScopesClient
func NewScopesClient(c *CallPropertiesCreator) *ScopesClient {
	return &ScopesClient{callPropertiesCreator: c}
}
