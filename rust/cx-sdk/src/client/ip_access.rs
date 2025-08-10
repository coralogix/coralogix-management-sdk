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

use cx_api::proto::com::coralogixapis::aaa::v1::{
    CreateCompanyIpAccessSettingsRequest,
    CreateCompanyIpAccessSettingsResponse,
    DeleteCompanyIpAccessSettingsRequest,
    DeleteCompanyIpAccessSettingsResponse,
    GetCompanyIpAccessSettingsRequest,
    GetCompanyIpAccessSettingsResponse,
    ReplaceCompanyIpAccessSettingsRequest,
    ReplaceCompanyIpAccessSettingsResponse,
    ip_access_service_client::IpAccessServiceClient,
};

use std::str::FromStr;

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

use tokio::sync::Mutex;
use tonic::{
    metadata::MetadataMap,
    transport::{
        Channel,
        ClientTlsConfig,
        Endpoint,
    },
};

pub use crate::com::coralogixapis::aaa::v1::{
    CompanyIpAccessSettings,
    CoralogixCustomerSupportAccess,
    IpAccess,
};

const IP_ACCESS_FEATURE_GROUP_ID: &str = "aaa";

/// IP Access API client.
/// Read more at <https://coralogix.com/docs/user-guides/account-management/account-settings/ip-access-control/>
pub struct IpAccessClient {
    metadata_map: MetadataMap,
    service_client: Mutex<IpAccessServiceClient<Channel>>,
}

impl IpAccessClient {
    /// Creates a new client for the IP Access API.
    ///
    /// # Arguments
    /// * `auth_context` - The [`AuthContext`] to use for authentication (team-level).
    /// * `region` - The [`CoralogixRegion`] to connect to.
    pub fn new(auth_context: AuthContext, region: CoralogixRegion) -> Result<Self> {
        let channel: Channel = Endpoint::from_str(region.grpc_endpoint().as_str())?
            .tls_config(ClientTlsConfig::new().with_native_roots())?
            .connect_lazy();

        // Use team-level API key semantics (same as ApiKeys client).
        let request_metadata: CallProperties = (&auth_context.team_level_api_key).into();

        Ok(Self {
            metadata_map: request_metadata.to_metadata_map(),
            service_client: Mutex::new(IpAccessServiceClient::new(channel)),
        })
    }

    /// Creates company IP access settings.
    ///
    /// # Arguments
    /// * `ip_access` - A vector of [`IpAccess`] settings to apply.
    /// * `support_access` - Optional [`CoralogixCustomerSupportAccess`] setting to enable or disable access for Coralogix customer support.
    pub async fn create(
        &self,
        ip_access: Vec<IpAccess>,
        support_access: CoralogixCustomerSupportAccess,
    ) -> Result<CreateCompanyIpAccessSettingsResponse> {
        let req = CreateCompanyIpAccessSettingsRequest {
            ip_access,
            enable_coralogix_customer_support_access: support_access.into(),
        };

        let request = make_request_with_metadata(req, &self.metadata_map);

        self.service_client
            .lock()
            .await
            .create_company_ip_access_settings(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint:
                        "/com.coralogixapis.aaa.v1.IpAccessService/CreateCompanyIpAccessSettings"
                            .into(),
                    feature_group: IP_ACCESS_FEATURE_GROUP_ID.into(),
                })
            })
    }

    /// Replaces company IP access settings.
    ///
    /// # Arguments
    /// * `id` - The optional ID of the settings to replace.
    /// * `ip_access` - A vector of [`IpAccess`] settings to apply.
    /// * `support_access` - Optional [`CoralogixCustomerSupportAccess`] setting to enable or disable access for Coralogix customer support.
    pub async fn replace(
        &self,
        id: Option<String>,
        ip_access: Vec<IpAccess>,
        support_access: CoralogixCustomerSupportAccess,
    ) -> Result<ReplaceCompanyIpAccessSettingsResponse> {
        let req = ReplaceCompanyIpAccessSettingsRequest {
            id,
            ip_access,
            enable_coralogix_customer_support_access: support_access.into(),
        };

        let request = make_request_with_metadata(req, &self.metadata_map);

        self.service_client
            .lock()
            .await
            .replace_company_ip_access_settings(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint:
                        "/com.coralogixapis.aaa.v1.IpAccessService/ReplaceCompanyIpAccessSettings"
                            .into(),
                    feature_group: IP_ACCESS_FEATURE_GROUP_ID.into(),
                })
            })
    }

    /// Retrieves company IP access settings.
    ///
    /// # Arguments
    /// * `id` - The optional ID of the settings to retrieve.
    pub async fn get(&self, id: Option<String>) -> Result<GetCompanyIpAccessSettingsResponse> {
        let req = GetCompanyIpAccessSettingsRequest { id };

        let request = make_request_with_metadata(req, &self.metadata_map);

        self.service_client
            .lock()
            .await
            .get_company_ip_access_settings(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint:
                        "/com.coralogixapis.aaa.v1.IpAccessService/GetCompanyIpAccessSettings"
                            .into(),
                    feature_group: IP_ACCESS_FEATURE_GROUP_ID.into(),
                })
            })
    }

    /// Deletes company IP access settings.
    ///
    /// # Arguments
    /// * `id` - The optional ID of the settings to delete.
    pub async fn delete(
        &self,
        id: Option<String>,
    ) -> Result<DeleteCompanyIpAccessSettingsResponse> {
        let req = DeleteCompanyIpAccessSettingsRequest { id };

        let request = make_request_with_metadata(req, &self.metadata_map);

        self.service_client
            .lock()
            .await
            .delete_company_ip_access_settings(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint:
                        "/com.coralogixapis.aaa.v1.IpAccessService/DeleteCompanyIpAccessSettings"
                            .into(),
                    feature_group: IP_ACCESS_FEATURE_GROUP_ID.into(),
                })
            })
    }
}
