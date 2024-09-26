// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.28.1
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/metric/common/metrics_filter.proto

package v3

import (
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

type MetricFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*MetricFilter_Promql
	Type isMetricFilter_Type `protobuf_oneof:"type"`
}

func (x *MetricFilter) Reset() {
	*x = MetricFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricFilter) ProtoMessage() {}

func (x *MetricFilter) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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

func (m *MetricFilter) GetType() isMetricFilter_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *MetricFilter) GetPromql() *wrapperspb.StringValue {
	if x, ok := x.GetType().(*MetricFilter_Promql); ok {
		return x.Promql
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

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_rawDesc = []byte{
	0x0a, 0x58, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x0c, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x71,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x71, 0x6c, 0x42,
	0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_rawDescData = file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_goTypes = []interface{}{
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
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricFilter); i {
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
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*MetricFilter_Promql)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_rawDesc,
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
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metrics_filter_proto_depIdxs = nil
}
