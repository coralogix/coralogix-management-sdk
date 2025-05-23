// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/logs/new_value/logs_new_value_timewindow.proto

package v3

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Type:
	//
	//	*LogsNewValueTimeWindow_LogsNewValueTimeWindowSpecificValue
	Type          isLogsNewValueTimeWindow_Type `protobuf_oneof:"type"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogsNewValueTimeWindow) Reset() {
	*x = LogsNewValueTimeWindow{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsNewValueTimeWindow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsNewValueTimeWindow) ProtoMessage() {}

func (x *LogsNewValueTimeWindow) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_msgTypes[0]
	if x != nil {
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

func (x *LogsNewValueTimeWindow) GetType() isLogsNewValueTimeWindow_Type {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *LogsNewValueTimeWindow) GetLogsNewValueTimeWindowSpecificValue() LogsNewValueTimeWindowValue {
	if x != nil {
		if x, ok := x.Type.(*LogsNewValueTimeWindow_LogsNewValueTimeWindowSpecificValue); ok {
			return x.LogsNewValueTimeWindowSpecificValue
		}
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

const file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_rawDesc = "" +
	"\n" +
	"dcom/coralogixapis/alerts/v3/alert_def_type_definition/logs/new_value/logs_new_value_timewindow.proto\x12\x1bcom.coralogixapis.alerts.v3\x1a.protoc-gen-openapiv2/options/annotations.proto\"\x90\x03\n" +
	"\x16LogsNewValueTimeWindow\x12\xfe\x01\n" +
	")logs_new_value_time_window_specific_value\x18\x01 \x01(\x0e28.com.coralogixapis.alerts.v3.LogsNewValueTimeWindowValueBj\x92Ag2)A time window defined by a specific valueJ:\"LOGS_NEW_VALUE_TIME_WINDOW_VALUE_HOURS_12_OR_UNSPECIFIED\"H\x00R#logsNewValueTimeWindowSpecificValue:m\x92Aj\n" +
	"h*%Log-based new value alert time window28Time window configuration for log-based new value alerts\xd2\x01\x04typeB\x06\n" +
	"\x04type*\xa1\x03\n" +
	"\x1bLogsNewValueTimeWindowValue\x12<\n" +
	"8LOGS_NEW_VALUE_TIME_WINDOW_VALUE_HOURS_12_OR_UNSPECIFIED\x10\x00\x12-\n" +
	")LOGS_NEW_VALUE_TIME_WINDOW_VALUE_HOURS_24\x10\x01\x12-\n" +
	")LOGS_NEW_VALUE_TIME_WINDOW_VALUE_HOURS_48\x10\x02\x12-\n" +
	")LOGS_NEW_VALUE_TIME_WINDOW_VALUE_HOURS_72\x10\x03\x12+\n" +
	"'LOGS_NEW_VALUE_TIME_WINDOW_VALUE_WEEK_1\x10\x04\x12,\n" +
	"(LOGS_NEW_VALUE_TIME_WINDOW_VALUE_MONTH_1\x10\x05\x12-\n" +
	")LOGS_NEW_VALUE_TIME_WINDOW_VALUE_MONTHS_2\x10\x06\x12-\n" +
	")LOGS_NEW_VALUE_TIME_WINDOW_VALUE_MONTHS_3\x10\ab\x06proto3"

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_goTypes = []any{
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
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_msgTypes[0].OneofWrappers = []any{
		(*LogsNewValueTimeWindow_LogsNewValueTimeWindowSpecificValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_rawDesc)),
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
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_new_value_logs_new_value_timewindow_proto_depIdxs = nil
}
