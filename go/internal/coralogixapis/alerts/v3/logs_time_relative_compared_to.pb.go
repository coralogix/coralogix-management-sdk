// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/logs/time_relative/logs_time_relative_compared_to.proto

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

type LogsTimeRelativeComparedTo int32

const (
	LogsTimeRelativeComparedTo_LOGS_TIME_RELATIVE_COMPARED_TO_PREVIOUS_HOUR_OR_UNSPECIFIED LogsTimeRelativeComparedTo = 0
	LogsTimeRelativeComparedTo_LOGS_TIME_RELATIVE_COMPARED_TO_SAME_HOUR_YESTERDAY          LogsTimeRelativeComparedTo = 1
	LogsTimeRelativeComparedTo_LOGS_TIME_RELATIVE_COMPARED_TO_SAME_HOUR_LAST_WEEK          LogsTimeRelativeComparedTo = 2
	LogsTimeRelativeComparedTo_LOGS_TIME_RELATIVE_COMPARED_TO_YESTERDAY                    LogsTimeRelativeComparedTo = 3
	LogsTimeRelativeComparedTo_LOGS_TIME_RELATIVE_COMPARED_TO_SAME_DAY_LAST_WEEK           LogsTimeRelativeComparedTo = 4
	LogsTimeRelativeComparedTo_LOGS_TIME_RELATIVE_COMPARED_TO_SAME_DAY_LAST_MONTH          LogsTimeRelativeComparedTo = 5
)

// Enum value maps for LogsTimeRelativeComparedTo.
var (
	LogsTimeRelativeComparedTo_name = map[int32]string{
		0: "LOGS_TIME_RELATIVE_COMPARED_TO_PREVIOUS_HOUR_OR_UNSPECIFIED",
		1: "LOGS_TIME_RELATIVE_COMPARED_TO_SAME_HOUR_YESTERDAY",
		2: "LOGS_TIME_RELATIVE_COMPARED_TO_SAME_HOUR_LAST_WEEK",
		3: "LOGS_TIME_RELATIVE_COMPARED_TO_YESTERDAY",
		4: "LOGS_TIME_RELATIVE_COMPARED_TO_SAME_DAY_LAST_WEEK",
		5: "LOGS_TIME_RELATIVE_COMPARED_TO_SAME_DAY_LAST_MONTH",
	}
	LogsTimeRelativeComparedTo_value = map[string]int32{
		"LOGS_TIME_RELATIVE_COMPARED_TO_PREVIOUS_HOUR_OR_UNSPECIFIED": 0,
		"LOGS_TIME_RELATIVE_COMPARED_TO_SAME_HOUR_YESTERDAY":          1,
		"LOGS_TIME_RELATIVE_COMPARED_TO_SAME_HOUR_LAST_WEEK":          2,
		"LOGS_TIME_RELATIVE_COMPARED_TO_YESTERDAY":                    3,
		"LOGS_TIME_RELATIVE_COMPARED_TO_SAME_DAY_LAST_WEEK":           4,
		"LOGS_TIME_RELATIVE_COMPARED_TO_SAME_DAY_LAST_MONTH":          5,
	}
)

func (x LogsTimeRelativeComparedTo) Enum() *LogsTimeRelativeComparedTo {
	p := new(LogsTimeRelativeComparedTo)
	*p = x
	return p
}

func (x LogsTimeRelativeComparedTo) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogsTimeRelativeComparedTo) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_time_relative_logs_time_relative_compared_to_proto_enumTypes[0].Descriptor()
}

func (LogsTimeRelativeComparedTo) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_time_relative_logs_time_relative_compared_to_proto_enumTypes[0]
}

func (x LogsTimeRelativeComparedTo) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogsTimeRelativeComparedTo.Descriptor instead.
func (LogsTimeRelativeComparedTo) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_time_relative_logs_time_relative_compared_to_proto_rawDescGZIP(), []int{0}
}

var File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_time_relative_logs_time_relative_compared_to_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_time_relative_logs_time_relative_compared_to_proto_rawDesc = []byte{
	0x0a, 0x6d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2a, 0xea, 0x02, 0x0a,
	0x1a, 0x4c, 0x6f, 0x67, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x64, 0x54, 0x6f, 0x12, 0x3f, 0x0a, 0x3b, 0x4c,
	0x4f, 0x47, 0x53, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x56,
	0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x52, 0x45, 0x44, 0x5f, 0x54, 0x4f, 0x5f, 0x50, 0x52,
	0x45, 0x56, 0x49, 0x4f, 0x55, 0x53, 0x5f, 0x48, 0x4f, 0x55, 0x52, 0x5f, 0x4f, 0x52, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x36, 0x0a, 0x32,
	0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49,
	0x56, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x52, 0x45, 0x44, 0x5f, 0x54, 0x4f, 0x5f, 0x53,
	0x41, 0x4d, 0x45, 0x5f, 0x48, 0x4f, 0x55, 0x52, 0x5f, 0x59, 0x45, 0x53, 0x54, 0x45, 0x52, 0x44,
	0x41, 0x59, 0x10, 0x01, 0x12, 0x36, 0x0a, 0x32, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x54, 0x49, 0x4d,
	0x45, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41,
	0x52, 0x45, 0x44, 0x5f, 0x54, 0x4f, 0x5f, 0x53, 0x41, 0x4d, 0x45, 0x5f, 0x48, 0x4f, 0x55, 0x52,
	0x5f, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x57, 0x45, 0x45, 0x4b, 0x10, 0x02, 0x12, 0x2c, 0x0a, 0x28,
	0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49,
	0x56, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x52, 0x45, 0x44, 0x5f, 0x54, 0x4f, 0x5f, 0x59,
	0x45, 0x53, 0x54, 0x45, 0x52, 0x44, 0x41, 0x59, 0x10, 0x03, 0x12, 0x35, 0x0a, 0x31, 0x4c, 0x4f,
	0x47, 0x53, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x56, 0x45,
	0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x52, 0x45, 0x44, 0x5f, 0x54, 0x4f, 0x5f, 0x53, 0x41, 0x4d,
	0x45, 0x5f, 0x44, 0x41, 0x59, 0x5f, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x57, 0x45, 0x45, 0x4b, 0x10,
	0x04, 0x12, 0x36, 0x0a, 0x32, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x52,
	0x45, 0x4c, 0x41, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x52, 0x45, 0x44,
	0x5f, 0x54, 0x4f, 0x5f, 0x53, 0x41, 0x4d, 0x45, 0x5f, 0x44, 0x41, 0x59, 0x5f, 0x4c, 0x41, 0x53,
	0x54, 0x5f, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x10, 0x05, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_time_relative_logs_time_relative_compared_to_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_time_relative_logs_time_relative_compared_to_proto_rawDescData = file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_time_relative_logs_time_relative_compared_to_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_time_relative_logs_time_relative_compared_to_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_time_relative_logs_time_relative_compared_to_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_time_relative_logs_time_relative_compared_to_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_time_relative_logs_time_relative_compared_to_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_time_relative_logs_time_relative_compared_to_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_time_relative_logs_time_relative_compared_to_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_time_relative_logs_time_relative_compared_to_proto_goTypes = []any{
	(LogsTimeRelativeComparedTo)(0), // 0: com.coralogixapis.alerts.v3.LogsTimeRelativeComparedTo
}
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_time_relative_logs_time_relative_compared_to_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_time_relative_logs_time_relative_compared_to_proto_init()
}
func file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_time_relative_logs_time_relative_compared_to_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_time_relative_logs_time_relative_compared_to_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_time_relative_logs_time_relative_compared_to_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_time_relative_logs_time_relative_compared_to_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_time_relative_logs_time_relative_compared_to_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_time_relative_logs_time_relative_compared_to_proto_enumTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_time_relative_logs_time_relative_compared_to_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_time_relative_logs_time_relative_compared_to_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_time_relative_logs_time_relative_compared_to_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_time_relative_logs_time_relative_compared_to_proto_depIdxs = nil
}
