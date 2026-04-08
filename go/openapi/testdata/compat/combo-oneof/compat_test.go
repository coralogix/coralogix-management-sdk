package combo_oneof

import (
	"encoding/json"
	"testing"
)

func TestComboOneOf_SelectsSingleBranchFromRequiredFieldSet(t *testing.T) {
	var value ComboOneOf

	err := json.Unmarshal([]byte(`{"common":"shared","alpha":"a","xray":"x","extra":"ignored"}`), &value)
	if err != nil {
		t.Fatalf("json.Unmarshal returned error: %v", err)
	}
	if value.AlphaXrayVariant == nil {
		t.Fatal("expected AlphaXrayVariant branch to match")
	}
	if value.AlphaYankeeVariant != nil || value.BetaXrayVariant != nil || value.BetaYankeeVariant != nil {
		t.Fatal("expected all non-matching branches to remain nil")
	}
}
