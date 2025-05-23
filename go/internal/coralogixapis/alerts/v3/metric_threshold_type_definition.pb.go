// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/metric/metric_threshold_type_definition.proto

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

type MetricThresholdType struct {
	state                      protoimpl.MessageState      `protogen:"open.v1"`
	MetricFilter               *MetricFilter               `protobuf:"bytes,1,opt,name=metric_filter,json=metricFilter,proto3" json:"metric_filter,omitempty"`
	Rules                      []*MetricThresholdRule      `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty"`
	UndetectedValuesManagement *UndetectedValuesManagement `protobuf:"bytes,3,opt,name=undetected_values_management,json=undetectedValuesManagement,proto3" json:"undetected_values_management,omitempty"`
	MissingValues              *MetricMissingValues        `protobuf:"bytes,4,opt,name=missing_values,json=missingValues,proto3" json:"missing_values,omitempty"`
	EvaluationDelayMs          *wrapperspb.Int32Value      `protobuf:"bytes,5,opt,name=evaluation_delay_ms,json=evaluationDelayMs,proto3" json:"evaluation_delay_ms,omitempty"`
	unknownFields              protoimpl.UnknownFields
	sizeCache                  protoimpl.SizeCache
}

func (x *MetricThresholdType) Reset() {
	*x = MetricThresholdType{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetricThresholdType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricThresholdType) ProtoMessage() {}

func (x *MetricThresholdType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricThresholdType.ProtoReflect.Descriptor instead.
func (*MetricThresholdType) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_rawDescGZIP(), []int{0}
}

func (x *MetricThresholdType) GetMetricFilter() *MetricFilter {
	if x != nil {
		return x.MetricFilter
	}
	return nil
}

func (x *MetricThresholdType) GetRules() []*MetricThresholdRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *MetricThresholdType) GetUndetectedValuesManagement() *UndetectedValuesManagement {
	if x != nil {
		return x.UndetectedValuesManagement
	}
	return nil
}

func (x *MetricThresholdType) GetMissingValues() *MetricMissingValues {
	if x != nil {
		return x.MissingValues
	}
	return nil
}

func (x *MetricThresholdType) GetEvaluationDelayMs() *wrapperspb.Int32Value {
	if x != nil {
		return x.EvaluationDelayMs
	}
	return nil
}

type MetricThresholdRule struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Condition     *MetricThresholdCondition `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition,omitempty"`
	Override      *AlertDefOverride         `protobuf:"bytes,2,opt,name=override,proto3" json:"override,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MetricThresholdRule) Reset() {
	*x = MetricThresholdRule{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetricThresholdRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricThresholdRule) ProtoMessage() {}

func (x *MetricThresholdRule) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricThresholdRule.ProtoReflect.Descriptor instead.
func (*MetricThresholdRule) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_rawDescGZIP(), []int{1}
}

func (x *MetricThresholdRule) GetCondition() *MetricThresholdCondition {
	if x != nil {
		return x.Condition
	}
	return nil
}

func (x *MetricThresholdRule) GetOverride() *AlertDefOverride {
	if x != nil {
		return x.Override
	}
	return nil
}

type MetricThresholdCondition struct {
	state         protoimpl.MessageState       `protogen:"open.v1"`
	Threshold     *wrapperspb.DoubleValue      `protobuf:"bytes,1,opt,name=threshold,proto3" json:"threshold,omitempty"`
	ForOverPct    *wrapperspb.UInt32Value      `protobuf:"bytes,2,opt,name=for_over_pct,json=forOverPct,proto3" json:"for_over_pct,omitempty"`
	OfTheLast     *MetricTimeWindow            `protobuf:"bytes,3,opt,name=of_the_last,json=ofTheLast,proto3" json:"of_the_last,omitempty"`
	ConditionType MetricThresholdConditionType `protobuf:"varint,5,opt,name=condition_type,json=conditionType,proto3,enum=com.coralogixapis.alerts.v3.MetricThresholdConditionType" json:"condition_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MetricThresholdCondition) Reset() {
	*x = MetricThresholdCondition{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetricThresholdCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricThresholdCondition) ProtoMessage() {}

func (x *MetricThresholdCondition) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricThresholdCondition.ProtoReflect.Descriptor instead.
func (*MetricThresholdCondition) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_rawDescGZIP(), []int{2}
}

func (x *MetricThresholdCondition) GetThreshold() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Threshold
	}
	return nil
}

func (x *MetricThresholdCondition) GetForOverPct() *wrapperspb.UInt32Value {
	if x != nil {
		return x.ForOverPct
	}
	return nil
}

func (x *MetricThresholdCondition) GetOfTheLast() *MetricTimeWindow {
	if x != nil {
		return x.OfTheLast
	}
	return nil
}

func (x *MetricThresholdCondition) GetConditionType() MetricThresholdConditionType {
	if x != nil {
		return x.ConditionType
	}
	return MetricThresholdConditionType_METRIC_THRESHOLD_CONDITION_TYPE_MORE_THAN_OR_UNSPECIFIED
}

var File_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto protoreflect.FileDescriptor

const file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_rawDesc = "" +
	"\n" +
	"ccom/coralogixapis/alerts/v3/alert_def_type_definition/metric/metric_threshold_type_definition.proto\x12\x1bcom.coralogixapis.alerts.v3\x1a[com/coralogixapis/alerts/v3/alert_def_type_definition/metric/common/metric_timewindow.proto\x1aXcom/coralogixapis/alerts/v3/alert_def_type_definition/metric/common/metrics_filter.proto\x1alcom/coralogixapis/alerts/v3/alert_def_type_definition/metric/threshold/metric_threshold_condition_type.proto\x1alcom/coralogixapis/alerts/v3/alert_def_type_definition/metric/threshold/metric_threshold_missing_values.proto\x1aEcom/coralogixapis/alerts/v3/alert_def_type_definition/overrides.proto\x1aXcom/coralogixapis/alerts/v3/alert_def_type_definition/undetected_values_management.proto\x1a\x1egoogle/protobuf/wrappers.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\xc3\b\n" +
	"\x13MetricThresholdType\x12\x86\x01\n" +
	"\rmetric_filter\x18\x01 \x01(\v2).com.coralogixapis.alerts.v3.MetricFilterB6\x92A321The filter to match metric entries for the alert.R\fmetricFilter\x12{\n" +
	"\x05rules\x18\x02 \x03(\v20.com.coralogixapis.alerts.v3.MetricThresholdRuleB3\x92A02(The rules for the metric threshold alert\xa0\x01\x05\xa8\x01\x01R\x05rules\x12\xba\x01\n" +
	"\x1cundetected_values_management\x18\x03 \x01(\v27.com.coralogixapis.alerts.v3.UndetectedValuesManagementB?\x92A<2:Configuration for handling undetected values in the alert.R\x1aundetectedValuesManagement\x12\x95\x01\n" +
	"\x0emissing_values\x18\x04 \x01(\v20.com.coralogixapis.alerts.v3.MetricMissingValuesB<\x92A927Configuration for handling missing values in the alert.R\rmissingValues\x12\x98\x01\n" +
	"\x13evaluation_delay_ms\x18\x05 \x01(\v2\x1b.google.protobuf.Int32ValueBK\x92AH2?The delay in milliseconds before evaluating the alert conditionJ\x0560000R\x11evaluationDelayMs:\xb5\x02\x92A\xb1\x02\n" +
	"\x8b\x01*!Metric-based threshold alert type2=Configuration for alerts based on metric threshold violations\xd2\x01\rmetric_filter\xd2\x01\x05rules\xd2\x01\x0emissing_values*\xa0\x01\n" +
	"CLearn more about metric-based threshold alerts in our documentation\x12Yhttps://coralogix.com/docs/user-guides/alerting/create-an-alert/metrics/threshold-alerts/\"\xee\x02\n" +
	"\x13MetricThresholdRule\x12\x86\x01\n" +
	"\tcondition\x18\x01 \x01(\v25.com.coralogixapis.alerts.v3.MetricThresholdConditionB1\x92A.2,The condition for the metric threshold alertR\tcondition\x12g\n" +
	"\boverride\x18\x02 \x01(\v2-.com.coralogixapis.alerts.v3.AlertDefOverrideB\x1c\x92A\x192\x17Alert override settingsR\boverride:e\x92Ab\n" +
	"`*\x15Metric Threshold Rule20Defines a rule for metric-based threshold alerts\xd2\x01\tcondition\xd2\x01\boverride\"\x98\x06\n" +
	"\x18MetricThresholdCondition\x12\x83\x01\n" +
	"\tthreshold\x18\x01 \x01(\v2\x1c.google.protobuf.DoubleValueBG\x92AD2+The threshold value for the alert conditionJ\x05100.0\x8a\x01\r^\\d+(\\.\\d+)?$R\tthreshold\x12\xa8\x01\n" +
	"\ffor_over_pct\x18\x02 \x01(\v2\x1c.google.protobuf.UInt32ValueBh\x92Ae2LThe percentage of values that must exceed the threshold to trigger the alertJ\x0280\x8a\x01\x10^(100|[1-9]?\\d)$R\n" +
	"forOverPct\x12{\n" +
	"\vof_the_last\x18\x03 \x01(\v2-.com.coralogixapis.alerts.v3.MetricTimeWindowB,\x92A)2'The time window for the alert conditionR\tofTheLast\x12\xbd\x01\n" +
	"\x0econdition_type\x18\x05 \x01(\x0e29.com.coralogixapis.alerts.v3.MetricThresholdConditionTypeB[\x92AX2\x1fThe type of the alert conditionJ5\"METRIC_THRESHOLD_CONDITION_TYPE_MORE_THAN_OR_EQUALS\"R\rconditionType:\x8d\x01\x92A\x89\x01\n" +
	"\x86\x01*\x1aMetric Threshold Condition2.Defines conditions for metric threshold alerts\xd2\x01\tthreshold\xd2\x01\ffor_over_pct\xd2\x01\vof_the_last\xd2\x01\x0econdition_typeb\x06proto3"

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_goTypes = []any{
	(*MetricThresholdType)(nil),        // 0: com.coralogixapis.alerts.v3.MetricThresholdType
	(*MetricThresholdRule)(nil),        // 1: com.coralogixapis.alerts.v3.MetricThresholdRule
	(*MetricThresholdCondition)(nil),   // 2: com.coralogixapis.alerts.v3.MetricThresholdCondition
	(*MetricFilter)(nil),               // 3: com.coralogixapis.alerts.v3.MetricFilter
	(*UndetectedValuesManagement)(nil), // 4: com.coralogixapis.alerts.v3.UndetectedValuesManagement
	(*MetricMissingValues)(nil),        // 5: com.coralogixapis.alerts.v3.MetricMissingValues
	(*wrapperspb.Int32Value)(nil),      // 6: google.protobuf.Int32Value
	(*AlertDefOverride)(nil),           // 7: com.coralogixapis.alerts.v3.AlertDefOverride
	(*wrapperspb.DoubleValue)(nil),     // 8: google.protobuf.DoubleValue
	(*wrapperspb.UInt32Value)(nil),     // 9: google.protobuf.UInt32Value
	(*MetricTimeWindow)(nil),           // 10: com.coralogixapis.alerts.v3.MetricTimeWindow
	(MetricThresholdConditionType)(0),  // 11: com.coralogixapis.alerts.v3.MetricThresholdConditionType
}
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_depIdxs = []int32{
	3,  // 0: com.coralogixapis.alerts.v3.MetricThresholdType.metric_filter:type_name -> com.coralogixapis.alerts.v3.MetricFilter
	1,  // 1: com.coralogixapis.alerts.v3.MetricThresholdType.rules:type_name -> com.coralogixapis.alerts.v3.MetricThresholdRule
	4,  // 2: com.coralogixapis.alerts.v3.MetricThresholdType.undetected_values_management:type_name -> com.coralogixapis.alerts.v3.UndetectedValuesManagement
	5,  // 3: com.coralogixapis.alerts.v3.MetricThresholdType.missing_values:type_name -> com.coralogixapis.alerts.v3.MetricMissingValues
	6,  // 4: com.coralogixapis.alerts.v3.MetricThresholdType.evaluation_delay_ms:type_name -> google.protobuf.Int32Value
	2,  // 5: com.coralogixapis.alerts.v3.MetricThresholdRule.condition:type_name -> com.coralogixapis.alerts.v3.MetricThresholdCondition
	7,  // 6: com.coralogixapis.alerts.v3.MetricThresholdRule.override:type_name -> com.coralogixapis.alerts.v3.AlertDefOverride
	8,  // 7: com.coralogixapis.alerts.v3.MetricThresholdCondition.threshold:type_name -> google.protobuf.DoubleValue
	9,  // 8: com.coralogixapis.alerts.v3.MetricThresholdCondition.for_over_pct:type_name -> google.protobuf.UInt32Value
	10, // 9: com.coralogixapis.alerts.v3.MetricThresholdCondition.of_the_last:type_name -> com.coralogixapis.alerts.v3.MetricTimeWindow
	11, // 10: com.coralogixapis.alerts.v3.MetricThresholdCondition.condition_type:type_name -> com.coralogixapis.alerts.v3.MetricThresholdConditionType
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_init()
}
func file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_condition_type_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_missing_values_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_undetected_values_management_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_depIdxs = nil
}
