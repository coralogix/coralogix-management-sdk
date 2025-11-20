package examples

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	alerts "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/alert_definitions_service"
	scheduler "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/alert_scheduler_rule_service"
	slos "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/slos_service"
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
	region, _ := cxsdk.URLFromRegion(cxsdk.RegionFromEnv())
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		region,
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewAlertsClient(cpc)
	ctx := context.Background()

	createReq := alerts.CreateAlertDefinitionRequest{
		AlertDefProperties: &alerts.AlertDefProperties{
			AlertDefPropertiesTracingImmediate: CreateTracingImmediateAlert(),
		},
	}
	created, httpResp, err := client.
		AlertDefsServiceCreateAlertDef(ctx).
		CreateAlertDefinitionRequest(createReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotEmpty(t, created.AlertDef.Id)

	retrieved, httpResp, err := client.
		AlertDefsServiceGetAlertDef(ctx, *created.AlertDef.Id).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, created.AlertDef.Id, retrieved.AlertDef.Id)

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
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, *updated.AlertDef.AlertDefProperties.AlertDefPropertiesTracingImmediate.Description, newDesc)

	_, httpResp, err = client.
		AlertDefsServiceDeleteAlertDef(ctx, *updated.AlertDef.Id).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	_, httpResp, err = client.
		AlertDefsServiceGetAlertDef(ctx, *updated.AlertDef.Id).
		Execute()
	require.NotNil(t, cxsdk.NewAPIError(httpResp, err))
}

func TestLogsRatioAlerts(t *testing.T) {
	region, _ := cxsdk.URLFromRegion(cxsdk.RegionFromEnv())
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		region,
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewAlertsClient(cpc)
	ctx := context.Background()

	createReq := alerts.CreateAlertDefinitionRequest{
		AlertDefProperties: &alerts.AlertDefProperties{
			AlertDefPropertiesLogsRatioThreshold: CreateLogsRatioAlert(),
		},
	}
	created, httpResp, err := client.
		AlertDefsServiceCreateAlertDef(ctx).
		CreateAlertDefinitionRequest(createReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotEmpty(t, created.AlertDef.Id)

	retrieved, httpResp, err := client.
		AlertDefsServiceGetAlertDef(ctx, *created.AlertDef.Id).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, created.AlertDef.Id, retrieved.AlertDef.Id)

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
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, *updated.AlertDef.AlertDefProperties.AlertDefPropertiesLogsRatioThreshold.Description, newDesc)

	_, httpResp, err = client.
		AlertDefsServiceDeleteAlertDef(ctx, *updated.AlertDef.Id).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	_, httpResp, err = client.
		AlertDefsServiceGetAlertDef(ctx, *updated.AlertDef.Id).
		Execute()
	require.NotNil(t, cxsdk.NewAPIError(httpResp, err))
}

func TestTracingThresholdAlerts(t *testing.T) {
	region, _ := cxsdk.URLFromRegion(cxsdk.RegionFromEnv())
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		region,
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewAlertsClient(cpc)
	ctx := context.Background()

	createReq := alerts.CreateAlertDefinitionRequest{
		AlertDefProperties: &alerts.AlertDefProperties{
			AlertDefPropertiesTracingThreshold: CreateTracingThresholdAlert(),
		},
	}
	created, httpResp, err := client.
		AlertDefsServiceCreateAlertDef(ctx).
		CreateAlertDefinitionRequest(createReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotEmpty(t, created.AlertDef.Id)

	retrieved, httpResp, err := client.
		AlertDefsServiceGetAlertDef(ctx, *created.AlertDef.Id).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, created.AlertDef.Id, retrieved.AlertDef.Id)

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
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, *updated.AlertDef.AlertDefProperties.AlertDefPropertiesTracingThreshold.Description, newDesc)

	_, httpResp, err = client.
		AlertDefsServiceDeleteAlertDef(ctx, *updated.AlertDef.Id).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	_, httpResp, err = client.
		AlertDefsServiceGetAlertDef(ctx, *updated.AlertDef.Id).
		Execute()
	require.NotNil(t, cxsdk.NewAPIError(httpResp, err))
}

func TestFlowAlerts(t *testing.T) {
	region, _ := cxsdk.URLFromRegion(cxsdk.RegionFromEnv())
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		region,
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewAlertsClient(cpc)
	ctx := context.Background()

	// First create a tracing threshold alert to be used in the flow
	tracingAlertReq := alerts.CreateAlertDefinitionRequest{
		AlertDefProperties: &alerts.AlertDefProperties{
			AlertDefPropertiesTracingThreshold: CreateTracingThresholdAlert(),
		},
	}
	tracingAlert, httpResp, err := client.
		AlertDefsServiceCreateAlertDef(ctx).
		CreateAlertDefinitionRequest(tracingAlertReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotEmpty(t, tracingAlert.AlertDef.Id)

	// Now create the flow alert using the created tracing alert
	flowAlertReq := alerts.CreateAlertDefinitionRequest{
		AlertDefProperties: &alerts.AlertDefProperties{
			AlertDefPropertiesFlow: CreateFlowAlert(*tracingAlert.AlertDef.Id),
		},
	}
	flowAlert, httpResp, err := client.
		AlertDefsServiceCreateAlertDef(ctx).
		CreateAlertDefinitionRequest(flowAlertReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotEmpty(t, flowAlert.AlertDef.Id)

	retrieved, httpResp, err := client.
		AlertDefsServiceGetAlertDef(ctx, *flowAlert.AlertDef.Id).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, flowAlert.AlertDef.Id, retrieved.AlertDef.Id)

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
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, *updated.AlertDef.AlertDefProperties.AlertDefPropertiesFlow.Description, newDesc)

	_, httpResp, err = client.
		AlertDefsServiceDeleteAlertDef(ctx, *updated.AlertDef.Id).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	_, httpResp, err = client.
		AlertDefsServiceGetAlertDef(ctx, *updated.AlertDef.Id).
		Execute()
	require.NotNil(t, cxsdk.NewAPIError(httpResp, err))

	// Clean up the tracing alert
	_, httpResp, err = client.
		AlertDefsServiceDeleteAlertDef(ctx, *tracingAlert.AlertDef.Id).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
}

func TestSloAlerts(t *testing.T) {
	ctx := context.Background()
	region, _ := cxsdk.URLFromRegion(cxsdk.RegionFromEnv())
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		region,
		cxsdk.APIKeyFromEnv(),
	)

	sloClient := cxsdk.NewClientSet(cpc).SLOs()
	sloPayload := getRequestBasedSlo("example_slo_for_alert")

	createSloReq := slos.SlosServiceCreateSloRequest{
		SloRequestBasedMetricSli: sloPayload,
	}
	sloResp, httpResp, err := sloClient.
		SlosServiceCreateSlo(ctx).
		SlosServiceCreateSloRequest(createSloReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	sloID := sloResp.GetSlo().SloRequestBasedMetricSli.GetId()
	require.NotEmpty(t, sloID)

	alertsClient := cxsdk.NewAlertsClient(cpc)
	createAlertReq := alerts.CreateAlertDefinitionRequest{
		AlertDefProperties: &alerts.AlertDefProperties{
			AlertDefPropertiesSloThreshold: CreateBurnRateSloAlert(sloID),
		},
	}

	created, httpResp, err := alertsClient.
		AlertDefsServiceCreateAlertDef(ctx).
		CreateAlertDefinitionRequest(createAlertReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotEmpty(t, created.AlertDef.Id)

	retrieved, httpResp, err := alertsClient.
		AlertDefsServiceGetAlertDef(ctx, *created.AlertDef.Id).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, created.AlertDef.Id, retrieved.AlertDef.Id)

	newDesc := "Updated SLO alert description via OpenAPI SDK"
	updateReq := alerts.ReplaceAlertDefinitionRequest{
		Id: created.AlertDef.Id,
		AlertDefProperties: &alerts.AlertDefProperties{
			AlertDefPropertiesSloThreshold: &alerts.AlertDefPropertiesSloThreshold{
				Name:         created.AlertDef.AlertDefProperties.AlertDefPropertiesSloThreshold.Name,
				Description:  &newDesc,
				Enabled:      created.AlertDef.AlertDefProperties.AlertDefPropertiesSloThreshold.Enabled,
				Priority:     created.AlertDef.AlertDefProperties.AlertDefPropertiesSloThreshold.Priority,
				Type:         created.AlertDef.AlertDefProperties.AlertDefPropertiesSloThreshold.Type,
				EntityLabels: created.AlertDef.AlertDefProperties.AlertDefPropertiesSloThreshold.EntityLabels,
				GroupByKeys:  created.AlertDef.AlertDefProperties.AlertDefPropertiesSloThreshold.GroupByKeys,
				PhantomMode:  created.AlertDef.AlertDefProperties.AlertDefPropertiesSloThreshold.PhantomMode,
				SloThreshold: created.AlertDef.AlertDefProperties.AlertDefPropertiesSloThreshold.SloThreshold,
			},
		},
	}
	updated, httpResp, err := alertsClient.
		AlertDefsServiceReplaceAlertDef(ctx).
		ReplaceAlertDefinitionRequest(updateReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, *updated.AlertDef.AlertDefProperties.AlertDefPropertiesSloThreshold.Description, newDesc)

	_, httpResp, err = alertsClient.
		AlertDefsServiceDeleteAlertDef(ctx, *updated.AlertDef.Id).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	_, httpResp, err = alertsClient.
		AlertDefsServiceGetAlertDef(ctx, *updated.AlertDef.Id).
		Execute()
	require.NotNil(t, cxsdk.NewAPIError(httpResp, err))

	_, httpResp, err = sloClient.
		SlosServiceDeleteSlo(ctx, sloID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
}

func TestAlertScheduler(t *testing.T) {
	region, _ := cxsdk.URLFromRegion(cxsdk.RegionFromEnv())
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		region,
		cxsdk.APIKeyFromEnv(),
	)

	ctx := context.Background()

	alertsClient := cxsdk.NewAlertsClient(cpc)
	createAlertReq := alerts.CreateAlertDefinitionRequest{
		AlertDefProperties: &alerts.AlertDefProperties{
			AlertDefPropertiesTracingImmediate: CreateTracingImmediateAlert(),
		},
	}

	createdAlert, httpResp, err := alertsClient.
		AlertDefsServiceCreateAlertDef(ctx).
		CreateAlertDefinitionRequest(createAlertReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotEmpty(t, createdAlert.AlertDef.Id)

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
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotNil(t, createResp.AlertSchedulerRule.UniqueIdentifier)

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
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, updatedName, *updateResp.AlertSchedulerRule.Name)

	getResp, httpResp, err := schedulerClient.
		AlertSchedulerRuleServiceGetAlertSchedulerRule(ctx, ruleID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, updatedName, *getResp.AlertSchedulerRule.Name)

	_, httpResp, err = schedulerClient.
		AlertSchedulerRuleServiceDeleteAlertSchedulerRule(ctx, ruleID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	_, httpResp, err = alertsClient.
		AlertDefsServiceDeleteAlertDef(ctx, *createdAlert.AlertDef.Id).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
}

// CreateSloAlert returns a reusable Burn Rate SLO alert definition payload
func CreateBurnRateSloAlert(sloID string) *alerts.AlertDefPropertiesSloThreshold {
	return &alerts.AlertDefPropertiesSloThreshold{
		Name:        slos.PtrString("SLO Burn Rate Alert Example"),
		Description: alerts.PtrString("SLO burn rate alert created via OpenAPI SDK"),
		Enabled:     alerts.PtrBool(true),
		Priority:    alerts.ALERTDEFPRIORITY_ALERT_DEF_PRIORITY_P1.Ptr(),
		Type:        alerts.ALERTDEFTYPE_ALERT_DEF_TYPE_SLO_THRESHOLD.Ptr(),
		EntityLabels: &map[string]string{
			"alert_type": "slo",
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
			StartTime: &alerts.TimeOfDay{Hours: slos.PtrInt32(8), Minutes: slos.PtrInt32(32)},
			EndTime:   &alerts.TimeOfDay{Hours: slos.PtrInt32(20), Minutes: slos.PtrInt32(30)},
		},
		SloThreshold: &alerts.SloThresholdType{
			SloThresholdTypeBurnRate: &alerts.SloThresholdTypeBurnRate{
				SloDefinition: &alerts.V3SloDefinition{
					SloId: &sloID,
				},
				BurnRate: &alerts.BurnRateThreshold{
					BurnRateThresholdSingle: &alerts.BurnRateThresholdSingle{
						Rules: []alerts.SloThresholdRule{
							{
								Condition: &alerts.SloThresholdCondition{
									Threshold: slos.PtrFloat64(1),
								},
								Override: &alerts.AlertDefOverride{
									Priority: alerts.ALERTDEFPRIORITY_ALERT_DEF_PRIORITY_P1.Ptr(),
								},
							},
						},
						Single: &alerts.BurnRateTypeSingle{
							TimeDuration: &alerts.TimeDuration{
								Duration: slos.PtrString("1"),
								Unit:     alerts.DURATIONUNIT_DURATION_UNIT_HOURS.Ptr(),
							},
						},
					},
				},
			},
		},
	}
}
