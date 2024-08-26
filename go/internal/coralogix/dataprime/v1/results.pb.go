// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.3
// source: com/coralogix/dataprime/v1/results.proto

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

type GetQueryResultsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueryId string `protobuf:"bytes,1,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
}

func (x *GetQueryResultsRequest) Reset() {
	*x = GetQueryResultsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_dataprime_v1_results_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQueryResultsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQueryResultsRequest) ProtoMessage() {}

func (x *GetQueryResultsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_results_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQueryResultsRequest.ProtoReflect.Descriptor instead.
func (*GetQueryResultsRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_results_proto_rawDescGZIP(), []int{0}
}

func (x *GetQueryResultsRequest) GetQueryId() string {
	if x != nil {
		return x.QueryId
	}
	return ""
}

type GetQueryResultsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata []*GetQueryResultsResponse_KeyValue `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty"`
	Labels   []*GetQueryResultsResponse_KeyValue `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"`
	UserData string                              `protobuf:"bytes,3,opt,name=user_data,json=userData,proto3" json:"user_data,omitempty"`
}

func (x *GetQueryResultsResponse) Reset() {
	*x = GetQueryResultsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_dataprime_v1_results_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQueryResultsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQueryResultsResponse) ProtoMessage() {}

func (x *GetQueryResultsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_results_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQueryResultsResponse.ProtoReflect.Descriptor instead.
func (*GetQueryResultsResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_results_proto_rawDescGZIP(), []int{1}
}

func (x *GetQueryResultsResponse) GetMetadata() []*GetQueryResultsResponse_KeyValue {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *GetQueryResultsResponse) GetLabels() []*GetQueryResultsResponse_KeyValue {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *GetQueryResultsResponse) GetUserData() string {
	if x != nil {
		return x.UserData
	}
	return ""
}

type GetQueryResultsResponse_KeyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetQueryResultsResponse_KeyValue) Reset() {
	*x = GetQueryResultsResponse_KeyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_dataprime_v1_results_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQueryResultsResponse_KeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQueryResultsResponse_KeyValue) ProtoMessage() {}

func (x *GetQueryResultsResponse_KeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_results_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQueryResultsResponse_KeyValue.ProtoReflect.Descriptor instead.
func (*GetQueryResultsResponse_KeyValue) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_results_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetQueryResultsResponse_KeyValue) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *GetQueryResultsResponse_KeyValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_com_coralogix_dataprime_v1_results_proto protoreflect.FileDescriptor

var file_com_coralogix_dataprime_v1_results_proto_rawDesc = []byte{
	0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x33, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x71, 0x75, 0x65, 0x72, 0x79, 0x49, 0x64, 0x22, 0x9a, 0x02, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4b,
	0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x54, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x3c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x1a, 0x32, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_dataprime_v1_results_proto_rawDescOnce sync.Once
	file_com_coralogix_dataprime_v1_results_proto_rawDescData = file_com_coralogix_dataprime_v1_results_proto_rawDesc
)

func file_com_coralogix_dataprime_v1_results_proto_rawDescGZIP() []byte {
	file_com_coralogix_dataprime_v1_results_proto_rawDescOnce.Do(func() {
		file_com_coralogix_dataprime_v1_results_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_dataprime_v1_results_proto_rawDescData)
	})
	return file_com_coralogix_dataprime_v1_results_proto_rawDescData
}

var file_com_coralogix_dataprime_v1_results_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogix_dataprime_v1_results_proto_goTypes = []interface{}{
	(*GetQueryResultsRequest)(nil),           // 0: com.coralogix.dataprime.v1.GetQueryResultsRequest
	(*GetQueryResultsResponse)(nil),          // 1: com.coralogix.dataprime.v1.GetQueryResultsResponse
	(*GetQueryResultsResponse_KeyValue)(nil), // 2: com.coralogix.dataprime.v1.GetQueryResultsResponse.KeyValue
}
var file_com_coralogix_dataprime_v1_results_proto_depIdxs = []int32{
	2, // 0: com.coralogix.dataprime.v1.GetQueryResultsResponse.metadata:type_name -> com.coralogix.dataprime.v1.GetQueryResultsResponse.KeyValue
	2, // 1: com.coralogix.dataprime.v1.GetQueryResultsResponse.labels:type_name -> com.coralogix.dataprime.v1.GetQueryResultsResponse.KeyValue
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_com_coralogix_dataprime_v1_results_proto_init() }
func file_com_coralogix_dataprime_v1_results_proto_init() {
	if File_com_coralogix_dataprime_v1_results_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_dataprime_v1_results_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQueryResultsRequest); i {
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
		file_com_coralogix_dataprime_v1_results_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQueryResultsResponse); i {
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
		file_com_coralogix_dataprime_v1_results_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQueryResultsResponse_KeyValue); i {
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
			RawDescriptor: file_com_coralogix_dataprime_v1_results_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_dataprime_v1_results_proto_goTypes,
		DependencyIndexes: file_com_coralogix_dataprime_v1_results_proto_depIdxs,
		MessageInfos:      file_com_coralogix_dataprime_v1_results_proto_msgTypes,
	}.Build()
	File_com_coralogix_dataprime_v1_results_proto = out.File
	file_com_coralogix_dataprime_v1_results_proto_rawDesc = nil
	file_com_coralogix_dataprime_v1_results_proto_goTypes = nil
	file_com_coralogix_dataprime_v1_results_proto_depIdxs = nil
}
