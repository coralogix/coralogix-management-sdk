// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: com/coralogix/dpxl/v1/dpxl.proto

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

type Expression struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Predicate string `protobuf:"bytes,1,opt,name=predicate,proto3" json:"predicate,omitempty"`
}

func (x *Expression) Reset() {
	*x = Expression{}
	mi := &file_com_coralogix_dpxl_v1_dpxl_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Expression) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Expression) ProtoMessage() {}

func (x *Expression) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dpxl_v1_dpxl_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Expression.ProtoReflect.Descriptor instead.
func (*Expression) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dpxl_v1_dpxl_proto_rawDescGZIP(), []int{0}
}

func (x *Expression) GetPredicate() string {
	if x != nil {
		return x.Predicate
	}
	return ""
}

type SerializedExpression struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *SerializedExpression) Reset() {
	*x = SerializedExpression{}
	mi := &file_com_coralogix_dpxl_v1_dpxl_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SerializedExpression) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SerializedExpression) ProtoMessage() {}

func (x *SerializedExpression) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dpxl_v1_dpxl_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SerializedExpression.ProtoReflect.Descriptor instead.
func (*SerializedExpression) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dpxl_v1_dpxl_proto_rawDescGZIP(), []int{1}
}

func (x *SerializedExpression) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type Diagnostic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Diagnostic) Reset() {
	*x = Diagnostic{}
	mi := &file_com_coralogix_dpxl_v1_dpxl_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Diagnostic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Diagnostic) ProtoMessage() {}

func (x *Diagnostic) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dpxl_v1_dpxl_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Diagnostic.ProtoReflect.Descriptor instead.
func (*Diagnostic) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dpxl_v1_dpxl_proto_rawDescGZIP(), []int{2}
}

type Source struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Source:
	//
	//	*Source_Logs_
	//	*Source_Spans_
	Source isSource_Source `protobuf_oneof:"source"`
}

func (x *Source) Reset() {
	*x = Source{}
	mi := &file_com_coralogix_dpxl_v1_dpxl_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Source) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Source) ProtoMessage() {}

func (x *Source) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dpxl_v1_dpxl_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Source.ProtoReflect.Descriptor instead.
func (*Source) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dpxl_v1_dpxl_proto_rawDescGZIP(), []int{3}
}

func (m *Source) GetSource() isSource_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *Source) GetLogs() *Source_Logs {
	if x, ok := x.GetSource().(*Source_Logs_); ok {
		return x.Logs
	}
	return nil
}

func (x *Source) GetSpans() *Source_Spans {
	if x, ok := x.GetSource().(*Source_Spans_); ok {
		return x.Spans
	}
	return nil
}

type isSource_Source interface {
	isSource_Source()
}

type Source_Logs_ struct {
	Logs *Source_Logs `protobuf:"bytes,1,opt,name=logs,proto3,oneof"`
}

type Source_Spans_ struct {
	Spans *Source_Spans `protobuf:"bytes,2,opt,name=spans,proto3,oneof"`
}

func (*Source_Logs_) isSource_Source() {}

func (*Source_Spans_) isSource_Source() {}

type EntityType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to EntityType:
	//
	//	*EntityType_Logs_
	//	*EntityType_Spans_
	EntityType isEntityType_EntityType `protobuf_oneof:"entity_type"`
}

func (x *EntityType) Reset() {
	*x = EntityType{}
	mi := &file_com_coralogix_dpxl_v1_dpxl_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntityType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityType) ProtoMessage() {}

func (x *EntityType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dpxl_v1_dpxl_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityType.ProtoReflect.Descriptor instead.
func (*EntityType) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dpxl_v1_dpxl_proto_rawDescGZIP(), []int{4}
}

func (m *EntityType) GetEntityType() isEntityType_EntityType {
	if m != nil {
		return m.EntityType
	}
	return nil
}

func (x *EntityType) GetLogs() *EntityType_Logs {
	if x, ok := x.GetEntityType().(*EntityType_Logs_); ok {
		return x.Logs
	}
	return nil
}

func (x *EntityType) GetSpans() *EntityType_Spans {
	if x, ok := x.GetEntityType().(*EntityType_Spans_); ok {
		return x.Spans
	}
	return nil
}

type isEntityType_EntityType interface {
	isEntityType_EntityType()
}

type EntityType_Logs_ struct {
	Logs *EntityType_Logs `protobuf:"bytes,1,opt,name=logs,proto3,oneof"`
}

type EntityType_Spans_ struct {
	Spans *EntityType_Spans `protobuf:"bytes,2,opt,name=spans,proto3,oneof"`
}

func (*EntityType_Logs_) isEntityType_EntityType() {}

func (*EntityType_Spans_) isEntityType_EntityType() {}

type Diagnostic_Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string            `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Range   *Diagnostic_Range `protobuf:"bytes,2,opt,name=range,proto3" json:"range,omitempty"`
}

func (x *Diagnostic_Error) Reset() {
	*x = Diagnostic_Error{}
	mi := &file_com_coralogix_dpxl_v1_dpxl_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Diagnostic_Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Diagnostic_Error) ProtoMessage() {}

func (x *Diagnostic_Error) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dpxl_v1_dpxl_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Diagnostic_Error.ProtoReflect.Descriptor instead.
func (*Diagnostic_Error) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dpxl_v1_dpxl_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Diagnostic_Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Diagnostic_Error) GetRange() *Diagnostic_Range {
	if x != nil {
		return x.Range
	}
	return nil
}

type Diagnostic_Warning struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string            `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Range   *Diagnostic_Range `protobuf:"bytes,2,opt,name=range,proto3" json:"range,omitempty"`
}

func (x *Diagnostic_Warning) Reset() {
	*x = Diagnostic_Warning{}
	mi := &file_com_coralogix_dpxl_v1_dpxl_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Diagnostic_Warning) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Diagnostic_Warning) ProtoMessage() {}

func (x *Diagnostic_Warning) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dpxl_v1_dpxl_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Diagnostic_Warning.ProtoReflect.Descriptor instead.
func (*Diagnostic_Warning) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dpxl_v1_dpxl_proto_rawDescGZIP(), []int{2, 1}
}

func (x *Diagnostic_Warning) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Diagnostic_Warning) GetRange() *Diagnostic_Range {
	if x != nil {
		return x.Range
	}
	return nil
}

type Diagnostic_Range struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From *Diagnostic_Position `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To   *Diagnostic_Position `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *Diagnostic_Range) Reset() {
	*x = Diagnostic_Range{}
	mi := &file_com_coralogix_dpxl_v1_dpxl_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Diagnostic_Range) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Diagnostic_Range) ProtoMessage() {}

func (x *Diagnostic_Range) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dpxl_v1_dpxl_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Diagnostic_Range.ProtoReflect.Descriptor instead.
func (*Diagnostic_Range) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dpxl_v1_dpxl_proto_rawDescGZIP(), []int{2, 2}
}

func (x *Diagnostic_Range) GetFrom() *Diagnostic_Position {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *Diagnostic_Range) GetTo() *Diagnostic_Position {
	if x != nil {
		return x.To
	}
	return nil
}

type Diagnostic_Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Line uint32 `protobuf:"varint,1,opt,name=line,proto3" json:"line,omitempty"`
	Col  uint32 `protobuf:"varint,2,opt,name=col,proto3" json:"col,omitempty"`
}

func (x *Diagnostic_Position) Reset() {
	*x = Diagnostic_Position{}
	mi := &file_com_coralogix_dpxl_v1_dpxl_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Diagnostic_Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Diagnostic_Position) ProtoMessage() {}

func (x *Diagnostic_Position) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dpxl_v1_dpxl_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Diagnostic_Position.ProtoReflect.Descriptor instead.
func (*Diagnostic_Position) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dpxl_v1_dpxl_proto_rawDescGZIP(), []int{2, 3}
}

func (x *Diagnostic_Position) GetLine() uint32 {
	if x != nil {
		return x.Line
	}
	return 0
}

func (x *Diagnostic_Position) GetCol() uint32 {
	if x != nil {
		return x.Col
	}
	return 0
}

type Source_Logs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Source_Logs) Reset() {
	*x = Source_Logs{}
	mi := &file_com_coralogix_dpxl_v1_dpxl_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Source_Logs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Source_Logs) ProtoMessage() {}

func (x *Source_Logs) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dpxl_v1_dpxl_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Source_Logs.ProtoReflect.Descriptor instead.
func (*Source_Logs) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dpxl_v1_dpxl_proto_rawDescGZIP(), []int{3, 0}
}

type Source_Spans struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Source_Spans) Reset() {
	*x = Source_Spans{}
	mi := &file_com_coralogix_dpxl_v1_dpxl_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Source_Spans) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Source_Spans) ProtoMessage() {}

func (x *Source_Spans) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dpxl_v1_dpxl_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Source_Spans.ProtoReflect.Descriptor instead.
func (*Source_Spans) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dpxl_v1_dpxl_proto_rawDescGZIP(), []int{3, 1}
}

type EntityType_Logs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EntityType_Logs) Reset() {
	*x = EntityType_Logs{}
	mi := &file_com_coralogix_dpxl_v1_dpxl_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntityType_Logs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityType_Logs) ProtoMessage() {}

func (x *EntityType_Logs) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dpxl_v1_dpxl_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityType_Logs.ProtoReflect.Descriptor instead.
func (*EntityType_Logs) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dpxl_v1_dpxl_proto_rawDescGZIP(), []int{4, 0}
}

type EntityType_Spans struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EntityType_Spans) Reset() {
	*x = EntityType_Spans{}
	mi := &file_com_coralogix_dpxl_v1_dpxl_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntityType_Spans) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityType_Spans) ProtoMessage() {}

func (x *EntityType_Spans) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dpxl_v1_dpxl_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityType_Spans.ProtoReflect.Descriptor instead.
func (*EntityType_Spans) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dpxl_v1_dpxl_proto_rawDescGZIP(), []int{4, 1}
}

var File_com_coralogix_dpxl_v1_dpxl_proto protoreflect.FileDescriptor

var file_com_coralogix_dpxl_v1_dpxl_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x64, 0x70, 0x78, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x70, 0x78, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2e, 0x64, 0x70, 0x78, 0x6c, 0x2e, 0x76, 0x31, 0x22, 0x2a, 0x0a, 0x0a, 0x45, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x64, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x22, 0x30, 0x0a, 0x14, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x64, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x8a, 0x03, 0x0a, 0x0a, 0x44, 0x69, 0x61, 0x67,
	0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x1a, 0x60, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3d, 0x0a, 0x05, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x70, 0x78, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x1a, 0x62, 0x0a, 0x07, 0x57, 0x61, 0x72, 0x6e,
	0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3d, 0x0a,
	0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x70, 0x78,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2e,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x1a, 0x83, 0x01, 0x0a,
	0x05, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x3e, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x70, 0x78, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x61,
	0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x3a, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x64, 0x70, 0x78, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x61, 0x67, 0x6e,
	0x6f, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x02,
	0x74, 0x6f, 0x1a, 0x30, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x6c, 0x69,
	0x6e, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x63, 0x6f, 0x6c, 0x22, 0x9a, 0x01, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x38, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x70,
	0x78, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x67,
	0x73, 0x48, 0x00, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x3b, 0x0a, 0x05, 0x73, 0x70, 0x61,
	0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x70, 0x78, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x73, 0x48, 0x00, 0x52,
	0x05, 0x73, 0x70, 0x61, 0x6e, 0x73, 0x1a, 0x06, 0x0a, 0x04, 0x4c, 0x6f, 0x67, 0x73, 0x1a, 0x07,
	0x0a, 0x05, 0x53, 0x70, 0x61, 0x6e, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x22, 0xb1, 0x01, 0x0a, 0x0a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x3c, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64,
	0x70, 0x78, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x48, 0x00, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x3f,
	0x0a, 0x05, 0x73, 0x70, 0x61, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x70,
	0x78, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x2e, 0x53, 0x70, 0x61, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x05, 0x73, 0x70, 0x61, 0x6e, 0x73, 0x1a,
	0x06, 0x0a, 0x04, 0x4c, 0x6f, 0x67, 0x73, 0x1a, 0x07, 0x0a, 0x05, 0x53, 0x70, 0x61, 0x6e, 0x73,
	0x42, 0x0d, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x4a,
	0x04, 0x08, 0x03, 0x10, 0x04, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_dpxl_v1_dpxl_proto_rawDescOnce sync.Once
	file_com_coralogix_dpxl_v1_dpxl_proto_rawDescData = file_com_coralogix_dpxl_v1_dpxl_proto_rawDesc
)

func file_com_coralogix_dpxl_v1_dpxl_proto_rawDescGZIP() []byte {
	file_com_coralogix_dpxl_v1_dpxl_proto_rawDescOnce.Do(func() {
		file_com_coralogix_dpxl_v1_dpxl_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_dpxl_v1_dpxl_proto_rawDescData)
	})
	return file_com_coralogix_dpxl_v1_dpxl_proto_rawDescData
}

var file_com_coralogix_dpxl_v1_dpxl_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_com_coralogix_dpxl_v1_dpxl_proto_goTypes = []any{
	(*Expression)(nil),           // 0: com.coralogix.dpxl.v1.Expression
	(*SerializedExpression)(nil), // 1: com.coralogix.dpxl.v1.SerializedExpression
	(*Diagnostic)(nil),           // 2: com.coralogix.dpxl.v1.Diagnostic
	(*Source)(nil),               // 3: com.coralogix.dpxl.v1.Source
	(*EntityType)(nil),           // 4: com.coralogix.dpxl.v1.EntityType
	(*Diagnostic_Error)(nil),     // 5: com.coralogix.dpxl.v1.Diagnostic.Error
	(*Diagnostic_Warning)(nil),   // 6: com.coralogix.dpxl.v1.Diagnostic.Warning
	(*Diagnostic_Range)(nil),     // 7: com.coralogix.dpxl.v1.Diagnostic.Range
	(*Diagnostic_Position)(nil),  // 8: com.coralogix.dpxl.v1.Diagnostic.Position
	(*Source_Logs)(nil),          // 9: com.coralogix.dpxl.v1.Source.Logs
	(*Source_Spans)(nil),         // 10: com.coralogix.dpxl.v1.Source.Spans
	(*EntityType_Logs)(nil),      // 11: com.coralogix.dpxl.v1.EntityType.Logs
	(*EntityType_Spans)(nil),     // 12: com.coralogix.dpxl.v1.EntityType.Spans
}
var file_com_coralogix_dpxl_v1_dpxl_proto_depIdxs = []int32{
	9,  // 0: com.coralogix.dpxl.v1.Source.logs:type_name -> com.coralogix.dpxl.v1.Source.Logs
	10, // 1: com.coralogix.dpxl.v1.Source.spans:type_name -> com.coralogix.dpxl.v1.Source.Spans
	11, // 2: com.coralogix.dpxl.v1.EntityType.logs:type_name -> com.coralogix.dpxl.v1.EntityType.Logs
	12, // 3: com.coralogix.dpxl.v1.EntityType.spans:type_name -> com.coralogix.dpxl.v1.EntityType.Spans
	7,  // 4: com.coralogix.dpxl.v1.Diagnostic.Error.range:type_name -> com.coralogix.dpxl.v1.Diagnostic.Range
	7,  // 5: com.coralogix.dpxl.v1.Diagnostic.Warning.range:type_name -> com.coralogix.dpxl.v1.Diagnostic.Range
	8,  // 6: com.coralogix.dpxl.v1.Diagnostic.Range.from:type_name -> com.coralogix.dpxl.v1.Diagnostic.Position
	8,  // 7: com.coralogix.dpxl.v1.Diagnostic.Range.to:type_name -> com.coralogix.dpxl.v1.Diagnostic.Position
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_com_coralogix_dpxl_v1_dpxl_proto_init() }
func file_com_coralogix_dpxl_v1_dpxl_proto_init() {
	if File_com_coralogix_dpxl_v1_dpxl_proto != nil {
		return
	}
	file_com_coralogix_dpxl_v1_dpxl_proto_msgTypes[3].OneofWrappers = []any{
		(*Source_Logs_)(nil),
		(*Source_Spans_)(nil),
	}
	file_com_coralogix_dpxl_v1_dpxl_proto_msgTypes[4].OneofWrappers = []any{
		(*EntityType_Logs_)(nil),
		(*EntityType_Spans_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_dpxl_v1_dpxl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_dpxl_v1_dpxl_proto_goTypes,
		DependencyIndexes: file_com_coralogix_dpxl_v1_dpxl_proto_depIdxs,
		MessageInfos:      file_com_coralogix_dpxl_v1_dpxl_proto_msgTypes,
	}.Build()
	File_com_coralogix_dpxl_v1_dpxl_proto = out.File
	file_com_coralogix_dpxl_v1_dpxl_proto_rawDesc = nil
	file_com_coralogix_dpxl_v1_dpxl_proto_goTypes = nil
	file_com_coralogix_dpxl_v1_dpxl_proto_depIdxs = nil
}
