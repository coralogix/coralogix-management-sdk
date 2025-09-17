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

import (
	"context"
	"fmt"
	"testing"

	globalrouters "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/global_routers_service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	connectors "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/connectors_service"
	presets "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/presets_service"
)

func TestHttpsConnector(t *testing.T) {
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewConnectorsClient(cpc)

	name := fmt.Sprintf("TestConnector-%v", uuid.NewString())

	connector := getHttpsConnector(name)

	created, httpResp, err := client.
		ConnectorsServiceCreateConnector(context.Background()).
		Connector1(*connector).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))

	connectorID := created.Connector.Id
	assert.NotNil(t, connectorID)

	got, httpResp, err := client.
		ConnectorsServiceGetConnector(context.Background(), *connectorID).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.Equal(t, name, got.Connector.Name)

	_, httpResp, err = client.
		ConnectorsServiceDeleteConnector(context.Background(), *connectorID).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
}

func TestSlackConnector(t *testing.T) {
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewConnectorsClient(cpc)

	connector := connectors.Connector1{
		Name:        "TestSlackConnector",
		Type:        "SLACK",
		Description: connectors.PtrString("This is the slack connector to use for Notification Center testing."),
		ConnectorConfig: &connectors.ConnectorConfig{
			Fields: []connectors.NotificationCenterConnectorConfigField{
				{FieldName: connectors.PtrString("integrationId"), Value: connectors.PtrString("60d87305-1110-4a0a-b388-85fb769892ee")},
				{FieldName: connectors.PtrString("fallbackChannel"), Value: connectors.PtrString("luigis-testing-grounds")},
				{FieldName: connectors.PtrString("channel"), Value: connectors.PtrString("luigis-testing-grounds")},
			},
		},
		ConfigOverrides: []connectors.EntityTypeConfigOverrides{
			{
				EntityType: connectors.NOTIFICATIONCENTERENTITYTYPE_ALERTS.Ptr(),
				Fields: []connectors.TemplatedConnectorConfigField{
					{
						FieldName: connectors.PtrString("channel"),
						Template:  connectors.PtrString("{{alertDef.priority}}"),
					},
				},
			},
		},
	}

	created, httpResp, err := client.
		ConnectorsServiceCreateConnector(context.Background()).
		Connector1(connector).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))

	connectorID := created.Connector.Id
	assert.NotNil(t, connectorID)

	got, httpResp, err := client.
		ConnectorsServiceGetConnector(context.Background(), *connectorID).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.Equal(t, "TestSlackConnector", got.Connector.Name)

	_, httpResp, err = client.
		ConnectorsServiceDeleteConnector(context.Background(), *connectorID).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
}

func TestPagerdutyConnector(t *testing.T) {
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewConnectorsClient(cpc)

	connector := connectors.Connector1{
		Name:        "TestPagerdutyConnector",
		Type:        "PAGERDUTY",
		Description: connectors.PtrString("This is the PagerDuty connector to use for Notification Center testing."),
		ConnectorConfig: &connectors.ConnectorConfig{
			Fields: []connectors.NotificationCenterConnectorConfigField{
				{FieldName: connectors.PtrString("integrationKey"), Value: connectors.PtrString("test")},
			},
		},
		ConfigOverrides: []connectors.EntityTypeConfigOverrides{
			{
				EntityType: connectors.NOTIFICATIONCENTERENTITYTYPE_ALERTS.Ptr(),
				Fields: []connectors.TemplatedConnectorConfigField{
					{
						FieldName: connectors.PtrString("integrationKey"),
						Template:  connectors.PtrString("integration_{{alertDef.priority}}"),
					},
				},
			},
		},
	}

	created, httpResp, err := client.
		ConnectorsServiceCreateConnector(context.Background()).
		Connector1(connector).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))

	connectorID := created.Connector.Id
	assert.NotNil(t, connectorID)

	got, httpResp, err := client.
		ConnectorsServiceGetConnector(context.Background(), *connectorID).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.Equal(t, "TestPagerdutyConnector", got.Connector.Name)

	_, httpResp, err = client.
		ConnectorsServiceDeleteConnector(context.Background(), *connectorID).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
}

func TestHttpsPreset(t *testing.T) {
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewPresetsClient(cpc)

	name := fmt.Sprintf("TestGoHttpsPreset-%v", uuid.NewString())

	preset := getHttpsPreset(name)

	created, httpResp, err := client.
		PresetsServiceCreateCustomPreset(context.Background()).
		Preset1(*preset).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))

	presetID := created.Preset.Id
	assert.NotNil(t, presetID)

	got, httpResp, err := client.
		PresetsServiceGetPreset(context.Background(), *presetID).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.Equal(t, name, got.Preset.Name)

	_, httpResp, err = client.
		PresetsServiceSetPresetAsDefault(context.Background(), *presetID).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))

	defaultPreset, httpResp, err := client.
		PresetsServiceGetDefaultPresetSummary(context.Background()).
		EntityType("ALERTS").
		ConnectorType("GENERIC_HTTPS").
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.NotNil(t, defaultPreset)
	assert.Equal(t, name, *defaultPreset.PresetSummary.Name)

	_, httpResp, err = client.
		PresetsServiceDeleteCustomPreset(context.Background(), *presetID).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
}

func TestSlackPreset(t *testing.T) {
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewPresetsClient(cpc)

	name := fmt.Sprintf("TestGoSlackPreset-%v", uuid.NewString())

	preset := presets.Preset1{
		Name:          name,
		Description:   presets.PtrString("This is the preset to use for Notification Center testing."),
		PresetType:    presets.PRESETTYPE_CUSTOM.Ptr(),
		EntityType:    "ALERTS",
		ParentId:      presets.PtrString("preset_system_slack_alerts_basic"),
		ConnectorType: presets.CONNECTORTYPE_SLACK.Ptr(),
		ConfigOverrides: []presets.ConfigOverrides{
			{
				ConditionType: &presets.NotificationCenterConditionType{
					NotificationCenterConditionTypeOneOf1: &presets.NotificationCenterConditionTypeOneOf1{
						MatchEntityTypeAndSubType: &presets.MatchEntityTypeAndSubTypeCondition{
							EntitySubType: presets.PtrString("logsImmediateResolved"),
						},
					},
				},
				MessageConfig: &presets.MessageConfig{
					Fields: []presets.NotificationCenterMessageConfigField{
						{
							FieldName: "title",
							Template:  "{{alert.status}} {{alertDef.priority}} - {{alertDef.name}}",
						},
						{
							FieldName: "description",
							Template:  "{{alert.timestamp |  date(format=\"%Y-%m-%d %H:%M\")}}\nWe've detected a log that matches the query.",
						},
					},
				},
			},
		},
	}

	created, httpResp, err := client.
		PresetsServiceCreateCustomPreset(context.Background()).
		Preset1(preset).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))

	presetID := created.Preset.Id
	assert.NotNil(t, presetID)

	got, httpResp, err := client.
		PresetsServiceGetPreset(context.Background(), *presetID).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.Equal(t, name, got.Preset.Name)

	_, httpResp, err = client.
		PresetsServiceSetPresetAsDefault(context.Background(), *presetID).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))

	defaultPreset, httpResp, err := client.
		PresetsServiceGetDefaultPresetSummary(context.Background()).
		ConnectorType("SLACK").
		EntityType("ALERTS").
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.NotNil(t, defaultPreset)
	assert.Equal(t, name, *defaultPreset.PresetSummary.Name)

	_, httpResp, err = client.
		PresetsServiceDeleteCustomPreset(context.Background(), *presetID).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
}

func TestPagerdutyPreset(t *testing.T) {
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewPresetsClient(cpc)

	name := fmt.Sprintf("TestPagerDutyPreset-%v", uuid.NewString())

	preset := presets.Preset1{
		Name:          name,
		Description:   presets.PtrString("This is the preset to use for Notification Center testing."),
		PresetType:    presets.PRESETTYPE_CUSTOM.Ptr(),
		EntityType:    presets.NOTIFICATIONCENTERENTITYTYPE_ALERTS,
		ParentId:      presets.PtrString("preset_system_pagerduty_alerts_basic"),
		ConnectorType: presets.CONNECTORTYPE_PAGERDUTY.Ptr(),
		ConfigOverrides: []presets.ConfigOverrides{
			{
				ConditionType: &presets.NotificationCenterConditionType{
					NotificationCenterConditionTypeOneOf1: &presets.NotificationCenterConditionTypeOneOf1{
						MatchEntityTypeAndSubType: &presets.MatchEntityTypeAndSubTypeCondition{
							EntitySubType: presets.PtrString("logsImmediateTriggered"),
						},
					},
				},
				MessageConfig: &presets.MessageConfig{
					Fields: []presets.NotificationCenterMessageConfigField{
						{
							FieldName: "summary",
							Template:  "{{ alertDef.name }}",
						},
					},
				},
			},
		},
	}

	created, httpResp, err := client.
		PresetsServiceCreateCustomPreset(context.Background()).
		Preset1(preset).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))

	presetID := created.Preset.Id
	assert.NotNil(t, presetID)

	got, httpResp, err := client.
		PresetsServiceGetPreset(context.Background(), *presetID).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.Equal(t, name, got.Preset.Name)

	_, httpResp, err = client.
		PresetsServiceSetPresetAsDefault(context.Background(), *presetID).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))

	defaultPreset, httpResp, err := client.
		PresetsServiceGetDefaultPresetSummary(context.Background()).
		ConnectorType("PAGERDUTY").
		EntityType("ALERTS").
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.NotNil(t, defaultPreset)
	assert.Equal(t, name, *defaultPreset.PresetSummary.Name)

	_, httpResp, err = client.
		PresetsServiceDeleteCustomPreset(context.Background(), *presetID).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
}

// TODO: Add Alert with routing
func TestGlobalRouter(t *testing.T) {
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)

	routersClient := cxsdk.NewGlobalRoutersClient(cpc)
	connectorsClient := cxsdk.NewConnectorsClient(cpc)
	presetsClient := cxsdk.NewPresetsClient(cpc)

	//_, _, err := routersClient.
	//	GlobalRoutersServiceListGlobalRouters(context.Background()).EntityType("ALERTS").
	//	Execute()
	//assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))

	routerId := "router_default"
	connector := getHttpsConnector(fmt.Sprintf("TestConnector-%v", uuid.NewString()))
	preset := getHttpsPreset(fmt.Sprintf("TestGoHttpsPreset-%v", uuid.NewString()))

	createdConnector, httpResp, err := connectorsClient.
		ConnectorsServiceCreateConnector(context.Background()).
		Connector1(*connector).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	connectorId := createdConnector.Connector.Id

	createdPreset, httpResp, err := presetsClient.
		PresetsServiceCreateCustomPreset(context.Background()).
		Preset1(*preset).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	presetId := createdPreset.Preset.Id

	router := globalrouters.GlobalRouter{
		Id:          &routerId,
		Name:        "global router",
		EntityType:  globalrouters.NOTIFICATIONCENTERENTITYTYPE_ALERTS,
		Description: globalrouters.PtrString("global router example"),
		Rules: []globalrouters.RoutingRule{
			{
				Name:      globalrouters.PtrString("TestRoutingRule"),
				Condition: "alertDef.priority == \"P1\"",
				Targets: []globalrouters.RoutingTarget{
					{
						ConnectorId: *connectorId,
						PresetId:    presetId,
					},
				},
			},
		},
	}
	replacedRouter, httpResp, err := routersClient.
		GlobalRoutersServiceReplaceGlobalRouter(context.Background()).
		GlobalRouter(router).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))

	//gotRouter, _, err := routersClient.
	//	GlobalRoutersServiceGetGlobalRouter(context.Background(), *replacedRouter.Router.Id).
	//	Execute()
	//assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	//
	//assert.Equal(t, "global router", gotRouter.Router.Name)

	_, httpResp, err = routersClient.
		GlobalRoutersServiceDeleteGlobalRouter(context.Background(), *replacedRouter.Router.Id).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))

	_, httpResp, err = connectorsClient.
		ConnectorsServiceDeleteConnector(context.Background(), *connectorId).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))

	_, httpResp, err = presetsClient.
		PresetsServiceDeleteCustomPreset(context.Background(), *presetId).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
}

func getHttpsConnector(name string) *connectors.Connector1 {
	return &connectors.Connector1{
		Name:        name,
		Type:        connectors.CONNECTORTYPE_GENERIC_HTTPS,
		Description: connectors.PtrString("This is the connector to use for Notification Center testing."),
		ConnectorConfig: &connectors.ConnectorConfig{
			Fields: []connectors.NotificationCenterConnectorConfigField{
				{FieldName: connectors.PtrString("url"), Value: connectors.PtrString("https://httpbun.org/post")},
				{FieldName: connectors.PtrString("method"), Value: connectors.PtrString("post")},
			},
		},
		ConfigOverrides: []connectors.EntityTypeConfigOverrides{
			{
				EntityType: connectors.NOTIFICATIONCENTERENTITYTYPE_ALERTS.Ptr(),
				Fields: []connectors.TemplatedConnectorConfigField{
					{
						FieldName: connectors.PtrString("additionalBodyFields"),
						Template:  connectors.PtrString("{\"priority\": \"{{alertDef.priority}}\"}"),
					},
				},
			},
		},
	}
}

func getHttpsPreset(name string) *presets.Preset1 {
	return &presets.Preset1{
		Name:          name,
		Description:   presets.PtrString("This is the preset to use for Notification Center testing."),
		PresetType:    presets.PRESETTYPE_CUSTOM.Ptr(),
		EntityType:    "ALERTS",
		ParentId:      presets.PtrString("preset_system_generic_https_alerts_empty"),
		ConnectorType: presets.CONNECTORTYPE_GENERIC_HTTPS.Ptr(),
		ConfigOverrides: []presets.ConfigOverrides{
			{
				PayloadType: presets.PtrString("generic_https_default"),
				MessageConfig: &presets.MessageConfig{
					Fields: []presets.NotificationCenterMessageConfigField{
						{
							FieldName: "headers",
							Template:  "{}",
						},
						{
							FieldName: "body",
							Template:  "{ \"groupingKey\": \"{{alert.groupingKey}}\", \"status\": \"{{alert.status}}\", \"groups\": \"{{alert.groups}}\" }",
						},
					},
				},
				ConditionType: &presets.NotificationCenterConditionType{
					NotificationCenterConditionTypeOneOf1: &presets.NotificationCenterConditionTypeOneOf1{
						MatchEntityTypeAndSubType: &presets.MatchEntityTypeAndSubTypeCondition{
							EntitySubType: presets.PtrString("logsImmediateResolved"),
						},
					},
				},
			},
		},
	}
}
