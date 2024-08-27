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
use serde::{Deserialize, Serialize};

use crate::{
    auth::{ApiKey, AuthContext},
    CoralogixRegion, RUSTC_VERSION, SDK_CORRELATION_ID_HEADER_NAME, SDK_LANGUAGE_HEADER_NAME,
    SDK_RUSTC_VERSION_HEADER_NAME, SDK_VERSION, SDK_VERSION_HEADER_NAME,
};

#[derive(Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
/// A SCIM User.
pub struct ScimUser {
    /// The schemas of the SCIM user.
    pub schemas: Vec<String>,
    /// The ID of the SCIM user.
    pub id: Option<String>,
    /// The name of the SCIM user.
    pub user_name: String,
    /// Whether the SCIM user is active.
    pub active: bool,
    /// The name of the SCIM user.
    pub name: ScimUserName,
    /// The emails of the SCIM user.
    pub emails: Vec<ScimUserEmail>,
    /// The groups of the SCIM user.
    pub groups: Vec<ScimUserGroup>,
}

#[derive(Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
/// The name of a SCIM user.
pub struct ScimUserName {
    /// The given name of the SCIM user.
    pub given_name: String,
    /// The family name of the SCIM user.
    pub family_name: String,
}

#[derive(Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
/// The email of a SCIM user.
pub struct ScimUserEmail {
    /// The email address of the SCIM user.
    pub value: String,
    /// The type of the email address.
    pub r#type: Option<String>,
    /// Whether the email address is primary.
    pub primary: bool,
}

#[derive(Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
/// The phone number of a SCIM user.
pub struct ScimUserGroup {
    /// The value of the SCIM user group.
    pub value: String,
}

/// UsersClient is a client for the users API.
pub struct UsersClient {
    target_url: String,
    http_client: reqwest::Client,
    api_key: ApiKey,
    correlation_id: String,
}

impl UsersClient {
    /// Create a new UsersClient.
    ///
    /// # Arguments
    /// * `region` - The [`CoralogixRegion`] to connect to.
    /// * `auth_context` - Th [`AuthContext`] to use for authentication.
    pub fn new(region: CoralogixRegion, auth_context: AuthContext) -> Self {
        Self {
            http_client: reqwest::Client::new(),
            api_key: auth_context.team_level_api_key,
            target_url: region.rest_endpoint() + "/scim/Users",
            correlation_id: uuid::Uuid::new_v4().to_string(),
        }
    }

    /// Create a [`ScimUser`].
    pub async fn create(&self, user: ScimUser) -> Result<ScimUser, reqwest::Error> {
        let response = self
            .http_client
            .post(&self.target_url)
            .header("Authorization", format!("Bearer {}", self.api_key.token()))
            .header(SDK_VERSION_HEADER_NAME, SDK_VERSION)
            .header(SDK_LANGUAGE_HEADER_NAME, "rust")
            .header(SDK_RUSTC_VERSION_HEADER_NAME, RUSTC_VERSION)
            .header(SDK_CORRELATION_ID_HEADER_NAME, &self.correlation_id)
            .json(&user)
            .send()
            .await?;
        response.json().await
    }

    /// Get a SCIM user by ID.
    ///
    /// # Arguments
    /// * `id` - The ID of the SCIM user.
    pub async fn get(&self, id: &str) -> Result<ScimUser, reqwest::Error> {
        let url = format!("{}/{}", self.target_url, id);
        let response = self
            .http_client
            .get(&url)
            .header("Authorization", format!("Bearer {}", self.api_key.token()))
            .header(SDK_VERSION_HEADER_NAME, SDK_VERSION)
            .header(SDK_LANGUAGE_HEADER_NAME, "rust")
            .header(SDK_RUSTC_VERSION_HEADER_NAME, RUSTC_VERSION)
            .header(SDK_CORRELATION_ID_HEADER_NAME, &self.correlation_id)
            .send()
            .await?;
        response.json().await
    }

    /// Update a SCIM user.
    ///
    /// # Arguments
    /// * `user` - The [`ScimUser`] to update.
    pub async fn update(&self, user: ScimUser) -> Result<ScimUser, reqwest::Error> {
        let url = format!("{}/{}", self.target_url, user.id.as_ref().unwrap());
        let response = self
            .http_client
            .put(&url)
            .header("Authorization", format!("Bearer {}", self.api_key.token()))
            .header(SDK_VERSION_HEADER_NAME, SDK_VERSION)
            .header(SDK_LANGUAGE_HEADER_NAME, "rust")
            .header(SDK_RUSTC_VERSION_HEADER_NAME, RUSTC_VERSION)
            .header(SDK_CORRELATION_ID_HEADER_NAME, &self.correlation_id)
            .json(&user)
            .send()
            .await?;
        response.json().await
    }

    /// Delete a SCIM user by ID.
    ///
    /// # Arguments
    /// * `id` - The ID of the SCIM user.
    pub async fn delete(&self, id: &str) -> Result<(), reqwest::Error> {
        let url = format!("{}/{}", self.target_url, id);
        let response = self
            .http_client
            .delete(&url)
            .header("Authorization", format!("Bearer {}", self.api_key.token()))
            .header(SDK_VERSION_HEADER_NAME, SDK_VERSION)
            .header(SDK_LANGUAGE_HEADER_NAME, "rust")
            .header(SDK_RUSTC_VERSION_HEADER_NAME, RUSTC_VERSION)
            .header(SDK_CORRELATION_ID_HEADER_NAME, &self.correlation_id)
            .send()
            .await?;
        response.error_for_status()?;
        Ok(())
    }
}
