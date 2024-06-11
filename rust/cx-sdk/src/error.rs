use http::uri::InvalidUri;
use tonic::transport::Error as TonicError;

#[derive(Debug, thiserror::Error)]
pub enum SdkError {
    #[error("Invalid endpoint URI")]
    TeamNotFound(#[from] InvalidUri),

    #[error("Endpoint not found")]
    EndpointNotFound(#[from] TonicError),
}
