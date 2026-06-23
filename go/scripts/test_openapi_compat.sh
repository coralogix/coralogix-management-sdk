#!/bin/bash

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/../.." && pwd)"
FIXTURES_DIR="$REPO_ROOT/go/openapi/testdata/compat"
WORK_DIR="$REPO_ROOT/go/openapi/testdata/.compat-tmp"
TEMPLATE_DIR="$REPO_ROOT/go/openapi/templates"
OPENAPI_TOOLS_CONFIG="$REPO_ROOT/openapitools.json"

source "$SCRIPT_DIR/openapi_generator_common.sh"

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

  normalize_regex_validator_tags "$outdir"

  cp "$fixture_dir/compat_test.go" "$outdir/compat_test.go"

  # Models with `additionalProperties: false` + required fields can emit a
  # duplicate `"bytes"` import. goimports cleans this up so each fixture
  # compiles. Install on demand if the binary isn't already available.
  if ! command -v goimports >/dev/null 2>&1; then
    GOPATH_BIN="$(go env GOPATH)/bin"
    if [ ! -x "$GOPATH_BIN/goimports" ]; then
      go install golang.org/x/tools/cmd/goimports@latest >/dev/null 2>&1 || true
    fi
    if [ -x "$GOPATH_BIN/goimports" ]; then
      "$GOPATH_BIN/goimports" -w "$outdir"
    fi
  else
    goimports -w "$outdir"
  fi

  (cd "$REPO_ROOT" && go test "./${outdir#$REPO_ROOT/}")
done
