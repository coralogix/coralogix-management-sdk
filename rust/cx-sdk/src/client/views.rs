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
        CreateViewRequest,
        CreateViewResponse,
        DeleteViewRequest,
        DeleteViewResponse,
        GetViewRequest,
        ListViewsRequest,
        ReplaceViewRequest,
        ReplaceViewResponse,
        views_service_client::ViewsServiceClient,
    },
    error::{
        Result,
        SdkApiError,
    },
    metadata::CallProperties,
    util::make_request_with_metadata,
};

pub use crate::com::coralogixapis::views::v1::{
    CustomTimeSelection,
    Filter,
    QuickTimeSelection,
    SearchQuery,
    SelectedFilters,
    TimeSelection,
    services::View,
    time_selection::SelectionType,
};

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

/// The Views API client.
/// Read more at <https://coralogix.com/docs/user-guides/monitoring-and-insights/explore-screen/custom-views//>
pub struct ViewsClient {
    teams_level_metadata_map: MetadataMap,
    service_client: Mutex<ViewsServiceClient<Channel>>,
}

impl ViewsClient {
    /// Creates a new client for the Views API.
    ///
    /// # Arguments
    /// * `auth_context` - The [`AuthContext`] to use for authentication.
    /// * `region` - The [`CoralogixRegion`] to connect to.
    pub fn new(auth_context: AuthContext, region: CoralogixRegion) -> Result<Self> {
        let channel: Channel = Endpoint::from_str(region.grpc_endpoint().as_str())?
            .tls_config(ClientTlsConfig::new().with_native_roots())?
            .connect_lazy();
        let request_metadata: CallProperties = (&auth_context.user_level_api_key).into();
        Ok(Self {
            teams_level_metadata_map: request_metadata.to_metadata_map(),
            service_client: Mutex::new(ViewsServiceClient::new(channel)),
        })
    }

    /// Creates a new View
    ///
    /// # Arguments
    /// * `view` - The [`View`] to create.
    pub async fn create(&self, view: View) -> Result<CreateViewResponse> {
        let request = make_request_with_metadata(
            CreateViewRequest {
                name: view.name,
                search_query: view.search_query,
                time_selection: view.time_selection,
                filters: view.filters,
                folder_id: view.folder_id,
                view_type: view.view_type,
            },
            &self.teams_level_metadata_map,
        );
        Ok(self
            .service_client
            .lock()
            .await
            .create_view(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.views.v1.services.ViewsService/CreateView".into(),
                feature_group: FEATURE_GROUP_ID.into(),
            })?
            .into_inner())
    }

    /// Replaces the existing [`View`] identified by its id.
    ///
    /// # Arguments
    /// * `view` - The view to replace.
    pub async fn replace(&self, view: View) -> Result<ReplaceViewResponse> {
        let request = make_request_with_metadata(
            ReplaceViewRequest { view: Some(view) },
            &self.teams_level_metadata_map,
        );
        Ok(self
            .service_client
            .lock()
            .await
            .replace_view(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.views.v1.services.ViewsService/ReplaceView"
                    .to_string(),
                feature_group: FEATURE_GROUP_ID.to_string(),
            })?
            .into_inner())
    }

    /// Deletes the [`View`] identified by its id.
    ///
    /// # Arguments
    /// * `id` - The id of the view to delete.
    pub async fn delete(&self, id: i32) -> Result<DeleteViewResponse> {
        let request = make_request_with_metadata(
            DeleteViewRequest { id: Some(id) },
            &self.teams_level_metadata_map,
        );
        Ok(self
            .service_client
            .lock()
            .await
            .delete_view(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.views.v1.services.ViewsService/DeleteView"
                    .to_string(),
                feature_group: FEATURE_GROUP_ID.to_string(),
            })?
            .into_inner())
    }

    /// Retrieves the [`View`] by id.
    ///
    /// # Arguments
    /// * `id` - The id of the view to retrieve.
    pub async fn get(&self, id: i32) -> Result<Option<View>> {
        let request = make_request_with_metadata(
            GetViewRequest { id: Some(id) },
            &self.teams_level_metadata_map,
        );

        Ok(self
            .service_client
            .lock()
            .await
            .get_view(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.views.v1.services.ViewsService/GetView".to_string(),
                feature_group: FEATURE_GROUP_ID.to_string(),
            })?
            .into_inner()
            .view)
    }

    /// Retrieves a list of all [`View`]s.
    ///     
    /// # Returns
    /// A list of all views.
    pub async fn list(&self) -> Result<Vec<View>> {
        let request =
            make_request_with_metadata(ListViewsRequest {}, &self.teams_level_metadata_map);

        Ok(self
            .service_client
            .lock()
            .await
            .list_views(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.views.v1.services.ViewsService/ListViews".to_string(),
                feature_group: FEATURE_GROUP_ID.to_string(),
            })?
            .into_inner()
            .views)
    }
}
