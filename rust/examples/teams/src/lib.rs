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
    use cx_sdk::{auth::AuthContext, client::teams::TeamsClient, CoralogixRegion};

    #[tokio::test]
    #[ignore = "organization key is not available"]
    async fn test_teams() {
        let client = TeamsClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();

        let team_1_creation_result = client
            .create(
                "team_1".to_string(),
                vec!["admin@coralogix.com".into(), "admin2@coralogix.com".into()],
                Some(100f64),
            )
            .await;

        assert!(team_1_creation_result.is_ok());

        let team_1_id = team_1_creation_result.unwrap().team_id.unwrap();

        let team_2_creation_result = client
            .create(
                "team_2".to_string(),
                vec!["admin@coralogix.com".into(), "admin2@coralogix.com".into()],
                Some(200f64),
            )
            .await;

        assert!(team_2_creation_result.is_ok());

        let team_2_id = team_2_creation_result.unwrap().team_id.unwrap();

        let team_1_update_result = client
            .replace(team_1_id.id, Some("team_1_updated".into()), Some(150f64))
            .await;

        assert!(team_1_update_result.is_ok());

        let team_1_get_result = client.get(team_1_id.id).await;

        assert!(team_1_get_result.is_ok());

        assert!(team_1_get_result.unwrap().team_name == "team_1_updated");

        let set_daily_quota_result = client.set_daily_quota(team_1_id.id, 250f32).await;

        assert!(set_daily_quota_result.is_ok());

        let move_quota_result = client.move_quota(team_1_id.id, team_2_id.id, 50f64).await;

        assert!(move_quota_result.is_ok());
    }
}
