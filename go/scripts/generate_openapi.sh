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

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SPECS_DIR="specs"
OUT_BASE="go/openapi/gen"
TEMPLATE_DIR="go/openapi/templates"
OPENAPI_TOOLS_CONFIG="openapitools.json"
CODEGEN_SPECS_DIR="$(mktemp -d ".openapi-codegen-specs.XXXXXX")"
trap 'rm -rf "$CODEGEN_SPECS_DIR"' EXIT

source "$SCRIPT_DIR/openapi_generator_common.sh"

generate_go_client() {
  local codegen_spec="$1"
  local outdir="$2"
  local package_name="$3"
  local global_property="$4"

  run_openapi_generator generate \
    -i "$codegen_spec" \
    -g go \
    -o "$outdir" \
    --template-dir="$TEMPLATE_DIR" \
    --additional-properties=withGoMod=false,packageName="$package_name",enumClassPrefix=true,disallowAdditionalPropertiesIfNotPresent=false \
    --global-property="$global_property"
}

generate_dashboard_support_files() {
  local codegen_spec="$1"
  local outdir="$2"

  generate_go_client \
    "$codegen_spec" \
    "$outdir" \
    "dashboard_service" \
    "apiTests=false,modelTests=false,apiDocs=false,modelDocs=false,supportingFiles=client.go"

  generate_go_client \
    "$codegen_spec" \
    "$outdir" \
    "dashboard_service" \
    "apiTests=false,modelTests=false,apiDocs=false,modelDocs=false,supportingFiles=configuration.go"

  generate_go_client \
    "$codegen_spec" \
    "$outdir" \
    "dashboard_service" \
    "apiTests=false,modelTests=false,apiDocs=false,modelDocs=false,supportingFiles=response.go"

  generate_go_client \
    "$codegen_spec" \
    "$outdir" \
    "dashboard_service" \
    "apiTests=false,modelTests=false,apiDocs=false,modelDocs=false,supportingFiles=utils.go"
}

ensure_dashboard_client_service_wiring() {
  local client_file="$1/client.go"

  if grep -q "DashboardServiceAPI \*DashboardServiceAPIService" "$client_file" && \
    grep -q "c.DashboardServiceAPI = (\*DashboardServiceAPIService)(&c.common)" "$client_file"; then
    return
  fi

  python3 - "$client_file" <<'PY'
import sys
from pathlib import Path

path = Path(sys.argv[1])
content = path.read_text()

content = content.replace(
    "\t// API Services\n}",
    "\t// API Services\n\n\tDashboardServiceAPI *DashboardServiceAPIService\n}",
    1,
)
content = content.replace(
    "\t// API Services\n\n\treturn c",
    "\t// API Services\n\tc.DashboardServiceAPI = (*DashboardServiceAPIService)(&c.common)\n\n\treturn c",
    1,
)

path.write_text(content)
PY

  if ! grep -q "DashboardServiceAPI \*DashboardServiceAPIService" "$client_file"; then
    echo "FAILED to populate dashboard APIClient service field" >&2
    exit 1
  fi

  if ! grep -q "c.DashboardServiceAPI = (\*DashboardServiceAPIService)(&c.common)" "$client_file"; then
    echo "FAILED to populate dashboard APIClient service initializer" >&2
    exit 1
  fi
}

rm -rf "$OUT_BASE"

for spec in "$SPECS_DIR"/*; do
  [ -f "$spec" ] || continue

  filename=$(basename -- "$spec")
  name="${filename%.*}"
  outdir="$OUT_BASE/$name"
  codegen_spec="$CODEGEN_SPECS_DIR/$filename"

  python3 "$SCRIPT_DIR/sanitize_openapi_for_go_codegen.py" "$spec" "$codegen_spec"

  echo ">>> Processing file: '$filename' (name='$name')"
  echo ">>> Outdir: '$outdir'"

  if ! generate_go_client "$codegen_spec" "$outdir" "$name" "apiTests=false,modelTests=false,apiDocs=false,modelDocs=false"; then
    # Dashboard is large enough that OpenAPI Generator can OOM while serializing
    # its generated api/openapi.yaml support file. The models/API are written
    # before that step, and we copy the canonical spec below anyway.
    if [ "$name" = "dashboard_service" ] && [ -f "$outdir/api_dashboard_service.go" ]; then
      echo "Dashboard models/APIs generated; continuing with separate support-file generation"
    else
      echo "FAILED to generate for $filename" >&2
      exit 1
    fi
  fi

  if [ "$name" = "dashboard_service" ]; then
    if ! generate_dashboard_support_files "$codegen_spec" "$outdir"; then
      echo "FAILED to generate dashboard support files" >&2
      exit 1
    fi
    ensure_dashboard_client_service_wiring "$outdir"
  fi

  normalize_regex_validator_tags "$outdir"

  # The sanitized spec is only a workaround for Go code generation. Keep the
  # package-local OpenAPI artifact identical to the canonical split spec so
  # downstream consumers still see the full validation constraints.
  cp "$spec" "$outdir/api/openapi.yaml"
done
