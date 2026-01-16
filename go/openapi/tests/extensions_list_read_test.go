package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	extensions "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/extension_service"
)

func TestExtensionsListRead(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewExtensionsClient(cfg)

	includeHidden := false
	listReq := extensions.GetAllExtensionsRequest{
		IncludeHiddenExtensions: &includeHidden,
	}

	listResp, httpResp, err := client.
		ExtensionServiceGetAllExtensions(context.Background()).
		GetAllExtensionsRequest(listReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	if len(listResp.Extensions) == 0 {
		t.Skip("no resources to read")
	}

	extensionID := listResp.Extensions[0].GetId()
	if extensionID == "" {
		t.Skip("no resources to read")
	}

	getResp, httpResp, err := client.
		ExtensionServiceGetExtension(context.Background(), extensionID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, extensionID, getResp.GetId())
}
