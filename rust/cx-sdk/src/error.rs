
#[derive(Debug, thiserror::Error)]
pub enum SdkError {
    #[from(InvalidUri)]
    #[error("Invalid endpoint URI")]
    TeamNotFound,
}