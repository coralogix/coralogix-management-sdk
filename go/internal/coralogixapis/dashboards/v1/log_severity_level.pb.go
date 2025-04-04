// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: com/coralogixapis/dashboards/v1/common/log_severity_level.proto

package v1

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

type LogSeverityLevel int32

const (
	LogSeverityLevel_LOG_SEVERITY_LEVEL_UNSPECIFIED LogSeverityLevel = 0
	LogSeverityLevel_LOG_SEVERITY_LEVEL_DEBUG       LogSeverityLevel = 1
	LogSeverityLevel_LOG_SEVERITY_LEVEL_VERBOSE     LogSeverityLevel = 2
	LogSeverityLevel_LOG_SEVERITY_LEVEL_INFO        LogSeverityLevel = 3
	LogSeverityLevel_LOG_SEVERITY_LEVEL_WARNING     LogSeverityLevel = 4
	LogSeverityLevel_LOG_SEVERITY_LEVEL_ERROR       LogSeverityLevel = 5
	LogSeverityLevel_LOG_SEVERITY_LEVEL_CRITICAL    LogSeverityLevel = 6
)

// Enum value maps for LogSeverityLevel.
var (
	LogSeverityLevel_name = map[int32]string{
		0: "LOG_SEVERITY_LEVEL_UNSPECIFIED",
		1: "LOG_SEVERITY_LEVEL_DEBUG",
		2: "LOG_SEVERITY_LEVEL_VERBOSE",
		3: "LOG_SEVERITY_LEVEL_INFO",
		4: "LOG_SEVERITY_LEVEL_WARNING",
		5: "LOG_SEVERITY_LEVEL_ERROR",
		6: "LOG_SEVERITY_LEVEL_CRITICAL",
	}
	LogSeverityLevel_value = map[string]int32{
		"LOG_SEVERITY_LEVEL_UNSPECIFIED": 0,
		"LOG_SEVERITY_LEVEL_DEBUG":       1,
		"LOG_SEVERITY_LEVEL_VERBOSE":     2,
		"LOG_SEVERITY_LEVEL_INFO":        3,
		"LOG_SEVERITY_LEVEL_WARNING":     4,
		"LOG_SEVERITY_LEVEL_ERROR":       5,
		"LOG_SEVERITY_LEVEL_CRITICAL":    6,
	}
)

func (x LogSeverityLevel) Enum() *LogSeverityLevel {
	p := new(LogSeverityLevel)
	*p = x
	return p
}

func (x LogSeverityLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogSeverityLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_dashboards_v1_common_log_severity_level_proto_enumTypes[0].Descriptor()
}

func (LogSeverityLevel) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_dashboards_v1_common_log_severity_level_proto_enumTypes[0]
}

func (x LogSeverityLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogSeverityLevel.Descriptor instead.
func (LogSeverityLevel) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_log_severity_level_proto_rawDescGZIP(), []int{0}
}

var File_com_coralogixapis_dashboards_v1_common_log_severity_level_proto protoreflect.FileDescriptor

var file_com_coralogixapis_dashboards_v1_common_log_severity_level_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65, 0x76,
	0x65, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2a, 0xf0, 0x01, 0x0a, 0x10, 0x4c, 0x6f,
	0x67, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x22,
	0x0a, 0x1e, 0x4c, 0x4f, 0x47, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4c,
	0x45, 0x56, 0x45, 0x4c, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x4c, 0x4f, 0x47, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49,
	0x54, 0x59, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x01,
	0x12, 0x1e, 0x0a, 0x1a, 0x4c, 0x4f, 0x47, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59,
	0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x56, 0x45, 0x52, 0x42, 0x4f, 0x53, 0x45, 0x10, 0x02,
	0x12, 0x1b, 0x0a, 0x17, 0x4c, 0x4f, 0x47, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59,
	0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x03, 0x12, 0x1e, 0x0a,
	0x1a, 0x4c, 0x4f, 0x47, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4c, 0x45,
	0x56, 0x45, 0x4c, 0x5f, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x1c, 0x0a,
	0x18, 0x4c, 0x4f, 0x47, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4c, 0x45,
	0x56, 0x45, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x05, 0x12, 0x1f, 0x0a, 0x1b, 0x4c,
	0x4f, 0x47, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4c, 0x45, 0x56, 0x45,
	0x4c, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x06, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_dashboards_v1_common_log_severity_level_proto_rawDescOnce sync.Once
	file_com_coralogixapis_dashboards_v1_common_log_severity_level_proto_rawDescData = file_com_coralogixapis_dashboards_v1_common_log_severity_level_proto_rawDesc
)

func file_com_coralogixapis_dashboards_v1_common_log_severity_level_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_dashboards_v1_common_log_severity_level_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_dashboards_v1_common_log_severity_level_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_dashboards_v1_common_log_severity_level_proto_rawDescData)
	})
	return file_com_coralogixapis_dashboards_v1_common_log_severity_level_proto_rawDescData
}

var file_com_coralogixapis_dashboards_v1_common_log_severity_level_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_dashboards_v1_common_log_severity_level_proto_goTypes = []any{
	(LogSeverityLevel)(0), // 0: com.coralogixapis.dashboards.v1.common.LogSeverityLevel
}
var file_com_coralogixapis_dashboards_v1_common_log_severity_level_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_dashboards_v1_common_log_severity_level_proto_init() }
func file_com_coralogixapis_dashboards_v1_common_log_severity_level_proto_init() {
	if File_com_coralogixapis_dashboards_v1_common_log_severity_level_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_dashboards_v1_common_log_severity_level_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_dashboards_v1_common_log_severity_level_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_dashboards_v1_common_log_severity_level_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_dashboards_v1_common_log_severity_level_proto_enumTypes,
	}.Build()
	File_com_coralogixapis_dashboards_v1_common_log_severity_level_proto = out.File
	file_com_coralogixapis_dashboards_v1_common_log_severity_level_proto_rawDesc = nil
	file_com_coralogixapis_dashboards_v1_common_log_severity_level_proto_goTypes = nil
	file_com_coralogixapis_dashboards_v1_common_log_severity_level_proto_depIdxs = nil
}
