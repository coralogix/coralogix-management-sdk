package simple_oneof

import (
	"encoding/json"
	"testing"
)

func TestSimpleOneOf_SelectsSingleBranchAfterLenientDecode(t *testing.T) {
	var value SimpleOneOf

	err := json.Unmarshal([]byte(`{"name":"policy","priority":"high","logRules":{},"extra":"ignored"}`), &value)
	if err != nil {
		t.Fatalf("json.Unmarshal returned error: %v", err)
	}
	if value.LogRulesVariant == nil {
		t.Fatal("expected LogRulesVariant branch to match")
	}
	if value.SpanRulesVariant != nil {
		t.Fatal("expected SpanRulesVariant branch to remain nil")
	}
}
