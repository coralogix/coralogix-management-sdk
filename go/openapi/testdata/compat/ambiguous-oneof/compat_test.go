package ambiguous_oneof

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestAmbiguousOneOf_RemainsAmbiguous(t *testing.T) {
	var value AmbiguousOneOf

	err := json.Unmarshal([]byte(`{"value":"shared","extra":"ignored"}`), &value)
	if err == nil {
		t.Fatal("expected ambiguous payload to fail")
	}
	if !strings.Contains(err.Error(), "matches more than one schema") {
		t.Fatalf("ambiguity error = %q", err)
	}
}
