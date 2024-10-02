// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/logs/logs_unusual_type_definition.proto

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

type LogsUnusualType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogsFilter                *LogsFilter               `protobuf:"bytes,1,opt,name=logs_filter,json=logsFilter,proto3" json:"logs_filter,omitempty"`
	Rules                     []*LogsUnusualRule        `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty"`
	NotificationPayloadFilter []*wrapperspb.StringValue `protobuf:"bytes,3,rep,name=notification_payload_filter,json=notificationPayloadFilter,proto3" json:"notification_payload_filter,omitempty"`
}

func (x *LogsUnusualType) Reset() {
	*x = LogsUnusualType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsUnusualType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsUnusualType) ProtoMessage() {}

func (x *LogsUnusualType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsUnusualType.ProtoReflect.Descriptor instead.
func (*LogsUnusualType) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_rawDescGZIP(), []int{0}
}

func (x *LogsUnusualType) GetLogsFilter() *LogsFilter {
	if x != nil {
		return x.LogsFilter
	}
	return nil
}

func (x *LogsUnusualType) GetRules() []*LogsUnusualRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *LogsUnusualType) GetNotificationPayloadFilter() []*wrapperspb.StringValue {
	if x != nil {
		return x.NotificationPayloadFilter
	}
	return nil
}

type LogsUnusualRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Condition *LogsUnusualCondition `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition,omitempty"`
}

func (x *LogsUnusualRule) Reset() {
	*x = LogsUnusualRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsUnusualRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsUnusualRule) ProtoMessage() {}

func (x *LogsUnusualRule) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsUnusualRule.ProtoReflect.Descriptor instead.
func (*LogsUnusualRule) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_rawDescGZIP(), []int{1}
}

func (x *LogsUnusualRule) GetCondition() *LogsUnusualCondition {
	if x != nil {
		return x.Condition
	}
	return nil
}

type LogsUnusualCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinimumThreshold *wrapperspb.DoubleValue  `protobuf:"bytes,1,opt,name=minimum_threshold,json=minimumThreshold,proto3" json:"minimum_threshold,omitempty"`
	TimeWindow       *LogsTimeWindow          `protobuf:"bytes,2,opt,name=time_window,json=timeWindow,proto3" json:"time_window,omitempty"`
	ConditionType    LogsUnusualConditionType `protobuf:"varint,3,opt,name=condition_type,json=conditionType,proto3,enum=com.coralogixapis.alerts.v3.LogsUnusualConditionType" json:"condition_type,omitempty"`
}

func (x *LogsUnusualCondition) Reset() {
	*x = LogsUnusualCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsUnusualCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsUnusualCondition) ProtoMessage() {}

func (x *LogsUnusualCondition) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsUnusualCondition.ProtoReflect.Descriptor instead.
func (*LogsUnusualCondition) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_rawDescGZIP(), []int{2}
}

func (x *LogsUnusualCondition) GetMinimumThreshold() *wrapperspb.DoubleValue {
	if x != nil {
		return x.MinimumThreshold
	}
	return nil
}

func (x *LogsUnusualCondition) GetTimeWindow() *LogsTimeWindow {
	if x != nil {
		return x.TimeWindow
	}
	return nil
}

func (x *LogsUnusualCondition) GetConditionType() LogsUnusualConditionType {
	if x != nil {
		return x.ConditionType
	}
	return LogsUnusualConditionType_LOGS_UNUSUAL_CONDITION_TYPE_MORE_THAN_USUAL_OR_UNSPECIFIED
}

var File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_rawDesc = []byte{
	0x0a, 0x5d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x67,
	0x73, 0x5f, 0x75, 0x6e, 0x75, 0x73, 0x75, 0x61, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64,
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
	0x6f, 0x1a, 0x57, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x77, 0x69,
	0x6e, 0x64, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x63, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x64, 0x65,
	0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x75, 0x6e, 0x75, 0x73, 0x75, 0x61, 0x6c, 0x2f, 0x6c,
	0x6f, 0x67, 0x73, 0x5f, 0x75, 0x6e, 0x73, 0x75, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xfd, 0x01, 0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x73, 0x55, 0x6e, 0x75, 0x73, 0x75, 0x61, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x42, 0x0a,
	0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x55,
	0x6e, 0x75, 0x73, 0x75, 0x61, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65,
	0x73, 0x12, 0x5c, 0x0a, 0x1b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x19, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22,
	0x62, 0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x73, 0x55, 0x6e, 0x75, 0x73, 0x75, 0x61, 0x6c, 0x52, 0x75,
	0x6c, 0x65, 0x12, 0x4f, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x55, 0x6e, 0x75, 0x73, 0x75, 0x61, 0x6c, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x8d, 0x02, 0x0a, 0x14, 0x4c, 0x6f, 0x67, 0x73, 0x55, 0x6e, 0x75, 0x73,
	0x75, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x11,
	0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x54, 0x68,
	0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x4c, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x54,
	0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x57,
	0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x5c, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x73,
	0x55, 0x6e, 0x75, 0x73, 0x75, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_rawDescData = file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_goTypes = []any{
	(*LogsUnusualType)(nil),        // 0: com.coralogixapis.alerts.v3.LogsUnusualType
	(*LogsUnusualRule)(nil),        // 1: com.coralogixapis.alerts.v3.LogsUnusualRule
	(*LogsUnusualCondition)(nil),   // 2: com.coralogixapis.alerts.v3.LogsUnusualCondition
	(*LogsFilter)(nil),             // 3: com.coralogixapis.alerts.v3.LogsFilter
	(*wrapperspb.StringValue)(nil), // 4: google.protobuf.StringValue
	(*wrapperspb.DoubleValue)(nil), // 5: google.protobuf.DoubleValue
	(*LogsTimeWindow)(nil),         // 6: com.coralogixapis.alerts.v3.LogsTimeWindow
	(LogsUnusualConditionType)(0),  // 7: com.coralogixapis.alerts.v3.LogsUnusualConditionType
}
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_depIdxs = []int32{
	3, // 0: com.coralogixapis.alerts.v3.LogsUnusualType.logs_filter:type_name -> com.coralogixapis.alerts.v3.LogsFilter
	1, // 1: com.coralogixapis.alerts.v3.LogsUnusualType.rules:type_name -> com.coralogixapis.alerts.v3.LogsUnusualRule
	4, // 2: com.coralogixapis.alerts.v3.LogsUnusualType.notification_payload_filter:type_name -> google.protobuf.StringValue
	2, // 3: com.coralogixapis.alerts.v3.LogsUnusualRule.condition:type_name -> com.coralogixapis.alerts.v3.LogsUnusualCondition
	5, // 4: com.coralogixapis.alerts.v3.LogsUnusualCondition.minimum_threshold:type_name -> google.protobuf.DoubleValue
	6, // 5: com.coralogixapis.alerts.v3.LogsUnusualCondition.time_window:type_name -> com.coralogixapis.alerts.v3.LogsTimeWindow
	7, // 6: com.coralogixapis.alerts.v3.LogsUnusualCondition.condition_type:type_name -> com.coralogixapis.alerts.v3.LogsUnusualConditionType
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_init()
}
func file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_timewindow_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_unusual_logs_unsual_condition_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*LogsUnusualType); i {
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
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*LogsUnusualRule); i {
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
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*LogsUnusualCondition); i {
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
			RawDescriptor: file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unusual_type_definition_proto_depIdxs = nil
}
