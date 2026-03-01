package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
)

func TestE2MsListRead(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewEvents2MetricsClient(cfg)

	listResp, httpResp, err := client.
		Events2MetricServiceListE2M(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	if len(listResp.E2m) == 0 {
		t.Skip("no resources to read")
	}

	var e2mID string
	if listResp.E2m[0].E2MLogsQuery != nil {
		e2mID = listResp.E2m[0].E2MLogsQuery.GetId()
	} else if listResp.E2m[0].E2MSpansQuery != nil {
		e2mID = listResp.E2m[0].E2MSpansQuery.GetId()
	}
	if e2mID == "" {
		t.Skip("no resources to read")
	}

	getResp, httpResp, err := client.
		Events2MetricServiceGetE2M(context.Background(), e2mID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	var readID string
	if getResp.E2m.E2MLogsQuery != nil {
		readID = getResp.E2m.E2MLogsQuery.GetId()
	} else if getResp.E2m.E2MSpansQuery != nil {
		readID = getResp.E2m.E2MSpansQuery.GetId()
	}
	require.Equal(t, e2mID, readID)
}
