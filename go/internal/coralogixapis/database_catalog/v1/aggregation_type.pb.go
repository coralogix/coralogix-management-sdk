// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.1
// source: com/coralogixapis/database_catalog/v1/aggregation_type.proto

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

type AggregationType int32

const (
	AggregationType_AGGREGATION_TYPE_UNSPECIFIED AggregationType = 0
	AggregationType_AGGREGATION_TYPE_AVG         AggregationType = 1
	AggregationType_AGGREGATION_TYPE_P95         AggregationType = 2
	AggregationType_AGGREGATION_TYPE_P99         AggregationType = 3
	AggregationType_AGGREGATION_TYPE_P50         AggregationType = 4
	AggregationType_AGGREGATION_TYPE_P75         AggregationType = 5
	AggregationType_AGGREGATION_TYPE_MAX         AggregationType = 6
)

// Enum value maps for AggregationType.
var (
	AggregationType_name = map[int32]string{
		0: "AGGREGATION_TYPE_UNSPECIFIED",
		1: "AGGREGATION_TYPE_AVG",
		2: "AGGREGATION_TYPE_P95",
		3: "AGGREGATION_TYPE_P99",
		4: "AGGREGATION_TYPE_P50",
		5: "AGGREGATION_TYPE_P75",
		6: "AGGREGATION_TYPE_MAX",
	}
	AggregationType_value = map[string]int32{
		"AGGREGATION_TYPE_UNSPECIFIED": 0,
		"AGGREGATION_TYPE_AVG":         1,
		"AGGREGATION_TYPE_P95":         2,
		"AGGREGATION_TYPE_P99":         3,
		"AGGREGATION_TYPE_P50":         4,
		"AGGREGATION_TYPE_P75":         5,
		"AGGREGATION_TYPE_MAX":         6,
	}
)

func (x AggregationType) Enum() *AggregationType {
	p := new(AggregationType)
	*p = x
	return p
}

func (x AggregationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AggregationType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_database_catalog_v1_aggregation_type_proto_enumTypes[0].Descriptor()
}

func (AggregationType) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_database_catalog_v1_aggregation_type_proto_enumTypes[0]
}

func (x AggregationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AggregationType.Descriptor instead.
func (AggregationType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_database_catalog_v1_aggregation_type_proto_rawDescGZIP(), []int{0}
}

var File_com_coralogixapis_database_catalog_v1_aggregation_type_proto protoreflect.FileDescriptor

var file_com_coralogixapis_database_catalog_v1_aggregation_type_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2a, 0xcf, 0x01, 0x0a, 0x0f, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x47, 0x47,
	0x52, 0x45, 0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x41,
	0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x41, 0x56, 0x47, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x39, 0x35, 0x10, 0x02, 0x12,
	0x18, 0x0a, 0x14, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x50, 0x39, 0x39, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x47, 0x47,
	0x52, 0x45, 0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x35,
	0x30, 0x10, 0x04, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x37, 0x35, 0x10, 0x05, 0x12, 0x18, 0x0a,
	0x14, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4d, 0x41, 0x58, 0x10, 0x06, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_database_catalog_v1_aggregation_type_proto_rawDescOnce sync.Once
	file_com_coralogixapis_database_catalog_v1_aggregation_type_proto_rawDescData = file_com_coralogixapis_database_catalog_v1_aggregation_type_proto_rawDesc
)

func file_com_coralogixapis_database_catalog_v1_aggregation_type_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_database_catalog_v1_aggregation_type_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_database_catalog_v1_aggregation_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_database_catalog_v1_aggregation_type_proto_rawDescData)
	})
	return file_com_coralogixapis_database_catalog_v1_aggregation_type_proto_rawDescData
}

var file_com_coralogixapis_database_catalog_v1_aggregation_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_database_catalog_v1_aggregation_type_proto_goTypes = []any{
	(AggregationType)(0), // 0: com.coralogixapis.database_catalog.v1.AggregationType
}
var file_com_coralogixapis_database_catalog_v1_aggregation_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_database_catalog_v1_aggregation_type_proto_init() }
func file_com_coralogixapis_database_catalog_v1_aggregation_type_proto_init() {
	if File_com_coralogixapis_database_catalog_v1_aggregation_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_database_catalog_v1_aggregation_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_database_catalog_v1_aggregation_type_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_database_catalog_v1_aggregation_type_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_database_catalog_v1_aggregation_type_proto_enumTypes,
	}.Build()
	File_com_coralogixapis_database_catalog_v1_aggregation_type_proto = out.File
	file_com_coralogixapis_database_catalog_v1_aggregation_type_proto_rawDesc = nil
	file_com_coralogixapis_database_catalog_v1_aggregation_type_proto_goTypes = nil
	file_com_coralogixapis_database_catalog_v1_aggregation_type_proto_depIdxs = nil
}
