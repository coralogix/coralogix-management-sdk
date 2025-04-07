// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.26.1
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
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Connector struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// A unique identifier automatically generated by the service (do not provide this from the client side)
	Id *string `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// A unique identifier provided by the user
	UserDefinedId *string              `protobuf:"bytes,2,opt,name=user_defined_id,json=userDefinedId,proto3,oneof" json:"user_defined_id,omitempty"`
	Type          common.ConnectorType `protobuf:"varint,3,opt,name=type,proto3,enum=com.coralogixapis.notification_center.ConnectorType" json:"type,omitempty"`
	TeamId        *uint32              `protobuf:"varint,4,opt,name=team_id,json=teamId,proto3,oneof" json:"team_id,omitempty"`
	Name          string               `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Description   string               `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// System-generated timestamp for when the connector was last updated
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3,oneof" json:"create_time,omitempty"`
	// System-generated timestamp for when the connector was last updated
	UpdateTime       *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=update_time,json=updateTime,proto3,oneof" json:"update_time,omitempty"`
	ConnectorConfigs []*ConnectorConfig     `protobuf:"bytes,9,rep,name=connector_configs,json=connectorConfigs,proto3" json:"connector_configs,omitempty"`
	// Configuration override templates for specific entity types, values from connector_configs will be used if not overridden
	ConfigOverrides []*EntityTypeConfigOverrides `protobuf:"bytes,11,rep,name=config_overrides,json=configOverrides,proto3" json:"config_overrides,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
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

func (x *Connector) GetUserDefinedId() string {
	if x != nil && x.UserDefinedId != nil {
		return *x.UserDefinedId
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
	state            protoimpl.MessageState `protogen:"open.v1"`
	EntityType       string                 `protobuf:"bytes,1,opt,name=entity_type,json=entityType,proto3" json:"entity_type,omitempty"`
	ConnectorConfigs []*ConnectorConfig     `protobuf:"bytes,2,rep,name=connector_configs,json=connectorConfigs,proto3" json:"connector_configs,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
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
	state protoimpl.MessageState `protogen:"open.v1"`
	Id    *string                `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// for declarative API
	UserDefinedId *string                `protobuf:"bytes,2,opt,name=user_defined_id,json=userDefinedId,proto3,oneof" json:"user_defined_id,omitempty"`
	Type          common.ConnectorType   `protobuf:"varint,3,opt,name=type,proto3,enum=com.coralogixapis.notification_center.ConnectorType" json:"type,omitempty"`
	TeamId        *uint32                `protobuf:"varint,4,opt,name=team_id,json=teamId,proto3,oneof" json:"team_id,omitempty"`
	Name          string                 `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	CreateTime    *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3,oneof" json:"create_time,omitempty"`
	UpdateTime    *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=update_time,json=updateTime,proto3,oneof" json:"update_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

func (x *ConnectorSummary) GetUserDefinedId() string {
	if x != nil && x.UserDefinedId != nil {
		return *x.UserDefinedId
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
	state          protoimpl.MessageState     `protogen:"open.v1"`
	OutputSchemaId string                     `protobuf:"bytes,1,opt,name=output_schema_id,json=outputSchemaId,proto3" json:"output_schema_id,omitempty"`
	Fields         []*v1.ConnectorConfigField `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          common.ConnectorType   `protobuf:"varint,1,opt,name=type,proto3,enum=com.coralogixapis.notification_center.ConnectorType" json:"type,omitempty"`
	Count         uint32                 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

const file_com_coralogixapis_notification_center_connectors_v1_connector_proto_rawDesc = "" +
	"\n" +
	"Ccom/coralogixapis/notification_center/connectors/v1/connector.proto\x123com.coralogixapis.notification_center.connectors.v1\x1a9com/coralogixapis/notification_center/common/common.proto\x1aCcom/coralogixapis/notification_center/common/v1/config_fields.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\xbe\b\n" +
	"\tConnector\x12@\n" +
	"\x02id\x18\x01 \x01(\tB+\x92A(J&\"a16e24c8-4db2-4abf-ba3c-c9e1fc35a3b9\"H\x00R\x02id\x88\x01\x01\x12E\n" +
	"\x0fuser_defined_id\x18\x02 \x01(\tB\x18\x92A\x15J\x13\"user-connector-id\"H\x01R\ruserDefinedId\x88\x01\x01\x12H\n" +
	"\x04type\x18\x03 \x01(\x0e24.com.coralogixapis.notification_center.ConnectorTypeR\x04type\x12*\n" +
	"\ateam_id\x18\x04 \x01(\rB\f\x92A\tJ\a\"12345\"H\x02R\x06teamId\x88\x01\x01\x120\n" +
	"\x04name\x18\x05 \x01(\tB\x1c\x92A\x19J\x14\"My Slack Connector\"x\xc8\x01R\x04name\x12L\n" +
	"\vdescription\x18\x06 \x01(\tB*\x92A'J\"\"Connector for team notifications\"x\x88'R\vdescription\x12@\n" +
	"\vcreate_time\x18\a \x01(\v2\x1a.google.protobuf.TimestampH\x03R\n" +
	"createTime\x88\x01\x01\x12@\n" +
	"\vupdate_time\x18\b \x01(\v2\x1a.google.protobuf.TimestampH\x04R\n" +
	"updateTime\x88\x01\x01\x12q\n" +
	"\x11connector_configs\x18\t \x03(\v2D.com.coralogixapis.notification_center.connectors.v1.ConnectorConfigR\x10connectorConfigs\x12y\n" +
	"\x10config_overrides\x18\v \x03(\v2N.com.coralogixapis.notification_center.connectors.v1.EntityTypeConfigOverridesR\x0fconfigOverrides:\xe5\x01\x92A\xe1\x01\n" +
	"b*\tConnector23A connector configuration for sending notifications\xd2\x01\x04type\xd2\x01\x04name\xd2\x01\x11connector_configs*{\n" +
	"'Find out more about notification center\x12Phttps://coralogix.com/docs/user-guides/notification-center/introduction/welcome/B\x05\n" +
	"\x03_idB\x12\n" +
	"\x10_user_defined_idB\n" +
	"\n" +
	"\b_team_idB\x0e\n" +
	"\f_create_timeB\x0e\n" +
	"\f_update_timeJ\x04\b\n" +
	"\x10\vR\ventity_type\"\xbe\x01\n" +
	"\x19EntityTypeConfigOverrides\x12.\n" +
	"\ventity_type\x18\x01 \x01(\tB\r\x92A\n" +
	"J\b\"alerts\"R\n" +
	"entityType\x12q\n" +
	"\x11connector_configs\x18\x02 \x03(\v2D.com.coralogixapis.notification_center.connectors.v1.ConnectorConfigR\x10connectorConfigs\"\xd0\x03\n" +
	"\x10ConnectorSummary\x12\x13\n" +
	"\x02id\x18\x01 \x01(\tH\x00R\x02id\x88\x01\x01\x12+\n" +
	"\x0fuser_defined_id\x18\x02 \x01(\tH\x01R\ruserDefinedId\x88\x01\x01\x12H\n" +
	"\x04type\x18\x03 \x01(\x0e24.com.coralogixapis.notification_center.ConnectorTypeR\x04type\x12\x1c\n" +
	"\ateam_id\x18\x04 \x01(\rH\x02R\x06teamId\x88\x01\x01\x12\x12\n" +
	"\x04name\x18\x05 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x06 \x01(\tR\vdescription\x12@\n" +
	"\vcreate_time\x18\a \x01(\v2\x1a.google.protobuf.TimestampH\x03R\n" +
	"createTime\x88\x01\x01\x12@\n" +
	"\vupdate_time\x18\b \x01(\v2\x1a.google.protobuf.TimestampH\x04R\n" +
	"updateTime\x88\x01\x01B\x05\n" +
	"\x03_idB\x12\n" +
	"\x10_user_defined_idB\n" +
	"\n" +
	"\b_team_idB\x0e\n" +
	"\f_create_timeB\x0e\n" +
	"\f_update_timeJ\x04\b\t\x10\n" +
	"R\ventity_type\"\xf5\x02\n" +
	"\x0fConnectorConfig\x12:\n" +
	"\x10output_schema_id\x18\x01 \x01(\tB\x10\x92A\rJ\v\"schema-id\"R\x0eoutputSchemaId\x12S\n" +
	"\x06fields\x18\x02 \x03(\v2;.com.coralogixapis.notification_center.ConnectorConfigFieldR\x06fields:\xd0\x01\x92A\xcc\x01\n" +
	"M*\x10Connector Config29Configuration for a specific output schema of a connector*{\n" +
	"'Find out more about notification center\x12Phttps://coralogix.com/docs/user-guides/notification-center/introduction/welcome/\"\xca\x02\n" +
	"\x14ConnectorTypeSummary\x12H\n" +
	"\x04type\x18\x01 \x01(\x0e24.com.coralogixapis.notification_center.ConnectorTypeR\x04type\x12\x1e\n" +
	"\x05count\x18\x02 \x01(\rB\b\x92A\x05J\x03\"5\"R\x05count:\xc7\x01\x92A\xc3\x01\n" +
	"D*\x16Connector Type Summary2*Summary information about a connector type*{\n" +
	"'Find out more about notification center\x12Phttps://coralogix.com/docs/user-guides/notification-center/introduction/welcome/b\x06proto3"

var (
	file_com_coralogixapis_notification_center_connectors_v1_connector_proto_rawDescOnce sync.Once
	file_com_coralogixapis_notification_center_connectors_v1_connector_proto_rawDescData []byte
)

func file_com_coralogixapis_notification_center_connectors_v1_connector_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_notification_center_connectors_v1_connector_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_notification_center_connectors_v1_connector_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_notification_center_connectors_v1_connector_proto_rawDesc), len(file_com_coralogixapis_notification_center_connectors_v1_connector_proto_rawDesc)))
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_notification_center_connectors_v1_connector_proto_rawDesc), len(file_com_coralogixapis_notification_center_connectors_v1_connector_proto_rawDesc)),
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
	file_com_coralogixapis_notification_center_connectors_v1_connector_proto_goTypes = nil
	file_com_coralogixapis_notification_center_connectors_v1_connector_proto_depIdxs = nil
}
