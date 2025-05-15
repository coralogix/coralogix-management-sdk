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

use crate::error::{
    SdkApiError,
    SdkError,
};
use crate::{
    auth::AuthContext,
    error::Result,
    metadata::CallProperties,
    util::make_request_with_metadata,
};

pub use cx_api::proto::com::coralogixapis::slo::v1::{
    IsFilterPredicate,
    Metric,
    RequestBasedMetricSli,
    Slo,
    SloConstantFilterField,
    SloFilter,
    SloFilterField,
    SloFilterPredicate,
    SloFilters,
    SloTimeFrame,
    WindowBasedMetricSli,
    slo::{
        Sli,
        Window,
    },
    slo_filter_field::Field,
    slo_filter_predicate::Predicate,
};

use cx_api::proto::com::coralogixapis::slo::v1::slos_service_client::SlosServiceClient;
use cx_api::proto::com::coralogixapis::slo::v1::{
    BatchGetSlosRequest,
    BatchGetSlosResponse,
    CreateSloRequest,
    CreateSloResponse,
    DeleteSloRequest,
    DeleteSloResponse,
    GetSloRequest,
    GetSloResponse,
    ListSlosRequest,
    ListSlosResponse,
    ReplaceSloRequest,
    ReplaceSloResponse,
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

const INFRA_MONITORING_FEATURE_GROUP_ID: &str = "infra-monitoring";

/// The Service Line Objectives (SLO) client.
/// Read more at <https://coralogix.com/docs/slo-management-api/>
pub struct SloClient {
    metadata_map: MetadataMap,
    service_client: Mutex<SlosServiceClient<Channel>>,
}

impl SloClient {
    /// Creates a new client for the SLO.
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
            metadata_map: request_metadata.to_metadata_map(),
            service_client: Mutex::new(SlosServiceClient::new(channel)),
        })
    }

    /// Creates a new Service SLO.
    ///
    /// # Arguments
    /// * `slo` - The [`ServiceSlo`] to create.
    pub async fn create(&self, slo: Slo) -> Result<CreateSloResponse> {
        let request =
            make_request_with_metadata(CreateSloRequest { slo: Some(slo) }, &self.metadata_map);
        self.service_client
            .lock()
            .await
            .create_slo(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogixapis.slo.v1.ServiceSloService/CreateServiceSlo"
                        .to_string(),
                    feature_group: INFRA_MONITORING_FEATURE_GROUP_ID.to_string(),
                })
            })
    }

    /// Updates an existing SLO.
    ///
    /// # Arguments
    /// * `slo` - The [`ServiceSlo`] to update.    
    pub async fn update(&self, slo: Slo) -> Result<ReplaceSloResponse> {
        let request =
            make_request_with_metadata(ReplaceSloRequest { slo: Some(slo) }, &self.metadata_map);
        self.service_client
            .lock()
            .await
            .replace_slo(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogixapis.slo.v1.ServiceSloService/ReplaceServiceSlo"
                        .to_string(),
                    feature_group: INFRA_MONITORING_FEATURE_GROUP_ID.to_string(),
                })
            })
    }

    /// Deletes a Service SLO.
    ///
    /// # Arguments
    /// * `id` - The id of the Service SLO to delete.
    pub async fn delete(&self, id: String) -> Result<DeleteSloResponse> {
        let request = make_request_with_metadata(DeleteSloRequest { id }, &self.metadata_map);
        self.service_client
            .lock()
            .await
            .delete_slo(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogixapis.slo.v1.ServiceSloService/DeleteServiceSlo"
                        .to_string(),
                    feature_group: INFRA_MONITORING_FEATURE_GROUP_ID.to_string(),
                })
            })
    }

    /// Get the Service SLO.
    ///
    /// # Arguments
    /// * `id` - The ID of the Service SLO
    pub async fn get(&self, id: String) -> Result<GetSloResponse> {
        let request = make_request_with_metadata(GetSloRequest { id }, &self.metadata_map);

        self.service_client
            .lock()
            .await
            .get_slo(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogixapis.slo.v1.ServiceSloService/GetServiceSlo"
                        .to_string(),
                    feature_group: INFRA_MONITORING_FEATURE_GROUP_ID.to_string(),
                })
            })
    }

    /// Get the Service SLO in bulk.
    ///
    /// # Arguments
    /// * `ids` - The IDs of the Service SLO
    pub async fn get_bulk(&self, ids: Vec<String>) -> Result<BatchGetSlosResponse> {
        let request = make_request_with_metadata(BatchGetSlosRequest { ids }, &self.metadata_map);

        self.service_client
            .lock()
            .await
            .batch_get_slos(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogixapis.slo.v1.ServiceSloService/BatchGetServiceSlos"
                        .to_string(),
                    feature_group: INFRA_MONITORING_FEATURE_GROUP_ID.to_string(),
                })
            })
    }

    /// List the Service SLOs.
    ///
    /// # Arguments
    /// * `filter` - The [`SloFilters`] .
    pub async fn list(&self, filters: Option<SloFilters>) -> Result<ListSlosResponse> {
        let request = make_request_with_metadata(ListSlosRequest { filters }, &self.metadata_map);

        self.service_client
            .lock()
            .await
            .list_slos(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogixapis.slo.v1.ServiceSloService/ListServiceSlos"
                        .to_string(),
                    feature_group: INFRA_MONITORING_FEATURE_GROUP_ID.to_string(),
                })
            })
    }
}
