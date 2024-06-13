use std::fmt::{Debug, Formatter};

use http::{HeaderMap, HeaderName, HeaderValue};
use tonic::metadata::MetadataMap;

const AUTHORIZATION_HEADER_NAME: &str = "authorization";

#[derive(Clone, Hash, PartialEq, Eq, PartialOrd, Ord)]
pub struct ApiKey(String);

impl ApiKey {
    pub fn token(&self) -> &str {
        &self.0
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
pub struct AuthData {
    header_map: HeaderMap,
}

impl AuthData {
    pub fn new(header_map: HeaderMap) -> Self {
        AuthData { header_map }
    }

    pub fn to_header_map(&self) -> HeaderMap {
        self.header_map.clone()
    }

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
