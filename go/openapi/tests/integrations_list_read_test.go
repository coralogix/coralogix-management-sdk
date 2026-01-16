package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	integrations "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/integration_service"
)

func integrationDetailsID(details integrations.IntegrationDetails) string {
	actual := details.GetActualInstance()
	switch typed := actual.(type) {
	case *integrations.IntegrationDetailsExternal:
		return typed.GetIntegration().GetId()
	case *integrations.IntegrationDetailsLocal:
		return typed.GetIntegration().GetId()
	default:
		return ""
	}
}

func TestIntegrationsListRead(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewIntegrationsClient(cfg)

	listResp, httpResp, err := client.
		IntegrationServiceGetIntegrations(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	if len(listResp.Integrations) == 0 {
		t.Skip("no resources to read")
	}

	integrationID := listResp.Integrations[0].GetIntegration().GetId()
	if integrationID == "" {
		t.Skip("no resources to read")
	}

	getResp, httpResp, err := client.
		IntegrationServiceGetIntegrationDetails(context.Background(), integrationID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, integrationID, integrationDetailsID(getResp.GetIntegrationDetail()))
}
