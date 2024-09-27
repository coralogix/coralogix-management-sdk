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
