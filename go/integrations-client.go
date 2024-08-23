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

	ext "github.com/coralogix/coralogix-management-sdk/internal/coralogix/integrations/v1"
)

// UpdateIntegrationRequest is a request to update an integration.
type UpdateIntegrationRequest = ext.UpdateIntegrationRequest

// SaveIntegrationRequest is a request to create an integration.
type SaveIntegrationRequest = ext.SaveIntegrationRequest

// DeleteIntegrationRequest is a request to delete an integration.
type DeleteIntegrationRequest = ext.DeleteIntegrationRequest

// GetIntegrationDetailsRequest is a request to get integration details.
type GetIntegrationDetailsRequest = ext.GetIntegrationDetailsRequest

// GetIntegrationDefinitionRequest is a request to get an integration definition.
type GetIntegrationDefinitionRequest = ext.GetIntegrationDefinitionRequest

// GetManagedIntegrationStatusRequest is a request to get the status of a managed integration.
type GetManagedIntegrationStatusRequest = ext.GetManagedIntegrationStatusRequest

// GetTemplateRequest is a request to get an integration template.
type GetTemplateRequest = ext.GetTemplateRequest

// GetRumApplicationVersionDataRequest is a request to get RUM application version data.
type GetRumApplicationVersionDataRequest = ext.GetRumApplicationVersionDataRequest

// SyncRumDataRequest is a request to sync RUM data.
type SyncRumDataRequest = ext.SyncRumDataRequest

// TestIntegrationRequest is a request to test an integration.
type TestIntegrationRequest = ext.TestIntegrationRequest

// IntegrationMetadata is an integration.
type IntegrationMetadata = ext.IntegrationMetadata

// IntegrationMetadataIntegrationParameters is integration parameters.
type IntegrationMetadataIntegrationParameters = ext.IntegrationMetadata_IntegrationParameters

// IntegrationParameter is a parameter for configuring an integration.
type IntegrationParameter = ext.Parameter

// GenericIntegrationParameters are the integration parameters.
type GenericIntegrationParameters = ext.GenericIntegrationParameters

// IntegrationParameterStringValue is a string value parameter.
type IntegrationParameterStringValue = ext.Parameter_StringValue

// IntegrationParameterBooleanValue is a boolean value parameter.
type IntegrationParameterBooleanValue = ext.Parameter_BooleanValue

// IntegrationParameterStringList is a string list parameter.
type IntegrationParameterStringList = ext.Parameter_StringList_

// IntegrationParameterStringListInner is the wrapped string list parameter.
type IntegrationParameterStringListInner = ext.Parameter_StringList

// IntegrationParameterAPIKey is an API key parameter.
type IntegrationParameterAPIKey = ext.Parameter_ApiKey

// IntegrationParameterNumericValue is a numeric value parameter.
type IntegrationParameterNumericValue = ext.Parameter_NumericValue

// IntegrationParameterSensitiveData is a sensitive data parameter.
type IntegrationParameterSensitiveData = ext.Parameter_SensitiveData

// IntegrationTestSuccess indicates whether a test was a success
type IntegrationTestSuccess = ext.TestIntegrationResult_Success_

// IntegrationTestFail indicates whether a test was a failure
type IntegrationTestFail = ext.TestIntegrationResult_Failure_

// IntegrationDetailsDefault the type for all registered instances
type IntegrationDetailsDefault = ext.IntegrationDetails_Default

// GetDeployedIntegrationRequest is a request to get a deployed integration.
type GetDeployedIntegrationRequest = ext.GetDeployedIntegrationRequest

// GetDeployedIntegrationResponse contains the response to a GetDeployedIntegrationRequest.
type GetDeployedIntegrationResponse = ext.GetDeployedIntegrationResponse

// IntegrationsClient is a client for the Coralogix Extensions API.
type IntegrationsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// Create creates a new integration.
func (c IntegrationsClient) Create(ctx context.Context, req *SaveIntegrationRequest) (*ext.SaveIntegrationResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := ext.NewIntegrationServiceClient(conn)

	return client.SaveIntegration(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Update updates an integration
func (c IntegrationsClient) Update(ctx context.Context, req *UpdateIntegrationRequest) (*ext.UpdateIntegrationResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := ext.NewIntegrationServiceClient(conn)

	return client.UpdateIntegration(callProperties.Ctx, req, callProperties.CallOptions...)
}

// GetDetails gets all deployed integrations
func (c IntegrationsClient) GetDetails(ctx context.Context, req *GetIntegrationDetailsRequest) (*ext.GetIntegrationDetailsResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := ext.NewIntegrationServiceClient(conn)

	return client.GetIntegrationDetails(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Get gets a deployed integration
func (c IntegrationsClient) Get(ctx context.Context, req *ext.GetDeployedIntegrationRequest) (*ext.GetDeployedIntegrationResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := ext.NewIntegrationServiceClient(conn)

	return client.GetDeployedIntegration(callProperties.Ctx, req, callProperties.CallOptions...)
}

// GetDefinition gets an integration definition
func (c IntegrationsClient) GetDefinition(ctx context.Context, req *GetIntegrationDefinitionRequest) (*ext.GetIntegrationDefinitionResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := ext.NewIntegrationServiceClient(conn)

	return client.GetIntegrationDefinition(callProperties.Ctx, req, callProperties.CallOptions...)
}

// GetIntegrationStatus gets the status of a integration
func (c IntegrationsClient) GetIntegrationStatus(ctx context.Context, req *GetManagedIntegrationStatusRequest) (*ext.GetManagedIntegrationStatusResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := ext.NewIntegrationServiceClient(conn)

	return client.GetManagedIntegrationStatus(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Delete deletes an integration
func (c IntegrationsClient) Delete(ctx context.Context, req *DeleteIntegrationRequest) (*ext.DeleteIntegrationResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := ext.NewIntegrationServiceClient(conn)

	return client.DeleteIntegration(callProperties.Ctx, req, callProperties.CallOptions...)
}

// GetTemplate gets an integration template
func (c IntegrationsClient) GetTemplate(ctx context.Context, req *GetTemplateRequest) (*ext.GetTemplateResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := ext.NewIntegrationServiceClient(conn)

	return client.GetTemplate(callProperties.Ctx, req, callProperties.CallOptions...)
}

// GetRumApplicationVersionData gets RUM application version data
func (c IntegrationsClient) GetRumApplicationVersionData(ctx context.Context, req *GetRumApplicationVersionDataRequest) (*ext.GetRumApplicationVersionDataResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := ext.NewIntegrationServiceClient(conn)

	return client.GetRumApplicationVersionData(callProperties.Ctx, req, callProperties.CallOptions...)
}

// SyncRumData syncs RUM data
func (c IntegrationsClient) SyncRumData(ctx context.Context, req *SyncRumDataRequest) (*ext.SyncRumDataResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := ext.NewIntegrationServiceClient(conn)

	return client.SyncRumData(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Test tests an integration
func (c IntegrationsClient) Test(ctx context.Context, req *TestIntegrationRequest) (*ext.TestIntegrationResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := ext.NewIntegrationServiceClient(conn)

	return client.TestIntegration(callProperties.Ctx, req, callProperties.CallOptions...)
}

// NewIntegrationsClient creates a new client.
func NewIntegrationsClient(c *CallPropertiesCreator) *IntegrationsClient {
	return &IntegrationsClient{callPropertiesCreator: c}
}
