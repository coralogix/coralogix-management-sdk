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

use std::str::FromStr;

use crate::{
    auth::{ApiKey, AuthData},
    error::Result,
    util::make_request_with_metadata,
};

pub use crate::com::coralogixapis::actions::v2::Action;

pub use cx_api::proto::com::coralogixapis::dashboards::v1::ast::*;
use cx_api::proto::com::coralogixapis::dashboards::v1::services::{
    dashboard_folders_service_client::DashboardFoldersServiceClient, CreateDashboardFolderRequest,
    CreateDashboardFolderResponse, DeleteDashboardFolderRequest, DeleteDashboardFolderResponse,
    GetDashboardFolderRequest, GetDashboardFolderResponse, ListDashboardFoldersRequest,
    ListDashboardFoldersResponse, ReplaceDashboardFolderRequest, ReplaceDashboardFolderResponse,
};
use tokio::sync::Mutex;
use tonic::{
    metadata::MetadataMap,
    transport::{Channel, Endpoint},
    Request,
};

pub use cx_api::proto::com::coralogixapis::dashboards::v1::common::DashboardFolder;

use crate::CoralogixRegion;

/// The Dashboard Folders API client.
/// Read more at <https://coralogix.com/docs/custom-dashboards/>
pub struct DashboardFoldersClient {
    metadata_map: MetadataMap,
    service_client: Mutex<DashboardFoldersServiceClient<Channel>>,
}

impl DashboardFoldersClient {
    /// Creates a new client for the Dashboard Folders API.
    ///
    /// # Arguments
    /// * `api_key` - The [`ApiKey`] to use for authentication.
    /// * `region` - The [`CoralogixRegion`] to connect to.
    pub fn new(api_key: ApiKey, region: CoralogixRegion) -> Result<Self> {
        let channel: Channel = Endpoint::from_str(region.grpc_endpoint().as_str())?.connect_lazy();
        let auth_data: AuthData = (&api_key).into();
        Ok(Self {
            metadata_map: auth_data.to_metadata_map(),
            service_client: Mutex::new(DashboardFoldersServiceClient::new(channel)),
        })
    }

    /// Creates a new dashboard folder.
    ///
    /// # Arguments
    /// * `folder` - The [`DashboardFolder`] to create.
    pub async fn create(&self, folder: DashboardFolder) -> Result<CreateDashboardFolderResponse> {
        let request: Request<CreateDashboardFolderRequest> = make_request_with_metadata(
            CreateDashboardFolderRequest {
                folder: Some(folder),
                request_id: None,
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .create_dashboard_folder(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    /// Replaces a dashboard folder.
    ///
    /// # Arguments
    /// * `folder` - The [`DashboardFolder`] to replace.
    pub async fn replace(&self, folder: DashboardFolder) -> Result<ReplaceDashboardFolderResponse> {
        let request: Request<ReplaceDashboardFolderRequest> = make_request_with_metadata(
            ReplaceDashboardFolderRequest {
                request_id: None,
                folder: Some(folder),
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .replace_dashboard_folder(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    /// Gets a dashboard folder by its ID.
    ///
    /// # Arguments
    /// * `folder_id` - The ID of the dashboard to get.
    pub async fn get(&self, folder_id: String) -> Result<GetDashboardFolderResponse> {
        let request: Request<GetDashboardFolderRequest> = make_request_with_metadata(
            GetDashboardFolderRequest {
                request_id: None,
                folder_id: Some(folder_id),
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .get_dashboard_folder(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    /// Deletes a dashboard folder by its ID.
    ///
    /// # Arguments
    /// * `folder_id` - The ID of the [`DashboardFolder`] to delete.
    pub async fn delete(&self, folder_id: String) -> Result<DeleteDashboardFolderResponse> {
        let request: Request<DeleteDashboardFolderRequest> = make_request_with_metadata(
            DeleteDashboardFolderRequest {
                request_id: None,
                folder_id: Some(folder_id),
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .delete_dashboard_folder(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    /// Lists all dashboard folders.
    pub async fn list(&self) -> Result<ListDashboardFoldersResponse> {
        let request = Request::new(ListDashboardFoldersRequest {});
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .list_dashboard_folders(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }
}
