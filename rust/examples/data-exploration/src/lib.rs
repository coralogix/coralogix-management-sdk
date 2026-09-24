// Copyright 2025 Coralogix Ltd.
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
    use cx_sdk::{
        CoralogixRegion,
        auth::AuthContext,
        client::{
            views::{
                QuickTimeSelection,
                SearchQuery,
                SelectionType,
                TimeSelection,
                View,
                ViewsClient,
            },
            views_folders::{
                ViewFolder,
                ViewFoldersClient,
            },
        },
    };

    #[tokio::test]
    async fn test_views() {
        let region = CoralogixRegion::from_env().expect("Failed to get region from env");
        let auth_context = AuthContext::from_env();

        let views_client = ViewsClient::new(auth_context.clone(), region.clone())
            .expect("Failed to create client");
        // Create a view
        let view = View {
            name: Some("RustTestView".to_string()),
            search_query: Some(SearchQuery {
                query: Some("source logs | filter $l.applicationname == 'default'".to_string()),
                ..Default::default()
            }),
            time_selection: Some(TimeSelection {
                selection_type: Some(SelectionType::QuickSelection(QuickTimeSelection {
                    seconds: 86400, // 24h
                    caption: None,
                })),
            }),
            filters: Some(cx_sdk::client::views::SelectedFilters { filters: vec![] }),
            id: None,
            folder_id: None,
            is_compact_mode: None,
            view_type: 0,
        };

        let create_response = views_client
            .create(view)
            .await
            .expect("Failed to create view");
        let mut created_view = create_response.view.expect("View not found in response");

        // Update view name
        created_view.name = Some("RustTestViewUpdated".to_string());
        views_client
            .replace(created_view.clone())
            .await
            .expect("Failed to replace view");

        // Get view
        let get_response = views_client
            .get(created_view.id.unwrap())
            .await
            .expect("Failed to get view");
        assert!(get_response.is_some());

        // Delete view
        views_client
            .delete(created_view.id.unwrap())
            .await
            .expect("Failed to delete view");
    }

    #[tokio::test]
    async fn test_view_folders() {
        let region = CoralogixRegion::from_env().expect("Failed to get region from env");
        let auth_context = AuthContext::from_env();

        let client = ViewFoldersClient::new(auth_context, region).expect("Failed to create client");

        // Get initial folder count
        let initial_folders = client.list().await.expect("Failed to list folders");
        let initial_count = initial_folders.len();

        let suffix = uuid::Uuid::new_v4();
        let folder_name = format!("RustTestViewFolder-{suffix}");
        let updated_name = format!("RustTestViewFolderUpdated-{suffix}");

        let create_response = client
            .create(ViewFolder {
                name: Some(folder_name),
                id: None,
            })
            .await
            .expect("Failed to create folder");
        let mut created_folder = create_response
            .folder
            .expect("Folder not found in response");
        let folder_id = created_folder
            .id
            .clone()
            .expect("Created folder is missing an id");

        created_folder.name = Some(updated_name.clone());
        let replace_result = client.replace(created_folder.clone()).await;
        let get_result = client.get(folder_id.clone()).await;
        let list_after_create = client.list().await;

        client
            .delete(folder_id)
            .await
            .expect("Failed to delete folder");

        replace_result.expect("Failed to replace folder");
        let updated_folder = get_result
            .expect("Failed to get folder")
            .expect("Folder not found");
        assert_eq!(updated_folder.name.unwrap(), updated_name);

        let folders_after_create = list_after_create.expect("Failed to list folders");
        assert_eq!(folders_after_create.len(), initial_count + 1);

        let folders_after_delete = client.list().await.expect("Failed to list folders");
        assert_eq!(folders_after_delete.len(), initial_count);
    }
}
