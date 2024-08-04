package examples

import (
	"context"
	"testing"

	cxsdk "github.com/coralogix/coralogix-management-sdk/go"

	"github.com/stretchr/testify/assert"
)

func TestUsers(t *testing.T) {
	endpoint, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	apiKey, err := cxsdk.CoralogixAPIKeyFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(endpoint, apiKey)
	c := cxsdk.NewUsersClient(creator)

	createdUser, err := c.Create(context.Background(), &cxsdk.SCIMUser{
		Schemas:  []string{},
		UserName: "yak@coralogix.com",
		Active:   true,
		Name: &cxsdk.SCIMUserName{
			GivenName:  "example",
			FamilyName: "example",
		},
		Groups: []cxsdk.SCIMUserGroup{
			{Value: "Admins"},
		},
		Emails: []cxsdk.SCIMUserEmail{
			{Value: "example@coralogix.com", Primary: true, Type: "work"},
		},
	})

	assert.Nil(t, err)

	_, retrievalError := c.Get(context.Background(), *createdUser.ID)

	assert.Nil(t, retrievalError)

	deletionError := c.Delete(context.Background(), *createdUser.ID)

	assert.Nil(t, deletionError)

}
