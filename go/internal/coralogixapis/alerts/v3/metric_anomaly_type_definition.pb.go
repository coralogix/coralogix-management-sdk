// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/metric/metric_anomaly_type_definition.proto

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

type MetricAnomalyType struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	MetricFilter         *MetricFilter          `protobuf:"bytes,1,opt,name=metric_filter,json=metricFilter,proto3" json:"metric_filter,omitempty"`
	Rules                []*MetricAnomalyRule   `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty"`
	EvaluationDelayMs    *wrapperspb.Int32Value `protobuf:"bytes,3,opt,name=evaluation_delay_ms,json=evaluationDelayMs,proto3" json:"evaluation_delay_ms,omitempty"`
	AnomalyAlertSettings *AnomalyAlertSettings  `protobuf:"bytes,4,opt,name=anomaly_alert_settings,json=anomalyAlertSettings,proto3" json:"anomaly_alert_settings,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *MetricAnomalyType) Reset() {
	*x = MetricAnomalyType{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetricAnomalyType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricAnomalyType) ProtoMessage() {}

func (x *MetricAnomalyType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricAnomalyType.ProtoReflect.Descriptor instead.
func (*MetricAnomalyType) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto_rawDescGZIP(), []int{0}
}

func (x *MetricAnomalyType) GetMetricFilter() *MetricFilter {
	if x != nil {
		return x.MetricFilter
	}
	return nil
}

func (x *MetricAnomalyType) GetRules() []*MetricAnomalyRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *MetricAnomalyType) GetEvaluationDelayMs() *wrapperspb.Int32Value {
	if x != nil {
		return x.EvaluationDelayMs
	}
	return nil
}

func (x *MetricAnomalyType) GetAnomalyAlertSettings() *AnomalyAlertSettings {
	if x != nil {
		return x.AnomalyAlertSettings
	}
	return nil
}

type MetricAnomalyRule struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Condition     *MetricAnomalyCondition `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MetricAnomalyRule) Reset() {
	*x = MetricAnomalyRule{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetricAnomalyRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricAnomalyRule) ProtoMessage() {}

func (x *MetricAnomalyRule) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricAnomalyRule.ProtoReflect.Descriptor instead.
func (*MetricAnomalyRule) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto_rawDescGZIP(), []int{1}
}

func (x *MetricAnomalyRule) GetCondition() *MetricAnomalyCondition {
	if x != nil {
		return x.Condition
	}
	return nil
}

type MetricAnomalyCondition struct {
	state               protoimpl.MessageState     `protogen:"open.v1"`
	Threshold           *wrapperspb.DoubleValue    `protobuf:"bytes,1,opt,name=threshold,proto3" json:"threshold,omitempty"`
	ForOverPct          *wrapperspb.UInt32Value    `protobuf:"bytes,2,opt,name=for_over_pct,json=forOverPct,proto3" json:"for_over_pct,omitempty"`
	OfTheLast           *MetricTimeWindow          `protobuf:"bytes,3,opt,name=of_the_last,json=ofTheLast,proto3" json:"of_the_last,omitempty"`
	MinNonNullValuesPct *wrapperspb.UInt32Value    `protobuf:"bytes,4,opt,name=min_non_null_values_pct,json=minNonNullValuesPct,proto3" json:"min_non_null_values_pct,omitempty"`
	ConditionType       MetricAnomalyConditionType `protobuf:"varint,5,opt,name=condition_type,json=conditionType,proto3,enum=com.coralogixapis.alerts.v3.MetricAnomalyConditionType" json:"condition_type,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *MetricAnomalyCondition) Reset() {
	*x = MetricAnomalyCondition{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetricAnomalyCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricAnomalyCondition) ProtoMessage() {}

func (x *MetricAnomalyCondition) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricAnomalyCondition.ProtoReflect.Descriptor instead.
func (*MetricAnomalyCondition) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto_rawDescGZIP(), []int{2}
}

func (x *MetricAnomalyCondition) GetThreshold() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Threshold
	}
	return nil
}

func (x *MetricAnomalyCondition) GetForOverPct() *wrapperspb.UInt32Value {
	if x != nil {
		return x.ForOverPct
	}
	return nil
}

func (x *MetricAnomalyCondition) GetOfTheLast() *MetricTimeWindow {
	if x != nil {
		return x.OfTheLast
	}
	return nil
}

func (x *MetricAnomalyCondition) GetMinNonNullValuesPct() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MinNonNullValuesPct
	}
	return nil
}

func (x *MetricAnomalyCondition) GetConditionType() MetricAnomalyConditionType {
	if x != nil {
		return x.ConditionType
	}
	return MetricAnomalyConditionType_METRIC_ANOMALY_CONDITION_TYPE_MORE_THAN_USUAL_OR_UNSPECIFIED
}

var File_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto protoreflect.FileDescriptor

const file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto_rawDesc = "" +
	"\n" +
	"acom/coralogixapis/alerts/v3/alert_def_type_definition/metric/metric_anomaly_type_definition.proto\x12\x1bcom.coralogixapis.alerts.v3\x1a\x1egoogle/protobuf/wrappers.proto\x1a[com/coralogixapis/alerts/v3/alert_def_type_definition/metric/common/metric_timewindow.proto\x1aXcom/coralogixapis/alerts/v3/alert_def_type_definition/metric/common/metrics_filter.proto\x1ahcom/coralogixapis/alerts/v3/alert_def_type_definition/metric/anomaly/metric_anomaly_condition_type.proto\x1aYcom/coralogixapis/alerts/v3/alert_def_type_definition/common/anomaly_alert_settings.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\x97\x05\n" +
	"\x11MetricAnomalyType\x12N\n" +
	"\rmetric_filter\x18\x01 \x01(\v2).com.coralogixapis.alerts.v3.MetricFilterR\fmetricFilter\x12D\n" +
	"\x05rules\x18\x02 \x03(\v2..com.coralogixapis.alerts.v3.MetricAnomalyRuleR\x05rules\x12W\n" +
	"\x13evaluation_delay_ms\x18\x03 \x01(\v2\x1b.google.protobuf.Int32ValueB\n" +
	"\x92A\aJ\x0560000R\x11evaluationDelayMs\x12g\n" +
	"\x16anomaly_alert_settings\x18\x04 \x01(\v21.com.coralogixapis.alerts.v3.AnomalyAlertSettingsR\x14anomalyAlertSettings:\xa9\x02\x92A\xa5\x02\n" +
	"z*\x1fMetric-based anomaly alert type2?Configuration for alerts triggered by anomalous metric patterns\xd2\x01\rmetric_source\xd2\x01\x05rules*\xa6\x01\n" +
	"ALearn more about metric-based anomaly alerts in our documentation\x12ahttps://coralogix.com/docs/user-guides/alerting/create-an-alert/metrics/anomaly-detection-alerts/\"f\n" +
	"\x11MetricAnomalyRule\x12Q\n" +
	"\tcondition\x18\x01 \x01(\v23.com.coralogixapis.alerts.v3.MetricAnomalyConditionR\tcondition\"\x9b\x04\n" +
	"\x16MetricAnomalyCondition\x12E\n" +
	"\tthreshold\x18\x01 \x01(\v2\x1c.google.protobuf.DoubleValueB\t\x92A\x06J\x0410.0R\tthreshold\x12G\n" +
	"\ffor_over_pct\x18\x02 \x01(\v2\x1c.google.protobuf.UInt32ValueB\a\x92A\x04J\x0210R\n" +
	"forOverPct\x12V\n" +
	"\vof_the_last\x18\x03 \x01(\v2-.com.coralogixapis.alerts.v3.MetricTimeWindowB\a\x92A\x04J\x0210R\tofTheLast\x12[\n" +
	"\x17min_non_null_values_pct\x18\x04 \x01(\v2\x1c.google.protobuf.UInt32ValueB\a\x92A\x04J\x0210R\x13minNonNullValuesPct\x12s\n" +
	"\x0econdition_type\x18\x05 \x01(\x0e27.com.coralogixapis.alerts.v3.MetricAnomalyConditionTypeB\x13\x92A\x10J\x0e\"GREATER_THAN\"R\rconditionType:G\x92AD\n" +
	"B*\x1eMetric-based anomaly condition\xd2\x01\x11minimum_threshold\xd2\x01\vtime_windowb\x06proto3"

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto_goTypes = []any{
	(*MetricAnomalyType)(nil),       // 0: com.coralogixapis.alerts.v3.MetricAnomalyType
	(*MetricAnomalyRule)(nil),       // 1: com.coralogixapis.alerts.v3.MetricAnomalyRule
	(*MetricAnomalyCondition)(nil),  // 2: com.coralogixapis.alerts.v3.MetricAnomalyCondition
	(*MetricFilter)(nil),            // 3: com.coralogixapis.alerts.v3.MetricFilter
	(*wrapperspb.Int32Value)(nil),   // 4: google.protobuf.Int32Value
	(*AnomalyAlertSettings)(nil),    // 5: com.coralogixapis.alerts.v3.AnomalyAlertSettings
	(*wrapperspb.DoubleValue)(nil),  // 6: google.protobuf.DoubleValue
	(*wrapperspb.UInt32Value)(nil),  // 7: google.protobuf.UInt32Value
	(*MetricTimeWindow)(nil),        // 8: com.coralogixapis.alerts.v3.MetricTimeWindow
	(MetricAnomalyConditionType)(0), // 9: com.coralogixapis.alerts.v3.MetricAnomalyConditionType
}
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto_depIdxs = []int32{
	3,  // 0: com.coralogixapis.alerts.v3.MetricAnomalyType.metric_filter:type_name -> com.coralogixapis.alerts.v3.MetricFilter
	1,  // 1: com.coralogixapis.alerts.v3.MetricAnomalyType.rules:type_name -> com.coralogixapis.alerts.v3.MetricAnomalyRule
	4,  // 2: com.coralogixapis.alerts.v3.MetricAnomalyType.evaluation_delay_ms:type_name -> google.protobuf.Int32Value
	5,  // 3: com.coralogixapis.alerts.v3.MetricAnomalyType.anomaly_alert_settings:type_name -> com.coralogixapis.alerts.v3.AnomalyAlertSettings
	2,  // 4: com.coralogixapis.alerts.v3.MetricAnomalyRule.condition:type_name -> com.coralogixapis.alerts.v3.MetricAnomalyCondition
	6,  // 5: com.coralogixapis.alerts.v3.MetricAnomalyCondition.threshold:type_name -> google.protobuf.DoubleValue
	7,  // 6: com.coralogixapis.alerts.v3.MetricAnomalyCondition.for_over_pct:type_name -> google.protobuf.UInt32Value
	8,  // 7: com.coralogixapis.alerts.v3.MetricAnomalyCondition.of_the_last:type_name -> com.coralogixapis.alerts.v3.MetricTimeWindow
	7,  // 8: com.coralogixapis.alerts.v3.MetricAnomalyCondition.min_non_null_values_pct:type_name -> google.protobuf.UInt32Value
	9,  // 9: com.coralogixapis.alerts.v3.MetricAnomalyCondition.condition_type:type_name -> com.coralogixapis.alerts.v3.MetricAnomalyConditionType
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto_init()
}
func file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_anomaly_metric_anomaly_condition_type_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_common_anomaly_alert_settings_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto_depIdxs = nil
}
