// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/logs/logs_ratio_threshold_type_definition.proto

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

type LogsRatioThresholdType struct {
	state                      protoimpl.MessageState      `protogen:"open.v1"`
	Numerator                  *LogsFilter                 `protobuf:"bytes,1,opt,name=numerator,proto3" json:"numerator,omitempty"`
	NumeratorAlias             *wrapperspb.StringValue     `protobuf:"bytes,2,opt,name=numerator_alias,json=numeratorAlias,proto3" json:"numerator_alias,omitempty"`
	Denominator                *LogsFilter                 `protobuf:"bytes,3,opt,name=denominator,proto3" json:"denominator,omitempty"`
	DenominatorAlias           *wrapperspb.StringValue     `protobuf:"bytes,4,opt,name=denominator_alias,json=denominatorAlias,proto3" json:"denominator_alias,omitempty"`
	Rules                      []*LogsRatioRules           `protobuf:"bytes,5,rep,name=rules,proto3" json:"rules,omitempty"`
	NotificationPayloadFilter  []*wrapperspb.StringValue   `protobuf:"bytes,6,rep,name=notification_payload_filter,json=notificationPayloadFilter,proto3" json:"notification_payload_filter,omitempty"`
	GroupByFor                 LogsRatioGroupByFor         `protobuf:"varint,7,opt,name=group_by_for,json=groupByFor,proto3,enum=com.coralogixapis.alerts.v3.LogsRatioGroupByFor" json:"group_by_for,omitempty"`
	UndetectedValuesManagement *UndetectedValuesManagement `protobuf:"bytes,8,opt,name=undetected_values_management,json=undetectedValuesManagement,proto3" json:"undetected_values_management,omitempty"`
	IgnoreInfinity             *wrapperspb.BoolValue       `protobuf:"bytes,9,opt,name=ignore_infinity,json=ignoreInfinity,proto3" json:"ignore_infinity,omitempty"`
	EvaluationDelayMs          *wrapperspb.Int32Value      `protobuf:"bytes,10,opt,name=evaluation_delay_ms,json=evaluationDelayMs,proto3" json:"evaluation_delay_ms,omitempty"`
	unknownFields              protoimpl.UnknownFields
	sizeCache                  protoimpl.SizeCache
}

func (x *LogsRatioThresholdType) Reset() {
	*x = LogsRatioThresholdType{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsRatioThresholdType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsRatioThresholdType) ProtoMessage() {}

func (x *LogsRatioThresholdType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsRatioThresholdType.ProtoReflect.Descriptor instead.
func (*LogsRatioThresholdType) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_rawDescGZIP(), []int{0}
}

func (x *LogsRatioThresholdType) GetNumerator() *LogsFilter {
	if x != nil {
		return x.Numerator
	}
	return nil
}

func (x *LogsRatioThresholdType) GetNumeratorAlias() *wrapperspb.StringValue {
	if x != nil {
		return x.NumeratorAlias
	}
	return nil
}

func (x *LogsRatioThresholdType) GetDenominator() *LogsFilter {
	if x != nil {
		return x.Denominator
	}
	return nil
}

func (x *LogsRatioThresholdType) GetDenominatorAlias() *wrapperspb.StringValue {
	if x != nil {
		return x.DenominatorAlias
	}
	return nil
}

func (x *LogsRatioThresholdType) GetRules() []*LogsRatioRules {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *LogsRatioThresholdType) GetNotificationPayloadFilter() []*wrapperspb.StringValue {
	if x != nil {
		return x.NotificationPayloadFilter
	}
	return nil
}

func (x *LogsRatioThresholdType) GetGroupByFor() LogsRatioGroupByFor {
	if x != nil {
		return x.GroupByFor
	}
	return LogsRatioGroupByFor_LOGS_RATIO_GROUP_BY_FOR_BOTH_OR_UNSPECIFIED
}

func (x *LogsRatioThresholdType) GetUndetectedValuesManagement() *UndetectedValuesManagement {
	if x != nil {
		return x.UndetectedValuesManagement
	}
	return nil
}

func (x *LogsRatioThresholdType) GetIgnoreInfinity() *wrapperspb.BoolValue {
	if x != nil {
		return x.IgnoreInfinity
	}
	return nil
}

func (x *LogsRatioThresholdType) GetEvaluationDelayMs() *wrapperspb.Int32Value {
	if x != nil {
		return x.EvaluationDelayMs
	}
	return nil
}

type LogsRatioRules struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Condition     *LogsRatioCondition    `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition,omitempty"`
	Override      *AlertDefOverride      `protobuf:"bytes,2,opt,name=override,proto3" json:"override,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogsRatioRules) Reset() {
	*x = LogsRatioRules{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsRatioRules) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsRatioRules) ProtoMessage() {}

func (x *LogsRatioRules) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsRatioRules.ProtoReflect.Descriptor instead.
func (*LogsRatioRules) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_rawDescGZIP(), []int{1}
}

func (x *LogsRatioRules) GetCondition() *LogsRatioCondition {
	if x != nil {
		return x.Condition
	}
	return nil
}

func (x *LogsRatioRules) GetOverride() *AlertDefOverride {
	if x != nil {
		return x.Override
	}
	return nil
}

type LogsRatioCondition struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Threshold     *wrapperspb.DoubleValue `protobuf:"bytes,1,opt,name=threshold,proto3" json:"threshold,omitempty"`
	TimeWindow    *LogsRatioTimeWindow    `protobuf:"bytes,2,opt,name=time_window,json=timeWindow,proto3" json:"time_window,omitempty"`
	ConditionType LogsRatioConditionType  `protobuf:"varint,4,opt,name=condition_type,json=conditionType,proto3,enum=com.coralogixapis.alerts.v3.LogsRatioConditionType" json:"condition_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogsRatioCondition) Reset() {
	*x = LogsRatioCondition{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsRatioCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsRatioCondition) ProtoMessage() {}

func (x *LogsRatioCondition) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsRatioCondition.ProtoReflect.Descriptor instead.
func (*LogsRatioCondition) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_rawDescGZIP(), []int{2}
}

func (x *LogsRatioCondition) GetThreshold() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Threshold
	}
	return nil
}

func (x *LogsRatioCondition) GetTimeWindow() *LogsRatioTimeWindow {
	if x != nil {
		return x.TimeWindow
	}
	return nil
}

func (x *LogsRatioCondition) GetConditionType() LogsRatioConditionType {
	if x != nil {
		return x.ConditionType
	}
	return LogsRatioConditionType_LOGS_RATIO_CONDITION_TYPE_MORE_THAN_OR_UNSPECIFIED
}

var File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto protoreflect.FileDescriptor

const file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_rawDesc = "" +
	"\n" +
	"ecom/coralogixapis/alerts/v3/alert_def_type_definition/logs/logs_ratio_threshold_type_definition.proto\x12\x1bcom.coralogixapis.alerts.v3\x1aScom/coralogixapis/alerts/v3/alert_def_type_definition/logs/common/logs_filter.proto\x1a`com/coralogixapis/alerts/v3/alert_def_type_definition/logs/ratio/logs_ratio_condition_type.proto\x1a^com/coralogixapis/alerts/v3/alert_def_type_definition/logs/ratio/logs_ratio_group_by_for.proto\x1a\\com/coralogixapis/alerts/v3/alert_def_type_definition/logs/ratio/logs_ratio_timewindow.proto\x1aEcom/coralogixapis/alerts/v3/alert_def_type_definition/overrides.proto\x1aXcom/coralogixapis/alerts/v3/alert_def_type_definition/undetected_values_management.proto\x1a\x1egoogle/protobuf/wrappers.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\xc6\x0e\n" +
	"\x16LogsRatioThresholdType\x12\x8a\x01\n" +
	"\tnumerator\x18\x01 \x01(\v2'.com.coralogixapis.alerts.v3.LogsFilterBC\x92A@2>The filter to match log entries for the numerator of the ratioR\tnumerator\x12\x9c\x01\n" +
	"\x0fnumerator_alias\x18\x02 \x01(\v2\x1c.google.protobuf.StringValueBU\x92AR2=The alias for the numerator filter, used for display purposesJ\x11\"numerator_alias\"R\x0enumeratorAlias\x12\x90\x01\n" +
	"\vdenominator\x18\x03 \x01(\v2'.com.coralogixapis.alerts.v3.LogsFilterBE\x92AB2@The filter to match log entries for the denominator of the ratioR\vdenominator\x12\xa4\x01\n" +
	"\x11denominator_alias\x18\x04 \x01(\v2\x1c.google.protobuf.StringValueBY\x92AV2?The alias for the denominator filter, used for display purposesJ\x13\"denominator_alias\"R\x10denominatorAlias\x12k\n" +
	"\x05rules\x18\x05 \x03(\v2+.com.coralogixapis.alerts.v3.LogsRatioRulesB(\x92A%2\x1dThe rules for the ratio alert\xa0\x01\x05\xa8\x01\x01R\x05rules\x12\x83\x01\n" +
	"\x1bnotification_payload_filter\x18\x06 \x03(\v2\x1c.google.protobuf.StringValueB%\x92A\"J\r[\"obj.field\"]\x8a\x01\x10^[a-zA-Z0-9_.]+$R\x19notificationPayloadFilter\x12\xc2\x01\n" +
	"\fgroup_by_for\x18\a \x01(\x0e20.com.coralogixapis.alerts.v3.LogsRatioGroupByForBn\x92Ak2?The group by settings for the numerator and denominator filtersJ(\"LOGS_RATIO_GROUP_BY_FOR_NUMERATOR_ONLY\"R\n" +
	"groupByFor\x12\xbd\x01\n" +
	"\x1cundetected_values_management\x18\b \x01(\v27.com.coralogixapis.alerts.v3.UndetectedValuesManagementBB\x92A?2=Configuration for handling the undetected values in the alertR\x1aundetectedValuesManagement\x12\x8b\x01\n" +
	"\x0fignore_infinity\x18\t \x01(\v2\x1a.google.protobuf.BoolValueBF\x92AC2;The configuration for ignoring infinity values in the ratioJ\x04trueR\x0eignoreInfinity\x12\x98\x01\n" +
	"\x13evaluation_delay_ms\x18\n" +
	" \x01(\v2\x1b.google.protobuf.Int32ValueBK\x92AH2?The delay in milliseconds before evaluating the alert conditionJ\x0560000R\x11evaluationDelayMs:\xa4\x02\x92A\xa0\x02\n" +
	"\x8d\x01*$Log-based ratio threshold alert type2CConfiguration for alerts based on the ratio between two log queries\xd2\x01\tnumerator\xd2\x01\vdenominator\xd2\x01\x05rules*\x8d\x01\n" +
	"7Learn more about logs ratio alerts in our documentation\x12Rhttps://coralogix.com/docs/user-guides/alerting/create-an-alert/logs/ratio-alerts/\"\xdf\x02\n" +
	"\x0eLogsRatioRules\x12u\n" +
	"\tcondition\x18\x01 \x01(\v2/.com.coralogixapis.alerts.v3.LogsRatioConditionB&\x92A#2!The condition for the ratio alertR\tcondition\x12s\n" +
	"\boverride\x18\x02 \x01(\v2-.com.coralogixapis.alerts.v3.AlertDefOverrideB(\x92A%2#The override settings for the alertR\boverride:a\x92A^\n" +
	"\\*\x15Log-based ratio rules2,Defines the rules for log-based ratio alerts\xd2\x01\tcondition\xd2\x01\boverride\"\xd4\x04\n" +
	"\x12LogsRatioCondition\x12\x82\x01\n" +
	"\tthreshold\x18\x01 \x01(\v2\x1c.google.protobuf.DoubleValueBF\x92AC2+The threshold value for the alert conditionJ\x0410.0\x8a\x01\r^\\d+(\\.\\d+)?$R\tthreshold\x12\x83\x01\n" +
	"\vtime_window\x18\x02 \x01(\v20.com.coralogixapis.alerts.v3.LogsRatioTimeWindowB0\x92A-2'The time window for the alert conditionJ\x0210R\n" +
	"timeWindow\x12\xba\x01\n" +
	"\x0econdition_type\x18\x04 \x01(\x0e23.com.coralogixapis.alerts.v3.LogsRatioConditionTypeB^\x92A[2#The type of condition for the alertJ4\"LOGS_RATIO_CONDITION_TYPE_MORE_THAN_OR_UNSPECIFIED\"R\rconditionType:v\x92As\n" +
	"q*\x19Log-based ratio condition2)Defines conditions for ratio-based alerts\xd2\x01\tthreshold\xd2\x01\vtime_window\xd2\x01\x0econdition_typeb\x06proto3"

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_goTypes = []any{
	(*LogsRatioThresholdType)(nil),     // 0: com.coralogixapis.alerts.v3.LogsRatioThresholdType
	(*LogsRatioRules)(nil),             // 1: com.coralogixapis.alerts.v3.LogsRatioRules
	(*LogsRatioCondition)(nil),         // 2: com.coralogixapis.alerts.v3.LogsRatioCondition
	(*LogsFilter)(nil),                 // 3: com.coralogixapis.alerts.v3.LogsFilter
	(*wrapperspb.StringValue)(nil),     // 4: google.protobuf.StringValue
	(LogsRatioGroupByFor)(0),           // 5: com.coralogixapis.alerts.v3.LogsRatioGroupByFor
	(*UndetectedValuesManagement)(nil), // 6: com.coralogixapis.alerts.v3.UndetectedValuesManagement
	(*wrapperspb.BoolValue)(nil),       // 7: google.protobuf.BoolValue
	(*wrapperspb.Int32Value)(nil),      // 8: google.protobuf.Int32Value
	(*AlertDefOverride)(nil),           // 9: com.coralogixapis.alerts.v3.AlertDefOverride
	(*wrapperspb.DoubleValue)(nil),     // 10: google.protobuf.DoubleValue
	(*LogsRatioTimeWindow)(nil),        // 11: com.coralogixapis.alerts.v3.LogsRatioTimeWindow
	(LogsRatioConditionType)(0),        // 12: com.coralogixapis.alerts.v3.LogsRatioConditionType
}
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_depIdxs = []int32{
	3,  // 0: com.coralogixapis.alerts.v3.LogsRatioThresholdType.numerator:type_name -> com.coralogixapis.alerts.v3.LogsFilter
	4,  // 1: com.coralogixapis.alerts.v3.LogsRatioThresholdType.numerator_alias:type_name -> google.protobuf.StringValue
	3,  // 2: com.coralogixapis.alerts.v3.LogsRatioThresholdType.denominator:type_name -> com.coralogixapis.alerts.v3.LogsFilter
	4,  // 3: com.coralogixapis.alerts.v3.LogsRatioThresholdType.denominator_alias:type_name -> google.protobuf.StringValue
	1,  // 4: com.coralogixapis.alerts.v3.LogsRatioThresholdType.rules:type_name -> com.coralogixapis.alerts.v3.LogsRatioRules
	4,  // 5: com.coralogixapis.alerts.v3.LogsRatioThresholdType.notification_payload_filter:type_name -> google.protobuf.StringValue
	5,  // 6: com.coralogixapis.alerts.v3.LogsRatioThresholdType.group_by_for:type_name -> com.coralogixapis.alerts.v3.LogsRatioGroupByFor
	6,  // 7: com.coralogixapis.alerts.v3.LogsRatioThresholdType.undetected_values_management:type_name -> com.coralogixapis.alerts.v3.UndetectedValuesManagement
	7,  // 8: com.coralogixapis.alerts.v3.LogsRatioThresholdType.ignore_infinity:type_name -> google.protobuf.BoolValue
	8,  // 9: com.coralogixapis.alerts.v3.LogsRatioThresholdType.evaluation_delay_ms:type_name -> google.protobuf.Int32Value
	2,  // 10: com.coralogixapis.alerts.v3.LogsRatioRules.condition:type_name -> com.coralogixapis.alerts.v3.LogsRatioCondition
	9,  // 11: com.coralogixapis.alerts.v3.LogsRatioRules.override:type_name -> com.coralogixapis.alerts.v3.AlertDefOverride
	10, // 12: com.coralogixapis.alerts.v3.LogsRatioCondition.threshold:type_name -> google.protobuf.DoubleValue
	11, // 13: com.coralogixapis.alerts.v3.LogsRatioCondition.time_window:type_name -> com.coralogixapis.alerts.v3.LogsRatioTimeWindow
	12, // 14: com.coralogixapis.alerts.v3.LogsRatioCondition.condition_type:type_name -> com.coralogixapis.alerts.v3.LogsRatioConditionType
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_init()
}
func file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_condition_type_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_group_by_for_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_undetected_values_management_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_depIdxs = nil
}
