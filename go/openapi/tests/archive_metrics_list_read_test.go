package tests

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
)

func TestArchiveMetricsListRead(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewArchiveMetricsClient(cfg)

	listResp, httpResp, err := client.
		MetricsConfiguratorPublicServiceGetTenantConfig(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	if listResp == nil || listResp.TenantConfig == nil {
		t.Skip("no resources to read")
	}

	listConfig := listResp.TenantConfig.GetActualInstance()
	if listConfig == nil {
		t.Skip("no resources to read")
	}

	getResp, httpResp, err := client.
		MetricsConfiguratorPublicServiceGetTenantConfig(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	if getResp == nil || getResp.TenantConfig == nil {
		t.Skip("no resources to read")
	}

	getConfig := getResp.TenantConfig.GetActualInstance()
	require.NotNil(t, getConfig)
	require.Equal(t, reflect.TypeOf(listConfig), reflect.TypeOf(getConfig))
}
