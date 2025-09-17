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

//
//import (
//	"context"
//	"fmt"
//	"testing"
//	"time"
//
//	"github.com/stretchr/testify/assert"
//
//		"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
//		recordingrules "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/metrics_rule_manager_service"
//)
//
//func TestRecordingRuleGroups(t *testing.T) {
//	t.Skip("Skipping because OpenAPI Facade is missing recording rule groups support")
//
//	client := cxsdk.NewClientBuilder().WithRegionFromEnv().WithAPIKeyFromEnv().Build()
//
//	setName := fmt.Sprintf("TestRecordingRuleGroupSet-%d", time.Now().UnixMilli())
//
//	interval := int64(180)
//	limit := "100"
//
//	req := cxsdk.RuleGroupSetsCreateRequest{
//		Name: &setName,
//		Groups: []cxsdk.ComCoralogixapisMetricsRuleManagerV1InRuleGroup{
//			{
//				Name:     cxsdk.PtrString("foo"),
//				Interval: &interval,
//				Limit:    &limit,
//				Rules: []cxsdk.ComCoralogixapisMetricsRuleManagerV1InRule{
//					{
//						Record: cxsdk.PtrString("ts3db_live_ingester_write_latency:3m"),
//						Expr:   cxsdk.PtrString("sum(rate(ts3db_live_ingester_write_latency_seconds_count{CX_LEVEL=\"staging\",pod=~\"ts3db-live-ingester.*\"}[2m])) by (pod)"),
//					},
//					{
//						Record: cxsdk.PtrString("job:http_requests_total:sum"),
//						Expr:   cxsdk.PtrString("sum(rate(http_requests_total[5m])) by (job)"),
//					},
//				},
//			},
//			{
//				Name:     cxsdk.PtrString("bar"),
//				Interval: &interval,
//				Limit:    &limit,
//				Rules: []cxsdk.ComCoralogixapisMetricsRuleManagerV1InRule{
//					{
//						Record: cxsdk.PtrString("ts3db_live_ingester_write_latency:3m"),
//						Expr:   cxsdk.PtrString("sum(rate(ts3db_live_ingester_write_latency_seconds_count{CX_LEVEL=\"staging\",pod=~\"ts3db-live-ingester.*\"}[2m])) by (pod)"),
//					},
//					{
//						Record: cxsdk.PtrString("job:http_requests_total:sum"),
//						Expr:   cxsdk.PtrString("sum(rate(http_requests_total[5m])) by (job)"),
//					},
//				},
//			},
//		},
//	}
//
//	created, httpResp, err := client.DefaultAPI.
//		RuleGroupSetsCreate(context.Background()).
//		RuleGroupSetsCreateRequest(req).
//		Execute()
//	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
//	assert.NotNil(t, created.Id)
//
//	got, httpResp, err := client.DefaultAPI.
//		RuleGroupSetsFetch(context.Background(), *created.Id).
//		Execute()
//	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
//	assert.Len(t, got.Groups, 2)
//
//	list, httpResp, err := client.DefaultAPI.
//		RuleGroupSetsList(context.Background()).
//		Execute()
//	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
//
//	found := false
//	for _, s := range list.Sets {
//		if s.Id != nil && *s.Id == *created.Id {
//			found = true
//			break
//		}
//	}
//	assert.True(t, found)
//
//	_, httpResp, err = client.DefaultAPI.
//		RuleGroupSetsDelete(context.Background(), *created.Id).
//		Execute()
//	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
//}
