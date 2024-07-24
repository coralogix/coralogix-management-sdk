#[cfg(test)]
mod tests {

    use cx_sdk::client::groups::{RoleId, TeamId};
    use cx_sdk::{auth::ApiKey, client::groups::GroupsClient, CoralogixRegion};

    #[tokio::test]
    async fn test_groups() {
        let team_id_value = std::env::var("TEAM_ID").unwrap().parse::<u32>().unwrap();
        let client = GroupsClient::new(
            ApiKey::from_env().unwrap(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();
        let team_id = TeamId { id: team_id_value };
        let groups = client.list(team_id.clone()).await.unwrap();
        assert!(!groups.groups.is_empty());

        let created_group = client
            .create(
                "Test Group".to_string(),
                team_id,
                "A Test Group".to_string(),
                None,
                vec![RoleId { id: 1 }],
                vec![],
                None,
                None,
            )
            .await
            .unwrap();

        let group_id = created_group.group_id.unwrap();

        let retrieved_group = client.get(group_id.clone()).await.unwrap();
        assert_eq!(
            group_id.id,
            retrieved_group.group.unwrap().group_id.unwrap().id
        );

        client
            .update(
                group_id.clone(),
                "Updated Test Group".to_string(),
                "A Test Group".to_string(),
                None,
                None,
                None,
                None,
                None,
            )
            .await
            .unwrap();

        client.delete(group_id).await.unwrap();
    }
}
