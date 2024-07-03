use http::uri::InvalidUri;
use tonic::transport::Error as TonicError;

pub type Result<T> = std::result::Result<T, SdkError>;

#[derive(Debug, thiserror::Error)]
pub enum SdkError {
    #[error("Invalid endpoint URI: {0}")]
    TeamNotFound(#[from] InvalidUri),

    #[error("Endpoint not found: {0}")]
    EndpointNotFound(#[from] TonicError),

    #[error("API error: {0}")]
    ApiError(#[from] tonic::Status),

    #[error("Invalid environment variable for CORALOGIX_REGION: {0}")]
    EnvConfigurationError(#[from] std::env::VarError),
}
