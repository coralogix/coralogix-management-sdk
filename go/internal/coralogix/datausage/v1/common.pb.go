// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.3
// source: com/coralogix/datausage/v1/common.proto

package v1

import (
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

// Deprecated: Marked as deprecated in com/coralogix/datausage/v1/common.proto.
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
	return file_com_coralogix_datausage_v1_common_proto_enumTypes[0].Descriptor()
}

func (TcoTier) Type() protoreflect.EnumType {
	return &file_com_coralogix_datausage_v1_common_proto_enumTypes[0]
}

func (x TcoTier) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TcoTier.Descriptor instead.
func (TcoTier) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_datausage_v1_common_proto_rawDescGZIP(), []int{0}
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
	return file_com_coralogix_datausage_v1_common_proto_enumTypes[1].Descriptor()
}

func (Priority) Type() protoreflect.EnumType {
	return &file_com_coralogix_datausage_v1_common_proto_enumTypes[1]
}

func (x Priority) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Priority.Descriptor instead.
func (Priority) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_datausage_v1_common_proto_rawDescGZIP(), []int{1}
}

type DataType int32

const (
	DataType_DATA_TYPE_UNSPECIFIED DataType = 0
	DataType_DATA_TYPE_LOGS        DataType = 1
	DataType_DATA_TYPE_METRICS     DataType = 2
	DataType_DATA_TYPE_TRACING     DataType = 3
)

// Enum value maps for DataType.
var (
	DataType_name = map[int32]string{
		0: "DATA_TYPE_UNSPECIFIED",
		1: "DATA_TYPE_LOGS",
		2: "DATA_TYPE_METRICS",
		3: "DATA_TYPE_TRACING",
	}
	DataType_value = map[string]int32{
		"DATA_TYPE_UNSPECIFIED": 0,
		"DATA_TYPE_LOGS":        1,
		"DATA_TYPE_METRICS":     2,
		"DATA_TYPE_TRACING":     3,
	}
)

func (x DataType) Enum() *DataType {
	p := new(DataType)
	*p = x
	return p
}

func (x DataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_datausage_v1_common_proto_enumTypes[2].Descriptor()
}

func (DataType) Type() protoreflect.EnumType {
	return &file_com_coralogix_datausage_v1_common_proto_enumTypes[2]
}

func (x DataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataType.Descriptor instead.
func (DataType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_datausage_v1_common_proto_rawDescGZIP(), []int{2}
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
	return file_com_coralogix_datausage_v1_common_proto_enumTypes[3].Descriptor()
}

func (Severity) Type() protoreflect.EnumType {
	return &file_com_coralogix_datausage_v1_common_proto_enumTypes[3]
}

func (x Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Severity.Descriptor instead.
func (Severity) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_datausage_v1_common_proto_rawDescGZIP(), []int{3}
}

type Range int32

const (
	Range_RANGE_UNSPECIFIED   Range = 0
	Range_RANGE_CURRENT_MONTH Range = 1
	Range_RANGE_LAST_30_DAYS  Range = 2
	Range_RANGE_LAST_90_DAYS  Range = 3
	Range_RANGE_LAST_WEEK     Range = 4
)

// Enum value maps for Range.
var (
	Range_name = map[int32]string{
		0: "RANGE_UNSPECIFIED",
		1: "RANGE_CURRENT_MONTH",
		2: "RANGE_LAST_30_DAYS",
		3: "RANGE_LAST_90_DAYS",
		4: "RANGE_LAST_WEEK",
	}
	Range_value = map[string]int32{
		"RANGE_UNSPECIFIED":   0,
		"RANGE_CURRENT_MONTH": 1,
		"RANGE_LAST_30_DAYS":  2,
		"RANGE_LAST_90_DAYS":  3,
		"RANGE_LAST_WEEK":     4,
	}
)

func (x Range) Enum() *Range {
	p := new(Range)
	*p = x
	return p
}

func (x Range) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Range) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_datausage_v1_common_proto_enumTypes[4].Descriptor()
}

func (Range) Type() protoreflect.EnumType {
	return &file_com_coralogix_datausage_v1_common_proto_enumTypes[4]
}

func (x Range) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Range.Descriptor instead.
func (Range) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_datausage_v1_common_proto_rawDescGZIP(), []int{4}
}

type Resolution int32

const (
	Resolution_RESOLUTION_UNSPECIFIED     Resolution = 0
	Resolution_RESOLUTION_ONE_MINUTE      Resolution = 1
	Resolution_RESOLUTION_FIVE_MINUTES    Resolution = 2
	Resolution_RESOLUTION_FIFTEEN_MINUTES Resolution = 3
	Resolution_RESOLUTION_ONE_HOUR        Resolution = 4
	Resolution_RESOLUTION_SIX_HOURS       Resolution = 5
	Resolution_RESOLUTION_ONE_DAY         Resolution = 6
	Resolution_RESOLUTION_ONE_WEEK        Resolution = 7
)

// Enum value maps for Resolution.
var (
	Resolution_name = map[int32]string{
		0: "RESOLUTION_UNSPECIFIED",
		1: "RESOLUTION_ONE_MINUTE",
		2: "RESOLUTION_FIVE_MINUTES",
		3: "RESOLUTION_FIFTEEN_MINUTES",
		4: "RESOLUTION_ONE_HOUR",
		5: "RESOLUTION_SIX_HOURS",
		6: "RESOLUTION_ONE_DAY",
		7: "RESOLUTION_ONE_WEEK",
	}
	Resolution_value = map[string]int32{
		"RESOLUTION_UNSPECIFIED":     0,
		"RESOLUTION_ONE_MINUTE":      1,
		"RESOLUTION_FIVE_MINUTES":    2,
		"RESOLUTION_FIFTEEN_MINUTES": 3,
		"RESOLUTION_ONE_HOUR":        4,
		"RESOLUTION_SIX_HOURS":       5,
		"RESOLUTION_ONE_DAY":         6,
		"RESOLUTION_ONE_WEEK":        7,
	}
)

func (x Resolution) Enum() *Resolution {
	p := new(Resolution)
	*p = x
	return p
}

func (x Resolution) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Resolution) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_datausage_v1_common_proto_enumTypes[5].Descriptor()
}

func (Resolution) Type() protoreflect.EnumType {
	return &file_com_coralogix_datausage_v1_common_proto_enumTypes[5]
}

func (x Resolution) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Resolution.Descriptor instead.
func (Resolution) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_datausage_v1_common_proto_rawDescGZIP(), []int{5}
}

type Team struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *wrapperspb.UInt64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Team) Reset() {
	*x = Team{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_datausage_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Team) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Team) ProtoMessage() {}

func (x *Team) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_datausage_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_com_coralogix_datausage_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *Team) GetId() *wrapperspb.UInt64Value {
	if x != nil {
		return x.Id
	}
	return nil
}

type Unit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *wrapperspb.FloatValue `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Unit) Reset() {
	*x = Unit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_datausage_v1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unit) ProtoMessage() {}

func (x *Unit) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_datausage_v1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unit.ProtoReflect.Descriptor instead.
func (*Unit) Descriptor() ([]byte, []int) {
	return file_com_coralogix_datausage_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *Unit) GetValue() *wrapperspb.FloatValue {
	if x != nil {
		return x.Value
	}
	return nil
}

type GB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *wrapperspb.FloatValue `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GB) Reset() {
	*x = GB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_datausage_v1_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GB) ProtoMessage() {}

func (x *GB) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_datausage_v1_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GB.ProtoReflect.Descriptor instead.
func (*GB) Descriptor() ([]byte, []int) {
	return file_com_coralogix_datausage_v1_common_proto_rawDescGZIP(), []int{2}
}

func (x *GB) GetValue() *wrapperspb.FloatValue {
	if x != nil {
		return x.Value
	}
	return nil
}

type Retention struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *wrapperspb.UInt64Value `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Retention) Reset() {
	*x = Retention{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_datausage_v1_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Retention) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Retention) ProtoMessage() {}

func (x *Retention) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_datausage_v1_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Retention.ProtoReflect.Descriptor instead.
func (*Retention) Descriptor() ([]byte, []int) {
	return file_com_coralogix_datausage_v1_common_proto_rawDescGZIP(), []int{3}
}

func (x *Retention) GetValue() *wrapperspb.UInt64Value {
	if x != nil {
		return x.Value
	}
	return nil
}

type DateRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromDate *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=from_date,json=fromDate,proto3" json:"from_date,omitempty"`
	ToDate   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=to_date,json=toDate,proto3" json:"to_date,omitempty"`
}

func (x *DateRange) Reset() {
	*x = DateRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_datausage_v1_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DateRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DateRange) ProtoMessage() {}

func (x *DateRange) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_datausage_v1_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_com_coralogix_datausage_v1_common_proto_rawDescGZIP(), []int{4}
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

var File_com_coralogix_datausage_v1_common_proto protoreflect.FileDescriptor

var file_com_coralogix_datausage_v1_common_proto_rawDesc = []byte{
	0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x75, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x04, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x2c,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x69, 0x64, 0x22, 0x39, 0x0a, 0x04,
	0x55, 0x6e, 0x69, 0x74, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x37, 0x0a, 0x02, 0x47, 0x42, 0x12, 0x31, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x3f, 0x0a, 0x09, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x79, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x37,
	0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x66,
	0x72, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x74, 0x6f, 0x44, 0x61, 0x74, 0x65, 0x2a, 0x77, 0x0a, 0x07,
	0x54, 0x63, 0x6f, 0x54, 0x69, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x43, 0x4f, 0x5f, 0x54,
	0x49, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x43, 0x4f, 0x5f, 0x54, 0x49, 0x45, 0x52, 0x5f, 0x4c, 0x4f,
	0x57, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x43, 0x4f, 0x5f, 0x54, 0x49, 0x45, 0x52, 0x5f,
	0x4d, 0x45, 0x44, 0x49, 0x55, 0x4d, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x54, 0x43, 0x4f, 0x5f,
	0x54, 0x49, 0x45, 0x52, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x54,
	0x43, 0x4f, 0x5f, 0x54, 0x49, 0x45, 0x52, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10,
	0x04, 0x1a, 0x02, 0x18, 0x01, 0x2a, 0x74, 0x0a, 0x08, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x50,
	0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4c, 0x4f, 0x57, 0x10, 0x01, 0x12, 0x13, 0x0a,
	0x0f, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x55, 0x4d,
	0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x48,
	0x49, 0x47, 0x48, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54,
	0x59, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x04, 0x2a, 0x67, 0x0a, 0x08, 0x44,
	0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x41, 0x54, 0x41, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4c, 0x4f, 0x47, 0x53, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x53, 0x10, 0x02, 0x12, 0x15, 0x0a,
	0x11, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x49,
	0x4e, 0x47, 0x10, 0x03, 0x2a, 0xa2, 0x01, 0x0a, 0x08, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53,
	0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x01, 0x12,
	0x14, 0x0a, 0x10, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x56, 0x45, 0x52, 0x42,
	0x4f, 0x53, 0x45, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54,
	0x59, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x45, 0x56, 0x45,
	0x52, 0x49, 0x54, 0x59, 0x5f, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x12,
	0x0a, 0x0e, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x05, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x43,
	0x52, 0x49, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x06, 0x2a, 0x7c, 0x0a, 0x05, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x41, 0x4e,
	0x47, 0x45, 0x5f, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x4f, 0x4e, 0x54, 0x48,
	0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x4c, 0x41, 0x53, 0x54,
	0x5f, 0x33, 0x30, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x41,
	0x4e, 0x47, 0x45, 0x5f, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x39, 0x30, 0x5f, 0x44, 0x41, 0x59, 0x53,
	0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x4c, 0x41, 0x53, 0x54,
	0x5f, 0x57, 0x45, 0x45, 0x4b, 0x10, 0x04, 0x2a, 0xe4, 0x01, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x55,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x55, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x4f, 0x4e, 0x45, 0x5f, 0x4d, 0x49, 0x4e, 0x55, 0x54, 0x45, 0x10, 0x01, 0x12, 0x1b, 0x0a,
	0x17, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x49, 0x56, 0x45,
	0x5f, 0x4d, 0x49, 0x4e, 0x55, 0x54, 0x45, 0x53, 0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x52, 0x45,
	0x53, 0x4f, 0x4c, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x49, 0x46, 0x54, 0x45, 0x45, 0x4e,
	0x5f, 0x4d, 0x49, 0x4e, 0x55, 0x54, 0x45, 0x53, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x45,
	0x53, 0x4f, 0x4c, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x4e, 0x45, 0x5f, 0x48, 0x4f, 0x55,
	0x52, 0x10, 0x04, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x55, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x53, 0x49, 0x58, 0x5f, 0x48, 0x4f, 0x55, 0x52, 0x53, 0x10, 0x05, 0x12, 0x16, 0x0a,
	0x12, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x4e, 0x45, 0x5f,
	0x44, 0x41, 0x59, 0x10, 0x06, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x55, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x4e, 0x45, 0x5f, 0x57, 0x45, 0x45, 0x4b, 0x10, 0x07, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_datausage_v1_common_proto_rawDescOnce sync.Once
	file_com_coralogix_datausage_v1_common_proto_rawDescData = file_com_coralogix_datausage_v1_common_proto_rawDesc
)

func file_com_coralogix_datausage_v1_common_proto_rawDescGZIP() []byte {
	file_com_coralogix_datausage_v1_common_proto_rawDescOnce.Do(func() {
		file_com_coralogix_datausage_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_datausage_v1_common_proto_rawDescData)
	})
	return file_com_coralogix_datausage_v1_common_proto_rawDescData
}

var file_com_coralogix_datausage_v1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 6)
var file_com_coralogix_datausage_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_com_coralogix_datausage_v1_common_proto_goTypes = []interface{}{
	(TcoTier)(0),                   // 0: com.coralogix.datausage.v1.TcoTier
	(Priority)(0),                  // 1: com.coralogix.datausage.v1.Priority
	(DataType)(0),                  // 2: com.coralogix.datausage.v1.DataType
	(Severity)(0),                  // 3: com.coralogix.datausage.v1.Severity
	(Range)(0),                     // 4: com.coralogix.datausage.v1.Range
	(Resolution)(0),                // 5: com.coralogix.datausage.v1.Resolution
	(*Team)(nil),                   // 6: com.coralogix.datausage.v1.Team
	(*Unit)(nil),                   // 7: com.coralogix.datausage.v1.Unit
	(*GB)(nil),                     // 8: com.coralogix.datausage.v1.GB
	(*Retention)(nil),              // 9: com.coralogix.datausage.v1.Retention
	(*DateRange)(nil),              // 10: com.coralogix.datausage.v1.DateRange
	(*wrapperspb.UInt64Value)(nil), // 11: google.protobuf.UInt64Value
	(*wrapperspb.FloatValue)(nil),  // 12: google.protobuf.FloatValue
	(*timestamppb.Timestamp)(nil),  // 13: google.protobuf.Timestamp
}
var file_com_coralogix_datausage_v1_common_proto_depIdxs = []int32{
	11, // 0: com.coralogix.datausage.v1.Team.id:type_name -> google.protobuf.UInt64Value
	12, // 1: com.coralogix.datausage.v1.Unit.value:type_name -> google.protobuf.FloatValue
	12, // 2: com.coralogix.datausage.v1.GB.value:type_name -> google.protobuf.FloatValue
	11, // 3: com.coralogix.datausage.v1.Retention.value:type_name -> google.protobuf.UInt64Value
	13, // 4: com.coralogix.datausage.v1.DateRange.from_date:type_name -> google.protobuf.Timestamp
	13, // 5: com.coralogix.datausage.v1.DateRange.to_date:type_name -> google.protobuf.Timestamp
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_com_coralogix_datausage_v1_common_proto_init() }
func file_com_coralogix_datausage_v1_common_proto_init() {
	if File_com_coralogix_datausage_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_datausage_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Team); i {
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
		file_com_coralogix_datausage_v1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unit); i {
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
		file_com_coralogix_datausage_v1_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GB); i {
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
		file_com_coralogix_datausage_v1_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Retention); i {
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
		file_com_coralogix_datausage_v1_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DateRange); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_datausage_v1_common_proto_rawDesc,
			NumEnums:      6,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_datausage_v1_common_proto_goTypes,
		DependencyIndexes: file_com_coralogix_datausage_v1_common_proto_depIdxs,
		EnumInfos:         file_com_coralogix_datausage_v1_common_proto_enumTypes,
		MessageInfos:      file_com_coralogix_datausage_v1_common_proto_msgTypes,
	}.Build()
	File_com_coralogix_datausage_v1_common_proto = out.File
	file_com_coralogix_datausage_v1_common_proto_rawDesc = nil
	file_com_coralogix_datausage_v1_common_proto_goTypes = nil
	file_com_coralogix_datausage_v1_common_proto_depIdxs = nil
}
