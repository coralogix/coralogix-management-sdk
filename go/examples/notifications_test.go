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

	cxsdk "github.com/coralogix/coralogix-management-sdk/go"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestHttpsConnector(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	name := fmt.Sprintf("TestConnector-%v", uuid.NewString())
	connectorRaw := cxsdk.Connector{
		Type:        cxsdk.ConnectorTypeGenericHTTPS,
		Name:        name,
		Description: "This is the connector to use for Notification Center testing.",
		ConnectorConfig: &cxsdk.ConnectorConfig{
			Fields: []*cxsdk.ConnectorConfigField{
				{FieldName: "url", Value: "https://httpbun.org/post"},
				{FieldName: "method", Value: "post"},
			},
		},
		ConfigOverrides: []*cxsdk.EntityTypeConfigOverrides{
			{
				EntityType: cxsdk.EntityTypeAlerts,
				Fields: []*cxsdk.TemplatedConnectorConfigField{
					{FieldName: "additionalBodyFields", Template: "{\"priority\": \"{{alertDef.priority}}\"}"},
				},
			},
		},
	}

	entityType := cxsdk.EntityTypeAlerts
	c := cxsdk.NewNotificationsClient(creator)
	success, err := c.TestConnectorConfig(context.Background(), &cxsdk.TestConnectorConfigRequest{
		PayloadType: "generic_https_default",
		Type:        connectorRaw.Type,
		Fields:      connectorRaw.ConnectorConfig.Fields,
		EntityType:  &entityType,
	})
	if err != nil {
		t.Fatal(err)
	}

	if success.Result.GetFailure() != nil {
		t.Logf("TestConnectorConfig is working, but the underlying webhook failed with error: %s", success.Result.GetFailure().Message)
	} else {
		assert.NotNil(t, success.Result.GetSuccess())
	}

	createRes, err := c.CreateConnector(context.Background(), &cxsdk.CreateConnectorRequest{
		Connector: &connectorRaw,
	})
	if err != nil {
		t.Fatal(err)
	}

	connectorId := createRes.Connector.Id

	connector, err := c.GetConnector(context.Background(), &cxsdk.GetConnectorRequest{
		Id: *connectorId,
	})

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, connector.Connector.Name, name)

	_, err = c.TestExistingConnector(context.Background(), &cxsdk.TestExistingConnectorRequest{
		ConnectorId: *connectorId,
		PayloadType: "generic_https_default",
	})

	if err != nil {
		t.Fatal(err)
	}

	_, err = c.DeleteConnector(context.Background(), &cxsdk.DeleteConnectorRequest{
		Id: *connectorId,
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestSlackConnector(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)

	connectorRaw := cxsdk.Connector{
		Type:        cxsdk.ConnectorTypeSlack,
		Name:        "TestSlackConnector",
		Description: "This is the slack connector to use for Notification Center testing.",
		ConnectorConfig: &cxsdk.ConnectorConfig{
			Fields: []*cxsdk.ConnectorConfigField{
				{FieldName: "integrationId", Value: "c6e72871-388b-4776-a7de-8b7d17ed4828"},
				{FieldName: "fallbackChannel", Value: "luigis-testing-grounds"},
				{FieldName: "channel", Value: "luigis-testing-grounds"},
			},
		},
		ConfigOverrides: []*cxsdk.EntityTypeConfigOverrides{
			{
				EntityType: cxsdk.EntityTypeAlerts,
				Fields: []*cxsdk.TemplatedConnectorConfigField{
					{FieldName: "channel", Template: "{{alertDef.priority}}"},
				},
			},
		},
	}

	entityType := cxsdk.EntityTypeAlerts
	c := cxsdk.NewNotificationsClient(creator)
	success, err := c.TestConnectorConfig(context.Background(), &cxsdk.TestConnectorConfigRequest{
		PayloadType: "slack_raw",
		Type:        connectorRaw.Type,
		Fields:      connectorRaw.ConnectorConfig.Fields,
		EntityType:  &entityType,
	})
	if err != nil {
		t.Fatal(err)
	}

	if success.Result.GetFailure() != nil {
		t.Fatal(success.Result.GetFailure().Message)
	} else {
		assert.NotNil(t, success.Result.GetSuccess())
	}

	createRes, err := c.CreateConnector(context.Background(), &cxsdk.CreateConnectorRequest{
		Connector: &connectorRaw,
	})
	if err != nil {
		t.Fatal(err)
	}

	connectorId := createRes.Connector.Id

	connector, err := c.GetConnector(context.Background(), &cxsdk.GetConnectorRequest{
		Id: *connectorId,
	})

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, connector.Connector.Name, "TestSlackConnector")

	_, err = c.TestExistingConnector(context.Background(), &cxsdk.TestExistingConnectorRequest{
		ConnectorId: *connectorId,
		PayloadType: "slack_raw",
	})

	if err != nil {
		t.Fatal(err)
	}

	_, err = c.DeleteConnector(context.Background(), &cxsdk.DeleteConnectorRequest{
		Id: *connectorId,
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestPagerdutyConnector(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)

	connectorRaw := cxsdk.Connector{
		Type:        cxsdk.ConnectorTypePagerDuty,
		Name:        "TestPagerdutyConnector",
		Description: "This is the PagerDuty connector to use for Notification Center testing.",
		ConnectorConfig: &cxsdk.ConnectorConfig{
			Fields: []*cxsdk.ConnectorConfigField{
				{FieldName: "integrationKey", Value: "test"},
			},
		},
		ConfigOverrides: []*cxsdk.EntityTypeConfigOverrides{
			{
				EntityType: cxsdk.EntityTypeAlerts,
				Fields: []*cxsdk.TemplatedConnectorConfigField{
					{FieldName: "integrationKey", Template: "integration_{{alertDef.priority}}"},
				},
			},
		},
	}

	c := cxsdk.NewNotificationsClient(creator)

	createRes, err := c.CreateConnector(context.Background(), &cxsdk.CreateConnectorRequest{
		Connector: &connectorRaw,
	})
	if err != nil {
		t.Fatal(err)
	}

	connectorId := createRes.Connector.Id

	connector, err := c.GetConnector(context.Background(), &cxsdk.GetConnectorRequest{
		Id: *connectorId,
	})

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, connector.Connector.Name, "TestPagerdutyConnector")

	if err != nil {
		t.Fatal(err)
	}

	_, err = c.DeleteConnector(context.Background(), &cxsdk.DeleteConnectorRequest{
		Id: *connectorId,
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestHttpsPreset(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)

	c := cxsdk.NewNotificationsClient(creator)
	presetName := "TestGoHttpsPreset"
	createRes, err := c.CreateCustomPreset(context.Background(), &cxsdk.CreateCustomPresetRequest{
		Preset: CreateHttpsPreset(presetName),
	})
	if err != nil {
		t.Fatal(err)
	}

	presetId := createRes.Preset.Id
	preset, err := c.GetPreset(context.Background(), &cxsdk.GetPresetRequest{
		Id: *presetId,
	})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, preset.Preset.Name, presetName)

	_, err = c.SetPresetAsDefault(context.Background(), &cxsdk.SetPresetAsDefaultRequest{
		Id: *presetId,
	})

	if err != nil {
		t.Fatal(err)
	}

	defaultPreset, err := c.GetDefaultPresetSummary(context.Background(), &cxsdk.GetDefaultPresetSummaryRequest{
		ConnectorType: cxsdk.ConnectorTypeGenericHTTPS,
		EntityType:    cxsdk.EntityTypeAlerts,
	})

	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, defaultPreset)
	assert.Equal(t, defaultPreset.PresetSummary.Name, presetName)

	_, err = c.DeleteCustomPreset(context.Background(), &cxsdk.DeleteCustomPresetRequest{
		Id: *presetId,
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestSlackPreset(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)

	c := cxsdk.NewNotificationsClient(creator)
	presetName := "TestGoSlackPreset"
	createRes, err := c.CreateCustomPreset(context.Background(), &cxsdk.CreateCustomPresetRequest{
		Preset: CreateSlackPreset(presetName),
	})
	if err != nil {
		t.Fatal(err)
	}

	presetId := createRes.Preset.Id
	preset, err := c.GetPreset(context.Background(), &cxsdk.GetPresetRequest{
		Id: *presetId,
	})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, preset.Preset.Name, presetName)

	_, err = c.SetPresetAsDefault(context.Background(), &cxsdk.SetPresetAsDefaultRequest{
		Id: *presetId,
	})

	if err != nil {
		t.Fatal(err)
	}

	defaultPreset, err := c.GetDefaultPresetSummary(context.Background(), &cxsdk.GetDefaultPresetSummaryRequest{
		ConnectorType: cxsdk.ConnectorTypeSlack,
		EntityType:    cxsdk.EntityTypeAlerts,
	})

	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, defaultPreset)
	assert.Equal(t, defaultPreset.PresetSummary.Name, presetName)

	_, err = c.DeleteCustomPreset(context.Background(), &cxsdk.DeleteCustomPresetRequest{
		Id: *presetId,
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestPagerdutyPreset(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)

	c := cxsdk.NewNotificationsClient(creator)
	presetName := "TestGoPagerdutyPreset"
	createRes, err := c.CreateCustomPreset(context.Background(), &cxsdk.CreateCustomPresetRequest{
		Preset: CreatePagerDutyPreset(presetName),
	})
	if err != nil {
		t.Fatal(err)
	}

	presetId := createRes.Preset.Id
	preset, err := c.GetPreset(context.Background(), &cxsdk.GetPresetRequest{
		Id: *presetId,
	})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, preset.Preset.Name, presetName)

	_, err = c.SetPresetAsDefault(context.Background(), &cxsdk.SetPresetAsDefaultRequest{
		Id: *presetId,
	})

	if err != nil {
		t.Fatal(err)
	}

	defaultPreset, err := c.GetDefaultPresetSummary(context.Background(), &cxsdk.GetDefaultPresetSummaryRequest{
		ConnectorType: cxsdk.ConnectorTypePagerDuty,
		EntityType:    cxsdk.EntityTypeAlerts,
	})

	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, defaultPreset)
	assert.Equal(t, defaultPreset.PresetSummary.Name, presetName)

	_, err = c.DeleteCustomPreset(context.Background(), &cxsdk.DeleteCustomPresetRequest{
		Id: *presetId,
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestGlobalRouter(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)

	notificationCenterClient := cxsdk.NewNotificationsClient(creator)
	alertsClient := cxsdk.NewAlertsClient(creator)

	_, err = notificationCenterClient.ListGlobalRouters(context.Background(), &cxsdk.ListGlobalRoutersRequest{
		EntityType: cxsdk.EntityTypeAlerts.Enum(),
	})
	if err != nil {
		t.Fatal(err)
	}

	routerId := "router_default"
	name := fmt.Sprintf("TestConnector-%v", uuid.NewString())
	connectorRaw := cxsdk.Connector{
		Type:        cxsdk.ConnectorTypeGenericHTTPS,
		Name:        name,
		Description: "This is the connector to use for Notification Center testing.",
		ConnectorConfig: &cxsdk.ConnectorConfig{
			Fields: []*cxsdk.ConnectorConfigField{
				{FieldName: "url", Value: "https://httpbun.org/post"},
				{FieldName: "method", Value: "post"},
			},
		},
		ConfigOverrides: []*cxsdk.EntityTypeConfigOverrides{
			{
				EntityType: cxsdk.EntityTypeAlerts,
				Fields: []*cxsdk.TemplatedConnectorConfigField{
					{FieldName: "additionalBodyFields", Template: "{\"priority\": \"{{alertDef.priority}}\"}"},
				},
			},
		},
	}

	createPresetResponse, err := notificationCenterClient.CreateCustomPreset(context.Background(), &cxsdk.CreateCustomPresetRequest{
		Preset: CreateHttpsPreset("TestGoHttpsPreset"),
	})

	if err != nil {
		t.Fatal(err)
	}
	presetId := createPresetResponse.Preset.Id

	createConnectorResponse, err := notificationCenterClient.CreateConnector(context.Background(), &cxsdk.CreateConnectorRequest{
		Connector: &connectorRaw,
	})

	if err != nil {
		t.Fatal(err)
	}

	routingRuleName := "TestRoutingRule"
	connectorId := createConnectorResponse.Connector.Id

	createOrReplaceRes, err := notificationCenterClient.CreateOrReplaceGlobalRouter(context.Background(), &cxsdk.CreateOrReplaceGlobalRouterRequest{
		Router: &cxsdk.GlobalRouter{
			Id:          &routerId,
			Name:        "TestGlobalRouter",
			EntityType:  cxsdk.EntityTypeAlerts,
			Description: "This is a test Global Router.",
			Rules: []*cxsdk.RoutingRule{
				{
					Name:      &routingRuleName,
					Condition: "alertDef.priority == \"P1\"",
					Targets: []*cxsdk.RoutingTarget{
						{
							ConnectorId: *connectorId,
							PresetId:    presetId,
						},
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	getRes, err := notificationCenterClient.GetGlobalRouter(context.Background(), &cxsdk.GetGlobalRouterRequest{
		Identifier: &cxsdk.GlobalRouterIdentifier{
			Value: &cxsdk.GlobalRouterIdentifierIDValue{
				Id: *createOrReplaceRes.Router.Id,
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	createdAlertDefWithRouter, err := alertsClient.Create(context.Background(), &cxsdk.CreateAlertDefRequest{
		AlertDefProperties: CreateAlertWithRouter(),
	})
	if err != nil {
		t.Fatal(err)
	}

	alertId := createdAlertDefWithRouter.AlertDef.Id
	_, err = alertsClient.Delete(context.Background(), &cxsdk.DeleteAlertDefRequest{
		Id: alertId,
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = notificationCenterClient.TestDestination(context.Background(), &cxsdk.TestDestinationRequest{
		EntityType:  cxsdk.EntityTypeAlerts,
		ConnectorId: *connectorId,
		PresetId:    *presetId,
		PayloadType: "generic_https_default",
	})

	if err != nil {
		t.Fatal(err)
	}

	_, err = notificationCenterClient.DeleteGlobalRouter(context.Background(), &cxsdk.DeleteGlobalRouterRequest{
		Id: *createOrReplaceRes.Router.Id,
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = notificationCenterClient.DeleteConnector(context.Background(), &cxsdk.DeleteConnectorRequest{
		Id: *connectorId,
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = notificationCenterClient.DeleteCustomPreset(context.Background(), &cxsdk.DeleteCustomPresetRequest{
		Id: *presetId,
	})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, getRes.Router.Name, "TestGlobalRouter")
}

func TestCreateAlertWithDestination(t *testing.T) {

	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)

	alertsClient := cxsdk.NewAlertsClient(creator)

	notificationCenterClient := cxsdk.NewNotificationsClient(creator)
	name := fmt.Sprintf("TestConnector-%v", uuid.NewString())
	connectorRaw := cxsdk.Connector{
		Type:        cxsdk.ConnectorTypeGenericHTTPS,
		Name:        name,
		Description: "This is the connector to use for Notification Center testing.",
		ConnectorConfig: &cxsdk.ConnectorConfig{
			Fields: []*cxsdk.ConnectorConfigField{
				{FieldName: "url", Value: "https://httpbun.org/post"},
				{FieldName: "method", Value: "post"},
			},
		},
		ConfigOverrides: []*cxsdk.EntityTypeConfigOverrides{
			{
				EntityType: cxsdk.EntityTypeAlerts,
				Fields: []*cxsdk.TemplatedConnectorConfigField{
					{FieldName: "additionalBodyFields", Template: "{\"priority\": \"{{alertDef.priority}}\"}"},
				},
			},
		},
	}

	createPresetResponse, err := notificationCenterClient.CreateCustomPreset(context.Background(), &cxsdk.CreateCustomPresetRequest{
		Preset: CreateHttpsPreset("TestGoHttpsPreset"),
	})

	if err != nil {
		t.Fatal(err)
	}
	presetId := createPresetResponse.Preset.Id

	createConnectorResponse, err := notificationCenterClient.CreateConnector(context.Background(), &cxsdk.CreateConnectorRequest{
		Connector: &connectorRaw,
	})

	if err != nil {
		t.Fatal(err)
	}

	connectorId := createConnectorResponse.Connector.Id

	createdAlertDefWithDestination, err := alertsClient.Create(context.Background(), &cxsdk.CreateAlertDefRequest{
		AlertDefProperties: CreateAlertWithDestination(
			*connectorId,
			presetId,
		),
	})

	if err != nil {
		t.Fatal(err)
	}

	_, err = notificationCenterClient.TestDestination(context.Background(), &cxsdk.TestDestinationRequest{
		EntityType:  cxsdk.EntityTypeAlerts,
		ConnectorId: *connectorId,
		PresetId:    *presetId,
		PayloadType: "generic_https_default",
	})

	if err != nil {
		t.Fatal(err)
	}

	alertId := createdAlertDefWithDestination.AlertDef.Id
	_, err = alertsClient.Delete(context.Background(), &cxsdk.DeleteAlertDefRequest{
		Id: alertId,
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = notificationCenterClient.DeleteConnector(context.Background(), &cxsdk.DeleteConnectorRequest{
		Id: *connectorId,
	})
	if err != nil {
		t.Fatal(err)
	}
	_, err = notificationCenterClient.DeleteCustomPreset(context.Background(), &cxsdk.DeleteCustomPresetRequest{
		Id: *presetId,
	})
	if err != nil {
		t.Fatal(err)
	}

}

func CreateAlertWithDestination(connectorId string, presetId *string) *cxsdk.AlertDefProperties {
	notifyOn := cxsdk.AlertNotifyOnTriggeredAndResolved
	notificationDestination := cxsdk.NotificationDestination{
		ConnectorId: connectorId,
		PresetId:    presetId,
	}
	return &cxsdk.AlertDefProperties{
		Name:              wrapperspb.String("Standard alert example"),
		Description:       wrapperspb.String("Standard alert example from terraform"),
		Enabled:           &wrapperspb.BoolValue{Value: true},
		Priority:          cxsdk.AlertDefPriorityP1,
		Deleted:           nil,
		Type:              cxsdk.AlertDefTypeLogsThreshold,
		IncidentsSettings: nil,
		PhantomMode:       &wrapperspb.BoolValue{Value: false},
		NotificationGroup: &cxsdk.AlertDefNotificationGroup{
			Destinations: []*cxsdk.NotificationDestination{
				&notificationDestination,
			},
			GroupByKeys: []*wrapperspb.StringValue{},
			Webhooks: []*cxsdk.AlertDefWebhooksSettings{
				{
					RetriggeringPeriod: &cxsdk.AlertDefWebhooksSettingsMinutes{
						Minutes: wrapperspb.UInt32(5),
					},
					NotifyOn: &notifyOn,
					Integration: &cxsdk.AlertDefIntegrationType{
						IntegrationType: &cxsdk.AlertDefIntegrationTypeRecipients{
							Recipients: &cxsdk.AlertDefRecipients{
								Emails: []*wrapperspb.StringValue{
									{Value: "example@coralogix.com"},
								},
							},
						},
					},
				},
			},
		},
		EntityLabels: map[string]string{
			"alert_type":        "security",
			"security_severity": "high",
		},
		Schedule: &cxsdk.AlertDefScheduleActiveOn{
			ActiveOn: &cxsdk.AlertsActivitySchedule{
				DayOfWeek: []cxsdk.AlertDayOfWeek{
					cxsdk.AlertDayOfWeekWednesday,
					cxsdk.AlertDayOfWeekThursday,
				},
				StartTime: &cxsdk.AlertTimeOfDay{
					Hours:   8,
					Minutes: 30,
				},
				EndTime: &cxsdk.AlertTimeOfDay{
					Hours:   20,
					Minutes: 30,
				},
			},
		},
		TypeDefinition: &cxsdk.AlertDefPropertiesLogsThreshold{
			LogsThreshold: &cxsdk.LogsThresholdType{
				Rules: []*cxsdk.LogsThresholdRule{
					{
						Override: &cxsdk.AlertDefPriorityOverride{
							Priority: cxsdk.AlertDefPriorityP1,
						},
						Condition: &cxsdk.LogsThresholdCondition{
							Threshold: wrapperspb.Double(10.0),
							TimeWindow: &cxsdk.LogsTimeWindow{
								Type: &cxsdk.LogsTimeWindowSpecificValue{
									LogsTimeWindowSpecificValue: cxsdk.LogsTimeWindowValue10Minutes,
								},
							},
							ConditionType: cxsdk.LogsThresholdConditionTypeMoreThanOrUnspecified,
						},
					},
				},
				LogsFilter: &cxsdk.LogsFilter{
					FilterType: &cxsdk.LogsFilterSimpleFilter{
						SimpleFilter: &cxsdk.SimpleFilter{
							LuceneQuery: wrapperspb.String("remote_addr_enriched:/.*/"),
							LabelFilters: &cxsdk.LabelFilters{
								ApplicationName: []*cxsdk.LabelFilterType{
									{Value: wrapperspb.String("nginx"), Operation: *cxsdk.LogFilterOperationIncludes.Enum()},
								},
								SubsystemName: []*cxsdk.LabelFilterType{
									{Value: wrapperspb.String("subsystem-name"), Operation: *cxsdk.LogFilterOperationStartsWith.Enum()},
								},
								Severities: []cxsdk.LogSeverity{
									*cxsdk.LogSeverityWarning.Enum(),
									*cxsdk.LogSeverityError.Enum(),
								},
							},
						},
					},
				},
			},
		},
		GroupByKeys: []*wrapperspb.StringValue{
			{Value: "coralogix.metadata.sdkId"},
			{Value: "EventType"},
		},
	}
}

func CreateAlertWithRouter() *cxsdk.AlertDefProperties {
	notifyOn := cxsdk.AlertNotifyOnTriggeredAndResolved
	return &cxsdk.AlertDefProperties{
		Name:              wrapperspb.String("Standard alert example"),
		Description:       wrapperspb.String("Standard alert example from terraform"),
		Enabled:           &wrapperspb.BoolValue{Value: true},
		Priority:          cxsdk.AlertDefPriorityP1,
		Deleted:           nil,
		Type:              cxsdk.AlertDefTypeLogsThreshold,
		IncidentsSettings: nil,
		PhantomMode:       &wrapperspb.BoolValue{Value: false},
		NotificationGroup: &cxsdk.AlertDefNotificationGroup{
			Router: &cxsdk.NotificationRouter{
				Id: "router_default",
			},
			GroupByKeys: []*wrapperspb.StringValue{},
			Webhooks: []*cxsdk.AlertDefWebhooksSettings{
				{
					RetriggeringPeriod: &cxsdk.AlertDefWebhooksSettingsMinutes{
						Minutes: wrapperspb.UInt32(5),
					},
					NotifyOn: &notifyOn,
					Integration: &cxsdk.AlertDefIntegrationType{
						IntegrationType: &cxsdk.AlertDefIntegrationTypeRecipients{
							Recipients: &cxsdk.AlertDefRecipients{
								Emails: []*wrapperspb.StringValue{
									{Value: "example@coralogix.com"},
								},
							},
						},
					},
				},
			},
		},
		EntityLabels: map[string]string{
			"alert_type":        "security",
			"security_severity": "high",
		},
		Schedule: &cxsdk.AlertDefScheduleActiveOn{
			ActiveOn: &cxsdk.AlertsActivitySchedule{
				DayOfWeek: []cxsdk.AlertDayOfWeek{
					cxsdk.AlertDayOfWeekWednesday,
					cxsdk.AlertDayOfWeekThursday,
				},
				StartTime: &cxsdk.AlertTimeOfDay{
					Hours:   8,
					Minutes: 30,
				},
				EndTime: &cxsdk.AlertTimeOfDay{
					Hours:   20,
					Minutes: 30,
				},
			},
		},
		TypeDefinition: &cxsdk.AlertDefPropertiesLogsThreshold{
			LogsThreshold: &cxsdk.LogsThresholdType{
				Rules: []*cxsdk.LogsThresholdRule{
					{
						Override: &cxsdk.AlertDefPriorityOverride{
							Priority: cxsdk.AlertDefPriorityP1,
						},
						Condition: &cxsdk.LogsThresholdCondition{
							Threshold: wrapperspb.Double(10.0),
							TimeWindow: &cxsdk.LogsTimeWindow{
								Type: &cxsdk.LogsTimeWindowSpecificValue{
									LogsTimeWindowSpecificValue: cxsdk.LogsTimeWindowValue10Minutes,
								},
							},
							ConditionType: cxsdk.LogsThresholdConditionTypeMoreThanOrUnspecified,
						},
					},
				},
				LogsFilter: &cxsdk.LogsFilter{
					FilterType: &cxsdk.LogsFilterSimpleFilter{
						SimpleFilter: &cxsdk.SimpleFilter{
							LuceneQuery: wrapperspb.String("remote_addr_enriched:/.*/"),
							LabelFilters: &cxsdk.LabelFilters{
								ApplicationName: []*cxsdk.LabelFilterType{
									{Value: wrapperspb.String("nginx"), Operation: *cxsdk.LogFilterOperationIncludes.Enum()},
								},
								SubsystemName: []*cxsdk.LabelFilterType{
									{Value: wrapperspb.String("subsystem-name"), Operation: *cxsdk.LogFilterOperationStartsWith.Enum()},
								},
								Severities: []cxsdk.LogSeverity{
									*cxsdk.LogSeverityWarning.Enum(),
									*cxsdk.LogSeverityError.Enum(),
								},
							},
						},
					},
				},
			},
		},
		GroupByKeys: []*wrapperspb.StringValue{
			{Value: "coralogix.metadata.sdkId"},
			{Value: "EventType"},
		},
	}
}

func CreateHttpsPreset(presetName string) *cxsdk.Preset {
	name := fmt.Sprintf("%v-%v", presetName, uuid.NewString())
	presetType := cxsdk.PresetTypeCustom
	parentId := "preset_system_generic_https_alerts_empty"
	return &cxsdk.Preset{
		Name:          name,
		Description:   "This is the preset to use for Notification Center testing.",
		PresetType:    &presetType,
		EntityType:    cxsdk.EntityTypeAlerts,
		ParentId:      &parentId,
		ConnectorType: cxsdk.ConnectorTypeGenericHTTPS,
		ConfigOverrides: []*cxsdk.ConfigOverrides{
			{
				ConditionType: &cxsdk.ConditionType{
					Condition: &cxsdk.ConditionTypeMatchEntityTypeAndSubType{
						MatchEntityTypeAndSubType: &cxsdk.MatchEntityTypeAndSubTypeCondition{
							EntityType:    cxsdk.EntityTypeAlerts,
							EntitySubType: "logsImmediateResolved",
						},
					},
				},
				MessageConfig: &cxsdk.MessageConfig{
					Fields: []*cxsdk.MessageConfigField{
						{
							Template:  "{}",
							FieldName: "headers",
						},
						{
							Template:  "{ \"groupingKey\": \"{{alert.groupingKey}}\", \"status\": \"{{alert.status}}\", \"groups\": \"{{alert.groups}}\" }",
							FieldName: "body",
						},
					},
				},
			},
		},
	}
}

func CreateSlackPreset(presetName string) *cxsdk.Preset {
	presetType := cxsdk.PresetTypeCustom
	parentId := "preset_system_slack_alerts_empty"
	return &cxsdk.Preset{
		Name:          presetName,
		Description:   "This is the preset to use for Notification Center testing.",
		PresetType:    &presetType,
		EntityType:    cxsdk.EntityTypeAlerts,
		ParentId:      &parentId,
		ConnectorType: cxsdk.ConnectorTypeSlack,
		ConfigOverrides: []*cxsdk.ConfigOverrides{
			{
				ConditionType: &cxsdk.ConditionType{
					Condition: &cxsdk.ConditionTypeMatchEntityTypeAndSubType{
						MatchEntityTypeAndSubType: &cxsdk.MatchEntityTypeAndSubTypeCondition{
							EntityType:    cxsdk.EntityTypeAlerts,
							EntitySubType: "logsImmediateResolved",
						},
					},
				},
				MessageConfig: &cxsdk.MessageConfig{
					Fields: []*cxsdk.MessageConfigField{
						{
							Template:  "{{alert.status}} {{alertDef.priority}} - {{alertDef.name}}",
							FieldName: "title",
						},
						{
							Template:  "{{alert.timestamp |  date(format=\"%Y-%m-%d %H:%M\")}}\nWe've detected a log that matches the query.",
							FieldName: "description",
						},
					},
				},
			},
		},
	}
}

func CreatePagerDutyPreset(presetName string) *cxsdk.Preset {
	presetType := cxsdk.PresetTypeCustom
	parentId := "preset_system_pagerduty_alerts_basic"
	return &cxsdk.Preset{
		Name:          presetName,
		Description:   "This is the preset to use for Notification Center testing.",
		PresetType:    &presetType,
		EntityType:    cxsdk.EntityTypeAlerts,
		ParentId:      &parentId,
		ConnectorType: cxsdk.ConnectorTypePagerDuty,
		ConfigOverrides: []*cxsdk.ConfigOverrides{
			{
				ConditionType: &cxsdk.ConditionType{
					Condition: &cxsdk.ConditionTypeMatchEntityTypeAndSubType{
						MatchEntityTypeAndSubType: &cxsdk.MatchEntityTypeAndSubTypeCondition{
							EntityType:    cxsdk.EntityTypeAlerts,
							EntitySubType: "logsImmediateTriggered",
						},
					},
				},
				MessageConfig: &cxsdk.MessageConfig{
					Fields: []*cxsdk.MessageConfigField{
						{
							Template:  "{{ alertDef.name }}",
							FieldName: "summary",
						},
					},
				},
			},
		},
	}
}
