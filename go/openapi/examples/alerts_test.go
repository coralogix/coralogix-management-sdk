package examples

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	alerts "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/alert_definitions_service"
	scheduler "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/alert_scheduler_rule_service"
)

// CreateTracingImmediateAlert returns a reusable OpenAPI alert definition payload
func CreateTracingImmediateAlert() *alerts.AlertDefPropertiesTracingImmediate {
	name := "Standard alert example"
	startHour := int32(8)
	startMinute := int32(30)
	endHour := int32(20)
	endMinute := int32(30)
	latencyThresholdMs := "30"
	// applicationName := "nginx"
	// subsystemName := "subsystem-name"
	return &alerts.AlertDefPropertiesTracingImmediate{
		Name:        &name,
		Description: alerts.PtrString("Standard alert example from OpenAPI SDK"),
		Enabled:     alerts.PtrBool(true),
		Priority:    alerts.ALERTDEFPRIORITY_ALERT_DEF_PRIORITY_P1.Ptr(),
		Type:        alerts.ALERTDEFTYPE_ALERT_DEF_TYPE_TRACING_IMMEDIATE.Ptr(),
		EntityLabels: &map[string]string{
			"alert_type":        "security",
			"security_severity": "high",
		},
		PhantomMode: alerts.PtrBool(false),
		NotificationGroup: &alerts.AlertDefNotificationGroup{
			Webhooks: []alerts.AlertDefWebhooksSettings{
				{
					Minutes:  alerts.PtrInt64(5),
					NotifyOn: alerts.NOTIFYON_NOTIFY_ON_TRIGGERED_AND_RESOLVED.Ptr(),
					Integration: &alerts.V3IntegrationType{
						V3IntegrationTypeRecipients: &alerts.V3IntegrationTypeRecipients{
							Recipients: &alerts.Recipients{
								Emails: []string{"example@coralogix.com"},
							},
						},
					},
				},
			},
		},
		ActiveOn: &alerts.ActivitySchedule{
			DayOfWeek: []alerts.DayOfWeek{
				alerts.DAYOFWEEK_DAY_OF_WEEK_WEDNESDAY,
				alerts.DAYOFWEEK_DAY_OF_WEEK_THURSDAY,
			},
			StartTime: &alerts.TimeOfDay{Hours: &startHour, Minutes: &startMinute},
			EndTime:   &alerts.TimeOfDay{Hours: &endHour, Minutes: &endMinute},
		},
		TracingImmediate: &alerts.TracingImmediateType{
			TracingFilter: &alerts.TracingFilter{
				SimpleFilter: &alerts.TracingSimpleFilter{
					LatencyThresholdMs: &latencyThresholdMs,
				},
			},
		},
	}
}

func CreateLogsRatioAlert() *alerts.AlertDefPropertiesLogsRatioThreshold {
	name := "Standard alert example"
	startHour := int32(8)
	startMinute := int32(30)
	endHour := int32(20)
	endMinute := int32(30)
	numeratorAlias := "numerator"
	denominatorAlias := "denominator"
	threshold := float64(2)
	return &alerts.AlertDefPropertiesLogsRatioThreshold{
		Name:        &name,
		Description: alerts.PtrString("Standard alert example from OpenAPI SDK"),
		Enabled:     alerts.PtrBool(true),
		Priority:    alerts.ALERTDEFPRIORITY_ALERT_DEF_PRIORITY_P1.Ptr(),
		Type:        alerts.ALERTDEFTYPE_ALERT_DEF_TYPE_LOGS_RATIO_THRESHOLD.Ptr(),
		EntityLabels: &map[string]string{
			"alert_type":        "security",
			"security_severity": "high",
		},
		PhantomMode: alerts.PtrBool(false),
		NotificationGroup: &alerts.AlertDefNotificationGroup{
			Webhooks: []alerts.AlertDefWebhooksSettings{
				{
					Minutes:  alerts.PtrInt64(5),
					NotifyOn: alerts.NOTIFYON_NOTIFY_ON_TRIGGERED_AND_RESOLVED.Ptr(),
					Integration: &alerts.V3IntegrationType{
						V3IntegrationTypeRecipients: &alerts.V3IntegrationTypeRecipients{
							Recipients: &alerts.Recipients{
								Emails: []string{"example@coralogix.com"},
							},
						},
					},
				},
			},
		},
		ActiveOn: &alerts.ActivitySchedule{
			DayOfWeek: []alerts.DayOfWeek{
				alerts.DAYOFWEEK_DAY_OF_WEEK_WEDNESDAY,
				alerts.DAYOFWEEK_DAY_OF_WEEK_THURSDAY,
			},
			StartTime: &alerts.TimeOfDay{Hours: &startHour, Minutes: &startMinute},
			EndTime:   &alerts.TimeOfDay{Hours: &endHour, Minutes: &endMinute},
		},
		LogsRatioThreshold: &alerts.LogsRatioThresholdType{
			NumeratorAlias:   &numeratorAlias,
			DenominatorAlias: &denominatorAlias,
			Rules: []alerts.LogsRatioRules{
				{
					Condition: &alerts.LogsRatioCondition{
						Threshold: &threshold,
						TimeWindow: &alerts.LogsRatioTimeWindow{
							LogsRatioTimeWindowSpecificValue: alerts.LOGSRATIOTIMEWINDOWVALUE_LOGS_RATIO_TIME_WINDOW_VALUE_MINUTES_5_OR_UNSPECIFIED.Ptr(),
						},
						ConditionType: alerts.LOGSRATIOCONDITIONTYPE_LOGS_RATIO_CONDITION_TYPE_LESS_THAN.Ptr(),
					},
					Override: &alerts.AlertDefOverride{
						Priority: alerts.ALERTDEFPRIORITY_ALERT_DEF_PRIORITY_P2.Ptr(),
					},
				},
			},
		},
	}
}

func CreateTracingThresholdAlert() *alerts.AlertDefPropertiesTracingThreshold {
	name := "Standard alert example"
	startHour := int32(8)
	startMinute := int32(30)
	endHour := int32(20)
	endMinute := int32(30)
	latencyThresholdMs := "30"
	spanAmount := float64(5)
	return &alerts.AlertDefPropertiesTracingThreshold{
		Name:        &name,
		Description: alerts.PtrString("Standard alert example from OpenAPI SDK"),
		Enabled:     alerts.PtrBool(true),
		Priority:    alerts.ALERTDEFPRIORITY_ALERT_DEF_PRIORITY_P1.Ptr(),
		Type:        alerts.ALERTDEFTYPE_ALERT_DEF_TYPE_TRACING_THRESHOLD.Ptr(),
		EntityLabels: &map[string]string{
			"alert_type":        "security",
			"security_severity": "high",
		},
		PhantomMode: alerts.PtrBool(false),
		NotificationGroup: &alerts.AlertDefNotificationGroup{
			Webhooks: []alerts.AlertDefWebhooksSettings{
				{
					Minutes:  alerts.PtrInt64(5),
					NotifyOn: alerts.NOTIFYON_NOTIFY_ON_TRIGGERED_AND_RESOLVED.Ptr(),
					Integration: &alerts.V3IntegrationType{
						V3IntegrationTypeRecipients: &alerts.V3IntegrationTypeRecipients{
							Recipients: &alerts.Recipients{
								Emails: []string{"example@coralogix.com"},
							},
						},
					},
				},
			},
		},
		ActiveOn: &alerts.ActivitySchedule{
			DayOfWeek: []alerts.DayOfWeek{
				alerts.DAYOFWEEK_DAY_OF_WEEK_WEDNESDAY,
				alerts.DAYOFWEEK_DAY_OF_WEEK_THURSDAY,
			},
			StartTime: &alerts.TimeOfDay{Hours: &startHour, Minutes: &startMinute},
			EndTime:   &alerts.TimeOfDay{Hours: &endHour, Minutes: &endMinute},
		},
		TracingThreshold: &alerts.TracingThresholdType{
			TracingFilter: &alerts.TracingFilter{
				SimpleFilter: &alerts.TracingSimpleFilter{
					LatencyThresholdMs: &latencyThresholdMs,
					TracingLabelFilters: &alerts.TracingLabelFilters{
						ApplicationName: []alerts.TracingFilterType{
							{
								Operation: alerts.TRACINGFILTEROPERATIONTYPE_TRACING_FILTER_OPERATION_TYPE_IS_OR_UNSPECIFIED.Ptr(),
								Values:    []string{"nginx", "apache"},
							},
							{
								Operation: alerts.TRACINGFILTEROPERATIONTYPE_TRACING_FILTER_OPERATION_TYPE_STARTS_WITH.Ptr(),
								Values:    []string{"application-name:"},
							},
						},
					},
				},
			},
			Rules: []alerts.TracingThresholdRule{
				{
					Condition: &alerts.TracingThresholdCondition{
						TimeWindow: &alerts.TracingTimeWindow{
							TracingTimeWindowValue: alerts.TRACINGTIMEWINDOWVALUE_TRACING_TIME_WINDOW_VALUE_MINUTES_10.Ptr(),
						},
						SpanAmount: &spanAmount,
					},
				},
			},
		},
	}
}

func CreateFlowAlert(alertId string) *alerts.AlertDefPropertiesFlow {
	name := "Standard alert example"
	startHour := int32(8)
	startMinute := int32(30)
	endHour := int32(20)
	endMinute := int32(30)
	enforceSuppresion := false
	timeframeMs := "10"
	return &alerts.AlertDefPropertiesFlow{
		Name:        &name,
		Description: alerts.PtrString("Standard alert example from OpenAPI SDK"),
		Enabled:     alerts.PtrBool(true),
		Priority:    alerts.ALERTDEFPRIORITY_ALERT_DEF_PRIORITY_P1.Ptr(),
		Type:        alerts.ALERTDEFTYPE_ALERT_DEF_TYPE_FLOW.Ptr(),
		EntityLabels: &map[string]string{
			"alert_type":        "security",
			"security_severity": "high",
		},
		PhantomMode: alerts.PtrBool(false),
		NotificationGroup: &alerts.AlertDefNotificationGroup{
			Webhooks: []alerts.AlertDefWebhooksSettings{
				{
					Minutes:  alerts.PtrInt64(5),
					NotifyOn: alerts.NOTIFYON_NOTIFY_ON_TRIGGERED_AND_RESOLVED.Ptr(),
					Integration: &alerts.V3IntegrationType{
						V3IntegrationTypeRecipients: &alerts.V3IntegrationTypeRecipients{
							Recipients: &alerts.Recipients{
								Emails: []string{"example@coralogix.com"},
							},
						},
					},
				},
			},
		},
		ActiveOn: &alerts.ActivitySchedule{
			DayOfWeek: []alerts.DayOfWeek{
				alerts.DAYOFWEEK_DAY_OF_WEEK_WEDNESDAY,
				alerts.DAYOFWEEK_DAY_OF_WEEK_THURSDAY,
			},
			StartTime: &alerts.TimeOfDay{Hours: &startHour, Minutes: &startMinute},
			EndTime:   &alerts.TimeOfDay{Hours: &endHour, Minutes: &endMinute},
		},
		Flow: &alerts.FlowType{
			EnforceSuppression: &enforceSuppresion,
			Stages: []alerts.FlowStages{
				{
					FlowStagesGroups: &alerts.FlowStagesGroups{
						Groups: []alerts.FlowStagesGroup{
							{
								AlertDefs: []alerts.FlowStagesGroupsAlertDefs{
									{
										Id: &alertId,
									},
								},
							},
						},
					},
					TimeframeMs:   &timeframeMs,
					TimeframeType: alerts.TIMEFRAMETYPE_TIMEFRAME_TYPE_UP_TO.Ptr(),
				},
			},
		},
	}
}

func TestTracingImmediateAlerts(t *testing.T) {
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)
	client := cxsdk.NewAlertsClient(cpc)
	ctx := context.Background()

	createReq := alerts.AlertDefsServiceCreateAlertDefRequest{
		AlertDefPropertiesTracingImmediate: CreateTracingImmediateAlert(),
	}
	created, httpResp, err := client.
		AlertDefsServiceCreateAlertDef(ctx).
		AlertDefsServiceCreateAlertDefRequest(createReq).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.NotEmpty(t, created.AlertDef.Id)

	retrieved, httpResp, err := client.
		AlertDefsServiceGetAlertDef(ctx, *created.AlertDef.Id).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.Equal(t, created.AlertDef.Id, retrieved.AlertDef.Id)

	newDesc := "Updated description via OpenAPI SDK"
	updateReq := alerts.ReplaceAlertDefinitionRequest{
		Id: created.AlertDef.Id,
		AlertDefProperties: &alerts.AlertDefProperties{
			AlertDefPropertiesTracingImmediate: &alerts.AlertDefPropertiesTracingImmediate{
				Name:             created.AlertDef.AlertDefProperties.AlertDefPropertiesTracingImmediate.Name,
				Description:      &newDesc,
				Enabled:          created.AlertDef.AlertDefProperties.AlertDefPropertiesTracingImmediate.Enabled,
				Priority:         created.AlertDef.AlertDefProperties.AlertDefPropertiesTracingImmediate.Priority,
				Type:             created.AlertDef.AlertDefProperties.AlertDefPropertiesTracingImmediate.Type,
				EntityLabels:     created.AlertDef.AlertDefProperties.AlertDefPropertiesTracingImmediate.EntityLabels,
				GroupByKeys:      created.AlertDef.AlertDefProperties.AlertDefPropertiesTracingImmediate.GroupByKeys,
				PhantomMode:      created.AlertDef.AlertDefProperties.AlertDefPropertiesTracingImmediate.PhantomMode,
				TracingImmediate: created.AlertDef.AlertDefProperties.AlertDefPropertiesTracingImmediate.TracingImmediate,
			},
		},
	}
	updated, httpResp, err := client.
		AlertDefsServiceReplaceAlertDef(ctx).
		ReplaceAlertDefinitionRequest(updateReq).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.Equal(t, *updated.AlertDef.AlertDefProperties.AlertDefPropertiesTracingImmediate.Description, newDesc)

	_, httpResp, err = client.
		AlertDefsServiceDeleteAlertDef(ctx, *updated.AlertDef.Id).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))

	_, httpResp, err = client.
		AlertDefsServiceGetAlertDef(ctx, *updated.AlertDef.Id).
		Execute()
	assert.NotNil(t, cxsdk.NewAPIError(httpResp, err))
}

func TestLogsRatioAlerts(t *testing.T) {
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)
	client := cxsdk.NewAlertsClient(cpc)
	ctx := context.Background()

	createReq := alerts.AlertDefsServiceCreateAlertDefRequest{
		AlertDefPropertiesLogsRatioThreshold: CreateLogsRatioAlert(),
	}
	created, httpResp, err := client.
		AlertDefsServiceCreateAlertDef(ctx).
		AlertDefsServiceCreateAlertDefRequest(createReq).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.NotEmpty(t, created.AlertDef.Id)

	retrieved, httpResp, err := client.
		AlertDefsServiceGetAlertDef(ctx, *created.AlertDef.Id).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.Equal(t, created.AlertDef.Id, retrieved.AlertDef.Id)

	newDesc := "Updated description via OpenAPI SDK"
	updateReq := alerts.ReplaceAlertDefinitionRequest{
		Id: created.AlertDef.Id,
		AlertDefProperties: &alerts.AlertDefProperties{
			AlertDefPropertiesLogsRatioThreshold: &alerts.AlertDefPropertiesLogsRatioThreshold{
				Name:               created.AlertDef.AlertDefProperties.AlertDefPropertiesLogsRatioThreshold.Name,
				Description:        &newDesc,
				Enabled:            created.AlertDef.AlertDefProperties.AlertDefPropertiesLogsRatioThreshold.Enabled,
				Priority:           created.AlertDef.AlertDefProperties.AlertDefPropertiesLogsRatioThreshold.Priority,
				Type:               created.AlertDef.AlertDefProperties.AlertDefPropertiesLogsRatioThreshold.Type,
				EntityLabels:       created.AlertDef.AlertDefProperties.AlertDefPropertiesLogsRatioThreshold.EntityLabels,
				GroupByKeys:        created.AlertDef.AlertDefProperties.AlertDefPropertiesLogsRatioThreshold.GroupByKeys,
				PhantomMode:        created.AlertDef.AlertDefProperties.AlertDefPropertiesLogsRatioThreshold.PhantomMode,
				LogsRatioThreshold: created.AlertDef.AlertDefProperties.AlertDefPropertiesLogsRatioThreshold.LogsRatioThreshold,
			},
		},
	}
	updated, httpResp, err := client.
		AlertDefsServiceReplaceAlertDef(ctx).
		ReplaceAlertDefinitionRequest(updateReq).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.Equal(t, *updated.AlertDef.AlertDefProperties.AlertDefPropertiesLogsRatioThreshold.Description, newDesc)

	_, httpResp, err = client.
		AlertDefsServiceDeleteAlertDef(ctx, *updated.AlertDef.Id).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))

	_, httpResp, err = client.
		AlertDefsServiceGetAlertDef(ctx, *updated.AlertDef.Id).
		Execute()
	assert.NotNil(t, cxsdk.NewAPIError(httpResp, err))
}

func TestTracingThresholdAlerts(t *testing.T) {
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)
	client := cxsdk.NewAlertsClient(cpc)
	ctx := context.Background()

	createReq := alerts.AlertDefsServiceCreateAlertDefRequest{
		AlertDefPropertiesTracingThreshold: CreateTracingThresholdAlert(),
	}
	created, httpResp, err := client.
		AlertDefsServiceCreateAlertDef(ctx).
		AlertDefsServiceCreateAlertDefRequest(createReq).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.NotEmpty(t, created.AlertDef.Id)

	retrieved, httpResp, err := client.
		AlertDefsServiceGetAlertDef(ctx, *created.AlertDef.Id).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.Equal(t, created.AlertDef.Id, retrieved.AlertDef.Id)

	newDesc := "Updated description via OpenAPI SDK"
	updateReq := alerts.ReplaceAlertDefinitionRequest{
		Id: created.AlertDef.Id,
		AlertDefProperties: &alerts.AlertDefProperties{
			AlertDefPropertiesTracingThreshold: &alerts.AlertDefPropertiesTracingThreshold{
				Name:             created.AlertDef.AlertDefProperties.AlertDefPropertiesTracingThreshold.Name,
				Description:      &newDesc,
				Enabled:          created.AlertDef.AlertDefProperties.AlertDefPropertiesTracingThreshold.Enabled,
				Priority:         created.AlertDef.AlertDefProperties.AlertDefPropertiesTracingThreshold.Priority,
				Type:             created.AlertDef.AlertDefProperties.AlertDefPropertiesTracingThreshold.Type,
				EntityLabels:     created.AlertDef.AlertDefProperties.AlertDefPropertiesTracingThreshold.EntityLabels,
				GroupByKeys:      created.AlertDef.AlertDefProperties.AlertDefPropertiesTracingThreshold.GroupByKeys,
				PhantomMode:      created.AlertDef.AlertDefProperties.AlertDefPropertiesTracingThreshold.PhantomMode,
				TracingThreshold: created.AlertDef.AlertDefProperties.AlertDefPropertiesTracingThreshold.TracingThreshold,
			},
		},
	}
	updated, httpResp, err := client.
		AlertDefsServiceReplaceAlertDef(ctx).
		ReplaceAlertDefinitionRequest(updateReq).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.Equal(t, *updated.AlertDef.AlertDefProperties.AlertDefPropertiesTracingThreshold.Description, newDesc)

	_, httpResp, err = client.
		AlertDefsServiceDeleteAlertDef(ctx, *updated.AlertDef.Id).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))

	_, httpResp, err = client.
		AlertDefsServiceGetAlertDef(ctx, *updated.AlertDef.Id).
		Execute()
	assert.NotNil(t, cxsdk.NewAPIError(httpResp, err))
}

func TestFlowAlerts(t *testing.T) {
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)
	client := cxsdk.NewAlertsClient(cpc)
	ctx := context.Background()

	// First create a tracing threshold alert to be used in the flow
	tracingAlertReq := alerts.AlertDefsServiceCreateAlertDefRequest{
		AlertDefPropertiesTracingThreshold: CreateTracingThresholdAlert(),
	}
	tracingAlert, httpResp, err := client.
		AlertDefsServiceCreateAlertDef(ctx).
		AlertDefsServiceCreateAlertDefRequest(tracingAlertReq).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.NotEmpty(t, tracingAlert.AlertDef.Id)

	// Now create the flow alert using the created tracing alert
	flowAlertReq := alerts.AlertDefsServiceCreateAlertDefRequest{
		AlertDefPropertiesFlow: CreateFlowAlert(*tracingAlert.AlertDef.Id),
	}
	flowAlert, httpResp, err := client.
		AlertDefsServiceCreateAlertDef(ctx).
		AlertDefsServiceCreateAlertDefRequest(flowAlertReq).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.NotEmpty(t, flowAlert.AlertDef.Id)

	retrieved, httpResp, err := client.
		AlertDefsServiceGetAlertDef(ctx, *flowAlert.AlertDef.Id).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.Equal(t, flowAlert.AlertDef.Id, retrieved.AlertDef.Id)

	newDesc := "Updated description via OpenAPI SDK"
	updateReq := alerts.ReplaceAlertDefinitionRequest{
		Id: flowAlert.AlertDef.Id,
		AlertDefProperties: &alerts.AlertDefProperties{
			AlertDefPropertiesFlow: &alerts.AlertDefPropertiesFlow{
				Name:         flowAlert.AlertDef.AlertDefProperties.AlertDefPropertiesFlow.Name,
				Description:  &newDesc,
				Enabled:      flowAlert.AlertDef.AlertDefProperties.AlertDefPropertiesFlow.Enabled,
				Priority:     flowAlert.AlertDef.AlertDefProperties.AlertDefPropertiesFlow.Priority,
				Type:         flowAlert.AlertDef.AlertDefProperties.AlertDefPropertiesFlow.Type,
				EntityLabels: flowAlert.AlertDef.AlertDefProperties.AlertDefPropertiesFlow.EntityLabels,
				GroupByKeys:  flowAlert.AlertDef.AlertDefProperties.AlertDefPropertiesFlow.GroupByKeys,
				PhantomMode:  flowAlert.AlertDef.AlertDefProperties.AlertDefPropertiesFlow.PhantomMode,
				Flow:         flowAlert.AlertDef.AlertDefProperties.AlertDefPropertiesFlow.Flow,
			},
		},
	}
	updated, httpResp, err := client.
		AlertDefsServiceReplaceAlertDef(ctx).
		ReplaceAlertDefinitionRequest(updateReq).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.Equal(t, *updated.AlertDef.AlertDefProperties.AlertDefPropertiesFlow.Description, newDesc)

	_, httpResp, err = client.
		AlertDefsServiceDeleteAlertDef(ctx, *updated.AlertDef.Id).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))

	_, httpResp, err = client.
		AlertDefsServiceGetAlertDef(ctx, *updated.AlertDef.Id).
		Execute()
	assert.NotNil(t, cxsdk.NewAPIError(httpResp, err))

	// Clean up the tracing alert
	_, httpResp, err = client.
		AlertDefsServiceDeleteAlertDef(ctx, *tracingAlert.AlertDef.Id).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
}

func TestAlertScheduler(t *testing.T) {
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)

	ctx := context.Background()

	alertsClient := cxsdk.NewAlertsClient(cpc)
	createAlertReq := alerts.AlertDefsServiceCreateAlertDefRequest{
		AlertDefPropertiesTracingImmediate: CreateTracingImmediateAlert(),
	}

	createdAlert, httpResp, err := alertsClient.
		AlertDefsServiceCreateAlertDef(ctx).
		AlertDefsServiceCreateAlertDefRequest(createAlertReq).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.NotEmpty(t, createdAlert.AlertDef.Id)

	schedulerClient := cxsdk.NewAlertSchedulerClient(cpc)

	name := "Scheduler Test " + time.Now().Format(time.RFC3339)
	enabled := true

	schedule := scheduler.ScheduleOneTimeAsSchedule(&scheduler.ScheduleOneTime{
		ScheduleOperation: scheduler.SCHEDULEOPERATION_SCHEDULE_OPERATION_MUTE.Ptr(),
		OneTime: &scheduler.OneTime{
			Timeframe: &scheduler.Timeframe{
				TimeframeEndTime: &scheduler.TimeframeEndTime{
					StartTime: scheduler.PtrString("2025-01-01T00:00:00.000Z"),
					EndTime:   scheduler.PtrString("2025-12-31T23:59:59.000Z"),
					Timezone:  scheduler.PtrString("UTC"),
				},
			},
		},
	})

	filter := scheduler.AlertSchedulerRuleProtobufV1Filter{
		AlertSchedulerRuleProtobufV1FilterAlertUniqueIds: &scheduler.AlertSchedulerRuleProtobufV1FilterAlertUniqueIds{
			WhatExpression: scheduler.PtrString("source logs | filter $d.cpodId:string == '122'"),
			AlertUniqueIds: &scheduler.AlertUniqueIds{
				Value: []string{*createdAlert.AlertDef.Id},
			},
		},
	}

	createReq := scheduler.NewCreateAlertSchedulerRuleRequestDataStructure(
		scheduler.AlertSchedulerRule{
			Name:        &name,
			Description: scheduler.PtrString("Created via OpenAPI SDK"),
			Enabled:     &enabled,
			Filter:      &filter,
			Schedule:    &schedule,
		},
	)

	createResp, httpResp, err := schedulerClient.
		AlertSchedulerRuleServiceCreateAlertSchedulerRule(ctx).
		CreateAlertSchedulerRuleRequestDataStructure(*createReq).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.NotNil(t, createResp.AlertSchedulerRule.UniqueIdentifier)

	ruleID := *createResp.AlertSchedulerRule.UniqueIdentifier

	updatedName := "Updated Scheduler " + time.Now().Format(time.RFC3339)
	updateReq := scheduler.UpdateAlertSchedulerRuleRequestDataStructure{
		AlertSchedulerRule: scheduler.AlertSchedulerRule{
			UniqueIdentifier: &ruleID,
			Name:             &updatedName,
			Description:      scheduler.PtrString("Updated description via OpenAPI SDK"),
			Enabled:          &enabled,
			Filter:           &filter,
			Schedule:         &schedule,
		},
	}

	updateResp, httpResp, err := schedulerClient.
		AlertSchedulerRuleServiceUpdateAlertSchedulerRule(ctx).
		UpdateAlertSchedulerRuleRequestDataStructure(updateReq).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.Equal(t, updatedName, *updateResp.AlertSchedulerRule.Name)

	getResp, httpResp, err := schedulerClient.
		AlertSchedulerRuleServiceGetAlertSchedulerRule(ctx, ruleID).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.Equal(t, updatedName, *getResp.AlertSchedulerRule.Name)

	_, httpResp, err = schedulerClient.
		AlertSchedulerRuleServiceDeleteAlertSchedulerRule(ctx, ruleID).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))

	_, httpResp, err = alertsClient.
		AlertDefsServiceDeleteAlertDef(ctx, *createdAlert.AlertDef.Id).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
}
