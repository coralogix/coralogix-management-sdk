use cx_sdk::com::coralogix::enrichment::v1::enrichment_service_client::EnrichmentServiceClient;
use cx_sdk::com::coralogix::enrichment::v1::{GetEnrichmentsRequest, RemoveEnrichmentsRequest, AddEnrichmentsRequest, EnrichmentRequestModel, EnrichmentType, CustomEnrichmentType};
use cx_sdk::com::coralogix::enrichment::v1::enrichment_type::Type;
use cx_sdk::ApiKey;
use std::str::FromStr;
use tokio::sync::Mutex;
use tonic::{
    metadata::MetadataMap,
    transport::{Channel, Endpoint},
    Code, Request, Status,
};
use anyhow::Result;

type ExternalEnrichmentClient = Mutex<EnrichmentServiceClient<Channel>>;

pub fn make_request_with_metadata<T>(request: T, new_metadata: &MetadataMap) -> Request<T> {
    let mut req = Request::new(request);
    let metadata = req.metadata_mut();
    *metadata = new_metadata.clone();
    req
}

pub struct ExternalEnrichmentService {
    metadata_map: MetadataMap,
    enrichment_client: ExternalEnrichmentClient,
}

impl ExternalEnrichmentService {
    #[allow(dead_code)]
    pub fn new(endpoint: &str, api_key: String) -> Self {
        let enrichments_service_channel: Channel =
            Endpoint::from_str(endpoint).unwrap().connect_lazy();
        let api_key: ApiKey = api_key.into();
        Self {
            metadata_map: api_key.to_metadata_map(),
            enrichment_client: Mutex::new(EnrichmentServiceClient::new(
                enrichments_service_channel,
            )),
        }
    }
}

impl ExternalEnrichmentService {
    async fn set_enrichment_key(
        &self,
        enrichment_id: u32,
        _company_id: u32,
    ) -> Result<()> {
        let add_enrichments_request = AddEnrichmentsRequest {
            request_enrichments: vec![EnrichmentRequestModel {
                field_name: Some("input".to_string()),
                enrichment_type: Some(EnrichmentType {
                    r#type: Some(Type::CustomEnrichment(CustomEnrichmentType {
                        id: Some(enrichment_id),
                    })),
                }),
                enriched_field_name: Some("enriched".to_string()),
            }],
        };
        let request: Request<AddEnrichmentsRequest> =
            make_request_with_metadata(add_enrichments_request, &self.metadata_map.clone());
        {
            let mut client = self.enrichment_client.lock().await.clone();
            let _ = client.add_enrichments(request).await;
        }
        Ok(())
    }

    async fn delete_enrichment_keys(
        &self,
        enrichment_ids: Vec<u32>,
        _company_id: u32,
    ) -> Result<()> {
        let remove_enrichments_request = RemoveEnrichmentsRequest { enrichment_ids };
        let request = make_request_with_metadata(remove_enrichments_request, &self.metadata_map);
        {
            let mut client = self.enrichment_client.lock().await.clone();
            let _ = client.remove_enrichments(request).await?;
        }
        Ok(())
    }

    pub async fn get_ids(&self) -> Result<Vec<u32>> {
        let get_enrichments_request = GetEnrichmentsRequest {};
        let request = make_request_with_metadata(get_enrichments_request, &self.metadata_map);
        {
            let mut client = self.enrichment_client.lock().await.clone();
            let response = client.get_enrichments(request).await?;
            Ok(response
                .into_inner()
                .enrichments
                .into_iter()
                .map(|e| e.id)
                .collect())
        }
    }
}

fn main() {
    println!("Hello, world!");
}
