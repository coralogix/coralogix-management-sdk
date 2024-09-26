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
        client::apikeys::{
            ApiKeysClient,
            Owner,
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
}
