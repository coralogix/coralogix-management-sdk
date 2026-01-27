package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
)

func TestScopesListRead(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewScopesClient(cfg)

	listResp, httpResp, err := client.
		ScopesServiceGetTeamScopes(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	if len(listResp.Scopes) == 0 {
		t.Skip("no resources to read")
	}

	scopeID := listResp.Scopes[0].GetId()
	if scopeID == "" {
		t.Skip("no resources to read")
	}

	getResp, httpResp, err := client.
		ScopesServiceGetTeamScopesByIds(context.Background()).
		Ids([]string{scopeID}).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	if len(getResp.Scopes) == 0 {
		t.Skip("no resources to read")
	}
	require.Equal(t, scopeID, getResp.Scopes[0].GetId())
}
