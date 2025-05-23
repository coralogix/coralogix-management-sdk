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

#![allow(clippy::result_large_err)]
#![allow(renamed_and_removed_lints)]

//! The Rust SDK for Coralogix  APIs.
use auth::ApiKey;
use cx_api::proto::*;
use error::SdkError;
use http::{
    HeaderMap,
    HeaderName,
    HeaderValue,
};
use std::{
    fmt::Debug,
    str::FromStr,
};

/// This module contains the authentication primitives for the SDK.
pub mod auth;

/// This module contains the clients for the underlying Coralogix APIs.
pub mod client;

/// This module contains the error types for the SDK.
pub mod error;

/// This module contains the utility functions for the SDK.
mod util;

mod metadata;

const ENV_CORALOGIX_REGION: &str = "CORALOGIX_REGION";
const AUTHORIZATION_HEADER_NAME: &str = "authorization";
const SDK_VERSION_HEADER_NAME: &str = "x-cx-sdk-version";
const SDK_LANGUAGE_HEADER_NAME: &str = "x-cx-sdk-language";
const SDK_RUSTC_VERSION_HEADER_NAME: &str = "x-cx-sdk-rustc-version";
const SDK_CORRELATION_ID_HEADER_NAME: &str = "x-cx-sdk-correlation-id";
const RUSTC_VERSION: &str = env!("RUSTC_VERSION");

/// The SDK version.
pub const SDK_VERSION: &str = env!("CARGO_PKG_VERSION");

/// From <https://coralogix.com/docs/coralogix-domain/>
#[derive(Debug, Clone)]
pub enum CoralogixRegion {
    /// US1 region. Associated to the endpoint `https://ng-api-grpc.coralogix.com`
    US1,
    /// US2 region. Associated to the endpoint `https://ng-api-grpc.cx498.coralogix.com`
    US2,
    /// EU1 region. Associated to the endpoint `https://ng-api-grpc.coralogix.com`
    EU1,
    /// EU2 region. Associated to the endpoint `https://ng-api-grpc.eu2.coralogix.com`
    EU2,
    /// AP1 region. Associated to the endpoint `https://ng-api-grpc.app.coralogix.in`
    AP1,
    /// AP2 region. Associated to the endpoint `https://ng-api-grpc.coralogixsg.com`
    AP2,
    /// AP3 region. Associated to the endpoint `https://ng-api-grpc.ap3.coralogix.com`
    AP3,
    /// Custom region. It's associated with a custom endpoint.
    Custom(String),
}

impl CoralogixRegion {
    /// gRPC endpoint for the corresponding Coralogix region
    /// <https://coralogix.com/docs/coralogix-domain/>
    pub fn grpc_endpoint(&self) -> String {
        match self {
            CoralogixRegion::US1 => "https://ng-api-grpc.coralogix.us".into(),
            CoralogixRegion::US2 => "https://ng-api-grpc.cx498.coralogix.com".into(),
            CoralogixRegion::EU1 => "https://ng-api-grpc.coralogix.com".into(),
            CoralogixRegion::EU2 => "https://ng-api-grpc.eu2.coralogix.com".into(),
            CoralogixRegion::AP1 => "https://ng-api-grpc.app.coralogix.in".into(),
            CoralogixRegion::AP2 => "https://ng-api-grpc.coralogixsg.com".into(),
            CoralogixRegion::AP3 => "https://ng-api-grpc.ap3.coralogix.com".into(),
            CoralogixRegion::Custom(custom) => format!("https://ng-api-grpc.{custom}:443"),
        }
    }

    /// REST endpoint for the corresponding Coralogix region
    /// <https://coralogix.com/docs/coralogix-domain/>
    pub fn rest_endpoint(&self) -> String {
        match self {
            CoralogixRegion::US1 => "https://ng-api-http.coralogix.us".into(),
            CoralogixRegion::US2 => "https://ng-api-http.cx498.coralogix.com".into(),
            CoralogixRegion::EU1 => "https://ng-api-http.coralogix.com".into(),
            CoralogixRegion::EU2 => "https://ng-api-http.eu2.coralogix.com".into(),
            CoralogixRegion::AP1 => "https://ng-api-http.app.coralogix.in".into(),
            CoralogixRegion::AP2 => "https://ng-api-http.coralogixsg.com".into(),
            CoralogixRegion::AP3 => "https://ng-api-http.ap3.coralogix.com".into(),
            CoralogixRegion::Custom(custom) => format!("https://ng-api-http.{custom}"),
        }
    }

    /// Creates a CoralogixRegion from the environment variable `ENV_CORALOGIX_REGION`.
    pub fn from_env() -> Result<Self, SdkError> {
        let region = std::env::var(ENV_CORALOGIX_REGION)?;
        Self::from_str(&region)
    }
}

impl FromStr for CoralogixRegion {
    type Err = SdkError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        match s.to_lowercase().as_str() {
            "us1" => Ok(CoralogixRegion::US1),
            "us2" => Ok(CoralogixRegion::US2),
            "eu1" => Ok(CoralogixRegion::EU1),
            "eu2" => Ok(CoralogixRegion::EU2),
            "ap1" => Ok(CoralogixRegion::AP1),
            "ap2" => Ok(CoralogixRegion::AP2),
            "ap3" => Ok(CoralogixRegion::AP3),
            custom => Ok(CoralogixRegion::Custom(custom.to_string())),
        }
    }
}

impl From<&ApiKey> for metadata::CallProperties {
    fn from(value: &ApiKey) -> Self {
        let value = bytes::Bytes::from(format!("Bearer {}", value.token()));
        let mut value = HeaderValue::from_maybe_shared(value).unwrap();
        value.set_sensitive(true);
        metadata::CallProperties::new(HeaderMap::from_iter(
            [
                (HeaderName::from_static(AUTHORIZATION_HEADER_NAME), value),
                (
                    HeaderName::from_static(SDK_RUSTC_VERSION_HEADER_NAME),
                    HeaderValue::from_str(RUSTC_VERSION).unwrap(),
                ),
                (
                    HeaderName::from_static(SDK_VERSION_HEADER_NAME),
                    HeaderValue::from_str(format!("sdk-{SDK_VERSION}").as_str()).unwrap(),
                ),
                (
                    HeaderName::from_static(SDK_LANGUAGE_HEADER_NAME),
                    HeaderValue::from_static("rust"),
                ),
                (
                    HeaderName::from_static(SDK_CORRELATION_ID_HEADER_NAME),
                    HeaderValue::from_str(&uuid::Uuid::new_v4().to_string()).unwrap(),
                ),
            ]
            .iter()
            .cloned(),
        ))
    }
}
