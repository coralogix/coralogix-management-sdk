// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: com/coralogixapis/alerts/v3/alert_def_notification_group.proto

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

type NotifyOn int32

const (
	NotifyOn_NOTIFY_ON_TRIGGERED_ONLY_UNSPECIFIED NotifyOn = 0
	NotifyOn_NOTIFY_ON_TRIGGERED_AND_RESOLVED     NotifyOn = 1
)

// Enum value maps for NotifyOn.
var (
	NotifyOn_name = map[int32]string{
		0: "NOTIFY_ON_TRIGGERED_ONLY_UNSPECIFIED",
		1: "NOTIFY_ON_TRIGGERED_AND_RESOLVED",
	}
	NotifyOn_value = map[string]int32{
		"NOTIFY_ON_TRIGGERED_ONLY_UNSPECIFIED": 0,
		"NOTIFY_ON_TRIGGERED_AND_RESOLVED":     1,
	}
)

func (x NotifyOn) Enum() *NotifyOn {
	p := new(NotifyOn)
	*p = x
	return p
}

func (x NotifyOn) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotifyOn) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_enumTypes[0].Descriptor()
}

func (NotifyOn) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_enumTypes[0]
}

func (x NotifyOn) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotifyOn.Descriptor instead.
func (NotifyOn) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDescGZIP(), []int{0}
}

type AlertDefIncidentSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to RetriggeringPeriod:
	//
	//	*AlertDefIncidentSettings_Minutes
	RetriggeringPeriod isAlertDefIncidentSettings_RetriggeringPeriod `protobuf_oneof:"retriggering_period"`
	NotifyOn           NotifyOn                                      `protobuf:"varint,2,opt,name=notify_on,json=notifyOn,proto3,enum=com.coralogixapis.alerts.v3.NotifyOn" json:"notify_on,omitempty"`
}

func (x *AlertDefIncidentSettings) Reset() {
	*x = AlertDefIncidentSettings{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlertDefIncidentSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertDefIncidentSettings) ProtoMessage() {}

func (x *AlertDefIncidentSettings) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertDefIncidentSettings.ProtoReflect.Descriptor instead.
func (*AlertDefIncidentSettings) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDescGZIP(), []int{0}
}

func (m *AlertDefIncidentSettings) GetRetriggeringPeriod() isAlertDefIncidentSettings_RetriggeringPeriod {
	if m != nil {
		return m.RetriggeringPeriod
	}
	return nil
}

func (x *AlertDefIncidentSettings) GetMinutes() *wrapperspb.UInt32Value {
	if x, ok := x.GetRetriggeringPeriod().(*AlertDefIncidentSettings_Minutes); ok {
		return x.Minutes
	}
	return nil
}

func (x *AlertDefIncidentSettings) GetNotifyOn() NotifyOn {
	if x != nil {
		return x.NotifyOn
	}
	return NotifyOn_NOTIFY_ON_TRIGGERED_ONLY_UNSPECIFIED
}

type isAlertDefIncidentSettings_RetriggeringPeriod interface {
	isAlertDefIncidentSettings_RetriggeringPeriod()
}

type AlertDefIncidentSettings_Minutes struct {
	Minutes *wrapperspb.UInt32Value `protobuf:"bytes,100,opt,name=minutes,proto3,oneof"`
}

func (*AlertDefIncidentSettings_Minutes) isAlertDefIncidentSettings_RetriggeringPeriod() {}

type AlertDefNotificationGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupByFields []*wrapperspb.StringValue `protobuf:"bytes,1,rep,name=group_by_fields,json=groupByFields,proto3" json:"group_by_fields,omitempty"`
	// Types that are assignable to Targets:
	//
	//	*AlertDefNotificationGroup_Advanced
	//	*AlertDefNotificationGroup_Simple
	Targets isAlertDefNotificationGroup_Targets `protobuf_oneof:"targets"`
}

func (x *AlertDefNotificationGroup) Reset() {
	*x = AlertDefNotificationGroup{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlertDefNotificationGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertDefNotificationGroup) ProtoMessage() {}

func (x *AlertDefNotificationGroup) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertDefNotificationGroup.ProtoReflect.Descriptor instead.
func (*AlertDefNotificationGroup) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDescGZIP(), []int{1}
}

func (x *AlertDefNotificationGroup) GetGroupByFields() []*wrapperspb.StringValue {
	if x != nil {
		return x.GroupByFields
	}
	return nil
}

func (m *AlertDefNotificationGroup) GetTargets() isAlertDefNotificationGroup_Targets {
	if m != nil {
		return m.Targets
	}
	return nil
}

func (x *AlertDefNotificationGroup) GetAdvanced() *AlertDefAdvancedTargets {
	if x, ok := x.GetTargets().(*AlertDefNotificationGroup_Advanced); ok {
		return x.Advanced
	}
	return nil
}

func (x *AlertDefNotificationGroup) GetSimple() *AlertDefTargetSimple {
	if x, ok := x.GetTargets().(*AlertDefNotificationGroup_Simple); ok {
		return x.Simple
	}
	return nil
}

type isAlertDefNotificationGroup_Targets interface {
	isAlertDefNotificationGroup_Targets()
}

type AlertDefNotificationGroup_Advanced struct {
	Advanced *AlertDefAdvancedTargets `protobuf:"bytes,2,opt,name=advanced,proto3,oneof"`
}

type AlertDefNotificationGroup_Simple struct {
	Simple *AlertDefTargetSimple `protobuf:"bytes,3,opt,name=simple,proto3,oneof"`
}

func (*AlertDefNotificationGroup_Advanced) isAlertDefNotificationGroup_Targets() {}

func (*AlertDefNotificationGroup_Simple) isAlertDefNotificationGroup_Targets() {}

type AlertDefTargetSimple struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Integrations []*IntegrationType `protobuf:"bytes,1,rep,name=integrations,proto3" json:"integrations,omitempty"`
}

func (x *AlertDefTargetSimple) Reset() {
	*x = AlertDefTargetSimple{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlertDefTargetSimple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertDefTargetSimple) ProtoMessage() {}

func (x *AlertDefTargetSimple) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertDefTargetSimple.ProtoReflect.Descriptor instead.
func (*AlertDefTargetSimple) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDescGZIP(), []int{2}
}

func (x *AlertDefTargetSimple) GetIntegrations() []*IntegrationType {
	if x != nil {
		return x.Integrations
	}
	return nil
}

type AlertDefAdvancedTargets struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdvancedTargetsSettings []*AlertDefAdvancedTargetSettings `protobuf:"bytes,1,rep,name=advanced_targets_settings,json=advancedTargetsSettings,proto3" json:"advanced_targets_settings,omitempty"`
}

func (x *AlertDefAdvancedTargets) Reset() {
	*x = AlertDefAdvancedTargets{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlertDefAdvancedTargets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertDefAdvancedTargets) ProtoMessage() {}

func (x *AlertDefAdvancedTargets) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertDefAdvancedTargets.ProtoReflect.Descriptor instead.
func (*AlertDefAdvancedTargets) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDescGZIP(), []int{3}
}

func (x *AlertDefAdvancedTargets) GetAdvancedTargetsSettings() []*AlertDefAdvancedTargetSettings {
	if x != nil {
		return x.AdvancedTargetsSettings
	}
	return nil
}

type AlertDefAdvancedTargetSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to RetriggeringPeriod:
	//
	//	*AlertDefAdvancedTargetSettings_Minutes
	RetriggeringPeriod isAlertDefAdvancedTargetSettings_RetriggeringPeriod `protobuf_oneof:"retriggering_period"`
	NotifyOn           *NotifyOn                                           `protobuf:"varint,1,opt,name=notify_on,json=notifyOn,proto3,enum=com.coralogixapis.alerts.v3.NotifyOn,oneof" json:"notify_on,omitempty"`
	Integration        *IntegrationType                                    `protobuf:"bytes,2,opt,name=integration,proto3" json:"integration,omitempty"`
}

func (x *AlertDefAdvancedTargetSettings) Reset() {
	*x = AlertDefAdvancedTargetSettings{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlertDefAdvancedTargetSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertDefAdvancedTargetSettings) ProtoMessage() {}

func (x *AlertDefAdvancedTargetSettings) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertDefAdvancedTargetSettings.ProtoReflect.Descriptor instead.
func (*AlertDefAdvancedTargetSettings) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDescGZIP(), []int{4}
}

func (m *AlertDefAdvancedTargetSettings) GetRetriggeringPeriod() isAlertDefAdvancedTargetSettings_RetriggeringPeriod {
	if m != nil {
		return m.RetriggeringPeriod
	}
	return nil
}

func (x *AlertDefAdvancedTargetSettings) GetMinutes() *wrapperspb.UInt32Value {
	if x, ok := x.GetRetriggeringPeriod().(*AlertDefAdvancedTargetSettings_Minutes); ok {
		return x.Minutes
	}
	return nil
}

func (x *AlertDefAdvancedTargetSettings) GetNotifyOn() NotifyOn {
	if x != nil && x.NotifyOn != nil {
		return *x.NotifyOn
	}
	return NotifyOn_NOTIFY_ON_TRIGGERED_ONLY_UNSPECIFIED
}

func (x *AlertDefAdvancedTargetSettings) GetIntegration() *IntegrationType {
	if x != nil {
		return x.Integration
	}
	return nil
}

type isAlertDefAdvancedTargetSettings_RetriggeringPeriod interface {
	isAlertDefAdvancedTargetSettings_RetriggeringPeriod()
}

type AlertDefAdvancedTargetSettings_Minutes struct {
	Minutes *wrapperspb.UInt32Value `protobuf:"bytes,100,opt,name=minutes,proto3,oneof"`
}

func (*AlertDefAdvancedTargetSettings_Minutes) isAlertDefAdvancedTargetSettings_RetriggeringPeriod() {
}

type IntegrationType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to IntegrationType:
	//
	//	*IntegrationType_IntegrationId
	//	*IntegrationType_Recipients
	IntegrationType isIntegrationType_IntegrationType `protobuf_oneof:"integration_type"`
}

func (x *IntegrationType) Reset() {
	*x = IntegrationType{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IntegrationType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegrationType) ProtoMessage() {}

func (x *IntegrationType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegrationType.ProtoReflect.Descriptor instead.
func (*IntegrationType) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDescGZIP(), []int{5}
}

func (m *IntegrationType) GetIntegrationType() isIntegrationType_IntegrationType {
	if m != nil {
		return m.IntegrationType
	}
	return nil
}

func (x *IntegrationType) GetIntegrationId() *wrapperspb.UInt32Value {
	if x, ok := x.GetIntegrationType().(*IntegrationType_IntegrationId); ok {
		return x.IntegrationId
	}
	return nil
}

func (x *IntegrationType) GetRecipients() *Recipients {
	if x, ok := x.GetIntegrationType().(*IntegrationType_Recipients); ok {
		return x.Recipients
	}
	return nil
}

type isIntegrationType_IntegrationType interface {
	isIntegrationType_IntegrationType()
}

type IntegrationType_IntegrationId struct {
	IntegrationId *wrapperspb.UInt32Value `protobuf:"bytes,2,opt,name=integration_id,json=integrationId,proto3,oneof"`
}

type IntegrationType_Recipients struct {
	Recipients *Recipients `protobuf:"bytes,3,opt,name=recipients,proto3,oneof"`
}

func (*IntegrationType_IntegrationId) isIntegrationType_IntegrationType() {}

func (*IntegrationType_Recipients) isIntegrationType_IntegrationType() {}

type Recipients struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Emails []*wrapperspb.StringValue `protobuf:"bytes,1,rep,name=emails,proto3" json:"emails,omitempty"`
}

func (x *Recipients) Reset() {
	*x = Recipients{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Recipients) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recipients) ProtoMessage() {}

func (x *Recipients) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recipients.ProtoReflect.Descriptor instead.
func (*Recipients) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDescGZIP(), []int{6}
}

func (x *Recipients) GetEmails() []*wrapperspb.StringValue {
	if x != nil {
		return x.Emails
	}
	return nil
}

var File_com_coralogixapis_alerts_v3_alert_def_notification_group_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x01,
	0x0a, 0x18, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x44, 0x65, 0x66, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x38, 0x0a, 0x07, 0x6d, 0x69,
	0x6e, 0x75, 0x74, 0x65, 0x73, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x69, 0x6e,
	0x75, 0x74, 0x65, 0x73, 0x12, 0x42, 0x0a, 0x09, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x6e, 0x52, 0x08,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x6e, 0x42, 0x15, 0x0a, 0x13, 0x72, 0x65, 0x74, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x22,
	0x8d, 0x02, 0x0a, 0x19, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x44, 0x65, 0x66, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x44, 0x0a,
	0x0f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x62, 0x79, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x79, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x12, 0x52, 0x0a, 0x08, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x2e, 0x76, 0x33, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x44, 0x65, 0x66, 0x41, 0x64, 0x76, 0x61,
	0x6e, 0x63, 0x65, 0x64, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x48, 0x00, 0x52, 0x08, 0x61,
	0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x12, 0x4b, 0x0a, 0x06, 0x73, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x44, 0x65, 0x66, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x06, 0x73, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x22,
	0x68, 0x0a, 0x14, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x44, 0x65, 0x66, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x50, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x49, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x92, 0x01, 0x0a, 0x17, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x44, 0x65, 0x66, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x77, 0x0a, 0x19, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65,
	0x64, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x44, 0x65, 0x66, 0x41,
	0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x17, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x98,
	0x02, 0x0a, 0x1e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x44, 0x65, 0x66, 0x41, 0x64, 0x76, 0x61, 0x6e,
	0x63, 0x65, 0x64, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x12, 0x38, 0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x18, 0x64, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x48, 0x00, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x12, 0x47, 0x0a, 0x09, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x4f, 0x6e, 0x48, 0x01, 0x52, 0x08, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f,
	0x6e, 0x88, 0x01, 0x01, 0x12, 0x4e, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x15, 0x0a, 0x13, 0x72, 0x65, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x6f, 0x6e, 0x22, 0xb7, 0x01, 0x0a, 0x0f, 0x49, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x45, 0x0a,
	0x0e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x49, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x48, 0x00, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x42,
	0x12, 0x0a, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x42, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x34, 0x0a, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x2a, 0x5a, 0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x4f, 0x6e, 0x12, 0x28, 0x0a, 0x24, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x4f, 0x4e,
	0x5f, 0x54, 0x52, 0x49, 0x47, 0x47, 0x45, 0x52, 0x45, 0x44, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x24, 0x0a,
	0x20, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x4f, 0x4e, 0x5f, 0x54, 0x52, 0x49, 0x47, 0x47,
	0x45, 0x52, 0x45, 0x44, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x56, 0x45,
	0x44, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDescData = file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_goTypes = []any{
	(NotifyOn)(0),                          // 0: com.coralogixapis.alerts.v3.NotifyOn
	(*AlertDefIncidentSettings)(nil),       // 1: com.coralogixapis.alerts.v3.AlertDefIncidentSettings
	(*AlertDefNotificationGroup)(nil),      // 2: com.coralogixapis.alerts.v3.AlertDefNotificationGroup
	(*AlertDefTargetSimple)(nil),           // 3: com.coralogixapis.alerts.v3.AlertDefTargetSimple
	(*AlertDefAdvancedTargets)(nil),        // 4: com.coralogixapis.alerts.v3.AlertDefAdvancedTargets
	(*AlertDefAdvancedTargetSettings)(nil), // 5: com.coralogixapis.alerts.v3.AlertDefAdvancedTargetSettings
	(*IntegrationType)(nil),                // 6: com.coralogixapis.alerts.v3.IntegrationType
	(*Recipients)(nil),                     // 7: com.coralogixapis.alerts.v3.Recipients
	(*wrapperspb.UInt32Value)(nil),         // 8: google.protobuf.UInt32Value
	(*wrapperspb.StringValue)(nil),         // 9: google.protobuf.StringValue
}
var file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_depIdxs = []int32{
	8,  // 0: com.coralogixapis.alerts.v3.AlertDefIncidentSettings.minutes:type_name -> google.protobuf.UInt32Value
	0,  // 1: com.coralogixapis.alerts.v3.AlertDefIncidentSettings.notify_on:type_name -> com.coralogixapis.alerts.v3.NotifyOn
	9,  // 2: com.coralogixapis.alerts.v3.AlertDefNotificationGroup.group_by_fields:type_name -> google.protobuf.StringValue
	4,  // 3: com.coralogixapis.alerts.v3.AlertDefNotificationGroup.advanced:type_name -> com.coralogixapis.alerts.v3.AlertDefAdvancedTargets
	3,  // 4: com.coralogixapis.alerts.v3.AlertDefNotificationGroup.simple:type_name -> com.coralogixapis.alerts.v3.AlertDefTargetSimple
	6,  // 5: com.coralogixapis.alerts.v3.AlertDefTargetSimple.integrations:type_name -> com.coralogixapis.alerts.v3.IntegrationType
	5,  // 6: com.coralogixapis.alerts.v3.AlertDefAdvancedTargets.advanced_targets_settings:type_name -> com.coralogixapis.alerts.v3.AlertDefAdvancedTargetSettings
	8,  // 7: com.coralogixapis.alerts.v3.AlertDefAdvancedTargetSettings.minutes:type_name -> google.protobuf.UInt32Value
	0,  // 8: com.coralogixapis.alerts.v3.AlertDefAdvancedTargetSettings.notify_on:type_name -> com.coralogixapis.alerts.v3.NotifyOn
	6,  // 9: com.coralogixapis.alerts.v3.AlertDefAdvancedTargetSettings.integration:type_name -> com.coralogixapis.alerts.v3.IntegrationType
	8,  // 10: com.coralogixapis.alerts.v3.IntegrationType.integration_id:type_name -> google.protobuf.UInt32Value
	7,  // 11: com.coralogixapis.alerts.v3.IntegrationType.recipients:type_name -> com.coralogixapis.alerts.v3.Recipients
	9,  // 12: com.coralogixapis.alerts.v3.Recipients.emails:type_name -> google.protobuf.StringValue
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_init() }
func file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_notification_group_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_msgTypes[0].OneofWrappers = []any{
		(*AlertDefIncidentSettings_Minutes)(nil),
	}
	file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_msgTypes[1].OneofWrappers = []any{
		(*AlertDefNotificationGroup_Advanced)(nil),
		(*AlertDefNotificationGroup_Simple)(nil),
	}
	file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_msgTypes[4].OneofWrappers = []any{
		(*AlertDefAdvancedTargetSettings_Minutes)(nil),
	}
	file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_msgTypes[5].OneofWrappers = []any{
		(*IntegrationType_IntegrationId)(nil),
		(*IntegrationType_Recipients)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_notification_group_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_depIdxs = nil
}
