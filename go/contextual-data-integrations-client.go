// Copyright 2025 Coralogix Ltd.
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

	v1 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/integrations/v1"
)

// RPC method names
const (
	contextualDataIntegrationsFeatureGroupID = "contextual_data_integrations"

	GetContextualDataIntegrationsRPC          = v1.ContextualDataIntegrationService_GetContextualDataIntegrations_FullMethodName
	GetContextualDataIntegrationDefinitionRPC = v1.ContextualDataIntegrationService_GetContextualDataIntegrationDefinition_FullMethodName
	GetContextualDataIntegrationDetailsRPC    = v1.ContextualDataIntegrationService_GetContextualDataIntegrationDetails_FullMethodName
	TestContextualDataIntegrationRPC          = v1.ContextualDataIntegrationService_TestContextualDataIntegration_FullMethodName
	SaveContextualDataIntegrationRPC          = v1.ContextualDataIntegrationService_SaveContextualDataIntegration_FullMethodName
	UpdateContextualDataIntegrationRPC        = v1.ContextualDataIntegrationService_UpdateContextualDataIntegration_FullMethodName
	DeleteContextualDataIntegrationRPC        = v1.ContextualDataIntegrationService_DeleteContextualDataIntegration_FullMethodName
)

// Request/Response type definitions
type (
	// GetContextualDataIntegrationsRequest is the request type for the GetContextualDataIntegrations method
	GetContextualDataIntegrationsRequest = v1.GetContextualDataIntegrationsRequest
	// GetContextualDataIntegrationsResponse is the response type for the GetContextualDataIntegrations method
	GetContextualDataIntegrationsResponse = v1.GetContextualDataIntegrationsResponse
	// GetContextualDataIntegrationDefinitionRequest is the request type for the GetContextualDataIntegrationDefinition method
	GetContextualDataIntegrationDefinitionRequest = v1.GetContextualDataIntegrationDefinitionRequest
	// GetContextualDataIntegrationDefinitionResponse is the response type for the GetContextualDataIntegrationDefinition method
	GetContextualDataIntegrationDefinitionResponse = v1.GetContextualDataIntegrationDefinitionResponse
	// GetContextualDataIntegrationDetailsRequest is the request type for the GetContextualDataIntegrationDetails method
	GetContextualDataIntegrationDetailsRequest = v1.GetContextualDataIntegrationDetailsRequest
	// GetContextualDataIntegrationDetailsResponse is the response type for the GetContextualDataIntegrationDetails method
	GetContextualDataIntegrationDetailsResponse = v1.GetContextualDataIntegrationDetailsResponse
	// TestContextualDataIntegrationRequest is the request type for the TestContextualDataIntegration method
	TestContextualDataIntegrationRequest = v1.TestContextualDataIntegrationRequest
	// TestContextualDataIntegrationResponse is the response type for the TestContextualDataIntegration method
	TestContextualDataIntegrationResponse = v1.TestContextualDataIntegrationResponse
	// SaveContextualDataIntegrationRequest is the request type for the SaveContextualDataIntegration method
	SaveContextualDataIntegrationRequest = v1.SaveContextualDataIntegrationRequest
	// SaveContextualDataIntegrationResponse is the response type for the SaveContextualDataIntegration method
	SaveContextualDataIntegrationResponse = v1.SaveContextualDataIntegrationResponse
	// UpdateContextualDataIntegrationRequest is the request type for the UpdateContextualDataIntegration method
	UpdateContextualDataIntegrationRequest = v1.UpdateContextualDataIntegrationRequest
	// UpdateContextualDataIntegrationResponse is the response type for the UpdateContextualDataIntegration method
	UpdateContextualDataIntegrationResponse = v1.UpdateContextualDataIntegrationResponse
	// DeleteContextualDataIntegrationRequest is the request type for the DeleteContextualDataIntegration method
	DeleteContextualDataIntegrationRequest = v1.DeleteContextualDataIntegrationRequest
	// DeleteContextualDataIntegrationResponse is the response type for the DeleteContextualDataIntegration method
	DeleteContextualDataIntegrationResponse = v1.DeleteContextualDataIntegrationResponse
)

// ContextualDataIntegrationsClient is a client for managing contextual data integrations
type ContextualDataIntegrationsClient struct {
	callPropertiesCreator CallPropertiesCreator
}

// List gets all contextual data integrations
func (c ContextualDataIntegrationsClient) List(ctx context.Context, req *GetContextualDataIntegrationsRequest) (*GetContextualDataIntegrationsResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := v1.NewContextualDataIntegrationServiceClient(conn)

	response, err := client.GetContextualDataIntegrations(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, GetContextualDataIntegrationsRPC, contextualDataIntegrationsFeatureGroupID)
	}
	return response, nil
}

// GetDefinition gets a contextual data integration definition
func (c ContextualDataIntegrationsClient) GetDefinition(ctx context.Context, req *GetContextualDataIntegrationDefinitionRequest) (*GetContextualDataIntegrationDefinitionResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := v1.NewContextualDataIntegrationServiceClient(conn)

	response, err := client.GetContextualDataIntegrationDefinition(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, GetContextualDataIntegrationDefinitionRPC, contextualDataIntegrationsFeatureGroupID)
	}
	return response, nil
}

// GetDetails gets details for a contextual data integration
func (c ContextualDataIntegrationsClient) GetDetails(ctx context.Context, req *GetContextualDataIntegrationDetailsRequest) (*GetContextualDataIntegrationDetailsResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := v1.NewContextualDataIntegrationServiceClient(conn)

	response, err := client.GetContextualDataIntegrationDetails(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, GetContextualDataIntegrationDetailsRPC, contextualDataIntegrationsFeatureGroupID)
	}
	return response, nil
}

// Test tests a contextual data integration
func (c ContextualDataIntegrationsClient) Test(ctx context.Context, req *TestContextualDataIntegrationRequest) (*TestContextualDataIntegrationResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := v1.NewContextualDataIntegrationServiceClient(conn)

	response, err := client.TestContextualDataIntegration(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, TestContextualDataIntegrationRPC, contextualDataIntegrationsFeatureGroupID)
	}
	return response, nil
}

// Create saves a new contextual data integration
func (c ContextualDataIntegrationsClient) Create(ctx context.Context, req *SaveContextualDataIntegrationRequest) (*SaveContextualDataIntegrationResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := v1.NewContextualDataIntegrationServiceClient(conn)

	response, err := client.SaveContextualDataIntegration(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, SaveContextualDataIntegrationRPC, contextualDataIntegrationsFeatureGroupID)
	}
	return response, nil
}

// Update updates an existing contextual data integration
func (c ContextualDataIntegrationsClient) Update(ctx context.Context, req *UpdateContextualDataIntegrationRequest) (*UpdateContextualDataIntegrationResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := v1.NewContextualDataIntegrationServiceClient(conn)

	response, err := client.UpdateContextualDataIntegration(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, UpdateContextualDataIntegrationRPC, contextualDataIntegrationsFeatureGroupID)
	}
	return response, nil
}

// Delete deletes a contextual data integration
func (c ContextualDataIntegrationsClient) Delete(ctx context.Context, req *DeleteContextualDataIntegrationRequest) (*DeleteContextualDataIntegrationResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := v1.NewContextualDataIntegrationServiceClient(conn)

	response, err := client.DeleteContextualDataIntegration(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, DeleteContextualDataIntegrationRPC, contextualDataIntegrationsFeatureGroupID)
	}
	return response, nil
}

// NewContextualDataIntegrationsClient creates a new contextual data integrations client
func NewContextualDataIntegrationsClient(c CallPropertiesCreator) *ContextualDataIntegrationsClient {
	return &ContextualDataIntegrationsClient{callPropertiesCreator: c}
}
