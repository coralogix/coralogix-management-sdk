use anyhow::Result;
use cx_sdk::com::coralogixapis::dataprime::v1::dataprime_query_service_client::DataprimeQueryServiceClient;
use cx_sdk::com::coralogixapis::dataprime::v1::query_response::Message;
use cx_sdk::com::coralogixapis::dataprime::v1::QueryResponse;
use cx_sdk::{ApiKey, AuthData};
use std::str::FromStr;
use tokio::sync::Mutex;
use tonic::Streaming;
use tonic::{
    metadata::MetadataMap,
    transport::{Channel, Endpoint},
    Request,
};

pub fn make_request_with_metadata<T>(request: T, new_metadata: &MetadataMap) -> Request<T> {
    let mut req = Request::new(request);
    let metadata = req.metadata_mut();
    *metadata = new_metadata.clone();
    req
}

pub struct DataprimeQueryService {
    metadata_map: MetadataMap,
    dataprime_query_client: Mutex<DataprimeQueryServiceClient<Channel>>,
}

impl DataprimeQueryService {
    pub fn new(endpoint: &str, api_key: String) -> Self {
        let enrichments_service_channel: Channel =
            Endpoint::from_str(endpoint).unwrap().connect_lazy();
        let api_key: ApiKey = api_key.into();
        let auth_data: AuthData = (&api_key).into();
        Self {
            metadata_map: auth_data.to_metadata_map(),
            dataprime_query_client: Mutex::new(DataprimeQueryServiceClient::new(
                enrichments_service_channel.clone(),
            )),
        }
    }

    async fn run_query(&self, query: String) -> Result<Streaming<QueryResponse>> {
        let request = make_request_with_metadata(
            cx_sdk::com::coralogixapis::dataprime::v1::QueryRequest {
                query: Some(query),
                ..Default::default()
            },
            &self.metadata_map,
        );
        {
            let mut client = self.dataprime_query_client.lock().await;
            Ok(client.query(request).await?.into_inner())
        }
    }
}

#[tokio::main]
async fn main() {
    let svc =
        DataprimeQueryService::new("https://ng-api-grpc.eu2.coralogix.com", "my-api-key".into());

    let mut stream = svc
        .run_query("filter log_obj.message ~ 'Hello world'".to_string())
        .await
        .unwrap();
    while let Some(response) = stream.message().await.unwrap() {
        if let Some(Message::Result(result)) = response.message {
            for r in result.results {
                println!("{:?}", r.user_data);
            }
        }
    }
}
