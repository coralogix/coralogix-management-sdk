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

// CreateAlert returns a reusable OpenAPI alert definition payload
func CreateAlert() *alerts.AlertDefPropertiesLogsThreshold {
	return &alerts.AlertDefPropertiesLogsThreshold{
		Name:        "Standard alert example",
		Description: alerts.PtrString("Standard alert example from OpenAPI SDK"),
		Enabled:     alerts.PtrBool(true),
		Priority:    alerts.ALERTDEFPRIORITY_ALERT_DEF_PRIORITY_P1,
		Type:        alerts.ALERTDEFTYPE_ALERT_DEF_TYPE_LOGS_THRESHOLD,
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
					Integration: alerts.V3IntegrationType{
						IntegrationTypeRecipients: &alerts.IntegrationTypeRecipients{
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
			StartTime: alerts.TimeOfDay{Hours: 8, Minutes: 30},
			EndTime:   alerts.TimeOfDay{Hours: 20, Minutes: 30},
		},
		LogsThreshold: &alerts.LogsThresholdType{
			Rules: []alerts.LogsThresholdRule{
				{
					Override: alerts.AlertDefOverride{
						Priority: alerts.ALERTDEFPRIORITY_ALERT_DEF_PRIORITY_P1.Ptr(),
					},
					Condition: alerts.LogsThresholdCondition{
						ConditionType: alerts.LOGSTHRESHOLDCONDITIONTYPE_LOGS_THRESHOLD_CONDITION_TYPE_MORE_THAN_OR_UNSPECIFIED,
						Threshold:     10,
						TimeWindow: alerts.LogsTimeWindow{
							LogsTimeWindowSpecificValue: alerts.LOGSTIMEWINDOWVALUE_LOGS_TIME_WINDOW_VALUE_MINUTES_10.Ptr(),
						},
					},
				},
			},
			LogsFilter: &alerts.V3LogsFilter{
				SimpleFilter: &alerts.LogsSimpleFilter{
					LuceneQuery: alerts.PtrString("remote_addr_enriched:/.*/"),
					LabelFilters: &alerts.LabelFilters{
						ApplicationName: []alerts.LabelFilterType{
							{Value: "nginx", Operation: alerts.LOGFILTEROPERATIONTYPE_LOG_FILTER_OPERATION_TYPE_INCLUDES},
						},
						SubsystemName: []alerts.LabelFilterType{
							{Value: "subsystem-name", Operation: alerts.LOGFILTEROPERATIONTYPE_LOG_FILTER_OPERATION_TYPE_STARTS_WITH},
						},
						Severities: []alerts.LogSeverity{
							alerts.LOGSEVERITY_LOG_SEVERITY_WARNING,
							alerts.LOGSEVERITY_LOG_SEVERITY_ERROR,
						},
					},
				},
			},
		},
		GroupByKeys: []string{
			"coralogix.metadata.sdkId",
			"EventType",
		},
	}
}

func TestAlerts(t *testing.T) {
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)
	client := cxsdk.NewAlertsClient(cpc)
	ctx := context.Background()

	createReq := alerts.AlertDefsServiceCreateAlertDefRequest{
		AlertDefPropertiesLogsThreshold: CreateAlert(),
	}
	created, httpResp, err := client.
		AlertDefsServiceCreateAlertDef(ctx).
		AlertDefsServiceCreateAlertDefRequest(createReq).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.NotEmpty(t, created.AlertDef.Id)

	retrieved, httpResp, err := client.
		AlertDefsServiceGetAlertDef(ctx, created.AlertDef.Id).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.Equal(t, created.AlertDef.Id, retrieved.AlertDef.Id)

	newDesc := "Updated description via OpenAPI SDK"
	updateReq := alerts.ReplaceAlertDefinitionRequest{
		Id: created.AlertDef.Id,
		AlertDefProperties: alerts.AlertDefProperties{
			AlertDefPropertiesLogsThreshold: &alerts.AlertDefPropertiesLogsThreshold{
				Name:          created.AlertDef.AlertDefProperties.AlertDefPropertiesLogsThreshold.Name,
				Description:   &newDesc,
				Enabled:       created.AlertDef.AlertDefProperties.AlertDefPropertiesLogsThreshold.Enabled,
				Priority:      created.AlertDef.AlertDefProperties.AlertDefPropertiesLogsThreshold.Priority,
				Type:          created.AlertDef.AlertDefProperties.AlertDefPropertiesLogsThreshold.Type,
				EntityLabels:  created.AlertDef.AlertDefProperties.AlertDefPropertiesLogsThreshold.EntityLabels,
				GroupByKeys:   created.AlertDef.AlertDefProperties.AlertDefPropertiesLogsThreshold.GroupByKeys,
				PhantomMode:   created.AlertDef.AlertDefProperties.AlertDefPropertiesLogsThreshold.PhantomMode,
				LogsThreshold: created.AlertDef.AlertDefProperties.AlertDefPropertiesLogsThreshold.LogsThreshold,
			},
		},
	}
	updated, httpResp, err := client.
		AlertDefsServiceReplaceAlertDef(ctx).
		ReplaceAlertDefinitionRequest(updateReq).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.Equal(t, *updated.AlertDef.AlertDefProperties.AlertDefPropertiesLogsThreshold.Description, newDesc)

	_, httpResp, err = client.
		AlertDefsServiceDeleteAlertDef(ctx, updated.AlertDef.Id).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))

	_, httpResp, err = client.
		AlertDefsServiceGetAlertDef(ctx, updated.AlertDef.Id).
		Execute()
	assert.NotNil(t, cxsdk.NewAPIError(httpResp, err))
}

func TestAlertScheduler(t *testing.T) {
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)

	ctx := context.Background()

	alertsClient := cxsdk.NewAlertsClient(cpc)
	createAlertReq := alerts.AlertDefsServiceCreateAlertDefRequest{
		AlertDefPropertiesLogsThreshold: CreateAlert(),
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
		FilterAlertUniqueIds: &scheduler.FilterAlertUniqueIds{
			WhatExpression: scheduler.PtrString("source logs | filter $d.cpodId:string == '122'"),
			AlertUniqueIds: &scheduler.AlertUniqueIds{
				Value: []string{createdAlert.AlertDef.Id},
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
		AlertDefsServiceDeleteAlertDef(ctx, createdAlert.AlertDef.Id).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
}
