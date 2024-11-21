// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: com/coralogixapis/apm_shared/common/queries.proto

package common

import (
	_ "github.com/golang/protobuf/ptypes/struct"
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

type GridQueryInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query  *wrappers.StringValue `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Column *wrappers.StringValue `protobuf:"bytes,2,opt,name=column,proto3" json:"column,omitempty"`
}

func (x *GridQueryInfo) Reset() {
	*x = GridQueryInfo{}
	mi := &file_com_coralogixapis_apm_shared_common_queries_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GridQueryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GridQueryInfo) ProtoMessage() {}

func (x *GridQueryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_shared_common_queries_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GridQueryInfo.ProtoReflect.Descriptor instead.
func (*GridQueryInfo) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_shared_common_queries_proto_rawDescGZIP(), []int{0}
}

func (x *GridQueryInfo) GetQuery() *wrappers.StringValue {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *GridQueryInfo) GetColumn() *wrappers.StringValue {
	if x != nil {
		return x.Column
	}
	return nil
}

type QueryInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query *wrappers.StringValue `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Name  *wrappers.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *QueryInfo) Reset() {
	*x = QueryInfo{}
	mi := &file_com_coralogixapis_apm_shared_common_queries_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryInfo) ProtoMessage() {}

func (x *QueryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_shared_common_queries_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryInfo.ProtoReflect.Descriptor instead.
func (*QueryInfo) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_shared_common_queries_proto_rawDescGZIP(), []int{1}
}

func (x *QueryInfo) GetQuery() *wrappers.StringValue {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *QueryInfo) GetName() *wrappers.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

var File_com_coralogixapis_apm_shared_common_queries_proto protoreflect.FileDescriptor

var file_com_coralogixapis_apm_shared_common_queries_proto_rawDesc = []byte{
	0x0a, 0x31, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x32, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x79, 0x0a, 0x0d, 0x47, 0x72, 0x69, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x32, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x34, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x22, 0x71, 0x0a,
	0x09, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x32, 0x0a, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x30,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_apm_shared_common_queries_proto_rawDescOnce sync.Once
	file_com_coralogixapis_apm_shared_common_queries_proto_rawDescData = file_com_coralogixapis_apm_shared_common_queries_proto_rawDesc
)

func file_com_coralogixapis_apm_shared_common_queries_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_apm_shared_common_queries_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_apm_shared_common_queries_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_apm_shared_common_queries_proto_rawDescData)
	})
	return file_com_coralogixapis_apm_shared_common_queries_proto_rawDescData
}

var file_com_coralogixapis_apm_shared_common_queries_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogixapis_apm_shared_common_queries_proto_goTypes = []any{
	(*GridQueryInfo)(nil),        // 0: com.coralogixapis.apm.common.v2.GridQueryInfo
	(*QueryInfo)(nil),            // 1: com.coralogixapis.apm.common.v2.QueryInfo
	(*wrappers.StringValue)(nil), // 2: google.protobuf.StringValue
}
var file_com_coralogixapis_apm_shared_common_queries_proto_depIdxs = []int32{
	2, // 0: com.coralogixapis.apm.common.v2.GridQueryInfo.query:type_name -> google.protobuf.StringValue
	2, // 1: com.coralogixapis.apm.common.v2.GridQueryInfo.column:type_name -> google.protobuf.StringValue
	2, // 2: com.coralogixapis.apm.common.v2.QueryInfo.query:type_name -> google.protobuf.StringValue
	2, // 3: com.coralogixapis.apm.common.v2.QueryInfo.name:type_name -> google.protobuf.StringValue
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_apm_shared_common_queries_proto_init() }
func file_com_coralogixapis_apm_shared_common_queries_proto_init() {
	if File_com_coralogixapis_apm_shared_common_queries_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_apm_shared_common_queries_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_apm_shared_common_queries_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_apm_shared_common_queries_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_apm_shared_common_queries_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_apm_shared_common_queries_proto = out.File
	file_com_coralogixapis_apm_shared_common_queries_proto_rawDesc = nil
	file_com_coralogixapis_apm_shared_common_queries_proto_goTypes = nil
	file_com_coralogixapis_apm_shared_common_queries_proto_depIdxs = nil
}
