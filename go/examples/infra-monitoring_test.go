// Copyright 2024 Coralogix Ltd.
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

	cxsdk "github.com/coralogix/coralogix-management-sdk/go"

	"github.com/stretchr/testify/assert"
)

func TestSlos(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	c := cxsdk.NewSLOsClient(creator)
	sloDescription := "description"

	createSloResponse, err := c.Create(context.Background(), &cxsdk.CreateServiceSloRequest{
		Slo: &cxsdk.Slo{
			Name:                      "coralogix_slo_go_example",
			Description:               &sloDescription,
			TargetThresholdPercentage: 30,
			Sli: &cxsdk.SloMetricSli{
				MetricSli: &cxsdk.MetricSli{
					GoodEvents: &cxsdk.Metric{
						Query: "avg(rate(cpu_usage_seconds_total[5m])) by instance",
					},
					TotalEvents: &cxsdk.Metric{
						Query: "avg(rate(cpu_usage_seconds_total[5m])) by instance",
					},
				},
			},
			Window: &cxsdk.SloTimeframe{
				SloTimeFrame: cxsdk.SloTimeframe7Days,
			},
		},
	})

	assertNilAndPrintError(t, err)

	_, retrievalError := c.Get(context.Background(), &cxsdk.GetServiceSloRequest{
		Id: createSloResponse.Slo.Id,
	})

	assertNilAndPrintError(t, retrievalError)

	updateSloResponse, updateError := c.Update(context.Background(), &cxsdk.ReplaceServiceSloRequest{
		Slo: &cxsdk.Slo{
			Name:                      "coralogix_slo_go_example_updated",
			Description:               &sloDescription,
			TargetThresholdPercentage: 95,
			Sli:                       &cxsdk.SloMetricSli{},
			Window: &cxsdk.SloTimeframe{
				SloTimeFrame: cxsdk.SloTimeframe7Days,
			},
		},
	})

	assertNilAndPrintError(t, updateError)

	assert.Equal(t, createSloResponse.Slo.Id, updateSloResponse.Slo.Id)

	_, deletionError := c.Delete(context.Background(), &cxsdk.DeleteServiceSloRequest{
		Id: createSloResponse.Slo.Id,
	})

	assertNilAndPrintError(t, deletionError)
}

// func TestSlosWithFilters(t *testing.T) {

// 	region, err := cxsdk.CoralogixRegionFromEnv()
// 	assertNilAndPrintError(t, err)
// 	authContext, err := cxsdk.AuthContextFromEnv()
// 	assertNilAndPrintError(t, err)
// 	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
// 	c := cxsdk.NewSLOsClient(creator)
// 	sloDescription := "description"
// 	timeFrame := cxsdk.SloTimeFrame{
// 		SloTimeFrame: cxsdk.SloTimeframe7Days,
// 	}

// 	createSloResponse, err := c.Create(context.Background(), &cxsdk.CreateServiceSloRequest{
// 		Slo: &cxsdk.Slo{
// 			Name:                      "coralogix_slo_example",
// 			Description:               &sloDescription,
// 			TargetThresholdPercentage: 30,
// 			Sli:                       &cxsdk.SloMetricSli{},
// 			Window:                    &timeFrame,
// 		},
// 	})

// 	assertNilAndPrintError(t, err)

// 	_, retrievalError := c.Get(context.Background(), &cxsdk.GetServiceSloRequest{
// 		Id: createSloResponse.Slo.Id,
// 	})

// 	assertNilAndPrintError(t, retrievalError)

// 	updateSloResponse, updateError := c.Update(context.Background(), &cxsdk.ReplaceServiceSloRequest{
// 		Slo: &cxsdk.ServiceSlo{
// 			Id:               createSloResponse.Slo.Id,
// 			Name:             &wrapperspb.StringValue{Value: "coralogix_slo_updated_example"},
// 			ServiceName:      &wrapperspb.StringValue{Value: "service_name"},
// 			Description:      &wrapperspb.StringValue{Value: "description"},
// 			TargetPercentage: &wrapperspb.UInt32Value{Value: 30},
// 			SliType:          &cxsdk.ServiceSloErrorSli{},
// 			Period:           v1.SloPeriod_SLO_PERIOD_7_DAYS,
// 		},
// 	})

// 	assertNilAndPrintError(t, updateError)

// 	assert.Equal(t, createSloResponse.Slo.Id.Value, updateSloResponse.Slo.Id.Value)

// 	_, deletionError := c.Delete(context.Background(), &cxsdk.DeleteServiceSloRequest{
// 		Id: createSloResponse.Slo.Id,
// 	})

// 	assertNilAndPrintError(t, deletionError)
// }
