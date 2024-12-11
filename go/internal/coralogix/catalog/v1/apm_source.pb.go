// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.1
// source: com/coralogixapis/service_catalog/v1/apm_source.proto

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

type ApmSource int32

const (
	ApmSource_APM_SOURCE_UNSPECIFIED  ApmSource = 0
	ApmSource_APM_SOURCE_E2M          ApmSource = 1
	ApmSource_APM_SOURCE_SPAN_METRICS ApmSource = 2
)

// Enum value maps for ApmSource.
var (
	ApmSource_name = map[int32]string{
		0: "APM_SOURCE_UNSPECIFIED",
		1: "APM_SOURCE_E2M",
		2: "APM_SOURCE_SPAN_METRICS",
	}
	ApmSource_value = map[string]int32{
		"APM_SOURCE_UNSPECIFIED":  0,
		"APM_SOURCE_E2M":          1,
		"APM_SOURCE_SPAN_METRICS": 2,
	}
)

func (x ApmSource) Enum() *ApmSource {
	p := new(ApmSource)
	*p = x
	return p
}

func (x ApmSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApmSource) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_service_catalog_v1_apm_source_proto_enumTypes[0].Descriptor()
}

func (ApmSource) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_service_catalog_v1_apm_source_proto_enumTypes[0]
}

func (x ApmSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApmSource.Descriptor instead.
func (ApmSource) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_service_catalog_v1_apm_source_proto_rawDescGZIP(), []int{0}
}

var File_com_coralogixapis_service_catalog_v1_apm_source_proto protoreflect.FileDescriptor

var file_com_coralogixapis_service_catalog_v1_apm_source_proto_rawDesc = []byte{
	0x0a, 0x35, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x6d, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2a, 0x58, 0x0a,
	0x09, 0x41, 0x70, 0x6d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x41, 0x50,
	0x4d, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x50, 0x4d, 0x5f, 0x53, 0x4f,
	0x55, 0x52, 0x43, 0x45, 0x5f, 0x45, 0x32, 0x4d, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x50,
	0x4d, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x53, 0x50, 0x41, 0x4e, 0x5f, 0x4d, 0x45,
	0x54, 0x52, 0x49, 0x43, 0x53, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_service_catalog_v1_apm_source_proto_rawDescOnce sync.Once
	file_com_coralogixapis_service_catalog_v1_apm_source_proto_rawDescData = file_com_coralogixapis_service_catalog_v1_apm_source_proto_rawDesc
)

func file_com_coralogixapis_service_catalog_v1_apm_source_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_service_catalog_v1_apm_source_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_service_catalog_v1_apm_source_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_service_catalog_v1_apm_source_proto_rawDescData)
	})
	return file_com_coralogixapis_service_catalog_v1_apm_source_proto_rawDescData
}

var file_com_coralogixapis_service_catalog_v1_apm_source_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_service_catalog_v1_apm_source_proto_goTypes = []any{
	(ApmSource)(0), // 0: com.coralogixapis.service_catalog.v1.ApmSource
}
var file_com_coralogixapis_service_catalog_v1_apm_source_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_service_catalog_v1_apm_source_proto_init() }
func file_com_coralogixapis_service_catalog_v1_apm_source_proto_init() {
	if File_com_coralogixapis_service_catalog_v1_apm_source_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_service_catalog_v1_apm_source_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_service_catalog_v1_apm_source_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_service_catalog_v1_apm_source_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_service_catalog_v1_apm_source_proto_enumTypes,
	}.Build()
	File_com_coralogixapis_service_catalog_v1_apm_source_proto = out.File
	file_com_coralogixapis_service_catalog_v1_apm_source_proto_rawDesc = nil
	file_com_coralogixapis_service_catalog_v1_apm_source_proto_goTypes = nil
	file_com_coralogixapis_service_catalog_v1_apm_source_proto_depIdxs = nil
}
