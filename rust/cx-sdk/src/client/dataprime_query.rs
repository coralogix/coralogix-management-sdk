use crate::{
    auth::{ApiKey, AuthData},
    error::Result,
    util::make_request_with_metadata,
    CoralogixRegion,
};
use cx_api::proto::com::coralogixapis::dataprime::v1::{
    dataprime_query_service_client::DataprimeQueryServiceClient, QueryRequest, QueryResponse,
};
use std::str::FromStr;
use tokio::sync::Mutex;
use tonic::Streaming;
use tonic::{
    metadata::MetadataMap,
    transport::{Channel, Endpoint},
};

pub use cx_api::proto::com::coralogixapis::dataprime::v1::query_response::Message;
pub use cx_api::proto::com::coralogixapis::dataprime::v1::Metadata;

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
    /// * `api_key` - The  to use for authentication.
    /// * `region` - The region to connect to.
    pub fn new(region: CoralogixRegion, api_key: ApiKey) -> Self {
        let enrichments_service_channel: Channel = Endpoint::from_str(&region.endpoint())
            .unwrap()
            .connect_lazy();
        let auth_data: AuthData = (&api_key).into();
        Self {
            metadata_map: auth_data.to_metadata_map(),
            service_client: Mutex::new(DataprimeQueryServiceClient::new(
                enrichments_service_channel.clone(),
            )),
        }
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
