// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: com/coralogix/archive/dataset/v2/schema_service.proto

package v2

import (
	_ "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/common/v1"
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

type SetNamedSchemaRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Schema        *NamedSchema           `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
	UpdateSchema  bool                   `protobuf:"varint,2,opt,name=update_schema,json=updateSchema,proto3" json:"update_schema,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetNamedSchemaRequest) Reset() {
	*x = SetNamedSchemaRequest{}
	mi := &file_com_coralogix_archive_dataset_v2_schema_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetNamedSchemaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetNamedSchemaRequest) ProtoMessage() {}

func (x *SetNamedSchemaRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SetNamedSchemaRequest.ProtoReflect.Descriptor instead.
func (*SetNamedSchemaRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDescGZIP(), []int{0}
}

func (x *SetNamedSchemaRequest) GetSchema() *NamedSchema {
	if x != nil {
		return x.Schema
	}
	return nil
}

func (x *SetNamedSchemaRequest) GetUpdateSchema() bool {
	if x != nil {
		return x.UpdateSchema
	}
	return false
}

type SetNamedSchemaResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Schema        *NamedSchema           `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetNamedSchemaResponse) Reset() {
	*x = SetNamedSchemaResponse{}
	mi := &file_com_coralogix_archive_dataset_v2_schema_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetNamedSchemaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetNamedSchemaResponse) ProtoMessage() {}

func (x *SetNamedSchemaResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SetNamedSchemaResponse.ProtoReflect.Descriptor instead.
func (*SetNamedSchemaResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDescGZIP(), []int{1}
}

func (x *SetNamedSchemaResponse) GetSchema() *NamedSchema {
	if x != nil {
		return x.Schema
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
	mi := &file_com_coralogix_archive_dataset_v2_schema_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NamedSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamedSchema) ProtoMessage() {}

func (x *NamedSchema) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use NamedSchema.ProtoReflect.Descriptor instead.
func (*NamedSchema) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDescGZIP(), []int{2}
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
	mi := &file_com_coralogix_archive_dataset_v2_schema_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NamedSchemaField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamedSchemaField) ProtoMessage() {}

func (x *NamedSchemaField) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use NamedSchemaField.ProtoReflect.Descriptor instead.
func (*NamedSchemaField) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDescGZIP(), []int{3}
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
	mi := &file_com_coralogix_archive_dataset_v2_schema_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NamedSchemaFieldSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamedSchemaFieldSettings) ProtoMessage() {}

func (x *NamedSchemaFieldSettings) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use NamedSchemaFieldSettings.ProtoReflect.Descriptor instead.
func (*NamedSchemaFieldSettings) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDescGZIP(), []int{4}
}

func (x *NamedSchemaFieldSettings) GetIsReserved() FieldReservation {
	if x != nil {
		return x.IsReserved
	}
	return FieldReservation_FIELD_RESERVATION_NOT_RESERVED_UNSPECIFIED
}

type GetNamedSchemaRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetNamedSchemaRequest) Reset() {
	*x = GetNamedSchemaRequest{}
	mi := &file_com_coralogix_archive_dataset_v2_schema_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetNamedSchemaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNamedSchemaRequest) ProtoMessage() {}

func (x *GetNamedSchemaRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetNamedSchemaRequest.ProtoReflect.Descriptor instead.
func (*GetNamedSchemaRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetNamedSchemaRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetNamedSchemaResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Schema        *NamedSchema           `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetNamedSchemaResponse) Reset() {
	*x = GetNamedSchemaResponse{}
	mi := &file_com_coralogix_archive_dataset_v2_schema_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetNamedSchemaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNamedSchemaResponse) ProtoMessage() {}

func (x *GetNamedSchemaResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetNamedSchemaResponse.ProtoReflect.Descriptor instead.
func (*GetNamedSchemaResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetNamedSchemaResponse) GetSchema() *NamedSchema {
	if x != nil {
		return x.Schema
	}
	return nil
}

var File_com_coralogix_archive_dataset_v2_schema_service_proto protoreflect.FileDescriptor

var file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDesc = []byte{
	0x0a, 0x35, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2f,
	0x76, 0x32, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x37, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a,
	0x15, 0x53, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x23, 0x0a,
	0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x22, 0x5f, 0x0a, 0x16, 0x53, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x06,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e,
	0x4e, 0x61, 0x6d, 0x65, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x06, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x22, 0xf5, 0x01, 0x0a, 0x0b, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4a, 0x0a, 0x06, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x4e, 0x61, 0x6d,
	0x65, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xd3, 0x02, 0x0a, 0x10,
	0x4e, 0x61, 0x6d, 0x65, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x61, 0x72, 0x72, 0x61, 0x79, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x74, 0x68, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12,
	0x47, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x61,
	0x0a, 0x0e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x52, 0x0d, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x22, 0x6f, 0x0a, 0x18, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x53, 0x0a,
	0x0b, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x69, 0x73, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x64, 0x22, 0x29, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x5f, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x64,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x32, 0xd7,
	0x02, 0x0a, 0x0d, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0xa1, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x12, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e,
	0x53, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0xc2, 0xb8, 0x02, 0x18, 0x0a, 0x16, 0x53, 0x65,
	0x74, 0x20, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x12, 0xa1, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0xc2, 0xb8, 0x02, 0x18,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x20, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x20, 0x66, 0x6f, 0x72,
	0x20, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDescOnce sync.Once
	file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDescData = file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDesc
)

func file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDescGZIP() []byte {
	file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDescOnce.Do(func() {
		file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDescData)
	})
	return file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDescData
}

var file_com_coralogix_archive_dataset_v2_schema_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_com_coralogix_archive_dataset_v2_schema_service_proto_goTypes = []any{
	(*SetNamedSchemaRequest)(nil),    // 0: com.coralogix.archive.dataset.v2.SetNamedSchemaRequest
	(*SetNamedSchemaResponse)(nil),   // 1: com.coralogix.archive.dataset.v2.SetNamedSchemaResponse
	(*NamedSchema)(nil),              // 2: com.coralogix.archive.dataset.v2.NamedSchema
	(*NamedSchemaField)(nil),         // 3: com.coralogix.archive.dataset.v2.NamedSchemaField
	(*NamedSchemaFieldSettings)(nil), // 4: com.coralogix.archive.dataset.v2.NamedSchemaFieldSettings
	(*GetNamedSchemaRequest)(nil),    // 5: com.coralogix.archive.dataset.v2.GetNamedSchemaRequest
	(*GetNamedSchemaResponse)(nil),   // 6: com.coralogix.archive.dataset.v2.GetNamedSchemaResponse
	(*timestamppb.Timestamp)(nil),    // 7: google.protobuf.Timestamp
	(DataType)(0),                    // 8: com.coralogix.archive.dataset.v2.DataType
	(FieldReservation)(0),            // 9: com.coralogix.archive.dataset.v2.FieldReservation
}
var file_com_coralogix_archive_dataset_v2_schema_service_proto_depIdxs = []int32{
	2,  // 0: com.coralogix.archive.dataset.v2.SetNamedSchemaRequest.schema:type_name -> com.coralogix.archive.dataset.v2.NamedSchema
	2,  // 1: com.coralogix.archive.dataset.v2.SetNamedSchemaResponse.schema:type_name -> com.coralogix.archive.dataset.v2.NamedSchema
	3,  // 2: com.coralogix.archive.dataset.v2.NamedSchema.fields:type_name -> com.coralogix.archive.dataset.v2.NamedSchemaField
	7,  // 3: com.coralogix.archive.dataset.v2.NamedSchema.created_at:type_name -> google.protobuf.Timestamp
	7,  // 4: com.coralogix.archive.dataset.v2.NamedSchema.updated_at:type_name -> google.protobuf.Timestamp
	8,  // 5: com.coralogix.archive.dataset.v2.NamedSchemaField.data_type:type_name -> com.coralogix.archive.dataset.v2.DataType
	7,  // 6: com.coralogix.archive.dataset.v2.NamedSchemaField.created_at:type_name -> google.protobuf.Timestamp
	7,  // 7: com.coralogix.archive.dataset.v2.NamedSchemaField.updated_at:type_name -> google.protobuf.Timestamp
	4,  // 8: com.coralogix.archive.dataset.v2.NamedSchemaField.field_settings:type_name -> com.coralogix.archive.dataset.v2.NamedSchemaFieldSettings
	9,  // 9: com.coralogix.archive.dataset.v2.NamedSchemaFieldSettings.is_reserved:type_name -> com.coralogix.archive.dataset.v2.FieldReservation
	2,  // 10: com.coralogix.archive.dataset.v2.GetNamedSchemaResponse.schema:type_name -> com.coralogix.archive.dataset.v2.NamedSchema
	0,  // 11: com.coralogix.archive.dataset.v2.SchemaService.SetNamedSchema:input_type -> com.coralogix.archive.dataset.v2.SetNamedSchemaRequest
	5,  // 12: com.coralogix.archive.dataset.v2.SchemaService.GetNamedSchema:input_type -> com.coralogix.archive.dataset.v2.GetNamedSchemaRequest
	1,  // 13: com.coralogix.archive.dataset.v2.SchemaService.SetNamedSchema:output_type -> com.coralogix.archive.dataset.v2.SetNamedSchemaResponse
	6,  // 14: com.coralogix.archive.dataset.v2.SchemaService.GetNamedSchema:output_type -> com.coralogix.archive.dataset.v2.GetNamedSchemaResponse
	13, // [13:15] is the sub-list for method output_type
	11, // [11:13] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
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
			RawDescriptor: file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogix_archive_dataset_v2_schema_service_proto_goTypes,
		DependencyIndexes: file_com_coralogix_archive_dataset_v2_schema_service_proto_depIdxs,
		MessageInfos:      file_com_coralogix_archive_dataset_v2_schema_service_proto_msgTypes,
	}.Build()
	File_com_coralogix_archive_dataset_v2_schema_service_proto = out.File
	file_com_coralogix_archive_dataset_v2_schema_service_proto_rawDesc = nil
	file_com_coralogix_archive_dataset_v2_schema_service_proto_goTypes = nil
	file_com_coralogix_archive_dataset_v2_schema_service_proto_depIdxs = nil
}
