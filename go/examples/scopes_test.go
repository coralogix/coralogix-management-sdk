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

func TestScopes(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	c := cxsdk.NewScopesClient(creator)
	all, e := c.List(context.Background(), &cxsdk.GetTeamScopesRequest{})
	assert.Nil(t, e)
	beginngingCount := len(all.Scopes)
	description := "Data Access Rule intended for testing"
	result, e := c.Create(context.Background(), &cxsdk.CreateScopeRequest{
		DisplayName: "Test Data Access Rule",
		Description: &description,
		Filters: []*cxsdk.Filter{
			{EntityType: cxsdk.EntityTypeLogs, Expression: "<v1> foo == 'bar'"},
		},
		DefaultExpression: "<v1>true",
	})
	assert.Nil(t, e)

	_, e = c.Update(context.Background(), &cxsdk.UpdateScopeRequest{
		Id:                result.Scope.Id,
		DisplayName:       "Updated Test Data Access Rule",
		Filters:           result.Scope.Filters,
		DefaultExpression: result.Scope.DefaultExpression,
	})
	assert.Nil(t, e)

	updated, _ := c.Get(context.Background(), &cxsdk.GetTeamScopesByIDsRequest{
		Ids: []string{result.Scope.Id},
	})

	assert.Equal(t, "Updated Test Data Access Rule", updated.Scopes[0].DisplayName)

	c.Delete(context.Background(), &cxsdk.DeleteScopeRequest{
		Id: result.Scope.Id,
	})

	all, e = c.List(context.Background(), &cxsdk.GetTeamScopesRequest{})
	assert.Nil(t, e)
	assert.Equal(t, beginngingCount, len(all.Scopes))
}
