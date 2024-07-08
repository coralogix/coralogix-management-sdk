use std::str::FromStr;

use crate::{
    auth::{ApiKey, AuthData},
    error::Result,
    util::make_request_with_metadata,
};

use cx_api::proto::com::coralogixapis::events2metrics::v2::{
    events2_metric_service_client::Events2MetricServiceClient, AtomicBatchExecuteE2mRequest,
    CreateE2mRequest, CreateE2mResponse, DeleteE2mRequest, DeleteE2mResponse, GetE2mRequest,
    GetE2mResponse, GetLimitsRequest, GetLimitsResponse, ListE2mRequest, ListE2mResponse,
    ListLabelsCardinalityRequest, ListLabelsCardinalityResponse, ReplaceE2mRequest,
    ReplaceE2mResponse,
};
use tokio::sync::Mutex;
use tonic::{
    metadata::MetadataMap,
    transport::{Channel, Endpoint},
};

pub use cx_api::proto::com::coralogixapis::events2metrics::v2::{
    list_labels_cardinality_request::Query, E2m, E2mCreateParams, MetricField, MetricLabel,
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
    /// * `api_key` - The [`ApiKey`] to use for authentication.
    /// * `region` - The [`CoralogixRegion`] to connect to.
    pub fn new(api_key: ApiKey, region: CoralogixRegion) -> Result<Self> {
        let channel: Channel = Endpoint::from_str(region.endpoint().as_str())?.connect_lazy();
        let auth_data: AuthData = (&api_key).into();
        Ok(Self {
            metadata_map: auth_data.to_metadata_map(),
            service_client: Mutex::new(Events2MetricServiceClient::new(channel)),
        })
    }

    pub async fn create(&self, params: E2mCreateParams) -> Result<CreateE2mResponse> {
        let request =
            make_request_with_metadata(CreateE2mRequest { e2m: Some(params) }, &self.metadata_map);
        Ok(self
            .service_client
            .lock()
            .await
            .create_e2m(request)
            .await?
            .into_inner())
    }

    pub async fn replace(&self, e2m: E2m) -> Result<ReplaceE2mResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .replace_e2m(make_request_with_metadata(
                ReplaceE2mRequest { e2m: Some(e2m) },
                &self.metadata_map,
            ))
            .await?
            .into_inner())
    }

    pub async fn delete(&self, id: String) -> Result<DeleteE2mResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .delete_e2m(make_request_with_metadata(
                DeleteE2mRequest { id: Some(id) },
                &self.metadata_map,
            ))
            .await?
            .into_inner())
    }

    pub async fn get(&self, id: String) -> Result<GetE2mResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .get_e2m(make_request_with_metadata(
                GetE2mRequest { id: Some(id) },
                &self.metadata_map,
            ))
            .await?
            .into_inner())
    }

    /// Retrieves the limits associated with this account.
    pub async fn get_limits(&self) -> Result<GetLimitsResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .get_limits(make_request_with_metadata(
                GetLimitsRequest {},
                &self.metadata_map,
            ))
            .await?
            .into_inner())
    }

    pub async fn list(&self) -> Result<ListE2mResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .list_e2m(make_request_with_metadata(
                ListE2mRequest {},
                &self.metadata_map,
            ))
            .await?
            .into_inner())
    }

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
            .await?
            .into_inner())
    }

    // TODO: Expose the AtomicBatchExecuteE2m API where CRUD requests can be batched together.
}
