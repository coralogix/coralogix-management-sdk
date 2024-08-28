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
	"github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/alerts/v3/alert_def_type_definition"
	"github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/alerts/v3/alert_def_type_definition/standard"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestAlerts(t *testing.T) {

	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	c := cxsdk.NewAlertsClient(creator)

	createdAlertDef, err := c.Create(context.Background(), &cxsdk.CreateAlertDefRequest{
		AlertDefProperties: &cxsdk.AlertDefProperties{
			Name:              wrapperspb.String("Standard alert example"),
			Description:       wrapperspb.String("Standard alert example from terraform"),
			Enabled:           &wrapperspb.BoolValue{Value: true},
			Priority:          cxsdk.AlertDefPriorityP1,
			Deleted:           nil,
			Type:              cxsdk.AlertDefTypeLogsThreshold,
			IncidentsSettings: nil,
			PhantomMode:       &wrapperspb.BoolValue{Value: false},
			NotificationGroup: &cxsdk.AlertDefNotificationGroup{
				GroupByFields: []*wrapperspb.StringValue{},
				Targets: &cxsdk.AlertDefNotificationGroupSimple{
					Simple: &cxsdk.AlertDefTargetSimple{
						Integrations: []*cxsdk.AlertDefIntegrationType{
							{IntegrationType: &cxsdk.AlertDefIntegrationTypeRecipients{
								Recipients: &cxsdk.AlertDefRecipients{
									Emails: []*wrapperspb.StringValue{
										{Value: "example@coralogix.com"},
									},
								},
							}},
						},
					},
				},
			},
			Labels: map[string]string{
				"alert_type":        "security",
				"security_severity": "high",
			},
			Schedule: &cxsdk.AlertDefScheduleActiveOn{
				ActiveOn: &cxsdk.AlertsActivitySchedule{
					DayOfWeek: []cxsdk.AlertsDayOfWeek{
						cxsdk.AlertsDayOfWeekWednesday,
						cxsdk.AlertsDayOfWeekThursday,
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
			TypeDefinition: &cxsdk.AlertDefPropertiesLogsMoreThan{
				LogsMoreThan: &standard.LogsMoreThanTypeDefinition{
					LogsFilter: &alert_def_type_definition.LogsFilter{
						FilterType: &alert_def_type_definition.LogsFilter_LuceneFilter{
							LuceneFilter: &alert_def_type_definition.LuceneFilter{
								LuceneQuery: &wrapperspb.StringValue{Value: "remote_addr_enriched:/.*/"},
								LabelFilters: &alert_def_type_definition.LabelFilters{
									ApplicationName: []*alert_def_type_definition.LabelFilterType{
										{Value: &wrapperspb.StringValue{Value: "nginx"}, Operation: *alert_def_type_definition.LogFilterOperationType_LOG_FILTER_OPERATION_TYPE_INCLUDES.Enum()},
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
			GroupBy: []*wrapperspb.StringValue{
				{Value: "coralogix.metadata.sdkId"},
				{Value: "EventType"},
			},
		},
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
	a := cxsdk.NewAlertsSchedulersClient(creator)
	description := "example"
	metaLabels := make([]*cxsdk.MetaLabel, 0)
	schedule := cxsdk.Schedule{
		ScheduleOperation: cxsdk.ScheduleOperationScheduleOperationMute,
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
						Value: []string{"55a457ed-5f23-407a-a724-12d7fe533a4e"},
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

}
