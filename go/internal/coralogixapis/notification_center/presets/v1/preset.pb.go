// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: com/coralogixapis/notification_center/presets/v1/preset.proto

package v1

import (
	common "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/notification_center/common"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PresetType int32

const (
	PresetType_PRESET_TYPE_UNSPECIFIED PresetType = 0
	PresetType_SYSTEM                  PresetType = 1
	PresetType_CUSTOM                  PresetType = 2
)

// Enum value maps for PresetType.
var (
	PresetType_name = map[int32]string{
		0: "PRESET_TYPE_UNSPECIFIED",
		1: "SYSTEM",
		2: "CUSTOM",
	}
	PresetType_value = map[string]int32{
		"PRESET_TYPE_UNSPECIFIED": 0,
		"SYSTEM":                  1,
		"CUSTOM":                  2,
	}
)

func (x PresetType) Enum() *PresetType {
	p := new(PresetType)
	*p = x
	return p
}

func (x PresetType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PresetType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_notification_center_presets_v1_preset_proto_enumTypes[0].Descriptor()
}

func (PresetType) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_notification_center_presets_v1_preset_proto_enumTypes[0]
}

func (x PresetType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PresetType.Descriptor instead.
func (PresetType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_presets_v1_preset_proto_rawDescGZIP(), []int{0}
}

type Preset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              *string                   `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	UserFacingId    *string                   `protobuf:"bytes,2,opt,name=user_facing_id,json=userFacingId,proto3,oneof" json:"user_facing_id,omitempty"`
	EntityType      string                    `protobuf:"bytes,3,opt,name=entity_type,json=entityType,proto3" json:"entity_type,omitempty"`
	ConnectorType   common.ConnectorType      `protobuf:"varint,4,opt,name=connector_type,json=connectorType,proto3,enum=com.coralogixapis.notification_center.ConnectorType" json:"connector_type,omitempty"`
	ConfigOverrides []*common.ConfigOverrides `protobuf:"bytes,5,rep,name=config_overrides,json=configOverrides,proto3" json:"config_overrides,omitempty"`
	Parent          *Preset                   `protobuf:"bytes,6,opt,name=parent,proto3,oneof" json:"parent,omitempty"`
	Name            string                    `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Description     string                    `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	CreateTime      *timestamppb.Timestamp    `protobuf:"bytes,9,opt,name=create_time,json=createTime,proto3,oneof" json:"create_time,omitempty"`
	UpdateTime      *timestamppb.Timestamp    `protobuf:"bytes,10,opt,name=update_time,json=updateTime,proto3,oneof" json:"update_time,omitempty"`
	PresetType      *PresetType               `protobuf:"varint,11,opt,name=preset_type,json=presetType,proto3,enum=com.coralogixapis.notification_center.presets.v1.PresetType,oneof" json:"preset_type,omitempty"`
}

func (x *Preset) Reset() {
	*x = Preset{}
	mi := &file_com_coralogixapis_notification_center_presets_v1_preset_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Preset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Preset) ProtoMessage() {}

func (x *Preset) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_presets_v1_preset_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Preset.ProtoReflect.Descriptor instead.
func (*Preset) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_presets_v1_preset_proto_rawDescGZIP(), []int{0}
}

func (x *Preset) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *Preset) GetUserFacingId() string {
	if x != nil && x.UserFacingId != nil {
		return *x.UserFacingId
	}
	return ""
}

func (x *Preset) GetEntityType() string {
	if x != nil {
		return x.EntityType
	}
	return ""
}

func (x *Preset) GetConnectorType() common.ConnectorType {
	if x != nil {
		return x.ConnectorType
	}
	return common.ConnectorType(0)
}

func (x *Preset) GetConfigOverrides() []*common.ConfigOverrides {
	if x != nil {
		return x.ConfigOverrides
	}
	return nil
}

func (x *Preset) GetParent() *Preset {
	if x != nil {
		return x.Parent
	}
	return nil
}

func (x *Preset) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Preset) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Preset) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Preset) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Preset) GetPresetType() PresetType {
	if x != nil && x.PresetType != nil {
		return *x.PresetType
	}
	return PresetType_PRESET_TYPE_UNSPECIFIED
}

type PresetSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserFacingId  string                 `protobuf:"bytes,2,opt,name=user_facing_id,json=userFacingId,proto3" json:"user_facing_id,omitempty"`
	EntityType    string                 `protobuf:"bytes,3,opt,name=entity_type,json=entityType,proto3" json:"entity_type,omitempty"`
	ConnectorType common.ConnectorType   `protobuf:"varint,4,opt,name=connector_type,json=connectorType,proto3,enum=com.coralogixapis.notification_center.ConnectorType" json:"connector_type,omitempty"`
	ParentId      *string                `protobuf:"bytes,5,opt,name=parent_id,json=parentId,proto3,oneof" json:"parent_id,omitempty"`
	Name          string                 `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	CreateTime    *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime    *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	PresetType    PresetType             `protobuf:"varint,10,opt,name=preset_type,json=presetType,proto3,enum=com.coralogixapis.notification_center.presets.v1.PresetType" json:"preset_type,omitempty"`
}

func (x *PresetSummary) Reset() {
	*x = PresetSummary{}
	mi := &file_com_coralogixapis_notification_center_presets_v1_preset_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PresetSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PresetSummary) ProtoMessage() {}

func (x *PresetSummary) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_presets_v1_preset_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PresetSummary.ProtoReflect.Descriptor instead.
func (*PresetSummary) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_presets_v1_preset_proto_rawDescGZIP(), []int{1}
}

func (x *PresetSummary) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PresetSummary) GetUserFacingId() string {
	if x != nil {
		return x.UserFacingId
	}
	return ""
}

func (x *PresetSummary) GetEntityType() string {
	if x != nil {
		return x.EntityType
	}
	return ""
}

func (x *PresetSummary) GetConnectorType() common.ConnectorType {
	if x != nil {
		return x.ConnectorType
	}
	return common.ConnectorType(0)
}

func (x *PresetSummary) GetParentId() string {
	if x != nil && x.ParentId != nil {
		return *x.ParentId
	}
	return ""
}

func (x *PresetSummary) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PresetSummary) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PresetSummary) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *PresetSummary) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *PresetSummary) GetPresetType() PresetType {
	if x != nil {
		return x.PresetType
	}
	return PresetType_PRESET_TYPE_UNSPECIFIED
}

type PresetIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*PresetIdentifier_Id
	//	*PresetIdentifier_UserFacingId
	Value isPresetIdentifier_Value `protobuf_oneof:"value"`
}

func (x *PresetIdentifier) Reset() {
	*x = PresetIdentifier{}
	mi := &file_com_coralogixapis_notification_center_presets_v1_preset_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PresetIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PresetIdentifier) ProtoMessage() {}

func (x *PresetIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_presets_v1_preset_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PresetIdentifier.ProtoReflect.Descriptor instead.
func (*PresetIdentifier) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_presets_v1_preset_proto_rawDescGZIP(), []int{2}
}

func (m *PresetIdentifier) GetValue() isPresetIdentifier_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *PresetIdentifier) GetId() string {
	if x, ok := x.GetValue().(*PresetIdentifier_Id); ok {
		return x.Id
	}
	return ""
}

func (x *PresetIdentifier) GetUserFacingId() string {
	if x, ok := x.GetValue().(*PresetIdentifier_UserFacingId); ok {
		return x.UserFacingId
	}
	return ""
}

type isPresetIdentifier_Value interface {
	isPresetIdentifier_Value()
}

type PresetIdentifier_Id struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type PresetIdentifier_UserFacingId struct {
	UserFacingId string `protobuf:"bytes,2,opt,name=user_facing_id,json=userFacingId,proto3,oneof"`
}

func (*PresetIdentifier_Id) isPresetIdentifier_Value() {}

func (*PresetIdentifier_UserFacingId) isPresetIdentifier_Value() {}

var File_com_coralogixapis_notification_center_presets_v1_preset_proto protoreflect.FileDescriptor

var file_com_coralogixapis_notification_center_presets_v1_preset_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x30, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x39, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x43, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb7, 0x08, 0x0a, 0x06, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x12, 0x40, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0x92, 0x41, 0x28, 0x4a, 0x26,
	0x22, 0x61, 0x31, 0x36, 0x65, 0x32, 0x34, 0x63, 0x38, 0x2d, 0x34, 0x64, 0x62, 0x32, 0x2d, 0x34,
	0x61, 0x62, 0x66, 0x2d, 0x62, 0x61, 0x33, 0x63, 0x2d, 0x63, 0x39, 0x65, 0x31, 0x66, 0x63, 0x33,
	0x35, 0x61, 0x33, 0x62, 0x39, 0x22, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x40, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x66, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0x92, 0x41, 0x12, 0x4a, 0x10, 0x22, 0x75,
	0x73, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x2d, 0x69, 0x64, 0x22, 0x48, 0x01,
	0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x46, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x2e, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0x92, 0x41, 0x0a, 0x4a, 0x08, 0x22, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x22, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x5b, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x61,
	0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73,
	0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65,
	0x73, 0x12, 0x55, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x48, 0x02, 0x52, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0x92, 0x41, 0x0d, 0x4a, 0x0b, 0x22, 0x4d, 0x79,
	0x20, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x22, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x41,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x1f, 0x92, 0x41, 0x1c, 0x4a, 0x1a, 0x22, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x20, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x73, 0x22, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x48, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x48, 0x04, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x62, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3c, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72,
	0x65, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x48, 0x05, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x73,
	0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x3a, 0xbb, 0x01, 0x92, 0x41, 0xb7, 0x01,
	0x0a, 0x38, 0x2a, 0x06, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74,
	0x20, 0x66, 0x6f, 0x72, 0x20, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x20, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2a, 0x7b, 0x0a, 0x27, 0x46, 0x69,
	0x6e, 0x64, 0x20, 0x6f, 0x75, 0x74, 0x20, 0x6d, 0x6f, 0x72, 0x65, 0x20, 0x61, 0x62, 0x6f, 0x75,
	0x74, 0x20, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x50, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x63,
	0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x67, 0x75, 0x69, 0x64, 0x65, 0x73, 0x2f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2f, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x77,
	0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x2f, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x11,
	0x0a, 0x0f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x66, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x69,
	0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0xe9, 0x06, 0x0a,
	0x0d, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x3b,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0x92, 0x41, 0x28, 0x4a,
	0x26, 0x22, 0x61, 0x31, 0x36, 0x65, 0x32, 0x34, 0x63, 0x38, 0x2d, 0x34, 0x64, 0x62, 0x32, 0x2d,
	0x34, 0x61, 0x62, 0x66, 0x2d, 0x62, 0x61, 0x33, 0x63, 0x2d, 0x63, 0x39, 0x65, 0x31, 0x66, 0x63,
	0x33, 0x35, 0x61, 0x33, 0x62, 0x39, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3b, 0x0a, 0x0e, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x66, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x15, 0x92, 0x41, 0x12, 0x4a, 0x10, 0x22, 0x75, 0x73, 0x65, 0x72, 0x2d,
	0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x2d, 0x69, 0x64, 0x22, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72,
	0x46, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0x92,
	0x41, 0x0a, 0x4a, 0x08, 0x22, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x22, 0x52, 0x0a, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x5b, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0x92, 0x41, 0x14, 0x4a, 0x12, 0x22,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x2d, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x2d, 0x69, 0x64,
	0x22, 0x48, 0x00, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10,
	0x92, 0x41, 0x0d, 0x4a, 0x0b, 0x22, 0x4d, 0x79, 0x20, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x22,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1f, 0x92, 0x41, 0x1c,
	0x4a, 0x1a, 0x22, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x20, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74,
	0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x22, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x5d, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x73,
	0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x3a, 0xc5, 0x01, 0x92, 0x41, 0xc1, 0x01, 0x0a, 0x42, 0x2a, 0x0e, 0x50, 0x72, 0x65,
	0x73, 0x65, 0x74, 0x20, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x32, 0x30, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x20, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x20, 0x61, 0x20, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x20,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x7b, 0x0a,
	0x27, 0x46, 0x69, 0x6e, 0x64, 0x20, 0x6f, 0x75, 0x74, 0x20, 0x6d, 0x6f, 0x72, 0x65, 0x20, 0x61,
	0x62, 0x6f, 0x75, 0x74, 0x20, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x50, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x6f, 0x63, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x67, 0x75, 0x69, 0x64, 0x65, 0x73,
	0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x77, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x2f, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x22, 0x55, 0x0a, 0x10, 0x50, 0x72, 0x65, 0x73,
	0x65, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26,
	0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x66, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x46, 0x61,
	0x63, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a,
	0x41, 0x0a, 0x0a, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a,
	0x17, 0x50, 0x52, 0x45, 0x53, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x59,
	0x53, 0x54, 0x45, 0x4d, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d,
	0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_notification_center_presets_v1_preset_proto_rawDescOnce sync.Once
	file_com_coralogixapis_notification_center_presets_v1_preset_proto_rawDescData = file_com_coralogixapis_notification_center_presets_v1_preset_proto_rawDesc
)

func file_com_coralogixapis_notification_center_presets_v1_preset_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_notification_center_presets_v1_preset_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_notification_center_presets_v1_preset_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_notification_center_presets_v1_preset_proto_rawDescData)
	})
	return file_com_coralogixapis_notification_center_presets_v1_preset_proto_rawDescData
}

var file_com_coralogixapis_notification_center_presets_v1_preset_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_notification_center_presets_v1_preset_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogixapis_notification_center_presets_v1_preset_proto_goTypes = []any{
	(PresetType)(0),                // 0: com.coralogixapis.notification_center.presets.v1.PresetType
	(*Preset)(nil),                 // 1: com.coralogixapis.notification_center.presets.v1.Preset
	(*PresetSummary)(nil),          // 2: com.coralogixapis.notification_center.presets.v1.PresetSummary
	(*PresetIdentifier)(nil),       // 3: com.coralogixapis.notification_center.presets.v1.PresetIdentifier
	(common.ConnectorType)(0),      // 4: com.coralogixapis.notification_center.ConnectorType
	(*common.ConfigOverrides)(nil), // 5: com.coralogixapis.notification_center.ConfigOverrides
	(*timestamppb.Timestamp)(nil),  // 6: google.protobuf.Timestamp
}
var file_com_coralogixapis_notification_center_presets_v1_preset_proto_depIdxs = []int32{
	4,  // 0: com.coralogixapis.notification_center.presets.v1.Preset.connector_type:type_name -> com.coralogixapis.notification_center.ConnectorType
	5,  // 1: com.coralogixapis.notification_center.presets.v1.Preset.config_overrides:type_name -> com.coralogixapis.notification_center.ConfigOverrides
	1,  // 2: com.coralogixapis.notification_center.presets.v1.Preset.parent:type_name -> com.coralogixapis.notification_center.presets.v1.Preset
	6,  // 3: com.coralogixapis.notification_center.presets.v1.Preset.create_time:type_name -> google.protobuf.Timestamp
	6,  // 4: com.coralogixapis.notification_center.presets.v1.Preset.update_time:type_name -> google.protobuf.Timestamp
	0,  // 5: com.coralogixapis.notification_center.presets.v1.Preset.preset_type:type_name -> com.coralogixapis.notification_center.presets.v1.PresetType
	4,  // 6: com.coralogixapis.notification_center.presets.v1.PresetSummary.connector_type:type_name -> com.coralogixapis.notification_center.ConnectorType
	6,  // 7: com.coralogixapis.notification_center.presets.v1.PresetSummary.create_time:type_name -> google.protobuf.Timestamp
	6,  // 8: com.coralogixapis.notification_center.presets.v1.PresetSummary.update_time:type_name -> google.protobuf.Timestamp
	0,  // 9: com.coralogixapis.notification_center.presets.v1.PresetSummary.preset_type:type_name -> com.coralogixapis.notification_center.presets.v1.PresetType
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_notification_center_presets_v1_preset_proto_init() }
func file_com_coralogixapis_notification_center_presets_v1_preset_proto_init() {
	if File_com_coralogixapis_notification_center_presets_v1_preset_proto != nil {
		return
	}
	file_com_coralogixapis_notification_center_presets_v1_preset_proto_msgTypes[0].OneofWrappers = []any{}
	file_com_coralogixapis_notification_center_presets_v1_preset_proto_msgTypes[1].OneofWrappers = []any{}
	file_com_coralogixapis_notification_center_presets_v1_preset_proto_msgTypes[2].OneofWrappers = []any{
		(*PresetIdentifier_Id)(nil),
		(*PresetIdentifier_UserFacingId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_notification_center_presets_v1_preset_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_notification_center_presets_v1_preset_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_notification_center_presets_v1_preset_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_notification_center_presets_v1_preset_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_notification_center_presets_v1_preset_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_notification_center_presets_v1_preset_proto = out.File
	file_com_coralogixapis_notification_center_presets_v1_preset_proto_rawDesc = nil
	file_com_coralogixapis_notification_center_presets_v1_preset_proto_goTypes = nil
	file_com_coralogixapis_notification_center_presets_v1_preset_proto_depIdxs = nil
}
