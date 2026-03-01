package tests

import (
	"context"
	"testing"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	"github.com/stretchr/testify/require"
)

func TestWebhooksListRead(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewWebhooksClient(cfg)

	_, httpResp, err := client.
		OutgoingWebhooksServiceListAllOutgoingWebhooks(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	// skip the remainder because each type has their own id field and that would become a massive if chain...
	// if len(listResp.Deployed) == 0 {
	// 	t.Skip("no resources to read")
	// }

	// webhookID := listResp.Deployed[0].GetId()
	// if webhookID == "" {
	// 	t.Skip("no resources to read")
	// }

	// getResp, httpResp, err := client.
	// 	OutgoingWebhooksServiceGetOutgoingWebhook(context.Background(), webhookID).
	// 	Execute()
	// require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	// require.Equal(t, webhookID, getResp.GetWebhook().())
}
