// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.1
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/logs/logs_ratio_threshold_type_definition.proto

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

type LogsRatioThresholdType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Numerator                  *LogsFilter                 `protobuf:"bytes,1,opt,name=numerator,proto3" json:"numerator,omitempty"`
	NumeratorAlias             *wrapperspb.StringValue     `protobuf:"bytes,2,opt,name=numerator_alias,json=numeratorAlias,proto3" json:"numerator_alias,omitempty"`
	Denominator                *LogsFilter                 `protobuf:"bytes,3,opt,name=denominator,proto3" json:"denominator,omitempty"`
	DenominatorAlias           *wrapperspb.StringValue     `protobuf:"bytes,4,opt,name=denominator_alias,json=denominatorAlias,proto3" json:"denominator_alias,omitempty"`
	Rules                      []*LogsRatioRules           `protobuf:"bytes,5,rep,name=rules,proto3" json:"rules,omitempty"`
	IgnoreInfinity             *wrapperspb.BoolValue       `protobuf:"bytes,9,opt,name=ignore_infinity,json=ignoreInfinity,proto3" json:"ignore_infinity,omitempty"`
	NotificationPayloadFilter  []*wrapperspb.StringValue   `protobuf:"bytes,6,rep,name=notification_payload_filter,json=notificationPayloadFilter,proto3" json:"notification_payload_filter,omitempty"`
	GroupByFor                 LogsRatioGroupByFor         `protobuf:"varint,7,opt,name=group_by_for,json=groupByFor,proto3,enum=com.coralogixapis.alerts.v3.LogsRatioGroupByFor" json:"group_by_for,omitempty"`
	UndetectedValuesManagement *UndetectedValuesManagement `protobuf:"bytes,8,opt,name=undetected_values_management,json=undetectedValuesManagement,proto3" json:"undetected_values_management,omitempty"`
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

func (x *LogsRatioThresholdType) GetIgnoreInfinity() *wrapperspb.BoolValue {
	if x != nil {
		return x.IgnoreInfinity
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

type LogsRatioRules struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Condition *LogsRatioCondition `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition,omitempty"`
	Override  *AlertDefOverride   `protobuf:"bytes,2,opt,name=override,proto3" json:"override,omitempty"`
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Threshold     *wrapperspb.DoubleValue `protobuf:"bytes,1,opt,name=threshold,proto3" json:"threshold,omitempty"`
	TimeWindow    *LogsRatioTimeWindow    `protobuf:"bytes,2,opt,name=time_window,json=timeWindow,proto3" json:"time_window,omitempty"`
	ConditionType LogsRatioConditionType  `protobuf:"varint,4,opt,name=condition_type,json=conditionType,proto3,enum=com.coralogixapis.alerts.v3.LogsRatioConditionType" json:"condition_type,omitempty"`
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

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_rawDesc = []byte{
	0x0a, 0x65, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x67,
	0x73, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x73, 0x2e, 0x76, 0x33, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x53, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76,
	0x33, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x73,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x5c, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x73,
	0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x77, 0x69, 0x6e, 0x64, 0x6f,
	0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x58, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x75,
	0x6e, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x5e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x5f, 0x62, 0x79, 0x5f, 0x66, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x60, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x5f, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x45, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33,
	0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x76, 0x65, 0x72, 0x72,
	0x69, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf1, 0x05, 0x0a, 0x16, 0x4c,
	0x6f, 0x67, 0x73, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x45, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x52, 0x09, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x45, 0x0a, 0x0f,
	0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0e, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x6c,
	0x69, 0x61, 0x73, 0x12, 0x49, 0x0a, 0x0b, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x74,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x52, 0x0b, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x49,
	0x0a, 0x11, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x6c,
	0x69, 0x61, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x6e,
	0x61, 0x74, 0x6f, 0x72, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x41, 0x0a, 0x05, 0x72, 0x75, 0x6c,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x61, 0x74, 0x69, 0x6f,
	0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x43, 0x0a, 0x0f,
	0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x79, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0e, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x79, 0x12, 0x5c, 0x0a, 0x1b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x19, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x52, 0x0a, 0x0c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x62, 0x79, 0x5f, 0x66, 0x6f, 0x72, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x42, 0x79, 0x46, 0x6f, 0x72, 0x52, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x79,
	0x46, 0x6f, 0x72, 0x12, 0x79, 0x0a, 0x1c, 0x75, 0x6e, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x6e, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x1a, 0x75, 0x6e, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xaa,
	0x01, 0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x52, 0x75, 0x6c, 0x65,
	0x73, 0x12, 0x4d, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e,
	0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x43, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x49, 0x0a, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33,
	0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x44, 0x65, 0x66, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64,
	0x65, 0x52, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x22, 0xff, 0x01, 0x0a, 0x12,
	0x4c, 0x6f, 0x67, 0x73, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x51,
	0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76,
	0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x57,
	0x69, 0x6e, 0x64, 0x6f, 0x77, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f,
	0x77, 0x12, 0x5a, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x61, 0x74, 0x69,
	0x6f, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d,
	0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_rawDescData = file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_rawDescData)
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
	(*wrapperspb.BoolValue)(nil),       // 5: google.protobuf.BoolValue
	(LogsRatioGroupByFor)(0),           // 6: com.coralogixapis.alerts.v3.LogsRatioGroupByFor
	(*UndetectedValuesManagement)(nil), // 7: com.coralogixapis.alerts.v3.UndetectedValuesManagement
	(*AlertDefOverride)(nil),           // 8: com.coralogixapis.alerts.v3.AlertDefOverride
	(*wrapperspb.DoubleValue)(nil),     // 9: google.protobuf.DoubleValue
	(*LogsRatioTimeWindow)(nil),        // 10: com.coralogixapis.alerts.v3.LogsRatioTimeWindow
	(LogsRatioConditionType)(0),        // 11: com.coralogixapis.alerts.v3.LogsRatioConditionType
}
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_depIdxs = []int32{
	3,  // 0: com.coralogixapis.alerts.v3.LogsRatioThresholdType.numerator:type_name -> com.coralogixapis.alerts.v3.LogsFilter
	4,  // 1: com.coralogixapis.alerts.v3.LogsRatioThresholdType.numerator_alias:type_name -> google.protobuf.StringValue
	3,  // 2: com.coralogixapis.alerts.v3.LogsRatioThresholdType.denominator:type_name -> com.coralogixapis.alerts.v3.LogsFilter
	4,  // 3: com.coralogixapis.alerts.v3.LogsRatioThresholdType.denominator_alias:type_name -> google.protobuf.StringValue
	1,  // 4: com.coralogixapis.alerts.v3.LogsRatioThresholdType.rules:type_name -> com.coralogixapis.alerts.v3.LogsRatioRules
	5,  // 5: com.coralogixapis.alerts.v3.LogsRatioThresholdType.ignore_infinity:type_name -> google.protobuf.BoolValue
	4,  // 6: com.coralogixapis.alerts.v3.LogsRatioThresholdType.notification_payload_filter:type_name -> google.protobuf.StringValue
	6,  // 7: com.coralogixapis.alerts.v3.LogsRatioThresholdType.group_by_for:type_name -> com.coralogixapis.alerts.v3.LogsRatioGroupByFor
	7,  // 8: com.coralogixapis.alerts.v3.LogsRatioThresholdType.undetected_values_management:type_name -> com.coralogixapis.alerts.v3.UndetectedValuesManagement
	2,  // 9: com.coralogixapis.alerts.v3.LogsRatioRules.condition:type_name -> com.coralogixapis.alerts.v3.LogsRatioCondition
	8,  // 10: com.coralogixapis.alerts.v3.LogsRatioRules.override:type_name -> com.coralogixapis.alerts.v3.AlertDefOverride
	9,  // 11: com.coralogixapis.alerts.v3.LogsRatioCondition.threshold:type_name -> google.protobuf.DoubleValue
	10, // 12: com.coralogixapis.alerts.v3.LogsRatioCondition.time_window:type_name -> com.coralogixapis.alerts.v3.LogsRatioTimeWindow
	11, // 13: com.coralogixapis.alerts.v3.LogsRatioCondition.condition_type:type_name -> com.coralogixapis.alerts.v3.LogsRatioConditionType
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_init()
}
func file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_undetected_values_management_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_group_by_for_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_condition_type_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_rawDesc,
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
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_depIdxs = nil
}
