// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/logs/ratio/logs_ratio_timewindow.proto

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

type LogsRatioTimeWindowValue int32

const (
	LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_MINUTES_5_OR_UNSPECIFIED LogsRatioTimeWindowValue = 0
	LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_MINUTES_10               LogsRatioTimeWindowValue = 1
	LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_MINUTES_15               LogsRatioTimeWindowValue = 2
	LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_MINUTES_30               LogsRatioTimeWindowValue = 3
	LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_HOUR_1                   LogsRatioTimeWindowValue = 4
	LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_HOURS_2                  LogsRatioTimeWindowValue = 5
	LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_HOURS_4                  LogsRatioTimeWindowValue = 6
	LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_HOURS_6                  LogsRatioTimeWindowValue = 7
	LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_HOURS_12                 LogsRatioTimeWindowValue = 8
	LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_HOURS_24                 LogsRatioTimeWindowValue = 9
	LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_HOURS_36                 LogsRatioTimeWindowValue = 10
)

// Enum value maps for LogsRatioTimeWindowValue.
var (
	LogsRatioTimeWindowValue_name = map[int32]string{
		0:  "LOGS_RATIO_TIME_WINDOW_VALUE_MINUTES_5_OR_UNSPECIFIED",
		1:  "LOGS_RATIO_TIME_WINDOW_VALUE_MINUTES_10",
		2:  "LOGS_RATIO_TIME_WINDOW_VALUE_MINUTES_15",
		3:  "LOGS_RATIO_TIME_WINDOW_VALUE_MINUTES_30",
		4:  "LOGS_RATIO_TIME_WINDOW_VALUE_HOUR_1",
		5:  "LOGS_RATIO_TIME_WINDOW_VALUE_HOURS_2",
		6:  "LOGS_RATIO_TIME_WINDOW_VALUE_HOURS_4",
		7:  "LOGS_RATIO_TIME_WINDOW_VALUE_HOURS_6",
		8:  "LOGS_RATIO_TIME_WINDOW_VALUE_HOURS_12",
		9:  "LOGS_RATIO_TIME_WINDOW_VALUE_HOURS_24",
		10: "LOGS_RATIO_TIME_WINDOW_VALUE_HOURS_36",
	}
	LogsRatioTimeWindowValue_value = map[string]int32{
		"LOGS_RATIO_TIME_WINDOW_VALUE_MINUTES_5_OR_UNSPECIFIED": 0,
		"LOGS_RATIO_TIME_WINDOW_VALUE_MINUTES_10":               1,
		"LOGS_RATIO_TIME_WINDOW_VALUE_MINUTES_15":               2,
		"LOGS_RATIO_TIME_WINDOW_VALUE_MINUTES_30":               3,
		"LOGS_RATIO_TIME_WINDOW_VALUE_HOUR_1":                   4,
		"LOGS_RATIO_TIME_WINDOW_VALUE_HOURS_2":                  5,
		"LOGS_RATIO_TIME_WINDOW_VALUE_HOURS_4":                  6,
		"LOGS_RATIO_TIME_WINDOW_VALUE_HOURS_6":                  7,
		"LOGS_RATIO_TIME_WINDOW_VALUE_HOURS_12":                 8,
		"LOGS_RATIO_TIME_WINDOW_VALUE_HOURS_24":                 9,
		"LOGS_RATIO_TIME_WINDOW_VALUE_HOURS_36":                 10,
	}
)

func (x LogsRatioTimeWindowValue) Enum() *LogsRatioTimeWindowValue {
	p := new(LogsRatioTimeWindowValue)
	*p = x
	return p
}

func (x LogsRatioTimeWindowValue) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogsRatioTimeWindowValue) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_enumTypes[0].Descriptor()
}

func (LogsRatioTimeWindowValue) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_enumTypes[0]
}

func (x LogsRatioTimeWindowValue) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogsRatioTimeWindowValue.Descriptor instead.
func (LogsRatioTimeWindowValue) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_rawDescGZIP(), []int{0}
}

type LogsRatioTimeWindow struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Type:
	//
	//	*LogsRatioTimeWindow_LogsRatioTimeWindowSpecificValue
	Type          isLogsRatioTimeWindow_Type `protobuf_oneof:"type"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogsRatioTimeWindow) Reset() {
	*x = LogsRatioTimeWindow{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsRatioTimeWindow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsRatioTimeWindow) ProtoMessage() {}

func (x *LogsRatioTimeWindow) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsRatioTimeWindow.ProtoReflect.Descriptor instead.
func (*LogsRatioTimeWindow) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_rawDescGZIP(), []int{0}
}

func (x *LogsRatioTimeWindow) GetType() isLogsRatioTimeWindow_Type {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *LogsRatioTimeWindow) GetLogsRatioTimeWindowSpecificValue() LogsRatioTimeWindowValue {
	if x != nil {
		if x, ok := x.Type.(*LogsRatioTimeWindow_LogsRatioTimeWindowSpecificValue); ok {
			return x.LogsRatioTimeWindowSpecificValue
		}
	}
	return LogsRatioTimeWindowValue_LOGS_RATIO_TIME_WINDOW_VALUE_MINUTES_5_OR_UNSPECIFIED
}

type isLogsRatioTimeWindow_Type interface {
	isLogsRatioTimeWindow_Type()
}

type LogsRatioTimeWindow_LogsRatioTimeWindowSpecificValue struct {
	LogsRatioTimeWindowSpecificValue LogsRatioTimeWindowValue `protobuf:"varint,1,opt,name=logs_ratio_time_window_specific_value,json=logsRatioTimeWindowSpecificValue,proto3,enum=com.coralogixapis.alerts.v3.LogsRatioTimeWindowValue,oneof"`
}

func (*LogsRatioTimeWindow_LogsRatioTimeWindowSpecificValue) isLogsRatioTimeWindow_Type() {}

var File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_rawDesc = []byte{
	0x0a, 0x5c, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x22, 0xa8, 0x01, 0x0a, 0x13,
	0x4c, 0x6f, 0x67, 0x73, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e,
	0x64, 0x6f, 0x77, 0x12, 0x88, 0x01, 0x0a, 0x25, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x73,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76,
	0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x57,
	0x69, 0x6e, 0x64, 0x6f, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x20, 0x6c, 0x6f,
	0x67, 0x73, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f,
	0x77, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x06,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x2a, 0x84, 0x04, 0x0a, 0x18, 0x4c, 0x6f, 0x67, 0x73, 0x52,
	0x61, 0x74, 0x69, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x39, 0x0a, 0x35, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x52, 0x41, 0x54, 0x49,
	0x4f, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41,
	0x4c, 0x55, 0x45, 0x5f, 0x4d, 0x49, 0x4e, 0x55, 0x54, 0x45, 0x53, 0x5f, 0x35, 0x5f, 0x4f, 0x52,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x2b,
	0x0a, 0x27, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x5f, 0x54, 0x49, 0x4d,
	0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4d,
	0x49, 0x4e, 0x55, 0x54, 0x45, 0x53, 0x5f, 0x31, 0x30, 0x10, 0x01, 0x12, 0x2b, 0x0a, 0x27, 0x4c,
	0x4f, 0x47, 0x53, 0x5f, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57,
	0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4d, 0x49, 0x4e, 0x55,
	0x54, 0x45, 0x53, 0x5f, 0x31, 0x35, 0x10, 0x02, 0x12, 0x2b, 0x0a, 0x27, 0x4c, 0x4f, 0x47, 0x53,
	0x5f, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44,
	0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4d, 0x49, 0x4e, 0x55, 0x54, 0x45, 0x53,
	0x5f, 0x33, 0x30, 0x10, 0x03, 0x12, 0x27, 0x0a, 0x23, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x52, 0x41,
	0x54, 0x49, 0x4f, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f,
	0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x48, 0x4f, 0x55, 0x52, 0x5f, 0x31, 0x10, 0x04, 0x12, 0x28,
	0x0a, 0x24, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x5f, 0x54, 0x49, 0x4d,
	0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x48,
	0x4f, 0x55, 0x52, 0x53, 0x5f, 0x32, 0x10, 0x05, 0x12, 0x28, 0x0a, 0x24, 0x4c, 0x4f, 0x47, 0x53,
	0x5f, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44,
	0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x48, 0x4f, 0x55, 0x52, 0x53, 0x5f, 0x34,
	0x10, 0x06, 0x12, 0x28, 0x0a, 0x24, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x52, 0x41, 0x54, 0x49, 0x4f,
	0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c,
	0x55, 0x45, 0x5f, 0x48, 0x4f, 0x55, 0x52, 0x53, 0x5f, 0x36, 0x10, 0x07, 0x12, 0x29, 0x0a, 0x25,
	0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f,
	0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x48, 0x4f, 0x55,
	0x52, 0x53, 0x5f, 0x31, 0x32, 0x10, 0x08, 0x12, 0x29, 0x0a, 0x25, 0x4c, 0x4f, 0x47, 0x53, 0x5f,
	0x52, 0x41, 0x54, 0x49, 0x4f, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f,
	0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x48, 0x4f, 0x55, 0x52, 0x53, 0x5f, 0x32, 0x34,
	0x10, 0x09, 0x12, 0x29, 0x0a, 0x25, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x52, 0x41, 0x54, 0x49, 0x4f,
	0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c,
	0x55, 0x45, 0x5f, 0x48, 0x4f, 0x55, 0x52, 0x53, 0x5f, 0x33, 0x36, 0x10, 0x0a, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_rawDescData = file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_goTypes = []any{
	(LogsRatioTimeWindowValue)(0), // 0: com.coralogixapis.alerts.v3.LogsRatioTimeWindowValue
	(*LogsRatioTimeWindow)(nil),   // 1: com.coralogixapis.alerts.v3.LogsRatioTimeWindow
}
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_depIdxs = []int32{
	0, // 0: com.coralogixapis.alerts.v3.LogsRatioTimeWindow.logs_ratio_time_window_specific_value:type_name -> com.coralogixapis.alerts.v3.LogsRatioTimeWindowValue
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_init()
}
func file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_msgTypes[0].OneofWrappers = []any{
		(*LogsRatioTimeWindow_LogsRatioTimeWindowSpecificValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_ratio_logs_ratio_timewindow_proto_depIdxs = nil
}
