// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.3
// source: com/coralogix/archive/v2/archive_target.proto

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

type ProviderType int32

const (
	ProviderType_PROVIDER_TYPE_UNSPECIFIED ProviderType = 0
	ProviderType_PROVIDER_TYPE_S3          ProviderType = 1
	ProviderType_PROVIDER_TYPE_IBM_COS     ProviderType = 2
)

// Enum value maps for ProviderType.
var (
	ProviderType_name = map[int32]string{
		0: "PROVIDER_TYPE_UNSPECIFIED",
		1: "PROVIDER_TYPE_S3",
		2: "PROVIDER_TYPE_IBM_COS",
	}
	ProviderType_value = map[string]int32{
		"PROVIDER_TYPE_UNSPECIFIED": 0,
		"PROVIDER_TYPE_S3":          1,
		"PROVIDER_TYPE_IBM_COS":     2,
	}
)

func (x ProviderType) Enum() *ProviderType {
	p := new(ProviderType)
	*p = x
	return p
}

func (x ProviderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProviderType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_archive_v2_archive_target_proto_enumTypes[0].Descriptor()
}

func (ProviderType) Type() protoreflect.EnumType {
	return &file_com_coralogix_archive_v2_archive_target_proto_enumTypes[0]
}

func (x ProviderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProviderType.Descriptor instead.
func (ProviderType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_archive_target_proto_rawDescGZIP(), []int{0}
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
	return file_com_coralogix_archive_v2_archive_target_proto_enumTypes[1].Descriptor()
}

func (PartitioningScheme) Type() protoreflect.EnumType {
	return &file_com_coralogix_archive_v2_archive_target_proto_enumTypes[1]
}

func (x PartitioningScheme) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PartitioningScheme.Descriptor instead.
func (PartitioningScheme) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_archive_target_proto_rawDescGZIP(), []int{1}
}

type ArchiveTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Format *Format `protobuf:"bytes,1,opt,name=format,proto3" json:"format,omitempty"`
	// Types that are assignable to PhysicalTarget:
	//
	//	*ArchiveTarget_S3
	//	*ArchiveTarget_IbmCos
	PhysicalTarget     isArchiveTarget_PhysicalTarget `protobuf_oneof:"physical_target"`
	EnableTags         bool                           `protobuf:"varint,4,opt,name=enable_tags,json=enableTags,proto3" json:"enable_tags,omitempty"`
	ProviderType       ProviderType                   `protobuf:"varint,5,opt,name=provider_type,json=providerType,proto3,enum=com.coralogix.archive.v2.ProviderType" json:"provider_type,omitempty"`
	Prefix             string                         `protobuf:"bytes,6,opt,name=prefix,proto3" json:"prefix,omitempty"`
	FilenamePrefix     *string                        `protobuf:"bytes,7,opt,name=filename_prefix,json=filenamePrefix,proto3,oneof" json:"filename_prefix,omitempty"`
	Tags               []*ObjectTag                   `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
	MaxCompactionLevel int32                          `protobuf:"varint,9,opt,name=max_compaction_level,json=maxCompactionLevel,proto3" json:"max_compaction_level,omitempty"`
	PartitioningScheme *PartitioningScheme            `protobuf:"varint,10,opt,name=partitioning_scheme,json=partitioningScheme,proto3,enum=com.coralogix.archive.v2.PartitioningScheme,oneof" json:"partitioning_scheme,omitempty"`
}

func (x *ArchiveTarget) Reset() {
	*x = ArchiveTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_archive_v2_archive_target_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArchiveTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArchiveTarget) ProtoMessage() {}

func (x *ArchiveTarget) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_archive_target_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArchiveTarget.ProtoReflect.Descriptor instead.
func (*ArchiveTarget) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_archive_target_proto_rawDescGZIP(), []int{0}
}

func (x *ArchiveTarget) GetFormat() *Format {
	if x != nil {
		return x.Format
	}
	return nil
}

func (m *ArchiveTarget) GetPhysicalTarget() isArchiveTarget_PhysicalTarget {
	if m != nil {
		return m.PhysicalTarget
	}
	return nil
}

func (x *ArchiveTarget) GetS3() *S3TargetSpec {
	if x, ok := x.GetPhysicalTarget().(*ArchiveTarget_S3); ok {
		return x.S3
	}
	return nil
}

func (x *ArchiveTarget) GetIbmCos() *IBMCosTargetSpec {
	if x, ok := x.GetPhysicalTarget().(*ArchiveTarget_IbmCos); ok {
		return x.IbmCos
	}
	return nil
}

func (x *ArchiveTarget) GetEnableTags() bool {
	if x != nil {
		return x.EnableTags
	}
	return false
}

func (x *ArchiveTarget) GetProviderType() ProviderType {
	if x != nil {
		return x.ProviderType
	}
	return ProviderType_PROVIDER_TYPE_UNSPECIFIED
}

func (x *ArchiveTarget) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *ArchiveTarget) GetFilenamePrefix() string {
	if x != nil && x.FilenamePrefix != nil {
		return *x.FilenamePrefix
	}
	return ""
}

func (x *ArchiveTarget) GetTags() []*ObjectTag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *ArchiveTarget) GetMaxCompactionLevel() int32 {
	if x != nil {
		return x.MaxCompactionLevel
	}
	return 0
}

func (x *ArchiveTarget) GetPartitioningScheme() PartitioningScheme {
	if x != nil && x.PartitioningScheme != nil {
		return *x.PartitioningScheme
	}
	return PartitioningScheme_PARTITIONING_SCHEME_UNSPECIFIED
}

type isArchiveTarget_PhysicalTarget interface {
	isArchiveTarget_PhysicalTarget()
}

type ArchiveTarget_S3 struct {
	S3 *S3TargetSpec `protobuf:"bytes,2,opt,name=s3,proto3,oneof"`
}

type ArchiveTarget_IbmCos struct {
	IbmCos *IBMCosTargetSpec `protobuf:"bytes,3,opt,name=ibm_cos,json=ibmCos,proto3,oneof"`
}

func (*ArchiveTarget_S3) isArchiveTarget_PhysicalTarget() {}

func (*ArchiveTarget_IbmCos) isArchiveTarget_PhysicalTarget() {}

var File_com_coralogix_archive_v2_archive_target_proto protoreflect.FileDescriptor

var file_com_coralogix_archive_v2_archive_target_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x1a, 0x25, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x2f, 0x76, 0x32, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x25, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x76,
	0x32, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8c, 0x05, 0x0a, 0x0d, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x12, 0x38, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x38,
	0x0a, 0x02, 0x73, 0x33, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x33, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x70,
	0x65, 0x63, 0x48, 0x00, 0x52, 0x02, 0x73, 0x33, 0x12, 0x45, 0x0a, 0x07, 0x69, 0x62, 0x6d, 0x5f,
	0x63, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x42, 0x4d, 0x43, 0x6f, 0x73, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x53, 0x70, 0x65, 0x63, 0x48, 0x00, 0x52, 0x06, 0x69, 0x62, 0x6d, 0x43, 0x6f, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x61, 0x67, 0x73,
	0x12, 0x4b, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e,
	0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x2c, 0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x0e, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x88, 0x01, 0x01, 0x12, 0x37, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x30, 0x0a, 0x14,
	0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x6d, 0x61, 0x78, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x62,
	0x0a, 0x13, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x69, 0x6e, 0x67, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x48, 0x02, 0x52, 0x12, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x11, 0x0a, 0x0f, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x65, 0x2a, 0x5e, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1d, 0x0a, 0x19, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x14, 0x0a, 0x10, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x53, 0x33, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44,
	0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x42, 0x4d, 0x5f, 0x43, 0x4f, 0x53, 0x10,
	0x02, 0x2a, 0x58, 0x0a, 0x12, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x69, 0x6e,
	0x67, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x1f, 0x50, 0x41, 0x52, 0x54, 0x49,
	0x54, 0x49, 0x4f, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19,
	0x50, 0x41, 0x52, 0x54, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x43, 0x48,
	0x45, 0x4d, 0x45, 0x5f, 0x44, 0x54, 0x5f, 0x48, 0x52, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_archive_v2_archive_target_proto_rawDescOnce sync.Once
	file_com_coralogix_archive_v2_archive_target_proto_rawDescData = file_com_coralogix_archive_v2_archive_target_proto_rawDesc
)

func file_com_coralogix_archive_v2_archive_target_proto_rawDescGZIP() []byte {
	file_com_coralogix_archive_v2_archive_target_proto_rawDescOnce.Do(func() {
		file_com_coralogix_archive_v2_archive_target_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_archive_v2_archive_target_proto_rawDescData)
	})
	return file_com_coralogix_archive_v2_archive_target_proto_rawDescData
}

var file_com_coralogix_archive_v2_archive_target_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_com_coralogix_archive_v2_archive_target_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogix_archive_v2_archive_target_proto_goTypes = []interface{}{
	(ProviderType)(0),        // 0: com.coralogix.archive.v2.ProviderType
	(PartitioningScheme)(0),  // 1: com.coralogix.archive.v2.PartitioningScheme
	(*ArchiveTarget)(nil),    // 2: com.coralogix.archive.v2.ArchiveTarget
	(*Format)(nil),           // 3: com.coralogix.archive.v2.Format
	(*S3TargetSpec)(nil),     // 4: com.coralogix.archive.v2.S3TargetSpec
	(*IBMCosTargetSpec)(nil), // 5: com.coralogix.archive.v2.IBMCosTargetSpec
	(*ObjectTag)(nil),        // 6: com.coralogix.archive.v2.ObjectTag
}
var file_com_coralogix_archive_v2_archive_target_proto_depIdxs = []int32{
	3, // 0: com.coralogix.archive.v2.ArchiveTarget.format:type_name -> com.coralogix.archive.v2.Format
	4, // 1: com.coralogix.archive.v2.ArchiveTarget.s3:type_name -> com.coralogix.archive.v2.S3TargetSpec
	5, // 2: com.coralogix.archive.v2.ArchiveTarget.ibm_cos:type_name -> com.coralogix.archive.v2.IBMCosTargetSpec
	0, // 3: com.coralogix.archive.v2.ArchiveTarget.provider_type:type_name -> com.coralogix.archive.v2.ProviderType
	6, // 4: com.coralogix.archive.v2.ArchiveTarget.tags:type_name -> com.coralogix.archive.v2.ObjectTag
	1, // 5: com.coralogix.archive.v2.ArchiveTarget.partitioning_scheme:type_name -> com.coralogix.archive.v2.PartitioningScheme
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_com_coralogix_archive_v2_archive_target_proto_init() }
func file_com_coralogix_archive_v2_archive_target_proto_init() {
	if File_com_coralogix_archive_v2_archive_target_proto != nil {
		return
	}
	file_com_coralogix_archive_v2_target_proto_init()
	file_com_coralogix_archive_v2_format_proto_init()
	file_com_coralogix_archive_v2_object_tag_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_archive_v2_archive_target_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArchiveTarget); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_com_coralogix_archive_v2_archive_target_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ArchiveTarget_S3)(nil),
		(*ArchiveTarget_IbmCos)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_archive_v2_archive_target_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_archive_v2_archive_target_proto_goTypes,
		DependencyIndexes: file_com_coralogix_archive_v2_archive_target_proto_depIdxs,
		EnumInfos:         file_com_coralogix_archive_v2_archive_target_proto_enumTypes,
		MessageInfos:      file_com_coralogix_archive_v2_archive_target_proto_msgTypes,
	}.Build()
	File_com_coralogix_archive_v2_archive_target_proto = out.File
	file_com_coralogix_archive_v2_archive_target_proto_rawDesc = nil
	file_com_coralogix_archive_v2_archive_target_proto_goTypes = nil
	file_com_coralogix_archive_v2_archive_target_proto_depIdxs = nil
}
