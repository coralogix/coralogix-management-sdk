// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.1
// source: com/coralogix/dataprime/v1/internal_metadata.proto

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

type MetadataEnums_Tier int32

const (
	MetadataEnums_TIER_UNSPECIFIED     MetadataEnums_Tier = 0
	MetadataEnums_TIER_ARCHIVE         MetadataEnums_Tier = 1
	MetadataEnums_TIER_FREQUENT_SEARCH MetadataEnums_Tier = 2
)

// Enum value maps for MetadataEnums_Tier.
var (
	MetadataEnums_Tier_name = map[int32]string{
		0: "TIER_UNSPECIFIED",
		1: "TIER_ARCHIVE",
		2: "TIER_FREQUENT_SEARCH",
	}
	MetadataEnums_Tier_value = map[string]int32{
		"TIER_UNSPECIFIED":     0,
		"TIER_ARCHIVE":         1,
		"TIER_FREQUENT_SEARCH": 2,
	}
)

func (x MetadataEnums_Tier) Enum() *MetadataEnums_Tier {
	p := new(MetadataEnums_Tier)
	*p = x
	return p
}

func (x MetadataEnums_Tier) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetadataEnums_Tier) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_dataprime_v1_internal_metadata_proto_enumTypes[0].Descriptor()
}

func (MetadataEnums_Tier) Type() protoreflect.EnumType {
	return &file_com_coralogix_dataprime_v1_internal_metadata_proto_enumTypes[0]
}

func (x MetadataEnums_Tier) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetadataEnums_Tier.Descriptor instead.
func (MetadataEnums_Tier) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_internal_metadata_proto_rawDescGZIP(), []int{0, 0}
}

type MetadataEnums_QuerySyntax int32

const (
	MetadataEnums_QUERY_SYNTAX_UNSPECIFIED MetadataEnums_QuerySyntax = 0
	MetadataEnums_QUERY_SYNTAX_LUCENE      MetadataEnums_QuerySyntax = 1
	MetadataEnums_QUERY_SYNTAX_DATAPRIME   MetadataEnums_QuerySyntax = 2
)

// Enum value maps for MetadataEnums_QuerySyntax.
var (
	MetadataEnums_QuerySyntax_name = map[int32]string{
		0: "QUERY_SYNTAX_UNSPECIFIED",
		1: "QUERY_SYNTAX_LUCENE",
		2: "QUERY_SYNTAX_DATAPRIME",
	}
	MetadataEnums_QuerySyntax_value = map[string]int32{
		"QUERY_SYNTAX_UNSPECIFIED": 0,
		"QUERY_SYNTAX_LUCENE":      1,
		"QUERY_SYNTAX_DATAPRIME":   2,
	}
)

func (x MetadataEnums_QuerySyntax) Enum() *MetadataEnums_QuerySyntax {
	p := new(MetadataEnums_QuerySyntax)
	*p = x
	return p
}

func (x MetadataEnums_QuerySyntax) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetadataEnums_QuerySyntax) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_dataprime_v1_internal_metadata_proto_enumTypes[1].Descriptor()
}

func (MetadataEnums_QuerySyntax) Type() protoreflect.EnumType {
	return &file_com_coralogix_dataprime_v1_internal_metadata_proto_enumTypes[1]
}

func (x MetadataEnums_QuerySyntax) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetadataEnums_QuerySyntax.Descriptor instead.
func (MetadataEnums_QuerySyntax) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_internal_metadata_proto_rawDescGZIP(), []int{0, 1}
}

type MetadataEnums struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MetadataEnums) Reset() {
	*x = MetadataEnums{}
	mi := &file_com_coralogix_dataprime_v1_internal_metadata_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetadataEnums) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataEnums) ProtoMessage() {}

func (x *MetadataEnums) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_internal_metadata_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataEnums.ProtoReflect.Descriptor instead.
func (*MetadataEnums) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_internal_metadata_proto_rawDescGZIP(), []int{0}
}

type QueryMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// beginning of the time range for the query
	StartDate *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start_date,json=startDate,proto3,oneof" json:"start_date,omitempty"`
	// end of the time range for the query
	EndDate *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end_date,json=endDate,proto3,oneof" json:"end_date,omitempty"`
	// tier on which query runs
	Tier *MetadataEnums_Tier `protobuf:"varint,3,opt,name=tier,proto3,enum=com.coralogix.dataprime.v1.MetadataEnums_Tier,oneof" json:"tier,omitempty"`
	// limit number of results
	Limit *int32 `protobuf:"varint,4,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	// dataprime strict fields validation in parser
	StrictFieldsValidation *bool `protobuf:"varint,5,opt,name=strict_fields_validation,json=strictFieldsValidation,proto3,oneof" json:"strict_fields_validation,omitempty"`
	// contextual `now` for the query, default: current time
	NowDate *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=now_date,json=nowDate,proto3,oneof" json:"now_date,omitempty"`
}

func (x *QueryMetadata) Reset() {
	*x = QueryMetadata{}
	mi := &file_com_coralogix_dataprime_v1_internal_metadata_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMetadata) ProtoMessage() {}

func (x *QueryMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_internal_metadata_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMetadata.ProtoReflect.Descriptor instead.
func (*QueryMetadata) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_internal_metadata_proto_rawDescGZIP(), []int{1}
}

func (x *QueryMetadata) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *QueryMetadata) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

func (x *QueryMetadata) GetTier() MetadataEnums_Tier {
	if x != nil && x.Tier != nil {
		return *x.Tier
	}
	return MetadataEnums_TIER_UNSPECIFIED
}

func (x *QueryMetadata) GetLimit() int32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *QueryMetadata) GetStrictFieldsValidation() bool {
	if x != nil && x.StrictFieldsValidation != nil {
		return *x.StrictFieldsValidation
	}
	return false
}

func (x *QueryMetadata) GetNowDate() *timestamppb.Timestamp {
	if x != nil {
		return x.NowDate
	}
	return nil
}

// only text related metadata
type TextQueryMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// text query syntax
	Syntax *MetadataEnums_QuerySyntax `protobuf:"varint,1,opt,name=syntax,proto3,enum=com.coralogix.dataprime.v1.MetadataEnums_QuerySyntax,oneof" json:"syntax,omitempty"`
	// default value for source to be used when source is omitted in a query
	DefaultSource *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=default_source,json=defaultSource,proto3,oneof" json:"default_source,omitempty"`
}

func (x *TextQueryMetadata) Reset() {
	*x = TextQueryMetadata{}
	mi := &file_com_coralogix_dataprime_v1_internal_metadata_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TextQueryMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextQueryMetadata) ProtoMessage() {}

func (x *TextQueryMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_internal_metadata_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextQueryMetadata.ProtoReflect.Descriptor instead.
func (*TextQueryMetadata) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_internal_metadata_proto_rawDescGZIP(), []int{2}
}

func (x *TextQueryMetadata) GetSyntax() MetadataEnums_QuerySyntax {
	if x != nil && x.Syntax != nil {
		return *x.Syntax
	}
	return MetadataEnums_QUERY_SYNTAX_UNSPECIFIED
}

func (x *TextQueryMetadata) GetDefaultSource() *wrapperspb.StringValue {
	if x != nil {
		return x.DefaultSource
	}
	return nil
}

var File_com_coralogix_dataprime_v1_internal_metadata_proto protoreflect.FileDescriptor

var file_com_coralogix_dataprime_v1_internal_metadata_proto_rawDesc = []byte{
	0x0a, 0x32, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xbb, 0x01, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0x22, 0x48, 0x0a, 0x04, 0x54, 0x69, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x10, 0x54,
	0x49, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x49, 0x45, 0x52, 0x5f, 0x41, 0x52, 0x43, 0x48, 0x49, 0x56,
	0x45, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x49, 0x45, 0x52, 0x5f, 0x46, 0x52, 0x45, 0x51,
	0x55, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x10, 0x02, 0x22, 0x60, 0x0a,
	0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x12, 0x1c, 0x0a, 0x18,
	0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x53, 0x59, 0x4e, 0x54, 0x41, 0x58, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x51, 0x55,
	0x45, 0x52, 0x59, 0x5f, 0x53, 0x59, 0x4e, 0x54, 0x41, 0x58, 0x5f, 0x4c, 0x55, 0x43, 0x45, 0x4e,
	0x45, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x53, 0x59, 0x4e,
	0x54, 0x41, 0x58, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x50, 0x52, 0x49, 0x4d, 0x45, 0x10, 0x02, 0x22,
	0xc9, 0x03, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x3e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x48, 0x00, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x3a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48,
	0x01, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x47, 0x0a,
	0x04, 0x74, 0x69, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x54, 0x69, 0x65, 0x72, 0x48, 0x02, 0x52, 0x04, 0x74,
	0x69, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x3d, 0x0a, 0x18, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x04, 0x52, 0x16, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01,
	0x12, 0x3a, 0x0a, 0x08, 0x6e, 0x6f, 0x77, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x05,
	0x52, 0x07, 0x6e, 0x6f, 0x77, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b,
	0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f,
	0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x69, 0x65,
	0x72, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x1b, 0x0a, 0x19, 0x5f,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x5f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6e, 0x6f, 0x77,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x04, 0x08, 0x06, 0x10, 0x07, 0x22, 0xcf, 0x01, 0x0a, 0x11,
	0x54, 0x65, 0x78, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x52, 0x0a, 0x06, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x53, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x48, 0x00, 0x52, 0x06, 0x73, 0x79, 0x6e, 0x74,
	0x61, 0x78, 0x88, 0x01, 0x01, 0x12, 0x48, 0x0a, 0x0e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x01, 0x52, 0x0d, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x88, 0x01, 0x01, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_dataprime_v1_internal_metadata_proto_rawDescOnce sync.Once
	file_com_coralogix_dataprime_v1_internal_metadata_proto_rawDescData = file_com_coralogix_dataprime_v1_internal_metadata_proto_rawDesc
)

func file_com_coralogix_dataprime_v1_internal_metadata_proto_rawDescGZIP() []byte {
	file_com_coralogix_dataprime_v1_internal_metadata_proto_rawDescOnce.Do(func() {
		file_com_coralogix_dataprime_v1_internal_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_dataprime_v1_internal_metadata_proto_rawDescData)
	})
	return file_com_coralogix_dataprime_v1_internal_metadata_proto_rawDescData
}

var file_com_coralogix_dataprime_v1_internal_metadata_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_com_coralogix_dataprime_v1_internal_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogix_dataprime_v1_internal_metadata_proto_goTypes = []any{
	(MetadataEnums_Tier)(0),        // 0: com.coralogix.dataprime.v1.MetadataEnums.Tier
	(MetadataEnums_QuerySyntax)(0), // 1: com.coralogix.dataprime.v1.MetadataEnums.QuerySyntax
	(*MetadataEnums)(nil),          // 2: com.coralogix.dataprime.v1.MetadataEnums
	(*QueryMetadata)(nil),          // 3: com.coralogix.dataprime.v1.QueryMetadata
	(*TextQueryMetadata)(nil),      // 4: com.coralogix.dataprime.v1.TextQueryMetadata
	(*timestamppb.Timestamp)(nil),  // 5: google.protobuf.Timestamp
	(*wrapperspb.StringValue)(nil), // 6: google.protobuf.StringValue
}
var file_com_coralogix_dataprime_v1_internal_metadata_proto_depIdxs = []int32{
	5, // 0: com.coralogix.dataprime.v1.QueryMetadata.start_date:type_name -> google.protobuf.Timestamp
	5, // 1: com.coralogix.dataprime.v1.QueryMetadata.end_date:type_name -> google.protobuf.Timestamp
	0, // 2: com.coralogix.dataprime.v1.QueryMetadata.tier:type_name -> com.coralogix.dataprime.v1.MetadataEnums.Tier
	5, // 3: com.coralogix.dataprime.v1.QueryMetadata.now_date:type_name -> google.protobuf.Timestamp
	1, // 4: com.coralogix.dataprime.v1.TextQueryMetadata.syntax:type_name -> com.coralogix.dataprime.v1.MetadataEnums.QuerySyntax
	6, // 5: com.coralogix.dataprime.v1.TextQueryMetadata.default_source:type_name -> google.protobuf.StringValue
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_com_coralogix_dataprime_v1_internal_metadata_proto_init() }
func file_com_coralogix_dataprime_v1_internal_metadata_proto_init() {
	if File_com_coralogix_dataprime_v1_internal_metadata_proto != nil {
		return
	}
	file_com_coralogix_dataprime_v1_internal_metadata_proto_msgTypes[1].OneofWrappers = []any{}
	file_com_coralogix_dataprime_v1_internal_metadata_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_dataprime_v1_internal_metadata_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_dataprime_v1_internal_metadata_proto_goTypes,
		DependencyIndexes: file_com_coralogix_dataprime_v1_internal_metadata_proto_depIdxs,
		EnumInfos:         file_com_coralogix_dataprime_v1_internal_metadata_proto_enumTypes,
		MessageInfos:      file_com_coralogix_dataprime_v1_internal_metadata_proto_msgTypes,
	}.Build()
	File_com_coralogix_dataprime_v1_internal_metadata_proto = out.File
	file_com_coralogix_dataprime_v1_internal_metadata_proto_rawDesc = nil
	file_com_coralogix_dataprime_v1_internal_metadata_proto_goTypes = nil
	file_com_coralogix_dataprime_v1_internal_metadata_proto_depIdxs = nil
}
