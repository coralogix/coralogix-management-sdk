package forward_compat_oneof

import (
	"encoding/json"
	"strings"
	"testing"
)

// When the server returns a list response containing an unknown oneOf variant
// (e.g. a brand-new alert kind shipped by the backend before the SDK knows
// about it), decoding must succeed and the unknown element must keep its
// sibling fields. The unknown oneOf field itself is left with all variant
// pointers nil — callers can detect this via GetActualInstance() == nil.
func TestForwardCompatOneOf_UnknownVariantInListKeepsAlertWithNilOneOf(t *testing.T) {
	payload := `{
		"alerts": [
			{"id": "alert-known", "alertProps": {"logFilter": "level:error"}},
			{"id": "alert-unknown", "alertProps": {"futureVariant": {"foo": "bar"}}}
		]
	}`

	var resp ListAlertsResponse
	if err := json.Unmarshal([]byte(payload), &resp); err != nil {
		t.Fatalf("expected unmarshal to succeed, got: %v", err)
	}

	if len(resp.Alerts) != 2 {
		t.Fatalf("expected 2 alerts, got %d", len(resp.Alerts))
	}

	known := resp.Alerts[0]
	if known.Id != "alert-known" {
		t.Fatalf("known alert id mismatch: %q", known.Id)
	}
	if known.AlertProps == nil || known.AlertProps.LogsImmediateProps == nil {
		t.Fatalf("expected known alert to decode LogsImmediateProps, got %+v", known.AlertProps)
	}
	if known.AlertProps.MetricThresholdProps != nil {
		t.Fatalf("expected MetricThresholdProps to remain nil")
	}

	unknown := resp.Alerts[1]
	if unknown.Id != "alert-unknown" {
		t.Fatalf("unknown alert id mismatch: %q", unknown.Id)
	}
	if unknown.AlertProps == nil {
		t.Fatalf("expected unknown alert to keep AlertProps wrapper (non-nil), got nil")
	}
	if unknown.AlertProps.GetActualInstance() != nil {
		t.Fatalf("expected unknown variant to leave GetActualInstance() == nil")
	}
	if unknown.AlertProps.LogsImmediateProps != nil || unknown.AlertProps.MetricThresholdProps != nil {
		t.Fatalf("expected all variant pointers to remain nil for unknown variant")
	}
}

// Regression guard: an all-known payload must still resolve variants exactly
// as before this change.
func TestForwardCompatOneOf_AllKnownVariantsUnchanged(t *testing.T) {
	payload := `{
		"alerts": [
			{"id": "a1", "alertProps": {"logFilter": "level:error"}},
			{"id": "a2", "alertProps": {"metricFilter": "rate(http_requests_total)"}}
		]
	}`

	var resp ListAlertsResponse
	if err := json.Unmarshal([]byte(payload), &resp); err != nil {
		t.Fatalf("expected unmarshal to succeed, got: %v", err)
	}
	if len(resp.Alerts) != 2 {
		t.Fatalf("expected 2 alerts, got %d", len(resp.Alerts))
	}

	if resp.Alerts[0].AlertProps == nil || resp.Alerts[0].AlertProps.LogsImmediateProps == nil {
		t.Fatal("expected a1 to resolve to LogsImmediateProps")
	}
	if resp.Alerts[1].AlertProps == nil || resp.Alerts[1].AlertProps.MetricThresholdProps == nil {
		t.Fatal("expected a2 to resolve to MetricThresholdProps")
	}
}

// Decoding a *single* unknown-variant oneOf (not nested in a list) must also
// succeed cleanly — the fix applies at the oneOf level and is symmetric across
// GET-by-id and LIST endpoints.
func TestForwardCompatOneOf_StandaloneUnknownVariantDecodesAsNil(t *testing.T) {
	payload := `{"futureVariant": {"foo": "bar"}}`

	var props AlertProps
	if err := json.Unmarshal([]byte(payload), &props); err != nil {
		t.Fatalf("expected unmarshal to succeed, got: %v", err)
	}
	if props.GetActualInstance() != nil {
		t.Fatalf("expected GetActualInstance() == nil for unknown variant")
	}
	if props.LogsImmediateProps != nil || props.MetricThresholdProps != nil {
		t.Fatalf("expected all variant pointers to remain nil")
	}
}
