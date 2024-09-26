// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.28.1
// source: com/coralogix/dataprime/v1/internal_query_request.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type DpxlScopesPlacementType int32

const (
	DpxlScopesPlacementType_DPXL_SCOPES_PLACEMENT_TYPE_UNSPECIFIED  DpxlScopesPlacementType = 0
	DpxlScopesPlacementType_DPXL_SCOPES_PLACEMENT_TYPE_FIRST_FILTER DpxlScopesPlacementType = 1
	DpxlScopesPlacementType_DPXL_SCOPES_PLACEMENT_TYPE_PLACEHOLDER  DpxlScopesPlacementType = 2
	DpxlScopesPlacementType_DPXL_SCOPES_PLACEMENT_TYPE_SKIP         DpxlScopesPlacementType = 3
)

// Enum value maps for DpxlScopesPlacementType.
var (
	DpxlScopesPlacementType_name = map[int32]string{
		0: "DPXL_SCOPES_PLACEMENT_TYPE_UNSPECIFIED",
		1: "DPXL_SCOPES_PLACEMENT_TYPE_FIRST_FILTER",
		2: "DPXL_SCOPES_PLACEMENT_TYPE_PLACEHOLDER",
		3: "DPXL_SCOPES_PLACEMENT_TYPE_SKIP",
	}
	DpxlScopesPlacementType_value = map[string]int32{
		"DPXL_SCOPES_PLACEMENT_TYPE_UNSPECIFIED":  0,
		"DPXL_SCOPES_PLACEMENT_TYPE_FIRST_FILTER": 1,
		"DPXL_SCOPES_PLACEMENT_TYPE_PLACEHOLDER":  2,
		"DPXL_SCOPES_PLACEMENT_TYPE_SKIP":         3,
	}
)

func (x DpxlScopesPlacementType) Enum() *DpxlScopesPlacementType {
	p := new(DpxlScopesPlacementType)
	*p = x
	return p
}

func (x DpxlScopesPlacementType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DpxlScopesPlacementType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_dataprime_v1_internal_query_request_proto_enumTypes[0].Descriptor()
}

func (DpxlScopesPlacementType) Type() protoreflect.EnumType {
	return &file_com_coralogix_dataprime_v1_internal_query_request_proto_enumTypes[0]
}

func (x DpxlScopesPlacementType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DpxlScopesPlacementType.Descriptor instead.
func (DpxlScopesPlacementType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_internal_query_request_proto_rawDescGZIP(), []int{0}
}

// internal dataprime query request
type QueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Query:
	//
	//	*QueryRequest_AstQuery
	//	*QueryRequest_TextQuery
	Query isQueryRequest_Query `protobuf_oneof:"query"`
	// readable string describing what feature it is used for
	InternalClientId string `protobuf:"bytes,3,opt,name=internal_client_id,json=internalClientId,proto3" json:"internal_client_id,omitempty"`
	// query associated metadata
	Metadata *QueryMetadata `protobuf:"bytes,4,opt,name=metadata,proto3,oneof" json:"metadata,omitempty"`
	// query id, identifies query in DQE, used in tracing (tags.query.id)
	QueryId             *string              `protobuf:"bytes,5,opt,name=query_id,json=queryId,proto3,oneof" json:"query_id,omitempty"`
	DpxlScopesPlacement *DpxlScopesPlacement `protobuf:"bytes,6,opt,name=dpxl_scopes_placement,json=dpxlScopesPlacement,proto3" json:"dpxl_scopes_placement,omitempty"`
}

func (x *QueryRequest) Reset() {
	*x = QueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_dataprime_v1_internal_query_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRequest) ProtoMessage() {}

func (x *QueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_internal_query_request_proto_msgTypes[0]
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
	return file_com_coralogix_dataprime_v1_internal_query_request_proto_rawDescGZIP(), []int{0}
}

func (m *QueryRequest) GetQuery() isQueryRequest_Query {
	if m != nil {
		return m.Query
	}
	return nil
}

func (x *QueryRequest) GetAstQuery() *InternalAstQueryRequest {
	if x, ok := x.GetQuery().(*QueryRequest_AstQuery); ok {
		return x.AstQuery
	}
	return nil
}

func (x *QueryRequest) GetTextQuery() *InternalTextQueryRequest {
	if x, ok := x.GetQuery().(*QueryRequest_TextQuery); ok {
		return x.TextQuery
	}
	return nil
}

func (x *QueryRequest) GetInternalClientId() string {
	if x != nil {
		return x.InternalClientId
	}
	return ""
}

func (x *QueryRequest) GetMetadata() *QueryMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *QueryRequest) GetQueryId() string {
	if x != nil && x.QueryId != nil {
		return *x.QueryId
	}
	return ""
}

func (x *QueryRequest) GetDpxlScopesPlacement() *DpxlScopesPlacement {
	if x != nil {
		return x.DpxlScopesPlacement
	}
	return nil
}

type isQueryRequest_Query interface {
	isQueryRequest_Query()
}

type QueryRequest_AstQuery struct {
	// AST query
	AstQuery *InternalAstQueryRequest `protobuf:"bytes,1,opt,name=ast_query,json=astQuery,proto3,oneof"`
}

type QueryRequest_TextQuery struct {
	// text query
	TextQuery *InternalTextQueryRequest `protobuf:"bytes,2,opt,name=text_query,json=textQuery,proto3,oneof"`
}

func (*QueryRequest_AstQuery) isQueryRequest_Query() {}

func (*QueryRequest_TextQuery) isQueryRequest_Query() {}

type DpxlScopesPlacement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        DpxlScopesPlacementType `protobuf:"varint,1,opt,name=type,proto3,enum=com.coralogix.dataprime.v1.DpxlScopesPlacementType" json:"type,omitempty"`
	Placeholder *string                 `protobuf:"bytes,2,opt,name=placeholder,proto3,oneof" json:"placeholder,omitempty"`
}

func (x *DpxlScopesPlacement) Reset() {
	*x = DpxlScopesPlacement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_dataprime_v1_internal_query_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DpxlScopesPlacement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DpxlScopesPlacement) ProtoMessage() {}

func (x *DpxlScopesPlacement) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_internal_query_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DpxlScopesPlacement.ProtoReflect.Descriptor instead.
func (*DpxlScopesPlacement) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_internal_query_request_proto_rawDescGZIP(), []int{1}
}

func (x *DpxlScopesPlacement) GetType() DpxlScopesPlacementType {
	if x != nil {
		return x.Type
	}
	return DpxlScopesPlacementType_DPXL_SCOPES_PLACEMENT_TYPE_UNSPECIFIED
}

func (x *DpxlScopesPlacement) GetPlaceholder() string {
	if x != nil && x.Placeholder != nil {
		return *x.Placeholder
	}
	return ""
}

type InternalAstQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// byte representation of dataprime AST
	Ast *SerializedDataprime `protobuf:"bytes,1,opt,name=ast,proto3" json:"ast,omitempty"`
}

func (x *InternalAstQueryRequest) Reset() {
	*x = InternalAstQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_dataprime_v1_internal_query_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalAstQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalAstQueryRequest) ProtoMessage() {}

func (x *InternalAstQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_internal_query_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalAstQueryRequest.ProtoReflect.Descriptor instead.
func (*InternalAstQueryRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_internal_query_request_proto_rawDescGZIP(), []int{2}
}

func (x *InternalAstQueryRequest) GetAst() *SerializedDataprime {
	if x != nil {
		return x.Ast
	}
	return nil
}

type InternalTextQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// dataprime query
	Text *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// text query associated metadata
	Metadata *TextQueryMetadata `protobuf:"bytes,2,opt,name=metadata,proto3,oneof" json:"metadata,omitempty"`
}

func (x *InternalTextQueryRequest) Reset() {
	*x = InternalTextQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_dataprime_v1_internal_query_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalTextQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalTextQueryRequest) ProtoMessage() {}

func (x *InternalTextQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_internal_query_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalTextQueryRequest.ProtoReflect.Descriptor instead.
func (*InternalTextQueryRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_internal_query_request_proto_rawDescGZIP(), []int{3}
}

func (x *InternalTextQueryRequest) GetText() *wrapperspb.StringValue {
	if x != nil {
		return x.Text
	}
	return nil
}

func (x *InternalTextQueryRequest) GetMetadata() *TextQueryMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_com_coralogix_dataprime_v1_internal_query_request_proto protoreflect.FileDescriptor

var file_com_coralogix_dataprime_v1_internal_query_request_proto_rawDesc = []byte{
	0x0a, 0x37, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69,
	0x6d, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x32, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69,
	0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xdb, 0x03, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x52, 0x0a, 0x09, 0x61, 0x73, 0x74, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x73, 0x74, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x08, 0x61, 0x73,
	0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x55, 0x0a, 0x0a, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70,
	0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x54, 0x65, 0x78, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x48, 0x00, 0x52, 0x09, 0x74, 0x65, 0x78, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2c, 0x0a,
	0x12, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x4a, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x01, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x07, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x63, 0x0a, 0x15, 0x64, 0x70, 0x78, 0x6c, 0x5f,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x70, 0x78, 0x6c, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x13, 0x64, 0x70, 0x78, 0x6c, 0x53, 0x63, 0x6f,
	0x70, 0x65, 0x73, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x07, 0x0a, 0x05,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x22,
	0x95, 0x01, 0x0a, 0x13, 0x44, 0x70, 0x78, 0x6c, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x47, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x70, 0x78, 0x6c, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x50, 0x6c, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x25, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x22, 0x5c, 0x0a, 0x17, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x41, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x41, 0x0a, 0x03, 0x61, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65,
	0x52, 0x03, 0x61, 0x73, 0x74, 0x22, 0xa9, 0x01, 0x0a, 0x18, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x54, 0x65, 0x78, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x4e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2a, 0xc3, 0x01, 0x0a, 0x17, 0x44, 0x70, 0x78, 0x6c, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73,
	0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a,
	0x26, 0x44, 0x50, 0x58, 0x4c, 0x5f, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x53, 0x5f, 0x50, 0x4c, 0x41,
	0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x2b, 0x0a, 0x27, 0x44, 0x50, 0x58,
	0x4c, 0x5f, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x53, 0x5f, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x4d, 0x45,
	0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x49, 0x52, 0x53, 0x54, 0x5f, 0x46, 0x49,
	0x4c, 0x54, 0x45, 0x52, 0x10, 0x01, 0x12, 0x2a, 0x0a, 0x26, 0x44, 0x50, 0x58, 0x4c, 0x5f, 0x53,
	0x43, 0x4f, 0x50, 0x45, 0x53, 0x5f, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x48, 0x4f, 0x4c, 0x44, 0x45, 0x52,
	0x10, 0x02, 0x12, 0x23, 0x0a, 0x1f, 0x44, 0x50, 0x58, 0x4c, 0x5f, 0x53, 0x43, 0x4f, 0x50, 0x45,
	0x53, 0x5f, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x53, 0x4b, 0x49, 0x50, 0x10, 0x03, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_dataprime_v1_internal_query_request_proto_rawDescOnce sync.Once
	file_com_coralogix_dataprime_v1_internal_query_request_proto_rawDescData = file_com_coralogix_dataprime_v1_internal_query_request_proto_rawDesc
)

func file_com_coralogix_dataprime_v1_internal_query_request_proto_rawDescGZIP() []byte {
	file_com_coralogix_dataprime_v1_internal_query_request_proto_rawDescOnce.Do(func() {
		file_com_coralogix_dataprime_v1_internal_query_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_dataprime_v1_internal_query_request_proto_rawDescData)
	})
	return file_com_coralogix_dataprime_v1_internal_query_request_proto_rawDescData
}

var file_com_coralogix_dataprime_v1_internal_query_request_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogix_dataprime_v1_internal_query_request_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_com_coralogix_dataprime_v1_internal_query_request_proto_goTypes = []interface{}{
	(DpxlScopesPlacementType)(0),     // 0: com.coralogix.dataprime.v1.DpxlScopesPlacementType
	(*QueryRequest)(nil),             // 1: com.coralogix.dataprime.v1.QueryRequest
	(*DpxlScopesPlacement)(nil),      // 2: com.coralogix.dataprime.v1.DpxlScopesPlacement
	(*InternalAstQueryRequest)(nil),  // 3: com.coralogix.dataprime.v1.InternalAstQueryRequest
	(*InternalTextQueryRequest)(nil), // 4: com.coralogix.dataprime.v1.InternalTextQueryRequest
	(*QueryMetadata)(nil),            // 5: com.coralogix.dataprime.v1.QueryMetadata
	(*SerializedDataprime)(nil),      // 6: com.coralogix.dataprime.v1.SerializedDataprime
	(*wrapperspb.StringValue)(nil),   // 7: google.protobuf.StringValue
	(*TextQueryMetadata)(nil),        // 8: com.coralogix.dataprime.v1.TextQueryMetadata
}
var file_com_coralogix_dataprime_v1_internal_query_request_proto_depIdxs = []int32{
	3, // 0: com.coralogix.dataprime.v1.QueryRequest.ast_query:type_name -> com.coralogix.dataprime.v1.InternalAstQueryRequest
	4, // 1: com.coralogix.dataprime.v1.QueryRequest.text_query:type_name -> com.coralogix.dataprime.v1.InternalTextQueryRequest
	5, // 2: com.coralogix.dataprime.v1.QueryRequest.metadata:type_name -> com.coralogix.dataprime.v1.QueryMetadata
	2, // 3: com.coralogix.dataprime.v1.QueryRequest.dpxl_scopes_placement:type_name -> com.coralogix.dataprime.v1.DpxlScopesPlacement
	0, // 4: com.coralogix.dataprime.v1.DpxlScopesPlacement.type:type_name -> com.coralogix.dataprime.v1.DpxlScopesPlacementType
	6, // 5: com.coralogix.dataprime.v1.InternalAstQueryRequest.ast:type_name -> com.coralogix.dataprime.v1.SerializedDataprime
	7, // 6: com.coralogix.dataprime.v1.InternalTextQueryRequest.text:type_name -> google.protobuf.StringValue
	8, // 7: com.coralogix.dataprime.v1.InternalTextQueryRequest.metadata:type_name -> com.coralogix.dataprime.v1.TextQueryMetadata
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_com_coralogix_dataprime_v1_internal_query_request_proto_init() }
func file_com_coralogix_dataprime_v1_internal_query_request_proto_init() {
	if File_com_coralogix_dataprime_v1_internal_query_request_proto != nil {
		return
	}
	file_com_coralogix_dataprime_v1_internal_metadata_proto_init()
	file_com_coralogix_dataprime_v1_query_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_dataprime_v1_internal_query_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_com_coralogix_dataprime_v1_internal_query_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DpxlScopesPlacement); i {
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
		file_com_coralogix_dataprime_v1_internal_query_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalAstQueryRequest); i {
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
		file_com_coralogix_dataprime_v1_internal_query_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalTextQueryRequest); i {
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
	file_com_coralogix_dataprime_v1_internal_query_request_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*QueryRequest_AstQuery)(nil),
		(*QueryRequest_TextQuery)(nil),
	}
	file_com_coralogix_dataprime_v1_internal_query_request_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_com_coralogix_dataprime_v1_internal_query_request_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_dataprime_v1_internal_query_request_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_dataprime_v1_internal_query_request_proto_goTypes,
		DependencyIndexes: file_com_coralogix_dataprime_v1_internal_query_request_proto_depIdxs,
		EnumInfos:         file_com_coralogix_dataprime_v1_internal_query_request_proto_enumTypes,
		MessageInfos:      file_com_coralogix_dataprime_v1_internal_query_request_proto_msgTypes,
	}.Build()
	File_com_coralogix_dataprime_v1_internal_query_request_proto = out.File
	file_com_coralogix_dataprime_v1_internal_query_request_proto_rawDesc = nil
	file_com_coralogix_dataprime_v1_internal_query_request_proto_goTypes = nil
	file_com_coralogix_dataprime_v1_internal_query_request_proto_depIdxs = nil
}
