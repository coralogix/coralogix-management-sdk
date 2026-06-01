// Copyright 2026 Coralogix Ltd.
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
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	dashboards "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/dashboard_service"
)

func TestDashboardVariantFixtures(t *testing.T) {
	timeFrames := []struct {
		field      string
		value      interface{}
		variant    string
		otherField string
	}{
		{
			field:      "relativeTimeFrame",
			value:      "900s",
			variant:    "RelativeTimeFrame",
			otherField: "absoluteTimeFrame",
		},
		{
			field: "absoluteTimeFrame",
			value: map[string]string{
				"from": "2026-05-28T10:00:00Z",
				"to":   "2026-05-28T11:00:00Z",
			},
			variant:    "AbsoluteTimeFrame",
			otherField: "relativeTimeFrame",
		},
	}

	refreshes := []struct {
		field   string
		variant string
	}{
		{field: "off", variant: "Off"},
		{field: "oneMinute", variant: "OneMinute"},
		{field: "twoMinutes", variant: "TwoMinutes"},
		{field: "fiveMinutes", variant: "FiveMinutes"},
		{field: "fifteenMinutes", variant: "FifteenMinutes"},
	}

	folders := []struct {
		name       string
		configure  func(map[string]interface{})
		assertions func(*testing.T, map[string]interface{})
	}{
		{
			name:      "root",
			configure: func(map[string]interface{}) {},
			assertions: func(t *testing.T, raw map[string]interface{}) {
				t.Helper()
				require.NotContains(t, raw, "folderId")
				require.NotContains(t, raw, "folderPath")
			},
		},
		{
			name: "folderId",
			configure: func(raw map[string]interface{}) {
				raw["folderId"] = map[string]string{"value": "00000000-0000-4000-8000-000000000001"}
			},
			assertions: func(t *testing.T, raw map[string]interface{}) {
				t.Helper()
				require.Contains(t, raw, "folderId")
				require.NotContains(t, raw, "folderPath")
			},
		},
		{
			name: "folderPath",
			configure: func(raw map[string]interface{}) {
				raw["folderPath"] = map[string][]string{"segments": {"team", "dashboards"}}
			},
			assertions: func(t *testing.T, raw map[string]interface{}) {
				t.Helper()
				require.NotContains(t, raw, "folderId")
				require.Contains(t, raw, "folderPath")
			},
		},
	}

	for _, timeFrame := range timeFrames {
		for _, refresh := range refreshes {
			for _, folder := range folders {
				name := fmt.Sprintf("%s/%s/%s", timeFrame.field, refresh.field, folder.name)
				t.Run(name, func(t *testing.T) {
					raw := minimalDashboardVariantJSON(name)
					raw[timeFrame.field] = timeFrame.value
					raw[refresh.field] = map[string]interface{}{}
					folder.configure(raw)

					data, err := json.Marshal(raw)
					require.NoError(t, err)

					var dashboard dashboards.Dashboard
					require.NoError(t, json.Unmarshal(data, &dashboard))

					actual := dashboard.GetActualInstance()
					require.NotNil(t, actual)
					require.Equal(t, "Dashboard"+refresh.variant+timeFrame.variant, dashboardVariantName(actual))

					roundTripData, err := json.Marshal(dashboard)
					require.NoError(t, err)

					var roundTrip map[string]interface{}
					require.NoError(t, json.Unmarshal(roundTripData, &roundTrip))
					require.Contains(t, roundTrip, timeFrame.field)
					require.NotContains(t, roundTrip, timeFrame.otherField)
					require.Contains(t, roundTrip, refresh.field)
					folder.assertions(t, roundTrip)
				})
			}
		}
	}
}

type generatedOneOf interface {
	GetActualInstance() interface{}
}

func TestDashboardNestedOneOfFixtures(t *testing.T) {
	tests := []struct {
		name    string
		payload string
		target  func() generatedOneOf
		want    string
	}{
		{
			name:    "widget line chart",
			payload: `{"lineChart":{"queryDefinitions":[]}}`,
			target:  func() generatedOneOf { return &dashboards.WidgetDefinition{} },
			want:    "WidgetDefinitionLineChart",
		},
		{
			name:    "widget bar chart",
			payload: `{"barChart":{}}`,
			target:  func() generatedOneOf { return &dashboards.WidgetDefinition{} },
			want:    "WidgetDefinitionBarChart",
		},
		{
			name:    "widget data table",
			payload: `{"dataTable":{}}`,
			target:  func() generatedOneOf { return &dashboards.WidgetDefinition{} },
			want:    "WidgetDefinitionDataTable",
		},
		{
			name:    "widget dynamic",
			payload: `{"dynamic":{}}`,
			target:  func() generatedOneOf { return &dashboards.WidgetDefinition{} },
			want:    "WidgetDefinitionDynamic",
		},
		{
			name:    "widget gauge",
			payload: `{"gauge":{}}`,
			target:  func() generatedOneOf { return &dashboards.WidgetDefinition{} },
			want:    "WidgetDefinitionGauge",
		},
		{
			name:    "widget hexagon",
			payload: `{"hexagon":{}}`,
			target:  func() generatedOneOf { return &dashboards.WidgetDefinition{} },
			want:    "WidgetDefinitionHexagon",
		},
		{
			name:    "widget horizontal bar chart",
			payload: `{"horizontalBarChart":{}}`,
			target:  func() generatedOneOf { return &dashboards.WidgetDefinition{} },
			want:    "WidgetDefinitionHorizontalBarChart",
		},
		{
			name:    "widget markdown",
			payload: `{"markdown":{}}`,
			target:  func() generatedOneOf { return &dashboards.WidgetDefinition{} },
			want:    "WidgetDefinitionMarkdown",
		},
		{
			name:    "widget pie chart",
			payload: `{"pieChart":{}}`,
			target:  func() generatedOneOf { return &dashboards.WidgetDefinition{} },
			want:    "WidgetDefinitionPieChart",
		},
		{
			name:    "bar chart dataprime query",
			payload: `{"dataprime":{}}`,
			target:  func() generatedOneOf { return &dashboards.BarChartQuery{} },
			want:    "BarChartQueryDataprime",
		},
		{
			name:    "bar chart logs query",
			payload: `{"logs":{}}`,
			target:  func() generatedOneOf { return &dashboards.BarChartQuery{} },
			want:    "BarChartQueryLogs",
		},
		{
			name:    "bar chart metrics query",
			payload: `{"metrics":{}}`,
			target:  func() generatedOneOf { return &dashboards.BarChartQuery{} },
			want:    "BarChartQueryMetrics",
		},
		{
			name:    "bar chart spans query",
			payload: `{"spans":{}}`,
			target:  func() generatedOneOf { return &dashboards.BarChartQuery{} },
			want:    "BarChartQuerySpans",
		},
		{
			name:    "dynamic dataprime query",
			payload: `{"dataprime":{}}`,
			target:  func() generatedOneOf { return &dashboards.DynamicQuery{} },
			want:    "DynamicQueryDataprime",
		},
		{
			name:    "dynamic logs query",
			payload: `{"logs":{}}`,
			target:  func() generatedOneOf { return &dashboards.DynamicQuery{} },
			want:    "DynamicQueryLogs",
		},
		{
			name:    "dynamic metrics query",
			payload: `{"metrics":{}}`,
			target:  func() generatedOneOf { return &dashboards.DynamicQuery{} },
			want:    "DynamicQueryMetrics",
		},
		{
			name:    "dynamic spans query",
			payload: `{"spans":{}}`,
			target:  func() generatedOneOf { return &dashboards.DynamicQuery{} },
			want:    "DynamicQuerySpans",
		},
		{
			name:    "gauge dataprime query",
			payload: `{"dataprime":{}}`,
			target:  func() generatedOneOf { return &dashboards.GaugeQuery{} },
			want:    "GaugeQueryDataprime",
		},
		{
			name:    "gauge logs query",
			payload: `{"logs":{}}`,
			target:  func() generatedOneOf { return &dashboards.GaugeQuery{} },
			want:    "GaugeQueryLogs",
		},
		{
			name:    "gauge metrics query",
			payload: `{"metrics":{}}`,
			target:  func() generatedOneOf { return &dashboards.GaugeQuery{} },
			want:    "GaugeQueryMetrics",
		},
		{
			name:    "gauge spans query",
			payload: `{"spans":{}}`,
			target:  func() generatedOneOf { return &dashboards.GaugeQuery{} },
			want:    "GaugeQuerySpans",
		},
		{
			name:    "hexagon dataprime query",
			payload: `{"dataprime":{}}`,
			target:  func() generatedOneOf { return &dashboards.HexagonQuery{} },
			want:    "HexagonQueryDataprime",
		},
		{
			name:    "hexagon logs query",
			payload: `{"logs":{}}`,
			target:  func() generatedOneOf { return &dashboards.HexagonQuery{} },
			want:    "HexagonQueryLogs",
		},
		{
			name:    "hexagon metrics query",
			payload: `{"metrics":{}}`,
			target:  func() generatedOneOf { return &dashboards.HexagonQuery{} },
			want:    "HexagonQueryMetrics",
		},
		{
			name:    "hexagon spans query",
			payload: `{"spans":{}}`,
			target:  func() generatedOneOf { return &dashboards.HexagonQuery{} },
			want:    "HexagonQuerySpans",
		},
		{
			name:    "horizontal bar chart dataprime query",
			payload: `{"dataprime":{}}`,
			target:  func() generatedOneOf { return &dashboards.HorizontalBarChartQuery{} },
			want:    "HorizontalBarChartQueryDataprime",
		},
		{
			name:    "horizontal bar chart logs query",
			payload: `{"logs":{}}`,
			target:  func() generatedOneOf { return &dashboards.HorizontalBarChartQuery{} },
			want:    "HorizontalBarChartQueryLogs",
		},
		{
			name:    "horizontal bar chart metrics query",
			payload: `{"metrics":{}}`,
			target:  func() generatedOneOf { return &dashboards.HorizontalBarChartQuery{} },
			want:    "HorizontalBarChartQueryMetrics",
		},
		{
			name:    "horizontal bar chart spans query",
			payload: `{"spans":{}}`,
			target:  func() generatedOneOf { return &dashboards.HorizontalBarChartQuery{} },
			want:    "HorizontalBarChartQuerySpans",
		},
		{
			name:    "pie chart dataprime query",
			payload: `{"dataprime":{}}`,
			target:  func() generatedOneOf { return &dashboards.PieChartQuery{} },
			want:    "PieChartQueryDataprime",
		},
		{
			name:    "pie chart logs query",
			payload: `{"logs":{}}`,
			target:  func() generatedOneOf { return &dashboards.PieChartQuery{} },
			want:    "PieChartQueryLogs",
		},
		{
			name:    "pie chart metrics query",
			payload: `{"metrics":{}}`,
			target:  func() generatedOneOf { return &dashboards.PieChartQuery{} },
			want:    "PieChartQueryMetrics",
		},
		{
			name:    "pie chart spans query",
			payload: `{"spans":{}}`,
			target:  func() generatedOneOf { return &dashboards.PieChartQuery{} },
			want:    "PieChartQuerySpans",
		},
		{
			name:    "line chart metrics query",
			payload: `{"metrics":{"promqlQuery":{"value":"up"}}}`,
			target:  func() generatedOneOf { return &dashboards.LineChartQuery{} },
			want:    "LineChartQueryMetrics",
		},
		{
			name:    "line chart logs query",
			payload: `{"logs":{"luceneQuery":{"value":"severity:warning"}}}`,
			target:  func() generatedOneOf { return &dashboards.LineChartQuery{} },
			want:    "LineChartQueryLogs",
		},
		{
			name:    "line chart dataprime query",
			payload: `{"dataprime":{}}`,
			target:  func() generatedOneOf { return &dashboards.LineChartQuery{} },
			want:    "LineChartQueryDataprime",
		},
		{
			name:    "line chart spans query",
			payload: `{"spans":{}}`,
			target:  func() generatedOneOf { return &dashboards.LineChartQuery{} },
			want:    "LineChartQuerySpans",
		},
		{
			name:    "data table logs query",
			payload: `{"logs":{"aggregations":[{"count":{}}]}}`,
			target:  func() generatedOneOf { return &dashboards.DataTableQuery{} },
			want:    "DataTableQueryLogs",
		},
		{
			name:    "data table metrics query",
			payload: `{"metrics":{"promqlQuery":{"value":"up"}}}`,
			target:  func() generatedOneOf { return &dashboards.DataTableQuery{} },
			want:    "DataTableQueryMetrics",
		},
		{
			name:    "data table dataprime query",
			payload: `{"dataprime":{}}`,
			target:  func() generatedOneOf { return &dashboards.DataTableQuery{} },
			want:    "DataTableQueryDataprime",
		},
		{
			name:    "data table spans query",
			payload: `{"spans":{}}`,
			target:  func() generatedOneOf { return &dashboards.DataTableQuery{} },
			want:    "DataTableQuerySpans",
		},
		{
			name:    "logs aggregation average",
			payload: `{"average":{"field":"duration"}}`,
			target:  func() generatedOneOf { return &dashboards.LogsAggregation{} },
			want:    "LogsAggregationAverage",
		},
		{
			name:    "logs aggregation count",
			payload: `{"count":{}}`,
			target:  func() generatedOneOf { return &dashboards.LogsAggregation{} },
			want:    "LogsAggregationCount",
		},
		{
			name:    "logs aggregation count distinct",
			payload: `{"countDistinct":{"field":"user"}}`,
			target:  func() generatedOneOf { return &dashboards.LogsAggregation{} },
			want:    "LogsAggregationCountDistinct",
		},
		{
			name:    "logs aggregation max",
			payload: `{"max":{"field":"duration"}}`,
			target:  func() generatedOneOf { return &dashboards.LogsAggregation{} },
			want:    "LogsAggregationMax",
		},
		{
			name:    "logs aggregation min",
			payload: `{"min":{"field":"duration"}}`,
			target:  func() generatedOneOf { return &dashboards.LogsAggregation{} },
			want:    "LogsAggregationMin",
		},
		{
			name:    "logs aggregation sum",
			payload: `{"sum":{"field":"duration"}}`,
			target:  func() generatedOneOf { return &dashboards.LogsAggregation{} },
			want:    "LogsAggregationSum",
		},
		{
			name:    "logs aggregation percentile",
			payload: `{"percentile":{"field":"duration","percent":99}}`,
			target:  func() generatedOneOf { return &dashboards.LogsAggregation{} },
			want:    "LogsAggregationPercentile",
		},
		{
			name:    "filter equals",
			payload: `{"equals":{"selection":{"list":{"values":["prod"]}}}}`,
			target:  func() generatedOneOf { return &dashboards.FilterOperator{} },
			want:    "FilterOperatorEquals",
		},
		{
			name:    "filter not equals",
			payload: `{"notEquals":{"selection":{"list":{"values":["dev"]}}}}`,
			target:  func() generatedOneOf { return &dashboards.FilterOperator{} },
			want:    "FilterOperatorNotEquals",
		},
		{
			name:    "equals selection all",
			payload: `{"all":{}}`,
			target:  func() generatedOneOf { return &dashboards.EqualsSelection{} },
			want:    "EqualsSelectionAll",
		},
		{
			name:    "equals selection list",
			payload: `{"list":{"values":["prod","staging"]}}`,
			target:  func() generatedOneOf { return &dashboards.EqualsSelection{} },
			want:    "EqualsSelectionList",
		},
		{
			name:    "time frame relative",
			payload: `{"relativeTimeFrame":"900s"}`,
			target:  func() generatedOneOf { return &dashboards.TimeFrameSelect{} },
			want:    "TimeFrameSelectRelativeTimeFrame",
		},
		{
			name:    "time frame absolute",
			payload: `{"absoluteTimeFrame":{"from":"2026-05-28T10:00:00Z","to":"2026-05-28T11:00:00Z"}}`,
			target:  func() generatedOneOf { return &dashboards.TimeFrameSelect{} },
			want:    "TimeFrameSelectAbsoluteTimeFrame",
		},
		{
			name:    "variable constant",
			payload: `{"constant":{"value":"production"}}`,
			target:  func() generatedOneOf { return &dashboards.VariableDefinition{} },
			want:    "VariableDefinitionConstant",
		},
		{
			name:    "variable multi select",
			payload: `{"multiSelect":{"selection":{"all":{}},"selected":["production"]}}`,
			target:  func() generatedOneOf { return &dashboards.VariableDefinition{} },
			want:    "VariableDefinitionMultiSelect",
		},
		{
			name:    "multi select selection all",
			payload: `{"all":{}}`,
			target:  func() generatedOneOf { return &dashboards.MultiSelectSelection{} },
			want:    "MultiSelectSelectionAll",
		},
		{
			name:    "multi select selection list",
			payload: `{"list":{"values":["production"]}}`,
			target:  func() generatedOneOf { return &dashboards.MultiSelectSelection{} },
			want:    "MultiSelectSelectionList",
		},
		{
			name:    "multi select query logs",
			payload: `{"logsQuery":{}}`,
			target:  func() generatedOneOf { return &dashboards.MultiSelectQuery{} },
			want:    "MultiSelectQueryLogsQuery",
		},
		{
			name:    "multi select query metrics",
			payload: `{"metricsQuery":{}}`,
			target:  func() generatedOneOf { return &dashboards.MultiSelectQuery{} },
			want:    "MultiSelectQueryMetricsQuery",
		},
		{
			name:    "multi select query spans",
			payload: `{"spansQuery":{}}`,
			target:  func() generatedOneOf { return &dashboards.MultiSelectQuery{} },
			want:    "MultiSelectQuerySpansQuery",
		},
		{
			name:    "multi select source constant list",
			payload: `{"constantList":{"values":["production"]}}`,
			target:  func() generatedOneOf { return &dashboards.MultiSelectSource{} },
			want:    "MultiSelectSourceConstantList",
		},
		{
			name:    "multi select source logs path",
			payload: `{"logsPath":{"path":"service"}}`,
			target:  func() generatedOneOf { return &dashboards.MultiSelectSource{} },
			want:    "MultiSelectSourceLogsPath",
		},
		{
			name:    "multi select source metric label",
			payload: `{"metricLabel":{"label":"service"}}`,
			target:  func() generatedOneOf { return &dashboards.MultiSelectSource{} },
			want:    "MultiSelectSourceMetricLabel",
		},
		{
			name:    "multi select source query",
			payload: `{"query":{}}`,
			target:  func() generatedOneOf { return &dashboards.MultiSelectSource{} },
			want:    "MultiSelectSourceQuery",
		},
		{
			name:    "multi select source span field",
			payload: `{"spanField":{"field":"service"}}`,
			target:  func() generatedOneOf { return &dashboards.MultiSelectSource{} },
			want:    "MultiSelectSourceSpanField",
		},
		{
			name:    "variable source query",
			payload: `{"query":{"logsQuery":{}}}`,
			target:  func() generatedOneOf { return &dashboards.VariableSourceV2{} },
			want:    "VariableSourceV2Query",
		},
		{
			name:    "variable source static",
			payload: `{"static":{"values":[{"value":"production","label":"Production"}]}}`,
			target:  func() generatedOneOf { return &dashboards.VariableSourceV2{} },
			want:    "VariableSourceV2Static",
		},
		{
			name:    "variable source textbox",
			payload: `{"textbox":{}}`,
			target:  func() generatedOneOf { return &dashboards.VariableSourceV2{} },
			want:    "VariableSourceV2Textbox",
		},
		{
			name:    "variable query source dataprime",
			payload: `{"dataprimeQuery":{}}`,
			target:  func() generatedOneOf { return &dashboards.VariableSourceV2QuerySource{} },
			want:    "VariableSourceV2QuerySourceDataprimeQuery",
		},
		{
			name:    "variable query source logs",
			payload: `{"logsQuery":{}}`,
			target:  func() generatedOneOf { return &dashboards.VariableSourceV2QuerySource{} },
			want:    "VariableSourceV2QuerySourceLogsQuery",
		},
		{
			name:    "variable query source metrics",
			payload: `{"metricsQuery":{}}`,
			target:  func() generatedOneOf { return &dashboards.VariableSourceV2QuerySource{} },
			want:    "VariableSourceV2QuerySourceMetricsQuery",
		},
		{
			name:    "variable query source spans",
			payload: `{"spansQuery":{}}`,
			target:  func() generatedOneOf { return &dashboards.VariableSourceV2QuerySource{} },
			want:    "VariableSourceV2QuerySourceSpansQuery",
		},
		{
			name:    "variable value interval",
			payload: `{"interval":{"value":{"value":"5m","label":"5m"}}}`,
			target:  func() generatedOneOf { return &dashboards.VariableValueV2{} },
			want:    "VariableValueV2Interval",
		},
		{
			name:    "variable value lucene",
			payload: `{"lucene":{"value":{"value":"severity:error","label":"Errors"}}}`,
			target:  func() generatedOneOf { return &dashboards.VariableValueV2{} },
			want:    "VariableValueV2Lucene",
		},
		{
			name:    "variable value multi string",
			payload: `{"multiString":{"list":{"values":[{"value":{"value":"production","label":"Production"}}]}}}`,
			target:  func() generatedOneOf { return &dashboards.VariableValueV2{} },
			want:    "VariableValueV2MultiString",
		},
		{
			name:    "variable value regex",
			payload: `{"regex":{"value":{"value":"prod.*","label":"Production"}}}`,
			target:  func() generatedOneOf { return &dashboards.VariableValueV2{} },
			want:    "VariableValueV2Regex",
		},
		{
			name:    "variable value single numeric",
			payload: `{"singleNumeric":{"value":{"value":42,"label":"42"}}}`,
			target:  func() generatedOneOf { return &dashboards.VariableValueV2{} },
			want:    "VariableValueV2SingleNumeric",
		},
		{
			name:    "variable value single string",
			payload: `{"singleString":{"value":{"value":"production","label":"Production"}}}`,
			target:  func() generatedOneOf { return &dashboards.VariableValueV2{} },
			want:    "VariableValueV2SingleString",
		},
		{
			name:    "metric query type metric name",
			payload: `{"metricName":{"metricRegex":"process_.*"}}`,
			target:  func() generatedOneOf { return &dashboards.QueryMetricsQueryType{} },
			want:    "QueryMetricsQueryTypeMetricNameVariant",
		},
		{
			name:    "metric query type label name",
			payload: `{"labelName":{"metricRegex":"process_.*"}}`,
			target:  func() generatedOneOf { return &dashboards.QueryMetricsQueryType{} },
			want:    "QueryMetricsQueryTypeLabelNameVariant",
		},
		{
			name:    "metric query type label value",
			payload: `{"labelValue":{}}`,
			target:  func() generatedOneOf { return &dashboards.QueryMetricsQueryType{} },
			want:    "QueryMetricsQueryTypeLabelValueVariant",
		},
		{
			name:    "source metric query type metric name",
			payload: `{"metricName":{"metricRegex":"process_.*"}}`,
			target:  func() generatedOneOf { return &dashboards.QuerySourceMetricsQueryType{} },
			want:    "QuerySourceMetricsQueryTypeMetricNameVariant",
		},
		{
			name:    "source metric query type label name",
			payload: `{"labelName":{"metricRegex":"process_.*"}}`,
			target:  func() generatedOneOf { return &dashboards.QuerySourceMetricsQueryType{} },
			want:    "QuerySourceMetricsQueryTypeLabelNameVariant",
		},
		{
			name:    "source metric query type label value",
			payload: `{"labelValue":{}}`,
			target:  func() generatedOneOf { return &dashboards.QuerySourceMetricsQueryType{} },
			want:    "QuerySourceMetricsQueryTypeLabelValueVariant",
		},
		{
			name:    "logs query type field name",
			payload: `{"fieldName":{"logRegex":"kubernetes.*"}}`,
			target:  func() generatedOneOf { return &dashboards.QueryLogsQueryType{} },
			want:    "QueryLogsQueryTypeFieldNameVariant",
		},
		{
			name:    "logs query type field value",
			payload: `{"fieldValue":{}}`,
			target:  func() generatedOneOf { return &dashboards.QueryLogsQueryType{} },
			want:    "QueryLogsQueryTypeFieldValueVariant",
		},
		{
			name:    "spans query type field name",
			payload: `{"fieldName":{"spanRegex":"http.*"}}`,
			target:  func() generatedOneOf { return &dashboards.QuerySpansQueryType{} },
			want:    "QuerySpansQueryTypeFieldNameVariant",
		},
		{
			name:    "spans query type field value",
			payload: `{"fieldValue":{}}`,
			target:  func() generatedOneOf { return &dashboards.QuerySpansQueryType{} },
			want:    "QuerySpansQueryTypeFieldValueVariant",
		},
		{
			name:    "metrics operator equals",
			payload: `{"equals":{"selection":{"list":{"values":[{"stringValue":"production"}]}}}}`,
			target:  func() generatedOneOf { return &dashboards.QueryMetricsQueryOperator{} },
			want:    "QueryMetricsQueryOperatorEquals",
		},
		{
			name:    "metrics operator not equals",
			payload: `{"notEquals":{"selection":{"list":{"values":[{"stringValue":"staging"}]}}}}`,
			target:  func() generatedOneOf { return &dashboards.QueryMetricsQueryOperator{} },
			want:    "QueryMetricsQueryOperatorNotEquals",
		},
		{
			name:    "metrics string or variable string",
			payload: `{"stringValue":"production"}`,
			target:  func() generatedOneOf { return &dashboards.QueryMetricsQueryStringOrVariable{} },
			want:    "QueryMetricsQueryStringOrVariableStringValue",
		},
		{
			name:    "metrics string or variable name",
			payload: `{"variableName":"env"}`,
			target:  func() generatedOneOf { return &dashboards.QueryMetricsQueryStringOrVariable{} },
			want:    "QueryMetricsQueryStringOrVariableVariableName",
		},
		{
			name:    "visualization gauge",
			payload: `{"gauge":{}}`,
			target:  func() generatedOneOf { return &dashboards.Visualization{} },
			want:    "VisualizationGaugeVariant",
		},
		{
			name:    "visualization geomap",
			payload: `{"geomap":{}}`,
			target:  func() generatedOneOf { return &dashboards.Visualization{} },
			want:    "VisualizationGeomap",
		},
		{
			name:    "visualization heatmap",
			payload: `{"heatmap":{"preset":"HEATMAP_COLOR_PRESET_BLUE"}}`,
			target:  func() generatedOneOf { return &dashboards.Visualization{} },
			want:    "VisualizationHeatmap",
		},
		{
			name:    "visualization hexagon bins",
			payload: `{"hexagonBins":{}}`,
			target:  func() generatedOneOf { return &dashboards.Visualization{} },
			want:    "VisualizationHexagonBins",
		},
		{
			name:    "visualization horizontal bars",
			payload: `{"horizontalBars":{}}`,
			target:  func() generatedOneOf { return &dashboards.Visualization{} },
			want:    "VisualizationHorizontalBars",
		},
		{
			name:    "visualization horizontal bars multi",
			payload: `{"horizontalBarsMulti":{}}`,
			target:  func() generatedOneOf { return &dashboards.Visualization{} },
			want:    "VisualizationHorizontalBarsMulti",
		},
		{
			name:    "visualization pie chart",
			payload: `{"pieChart":{}}`,
			target:  func() generatedOneOf { return &dashboards.Visualization{} },
			want:    "VisualizationPieChartVariant",
		},
		{
			name:    "visualization stat",
			payload: `{"stat":{}}`,
			target:  func() generatedOneOf { return &dashboards.Visualization{} },
			want:    "VisualizationStat",
		},
		{
			name:    "visualization stat card",
			payload: `{"statCard":{}}`,
			target:  func() generatedOneOf { return &dashboards.Visualization{} },
			want:    "VisualizationStatCard",
		},
		{
			name:    "visualization table",
			payload: `{"table":{}}`,
			target:  func() generatedOneOf { return &dashboards.Visualization{} },
			want:    "VisualizationTable",
		},
		{
			name:    "visualization time series bars",
			payload: `{"timeSeriesBars":{}}`,
			target:  func() generatedOneOf { return &dashboards.Visualization{} },
			want:    "VisualizationTimeSeriesBars",
		},
		{
			name:    "visualization time series lines",
			payload: `{"timeSeriesLines":{}}`,
			target:  func() generatedOneOf { return &dashboards.Visualization{} },
			want:    "VisualizationTimeSeriesLines",
		},
		{
			name:    "visualization time series lines multi",
			payload: `{"timeSeriesLinesMulti":{}}`,
			target:  func() generatedOneOf { return &dashboards.Visualization{} },
			want:    "VisualizationTimeSeriesLinesMulti",
		},
		{
			name:    "visualization vertical bars",
			payload: `{"verticalBars":{}}`,
			target:  func() generatedOneOf { return &dashboards.Visualization{} },
			want:    "VisualizationVerticalBars",
		},
		{
			name:    "visualization vertical bars multi",
			payload: `{"verticalBarsMulti":{}}`,
			target:  func() generatedOneOf { return &dashboards.Visualization{} },
			want:    "VisualizationVerticalBarsMulti",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			target := tt.target()
			require.NoError(t, json.Unmarshal([]byte(tt.payload), target))
			actual := target.GetActualInstance()
			require.NotNil(t, actual)
			require.Equal(t, tt.want, dashboardVariantName(actual))

			data, err := json.Marshal(target)
			require.NoError(t, err)
			require.NotEqual(t, "null", string(data))
		})
	}
}

func minimalDashboardVariantJSON(name string) map[string]interface{} {
	return map[string]interface{}{
		"name":        name,
		"description": "dashboard OpenAPI variant fixture",
		"layout": map[string]interface{}{
			"sections": []interface{}{},
		},
	}
}

func dashboardVariantName(v interface{}) string {
	return strings.TrimPrefix(reflect.TypeOf(v).String(), "*dashboard_service.")
}
