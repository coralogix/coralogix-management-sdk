// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.26.1
// source: com/coralogixapis/notification_center/common/v1/routing.proto

package v1

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

type RoutingRule struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Condition     string                 `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition,omitempty"`
	Targets       []*RoutingTarget       `protobuf:"bytes,2,rep,name=targets,proto3" json:"targets,omitempty"`
	CustomDetails map[string]string      `protobuf:"bytes,3,rep,name=custom_details,json=customDetails,proto3" json:"custom_details,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Name          *string                `protobuf:"bytes,4,opt,name=name,proto3,oneof" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RoutingRule) Reset() {
	*x = RoutingRule{}
	mi := &file_com_coralogixapis_notification_center_common_v1_routing_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoutingRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoutingRule) ProtoMessage() {}

func (x *RoutingRule) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_common_v1_routing_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoutingRule.ProtoReflect.Descriptor instead.
func (*RoutingRule) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_common_v1_routing_proto_rawDescGZIP(), []int{0}
}

func (x *RoutingRule) GetCondition() string {
	if x != nil {
		return x.Condition
	}
	return ""
}

func (x *RoutingRule) GetTargets() []*RoutingTarget {
	if x != nil {
		return x.Targets
	}
	return nil
}

func (x *RoutingRule) GetCustomDetails() map[string]string {
	if x != nil {
		return x.CustomDetails
	}
	return nil
}

func (x *RoutingRule) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type RoutingTarget struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	ConnectorId     string                 `protobuf:"bytes,1,opt,name=connector_id,json=connectorId,proto3" json:"connector_id,omitempty"`
	PresetId        *string                `protobuf:"bytes,2,opt,name=preset_id,json=presetId,proto3,oneof" json:"preset_id,omitempty"`
	ConfigOverrides *SourceOverrides       `protobuf:"bytes,3,opt,name=config_overrides,json=configOverrides,proto3,oneof" json:"config_overrides,omitempty"`
	CustomDetails   map[string]string      `protobuf:"bytes,4,rep,name=custom_details,json=customDetails,proto3" json:"custom_details,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *RoutingTarget) Reset() {
	*x = RoutingTarget{}
	mi := &file_com_coralogixapis_notification_center_common_v1_routing_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoutingTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoutingTarget) ProtoMessage() {}

func (x *RoutingTarget) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_common_v1_routing_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoutingTarget.ProtoReflect.Descriptor instead.
func (*RoutingTarget) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_common_v1_routing_proto_rawDescGZIP(), []int{1}
}

func (x *RoutingTarget) GetConnectorId() string {
	if x != nil {
		return x.ConnectorId
	}
	return ""
}

func (x *RoutingTarget) GetPresetId() string {
	if x != nil && x.PresetId != nil {
		return *x.PresetId
	}
	return ""
}

func (x *RoutingTarget) GetConfigOverrides() *SourceOverrides {
	if x != nil {
		return x.ConfigOverrides
	}
	return nil
}

func (x *RoutingTarget) GetCustomDetails() map[string]string {
	if x != nil {
		return x.CustomDetails
	}
	return nil
}

type SourceOverrides struct {
	state                 protoimpl.MessageState  `protogen:"open.v1"`
	OutputSchemaId        string                  `protobuf:"bytes,1,opt,name=output_schema_id,json=outputSchemaId,proto3" json:"output_schema_id,omitempty"`
	MessageConfigFields   []*MessageConfigField   `protobuf:"bytes,3,rep,name=message_config_fields,json=messageConfigFields,proto3" json:"message_config_fields,omitempty"`
	ConnectorConfigFields []*ConnectorConfigField `protobuf:"bytes,4,rep,name=connector_config_fields,json=connectorConfigFields,proto3" json:"connector_config_fields,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *SourceOverrides) Reset() {
	*x = SourceOverrides{}
	mi := &file_com_coralogixapis_notification_center_common_v1_routing_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SourceOverrides) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceOverrides) ProtoMessage() {}

func (x *SourceOverrides) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_common_v1_routing_proto_msgTypes[2]
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
	return file_com_coralogixapis_notification_center_common_v1_routing_proto_rawDescGZIP(), []int{2}
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

type GlobalRouterIdentifier struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Value:
	//
	//	*GlobalRouterIdentifier_Id
	//	*GlobalRouterIdentifier_UserFacingId
	Value         isGlobalRouterIdentifier_Value `protobuf_oneof:"value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GlobalRouterIdentifier) Reset() {
	*x = GlobalRouterIdentifier{}
	mi := &file_com_coralogixapis_notification_center_common_v1_routing_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GlobalRouterIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobalRouterIdentifier) ProtoMessage() {}

func (x *GlobalRouterIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_common_v1_routing_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobalRouterIdentifier.ProtoReflect.Descriptor instead.
func (*GlobalRouterIdentifier) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_common_v1_routing_proto_rawDescGZIP(), []int{3}
}

func (x *GlobalRouterIdentifier) GetValue() isGlobalRouterIdentifier_Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *GlobalRouterIdentifier) GetId() string {
	if x != nil {
		if x, ok := x.Value.(*GlobalRouterIdentifier_Id); ok {
			return x.Id
		}
	}
	return ""
}

func (x *GlobalRouterIdentifier) GetUserFacingId() string {
	if x != nil {
		if x, ok := x.Value.(*GlobalRouterIdentifier_UserFacingId); ok {
			return x.UserFacingId
		}
	}
	return ""
}

type isGlobalRouterIdentifier_Value interface {
	isGlobalRouterIdentifier_Value()
}

type GlobalRouterIdentifier_Id struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type GlobalRouterIdentifier_UserFacingId struct {
	UserFacingId string `protobuf:"bytes,2,opt,name=user_facing_id,json=userFacingId,proto3,oneof"`
}

func (*GlobalRouterIdentifier_Id) isGlobalRouterIdentifier_Value() {}

func (*GlobalRouterIdentifier_UserFacingId) isGlobalRouterIdentifier_Value() {}

var File_com_coralogixapis_notification_center_common_v1_routing_proto protoreflect.FileDescriptor

var file_com_coralogixapis_notification_center_common_v1_routing_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x2d, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x43,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x02, 0x0a, 0x0b, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x52,
	0x75, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x56, 0x0a, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x52, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x74, 0x0a, 0x0e, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x4d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x1a, 0x40, 0x0a, 0x12, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0xa1, 0x03, 0x0a, 0x0d, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x73,
	0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x6e, 0x0a, 0x10, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x72,
	0x69, 0x64, 0x65, 0x73, 0x48, 0x01, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f, 0x76,
	0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x76, 0x0a, 0x0e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x4f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x1a, 0x40, 0x0a, 0x12, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f,
	0x69, 0x64, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6f, 0x76,
	0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x22, 0x9f, 0x02, 0x0a, 0x0f, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x49, 0x64, 0x12, 0x6d, 0x0a, 0x15, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52,
	0x13, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x12, 0x73, 0x0a, 0x17, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x52, 0x15, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x5b, 0x0a, 0x16, 0x47, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x66, 0x61,
	0x63, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0c, 0x75, 0x73, 0x65, 0x72, 0x46, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x42, 0x07, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_notification_center_common_v1_routing_proto_rawDescOnce sync.Once
	file_com_coralogixapis_notification_center_common_v1_routing_proto_rawDescData = file_com_coralogixapis_notification_center_common_v1_routing_proto_rawDesc
)

func file_com_coralogixapis_notification_center_common_v1_routing_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_notification_center_common_v1_routing_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_notification_center_common_v1_routing_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_notification_center_common_v1_routing_proto_rawDescData)
	})
	return file_com_coralogixapis_notification_center_common_v1_routing_proto_rawDescData
}

var file_com_coralogixapis_notification_center_common_v1_routing_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_com_coralogixapis_notification_center_common_v1_routing_proto_goTypes = []any{
	(*RoutingRule)(nil),            // 0: com.coralogixapis.notification_center.routing.RoutingRule
	(*RoutingTarget)(nil),          // 1: com.coralogixapis.notification_center.routing.RoutingTarget
	(*SourceOverrides)(nil),        // 2: com.coralogixapis.notification_center.routing.SourceOverrides
	(*GlobalRouterIdentifier)(nil), // 3: com.coralogixapis.notification_center.routing.GlobalRouterIdentifier
	nil,                            // 4: com.coralogixapis.notification_center.routing.RoutingRule.CustomDetailsEntry
	nil,                            // 5: com.coralogixapis.notification_center.routing.RoutingTarget.CustomDetailsEntry
	(*MessageConfigField)(nil),     // 6: com.coralogixapis.notification_center.MessageConfigField
	(*ConnectorConfigField)(nil),   // 7: com.coralogixapis.notification_center.ConnectorConfigField
}
var file_com_coralogixapis_notification_center_common_v1_routing_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.notification_center.routing.RoutingRule.targets:type_name -> com.coralogixapis.notification_center.routing.RoutingTarget
	4, // 1: com.coralogixapis.notification_center.routing.RoutingRule.custom_details:type_name -> com.coralogixapis.notification_center.routing.RoutingRule.CustomDetailsEntry
	2, // 2: com.coralogixapis.notification_center.routing.RoutingTarget.config_overrides:type_name -> com.coralogixapis.notification_center.routing.SourceOverrides
	5, // 3: com.coralogixapis.notification_center.routing.RoutingTarget.custom_details:type_name -> com.coralogixapis.notification_center.routing.RoutingTarget.CustomDetailsEntry
	6, // 4: com.coralogixapis.notification_center.routing.SourceOverrides.message_config_fields:type_name -> com.coralogixapis.notification_center.MessageConfigField
	7, // 5: com.coralogixapis.notification_center.routing.SourceOverrides.connector_config_fields:type_name -> com.coralogixapis.notification_center.ConnectorConfigField
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_notification_center_common_v1_routing_proto_init() }
func file_com_coralogixapis_notification_center_common_v1_routing_proto_init() {
	if File_com_coralogixapis_notification_center_common_v1_routing_proto != nil {
		return
	}
	file_com_coralogixapis_notification_center_common_v1_config_fields_proto_init()
	file_com_coralogixapis_notification_center_common_v1_routing_proto_msgTypes[0].OneofWrappers = []any{}
	file_com_coralogixapis_notification_center_common_v1_routing_proto_msgTypes[1].OneofWrappers = []any{}
	file_com_coralogixapis_notification_center_common_v1_routing_proto_msgTypes[3].OneofWrappers = []any{
		(*GlobalRouterIdentifier_Id)(nil),
		(*GlobalRouterIdentifier_UserFacingId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_notification_center_common_v1_routing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_notification_center_common_v1_routing_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_notification_center_common_v1_routing_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_notification_center_common_v1_routing_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_notification_center_common_v1_routing_proto = out.File
	file_com_coralogixapis_notification_center_common_v1_routing_proto_rawDesc = nil
	file_com_coralogixapis_notification_center_common_v1_routing_proto_goTypes = nil
	file_com_coralogixapis_notification_center_common_v1_routing_proto_depIdxs = nil
}
