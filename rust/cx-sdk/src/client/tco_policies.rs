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

use cx_api::proto::com::coralogix::quota::v1::{
    CreatePolicyRequest,
    CreatePolicyResponse,
    DeletePolicyRequest,
    DeletePolicyResponse,
    GetCompanyPoliciesRequest,
    GetCompanyPoliciesResponse,
    GetPolicyRequest,
    GetPolicyResponse,
    UpdatePolicyRequest,
    UpdatePolicyResponse,
    policies_service_client::PoliciesServiceClient,
    update_policy_request,
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

pub use cx_api::proto::com::coralogix::quota::v1::{
    ArchiveRetention,
    LogRules,
    Rule,
    RuleTypeId,
    SourceType,
    SpanRules,
    create_policy_request::SourceTypeRules,
};

fn convert_source_types(a: SourceTypeRules) -> update_policy_request::SourceTypeRules {
    match a {
        SourceTypeRules::LogRules(r) => update_policy_request::SourceTypeRules::LogRules(r),
        SourceTypeRules::SpanRules(r) => update_policy_request::SourceTypeRules::SpanRules(r),
    }
}

use crate::CoralogixRegion;

/// The TCO client.
/// Read more at <https://coralogix.com/docs/tco-tracing-policy-grpc-api/>
pub struct TcoPoliciesClient {
    metadata_map: MetadataMap,
    service_client: Mutex<PoliciesServiceClient<Channel>>,
}

impl TcoPoliciesClient {
    /// Creates a new client for the TCO.
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
            service_client: Mutex::new(PoliciesServiceClient::new(channel)),
        })
    }

    /// Creates a new TCO policy.
    ///
    /// # Arguments
    /// * `name` - The name of the policy.
    /// * `description` - The description of the policy.
    /// * `priority` - The priority of the policy.
    /// * `application_rule` - The application [`Rule`] of the policy.
    /// * `subsystem_rule` - The subsystem [`Rule`] of the policy.
    /// * `archive_retention` - The [`ArchiveRetention`] of the policy.
    /// * `source_type_rules` - The [`SourceTypeRules`] of the policy.
    #[allow(clippy::too_many_arguments)]
    pub async fn create(
        &self,
        name: Option<String>,
        description: Option<String>,
        priority: i32,
        application_rule: Option<Rule>,
        subsystem_rule: Option<Rule>,
        archive_retention: Option<ArchiveRetention>,
        source_type_rules: Option<SourceTypeRules>,
    ) -> Result<CreatePolicyResponse> {
        let request = make_request_with_metadata(
            CreatePolicyRequest {
                name,
                description,
                priority,
                application_rule,
                subsystem_rule,
                archive_retention,
                source_type_rules,
            },
            &self.metadata_map,
        );
        self.service_client
            .lock()
            .await
            .create_policy(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogixapis.quota.v1.PoliciesService/CreatePolicy".into(),
                })
            })
    }

    /// Creates a new TCO policy.
    ///
    /// # Arguments
    /// * `id` - The id of the policy.
    /// * `name` - The name of the policy.
    /// * `description` - The description of the policy.
    /// * `priority` - The priority of the policy.
    /// * `application_rule` - The application [`Rule`] of the policy.
    /// * `subsystem_rule` - The subsystem [`Rule`] of the policy.
    /// * `archive_retention` - The [`ArchiveRetention`] of the policy.
    /// * `source_type_rules` - The [`SourceTypeRules`] of the policy.
    #[allow(clippy::too_many_arguments)]
    pub async fn update(
        &self,
        id: String,
        name: Option<String>,
        description: Option<String>,
        priority: i32,
        application_rule: Option<Rule>,
        subsystem_rule: Option<Rule>,
        archive_retention: Option<ArchiveRetention>,
        enabled: Option<bool>,
        source_type_rules: Option<SourceTypeRules>,
    ) -> Result<UpdatePolicyResponse> {
        let request = make_request_with_metadata(
            UpdatePolicyRequest {
                id: Some(id),
                name,
                description,
                priority,
                application_rule,
                subsystem_rule,
                archive_retention,
                enabled,
                source_type_rules: source_type_rules.map(convert_source_types),
            },
            &self.metadata_map,
        );
        self.service_client
            .lock()
            .await
            .update_policy(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogixapis.quota.v1.PoliciesService/UpdatePolicy".into(),
                })
            })
    }

    /// Deletes a TCO policy.
    ///
    /// # Arguments
    /// * `id` - The id of the policy to delete.
    pub async fn delete(&self, id: String) -> Result<DeletePolicyResponse> {
        let request =
            make_request_with_metadata(DeletePolicyRequest { id: Some(id) }, &self.metadata_map);
        self.service_client
            .lock()
            .await
            .delete_policy(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogixapis.quota.v1.PoliciesService/DeletePolicy".into(),
                })
            })
    }

    /// Retrieves a TCO policy.
    ///
    /// # Arguments
    /// * `id` - The id of the policy to retrieve.
    pub async fn get(&self, id: String) -> Result<GetPolicyResponse> {
        let request =
            make_request_with_metadata(GetPolicyRequest { id: Some(id) }, &self.metadata_map);

        self.service_client
            .lock()
            .await
            .get_policy(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogixapis.quota.v1.PoliciesService/GetPolicy".into(),
                })
            })
    }

    /// Retrieves a list of TCO policies.
    ///
    /// # Arguments
    /// * `source_type` - The [`SourceType`] of the policies to retrieve.
    /// * `enabled_only` - Whether to retrieve only enabled policies.
    pub async fn list(
        &self,
        source_type: SourceType,
        enabled_only: bool,
    ) -> Result<GetCompanyPoliciesResponse> {
        let request = make_request_with_metadata(
            GetCompanyPoliciesRequest {
                enabled_only: Some(enabled_only),
                source_type: Some(source_type.into()),
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .get_company_policies(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogixapis.quota.v1.PoliciesService/GetCompanyPolicies"
                        .into(),
                })
            })
    }
}
