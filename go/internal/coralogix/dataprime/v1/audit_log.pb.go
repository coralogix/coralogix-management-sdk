// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: com/coralogix/dataprime/v1/audit_log.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuditLogDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description *string `protobuf:"bytes,1,opt,name=description,proto3,oneof" json:"description,omitempty"`
}

func (x *AuditLogDescription) Reset() {
	*x = AuditLogDescription{}
	mi := &file_com_coralogix_dataprime_v1_audit_log_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuditLogDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditLogDescription) ProtoMessage() {}

func (x *AuditLogDescription) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_audit_log_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditLogDescription.ProtoReflect.Descriptor instead.
func (*AuditLogDescription) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_audit_log_proto_rawDescGZIP(), []int{0}
}

func (x *AuditLogDescription) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

var file_com_coralogix_dataprime_v1_audit_log_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*AuditLogDescription)(nil),
		Field:         5000,
		Name:          "com.coralogix.dataprime.v1.audit_log_description",
		Tag:           "bytes,5000,opt,name=audit_log_description",
		Filename:      "com/coralogix/dataprime/v1/audit_log.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional com.coralogix.dataprime.v1.AuditLogDescription audit_log_description = 5000;
	E_AuditLogDescription = &file_com_coralogix_dataprime_v1_audit_log_proto_extTypes[0]
)

var File_com_coralogix_dataprime_v1_audit_log_proto protoreflect.FileDescriptor

var file_com_coralogix_dataprime_v1_audit_log_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x13, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x84, 0x01, 0x0a, 0x15, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x88, 0x27, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x13, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x4c, 0x6f, 0x67, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_dataprime_v1_audit_log_proto_rawDescOnce sync.Once
	file_com_coralogix_dataprime_v1_audit_log_proto_rawDescData = file_com_coralogix_dataprime_v1_audit_log_proto_rawDesc
)

func file_com_coralogix_dataprime_v1_audit_log_proto_rawDescGZIP() []byte {
	file_com_coralogix_dataprime_v1_audit_log_proto_rawDescOnce.Do(func() {
		file_com_coralogix_dataprime_v1_audit_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_dataprime_v1_audit_log_proto_rawDescData)
	})
	return file_com_coralogix_dataprime_v1_audit_log_proto_rawDescData
}

var file_com_coralogix_dataprime_v1_audit_log_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogix_dataprime_v1_audit_log_proto_goTypes = []any{
	(*AuditLogDescription)(nil),        // 0: com.coralogix.dataprime.v1.AuditLogDescription
	(*descriptorpb.MethodOptions)(nil), // 1: google.protobuf.MethodOptions
}
var file_com_coralogix_dataprime_v1_audit_log_proto_depIdxs = []int32{
	1, // 0: com.coralogix.dataprime.v1.audit_log_description:extendee -> google.protobuf.MethodOptions
	0, // 1: com.coralogix.dataprime.v1.audit_log_description:type_name -> com.coralogix.dataprime.v1.AuditLogDescription
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogix_dataprime_v1_audit_log_proto_init() }
func file_com_coralogix_dataprime_v1_audit_log_proto_init() {
	if File_com_coralogix_dataprime_v1_audit_log_proto != nil {
		return
	}
	file_com_coralogix_dataprime_v1_audit_log_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_dataprime_v1_audit_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_dataprime_v1_audit_log_proto_goTypes,
		DependencyIndexes: file_com_coralogix_dataprime_v1_audit_log_proto_depIdxs,
		MessageInfos:      file_com_coralogix_dataprime_v1_audit_log_proto_msgTypes,
		ExtensionInfos:    file_com_coralogix_dataprime_v1_audit_log_proto_extTypes,
	}.Build()
	File_com_coralogix_dataprime_v1_audit_log_proto = out.File
	file_com_coralogix_dataprime_v1_audit_log_proto_rawDesc = nil
	file_com_coralogix_dataprime_v1_audit_log_proto_goTypes = nil
	file_com_coralogix_dataprime_v1_audit_log_proto_depIdxs = nil
}
