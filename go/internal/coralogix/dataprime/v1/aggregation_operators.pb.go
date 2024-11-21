// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: com/coralogix/dataprime/v1/aggregation_operators.proto

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

type Operator int32

const (
	Operator_OPERATOR_COUNT_UNSPECIFIED Operator = 0
	Operator_OPERATOR_AVG               Operator = 1
	Operator_OPERATOR_SUM               Operator = 2
	Operator_OPERATOR_MIN               Operator = 3
	Operator_OPERATOR_MAX               Operator = 4
)

// Enum value maps for Operator.
var (
	Operator_name = map[int32]string{
		0: "OPERATOR_COUNT_UNSPECIFIED",
		1: "OPERATOR_AVG",
		2: "OPERATOR_SUM",
		3: "OPERATOR_MIN",
		4: "OPERATOR_MAX",
	}
	Operator_value = map[string]int32{
		"OPERATOR_COUNT_UNSPECIFIED": 0,
		"OPERATOR_AVG":               1,
		"OPERATOR_SUM":               2,
		"OPERATOR_MIN":               3,
		"OPERATOR_MAX":               4,
	}
)

func (x Operator) Enum() *Operator {
	p := new(Operator)
	*p = x
	return p
}

func (x Operator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Operator) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_dataprime_v1_aggregation_operators_proto_enumTypes[0].Descriptor()
}

func (Operator) Type() protoreflect.EnumType {
	return &file_com_coralogix_dataprime_v1_aggregation_operators_proto_enumTypes[0]
}

func (x Operator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operator.Descriptor instead.
func (Operator) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_aggregation_operators_proto_rawDescGZIP(), []int{0}
}

var File_com_coralogix_dataprime_v1_aggregation_operators_proto protoreflect.FileDescriptor

var file_com_coralogix_dataprime_v1_aggregation_operators_proto_rawDesc = []byte{
	0x0a, 0x36, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d,
	0x65, 0x2e, 0x76, 0x31, 0x2a, 0x72, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x12, 0x1e, 0x0a, 0x1a, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x55,
	0x4e, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x10, 0x0a, 0x0c, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x41, 0x56, 0x47,
	0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x53,
	0x55, 0x4d, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52,
	0x5f, 0x4d, 0x49, 0x4e, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54,
	0x4f, 0x52, 0x5f, 0x4d, 0x41, 0x58, 0x10, 0x04, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_dataprime_v1_aggregation_operators_proto_rawDescOnce sync.Once
	file_com_coralogix_dataprime_v1_aggregation_operators_proto_rawDescData = file_com_coralogix_dataprime_v1_aggregation_operators_proto_rawDesc
)

func file_com_coralogix_dataprime_v1_aggregation_operators_proto_rawDescGZIP() []byte {
	file_com_coralogix_dataprime_v1_aggregation_operators_proto_rawDescOnce.Do(func() {
		file_com_coralogix_dataprime_v1_aggregation_operators_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_dataprime_v1_aggregation_operators_proto_rawDescData)
	})
	return file_com_coralogix_dataprime_v1_aggregation_operators_proto_rawDescData
}

var file_com_coralogix_dataprime_v1_aggregation_operators_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogix_dataprime_v1_aggregation_operators_proto_goTypes = []any{
	(Operator)(0), // 0: com.coralogix.dataprime.v1.Operator
}
var file_com_coralogix_dataprime_v1_aggregation_operators_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogix_dataprime_v1_aggregation_operators_proto_init() }
func file_com_coralogix_dataprime_v1_aggregation_operators_proto_init() {
	if File_com_coralogix_dataprime_v1_aggregation_operators_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_dataprime_v1_aggregation_operators_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_dataprime_v1_aggregation_operators_proto_goTypes,
		DependencyIndexes: file_com_coralogix_dataprime_v1_aggregation_operators_proto_depIdxs,
		EnumInfos:         file_com_coralogix_dataprime_v1_aggregation_operators_proto_enumTypes,
	}.Build()
	File_com_coralogix_dataprime_v1_aggregation_operators_proto = out.File
	file_com_coralogix_dataprime_v1_aggregation_operators_proto_rawDesc = nil
	file_com_coralogix_dataprime_v1_aggregation_operators_proto_goTypes = nil
	file_com_coralogix_dataprime_v1_aggregation_operators_proto_depIdxs = nil
}
