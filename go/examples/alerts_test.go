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

func TestAlerts(t *testing.T) {

	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	c := cxsdk.NewAlertsClient(creator)

	createdAlertDef, err := c.Create(context.Background(), &cxsdk.CreateAlertDefRequest{
		AlertDefProperties: &cxsdk.AlertDefProperties{
			Name:        &wrapperspb.StringValue{Value: "Standard alert example"},
			Description: &wrapperspb.StringValue{Value: "Standard alert example from terraform"},
			Enabled:     &wrapperspb.BoolValue{Value: true},
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
			Priority: cxsdk.AlertDefPriorityP1,
			GroupBy: []*wrapperspb.StringValue{
				{Value: "coralogix.metadata.sdkId"},
				{Value: "EventType"},
			},
			IncidentsSettings: nil,
			NotificationGroup: &cxsdk.AlertDefNotificationGroup{
				GroupByFields: []*wrapperspb.StringValue{
					{Value: "coralogix.metadata.sdkId"},
					{Value: "EventType"},
				},
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
