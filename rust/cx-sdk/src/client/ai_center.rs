use std::str::FromStr;

use crate::{
    auth::AuthContext,
    error::{Result, SdkApiError, SdkError},
    metadata::CallProperties,
    util::make_request_with_metadata,
};

use cx_api::proto::com::coralogixapis::ai::v3::{
    AiEvaluation, CountAiAppsPerEvalTypeRequest, CountAiAppsPerEvalTypeResponse,
    CreateAiEvaluationRequest, DeleteAiEvaluationRequest, GetAiEvaluationRequest,
    ListAiApplicationsRequest, ListAiApplicationsResponse, ListAiEvaluationsRequest,
    ListAiEvaluationsResponse, UpdateAiEvaluationRequest,
    ai_applications_service_client::AiApplicationsServiceClient,
    ai_evaluations_service_client::AiEvaluationsServiceClient,
};

use tonic::{metadata::MetadataMap, transport::{Channel, ClientTlsConfig, Endpoint}};
use tokio::sync::Mutex;

use crate::CoralogixRegion;

const AI_CENTER_FEATURE_GROUP_ID: &str = "ai";

/// A client for interacting with the Coralogix AI Center API.
pub struct AiCenterClient {
    metadata_map: MetadataMap,
    evaluations_client: Mutex<AiEvaluationsServiceClient<Channel>>,
    applications_client: Mutex<AiApplicationsServiceClient<Channel>>,
}

impl AiCenterClient {
    /// Create a new AI Center client.
    pub fn new(auth_context: AuthContext, region: CoralogixRegion) -> Result<Self> {
        let channel: Channel = Endpoint::from_str(region.grpc_endpoint().as_str())?
            .tls_config(ClientTlsConfig::new().with_native_roots())?
            .connect_lazy();
        let request_metadata: CallProperties = (&auth_context.team_level_api_key).into();
        Ok(Self {
            metadata_map: request_metadata.to_metadata_map(),
            evaluations_client: Mutex::new(AiEvaluationsServiceClient::new(channel.clone())),
            applications_client: Mutex::new(AiApplicationsServiceClient::new(channel)),
        })
    }

    pub async fn create_evaluation(&self, req: CreateAiEvaluationRequest) -> Result<AiEvaluation> {
        Ok(self
            .evaluations_client
            .lock()
            .await
            .create_ai_evaluation(make_request_with_metadata(req, &self.metadata_map))
            .await
            .map_err(|status| SdkError::ApiError(SdkApiError {
                status,
                endpoint: "/com.coralogixapis.ai.v3.AiEvaluationsService/CreateAiEvaluation"
                    .to_string(),
                feature_group: AI_CENTER_FEATURE_GROUP_ID.into(),
            }))?
            .into_inner())
    }

    pub async fn get_evaluation(&self, req: GetAiEvaluationRequest) -> Result<AiEvaluation> {
        Ok(self
            .evaluations_client
            .lock()
            .await
            .get_ai_evaluation(make_request_with_metadata(req, &self.metadata_map))
            .await
            .map_err(|status| SdkError::ApiError(SdkApiError {
                status,
                endpoint: "/com.coralogixapis.ai.v3.AiEvaluationsService/GetAiEvaluation"
                    .to_string(),
                feature_group: AI_CENTER_FEATURE_GROUP_ID.into(),
            }))?
            .into_inner())
    }

    pub async fn update_evaluation(&self, req: UpdateAiEvaluationRequest) -> Result<AiEvaluation> {
        Ok(self
            .evaluations_client
            .lock()
            .await
            .update_ai_evaluation(make_request_with_metadata(req, &self.metadata_map))
            .await
            .map_err(|status| SdkError::ApiError(SdkApiError {
                status,
                endpoint: "/com.coralogixapis.ai.v3.AiEvaluationsService/UpdateAiEvaluation"
                    .to_string(),
                feature_group: AI_CENTER_FEATURE_GROUP_ID.into(),
            }))?
            .into_inner())
    }

    pub async fn delete_evaluation(&self, req: DeleteAiEvaluationRequest) -> Result<()> {
        self.evaluations_client
            .lock()
            .await
            .delete_ai_evaluation(make_request_with_metadata(req, &self.metadata_map))
            .await
            .map_err(|status| SdkError::ApiError(SdkApiError {
                status,
                endpoint: "/com.coralogixapis.ai.v3.AiEvaluationsService/DeleteAiEvaluation"
                    .to_string(),
                feature_group: AI_CENTER_FEATURE_GROUP_ID.into(),
            }))?;
        Ok(())
    }

    pub async fn list_evaluations(
        &self,
        req: ListAiEvaluationsRequest,
    ) -> Result<ListAiEvaluationsResponse> {
        Ok(self
            .evaluations_client
            .lock()
            .await
            .list_ai_evaluations(make_request_with_metadata(req, &self.metadata_map))
            .await
            .map_err(|status| SdkError::ApiError(SdkApiError {
                status,
                endpoint: "/com.coralogixapis.ai.v3.AiEvaluationsService/ListAiEvaluations"
                    .to_string(),
                feature_group: AI_CENTER_FEATURE_GROUP_ID.into(),
            }))?
            .into_inner())
    }

    pub async fn count_apps_per_eval_type(
        &self,
        req: CountAiAppsPerEvalTypeRequest,
    ) -> Result<CountAiAppsPerEvalTypeResponse> {
        Ok(self
            .evaluations_client
            .lock()
            .await
            .count_ai_apps_per_eval_type(make_request_with_metadata(req, &self.metadata_map))
            .await
            .map_err(|status| SdkError::ApiError(SdkApiError {
                status,
                endpoint:
                    "/com.coralogixapis.ai.v3.AiEvaluationsService/CountAiAppsPerEvalType"
                        .to_string(),
                feature_group: AI_CENTER_FEATURE_GROUP_ID.into(),
            }))?
            .into_inner())
    }

    pub async fn list_applications(
        &self,
        req: ListAiApplicationsRequest,
    ) -> Result<ListAiApplicationsResponse> {
        Ok(self
            .applications_client
            .lock()
            .await
            .list_ai_applications(make_request_with_metadata(req, &self.metadata_map))
            .await
            .map_err(|status| SdkError::ApiError(SdkApiError {
                status,
                endpoint: "/com.coralogixapis.ai.v3.AiApplicationsService/ListAiApplications"
                    .to_string(),
                feature_group: AI_CENTER_FEATURE_GROUP_ID.into(),
            }))?
            .into_inner())
    }
}
