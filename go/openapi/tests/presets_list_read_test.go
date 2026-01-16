package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
)

func TestPresetsListRead(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewPresetsClient(cfg)

	listResp, httpResp, err := client.
		PresetsServiceListPresetSummaries(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	if len(listResp.PresetSummaries) == 0 {
		t.Skip("no resources to read")
	}

	presetID := listResp.PresetSummaries[0].GetId()
	if presetID == "" {
		t.Skip("no resources to read")
	}

	getResp, httpResp, err := client.
		PresetsServiceGetPreset(context.Background(), presetID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, presetID, getResp.GetPreset().GetId())
}
