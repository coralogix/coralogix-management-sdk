// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: com/coralogix/global_mapping/v1/data_source_type.proto

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

type DataSourceType int32

const (
	DataSourceType_DATA_SOURCE_TYPE_UNSPECIFIED DataSourceType = 0
	DataSourceType_DATA_SOURCE_TYPE_LOGS        DataSourceType = 1
	DataSourceType_DATA_SOURCE_TYPE_SPAN        DataSourceType = 2
	DataSourceType_DATA_SOURCE_TYPE_METRICS     DataSourceType = 3
	DataSourceType_DATA_SOURCE_TYPE_EVENTS      DataSourceType = 4
	DataSourceType_DATA_SOURCE_TYPE_PROFILES    DataSourceType = 5
)

// Enum value maps for DataSourceType.
var (
	DataSourceType_name = map[int32]string{
		0: "DATA_SOURCE_TYPE_UNSPECIFIED",
		1: "DATA_SOURCE_TYPE_LOGS",
		2: "DATA_SOURCE_TYPE_SPAN",
		3: "DATA_SOURCE_TYPE_METRICS",
		4: "DATA_SOURCE_TYPE_EVENTS",
		5: "DATA_SOURCE_TYPE_PROFILES",
	}
	DataSourceType_value = map[string]int32{
		"DATA_SOURCE_TYPE_UNSPECIFIED": 0,
		"DATA_SOURCE_TYPE_LOGS":        1,
		"DATA_SOURCE_TYPE_SPAN":        2,
		"DATA_SOURCE_TYPE_METRICS":     3,
		"DATA_SOURCE_TYPE_EVENTS":      4,
		"DATA_SOURCE_TYPE_PROFILES":    5,
	}
)

func (x DataSourceType) Enum() *DataSourceType {
	p := new(DataSourceType)
	*p = x
	return p
}

func (x DataSourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataSourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_global_mapping_v1_data_source_type_proto_enumTypes[0].Descriptor()
}

func (DataSourceType) Type() protoreflect.EnumType {
	return &file_com_coralogix_global_mapping_v1_data_source_type_proto_enumTypes[0]
}

func (x DataSourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataSourceType.Descriptor instead.
func (DataSourceType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_global_mapping_v1_data_source_type_proto_rawDescGZIP(), []int{0}
}

var File_com_coralogix_global_mapping_v1_data_source_type_proto protoreflect.FileDescriptor

var file_com_coralogix_global_mapping_v1_data_source_type_proto_rawDesc = []byte{
	0x0a, 0x36, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2a, 0xc2, 0x01, 0x0a, 0x0e, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x1c,
	0x44, 0x41, 0x54, 0x41, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19,
	0x0a, 0x15, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4c, 0x4f, 0x47, 0x53, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x41, 0x54,
	0x41, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x50,
	0x41, 0x4e, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x53,
	0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x53, 0x10, 0x04, 0x12,
	0x1d, 0x0a, 0x19, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x45, 0x53, 0x10, 0x05, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_global_mapping_v1_data_source_type_proto_rawDescOnce sync.Once
	file_com_coralogix_global_mapping_v1_data_source_type_proto_rawDescData = file_com_coralogix_global_mapping_v1_data_source_type_proto_rawDesc
)

func file_com_coralogix_global_mapping_v1_data_source_type_proto_rawDescGZIP() []byte {
	file_com_coralogix_global_mapping_v1_data_source_type_proto_rawDescOnce.Do(func() {
		file_com_coralogix_global_mapping_v1_data_source_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_global_mapping_v1_data_source_type_proto_rawDescData)
	})
	return file_com_coralogix_global_mapping_v1_data_source_type_proto_rawDescData
}

var file_com_coralogix_global_mapping_v1_data_source_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogix_global_mapping_v1_data_source_type_proto_goTypes = []any{
	(DataSourceType)(0), // 0: com.coralogix.global_mapping.v1.DataSourceType
}
var file_com_coralogix_global_mapping_v1_data_source_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogix_global_mapping_v1_data_source_type_proto_init() }
func file_com_coralogix_global_mapping_v1_data_source_type_proto_init() {
	if File_com_coralogix_global_mapping_v1_data_source_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_global_mapping_v1_data_source_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_global_mapping_v1_data_source_type_proto_goTypes,
		DependencyIndexes: file_com_coralogix_global_mapping_v1_data_source_type_proto_depIdxs,
		EnumInfos:         file_com_coralogix_global_mapping_v1_data_source_type_proto_enumTypes,
	}.Build()
	File_com_coralogix_global_mapping_v1_data_source_type_proto = out.File
	file_com_coralogix_global_mapping_v1_data_source_type_proto_rawDesc = nil
	file_com_coralogix_global_mapping_v1_data_source_type_proto_goTypes = nil
	file_com_coralogix_global_mapping_v1_data_source_type_proto_depIdxs = nil
}
