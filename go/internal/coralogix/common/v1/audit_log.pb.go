// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogix/common/v1/audit_log.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuditLogDescription struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Description   string                 `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuditLogDescription) Reset() {
	*x = AuditLogDescription{}
	mi := &file_com_coralogix_common_v1_audit_log_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuditLogDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditLogDescription) ProtoMessage() {}

func (x *AuditLogDescription) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_common_v1_audit_log_proto_msgTypes[0]
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
	return file_com_coralogix_common_v1_audit_log_proto_rawDescGZIP(), []int{0}
}

func (x *AuditLogDescription) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var file_com_coralogix_common_v1_audit_log_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*AuditLogDescription)(nil),
		Field:         5000,
		Name:          "com.coralogix.common.v1.audit_log_description",
		Tag:           "bytes,5000,opt,name=audit_log_description",
		Filename:      "com/coralogix/common/v1/audit_log.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional com.coralogix.common.v1.AuditLogDescription audit_log_description = 5000;
	E_AuditLogDescription = &file_com_coralogix_common_v1_audit_log_proto_extTypes[0]
)

var File_com_coralogix_common_v1_audit_log_proto protoreflect.FileDescriptor

const file_com_coralogix_common_v1_audit_log_proto_rawDesc = "" +
	"\n" +
	"'com/coralogix/common/v1/audit_log.proto\x12\x17com.coralogix.common.v1\x1a google/protobuf/descriptor.proto\"7\n" +
	"\x13AuditLogDescription\x12 \n" +
	"\vdescription\x18\x01 \x01(\tR\vdescription:\x81\x01\n" +
	"\x15audit_log_description\x12\x1e.google.protobuf.MethodOptions\x18\x88' \x01(\v2,.com.coralogix.common.v1.AuditLogDescriptionR\x13auditLogDescriptionb\x06proto3"

var (
	file_com_coralogix_common_v1_audit_log_proto_rawDescOnce sync.Once
	file_com_coralogix_common_v1_audit_log_proto_rawDescData []byte
)

func file_com_coralogix_common_v1_audit_log_proto_rawDescGZIP() []byte {
	file_com_coralogix_common_v1_audit_log_proto_rawDescOnce.Do(func() {
		file_com_coralogix_common_v1_audit_log_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogix_common_v1_audit_log_proto_rawDesc), len(file_com_coralogix_common_v1_audit_log_proto_rawDesc)))
	})
	return file_com_coralogix_common_v1_audit_log_proto_rawDescData
}

var file_com_coralogix_common_v1_audit_log_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogix_common_v1_audit_log_proto_goTypes = []any{
	(*AuditLogDescription)(nil),        // 0: com.coralogix.common.v1.AuditLogDescription
	(*descriptorpb.MethodOptions)(nil), // 1: google.protobuf.MethodOptions
}
var file_com_coralogix_common_v1_audit_log_proto_depIdxs = []int32{
	1, // 0: com.coralogix.common.v1.audit_log_description:extendee -> google.protobuf.MethodOptions
	0, // 1: com.coralogix.common.v1.audit_log_description:type_name -> com.coralogix.common.v1.AuditLogDescription
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogix_common_v1_audit_log_proto_init() }
func file_com_coralogix_common_v1_audit_log_proto_init() {
	if File_com_coralogix_common_v1_audit_log_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogix_common_v1_audit_log_proto_rawDesc), len(file_com_coralogix_common_v1_audit_log_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_common_v1_audit_log_proto_goTypes,
		DependencyIndexes: file_com_coralogix_common_v1_audit_log_proto_depIdxs,
		MessageInfos:      file_com_coralogix_common_v1_audit_log_proto_msgTypes,
		ExtensionInfos:    file_com_coralogix_common_v1_audit_log_proto_extTypes,
	}.Build()
	File_com_coralogix_common_v1_audit_log_proto = out.File
	file_com_coralogix_common_v1_audit_log_proto_goTypes = nil
	file_com_coralogix_common_v1_audit_log_proto_depIdxs = nil
}
