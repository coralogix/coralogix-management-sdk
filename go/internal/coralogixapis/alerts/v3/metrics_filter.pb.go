// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/metric/common/metrics_filter.proto

package v3

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

type MetricFilter struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Type:
	//
	//	*MetricFilter_Promql
	Type          isMetricFilter_Type `protobuf_oneof:"type"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MetricFilter) Reset() {
	*x = MetricFilter{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetricFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricFilter) ProtoMessage() {}

func (x *MetricFilter) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricFilter.ProtoReflect.Descriptor instead.
func (*MetricFilter) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_rawDescGZIP(), []int{0}
}

func (x *MetricFilter) GetType() isMetricFilter_Type {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *MetricFilter) GetPromql() *wrapperspb.StringValue {
	if x != nil {
		if x, ok := x.Type.(*MetricFilter_Promql); ok {
			return x.Promql
		}
	}
	return nil
}

type isMetricFilter_Type interface {
	isMetricFilter_Type()
}

type MetricFilter_Promql struct {
	Promql *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=promql,proto3,oneof"`
}

func (*MetricFilter_Promql) isMetricFilter_Type() {}

var File_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto protoreflect.FileDescriptor

const file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_rawDesc = "" +
	"\n" +
	"Xcom/coralogixapis/alerts/v3/alert_def_type_definition/metric/common/metrics_filter.proto\x12\x1bcom.coralogixapis.alerts.v3\x1a\x1egoogle/protobuf/wrappers.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\xd8\x01\n" +
	"\fMetricFilter\x12\x88\x01\n" +
	"\x06promql\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueBP\x92AM2\x1bA PromQL filter for metricsJ%\"avg_over_time(metric_name[5m]) > 10\"\xa2\x02\x06promqlH\x00R\x06promql:5\x92A2\n" +
	"0*%Metric filter configuration in alerts\xd2\x01\x06promqlB\x06\n" +
	"\x04typeb\x06proto3"

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_goTypes = []any{
	(*MetricFilter)(nil),           // 0: com.coralogixapis.alerts.v3.MetricFilter
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
}
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.alerts.v3.MetricFilter.promql:type_name -> google.protobuf.StringValue
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_init()
}
func file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_msgTypes[0].OneofWrappers = []any{
		(*MetricFilter_Promql)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_depIdxs = nil
}
