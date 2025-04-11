// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/alerts/v3/alert_def_notification_group.proto

package v3

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

type AlertDefIncidentSettings struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to RetriggeringPeriod:
	//
	//	*AlertDefIncidentSettings_Minutes
	RetriggeringPeriod isAlertDefIncidentSettings_RetriggeringPeriod `protobuf_oneof:"retriggering_period"`
	NotifyOn           NotifyOn                                      `protobuf:"varint,2,opt,name=notify_on,json=notifyOn,proto3,enum=com.coralogixapis.alerts.v3.NotifyOn" json:"notify_on,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
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

func (x *AlertDefIncidentSettings) GetRetriggeringPeriod() isAlertDefIncidentSettings_RetriggeringPeriod {
	if x != nil {
		return x.RetriggeringPeriod
	}
	return nil
}

func (x *AlertDefIncidentSettings) GetMinutes() *wrapperspb.UInt32Value {
	if x != nil {
		if x, ok := x.RetriggeringPeriod.(*AlertDefIncidentSettings_Minutes); ok {
			return x.Minutes
		}
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
	state         protoimpl.MessageState      `protogen:"open.v1"`
	GroupByKeys   []*wrapperspb.StringValue   `protobuf:"bytes,1,rep,name=group_by_keys,json=groupByKeys,proto3" json:"group_by_keys,omitempty"`
	Webhooks      []*AlertDefWebhooksSettings `protobuf:"bytes,2,rep,name=webhooks,proto3" json:"webhooks,omitempty"`
	Destinations  []*NotificationDestination  `protobuf:"bytes,3,rep,name=destinations,proto3" json:"destinations,omitempty"`
	Router        *NotificationRouter         `protobuf:"bytes,4,opt,name=router,proto3,oneof" json:"router,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

func (x *AlertDefNotificationGroup) GetGroupByKeys() []*wrapperspb.StringValue {
	if x != nil {
		return x.GroupByKeys
	}
	return nil
}

func (x *AlertDefNotificationGroup) GetWebhooks() []*AlertDefWebhooksSettings {
	if x != nil {
		return x.Webhooks
	}
	return nil
}

func (x *AlertDefNotificationGroup) GetDestinations() []*NotificationDestination {
	if x != nil {
		return x.Destinations
	}
	return nil
}

func (x *AlertDefNotificationGroup) GetRouter() *NotificationRouter {
	if x != nil {
		return x.Router
	}
	return nil
}

type AlertDefWebhooksSettings struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to RetriggeringPeriod:
	//
	//	*AlertDefWebhooksSettings_Minutes
	RetriggeringPeriod isAlertDefWebhooksSettings_RetriggeringPeriod `protobuf_oneof:"retriggering_period"`
	NotifyOn           *NotifyOn                                     `protobuf:"varint,1,opt,name=notify_on,json=notifyOn,proto3,enum=com.coralogixapis.alerts.v3.NotifyOn,oneof" json:"notify_on,omitempty"`
	Integration        *IntegrationType                              `protobuf:"bytes,2,opt,name=integration,proto3" json:"integration,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *AlertDefWebhooksSettings) Reset() {
	*x = AlertDefWebhooksSettings{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlertDefWebhooksSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertDefWebhooksSettings) ProtoMessage() {}

func (x *AlertDefWebhooksSettings) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AlertDefWebhooksSettings.ProtoReflect.Descriptor instead.
func (*AlertDefWebhooksSettings) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDescGZIP(), []int{2}
}

func (x *AlertDefWebhooksSettings) GetRetriggeringPeriod() isAlertDefWebhooksSettings_RetriggeringPeriod {
	if x != nil {
		return x.RetriggeringPeriod
	}
	return nil
}

func (x *AlertDefWebhooksSettings) GetMinutes() *wrapperspb.UInt32Value {
	if x != nil {
		if x, ok := x.RetriggeringPeriod.(*AlertDefWebhooksSettings_Minutes); ok {
			return x.Minutes
		}
	}
	return nil
}

func (x *AlertDefWebhooksSettings) GetNotifyOn() NotifyOn {
	if x != nil && x.NotifyOn != nil {
		return *x.NotifyOn
	}
	return NotifyOn_NOTIFY_ON_TRIGGERED_ONLY_UNSPECIFIED
}

func (x *AlertDefWebhooksSettings) GetIntegration() *IntegrationType {
	if x != nil {
		return x.Integration
	}
	return nil
}

type isAlertDefWebhooksSettings_RetriggeringPeriod interface {
	isAlertDefWebhooksSettings_RetriggeringPeriod()
}

type AlertDefWebhooksSettings_Minutes struct {
	Minutes *wrapperspb.UInt32Value `protobuf:"bytes,100,opt,name=minutes,proto3,oneof"`
}

func (*AlertDefWebhooksSettings_Minutes) isAlertDefWebhooksSettings_RetriggeringPeriod() {}

type IntegrationType struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to IntegrationType:
	//
	//	*IntegrationType_IntegrationId
	//	*IntegrationType_Recipients
	IntegrationType isIntegrationType_IntegrationType `protobuf_oneof:"integration_type"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *IntegrationType) Reset() {
	*x = IntegrationType{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IntegrationType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegrationType) ProtoMessage() {}

func (x *IntegrationType) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use IntegrationType.ProtoReflect.Descriptor instead.
func (*IntegrationType) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDescGZIP(), []int{3}
}

func (x *IntegrationType) GetIntegrationType() isIntegrationType_IntegrationType {
	if x != nil {
		return x.IntegrationType
	}
	return nil
}

func (x *IntegrationType) GetIntegrationId() *wrapperspb.UInt32Value {
	if x != nil {
		if x, ok := x.IntegrationType.(*IntegrationType_IntegrationId); ok {
			return x.IntegrationId
		}
	}
	return nil
}

func (x *IntegrationType) GetRecipients() *Recipients {
	if x != nil {
		if x, ok := x.IntegrationType.(*IntegrationType_Recipients); ok {
			return x.Recipients
		}
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
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Emails        []*wrapperspb.StringValue `protobuf:"bytes,1,rep,name=emails,proto3" json:"emails,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Recipients) Reset() {
	*x = Recipients{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Recipients) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recipients) ProtoMessage() {}

func (x *Recipients) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Recipients.ProtoReflect.Descriptor instead.
func (*Recipients) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDescGZIP(), []int{4}
}

func (x *Recipients) GetEmails() []*wrapperspb.StringValue {
	if x != nil {
		return x.Emails
	}
	return nil
}

var File_com_coralogixapis_alerts_v3_alert_def_notification_group_proto protoreflect.FileDescriptor

const file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDesc = "" +
	"\n" +
	">com/coralogixapis/alerts/v3/alert_def_notification_group.proto\x12\x1bcom.coralogixapis.alerts.v3\x1a\x1egoogle/protobuf/wrappers.proto\x1a)com/coralogixapis/alerts/v3/commons.proto\x1aTcom/coralogixapis/alerts/v3/alert_def_type_definition/notification_destination.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\x85\x02\n" +
	"\x18AlertDefIncidentSettings\x12A\n" +
	"\aminutes\x18d \x01(\v2\x1c.google.protobuf.UInt32ValueB\a\x92A\x04J\x0210H\x00R\aminutes\x12B\n" +
	"\tnotify_on\x18\x02 \x01(\x0e2%.com.coralogixapis.alerts.v3.NotifyOnR\bnotifyOn:K\x92AH\n" +
	"F*\"Alert definition incident settings\xd2\x01\x13retriggering_period\xd2\x01\tnotify_onB\x15\n" +
	"\x13retriggering_period\"\xb9\x03\n" +
	"\x19AlertDefNotificationGroup\x12@\n" +
	"\rgroup_by_keys\x18\x01 \x03(\v2\x1c.google.protobuf.StringValueR\vgroupByKeys\x12Q\n" +
	"\bwebhooks\x18\x02 \x03(\v25.com.coralogixapis.alerts.v3.AlertDefWebhooksSettingsR\bwebhooks\x12X\n" +
	"\fdestinations\x18\x03 \x03(\v24.com.coralogixapis.alerts.v3.NotificationDestinationR\fdestinations\x12L\n" +
	"\x06router\x18\x04 \x01(\v2/.com.coralogixapis.alerts.v3.NotificationRouterH\x00R\x06router\x88\x01\x01:T\x92AQ\n" +
	"O*#Alert definition notification group\xd2\x01\rgroup_by_keys\xd2\x01\bwebhooks\xd2\x01\fdestinationsB\t\n" +
	"\a_router\"\xae\x03\n" +
	"\x18AlertDefWebhooksSettings\x12A\n" +
	"\aminutes\x18d \x01(\v2\x1c.google.protobuf.UInt32ValueB\a\x92A\x04J\x0210H\x00R\aminutes\x12G\n" +
	"\tnotify_on\x18\x01 \x01(\x0e2%.com.coralogixapis.alerts.v3.NotifyOnH\x01R\bnotifyOn\x88\x01\x01\x12N\n" +
	"\vintegration\x18\x02 \x01(\v2,.com.coralogixapis.alerts.v3.IntegrationTypeR\vintegration:\x90\x01\x92A\x8c\x01\n" +
	"\x89\x01*!Alert definition webhook settings24Configuration for webhook notifications for an alert\xd2\x01\x13retriggering_period\xd2\x01\tnotify_on\xd2\x01\vintegrationB\x15\n" +
	"\x13retriggering_periodB\f\n" +
	"\n" +
	"_notify_on\"\xa7\x02\n" +
	"\x0fIntegrationType\x12O\n" +
	"\x0eintegration_id\x18\x02 \x01(\v2\x1c.google.protobuf.UInt32ValueB\b\x92A\x05J\x03123H\x00R\rintegrationId\x12I\n" +
	"\n" +
	"recipients\x18\x03 \x01(\v2'.com.coralogixapis.alerts.v3.RecipientsH\x00R\n" +
	"recipients:d\x92Aa\n" +
	"_*\x10Integration type28Defines the type of integration to use for notifications\xd2\x01\x10integration_typeB\x12\n" +
	"\x10integration_type\"\x90\x01\n" +
	"\n" +
	"Recipients\x124\n" +
	"\x06emails\x18\x01 \x03(\v2\x1c.google.protobuf.StringValueR\x06emails:L\x92AI\n" +
	"G*\n" +
	"Recipients20List of email recipients for alert notifications\xd2\x01\x06emailsb\x06proto3"

var (
	file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_goTypes = []any{
	(*AlertDefIncidentSettings)(nil),  // 0: com.coralogixapis.alerts.v3.AlertDefIncidentSettings
	(*AlertDefNotificationGroup)(nil), // 1: com.coralogixapis.alerts.v3.AlertDefNotificationGroup
	(*AlertDefWebhooksSettings)(nil),  // 2: com.coralogixapis.alerts.v3.AlertDefWebhooksSettings
	(*IntegrationType)(nil),           // 3: com.coralogixapis.alerts.v3.IntegrationType
	(*Recipients)(nil),                // 4: com.coralogixapis.alerts.v3.Recipients
	(*wrapperspb.UInt32Value)(nil),    // 5: google.protobuf.UInt32Value
	(NotifyOn)(0),                     // 6: com.coralogixapis.alerts.v3.NotifyOn
	(*wrapperspb.StringValue)(nil),    // 7: google.protobuf.StringValue
	(*NotificationDestination)(nil),   // 8: com.coralogixapis.alerts.v3.NotificationDestination
	(*NotificationRouter)(nil),        // 9: com.coralogixapis.alerts.v3.NotificationRouter
}
var file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_depIdxs = []int32{
	5,  // 0: com.coralogixapis.alerts.v3.AlertDefIncidentSettings.minutes:type_name -> google.protobuf.UInt32Value
	6,  // 1: com.coralogixapis.alerts.v3.AlertDefIncidentSettings.notify_on:type_name -> com.coralogixapis.alerts.v3.NotifyOn
	7,  // 2: com.coralogixapis.alerts.v3.AlertDefNotificationGroup.group_by_keys:type_name -> google.protobuf.StringValue
	2,  // 3: com.coralogixapis.alerts.v3.AlertDefNotificationGroup.webhooks:type_name -> com.coralogixapis.alerts.v3.AlertDefWebhooksSettings
	8,  // 4: com.coralogixapis.alerts.v3.AlertDefNotificationGroup.destinations:type_name -> com.coralogixapis.alerts.v3.NotificationDestination
	9,  // 5: com.coralogixapis.alerts.v3.AlertDefNotificationGroup.router:type_name -> com.coralogixapis.alerts.v3.NotificationRouter
	5,  // 6: com.coralogixapis.alerts.v3.AlertDefWebhooksSettings.minutes:type_name -> google.protobuf.UInt32Value
	6,  // 7: com.coralogixapis.alerts.v3.AlertDefWebhooksSettings.notify_on:type_name -> com.coralogixapis.alerts.v3.NotifyOn
	3,  // 8: com.coralogixapis.alerts.v3.AlertDefWebhooksSettings.integration:type_name -> com.coralogixapis.alerts.v3.IntegrationType
	5,  // 9: com.coralogixapis.alerts.v3.IntegrationType.integration_id:type_name -> google.protobuf.UInt32Value
	4,  // 10: com.coralogixapis.alerts.v3.IntegrationType.recipients:type_name -> com.coralogixapis.alerts.v3.Recipients
	7,  // 11: com.coralogixapis.alerts.v3.Recipients.emails:type_name -> google.protobuf.StringValue
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_init() }
func file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_notification_group_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_commons_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_msgTypes[0].OneofWrappers = []any{
		(*AlertDefIncidentSettings_Minutes)(nil),
	}
	file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_msgTypes[1].OneofWrappers = []any{}
	file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_msgTypes[2].OneofWrappers = []any{
		(*AlertDefWebhooksSettings_Minutes)(nil),
	}
	file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_msgTypes[3].OneofWrappers = []any{
		(*IntegrationType_IntegrationId)(nil),
		(*IntegrationType_Recipients)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_notification_group_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_depIdxs = nil
}
