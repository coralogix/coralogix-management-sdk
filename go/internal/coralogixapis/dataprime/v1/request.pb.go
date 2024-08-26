// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.3
// source: com/coralogixapis/dataprime/v1/request.proto

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

// tier on which query runs, default: TIER_FREQUENT_SEARCH
type Metadata_Tier int32

const (
	Metadata_TIER_UNSPECIFIED     Metadata_Tier = 0
	Metadata_TIER_ARCHIVE         Metadata_Tier = 1
	Metadata_TIER_FREQUENT_SEARCH Metadata_Tier = 2
)

// Enum value maps for Metadata_Tier.
var (
	Metadata_Tier_name = map[int32]string{
		0: "TIER_UNSPECIFIED",
		1: "TIER_ARCHIVE",
		2: "TIER_FREQUENT_SEARCH",
	}
	Metadata_Tier_value = map[string]int32{
		"TIER_UNSPECIFIED":     0,
		"TIER_ARCHIVE":         1,
		"TIER_FREQUENT_SEARCH": 2,
	}
)

func (x Metadata_Tier) Enum() *Metadata_Tier {
	p := new(Metadata_Tier)
	*p = x
	return p
}

func (x Metadata_Tier) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Metadata_Tier) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_dataprime_v1_request_proto_enumTypes[0].Descriptor()
}

func (Metadata_Tier) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_dataprime_v1_request_proto_enumTypes[0]
}

func (x Metadata_Tier) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Metadata_Tier.Descriptor instead.
func (Metadata_Tier) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_dataprime_v1_request_proto_rawDescGZIP(), []int{1, 0}
}

// syntax of the query, default: QUERY_SYNTAX_DATAPRIME
type Metadata_QuerySyntax int32

const (
	Metadata_QUERY_SYNTAX_UNSPECIFIED Metadata_QuerySyntax = 0
	Metadata_QUERY_SYNTAX_LUCENE      Metadata_QuerySyntax = 1
	Metadata_QUERY_SYNTAX_DATAPRIME   Metadata_QuerySyntax = 2
)

// Enum value maps for Metadata_QuerySyntax.
var (
	Metadata_QuerySyntax_name = map[int32]string{
		0: "QUERY_SYNTAX_UNSPECIFIED",
		1: "QUERY_SYNTAX_LUCENE",
		2: "QUERY_SYNTAX_DATAPRIME",
	}
	Metadata_QuerySyntax_value = map[string]int32{
		"QUERY_SYNTAX_UNSPECIFIED": 0,
		"QUERY_SYNTAX_LUCENE":      1,
		"QUERY_SYNTAX_DATAPRIME":   2,
	}
)

func (x Metadata_QuerySyntax) Enum() *Metadata_QuerySyntax {
	p := new(Metadata_QuerySyntax)
	*p = x
	return p
}

func (x Metadata_QuerySyntax) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Metadata_QuerySyntax) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_dataprime_v1_request_proto_enumTypes[1].Descriptor()
}

func (Metadata_QuerySyntax) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_dataprime_v1_request_proto_enumTypes[1]
}

func (x Metadata_QuerySyntax) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Metadata_QuerySyntax.Descriptor instead.
func (Metadata_QuerySyntax) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_dataprime_v1_request_proto_rawDescGZIP(), []int{1, 1}
}

// dataprime text query request
type QueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// query for which you seek results
	Query *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// configuration of query execution
	Metadata *Metadata `protobuf:"bytes,2,opt,name=metadata,proto3,oneof" json:"metadata,omitempty"`
}

func (x *QueryRequest) Reset() {
	*x = QueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_dataprime_v1_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRequest) ProtoMessage() {}

func (x *QueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dataprime_v1_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRequest.ProtoReflect.Descriptor instead.
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dataprime_v1_request_proto_rawDescGZIP(), []int{0}
}

func (x *QueryRequest) GetQuery() *wrapperspb.StringValue {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *QueryRequest) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// configuration of query execution
type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// beginning of the time range for the query, default: end - 15 min or current time - 15 min if end is not defined
	StartDate *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start_date,json=startDate,proto3,oneof" json:"start_date,omitempty"`
	// end of the time range for the query, default: start + 15 min or current time if start is not defined
	EndDate *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end_date,json=endDate,proto3,oneof" json:"end_date,omitempty"`
	// default value for source to be used when source is omitted in a query
	DefaultSource *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=default_source,json=defaultSource,proto3,oneof" json:"default_source,omitempty"`
	// tier on which query runs
	Tier *Metadata_Tier `protobuf:"varint,4,opt,name=tier,proto3,enum=com.coralogixapis.dataprime.v1.Metadata_Tier,oneof" json:"tier,omitempty"`
	// which syntax query is written in
	Syntax *Metadata_QuerySyntax `protobuf:"varint,5,opt,name=syntax,proto3,enum=com.coralogixapis.dataprime.v1.Metadata_QuerySyntax,oneof" json:"syntax,omitempty"`
	// limit number of results, default: 2000; max for TIER_FREQUENT_SEARCH: 12000;  max for TIER_ARCHIVE: 50000
	Limit *int32 `protobuf:"varint,6,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	// prohibit using unknown fields, ones which were not detected in the ingested data, default = false
	StrictFieldsValidation *bool `protobuf:"varint,7,opt,name=strict_fields_validation,json=strictFieldsValidation,proto3,oneof" json:"strict_fields_validation,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_dataprime_v1_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dataprime_v1_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dataprime_v1_request_proto_rawDescGZIP(), []int{1}
}

func (x *Metadata) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *Metadata) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

func (x *Metadata) GetDefaultSource() *wrapperspb.StringValue {
	if x != nil {
		return x.DefaultSource
	}
	return nil
}

func (x *Metadata) GetTier() Metadata_Tier {
	if x != nil && x.Tier != nil {
		return *x.Tier
	}
	return Metadata_TIER_UNSPECIFIED
}

func (x *Metadata) GetSyntax() Metadata_QuerySyntax {
	if x != nil && x.Syntax != nil {
		return *x.Syntax
	}
	return Metadata_QUERY_SYNTAX_UNSPECIFIED
}

func (x *Metadata) GetLimit() int32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *Metadata) GetStrictFieldsValidation() bool {
	if x != nil && x.StrictFieldsValidation != nil {
		return *x.StrictFieldsValidation
	}
	return false
}

var File_com_coralogixapis_dataprime_v1_request_proto protoreflect.FileDescriptor

var file_com_coralogixapis_dataprime_v1_request_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9a, 0x01, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x32, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x49, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70,
	0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x48, 0x00, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0xdb, 0x05, 0x0a,
	0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3e, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x3a, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61,
	0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x48, 0x0a, 0x0e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x02, 0x52, 0x0d, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x46, 0x0a, 0x04, 0x74, 0x69, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x69, 0x65, 0x72, 0x48, 0x03, 0x52, 0x04,
	0x74, 0x69, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x51, 0x0a, 0x06, 0x73, 0x79, 0x6e, 0x74, 0x61,
	0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x48, 0x04, 0x52,
	0x06, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x48, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x88, 0x01, 0x01, 0x12, 0x3d, 0x0a, 0x18, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x48, 0x06, 0x52, 0x16, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x88, 0x01, 0x01, 0x22, 0x48, 0x0a, 0x04, 0x54, 0x69, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x10,
	0x54, 0x49, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x49, 0x45, 0x52, 0x5f, 0x41, 0x52, 0x43, 0x48, 0x49,
	0x56, 0x45, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x49, 0x45, 0x52, 0x5f, 0x46, 0x52, 0x45,
	0x51, 0x55, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x10, 0x02, 0x22, 0x60,
	0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x12, 0x1c, 0x0a,
	0x18, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x53, 0x59, 0x4e, 0x54, 0x41, 0x58, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x51,
	0x55, 0x45, 0x52, 0x59, 0x5f, 0x53, 0x59, 0x4e, 0x54, 0x41, 0x58, 0x5f, 0x4c, 0x55, 0x43, 0x45,
	0x4e, 0x45, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x53, 0x59,
	0x4e, 0x54, 0x41, 0x58, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x50, 0x52, 0x49, 0x4d, 0x45, 0x10, 0x02,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x42, 0x11, 0x0a, 0x0f,
	0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x74, 0x69, 0x65, 0x72, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x79, 0x6e,
	0x74, 0x61, 0x78, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x1b, 0x0a,
	0x19, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x5f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_com_coralogixapis_dataprime_v1_request_proto_rawDescOnce sync.Once
	file_com_coralogixapis_dataprime_v1_request_proto_rawDescData = file_com_coralogixapis_dataprime_v1_request_proto_rawDesc
)

func file_com_coralogixapis_dataprime_v1_request_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_dataprime_v1_request_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_dataprime_v1_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_dataprime_v1_request_proto_rawDescData)
	})
	return file_com_coralogixapis_dataprime_v1_request_proto_rawDescData
}

var file_com_coralogixapis_dataprime_v1_request_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_com_coralogixapis_dataprime_v1_request_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogixapis_dataprime_v1_request_proto_goTypes = []interface{}{
	(Metadata_Tier)(0),             // 0: com.coralogixapis.dataprime.v1.Metadata.Tier
	(Metadata_QuerySyntax)(0),      // 1: com.coralogixapis.dataprime.v1.Metadata.QuerySyntax
	(*QueryRequest)(nil),           // 2: com.coralogixapis.dataprime.v1.QueryRequest
	(*Metadata)(nil),               // 3: com.coralogixapis.dataprime.v1.Metadata
	(*wrapperspb.StringValue)(nil), // 4: google.protobuf.StringValue
	(*timestamppb.Timestamp)(nil),  // 5: google.protobuf.Timestamp
}
var file_com_coralogixapis_dataprime_v1_request_proto_depIdxs = []int32{
	4, // 0: com.coralogixapis.dataprime.v1.QueryRequest.query:type_name -> google.protobuf.StringValue
	3, // 1: com.coralogixapis.dataprime.v1.QueryRequest.metadata:type_name -> com.coralogixapis.dataprime.v1.Metadata
	5, // 2: com.coralogixapis.dataprime.v1.Metadata.start_date:type_name -> google.protobuf.Timestamp
	5, // 3: com.coralogixapis.dataprime.v1.Metadata.end_date:type_name -> google.protobuf.Timestamp
	4, // 4: com.coralogixapis.dataprime.v1.Metadata.default_source:type_name -> google.protobuf.StringValue
	0, // 5: com.coralogixapis.dataprime.v1.Metadata.tier:type_name -> com.coralogixapis.dataprime.v1.Metadata.Tier
	1, // 6: com.coralogixapis.dataprime.v1.Metadata.syntax:type_name -> com.coralogixapis.dataprime.v1.Metadata.QuerySyntax
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_dataprime_v1_request_proto_init() }
func file_com_coralogixapis_dataprime_v1_request_proto_init() {
	if File_com_coralogixapis_dataprime_v1_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_dataprime_v1_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRequest); i {
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
		file_com_coralogixapis_dataprime_v1_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
	file_com_coralogixapis_dataprime_v1_request_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_com_coralogixapis_dataprime_v1_request_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_dataprime_v1_request_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_dataprime_v1_request_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_dataprime_v1_request_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_dataprime_v1_request_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_dataprime_v1_request_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_dataprime_v1_request_proto = out.File
	file_com_coralogixapis_dataprime_v1_request_proto_rawDesc = nil
	file_com_coralogixapis_dataprime_v1_request_proto_goTypes = nil
	file_com_coralogixapis_dataprime_v1_request_proto_depIdxs = nil
}
