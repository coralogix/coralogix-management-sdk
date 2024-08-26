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
	v3 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/alerts/v3"
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

	createdAlertDef, err := c.Create(context.Background(), &cxsdk.CreateAlertRequest{
		AlertDefProperties: &cxsdk.AlertDefProperties{
			Name:         &wrapperspb.StringValue{Value: "Standard alert example"},
			Description:  &wrapperspb.StringValue{Value: "Standard alert example from terraform"},
			Enabled:      &wrapperspb.BoolValue{Value: true},
			Priority:     cxsdk.AlertDefPriorityP1,
			AlertDefType: cxsdk.AlertDefTypeLogsMoreThan,
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
				Targets: &cxsdk.AlertDefSimpleNotificationGroup{
					Simple: &cxsdk.AlertDefSimpleTarget{
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
			Schedule: &v3.AlertDefProperties_ActiveOn{
				ActiveOn: &v3.ActivitySchedule{
					DayOfWeek: []v3.DayOfWeek{
						v3.DayOfWeek_DAY_OF_WEEK_WEDNESDAY,
						v3.DayOfWeek_DAY_OF_WEEK_THURSDAY,
					},
					StartTime: &v3.TimeOfDay{
						Hours:   8,
						Minutes: 30,
					},
					EndTime: &v3.TimeOfDay{
						Hours:   20,
						Minutes: 30,
					},
				},
			},
			TypeDefinition: &v3.AlertDefProperties_LogsMoreThan{
				LogsMoreThan: &standard.LogsMoreThanTypeDefinition{
					LogsFilter: &alert_def_type_definition.LogsFilter{
						FilterType: &alert_def_type_definition.LogsFilter_LuceneFilter{
							LuceneFilter: &alert_def_type_definition.LuceneFilter{
								LuceneQuery: &wrapperspb.StringValue{Value: "remote_addr_enriched:/.*/"},
								LabelFilters: &alert_def_type_definition.LabelFilters{
									ApplicationName: []*alert_def_type_definition.LabelFilterType{
										{Value: &wrapperspb.StringValue{Value: "nginx"}, Operation: *alert_def_type_definition.LogFilterOperationType_LOG_FILTER_OPERATION_TYPE_INCLUDES.Enum()},
									},
									SubsystemName: []*alert_def_type_definition.LabelFilterType{
										{Value: &wrapperspb.StringValue{Value: "subsystem-name"}, Operation: *alert_def_type_definition.LogFilterOperationType_LOG_FILTER_OPERATION_TYPE_STARTS_WITH.Enum()},
									},
									Severities: []alert_def_type_definition.LogSeverity{
										*alert_def_type_definition.LogSeverity_LOG_SEVERITY_WARNING.Enum(),
										*alert_def_type_definition.LogSeverity_LOG_SEVERITY_ERROR.Enum(),
									},
								},
							},
						},
					},
					Threshold: &wrapperspb.UInt32Value{Value: 5},
					TimeWindow: &standard.LogsTimeWindow{
						Type: &standard.LogsTimeWindow_LogsTimeWindowSpecificValue{
							LogsTimeWindowSpecificValue: standard.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_MINUTES_30,
						},
					},
					EvaluationWindow:          standard.EvaluationWindow_EVALUATION_WINDOW_ROLLING_OR_UNSPECIFIED,
					NotificationPayloadFilter: []*wrapperspb.StringValue{},
				},
			},
		},
	})
	assert.Nil(t, err)

	retrievedAlert, err := c.Get(context.Background(), &v3.GetAlertDefRequest{
		Id: createdAlertDef.AlertDef.Id,
	})

	assert.Nil(t, err)

	updatedAlertDef := retrievedAlert.AlertDef
	updatedAlertDef.AlertDefProperties.Description = &wrapperspb.StringValue{Value: "Updated description"}

	updatedAlert, err := c.Replace(context.Background(), &v3.ReplaceAlertDefRequest{
		AlertDefProperties: updatedAlertDef.AlertDefProperties,
		Id:                 updatedAlertDef.Id,
	})

	assert.Nil(t, err)

	_, e := c.Delete(context.Background(), &v3.DeleteAlertDefRequest{
		Id: updatedAlert.AlertDef.Id,
	})

	assert.Nil(t, e)
}
