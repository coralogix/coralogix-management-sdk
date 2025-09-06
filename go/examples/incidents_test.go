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
	"time"

	cxsdk "github.com/coralogix/coralogix-management-sdk/go"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestIncidents(t *testing.T) {
	// Setup client
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator, err := cxsdk.NewSDKCallPropertiesCreator(region, authContext)
	assertNilAndPrintError(t, err)
	c := cxsdk.NewIncidentsClient(creator)
	defer creator.CloseConnection()

	// List incidents to get initial state
	listResponse, err := c.ListIncidents(context.Background(), &cxsdk.ListIncidentsRequest{})
	assertNilAndPrintError(t, err)
	initialCount := len(listResponse.Incidents)

	// Get incident events
	_, err = c.ListIncidentEvents(context.Background(), &cxsdk.ListIncidentEventsRequest{
		Filter: &cxsdk.IncidentEventQueryFilter{
			Name: &wrapperspb.StringValue{
				Value: "test",
			},
			IsMuted: &wrapperspb.BoolValue{
				Value: false,
			},
			Labels: &cxsdk.IncidentEventLabelsFilter{
				MetaLabels: []*cxsdk.IncidentEventMetaLabel{
					{
						Key: &wrapperspb.StringValue{
							Value: "test",
						},
						Value: &wrapperspb.StringValue{
							Value: "a test value",
						},
					},
				},
				Operator: cxsdk.IncidentEventFilterOperatorOrOrUnspecified,
			},

			Timestamp: &cxsdk.IncidentEventTimeRange{
				StartTime: timestamppb.New(time.Now().Add(-24 * time.Hour)),
				EndTime:   timestamppb.New(time.Now()),
			},
		},
		OrderBy: &cxsdk.ListIncidentEventOrderBy{
			Field:     cxsdk.IncidentEventOrderByFieldTypeTimestampOrUnspecified,
			Direction: cxsdk.IncidentEventOrderByDirectionDesc,
		},
		Pagination: &cxsdk.PaginationRequest{
			PageSize: &wrapperspb.UInt32Value{
				Value: 10,
			},
		},
	})
	assertNilAndPrintError(t, err)

	// Get events count
	_, err = c.ListIncidentEventsTotalCount(context.Background(), &cxsdk.ListIncidentEventsTotalCountRequest{
		Filter: &cxsdk.IncidentEventQueryFilter{
			Name: &wrapperspb.StringValue{
				Value: "test",
			},
			IsMuted: &wrapperspb.BoolValue{
				Value: false,
			},
			Labels: &cxsdk.IncidentEventLabelsFilter{
				MetaLabels: []*cxsdk.IncidentEventMetaLabel{
					{
						Key: &wrapperspb.StringValue{
							Value: "test",
						},
						Value: &wrapperspb.StringValue{
							Value: "a test value",
						},
					},
				},
				Operator: cxsdk.IncidentEventFilterOperatorOrOrUnspecified,
			},

			Timestamp: &cxsdk.IncidentEventTimeRange{
				StartTime: timestamppb.New(time.Now().Add(-24 * time.Hour)),
				EndTime:   timestamppb.New(time.Now()),
			},
		},
	})
	assertNilAndPrintError(t, err)

	// Get filter values
	_, err = c.ListIncidentEventsFilterValues(context.Background(), &cxsdk.ListIncidentEventsFilterValuesRequest{
		Filter: &cxsdk.IncidentEventQueryFilter{
			Name: &wrapperspb.StringValue{
				Value: "test",
			},
			IsMuted: &wrapperspb.BoolValue{
				Value: false,
			},
			Labels: &cxsdk.IncidentEventLabelsFilter{
				MetaLabels: []*cxsdk.IncidentEventMetaLabel{
					{
						Key: &wrapperspb.StringValue{
							Value: "test",
						},
						Value: &wrapperspb.StringValue{
							Value: "a test value",
						},
					},
				},
				Operator: cxsdk.IncidentEventFilterOperatorOrOrUnspecified,
			},

			Timestamp: &cxsdk.IncidentEventTimeRange{
				StartTime: timestamppb.New(time.Now().Add(-24 * time.Hour)),
				EndTime:   timestamppb.New(time.Now()),
			},
		},
	})
	assertNilAndPrintError(t, err)

	// If there are incidents, test operations on the first one
	if initialCount > 0 && len(listResponse.Incidents) > 0 {
		incident := listResponse.Incidents[0]

		getResponse, err := c.GetIncident(context.Background(), &cxsdk.GetIncidentRequest{
			Id: incident.Id,
		})
		assertNilAndPrintError(t, err)
		assert.Equal(t, incident.Id.Value, getResponse.Incident.Id.Value)

		batchResponse, err := c.BatchGetIncident(context.Background(), &cxsdk.BatchGetIncidentRequest{
			Ids: []*wrapperspb.StringValue{
				incident.Id,
			},
		})
		assertNilAndPrintError(t, err)
		assert.Equal(t, 1, len(batchResponse.Incidents))

		_, err = c.ResolveIncidents(context.Background(), &cxsdk.ResolveIncidentsRequest{
			IncidentIds: []*wrapperspb.StringValue{
				incident.Id,
			},
		})
		assertNilAndPrintError(t, err)
	}

	_, err = c.ListIncidentAggregations(context.Background(), &cxsdk.ListIncidentAggregationsRequest{})
	assertNilAndPrintError(t, err)
}
