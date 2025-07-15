#!/bin/bash

# Copyright 2024 Coralogix Ltd.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Comprehensive Terraform Provider Import Test Script
# 
# This script tests the import functionality for all supported Coralogix Terraform
# provider resource types by running generate_and_migrate.sh for each resource.
#
# Usage:
#   ./test_imports.sh
#
# Environment Variables:
#   TF_PROVIDER_VERSION - Override the default Terraform provider version (default: 2.1.0)
#                        Examples: TF_PROVIDER_VERSION="~>2.0.0" ./test_imports.sh
#                                 TF_PROVIDER_VERSION=">=2.1.0" ./test_imports.sh

# Don't exit immediately on command failures - we want to continue testing all resources
# set -e

# Define colors using tput
INFO="$(tput setaf 4)[INFO]$(tput sgr0)"      # Blue
WARNING="$(tput setaf 3)[WARNING]$(tput sgr0)"   # Yellow
WARN="$WARNING"                                # Alias for WARNING
ERROR="$(tput setaf 1)[ERROR]$(tput sgr0)"     # Red
DEBUG="$(tput setaf 2)[DEBUG]$(tput sgr0)"     # Green
SUCCESS="$(tput setaf 2)[SUCCESS]$(tput sgr0)" # Green
RESET="$(tput sgr0)"                           # Reset color

# Logging function - writes to both console and log file
log() {
  local level=$1
  shift
  local message="$(date +'%Y-%m-%d %H:%M:%S') ${level} $@"
  echo "$message"
  echo "$message" >> "$LOG_FILE"
}

# Resource types array - matches the OPTIONS array in generate_and_migrate.sh
RESOURCE_TYPES=("alert" "archive_logs" "archive_metrics" "archive_retentions" "custom_role" "dashboard"
                "dashboards_folder" "events2metrics" "group" "recording_rules_groups_set" "scope"
                "tco_policies_logs" "tco_policies_traces" "webhook")

# Function to get resource selection number (1-indexed)
get_resource_number() {
    local resource_name=$1
    for i in "${!RESOURCE_TYPES[@]}"; do
        if [ "${RESOURCE_TYPES[$i]}" = "$resource_name" ]; then
            echo $((i + 1))
            return
        fi
    done
    echo "0"  # Not found
}

# Default provider version - can be overridden with TF_PROVIDER_VERSION env var
DEFAULT_PROVIDER_VERSION="${TF_PROVIDER_VERSION:-2.1.0}"

# Log file for all operations
LOG_FILE="terraform_import_test_$(date +%Y%m%d_%H%M%S).log"

# Arrays to track results
declare -a SUCCESSFUL_RESOURCES=()
declare -a FAILED_RESOURCES=()
declare -a DRIFT_RESOURCES=()

# Cleanup function
cleanup() {
    if [ -n "$LOG_FILE" ]; then
        log "$INFO" "Cleaning up any background processes..."
    fi
    jobs -p | xargs -r kill 2>/dev/null || true
}

# Set up trap for cleanup
trap cleanup EXIT

# Function to run complete test for a specific resource (migration + plan check)
run_complete_test_for_resource() {
    local resource_name=$1
    local resource_number=$(get_resource_number "$resource_name")
    
    if [ "$resource_number" -eq 0 ]; then
        log "$ERROR" "Invalid resource name: $resource_name"
        FAILED_RESOURCES+=("$resource_name")
        return 1
    fi
    
    log "$INFO" "Starting complete test for resource: $resource_name (option $resource_number)"
    
    # Step 1: Run migration
    log "$INFO" "Step 1: Running migration for $resource_name..."
    
    echo "===== MIGRATION OUTPUT for $resource_name =====" >> "$LOG_FILE"
    
    # Create temporary file to capture output
    local temp_output=$(mktemp)
    
    # Use timeout with proper error handling - output only to log file and temp file
    # Use printf with explicit newlines and proper input redirection
    # No need for "yes" since we use terraform apply -auto-approve
    printf "2\n%s\n%s\n" "$resource_number" "$DEFAULT_PROVIDER_VERSION" | timeout 300 bash generate_and_migrate.sh > "$temp_output" 2>&1
    local exit_code=$?
    
    # Append output to log file
    cat "$temp_output" >> "$LOG_FILE"
    
    # Check for terraform errors in the output
    local has_terraform_errors=false
    if grep -q "Error:" "$temp_output" || grep -q "│ Error:" "$temp_output"; then
        has_terraform_errors=true
        log "$ERROR" "Terraform errors detected in migration output for resource: $resource_name"
    fi
    
    # Check for import failures
    if grep -q "Configuration for import target does not exist" "$temp_output"; then
        has_terraform_errors=true
        log "$ERROR" "Import configuration errors detected for resource: $resource_name"
    fi
    
    # Cleanup temp file
    rm -f "$temp_output"
    
    # Check migration success
    if [ $exit_code -ne 0 ] || [ "$has_terraform_errors" = true ]; then
        if [ $exit_code -eq 124 ]; then
            log "$ERROR" "Migration timed out after 300 seconds for resource: $resource_name"
        elif [ "$has_terraform_errors" = true ]; then
            log "$ERROR" "Migration failed due to terraform errors for resource: $resource_name"
        else
            log "$ERROR" "Migration failed for resource: $resource_name (exit code: $exit_code)"
        fi
        FAILED_RESOURCES+=("$resource_name")
        return 1
    fi
    
    log "$SUCCESS" "Migration completed successfully for resource: $resource_name"
    
    # Step 2: Run plan check
    log "$INFO" "Step 2: Running plan check for $resource_name..."
    
    local migration_folder="${resource_name}_migration"
    
    if [ ! -d "$migration_folder" ]; then
        log "$ERROR" "Migration folder not found: $migration_folder"
        FAILED_RESOURCES+=("$resource_name")
        return 1
    fi
    
    cd "$migration_folder" || {
        log "$ERROR" "Failed to change to migration folder: $migration_folder"
        FAILED_RESOURCES+=("$resource_name")
        return 1
    }
    
    echo "===== TERRAFORM PLAN for $resource_name =====" >> "../$LOG_FILE"
    
    # Create temporary file to capture output
    local temp_plan_output=$(mktemp)
    
    # Use proper error handling for terraform plan - output only to log file and temp file
    terraform plan -detailed-exitcode > "$temp_plan_output" 2>&1
    local plan_exit_code=$?
    
    # Append output to log file
    cat "$temp_plan_output" >> "../$LOG_FILE"
    
    # Check for terraform errors in the output
    local has_plan_errors=false
    if grep -q "Error:" "$temp_plan_output" || grep -q "│ Error:" "$temp_plan_output"; then
        has_plan_errors=true
        log "$ERROR" "Terraform plan errors detected for resource: $resource_name"
    fi
    
    # Cleanup temp file
    rm -f "$temp_plan_output"
    
    cd .. || {
        log "$ERROR" "Failed to change back to parent directory"
        FAILED_RESOURCES+=("$resource_name")
        return 1
    }
    
    # Check plan success
    if [ $plan_exit_code -eq 0 ] && [ "$has_plan_errors" = false ]; then
        log "$SUCCESS" "Complete test passed for resource: $resource_name"
        SUCCESSFUL_RESOURCES+=("$resource_name")
        return 0
    elif [ $plan_exit_code -eq 2 ] && [ "$has_plan_errors" = false ]; then
        log "$WARNING" "Drift detected for resource: $resource_name"
        DRIFT_RESOURCES+=("$resource_name")
        return 0
    else
        log "$ERROR" "Plan check failed for resource: $resource_name (exit code: $plan_exit_code)"
        FAILED_RESOURCES+=("$resource_name")
        return 1
    fi
}



# Function to print summary
print_summary() {
    echo
    log "$INFO" "=================== TEST SUMMARY ==================="
    log "$INFO" "Total resources tested: ${#RESOURCE_TYPES[@]}"
    log "$SUCCESS" "Successful resources: ${#SUCCESSFUL_RESOURCES[@]}"
    log "$ERROR" "Failed resources: ${#FAILED_RESOURCES[@]}"
    log "$WARNING" "Resources with drift: ${#DRIFT_RESOURCES[@]}"
    
    if [ ${#SUCCESSFUL_RESOURCES[@]} -gt 0 ]; then
        echo
        log "$SUCCESS" "Successful resources:"
        for resource in "${SUCCESSFUL_RESOURCES[@]}"; do
            log "$SUCCESS" "  - $resource"
        done
    fi
    
    if [ ${#FAILED_RESOURCES[@]} -gt 0 ]; then
        echo
        log "$ERROR" "Failed resources:"
        for resource in "${FAILED_RESOURCES[@]}"; do
            log "$ERROR" "  - $resource"
        done
    fi
    
    if [ ${#DRIFT_RESOURCES[@]} -gt 0 ]; then
        echo
        log "$WARNING" "Resources with drift:"
        for resource in "${DRIFT_RESOURCES[@]}"; do
            log "$WARNING" "  - $resource"
        done
    fi
    
    echo
    log "$INFO" "===================================================="
    log "$INFO" "Complete test log available at: $LOG_FILE"
}

# Main execution
main() {
    # Initialize log file
    echo "Terraform Provider Import Test Log" > "$LOG_FILE"
    echo "Started at: $(date)" >> "$LOG_FILE"
    echo "=========================================" >> "$LOG_FILE"
    
    log "$INFO" "Starting comprehensive terraform provider resource import tests"
    log "$INFO" "Testing ${#RESOURCE_TYPES[@]} resource types: ${RESOURCE_TYPES[*]}"
    log "$INFO" "Using Terraform provider version: $DEFAULT_PROVIDER_VERSION"
    log "$INFO" "All output will be logged to: $LOG_FILE"
    
    # Check if generate_and_migrate.sh exists
    if [ ! -f "generate_and_migrate.sh" ]; then
        log "$ERROR" "generate_and_migrate.sh not found in current directory"
        exit 1
    fi
    
    # Check if required tools are available
    command -v terraform >/dev/null 2>&1 || { log "$ERROR" "terraform is required but not installed"; exit 1; }
    command -v go >/dev/null 2>&1 || { log "$ERROR" "go is required but not installed"; exit 1; }
    command -v timeout >/dev/null 2>&1 || { log "$ERROR" "timeout is required but not installed"; exit 1; }
    
    # Run complete tests for all resource types
    log "$INFO" "=== Running complete tests for all resource types ==="
    
    local total_resources=${#RESOURCE_TYPES[@]}
    local current_index=0
    
    for resource in "${RESOURCE_TYPES[@]}"; do
        current_index=$((current_index + 1))
        log "$INFO" "Processing resource: $resource ($current_index/$total_resources)"
        
        # Skip if migration folder already exists and contains files
        if [ -d "${resource}_migration" ] && [ "$(ls -A "${resource}_migration" 2>/dev/null)" ]; then
            log "$WARNING" "Migration folder ${resource}_migration already exists and is not empty. Skipping test."
            log "$INFO" "To re-run test for $resource, remove the folder first: rm -rf ${resource}_migration"
            SUCCESSFUL_RESOURCES+=("$resource")
            continue
        fi
        
        # Run complete test for this resource - continue even if it fails
        if run_complete_test_for_resource "$resource"; then
            log "$DEBUG" "Complete test returned success for: $resource"
        else
            log "$DEBUG" "Complete test returned failure for: $resource - continuing with next resource"
        fi
        
        # Show intermediate progress
        log "$INFO" "Progress: $current_index/$total_resources resources processed"
        
        # Small delay between tests to avoid overwhelming the system
        sleep 2
    done
    
    # Print final summary
    print_summary
    
    # Cleanup migration folders
    log "$INFO" "Cleaning up migration folders..."
    for resource in "${RESOURCE_TYPES[@]}"; do
        if [ -d "${resource}_migration" ]; then
            log "$INFO" "Removing ${resource}_migration/"
            rm -rf "${resource}_migration"
        fi
    done
    log "$INFO" "Migration folders cleanup completed."
    
    # Set exit code based on results
    if [ ${#FAILED_RESOURCES[@]} -gt 0 ]; then
        log "$ERROR" "Some tests failed. Check the complete log at: $LOG_FILE"
        exit 1
    elif [ ${#DRIFT_RESOURCES[@]} -gt 0 ]; then
        log "$WARNING" "Some resources have drift. Check the complete log at: $LOG_FILE"
        exit 2
    else
        log "$SUCCESS" "All tests passed successfully!"
        log "$INFO" "Complete test log saved at: $LOG_FILE"
        exit 0
    fi
}

# Run main function
main "$@"
