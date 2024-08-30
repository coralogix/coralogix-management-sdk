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

	flow "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/alerts/v3/alert_def_type_definition/flow"
	metric "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/alerts/v3/alert_def_type_definition/metric"
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

// AlertEvaluationWindowRollingOrUnspecified is a type of evaluation window.
const AlertEvaluationWindowRollingOrUnspecified = standard.EvaluationWindow_EVALUATION_WINDOW_ROLLING_OR_UNSPECIFIED

// AlertEvaluationWindowDynamic is a type of evaluation window.
const AlertEvaluationWindowDynamic = standard.EvaluationWindow_EVALUATION_WINDOW_DYNAMIC

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

// AlertDayOfWeekUnspecified is unspecified.
const AlertDayOfWeekUnspecified = alerts.DayOfWeek_DAY_OF_WEEK_MONDAY_OR_UNSPECIFIED

// AlertDayOfWeekMonday is Monday.
const AlertDayOfWeekMonday = alerts.DayOfWeek_DAY_OF_WEEK_MONDAY_OR_UNSPECIFIED

// AlertDayOfWeekTuesday is Tuesday.
const AlertDayOfWeekTuesday = alerts.DayOfWeek_DAY_OF_WEEK_TUESDAY

// AlertDayOfWeekWednesday is Wednesday.
const AlertDayOfWeekWednesday = alerts.DayOfWeek_DAY_OF_WEEK_WEDNESDAY

// AlertDayOfWeekThursday is Thursday.
const AlertDayOfWeekThursday = alerts.DayOfWeek_DAY_OF_WEEK_THURSDAY

// AlertDayOfWeekFriday is Friday.
const AlertDayOfWeekFriday = alerts.DayOfWeek_DAY_OF_WEEK_FRIDAY

// AlertDayOfWeekSaturday is Saturday.
const AlertDayOfWeekSaturday = alerts.DayOfWeek_DAY_OF_WEEK_SATURDAY

// AlertDayOfWeekSunday is Sunday.
const AlertDayOfWeekSunday = alerts.DayOfWeek_DAY_OF_WEEK_SUNDAY

// CreateAlertDefRpc is the name of the respective RPC.
const CreateAlertDefRpc = alerts.AlertDefsService_CreateAlertDef_FullMethodName

// DeleteAlertDefRpc is the name of the respective RPC.
const DeleteAlertDefRpc = alerts.AlertDefsService_DeleteAlertDef_FullMethodName

// GetAlertDefRpc is the name of the respective RPC.
const GetAlertDefRpc = alerts.AlertDefsService_GetAlertDef_FullMethodName

// ListAlertDefsRpc is the name of the respective RPC.
const ListAlertDefsRpc = alerts.AlertDefsService_ListAlertDefs_FullMethodName

// ReplaceAlertDefRpc is the name of the respective RPC.
const ReplaceAlertDefRpc = alerts.AlertDefsService_ReplaceAlertDef_FullMethodName

// SetAlertDefActiveRpc is the name of the respective RPC.
const SetAlertDefActiveRpc = alerts.AlertDefsService_SetActive_FullMethodName

// AlertNotifyOnTriggeredOnlyUnspecified is a notification setting for an alert.
const AlertNotifyOnTriggeredOnlyUnspecified = alerts.NotifyOn_NOTIFY_ON_TRIGGERED_ONLY_UNSPECIFIED

// AlertNotifyOnTriggeredAndResolved is a notification setting for an alert.
const AlertNotifyOnTriggeredAndResolved = alerts.NotifyOn_NOTIFY_ON_TRIGGERED_AND_RESOLVED

// LogFilterOperationIsOrUnspecified is a filter operation for an alert.
const LogFilterOperationIsOrUnspecified = typedefs.LogFilterOperationType_LOG_FILTER_OPERATION_TYPE_IS_OR_UNSPECIFIED

// LogFilterOperationIncludes is a filter operation for an alert.
const LogFilterOperationIncludes = typedefs.LogFilterOperationType_LOG_FILTER_OPERATION_TYPE_INCLUDES

// LogFilterOperationEndsWith is a filter operation for an alert.
const LogFilterOperationEndsWith = typedefs.LogFilterOperationType_LOG_FILTER_OPERATION_TYPE_ENDS_WITH

// LogFilterOperationStartsWith is a filter operation for an alert.
const LogFilterOperationStartsWith = typedefs.LogFilterOperationType_LOG_FILTER_OPERATION_TYPE_STARTS_WITH

// LogSeverityVerboseUnspecified is a log level for the logs triggering the alert.
const LogSeverityVerboseUnspecified = typedefs.LogSeverity_LOG_SEVERITY_VERBOSE_UNSPECIFIED

// LogSeverityDebug is a log level for the logs triggering the alert.
const LogSeverityDebug = typedefs.LogSeverity_LOG_SEVERITY_DEBUG

// LogSeverityInfo is a log level for the logs triggering the alert.
const LogSeverityInfo = typedefs.LogSeverity_LOG_SEVERITY_INFO

// LogSeverityWarning is a log level for the logs triggering the alert.
const LogSeverityWarning = typedefs.LogSeverity_LOG_SEVERITY_WARNING

// LogSeverityError is a log level for the logs triggering the alert.
const LogSeverityError = typedefs.LogSeverity_LOG_SEVERITY_ERROR

// LogSeverityCritical is a log level for the logs triggering the alert.
const LogSeverityCritical = typedefs.LogSeverity_LOG_SEVERITY_CRITICAL

// LogsTimeWindow5MinutesOrUnspecified is a time window setting for logs.
const LogsTimeWindow5MinutesOrUnspecified = standard.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_MINUTES_5_OR_UNSPECIFIED

// LogsTimeWindow10Minutes is a time window setting for logs.
const LogsTimeWindow10Minutes = standard.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_MINUTES_10

// LogsTimeWindow15Minutes is a time window setting for logs.
const LogsTimeWindow15Minutes = standard.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_MINUTES_15

// LogsTimeWindow30Minutes is a time window setting for logs.
const LogsTimeWindow30Minutes = standard.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_MINUTES_30

// LogsTimeWindow1Hour is a time window setting for logs.
const LogsTimeWindow1Hour = standard.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_HOUR_1

// LogsTimeWindow2Hour is a time window setting for logs.
const LogsTimeWindow2Hours = standard.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_HOURS_2

// LogsTimeWindow4Hour is a time window setting for logs.
const LogsTimeWindow4Hours = standard.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_HOURS_4

// LogsTimeWindow6Hour is a time window setting for logs.
const LogsTimeWindow6Hours = standard.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_HOURS_6

// LogsTimeWindow12Hours is a time window setting for logs.
const LogsTimeWindow12Hours = standard.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_HOURS_12

// LogsTimeWindow24Hours is a time window setting for logs.
const LogsTimeWindow24Hours = standard.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_HOURS_24

// LogsTimeWindow36Hours is a time window setting for logs.
const LogsTimeWindow36Hours = standard.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_HOURS_36

// LogsTimeRelativeComparedToPreviousHourOrUnspecified is a relative time setting for logs.
const LogsTimeRelativeComparedToPreviousHourOrUnspecified = timerelative.LogsTimeRelativeComparedTo_LOGS_TIME_RELATIVE_COMPARED_TO_PREVIOUS_HOUR_OR_UNSPECIFIED

// LogsTimeRelativeComparedToSameHourYesterday is a relative time setting for logs.
const LogsTimeRelativeComparedToSameHourYesterday = timerelative.LogsTimeRelativeComparedTo_LOGS_TIME_RELATIVE_COMPARED_TO_SAME_HOUR_YESTERDAY

// LogsTimeRelativeComparedToSameHourLastWeek is a relative time setting for logs.
const LogsTimeRelativeComparedToSameHourLastWeek = timerelative.LogsTimeRelativeComparedTo_LOGS_TIME_RELATIVE_COMPARED_TO_SAME_HOUR_LAST_WEEK

// LogsTimeRelativeComparedToYesterday is a relative time setting for logs.
const LogsTimeRelativeComparedToYesterday = timerelative.LogsTimeRelativeComparedTo_LOGS_TIME_RELATIVE_COMPARED_TO_YESTERDAY

// LogsTimeRelativeComparedToSameDayLastWeek is a relative time setting for logs.
const LogsTimeRelativeComparedToSameDayLastWeek = timerelative.LogsTimeRelativeComparedTo_LOGS_TIME_RELATIVE_COMPARED_TO_SAME_DAY_LAST_WEEK

// LogsTimeRelativeComparedToSameDayLastMonth is a relative time setting for logs.
const LogsTimeRelativeComparedToSameDayLastMonth = timerelative.LogsTimeRelativeComparedTo_LOGS_TIME_RELATIVE_COMPARED_TO_SAME_DAY_LAST_MONTH

// MetricTimeWindowValue1MinuteOrUnspecified is a time window setting for metrics.
const MetricTimeWindowValue1MinuteOrUnspecified = metric.MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_MINUTES_1_OR_UNSPECIFIED

// MetricTimeWindowValue5Minutes is a time window setting for metrics.
const MetricTimeWindowValue5Minutes = metric.MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_MINUTES_5

// MetricTimeWindowValue10Minutes is a time window setting for metrics.
const MetricTimeWindowValue10Minutes = metric.MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_MINUTES_10

// MetricTimeWindowValue15Minutes is a time window setting for metrics.
const MetricTimeWindowValue15Minutes = metric.MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_MINUTES_15

// MetricTimeWindowValue30Minutes is a time window setting for metrics.
const MetricTimeWindowValue30Minutes = metric.MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_MINUTES_30

// MetricTimeWindowValue1Hour is a time window setting for metrics.
const MetricTimeWindowValue1Hour = metric.MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_HOUR_1

// MetricTimeWindowValue2Hours is a time window setting for metrics.
const MetricTimeWindowValue2Hours = metric.MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_HOURS_2

// MetricTimeWindowValue4Hours is a time window setting for metrics.
const MetricTimeWindowValue4Hours = metric.MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_HOURS_4

// MetricTimeWindowValue6Hours is a time window setting for metrics.
const MetricTimeWindowValue6Hours = metric.MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_HOURS_6

// MetricTimeWindowValue12Hours is a time window setting for metrics.
const MetricTimeWindowValue12Hours = metric.MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_HOURS_12

// MetricTimeWindowValue24Hours is a time window setting for metrics.
const MetricTimeWindowValue24Hours = metric.MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_HOURS_24

// TracingTimeWindowValue5MinutesOrUnspecified is a time window setting for tracing.
const TracingTimeWindowValue5MinutesOrUnspecified = tracing.TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_MINUTES_5_OR_UNSPECIFIED

// TracingTimeWindowValue10Minutes is a time window setting for tracing.
const TracingTimeWindowValue10Minutes = tracing.TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_MINUTES_10

// TracingTimeWindowValue15Minutes is a time window setting for tracing.
const TracingTimeWindowValue15Minutes = tracing.TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_MINUTES_15

// TracingTimeWindowValue30Minutes is a time window setting for tracing.
const TracingTimeWindowValue30Minutes = tracing.TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_MINUTES_30

// TracingTimeWindowValue1Hour is a time window setting for tracing.
const TracingTimeWindowValue1Hour = tracing.TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_HOUR_1

// TracingTimeWindowValue2Hours is a time window setting for tracing.
const TracingTimeWindowValue2Hours = tracing.TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_HOURS_2

// TracingTimeWindowValue4Hours is a time window setting for tracing.
const TracingTimeWindowValue4Hours = tracing.TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_HOURS_4

// TracingTimeWindowValue6Hours is a time window setting for tracing.
const TracingTimeWindowValue6Hours = tracing.TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_HOURS_6

// TracingTimeWindowValue12Hours is a time window setting for tracing.
const TracingTimeWindowValue12Hours = tracing.TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_HOURS_12

// TracingTimeWindowValue24Hours is a time window setting for tracing.
const TracingTimeWindowValue24Hours = tracing.TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_HOURS_24

// TracingTimeWindowValue36Hours is a time window setting for tracing.
const TracingTimeWindowValue36Hours = tracing.TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_HOURS_36

// TracingFilterOperationTypeIsOrUnspecified is a filter operation for tracing.
const TracingFilterOperationTypeIsOrUnspecified = tracing.TracingFilterOperationType_TRACING_FILTER_OPERATION_TYPE_IS_OR_UNSPECIFIED

// TracingFilterOperationTypeIncludes is a filter operation for tracing.
const TracingFilterOperationTypeIncludes = tracing.TracingFilterOperationType_TRACING_FILTER_OPERATION_TYPE_INCLUDES

// TracingFilterOperationTypeEndsWith is a filter operation for tracing.
const TracingFilterOperationTypeEndsWith = tracing.TracingFilterOperationType_TRACING_FILTER_OPERATION_TYPE_ENDS_WITH

// TracingFilterOperationTypeStartsWith is a filter operation for tracing.
const TracingFilterOperationTypeStartsWith = tracing.TracingFilterOperationType_TRACING_FILTER_OPERATION_TYPE_STARTS_WITH

// TimeframeTypeUnspecified is a type of timeframe.
const TimeframeTypeUnspecified = flow.TimeframeType_TIMEFRAME_TYPE_UNSPECIFIED

// TimeframeTypeUpTo is a type of timeframe.
const TimeframeTypeUpTo = flow.TimeframeType_TIMEFRAME_TYPE_UP_TO

// NextOpAndOrUnspecified is a declaration for the next operation (AND).
const NextOpAndOrUnspecified = flow.NextOp_NEXT_OP_AND_OR_UNSPECIFIED

// NextOpOr is a declaration for the next operation (OR).
const NextOpOr = flow.NextOp_NEXT_OP_OR

// AutoRetireTimeframe values.
const (
	AutoRetireTimeframeNeverOrUnspecified = typedefs.AutoRetireTimeframe_AUTO_RETIRE_TIMEFRAME_NEVER_OR_UNSPECIFIED
	AutoRetireTimeframe5Minutes           = typedefs.AutoRetireTimeframe_AUTO_RETIRE_TIMEFRAME_MINUTES_5
	AutoRetireTimeframe10Minutes          = typedefs.AutoRetireTimeframe_AUTO_RETIRE_TIMEFRAME_MINUTES_10
	AutoRetireTimeframe1Hour              = typedefs.AutoRetireTimeframe_AUTO_RETIRE_TIMEFRAME_HOUR_1
	AutoRetireTimeframe2Hours             = typedefs.AutoRetireTimeframe_AUTO_RETIRE_TIMEFRAME_HOURS_2
	AutoRetireTimeframe6Hours             = typedefs.AutoRetireTimeframe_AUTO_RETIRE_TIMEFRAME_HOURS_6
	AutoRetireTimeframe12Hours            = typedefs.AutoRetireTimeframe_AUTO_RETIRE_TIMEFRAME_HOURS_12
	AutoRetireTimeframe24Hours            = typedefs.AutoRetireTimeframe_AUTO_RETIRE_TIMEFRAME_HOURS_24
)

// LogsRatioTimeWindow values.
const (
	LogsRatioTimeWindowValue5MinutesOrUnspecified = ratio.LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_MINUTES_5_OR_UNSPECIFIED
	LogsRatioTimeWindowValue10Minutes             = ratio.LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_MINUTES_10
	LogsRatioTimeWindowValue15Minutes             = ratio.LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_MINUTES_15
	LogsRatioTimeWindowValue30Minutes             = ratio.LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_MINUTES_30
	LogsRatioTimeWindowValue1Hour                 = ratio.LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_HOUR_1
	LogsRatioTimeWindowValue2Hours                = ratio.LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_HOURS_2
	LogsRatioTimeWindowValue4Hours                = ratio.LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_HOURS_4
	LogsRatioTimeWindowValue6Hours                = ratio.LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_HOURS_6
	LogsRatioTimeWindowValue12Hours               = ratio.LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_HOURS_12
	LogsRatioTimeWindowValue24Hours               = ratio.LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_HOURS_24
	LogsRatioTimeWindowValue36Hours               = ratio.LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_HOURS_36
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
func (a AlertsClient) Replace(ctx context.Context, req *alerts.ReplaceAlertDefRequest) (*alerts.ReplaceAlertDefResponse, error) {
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
func (a AlertsClient) Set(ctx context.Context, req *SetActiveRequest) (*alerts.SetActiveResponse, error) {
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
func (a AlertsClient) List(ctx context.Context, req *ListAlertDefsRequest) (*alerts.ListAlertDefsResponse, error) {
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
