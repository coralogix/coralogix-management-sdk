// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/notification_destination.proto

package v3

import (
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

type NotificationDestination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectorId               string               `protobuf:"bytes,1,opt,name=connector_id,json=connectorId,proto3" json:"connector_id,omitempty"`
	PresetId                  *string              `protobuf:"bytes,2,opt,name=preset_id,json=presetId,proto3,oneof" json:"preset_id,omitempty"`
	NotifyOn                  NotifyOn             `protobuf:"varint,3,opt,name=notify_on,json=notifyOn,proto3,enum=com.coralogixapis.alerts.v3.NotifyOn" json:"notify_on,omitempty"`
	TriggeredRoutingOverrides *NotificationRouting `protobuf:"bytes,4,opt,name=triggeredRoutingOverrides,proto3" json:"triggeredRoutingOverrides,omitempty"`
	ResolvedRouteOverrides    *NotificationRouting `protobuf:"bytes,5,opt,name=resolvedRouteOverrides,proto3,oneof" json:"resolvedRouteOverrides,omitempty"`
}

func (x *NotificationDestination) Reset() {
	*x = NotificationDestination{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotificationDestination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationDestination) ProtoMessage() {}

func (x *NotificationDestination) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationDestination.ProtoReflect.Descriptor instead.
func (*NotificationDestination) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_rawDescGZIP(), []int{0}
}

func (x *NotificationDestination) GetConnectorId() string {
	if x != nil {
		return x.ConnectorId
	}
	return ""
}

func (x *NotificationDestination) GetPresetId() string {
	if x != nil && x.PresetId != nil {
		return *x.PresetId
	}
	return ""
}

func (x *NotificationDestination) GetNotifyOn() NotifyOn {
	if x != nil {
		return x.NotifyOn
	}
	return NotifyOn_NOTIFY_ON_TRIGGERED_ONLY_UNSPECIFIED
}

func (x *NotificationDestination) GetTriggeredRoutingOverrides() *NotificationRouting {
	if x != nil {
		return x.TriggeredRoutingOverrides
	}
	return nil
}

func (x *NotificationDestination) GetResolvedRouteOverrides() *NotificationRouting {
	if x != nil {
		return x.ResolvedRouteOverrides
	}
	return nil
}

type NotificationRouting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigOverrides *SourceOverrides `protobuf:"bytes,1,opt,name=config_overrides,json=configOverrides,proto3,oneof" json:"config_overrides,omitempty"`
}

func (x *NotificationRouting) Reset() {
	*x = NotificationRouting{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotificationRouting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationRouting) ProtoMessage() {}

func (x *NotificationRouting) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationRouting.ProtoReflect.Descriptor instead.
func (*NotificationRouting) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_rawDescGZIP(), []int{1}
}

func (x *NotificationRouting) GetConfigOverrides() *SourceOverrides {
	if x != nil {
		return x.ConfigOverrides
	}
	return nil
}

type SourceOverrides struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutputSchemaId        string                  `protobuf:"bytes,1,opt,name=output_schema_id,json=outputSchemaId,proto3" json:"output_schema_id,omitempty"`
	MessageConfigFields   []*MessageConfigField   `protobuf:"bytes,2,rep,name=message_config_fields,json=messageConfigFields,proto3" json:"message_config_fields,omitempty"`
	ConnectorConfigFields []*ConnectorConfigField `protobuf:"bytes,3,rep,name=connector_config_fields,json=connectorConfigFields,proto3" json:"connector_config_fields,omitempty"`
}

func (x *SourceOverrides) Reset() {
	*x = SourceOverrides{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SourceOverrides) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceOverrides) ProtoMessage() {}

func (x *SourceOverrides) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceOverrides.ProtoReflect.Descriptor instead.
func (*SourceOverrides) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_rawDescGZIP(), []int{2}
}

func (x *SourceOverrides) GetOutputSchemaId() string {
	if x != nil {
		return x.OutputSchemaId
	}
	return ""
}

func (x *SourceOverrides) GetMessageConfigFields() []*MessageConfigField {
	if x != nil {
		return x.MessageConfigFields
	}
	return nil
}

func (x *SourceOverrides) GetConnectorConfigFields() []*ConnectorConfigField {
	if x != nil {
		return x.ConnectorConfigFields
	}
	return nil
}

type ConnectorConfigField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldName string `protobuf:"bytes,1,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	Template  string `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *ConnectorConfigField) Reset() {
	*x = ConnectorConfigField{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectorConfigField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectorConfigField) ProtoMessage() {}

func (x *ConnectorConfigField) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectorConfigField.ProtoReflect.Descriptor instead.
func (*ConnectorConfigField) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_rawDescGZIP(), []int{3}
}

func (x *ConnectorConfigField) GetFieldName() string {
	if x != nil {
		return x.FieldName
	}
	return ""
}

func (x *ConnectorConfigField) GetTemplate() string {
	if x != nil {
		return x.Template
	}
	return ""
}

type MessageConfigField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldName string `protobuf:"bytes,1,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	Template  string `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *MessageConfigField) Reset() {
	*x = MessageConfigField{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageConfigField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageConfigField) ProtoMessage() {}

func (x *MessageConfigField) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageConfigField.ProtoReflect.Descriptor instead.
func (*MessageConfigField) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_rawDescGZIP(), []int{4}
}

func (x *MessageConfigField) GetFieldName() string {
	if x != nil {
		return x.FieldName
	}
	return ""
}

func (x *MessageConfigField) GetTemplate() string {
	if x != nil {
		return x.Template
	}
	return ""
}

var File_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_rawDesc = []byte{
	0x0a, 0x54, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x2e, 0x76, 0x33, 0x1a, 0x29, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa,
	0x03, 0x0a, 0x17, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a,
	0x09, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x08, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x42, 0x0a, 0x09, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33,
	0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x6e, 0x52, 0x08, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x4f, 0x6e, 0x12, 0x6e, 0x0a, 0x19, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x65, 0x64,
	0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x19, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69,
	0x64, 0x65, 0x73, 0x12, 0x6d, 0x0a, 0x16, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76,
	0x33, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x48, 0x01, 0x52, 0x16, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x88,
	0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64,
	0x42, 0x19, 0x0a, 0x17, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x22, 0x88, 0x01, 0x0a, 0x13,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x5c, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6f, 0x76,
	0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x48, 0x00, 0x52, 0x0f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x88, 0x01,
	0x01, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6f, 0x76, 0x65,
	0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x22, 0x8b, 0x02, 0x0a, 0x0f, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x49, 0x64, 0x12, 0x63, 0x0a, 0x15, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76,
	0x33, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x52, 0x13, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x69, 0x0a, 0x17, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x15, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x22, 0x51, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x4f, 0x0a, 0x12, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_rawDescData = file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_goTypes = []any{
	(*NotificationDestination)(nil), // 0: com.coralogixapis.alerts.v3.NotificationDestination
	(*NotificationRouting)(nil),     // 1: com.coralogixapis.alerts.v3.NotificationRouting
	(*SourceOverrides)(nil),         // 2: com.coralogixapis.alerts.v3.SourceOverrides
	(*ConnectorConfigField)(nil),    // 3: com.coralogixapis.alerts.v3.ConnectorConfigField
	(*MessageConfigField)(nil),      // 4: com.coralogixapis.alerts.v3.MessageConfigField
	(NotifyOn)(0),                   // 5: com.coralogixapis.alerts.v3.NotifyOn
}
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_depIdxs = []int32{
	5, // 0: com.coralogixapis.alerts.v3.NotificationDestination.notify_on:type_name -> com.coralogixapis.alerts.v3.NotifyOn
	1, // 1: com.coralogixapis.alerts.v3.NotificationDestination.triggeredRoutingOverrides:type_name -> com.coralogixapis.alerts.v3.NotificationRouting
	1, // 2: com.coralogixapis.alerts.v3.NotificationDestination.resolvedRouteOverrides:type_name -> com.coralogixapis.alerts.v3.NotificationRouting
	2, // 3: com.coralogixapis.alerts.v3.NotificationRouting.config_overrides:type_name -> com.coralogixapis.alerts.v3.SourceOverrides
	4, // 4: com.coralogixapis.alerts.v3.SourceOverrides.message_config_fields:type_name -> com.coralogixapis.alerts.v3.MessageConfigField
	3, // 5: com.coralogixapis.alerts.v3.SourceOverrides.connector_config_fields:type_name -> com.coralogixapis.alerts.v3.ConnectorConfigField
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_init()
}
func file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_commons_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_msgTypes[0].OneofWrappers = []any{}
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_notification_destination_proto_depIdxs = nil
}
