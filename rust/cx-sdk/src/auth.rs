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

use crate::error::Result;
use std::fmt::{
    Debug,
    Formatter,
};

const ENV_TEAM_API_KEY: &str = "CORALOGIX_TEAM_API_KEY";
const ENV_USER_API_KEY: &str = "CORALOGIX_USER_API_KEY";

#[derive(Clone, Hash, PartialEq, Eq, PartialOrd, Ord, Default)]
/// The authentication context.
pub struct AuthContext {
    /// The team level API key.
    pub(crate) team_level_api_key: ApiKey,
    /// The user level API key.
    pub(crate) user_level_api_key: ApiKey,
}

impl AuthContext {
    /// Creates a new AuthContext instance.
    ///
    /// # Arguments
    /// * `team_level_api_key` - The team level API key.
    /// * `user_level_api_key` - The user level API key.
    pub fn new(team_level_api_key: Option<ApiKey>, user_level_api_key: Option<ApiKey>) -> Self {
        if team_level_api_key.is_none() {
            log::warn!("Team level API key is not set. some functionality may not be available.");
        }
        if user_level_api_key.is_none() {
            log::warn!("User level API key is not set. some functionality may not be available.");
        }
        AuthContext {
            team_level_api_key: team_level_api_key.unwrap_or_default(),
            user_level_api_key: user_level_api_key.unwrap_or_default(),
        }
    }

    /// Creates a new AuthContext instance from environment variables.
    pub fn from_env() -> Self {
        AuthContext::new(
            ApiKey::teams_level_from_env().ok(),
            ApiKey::user_level_from_env().ok(),
        )
    }
}

#[derive(Clone, Hash, PartialEq, Eq, PartialOrd, Ord, Default)]
/// The API key used for authentication.
pub struct ApiKey(String);

impl ApiKey {
    /// Returns the API key as a string.
    pub fn token(&self) -> &str {
        &self.0
    }

    /// Creates a new teams-level API key from the ENV_API_KEY environment variable.
    pub fn teams_level_from_env() -> Result<Self> {
        std::env::var(ENV_TEAM_API_KEY)
            .map(ApiKey)
            .map_err(From::from)
    }

    /// Creates a new user-level API key from the ENV_API_KEY environment variable.
    pub fn user_level_from_env() -> Result<Self> {
        std::env::var(ENV_USER_API_KEY)
            .map(ApiKey)
            .map_err(From::from)
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
