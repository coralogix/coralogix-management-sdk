#!/bin/bash

# Directory containing the protobuf files
proto_dir="../proto"
proto_files=($(find "$proto_dir" -name "*.proto" -print))

# Generate go files one by one. All together only works if there are no duplicates
for proto_file in "${proto_files[@]}" 
do
    filename=$(basename "$proto_file")
    python3 -m grpc_tools.protoc -I../proto --python_out=proto --pyi_out=proto --grpc_python_out=proto "$proto_file"
    echo "Generated python files for $filename"
done
