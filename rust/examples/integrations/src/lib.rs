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
        auth::ApiKey,
        client::integration::{IntegrationsClient, Parameter, StringList},
        CoralogixRegion,
    };

    #[tokio::test]
    async fn test_groups() {
        let client = IntegrationsClient::new(
            ApiKey::from_env().unwrap(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();

        let created = client
            .create(
                "".into(),
                Some("asdf"),
                Some(vec![Parameter {
                    key: "param1".to_string(),
                    value: StringList(vec!["value1".to_string()]),
                }]),
            )
            .await
            .unwrap();

        let group_id = created.group_id.unwrap();

        let retrieved_group = client.get(group_id.clone()).await.unwrap();
        assert_eq!(
            group_id.id,
            retrieved_group.group.unwrap().group_id.unwrap().id
        );

        client
            .update(
                group_id.clone(),
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
}
