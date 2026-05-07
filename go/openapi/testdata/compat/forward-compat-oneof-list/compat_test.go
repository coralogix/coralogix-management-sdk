package forward_compat_oneof_list

import (
	"encoding/json"
	"testing"
)

// Simulates an old SDK reading a list response from a server that has
// introduced a brand-new oneOf variant ("traceRule"). The old SDK only
// knows logsRule and metricRule, so the unknown alert must be dropped
// while the rest of the list survives.
func TestForwardCompatOneOfList_DropsUnknownVariantAndKeepsRest(t *testing.T) {
	payload := []byte(`{
		"alerts": [
			{"id": "1", "props": {"kind": "logs", "logsRule": {"q": "error"}}},
			{"id": "2", "props": {"kind": "trace", "traceRule": {"latencyMs": 500}}},
			{"id": "3", "props": {"kind": "metric", "metricRule": {"threshold": 0.9}}}
		]
	}`)

	var resp ListAlertsResponse
	if err := json.Unmarshal(payload, &resp); err != nil {
		t.Fatalf("expected unmarshal to succeed despite unknown variant, got: %v", err)
	}

	if got, want := len(resp.Alerts), 2; got != want {
		t.Fatalf("len(Alerts) = %d, want %d (unknown variant should be dropped)", got, want)
	}

	if resp.Alerts[0].Id != "1" {
		t.Errorf("Alerts[0].Id = %q, want %q", resp.Alerts[0].Id, "1")
	}
	if resp.Alerts[1].Id != "3" {
		t.Errorf("Alerts[1].Id = %q, want %q", resp.Alerts[1].Id, "3")
	}

	if resp.Alerts[0].Props.LogsAlertProps == nil {
		t.Error("Alerts[0].Props.LogsAlertProps was nil; want LogsAlertProps to be set")
	}
	if resp.Alerts[1].Props.MetricAlertProps == nil {
		t.Error("Alerts[1].Props.MetricAlertProps was nil; want MetricAlertProps to be set")
	}
}

// Even when every list element is unknown the response itself decodes
// successfully — callers see an empty list rather than a hard error.
func TestForwardCompatOneOfList_AllUnknownYieldsEmptyList(t *testing.T) {
	payload := []byte(`{
		"alerts": [
			{"id": "1", "props": {"kind": "trace", "traceRule": {}}},
			{"id": "2", "props": {"kind": "audit", "auditRule": {}}}
		]
	}`)

	var resp ListAlertsResponse
	if err := json.Unmarshal(payload, &resp); err != nil {
		t.Fatalf("expected unmarshal to succeed when all variants unknown, got: %v", err)
	}
	if len(resp.Alerts) != 0 {
		t.Fatalf("len(Alerts) = %d, want 0 — all elements should be dropped as unknown", len(resp.Alerts))
	}
}

// A purely-known list keeps every element.
func TestForwardCompatOneOfList_AllKnownIsUnchanged(t *testing.T) {
	payload := []byte(`{
		"alerts": [
			{"id": "1", "props": {"kind": "logs", "logsRule": {"q": "error"}}},
			{"id": "2", "props": {"kind": "metric", "metricRule": {"threshold": 0.5}}}
		]
	}`)

	var resp ListAlertsResponse
	if err := json.Unmarshal(payload, &resp); err != nil {
		t.Fatalf("unmarshal failed on all-known payload: %v", err)
	}
	if len(resp.Alerts) != 2 {
		t.Fatalf("len(Alerts) = %d, want 2", len(resp.Alerts))
	}
}
