// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogix/archive/dataset/v2/schema_service.proto

package v2

import (
	_ "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/common/v1"
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

type SetNamedLogsSchemaRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Schema        *Schema                `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
	UpdateSchema  bool                   `protobuf:"varint,2,opt,name=update_schema,json=updateSchema,proto3" json:"update_schema,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetNamedLogsSchemaRequest) Reset() {
	*x = SetNamedLogsSchemaRequest{}
	mi := &file_com_coralogix_archive_dataset_v2_schema_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetNamedLogsSchemaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetNamedLogsSchemaRequest) ProtoMessage() {}

func (x *SetNamedLogsSchemaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_dataset_v2_schema_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetNamedLogsSchemaRequest.ProtoReflect.Descriptor instead.
func (*SetNamedLogsSchemaRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDescGZIP(), []int{0}
}

func (x *SetNamedLogsSchemaRequest) GetSchema() *Schema {
	if x != nil {
		return x.Schema
	}
	return nil
}

func (x *SetNamedLogsSchemaRequest) GetUpdateSchema() bool {
	if x != nil {
		return x.UpdateSchema
	}
	return false
}

type SetNamedLogsSchemaResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Schema        *NamedSchema           `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetNamedLogsSchemaResponse) Reset() {
	*x = SetNamedLogsSchemaResponse{}
	mi := &file_com_coralogix_archive_dataset_v2_schema_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetNamedLogsSchemaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetNamedLogsSchemaResponse) ProtoMessage() {}

func (x *SetNamedLogsSchemaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_dataset_v2_schema_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetNamedLogsSchemaResponse.ProtoReflect.Descriptor instead.
func (*SetNamedLogsSchemaResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDescGZIP(), []int{1}
}

func (x *SetNamedLogsSchemaResponse) GetSchema() *NamedSchema {
	if x != nil {
		return x.Schema
	}
	return nil
}

type Schema struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Fields        []*NamedSchemaField    `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Schema) Reset() {
	*x = Schema{}
	mi := &file_com_coralogix_archive_dataset_v2_schema_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Schema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schema) ProtoMessage() {}

func (x *Schema) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_dataset_v2_schema_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schema.ProtoReflect.Descriptor instead.
func (*Schema) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDescGZIP(), []int{2}
}

func (x *Schema) GetFields() []*NamedSchemaField {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *Schema) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Schema) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type NamedSchema struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Name          string                 `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Fields        []*NamedSchemaField    `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NamedSchema) Reset() {
	*x = NamedSchema{}
	mi := &file_com_coralogix_archive_dataset_v2_schema_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NamedSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamedSchema) ProtoMessage() {}

func (x *NamedSchema) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_dataset_v2_schema_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamedSchema.ProtoReflect.Descriptor instead.
func (*NamedSchema) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDescGZIP(), []int{3}
}

func (x *NamedSchema) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *NamedSchema) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NamedSchema) GetFields() []*NamedSchemaField {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *NamedSchema) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *NamedSchema) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type NamedSchemaField struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	PathArray     []string                  `protobuf:"bytes,1,rep,name=path_array,json=pathArray,proto3" json:"path_array,omitempty"`
	DataType      DataType                  `protobuf:"varint,2,opt,name=data_type,json=dataType,proto3,enum=com.coralogix.archive.dataset.v2.DataType" json:"data_type,omitempty"`
	CreatedAt     *timestamppb.Timestamp    `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp    `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	FieldSettings *NamedSchemaFieldSettings `protobuf:"bytes,5,opt,name=field_settings,json=fieldSettings,proto3" json:"field_settings,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NamedSchemaField) Reset() {
	*x = NamedSchemaField{}
	mi := &file_com_coralogix_archive_dataset_v2_schema_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NamedSchemaField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamedSchemaField) ProtoMessage() {}

func (x *NamedSchemaField) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_dataset_v2_schema_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamedSchemaField.ProtoReflect.Descriptor instead.
func (*NamedSchemaField) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDescGZIP(), []int{4}
}

func (x *NamedSchemaField) GetPathArray() []string {
	if x != nil {
		return x.PathArray
	}
	return nil
}

func (x *NamedSchemaField) GetDataType() DataType {
	if x != nil {
		return x.DataType
	}
	return DataType_DATA_TYPE_UNSPECIFIED
}

func (x *NamedSchemaField) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *NamedSchemaField) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *NamedSchemaField) GetFieldSettings() *NamedSchemaFieldSettings {
	if x != nil {
		return x.FieldSettings
	}
	return nil
}

type NamedSchemaFieldSettings struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsReserved    FieldReservation       `protobuf:"varint,1,opt,name=is_reserved,json=isReserved,proto3,enum=com.coralogix.archive.dataset.v2.FieldReservation" json:"is_reserved,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NamedSchemaFieldSettings) Reset() {
	*x = NamedSchemaFieldSettings{}
	mi := &file_com_coralogix_archive_dataset_v2_schema_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NamedSchemaFieldSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamedSchemaFieldSettings) ProtoMessage() {}

func (x *NamedSchemaFieldSettings) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_dataset_v2_schema_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamedSchemaFieldSettings.ProtoReflect.Descriptor instead.
func (*NamedSchemaFieldSettings) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDescGZIP(), []int{5}
}

func (x *NamedSchemaFieldSettings) GetIsReserved() FieldReservation {
	if x != nil {
		return x.IsReserved
	}
	return FieldReservation_FIELD_RESERVATION_NOT_RESERVED_UNSPECIFIED
}

type GetNamedLogsSchemaRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetNamedLogsSchemaRequest) Reset() {
	*x = GetNamedLogsSchemaRequest{}
	mi := &file_com_coralogix_archive_dataset_v2_schema_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetNamedLogsSchemaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNamedLogsSchemaRequest) ProtoMessage() {}

func (x *GetNamedLogsSchemaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_dataset_v2_schema_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNamedLogsSchemaRequest.ProtoReflect.Descriptor instead.
func (*GetNamedLogsSchemaRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDescGZIP(), []int{6}
}

type GetNamedLogsSchemaResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Schema        *NamedSchema           `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetNamedLogsSchemaResponse) Reset() {
	*x = GetNamedLogsSchemaResponse{}
	mi := &file_com_coralogix_archive_dataset_v2_schema_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetNamedLogsSchemaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNamedLogsSchemaResponse) ProtoMessage() {}

func (x *GetNamedLogsSchemaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_dataset_v2_schema_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNamedLogsSchemaResponse.ProtoReflect.Descriptor instead.
func (*GetNamedLogsSchemaResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetNamedLogsSchemaResponse) GetSchema() *NamedSchema {
	if x != nil {
		return x.Schema
	}
	return nil
}

var File_com_coralogix_archive_dataset_v2_schema_service_proto protoreflect.FileDescriptor

const file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDesc = "" +
	"\n" +
	"5com/coralogix/archive/dataset/v2/schema_service.proto\x12 com.coralogix.archive.dataset.v2\x1a\x1fgoogle/protobuf/timestamp.proto\x1a7com/coralogix/archive/dataset/v2/preferred_schema.proto\x1a'com/coralogix/common/v1/audit_log.proto\"\x82\x01\n" +
	"\x19SetNamedLogsSchemaRequest\x12@\n" +
	"\x06schema\x18\x01 \x01(\v2(.com.coralogix.archive.dataset.v2.SchemaR\x06schema\x12#\n" +
	"\rupdate_schema\x18\x02 \x01(\bR\fupdateSchema\"c\n" +
	"\x1aSetNamedLogsSchemaResponse\x12E\n" +
	"\x06schema\x18\x01 \x01(\v2-.com.coralogix.archive.dataset.v2.NamedSchemaR\x06schema\"\xca\x01\n" +
	"\x06Schema\x12J\n" +
	"\x06fields\x18\x02 \x03(\v22.com.coralogix.archive.dataset.v2.NamedSchemaFieldR\x06fields\x129\n" +
	"\n" +
	"created_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\"\xf5\x01\n" +
	"\vNamedSchema\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x12\n" +
	"\x04name\x18\x05 \x01(\tR\x04name\x12J\n" +
	"\x06fields\x18\x02 \x03(\v22.com.coralogix.archive.dataset.v2.NamedSchemaFieldR\x06fields\x129\n" +
	"\n" +
	"created_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\"\xd3\x02\n" +
	"\x10NamedSchemaField\x12\x1d\n" +
	"\n" +
	"path_array\x18\x01 \x03(\tR\tpathArray\x12G\n" +
	"\tdata_type\x18\x02 \x01(\x0e2*.com.coralogix.archive.dataset.v2.DataTypeR\bdataType\x129\n" +
	"\n" +
	"created_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\x12a\n" +
	"\x0efield_settings\x18\x05 \x01(\v2:.com.coralogix.archive.dataset.v2.NamedSchemaFieldSettingsR\rfieldSettings\"o\n" +
	"\x18NamedSchemaFieldSettings\x12S\n" +
	"\vis_reserved\x18\x01 \x01(\x0e22.com.coralogix.archive.dataset.v2.FieldReservationR\n" +
	"isReserved\"\x1b\n" +
	"\x19GetNamedLogsSchemaRequest\"c\n" +
	"\x1aGetNamedLogsSchemaResponse\x12E\n" +
	"\x06schema\x18\x01 \x01(\v2-.com.coralogix.archive.dataset.v2.NamedSchemaR\x06schema2\xef\x02\n" +
	"\rSchemaService\x12\xad\x01\n" +
	"\x12SetNamedLogsSchema\x12;.com.coralogix.archive.dataset.v2.SetNamedLogsSchemaRequest\x1a<.com.coralogix.archive.dataset.v2.SetNamedLogsSchemaResponse\"\x1c¸\x02\x18\n" +
	"\x16Set schema for company\x12\xad\x01\n" +
	"\x12GetNamedLogsSchema\x12;.com.coralogix.archive.dataset.v2.GetNamedLogsSchemaRequest\x1a<.com.coralogix.archive.dataset.v2.GetNamedLogsSchemaResponse\"\x1c¸\x02\x18\n" +
	"\x16Get schema for companyb\x06proto3"

var (
	file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDescOnce sync.Once
	file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDescData []byte
)

func file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDescGZIP() []byte {
	file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDescOnce.Do(func() {
		file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDesc), len(file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDesc)))
	})
	return file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDescData
}

var file_com_coralogix_archive_dataset_v2_schema_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_com_coralogix_archive_dataset_v2_schema_service_proto_goTypes = []any{
	(*SetNamedLogsSchemaRequest)(nil),  // 0: com.coralogix.archive.dataset.v2.SetNamedLogsSchemaRequest
	(*SetNamedLogsSchemaResponse)(nil), // 1: com.coralogix.archive.dataset.v2.SetNamedLogsSchemaResponse
	(*Schema)(nil),                     // 2: com.coralogix.archive.dataset.v2.Schema
	(*NamedSchema)(nil),                // 3: com.coralogix.archive.dataset.v2.NamedSchema
	(*NamedSchemaField)(nil),           // 4: com.coralogix.archive.dataset.v2.NamedSchemaField
	(*NamedSchemaFieldSettings)(nil),   // 5: com.coralogix.archive.dataset.v2.NamedSchemaFieldSettings
	(*GetNamedLogsSchemaRequest)(nil),  // 6: com.coralogix.archive.dataset.v2.GetNamedLogsSchemaRequest
	(*GetNamedLogsSchemaResponse)(nil), // 7: com.coralogix.archive.dataset.v2.GetNamedLogsSchemaResponse
	(*timestamppb.Timestamp)(nil),      // 8: google.protobuf.Timestamp
	(DataType)(0),                      // 9: com.coralogix.archive.dataset.v2.DataType
	(FieldReservation)(0),              // 10: com.coralogix.archive.dataset.v2.FieldReservation
}
var file_com_coralogix_archive_dataset_v2_schema_service_proto_depIdxs = []int32{
	2,  // 0: com.coralogix.archive.dataset.v2.SetNamedLogsSchemaRequest.schema:type_name -> com.coralogix.archive.dataset.v2.Schema
	3,  // 1: com.coralogix.archive.dataset.v2.SetNamedLogsSchemaResponse.schema:type_name -> com.coralogix.archive.dataset.v2.NamedSchema
	4,  // 2: com.coralogix.archive.dataset.v2.Schema.fields:type_name -> com.coralogix.archive.dataset.v2.NamedSchemaField
	8,  // 3: com.coralogix.archive.dataset.v2.Schema.created_at:type_name -> google.protobuf.Timestamp
	8,  // 4: com.coralogix.archive.dataset.v2.Schema.updated_at:type_name -> google.protobuf.Timestamp
	4,  // 5: com.coralogix.archive.dataset.v2.NamedSchema.fields:type_name -> com.coralogix.archive.dataset.v2.NamedSchemaField
	8,  // 6: com.coralogix.archive.dataset.v2.NamedSchema.created_at:type_name -> google.protobuf.Timestamp
	8,  // 7: com.coralogix.archive.dataset.v2.NamedSchema.updated_at:type_name -> google.protobuf.Timestamp
	9,  // 8: com.coralogix.archive.dataset.v2.NamedSchemaField.data_type:type_name -> com.coralogix.archive.dataset.v2.DataType
	8,  // 9: com.coralogix.archive.dataset.v2.NamedSchemaField.created_at:type_name -> google.protobuf.Timestamp
	8,  // 10: com.coralogix.archive.dataset.v2.NamedSchemaField.updated_at:type_name -> google.protobuf.Timestamp
	5,  // 11: com.coralogix.archive.dataset.v2.NamedSchemaField.field_settings:type_name -> com.coralogix.archive.dataset.v2.NamedSchemaFieldSettings
	10, // 12: com.coralogix.archive.dataset.v2.NamedSchemaFieldSettings.is_reserved:type_name -> com.coralogix.archive.dataset.v2.FieldReservation
	3,  // 13: com.coralogix.archive.dataset.v2.GetNamedLogsSchemaResponse.schema:type_name -> com.coralogix.archive.dataset.v2.NamedSchema
	0,  // 14: com.coralogix.archive.dataset.v2.SchemaService.SetNamedLogsSchema:input_type -> com.coralogix.archive.dataset.v2.SetNamedLogsSchemaRequest
	6,  // 15: com.coralogix.archive.dataset.v2.SchemaService.GetNamedLogsSchema:input_type -> com.coralogix.archive.dataset.v2.GetNamedLogsSchemaRequest
	1,  // 16: com.coralogix.archive.dataset.v2.SchemaService.SetNamedLogsSchema:output_type -> com.coralogix.archive.dataset.v2.SetNamedLogsSchemaResponse
	7,  // 17: com.coralogix.archive.dataset.v2.SchemaService.GetNamedLogsSchema:output_type -> com.coralogix.archive.dataset.v2.GetNamedLogsSchemaResponse
	16, // [16:18] is the sub-list for method output_type
	14, // [14:16] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_com_coralogix_archive_dataset_v2_schema_service_proto_init() }
func file_com_coralogix_archive_dataset_v2_schema_service_proto_init() {
	if File_com_coralogix_archive_dataset_v2_schema_service_proto != nil {
		return
	}
	file_com_coralogix_archive_dataset_v2_preferred_schema_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDesc), len(file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogix_archive_dataset_v2_schema_service_proto_goTypes,
		DependencyIndexes: file_com_coralogix_archive_dataset_v2_schema_service_proto_depIdxs,
		MessageInfos:      file_com_coralogix_archive_dataset_v2_schema_service_proto_msgTypes,
	}.Build()
	File_com_coralogix_archive_dataset_v2_schema_service_proto = out.File
	file_com_coralogix_archive_dataset_v2_schema_service_proto_goTypes = nil
	file_com_coralogix_archive_dataset_v2_schema_service_proto_depIdxs = nil
}
