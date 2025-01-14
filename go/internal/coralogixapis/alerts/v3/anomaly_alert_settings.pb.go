// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/common/anomaly_alert_settings.proto

package v3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AnomalyAlertSettings struct {
	state                 protoimpl.MessageState `protogen:"open.v1"`
	PercentageOfDeviation *wrapperspb.FloatValue `protobuf:"bytes,1,opt,name=percentage_of_deviation,json=percentageOfDeviation,proto3" json:"percentage_of_deviation,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *AnomalyAlertSettings) Reset() {
	*x = AnomalyAlertSettings{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_common_anomaly_alert_settings_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnomalyAlertSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnomalyAlertSettings) ProtoMessage() {}

func (x *AnomalyAlertSettings) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_common_anomaly_alert_settings_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnomalyAlertSettings.ProtoReflect.Descriptor instead.
func (*AnomalyAlertSettings) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_common_anomaly_alert_settings_proto_rawDescGZIP(), []int{0}
}

func (x *AnomalyAlertSettings) GetPercentageOfDeviation() *wrapperspb.FloatValue {
	if x != nil {
		return x.PercentageOfDeviation
	}
	return nil
}

var File_com_coralogixapis_alerts_v3_alert_def_type_definition_common_anomaly_alert_settings_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_common_anomaly_alert_settings_proto_rawDesc = []byte{
	0x0a, 0x59, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61,
	0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x5f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x14, 0x41, 0x6e, 0x6f, 0x6d,
	0x61, 0x6c, 0x79, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x12, 0x53, 0x0a, 0x17, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x6f,
	0x66, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x15,
	0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x4f, 0x66, 0x44, 0x65, 0x76, 0x69,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_common_anomaly_alert_settings_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_common_anomaly_alert_settings_proto_rawDescData = file_com_coralogixapis_alerts_v3_alert_def_type_definition_common_anomaly_alert_settings_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_common_anomaly_alert_settings_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_common_anomaly_alert_settings_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_common_anomaly_alert_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_alert_def_type_definition_common_anomaly_alert_settings_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_common_anomaly_alert_settings_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_common_anomaly_alert_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_common_anomaly_alert_settings_proto_goTypes = []any{
	(*AnomalyAlertSettings)(nil),  // 0: com.coralogixapis.alerts.v3.AnomalyAlertSettings
	(*wrapperspb.FloatValue)(nil), // 1: google.protobuf.FloatValue
}
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_common_anomaly_alert_settings_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.alerts.v3.AnomalyAlertSettings.percentage_of_deviation:type_name -> google.protobuf.FloatValue
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_common_anomaly_alert_settings_proto_init()
}
func file_com_coralogixapis_alerts_v3_alert_def_type_definition_common_anomaly_alert_settings_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_type_definition_common_anomaly_alert_settings_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_alert_def_type_definition_common_anomaly_alert_settings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_type_definition_common_anomaly_alert_settings_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_type_definition_common_anomaly_alert_settings_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_alert_def_type_definition_common_anomaly_alert_settings_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_type_definition_common_anomaly_alert_settings_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_common_anomaly_alert_settings_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_common_anomaly_alert_settings_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_common_anomaly_alert_settings_proto_depIdxs = nil
}
