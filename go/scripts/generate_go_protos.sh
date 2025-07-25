#!/bin/bash
# Copyright 2024 Coralogix Ltd.
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     https://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


# Directory containing the protobuf files
proto_dir="../proto"
go_out_dir="internal"
mod_prefix="github.com/coralogix"
mod_name="$mod_prefix/coralogix-management-sdk/go"
proto_files=($(find "$proto_dir" -name "*.proto" -print))
args=""
# Build arguments for import paths of all modules
for proto_file in "${proto_files[@]}" 
do
    out_module=$(dirname $proto_file)
  
    if [[ $out_module == *"coralogix"* ]]; then
        mod_path="${out_module##*/com/}"
        case "$mod_path" in
            # The archive v2 protos contain circular dependencies, so we need to make sure that all dashboard files end up in the same package
            *archive/dataset/v2* )
                args+="--go_opt=M${proto_file##*$proto_dir/}=${mod_name}/${go_out_dir}/coralogix/archive/v2 "
                args+="--go-grpc_opt=M${proto_file##*$proto_dir/}=${mod_name}/${go_out_dir}/coralogix/archive/v2 "
                ;;
            # The alerts v3 protos contain circular dependencies, so the files are going into the same namespace
            *alerts/v3* )
                args+="--go_opt=M${proto_file##*$proto_dir/}=${mod_name}/${go_out_dir}/coralogixapis/alerts/v3 "
                args+="--go-grpc_opt=M${proto_file##*$proto_dir/}=${mod_name}/${go_out_dir}/coralogixapis/alerts/v3 "
                ;;
            # The permissions proto contain circular dependencies, so we need to make sure that all permissions files end up in the same package
            *permissions* )
                args+="--go_opt=M${proto_file##*$proto_dir/}=${mod_name}/${go_out_dir}/coralogix/permissions/v1 "
                args+="--go-grpc_opt=M${proto_file##*$proto_dir/}=${mod_name}/${go_out_dir}/coralogix/permissions/v1 "
                ;;
            *incidents* )
                args+="--go_opt=M${proto_file##*$proto_dir/}=${mod_name}/${go_out_dir}/coralogixapis/incidents/v1 "
                args+="--go-grpc_opt=M${proto_file##*$proto_dir/}=${mod_name}/${go_out_dir}/coralogixapis/incidents/v1 "
                ;;
            *)    
            # For all other protos, the package path is the same as the directory path
                args+="--go_opt=M${proto_file##*$proto_dir/}=${mod_name}/${go_out_dir}/${mod_path} "
                args+="--go-grpc_opt=M${proto_file##*$proto_dir/}=${mod_name}/${go_out_dir}/${mod_path} "
                ;;
        esac
    fi
done

# Generate go files one by one. All together only works if there are no duplicates
for proto_file in "${proto_files[@]}" 
do
    if [[ $proto_file == *"coralogix"* ]]; then 
        protoc --proto_path=$proto_dir --go_out=../.. --go-grpc_out=../.. --go-grpc_opt=module=$mod_prefix --go_opt=module=$mod_prefix $args $proto_file
    fi
done


