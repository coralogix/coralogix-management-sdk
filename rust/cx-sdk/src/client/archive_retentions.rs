use cx_api::proto::com::coralogix::archive::v1::{
    retentions_service_client::RetentionsServiceClient, ActivateRetentionsRequest,
    ActivateRetentionsResponse, GetRetentionsEnabledRequest, GetRetentionsEnabledResponse,
    GetRetentionsRequest, GetRetentionsResponse, UpdateRetentionsRequest, UpdateRetentionsResponse,
};
use std::str::FromStr;
use tokio::sync::Mutex;
use tonic::{
    metadata::MetadataMap,
    transport::{Channel, Endpoint},
};

use crate::{
    auth::{ApiKey, AuthData},
    error::Result,
    util::make_request_with_metadata,
    CoralogixRegion,
};

pub use cx_api::proto::com::coralogix::archive::v1::RetentionUpdateElement;

/// The Archive Retention API client.
/// Read more at <https://coralogix.com/docs/archive-retention/>
pub struct ArchiveRetentionClient {
    metadata_map: MetadataMap,
    service_client: Mutex<RetentionsServiceClient<Channel>>,
}

impl ArchiveRetentionClient {
    /// Creates a new client for the Archive Retention API.
    ///
    /// # Arguments
    /// * `api_key` - The [`ApiKey`] to use for authentication.
    /// * `region` - The [`CoralogixRegion`] to connect to.
    pub fn new(api_key: ApiKey, region: CoralogixRegion) -> Result<Self> {
        let channel: Channel = Endpoint::from_str(region.endpoint().as_str())?.connect_lazy();
        let auth_data: AuthData = (&api_key).into();
        Ok(Self {
            metadata_map: auth_data.to_metadata_map(),
            service_client: Mutex::new(RetentionsServiceClient::new(channel)),
        })
    }

    /// Retrieves the retention settings for a tenant.
    pub async fn get_retentions(&self) -> Result<GetRetentionsResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .get_retentions(make_request_with_metadata(
                GetRetentionsRequest {},
                &self.metadata_map,
            ))
            .await?
            .into_inner())
    }

    /// Retrieves the enabled retention settings for a tenant.
    pub async fn get_enabled_retentions(&self) -> Result<GetRetentionsEnabledResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .get_retentions_enabled(make_request_with_metadata(
                GetRetentionsEnabledRequest {},
                &self.metadata_map,
            ))
            .await?
            .into_inner())
    }

    /// Updates the retention settings for a tenant.
    pub async fn update(
        &self,
        retention_update_elements: Vec<RetentionUpdateElement>,
    ) -> Result<UpdateRetentionsResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .update_retentions(make_request_with_metadata(
                UpdateRetentionsRequest {
                    retention_update_elements,
                },
                &self.metadata_map,
            ))
            .await?
            .into_inner())
    }

    /// Activates all archive retentions
    pub async fn activate(&self) -> Result<ActivateRetentionsResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .activate_retentions(make_request_with_metadata(
                ActivateRetentionsRequest {},
                &self.metadata_map,
            ))
            .await?
            .into_inner())
    }
}
