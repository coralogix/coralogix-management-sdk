// Copyright 2024 Coralogix Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.apache.org/licenses/LICENSE-2.0
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
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestDataUsage(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator, err := cxsdk.NewSDKCallPropertiesCreator(region, authContext)
	assertNilAndPrintError(t, err)
	c := cxsdk.NewDataUsageClient(creator)
	defer creator.CloseConnection()

	dateRangeBegin, err := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
	assertNilAndPrintError(t, err)
	dateRangeEnd, err := time.Parse(time.RFC3339, "2021-01-02T00:00:00Z")
	assertNilAndPrintError(t, err)

	logsCount, err := c.GetLogsCount(context.Background(), &cxsdk.GetLogsCountRequest{
		DateRange: &cxsdk.DateRange{
			FromDate: &timestamppb.Timestamp{
				Seconds: dateRangeBegin.Unix(),
			},
			ToDate: &timestamppb.Timestamp{
				Seconds: dateRangeEnd.Unix(),
			},
		},
		Resolution: durationpb.New(time.Hour),
		Filters: &cxsdk.ScopesFilter{
			Application: []string{"myapp"},
			Subsystem:   []string{"myapp"},
		},
	})

	assertNilAndPrintError(t, err)
	assert.NotNil(t, logsCount)

	spansCount, err := c.GetSpansCount(context.Background(), &cxsdk.GetSpansCountRequest{
		DateRange: &cxsdk.DateRange{
			FromDate: &timestamppb.Timestamp{
				Seconds: dateRangeBegin.Unix(),
			},
			ToDate: &timestamppb.Timestamp{
				Seconds: dateRangeEnd.Unix(),
			},
		},
		Resolution: durationpb.New(time.Hour),
		Filters: &cxsdk.ScopesFilter{
			Application: []string{"myapp"},
			Subsystem:   []string{"myapp"},
		},
	})

	assertNilAndPrintError(t, err)
	assert.NotNil(t, spansCount)

	dataUsageStream, err := c.GetDataUsage(context.Background(), &cxsdk.GetDataUsageRequest{
		DateRange: &cxsdk.DateRange{
			FromDate: &timestamppb.Timestamp{
				Seconds: dateRangeBegin.Unix(),
			},
			ToDate: &timestamppb.Timestamp{
				Seconds: dateRangeEnd.Unix(),
			},
		},
		Resolution:       durationpb.New(time.Hour),
		Aggregate:        []cxsdk.AggregateBy{},
		DimensionFilters: []*cxsdk.Dimension{},
	})

	assertNilAndPrintError(t, err)
	assert.NotNil(t, dataUsageStream)

}
