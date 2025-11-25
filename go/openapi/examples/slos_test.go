// Copyright 2025 Coralogix Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package examples

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	slos "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/slos_service"
)

func TestSLOs(t *testing.T) {
	ctx := context.Background()
	region, _ := cxsdk.URLFromRegion(cxsdk.RegionFromEnv())
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		region,
		cxsdk.APIKeyFromEnv(),
		true,
	)

	client := cxsdk.NewClientSet(cpc).SLOs()

	sloName := "example_slo_" + uuid.NewString()
	sloPayload := getRequestBasedSlo(sloName)
	createReq := slos.SlosServiceCreateSloRequest{
		SloRequestBasedMetricSli: sloPayload,
	}
	createResp, httpResp, err := client.SlosServiceCreateSlo(ctx).SlosServiceCreateSloRequest(createReq).Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	sloID := createResp.GetSlo().SloRequestBasedMetricSli.GetId()
	getResp, httpResp, err := client.
		SlosServiceGetSlo(ctx, sloID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, sloName, getResp.GetSlo().SloRequestBasedMetricSli.GetName())

	updatedName := "updated_example_slo_" + uuid.NewString()
	updatePayload := sloPayload
	updatePayload.Id = &sloID
	updatePayload.Name = updatedName

	updateReq := slos.SlosServiceReplaceSloRequest{
		SloRequestBasedMetricSli: updatePayload,
	}

	updateResp, httpResp, err := client.
		SlosServiceReplaceSlo(ctx).
		SlosServiceReplaceSloRequest(updateReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, updatedName, updateResp.GetSlo().SloRequestBasedMetricSli.GetName())

	listResp, httpResp, err := client.
		SlosServiceListSlos(ctx).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	found := false
	for _, s := range listResp.GetSlos() {
		if s.SloRequestBasedMetricSli.GetId() == sloID {
			found = true
			break
		}
	}
	require.True(t, found, "created SLO not found in list response")

	_, httpResp, err = client.
		SlosServiceDeleteSlo(ctx, sloID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
}

func getRequestBasedSlo(name string) *slos.SloRequestBasedMetricSli {
	desc := "example SLO created via OpenAPI SDK"
	target := float32(30.0)

	return &slos.SloRequestBasedMetricSli{
		Name:                      name,
		Description:               slos.PtrString(desc),
		TargetThresholdPercentage: target,
		RequestBasedMetricSli: &slos.RequestBasedMetricSli{
			GoodEvents: slos.Metric{
				Query: "avg(rate(cpu_usage_seconds_total[5m])) by (instance)",
			},
			TotalEvents: slos.Metric{
				Query: "avg(rate(cpu_usage_seconds_total[5m])) by (instance)",
			},
		},
		SloTimeFrame: slos.SLOTIMEFRAME_SLO_TIME_FRAME_7_DAYS.Ptr(),
		Labels:       &map[string]string{"label1": "value1"},
	}
}
