package tests

import (
	"context"
	"testing"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	"github.com/coralogix/coralogix-management-sdk/go/openapi/gen/slos_service"
	"github.com/stretchr/testify/require"
)

func TestSLOsListRead(t *testing.T) {
	t.Skip("later")
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewSLOsClient(cfg)

	listResp, httpResp, err := client.
		SlosServiceListSlos(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	if len(listResp.Slos) == 0 {
		t.Skip("no resources to read")
	}

	sloID := getIdFromSlo(listResp.Slos[0])

	if sloID == "" {
		t.Skip("no resources to read")
	}

	getResp, httpResp, err := client.
		SlosServiceGetSlo(context.Background(), sloID).
		Execute()

	id := getIdFromSlo(getResp.Slo)

	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, sloID, id)
}

func getIdFromSlo(slo slos_service.Slo) string {
	id := ""
	if slo.SloRequestBasedMetricSli != nil {
		id = *slo.SloRequestBasedMetricSli.Id
	}

	if slo.SloWindowBasedMetricSli != nil {
		id = *slo.SloWindowBasedMetricSli.Id
	}
	return id
}
