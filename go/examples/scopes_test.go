package examples

import (
	"context"
	cxsdk "coralogix-management-sdk/go"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScopes(t *testing.T) {
	creator := cxsdk.NewCallPropertiesCreator("https://ng-api-grpc.coralogix.com", "my-secret-token")
	c := cxsdk.NewScopesClient(creator)

	description := "Data Access Rule intended for testing"
	result, e := c.Create(context.Background(), &cxsdk.CreateScopeRequest{
		DisplayName: "Test Data Access Rule",
		Description: &description,
		Filters: []*cxsdk.Filter{
			{EntityType: cxsdk.EntityTypeLogs, Expression: "<v1> foo == 'bar'"},
		},
		DefaultExpression: "<v1> foo == 'bar'",
	})
	assert.Nil(t, e)

	_, e = c.Update(context.Background(), &cxsdk.UpdateScopeRequest{
		Id:          result.Scope.Id,
		DisplayName: "Updated Test Data Access Rule",
	})
	assert.Nil(t, e)

	updated, _ := c.Get(context.Background(), &cxsdk.GetTeamScopesByIDsRequest{
		Ids: []string{result.Scope.Id},
	})

	assert.Equal(t, updated.Scopes[0].DisplayName, "Updated Test Data Access Rule")

	c.Delete(context.Background(), &cxsdk.DeleteScopeRequest{
		Id: result.Scope.Id,
	})

	all, e := c.List(context.Background(), &cxsdk.GetTeamScopesRequest{})
	assert.Nil(t, e)
	assert.Empty(t, all.Scopes)
}
