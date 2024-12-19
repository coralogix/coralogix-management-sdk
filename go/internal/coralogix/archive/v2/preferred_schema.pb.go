// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.26.1
// source: com/coralogix/archive/dataset/v2/preferred_schema.proto

package v2

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

type DataType int32

const (
	DataType_DATA_TYPE_UNSPECIFIED     DataType = 0
	DataType_DATA_TYPE_NUM             DataType = 1
	DataType_DATA_TYPE_STRING          DataType = 2
	DataType_DATA_TYPE_BOOL            DataType = 3
	DataType_DATA_TYPE_UNTRACKED_ARRAY DataType = 4
	DataType_DATA_TYPE_OBJECT          DataType = 5
	DataType_DATA_TYPE_NUM_ARRAY       DataType = 6
	DataType_DATA_TYPE_STRING_ARRAY    DataType = 7
	DataType_DATA_TYPE_BOOL_ARRAY      DataType = 8
	DataType_DATA_TYPE_ARRAY_ARRAY     DataType = 9
	DataType_DATA_TYPE_OBJECT_ARRAY    DataType = 10
	DataType_DATA_TYPE_NULL            DataType = 11
	DataType_DATA_TYPE_NULL_ARRAY      DataType = 12
	DataType_DATA_TYPE_EMPTY_ARRAY     DataType = 13
)

// Enum value maps for DataType.
var (
	DataType_name = map[int32]string{
		0:  "DATA_TYPE_UNSPECIFIED",
		1:  "DATA_TYPE_NUM",
		2:  "DATA_TYPE_STRING",
		3:  "DATA_TYPE_BOOL",
		4:  "DATA_TYPE_UNTRACKED_ARRAY",
		5:  "DATA_TYPE_OBJECT",
		6:  "DATA_TYPE_NUM_ARRAY",
		7:  "DATA_TYPE_STRING_ARRAY",
		8:  "DATA_TYPE_BOOL_ARRAY",
		9:  "DATA_TYPE_ARRAY_ARRAY",
		10: "DATA_TYPE_OBJECT_ARRAY",
		11: "DATA_TYPE_NULL",
		12: "DATA_TYPE_NULL_ARRAY",
		13: "DATA_TYPE_EMPTY_ARRAY",
	}
	DataType_value = map[string]int32{
		"DATA_TYPE_UNSPECIFIED":     0,
		"DATA_TYPE_NUM":             1,
		"DATA_TYPE_STRING":          2,
		"DATA_TYPE_BOOL":            3,
		"DATA_TYPE_UNTRACKED_ARRAY": 4,
		"DATA_TYPE_OBJECT":          5,
		"DATA_TYPE_NUM_ARRAY":       6,
		"DATA_TYPE_STRING_ARRAY":    7,
		"DATA_TYPE_BOOL_ARRAY":      8,
		"DATA_TYPE_ARRAY_ARRAY":     9,
		"DATA_TYPE_OBJECT_ARRAY":    10,
		"DATA_TYPE_NULL":            11,
		"DATA_TYPE_NULL_ARRAY":      12,
		"DATA_TYPE_EMPTY_ARRAY":     13,
	}
)

func (x DataType) Enum() *DataType {
	p := new(DataType)
	*p = x
	return p
}

func (x DataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_archive_dataset_v2_preferred_schema_proto_enumTypes[0].Descriptor()
}

func (DataType) Type() protoreflect.EnumType {
	return &file_com_coralogix_archive_dataset_v2_preferred_schema_proto_enumTypes[0]
}

func (x DataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataType.Descriptor instead.
func (DataType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_preferred_schema_proto_rawDescGZIP(), []int{0}
}

type NamedSchemaTemplateField struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	PathArray       []string               `protobuf:"bytes,1,rep,name=path_array,json=pathArray,proto3" json:"path_array,omitempty"`
	DataType        DataType               `protobuf:"varint,2,opt,name=data_type,json=dataType,proto3,enum=com.coralogix.archive.dataset.v2.DataType" json:"data_type,omitempty"`
	LogicalDataType string                 `protobuf:"bytes,3,opt,name=logical_data_type,json=logicalDataType,proto3" json:"logical_data_type,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *NamedSchemaTemplateField) Reset() {
	*x = NamedSchemaTemplateField{}
	mi := &file_com_coralogix_archive_dataset_v2_preferred_schema_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NamedSchemaTemplateField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamedSchemaTemplateField) ProtoMessage() {}

func (x *NamedSchemaTemplateField) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_dataset_v2_preferred_schema_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamedSchemaTemplateField.ProtoReflect.Descriptor instead.
func (*NamedSchemaTemplateField) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_preferred_schema_proto_rawDescGZIP(), []int{0}
}

func (x *NamedSchemaTemplateField) GetPathArray() []string {
	if x != nil {
		return x.PathArray
	}
	return nil
}

func (x *NamedSchemaTemplateField) GetDataType() DataType {
	if x != nil {
		return x.DataType
	}
	return DataType_DATA_TYPE_UNSPECIFIED
}

func (x *NamedSchemaTemplateField) GetLogicalDataType() string {
	if x != nil {
		return x.LogicalDataType
	}
	return ""
}

var File_com_coralogix_archive_dataset_v2_preferred_schema_proto protoreflect.FileDescriptor

var file_com_coralogix_archive_dataset_v2_preferred_schema_proto_rawDesc = []byte{
	0x0a, 0x37, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2f,
	0x76, 0x32, 0x2f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x22, 0xae, 0x01, 0x0a, 0x18,
	0x4e, 0x61, 0x6d, 0x65, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x68,
	0x5f, 0x61, 0x72, 0x72, 0x61, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61,
	0x74, 0x68, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x47, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6c, 0x6f, 0x67,
	0x69, 0x63, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x2a, 0xe6, 0x02, 0x0a,
	0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x41, 0x54,
	0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4e, 0x55, 0x4d, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x44, 0x41, 0x54, 0x41, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x12, 0x0a,
	0x0e, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4c, 0x10,
	0x03, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x45, 0x44, 0x5f, 0x41, 0x52, 0x52, 0x41, 0x59, 0x10, 0x04,
	0x12, 0x14, 0x0a, 0x10, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x42,
	0x4a, 0x45, 0x43, 0x54, 0x10, 0x05, 0x12, 0x17, 0x0a, 0x13, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4e, 0x55, 0x4d, 0x5f, 0x41, 0x52, 0x52, 0x41, 0x59, 0x10, 0x06, 0x12,
	0x1a, 0x0a, 0x16, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x52,
	0x49, 0x4e, 0x47, 0x5f, 0x41, 0x52, 0x52, 0x41, 0x59, 0x10, 0x07, 0x12, 0x18, 0x0a, 0x14, 0x44,
	0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4c, 0x5f, 0x41, 0x52,
	0x52, 0x41, 0x59, 0x10, 0x08, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x41, 0x52, 0x52, 0x41, 0x59, 0x5f, 0x41, 0x52, 0x52, 0x41, 0x59, 0x10, 0x09,
	0x12, 0x1a, 0x0a, 0x16, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x42,
	0x4a, 0x45, 0x43, 0x54, 0x5f, 0x41, 0x52, 0x52, 0x41, 0x59, 0x10, 0x0a, 0x12, 0x12, 0x0a, 0x0e,
	0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x55, 0x4c, 0x4c, 0x10, 0x0b,
	0x12, 0x18, 0x0a, 0x14, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x55,
	0x4c, 0x4c, 0x5f, 0x41, 0x52, 0x52, 0x41, 0x59, 0x10, 0x0c, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x41,
	0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x5f, 0x41, 0x52,
	0x52, 0x41, 0x59, 0x10, 0x0d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_archive_dataset_v2_preferred_schema_proto_rawDescOnce sync.Once
	file_com_coralogix_archive_dataset_v2_preferred_schema_proto_rawDescData = file_com_coralogix_archive_dataset_v2_preferred_schema_proto_rawDesc
)

func file_com_coralogix_archive_dataset_v2_preferred_schema_proto_rawDescGZIP() []byte {
	file_com_coralogix_archive_dataset_v2_preferred_schema_proto_rawDescOnce.Do(func() {
		file_com_coralogix_archive_dataset_v2_preferred_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_archive_dataset_v2_preferred_schema_proto_rawDescData)
	})
	return file_com_coralogix_archive_dataset_v2_preferred_schema_proto_rawDescData
}

var file_com_coralogix_archive_dataset_v2_preferred_schema_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogix_archive_dataset_v2_preferred_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogix_archive_dataset_v2_preferred_schema_proto_goTypes = []any{
	(DataType)(0),                    // 0: com.coralogix.archive.dataset.v2.DataType
	(*NamedSchemaTemplateField)(nil), // 1: com.coralogix.archive.dataset.v2.NamedSchemaTemplateField
}
var file_com_coralogix_archive_dataset_v2_preferred_schema_proto_depIdxs = []int32{
	0, // 0: com.coralogix.archive.dataset.v2.NamedSchemaTemplateField.data_type:type_name -> com.coralogix.archive.dataset.v2.DataType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_com_coralogix_archive_dataset_v2_preferred_schema_proto_init() }
func file_com_coralogix_archive_dataset_v2_preferred_schema_proto_init() {
	if File_com_coralogix_archive_dataset_v2_preferred_schema_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_archive_dataset_v2_preferred_schema_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_archive_dataset_v2_preferred_schema_proto_goTypes,
		DependencyIndexes: file_com_coralogix_archive_dataset_v2_preferred_schema_proto_depIdxs,
		EnumInfos:         file_com_coralogix_archive_dataset_v2_preferred_schema_proto_enumTypes,
		MessageInfos:      file_com_coralogix_archive_dataset_v2_preferred_schema_proto_msgTypes,
	}.Build()
	File_com_coralogix_archive_dataset_v2_preferred_schema_proto = out.File
	file_com_coralogix_archive_dataset_v2_preferred_schema_proto_rawDesc = nil
	file_com_coralogix_archive_dataset_v2_preferred_schema_proto_goTypes = nil
	file_com_coralogix_archive_dataset_v2_preferred_schema_proto_depIdxs = nil
}
