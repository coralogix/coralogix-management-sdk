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
	"os"
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
				raw["folderPath"] = map[string][]string{"segments": []string{"team", "dashboards"}}
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
					raw := dashboardFixtureMap(t)
					raw["name"] = name
					raw["description"] = "dashboard OpenAPI variant fixture"
					for _, candidate := range timeFrames {
						delete(raw, candidate.field)
					}
					for _, candidate := range refreshes {
						delete(raw, candidate.field)
					}
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

func TestDashboardFullObjectVariantFixtures(t *testing.T) {
	tests := []struct {
		name    string
		fixture string
		paths   [][]interface{}
	}{
		{
			name:    "mixed widget definitions and queries",
			fixture: "dashboard_variant_mixed_widgets.json",
			paths: [][]interface{}{
				{"layout", "sections", 0, "rows", 0, "widgets", 0, "definition", "lineChart", "queryDefinitions", 0, "query", "metrics"},
				{"layout", "sections", 0, "rows", 0, "widgets", 1, "definition", "barChart", "colorsBy", "aggregation"},
				{"layout", "sections", 0, "rows", 0, "widgets", 2, "definition", "dataTable", "query", "logs"},
				{"layout", "sections", 0, "rows", 0, "widgets", 3, "definition", "dynamic", "visualization", "table"},
				{"layout", "sections", 0, "rows", 0, "widgets", 4, "definition", "gauge", "query", "dataprime"},
				{"layout", "sections", 0, "rows", 0, "widgets", 5, "definition", "hexagon", "query", "logs"},
				{"layout", "sections", 0, "rows", 0, "widgets", 6, "definition", "horizontalBarChart", "yAxisViewBy", "category"},
				{"layout", "sections", 0, "rows", 0, "widgets", 7, "definition", "markdown", "markdownText"},
				{"layout", "sections", 0, "rows", 0, "widgets", 8, "definition", "pieChart", "query", "metrics"},
			},
		},
		{
			name:    "dashboard variables and filters",
			fixture: "dashboard_variant_variables_filters.json",
			paths: [][]interface{}{
				{"variables", 0, "definition", "constant"},
				{"variables", 1, "definition", "multiSelect", "selection", "list"},
				{"variablesV2", 0, "source", "query", "logsQuery"},
				{"variablesV2", 1, "source", "static", "values", 0, "value"},
				{"variablesV2", 2, "source", "textbox"},
				{"filters", 0, "source", "logs"},
				{"filters", 1, "source", "metrics"},
				{"filters", 2, "source", "spans"},
			},
		},
		{
			name:    "dashboard actions and annotations",
			fixture: "dashboard_variant_actions_annotations.json",
			paths: [][]interface{}{
				{"actions", 0, "definition", "customAction"},
				{"actions", 1, "definition", "goToDashboardAction"},
				{"annotations", 0, "scope", "allWidgets"},
				{"annotations", 0, "source", "logs"},
				{"annotations", 1, "scope", "specificWidgets"},
				{"annotations", 1, "source", "metrics"},
				{"annotations", 2, "source", "manual"},
				{"annotations", 3, "source", "dataprime"},
				{"annotations", 4, "source", "spans"},
			},
		},
		{
			name:    "dynamic visualizations",
			fixture: "dashboard_variant_dynamic_visualizations.json",
			paths: [][]interface{}{
				{"layout", "sections", 0, "rows", 0, "widgets", 0, "definition", "dynamic", "visualization", "gauge"},
				{"layout", "sections", 0, "rows", 0, "widgets", 1, "definition", "dynamic", "visualization", "geomap"},
				{"layout", "sections", 0, "rows", 0, "widgets", 2, "definition", "dynamic", "visualization", "heatmap"},
				{"layout", "sections", 0, "rows", 0, "widgets", 9, "definition", "dynamic", "visualization", "table"},
				{"layout", "sections", 0, "rows", 0, "widgets", 14, "definition", "dynamic", "visualization", "verticalBarsMulti"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			raw := dashboardVariantFixtureMap(t, tt.fixture)
			roundTrip := roundTripDashboardFixture(t, raw)
			for _, path := range tt.paths {
				requireDashboardPath(t, roundTrip, path...)
			}
		})
	}
}

func dashboardVariantName(v interface{}) string {
	return strings.TrimPrefix(reflect.TypeOf(v).String(), "*dashboard_service.")
}

func dashboardVariantFixtureMap(t *testing.T, fixture string) map[string]interface{} {
	t.Helper()

	data, err := os.ReadFile(fixture)
	require.NoError(t, err)

	var raw map[string]interface{}
	require.NoError(t, json.Unmarshal(data, &raw))
	return raw
}

func roundTripDashboardFixture(t *testing.T, raw map[string]interface{}) map[string]interface{} {
	t.Helper()

	data, err := json.Marshal(raw)
	require.NoError(t, err)

	var dashboard dashboards.Dashboard
	require.NoError(t, json.Unmarshal(data, &dashboard))
	require.NotNil(t, dashboard.GetActualInstance())

	data, err = json.Marshal(dashboard)
	require.NoError(t, err)

	var roundTrip map[string]interface{}
	require.NoError(t, json.Unmarshal(data, &roundTrip))
	return roundTrip
}

func requireDashboardPath(t *testing.T, raw map[string]interface{}, path ...interface{}) interface{} {
	t.Helper()

	var current interface{} = raw
	for _, segment := range path {
		switch key := segment.(type) {
		case string:
			asMap, ok := current.(map[string]interface{})
			require.Truef(t, ok, "expected object at %v", path)
			next, ok := asMap[key]
			require.Truef(t, ok, "missing key %q at %v", key, path)
			require.NotNilf(t, next, "nil value for key %q at %v", key, path)
			current = next
		case int:
			asSlice, ok := current.([]interface{})
			require.Truef(t, ok, "expected array at %v", path)
			require.Lessf(t, key, len(asSlice), "array index out of range at %v", path)
			current = asSlice[key]
		default:
			t.Fatalf("unsupported path segment %T in %v", segment, path)
		}
	}
	return current
}
