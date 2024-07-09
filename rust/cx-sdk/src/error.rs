use http::uri::InvalidUri;
use tonic::transport::Error as TonicError;

/// The result type for the SDK.
pub type Result<T> = std::result::Result<T, SdkError>;

/// The error type for the SDK.
#[derive(Debug, thiserror::Error)]
pub enum SdkError {
    /// This error is returned when the team that an operation refers to is not found.
    #[error("Invalid endpoint URI: {0}")]
    TeamNotFound(#[from] InvalidUri),

    /// This error is returned when an endpoint is not found.
    #[error("Endpoint not found: {0}")]
    EndpointNotFound(#[from] TonicError),

    /// This error is returned when an underying API throws an error.
    #[error("API error: {0}")]
    ApiError(#[from] tonic::Status),

    /// This error is returned when an environment variable is not found or is invalid.
    #[error("Invalid environment variable for CORALOGIX_REGION: {0}")]
    EnvConfigurationError(#[from] std::env::VarError),
}
