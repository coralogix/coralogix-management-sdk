// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.26.1
// source: com/coralogixapis/alerts/v3/notification/notification_event.proto

package v3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type AlertNotificationEvent struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Timestamp     *timestamppb.Timestamp  `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Id            *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Status        AlertStatus             `protobuf:"varint,3,opt,name=status,proto3,enum=com.coralogixapis.alerts.v3.AlertStatus" json:"status,omitempty"`
	Attachments   *Attachments            `protobuf:"bytes,4,opt,name=attachments,proto3" json:"attachments,omitempty"`
	Groups        map[string]*Group       `protobuf:"bytes,5,rep,name=groups,proto3" json:"groups,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AlertNotificationEvent) Reset() {
	*x = AlertNotificationEvent{}
	mi := &file_com_coralogixapis_alerts_v3_notification_notification_event_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlertNotificationEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertNotificationEvent) ProtoMessage() {}

func (x *AlertNotificationEvent) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_notification_notification_event_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertNotificationEvent.ProtoReflect.Descriptor instead.
func (*AlertNotificationEvent) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_notification_notification_event_proto_rawDescGZIP(), []int{0}
}

func (x *AlertNotificationEvent) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *AlertNotificationEvent) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *AlertNotificationEvent) GetStatus() AlertStatus {
	if x != nil {
		return x.Status
	}
	return AlertStatus_ALERT_STATUS_RESOLVED_OR_UNSPECIFIED
}

func (x *AlertNotificationEvent) GetAttachments() *Attachments {
	if x != nil {
		return x.Attachments
	}
	return nil
}

func (x *AlertNotificationEvent) GetGroups() map[string]*Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

type Group struct {
	state      protoimpl.MessageState `protogen:"open.v1"`
	Status     AlertStatus            `protobuf:"varint,1,opt,name=status,proto3,enum=com.coralogixapis.alerts.v3.AlertStatus" json:"status,omitempty"`
	Suppressed *wrapperspb.BoolValue  `protobuf:"bytes,2,opt,name=suppressed,proto3" json:"suppressed,omitempty"`
	Priority   AlertDefPriority       `protobuf:"varint,3,opt,name=priority,proto3,enum=com.coralogixapis.alerts.v3.AlertDefPriority" json:"priority,omitempty"`
	KeyValues  map[string]string      `protobuf:"bytes,4,rep,name=key_values,json=keyValues,proto3" json:"key_values,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Types that are valid to be assigned to Details:
	//
	//	*Group_LogsImmediate
	//	*Group_LogsThreshold
	//	*Group_MetricThreshold
	Details       isGroup_Details `protobuf_oneof:"details"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Group) Reset() {
	*x = Group{}
	mi := &file_com_coralogixapis_alerts_v3_notification_notification_event_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_notification_notification_event_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_notification_notification_event_proto_rawDescGZIP(), []int{1}
}

func (x *Group) GetStatus() AlertStatus {
	if x != nil {
		return x.Status
	}
	return AlertStatus_ALERT_STATUS_RESOLVED_OR_UNSPECIFIED
}

func (x *Group) GetSuppressed() *wrapperspb.BoolValue {
	if x != nil {
		return x.Suppressed
	}
	return nil
}

func (x *Group) GetPriority() AlertDefPriority {
	if x != nil {
		return x.Priority
	}
	return AlertDefPriority_ALERT_DEF_PRIORITY_P5_OR_UNSPECIFIED
}

func (x *Group) GetKeyValues() map[string]string {
	if x != nil {
		return x.KeyValues
	}
	return nil
}

func (x *Group) GetDetails() isGroup_Details {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *Group) GetLogsImmediate() *LogsImmediateNotification {
	if x != nil {
		if x, ok := x.Details.(*Group_LogsImmediate); ok {
			return x.LogsImmediate
		}
	}
	return nil
}

func (x *Group) GetLogsThreshold() *LogsThresholdNotification {
	if x != nil {
		if x, ok := x.Details.(*Group_LogsThreshold); ok {
			return x.LogsThreshold
		}
	}
	return nil
}

func (x *Group) GetMetricThreshold() *MetricThresholdNotification {
	if x != nil {
		if x, ok := x.Details.(*Group_MetricThreshold); ok {
			return x.MetricThreshold
		}
	}
	return nil
}

type isGroup_Details interface {
	isGroup_Details()
}

type Group_LogsImmediate struct {
	LogsImmediate *LogsImmediateNotification `protobuf:"bytes,100,opt,name=logs_immediate,json=logsImmediate,proto3,oneof"`
}

type Group_LogsThreshold struct {
	LogsThreshold *LogsThresholdNotification `protobuf:"bytes,101,opt,name=logs_threshold,json=logsThreshold,proto3,oneof"`
}

type Group_MetricThreshold struct {
	MetricThreshold *MetricThresholdNotification `protobuf:"bytes,102,opt,name=metric_threshold,json=metricThreshold,proto3,oneof"`
}

func (*Group_LogsImmediate) isGroup_Details() {}

func (*Group_LogsThreshold) isGroup_Details() {}

func (*Group_MetricThreshold) isGroup_Details() {}

type Attachments struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	LogExample    *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=logExample,proto3" json:"logExample,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Attachments) Reset() {
	*x = Attachments{}
	mi := &file_com_coralogixapis_alerts_v3_notification_notification_event_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Attachments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attachments) ProtoMessage() {}

func (x *Attachments) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_notification_notification_event_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attachments.ProtoReflect.Descriptor instead.
func (*Attachments) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_notification_notification_event_proto_rawDescGZIP(), []int{2}
}

func (x *Attachments) GetLogExample() *wrapperspb.StringValue {
	if x != nil {
		return x.LogExample
	}
	return nil
}

var File_com_coralogixapis_alerts_v3_notification_notification_event_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_notification_notification_event_proto_rawDesc = []byte{
	0x0a, 0x41, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x33, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x70, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4a, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x69, 0x6d, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4a, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4c, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76,
	0x33, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x5f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc6, 0x03, 0x0a, 0x16, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x02, 0x69, 0x64, 0x12, 0x40, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x2e, 0x76, 0x33, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4a, 0x0a, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x57, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76,
	0x33, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x1a, 0x5d, 0x0a, 0x0b,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x38, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x94, 0x05, 0x0a, 0x05,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x40, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x2e, 0x76, 0x33, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3a, 0x0a, 0x0a, 0x73, 0x75, 0x70, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x73, 0x75, 0x70, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x65, 0x64, 0x12, 0x49, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x2e, 0x76, 0x33, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x44, 0x65, 0x66, 0x50, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x50,
	0x0a, 0x0a, 0x6b, 0x65, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33,
	0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x6b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x12, 0x5f, 0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x69, 0x6d, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x74, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x49, 0x6d, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x00, 0x52, 0x0d, 0x6c, 0x6f, 0x67, 0x73, 0x49, 0x6d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74,
	0x65, 0x12, 0x5f, 0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68,
	0x6f, 0x6c, 0x64, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x54, 0x68, 0x72, 0x65,
	0x73, 0x68, 0x6f, 0x6c, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x48, 0x00, 0x52, 0x0d, 0x6c, 0x6f, 0x67, 0x73, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f,
	0x6c, 0x64, 0x12, 0x65, 0x0a, 0x10, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x74, 0x68, 0x72,
	0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x1a, 0x3c, 0x0a, 0x0e, 0x4b, 0x65, 0x79,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x22, 0x4b, 0x0a, 0x0b, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x3c, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_notification_notification_event_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_notification_notification_event_proto_rawDescData = file_com_coralogixapis_alerts_v3_notification_notification_event_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_notification_notification_event_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_notification_notification_event_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_notification_notification_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_notification_notification_event_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_notification_notification_event_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_notification_notification_event_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_com_coralogixapis_alerts_v3_notification_notification_event_proto_goTypes = []any{
	(*AlertNotificationEvent)(nil),      // 0: com.coralogixapis.alerts.v3.AlertNotificationEvent
	(*Group)(nil),                       // 1: com.coralogixapis.alerts.v3.Group
	(*Attachments)(nil),                 // 2: com.coralogixapis.alerts.v3.Attachments
	nil,                                 // 3: com.coralogixapis.alerts.v3.AlertNotificationEvent.GroupsEntry
	nil,                                 // 4: com.coralogixapis.alerts.v3.Group.KeyValuesEntry
	(*timestamppb.Timestamp)(nil),       // 5: google.protobuf.Timestamp
	(*wrapperspb.StringValue)(nil),      // 6: google.protobuf.StringValue
	(AlertStatus)(0),                    // 7: com.coralogixapis.alerts.v3.AlertStatus
	(*wrapperspb.BoolValue)(nil),        // 8: google.protobuf.BoolValue
	(AlertDefPriority)(0),               // 9: com.coralogixapis.alerts.v3.AlertDefPriority
	(*LogsImmediateNotification)(nil),   // 10: com.coralogixapis.alerts.v3.LogsImmediateNotification
	(*LogsThresholdNotification)(nil),   // 11: com.coralogixapis.alerts.v3.LogsThresholdNotification
	(*MetricThresholdNotification)(nil), // 12: com.coralogixapis.alerts.v3.MetricThresholdNotification
}
var file_com_coralogixapis_alerts_v3_notification_notification_event_proto_depIdxs = []int32{
	5,  // 0: com.coralogixapis.alerts.v3.AlertNotificationEvent.timestamp:type_name -> google.protobuf.Timestamp
	6,  // 1: com.coralogixapis.alerts.v3.AlertNotificationEvent.id:type_name -> google.protobuf.StringValue
	7,  // 2: com.coralogixapis.alerts.v3.AlertNotificationEvent.status:type_name -> com.coralogixapis.alerts.v3.AlertStatus
	2,  // 3: com.coralogixapis.alerts.v3.AlertNotificationEvent.attachments:type_name -> com.coralogixapis.alerts.v3.Attachments
	3,  // 4: com.coralogixapis.alerts.v3.AlertNotificationEvent.groups:type_name -> com.coralogixapis.alerts.v3.AlertNotificationEvent.GroupsEntry
	7,  // 5: com.coralogixapis.alerts.v3.Group.status:type_name -> com.coralogixapis.alerts.v3.AlertStatus
	8,  // 6: com.coralogixapis.alerts.v3.Group.suppressed:type_name -> google.protobuf.BoolValue
	9,  // 7: com.coralogixapis.alerts.v3.Group.priority:type_name -> com.coralogixapis.alerts.v3.AlertDefPriority
	4,  // 8: com.coralogixapis.alerts.v3.Group.key_values:type_name -> com.coralogixapis.alerts.v3.Group.KeyValuesEntry
	10, // 9: com.coralogixapis.alerts.v3.Group.logs_immediate:type_name -> com.coralogixapis.alerts.v3.LogsImmediateNotification
	11, // 10: com.coralogixapis.alerts.v3.Group.logs_threshold:type_name -> com.coralogixapis.alerts.v3.LogsThresholdNotification
	12, // 11: com.coralogixapis.alerts.v3.Group.metric_threshold:type_name -> com.coralogixapis.alerts.v3.MetricThresholdNotification
	6,  // 12: com.coralogixapis.alerts.v3.Attachments.logExample:type_name -> google.protobuf.StringValue
	1,  // 13: com.coralogixapis.alerts.v3.AlertNotificationEvent.GroupsEntry.value:type_name -> com.coralogixapis.alerts.v3.Group
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_alerts_v3_notification_notification_event_proto_init() }
func file_com_coralogixapis_alerts_v3_notification_notification_event_proto_init() {
	if File_com_coralogixapis_alerts_v3_notification_notification_event_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_event_alert_event_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_priority_proto_init()
	file_com_coralogixapis_alerts_v3_notification_logs_immediate_notification_proto_init()
	file_com_coralogixapis_alerts_v3_notification_logs_threshold_notification_proto_init()
	file_com_coralogixapis_alerts_v3_notification_metric_threshold_notification_proto_init()
	file_com_coralogixapis_alerts_v3_notification_notification_event_proto_msgTypes[1].OneofWrappers = []any{
		(*Group_LogsImmediate)(nil),
		(*Group_LogsThreshold)(nil),
		(*Group_MetricThreshold)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_notification_notification_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_notification_notification_event_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_notification_notification_event_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_notification_notification_event_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_notification_notification_event_proto = out.File
	file_com_coralogixapis_alerts_v3_notification_notification_event_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_notification_notification_event_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_notification_notification_event_proto_depIdxs = nil
}
