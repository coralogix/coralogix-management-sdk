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
        client::scopes::{EntityType, Filter, ScopesClient},
        CoralogixRegion,
    };

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
}
