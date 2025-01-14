// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/metric/common/metric_timewindow.proto

package v3

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

type MetricTimeWindowValue int32

const (
	MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_MINUTES_1_OR_UNSPECIFIED MetricTimeWindowValue = 0
	MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_MINUTES_5                MetricTimeWindowValue = 1
	MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_MINUTES_10               MetricTimeWindowValue = 2
	MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_MINUTES_15               MetricTimeWindowValue = 3
	MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_MINUTES_20               MetricTimeWindowValue = 11
	MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_MINUTES_30               MetricTimeWindowValue = 4
	MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_HOUR_1                   MetricTimeWindowValue = 5
	MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_HOURS_2                  MetricTimeWindowValue = 6
	MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_HOURS_4                  MetricTimeWindowValue = 7
	MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_HOURS_6                  MetricTimeWindowValue = 8
	MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_HOURS_12                 MetricTimeWindowValue = 9
	MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_HOURS_24                 MetricTimeWindowValue = 10
	MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_HOURS_36                 MetricTimeWindowValue = 12
)

// Enum value maps for MetricTimeWindowValue.
var (
	MetricTimeWindowValue_name = map[int32]string{
		0:  "METRIC_TIME_WINDOW_VALUE_MINUTES_1_OR_UNSPECIFIED",
		1:  "METRIC_TIME_WINDOW_VALUE_MINUTES_5",
		2:  "METRIC_TIME_WINDOW_VALUE_MINUTES_10",
		3:  "METRIC_TIME_WINDOW_VALUE_MINUTES_15",
		11: "METRIC_TIME_WINDOW_VALUE_MINUTES_20",
		4:  "METRIC_TIME_WINDOW_VALUE_MINUTES_30",
		5:  "METRIC_TIME_WINDOW_VALUE_HOUR_1",
		6:  "METRIC_TIME_WINDOW_VALUE_HOURS_2",
		7:  "METRIC_TIME_WINDOW_VALUE_HOURS_4",
		8:  "METRIC_TIME_WINDOW_VALUE_HOURS_6",
		9:  "METRIC_TIME_WINDOW_VALUE_HOURS_12",
		10: "METRIC_TIME_WINDOW_VALUE_HOURS_24",
		12: "METRIC_TIME_WINDOW_VALUE_HOURS_36",
	}
	MetricTimeWindowValue_value = map[string]int32{
		"METRIC_TIME_WINDOW_VALUE_MINUTES_1_OR_UNSPECIFIED": 0,
		"METRIC_TIME_WINDOW_VALUE_MINUTES_5":                1,
		"METRIC_TIME_WINDOW_VALUE_MINUTES_10":               2,
		"METRIC_TIME_WINDOW_VALUE_MINUTES_15":               3,
		"METRIC_TIME_WINDOW_VALUE_MINUTES_20":               11,
		"METRIC_TIME_WINDOW_VALUE_MINUTES_30":               4,
		"METRIC_TIME_WINDOW_VALUE_HOUR_1":                   5,
		"METRIC_TIME_WINDOW_VALUE_HOURS_2":                  6,
		"METRIC_TIME_WINDOW_VALUE_HOURS_4":                  7,
		"METRIC_TIME_WINDOW_VALUE_HOURS_6":                  8,
		"METRIC_TIME_WINDOW_VALUE_HOURS_12":                 9,
		"METRIC_TIME_WINDOW_VALUE_HOURS_24":                 10,
		"METRIC_TIME_WINDOW_VALUE_HOURS_36":                 12,
	}
)

func (x MetricTimeWindowValue) Enum() *MetricTimeWindowValue {
	p := new(MetricTimeWindowValue)
	*p = x
	return p
}

func (x MetricTimeWindowValue) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetricTimeWindowValue) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_enumTypes[0].Descriptor()
}

func (MetricTimeWindowValue) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_enumTypes[0]
}

func (x MetricTimeWindowValue) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetricTimeWindowValue.Descriptor instead.
func (MetricTimeWindowValue) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_rawDescGZIP(), []int{0}
}

type MetricTimeWindow struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Type:
	//
	//	*MetricTimeWindow_MetricTimeWindowSpecificValue
	Type          isMetricTimeWindow_Type `protobuf_oneof:"type"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MetricTimeWindow) Reset() {
	*x = MetricTimeWindow{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetricTimeWindow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricTimeWindow) ProtoMessage() {}

func (x *MetricTimeWindow) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricTimeWindow.ProtoReflect.Descriptor instead.
func (*MetricTimeWindow) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_rawDescGZIP(), []int{0}
}

func (x *MetricTimeWindow) GetType() isMetricTimeWindow_Type {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *MetricTimeWindow) GetMetricTimeWindowSpecificValue() MetricTimeWindowValue {
	if x != nil {
		if x, ok := x.Type.(*MetricTimeWindow_MetricTimeWindowSpecificValue); ok {
			return x.MetricTimeWindowSpecificValue
		}
	}
	return MetricTimeWindowValue_METRIC_TIME_WINDOW_VALUE_MINUTES_1_OR_UNSPECIFIED
}

type isMetricTimeWindow_Type interface {
	isMetricTimeWindow_Type()
}

type MetricTimeWindow_MetricTimeWindowSpecificValue struct {
	MetricTimeWindowSpecificValue MetricTimeWindowValue `protobuf:"varint,1,opt,name=metric_time_window_specific_value,json=metricTimeWindowSpecificValue,proto3,enum=com.coralogixapis.alerts.v3.MetricTimeWindowValue,oneof"`
}

func (*MetricTimeWindow_MetricTimeWindowSpecificValue) isMetricTimeWindow_Type() {}

var File_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_rawDesc = []byte{
	0x0a, 0x5b, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x22, 0x9a, 0x01, 0x0a, 0x10, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12,
	0x7e, 0x0a, 0x21, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x77,
	0x69, 0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54,
	0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00,
	0x52, 0x1d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64,
	0x6f, 0x77, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x2a, 0xa6, 0x04, 0x0a, 0x15, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x35, 0x0a, 0x31, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x54, 0x49, 0x4d, 0x45,
	0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4d, 0x49,
	0x4e, 0x55, 0x54, 0x45, 0x53, 0x5f, 0x31, 0x5f, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x26, 0x0a, 0x22, 0x4d, 0x45, 0x54, 0x52,
	0x49, 0x43, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56,
	0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4d, 0x49, 0x4e, 0x55, 0x54, 0x45, 0x53, 0x5f, 0x35, 0x10, 0x01,
	0x12, 0x27, 0x0a, 0x23, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f,
	0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4d, 0x49, 0x4e,
	0x55, 0x54, 0x45, 0x53, 0x5f, 0x31, 0x30, 0x10, 0x02, 0x12, 0x27, 0x0a, 0x23, 0x4d, 0x45, 0x54,
	0x52, 0x49, 0x43, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f,
	0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4d, 0x49, 0x4e, 0x55, 0x54, 0x45, 0x53, 0x5f, 0x31, 0x35,
	0x10, 0x03, 0x12, 0x27, 0x0a, 0x23, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x54, 0x49, 0x4d,
	0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4d,
	0x49, 0x4e, 0x55, 0x54, 0x45, 0x53, 0x5f, 0x32, 0x30, 0x10, 0x0b, 0x12, 0x27, 0x0a, 0x23, 0x4d,
	0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f,
	0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4d, 0x49, 0x4e, 0x55, 0x54, 0x45, 0x53, 0x5f,
	0x33, 0x30, 0x10, 0x04, 0x12, 0x23, 0x0a, 0x1f, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x54,
	0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45,
	0x5f, 0x48, 0x4f, 0x55, 0x52, 0x5f, 0x31, 0x10, 0x05, 0x12, 0x24, 0x0a, 0x20, 0x4d, 0x45, 0x54,
	0x52, 0x49, 0x43, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f,
	0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x48, 0x4f, 0x55, 0x52, 0x53, 0x5f, 0x32, 0x10, 0x06, 0x12,
	0x24, 0x0a, 0x20, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57,
	0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x48, 0x4f, 0x55, 0x52,
	0x53, 0x5f, 0x34, 0x10, 0x07, 0x12, 0x24, 0x0a, 0x20, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f,
	0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55,
	0x45, 0x5f, 0x48, 0x4f, 0x55, 0x52, 0x53, 0x5f, 0x36, 0x10, 0x08, 0x12, 0x25, 0x0a, 0x21, 0x4d,
	0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f,
	0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x48, 0x4f, 0x55, 0x52, 0x53, 0x5f, 0x31, 0x32,
	0x10, 0x09, 0x12, 0x25, 0x0a, 0x21, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x54, 0x49, 0x4d,
	0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x48,
	0x4f, 0x55, 0x52, 0x53, 0x5f, 0x32, 0x34, 0x10, 0x0a, 0x12, 0x25, 0x0a, 0x21, 0x4d, 0x45, 0x54,
	0x52, 0x49, 0x43, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f,
	0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x48, 0x4f, 0x55, 0x52, 0x53, 0x5f, 0x33, 0x36, 0x10, 0x0c,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_rawDescData = file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_goTypes = []any{
	(MetricTimeWindowValue)(0), // 0: com.coralogixapis.alerts.v3.MetricTimeWindowValue
	(*MetricTimeWindow)(nil),   // 1: com.coralogixapis.alerts.v3.MetricTimeWindow
}
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_depIdxs = []int32{
	0, // 0: com.coralogixapis.alerts.v3.MetricTimeWindow.metric_time_window_specific_value:type_name -> com.coralogixapis.alerts.v3.MetricTimeWindowValue
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_init()
}
func file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_msgTypes[0].OneofWrappers = []any{
		(*MetricTimeWindow_MetricTimeWindowSpecificValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_common_metric_timewindow_proto_depIdxs = nil
}
