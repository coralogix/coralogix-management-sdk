// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: com/coralogix/dpxl/v1/unauthenticated_internal_dpxl_service.proto

package v1

import (
	_ "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/dataprime/v1"
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

var File_com_coralogix_dpxl_v1_unauthenticated_internal_dpxl_service_proto protoreflect.FileDescriptor

var file_com_coralogix_dpxl_v1_unauthenticated_internal_dpxl_service_proto_rawDesc = []byte{
	0x0a, 0x41, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x64, 0x70, 0x78, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x5f, 0x64, 0x70, 0x78, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x64, 0x70, 0x78, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x41, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x64, 0x70, 0x78, 0x6c, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x64, 0x70, 0x78, 0x6c, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x42, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x64, 0x70, 0x78,
	0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x64, 0x70,
	0x78, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2a, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xc6, 0x0c, 0x0a, 0x22, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x70, 0x78, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xc3, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70, 0x69,
	0x6c, 0x65, 0x12, 0x47, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x64, 0x70, 0x78, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x61, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x44, 0x70, 0x78, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x48, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x70, 0x78, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x70, 0x78, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0xc2, 0xb8, 0x02, 0x21, 0x0a, 0x1f, 0x43, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x20, 0x44, 0x50, 0x58, 0x4c, 0x20, 0x74, 0x6f, 0x20, 0x73, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0xd1, 0x01, 0x0a,
	0x0a, 0x54, 0x6f, 0x44, 0x70, 0x78, 0x6c, 0x54, 0x65, 0x78, 0x74, 0x12, 0x4a, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x70, 0x78, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x70, 0x78, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x44, 0x70, 0x78, 0x6c, 0x54, 0x65, 0x78, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x4b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x70, 0x78, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x70, 0x78, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x54, 0x6f, 0x44, 0x70, 0x78, 0x6c, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0xc2, 0xb8, 0x02, 0x26, 0x0a, 0x24, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x74, 0x20, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x44,
	0x50, 0x58, 0x4c, 0x20, 0x74, 0x6f, 0x20, 0x44, 0x50, 0x58, 0x4c, 0x20, 0x74, 0x65, 0x78, 0x74,
	0x12, 0xec, 0x01, 0x0a, 0x12, 0x54, 0x6f, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x43, 0x6c, 0x61, 0x75, 0x73, 0x65, 0x12, 0x52, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x70, 0x78, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x70, 0x78, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x54, 0x6f, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x6c,
	0x61, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x53, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x70, 0x78, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x70, 0x78, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x43, 0x6c, 0x61, 0x75, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x2d, 0xc2, 0xb8, 0x02, 0x29, 0x0a, 0x27, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74,
	0x65, 0x20, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x44, 0x50, 0x58,
	0x4c, 0x20, 0x74, 0x6f, 0x20, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x9d, 0x02, 0x0a, 0x1c, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x41, 0x73, 0x74,
	0x12, 0x5c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x64, 0x70, 0x78, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x44, 0x70, 0x78, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x53, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65,
	0x43, 0x6f, 0x72, 0x65, 0x41, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x5d,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64,
	0x70, 0x78, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x70,
	0x78, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x43, 0x6f,
	0x72, 0x65, 0x41, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x40, 0xc2,
	0xb8, 0x02, 0x3c, 0x0a, 0x3a, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x20, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x44, 0x50, 0x58, 0x4c, 0x20, 0x74,
	0x6f, 0x20, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x44, 0x61, 0x74,
	0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x20, 0x43, 0x6f, 0x72, 0x65, 0x20, 0x41, 0x53, 0x54, 0x12,
	0xc9, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x12, 0x47, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x70, 0x78, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x70, 0x78, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x48, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x70, 0x78, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x61,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x44, 0x70, 0x78, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43,
	0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b,
	0xc2, 0xb8, 0x02, 0x27, 0x0a, 0x25, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x20, 0x74, 0x77,
	0x6f, 0x20, 0x44, 0x50, 0x58, 0x4c, 0x20, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x20, 0x69, 0x6e, 0x74, 0x6f, 0x20, 0x6f, 0x6e, 0x65, 0x12, 0xda, 0x01, 0x0a, 0x0b,
	0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x4d, 0x61, 0x6e, 0x79, 0x12, 0x4b, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x70, 0x78, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x70, 0x78, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x4d, 0x61, 0x6e,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x4c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x70, 0x78, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x70, 0x78, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0xc2, 0xb8, 0x02, 0x2c, 0x0a, 0x2a, 0x43, 0x6f,
	0x6d, 0x62, 0x69, 0x6e, 0x65, 0x20, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x20, 0x44,
	0x50, 0x58, 0x4c, 0x20, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x20,
	0x69, 0x6e, 0x74, 0x6f, 0x20, 0x6f, 0x6e, 0x65, 0x12, 0xcd, 0x01, 0x0a, 0x09, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x49, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x70, 0x78, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x6e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x70, 0x78, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x4a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2e, 0x64, 0x70, 0x78, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x44, 0x70, 0x78, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0xc2,
	0xb8, 0x02, 0x25, 0x0a, 0x23, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74,
	0x79, 0x70, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x20, 0x44, 0x50, 0x58, 0x4c, 0x20, 0x65, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_com_coralogix_dpxl_v1_unauthenticated_internal_dpxl_service_proto_goTypes = []any{
	(*UnauthenticatedInternalDpxlServiceCompileRequest)(nil),                       // 0: com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlServiceCompileRequest
	(*UnauthenticatedInternalDpxlServiceToDpxlTextRequest)(nil),                    // 1: com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlServiceToDpxlTextRequest
	(*UnauthenticatedInternalDpxlServiceToOpenSearchClauseRequest)(nil),            // 2: com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlServiceToOpenSearchClauseRequest
	(*UnauthenticatedInternalDpxlServiceToSerializedDataprimeCoreAstRequest)(nil),  // 3: com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlServiceToSerializedDataprimeCoreAstRequest
	(*UnauthenticatedInternalDpxlServiceCombineRequest)(nil),                       // 4: com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlServiceCombineRequest
	(*UnauthenticatedInternalDpxlServiceCombineManyRequest)(nil),                   // 5: com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlServiceCombineManyRequest
	(*UnauthenticatedInternalDpxlServiceCheckTypeRequest)(nil),                     // 6: com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlServiceCheckTypeRequest
	(*UnauthenticatedInternalDpxlServiceCompileResponse)(nil),                      // 7: com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlServiceCompileResponse
	(*UnauthenticatedInternalDpxlServiceToDpxlTextResponse)(nil),                   // 8: com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlServiceToDpxlTextResponse
	(*UnauthenticatedInternalDpxlServiceToOpenSearchClauseResponse)(nil),           // 9: com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlServiceToOpenSearchClauseResponse
	(*UnauthenticatedInternalDpxlServiceToSerializedDataprimeCoreAstResponse)(nil), // 10: com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlServiceToSerializedDataprimeCoreAstResponse
	(*UnauthenticatedInternalDpxlServiceCombineResponse)(nil),                      // 11: com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlServiceCombineResponse
	(*UnauthenticatedInternalDpxlServiceCombineManyResponse)(nil),                  // 12: com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlServiceCombineManyResponse
	(*UnauthenticatedInternalDpxlServiceCheckTypeResponse)(nil),                    // 13: com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlServiceCheckTypeResponse
}
var file_com_coralogix_dpxl_v1_unauthenticated_internal_dpxl_service_proto_depIdxs = []int32{
	0,  // 0: com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlService.Compile:input_type -> com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlServiceCompileRequest
	1,  // 1: com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlService.ToDpxlText:input_type -> com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlServiceToDpxlTextRequest
	2,  // 2: com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlService.ToOpenSearchClause:input_type -> com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlServiceToOpenSearchClauseRequest
	3,  // 3: com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlService.ToSerializedDataprimeCoreAst:input_type -> com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlServiceToSerializedDataprimeCoreAstRequest
	4,  // 4: com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlService.Combine:input_type -> com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlServiceCombineRequest
	5,  // 5: com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlService.CombineMany:input_type -> com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlServiceCombineManyRequest
	6,  // 6: com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlService.CheckType:input_type -> com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlServiceCheckTypeRequest
	7,  // 7: com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlService.Compile:output_type -> com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlServiceCompileResponse
	8,  // 8: com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlService.ToDpxlText:output_type -> com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlServiceToDpxlTextResponse
	9,  // 9: com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlService.ToOpenSearchClause:output_type -> com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlServiceToOpenSearchClauseResponse
	10, // 10: com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlService.ToSerializedDataprimeCoreAst:output_type -> com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlServiceToSerializedDataprimeCoreAstResponse
	11, // 11: com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlService.Combine:output_type -> com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlServiceCombineResponse
	12, // 12: com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlService.CombineMany:output_type -> com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlServiceCombineManyResponse
	13, // 13: com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlService.CheckType:output_type -> com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlServiceCheckTypeResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogix_dpxl_v1_unauthenticated_internal_dpxl_service_proto_init() }
func file_com_coralogix_dpxl_v1_unauthenticated_internal_dpxl_service_proto_init() {
	if File_com_coralogix_dpxl_v1_unauthenticated_internal_dpxl_service_proto != nil {
		return
	}
	file_com_coralogix_dpxl_v1_unauthenticated_internal_dpxl_request_proto_init()
	file_com_coralogix_dpxl_v1_unauthenticated_internal_dpxl_response_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_dpxl_v1_unauthenticated_internal_dpxl_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogix_dpxl_v1_unauthenticated_internal_dpxl_service_proto_goTypes,
		DependencyIndexes: file_com_coralogix_dpxl_v1_unauthenticated_internal_dpxl_service_proto_depIdxs,
	}.Build()
	File_com_coralogix_dpxl_v1_unauthenticated_internal_dpxl_service_proto = out.File
	file_com_coralogix_dpxl_v1_unauthenticated_internal_dpxl_service_proto_rawDesc = nil
	file_com_coralogix_dpxl_v1_unauthenticated_internal_dpxl_service_proto_goTypes = nil
	file_com_coralogix_dpxl_v1_unauthenticated_internal_dpxl_service_proto_depIdxs = nil
}
