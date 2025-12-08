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

	dashboards "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/dashboards/v1"
	ast "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/dashboards/v1/ast"
	annotations "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/dashboards/v1/ast/annotations"
	filters "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/dashboards/v1/ast/filters"
	variables "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/dashboards/v1/ast/variables"
	widgets "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/dashboards/v1/ast/widgets"
	widgetsCommon "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/dashboards/v1/ast/widgets/common"
	common "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/dashboards/v1/common"
	services "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/dashboards/v1/services"
)

// DashboardsClient is a client for the Coralogix Dashboards API.
type DashboardsClient struct {
	callPropertiesCreator CallPropertiesCreator
}

// CreateDashboardRequest is a request to create a dashboard.
type CreateDashboardRequest = services.CreateDashboardRequest

// ReplaceDashboardRequest is a request to replace a dashboard.
type ReplaceDashboardRequest = services.ReplaceDashboardRequest

// DeleteDashboardRequest is a request to delete a dashboard.
type DeleteDashboardRequest = services.DeleteDashboardRequest

// GetDashboardRequest is a request to get a dashboard.
type GetDashboardRequest = services.GetDashboardRequest

// PinDashboardRequest is a request to pin a dashboard.
type PinDashboardRequest = services.PinDashboardRequest

// UnpinDashboardRequest is a request to unpin a dashboard.
type UnpinDashboardRequest = services.UnpinDashboardRequest

// FolderPath is a dashboard folder path.
type FolderPath = ast.FolderPath

// Dashboard is a dashboard.
type Dashboard = ast.Dashboard

// RowStyle is a style for a row.
type RowStyle = widgets.RowStyle

// PromQLQueryType is a type for promql query.
type PromQLQueryType = common.PromQLQueryType

// PromQLQueryType values.
const (
	PromQLQueryTypeInstant     = common.PromQLQueryType_PROM_QL_QUERY_TYPE_INSTANT
	PromQLQueryTypeRange       = common.PromQLQueryType_PROM_QL_QUERY_TYPE_RANGE
	PromQLQueryTypeUnspecified = common.PromQLQueryType_PROM_QL_QUERY_TYPE_UNSPECIFIED
)

// RowStyle values.
const (
	RowStyleOneLine     = widgets.RowStyle_ROW_STYLE_ONE_LINE
	RowStyleTwoLine     = widgets.RowStyle_ROW_STYLE_TWO_LINE
	RowStyleCondensed   = widgets.RowStyle_ROW_STYLE_CONDENSED
	RowStyleJSON        = widgets.RowStyle_ROW_STYLE_JSON
	RowStyleList        = widgets.RowStyle_ROW_STYLE_LIST
	RowStyleUnspecified = widgets.RowStyle_ROW_STYLE_UNSPECIFIED
)

// DashboardLegend is the type for the legends.
type DashboardLegend = widgetsCommon.Legend

// DashboardLegendColumn is the column type for the legends.
type DashboardLegendColumn = widgetsCommon.Legend_LegendColumn

// LegendColumn values.
const (
	LegendColumnUnspecified = widgetsCommon.Legend_LEGEND_COLUMN_UNSPECIFIED
	LegendColumnMin         = widgetsCommon.Legend_LEGEND_COLUMN_MIN
	LegendColumnMax         = widgetsCommon.Legend_LEGEND_COLUMN_MAX
	LegendColumnSum         = widgetsCommon.Legend_LEGEND_COLUMN_SUM
	LegendColumnAvg         = widgetsCommon.Legend_LEGEND_COLUMN_AVG
	LegendColumnLast        = widgetsCommon.Legend_LEGEND_COLUMN_LAST
	LegendColumnName        = widgetsCommon.Legend_LEGEND_COLUMN_NAME
)

// LegendPlacement is a type for the legend placement.
type LegendPlacement = widgetsCommon.Legend_LegendPlacement

// LegendPlacement values.
const (
	LegendPlacementUnspecified = widgetsCommon.Legend_LEGEND_PLACEMENT_UNSPECIFIED
	LegendPlacementAuto        = widgetsCommon.Legend_LEGEND_PLACEMENT_AUTO
	LegendPlacementBottom      = widgetsCommon.Legend_LEGEND_PLACEMENT_BOTTOM
	LegendPlacementSide        = widgetsCommon.Legend_LEGEND_PLACEMENT_SIDE
	LegendPlacementHidden      = widgetsCommon.Legend_LEGEND_PLACEMENT_HIDDEN
)

// OrderDirection is the sorting direction
type OrderDirection = common.OrderDirection

// OrderDirection values.
const (
	OrderDirectionAsc         = common.OrderDirection_ORDER_DIRECTION_ASC
	OrderDirectionDesc        = common.OrderDirection_ORDER_DIRECTION_DESC
	OrderDirectionUnspecified = common.OrderDirection_ORDER_DIRECTION_UNSPECIFIED
)

// LineChartStackedLine is a visual configuration for the line chart.
type LineChartStackedLine = widgets.LineChart_StackedLine

// LineChartStackedLineType values.
const (
	LineChartStackedLineUnspecified = widgets.LineChart_STACKED_LINE_UNSPECIFIED
	LineChartStackedLineAbsolute    = widgets.LineChart_STACKED_LINE_ABSOLUTE
	LineChartStackedLineRelative    = widgets.LineChart_STACKED_LINE_RELATIVE
)

// LineChartTooltip is a tooltip for the line chart.
type LineChartTooltip = widgets.LineChart_Tooltip

// LineChartTooltipType is the type of the line chart tooltip.
type LineChartTooltipType = widgets.LineChart_TooltipType

// LineChartTooltipType values.
const (
	LineChartToolTipTypeUnspecified = widgets.LineChart_TOOLTIP_TYPE_UNSPECIFIED
	LineChartToolTipTypeAll         = widgets.LineChart_TOOLTIP_TYPE_ALL
	LineChartToolTipTypeSingle      = widgets.LineChart_TOOLTIP_TYPE_SINGLE
)

// ScaleType is the type of the scale.
type ScaleType = widgetsCommon.ScaleType

// ScaleType values.
const (
	ScaleTypeUnspecified = widgetsCommon.ScaleType_SCALE_TYPE_UNSPECIFIED
	ScaleTypeLinear      = widgetsCommon.ScaleType_SCALE_TYPE_LINEAR
	ScaleTypeLogarithmic = widgetsCommon.ScaleType_SCALE_TYPE_LOGARITHMIC
)

// Unit is the unit for a widget.
type Unit = widgetsCommon.Unit

// Unit values.
const (
	UnitUnspecified  = widgetsCommon.Unit_UNIT_UNSPECIFIED
	UnitMicroseconds = widgetsCommon.Unit_UNIT_MICROSECONDS
	UnitMilliseconds = widgetsCommon.Unit_UNIT_MILLISECONDS
	UnitNanoseconds  = widgetsCommon.Unit_UNIT_NANOSECONDS
	UnitSeconds      = widgetsCommon.Unit_UNIT_SECONDS
	UnitBytes        = widgetsCommon.Unit_UNIT_BYTES
	UnitKbytes       = widgetsCommon.Unit_UNIT_KBYTES
	UnitMbytes       = widgetsCommon.Unit_UNIT_MBYTES
	UnitGbytes       = widgetsCommon.Unit_UNIT_GBYTES
	UnitBytesIec     = widgetsCommon.Unit_UNIT_BYTES_IEC
	UnitKibytes      = widgetsCommon.Unit_UNIT_KIBYTES
	UnitMibytes      = widgetsCommon.Unit_UNIT_MIBYTES
	UnitGibytes      = widgetsCommon.Unit_UNIT_GIBYTES
	UnitEurCents     = widgetsCommon.Unit_UNIT_EUR_CENTS
	UnitEur          = widgetsCommon.Unit_UNIT_EUR
	UnitUsdCents     = widgetsCommon.Unit_UNIT_USD_CENTS
	UnitUsd          = widgetsCommon.Unit_UNIT_USD
	UnitCustom       = widgetsCommon.Unit_UNIT_CUSTOM
	UnitPercent01    = widgetsCommon.Unit_UNIT_PERCENT_ZERO_ONE
	UnitPercent100   = widgetsCommon.Unit_UNIT_PERCENT_ZERO_HUNDRED
)

// GaugeUnit is the unit for a widget.
type GaugeUnit = widgets.Gauge_Unit

// GaugeUnit values.
const (
	GaugeUnitUnspecified  = widgets.Gauge_UNIT_UNSPECIFIED
	GaugeUnitMicroseconds = widgets.Gauge_UNIT_MICROSECONDS
	GaugeUnitMilliseconds = widgets.Gauge_UNIT_MILLISECONDS
	GaugeUnitNanoseconds  = widgets.Gauge_UNIT_NANOSECONDS
	GaugeUnitNumber       = widgets.Gauge_UNIT_NUMBER
	GaugeUnitPercent      = widgets.Gauge_UNIT_PERCENT
	GaugeUnitSeconds      = widgets.Gauge_UNIT_SECONDS
	GaugeUnitBytes        = widgets.Gauge_UNIT_BYTES
	GaugeUnitKbytes       = widgets.Gauge_UNIT_KBYTES
	GaugeUnitMbytes       = widgets.Gauge_UNIT_MBYTES
	GaugeUnitGbytes       = widgets.Gauge_UNIT_GBYTES
	GaugeUnitBytesIec     = widgets.Gauge_UNIT_BYTES_IEC
	GaugeUnitKibytes      = widgets.Gauge_UNIT_KIBYTES
	GaugeUnitMibytes      = widgets.Gauge_UNIT_MIBYTES
	GaugeUnitGibytes      = widgets.Gauge_UNIT_GIBYTES
	GaugeUnitEurCents     = widgets.Gauge_UNIT_EUR_CENTS
	GaugeUnitEur          = widgets.Gauge_UNIT_EUR
	GaugeUnitUsdCents     = widgets.Gauge_UNIT_USD_CENTS
	GaugeUnitUsd          = widgets.Gauge_UNIT_USD
	GaugeUnitCustom       = widgets.Gauge_UNIT_CUSTOM
	GaugeUnitPercent01    = widgets.Gauge_UNIT_PERCENT_ZERO_ONE
	GaugeUnitPercent100   = widgets.Gauge_UNIT_PERCENT_ZERO_HUNDRED
)

// PieChartLabelSource is the source of the labels for the pie chart.
type PieChartLabelSource = widgets.PieChart_LabelSource

// PieChartLabelSource values.
const (
	PieChartLabelSourceUnspecified = widgets.PieChart_LABEL_SOURCE_UNSPECIFIED
	PieChartLabelSourceInner       = widgets.PieChart_LABEL_SOURCE_INNER
	PieChartLabelSourceStack       = widgets.PieChart_LABEL_SOURCE_STACK
)

// PieChartLabelDefinition is a label definition type.
type PieChartLabelDefinition = widgets.PieChart_LabelDefinition

// PieChartStackDefinition is a stack definition type.
type PieChartStackDefinition = widgets.PieChart_StackDefinition

// BarChartStackDefinition is a stack definition type.
type BarChartStackDefinition = widgets.BarChart_StackDefinition

// BarChartDataprimeQuery is a type for bar chart data prime query.
type BarChartDataprimeQuery = widgets.BarChart_DataprimeQuery

// HorizontalBarChartStackDefinition is the source of the labels for the horizontal bar chart.
type HorizontalBarChartStackDefinition = widgets.HorizontalBarChart_StackDefinition

// SpansAggregation is a type for spans aggregation.
type SpansAggregation = common.SpansAggregation

// SpansAggregationMetricAggregation is a type for metric aggregation.
type SpansAggregationMetricAggregation = common.SpansAggregation_MetricAggregation_

// SpansAggregationMetricAggregationInner is a type for spans aggregation metric aggregation.
type SpansAggregationMetricAggregationInner = common.SpansAggregation_MetricAggregation

// SpansAggregationDimensionAggregation is a type for aggregation.
type SpansAggregationDimensionAggregation = common.SpansAggregation_DimensionAggregation_

// SpansAggregationDimensionAggregationInner is a type for spans aggregation metric aggregation.
type SpansAggregationDimensionAggregationInner = common.SpansAggregation_DimensionAggregation

// SpansAggregationDimensionAggregationDimensionField is a type for spans aggregation dimension aggregation dimension field.
type SpansAggregationDimensionAggregationDimensionField = common.SpansAggregation_DimensionAggregation_DimensionField

// SpansAggregationDimensionAggregationDimensionField values.
const (
	SpansAggregationDimensionAggregationDimensionFieldUnspecified = common.SpansAggregation_DimensionAggregation_DIMENSION_FIELD_UNSPECIFIED
	SpansAggregationDimensionAggregationDimensionFieldTraceID     = common.SpansAggregation_DimensionAggregation_DIMENSION_FIELD_TRACE_ID
)

// SpansAggregationMetricAggregationMetricAggregationType is a type for spans aggregation metric aggregation metric aggregation type.
type SpansAggregationMetricAggregationMetricAggregationType = common.SpansAggregation_MetricAggregation_MetricAggregationType

// SpansAggregationMetricAggregationMetricTypeUnspecified values.
const (
	SpansAggregationMetricAggregationMetricTypeUnspecified  = common.SpansAggregation_MetricAggregation_METRIC_AGGREGATION_TYPE_UNSPECIFIED
	SpansAggregationMetricAggregationMetricTypeMin          = common.SpansAggregation_MetricAggregation_METRIC_AGGREGATION_TYPE_MIN
	SpansAggregationMetricAggregationMetricTypeMax          = common.SpansAggregation_MetricAggregation_METRIC_AGGREGATION_TYPE_MAX
	SpansAggregationMetricAggregationMetricTypeAverage      = common.SpansAggregation_MetricAggregation_METRIC_AGGREGATION_TYPE_AVERAGE
	SpansAggregationMetricAggregationMetricTypeSum          = common.SpansAggregation_MetricAggregation_METRIC_AGGREGATION_TYPE_SUM
	SpansAggregationMetricAggregationMetricTypePercentile99 = common.SpansAggregation_MetricAggregation_METRIC_AGGREGATION_TYPE_PERCENTILE_99
	SpansAggregationMetricAggregationMetricTypePercentile95 = common.SpansAggregation_MetricAggregation_METRIC_AGGREGATION_TYPE_PERCENTILE_95
	SpansAggregationMetricAggregationMetricTypePercentile50 = common.SpansAggregation_MetricAggregation_METRIC_AGGREGATION_TYPE_PERCENTILE_50
)

// SpansAggregationMetricAggregationMetricField is a type for spans aggregation metric aggregation metric field.
type SpansAggregationMetricAggregationMetricField = common.SpansAggregation_MetricAggregation_MetricField

// SpansAggregationMetricAggregationMetricField values.
const (
	SpansAggregationMetricAggregationMetricFieldUnspecified = common.SpansAggregation_MetricAggregation_METRIC_FIELD_UNSPECIFIED
	SpansAggregationMetricAggregationMetricFieldDuration    = common.SpansAggregation_MetricAggregation_METRIC_FIELD_DURATION
)

// UUID is a type wrapper for UUID.
type UUID = dashboards.UUID

// GaugeQueryMetrics is a type for gauge query for metrics.
type GaugeQueryMetrics = widgets.Gauge_Query_Metrics

// GaugeQueryLogs is a type for gauge query for logs.
type GaugeQueryLogs = widgets.Gauge_Query_Logs

// GaugeQuerySpans is a type for gauge query for spans.
type GaugeQuerySpans = widgets.Gauge_Query_Spans

// GaugeSpansQuery is a type of gauge query for spans.
type GaugeSpansQuery = widgets.Gauge_SpansQuery

// GaugeQuery is a type for gauge query.
type GaugeQuery = widgets.Gauge_Query

// GaugeAggregation is a type for gauge aggregation.
type GaugeAggregation = widgets.Gauge_Aggregation

// GaugeAggregation values.
const (
	GaugeAggregationUnspecified = widgets.Gauge_AGGREGATION_UNSPECIFIED
	GaugeAggregationLast        = widgets.Gauge_AGGREGATION_LAST
	GaugeAggregationMin         = widgets.Gauge_AGGREGATION_MIN
	GaugeAggregationMax         = widgets.Gauge_AGGREGATION_MAX
	GaugeAggregationAvg         = widgets.Gauge_AGGREGATION_AVG
	GaugeAggregationSum         = widgets.Gauge_AGGREGATION_SUM
)

// RPC names.
const (
	CreateDashboardRPC         = services.DashboardsService_CreateDashboard_FullMethodName
	ReplaceDashboardRPC        = services.DashboardsService_ReplaceDashboard_FullMethodName
	DeleteDashboardRPC         = services.DashboardsService_DeleteDashboard_FullMethodName
	GetDashboardRPC            = services.DashboardsService_GetDashboard_FullMethodName
	PinDashboardRPC            = services.DashboardsService_PinDashboard_FullMethodName
	UnpinDashboardRPC          = services.DashboardsService_UnpinDashboard_FullMethodName
	ReplaceDefaultDashboardRPC = services.DashboardsService_ReplaceDefaultDashboard_FullMethodName
	AssignDashboardFolderRPC   = services.DashboardsService_AssignDashboardFolder_FullMethodName
)

// SortByType is the type for sorting by.
type SortByType = widgetsCommon.SortByType

// SortByType values.
const (
	SortByTypeUnspecified = widgetsCommon.SortByType_SORT_BY_TYPE_UNSPECIFIED
	SortByTypeValue       = widgetsCommon.SortByType_SORT_BY_TYPE_VALUE
	SortByTypeName        = widgetsCommon.SortByType_SORT_BY_TYPE_NAME
)

// SpanFieldMetadataField values.
const (
	SpanFieldMetadataFieldUnspecified     = common.SpanField_METADATA_FIELD_UNSPECIFIED
	SpanFieldMetadataFieldApplicationName = common.SpanField_METADATA_FIELD_APPLICATION_NAME
	SpanFieldMetadataFieldSubsystemName   = common.SpanField_METADATA_FIELD_SUBSYSTEM_NAME
	SpanFieldMetadataFieldServiceName     = common.SpanField_METADATA_FIELD_SERVICE_NAME
	SpanFieldMetadataFieldOperationName   = common.SpanField_METADATA_FIELD_OPERATION_NAME
)

// SpansAggregationDimensionAggregationType is the type for spans aggregation dimension aggregation type.
type SpansAggregationDimensionAggregationType = common.SpansAggregation_DimensionAggregation_DimensionAggregationType

// SpansAggregationDimensionAggregationType values.
const (
	SpansAggregationDimensionAggregationTypeUnspecified = common.SpansAggregation_DimensionAggregation_DIMENSION_AGGREGATION_TYPE_UNSPECIFIED
	SpansAggregationDimensionAggregationTypeUniqueCount = common.SpansAggregation_DimensionAggregation_DIMENSION_AGGREGATION_TYPE_UNIQUE_COUNT
	SpansAggregationDimensionAggregationTypeErrorCount  = common.SpansAggregation_DimensionAggregation_DIMENSION_AGGREGATION_TYPE_ERROR_COUNT
)

// DatasetScope is the scope of the dataset.
type DatasetScope = common.DatasetScope

// DatasetScope values.
const (
	DatasetScopeUnspecified = common.DatasetScope_DATASET_SCOPE_UNSPECIFIED
	DatasetScopeUserData    = common.DatasetScope_DATASET_SCOPE_USER_DATA
	DatasetScopeLabel       = common.DatasetScope_DATASET_SCOPE_LABEL
	DatasetScopeMetadata    = common.DatasetScope_DATASET_SCOPE_METADATA
)

// DataModeType is the type for data mode.
type DataModeType = widgetsCommon.DataModeType

// DataModeType values.
const (
	DataModeTypeHighUnspecified = widgetsCommon.DataModeType_DATA_MODE_TYPE_HIGH_UNSPECIFIED
	DataModeTypeArchive         = widgetsCommon.DataModeType_DATA_MODE_TYPE_ARCHIVE
)

// GaugeThreshold is the type for gauge threshold.
type GaugeThreshold = widgets.Gauge_Threshold

// GaugeThresholdBy is the type for gauge threshold by.
type GaugeThresholdBy = widgets.Gauge_ThresholdBy

// GaugeThresholdBy values.
const (
	GaugeThresholdByUnspecified = widgets.Gauge_THRESHOLD_BY_UNSPECIFIED
	GaugeThresholdByValue       = widgets.Gauge_THRESHOLD_BY_VALUE
	GaugeThresholdByBackground  = widgets.Gauge_THRESHOLD_BY_BACKGROUND
)

// MultiSelectRefreshStrategy is the type for multi select refresh strategy.
type MultiSelectRefreshStrategy = variables.MultiSelect_RefreshStrategy

// MultiSelectRefreshStrategy values.
const (
	MultiSelectRefreshStrategyUnspecified       = variables.MultiSelect_REFRESH_STRATEGY_UNSPECIFIED
	MultiSelectRefreshStrategyOnDashboardLoad   = variables.MultiSelect_REFRESH_STRATEGY_ON_DASHBOARD_LOAD
	MultiSelectRefreshStrategyOnTimeFrameChange = variables.MultiSelect_REFRESH_STRATEGY_ON_TIME_FRAME_CHANGE
)

// Annotation is an annotation for dashboards.
type Annotation = annotations.Annotation

// DashboardFilter is a filter for dashboards.
type DashboardFilter = filters.Filter

// DashboardFilterSource is a filter source for dashboards.
type DashboardFilterSource = filters.Filter_Source

// DashboardFilterSourceSpans is a filter source for dashboards.
type DashboardFilterSourceSpans = filters.Filter_Source_Spans

// DashboardFilterSourceLogs is a filter source for dashboards.
type DashboardFilterSourceLogs = filters.Filter_Source_Logs

// DashboardFilterSourceMetrics is a filter source for dashboards.
type DashboardFilterSourceMetrics = filters.Filter_Source_Metrics

// DashboardFilterSpansFilter is a filter type.
type DashboardFilterSpansFilter = filters.Filter_SpansFilter

// DashboardFilterLogsFilter is a filter type.
type DashboardFilterLogsFilter = filters.Filter_LogsFilter

// DashboardFilterMetricsFilter is a filter type.
type DashboardFilterMetricsFilter = filters.Filter_MetricsFilter

// SpanField is a filter type.
type SpanField = common.SpanField

// SpanFieldMetadataField is the type for span field metadata field.
type SpanFieldMetadataField = common.SpanField_MetadataField_

// SpanFieldMetadataFieldInner is the type for span field metadata field.
type SpanFieldMetadataFieldInner = common.SpanField_MetadataField

// SpanFieldTagField is a tag field type for spans.
type SpanFieldTagField = common.SpanField_TagField

// SpanFieldProcessTagField is a tag field type for spans.
type SpanFieldProcessTagField = common.SpanField_ProcessTagField

// MultiSelectValueDisplayOptions is a type for display options.
type MultiSelectValueDisplayOptions = variables.MultiSelect_ValueDisplayOptions

// MultiSelectSource is a source for multi select.
type MultiSelectSource = variables.MultiSelect_Source

// MultiSelectSourceQuery is a source for multi select.
type MultiSelectSourceQuery = variables.MultiSelect_Source_Query

// MultiSelectQuery is a query type for multi select.
type MultiSelectQuery = variables.MultiSelect_Query

// MultiSelectQuerySource is a source for multi select.
type MultiSelectQuerySource = variables.MultiSelect_QuerySource

// MultiSelectQueryLogsQuery is a query type for multi select.
type MultiSelectQueryLogsQuery = variables.MultiSelect_Query_LogsQuery_

// MultiSelectQueryLogsQueryInner is a query type for multi select.
type MultiSelectQueryLogsQueryInner = variables.MultiSelect_Query_LogsQuery

// MultiSelectQueryLogsQueryType is a query type for multi select.
type MultiSelectQueryLogsQueryType = variables.MultiSelect_Query_LogsQuery_Type

// MultiSelectSourceSpanField is a source for multi select.
type MultiSelectSourceSpanField = variables.MultiSelect_Source_SpanField

// MultiSelectSpanFieldSource is a source for multi select.
type MultiSelectSpanFieldSource = variables.MultiSelect_SpanFieldSource

// MultiSelectSourceConstantList is a source for multi select.
type MultiSelectSourceConstantList = variables.MultiSelect_Source_ConstantList

// MultiSelectConstantListSource is a source for multi select.
type MultiSelectConstantListSource = variables.MultiSelect_ConstantListSource

// MultiSelectSourceMetricLabel is a source for multi select.
type MultiSelectSourceMetricLabel = variables.MultiSelect_Source_MetricLabel

// MultiSelectMetricLabelSource is a source for multi select.
type MultiSelectMetricLabelSource = variables.MultiSelect_MetricLabelSource

// MultiSelectLogsPathSource is a source for multi select.
type MultiSelectLogsPathSource = variables.MultiSelect_LogsPathSource

// MultiSelectSourceLogsPath is a source for multi select.
type MultiSelectSourceLogsPath = variables.MultiSelect_Source_LogsPath

// WidgetDefinitionMarkdown is a widget definition.
type WidgetDefinitionMarkdown = ast.Widget_Definition_Markdown

// WidgetDefinitionHorizontalBarChart is a widget definition.
type WidgetDefinitionHorizontalBarChart = ast.Widget_Definition_HorizontalBarChart

// WidgetDefinitionLineChart is a widget definition.
type WidgetDefinitionLineChart = ast.Widget_Definition_LineChart

// WidgetDefinitionDataTable is a widget definition.
type WidgetDefinitionDataTable = ast.Widget_Definition_DataTable

// WidgetDefinitionGauge is a widget definition.
type WidgetDefinitionGauge = ast.Widget_Definition_Gauge

// WidgetDefinitionPieChart is a widget definition.
type WidgetDefinitionPieChart = ast.Widget_Definition_PieChart

// WidgetDefinitionBarChart is a widget definition.
type WidgetDefinitionBarChart = ast.Widget_Definition_BarChart

// WidgetDefinitionHexagon is a widget definition for hexagons.
type WidgetDefinitionHexagon = ast.Widget_Definition_Hexagon

// Markdown is a widget definition inner type.
type Markdown = widgets.Markdown

// HorizontalBarChart is a widget definition inner type.
type HorizontalBarChart = widgets.HorizontalBarChart

// LineChart is a widget definition inner type.
type LineChart = widgets.LineChart

// DataTable is a widget definition inner type.
type DataTable = widgets.DataTable

// Gauge is a widget definition inner type.
type Gauge = widgets.Gauge

// PieChart is a widget definition inner type.
type PieChart = widgets.PieChart

// BarChart is a widget definition inner type.
type BarChart = widgets.BarChart

// WidgetDefinition is a widget definition.
type WidgetDefinition = ast.Widget_Definition

// Hexagon is a widget definition inner type.
type Hexagon = widgets.Hexagon

// DashboardRow is a row in a dashboard.
type DashboardRow = ast.Row

// DashboardRowAppearance is the appearance of a row in a dashboard.
type DashboardRowAppearance = ast.Row_Appearance

// DashboardWidget is a widget in a dashboard.
type DashboardWidget = ast.Widget

// DashboardWidgetAppearance is the appearance of a widget in a dashboard.
type DashboardWidgetAppearance = ast.Widget_Appearance

// MultiSelectQueryMetricsQueryStringOrVariable is a query string type.
type MultiSelectQueryMetricsQueryStringOrVariable = variables.MultiSelect_Query_MetricsQuery_StringOrVariable

// MultiSelectQueryMetricsQueryStringOrVariableString is a string query type.
type MultiSelectQueryMetricsQueryStringOrVariableString = variables.MultiSelect_Query_MetricsQuery_StringOrVariable_StringValue

// MultiSelectQueryMetricsQueryStringOrVariableVariable is a variable query type.
type MultiSelectQueryMetricsQueryStringOrVariableVariable = variables.MultiSelect_Query_MetricsQuery_StringOrVariable_VariableName

// MultiSelectQueryMetricsQuerySelection is a selection for a multi select query.
type MultiSelectQueryMetricsQuerySelection = variables.MultiSelect_Query_MetricsQuery_Selection

// MultiSelectQueryMetricsQuerySelectionList is a selection for a multi select query.
type MultiSelectQueryMetricsQuerySelectionList = variables.MultiSelect_Query_MetricsQuery_Selection_List

// MultiSelectQueryMetricsQuerySelectionListSelection is a selection for a multi select query.
type MultiSelectQueryMetricsQuerySelectionListSelection = variables.MultiSelect_Query_MetricsQuery_Selection_ListSelection

// MultiSelectQueryMetricsQueryOperatorEquals is a selection for a multi select query.
type MultiSelectQueryMetricsQueryOperatorEquals = variables.MultiSelect_Query_MetricsQuery_Operator_Equals

// MultiSelectQueryMetricsQueryEquals is a selection for a multi select query.
type MultiSelectQueryMetricsQueryEquals = variables.MultiSelect_Query_MetricsQuery_Equals

// MultiSelectQueryMetricsQueryOperator is a selection for a multi select query.
type MultiSelectQueryMetricsQueryOperator = variables.MultiSelect_Query_MetricsQuery_Operator

// MultiSelectQueryMetricsQueryOperatorNotEquals is a selection for a multi select query.
type MultiSelectQueryMetricsQueryOperatorNotEquals = variables.MultiSelect_Query_MetricsQuery_Operator_NotEquals

// MultiSelectQueryMetricsQueryNotEquals is a selection for a multi select query.
// MultiSelectQueryMetricsQueryNotEquals is a metrics query operator.
type MultiSelectQueryMetricsQueryNotEquals = variables.MultiSelect_Query_MetricsQuery_NotEquals

// MultiSelectQuerySpansQuery is a spans query.
type MultiSelectQuerySpansQuery = variables.MultiSelect_Query_SpansQuery_

// MultiSelectQuerySpansQueryInner is a spans query inner type.
type MultiSelectQuerySpansQueryInner = variables.MultiSelect_Query_SpansQuery

// MultiSelectQuerySpansQueryType is a spans query type.
type MultiSelectQuerySpansQueryType = variables.MultiSelect_Query_SpansQuery_Type

// MultiSelectQuerySpansQueryTypeFieldName is a field name for the spans query.
type MultiSelectQuerySpansQueryTypeFieldName = variables.MultiSelect_Query_SpansQuery_Type_FieldName_

// MultiSelectQuerySpansQueryTypeFieldNameInner is a field name for the spans query.
type MultiSelectQuerySpansQueryTypeFieldNameInner = variables.MultiSelect_Query_SpansQuery_Type_FieldName

// MultiSelectQuerySpansQueryTypeFieldValue is a field value for the spans query.
type MultiSelectQuerySpansQueryTypeFieldValue = variables.MultiSelect_Query_SpansQuery_Type_FieldValue_

// MultiSelectQuerySpansQueryTypeFieldValueInner is a field value for the spans query.
type MultiSelectQuerySpansQueryTypeFieldValueInner = variables.MultiSelect_Query_SpansQuery_Type_FieldValue

// MultiSelectQueryLogsQueryTypeFieldName is a field name for the logs query.
type MultiSelectQueryLogsQueryTypeFieldName = variables.MultiSelect_Query_LogsQuery_Type_FieldName_

// MultiSelectQueryLogsQueryTypeFieldNameInner is a field name for the logs query.
type MultiSelectQueryLogsQueryTypeFieldNameInner = variables.MultiSelect_Query_LogsQuery_Type_FieldName

// MultiSelectQueryLogsQueryTypeFieldValue is a field value for the logs query.
type MultiSelectQueryLogsQueryTypeFieldValue = variables.MultiSelect_Query_LogsQuery_Type_FieldValue_

// MultiSelectQueryLogsQueryTypeFieldValueInner is a field value for the logs query.
type MultiSelectQueryLogsQueryTypeFieldValueInner = variables.MultiSelect_Query_LogsQuery_Type_FieldValue

// MultiSelectQueryMetricsQuery is a metrics query.
type MultiSelectQueryMetricsQuery = variables.MultiSelect_Query_MetricsQuery_

// MultiSelectQueryMetricsQueryInner is a metrics query.
type MultiSelectQueryMetricsQueryInner = variables.MultiSelect_Query_MetricsQuery

// MultiSelectQueryMetricsQueryType is a metrics query.
type MultiSelectQueryMetricsQueryType = variables.MultiSelect_Query_MetricsQuery_Type

// MultiSelectQueryMetricsQueryTypeMetricName is a field name for the metrics query.
type MultiSelectQueryMetricsQueryTypeMetricName = variables.MultiSelect_Query_MetricsQuery_Type_MetricName_

// MultiSelectQueryMetricsQueryTypeMetricNameInner is a field name for the metrics query.
type MultiSelectQueryMetricsQueryTypeMetricNameInner = variables.MultiSelect_Query_MetricsQuery_Type_MetricName

// MultiSelectQueryMetricsQueryTypeLabelValue is a field value for the metrics query.
type MultiSelectQueryMetricsQueryTypeLabelValue = variables.MultiSelect_Query_MetricsQuery_Type_LabelValue_

// MultiSelectQueryMetricsQueryTypeLabelValueInner is a field value for the metrics query.
type MultiSelectQueryMetricsQueryTypeLabelValueInner = variables.MultiSelect_Query_MetricsQuery_Type_LabelValue

// MultiSelectQueryMetricsQueryTypeLabelName is a field value for the metrics query.
type MultiSelectQueryMetricsQueryTypeLabelName = variables.MultiSelect_Query_MetricsQuery_Type_LabelName_

// MultiSelectQueryMetricsQueryTypeLabelNameInner is a field value for the metrics query.
type MultiSelectQueryMetricsQueryTypeLabelNameInner = variables.MultiSelect_Query_MetricsQuery_Type_LabelName

// MultiSelectQueryMetricsQueryMetricsLabelFilter is a metrics query label filter.
type MultiSelectQueryMetricsQueryMetricsLabelFilter = variables.MultiSelect_Query_MetricsQuery_MetricsLabelFilter

// GaugeMetricsQuery is a gauge metrics query.
type GaugeMetricsQuery = widgets.Gauge_MetricsQuery

// DashboardMetricsFilter is a metrics filter.
type DashboardMetricsFilter = filters.Filter_MetricsFilter

// DashboardFilterOperatorEquals is a filter operator.
type DashboardFilterOperatorEquals = filters.Filter_Operator_Equals

// DashboardFilterEquals is a filter operator: equals.
type DashboardFilterEquals = filters.Filter_Equals

// DashboardFilterEqualsSelection is a filter equals selection.
type DashboardFilterEqualsSelection = filters.Filter_Equals_Selection

// DashboardFilterEqualsSelectionList is a filter equals selection list.
type DashboardFilterEqualsSelectionList = filters.Filter_Equals_Selection_List

// DashboardFilterEqualsSelectionListSelection is a filter equals selection list selection.
type DashboardFilterEqualsSelectionListSelection = filters.Filter_Equals_Selection_ListSelection

// DashboardFilterEqualsSelectionAll is a filter equals selection all.
type DashboardFilterEqualsSelectionAll = filters.Filter_Equals_Selection_All

// DashboardFilterEqualsSelectionAllSelection is a filter equals selection all selection.
type DashboardFilterEqualsSelectionAllSelection = filters.Filter_Equals_Selection_AllSelection

// DashboardFilterOperator is a filter operator.
type DashboardFilterOperator = filters.Filter_Operator

// DashboardFilterOperatorNotEquals is a filter operator.
type DashboardFilterOperatorNotEquals = filters.Filter_Operator_NotEquals

// DashboardFilterNotEquals is a filter operator not equals.
type DashboardFilterNotEquals = filters.Filter_NotEquals

// DashboardFilterNotEqualsSelection is a filter not equals selection.
type DashboardFilterNotEqualsSelection = filters.Filter_NotEquals_Selection

// DashboardFilterNotEqualsSelectionList is a filter not equals selection list.
type DashboardFilterNotEqualsSelectionList = filters.Filter_NotEquals_Selection_List

// DashboardFilterNotEqualsSelectionListSelection is a filter not equals selection list selection.
type DashboardFilterNotEqualsSelectionListSelection = filters.Filter_NotEquals_Selection_ListSelection

// DashboardPromQlQuery is a promQL query.
// Deprecated: Use DashboardPromQLQuery (upper case 'L') instead.
type DashboardPromQlQuery = common.PromQlQuery

// DashboardPromQLQuery is a promQL query.
type DashboardPromQLQuery = common.PromQlQuery

// DashboardLuceneQuery is a lucene query.
type DashboardLuceneQuery = common.LuceneQuery

// GaugeLogsQuery is a logs query for gauges.
type GaugeLogsQuery = widgets.Gauge_LogsQuery

// LogsAggregation is a logs aggregation.
type LogsAggregation = common.LogsAggregation

// LogsAggregationCount is a logs aggregation type.
type LogsAggregationCount = common.LogsAggregation_Count_

// LogsAggregationCountInner is a logs aggregation type.
type LogsAggregationCountInner = common.LogsAggregation_Count

// LogsAggregationCountDistinct is a logs aggregation type.
type LogsAggregationCountDistinct = common.LogsAggregation_CountDistinct_

// LogsAggregationCountDistinctInner is a logs aggregation type.
type LogsAggregationCountDistinctInner = common.LogsAggregation_CountDistinct

// LogsAggregationSum is a logs aggregation type.
type LogsAggregationSum = common.LogsAggregation_Sum_

// LogsAggregationSumInner is a logs aggregation type.
type LogsAggregationSumInner = common.LogsAggregation_Sum

// LogsAggregationAverage is a logs aggregation type.
type LogsAggregationAverage = common.LogsAggregation_Average_

// LogsAggregationAverageInner is a logs aggregation type.
type LogsAggregationAverageInner = common.LogsAggregation_Average

// LogsAggregationMin is a logs aggregation type.
type LogsAggregationMin = common.LogsAggregation_Min_

// LogsAggregationMinInner is a logs aggregation type.
type LogsAggregationMinInner = common.LogsAggregation_Min

// LogsAggregationMax is a logs aggregation type.
type LogsAggregationMax = common.LogsAggregation_Max_

// LogsAggregationMaxInner is a logs aggregation type.
type LogsAggregationMaxInner = common.LogsAggregation_Max

// LogsAggregationPercentile is a logs aggregation type.
type LogsAggregationPercentile = common.LogsAggregation_Percentile_

// LogsAggregationPercentileInner is a logs aggregation type.
type LogsAggregationPercentileInner = common.LogsAggregation_Percentile

// AnnotationSource is an annotation variant.
type AnnotationSource = annotations.Annotation_Source

// AnnotationSourceLogs is an annotation variant.
type AnnotationSourceLogs = annotations.Annotation_Source_Logs

// AnnotationLogsSource is an annotation variant.
type AnnotationLogsSource = annotations.Annotation_LogsSource

// AnnotationLogsSourceStrategy is an annotation variant.
type AnnotationLogsSourceStrategy = annotations.Annotation_LogsSource_Strategy

// AnnotationSourceSpans is an annotation variant.
type AnnotationSourceSpans = annotations.Annotation_Source_Spans

// AnnotationSpansSource is an annotation variant.
type AnnotationSpansSource = annotations.Annotation_SpansSource

// AnnotationSpansSourceStrategy is an annotation variant.
type AnnotationSpansSourceStrategy = annotations.Annotation_SpansSource_Strategy

// AnnotationSourceMetrics is an annotation variant.
type AnnotationSourceMetrics = annotations.Annotation_Source_Metrics

// AnnotationMetricsSource is an annotation variant.
type AnnotationMetricsSource = annotations.Annotation_MetricsSource

// AnnotationMetricsSourceStrategy is an annotation variant.
type AnnotationMetricsSourceStrategy = annotations.Annotation_MetricsSource_Strategy

// AnnotationSpansSourceStrategyDurationInner is an annotation variant.
type AnnotationSpansSourceStrategyDurationInner = annotations.Annotation_SpansSource_Strategy_Duration

// AnnotationSpansSourceStrategyDuration is an annotation variant.
type AnnotationSpansSourceStrategyDuration = annotations.Annotation_SpansSource_Strategy_Duration_

// AnnotationSpansSourceStrategyRange is an annotation variant.
type AnnotationSpansSourceStrategyRange = annotations.Annotation_SpansSource_Strategy_Range_

// AnnotationSpansSourceStrategyRangeInner is an annotation variant.
type AnnotationSpansSourceStrategyRangeInner = annotations.Annotation_SpansSource_Strategy_Range

// AnnotationSpansSourceStrategyInstant is an annotation variant.
type AnnotationSpansSourceStrategyInstant = annotations.Annotation_SpansSource_Strategy_Instant_

// AnnotationSpansSourceStrategyInstantInner is an annotation variant.
type AnnotationSpansSourceStrategyInstantInner = annotations.Annotation_SpansSource_Strategy_Instant

// AnnotationLogsSourceStrategyDurationInner is an annotation variant.
type AnnotationLogsSourceStrategyDurationInner = annotations.Annotation_LogsSource_Strategy_Duration

// AnnotationLogsSourceStrategyDuration is an annotation variant.
type AnnotationLogsSourceStrategyDuration = annotations.Annotation_LogsSource_Strategy_Duration_

// AnnotationLogsSourceStrategyRange is an annotation variant.
type AnnotationLogsSourceStrategyRange = annotations.Annotation_LogsSource_Strategy_Range_

// AnnotationLogsSourceStrategyRangeInner is an annotation variant.
type AnnotationLogsSourceStrategyRangeInner = annotations.Annotation_LogsSource_Strategy_Range

// AnnotationLogsSourceStrategyInstant is an annotation variant.
type AnnotationLogsSourceStrategyInstant = annotations.Annotation_LogsSource_Strategy_Instant_

// AnnotationLogsSourceStrategyInstantInner is an annotation variant.
type AnnotationLogsSourceStrategyInstantInner = annotations.Annotation_LogsSource_Strategy_Instant

// AnnotationMetricsSourceStartTimeMetric is an annotation variant.
type AnnotationMetricsSourceStartTimeMetric = annotations.Annotation_MetricsSource_StartTimeMetric

// AnnotationMetricsSourceStrategyStartTimeMetric is an annotation variant.
type AnnotationMetricsSourceStrategyStartTimeMetric = annotations.Annotation_MetricsSource_Strategy_StartTimeMetric

// DashboardLayout is a dashboard layout type.
type DashboardLayout = ast.Layout

// DashboardSection is a dashboard layout type.
type DashboardSection = ast.Section

// DashboardSectionOptions is a dashboard layout type.
type DashboardSectionOptions = ast.SectionOptions

// DashboardSectionOptionsCustom is a customizable dashboard layout option.
type DashboardSectionOptionsCustom = ast.SectionOptions_Custom

// CustomSectionOptions is a type for customizing section options.
type CustomSectionOptions = ast.CustomSectionOptions

// DashboardSectionColor is a dashboard layout type.
type DashboardSectionColor = ast.SectionColor

// DashboardSectionColorPredefined is a dashboard layout type.
type DashboardSectionColorPredefined = ast.SectionColor_Predefined

// DashboardSectionColorPredefinedColor is a dashboard layout type.
type DashboardSectionColorPredefinedColor = ast.SectionPredefinedColor

var (
	// DashboardSectionPredefinedColorValueLookup is a map of predefined color values to strings.
	DashboardSectionPredefinedColorValueLookup = ast.SectionPredefinedColor_value
	// DashboardSectionPredefinedColorNameLookup is a map of predefined strings to color values.
	DashboardSectionPredefinedColorNameLookup = ast.SectionPredefinedColor_name
)

// DashboardTwoMinutes is a auto refresh setting
type DashboardTwoMinutes = ast.Dashboard_TwoMinutes

// DashboardAutoRefreshTwoMinutes is a auto refresh setting
type DashboardAutoRefreshTwoMinutes = ast.Dashboard_AutoRefreshTwoMinutes

// DashboardFiveMinutes is a auto refresh setting
type DashboardFiveMinutes = ast.Dashboard_FiveMinutes

// DashboardAutoRefreshFiveMinutes is a auto refresh setting
type DashboardAutoRefreshFiveMinutes = ast.Dashboard_AutoRefreshFiveMinutes

// DashboardOff is a auto refresh setting
type DashboardOff = ast.Dashboard_Off

// DashboardAutoRefreshOff is a auto refresh setting
type DashboardAutoRefreshOff = ast.Dashboard_AutoRefreshOff

// HorizontalBarChartYAxisViewBy is a type for the horizontal bar chart y axis.
type HorizontalBarChartYAxisViewBy = widgets.HorizontalBarChart_YAxisViewBy

// HorizontalBarChartYAxisViewByCategory is a type for the horizontal bar chart y axis.
type HorizontalBarChartYAxisViewByCategory = widgets.HorizontalBarChart_YAxisViewBy_Category

// HorizontalBarChartYAxisViewByValue is a type for the horizontal bar chart y axis.
type HorizontalBarChartYAxisViewByValue = widgets.HorizontalBarChart_YAxisViewBy_Value

// BarChartXAxis is a type for the  bar chart x axis.
type BarChartXAxis = widgets.BarChart_XAxis

// BarChartXAxisTime is a type for the  bar chart x axis.
type BarChartXAxisTime = widgets.BarChart_XAxis_Time

// BarChartXAxisValue is a type for the  bar chart x axis.
type BarChartXAxisValue = widgets.BarChart_XAxis_Value

// BarChartXAxisByTime is a type for the  bar chart x axis.
type BarChartXAxisByTime = widgets.BarChart_XAxis_XAxisByTime

// BarChartXAxisByValue is a type for the  bar chart x axis.
type BarChartXAxisByValue = widgets.BarChart_XAxis_XAxisByValue

// BarChartQuery is a type for bar chart data.
type BarChartQuery = widgets.BarChart_Query

// BarChartQueryLogs is a type for bar chart data.
type BarChartQueryLogs = widgets.BarChart_Query_Logs

// BarChartQueryMetrics is a type for bar chart data.
type BarChartQueryMetrics = widgets.BarChart_Query_Metrics

// BarChartQuerySpans is a type for bar chart data.
type BarChartQuerySpans = widgets.BarChart_Query_Spans

// BarChartQueryDataprime is a type for bar chart data.
type BarChartQueryDataprime = widgets.BarChart_Query_Dataprime

// HorizontalBarChartQuery is a type for bar chart data.
type HorizontalBarChartQuery = widgets.HorizontalBarChart_Query

// HorizontalBarChartQueryLogs is a type for bar chart data.
type HorizontalBarChartQueryLogs = widgets.HorizontalBarChart_Query_Logs

// HorizontalBarChartQueryMetrics is a type for bar chart data.
type HorizontalBarChartQueryMetrics = widgets.HorizontalBarChart_Query_Metrics

// HorizontalBarChartQuerySpans is a type for bar chart data.
type HorizontalBarChartQuerySpans = widgets.HorizontalBarChart_Query_Spans

// HorizontalBarChartQueryDataprime is a type for bar chart data.
type HorizontalBarChartQueryDataprime = widgets.HorizontalBarChart_Query_Dataprime

// HorizontalBarChartLogsQuery is a type for bar chart data.
type HorizontalBarChartLogsQuery = widgets.HorizontalBarChart_LogsQuery

// HorizontalBarChartMetricsQuery is a type for bar chart data.
type HorizontalBarChartMetricsQuery = widgets.HorizontalBarChart_MetricsQuery

// HorizontalBarChartSpansQuery is a type for bar chart data.
type HorizontalBarChartSpansQuery = widgets.HorizontalBarChart_SpansQuery

// HorizontalBarChartDataprimeQuery is a type for bar chart data.
type HorizontalBarChartDataprimeQuery = widgets.HorizontalBarChart_DataprimeQuery

// BarChartLogsQuery is a type for data in bar chart.
type BarChartLogsQuery = widgets.BarChart_LogsQuery

// BarChartMetricsQuery is a type for data in bar chart.
type BarChartMetricsQuery = widgets.BarChart_MetricsQuery

// BarChartSpansQuery is a type for data in bar chart.
type BarChartSpansQuery = widgets.BarChart_SpansQuery

// ObservationField is a type for data in bar chart.
type ObservationField = common.ObservationField

// DashboardsColorsBy is a type for dashboard coloring.
type DashboardsColorsBy = widgetsCommon.ColorsBy

// DashboardsColorsByStack is a type for dashboard coloring.
type DashboardsColorsByStack = widgetsCommon.ColorsBy_Stack

// DashboardsColorsByGroupBy is a type for dashboard coloring.
type DashboardsColorsByGroupBy = widgetsCommon.ColorsBy_GroupBy

// DashboardsColorsByAggregation is a type for dashboard coloring.
type DashboardsColorsByAggregation = widgetsCommon.ColorsBy_Aggregation

// DashboardsColorsByStackInner is a type for dashboard coloring.
type DashboardsColorsByStackInner = widgetsCommon.ColorsBy_ColorsByStack

// DashboardsColorsByGroupByInner is a type for dashboard coloring.
type DashboardsColorsByGroupByInner = widgetsCommon.ColorsBy_ColorsByGroupBy

// DashboardsColorsByAggregationInner is a type for dashboard coloring.
type DashboardsColorsByAggregationInner = widgetsCommon.ColorsBy_ColorsByAggregation

// PieChartQuery is a type for data for charts.
type PieChartQuery = widgets.PieChart_Query

// PieChartQueryDataprime is a type for data for charts.
type PieChartQueryDataprime = widgets.PieChart_Query_Dataprime

// PieChartDataprimeQuery is a type for data for charts.
type PieChartDataprimeQuery = widgets.PieChart_DataprimeQuery

// DashboardDataprimeQuery is a type for data for charts.
type DashboardDataprimeQuery = common.DataprimeQuery

// PieChartQuerySpans is a type for data for charts.
type PieChartQuerySpans = widgets.PieChart_Query_Spans

// PieChartSpansQuery is a type for data for charts.
type PieChartSpansQuery = widgets.PieChart_SpansQuery

// PieChartQueryMetrics is a type for data for charts.
type PieChartQueryMetrics = widgets.PieChart_Query_Metrics

// PieChartMetricsQuery is a type for data for charts.
type PieChartMetricsQuery = widgets.PieChart_MetricsQuery

// PieChartQueryLogs is a type for data for charts.
type PieChartQueryLogs = widgets.PieChart_Query_Logs

// PieChartLogsQuery is a type for data for charts.
type PieChartLogsQuery = widgets.PieChart_LogsQuery

// DashboardVariable is a type for data for charts.
type DashboardVariable = variables.Variable

// DashboardConstant is a type for data for charts.
type DashboardConstant = variables.Constant

// DashboardMultiSelect is a type for data for charts.
type DashboardMultiSelect = variables.MultiSelect

// DashboardMultiSelectSource is a type for data for charts.
type DashboardMultiSelectSource = variables.MultiSelect_Source

// DashboardMultiSelectSelectionAll is a type for data for charts.
type DashboardMultiSelectSelectionAll = variables.MultiSelect_Selection_All

// DashboardMultiSelectAllSelection is a type for data for charts.
type DashboardMultiSelectAllSelection = variables.MultiSelect_Selection_AllSelection

// DashboardMultiSelectSelectionList is a type for data for charts.
type DashboardMultiSelectSelectionList = variables.MultiSelect_Selection_List

// DashboardMultiSelectListSelection is a type for data for charts.
type DashboardMultiSelectListSelection = variables.MultiSelect_Selection_ListSelection

// DashboardMultiSelectSelection is a type for data for charts.
type DashboardMultiSelectSelection = variables.MultiSelect_Selection

// DashboardVariableDefinition is a type for data for charts.
type DashboardVariableDefinition = variables.Variable_Definition

// DashboardVariableDefinitionConstant is a type for data for charts.
type DashboardVariableDefinitionConstant = variables.Variable_Definition_Constant

// DashboardVariableDefinitionMultiSelect is a type for data for charts.
type DashboardVariableDefinitionMultiSelect = variables.Variable_Definition_MultiSelect

// LineChartQuery is a type for data for charts.
type LineChartQuery = widgets.LineChart_Query

// LineChartQueryDefinition is a type for data for charts.
type LineChartQueryDefinition = widgets.LineChart_QueryDefinition

// LineChartResolution is a type for data for charts.
type LineChartResolution = widgets.LineChart_Resolution

// LineChartQueryLogs is a type for data for charts.
type LineChartQueryLogs = widgets.LineChart_Query_Logs

// LineChartLogsQuery is a type for data for charts.
type LineChartLogsQuery = widgets.LineChart_LogsQuery

// LineChartQueryDataprime is a type for data for charts.
type LineChartQueryDataprime = widgets.LineChart_Query_Dataprime

// LineChartDataprimeQuery is the inner type for data for charts.
type LineChartDataprimeQuery = widgets.LineChart_DataprimeQuery

// LineChartQueryMetrics is a type for data for charts.
type LineChartQueryMetrics = widgets.LineChart_Query_Metrics

// LineChartMetricsQuery is a type for data for charts.
type LineChartMetricsQuery = widgets.LineChart_MetricsQuery

// LineChartQuerySpans is a type for data for charts.
type LineChartQuerySpans = widgets.LineChart_Query_Spans

// LineChartSpansQuery is a type for data for charts.
type LineChartSpansQuery = widgets.LineChart_SpansQuery

// DashboardDataTable is a type for dashboard charts.
type DashboardDataTable = widgets.DataTable

// DashboardDataTableQuery is a type for dashboard charts.
type DashboardDataTableQuery = widgets.DataTable_Query

// DashboardDataTableQueryDataprime is a type for dashboard charts.
type DashboardDataTableQueryDataprime = widgets.DataTable_Query_Dataprime

// DashboardDataTableDataprimeQuery is a type for dashboard charts.
type DashboardDataTableDataprimeQuery = widgets.DataTable_DataprimeQuery

// DashboardDataTableQueryMetrics is a type for dashboard charts.
type DashboardDataTableQueryMetrics = widgets.DataTable_Query_Metrics

// DashboardDataTableMetricsQuery is a type for dashboard charts.
type DashboardDataTableMetricsQuery = widgets.DataTable_MetricsQuery

// DashboardDataTableQueryLogs is a type for dashboard charts.
type DashboardDataTableQueryLogs = widgets.DataTable_Query_Logs

// DashboardDataTableLogsQuery is a type for dashboard charts.
type DashboardDataTableLogsQuery = widgets.DataTable_LogsQuery

// DashboardDataTableLogsQueryGrouping is a type for dashboard charts.
type DashboardDataTableLogsQueryGrouping = widgets.DataTable_LogsQuery_Grouping

// DashboardDataTableLogsQueryAggregation is a type for dashboard charts.
type DashboardDataTableLogsQueryAggregation = widgets.DataTable_LogsQuery_Aggregation

// DashboardDataTableQuerySpans is a type for dashboard charts.
type DashboardDataTableQuerySpans = widgets.DataTable_Query_Spans

// DashboardDataTableSpansQuery is a type for dashboard charts.
type DashboardDataTableSpansQuery = widgets.DataTable_SpansQuery

// DashboardDataTableSpansQueryGrouping is a type for dashboard charts.
type DashboardDataTableSpansQueryGrouping = widgets.DataTable_SpansQuery_Grouping

// DashboardDataTableSpansQueryAggregation is a type for dashboard charts.
type DashboardDataTableSpansQueryAggregation = widgets.DataTable_SpansQuery_Aggregation

// DashboardDataTableColumn is a type for dashboard charts.
type DashboardDataTableColumn = widgets.DataTable_Column

// TimeframeSelectRelative is a type for dashboard charts.
type TimeframeSelectRelative = common.TimeFrameSelect_RelativeTimeFrame

// TimeframeSelectAbsolute is a type for dashboard charts.
type TimeframeSelectAbsolute = common.TimeFrameSelect_AbsoluteTimeFrame

// TimeframeSelect is a type for dashboard charts.
type TimeframeSelect = common.TimeFrameSelect

// DashboardRelativeTimeFrame is a type for dashboard charts.
type DashboardRelativeTimeFrame = ast.Dashboard_RelativeTimeFrame

// DashboardAbsoluteTimeFrame is a type for dashboard charts.
type DashboardAbsoluteTimeFrame = ast.Dashboard_AbsoluteTimeFrame

// DashboardTimeFrame is a type for dashboard charts.
type DashboardTimeFrame = common.TimeFrame

// DashboardOrderingField is a type for dashboard charts.
type DashboardOrderingField = common.OrderingField

// LegendBy is a legend configuration option.
type LegendBy = widgetsCommon.LegendBy

// LegendBy values.
const (
	LegendByUnspecified = widgetsCommon.LegendBy_LEGEND_BY_UNSPECIFIED
	LegendByThresholds  = widgetsCommon.LegendBy_LEGEND_BY_THRESHOLDS
	LegendByGroups      = widgetsCommon.LegendBy_LEGEND_BY_GROUPS
)

// HexagonQuery is a query class for hexagons.
type HexagonQuery = widgets.Hexagon_Query

// HexagonDataprimeQuery is a dataprime query for hexagon widgets.
type HexagonDataprimeQuery = widgets.Hexagon_DataprimeQuery

// HexagonQueryDataprime is a dataprime query for hexagon widgets.
type HexagonQueryDataprime = widgets.Hexagon_Query_Dataprime

// HexagonLogsQuery is a logs query for hexagon widgets.
type HexagonLogsQuery = widgets.Hexagon_LogsQuery

// HexagonQueryLogs is a logs query for hexagon widgets.
type HexagonQueryLogs = widgets.Hexagon_Query_Logs

// HexagonMetricsQuery is a metrics query for hexagon widgets.
type HexagonMetricsQuery = widgets.Hexagon_MetricsQuery

// HexagonQueryMetrics is a metrics query for hexagon widgets.
type HexagonQueryMetrics = widgets.Hexagon_Query_Metrics

// HexagonSpansQuery is a spans query for hexagon widgets.
type HexagonSpansQuery = widgets.Hexagon_SpansQuery

// HexagonQuerySpans is a spans query for hexagon widgets.
type HexagonQuerySpans = widgets.Hexagon_Query_Spans

// HexagonSpansQueryGrouping is a spans query grouping for hexagon widgets.
type HexagonSpansQueryGrouping = widgets.Hexagon_SpansQuery

// Threshold is a hexagon threshold configuration option.
type Threshold = widgetsCommon.Threshold

// ThresholdType is a legend configuration option.
type ThresholdType = widgetsCommon.ThresholdType

// ThresholdType values.
const (
	ThresholdTypeUnspecified = widgetsCommon.ThresholdType_THRESHOLD_TYPE_UNSPECIFIED
	ThresholdTypeAbsolute    = widgetsCommon.ThresholdType_THRESHOLD_TYPE_ABSOLUTE
	ThresholdTypeRelative    = widgetsCommon.ThresholdType_THRESHOLD_TYPE_RELATIVE
)

// HexagonMetricAggregation is a type for aggregating values.
type HexagonMetricAggregation = widgetsCommon.Aggregation

// HexagonMetricAggregation values
const (
	HexagonMetricAggregationUnspecified = widgetsCommon.Aggregation_AGGREGATION_UNSPECIFIED
	HexagonMetricAggregationLast        = widgetsCommon.Aggregation_AGGREGATION_LAST
	HexagonMetricAggregationMin         = widgetsCommon.Aggregation_AGGREGATION_MIN
	HexagonMetricAggregationMax         = widgetsCommon.Aggregation_AGGREGATION_MAX
	HexagonMetricAggregationAvg         = widgetsCommon.Aggregation_AGGREGATION_AVG
	HexagonMetricAggregationSum         = widgetsCommon.Aggregation_AGGREGATION_SUM
)

// AssignDashboardFolderRequest is a request to assign a dashboard to a folder.
type AssignDashboardFolderRequest = services.AssignDashboardFolderRequest

const dashboardsFeatureGroupID = "dashboards"

// Create Creates a new dashboard.
func (d DashboardsClient) Create(ctx context.Context, req *CreateDashboardRequest) (*services.CreateDashboardResponse, error) {
	callProperties, err := d.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := services.NewDashboardsServiceClient(conn)

	response, err := client.CreateDashboard(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, CreateDashboardRPC, dashboardsFeatureGroupID)
	}
	return response, nil
}

// Get gets a dashboard.
func (d DashboardsClient) Get(ctx context.Context, req *GetDashboardRequest) (*services.GetDashboardResponse, error) {
	callProperties, err := d.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := services.NewDashboardsServiceClient(conn)

	response, err := client.GetDashboard(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, GetDashboardRPC, dashboardsFeatureGroupID)
	}
	return response, nil
}

// List lists all dashboards.
func (d DashboardsClient) List(ctx context.Context) (*services.GetDashboardCatalogResponse, error) {
	callProperties, err := d.callPropertiesCreator.GetUserLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := services.NewDashboardCatalogServiceClient(conn)

	response, err := client.GetDashboardCatalog(callProperties.Ctx, &services.GetDashboardCatalogRequest{}, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, services.DashboardCatalogService_GetDashboardCatalog_FullMethodName, dashboardsFeatureGroupID)
	}
	return response, nil
}

// Replace replaces a dashboard.
func (d DashboardsClient) Replace(ctx context.Context, req *ReplaceDashboardRequest) (*services.ReplaceDashboardResponse, error) {
	callProperties, err := d.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := services.NewDashboardsServiceClient(conn)
	response, err := client.ReplaceDashboard(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, services.DashboardsService_ReplaceDashboard_FullMethodName, dashboardsFeatureGroupID)
	}
	return response, nil
}

// Delete deletes a dashboard.
func (d DashboardsClient) Delete(ctx context.Context, req *DeleteDashboardRequest) (*services.DeleteDashboardResponse, error) {
	callProperties, err := d.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := services.NewDashboardsServiceClient(conn)

	response, err := client.DeleteDashboard(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, DeleteDashboardRPC, dashboardsFeatureGroupID)
	}
	return response, nil
}

// Pin pins a dashboard.
func (d DashboardsClient) Pin(ctx context.Context, req *PinDashboardRequest) (*services.PinDashboardResponse, error) {
	callProperties, err := d.callPropertiesCreator.GetUserLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := services.NewDashboardsServiceClient(conn)

	response, err := client.PinDashboard(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, PinDashboardRPC, dashboardsFeatureGroupID)
	}
	return response, nil
}

// Unpin unpins a dashboard.
func (d DashboardsClient) Unpin(ctx context.Context, req *UnpinDashboardRequest) (*services.UnpinDashboardResponse, error) {
	callProperties, err := d.callPropertiesCreator.GetUserLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := services.NewDashboardsServiceClient(conn)

	response, err := client.UnpinDashboard(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, UnpinDashboardRPC, dashboardsFeatureGroupID)
	}
	return response, nil
}

// AssignToFolder assigns a dashboard to a folder.
func (d DashboardsClient) AssignToFolder(ctx context.Context, req *AssignDashboardFolderRequest) (*services.AssignDashboardFolderResponse, error) {
	callProperties, err := d.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := services.NewDashboardsServiceClient(conn)

	response, err := client.AssignDashboardFolder(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, AssignDashboardFolderRPC, dashboardsFeatureGroupID)
	}
	return response, nil
}

// NewDashboardsClient creates a new DashboardsClient.
func NewDashboardsClient(c CallPropertiesCreator) *DashboardsClient {
	return &DashboardsClient{callPropertiesCreator: c}
}
