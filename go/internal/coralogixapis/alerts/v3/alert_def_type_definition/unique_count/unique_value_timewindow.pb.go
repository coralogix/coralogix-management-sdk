// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.3
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/unique_count/unique_value_timewindow.proto

package unique_count

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

type LogsUniqueValueTimeWindowValue int32

const (
	LogsUniqueValueTimeWindowValue_LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_MINUTE_1_OR_UNSPECIFIED LogsUniqueValueTimeWindowValue = 0
	LogsUniqueValueTimeWindowValue_LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_MINUTES_15              LogsUniqueValueTimeWindowValue = 1
	LogsUniqueValueTimeWindowValue_LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_MINUTES_20              LogsUniqueValueTimeWindowValue = 2
	LogsUniqueValueTimeWindowValue_LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_MINUTES_30              LogsUniqueValueTimeWindowValue = 3
	LogsUniqueValueTimeWindowValue_LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_HOURS_1                 LogsUniqueValueTimeWindowValue = 4
	LogsUniqueValueTimeWindowValue_LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_HOURS_2                 LogsUniqueValueTimeWindowValue = 5
	LogsUniqueValueTimeWindowValue_LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_HOURS_4                 LogsUniqueValueTimeWindowValue = 6
	LogsUniqueValueTimeWindowValue_LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_HOURS_6                 LogsUniqueValueTimeWindowValue = 7
	LogsUniqueValueTimeWindowValue_LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_HOURS_12                LogsUniqueValueTimeWindowValue = 8
	LogsUniqueValueTimeWindowValue_LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_HOURS_24                LogsUniqueValueTimeWindowValue = 9
)

// Enum value maps for LogsUniqueValueTimeWindowValue.
var (
	LogsUniqueValueTimeWindowValue_name = map[int32]string{
		0: "LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_MINUTE_1_OR_UNSPECIFIED",
		1: "LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_MINUTES_15",
		2: "LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_MINUTES_20",
		3: "LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_MINUTES_30",
		4: "LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_HOURS_1",
		5: "LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_HOURS_2",
		6: "LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_HOURS_4",
		7: "LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_HOURS_6",
		8: "LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_HOURS_12",
		9: "LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_HOURS_24",
	}
	LogsUniqueValueTimeWindowValue_value = map[string]int32{
		"LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_MINUTE_1_OR_UNSPECIFIED": 0,
		"LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_MINUTES_15":              1,
		"LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_MINUTES_20":              2,
		"LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_MINUTES_30":              3,
		"LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_HOURS_1":                 4,
		"LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_HOURS_2":                 5,
		"LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_HOURS_4":                 6,
		"LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_HOURS_6":                 7,
		"LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_HOURS_12":                8,
		"LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_HOURS_24":                9,
	}
)

func (x LogsUniqueValueTimeWindowValue) Enum() *LogsUniqueValueTimeWindowValue {
	p := new(LogsUniqueValueTimeWindowValue)
	*p = x
	return p
}

func (x LogsUniqueValueTimeWindowValue) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogsUniqueValueTimeWindowValue) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto_enumTypes[0].Descriptor()
}

func (LogsUniqueValueTimeWindowValue) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto_enumTypes[0]
}

func (x LogsUniqueValueTimeWindowValue) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogsUniqueValueTimeWindowValue.Descriptor instead.
func (LogsUniqueValueTimeWindowValue) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto_rawDescGZIP(), []int{0}
}

type LogsUniqueValueTimeWindow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*LogsUniqueValueTimeWindow_LogsUniqueValueTimeWindowSpecificValue
	Type isLogsUniqueValueTimeWindow_Type `protobuf_oneof:"type"`
}

func (x *LogsUniqueValueTimeWindow) Reset() {
	*x = LogsUniqueValueTimeWindow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsUniqueValueTimeWindow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsUniqueValueTimeWindow) ProtoMessage() {}

func (x *LogsUniqueValueTimeWindow) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsUniqueValueTimeWindow.ProtoReflect.Descriptor instead.
func (*LogsUniqueValueTimeWindow) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto_rawDescGZIP(), []int{0}
}

func (m *LogsUniqueValueTimeWindow) GetType() isLogsUniqueValueTimeWindow_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *LogsUniqueValueTimeWindow) GetLogsUniqueValueTimeWindowSpecificValue() LogsUniqueValueTimeWindowValue {
	if x, ok := x.GetType().(*LogsUniqueValueTimeWindow_LogsUniqueValueTimeWindowSpecificValue); ok {
		return x.LogsUniqueValueTimeWindowSpecificValue
	}
	return LogsUniqueValueTimeWindowValue_LOGS_UNIQUE_VALUE_TIME_WINDOW_VALUE_MINUTE_1_OR_UNSPECIFIED
}

type isLogsUniqueValueTimeWindow_Type interface {
	isLogsUniqueValueTimeWindow_Type()
}

type LogsUniqueValueTimeWindow_LogsUniqueValueTimeWindowSpecificValue struct {
	LogsUniqueValueTimeWindowSpecificValue LogsUniqueValueTimeWindowValue `protobuf:"varint,1,opt,name=logs_unique_value_time_window_specific_value,json=logsUniqueValueTimeWindowSpecificValue,proto3,enum=com.coralogixapis.alerts.v3.LogsUniqueValueTimeWindowValue,oneof"`
}

func (*LogsUniqueValueTimeWindow_LogsUniqueValueTimeWindowSpecificValue) isLogsUniqueValueTimeWindow_Type() {
}

var File_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto_rawDesc = []byte{
	0x0a, 0x60, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x22,
	0xc1, 0x01, 0x0a, 0x19, 0x4c, 0x6f, 0x67, 0x73, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x9b, 0x01,
	0x0a, 0x2c, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x3b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e,
	0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x48, 0x00, 0x52, 0x26, 0x6c, 0x6f, 0x67, 0x73, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x53, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x2a, 0xa5, 0x04, 0x0a, 0x1e, 0x4c, 0x6f, 0x67, 0x73, 0x55, 0x6e, 0x69, 0x71,
	0x75, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f,
	0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3f, 0x0a, 0x3b, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x55,
	0x4e, 0x49, 0x51, 0x55, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45,
	0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4d, 0x49,
	0x4e, 0x55, 0x54, 0x45, 0x5f, 0x31, 0x5f, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x32, 0x0a, 0x2e, 0x4c, 0x4f, 0x47, 0x53, 0x5f,
	0x55, 0x4e, 0x49, 0x51, 0x55, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54, 0x49, 0x4d,
	0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4d,
	0x49, 0x4e, 0x55, 0x54, 0x45, 0x53, 0x5f, 0x31, 0x35, 0x10, 0x01, 0x12, 0x32, 0x0a, 0x2e, 0x4c,
	0x4f, 0x47, 0x53, 0x5f, 0x55, 0x4e, 0x49, 0x51, 0x55, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45,
	0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c,
	0x55, 0x45, 0x5f, 0x4d, 0x49, 0x4e, 0x55, 0x54, 0x45, 0x53, 0x5f, 0x32, 0x30, 0x10, 0x02, 0x12,
	0x32, 0x0a, 0x2e, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x55, 0x4e, 0x49, 0x51, 0x55, 0x45, 0x5f, 0x56,
	0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57,
	0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4d, 0x49, 0x4e, 0x55, 0x54, 0x45, 0x53, 0x5f, 0x33,
	0x30, 0x10, 0x03, 0x12, 0x2f, 0x0a, 0x2b, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x55, 0x4e, 0x49, 0x51,
	0x55, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49,
	0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x48, 0x4f, 0x55, 0x52, 0x53,
	0x5f, 0x31, 0x10, 0x04, 0x12, 0x2f, 0x0a, 0x2b, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x55, 0x4e, 0x49,
	0x51, 0x55, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x57,
	0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x48, 0x4f, 0x55, 0x52,
	0x53, 0x5f, 0x32, 0x10, 0x05, 0x12, 0x2f, 0x0a, 0x2b, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x55, 0x4e,
	0x49, 0x51, 0x55, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f,
	0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x48, 0x4f, 0x55,
	0x52, 0x53, 0x5f, 0x34, 0x10, 0x06, 0x12, 0x2f, 0x0a, 0x2b, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x55,
	0x4e, 0x49, 0x51, 0x55, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45,
	0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x48, 0x4f,
	0x55, 0x52, 0x53, 0x5f, 0x36, 0x10, 0x07, 0x12, 0x30, 0x0a, 0x2c, 0x4c, 0x4f, 0x47, 0x53, 0x5f,
	0x55, 0x4e, 0x49, 0x51, 0x55, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54, 0x49, 0x4d,
	0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x48,
	0x4f, 0x55, 0x52, 0x53, 0x5f, 0x31, 0x32, 0x10, 0x08, 0x12, 0x30, 0x0a, 0x2c, 0x4c, 0x4f, 0x47,
	0x53, 0x5f, 0x55, 0x4e, 0x49, 0x51, 0x55, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54,
	0x49, 0x4d, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45,
	0x5f, 0x48, 0x4f, 0x55, 0x52, 0x53, 0x5f, 0x32, 0x34, 0x10, 0x09, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto_rawDescData = file_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto_goTypes = []interface{}{
	(LogsUniqueValueTimeWindowValue)(0), // 0: com.coralogixapis.alerts.v3.LogsUniqueValueTimeWindowValue
	(*LogsUniqueValueTimeWindow)(nil),   // 1: com.coralogixapis.alerts.v3.LogsUniqueValueTimeWindow
}
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto_depIdxs = []int32{
	0, // 0: com.coralogixapis.alerts.v3.LogsUniqueValueTimeWindow.logs_unique_value_time_window_specific_value:type_name -> com.coralogixapis.alerts.v3.LogsUniqueValueTimeWindowValue
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto_init()
}
func file_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogsUniqueValueTimeWindow); i {
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
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*LogsUniqueValueTimeWindow_LogsUniqueValueTimeWindowSpecificValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_unique_count_unique_value_timewindow_proto_depIdxs = nil
}
