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

package examples

import (
	"context"
	"testing"

	cxsdk "github.com/coralogix/coralogix-management-sdk/go"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func CreateAlert() *cxsdk.AlertDefProperties {
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

func CreateBurnRateSloAlert(sloId string) *cxsdk.AlertDefProperties {

	notifyOn := cxsdk.AlertNotifyOnTriggeredAndResolved
	return &cxsdk.AlertDefProperties{
		Name:              wrapperspb.String("Standard alert example"),
		Description:       wrapperspb.String("Standard alert example from terraform"),
		Enabled:           &wrapperspb.BoolValue{Value: true},
		Priority:          cxsdk.AlertDefPriorityP1,
		Deleted:           nil,
		Type:              cxsdk.AlertDefTypeSloThreshold,
		IncidentsSettings: nil,
		PhantomMode:       &wrapperspb.BoolValue{Value: false},
		NotificationGroup: &cxsdk.AlertDefNotificationGroup{
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
		TypeDefinition: &cxsdk.AlertDefPropertiesSlo{
			SloThreshold: &cxsdk.SloThresholdType{
				SloDefinition: &cxsdk.AlertSloDefinition{
					SloId: &wrapperspb.StringValue{Value: sloId},
				},
				Threshold: &cxsdk.SloBurnRateThresholdType{
					BurnRate: &cxsdk.SloBurnRateThreshold{
						Rules: []*cxsdk.SloThresholdRule{
							{
								Condition: &cxsdk.SloThresholdCondition{
									Threshold: wrapperspb.Double(1.0),
								},
								Override: &cxsdk.AlertDefPriorityOverride{
									Priority: cxsdk.AlertDefPriorityP1,
								},
							},
						},
						Type: &cxsdk.SingleBurnRateThresholdType{
							Single: &cxsdk.SingleBurnRateThreshold{
								TimeDuration: &cxsdk.TimeDuration{
									Duration: &wrapperspb.UInt64Value{Value: 1},
									Unit:     cxsdk.DurationUnitHours,
								},
							},
						},
					},
				},
			},
		},
		GroupByKeys: []*wrapperspb.StringValue{},
	}
}

func CreateErrorBudgetSloAlert(sloId string) *cxsdk.AlertDefProperties {
	notifyOn := cxsdk.AlertNotifyOnTriggeredAndResolved
	return &cxsdk.AlertDefProperties{
		Name:              wrapperspb.String("Standard alert example"),
		Description:       wrapperspb.String("Standard alert example from terraform"),
		Enabled:           &wrapperspb.BoolValue{Value: true},
		Priority:          cxsdk.AlertDefPriorityP1,
		Deleted:           nil,
		Type:              cxsdk.AlertDefTypeSloThreshold,
		IncidentsSettings: nil,
		PhantomMode:       &wrapperspb.BoolValue{Value: false},
		NotificationGroup: &cxsdk.AlertDefNotificationGroup{
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
		TypeDefinition: &cxsdk.AlertDefPropertiesSlo{
			SloThreshold: &cxsdk.SloThresholdType{
				SloDefinition: &cxsdk.AlertSloDefinition{
					SloId: &wrapperspb.StringValue{Value: sloId},
				},
				Threshold: &cxsdk.SloBurnRateThresholdType{
					BurnRate: &cxsdk.SloBurnRateThreshold{
						Rules: []*cxsdk.SloThresholdRule{
							{
								Condition: &cxsdk.SloThresholdCondition{
									Threshold: wrapperspb.Double(1.0),
								},
								Override: &cxsdk.AlertDefPriorityOverride{
									Priority: cxsdk.AlertDefPriorityP1,
								},
							},
						},
						Type: &cxsdk.SingleBurnRateThresholdType{
							Single: &cxsdk.SingleBurnRateThreshold{
								TimeDuration: &cxsdk.TimeDuration{
									Duration: &wrapperspb.UInt64Value{Value: 1},
									Unit:     cxsdk.DurationUnitHours,
								},
							},
						},
					},
				},
			},
		},
		GroupByKeys: []*wrapperspb.StringValue{},
	}
}

func TestAlerts(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	c := cxsdk.NewAlertsClient(creator)

	createdAlertDef, err := c.Create(context.Background(), &cxsdk.CreateAlertDefRequest{
		AlertDefProperties: CreateAlert(),
	})

	assertNilAndPrintError(t, err)

	retrievedAlert, err := c.Get(context.Background(), &cxsdk.GetAlertDefRequest{
		Id: createdAlertDef.AlertDef.Id,
	})

	assertNilAndPrintError(t, err)

	alerts, err := c.List(context.Background(), &cxsdk.ListAlertDefsRequest{})
	assertNilAndPrintError(t, err)
	assert.Greater(t, len(alerts.AlertDefs), 0)

	updatedAlertDef := retrievedAlert.AlertDef
	updatedAlertDef.AlertDefProperties.Description = &wrapperspb.StringValue{Value: "Updated description"}

	updatedAlert, err := c.Replace(context.Background(), &cxsdk.ReplaceAlertDefRequest{
		AlertDefProperties: updatedAlertDef.AlertDefProperties,
		Id:                 updatedAlertDef.Id,
	})

	assertNilAndPrintError(t, err)

	_, e := c.Delete(context.Background(), &cxsdk.DeleteAlertDefRequest{
		Id: updatedAlert.AlertDef.Id,
	})

	assertNilAndPrintError(t, e)

	_, err = c.Get(context.Background(), &cxsdk.GetAlertDefRequest{
		Id: updatedAlert.AlertDef.Id,
	})
	// Returns NotFound
	assert.NotNil(t, err)
}

func TestBurnRateSloAlerts(t *testing.T) {
	t.Skip("Unstable Test")
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	alertsClient := cxsdk.NewAlertsClient(creator)
	slosClient := cxsdk.NewSLOsClient(creator)
	sloDescription := "description"

	createSloResponse, err := slosClient.Create(context.Background(), &cxsdk.CreateServiceSloRequest{
		Slo: &cxsdk.Slo{
			Name:                      "coralogix_slo_go_example",
			Description:               &sloDescription,
			TargetThresholdPercentage: 30,
			Sli: &cxsdk.SloRequestBasedMetricSli{
				RequestBasedMetricSli: &cxsdk.RequestBasedMetricSli{
					GoodEvents: &cxsdk.Metric{
						Query: "avg(rate(cpu_usage_seconds_total[5m])) by (instance)",
					},
					TotalEvents: &cxsdk.Metric{
						Query: "avg(rate(cpu_usage_seconds_total[5m])) by (instance)",
					},
				},
			},
			Window: &cxsdk.SloTimeframe{
				SloTimeFrame: cxsdk.SloTimeframe7Days,
			},
			Labels: map[string]string{
				"label1": "value1",
			},
		},
	})
	if err != nil {
		t.Fatalf("Failed to create SLO: %v", err)
	}

	createdAlertDef, err := alertsClient.Create(context.Background(), &cxsdk.CreateAlertDefRequest{
		AlertDefProperties: CreateBurnRateSloAlert(*createSloResponse.Slo.Id),
	})

	assertNilAndPrintError(t, err)

	retrievedAlert, err := alertsClient.Get(context.Background(), &cxsdk.GetAlertDefRequest{
		Id: createdAlertDef.AlertDef.Id,
	})

	assertNilAndPrintError(t, err)

	alerts, err := alertsClient.List(context.Background(), &cxsdk.ListAlertDefsRequest{})
	assertNilAndPrintError(t, err)
	assert.Greater(t, len(alerts.AlertDefs), 0)

	updatedAlertDef := retrievedAlert.AlertDef
	updatedAlertDef.AlertDefProperties.Description = &wrapperspb.StringValue{Value: "Updated description"}

	updatedAlert, err := alertsClient.Replace(context.Background(), &cxsdk.ReplaceAlertDefRequest{
		AlertDefProperties: updatedAlertDef.AlertDefProperties,
		Id:                 updatedAlertDef.Id,
	})

	assertNilAndPrintError(t, err)

	_, e := alertsClient.Delete(context.Background(), &cxsdk.DeleteAlertDefRequest{
		Id: updatedAlert.AlertDef.Id,
	})

	assertNilAndPrintError(t, e)
	_, e = slosClient.Delete(context.Background(), &cxsdk.DeleteServiceSloRequest{
		Id: *createSloResponse.Slo.Id,
	})
	assertNilAndPrintError(t, e)
}

func TestAlertGetsDeletedOnSloDeletion(t *testing.T) {
	t.Skip("Unstable Test")

	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	alertsClient := cxsdk.NewAlertsClient(creator)
	slosClient := cxsdk.NewSLOsClient(creator)
	sloDescription := "description"

	createSloResponse, err := slosClient.Create(context.Background(), &cxsdk.CreateServiceSloRequest{
		Slo: &cxsdk.Slo{
			Name:                      "coralogix_slo_go_example",
			Description:               &sloDescription,
			TargetThresholdPercentage: 30,
			Sli: &cxsdk.SloRequestBasedMetricSli{
				RequestBasedMetricSli: &cxsdk.RequestBasedMetricSli{
					GoodEvents: &cxsdk.Metric{
						Query: "avg(rate(cpu_usage_seconds_total[5m])) by (instance)",
					},
					TotalEvents: &cxsdk.Metric{
						Query: "avg(rate(cpu_usage_seconds_total[5m])) by (instance)",
					},
				},
			},
			Window: &cxsdk.SloTimeframe{
				SloTimeFrame: cxsdk.SloTimeframe7Days,
			},
			Labels: map[string]string{
				"label1": "value1",
			},
		},
	})
	if err != nil {
		t.Fatalf("Failed to create SLO: %v", err)
	}

	createdAlertDef, err := alertsClient.Create(context.Background(), &cxsdk.CreateAlertDefRequest{
		AlertDefProperties: CreateBurnRateSloAlert(*createSloResponse.Slo.Id),
	})

	assertNilAndPrintError(t, err)

	retrievedAlert, err := alertsClient.Get(context.Background(), &cxsdk.GetAlertDefRequest{
		Id: createdAlertDef.AlertDef.Id,
	})

	assertNilAndPrintError(t, err)

	alerts, err := alertsClient.List(context.Background(), &cxsdk.ListAlertDefsRequest{})
	assertNilAndPrintError(t, err)
	assert.Greater(t, len(alerts.AlertDefs), 0)

	updatedAlertDef := retrievedAlert.AlertDef
	updatedAlertDef.AlertDefProperties.Description = &wrapperspb.StringValue{Value: "Updated description"}

	updateAlertResponse, err := alertsClient.Replace(context.Background(), &cxsdk.ReplaceAlertDefRequest{
		AlertDefProperties: updatedAlertDef.AlertDefProperties,
		Id:                 updatedAlertDef.Id,
	})

	assertNilAndPrintError(t, err)

	_, err = slosClient.Delete(context.Background(), &cxsdk.DeleteServiceSloRequest{
		Id: *createSloResponse.Slo.Id,
	})
	assertNilAndPrintError(t, err)

	_, err = alertsClient.Get(context.Background(), &cxsdk.GetAlertDefRequest{
		Id: updateAlertResponse.AlertDef.Id,
	})

	// Returns NotFound
	assert.NotNil(t, err)
}

func TestAlertScheduler(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	a := cxsdk.NewAlertSchedulerClient(creator)
	c := cxsdk.NewAlertsClient(creator)

	createdAlertDef, err := c.Create(context.Background(), &cxsdk.CreateAlertDefRequest{
		AlertDefProperties: CreateAlert(),
	})

	assertNilAndPrintError(t, err)

	description := "example"
	metaLabels := make([]*cxsdk.MetaLabel, 0)
	schedule := cxsdk.Schedule{
		ScheduleOperation: cxsdk.ScheduleOperationMute,
		Scheduler: &cxsdk.ScheduleOneTime{
			OneTime: &cxsdk.OneTime{
				Timeframe: &cxsdk.Timeframe{
					StartTime: "2021-01-04T00:00:00.000",
					Until: &cxsdk.TimeframeEndTime{
						EndTime: "2025-01-01T00:00:50.000",
					},
					Timezone: "UTC+2",
				},
			},
		},
	}
	createAlertSchedulerResponse, e := a.Create(context.Background(), &cxsdk.CreateAlertSchedulerRuleRequest{
		AlertSchedulerRule: &cxsdk.AlertSchedulerRule{
			Name:        "example",
			Description: &description,
			MetaLabels:  metaLabels,
			Schedule:    &schedule,
			Filter: &cxsdk.AlertSchedulerFilter{
				WhatExpression: "source logs | filter $d.cpodId:string == '122'",
				WhichAlerts: &cxsdk.AlertSchedulerFilterUniqueIDs{
					AlertUniqueIds: &cxsdk.AlertUniqueIDs{
						Value: []string{createdAlertDef.AlertDef.Id.Value},
					},
				},
			},
			Enabled:   true,
			CreatedAt: nil,
			UpdatedAt: nil,
		},
	})

	assertNilAndPrintError(t, e)

	alertScheduler := createAlertSchedulerResponse.AlertSchedulerRule
	alertScheduler.Name = "MyAlertUpdated"
	alertScheduler.Id = nil

	updateAlertSchedulerResponse, e := a.Update(context.Background(), &cxsdk.UpdateAlertSchedulerRuleRequest{
		AlertSchedulerRule: alertScheduler,
	})

	assertNilAndPrintError(t, e)

	assert.Equal(t, updateAlertSchedulerResponse.AlertSchedulerRule.Name, "MyAlertUpdated")

	getAlertSchedulerResponse, e := a.Get(context.Background(), &cxsdk.GetAlertSchedulerRuleRequest{
		AlertSchedulerRuleId: *updateAlertSchedulerResponse.AlertSchedulerRule.UniqueIdentifier,
	})

	assertNilAndPrintError(t, e)

	assert.Equal(t, getAlertSchedulerResponse.AlertSchedulerRule.Name, "MyAlertUpdated")

	_, error := a.Delete(context.Background(), &cxsdk.DeleteAlertSchedulerRuleRequest{
		AlertSchedulerRuleId: *updateAlertSchedulerResponse.AlertSchedulerRule.UniqueIdentifier,
	})
	assertNilAndPrintError(t, error)

	_, e = c.Delete(context.Background(), &cxsdk.DeleteAlertDefRequest{
		Id: createdAlertDef.AlertDef.Id,
	})

	assertNilAndPrintError(t, e)
}
