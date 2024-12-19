// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.26.1
// source: com/coralogix/archive/private/state/v1/schema.proto

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

type Schema struct {
	state         protoimpl.MessageState           `protogen:"open.v1"`
	Payloads      map[string]*GenericPayloadSchema `protobuf:"bytes,1,rep,name=payloads,proto3" json:"payloads,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Truncated     *bool                            `protobuf:"varint,2,opt,name=truncated,proto3,oneof" json:"truncated,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Schema) Reset() {
	*x = Schema{}
	mi := &file_com_coralogix_archive_private_state_v1_schema_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Schema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schema) ProtoMessage() {}

func (x *Schema) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_private_state_v1_schema_proto_msgTypes[0]
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
	return file_com_coralogix_archive_private_state_v1_schema_proto_rawDescGZIP(), []int{0}
}

func (x *Schema) GetPayloads() map[string]*GenericPayloadSchema {
	if x != nil {
		return x.Payloads
	}
	return nil
}

func (x *Schema) GetTruncated() bool {
	if x != nil && x.Truncated != nil {
		return *x.Truncated
	}
	return false
}

type GenericPayloadSchema struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	Name          string                      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Fields        *GenericPayloadSchemaFields `protobuf:"bytes,2,opt,name=fields,proto3" json:"fields,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenericPayloadSchema) Reset() {
	*x = GenericPayloadSchema{}
	mi := &file_com_coralogix_archive_private_state_v1_schema_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenericPayloadSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericPayloadSchema) ProtoMessage() {}

func (x *GenericPayloadSchema) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_private_state_v1_schema_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericPayloadSchema.ProtoReflect.Descriptor instead.
func (*GenericPayloadSchema) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_private_state_v1_schema_proto_rawDescGZIP(), []int{1}
}

func (x *GenericPayloadSchema) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GenericPayloadSchema) GetFields() *GenericPayloadSchemaFields {
	if x != nil {
		return x.Fields
	}
	return nil
}

type GenericPayloadSchemaFields struct {
	state         protoimpl.MessageState           `protogen:"open.v1"`
	FieldsMap     []*GenericPayloadSchemaFieldsMap `protobuf:"bytes,1,rep,name=fields_map,json=fieldsMap,proto3" json:"fields_map,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenericPayloadSchemaFields) Reset() {
	*x = GenericPayloadSchemaFields{}
	mi := &file_com_coralogix_archive_private_state_v1_schema_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenericPayloadSchemaFields) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericPayloadSchemaFields) ProtoMessage() {}

func (x *GenericPayloadSchemaFields) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_private_state_v1_schema_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericPayloadSchemaFields.ProtoReflect.Descriptor instead.
func (*GenericPayloadSchemaFields) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_private_state_v1_schema_proto_rawDescGZIP(), []int{2}
}

func (x *GenericPayloadSchemaFields) GetFieldsMap() []*GenericPayloadSchemaFieldsMap {
	if x != nil {
		return x.FieldsMap
	}
	return nil
}

type GenericPayloadSchemaFieldsKey struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FieldName     *FieldName             `protobuf:"bytes,1,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	FieldType     FieldType              `protobuf:"varint,2,opt,name=field_type,json=fieldType,proto3,enum=com.coralogix.archive.private.state.v1.FieldType" json:"field_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenericPayloadSchemaFieldsKey) Reset() {
	*x = GenericPayloadSchemaFieldsKey{}
	mi := &file_com_coralogix_archive_private_state_v1_schema_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenericPayloadSchemaFieldsKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericPayloadSchemaFieldsKey) ProtoMessage() {}

func (x *GenericPayloadSchemaFieldsKey) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_private_state_v1_schema_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericPayloadSchemaFieldsKey.ProtoReflect.Descriptor instead.
func (*GenericPayloadSchemaFieldsKey) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_private_state_v1_schema_proto_rawDescGZIP(), []int{3}
}

func (x *GenericPayloadSchemaFieldsKey) GetFieldName() *FieldName {
	if x != nil {
		return x.FieldName
	}
	return nil
}

func (x *GenericPayloadSchemaFieldsKey) GetFieldType() FieldType {
	if x != nil {
		return x.FieldType
	}
	return FieldType_FIELD_TYPE_UNSPECIFIED
}

type GenericPayloadSchemaFieldsValue struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FieldData     *FieldData             `protobuf:"bytes,1,opt,name=field_data,json=fieldData,proto3" json:"field_data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenericPayloadSchemaFieldsValue) Reset() {
	*x = GenericPayloadSchemaFieldsValue{}
	mi := &file_com_coralogix_archive_private_state_v1_schema_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenericPayloadSchemaFieldsValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericPayloadSchemaFieldsValue) ProtoMessage() {}

func (x *GenericPayloadSchemaFieldsValue) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_private_state_v1_schema_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericPayloadSchemaFieldsValue.ProtoReflect.Descriptor instead.
func (*GenericPayloadSchemaFieldsValue) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_private_state_v1_schema_proto_rawDescGZIP(), []int{4}
}

func (x *GenericPayloadSchemaFieldsValue) GetFieldData() *FieldData {
	if x != nil {
		return x.FieldData
	}
	return nil
}

type GenericPayloadSchemaFieldsMap struct {
	state         protoimpl.MessageState           `protogen:"open.v1"`
	Key           *GenericPayloadSchemaFieldsKey   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value         *GenericPayloadSchemaFieldsValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenericPayloadSchemaFieldsMap) Reset() {
	*x = GenericPayloadSchemaFieldsMap{}
	mi := &file_com_coralogix_archive_private_state_v1_schema_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenericPayloadSchemaFieldsMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericPayloadSchemaFieldsMap) ProtoMessage() {}

func (x *GenericPayloadSchemaFieldsMap) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_private_state_v1_schema_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericPayloadSchemaFieldsMap.ProtoReflect.Descriptor instead.
func (*GenericPayloadSchemaFieldsMap) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_private_state_v1_schema_proto_rawDescGZIP(), []int{5}
}

func (x *GenericPayloadSchemaFieldsMap) GetKey() *GenericPayloadSchemaFieldsKey {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *GenericPayloadSchemaFieldsMap) GetValue() *GenericPayloadSchemaFieldsValue {
	if x != nil {
		return x.Value
	}
	return nil
}

type FieldData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          *FieldName             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type          FieldType              `protobuf:"varint,2,opt,name=type,proto3,enum=com.coralogix.archive.private.state.v1.FieldType" json:"type,omitempty"`
	Required      bool                   `protobuf:"varint,3,opt,name=required,proto3" json:"required,omitempty"`
	Occurrences   int32                  `protobuf:"varint,4,opt,name=occurrences,proto3" json:"occurrences,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FieldData) Reset() {
	*x = FieldData{}
	mi := &file_com_coralogix_archive_private_state_v1_schema_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FieldData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldData) ProtoMessage() {}

func (x *FieldData) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_private_state_v1_schema_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldData.ProtoReflect.Descriptor instead.
func (*FieldData) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_private_state_v1_schema_proto_rawDescGZIP(), []int{6}
}

func (x *FieldData) GetName() *FieldName {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *FieldData) GetType() FieldType {
	if x != nil {
		return x.Type
	}
	return FieldType_FIELD_TYPE_UNSPECIFIED
}

func (x *FieldData) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *FieldData) GetOccurrences() int32 {
	if x != nil {
		return x.Occurrences
	}
	return 0
}

type FieldName struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Parts         []string               `protobuf:"bytes,1,rep,name=parts,proto3" json:"parts,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FieldName) Reset() {
	*x = FieldName{}
	mi := &file_com_coralogix_archive_private_state_v1_schema_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FieldName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldName) ProtoMessage() {}

func (x *FieldName) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_private_state_v1_schema_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldName.ProtoReflect.Descriptor instead.
func (*FieldName) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_private_state_v1_schema_proto_rawDescGZIP(), []int{7}
}

func (x *FieldName) GetParts() []string {
	if x != nil {
		return x.Parts
	}
	return nil
}

var File_com_coralogix_archive_private_state_v1_schema_proto protoreflect.FileDescriptor

var file_com_coralogix_archive_private_state_v1_schema_proto_rawDesc = []byte{
	0x0a, 0x33, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x37, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x02, 0x0a, 0x06, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x12, 0x58, 0x0a, 0x08, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x08, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x12, 0x21, 0x0a, 0x09, 0x74,
	0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00,
	0x52, 0x09, 0x74, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01, 0x1a, 0x79,
	0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x52, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69,
	0x63, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x74, 0x72,
	0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x64, 0x22, 0x86, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x69, 0x63, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x5a, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x69, 0x63, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x22, 0x82, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12,
	0x64, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x69, 0x63, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x4d, 0x61, 0x70, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x4d, 0x61, 0x70, 0x22, 0xc3, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69,
	0x63, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x50, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x09,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x50, 0x0a, 0x0a, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x22, 0x73, 0x0a, 0x1f, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x50,
	0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x22, 0xd7, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x4d,
	0x61, 0x70, 0x12, 0x57, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x45, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x5d, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xd7, 0x01, 0x0a, 0x09, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x45, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x45, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x73, 0x22, 0x21, 0x0a, 0x09, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x61, 0x72, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_archive_private_state_v1_schema_proto_rawDescOnce sync.Once
	file_com_coralogix_archive_private_state_v1_schema_proto_rawDescData = file_com_coralogix_archive_private_state_v1_schema_proto_rawDesc
)

func file_com_coralogix_archive_private_state_v1_schema_proto_rawDescGZIP() []byte {
	file_com_coralogix_archive_private_state_v1_schema_proto_rawDescOnce.Do(func() {
		file_com_coralogix_archive_private_state_v1_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_archive_private_state_v1_schema_proto_rawDescData)
	})
	return file_com_coralogix_archive_private_state_v1_schema_proto_rawDescData
}

var file_com_coralogix_archive_private_state_v1_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_com_coralogix_archive_private_state_v1_schema_proto_goTypes = []any{
	(*Schema)(nil),                          // 0: com.coralogix.archive.private.state.v1.Schema
	(*GenericPayloadSchema)(nil),            // 1: com.coralogix.archive.private.state.v1.GenericPayloadSchema
	(*GenericPayloadSchemaFields)(nil),      // 2: com.coralogix.archive.private.state.v1.GenericPayloadSchemaFields
	(*GenericPayloadSchemaFieldsKey)(nil),   // 3: com.coralogix.archive.private.state.v1.GenericPayloadSchemaFieldsKey
	(*GenericPayloadSchemaFieldsValue)(nil), // 4: com.coralogix.archive.private.state.v1.GenericPayloadSchemaFieldsValue
	(*GenericPayloadSchemaFieldsMap)(nil),   // 5: com.coralogix.archive.private.state.v1.GenericPayloadSchemaFieldsMap
	(*FieldData)(nil),                       // 6: com.coralogix.archive.private.state.v1.FieldData
	(*FieldName)(nil),                       // 7: com.coralogix.archive.private.state.v1.FieldName
	nil,                                     // 8: com.coralogix.archive.private.state.v1.Schema.PayloadsEntry
	(FieldType)(0),                          // 9: com.coralogix.archive.private.state.v1.FieldType
}
var file_com_coralogix_archive_private_state_v1_schema_proto_depIdxs = []int32{
	8,  // 0: com.coralogix.archive.private.state.v1.Schema.payloads:type_name -> com.coralogix.archive.private.state.v1.Schema.PayloadsEntry
	2,  // 1: com.coralogix.archive.private.state.v1.GenericPayloadSchema.fields:type_name -> com.coralogix.archive.private.state.v1.GenericPayloadSchemaFields
	5,  // 2: com.coralogix.archive.private.state.v1.GenericPayloadSchemaFields.fields_map:type_name -> com.coralogix.archive.private.state.v1.GenericPayloadSchemaFieldsMap
	7,  // 3: com.coralogix.archive.private.state.v1.GenericPayloadSchemaFieldsKey.field_name:type_name -> com.coralogix.archive.private.state.v1.FieldName
	9,  // 4: com.coralogix.archive.private.state.v1.GenericPayloadSchemaFieldsKey.field_type:type_name -> com.coralogix.archive.private.state.v1.FieldType
	6,  // 5: com.coralogix.archive.private.state.v1.GenericPayloadSchemaFieldsValue.field_data:type_name -> com.coralogix.archive.private.state.v1.FieldData
	3,  // 6: com.coralogix.archive.private.state.v1.GenericPayloadSchemaFieldsMap.key:type_name -> com.coralogix.archive.private.state.v1.GenericPayloadSchemaFieldsKey
	4,  // 7: com.coralogix.archive.private.state.v1.GenericPayloadSchemaFieldsMap.value:type_name -> com.coralogix.archive.private.state.v1.GenericPayloadSchemaFieldsValue
	7,  // 8: com.coralogix.archive.private.state.v1.FieldData.name:type_name -> com.coralogix.archive.private.state.v1.FieldName
	9,  // 9: com.coralogix.archive.private.state.v1.FieldData.type:type_name -> com.coralogix.archive.private.state.v1.FieldType
	1,  // 10: com.coralogix.archive.private.state.v1.Schema.PayloadsEntry.value:type_name -> com.coralogix.archive.private.state.v1.GenericPayloadSchema
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_com_coralogix_archive_private_state_v1_schema_proto_init() }
func file_com_coralogix_archive_private_state_v1_schema_proto_init() {
	if File_com_coralogix_archive_private_state_v1_schema_proto != nil {
		return
	}
	file_com_coralogix_archive_private_state_v1_field_type_proto_init()
	file_com_coralogix_archive_private_state_v1_schema_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_archive_private_state_v1_schema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_archive_private_state_v1_schema_proto_goTypes,
		DependencyIndexes: file_com_coralogix_archive_private_state_v1_schema_proto_depIdxs,
		MessageInfos:      file_com_coralogix_archive_private_state_v1_schema_proto_msgTypes,
	}.Build()
	File_com_coralogix_archive_private_state_v1_schema_proto = out.File
	file_com_coralogix_archive_private_state_v1_schema_proto_rawDesc = nil
	file_com_coralogix_archive_private_state_v1_schema_proto_goTypes = nil
	file_com_coralogix_archive_private_state_v1_schema_proto_depIdxs = nil
}
