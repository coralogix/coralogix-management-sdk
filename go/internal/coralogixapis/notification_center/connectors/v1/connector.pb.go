// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: com/coralogixapis/notification_center/connectors/v1/connector.proto

package v1

import (
	common "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/notification_center/common"
	v1 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/notification_center/common/v1"
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

type Connector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique identifier automatically generated by the service (do not provide this from the client side)
	Id *string `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// A unique identifier provided by the user
	UserFacingId *string              `protobuf:"bytes,2,opt,name=user_facing_id,json=userFacingId,proto3,oneof" json:"user_facing_id,omitempty"`
	Type         common.ConnectorType `protobuf:"varint,3,opt,name=type,proto3,enum=com.coralogixapis.notification_center.ConnectorType" json:"type,omitempty"`
	TeamId       *uint32              `protobuf:"varint,4,opt,name=team_id,json=teamId,proto3,oneof" json:"team_id,omitempty"`
	Name         string               `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Description  string               `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// System-generated timestamp for when the connector was last updated
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3,oneof" json:"create_time,omitempty"`
	// System-generated timestamp for when the connector was last updated
	UpdateTime       *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=update_time,json=updateTime,proto3,oneof" json:"update_time,omitempty"`
	ConnectorConfigs []*ConnectorConfig     `protobuf:"bytes,9,rep,name=connector_configs,json=connectorConfigs,proto3" json:"connector_configs,omitempty"`
	// Configuration override templates for specific entity types, values from connector_configs will be used if not overridden
	ConfigOverrides []*EntityTypeConfigOverrides `protobuf:"bytes,11,rep,name=config_overrides,json=configOverrides,proto3" json:"config_overrides,omitempty"`
}

func (x *Connector) Reset() {
	*x = Connector{}
	mi := &file_com_coralogixapis_notification_center_connectors_v1_connector_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Connector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connector) ProtoMessage() {}

func (x *Connector) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_connectors_v1_connector_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connector.ProtoReflect.Descriptor instead.
func (*Connector) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_connectors_v1_connector_proto_rawDescGZIP(), []int{0}
}

func (x *Connector) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *Connector) GetUserFacingId() string {
	if x != nil && x.UserFacingId != nil {
		return *x.UserFacingId
	}
	return ""
}

func (x *Connector) GetType() common.ConnectorType {
	if x != nil {
		return x.Type
	}
	return common.ConnectorType(0)
}

func (x *Connector) GetTeamId() uint32 {
	if x != nil && x.TeamId != nil {
		return *x.TeamId
	}
	return 0
}

func (x *Connector) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Connector) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Connector) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Connector) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Connector) GetConnectorConfigs() []*ConnectorConfig {
	if x != nil {
		return x.ConnectorConfigs
	}
	return nil
}

func (x *Connector) GetConfigOverrides() []*EntityTypeConfigOverrides {
	if x != nil {
		return x.ConfigOverrides
	}
	return nil
}

type EntityTypeConfigOverrides struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityType       string             `protobuf:"bytes,1,opt,name=entity_type,json=entityType,proto3" json:"entity_type,omitempty"`
	ConnectorConfigs []*ConnectorConfig `protobuf:"bytes,2,rep,name=connector_configs,json=connectorConfigs,proto3" json:"connector_configs,omitempty"`
}

func (x *EntityTypeConfigOverrides) Reset() {
	*x = EntityTypeConfigOverrides{}
	mi := &file_com_coralogixapis_notification_center_connectors_v1_connector_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntityTypeConfigOverrides) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityTypeConfigOverrides) ProtoMessage() {}

func (x *EntityTypeConfigOverrides) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_connectors_v1_connector_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityTypeConfigOverrides.ProtoReflect.Descriptor instead.
func (*EntityTypeConfigOverrides) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_connectors_v1_connector_proto_rawDescGZIP(), []int{1}
}

func (x *EntityTypeConfigOverrides) GetEntityType() string {
	if x != nil {
		return x.EntityType
	}
	return ""
}

func (x *EntityTypeConfigOverrides) GetConnectorConfigs() []*ConnectorConfig {
	if x != nil {
		return x.ConnectorConfigs
	}
	return nil
}

type ConnectorSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *string `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// for declarative API
	UserFacingId *string                `protobuf:"bytes,2,opt,name=user_facing_id,json=userFacingId,proto3,oneof" json:"user_facing_id,omitempty"`
	Type         common.ConnectorType   `protobuf:"varint,3,opt,name=type,proto3,enum=com.coralogixapis.notification_center.ConnectorType" json:"type,omitempty"`
	TeamId       *uint32                `protobuf:"varint,4,opt,name=team_id,json=teamId,proto3,oneof" json:"team_id,omitempty"`
	Name         string                 `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Description  string                 `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	CreateTime   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3,oneof" json:"create_time,omitempty"`
	UpdateTime   *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=update_time,json=updateTime,proto3,oneof" json:"update_time,omitempty"`
}

func (x *ConnectorSummary) Reset() {
	*x = ConnectorSummary{}
	mi := &file_com_coralogixapis_notification_center_connectors_v1_connector_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectorSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectorSummary) ProtoMessage() {}

func (x *ConnectorSummary) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_connectors_v1_connector_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectorSummary.ProtoReflect.Descriptor instead.
func (*ConnectorSummary) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_connectors_v1_connector_proto_rawDescGZIP(), []int{2}
}

func (x *ConnectorSummary) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *ConnectorSummary) GetUserFacingId() string {
	if x != nil && x.UserFacingId != nil {
		return *x.UserFacingId
	}
	return ""
}

func (x *ConnectorSummary) GetType() common.ConnectorType {
	if x != nil {
		return x.Type
	}
	return common.ConnectorType(0)
}

func (x *ConnectorSummary) GetTeamId() uint32 {
	if x != nil && x.TeamId != nil {
		return *x.TeamId
	}
	return 0
}

func (x *ConnectorSummary) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConnectorSummary) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ConnectorSummary) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *ConnectorSummary) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

type ConnectorConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutputSchemaId string                     `protobuf:"bytes,1,opt,name=output_schema_id,json=outputSchemaId,proto3" json:"output_schema_id,omitempty"`
	Fields         []*v1.ConnectorConfigField `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *ConnectorConfig) Reset() {
	*x = ConnectorConfig{}
	mi := &file_com_coralogixapis_notification_center_connectors_v1_connector_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectorConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectorConfig) ProtoMessage() {}

func (x *ConnectorConfig) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_connectors_v1_connector_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectorConfig.ProtoReflect.Descriptor instead.
func (*ConnectorConfig) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_connectors_v1_connector_proto_rawDescGZIP(), []int{3}
}

func (x *ConnectorConfig) GetOutputSchemaId() string {
	if x != nil {
		return x.OutputSchemaId
	}
	return ""
}

func (x *ConnectorConfig) GetFields() []*v1.ConnectorConfigField {
	if x != nil {
		return x.Fields
	}
	return nil
}

type ConnectorTypeSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  common.ConnectorType `protobuf:"varint,1,opt,name=type,proto3,enum=com.coralogixapis.notification_center.ConnectorType" json:"type,omitempty"`
	Count uint32               `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ConnectorTypeSummary) Reset() {
	*x = ConnectorTypeSummary{}
	mi := &file_com_coralogixapis_notification_center_connectors_v1_connector_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectorTypeSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectorTypeSummary) ProtoMessage() {}

func (x *ConnectorTypeSummary) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_connectors_v1_connector_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectorTypeSummary.ProtoReflect.Descriptor instead.
func (*ConnectorTypeSummary) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_connectors_v1_connector_proto_rawDescGZIP(), []int{4}
}

func (x *ConnectorTypeSummary) GetType() common.ConnectorType {
	if x != nil {
		return x.Type
	}
	return common.ConnectorType(0)
}

func (x *ConnectorTypeSummary) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_com_coralogixapis_notification_center_connectors_v1_connector_proto protoreflect.FileDescriptor

var file_com_coralogixapis_notification_center_connectors_v1_connector_proto_rawDesc = []byte{
	0x0a, 0x43, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x33, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x39, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x43, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76,
	0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x08, 0x0a, 0x09,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x40, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0x92, 0x41, 0x28, 0x4a, 0x26, 0x22, 0x61, 0x31, 0x36,
	0x65, 0x32, 0x34, 0x63, 0x38, 0x2d, 0x34, 0x64, 0x62, 0x32, 0x2d, 0x34, 0x61, 0x62, 0x66, 0x2d,
	0x62, 0x61, 0x33, 0x63, 0x2d, 0x63, 0x39, 0x65, 0x31, 0x66, 0x63, 0x33, 0x35, 0x61, 0x33, 0x62,
	0x39, 0x22, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x43, 0x0a, 0x0e, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x66, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x18, 0x92, 0x41, 0x15, 0x4a, 0x13, 0x22, 0x75, 0x73, 0x65, 0x72, 0x2d,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2d, 0x69, 0x64, 0x22, 0x48, 0x01, 0x52,
	0x0c, 0x75, 0x73, 0x65, 0x72, 0x46, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x48, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x74, 0x65,
	0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0c, 0x92, 0x41, 0x09,
	0x4a, 0x07, 0x22, 0x31, 0x32, 0x33, 0x34, 0x35, 0x22, 0x48, 0x02, 0x52, 0x06, 0x74, 0x65, 0x61,
	0x6d, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x1c, 0x92, 0x41, 0x19, 0x4a, 0x14, 0x22, 0x4d, 0x79, 0x20, 0x53,
	0x6c, 0x61, 0x63, 0x6b, 0x20, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x78,
	0xc8, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2a, 0x92,
	0x41, 0x27, 0x4a, 0x22, 0x22, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x74, 0x65, 0x61, 0x6d, 0x20, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x78, 0x88, 0x27, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x04, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x71, 0x0a, 0x11, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x10, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x79, 0x0a,
	0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65,
	0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f, 0x76,
	0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f,
	0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x3a, 0xe5, 0x01, 0x92, 0x41, 0xe1, 0x01, 0x0a,
	0x62, 0x2a, 0x09, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x32, 0x33, 0x41, 0x20,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x20, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x73, 0x65, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x20, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0xd2, 0x01, 0x04, 0x74, 0x79, 0x70, 0x65, 0xd2, 0x01, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xd2,
	0x01, 0x11, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x73, 0x2a, 0x7b, 0x0a, 0x27, 0x46, 0x69, 0x6e, 0x64, 0x20, 0x6f, 0x75, 0x74, 0x20,
	0x6d, 0x6f, 0x72, 0x65, 0x20, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x20, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x50,
	0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x63, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d,
	0x67, 0x75, 0x69, 0x64, 0x65, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x77, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x2f,
	0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x66, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x74,
	0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x4a, 0x04, 0x08, 0x0a, 0x10, 0x0b, 0x52, 0x0b, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0xbe, 0x01, 0x0a, 0x19, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f, 0x76,
	0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0x92, 0x41,
	0x0a, 0x4a, 0x08, 0x22, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x22, 0x52, 0x0a, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x71, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x44, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x22, 0xcd, 0x03, 0x0a, 0x10, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12,
	0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x66, 0x61, 0x63,
	0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0c,
	0x75, 0x73, 0x65, 0x72, 0x46, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x48, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x07, 0x74, 0x65, 0x61,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x02, 0x52, 0x06, 0x74, 0x65,
	0x61, 0x6d, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x03,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x48, 0x04, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x66, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x4a, 0x04, 0x08, 0x09, 0x10, 0x0a, 0x52, 0x0b, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0xf5, 0x02, 0x0a, 0x0f, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3a,
	0x0a, 0x10, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0x92, 0x41, 0x0d, 0x4a, 0x0b, 0x22,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2d, 0x69, 0x64, 0x22, 0x52, 0x0e, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64, 0x12, 0x53, 0x0a, 0x06, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x3a,
	0xd0, 0x01, 0x92, 0x41, 0xcc, 0x01, 0x0a, 0x4d, 0x2a, 0x10, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x20, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x32, 0x39, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x20,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x20,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x20, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x2a, 0x7b, 0x0a, 0x27, 0x46, 0x69, 0x6e, 0x64, 0x20, 0x6f, 0x75,
	0x74, 0x20, 0x6d, 0x6f, 0x72, 0x65, 0x20, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x20, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x12, 0x50, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x63, 0x73, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2d, 0x67, 0x75, 0x69, 0x64, 0x65, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x77, 0x65, 0x6c, 0x63, 0x6f, 0x6d,
	0x65, 0x2f, 0x22, 0xca, 0x02, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x48, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x42, 0x08, 0x92, 0x41, 0x05, 0x4a, 0x03, 0x22, 0x35, 0x22, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0xc7, 0x01, 0x92, 0x41, 0xc3, 0x01, 0x0a, 0x44, 0x2a, 0x16,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x20, 0x54, 0x79, 0x70, 0x65, 0x20, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x32, 0x2a, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x20,
	0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x61, 0x62, 0x6f, 0x75,
	0x74, 0x20, 0x61, 0x20, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x20, 0x74, 0x79,
	0x70, 0x65, 0x2a, 0x7b, 0x0a, 0x27, 0x46, 0x69, 0x6e, 0x64, 0x20, 0x6f, 0x75, 0x74, 0x20, 0x6d,
	0x6f, 0x72, 0x65, 0x20, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x20, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x50, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x63, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x67,
	0x75, 0x69, 0x64, 0x65, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x77, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x2f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_notification_center_connectors_v1_connector_proto_rawDescOnce sync.Once
	file_com_coralogixapis_notification_center_connectors_v1_connector_proto_rawDescData = file_com_coralogixapis_notification_center_connectors_v1_connector_proto_rawDesc
)

func file_com_coralogixapis_notification_center_connectors_v1_connector_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_notification_center_connectors_v1_connector_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_notification_center_connectors_v1_connector_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_notification_center_connectors_v1_connector_proto_rawDescData)
	})
	return file_com_coralogixapis_notification_center_connectors_v1_connector_proto_rawDescData
}

var file_com_coralogixapis_notification_center_connectors_v1_connector_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_com_coralogixapis_notification_center_connectors_v1_connector_proto_goTypes = []any{
	(*Connector)(nil),                 // 0: com.coralogixapis.notification_center.connectors.v1.Connector
	(*EntityTypeConfigOverrides)(nil), // 1: com.coralogixapis.notification_center.connectors.v1.EntityTypeConfigOverrides
	(*ConnectorSummary)(nil),          // 2: com.coralogixapis.notification_center.connectors.v1.ConnectorSummary
	(*ConnectorConfig)(nil),           // 3: com.coralogixapis.notification_center.connectors.v1.ConnectorConfig
	(*ConnectorTypeSummary)(nil),      // 4: com.coralogixapis.notification_center.connectors.v1.ConnectorTypeSummary
	(common.ConnectorType)(0),         // 5: com.coralogixapis.notification_center.ConnectorType
	(*timestamppb.Timestamp)(nil),     // 6: google.protobuf.Timestamp
	(*v1.ConnectorConfigField)(nil),   // 7: com.coralogixapis.notification_center.ConnectorConfigField
}
var file_com_coralogixapis_notification_center_connectors_v1_connector_proto_depIdxs = []int32{
	5,  // 0: com.coralogixapis.notification_center.connectors.v1.Connector.type:type_name -> com.coralogixapis.notification_center.ConnectorType
	6,  // 1: com.coralogixapis.notification_center.connectors.v1.Connector.create_time:type_name -> google.protobuf.Timestamp
	6,  // 2: com.coralogixapis.notification_center.connectors.v1.Connector.update_time:type_name -> google.protobuf.Timestamp
	3,  // 3: com.coralogixapis.notification_center.connectors.v1.Connector.connector_configs:type_name -> com.coralogixapis.notification_center.connectors.v1.ConnectorConfig
	1,  // 4: com.coralogixapis.notification_center.connectors.v1.Connector.config_overrides:type_name -> com.coralogixapis.notification_center.connectors.v1.EntityTypeConfigOverrides
	3,  // 5: com.coralogixapis.notification_center.connectors.v1.EntityTypeConfigOverrides.connector_configs:type_name -> com.coralogixapis.notification_center.connectors.v1.ConnectorConfig
	5,  // 6: com.coralogixapis.notification_center.connectors.v1.ConnectorSummary.type:type_name -> com.coralogixapis.notification_center.ConnectorType
	6,  // 7: com.coralogixapis.notification_center.connectors.v1.ConnectorSummary.create_time:type_name -> google.protobuf.Timestamp
	6,  // 8: com.coralogixapis.notification_center.connectors.v1.ConnectorSummary.update_time:type_name -> google.protobuf.Timestamp
	7,  // 9: com.coralogixapis.notification_center.connectors.v1.ConnectorConfig.fields:type_name -> com.coralogixapis.notification_center.ConnectorConfigField
	5,  // 10: com.coralogixapis.notification_center.connectors.v1.ConnectorTypeSummary.type:type_name -> com.coralogixapis.notification_center.ConnectorType
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_notification_center_connectors_v1_connector_proto_init() }
func file_com_coralogixapis_notification_center_connectors_v1_connector_proto_init() {
	if File_com_coralogixapis_notification_center_connectors_v1_connector_proto != nil {
		return
	}
	file_com_coralogixapis_notification_center_connectors_v1_connector_proto_msgTypes[0].OneofWrappers = []any{}
	file_com_coralogixapis_notification_center_connectors_v1_connector_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_notification_center_connectors_v1_connector_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_notification_center_connectors_v1_connector_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_notification_center_connectors_v1_connector_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_notification_center_connectors_v1_connector_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_notification_center_connectors_v1_connector_proto = out.File
	file_com_coralogixapis_notification_center_connectors_v1_connector_proto_rawDesc = nil
	file_com_coralogixapis_notification_center_connectors_v1_connector_proto_goTypes = nil
	file_com_coralogixapis_notification_center_connectors_v1_connector_proto_depIdxs = nil
}
