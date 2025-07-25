// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
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
	unsafe "unsafe"
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// A unique identifier - automatically generated unless provided by the user
	Id            *string              `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	ConnectorType common.ConnectorType `protobuf:"varint,4,opt,name=connector_type,json=connectorType,proto3,enum=com.coralogixapis.notification_center.ConnectorType" json:"connector_type,omitempty"`
	// A list of of configuration override templates, each associated with a specific output schema and condition
	ConfigOverrides []*common.ConfigOverrides `protobuf:"bytes,5,rep,name=config_overrides,json=configOverrides,proto3" json:"config_overrides,omitempty"`
	Name            string                    `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Description     string                    `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	// System-generated timestamp for when the preset was created
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=create_time,json=createTime,proto3,oneof" json:"create_time,omitempty"`
	// System-generated timestamp for when the preset was last updated
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=update_time,json=updateTime,proto3,oneof" json:"update_time,omitempty"`
	// Indicates whether the preset is system-provided or custom-created by the user
	PresetType    *PresetType       `protobuf:"varint,11,opt,name=preset_type,json=presetType,proto3,enum=com.coralogixapis.notification_center.presets.v1.PresetType,oneof" json:"preset_type,omitempty"`
	EntityType    common.EntityType `protobuf:"varint,12,opt,name=entity_type,json=entityType,proto3,enum=com.coralogixapis.notification_center.EntityType" json:"entity_type,omitempty"`
	ParentId      *string           `protobuf:"bytes,13,opt,name=parent_id,json=parentId,proto3,oneof" json:"parent_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

func (x *Preset) GetEntityType() common.EntityType {
	if x != nil {
		return x.EntityType
	}
	return common.EntityType(0)
}

func (x *Preset) GetParentId() string {
	if x != nil && x.ParentId != nil {
		return *x.ParentId
	}
	return ""
}

type PresetSummary struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ConnectorType common.ConnectorType   `protobuf:"varint,4,opt,name=connector_type,json=connectorType,proto3,enum=com.coralogixapis.notification_center.ConnectorType" json:"connector_type,omitempty"`
	ParentId      *string                `protobuf:"bytes,5,opt,name=parent_id,json=parentId,proto3,oneof" json:"parent_id,omitempty"`
	Name          string                 `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	CreateTime    *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime    *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	PresetType    PresetType             `protobuf:"varint,10,opt,name=preset_type,json=presetType,proto3,enum=com.coralogixapis.notification_center.presets.v1.PresetType" json:"preset_type,omitempty"`
	EntityType    common.EntityType      `protobuf:"varint,11,opt,name=entity_type,json=entityType,proto3,enum=com.coralogixapis.notification_center.EntityType" json:"entity_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

func (x *PresetSummary) GetEntityType() common.EntityType {
	if x != nil {
		return x.EntityType
	}
	return common.EntityType(0)
}

type PresetIdentifier struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Value:
	//
	//	*PresetIdentifier_Id
	//	*PresetIdentifier_UserFacingId
	Value         isPresetIdentifier_Value `protobuf_oneof:"value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

func (x *PresetIdentifier) GetValue() isPresetIdentifier_Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *PresetIdentifier) GetId() string {
	if x != nil {
		if x, ok := x.Value.(*PresetIdentifier_Id); ok {
			return x.Id
		}
	}
	return ""
}

func (x *PresetIdentifier) GetUserFacingId() string {
	if x != nil {
		if x, ok := x.Value.(*PresetIdentifier_UserFacingId); ok {
			return x.UserFacingId
		}
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

const file_com_coralogixapis_notification_center_presets_v1_preset_proto_rawDesc = "" +
	"\n" +
	"=com/coralogixapis/notification_center/presets/v1/preset.proto\x120com.coralogixapis.notification_center.presets.v1\x1a9com/coralogixapis/notification_center/common/common.proto\x1aCcom/coralogixapis/notification_center/common/config_overrides.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\x8d\t\n" +
	"\x06Preset\x12@\n" +
	"\x02id\x18\x01 \x01(\tB+\x92A(J&\"a16e24c8-4db2-4abf-ba3c-c9e1fc35a3b9\"H\x00R\x02id\x88\x01\x01\x12[\n" +
	"\x0econnector_type\x18\x04 \x01(\x0e24.com.coralogixapis.notification_center.ConnectorTypeR\rconnectorType\x12a\n" +
	"\x10config_overrides\x18\x05 \x03(\v26.com.coralogixapis.notification_center.ConfigOverridesR\x0fconfigOverrides\x12$\n" +
	"\x04name\x18\a \x01(\tB\x10\x92A\rJ\v\"My Preset\"R\x04name\x12A\n" +
	"\vdescription\x18\b \x01(\tB\x1f\x92A\x1cJ\x1a\"Custom preset for alerts\"R\vdescription\x12@\n" +
	"\vcreate_time\x18\t \x01(\v2\x1a.google.protobuf.TimestampH\x01R\n" +
	"createTime\x88\x01\x01\x12@\n" +
	"\vupdate_time\x18\n" +
	" \x01(\v2\x1a.google.protobuf.TimestampH\x02R\n" +
	"updateTime\x88\x01\x01\x12b\n" +
	"\vpreset_type\x18\v \x01(\x0e2<.com.coralogixapis.notification_center.presets.v1.PresetTypeH\x03R\n" +
	"presetType\x88\x01\x01\x12a\n" +
	"\ventity_type\x18\f \x01(\x0e21.com.coralogixapis.notification_center.EntityTypeB\r\x92A\n" +
	"J\b\"ALERTS\"R\n" +
	"entityType\x12L\n" +
	"\tparent_id\x18\r \x01(\tB*\x92A'J%\"preset_system_slack_alerts_detailed\"H\x04R\bparentId\x88\x01\x01:\xf6\x01\x92A\xf2\x01\n" +
	"s*\x06Preset2ASet of preconfigured templates for notification content rendering\xd2\x01\ventity_type\xd2\x01\x10config_overrides\xd2\x01\x04name*{\n" +
	"'Find out more about notification center\x12Phttps://coralogix.com/docs/user-guides/notification-center/introduction/welcome/B\x05\n" +
	"\x03_idB\x0e\n" +
	"\f_create_timeB\x0e\n" +
	"\f_update_timeB\x0e\n" +
	"\f_preset_typeB\f\n" +
	"\n" +
	"_parent_idJ\x04\b\x02\x10\x03J\x04\b\x03\x10\x04J\x04\b\x06\x10\aR\x0fuser_defined_idR\x16deprecated_entity_typeR\x06parent\"\xa5\a\n" +
	"\rPresetSummary\x12;\n" +
	"\x02id\x18\x01 \x01(\tB+\x92A(J&\"a16e24c8-4db2-4abf-ba3c-c9e1fc35a3b9\"R\x02id\x12[\n" +
	"\x0econnector_type\x18\x04 \x01(\x0e24.com.coralogixapis.notification_center.ConnectorTypeR\rconnectorType\x12M\n" +
	"\tparent_id\x18\x05 \x01(\tB+\x92A(J&\"c246e826-10c2-405e-8d3f-afcc24ad4d15\"H\x00R\bparentId\x88\x01\x01\x12'\n" +
	"\x04name\x18\x06 \x01(\tB\x13\x92A\x10J\v\"My Preset\"x\xc8\x01R\x04name\x12D\n" +
	"\vdescription\x18\a \x01(\tB\"\x92A\x1fJ\x1a\"Custom preset for alerts\"x\x88'R\vdescription\x12;\n" +
	"\vcreate_time\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"createTime\x12;\n" +
	"\vupdate_time\x18\t \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"updateTime\x12]\n" +
	"\vpreset_type\x18\n" +
	" \x01(\x0e2<.com.coralogixapis.notification_center.presets.v1.PresetTypeR\n" +
	"presetType\x12a\n" +
	"\ventity_type\x18\v \x01(\x0e21.com.coralogixapis.notification_center.EntityTypeB\r\x92A\n" +
	"J\b\"ALERTS\"R\n" +
	"entityType:\xbc\x01\x92A\xb8\x01\n" +
	"9*\x0ePreset Summary2'Provides a concise overview of a preset*{\n" +
	"'Find out more about notification center\x12Phttps://coralogix.com/docs/user-guides/notification-center/introduction/welcome/B\f\n" +
	"\n" +
	"_parent_idJ\x04\b\x02\x10\x03J\x04\b\x03\x10\x04R\x0fuser_defined_idR\x16deprecated_entity_type\"U\n" +
	"\x10PresetIdentifier\x12\x10\n" +
	"\x02id\x18\x01 \x01(\tH\x00R\x02id\x12&\n" +
	"\x0euser_facing_id\x18\x02 \x01(\tH\x00R\fuserFacingIdB\a\n" +
	"\x05value*A\n" +
	"\n" +
	"PresetType\x12\x1b\n" +
	"\x17PRESET_TYPE_UNSPECIFIED\x10\x00\x12\n" +
	"\n" +
	"\x06SYSTEM\x10\x01\x12\n" +
	"\n" +
	"\x06CUSTOM\x10\x02B2Z0com/coralogixapis/notification_center/presets/v1b\x06proto3"

var (
	file_com_coralogixapis_notification_center_presets_v1_preset_proto_rawDescOnce sync.Once
	file_com_coralogixapis_notification_center_presets_v1_preset_proto_rawDescData []byte
)

func file_com_coralogixapis_notification_center_presets_v1_preset_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_notification_center_presets_v1_preset_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_notification_center_presets_v1_preset_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_notification_center_presets_v1_preset_proto_rawDesc), len(file_com_coralogixapis_notification_center_presets_v1_preset_proto_rawDesc)))
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
	(common.EntityType)(0),         // 7: com.coralogixapis.notification_center.EntityType
}
var file_com_coralogixapis_notification_center_presets_v1_preset_proto_depIdxs = []int32{
	4,  // 0: com.coralogixapis.notification_center.presets.v1.Preset.connector_type:type_name -> com.coralogixapis.notification_center.ConnectorType
	5,  // 1: com.coralogixapis.notification_center.presets.v1.Preset.config_overrides:type_name -> com.coralogixapis.notification_center.ConfigOverrides
	6,  // 2: com.coralogixapis.notification_center.presets.v1.Preset.create_time:type_name -> google.protobuf.Timestamp
	6,  // 3: com.coralogixapis.notification_center.presets.v1.Preset.update_time:type_name -> google.protobuf.Timestamp
	0,  // 4: com.coralogixapis.notification_center.presets.v1.Preset.preset_type:type_name -> com.coralogixapis.notification_center.presets.v1.PresetType
	7,  // 5: com.coralogixapis.notification_center.presets.v1.Preset.entity_type:type_name -> com.coralogixapis.notification_center.EntityType
	4,  // 6: com.coralogixapis.notification_center.presets.v1.PresetSummary.connector_type:type_name -> com.coralogixapis.notification_center.ConnectorType
	6,  // 7: com.coralogixapis.notification_center.presets.v1.PresetSummary.create_time:type_name -> google.protobuf.Timestamp
	6,  // 8: com.coralogixapis.notification_center.presets.v1.PresetSummary.update_time:type_name -> google.protobuf.Timestamp
	0,  // 9: com.coralogixapis.notification_center.presets.v1.PresetSummary.preset_type:type_name -> com.coralogixapis.notification_center.presets.v1.PresetType
	7,  // 10: com.coralogixapis.notification_center.presets.v1.PresetSummary.entity_type:type_name -> com.coralogixapis.notification_center.EntityType
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_notification_center_presets_v1_preset_proto_rawDesc), len(file_com_coralogixapis_notification_center_presets_v1_preset_proto_rawDesc)),
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
	file_com_coralogixapis_notification_center_presets_v1_preset_proto_goTypes = nil
	file_com_coralogixapis_notification_center_presets_v1_preset_proto_depIdxs = nil
}
