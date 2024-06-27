use std::str::FromStr;

use anyhow::Result;
use cx_sdk::{
    auth::{ApiKey, AuthData},
    com::coralogixapis::aaa::apikeys::v3::{
        api_keys_service_client::ApiKeysServiceClient,
        create_api_key_request::KeyPermissions,
        update_api_key_request::{Permissions, Presets},
        CreateApiKeyRequest, CreateApiKeyResponse, DeleteApiKeyRequest, GetApiKeyRequest,
        GetApiKeyResponse, Owner, UpdateApiKeyRequest, UpdateApiKeyResponse,
    },
};
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

pub struct ApiKeysService {
    metadata_map: MetadataMap,
    service_client: Mutex<ApiKeysServiceClient<Channel>>,
}

impl ApiKeysService {
    pub fn new(endpoint: &str, api_key: String) -> Self {
        let channel: Channel = Endpoint::from_str(endpoint).unwrap().connect_lazy();
        let api_key: ApiKey = api_key.into();
        let auth_data: AuthData = (&api_key).into();
        Self {
            metadata_map: auth_data.to_metadata_map(),
            service_client: Mutex::new(ApiKeysServiceClient::new(channel)),
        }
    }

    pub async fn create(
        &self,
        name: String,
        owner: Option<Owner>,
        presets: Vec<String>,
        permissions: Vec<String>,
        hashed: bool,
    ) -> Result<CreateApiKeyResponse> {
        let request = make_request_with_metadata(
            CreateApiKeyRequest {
                hashed,
                name,
                owner,
                key_permissions: Some(KeyPermissions {
                    presets,
                    permissions,
                }),
            },
            &self.metadata_map,
        );
        self.service_client
            .lock()
            .await
            .create_api_key(request)
            .await
            .map(|r| r.into_inner())
            .map_err(From::from)
    }

    pub async fn update(
        &self,
        key_id: String,
        is_active: Option<bool>,
        new_name: Option<String>,
        presets: Option<Presets>,
        permissions: Option<Permissions>,
    ) -> Result<UpdateApiKeyResponse> {
        let request = make_request_with_metadata(
            UpdateApiKeyRequest {
                key_id,
                new_name,
                is_active,
                presets,
                permissions,
            },
            &self.metadata_map,
        );
        self.service_client
            .lock()
            .await
            .update_api_key(request)
            .await
            .map(|r| r.into_inner())
            .map_err(From::from)
    }

    pub async fn delete(&self, key_id: String) -> Result<()> {
        let request =
            make_request_with_metadata(DeleteApiKeyRequest { key_id }, &self.metadata_map);
        self.service_client
            .lock()
            .await
            .delete_api_key(request)
            .await
            .map(|_| ())
            .map_err(From::from)
    }

    pub async fn get(&self, key_id: String) -> Result<GetApiKeyResponse> {
        let request =
            make_request_with_metadata(GetApiKeyRequest { key_id: key_id }, &self.metadata_map);
        self.service_client
            .lock()
            .await
            .get_api_key(request)
            .await
            .map(|r| r.into_inner())
            .map_err(From::from)
    }
}

#[cfg(test)]
mod tests {
    use cx_sdk::com::coralogixapis::aaa::apikeys::v3::owner::Owner;

    use crate::ApiKeysService;

    #[tokio::test]
    async fn test_actions_service() {
        let service = ApiKeysService::new(
            "https://ng-api-grpc.eu2.coralogix.com",
            "api-key".to_string(),
        );

        let create_result = service
            .create(
                "My APM KEY".to_string(),
                Some(cx_sdk::com::coralogixapis::aaa::apikeys::v3::Owner {
                    owner: Some(Owner::UserId("4013254".to_string())),
                }),
                vec!["APM".to_string()],
                vec![],
                false,
            )
            .await;

        assert!(create_result.is_ok());

        let key_id = create_result.unwrap().key_id;

        let update_result = service
            .update(
                key_id.clone(),
                None,
                Some("new-name".to_string()),
                None,
                None,
            )
            .await;
        assert!(update_result.is_ok());

        let new_api_key = service.get(key_id.clone()).await;

        assert!(new_api_key.is_ok());
        assert_eq!(
            new_api_key.unwrap().key_info.map(|k| k.name),
            Some("new-name".to_string())
        );

        let delete_action_result = service.delete(key_id).await;
        assert!(delete_action_result.is_ok());
    }
}
