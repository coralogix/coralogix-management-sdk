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
        CoralogixRegion,
        auth::AuthContext,
        client::dataprime_query::{
            DataprimeQueryClient,
            Message,
        },
    };

    #[tokio::test]
    async fn test_dataprime_query() {
        let svc = DataprimeQueryClient::new(
            CoralogixRegion::from_env().unwrap(),
            AuthContext::from_env(),
        )
        .unwrap();
        let mut stream = svc
            .run("filter log_obj.message ~ 'Hello world'".to_string(), None)
            .await
            .unwrap();
        while let Some(response) = stream.message().await.unwrap() {
            if let Some(Message::Result(_)) = response.message {}
        }
    }
}
