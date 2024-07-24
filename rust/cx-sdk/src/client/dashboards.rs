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
    dashboards_service_client::DashboardsServiceClient, AssignDashboardFolderRequest,
    AssignDashboardFolderResponse, CreateDashboardRequest, CreateDashboardResponse,
    DeleteDashboardRequest, DeleteDashboardResponse, GetDashboardRequest, GetDashboardResponse,
    PinDashboardRequest, PinDashboardResponse, ReplaceDashboardRequest, ReplaceDashboardResponse,
    UnpinDashboardRequest, UnpinDashboardResponse,
};
use tokio::sync::Mutex;
use tonic::{
    metadata::MetadataMap,
    transport::{Channel, Endpoint},
    Request,
};

use crate::CoralogixRegion;

/// The Dashboards API client.
/// Read more at <https://coralogix.com/docs/custom-dashboards/>
pub struct DashboardsClient {
    metadata_map: MetadataMap,
    service_client: Mutex<DashboardsServiceClient<Channel>>,
}

impl DashboardsClient {
    /// Creates a new client for the Dashboards API.
    ///
    /// # Arguments
    /// * `api_key` - The [`ApiKey`] to use for authentication.
    /// * `region` - The [`CoralogixRegion`] to connect to.
    pub fn new(api_key: ApiKey, region: CoralogixRegion) -> Result<Self> {
        let channel: Channel = Endpoint::from_str(region.grpc_endpoint().as_str())?.connect_lazy();
        let auth_data: AuthData = (&api_key).into();
        Ok(Self {
            metadata_map: auth_data.to_metadata_map(),
            service_client: Mutex::new(DashboardsServiceClient::new(channel)),
        })
    }

    /// Creates a new dashboard.
    ///
    /// # Arguments
    /// * `dashboard` - The dashboard to create.
    pub async fn create(&self, dashboard: Dashboard) -> Result<CreateDashboardResponse> {
        let request: Request<CreateDashboardRequest> = make_request_with_metadata(
            CreateDashboardRequest {
                request_id: None,
                dashboard: Some(dashboard),
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .create_dashboard(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    /// Replaces a dashboard.
    ///
    /// # Arguments
    /// * `dashboard` - The [`Dashboard`] to replace.
    pub async fn replace(&self, dashboard: Dashboard) -> Result<ReplaceDashboardResponse> {
        let request: Request<ReplaceDashboardRequest> = make_request_with_metadata(
            ReplaceDashboardRequest {
                request_id: None,
                dashboard: Some(dashboard),
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .replace_dashboard(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    /// Gets a dashboard by its ID.
    ///
    /// # Arguments
    /// * `dashboard_id` - The ID of the dashboard to get.
    pub async fn get(&self, dashboard_id: String) -> Result<GetDashboardResponse> {
        let request: Request<GetDashboardRequest> = make_request_with_metadata(
            GetDashboardRequest {
                dashboard_id: Some(dashboard_id),
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .get_dashboard(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    /// Deletes a dashboard by its ID.
    ///
    /// # Arguments
    /// * `dashboard_id` - The ID of the dashboard to delete.
    pub async fn delete(&self, dashboard_id: String) -> Result<DeleteDashboardResponse> {
        let request: Request<DeleteDashboardRequest> = make_request_with_metadata(
            DeleteDashboardRequest {
                request_id: None,
                dashboard_id: Some(dashboard_id),
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .delete_dashboard(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    /// Pins a dashboard by its ID.
    ///
    /// # Arguments
    /// * `dashboard_id` - The ID of the dashboard to pin.
    pub async fn pin_dashboard(&self, dashboard_id: String) -> Result<PinDashboardResponse> {
        let request: Request<PinDashboardRequest> = make_request_with_metadata(
            PinDashboardRequest {
                request_id: None,
                dashboard_id: Some(dashboard_id),
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .pin_dashboard(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    /// Unpins a dashboard by its ID.
    ///
    /// # Arguments
    /// * `dashboard_id` - The ID of the dashboard to unpin.
    pub async fn unpin_dashboard(&self, dashboard_id: String) -> Result<UnpinDashboardResponse> {
        let request: Request<UnpinDashboardRequest> = make_request_with_metadata(
            UnpinDashboardRequest {
                request_id: None,
                dashboard_id: Some(dashboard_id),
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .unpin_dashboard(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    /// Assigns a dashboard to a folder.
    ///
    /// # Arguments
    /// * `dashboard_id` - The ID of the dashboard to assign.
    /// * `folder_id` - The ID of the folder to assign the dashboard to.
    pub async fn assign_dashboard_to_folder(
        &self,
        dashboard_id: String,
        folder_id: String,
    ) -> Result<AssignDashboardFolderResponse> {
        let request: Request<AssignDashboardFolderRequest> = make_request_with_metadata(
            AssignDashboardFolderRequest {
                request_id: None,
                dashboard_id: Some(dashboard_id),
                folder_id: Some(folder_id),
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .assign_dashboard_folder(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }
}
