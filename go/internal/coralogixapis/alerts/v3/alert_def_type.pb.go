// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: com/coralogixapis/alerts/v3/alert_def_type.proto

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

type AlertDefType int32

const (
	AlertDefType_ALERT_DEF_TYPE_LOGS_IMMEDIATE_OR_UNSPECIFIED AlertDefType = 0
	AlertDefType_ALERT_DEF_TYPE_LOGS_THRESHOLD                AlertDefType = 1
	AlertDefType_ALERT_DEF_TYPE_LOGS_ANOMALY                  AlertDefType = 3
	AlertDefType_ALERT_DEF_TYPE_LOGS_RATIO_THRESHOLD          AlertDefType = 4
	AlertDefType_ALERT_DEF_TYPE_LOGS_NEW_VALUE                AlertDefType = 6
	AlertDefType_ALERT_DEF_TYPE_LOGS_UNIQUE_COUNT             AlertDefType = 7
	AlertDefType_ALERT_DEF_TYPE_LOGS_TIME_RELATIVE_THRESHOLD  AlertDefType = 8
	AlertDefType_ALERT_DEF_TYPE_METRIC_THRESHOLD              AlertDefType = 10
	AlertDefType_ALERT_DEF_TYPE_METRIC_ANOMALY                AlertDefType = 14
	AlertDefType_ALERT_DEF_TYPE_TRACING_IMMEDIATE             AlertDefType = 15
	AlertDefType_ALERT_DEF_TYPE_TRACING_THRESHOLD             AlertDefType = 16
	AlertDefType_ALERT_DEF_TYPE_FLOW                          AlertDefType = 17
)

// Enum value maps for AlertDefType.
var (
	AlertDefType_name = map[int32]string{
		0:  "ALERT_DEF_TYPE_LOGS_IMMEDIATE_OR_UNSPECIFIED",
		1:  "ALERT_DEF_TYPE_LOGS_THRESHOLD",
		3:  "ALERT_DEF_TYPE_LOGS_ANOMALY",
		4:  "ALERT_DEF_TYPE_LOGS_RATIO_THRESHOLD",
		6:  "ALERT_DEF_TYPE_LOGS_NEW_VALUE",
		7:  "ALERT_DEF_TYPE_LOGS_UNIQUE_COUNT",
		8:  "ALERT_DEF_TYPE_LOGS_TIME_RELATIVE_THRESHOLD",
		10: "ALERT_DEF_TYPE_METRIC_THRESHOLD",
		14: "ALERT_DEF_TYPE_METRIC_ANOMALY",
		15: "ALERT_DEF_TYPE_TRACING_IMMEDIATE",
		16: "ALERT_DEF_TYPE_TRACING_THRESHOLD",
		17: "ALERT_DEF_TYPE_FLOW",
	}
	AlertDefType_value = map[string]int32{
		"ALERT_DEF_TYPE_LOGS_IMMEDIATE_OR_UNSPECIFIED": 0,
		"ALERT_DEF_TYPE_LOGS_THRESHOLD":                1,
		"ALERT_DEF_TYPE_LOGS_ANOMALY":                  3,
		"ALERT_DEF_TYPE_LOGS_RATIO_THRESHOLD":          4,
		"ALERT_DEF_TYPE_LOGS_NEW_VALUE":                6,
		"ALERT_DEF_TYPE_LOGS_UNIQUE_COUNT":             7,
		"ALERT_DEF_TYPE_LOGS_TIME_RELATIVE_THRESHOLD":  8,
		"ALERT_DEF_TYPE_METRIC_THRESHOLD":              10,
		"ALERT_DEF_TYPE_METRIC_ANOMALY":                14,
		"ALERT_DEF_TYPE_TRACING_IMMEDIATE":             15,
		"ALERT_DEF_TYPE_TRACING_THRESHOLD":             16,
		"ALERT_DEF_TYPE_FLOW":                          17,
	}
)

func (x AlertDefType) Enum() *AlertDefType {
	p := new(AlertDefType)
	*p = x
	return p
}

func (x AlertDefType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AlertDefType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_alerts_v3_alert_def_type_proto_enumTypes[0].Descriptor()
}

func (AlertDefType) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_alerts_v3_alert_def_type_proto_enumTypes[0]
}

func (x AlertDefType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AlertDefType.Descriptor instead.
func (AlertDefType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_proto_rawDescGZIP(), []int{0}
}

var File_com_coralogixapis_alerts_v3_alert_def_type_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_alert_def_type_proto_rawDesc = []byte{
	0x0a, 0x30, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2a,
	0xd4, 0x03, 0x0a, 0x0c, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x44, 0x65, 0x66, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x30, 0x0a, 0x2c, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x5f, 0x44, 0x45, 0x46, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x49, 0x4d, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x54,
	0x45, 0x5f, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x5f, 0x44, 0x45, 0x46, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x54, 0x48, 0x52, 0x45, 0x53, 0x48,
	0x4f, 0x4c, 0x44, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x5f, 0x44,
	0x45, 0x46, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x41, 0x4e, 0x4f,
	0x4d, 0x41, 0x4c, 0x59, 0x10, 0x03, 0x12, 0x27, 0x0a, 0x23, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x5f,
	0x44, 0x45, 0x46, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x52, 0x41,
	0x54, 0x49, 0x4f, 0x5f, 0x54, 0x48, 0x52, 0x45, 0x53, 0x48, 0x4f, 0x4c, 0x44, 0x10, 0x04, 0x12,
	0x21, 0x0a, 0x1d, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x5f, 0x44, 0x45, 0x46, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x4e, 0x45, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45,
	0x10, 0x06, 0x12, 0x24, 0x0a, 0x20, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x5f, 0x44, 0x45, 0x46, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x55, 0x4e, 0x49, 0x51, 0x55, 0x45,
	0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x07, 0x12, 0x2f, 0x0a, 0x2b, 0x41, 0x4c, 0x45, 0x52,
	0x54, 0x5f, 0x44, 0x45, 0x46, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x4f, 0x47, 0x53, 0x5f,
	0x54, 0x49, 0x4d, 0x45, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x48,
	0x52, 0x45, 0x53, 0x48, 0x4f, 0x4c, 0x44, 0x10, 0x08, 0x12, 0x23, 0x0a, 0x1f, 0x41, 0x4c, 0x45,
	0x52, 0x54, 0x5f, 0x44, 0x45, 0x46, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x45, 0x54, 0x52,
	0x49, 0x43, 0x5f, 0x54, 0x48, 0x52, 0x45, 0x53, 0x48, 0x4f, 0x4c, 0x44, 0x10, 0x0a, 0x12, 0x21,
	0x0a, 0x1d, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x5f, 0x44, 0x45, 0x46, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x41, 0x4e, 0x4f, 0x4d, 0x41, 0x4c, 0x59, 0x10,
	0x0e, 0x12, 0x24, 0x0a, 0x20, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x5f, 0x44, 0x45, 0x46, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x49, 0x4e, 0x47, 0x5f, 0x49, 0x4d, 0x4d, 0x45,
	0x44, 0x49, 0x41, 0x54, 0x45, 0x10, 0x0f, 0x12, 0x24, 0x0a, 0x20, 0x41, 0x4c, 0x45, 0x52, 0x54,
	0x5f, 0x44, 0x45, 0x46, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x49, 0x4e,
	0x47, 0x5f, 0x54, 0x48, 0x52, 0x45, 0x53, 0x48, 0x4f, 0x4c, 0x44, 0x10, 0x10, 0x12, 0x17, 0x0a,
	0x13, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x5f, 0x44, 0x45, 0x46, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x46, 0x4c, 0x4f, 0x57, 0x10, 0x11, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_proto_rawDescData = file_com_coralogixapis_alerts_v3_alert_def_type_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_alert_def_type_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_alert_def_type_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_alerts_v3_alert_def_type_proto_goTypes = []any{
	(AlertDefType)(0), // 0: com.coralogixapis.alerts.v3.AlertDefType
}
var file_com_coralogixapis_alerts_v3_alert_def_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_alerts_v3_alert_def_type_proto_init() }
func file_com_coralogixapis_alerts_v3_alert_def_type_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_alert_def_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_type_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_type_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_alerts_v3_alert_def_type_proto_enumTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_type_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_type_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_proto_depIdxs = nil
}
