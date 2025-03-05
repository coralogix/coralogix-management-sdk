// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: com/coralogix/archive/v2/schema.proto

package v2

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

type Namespace int32

const (
	Namespace_NAMESPACE_UNSPECIFIED Namespace = 0
	Namespace_NAMESPACE_METADATA    Namespace = 1
	Namespace_NAMESPACE_LABELS      Namespace = 2
	Namespace_NAMESPACE_USER_DATA   Namespace = 3
)

// Enum value maps for Namespace.
var (
	Namespace_name = map[int32]string{
		0: "NAMESPACE_UNSPECIFIED",
		1: "NAMESPACE_METADATA",
		2: "NAMESPACE_LABELS",
		3: "NAMESPACE_USER_DATA",
	}
	Namespace_value = map[string]int32{
		"NAMESPACE_UNSPECIFIED": 0,
		"NAMESPACE_METADATA":    1,
		"NAMESPACE_LABELS":      2,
		"NAMESPACE_USER_DATA":   3,
	}
)

func (x Namespace) Enum() *Namespace {
	p := new(Namespace)
	*p = x
	return p
}

func (x Namespace) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Namespace) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_archive_v2_schema_proto_enumTypes[0].Descriptor()
}

func (Namespace) Type() protoreflect.EnumType {
	return &file_com_coralogix_archive_v2_schema_proto_enumTypes[0]
}

func (x Namespace) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Namespace.Descriptor instead.
func (Namespace) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_schema_proto_rawDescGZIP(), []int{0}
}

type SchemaType int32

const (
	SchemaType_SCHEMA_TYPE_UNSPECIFIED      SchemaType = 0
	SchemaType_SCHEMA_TYPE_STRING           SchemaType = 1
	SchemaType_SCHEMA_TYPE_TIMESTAMP_NANOS  SchemaType = 2
	SchemaType_SCHEMA_TYPE_TIMESTAMP_MICROS SchemaType = 3
)

// Enum value maps for SchemaType.
var (
	SchemaType_name = map[int32]string{
		0: "SCHEMA_TYPE_UNSPECIFIED",
		1: "SCHEMA_TYPE_STRING",
		2: "SCHEMA_TYPE_TIMESTAMP_NANOS",
		3: "SCHEMA_TYPE_TIMESTAMP_MICROS",
	}
	SchemaType_value = map[string]int32{
		"SCHEMA_TYPE_UNSPECIFIED":      0,
		"SCHEMA_TYPE_STRING":           1,
		"SCHEMA_TYPE_TIMESTAMP_NANOS":  2,
		"SCHEMA_TYPE_TIMESTAMP_MICROS": 3,
	}
)

func (x SchemaType) Enum() *SchemaType {
	p := new(SchemaType)
	*p = x
	return p
}

func (x SchemaType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SchemaType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_archive_v2_schema_proto_enumTypes[1].Descriptor()
}

func (SchemaType) Type() protoreflect.EnumType {
	return &file_com_coralogix_archive_v2_schema_proto_enumTypes[1]
}

func (x SchemaType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SchemaType.Descriptor instead.
func (SchemaType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_schema_proto_rawDescGZIP(), []int{1}
}

type SubSchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace Namespace             `protobuf:"varint,1,opt,name=namespace,proto3,enum=com.coralogix.archive.v2.Namespace" json:"namespace,omitempty"`
	Fields    map[string]*FieldNode `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SubSchema) Reset() {
	*x = SubSchema{}
	mi := &file_com_coralogix_archive_v2_schema_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubSchema) ProtoMessage() {}

func (x *SubSchema) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_schema_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubSchema.ProtoReflect.Descriptor instead.
func (*SubSchema) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_schema_proto_rawDescGZIP(), []int{0}
}

func (x *SubSchema) GetNamespace() Namespace {
	if x != nil {
		return x.Namespace
	}
	return Namespace_NAMESPACE_UNSPECIFIED
}

func (x *SubSchema) GetFields() map[string]*FieldNode {
	if x != nil {
		return x.Fields
	}
	return nil
}

type SchemaTree struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schemas []*SubSchema `protobuf:"bytes,1,rep,name=schemas,proto3" json:"schemas,omitempty"`
}

func (x *SchemaTree) Reset() {
	*x = SchemaTree{}
	mi := &file_com_coralogix_archive_v2_schema_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SchemaTree) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchemaTree) ProtoMessage() {}

func (x *SchemaTree) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_schema_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchemaTree.ProtoReflect.Descriptor instead.
func (*SchemaTree) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_schema_proto_rawDescGZIP(), []int{1}
}

func (x *SchemaTree) GetSchemas() []*SubSchema {
	if x != nil {
		return x.Schemas
	}
	return nil
}

type FieldNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Children map[string]*FieldNode `protobuf:"bytes,2,rep,name=children,proto3" json:"children,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Metadata *FieldMetadata        `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *FieldNode) Reset() {
	*x = FieldNode{}
	mi := &file_com_coralogix_archive_v2_schema_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FieldNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldNode) ProtoMessage() {}

func (x *FieldNode) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_schema_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldNode.ProtoReflect.Descriptor instead.
func (*FieldNode) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_schema_proto_rawDescGZIP(), []int{2}
}

func (x *FieldNode) GetChildren() map[string]*FieldNode {
	if x != nil {
		return x.Children
	}
	return nil
}

func (x *FieldNode) GetMetadata() *FieldMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type FieldMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location *FieldLocation `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Required bool           `protobuf:"varint,2,opt,name=required,proto3" json:"required,omitempty"`
}

func (x *FieldMetadata) Reset() {
	*x = FieldMetadata{}
	mi := &file_com_coralogix_archive_v2_schema_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FieldMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldMetadata) ProtoMessage() {}

func (x *FieldMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_schema_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldMetadata.ProtoReflect.Descriptor instead.
func (*FieldMetadata) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_schema_proto_rawDescGZIP(), []int{3}
}

func (x *FieldMetadata) GetLocation() *FieldLocation {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *FieldMetadata) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

type FieldLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Location:
	//
	//	*FieldLocation_Payload
	//	*FieldLocation_Column
	Location isFieldLocation_Location `protobuf_oneof:"location"`
}

func (x *FieldLocation) Reset() {
	*x = FieldLocation{}
	mi := &file_com_coralogix_archive_v2_schema_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FieldLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldLocation) ProtoMessage() {}

func (x *FieldLocation) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_schema_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldLocation.ProtoReflect.Descriptor instead.
func (*FieldLocation) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_schema_proto_rawDescGZIP(), []int{4}
}

func (m *FieldLocation) GetLocation() isFieldLocation_Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (x *FieldLocation) GetPayload() *Payload {
	if x, ok := x.GetLocation().(*FieldLocation_Payload); ok {
		return x.Payload
	}
	return nil
}

func (x *FieldLocation) GetColumn() *Column {
	if x, ok := x.GetLocation().(*FieldLocation_Column); ok {
		return x.Column
	}
	return nil
}

type isFieldLocation_Location interface {
	isFieldLocation_Location()
}

type FieldLocation_Payload struct {
	Payload *Payload `protobuf:"bytes,1,opt,name=payload,proto3,oneof"`
}

type FieldLocation_Column struct {
	Column *Column `protobuf:"bytes,2,opt,name=column,proto3,oneof"`
}

func (*FieldLocation_Payload) isFieldLocation_Location() {}

func (*FieldLocation_Column) isFieldLocation_Location() {}

type Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Payload) Reset() {
	*x = Payload{}
	mi := &file_com_coralogix_archive_v2_schema_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payload) ProtoMessage() {}

func (x *Payload) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_schema_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payload.ProtoReflect.Descriptor instead.
func (*Payload) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_schema_proto_rawDescGZIP(), []int{5}
}

type Column struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index int32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *Column) Reset() {
	*x = Column{}
	mi := &file_com_coralogix_archive_v2_schema_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Column) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Column) ProtoMessage() {}

func (x *Column) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_schema_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Column.ProtoReflect.Descriptor instead.
func (*Column) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_schema_proto_rawDescGZIP(), []int{6}
}

func (x *Column) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

var File_com_coralogix_archive_v2_schema_proto protoreflect.FileDescriptor

var file_com_coralogix_archive_v2_schema_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76,
	0x32, 0x22, 0xf7, 0x01, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12,
	0x41, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x47, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x75,
	0x62, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a, 0x5e, 0x0a, 0x0b, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x39, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4b, 0x0a, 0x0a, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x54, 0x72, 0x65, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x75, 0x62, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52,
	0x07, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x22, 0x87, 0x02, 0x0a, 0x09, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x4d, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x2e, 0x76, 0x32, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x43, 0x68,
	0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x63, 0x68, 0x69,
	0x6c, 0x64, 0x72, 0x65, 0x6e, 0x12, 0x43, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e,
	0x76, 0x32, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x60, 0x0a, 0x0d, 0x43, 0x68,
	0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x39, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x6f, 0x64,
	0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x4a, 0x04, 0x08, 0x01,
	0x10, 0x02, 0x22, 0x70, 0x0a, 0x0d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x43, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x22, 0x96, 0x01, 0x0a, 0x0d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e,
	0x76, 0x32, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x00, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x3a, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32,
	0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x48, 0x00, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x42, 0x0a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x09, 0x0a,
	0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x1e, 0x0a, 0x06, 0x43, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2a, 0x6d, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x4e, 0x41, 0x4d, 0x45, 0x53, 0x50, 0x41,
	0x43, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x16, 0x0a, 0x12, 0x4e, 0x41, 0x4d, 0x45, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x4d, 0x45,
	0x54, 0x41, 0x44, 0x41, 0x54, 0x41, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x4e, 0x41, 0x4d, 0x45,
	0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x53, 0x10, 0x02, 0x12, 0x17,
	0x0a, 0x13, 0x4e, 0x41, 0x4d, 0x45, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x52,
	0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0x03, 0x2a, 0x84, 0x01, 0x0a, 0x0a, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x53,
	0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x53,
	0x54, 0x41, 0x4d, 0x50, 0x5f, 0x4e, 0x41, 0x4e, 0x4f, 0x53, 0x10, 0x02, 0x12, 0x20, 0x0a, 0x1c,
	0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45,
	0x53, 0x54, 0x41, 0x4d, 0x50, 0x5f, 0x4d, 0x49, 0x43, 0x52, 0x4f, 0x53, 0x10, 0x03, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_archive_v2_schema_proto_rawDescOnce sync.Once
	file_com_coralogix_archive_v2_schema_proto_rawDescData = file_com_coralogix_archive_v2_schema_proto_rawDesc
)

func file_com_coralogix_archive_v2_schema_proto_rawDescGZIP() []byte {
	file_com_coralogix_archive_v2_schema_proto_rawDescOnce.Do(func() {
		file_com_coralogix_archive_v2_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_archive_v2_schema_proto_rawDescData)
	})
	return file_com_coralogix_archive_v2_schema_proto_rawDescData
}

var file_com_coralogix_archive_v2_schema_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_com_coralogix_archive_v2_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_com_coralogix_archive_v2_schema_proto_goTypes = []any{
	(Namespace)(0),        // 0: com.coralogix.archive.v2.Namespace
	(SchemaType)(0),       // 1: com.coralogix.archive.v2.SchemaType
	(*SubSchema)(nil),     // 2: com.coralogix.archive.v2.SubSchema
	(*SchemaTree)(nil),    // 3: com.coralogix.archive.v2.SchemaTree
	(*FieldNode)(nil),     // 4: com.coralogix.archive.v2.FieldNode
	(*FieldMetadata)(nil), // 5: com.coralogix.archive.v2.FieldMetadata
	(*FieldLocation)(nil), // 6: com.coralogix.archive.v2.FieldLocation
	(*Payload)(nil),       // 7: com.coralogix.archive.v2.Payload
	(*Column)(nil),        // 8: com.coralogix.archive.v2.Column
	nil,                   // 9: com.coralogix.archive.v2.SubSchema.FieldsEntry
	nil,                   // 10: com.coralogix.archive.v2.FieldNode.ChildrenEntry
}
var file_com_coralogix_archive_v2_schema_proto_depIdxs = []int32{
	0,  // 0: com.coralogix.archive.v2.SubSchema.namespace:type_name -> com.coralogix.archive.v2.Namespace
	9,  // 1: com.coralogix.archive.v2.SubSchema.fields:type_name -> com.coralogix.archive.v2.SubSchema.FieldsEntry
	2,  // 2: com.coralogix.archive.v2.SchemaTree.schemas:type_name -> com.coralogix.archive.v2.SubSchema
	10, // 3: com.coralogix.archive.v2.FieldNode.children:type_name -> com.coralogix.archive.v2.FieldNode.ChildrenEntry
	5,  // 4: com.coralogix.archive.v2.FieldNode.metadata:type_name -> com.coralogix.archive.v2.FieldMetadata
	6,  // 5: com.coralogix.archive.v2.FieldMetadata.location:type_name -> com.coralogix.archive.v2.FieldLocation
	7,  // 6: com.coralogix.archive.v2.FieldLocation.payload:type_name -> com.coralogix.archive.v2.Payload
	8,  // 7: com.coralogix.archive.v2.FieldLocation.column:type_name -> com.coralogix.archive.v2.Column
	4,  // 8: com.coralogix.archive.v2.SubSchema.FieldsEntry.value:type_name -> com.coralogix.archive.v2.FieldNode
	4,  // 9: com.coralogix.archive.v2.FieldNode.ChildrenEntry.value:type_name -> com.coralogix.archive.v2.FieldNode
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_com_coralogix_archive_v2_schema_proto_init() }
func file_com_coralogix_archive_v2_schema_proto_init() {
	if File_com_coralogix_archive_v2_schema_proto != nil {
		return
	}
	file_com_coralogix_archive_v2_schema_proto_msgTypes[4].OneofWrappers = []any{
		(*FieldLocation_Payload)(nil),
		(*FieldLocation_Column)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_archive_v2_schema_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_archive_v2_schema_proto_goTypes,
		DependencyIndexes: file_com_coralogix_archive_v2_schema_proto_depIdxs,
		EnumInfos:         file_com_coralogix_archive_v2_schema_proto_enumTypes,
		MessageInfos:      file_com_coralogix_archive_v2_schema_proto_msgTypes,
	}.Build()
	File_com_coralogix_archive_v2_schema_proto = out.File
	file_com_coralogix_archive_v2_schema_proto_rawDesc = nil
	file_com_coralogix_archive_v2_schema_proto_goTypes = nil
	file_com_coralogix_archive_v2_schema_proto_depIdxs = nil
}
