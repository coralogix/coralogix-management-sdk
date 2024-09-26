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

#[cfg(test)]
mod tests {
    use std::vec;

    use cx_sdk::{
        auth::AuthContext,
        client::users::{
            ScimUser,
            ScimUserEmail,
            ScimUserGroup,
            ScimUserName,
            UsersClient,
        },
        CoralogixRegion,
    };

    #[tokio::test]
    async fn test_users() {
        let client = UsersClient::new(
            CoralogixRegion::from_env().unwrap(),
            AuthContext::from_env(),
        );

        let scim_user = ScimUser {
            schemas: vec![],
            active: true,
            id: None,
            user_name: "yak@coralogix.com".into(),
            name: ScimUserName {
                given_name: "example".into(),
                family_name: "example".into(),
            },
            emails: vec![ScimUserEmail {
                value: "example@coralogix.com".into(),
                r#type: Some("work".into()),
                primary: true,
            }],
            groups: vec![ScimUserGroup {
                value: "Admins".into(),
            }],
        };

        let created_user = client.create(scim_user).await.unwrap();

        client
            .get(created_user.id.clone().unwrap().as_str())
            .await
            .unwrap();

        client
            .delete(created_user.id.clone().unwrap().as_str())
            .await
            .unwrap();
    }
}
