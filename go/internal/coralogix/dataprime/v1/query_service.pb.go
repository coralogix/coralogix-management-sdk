// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: com/coralogix/dataprime/v1/query_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_com_coralogix_dataprime_v1_query_service_proto protoreflect.FileDescriptor

var file_com_coralogix_dataprime_v1_query_service_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x2a, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c,
	0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x26, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0x99, 0x12, 0x0a, 0x10, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xb2, 0x01, 0x0a, 0x0b, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69,
	0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69,
	0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x42, 0xc2, 0xb8, 0x02, 0x18, 0x0a,
	0x16, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x20, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d,
	0x65, 0x20, 0x71, 0x75, 0x65, 0x72, 0x79, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22, 0x1e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65,
	0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x12, 0xc3, 0x01,
	0x0a, 0x0e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x44, 0x64, 0x6c, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x44, 0x64, 0x6c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x44, 0x64, 0x6c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4a, 0xc2, 0xb8, 0x02, 0x1c, 0x0a, 0x1a, 0x73,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x20, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x20,
	0x64, 0x64, 0x6c, 0x20, 0x71, 0x75, 0x65, 0x72, 0x79, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x22,
	0x22, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69,
	0x6d, 0x65, 0x2f, 0x64, 0x64, 0x6c, 0x2d, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x73, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x12, 0xbc, 0x01, 0x0a, 0x0c, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x49, 0xc2, 0xb8, 0x02, 0x1e, 0x0a, 0x1c, 0x65,
	0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x20, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65,
	0x20, 0x71, 0x75, 0x65, 0x72, 0x79, 0x20, 0x70, 0x6c, 0x61, 0x6e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x21, 0x22, 0x1f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70,
	0x72, 0x69, 0x6d, 0x65, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x65, 0x78, 0x70, 0x6c, 0x61,
	0x69, 0x6e, 0x12, 0xb7, 0x01, 0x0a, 0x0a, 0x41, 0x77, 0x61, 0x69, 0x74, 0x52, 0x65, 0x61, 0x64,
	0x79, 0x12, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x77, 0x61, 0x69, 0x74, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x77,
	0x61, 0x69, 0x74, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x4a, 0xc2, 0xb8, 0x02, 0x21, 0x0a, 0x1f, 0x41, 0x77, 0x61, 0x69, 0x74, 0x20, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x20, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x62,
	0x65, 0x20, 0x72, 0x65, 0x61, 0x64, 0x79, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x22, 0x1d, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65,
	0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x61, 0x77, 0x61, 0x69, 0x74, 0x12, 0xb8, 0x01, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x2d, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70,
	0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4b, 0xc2, 0xb8, 0x02, 0x20,
	0x0a, 0x1e, 0x67, 0x61, 0x74, 0x68, 0x65, 0x72, 0x20, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69,
	0x6d, 0x65, 0x20, 0x71, 0x75, 0x65, 0x72, 0x79, 0x20, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22, 0x1f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0xd1, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x32, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x53, 0xc2, 0xb8, 0x02, 0x28, 0x0a, 0x26, 0x67, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x20, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x20, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x20, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x64, 0x20, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22, 0x1f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x30, 0x01, 0x12, 0xb2, 0x01, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x2d, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70,
	0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x43, 0xc2, 0xb8, 0x02, 0x18, 0x0a,
	0x16, 0x71, 0x75, 0x65, 0x72, 0x79, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x20,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22, 0x1f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65,
	0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x30, 0x01,
	0x12, 0xed, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x3a, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69,
	0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x59, 0xc2, 0xb8, 0x02, 0x1f, 0x0a, 0x1d, 0x6f, 0x62, 0x74,
	0x61, 0x69, 0x6e, 0x20, 0x70, 0x72, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x20, 0x64, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x20, 0x75, 0x72, 0x6c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30,
	0x22, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x69, 0x6d, 0x65, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x2d, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x2d, 0x75, 0x72, 0x6c,
	0x12, 0xac, 0x01, 0x0a, 0x0b, 0x44, 0x72, 0x6f, 0x70, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x12, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x72,
	0x6f, 0x70, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x72,
	0x6f, 0x70, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x3c, 0xc2, 0xb8, 0x02, 0x0e, 0x0a, 0x0c, 0x64, 0x72, 0x6f, 0x70, 0x20, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x22, 0x22, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x64,
	0x64, 0x6c, 0x2d, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x12,
	0xc2, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x49, 0xc2, 0xb8, 0x02, 0x1f, 0x0a,
	0x1d, 0x67, 0x61, 0x74, 0x68, 0x65, 0x72, 0x20, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d,
	0x65, 0x20, 0x71, 0x75, 0x65, 0x72, 0x79, 0x20, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x20, 0x22, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0xb3, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x43, 0xc2, 0xb8, 0x02, 0x19, 0x0a, 0x17, 0x63, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x20, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x61, 0x6d, 0x65, 0x20, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22, 0x1e, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x12, 0xb2, 0x01, 0x0a, 0x09, 0x4c,
	0x69, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69,
	0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x48, 0xc2, 0xb8, 0x02, 0x20, 0x0a, 0x1e, 0x6c, 0x69, 0x73,
	0x74, 0x20, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x20, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x69, 0x6d, 0x65, 0x20, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1e, 0x22, 0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70,
	0x72, 0x69, 0x6d, 0x65, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_com_coralogix_dataprime_v1_query_service_proto_goTypes = []any{
	(*SubmitQueryRequest)(nil),              // 0: com.coralogix.dataprime.v1.SubmitQueryRequest
	(*SubmitDdlQueryRequest)(nil),           // 1: com.coralogix.dataprime.v1.SubmitDdlQueryRequest
	(*ExplainQueryRequest)(nil),             // 2: com.coralogix.dataprime.v1.ExplainQueryRequest
	(*AwaitReadyRequest)(nil),               // 3: com.coralogix.dataprime.v1.AwaitReadyRequest
	(*GetMetricsRequest)(nil),               // 4: com.coralogix.dataprime.v1.GetMetricsRequest
	(*GetQueryResultsRequest)(nil),          // 5: com.coralogix.dataprime.v1.GetQueryResultsRequest
	(*GetDatasetRequest)(nil),               // 6: com.coralogix.dataprime.v1.GetDatasetRequest
	(*GetPresignedDownloadUrlRequest)(nil),  // 7: com.coralogix.dataprime.v1.GetPresignedDownloadUrlRequest
	(*DropDatasetRequest)(nil),              // 8: com.coralogix.dataprime.v1.DropDatasetRequest
	(*GetQueryStatusRequest)(nil),           // 9: com.coralogix.dataprime.v1.GetQueryStatusRequest
	(*CancelQueryRequest)(nil),              // 10: com.coralogix.dataprime.v1.CancelQueryRequest
	(*ListQueryRequest)(nil),                // 11: com.coralogix.dataprime.v1.ListQueryRequest
	(*SubmitQueryResponse)(nil),             // 12: com.coralogix.dataprime.v1.SubmitQueryResponse
	(*SubmitDdlQueryResponse)(nil),          // 13: com.coralogix.dataprime.v1.SubmitDdlQueryResponse
	(*ExplainQueryResponse)(nil),            // 14: com.coralogix.dataprime.v1.ExplainQueryResponse
	(*AwaitReadyResponse)(nil),              // 15: com.coralogix.dataprime.v1.AwaitReadyResponse
	(*GetMetricsResponse)(nil),              // 16: com.coralogix.dataprime.v1.GetMetricsResponse
	(*GetQueryResultsResponse)(nil),         // 17: com.coralogix.dataprime.v1.GetQueryResultsResponse
	(*GetDatasetResponse)(nil),              // 18: com.coralogix.dataprime.v1.GetDatasetResponse
	(*GetPresignedDownloadUrlResponse)(nil), // 19: com.coralogix.dataprime.v1.GetPresignedDownloadUrlResponse
	(*DropDatasetResponse)(nil),             // 20: com.coralogix.dataprime.v1.DropDatasetResponse
	(*GetQueryStatusResponse)(nil),          // 21: com.coralogix.dataprime.v1.GetQueryStatusResponse
	(*CancelQueryResponse)(nil),             // 22: com.coralogix.dataprime.v1.CancelQueryResponse
	(*ListQueryResponse)(nil),               // 23: com.coralogix.dataprime.v1.ListQueryResponse
}
var file_com_coralogix_dataprime_v1_query_service_proto_depIdxs = []int32{
	0,  // 0: com.coralogix.dataprime.v1.DataprimeService.SubmitQuery:input_type -> com.coralogix.dataprime.v1.SubmitQueryRequest
	1,  // 1: com.coralogix.dataprime.v1.DataprimeService.SubmitDdlQuery:input_type -> com.coralogix.dataprime.v1.SubmitDdlQueryRequest
	2,  // 2: com.coralogix.dataprime.v1.DataprimeService.ExplainQuery:input_type -> com.coralogix.dataprime.v1.ExplainQueryRequest
	3,  // 3: com.coralogix.dataprime.v1.DataprimeService.AwaitReady:input_type -> com.coralogix.dataprime.v1.AwaitReadyRequest
	4,  // 4: com.coralogix.dataprime.v1.DataprimeService.GetMetrics:input_type -> com.coralogix.dataprime.v1.GetMetricsRequest
	5,  // 5: com.coralogix.dataprime.v1.DataprimeService.GetQueryResults:input_type -> com.coralogix.dataprime.v1.GetQueryResultsRequest
	6,  // 6: com.coralogix.dataprime.v1.DataprimeService.GetDataset:input_type -> com.coralogix.dataprime.v1.GetDatasetRequest
	7,  // 7: com.coralogix.dataprime.v1.DataprimeService.GetPresignedDownloadUrl:input_type -> com.coralogix.dataprime.v1.GetPresignedDownloadUrlRequest
	8,  // 8: com.coralogix.dataprime.v1.DataprimeService.DropDataset:input_type -> com.coralogix.dataprime.v1.DropDatasetRequest
	9,  // 9: com.coralogix.dataprime.v1.DataprimeService.GetQueryStatus:input_type -> com.coralogix.dataprime.v1.GetQueryStatusRequest
	10, // 10: com.coralogix.dataprime.v1.DataprimeService.CancelQuery:input_type -> com.coralogix.dataprime.v1.CancelQueryRequest
	11, // 11: com.coralogix.dataprime.v1.DataprimeService.ListQuery:input_type -> com.coralogix.dataprime.v1.ListQueryRequest
	12, // 12: com.coralogix.dataprime.v1.DataprimeService.SubmitQuery:output_type -> com.coralogix.dataprime.v1.SubmitQueryResponse
	13, // 13: com.coralogix.dataprime.v1.DataprimeService.SubmitDdlQuery:output_type -> com.coralogix.dataprime.v1.SubmitDdlQueryResponse
	14, // 14: com.coralogix.dataprime.v1.DataprimeService.ExplainQuery:output_type -> com.coralogix.dataprime.v1.ExplainQueryResponse
	15, // 15: com.coralogix.dataprime.v1.DataprimeService.AwaitReady:output_type -> com.coralogix.dataprime.v1.AwaitReadyResponse
	16, // 16: com.coralogix.dataprime.v1.DataprimeService.GetMetrics:output_type -> com.coralogix.dataprime.v1.GetMetricsResponse
	17, // 17: com.coralogix.dataprime.v1.DataprimeService.GetQueryResults:output_type -> com.coralogix.dataprime.v1.GetQueryResultsResponse
	18, // 18: com.coralogix.dataprime.v1.DataprimeService.GetDataset:output_type -> com.coralogix.dataprime.v1.GetDatasetResponse
	19, // 19: com.coralogix.dataprime.v1.DataprimeService.GetPresignedDownloadUrl:output_type -> com.coralogix.dataprime.v1.GetPresignedDownloadUrlResponse
	20, // 20: com.coralogix.dataprime.v1.DataprimeService.DropDataset:output_type -> com.coralogix.dataprime.v1.DropDatasetResponse
	21, // 21: com.coralogix.dataprime.v1.DataprimeService.GetQueryStatus:output_type -> com.coralogix.dataprime.v1.GetQueryStatusResponse
	22, // 22: com.coralogix.dataprime.v1.DataprimeService.CancelQuery:output_type -> com.coralogix.dataprime.v1.CancelQueryResponse
	23, // 23: com.coralogix.dataprime.v1.DataprimeService.ListQuery:output_type -> com.coralogix.dataprime.v1.ListQueryResponse
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogix_dataprime_v1_query_service_proto_init() }
func file_com_coralogix_dataprime_v1_query_service_proto_init() {
	if File_com_coralogix_dataprime_v1_query_service_proto != nil {
		return
	}
	file_com_coralogix_dataprime_v1_audit_log_proto_init()
	file_com_coralogix_dataprime_v1_metrics_proto_init()
	file_com_coralogix_dataprime_v1_query_proto_init()
	file_com_coralogix_dataprime_v1_results_proto_init()
	file_com_coralogix_dataprime_v1_list_queries_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_dataprime_v1_query_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogix_dataprime_v1_query_service_proto_goTypes,
		DependencyIndexes: file_com_coralogix_dataprime_v1_query_service_proto_depIdxs,
	}.Build()
	File_com_coralogix_dataprime_v1_query_service_proto = out.File
	file_com_coralogix_dataprime_v1_query_service_proto_rawDesc = nil
	file_com_coralogix_dataprime_v1_query_service_proto_goTypes = nil
	file_com_coralogix_dataprime_v1_query_service_proto_depIdxs = nil
}
