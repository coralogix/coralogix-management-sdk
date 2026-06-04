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
		name      string
		configure func(map[string]interface{})
		paths     [][]interface{}
	}{
		{
			name: "mixed widget definitions and queries",
			configure: func(raw map[string]interface{}) {
				setDashboardWidgets(raw, []interface{}{
					dashboardWidget("line", "line", map[string]interface{}{
						"lineChart": map[string]interface{}{
							"queryDefinitions": []interface{}{
								map[string]interface{}{"id": "line-metrics", "query": map[string]interface{}{"metrics": map[string]interface{}{"promqlQuery": map[string]interface{}{"value": "up"}}}},
								map[string]interface{}{"id": "line-logs", "query": map[string]interface{}{"logs": map[string]interface{}{"aggregations": []interface{}{map[string]interface{}{"count": map[string]interface{}{}}}}}},
							},
						},
					}),
					dashboardWidget("bar", "bar", map[string]interface{}{
						"barChart": map[string]interface{}{
							"query":    map[string]interface{}{"metrics": map[string]interface{}{"promqlQuery": map[string]interface{}{"value": "up"}}},
							"colorsBy": map[string]interface{}{"aggregation": map[string]interface{}{}},
							"xAxis":    map[string]interface{}{"time": map[string]interface{}{}},
						},
					}),
					dashboardWidget("table", "table", map[string]interface{}{
						"dataTable": map[string]interface{}{
							"query": map[string]interface{}{"logs": map[string]interface{}{"aggregations": []interface{}{map[string]interface{}{"count": map[string]interface{}{}}}}},
						},
					}),
					dashboardWidget("dynamic", "dynamic", map[string]interface{}{
						"dynamic": map[string]interface{}{
							"queryDefinitions": []interface{}{map[string]interface{}{"id": "dynamic-spans", "query": map[string]interface{}{"spans": map[string]interface{}{}}}},
							"visualization":    map[string]interface{}{"table": map[string]interface{}{}},
						},
					}),
					dashboardWidget("gauge", "gauge", map[string]interface{}{"gauge": map[string]interface{}{"query": map[string]interface{}{"dataprime": map[string]interface{}{}}}}),
					dashboardWidget("hexagon", "hexagon", map[string]interface{}{"hexagon": map[string]interface{}{"query": map[string]interface{}{"logs": map[string]interface{}{}}}}),
					dashboardWidget("horizontal", "horizontal", map[string]interface{}{"horizontalBarChart": map[string]interface{}{"query": map[string]interface{}{"spans": map[string]interface{}{}}, "yAxisViewBy": map[string]interface{}{"category": map[string]interface{}{}}}}),
					dashboardWidget("markdown", "markdown", map[string]interface{}{"markdown": map[string]interface{}{"markdownText": "OpenAPI dashboard fixture"}}),
					dashboardWidget("pie", "pie", map[string]interface{}{"pieChart": map[string]interface{}{"query": map[string]interface{}{"metrics": map[string]interface{}{"promqlQuery": map[string]interface{}{"value": "up"}}}}}),
				})
			},
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
			name: "dashboard variables and filters",
			configure: func(raw map[string]interface{}) {
				raw["variables"] = []interface{}{
					map[string]interface{}{"name": "env", "definition": map[string]interface{}{"constant": map[string]interface{}{"value": "production"}}},
					map[string]interface{}{"name": "service", "definition": map[string]interface{}{"multiSelect": map[string]interface{}{"selection": map[string]interface{}{"list": map[string]interface{}{"values": []interface{}{"api", "worker"}}}, "selected": []interface{}{"api"}}}},
				}
				raw["variablesV2"] = []interface{}{
					map[string]interface{}{"name": "query", "source": map[string]interface{}{"query": map[string]interface{}{"logsQuery": map[string]interface{}{}}}, "value": map[string]interface{}{"singleString": map[string]interface{}{"value": map[string]interface{}{"value": "api", "label": "API"}}}},
					map[string]interface{}{"name": "static", "source": map[string]interface{}{"static": map[string]interface{}{"values": []interface{}{map[string]interface{}{"value": "api", "label": "API"}}}}, "value": map[string]interface{}{"multiString": map[string]interface{}{"list": map[string]interface{}{"values": []interface{}{map[string]interface{}{"value": map[string]interface{}{"value": "api", "label": "API"}}}}}}},
					map[string]interface{}{"name": "textbox", "source": map[string]interface{}{"textbox": map[string]interface{}{}}, "value": map[string]interface{}{"regex": map[string]interface{}{"value": map[string]interface{}{"value": "api.*", "label": "API services"}}}},
				}
				raw["filters"] = []interface{}{
					map[string]interface{}{"displayName": "logs", "source": map[string]interface{}{"logs": map[string]interface{}{}}},
					map[string]interface{}{"displayName": "metrics", "source": map[string]interface{}{"metrics": map[string]interface{}{}}},
					map[string]interface{}{"displayName": "spans", "source": map[string]interface{}{"spans": map[string]interface{}{}}},
				}
			},
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
			name: "dashboard actions and annotations",
			configure: func(raw map[string]interface{}) {
				raw["actions"] = []interface{}{
					map[string]interface{}{"name": "runbook", "definition": map[string]interface{}{"customAction": map[string]interface{}{"url": "https://example.com/runbook"}}},
					map[string]interface{}{"name": "related", "definition": map[string]interface{}{"goToDashboardAction": map[string]interface{}{"dashboardId": "00000000-0000-4000-8000-000000000002"}}},
				}
				raw["annotations"] = []interface{}{
					map[string]interface{}{"name": "logs", "scope": map[string]interface{}{"allWidgets": map[string]interface{}{}}, "source": map[string]interface{}{"logs": map[string]interface{}{}}},
					map[string]interface{}{"name": "metrics", "scope": map[string]interface{}{"specificWidgets": map[string]interface{}{}}, "source": map[string]interface{}{"metrics": map[string]interface{}{}}},
					map[string]interface{}{"name": "manual", "source": map[string]interface{}{"manual": map[string]interface{}{}}},
					map[string]interface{}{"name": "dataprime", "source": map[string]interface{}{"dataprime": map[string]interface{}{}}},
					map[string]interface{}{"name": "spans", "source": map[string]interface{}{"spans": map[string]interface{}{}}},
				}
			},
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
			name: "dynamic visualizations",
			configure: func(raw map[string]interface{}) {
				widgets := make([]interface{}, 0)
				for _, visualization := range []string{
					"gauge", "geomap", "heatmap", "hexagonBins", "horizontalBars",
					"horizontalBarsMulti", "pieChart", "stat", "statCard", "table",
					"timeSeriesBars", "timeSeriesLines", "timeSeriesLinesMulti",
					"verticalBars", "verticalBarsMulti",
				} {
					value := map[string]interface{}{}
					if visualization == "heatmap" {
						value["preset"] = "HEATMAP_COLOR_PRESET_BLUE"
					}
					widgets = append(widgets, dashboardWidget(visualization, visualization, map[string]interface{}{
						"dynamic": map[string]interface{}{
							"queryDefinitions": []interface{}{map[string]interface{}{"id": visualization, "query": map[string]interface{}{"logs": map[string]interface{}{}}}},
							"visualization":    map[string]interface{}{visualization: value},
						},
					}))
				}
				setDashboardWidgets(raw, widgets)
			},
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
			raw := dashboardFixtureMap(t)
			raw["name"] = "dashboard OpenAPI " + tt.name
			raw["description"] = "dashboard OpenAPI full object variant fixture"
			tt.configure(raw)

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

func setDashboardWidgets(raw map[string]interface{}, widgets []interface{}) {
	raw["layout"] = map[string]interface{}{
		"sections": []interface{}{
			map[string]interface{}{
				"id": map[string]interface{}{"value": "section"},
				"rows": []interface{}{
					map[string]interface{}{
						"id":         map[string]interface{}{"value": "row"},
						"appearance": map[string]interface{}{"height": 12},
						"widgets":    widgets,
					},
				},
			},
		},
	}
}

func dashboardWidget(id, title string, definition map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"id":         map[string]interface{}{"value": id},
		"title":      title,
		"definition": definition,
		"appearance": map[string]interface{}{"width": 0},
	}
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
