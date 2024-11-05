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

	cxsdk "github.com/coralogix/coralogix-management-sdk/go"
	v1 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/permissions/v1"

	"github.com/stretchr/testify/assert"
)

func TestApiKeys(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	k := cxsdk.NewAPIKeysClient(creator)
	teamId, e := strconv.ParseUint(os.Getenv("TEAM_ID"), 10, 32)
	assert.Nil(t, e)

	key, e := k.Create(context.Background(), &cxsdk.CreateAPIKeyRequest{
		Name: "My APM KEY",
		Owner: &cxsdk.Owner{
			Owner: &cxsdk.OwnerTeamID{
				TeamId: uint32(teamId),
			},
		},
		KeyPermissions: &cxsdk.APIKeyPermissions{
			Presets:     []string{"APM"},
			Permissions: []string{},
		},
	})
	assert.Nil(t, e)

	newName := "new-name"
	k.Update(context.Background(), &cxsdk.UpdateAPIKeyRequest{
		KeyId:   key.KeyId,
		NewName: &newName,
	})

	updated, _ := k.Get(context.Background(), &cxsdk.GetAPIKeyRequest{
		KeyId: key.KeyId,
	})

	assert.Equal(t, updated.KeyInfo.Name, newName)

	k.Delete(context.Background(), &cxsdk.DeleteAPIKeyRequest{
		KeyId: key.KeyId,
	})
}

func TestUsers(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)

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
		Filters: []*cxsdk.ScopeFilter{
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

func TestSamlConfigurationRetrieval(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	c := cxsdk.NewSamlClient(creator)

	teamId, err := strconv.ParseUint(os.Getenv("TEAM_ID"), 10, 32)
	if err != nil {
		t.Error(err)
	}

	_, e := c.GetSPParameters(context.Background(), &cxsdk.GetSPParametersRequest{TeamId: uint32(teamId)})
	assert.Nil(t, e)

	_, e = c.GetConfiguration(context.Background(), &cxsdk.GetSamlConfigurationRequest{TeamId: uint32(teamId)})
	assert.Nil(t, e)

	_, e = c.SetActive(context.Background(), &cxsdk.SetSamlActiveRequest{TeamId: uint32(teamId), IsActive: false})
}

func TestSamlSetUpWithContent(t *testing.T) {
	t.Skip("Skipping integration test")

	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	c := cxsdk.NewSamlClient(creator)

	teamId, err := strconv.ParseUint(os.Getenv("TEAM_ID"), 10, 32)
	if err != nil {
		t.Error(err)
	}

	teamEntityId := uint32(teamId)

	_, e := c.SetIDPParameters(context.Background(), &cxsdk.SetIDPParametersRequest{
		TeamId: uint32(teamId),
		Params: &cxsdk.IDPParameters{
			Active:     true,
			GroupNames: []string{"ReadOnlyUsers"},
			Metadata: &cxsdk.IDPParametersContent{
				MetadataContent: "<?xml version=\"1.0\" encoding=\"UTF-8\"?><md:EntityDescriptor entityID=\"http://www.okta.com/<...>\" xmlns:md=\"urn:oasis:names:tc:SAML:2.0:metadata\"><md:IDPSSODescriptor WantAuthnRequestsSigned=\"false\" protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\"><md:KeyDescriptor use=\"signing\"><ds:KeyInfo xmlns:ds=\"http://www.w3.org/2000/09/xmldsig#\"><ds:X509Data><ds:X509Certificate>MIIDqDCCApCgAwIBAgIGAY1FD/bXMA0GCSqGSIb3DQEBCwUAMIGUMQswCQYDVQQGEwJVUzETMBEG\nA1UECAwKQ2FsaWZvcm5pYTEWMBQGA1UEBwwNU2FuIEZyYW5jaXNjbzENMAsGA1UECgwET2t0YTEU\nMBIGA1UECwwLU1NPUHJvdmlkZXIxFTATBgNVBAMMDGRldi01OTMyOTA1NzEcMBoGCSqGSIb3DQEJ\nARYNaW5mb0Bva3RhLmNvbTAeFw0yNDAxMjYwOTE3MTBaFw0zNDAxMjYwOTE4MTBaMIGUMQswCQYD\nVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTEWMBQGA1UEBwwNU2FuIEZyYW5jaXNjbzENMAsG\nA1UECgwET2t0YTEUMBIGA1UECwwLU1NPUHJvdmlkZXIxFTATBgNVBAMMDGRldi01OTMyOTA1NzEc\nMBoGCSqGSIb3DQEJARYNaW5mb0Bva3RhLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoC\nggEBAL7Ict7Pv1vRJsFRJCuygCYM/ELN+I6Am9vLQ8dbXUNphdD1qPxAXjjOR9zs9SetVZfrtAmw\n/o7zpJeIYEEQ9fVd2ayDY3lm2WgzK9NS3aO/9Lti0Z17Ppxih6S76FnQT3/4B5CRXNpw9cC11QGj\nmzNirZ3I2h6F9qNGZ3DSkyG6PdvcdX4J/AFcKqvm6l9dwfnRDV3pBUZjHMoR/IrwosEkLe20zxHM\nLrkKaxTk0hzXKSFWxWw+qJJv3IIMG02iVD59zxsVP07FsD5ThJ8tU+FWuAi+P//P3upHdqpfViXr\n7G9ydgVnedi2cIua78JAqcK8W0hzEpJgy89i0q4JwRUCAwEAATANBgkqhkiG9w0BAQsFAAOCAQEA\nNAQ2nQ7RSN1gW/pBvSxwSy7NkVEbVFygJCwBdaLE3ksqcz9wKsh7aL6AKNC44ry5CkONW8EZ2bGz\niVcZgg4fyEHNbBOnnUojadVszueOtijrnaiHCMwZRumhM9p/LJQ6trUvWZTsarYJdrLd+fDFtbfS\nMbKMSAt/jrmJ+okRCbu8yscB8BRcOuJ0tM0ZDstzCJ7O4P77o8iGTu5W8Itx0FMiy3aL3BT/7qaP\n1vYCJs5TFYHTaQe5GURPhJEQ8RCKy8WLN+KfyK1mz6slSmO/Jaqu6ppPc4YVPVClejLOv05hx0bs\nblLzVsAZsNpbTBo77bFopxaUk9fzuowO9xukpw==</ds:X509Certificate></ds:X509Data></ds:KeyInfo></md:KeyDescriptor><md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress</md:NameIDFormat><md:SingleSignOnService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://<...>.okta.com/app/<...>/sso/saml\"/><md:SingleSignOnService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\" Location=\"https://<...>.okta.com/app/<...>/sso/saml\"/></md:IDPSSODescriptor></md:EntityDescriptor>",
			},
			TeamEntityId: &teamEntityId,
		},
	})
	assert.Nil(t, e)
}

func TestSamlSetUpWithUrl(t *testing.T) {
	t.Skip("Skipping integration test")

	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	c := cxsdk.NewSamlClient(creator)

	teamId, err := strconv.ParseUint(os.Getenv("TEAM_ID"), 10, 32)
	if err != nil {
		t.Error(err)
	}

	teamEntityId := uint32(teamId)

	_, e := c.SetIDPParameters(context.Background(), &cxsdk.SetIDPParametersRequest{
		TeamId: uint32(teamId),
		Params: &cxsdk.IDPParameters{
			Active:     true,
			GroupNames: []string{"ReadOnlyUsers"},
			Metadata: &cxsdk.IDPParametersURL{
				MetadataUrl: "https://<...>.okta.com/app/<...>/sso/saml/metadata",
			},
			TeamEntityId: &teamEntityId,
		},
	})
	assert.Nil(t, e)
}
