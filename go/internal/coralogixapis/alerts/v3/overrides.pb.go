// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/overrides.proto

package v3

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

type AlertDefOverride struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Priority      AlertDefPriority       `protobuf:"varint,1,opt,name=priority,proto3,enum=com.coralogixapis.alerts.v3.AlertDefPriority" json:"priority,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AlertDefOverride) Reset() {
	*x = AlertDefOverride{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlertDefOverride) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertDefOverride) ProtoMessage() {}

func (x *AlertDefOverride) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertDefOverride.ProtoReflect.Descriptor instead.
func (*AlertDefOverride) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto_rawDescGZIP(), []int{0}
}

func (x *AlertDefOverride) GetPriority() AlertDefPriority {
	if x != nil {
		return x.Priority
	}
	return AlertDefPriority_ALERT_DEF_PRIORITY_P5_OR_UNSPECIFIED
}

var File_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto_rawDesc = []byte{
	0x0a, 0x45, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x73, 0x2e, 0x76, 0x33, 0x1a, 0x34, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76,
	0x33, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x70, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x10, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x44, 0x65, 0x66, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x12,
	0x49, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x44, 0x65, 0x66, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x3a, 0x27, 0x92, 0x41, 0x24, 0x0a,
	0x22, 0x2a, 0x20, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x20, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x20, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto_rawDescData = file_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto_goTypes = []any{
	(*AlertDefOverride)(nil), // 0: com.coralogixapis.alerts.v3.AlertDefOverride
	(AlertDefPriority)(0),    // 1: com.coralogixapis.alerts.v3.AlertDefPriority
}
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.alerts.v3.AlertDefOverride.priority:type_name -> com.coralogixapis.alerts.v3.AlertDefPriority
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto_init() }
func file_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_alert_def_priority_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto_depIdxs = nil
}
