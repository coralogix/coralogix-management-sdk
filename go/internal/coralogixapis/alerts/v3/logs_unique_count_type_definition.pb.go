// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.26.1
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/logs/logs_unique_count_type_definition.proto

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

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_rawDesc = []byte{
	0x0a, 0x62, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x67,
	0x73, 0x5f, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76,
	0x33, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x53, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x6a, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f,
	0x67, 0x73, 0x2f, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xba, 0x03, 0x0a, 0x13, 0x4c, 0x6f, 0x67, 0x73, 0x55, 0x6e, 0x69, 0x71, 0x75,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x6c, 0x6f,
	0x67, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f,
	0x67, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x73, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76,
	0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x5c, 0x0a, 0x1b,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x19, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x63, 0x0a, 0x21, 0x6d, 0x61,
	0x78, 0x5f, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70,
	0x65, 0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x62, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x1b, 0x6d, 0x61, 0x78, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x50, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x12,
	0x4e, 0x0a, 0x14, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x6b, 0x65, 0x79, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x12, 0x75, 0x6e, 0x69,
	0x71, 0x75, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x70, 0x61, 0x74, 0x68, 0x22,
	0x6a, 0x0a, 0x13, 0x4c, 0x6f, 0x67, 0x73, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x53, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x55, 0x6e, 0x69, 0x71,
	0x75, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xba, 0x01, 0x0a, 0x18,
	0x4c, 0x6f, 0x67, 0x73, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x10, 0x6d, 0x61, 0x78, 0x5f,
	0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x0e, 0x6d, 0x61, 0x78, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x57, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e,
	0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x52, 0x0a, 0x74, 0x69,
	0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_rawDescData = file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_rawDescData)
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
			RawDescriptor: file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_rawDesc,
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
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_depIdxs = nil
}
