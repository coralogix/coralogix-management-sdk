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

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	alerts "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/alert_definitions_service"
	connectors "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/connectors_service"
	globalrouters "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/global_routers_service"
	presets "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/presets_service"
)

func TestHttpsConnector(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().WithHTTPLogging().Build()
	client := cxsdk.NewConnectorsClient(cfg)

	name := fmt.Sprintf("TestConnector-%v", uuid.NewString())

	connector := getHttpsConnector(name)

	created, httpResp, err := client.
		ConnectorsServiceCreateConnector(context.Background()).
		CreateConnectorRequest(*connector).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	connectorID := created.Connector.Id
	require.NotNil(t, connectorID)

	got, httpResp, err := client.
		ConnectorsServiceGetConnector(context.Background(), *connectorID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, name, *got.Connector.Name)

	_, httpResp, err = client.
		ConnectorsServiceDeleteConnector(context.Background(), *connectorID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
}

func TestSlackConnector(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().WithHTTPLogging().Build()
	client := cxsdk.NewConnectorsClient(cfg)

	connector := connectors.Connector{
		Name:        connectors.PtrString("TestSlackConnector"),
		Type:        connectors.CONNECTORTYPE_SLACK.Ptr(),
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

	createConnectorRequest := connectors.CreateConnectorRequest{
		Connector: &connector,
	}

	created, httpResp, err := client.
		ConnectorsServiceCreateConnector(context.Background()).
		CreateConnectorRequest(createConnectorRequest).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	connectorID := created.Connector.Id
	require.NotNil(t, connectorID)

	got, httpResp, err := client.
		ConnectorsServiceGetConnector(context.Background(), *connectorID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, "TestSlackConnector", *got.Connector.Name)

	_, httpResp, err = client.
		ConnectorsServiceDeleteConnector(context.Background(), *connectorID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
}

func TestPagerdutyConnector(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().WithHTTPLogging().Build()
	client := cxsdk.NewConnectorsClient(cfg)

	connector := connectors.Connector{
		Name:        connectors.PtrString("TestPagerdutyConnector"),
		Type:        connectors.CONNECTORTYPE_PAGERDUTY.Ptr(),
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

	createConnectorRequest := connectors.CreateConnectorRequest{
		Connector: &connector,
	}

	created, httpResp, err := client.
		ConnectorsServiceCreateConnector(context.Background()).
		CreateConnectorRequest(createConnectorRequest).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	connectorID := created.Connector.Id
	require.NotNil(t, connectorID)

	got, httpResp, err := client.
		ConnectorsServiceGetConnector(context.Background(), *connectorID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, "TestPagerdutyConnector", *got.Connector.Name)

	_, httpResp, err = client.
		ConnectorsServiceDeleteConnector(context.Background(), *connectorID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
}

func TestHttpsPreset(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().WithHTTPLogging().Build()
	client := cxsdk.NewPresetsClient(cfg)

	name := fmt.Sprintf("TestGoHttpsPreset-%v", uuid.NewString())

	createPresetRequest := getHttpsPreset(name)

	created, httpResp, err := client.
		PresetsServiceCreateCustomPreset(context.Background()).
		CreateCustomPresetRequest(*createPresetRequest).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	presetID := created.Preset.Id
	require.NotNil(t, presetID)

	got, httpResp, err := client.
		PresetsServiceGetPreset(context.Background(), *presetID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, name, *got.Preset.Name)

	_, httpResp, err = client.
		PresetsServiceSetPresetAsDefault(context.Background(), *presetID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	defaultPreset, httpResp, err := client.
		PresetsServiceGetDefaultPresetSummary(context.Background()).
		EntityType("ALERTS").
		ConnectorType("GENERIC_HTTPS").
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotNil(t, defaultPreset)
	require.Equal(t, name, *defaultPreset.PresetSummary.Name)

	_, httpResp, err = client.
		PresetsServiceDeleteCustomPreset(context.Background(), *presetID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
}

func TestSlackPreset(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().WithHTTPLogging().Build()
	client := cxsdk.NewPresetsClient(cfg)

	name := fmt.Sprintf("TestGoSlackPreset-%v", uuid.NewString())

	preset := presets.Preset{
		Name:          &name,
		Description:   presets.PtrString("This is the preset to use for Notification Center testing."),
		PresetType:    presets.PRESETTYPE_CUSTOM.Ptr(),
		EntityType:    presets.NOTIFICATIONCENTERENTITYTYPE_ALERTS.Ptr(),
		ParentId:      presets.PtrString("preset_system_slack_alerts_basic"),
		ConnectorType: presets.CONNECTORTYPE_SLACK.Ptr(),
		ConfigOverrides: []presets.ConfigOverrides{
			{
				ConditionType: &presets.NotificationCenterConditionType{
					NotificationCenterConditionTypeMatchEntityTypeAndSubType: &presets.NotificationCenterConditionTypeMatchEntityTypeAndSubType{
						MatchEntityTypeAndSubType: &presets.MatchEntityTypeAndSubTypeCondition{
							EntitySubType: presets.PtrString("logsImmediateResolved"),
						},
					},
				},
				MessageConfig: &presets.MessageConfig{
					Fields: []presets.NotificationCenterMessageConfigField{
						{
							FieldName: presets.PtrString("title"),
							Template:  presets.PtrString("{{alert.status}} {{alertDef.priority}} - {{alertDef.name}}"),
						},
						{
							FieldName: presets.PtrString("description"),
							Template:  presets.PtrString("{{alert.timestamp |  date(format=\"%Y-%m-%d %H:%M\")}}\nWe've detected a log that matches the query."),
						},
					},
				},
			},
		},
	}

	createPresetRequest := presets.CreateCustomPresetRequest{
		Preset: &preset,
	}

	created, httpResp, err := client.
		PresetsServiceCreateCustomPreset(context.Background()).
		CreateCustomPresetRequest(createPresetRequest).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	presetID := created.Preset.Id
	require.NotNil(t, presetID)

	got, httpResp, err := client.
		PresetsServiceGetPreset(context.Background(), *presetID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, name, *got.Preset.Name)

	_, httpResp, err = client.
		PresetsServiceSetPresetAsDefault(context.Background(), *presetID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	defaultPreset, httpResp, err := client.
		PresetsServiceGetDefaultPresetSummary(context.Background()).
		ConnectorType("SLACK").
		EntityType("ALERTS").
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotNil(t, defaultPreset)
	require.Equal(t, name, *defaultPreset.PresetSummary.Name)

	_, httpResp, err = client.
		PresetsServiceDeleteCustomPreset(context.Background(), *presetID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
}

func TestPagerdutyPreset(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().WithHTTPLogging().Build()
	client := cxsdk.NewPresetsClient(cfg)

	name := fmt.Sprintf("TestPagerDutyPreset-%v", uuid.NewString())

	preset := presets.Preset{
		Name:          &name,
		Description:   presets.PtrString("This is the preset to use for Notification Center testing."),
		PresetType:    presets.PRESETTYPE_CUSTOM.Ptr(),
		EntityType:    presets.NOTIFICATIONCENTERENTITYTYPE_ALERTS.Ptr(),
		ParentId:      presets.PtrString("preset_system_pagerduty_alerts_basic"),
		ConnectorType: presets.CONNECTORTYPE_PAGERDUTY.Ptr(),
		ConfigOverrides: []presets.ConfigOverrides{
			{
				ConditionType: &presets.NotificationCenterConditionType{
					NotificationCenterConditionTypeMatchEntityTypeAndSubType: &presets.NotificationCenterConditionTypeMatchEntityTypeAndSubType{
						MatchEntityTypeAndSubType: &presets.MatchEntityTypeAndSubTypeCondition{
							EntitySubType: presets.PtrString("logsImmediateTriggered"),
						},
					},
				},
				MessageConfig: &presets.MessageConfig{
					Fields: []presets.NotificationCenterMessageConfigField{
						{
							FieldName: presets.PtrString("summary"),
							Template:  presets.PtrString("{{ alertDef.name }}"),
						},
					},
				},
			},
		},
	}

	createPresetRequest := presets.CreateCustomPresetRequest{
		Preset: &preset,
	}

	created, httpResp, err := client.
		PresetsServiceCreateCustomPreset(context.Background()).
		CreateCustomPresetRequest(createPresetRequest).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	presetID := created.Preset.Id
	require.NotNil(t, presetID)

	got, httpResp, err := client.
		PresetsServiceGetPreset(context.Background(), *presetID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, name, *got.Preset.Name)

	_, httpResp, err = client.
		PresetsServiceSetPresetAsDefault(context.Background(), *presetID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	defaultPreset, httpResp, err := client.
		PresetsServiceGetDefaultPresetSummary(context.Background()).
		ConnectorType("PAGERDUTY").
		EntityType("ALERTS").
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotNil(t, defaultPreset)
	require.Equal(t, name, *defaultPreset.PresetSummary.Name)

	_, httpResp, err = client.
		PresetsServiceDeleteCustomPreset(context.Background(), *presetID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
}

func TestGlobalRouter(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().WithHTTPLogging().Build()
	clientSet := cxsdk.NewClientSet(cfg)

	routersClient := clientSet.GlobalRouters()
	connectorsClient := clientSet.Connectors()
	presetsClient := clientSet.Presets()
	alertsClient := clientSet.Alerts()

	createdConnector, httpResp, err := connectorsClient.
		ConnectorsServiceCreateConnector(context.Background()).
		CreateConnectorRequest(*getHttpsConnector(fmt.Sprintf("TestConnector-%v", uuid.NewString()))).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	connectorId := createdConnector.Connector.Id

	createdPreset, httpResp, err := presetsClient.
		PresetsServiceCreateCustomPreset(context.Background()).
		CreateCustomPresetRequest(*getHttpsPreset(fmt.Sprintf("TestGoHttpsPreset-%v", uuid.NewString()))).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	presetId := createdPreset.Preset.Id

	router := globalrouters.GlobalRouter{
		Name: globalrouters.PtrString("global router" + uuid.NewString()),
		RoutingLabels: &globalrouters.RoutingLabels{
			Environment: globalrouters.PtrString(uuid.NewString()),
			Service:     globalrouters.PtrString(uuid.NewString()),
			Team:        globalrouters.PtrString(uuid.NewString()),
		},
		Description: globalrouters.PtrString("global router example"),
		Rules: []globalrouters.RoutingRule{
			{
				Name:      globalrouters.PtrString("TestRoutingRule"),
				Condition: globalrouters.PtrString("alertDef.priority == \"P1\""),
				Targets: []globalrouters.RoutingTarget{
					{
						ConnectorId: connectorId,
						PresetId:    presetId,
					},
				},
			},
		},
	}

	createRouterRequest := globalrouters.CreateGlobalRouterRequest{Router: &router}
	createdRouter, httpResp, err := routersClient.
		GlobalRoutersServiceCreateGlobalRouter(context.Background()).
		CreateGlobalRouterRequest(createRouterRequest).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	router.Id = createdRouter.Router.Id
	router.Name = globalrouters.PtrString("global router updated" + uuid.NewString())
	replaceRouterRequest := globalrouters.ReplaceGlobalRouterRequest{Router: &router}
	replacedRouter, httpResp, err := routersClient.
		GlobalRoutersServiceReplaceGlobalRouter(context.Background()).
		ReplaceGlobalRouterRequest(replaceRouterRequest).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	createAlertReq := alerts.CreateAlertDefinitionRequest{
		AlertDefProperties: &alerts.AlertDefProperties{
			AlertDefPropertiesLogsRatioThreshold: CreateLogsRatioAlert(),
		},
	}
	createAlertReq.AlertDefProperties.
		AlertDefPropertiesLogsRatioThreshold.
		NotificationGroup.
		Router = &alerts.NotificationRouter{
		Id: replacedRouter.Router.Id,
	}
	_, httpResp, err = alertsClient.
		AlertDefsServiceCreateAlertDef(context.Background()).
		CreateAlertDefinitionRequest(createAlertReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	got, httpResp, err := routersClient.
		GlobalRoutersServiceGetGlobalRouter(context.Background(), *replacedRouter.Router.Id).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, got.Router.Name, replaceRouterRequest.Router.Name)

	list, httpResp, err := routersClient.
		GlobalRoutersServiceListGlobalRouters(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	found := false
	for _, r := range list.Routers {
		if *r.Id == *replacedRouter.Router.Id {
			found = true
			break
		}
	}
	require.True(t, found, "created router not found in list")

	_, httpResp, err = routersClient.
		GlobalRoutersServiceDeleteGlobalRouter(context.Background(), *replacedRouter.Router.Id).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	_, httpResp, err = connectorsClient.
		ConnectorsServiceDeleteConnector(context.Background(), *connectorId).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	_, httpResp, err = presetsClient.
		PresetsServiceDeleteCustomPreset(context.Background(), *presetId).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
}

func getHttpsConnector(name string) *connectors.CreateConnectorRequest {
	return &connectors.CreateConnectorRequest{
		Connector: &connectors.Connector{
			Name:        &name,
			Type:        connectors.CONNECTORTYPE_GENERIC_HTTPS.Ptr(),
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
		},
	}

}

func getHttpsPreset(name string) *presets.CreateCustomPresetRequest {
	return &presets.CreateCustomPresetRequest{
		Preset: &presets.Preset{
			Name:          &name,
			Description:   presets.PtrString("This is the preset to use for Notification Center testing."),
			PresetType:    presets.PRESETTYPE_CUSTOM.Ptr(),
			EntityType:    presets.NOTIFICATIONCENTERENTITYTYPE_ALERTS.Ptr(),
			ParentId:      presets.PtrString("preset_system_generic_https_alerts_empty"),
			ConnectorType: presets.CONNECTORTYPE_GENERIC_HTTPS.Ptr(),
			ConfigOverrides: []presets.ConfigOverrides{
				{
					PayloadType: presets.PtrString("generic_https_default"),
					MessageConfig: &presets.MessageConfig{
						Fields: []presets.NotificationCenterMessageConfigField{
							{
								FieldName: presets.PtrString("headers"),
								Template:  presets.PtrString("{}"),
							},
							{
								FieldName: presets.PtrString("body"),
								Template:  presets.PtrString("{ \"groupingKey\": \"{{alert.groupingKey}}\", \"status\": \"{{alert.status}}\", \"groups\": \"{{alert.groups}}\" }"),
							},
						},
					},
					ConditionType: &presets.NotificationCenterConditionType{
						NotificationCenterConditionTypeMatchEntityTypeAndSubType: &presets.NotificationCenterConditionTypeMatchEntityTypeAndSubType{
							MatchEntityTypeAndSubType: &presets.MatchEntityTypeAndSubTypeCondition{
								EntitySubType: presets.PtrString("logsImmediateResolved"),
							},
						},
					},
				},
			},
		},
	}
}
