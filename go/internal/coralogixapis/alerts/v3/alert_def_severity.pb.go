// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: com/coralogixapis/alerts/v3/alert_def_severity.proto

package v3

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AlertDefSeverity struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Severity      *wrapperspb.Int32Value `protobuf:"bytes,1,opt,name=severity,proto3" json:"severity,omitempty"`
	Priority      AlertDefPriority       `protobuf:"varint,2,opt,name=priority,proto3,enum=com.coralogixapis.alerts.v3.AlertDefPriority" json:"priority,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AlertDefSeverity) Reset() {
	*x = AlertDefSeverity{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_severity_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlertDefSeverity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertDefSeverity) ProtoMessage() {}

func (x *AlertDefSeverity) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_severity_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertDefSeverity.ProtoReflect.Descriptor instead.
func (*AlertDefSeverity) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_severity_proto_rawDescGZIP(), []int{0}
}

func (x *AlertDefSeverity) GetSeverity() *wrapperspb.Int32Value {
	if x != nil {
		return x.Severity
	}
	return nil
}

func (x *AlertDefSeverity) GetPriority() AlertDefPriority {
	if x != nil {
		return x.Priority
	}
	return AlertDefPriority_ALERT_DEF_PRIORITY_P5_OR_UNSPECIFIED
}

var File_com_coralogixapis_alerts_v3_alert_def_severity_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_alert_def_severity_proto_rawDesc = []byte{
	0x0a, 0x34, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x2e, 0x76, 0x33, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33,
	0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x02, 0x0a, 0x10, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x44, 0x65, 0x66, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x3f,
	0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x06, 0x92,
	0x41, 0x03, 0x4a, 0x01, 0x33, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12,
	0x49, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x44, 0x65, 0x66, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x3a, 0x6b, 0x92, 0x41, 0x68, 0x0a,
	0x66, 0x2a, 0x19, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x20, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x32, 0x33, 0x44, 0x65,
	0x66, 0x69, 0x6e, 0x65, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69,
	0x74, 0x79, 0x20, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x70, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x6e, 0x20, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0xd2, 0x01, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0xd2, 0x01, 0x08, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_alert_def_severity_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_severity_proto_rawDescData = file_com_coralogixapis_alerts_v3_alert_def_severity_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_alert_def_severity_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_severity_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_severity_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_alert_def_severity_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_alert_def_severity_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_severity_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_alerts_v3_alert_def_severity_proto_goTypes = []any{
	(*AlertDefSeverity)(nil),      // 0: com.coralogixapis.alerts.v3.AlertDefSeverity
	(*wrapperspb.Int32Value)(nil), // 1: google.protobuf.Int32Value
	(AlertDefPriority)(0),         // 2: com.coralogixapis.alerts.v3.AlertDefPriority
}
var file_com_coralogixapis_alerts_v3_alert_def_severity_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.alerts.v3.AlertDefSeverity.severity:type_name -> google.protobuf.Int32Value
	2, // 1: com.coralogixapis.alerts.v3.AlertDefSeverity.priority:type_name -> com.coralogixapis.alerts.v3.AlertDefPriority
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_alerts_v3_alert_def_severity_proto_init() }
func file_com_coralogixapis_alerts_v3_alert_def_severity_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_severity_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_alert_def_priority_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_alert_def_severity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_severity_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_severity_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_alert_def_severity_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_severity_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_severity_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_alert_def_severity_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_severity_proto_depIdxs = nil
}
