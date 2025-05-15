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

#[cfg(test)]
mod tests {
    use cx_sdk::{
        CoralogixRegion,
        auth::AuthContext,
        client::{
            apikeys::{
                ApiKeysClient,
                Owner,
            },
            groups::{
                GroupsClient,
                RoleId,
                TeamId,
            },
            saml::{
                IdpParameters,
                Metadata,
            },
            scopes::{
                EntityType,
                Filter,
                ScopesClient,
            },
            teams::TeamsClient,
            users::{
                ScimUser,
                ScimUserEmail,
                ScimUserGroup,
                ScimUserName,
                UsersClient,
            },
        },
    };
    #[tokio::test]
    async fn test_api_keys_client() {
        let team_id = std::env::var("TEAM_ID").unwrap().parse().unwrap();
        let client = ApiKeysClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();

        let create_response = client
            .create(
                "My APM KEY".to_string(),
                Some(Owner::TeamId(team_id)),
                vec!["APM".to_string()],
                vec![],
                false,
            )
            .await
            .unwrap();

        let key_id = create_response.key_id;

        let update_result = client
            .update(
                key_id.clone(),
                None,
                Some("new-name".to_string()),
                None,
                None,
            )
            .await;
        assert!(update_result.is_ok());

        let new_api_key = client.get(key_id.clone()).await;

        assert!(new_api_key.is_ok());
        assert_eq!(
            new_api_key.unwrap().key_info.map(|k| k.name),
            Some("new-name".to_string())
        );

        let delete_action_result = client.delete(key_id).await;
        assert!(delete_action_result.is_ok());
    }

    #[tokio::test]
    async fn test_groups() {
        let team_id_value = std::env::var("TEAM_ID").unwrap().parse::<u32>().unwrap();
        let client = GroupsClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();
        let team_id = TeamId { id: team_id_value };
        let groups = client.list(team_id).await.unwrap();
        assert!(!groups.groups.is_empty());

        let created_group = client
            .create(
                format!("Test group {}", chrono::Utc::now().timestamp_millis()),
                team_id,
                "A Test Group".to_string(),
                None,
                vec![RoleId { id: 1 }],
                vec![],
                None,
                None,
            )
            .await
            .unwrap();

        let group_id = created_group.group_id.unwrap();

        let retrieved_group = client.get(group_id).await.unwrap();
        assert_eq!(
            group_id.id,
            retrieved_group.group.unwrap().group_id.unwrap().id
        );

        client
            .update(
                group_id,
                "Updated Test Group".to_string(),
                "A Test Group".to_string(),
                None,
                None,
                None,
                None,
                None,
            )
            .await
            .unwrap();

        client.delete(group_id).await.unwrap();
    }

    #[tokio::test]
    async fn test_scopes_client() {
        let client = ScopesClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();

        let create = client
            .create(
                "Test Data Access Rule".into(),
                Some("Data Access Rule intended for testing".into()),
                vec![Filter {
                    entity_type: EntityType::Logs.into(),
                    expression: "<v1> foo == 'bar'".into(),
                }],
                "<v1>true".into(), // version -> <v1>true <- can only be true or false right now
            )
            .await
            .unwrap();

        let scope = create.scope.unwrap();

        let update_result = client
            .update(
                scope.id,
                "Updated Test Data Access Rule".into(),
                scope.description,
                scope.filters,
                scope.default_expression,
            )
            .await
            .unwrap();

        let new_scope = update_result.scope.unwrap();
        assert_eq!(new_scope.display_name, "Updated Test Data Access Rule");

        let _ = client.get(vec![new_scope.id.clone()]).await.unwrap();
        let deleted_scope_id = new_scope.id.clone();
        let _ = client.delete(new_scope.id).await.unwrap();

        let scopes_list = client.list().await.unwrap();
        for scope in scopes_list.scopes {
            assert_ne!(
                scope.id, deleted_scope_id,
                "Deleted scope should not be in the list"
            );
        }
    }

    #[tokio::test]
    #[ignore = "organization key is not available"]
    async fn test_teams() {
        let client = TeamsClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();

        let team_1_creation_result = client
            .create(
                "team_1".to_string(),
                vec!["admin@coralogix.com".into(), "admin2@coralogix.com".into()],
                Some(100f64),
            )
            .await;

        assert!(team_1_creation_result.is_ok());

        let team_1_id = team_1_creation_result.unwrap().team_id.unwrap();

        let team_2_creation_result = client
            .create(
                "team_2".to_string(),
                vec!["admin@coralogix.com".into(), "admin2@coralogix.com".into()],
                Some(200f64),
            )
            .await;

        assert!(team_2_creation_result.is_ok());

        let team_2_id = team_2_creation_result.unwrap().team_id.unwrap();

        let team_1_update_result = client
            .replace(team_1_id.id, Some("team_1_updated".into()), Some(150f64))
            .await;

        assert!(team_1_update_result.is_ok());

        let team_1_get_result = client.get(team_1_id.id).await;

        assert!(team_1_get_result.is_ok());

        assert!(team_1_get_result.unwrap().team_name == "team_1_updated");

        let set_daily_quota_result = client.set_daily_quota(team_1_id.id, 250f32).await;

        assert!(set_daily_quota_result.is_ok());

        let move_quota_result = client.move_quota(team_1_id.id, team_2_id.id, 50f64).await;

        assert!(move_quota_result.is_ok());
    }

    #[tokio::test]
    async fn test_users() {
        let client = UsersClient::new(
            CoralogixRegion::from_env().unwrap(),
            AuthContext::from_env(),
        );

        let scim_user = ScimUser {
            schemas: vec![],
            active: true,
            id: None,
            user_name: format!("yak{}@coralogix.com", chrono::Utc::now().timestamp_millis()),
            name: ScimUserName {
                given_name: "example".into(),
                family_name: "example".into(),
            },
            emails: vec![ScimUserEmail {
                value: format!(
                    "example{}@coralogix.com",
                    chrono::Utc::now().timestamp_millis()
                ),
                r#type: Some("work".into()),
                primary: true,
            }],
            groups: vec![ScimUserGroup {
                value: "Admins".into(),
            }],
        };

        let created_user = client.create(scim_user).await.unwrap();

        client
            .get(created_user.id.clone().unwrap().as_str())
            .await
            .unwrap();

        let users = client.list().await.unwrap();

        let found = users.iter().any(|user| {
            user.id
                .as_ref()
                .map_or(false, |id| id == &created_user.id.clone().unwrap())
        });

        assert!(found, "The created user was not found in the list.");

        client
            .delete(created_user.id.clone().unwrap().as_str())
            .await
            .unwrap();
    }

    #[tokio::test]
    async fn test_saml_configuration() {
        let client = cx_sdk::client::saml::SamlClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();

        let team_id = std::env::var("TEAM_ID").unwrap().parse::<u32>().unwrap();

        let _ = client.get_configuration(team_id).await.unwrap();

        let _ = client.get_sp_parameters(team_id).await.unwrap();

        client.set_active(team_id, false).await.unwrap();
    }

    #[ignore = "SAML integration testing is too complex to automate"]
    #[tokio::test]
    async fn test_saml_integration_with_content() {
        let client = cx_sdk::client::saml::SamlClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();

        let team_id = std::env::var("TEAM_ID").unwrap().parse::<u32>().unwrap();

        client
            .set_idp_parameters(
                team_id,
                Some(IdpParameters {
                    active: true,
                    team_entity_id: Some(team_id),
                    group_names: vec!["ReadOnlyUsers".to_string()],
                    metadata: Some(Metadata::MetadataContent("<?xml version=\"1.0\" encoding=\"UTF-8\"?><md:EntityDescriptor entityID=\"http://www.okta.com/<...>\" xmlns:md=\"urn:oasis:names:tc:SAML:2.0:metadata\"><md:IDPSSODescriptor WantAuthnRequestsSigned=\"false\" protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\"><md:KeyDescriptor use=\"signing\"><ds:KeyInfo xmlns:ds=\"http://www.w3.org/2000/09/xmldsig#\"><ds:X509Data><ds:X509Certificate>MIIDqDCCApCgAwIBAgIGAY1FD/bXMA0GCSqGSIb3DQEBCwUAMIGUMQswCQYDVQQGEwJVUzETMBEG\nA1UECAwKQ2FsaWZvcm5pYTEWMBQGA1UEBwwNU2FuIEZyYW5jaXNjbzENMAsGA1UECgwET2t0YTEU\nMBIGA1UECwwLU1NPUHJvdmlkZXIxFTATBgNVBAMMDGRldi01OTMyOTA1NzEcMBoGCSqGSIb3DQEJ\nARYNaW5mb0Bva3RhLmNvbTAeFw0yNDAxMjYwOTE3MTBaFw0zNDAxMjYwOTE4MTBaMIGUMQswCQYD\nVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTEWMBQGA1UEBwwNU2FuIEZyYW5jaXNjbzENMAsG\nA1UECgwET2t0YTEUMBIGA1UECwwLU1NPUHJvdmlkZXIxFTATBgNVBAMMDGRldi01OTMyOTA1NzEc\nMBoGCSqGSIb3DQEJARYNaW5mb0Bva3RhLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoC\nggEBAL7Ict7Pv1vRJsFRJCuygCYM/ELN+I6Am9vLQ8dbXUNphdD1qPxAXjjOR9zs9SetVZfrtAmw\n/o7zpJeIYEEQ9fVd2ayDY3lm2WgzK9NS3aO/9Lti0Z17Ppxih6S76FnQT3/4B5CRXNpw9cC11QGj\nmzNirZ3I2h6F9qNGZ3DSkyG6PdvcdX4J/AFcKqvm6l9dwfnRDV3pBUZjHMoR/IrwosEkLe20zxHM\nLrkKaxTk0hzXKSFWxWw+qJJv3IIMG02iVD59zxsVP07FsD5ThJ8tU+FWuAi+P//P3upHdqpfViXr\n7G9ydgVnedi2cIua78JAqcK8W0hzEpJgy89i0q4JwRUCAwEAATANBgkqhkiG9w0BAQsFAAOCAQEA\nNAQ2nQ7RSN1gW/pBvSxwSy7NkVEbVFygJCwBdaLE3ksqcz9wKsh7aL6AKNC44ry5CkONW8EZ2bGz\niVcZgg4fyEHNbBOnnUojadVszueOtijrnaiHCMwZRumhM9p/LJQ6trUvWZTsarYJdrLd+fDFtbfS\nMbKMSAt/jrmJ+okRCbu8yscB8BRcOuJ0tM0ZDstzCJ7O4P77o8iGTu5W8Itx0FMiy3aL3BT/7qaP\n1vYCJs5TFYHTaQe5GURPhJEQ8RCKy8WLN+KfyK1mz6slSmO/Jaqu6ppPc4YVPVClejLOv05hx0bs\nblLzVsAZsNpbTBo77bFopxaUk9fzuowO9xukpw==</ds:X509Certificate></ds:X509Data></ds:KeyInfo></md:KeyDescriptor><md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress</md:NameIDFormat><md:SingleSignOnService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://<...>.okta.com/app/<...>/sso/saml\"/><md:SingleSignOnService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\" Location=\"https://<...>.okta.com/app/<...>/sso/saml\"/></md:IDPSSODescriptor></md:EntityDescriptor>".to_string())),
                }),
            )
            .await
            .unwrap();
    }

    #[ignore = "SAML integration testing is too complex to automate"]
    #[tokio::test]
    async fn test_saml_integration_with_url() {
        let client = cx_sdk::client::saml::SamlClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();

        let team_id = std::env::var("TEAM_ID").unwrap().parse::<u32>().unwrap();

        client
            .set_idp_parameters(
                team_id,
                Some(IdpParameters {
                    active: true,
                    team_entity_id: Some(team_id),
                    group_names: vec!["ReadOnlyUsers".to_string()],
                    metadata: Some(Metadata::MetadataUrl(
                        "https://<...>.okta.com/app/<...>/sso/saml/metadata".to_string(),
                    )),
                }),
            )
            .await
            .unwrap();
    }
}
