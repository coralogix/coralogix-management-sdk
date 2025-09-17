// Copyright 2024 Coralogix Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package examples

//
//import (
//	"context"
//	"fmt"
//	"testing"
//
//	"github.com/google/uuid"
//	"github.com/stretchr/testify/assert"
//
//	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
//	slos "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/slo_service"
//)
//
//func TestSlos(t *testing.T) {
//	t.Skip("Unstable Test")
//	cpc := cxsdk.NewSDKCallPropertiesCreator(
//		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
//		cxsdk.APIKeyFromEnv(),
//	)
//
//	client := cxsdk.NewSLOsClient(cpc)
//	name := fmt.Sprintf("coralogix_slo_go_example_%v", uuid.NewString())
//
//	slo := slos.ServiceSlo{
//		Name:                      name,
//		Description:               slos.PtrString("Example SLO created via Go SDK"),
//		TargetThresholdPercentage: 30,
//		RequestBasedMetricSli: &slos.RequestBasedMetricSli{
//			GoodEvents: slos.Metric{
//				Query: "avg(rate(cpu_usage_seconds_total[5m])) by (instance)",
//			},
//			TotalEvents: slos.Metric{
//				Query: "avg(rate(cpu_usage_seconds_total[5m])) by (instance)",
//			},
//		},
//		SloTimeFrame: slos.PtrString("SLO_TIME_FRAME_7_DAYS"),
//		Labels: []slos.SloLabelsInner{
//			{Key: slos.PtrString("label1"), Value: slos.PtrString("value1")},
//		},
//	}
//
//	req := slos.CreateServiceSLORequest{
//		Slo: &slo,
//	}
//
//	createResp, httpResp, err := client.
//		ServiceSloServiceCreateServiceSlo(context.Background()).
//		CreateServiceSLORequest(req).
//		Execute()
//	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
//
//	sloID := createResp.Slo.Slo.Id
//	assert.NotNil(t, sloID)
//
//	got, httpResp, err := client.
//		SlosServiceGetSlo(context.Background(), *sloID).
//		Execute()
//	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
//	assert.Equal(t, name, got.Slo.Slo.Name)
//
//	updatedName := name + "_updated"
//	slo.Id = sloID
//	slo.Revision = createResp.Slo.Slo.Revision
//	slo.Name = updatedName
//
//	req = slos.SlosServiceReplaceSloRequest{
//		Slo2: &slo,
//	}
//	updateResp, httpResp, err := client.
//		SlosServiceReplaceSlo(context.Background()).
//		SlosServiceReplaceSloRequest(req).
//		Execute()
//	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
//	assert.Equal(t, updatedName, updateResp.Slo.Slo.Name)
//
//	listResp, httpResp, err := client.
//		SlosServiceListSlos(context.Background()).
//		Execute()
//	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
//	assert.Greater(t, len(listResp.Slos), 0)
//
//	_, httpResp, err = client.
//		SlosServiceDeleteSlo(context.Background(), *sloID).
//		Execute()
//	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
//}
