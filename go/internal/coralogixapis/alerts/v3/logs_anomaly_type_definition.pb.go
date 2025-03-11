// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/logs/logs_anomaly_type_definition.proto

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

type LogsAnomalyType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogsFilter                *LogsFilter               `protobuf:"bytes,1,opt,name=logs_filter,json=logsFilter,proto3" json:"logs_filter,omitempty"`
	Rules                     []*LogsAnomalyRule        `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty"`
	NotificationPayloadFilter []*wrapperspb.StringValue `protobuf:"bytes,3,rep,name=notification_payload_filter,json=notificationPayloadFilter,proto3" json:"notification_payload_filter,omitempty"`
	EvaluationDelayMs         *wrapperspb.Int32Value    `protobuf:"bytes,4,opt,name=evaluation_delay_ms,json=evaluationDelayMs,proto3" json:"evaluation_delay_ms,omitempty"`
	AnomalyAlertSettings      *AnomalyAlertSettings     `protobuf:"bytes,5,opt,name=anomaly_alert_settings,json=anomalyAlertSettings,proto3" json:"anomaly_alert_settings,omitempty"`
}

func (x *LogsAnomalyType) Reset() {
	*x = LogsAnomalyType{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsAnomalyType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsAnomalyType) ProtoMessage() {}

func (x *LogsAnomalyType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsAnomalyType.ProtoReflect.Descriptor instead.
func (*LogsAnomalyType) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto_rawDescGZIP(), []int{0}
}

func (x *LogsAnomalyType) GetLogsFilter() *LogsFilter {
	if x != nil {
		return x.LogsFilter
	}
	return nil
}

func (x *LogsAnomalyType) GetRules() []*LogsAnomalyRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *LogsAnomalyType) GetNotificationPayloadFilter() []*wrapperspb.StringValue {
	if x != nil {
		return x.NotificationPayloadFilter
	}
	return nil
}

func (x *LogsAnomalyType) GetEvaluationDelayMs() *wrapperspb.Int32Value {
	if x != nil {
		return x.EvaluationDelayMs
	}
	return nil
}

func (x *LogsAnomalyType) GetAnomalyAlertSettings() *AnomalyAlertSettings {
	if x != nil {
		return x.AnomalyAlertSettings
	}
	return nil
}

type LogsAnomalyRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Condition *LogsAnomalyCondition `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition,omitempty"`
}

func (x *LogsAnomalyRule) Reset() {
	*x = LogsAnomalyRule{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsAnomalyRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsAnomalyRule) ProtoMessage() {}

func (x *LogsAnomalyRule) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsAnomalyRule.ProtoReflect.Descriptor instead.
func (*LogsAnomalyRule) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto_rawDescGZIP(), []int{1}
}

func (x *LogsAnomalyRule) GetCondition() *LogsAnomalyCondition {
	if x != nil {
		return x.Condition
	}
	return nil
}

type LogsAnomalyCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinimumThreshold *wrapperspb.DoubleValue  `protobuf:"bytes,1,opt,name=minimum_threshold,json=minimumThreshold,proto3" json:"minimum_threshold,omitempty"`
	TimeWindow       *LogsTimeWindow          `protobuf:"bytes,2,opt,name=time_window,json=timeWindow,proto3" json:"time_window,omitempty"`
	ConditionType    LogsAnomalyConditionType `protobuf:"varint,3,opt,name=condition_type,json=conditionType,proto3,enum=com.coralogixapis.alerts.v3.LogsAnomalyConditionType" json:"condition_type,omitempty"`
}

func (x *LogsAnomalyCondition) Reset() {
	*x = LogsAnomalyCondition{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsAnomalyCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsAnomalyCondition) ProtoMessage() {}

func (x *LogsAnomalyCondition) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsAnomalyCondition.ProtoReflect.Descriptor instead.
func (*LogsAnomalyCondition) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto_rawDescGZIP(), []int{2}
}

func (x *LogsAnomalyCondition) GetMinimumThreshold() *wrapperspb.DoubleValue {
	if x != nil {
		return x.MinimumThreshold
	}
	return nil
}

func (x *LogsAnomalyCondition) GetTimeWindow() *LogsTimeWindow {
	if x != nil {
		return x.TimeWindow
	}
	return nil
}

func (x *LogsAnomalyCondition) GetConditionType() LogsAnomalyConditionType {
	if x != nil {
		return x.ConditionType
	}
	return LogsAnomalyConditionType_LOGS_ANOMALY_CONDITION_TYPE_MORE_THAN_USUAL_OR_UNSPECIFIED
}

var File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto_rawDesc = []byte{
	0x0a, 0x5d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x67,
	0x73, 0x5f, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64,
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
	0x6e, 0x64, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x64, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x64, 0x65,
	0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x2f, 0x6c,
	0x6f, 0x67, 0x73, 0x5f, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x59, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61,
	0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x5f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76,
	0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x04, 0x0a, 0x0f,
	0x4c, 0x6f, 0x67, 0x73, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x48, 0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e,
	0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0a, 0x6c,
	0x6f, 0x67, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x05, 0x72, 0x75, 0x6c,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x41, 0x6e, 0x6f, 0x6d, 0x61,
	0x6c, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x69, 0x0a,
	0x1b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x0b, 0x92, 0x41, 0x08, 0x4a, 0x06, 0x22, 0x74, 0x65, 0x78, 0x74, 0x22, 0x52, 0x19, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x57, 0x0a, 0x13, 0x65, 0x76, 0x61, 0x6c,
	0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x6d, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x0a, 0x92, 0x41, 0x07, 0x4a, 0x05, 0x36, 0x30, 0x30, 0x30, 0x30, 0x52, 0x11,
	0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x4d,
	0x73, 0x12, 0x67, 0x0a, 0x16, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x5f, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e,
	0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x14, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x3a, 0x72, 0x92, 0x41, 0x6f, 0x0a,
	0x6d, 0x2a, 0x17, 0x4c, 0x6f, 0x67, 0x73, 0x20, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x20,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x20, 0x54, 0x79, 0x70, 0x65, 0x32, 0x3c, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x20, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x65, 0x64, 0x20, 0x62,
	0x79, 0x20, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x6f, 0x75, 0x73, 0x20, 0x6c, 0x6f, 0x67, 0x20,
	0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0xd2, 0x01, 0x0b, 0x6c, 0x6f, 0x67, 0x73, 0x5f,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0xd2, 0x01, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x22, 0xb4,
	0x01, 0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x73, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x52, 0x75,
	0x6c, 0x65, 0x12, 0x4f, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x3a, 0x50, 0x92, 0x41, 0x4d, 0x0a, 0x4b, 0x2a, 0x11, 0x4c, 0x6f, 0x67, 0x73,
	0x20, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x20, 0x52, 0x75, 0x6c, 0x65, 0x32, 0x2a, 0x44,
	0x65, 0x66, 0x69, 0x6e, 0x65, 0x73, 0x20, 0x61, 0x20, 0x72, 0x75, 0x6c, 0x65, 0x20, 0x66, 0x6f,
	0x72, 0x20, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x6c, 0x6f, 0x67, 0x20,
	0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x69, 0x65, 0x73, 0xd2, 0x01, 0x09, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x89, 0x03, 0x0a, 0x14, 0x4c, 0x6f, 0x67, 0x73, 0x41, 0x6e,
	0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x54,
	0x0a, 0x11, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68,
	0x6f, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x09, 0x92, 0x41, 0x06, 0x4a, 0x04, 0x31, 0x30,
	0x2e, 0x30, 0x52, 0x10, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x54, 0x68, 0x72, 0x65, 0x73,
	0x68, 0x6f, 0x6c, 0x64, 0x12, 0x4c, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x77, 0x69, 0x6e,
	0x64, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x54, 0x69, 0x6d, 0x65,
	0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64,
	0x6f, 0x77, 0x12, 0x5c, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x41, 0x6e, 0x6f,
	0x6d, 0x61, 0x6c, 0x79, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x3a, 0x6f, 0x92, 0x41, 0x6c, 0x0a, 0x6a, 0x2a, 0x16, 0x4c, 0x6f, 0x67, 0x73, 0x20, 0x41, 0x6e,
	0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x20, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x32,
	0x2e, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x73, 0x20, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6e,
	0x67, 0x20, 0x6c, 0x6f, 0x67, 0x20, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x69, 0x65, 0x73, 0xd2,
	0x01, 0x11, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68,
	0x6f, 0x6c, 0x64, 0xd2, 0x01, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f,
	0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto_rawDescData = file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto_goTypes = []any{
	(*LogsAnomalyType)(nil),        // 0: com.coralogixapis.alerts.v3.LogsAnomalyType
	(*LogsAnomalyRule)(nil),        // 1: com.coralogixapis.alerts.v3.LogsAnomalyRule
	(*LogsAnomalyCondition)(nil),   // 2: com.coralogixapis.alerts.v3.LogsAnomalyCondition
	(*LogsFilter)(nil),             // 3: com.coralogixapis.alerts.v3.LogsFilter
	(*wrapperspb.StringValue)(nil), // 4: google.protobuf.StringValue
	(*wrapperspb.Int32Value)(nil),  // 5: google.protobuf.Int32Value
	(*AnomalyAlertSettings)(nil),   // 6: com.coralogixapis.alerts.v3.AnomalyAlertSettings
	(*wrapperspb.DoubleValue)(nil), // 7: google.protobuf.DoubleValue
	(*LogsTimeWindow)(nil),         // 8: com.coralogixapis.alerts.v3.LogsTimeWindow
	(LogsAnomalyConditionType)(0),  // 9: com.coralogixapis.alerts.v3.LogsAnomalyConditionType
}
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto_depIdxs = []int32{
	3, // 0: com.coralogixapis.alerts.v3.LogsAnomalyType.logs_filter:type_name -> com.coralogixapis.alerts.v3.LogsFilter
	1, // 1: com.coralogixapis.alerts.v3.LogsAnomalyType.rules:type_name -> com.coralogixapis.alerts.v3.LogsAnomalyRule
	4, // 2: com.coralogixapis.alerts.v3.LogsAnomalyType.notification_payload_filter:type_name -> google.protobuf.StringValue
	5, // 3: com.coralogixapis.alerts.v3.LogsAnomalyType.evaluation_delay_ms:type_name -> google.protobuf.Int32Value
	6, // 4: com.coralogixapis.alerts.v3.LogsAnomalyType.anomaly_alert_settings:type_name -> com.coralogixapis.alerts.v3.AnomalyAlertSettings
	2, // 5: com.coralogixapis.alerts.v3.LogsAnomalyRule.condition:type_name -> com.coralogixapis.alerts.v3.LogsAnomalyCondition
	7, // 6: com.coralogixapis.alerts.v3.LogsAnomalyCondition.minimum_threshold:type_name -> google.protobuf.DoubleValue
	8, // 7: com.coralogixapis.alerts.v3.LogsAnomalyCondition.time_window:type_name -> com.coralogixapis.alerts.v3.LogsTimeWindow
	9, // 8: com.coralogixapis.alerts.v3.LogsAnomalyCondition.condition_type:type_name -> com.coralogixapis.alerts.v3.LogsAnomalyConditionType
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto_init()
}
func file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_timewindow_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_anomaly_logs_anomaly_condition_type_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_common_anomaly_alert_settings_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto_depIdxs = nil
}
