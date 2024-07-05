use std::str::FromStr;

use crate::{
    auth::{ApiKey, AuthData},
    error::Result,
    util::make_request_with_metadata,
};

pub use crate::com::coralogixapis::actions::v2::Action;

use cx_api::proto::com::coralogix::enrichment::v1::{
    custom_enrichment_service_client::CustomEnrichmentServiceClient, file::Content,
    CreateCustomEnrichmentRequest, CreateCustomEnrichmentResponse, DeleteCustomEnrichmentRequest,
    DeleteCustomEnrichmentResponse, File, GetCustomEnrichmentRequest, GetCustomEnrichmentsRequest,
    UpdateCustomEnrichmentRequest, UpdateCustomEnrichmentResponse,
};
use tokio::sync::Mutex;
use tonic::{
    metadata::MetadataMap,
    transport::{Channel, Endpoint},
    Request,
};

use crate::CoralogixRegion;

pub use crate::com::coralogix::enrichment::v1::CustomEnrichment;

/// The Custom Enrichments API client.
/// Read more at <https://coralogix.com/docs/custom-enrichment-api/>
pub struct DatasetClient {
    metadata_map: MetadataMap,
    service_client: Mutex<CustomEnrichmentServiceClient<Channel>>,
}

impl DatasetClient {
    /// Creates a new client for the Custom Enrichments API.
    ///
    /// # Arguments
    /// * `api_key` - The API key to use for authentication.
    /// * `region` - The region to connect to.
    pub fn new(api_key: ApiKey, region: CoralogixRegion) -> Result<Self> {
        let channel: Channel = Endpoint::from_str(region.endpoint().as_str())?.connect_lazy();
        let auth_data: AuthData = (&api_key).into();
        Ok(Self {
            metadata_map: auth_data.to_metadata_map(),
            service_client: Mutex::new(CustomEnrichmentServiceClient::new(channel)),
        })
    }

    /// Creates a new dataset.
    ///
    /// # Arguments
    /// * `name` - The name of the Custom Enrichment.
    /// * `description` - The description of the Custom Enrichment.
    /// * `data` - The data of the Custom Enrichment.
    pub async fn create(
        &self,
        name: String,
        description: Option<String>,
        data: Vec<u8>,
    ) -> Result<CreateCustomEnrichmentResponse> {
        let file = File {
            name: Some(name.clone()),
            content: Some(Content::Binary(data)),
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
            let mut client = self.service_client.lock().await.clone();

            client
                .create_custom_enrichment(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    /// Replaces the existing dataset identified by its id.
    ///
    /// # Arguments
    /// * `id` - The id of the Custom Enrichment.
    /// * `name` - The name of the Custom Enrichment.
    /// * `description` - The description of the Custom Enrichment.
    /// * `data` - The data of the Custom Enrichment.
    ///
    /// Note that the existing dataset will be replaced with a new version of the data.
    pub async fn replace(
        &self,
        id: u32,
        name: String,
        description: Option<String>,
        data: Vec<u8>,
    ) -> Result<UpdateCustomEnrichmentResponse> {
        let file = File {
            name: Some(name.clone()),
            content: Some(Content::Binary(data)),
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
            let mut client = self.service_client.lock().await.clone();

            client
                .update_custom_enrichment(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    /// Deletes the dataset identified by its id.
    ///
    /// # Arguments
    /// * `custom_enrichment_id` - The id of the Custom Enrichment.
    pub async fn delete(
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
            let mut client = self.service_client.lock().await.clone();

            client
                .delete_custom_enrichment(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    /// Retrieves the dataset by id.
    ///
    /// # Arguments
    /// * `id` - The id of the Custom Enrichment.
    pub async fn get(&self, id: u32) -> Result<Option<CustomEnrichment>> {
        let get_enrichments_request = GetCustomEnrichmentRequest { id: Some(id) };
        let request = make_request_with_metadata(get_enrichments_request, &self.metadata_map);
        {
            let mut client = self.service_client.lock().await.clone();
            let response = client.get_custom_enrichment(request).await?;
            Ok(response.into_inner().custom_enrichment)
        }
    }

    /// Retrieves all datasets.
    ///
    /// # Returns
    /// A list of all Custom Enrichments.
    pub async fn list(&self) -> Result<Vec<CustomEnrichment>> {
        let get_enrichments_request = GetCustomEnrichmentsRequest {};
        let request = make_request_with_metadata(get_enrichments_request, &self.metadata_map);
        {
            let mut client = self.service_client.lock().await.clone();
            let response = client.get_custom_enrichments(request).await?;
            Ok(response.into_inner().custom_enrichments)
        }
    }
}
