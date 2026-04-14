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

SPECS_DIR="specs"
OUT_BASE="go/openapi/gen"
TEMPLATE_DIR="go/openapi/templates"
OPENAPI_TOOLS_CONFIG="openapitools.json"

run_openapi_generator() {
  if command -v openapi-generator-cli >/dev/null 2>&1; then
    openapi-generator-cli "$@"
    return
  fi

  if java -version >/dev/null 2>&1; then
    npx --yes @openapitools/openapi-generator-cli "$@"
    return
  fi

  if command -v docker >/dev/null 2>&1 && docker info >/dev/null 2>&1; then
    local generator_version
    generator_version="$(python3 - "$OPENAPI_TOOLS_CONFIG" <<'PY'
import json
import sys

config = json.load(open(sys.argv[1]))
print(config["generator-cli"]["version"])
PY
)"
    docker run --rm \
      -v "$PWD:/local" \
      -w /local \
      "openapitools/openapi-generator-cli:v${generator_version}" \
      "$@"
    return
  fi

  echo "openapi-generator-cli requires either a local CLI, a working Java runtime, or Docker" >&2
  return 127
}

rm -rf "$OUT_BASE"

for spec in "$SPECS_DIR"/*; do
  [ -f "$spec" ] || continue

  filename=$(basename -- "$spec")
  name="${filename%.*}"
  outdir="$OUT_BASE/$name"

  echo ">>> Processing file: '$filename' (name='$name')"
  echo ">>> Outdir: '$outdir'"

  if ! run_openapi_generator generate \
    -i "$spec" \
    -g go \
    -o "$outdir" \
    --template-dir="$TEMPLATE_DIR" \
    --additional-properties=withGoMod=false,packageName="$name",enumClassPrefix=true, disallowAdditionalPropertiesIfNotPresent=false \
    --global-property=apiTests=false,modelTests=false,apiDocs=false,modelDocs=false; then
      echo "FAILED to generate for $filename"
      continue
  fi
done
