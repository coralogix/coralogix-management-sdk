// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.26.1
// source: com/coralogixapis/notification_center/common/v1/config_fields.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type ConnectorConfigField struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FieldName     string                 `protobuf:"bytes,1,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	Template      string                 `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConnectorConfigField) Reset() {
	*x = ConnectorConfigField{}
	mi := &file_com_coralogixapis_notification_center_common_v1_config_fields_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectorConfigField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectorConfigField) ProtoMessage() {}

func (x *ConnectorConfigField) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_common_v1_config_fields_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectorConfigField.ProtoReflect.Descriptor instead.
func (*ConnectorConfigField) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_common_v1_config_fields_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectorConfigField) GetFieldName() string {
	if x != nil {
		return x.FieldName
	}
	return ""
}

func (x *ConnectorConfigField) GetTemplate() string {
	if x != nil {
		return x.Template
	}
	return ""
}

type MessageConfigField struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FieldName     string                 `protobuf:"bytes,1,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	Template      string                 `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageConfigField) Reset() {
	*x = MessageConfigField{}
	mi := &file_com_coralogixapis_notification_center_common_v1_config_fields_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageConfigField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageConfigField) ProtoMessage() {}

func (x *MessageConfigField) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_common_v1_config_fields_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageConfigField.ProtoReflect.Descriptor instead.
func (*MessageConfigField) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_common_v1_config_fields_proto_rawDescGZIP(), []int{1}
}

func (x *MessageConfigField) GetFieldName() string {
	if x != nil {
		return x.FieldName
	}
	return ""
}

func (x *MessageConfigField) GetTemplate() string {
	if x != nil {
		return x.Template
	}
	return ""
}

type RenderedConnectorConfigField struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FieldName     string                 `protobuf:"bytes,1,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RenderedConnectorConfigField) Reset() {
	*x = RenderedConnectorConfigField{}
	mi := &file_com_coralogixapis_notification_center_common_v1_config_fields_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RenderedConnectorConfigField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenderedConnectorConfigField) ProtoMessage() {}

func (x *RenderedConnectorConfigField) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_common_v1_config_fields_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenderedConnectorConfigField.ProtoReflect.Descriptor instead.
func (*RenderedConnectorConfigField) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_common_v1_config_fields_proto_rawDescGZIP(), []int{2}
}

func (x *RenderedConnectorConfigField) GetFieldName() string {
	if x != nil {
		return x.FieldName
	}
	return ""
}

func (x *RenderedConnectorConfigField) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type RenderedMessageConfigField struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FieldName     string                 `protobuf:"bytes,1,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RenderedMessageConfigField) Reset() {
	*x = RenderedMessageConfigField{}
	mi := &file_com_coralogixapis_notification_center_common_v1_config_fields_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RenderedMessageConfigField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenderedMessageConfigField) ProtoMessage() {}

func (x *RenderedMessageConfigField) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_common_v1_config_fields_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenderedMessageConfigField.ProtoReflect.Descriptor instead.
func (*RenderedMessageConfigField) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_common_v1_config_fields_proto_rawDescGZIP(), []int{3}
}

func (x *RenderedMessageConfigField) GetFieldName() string {
	if x != nil {
		return x.FieldName
	}
	return ""
}

func (x *RenderedMessageConfigField) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_com_coralogixapis_notification_center_common_v1_config_fields_proto protoreflect.FileDescriptor

var file_com_coralogixapis_notification_center_common_v1_config_fields_proto_rawDesc = []byte{
	0x0a, 0x43, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x14,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22,
	0xc3, 0x03, 0x0a, 0x12, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x2b, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0x92, 0x41, 0x09, 0x4a,
	0x07, 0x22, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x5d, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x41, 0x92, 0x41, 0x3e, 0x4a, 0x3c, 0x22, 0x7b, 0x7b, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x7d, 0x7d, 0x20, 0x7b, 0x7b,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x44, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x7d, 0x7d, 0x20, 0x2d, 0x20, 0x7b, 0x7b, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x44, 0x65, 0x66,
	0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x7d, 0x22, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x3a, 0xa0, 0x02, 0x92, 0x41, 0x9c, 0x02, 0x0a, 0x9c, 0x01, 0x2a, 0x14, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x20, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x32, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x20, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x73, 0x20, 0x61, 0x20, 0x77, 0x61, 0x79, 0x20, 0x74, 0x6f, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x65, 0x20, 0x61, 0x20, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x20, 0x74, 0x68, 0x61,
	0x74, 0x20, 0x63, 0x61, 0x6e, 0x20, 0x62, 0x65, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x74, 0x6f,
	0x20, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0xd2, 0x01, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0xd2, 0x01, 0x08,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2a, 0x7b, 0x0a, 0x27, 0x46, 0x69, 0x6e, 0x64,
	0x20, 0x6f, 0x75, 0x74, 0x20, 0x6d, 0x6f, 0x72, 0x65, 0x20, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x20,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x12, 0x50, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x63, 0x73, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2d, 0x67, 0x75, 0x69, 0x64, 0x65, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f,
	0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x77, 0x65, 0x6c,
	0x63, 0x6f, 0x6d, 0x65, 0x2f, 0x22, 0x53, 0x0a, 0x1c, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65,
	0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x51, 0x0a, 0x1a, 0x52, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_notification_center_common_v1_config_fields_proto_rawDescOnce sync.Once
	file_com_coralogixapis_notification_center_common_v1_config_fields_proto_rawDescData = file_com_coralogixapis_notification_center_common_v1_config_fields_proto_rawDesc
)

func file_com_coralogixapis_notification_center_common_v1_config_fields_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_notification_center_common_v1_config_fields_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_notification_center_common_v1_config_fields_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_notification_center_common_v1_config_fields_proto_rawDescData)
	})
	return file_com_coralogixapis_notification_center_common_v1_config_fields_proto_rawDescData
}

var file_com_coralogixapis_notification_center_common_v1_config_fields_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_com_coralogixapis_notification_center_common_v1_config_fields_proto_goTypes = []any{
	(*ConnectorConfigField)(nil),         // 0: com.coralogixapis.notification_center.ConnectorConfigField
	(*MessageConfigField)(nil),           // 1: com.coralogixapis.notification_center.MessageConfigField
	(*RenderedConnectorConfigField)(nil), // 2: com.coralogixapis.notification_center.RenderedConnectorConfigField
	(*RenderedMessageConfigField)(nil),   // 3: com.coralogixapis.notification_center.RenderedMessageConfigField
}
var file_com_coralogixapis_notification_center_common_v1_config_fields_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_notification_center_common_v1_config_fields_proto_init() }
func file_com_coralogixapis_notification_center_common_v1_config_fields_proto_init() {
	if File_com_coralogixapis_notification_center_common_v1_config_fields_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_notification_center_common_v1_config_fields_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_notification_center_common_v1_config_fields_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_notification_center_common_v1_config_fields_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_notification_center_common_v1_config_fields_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_notification_center_common_v1_config_fields_proto = out.File
	file_com_coralogixapis_notification_center_common_v1_config_fields_proto_rawDesc = nil
	file_com_coralogixapis_notification_center_common_v1_config_fields_proto_goTypes = nil
	file_com_coralogixapis_notification_center_common_v1_config_fields_proto_depIdxs = nil
}
