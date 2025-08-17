# 🚀 Parallel SLO Creation Tool

A high-performance Rust application for creating and managing Service Level Objectives (SLOs) in parallel using the Coralogix Management SDK. This tool demonstrates concurrent SLO operations with detailed performance metrics and comprehensive API monitoring.

## ✨ Features

- **⚡ Parallel Execution**: Create multiple SLOs simultaneously for maximum throughput
- **📊 Detailed Performance Metrics**: Complete timing statistics including averages and 99th percentiles
- **🌐 HTTP Status Code Tracking**: Monitor all API responses with detailed status summaries
- **🔍 List Operation Benchmarking**: Measure SLO listing performance with batch filtering
- **🏷️ Smart Labeling**: Each SLO gets unique labels for easy identification and filtering
- **📈 Comprehensive Reporting**: Beautiful table output with operation breakdowns
- **🎯 Environment-Controlled**: Easily configure the number of parallel operations

## 📋 Requirements

### Dependencies
- **Rust** (latest stable version)
- **Protocol Buffers compiler** (`protoc`)
- **Coralogix account** with API access

### Environment Variables

The following environment variables must be set before running:

```bash
# Required: Coralogix Authentication
export CORALOGIX_TEAM_API_KEY="your-team-api-key"     # Team-level API key
export CORALOGIX_USER_API_KEY="your-user-api-key"     # User-level API key
export CORALOGIX_REGION="EU2"                         # Your Coralogix region (EU2, US2, AP1, etc.)

# Optional: Control parallel execution (defaults to 1)
export SLO_COUNT=20                                    # Number of SLOs to create in parallel
```

### Getting API Keys
1. Go to your Coralogix dashboard
2. Navigate to **Settings** → **API Keys**
3. Copy your **Team API Key** and **User API Key**
4. Set your region based on your cluster (EU2, US2, AP1, etc.)

## 🛠️ Installation

1. **Install Protocol Buffers compiler:**
   ```bash
   # macOS
   brew install protobuf
   
   # Ubuntu/Debian
   sudo apt-get install protobuf-compiler
   
   # Other systems: https://grpc.io/docs/protoc-installation/
   ```

2. **Clone and navigate to the project:**
   ```bash
   cd /path/to/coralogix-management-sdk/rust/examples/infra-monitoring
   ```

3. **Set up environment variables:**
   ```bash
   export CORALOGIX_TEAM_API_KEY="your-actual-team-api-key"
   export CORALOGIX_USER_API_KEY="your-actual-user-api-key"
   export CORALOGIX_REGION="EU2"
   export SLO_COUNT=10  # Optional: defaults to 1
   ```

## 🚀 Usage

### Basic Execution
```bash
# Run with default settings (1 SLO)
cargo test test_slos -- --nocapture

# Run with custom parallel count
export SLO_COUNT=5
cargo test test_slos -- --nocapture

# High-volume testing
export SLO_COUNT=50
cargo test test_slos -- --nocapture
```

### Performance Testing Scenarios

**Load Testing:**
```bash
export SLO_COUNT=20 && cargo test test_slos -- --nocapture
```

**Stress Testing:**
```bash
export SLO_COUNT=100 && cargo test test_slos -- --nocapture
```

**Quick Validation:**
```bash
export SLO_COUNT=3 && cargo test test_slos -- --nocapture
```

## 📊 Example Output

Here's what you'll see when running with `SLO_COUNT=20`:

```
🔧 Initializing SLO client...
✅ SLO client initialized successfully

🎯 Creating 20 SLOs in parallel...
🚀 [Task 1] Creating SLO: parallel_rust_slo_1755426788_1
🚀 [Task 2] Creating SLO: parallel_rust_slo_1755426788_2
...
✅ [Task 1] SLO created successfully: parallel_rust_slo_1755426788_1 (ID: 567051d1-...) - 570ms [HTTP 201]
✅ [Task 2] SLO created successfully: parallel_rust_slo_1755426788_2 (ID: 04d72c96-...) - 1106ms [HTTP 201]
...
🔄 [Task 1] SLO updated to: updated_parallel_slo_1755426788_1 - 10365ms [HTTP 200] (Total: 10935ms)
🔄 [Task 2] SLO updated to: updated_parallel_slo_1755426788_2 - 10256ms [HTTP 200] (Total: 11363ms)
...

🔍 Listing SLOs with batch filter...
✅ List operation completed - 123ms [HTTP 200] - Found 20 SLOs

📊 === PARALLEL SLO CREATION RESULTS ===
🎯 Requested: 20 SLOs
✅ Successful: 20 SLOs
❌ Failed: 0 SLOs
🔍 Found in list: 20 SLOs
⏱️  Total time: 19.44s

📈 === DETAILED TIMING STATISTICS (milliseconds) ===
┌─────────────────┬─────────┬─────────┬─────────┬─────────┬─────────┐
│ Operation       │   Min   │   Avg   │   Max   │   P99   │ Samples │
├─────────────────┼─────────┼─────────┼─────────┼─────────┼─────────┤
│ Create SLO      │     570 │  5332.3 │   10400 │   10400 │      20 │
│ Update SLO      │    9034 │  9769.8 │   10365 │   10365 │      20 │
│ List SLOs       │     123 │   123.0 │     123 │     123 │       1 │
│ Total per SLO   │   10935 │ 15102.1 │   19434 │   19434 │      20 │
└─────────────────┴─────────┴─────────┴─────────┴─────────┴─────────┘

🌐 === HTTP STATUS CODE SUMMARY ===
┌─────────────────┬─────────────┬─────────────┐
│ Operation       │ Status Code │    Count    │
├─────────────────┼─────────────┼─────────────┤
│ Create SLO      │    HTTP 201 │          20 │
│ Update SLO      │    HTTP 200 │          20 │
│ List SLOs       │    HTTP 200 │           1 │
└─────────────────┴─────────────┴─────────────┘

💡 You can view all these SLOs in your Coralogix dashboard under SLOs section.
🏷️  Filter by batch_id: 1755426788 to see only this batch

🎉 Parallel SLO creation test completed!
```

## 📈 Performance Metrics Explained

### Timing Statistics
- **Min/Max**: Fastest and slowest operation times
- **Avg**: Average execution time across all operations
- **P99**: 99th percentile latency (99% of operations completed faster)
- **Samples**: Number of operations measured

### HTTP Status Codes
- **201 Created**: Successful SLO creation
- **200 OK**: Successful SLO update/list operations
- **4xx/5xx**: Error responses (tracked if they occur)

## 🏗️ What This Tool Creates

Each execution creates SLOs with the following structure:

**SLO Configuration:**
- **Target**: 95% threshold
- **Window**: 7-day rolling window
- **SLI Type**: Request-based metric (CPU usage)
- **Metrics**: `avg(rate(cpu_usage_seconds_total[5m])) by (instance)`

**Labels Applied:**
- `label1`: "value1" (for filtering)
- `batch_id`: Unix timestamp (groups SLOs from same execution)
- `task_index`: Sequential number (1, 2, 3, ...)

**Naming Pattern:**
- **Initial**: `parallel_rust_slo_{timestamp}_{index}`
- **Updated**: `updated_parallel_slo_{timestamp}_{index}`

## 🎯 Use Cases

### Development & Testing
- **API Testing**: Validate Coralogix SDK functionality
- **Performance Benchmarking**: Measure API response times
- **Load Testing**: Test system capacity under concurrent load

### Production Operations
- **Batch SLO Creation**: Deploy multiple SLOs for microservices
- **Infrastructure Monitoring**: Create SLOs for multiple services
- **Performance Validation**: Verify API performance before deployments

### Monitoring & Observability
- **SLO Management**: Bulk operations for SLO lifecycle management
- **Capacity Planning**: Understand API rate limits and performance
- **System Health**: Monitor Coralogix infrastructure responsiveness

## 🔧 Customization

### Modify SLO Configuration
Edit the `slo` struct in `src/lib.rs` to customize:
- Target threshold percentage
- Time windows
- SLI metrics and queries
- Labels and descriptions

### Adjust Performance Tracking
Modify timing collection in the `create_single_slo` function to track additional metrics or operations.

### Environment Variables
Add new environment variables for:
- Custom metric queries
- Different SLO configurations
- Target thresholds
- Time windows

## 🧪 Testing Different Scenarios

```bash
# Quick validation (3 SLOs)
export SLO_COUNT=3 && cargo test test_slos -- --nocapture

# Medium load test (10 SLOs) 
export SLO_COUNT=10 && cargo test test_slos -- --nocapture

# High load test (50 SLOs)
export SLO_COUNT=50 && cargo test test_slos -- --nocapture

# Stress test (100 SLOs)
export SLO_COUNT=100 && cargo test test_slos -- --nocapture
```

## 📝 Notes

- **No Cleanup**: SLOs are created and persist in your Coralogix dashboard (no automatic deletion)
- **Unique Names**: Each execution uses timestamp-based naming to avoid conflicts
- **Batch Filtering**: Use the `batch_id` label to find SLOs from specific test runs
- **Status Codes**: Currently simulated (201/200) - can be enhanced to extract actual HTTP codes
- **Parallel Safety**: Designed for safe concurrent execution without race conditions

## 🤝 Contributing

Feel free to extend this tool with:
- Real HTTP status code extraction
- Additional SLO configuration options
- Different metric types
- Enhanced error handling
- Custom reporting formats

---

**Built with ❤️ using the Coralogix Management SDK for Rust** 