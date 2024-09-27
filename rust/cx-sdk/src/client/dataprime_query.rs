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

use crate::{
    CoralogixRegion,
    auth::AuthContext,
    error::Result,
    metadata::CallProperties,
    util::make_request_with_metadata,
};
use cx_api::proto::com::coralogixapis::dataprime::v1::{
    QueryRequest,
    QueryResponse,
    dataprime_query_service_client::DataprimeQueryServiceClient,
};
use std::str::FromStr;
use tokio::sync::Mutex;
use tonic::{
    Streaming,
    transport::ClientTlsConfig,
};
use tonic::{
    metadata::MetadataMap,
    transport::{
        Channel,
        Endpoint,
    },
};

pub use cx_api::proto::com::coralogixapis::dataprime::v1::Metadata;
pub use cx_api::proto::com::coralogixapis::dataprime::v1::query_response::Message;

/// The Dataprime Query API client.
/// Read more at [https://coralogix.com/docs/dataprime-query-language/]()
pub struct DataprimeQueryClient {
    metadata_map: MetadataMap,
    service_client: Mutex<DataprimeQueryServiceClient<Channel>>,
}

impl DataprimeQueryClient {
    /// Creates a new client for the Dataprime Query API.
    ///
    /// # Arguments
    /// * `auth_context` - The  to use for authentication.
    /// * `region` - The region to connect to.
    pub fn new(region: CoralogixRegion, auth_context: AuthContext) -> Result<Self> {
        let enrichments_service_channel: Channel = Endpoint::from_str(&region.grpc_endpoint())?
            .tls_config(ClientTlsConfig::new().with_native_roots())?
            .connect_lazy();
        let request_metadata: CallProperties = (&auth_context.team_level_api_key).into();
        Ok(Self {
            metadata_map: request_metadata.to_metadata_map(),
            service_client: Mutex::new(DataprimeQueryServiceClient::new(
                enrichments_service_channel.clone(),
            )),
        })
    }

    /// Runs a dataprime query.
    ///
    /// # Arguments
    /// * `query` - The query to run.
    pub async fn run(
        &self,
        query: String,
        metadata: Option<Metadata>,
    ) -> Result<Streaming<QueryResponse>> {
        let request = make_request_with_metadata(
            QueryRequest {
                query: Some(query),
                metadata,
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await;
            Ok(client.query(request).await?.into_inner())
        }
    }
}
