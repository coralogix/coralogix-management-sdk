use crate::error::Result;
use std::fmt::{Debug, Formatter};

use http::{HeaderMap, HeaderName, HeaderValue};
use tonic::metadata::MetadataMap;

const AUTHORIZATION_HEADER_NAME: &str = "authorization";
const ENV_API_KEY: &str = "CORALOGIX_API_KEY";

#[derive(Clone, Hash, PartialEq, Eq, PartialOrd, Ord)]
/// The API key used for authentication.
pub struct ApiKey(String);

impl ApiKey {
    /// Returns the API key as a string.
    pub fn token(&self) -> &str {
        &self.0
    }

    /// Creates a new API key from the ENV_API_KEY environment variable.
    pub fn from_env() -> Result<Self> {
        std::env::var(ENV_API_KEY).map(ApiKey).map_err(From::from)
    }
}

impl From<&str> for ApiKey {
    fn from(s: &str) -> Self {
        ApiKey(String::from(s))
    }
}

impl From<String> for ApiKey {
    fn from(s: String) -> Self {
        ApiKey(s)
    }
}

impl Debug for ApiKey {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        f.write_str("ApiKey(***)")
    }
}

#[derive(Default, Debug, Clone)]
/// The authentication data.
pub struct AuthData {
    header_map: HeaderMap,
}

impl AuthData {
    /// Creates a new AuthData instance.
    ///
    /// # Arguments
    /// * `header_map` - The header map to use for authentication.
    pub fn new(header_map: HeaderMap) -> Self {
        AuthData { header_map }
    }

    /// Returns the underlying header map.
    pub fn to_header_map(&self) -> HeaderMap {
        self.header_map.clone()
    }

    /// Returns the underlying header map as a MetadataMap.
    pub fn to_metadata_map(&self) -> MetadataMap {
        MetadataMap::from_headers(self.header_map.clone())
    }
}

impl From<&ApiKey> for AuthData {
    fn from(value: &ApiKey) -> Self {
        let value = bytes::Bytes::from(format!("Bearer {}", value.token()));
        let mut value = HeaderValue::from_maybe_shared(value).unwrap();
        value.set_sensitive(true);
        AuthData::new(HeaderMap::from_iter([(
            HeaderName::from_static(AUTHORIZATION_HEADER_NAME),
            value,
        )]))
    }
}
