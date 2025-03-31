// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/logs/logs_unique_count_type_definition.proto

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

type LogsUniqueCountType struct {
	state                       protoimpl.MessageState    `protogen:"open.v1"`
	LogsFilter                  *LogsFilter               `protobuf:"bytes,1,opt,name=logs_filter,json=logsFilter,proto3" json:"logs_filter,omitempty"`
	Rules                       []*LogsUniqueCountRule    `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty"`
	NotificationPayloadFilter   []*wrapperspb.StringValue `protobuf:"bytes,3,rep,name=notification_payload_filter,json=notificationPayloadFilter,proto3" json:"notification_payload_filter,omitempty"`
	MaxUniqueCountPerGroupByKey *wrapperspb.Int64Value    `protobuf:"bytes,4,opt,name=max_unique_count_per_group_by_key,json=maxUniqueCountPerGroupByKey,proto3" json:"max_unique_count_per_group_by_key,omitempty"`
	UniqueCountKeypath          *wrapperspb.StringValue   `protobuf:"bytes,5,opt,name=unique_count_keypath,json=uniqueCountKeypath,proto3" json:"unique_count_keypath,omitempty"`
	unknownFields               protoimpl.UnknownFields
	sizeCache                   protoimpl.SizeCache
}

func (x *LogsUniqueCountType) Reset() {
	*x = LogsUniqueCountType{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsUniqueCountType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsUniqueCountType) ProtoMessage() {}

func (x *LogsUniqueCountType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsUniqueCountType.ProtoReflect.Descriptor instead.
func (*LogsUniqueCountType) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_rawDescGZIP(), []int{0}
}

func (x *LogsUniqueCountType) GetLogsFilter() *LogsFilter {
	if x != nil {
		return x.LogsFilter
	}
	return nil
}

func (x *LogsUniqueCountType) GetRules() []*LogsUniqueCountRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *LogsUniqueCountType) GetNotificationPayloadFilter() []*wrapperspb.StringValue {
	if x != nil {
		return x.NotificationPayloadFilter
	}
	return nil
}

func (x *LogsUniqueCountType) GetMaxUniqueCountPerGroupByKey() *wrapperspb.Int64Value {
	if x != nil {
		return x.MaxUniqueCountPerGroupByKey
	}
	return nil
}

func (x *LogsUniqueCountType) GetUniqueCountKeypath() *wrapperspb.StringValue {
	if x != nil {
		return x.UniqueCountKeypath
	}
	return nil
}

type LogsUniqueCountRule struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Condition     *LogsUniqueCountCondition `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogsUniqueCountRule) Reset() {
	*x = LogsUniqueCountRule{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsUniqueCountRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsUniqueCountRule) ProtoMessage() {}

func (x *LogsUniqueCountRule) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsUniqueCountRule.ProtoReflect.Descriptor instead.
func (*LogsUniqueCountRule) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_rawDescGZIP(), []int{1}
}

func (x *LogsUniqueCountRule) GetCondition() *LogsUniqueCountCondition {
	if x != nil {
		return x.Condition
	}
	return nil
}

type LogsUniqueCountCondition struct {
	state          protoimpl.MessageState     `protogen:"open.v1"`
	MaxUniqueCount *wrapperspb.Int64Value     `protobuf:"bytes,2,opt,name=max_unique_count,json=maxUniqueCount,proto3" json:"max_unique_count,omitempty"`
	TimeWindow     *LogsUniqueValueTimeWindow `protobuf:"bytes,3,opt,name=time_window,json=timeWindow,proto3" json:"time_window,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *LogsUniqueCountCondition) Reset() {
	*x = LogsUniqueCountCondition{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsUniqueCountCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsUniqueCountCondition) ProtoMessage() {}

func (x *LogsUniqueCountCondition) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsUniqueCountCondition.ProtoReflect.Descriptor instead.
func (*LogsUniqueCountCondition) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_rawDescGZIP(), []int{2}
}

func (x *LogsUniqueCountCondition) GetMaxUniqueCount() *wrapperspb.Int64Value {
	if x != nil {
		return x.MaxUniqueCount
	}
	return nil
}

func (x *LogsUniqueCountCondition) GetTimeWindow() *LogsUniqueValueTimeWindow {
	if x != nil {
		return x.TimeWindow
	}
	return nil
}

var File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto protoreflect.FileDescriptor

const file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_rawDesc = "" +
	"\n" +
	"bcom/coralogixapis/alerts/v3/alert_def_type_definition/logs/logs_unique_count_type_definition.proto\x12\x1bcom.coralogixapis.alerts.v3\x1a\x1egoogle/protobuf/wrappers.proto\x1aScom/coralogixapis/alerts/v3/alert_def_type_definition/logs/common/logs_filter.proto\x1ajcom/coralogixapis/alerts/v3/alert_def_type_definition/logs/unique_count/logs_unique_value_timewindow.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\xfa\x04\n" +
	"\x13LogsUniqueCountType\x12H\n" +
	"\vlogs_filter\x18\x01 \x01(\v2'.com.coralogixapis.alerts.v3.LogsFilterR\n" +
	"logsFilter\x12F\n" +
	"\x05rules\x18\x02 \x03(\v20.com.coralogixapis.alerts.v3.LogsUniqueCountRuleR\x05rules\x12i\n" +
	"\x1bnotification_payload_filter\x18\x03 \x03(\v2\x1c.google.protobuf.StringValueB\v\x92A\bJ\x06\"text\"R\x19notificationPayloadFilter\x12m\n" +
	"!max_unique_count_per_group_by_key\x18\x04 \x01(\v2\x1b.google.protobuf.Int64ValueB\b\x92A\x05J\x03100R\x1bmaxUniqueCountPerGroupByKey\x12c\n" +
	"\x14unique_count_keypath\x18\x05 \x01(\v2\x1c.google.protobuf.StringValueB\x13\x92A\x10J\x0e\"unique_count\"R\x12uniqueCountKeypath:\x91\x01\x92A\x8d\x01\n" +
	"\x8a\x01*\x1cLogs Unique Count Alert Type2=Configuration for alerts based on unique value counts in logs\xd2\x01\vlogs_filter\xd2\x01\x05rules\xd2\x01\x14unique_count_keypath\"j\n" +
	"\x13LogsUniqueCountRule\x12S\n" +
	"\tcondition\x18\x01 \x01(\v25.com.coralogixapis.alerts.v3.LogsUniqueCountConditionR\tcondition\"\xab\x02\n" +
	"\x18LogsUniqueCountCondition\x12E\n" +
	"\x10max_unique_count\x18\x02 \x01(\v2\x1b.google.protobuf.Int64ValueR\x0emaxUniqueCount\x12W\n" +
	"\vtime_window\x18\x03 \x01(\v26.com.coralogixapis.alerts.v3.LogsUniqueValueTimeWindowR\n" +
	"timeWindow:o\x92Al\n" +
	"j*\x1bLogs Unique Count Condition2*Defines conditions for unique count alerts\xd2\x01\x10max_unique_count\xd2\x01\vtime_windowb\x06proto3"

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_goTypes = []any{
	(*LogsUniqueCountType)(nil),       // 0: com.coralogixapis.alerts.v3.LogsUniqueCountType
	(*LogsUniqueCountRule)(nil),       // 1: com.coralogixapis.alerts.v3.LogsUniqueCountRule
	(*LogsUniqueCountCondition)(nil),  // 2: com.coralogixapis.alerts.v3.LogsUniqueCountCondition
	(*LogsFilter)(nil),                // 3: com.coralogixapis.alerts.v3.LogsFilter
	(*wrapperspb.StringValue)(nil),    // 4: google.protobuf.StringValue
	(*wrapperspb.Int64Value)(nil),     // 5: google.protobuf.Int64Value
	(*LogsUniqueValueTimeWindow)(nil), // 6: com.coralogixapis.alerts.v3.LogsUniqueValueTimeWindow
}
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_depIdxs = []int32{
	3, // 0: com.coralogixapis.alerts.v3.LogsUniqueCountType.logs_filter:type_name -> com.coralogixapis.alerts.v3.LogsFilter
	1, // 1: com.coralogixapis.alerts.v3.LogsUniqueCountType.rules:type_name -> com.coralogixapis.alerts.v3.LogsUniqueCountRule
	4, // 2: com.coralogixapis.alerts.v3.LogsUniqueCountType.notification_payload_filter:type_name -> google.protobuf.StringValue
	5, // 3: com.coralogixapis.alerts.v3.LogsUniqueCountType.max_unique_count_per_group_by_key:type_name -> google.protobuf.Int64Value
	4, // 4: com.coralogixapis.alerts.v3.LogsUniqueCountType.unique_count_keypath:type_name -> google.protobuf.StringValue
	2, // 5: com.coralogixapis.alerts.v3.LogsUniqueCountRule.condition:type_name -> com.coralogixapis.alerts.v3.LogsUniqueCountCondition
	5, // 6: com.coralogixapis.alerts.v3.LogsUniqueCountCondition.max_unique_count:type_name -> google.protobuf.Int64Value
	6, // 7: com.coralogixapis.alerts.v3.LogsUniqueCountCondition.time_window:type_name -> com.coralogixapis.alerts.v3.LogsUniqueValueTimeWindow
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_init()
}
func file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_unique_count_logs_unique_value_timewindow_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_depIdxs = nil
}
