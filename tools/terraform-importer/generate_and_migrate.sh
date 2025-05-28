#!/bin/bash

# Copyright 2024 Coralogix Ltd.
# Licensed under the Apache License, Version 2.0

set -e

# Colors
INFO="$(tput setaf 4)[INFO]$(tput sgr0)"
WARNING="$(tput setaf 3)[WARNING]$(tput sgr0)"
ERROR="$(tput setaf 1)[ERROR]$(tput sgr0)"
RESET="$(tput sgr0)"

# Logger
log() {
  local level=$1; shift
  echo "$(date +'%Y-%m-%d %H:%M:%S') ${level} $@"
}

# Log Colorization
colorize_logs() {
  while IFS= read -r line; do
    case "$line" in
      *INFO*) echo -e "$(date +'%Y-%m-%d %H:%M:%S') ${INFO} ${line}" ;;
      *WARNING*) echo -e "$(date +'%Y-%m-%d %H:%M:%S') ${WARNING} ${line}" ;;
      *ERROR*) echo -e "$(date +'%Y-%m-%d %H:%M:%S') ${ERROR} ${line}" ;;
      *) echo -e "$(date +'%Y-%m-%d %H:%M:%S') ${RESET} ${line}" ;;
    esac
  done
}

SCRIPT_DIR=$(pwd)
CLEANED_JSON_FILE="cleaned_config.json"

log "$INFO" "Select the migration type:"
log "$INFO" "1) Migrate from folder with terraform.tfstate"
log "$INFO" "2) Migrate from a specific resource name"

read -rp "Enter your choice (1 or 2): " CHOICE

if [ "$CHOICE" = "1" ]; then
  read -rp "Enter path to the folder with terraform.tfstate: " INPUT
  if [ ! -f "$INPUT/terraform.tfstate" ]; then
    log "$ERROR" "No terraform.tfstate found at: $INPUT"
    exit 1
  fi
  MIGRATION_FOLDER="${INPUT}_migration"
  GENERATE_FLAG="-folder"
elif [ "$CHOICE" = "2" ]; then
  log "$INFO" "Choose a resource type:"
  RESOURCES="alert archive_logs archive_metrics archive_retentions custom_role dashboard dashboards_folder events2metrics group recording_rules_groups_set scope tco_policies_logs tco_policies_traces webhook"
  select RESOURCE in $RESOURCES; do
    if [ -n "$RESOURCE" ]; then
      INPUT="$RESOURCE"
      MIGRATION_FOLDER="./${RESOURCE}_migration"
      GENERATE_FLAG="-type"
      break
    else
      log "$WARNING" "Invalid selection. Try again."
    fi
  done
else
  log "$ERROR" "Invalid choice. Exiting."
  exit 1
fi

read -rp "Enter provider version (default: >=2.0.0): " PROVIDER_VERSION
PROVIDER_VERSION="${PROVIDER_VERSION:-">=2.0.0"}"

log "$INFO" "Creating migration folder: $MIGRATION_FOLDER"
mkdir -p "$MIGRATION_FOLDER"

log "$INFO" "Generating imports.tf..."
go run -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn" generate_imports.go "$GENERATE_FLAG=$INPUT" -output="$MIGRATION_FOLDER/imports.tf" || {
  log "$ERROR" "Failed to run generate_imports.go"
  exit 1
}

log "$INFO" "Writing provider.tf..."
cat <<EOF > "$MIGRATION_FOLDER/provider.tf"
terraform {
  required_providers {
    coralogix = {
      version = "$PROVIDER_VERSION"
      source  = "coralogix/coralogix"
    }
  }
}

provider "coralogix" {
  #api_key = "<use CORALOGIX_API_KEY env>"
  #env     = "<use CORALOGIX_ENV env>"
}
EOF

cd "$MIGRATION_FOLDER"

log "$INFO" "Running terraform init..."
terraform init 2>&1 | colorize_logs || {
  log "$ERROR" "Terraform init failed"
  exit 1
}

log "$INFO" "Running terraform plan..."
terraform plan -generate-config-out=generated.tf 2>&1 | colorize_logs || {
  log "$ERROR" "Terraform plan failed"
  exit 1
}

log "$INFO" "Converting generated.tf to JSON..."
hcl2json < generated.tf > config.json || {
  log "$ERROR" "hcl2json conversion failed"
  exit 1
}

log "$INFO" "Cleaning JSON with Python..."
python3 <<EOF
import json

def remove_nulls(obj):
    if isinstance(obj, dict):
        return {k: remove_nulls(v) for k, v in obj.items() if v is not None}
    elif isinstance(obj, list):
        return [remove_nulls(v) for v in obj if v is not None]
    return obj

with open("config.json") as f:
    data = json.load(f)

with open("$CLEANED_JSON_FILE", "w") as f:
    json.dump(remove_nulls(data), f, indent=2)
EOF

cd "$SCRIPT_DIR"
log "$INFO" "Converting cleaned JSON to HCL..."
go run json_to_hcl.go "$MIGRATION_FOLDER/$CLEANED_JSON_FILE" "$MIGRATION_FOLDER/cleaned_config.tf" || {
  log "$ERROR" "Failed to convert JSON to HCL"
  exit 1
}

mv "$MIGRATION_FOLDER/cleaned_config.tf" "$MIGRATION_FOLDER/generated.tf"

cd "$MIGRATION_FOLDER"

read -rp "Do you want to apply the Terraform changes now? (yes/no): " APPLY_RESPONSE
if [ "$APPLY_RESPONSE" = "yes" ]; then
  APPLY_TERRAFORM="yes"
else
  APPLY_TERRAFORM="no"
fi

if [ "$APPLY_TERRAFORM" = "yes" ]; then
  log "$INFO" "Running terraform apply..."
  terraform apply 2>&1 | tee tf_apply.log | colorize_logs
  exit_code=${PIPESTATUS[0]}
  if [ "$exit_code" -ne 0 ]; then
    log "$ERROR" "Terraform apply failed with exit code $exit_code"
    exit "$exit_code"
  fi
  log "$INFO" "Terraform apply completed."
else
  log "$WARNING" "Skipping terraform apply (APPLY_TERRAFORM != yes)"
fi

log "$INFO" "Cleaning up..."
rm -f imports.tf config.json "$CLEANED_JSON_FILE"
log "$INFO" "Done. Migration folder: $MIGRATION_FOLDER"