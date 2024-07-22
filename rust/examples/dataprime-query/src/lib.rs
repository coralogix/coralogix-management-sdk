#[cfg(test)]
mod tests {
    use cx_sdk::{
        auth::ApiKey,
        client::dataprime_query::{DataprimeQueryClient, Message},
        CoralogixRegion,
    };

    #[tokio::test]
    async fn test_dataprime_query() {
        let svc = DataprimeQueryClient::new(
            CoralogixRegion::from_env().unwrap(),
            ApiKey::from_env(api_key.as_str()),
        );

        let mut stream = svc
            .run("filter log_obj.message ~ 'Hello world'".to_string(), None)
            .await
            .unwrap();
        while let Some(response) = stream.message().await.unwrap() {
            if let Some(Message::Result(result)) = response.message {
                for r in result.results {
                    println!("{:?}", r.user_data);
                }
            }
        }
    }
}
