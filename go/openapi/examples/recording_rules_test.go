// Copyright 2024 Coralogix Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied. See the License for the specific language
// governing permissions and limitations under the License.

package examples

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	recordingrules "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/recording_rules_service"
)

func TestRecordingRuleGroups(t *testing.T) {
	ctx := context.Background()
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().WithHTTPLogging().Build()
	client := cxsdk.NewRecordingRulesClient(cfg)

	setName := fmt.Sprintf("TestRecordingRuleGroup-%d", time.Now().UnixMilli())

	interval := int64(180)
	limit := "100"

	req := recordingrules.CreateRuleGroupSet{
		Name: &setName,
		Groups: []recordingrules.InRuleGroup{
			{
				Name:     recordingrules.PtrString("foo"),
				Interval: &interval,
				Limit:    &limit,
				Rules: []recordingrules.InRule{
					{
						Record: recordingrules.PtrString("ts3db_live_ingester_write_latency:3m"),
						Expr:   recordingrules.PtrString("sum(rate(ts3db_live_ingester_write_latency_seconds_count{CX_LEVEL=\"staging\",pod=~\"ts3db-live-ingester.*\"}[2m])) by (pod)"),
					},
					{
						Record: recordingrules.PtrString("job:http_requests_total:sum"),
						Expr:   recordingrules.PtrString("sum(rate(http_requests_total[5m])) by (job)"),
					},
				},
			},
			{
				Name:     recordingrules.PtrString("bar"),
				Interval: &interval,
				Limit:    &limit,
				Rules: []recordingrules.InRule{
					{
						Record: recordingrules.PtrString("ts3db_live_ingester_write_latency:3m"),
						Expr:   recordingrules.PtrString("sum(rate(ts3db_live_ingester_write_latency_seconds_count{CX_LEVEL=\"staging\",pod=~\"ts3db-live-ingester.*\"}[2m])) by (pod)"),
					},
					{
						Record: recordingrules.PtrString("job:http_requests_total:sum"),
						Expr:   recordingrules.PtrString("sum(rate(http_requests_total[5m])) by (job)"),
					},
				},
			},
		},
	}

	created, httpResp, err := client.
		RuleGroupSetsCreate(ctx).
		CreateRuleGroupSet(req).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotNil(t, created.Id)

	updatedName := setName + "-updated"
	updateReq := recordingrules.UpdateRuleGroupSet{
		Name: &updatedName,
		Groups: []recordingrules.InRuleGroup{
			{
				Name:     recordingrules.PtrString("baz"),
				Interval: &interval,
				Limit:    &limit,
				Rules: []recordingrules.InRule{
					{
						Record: recordingrules.PtrString("job:http_requests_total:sum"),
						Expr:   recordingrules.PtrString("sum(rate(http_requests_total[5m])) by (job)"),
					},
				},
			},
		},
	}

	_, httpResp, err = client.
		RuleGroupSetsUpdate(ctx, *created.Id).
		UpdateRuleGroupSet(updateReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	got, httpResp, err := client.
		RuleGroupSetsFetch(ctx, *created.Id).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Len(t, got.Groups, 1)
	require.Equal(t, updatedName, *got.Name)

	list, httpResp, err := client.
		RuleGroupSetsList(ctx).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	found := false
	for _, s := range list.Sets {
		if s.Id != nil && *s.Id == *created.Id {
			found = true
			break
		}
	}
	require.True(t, found)

	_, httpResp, err = client.
		RuleGroupSetsDelete(ctx, *created.Id).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
}
