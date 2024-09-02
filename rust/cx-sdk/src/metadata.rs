use tonic::metadata::MetadataMap;

use http::HeaderMap;

#[derive(Default, Debug, Clone)]
/// Request metadata. It contains the metadata needed for authorization and metrics gathering.
pub struct CallProperties {
    pub(crate) header_map: HeaderMap,
}

impl CallProperties {
    /// Creates a new AuthData instance.
    ///
    /// # Arguments
    /// * `header_map` - The header map to use for authentication.
    pub fn new(header_map: HeaderMap) -> Self {
        CallProperties { header_map }
    }

    /// Returns the underlying header map as a MetadataMap.
    pub fn to_metadata_map(&self) -> MetadataMap {
        MetadataMap::from_headers(self.header_map.clone())
    }
}
