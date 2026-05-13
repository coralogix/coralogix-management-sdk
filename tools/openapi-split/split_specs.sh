#!/bin/bash
set -e

# Get absolute paths
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/../.." && pwd)"

cd "$SCRIPT_DIR"

# Copy OpenAPI spec from repo root
cp "$REPO_ROOT/openapi.yaml" .

# Run the Rust splitter (local Cargo.toml)
cargo run --manifest-path "$SCRIPT_DIR/Cargo.toml"

SPEC_DIR="specs"

if [ ! -d "$SPEC_DIR" ]; then
    echo "Error: Directory '$SPEC_DIR' not found."
    echo "Please create the directory and place your files inside it."
    exit 1
fi

echo "Starting batch processing for files in '$SPEC_DIR/'..."
echo "--------------------------------------------------------"

# Use the globally-installed `redocly` binary (assumed on PATH; CI does
# `npm install -g @redocly/cli@latest`) instead of `npx @redocly/cli@latest`.
# Bundling files in parallel via `npx` corrupts its shared npm cache and
# OOMs the runner with 46+ concurrent processes. Limit to 4 concurrent
# bundles via `xargs -P 4`.
find "$SPEC_DIR" -maxdepth 1 -type f -print0 \
    | xargs -0 -n 1 -P 4 -I {} \
        redocly bundle {} --output {} --remove-unused-components

echo "--------------------------------------------------------"
echo "Batch processing complete."

# Remove existing specs directory if it exists
if [ -d "$REPO_ROOT/specs" ]; then
    rm -rf "$REPO_ROOT/specs"
fi
mv specs "$REPO_ROOT"
