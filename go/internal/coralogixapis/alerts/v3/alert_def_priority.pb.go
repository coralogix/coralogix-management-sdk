// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: com/coralogixapis/alerts/v3/alert_def_priority.proto

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

type AlertDefPriority int32

const (
	AlertDefPriority_ALERT_DEF_PRIORITY_P5_OR_UNSPECIFIED AlertDefPriority = 0
	AlertDefPriority_ALERT_DEF_PRIORITY_P4                AlertDefPriority = 1
	AlertDefPriority_ALERT_DEF_PRIORITY_P3                AlertDefPriority = 2
	AlertDefPriority_ALERT_DEF_PRIORITY_P2                AlertDefPriority = 3
	AlertDefPriority_ALERT_DEF_PRIORITY_P1                AlertDefPriority = 4
)

// Enum value maps for AlertDefPriority.
var (
	AlertDefPriority_name = map[int32]string{
		0: "ALERT_DEF_PRIORITY_P5_OR_UNSPECIFIED",
		1: "ALERT_DEF_PRIORITY_P4",
		2: "ALERT_DEF_PRIORITY_P3",
		3: "ALERT_DEF_PRIORITY_P2",
		4: "ALERT_DEF_PRIORITY_P1",
	}
	AlertDefPriority_value = map[string]int32{
		"ALERT_DEF_PRIORITY_P5_OR_UNSPECIFIED": 0,
		"ALERT_DEF_PRIORITY_P4":                1,
		"ALERT_DEF_PRIORITY_P3":                2,
		"ALERT_DEF_PRIORITY_P2":                3,
		"ALERT_DEF_PRIORITY_P1":                4,
	}
)

func (x AlertDefPriority) Enum() *AlertDefPriority {
	p := new(AlertDefPriority)
	*p = x
	return p
}

func (x AlertDefPriority) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AlertDefPriority) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_alerts_v3_alert_def_priority_proto_enumTypes[0].Descriptor()
}

func (AlertDefPriority) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_alerts_v3_alert_def_priority_proto_enumTypes[0]
}

func (x AlertDefPriority) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AlertDefPriority.Descriptor instead.
func (AlertDefPriority) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_priority_proto_rawDescGZIP(), []int{0}
}

var File_com_coralogixapis_alerts_v3_alert_def_priority_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_alert_def_priority_proto_rawDesc = []byte{
	0x0a, 0x34, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x2e, 0x76, 0x33, 0x2a, 0xa8, 0x01, 0x0a, 0x10, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x44, 0x65, 0x66,
	0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x28, 0x0a, 0x24, 0x41, 0x4c, 0x45, 0x52,
	0x54, 0x5f, 0x44, 0x45, 0x46, 0x5f, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x50,
	0x35, 0x5f, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x5f, 0x44, 0x45, 0x46, 0x5f,
	0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x50, 0x34, 0x10, 0x01, 0x12, 0x19, 0x0a,
	0x15, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x5f, 0x44, 0x45, 0x46, 0x5f, 0x50, 0x52, 0x49, 0x4f, 0x52,
	0x49, 0x54, 0x59, 0x5f, 0x50, 0x33, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x4c, 0x45, 0x52,
	0x54, 0x5f, 0x44, 0x45, 0x46, 0x5f, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x50,
	0x32, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x5f, 0x44, 0x45, 0x46,
	0x5f, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x50, 0x31, 0x10, 0x04, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_alert_def_priority_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_priority_proto_rawDescData = file_com_coralogixapis_alerts_v3_alert_def_priority_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_alert_def_priority_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_priority_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_priority_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_alert_def_priority_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_alert_def_priority_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_priority_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_alerts_v3_alert_def_priority_proto_goTypes = []any{
	(AlertDefPriority)(0), // 0: com.coralogixapis.alerts.v3.AlertDefPriority
}
var file_com_coralogixapis_alerts_v3_alert_def_priority_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_alerts_v3_alert_def_priority_proto_init() }
func file_com_coralogixapis_alerts_v3_alert_def_priority_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_priority_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_alert_def_priority_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_priority_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_priority_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_alerts_v3_alert_def_priority_proto_enumTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_priority_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_priority_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_alert_def_priority_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_priority_proto_depIdxs = nil
}
