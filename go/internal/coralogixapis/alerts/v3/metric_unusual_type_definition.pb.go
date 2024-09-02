// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.3
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/metric/metric_unusual_type_definition.proto

package v3

import (
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

type MetricUnusualType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetricFilter *MetricFilter        `protobuf:"bytes,1,opt,name=metric_filter,json=metricFilter,proto3" json:"metric_filter,omitempty"`
	Rules        []*MetricUnusualRule `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *MetricUnusualType) Reset() {
	*x = MetricUnusualType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricUnusualType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricUnusualType) ProtoMessage() {}

func (x *MetricUnusualType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricUnusualType.ProtoReflect.Descriptor instead.
func (*MetricUnusualType) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_rawDescGZIP(), []int{0}
}

func (x *MetricUnusualType) GetMetricFilter() *MetricFilter {
	if x != nil {
		return x.MetricFilter
	}
	return nil
}

func (x *MetricUnusualType) GetRules() []*MetricUnusualRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

type MetricUnusualRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Condition *MetricUnusualCondition `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition,omitempty"`
}

func (x *MetricUnusualRule) Reset() {
	*x = MetricUnusualRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricUnusualRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricUnusualRule) ProtoMessage() {}

func (x *MetricUnusualRule) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricUnusualRule.ProtoReflect.Descriptor instead.
func (*MetricUnusualRule) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_rawDescGZIP(), []int{1}
}

func (x *MetricUnusualRule) GetCondition() *MetricUnusualCondition {
	if x != nil {
		return x.Condition
	}
	return nil
}

type MetricUnusualCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Threshold           *wrapperspb.DoubleValue    `protobuf:"bytes,1,opt,name=threshold,proto3" json:"threshold,omitempty"`
	ForOverPct          *wrapperspb.UInt32Value    `protobuf:"bytes,2,opt,name=for_over_pct,json=forOverPct,proto3" json:"for_over_pct,omitempty"`
	OfTheLast           *MetricTimeWindow          `protobuf:"bytes,3,opt,name=of_the_last,json=ofTheLast,proto3" json:"of_the_last,omitempty"`
	MinNonNullValuesPct *wrapperspb.UInt32Value    `protobuf:"bytes,4,opt,name=min_non_null_values_pct,json=minNonNullValuesPct,proto3" json:"min_non_null_values_pct,omitempty"`
	ConditionType       MetricUnusualConditionType `protobuf:"varint,5,opt,name=condition_type,json=conditionType,proto3,enum=com.coralogixapis.alerts.v3.MetricUnusualConditionType" json:"condition_type,omitempty"`
}

func (x *MetricUnusualCondition) Reset() {
	*x = MetricUnusualCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricUnusualCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricUnusualCondition) ProtoMessage() {}

func (x *MetricUnusualCondition) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricUnusualCondition.ProtoReflect.Descriptor instead.
func (*MetricUnusualCondition) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_rawDescGZIP(), []int{2}
}

func (x *MetricUnusualCondition) GetThreshold() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Threshold
	}
	return nil
}

func (x *MetricUnusualCondition) GetForOverPct() *wrapperspb.UInt32Value {
	if x != nil {
		return x.ForOverPct
	}
	return nil
}

func (x *MetricUnusualCondition) GetOfTheLast() *MetricTimeWindow {
	if x != nil {
		return x.OfTheLast
	}
	return nil
}

func (x *MetricUnusualCondition) GetMinNonNullValuesPct() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MinNonNullValuesPct
	}
	return nil
}

func (x *MetricUnusualCondition) GetConditionType() MetricUnusualConditionType {
	if x != nil {
		return x.ConditionType
	}
	return MetricUnusualConditionType_METRIC_UNUSUAL_CONDITION_TYPE_MORE_THAN_USUAL_OR_UNSPECIFIED
}

var File_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_rawDesc = []byte{
	0x0a, 0x61, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x75, 0x6e, 0x75, 0x73, 0x75, 0x61, 0x6c, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x5b, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x58, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x67, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x2f, 0x75, 0x6e, 0x75, 0x73, 0x75, 0x61, 0x6c, 0x2f, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x5f, 0x75, 0x6e, 0x73, 0x75, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa9, 0x01, 0x0a, 0x11, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x55, 0x6e, 0x75, 0x73, 0x75,
	0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x55, 0x6e, 0x75, 0x73, 0x75, 0x61,
	0x6c, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x66, 0x0a, 0x11,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x55, 0x6e, 0x75, 0x73, 0x75, 0x61, 0x6c, 0x52, 0x75, 0x6c,
	0x65, 0x12, 0x51, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e,
	0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x55, 0x6e, 0x75, 0x73, 0x75, 0x61, 0x6c,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x97, 0x03, 0x0a, 0x16, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x55,
	0x6e, 0x75, 0x73, 0x75, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x3a, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x3e, 0x0a, 0x0c, 0x66,
	0x6f, 0x72, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x0a, 0x66, 0x6f, 0x72, 0x4f, 0x76, 0x65, 0x72, 0x50, 0x63, 0x74, 0x12, 0x4d, 0x0a, 0x0b, 0x6f,
	0x66, 0x5f, 0x74, 0x68, 0x65, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x52,
	0x09, 0x6f, 0x66, 0x54, 0x68, 0x65, 0x4c, 0x61, 0x73, 0x74, 0x12, 0x52, 0x0a, 0x17, 0x6d, 0x69,
	0x6e, 0x5f, 0x6e, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6c, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x5f, 0x70, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x13, 0x6d, 0x69, 0x6e, 0x4e, 0x6f,
	0x6e, 0x4e, 0x75, 0x6c, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x50, 0x63, 0x74, 0x12, 0x5e,
	0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x55, 0x6e, 0x75, 0x73, 0x75,
	0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0d, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_rawDescData = file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_goTypes = []interface{}{
	(*MetricUnusualType)(nil),       // 0: com.coralogixapis.alerts.v3.MetricUnusualType
	(*MetricUnusualRule)(nil),       // 1: com.coralogixapis.alerts.v3.MetricUnusualRule
	(*MetricUnusualCondition)(nil),  // 2: com.coralogixapis.alerts.v3.MetricUnusualCondition
	(*MetricFilter)(nil),            // 3: com.coralogixapis.alerts.v3.MetricFilter
	(*wrapperspb.DoubleValue)(nil),  // 4: google.protobuf.DoubleValue
	(*wrapperspb.UInt32Value)(nil),  // 5: google.protobuf.UInt32Value
	(*MetricTimeWindow)(nil),        // 6: com.coralogixapis.alerts.v3.MetricTimeWindow
	(MetricUnusualConditionType)(0), // 7: com.coralogixapis.alerts.v3.MetricUnusualConditionType
}
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_depIdxs = []int32{
	3, // 0: com.coralogixapis.alerts.v3.MetricUnusualType.metric_filter:type_name -> com.coralogixapis.alerts.v3.MetricFilter
	1, // 1: com.coralogixapis.alerts.v3.MetricUnusualType.rules:type_name -> com.coralogixapis.alerts.v3.MetricUnusualRule
	2, // 2: com.coralogixapis.alerts.v3.MetricUnusualRule.condition:type_name -> com.coralogixapis.alerts.v3.MetricUnusualCondition
	4, // 3: com.coralogixapis.alerts.v3.MetricUnusualCondition.threshold:type_name -> google.protobuf.DoubleValue
	5, // 4: com.coralogixapis.alerts.v3.MetricUnusualCondition.for_over_pct:type_name -> google.protobuf.UInt32Value
	6, // 5: com.coralogixapis.alerts.v3.MetricUnusualCondition.of_the_last:type_name -> com.coralogixapis.alerts.v3.MetricTimeWindow
	5, // 6: com.coralogixapis.alerts.v3.MetricUnusualCondition.min_non_null_values_pct:type_name -> google.protobuf.UInt32Value
	7, // 7: com.coralogixapis.alerts.v3.MetricUnusualCondition.condition_type:type_name -> com.coralogixapis.alerts.v3.MetricUnusualConditionType
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_init()
}
func file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_unusual_metric_unsual_condition_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricUnusualType); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricUnusualRule); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricUnusualCondition); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_unusual_type_definition_proto_depIdxs = nil
}
