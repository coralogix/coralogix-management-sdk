package required_object

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestRequiredObject_AllowsAdditiveFieldsButRejectsMissingRequired(t *testing.T) {
	var value RequiredObject

	if err := json.Unmarshal([]byte(`{"name":"policy","extra":"ignored"}`), &value); err != nil {
		t.Fatalf("json.Unmarshal returned error for additive field: %v", err)
	}
	if value.Name != "policy" {
		t.Fatalf("value.Name = %q, want %q", value.Name, "policy")
	}

	err := json.Unmarshal([]byte(`{"extra":"ignored"}`), &value)
	if err == nil {
		t.Fatal("expected missing required property to fail")
	}
	if !strings.Contains(err.Error(), "no value given for required property name") {
		t.Fatalf("missing required property error = %q", err)
	}
}
