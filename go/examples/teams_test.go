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

	cxsdk "github.com/coralogix/coralogix-management-sdk"

	"github.com/stretchr/testify/assert"
)

func TestTeams(t *testing.T) {

	t.Skip("Skipping integration test")

	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)

	c := cxsdk.NewTeamsClient(creator)

	team1, e := c.Create(context.Background(), &cxsdk.CreateTeamInOrgRequest{
		TeamName: "team_1",
		TeamAdminsEmail: []string{
			"admin@coralogix.com",
			"admin2@coralogix.com",
		},
	})
	assert.Nil(t, e)

	team2, e := c.Create(context.Background(), &cxsdk.CreateTeamInOrgRequest{
		TeamName: "team_2",
		TeamAdminsEmail: []string{
			"admin@coralogix.com",
			"admin2@coralogix.com",
		},
	})
	assert.Nil(t, e)

	team_1_updated_name := "team_1_updated"
	_, e = c.Update(context.Background(), &cxsdk.UpdateTeamRequest{
		TeamName: &team_1_updated_name,
	})

	assert.Nil(t, e)

	updated, e := c.Get(context.Background(), &cxsdk.GetTeamRequest{
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

	_, deletionError := c.Delete(context.Background(), &cxsdk.DeleteTeamRequest{
		TeamId: team1.TeamId,
	})

	assert.Nil(t, deletionError)
}
