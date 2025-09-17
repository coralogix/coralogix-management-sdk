// Copyright 2024 Coralogix Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package examples

//
//import (
//	"context"
//	"fmt"
//	"testing"
//
//	"github.com/google/uuid"
//	"github.com/stretchr/testify/assert"
//
//	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
//	connectors "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/connectors_service"
//	presets "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/presets_service"
//	routing "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/global_routers_service"
//)
//
//func TestHttpsConnector(t *testing.T) {
//	t.Skip("Skipping because OpenAPI Facade is missing notification center APIs")
//	client := cxsdk.NewClientBuilder().WithRegionFromEnv().WithAPIKeyFromEnv().Build()
//
//	name := fmt.Sprintf("TestConnector-%v", uuid.NewString())
//
//	connector := cxsdk.Connector{
//		Name:        name,
//		Type:        "GENERIC_HTTPS",
//		Description: cxsdk.PtrString("This is the connector to use for Notification Center testing."),
//		ConnectorConfig: &cxsdk.ComCoralogixapisNotificationCenterConnectorsV1ConnectorConfig{
//			Fields: []cxsdk.ComCoralogixapisNotificationCenterConnectorConfigField{
//				{FieldName: cxsdk.PtrString("url"), Value: cxsdk.PtrString("https://httpbun.org/post")},
//				{FieldName: cxsdk.PtrString("method"), Value: cxsdk.PtrString("post")},
//			},
//		},
//		ConfigOverrides: []cxsdk.ComCoralogixapisNotificationCenterConnectorsV1EntityTypeConfigOverrides{
//			{
//				EntityType: cxsdk.PtrString("ALERTS"),
//				Fields: []cxsdk.ComCoralogixapisNotificationCenterTemplatedConnectorConfigField{
//					{
//						FieldName: cxsdk.PtrString("additionalBodyFields"),
//						Template:  cxsdk.PtrString("{\"priority\": \"{{alertDef.priority}}\"}"),
//					},
//				},
//			},
//		},
//	}
//
//	created, _, err := client.DefaultAPI.
//		ConnectorsServiceCreateConnector(context.Background()).
//		Connector(connector).
//		Execute()
//	assertNilAndPrintError(t, err)
//
//	connectorID := created.Connector.Id
//	assert.NotNil(t, connectorID)
//
//	got, _, err := client.DefaultAPI.
//		ConnectorsServiceGetConnector(context.Background(), *connectorID).
//		Execute()
//	assertNilAndPrintError(t, err)
//	assert.Equal(t, name, got.Connector.Name)
//
//	testReq := cxsdk.TestingServiceTestExistingConnectorRequest{
//		ConnectorId: connectorID,
//		PayloadType: cxsdk.PtrString("generic_https_default"),
//	}
//
//	_, _, err = client.DefaultAPI.
//		TestingServiceTestExistingConnector(context.Background()).
//		TestingServiceTestExistingConnectorRequest(testReq).
//		Execute()
//	assertNilAndPrintError(t, err)
//
//	_, _, err = client.DefaultAPI.
//		ConnectorsServiceDeleteConnector(context.Background(), *connectorID).
//		Execute()
//	assertNilAndPrintError(t, err)
//}
//
//func TestSlackConnector(t *testing.T) {
//	t.Skip("Skipping because OpenAPI Facade is missing notification center APIs")
//
//	client := cxsdk.NewClientBuilder().WithRegionFromEnv().WithAPIKeyFromEnv().Build()
//
//	connector := cxsdk.Connector{
//		Name:        "TestSlackConnector",
//		Type:        "SLACK",
//		Description: cxsdk.PtrString("This is the slack connector to use for Notification Center testing."),
//		ConnectorConfig: &cxsdk.ComCoralogixapisNotificationCenterConnectorsV1ConnectorConfig{
//			Fields: []cxsdk.ComCoralogixapisNotificationCenterConnectorConfigField{
//				{FieldName: cxsdk.PtrString("integrationId"), Value: cxsdk.PtrString("c6e72871-388b-4776-a7de-8b7d17ed4828")},
//				{FieldName: cxsdk.PtrString("fallbackChannel"), Value: cxsdk.PtrString("luigis-testing-grounds")},
//				{FieldName: cxsdk.PtrString("channel"), Value: cxsdk.PtrString("luigis-testing-grounds")},
//			},
//		},
//		ConfigOverrides: []cxsdk.ComCoralogixapisNotificationCenterConnectorsV1EntityTypeConfigOverrides{
//			{
//				EntityType: cxsdk.PtrString("ALERTS"),
//				Fields: []cxsdk.ComCoralogixapisNotificationCenterTemplatedConnectorConfigField{
//					{
//						FieldName: cxsdk.PtrString("channel"),
//						Template:  cxsdk.PtrString("{{alertDef.priority}}"),
//					},
//				},
//			},
//		},
//	}
//
//	created, _, err := client.DefaultAPI.
//		ConnectorsServiceCreateConnector(context.Background()).
//		Connector(connector).
//		Execute()
//	assertNilAndPrintError(t, err)
//
//	connectorID := created.Connector.Id
//	assert.NotNil(t, connectorID)
//
//	got, _, err := client.DefaultAPI.
//		ConnectorsServiceGetConnector(context.Background(), *connectorID).
//		Execute()
//	assertNilAndPrintError(t, err)
//	assert.Equal(t, "TestSlackConnector", got.Connector.Name)
//
//	testReq := cxsdk.TestingServiceTestExistingConnectorRequest{
//		ConnectorId: connectorID,
//		PayloadType: cxsdk.PtrString("slack_raw"),
//	}
//
//	_, _, err = client.DefaultAPI.
//		TestingServiceTestExistingConnector(context.Background()).
//		TestingServiceTestExistingConnectorRequest(testReq).
//		Execute()
//	assertNilAndPrintError(t, err)
//
//	_, _, err = client.DefaultAPI.
//		ConnectorsServiceDeleteConnector(context.Background(), *connectorID).
//		Execute()
//	assertNilAndPrintError(t, err)
//}
//
//func TestPagerdutyConnector(t *testing.T) {
//	t.Skip("Skipping because OpenAPI Facade is missing notification center APIs")
//
//	client := cxsdk.NewClientBuilder().WithRegionFromEnv().WithAPIKeyFromEnv().Build()
//
//	connector := cxsdk.Connector{
//		Name:        "TestPagerdutyConnector",
//		Type:        "PAGERDUTY",
//		Description: cxsdk.PtrString("This is the PagerDuty connector to use for Notification Center testing."),
//		ConnectorConfig: &cxsdk.ComCoralogixapisNotificationCenterConnectorsV1ConnectorConfig{
//			Fields: []cxsdk.ComCoralogixapisNotificationCenterConnectorConfigField{
//				{FieldName: cxsdk.PtrString("integrationKey"), Value: cxsdk.PtrString("test")},
//			},
//		},
//		ConfigOverrides: []cxsdk.ComCoralogixapisNotificationCenterConnectorsV1EntityTypeConfigOverrides{
//			{
//				EntityType: cxsdk.PtrString("ALERTS"),
//				Fields: []cxsdk.ComCoralogixapisNotificationCenterTemplatedConnectorConfigField{
//					{
//						FieldName: cxsdk.PtrString("integrationKey"),
//						Template:  cxsdk.PtrString("integration_{{alertDef.priority}}"),
//					},
//				},
//			},
//		},
//	}
//
//	created, _, err := client.DefaultAPI.
//		ConnectorsServiceCreateConnector(context.Background()).
//		Connector(connector).
//		Execute()
//	assertNilAndPrintError(t, err)
//
//	connectorID := created.Connector.Id
//	assert.NotNil(t, connectorID)
//
//	got, _, err := client.DefaultAPI.
//		ConnectorsServiceGetConnector(context.Background(), *connectorID).
//		Execute()
//	assertNilAndPrintError(t, err)
//	assert.Equal(t, "TestPagerdutyConnector", got.Connector.Name)
//
//	_, _, err = client.DefaultAPI.
//		ConnectorsServiceDeleteConnector(context.Background(), *connectorID).
//		Execute()
//	assertNilAndPrintError(t, err)
//}
//
//func TestHttpsPreset(t *testing.T) {
//	t.Skip("Skipping because OpenAPI Facade is missing notification center APIs")
//
//	client := cxsdk.NewClientBuilder().
//		WithRegionFromEnv().
//		WithAPIKeyFromEnv().
//		Build()
//
//	name := fmt.Sprintf("TestGoHttpsPreset-%v", uuid.NewString())
//
//	preset := cxsdk.Preset{
//		Name:          name,
//		Description:   cxsdk.PtrString("This is the preset to use for Notification Center testing."),
//		PresetType:    cxsdk.PtrString("CUSTOM"),
//		EntityType:    "ALERTS",
//		ParentId:      cxsdk.PtrString("preset_system_generic_https_alerts_empty"),
//		ConnectorType: cxsdk.PtrString("GENERIC_HTTPS"),
//		ConfigOverrides: []cxsdk.ComCoralogixapisNotificationCenterConfigOverrides{
//			{
//				PayloadType: cxsdk.PtrString("generic_https_default"),
//				MessageConfig: &cxsdk.ComCoralogixapisNotificationCenterMessageConfig{
//					Fields: []cxsdk.ComCoralogixapisNotificationCenterMessageConfigField{
//						{
//							FieldName: "headers",
//							Template:  "{}",
//						},
//						{
//							FieldName: "body",
//							Template:  "{ \"groupingKey\": \"{{alert.groupingKey}}\", \"status\": \"{{alert.status}}\", \"groups\": \"{{alert.groups}}\" }",
//						},
//					},
//				},
//			},
//		},
//	}
//
//	created, _, err := client.DefaultAPI.
//		PresetsServiceCreateCustomPreset(context.Background()).
//		Preset(preset).
//		Execute()
//	assertNilAndPrintError(t, err)
//
//	presetID := created.Preset.Id
//	assert.NotNil(t, presetID)
//
//	got, _, err := client.DefaultAPI.
//		PresetsServiceGetPreset(context.Background(), *presetID).
//		Execute()
//	assertNilAndPrintError(t, err)
//	assert.Equal(t, name, got.Preset.Name)
//
//	_, _, err = client.DefaultAPI.
//		PresetsServiceSetPresetAsDefault(context.Background(), *presetID).
//		Execute()
//	assertNilAndPrintError(t, err)
//
//	defaultPreset, _, err := client.DefaultAPI.
//		PresetsServiceGetDefaultPresetSummary(context.Background()).
//		Execute()
//	assertNilAndPrintError(t, err)
//	assert.NotNil(t, defaultPreset)
//	assert.Equal(t, name, defaultPreset.PresetSummary.Name)
//
//	_, _, err = client.DefaultAPI.
//		PresetsServiceDeleteCustomPreset(context.Background(), *presetID).
//		Execute()
//	assertNilAndPrintError(t, err)
//}
//
//func TestSlackPreset(t *testing.T) {
//	t.Skip("Skipping because OpenAPI Facade is missing notification center APIs")
//
//	client := cxsdk.NewClientBuilder().
//		WithRegionFromEnv().
//		WithAPIKeyFromEnv().
//		Build()
//
//	name := fmt.Sprintf("TestGoSlackPreset-%v", uuid.NewString())
//	presetType := "CUSTOM"
//	parentId := "preset_system_slack_alerts_basic"
//
//	preset := cxsdk.Preset{
//		Name:          name,
//		Description:   cxsdk.PtrString("This is the preset to use for Notification Center testing."),
//		PresetType:    &presetType,
//		EntityType:    "ALERTS",
//		ParentId:      &parentId,
//		ConnectorType: cxsdk.PtrString("SLACK"),
//		ConfigOverrides: []cxsdk.ComCoralogixapisNotificationCenterConfigOverrides{
//			{
//				ConditionType: &cxsdk.ComCoralogixapisNotificationCenterConditionType{
//					ComCoralogixapisNotificationCenterConditionTypeOneOf1: &cxsdk.ComCoralogixapisNotificationCenterConditionTypeOneOf1{
//						MatchEntityTypeAndSubType: &cxsdk.ComCoralogixapisNotificationCenterMatchEntityTypeAndSubTypeCondition{
//							EntitySubType: cxsdk.PtrString("logsImmediateResolved"),
//						},
//					},
//				},
//				MessageConfig: &cxsdk.ComCoralogixapisNotificationCenterMessageConfig{
//					Fields: []cxsdk.ComCoralogixapisNotificationCenterMessageConfigField{
//						{
//							FieldName: "title",
//							Template:  "{{alert.status}} {{alertDef.priority}} - {{alertDef.name}}",
//						},
//						{
//							FieldName: "description",
//							Template:  "{{alert.timestamp |  date(format=\"%Y-%m-%d %H:%M\")}}\nWe've detected a log that matches the query.",
//						},
//					},
//				},
//			},
//		},
//	}
//
//	created, _, err := client.DefaultAPI.
//		PresetsServiceCreateCustomPreset(context.Background()).
//		Preset(preset).
//		Execute()
//	assertNilAndPrintError(t, err)
//
//	presetID := created.Preset.Id
//	assert.NotNil(t, presetID)
//
//	got, _, err := client.DefaultAPI.
//		PresetsServiceGetPreset(context.Background(), *presetID).
//		Execute()
//	assertNilAndPrintError(t, err)
//	assert.Equal(t, name, got.Preset.Name)
//
//	_, _, err = client.DefaultAPI.
//		PresetsServiceSetPresetAsDefault(context.Background(), *presetID).
//		Execute()
//	assertNilAndPrintError(t, err)
//
//	defaultPreset, _, err := client.DefaultAPI.
//		PresetsServiceGetDefaultPresetSummary(context.Background()).
//		Execute()
//	assertNilAndPrintError(t, err)
//	assert.NotNil(t, defaultPreset)
//	assert.Equal(t, name, defaultPreset.PresetSummary.Name)
//
//	_, _, err = client.DefaultAPI.
//		PresetsServiceDeleteCustomPreset(context.Background(), *presetID).
//		Execute()
//	assertNilAndPrintError(t, err)
//}
//
//func TestPagerdutyPreset(t *testing.T) {
//	t.Skip("Skipping because OpenAPI Facade is missing notification center APIs")
//
//	client := cxsdk.NewClientBuilder().
//		WithRegionFromEnv().
//		WithAPIKeyFromEnv().
//		Build()
//
//	name := fmt.Sprintf("TestPagerDutyPreset-%v", uuid.NewString())
//
//	preset := cxsdk.Preset{
//		Name:          name,
//		Description:   cxsdk.PtrString("This is the preset to use for Notification Center testing."),
//		PresetType:    cxsdk.PtrString("CUSTOM"),
//		EntityType:    "ALERTS",
//		ParentId:      cxsdk.PtrString("preset_system_pagerduty_alerts_basic"),
//		ConnectorType: cxsdk.PtrString("PAGERDUTY"),
//		ConfigOverrides: []cxsdk.ComCoralogixapisNotificationCenterConfigOverrides{
//			{
//				ConditionType: &cxsdk.ComCoralogixapisNotificationCenterConditionType{
//					ComCoralogixapisNotificationCenterConditionTypeOneOf1: &cxsdk.ComCoralogixapisNotificationCenterConditionTypeOneOf1{
//						MatchEntityTypeAndSubType: &cxsdk.ComCoralogixapisNotificationCenterMatchEntityTypeAndSubTypeCondition{
//							EntitySubType: cxsdk.PtrString("logsImmediateResolved"),
//						},
//					},
//				},
//				MessageConfig: &cxsdk.ComCoralogixapisNotificationCenterMessageConfig{
//					Fields: []cxsdk.ComCoralogixapisNotificationCenterMessageConfigField{
//						{
//							FieldName: "summary",
//							Template:  "{{ alertDef.name }}",
//						},
//					},
//				},
//			},
//		},
//	}
//
//	created, _, err := client.DefaultAPI.
//		PresetsServiceCreateCustomPreset(context.Background()).
//		Preset(preset).
//		Execute()
//	assertNilAndPrintError(t, err)
//
//	presetID := created.Preset.Id
//	assert.NotNil(t, presetID)
//
//	got, _, err := client.DefaultAPI.
//		PresetsServiceGetPreset(context.Background(), *presetID).
//		Execute()
//	assertNilAndPrintError(t, err)
//	assert.Equal(t, name, got.Preset.Name)
//
//	_, _, err = client.DefaultAPI.
//		PresetsServiceSetPresetAsDefault(context.Background(), *presetID).
//		Execute()
//	assertNilAndPrintError(t, err)
//
//	defaultPreset, _, err := client.DefaultAPI.
//		PresetsServiceGetDefaultPresetSummary(context.Background()).
//		Execute()
//	assertNilAndPrintError(t, err)
//	assert.NotNil(t, defaultPreset)
//	assert.Equal(t, name, defaultPreset.PresetSummary.Name)
//
//	_, _, err = client.DefaultAPI.
//		PresetsServiceDeleteCustomPreset(context.Background(), *presetID).
//		Execute()
//	assertNilAndPrintError(t, err)
//}
//
//// TODO: Add Alert with routing
//func TestGlobalRouter(t *testing.T) {
//	t.Skip("Skipping because OpenAPI Facade is missing notification center APIs")
//
//	client := cxsdk.NewClientBuilder().
//		WithRegionFromEnv().
//		WithAPIKeyFromEnv().
//		Build()
//
//	_, _, err := client.DefaultAPI.
//		GlobalRoutersServiceListGlobalRouters(context.Background()).
//		Execute()
//	assertNilAndPrintError(t, err)
//
//	routerId := "router_default"
//	name := fmt.Sprintf("TestConnector-%v", uuid.NewString())
//
//	// --- Create HTTPS connector
//	connector := cxsdk.Connector{
//		Type:        "GENERIC_HTTPS",
//		Name:        name,
//		Description: cxsdk.PtrString("This is the connector to use for Notification Center testing."),
//		ConnectorConfig: &cxsdk.ComCoralogixapisNotificationCenterConnectorsV1ConnectorConfig{
//			Fields: []cxsdk.ComCoralogixapisNotificationCenterConnectorConfigField{
//				{FieldName: cxsdk.PtrString("url"), Value: cxsdk.PtrString("https://httpbun.org/post")},
//				{FieldName: cxsdk.PtrString("method"), Value: cxsdk.PtrString("post")},
//			},
//		},
//		ConfigOverrides: []cxsdk.ComCoralogixapisNotificationCenterConnectorsV1EntityTypeConfigOverrides{
//			{
//				EntityType: cxsdk.PtrString("ALERTS"),
//				Fields: []cxsdk.ComCoralogixapisNotificationCenterTemplatedConnectorConfigField{
//					{
//						FieldName: cxsdk.PtrString("additionalBodyFields"),
//						Template:  cxsdk.PtrString("{\"priority\": \"{{alertDef.priority}}\"}"),
//					},
//				},
//			},
//		},
//	}
//
//	preset := cxsdk.Preset{
//		Name:          fmt.Sprintf("TestGoHttpsPreset-%v", uuid.NewString()),
//		EntityType:    "ALERTS",
//		ConnectorType: cxsdk.PtrString("GENERIC_HTTPS"),
//		ConfigOverrides: []cxsdk.ComCoralogixapisNotificationCenterConfigOverrides{
//			{
//				PayloadType: cxsdk.PtrString("generic_https_default"),
//				MessageConfig: &cxsdk.ComCoralogixapisNotificationCenterMessageConfig{
//					Fields: []cxsdk.ComCoralogixapisNotificationCenterMessageConfigField{
//						{FieldName: "url", Template: "https://httpbun.org/post"},
//						{FieldName: "method", Template: "POST"},
//						{FieldName: "body", Template: "{\"priority\": \"{{alertDef.priority}}\"}"},
//					},
//				},
//			},
//		},
//	}
//
//	createdPreset, _, err := client.DefaultAPI.
//		PresetsServiceCreateCustomPreset(context.Background()).
//		Preset(preset).
//		Execute()
//	assertNilAndPrintError(t, err)
//	presetId := createdPreset.Preset.Id
//
//	createdConnector, _, err := client.DefaultAPI.
//		ConnectorsServiceCreateConnector(context.Background()).
//		Connector(connector).
//		Execute()
//	assertNilAndPrintError(t, err)
//	connectorId := createdConnector.Connector.Id
//
//	router := cxsdk.GlobalRouter{
//		Id:          &routerId,
//		Name:        "global router",
//		Description: cxsdk.PtrString("global router example"),
//		Rules: []cxsdk.ComCoralogixapisNotificationCenterRoutingRoutingRule{
//			{
//				Name:      cxsdk.PtrString("TestRoutingRule"),
//				Condition: "alertDef.priority == \"P1\"",
//				Targets: []cxsdk.ComCoralogixapisNotificationCenterRoutingRoutingTarget{
//					{
//						ConnectorId: *connectorId,
//						PresetId:    presetId,
//					},
//				},
//			},
//		},
//	}
//	createdRouter, _, err := client.DefaultAPI.
//		GlobalRoutersServiceCreateGlobalRouter(context.Background()).
//		GlobalRouter(router).
//		Execute()
//	assertNilAndPrintError(t, err)
//
//	gotRouter, _, err := client.DefaultAPI.
//		GlobalRoutersServiceGetGlobalRouter(context.Background(), *createdRouter.Router.Id).
//		Execute()
//	assertNilAndPrintError(t, err)
//
//	assert.Equal(t, "global router", gotRouter.Router.Name)
//
//	_, _, err = client.DefaultAPI.
//		TestingServiceTestDestination(context.Background()).
//		TestingServiceTestDestinationRequest(cxsdk.TestingServiceTestDestinationRequest{
//			EntityType:  cxsdk.PtrString("ALERTS"),
//			ConnectorId: connectorId,
//			PresetId:    presetId,
//			PayloadType: cxsdk.PtrString("generic_https_default"),
//		}).
//		Execute()
//	assertNilAndPrintError(t, err)
//
//	_, _, err = client.DefaultAPI.
//		GlobalRoutersServiceDeleteGlobalRouter(context.Background(), *createdRouter.Router.Id).
//		Execute()
//	assertNilAndPrintError(t, err)
//
//	_, _, err = client.DefaultAPI.
//		ConnectorsServiceDeleteConnector(context.Background(), *connectorId).
//		Execute()
//	assertNilAndPrintError(t, err)
//
//	_, _, err = client.DefaultAPI.
//		PresetsServiceDeleteCustomPreset(context.Background(), *presetId).
//		Execute()
//	assertNilAndPrintError(t, err)
//}
