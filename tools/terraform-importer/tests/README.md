# Terraform Provider Import Test

This directory contains a comprehensive test suite for validating the import functionality of all supported Coralogix Terraform provider resource types.

## Purpose

The `import_test.sh` script automates testing of the Terraform import workflow by:

1. **Testing all supported resource types** - Validates import functionality for different Coralogix resources
2. **Detecting configuration drift** - Runs `terraform plan` to ensure imported resources match their expected configuration

## Supported Resource Types

The script tests the following Coralogix Terraform provider resources:
- `alert`
- `archive_logs`
- `archive_metrics`
- `archive_retentions`
- `custom_role`
- `dashboard`
- `dashboards_folder`
- `events2metrics`
- `group`
- `recording_rules_groups_set`
- `scope`
- `tco_policies_logs`
- `tco_policies_traces`
- `webhook`

## Usage

### Basic Usage
```bash
cd tools/terraform-importer
./tests/import_test.sh
```

### Custom Provider Version
```bash
# Use a specific version
TF_PROVIDER_VERSION="2.1.0" ./import_test.sh

# Use version constraints
TF_PROVIDER_VERSION="~>2.0.0" ./import_test.sh
```

## Environment Variables

- `TF_PROVIDER_VERSION` - Override the default Terraform provider version (default: `2.1.0`)
– `CORALOGIX_API_KEY`
– `CORALOGIX_ENV`

## Output

The script provides:

1. **Real-time console output** with color-coded status messages
2. **Comprehensive log file** named `terraform_import_test_YYYYMMDD_HHMMSS.log`
3. **Final summary report** showing:
   - Total resources tested
   - Successful imports
   - Failed imports
   - Resources with configuration drift

## Test Process

For each resource type, the script runs:

1. **Migration Phase**: Runs `generate_and_migrate.sh` with the resource type
2. **Validation Phase**: Executes `terraform plan` to detect drift
3. **Result Classification**: Categorizes the outcome as success, failure, or drift
4. **Cleanup**: Removes temporary migration folders

## GitHub Actions Pipeline

The script runs in GitHub Actions workflow (test-tf-import)[../../../.github/workflows/test-tf-import.yml] on schedule, testing the latest patch version of each minor version declared in the workflow matrix.

Full test logs with all output from Terraform are saved as an artifact and can be downloaded from the workflow's page by using the link from `Upload test logs` action. This can be particularly useful for troubleshooting import errors.
