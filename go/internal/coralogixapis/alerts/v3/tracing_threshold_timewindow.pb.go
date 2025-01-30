// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.2
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/tracing/threshold/tracing_threshold_timewindow.proto

package v3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type TracingTimeWindowValue int32

const (
	TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_MINUTES_5_OR_UNSPECIFIED TracingTimeWindowValue = 0
	TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_MINUTES_10               TracingTimeWindowValue = 1
	TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_MINUTES_15               TracingTimeWindowValue = 2
	TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_MINUTES_20               TracingTimeWindowValue = 11
	TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_MINUTES_30               TracingTimeWindowValue = 3
	TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_HOUR_1                   TracingTimeWindowValue = 4
	TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_HOURS_2                  TracingTimeWindowValue = 5
	TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_HOURS_4                  TracingTimeWindowValue = 6
	TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_HOURS_6                  TracingTimeWindowValue = 7
	TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_HOURS_12                 TracingTimeWindowValue = 8
	TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_HOURS_24                 TracingTimeWindowValue = 9
	TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_HOURS_36                 TracingTimeWindowValue = 10
)

// Enum value maps for TracingTimeWindowValue.
var (
	TracingTimeWindowValue_name = map[int32]string{
		0:  "TRACING_TIME_WINDOW_VALUE_MINUTES_5_OR_UNSPECIFIED",
		1:  "TRACING_TIME_WINDOW_VALUE_MINUTES_10",
		2:  "TRACING_TIME_WINDOW_VALUE_MINUTES_15",
		11: "TRACING_TIME_WINDOW_VALUE_MINUTES_20",
		3:  "TRACING_TIME_WINDOW_VALUE_MINUTES_30",
		4:  "TRACING_TIME_WINDOW_VALUE_HOUR_1",
		5:  "TRACING_TIME_WINDOW_VALUE_HOURS_2",
		6:  "TRACING_TIME_WINDOW_VALUE_HOURS_4",
		7:  "TRACING_TIME_WINDOW_VALUE_HOURS_6",
		8:  "TRACING_TIME_WINDOW_VALUE_HOURS_12",
		9:  "TRACING_TIME_WINDOW_VALUE_HOURS_24",
		10: "TRACING_TIME_WINDOW_VALUE_HOURS_36",
	}
	TracingTimeWindowValue_value = map[string]int32{
		"TRACING_TIME_WINDOW_VALUE_MINUTES_5_OR_UNSPECIFIED": 0,
		"TRACING_TIME_WINDOW_VALUE_MINUTES_10":               1,
		"TRACING_TIME_WINDOW_VALUE_MINUTES_15":               2,
		"TRACING_TIME_WINDOW_VALUE_MINUTES_20":               11,
		"TRACING_TIME_WINDOW_VALUE_MINUTES_30":               3,
		"TRACING_TIME_WINDOW_VALUE_HOUR_1":                   4,
		"TRACING_TIME_WINDOW_VALUE_HOURS_2":                  5,
		"TRACING_TIME_WINDOW_VALUE_HOURS_4":                  6,
		"TRACING_TIME_WINDOW_VALUE_HOURS_6":                  7,
		"TRACING_TIME_WINDOW_VALUE_HOURS_12":                 8,
		"TRACING_TIME_WINDOW_VALUE_HOURS_24":                 9,
		"TRACING_TIME_WINDOW_VALUE_HOURS_36":                 10,
	}
)

func (x TracingTimeWindowValue) Enum() *TracingTimeWindowValue {
	p := new(TracingTimeWindowValue)
	*p = x
	return p
}

func (x TracingTimeWindowValue) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TracingTimeWindowValue) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto_enumTypes[0].Descriptor()
}

func (TracingTimeWindowValue) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto_enumTypes[0]
}

func (x TracingTimeWindowValue) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TracingTimeWindowValue.Descriptor instead.
func (TracingTimeWindowValue) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto_rawDescGZIP(), []int{0}
}

type TracingTimeWindow struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Type:
	//
	//	*TracingTimeWindow_TracingTimeWindowValue
	Type          isTracingTimeWindow_Type `protobuf_oneof:"type"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TracingTimeWindow) Reset() {
	*x = TracingTimeWindow{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TracingTimeWindow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TracingTimeWindow) ProtoMessage() {}

func (x *TracingTimeWindow) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TracingTimeWindow.ProtoReflect.Descriptor instead.
func (*TracingTimeWindow) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto_rawDescGZIP(), []int{0}
}

func (x *TracingTimeWindow) GetType() isTracingTimeWindow_Type {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *TracingTimeWindow) GetTracingTimeWindowValue() TracingTimeWindowValue {
	if x != nil {
		if x, ok := x.Type.(*TracingTimeWindow_TracingTimeWindowValue); ok {
			return x.TracingTimeWindowValue
		}
	}
	return TracingTimeWindowValue_TRACING_TIME_WINDOW_VALUE_MINUTES_5_OR_UNSPECIFIED
}

type isTracingTimeWindow_Type interface {
	isTracingTimeWindow_Type()
}

type TracingTimeWindow_TracingTimeWindowValue struct {
	TracingTimeWindowValue TracingTimeWindowValue `protobuf:"varint,1,opt,name=tracing_time_window_value,json=tracingTimeWindowValue,proto3,enum=com.coralogixapis.alerts.v3.TracingTimeWindowValue,oneof"`
}

func (*TracingTimeWindow_TracingTimeWindowValue) isTracingTimeWindow_Type() {}

var File_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto_rawDesc = string([]byte{
	0x0a, 0x6a, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2f,
	0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e,
	0x67, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x22, 0x8d, 0x01, 0x0a, 0x11, 0x54, 0x72,
	0x61, 0x63, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12,
	0x70, 0x0a, 0x19, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33,
	0x2e, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64,
	0x6f, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x16, 0x74, 0x72, 0x61, 0x63, 0x69,
	0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x2a, 0x8b, 0x04, 0x0a, 0x16, 0x54, 0x72,
	0x61, 0x63, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x36, 0x0a, 0x32, 0x54, 0x52, 0x41, 0x43, 0x49, 0x4e, 0x47, 0x5f,
	0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55,
	0x45, 0x5f, 0x4d, 0x49, 0x4e, 0x55, 0x54, 0x45, 0x53, 0x5f, 0x35, 0x5f, 0x4f, 0x52, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x28, 0x0a, 0x24,
	0x54, 0x52, 0x41, 0x43, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e,
	0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4d, 0x49, 0x4e, 0x55, 0x54, 0x45,
	0x53, 0x5f, 0x31, 0x30, 0x10, 0x01, 0x12, 0x28, 0x0a, 0x24, 0x54, 0x52, 0x41, 0x43, 0x49, 0x4e,
	0x47, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41,
	0x4c, 0x55, 0x45, 0x5f, 0x4d, 0x49, 0x4e, 0x55, 0x54, 0x45, 0x53, 0x5f, 0x31, 0x35, 0x10, 0x02,
	0x12, 0x28, 0x0a, 0x24, 0x54, 0x52, 0x41, 0x43, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x49, 0x4d, 0x45,
	0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4d, 0x49,
	0x4e, 0x55, 0x54, 0x45, 0x53, 0x5f, 0x32, 0x30, 0x10, 0x0b, 0x12, 0x28, 0x0a, 0x24, 0x54, 0x52,
	0x41, 0x43, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f,
	0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4d, 0x49, 0x4e, 0x55, 0x54, 0x45, 0x53, 0x5f,
	0x33, 0x30, 0x10, 0x03, 0x12, 0x24, 0x0a, 0x20, 0x54, 0x52, 0x41, 0x43, 0x49, 0x4e, 0x47, 0x5f,
	0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55,
	0x45, 0x5f, 0x48, 0x4f, 0x55, 0x52, 0x5f, 0x31, 0x10, 0x04, 0x12, 0x25, 0x0a, 0x21, 0x54, 0x52,
	0x41, 0x43, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f,
	0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x48, 0x4f, 0x55, 0x52, 0x53, 0x5f, 0x32, 0x10,
	0x05, 0x12, 0x25, 0x0a, 0x21, 0x54, 0x52, 0x41, 0x43, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x49, 0x4d,
	0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x48,
	0x4f, 0x55, 0x52, 0x53, 0x5f, 0x34, 0x10, 0x06, 0x12, 0x25, 0x0a, 0x21, 0x54, 0x52, 0x41, 0x43,
	0x49, 0x4e, 0x47, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f,
	0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x48, 0x4f, 0x55, 0x52, 0x53, 0x5f, 0x36, 0x10, 0x07, 0x12,
	0x26, 0x0a, 0x22, 0x54, 0x52, 0x41, 0x43, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f,
	0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x48, 0x4f, 0x55,
	0x52, 0x53, 0x5f, 0x31, 0x32, 0x10, 0x08, 0x12, 0x26, 0x0a, 0x22, 0x54, 0x52, 0x41, 0x43, 0x49,
	0x4e, 0x47, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56,
	0x41, 0x4c, 0x55, 0x45, 0x5f, 0x48, 0x4f, 0x55, 0x52, 0x53, 0x5f, 0x32, 0x34, 0x10, 0x09, 0x12,
	0x26, 0x0a, 0x22, 0x54, 0x52, 0x41, 0x43, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f,
	0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x48, 0x4f, 0x55,
	0x52, 0x53, 0x5f, 0x33, 0x36, 0x10, 0x0a, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto_goTypes = []any{
	(TracingTimeWindowValue)(0), // 0: com.coralogixapis.alerts.v3.TracingTimeWindowValue
	(*TracingTimeWindow)(nil),   // 1: com.coralogixapis.alerts.v3.TracingTimeWindow
}
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto_depIdxs = []int32{
	0, // 0: com.coralogixapis.alerts.v3.TracingTimeWindow.tracing_time_window_value:type_name -> com.coralogixapis.alerts.v3.TracingTimeWindowValue
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto_init()
}
func file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto_msgTypes[0].OneofWrappers = []any{
		(*TracingTimeWindow_TracingTimeWindowValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_timewindow_proto_depIdxs = nil
}
