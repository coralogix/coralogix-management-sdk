use std::str::FromStr;

use crate::{
    auth::{ApiKey, AuthData},
    error::Result,
    util::make_request_with_metadata,
};

use cx_api::proto::com::coralogix::catalog::v1::{
    sli_service_client::SliServiceClient, CreateSliRequest, CreateSliResponse, CreateSlisRequest,
    CreateSlisResponse, DeleteSliRequest, DeleteSliResponse, GetSliE2mQueryRequest,
    GetSliE2mQueryResponse, GetSliStatusHistoryRequest, GetSliStatusHistoryResponse,
    GetSlisRequest, GetSlisResponse, UpdateSliRequest, UpdateSliResponse, UpdateSlisRequest,
    UpdateSlisResponse,
};
pub use cx_api::proto::com::coralogix::catalog::v1::{Sli, SliFilter};
use time::OffsetDateTime;
use tokio::sync::Mutex;
use tonic::{
    metadata::MetadataMap,
    transport::{Channel, Endpoint},
};

use crate::CoralogixRegion;

/// The Service Line Indicator (SLI) client.
/// Read more at <https://coralogix.com/docs/sli/>
pub struct SliClient {
    metadata_map: MetadataMap,
    service_client: Mutex<SliServiceClient<Channel>>,
}

impl SliClient {
    /// Creates a new client for the SLI.
    ///
    /// # Arguments
    /// * `api_key` - The [`ApiKey`] to use for authentication.
    /// * `region` - The [`CoralogixRegion`] to connect to.
    pub fn new(api_key: ApiKey, region: CoralogixRegion) -> Result<Self> {
        let channel: Channel = Endpoint::from_str(region.endpoint().as_str())?.connect_lazy();
        let auth_data: AuthData = (&api_key).into();
        Ok(Self {
            metadata_map: auth_data.to_metadata_map(),
            service_client: Mutex::new(SliServiceClient::new(channel)),
        })
    }

    /// Creates a new SLI.
    ///
    /// # Arguments
    /// * `sli` - The [`Sli`] to create.
    pub async fn create(&self, sli: Sli) -> Result<CreateSliResponse> {
        let request =
            make_request_with_metadata(CreateSliRequest { sli: Some(sli) }, &self.metadata_map);
        self.service_client
            .lock()
            .await
            .create_sli(request)
            .await
            .map(|r| r.into_inner())
            .map_err(From::from)
    }

    /// Creates a new SLI in bulk.
    ///
    /// # Arguments
    /// * `sli` - The [`Sli`]s to create.
    pub async fn create_bulk(&self, slis: Vec<Sli>) -> Result<CreateSlisResponse> {
        let request = make_request_with_metadata(CreateSlisRequest { slis }, &self.metadata_map);
        self.service_client
            .lock()
            .await
            .create_slis(request)
            .await
            .map(|r| r.into_inner())
            .map_err(From::from)
    }

    /// Updates an existing SLI.
    ///
    /// # Arguments
    /// * `sli` - The [`Sli`] to update.    
    pub async fn update(&self, sli: Sli) -> Result<UpdateSliResponse> {
        let request =
            make_request_with_metadata(UpdateSliRequest { sli: Some(sli) }, &self.metadata_map);
        self.service_client
            .lock()
            .await
            .update_sli(request)
            .await
            .map(|r| r.into_inner())
            .map_err(From::from)
    }

    /// Updates an existing SLIs in bulk.
    ///
    /// # Arguments
    /// * `slis` - The [`Sli`]s to update.    
    pub async fn update_bulk(&self, slis: Vec<Sli>) -> Result<UpdateSlisResponse> {
        let request = make_request_with_metadata(UpdateSlisRequest { slis }, &self.metadata_map);
        self.service_client
            .lock()
            .await
            .update_slis(request)
            .await
            .map(|r| r.into_inner())
            .map_err(From::from)
    }

    /// Deletes an SLI.
    ///
    /// # Arguments
    /// * `id` - The id of the SLI to delete.
    pub async fn delete(&self, id: String) -> Result<DeleteSliResponse> {
        let request =
            make_request_with_metadata(DeleteSliRequest { sli_id: Some(id) }, &self.metadata_map);
        self.service_client
            .lock()
            .await
            .delete_sli(request)
            .await
            .map(|r| r.into_inner())
            .map_err(From::from)
    }

    /// Retrieves a list of SLIs by service name.
    ///
    /// # Arguments
    /// * `service_name` - The name of the service to retrieve SLIs for.
    pub async fn get_by_service_name(&self, service_name: String) -> Result<GetSlisResponse> {
        let request = make_request_with_metadata(
            GetSlisRequest {
                service_name: Some(service_name),
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .get_slis(request)
            .await
            .map(|r| r.into_inner())
            .map_err(From::from)
    }

    /// Get the SLI status history.
    ///
    /// # Arguments
    /// * `sli_id` - The ID of the SLI to retrieve the status history for.
    /// * `start` - The start time of the history.
    /// * `end` - The end time of the history.
    pub async fn get_sli_status_history(
        &self,
        sli_id: String,
        start: Option<OffsetDateTime>,
        end: Option<OffsetDateTime>,
    ) -> Result<GetSliStatusHistoryResponse> {
        let request = make_request_with_metadata(
            GetSliStatusHistoryRequest {
                sli_id: Some(sli_id),
                time_range: Some(cx_api::proto::com::coralogix::catalog::v1::TimeRange {
                    start: start.map(|f| f.unix_timestamp()),
                    end: end.map(|f| f.unix_timestamp()),
                }),
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .get_sli_status_history(request)
            .await
            .map(|r| r.into_inner())
            .map_err(From::from)
    }

    /// Get the SLI events2metrics query.
    ///
    /// # Arguments
    /// * `sli_id` - The ID of the SLI to retrieve the status history for.
    pub async fn get_sli_e2m_query(&self, sli_id: String) -> Result<GetSliE2mQueryResponse> {
        let request = make_request_with_metadata(
            GetSliE2mQueryRequest {
                sli_id: Some(sli_id),
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .get_sli_e2m_query(request)
            .await
            .map(|r| r.into_inner())
            .map_err(From::from)
    }
}
