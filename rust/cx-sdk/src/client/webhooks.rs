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
    auth::{AuthContext, AuthData},
    error::Result,
    util::make_request_with_metadata,
};

use cx_api::proto::com::coralogix::outgoing_webhooks::v1::{
    outgoing_webhooks_service_client::OutgoingWebhooksServiceClient, CreateOutgoingWebhookRequest,
    CreateOutgoingWebhookResponse, DeleteOutgoingWebhookRequest, DeleteOutgoingWebhookResponse,
    GetOutgoingWebhookRequest, GetOutgoingWebhookResponse, GetOutgoingWebhookTypeDetailsRequest,
    GetOutgoingWebhookTypeDetailsResponse, ListAllOutgoingWebhooksRequest,
    ListAllOutgoingWebhooksResponse, ListOutgoingWebhookTypesRequest,
    ListOutgoingWebhookTypesResponse, ListOutgoingWebhooksRequest, ListOutgoingWebhooksResponse,
    OutgoingWebhookInputData, TestExistingOutgoingWebhookRequest, TestOutgoingWebhookRequest,
    TestOutgoingWebhookResponse, UpdateOutgoingWebhookRequest, UpdateOutgoingWebhookResponse,
};
use tokio::sync::Mutex;
use tonic::{
    metadata::MetadataMap,
    transport::{Channel, ClientTlsConfig, Endpoint},
};

pub use cx_api::proto::com::coralogix::outgoing_webhooks::v1::{
    generic_webhook_config, outgoing_webhook_input_data::Config, slack_config,
    test_outgoing_webhook_response::Result as WebhookTestResult, AwsEventBridgeConfig,
    DemistoConfig, EmailGroupConfig, GenericWebhookConfig, IbmEventNotificationsConfig, JiraConfig,
    MicrosoftTeamsConfig, OpsgenieConfig, PagerDutyConfig, SendLogConfig, SlackConfig, WebhookType,
};
use url::Url;

use crate::CoralogixRegion;

/// The Webhooks API client.
/// Read more at <https://coralogix.com/docs/webhooks-api/>
pub struct WebhooksClient {
    metadata_map: MetadataMap,
    service_client: Mutex<OutgoingWebhooksServiceClient<Channel>>,
}

impl WebhooksClient {
    /// Creates a new client for the Webhooks API.
    ///
    /// # Arguments
    /// * `auth_context` - The [`AuthContext`] to use for authentication.
    /// * `region` - The [`CoralogixRegion`] to connect to.
    pub fn new(auth_context: AuthContext, region: CoralogixRegion) -> Result<Self> {
        let channel: Channel = Endpoint::from_str(region.grpc_endpoint().as_str())?
            .tls_config(ClientTlsConfig::new().with_native_roots())?
            .connect_lazy();
        let auth_data: AuthData = (&auth_context.team_level_api_key).into();
        Ok(Self {
            metadata_map: auth_data.to_metadata_map(),
            service_client: Mutex::new(OutgoingWebhooksServiceClient::new(channel)),
        })
    }

    /// Creates a new Outgoing Webhook.
    ///
    /// # Arguments
    /// * `webhook_type` - The [`WebhookType`].
    /// * `name` - The name of the webhook.
    /// * `url` - The [`Url`] of the webhook.
    /// * `config` - The [`Config`] of the webhook.
    pub async fn create(
        &self,
        webhook_type: WebhookType,
        name: Option<String>,
        url: Url,
        config: Config,
    ) -> Result<CreateOutgoingWebhookResponse> {
        let request = make_request_with_metadata(
            CreateOutgoingWebhookRequest {
                data: Some(OutgoingWebhookInputData {
                    r#type: webhook_type.into(),
                    name,
                    url: Some(url.into()),
                    config: Some(config),
                }),
            },
            &self.metadata_map,
        );
        self.service_client
            .lock()
            .await
            .create_outgoing_webhook(request)
            .await
            .map(|r| r.into_inner())
            .map_err(From::from)
    }

    /// Update the Outgoing Webhook identified by its id.
    ///
    /// # Arguments
    /// * `webhook_id` - The id of the webhook to update.
    /// * `webhook_type` - The [`WebhookType`].
    /// * `name` - The name of the webhook.
    /// * `url` - The [`Url`] of the webhook.
    /// * `config` - The [`Config`] of the webhook.
    pub async fn replace(
        &self,
        webhook_id: String,
        name: Option<String>,
        webhook_type: WebhookType,
        url: Url,
        config: Config,
    ) -> Result<UpdateOutgoingWebhookResponse> {
        let request = make_request_with_metadata(
            UpdateOutgoingWebhookRequest {
                id: webhook_id,
                data: Some(OutgoingWebhookInputData {
                    r#type: webhook_type.into(),
                    name,
                    url: Some(url.into()),
                    config: Some(config),
                }),
            },
            &self.metadata_map,
        );
        self.service_client
            .lock()
            .await
            .update_outgoing_webhook(request)
            .await
            .map(|r| r.into_inner())
            .map_err(From::from)
    }

    /// Deletes the Outgoing Webhook.
    ///
    /// # Arguments
    /// * `webhook_id` - The id of the webhook to delete.
    pub async fn delete(&self, webhook_id: String) -> Result<DeleteOutgoingWebhookResponse> {
        let request = make_request_with_metadata(
            DeleteOutgoingWebhookRequest {
                id: Some(webhook_id),
            },
            &self.metadata_map,
        );
        self.service_client
            .lock()
            .await
            .delete_outgoing_webhook(request)
            .await
            .map(|r| r.into_inner())
            .map_err(From::from)
    }

    /// Retrieves the Outgoing Webhook by id.
    ///
    /// # Arguments
    /// * `webhook_id` - The id of the webhook to get.
    pub async fn get(&self, webhook_id: String) -> Result<GetOutgoingWebhookResponse> {
        let request = make_request_with_metadata(
            GetOutgoingWebhookRequest {
                id: Some(webhook_id),
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .get_outgoing_webhook(request)
            .await
            .map(|r| r.into_inner())
            .map_err(From::from)
    }

    /// Get outgoing [`WebhookType`] details.
    ///
    /// # Arguments
    /// * `webhook_type` - The type of the webhook.
    pub async fn get_outgoing_webhook_type_details(
        &self,
        webhook_type: WebhookType,
    ) -> Result<GetOutgoingWebhookTypeDetailsResponse> {
        let request = make_request_with_metadata(
            GetOutgoingWebhookTypeDetailsRequest {
                r#type: webhook_type.into(),
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .get_outgoing_webhook_type_details(request)
            .await
            .map(|r| r.into_inner())
            .map_err(From::from)
    }

    /// Tests the provided outgoing webhook
    ///
    /// # Arguments
    /// * `webhook_type` - The [`WebhookType`].
    /// * `name` - The name of the webhook.
    /// * `url` - The [`Url`] of the webhook.
    /// * `config` - The [`Config`] of the webhook.
    pub async fn test_webhook(
        &self,
        webhook_type: WebhookType,
        name: Option<String>,
        url: Url,
        config: Config,
    ) -> Result<TestOutgoingWebhookResponse> {
        let request = make_request_with_metadata(
            TestOutgoingWebhookRequest {
                data: Some(OutgoingWebhookInputData {
                    r#type: webhook_type.into(),
                    name,
                    url: Some(url.into()),
                    config: Some(config),
                }),
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .test_outgoing_webhook(request)
            .await
            .map(|r| r.into_inner())
            .map_err(From::from)
    }

    /// Tests the existing outgoing webhook
    ///
    /// # Arguments
    /// * `webhook_id` - The id of the webhook to test.
    pub async fn test_webhook_by_id(
        &self,
        webhook_id: String,
    ) -> Result<TestOutgoingWebhookResponse> {
        let request = make_request_with_metadata(
            TestExistingOutgoingWebhookRequest {
                id: Some(webhook_id),
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .test_existing_outgoing_webhook(request)
            .await
            .map(|r| r.into_inner())
            .map_err(From::from)
    }

    /// Get [`WebhookType`] details.
    ///     
    /// # Arguments
    /// * `webhook_type` - The [`WebhookType`].
    pub async fn get_type(
        &self,
        webhook_type: WebhookType,
    ) -> Result<ListOutgoingWebhooksResponse> {
        let request = make_request_with_metadata(
            ListOutgoingWebhooksRequest {
                r#type: webhook_type.into(),
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .list_outgoing_webhooks(request)
            .await
            .map(|r| r.into_inner())
            .map_err(From::from)
    }

    /// Retrieves a list of all outgoing webhooks.
    ///     
    /// # Returns  
    /// A list of all outgoing webhooks.
    pub async fn list(&self) -> Result<ListAllOutgoingWebhooksResponse> {
        let request =
            make_request_with_metadata(ListAllOutgoingWebhooksRequest {}, &self.metadata_map);

        self.service_client
            .lock()
            .await
            .list_all_outgoing_webhooks(request)
            .await
            .map(|r| r.into_inner())
            .map_err(From::from)
    }

    /// Retrieves a list of all [`WebhookType`]s.
    ///     
    /// # Returns
    /// A list of all webhook types.
    pub async fn list_webhook_types(&self) -> Result<ListOutgoingWebhookTypesResponse> {
        let request =
            make_request_with_metadata(ListOutgoingWebhookTypesRequest {}, &self.metadata_map);

        self.service_client
            .lock()
            .await
            .list_outgoing_webhook_types(request)
            .await
            .map(|r| r.into_inner())
            .map_err(From::from)
    }
}
