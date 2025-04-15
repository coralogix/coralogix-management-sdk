// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
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
	unsafe "unsafe"
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

const file_com_coralogixapis_alerts_v3_alert_def_severity_proto_rawDesc = "" +
	"\n" +
	"4com/coralogixapis/alerts/v3/alert_def_severity.proto\x12\x1bcom.coralogixapis.alerts.v3\x1a\x1egoogle/protobuf/wrappers.proto\x1a4com/coralogixapis/alerts/v3/alert_def_priority.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\x8b\x02\n" +
	"\x10AlertDefSeverity\x12?\n" +
	"\bseverity\x18\x01 \x01(\v2\x1b.google.protobuf.Int32ValueB\x06\x92A\x03J\x013R\bseverity\x12I\n" +
	"\bpriority\x18\x02 \x01(\x0e2-.com.coralogixapis.alerts.v3.AlertDefPriorityR\bpriority:k\x92Ah\n" +
	"f*\x19Alert definition severity23Defines the severity level and priority of an alert\xd2\x01\bseverity\xd2\x01\bpriorityb\x06proto3"

var (
	file_com_coralogixapis_alerts_v3_alert_def_severity_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_severity_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v3_alert_def_severity_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_severity_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_severity_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_severity_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_severity_proto_rawDesc)))
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_severity_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_severity_proto_rawDesc)),
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
	file_com_coralogixapis_alerts_v3_alert_def_severity_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_severity_proto_depIdxs = nil
}
