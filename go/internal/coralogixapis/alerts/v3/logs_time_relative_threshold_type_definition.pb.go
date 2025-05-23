// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/logs/logs_time_relative_threshold_type_definition.proto

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

type LogsTimeRelativeThresholdType struct {
	state                      protoimpl.MessageState      `protogen:"open.v1"`
	LogsFilter                 *LogsFilter                 `protobuf:"bytes,1,opt,name=logs_filter,json=logsFilter,proto3" json:"logs_filter,omitempty"`
	Rules                      []*LogsTimeRelativeRule     `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty"`
	IgnoreInfinity             *wrapperspb.BoolValue       `protobuf:"bytes,3,opt,name=ignore_infinity,json=ignoreInfinity,proto3" json:"ignore_infinity,omitempty"`
	NotificationPayloadFilter  []*wrapperspb.StringValue   `protobuf:"bytes,4,rep,name=notification_payload_filter,json=notificationPayloadFilter,proto3" json:"notification_payload_filter,omitempty"`
	UndetectedValuesManagement *UndetectedValuesManagement `protobuf:"bytes,5,opt,name=undetected_values_management,json=undetectedValuesManagement,proto3" json:"undetected_values_management,omitempty"`
	EvaluationDelayMs          *wrapperspb.Int32Value      `protobuf:"bytes,6,opt,name=evaluation_delay_ms,json=evaluationDelayMs,proto3" json:"evaluation_delay_ms,omitempty"`
	unknownFields              protoimpl.UnknownFields
	sizeCache                  protoimpl.SizeCache
}

func (x *LogsTimeRelativeThresholdType) Reset() {
	*x = LogsTimeRelativeThresholdType{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsTimeRelativeThresholdType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsTimeRelativeThresholdType) ProtoMessage() {}

func (x *LogsTimeRelativeThresholdType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsTimeRelativeThresholdType.ProtoReflect.Descriptor instead.
func (*LogsTimeRelativeThresholdType) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_rawDescGZIP(), []int{0}
}

func (x *LogsTimeRelativeThresholdType) GetLogsFilter() *LogsFilter {
	if x != nil {
		return x.LogsFilter
	}
	return nil
}

func (x *LogsTimeRelativeThresholdType) GetRules() []*LogsTimeRelativeRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *LogsTimeRelativeThresholdType) GetIgnoreInfinity() *wrapperspb.BoolValue {
	if x != nil {
		return x.IgnoreInfinity
	}
	return nil
}

func (x *LogsTimeRelativeThresholdType) GetNotificationPayloadFilter() []*wrapperspb.StringValue {
	if x != nil {
		return x.NotificationPayloadFilter
	}
	return nil
}

func (x *LogsTimeRelativeThresholdType) GetUndetectedValuesManagement() *UndetectedValuesManagement {
	if x != nil {
		return x.UndetectedValuesManagement
	}
	return nil
}

func (x *LogsTimeRelativeThresholdType) GetEvaluationDelayMs() *wrapperspb.Int32Value {
	if x != nil {
		return x.EvaluationDelayMs
	}
	return nil
}

type LogsTimeRelativeRule struct {
	state         protoimpl.MessageState     `protogen:"open.v1"`
	Condition     *LogsTimeRelativeCondition `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition,omitempty"`
	Override      *AlertDefOverride          `protobuf:"bytes,2,opt,name=override,proto3" json:"override,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogsTimeRelativeRule) Reset() {
	*x = LogsTimeRelativeRule{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsTimeRelativeRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsTimeRelativeRule) ProtoMessage() {}

func (x *LogsTimeRelativeRule) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsTimeRelativeRule.ProtoReflect.Descriptor instead.
func (*LogsTimeRelativeRule) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_rawDescGZIP(), []int{1}
}

func (x *LogsTimeRelativeRule) GetCondition() *LogsTimeRelativeCondition {
	if x != nil {
		return x.Condition
	}
	return nil
}

func (x *LogsTimeRelativeRule) GetOverride() *AlertDefOverride {
	if x != nil {
		return x.Override
	}
	return nil
}

type LogsTimeRelativeCondition struct {
	state         protoimpl.MessageState        `protogen:"open.v1"`
	Threshold     *wrapperspb.DoubleValue       `protobuf:"bytes,1,opt,name=threshold,proto3" json:"threshold,omitempty"`
	ComparedTo    LogsTimeRelativeComparedTo    `protobuf:"varint,2,opt,name=compared_to,json=comparedTo,proto3,enum=com.coralogixapis.alerts.v3.LogsTimeRelativeComparedTo" json:"compared_to,omitempty"`
	ConditionType LogsTimeRelativeConditionType `protobuf:"varint,4,opt,name=condition_type,json=conditionType,proto3,enum=com.coralogixapis.alerts.v3.LogsTimeRelativeConditionType" json:"condition_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogsTimeRelativeCondition) Reset() {
	*x = LogsTimeRelativeCondition{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsTimeRelativeCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsTimeRelativeCondition) ProtoMessage() {}

func (x *LogsTimeRelativeCondition) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsTimeRelativeCondition.ProtoReflect.Descriptor instead.
func (*LogsTimeRelativeCondition) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_rawDescGZIP(), []int{2}
}

func (x *LogsTimeRelativeCondition) GetThreshold() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Threshold
	}
	return nil
}

func (x *LogsTimeRelativeCondition) GetComparedTo() LogsTimeRelativeComparedTo {
	if x != nil {
		return x.ComparedTo
	}
	return LogsTimeRelativeComparedTo_LOGS_TIME_RELATIVE_COMPARED_TO_PREVIOUS_HOUR_OR_UNSPECIFIED
}

func (x *LogsTimeRelativeCondition) GetConditionType() LogsTimeRelativeConditionType {
	if x != nil {
		return x.ConditionType
	}
	return LogsTimeRelativeConditionType_LOGS_TIME_RELATIVE_CONDITION_TYPE_MORE_THAN_OR_UNSPECIFIED
}

var File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto protoreflect.FileDescriptor

const file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_rawDesc = "" +
	"\n" +
	"mcom/coralogixapis/alerts/v3/alert_def_type_definition/logs/logs_time_relative_threshold_type_definition.proto\x12\x1bcom.coralogixapis.alerts.v3\x1aScom/coralogixapis/alerts/v3/alert_def_type_definition/logs/common/logs_filter.proto\x1amcom/coralogixapis/alerts/v3/alert_def_type_definition/logs/time_relative/logs_time_relative_compared_to.proto\x1apcom/coralogixapis/alerts/v3/alert_def_type_definition/logs/time_relative/logs_time_relative_condition_type.proto\x1aEcom/coralogixapis/alerts/v3/alert_def_type_definition/overrides.proto\x1aXcom/coralogixapis/alerts/v3/alert_def_type_definition/undetected_values_management.proto\x1a\x1egoogle/protobuf/wrappers.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\xcd\t\n" +
	"\x1dLogsTimeRelativeThresholdType\x12}\n" +
	"\vlogs_filter\x18\x01 \x01(\v2'.com.coralogixapis.alerts.v3.LogsFilterB3\x92A02.The filter to match log entries for the alert.R\n" +
	"logsFilter\x12y\n" +
	"\x05rules\x18\x02 \x03(\v21.com.coralogixapis.alerts.v3.LogsTimeRelativeRuleB0\x92A-2%The rules for the time-relative alert\xa0\x01\x05\xa8\x01\x01R\x05rules\x12s\n" +
	"\x0fignore_infinity\x18\x03 \x01(\v2\x1a.google.protobuf.BoolValueB.\x92A+2#Ignore infinity values in the alertJ\x04trueR\x0eignoreInfinity\x12\x83\x01\n" +
	"\x1bnotification_payload_filter\x18\x04 \x03(\v2\x1c.google.protobuf.StringValueB%\x92A\"J\r[\"obj.field\"]\x8a\x01\x10^[a-zA-Z0-9_.]+$R\x19notificationPayloadFilter\x12\xbe\x01\n" +
	"\x1cundetected_values_management\x18\x05 \x01(\v27.com.coralogixapis.alerts.v3.UndetectedValuesManagementBC\x92A@2>Configuration for handling the undetected values in the alert.R\x1aundetectedValuesManagement\x12\x98\x01\n" +
	"\x13evaluation_delay_ms\x18\x06 \x01(\v2\x1b.google.protobuf.Int32ValueBK\x92AH2?The delay in milliseconds before evaluating the alert conditionJ\x0560000R\x11evaluationDelayMs:\xda\x02\x92A\xd6\x02\n" +
	"\xad\x01*,Log-based time-relative threshold alert type2uConfiguration for alerts that are triggered when a fixed ratio reaches a set threshold compared to a past time frame.\xd2\x01\x05rules*\xa3\x01\n" +
	"ELearn more about log-based, time-relative alerts in our documentation\x12Zhttps://coralogix.com/docs/user-guides/alerting/create-an-alert/logs/time-relative-alerts/\"\xa8\x03\n" +
	"\x14LogsTimeRelativeRule\x12\x85\x01\n" +
	"\tcondition\x18\x01 \x01(\v26.com.coralogixapis.alerts.v3.LogsTimeRelativeConditionB/\x92A,2*The condition for the time-relative alert.R\tcondition\x12t\n" +
	"\boverride\x18\x02 \x01(\v2-.com.coralogixapis.alerts.v3.AlertDefOverrideB)\x92A&2$The override settings for the alert.R\boverride:\x91\x01\x92A\x8d\x01\n" +
	"\x8a\x01*\x17Logs Time Relative Rule2XLogsTimeRelativeRule is a message that defines a rule for log-based time-relative alerts\xd2\x01\tcondition\xd2\x01\boverride\"\xa7\x05\n" +
	"\x19LogsTimeRelativeCondition\x12m\n" +
	"\tthreshold\x18\x01 \x01(\v2\x1c.google.protobuf.DoubleValueB1\x92A.2,The threshold value for the alert condition.R\tthreshold\x12\xbe\x01\n" +
	"\vcompared_to\x18\x02 \x01(\x0e27.com.coralogixapis.alerts.v3.LogsTimeRelativeComparedToBd\x92Aa23The time frame to compare the current value againstJ*\"LOGS_TIME_RELATIVE_COMPARED_TO_YESTERDAY\"R\n" +
	"comparedTo\x12\xc9\x01\n" +
	"\x0econdition_type\x18\x04 \x01(\x0e2:.com.coralogixapis.alerts.v3.LogsTimeRelativeConditionTypeBf\x92Ac2#The type of condition for the alertJ<\"LOGS_TIME_RELATIVE_CONDITION_TYPE_MORE_THAN_OR_UNSPECIFIED\"R\rconditionType:\x8d\x01\x92A\x89\x01\n" +
	"\x86\x01*!Log-based time-relative condition26Defines conditions for time-relative comparison alerts\xd2\x01\tthreshold\xd2\x01\vcompared_to\xd2\x01\x0econdition_typeb\x06proto3"

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_goTypes = []any{
	(*LogsTimeRelativeThresholdType)(nil), // 0: com.coralogixapis.alerts.v3.LogsTimeRelativeThresholdType
	(*LogsTimeRelativeRule)(nil),          // 1: com.coralogixapis.alerts.v3.LogsTimeRelativeRule
	(*LogsTimeRelativeCondition)(nil),     // 2: com.coralogixapis.alerts.v3.LogsTimeRelativeCondition
	(*LogsFilter)(nil),                    // 3: com.coralogixapis.alerts.v3.LogsFilter
	(*wrapperspb.BoolValue)(nil),          // 4: google.protobuf.BoolValue
	(*wrapperspb.StringValue)(nil),        // 5: google.protobuf.StringValue
	(*UndetectedValuesManagement)(nil),    // 6: com.coralogixapis.alerts.v3.UndetectedValuesManagement
	(*wrapperspb.Int32Value)(nil),         // 7: google.protobuf.Int32Value
	(*AlertDefOverride)(nil),              // 8: com.coralogixapis.alerts.v3.AlertDefOverride
	(*wrapperspb.DoubleValue)(nil),        // 9: google.protobuf.DoubleValue
	(LogsTimeRelativeComparedTo)(0),       // 10: com.coralogixapis.alerts.v3.LogsTimeRelativeComparedTo
	(LogsTimeRelativeConditionType)(0),    // 11: com.coralogixapis.alerts.v3.LogsTimeRelativeConditionType
}
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_depIdxs = []int32{
	3,  // 0: com.coralogixapis.alerts.v3.LogsTimeRelativeThresholdType.logs_filter:type_name -> com.coralogixapis.alerts.v3.LogsFilter
	1,  // 1: com.coralogixapis.alerts.v3.LogsTimeRelativeThresholdType.rules:type_name -> com.coralogixapis.alerts.v3.LogsTimeRelativeRule
	4,  // 2: com.coralogixapis.alerts.v3.LogsTimeRelativeThresholdType.ignore_infinity:type_name -> google.protobuf.BoolValue
	5,  // 3: com.coralogixapis.alerts.v3.LogsTimeRelativeThresholdType.notification_payload_filter:type_name -> google.protobuf.StringValue
	6,  // 4: com.coralogixapis.alerts.v3.LogsTimeRelativeThresholdType.undetected_values_management:type_name -> com.coralogixapis.alerts.v3.UndetectedValuesManagement
	7,  // 5: com.coralogixapis.alerts.v3.LogsTimeRelativeThresholdType.evaluation_delay_ms:type_name -> google.protobuf.Int32Value
	2,  // 6: com.coralogixapis.alerts.v3.LogsTimeRelativeRule.condition:type_name -> com.coralogixapis.alerts.v3.LogsTimeRelativeCondition
	8,  // 7: com.coralogixapis.alerts.v3.LogsTimeRelativeRule.override:type_name -> com.coralogixapis.alerts.v3.AlertDefOverride
	9,  // 8: com.coralogixapis.alerts.v3.LogsTimeRelativeCondition.threshold:type_name -> google.protobuf.DoubleValue
	10, // 9: com.coralogixapis.alerts.v3.LogsTimeRelativeCondition.compared_to:type_name -> com.coralogixapis.alerts.v3.LogsTimeRelativeComparedTo
	11, // 10: com.coralogixapis.alerts.v3.LogsTimeRelativeCondition.condition_type:type_name -> com.coralogixapis.alerts.v3.LogsTimeRelativeConditionType
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_init()
}
func file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_time_relative_logs_time_relative_compared_to_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_time_relative_logs_time_relative_condition_type_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_undetected_values_management_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_depIdxs = nil
}
