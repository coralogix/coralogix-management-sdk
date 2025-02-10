// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: src/com/coralogixapis/spans2metrics/v2/spans_query.proto

package v2

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type SpansQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// lucene query
	Lucene *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=lucene,proto3" json:"lucene,omitempty"`
	// application name filters
	ApplicationnameFilters []*wrapperspb.StringValue `protobuf:"bytes,2,rep,name=applicationname_filters,json=applicationnameFilters,proto3" json:"applicationname_filters,omitempty"`
	// subsystem name filters
	SubsystemnameFilters []*wrapperspb.StringValue `protobuf:"bytes,3,rep,name=subsystemname_filters,json=subsystemnameFilters,proto3" json:"subsystemname_filters,omitempty"`
	// action filters
	ActionFilters []*wrapperspb.StringValue `protobuf:"bytes,4,rep,name=action_filters,json=actionFilters,proto3" json:"action_filters,omitempty"`
	// service filters
	ServiceFilters []*wrapperspb.StringValue `protobuf:"bytes,5,rep,name=service_filters,json=serviceFilters,proto3" json:"service_filters,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *SpansQuery) Reset() {
	*x = SpansQuery{}
	mi := &file_src_com_coralogixapis_spans2metrics_v2_spans_query_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SpansQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpansQuery) ProtoMessage() {}

func (x *SpansQuery) ProtoReflect() protoreflect.Message {
	mi := &file_src_com_coralogixapis_spans2metrics_v2_spans_query_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpansQuery.ProtoReflect.Descriptor instead.
func (*SpansQuery) Descriptor() ([]byte, []int) {
	return file_src_com_coralogixapis_spans2metrics_v2_spans_query_proto_rawDescGZIP(), []int{0}
}

func (x *SpansQuery) GetLucene() *wrapperspb.StringValue {
	if x != nil {
		return x.Lucene
	}
	return nil
}

func (x *SpansQuery) GetApplicationnameFilters() []*wrapperspb.StringValue {
	if x != nil {
		return x.ApplicationnameFilters
	}
	return nil
}

func (x *SpansQuery) GetSubsystemnameFilters() []*wrapperspb.StringValue {
	if x != nil {
		return x.SubsystemnameFilters
	}
	return nil
}

func (x *SpansQuery) GetActionFilters() []*wrapperspb.StringValue {
	if x != nil {
		return x.ActionFilters
	}
	return nil
}

func (x *SpansQuery) GetServiceFilters() []*wrapperspb.StringValue {
	if x != nil {
		return x.ServiceFilters
	}
	return nil
}

var File_src_com_coralogixapis_spans2metrics_v2_spans_query_proto protoreflect.FileDescriptor

var file_src_com_coralogixapis_spans2metrics_v2_spans_query_proto_rawDesc = []byte{
	0x0a, 0x38, 0x73, 0x72, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x73, 0x32, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x73, 0x5f, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x70,
	0x61, 0x6e, 0x73, 0x32, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99,
	0x05, 0x0a, 0x0a, 0x53, 0x70, 0x61, 0x6e, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x52, 0x0a,
	0x06, 0x6c, 0x75, 0x63, 0x65, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x1c, 0x92, 0x41, 0x19,
	0x4a, 0x17, 0x22, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x3a, 0x6d, 0x79, 0x41, 0x70, 0x70, 0x22, 0x52, 0x06, 0x6c, 0x75, 0x63, 0x65, 0x6e,
	0x65, 0x12, 0x63, 0x0a, 0x17, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x0c, 0x92, 0x41, 0x09, 0x4a, 0x07, 0x22, 0x6d, 0x79, 0x41, 0x70, 0x70, 0x22, 0x52, 0x16,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x6e, 0x61, 0x6d, 0x65, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x65, 0x0a, 0x15, 0x73, 0x75, 0x62, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x12, 0x92, 0x41, 0x0f, 0x4a, 0x0d, 0x22, 0x6d, 0x79, 0x53, 0x75, 0x62,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x22, 0x52, 0x14, 0x73, 0x75, 0x62, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x6e, 0x61, 0x6d, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x54, 0x0a,
	0x0e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x0f, 0x92, 0x41, 0x0c, 0x4a, 0x0a, 0x22, 0x6d, 0x79, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x52, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x12, 0x57, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x10, 0x92, 0x41, 0x0d, 0x4a,
	0x0b, 0x22, 0x6d, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x52, 0x0e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x3a, 0xbb, 0x01, 0x92,
	0x41, 0xb7, 0x01, 0x0a, 0x3f, 0x2a, 0x0a, 0x53, 0x70, 0x61, 0x6e, 0x73, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x32, 0x31, 0x54, 0x68, 0x69, 0x73, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x20, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74,
	0x73, 0x20, 0x61, 0x20, 0x71, 0x75, 0x65, 0x72, 0x79, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x73, 0x70,
	0x61, 0x6e, 0x73, 0x2e, 0x2a, 0x74, 0x0a, 0x22, 0x46, 0x69, 0x6e, 0x64, 0x20, 0x6f, 0x75, 0x74,
	0x20, 0x6d, 0x6f, 0x72, 0x65, 0x20, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x20, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x32, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x4e, 0x68, 0x74, 0x74, 0x70,
	0x73, 0x3a, 0x2f, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x6f, 0x63, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x67, 0x75, 0x69, 0x64,
	0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2d, 0x61, 0x6e,
	0x64, 0x2d, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x32, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_src_com_coralogixapis_spans2metrics_v2_spans_query_proto_rawDescOnce sync.Once
	file_src_com_coralogixapis_spans2metrics_v2_spans_query_proto_rawDescData = file_src_com_coralogixapis_spans2metrics_v2_spans_query_proto_rawDesc
)

func file_src_com_coralogixapis_spans2metrics_v2_spans_query_proto_rawDescGZIP() []byte {
	file_src_com_coralogixapis_spans2metrics_v2_spans_query_proto_rawDescOnce.Do(func() {
		file_src_com_coralogixapis_spans2metrics_v2_spans_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_com_coralogixapis_spans2metrics_v2_spans_query_proto_rawDescData)
	})
	return file_src_com_coralogixapis_spans2metrics_v2_spans_query_proto_rawDescData
}

var file_src_com_coralogixapis_spans2metrics_v2_spans_query_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_src_com_coralogixapis_spans2metrics_v2_spans_query_proto_goTypes = []any{
	(*SpansQuery)(nil),             // 0: com.coralogixapis.spans2metrics.v2.SpansQuery
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
}
var file_src_com_coralogixapis_spans2metrics_v2_spans_query_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.spans2metrics.v2.SpansQuery.lucene:type_name -> google.protobuf.StringValue
	1, // 1: com.coralogixapis.spans2metrics.v2.SpansQuery.applicationname_filters:type_name -> google.protobuf.StringValue
	1, // 2: com.coralogixapis.spans2metrics.v2.SpansQuery.subsystemname_filters:type_name -> google.protobuf.StringValue
	1, // 3: com.coralogixapis.spans2metrics.v2.SpansQuery.action_filters:type_name -> google.protobuf.StringValue
	1, // 4: com.coralogixapis.spans2metrics.v2.SpansQuery.service_filters:type_name -> google.protobuf.StringValue
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_src_com_coralogixapis_spans2metrics_v2_spans_query_proto_init() }
func file_src_com_coralogixapis_spans2metrics_v2_spans_query_proto_init() {
	if File_src_com_coralogixapis_spans2metrics_v2_spans_query_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_src_com_coralogixapis_spans2metrics_v2_spans_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_com_coralogixapis_spans2metrics_v2_spans_query_proto_goTypes,
		DependencyIndexes: file_src_com_coralogixapis_spans2metrics_v2_spans_query_proto_depIdxs,
		MessageInfos:      file_src_com_coralogixapis_spans2metrics_v2_spans_query_proto_msgTypes,
	}.Build()
	File_src_com_coralogixapis_spans2metrics_v2_spans_query_proto = out.File
	file_src_com_coralogixapis_spans2metrics_v2_spans_query_proto_rawDesc = nil
	file_src_com_coralogixapis_spans2metrics_v2_spans_query_proto_goTypes = nil
	file_src_com_coralogixapis_spans2metrics_v2_spans_query_proto_depIdxs = nil
}
