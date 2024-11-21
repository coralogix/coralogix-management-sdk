// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: com/coralogixapis/dataprime/background_queries/v1/request.proto

package v1

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type QuerySyntax int32

const (
	QuerySyntax_QUERY_SYNTAX_UNSPECIFIED QuerySyntax = 0
	QuerySyntax_QUERY_SYNTAX_LUCENE      QuerySyntax = 1
	QuerySyntax_QUERY_SYNTAX_DATAPRIME   QuerySyntax = 2
)

// Enum value maps for QuerySyntax.
var (
	QuerySyntax_name = map[int32]string{
		0: "QUERY_SYNTAX_UNSPECIFIED",
		1: "QUERY_SYNTAX_LUCENE",
		2: "QUERY_SYNTAX_DATAPRIME",
	}
	QuerySyntax_value = map[string]int32{
		"QUERY_SYNTAX_UNSPECIFIED": 0,
		"QUERY_SYNTAX_LUCENE":      1,
		"QUERY_SYNTAX_DATAPRIME":   2,
	}
)

func (x QuerySyntax) Enum() *QuerySyntax {
	p := new(QuerySyntax)
	*p = x
	return p
}

func (x QuerySyntax) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QuerySyntax) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_dataprime_background_queries_v1_request_proto_enumTypes[0].Descriptor()
}

func (QuerySyntax) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_dataprime_background_queries_v1_request_proto_enumTypes[0]
}

func (x QuerySyntax) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QuerySyntax.Descriptor instead.
func (QuerySyntax) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_dataprime_background_queries_v1_request_proto_rawDescGZIP(), []int{0}
}

type RunQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query     *wrappers.StringValue `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Syntax    QuerySyntax           `protobuf:"varint,2,opt,name=syntax,proto3,enum=com.coralogixapis.dataprime.background_queries.v1.QuerySyntax" json:"syntax,omitempty"`
	StartDate *timestamp.Timestamp  `protobuf:"bytes,3,opt,name=start_date,json=startDate,proto3,oneof" json:"start_date,omitempty"`
	EndDate   *timestamp.Timestamp  `protobuf:"bytes,4,opt,name=end_date,json=endDate,proto3,oneof" json:"end_date,omitempty"`
	NowDate   *timestamp.Timestamp  `protobuf:"bytes,5,opt,name=now_date,json=nowDate,proto3,oneof" json:"now_date,omitempty"`
}

func (x *RunQueryRequest) Reset() {
	*x = RunQueryRequest{}
	mi := &file_com_coralogixapis_dataprime_background_queries_v1_request_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RunQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunQueryRequest) ProtoMessage() {}

func (x *RunQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dataprime_background_queries_v1_request_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunQueryRequest.ProtoReflect.Descriptor instead.
func (*RunQueryRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dataprime_background_queries_v1_request_proto_rawDescGZIP(), []int{0}
}

func (x *RunQueryRequest) GetQuery() *wrappers.StringValue {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *RunQueryRequest) GetSyntax() QuerySyntax {
	if x != nil {
		return x.Syntax
	}
	return QuerySyntax_QUERY_SYNTAX_UNSPECIFIED
}

func (x *RunQueryRequest) GetStartDate() *timestamp.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *RunQueryRequest) GetEndDate() *timestamp.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

func (x *RunQueryRequest) GetNowDate() *timestamp.Timestamp {
	if x != nil {
		return x.NowDate
	}
	return nil
}

type GetStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueryId string `protobuf:"bytes,1,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
}

func (x *GetStatusRequest) Reset() {
	*x = GetStatusRequest{}
	mi := &file_com_coralogixapis_dataprime_background_queries_v1_request_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatusRequest) ProtoMessage() {}

func (x *GetStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dataprime_background_queries_v1_request_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatusRequest.ProtoReflect.Descriptor instead.
func (*GetStatusRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dataprime_background_queries_v1_request_proto_rawDescGZIP(), []int{1}
}

func (x *GetStatusRequest) GetQueryId() string {
	if x != nil {
		return x.QueryId
	}
	return ""
}

type GetResultsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueryId string `protobuf:"bytes,1,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
}

func (x *GetResultsRequest) Reset() {
	*x = GetResultsRequest{}
	mi := &file_com_coralogixapis_dataprime_background_queries_v1_request_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetResultsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResultsRequest) ProtoMessage() {}

func (x *GetResultsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dataprime_background_queries_v1_request_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResultsRequest.ProtoReflect.Descriptor instead.
func (*GetResultsRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dataprime_background_queries_v1_request_proto_rawDescGZIP(), []int{2}
}

func (x *GetResultsRequest) GetQueryId() string {
	if x != nil {
		return x.QueryId
	}
	return ""
}

var File_com_coralogixapis_dataprime_background_queries_v1_request_proto protoreflect.FileDescriptor

var file_com_coralogixapis_dataprime_background_queries_v1_request_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x62, 0x61,
	0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x31, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x62,
	0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x02, 0x0a, 0x0f, 0x52, 0x75, 0x6e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x56, 0x0a,
	0x06, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3e, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x62, 0x61, 0x63, 0x6b,
	0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x52, 0x06, 0x73,
	0x79, 0x6e, 0x74, 0x61, 0x78, 0x12, 0x3e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61,
	0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x3a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x3a, 0x0a, 0x08, 0x6e, 0x6f, 0x77, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48,
	0x02, 0x52, 0x07, 0x6e, 0x6f, 0x77, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6e, 0x6f,
	0x77, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x22, 0x2d, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x49, 0x64, 0x2a, 0x60, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x79,
	0x6e, 0x74, 0x61, 0x78, 0x12, 0x1c, 0x0a, 0x18, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x53, 0x59,
	0x4e, 0x54, 0x41, 0x58, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x53, 0x59, 0x4e, 0x54,
	0x41, 0x58, 0x5f, 0x4c, 0x55, 0x43, 0x45, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x51,
	0x55, 0x45, 0x52, 0x59, 0x5f, 0x53, 0x59, 0x4e, 0x54, 0x41, 0x58, 0x5f, 0x44, 0x41, 0x54, 0x41,
	0x50, 0x52, 0x49, 0x4d, 0x45, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_dataprime_background_queries_v1_request_proto_rawDescOnce sync.Once
	file_com_coralogixapis_dataprime_background_queries_v1_request_proto_rawDescData = file_com_coralogixapis_dataprime_background_queries_v1_request_proto_rawDesc
)

func file_com_coralogixapis_dataprime_background_queries_v1_request_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_dataprime_background_queries_v1_request_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_dataprime_background_queries_v1_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_dataprime_background_queries_v1_request_proto_rawDescData)
	})
	return file_com_coralogixapis_dataprime_background_queries_v1_request_proto_rawDescData
}

var file_com_coralogixapis_dataprime_background_queries_v1_request_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_dataprime_background_queries_v1_request_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogixapis_dataprime_background_queries_v1_request_proto_goTypes = []any{
	(QuerySyntax)(0),             // 0: com.coralogixapis.dataprime.background_queries.v1.QuerySyntax
	(*RunQueryRequest)(nil),      // 1: com.coralogixapis.dataprime.background_queries.v1.RunQueryRequest
	(*GetStatusRequest)(nil),     // 2: com.coralogixapis.dataprime.background_queries.v1.GetStatusRequest
	(*GetResultsRequest)(nil),    // 3: com.coralogixapis.dataprime.background_queries.v1.GetResultsRequest
	(*wrappers.StringValue)(nil), // 4: google.protobuf.StringValue
	(*timestamp.Timestamp)(nil),  // 5: google.protobuf.Timestamp
}
var file_com_coralogixapis_dataprime_background_queries_v1_request_proto_depIdxs = []int32{
	4, // 0: com.coralogixapis.dataprime.background_queries.v1.RunQueryRequest.query:type_name -> google.protobuf.StringValue
	0, // 1: com.coralogixapis.dataprime.background_queries.v1.RunQueryRequest.syntax:type_name -> com.coralogixapis.dataprime.background_queries.v1.QuerySyntax
	5, // 2: com.coralogixapis.dataprime.background_queries.v1.RunQueryRequest.start_date:type_name -> google.protobuf.Timestamp
	5, // 3: com.coralogixapis.dataprime.background_queries.v1.RunQueryRequest.end_date:type_name -> google.protobuf.Timestamp
	5, // 4: com.coralogixapis.dataprime.background_queries.v1.RunQueryRequest.now_date:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_dataprime_background_queries_v1_request_proto_init() }
func file_com_coralogixapis_dataprime_background_queries_v1_request_proto_init() {
	if File_com_coralogixapis_dataprime_background_queries_v1_request_proto != nil {
		return
	}
	file_com_coralogixapis_dataprime_background_queries_v1_request_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_dataprime_background_queries_v1_request_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_dataprime_background_queries_v1_request_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_dataprime_background_queries_v1_request_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_dataprime_background_queries_v1_request_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_dataprime_background_queries_v1_request_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_dataprime_background_queries_v1_request_proto = out.File
	file_com_coralogixapis_dataprime_background_queries_v1_request_proto_rawDesc = nil
	file_com_coralogixapis_dataprime_background_queries_v1_request_proto_goTypes = nil
	file_com_coralogixapis_dataprime_background_queries_v1_request_proto_depIdxs = nil
}
