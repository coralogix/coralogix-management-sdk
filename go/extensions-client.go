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

	extensions "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/extensions/v1"
)

const extensionsFeatureGroupID = "extensions"

// GetAllExtensionsRequest is the request type for the GetAllExtensions method.
type GetAllExtensionsRequest = extensions.GetAllExtensionsRequest

// GetAllExtensionsRequestFilter is the filter type for the GetAllExtensions method.
type GetAllExtensionsRequestFilter = extensions.GetAllExtensionsRequest_Filter

// GetAllExtensionsResponse is the response type for the GetAllExtensions method.
type GetAllExtensionsResponse = extensions.GetAllExtensionsResponse

// Extension is the type for an extension.
type Extension = extensions.Extension

// ExtensionDescriptor is the type for an extension description.
type ExtensionDescriptor = extensions.GetAllExtensionsResponse_Extension

// GetExtensionResponse is the response type for the GetExtension method.
type GetExtensionResponse = extensions.GetExtensionResponse

// GetExtensionRequest is the request type for the GetExtension method.
type GetExtensionRequest = extensions.GetExtensionRequest

// GetDeployedExtensionsRequest is the request type for the GetDeployedExtensions method.
type GetDeployedExtensionsRequest = extensions.GetDeployedExtensionsRequest

// GetDeployedExtensionsResponse is the response type for the GetDeployedExtensions method.
type GetDeployedExtensionsResponse = extensions.GetDeployedExtensionsResponse

// DeployExtensionRequest is the request type for the DeployExtension method.
type DeployExtensionRequest = extensions.DeployExtensionRequest

// ExtensionDeployment is the type for an extension deployment.
type ExtensionDeployment = extensions.ExtensionDeployment

// DeployExtensionResponse is the response type for the DeployExtension method.
type DeployExtensionResponse = extensions.DeployExtensionResponse

// UndeployExtensionRequest is the request type for the UndeployExtension method.
type UndeployExtensionRequest = extensions.UndeployExtensionRequest

// UpdateExtensionRequest is the request type for the UpdateExtension method.
type UpdateExtensionRequest = extensions.UpdateExtensionRequest

// RPC methods
const (
	GetAllExtensionsRPC      = extensions.ExtensionService_GetAllExtensions_FullMethodName
	GetExtensionRPC          = extensions.ExtensionService_GetExtension_FullMethodName
	GetDeployedExtensionsRPC = extensions.ExtensionDeploymentService_GetDeployedExtensions_FullMethodName
	DeployExtensionRPC       = extensions.ExtensionDeploymentService_DeployExtension_FullMethodName
	UndeployExtensionRPC     = extensions.ExtensionDeploymentService_UndeployExtension_FullMethodName
	UpdateExtensionRPC       = extensions.ExtensionDeploymentService_UpdateExtension_FullMethodName
)

// The target domains for extensions.
const (
	TargetDomainAlert            = extensions.TargetDomain_ALERT_V3
	TargetDomainAction           = extensions.TargetDomain_ACTION
	TargetDomainCustomDashboard  = extensions.TargetDomain_CX_CUSTOM_DASHBOARD
	TargetDomainEnrichment       = extensions.TargetDomain_ENRICHMENT
	TargetDomainEventsToMetrics  = extensions.TargetDomain_EVENTS_TO_METRICS
	TargetDomainGrafanaDashboard = extensions.TargetDomain_GRAFANA_DASHBOARD
	TargetDomainKibanaDashboard  = extensions.TargetDomain_KIBANA_DASHBOARD
	TargetDomainView             = extensions.TargetDomain_SAVED_VIEW
	TargetDomainParsingRule      = extensions.TargetDomain_PARSING_RULE
	TargetDomainMetricsRuleGroup = extensions.TargetDomain_METRICS_RULE_GROUP
)

// ExtensionsClient is a client for the Coralogix Extensions API.
type ExtensionsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// GetAll returns all extensions.
func (e ExtensionsClient) GetAll(ctx context.Context, req *GetAllExtensionsRequest) (*GetAllExtensionsResponse, error) {
	callProperties, err := e.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := extensions.NewExtensionServiceClient(conn)

	response, err := client.GetAllExtensions(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, GetAllExtensionsRPC, extensionsFeatureGroupID)
	}
	return response, nil
}

// Get returns a single extension by ID.
func (e ExtensionsClient) Get(ctx context.Context, req *GetExtensionRequest) (*GetExtensionResponse, error) {
	callProperties, err := e.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := extensions.NewExtensionServiceClient(conn)

	response, err := client.GetExtension(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, GetExtensionRPC, extensionsFeatureGroupID)
	}
	return response, nil
}

// GetDeployed returns all deployed extensions.
func (e ExtensionsClient) GetDeployed(ctx context.Context, req *GetDeployedExtensionsRequest) (*GetDeployedExtensionsResponse, error) {
	callProperties, err := e.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := extensions.NewExtensionDeploymentServiceClient(conn)

	response, err := client.GetDeployedExtensions(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, GetDeployedExtensionsRPC, extensionsFeatureGroupID)
	}
	return response, nil
}

// Deploy deploys an extension.
func (e ExtensionsClient) Deploy(ctx context.Context, req *DeployExtensionRequest) (*DeployExtensionResponse, error) {
	callProperties, err := e.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := extensions.NewExtensionDeploymentServiceClient(conn)

	response, err := client.DeployExtension(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, DeployExtensionRPC, extensionsFeatureGroupID)
	}
	return response, nil
}

// Undeploy undeploys an extension.
func (e ExtensionsClient) Undeploy(ctx context.Context, req *UndeployExtensionRequest) (*extensions.UndeployExtensionResponse, error) {
	callProperties, err := e.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := extensions.NewExtensionDeploymentServiceClient(conn)

	response, err := client.UndeployExtension(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, UndeployExtensionRPC, extensionsFeatureGroupID)
	}
	return response, nil
}

// Update updates an extension.
func (e ExtensionsClient) Update(ctx context.Context, req *UpdateExtensionRequest) (*extensions.UpdateExtensionResponse, error) {
	callProperties, err := e.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := extensions.NewExtensionDeploymentServiceClient(conn)

	response, err := client.UpdateExtension(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, UpdateExtensionRPC, extensionsFeatureGroupID)
	}
	return response, nil
}

// NewExtensionsClient creates a new ExtensionsClient.
func NewExtensionsClient(callPropertiesCreator *CallPropertiesCreator) *ExtensionsClient {
	return &ExtensionsClient{
		callPropertiesCreator: callPropertiesCreator,
	}
}
