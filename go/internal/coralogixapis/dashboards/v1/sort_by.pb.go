// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: src/com/coralogixapis/dashboards/v1/ast/widgets/common/sort_by.proto

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

type SortByType int32

const (
	SortByType_SORT_BY_TYPE_UNSPECIFIED SortByType = 0
	SortByType_SORT_BY_TYPE_VALUE       SortByType = 1
	SortByType_SORT_BY_TYPE_NAME        SortByType = 2
)

// Enum value maps for SortByType.
var (
	SortByType_name = map[int32]string{
		0: "SORT_BY_TYPE_UNSPECIFIED",
		1: "SORT_BY_TYPE_VALUE",
		2: "SORT_BY_TYPE_NAME",
	}
	SortByType_value = map[string]int32{
		"SORT_BY_TYPE_UNSPECIFIED": 0,
		"SORT_BY_TYPE_VALUE":       1,
		"SORT_BY_TYPE_NAME":        2,
	}
)

func (x SortByType) Enum() *SortByType {
	p := new(SortByType)
	*p = x
	return p
}

func (x SortByType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortByType) Descriptor() protoreflect.EnumDescriptor {
	return file_src_com_coralogixapis_dashboards_v1_ast_widgets_common_sort_by_proto_enumTypes[0].Descriptor()
}

func (SortByType) Type() protoreflect.EnumType {
	return &file_src_com_coralogixapis_dashboards_v1_ast_widgets_common_sort_by_proto_enumTypes[0]
}

func (x SortByType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortByType.Descriptor instead.
func (SortByType) EnumDescriptor() ([]byte, []int) {
	return file_src_com_coralogixapis_dashboards_v1_ast_widgets_common_sort_by_proto_rawDescGZIP(), []int{0}
}

var File_src_com_coralogixapis_dashboards_v1_ast_widgets_common_sort_by_proto protoreflect.FileDescriptor

var file_src_com_coralogixapis_dashboards_v1_ast_widgets_common_sort_by_proto_rawDesc = []byte{
	0x0a, 0x44, 0x73, 0x72, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x74, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74,
	0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x32, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x77, 0x69, 0x64, 0x67,
	0x65, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2a, 0x59, 0x0a, 0x0a, 0x53, 0x6f,
	0x72, 0x74, 0x42, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x4f, 0x52, 0x54,
	0x5f, 0x42, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x4f, 0x52, 0x54, 0x5f, 0x42,
	0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x01, 0x12, 0x15,
	0x0a, 0x11, 0x53, 0x4f, 0x52, 0x54, 0x5f, 0x42, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e,
	0x41, 0x4d, 0x45, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_com_coralogixapis_dashboards_v1_ast_widgets_common_sort_by_proto_rawDescOnce sync.Once
	file_src_com_coralogixapis_dashboards_v1_ast_widgets_common_sort_by_proto_rawDescData = file_src_com_coralogixapis_dashboards_v1_ast_widgets_common_sort_by_proto_rawDesc
)

func file_src_com_coralogixapis_dashboards_v1_ast_widgets_common_sort_by_proto_rawDescGZIP() []byte {
	file_src_com_coralogixapis_dashboards_v1_ast_widgets_common_sort_by_proto_rawDescOnce.Do(func() {
		file_src_com_coralogixapis_dashboards_v1_ast_widgets_common_sort_by_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_com_coralogixapis_dashboards_v1_ast_widgets_common_sort_by_proto_rawDescData)
	})
	return file_src_com_coralogixapis_dashboards_v1_ast_widgets_common_sort_by_proto_rawDescData
}

var file_src_com_coralogixapis_dashboards_v1_ast_widgets_common_sort_by_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_src_com_coralogixapis_dashboards_v1_ast_widgets_common_sort_by_proto_goTypes = []any{
	(SortByType)(0), // 0: com.coralogixapis.dashboards.v1.ast.widgets.common.SortByType
}
var file_src_com_coralogixapis_dashboards_v1_ast_widgets_common_sort_by_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_src_com_coralogixapis_dashboards_v1_ast_widgets_common_sort_by_proto_init() }
func file_src_com_coralogixapis_dashboards_v1_ast_widgets_common_sort_by_proto_init() {
	if File_src_com_coralogixapis_dashboards_v1_ast_widgets_common_sort_by_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_src_com_coralogixapis_dashboards_v1_ast_widgets_common_sort_by_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_com_coralogixapis_dashboards_v1_ast_widgets_common_sort_by_proto_goTypes,
		DependencyIndexes: file_src_com_coralogixapis_dashboards_v1_ast_widgets_common_sort_by_proto_depIdxs,
		EnumInfos:         file_src_com_coralogixapis_dashboards_v1_ast_widgets_common_sort_by_proto_enumTypes,
	}.Build()
	File_src_com_coralogixapis_dashboards_v1_ast_widgets_common_sort_by_proto = out.File
	file_src_com_coralogixapis_dashboards_v1_ast_widgets_common_sort_by_proto_rawDesc = nil
	file_src_com_coralogixapis_dashboards_v1_ast_widgets_common_sort_by_proto_goTypes = nil
	file_src_com_coralogixapis_dashboards_v1_ast_widgets_common_sort_by_proto_depIdxs = nil
}
