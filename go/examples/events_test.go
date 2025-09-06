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

func TestEvents(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator, err := cxsdk.NewSDKCallPropertiesCreator(region, authContext)
	assertNilAndPrintError(t, err)
	c := cxsdk.NewEventsClient(creator)
	defer creator.CloseConnection()

	filter := &cxsdk.EventsFilter{
		Timestamp: &cxsdk.EventsQueryFilterTimestampRange{
			From: timestamppb.New(time.Now().Add(-24 * time.Hour)),
			To:   timestamppb.New(time.Now()),
		},
	}

	// Test ListCount with filters
	countResponse, err := c.ListCount(context.Background(), &cxsdk.ListEventsCountRequest{
		Filter: &cxsdk.EventsFilter{
			Timestamp: &cxsdk.EventsQueryFilterTimestampRange{
				From: timestamppb.New(time.Now().Add(-24 * time.Hour)),
				To:   timestamppb.New(time.Now()),
			},
		},
	})
	assertNilAndPrintError(t, err)
	assert.NotNil(t, countResponse)

	// Test List with filters
	_, err = c.List(context.Background(), &cxsdk.ListEventsRequest{
		Filter:   filter,
		OrderBys: []*cxsdk.OrderBy{},
		Pagination: &cxsdk.EventsPaginationRequest{
			PageSize: wrapperspb.UInt32(10),
		},
	})
	assertNilAndPrintError(t, err)

	_, err = c.ListCount(context.Background(), &cxsdk.ListEventsCountRequest{
		Filter: filter,
	})
	assertNilAndPrintError(t, err)

	_, err = c.GetStatistics(context.Background(), &cxsdk.GetEventsStatisticsRequest{
		Filter: filter,
	})
	assertNilAndPrintError(t, err)

}
