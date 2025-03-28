// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: com/coralogix/datausage/v2/data_usage.proto

package v2

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Pillar int32

const (
	Pillar_PILLAR_UNSPECIFIED Pillar = 0
	Pillar_PILLAR_METRICS     Pillar = 1
	Pillar_PILLAR_LOGS        Pillar = 2
	Pillar_PILLAR_SPANS       Pillar = 3
	Pillar_PILLAR_BINARY      Pillar = 4
)

// Enum value maps for Pillar.
var (
	Pillar_name = map[int32]string{
		0: "PILLAR_UNSPECIFIED",
		1: "PILLAR_METRICS",
		2: "PILLAR_LOGS",
		3: "PILLAR_SPANS",
		4: "PILLAR_BINARY",
	}
	Pillar_value = map[string]int32{
		"PILLAR_UNSPECIFIED": 0,
		"PILLAR_METRICS":     1,
		"PILLAR_LOGS":        2,
		"PILLAR_SPANS":       3,
		"PILLAR_BINARY":      4,
	}
)

func (x Pillar) Enum() *Pillar {
	p := new(Pillar)
	*p = x
	return p
}

func (x Pillar) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Pillar) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_datausage_v2_data_usage_proto_enumTypes[0].Descriptor()
}

func (Pillar) Type() protoreflect.EnumType {
	return &file_com_coralogix_datausage_v2_data_usage_proto_enumTypes[0]
}

func (x Pillar) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Pillar.Descriptor instead.
func (Pillar) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_datausage_v2_data_usage_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Marked as deprecated in com/coralogix/datausage/v2/data_usage.proto.
type TcoTier int32

const (
	TcoTier_TCO_TIER_UNSPECIFIED TcoTier = 0
	TcoTier_TCO_TIER_LOW         TcoTier = 1
	TcoTier_TCO_TIER_MEDIUM      TcoTier = 2
	TcoTier_TCO_TIER_HIGH        TcoTier = 3
	TcoTier_TCO_TIER_BLOCKED     TcoTier = 4
)

// Enum value maps for TcoTier.
var (
	TcoTier_name = map[int32]string{
		0: "TCO_TIER_UNSPECIFIED",
		1: "TCO_TIER_LOW",
		2: "TCO_TIER_MEDIUM",
		3: "TCO_TIER_HIGH",
		4: "TCO_TIER_BLOCKED",
	}
	TcoTier_value = map[string]int32{
		"TCO_TIER_UNSPECIFIED": 0,
		"TCO_TIER_LOW":         1,
		"TCO_TIER_MEDIUM":      2,
		"TCO_TIER_HIGH":        3,
		"TCO_TIER_BLOCKED":     4,
	}
)

func (x TcoTier) Enum() *TcoTier {
	p := new(TcoTier)
	*p = x
	return p
}

func (x TcoTier) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TcoTier) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_datausage_v2_data_usage_proto_enumTypes[1].Descriptor()
}

func (TcoTier) Type() protoreflect.EnumType {
	return &file_com_coralogix_datausage_v2_data_usage_proto_enumTypes[1]
}

func (x TcoTier) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TcoTier.Descriptor instead.
func (TcoTier) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_datausage_v2_data_usage_proto_rawDescGZIP(), []int{1}
}

type Priority int32

const (
	Priority_PRIORITY_UNSPECIFIED Priority = 0
	Priority_PRIORITY_LOW         Priority = 1
	Priority_PRIORITY_MEDIUM      Priority = 2
	Priority_PRIORITY_HIGH        Priority = 3
	Priority_PRIORITY_BLOCKED     Priority = 4
)

// Enum value maps for Priority.
var (
	Priority_name = map[int32]string{
		0: "PRIORITY_UNSPECIFIED",
		1: "PRIORITY_LOW",
		2: "PRIORITY_MEDIUM",
		3: "PRIORITY_HIGH",
		4: "PRIORITY_BLOCKED",
	}
	Priority_value = map[string]int32{
		"PRIORITY_UNSPECIFIED": 0,
		"PRIORITY_LOW":         1,
		"PRIORITY_MEDIUM":      2,
		"PRIORITY_HIGH":        3,
		"PRIORITY_BLOCKED":     4,
	}
)

func (x Priority) Enum() *Priority {
	p := new(Priority)
	*p = x
	return p
}

func (x Priority) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Priority) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_datausage_v2_data_usage_proto_enumTypes[2].Descriptor()
}

func (Priority) Type() protoreflect.EnumType {
	return &file_com_coralogix_datausage_v2_data_usage_proto_enumTypes[2]
}

func (x Priority) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Priority.Descriptor instead.
func (Priority) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_datausage_v2_data_usage_proto_rawDescGZIP(), []int{2}
}

type Severity int32

const (
	Severity_SEVERITY_UNSPECIFIED Severity = 0
	Severity_SEVERITY_DEBUG       Severity = 1
	Severity_SEVERITY_VERBOSE     Severity = 2
	Severity_SEVERITY_INFO        Severity = 3
	Severity_SEVERITY_WARNING     Severity = 4
	Severity_SEVERITY_ERROR       Severity = 5
	Severity_SEVERITY_CRITICAL    Severity = 6
)

// Enum value maps for Severity.
var (
	Severity_name = map[int32]string{
		0: "SEVERITY_UNSPECIFIED",
		1: "SEVERITY_DEBUG",
		2: "SEVERITY_VERBOSE",
		3: "SEVERITY_INFO",
		4: "SEVERITY_WARNING",
		5: "SEVERITY_ERROR",
		6: "SEVERITY_CRITICAL",
	}
	Severity_value = map[string]int32{
		"SEVERITY_UNSPECIFIED": 0,
		"SEVERITY_DEBUG":       1,
		"SEVERITY_VERBOSE":     2,
		"SEVERITY_INFO":        3,
		"SEVERITY_WARNING":     4,
		"SEVERITY_ERROR":       5,
		"SEVERITY_CRITICAL":    6,
	}
)

func (x Severity) Enum() *Severity {
	p := new(Severity)
	*p = x
	return p
}

func (x Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_datausage_v2_data_usage_proto_enumTypes[3].Descriptor()
}

func (Severity) Type() protoreflect.EnumType {
	return &file_com_coralogix_datausage_v2_data_usage_proto_enumTypes[3]
}

func (x Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Severity.Descriptor instead.
func (Severity) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_datausage_v2_data_usage_proto_rawDescGZIP(), []int{3}
}

type Team struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Id            *wrapperspb.UInt64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Team) Reset() {
	*x = Team{}
	mi := &file_com_coralogix_datausage_v2_data_usage_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Team) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Team) ProtoMessage() {}

func (x *Team) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_datausage_v2_data_usage_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Team.ProtoReflect.Descriptor instead.
func (*Team) Descriptor() ([]byte, []int) {
	return file_com_coralogix_datausage_v2_data_usage_proto_rawDescGZIP(), []int{0}
}

func (x *Team) GetId() *wrapperspb.UInt64Value {
	if x != nil {
		return x.Id
	}
	return nil
}

type DateRange struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FromDate      *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=from_date,json=fromDate,proto3" json:"from_date,omitempty"`
	ToDate        *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=to_date,json=toDate,proto3" json:"to_date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DateRange) Reset() {
	*x = DateRange{}
	mi := &file_com_coralogix_datausage_v2_data_usage_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DateRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DateRange) ProtoMessage() {}

func (x *DateRange) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_datausage_v2_data_usage_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DateRange.ProtoReflect.Descriptor instead.
func (*DateRange) Descriptor() ([]byte, []int) {
	return file_com_coralogix_datausage_v2_data_usage_proto_rawDescGZIP(), []int{1}
}

func (x *DateRange) GetFromDate() *timestamppb.Timestamp {
	if x != nil {
		return x.FromDate
	}
	return nil
}

func (x *DateRange) GetToDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ToDate
	}
	return nil
}

type GenericDimension struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenericDimension) Reset() {
	*x = GenericDimension{}
	mi := &file_com_coralogix_datausage_v2_data_usage_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenericDimension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericDimension) ProtoMessage() {}

func (x *GenericDimension) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_datausage_v2_data_usage_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericDimension.ProtoReflect.Descriptor instead.
func (*GenericDimension) Descriptor() ([]byte, []int) {
	return file_com_coralogix_datausage_v2_data_usage_proto_rawDescGZIP(), []int{2}
}

func (x *GenericDimension) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *GenericDimension) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Dimension struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Dimension:
	//
	//	*Dimension_Pillar
	//	*Dimension_GenericDimension
	//	*Dimension_Tier
	//	*Dimension_Severity
	//	*Dimension_Priority
	Dimension     isDimension_Dimension `protobuf_oneof:"dimension"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Dimension) Reset() {
	*x = Dimension{}
	mi := &file_com_coralogix_datausage_v2_data_usage_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Dimension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dimension) ProtoMessage() {}

func (x *Dimension) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_datausage_v2_data_usage_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dimension.ProtoReflect.Descriptor instead.
func (*Dimension) Descriptor() ([]byte, []int) {
	return file_com_coralogix_datausage_v2_data_usage_proto_rawDescGZIP(), []int{3}
}

func (x *Dimension) GetDimension() isDimension_Dimension {
	if x != nil {
		return x.Dimension
	}
	return nil
}

func (x *Dimension) GetPillar() Pillar {
	if x != nil {
		if x, ok := x.Dimension.(*Dimension_Pillar); ok {
			return x.Pillar
		}
	}
	return Pillar_PILLAR_UNSPECIFIED
}

func (x *Dimension) GetGenericDimension() *GenericDimension {
	if x != nil {
		if x, ok := x.Dimension.(*Dimension_GenericDimension); ok {
			return x.GenericDimension
		}
	}
	return nil
}

// Deprecated: Marked as deprecated in com/coralogix/datausage/v2/data_usage.proto.
func (x *Dimension) GetTier() TcoTier {
	if x != nil {
		if x, ok := x.Dimension.(*Dimension_Tier); ok {
			return x.Tier
		}
	}
	return TcoTier_TCO_TIER_UNSPECIFIED
}

func (x *Dimension) GetSeverity() Severity {
	if x != nil {
		if x, ok := x.Dimension.(*Dimension_Severity); ok {
			return x.Severity
		}
	}
	return Severity_SEVERITY_UNSPECIFIED
}

func (x *Dimension) GetPriority() Priority {
	if x != nil {
		if x, ok := x.Dimension.(*Dimension_Priority); ok {
			return x.Priority
		}
	}
	return Priority_PRIORITY_UNSPECIFIED
}

type isDimension_Dimension interface {
	isDimension_Dimension()
}

type Dimension_Pillar struct {
	Pillar Pillar `protobuf:"varint,1,opt,name=pillar,proto3,enum=com.coralogix.datausage.v2.Pillar,oneof"`
}

type Dimension_GenericDimension struct {
	GenericDimension *GenericDimension `protobuf:"bytes,2,opt,name=generic_dimension,json=genericDimension,proto3,oneof"`
}

type Dimension_Tier struct {
	// Deprecated: Marked as deprecated in com/coralogix/datausage/v2/data_usage.proto.
	Tier TcoTier `protobuf:"varint,3,opt,name=tier,proto3,enum=com.coralogix.datausage.v2.TcoTier,oneof"`
}

type Dimension_Severity struct {
	Severity Severity `protobuf:"varint,4,opt,name=severity,proto3,enum=com.coralogix.datausage.v2.Severity,oneof"`
}

type Dimension_Priority struct {
	Priority Priority `protobuf:"varint,5,opt,name=priority,proto3,enum=com.coralogix.datausage.v2.Priority,oneof"`
}

func (*Dimension_Pillar) isDimension_Dimension() {}

func (*Dimension_GenericDimension) isDimension_Dimension() {}

func (*Dimension_Tier) isDimension_Dimension() {}

func (*Dimension_Severity) isDimension_Dimension() {}

func (*Dimension_Priority) isDimension_Dimension() {}

var File_com_coralogix_datausage_v2_data_usage_proto protoreflect.FileDescriptor

var file_com_coralogix_datausage_v2_data_usage_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x04, 0x54, 0x65,
	0x61, 0x6d, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x69, 0x64,
	0x22, 0xf4, 0x02, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x54,
	0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x1b, 0x92,
	0x41, 0x18, 0x4a, 0x16, 0x22, 0x32, 0x30, 0x32, 0x31, 0x2d, 0x30, 0x31, 0x2d, 0x30, 0x31, 0x54,
	0x30, 0x30, 0x3a, 0x30, 0x30, 0x3a, 0x30, 0x30, 0x5a, 0x22, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x50, 0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x42, 0x1b, 0x92, 0x41, 0x18, 0x4a, 0x16, 0x22, 0x32, 0x30, 0x32, 0x31, 0x2d, 0x30, 0x31,
	0x2d, 0x30, 0x31, 0x54, 0x30, 0x30, 0x3a, 0x30, 0x30, 0x3a, 0x30, 0x30, 0x5a, 0x22, 0x52, 0x06,
	0x74, 0x6f, 0x44, 0x61, 0x74, 0x65, 0x3a, 0xbe, 0x01, 0x92, 0x41, 0xba, 0x01, 0x0a, 0x3a, 0x2a,
	0x0a, 0x44, 0x61, 0x74, 0x65, 0x20, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x32, 0x2c, 0x54, 0x68, 0x69,
	0x73, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x20, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x73, 0x20, 0x61, 0x20, 0x64, 0x61,
	0x74, 0x65, 0x20, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x2a, 0x7c, 0x0a, 0x1f, 0x46, 0x69, 0x6e,
	0x64, 0x20, 0x6f, 0x75, 0x74, 0x20, 0x6d, 0x6f, 0x72, 0x65, 0x20, 0x61, 0x62, 0x6f, 0x75, 0x74,
	0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x12, 0x59, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x63, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x67, 0x75,
	0x69, 0x64, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2d, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2d,
	0x61, 0x6e, 0x64, 0x2d, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x2d, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x22, 0xa3, 0x02, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x69, 0x63, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x92, 0x41, 0x07, 0x4a, 0x05,
	0x22, 0x6b, 0x65, 0x79, 0x22, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0x92, 0x41, 0x09, 0x4a, 0x07,
	0x22, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0xcc,
	0x01, 0x92, 0x41, 0xc8, 0x01, 0x0a, 0x48, 0x2a, 0x11, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x20, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0x33, 0x54, 0x68, 0x69, 0x73,
	0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x20,
	0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x73, 0x20, 0x61, 0x20, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x69, 0x63, 0x20, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x2a,
	0x7c, 0x0a, 0x1f, 0x46, 0x69, 0x6e, 0x64, 0x20, 0x6f, 0x75, 0x74, 0x20, 0x6d, 0x6f, 0x72, 0x65,
	0x20, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x75, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x12, 0x59, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x63, 0x73, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2d, 0x67, 0x75, 0x69, 0x64, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x61, 0x6e, 0x64, 0x2d, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x22, 0xfa, 0x02,
	0x0a, 0x09, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x06, 0x70,
	0x69, 0x6c, 0x6c, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x69, 0x6c, 0x6c, 0x61, 0x72, 0x48,
	0x00, 0x52, 0x06, 0x70, 0x69, 0x6c, 0x6c, 0x61, 0x72, 0x12, 0x5b, 0x0a, 0x11, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x69, 0x63, 0x5f, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x32, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x48, 0x00, 0x52, 0x10, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x44, 0x69, 0x6d,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x04, 0x74, 0x69, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x32, 0x2e, 0x54, 0x63, 0x6f, 0x54, 0x69, 0x65, 0x72, 0x42, 0x02, 0x18, 0x01, 0x48, 0x00, 0x52,
	0x04, 0x74, 0x69, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x75, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x48, 0x00, 0x52,
	0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x42, 0x0a, 0x08, 0x70, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x48, 0x00, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x42, 0x0b, 0x0a,
	0x09, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2a, 0x6a, 0x0a, 0x06, 0x50, 0x69,
	0x6c, 0x6c, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x49, 0x4c, 0x4c, 0x41, 0x52, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e,
	0x50, 0x49, 0x4c, 0x4c, 0x41, 0x52, 0x5f, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x53, 0x10, 0x01,
	0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x49, 0x4c, 0x4c, 0x41, 0x52, 0x5f, 0x4c, 0x4f, 0x47, 0x53, 0x10,
	0x02, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x49, 0x4c, 0x4c, 0x41, 0x52, 0x5f, 0x53, 0x50, 0x41, 0x4e,
	0x53, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x49, 0x4c, 0x4c, 0x41, 0x52, 0x5f, 0x42, 0x49,
	0x4e, 0x41, 0x52, 0x59, 0x10, 0x04, 0x2a, 0x77, 0x0a, 0x07, 0x54, 0x63, 0x6f, 0x54, 0x69, 0x65,
	0x72, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x43, 0x4f, 0x5f, 0x54, 0x49, 0x45, 0x52, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x54,
	0x43, 0x4f, 0x5f, 0x54, 0x49, 0x45, 0x52, 0x5f, 0x4c, 0x4f, 0x57, 0x10, 0x01, 0x12, 0x13, 0x0a,
	0x0f, 0x54, 0x43, 0x4f, 0x5f, 0x54, 0x49, 0x45, 0x52, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x55, 0x4d,
	0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x54, 0x43, 0x4f, 0x5f, 0x54, 0x49, 0x45, 0x52, 0x5f, 0x48,
	0x49, 0x47, 0x48, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x43, 0x4f, 0x5f, 0x54, 0x49, 0x45,
	0x52, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x04, 0x1a, 0x02, 0x18, 0x01, 0x2a,
	0x74, 0x0a, 0x08, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x14, 0x50,
	0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54,
	0x59, 0x5f, 0x4c, 0x4f, 0x57, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x52, 0x49, 0x4f, 0x52,
	0x49, 0x54, 0x59, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x55, 0x4d, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d,
	0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x10, 0x03, 0x12,
	0x14, 0x0a, 0x10, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x42, 0x4c, 0x4f, 0x43,
	0x4b, 0x45, 0x44, 0x10, 0x04, 0x2a, 0xa2, 0x01, 0x0a, 0x08, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e,
	0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x01,
	0x12, 0x14, 0x0a, 0x10, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x56, 0x45, 0x52,
	0x42, 0x4f, 0x53, 0x45, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49,
	0x54, 0x59, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x45, 0x56,
	0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12,
	0x12, 0x0a, 0x0e, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x05, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f,
	0x43, 0x52, 0x49, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x06, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_com_coralogix_datausage_v2_data_usage_proto_rawDescOnce sync.Once
	file_com_coralogix_datausage_v2_data_usage_proto_rawDescData = file_com_coralogix_datausage_v2_data_usage_proto_rawDesc
)

func file_com_coralogix_datausage_v2_data_usage_proto_rawDescGZIP() []byte {
	file_com_coralogix_datausage_v2_data_usage_proto_rawDescOnce.Do(func() {
		file_com_coralogix_datausage_v2_data_usage_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_datausage_v2_data_usage_proto_rawDescData)
	})
	return file_com_coralogix_datausage_v2_data_usage_proto_rawDescData
}

var file_com_coralogix_datausage_v2_data_usage_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_com_coralogix_datausage_v2_data_usage_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_com_coralogix_datausage_v2_data_usage_proto_goTypes = []any{
	(Pillar)(0),                    // 0: com.coralogix.datausage.v2.Pillar
	(TcoTier)(0),                   // 1: com.coralogix.datausage.v2.TcoTier
	(Priority)(0),                  // 2: com.coralogix.datausage.v2.Priority
	(Severity)(0),                  // 3: com.coralogix.datausage.v2.Severity
	(*Team)(nil),                   // 4: com.coralogix.datausage.v2.Team
	(*DateRange)(nil),              // 5: com.coralogix.datausage.v2.DateRange
	(*GenericDimension)(nil),       // 6: com.coralogix.datausage.v2.GenericDimension
	(*Dimension)(nil),              // 7: com.coralogix.datausage.v2.Dimension
	(*wrapperspb.UInt64Value)(nil), // 8: google.protobuf.UInt64Value
	(*timestamppb.Timestamp)(nil),  // 9: google.protobuf.Timestamp
}
var file_com_coralogix_datausage_v2_data_usage_proto_depIdxs = []int32{
	8, // 0: com.coralogix.datausage.v2.Team.id:type_name -> google.protobuf.UInt64Value
	9, // 1: com.coralogix.datausage.v2.DateRange.from_date:type_name -> google.protobuf.Timestamp
	9, // 2: com.coralogix.datausage.v2.DateRange.to_date:type_name -> google.protobuf.Timestamp
	0, // 3: com.coralogix.datausage.v2.Dimension.pillar:type_name -> com.coralogix.datausage.v2.Pillar
	6, // 4: com.coralogix.datausage.v2.Dimension.generic_dimension:type_name -> com.coralogix.datausage.v2.GenericDimension
	1, // 5: com.coralogix.datausage.v2.Dimension.tier:type_name -> com.coralogix.datausage.v2.TcoTier
	3, // 6: com.coralogix.datausage.v2.Dimension.severity:type_name -> com.coralogix.datausage.v2.Severity
	2, // 7: com.coralogix.datausage.v2.Dimension.priority:type_name -> com.coralogix.datausage.v2.Priority
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_com_coralogix_datausage_v2_data_usage_proto_init() }
func file_com_coralogix_datausage_v2_data_usage_proto_init() {
	if File_com_coralogix_datausage_v2_data_usage_proto != nil {
		return
	}
	file_com_coralogix_datausage_v2_data_usage_proto_msgTypes[3].OneofWrappers = []any{
		(*Dimension_Pillar)(nil),
		(*Dimension_GenericDimension)(nil),
		(*Dimension_Tier)(nil),
		(*Dimension_Severity)(nil),
		(*Dimension_Priority)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_datausage_v2_data_usage_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_datausage_v2_data_usage_proto_goTypes,
		DependencyIndexes: file_com_coralogix_datausage_v2_data_usage_proto_depIdxs,
		EnumInfos:         file_com_coralogix_datausage_v2_data_usage_proto_enumTypes,
		MessageInfos:      file_com_coralogix_datausage_v2_data_usage_proto_msgTypes,
	}.Build()
	File_com_coralogix_datausage_v2_data_usage_proto = out.File
	file_com_coralogix_datausage_v2_data_usage_proto_rawDesc = nil
	file_com_coralogix_datausage_v2_data_usage_proto_goTypes = nil
	file_com_coralogix_datausage_v2_data_usage_proto_depIdxs = nil
}
