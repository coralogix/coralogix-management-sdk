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

func TestUsers(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewSDKCallPropertiesCreator(region, authContext)

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
	t.Cleanup(func() {
		if createdUser != nil && createdUser.ID != nil {
			_ = c.Delete(context.Background(), *createdUser.ID)
		}
	})

	assertNilAndPrintError(t, err)

	_, retrievalError := c.Get(context.Background(), *createdUser.ID)

	assertNilAndPrintError(t, retrievalError)

	users, err := c.List(context.Background())
	assert.Nil(t, err)

	found := false
	for _, user := range users {
		if user.ID != nil && *user.ID == *createdUser.ID {
			found = true
			break
		}
	}
	assert.True(t, found)

	deletionError := c.Delete(context.Background(), *createdUser.ID)

	assertNilAndPrintError(t, deletionError)
}
