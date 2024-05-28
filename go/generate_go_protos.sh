#!/bin/bash

# Directory containing the protobuf files
proto_dir="../proto"
proto_files=($(find "$proto_dir" -name "*.proto" -print))
args=""
# Build arguments for import paths of all modules. TODO: remove Google's
for proto_file in "${proto_files[@]}" 
do
    args+="--go_opt=M${proto_file##*$proto_dir/}=pb/${proto_file##*../proto} "
done

# Generate go files one by one. All together only works if there are no duplicates
for proto_file in "${proto_files[@]}" 
do
    protoc --proto_path=$proto_dir --go_out=.  $args $proto_file
done
