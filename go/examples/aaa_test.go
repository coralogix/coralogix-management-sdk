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
	"time"

	cxsdk "github.com/coralogix/coralogix-management-sdk/go"

	"github.com/stretchr/testify/assert"
)

func TestApiKeys(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewSDKCallPropertiesCreator(region, authContext)
	k := cxsdk.NewAPIKeysClient(creator)
	teamId, e := strconv.ParseUint(os.Getenv("TEAM_ID"), 10, 32)
	assertNilAndPrintError(t, e)

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
	assertNilAndPrintError(t, e)

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

func TestScopes(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewSDKCallPropertiesCreator(region, authContext)
	c := cxsdk.NewScopesClient(creator)
	description := "Data Access Rule intended for testing"
	result, e := c.Create(context.Background(), &cxsdk.CreateScopeRequest{
		DisplayName: "Test Data Access Rule",
		Description: &description,
		Filters: []*cxsdk.ScopeFilter{
			{EntityType: cxsdk.EntityTypeLogs, Expression: "<v1> foo == 'bar'"},
		},
		DefaultExpression: "<v1>true",
	})
	assertNilAndPrintError(t, e)

	_, e = c.Update(context.Background(), &cxsdk.UpdateScopeRequest{
		Id:                result.Scope.Id,
		DisplayName:       "Updated Test Data Access Rule",
		Filters:           result.Scope.Filters,
		DefaultExpression: result.Scope.DefaultExpression,
	})
	assertNilAndPrintError(t, e)

	updated, _ := c.Get(context.Background(), &cxsdk.GetTeamScopesByIDsRequest{
		Ids: []string{result.Scope.Id},
	})

	assert.Equal(t, "Updated Test Data Access Rule", updated.Scopes[0].DisplayName)

	c.Delete(context.Background(), &cxsdk.DeleteScopeRequest{
		Id: result.Scope.Id,
	})

	scopes, err := c.List(context.Background(), &cxsdk.GetTeamScopesRequest{})
	assert.Nil(t, err)

	for _, scope := range scopes.Scopes {
		assert.NotEqual(t, result.Scope.Id, scope.Id)
	}
}

func TestGroups(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewSDKCallPropertiesCreator(region, authContext)
	teamId, err := strconv.ParseUint(os.Getenv("TEAM_ID"), 10, 32)
	assertNilAndPrintError(t, err)
	c := cxsdk.NewGroupsClient(creator)

	groups, err := c.List(context.Background(), &cxsdk.GetTeamGroupsRequest{
		TeamId: &cxsdk.GroupsTeamID{
			Id: uint32(teamId),
		},
	})
	assertNilAndPrintError(t, err)
	assert.Greater(t, len(groups.Groups), 0)

	groupDesc := "A Test Group"

	createdGroup, err := c.Create(context.Background(), &cxsdk.CreateTeamGroupRequest{
		Name: "Test Group " + strconv.FormatInt(time.Now().UnixMilli(), 10),
		TeamId: &cxsdk.GroupsTeamID{
			Id: uint32(teamId),
		},
		Description: &groupDesc,
		ExternalId:  nil,
		RoleIds: []*cxsdk.RoleID{
			{Id: 1},
		},
		UserIds: []*cxsdk.UserID{},
	})

	assertNilAndPrintError(t, err)

	retrievedGroup, err := c.Get(context.Background(), &cxsdk.GetTeamGroupRequest{
		GroupId: &cxsdk.TeamGroupID{
			Id: createdGroup.GroupId.Id,
		},
	})

	assertNilAndPrintError(t, err)
	assert.Equal(t, createdGroup.GroupId.Id, retrievedGroup.Group.GroupId.Id)

	_, updateError := c.Update(context.Background(), &cxsdk.UpdateTeamGroupRequest{
		GroupId:     createdGroup.GroupId,
		Name:        "Updated Test Group " + strconv.FormatInt(time.Now().UnixMilli(), 10),
		Description: &groupDesc,
	})

	assertNilAndPrintError(t, updateError)

	_, deletionError := c.Delete(context.Background(), &cxsdk.DeleteTeamGroupRequest{
		GroupId: &cxsdk.TeamGroupID{
			Id: createdGroup.GroupId.Id,
		},
	})

	assertNilAndPrintError(t, deletionError)
}

func TestTeams(t *testing.T) {

	t.Skip("Skipping integration test")

	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	creator := cxsdk.NewSDKCallPropertiesCreator(region, authContext)

	c := cxsdk.NewTeamsClient(creator)

	team1, e := c.Create(context.Background(), &cxsdk.CreateTeamInOrgRequest{
		TeamName: "team_1",
		TeamAdminsEmail: []string{
			"admin@coralogix.com",
			"admin2@coralogix.com",
		},
	})
	assertNilAndPrintError(t, e)

	team2, e := c.Create(context.Background(), &cxsdk.CreateTeamInOrgRequest{
		TeamName: "team_2",
		TeamAdminsEmail: []string{
			"admin@coralogix.com",
			"admin2@coralogix.com",
		},
	})
	assertNilAndPrintError(t, e)

	team_1_updated_name := "team_1_updated"
	_, e = c.Update(context.Background(), &cxsdk.UpdateTeamRequest{
		TeamName: &team_1_updated_name,
	})

	assertNilAndPrintError(t, e)

	updated, e := c.Get(context.Background(), &cxsdk.GetTeamRequest{
		TeamId: team1.TeamId,
	})

	assertNilAndPrintError(t, e)

	assert.Equal(t, updated.TeamName, "team_1_updated")

	_, setQuotaError := c.SetDailyQuota(context.Background(), &cxsdk.SetDailyQuotaRequest{
		TeamId:           team1.TeamId,
		TargetDailyQuota: 250,
	})

	assertNilAndPrintError(t, setQuotaError)

	_, moveQuotaError := c.MoveQuota(context.Background(), &cxsdk.MoveQuotaRequest{
		SourceTeam:      team1.TeamId,
		DestinationTeam: team2.TeamId,
		UnitsToMove:     50,
	})

	assertNilAndPrintError(t, moveQuotaError)

	_, deletionError := c.Delete(context.Background(), &cxsdk.DeleteTeamRequest{
		TeamId: team1.TeamId,
	})

	assertNilAndPrintError(t, deletionError)
}

func TestSamlConfigurationRetrieval(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewSDKCallPropertiesCreator(region, authContext)
	c := cxsdk.NewSamlClient(creator)

	teamId, err := strconv.ParseUint(os.Getenv("TEAM_ID"), 10, 32)
	if err != nil {
		t.Error(err)
	}

	_, e := c.GetSPParameters(context.Background(), &cxsdk.GetSPParametersRequest{TeamId: uint32(teamId)})
	assertNilAndPrintError(t, e)

	_, e = c.GetConfiguration(context.Background(), &cxsdk.GetSamlConfigurationRequest{TeamId: uint32(teamId)})
	assertNilAndPrintError(t, e)

	_, _ = c.SetActive(context.Background(), &cxsdk.SetSamlActiveRequest{TeamId: uint32(teamId), IsActive: false})
}

func TestSamlSetUpWithContent(t *testing.T) {
	t.Skip("Skipping integration test")

	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewSDKCallPropertiesCreator(region, authContext)
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
	assertNilAndPrintError(t, e)
}

func TestSamlSetUpWithUrl(t *testing.T) {
	t.Skip("Skipping integration test")

	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewSDKCallPropertiesCreator(region, authContext)
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
	assertNilAndPrintError(t, e)
}

func TestIpAccess(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewSDKCallPropertiesCreator(region, authContext)
	c := cxsdk.NewIPAccessClient(creator)

	// replace without ID, so it will create a new settings in case it doesn't exist
	// if it exists, it will replace the existing settings
	createRes, err := c.Replace(context.Background(), &cxsdk.ReplaceCompanyIPAccessSettingsRequest{
		IpAccess: []*cxsdk.IPAccess{
			{
				Name:    "Office Network",
				IpRange: "31.154.215.114/32",
				Enabled: false,
			},
			{
				Name:    "VPN",
				IpRange: "198.51.100.0/24",
				Enabled: false,
			},
		},
		EnableCoralogixCustomerSupportAccess: cxsdk.CoralogixCustomerSupportAccessDisabled,
	})
	assertNilAndPrintError(t, err)

	_, err = c.Get(context.Background(), &cxsdk.GetCompanyIPAccessSettingsRequest{
		Id: &createRes.Settings.Id,
	})
	assertNilAndPrintError(t, err)

	replaceRes, err := c.Replace(context.Background(), &cxsdk.ReplaceCompanyIPAccessSettingsRequest{
		Id: &createRes.Settings.Id,
		IpAccess: []*cxsdk.IPAccess{
			{
				Name:    "Office Network",
				IpRange: "31.154.215.114/32",
				Enabled: false,
			},
			{
				Name:    "VPN",
				IpRange: "198.51.100.0/18",
				Enabled: false,
			},
		},
		EnableCoralogixCustomerSupportAccess: cxsdk.CoralogixCustomerSupportAccessDisabled,
	})
	assertNilAndPrintError(t, err)

	found := false
	for _, ipAccess := range replaceRes.Settings.IpAccess {
		if ipAccess.Name == "VPN" {
			assert.Equal(t, "198.51.100.0/18", ipAccess.IpRange)
		}
		found = true
		break
	}
	if !found {
		t.Error("Expected 'VPN' entry not found in IpAccess")
	}

	_, err = c.Delete(context.Background(), &cxsdk.DeleteCompanyIPAccessSettingsRequest{
		Id: &replaceRes.Settings.Id,
	})
	assertNilAndPrintError(t, err)
}
