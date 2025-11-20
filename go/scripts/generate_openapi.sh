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

rm -rf "$OUT_BASE"

for spec in "$SPECS_DIR"/*; do
  [ -f "$spec" ] || continue

  filename=$(basename -- "$spec")
  name="${filename%.*}"
  outdir="$OUT_BASE/$name"

  echo ">>> Processing file: '$filename' (name='$name')"
  echo ">>> Outdir: '$outdir'"

  if ! openapi-generator-cli generate \
    -i "$spec" \
    -g go \
    -o "$outdir" \
    --additional-properties=withGoMod=false,packageName="$name",enumClassPrefix=true\
    --global-property=apiTests=false,modelTests=false,apiDocs=false,modelDocs=false; then
      echo "FAILED to generate for $filename"
      continue
  fi
done
