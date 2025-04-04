// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/activity_analysis/activity_analysis.proto

package v3

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type ActivityAnalysisStatus int32

const (
	ActivityAnalysisStatus_ACTIVITY_ANALYSIS_STATUS_ACTIVATE_OR_UNSPECIFIED ActivityAnalysisStatus = 0
	ActivityAnalysisStatus_ACTIVITY_ANALYSIS_STATUS_MUTE                    ActivityAnalysisStatus = 1
)

// Enum value maps for ActivityAnalysisStatus.
var (
	ActivityAnalysisStatus_name = map[int32]string{
		0: "ACTIVITY_ANALYSIS_STATUS_ACTIVATE_OR_UNSPECIFIED",
		1: "ACTIVITY_ANALYSIS_STATUS_MUTE",
	}
	ActivityAnalysisStatus_value = map[string]int32{
		"ACTIVITY_ANALYSIS_STATUS_ACTIVATE_OR_UNSPECIFIED": 0,
		"ACTIVITY_ANALYSIS_STATUS_MUTE":                    1,
	}
)

func (x ActivityAnalysisStatus) Enum() *ActivityAnalysisStatus {
	p := new(ActivityAnalysisStatus)
	*p = x
	return p
}

func (x ActivityAnalysisStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActivityAnalysisStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto_enumTypes[0].Descriptor()
}

func (ActivityAnalysisStatus) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto_enumTypes[0]
}

func (x ActivityAnalysisStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActivityAnalysisStatus.Descriptor instead.
func (ActivityAnalysisStatus) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto_rawDescGZIP(), []int{0}
}

type ActivityAnalysis struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Rules         []string               `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	Status        ActivityAnalysisStatus `protobuf:"varint,2,opt,name=status,proto3,enum=com.coralogixapis.alerts.v3.ActivityAnalysisStatus" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ActivityAnalysis) Reset() {
	*x = ActivityAnalysis{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActivityAnalysis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityAnalysis) ProtoMessage() {}

func (x *ActivityAnalysis) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityAnalysis.ProtoReflect.Descriptor instead.
func (*ActivityAnalysis) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto_rawDescGZIP(), []int{0}
}

func (x *ActivityAnalysis) GetRules() []string {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *ActivityAnalysis) GetStatus() ActivityAnalysisStatus {
	if x != nil {
		return x.Status
	}
	return ActivityAnalysisStatus_ACTIVITY_ANALYSIS_STATUS_ACTIVATE_OR_UNSPECIFIED
}

var File_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto_rawDesc = []byte{
	0x0a, 0x5f, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92,
	0x02, 0x0a, 0x10, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x41, 0x6e, 0x61, 0x6c, 0x79,
	0x73, 0x69, 0x73, 0x12, 0x2d, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x42, 0x17, 0x92, 0x41, 0x14, 0x4a, 0x12, 0x5b, 0x22, 0x72, 0x75, 0x6c, 0x65, 0x31,
	0x22, 0x2c, 0x20, 0x22, 0x72, 0x75, 0x6c, 0x65, 0x32, 0x22, 0x5d, 0x52, 0x05, 0x72, 0x75, 0x6c,
	0x65, 0x73, 0x12, 0x4b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69,
	0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a,
	0x81, 0x01, 0x92, 0x41, 0x7e, 0x0a, 0x7c, 0x2a, 0x20, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x20, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x32, 0x47, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x20, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x20, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x73, 0x69, 0x73, 0x2c, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0x20, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0xd2, 0x01, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0xd2, 0x01, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2a, 0x71, 0x0a, 0x16, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x41,
	0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x34, 0x0a,
	0x30, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x41, 0x4e, 0x41, 0x4c, 0x59, 0x53,
	0x49, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x41,
	0x54, 0x45, 0x5f, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f,
	0x41, 0x4e, 0x41, 0x4c, 0x59, 0x53, 0x49, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x4d, 0x55, 0x54, 0x45, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto_rawDescData = file_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto_goTypes = []any{
	(ActivityAnalysisStatus)(0), // 0: com.coralogixapis.alerts.v3.ActivityAnalysisStatus
	(*ActivityAnalysis)(nil),    // 1: com.coralogixapis.alerts.v3.ActivityAnalysis
}
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto_depIdxs = []int32{
	0, // 0: com.coralogixapis.alerts.v3.ActivityAnalysis.status:type_name -> com.coralogixapis.alerts.v3.ActivityAnalysisStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto_init()
}
func file_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_activity_analysis_activity_analysis_proto_depIdxs = nil
}
