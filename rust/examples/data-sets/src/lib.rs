use anyhow::Result;
use cx_sdk::auth::{ApiKey, AuthData};
use cx_sdk::com::coralogix::enrichment::v1::custom_enrichment_service_client::CustomEnrichmentServiceClient;
use cx_sdk::com::coralogix::enrichment::v1::{file::Content as FileContent, File};
use cx_sdk::com::coralogix::enrichment::v1::{
     UpdateCustomEnrichmentRequest,
    UpdateCustomEnrichmentResponse,
};
use cx_sdk::com::coralogix::enrichment::v1::{
    CreateCustomEnrichmentRequest, CustomEnrichment, DeleteCustomEnrichmentResponse,
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

pub struct DataSetService {
    metadata_map: MetadataMap,
    data_sets: Mutex<CustomEnrichmentServiceClient<Channel>>,
}

impl DataSetService {
    pub fn new(endpoint: &str, api_key: String) -> Self {
        let enrichments_service_channel: Channel =
            Endpoint::from_str(endpoint).unwrap().connect_lazy();
        let api_key: ApiKey = api_key.into();
        let auth_data: AuthData = (&api_key).into();
        Self {
            metadata_map: auth_data.to_metadata_map(),
            data_sets: Mutex::new(CustomEnrichmentServiceClient::new(
                enrichments_service_channel,
            )),
        }
    }

    pub async fn create_data_set(
        &self,
        name: String,
        description: Option<String>,
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
                description,
            },
            &self.metadata_map,
        );
        {
            let mut client = self.data_sets.lock().await.clone();

            client
                .create_custom_enrichment(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    pub async fn update_data_set(
        &self,
        id: u32,
        name: String,
        description: Option<String>,
        data: String,
    ) -> Result<UpdateCustomEnrichmentResponse> {
        let file = File {
            name: Some(name.clone()),
            content: Some(FileContent::Binary(data.into_bytes())),
            extension: Some("csv".into()),
        };
        let request: Request<UpdateCustomEnrichmentRequest> = make_request_with_metadata(
            UpdateCustomEnrichmentRequest {
                custom_enrichment_id: Some(id),
                file: Some(file),
                name: Some(name.clone()),
                description,
            },
            &self.metadata_map,
        );
        {
            let mut client = self.data_sets.lock().await.clone();

            client
                .update_custom_enrichment(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    pub async fn delete_data_set(
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
            let mut client = self.data_sets.lock().await.clone();

            client
                .delete_custom_enrichment(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    pub async fn get_data_set(&self, id: u32) -> Result<Option<CustomEnrichment>> {
        let get_enrichments_request = GetCustomEnrichmentRequest { id: Some(id) };
        let request = make_request_with_metadata(get_enrichments_request, &self.metadata_map);
        {
            let mut client = self.data_sets.lock().await.clone();
            let response = client.get_custom_enrichment(request).await?;
            Ok(response.into_inner().custom_enrichment)
        }
    }

    pub async fn get_data_sets(&self) -> Result<Vec<CustomEnrichment>> {
        let get_enrichments_request = GetCustomEnrichmentsRequest {};
        let request = make_request_with_metadata(get_enrichments_request, &self.metadata_map);
        {
            let mut client = self.data_sets.lock().await.clone();
            let response = client.get_custom_enrichments(request).await?;
            Ok(response.into_inner().custom_enrichments)
        }
    }
}

#[cfg(test)]
mod tests {
    use tokio::fs;

    use crate::DataSetService;

    #[tokio::test]
    async fn test_custom_enrichment_service() {
        let svc = DataSetService::new("https://ng-api-grpc.eu2.coralogix.com", "my-api-key".into());
        let dataset = svc
            .create_data_set(
                "my-enrichment".into(),
                Some("My custom enrichment description".to_string()),
                fs::read_to_string("date-to-day-of-the-week.csv")
                    .await
                    .unwrap(),
            )
            .await
            .unwrap()
            .custom_enrichment
            .unwrap();
        let dataset_id = dataset.id;
        let updated = svc
            .update_data_set(
                dataset_id,
                "my-enrichment".into(),
                Some("My updated enrichment description".to_string()),
                fs::read_to_string("date-to-day-of-the-week.csv")
                    .await
                    .unwrap(),
            )
            .await
            .unwrap()
            .custom_enrichment
            .unwrap();

        let fetched = svc.get_enrichment(dataset_id).await.unwrap().unwrap();
        assert_eq!(updated.description, fetched.description);
        assert_eq!(fetched.version, dataset.version + 1);

        let _ = svc.delete_data_set(dataset_id).await.unwrap();
    }
}
