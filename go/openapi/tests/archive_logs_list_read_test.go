package tests

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
)

func TestArchiveLogsListRead(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewArchiveLogsClient(cfg)

	listResp, httpResp, err := client.
		S3TargetServiceGetTarget(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	listTarget := listResp.Target.GetActualInstance()
	if listTarget == nil {
		t.Skip("no resources to read")
	}

	getResp, httpResp, err := client.
		S3TargetServiceGetTarget(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	getTarget := getResp.Target.GetActualInstance()
	require.NotNil(t, getTarget)
	require.Equal(t, reflect.TypeOf(listTarget), reflect.TypeOf(getTarget))
}
