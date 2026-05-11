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

// Package forwardcompat exercises forward-compatible decoding behavior
// against the *real* generated alert_definitions_service package, not synthetic
// compat fixtures. These tests guard against regressions where a server-side
// addition of a new oneOf variant would crash older SDK clients.
package forwardcompat

import (
	"encoding/json"
	"testing"

	alerts "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/alert_definitions_service"
)

// A list response that mixes a known alert kind with a server-side new alert
// kind must decode cleanly. The unknown alert is kept; its AlertDefProperties
// wrapper exists with all variant pointers nil so callers can detect "I don't
// understand this one" via GetActualInstance() == nil.
func TestListAlertDefsResponse_UnknownVariantKeepsAlertWithNilOneOf(t *testing.T) {
	payload := `{
		"alertDefs": [
			{
				"id": "alert-known",
				"alertDefProperties": {
					"name": "known-alert",
					"logsImmediate": {}
				}
			},
			{
				"id": "alert-unknown",
				"alertDefProperties": {
					"name": "future-alert",
					"futureKindProps": {"some": "field"}
				}
			}
		]
	}`

	var resp alerts.ListAlertDefsResponse
	if err := json.Unmarshal([]byte(payload), &resp); err != nil {
		t.Fatalf("expected unmarshal to succeed for mixed payload, got: %v", err)
	}

	if got := len(resp.AlertDefs); got != 2 {
		t.Fatalf("expected 2 alert defs, got %d", got)
	}

	known := resp.AlertDefs[0]
	if known.Id == nil || *known.Id != "alert-known" {
		t.Fatalf("known alert id mismatch: %+v", known.Id)
	}
	if known.AlertDefProperties == nil {
		t.Fatal("expected known alert to have non-nil AlertDefProperties wrapper")
	}
	if known.AlertDefProperties.GetActualInstance() == nil {
		t.Fatal("expected known alert to resolve to a variant (got GetActualInstance == nil)")
	}
	if known.AlertDefProperties.AlertDefPropertiesLogsImmediate == nil {
		t.Fatal("expected known alert to resolve to AlertDefPropertiesLogsImmediate")
	}

	unknown := resp.AlertDefs[1]
	if unknown.Id == nil || *unknown.Id != "alert-unknown" {
		t.Fatalf("unknown alert id mismatch: %+v", unknown.Id)
	}
	if unknown.AlertDefProperties == nil {
		t.Fatal("expected unknown alert to keep non-nil AlertDefProperties wrapper")
	}
	if unknown.AlertDefProperties.GetActualInstance() != nil {
		t.Fatalf("expected unknown variant to leave GetActualInstance() == nil, got %T",
			unknown.AlertDefProperties.GetActualInstance())
	}
}

// Regression guard: a single standalone unknown-variant AlertDefProperties
// payload (the GET-by-id case) must also decode without error.
func TestAlertDefProperties_StandaloneUnknownVariantDecodesAsNil(t *testing.T) {
	payload := `{"name": "future-alert", "futureKindProps": {"some": "field"}}`

	var props alerts.AlertDefProperties
	if err := json.Unmarshal([]byte(payload), &props); err != nil {
		t.Fatalf("expected unmarshal to succeed, got: %v", err)
	}
	if props.GetActualInstance() != nil {
		t.Fatalf("expected GetActualInstance() == nil for unknown variant, got %T",
			props.GetActualInstance())
	}
}

// Regression guard: an all-known list response still resolves variants
// correctly. Ensures the no-match path didn't accidentally silence valid
// match paths.
func TestListAlertDefsResponse_AllKnownVariantsUnchanged(t *testing.T) {
	payload := `{
		"alertDefs": [
			{"id": "a1", "alertDefProperties": {"name": "n1", "logsImmediate": {}}},
			{"id": "a2", "alertDefProperties": {"name": "n2", "metricThreshold": {"metricFilter": {}, "rules": []}}}
		]
	}`

	var resp alerts.ListAlertDefsResponse
	if err := json.Unmarshal([]byte(payload), &resp); err != nil {
		t.Fatalf("expected unmarshal to succeed, got: %v", err)
	}
	if len(resp.AlertDefs) != 2 {
		t.Fatalf("expected 2 alerts, got %d", len(resp.AlertDefs))
	}
	if resp.AlertDefs[0].AlertDefProperties == nil ||
		resp.AlertDefs[0].AlertDefProperties.AlertDefPropertiesLogsImmediate == nil {
		t.Fatal("expected a1 to resolve to AlertDefPropertiesLogsImmediate")
	}
	if resp.AlertDefs[1].AlertDefProperties == nil ||
		resp.AlertDefs[1].AlertDefProperties.AlertDefPropertiesMetricThreshold == nil {
		t.Fatal("expected a2 to resolve to AlertDefPropertiesMetricThreshold")
	}
}
