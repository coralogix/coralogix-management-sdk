// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: com/coralogixapis/apm/widgets/v1/units.proto

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
	Unit_UNIT_UNSPECIFIED         Unit = 0
	Unit_UNIT_BYTES               Unit = 1
	Unit_UNIT_MESSAGES            Unit = 2
	Unit_UNIT_MESSAGES_PER_SECOND Unit = 3
	Unit_UNIT_BYTES_PER_SECOND    Unit = 4
	Unit_UNIT_SECONDS             Unit = 5
	Unit_UNIT_MILLISECONDS        Unit = 6
	Unit_UNIT_X_PER_SECOND        Unit = 7
	Unit_UNIT_PERCENT             Unit = 8
	Unit_UNIT_MICROSECONDS        Unit = 9
	Unit_UNIT_MESSAGES_PER_MINUTE Unit = 10
	Unit_UNIT_X_PER_MINUTE        Unit = 11
	Unit_UNIT_MEGABYTES           Unit = 12
)

// Enum value maps for Unit.
var (
	Unit_name = map[int32]string{
		0:  "UNIT_UNSPECIFIED",
		1:  "UNIT_BYTES",
		2:  "UNIT_MESSAGES",
		3:  "UNIT_MESSAGES_PER_SECOND",
		4:  "UNIT_BYTES_PER_SECOND",
		5:  "UNIT_SECONDS",
		6:  "UNIT_MILLISECONDS",
		7:  "UNIT_X_PER_SECOND",
		8:  "UNIT_PERCENT",
		9:  "UNIT_MICROSECONDS",
		10: "UNIT_MESSAGES_PER_MINUTE",
		11: "UNIT_X_PER_MINUTE",
		12: "UNIT_MEGABYTES",
	}
	Unit_value = map[string]int32{
		"UNIT_UNSPECIFIED":         0,
		"UNIT_BYTES":               1,
		"UNIT_MESSAGES":            2,
		"UNIT_MESSAGES_PER_SECOND": 3,
		"UNIT_BYTES_PER_SECOND":    4,
		"UNIT_SECONDS":             5,
		"UNIT_MILLISECONDS":        6,
		"UNIT_X_PER_SECOND":        7,
		"UNIT_PERCENT":             8,
		"UNIT_MICROSECONDS":        9,
		"UNIT_MESSAGES_PER_MINUTE": 10,
		"UNIT_X_PER_MINUTE":        11,
		"UNIT_MEGABYTES":           12,
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
	return file_com_coralogixapis_apm_widgets_v1_units_proto_enumTypes[0].Descriptor()
}

func (Unit) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_apm_widgets_v1_units_proto_enumTypes[0]
}

func (x Unit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Unit.Descriptor instead.
func (Unit) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_widgets_v1_units_proto_rawDescGZIP(), []int{0}
}

var File_com_coralogixapis_apm_widgets_v1_units_proto protoreflect.FileDescriptor

var file_com_coralogixapis_apm_widgets_v1_units_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2a, 0xaa, 0x02, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x4e, 0x49,
	0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0e, 0x0a, 0x0a, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x42, 0x59, 0x54, 0x45, 0x53, 0x10, 0x01, 0x12,
	0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x53,
	0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41,
	0x47, 0x45, 0x53, 0x5f, 0x50, 0x45, 0x52, 0x5f, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x10, 0x03,
	0x12, 0x19, 0x0a, 0x15, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x42, 0x59, 0x54, 0x45, 0x53, 0x5f, 0x50,
	0x45, 0x52, 0x5f, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x55,
	0x4e, 0x49, 0x54, 0x5f, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x53, 0x10, 0x05, 0x12, 0x15, 0x0a,
	0x11, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x4d, 0x49, 0x4c, 0x4c, 0x49, 0x53, 0x45, 0x43, 0x4f, 0x4e,
	0x44, 0x53, 0x10, 0x06, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x58, 0x5f, 0x50,
	0x45, 0x52, 0x5f, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x10, 0x07, 0x12, 0x10, 0x0a, 0x0c, 0x55,
	0x4e, 0x49, 0x54, 0x5f, 0x50, 0x45, 0x52, 0x43, 0x45, 0x4e, 0x54, 0x10, 0x08, 0x12, 0x15, 0x0a,
	0x11, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x4d, 0x49, 0x43, 0x52, 0x4f, 0x53, 0x45, 0x43, 0x4f, 0x4e,
	0x44, 0x53, 0x10, 0x09, 0x12, 0x1c, 0x0a, 0x18, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x4d, 0x45, 0x53,
	0x53, 0x41, 0x47, 0x45, 0x53, 0x5f, 0x50, 0x45, 0x52, 0x5f, 0x4d, 0x49, 0x4e, 0x55, 0x54, 0x45,
	0x10, 0x0a, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x58, 0x5f, 0x50, 0x45, 0x52,
	0x5f, 0x4d, 0x49, 0x4e, 0x55, 0x54, 0x45, 0x10, 0x0b, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x49,
	0x54, 0x5f, 0x4d, 0x45, 0x47, 0x41, 0x42, 0x59, 0x54, 0x45, 0x53, 0x10, 0x0c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_apm_widgets_v1_units_proto_rawDescOnce sync.Once
	file_com_coralogixapis_apm_widgets_v1_units_proto_rawDescData = file_com_coralogixapis_apm_widgets_v1_units_proto_rawDesc
)

func file_com_coralogixapis_apm_widgets_v1_units_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_apm_widgets_v1_units_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_apm_widgets_v1_units_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_apm_widgets_v1_units_proto_rawDescData)
	})
	return file_com_coralogixapis_apm_widgets_v1_units_proto_rawDescData
}

var file_com_coralogixapis_apm_widgets_v1_units_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_apm_widgets_v1_units_proto_goTypes = []any{
	(Unit)(0), // 0: com.coralogixapis.apm.widgets.v1.Unit
}
var file_com_coralogixapis_apm_widgets_v1_units_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_apm_widgets_v1_units_proto_init() }
func file_com_coralogixapis_apm_widgets_v1_units_proto_init() {
	if File_com_coralogixapis_apm_widgets_v1_units_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_apm_widgets_v1_units_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_apm_widgets_v1_units_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_apm_widgets_v1_units_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_apm_widgets_v1_units_proto_enumTypes,
	}.Build()
	File_com_coralogixapis_apm_widgets_v1_units_proto = out.File
	file_com_coralogixapis_apm_widgets_v1_units_proto_rawDesc = nil
	file_com_coralogixapis_apm_widgets_v1_units_proto_goTypes = nil
	file_com_coralogixapis_apm_widgets_v1_units_proto_depIdxs = nil
}
