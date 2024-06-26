use anyhow::Result;
use cx_sdk::auth::{ApiKey, AuthData};
use cx_sdk::com::coralogix::archive::v2::set_target_request::TargetSpec;
use cx_sdk::com::coralogix::archive::v2::target_service_client::TargetServiceClient;
use cx_sdk::com::coralogix::archive::v2::{
    validate_target_request, GetTargetRequest, GetTargetResponse, SetTargetRequest,
    SetTargetResponse, ValidateTargetRequest,
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

pub struct LogsArchiveService {
    metadata_map: MetadataMap,
    service_client: Mutex<TargetServiceClient<Channel>>,
}

impl LogsArchiveService {
    pub fn new(endpoint: &str, api_key: String) -> Self {
        let channel: Channel = Endpoint::from_str(endpoint).unwrap().connect_lazy();
        let api_key: ApiKey = api_key.into();
        let auth_data: AuthData = (&api_key).into();
        Self {
            metadata_map: auth_data.to_metadata_map(),
            service_client: Mutex::new(TargetServiceClient::new(channel)),
        }
    }

    pub async fn get_target(&self) -> Result<GetTargetResponse> {
        let request = make_request_with_metadata(GetTargetRequest {}, &self.metadata_map);
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .get_target(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    pub async fn set_target(
        &self,
        is_active: bool,
        target_spec: TargetSpec,
    ) -> Result<SetTargetResponse> {
        let request = make_request_with_metadata(
            SetTargetRequest {
                is_active,
                target_spec: Some(target_spec),
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .set_target(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    pub async fn validate_target(
        &self,
        is_active: bool,
        target_spec: validate_target_request::TargetSpec,
    ) -> Result<()> {
        let request = make_request_with_metadata(
            ValidateTargetRequest {
                target_spec: Some(target_spec),
                is_active,
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .validate_target(request)
                .await
                .map(|_| ())
                .map_err(From::from)
        }
    }
}

#[cfg(test)]
mod tests {
    use cx_sdk::com::coralogix::archive::v2::{
        set_target_request, validate_target_request, S3TargetSpec,
    };

    #[tokio::test]
    async fn test_logs_archive() {
        let endpoint = "https://ng-api-grpc.eu2.coralogix.com";
        let api_key = "api_key".to_string();
        let logs_archive_service = super::LogsArchiveService::new(endpoint, api_key);

        let validate_target_result = logs_archive_service
            .validate_target(
                true,
                validate_target_request::TargetSpec::S3(S3TargetSpec {
                    bucket: "my-very-cool-bucket".to_string(),
                    region: Some("us1".to_string()),
                }),
            )
            .await;

        assert!(validate_target_result.is_ok());

        let set_target_result = logs_archive_service
            .set_target(
                true,
                set_target_request::TargetSpec::S3(S3TargetSpec {
                    bucket: "my-very-cool-bucket".to_string(),
                    region: Some("us1".to_string()),
                }),
            )
            .await;

        assert!(set_target_result.is_ok());

        let expected_target = set_target_result.unwrap().target.unwrap();

        let get_target_response = logs_archive_service.get_target().await;

        assert!(get_target_response.is_ok());

        let actual_target = get_target_response.unwrap().target.unwrap();

        assert!(expected_target == actual_target);
    }
}
