// Copyright 2024 Coralogix Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
