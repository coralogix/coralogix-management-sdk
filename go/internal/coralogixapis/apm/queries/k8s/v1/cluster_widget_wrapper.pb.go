// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: com/coralogixapis/apm/queries/k8s/v1/cluster_widget_wrapper.proto

package v1

import (
	v1 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/apm/widgets/v1"
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

type ClusterWidgetWrapper struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Widget:
	//
	//	*ClusterWidgetWrapper_Donut
	//	*ClusterWidgetWrapper_Gauge
	//	*ClusterWidgetWrapper_LineChart
	//	*ClusterWidgetWrapper_MapChart
	//	*ClusterWidgetWrapper_Topk
	//	*ClusterWidgetWrapper_TopkMultiValue
	Widget        isClusterWidgetWrapper_Widget `protobuf_oneof:"widget"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClusterWidgetWrapper) Reset() {
	*x = ClusterWidgetWrapper{}
	mi := &file_com_coralogixapis_apm_queries_k8s_v1_cluster_widget_wrapper_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClusterWidgetWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterWidgetWrapper) ProtoMessage() {}

func (x *ClusterWidgetWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_queries_k8s_v1_cluster_widget_wrapper_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterWidgetWrapper.ProtoReflect.Descriptor instead.
func (*ClusterWidgetWrapper) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_queries_k8s_v1_cluster_widget_wrapper_proto_rawDescGZIP(), []int{0}
}

func (x *ClusterWidgetWrapper) GetWidget() isClusterWidgetWrapper_Widget {
	if x != nil {
		return x.Widget
	}
	return nil
}

func (x *ClusterWidgetWrapper) GetDonut() *v1.Donut {
	if x != nil {
		if x, ok := x.Widget.(*ClusterWidgetWrapper_Donut); ok {
			return x.Donut
		}
	}
	return nil
}

func (x *ClusterWidgetWrapper) GetGauge() *v1.Gauge {
	if x != nil {
		if x, ok := x.Widget.(*ClusterWidgetWrapper_Gauge); ok {
			return x.Gauge
		}
	}
	return nil
}

func (x *ClusterWidgetWrapper) GetLineChart() *v1.LineChart {
	if x != nil {
		if x, ok := x.Widget.(*ClusterWidgetWrapper_LineChart); ok {
			return x.LineChart
		}
	}
	return nil
}

func (x *ClusterWidgetWrapper) GetMapChart() *v1.MapChart {
	if x != nil {
		if x, ok := x.Widget.(*ClusterWidgetWrapper_MapChart); ok {
			return x.MapChart
		}
	}
	return nil
}

func (x *ClusterWidgetWrapper) GetTopk() *v1.Topk {
	if x != nil {
		if x, ok := x.Widget.(*ClusterWidgetWrapper_Topk); ok {
			return x.Topk
		}
	}
	return nil
}

func (x *ClusterWidgetWrapper) GetTopkMultiValue() *v1.TopkMultiValue {
	if x != nil {
		if x, ok := x.Widget.(*ClusterWidgetWrapper_TopkMultiValue); ok {
			return x.TopkMultiValue
		}
	}
	return nil
}

type isClusterWidgetWrapper_Widget interface {
	isClusterWidgetWrapper_Widget()
}

type ClusterWidgetWrapper_Donut struct {
	Donut *v1.Donut `protobuf:"bytes,1,opt,name=donut,proto3,oneof"`
}

type ClusterWidgetWrapper_Gauge struct {
	Gauge *v1.Gauge `protobuf:"bytes,2,opt,name=gauge,proto3,oneof"`
}

type ClusterWidgetWrapper_LineChart struct {
	LineChart *v1.LineChart `protobuf:"bytes,3,opt,name=line_chart,json=lineChart,proto3,oneof"`
}

type ClusterWidgetWrapper_MapChart struct {
	MapChart *v1.MapChart `protobuf:"bytes,4,opt,name=map_chart,json=mapChart,proto3,oneof"`
}

type ClusterWidgetWrapper_Topk struct {
	Topk *v1.Topk `protobuf:"bytes,5,opt,name=topk,proto3,oneof"`
}

type ClusterWidgetWrapper_TopkMultiValue struct {
	TopkMultiValue *v1.TopkMultiValue `protobuf:"bytes,6,opt,name=topk_multi_value,json=topkMultiValue,proto3,oneof"`
}

func (*ClusterWidgetWrapper_Donut) isClusterWidgetWrapper_Widget() {}

func (*ClusterWidgetWrapper_Gauge) isClusterWidgetWrapper_Widget() {}

func (*ClusterWidgetWrapper_LineChart) isClusterWidgetWrapper_Widget() {}

func (*ClusterWidgetWrapper_MapChart) isClusterWidgetWrapper_Widget() {}

func (*ClusterWidgetWrapper_Topk) isClusterWidgetWrapper_Widget() {}

func (*ClusterWidgetWrapper_TopkMultiValue) isClusterWidgetWrapper_Widget() {}

var File_com_coralogixapis_apm_queries_k8s_v1_cluster_widget_wrapper_proto protoreflect.FileDescriptor

var file_com_coralogixapis_apm_queries_k8s_v1_cluster_widget_wrapper_proto_rawDesc = []byte{
	0x0a, 0x41, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2f,
	0x6b, 0x38, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x77,
	0x69, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x2c, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d,
	0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x6e, 0x75,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x77,
	0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x77, 0x69, 0x64,
	0x67, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x63, 0x68, 0x61,
	0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f,
	0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x70, 0x5f, 0x63,
	0x68, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70,
	0x6d, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x70,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x37, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x77,
	0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x70, 0x6b, 0x5f, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd7, 0x03, 0x0a, 0x14, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x57, 0x69, 0x64, 0x67,
	0x65, 0x74, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x05, 0x64, 0x6f, 0x6e,
	0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d,
	0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x6e, 0x75,
	0x74, 0x48, 0x00, 0x52, 0x05, 0x64, 0x6f, 0x6e, 0x75, 0x74, 0x12, 0x3f, 0x0a, 0x05, 0x67, 0x61,
	0x75, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70,
	0x6d, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x75,
	0x67, 0x65, 0x48, 0x00, 0x52, 0x05, 0x67, 0x61, 0x75, 0x67, 0x65, 0x12, 0x4c, 0x0a, 0x0a, 0x6c,
	0x69, 0x6e, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x72, 0x74, 0x48, 0x00, 0x52, 0x09,
	0x6c, 0x69, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x72, 0x74, 0x12, 0x49, 0x0a, 0x09, 0x6d, 0x61, 0x70,
	0x5f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x61, 0x70, 0x43, 0x68, 0x61, 0x72, 0x74, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x61, 0x70, 0x43,
	0x68, 0x61, 0x72, 0x74, 0x12, 0x3c, 0x0a, 0x04, 0x74, 0x6f, 0x70, 0x6b, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70, 0x6b, 0x48, 0x00, 0x52, 0x04, 0x74, 0x6f,
	0x70, 0x6b, 0x12, 0x5c, 0x0a, 0x10, 0x74, 0x6f, 0x70, 0x6b, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x6f, 0x70, 0x6b, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00,
	0x52, 0x0e, 0x74, 0x6f, 0x70, 0x6b, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x08, 0x0a, 0x06, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_com_coralogixapis_apm_queries_k8s_v1_cluster_widget_wrapper_proto_rawDescOnce sync.Once
	file_com_coralogixapis_apm_queries_k8s_v1_cluster_widget_wrapper_proto_rawDescData = file_com_coralogixapis_apm_queries_k8s_v1_cluster_widget_wrapper_proto_rawDesc
)

func file_com_coralogixapis_apm_queries_k8s_v1_cluster_widget_wrapper_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_apm_queries_k8s_v1_cluster_widget_wrapper_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_apm_queries_k8s_v1_cluster_widget_wrapper_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_apm_queries_k8s_v1_cluster_widget_wrapper_proto_rawDescData)
	})
	return file_com_coralogixapis_apm_queries_k8s_v1_cluster_widget_wrapper_proto_rawDescData
}

var file_com_coralogixapis_apm_queries_k8s_v1_cluster_widget_wrapper_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_apm_queries_k8s_v1_cluster_widget_wrapper_proto_goTypes = []any{
	(*ClusterWidgetWrapper)(nil), // 0: com.coralogixapis.apm.queries.k8s.v1.ClusterWidgetWrapper
	(*v1.Donut)(nil),             // 1: com.coralogixapis.apm.widgets.v1.Donut
	(*v1.Gauge)(nil),             // 2: com.coralogixapis.apm.widgets.v1.Gauge
	(*v1.LineChart)(nil),         // 3: com.coralogixapis.apm.widgets.v1.LineChart
	(*v1.MapChart)(nil),          // 4: com.coralogixapis.apm.widgets.v1.MapChart
	(*v1.Topk)(nil),              // 5: com.coralogixapis.apm.widgets.v1.Topk
	(*v1.TopkMultiValue)(nil),    // 6: com.coralogixapis.apm.widgets.v1.TopkMultiValue
}
var file_com_coralogixapis_apm_queries_k8s_v1_cluster_widget_wrapper_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.apm.queries.k8s.v1.ClusterWidgetWrapper.donut:type_name -> com.coralogixapis.apm.widgets.v1.Donut
	2, // 1: com.coralogixapis.apm.queries.k8s.v1.ClusterWidgetWrapper.gauge:type_name -> com.coralogixapis.apm.widgets.v1.Gauge
	3, // 2: com.coralogixapis.apm.queries.k8s.v1.ClusterWidgetWrapper.line_chart:type_name -> com.coralogixapis.apm.widgets.v1.LineChart
	4, // 3: com.coralogixapis.apm.queries.k8s.v1.ClusterWidgetWrapper.map_chart:type_name -> com.coralogixapis.apm.widgets.v1.MapChart
	5, // 4: com.coralogixapis.apm.queries.k8s.v1.ClusterWidgetWrapper.topk:type_name -> com.coralogixapis.apm.widgets.v1.Topk
	6, // 5: com.coralogixapis.apm.queries.k8s.v1.ClusterWidgetWrapper.topk_multi_value:type_name -> com.coralogixapis.apm.widgets.v1.TopkMultiValue
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_apm_queries_k8s_v1_cluster_widget_wrapper_proto_init() }
func file_com_coralogixapis_apm_queries_k8s_v1_cluster_widget_wrapper_proto_init() {
	if File_com_coralogixapis_apm_queries_k8s_v1_cluster_widget_wrapper_proto != nil {
		return
	}
	file_com_coralogixapis_apm_queries_k8s_v1_cluster_widget_wrapper_proto_msgTypes[0].OneofWrappers = []any{
		(*ClusterWidgetWrapper_Donut)(nil),
		(*ClusterWidgetWrapper_Gauge)(nil),
		(*ClusterWidgetWrapper_LineChart)(nil),
		(*ClusterWidgetWrapper_MapChart)(nil),
		(*ClusterWidgetWrapper_Topk)(nil),
		(*ClusterWidgetWrapper_TopkMultiValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_apm_queries_k8s_v1_cluster_widget_wrapper_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_apm_queries_k8s_v1_cluster_widget_wrapper_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_apm_queries_k8s_v1_cluster_widget_wrapper_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_apm_queries_k8s_v1_cluster_widget_wrapper_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_apm_queries_k8s_v1_cluster_widget_wrapper_proto = out.File
	file_com_coralogixapis_apm_queries_k8s_v1_cluster_widget_wrapper_proto_rawDesc = nil
	file_com_coralogixapis_apm_queries_k8s_v1_cluster_widget_wrapper_proto_goTypes = nil
	file_com_coralogixapis_apm_queries_k8s_v1_cluster_widget_wrapper_proto_depIdxs = nil
}
