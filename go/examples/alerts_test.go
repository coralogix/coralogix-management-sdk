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
					{Condition: &cxsdk.LogsThresholdCondition{
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

func TestAlerts(t *testing.T) {

	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	c := cxsdk.NewAlertsClient(creator)

	createdAlertDef, err := c.Create(context.Background(), &cxsdk.CreateAlertDefRequest{
		AlertDefProperties: CreateAlert(),
	})

	if err != nil {
		t.Errorf("Error creating alert: %v", err)
	}
	assert.Nil(t, err)

	retrievedAlert, err := c.Get(context.Background(), &cxsdk.GetAlertDefRequest{
		Id: createdAlertDef.AlertDef.Id,
	})

	assert.Nil(t, err)

	updatedAlertDef := retrievedAlert.AlertDef
	updatedAlertDef.AlertDefProperties.Description = &wrapperspb.StringValue{Value: "Updated description"}

	updatedAlert, err := c.Replace(context.Background(), &cxsdk.ReplaceAlertDefRequest{
		AlertDefProperties: updatedAlertDef.AlertDefProperties,
		Id:                 updatedAlertDef.Id,
	})

	assert.Nil(t, err)

	_, e := c.Delete(context.Background(), &cxsdk.DeleteAlertDefRequest{
		Id: updatedAlert.AlertDef.Id,
	})

	assert.Nil(t, e)
}

func TestAlertScheduler(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	a := cxsdk.NewAlertSchedulerClient(creator)
	c := cxsdk.NewAlertsClient(creator)

	createdAlertDef, err := c.Create(context.Background(), &cxsdk.CreateAlertDefRequest{
		AlertDefProperties: CreateAlert(),
	})

	if err != nil {
		t.Errorf("Error creating alert: %v", err)
	}
	assert.Nil(t, err)

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

	assert.Nil(t, e)

	alertScheduler := createAlertSchedulerResponse.AlertSchedulerRule
	alertScheduler.Name = "MyAlertUpdated"

	updateAlertSchedulerResponse, e := a.Update(context.Background(), &cxsdk.UpdateAlertSchedulerRuleRequest{
		AlertSchedulerRule: alertScheduler,
	})

	assert.Nil(t, e)

	assert.Equal(t, updateAlertSchedulerResponse.AlertSchedulerRule.Name, "MyAlertUpdated")

	getAlertSchedulerResponse, e := a.Get(context.Background(), &cxsdk.GetAlertSchedulerRuleRequest{
		AlertSchedulerRuleId: *updateAlertSchedulerResponse.AlertSchedulerRule.UniqueIdentifier,
	})

	assert.Nil(t, e)

	assert.Equal(t, getAlertSchedulerResponse.AlertSchedulerRule.Name, "MyAlertUpdated")

	_, error := a.Delete(context.Background(), &cxsdk.DeleteAlertSchedulerRuleRequest{
		AlertSchedulerRuleId: *updateAlertSchedulerResponse.AlertSchedulerRule.UniqueIdentifier,
	})
	assert.Nil(t, error)

	_, e = c.Delete(context.Background(), &cxsdk.DeleteAlertDefRequest{
		Id: createdAlertDef.AlertDef.Id,
	})

	assert.Nil(t, e)
}
