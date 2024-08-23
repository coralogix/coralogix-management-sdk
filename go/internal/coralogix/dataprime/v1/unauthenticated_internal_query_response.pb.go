// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: com/coralogix/dataprime/v1/unauthenticated_internal_query_response.proto

package v1

import (
	v1 "github.com/coralogix/coralogix-management-sdk/internal/coralogixapis/dataprime/v1"
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

// dataprime response for text query
type UnauthenticatedInternalDataprimeQueryServiceQueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//
	//	*UnauthenticatedInternalDataprimeQueryServiceQueryResponse_Error
	//	*UnauthenticatedInternalDataprimeQueryServiceQueryResponse_Result
	//	*UnauthenticatedInternalDataprimeQueryServiceQueryResponse_Warning
	//	*UnauthenticatedInternalDataprimeQueryServiceQueryResponse_QueryId
	Message isUnauthenticatedInternalDataprimeQueryServiceQueryResponse_Message `protobuf_oneof:"message"`
}

func (x *UnauthenticatedInternalDataprimeQueryServiceQueryResponse) Reset() {
	*x = UnauthenticatedInternalDataprimeQueryServiceQueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_dataprime_v1_unauthenticated_internal_query_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnauthenticatedInternalDataprimeQueryServiceQueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnauthenticatedInternalDataprimeQueryServiceQueryResponse) ProtoMessage() {}

func (x *UnauthenticatedInternalDataprimeQueryServiceQueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_unauthenticated_internal_query_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnauthenticatedInternalDataprimeQueryServiceQueryResponse.ProtoReflect.Descriptor instead.
func (*UnauthenticatedInternalDataprimeQueryServiceQueryResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_unauthenticated_internal_query_response_proto_rawDescGZIP(), []int{0}
}

func (m *UnauthenticatedInternalDataprimeQueryServiceQueryResponse) GetMessage() isUnauthenticatedInternalDataprimeQueryServiceQueryResponse_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *UnauthenticatedInternalDataprimeQueryServiceQueryResponse) GetError() *v1.DataprimeError {
	if x, ok := x.GetMessage().(*UnauthenticatedInternalDataprimeQueryServiceQueryResponse_Error); ok {
		return x.Error
	}
	return nil
}

func (x *UnauthenticatedInternalDataprimeQueryServiceQueryResponse) GetResult() *v1.DataprimeResult {
	if x, ok := x.GetMessage().(*UnauthenticatedInternalDataprimeQueryServiceQueryResponse_Result); ok {
		return x.Result
	}
	return nil
}

func (x *UnauthenticatedInternalDataprimeQueryServiceQueryResponse) GetWarning() *v1.DataprimeWarning {
	if x, ok := x.GetMessage().(*UnauthenticatedInternalDataprimeQueryServiceQueryResponse_Warning); ok {
		return x.Warning
	}
	return nil
}

func (x *UnauthenticatedInternalDataprimeQueryServiceQueryResponse) GetQueryId() *v1.QueryId {
	if x, ok := x.GetMessage().(*UnauthenticatedInternalDataprimeQueryServiceQueryResponse_QueryId); ok {
		return x.QueryId
	}
	return nil
}

type isUnauthenticatedInternalDataprimeQueryServiceQueryResponse_Message interface {
	isUnauthenticatedInternalDataprimeQueryServiceQueryResponse_Message()
}

type UnauthenticatedInternalDataprimeQueryServiceQueryResponse_Error struct {
	Error *v1.DataprimeError `protobuf:"bytes,1,opt,name=error,proto3,oneof"`
}

type UnauthenticatedInternalDataprimeQueryServiceQueryResponse_Result struct {
	Result *v1.DataprimeResult `protobuf:"bytes,2,opt,name=result,proto3,oneof"`
}

type UnauthenticatedInternalDataprimeQueryServiceQueryResponse_Warning struct {
	Warning *v1.DataprimeWarning `protobuf:"bytes,3,opt,name=warning,proto3,oneof"`
}

type UnauthenticatedInternalDataprimeQueryServiceQueryResponse_QueryId struct {
	QueryId *v1.QueryId `protobuf:"bytes,4,opt,name=query_id,json=queryId,proto3,oneof"`
}

func (*UnauthenticatedInternalDataprimeQueryServiceQueryResponse_Error) isUnauthenticatedInternalDataprimeQueryServiceQueryResponse_Message() {
}

func (*UnauthenticatedInternalDataprimeQueryServiceQueryResponse_Result) isUnauthenticatedInternalDataprimeQueryServiceQueryResponse_Message() {
}

func (*UnauthenticatedInternalDataprimeQueryServiceQueryResponse_Warning) isUnauthenticatedInternalDataprimeQueryServiceQueryResponse_Message() {
}

func (*UnauthenticatedInternalDataprimeQueryServiceQueryResponse_QueryId) isUnauthenticatedInternalDataprimeQueryServiceQueryResponse_Message() {
}

var File_com_coralogix_dataprime_v1_unauthenticated_internal_query_response_proto protoreflect.FileDescriptor

var file_com_coralogix_dataprime_v1_unauthenticated_internal_query_response_proto_rawDesc = []byte{
	0x0a, 0x48, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x6e, 0x61,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x2d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed, 0x02, 0x0a, 0x39, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x49, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x4c, 0x0a, 0x07, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69,
	0x6d, 0x65, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x07, 0x77, 0x61, 0x72,
	0x6e, 0x69, 0x6e, 0x67, 0x12, 0x44, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70,
	0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x64, 0x48,
	0x00, 0x52, 0x07, 0x71, 0x75, 0x65, 0x72, 0x79, 0x49, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_dataprime_v1_unauthenticated_internal_query_response_proto_rawDescOnce sync.Once
	file_com_coralogix_dataprime_v1_unauthenticated_internal_query_response_proto_rawDescData = file_com_coralogix_dataprime_v1_unauthenticated_internal_query_response_proto_rawDesc
)

func file_com_coralogix_dataprime_v1_unauthenticated_internal_query_response_proto_rawDescGZIP() []byte {
	file_com_coralogix_dataprime_v1_unauthenticated_internal_query_response_proto_rawDescOnce.Do(func() {
		file_com_coralogix_dataprime_v1_unauthenticated_internal_query_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_dataprime_v1_unauthenticated_internal_query_response_proto_rawDescData)
	})
	return file_com_coralogix_dataprime_v1_unauthenticated_internal_query_response_proto_rawDescData
}

var file_com_coralogix_dataprime_v1_unauthenticated_internal_query_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogix_dataprime_v1_unauthenticated_internal_query_response_proto_goTypes = []any{
	(*UnauthenticatedInternalDataprimeQueryServiceQueryResponse)(nil), // 0: com.coralogix.dataprime.v1.UnauthenticatedInternalDataprimeQueryServiceQueryResponse
	(*v1.DataprimeError)(nil),   // 1: com.coralogixapis.dataprime.v1.DataprimeError
	(*v1.DataprimeResult)(nil),  // 2: com.coralogixapis.dataprime.v1.DataprimeResult
	(*v1.DataprimeWarning)(nil), // 3: com.coralogixapis.dataprime.v1.DataprimeWarning
	(*v1.QueryId)(nil),          // 4: com.coralogixapis.dataprime.v1.QueryId
}
var file_com_coralogix_dataprime_v1_unauthenticated_internal_query_response_proto_depIdxs = []int32{
	1, // 0: com.coralogix.dataprime.v1.UnauthenticatedInternalDataprimeQueryServiceQueryResponse.error:type_name -> com.coralogixapis.dataprime.v1.DataprimeError
	2, // 1: com.coralogix.dataprime.v1.UnauthenticatedInternalDataprimeQueryServiceQueryResponse.result:type_name -> com.coralogixapis.dataprime.v1.DataprimeResult
	3, // 2: com.coralogix.dataprime.v1.UnauthenticatedInternalDataprimeQueryServiceQueryResponse.warning:type_name -> com.coralogixapis.dataprime.v1.DataprimeWarning
	4, // 3: com.coralogix.dataprime.v1.UnauthenticatedInternalDataprimeQueryServiceQueryResponse.query_id:type_name -> com.coralogixapis.dataprime.v1.QueryId
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_com_coralogix_dataprime_v1_unauthenticated_internal_query_response_proto_init() }
func file_com_coralogix_dataprime_v1_unauthenticated_internal_query_response_proto_init() {
	if File_com_coralogix_dataprime_v1_unauthenticated_internal_query_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_dataprime_v1_unauthenticated_internal_query_response_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UnauthenticatedInternalDataprimeQueryServiceQueryResponse); i {
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
	file_com_coralogix_dataprime_v1_unauthenticated_internal_query_response_proto_msgTypes[0].OneofWrappers = []any{
		(*UnauthenticatedInternalDataprimeQueryServiceQueryResponse_Error)(nil),
		(*UnauthenticatedInternalDataprimeQueryServiceQueryResponse_Result)(nil),
		(*UnauthenticatedInternalDataprimeQueryServiceQueryResponse_Warning)(nil),
		(*UnauthenticatedInternalDataprimeQueryServiceQueryResponse_QueryId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_dataprime_v1_unauthenticated_internal_query_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_dataprime_v1_unauthenticated_internal_query_response_proto_goTypes,
		DependencyIndexes: file_com_coralogix_dataprime_v1_unauthenticated_internal_query_response_proto_depIdxs,
		MessageInfos:      file_com_coralogix_dataprime_v1_unauthenticated_internal_query_response_proto_msgTypes,
	}.Build()
	File_com_coralogix_dataprime_v1_unauthenticated_internal_query_response_proto = out.File
	file_com_coralogix_dataprime_v1_unauthenticated_internal_query_response_proto_rawDesc = nil
	file_com_coralogix_dataprime_v1_unauthenticated_internal_query_response_proto_goTypes = nil
	file_com_coralogix_dataprime_v1_unauthenticated_internal_query_response_proto_depIdxs = nil
}
