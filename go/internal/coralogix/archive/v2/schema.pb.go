// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.1
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
	SchemaType_SCHEMA_TYPE_BINARY           SchemaType = 4
)

// Enum value maps for SchemaType.
var (
	SchemaType_name = map[int32]string{
		0: "SCHEMA_TYPE_UNSPECIFIED",
		1: "SCHEMA_TYPE_STRING",
		2: "SCHEMA_TYPE_TIMESTAMP_NANOS",
		3: "SCHEMA_TYPE_TIMESTAMP_MICROS",
		4: "SCHEMA_TYPE_BINARY",
	}
	SchemaType_value = map[string]int32{
		"SCHEMA_TYPE_UNSPECIFIED":      0,
		"SCHEMA_TYPE_STRING":           1,
		"SCHEMA_TYPE_TIMESTAMP_NANOS":  2,
		"SCHEMA_TYPE_TIMESTAMP_MICROS": 3,
		"SCHEMA_TYPE_BINARY":           4,
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

type Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path        []string   `protobuf:"bytes,2,rep,name=path,proto3" json:"path,omitempty"`
	Occurrences int32      `protobuf:"varint,3,opt,name=occurrences,proto3" json:"occurrences,omitempty"`
	Type        SchemaType `protobuf:"varint,4,opt,name=type,proto3,enum=com.coralogix.archive.v2.SchemaType" json:"type,omitempty"`
}

func (x *Field) Reset() {
	*x = Field{}
	mi := &file_com_coralogix_archive_v2_schema_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Field.ProtoReflect.Descriptor instead.
func (*Field) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_schema_proto_rawDescGZIP(), []int{0}
}

func (x *Field) GetPath() []string {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *Field) GetOccurrences() int32 {
	if x != nil {
		return x.Occurrences
	}
	return 0
}

func (x *Field) GetType() SchemaType {
	if x != nil {
		return x.Type
	}
	return SchemaType_SCHEMA_TYPE_UNSPECIFIED
}

type SubSchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace Namespace `protobuf:"varint,1,opt,name=namespace,proto3,enum=com.coralogix.archive.v2.Namespace" json:"namespace,omitempty"`
	Fields    []*Field  `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *SubSchema) Reset() {
	*x = SubSchema{}
	mi := &file_com_coralogix_archive_v2_schema_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubSchema) ProtoMessage() {}

func (x *SubSchema) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SubSchema.ProtoReflect.Descriptor instead.
func (*SubSchema) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_schema_proto_rawDescGZIP(), []int{1}
}

func (x *SubSchema) GetNamespace() Namespace {
	if x != nil {
		return x.Namespace
	}
	return Namespace_NAMESPACE_UNSPECIFIED
}

func (x *SubSchema) GetFields() []*Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

type Schema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schemas []*SubSchema `protobuf:"bytes,1,rep,name=schemas,proto3" json:"schemas,omitempty"`
}

func (x *Schema) Reset() {
	*x = Schema{}
	mi := &file_com_coralogix_archive_v2_schema_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Schema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schema) ProtoMessage() {}

func (x *Schema) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Schema.ProtoReflect.Descriptor instead.
func (*Schema) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_schema_proto_rawDescGZIP(), []int{2}
}

func (x *Schema) GetSchemas() []*SubSchema {
	if x != nil {
		return x.Schemas
	}
	return nil
}

var File_com_coralogix_archive_v2_schema_proto protoreflect.FileDescriptor

var file_com_coralogix_archive_v2_schema_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76,
	0x32, 0x22, 0x77, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x20,
	0x0a, 0x0b, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73,
	0x12, 0x38, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x87, 0x01, 0x0a, 0x09, 0x53,
	0x75, 0x62, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x41, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x06, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x22, 0x47, 0x0a, 0x06, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x3d,
	0x0a, 0x07, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x75, 0x62, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x52, 0x07, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2a, 0x6d, 0x0a,
	0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x4e, 0x41,
	0x4d, 0x45, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x4e, 0x41, 0x4d, 0x45, 0x53, 0x50, 0x41,
	0x43, 0x45, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x44, 0x41, 0x54, 0x41, 0x10, 0x01, 0x12, 0x14, 0x0a,
	0x10, 0x4e, 0x41, 0x4d, 0x45, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x4c, 0x41, 0x42, 0x45, 0x4c,
	0x53, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x4e, 0x41, 0x4d, 0x45, 0x53, 0x50, 0x41, 0x43, 0x45,
	0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0x03, 0x2a, 0x9c, 0x01, 0x0a,
	0x0a, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x53,
	0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x43, 0x48, 0x45,
	0x4d, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x01,
	0x12, 0x1f, 0x0a, 0x1b, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x5f, 0x4e, 0x41, 0x4e, 0x4f, 0x53, 0x10,
	0x02, 0x12, 0x20, 0x0a, 0x1c, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x5f, 0x4d, 0x49, 0x43, 0x52, 0x4f,
	0x53, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x42, 0x49, 0x4e, 0x41, 0x52, 0x59, 0x10, 0x04, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
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
var file_com_coralogix_archive_v2_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogix_archive_v2_schema_proto_goTypes = []any{
	(Namespace)(0),    // 0: com.coralogix.archive.v2.Namespace
	(SchemaType)(0),   // 1: com.coralogix.archive.v2.SchemaType
	(*Field)(nil),     // 2: com.coralogix.archive.v2.Field
	(*SubSchema)(nil), // 3: com.coralogix.archive.v2.SubSchema
	(*Schema)(nil),    // 4: com.coralogix.archive.v2.Schema
}
var file_com_coralogix_archive_v2_schema_proto_depIdxs = []int32{
	1, // 0: com.coralogix.archive.v2.Field.type:type_name -> com.coralogix.archive.v2.SchemaType
	0, // 1: com.coralogix.archive.v2.SubSchema.namespace:type_name -> com.coralogix.archive.v2.Namespace
	2, // 2: com.coralogix.archive.v2.SubSchema.fields:type_name -> com.coralogix.archive.v2.Field
	3, // 3: com.coralogix.archive.v2.Schema.schemas:type_name -> com.coralogix.archive.v2.SubSchema
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_com_coralogix_archive_v2_schema_proto_init() }
func file_com_coralogix_archive_v2_schema_proto_init() {
	if File_com_coralogix_archive_v2_schema_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_archive_v2_schema_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
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
