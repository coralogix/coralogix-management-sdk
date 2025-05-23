// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/logs/anomaly/logs_anomaly_condition_type.proto

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

type LogsAnomalyConditionType int32

const (
	LogsAnomalyConditionType_LOGS_ANOMALY_CONDITION_TYPE_MORE_THAN_USUAL_OR_UNSPECIFIED LogsAnomalyConditionType = 0
)

// Enum value maps for LogsAnomalyConditionType.
var (
	LogsAnomalyConditionType_name = map[int32]string{
		0: "LOGS_ANOMALY_CONDITION_TYPE_MORE_THAN_USUAL_OR_UNSPECIFIED",
	}
	LogsAnomalyConditionType_value = map[string]int32{
		"LOGS_ANOMALY_CONDITION_TYPE_MORE_THAN_USUAL_OR_UNSPECIFIED": 0,
	}
)

func (x LogsAnomalyConditionType) Enum() *LogsAnomalyConditionType {
	p := new(LogsAnomalyConditionType)
	*p = x
	return p
}

func (x LogsAnomalyConditionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogsAnomalyConditionType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_anomaly_logs_anomaly_condition_type_proto_enumTypes[0].Descriptor()
}

func (LogsAnomalyConditionType) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_anomaly_logs_anomaly_condition_type_proto_enumTypes[0]
}

func (x LogsAnomalyConditionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogsAnomalyConditionType.Descriptor instead.
func (LogsAnomalyConditionType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_anomaly_logs_anomaly_condition_type_proto_rawDescGZIP(), []int{0}
}

var File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_anomaly_logs_anomaly_condition_type_proto protoreflect.FileDescriptor

const file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_anomaly_logs_anomaly_condition_type_proto_rawDesc = "" +
	"\n" +
	"dcom/coralogixapis/alerts/v3/alert_def_type_definition/logs/anomaly/logs_anomaly_condition_type.proto\x12\x1bcom.coralogixapis.alerts.v3*Z\n" +
	"\x18LogsAnomalyConditionType\x12>\n" +
	":LOGS_ANOMALY_CONDITION_TYPE_MORE_THAN_USUAL_OR_UNSPECIFIED\x10\x00b\x06proto3"

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_anomaly_logs_anomaly_condition_type_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_anomaly_logs_anomaly_condition_type_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_anomaly_logs_anomaly_condition_type_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_anomaly_logs_anomaly_condition_type_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_anomaly_logs_anomaly_condition_type_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_anomaly_logs_anomaly_condition_type_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_anomaly_logs_anomaly_condition_type_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_anomaly_logs_anomaly_condition_type_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_anomaly_logs_anomaly_condition_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_anomaly_logs_anomaly_condition_type_proto_goTypes = []any{
	(LogsAnomalyConditionType)(0), // 0: com.coralogixapis.alerts.v3.LogsAnomalyConditionType
}
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_anomaly_logs_anomaly_condition_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_anomaly_logs_anomaly_condition_type_proto_init()
}
func file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_anomaly_logs_anomaly_condition_type_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_anomaly_logs_anomaly_condition_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_anomaly_logs_anomaly_condition_type_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_anomaly_logs_anomaly_condition_type_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_anomaly_logs_anomaly_condition_type_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_anomaly_logs_anomaly_condition_type_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_anomaly_logs_anomaly_condition_type_proto_enumTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_anomaly_logs_anomaly_condition_type_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_anomaly_logs_anomaly_condition_type_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_anomaly_logs_anomaly_condition_type_proto_depIdxs = nil
}
