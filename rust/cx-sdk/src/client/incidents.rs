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

use cx_api::proto::com::coralogixapis::incidents::v1::{
    AcknowledgeIncidentsRequest,
    AcknowledgeIncidentsResponse,
    AssignIncidentsRequest,
    AssignIncidentsResponse,
    BatchGetIncidentRequest,
    BatchGetIncidentResponse,
    CloseIncidentsRequest,
    CloseIncidentsResponse,
    GetFilterValuesRequest,
    GetFilterValuesResponse,
    GetIncidentEventsRequest,
    GetIncidentEventsResponse,
    GetIncidentRequest,
    GetIncidentResponse,
    GetIncidentUsingCorrelationKeyRequest,
    GetIncidentUsingCorrelationKeyResponse,
    ListIncidentAggregationsRequest,
    ListIncidentAggregationsResponse,
    ListIncidentEventsFilterValuesRequest,
    ListIncidentEventsFilterValuesResponse,
    ListIncidentEventsRequest,
    ListIncidentEventsResponse,
    ListIncidentEventsTotalCountRequest,
    ListIncidentEventsTotalCountResponse,
    ListIncidentsRequest,
    ResolveIncidentsRequest,
    ResolveIncidentsResponse,
    UnassignIncidentsRequest,
    UnassignIncidentsResponse,
    incidents_service_client::IncidentsServiceClient,
};
pub use cx_api::proto::com::coralogixapis::incidents::v1::{
    ContextualLabelValues,
    FilterOperator,
    GroupBy,
    IncidentEventOrderByFieldType,
    IncidentEventQueryFilter,
    IncidentFields,
    IncidentQueryFilter,
    IncidentSearchQuery,
    IncidentSeverity,
    IncidentStatus,
    LabelsFilter,
    ListIncidentEventRequestOrderBy,
    ListIncidentsResponse,
    MetaLabel,
    OrderBy,
    OrderByDirection,
    OrderByFields,
    PaginationRequest,
    TimeRange,
    UserDetails,
    incident_search_query::Field,
    order_by::Field as OrderByField,
};
use time::OffsetDateTime;
use tokio::sync::Mutex;
use tonic::{
    metadata::MetadataMap,
    transport::{
        Channel,
        ClientTlsConfig,
        Endpoint,
    },
};

use crate::{
    CoralogixRegion,
    auth::AuthContext,
    error::SdkApiError,
    metadata::CallProperties,
    util::make_request_with_metadata,
};

const INCIDENTS_FEATURE_GROUP_ID: &str = "incidents";

/// IncidentsClient is a client for the Incidents API.
pub struct IncidentsClient {
    service_client: Mutex<IncidentsServiceClient<Channel>>,
    metadata_map: MetadataMap,
}

impl IncidentsClient {
    /// Creates a new IncidentsClient.
    /// # Arguments
    /// * `auth_context` - The [`AuthContext`] to use for authentication.
    /// * `region` - The [`CoralogixRegion`] to connect to.
    pub async fn new(
        auth_context: AuthContext,
        region: CoralogixRegion,
    ) -> Result<Self, Box<dyn std::error::Error>> {
        let channel: Channel = Endpoint::from_str(region.grpc_endpoint().as_str())?
            .tls_config(ClientTlsConfig::new().with_native_roots())?
            .connect_lazy();

        let request_metadata: CallProperties = (&auth_context.team_level_api_key).into();
        Ok(Self {
            metadata_map: request_metadata.to_metadata_map(),
            service_client: Mutex::new(IncidentsServiceClient::new(channel)),
        })
    }

    /// Gets an incident by its ID.
    /// # Arguments
    /// * `id` - The ID of the incident to get.
    pub async fn get_incident(
        &self,
        id: &str,
    ) -> Result<GetIncidentResponse, Box<dyn std::error::Error>> {
        let request = make_request_with_metadata(
            GetIncidentRequest {
                id: Some(id.to_string()),
            },
            &self.metadata_map,
        );
        let response = self
            .service_client
            .lock()
            .await
            .get_incident(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.incidents.v1.IncidentsService/GetIncident".into(),
                feature_group: INCIDENTS_FEATURE_GROUP_ID.into(),
            })?;
        Ok(response.into_inner())
    }

    /// Batch gets incidents by their IDs.
    /// # Arguments
    /// * `ids` - The IDs of the incidents to get.
    pub async fn batch_get_incident(
        &self,
        ids: Vec<String>,
    ) -> Result<BatchGetIncidentResponse, Box<dyn std::error::Error>> {
        let request =
            make_request_with_metadata(BatchGetIncidentRequest { ids }, &self.metadata_map);
        let response = self
            .service_client
            .lock()
            .await
            .batch_get_incident(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.incidents.v1.IncidentsService/BatchGetIncident"
                    .into(),
                feature_group: INCIDENTS_FEATURE_GROUP_ID.into(),
            })?;
        Ok(response.into_inner())
    }

    /// Lists incidents.
    /// # Arguments
    /// * `filter` - The [`IncidentQueryFilter`] to filter the incidents by.
    /// * `pagination` - The [`PaginationRequest`] to paginate the incidents.
    /// * `order_bys` - The [`OrderBy`]s to order the incidents by.
    pub async fn list_incidents(
        &self,
        filter: Option<IncidentQueryFilter>,
        pagination: Option<PaginationRequest>,
        order_bys: Vec<OrderBy>,
    ) -> Result<ListIncidentsResponse, Box<dyn std::error::Error>> {
        let request = make_request_with_metadata(
            ListIncidentsRequest {
                filter,
                pagination,
                order_bys,
            },
            &self.metadata_map,
        );

        let response = self
            .service_client
            .lock()
            .await
            .list_incidents(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.incidents.v1.IncidentsService/ListIncidents".into(),
                feature_group: INCIDENTS_FEATURE_GROUP_ID.into(),
            })?;
        Ok(response.into_inner())
    }

    /// Lists incident aggregations.
    /// # Arguments
    /// * `request` - The [`ListIncidentAggregationsRequest`] to list the incident aggregations.
    pub async fn list_incident_aggregations(
        &self,
        filter: Option<IncidentQueryFilter>,
        group_bys: Vec<GroupBy>,
        pagination: Option<PaginationRequest>,
    ) -> Result<ListIncidentAggregationsResponse, Box<dyn std::error::Error>> {
        let request = make_request_with_metadata(
            ListIncidentAggregationsRequest {
                filter,
                group_bys,
                pagination,
            },
            &self.metadata_map,
        );
        let response = self
            .service_client
            .lock()
            .await
            .list_incident_aggregations(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint:
                    "/com.coralogixapis.incidents.v1.IncidentsService/ListIncidentAggregations"
                        .into(),
                feature_group: INCIDENTS_FEATURE_GROUP_ID.into(),
            })?;
        Ok(response.into_inner())
    }

    /// Gets the filter values for a given filter.
    /// # Arguments
    /// * `filter` - The [`IncidentQueryFilter`] to get the filter values for.
    pub async fn get_filter_values(
        &self,
        filter: IncidentQueryFilter,
    ) -> Result<GetFilterValuesResponse, Box<dyn std::error::Error>> {
        let request = make_request_with_metadata(
            GetFilterValuesRequest {
                filter: Some(filter),
            },
            &self.metadata_map,
        );
        let response = self
            .service_client
            .lock()
            .await
            .get_filter_values(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.incidents.v1.IncidentsService/GetFilterValues".into(),
                feature_group: INCIDENTS_FEATURE_GROUP_ID.into(),
            })?;
        Ok(response.into_inner())
    }

    /// Assigns incidents to a user.
    /// # Arguments
    /// * `incident_ids` - The IDs of the incidents to assign.
    /// * `assigned_to` - The [`UserDetails`] to assign the incidents to.
    pub async fn assign_incidents(
        &self,
        incident_ids: Vec<String>,
        assigned_to: Option<UserDetails>,
    ) -> Result<AssignIncidentsResponse, Box<dyn std::error::Error>> {
        let request = make_request_with_metadata(
            AssignIncidentsRequest {
                incident_ids,
                assigned_to,
            },
            &self.metadata_map,
        );
        let response = self
            .service_client
            .lock()
            .await
            .assign_incidents(request)
            .await?;
        Ok(response.into_inner())
    }

    /// Unassigns incidents from a user.
    /// # Arguments
    /// * `incident_ids` - The IDs of the incidents to unassign.
    pub async fn unassign_incidents(
        &self,
        incident_ids: Vec<String>,
    ) -> Result<UnassignIncidentsResponse, Box<dyn std::error::Error>> {
        let request = make_request_with_metadata(
            UnassignIncidentsRequest { incident_ids },
            &self.metadata_map,
        );
        let response = self
            .service_client
            .lock()
            .await
            .unassign_incidents(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.incidents.v1.IncidentsService/UnassignIncidents"
                    .into(),
                feature_group: INCIDENTS_FEATURE_GROUP_ID.into(),
            })?;
        Ok(response.into_inner())
    }

    /// Acknowledges incidents.
    /// # Arguments
    /// * `incident_ids` - The IDs of the incidents to acknowledge.
    pub async fn acknowledge_incidents(
        &self,
        incident_ids: Vec<String>,
    ) -> Result<AcknowledgeIncidentsResponse, Box<dyn std::error::Error>> {
        let request = make_request_with_metadata(
            AcknowledgeIncidentsRequest { incident_ids },
            &self.metadata_map,
        );
        let response = self
            .service_client
            .lock()
            .await
            .acknowledge_incidents(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.incidents.v1.IncidentsService/AcknowledgeIncidents"
                    .into(),
                feature_group: INCIDENTS_FEATURE_GROUP_ID.into(),
            })?;
        Ok(response.into_inner())
    }

    /// Closes incidents.
    /// # Arguments
    /// * `incident_ids` - The IDs of the incidents to close.
    pub async fn close_incidents(
        &self,
        incident_ids: Vec<String>,
    ) -> Result<CloseIncidentsResponse, Box<dyn std::error::Error>> {
        let request =
            make_request_with_metadata(CloseIncidentsRequest { incident_ids }, &self.metadata_map);
        let response = self
            .service_client
            .lock()
            .await
            .close_incidents(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.incidents.v1.IncidentsService/CloseIncidents".into(),
                feature_group: INCIDENTS_FEATURE_GROUP_ID.into(),
            })?;
        Ok(response.into_inner())
    }

    /// Gets the events for an incident.
    /// # Arguments
    /// * `incident_id` - The ID of the incident to get the events for.
    pub async fn get_incident_events(
        &self,
        incident_id: Option<String>,
    ) -> Result<GetIncidentEventsResponse, Box<dyn std::error::Error>> {
        let request = make_request_with_metadata(
            GetIncidentEventsRequest { incident_id },
            &self.metadata_map,
        );
        let response = self
            .service_client
            .lock()
            .await
            .get_incident_events(request)
            .await?;
        Ok(response.into_inner())
    }

    /// Resolves incidents.
    /// # Arguments
    /// * `incident_ids` - The IDs of the incidents to resolve.
    pub async fn resolve_incidents(
        &self,
        incident_ids: Vec<String>,
    ) -> Result<ResolveIncidentsResponse, Box<dyn std::error::Error>> {
        let request = make_request_with_metadata(
            ResolveIncidentsRequest { incident_ids },
            &self.metadata_map,
        );
        let response = self
            .service_client
            .lock()
            .await
            .resolve_incidents(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.incidents.v1.IncidentsService/ResolveIncidents"
                    .into(),
                feature_group: INCIDENTS_FEATURE_GROUP_ID.into(),
            })?;
        Ok(response.into_inner())
    }

    /// Gets an incident by its correlation key and specific point in time.
    /// # Arguments
    /// * `correlation_key` - The correlation key of the incident to get.
    /// * `timestamp` - The timestamp to get the incident at.
    pub async fn get_incident_using_correlation_key(
        &self,
        correlation_key: String,
        timestamp: OffsetDateTime,
    ) -> Result<GetIncidentUsingCorrelationKeyResponse, Box<dyn std::error::Error>> {
        let request = make_request_with_metadata(
            GetIncidentUsingCorrelationKeyRequest {
                correlation_key: Some(correlation_key),
                incident_point_in_time: Some(timestamp.to_string().parse().unwrap()),
            },
            &self.metadata_map,
        );
        let response = self
            .service_client
            .lock()
            .await
            .get_incident_using_correlation_key(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint:
                    "/com.coralogixapis.incidents.v1.IncidentsService/GetIncidentUsingCorrelationKey"
                        .into(),
                feature_group: INCIDENTS_FEATURE_GROUP_ID.into(),
            })?;
        Ok(response.into_inner())
    }

    /// Lists the events for an incident.
    /// # Arguments
    /// * `incident_id` - The ID of the incident to get the events for.
    pub async fn list_incident_events(
        &self,
        filter: Option<IncidentEventQueryFilter>,
        pagination: Option<PaginationRequest>,
        order_by: Option<ListIncidentEventRequestOrderBy>,
    ) -> Result<ListIncidentEventsResponse, Box<dyn std::error::Error>> {
        let request = make_request_with_metadata(
            ListIncidentEventsRequest {
                filter,
                pagination,
                order_by,
            },
            &self.metadata_map,
        );
        let response = self
            .service_client
            .lock()
            .await
            .list_incident_events(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.incidents.v1.IncidentsService/ListIncidentEvents"
                    .into(),
                feature_group: INCIDENTS_FEATURE_GROUP_ID.into(),
            })?;
        Ok(response.into_inner())
    }

    /// Lists the total count of events for an incident.
    /// # Arguments
    /// * `incident_id` - The ID of the incident to get the total count of events for.
    pub async fn list_incident_events_total_count(
        &self,
        filter: Option<IncidentEventQueryFilter>,
    ) -> Result<ListIncidentEventsTotalCountResponse, Box<dyn std::error::Error>> {
        let request = make_request_with_metadata(
            ListIncidentEventsTotalCountRequest { filter },
            &self.metadata_map,
        );
        let response = self
            .service_client
            .lock()
            .await
            .list_incident_events_total_count(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint:
                    "/com.coralogixapis.incidents.v1.IncidentsService/ListIncidentEventsTotalCount"
                        .into(),
                feature_group: INCIDENTS_FEATURE_GROUP_ID.into(),
            })?;
        Ok(response.into_inner())
    }

    /// Lists the filter values for an incident event.
    /// # Arguments
    /// * `filter` - The [`IncidentEventQueryFilter`] to get the filter values for.
    pub async fn list_incident_events_filter_values(
        &self,
        filter: Option<IncidentEventQueryFilter>,
    ) -> Result<ListIncidentEventsFilterValuesResponse, Box<dyn std::error::Error>> {
        let request = make_request_with_metadata(
            ListIncidentEventsFilterValuesRequest { filter },
            &self.metadata_map,
        );
        let response = self
            .service_client
            .lock()
            .await
            .list_incident_events_filter_values(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint:
                    "/com.coralogixapis.incidents.v1.IncidentsService/ListIncidentEventsFilterValues"
                        .into(),
                feature_group: INCIDENTS_FEATURE_GROUP_ID.into(),
            })?;
        Ok(response.into_inner())
    }
}
