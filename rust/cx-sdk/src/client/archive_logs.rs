pub use cx_api::proto::com::coralogix::archive::v2::set_target_request::TargetSpec;
use cx_api::proto::com::coralogix::archive::v2::target_service_client::TargetServiceClient;
use cx_api::proto::com::coralogix::archive::v2::{
    GetTargetRequest, GetTargetResponse, SetTargetRequest, SetTargetResponse, ValidateTargetRequest,
};
use std::str::FromStr;
use tokio::sync::Mutex;
use tonic::{
    metadata::MetadataMap,
    transport::{Channel, Endpoint},
};

pub use cx_api::proto::com::coralogix::archive::v2::validate_target_request::TargetSpec as TargetSpecValidation;
pub use cx_api::proto::com::coralogix::archive::v2::S3TargetSpec;

use crate::CoralogixRegion;
use crate::{
    auth::{ApiKey, AuthData},
    error::Result,
    util::make_request_with_metadata,
};

/// The logs archive API client.
/// Read more at [https://coralogix.com/docs/archive-s3-bucket-forever/]()
pub struct LogsArchiveClient {
    metadata_map: MetadataMap,
    service_client: Mutex<TargetServiceClient<Channel>>,
}

impl LogsArchiveClient {
    /// Creates a new client for the Logs Archive API.
    ///
    /// # Arguments
    /// * `api_key` - The  to use for authentication.
    /// * `region` - The region to connect to.
    pub fn new(region: CoralogixRegion, api_key: ApiKey) -> Self {
        let channel: Channel = Endpoint::from_str(&region.endpoint())
            .unwrap()
            .connect_lazy();
        let auth_data: AuthData = (&api_key).into();
        Self {
            metadata_map: auth_data.to_metadata_map(),
            service_client: Mutex::new(TargetServiceClient::new(channel)),
        }
    }

    /// Retrieves the current target configuration.
    pub async fn get_target(&self) -> Result<GetTargetResponse> {
        let request = make_request_with_metadata(GetTargetRequest {}, &self.metadata_map);
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .get_target(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    /// Sets the target configuration.
    ///
    /// # Arguments
    /// * `is_active` - Whether the target should be active.
    /// * `target_spec` - The target configuration.
    pub async fn set_target(
        &self,
        is_active: bool,
        target_spec: TargetSpec,
    ) -> Result<SetTargetResponse> {
        let request = make_request_with_metadata(
            SetTargetRequest {
                is_active,
                target_spec: Some(target_spec),
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .set_target(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    /// Validates a target configuration.
    ///
    /// # Arguments
    /// * `is_active` - Whether the target should be active.
    /// * `target_spec` - The target configuration to validate.
    pub async fn validate_target(
        &self,
        is_active: bool,
        target_spec: TargetSpecValidation,
    ) -> Result<()> {
        let request = make_request_with_metadata(
            ValidateTargetRequest {
                target_spec: Some(target_spec),
                is_active,
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .validate_target(request)
                .await
                .map(|_| ())
                .map_err(From::from)
        }
    }
}
