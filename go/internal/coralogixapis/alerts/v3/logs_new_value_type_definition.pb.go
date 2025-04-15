// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/logs/logs_new_value_type_definition.proto

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

type LogsNewValueType struct {
	state                     protoimpl.MessageState    `protogen:"open.v1"`
	LogsFilter                *LogsFilter               `protobuf:"bytes,1,opt,name=logs_filter,json=logsFilter,proto3" json:"logs_filter,omitempty"`
	Rules                     []*LogsNewValueRule       `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty"`
	NotificationPayloadFilter []*wrapperspb.StringValue `protobuf:"bytes,3,rep,name=notification_payload_filter,json=notificationPayloadFilter,proto3" json:"notification_payload_filter,omitempty"`
	unknownFields             protoimpl.UnknownFields
	sizeCache                 protoimpl.SizeCache
}

func (x *LogsNewValueType) Reset() {
	*x = LogsNewValueType{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsNewValueType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsNewValueType) ProtoMessage() {}

func (x *LogsNewValueType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsNewValueType.ProtoReflect.Descriptor instead.
func (*LogsNewValueType) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_rawDescGZIP(), []int{0}
}

func (x *LogsNewValueType) GetLogsFilter() *LogsFilter {
	if x != nil {
		return x.LogsFilter
	}
	return nil
}

func (x *LogsNewValueType) GetRules() []*LogsNewValueRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *LogsNewValueType) GetNotificationPayloadFilter() []*wrapperspb.StringValue {
	if x != nil {
		return x.NotificationPayloadFilter
	}
	return nil
}

type LogsNewValueRule struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Condition     *LogsNewValueCondition `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogsNewValueRule) Reset() {
	*x = LogsNewValueRule{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsNewValueRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsNewValueRule) ProtoMessage() {}

func (x *LogsNewValueRule) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsNewValueRule.ProtoReflect.Descriptor instead.
func (*LogsNewValueRule) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_rawDescGZIP(), []int{1}
}

func (x *LogsNewValueRule) GetCondition() *LogsNewValueCondition {
	if x != nil {
		return x.Condition
	}
	return nil
}

type LogsNewValueCondition struct {
	state          protoimpl.MessageState  `protogen:"open.v1"`
	KeypathToTrack *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=keypath_to_track,json=keypathToTrack,proto3" json:"keypath_to_track,omitempty"`
	TimeWindow     *LogsNewValueTimeWindow `protobuf:"bytes,2,opt,name=time_window,json=timeWindow,proto3" json:"time_window,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *LogsNewValueCondition) Reset() {
	*x = LogsNewValueCondition{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsNewValueCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsNewValueCondition) ProtoMessage() {}

func (x *LogsNewValueCondition) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsNewValueCondition.ProtoReflect.Descriptor instead.
func (*LogsNewValueCondition) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_rawDescGZIP(), []int{2}
}

func (x *LogsNewValueCondition) GetKeypathToTrack() *wrapperspb.StringValue {
	if x != nil {
		return x.KeypathToTrack
	}
	return nil
}

func (x *LogsNewValueCondition) GetTimeWindow() *LogsNewValueTimeWindow {
	if x != nil {
		return x.TimeWindow
	}
	return nil
}

var File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto protoreflect.FileDescriptor

const file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_rawDesc = "" +
	"\n" +
	"_com/coralogixapis/alerts/v3/alert_def_type_definition/logs/logs_new_value_type_definition.proto\x12\x1bcom.coralogixapis.alerts.v3\x1a\x1egoogle/protobuf/wrappers.proto\x1adcom/coralogixapis/alerts/v3/alert_def_type_definition/logs/new_value/logs_new_value_timewindow.proto\x1aScom/coralogixapis/alerts/v3/alert_def_type_definition/logs/common/logs_filter.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\x8d\x03\n" +
	"\x10LogsNewValueType\x12H\n" +
	"\vlogs_filter\x18\x01 \x01(\v2'.com.coralogixapis.alerts.v3.LogsFilterR\n" +
	"logsFilter\x12C\n" +
	"\x05rules\x18\x02 \x03(\v2-.com.coralogixapis.alerts.v3.LogsNewValueRuleR\x05rules\x12i\n" +
	"\x1bnotification_payload_filter\x18\x03 \x03(\v2\x1c.google.protobuf.StringValueB\v\x92A\bJ\x06\"text\"R\x19notificationPayloadFilter:\x7f\x92A|\n" +
	"z*\x1eLog-based new value alert type2BConfiguration for alerts triggered by new values appearing in logs\xd2\x01\vlogs_filter\xd2\x01\x05rules\"d\n" +
	"\x10LogsNewValueRule\x12P\n" +
	"\tcondition\x18\x01 \x01(\v22.com.coralogixapis.alerts.v3.LogsNewValueConditionR\tcondition\"\xc8\x02\n" +
	"\x15LogsNewValueCondition\x12]\n" +
	"\x10keypath_to_track\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueB\x15\x92A\x12J\x10\"metadata.field\"R\x0ekeypathToTrack\x12T\n" +
	"\vtime_window\x18\x02 \x01(\v23.com.coralogixapis.alerts.v3.LogsNewValueTimeWindowR\n" +
	"timeWindow:z\x92Aw\n" +
	"u*\x1dLog-based new value condition23Defines conditions for detecting new values in logs\xd2\x01\x10keypath_to_track\xd2\x01\vtime_windowb\x06proto3"

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_goTypes = []any{
	(*LogsNewValueType)(nil),       // 0: com.coralogixapis.alerts.v3.LogsNewValueType
	(*LogsNewValueRule)(nil),       // 1: com.coralogixapis.alerts.v3.LogsNewValueRule
	(*LogsNewValueCondition)(nil),  // 2: com.coralogixapis.alerts.v3.LogsNewValueCondition
	(*LogsFilter)(nil),             // 3: com.coralogixapis.alerts.v3.LogsFilter
	(*wrapperspb.StringValue)(nil), // 4: google.protobuf.StringValue
	(*LogsNewValueTimeWindow)(nil), // 5: com.coralogixapis.alerts.v3.LogsNewValueTimeWindow
}
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_depIdxs = []int32{
	3, // 0: com.coralogixapis.alerts.v3.LogsNewValueType.logs_filter:type_name -> com.coralogixapis.alerts.v3.LogsFilter
	1, // 1: com.coralogixapis.alerts.v3.LogsNewValueType.rules:type_name -> com.coralogixapis.alerts.v3.LogsNewValueRule
	4, // 2: com.coralogixapis.alerts.v3.LogsNewValueType.notification_payload_filter:type_name -> google.protobuf.StringValue
	2, // 3: com.coralogixapis.alerts.v3.LogsNewValueRule.condition:type_name -> com.coralogixapis.alerts.v3.LogsNewValueCondition
	4, // 4: com.coralogixapis.alerts.v3.LogsNewValueCondition.keypath_to_track:type_name -> google.protobuf.StringValue
	5, // 5: com.coralogixapis.alerts.v3.LogsNewValueCondition.time_window:type_name -> com.coralogixapis.alerts.v3.LogsNewValueTimeWindow
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_init()
}
func file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_depIdxs = nil
}
