// Copyright 2026 Coralogix Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package forwardcompat verifies that list responses produced by the OpenAPI
// generator drop list elements whose oneOf variant the SDK does not recognize,
// instead of failing the entire response. This protects customers on older SDK
// versions when the server adds a new oneOf variant.
package forwardcompat

import (
	"encoding/json"
	"testing"

	alerts "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/alert_definitions_service"
)

// listAlertsPayload simulates a server response from a future server that has
// added a new alert kind (`futureUnknownAlert`). Sandwiched between two known
// variants — a logsImmediate and a metricThreshold — it must be silently
// dropped, never poison the whole list.
const listAlertsPayload = `{
  "alertDefs": [
    {
      "id": "alert-known-logs",
      "alertVersionId": "v1",
      "alertDefProperties": {
        "name": "logs immediate",
        "enabled": true,
        "priority": "ALERT_DEF_PRIORITY_P3",
        "type": "ALERT_DEF_TYPE_LOGS_IMMEDIATE_OR_UNSPECIFIED",
        "logsImmediate": {}
      }
    },
    {
      "id": "alert-future-unknown",
      "alertVersionId": "v1",
      "alertDefProperties": {
        "name": "future unknown",
        "enabled": true,
        "priority": "ALERT_DEF_PRIORITY_P3",
        "type": "ALERT_DEF_TYPE_FUTURE_UNKNOWN",
        "futureUnknownAlert": {
          "windowMs": 60000,
          "shape": "spike"
        }
      }
    },
    {
      "id": "alert-known-metric",
      "alertVersionId": "v1",
      "alertDefProperties": {
        "name": "metric threshold",
        "enabled": true,
        "priority": "ALERT_DEF_PRIORITY_P3",
        "type": "ALERT_DEF_TYPE_METRIC_THRESHOLD",
        "metricThreshold": {}
      }
    }
  ],
  "pagination": {
    "totalSize": 3
  }
}`

// TestListAlertDefsResponse_DropsUnknownVariant simulates a server that has
// introduced a brand-new alert kind. The old SDK doesn't know the new kind and
// must drop just that element — the response must still decode successfully and
// must keep both known elements with their type-specific data intact.
func TestListAlertDefsResponse_DropsUnknownVariant(t *testing.T) {
	var resp alerts.ListAlertDefsResponse
	if err := json.Unmarshal([]byte(listAlertsPayload), &resp); err != nil {
		t.Fatalf("ListAlertDefsResponse failed to decode despite unknown variant: %v", err)
	}

	if got, want := len(resp.AlertDefs), 2; got != want {
		t.Fatalf("len(AlertDefs) = %d, want %d (unknown variant should be dropped)", got, want)
	}

	if id := resp.AlertDefs[0].GetId(); id != "alert-known-logs" {
		t.Errorf("AlertDefs[0].Id = %q, want %q", id, "alert-known-logs")
	}
	if id := resp.AlertDefs[1].GetId(); id != "alert-known-metric" {
		t.Errorf("AlertDefs[1].Id = %q, want %q", id, "alert-known-metric")
	}

	if props := resp.AlertDefs[0].AlertDefProperties; props == nil || props.AlertDefPropertiesLogsImmediate == nil {
		t.Errorf("known logs alert lost its variant data; AlertDefProperties=%+v", props)
	}
	if props := resp.AlertDefs[1].AlertDefProperties; props == nil || props.AlertDefPropertiesMetricThreshold == nil {
		t.Errorf("known metric alert lost its variant data; AlertDefProperties=%+v", props)
	}

	if resp.Pagination == nil || resp.Pagination.GetTotalSize() != 3 {
		t.Errorf("pagination.TotalSize = %v, want 3 — sibling fields should survive unknown elements", resp.Pagination)
	}
}

// TestListAlertDefsResponse_AllUnknownYieldsEmptyList — even if every alert is
// of an unknown type the response itself decodes; callers see an empty list
// instead of a hard error.
func TestListAlertDefsResponse_AllUnknownYieldsEmptyList(t *testing.T) {
	payload := `{
		"alertDefs": [
			{"id": "u1", "alertDefProperties": {"name": "x", "type": "ALERT_DEF_TYPE_FUTURE_X", "futureX": {}}},
			{"id": "u2", "alertDefProperties": {"name": "y", "type": "ALERT_DEF_TYPE_FUTURE_Y", "futureY": {}}}
		]
	}`
	var resp alerts.ListAlertDefsResponse
	if err := json.Unmarshal([]byte(payload), &resp); err != nil {
		t.Fatalf("ListAlertDefsResponse failed to decode all-unknown list: %v", err)
	}
	if len(resp.AlertDefs) != 0 {
		t.Fatalf("len(AlertDefs) = %d, want 0 — all-unknown should be dropped", len(resp.AlertDefs))
	}
}

// TestListAlertDefsResponse_AllKnownIsUnchanged — pure-known list keeps every element.
func TestListAlertDefsResponse_AllKnownIsUnchanged(t *testing.T) {
	payload := `{
		"alertDefs": [
			{"id": "a", "alertDefProperties": {"name": "n1", "type": "ALERT_DEF_TYPE_LOGS_IMMEDIATE_OR_UNSPECIFIED", "logsImmediate": {}}},
			{"id": "b", "alertDefProperties": {"name": "n2", "type": "ALERT_DEF_TYPE_METRIC_THRESHOLD", "metricThreshold": {}}}
		]
	}`
	var resp alerts.ListAlertDefsResponse
	if err := json.Unmarshal([]byte(payload), &resp); err != nil {
		t.Fatalf("decode failed on all-known payload: %v", err)
	}
	if len(resp.AlertDefs) != 2 {
		t.Fatalf("len(AlertDefs) = %d, want 2", len(resp.AlertDefs))
	}
}
