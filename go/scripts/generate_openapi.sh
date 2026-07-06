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

  if ! run_openapi_generator generate \
    -i "$codegen_spec" \
    -g go \
    -o "$outdir" \
    --template-dir="$TEMPLATE_DIR" \
    --additional-properties=withGoMod=false,packageName="$name",enumClassPrefix=true,disallowAdditionalPropertiesIfNotPresent=false \
    --global-property=apiTests=false,modelTests=false,apiDocs=false,modelDocs=false; then
      echo "FAILED to generate for $filename"
      continue
  fi

  normalize_regex_validator_tags "$outdir"

  # The sanitized spec is only a workaround for Go code generation. Keep the
  # package-local OpenAPI artifact identical to the canonical split spec so
  # downstream consumers still see the full validation constraints.
  cp "$spec" "$outdir/api/openapi.yaml"
done
