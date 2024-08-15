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
    use cx_sdk::{
        auth::AuthContext,
        client::archive_logs::{LogsArchiveClient, S3TargetSpec, TargetSpec, TargetSpecValidation},
        CoralogixRegion,
    };

    #[tokio::test]
    async fn test_logs_archive() {
        let service = LogsArchiveClient::new(
            CoralogixRegion::from_env().unwrap(),
            AuthContext::from_env(),
        )
        .unwrap();

        let target_spec = S3TargetSpec {
            bucket: "yak-2-bucket".to_string(),
            region: Some("eu-north-1".to_string()),
        };
        service
            .validate_target(true, TargetSpecValidation::S3(target_spec.clone()))
            .await
            .unwrap();

        service
            .set_target(true, TargetSpec::S3(target_spec))
            .await
            .unwrap();

        let _ = service.get_target().await.unwrap();
    }
}
