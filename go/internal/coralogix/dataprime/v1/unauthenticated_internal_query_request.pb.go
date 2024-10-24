// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: com/coralogix/dataprime/v1/unauthenticated_internal_query_request.proto

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

// internal dataprime query request
type UnauthenticatedInternalDataprimeQueryServiceQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Query:
	//
	//	*UnauthenticatedInternalDataprimeQueryServiceQueryRequest_AstQuery
	//	*UnauthenticatedInternalDataprimeQueryServiceQueryRequest_TextQuery
	Query isUnauthenticatedInternalDataprimeQueryServiceQueryRequest_Query `protobuf_oneof:"query"`
	// readable string describing what feature it is used for
	InternalClientId string `protobuf:"bytes,3,opt,name=internal_client_id,json=internalClientId,proto3" json:"internal_client_id,omitempty"`
	// query associated metadata
	Metadata *QueryMetadata `protobuf:"bytes,4,opt,name=metadata,proto3,oneof" json:"metadata,omitempty"`
	// team id for which the query is run
	TeamId *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	// query id, identifies query in DQE, used in tracing (tags.query.id)
	QueryId *string `protobuf:"bytes,7,opt,name=query_id,json=queryId,proto3,oneof" json:"query_id,omitempty"`
}

func (x *UnauthenticatedInternalDataprimeQueryServiceQueryRequest) Reset() {
	*x = UnauthenticatedInternalDataprimeQueryServiceQueryRequest{}
	mi := &file_com_coralogix_dataprime_v1_unauthenticated_internal_query_request_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnauthenticatedInternalDataprimeQueryServiceQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnauthenticatedInternalDataprimeQueryServiceQueryRequest) ProtoMessage() {}

func (x *UnauthenticatedInternalDataprimeQueryServiceQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_unauthenticated_internal_query_request_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnauthenticatedInternalDataprimeQueryServiceQueryRequest.ProtoReflect.Descriptor instead.
func (*UnauthenticatedInternalDataprimeQueryServiceQueryRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_unauthenticated_internal_query_request_proto_rawDescGZIP(), []int{0}
}

func (m *UnauthenticatedInternalDataprimeQueryServiceQueryRequest) GetQuery() isUnauthenticatedInternalDataprimeQueryServiceQueryRequest_Query {
	if m != nil {
		return m.Query
	}
	return nil
}

func (x *UnauthenticatedInternalDataprimeQueryServiceQueryRequest) GetAstQuery() *InternalAstQueryRequest {
	if x, ok := x.GetQuery().(*UnauthenticatedInternalDataprimeQueryServiceQueryRequest_AstQuery); ok {
		return x.AstQuery
	}
	return nil
}

func (x *UnauthenticatedInternalDataprimeQueryServiceQueryRequest) GetTextQuery() *InternalTextQueryRequest {
	if x, ok := x.GetQuery().(*UnauthenticatedInternalDataprimeQueryServiceQueryRequest_TextQuery); ok {
		return x.TextQuery
	}
	return nil
}

func (x *UnauthenticatedInternalDataprimeQueryServiceQueryRequest) GetInternalClientId() string {
	if x != nil {
		return x.InternalClientId
	}
	return ""
}

func (x *UnauthenticatedInternalDataprimeQueryServiceQueryRequest) GetMetadata() *QueryMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *UnauthenticatedInternalDataprimeQueryServiceQueryRequest) GetTeamId() *wrapperspb.StringValue {
	if x != nil {
		return x.TeamId
	}
	return nil
}

func (x *UnauthenticatedInternalDataprimeQueryServiceQueryRequest) GetQueryId() string {
	if x != nil && x.QueryId != nil {
		return *x.QueryId
	}
	return ""
}

type isUnauthenticatedInternalDataprimeQueryServiceQueryRequest_Query interface {
	isUnauthenticatedInternalDataprimeQueryServiceQueryRequest_Query()
}

type UnauthenticatedInternalDataprimeQueryServiceQueryRequest_AstQuery struct {
	// AST query
	AstQuery *InternalAstQueryRequest `protobuf:"bytes,1,opt,name=ast_query,json=astQuery,proto3,oneof"`
}

type UnauthenticatedInternalDataprimeQueryServiceQueryRequest_TextQuery struct {
	// text query
	TextQuery *InternalTextQueryRequest `protobuf:"bytes,2,opt,name=text_query,json=textQuery,proto3,oneof"`
}

func (*UnauthenticatedInternalDataprimeQueryServiceQueryRequest_AstQuery) isUnauthenticatedInternalDataprimeQueryServiceQueryRequest_Query() {
}

func (*UnauthenticatedInternalDataprimeQueryServiceQueryRequest_TextQuery) isUnauthenticatedInternalDataprimeQueryServiceQueryRequest_Query() {
}

var File_com_coralogix_dataprime_v1_unauthenticated_internal_query_request_proto protoreflect.FileDescriptor

var file_com_coralogix_dataprime_v1_unauthenticated_internal_query_request_proto_rawDesc = []byte{
	0x0a, 0x47, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x6e, 0x61,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69,
	0x6d, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x32, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x37, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69,
	0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xdf, 0x03, 0x0a, 0x38, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x61,
	0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x52, 0x0a, 0x09, 0x61, 0x73, 0x74, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x08, 0x61, 0x73, 0x74, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x55, 0x0a, 0x0a, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x65, 0x78,
	0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52,
	0x09, 0x74, 0x65, 0x78, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2c, 0x0a, 0x12, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x4a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70,
	0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x01, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x88, 0x01, 0x01, 0x12, 0x35, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x08, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52,
	0x07, 0x71, 0x75, 0x65, 0x72, 0x79, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x4a, 0x04,
	0x08, 0x06, 0x10, 0x07, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_dataprime_v1_unauthenticated_internal_query_request_proto_rawDescOnce sync.Once
	file_com_coralogix_dataprime_v1_unauthenticated_internal_query_request_proto_rawDescData = file_com_coralogix_dataprime_v1_unauthenticated_internal_query_request_proto_rawDesc
)

func file_com_coralogix_dataprime_v1_unauthenticated_internal_query_request_proto_rawDescGZIP() []byte {
	file_com_coralogix_dataprime_v1_unauthenticated_internal_query_request_proto_rawDescOnce.Do(func() {
		file_com_coralogix_dataprime_v1_unauthenticated_internal_query_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_dataprime_v1_unauthenticated_internal_query_request_proto_rawDescData)
	})
	return file_com_coralogix_dataprime_v1_unauthenticated_internal_query_request_proto_rawDescData
}

var file_com_coralogix_dataprime_v1_unauthenticated_internal_query_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogix_dataprime_v1_unauthenticated_internal_query_request_proto_goTypes = []any{
	(*UnauthenticatedInternalDataprimeQueryServiceQueryRequest)(nil), // 0: com.coralogix.dataprime.v1.UnauthenticatedInternalDataprimeQueryServiceQueryRequest
	(*InternalAstQueryRequest)(nil),                                  // 1: com.coralogix.dataprime.v1.InternalAstQueryRequest
	(*InternalTextQueryRequest)(nil),                                 // 2: com.coralogix.dataprime.v1.InternalTextQueryRequest
	(*QueryMetadata)(nil),                                            // 3: com.coralogix.dataprime.v1.QueryMetadata
	(*wrapperspb.StringValue)(nil),                                   // 4: google.protobuf.StringValue
}
var file_com_coralogix_dataprime_v1_unauthenticated_internal_query_request_proto_depIdxs = []int32{
	1, // 0: com.coralogix.dataprime.v1.UnauthenticatedInternalDataprimeQueryServiceQueryRequest.ast_query:type_name -> com.coralogix.dataprime.v1.InternalAstQueryRequest
	2, // 1: com.coralogix.dataprime.v1.UnauthenticatedInternalDataprimeQueryServiceQueryRequest.text_query:type_name -> com.coralogix.dataprime.v1.InternalTextQueryRequest
	3, // 2: com.coralogix.dataprime.v1.UnauthenticatedInternalDataprimeQueryServiceQueryRequest.metadata:type_name -> com.coralogix.dataprime.v1.QueryMetadata
	4, // 3: com.coralogix.dataprime.v1.UnauthenticatedInternalDataprimeQueryServiceQueryRequest.team_id:type_name -> google.protobuf.StringValue
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_com_coralogix_dataprime_v1_unauthenticated_internal_query_request_proto_init() }
func file_com_coralogix_dataprime_v1_unauthenticated_internal_query_request_proto_init() {
	if File_com_coralogix_dataprime_v1_unauthenticated_internal_query_request_proto != nil {
		return
	}
	file_com_coralogix_dataprime_v1_internal_metadata_proto_init()
	file_com_coralogix_dataprime_v1_internal_query_request_proto_init()
	file_com_coralogix_dataprime_v1_unauthenticated_internal_query_request_proto_msgTypes[0].OneofWrappers = []any{
		(*UnauthenticatedInternalDataprimeQueryServiceQueryRequest_AstQuery)(nil),
		(*UnauthenticatedInternalDataprimeQueryServiceQueryRequest_TextQuery)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_dataprime_v1_unauthenticated_internal_query_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_dataprime_v1_unauthenticated_internal_query_request_proto_goTypes,
		DependencyIndexes: file_com_coralogix_dataprime_v1_unauthenticated_internal_query_request_proto_depIdxs,
		MessageInfos:      file_com_coralogix_dataprime_v1_unauthenticated_internal_query_request_proto_msgTypes,
	}.Build()
	File_com_coralogix_dataprime_v1_unauthenticated_internal_query_request_proto = out.File
	file_com_coralogix_dataprime_v1_unauthenticated_internal_query_request_proto_rawDesc = nil
	file_com_coralogix_dataprime_v1_unauthenticated_internal_query_request_proto_goTypes = nil
	file_com_coralogix_dataprime_v1_unauthenticated_internal_query_request_proto_depIdxs = nil
}
