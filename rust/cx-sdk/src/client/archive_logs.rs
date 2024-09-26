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

pub use cx_api::proto::com::coralogix::archive::v2::set_target_request::TargetSpec;
use cx_api::proto::com::coralogix::archive::v2::target_service_client::TargetServiceClient;
use cx_api::proto::com::coralogix::archive::v2::{
    GetTargetRequest,
    GetTargetResponse,
    SetTargetRequest,
    SetTargetResponse,
    ValidateTargetRequest,
};
use std::str::FromStr;
use tokio::sync::Mutex;
use tonic::transport::ClientTlsConfig;
use tonic::{
    metadata::MetadataMap,
    transport::{
        Channel,
        Endpoint,
    },
};

pub use cx_api::proto::com::coralogix::archive::v2::validate_target_request::TargetSpec as TargetSpecValidation;
pub use cx_api::proto::com::coralogix::archive::v2::S3TargetSpec;

use crate::auth::AuthContext;
use crate::CoralogixRegion;
use crate::{
    error::Result,
    metadata::CallProperties,
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
    /// * `auth_context` - The  to use for authentication.
    /// * `region` - The region to connect to.
    pub fn new(region: CoralogixRegion, auth_context: AuthContext) -> Result<Self> {
        let channel: Channel = Endpoint::from_str(&region.grpc_endpoint())?
            .tls_config(ClientTlsConfig::new().with_native_roots())?
            .connect_lazy();
        let request_metadata: CallProperties = (&auth_context.team_level_api_key).into();
        Ok(Self {
            metadata_map: request_metadata.to_metadata_map(),
            service_client: Mutex::new(TargetServiceClient::new(channel)),
        })
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
