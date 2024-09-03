// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.3
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/logs/new_value/logs_new_value_timewindow.proto

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

type LogsNewValueTimeWindowValue int32

const (
	LogsNewValueTimeWindowValue_LOGS_NEW_VALUE_TIME_WINDOW_VALUE_HOURS_12_OR_UNSPECIFIED LogsNewValueTimeWindowValue = 0
	LogsNewValueTimeWindowValue_LOGS_NEW_VALUE_TIME_WINDOW_VALUE_HOURS_24                LogsNewValueTimeWindowValue = 1
	LogsNewValueTimeWindowValue_LOGS_NEW_VALUE_TIME_WINDOW_VALUE_HOURS_48                LogsNewValueTimeWindowValue = 2
	LogsNewValueTimeWindowValue_LOGS_NEW_VALUE_TIME_WINDOW_VALUE_HOURS_72                LogsNewValueTimeWindowValue = 3
	LogsNewValueTimeWindowValue_LOGS_NEW_VALUE_TIME_WINDOW_VALUE_WEEK_1                  LogsNewValueTimeWindowValue = 4
	LogsNewValueTimeWindowValue_LOGS_NEW_VALUE_TIME_WINDOW_VALUE_MONTH_1                 LogsNewValueTimeWindowValue = 5
	LogsNewValueTimeWindowValue_LOGS_NEW_VALUE_TIME_WINDOW_VALUE_MONTHS_2                LogsNewValueTimeWindowValue = 6
	LogsNewValueTimeWindowValue_LOGS_NEW_VALUE_TIME_WINDOW_VALUE_MONTHS_3                LogsNewValueTimeWindowValue = 7
)

// Enum value maps for LogsNewValueTimeWindowValue.
var (
	LogsNewValueTimeWindowValue_name = map[int32]string{
		0: "LOGS_NEW_VALUE_TIME_WINDOW_VALUE_HOURS_12_OR_UNSPECIFIED",
		1: "LOGS_NEW_VALUE_TIME_WINDOW_VALUE_HOURS_24",
		2: "LOGS_NEW_VALUE_TIME_WINDOW_VALUE_HOURS_48",
		3: "LOGS_NEW_VALUE_TIME_WINDOW_VALUE_HOURS_72",
		4: "LOGS_NEW_VALUE_TIME_WINDOW_VALUE_WEEK_1",
		5: "LOGS_NEW_VALUE_TIME_WINDOW_VALUE_MONTH_1",
		6: "LOGS_NEW_VALUE_TIME_WINDOW_VALUE_MONTHS_2",
		7: "LOGS_NEW_VALUE_TIME_WINDOW_VALUE_MONTHS_3",
	}
	LogsNewValueTimeWindowValue_value = map[string]int32{
		"LOGS_NEW_VALUE_TIME_WINDOW_VALUE_HOURS_12_OR_UNSPECIFIED": 0,
		"LOGS_NEW_VALUE_TIME_WINDOW_VALUE_HOURS_24":                1,
		"LOGS_NEW_VALUE_TIME_WINDOW_VALUE_HOURS_48":                2,
		"LOGS_NEW_VALUE_TIME_WINDOW_VALUE_HOURS_72":                3,
		"LOGS_NEW_VALUE_TIME_WINDOW_VALUE_WEEK_1":                  4,
		"LOGS_NEW_VALUE_TIME_WINDOW_VALUE_MONTH_1":                 5,
		"LOGS_NEW_VALUE_TIME_WINDOW_VALUE_MONTHS_2":                6,
		"LOGS_NEW_VALUE_TIME_WINDOW_VALUE_MONTHS_3":                7,
	}
)

func (x LogsNewValueTimeWindowValue) Enum() *LogsNewValueTimeWindowValue {
	p := new(LogsNewValueTimeWindowValue)
	*p = x
	return p
}

func (x LogsNewValueTimeWindowValue) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogsNewValueTimeWindowValue) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_enumTypes[0].Descriptor()
}

func (LogsNewValueTimeWindowValue) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_enumTypes[0]
}

func (x LogsNewValueTimeWindowValue) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogsNewValueTimeWindowValue.Descriptor instead.
func (LogsNewValueTimeWindowValue) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_rawDescGZIP(), []int{0}
}

type LogsNewValueTimeWindow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*LogsNewValueTimeWindow_LogsNewValueTimeWindowSpecificValue
	Type isLogsNewValueTimeWindow_Type `protobuf_oneof:"type"`
}

func (x *LogsNewValueTimeWindow) Reset() {
	*x = LogsNewValueTimeWindow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsNewValueTimeWindow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsNewValueTimeWindow) ProtoMessage() {}

func (x *LogsNewValueTimeWindow) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsNewValueTimeWindow.ProtoReflect.Descriptor instead.
func (*LogsNewValueTimeWindow) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_rawDescGZIP(), []int{0}
}

func (m *LogsNewValueTimeWindow) GetType() isLogsNewValueTimeWindow_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *LogsNewValueTimeWindow) GetLogsNewValueTimeWindowSpecificValue() LogsNewValueTimeWindowValue {
	if x, ok := x.GetType().(*LogsNewValueTimeWindow_LogsNewValueTimeWindowSpecificValue); ok {
		return x.LogsNewValueTimeWindowSpecificValue
	}
	return LogsNewValueTimeWindowValue_LOGS_NEW_VALUE_TIME_WINDOW_VALUE_HOURS_12_OR_UNSPECIFIED
}

type isLogsNewValueTimeWindow_Type interface {
	isLogsNewValueTimeWindow_Type()
}

type LogsNewValueTimeWindow_LogsNewValueTimeWindowSpecificValue struct {
	LogsNewValueTimeWindowSpecificValue LogsNewValueTimeWindowValue `protobuf:"varint,1,opt,name=logs_new_value_time_window_specific_value,json=logsNewValueTimeWindowSpecificValue,proto3,enum=com.coralogixapis.alerts.v3.LogsNewValueTimeWindowValue,oneof"`
}

func (*LogsNewValueTimeWindow_LogsNewValueTimeWindowSpecificValue) isLogsNewValueTimeWindow_Type() {}

var File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_rawDesc = []byte{
	0x0a, 0x64, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x6e, 0x65, 0x77,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x6e, 0x65, 0x77, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x2e, 0x76, 0x33, 0x22, 0xb5, 0x01, 0x0a, 0x16, 0x4c, 0x6f, 0x67, 0x73, 0x4e, 0x65, 0x77, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x92,
	0x01, 0x0a, 0x29, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33,
	0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x4e, 0x65, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x23,
	0x6c, 0x6f, 0x67, 0x73, 0x4e, 0x65, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x2a, 0xa1, 0x03, 0x0a, 0x1b,
	0x4c, 0x6f, 0x67, 0x73, 0x4e, 0x65, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3c, 0x0a, 0x38, 0x4c,
	0x4f, 0x47, 0x53, 0x5f, 0x4e, 0x45, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54, 0x49,
	0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f,
	0x48, 0x4f, 0x55, 0x52, 0x53, 0x5f, 0x31, 0x32, 0x5f, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x2d, 0x0a, 0x29, 0x4c, 0x4f, 0x47,
	0x53, 0x5f, 0x4e, 0x45, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45,
	0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x48, 0x4f,
	0x55, 0x52, 0x53, 0x5f, 0x32, 0x34, 0x10, 0x01, 0x12, 0x2d, 0x0a, 0x29, 0x4c, 0x4f, 0x47, 0x53,
	0x5f, 0x4e, 0x45, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f,
	0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x48, 0x4f, 0x55,
	0x52, 0x53, 0x5f, 0x34, 0x38, 0x10, 0x02, 0x12, 0x2d, 0x0a, 0x29, 0x4c, 0x4f, 0x47, 0x53, 0x5f,
	0x4e, 0x45, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57,
	0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x48, 0x4f, 0x55, 0x52,
	0x53, 0x5f, 0x37, 0x32, 0x10, 0x03, 0x12, 0x2b, 0x0a, 0x27, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x4e,
	0x45, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49,
	0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x57, 0x45, 0x45, 0x4b, 0x5f,
	0x31, 0x10, 0x04, 0x12, 0x2c, 0x0a, 0x28, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x4e, 0x45, 0x57, 0x5f,
	0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f,
	0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x5f, 0x31, 0x10,
	0x05, 0x12, 0x2d, 0x0a, 0x29, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x4e, 0x45, 0x57, 0x5f, 0x56, 0x41,
	0x4c, 0x55, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f,
	0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x53, 0x5f, 0x32, 0x10, 0x06,
	0x12, 0x2d, 0x0a, 0x29, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x4e, 0x45, 0x57, 0x5f, 0x56, 0x41, 0x4c,
	0x55, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56,
	0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x53, 0x5f, 0x33, 0x10, 0x07, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_rawDescData = file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_goTypes = []interface{}{
	(LogsNewValueTimeWindowValue)(0), // 0: com.coralogixapis.alerts.v3.LogsNewValueTimeWindowValue
	(*LogsNewValueTimeWindow)(nil),   // 1: com.coralogixapis.alerts.v3.LogsNewValueTimeWindow
}
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_depIdxs = []int32{
	0, // 0: com.coralogixapis.alerts.v3.LogsNewValueTimeWindow.logs_new_value_time_window_specific_value:type_name -> com.coralogixapis.alerts.v3.LogsNewValueTimeWindowValue
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_init()
}
func file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogsNewValueTimeWindow); i {
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
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*LogsNewValueTimeWindow_LogsNewValueTimeWindowSpecificValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_depIdxs = nil
}