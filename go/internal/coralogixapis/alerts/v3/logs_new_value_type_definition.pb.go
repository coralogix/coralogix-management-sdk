// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/logs/logs_new_value_type_definition.proto

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

type LogsNewValueType struct {
	state                     protoimpl.MessageState    `protogen:"open.v1"`
	LogsFilter                *LogsFilter               `protobuf:"bytes,1,opt,name=logs_filter,json=logsFilter,proto3" json:"logs_filter,omitempty"`
	Rules                     []*LogsNewValueRule       `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty"`
	NotificationPayloadFilter []*wrapperspb.StringValue `protobuf:"bytes,3,rep,name=notification_payload_filter,json=notificationPayloadFilter,proto3" json:"notification_payload_filter,omitempty"`
	unknownFields             protoimpl.UnknownFields
	sizeCache                 protoimpl.SizeCache
}

func (x *LogsNewValueType) Reset() {
	*x = LogsNewValueType{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsNewValueType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsNewValueType) ProtoMessage() {}

func (x *LogsNewValueType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsNewValueType.ProtoReflect.Descriptor instead.
func (*LogsNewValueType) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_rawDescGZIP(), []int{0}
}

func (x *LogsNewValueType) GetLogsFilter() *LogsFilter {
	if x != nil {
		return x.LogsFilter
	}
	return nil
}

func (x *LogsNewValueType) GetRules() []*LogsNewValueRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *LogsNewValueType) GetNotificationPayloadFilter() []*wrapperspb.StringValue {
	if x != nil {
		return x.NotificationPayloadFilter
	}
	return nil
}

type LogsNewValueRule struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Condition     *LogsNewValueCondition `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogsNewValueRule) Reset() {
	*x = LogsNewValueRule{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsNewValueRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsNewValueRule) ProtoMessage() {}

func (x *LogsNewValueRule) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsNewValueRule.ProtoReflect.Descriptor instead.
func (*LogsNewValueRule) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_rawDescGZIP(), []int{1}
}

func (x *LogsNewValueRule) GetCondition() *LogsNewValueCondition {
	if x != nil {
		return x.Condition
	}
	return nil
}

type LogsNewValueCondition struct {
	state          protoimpl.MessageState  `protogen:"open.v1"`
	KeypathToTrack *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=keypath_to_track,json=keypathToTrack,proto3" json:"keypath_to_track,omitempty"`
	TimeWindow     *LogsNewValueTimeWindow `protobuf:"bytes,2,opt,name=time_window,json=timeWindow,proto3" json:"time_window,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *LogsNewValueCondition) Reset() {
	*x = LogsNewValueCondition{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsNewValueCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsNewValueCondition) ProtoMessage() {}

func (x *LogsNewValueCondition) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsNewValueCondition.ProtoReflect.Descriptor instead.
func (*LogsNewValueCondition) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_rawDescGZIP(), []int{2}
}

func (x *LogsNewValueCondition) GetKeypathToTrack() *wrapperspb.StringValue {
	if x != nil {
		return x.KeypathToTrack
	}
	return nil
}

func (x *LogsNewValueCondition) GetTimeWindow() *LogsNewValueTimeWindow {
	if x != nil {
		return x.TimeWindow
	}
	return nil
}

var File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_rawDesc = []byte{
	0x0a, 0x5f, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x67,
	0x73, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x64,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x6e, 0x65, 0x77, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x53, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76,
	0x33, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x73,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x03, 0x0a, 0x10, 0x4c, 0x6f,
	0x67, 0x73, 0x4e, 0x65, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x48,
	0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76,
	0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0a, 0x6c, 0x6f,
	0x67, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x4e, 0x65, 0x77, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x69, 0x0a,
	0x1b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x0b, 0x92, 0x41, 0x08, 0x4a, 0x06, 0x22, 0x74, 0x65, 0x78, 0x74, 0x22, 0x52, 0x19, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x3a, 0x7f, 0x92, 0x41, 0x7c, 0x0a, 0x7a, 0x2a,
	0x1e, 0x4c, 0x6f, 0x67, 0x2d, 0x62, 0x61, 0x73, 0x65, 0x64, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x20, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x20, 0x74, 0x79, 0x70, 0x65, 0x32,
	0x42, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x20, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x20, 0x61, 0x70, 0x70, 0x65, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x69, 0x6e, 0x20, 0x6c,
	0x6f, 0x67, 0x73, 0xd2, 0x01, 0x0b, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0xd2, 0x01, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x64, 0x0a, 0x10, 0x4c, 0x6f, 0x67,
	0x73, 0x4e, 0x65, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x50, 0x0a,
	0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4c,
	0x6f, 0x67, 0x73, 0x4e, 0x65, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0xc8, 0x02, 0x0a, 0x15, 0x4c, 0x6f, 0x67, 0x73, 0x4e, 0x65, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5d, 0x0a, 0x10, 0x6b, 0x65, 0x79,
	0x70, 0x61, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x15, 0x92, 0x41, 0x12, 0x4a, 0x10, 0x22, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x52, 0x0e, 0x6b, 0x65, 0x79, 0x70, 0x61, 0x74,
	0x68, 0x54, 0x6f, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x54, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x73,
	0x4e, 0x65, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64,
	0x6f, 0x77, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x3a, 0x7a,
	0x92, 0x41, 0x77, 0x0a, 0x75, 0x2a, 0x1d, 0x4c, 0x6f, 0x67, 0x2d, 0x62, 0x61, 0x73, 0x65, 0x64,
	0x20, 0x6e, 0x65, 0x77, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x63, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x32, 0x33, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x73, 0x20, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x64, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x20, 0x69, 0x6e, 0x20, 0x6c, 0x6f, 0x67, 0x73, 0xd2, 0x01, 0x10, 0x6b, 0x65, 0x79, 0x70,
	0x61, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0xd2, 0x01, 0x0b, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_rawDescData = file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_goTypes = []any{
	(*LogsNewValueType)(nil),       // 0: com.coralogixapis.alerts.v3.LogsNewValueType
	(*LogsNewValueRule)(nil),       // 1: com.coralogixapis.alerts.v3.LogsNewValueRule
	(*LogsNewValueCondition)(nil),  // 2: com.coralogixapis.alerts.v3.LogsNewValueCondition
	(*LogsFilter)(nil),             // 3: com.coralogixapis.alerts.v3.LogsFilter
	(*wrapperspb.StringValue)(nil), // 4: google.protobuf.StringValue
	(*LogsNewValueTimeWindow)(nil), // 5: com.coralogixapis.alerts.v3.LogsNewValueTimeWindow
}
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_depIdxs = []int32{
	3, // 0: com.coralogixapis.alerts.v3.LogsNewValueType.logs_filter:type_name -> com.coralogixapis.alerts.v3.LogsFilter
	1, // 1: com.coralogixapis.alerts.v3.LogsNewValueType.rules:type_name -> com.coralogixapis.alerts.v3.LogsNewValueRule
	4, // 2: com.coralogixapis.alerts.v3.LogsNewValueType.notification_payload_filter:type_name -> google.protobuf.StringValue
	2, // 3: com.coralogixapis.alerts.v3.LogsNewValueRule.condition:type_name -> com.coralogixapis.alerts.v3.LogsNewValueCondition
	4, // 4: com.coralogixapis.alerts.v3.LogsNewValueCondition.keypath_to_track:type_name -> google.protobuf.StringValue
	5, // 5: com.coralogixapis.alerts.v3.LogsNewValueCondition.time_window:type_name -> com.coralogixapis.alerts.v3.LogsNewValueTimeWindow
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_init()
}
func file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_depIdxs = nil
}
