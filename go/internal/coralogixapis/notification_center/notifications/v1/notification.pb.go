// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: com/coralogixapis/notification_center/notifications/v1/notification.proto

package v1

import (
	_ "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/notification_center/common/v1"
	_ "github.com/golang/protobuf/ptypes/struct"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TriggerType int32

const (
	TriggerType_TRIGGER_TYPE_UNSPECIFIED TriggerType = 0
	TriggerType_MANUAL                   TriggerType = 1
	TriggerType_AUTOMATIC                TriggerType = 2
)

// Enum value maps for TriggerType.
var (
	TriggerType_name = map[int32]string{
		0: "TRIGGER_TYPE_UNSPECIFIED",
		1: "MANUAL",
		2: "AUTOMATIC",
	}
	TriggerType_value = map[string]int32{
		"TRIGGER_TYPE_UNSPECIFIED": 0,
		"MANUAL":                   1,
		"AUTOMATIC":                2,
	}
)

func (x TriggerType) Enum() *TriggerType {
	p := new(TriggerType)
	*p = x
	return p
}

func (x TriggerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TriggerType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_notification_center_notifications_v1_notification_proto_enumTypes[0].Descriptor()
}

func (TriggerType) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_notification_center_notifications_v1_notification_proto_enumTypes[0]
}

func (x TriggerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TriggerType.Descriptor instead.
func (TriggerType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDescGZIP(), []int{0}
}

type NotificationTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Target:
	//
	//	*NotificationTarget_Private
	//	*NotificationTarget_Global
	Target isNotificationTarget_Target `protobuf_oneof:"target"`
}

func (x *NotificationTarget) Reset() {
	*x = NotificationTarget{}
	mi := &file_com_coralogixapis_notification_center_notifications_v1_notification_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotificationTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationTarget) ProtoMessage() {}

func (x *NotificationTarget) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_notifications_v1_notification_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationTarget.ProtoReflect.Descriptor instead.
func (*NotificationTarget) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDescGZIP(), []int{0}
}

func (m *NotificationTarget) GetTarget() isNotificationTarget_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (x *NotificationTarget) GetPrivate() *NotificationTarget_PrivateRouterConfig {
	if x, ok := x.GetTarget().(*NotificationTarget_Private); ok {
		return x.Private
	}
	return nil
}

func (x *NotificationTarget) GetGlobal() *NotificationTarget_GlobalRouterConfig {
	if x, ok := x.GetTarget().(*NotificationTarget_Global); ok {
		return x.Global
	}
	return nil
}

type isNotificationTarget_Target interface {
	isNotificationTarget_Target()
}

type NotificationTarget_Private struct {
	Private *NotificationTarget_PrivateRouterConfig `protobuf:"bytes,1,opt,name=private,proto3,oneof"`
}

type NotificationTarget_Global struct {
	Global *NotificationTarget_GlobalRouterConfig `protobuf:"bytes,2,opt,name=global,proto3,oneof"`
}

func (*NotificationTarget_Private) isNotificationTarget_Target() {}

func (*NotificationTarget_Global) isNotificationTarget_Target() {}

type NotificationOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key         string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Types that are assignable to OperationType:
	//
	//	*NotificationOperation_Http
	OperationType isNotificationOperation_OperationType `protobuf_oneof:"operation_type"`
}

func (x *NotificationOperation) Reset() {
	*x = NotificationOperation{}
	mi := &file_com_coralogixapis_notification_center_notifications_v1_notification_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotificationOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationOperation) ProtoMessage() {}

func (x *NotificationOperation) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_notifications_v1_notification_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationOperation.ProtoReflect.Descriptor instead.
func (*NotificationOperation) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDescGZIP(), []int{1}
}

func (x *NotificationOperation) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *NotificationOperation) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (m *NotificationOperation) GetOperationType() isNotificationOperation_OperationType {
	if m != nil {
		return m.OperationType
	}
	return nil
}

func (x *NotificationOperation) GetHttp() *NotificationOperation_HttpOperation {
	if x, ok := x.GetOperationType().(*NotificationOperation_Http); ok {
		return x.Http
	}
	return nil
}

type isNotificationOperation_OperationType interface {
	isNotificationOperation_OperationType()
}

type NotificationOperation_Http struct {
	Http *NotificationOperation_HttpOperation `protobuf:"bytes,100,opt,name=http,proto3,oneof"`
}

func (*NotificationOperation_Http) isNotificationOperation_OperationType() {}

type NotificationAttachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key         string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	ContentType string `protobuf:"bytes,3,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// Types that are assignable to Content:
	//
	//	*NotificationAttachment_Raw_
	//	*NotificationAttachment_Url_
	Content isNotificationAttachment_Content `protobuf_oneof:"content"`
}

func (x *NotificationAttachment) Reset() {
	*x = NotificationAttachment{}
	mi := &file_com_coralogixapis_notification_center_notifications_v1_notification_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotificationAttachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationAttachment) ProtoMessage() {}

func (x *NotificationAttachment) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_notifications_v1_notification_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationAttachment.ProtoReflect.Descriptor instead.
func (*NotificationAttachment) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDescGZIP(), []int{2}
}

func (x *NotificationAttachment) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *NotificationAttachment) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *NotificationAttachment) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (m *NotificationAttachment) GetContent() isNotificationAttachment_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *NotificationAttachment) GetRaw() *NotificationAttachment_Raw {
	if x, ok := x.GetContent().(*NotificationAttachment_Raw_); ok {
		return x.Raw
	}
	return nil
}

func (x *NotificationAttachment) GetUrl() *NotificationAttachment_Url {
	if x, ok := x.GetContent().(*NotificationAttachment_Url_); ok {
		return x.Url
	}
	return nil
}

type isNotificationAttachment_Content interface {
	isNotificationAttachment_Content()
}

type NotificationAttachment_Raw_ struct {
	Raw *NotificationAttachment_Raw `protobuf:"bytes,100,opt,name=raw,proto3,oneof"`
}

type NotificationAttachment_Url_ struct {
	Url *NotificationAttachment_Url `protobuf:"bytes,101,opt,name=url,proto3,oneof"`
}

func (*NotificationAttachment_Raw_) isNotificationAttachment_Content() {}

func (*NotificationAttachment_Url_) isNotificationAttachment_Content() {}

type Trigger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Activator string      `protobuf:"bytes,1,opt,name=activator,proto3" json:"activator,omitempty"`
	Origin    string      `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
	Type      TriggerType `protobuf:"varint,3,opt,name=type,proto3,enum=com.coralogixapis.notification_center.notifications.v1.TriggerType" json:"type,omitempty"`
}

func (x *Trigger) Reset() {
	*x = Trigger{}
	mi := &file_com_coralogixapis_notification_center_notifications_v1_notification_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Trigger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trigger) ProtoMessage() {}

func (x *Trigger) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_notifications_v1_notification_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trigger.ProtoReflect.Descriptor instead.
func (*Trigger) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDescGZIP(), []int{3}
}

func (x *Trigger) GetActivator() string {
	if x != nil {
		return x.Activator
	}
	return ""
}

func (x *Trigger) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *Trigger) GetType() TriggerType {
	if x != nil {
		return x.Type
	}
	return TriggerType_TRIGGER_TYPE_UNSPECIFIED
}

type NotificationTarget_GlobalRouterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouterId string `protobuf:"bytes,1,opt,name=router_id,json=routerId,proto3" json:"router_id,omitempty"`
}

func (x *NotificationTarget_GlobalRouterConfig) Reset() {
	*x = NotificationTarget_GlobalRouterConfig{}
	mi := &file_com_coralogixapis_notification_center_notifications_v1_notification_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotificationTarget_GlobalRouterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationTarget_GlobalRouterConfig) ProtoMessage() {}

func (x *NotificationTarget_GlobalRouterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_notifications_v1_notification_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationTarget_GlobalRouterConfig.ProtoReflect.Descriptor instead.
func (*NotificationTarget_GlobalRouterConfig) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDescGZIP(), []int{0, 0}
}

func (x *NotificationTarget_GlobalRouterConfig) GetRouterId() string {
	if x != nil {
		return x.RouterId
	}
	return ""
}

type NotificationTarget_PrivateRouterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouterId string `protobuf:"bytes,1,opt,name=router_id,json=routerId,proto3" json:"router_id,omitempty"`
}

func (x *NotificationTarget_PrivateRouterConfig) Reset() {
	*x = NotificationTarget_PrivateRouterConfig{}
	mi := &file_com_coralogixapis_notification_center_notifications_v1_notification_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotificationTarget_PrivateRouterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationTarget_PrivateRouterConfig) ProtoMessage() {}

func (x *NotificationTarget_PrivateRouterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_notifications_v1_notification_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationTarget_PrivateRouterConfig.ProtoReflect.Descriptor instead.
func (*NotificationTarget_PrivateRouterConfig) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDescGZIP(), []int{0, 1}
}

func (x *NotificationTarget_PrivateRouterConfig) GetRouterId() string {
	if x != nil {
		return x.RouterId
	}
	return ""
}

type NotificationOperation_HttpOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url     string            `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Method  string            `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Headers map[string]string `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *NotificationOperation_HttpOperation) Reset() {
	*x = NotificationOperation_HttpOperation{}
	mi := &file_com_coralogixapis_notification_center_notifications_v1_notification_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotificationOperation_HttpOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationOperation_HttpOperation) ProtoMessage() {}

func (x *NotificationOperation_HttpOperation) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_notifications_v1_notification_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationOperation_HttpOperation.ProtoReflect.Descriptor instead.
func (*NotificationOperation_HttpOperation) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDescGZIP(), []int{1, 0}
}

func (x *NotificationOperation_HttpOperation) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *NotificationOperation_HttpOperation) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *NotificationOperation_HttpOperation) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

type NotificationAttachment_Raw struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bytes []byte `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
}

func (x *NotificationAttachment_Raw) Reset() {
	*x = NotificationAttachment_Raw{}
	mi := &file_com_coralogixapis_notification_center_notifications_v1_notification_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotificationAttachment_Raw) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationAttachment_Raw) ProtoMessage() {}

func (x *NotificationAttachment_Raw) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_notifications_v1_notification_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationAttachment_Raw.ProtoReflect.Descriptor instead.
func (*NotificationAttachment_Raw) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDescGZIP(), []int{2, 0}
}

func (x *NotificationAttachment_Raw) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

type NotificationAttachment_Url struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url           string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	ContentLength uint32 `protobuf:"varint,2,opt,name=content_length,json=contentLength,proto3" json:"content_length,omitempty"`
}

func (x *NotificationAttachment_Url) Reset() {
	*x = NotificationAttachment_Url{}
	mi := &file_com_coralogixapis_notification_center_notifications_v1_notification_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotificationAttachment_Url) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationAttachment_Url) ProtoMessage() {}

func (x *NotificationAttachment_Url) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_notifications_v1_notification_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationAttachment_Url.ProtoReflect.Descriptor instead.
func (*NotificationAttachment_Url) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDescGZIP(), []int{2, 1}
}

func (x *NotificationAttachment_Url) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *NotificationAttachment_Url) GetContentLength() uint32 {
	if x != nil {
		return x.ContentLength
	}
	return 0
}

var File_com_coralogixapis_notification_center_notifications_v1_notification_proto protoreflect.FileDescriptor

var file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDesc = []byte{
	0x0a, 0x49, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x36, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x43, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x02, 0x0a, 0x12, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x7a, 0x0a,
	0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x5e,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00,
	0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x77, 0x0a, 0x06, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x5d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x06, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x1a, 0x31, 0x0a, 0x12, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x32, 0x0a, 0x13, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x49, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x22, 0xce, 0x03, 0x0a, 0x15, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x71, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x5b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x48, 0x74, 0x74, 0x70, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52,
	0x04, 0x68, 0x74, 0x74, 0x70, 0x1a, 0xfa, 0x01, 0x0a, 0x0d, 0x48, 0x74, 0x74, 0x70, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x82, 0x01, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x68, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x22, 0xa8, 0x03, 0x0a, 0x16, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x66, 0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x64,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x52, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x61, 0x77, 0x48, 0x00, 0x52, 0x03, 0x72, 0x61, 0x77, 0x12,
	0x66, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x52, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x72, 0x6c,
	0x48, 0x00, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x1a, 0x1b, 0x0a, 0x03, 0x52, 0x61, 0x77, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x1a, 0x3e, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x25, 0x0a,
	0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22,
	0x98, 0x01, 0x0a, 0x07, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x12, 0x57, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x43, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x2a, 0x46, 0x0a, 0x0b, 0x54, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x54, 0x52, 0x49,
	0x47, 0x47, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x41, 0x4e, 0x55, 0x41,
	0x4c, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x55, 0x54, 0x4f, 0x4d, 0x41, 0x54, 0x49, 0x43,
	0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDescOnce sync.Once
	file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDescData = file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDesc
)

func file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDescData)
	})
	return file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDescData
}

var file_com_coralogixapis_notification_center_notifications_v1_notification_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_notification_center_notifications_v1_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_com_coralogixapis_notification_center_notifications_v1_notification_proto_goTypes = []any{
	(TriggerType)(0),                               // 0: com.coralogixapis.notification_center.notifications.v1.TriggerType
	(*NotificationTarget)(nil),                     // 1: com.coralogixapis.notification_center.notifications.v1.NotificationTarget
	(*NotificationOperation)(nil),                  // 2: com.coralogixapis.notification_center.notifications.v1.NotificationOperation
	(*NotificationAttachment)(nil),                 // 3: com.coralogixapis.notification_center.notifications.v1.NotificationAttachment
	(*Trigger)(nil),                                // 4: com.coralogixapis.notification_center.notifications.v1.Trigger
	(*NotificationTarget_GlobalRouterConfig)(nil),  // 5: com.coralogixapis.notification_center.notifications.v1.NotificationTarget.GlobalRouterConfig
	(*NotificationTarget_PrivateRouterConfig)(nil), // 6: com.coralogixapis.notification_center.notifications.v1.NotificationTarget.PrivateRouterConfig
	(*NotificationOperation_HttpOperation)(nil),    // 7: com.coralogixapis.notification_center.notifications.v1.NotificationOperation.HttpOperation
	nil,                                // 8: com.coralogixapis.notification_center.notifications.v1.NotificationOperation.HttpOperation.HeadersEntry
	(*NotificationAttachment_Raw)(nil), // 9: com.coralogixapis.notification_center.notifications.v1.NotificationAttachment.Raw
	(*NotificationAttachment_Url)(nil), // 10: com.coralogixapis.notification_center.notifications.v1.NotificationAttachment.Url
}
var file_com_coralogixapis_notification_center_notifications_v1_notification_proto_depIdxs = []int32{
	6,  // 0: com.coralogixapis.notification_center.notifications.v1.NotificationTarget.private:type_name -> com.coralogixapis.notification_center.notifications.v1.NotificationTarget.PrivateRouterConfig
	5,  // 1: com.coralogixapis.notification_center.notifications.v1.NotificationTarget.global:type_name -> com.coralogixapis.notification_center.notifications.v1.NotificationTarget.GlobalRouterConfig
	7,  // 2: com.coralogixapis.notification_center.notifications.v1.NotificationOperation.http:type_name -> com.coralogixapis.notification_center.notifications.v1.NotificationOperation.HttpOperation
	9,  // 3: com.coralogixapis.notification_center.notifications.v1.NotificationAttachment.raw:type_name -> com.coralogixapis.notification_center.notifications.v1.NotificationAttachment.Raw
	10, // 4: com.coralogixapis.notification_center.notifications.v1.NotificationAttachment.url:type_name -> com.coralogixapis.notification_center.notifications.v1.NotificationAttachment.Url
	0,  // 5: com.coralogixapis.notification_center.notifications.v1.Trigger.type:type_name -> com.coralogixapis.notification_center.notifications.v1.TriggerType
	8,  // 6: com.coralogixapis.notification_center.notifications.v1.NotificationOperation.HttpOperation.headers:type_name -> com.coralogixapis.notification_center.notifications.v1.NotificationOperation.HttpOperation.HeadersEntry
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_notification_center_notifications_v1_notification_proto_init() }
func file_com_coralogixapis_notification_center_notifications_v1_notification_proto_init() {
	if File_com_coralogixapis_notification_center_notifications_v1_notification_proto != nil {
		return
	}
	file_com_coralogixapis_notification_center_notifications_v1_notification_proto_msgTypes[0].OneofWrappers = []any{
		(*NotificationTarget_Private)(nil),
		(*NotificationTarget_Global)(nil),
	}
	file_com_coralogixapis_notification_center_notifications_v1_notification_proto_msgTypes[1].OneofWrappers = []any{
		(*NotificationOperation_Http)(nil),
	}
	file_com_coralogixapis_notification_center_notifications_v1_notification_proto_msgTypes[2].OneofWrappers = []any{
		(*NotificationAttachment_Raw_)(nil),
		(*NotificationAttachment_Url_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_notification_center_notifications_v1_notification_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_notification_center_notifications_v1_notification_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_notification_center_notifications_v1_notification_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_notification_center_notifications_v1_notification_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_notification_center_notifications_v1_notification_proto = out.File
	file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDesc = nil
	file_com_coralogixapis_notification_center_notifications_v1_notification_proto_goTypes = nil
	file_com_coralogixapis_notification_center_notifications_v1_notification_proto_depIdxs = nil
}