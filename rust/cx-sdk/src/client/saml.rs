// Copyright 2024 Coralogix Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

use std::str::FromStr;

use cx_api::proto::com::coralogixapis::aaa::sso::v2::{
    GetConfigurationRequest,
    GetConfigurationResponse,
    GetSpParametersRequest,
    GetSpParametersResponse,
    SetActiveRequest,
    SetIdpParametersRequest,
    saml_configuration_service_client::SamlConfigurationServiceClient,
};

pub use cx_api::proto::com::coralogixapis::aaa::sso::v2::{
    IdpParameters,
    idp_parameters::Metadata,
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

use crate::{
    CoralogixRegion,
    auth::AuthContext,
    error::{
        Result,
        SdkApiError,
        SdkError,
    },
    metadata::CallProperties,
    util::make_request_with_metadata,
};

const SAML_FEATURE_GROUP_ID: &str = "aaa";
/// A client for the SAML Configuration API.
/// Read more at <https://coralogix.com/docs/slo-management-api/>
pub struct SamlClient {
    metadata_map: MetadataMap,
    service_client: Mutex<SamlConfigurationServiceClient<Channel>>,
}

impl SamlClient {
    /// Creates a new client for the SAML Configuration.
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
            service_client: Mutex::new(SamlConfigurationServiceClient::new(channel)),
        })
    }

    /// Retrieves the SP parameters for a given team.
    ///
    /// # Arguments
    /// * `team_id` - The ID of the team to retrieve parameters for.
    pub async fn get_sp_parameters(&self, team_id: u32) -> Result<GetSpParametersResponse> {
        let request =
            make_request_with_metadata(GetSpParametersRequest { team_id }, &self.metadata_map);
        self.service_client
            .lock()
            .await
            .get_sp_parameters(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogix.aaa.sso.v2.SamlConfigurationService/GetSpParameters"
                        .to_string(),
                    feature_group: SAML_FEATURE_GROUP_ID.into(),
                })
            })
    }

    /// Sets the IDP parameters for a given team.
    ///
    /// # Arguments
    /// * `team_id` - The ID of the team to set parameters for.
    /// * `idp_parameters` - The [`IdpParameters`] to set.
    pub async fn set_idp_parameters(
        &self,
        team_id: u32,
        idp_parameters: Option<IdpParameters>,
    ) -> Result<()> {
        let request = make_request_with_metadata(
            SetIdpParametersRequest {
                team_id,
                params: idp_parameters,
            },
            &self.metadata_map,
        );
        self.service_client
            .lock()
            .await
            .set_idp_parameters(request)
            .await
            .map(|_| ())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogix.aaa.sso.v2.SamlConfigurationService/SetIdpParameters"
                        .to_string(),
                    feature_group: SAML_FEATURE_GROUP_ID.into(),
                })
            })
    }

    /// Sets the active status for a given team.
    ///
    /// # Arguments
    /// * `team_id` - The ID of the team to set the active status for.
    /// * `is_active` - The active status to set.
    pub async fn set_active(&self, team_id: u32, is_active: bool) -> Result<()> {
        let request =
            make_request_with_metadata(SetActiveRequest { team_id, is_active }, &self.metadata_map);
        self.service_client
            .lock()
            .await
            .set_active(request)
            .await
            .map(|_| ())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogix.aaa.sso.v2.SamlConfigurationService/SetActive"
                        .to_string(),
                    feature_group: SAML_FEATURE_GROUP_ID.into(),
                })
            })
    }

    /// Retrieves the configuration for a given team.
    ///
    /// # Arguments
    /// * `team_id` - The ID of the team to retrieve the configuration for.
    pub async fn get_configuration(&self, team_id: u32) -> Result<GetConfigurationResponse> {
        let request =
            make_request_with_metadata(GetConfigurationRequest { team_id }, &self.metadata_map);

        self.service_client
            .lock()
            .await
            .get_configuration(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogix.aaa.sso.v2.SamlConfigurationService/GetConfiguration"
                        .to_string(),
                    feature_group: SAML_FEATURE_GROUP_ID.into(),
                })
            })
    }
}
