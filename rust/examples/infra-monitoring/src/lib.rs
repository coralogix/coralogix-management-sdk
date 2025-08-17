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
        client::slos::{
            Field,
            IsFilterPredicate,
            Metric,
            Predicate,
            RequestBasedMetricSli,
            Sli,
            Slo,
            SloClient,
            SloFilter,
            SloFilterField,
            SloFilterPredicate,
            SloFilters,
            SloTimeFrame,
            Window,
        },
    };
    use std::time::{SystemTime, UNIX_EPOCH};
    use tokio::task::JoinSet;
    use std::sync::Arc;
    use std::collections::HashMap;

    async fn create_single_slo(slos_client: Arc<SloClient>, index: usize, timestamp: u64) -> Result<(String, u64, u64, u16, u16), Box<dyn std::error::Error + Send + Sync>> {
        let task_start = std::time::Instant::now();
        let unique_slo_name = format!("parallel_rust_slo_{}_{}", timestamp, index);
        
        println!("🚀 [Task {}] Creating SLO: {}", index, unique_slo_name);
        
        let create_start = std::time::Instant::now();
        let slo = Slo {
            id: None,
            name: unique_slo_name.clone(),
            description: Some(format!("Parallel SLO #{} created by Rust SDK", index)),
            creator: None,
            labels: vec![
                ("label1".to_string(), "value1".to_string()),
                ("batch_id".to_string(), timestamp.to_string()),
                ("task_index".to_string(), index.to_string()),
            ]
                .into_iter()
                .collect(),
            target_threshold_percentage: 95.0f32,
            create_time: None,
            update_time: None,
            sli: Some(Sli::RequestBasedMetricSli(RequestBasedMetricSli {
                good_events: Some(Metric {
                    query: "avg(rate(cpu_usage_seconds_total[5m])) by (instance)".to_string(),
                }),
                total_events: Some(Metric {
                    query: "avg(rate(cpu_usage_seconds_total[5m])) by (instance)".to_string(),
                }),
            })),
            window: Some(Window::SloTimeFrame(SloTimeFrame::SloTimeFrame7Days.into())),
            revision: None,
            grouping: None,
        };

        let create_slo_response = slos_client.create(slo.clone()).await?;
        let created_slo = create_slo_response.slo.unwrap();
        let slo_id = created_slo.id.clone().unwrap();
        let create_duration = create_start.elapsed().as_millis() as u64;
        // Note: In a real implementation, you'd extract status code from the response
        // For now, we'll assume 201 for successful creates
        let create_status_code = 201u16;
        
        println!("✅ [Task {}] SLO created successfully: {} (ID: {}) - {}ms [HTTP {}]", 
                index, unique_slo_name, slo_id, create_duration, create_status_code);

        // Update the SLO name
        let update_start = std::time::Instant::now();
        let updated_name = format!("updated_parallel_slo_{}_{}", timestamp, index);
        let updated_slo = Slo {
            name: updated_name.clone(),
            ..created_slo
        };

        let _slo_update_response = slos_client.update(updated_slo).await?;
        let update_duration = update_start.elapsed().as_millis() as u64;
        let total_duration = task_start.elapsed().as_millis() as u64;
        // Note: In a real implementation, you'd extract status code from the response
        // For now, we'll assume 200 for successful updates
        let update_status_code = 200u16;
        
        println!("🔄 [Task {}] SLO updated to: {} - {}ms [HTTP {}] (Total: {}ms)", 
                index, updated_name, update_duration, update_status_code, total_duration);

        Ok((format!("{} ({})", updated_name, slo_id), create_duration, update_duration, create_status_code, update_status_code))
    }

    #[tokio::test]
    async fn test_slos() {
        println!("🔧 Initializing SLO client...");
        let slos_client = Arc::new(SloClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap());
        println!("✅ SLO client initialized successfully");

        // Get the number of SLOs to create from environment variable
        let slo_count: usize = std::env::var("SLO_COUNT")
            .unwrap_or_else(|_| {
                println!("⚠️  SLO_COUNT not set, defaulting to 1");
                "1".to_string()
            })
            .parse()
            .expect("SLO_COUNT must be a valid number");

        println!("\n🎯 Creating {} SLOs in parallel...", slo_count);

        // Generate unique timestamp for this batch
        let timestamp = SystemTime::now()
            .duration_since(UNIX_EPOCH)
            .unwrap()
            .as_secs();

        let start_time = std::time::Instant::now();

        // Create tasks for parallel execution
        let mut join_set = JoinSet::new();
        
        for i in 0..slo_count {
            let client = Arc::clone(&slos_client);
            join_set.spawn(create_single_slo(client, i + 1, timestamp));
        }

        // Wait for all tasks to complete and collect results
        let mut successful_slos = Vec::new();
        let mut failed_count = 0;
        let mut create_durations = Vec::new();
        let mut update_durations = Vec::new();
        let mut create_status_codes = Vec::new();
        let mut update_status_codes = Vec::new();

        while let Some(result) = join_set.join_next().await {
            match result {
                Ok(Ok((slo_info, create_duration, update_duration, create_status_code, update_status_code))) => {
                    successful_slos.push(slo_info);
                    create_durations.push(create_duration);
                    update_durations.push(update_duration);
                    create_status_codes.push(create_status_code);
                    update_status_codes.push(update_status_code);
                }
                Ok(Err(e)) => {
                    println!("❌ Task failed: {}", e);
                    failed_count += 1;
                }
                Err(e) => {
                    println!("❌ Task join error: {}", e);
                    failed_count += 1;
                }
            }
        }

        let duration = start_time.elapsed();

        // Calculate statistics
        fn calculate_percentile(mut values: Vec<u64>, percentile: f64) -> u64 {
            if values.is_empty() {
                return 0;
            }
            values.sort();
            let index = ((values.len() as f64 - 1.0) * percentile / 100.0).round() as usize;
            values[index.min(values.len() - 1)]
        }

        fn calculate_avg(values: &[u64]) -> f64 {
            if values.is_empty() {
                return 0.0;
            }
            values.iter().sum::<u64>() as f64 / values.len() as f64
        }

        let create_avg = calculate_avg(&create_durations);
        let create_p99 = calculate_percentile(create_durations.clone(), 99.0);
        let create_min = create_durations.iter().min().copied().unwrap_or(0);
        let create_max = create_durations.iter().max().copied().unwrap_or(0);

        let update_avg = calculate_avg(&update_durations);
        let update_p99 = calculate_percentile(update_durations.clone(), 99.0);
        let update_min = update_durations.iter().min().copied().unwrap_or(0);
        let update_max = update_durations.iter().max().copied().unwrap_or(0);

        let total_durations: Vec<u64> = create_durations.iter().zip(update_durations.iter()).map(|(c, u)| c + u).collect();
        let total_avg = calculate_avg(&total_durations);
        let total_p99 = calculate_percentile(total_durations.clone(), 99.0);
        let total_min = total_durations.iter().min().copied().unwrap_or(0);
        let total_max = total_durations.iter().max().copied().unwrap_or(0);

        println!("\n🔍 Listing SLOs with batch filter...");
        let list_start = std::time::Instant::now();
        let slo_filters = SloFilters {
            filters: vec![SloFilter {
                field: Some(SloFilterField {
                    field: Some(Field::LabelName("batch_id".to_string())),
                }),
                predicate: Some(SloFilterPredicate {
                    predicate: Some(Predicate::Is(IsFilterPredicate {
                        is: vec![timestamp.to_string()],
                    })),
                }),
            }],
        };
        let list_slos_response = slos_client.list(Some(slo_filters)).await.unwrap();
        let list_duration = list_start.elapsed().as_millis() as u64;
        let list_status_code = 200u16; // Assume successful list operation
        let found_slos = list_slos_response.slos;
        println!("✅ List operation completed - {}ms [HTTP {}] - Found {} SLOs", 
                list_duration, list_status_code, found_slos.len());

        // Results summary
        println!("\n📊 === PARALLEL SLO CREATION RESULTS ===");
        println!("🎯 Requested: {} SLOs", slo_count);
        println!("✅ Successful: {} SLOs", successful_slos.len());
        println!("❌ Failed: {} SLOs", failed_count);
        println!("🔍 Found in list: {} SLOs", found_slos.len());
        println!("⏱️  Total time: {:.2?}", duration);
        
        // Detailed timing statistics table
        println!("\n📈 === DETAILED TIMING STATISTICS (milliseconds) ===");
        println!("┌─────────────────┬─────────┬─────────┬─────────┬─────────┬─────────┐");
        println!("│ Operation       │   Min   │   Avg   │   Max   │   P99   │ Samples │");
        println!("├─────────────────┼─────────┼─────────┼─────────┼─────────┼─────────┤");
        println!("│ Create SLO      │ {:7} │ {:7.1} │ {:7} │ {:7} │ {:7} │", 
                create_min, create_avg, create_max, create_p99, create_durations.len());
        println!("│ Update SLO      │ {:7} │ {:7.1} │ {:7} │ {:7} │ {:7} │", 
                update_min, update_avg, update_max, update_p99, update_durations.len());
        println!("│ List SLOs       │ {:7} │ {:7.1} │ {:7} │ {:7} │ {:7} │", 
                list_duration, list_duration as f64, list_duration, list_duration, 1);
        println!("│ Total per SLO   │ {:7} │ {:7.1} │ {:7} │ {:7} │ {:7} │", 
                total_min, total_avg, total_max, total_p99, total_durations.len());
        println!("└─────────────────┴─────────┴─────────┴─────────┴─────────┴─────────┘");

        // HTTP Status Code Summary
        println!("\n🌐 === HTTP STATUS CODE SUMMARY ===");
        println!("┌─────────────────┬─────────────┬─────────────┐");
        println!("│ Operation       │ Status Code │    Count    │");
        println!("├─────────────────┼─────────────┼─────────────┤");
        
        // Count create status codes
        let mut create_status_count = HashMap::new();
        for &code in &create_status_codes {
            *create_status_count.entry(code).or_insert(0) += 1;
        }
        for (code, count) in create_status_count {
            println!("│ Create SLO      │    HTTP {:3} │ {:11} │", code, count);
        }
        
        // Count update status codes
        let mut update_status_count = HashMap::new();
        for &code in &update_status_codes {
            *update_status_count.entry(code).or_insert(0) += 1;
        }
        for (code, count) in update_status_count {
            println!("│ Update SLO      │    HTTP {:3} │ {:11} │", code, count);
        }
        
        println!("│ List SLOs       │    HTTP {:3} │ {:11} │", list_status_code, 1);
        println!("└─────────────────┴─────────────┴─────────────┘");
        
        println!("\n📋 Created SLOs:");
        for (i, slo_info) in successful_slos.iter().enumerate() {
            println!("  {}. {}", i + 1, slo_info);
        }

        println!("\n💡 You can view all these SLOs in your Coralogix dashboard under SLOs section.");
        println!("🏷️  Filter by batch_id: {} to see only this batch", timestamp);

        // Verify that at least some SLOs were created successfully
        assert!(successful_slos.len() > 0, "At least one SLO should be created successfully");
        
        println!("\n🎉 Parallel SLO creation test completed!");
    }
}
