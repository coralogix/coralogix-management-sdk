// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/apm/queries/k8s/v1/cluster_filters.proto

package v1

import (
	v2 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/apm/common/v2"
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

type ClusterFilters struct {
	state               protoimpl.MessageState  `protogen:"open.v1"`
	SelectedClusterName *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=selected_cluster_name,json=selectedClusterName,proto3" json:"selected_cluster_name,omitempty"`
	TimeRange           *v2.TimeRange           `protobuf:"bytes,2,opt,name=time_range,json=timeRange,proto3" json:"time_range,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *ClusterFilters) Reset() {
	*x = ClusterFilters{}
	mi := &file_com_coralogixapis_apm_queries_k8s_v1_cluster_filters_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClusterFilters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterFilters) ProtoMessage() {}

func (x *ClusterFilters) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_queries_k8s_v1_cluster_filters_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterFilters.ProtoReflect.Descriptor instead.
func (*ClusterFilters) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_queries_k8s_v1_cluster_filters_proto_rawDescGZIP(), []int{0}
}

func (x *ClusterFilters) GetSelectedClusterName() *wrapperspb.StringValue {
	if x != nil {
		return x.SelectedClusterName
	}
	return nil
}

func (x *ClusterFilters) GetTimeRange() *v2.TimeRange {
	if x != nil {
		return x.TimeRange
	}
	return nil
}

var File_com_coralogixapis_apm_queries_k8s_v1_cluster_filters_proto protoreflect.FileDescriptor

const file_com_coralogixapis_apm_queries_k8s_v1_cluster_filters_proto_rawDesc = "" +
	"\n" +
	":com/coralogixapis/apm/queries/k8s/v1/cluster_filters.proto\x12$com.coralogixapis.apm.queries.k8s.v1\x1a\x1egoogle/protobuf/wrappers.proto\x1a0com/coralogixapis/apm/common/v2/time_range.proto\"\xad\x01\n" +
	"\x0eClusterFilters\x12P\n" +
	"\x15selected_cluster_name\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueR\x13selectedClusterName\x12I\n" +
	"\n" +
	"time_range\x18\x02 \x01(\v2*.com.coralogixapis.apm.common.v2.TimeRangeR\ttimeRangeb\x06proto3"

var (
	file_com_coralogixapis_apm_queries_k8s_v1_cluster_filters_proto_rawDescOnce sync.Once
	file_com_coralogixapis_apm_queries_k8s_v1_cluster_filters_proto_rawDescData []byte
)

func file_com_coralogixapis_apm_queries_k8s_v1_cluster_filters_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_apm_queries_k8s_v1_cluster_filters_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_apm_queries_k8s_v1_cluster_filters_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_apm_queries_k8s_v1_cluster_filters_proto_rawDesc), len(file_com_coralogixapis_apm_queries_k8s_v1_cluster_filters_proto_rawDesc)))
	})
	return file_com_coralogixapis_apm_queries_k8s_v1_cluster_filters_proto_rawDescData
}

var file_com_coralogixapis_apm_queries_k8s_v1_cluster_filters_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_apm_queries_k8s_v1_cluster_filters_proto_goTypes = []any{
	(*ClusterFilters)(nil),         // 0: com.coralogixapis.apm.queries.k8s.v1.ClusterFilters
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
	(*v2.TimeRange)(nil),           // 2: com.coralogixapis.apm.common.v2.TimeRange
}
var file_com_coralogixapis_apm_queries_k8s_v1_cluster_filters_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.apm.queries.k8s.v1.ClusterFilters.selected_cluster_name:type_name -> google.protobuf.StringValue
	2, // 1: com.coralogixapis.apm.queries.k8s.v1.ClusterFilters.time_range:type_name -> com.coralogixapis.apm.common.v2.TimeRange
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_apm_queries_k8s_v1_cluster_filters_proto_init() }
func file_com_coralogixapis_apm_queries_k8s_v1_cluster_filters_proto_init() {
	if File_com_coralogixapis_apm_queries_k8s_v1_cluster_filters_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_apm_queries_k8s_v1_cluster_filters_proto_rawDesc), len(file_com_coralogixapis_apm_queries_k8s_v1_cluster_filters_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_apm_queries_k8s_v1_cluster_filters_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_apm_queries_k8s_v1_cluster_filters_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_apm_queries_k8s_v1_cluster_filters_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_apm_queries_k8s_v1_cluster_filters_proto = out.File
	file_com_coralogixapis_apm_queries_k8s_v1_cluster_filters_proto_goTypes = nil
	file_com_coralogixapis_apm_queries_k8s_v1_cluster_filters_proto_depIdxs = nil
}
