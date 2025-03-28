// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: com/coralogix/datausage/models/v1/units_ratios.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GbUnitsRatios int32

const (
	GbUnitsRatios_GB_UNITS_RATIOS_UNSPECIFIED GbUnitsRatios = 0
	// Metrics Ratios
	GbUnitsRatios_GB_UNITS_RATIOS_METRICS_HIGH    GbUnitsRatios = 1
	GbUnitsRatios_GB_UNITS_RATIOS_METRICS_BLOCKED GbUnitsRatios = 2
	// Spans Ratios
	GbUnitsRatios_GB_UNITS_RATIOS_SPANS_HIGH   GbUnitsRatios = 10
	GbUnitsRatios_GB_UNITS_RATIOS_SPANS_MEDIUM GbUnitsRatios = 11
	GbUnitsRatios_GB_UNITS_RATIOS_SPANS_LOW    GbUnitsRatios = 12
	// Logs Ratios
	GbUnitsRatios_GB_UNITS_RATIOS_LOGS_HIGH    GbUnitsRatios = 20
	GbUnitsRatios_GB_UNITS_RATIOS_LOGS_MEDIUM  GbUnitsRatios = 21 // 0.4 * 0.8
	GbUnitsRatios_GB_UNITS_RATIOS_LOGS_LOW     GbUnitsRatios = 22 // 0.15 * 0.8
	GbUnitsRatios_GB_UNITS_RATIOS_LOGS_BLOCKED GbUnitsRatios = 23 // 0.1 * 0.8
	// Session Recording Ratios
	GbUnitsRatios_GB_UNITS_RATIOS_SESSION_RECORDING_MEDIUM GbUnitsRatios = 30
	GbUnitsRatios_GB_UNITS_RATIOS_SESSION_RECORDING_LOW    GbUnitsRatios = 31
)

// Enum value maps for GbUnitsRatios.
var (
	GbUnitsRatios_name = map[int32]string{
		0:  "GB_UNITS_RATIOS_UNSPECIFIED",
		1:  "GB_UNITS_RATIOS_METRICS_HIGH",
		2:  "GB_UNITS_RATIOS_METRICS_BLOCKED",
		10: "GB_UNITS_RATIOS_SPANS_HIGH",
		11: "GB_UNITS_RATIOS_SPANS_MEDIUM",
		12: "GB_UNITS_RATIOS_SPANS_LOW",
		20: "GB_UNITS_RATIOS_LOGS_HIGH",
		21: "GB_UNITS_RATIOS_LOGS_MEDIUM",
		22: "GB_UNITS_RATIOS_LOGS_LOW",
		23: "GB_UNITS_RATIOS_LOGS_BLOCKED",
		30: "GB_UNITS_RATIOS_SESSION_RECORDING_MEDIUM",
		31: "GB_UNITS_RATIOS_SESSION_RECORDING_LOW",
	}
	GbUnitsRatios_value = map[string]int32{
		"GB_UNITS_RATIOS_UNSPECIFIED":              0,
		"GB_UNITS_RATIOS_METRICS_HIGH":             1,
		"GB_UNITS_RATIOS_METRICS_BLOCKED":          2,
		"GB_UNITS_RATIOS_SPANS_HIGH":               10,
		"GB_UNITS_RATIOS_SPANS_MEDIUM":             11,
		"GB_UNITS_RATIOS_SPANS_LOW":                12,
		"GB_UNITS_RATIOS_LOGS_HIGH":                20,
		"GB_UNITS_RATIOS_LOGS_MEDIUM":              21,
		"GB_UNITS_RATIOS_LOGS_LOW":                 22,
		"GB_UNITS_RATIOS_LOGS_BLOCKED":             23,
		"GB_UNITS_RATIOS_SESSION_RECORDING_MEDIUM": 30,
		"GB_UNITS_RATIOS_SESSION_RECORDING_LOW":    31,
	}
)

func (x GbUnitsRatios) Enum() *GbUnitsRatios {
	p := new(GbUnitsRatios)
	*p = x
	return p
}

func (x GbUnitsRatios) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GbUnitsRatios) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_datausage_models_v1_units_ratios_proto_enumTypes[0].Descriptor()
}

func (GbUnitsRatios) Type() protoreflect.EnumType {
	return &file_com_coralogix_datausage_models_v1_units_ratios_proto_enumTypes[0]
}

func (x GbUnitsRatios) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GbUnitsRatios.Descriptor instead.
func (GbUnitsRatios) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_datausage_models_v1_units_ratios_proto_rawDescGZIP(), []int{0}
}

var file_com_coralogix_datausage_models_v1_units_ratios_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*float32)(nil),
		Field:         50001,
		Name:          "com.coralogix.datausage.models.v1.ratio",
		Tag:           "fixed32,50001,opt,name=ratio",
		Filename:      "com/coralogix/datausage/models/v1/units_ratios.proto",
	},
}

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional float ratio = 50001;
	E_Ratio = &file_com_coralogix_datausage_models_v1_units_ratios_proto_extTypes[0]
)

var File_com_coralogix_datausage_models_v1_units_ratios_proto protoreflect.FileDescriptor

var file_com_coralogix_datausage_models_v1_units_ratios_proto_rawDesc = []byte{
	0x0a, 0x34, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x94, 0x04, 0x0a, 0x0d,
	0x47, 0x62, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x73, 0x12, 0x1f, 0x0a,
	0x1b, 0x47, 0x42, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x53, 0x5f, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x53,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x29,
	0x0a, 0x1c, 0x47, 0x42, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x53, 0x5f, 0x52, 0x41, 0x54, 0x49, 0x4f,
	0x53, 0x5f, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x53, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x10, 0x01,
	0x1a, 0x07, 0x8d, 0xb5, 0x18, 0x2f, 0x88, 0x08, 0x3d, 0x12, 0x2c, 0x0a, 0x1f, 0x47, 0x42, 0x5f,
	0x55, 0x4e, 0x49, 0x54, 0x53, 0x5f, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x53, 0x5f, 0x4d, 0x45, 0x54,
	0x52, 0x49, 0x43, 0x53, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x02, 0x1a, 0x07,
	0x8d, 0xb5, 0x18, 0x7e, 0x73, 0x5a, 0x3b, 0x12, 0x27, 0x0a, 0x1a, 0x47, 0x42, 0x5f, 0x55, 0x4e,
	0x49, 0x54, 0x53, 0x5f, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x53, 0x5f, 0x53, 0x50, 0x41, 0x4e, 0x53,
	0x5f, 0x48, 0x49, 0x47, 0x48, 0x10, 0x0a, 0x1a, 0x07, 0x8d, 0xb5, 0x18, 0x00, 0x00, 0x00, 0x3f,
	0x12, 0x29, 0x0a, 0x1c, 0x47, 0x42, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x53, 0x5f, 0x52, 0x41, 0x54,
	0x49, 0x4f, 0x53, 0x5f, 0x53, 0x50, 0x41, 0x4e, 0x53, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x55, 0x4d,
	0x10, 0x0b, 0x1a, 0x07, 0x8d, 0xb5, 0x18, 0x00, 0x00, 0x80, 0x3e, 0x12, 0x26, 0x0a, 0x19, 0x47,
	0x42, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x53, 0x5f, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x53, 0x5f, 0x53,
	0x50, 0x41, 0x4e, 0x53, 0x5f, 0x4c, 0x4f, 0x57, 0x10, 0x0c, 0x1a, 0x07, 0x8d, 0xb5, 0x18, 0xcd,
	0xcc, 0xcc, 0x3d, 0x12, 0x26, 0x0a, 0x19, 0x47, 0x42, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x53, 0x5f,
	0x52, 0x41, 0x54, 0x49, 0x4f, 0x53, 0x5f, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x48, 0x49, 0x47, 0x48,
	0x10, 0x14, 0x1a, 0x07, 0x8d, 0xb5, 0x18, 0x00, 0x00, 0x40, 0x3f, 0x12, 0x28, 0x0a, 0x1b, 0x47,
	0x42, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x53, 0x5f, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x53, 0x5f, 0x4c,
	0x4f, 0x47, 0x53, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x55, 0x4d, 0x10, 0x15, 0x1a, 0x07, 0x8d, 0xb5,
	0x18, 0x0a, 0xd7, 0xa3, 0x3e, 0x12, 0x25, 0x0a, 0x18, 0x47, 0x42, 0x5f, 0x55, 0x4e, 0x49, 0x54,
	0x53, 0x5f, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x53, 0x5f, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x4c, 0x4f,
	0x57, 0x10, 0x16, 0x1a, 0x07, 0x8d, 0xb5, 0x18, 0x8f, 0xc2, 0xf5, 0x3d, 0x12, 0x29, 0x0a, 0x1c,
	0x47, 0x42, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x53, 0x5f, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x53, 0x5f,
	0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x17, 0x1a, 0x07,
	0x8d, 0xb5, 0x18, 0x0a, 0xd7, 0xa3, 0x3d, 0x12, 0x35, 0x0a, 0x28, 0x47, 0x42, 0x5f, 0x55, 0x4e,
	0x49, 0x54, 0x53, 0x5f, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x53, 0x5f, 0x53, 0x45, 0x53, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x4d, 0x45, 0x44,
	0x49, 0x55, 0x4d, 0x10, 0x1e, 0x1a, 0x07, 0x8d, 0xb5, 0x18, 0x00, 0x00, 0x80, 0x3e, 0x12, 0x32,
	0x0a, 0x25, 0x47, 0x42, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x53, 0x5f, 0x52, 0x41, 0x54, 0x49, 0x4f,
	0x53, 0x5f, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44,
	0x49, 0x4e, 0x47, 0x5f, 0x4c, 0x4f, 0x57, 0x10, 0x1f, 0x1a, 0x07, 0x8d, 0xb5, 0x18, 0xcd, 0xcc,
	0xcc, 0x3d, 0x3a, 0x39, 0x0a, 0x05, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x21, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e,
	0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd1,
	0x86, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_datausage_models_v1_units_ratios_proto_rawDescOnce sync.Once
	file_com_coralogix_datausage_models_v1_units_ratios_proto_rawDescData = file_com_coralogix_datausage_models_v1_units_ratios_proto_rawDesc
)

func file_com_coralogix_datausage_models_v1_units_ratios_proto_rawDescGZIP() []byte {
	file_com_coralogix_datausage_models_v1_units_ratios_proto_rawDescOnce.Do(func() {
		file_com_coralogix_datausage_models_v1_units_ratios_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_datausage_models_v1_units_ratios_proto_rawDescData)
	})
	return file_com_coralogix_datausage_models_v1_units_ratios_proto_rawDescData
}

var file_com_coralogix_datausage_models_v1_units_ratios_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogix_datausage_models_v1_units_ratios_proto_goTypes = []any{
	(GbUnitsRatios)(0),                    // 0: com.coralogix.datausage.models.v1.GbUnitsRatios
	(*descriptorpb.EnumValueOptions)(nil), // 1: google.protobuf.EnumValueOptions
}
var file_com_coralogix_datausage_models_v1_units_ratios_proto_depIdxs = []int32{
	1, // 0: com.coralogix.datausage.models.v1.ratio:extendee -> google.protobuf.EnumValueOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogix_datausage_models_v1_units_ratios_proto_init() }
func file_com_coralogix_datausage_models_v1_units_ratios_proto_init() {
	if File_com_coralogix_datausage_models_v1_units_ratios_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_datausage_models_v1_units_ratios_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_datausage_models_v1_units_ratios_proto_goTypes,
		DependencyIndexes: file_com_coralogix_datausage_models_v1_units_ratios_proto_depIdxs,
		EnumInfos:         file_com_coralogix_datausage_models_v1_units_ratios_proto_enumTypes,
		ExtensionInfos:    file_com_coralogix_datausage_models_v1_units_ratios_proto_extTypes,
	}.Build()
	File_com_coralogix_datausage_models_v1_units_ratios_proto = out.File
	file_com_coralogix_datausage_models_v1_units_ratios_proto_rawDesc = nil
	file_com_coralogix_datausage_models_v1_units_ratios_proto_goTypes = nil
	file_com_coralogix_datausage_models_v1_units_ratios_proto_depIdxs = nil
}
