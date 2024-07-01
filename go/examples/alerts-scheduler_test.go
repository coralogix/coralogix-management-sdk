package examples

import (
	"context"
	cxsdk "coralogix-management-sdk/go"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlertScheduler(t *testing.T) {
	creator := cxsdk.NewCallPropertiesCreator("https://ng-api-grpc.coralogix.com", "my-secret-token")
	a := cxsdk.NewAlertsSchedulersClient(creator)
	description := "example"
	metaLabels := make([]*cxsdk.MetaLabel, 0)
	schedule := cxsdk.Schedule{
		ScheduleOperation: cxsdk.ScheduleOperation_SCHEDULE_OPERATION_MUTE,
		Scheduler: &cxsdk.Schedule_OneTime{
			OneTime: &cxsdk.OneTime{
				Timeframe: &cxsdk.Timeframe{
					StartTime: "2021-01-04T00:00:00.000",
					Until: &cxsdk.Timeframe_EndTime{
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
		},
	})

	assert.Nil(t, e)

	alertScheduler := createAlertSchedulerResponse.AlertSchedulerRule
	alertScheduler.Name = "example_updated"

	updateAlertSchedulerResponse, e := a.UpdateAlertScheduler(context.Background(), &cxsdk.UpdateAlertSchedulerRuleRequest{
		AlertSchedulerRule: alertScheduler,
	})

	assert.Nil(t, e)

	assert.Equal(t, updateAlertSchedulerResponse.AlertSchedulerRule.Name, "NewName")

	getAlertSchedulerResponse, e := a.GetAlertScheduler(context.Background(), &cxsdk.GetAlertSchedulerRuleRequest{
		AlertSchedulerRuleId: *updateAlertSchedulerResponse.AlertSchedulerRule.Id,
	})

	assert.Nil(t, e)

	assert.Equal(t, getAlertSchedulerResponse.AlertSchedulerRule.Name, "NewName")

	_, err := a.DeleteAlertScheduler(context.Background(), &cxsdk.DeleteAlertSchedulerRuleRequest{})
	assert.Nil(t, err)

}
