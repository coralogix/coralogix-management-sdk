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

use std::str::FromStr;

use crate::{
    auth::AuthContext,
    com::coralogixapis::views::v1::services::{
        CreateViewFolderRequest,
        CreateViewFolderResponse,
        DeleteViewFolderRequest,
        DeleteViewFolderResponse,
        GetViewFolderRequest,
        ListViewFoldersRequest,
        ReplaceViewFolderRequest,
        ReplaceViewFolderResponse,
        views_folders_service_client::ViewsFoldersServiceClient,
    },
    error::{
        Result,
        SdkApiError,
    },
    metadata::CallProperties,
    util::make_request_with_metadata,
};

pub use crate::com::coralogixapis::views::v1::ViewFolder;

use tokio::sync::Mutex;
use tonic::{
    metadata::MetadataMap,
    transport::{
        Channel,
        ClientTlsConfig,
        Endpoint,
    },
};

use crate::CoralogixRegion;

const FEATURE_GROUP_ID: &str = "data-exploration";

/// The ViewFolders API client.
/// Read more at <https://coralogix.com/docs/user-guides/monitoring-and-insights/explore-screen/custom-views/>
pub struct ViewFoldersClient {
    teams_level_metadata_map: MetadataMap,
    service_client: Mutex<ViewsFoldersServiceClient<Channel>>,
}

impl ViewFoldersClient {
    /// Creates a new client for the ViewFolders API.
    ///
    /// # Arguments
    /// * `auth_context` - The [`AuthContext`] to use for authentication.
    /// * `region` - The [`CoralogixRegion`] to connect to.
    pub fn new(auth_context: AuthContext, region: CoralogixRegion) -> Result<Self> {
        let channel: Channel = Endpoint::from_str(region.grpc_endpoint().as_str())?
            .tls_config(ClientTlsConfig::new().with_native_roots())?
            .connect_lazy();
        let request_metadata: CallProperties = (&auth_context.team_level_api_key).into();
        Ok(Self {
            teams_level_metadata_map: request_metadata.to_metadata_map(),
            service_client: Mutex::new(ViewsFoldersServiceClient::new(channel)),
        })
    }

    /// Creates a new ViewFolder
    ///
    /// # Arguments
    /// * `folder` - The [`ViewFolder`] to create.
    pub async fn create(&self, folder: ViewFolder) -> Result<CreateViewFolderResponse> {
        let request = make_request_with_metadata(
            CreateViewFolderRequest { name: folder.name },
            &self.teams_level_metadata_map,
        );
        Ok(self
            .service_client
            .lock()
            .await
            .create_view_folder(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint:
                    "/com.coralogixapis.views.v1.services.ViewsFoldersService/CreateViewFolder"
                        .into(),
                feature_group: FEATURE_GROUP_ID.into(),
            })?
            .into_inner())
    }

    /// Replaces the existing [`ViewFolder`] identified by its id.
    ///
    /// # Arguments
    /// * `folder` - The view folder to replace.
    pub async fn replace(&self, folder: ViewFolder) -> Result<ReplaceViewFolderResponse> {
        let request = make_request_with_metadata(
            ReplaceViewFolderRequest {
                folder: Some(folder),
            },
            &self.teams_level_metadata_map,
        );
        Ok(self
            .service_client
            .lock()
            .await
            .replace_view_folder(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.views.v1.services.ViewFoldersService/ReplaceView"
                    .to_string(),
                feature_group: FEATURE_GROUP_ID.to_string(),
            })?
            .into_inner())
    }

    /// Deletes the [`ViewFolder`] identified by its id.
    ///
    /// # Arguments
    /// * `id` - The id of the view folder to delete.
    pub async fn delete(&self, id: String) -> Result<DeleteViewFolderResponse> {
        let request = make_request_with_metadata(
            DeleteViewFolderRequest { id: Some(id) },
            &self.teams_level_metadata_map,
        );
        Ok(self
            .service_client
            .lock()
            .await
            .delete_view_folder(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint:
                    "/com.coralogixapis.views.v1.services.ViewsFoldersService/DeleteViewFolder"
                        .to_string(),
                feature_group: FEATURE_GROUP_ID.to_string(),
            })?
            .into_inner())
    }

    /// Retrieves the [`ViewFolder`] by id.
    ///
    /// # Arguments
    /// * `id` - The id of the view folder to retrieve.
    pub async fn get(&self, id: String) -> Result<Option<ViewFolder>> {
        let request = make_request_with_metadata(
            GetViewFolderRequest { id: Some(id) },
            &self.teams_level_metadata_map,
        );

        Ok(self
            .service_client
            .lock()
            .await
            .get_view_folder(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.views.v1.services.ViewsFoldersService/GetViewFolder"
                    .to_string(),
                feature_group: FEATURE_GROUP_ID.to_string(),
            })?
            .into_inner()
            .folder)
    }

    /// Retrieves a list of all [`ViewFolder`]s.
    ///     
    /// # Returns
    /// A list of all views.
    pub async fn list(&self) -> Result<Vec<ViewFolder>> {
        let request =
            make_request_with_metadata(ListViewFoldersRequest {}, &self.teams_level_metadata_map);

        Ok(self
            .service_client
            .lock()
            .await
            .list_view_folders(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint:
                    "/com.coralogixapis.views.v1.services.ViewsFoldersService/ListViewFolders"
                        .to_string(),
                feature_group: FEATURE_GROUP_ID.to_string(),
            })?
            .into_inner()
            .folders)
    }
}
