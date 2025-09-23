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

	common "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/notification_center/common"
	commonv1 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/notification_center/common/v1"
	"github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/notification_center/common/v1/routing"
	connectors "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/notification_center/connectors/v1"
	notifications "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/notification_center/notifications/v1"
	presets "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/notification_center/presets/v1"
	routers "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/notification_center/routers/v1"
)

// CreateConnectorRequest represents a request to create a connector.
type CreateConnectorRequest = connectors.CreateConnectorRequest

// CreateConnectorResponse represents a response to create a connector.
type CreateConnectorResponse = connectors.CreateConnectorResponse

// ReplaceConnectorRequest represents a request to replace a connector.
type ReplaceConnectorRequest = connectors.ReplaceConnectorRequest

// ReplaceConnectorResponse represents a response to replace a connector.
type ReplaceConnectorResponse = connectors.ReplaceConnectorResponse

// DeleteConnectorRequest represents a request to delete a connector.
type DeleteConnectorRequest = connectors.DeleteConnectorRequest

// DeleteConnectorResponse represents a response to delete a connector.
type DeleteConnectorResponse = connectors.DeleteConnectorResponse

// GetConnectorRequest represents a request to get a connector.
type GetConnectorRequest = connectors.GetConnectorRequest

// GetConnectorResponse represents a response to get a connector.
type GetConnectorResponse = connectors.GetConnectorResponse

// ListConnectorsRequest represents a request to list connectors.
type ListConnectorsRequest = connectors.ListConnectorsRequest

// ListConnectorsResponse represents a response to list connectors.
type ListConnectorsResponse = connectors.ListConnectorsResponse

// BatchGetConnectorsRequest represents a request to batch get connectors.
type BatchGetConnectorsRequest = connectors.BatchGetConnectorsRequest

// BatchGetConnectorsResponse represents a response to batch get connectors.
type BatchGetConnectorsResponse = connectors.BatchGetConnectorsResponse

// GetConnectorTypeSummariesRequest represents a request to get connector type summaries.
type GetConnectorTypeSummariesRequest = connectors.GetConnectorTypeSummariesRequest

// GetConnectorTypeSummariesResponse represents a response to get connector type summaries.
type GetConnectorTypeSummariesResponse = connectors.GetConnectorTypeSummariesResponse

// Connector represents a connector.
type Connector = connectors.Connector

// ConnectorConfig represents a connector configuration.
type ConnectorConfig = connectors.ConnectorConfig

// EntityTypeConfigOverrides represents a connector configuration overrides.
type EntityTypeConfigOverrides = connectors.EntityTypeConfigOverrides

// ConnectorConfigField represents a connector configuration field.
type ConnectorConfigField = commonv1.ConnectorConfigField

// ConnectorType is a connector type.
type ConnectorType = common.ConnectorType

// SetPresetAsDefaultRequest is a request to set a preset as default.
type SetPresetAsDefaultRequest = presets.SetPresetAsDefaultRequest

// EntityType values.
const (
	EntityTypeAlerts = common.EntityType_ALERTS
)

// ConnectorType values.
const (
	ConnectorTypeUnSpecified  = common.ConnectorType_CONNECTOR_TYPE_UNSPECIFIED
	ConnectorTypeSlack        = common.ConnectorType_SLACK
	ConnectorTypeGenericHTTPS = common.ConnectorType_GENERIC_HTTPS
	ConnectorTypePagerDuty    = common.ConnectorType_PAGERDUTY
)

// CreateCustomPresetRequest is a request to create a custom preset.
type CreateCustomPresetRequest = presets.CreateCustomPresetRequest

// CreateCustomPresetResponse is a response to create a custom preset.
type CreateCustomPresetResponse = presets.CreateCustomPresetResponse

// ReplaceCustomPresetRequest is a request to replace a custom preset.
type ReplaceCustomPresetRequest = presets.ReplaceCustomPresetRequest

// ReplaceCustomPresetResponse is a response to replace a custom preset.
type ReplaceCustomPresetResponse = presets.ReplaceCustomPresetResponse

// DeleteCustomPresetRequest is a request to delete a custom preset.
type DeleteCustomPresetRequest = presets.DeleteCustomPresetRequest

// DeleteCustomPresetResponse is a response to delete a custom preset.
type DeleteCustomPresetResponse = presets.DeleteCustomPresetResponse

// SetCustomPresetAsDefaultRequest is a request to set a custom preset as default.
type SetCustomPresetAsDefaultRequest = presets.SetCustomPresetAsDefaultRequest

// SetCustomPresetAsDefaultResponse is a response to set a custom preset as default.
type SetCustomPresetAsDefaultResponse = presets.SetCustomPresetAsDefaultResponse

// GetPresetRequest is a request to get a preset.
type GetPresetRequest = presets.GetPresetRequest

// GetPresetResponse is a response to get a preset.
type GetPresetResponse = presets.GetPresetResponse

// ListPresetSummariesRequest is a request to list preset summaries.
type ListPresetSummariesRequest = presets.ListPresetSummariesRequest

// ListPresetSummariesResponse is a response to list preset summaries.
type ListPresetSummariesResponse = presets.ListPresetSummariesResponse

// BatchGetPresetsRequest is a request to batch get presets.
type BatchGetPresetsRequest = presets.BatchGetPresetsRequest

// BatchGetPresetsResponse is a response to batch get presets.
type BatchGetPresetsResponse = presets.BatchGetPresetsResponse

// GetDefaultPresetSummaryRequest is a request to get the default preset summary.
type GetDefaultPresetSummaryRequest = presets.GetDefaultPresetSummaryRequest

// GetDefaultPresetSummaryResponse is a response to get the default preset summary.
type GetDefaultPresetSummaryResponse = presets.GetDefaultPresetSummaryResponse

// GetSystemDefaultPresetSummaryRequest is a request to get the system default preset summary.
type GetSystemDefaultPresetSummaryRequest = presets.GetSystemDefaultPresetSummaryRequest

// GetSystemDefaultPresetSummaryResponse is a response to get the system default preset summary.
type GetSystemDefaultPresetSummaryResponse = presets.GetSystemDefaultPresetSummaryResponse

// Preset represents a preset.
type Preset = presets.Preset

// ConfigOverrides is a configuration override.
type ConfigOverrides = common.ConfigOverrides

// ConditionType is a condition type.
type ConditionType = common.ConditionType

// ConditionTypeMatchEntityType is a match entity type condition type.
type ConditionTypeMatchEntityType = common.ConditionType_MatchEntityType

// MatchEntityTypeCondition is a match entity type condition.
type MatchEntityTypeCondition = common.MatchEntityTypeCondition

// ConditionTypeMatchEntityTypeAndSubType is a match entity type and subtype condition type.
type ConditionTypeMatchEntityTypeAndSubType = common.ConditionType_MatchEntityTypeAndSubType

// MatchEntityTypeAndSubTypeCondition is a match entity type and subtype condition.
type MatchEntityTypeAndSubTypeCondition = common.MatchEntityTypeAndSubTypeCondition

// RoutingRule is a routing target.
type RoutingRule = routing.RoutingRule

// RoutingTarget is a routing target.
type RoutingTarget = routing.RoutingTarget

// PresetType is a preset type.
type PresetType = presets.PresetType

// PresetType values.
const (
	PresetTypeUnspecified = presets.PresetType_PRESET_TYPE_UNSPECIFIED
	PresetTypeSystem      = presets.PresetType_SYSTEM
	PresetTypeCustom      = presets.PresetType_CUSTOM
)

// NotificationsEntityType is a type of entity.
type NotificationsEntityType = common.EntityType

// NotificationsEntityType values.
const (
	NotificationsEntityTypeUnspecified = common.EntityType_ENTITY_TYPE_UNSPECIFIED
	NotificationsEntityTypeAlerts      = common.EntityType_ALERTS
)

const (
	RouterEvaluationModeUnspecified      = routers.RouterEvaluationMode_ROUTER_EVALUATION_MODE_UNSPECIFIED
	RouterEvaluationModeEvaluateAll      = routers.RouterEvaluationMode_EVALUATE_ALL
	RouterEvaluationModeStopOnFirstMatch = routers.RouterEvaluationMode_STOP_ON_FIRST_MATCH
)

// CreateGlobalRouterRequest is a request to create a global router.
type CreateGlobalRouterRequest = routers.CreateGlobalRouterRequest

// CreateGlobalRouterResponse is a response to create a global router.
type CreateGlobalRouterResponse = routers.CreateGlobalRouterResponse

// ReplaceGlobalRouterRequest is a request to replace a global router.
type ReplaceGlobalRouterRequest = routers.ReplaceGlobalRouterRequest

// ReplaceGlobalRouterResponse is a response to replace a global router.
type ReplaceGlobalRouterResponse = routers.ReplaceGlobalRouterResponse

// CreateOrReplaceGlobalRouterRequest is a request to create or replace a global router.
type CreateOrReplaceGlobalRouterRequest = routers.CreateOrReplaceGlobalRouterRequest

// CreateOrReplaceGlobalRouterResponse is a response to create or replace a global router.
type CreateOrReplaceGlobalRouterResponse = routers.CreateOrReplaceGlobalRouterResponse

// DeleteGlobalRouterRequest is a request to delete a global router.
type DeleteGlobalRouterRequest = routers.DeleteGlobalRouterRequest

// DeleteGlobalRouterResponse is a response to delete a global router.
type DeleteGlobalRouterResponse = routers.DeleteGlobalRouterResponse

// GetGlobalRouterRequest is a request to get a global router.
type GetGlobalRouterRequest = routers.GetGlobalRouterRequest

// GetGlobalRouterResponse is a response to get a global router.
type GetGlobalRouterResponse = routers.GetGlobalRouterResponse

// ListGlobalRoutersRequest is a request to list global routers.
type ListGlobalRoutersRequest = routers.ListGlobalRoutersRequest

// ListGlobalRoutersResponse is a response to list global routers.
type ListGlobalRoutersResponse = routers.ListGlobalRoutersResponse

// BatchGetGlobalRoutersRequest is a request to batch get global routers.
type BatchGetGlobalRoutersRequest = routers.BatchGetGlobalRoutersRequest

// BatchGetGlobalRoutersResponse is a response to batch get global routers.
type BatchGetGlobalRoutersResponse = routers.BatchGetGlobalRoutersResponse

// GlobalRouter represents a global router.
type GlobalRouter = routers.GlobalRouter

// MessageConfig is a message configuration.
type MessageConfig = common.MessageConfig

// MessageConfigField is a message configuration field.
type MessageConfigField = commonv1.MessageConfigField

// TemplatedConnectorConfigField is a templated message configuration field.
type TemplatedConnectorConfigField = commonv1.TemplatedConnectorConfigField

// TestConnectorConfigRequest is a request to test a connector configuration.
type TestConnectorConfigRequest = notifications.TestConnectorConfigRequest

// TestConnectorConfigResponse is a response to test a connector configuration.
type TestConnectorConfigResponse = notifications.TestConnectorConfigResponse

// TestExistingConnectorRequest is a request to test an existing connector.
type TestExistingConnectorRequest = notifications.TestExistingConnectorRequest

// TestExistingConnectorResponse is a response to test an existing connector.
type TestExistingConnectorResponse = notifications.TestExistingConnectorResponse

// TestExistingPresetConfigRequest is a request to test an existing preset configuration.
type TestExistingPresetConfigRequest = notifications.TestExistingPresetRequest

// TestExistingPresetConfigResponse is a response to test an existing preset configuration.
type TestExistingPresetConfigResponse = notifications.TestExistingPresetResponse

// TestPresetConfigRequest is a request to test a preset configuration.
type TestPresetConfigRequest = notifications.TestPresetConfigRequest

// TestPresetConfigResponse is a response to test a preset configuration.
type TestPresetConfigResponse = notifications.TestPresetConfigResponse

// TestTemplateRenderRequest is a request to test a template rendering.
type TestTemplateRenderRequest = notifications.TestTemplateRenderRequest

// TestTemplateRenderResponse is a response to test a template rendering.
type TestTemplateRenderResponse = notifications.TestTemplateRenderResponse

// TestDestinationRequest is a request for testing a destination.
type TestDestinationRequest = notifications.TestDestinationRequest

// TestDestinationResponse is a response for testing a destination.
type TestDestinationResponse = notifications.TestDestinationResponse

// TestRoutingConditionValidRequest is a request for testing a routing condition.
type TestRoutingConditionValidRequest = notifications.TestRoutingConditionValidRequest

// TestRoutingConditionValidResponse is a response for testing a routing condition.
type TestRoutingConditionValidResponse = notifications.TestRoutingConditionValidResponse

// SetPresetAsDefaultResponse is a response to set a preset as default.
type SetPresetAsDefaultResponse = presets.SetPresetAsDefaultResponse

const notificationsFeatureGroupID = "notifications"

// RPC names.
const (
	ConnectorsCreateRPC                    = connectors.ConnectorsService_CreateConnector_FullMethodName
	ConnectorsReplaceRPC                   = connectors.ConnectorsService_ReplaceConnector_FullMethodName
	ConnectorsDeleteRPC                    = connectors.ConnectorsService_DeleteConnector_FullMethodName
	ConnectorsGetRPC                       = connectors.ConnectorsService_GetConnector_FullMethodName
	ConnectorsListRPC                      = connectors.ConnectorsService_ListConnectors_FullMethodName
	ConnectorsBatchGetRPC                  = connectors.ConnectorsService_BatchGetConnectors_FullMethodName
	ConnectorsGetConnectorTypeSummariesRPC = connectors.ConnectorsService_GetConnectorTypeSummaries_FullMethodName
	PresetsCreateRPC                       = presets.PresetsService_CreateCustomPreset_FullMethodName
	PresetsReplaceRPC                      = presets.PresetsService_ReplaceCustomPreset_FullMethodName
	PresetsDeleteRPC                       = presets.PresetsService_DeleteCustomPreset_FullMethodName
	PresetsSetAsDefaultRPC                 = presets.PresetsService_SetCustomPresetAsDefault_FullMethodName
	PresetsGetRPC                          = presets.PresetsService_GetPreset_FullMethodName
	PresetsListRPC                         = presets.PresetsService_ListPresetSummaries_FullMethodName
	PresetsBatchGetRPC                     = presets.PresetsService_BatchGetPresets_FullMethodName
	PresetsGetDefaultRPC                   = presets.PresetsService_GetDefaultPresetSummary_FullMethodName
	PresetsGetSystemDefaultRPC             = presets.PresetsService_GetSystemDefaultPresetSummary_FullMethodName
	GlobalRoutersCreateRPC                 = routers.GlobalRoutersService_CreateGlobalRouter_FullMethodName
	GlobalRoutersReplaceRPC                = routers.GlobalRoutersService_ReplaceGlobalRouter_FullMethodName
	GlobalRoutersDeleteRPC                 = routers.GlobalRoutersService_DeleteGlobalRouter_FullMethodName
	GlobalRoutersGetRPC                    = routers.GlobalRoutersService_GetGlobalRouter_FullMethodName
	GlobalRoutersListRPC                   = routers.GlobalRoutersService_ListGlobalRouters_FullMethodName
	GlobalRoutersBatchGetRPC               = routers.GlobalRoutersService_BatchGetGlobalRouters_FullMethodName
	TestingTestConnectorConfigRPC          = notifications.TestingService_TestConnectorConfig_FullMethodName
	TestingTestExistingConnectorRPC        = notifications.TestingService_TestExistingConnector_FullMethodName
	TestingTestPresetConfigRPC             = notifications.TestingService_TestPresetConfig_FullMethodName
	TestingTestTemplateRenderRPC           = notifications.TestingService_TestTemplateRender_FullMethodName
	TestingTestDestinationRPC              = notifications.TestingService_TestDestination_FullMethodName
)

// NotificationsClient is a client for the Coralogix Notifications API.
type NotificationsClient struct {
	callPropertiesCreator CallPropertiesCreator
}

// CreateConnector creates a new connector.
func (c NotificationsClient) CreateConnector(ctx context.Context, req *CreateConnectorRequest) (*CreateConnectorResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := connectors.NewConnectorsServiceClient(conn)

	response, err := client.CreateConnector(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, ConnectorsCreateRPC, notificationsFeatureGroupID)
	}
	return response, nil
}

// ReplaceConnector replaces a connector.
func (c NotificationsClient) ReplaceConnector(ctx context.Context, req *ReplaceConnectorRequest) (*ReplaceConnectorResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := connectors.NewConnectorsServiceClient(conn)

	response, err := client.ReplaceConnector(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, ConnectorsReplaceRPC, notificationsFeatureGroupID)
	}
	return response, nil
}

// DeleteConnector deletes a connector.
func (c NotificationsClient) DeleteConnector(ctx context.Context, req *DeleteConnectorRequest) (*DeleteConnectorResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := connectors.NewConnectorsServiceClient(conn)

	response, err := client.DeleteConnector(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, ConnectorsDeleteRPC, notificationsFeatureGroupID)
	}
	return response, nil
}

// GetConnector retrieves a connector by ID.
func (c NotificationsClient) GetConnector(ctx context.Context, req *GetConnectorRequest) (*GetConnectorResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetUserLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := connectors.NewConnectorsServiceClient(conn)

	response, err := client.GetConnector(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, ConnectorsGetRPC, notificationsFeatureGroupID)
	}
	return response, nil
}

// ListConnectors lists connectors.
func (c NotificationsClient) ListConnectors(ctx context.Context, req *ListConnectorsRequest) (*ListConnectorsResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := connectors.NewConnectorsServiceClient(conn)

	response, err := client.ListConnectors(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, ConnectorsListRPC, notificationsFeatureGroupID)
	}
	return response, nil
}

// BatchGetConnectors retrieves connectors by IDs.
func (c NotificationsClient) BatchGetConnectors(ctx context.Context, req *BatchGetConnectorsRequest) (*BatchGetConnectorsResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}
	conn := callProperties.Connection
	defer conn.Close()
	client := connectors.NewConnectorsServiceClient(conn)

	response, err := client.BatchGetConnectors(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, ConnectorsBatchGetRPC, notificationsFeatureGroupID)
	}
	return response, nil
}

// GetConnectorTypeSummaries retrieves connector type summaries.
func (c NotificationsClient) GetConnectorTypeSummaries(ctx context.Context, req *GetConnectorTypeSummariesRequest) (*GetConnectorTypeSummariesResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}
	conn := callProperties.Connection
	defer conn.Close()
	client := connectors.NewConnectorsServiceClient(conn)

	response, err := client.GetConnectorTypeSummaries(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, ConnectorsGetConnectorTypeSummariesRPC, notificationsFeatureGroupID)
	}
	return response, nil
}

// CreateCustomPreset creates a new custom preset.
func (c NotificationsClient) CreateCustomPreset(ctx context.Context, req *CreateCustomPresetRequest) (*CreateCustomPresetResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}
	conn := callProperties.Connection
	defer conn.Close()
	client := presets.NewPresetsServiceClient(conn)

	response, err := client.CreateCustomPreset(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, PresetsCreateRPC, notificationsFeatureGroupID)
	}
	return response, nil
}

// ReplaceCustomPreset replaces a custom preset.
func (c NotificationsClient) ReplaceCustomPreset(ctx context.Context, req *ReplaceCustomPresetRequest) (*ReplaceCustomPresetResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}
	conn := callProperties.Connection
	defer conn.Close()
	client := presets.NewPresetsServiceClient(conn)

	response, err := client.ReplaceCustomPreset(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, PresetsReplaceRPC, notificationsFeatureGroupID)
	}
	return response, nil
}

// DeleteCustomPreset deletes a custom preset.
func (c NotificationsClient) DeleteCustomPreset(ctx context.Context, req *DeleteCustomPresetRequest) (*DeleteCustomPresetResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}
	conn := callProperties.Connection
	defer conn.Close()
	client := presets.NewPresetsServiceClient(conn)

	response, err := client.DeleteCustomPreset(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, PresetsDeleteRPC, notificationsFeatureGroupID)
	}
	return response, nil
}

// SetCustomPresetAsDefault sets a custom preset as default.
func (c NotificationsClient) SetCustomPresetAsDefault(ctx context.Context, req *SetCustomPresetAsDefaultRequest) (*SetCustomPresetAsDefaultResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}
	conn := callProperties.Connection
	defer conn.Close()
	client := presets.NewPresetsServiceClient(conn)

	response, err := client.SetCustomPresetAsDefault(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, PresetsSetAsDefaultRPC, notificationsFeatureGroupID)
	}
	return response, nil
}

// GetPreset retrieves a preset.
func (c NotificationsClient) GetPreset(ctx context.Context, req *GetPresetRequest) (*GetPresetResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}
	conn := callProperties.Connection
	defer conn.Close()
	client := presets.NewPresetsServiceClient(conn)

	response, err := client.GetPreset(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, PresetsGetRPC, notificationsFeatureGroupID)
	}
	return response, nil
}

// ListPresetSummaries lists preset summaries.
func (c NotificationsClient) ListPresetSummaries(ctx context.Context, req *ListPresetSummariesRequest) (*ListPresetSummariesResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}
	conn := callProperties.Connection
	defer conn.Close()
	client := presets.NewPresetsServiceClient(conn)

	response, err := client.ListPresetSummaries(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, PresetsListRPC, notificationsFeatureGroupID)
	}
	return response, nil
}

// BatchGetPresets retrieves a group of presets.
func (c NotificationsClient) BatchGetPresets(ctx context.Context, req *BatchGetPresetsRequest) (*BatchGetPresetsResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}
	conn := callProperties.Connection
	defer conn.Close()
	client := presets.NewPresetsServiceClient(conn)

	response, err := client.BatchGetPresets(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, PresetsBatchGetRPC, notificationsFeatureGroupID)
	}
	return response, nil
}

// GetDefaultPresetSummary retrieves the default preset summary.
func (c NotificationsClient) GetDefaultPresetSummary(ctx context.Context, req *GetDefaultPresetSummaryRequest) (*GetDefaultPresetSummaryResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}
	conn := callProperties.Connection
	defer conn.Close()
	client := presets.NewPresetsServiceClient(conn)
	response, err := client.GetDefaultPresetSummary(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, PresetsGetDefaultRPC, notificationsFeatureGroupID)
	}
	return response, nil
}

// SetPresetAsDefault sets a preset as default.
func (c NotificationsClient) SetPresetAsDefault(ctx context.Context, req *SetPresetAsDefaultRequest) (*SetPresetAsDefaultResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}
	conn := callProperties.Connection
	defer conn.Close()
	client := presets.NewPresetsServiceClient(conn)

	response, err := client.SetPresetAsDefault(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, PresetsSetAsDefaultRPC, notificationsFeatureGroupID)
	}
	return response, nil
}

// GetSystemDefaultPresetSummary retrieves the system default preset summary.
func (c NotificationsClient) GetSystemDefaultPresetSummary(ctx context.Context, req *GetSystemDefaultPresetSummaryRequest) (*GetSystemDefaultPresetSummaryResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}
	conn := callProperties.Connection
	defer conn.Close()
	client := presets.NewPresetsServiceClient(conn)

	response, err := client.GetSystemDefaultPresetSummary(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, PresetsGetSystemDefaultRPC, notificationsFeatureGroupID)
	}
	return response, nil
}

// CreateOrReplaceGlobalRouter creates or replaces a global router.
func (c NotificationsClient) CreateOrReplaceGlobalRouter(ctx context.Context, req *CreateOrReplaceGlobalRouterRequest) (*CreateOrReplaceGlobalRouterResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}
	conn := callProperties.Connection
	defer conn.Close()
	client := routers.NewGlobalRoutersServiceClient(conn)

	response, err := client.CreateOrReplaceGlobalRouter(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, GlobalRoutersReplaceRPC, notificationsFeatureGroupID)
	}
	return response, nil
}

// DeleteGlobalRouter deletes a global router.
func (c NotificationsClient) DeleteGlobalRouter(ctx context.Context, req *DeleteGlobalRouterRequest) (*DeleteGlobalRouterResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}
	conn := callProperties.Connection
	defer conn.Close()
	client := routers.NewGlobalRoutersServiceClient(conn)
	response, err := client.DeleteGlobalRouter(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, GlobalRoutersDeleteRPC, notificationsFeatureGroupID)
	}
	return response, nil
}

// GetGlobalRouter retrieves a global router.
func (c NotificationsClient) GetGlobalRouter(ctx context.Context, req *GetGlobalRouterRequest) (*GetGlobalRouterResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}
	conn := callProperties.Connection
	defer conn.Close()
	client := routers.NewGlobalRoutersServiceClient(conn)

	response, err := client.GetGlobalRouter(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, GlobalRoutersGetRPC, notificationsFeatureGroupID)
	}
	return response, nil
}

// ListGlobalRouters lists global routers.
func (c NotificationsClient) ListGlobalRouters(ctx context.Context, req *ListGlobalRoutersRequest) (*ListGlobalRoutersResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}
	conn := callProperties.Connection
	defer conn.Close()
	client := routers.NewGlobalRoutersServiceClient(conn)

	response, err := client.ListGlobalRouters(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, GlobalRoutersListRPC, notificationsFeatureGroupID)
	}
	return response, nil
}

// BatchGetGlobalRouters retrieves global routers by IDs.
func (c NotificationsClient) BatchGetGlobalRouters(ctx context.Context, req *BatchGetGlobalRoutersRequest) (*BatchGetGlobalRoutersResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}
	conn := callProperties.Connection
	defer conn.Close()
	client := routers.NewGlobalRoutersServiceClient(conn)

	response, err := client.BatchGetGlobalRouters(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, GlobalRoutersBatchGetRPC, notificationsFeatureGroupID)
	}
	return response, nil
}

// TestConnectorConfig tests a connector configuration.
func (c NotificationsClient) TestConnectorConfig(ctx context.Context, req *TestConnectorConfigRequest) (*TestConnectorConfigResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetUserLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := notifications.NewTestingServiceClient(conn)

	response, err := client.TestConnectorConfig(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, TestingTestDestinationRPC, notificationsFeatureGroupID)
	}
	return response, nil
}

// TestExistingConnector tests an existing connector.
func (c NotificationsClient) TestExistingConnector(ctx context.Context, req *TestExistingConnectorRequest) (*TestExistingConnectorResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetUserLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := notifications.NewTestingServiceClient(conn)

	response, err := client.TestExistingConnector(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, TestingTestExistingConnectorRPC, notificationsFeatureGroupID)
	}
	return response, nil
}

// TestPresetConfig tests a preset configuration.
func (c NotificationsClient) TestPresetConfig(ctx context.Context, req *TestPresetConfigRequest) (*TestPresetConfigResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := notifications.NewTestingServiceClient(conn)

	response, err := client.TestPresetConfig(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, TestingTestPresetConfigRPC, notificationsFeatureGroupID)
	}
	return response, nil
}

// TestExistingPresetConfig tests a preset configuration.
func (c NotificationsClient) TestExistingPresetConfig(ctx context.Context, req *TestExistingPresetConfigRequest) (*TestExistingPresetConfigResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := notifications.NewTestingServiceClient(conn)

	response, err := client.TestExistingPreset(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, TestingTestPresetConfigRPC, notificationsFeatureGroupID)
	}
	return response, nil
}

// TestTemplateRender tests a template rendering.
func (c NotificationsClient) TestTemplateRender(ctx context.Context, req *TestTemplateRenderRequest) (*TestTemplateRenderResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := notifications.NewTestingServiceClient(conn)

	response, err := client.TestTemplateRender(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, TestingTestTemplateRenderRPC, notificationsFeatureGroupID)
	}
	return response, nil
}

// TestDestination tests a destination.
func (c NotificationsClient) TestDestination(ctx context.Context, req *TestDestinationRequest) (*TestDestinationResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := notifications.NewTestingServiceClient(conn)

	response, err := client.TestDestination(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, TestingTestDestinationRPC, notificationsFeatureGroupID)
	}
	return response, nil
}

// TestRoutingConditionValid tests a routing condition.
func (c NotificationsClient) TestRoutingConditionValid(ctx context.Context, req *TestRoutingConditionValidRequest) (*TestRoutingConditionValidResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := notifications.NewTestingServiceClient(conn)

	response, err := client.TestRoutingConditionValid(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, TestingTestTemplateRenderRPC, notificationsFeatureGroupID)
	}
	return response, nil
}

// NewNotificationsClient creates a new notifications' client.
func NewNotificationsClient(c CallPropertiesCreator) *NotificationsClient {
	return &NotificationsClient{callPropertiesCreator: c}
}
