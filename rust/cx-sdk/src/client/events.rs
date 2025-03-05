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

use tokio::sync::Mutex;
use tonic::transport::{
    Channel,
    ClientTlsConfig,
    Endpoint,
};

use cx_api::proto::com::coralogixapis::events::v3::{
    BatchGetEventRequest,
    BatchGetEventResponse,
    GetEventRequest,
    GetEventResponse,
    GetEventsStatisticsRequest,
    GetEventsStatisticsResponse,
    ListEventsCountRequest,
    ListEventsCountResponse,
    ListEventsRequest,
    ListEventsResponse,
    events_service_client::EventsServiceClient,
};

pub use cx_api::proto::com::coralogixapis::events::v3::{
    EventsFilter,
    EventsQueryFilter,
    FilterOperator,
    Filters,
    OrderBy,
    OrderByDirection,
    OrderByFields,
    PaginationRequest,
    TimestampRange,
};

use crate::{
    CoralogixRegion,
    auth::AuthContext,
    error::{
        Result,
        SdkApiError,
    },
    metadata::CallProperties,
    util::make_request_with_metadata,
};

const EVENTS_FEATURE_GROUP_ID: &str = "events";

/// The Events API client.
pub struct EventsClient {
    metadata_map: tonic::metadata::MetadataMap,
    service_client: Mutex<EventsServiceClient<Channel>>,
}

impl EventsClient {
    /// Creates a new client for the Events API.
    ///
    /// # Arguments
    /// * `region` - The region to connect to.
    /// * `auth_context` - The authentication context to use.
    pub fn new(region: CoralogixRegion, auth_context: AuthContext) -> Result<Self> {
        let channel = Endpoint::from_str(&region.grpc_endpoint())?
            .tls_config(ClientTlsConfig::new().with_native_roots())?
            .connect_lazy();

        let request_metadata: CallProperties = (&auth_context.team_level_api_key).into();
        Ok(Self {
            metadata_map: request_metadata.to_metadata_map(),
            service_client: Mutex::new(EventsServiceClient::new(channel)),
        })
    }

    /// Gets a single event by ID.
    ///
    /// # Arguments
    /// * `id` - The ID of the event to get.
    /// * `order_bys` - The order by fields to use.
    /// * `pagination` - The pagination request to use.
    pub async fn get(
        &self,
        id: Option<String>,
        order_bys: Vec<OrderBy>,
        pagination: Option<PaginationRequest>,
    ) -> Result<GetEventResponse> {
        let request = make_request_with_metadata(
            GetEventRequest {
                id,
                order_bys,
                pagination,
            },
            &self.metadata_map,
        );

        Ok(self
            .service_client
            .lock()
            .await
            .get_event(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.events.v3.EventsService/GetEvent".into(),
                feature_group: EVENTS_FEATURE_GROUP_ID.into(),
            })?
            .into_inner())
    }

    /// Gets multiple events by their IDs.
    ///
    /// # Arguments
    /// * `ids` - The IDs of the events to get.
    /// * `order_bys` - The order by fields to use.
    /// * `pagination` - The pagination request to use.
    /// * `filter` - The filter to use.
    pub async fn batch_get(
        &self,
        ids: Vec<String>,
        order_bys: Vec<OrderBy>,
        pagination: Option<PaginationRequest>,
        filter: Option<EventsQueryFilter>,
    ) -> Result<BatchGetEventResponse> {
        let request = make_request_with_metadata(
            BatchGetEventRequest {
                ids,
                order_bys,
                pagination,
                filter,
            },
            &self.metadata_map,
        );

        Ok(self
            .service_client
            .lock()
            .await
            .batch_get_event(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.events.v3.EventsService/BatchGetEvent".into(),
                feature_group: EVENTS_FEATURE_GROUP_ID.into(),
            })?
            .into_inner())
    }

    /// Lists events based on the provided filter criteria.
    ///
    /// # Arguments
    /// * `filter` - The filter to use.
    /// * `order_bys` - The order by fields to use.
    /// * `pagination` - The pagination request to use.
    pub async fn list(
        &self,
        filter: Option<EventsFilter>,
        order_bys: Vec<OrderBy>,
        pagination: Option<PaginationRequest>,
    ) -> Result<ListEventsResponse> {
        let request = make_request_with_metadata(
            ListEventsRequest {
                filter,
                order_bys,
                pagination,
            },
            &self.metadata_map,
        );

        Ok(self
            .service_client
            .lock()
            .await
            .list_events(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.events.v3.EventsService/ListEvents".into(),
                feature_group: EVENTS_FEATURE_GROUP_ID.into(),
            })?
            .into_inner())
    }

    /// Gets the count of events matching the filter criteria.
    ///
    /// # Arguments
    /// * `filter` - The filter to use.
    pub async fn list_count(
        &self,
        filter: Option<EventsFilter>,
    ) -> Result<ListEventsCountResponse> {
        let request =
            make_request_with_metadata(ListEventsCountRequest { filter }, &self.metadata_map);

        Ok(self
            .service_client
            .lock()
            .await
            .list_events_count(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.events.v3.EventsService/ListEventsCount".into(),
                feature_group: EVENTS_FEATURE_GROUP_ID.into(),
            })?
            .into_inner())
    }

    /// Gets statistics about events matching the filter criteria.
    ///
    /// # Arguments
    /// * `filter` - The filter to use.
    pub async fn get_statistics(
        &self,
        filter: Option<EventsFilter>,
    ) -> Result<GetEventsStatisticsResponse> {
        let request =
            make_request_with_metadata(GetEventsStatisticsRequest { filter }, &self.metadata_map);

        Ok(self
            .service_client
            .lock()
            .await
            .get_events_statistics(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.events.v3.EventsService/GetEventsStatistics".into(),
                feature_group: EVENTS_FEATURE_GROUP_ID.into(),
            })?
            .into_inner())
    }
}
