// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: com/coralogix/archive/dataset/v2/physical_location.proto

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

type ObjectStoreLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Bucket:
	//
	//	*ObjectStoreLocation_S3
	//	*ObjectStoreLocation_IbmCos
	Bucket             isObjectStoreLocation_Bucket `protobuf_oneof:"bucket"`
	Prefix             string                       `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
	EnableTags         bool                         `protobuf:"varint,4,opt,name=enable_tags,json=enableTags,proto3" json:"enable_tags,omitempty"`
	Tags               []*ObjectTag                 `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	PartitioningScheme *PartitioningScheme          `protobuf:"varint,6,opt,name=partitioning_scheme,json=partitioningScheme,proto3,enum=com.coralogix.archive.v2.PartitioningScheme,oneof" json:"partitioning_scheme,omitempty"`
}

func (x *ObjectStoreLocation) Reset() {
	*x = ObjectStoreLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_archive_dataset_v2_physical_location_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectStoreLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectStoreLocation) ProtoMessage() {}

func (x *ObjectStoreLocation) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_dataset_v2_physical_location_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectStoreLocation.ProtoReflect.Descriptor instead.
func (*ObjectStoreLocation) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_physical_location_proto_rawDescGZIP(), []int{0}
}

func (m *ObjectStoreLocation) GetBucket() isObjectStoreLocation_Bucket {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (x *ObjectStoreLocation) GetS3() *S3TargetSpec {
	if x, ok := x.GetBucket().(*ObjectStoreLocation_S3); ok {
		return x.S3
	}
	return nil
}

func (x *ObjectStoreLocation) GetIbmCos() *IBMCosTargetSpec {
	if x, ok := x.GetBucket().(*ObjectStoreLocation_IbmCos); ok {
		return x.IbmCos
	}
	return nil
}

func (x *ObjectStoreLocation) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *ObjectStoreLocation) GetEnableTags() bool {
	if x != nil {
		return x.EnableTags
	}
	return false
}

func (x *ObjectStoreLocation) GetTags() []*ObjectTag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *ObjectStoreLocation) GetPartitioningScheme() PartitioningScheme {
	if x != nil && x.PartitioningScheme != nil {
		return *x.PartitioningScheme
	}
	return PartitioningScheme_PARTITIONING_SCHEME_UNSPECIFIED
}

type isObjectStoreLocation_Bucket interface {
	isObjectStoreLocation_Bucket()
}

type ObjectStoreLocation_S3 struct {
	S3 *S3TargetSpec `protobuf:"bytes,1,opt,name=s3,proto3,oneof"`
}

type ObjectStoreLocation_IbmCos struct {
	IbmCos *IBMCosTargetSpec `protobuf:"bytes,2,opt,name=ibm_cos,json=ibmCos,proto3,oneof"`
}

func (*ObjectStoreLocation_S3) isObjectStoreLocation_Bucket() {}

func (*ObjectStoreLocation_IbmCos) isObjectStoreLocation_Bucket() {}

// Publicly known location spec. Customers could use this model in future "set location" API.
type LocationSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	IsActive bool    `protobuf:"varint,2,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	// Types that are assignable to Location:
	//
	//	*LocationSpec_ObjectStoreLocation
	Location isLocationSpec_Location `protobuf_oneof:"location"`
	Format   *Format                 `protobuf:"bytes,4,opt,name=format,proto3" json:"format,omitempty"`
}

func (x *LocationSpec) Reset() {
	*x = LocationSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_archive_dataset_v2_physical_location_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationSpec) ProtoMessage() {}

func (x *LocationSpec) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_dataset_v2_physical_location_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationSpec.ProtoReflect.Descriptor instead.
func (*LocationSpec) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_physical_location_proto_rawDescGZIP(), []int{1}
}

func (x *LocationSpec) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *LocationSpec) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (m *LocationSpec) GetLocation() isLocationSpec_Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (x *LocationSpec) GetObjectStoreLocation() *ObjectStoreLocation {
	if x, ok := x.GetLocation().(*LocationSpec_ObjectStoreLocation); ok {
		return x.ObjectStoreLocation
	}
	return nil
}

func (x *LocationSpec) GetFormat() *Format {
	if x != nil {
		return x.Format
	}
	return nil
}

type isLocationSpec_Location interface {
	isLocationSpec_Location()
}

type LocationSpec_ObjectStoreLocation struct {
	ObjectStoreLocation *ObjectStoreLocation `protobuf:"bytes,3,opt,name=object_store_location,json=objectStoreLocation,proto3,oneof"`
}

func (*LocationSpec_ObjectStoreLocation) isLocationSpec_Location() {}

// Internal fields that shouldn't be known outside ingestion pipeline.
// Mostly properties that can be set only for predefined datasets.
type IngestionLocationSpecProps struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxCompactionLevel int32    `protobuf:"varint,1,opt,name=max_compaction_level,json=maxCompactionLevel,proto3" json:"max_compaction_level,omitempty"`
	CollectLabels      bool     `protobuf:"varint,2,opt,name=collect_labels,json=collectLabels,proto3" json:"collect_labels,omitempty"`
	SplitByEventLabels []string `protobuf:"bytes,3,rep,name=split_by_event_labels,json=splitByEventLabels,proto3" json:"split_by_event_labels,omitempty"`
	// Deprecated: Marked as deprecated in com/coralogix/archive/dataset/v2/physical_location.proto.
	PartitioningScheme *PartitioningScheme `protobuf:"varint,4,opt,name=partitioning_scheme,json=partitioningScheme,proto3,enum=com.coralogix.archive.v2.PartitioningScheme,oneof" json:"partitioning_scheme,omitempty"`
}

func (x *IngestionLocationSpecProps) Reset() {
	*x = IngestionLocationSpecProps{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_archive_dataset_v2_physical_location_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngestionLocationSpecProps) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngestionLocationSpecProps) ProtoMessage() {}

func (x *IngestionLocationSpecProps) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_dataset_v2_physical_location_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngestionLocationSpecProps.ProtoReflect.Descriptor instead.
func (*IngestionLocationSpecProps) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_physical_location_proto_rawDescGZIP(), []int{2}
}

func (x *IngestionLocationSpecProps) GetMaxCompactionLevel() int32 {
	if x != nil {
		return x.MaxCompactionLevel
	}
	return 0
}

func (x *IngestionLocationSpecProps) GetCollectLabels() bool {
	if x != nil {
		return x.CollectLabels
	}
	return false
}

func (x *IngestionLocationSpecProps) GetSplitByEventLabels() []string {
	if x != nil {
		return x.SplitByEventLabels
	}
	return nil
}

// Deprecated: Marked as deprecated in com/coralogix/archive/dataset/v2/physical_location.proto.
func (x *IngestionLocationSpecProps) GetPartitioningScheme() PartitioningScheme {
	if x != nil && x.PartitioningScheme != nil {
		return *x.PartitioningScheme
	}
	return PartitioningScheme_PARTITIONING_SCHEME_UNSPECIFIED
}

// Specifies where specific dataset data can be found at the location.
type FinalLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spec *LocationSpec `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	// In case of data saved to S3 - base prefix of the dataset inside the bucket, without the dynamic dt/hr part
	FinalLocationBasePrefix *string `protobuf:"bytes,2,opt,name=final_location_base_prefix,json=finalLocationBasePrefix,proto3,oneof" json:"final_location_base_prefix,omitempty"`
}

func (x *FinalLocation) Reset() {
	*x = FinalLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_archive_dataset_v2_physical_location_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinalLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinalLocation) ProtoMessage() {}

func (x *FinalLocation) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_dataset_v2_physical_location_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinalLocation.ProtoReflect.Descriptor instead.
func (*FinalLocation) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_physical_location_proto_rawDescGZIP(), []int{3}
}

func (x *FinalLocation) GetSpec() *LocationSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *FinalLocation) GetFinalLocationBasePrefix() string {
	if x != nil && x.FinalLocationBasePrefix != nil {
		return *x.FinalLocationBasePrefix
	}
	return ""
}

// Intended for internal use inside ingestion pipeline. Only part that may change after dataset-ingress.
// As of now fully filled out by stage-writer, but if there will be more changes to location properties through the pipeline, they will go here too
type IngestionFinalLocationProps struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spec                    *IngestionLocationSpecProps `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	FilenamePrefix          *string                     `protobuf:"bytes,2,opt,name=filename_prefix,json=filenamePrefix,proto3,oneof" json:"filename_prefix,omitempty"`
	Tags                    []*ObjectTag                `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	FinalLocationFullPrefix *string                     `protobuf:"bytes,4,opt,name=final_location_full_prefix,json=finalLocationFullPrefix,proto3,oneof" json:"final_location_full_prefix,omitempty"` // includes the dynamic dt/hr part (if there is one)
}

func (x *IngestionFinalLocationProps) Reset() {
	*x = IngestionFinalLocationProps{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_archive_dataset_v2_physical_location_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngestionFinalLocationProps) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngestionFinalLocationProps) ProtoMessage() {}

func (x *IngestionFinalLocationProps) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_dataset_v2_physical_location_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngestionFinalLocationProps.ProtoReflect.Descriptor instead.
func (*IngestionFinalLocationProps) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_physical_location_proto_rawDescGZIP(), []int{4}
}

func (x *IngestionFinalLocationProps) GetSpec() *IngestionLocationSpecProps {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *IngestionFinalLocationProps) GetFilenamePrefix() string {
	if x != nil && x.FilenamePrefix != nil {
		return *x.FilenamePrefix
	}
	return ""
}

func (x *IngestionFinalLocationProps) GetTags() []*ObjectTag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *IngestionFinalLocationProps) GetFinalLocationFullPrefix() string {
	if x != nil && x.FinalLocationFullPrefix != nil {
		return *x.FinalLocationFullPrefix
	}
	return ""
}

// LocationSpec + internal fields needed for ingestion (dataset-ingress)
type IngestionLocationSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location *LocationSpec               `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Internal *IngestionLocationSpecProps `protobuf:"bytes,2,opt,name=internal,proto3,oneof" json:"internal,omitempty"`
}

func (x *IngestionLocationSpec) Reset() {
	*x = IngestionLocationSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_archive_dataset_v2_physical_location_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngestionLocationSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngestionLocationSpec) ProtoMessage() {}

func (x *IngestionLocationSpec) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_dataset_v2_physical_location_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngestionLocationSpec.ProtoReflect.Descriptor instead.
func (*IngestionLocationSpec) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_physical_location_proto_rawDescGZIP(), []int{5}
}

func (x *IngestionLocationSpec) GetLocation() *LocationSpec {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *IngestionLocationSpec) GetInternal() *IngestionLocationSpecProps {
	if x != nil {
		return x.Internal
	}
	return nil
}

// Goes through the whole ingestion pipeline, and is used externally for metastores
type IngestionFinalLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location *FinalLocation               `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Internal *IngestionFinalLocationProps `protobuf:"bytes,2,opt,name=internal,proto3,oneof" json:"internal,omitempty"`
}

func (x *IngestionFinalLocation) Reset() {
	*x = IngestionFinalLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_archive_dataset_v2_physical_location_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngestionFinalLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngestionFinalLocation) ProtoMessage() {}

func (x *IngestionFinalLocation) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_dataset_v2_physical_location_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngestionFinalLocation.ProtoReflect.Descriptor instead.
func (*IngestionFinalLocation) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_physical_location_proto_rawDescGZIP(), []int{6}
}

func (x *IngestionFinalLocation) GetLocation() *FinalLocation {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *IngestionFinalLocation) GetInternal() *IngestionFinalLocationProps {
	if x != nil {
		return x.Internal
	}
	return nil
}

var File_com_coralogix_archive_dataset_v2_physical_location_proto protoreflect.FileDescriptor

var file_com_coralogix_archive_dataset_v2_physical_location_proto_rawDesc = []byte{
	0x0a, 0x38, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2f,
	0x76, 0x32, 0x2f, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x1a, 0x25, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x61, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x03, 0x0a, 0x13, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x02,
	0x73, 0x33, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x2e, 0x76, 0x32, 0x2e, 0x53, 0x33, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63,
	0x48, 0x00, 0x52, 0x02, 0x73, 0x33, 0x12, 0x45, 0x0a, 0x07, 0x69, 0x62, 0x6d, 0x5f, 0x63, 0x6f,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e,
	0x76, 0x32, 0x2e, 0x49, 0x42, 0x4d, 0x43, 0x6f, 0x73, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53,
	0x70, 0x65, 0x63, 0x48, 0x00, 0x52, 0x06, 0x69, 0x62, 0x6d, 0x43, 0x6f, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x54, 0x61, 0x67, 0x73, 0x12, 0x37, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12,
	0x62, 0x0a, 0x13, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x5f,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x48, 0x01, 0x52, 0x12, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x16, 0x0a,
	0x14, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x65, 0x22, 0x80, 0x02, 0x0a, 0x0c, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x6b, 0x0a, 0x15,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x00, 0x52, 0x13, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x06, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x06, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xa8, 0x02, 0x0a, 0x1a, 0x49, 0x6e, 0x67,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70,
	0x65, 0x63, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x6d, 0x61, 0x78, 0x5f, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x12, 0x31, 0x0a, 0x15, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x12, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x42, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x12, 0x66, 0x0a, 0x13, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x69, 0x6e, 0x67, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x61, 0x72, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x42, 0x02,
	0x18, 0x01, 0x48, 0x00, 0x52, 0x12, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x69,
	0x6e, 0x67, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x16, 0x0a, 0x14, 0x5f,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x65, 0x22, 0xb4, 0x01, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x40, 0x0a, 0x1a, 0x66, 0x69, 0x6e,
	0x61, 0x6c, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x17, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61,
	0x73, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x88, 0x01, 0x01, 0x42, 0x1d, 0x0a, 0x1b, 0x5f,
	0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62,
	0x61, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0xcb, 0x02, 0x0a, 0x1b, 0x49,
	0x6e, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x12, 0x50, 0x0a, 0x04, 0x73, 0x70,
	0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6e, 0x67, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65,
	0x63, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x2c, 0x0a, 0x0f,
	0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0e, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x88, 0x01, 0x01, 0x12, 0x37, 0x0a, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x12, 0x40, 0x0a, 0x1a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x17, 0x66, 0x69, 0x6e, 0x61, 0x6c,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x75, 0x6c, 0x6c, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x88, 0x01, 0x01, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x42, 0x1d, 0x0a, 0x1b, 0x5f, 0x66, 0x69,
	0x6e, 0x61, 0x6c, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x75, 0x6c,
	0x6c, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0xcf, 0x01, 0x0a, 0x15, 0x49, 0x6e, 0x67,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70,
	0x65, 0x63, 0x12, 0x4a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x70, 0x65, 0x63, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5d,
	0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x48, 0x00,
	0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a,
	0x09, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x22, 0xd2, 0x01, 0x0a, 0x16, 0x49,
	0x6e, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x69, 0x6e, 0x61, 0x6c,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x5e, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x70, 0x73, 0x48, 0x00, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x88,
	0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_archive_dataset_v2_physical_location_proto_rawDescOnce sync.Once
	file_com_coralogix_archive_dataset_v2_physical_location_proto_rawDescData = file_com_coralogix_archive_dataset_v2_physical_location_proto_rawDesc
)

func file_com_coralogix_archive_dataset_v2_physical_location_proto_rawDescGZIP() []byte {
	file_com_coralogix_archive_dataset_v2_physical_location_proto_rawDescOnce.Do(func() {
		file_com_coralogix_archive_dataset_v2_physical_location_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_archive_dataset_v2_physical_location_proto_rawDescData)
	})
	return file_com_coralogix_archive_dataset_v2_physical_location_proto_rawDescData
}

var file_com_coralogix_archive_dataset_v2_physical_location_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_com_coralogix_archive_dataset_v2_physical_location_proto_goTypes = []any{
	(*ObjectStoreLocation)(nil),         // 0: com.coralogix.archive.dataset.v2.ObjectStoreLocation
	(*LocationSpec)(nil),                // 1: com.coralogix.archive.dataset.v2.LocationSpec
	(*IngestionLocationSpecProps)(nil),  // 2: com.coralogix.archive.dataset.v2.IngestionLocationSpecProps
	(*FinalLocation)(nil),               // 3: com.coralogix.archive.dataset.v2.FinalLocation
	(*IngestionFinalLocationProps)(nil), // 4: com.coralogix.archive.dataset.v2.IngestionFinalLocationProps
	(*IngestionLocationSpec)(nil),       // 5: com.coralogix.archive.dataset.v2.IngestionLocationSpec
	(*IngestionFinalLocation)(nil),      // 6: com.coralogix.archive.dataset.v2.IngestionFinalLocation
	(*S3TargetSpec)(nil),                // 7: com.coralogix.archive.v2.S3TargetSpec
	(*IBMCosTargetSpec)(nil),            // 8: com.coralogix.archive.v2.IBMCosTargetSpec
	(*ObjectTag)(nil),                   // 9: com.coralogix.archive.v2.ObjectTag
	(PartitioningScheme)(0),             // 10: com.coralogix.archive.v2.PartitioningScheme
	(*Format)(nil),                      // 11: com.coralogix.archive.v2.Format
}
var file_com_coralogix_archive_dataset_v2_physical_location_proto_depIdxs = []int32{
	7,  // 0: com.coralogix.archive.dataset.v2.ObjectStoreLocation.s3:type_name -> com.coralogix.archive.v2.S3TargetSpec
	8,  // 1: com.coralogix.archive.dataset.v2.ObjectStoreLocation.ibm_cos:type_name -> com.coralogix.archive.v2.IBMCosTargetSpec
	9,  // 2: com.coralogix.archive.dataset.v2.ObjectStoreLocation.tags:type_name -> com.coralogix.archive.v2.ObjectTag
	10, // 3: com.coralogix.archive.dataset.v2.ObjectStoreLocation.partitioning_scheme:type_name -> com.coralogix.archive.v2.PartitioningScheme
	0,  // 4: com.coralogix.archive.dataset.v2.LocationSpec.object_store_location:type_name -> com.coralogix.archive.dataset.v2.ObjectStoreLocation
	11, // 5: com.coralogix.archive.dataset.v2.LocationSpec.format:type_name -> com.coralogix.archive.v2.Format
	10, // 6: com.coralogix.archive.dataset.v2.IngestionLocationSpecProps.partitioning_scheme:type_name -> com.coralogix.archive.v2.PartitioningScheme
	1,  // 7: com.coralogix.archive.dataset.v2.FinalLocation.spec:type_name -> com.coralogix.archive.dataset.v2.LocationSpec
	2,  // 8: com.coralogix.archive.dataset.v2.IngestionFinalLocationProps.spec:type_name -> com.coralogix.archive.dataset.v2.IngestionLocationSpecProps
	9,  // 9: com.coralogix.archive.dataset.v2.IngestionFinalLocationProps.tags:type_name -> com.coralogix.archive.v2.ObjectTag
	1,  // 10: com.coralogix.archive.dataset.v2.IngestionLocationSpec.location:type_name -> com.coralogix.archive.dataset.v2.LocationSpec
	2,  // 11: com.coralogix.archive.dataset.v2.IngestionLocationSpec.internal:type_name -> com.coralogix.archive.dataset.v2.IngestionLocationSpecProps
	3,  // 12: com.coralogix.archive.dataset.v2.IngestionFinalLocation.location:type_name -> com.coralogix.archive.dataset.v2.FinalLocation
	4,  // 13: com.coralogix.archive.dataset.v2.IngestionFinalLocation.internal:type_name -> com.coralogix.archive.dataset.v2.IngestionFinalLocationProps
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_com_coralogix_archive_dataset_v2_physical_location_proto_init() }
func file_com_coralogix_archive_dataset_v2_physical_location_proto_init() {
	if File_com_coralogix_archive_dataset_v2_physical_location_proto != nil {
		return
	}
	file_com_coralogix_archive_v2_target_proto_init()
	file_com_coralogix_archive_v2_object_tag_proto_init()
	file_com_coralogix_archive_v2_archive_target_proto_init()
	file_com_coralogix_archive_v2_format_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_archive_dataset_v2_physical_location_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ObjectStoreLocation); i {
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
		file_com_coralogix_archive_dataset_v2_physical_location_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*LocationSpec); i {
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
		file_com_coralogix_archive_dataset_v2_physical_location_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*IngestionLocationSpecProps); i {
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
		file_com_coralogix_archive_dataset_v2_physical_location_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*FinalLocation); i {
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
		file_com_coralogix_archive_dataset_v2_physical_location_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*IngestionFinalLocationProps); i {
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
		file_com_coralogix_archive_dataset_v2_physical_location_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*IngestionLocationSpec); i {
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
		file_com_coralogix_archive_dataset_v2_physical_location_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*IngestionFinalLocation); i {
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
	file_com_coralogix_archive_dataset_v2_physical_location_proto_msgTypes[0].OneofWrappers = []any{
		(*ObjectStoreLocation_S3)(nil),
		(*ObjectStoreLocation_IbmCos)(nil),
	}
	file_com_coralogix_archive_dataset_v2_physical_location_proto_msgTypes[1].OneofWrappers = []any{
		(*LocationSpec_ObjectStoreLocation)(nil),
	}
	file_com_coralogix_archive_dataset_v2_physical_location_proto_msgTypes[2].OneofWrappers = []any{}
	file_com_coralogix_archive_dataset_v2_physical_location_proto_msgTypes[3].OneofWrappers = []any{}
	file_com_coralogix_archive_dataset_v2_physical_location_proto_msgTypes[4].OneofWrappers = []any{}
	file_com_coralogix_archive_dataset_v2_physical_location_proto_msgTypes[5].OneofWrappers = []any{}
	file_com_coralogix_archive_dataset_v2_physical_location_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_archive_dataset_v2_physical_location_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_archive_dataset_v2_physical_location_proto_goTypes,
		DependencyIndexes: file_com_coralogix_archive_dataset_v2_physical_location_proto_depIdxs,
		MessageInfos:      file_com_coralogix_archive_dataset_v2_physical_location_proto_msgTypes,
	}.Build()
	File_com_coralogix_archive_dataset_v2_physical_location_proto = out.File
	file_com_coralogix_archive_dataset_v2_physical_location_proto_rawDesc = nil
	file_com_coralogix_archive_dataset_v2_physical_location_proto_goTypes = nil
	file_com_coralogix_archive_dataset_v2_physical_location_proto_depIdxs = nil
}
