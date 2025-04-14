// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: com/coralogixapis/notification_center/entities/v1/entity_type.proto

package v1

import (
	common "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/notification_center/common"
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

type EntityType struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	EntityName     string                 `protobuf:"bytes,1,opt,name=entity_name,json=entityName,proto3" json:"entity_name,omitempty"`
	Version        string                 `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Schema         string                 `protobuf:"bytes,3,opt,name=schema,proto3" json:"schema,omitempty"`
	EntitySubTypes []*EntitySubType       `protobuf:"bytes,4,rep,name=entity_sub_types,json=entitySubTypes,proto3" json:"entity_sub_types,omitempty"`
	// Fallback if example is not provided for entity subtype
	DefaultExample         string                    `protobuf:"bytes,5,opt,name=default_example,json=defaultExample,proto3" json:"default_example,omitempty"`
	ConnectorSystemPresets []*ConnectorSystemPresets `protobuf:"bytes,6,rep,name=connector_system_presets,json=connectorSystemPresets,proto3" json:"connector_system_presets,omitempty"`
	PresentationDetails    *PresentationDetails      `protobuf:"bytes,7,opt,name=presentation_details,json=presentationDetails,proto3" json:"presentation_details,omitempty"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *EntityType) Reset() {
	*x = EntityType{}
	mi := &file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntityType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityType) ProtoMessage() {}

func (x *EntityType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityType.ProtoReflect.Descriptor instead.
func (*EntityType) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_rawDescGZIP(), []int{0}
}

func (x *EntityType) GetEntityName() string {
	if x != nil {
		return x.EntityName
	}
	return ""
}

func (x *EntityType) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *EntityType) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *EntityType) GetEntitySubTypes() []*EntitySubType {
	if x != nil {
		return x.EntitySubTypes
	}
	return nil
}

func (x *EntityType) GetDefaultExample() string {
	if x != nil {
		return x.DefaultExample
	}
	return ""
}

func (x *EntityType) GetConnectorSystemPresets() []*ConnectorSystemPresets {
	if x != nil {
		return x.ConnectorSystemPresets
	}
	return nil
}

func (x *EntityType) GetPresentationDetails() *PresentationDetails {
	if x != nil {
		return x.PresentationDetails
	}
	return nil
}

type EntitySubType struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Example       *string                `protobuf:"bytes,2,opt,name=example,proto3,oneof" json:"example,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EntitySubType) Reset() {
	*x = EntitySubType{}
	mi := &file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntitySubType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntitySubType) ProtoMessage() {}

func (x *EntitySubType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntitySubType.ProtoReflect.Descriptor instead.
func (*EntitySubType) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_rawDescGZIP(), []int{1}
}

func (x *EntitySubType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EntitySubType) GetExample() string {
	if x != nil && x.Example != nil {
		return *x.Example
	}
	return ""
}

type ConnectorSystemPresets struct {
	state                      protoimpl.MessageState           `protogen:"open.v1"`
	Connector                  common.ConnectorType             `protobuf:"varint,1,opt,name=connector,proto3,enum=com.coralogixapis.notification_center.ConnectorType" json:"connector,omitempty"`
	DefaultPresetUserDefinedId string                           `protobuf:"bytes,2,opt,name=default_preset_user_defined_id,json=defaultPresetUserDefinedId,proto3" json:"default_preset_user_defined_id,omitempty"`
	Presets                    []*ConnectorSystemPresets_Preset `protobuf:"bytes,3,rep,name=presets,proto3" json:"presets,omitempty"`
	unknownFields              protoimpl.UnknownFields
	sizeCache                  protoimpl.SizeCache
}

func (x *ConnectorSystemPresets) Reset() {
	*x = ConnectorSystemPresets{}
	mi := &file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectorSystemPresets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectorSystemPresets) ProtoMessage() {}

func (x *ConnectorSystemPresets) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectorSystemPresets.ProtoReflect.Descriptor instead.
func (*ConnectorSystemPresets) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_rawDescGZIP(), []int{2}
}

func (x *ConnectorSystemPresets) GetConnector() common.ConnectorType {
	if x != nil {
		return x.Connector
	}
	return common.ConnectorType(0)
}

func (x *ConnectorSystemPresets) GetDefaultPresetUserDefinedId() string {
	if x != nil {
		return x.DefaultPresetUserDefinedId
	}
	return ""
}

func (x *ConnectorSystemPresets) GetPresets() []*ConnectorSystemPresets_Preset {
	if x != nil {
		return x.Presets
	}
	return nil
}

type ConnectorSystemPresets_Preset struct {
	state           protoimpl.MessageState    `protogen:"open.v1"`
	UserDefinedId   string                    `protobuf:"bytes,1,opt,name=user_defined_id,json=userDefinedId,proto3" json:"user_defined_id,omitempty"`
	Name            string                    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description     string                    `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	ConfigOverrides []*common.ConfigOverrides `protobuf:"bytes,4,rep,name=config_overrides,json=configOverrides,proto3" json:"config_overrides,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ConnectorSystemPresets_Preset) Reset() {
	*x = ConnectorSystemPresets_Preset{}
	mi := &file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectorSystemPresets_Preset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectorSystemPresets_Preset) ProtoMessage() {}

func (x *ConnectorSystemPresets_Preset) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectorSystemPresets_Preset.ProtoReflect.Descriptor instead.
func (*ConnectorSystemPresets_Preset) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ConnectorSystemPresets_Preset) GetUserDefinedId() string {
	if x != nil {
		return x.UserDefinedId
	}
	return ""
}

func (x *ConnectorSystemPresets_Preset) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConnectorSystemPresets_Preset) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ConnectorSystemPresets_Preset) GetConfigOverrides() []*common.ConfigOverrides {
	if x != nil {
		return x.ConfigOverrides
	}
	return nil
}

var File_com_coralogixapis_notification_center_entities_v1_entity_type_proto protoreflect.FileDescriptor

var file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_rawDesc = []byte{
	0x0a, 0x43, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x31, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x39, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x43, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4c, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x65, 0x73,
	0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x03, 0x0a, 0x0a, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x6a, 0x0a, 0x10, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x5f, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x40, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x75, 0x62,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x75, 0x62, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x83, 0x01,
	0x0a, 0x18, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x49, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x73, 0x52, 0x16, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x65, 0x73,
	0x65, 0x74, 0x73, 0x12, 0x79, 0x0a, 0x14, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x46, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x13, 0x70, 0x72, 0x65, 0x73, 0x65,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x4e,
	0x0a, 0x0d, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0xe8,
	0x03, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x73, 0x12, 0x52, 0x0a, 0x09, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x42, 0x0a,
	0x1e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1a, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x50, 0x72,
	0x65, 0x73, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x49,
	0x64, 0x12, 0x6a, 0x0a, 0x07, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x50, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x50, 0x72,
	0x65, 0x73, 0x65, 0x74, 0x52, 0x07, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x73, 0x1a, 0xc9, 0x01,
	0x0a, 0x06, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x61, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f,
	0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x42, 0x33, 0x5a, 0x31, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_rawDescOnce sync.Once
	file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_rawDescData = file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_rawDesc
)

func file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_rawDescData)
	})
	return file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_rawDescData
}

var file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_goTypes = []any{
	(*EntityType)(nil),                    // 0: com.coralogixapis.notification_center.entities.v1.EntityType
	(*EntitySubType)(nil),                 // 1: com.coralogixapis.notification_center.entities.v1.EntitySubType
	(*ConnectorSystemPresets)(nil),        // 2: com.coralogixapis.notification_center.entities.v1.ConnectorSystemPresets
	(*ConnectorSystemPresets_Preset)(nil), // 3: com.coralogixapis.notification_center.entities.v1.ConnectorSystemPresets.Preset
	(*PresentationDetails)(nil),           // 4: com.coralogixapis.notification_center.entities.v1.PresentationDetails
	(common.ConnectorType)(0),             // 5: com.coralogixapis.notification_center.ConnectorType
	(*common.ConfigOverrides)(nil),        // 6: com.coralogixapis.notification_center.ConfigOverrides
}
var file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.notification_center.entities.v1.EntityType.entity_sub_types:type_name -> com.coralogixapis.notification_center.entities.v1.EntitySubType
	2, // 1: com.coralogixapis.notification_center.entities.v1.EntityType.connector_system_presets:type_name -> com.coralogixapis.notification_center.entities.v1.ConnectorSystemPresets
	4, // 2: com.coralogixapis.notification_center.entities.v1.EntityType.presentation_details:type_name -> com.coralogixapis.notification_center.entities.v1.PresentationDetails
	5, // 3: com.coralogixapis.notification_center.entities.v1.ConnectorSystemPresets.connector:type_name -> com.coralogixapis.notification_center.ConnectorType
	3, // 4: com.coralogixapis.notification_center.entities.v1.ConnectorSystemPresets.presets:type_name -> com.coralogixapis.notification_center.entities.v1.ConnectorSystemPresets.Preset
	6, // 5: com.coralogixapis.notification_center.entities.v1.ConnectorSystemPresets.Preset.config_overrides:type_name -> com.coralogixapis.notification_center.ConfigOverrides
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_init() }
func file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_init() {
	if File_com_coralogixapis_notification_center_entities_v1_entity_type_proto != nil {
		return
	}
	file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_init()
	file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_notification_center_entities_v1_entity_type_proto = out.File
	file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_rawDesc = nil
	file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_goTypes = nil
	file_com_coralogixapis_notification_center_entities_v1_entity_type_proto_depIdxs = nil
}
