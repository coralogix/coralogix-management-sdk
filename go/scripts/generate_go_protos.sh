#!/bin/bash

# Directory containing the protobuf files
proto_dir="../proto"
proto_files=($(find "$proto_dir" -name "*.proto" -print))
args=""
# Build arguments for import paths of all modules
for proto_file in "${proto_files[@]}" 
do
    out_module=$(dirname $proto_file)
    out_module=${out_module/coralogixapis/coralogix}
    if [[ $out_module == *"coralogix"* ]]; then
        mod_path="${out_module##*../proto/com/}"
        args+="--go_opt=M${proto_file##*$proto_dir/}=coralogix-management-sdk/go/internal/${mod_path} "
        args+="--go-grpc_opt=M${proto_file##*$proto_dir/}=coralogix-management-sdk/go/internal/${mod_path} "

    fi
done

# Generate go files one by one. All together only works if there are no duplicates
for proto_file in "${proto_files[@]}" 
do
    if [[ $proto_file == *"coralogix"* ]]; then
        protoc --proto_path=$proto_dir --go_out=../.. --go-grpc_out=../.. $args $proto_file
    fi
done
