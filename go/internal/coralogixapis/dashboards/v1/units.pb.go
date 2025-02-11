// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: com/coralogixapis/dashboards/v1/ast/widgets/common/units.proto

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

type Unit int32

const (
	Unit_UNIT_UNSPECIFIED          Unit = 0
	Unit_UNIT_MICROSECONDS         Unit = 1
	Unit_UNIT_MILLISECONDS         Unit = 2
	Unit_UNIT_SECONDS              Unit = 3
	Unit_UNIT_BYTES                Unit = 4
	Unit_UNIT_KBYTES               Unit = 5
	Unit_UNIT_MBYTES               Unit = 6
	Unit_UNIT_GBYTES               Unit = 7
	Unit_UNIT_BYTES_IEC            Unit = 8
	Unit_UNIT_KIBYTES              Unit = 9
	Unit_UNIT_MIBYTES              Unit = 10
	Unit_UNIT_GIBYTES              Unit = 11
	Unit_UNIT_EUR_CENTS            Unit = 12
	Unit_UNIT_EUR                  Unit = 13
	Unit_UNIT_USD_CENTS            Unit = 14
	Unit_UNIT_USD                  Unit = 15
	Unit_UNIT_NANOSECONDS          Unit = 16
	Unit_UNIT_CUSTOM               Unit = 17
	Unit_UNIT_PERCENT_ZERO_ONE     Unit = 18
	Unit_UNIT_PERCENT_ZERO_HUNDRED Unit = 19
	Unit_UNIT_PERCENT              Unit = 20
)

// Enum value maps for Unit.
var (
	Unit_name = map[int32]string{
		0:  "UNIT_UNSPECIFIED",
		1:  "UNIT_MICROSECONDS",
		2:  "UNIT_MILLISECONDS",
		3:  "UNIT_SECONDS",
		4:  "UNIT_BYTES",
		5:  "UNIT_KBYTES",
		6:  "UNIT_MBYTES",
		7:  "UNIT_GBYTES",
		8:  "UNIT_BYTES_IEC",
		9:  "UNIT_KIBYTES",
		10: "UNIT_MIBYTES",
		11: "UNIT_GIBYTES",
		12: "UNIT_EUR_CENTS",
		13: "UNIT_EUR",
		14: "UNIT_USD_CENTS",
		15: "UNIT_USD",
		16: "UNIT_NANOSECONDS",
		17: "UNIT_CUSTOM",
		18: "UNIT_PERCENT_ZERO_ONE",
		19: "UNIT_PERCENT_ZERO_HUNDRED",
		20: "UNIT_PERCENT",
	}
	Unit_value = map[string]int32{
		"UNIT_UNSPECIFIED":          0,
		"UNIT_MICROSECONDS":         1,
		"UNIT_MILLISECONDS":         2,
		"UNIT_SECONDS":              3,
		"UNIT_BYTES":                4,
		"UNIT_KBYTES":               5,
		"UNIT_MBYTES":               6,
		"UNIT_GBYTES":               7,
		"UNIT_BYTES_IEC":            8,
		"UNIT_KIBYTES":              9,
		"UNIT_MIBYTES":              10,
		"UNIT_GIBYTES":              11,
		"UNIT_EUR_CENTS":            12,
		"UNIT_EUR":                  13,
		"UNIT_USD_CENTS":            14,
		"UNIT_USD":                  15,
		"UNIT_NANOSECONDS":          16,
		"UNIT_CUSTOM":               17,
		"UNIT_PERCENT_ZERO_ONE":     18,
		"UNIT_PERCENT_ZERO_HUNDRED": 19,
		"UNIT_PERCENT":              20,
	}
)

func (x Unit) Enum() *Unit {
	p := new(Unit)
	*p = x
	return p
}

func (x Unit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Unit) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_dashboards_v1_ast_widgets_common_units_proto_enumTypes[0].Descriptor()
}

func (Unit) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_dashboards_v1_ast_widgets_common_units_proto_enumTypes[0]
}

func (x Unit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Unit.Descriptor instead.
func (Unit) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_widgets_common_units_proto_rawDescGZIP(), []int{0}
}

var File_com_coralogixapis_dashboards_v1_ast_widgets_common_units_proto protoreflect.FileDescriptor

var file_com_coralogixapis_dashboards_v1_ast_widgets_common_units_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x73, 0x74, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x32, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2a, 0xa0, 0x03, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x14, 0x0a,
	0x10, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x4d, 0x49, 0x43, 0x52,
	0x4f, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x53, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x4e,
	0x49, 0x54, 0x5f, 0x4d, 0x49, 0x4c, 0x4c, 0x49, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x53, 0x10,
	0x02, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44,
	0x53, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x42, 0x59, 0x54, 0x45,
	0x53, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x4b, 0x42, 0x59, 0x54,
	0x45, 0x53, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x4d, 0x42, 0x59,
	0x54, 0x45, 0x53, 0x10, 0x06, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x47, 0x42,
	0x59, 0x54, 0x45, 0x53, 0x10, 0x07, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x42,
	0x59, 0x54, 0x45, 0x53, 0x5f, 0x49, 0x45, 0x43, 0x10, 0x08, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x4e,
	0x49, 0x54, 0x5f, 0x4b, 0x49, 0x42, 0x59, 0x54, 0x45, 0x53, 0x10, 0x09, 0x12, 0x10, 0x0a, 0x0c,
	0x55, 0x4e, 0x49, 0x54, 0x5f, 0x4d, 0x49, 0x42, 0x59, 0x54, 0x45, 0x53, 0x10, 0x0a, 0x12, 0x10,
	0x0a, 0x0c, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x47, 0x49, 0x42, 0x59, 0x54, 0x45, 0x53, 0x10, 0x0b,
	0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x45, 0x55, 0x52, 0x5f, 0x43, 0x45, 0x4e,
	0x54, 0x53, 0x10, 0x0c, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x45, 0x55, 0x52,
	0x10, 0x0d, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x55, 0x53, 0x44, 0x5f, 0x43,
	0x45, 0x4e, 0x54, 0x53, 0x10, 0x0e, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x55,
	0x53, 0x44, 0x10, 0x0f, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x4e, 0x41, 0x4e,
	0x4f, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x53, 0x10, 0x10, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e,
	0x49, 0x54, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x10, 0x11, 0x12, 0x19, 0x0a, 0x15, 0x55,
	0x4e, 0x49, 0x54, 0x5f, 0x50, 0x45, 0x52, 0x43, 0x45, 0x4e, 0x54, 0x5f, 0x5a, 0x45, 0x52, 0x4f,
	0x5f, 0x4f, 0x4e, 0x45, 0x10, 0x12, 0x12, 0x1d, 0x0a, 0x19, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x50,
	0x45, 0x52, 0x43, 0x45, 0x4e, 0x54, 0x5f, 0x5a, 0x45, 0x52, 0x4f, 0x5f, 0x48, 0x55, 0x4e, 0x44,
	0x52, 0x45, 0x44, 0x10, 0x13, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x50, 0x45,
	0x52, 0x43, 0x45, 0x4e, 0x54, 0x10, 0x14, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_units_proto_rawDescOnce sync.Once
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_units_proto_rawDescData = file_com_coralogixapis_dashboards_v1_ast_widgets_common_units_proto_rawDesc
)

func file_com_coralogixapis_dashboards_v1_ast_widgets_common_units_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_units_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_dashboards_v1_ast_widgets_common_units_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_dashboards_v1_ast_widgets_common_units_proto_rawDescData)
	})
	return file_com_coralogixapis_dashboards_v1_ast_widgets_common_units_proto_rawDescData
}

var file_com_coralogixapis_dashboards_v1_ast_widgets_common_units_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_dashboards_v1_ast_widgets_common_units_proto_goTypes = []any{
	(Unit)(0), // 0: com.coralogixapis.dashboards.v1.ast.widgets.common.Unit
}
var file_com_coralogixapis_dashboards_v1_ast_widgets_common_units_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_dashboards_v1_ast_widgets_common_units_proto_init() }
func file_com_coralogixapis_dashboards_v1_ast_widgets_common_units_proto_init() {
	if File_com_coralogixapis_dashboards_v1_ast_widgets_common_units_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_dashboards_v1_ast_widgets_common_units_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_dashboards_v1_ast_widgets_common_units_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_dashboards_v1_ast_widgets_common_units_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_dashboards_v1_ast_widgets_common_units_proto_enumTypes,
	}.Build()
	File_com_coralogixapis_dashboards_v1_ast_widgets_common_units_proto = out.File
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_units_proto_rawDesc = nil
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_units_proto_goTypes = nil
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_units_proto_depIdxs = nil
}
