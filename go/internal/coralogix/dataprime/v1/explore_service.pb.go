// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: com/coralogix/dataprime/v1/explore_service.proto

package v1

import (
	_ "github.com/golang/protobuf/protoc-gen-go/descriptor"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_com_coralogix_dataprime_v1_explore_service_proto protoreflect.FileDescriptor

var file_com_coralogix_dataprime_v1_explore_service_proto_rawDesc = []byte{
	0x0a, 0x30, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x70,
	0x6c, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x2a,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd2, 0x01, 0x0a, 0x17, 0x44, 0x61, 0x74, 0x61,
	0x70, 0x72, 0x69, 0x6d, 0x65, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0xb6, 0x01, 0x0a, 0x11, 0x52, 0x75, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x70,
	0x72, 0x69, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x69, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x32, 0xc2, 0xb8, 0x02, 0x15, 0x0a, 0x13, 0x72, 0x75,
	0x6e, 0x20, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x20, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x72, 0x75, 0x6e, 0x30, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_com_coralogix_dataprime_v1_explore_service_proto_goTypes = []any{
	(*RunDataprimeQueryRequest)(nil),  // 0: com.coralogix.dataprime.v1.RunDataprimeQueryRequest
	(*RunDataprimeQueryResponse)(nil), // 1: com.coralogix.dataprime.v1.RunDataprimeQueryResponse
}
var file_com_coralogix_dataprime_v1_explore_service_proto_depIdxs = []int32{
	0, // 0: com.coralogix.dataprime.v1.DataprimeExploreService.RunDataprimeQuery:input_type -> com.coralogix.dataprime.v1.RunDataprimeQueryRequest
	1, // 1: com.coralogix.dataprime.v1.DataprimeExploreService.RunDataprimeQuery:output_type -> com.coralogix.dataprime.v1.RunDataprimeQueryResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogix_dataprime_v1_explore_service_proto_init() }
func file_com_coralogix_dataprime_v1_explore_service_proto_init() {
	if File_com_coralogix_dataprime_v1_explore_service_proto != nil {
		return
	}
	file_com_coralogix_dataprime_v1_audit_log_proto_init()
	file_com_coralogix_dataprime_v1_run_request_proto_init()
	file_com_coralogix_dataprime_v1_run_response_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_dataprime_v1_explore_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogix_dataprime_v1_explore_service_proto_goTypes,
		DependencyIndexes: file_com_coralogix_dataprime_v1_explore_service_proto_depIdxs,
	}.Build()
	File_com_coralogix_dataprime_v1_explore_service_proto = out.File
	file_com_coralogix_dataprime_v1_explore_service_proto_rawDesc = nil
	file_com_coralogix_dataprime_v1_explore_service_proto_goTypes = nil
	file_com_coralogix_dataprime_v1_explore_service_proto_depIdxs = nil
}
