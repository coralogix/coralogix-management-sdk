// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.1
// source: com/coralogix/archive/dataset/v1/dataset.proto

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

type DatasetClass int32

const (
	DatasetClass_DATASET_CLASS_UNSPECIFIED        DatasetClass = 0
	DatasetClass_DATASET_CLASS_ENV_LEVEL_DATASET  DatasetClass = 1
	DatasetClass_DATASET_CLASS_TEAM_LEVEL_DATASET DatasetClass = 2
)

// Enum value maps for DatasetClass.
var (
	DatasetClass_name = map[int32]string{
		0: "DATASET_CLASS_UNSPECIFIED",
		1: "DATASET_CLASS_ENV_LEVEL_DATASET",
		2: "DATASET_CLASS_TEAM_LEVEL_DATASET",
	}
	DatasetClass_value = map[string]int32{
		"DATASET_CLASS_UNSPECIFIED":        0,
		"DATASET_CLASS_ENV_LEVEL_DATASET":  1,
		"DATASET_CLASS_TEAM_LEVEL_DATASET": 2,
	}
)

func (x DatasetClass) Enum() *DatasetClass {
	p := new(DatasetClass)
	*p = x
	return p
}

func (x DatasetClass) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DatasetClass) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_archive_dataset_v1_dataset_proto_enumTypes[0].Descriptor()
}

func (DatasetClass) Type() protoreflect.EnumType {
	return &file_com_coralogix_archive_dataset_v1_dataset_proto_enumTypes[0]
}

func (x DatasetClass) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DatasetClass.Descriptor instead.
func (DatasetClass) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v1_dataset_proto_rawDescGZIP(), []int{0}
}

type PartitioningScheme int32

const (
	PartitioningScheme_PARTITIONING_SCHEME_UNSPECIFIED PartitioningScheme = 0
	PartitioningScheme_PARTITIONING_SCHEME_DT_HR       PartitioningScheme = 1
)

// Enum value maps for PartitioningScheme.
var (
	PartitioningScheme_name = map[int32]string{
		0: "PARTITIONING_SCHEME_UNSPECIFIED",
		1: "PARTITIONING_SCHEME_DT_HR",
	}
	PartitioningScheme_value = map[string]int32{
		"PARTITIONING_SCHEME_UNSPECIFIED": 0,
		"PARTITIONING_SCHEME_DT_HR":       1,
	}
)

func (x PartitioningScheme) Enum() *PartitioningScheme {
	p := new(PartitioningScheme)
	*p = x
	return p
}

func (x PartitioningScheme) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PartitioningScheme) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_archive_dataset_v1_dataset_proto_enumTypes[1].Descriptor()
}

func (PartitioningScheme) Type() protoreflect.EnumType {
	return &file_com_coralogix_archive_dataset_v1_dataset_proto_enumTypes[1]
}

func (x PartitioningScheme) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PartitioningScheme.Descriptor instead.
func (PartitioningScheme) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v1_dataset_proto_rawDescGZIP(), []int{1}
}

type Dataset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NamespaceId        string             `protobuf:"bytes,1,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	Id                 string             `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	ClassId            DatasetClass       `protobuf:"varint,3,opt,name=class_id,json=classId,proto3,enum=com.coralogix.archive.dataset.v1.DatasetClass" json:"class_id,omitempty"`
	TeamId             uint64             `protobuf:"varint,4,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	PartitioningScheme PartitioningScheme `protobuf:"varint,5,opt,name=partitioning_scheme,json=partitioningScheme,proto3,enum=com.coralogix.archive.dataset.v1.PartitioningScheme" json:"partitioning_scheme,omitempty"`
	SplitByEventLabels []string           `protobuf:"bytes,6,rep,name=split_by_event_labels,json=splitByEventLabels,proto3" json:"split_by_event_labels,omitempty"`
	DataspaceId        string             `protobuf:"bytes,7,opt,name=dataspace_id,json=dataspaceId,proto3" json:"dataspace_id,omitempty"`
}

func (x *Dataset) Reset() {
	*x = Dataset{}
	mi := &file_com_coralogix_archive_dataset_v1_dataset_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Dataset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dataset) ProtoMessage() {}

func (x *Dataset) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_dataset_v1_dataset_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dataset.ProtoReflect.Descriptor instead.
func (*Dataset) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v1_dataset_proto_rawDescGZIP(), []int{0}
}

func (x *Dataset) GetNamespaceId() string {
	if x != nil {
		return x.NamespaceId
	}
	return ""
}

func (x *Dataset) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Dataset) GetClassId() DatasetClass {
	if x != nil {
		return x.ClassId
	}
	return DatasetClass_DATASET_CLASS_UNSPECIFIED
}

func (x *Dataset) GetTeamId() uint64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *Dataset) GetPartitioningScheme() PartitioningScheme {
	if x != nil {
		return x.PartitioningScheme
	}
	return PartitioningScheme_PARTITIONING_SCHEME_UNSPECIFIED
}

func (x *Dataset) GetSplitByEventLabels() []string {
	if x != nil {
		return x.SplitByEventLabels
	}
	return nil
}

func (x *Dataset) GetDataspaceId() string {
	if x != nil {
		return x.DataspaceId
	}
	return ""
}

var File_com_coralogix_archive_dataset_v1_dataset_proto protoreflect.FileDescriptor

var file_com_coralogix_archive_dataset_v1_dataset_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e,
	0x76, 0x31, 0x22, 0xdd, 0x02, 0x0a, 0x07, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x49, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x74,
	0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x65, 0x0a, 0x13, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x69,
	0x6e, 0x67, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x12, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x15,
	0x73, 0x70, 0x6c, 0x69, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x73, 0x70, 0x6c,
	0x69, 0x74, 0x42, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x49, 0x64, 0x2a, 0x78, 0x0a, 0x0c, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x41, 0x54, 0x41, 0x53, 0x45, 0x54, 0x5f, 0x43, 0x4c,
	0x41, 0x53, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x23, 0x0a, 0x1f, 0x44, 0x41, 0x54, 0x41, 0x53, 0x45, 0x54, 0x5f, 0x43, 0x4c, 0x41,
	0x53, 0x53, 0x5f, 0x45, 0x4e, 0x56, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x44, 0x41, 0x54,
	0x41, 0x53, 0x45, 0x54, 0x10, 0x01, 0x12, 0x24, 0x0a, 0x20, 0x44, 0x41, 0x54, 0x41, 0x53, 0x45,
	0x54, 0x5f, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x5f, 0x54, 0x45, 0x41, 0x4d, 0x5f, 0x4c, 0x45, 0x56,
	0x45, 0x4c, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x53, 0x45, 0x54, 0x10, 0x02, 0x2a, 0x58, 0x0a, 0x12,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x65, 0x12, 0x23, 0x0a, 0x1f, 0x50, 0x41, 0x52, 0x54, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x49,
	0x4e, 0x47, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x50, 0x41, 0x52, 0x54, 0x49,
	0x54, 0x49, 0x4f, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x45, 0x5f, 0x44,
	0x54, 0x5f, 0x48, 0x52, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_archive_dataset_v1_dataset_proto_rawDescOnce sync.Once
	file_com_coralogix_archive_dataset_v1_dataset_proto_rawDescData = file_com_coralogix_archive_dataset_v1_dataset_proto_rawDesc
)

func file_com_coralogix_archive_dataset_v1_dataset_proto_rawDescGZIP() []byte {
	file_com_coralogix_archive_dataset_v1_dataset_proto_rawDescOnce.Do(func() {
		file_com_coralogix_archive_dataset_v1_dataset_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_archive_dataset_v1_dataset_proto_rawDescData)
	})
	return file_com_coralogix_archive_dataset_v1_dataset_proto_rawDescData
}

var file_com_coralogix_archive_dataset_v1_dataset_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_com_coralogix_archive_dataset_v1_dataset_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogix_archive_dataset_v1_dataset_proto_goTypes = []any{
	(DatasetClass)(0),       // 0: com.coralogix.archive.dataset.v1.DatasetClass
	(PartitioningScheme)(0), // 1: com.coralogix.archive.dataset.v1.PartitioningScheme
	(*Dataset)(nil),         // 2: com.coralogix.archive.dataset.v1.Dataset
}
var file_com_coralogix_archive_dataset_v1_dataset_proto_depIdxs = []int32{
	0, // 0: com.coralogix.archive.dataset.v1.Dataset.class_id:type_name -> com.coralogix.archive.dataset.v1.DatasetClass
	1, // 1: com.coralogix.archive.dataset.v1.Dataset.partitioning_scheme:type_name -> com.coralogix.archive.dataset.v1.PartitioningScheme
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_com_coralogix_archive_dataset_v1_dataset_proto_init() }
func file_com_coralogix_archive_dataset_v1_dataset_proto_init() {
	if File_com_coralogix_archive_dataset_v1_dataset_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_archive_dataset_v1_dataset_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_archive_dataset_v1_dataset_proto_goTypes,
		DependencyIndexes: file_com_coralogix_archive_dataset_v1_dataset_proto_depIdxs,
		EnumInfos:         file_com_coralogix_archive_dataset_v1_dataset_proto_enumTypes,
		MessageInfos:      file_com_coralogix_archive_dataset_v1_dataset_proto_msgTypes,
	}.Build()
	File_com_coralogix_archive_dataset_v1_dataset_proto = out.File
	file_com_coralogix_archive_dataset_v1_dataset_proto_rawDesc = nil
	file_com_coralogix_archive_dataset_v1_dataset_proto_goTypes = nil
	file_com_coralogix_archive_dataset_v1_dataset_proto_depIdxs = nil
}
