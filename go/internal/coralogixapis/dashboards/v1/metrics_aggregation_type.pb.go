// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: com/coralogixapis/dashboards/v1/ast/widgets/common/metrics_aggregation_type.proto

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

type Aggregation int32

const (
	Aggregation_AGGREGATION_UNSPECIFIED Aggregation = 0
	Aggregation_AGGREGATION_LAST        Aggregation = 1
	Aggregation_AGGREGATION_MIN         Aggregation = 2
	Aggregation_AGGREGATION_MAX         Aggregation = 3
	Aggregation_AGGREGATION_AVG         Aggregation = 4
	Aggregation_AGGREGATION_SUM         Aggregation = 5
)

// Enum value maps for Aggregation.
var (
	Aggregation_name = map[int32]string{
		0: "AGGREGATION_UNSPECIFIED",
		1: "AGGREGATION_LAST",
		2: "AGGREGATION_MIN",
		3: "AGGREGATION_MAX",
		4: "AGGREGATION_AVG",
		5: "AGGREGATION_SUM",
	}
	Aggregation_value = map[string]int32{
		"AGGREGATION_UNSPECIFIED": 0,
		"AGGREGATION_LAST":        1,
		"AGGREGATION_MIN":         2,
		"AGGREGATION_MAX":         3,
		"AGGREGATION_AVG":         4,
		"AGGREGATION_SUM":         5,
	}
)

func (x Aggregation) Enum() *Aggregation {
	p := new(Aggregation)
	*p = x
	return p
}

func (x Aggregation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Aggregation) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_aggregation_type_proto_enumTypes[0].Descriptor()
}

func (Aggregation) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_aggregation_type_proto_enumTypes[0]
}

func (x Aggregation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Aggregation.Descriptor instead.
func (Aggregation) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_aggregation_type_proto_rawDescGZIP(), []int{0}
}

var File_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_aggregation_type_proto protoreflect.FileDescriptor

var file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_aggregation_type_proto_rawDesc = []byte{
	0x0a, 0x51, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x73, 0x74, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x32, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2a, 0x94, 0x01, 0x0a, 0x0b, 0x41, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x47, 0x47, 0x52, 0x45,
	0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x41, 0x53, 0x54, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x47,
	0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x49, 0x4e, 0x10, 0x02, 0x12,
	0x13, 0x0a, 0x0f, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d,
	0x41, 0x58, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x56, 0x47, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x47, 0x47,
	0x52, 0x45, 0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x55, 0x4d, 0x10, 0x05, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_aggregation_type_proto_rawDescOnce sync.Once
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_aggregation_type_proto_rawDescData = file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_aggregation_type_proto_rawDesc
)

func file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_aggregation_type_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_aggregation_type_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_aggregation_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_aggregation_type_proto_rawDescData)
	})
	return file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_aggregation_type_proto_rawDescData
}

var file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_aggregation_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_aggregation_type_proto_goTypes = []any{
	(Aggregation)(0), // 0: com.coralogixapis.dashboards.v1.ast.widgets.common.Aggregation
}
var file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_aggregation_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_aggregation_type_proto_init()
}
func file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_aggregation_type_proto_init() {
	if File_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_aggregation_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_aggregation_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_aggregation_type_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_aggregation_type_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_aggregation_type_proto_enumTypes,
	}.Build()
	File_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_aggregation_type_proto = out.File
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_aggregation_type_proto_rawDesc = nil
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_aggregation_type_proto_goTypes = nil
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_aggregation_type_proto_depIdxs = nil
}
