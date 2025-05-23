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

mod tests {

    #[tokio::test]
    async fn test_extensions() {
        let extensions_client = ExtensionsClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();

        let get_all_extensions_response = extensions_client.get_all(None, true).await.unwrap();

        let extensions_descriptors = get_all_extensions_response.extensions;
        assert!(!extensions_descriptors.is_empty());

        let extension_to_deploy_id = extensions_descriptors
            .iter()
            .find(|extension| extension.name.as_ref().unwrap() == "Kubernetes Observability")
            .map(|extension| extension.id.clone())
            .unwrap()
            .unwrap();

        let get_deployed_extensions_response = extensions_client.get_deployed().await.unwrap();

        if let Some(extension_to_deploy) = get_deployed_extensions_response
            .deployed_extensions
            .iter()
            .find(|deployed_extension| {
                deployed_extension
                    .id
                    .as_ref()
                    .is_some_and(|id| id == &extension_to_deploy_id)
            })
        {
            extensions_client
                .undeploy(extension_to_deploy.id.as_ref().unwrap().clone(), vec![])
                .await
                .unwrap();
        }

        let get_extension_response = extensions_client
            .get(extension_to_deploy_id.clone())
            .await
            .unwrap();

        let extension_to_deploy = get_extension_response.extension.unwrap();

        let item_ids: Vec<String> = extension_to_deploy
            .revisions
            .iter()
            .find(|revision| revision.version.as_ref().unwrap() == "0.0.2")
            .map(|revision| {
                revision
                    .items
                    .iter()
                    .filter(|item| item.target_domain != TargetDomain::ParsingRule as i32)
                    .map(|item| item.id.clone().unwrap())
                    .collect()
            })
            .unwrap_or_default();

        extensions_client
            .deploy(
                extension_to_deploy_id,
                "0.0.2".to_string(),
                vec![],
                vec![],
                item_ids,
                None,
            )
            .await
            .unwrap();
    }
}
