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
	creator := cxsdk.NewSDKCallPropertiesCreator(region, authContext)
	c := cxsdk.NewSLOsClient(creator)
	sloDescription := "description"
	// sloId := "coralogix_slo_go_example"

	createSloResponse, err := c.Create(context.Background(), &cxsdk.CreateServiceSloRequest{
		Slo: &cxsdk.Slo{
			Name:                      "coralogix_slo_go_example",
			Description:               &sloDescription,
			TargetThresholdPercentage: 30,
			Sli: &cxsdk.SloRequestBasedMetricSli{
				RequestBasedMetricSli: &cxsdk.RequestBasedMetricSli{
					GoodEvents: &cxsdk.Metric{
						Query: "avg(rate(cpu_usage_seconds_total[1m])) by (instance)",
					},
					TotalEvents: &cxsdk.Metric{
						Query: "avg(rate(cpu_usage_seconds_total[1m])) by (instance)",
					},
				},
			},
			Window: &cxsdk.SloTimeframe{
				SloTimeFrame: cxsdk.SloTimeframe7Days,
			},
			Labels: map[string]string{
				"label1": "value1",
			},
		},
	})

	assertNilAndPrintError(t, err)

	_, retrievalError := c.Get(context.Background(), &cxsdk.GetServiceSloRequest{
		Id: *createSloResponse.Slo.Id,
	})

	assertNilAndPrintError(t, retrievalError)

	updateSloResponse, updateError := c.Update(context.Background(), &cxsdk.ReplaceServiceSloRequest{
		Slo: &cxsdk.Slo{
			Id:                        createSloResponse.Slo.Id,
			Revision:                  createSloResponse.Slo.Revision,
			Name:                      "coralogix_slo_go_example_updated",
			Description:               &sloDescription,
			TargetThresholdPercentage: 30,
			Sli: &cxsdk.SloRequestBasedMetricSli{
				RequestBasedMetricSli: &cxsdk.RequestBasedMetricSli{
					GoodEvents: &cxsdk.Metric{
						Query: "avg(rate(cpu_usage_seconds_total[1m])) by (instance)",
					},
					TotalEvents: &cxsdk.Metric{
						Query: "avg(rate(cpu_usage_seconds_total[1m])) by (instance)",
					},
				},
			},
			Window: &cxsdk.SloTimeframe{
				SloTimeFrame: cxsdk.SloTimeframe7Days,
			},
			Labels: map[string]string{
				"label1": "value1",
			},
		},
	})

	assertNilAndPrintError(t, updateError)

	listResponse, err := c.List(context.Background(), &cxsdk.ListSlosRequest{
		Filters: &cxsdk.SloFilters{
			Filters: []*cxsdk.SloFilter{
				{
					Field: &cxsdk.SloFilterField{
						Field: &cxsdk.SloLabelNameFilterField{
							LabelName: "label1",
						},
					},
					Predicate: &cxsdk.SloFilterPredicate{
						Predicate: &cxsdk.SloFilterPredicateIs{
							Is: &cxsdk.IsSloFilterPredicate{
								Is: []string{"value1"},
							},
						},
					},
				},
			},
		},
	})

	slos := listResponse.Slos
	assert.Greater(t, len(slos), 0)

	assertNilAndPrintError(t, err)

	assert.Equal(t, updateSloResponse.Slo.Id, updateSloResponse.Slo.Id)
	assert.Equal(t, updateSloResponse.Slo.Revision.Revision, updateSloResponse.Slo.Revision.Revision)
	assert.Equal(t, updateSloResponse.Slo.Name, updateSloResponse.Slo.Name)

	_, deletionError := c.Delete(context.Background(), &cxsdk.DeleteServiceSloRequest{
		Id: *createSloResponse.Slo.Id,
	})

	assertNilAndPrintError(t, deletionError)
}
