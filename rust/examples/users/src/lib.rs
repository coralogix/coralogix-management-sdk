#[cfg(test)]
mod tests {
    use std::vec;

    use cx_sdk::{
        auth::ApiKey,
        client::users::{ScimUser, ScimUserEmail, ScimUserGroup, ScimUserName, UsersClient},
        CoralogixRegion,
    };

    #[tokio::test]
    async fn users_test() {
        let client = UsersClient::new(
            CoralogixRegion::from_env().unwrap(),
            ApiKey::from_env().unwrap(),
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
