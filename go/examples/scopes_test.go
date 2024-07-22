package examples

import (
	"context"
	cxsdk "coralogix-management-sdk/go"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScopes(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	apiKey, err := cxsdk.CoralogixAPIKeyFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, apiKey)
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
