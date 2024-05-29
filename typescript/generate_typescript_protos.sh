#!/bin/bash

# Directory containing the protobuf files
proto_dir="../proto"
proto_files=($(find "$proto_dir" -name "*.proto" -print))
args=""

# Generate go files one by one. All together only works if there are no duplicates
for proto_file in "${proto_files[@]}" 
do
    filename=$(basename "$proto_file")
    protoc -I../proto --plugin=./node_modules/.bin/protoc-gen-ts_proto --ts_proto_out=src/grpc $proto_file
    echo "Generated typescript files for $filename"
done

