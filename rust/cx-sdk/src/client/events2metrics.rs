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
    auth::AuthContext,
    error::{
        Result,
        SdkApiError,
    },
    metadata::CallProperties,
    util::make_request_with_metadata,
};

use cx_api::proto::com::coralogixapis::events2metrics::v2::{
    CreateE2mRequest,
    CreateE2mResponse,
    DeleteE2mRequest,
    DeleteE2mResponse,
    GetE2mRequest,
    GetE2mResponse,
    GetLimitsRequest,
    GetLimitsResponse,
    ListE2mRequest,
    ListE2mResponse,
    ListLabelsCardinalityRequest,
    ListLabelsCardinalityResponse,
    ReplaceE2mRequest,
    ReplaceE2mResponse,
    events2_metric_service_client::Events2MetricServiceClient,
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

pub use cx_api::proto::com::coralogixapis::events2metrics::v2::{
    E2m,
    E2mCreateParams,
    MetricField,
    MetricLabel,
    list_labels_cardinality_request::Query,
};

use crate::CoralogixRegion;

/// The Events2Metrics API client.
/// Read more at <https://coralogix.com/docs/events2metrics/>
pub struct Events2MetricsClient {
    metadata_map: MetadataMap,
    service_client: Mutex<Events2MetricServiceClient<Channel>>,
}

impl Events2MetricsClient {
    /// Creates a new client for the Events2Metrics API.
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
            service_client: Mutex::new(Events2MetricServiceClient::new(channel)),
        })
    }

    /// Creates a new E2M mapping.
    ///
    /// # Arguments
    /// * `params` - The parameters for the E2M mapping.
    pub async fn create(&self, params: E2mCreateParams) -> Result<CreateE2mResponse> {
        let request =
            make_request_with_metadata(CreateE2mRequest { e2m: Some(params) }, &self.metadata_map);
        Ok(self
            .service_client
            .lock()
            .await
            .create_e2m(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.events2metrics.v2.Events2MetricService/CreateE2M"
                    .to_string(),
            })?
            .into_inner())
    }

    /// Replaces an existing E2M mapping.
    ///
    /// # Arguments
    /// * `e2m` - The E2M mapping to replace.
    pub async fn replace(&self, e2m: E2m) -> Result<ReplaceE2mResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .replace_e2m(make_request_with_metadata(
                ReplaceE2mRequest { e2m: Some(e2m) },
                &self.metadata_map,
            ))
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.events2metrics.v2.Events2MetricService/ReplaceE2M"
                    .to_string(),
            })?
            .into_inner())
    }

    /// Deletes an E2M mapping.
    ///
    /// # Arguments
    /// * `id` - The ID of the E2M mapping to delete.
    pub async fn delete(&self, id: String) -> Result<DeleteE2mResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .delete_e2m(make_request_with_metadata(
                DeleteE2mRequest { id: Some(id) },
                &self.metadata_map,
            ))
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.events2metrics.v2.Events2MetricService/DeleteE2M"
                    .to_string(),
            })?
            .into_inner())
    }

    /// Retrieves an E2M mapping by its ID.
    ///
    /// # Arguments
    /// * `id` - The ID of the E2M mapping to get.
    pub async fn get(&self, id: String) -> Result<GetE2mResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .get_e2m(make_request_with_metadata(
                GetE2mRequest { id: Some(id) },
                &self.metadata_map,
            ))
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.events2metrics.v2.Events2MetricService/GetE2M"
                    .to_string(),
            })?
            .into_inner())
    }

    /// Retrieves the limits associated with this account.
    ///
    /// # Returns
    /// The limits associated with this account.
    pub async fn get_limits(&self) -> Result<GetLimitsResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .get_limits(make_request_with_metadata(
                GetLimitsRequest {},
                &self.metadata_map,
            ))
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.events2metrics.v2.Events2MetricService/GetLimits"
                    .to_string(),
            })?
            .into_inner())
    }

    /// Lists all E2M mappings.
    ///
    /// # Returns
    /// A list of all E2M mappings.
    pub async fn list(&self) -> Result<ListE2mResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .list_e2m(make_request_with_metadata(
                ListE2mRequest {},
                &self.metadata_map,
            ))
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.events2metrics.v2.Events2MetricService/ListE2M"
                    .to_string(),
            })?
            .into_inner())
    }

    /// Lists the cardinality of labels for a given metric.
    ///
    /// # Arguments
    /// * `metric_labels` - The labels of the metric to list cardinality for.
    /// * `query` - The query to filter the labels by.
    pub async fn list_labels_cardinality(
        &self,
        metric_labels: Vec<MetricLabel>,
        query: Query,
    ) -> Result<ListLabelsCardinalityResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .list_labels_cardinality(make_request_with_metadata(
                ListLabelsCardinalityRequest {
                    metric_labels,
                    query: Some(query),
                },
                &self.metadata_map,
            ))
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.events2metrics.v2.Events2MetricService/ListLabelsCardinality".to_string(),
            })?
            .into_inner())
    }

    // TODO: Expose the AtomicBatchExecuteE2m API where CRUD requests can be batched together.
}
