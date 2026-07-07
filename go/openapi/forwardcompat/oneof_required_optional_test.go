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

package forwardcompat

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	scheduler "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/alert_scheduler_rule_service"
	events "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/events_service"
	scopes "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/scopes_service"
)

func TestQueryParameterOneOfValidationErrorIsReturned(t *testing.T) {
	filter := scheduler.AlertSchedulerRuleServiceGetBulkAlertSchedulerRuleAlertSchedulerRulesIdsParameter{}
	filter.SetAlertSchedulerIds(scheduler.AlertSchedulerRuleServiceGetBulkAlertSchedulerRuleAlertSchedulerRulesIdsParameterAlertSchedulerIds{
		AlertSchedulerRuleIds: []string{"rule-id"},
	})
	filter.SetAlertSchedulerVersionIds(scheduler.AlertSchedulerRuleServiceGetBulkAlertSchedulerRuleAlertSchedulerRulesIdsParameterAlertSchedulerVersionIds{
		AlertSchedulerRuleVersionIds: []string{"version-id"},
	})

	client := scheduler.NewAPIClient(scheduler.NewConfiguration())
	_, _, err := client.AlertSchedulerRuleServiceAPI.
		AlertSchedulerRuleServiceGetBulkAlertSchedulerRule(context.Background()).
		AlertSchedulerRulesIds(filter).
		Execute()
	if err == nil {
		t.Fatal("expected invalid query oneOf filter to fail before sending the request")
	}
	if !strings.Contains(err.Error(), "at most one of [alertSchedulerIds, alertSchedulerVersionIds] may be set") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRequiredOneOf_ResponseUnknownFutureArmIsForwardCompatible(t *testing.T) {
	payload := `{
		"path": "service.name",
		"futureValues": {"value": "api"}
	}`

	var got scopes.FilterPathAndValues
	if err := json.Unmarshal([]byte(payload), &got); err != nil {
		t.Fatalf("expected response unmarshal to tolerate unknown future oneOf arm, got: %v", err)
	}

	if got.Path != "service.name" {
		t.Fatalf("path mismatch: %q", got.Path)
	}
	if got.Filters != nil || got.MultipleValues != nil {
		t.Fatalf("expected known oneOf arms to remain nil for unknown future arm: %+v", got)
	}
	if _, ok := got.AdditionalProperties["futureValues"]; !ok {
		t.Fatal("expected futureValues to be preserved in AdditionalProperties")
	}
}

func TestRequiredOneOf_RequestMarshalStillRequiresExactlyOneKnownArm(t *testing.T) {
	missingArm := scopes.FilterPathAndValues{Path: "service.name"}
	if _, err := json.Marshal(missingArm); err == nil {
		t.Fatal("expected request marshal to reject required oneOf with no known arm")
	}
}

func TestOptionalOneOf_ResponseUnsetAndUnknownFutureArmAreForwardCompatible(t *testing.T) {
	tests := []struct {
		name    string
		payload string
		unknown string
	}{
		{
			name:    "unset",
			payload: `{}`,
		},
		{
			name:    "future arm",
			payload: `{"futureEvent": {"id": "e1"}}`,
			unknown: "futureEvent",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got events.CxEventSingleOrMultiple
			if err := json.Unmarshal([]byte(tt.payload), &got); err != nil {
				t.Fatalf("expected optional oneOf response unmarshal to succeed, got: %v", err)
			}
			if got.SingleEvent != nil || got.MultipleEvents != nil {
				t.Fatalf("expected known optional oneOf arms to remain nil: %+v", got)
			}
			if tt.unknown != "" {
				if _, ok := got.AdditionalProperties[tt.unknown]; !ok {
					t.Fatalf("expected %s to be preserved in AdditionalProperties", tt.unknown)
				}
			}
		})
	}
}
