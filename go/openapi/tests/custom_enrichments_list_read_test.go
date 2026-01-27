package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
)

func TestCustomEnrichmentsListRead(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewCustomEnrichmentsClient(cfg)

	listResp, httpResp, err := client.
		CustomEnrichmentServiceGetCustomEnrichments(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	if len(listResp.CustomEnrichments) == 0 {
		t.Skip("no resources to read")
	}

	enrichmentID := listResp.CustomEnrichments[0].GetId()
	if enrichmentID == 0 {
		t.Skip("no resources to read")
	}

	getResp, httpResp, err := client.
		CustomEnrichmentServiceGetCustomEnrichment(context.Background(), enrichmentID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, enrichmentID, *getResp.GetCustomEnrichment().Id)
}
