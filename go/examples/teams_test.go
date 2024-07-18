package examples

import (
	"context"
	cxsdk "coralogix-management-sdk/go"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeams(t *testing.T) {
	t.Skip("Skipping integration test")

	region, err := cxsdk.CoralogixRegionFromEnv()
	t.Log(region)
	assert.Nil(t, err)
	apiKey, err := cxsdk.CoralogixAPIKeyFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, apiKey)
	c := cxsdk.NewTeamsClient(creator)

	team1, e := c.CreateTeam(context.Background(), &cxsdk.CreateTeamInOrgRequest{
		TeamName: "team_1",
		TeamAdminsEmail: []string{
			"admin@coralogix.com",
			"admin2@coralogix.com",
		},
	})
	assert.Nil(t, e)

	team2, e := c.CreateTeam(context.Background(), &cxsdk.CreateTeamInOrgRequest{
		TeamName: "team_2",
		TeamAdminsEmail: []string{
			"admin@coralogix.com",
			"admin2@coralogix.com",
		},
	})
	assert.Nil(t, e)

	team_1_updated_name := "team_1_updated"
	_, e = c.UpdateTeam(context.Background(), &cxsdk.UpdateTeamRequest{
		TeamName: &team_1_updated_name,
	})

	assert.Nil(t, e)

	updated, e := c.GetTeam(context.Background(), &cxsdk.GetTeamRequest{
		TeamId: team1.TeamId,
	})

	assert.Nil(t, e)

	assert.Equal(t, updated.TeamName, "team_1_updated")

	_, setQuotaError := c.SetDailyQuota(context.Background(), &cxsdk.SetDailyQuotaRequest{
		TeamId:           team1.TeamId,
		TargetDailyQuota: 250,
	})

	assert.Nil(t, setQuotaError)

	_, moveQuotaError := c.MoveQuota(context.Background(), &cxsdk.MoveQuotaRequest{
		SourceTeam:      team1.TeamId,
		DestinationTeam: team2.TeamId,
		UnitsToMove:     50,
	})

	assert.Nil(t, moveQuotaError)

	_, deletionError := c.DeleteTeam(context.Background(), &cxsdk.DeleteTeamRequest{
		TeamId: team1.TeamId,
	})

	assert.Nil(t, deletionError)
}
