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

#[cfg(test)]
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

        // Create folder
        let folder = ViewFolder {
            name: Some("RustTestViewFolder".to_string()),
            id: None,
        };

        let create_response = client
            .create(folder)
            .await
            .expect("Failed to create folder");
        let mut created_folder = create_response
            .folder
            .expect("Folder not found in response");

        // Update folder name
        created_folder.name = Some("RustTestViewFolderUpdated".to_string());
        client
            .replace(created_folder.clone())
            .await
            .expect("Failed to replace folder");

        // Get folder
        let get_response = client
            .get(created_folder.id.clone().unwrap())
            .await
            .expect("Failed to get folder");

        let updated_folder = get_response.expect("Folder not found");
        assert_eq!(updated_folder.name.unwrap(), "RustTestViewFolderUpdated");

        // Verify folder count increased
        let folders_after_create = client.list().await.expect("Failed to list folders");
        assert_eq!(folders_after_create.len(), initial_count + 1);

        // Delete folder
        client
            .delete(created_folder.id.unwrap())
            .await
            .expect("Failed to delete folder");

        // Verify folder count returned to initial
        let folders_after_delete = client.list().await.expect("Failed to list folders");
        assert_eq!(folders_after_delete.len(), initial_count);
    }
}
