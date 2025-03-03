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
	"strconv"
	"testing"
	"time"

	cxsdk "github.com/coralogix/coralogix-management-sdk/go"
	"github.com/stretchr/testify/assert"
)

func TestRecordingRuleGroups(t *testing.T) {

	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	c := cxsdk.NewRecordingRuleGroupSetsClient(creator)
	setName := "TestRecordingRuleGroupSet " + strconv.FormatInt(time.Now().UnixMilli(), 10)
	interval := uint32(180)
	limit := uint64(100)

	createRuleGroupSet, err := c.Create(context.Background(), &cxsdk.CreateRuleGroupSetRequest{
		Name: &setName,
		Groups: []*cxsdk.InRuleGroup{
			{
				Name:     "foo",
				Interval: &interval,
				Limit:    &limit,
				Rules: []*cxsdk.InRule{
					{
						Record: "ts3db_live_ingester_write_latency:3m",
						Expr:   "sum(rate(ts3db_live_ingester_write_latency_seconds_count{CX_LEVEL=\"staging\",pod=~\"ts3db-live-ingester.*\"}[2m])) by (pod)",
					},
					{
						Record: "job:http_requests_total:sum",
						Expr:   "sum(rate(http_requests_total[5m])) by (job)",
					},
				},
			},
			{
				Name:     "bar",
				Interval: &interval,
				Limit:    &limit,
				Rules: []*cxsdk.InRule{
					{
						Record: "ts3db_live_ingester_write_latency:3m",
						Expr:   "sum(rate(ts3db_live_ingester_write_latency_seconds_count{CX_LEVEL=\"staging\",pod=~\"ts3db-live-ingester.*\"}[2m])) by (pod)",
					},
					{
						Record: "job:http_requests_total:sum",
						Expr:   "sum(rate(http_requests_total[5m])) by (job)",
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	recordingRuleGroupSet, err := c.Get(context.Background(), &cxsdk.GetRuleGroupSetRequest{
		Id: createRuleGroupSet.Id,
	})

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(&testing.T{}, (recordingRuleGroupSet.Groups), 2)

	_, err = c.Delete(context.Background(), &cxsdk.DeleteRuleGroupSetRequest{Id: createRuleGroupSet.Id})

	if err != nil {
		t.Fatal(err)
	}
}
