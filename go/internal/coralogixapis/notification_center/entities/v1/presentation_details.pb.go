// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: com/coralogixapis/notification_center/entities/v1/presentation_details.proto

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

type PresentationDetails struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Deprecated: Marked as deprecated in com/coralogixapis/notification_center/entities/v1/presentation_details.proto.
	DeprecatedEntityType     string                     `protobuf:"bytes,1,opt,name=deprecated_entity_type,json=deprecatedEntityType,proto3" json:"deprecated_entity_type,omitempty"`
	EntityDisplayName        string                     `protobuf:"bytes,2,opt,name=entity_display_name,json=entityDisplayName,proto3" json:"entity_display_name,omitempty"`
	EntityTypeDescription    string                     `protobuf:"bytes,3,opt,name=entity_type_description,json=entityTypeDescription,proto3" json:"entity_type_description,omitempty"`
	OverrideAll              *OverrideAllDefinition     `protobuf:"bytes,4,opt,name=override_all,json=overrideAll,proto3" json:"override_all,omitempty"`
	SubTypeGrouping          *SubTypeGroupingDefinition `protobuf:"bytes,5,opt,name=sub_type_grouping,json=subTypeGrouping,proto3" json:"sub_type_grouping,omitempty"`
	EntitySubTypeDefinitions []*EntitySubTypeDefinition `protobuf:"bytes,6,rep,name=entity_sub_type_definitions,json=entitySubTypeDefinitions,proto3" json:"entity_sub_type_definitions,omitempty"`
	EntityType               common.EntityType          `protobuf:"varint,7,opt,name=entity_type,json=entityType,proto3,enum=com.coralogixapis.notification_center.EntityType" json:"entity_type,omitempty"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *PresentationDetails) Reset() {
	*x = PresentationDetails{}
	mi := &file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PresentationDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PresentationDetails) ProtoMessage() {}

func (x *PresentationDetails) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PresentationDetails.ProtoReflect.Descriptor instead.
func (*PresentationDetails) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Marked as deprecated in com/coralogixapis/notification_center/entities/v1/presentation_details.proto.
func (x *PresentationDetails) GetDeprecatedEntityType() string {
	if x != nil {
		return x.DeprecatedEntityType
	}
	return ""
}

func (x *PresentationDetails) GetEntityDisplayName() string {
	if x != nil {
		return x.EntityDisplayName
	}
	return ""
}

func (x *PresentationDetails) GetEntityTypeDescription() string {
	if x != nil {
		return x.EntityTypeDescription
	}
	return ""
}

func (x *PresentationDetails) GetOverrideAll() *OverrideAllDefinition {
	if x != nil {
		return x.OverrideAll
	}
	return nil
}

func (x *PresentationDetails) GetSubTypeGrouping() *SubTypeGroupingDefinition {
	if x != nil {
		return x.SubTypeGrouping
	}
	return nil
}

func (x *PresentationDetails) GetEntitySubTypeDefinitions() []*EntitySubTypeDefinition {
	if x != nil {
		return x.EntitySubTypeDefinitions
	}
	return nil
}

func (x *PresentationDetails) GetEntityType() common.EntityType {
	if x != nil {
		return x.EntityType
	}
	return common.EntityType(0)
}

type OverrideAllDefinition struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DisplayName   string                 `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Summary       string                 `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OverrideAllDefinition) Reset() {
	*x = OverrideAllDefinition{}
	mi := &file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OverrideAllDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverrideAllDefinition) ProtoMessage() {}

func (x *OverrideAllDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverrideAllDefinition.ProtoReflect.Descriptor instead.
func (*OverrideAllDefinition) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_rawDescGZIP(), []int{1}
}

func (x *OverrideAllDefinition) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *OverrideAllDefinition) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *OverrideAllDefinition) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type SubTypeGroupingDefinition struct {
	state                 protoimpl.MessageState `protogen:"open.v1"`
	PrimaryGroupingKey    *GroupingKey           `protobuf:"bytes,1,opt,name=primary_grouping_key,json=primaryGroupingKey,proto3" json:"primary_grouping_key,omitempty"`
	SecondaryGroupingKeys []*GroupingKey         `protobuf:"bytes,2,rep,name=secondary_grouping_keys,json=secondaryGroupingKeys,proto3" json:"secondary_grouping_keys,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *SubTypeGroupingDefinition) Reset() {
	*x = SubTypeGroupingDefinition{}
	mi := &file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubTypeGroupingDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubTypeGroupingDefinition) ProtoMessage() {}

func (x *SubTypeGroupingDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubTypeGroupingDefinition.ProtoReflect.Descriptor instead.
func (*SubTypeGroupingDefinition) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_rawDescGZIP(), []int{2}
}

func (x *SubTypeGroupingDefinition) GetPrimaryGroupingKey() *GroupingKey {
	if x != nil {
		return x.PrimaryGroupingKey
	}
	return nil
}

func (x *SubTypeGroupingDefinition) GetSecondaryGroupingKeys() []*GroupingKey {
	if x != nil {
		return x.SecondaryGroupingKeys
	}
	return nil
}

type GroupingKey struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DisplayName   string                 `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GroupingKey) Reset() {
	*x = GroupingKey{}
	mi := &file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GroupingKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupingKey) ProtoMessage() {}

func (x *GroupingKey) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupingKey.ProtoReflect.Descriptor instead.
func (*GroupingKey) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_rawDescGZIP(), []int{3}
}

func (x *GroupingKey) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

type EntitySubTypeDefinition struct {
	state                   protoimpl.MessageState `protogen:"open.v1"`
	EntitySubType           string                 `protobuf:"bytes,1,opt,name=entity_sub_type,json=entitySubType,proto3" json:"entity_sub_type,omitempty"`
	PrimaryGroupingValue    string                 `protobuf:"bytes,2,opt,name=primary_grouping_value,json=primaryGroupingValue,proto3" json:"primary_grouping_value,omitempty"`
	SecondaryGroupingValues map[string]string      `protobuf:"bytes,3,rep,name=secondary_grouping_values,json=secondaryGroupingValues,proto3" json:"secondary_grouping_values,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Variant                 string                 `protobuf:"bytes,4,opt,name=variant,proto3" json:"variant,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *EntitySubTypeDefinition) Reset() {
	*x = EntitySubTypeDefinition{}
	mi := &file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntitySubTypeDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntitySubTypeDefinition) ProtoMessage() {}

func (x *EntitySubTypeDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntitySubTypeDefinition.ProtoReflect.Descriptor instead.
func (*EntitySubTypeDefinition) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_rawDescGZIP(), []int{4}
}

func (x *EntitySubTypeDefinition) GetEntitySubType() string {
	if x != nil {
		return x.EntitySubType
	}
	return ""
}

func (x *EntitySubTypeDefinition) GetPrimaryGroupingValue() string {
	if x != nil {
		return x.PrimaryGroupingValue
	}
	return ""
}

func (x *EntitySubTypeDefinition) GetSecondaryGroupingValues() map[string]string {
	if x != nil {
		return x.SecondaryGroupingValues
	}
	return nil
}

func (x *EntitySubTypeDefinition) GetVariant() string {
	if x != nil {
		return x.Variant
	}
	return ""
}

var File_com_coralogixapis_notification_center_entities_v1_presentation_details_proto protoreflect.FileDescriptor

var file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_rawDesc = []byte{
	0x0a, 0x4c, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x31,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x39, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x04, 0x0a,
	0x13, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x38, 0x0a, 0x16, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x14, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e,
	0x0a, 0x13, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x36,
	0x0a, 0x17, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x15, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6b, 0x0a, 0x0c, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69,
	0x64, 0x65, 0x5f, 0x61, 0x6c, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x41, 0x6c, 0x6c, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65,
	0x41, 0x6c, 0x6c, 0x12, 0x78, 0x0a, 0x11, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4c,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x69,
	0x6e, 0x67, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x73, 0x75,
	0x62, 0x54, 0x79, 0x70, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x89, 0x01,
	0x0a, 0x1b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x4a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x75,
	0x62, 0x54, 0x79, 0x70, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x18, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x44, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x52, 0x0a, 0x0b, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x22, 0x76, 0x0a,
	0x15, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x41, 0x6c, 0x6c, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x85, 0x02, 0x0a, 0x19, 0x53, 0x75, 0x62, 0x54, 0x79, 0x70,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x70, 0x0a, 0x14, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x3e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x67, 0x4b, 0x65,
	0x79, 0x52, 0x12, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x69,
	0x6e, 0x67, 0x4b, 0x65, 0x79, 0x12, 0x76, 0x0a, 0x17, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61,
	0x72, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x6b, 0x65, 0x79, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x52, 0x15, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72,
	0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x73, 0x22, 0x30, 0x0a,
	0x0b, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x0c,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x83, 0x03, 0x0a, 0x17, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x75, 0x62, 0x54, 0x79, 0x70,
	0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x75, 0x62, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x14, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0xa3, 0x01, 0x0a, 0x19, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x67,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x67, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x44,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x61, 0x72, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x17, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72,
	0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x1a, 0x4a, 0x0a, 0x1c, 0x53, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x33, 0x5a, 0x31, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_rawDescOnce sync.Once
	file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_rawDescData = file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_rawDesc
)

func file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_rawDescData)
	})
	return file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_rawDescData
}

var file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_goTypes = []any{
	(*PresentationDetails)(nil),       // 0: com.coralogixapis.notification_center.entities.v1.PresentationDetails
	(*OverrideAllDefinition)(nil),     // 1: com.coralogixapis.notification_center.entities.v1.OverrideAllDefinition
	(*SubTypeGroupingDefinition)(nil), // 2: com.coralogixapis.notification_center.entities.v1.SubTypeGroupingDefinition
	(*GroupingKey)(nil),               // 3: com.coralogixapis.notification_center.entities.v1.GroupingKey
	(*EntitySubTypeDefinition)(nil),   // 4: com.coralogixapis.notification_center.entities.v1.EntitySubTypeDefinition
	nil,                               // 5: com.coralogixapis.notification_center.entities.v1.EntitySubTypeDefinition.SecondaryGroupingValuesEntry
	(common.EntityType)(0),            // 6: com.coralogixapis.notification_center.EntityType
}
var file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.notification_center.entities.v1.PresentationDetails.override_all:type_name -> com.coralogixapis.notification_center.entities.v1.OverrideAllDefinition
	2, // 1: com.coralogixapis.notification_center.entities.v1.PresentationDetails.sub_type_grouping:type_name -> com.coralogixapis.notification_center.entities.v1.SubTypeGroupingDefinition
	4, // 2: com.coralogixapis.notification_center.entities.v1.PresentationDetails.entity_sub_type_definitions:type_name -> com.coralogixapis.notification_center.entities.v1.EntitySubTypeDefinition
	6, // 3: com.coralogixapis.notification_center.entities.v1.PresentationDetails.entity_type:type_name -> com.coralogixapis.notification_center.EntityType
	3, // 4: com.coralogixapis.notification_center.entities.v1.SubTypeGroupingDefinition.primary_grouping_key:type_name -> com.coralogixapis.notification_center.entities.v1.GroupingKey
	3, // 5: com.coralogixapis.notification_center.entities.v1.SubTypeGroupingDefinition.secondary_grouping_keys:type_name -> com.coralogixapis.notification_center.entities.v1.GroupingKey
	5, // 6: com.coralogixapis.notification_center.entities.v1.EntitySubTypeDefinition.secondary_grouping_values:type_name -> com.coralogixapis.notification_center.entities.v1.EntitySubTypeDefinition.SecondaryGroupingValuesEntry
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_init() }
func file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_init() {
	if File_com_coralogixapis_notification_center_entities_v1_presentation_details_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_notification_center_entities_v1_presentation_details_proto = out.File
	file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_rawDesc = nil
	file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_goTypes = nil
	file_com_coralogixapis_notification_center_entities_v1_presentation_details_proto_depIdxs = nil
}
