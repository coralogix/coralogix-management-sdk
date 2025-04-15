// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/notification_center/common/rpc_common.proto

package common

import (
	code "google.golang.org/genproto/googleapis/rpc/code"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type ResponseStatus struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StatusCode    code.Code              `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3,enum=google.rpc.Code" json:"status_code,omitempty"`
	Message       *string                `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"`
	Details       map[string]string      `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

const file_com_coralogixapis_notification_center_common_rpc_common_proto_rawDesc = "" +
	"\n" +
	"=com/coralogixapis/notification_center/common/rpc_common.proto\x12%com.coralogixapis.notification_center\x1a\x15google/rpc/code.proto\"\x88\x02\n" +
	"\x0eResponseStatus\x121\n" +
	"\vstatus_code\x18\x01 \x01(\x0e2\x10.google.rpc.CodeR\n" +
	"statusCode\x12\x1d\n" +
	"\amessage\x18\x02 \x01(\tH\x00R\amessage\x88\x01\x01\x12\\\n" +
	"\adetails\x18\x03 \x03(\v2B.com.coralogixapis.notification_center.ResponseStatus.DetailsEntryR\adetails\x1a:\n" +
	"\fDetailsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01B\n" +
	"\n" +
	"\b_messageb\x06proto3"

var (
	file_com_coralogixapis_notification_center_common_rpc_common_proto_rawDescOnce sync.Once
	file_com_coralogixapis_notification_center_common_rpc_common_proto_rawDescData []byte
)

func file_com_coralogixapis_notification_center_common_rpc_common_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_notification_center_common_rpc_common_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_notification_center_common_rpc_common_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_notification_center_common_rpc_common_proto_rawDesc), len(file_com_coralogixapis_notification_center_common_rpc_common_proto_rawDesc)))
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_notification_center_common_rpc_common_proto_rawDesc), len(file_com_coralogixapis_notification_center_common_rpc_common_proto_rawDesc)),
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
	file_com_coralogixapis_notification_center_common_rpc_common_proto_goTypes = nil
	file_com_coralogixapis_notification_center_common_rpc_common_proto_depIdxs = nil
}
