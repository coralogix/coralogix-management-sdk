// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.28.1
// source: com/coralogix/archive/v2/target.proto

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

type IbmBucketType int32

const (
	IbmBucketType_IBM_BUCKET_TYPE_UNSPECIFIED IbmBucketType = 0
	IbmBucketType_IBM_BUCKET_TYPE_EXTERNAL    IbmBucketType = 1
	IbmBucketType_IBM_BUCKET_TYPE_INTERNAL    IbmBucketType = 2
)

// Enum value maps for IbmBucketType.
var (
	IbmBucketType_name = map[int32]string{
		0: "IBM_BUCKET_TYPE_UNSPECIFIED",
		1: "IBM_BUCKET_TYPE_EXTERNAL",
		2: "IBM_BUCKET_TYPE_INTERNAL",
	}
	IbmBucketType_value = map[string]int32{
		"IBM_BUCKET_TYPE_UNSPECIFIED": 0,
		"IBM_BUCKET_TYPE_EXTERNAL":    1,
		"IBM_BUCKET_TYPE_INTERNAL":    2,
	}
)

func (x IbmBucketType) Enum() *IbmBucketType {
	p := new(IbmBucketType)
	*p = x
	return p
}

func (x IbmBucketType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IbmBucketType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_archive_v2_target_proto_enumTypes[0].Descriptor()
}

func (IbmBucketType) Type() protoreflect.EnumType {
	return &file_com_coralogix_archive_v2_target_proto_enumTypes[0]
}

func (x IbmBucketType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IbmBucketType.Descriptor instead.
func (IbmBucketType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_target_proto_rawDescGZIP(), []int{0}
}

type IBMCosTargetSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BucketCrn  string         `protobuf:"bytes,1,opt,name=bucket_crn,json=bucketCrn,proto3" json:"bucket_crn,omitempty"`
	Endpoint   string         `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	ServiceCrn *string        `protobuf:"bytes,3,opt,name=service_crn,json=serviceCrn,proto3,oneof" json:"service_crn,omitempty"`
	BucketType *IbmBucketType `protobuf:"varint,4,opt,name=bucket_type,json=bucketType,proto3,enum=com.coralogix.archive.v2.IbmBucketType,oneof" json:"bucket_type,omitempty"`
}

func (x *IBMCosTargetSpec) Reset() {
	*x = IBMCosTargetSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_archive_v2_target_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IBMCosTargetSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IBMCosTargetSpec) ProtoMessage() {}

func (x *IBMCosTargetSpec) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_target_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IBMCosTargetSpec.ProtoReflect.Descriptor instead.
func (*IBMCosTargetSpec) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_target_proto_rawDescGZIP(), []int{0}
}

func (x *IBMCosTargetSpec) GetBucketCrn() string {
	if x != nil {
		return x.BucketCrn
	}
	return ""
}

func (x *IBMCosTargetSpec) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *IBMCosTargetSpec) GetServiceCrn() string {
	if x != nil && x.ServiceCrn != nil {
		return *x.ServiceCrn
	}
	return ""
}

func (x *IBMCosTargetSpec) GetBucketType() IbmBucketType {
	if x != nil && x.BucketType != nil {
		return *x.BucketType
	}
	return IbmBucketType_IBM_BUCKET_TYPE_UNSPECIFIED
}

type S3TargetSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bucket string  `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Region *string `protobuf:"bytes,2,opt,name=region,proto3,oneof" json:"region,omitempty"`
}

func (x *S3TargetSpec) Reset() {
	*x = S3TargetSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_archive_v2_target_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3TargetSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3TargetSpec) ProtoMessage() {}

func (x *S3TargetSpec) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_target_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3TargetSpec.ProtoReflect.Descriptor instead.
func (*S3TargetSpec) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_target_proto_rawDescGZIP(), []int{1}
}

func (x *S3TargetSpec) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *S3TargetSpec) GetRegion() string {
	if x != nil && x.Region != nil {
		return *x.Region
	}
	return ""
}

type ArchiveSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArchivingFormatId string `protobuf:"bytes,1,opt,name=archiving_format_id,json=archivingFormatId,proto3" json:"archiving_format_id,omitempty"`
	IsActive          bool   `protobuf:"varint,2,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	EnableTags        bool   `protobuf:"varint,3,opt,name=enable_tags,json=enableTags,proto3" json:"enable_tags,omitempty"`
}

func (x *ArchiveSpec) Reset() {
	*x = ArchiveSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_archive_v2_target_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArchiveSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArchiveSpec) ProtoMessage() {}

func (x *ArchiveSpec) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_target_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArchiveSpec.ProtoReflect.Descriptor instead.
func (*ArchiveSpec) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_target_proto_rawDescGZIP(), []int{2}
}

func (x *ArchiveSpec) GetArchivingFormatId() string {
	if x != nil {
		return x.ArchivingFormatId
	}
	return ""
}

func (x *ArchiveSpec) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *ArchiveSpec) GetEnableTags() bool {
	if x != nil {
		return x.EnableTags
	}
	return false
}

type Target struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to TargetSpec:
	//
	//	*Target_S3
	//	*Target_IbmCos
	TargetSpec  isTarget_TargetSpec `protobuf_oneof:"target_spec"`
	ArchiveSpec *ArchiveSpec        `protobuf:"bytes,3,opt,name=archive_spec,json=archiveSpec,proto3" json:"archive_spec,omitempty"`
}

func (x *Target) Reset() {
	*x = Target{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_archive_v2_target_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Target) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Target) ProtoMessage() {}

func (x *Target) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_target_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Target.ProtoReflect.Descriptor instead.
func (*Target) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_target_proto_rawDescGZIP(), []int{3}
}

func (m *Target) GetTargetSpec() isTarget_TargetSpec {
	if m != nil {
		return m.TargetSpec
	}
	return nil
}

func (x *Target) GetS3() *S3TargetSpec {
	if x, ok := x.GetTargetSpec().(*Target_S3); ok {
		return x.S3
	}
	return nil
}

func (x *Target) GetIbmCos() *IBMCosTargetSpec {
	if x, ok := x.GetTargetSpec().(*Target_IbmCos); ok {
		return x.IbmCos
	}
	return nil
}

func (x *Target) GetArchiveSpec() *ArchiveSpec {
	if x != nil {
		return x.ArchiveSpec
	}
	return nil
}

type isTarget_TargetSpec interface {
	isTarget_TargetSpec()
}

type Target_S3 struct {
	S3 *S3TargetSpec `protobuf:"bytes,1,opt,name=s3,proto3,oneof"`
}

type Target_IbmCos struct {
	IbmCos *IBMCosTargetSpec `protobuf:"bytes,2,opt,name=ibm_cos,json=ibmCos,proto3,oneof"`
}

func (*Target_S3) isTarget_TargetSpec() {}

func (*Target_IbmCos) isTarget_TargetSpec() {}

var File_com_coralogix_archive_v2_target_proto protoreflect.FileDescriptor

var file_com_coralogix_archive_v2_target_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76,
	0x32, 0x22, 0xe2, 0x01, 0x0a, 0x10, 0x49, 0x42, 0x4d, 0x43, 0x6f, 0x73, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x5f, 0x63, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x43, 0x72, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x24, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x72, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x43, 0x72, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x4d, 0x0a, 0x0b, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x62, 0x6d, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x48, 0x01, 0x52, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x63, 0x72, 0x6e, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x4e, 0x0a, 0x0c, 0x53, 0x33, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1b,
	0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0x7b, 0x0a, 0x0b, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x2e, 0x0a, 0x13, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69,
	0x6e, 0x67, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x46, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x61, 0x67,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x54,
	0x61, 0x67, 0x73, 0x22, 0xe2, 0x01, 0x0a, 0x06, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x38,
	0x0a, 0x02, 0x73, 0x33, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x33, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x70,
	0x65, 0x63, 0x48, 0x00, 0x52, 0x02, 0x73, 0x33, 0x12, 0x45, 0x0a, 0x07, 0x69, 0x62, 0x6d, 0x5f,
	0x63, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x42, 0x4d, 0x43, 0x6f, 0x73, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x53, 0x70, 0x65, 0x63, 0x48, 0x00, 0x52, 0x06, 0x69, 0x62, 0x6d, 0x43, 0x6f, 0x73, 0x12,
	0x48, 0x0a, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32,
	0x2e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0b, 0x61, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x53, 0x70, 0x65, 0x63, 0x42, 0x0d, 0x0a, 0x0b, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x2a, 0x6c, 0x0a, 0x0d, 0x49, 0x62, 0x6d, 0x42,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x49, 0x42, 0x4d,
	0x5f, 0x42, 0x55, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x49, 0x42,
	0x4d, 0x5f, 0x42, 0x55, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x58,
	0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x49, 0x42, 0x4d, 0x5f,
	0x42, 0x55, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45,
	0x52, 0x4e, 0x41, 0x4c, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_archive_v2_target_proto_rawDescOnce sync.Once
	file_com_coralogix_archive_v2_target_proto_rawDescData = file_com_coralogix_archive_v2_target_proto_rawDesc
)

func file_com_coralogix_archive_v2_target_proto_rawDescGZIP() []byte {
	file_com_coralogix_archive_v2_target_proto_rawDescOnce.Do(func() {
		file_com_coralogix_archive_v2_target_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_archive_v2_target_proto_rawDescData)
	})
	return file_com_coralogix_archive_v2_target_proto_rawDescData
}

var file_com_coralogix_archive_v2_target_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogix_archive_v2_target_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_com_coralogix_archive_v2_target_proto_goTypes = []interface{}{
	(IbmBucketType)(0),       // 0: com.coralogix.archive.v2.IbmBucketType
	(*IBMCosTargetSpec)(nil), // 1: com.coralogix.archive.v2.IBMCosTargetSpec
	(*S3TargetSpec)(nil),     // 2: com.coralogix.archive.v2.S3TargetSpec
	(*ArchiveSpec)(nil),      // 3: com.coralogix.archive.v2.ArchiveSpec
	(*Target)(nil),           // 4: com.coralogix.archive.v2.Target
}
var file_com_coralogix_archive_v2_target_proto_depIdxs = []int32{
	0, // 0: com.coralogix.archive.v2.IBMCosTargetSpec.bucket_type:type_name -> com.coralogix.archive.v2.IbmBucketType
	2, // 1: com.coralogix.archive.v2.Target.s3:type_name -> com.coralogix.archive.v2.S3TargetSpec
	1, // 2: com.coralogix.archive.v2.Target.ibm_cos:type_name -> com.coralogix.archive.v2.IBMCosTargetSpec
	3, // 3: com.coralogix.archive.v2.Target.archive_spec:type_name -> com.coralogix.archive.v2.ArchiveSpec
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_com_coralogix_archive_v2_target_proto_init() }
func file_com_coralogix_archive_v2_target_proto_init() {
	if File_com_coralogix_archive_v2_target_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_archive_v2_target_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IBMCosTargetSpec); i {
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
		file_com_coralogix_archive_v2_target_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3TargetSpec); i {
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
		file_com_coralogix_archive_v2_target_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArchiveSpec); i {
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
		file_com_coralogix_archive_v2_target_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Target); i {
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
	file_com_coralogix_archive_v2_target_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_com_coralogix_archive_v2_target_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_com_coralogix_archive_v2_target_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Target_S3)(nil),
		(*Target_IbmCos)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_archive_v2_target_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_archive_v2_target_proto_goTypes,
		DependencyIndexes: file_com_coralogix_archive_v2_target_proto_depIdxs,
		EnumInfos:         file_com_coralogix_archive_v2_target_proto_enumTypes,
		MessageInfos:      file_com_coralogix_archive_v2_target_proto_msgTypes,
	}.Build()
	File_com_coralogix_archive_v2_target_proto = out.File
	file_com_coralogix_archive_v2_target_proto_rawDesc = nil
	file_com_coralogix_archive_v2_target_proto_goTypes = nil
	file_com_coralogix_archive_v2_target_proto_depIdxs = nil
}
