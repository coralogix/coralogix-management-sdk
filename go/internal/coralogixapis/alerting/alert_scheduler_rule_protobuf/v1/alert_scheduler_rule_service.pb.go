// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: com/coralogixapis/alerting/alert_scheduler_rule_protobuf/v1/alert_scheduler_rule_service.proto

package v1

import (
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

var File_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_service_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_service_proto_rawDesc = []byte{
	0x0a, 0x5e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75,
	0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x3b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x1a, 0x59, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x5a, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb1, 0x0c, 0x0a, 0x19, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0xd0, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x59, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x5a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xd9, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x75,
	0x6c, 0x65, 0x12, 0x5c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f,
	0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x5d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75,
	0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0xd9, 0x01, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x5c,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x5d, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52,
	0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xd9, 0x01,
	0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x5c, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x5d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xdc, 0x01, 0x0a, 0x19, 0x47, 0x65,
	0x74, 0x42, 0x75, 0x6c, 0x6b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x5d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x5e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xe5, 0x01, 0x0a, 0x1c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x75, 0x6c, 0x6b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x60, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75,
	0x6c, 0x6b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x61, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x75, 0x6c, 0x6b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0xe5, 0x01, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x6c, 0x6b, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c,
	0x65, 0x12, 0x60, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72,
	0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x6c, 0x6b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x6c, 0x6b, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_service_proto_goTypes = []any{
	(*GetAlertSchedulerRuleRequest)(nil),         // 0: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.GetAlertSchedulerRuleRequest
	(*CreateAlertSchedulerRuleRequest)(nil),      // 1: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.CreateAlertSchedulerRuleRequest
	(*UpdateAlertSchedulerRuleRequest)(nil),      // 2: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.UpdateAlertSchedulerRuleRequest
	(*DeleteAlertSchedulerRuleRequest)(nil),      // 3: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.DeleteAlertSchedulerRuleRequest
	(*GetBulkAlertSchedulerRuleRequest)(nil),     // 4: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.GetBulkAlertSchedulerRuleRequest
	(*CreateBulkAlertSchedulerRuleRequest)(nil),  // 5: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.CreateBulkAlertSchedulerRuleRequest
	(*UpdateBulkAlertSchedulerRuleRequest)(nil),  // 6: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.UpdateBulkAlertSchedulerRuleRequest
	(*GetAlertSchedulerRuleResponse)(nil),        // 7: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.GetAlertSchedulerRuleResponse
	(*CreateAlertSchedulerRuleResponse)(nil),     // 8: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.CreateAlertSchedulerRuleResponse
	(*UpdateAlertSchedulerRuleResponse)(nil),     // 9: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.UpdateAlertSchedulerRuleResponse
	(*DeleteAlertSchedulerRuleResponse)(nil),     // 10: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.DeleteAlertSchedulerRuleResponse
	(*GetBulkAlertSchedulerRuleResponse)(nil),    // 11: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.GetBulkAlertSchedulerRuleResponse
	(*CreateBulkAlertSchedulerRuleResponse)(nil), // 12: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.CreateBulkAlertSchedulerRuleResponse
	(*UpdateBulkAlertSchedulerRuleResponse)(nil), // 13: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.UpdateBulkAlertSchedulerRuleResponse
}
var file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_service_proto_depIdxs = []int32{
	0,  // 0: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService.GetAlertSchedulerRule:input_type -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.GetAlertSchedulerRuleRequest
	1,  // 1: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService.CreateAlertSchedulerRule:input_type -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.CreateAlertSchedulerRuleRequest
	2,  // 2: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService.UpdateAlertSchedulerRule:input_type -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.UpdateAlertSchedulerRuleRequest
	3,  // 3: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService.DeleteAlertSchedulerRule:input_type -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.DeleteAlertSchedulerRuleRequest
	4,  // 4: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService.GetBulkAlertSchedulerRule:input_type -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.GetBulkAlertSchedulerRuleRequest
	5,  // 5: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService.CreateBulkAlertSchedulerRule:input_type -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.CreateBulkAlertSchedulerRuleRequest
	6,  // 6: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService.UpdateBulkAlertSchedulerRule:input_type -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.UpdateBulkAlertSchedulerRuleRequest
	7,  // 7: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService.GetAlertSchedulerRule:output_type -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.GetAlertSchedulerRuleResponse
	8,  // 8: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService.CreateAlertSchedulerRule:output_type -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.CreateAlertSchedulerRuleResponse
	9,  // 9: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService.UpdateAlertSchedulerRule:output_type -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.UpdateAlertSchedulerRuleResponse
	10, // 10: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService.DeleteAlertSchedulerRule:output_type -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.DeleteAlertSchedulerRuleResponse
	11, // 11: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService.GetBulkAlertSchedulerRule:output_type -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.GetBulkAlertSchedulerRuleResponse
	12, // 12: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService.CreateBulkAlertSchedulerRule:output_type -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.CreateBulkAlertSchedulerRuleResponse
	13, // 13: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService.UpdateBulkAlertSchedulerRule:output_type -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.UpdateBulkAlertSchedulerRuleResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_service_proto_init()
}
func file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_service_proto_init() {
	if File_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_service_proto != nil {
		return
	}
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_init()
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_response_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_service_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_service_proto_depIdxs,
	}.Build()
	File_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_service_proto = out.File
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_service_proto_rawDesc = nil
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_service_proto_goTypes = nil
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_service_proto_depIdxs = nil
}
