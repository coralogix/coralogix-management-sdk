// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/notification_center/notifications/v1/notification.proto

package v1

import (
	v1 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/notification_center/common/v1"
	routing "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/notification_center/common/v1/routing"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Target:
	//
	//	*NotificationTarget_Private
	//	*NotificationTarget_Global
	Target        isNotificationTarget_Target `protobuf_oneof:"target"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

func (x *NotificationTarget) GetTarget() isNotificationTarget_Target {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *NotificationTarget) GetPrivate() *NotificationTarget_PrivateRouterConfig {
	if x != nil {
		if x, ok := x.Target.(*NotificationTarget_Private); ok {
			return x.Private
		}
	}
	return nil
}

func (x *NotificationTarget) GetGlobal() *NotificationTarget_GlobalRouterConfig {
	if x != nil {
		if x, ok := x.Target.(*NotificationTarget_Global); ok {
			return x.Global
		}
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
	state       protoimpl.MessageState `protogen:"open.v1"`
	Key         string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	DisplayName string                 `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Types that are valid to be assigned to OperationType:
	//
	//	*NotificationOperation_Http
	OperationType isNotificationOperation_OperationType `protobuf_oneof:"operation_type"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

func (x *NotificationOperation) GetOperationType() isNotificationOperation_OperationType {
	if x != nil {
		return x.OperationType
	}
	return nil
}

func (x *NotificationOperation) GetHttp() *NotificationOperation_HttpOperation {
	if x != nil {
		if x, ok := x.OperationType.(*NotificationOperation_Http); ok {
			return x.Http
		}
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
	state       protoimpl.MessageState `protogen:"open.v1"`
	Key         string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	DisplayName string                 `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	ContentType string                 `protobuf:"bytes,3,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// Types that are valid to be assigned to Content:
	//
	//	*NotificationAttachment_Raw_
	//	*NotificationAttachment_Url_
	Content       isNotificationAttachment_Content `protobuf_oneof:"content"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

func (x *NotificationAttachment) GetContent() isNotificationAttachment_Content {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *NotificationAttachment) GetRaw() *NotificationAttachment_Raw {
	if x != nil {
		if x, ok := x.Content.(*NotificationAttachment_Raw_); ok {
			return x.Raw
		}
	}
	return nil
}

func (x *NotificationAttachment) GetUrl() *NotificationAttachment_Url {
	if x != nil {
		if x, ok := x.Content.(*NotificationAttachment_Url_); ok {
			return x.Url
		}
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// client id is a user-defined string identifying the use case (e.g., "alert-notification-sender")
	ClientId      string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	CxServiceName string `protobuf:"bytes,2,opt,name=cx_service_name,json=cxServiceName,proto3" json:"cx_service_name,omitempty"`
	// Types that are valid to be assigned to TriggerDetail:
	//
	//	*Trigger_ManualTrigger_
	//	*Trigger_AutomaticTrigger_
	TriggerDetail isTrigger_TriggerDetail `protobuf_oneof:"trigger_detail"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

func (x *Trigger) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Trigger) GetCxServiceName() string {
	if x != nil {
		return x.CxServiceName
	}
	return ""
}

func (x *Trigger) GetTriggerDetail() isTrigger_TriggerDetail {
	if x != nil {
		return x.TriggerDetail
	}
	return nil
}

func (x *Trigger) GetManualTrigger() *Trigger_ManualTrigger {
	if x != nil {
		if x, ok := x.TriggerDetail.(*Trigger_ManualTrigger_); ok {
			return x.ManualTrigger
		}
	}
	return nil
}

func (x *Trigger) GetAutomaticTrigger() *Trigger_AutomaticTrigger {
	if x != nil {
		if x, ok := x.TriggerDetail.(*Trigger_AutomaticTrigger_); ok {
			return x.AutomaticTrigger
		}
	}
	return nil
}

type isTrigger_TriggerDetail interface {
	isTrigger_TriggerDetail()
}

type Trigger_ManualTrigger_ struct {
	ManualTrigger *Trigger_ManualTrigger `protobuf:"bytes,100,opt,name=manual_trigger,json=manualTrigger,proto3,oneof"`
}

type Trigger_AutomaticTrigger_ struct {
	AutomaticTrigger *Trigger_AutomaticTrigger `protobuf:"bytes,101,opt,name=automatic_trigger,json=automaticTrigger,proto3,oneof"`
}

func (*Trigger_ManualTrigger_) isTrigger_TriggerDetail() {}

func (*Trigger_AutomaticTrigger_) isTrigger_TriggerDetail() {}

type NotificationTarget_GlobalRouterConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Deprecated: Marked as deprecated in com/coralogixapis/notification_center/notifications/v1/notification.proto.
	DeprecatedIdentifier *routing.GlobalRouterIdentifier `protobuf:"bytes,1,opt,name=deprecated_identifier,json=deprecatedIdentifier,proto3" json:"deprecated_identifier,omitempty"`
	// Deprecated: Marked as deprecated in com/coralogixapis/notification_center/notifications/v1/notification.proto.
	Identifier    *v1.GlobalRouterIdentifier `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Id            string                     `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

// Deprecated: Marked as deprecated in com/coralogixapis/notification_center/notifications/v1/notification.proto.
func (x *NotificationTarget_GlobalRouterConfig) GetDeprecatedIdentifier() *routing.GlobalRouterIdentifier {
	if x != nil {
		return x.DeprecatedIdentifier
	}
	return nil
}

// Deprecated: Marked as deprecated in com/coralogixapis/notification_center/notifications/v1/notification.proto.
func (x *NotificationTarget_GlobalRouterConfig) GetIdentifier() *v1.GlobalRouterIdentifier {
	if x != nil {
		return x.Identifier
	}
	return nil
}

func (x *NotificationTarget_GlobalRouterConfig) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type NotificationTarget_PrivateRouterConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RouterKey     string                 `protobuf:"bytes,1,opt,name=router_key,json=routerKey,proto3" json:"router_key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

func (x *NotificationTarget_PrivateRouterConfig) GetRouterKey() string {
	if x != nil {
		return x.RouterKey
	}
	return ""
}

type NotificationOperation_HttpOperation struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Method        string                 `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Headers       map[string]string      `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Bytes         []byte                 `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	ContentLength uint32                 `protobuf:"varint,2,opt,name=content_length,json=contentLength,proto3" json:"content_length,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

type Trigger_ManualTrigger struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserEmail     string                 `protobuf:"bytes,1,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Trigger_ManualTrigger) Reset() {
	*x = Trigger_ManualTrigger{}
	mi := &file_com_coralogixapis_notification_center_notifications_v1_notification_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Trigger_ManualTrigger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trigger_ManualTrigger) ProtoMessage() {}

func (x *Trigger_ManualTrigger) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_notifications_v1_notification_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trigger_ManualTrigger.ProtoReflect.Descriptor instead.
func (*Trigger_ManualTrigger) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDescGZIP(), []int{3, 0}
}

func (x *Trigger_ManualTrigger) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

type Trigger_AutomaticTrigger struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Trigger_AutomaticTrigger) Reset() {
	*x = Trigger_AutomaticTrigger{}
	mi := &file_com_coralogixapis_notification_center_notifications_v1_notification_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Trigger_AutomaticTrigger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trigger_AutomaticTrigger) ProtoMessage() {}

func (x *Trigger_AutomaticTrigger) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_notifications_v1_notification_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trigger_AutomaticTrigger.ProtoReflect.Descriptor instead.
func (*Trigger_AutomaticTrigger) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDescGZIP(), []int{3, 1}
}

var File_com_coralogixapis_notification_center_notifications_v1_notification_proto protoreflect.FileDescriptor

const file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDesc = "" +
	"\n" +
	"Icom/coralogixapis/notification_center/notifications/v1/notification.proto\x126com.coralogixapis.notification_center.notifications.v1\x1a@com/coralogixapis/notification_center/common/v1/identifier.proto\x1aEcom/coralogixapis/notification_center/common/v1/routing/routing.proto\"\xd3\x04\n" +
	"\x12NotificationTarget\x12z\n" +
	"\aprivate\x18\x01 \x01(\v2^.com.coralogixapis.notification_center.notifications.v1.NotificationTarget.PrivateRouterConfigH\x00R\aprivate\x12w\n" +
	"\x06global\x18\x02 \x01(\v2].com.coralogixapis.notification_center.notifications.v1.NotificationTarget.GlobalRouterConfigH\x00R\x06global\x1a\x87\x02\n" +
	"\x12GlobalRouterConfig\x12~\n" +
	"\x15deprecated_identifier\x18\x01 \x01(\v2E.com.coralogixapis.notification_center.routing.GlobalRouterIdentifierB\x02\x18\x01R\x14deprecatedIdentifier\x12a\n" +
	"\n" +
	"identifier\x18\x02 \x01(\v2=.com.coralogixapis.notification_center.GlobalRouterIdentifierB\x02\x18\x01R\n" +
	"identifier\x12\x0e\n" +
	"\x02id\x18\x03 \x01(\tR\x02id\x1a4\n" +
	"\x13PrivateRouterConfig\x12\x1d\n" +
	"\n" +
	"router_key\x18\x01 \x01(\tR\trouterKeyB\b\n" +
	"\x06target\"\xce\x03\n" +
	"\x15NotificationOperation\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12!\n" +
	"\fdisplay_name\x18\x02 \x01(\tR\vdisplayName\x12q\n" +
	"\x04http\x18d \x01(\v2[.com.coralogixapis.notification_center.notifications.v1.NotificationOperation.HttpOperationH\x00R\x04http\x1a\xfa\x01\n" +
	"\rHttpOperation\x12\x10\n" +
	"\x03url\x18\x01 \x01(\tR\x03url\x12\x16\n" +
	"\x06method\x18\x02 \x01(\tR\x06method\x12\x82\x01\n" +
	"\aheaders\x18\x03 \x03(\v2h.com.coralogixapis.notification_center.notifications.v1.NotificationOperation.HttpOperation.HeadersEntryR\aheaders\x1a:\n" +
	"\fHeadersEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01B\x10\n" +
	"\x0eoperation_type\"\xa8\x03\n" +
	"\x16NotificationAttachment\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12!\n" +
	"\fdisplay_name\x18\x02 \x01(\tR\vdisplayName\x12!\n" +
	"\fcontent_type\x18\x03 \x01(\tR\vcontentType\x12f\n" +
	"\x03raw\x18d \x01(\v2R.com.coralogixapis.notification_center.notifications.v1.NotificationAttachment.RawH\x00R\x03raw\x12f\n" +
	"\x03url\x18e \x01(\v2R.com.coralogixapis.notification_center.notifications.v1.NotificationAttachment.UrlH\x00R\x03url\x1a\x1b\n" +
	"\x03Raw\x12\x14\n" +
	"\x05bytes\x18\x01 \x01(\fR\x05bytes\x1a>\n" +
	"\x03Url\x12\x10\n" +
	"\x03url\x18\x01 \x01(\tR\x03url\x12%\n" +
	"\x0econtent_length\x18\x02 \x01(\rR\rcontentLengthB\t\n" +
	"\acontent\"\x9d\x03\n" +
	"\aTrigger\x12\x1b\n" +
	"\tclient_id\x18\x01 \x01(\tR\bclientId\x12&\n" +
	"\x0fcx_service_name\x18\x02 \x01(\tR\rcxServiceName\x12v\n" +
	"\x0emanual_trigger\x18d \x01(\v2M.com.coralogixapis.notification_center.notifications.v1.Trigger.ManualTriggerH\x00R\rmanualTrigger\x12\x7f\n" +
	"\x11automatic_trigger\x18e \x01(\v2P.com.coralogixapis.notification_center.notifications.v1.Trigger.AutomaticTriggerH\x00R\x10automaticTrigger\x1a.\n" +
	"\rManualTrigger\x12\x1d\n" +
	"\n" +
	"user_email\x18\x01 \x01(\tR\tuserEmail\x1a\x12\n" +
	"\x10AutomaticTriggerB\x10\n" +
	"\x0etrigger_detail*F\n" +
	"\vTriggerType\x12\x1c\n" +
	"\x18TRIGGER_TYPE_UNSPECIFIED\x10\x00\x12\n" +
	"\n" +
	"\x06MANUAL\x10\x01\x12\r\n" +
	"\tAUTOMATIC\x10\x02B8Z6com/coralogixapis/notification_center/notifications/v1b\x06proto3"

var (
	file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDescOnce sync.Once
	file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDescData []byte
)

func file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDesc), len(file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDesc)))
	})
	return file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDescData
}

var file_com_coralogixapis_notification_center_notifications_v1_notification_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_notification_center_notifications_v1_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_com_coralogixapis_notification_center_notifications_v1_notification_proto_goTypes = []any{
	(TriggerType)(0),                               // 0: com.coralogixapis.notification_center.notifications.v1.TriggerType
	(*NotificationTarget)(nil),                     // 1: com.coralogixapis.notification_center.notifications.v1.NotificationTarget
	(*NotificationOperation)(nil),                  // 2: com.coralogixapis.notification_center.notifications.v1.NotificationOperation
	(*NotificationAttachment)(nil),                 // 3: com.coralogixapis.notification_center.notifications.v1.NotificationAttachment
	(*Trigger)(nil),                                // 4: com.coralogixapis.notification_center.notifications.v1.Trigger
	(*NotificationTarget_GlobalRouterConfig)(nil),  // 5: com.coralogixapis.notification_center.notifications.v1.NotificationTarget.GlobalRouterConfig
	(*NotificationTarget_PrivateRouterConfig)(nil), // 6: com.coralogixapis.notification_center.notifications.v1.NotificationTarget.PrivateRouterConfig
	(*NotificationOperation_HttpOperation)(nil),    // 7: com.coralogixapis.notification_center.notifications.v1.NotificationOperation.HttpOperation
	nil,                                    // 8: com.coralogixapis.notification_center.notifications.v1.NotificationOperation.HttpOperation.HeadersEntry
	(*NotificationAttachment_Raw)(nil),     // 9: com.coralogixapis.notification_center.notifications.v1.NotificationAttachment.Raw
	(*NotificationAttachment_Url)(nil),     // 10: com.coralogixapis.notification_center.notifications.v1.NotificationAttachment.Url
	(*Trigger_ManualTrigger)(nil),          // 11: com.coralogixapis.notification_center.notifications.v1.Trigger.ManualTrigger
	(*Trigger_AutomaticTrigger)(nil),       // 12: com.coralogixapis.notification_center.notifications.v1.Trigger.AutomaticTrigger
	(*routing.GlobalRouterIdentifier)(nil), // 13: com.coralogixapis.notification_center.routing.GlobalRouterIdentifier
	(*v1.GlobalRouterIdentifier)(nil),      // 14: com.coralogixapis.notification_center.GlobalRouterIdentifier
}
var file_com_coralogixapis_notification_center_notifications_v1_notification_proto_depIdxs = []int32{
	6,  // 0: com.coralogixapis.notification_center.notifications.v1.NotificationTarget.private:type_name -> com.coralogixapis.notification_center.notifications.v1.NotificationTarget.PrivateRouterConfig
	5,  // 1: com.coralogixapis.notification_center.notifications.v1.NotificationTarget.global:type_name -> com.coralogixapis.notification_center.notifications.v1.NotificationTarget.GlobalRouterConfig
	7,  // 2: com.coralogixapis.notification_center.notifications.v1.NotificationOperation.http:type_name -> com.coralogixapis.notification_center.notifications.v1.NotificationOperation.HttpOperation
	9,  // 3: com.coralogixapis.notification_center.notifications.v1.NotificationAttachment.raw:type_name -> com.coralogixapis.notification_center.notifications.v1.NotificationAttachment.Raw
	10, // 4: com.coralogixapis.notification_center.notifications.v1.NotificationAttachment.url:type_name -> com.coralogixapis.notification_center.notifications.v1.NotificationAttachment.Url
	11, // 5: com.coralogixapis.notification_center.notifications.v1.Trigger.manual_trigger:type_name -> com.coralogixapis.notification_center.notifications.v1.Trigger.ManualTrigger
	12, // 6: com.coralogixapis.notification_center.notifications.v1.Trigger.automatic_trigger:type_name -> com.coralogixapis.notification_center.notifications.v1.Trigger.AutomaticTrigger
	13, // 7: com.coralogixapis.notification_center.notifications.v1.NotificationTarget.GlobalRouterConfig.deprecated_identifier:type_name -> com.coralogixapis.notification_center.routing.GlobalRouterIdentifier
	14, // 8: com.coralogixapis.notification_center.notifications.v1.NotificationTarget.GlobalRouterConfig.identifier:type_name -> com.coralogixapis.notification_center.GlobalRouterIdentifier
	8,  // 9: com.coralogixapis.notification_center.notifications.v1.NotificationOperation.HttpOperation.headers:type_name -> com.coralogixapis.notification_center.notifications.v1.NotificationOperation.HttpOperation.HeadersEntry
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
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
	file_com_coralogixapis_notification_center_notifications_v1_notification_proto_msgTypes[3].OneofWrappers = []any{
		(*Trigger_ManualTrigger_)(nil),
		(*Trigger_AutomaticTrigger_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDesc), len(file_com_coralogixapis_notification_center_notifications_v1_notification_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_notification_center_notifications_v1_notification_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_notification_center_notifications_v1_notification_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_notification_center_notifications_v1_notification_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_notification_center_notifications_v1_notification_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_notification_center_notifications_v1_notification_proto = out.File
	file_com_coralogixapis_notification_center_notifications_v1_notification_proto_goTypes = nil
	file_com_coralogixapis_notification_center_notifications_v1_notification_proto_depIdxs = nil
}
