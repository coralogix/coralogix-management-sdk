package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
)

func TestEnrichmentsListRead(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewEnrichmentsClient(cfg)

	listResp, httpResp, err := client.
		EnrichmentServiceGetEnrichments(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	if len(listResp.Enrichments) == 0 {
		t.Skip("no resources to read")
	}

	enrichmentID := listResp.Enrichments[0].GetId()
	if enrichmentID == 0 {
		t.Skip("no resources to read")
	}

	getResp, httpResp, err := client.
		EnrichmentServiceGetEnrichments(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	found := false
	for _, enrichment := range getResp.Enrichments {
		if enrichment.GetId() == enrichmentID {
			found = true
			break
		}
	}
	require.True(t, found)
}
