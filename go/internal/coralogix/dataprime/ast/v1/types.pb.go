// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.28.1
// source: com/coralogix/dataprime/ast/v1/types.proto

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
	Datatype_PRIMITIVE_TYPE_REGEX       Datatype_PrimitiveType = 5
	Datatype_PRIMITIVE_TYPE_INTERVAL    Datatype_PrimitiveType = 6
	Datatype_PRIMITIVE_TYPE_DATATYPE    Datatype_PrimitiveType = 7
	Datatype_PRIMITIVE_TYPE_TIME_UNIT   Datatype_PrimitiveType = 8
	Datatype_PRIMITIVE_TYPE_DATE_UNIT   Datatype_PrimitiveType = 9
	// The type of `null` (but also a bottom type).
	Datatype_PRIMITIVE_TYPE_NULL Datatype_PrimitiveType = 10
)

// Enum value maps for Datatype_PrimitiveType.
var (
	Datatype_PrimitiveType_name = map[int32]string{
		0:  "PRIMITIVE_TYPE_UNSPECIFIED",
		1:  "PRIMITIVE_TYPE_BOOL",
		2:  "PRIMITIVE_TYPE_NUM",
		3:  "PRIMITIVE_TYPE_STRING",
		4:  "PRIMITIVE_TYPE_TIMESTAMP",
		5:  "PRIMITIVE_TYPE_REGEX",
		6:  "PRIMITIVE_TYPE_INTERVAL",
		7:  "PRIMITIVE_TYPE_DATATYPE",
		8:  "PRIMITIVE_TYPE_TIME_UNIT",
		9:  "PRIMITIVE_TYPE_DATE_UNIT",
		10: "PRIMITIVE_TYPE_NULL",
	}
	Datatype_PrimitiveType_value = map[string]int32{
		"PRIMITIVE_TYPE_UNSPECIFIED": 0,
		"PRIMITIVE_TYPE_BOOL":        1,
		"PRIMITIVE_TYPE_NUM":         2,
		"PRIMITIVE_TYPE_STRING":      3,
		"PRIMITIVE_TYPE_TIMESTAMP":   4,
		"PRIMITIVE_TYPE_REGEX":       5,
		"PRIMITIVE_TYPE_INTERVAL":    6,
		"PRIMITIVE_TYPE_DATATYPE":    7,
		"PRIMITIVE_TYPE_TIME_UNIT":   8,
		"PRIMITIVE_TYPE_DATE_UNIT":   9,
		"PRIMITIVE_TYPE_NULL":        10,
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
	return file_com_coralogix_dataprime_ast_v1_types_proto_enumTypes[0].Descriptor()
}

func (Datatype_PrimitiveType) Type() protoreflect.EnumType {
	return &file_com_coralogix_dataprime_ast_v1_types_proto_enumTypes[0]
}

func (x Datatype_PrimitiveType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Datatype_PrimitiveType.Descriptor instead.
func (Datatype_PrimitiveType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_ast_v1_types_proto_rawDescGZIP(), []int{0, 0}
}

type Datatype_CompositeType int32

const (
	Datatype_COMPOSITE_TYPE_UNSPECIFIED Datatype_CompositeType = 0
	Datatype_COMPOSITE_TYPE_OBJECT      Datatype_CompositeType = 2
)

// Enum value maps for Datatype_CompositeType.
var (
	Datatype_CompositeType_name = map[int32]string{
		0: "COMPOSITE_TYPE_UNSPECIFIED",
		2: "COMPOSITE_TYPE_OBJECT",
	}
	Datatype_CompositeType_value = map[string]int32{
		"COMPOSITE_TYPE_UNSPECIFIED": 0,
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
	return file_com_coralogix_dataprime_ast_v1_types_proto_enumTypes[1].Descriptor()
}

func (Datatype_CompositeType) Type() protoreflect.EnumType {
	return &file_com_coralogix_dataprime_ast_v1_types_proto_enumTypes[1]
}

func (x Datatype_CompositeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Datatype_CompositeType.Descriptor instead.
func (Datatype_CompositeType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_ast_v1_types_proto_rawDescGZIP(), []int{0, 1}
}

type Datatype struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Datatype:
	//
	//	*Datatype_PrimitiveType_
	//	*Datatype_CompositeType_
	//	*Datatype_UnionType_
	//	*Datatype_ArrayType_
	//	*Datatype_EnumType_
	Datatype isDatatype_Datatype `protobuf_oneof:"datatype"`
}

func (x *Datatype) Reset() {
	*x = Datatype{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_dataprime_ast_v1_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Datatype) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Datatype) ProtoMessage() {}

func (x *Datatype) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_ast_v1_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_com_coralogix_dataprime_ast_v1_types_proto_rawDescGZIP(), []int{0}
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

func (x *Datatype) GetUnionType() *Datatype_UnionType {
	if x, ok := x.GetDatatype().(*Datatype_UnionType_); ok {
		return x.UnionType
	}
	return nil
}

func (x *Datatype) GetArrayType() *Datatype_ArrayType {
	if x, ok := x.GetDatatype().(*Datatype_ArrayType_); ok {
		return x.ArrayType
	}
	return nil
}

func (x *Datatype) GetEnumType() *Datatype_EnumType {
	if x, ok := x.GetDatatype().(*Datatype_EnumType_); ok {
		return x.EnumType
	}
	return nil
}

type isDatatype_Datatype interface {
	isDatatype_Datatype()
}

type Datatype_PrimitiveType_ struct {
	PrimitiveType Datatype_PrimitiveType `protobuf:"varint,1,opt,name=primitive_type,json=primitiveType,proto3,enum=com.coralogix.dataprime.ast.v1.Datatype_PrimitiveType,oneof"`
}

type Datatype_CompositeType_ struct {
	CompositeType Datatype_CompositeType `protobuf:"varint,2,opt,name=composite_type,json=compositeType,proto3,enum=com.coralogix.dataprime.ast.v1.Datatype_CompositeType,oneof"`
}

type Datatype_UnionType_ struct {
	UnionType *Datatype_UnionType `protobuf:"bytes,3,opt,name=union_type,json=unionType,proto3,oneof"`
}

type Datatype_ArrayType_ struct {
	ArrayType *Datatype_ArrayType `protobuf:"bytes,4,opt,name=array_type,json=arrayType,proto3,oneof"`
}

type Datatype_EnumType_ struct {
	EnumType *Datatype_EnumType `protobuf:"bytes,5,opt,name=enum_type,json=enumType,proto3,oneof"`
}

func (*Datatype_PrimitiveType_) isDatatype_Datatype() {}

func (*Datatype_CompositeType_) isDatatype_Datatype() {}

func (*Datatype_UnionType_) isDatatype_Datatype() {}

func (*Datatype_ArrayType_) isDatatype_Datatype() {}

func (*Datatype_EnumType_) isDatatype_Datatype() {}

type Datatype_UnionType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Types []*Datatype `protobuf:"bytes,1,rep,name=types,proto3" json:"types,omitempty"`
}

func (x *Datatype_UnionType) Reset() {
	*x = Datatype_UnionType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_dataprime_ast_v1_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Datatype_UnionType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Datatype_UnionType) ProtoMessage() {}

func (x *Datatype_UnionType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_ast_v1_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Datatype_UnionType.ProtoReflect.Descriptor instead.
func (*Datatype_UnionType) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_ast_v1_types_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Datatype_UnionType) GetTypes() []*Datatype {
	if x != nil {
		return x.Types
	}
	return nil
}

type Datatype_ArrayType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// No element type is equivalent to Any (the top type).
	Element *Datatype `protobuf:"bytes,1,opt,name=element,proto3,oneof" json:"element,omitempty"`
}

func (x *Datatype_ArrayType) Reset() {
	*x = Datatype_ArrayType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_dataprime_ast_v1_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Datatype_ArrayType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Datatype_ArrayType) ProtoMessage() {}

func (x *Datatype_ArrayType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_ast_v1_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Datatype_ArrayType.ProtoReflect.Descriptor instead.
func (*Datatype_ArrayType) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_ast_v1_types_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Datatype_ArrayType) GetElement() *Datatype {
	if x != nil {
		return x.Element
	}
	return nil
}

type Datatype_EnumType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Datatype_EnumType) Reset() {
	*x = Datatype_EnumType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_dataprime_ast_v1_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Datatype_EnumType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Datatype_EnumType) ProtoMessage() {}

func (x *Datatype_EnumType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_ast_v1_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Datatype_EnumType.ProtoReflect.Descriptor instead.
func (*Datatype_EnumType) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_ast_v1_types_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Datatype_EnumType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_com_coralogix_dataprime_ast_v1_types_proto protoreflect.FileDescriptor

var file_com_coralogix_dataprime_ast_v1_types_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x61, 0x73, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x22, 0xba, 0x08, 0x0a,
	0x08, 0x44, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x12, 0x5f, 0x0a, 0x0e, 0x70, 0x72, 0x69,
	0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x61, 0x73, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x72, 0x69, 0x6d,
	0x69, 0x74, 0x69, 0x76, 0x65, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x70, 0x72, 0x69,
	0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x5f, 0x0a, 0x0e, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x61, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x53, 0x0a, 0x0a, 0x75,
	0x6e, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x09, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x53, 0x0a, 0x0a, 0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x61,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x41,
	0x72, 0x72, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x09, 0x61, 0x72, 0x72, 0x61,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x50, 0x0a, 0x09, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69,
	0x6d, 0x65, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x08, 0x65,
	0x6e, 0x75, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x4b, 0x0a, 0x09, 0x55, 0x6e, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x3e, 0x0a, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x61, 0x73,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x52, 0x05, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x1a, 0x60, 0x0a, 0x09, 0x41, 0x72, 0x72, 0x61, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x47, 0x0a, 0x07, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x61, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x07,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x1e, 0x0a, 0x08, 0x45, 0x6e, 0x75, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xc2, 0x02, 0x0a, 0x0d, 0x50, 0x72, 0x69, 0x6d, 0x69,
	0x74, 0x69, 0x76, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x52, 0x49, 0x4d,
	0x49, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x52, 0x49, 0x4d,
	0x49, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4c, 0x10,
	0x01, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x52, 0x49, 0x4d, 0x49, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4e, 0x55, 0x4d, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x52, 0x49,
	0x4d, 0x49, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x49,
	0x4e, 0x47, 0x10, 0x03, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x52, 0x49, 0x4d, 0x49, 0x54, 0x49, 0x56,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50,
	0x10, 0x04, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x52, 0x49, 0x4d, 0x49, 0x54, 0x49, 0x56, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x47, 0x45, 0x58, 0x10, 0x05, 0x12, 0x1b, 0x0a, 0x17,
	0x50, 0x52, 0x49, 0x4d, 0x49, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49,
	0x4e, 0x54, 0x45, 0x52, 0x56, 0x41, 0x4c, 0x10, 0x06, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x52, 0x49,
	0x4d, 0x49, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x41,
	0x54, 0x59, 0x50, 0x45, 0x10, 0x07, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x52, 0x49, 0x4d, 0x49, 0x54,
	0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x55, 0x4e,
	0x49, 0x54, 0x10, 0x08, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x52, 0x49, 0x4d, 0x49, 0x54, 0x49, 0x56,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x49, 0x54,
	0x10, 0x09, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x52, 0x49, 0x4d, 0x49, 0x54, 0x49, 0x56, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x55, 0x4c, 0x4c, 0x10, 0x0a, 0x22, 0x50, 0x0a, 0x0d, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x1a,
	0x43, 0x4f, 0x4d, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15,
	0x43, 0x4f, 0x4d, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f,
	0x42, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x02, 0x22, 0x04, 0x08, 0x01, 0x10, 0x01, 0x42, 0x0a, 0x0a,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_com_coralogix_dataprime_ast_v1_types_proto_rawDescOnce sync.Once
	file_com_coralogix_dataprime_ast_v1_types_proto_rawDescData = file_com_coralogix_dataprime_ast_v1_types_proto_rawDesc
)

func file_com_coralogix_dataprime_ast_v1_types_proto_rawDescGZIP() []byte {
	file_com_coralogix_dataprime_ast_v1_types_proto_rawDescOnce.Do(func() {
		file_com_coralogix_dataprime_ast_v1_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_dataprime_ast_v1_types_proto_rawDescData)
	})
	return file_com_coralogix_dataprime_ast_v1_types_proto_rawDescData
}

var file_com_coralogix_dataprime_ast_v1_types_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_com_coralogix_dataprime_ast_v1_types_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_com_coralogix_dataprime_ast_v1_types_proto_goTypes = []interface{}{
	(Datatype_PrimitiveType)(0), // 0: com.coralogix.dataprime.ast.v1.Datatype.PrimitiveType
	(Datatype_CompositeType)(0), // 1: com.coralogix.dataprime.ast.v1.Datatype.CompositeType
	(*Datatype)(nil),            // 2: com.coralogix.dataprime.ast.v1.Datatype
	(*Datatype_UnionType)(nil),  // 3: com.coralogix.dataprime.ast.v1.Datatype.UnionType
	(*Datatype_ArrayType)(nil),  // 4: com.coralogix.dataprime.ast.v1.Datatype.ArrayType
	(*Datatype_EnumType)(nil),   // 5: com.coralogix.dataprime.ast.v1.Datatype.EnumType
}
var file_com_coralogix_dataprime_ast_v1_types_proto_depIdxs = []int32{
	0, // 0: com.coralogix.dataprime.ast.v1.Datatype.primitive_type:type_name -> com.coralogix.dataprime.ast.v1.Datatype.PrimitiveType
	1, // 1: com.coralogix.dataprime.ast.v1.Datatype.composite_type:type_name -> com.coralogix.dataprime.ast.v1.Datatype.CompositeType
	3, // 2: com.coralogix.dataprime.ast.v1.Datatype.union_type:type_name -> com.coralogix.dataprime.ast.v1.Datatype.UnionType
	4, // 3: com.coralogix.dataprime.ast.v1.Datatype.array_type:type_name -> com.coralogix.dataprime.ast.v1.Datatype.ArrayType
	5, // 4: com.coralogix.dataprime.ast.v1.Datatype.enum_type:type_name -> com.coralogix.dataprime.ast.v1.Datatype.EnumType
	2, // 5: com.coralogix.dataprime.ast.v1.Datatype.UnionType.types:type_name -> com.coralogix.dataprime.ast.v1.Datatype
	2, // 6: com.coralogix.dataprime.ast.v1.Datatype.ArrayType.element:type_name -> com.coralogix.dataprime.ast.v1.Datatype
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_com_coralogix_dataprime_ast_v1_types_proto_init() }
func file_com_coralogix_dataprime_ast_v1_types_proto_init() {
	if File_com_coralogix_dataprime_ast_v1_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_dataprime_ast_v1_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Datatype); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_com_coralogix_dataprime_ast_v1_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Datatype_UnionType); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_com_coralogix_dataprime_ast_v1_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Datatype_ArrayType); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_com_coralogix_dataprime_ast_v1_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Datatype_EnumType); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_com_coralogix_dataprime_ast_v1_types_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Datatype_PrimitiveType_)(nil),
		(*Datatype_CompositeType_)(nil),
		(*Datatype_UnionType_)(nil),
		(*Datatype_ArrayType_)(nil),
		(*Datatype_EnumType_)(nil),
	}
	file_com_coralogix_dataprime_ast_v1_types_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_dataprime_ast_v1_types_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_dataprime_ast_v1_types_proto_goTypes,
		DependencyIndexes: file_com_coralogix_dataprime_ast_v1_types_proto_depIdxs,
		EnumInfos:         file_com_coralogix_dataprime_ast_v1_types_proto_enumTypes,
		MessageInfos:      file_com_coralogix_dataprime_ast_v1_types_proto_msgTypes,
	}.Build()
	File_com_coralogix_dataprime_ast_v1_types_proto = out.File
	file_com_coralogix_dataprime_ast_v1_types_proto_rawDesc = nil
	file_com_coralogix_dataprime_ast_v1_types_proto_goTypes = nil
	file_com_coralogix_dataprime_ast_v1_types_proto_depIdxs = nil
}
