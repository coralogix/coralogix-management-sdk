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

func TestAlertScheduler(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	a := cxsdk.NewAlertsSchedulersClient(creator)
	description := "example"
	metaLabels := make([]*cxsdk.MetaLabel, 0)
	schedule := cxsdk.Schedule{
		ScheduleOperation: cxsdk.ScheduleOperationScheduleOperationMute,
		Scheduler: &cxsdk.ScheduleOneTime{
			OneTime: &cxsdk.OneTime{
				Timeframe: &cxsdk.Timeframe{
					StartTime: "2021-01-04T00:00:00.000",
					Until: &cxsdk.TimeframeEndTime{
						EndTime: "2025-01-01T00:00:50.000",
					},
					Timezone: "UTC+2",
				},
			},
		},
	}
	createAlertSchedulerResponse, e := a.Create(context.Background(), &cxsdk.CreateAlertSchedulerRuleRequest{
		AlertSchedulerRule: &cxsdk.AlertSchedulerRule{
			Name:        "example",
			Description: &description,
			MetaLabels:  metaLabels,
			Schedule:    &schedule,
			Filter: &cxsdk.AlertSchedulerFilter{
				WhatExpression: "source logs | filter $d.cpodId:string == '122'",
				WhichAlerts: &cxsdk.AlertSchedulerFilterUniqueIDs{
					AlertUniqueIds: &cxsdk.AlertUniqueIDs{
						Value: []string{"55a457ed-5f23-407a-a724-12d7fe533a4e"},
					},
				},
			},
			Enabled:   true,
			CreatedAt: nil,
			UpdatedAt: nil,
		},
	})

	assert.Nil(t, e)

	alertScheduler := createAlertSchedulerResponse.AlertSchedulerRule
	alertScheduler.Name = "MyAlertUpdated"

	updateAlertSchedulerResponse, e := a.Update(context.Background(), &cxsdk.UpdateAlertSchedulerRuleRequest{
		AlertSchedulerRule: alertScheduler,
	})

	assert.Nil(t, e)

	assert.Equal(t, updateAlertSchedulerResponse.AlertSchedulerRule.Name, "MyAlertUpdated")

	getAlertSchedulerResponse, e := a.Get(context.Background(), &cxsdk.GetAlertSchedulerRuleRequest{
		AlertSchedulerRuleId: *updateAlertSchedulerResponse.AlertSchedulerRule.UniqueIdentifier,
	})

	assert.Nil(t, e)

	assert.Equal(t, getAlertSchedulerResponse.AlertSchedulerRule.Name, "MyAlertUpdated")

	_, error := a.Delete(context.Background(), &cxsdk.DeleteAlertSchedulerRuleRequest{
		AlertSchedulerRuleId: *updateAlertSchedulerResponse.AlertSchedulerRule.UniqueIdentifier,
	})
	assert.Nil(t, error)

}
