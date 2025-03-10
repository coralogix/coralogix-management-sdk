// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.26.1
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/metric/metric_threshold_type_definition.proto

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

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_rawDesc = []byte{
	0x0a, 0x63, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e,
	0x76, 0x33, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x5b, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x58, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x75, 0x6e, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x6c, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f,
	0x6c, 0x64, 0x5f, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x58, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x6c, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2f,
	0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x45, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x04, 0x0a, 0x13, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4e,
	0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x52, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x46,
	0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x52,
	0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x79, 0x0a, 0x1c, 0x75, 0x6e, 0x64, 0x65, 0x74, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x6e, 0x64, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x1a, 0x75, 0x6e, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x57, 0x0a, 0x0e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4d, 0x69,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x0d, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x4b, 0x0a, 0x13, 0x65, 0x76,
	0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x6d,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x11, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x65, 0x6c, 0x61, 0x79, 0x4d, 0x73, 0x3a, 0x79, 0x92, 0x41, 0x76, 0x0a, 0x74, 0x2a, 0x1b,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x20, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x20, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x20, 0x54, 0x79, 0x70, 0x65, 0x32, 0x3d, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x73, 0x20, 0x62, 0x61, 0x73, 0x65, 0x64, 0x20, 0x6f, 0x6e, 0x20, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x20, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x20,
	0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0xd2, 0x01, 0x0d, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0xd2, 0x01, 0x05, 0x72, 0x75, 0x6c,
	0x65, 0x73, 0x22, 0xb5, 0x01, 0x0a, 0x13, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x68, 0x72,
	0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x53, 0x0a, 0x09, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x49, 0x0a, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x44, 0x65, 0x66, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65,
	0x52, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x22, 0xb4, 0x03, 0x0a, 0x18, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x43, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73,
	0x68, 0x6f, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68,
	0x6f, 0x6c, 0x64, 0x12, 0x3e, 0x0a, 0x0c, 0x66, 0x6f, 0x72, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x5f,
	0x70, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x66, 0x6f, 0x72, 0x4f, 0x76, 0x65, 0x72,
	0x50, 0x63, 0x74, 0x12, 0x4d, 0x0a, 0x0b, 0x6f, 0x66, 0x5f, 0x74, 0x68, 0x65, 0x5f, 0x6c, 0x61,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x69, 0x6d,
	0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x52, 0x09, 0x6f, 0x66, 0x54, 0x68, 0x65, 0x4c, 0x61,
	0x73, 0x74, 0x12, 0x60, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x39, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54,
	0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x3a, 0x6b, 0x92, 0x41, 0x68, 0x0a, 0x66, 0x2a, 0x1a, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x20, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x20, 0x43, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x2e, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x73,
	0x20, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x20, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x20, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0xd2, 0x01, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68,
	0x6f, 0x6c, 0x64, 0xd2, 0x01, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f,
	0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_rawDescData = file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_rawDescData)
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
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_undetected_values_management_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_missing_values_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_condition_type_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_rawDesc,
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
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_depIdxs = nil
}
