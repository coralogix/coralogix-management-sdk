// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/spans2metrics/v2/spans_query.proto

package v2

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
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
	mi := &file_com_coralogixapis_spans2metrics_v2_spans_query_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SpansQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpansQuery) ProtoMessage() {}

func (x *SpansQuery) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_spans2metrics_v2_spans_query_proto_msgTypes[0]
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
	return file_com_coralogixapis_spans2metrics_v2_spans_query_proto_rawDescGZIP(), []int{0}
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

var File_com_coralogixapis_spans2metrics_v2_spans_query_proto protoreflect.FileDescriptor

const file_com_coralogixapis_spans2metrics_v2_spans_query_proto_rawDesc = "" +
	"\n" +
	"4com/coralogixapis/spans2metrics/v2/spans_query.proto\x12\"com.coralogixapis.spans2metrics.v2\x1a\x1egoogle/protobuf/wrappers.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\x99\x05\n" +
	"\n" +
	"SpansQuery\x12R\n" +
	"\x06lucene\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueB\x1c\x92A\x19J\x17\"applicationName:myApp\"R\x06lucene\x12c\n" +
	"\x17applicationname_filters\x18\x02 \x03(\v2\x1c.google.protobuf.StringValueB\f\x92A\tJ\a\"myApp\"R\x16applicationnameFilters\x12e\n" +
	"\x15subsystemname_filters\x18\x03 \x03(\v2\x1c.google.protobuf.StringValueB\x12\x92A\x0fJ\r\"mySubsystem\"R\x14subsystemnameFilters\x12T\n" +
	"\x0eaction_filters\x18\x04 \x03(\v2\x1c.google.protobuf.StringValueB\x0f\x92A\fJ\n" +
	"\"myAction\"R\ractionFilters\x12W\n" +
	"\x0fservice_filters\x18\x05 \x03(\v2\x1c.google.protobuf.StringValueB\x10\x92A\rJ\v\"myService\"R\x0eserviceFilters:\xbb\x01\x92A\xb7\x01\n" +
	"?*\n" +
	"SpansQuery21This data structure represents a query for spans.*t\n" +
	"\"Find out more about events2metrics\x12Nhttps://coralogix.com/docs/user-guides/monitoring-and-insights/events2metrics/b\x06proto3"

var (
	file_com_coralogixapis_spans2metrics_v2_spans_query_proto_rawDescOnce sync.Once
	file_com_coralogixapis_spans2metrics_v2_spans_query_proto_rawDescData []byte
)

func file_com_coralogixapis_spans2metrics_v2_spans_query_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_spans2metrics_v2_spans_query_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_spans2metrics_v2_spans_query_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_spans2metrics_v2_spans_query_proto_rawDesc), len(file_com_coralogixapis_spans2metrics_v2_spans_query_proto_rawDesc)))
	})
	return file_com_coralogixapis_spans2metrics_v2_spans_query_proto_rawDescData
}

var file_com_coralogixapis_spans2metrics_v2_spans_query_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_spans2metrics_v2_spans_query_proto_goTypes = []any{
	(*SpansQuery)(nil),             // 0: com.coralogixapis.spans2metrics.v2.SpansQuery
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
}
var file_com_coralogixapis_spans2metrics_v2_spans_query_proto_depIdxs = []int32{
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

func init() { file_com_coralogixapis_spans2metrics_v2_spans_query_proto_init() }
func file_com_coralogixapis_spans2metrics_v2_spans_query_proto_init() {
	if File_com_coralogixapis_spans2metrics_v2_spans_query_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_spans2metrics_v2_spans_query_proto_rawDesc), len(file_com_coralogixapis_spans2metrics_v2_spans_query_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_spans2metrics_v2_spans_query_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_spans2metrics_v2_spans_query_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_spans2metrics_v2_spans_query_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_spans2metrics_v2_spans_query_proto = out.File
	file_com_coralogixapis_spans2metrics_v2_spans_query_proto_goTypes = nil
	file_com_coralogixapis_spans2metrics_v2_spans_query_proto_depIdxs = nil
}
