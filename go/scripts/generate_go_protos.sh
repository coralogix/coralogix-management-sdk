#!/bin/bash

# Directory containing the protobuf files
proto_dir="../proto"
go_out_dir="internal"
mod_name="coralogix-management-sdk/go"
proto_files=($(find "$proto_dir" -name "*.proto" -print))
args=""
# Build arguments for import paths of all modules
for proto_file in "${proto_files[@]}" 
do
    out_module=$(dirname $proto_file)
  
    if [[ $out_module == *"coralogix"* ]]; then
        mod_path="${out_module##*/com/}"
        case "$mod_path" in
            *dashboards/v1/ast/widgets/common/queries.proto ) 
            ;; 
            *dashboards* )
                args+="--go_opt=M${proto_file##*$proto_dir/}=${mod_name}/${go_out_dir}/coralogixapis/dashboards/v1 "
                args+="--go-grpc_opt=M${proto_file##*$proto_dir/}=${mod_name}/${go_out_dir}/coralogixapis/dashboards/v1 "
                ;;
            *)    
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
        case "$proto_file" in
            *dashboards/v1/ast/widgets/common/queries.proto ) 
            ;; 
        *) 
            protoc --proto_path=$proto_dir --go_out=../.. --go-grpc_out=../.. $args $proto_file
            ;;
        esac
    fi
done

sed -i'.bak' -e 's/file_com_coralogixapis_dashboards_v1_ast_widgets_common_queries_proto_init()/\/\/file_com_coralogixapis_dashboards_v1_ast_widgets_common_queries_proto_init()/g' ${go_out_dir}/coralogixapis/dashboards/v1/*

rm -rf ${go_out_dir}/coralogixapis/dashboards/v1/*.bak