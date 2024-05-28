#!/bin/bash

# Directory containing the protobuf files
proto_dir="../proto/com/"

# Loop through subdirectories
for subdir in "$proto_dir"/*; do
  # Check if it's a directory
  if [[ -d "$subdir" ]]; then
    # Find all .proto files with "service" keyword and store in an array
    proto_files=($(find "$subdir" -name "*.proto" -print | grep -E "service"))
    # Loop through each proto file path
    for proto_file in "${proto_files[@]}" 
    do
      # Extract filename
      filename=$(basename "$proto_file")

      # Run the protoc command
      python3 -m grpc_tools.protoc -I../proto --python_out=grpc --pyi_out=grpc --grpc_python_out=grpc "$proto_file"
      echo "Generated python files for $filename"
    done
  fi
done