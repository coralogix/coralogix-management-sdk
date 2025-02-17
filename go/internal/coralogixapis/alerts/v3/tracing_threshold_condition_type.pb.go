// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.26.1
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/tracing/threshold/tracing_threshold_condition_type.proto

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

type TracingThresholdConditionType int32

const (
	TracingThresholdConditionType_TRACING_THRESHOLD_CONDITION_TYPE_MORE_THAN_OR_UNSPECIFIED TracingThresholdConditionType = 0
)

// Enum value maps for TracingThresholdConditionType.
var (
	TracingThresholdConditionType_name = map[int32]string{
		0: "TRACING_THRESHOLD_CONDITION_TYPE_MORE_THAN_OR_UNSPECIFIED",
	}
	TracingThresholdConditionType_value = map[string]int32{
		"TRACING_THRESHOLD_CONDITION_TYPE_MORE_THAN_OR_UNSPECIFIED": 0,
	}
)

func (x TracingThresholdConditionType) Enum() *TracingThresholdConditionType {
	p := new(TracingThresholdConditionType)
	*p = x
	return p
}

func (x TracingThresholdConditionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TracingThresholdConditionType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_condition_type_proto_enumTypes[0].Descriptor()
}

func (TracingThresholdConditionType) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_condition_type_proto_enumTypes[0]
}

func (x TracingThresholdConditionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TracingThresholdConditionType.Descriptor instead.
func (TracingThresholdConditionType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_condition_type_proto_rawDescGZIP(), []int{0}
}

var File_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_condition_type_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_condition_type_proto_rawDesc = []byte{
	0x0a, 0x6e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2f,
	0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e,
	0x67, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2a, 0x5e, 0x0a,
	0x1d, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3d,
	0x0a, 0x39, 0x54, 0x52, 0x41, 0x43, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x48, 0x52, 0x45, 0x53, 0x48,
	0x4f, 0x4c, 0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4d, 0x4f, 0x52, 0x45, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x5f, 0x4f, 0x52, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_condition_type_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_condition_type_proto_rawDescData = file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_condition_type_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_condition_type_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_condition_type_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_condition_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_condition_type_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_condition_type_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_condition_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_condition_type_proto_goTypes = []any{
	(TracingThresholdConditionType)(0), // 0: com.coralogixapis.alerts.v3.TracingThresholdConditionType
}
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_condition_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_condition_type_proto_init()
}
func file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_condition_type_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_condition_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_condition_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_condition_type_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_condition_type_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_condition_type_proto_enumTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_condition_type_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_condition_type_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_condition_type_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_threshold_tracing_threshold_condition_type_proto_depIdxs = nil
}
