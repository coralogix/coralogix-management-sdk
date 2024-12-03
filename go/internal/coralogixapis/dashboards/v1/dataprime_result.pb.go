// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: com/coralogixapis/dashboards/v1/common/dataprime_result.proto

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

type DataprimeResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata []*DataprimeResult_KeyValue `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty"`
	Labels   []*DataprimeResult_KeyValue `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"`
	UserData string                      `protobuf:"bytes,3,opt,name=user_data,json=userData,proto3" json:"user_data,omitempty"`
}

func (x *DataprimeResult) Reset() {
	*x = DataprimeResult{}
	mi := &file_com_coralogixapis_dashboards_v1_common_dataprime_result_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataprimeResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataprimeResult) ProtoMessage() {}

func (x *DataprimeResult) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_common_dataprime_result_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataprimeResult.ProtoReflect.Descriptor instead.
func (*DataprimeResult) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_dataprime_result_proto_rawDescGZIP(), []int{0}
}

func (x *DataprimeResult) GetMetadata() []*DataprimeResult_KeyValue {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *DataprimeResult) GetLabels() []*DataprimeResult_KeyValue {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *DataprimeResult) GetUserData() string {
	if x != nil {
		return x.UserData
	}
	return ""
}

type DataprimeResult_KeyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *DataprimeResult_KeyValue) Reset() {
	*x = DataprimeResult_KeyValue{}
	mi := &file_com_coralogixapis_dashboards_v1_common_dataprime_result_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataprimeResult_KeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataprimeResult_KeyValue) ProtoMessage() {}

func (x *DataprimeResult_KeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_common_dataprime_result_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataprimeResult_KeyValue.ProtoReflect.Descriptor instead.
func (*DataprimeResult_KeyValue) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_dataprime_result_proto_rawDescGZIP(), []int{0, 0}
}

func (x *DataprimeResult_KeyValue) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *DataprimeResult_KeyValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_com_coralogixapis_dashboards_v1_common_dataprime_result_proto protoreflect.FileDescriptor

var file_com_coralogixapis_dashboards_v1_common_dataprime_result_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69,
	0x6d, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x9a, 0x02, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61,
	0x70, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x5c, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x40, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x58, 0x0a, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x1a, 0x32, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_dashboards_v1_common_dataprime_result_proto_rawDescOnce sync.Once
	file_com_coralogixapis_dashboards_v1_common_dataprime_result_proto_rawDescData = file_com_coralogixapis_dashboards_v1_common_dataprime_result_proto_rawDesc
)

func file_com_coralogixapis_dashboards_v1_common_dataprime_result_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_dashboards_v1_common_dataprime_result_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_dashboards_v1_common_dataprime_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_dashboards_v1_common_dataprime_result_proto_rawDescData)
	})
	return file_com_coralogixapis_dashboards_v1_common_dataprime_result_proto_rawDescData
}

var file_com_coralogixapis_dashboards_v1_common_dataprime_result_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogixapis_dashboards_v1_common_dataprime_result_proto_goTypes = []any{
	(*DataprimeResult)(nil),          // 0: com.coralogixapis.dashboards.v1.common.DataprimeResult
	(*DataprimeResult_KeyValue)(nil), // 1: com.coralogixapis.dashboards.v1.common.DataprimeResult.KeyValue
}
var file_com_coralogixapis_dashboards_v1_common_dataprime_result_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.dashboards.v1.common.DataprimeResult.metadata:type_name -> com.coralogixapis.dashboards.v1.common.DataprimeResult.KeyValue
	1, // 1: com.coralogixapis.dashboards.v1.common.DataprimeResult.labels:type_name -> com.coralogixapis.dashboards.v1.common.DataprimeResult.KeyValue
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_dashboards_v1_common_dataprime_result_proto_init() }
func file_com_coralogixapis_dashboards_v1_common_dataprime_result_proto_init() {
	if File_com_coralogixapis_dashboards_v1_common_dataprime_result_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_dashboards_v1_common_dataprime_result_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_dashboards_v1_common_dataprime_result_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_dashboards_v1_common_dataprime_result_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_dashboards_v1_common_dataprime_result_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_dashboards_v1_common_dataprime_result_proto = out.File
	file_com_coralogixapis_dashboards_v1_common_dataprime_result_proto_rawDesc = nil
	file_com_coralogixapis_dashboards_v1_common_dataprime_result_proto_goTypes = nil
	file_com_coralogixapis_dashboards_v1_common_dataprime_result_proto_depIdxs = nil
}
