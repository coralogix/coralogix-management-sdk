package examples

import (
	"context"
	cxsdk "coralogix-management-sdk/go"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlertScheduler(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	apiKey, err := cxsdk.CoralogixAPIKeyFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, apiKey)
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
	createAlertSchedulerResponse, e := a.CreateAlertScheduler(context.Background(), &cxsdk.CreateAlertSchedulerRuleRequest{
		AlertSchedulerRule: &cxsdk.AlertSchedulerRule{
			Name:        "example",
			Description: &description,
			MetaLabels:  metaLabels,
			Schedule:    &schedule,
			Filter: &cxsdk.AlertSchedulerFilter{
				WhatExpression: "source logs | filter $d.cpodId:string == '122'",
				WhichAlerts: &cxsdk.AlertSchedulerFilterUniqueIDs{
					AlertUniqueIds: &cxsdk.AlertUniqueIDs{
						Value: []string{"07d5714a-c847-4afa-87db-5e5f6986688c"},
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

	updateAlertSchedulerResponse, e := a.UpdateAlertScheduler(context.Background(), &cxsdk.UpdateAlertSchedulerRuleRequest{
		AlertSchedulerRule: alertScheduler,
	})

	assert.Nil(t, e)

	assert.Equal(t, updateAlertSchedulerResponse.AlertSchedulerRule.Name, "MyAlertUpdated")

	getAlertSchedulerResponse, e := a.GetAlertScheduler(context.Background(), &cxsdk.GetAlertSchedulerRuleRequest{
		AlertSchedulerRuleId: *updateAlertSchedulerResponse.AlertSchedulerRule.UniqueIdentifier,
	})

	assert.Nil(t, e)

	assert.Equal(t, getAlertSchedulerResponse.AlertSchedulerRule.Name, "MyAlertUpdated")

	_, error := a.DeleteAlertScheduler(context.Background(), &cxsdk.DeleteAlertSchedulerRuleRequest{
		AlertSchedulerRuleId: *updateAlertSchedulerResponse.AlertSchedulerRule.UniqueIdentifier,
	})
	assert.Nil(t, error)

}
