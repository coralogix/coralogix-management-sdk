// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/tracing/tracing_immediate_type_definition.proto

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

type TracingImmediateType struct {
	state                     protoimpl.MessageState    `protogen:"open.v1"`
	TracingFilter             *TracingFilter            `protobuf:"bytes,1,opt,name=tracing_filter,json=tracingFilter,proto3" json:"tracing_filter,omitempty"`
	NotificationPayloadFilter []*wrapperspb.StringValue `protobuf:"bytes,2,rep,name=notification_payload_filter,json=notificationPayloadFilter,proto3" json:"notification_payload_filter,omitempty"`
	unknownFields             protoimpl.UnknownFields
	sizeCache                 protoimpl.SizeCache
}

func (x *TracingImmediateType) Reset() {
	*x = TracingImmediateType{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_tracing_immediate_type_definition_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TracingImmediateType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TracingImmediateType) ProtoMessage() {}

func (x *TracingImmediateType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_tracing_immediate_type_definition_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TracingImmediateType.ProtoReflect.Descriptor instead.
func (*TracingImmediateType) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_tracing_immediate_type_definition_proto_rawDescGZIP(), []int{0}
}

func (x *TracingImmediateType) GetTracingFilter() *TracingFilter {
	if x != nil {
		return x.TracingFilter
	}
	return nil
}

func (x *TracingImmediateType) GetNotificationPayloadFilter() []*wrapperspb.StringValue {
	if x != nil {
		return x.NotificationPayloadFilter
	}
	return nil
}

var File_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_tracing_immediate_type_definition_proto protoreflect.FileDescriptor

const file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_tracing_immediate_type_definition_proto_rawDesc = "" +
	"\n" +
	"ecom/coralogixapis/alerts/v3/alert_def_type_definition/tracing/tracing_immediate_type_definition.proto\x12\x1bcom.coralogixapis.alerts.v3\x1a\x1egoogle/protobuf/wrappers.proto\x1aYcom/coralogixapis/alerts/v3/alert_def_type_definition/tracing/common/tracing_filter.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\xbc\x02\n" +
	"\x14TracingImmediateType\x12Q\n" +
	"\x0etracing_filter\x18\x01 \x01(\v2*.com.coralogixapis.alerts.v3.TracingFilterR\rtracingFilter\x12\\\n" +
	"\x1bnotification_payload_filter\x18\x02 \x03(\v2\x1c.google.protobuf.StringValueR\x19notificationPayloadFilter:s\x92Ap\n" +
	"n*\x1cTracing Immediate Alert Type2=Configuration for immediate alerts triggered on trace entries\xd2\x01\x0etracing_filterb\x06proto3"

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_tracing_immediate_type_definition_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_tracing_immediate_type_definition_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_tracing_immediate_type_definition_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_tracing_immediate_type_definition_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_tracing_immediate_type_definition_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_tracing_immediate_type_definition_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_tracing_immediate_type_definition_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_tracing_immediate_type_definition_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_tracing_immediate_type_definition_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_tracing_immediate_type_definition_proto_goTypes = []any{
	(*TracingImmediateType)(nil),   // 0: com.coralogixapis.alerts.v3.TracingImmediateType
	(*TracingFilter)(nil),          // 1: com.coralogixapis.alerts.v3.TracingFilter
	(*wrapperspb.StringValue)(nil), // 2: google.protobuf.StringValue
}
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_tracing_immediate_type_definition_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.alerts.v3.TracingImmediateType.tracing_filter:type_name -> com.coralogixapis.alerts.v3.TracingFilter
	2, // 1: com.coralogixapis.alerts.v3.TracingImmediateType.notification_payload_filter:type_name -> google.protobuf.StringValue
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_tracing_immediate_type_definition_proto_init()
}
func file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_tracing_immediate_type_definition_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_tracing_immediate_type_definition_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_tracing_immediate_type_definition_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_tracing_immediate_type_definition_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_tracing_immediate_type_definition_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_tracing_immediate_type_definition_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_tracing_immediate_type_definition_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_tracing_immediate_type_definition_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_tracing_immediate_type_definition_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_tracing_immediate_type_definition_proto_depIdxs = nil
}
