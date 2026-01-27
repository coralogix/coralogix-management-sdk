package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
)

func TestArchiveRetentionsListRead(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewArchiveRetentionsClient(cfg)

	listResp, httpResp, err := client.
		RetentionsServiceGetRetentions(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	if len(listResp.Retentions) == 0 {
		t.Skip("no resources to read")
	}

	retentionID := listResp.Retentions[0].GetId()
	if retentionID == "" {
		t.Skip("no resources to read")
	}

	getResp, httpResp, err := client.
		RetentionsServiceGetRetentions(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	found := false
	for _, retention := range getResp.Retentions {
		if retention.GetId() == retentionID {
			found = true
			break
		}
	}
	require.True(t, found)
}
