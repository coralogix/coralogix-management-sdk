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

package cxsdk

import (
	"context"

	alerts "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/alerts/v3"
	typedefs "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/alerts/v3/alert_def_type_definition"
	ratio "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/alerts/v3/alert_def_type_definition/ratio"

	metric "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/alerts/v3/alert_def_type_definition/metric"
	newValue "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/alerts/v3/alert_def_type_definition/new_value"
	standard "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/alerts/v3/alert_def_type_definition/standard"
	timerelative "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/alerts/v3/alert_def_type_definition/time_relative"
	tracing "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/alerts/v3/alert_def_type_definition/tracing"
)

// AlertDefNotificationGroup represents a notification group.
type AlertDefNotificationGroup = alerts.AlertDefNotificationGroup

// AlertDefNotificationGroupSimple represents a simple notification group.
type AlertDefNotificationGroupSimple = alerts.AlertDefNotificationGroup_Simple

// AlertDefNotificationGroupAdvanced represents an advanced notification group.
type AlertDefNotificationGroupAdvanced = alerts.AlertDefNotificationGroup_Advanced

// AlertDefTargetSimple represents a simple target.
type AlertDefTargetSimple = alerts.AlertDefTargetSimple

// AlertDefAdvancedTargetSettings represents a simple target.
type AlertDefAdvancedTargetSettings = alerts.AlertDefAdvancedTargetSettings

// AlertDefIntegrationType represents an integration type.
type AlertDefIntegrationType = alerts.IntegrationType

// AlertDefIntegrationTypeIntegrationID represents an integration of type integration ID.
type AlertDefIntegrationTypeIntegrationID = alerts.IntegrationType_IntegrationId

// AlertDefIntegrationTypeRecipients represents an integration of type recipients.
type AlertDefIntegrationTypeRecipients = alerts.IntegrationType_Recipients

// AlertDefRecipients represents the recipients of a notification.
type AlertDefRecipients = alerts.Recipients

// AlertDefAdvancedTargets represents a set of advanced targets.
type AlertDefAdvancedTargets = alerts.AlertDefAdvancedTargets

// AlertDefAdvancedTargetSettingsMinutes represents the minute settings of an advanced target.
type AlertDefAdvancedTargetSettingsMinutes = alerts.AlertDefAdvancedTargetSettings_Minutes

// AlertDefScheduleActiveOn represents the active on schedule of an alert.
type AlertDefScheduleActiveOn = alerts.AlertDefProperties_ActiveOn

// AlertDefActivitySchedule is a schedule for an alert.
type AlertDefActivitySchedule = alerts.ActivitySchedule

// CreateAlertDefRequest is a request to create an alert.
type CreateAlertDefRequest = alerts.CreateAlertDefRequest

// CreateAlertDefRequest is a request to create an alert.
type GetAlertDefRequest = alerts.GetAlertDefRequest

// CreateAlertDefRequest is a request to create an alert.
type ReplaceAlertDefRequest = alerts.ReplaceAlertDefRequest

// CreateAlertDefRequest is a request to create an alert.
type DeleteAlertDefRequest = alerts.DeleteAlertDefRequest

// AlertDefProperties is the properties of an alert.
type AlertDefProperties = alerts.AlertDefProperties

// AlertDefPropertiesLogsMoreThan is a property of an alert.
type AlertDefPropertiesLogsMoreThan = alerts.AlertDefProperties_LogsMoreThan

// AlertDefPropertiesLogsLessThan is a property of an alert.
type AlertDefPropertiesLogsLessThan = alerts.AlertDefProperties_LogsLessThan

// AlertDefPropertiesLogsMoreThanUsual is a property of an alert.
type AlertDefPropertiesLogsMoreThanUsual = alerts.AlertDefProperties_LogsMoreThanUsual

// AlertDefPropertiesLogsNewValue is a property of an alert.
type AlertDefPropertiesLogsNewValue = alerts.AlertDefProperties_LogsNewValue

// AlertDefPropertiesLogsRatioLessThan is a property of an alert.
type AlertDefPropertiesLogsRatioLessThan = alerts.AlertDefProperties_LogsRatioLessThan

// AlertDefPropertiesLogsRatioMoreThan is a property of an alert.
type AlertDefPropertiesLogsRatioMoreThan = alerts.AlertDefProperties_LogsRatioMoreThan

// AlertDefPropertiesLogsTimeRelativeLessThan is a property of an alert.
type AlertDefPropertiesLogsTimeRelativeLessThan = alerts.AlertDefProperties_LogsTimeRelativeLessThan

// AlertDefPropertiesLogsTimeRelativeMoreThan is a property of an alert.
type AlertDefPropertiesLogsTimeRelativeMoreThan = alerts.AlertDefProperties_LogsTimeRelativeMoreThan

// AlertDefPropertiesLogsUniqueCount is a property of an alert.
type AlertDefPropertiesLogsUniqueCount = alerts.AlertDefProperties_LogsUniqueCount

// AlertDefPropertiesMetricLessThan is a property of an alert.
type AlertDefPropertiesMetricLessThan = alerts.AlertDefProperties_MetricLessThan

// AlertDefPropertiesMetricLessThanOrEquals is a property of an alert.
type AlertDefPropertiesMetricLessThanOrEquals = alerts.AlertDefProperties_MetricLessThanOrEquals

// AlertDefPropertiesMetricMoreThan is a property of an alert.
type AlertDefPropertiesMetricMoreThan = alerts.AlertDefProperties_MetricMoreThan

// AlertDefPropertiesMetricMoreThanOrEquals is a property of an alert.
type AlertDefPropertiesMetricMoreThanOrEquals = alerts.AlertDefProperties_MetricMoreThanOrEquals

// AlertDefPropertiesMetricLessThanUsual is a property of an alert.
type AlertDefPropertiesMetricLessThanUsual = alerts.AlertDefProperties_MetricLessThanUsual

// AlertDefPropertiesMetricMoreThanUsual is a property of an alert.
type AlertDefPropertiesMetricMoreThanUsual = alerts.AlertDefProperties_MetricMoreThanUsual

// AlertDefPropertiesTracingImmediate is a property of an alert.
type AlertDefPropertiesTracingImmediate = alerts.AlertDefProperties_TracingImmediate

// AlertDefPropertiesTracingMoreThan is a property of an alert.
type AlertDefPropertiesTracingMoreThan = alerts.AlertDefProperties_TracingMoreThan

// AlertTimeOfDay is a clock setting for an alert.
type AlertTimeOfDay = alerts.TimeOfDay

// SetActiveRequest is a request to set the active status of an alert.
type SetActiveRequest = alerts.SetActiveRequest

// ListAlertDefsRequest is a request to list alerts.
type ListAlertDefsRequest = alerts.ListAlertDefsRequest

// AlertsActivitySchedule is an activity schedule for Alerts.
type AlertsActivitySchedule = alerts.ActivitySchedule

// AlertDayOfWeek is a day enum used in Alerts.
type AlertDayOfWeek = alerts.DayOfWeek

// LogsFilter is a filter
type LogsFilter = typedefs.LogsFilter

// LogsFilterLuceneFilter is a filter type
type LogsFilterLuceneFilter = typedefs.LogsFilter_LuceneFilter

// LabelFilterType is a filter type
type LabelFilterType = typedefs.LabelFilterType

// LuceneFilter is a filter
type LuceneFilter = typedefs.LuceneFilter

// LabelFilters is a filter
type LabelFilters = typedefs.LabelFilters

// LogSeverity is a filter
type LogSeverity = typedefs.LogSeverity

// AlertDefIncidentSettings is the incident settings of an alert.
type AlertDefIncidentSettings = alerts.AlertDefIncidentSettings

// LogsMoreThanType is a logs filter type
type LogsMoreThanTypeDefinition = standard.LogsMoreThanTypeDefinition

// AlertDefIncidentSettingsMinutes is the incident settings of an alert.
type AlertDefIncidentSettingsMinutes = alerts.AlertDefIncidentSettings_Minutes

// LogsTimeWindowValue is a time window setting for logs.
type LogsTimeWindowValue = standard.LogsTimeWindowValue

// AlertNotifyOn is a trigger type.
type AlertNotifyOn = alerts.NotifyOn

// AlertEvaluationWindow is an evaluation window type for alerts.
type AlertEvaluationWindow = standard.EvaluationWindow

// LogFilterOperationType is a filter operation for logs.
type LogFilterOperationType = typedefs.LogFilterOperationType

// AlertDefPropertiesLogsImmediate is a property of an alert.
type AlertDefPropertiesLogsImmediate = alerts.AlertDefProperties_LogsImmediate

// AlertDefPropertiesLogsImmediateTypeDefinition is a type definition for logs immediate.
type AlertDefPropertiesLogsImmediateTypeDefinition = standard.LogsImmediateTypeDefinition

// LogsTimeRelativeComparedTo is a relative time setting for logs.
type LogsTimeRelativeComparedTo = timerelative.LogsTimeRelativeComparedTo

// MetricTimeWindowValue is a time window setting for metrics.
type MetricTimeWindowValue = metric.MetricTimeWindowValue

// TracingTimeWindowValue is a time window setting for tracing.
type TracingTimeWindowValue = tracing.TracingTimeWindowValue

// TracingFilterOperationType is an operation type setting for tracing.
type TracingFilterOperationType = tracing.TracingFilterOperationType

// AutoRetireTimeframe is a type of timeframe.
type AutoRetireTimeframe = typedefs.AutoRetireTimeframe

// LogsRatioTimeWindowValue is a time window setting for logs.
type LogsRatioTimeWindowValue = ratio.LogsRatioTimeWindowValue

// AlertDefPriority is the alert priority.
type AlertDefPriority = alerts.AlertDefPriority

// LogsRatioGroupBy is a group by setting for logs.
type LogsRatioGroupByFor = ratio.LogsRatioGroupByFor

// LogsNewValueTimeWindowValue is a time window setting for logs.
type LogsNewValueTimeWindowValue = newValue.LogsNewValueTimeWindowValue

// AlertDefPriorityP1 is the alert priority P1.
const AlertDefPriorityP1 = alerts.AlertDefPriority_ALERT_DEF_PRIORITY_P1

// AlertDefPriorityP2 is the alert priority P2.
const AlertDefPriorityP2 = alerts.AlertDefPriority_ALERT_DEF_PRIORITY_P2

// AlertDefPriorityP3 is the alert priority P3.
const AlertDefPriorityP3 = alerts.AlertDefPriority_ALERT_DEF_PRIORITY_P3

// AlertDefPriorityP4 is the alert priority P4.
const AlertDefPriorityP4 = alerts.AlertDefPriority_ALERT_DEF_PRIORITY_P4

// AlertDefPriorityP5OrUnspecified is the alert priority P5 or unspecified.
const AlertDefPriorityP5OrUnspecified = alerts.AlertDefPriority_ALERT_DEF_PRIORITY_P5_OR_UNSPECIFIED

// AlertDefTypeLogsMoreThan is the alert type "logs more than".
const AlertDefTypeLogsMoreThan = alerts.AlertDefType_ALERT_DEF_TYPE_LOGS_MORE_THAN

// AlertDefTypeLogsLessThan is the alert type "logs less than".
const AlertDefTypeLogsLessThan = alerts.AlertDefType_ALERT_DEF_TYPE_LOGS_LESS_THAN

// AlertDefTypeFlow is the alert type "flow".
const AlertDefTypeFlow = alerts.AlertDefType_ALERT_DEF_TYPE_FLOW

// AlertDefTypeLogsNewValue is the alert type "logs new value".
const AlertDefTypeLogsNewValue = alerts.AlertDefType_ALERT_DEF_TYPE_LOGS_NEW_VALUE

// AlertDefTypeLogsMoreThanUsual is the alert type "logs more than usual".
const AlertDefTypeLogsMoreThanUsual = alerts.AlertDefType_ALERT_DEF_TYPE_LOGS_MORE_THAN_USUAL

// AlertDefTypeLogsRatioMoreThan is the alert type "logs ratio more than".
const AlertDefTypeLogsRatioMoreThan = alerts.AlertDefType_ALERT_DEF_TYPE_LOGS_RATIO_MORE_THAN

// AlertDefTypeLogsRatioLessThan is the alert type "logs ratio less than".
const AlertDefTypeLogsRatioLessThan = alerts.AlertDefType_ALERT_DEF_TYPE_LOGS_RATIO_LESS_THAN

// AlertDefTypeLogsTimeRelativeLessThan is the alert type "logs time relative less than".
const AlertDefTypeLogsTimeRelativeLessThan = alerts.AlertDefType_ALERT_DEF_TYPE_LOGS_TIME_RELATIVE_LESS_THAN

// AlertDefTypeLogsTimeRelativeMoreThan is the alert type "logs time relative more than".
const AlertDefTypeLogsTimeRelativeMoreThan = alerts.AlertDefType_ALERT_DEF_TYPE_LOGS_TIME_RELATIVE_MORE_THAN

// AlertDefTypeLogsUniqueCount is the alert type "logs unique count".
const AlertDefTypeLogsUniqueCount = alerts.AlertDefType_ALERT_DEF_TYPE_LOGS_UNIQUE_COUNT

// AlertDefTypeLogsImmediateOrUnspecified is the alert type "logs immediate or unspecified".
const AlertDefTypeLogsImmediateOrUnspecified = alerts.AlertDefType_ALERT_DEF_TYPE_LOGS_IMMEDIATE_OR_UNSPECIFIED

// AlertDefTypeMetricMoreThan is the alert type "metric less than".
const AlertDefTypeMetricMoreThan = alerts.AlertDefType_ALERT_DEF_TYPE_METRIC_MORE_THAN

// AlertDefTypeMetricLessThan is the alert type "metric less than".
const AlertDefTypeMetricLessThan = alerts.AlertDefType_ALERT_DEF_TYPE_METRIC_LESS_THAN

// AlertDefTypeMetricLessThanOrEquals is the alert type "metric less than or equals".
const AlertDefTypeMetricLessThanOrEquals = alerts.AlertDefType_ALERT_DEF_TYPE_METRIC_LESS_THAN_OR_EQUALS

// AlertDefTypeMetricMoreThanOrEquals is the alert type "metric more than or equals".
const AlertDefTypeMetricMoreThanOrEquals = alerts.AlertDefType_ALERT_DEF_TYPE_METRIC_MORE_THAN_OR_EQUALS

// AlertDefTypeMetricMoreThanUsual is the alert type "metric more than usual".
const AlertDefTypeMetricMoreThanUsual = alerts.AlertDefType_ALERT_DEF_TYPE_METRIC_MORE_THAN_USUAL

// AlertDefTypeMetricLessThanUsual is the alert type "metric less than usual".
const AlertDefTypeMetricLessThanUsual = alerts.AlertDefType_ALERT_DEF_TYPE_METRIC_LESS_THAN_USUAL

// AlertDefTypeTracingImmediate is the alert type "tracing immediate".
const AlertDefTypeTracingImmediate = alerts.AlertDefType_ALERT_DEF_TYPE_TRACING_IMMEDIATE

// AlertDefTypeTracingMoreThan is the alert type "tracing more than".
const AlertDefTypeTracingMoreThan = alerts.AlertDefType_ALERT_DEF_TYPE_TRACING_MORE_THAN

// AlertDefSimpleNotificationGroup represents a simple notification group.
type AlertDefSimpleNotificationGroup = alerts.AlertDefNotificationGroup_Simple

// AlertDefAdvancedNotificationGroup represents an advanced notification group.
type AlertDefAdvancedNotificationGroup = alerts.AlertDefNotificationGroup_Advanced

// AlertDefSimpleTarget represents a simple target.
type AlertDefSimpleTarget = alerts.AlertDefTargetSimple

// CreateAlertRequest is a request to create an alert.
type CreateAlertRequest = alerts.CreateAlertDefRequest

// LogsRatioGroupByFor values.
const (
	LogsRatioGroupByForBothOrUnspecified = ratio.LogsRatioGroupByFor_LOGS_RATIO_GROUP_BY_FOR_BOTH_OR_UNSPECIFIED
	LogsRatioGroupByForNumeratorOnly     = ratio.LogsRatioGroupByFor_LOGS_RATIO_GROUP_BY_FOR_NUMERATOR_ONLY
	LogsRatioGroupByForDenumeratorOnly   = ratio.LogsRatioGroupByFor_LOGS_RATIO_GROUP_BY_FOR_DENUMERATOR_ONLY
)

// AlertsClient is a client for the Coralogix Alerts API.
type AlertsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// Create creates an alert.
func (a AlertsClient) Create(ctx context.Context, req *CreateAlertDefRequest) (*alerts.CreateAlertDefResponse, error) {
	callProperties, err := a.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := alerts.NewAlertDefsServiceClient(conn)

	return client.CreateAlertDef(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Get gets an alert.
func (a AlertsClient) Get(ctx context.Context, req *GetAlertDefRequest) (*alerts.GetAlertDefResponse, error) {
	callProperties, err := a.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := alerts.NewAlertDefsServiceClient(conn)

	return client.GetAlertDef(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Replace replaces an alert.
func (a AlertsClient) Replace(ctx context.Context, req *ReplaceAlertDefRequest) (*alerts.ReplaceAlertDefResponse, error) {
	callProperties, err := a.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := alerts.NewAlertDefsServiceClient(conn)

	return client.ReplaceAlertDef(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Delete deletes an alert.
func (a AlertsClient) Delete(ctx context.Context, req *DeleteAlertDefRequest) (*alerts.DeleteAlertDefResponse, error) {
	callProperties, err := a.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := alerts.NewAlertDefsServiceClient(conn)

	return client.DeleteAlertDef(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Set sets the active status of an alert.
func (a AlertsClient) Set(ctx context.Context, req *alerts.SetActiveRequest) (*alerts.SetActiveResponse, error) {
	callProperties, err := a.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := alerts.NewAlertDefsServiceClient(conn)

	return client.SetActive(callProperties.Ctx, req, callProperties.CallOptions...)
}

// List lists the alerts.
func (a AlertsClient) List(ctx context.Context, req *alerts.ListAlertDefsRequest) (*alerts.ListAlertDefsResponse, error) {
	callProperties, err := a.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := alerts.NewAlertDefsServiceClient(conn)

	return client.ListAlertDefs(callProperties.Ctx, req, callProperties.CallOptions...)
}

// NewAlertsClient creates a new alerts client.
func NewAlertsClient(c *CallPropertiesCreator) *AlertsClient {
	return &AlertsClient{callPropertiesCreator: c}
}
