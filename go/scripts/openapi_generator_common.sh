#!/bin/bash

run_openapi_generator() {
  if command -v openapi-generator-cli >/dev/null 2>&1; then
    openapi-generator-cli "$@"
    return
  fi

  if command -v openapi-generator >/dev/null 2>&1; then
    openapi-generator "$@"
    return
  fi

  if command -v npx >/dev/null 2>&1 && java -version >/dev/null 2>&1; then
    npx --yes @openapitools/openapi-generator-cli "$@"
    return
  fi

  if command -v docker >/dev/null 2>&1 && docker info >/dev/null 2>&1; then
    local generator_version
    generator_version="$(python3 - "${OPENAPI_TOOLS_CONFIG:?OPENAPI_TOOLS_CONFIG must be set}" <<'PY'
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

normalize_regex_validator_tags() {
  local outdir="$1"

  # OpenAPI Generator over-escapes \s and \S in Go struct tags. Four
  # backslashes in the generated source become two backslashes after
  # reflect.StructTag parsing, which changes the regexp semantics.
  find "$outdir" -type f -name '*.go' -exec perl -0pi -e '1 while s/(validate:"regexp=[^"`]*)\\{4}([sS])/$1\\\\$2/g' {} +
}
