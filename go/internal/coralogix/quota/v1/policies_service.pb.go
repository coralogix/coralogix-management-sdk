// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: com/coralogix/quota/v1/policies_service.proto

package v1

import (
	_ "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/common/v1"
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

var File_com_coralogix_quota_v1_policies_service_proto protoreflect.FileDescriptor

var file_com_coralogix_quota_v1_policies_service_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71,
	0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x27, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x30, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x71, 0x75, 0x6f, 0x74,
	0x61, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x33, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65,
	0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x3a, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x71, 0x75, 0x6f, 0x74,
	0x61, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x33, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x71, 0x75, 0x6f, 0x74,
	0x61, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3b, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x6c,
	0x6b, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x3c, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x6c, 0x6b, 0x5f, 0x74,
	0x65, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x32, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x71,
	0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x67, 0x67,
	0x6c, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x37, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31,
	0x2f, 0x62, 0x75, 0x6c, 0x6b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x38, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x6c, 0x6b, 0x5f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3f, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x71, 0x75, 0x6f, 0x74,
	0x61, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x5f, 0x6f, 0x76, 0x65, 0x72,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb2, 0x10, 0x0a, 0x0f, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8b,
	0x01, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x28, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x29, 0xc2, 0xb8, 0x02, 0x0c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x20, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x95, 0x01, 0x0a,
	0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x2b, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75,
	0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0xc2, 0xb8, 0x02, 0x0f, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x65, 0x73, 0x12, 0x95, 0x01, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x2a, 0xc2, 0xb8, 0x02, 0x0f, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x1a, 0x0c,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0xab, 0x01, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x65, 0x73, 0x12, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0xc2, 0xb8, 0x02, 0x16,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x20, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x20, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x97, 0x01, 0x0a, 0x0c, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x2b, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0xc2, 0xb8, 0x02, 0x0f, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x20, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x13, 0x2a, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0xa9, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0xc2, 0xb8, 0x02, 0x12, 0x0a,
	0x10, 0x52, 0x65, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x20, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x3a, 0x72, 0x65, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0xbf, 0x01, 0x0a, 0x13, 0x42, 0x75, 0x6c, 0x6b, 0x54, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x75, 0x6c, 0x6b, 0x54, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x6c, 0x6b, 0x54, 0x65, 0x73, 0x74, 0x4c, 0x6f,
	0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x3f, 0xc2, 0xb8, 0x02, 0x18, 0x0a, 0x16, 0x42, 0x75, 0x6c, 0x6b, 0x20, 0x54, 0x65,
	0x73, 0x74, 0x20, 0x4c, 0x6f, 0x67, 0x20, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x3a, 0x62, 0x75, 0x6c, 0x6b, 0x54, 0x65, 0x73, 0x74, 0x4c,
	0x6f, 0x67, 0x12, 0x9c, 0x01, 0x0a, 0x0c, 0x54, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x12, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x67,
	0x67, 0x6c, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x67, 0x67, 0x6c, 0x65,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x31,
	0xc2, 0xb8, 0x02, 0x0f, 0x0a, 0x0d, 0x54, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x20, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x3a, 0x74, 0x6f, 0x67, 0x67, 0x6c,
	0x65, 0x12, 0xc6, 0x01, 0x0a, 0x17, 0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x36, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75,
	0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x74, 0x6f, 0x6d, 0x69, 0x63, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3a,
	0xc2, 0xb8, 0x02, 0x14, 0x0a, 0x12, 0x42, 0x75, 0x6c, 0x6b, 0x20, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x20, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01,
	0x2a, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x3a,
	0x62, 0x75, 0x6c, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x9d, 0x02, 0x0a, 0x1a, 0x41,
	0x74, 0x6f, 0x6d, 0x69, 0x63, 0x4f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x6f,
	0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x4f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x4c, 0x6f, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74,
	0x6f, 0x6d, 0x69, 0x63, 0x4f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x6f, 0x67,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x87, 0x01, 0xc2, 0xb8, 0x02, 0x51, 0x0a, 0x4f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x73,
	0x20, 0x61, 0x6c, 0x6c, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x6c, 0x6f,
	0x67, 0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x6e, 0x65, 0x77, 0x6c, 0x79,
	0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x20, 0x6f, 0x6e, 0x65, 0x20, 0x61, 0x74,
	0x6f, 0x6d, 0x69, 0x63, 0x61, 0x6c, 0x6c, 0x79, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x3a, 0x01,
	0x2a, 0x22, 0x27, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x3a,
	0x61, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x4f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x4c,
	0x6f, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0xa2, 0x02, 0x0a, 0x1b, 0x41,
	0x74, 0x6f, 0x6d, 0x69, 0x63, 0x4f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x53, 0x70,
	0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x3a, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x4f, 0x76, 0x65, 0x72, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x53, 0x70, 0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x4f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x53,
	0x70, 0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x89, 0x01, 0xc2, 0xb8, 0x02, 0x52, 0x0a, 0x50, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x73, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x20, 0x73, 0x70, 0x61, 0x6e, 0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x20, 0x61,
	0x6e, 0x64, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x6e,
	0x65, 0x77, 0x6c, 0x79, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x20, 0x6f, 0x6e,
	0x65, 0x20, 0x61, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x61, 0x6c, 0x6c, 0x79, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x2d, 0x3a, 0x01, 0x2a, 0x22, 0x28, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x65, 0x73, 0x3a, 0x61, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x4f, 0x76, 0x65, 0x72, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x53, 0x70, 0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_com_coralogix_quota_v1_policies_service_proto_goTypes = []any{
	(*GetPolicyRequest)(nil),                    // 0: com.coralogix.quota.v1.GetPolicyRequest
	(*CreatePolicyRequest)(nil),                 // 1: com.coralogix.quota.v1.CreatePolicyRequest
	(*UpdatePolicyRequest)(nil),                 // 2: com.coralogix.quota.v1.UpdatePolicyRequest
	(*GetCompanyPoliciesRequest)(nil),           // 3: com.coralogix.quota.v1.GetCompanyPoliciesRequest
	(*DeletePolicyRequest)(nil),                 // 4: com.coralogix.quota.v1.DeletePolicyRequest
	(*ReorderPoliciesRequest)(nil),              // 5: com.coralogix.quota.v1.ReorderPoliciesRequest
	(*BulkTestLogPoliciesRequest)(nil),          // 6: com.coralogix.quota.v1.BulkTestLogPoliciesRequest
	(*TogglePolicyRequest)(nil),                 // 7: com.coralogix.quota.v1.TogglePolicyRequest
	(*AtomicBatchCreatePolicyRequest)(nil),      // 8: com.coralogix.quota.v1.AtomicBatchCreatePolicyRequest
	(*AtomicOverwriteLogPoliciesRequest)(nil),   // 9: com.coralogix.quota.v1.AtomicOverwriteLogPoliciesRequest
	(*AtomicOverwriteSpanPoliciesRequest)(nil),  // 10: com.coralogix.quota.v1.AtomicOverwriteSpanPoliciesRequest
	(*GetPolicyResponse)(nil),                   // 11: com.coralogix.quota.v1.GetPolicyResponse
	(*CreatePolicyResponse)(nil),                // 12: com.coralogix.quota.v1.CreatePolicyResponse
	(*UpdatePolicyResponse)(nil),                // 13: com.coralogix.quota.v1.UpdatePolicyResponse
	(*GetCompanyPoliciesResponse)(nil),          // 14: com.coralogix.quota.v1.GetCompanyPoliciesResponse
	(*DeletePolicyResponse)(nil),                // 15: com.coralogix.quota.v1.DeletePolicyResponse
	(*ReorderPoliciesResponse)(nil),             // 16: com.coralogix.quota.v1.ReorderPoliciesResponse
	(*BulkTestLogPoliciesResponse)(nil),         // 17: com.coralogix.quota.v1.BulkTestLogPoliciesResponse
	(*TogglePolicyResponse)(nil),                // 18: com.coralogix.quota.v1.TogglePolicyResponse
	(*AtomicBatchCreatePolicyResponse)(nil),     // 19: com.coralogix.quota.v1.AtomicBatchCreatePolicyResponse
	(*AtomicOverwriteLogPoliciesResponse)(nil),  // 20: com.coralogix.quota.v1.AtomicOverwriteLogPoliciesResponse
	(*AtomicOverwriteSpanPoliciesResponse)(nil), // 21: com.coralogix.quota.v1.AtomicOverwriteSpanPoliciesResponse
}
var file_com_coralogix_quota_v1_policies_service_proto_depIdxs = []int32{
	0,  // 0: com.coralogix.quota.v1.PoliciesService.GetPolicy:input_type -> com.coralogix.quota.v1.GetPolicyRequest
	1,  // 1: com.coralogix.quota.v1.PoliciesService.CreatePolicy:input_type -> com.coralogix.quota.v1.CreatePolicyRequest
	2,  // 2: com.coralogix.quota.v1.PoliciesService.UpdatePolicy:input_type -> com.coralogix.quota.v1.UpdatePolicyRequest
	3,  // 3: com.coralogix.quota.v1.PoliciesService.GetCompanyPolicies:input_type -> com.coralogix.quota.v1.GetCompanyPoliciesRequest
	4,  // 4: com.coralogix.quota.v1.PoliciesService.DeletePolicy:input_type -> com.coralogix.quota.v1.DeletePolicyRequest
	5,  // 5: com.coralogix.quota.v1.PoliciesService.ReorderPolicies:input_type -> com.coralogix.quota.v1.ReorderPoliciesRequest
	6,  // 6: com.coralogix.quota.v1.PoliciesService.BulkTestLogPolicies:input_type -> com.coralogix.quota.v1.BulkTestLogPoliciesRequest
	7,  // 7: com.coralogix.quota.v1.PoliciesService.TogglePolicy:input_type -> com.coralogix.quota.v1.TogglePolicyRequest
	8,  // 8: com.coralogix.quota.v1.PoliciesService.AtomicBatchCreatePolicy:input_type -> com.coralogix.quota.v1.AtomicBatchCreatePolicyRequest
	9,  // 9: com.coralogix.quota.v1.PoliciesService.AtomicOverwriteLogPolicies:input_type -> com.coralogix.quota.v1.AtomicOverwriteLogPoliciesRequest
	10, // 10: com.coralogix.quota.v1.PoliciesService.AtomicOverwriteSpanPolicies:input_type -> com.coralogix.quota.v1.AtomicOverwriteSpanPoliciesRequest
	11, // 11: com.coralogix.quota.v1.PoliciesService.GetPolicy:output_type -> com.coralogix.quota.v1.GetPolicyResponse
	12, // 12: com.coralogix.quota.v1.PoliciesService.CreatePolicy:output_type -> com.coralogix.quota.v1.CreatePolicyResponse
	13, // 13: com.coralogix.quota.v1.PoliciesService.UpdatePolicy:output_type -> com.coralogix.quota.v1.UpdatePolicyResponse
	14, // 14: com.coralogix.quota.v1.PoliciesService.GetCompanyPolicies:output_type -> com.coralogix.quota.v1.GetCompanyPoliciesResponse
	15, // 15: com.coralogix.quota.v1.PoliciesService.DeletePolicy:output_type -> com.coralogix.quota.v1.DeletePolicyResponse
	16, // 16: com.coralogix.quota.v1.PoliciesService.ReorderPolicies:output_type -> com.coralogix.quota.v1.ReorderPoliciesResponse
	17, // 17: com.coralogix.quota.v1.PoliciesService.BulkTestLogPolicies:output_type -> com.coralogix.quota.v1.BulkTestLogPoliciesResponse
	18, // 18: com.coralogix.quota.v1.PoliciesService.TogglePolicy:output_type -> com.coralogix.quota.v1.TogglePolicyResponse
	19, // 19: com.coralogix.quota.v1.PoliciesService.AtomicBatchCreatePolicy:output_type -> com.coralogix.quota.v1.AtomicBatchCreatePolicyResponse
	20, // 20: com.coralogix.quota.v1.PoliciesService.AtomicOverwriteLogPolicies:output_type -> com.coralogix.quota.v1.AtomicOverwriteLogPoliciesResponse
	21, // 21: com.coralogix.quota.v1.PoliciesService.AtomicOverwriteSpanPolicies:output_type -> com.coralogix.quota.v1.AtomicOverwriteSpanPoliciesResponse
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogix_quota_v1_policies_service_proto_init() }
func file_com_coralogix_quota_v1_policies_service_proto_init() {
	if File_com_coralogix_quota_v1_policies_service_proto != nil {
		return
	}
	file_com_coralogix_quota_v1_get_policy_request_proto_init()
	file_com_coralogix_quota_v1_get_policy_response_proto_init()
	file_com_coralogix_quota_v1_create_policy_request_proto_init()
	file_com_coralogix_quota_v1_create_policy_response_proto_init()
	file_com_coralogix_quota_v1_update_policy_request_proto_init()
	file_com_coralogix_quota_v1_update_policy_response_proto_init()
	file_com_coralogix_quota_v1_get_company_policies_request_proto_init()
	file_com_coralogix_quota_v1_get_company_policies_response_proto_init()
	file_com_coralogix_quota_v1_delete_policy_request_proto_init()
	file_com_coralogix_quota_v1_delete_policy_response_proto_init()
	file_com_coralogix_quota_v1_reorder_policies_request_proto_init()
	file_com_coralogix_quota_v1_reorder_policies_response_proto_init()
	file_com_coralogix_quota_v1_bulk_test_log_policies_request_proto_init()
	file_com_coralogix_quota_v1_bulk_test_log_policies_response_proto_init()
	file_com_coralogix_quota_v1_toggle_policy_request_proto_init()
	file_com_coralogix_quota_v1_toggle_policy_response_proto_init()
	file_com_coralogix_quota_v1_bulk_create_policy_request_proto_init()
	file_com_coralogix_quota_v1_bulk_create_policy_response_proto_init()
	file_com_coralogix_quota_v1_atomic_overwrite_policies_response_proto_init()
	file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_quota_v1_policies_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogix_quota_v1_policies_service_proto_goTypes,
		DependencyIndexes: file_com_coralogix_quota_v1_policies_service_proto_depIdxs,
	}.Build()
	File_com_coralogix_quota_v1_policies_service_proto = out.File
	file_com_coralogix_quota_v1_policies_service_proto_rawDesc = nil
	file_com_coralogix_quota_v1_policies_service_proto_goTypes = nil
	file_com_coralogix_quota_v1_policies_service_proto_depIdxs = nil
}
