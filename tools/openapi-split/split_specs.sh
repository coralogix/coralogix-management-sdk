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

for file in "$SPEC_DIR"/*; do
    if [ -f "$file" ]; then
        echo "--> Found file: $file"
        npx @redocly/cli@latest bundle "$file" --output "$file" --remove-unused-components
        echo "Placeholder: Running your command..."
    fi
done

echo "--------------------------------------------------------"
echo "Batch processing complete."
mv specs "$REPO_ROOT"
