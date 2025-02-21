// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: com/coralogixapis/notification_center/common/rpc_common.proto

package common

import (
	code "google.golang.org/genproto/googleapis/rpc/code"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResponseStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode code.Code         `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3,enum=google.rpc.Code" json:"status_code,omitempty"`
	Message    *string           `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"`
	Details    map[string]string `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ResponseStatus) Reset() {
	*x = ResponseStatus{}
	mi := &file_com_coralogixapis_notification_center_common_rpc_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResponseStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseStatus) ProtoMessage() {}

func (x *ResponseStatus) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_common_rpc_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseStatus.ProtoReflect.Descriptor instead.
func (*ResponseStatus) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_common_rpc_common_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseStatus) GetStatusCode() code.Code {
	if x != nil {
		return x.StatusCode
	}
	return code.Code(0)
}

func (x *ResponseStatus) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

func (x *ResponseStatus) GetDetails() map[string]string {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_com_coralogixapis_notification_center_common_rpc_common_proto protoreflect.FileDescriptor

var file_com_coralogixapis_notification_center_common_rpc_common_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72,
	0x70, 0x63, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x1a, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72,
	0x70, 0x63, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x02,
	0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x31, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x5c, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x1a, 0x3a, 0x0a, 0x0c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_notification_center_common_rpc_common_proto_rawDescOnce sync.Once
	file_com_coralogixapis_notification_center_common_rpc_common_proto_rawDescData = file_com_coralogixapis_notification_center_common_rpc_common_proto_rawDesc
)

func file_com_coralogixapis_notification_center_common_rpc_common_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_notification_center_common_rpc_common_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_notification_center_common_rpc_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_notification_center_common_rpc_common_proto_rawDescData)
	})
	return file_com_coralogixapis_notification_center_common_rpc_common_proto_rawDescData
}

var file_com_coralogixapis_notification_center_common_rpc_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogixapis_notification_center_common_rpc_common_proto_goTypes = []any{
	(*ResponseStatus)(nil), // 0: com.coralogixapis.notification_center.ResponseStatus
	nil,                    // 1: com.coralogixapis.notification_center.ResponseStatus.DetailsEntry
	(code.Code)(0),         // 2: google.rpc.Code
}
var file_com_coralogixapis_notification_center_common_rpc_common_proto_depIdxs = []int32{
	2, // 0: com.coralogixapis.notification_center.ResponseStatus.status_code:type_name -> google.rpc.Code
	1, // 1: com.coralogixapis.notification_center.ResponseStatus.details:type_name -> com.coralogixapis.notification_center.ResponseStatus.DetailsEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_notification_center_common_rpc_common_proto_init() }
func file_com_coralogixapis_notification_center_common_rpc_common_proto_init() {
	if File_com_coralogixapis_notification_center_common_rpc_common_proto != nil {
		return
	}
	file_com_coralogixapis_notification_center_common_rpc_common_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_notification_center_common_rpc_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_notification_center_common_rpc_common_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_notification_center_common_rpc_common_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_notification_center_common_rpc_common_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_notification_center_common_rpc_common_proto = out.File
	file_com_coralogixapis_notification_center_common_rpc_common_proto_rawDesc = nil
	file_com_coralogixapis_notification_center_common_rpc_common_proto_goTypes = nil
	file_com_coralogixapis_notification_center_common_rpc_common_proto_depIdxs = nil
}
