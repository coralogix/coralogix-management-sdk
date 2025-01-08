// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.1
// source: com/coralogix/archive/format/csv/v1/csv.proto

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

type Version int32

const (
	Version_VERSION_UNSPECIFIED Version = 0
	Version_VERSION_V1          Version = 1
)

// Enum value maps for Version.
var (
	Version_name = map[int32]string{
		0: "VERSION_UNSPECIFIED",
		1: "VERSION_V1",
	}
	Version_value = map[string]int32{
		"VERSION_UNSPECIFIED": 0,
		"VERSION_V1":          1,
	}
)

func (x Version) Enum() *Version {
	p := new(Version)
	*p = x
	return p
}

func (x Version) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Version) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_archive_format_csv_v1_csv_proto_enumTypes[0].Descriptor()
}

func (Version) Type() protoreflect.EnumType {
	return &file_com_coralogix_archive_format_csv_v1_csv_proto_enumTypes[0]
}

func (x Version) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Version.Descriptor instead.
func (Version) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_archive_format_csv_v1_csv_proto_rawDescGZIP(), []int{0}
}

type Csv struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version Version `protobuf:"varint,1,opt,name=version,proto3,enum=com.coralogix.archive.format.csv.v1.Version" json:"version,omitempty"`
}

func (x *Csv) Reset() {
	*x = Csv{}
	mi := &file_com_coralogix_archive_format_csv_v1_csv_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Csv) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Csv) ProtoMessage() {}

func (x *Csv) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_format_csv_v1_csv_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Csv.ProtoReflect.Descriptor instead.
func (*Csv) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_format_csv_v1_csv_proto_rawDescGZIP(), []int{0}
}

func (x *Csv) GetVersion() Version {
	if x != nil {
		return x.Version
	}
	return Version_VERSION_UNSPECIFIED
}

var File_com_coralogix_archive_format_csv_v1_csv_proto protoreflect.FileDescriptor

var file_com_coralogix_archive_format_csv_v1_csv_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2f, 0x63,
	0x73, 0x76, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x73, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2e, 0x63, 0x73,
	0x76, 0x2e, 0x76, 0x31, 0x22, 0x4d, 0x0a, 0x03, 0x43, 0x73, 0x76, 0x12, 0x46, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x2e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2e, 0x63, 0x73, 0x76, 0x2e,
	0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x2a, 0x32, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x17,
	0x0a, 0x13, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x56, 0x45, 0x52, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x56, 0x31, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_archive_format_csv_v1_csv_proto_rawDescOnce sync.Once
	file_com_coralogix_archive_format_csv_v1_csv_proto_rawDescData = file_com_coralogix_archive_format_csv_v1_csv_proto_rawDesc
)

func file_com_coralogix_archive_format_csv_v1_csv_proto_rawDescGZIP() []byte {
	file_com_coralogix_archive_format_csv_v1_csv_proto_rawDescOnce.Do(func() {
		file_com_coralogix_archive_format_csv_v1_csv_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_archive_format_csv_v1_csv_proto_rawDescData)
	})
	return file_com_coralogix_archive_format_csv_v1_csv_proto_rawDescData
}

var file_com_coralogix_archive_format_csv_v1_csv_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogix_archive_format_csv_v1_csv_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogix_archive_format_csv_v1_csv_proto_goTypes = []any{
	(Version)(0), // 0: com.coralogix.archive.format.csv.v1.Version
	(*Csv)(nil),  // 1: com.coralogix.archive.format.csv.v1.Csv
}
var file_com_coralogix_archive_format_csv_v1_csv_proto_depIdxs = []int32{
	0, // 0: com.coralogix.archive.format.csv.v1.Csv.version:type_name -> com.coralogix.archive.format.csv.v1.Version
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_com_coralogix_archive_format_csv_v1_csv_proto_init() }
func file_com_coralogix_archive_format_csv_v1_csv_proto_init() {
	if File_com_coralogix_archive_format_csv_v1_csv_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_archive_format_csv_v1_csv_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_archive_format_csv_v1_csv_proto_goTypes,
		DependencyIndexes: file_com_coralogix_archive_format_csv_v1_csv_proto_depIdxs,
		EnumInfos:         file_com_coralogix_archive_format_csv_v1_csv_proto_enumTypes,
		MessageInfos:      file_com_coralogix_archive_format_csv_v1_csv_proto_msgTypes,
	}.Build()
	File_com_coralogix_archive_format_csv_v1_csv_proto = out.File
	file_com_coralogix_archive_format_csv_v1_csv_proto_rawDesc = nil
	file_com_coralogix_archive_format_csv_v1_csv_proto_goTypes = nil
	file_com_coralogix_archive_format_csv_v1_csv_proto_depIdxs = nil
}