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

use std::{
    str::FromStr,
    time::Duration,
};

use crate::{
    auth::AuthContext,
    error::{
        Result,
        SdkApiError,
        SdkError,
    },
    metadata::CallProperties,
    util::make_request_with_metadata,
};

pub use crate::com::coralogixapis::actions::v2::Action;

pub use cx_api::proto::com::coralogix::datausage::v2::{
    DateRange,
    Dimension,
    ScopesFilter,
};
use cx_api::proto::com::coralogix::datausage::v2::{
    GetDataUsageMetricsExportStatusRequest,
    GetDataUsageMetricsExportStatusResponse,
    GetDataUsageRequest,
    GetDataUsageResponse,
    GetLogsCountRequest,
    GetLogsCountResponse,
    GetSpansCountRequest,
    GetSpansCountResponse,
    UpdateDataUsageMetricsExportStatusRequest,
    UpdateDataUsageMetricsExportStatusResponse,
    data_usage_service_client::DataUsageServiceClient,
};
use tokio::sync::Mutex;
use tonic::{
    Streaming,
    metadata::MetadataMap,
    transport::{
        Channel,
        ClientTlsConfig,
        Endpoint,
    },
};

use crate::CoralogixRegion;

pub use crate::com::coralogix::enrichment::v1::{
    EnrichmentFieldDefinition,
    EnrichmentRequestModel as EnrichmentMapping,
    EnrichmentType,
    enrichment_type::Type,
};

const DATA_PLANS_FEATURE_GROUP_ID: &str = "dataplans";

/// A client for the Data Usage API.
/// Read more at <https://coralogix.com/docs/user-guides/account-management/payment-and-billing/data-usage/>
pub struct DataUsageClient {
    metadata_map: MetadataMap,
    service_client: Mutex<DataUsageServiceClient<Channel>>,
}

impl DataUsageClient {
    /// Creates a new client for the Enrichments API.
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
            service_client: Mutex::new(DataUsageServiceClient::new(channel)),
        })
    }

    /// Gets the data usage for the provided filters.
    ///
    /// # Arguments
    /// * `date_range` - The date range to get the data usage for.
    /// * `resolution` - The resolution to get the data usage for.
    /// * `aggregate` - The aggregate to get the data usage for.
    /// * `dimension_filters` - The dimension filters to apply.
    pub async fn get_data_usage(
        &self,
        date_range: Option<DateRange>,
        resolution: Option<Duration>,
        aggregate: Vec<i32>,
        dimension_filters: Vec<Dimension>,
    ) -> Result<Streaming<GetDataUsageResponse>> {
        self.service_client
            .lock()
            .await
            .get_data_usage(make_request_with_metadata(
                GetDataUsageRequest {
                    date_range,
                    resolution: resolution.map(|r| r.try_into().unwrap()),
                    aggregate,
                    dimension_filters,
                },
                &self.metadata_map,
            ))
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogix.datausage.v2.DataUsageService/GetDataUsage".into(),
                    feature_group: DATA_PLANS_FEATURE_GROUP_ID.into(),
                })
            })
    }

    /// Gets the span count for the provided filters.
    ///
    /// # Arguments
    /// * `date_range` - The date range to get the span count for.
    /// * `resolution` - The resolution to get the span count for.
    /// * `filters` - The filters to apply.
    pub async fn get_spans_count(
        &self,
        date_range: Option<DateRange>,
        resolution: Option<Duration>,
        filters: Option<ScopesFilter>,
    ) -> Result<Streaming<GetSpansCountResponse>> {
        Ok(self
            .service_client
            .lock()
            .await
            .get_spans_count(make_request_with_metadata(
                GetSpansCountRequest {
                    date_range,
                    resolution: resolution.map(|r| r.try_into().unwrap()),
                    filters,
                },
                &self.metadata_map,
            ))
            .await
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogix.datausage.v2.DataUsageService/GetSpansCount".into(),
                    feature_group: DATA_PLANS_FEATURE_GROUP_ID.into(),
                })
            })?
            .into_inner())
    }

    /// Gets the log count for the provided filters.
    ///
    /// # Arguments
    /// * `date_range` - The date range to get the log count for.
    /// * `resolution` - The resolution to get the log count for.
    /// * `filters` - The filters to apply.
    /// * `subsystem_aggregation` - Whether to aggregate subsystems.
    /// * `application_aggregation` - Whether to aggregate applications.
    pub async fn get_logs_count(
        &self,
        date_range: Option<DateRange>,
        resolution: Option<Duration>,
        filters: Option<ScopesFilter>,
        subsystem_aggregation: Option<bool>,
        application_aggregation: Option<bool>,
    ) -> Result<Streaming<GetLogsCountResponse>> {
        Ok(self
            .service_client
            .lock()
            .await
            .get_logs_count(make_request_with_metadata(
                GetLogsCountRequest {
                    date_range,
                    resolution: resolution.map(|r| r.try_into().unwrap()),
                    filters,
                    subsystem_aggregation,
                    application_aggregation,
                },
                &self.metadata_map,
            ))
            .await
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogix.datausage.v2.DataUsageService/GetLogsCount".into(),
                    feature_group: DATA_PLANS_FEATURE_GROUP_ID.into(),
                })
            })?
            .into_inner())
    }

    /// Updates the data usage metrics export status.
    ///
    /// # Arguments
    /// * `enabled` - Whether to enable or disable the data usage metrics export.
    pub async fn update_data_usage_metrics_export_status(
        &self,
        enabled: bool,
    ) -> Result<UpdateDataUsageMetricsExportStatusResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .update_data_usage_metrics_export_status(make_request_with_metadata(
                UpdateDataUsageMetricsExportStatusRequest {
                    enabled
                },
                &self.metadata_map,
            ))
            .await
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogix.datausage.v2.DataUsageService/UpdateDataUsageMetricsExportStatus"
                        .into(),
                        feature_group: DATA_PLANS_FEATURE_GROUP_ID.into(),
                })
            })?
            .into_inner())
    }

    /// Gets the data usage metrics export status.
    pub async fn get_data_usage_metrics_export_status(
        &self,
    ) -> Result<GetDataUsageMetricsExportStatusResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .get_data_usage_metrics_export_status(make_request_with_metadata(
                GetDataUsageMetricsExportStatusRequest {},
                &self.metadata_map,
            ))
            .await
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogix.datausage.v2.DataUsageService/GetDataUsageMetricsExportStatus"
                        .into(),
                        feature_group: DATA_PLANS_FEATURE_GROUP_ID.into(),
                })
            })?
            .into_inner())
    }
}
