use serde::{Deserialize, Serialize};

use crate::{auth::ApiKey, CoralogixRegion};

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
}

impl UsersClient {
    /// Create a new UsersClient.
    ///
    /// # Arguments
    /// * `region` - The [`CoralogixRegion`] to connect to.
    /// * `api_key` - Th [`ApiKey`] to use for authentication.
    pub fn new(region: CoralogixRegion, api_key: ApiKey) -> Self {
        Self {
            http_client: reqwest::Client::new(),
            api_key,
            target_url: region.rest_endpoint() + "/scim/Users",
        }
    }

    /// Create a [`ScimUser`].
    pub async fn create(&self, user: ScimUser) -> Result<ScimUser, reqwest::Error> {
        let json_request = serde_json::to_string(&user).unwrap();
        let response = self
            .http_client
            .post(&self.target_url)
            .header("Authorization", format!("Bearer {}", self.api_key.token()))
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
            .send()
            .await?;
        response.error_for_status()?;
        Ok(())
    }
}
