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
)

// AlertDef is an alert definition.
type AlertDef = alerts.AlertDef

// AlertDefPropertiesActiveOn is a type to set when an alert is active.
type AlertDefPropertiesActiveOn = alerts.AlertDefProperties_ActiveOn

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

// LogsFilterSimpleFilter is a filter type for logs.
type LogsFilterSimpleFilter = alerts.LogsFilter_SimpleFilter

// SimpleFilter is a filter type for logs.
type SimpleFilter = alerts.LogsSimpleFilter

// LogsThresholdConditionType is a type of condition for logs matching an alert.
type LogsThresholdConditionType = alerts.LogsThresholdConditionType

// LogsRatioTimeWindow is a time window setting for logs.
type LogsRatioTimeWindow = alerts.LogsRatioTimeWindow

// LogsRatioTimeWindowValue is a specific time window setting for logs.
type LogsRatioTimeWindowSpecificValue = alerts.LogsRatioTimeWindow_LogsRatioTimeWindowSpecificValue

// LogsThresholdRule is a rule for logs of the Threshold type.
type LogsThresholdRule = alerts.LogsThresholdRule

// LogsThresholdCondition is a condition for logs of the Threshold type.
type LogsThresholdCondition = alerts.LogsThresholdCondition

// LogsTimeWindow is a condition for logs of the Threshold type.
type LogsTimeWindow = alerts.LogsTimeWindow

// AlertDefPropertiesLogsImmediate is a property of an alert.
type AlertDefPropertiesLogsImmediate = alerts.AlertDefProperties_LogsImmediate

// AlertDefPropertiesTracingImmediate is a property of an alert.
type AlertDefPropertiesTracingImmediate = alerts.AlertDefProperties_TracingImmediate

// AlertDefPropertiesLogsThreshold is a property of an alert.
type AlertDefPropertiesLogsThreshold = alerts.AlertDefProperties_LogsThreshold

// AlertDefPropertiesLogsRatioThreshold is a property of an alert.
type AlertDefPropertiesLogsRatioThreshold = alerts.AlertDefProperties_LogsRatioThreshold

// AlertDefPropertiesLogsTimeRelativeThreshold is a property of an alert.
type AlertDefPropertiesLogsTimeRelativeThreshold = alerts.AlertDefProperties_LogsTimeRelativeThreshold

// AlertDefPropertiesMetricThreshold is a property of an alert.
type AlertDefPropertiesMetricThreshold = alerts.AlertDefProperties_MetricThreshold

// AlertDefPropertiesTracingThreshold is a property of an alert.
type AlertDefPropertiesTracingThreshold = alerts.AlertDefProperties_TracingThreshold

// AlertDefPropertiesFlow is a property of an alert.
type AlertDefPropertiesFlow = alerts.AlertDefProperties_Flow

// AlertDefPropertiesLogsUnusual is a property of an alert.
type AlertDefPropertiesLogsUnusual = alerts.AlertDefProperties_LogsUnusual

// AlertDefPropertiesMetricUnusual is a property of an alert.
type AlertDefPropertiesMetricUnusual = alerts.AlertDefProperties_MetricUnusual

// AlertDefPropertiesLogsNewValue is a property of an alert.
type AlertDefPropertiesLogsNewValue = alerts.AlertDefProperties_LogsNewValue

// AlertDefPropertiesLogsUniqueCount is a property of an alert.
type AlertDefPropertiesLogsUniqueCount = alerts.AlertDefProperties_LogsUniqueCount

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
type LogsFilter = alerts.LogsFilter

// LabelFilterType is a filter type
type LabelFilterType = alerts.LabelFilterType

// LabelFilters is a filter
type LabelFilters = alerts.LabelFilters

// LogSeverity is a filter
type LogSeverity = alerts.LogSeverity

// AlertDefIncidentSettings is the incident settings of an alert.
type AlertDefIncidentSettings = alerts.AlertDefIncidentSettings

// AlertDefIncidentSettingsMinutes is the incident settings of an alert.
type AlertDefIncidentSettingsMinutes = alerts.AlertDefIncidentSettings_Minutes

// LogsTimeWindowValue is a time window setting for logs.
type LogsTimeWindowValue = alerts.LogsTimeWindowValue

// AlertNotifyOn is a trigger type.
type AlertNotifyOn = alerts.NotifyOn

// LogsRatioCondition is a condition type for logs in ratio alerts.
type LogsRatioCondition = alerts.LogsRatioCondition

// LogsRatioRules is a rule type for logs in ratio alerts.
type LogsRatioRules = alerts.LogsRatioRules

// AlertEvaluationWindow is an evaluation window type for alerts.
// type AlertEvaluationWindow = alerts.EvaluationWindow

// LogFilterOperationType is a filter operation for logs.
type LogFilterOperationType = alerts.LogFilterOperationType

// LogsTimeRelativeComparedTo is a relative time setting for logs.
type LogsTimeRelativeComparedTo = alerts.LogsTimeRelativeComparedTo

// MetricTimeWindowValue is a time window setting for metrics.
type MetricTimeWindowValue = alerts.MetricTimeWindowValue

// TracingTimeWindowValue is a time window setting for tracing.
type TracingTimeWindowValue = alerts.TracingTimeWindowValue

// TracingFilterOperationType is an operation type setting for tracing.
type TracingFilterOperationType = alerts.TracingFilterOperationType

// AutoRetireTimeframe is a type of timeframe.
type AutoRetireTimeframe = alerts.AutoRetireTimeframe

// LogsRatioTimeWindowValue is a time window setting for logs.
type LogsRatioTimeWindowValue = alerts.LogsRatioTimeWindowValue

// LogsTimeWindowSpecificValue is a specific value for the time window setting for logs.
type LogsTimeWindowSpecificValue = alerts.LogsTimeWindow_LogsTimeWindowSpecificValue

// LogsThresholdType is a type of alert type
type LogsThresholdType = alerts.LogsThresholdType

// TracingImmediateType is a type of alert type
type TracingImmediateType = alerts.TracingImmediateType

// LogsImmediateType is a type of alert type
type LogsImmediateType = alerts.LogsImmediateType

// LogsRatioThresholdType is a type of alert type
type LogsRatioThresholdType = alerts.LogsRatioThresholdType

// MetricUnusualType is a type of alert type
type MetricUnusualType = alerts.MetricUnusualType

// MetricThresholdType is a type of alert type
type MetricThresholdType = alerts.MetricThresholdType

// TracingThresholdType is a type of alert type
type TracingThresholdType = alerts.TracingThresholdType

// LogsUnusualRule is a rule for ususual alert types.
type LogsUnusualRule = alerts.LogsUnusualRule

// LogsUnusualCondition is a condition type for ususual alert types.
type LogsUnusualCondition = alerts.LogsUnusualCondition

// UndetectedValuesManagement is a type for undetected values management.
type UndetectedValuesManagement = alerts.UndetectedValuesManagement

// LogsUnusualConditionType is the condition type for unusual alert types.
type LogsUnusualConditionType = alerts.LogsUnusualConditionType

// LogsUnusualConditionType is a value for unusual alert type conditions.
const LogsUnusualConditionTypeMoreThanOrUnspecified = alerts.LogsUnusualConditionType_LOGS_UNUSUAL_CONDITION_TYPE_MORE_THAN_USUAL_OR_UNSPECIFIED

// LogsUnusualType is a type of alert.
type LogsUnusualType = alerts.LogsUnusualType

// FlowType is a type of alert.
type FlowType = alerts.FlowType

// LogsNewValueType is a type of alert.
type LogsNewValueType = alerts.LogsNewValueType

// LogsNewValueCondition is a condition for the type of alert.
type LogsNewValueCondition = alerts.LogsNewValueCondition

// LogsNewValueRule is a rule for new log alerts.
type LogsNewValueRule = alerts.LogsNewValueRule

// LogsUniqueCountType is a type of alert type
type LogsUniqueCountType = alerts.LogsUniqueCountType

// LogsUniqueCountRule is a rule for a condition of the unique count alert type.
type LogsUniqueCountRule = alerts.LogsUniqueCountRule

// LogsUniqueCountCondition is a condition for the unique count alert type.
type LogsUniqueCountCondition = alerts.LogsUniqueCountCondition

// AlertDefType is a type of alert type
type AlertDefType = alerts.AlertDefType

// AlertDefType values.
const (
	AlertDefTypeLogsImmediateOrUnspecified = alerts.AlertDefType_ALERT_DEF_TYPE_LOGS_IMMEDIATE_OR_UNSPECIFIED
	AlertDefTypeLogsThreshold              = alerts.AlertDefType_ALERT_DEF_TYPE_LOGS_THRESHOLD
	AlertDefTypeLogsUnusual                = alerts.AlertDefType_ALERT_DEF_TYPE_LOGS_UNUSUAL
	AlertDefTypeLogsRatioThreshold         = alerts.AlertDefType_ALERT_DEF_TYPE_LOGS_RATIO_THRESHOLD
	AlertDefTypeLogsNewValue               = alerts.AlertDefType_ALERT_DEF_TYPE_LOGS_NEW_VALUE
	AlertDefTypeLogsUniqueCount            = alerts.AlertDefType_ALERT_DEF_TYPE_LOGS_UNIQUE_COUNT
	AlertDefTypeLogsTimeRelativeThreshold  = alerts.AlertDefType_ALERT_DEF_TYPE_LOGS_TIME_RELATIVE_THRESHOLD
	AlertDefTypeMetricThreshold            = alerts.AlertDefType_ALERT_DEF_TYPE_METRIC_THRESHOLD
	AlertDefTypeMetricUnusual              = alerts.AlertDefType_ALERT_DEF_TYPE_METRIC_UNUSUAL
	AlertDefTypeTracingImmediate           = alerts.AlertDefType_ALERT_DEF_TYPE_TRACING_IMMEDIATE
	AlertDefTypeTracingThreshold           = alerts.AlertDefType_ALERT_DEF_TYPE_TRACING_THRESHOLD
	AlertDefTypeFlow                       = alerts.AlertDefType_ALERT_DEF_TYPE_FLOW
)

// LogsTimeWindowValue are values for the time window setting for logs.
const (
	LogsTimeWindowValue5MinutesOrUnspecified = alerts.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_MINUTES_5_OR_UNSPECIFIED
	LogsTimeWindowValue10Minutes             = alerts.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_MINUTES_10
	LogsTimeWindowValue20Minutes             = alerts.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_MINUTES_20
	LogsTimeWindowValue15Minutes             = alerts.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_MINUTES_15
	LogsTimeWindowValue30Minutes             = alerts.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_MINUTES_30
	LogsTimeWindowValue1Hour                 = alerts.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_HOUR_1
	LogsTimeWindowValue2Hours                = alerts.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_HOURS_2
	LogsTimeWindowValue4Hours                = alerts.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_HOURS_4
	LogsTimeWindowValue6Hours                = alerts.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_HOURS_6
	LogsTimeWindowValue12Hours               = alerts.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_HOURS_12
	LogsTimeWindowValue24Hours               = alerts.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_HOURS_24
	LogsTimeWindowValue36Hours               = alerts.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_HOURS_36
)

// AlertDefPriority is the alert priority.
type AlertDefPriority = alerts.AlertDefPriority

// AlertDefPriority value.
const (
	AlertDefPriorityP1              = alerts.AlertDefPriority_ALERT_DEF_PRIORITY_P1
	AlertDefPriorityP2              = alerts.AlertDefPriority_ALERT_DEF_PRIORITY_P2
	AlertDefPriorityP3              = alerts.AlertDefPriority_ALERT_DEF_PRIORITY_P3
	AlertDefPriorityP4              = alerts.AlertDefPriority_ALERT_DEF_PRIORITY_P4
	AlertDefPriorityP5OrUnspecified = alerts.AlertDefPriority_ALERT_DEF_PRIORITY_P5_OR_UNSPECIFIED
)

// LogsRatioConditionType is a type of condition for logs in ratio alerts.
type LogsRatioConditionType = alerts.LogsRatioConditionType

// LogsRatioConditionType values.
const (
	LogsRatioConditionTypeMoreThanOrUnspecified = alerts.LogsRatioConditionType_LOGS_RATIO_CONDITION_TYPE_MORE_THAN_OR_UNSPECIFIED
	LogsRatioConditionTypeLessThan              = alerts.LogsRatioConditionType_LOGS_RATIO_CONDITION_TYPE_LESS_THAN
)

// AlertDayOfWeek values.
const (
	AlertDayOfWeekUnspecified = alerts.DayOfWeek_DAY_OF_WEEK_MONDAY_OR_UNSPECIFIED
	AlertDayOfWeekMonday      = alerts.DayOfWeek_DAY_OF_WEEK_MONDAY_OR_UNSPECIFIED
	AlertDayOfWeekTuesday     = alerts.DayOfWeek_DAY_OF_WEEK_TUESDAY
	AlertDayOfWeekWednesday   = alerts.DayOfWeek_DAY_OF_WEEK_WEDNESDAY
	AlertDayOfWeekThursday    = alerts.DayOfWeek_DAY_OF_WEEK_THURSDAY
	AlertDayOfWeekFriday      = alerts.DayOfWeek_DAY_OF_WEEK_FRIDAY
	AlertDayOfWeekSaturday    = alerts.DayOfWeek_DAY_OF_WEEK_SATURDAY
	AlertDayOfWeekSunday      = alerts.DayOfWeek_DAY_OF_WEEK_SUNDAY
)

// Alert API RPC names.
const (
	CreateAlertDefRpc    = alerts.AlertDefsService_CreateAlertDef_FullMethodName
	DeleteAlertDefRpc    = alerts.AlertDefsService_DeleteAlertDef_FullMethodName
	GetAlertDefRpc       = alerts.AlertDefsService_GetAlertDef_FullMethodName
	ListAlertDefsRpc     = alerts.AlertDefsService_ListAlertDefs_FullMethodName
	ReplaceAlertDefRpc   = alerts.AlertDefsService_ReplaceAlertDef_FullMethodName
	SetAlertDefActiveRpc = alerts.AlertDefsService_SetActive_FullMethodName
)

// AlertNotifyOnTriggeredOnlyUnspecified is a notification setting for an alert.
const AlertNotifyOnTriggeredOnlyUnspecified = alerts.NotifyOn_NOTIFY_ON_TRIGGERED_ONLY_UNSPECIFIED

// AlertNotifyOnTriggeredAndResolved is a notification setting for an alert.
const AlertNotifyOnTriggeredAndResolved = alerts.NotifyOn_NOTIFY_ON_TRIGGERED_AND_RESOLVED

// LogFilterOperation values.
const (
	LogFilterOperationIsOrUnspecified = alerts.LogFilterOperationType_LOG_FILTER_OPERATION_TYPE_IS_OR_UNSPECIFIED
	LogFilterOperationIncludes        = alerts.LogFilterOperationType_LOG_FILTER_OPERATION_TYPE_INCLUDES
	LogFilterOperationEndsWith        = alerts.LogFilterOperationType_LOG_FILTER_OPERATION_TYPE_ENDS_WITH
	LogFilterOperationStartsWith      = alerts.LogFilterOperationType_LOG_FILTER_OPERATION_TYPE_STARTS_WITH
)

// LogsSeverity values.
const (
	LogSeverityVerboseUnspecified = alerts.LogSeverity_LOG_SEVERITY_VERBOSE_UNSPECIFIED
	LogSeverityDebug              = alerts.LogSeverity_LOG_SEVERITY_DEBUG
	LogSeverityInfo               = alerts.LogSeverity_LOG_SEVERITY_INFO
	LogSeverityWarning            = alerts.LogSeverity_LOG_SEVERITY_WARNING
	LogSeverityError              = alerts.LogSeverity_LOG_SEVERITY_ERROR
	LogSeverityCritical           = alerts.LogSeverity_LOG_SEVERITY_CRITICAL
)

// MetricUnusualConditionType is a type of condition for unusual metrics.
type MetricUnusualConditionType = alerts.MetricUnusualConditionType

// MetricUnusualConditionType values.
const (
	MetricUnusualConditionTypeMoreThanOrUnspecified = alerts.MetricUnusualConditionType_METRIC_UNUSUAL_CONDITION_TYPE_MORE_THAN_USUAL_OR_UNSPECIFIED
	MetricUnusualConditionTypeLessThan              = alerts.MetricUnusualConditionType_METRIC_UNUSUAL_CONDITION_TYPE_LESS_THAN_USUAL
)

// LogsNewValueTimeWindow is a time window setting for new log alerts.
type LogsNewValueTimeWindow = alerts.LogsNewValueTimeWindow

// LogsNewValueTimeWindowValue is a time window setting for new logs.
type LogsNewValueTimeWindowValue = alerts.LogsNewValueTimeWindowValue

// LogsNewValueTimeWindowSpecificValue is a specific time window setting for new log alerts.
type LogsNewValueTimeWindowSpecificValue = alerts.LogsNewValueTimeWindow_LogsNewValueTimeWindowSpecificValue

// LogsNewValueTimeWindow values.
const (
	LogsNewValueTimeWindowValue12HoursOrUnspecified = alerts.LogsNewValueTimeWindowValue_LOGS_NEW_VALUE_TIME_WINDOW_VALUE_HOURS_12_OR_UNSPECIFIED
	LogsNewValueTimeWindowValue24Hours              = alerts.LogsNewValueTimeWindowValue_LOGS_NEW_VALUE_TIME_WINDOW_VALUE_HOURS_24
	LogsNewValueTimeWindowValue48Hours              = alerts.LogsNewValueTimeWindowValue_LOGS_NEW_VALUE_TIME_WINDOW_VALUE_HOURS_48
	LogsNewValueTimeWindowValue72Hours              = alerts.LogsNewValueTimeWindowValue_LOGS_NEW_VALUE_TIME_WINDOW_VALUE_HOURS_72
	LogsNewValueTimeWindowValue1Week                = alerts.LogsNewValueTimeWindowValue_LOGS_NEW_VALUE_TIME_WINDOW_VALUE_WEEK_1
	LogsNewValueTimeWindowValue1Month               = alerts.LogsNewValueTimeWindowValue_LOGS_NEW_VALUE_TIME_WINDOW_VALUE_MONTH_1
	LogsNewValueTimeWindowValue2Months              = alerts.LogsNewValueTimeWindowValue_LOGS_NEW_VALUE_TIME_WINDOW_VALUE_MONTHS_2
	LogsNewValueTimeWindowValue_3Months             = alerts.LogsNewValueTimeWindowValue_LOGS_NEW_VALUE_TIME_WINDOW_VALUE_MONTHS_3
)

// LogsTimeWindow values.
const (
	LogsTimeWindow5MinutesOrUnspecified = alerts.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_MINUTES_5_OR_UNSPECIFIED
	LogsTimeWindow10Minutes             = alerts.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_MINUTES_10
	LogsTimeWindow15Minutes             = alerts.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_MINUTES_15
	LogsTimeWindow30Minutes             = alerts.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_MINUTES_30
	LogsTimeWindow1Hour                 = alerts.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_HOUR_1
	LogsTimeWindow2Hours                = alerts.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_HOURS_2
	LogsTimeWindow4Hours                = alerts.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_HOURS_4
	LogsTimeWindow6Hours                = alerts.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_HOURS_6
	LogsTimeWindow12Hours               = alerts.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_HOURS_12
	LogsTimeWindow24Hours               = alerts.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_HOURS_24
	LogsTimeWindow36Hours               = alerts.LogsTimeWindowValue_LOGS_TIME_WINDOW_VALUE_HOURS_36
)

// AlertsOp is a type of operation for alerts.
type AlertsOp = alerts.AlertsOp

const (
	AlertsOpAndOrUnspecified = alerts.AlertsOp_ALERTS_OP_AND_OR_UNSPECIFIED
	AlertsOpOr               = alerts.AlertsOp_ALERTS_OP_OR
)

// NextOp is a next operation setting for alerts.
type NextOp = alerts.NextOp

const (
	NextOpAndOrUnspecified = alerts.NextOp_NEXT_OP_AND_OR_UNSPECIFIED
	NextOpOr               = alerts.NextOp_NEXT_OP_OR
)

// LogsUniqueValueTimeWindow is a time window for logs unique count alerts.
type LogsUniqueValueTimeWindow = alerts.LogsUniqueValueTimeWindow

// LogsUniqueValueTimeWindowSpecificValue is a specific time window setting for unique logs alerts.
type LogsUniqueValueTimeWindowSpecificValue = alerts.LogsUniqueValueTimeWindow_LogsUniqueValueTimeWindowSpecificValue

// LogsUniqueValueTimeWindowValue is a time window setting for unique logs.
type LogsUniqueValueTimeWindowValue = alerts.LogsUniqueValueTimeWindowValue

// LogsUniqueValueTimeWindow values.
const (
	LogsUniqueValueTimeWindowValue1MinuteOrUnspecified = alerts.LogsUniqueValueTimeWindowValue_LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_MINUTE_1_OR_UNSPECIFIED
	LogsUniqueValueTimeWindowValue15Minutes            = alerts.LogsUniqueValueTimeWindowValue_LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_MINUTES_15
	LogsUniqueValueTimeWindowValue20Minutes            = alerts.LogsUniqueValueTimeWindowValue_LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_MINUTES_20
	LogsUniqueValueTimeWindowValue30Minutes            = alerts.LogsUniqueValueTimeWindowValue_LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_MINUTES_30
	LogsUniqueValueTimeWindowValue1Hour                = alerts.LogsUniqueValueTimeWindowValue_LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_HOURS_1
	LogsUniqueValueTimeWindowValue2Hours               = alerts.LogsUniqueValueTimeWindowValue_LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_HOURS_2
	LogsUniqueValueTimeWindowValue4Hours               = alerts.LogsUniqueValueTimeWindowValue_LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_HOURS_4
	LogsUniqueValueTimeWindowValue6Hours               = alerts.LogsUniqueValueTimeWindowValue_LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_HOURS_6
	LogsUniqueValueTimeWindowValue12Hours              = alerts.LogsUniqueValueTimeWindowValue_LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_HOURS_12
	LogsUniqueValueTimeWindowValue24Hours              = alerts.LogsUniqueValueTimeWindowValue_LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_HOURS_24
)

// LogsTimeRelativeRule is a rule for the time relative alert type.
type LogsTimeRelativeRule = alerts.LogsTimeRelativeRule

// LogsTimeRelativeCondition is a condition for the time relative alert type.
type LogsTimeRelativeCondition = alerts.LogsTimeRelativeCondition

// LogsTimeRelativeConditionType is a condition for time relative alert type.
type LogsTimeRelativeConditionType = alerts.LogsTimeRelativeConditionType

// LogsTimeRelativeConditionType values.
const (
	LogsTimeRelativeConditionTypeMoreThanOrUnspecified = alerts.LogsTimeRelativeConditionType_LOGS_TIME_RELATIVE_CONDITION_TYPE_MORE_THAN_OR_UNSPECIFIED
	LogsTimeRelativeConditionTypeLessThan              = alerts.LogsTimeRelativeConditionType_LOGS_TIME_RELATIVE_CONDITION_TYPE_LESS_THAN
)

// LogsTimeRelativeThresholdType is a time relative alert type.
type LogsTimeRelativeThresholdType = alerts.LogsTimeRelativeThresholdType

// LogsTimeRelativeComparedTo is a time setting for time relative alerts.
const (
	LogsTimeRelativeComparedToPreviousHourOrUnspecified = alerts.LogsTimeRelativeComparedTo_LOGS_TIME_RELATIVE_COMPARED_TO_PREVIOUS_HOUR_OR_UNSPECIFIED
	LogsTimeRelativeComparedToSameHourYesterday         = alerts.LogsTimeRelativeComparedTo_LOGS_TIME_RELATIVE_COMPARED_TO_SAME_HOUR_YESTERDAY
	LogsTimeRelativeComparedToSameHourLastWeek          = alerts.LogsTimeRelativeComparedTo_LOGS_TIME_RELATIVE_COMPARED_TO_SAME_HOUR_LAST_WEEK
	LogsTimeRelativeComparedToYesterday                 = alerts.LogsTimeRelativeComparedTo_LOGS_TIME_RELATIVE_COMPARED_TO_YESTERDAY
	LogsTimeRelativeComparedToSameDayLastWeek           = alerts.LogsTimeRelativeComparedTo_LOGS_TIME_RELATIVE_COMPARED_TO_SAME_DAY_LAST_WEEK
	LogsTimeRelativeComparedToSameDayLastMonth          = alerts.LogsTimeRelativeComparedTo_LOGS_TIME_RELATIVE_COMPARED_TO_SAME_DAY_LAST_MONTH
)

// MetricTimeWindowValue values.
const (
	MetricTimeWindowValue1MinuteOrUnspecified = alerts.MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_MINUTES_1_OR_UNSPECIFIED
	MetricTimeWindowValue5Minutes             = alerts.MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_MINUTES_5
	MetricTimeWindowValue10Minutes            = alerts.MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_MINUTES_10
	MetricTimeWindowValue15Minutes            = alerts.MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_MINUTES_15
	MetricTimeWindowValue30Minutes            = alerts.MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_MINUTES_30
	MetricTimeWindowValue1Hour                = alerts.MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_HOUR_1
	MetricTimeWindowValue2Hours               = alerts.MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_HOURS_2
	MetricTimeWindowValue4Hours               = alerts.MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_HOURS_4
	MetricTimeWindowValue6Hours               = alerts.MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_HOURS_6
	MetricTimeWindowValue12Hours              = alerts.MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_HOURS_12
	MetricTimeWindowValue24Hours              = alerts.MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_HOURS_24
)

// TracingTimeWindowValue values.
const (
	TracingTimeWindowValue5MinutesOrUnspecified = alerts.TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_MINUTES_5_OR_UNSPECIFIED
	TracingTimeWindowValue10Minutes             = alerts.TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_MINUTES_10
	TracingTimeWindowValue15Minutes             = alerts.TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_MINUTES_15
	TracingTimeWindowValue30Minutes             = alerts.TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_MINUTES_30
	TracingTimeWindowValue1Hour                 = alerts.TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_HOUR_1
	TracingTimeWindowValue2Hours                = alerts.TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_HOURS_2
	TracingTimeWindowValue4Hours                = alerts.TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_HOURS_4
	TracingTimeWindowValue6Hours                = alerts.TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_HOURS_6
	TracingTimeWindowValue12Hours               = alerts.TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_HOURS_12
	TracingTimeWindowValue24Hours               = alerts.TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_HOURS_24
	TracingTimeWindowValue36Hours               = alerts.TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_HOURS_36
)

// TracingFilterOperationType values.
const (
	TracingFilterOperationTypeIsOrUnspecified = alerts.TracingFilterOperationType_TRACING_FILTER_OPERATION_TYPE_IS_OR_UNSPECIFIED
	TracingFilterOperationTypeIncludes        = alerts.TracingFilterOperationType_TRACING_FILTER_OPERATION_TYPE_INCLUDES
	TracingFilterOperationTypeEndsWith        = alerts.TracingFilterOperationType_TRACING_FILTER_OPERATION_TYPE_ENDS_WITH
	TracingFilterOperationTypeStartsWith      = alerts.TracingFilterOperationType_TRACING_FILTER_OPERATION_TYPE_STARTS_WITH
)

// TimeframeType is a type for timeframes.
type TimeframeType = alerts.TimeframeType

// TimeframeType values.
const (
	TimeframeTypeUnspecified = alerts.TimeframeType_TIMEFRAME_TYPE_UNSPECIFIED
	TimeframeTypeUpTo        = alerts.TimeframeType_TIMEFRAME_TYPE_UP_TO
)

// AutoRetireTimeframe values.
const (
	AutoRetireTimeframeNeverOrUnspecified = alerts.AutoRetireTimeframe_AUTO_RETIRE_TIMEFRAME_NEVER_OR_UNSPECIFIED
	AutoRetireTimeframe5Minutes           = alerts.AutoRetireTimeframe_AUTO_RETIRE_TIMEFRAME_MINUTES_5
	AutoRetireTimeframe10Minutes          = alerts.AutoRetireTimeframe_AUTO_RETIRE_TIMEFRAME_MINUTES_10
	AutoRetireTimeframe1Hour              = alerts.AutoRetireTimeframe_AUTO_RETIRE_TIMEFRAME_HOUR_1
	AutoRetireTimeframe2Hours             = alerts.AutoRetireTimeframe_AUTO_RETIRE_TIMEFRAME_HOURS_2
	AutoRetireTimeframe6Hours             = alerts.AutoRetireTimeframe_AUTO_RETIRE_TIMEFRAME_HOURS_6
	AutoRetireTimeframe12Hours            = alerts.AutoRetireTimeframe_AUTO_RETIRE_TIMEFRAME_HOURS_12
	AutoRetireTimeframe24Hours            = alerts.AutoRetireTimeframe_AUTO_RETIRE_TIMEFRAME_HOURS_24
)

// LogsRatioTimeWindow values.
const (
	LogsRatioTimeWindowValue5MinutesOrUnspecified = alerts.LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_MINUTES_5_OR_UNSPECIFIED
	LogsRatioTimeWindowValue10Minutes             = alerts.LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_MINUTES_10
	LogsRatioTimeWindowValue15Minutes             = alerts.LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_MINUTES_15
	LogsRatioTimeWindowValue30Minutes             = alerts.LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_MINUTES_30
	LogsRatioTimeWindowValue1Hour                 = alerts.LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_HOUR_1
	LogsRatioTimeWindowValue2Hours                = alerts.LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_HOURS_2
	LogsRatioTimeWindowValue4Hours                = alerts.LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_HOURS_4
	LogsRatioTimeWindowValue6Hours                = alerts.LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_HOURS_6
	LogsRatioTimeWindowValue12Hours               = alerts.LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_HOURS_12
	LogsRatioTimeWindowValue24Hours               = alerts.LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_HOURS_24
	LogsRatioTimeWindowValue36Hours               = alerts.LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_HOURS_36
)

// MetricThresholdConditionType is a type of condition for metrics.
type MetricThresholdConditionType = alerts.MetricThresholdConditionType

// MetricThresholdConditionType values.
const (
	MetricThresholdConditionTypeMoreThanOrUnspecified = alerts.MetricThresholdConditionType_METRIC_THRESHOLD_CONDITION_TYPE_MORE_THAN_OR_UNSPECIFIED
	MetricThresholdConditionTypeLessThan              = alerts.MetricThresholdConditionType_METRIC_THRESHOLD_CONDITION_TYPE_LESS_THAN
	MetricThresholdConditionTypeMoreThanOrEquals      = alerts.MetricThresholdConditionType_METRIC_THRESHOLD_CONDITION_TYPE_MORE_THAN_OR_EQUALS
	MetricThresholdConditionTypeLessThanOrEquals      = alerts.MetricThresholdConditionType_METRIC_THRESHOLD_CONDITION_TYPE_LESS_THAN_OR_EQUALS
)

// LogsRatioGroupByFor is a group by setting for logs.
type LogsRatioGroupByFor = alerts.LogsRatioGroupByFor

// LogsRatioGroupByFor values.
const (
	LogsRatioGroupByForBothOrUnspecified = alerts.LogsRatioGroupByFor_LOGS_RATIO_GROUP_BY_FOR_BOTH_OR_UNSPECIFIED
	LogsRatioGroupByForNumeratorOnly     = alerts.LogsRatioGroupByFor_LOGS_RATIO_GROUP_BY_FOR_NUMERATOR_ONLY
	LogsRatioGroupByForDenumeratorOnly   = alerts.LogsRatioGroupByFor_LOGS_RATIO_GROUP_BY_FOR_DENUMERATOR_ONLY
)

const (
	LogsThresholdConditionTypeMoreThanOrUnspecified = alerts.LogsThresholdConditionType_LOGS_THRESHOLD_CONDITION_TYPE_MORE_THAN_OR_UNSPECIFIED
	LogsThresholdConditionTypeLessThan              = alerts.LogsThresholdConditionType_LOGS_THRESHOLD_CONDITION_TYPE_LESS_THAN
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
