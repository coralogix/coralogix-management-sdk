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
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	dashboards "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/dashboard_service"
)

func TestDashboardFullObjectVariantFixtures(t *testing.T) {
	tests := []struct {
		name    string
		fixture string
	}{
		{
			name:    "mixed widget definitions and queries",
			fixture: "dashboard_variant_mixed_widgets.json",
		},
		{
			name:    "dashboard variables and filters",
			fixture: "dashboard_variant_variables_filters.json",
		},
		{
			name:    "dashboard actions and annotations",
			fixture: "dashboard_variant_actions_annotations.json",
		},
		{
			name:    "dynamic visualizations",
			fixture: "dashboard_variant_dynamic_visualizations.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dashboard := dashboardVariantFixture(t, tt.fixture)
			require.NotNil(t, dashboard.GetActualInstance())
		})
	}
}

func dashboardVariantFixture(t *testing.T, fixture string) dashboards.Dashboard {
	t.Helper()

	data, err := os.ReadFile(fixture)
	require.NoError(t, err)

	var dashboard dashboards.Dashboard
	require.NoError(t, json.Unmarshal(data, &dashboard))
	return dashboard
}
