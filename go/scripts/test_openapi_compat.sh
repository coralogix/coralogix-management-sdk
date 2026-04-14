#!/bin/bash

set -euo pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
FIXTURES_DIR="$REPO_ROOT/go/openapi/testdata/compat"
WORK_DIR="$REPO_ROOT/go/openapi/testdata/.compat-tmp"
TEMPLATE_DIR="$REPO_ROOT/go/openapi/templates"
OPENAPI_TOOLS_CONFIG="$REPO_ROOT/openapitools.json"

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

rm -rf "$WORK_DIR"
mkdir -p "$WORK_DIR"

cleanup() {
  rm -rf "$WORK_DIR"
}
trap cleanup EXIT

for fixture_dir in "$FIXTURES_DIR"/*; do
  [ -d "$fixture_dir" ] || continue

  fixture_name="$(basename "$fixture_dir")"
  package_name="${fixture_name//-/_}"
  outdir="$WORK_DIR/$fixture_name"
  fixture_rel="${fixture_dir#$REPO_ROOT/}"
  outdir_rel="${outdir#$REPO_ROOT/}"
  template_rel="${TEMPLATE_DIR#$REPO_ROOT/}"

  run_openapi_generator generate \
    -i "$fixture_rel/openapi.yaml" \
    -g go \
    -o "$outdir_rel" \
    --template-dir="$template_rel" \
    --additional-properties=withGoMod=false,packageName="$package_name",enumClassPrefix=true,disallowAdditionalPropertiesIfNotPresent=false \
    --global-property=apiTests=false,modelTests=false,apiDocs=false,modelDocs=false >/dev/null

  cp "$fixture_dir/compat_test.go" "$outdir/compat_test.go"
  (cd "$REPO_ROOT" && go test "./${outdir#$REPO_ROOT/}")
done
