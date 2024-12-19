// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.26.1
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/logs/logs_time_relative_threshold_type_definition.proto

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

type LogsTimeRelativeThresholdType struct {
	state                      protoimpl.MessageState      `protogen:"open.v1"`
	LogsFilter                 *LogsFilter                 `protobuf:"bytes,1,opt,name=logs_filter,json=logsFilter,proto3" json:"logs_filter,omitempty"`
	Rules                      []*LogsTimeRelativeRule     `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty"`
	IgnoreInfinity             *wrapperspb.BoolValue       `protobuf:"bytes,3,opt,name=ignore_infinity,json=ignoreInfinity,proto3" json:"ignore_infinity,omitempty"`
	NotificationPayloadFilter  []*wrapperspb.StringValue   `protobuf:"bytes,4,rep,name=notification_payload_filter,json=notificationPayloadFilter,proto3" json:"notification_payload_filter,omitempty"`
	UndetectedValuesManagement *UndetectedValuesManagement `protobuf:"bytes,5,opt,name=undetected_values_management,json=undetectedValuesManagement,proto3" json:"undetected_values_management,omitempty"`
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

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_rawDesc = []byte{
	0x0a, 0x6d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x67,
	0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f,
	0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x53, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f,
	0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x6d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x2f, 0x6c, 0x6f, 0x67, 0x73,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x58, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x75, 0x6e, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x70, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x64, 0x65,
	0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x76, 0x65, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x45, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f,
	0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x03, 0x0a, 0x1d, 0x4c, 0x6f, 0x67, 0x73, 0x54, 0x69, 0x6d, 0x65,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x47, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67,
	0x73, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x52, 0x75, 0x6c,
	0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x43, 0x0a, 0x0f, 0x69, 0x67, 0x6e, 0x6f,
	0x72, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x69,
	0x67, 0x6e, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x79, 0x12, 0x5c, 0x0a,
	0x1b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x19, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x79, 0x0a, 0x1c, 0x75,
	0x6e, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e,
	0x55, 0x6e, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x1a, 0x75, 0x6e, 0x64, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xb7, 0x01, 0x0a, 0x14, 0x4c, 0x6f, 0x67, 0x73, 0x54,
	0x69, 0x6d, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x12,
	0x54, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33,
	0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x44, 0x65, 0x66, 0x4f, 0x76,
	0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x52, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65,
	0x22, 0x94, 0x02, 0x0a, 0x19, 0x4c, 0x6f, 0x67, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a,
	0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x58, 0x0a, 0x0b, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f,
	0x67, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x72, 0x65, 0x64, 0x54, 0x6f, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72,
	0x65, 0x64, 0x54, 0x6f, 0x12, 0x61, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3a, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x54,
	0x69, 0x6d, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_rawDescData = file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_rawDescData)
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
	(*AlertDefOverride)(nil),              // 7: com.coralogixapis.alerts.v3.AlertDefOverride
	(*wrapperspb.DoubleValue)(nil),        // 8: google.protobuf.DoubleValue
	(LogsTimeRelativeComparedTo)(0),       // 9: com.coralogixapis.alerts.v3.LogsTimeRelativeComparedTo
	(LogsTimeRelativeConditionType)(0),    // 10: com.coralogixapis.alerts.v3.LogsTimeRelativeConditionType
}
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_depIdxs = []int32{
	3,  // 0: com.coralogixapis.alerts.v3.LogsTimeRelativeThresholdType.logs_filter:type_name -> com.coralogixapis.alerts.v3.LogsFilter
	1,  // 1: com.coralogixapis.alerts.v3.LogsTimeRelativeThresholdType.rules:type_name -> com.coralogixapis.alerts.v3.LogsTimeRelativeRule
	4,  // 2: com.coralogixapis.alerts.v3.LogsTimeRelativeThresholdType.ignore_infinity:type_name -> google.protobuf.BoolValue
	5,  // 3: com.coralogixapis.alerts.v3.LogsTimeRelativeThresholdType.notification_payload_filter:type_name -> google.protobuf.StringValue
	6,  // 4: com.coralogixapis.alerts.v3.LogsTimeRelativeThresholdType.undetected_values_management:type_name -> com.coralogixapis.alerts.v3.UndetectedValuesManagement
	2,  // 5: com.coralogixapis.alerts.v3.LogsTimeRelativeRule.condition:type_name -> com.coralogixapis.alerts.v3.LogsTimeRelativeCondition
	7,  // 6: com.coralogixapis.alerts.v3.LogsTimeRelativeRule.override:type_name -> com.coralogixapis.alerts.v3.AlertDefOverride
	8,  // 7: com.coralogixapis.alerts.v3.LogsTimeRelativeCondition.threshold:type_name -> google.protobuf.DoubleValue
	9,  // 8: com.coralogixapis.alerts.v3.LogsTimeRelativeCondition.compared_to:type_name -> com.coralogixapis.alerts.v3.LogsTimeRelativeComparedTo
	10, // 9: com.coralogixapis.alerts.v3.LogsTimeRelativeCondition.condition_type:type_name -> com.coralogixapis.alerts.v3.LogsTimeRelativeConditionType
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
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
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_undetected_values_management_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_time_relative_logs_time_relative_condition_type_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_overrides_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_rawDesc,
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
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_depIdxs = nil
}
