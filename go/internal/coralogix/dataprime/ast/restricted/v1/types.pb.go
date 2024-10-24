// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: com/coralogix/dataprime/ast/restricted/v1/types.proto

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

type Datatype_PrimitiveType int32

const (
	Datatype_PRIMITIVE_TYPE_UNSPECIFIED Datatype_PrimitiveType = 0
	Datatype_PRIMITIVE_TYPE_BOOL        Datatype_PrimitiveType = 1
	Datatype_PRIMITIVE_TYPE_NUM         Datatype_PrimitiveType = 2
	Datatype_PRIMITIVE_TYPE_STRING      Datatype_PrimitiveType = 3
	Datatype_PRIMITIVE_TYPE_TIMESTAMP   Datatype_PrimitiveType = 4
)

// Enum value maps for Datatype_PrimitiveType.
var (
	Datatype_PrimitiveType_name = map[int32]string{
		0: "PRIMITIVE_TYPE_UNSPECIFIED",
		1: "PRIMITIVE_TYPE_BOOL",
		2: "PRIMITIVE_TYPE_NUM",
		3: "PRIMITIVE_TYPE_STRING",
		4: "PRIMITIVE_TYPE_TIMESTAMP",
	}
	Datatype_PrimitiveType_value = map[string]int32{
		"PRIMITIVE_TYPE_UNSPECIFIED": 0,
		"PRIMITIVE_TYPE_BOOL":        1,
		"PRIMITIVE_TYPE_NUM":         2,
		"PRIMITIVE_TYPE_STRING":      3,
		"PRIMITIVE_TYPE_TIMESTAMP":   4,
	}
)

func (x Datatype_PrimitiveType) Enum() *Datatype_PrimitiveType {
	p := new(Datatype_PrimitiveType)
	*p = x
	return p
}

func (x Datatype_PrimitiveType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Datatype_PrimitiveType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_dataprime_ast_restricted_v1_types_proto_enumTypes[0].Descriptor()
}

func (Datatype_PrimitiveType) Type() protoreflect.EnumType {
	return &file_com_coralogix_dataprime_ast_restricted_v1_types_proto_enumTypes[0]
}

func (x Datatype_PrimitiveType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Datatype_PrimitiveType.Descriptor instead.
func (Datatype_PrimitiveType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_ast_restricted_v1_types_proto_rawDescGZIP(), []int{0, 0}
}

type Datatype_CompositeType int32

const (
	Datatype_COMPOSITE_TYPE_UNSPECIFIED Datatype_CompositeType = 0
	Datatype_COMPOSITE_TYPE_ARRAY       Datatype_CompositeType = 1
	Datatype_COMPOSITE_TYPE_OBJECT      Datatype_CompositeType = 2
)

// Enum value maps for Datatype_CompositeType.
var (
	Datatype_CompositeType_name = map[int32]string{
		0: "COMPOSITE_TYPE_UNSPECIFIED",
		1: "COMPOSITE_TYPE_ARRAY",
		2: "COMPOSITE_TYPE_OBJECT",
	}
	Datatype_CompositeType_value = map[string]int32{
		"COMPOSITE_TYPE_UNSPECIFIED": 0,
		"COMPOSITE_TYPE_ARRAY":       1,
		"COMPOSITE_TYPE_OBJECT":      2,
	}
)

func (x Datatype_CompositeType) Enum() *Datatype_CompositeType {
	p := new(Datatype_CompositeType)
	*p = x
	return p
}

func (x Datatype_CompositeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Datatype_CompositeType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_dataprime_ast_restricted_v1_types_proto_enumTypes[1].Descriptor()
}

func (Datatype_CompositeType) Type() protoreflect.EnumType {
	return &file_com_coralogix_dataprime_ast_restricted_v1_types_proto_enumTypes[1]
}

func (x Datatype_CompositeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Datatype_CompositeType.Descriptor instead.
func (Datatype_CompositeType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_ast_restricted_v1_types_proto_rawDescGZIP(), []int{0, 1}
}

type Datatype struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Datatype:
	//
	//	*Datatype_PrimitiveType_
	//	*Datatype_CompositeType_
	Datatype isDatatype_Datatype `protobuf_oneof:"datatype"`
}

func (x *Datatype) Reset() {
	*x = Datatype{}
	mi := &file_com_coralogix_dataprime_ast_restricted_v1_types_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Datatype) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Datatype) ProtoMessage() {}

func (x *Datatype) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_ast_restricted_v1_types_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Datatype.ProtoReflect.Descriptor instead.
func (*Datatype) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_ast_restricted_v1_types_proto_rawDescGZIP(), []int{0}
}

func (m *Datatype) GetDatatype() isDatatype_Datatype {
	if m != nil {
		return m.Datatype
	}
	return nil
}

func (x *Datatype) GetPrimitiveType() Datatype_PrimitiveType {
	if x, ok := x.GetDatatype().(*Datatype_PrimitiveType_); ok {
		return x.PrimitiveType
	}
	return Datatype_PRIMITIVE_TYPE_UNSPECIFIED
}

func (x *Datatype) GetCompositeType() Datatype_CompositeType {
	if x, ok := x.GetDatatype().(*Datatype_CompositeType_); ok {
		return x.CompositeType
	}
	return Datatype_COMPOSITE_TYPE_UNSPECIFIED
}

type isDatatype_Datatype interface {
	isDatatype_Datatype()
}

type Datatype_PrimitiveType_ struct {
	PrimitiveType Datatype_PrimitiveType `protobuf:"varint,1,opt,name=primitive_type,json=primitiveType,proto3,enum=com.coralogix.dataprime.ast.restricted.v1.Datatype_PrimitiveType,oneof"`
}

type Datatype_CompositeType_ struct {
	CompositeType Datatype_CompositeType `protobuf:"varint,2,opt,name=composite_type,json=compositeType,proto3,enum=com.coralogix.dataprime.ast.restricted.v1.Datatype_CompositeType,oneof"`
}

func (*Datatype_PrimitiveType_) isDatatype_Datatype() {}

func (*Datatype_CompositeType_) isDatatype_Datatype() {}

var File_com_coralogix_dataprime_ast_restricted_v1_types_proto protoreflect.FileDescriptor

var file_com_coralogix_dataprime_ast_restricted_v1_types_proto_rawDesc = []byte{
	0x0a, 0x35, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x61, 0x73, 0x74, 0x2f, 0x72, 0x65,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65,
	0x2e, 0x61, 0x73, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x65, 0x64, 0x2e,
	0x76, 0x31, 0x22, 0xf0, 0x03, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x6a, 0x0a, 0x0e, 0x70, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x41, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d,
	0x65, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x65, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x72, 0x69,
	0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x70, 0x72,
	0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x6a, 0x0a, 0x0e, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x41, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x61, 0x73,
	0x74, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x99, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x69, 0x6d,
	0x69, 0x74, 0x69, 0x76, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x52, 0x49,
	0x4d, 0x49, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x52, 0x49,
	0x4d, 0x49, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4c,
	0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x52, 0x49, 0x4d, 0x49, 0x54, 0x49, 0x56, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x55, 0x4d, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x52,
	0x49, 0x4d, 0x49, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x52,
	0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x52, 0x49, 0x4d, 0x49, 0x54, 0x49,
	0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d,
	0x50, 0x10, 0x04, 0x22, 0x64, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x4f, 0x4d, 0x50, 0x4f, 0x53, 0x49, 0x54,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x4f, 0x4d, 0x50, 0x4f, 0x53, 0x49, 0x54,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x52, 0x52, 0x41, 0x59, 0x10, 0x01, 0x12, 0x19,
	0x0a, 0x15, 0x43, 0x4f, 0x4d, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x02, 0x42, 0x0a, 0x0a, 0x08, 0x64, 0x61, 0x74,
	0x61, 0x74, 0x79, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_dataprime_ast_restricted_v1_types_proto_rawDescOnce sync.Once
	file_com_coralogix_dataprime_ast_restricted_v1_types_proto_rawDescData = file_com_coralogix_dataprime_ast_restricted_v1_types_proto_rawDesc
)

func file_com_coralogix_dataprime_ast_restricted_v1_types_proto_rawDescGZIP() []byte {
	file_com_coralogix_dataprime_ast_restricted_v1_types_proto_rawDescOnce.Do(func() {
		file_com_coralogix_dataprime_ast_restricted_v1_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_dataprime_ast_restricted_v1_types_proto_rawDescData)
	})
	return file_com_coralogix_dataprime_ast_restricted_v1_types_proto_rawDescData
}

var file_com_coralogix_dataprime_ast_restricted_v1_types_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_com_coralogix_dataprime_ast_restricted_v1_types_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogix_dataprime_ast_restricted_v1_types_proto_goTypes = []any{
	(Datatype_PrimitiveType)(0), // 0: com.coralogix.dataprime.ast.restricted.v1.Datatype.PrimitiveType
	(Datatype_CompositeType)(0), // 1: com.coralogix.dataprime.ast.restricted.v1.Datatype.CompositeType
	(*Datatype)(nil),            // 2: com.coralogix.dataprime.ast.restricted.v1.Datatype
}
var file_com_coralogix_dataprime_ast_restricted_v1_types_proto_depIdxs = []int32{
	0, // 0: com.coralogix.dataprime.ast.restricted.v1.Datatype.primitive_type:type_name -> com.coralogix.dataprime.ast.restricted.v1.Datatype.PrimitiveType
	1, // 1: com.coralogix.dataprime.ast.restricted.v1.Datatype.composite_type:type_name -> com.coralogix.dataprime.ast.restricted.v1.Datatype.CompositeType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_com_coralogix_dataprime_ast_restricted_v1_types_proto_init() }
func file_com_coralogix_dataprime_ast_restricted_v1_types_proto_init() {
	if File_com_coralogix_dataprime_ast_restricted_v1_types_proto != nil {
		return
	}
	file_com_coralogix_dataprime_ast_restricted_v1_types_proto_msgTypes[0].OneofWrappers = []any{
		(*Datatype_PrimitiveType_)(nil),
		(*Datatype_CompositeType_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_dataprime_ast_restricted_v1_types_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_dataprime_ast_restricted_v1_types_proto_goTypes,
		DependencyIndexes: file_com_coralogix_dataprime_ast_restricted_v1_types_proto_depIdxs,
		EnumInfos:         file_com_coralogix_dataprime_ast_restricted_v1_types_proto_enumTypes,
		MessageInfos:      file_com_coralogix_dataprime_ast_restricted_v1_types_proto_msgTypes,
	}.Build()
	File_com_coralogix_dataprime_ast_restricted_v1_types_proto = out.File
	file_com_coralogix_dataprime_ast_restricted_v1_types_proto_rawDesc = nil
	file_com_coralogix_dataprime_ast_restricted_v1_types_proto_goTypes = nil
	file_com_coralogix_dataprime_ast_restricted_v1_types_proto_depIdxs = nil
}
