// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.1
// source: com/coralogix/catalog/v1/span_kind.proto

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

type SpanKind int32

const (
	SpanKind_SPAN_KIND_UNSPECIFIED SpanKind = 0
	SpanKind_SPAN_KIND_SERVER      SpanKind = 1
	SpanKind_SPAN_KIND_CLIENT      SpanKind = 2
	SpanKind_SPAN_KIND_CONSUMER    SpanKind = 3
	SpanKind_SPAN_KIND_PRODUCER    SpanKind = 4
	SpanKind_SPAN_KIND_INTERNAL    SpanKind = 5
)

// Enum value maps for SpanKind.
var (
	SpanKind_name = map[int32]string{
		0: "SPAN_KIND_UNSPECIFIED",
		1: "SPAN_KIND_SERVER",
		2: "SPAN_KIND_CLIENT",
		3: "SPAN_KIND_CONSUMER",
		4: "SPAN_KIND_PRODUCER",
		5: "SPAN_KIND_INTERNAL",
	}
	SpanKind_value = map[string]int32{
		"SPAN_KIND_UNSPECIFIED": 0,
		"SPAN_KIND_SERVER":      1,
		"SPAN_KIND_CLIENT":      2,
		"SPAN_KIND_CONSUMER":    3,
		"SPAN_KIND_PRODUCER":    4,
		"SPAN_KIND_INTERNAL":    5,
	}
)

func (x SpanKind) Enum() *SpanKind {
	p := new(SpanKind)
	*p = x
	return p
}

func (x SpanKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SpanKind) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_catalog_v1_span_kind_proto_enumTypes[0].Descriptor()
}

func (SpanKind) Type() protoreflect.EnumType {
	return &file_com_coralogix_catalog_v1_span_kind_proto_enumTypes[0]
}

func (x SpanKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SpanKind.Descriptor instead.
func (SpanKind) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_catalog_v1_span_kind_proto_rawDescGZIP(), []int{0}
}

var File_com_coralogix_catalog_v1_span_kind_proto protoreflect.FileDescriptor

var file_com_coralogix_catalog_v1_span_kind_proto_rawDesc = []byte{
	0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x5f,
	0x6b, 0x69, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x2e, 0x76, 0x31, 0x2a, 0x99, 0x01, 0x0a, 0x08, 0x53, 0x70, 0x61, 0x6e, 0x4b, 0x69, 0x6e,
	0x64, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x50, 0x41, 0x4e, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10,
	0x53, 0x50, 0x41, 0x4e, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52,
	0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x50, 0x41, 0x4e, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f,
	0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x50, 0x41, 0x4e,
	0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x55, 0x4d, 0x45, 0x52, 0x10, 0x03,
	0x12, 0x16, 0x0a, 0x12, 0x53, 0x50, 0x41, 0x4e, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x50, 0x52,
	0x4f, 0x44, 0x55, 0x43, 0x45, 0x52, 0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x50, 0x41, 0x4e,
	0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x05,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_catalog_v1_span_kind_proto_rawDescOnce sync.Once
	file_com_coralogix_catalog_v1_span_kind_proto_rawDescData = file_com_coralogix_catalog_v1_span_kind_proto_rawDesc
)

func file_com_coralogix_catalog_v1_span_kind_proto_rawDescGZIP() []byte {
	file_com_coralogix_catalog_v1_span_kind_proto_rawDescOnce.Do(func() {
		file_com_coralogix_catalog_v1_span_kind_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_catalog_v1_span_kind_proto_rawDescData)
	})
	return file_com_coralogix_catalog_v1_span_kind_proto_rawDescData
}

var file_com_coralogix_catalog_v1_span_kind_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogix_catalog_v1_span_kind_proto_goTypes = []any{
	(SpanKind)(0), // 0: com.coralogix.catalog.v1.SpanKind
}
var file_com_coralogix_catalog_v1_span_kind_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogix_catalog_v1_span_kind_proto_init() }
func file_com_coralogix_catalog_v1_span_kind_proto_init() {
	if File_com_coralogix_catalog_v1_span_kind_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_catalog_v1_span_kind_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_catalog_v1_span_kind_proto_goTypes,
		DependencyIndexes: file_com_coralogix_catalog_v1_span_kind_proto_depIdxs,
		EnumInfos:         file_com_coralogix_catalog_v1_span_kind_proto_enumTypes,
	}.Build()
	File_com_coralogix_catalog_v1_span_kind_proto = out.File
	file_com_coralogix_catalog_v1_span_kind_proto_rawDesc = nil
	file_com_coralogix_catalog_v1_span_kind_proto_goTypes = nil
	file_com_coralogix_catalog_v1_span_kind_proto_depIdxs = nil
}
