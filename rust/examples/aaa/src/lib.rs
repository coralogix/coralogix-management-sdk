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
        CoralogixRegion,
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
                "Test Group".to_string(),
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

        let before_count = client.list().await.unwrap().scopes.len();
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
        let _ = client.delete(new_scope.id).await.unwrap();
        assert!(client.list().await.unwrap().scopes.len() - before_count == 0);
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
            user_name: "yak@coralogix.com".into(),
            name: ScimUserName {
                given_name: "example".into(),
                family_name: "example".into(),
            },
            emails: vec![ScimUserEmail {
                value: "example@coralogix.com".into(),
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

        client
            .delete(created_user.id.clone().unwrap().as_str())
            .await
            .unwrap();
    }
}
