// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.26.1
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/metric/threshold/metric_threshold_condition_type.proto

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

type MetricThresholdConditionType int32

const (
	MetricThresholdConditionType_METRIC_THRESHOLD_CONDITION_TYPE_MORE_THAN_OR_UNSPECIFIED MetricThresholdConditionType = 0
	MetricThresholdConditionType_METRIC_THRESHOLD_CONDITION_TYPE_LESS_THAN                MetricThresholdConditionType = 1
	MetricThresholdConditionType_METRIC_THRESHOLD_CONDITION_TYPE_MORE_THAN_OR_EQUALS      MetricThresholdConditionType = 2
	MetricThresholdConditionType_METRIC_THRESHOLD_CONDITION_TYPE_LESS_THAN_OR_EQUALS      MetricThresholdConditionType = 3
)

// Enum value maps for MetricThresholdConditionType.
var (
	MetricThresholdConditionType_name = map[int32]string{
		0: "METRIC_THRESHOLD_CONDITION_TYPE_MORE_THAN_OR_UNSPECIFIED",
		1: "METRIC_THRESHOLD_CONDITION_TYPE_LESS_THAN",
		2: "METRIC_THRESHOLD_CONDITION_TYPE_MORE_THAN_OR_EQUALS",
		3: "METRIC_THRESHOLD_CONDITION_TYPE_LESS_THAN_OR_EQUALS",
	}
	MetricThresholdConditionType_value = map[string]int32{
		"METRIC_THRESHOLD_CONDITION_TYPE_MORE_THAN_OR_UNSPECIFIED": 0,
		"METRIC_THRESHOLD_CONDITION_TYPE_LESS_THAN":                1,
		"METRIC_THRESHOLD_CONDITION_TYPE_MORE_THAN_OR_EQUALS":      2,
		"METRIC_THRESHOLD_CONDITION_TYPE_LESS_THAN_OR_EQUALS":      3,
	}
)

func (x MetricThresholdConditionType) Enum() *MetricThresholdConditionType {
	p := new(MetricThresholdConditionType)
	*p = x
	return p
}

func (x MetricThresholdConditionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetricThresholdConditionType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_condition_type_proto_enumTypes[0].Descriptor()
}

func (MetricThresholdConditionType) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_condition_type_proto_enumTypes[0]
}

func (x MetricThresholdConditionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetricThresholdConditionType.Descriptor instead.
func (MetricThresholdConditionType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_condition_type_proto_rawDescGZIP(), []int{0}
}

var File_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_condition_type_proto protoreflect.FileDescriptor

const file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_condition_type_proto_rawDesc = "" +
	"\n" +
	"lcom/coralogixapis/alerts/v3/alert_def_type_definition/metric/threshold/metric_threshold_condition_type.proto\x12\x1bcom.coralogixapis.alerts.v3*\xfd\x01\n" +
	"\x1cMetricThresholdConditionType\x12<\n" +
	"8METRIC_THRESHOLD_CONDITION_TYPE_MORE_THAN_OR_UNSPECIFIED\x10\x00\x12-\n" +
	")METRIC_THRESHOLD_CONDITION_TYPE_LESS_THAN\x10\x01\x127\n" +
	"3METRIC_THRESHOLD_CONDITION_TYPE_MORE_THAN_OR_EQUALS\x10\x02\x127\n" +
	"3METRIC_THRESHOLD_CONDITION_TYPE_LESS_THAN_OR_EQUALS\x10\x03b\x06proto3"

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_condition_type_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_condition_type_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_condition_type_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_condition_type_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_condition_type_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_condition_type_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_condition_type_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_condition_type_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_condition_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_condition_type_proto_goTypes = []any{
	(MetricThresholdConditionType)(0), // 0: com.coralogixapis.alerts.v3.MetricThresholdConditionType
}
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_condition_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_condition_type_proto_init()
}
func file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_condition_type_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_condition_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_condition_type_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_condition_type_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_condition_type_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_condition_type_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_condition_type_proto_enumTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_condition_type_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_condition_type_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_condition_type_proto_depIdxs = nil
}
