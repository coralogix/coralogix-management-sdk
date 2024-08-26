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
	"os"
	"strconv"
	"testing"

	cxsdk "github.com/coralogix/coralogix-management-sdk"
	v1 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/permissions/v1"

	"github.com/stretchr/testify/assert"
)

func TestGroups(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	teamId, err := strconv.ParseUint(os.Getenv("TEAM_ID"), 10, 32)
	assert.Nil(t, err)
	c := cxsdk.NewGroupsClient(creator)

	groups, err := c.List(context.Background(), &v1.GetTeamGroupsRequest{
		TeamId: &v1.TeamId{
			Id: uint32(teamId),
		},
	})
	assert.Nil(t, err)
	assert.Greater(t, len(groups.Groups), 0)

	groupDesc := "A Test Group"

	createdGroup, err := c.Create(context.Background(), &v1.CreateTeamGroupRequest{
		Name: "Test Group",
		TeamId: &v1.TeamId{
			Id: uint32(teamId),
		},
		Description: &groupDesc,
		ExternalId:  nil,
		RoleIds: []*v1.RoleId{
			{Id: 1},
		},
		UserIds: []*v1.UserId{},
	})

	assert.Nil(t, err)

	retrievedGroup, err := c.Get(context.Background(), &v1.GetTeamGroupRequest{
		GroupId: &v1.TeamGroupId{
			Id: createdGroup.GroupId.Id,
		},
	})

	assert.Nil(t, err)
	assert.Equal(t, createdGroup.GroupId.Id, retrievedGroup.Group.GroupId.Id)

	_, updateError := c.Update(context.Background(), &v1.UpdateTeamGroupRequest{
		GroupId:     createdGroup.GroupId,
		Name:        "Updated Test Group",
		Description: &groupDesc,
	})

	assert.Nil(t, updateError)

	_, deletionError := c.Delete(context.Background(), &v1.DeleteTeamGroupRequest{
		GroupId: &v1.TeamGroupId{
			Id: createdGroup.GroupId.Id,
		},
	})

	assert.Nil(t, deletionError)
}
