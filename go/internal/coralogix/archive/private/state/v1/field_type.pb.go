// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogix/archive/private/state/v1/field_type.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FieldType int32

const (
	FieldType_FIELD_TYPE_UNSPECIFIED      FieldType = 0
	FieldType_FIELD_TYPE_STRING           FieldType = 1
	FieldType_FIELD_TYPE_STRING_ARRAY     FieldType = 2
	FieldType_FIELD_TYPE_TIMESTAMP_NANOS  FieldType = 3
	FieldType_FIELD_TYPE_TIMESTAMP_MICROS FieldType = 4
)

// Enum value maps for FieldType.
var (
	FieldType_name = map[int32]string{
		0: "FIELD_TYPE_UNSPECIFIED",
		1: "FIELD_TYPE_STRING",
		2: "FIELD_TYPE_STRING_ARRAY",
		3: "FIELD_TYPE_TIMESTAMP_NANOS",
		4: "FIELD_TYPE_TIMESTAMP_MICROS",
	}
	FieldType_value = map[string]int32{
		"FIELD_TYPE_UNSPECIFIED":      0,
		"FIELD_TYPE_STRING":           1,
		"FIELD_TYPE_STRING_ARRAY":     2,
		"FIELD_TYPE_TIMESTAMP_NANOS":  3,
		"FIELD_TYPE_TIMESTAMP_MICROS": 4,
	}
)

func (x FieldType) Enum() *FieldType {
	p := new(FieldType)
	*p = x
	return p
}

func (x FieldType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FieldType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_archive_private_state_v1_field_type_proto_enumTypes[0].Descriptor()
}

func (FieldType) Type() protoreflect.EnumType {
	return &file_com_coralogix_archive_private_state_v1_field_type_proto_enumTypes[0]
}

func (x FieldType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FieldType.Descriptor instead.
func (FieldType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_archive_private_state_v1_field_type_proto_rawDescGZIP(), []int{0}
}

var File_com_coralogix_archive_private_state_v1_field_type_proto protoreflect.FileDescriptor

const file_com_coralogix_archive_private_state_v1_field_type_proto_rawDesc = "" +
	"\n" +
	"7com/coralogix/archive/private/state/v1/field_type.proto\x12&com.coralogix.archive.private.state.v1*\x9c\x01\n" +
	"\tFieldType\x12\x1a\n" +
	"\x16FIELD_TYPE_UNSPECIFIED\x10\x00\x12\x15\n" +
	"\x11FIELD_TYPE_STRING\x10\x01\x12\x1b\n" +
	"\x17FIELD_TYPE_STRING_ARRAY\x10\x02\x12\x1e\n" +
	"\x1aFIELD_TYPE_TIMESTAMP_NANOS\x10\x03\x12\x1f\n" +
	"\x1bFIELD_TYPE_TIMESTAMP_MICROS\x10\x04b\x06proto3"

var (
	file_com_coralogix_archive_private_state_v1_field_type_proto_rawDescOnce sync.Once
	file_com_coralogix_archive_private_state_v1_field_type_proto_rawDescData []byte
)

func file_com_coralogix_archive_private_state_v1_field_type_proto_rawDescGZIP() []byte {
	file_com_coralogix_archive_private_state_v1_field_type_proto_rawDescOnce.Do(func() {
		file_com_coralogix_archive_private_state_v1_field_type_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogix_archive_private_state_v1_field_type_proto_rawDesc), len(file_com_coralogix_archive_private_state_v1_field_type_proto_rawDesc)))
	})
	return file_com_coralogix_archive_private_state_v1_field_type_proto_rawDescData
}

var file_com_coralogix_archive_private_state_v1_field_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogix_archive_private_state_v1_field_type_proto_goTypes = []any{
	(FieldType)(0), // 0: com.coralogix.archive.private.state.v1.FieldType
}
var file_com_coralogix_archive_private_state_v1_field_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogix_archive_private_state_v1_field_type_proto_init() }
func file_com_coralogix_archive_private_state_v1_field_type_proto_init() {
	if File_com_coralogix_archive_private_state_v1_field_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogix_archive_private_state_v1_field_type_proto_rawDesc), len(file_com_coralogix_archive_private_state_v1_field_type_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_archive_private_state_v1_field_type_proto_goTypes,
		DependencyIndexes: file_com_coralogix_archive_private_state_v1_field_type_proto_depIdxs,
		EnumInfos:         file_com_coralogix_archive_private_state_v1_field_type_proto_enumTypes,
	}.Build()
	File_com_coralogix_archive_private_state_v1_field_type_proto = out.File
	file_com_coralogix_archive_private_state_v1_field_type_proto_goTypes = nil
	file_com_coralogix_archive_private_state_v1_field_type_proto_depIdxs = nil
}
