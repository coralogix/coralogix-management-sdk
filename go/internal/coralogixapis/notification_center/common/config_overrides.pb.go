// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/notification_center/common/config_overrides.proto

package common

import (
	v1 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/notification_center/common/v1"
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

type ConfigOverrides struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	ConditionType  *ConditionType         `protobuf:"bytes,1,opt,name=condition_type,json=conditionType,proto3" json:"condition_type,omitempty"`
	OutputSchemaId string                 `protobuf:"bytes,2,opt,name=output_schema_id,json=outputSchemaId,proto3" json:"output_schema_id,omitempty"`
	MessageConfig  *MessageConfig         `protobuf:"bytes,3,opt,name=message_config,json=messageConfig,proto3" json:"message_config,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ConfigOverrides) Reset() {
	*x = ConfigOverrides{}
	mi := &file_com_coralogixapis_notification_center_common_config_overrides_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfigOverrides) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigOverrides) ProtoMessage() {}

func (x *ConfigOverrides) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_common_config_overrides_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigOverrides.ProtoReflect.Descriptor instead.
func (*ConfigOverrides) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_common_config_overrides_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigOverrides) GetConditionType() *ConditionType {
	if x != nil {
		return x.ConditionType
	}
	return nil
}

func (x *ConfigOverrides) GetOutputSchemaId() string {
	if x != nil {
		return x.OutputSchemaId
	}
	return ""
}

func (x *ConfigOverrides) GetMessageConfig() *MessageConfig {
	if x != nil {
		return x.MessageConfig
	}
	return nil
}

type ConditionType struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Condition:
	//
	//	*ConditionType_MatchEntityType
	//	*ConditionType_MatchEntityTypeAndSubType
	Condition     isConditionType_Condition `protobuf_oneof:"condition"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConditionType) Reset() {
	*x = ConditionType{}
	mi := &file_com_coralogixapis_notification_center_common_config_overrides_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConditionType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConditionType) ProtoMessage() {}

func (x *ConditionType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_common_config_overrides_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConditionType.ProtoReflect.Descriptor instead.
func (*ConditionType) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_common_config_overrides_proto_rawDescGZIP(), []int{1}
}

func (x *ConditionType) GetCondition() isConditionType_Condition {
	if x != nil {
		return x.Condition
	}
	return nil
}

func (x *ConditionType) GetMatchEntityType() *MatchEntityTypeCondition {
	if x != nil {
		if x, ok := x.Condition.(*ConditionType_MatchEntityType); ok {
			return x.MatchEntityType
		}
	}
	return nil
}

func (x *ConditionType) GetMatchEntityTypeAndSubType() *MatchEntityTypeAndSubTypeCondition {
	if x != nil {
		if x, ok := x.Condition.(*ConditionType_MatchEntityTypeAndSubType); ok {
			return x.MatchEntityTypeAndSubType
		}
	}
	return nil
}

type isConditionType_Condition interface {
	isConditionType_Condition()
}

type ConditionType_MatchEntityType struct {
	MatchEntityType *MatchEntityTypeCondition `protobuf:"bytes,1,opt,name=match_entity_type,json=matchEntityType,proto3,oneof"`
}

type ConditionType_MatchEntityTypeAndSubType struct {
	MatchEntityTypeAndSubType *MatchEntityTypeAndSubTypeCondition `protobuf:"bytes,2,opt,name=match_entity_type_and_sub_type,json=matchEntityTypeAndSubType,proto3,oneof"`
}

func (*ConditionType_MatchEntityType) isConditionType_Condition() {}

func (*ConditionType_MatchEntityTypeAndSubType) isConditionType_Condition() {}

type MessageConfig struct {
	state         protoimpl.MessageState   `protogen:"open.v1"`
	Fields        []*v1.MessageConfigField `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageConfig) Reset() {
	*x = MessageConfig{}
	mi := &file_com_coralogixapis_notification_center_common_config_overrides_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageConfig) ProtoMessage() {}

func (x *MessageConfig) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_common_config_overrides_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageConfig.ProtoReflect.Descriptor instead.
func (*MessageConfig) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_common_config_overrides_proto_rawDescGZIP(), []int{2}
}

func (x *MessageConfig) GetFields() []*v1.MessageConfigField {
	if x != nil {
		return x.Fields
	}
	return nil
}

type MatchEntityTypeCondition struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EntityType    string                 `protobuf:"bytes,1,opt,name=entity_type,json=entityType,proto3" json:"entity_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MatchEntityTypeCondition) Reset() {
	*x = MatchEntityTypeCondition{}
	mi := &file_com_coralogixapis_notification_center_common_config_overrides_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MatchEntityTypeCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchEntityTypeCondition) ProtoMessage() {}

func (x *MatchEntityTypeCondition) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_common_config_overrides_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchEntityTypeCondition.ProtoReflect.Descriptor instead.
func (*MatchEntityTypeCondition) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_common_config_overrides_proto_rawDescGZIP(), []int{3}
}

func (x *MatchEntityTypeCondition) GetEntityType() string {
	if x != nil {
		return x.EntityType
	}
	return ""
}

type MatchEntityTypeAndSubTypeCondition struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EntityType    string                 `protobuf:"bytes,1,opt,name=entity_type,json=entityType,proto3" json:"entity_type,omitempty"`
	EntitySubType string                 `protobuf:"bytes,2,opt,name=entity_sub_type,json=entitySubType,proto3" json:"entity_sub_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MatchEntityTypeAndSubTypeCondition) Reset() {
	*x = MatchEntityTypeAndSubTypeCondition{}
	mi := &file_com_coralogixapis_notification_center_common_config_overrides_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MatchEntityTypeAndSubTypeCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchEntityTypeAndSubTypeCondition) ProtoMessage() {}

func (x *MatchEntityTypeAndSubTypeCondition) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_common_config_overrides_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchEntityTypeAndSubTypeCondition.ProtoReflect.Descriptor instead.
func (*MatchEntityTypeAndSubTypeCondition) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_common_config_overrides_proto_rawDescGZIP(), []int{4}
}

func (x *MatchEntityTypeAndSubTypeCondition) GetEntityType() string {
	if x != nil {
		return x.EntityType
	}
	return ""
}

func (x *MatchEntityTypeAndSubTypeCondition) GetEntitySubType() string {
	if x != nil {
		return x.EntitySubType
	}
	return ""
}

var File_com_coralogixapis_notification_center_common_config_overrides_proto protoreflect.FileDescriptor

const file_com_coralogixapis_notification_center_common_config_overrides_proto_rawDesc = "" +
	"\n" +
	"Ccom/coralogixapis/notification_center/common/config_overrides.proto\x12%com.coralogixapis.notification_center\x1aCcom/coralogixapis/notification_center/common/v1/config_fields.proto\"\xf5\x01\n" +
	"\x0fConfigOverrides\x12[\n" +
	"\x0econdition_type\x18\x01 \x01(\v24.com.coralogixapis.notification_center.ConditionTypeR\rconditionType\x12(\n" +
	"\x10output_schema_id\x18\x02 \x01(\tR\x0eoutputSchemaId\x12[\n" +
	"\x0emessage_config\x18\x03 \x01(\v24.com.coralogixapis.notification_center.MessageConfigR\rmessageConfig\"\x9c\x02\n" +
	"\rConditionType\x12m\n" +
	"\x11match_entity_type\x18\x01 \x01(\v2?.com.coralogixapis.notification_center.MatchEntityTypeConditionH\x00R\x0fmatchEntityType\x12\x8e\x01\n" +
	"\x1ematch_entity_type_and_sub_type\x18\x02 \x01(\v2I.com.coralogixapis.notification_center.MatchEntityTypeAndSubTypeConditionH\x00R\x19matchEntityTypeAndSubTypeB\v\n" +
	"\tcondition\"b\n" +
	"\rMessageConfig\x12Q\n" +
	"\x06fields\x18\x01 \x03(\v29.com.coralogixapis.notification_center.MessageConfigFieldR\x06fields\";\n" +
	"\x18MatchEntityTypeCondition\x12\x1f\n" +
	"\ventity_type\x18\x01 \x01(\tR\n" +
	"entityType\"m\n" +
	"\"MatchEntityTypeAndSubTypeCondition\x12\x1f\n" +
	"\ventity_type\x18\x01 \x01(\tR\n" +
	"entityType\x12&\n" +
	"\x0fentity_sub_type\x18\x02 \x01(\tR\rentitySubTypeb\x06proto3"

var (
	file_com_coralogixapis_notification_center_common_config_overrides_proto_rawDescOnce sync.Once
	file_com_coralogixapis_notification_center_common_config_overrides_proto_rawDescData []byte
)

func file_com_coralogixapis_notification_center_common_config_overrides_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_notification_center_common_config_overrides_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_notification_center_common_config_overrides_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_notification_center_common_config_overrides_proto_rawDesc), len(file_com_coralogixapis_notification_center_common_config_overrides_proto_rawDesc)))
	})
	return file_com_coralogixapis_notification_center_common_config_overrides_proto_rawDescData
}

var file_com_coralogixapis_notification_center_common_config_overrides_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_com_coralogixapis_notification_center_common_config_overrides_proto_goTypes = []any{
	(*ConfigOverrides)(nil),                    // 0: com.coralogixapis.notification_center.ConfigOverrides
	(*ConditionType)(nil),                      // 1: com.coralogixapis.notification_center.ConditionType
	(*MessageConfig)(nil),                      // 2: com.coralogixapis.notification_center.MessageConfig
	(*MatchEntityTypeCondition)(nil),           // 3: com.coralogixapis.notification_center.MatchEntityTypeCondition
	(*MatchEntityTypeAndSubTypeCondition)(nil), // 4: com.coralogixapis.notification_center.MatchEntityTypeAndSubTypeCondition
	(*v1.MessageConfigField)(nil),              // 5: com.coralogixapis.notification_center.MessageConfigField
}
var file_com_coralogixapis_notification_center_common_config_overrides_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.notification_center.ConfigOverrides.condition_type:type_name -> com.coralogixapis.notification_center.ConditionType
	2, // 1: com.coralogixapis.notification_center.ConfigOverrides.message_config:type_name -> com.coralogixapis.notification_center.MessageConfig
	3, // 2: com.coralogixapis.notification_center.ConditionType.match_entity_type:type_name -> com.coralogixapis.notification_center.MatchEntityTypeCondition
	4, // 3: com.coralogixapis.notification_center.ConditionType.match_entity_type_and_sub_type:type_name -> com.coralogixapis.notification_center.MatchEntityTypeAndSubTypeCondition
	5, // 4: com.coralogixapis.notification_center.MessageConfig.fields:type_name -> com.coralogixapis.notification_center.MessageConfigField
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_notification_center_common_config_overrides_proto_init() }
func file_com_coralogixapis_notification_center_common_config_overrides_proto_init() {
	if File_com_coralogixapis_notification_center_common_config_overrides_proto != nil {
		return
	}
	file_com_coralogixapis_notification_center_common_config_overrides_proto_msgTypes[1].OneofWrappers = []any{
		(*ConditionType_MatchEntityType)(nil),
		(*ConditionType_MatchEntityTypeAndSubType)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_notification_center_common_config_overrides_proto_rawDesc), len(file_com_coralogixapis_notification_center_common_config_overrides_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_notification_center_common_config_overrides_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_notification_center_common_config_overrides_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_notification_center_common_config_overrides_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_notification_center_common_config_overrides_proto = out.File
	file_com_coralogixapis_notification_center_common_config_overrides_proto_goTypes = nil
	file_com_coralogixapis_notification_center_common_config_overrides_proto_depIdxs = nil
}
