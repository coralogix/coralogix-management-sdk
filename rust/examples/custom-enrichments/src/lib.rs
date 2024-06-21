use anyhow::Result;
use cx_sdk::auth::{ApiKey, AuthData};
use cx_sdk::com::coralogix::enrichment::v1::custom_enrichment_service_client::CustomEnrichmentServiceClient;
use cx_sdk::com::coralogix::enrichment::v1::enrichment_service_client::EnrichmentServiceClient;
use cx_sdk::com::coralogix::enrichment::v1::enrichment_type::Type;
use cx_sdk::com::coralogix::enrichment::v1::{file::Content as FileContent, File};
use cx_sdk::com::coralogix::enrichment::v1::{
    AddEnrichmentsRequest, CustomEnrichmentType, EnrichmentRequestModel, EnrichmentType,
    GetEnrichmentsRequest, RemoveEnrichmentsRequest,
};
use cx_sdk::com::coralogix::enrichment::v1::{
    CreateCustomEnrichmentRequest, CustomEnrichment, DeleteCustomEnrichmentResponse, Enrichment,
    GetCustomEnrichmentRequest, GetCustomEnrichmentsRequest,
};
use cx_sdk::com::coralogix::enrichment::v1::{
    CreateCustomEnrichmentResponse, DeleteCustomEnrichmentRequest,
};
use std::str::FromStr;
use tokio::sync::Mutex;
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

pub struct EnrichmentService {
    metadata_map: MetadataMap,
    enrichment_key_client: Mutex<EnrichmentServiceClient<Channel>>,
    custom_enrichment_client: Mutex<CustomEnrichmentServiceClient<Channel>>,
}

impl EnrichmentService {
    pub fn new(endpoint: &str, api_key: String) -> Self {
        let enrichments_service_channel: Channel =
            Endpoint::from_str(endpoint).unwrap().connect_lazy();
        let api_key: ApiKey = api_key.into();
        let auth_data: AuthData = (&api_key).into();
        Self {
            metadata_map: auth_data.to_metadata_map(),
            enrichment_key_client: Mutex::new(EnrichmentServiceClient::new(
                enrichments_service_channel.clone(),
            )),
            custom_enrichment_client: Mutex::new(CustomEnrichmentServiceClient::new(
                enrichments_service_channel,
            )),
        }
    }

    #[allow(dead_code)]
    async fn create_enrichment(
        &self,
        name: String,
        data: String,
    ) -> Result<CreateCustomEnrichmentResponse> {
        let file = File {
            name: Some(name.clone()),
            content: Some(FileContent::Binary(data.into_bytes())),
            extension: Some("csv".into()),
        };
        let request: Request<CreateCustomEnrichmentRequest> = make_request_with_metadata(
            CreateCustomEnrichmentRequest {
                file: Some(file),
                name: Some(name.clone()),
                description: Some("description".to_string()),
            },
            &self.metadata_map,
        );
        {
            let mut client = self.custom_enrichment_client.lock().await.clone();

            client
                .create_custom_enrichment(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    #[allow(dead_code)]
    async fn delete_enrichment(
        &self,
        custom_enrichment_id: u32,
    ) -> Result<DeleteCustomEnrichmentResponse> {
        let request: Request<DeleteCustomEnrichmentRequest> = make_request_with_metadata(
            DeleteCustomEnrichmentRequest {
                custom_enrichment_id: Some(custom_enrichment_id),
            },
            &self.metadata_map,
        );
        {
            let mut client = self.custom_enrichment_client.lock().await.clone();

            client
                .delete_custom_enrichment(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    pub async fn list_enrichment_keys(&self) -> Result<Vec<Enrichment>> {
        let get_enrichments_request = GetEnrichmentsRequest {};
        let request = make_request_with_metadata(get_enrichments_request, &self.metadata_map);
        {
            let mut client = self.enrichment_key_client.lock().await.clone();
            let response = client.get_enrichments(request).await?;
            Ok(response.into_inner().enrichments.into_iter().collect())
        }
    }

    pub async fn get_enrichment(&self, id: u32) -> Result<Option<CustomEnrichment>> {
        let get_enrichments_request = GetCustomEnrichmentRequest { id: Some(id) };
        let request = make_request_with_metadata(get_enrichments_request, &self.metadata_map);
        {
            let mut client = self.custom_enrichment_client.lock().await.clone();
            let response = client.get_custom_enrichment(request).await?;
            Ok(response.into_inner().custom_enrichment)
        }
    }

    pub async fn get_enrichments(&self) -> Result<Vec<CustomEnrichment>> {
        let get_enrichments_request = GetCustomEnrichmentsRequest {};
        let request = make_request_with_metadata(get_enrichments_request, &self.metadata_map);
        {
            let mut client = self.custom_enrichment_client.lock().await.clone();
            let response = client.get_custom_enrichments(request).await?;
            Ok(response.into_inner().custom_enrichments)
        }
    }

    #[allow(dead_code)]
    async fn add_enrichment_key_mapping(
        &self,
        enrichment_id: u32,
        from: String,
        to: String,
    ) -> Result<()> {
        let add_enrichments_request = AddEnrichmentsRequest {
            request_enrichments: vec![EnrichmentRequestModel {
                field_name: Some(from),
                enrichment_type: Some(EnrichmentType {
                    r#type: Some(Type::CustomEnrichment(CustomEnrichmentType {
                        id: Some(enrichment_id),
                    })),
                }),
                enriched_field_name: Some(to),
            }],
        };
        let request: Request<AddEnrichmentsRequest> =
            make_request_with_metadata(add_enrichments_request, &self.metadata_map.clone());
        {
            let mut client = self.enrichment_key_client.lock().await.clone();
            let _ = client.add_enrichments(request).await;
        }
        Ok(())
    }

    #[allow(dead_code)]
    async fn delete_enrichment_keys(&self, enrichment_ids: Vec<u32>) -> Result<()> {
        let remove_enrichments_request = RemoveEnrichmentsRequest { enrichment_ids };
        let request = make_request_with_metadata(remove_enrichments_request, &self.metadata_map);
        {
            let mut client = self.enrichment_key_client.lock().await.clone();
            let _ = client.remove_enrichments(request).await?;
        }
        Ok(())
    }
}

#[cfg(test)]
mod tests {
    use crate::EnrichmentService;

    #[tokio::test]
    async fn test_custom_enrichment_service() {
        let svc =
            EnrichmentService::new("https://ng-api-grpc.eu2.coralogix.com", "my-api-key".into());
        let enrichment = svc
            .create_enrichment("my-enrichment".into(), "a,b,c\n1,2,3".into())
            .await
            .unwrap()
            .custom_enrichment
            .unwrap();
        let ids = svc.list_enrichment_keys().await.unwrap();
        let enrichment_id = enrichment.id;
        assert!(ids.iter().any(|e| e.id == enrichment.id));

        svc.add_enrichment_key_mapping(enrichment_id, "from".into(), "to".into())
            .await
            .unwrap();
        svc.delete_enrichment_keys(ids.iter().map(|e| e.id).collect())
            .await
            .unwrap();
        let _ = svc.delete_enrichment(enrichment_id).await.unwrap();
    }
}
