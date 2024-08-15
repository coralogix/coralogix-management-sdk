// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: com/coralogix/archive/v2/event.proto

package v2

import (
	v1 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/archive/dataset/v1"
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

type Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectKey            string         `protobuf:"bytes,2,opt,name=object_key,json=objectKey,proto3" json:"object_key,omitempty"`
	ObjectSize           uint64         `protobuf:"varint,3,opt,name=object_size,json=objectSize,proto3" json:"object_size,omitempty"`
	CompactionLevel      uint32         `protobuf:"varint,6,opt,name=compaction_level,json=compactionLevel,proto3" json:"compaction_level,omitempty"`
	Dataset              *v1.Dataset    `protobuf:"bytes,7,opt,name=dataset,proto3" json:"dataset,omitempty"`
	RowCount             uint64         `protobuf:"varint,8,opt,name=row_count,json=rowCount,proto3" json:"row_count,omitempty"`
	CollectedEventLabels []*Labels      `protobuf:"bytes,9,rep,name=collected_event_labels,json=collectedEventLabels,proto3" json:"collected_event_labels,omitempty"`
	SplitByEventLabels   *Labels        `protobuf:"bytes,10,opt,name=split_by_event_labels,json=splitByEventLabels,proto3" json:"split_by_event_labels,omitempty"`
	ArchiveTarget        *ArchiveTarget `protobuf:"bytes,11,opt,name=archive_target,json=archiveTarget,proto3" json:"archive_target,omitempty"`
}

func (x *Object) Reset() {
	*x = Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_event_proto_rawDescGZIP(), []int{0}
}

func (x *Object) GetObjectKey() string {
	if x != nil {
		return x.ObjectKey
	}
	return ""
}

func (x *Object) GetObjectSize() uint64 {
	if x != nil {
		return x.ObjectSize
	}
	return 0
}

func (x *Object) GetCompactionLevel() uint32 {
	if x != nil {
		return x.CompactionLevel
	}
	return 0
}

func (x *Object) GetDataset() *v1.Dataset {
	if x != nil {
		return x.Dataset
	}
	return nil
}

func (x *Object) GetRowCount() uint64 {
	if x != nil {
		return x.RowCount
	}
	return 0
}

func (x *Object) GetCollectedEventLabels() []*Labels {
	if x != nil {
		return x.CollectedEventLabels
	}
	return nil
}

func (x *Object) GetSplitByEventLabels() *Labels {
	if x != nil {
		return x.SplitByEventLabels
	}
	return nil
}

func (x *Object) GetArchiveTarget() *ArchiveTarget {
	if x != nil {
		return x.ArchiveTarget
	}
	return nil
}

type ObjectCreated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object *Object `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *ObjectCreated) Reset() {
	*x = ObjectCreated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectCreated) ProtoMessage() {}

func (x *ObjectCreated) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectCreated.ProtoReflect.Descriptor instead.
func (*ObjectCreated) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_event_proto_rawDescGZIP(), []int{1}
}

func (x *ObjectCreated) GetObject() *Object {
	if x != nil {
		return x.Object
	}
	return nil
}

type ObjectRemoved struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object *Object `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *ObjectRemoved) Reset() {
	*x = ObjectRemoved{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectRemoved) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectRemoved) ProtoMessage() {}

func (x *ObjectRemoved) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectRemoved.ProtoReflect.Descriptor instead.
func (*ObjectRemoved) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_event_proto_rawDescGZIP(), []int{2}
}

func (x *ObjectRemoved) GetObject() *Object {
	if x != nil {
		return x.Object
	}
	return nil
}

type ObjectsReplaced struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Created []*ObjectCreated `protobuf:"bytes,1,rep,name=created,proto3" json:"created,omitempty"`
	Removed []*ObjectRemoved `protobuf:"bytes,2,rep,name=removed,proto3" json:"removed,omitempty"`
}

func (x *ObjectsReplaced) Reset() {
	*x = ObjectsReplaced{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectsReplaced) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectsReplaced) ProtoMessage() {}

func (x *ObjectsReplaced) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectsReplaced.ProtoReflect.Descriptor instead.
func (*ObjectsReplaced) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_event_proto_rawDescGZIP(), []int{3}
}

func (x *ObjectsReplaced) GetCreated() []*ObjectCreated {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *ObjectsReplaced) GetRemoved() []*ObjectRemoved {
	if x != nil {
		return x.Removed
	}
	return nil
}

type CompactionEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//
	//	*CompactionEvent_Created
	//	*CompactionEvent_Replaced
	Event     isCompactionEvent_Event `protobuf_oneof:"event"`
	ArchiveId string                  `protobuf:"bytes,3,opt,name=archive_id,json=archiveId,proto3" json:"archive_id,omitempty"`
}

func (x *CompactionEvent) Reset() {
	*x = CompactionEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompactionEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompactionEvent) ProtoMessage() {}

func (x *CompactionEvent) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompactionEvent.ProtoReflect.Descriptor instead.
func (*CompactionEvent) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_event_proto_rawDescGZIP(), []int{4}
}

func (m *CompactionEvent) GetEvent() isCompactionEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *CompactionEvent) GetCreated() *ObjectCreated {
	if x, ok := x.GetEvent().(*CompactionEvent_Created); ok {
		return x.Created
	}
	return nil
}

func (x *CompactionEvent) GetReplaced() *ObjectsReplaced {
	if x, ok := x.GetEvent().(*CompactionEvent_Replaced); ok {
		return x.Replaced
	}
	return nil
}

func (x *CompactionEvent) GetArchiveId() string {
	if x != nil {
		return x.ArchiveId
	}
	return ""
}

type isCompactionEvent_Event interface {
	isCompactionEvent_Event()
}

type CompactionEvent_Created struct {
	Created *ObjectCreated `protobuf:"bytes,1,opt,name=created,proto3,oneof"`
}

type CompactionEvent_Replaced struct {
	Replaced *ObjectsReplaced `protobuf:"bytes,2,opt,name=replaced,proto3,oneof"`
}

func (*CompactionEvent_Created) isCompactionEvent_Event() {}

func (*CompactionEvent_Replaced) isCompactionEvent_Event() {}

type Labels struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kvs []*KeyVal `protobuf:"bytes,1,rep,name=kvs,proto3" json:"kvs,omitempty"`
}

func (x *Labels) Reset() {
	*x = Labels{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Labels) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Labels) ProtoMessage() {}

func (x *Labels) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Labels.ProtoReflect.Descriptor instead.
func (*Labels) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_event_proto_rawDescGZIP(), []int{5}
}

func (x *Labels) GetKvs() []*KeyVal {
	if x != nil {
		return x.Kvs
	}
	return nil
}

type KeyVal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KeyVal) Reset() {
	*x = KeyVal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyVal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyVal) ProtoMessage() {}

func (x *KeyVal) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyVal.ProtoReflect.Descriptor instead.
func (*KeyVal) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_event_proto_rawDescGZIP(), []int{6}
}

func (x *KeyVal) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyVal) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_com_coralogix_archive_v2_event_proto protoreflect.FileDescriptor

var file_com_coralogix_archive_v2_event_proto_rawDesc = []byte{
	0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32,
	0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf0, 0x03, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x43, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f,
	0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x72,
	0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x56, 0x0a, 0x16, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e,
	0x76, 0x32, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x14, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12,
	0x53, 0x0a, 0x15, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x52, 0x12, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x42, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x4e, 0x0a, 0x0e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x5f,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x0d, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05,
	0x4a, 0x04, 0x08, 0x05, 0x10, 0x06, 0x4a, 0x04, 0x08, 0x0c, 0x10, 0x0d, 0x4a, 0x04, 0x08, 0x0d,
	0x10, 0x0e, 0x22, 0x49, 0x0a, 0x0d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x49, 0x0a,
	0x0d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x38,
	0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x97, 0x01, 0x0a, 0x0f, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x64, 0x12, 0x41, 0x0a, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x41, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x64, 0x22, 0xc7, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x43, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e,
	0x76, 0x32, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x48, 0x00, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x47, 0x0a, 0x08, 0x72,
	0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x64, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x49, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x3c, 0x0a, 0x06,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x32, 0x0a, 0x03, 0x6b, 0x76, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x4b,
	0x65, 0x79, 0x56, 0x61, 0x6c, 0x52, 0x03, 0x6b, 0x76, 0x73, 0x22, 0x30, 0x0a, 0x06, 0x4b, 0x65,
	0x79, 0x56, 0x61, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_archive_v2_event_proto_rawDescOnce sync.Once
	file_com_coralogix_archive_v2_event_proto_rawDescData = file_com_coralogix_archive_v2_event_proto_rawDesc
)

func file_com_coralogix_archive_v2_event_proto_rawDescGZIP() []byte {
	file_com_coralogix_archive_v2_event_proto_rawDescOnce.Do(func() {
		file_com_coralogix_archive_v2_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_archive_v2_event_proto_rawDescData)
	})
	return file_com_coralogix_archive_v2_event_proto_rawDescData
}

var file_com_coralogix_archive_v2_event_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_com_coralogix_archive_v2_event_proto_goTypes = []any{
	(*Object)(nil),          // 0: com.coralogix.archive.v2.Object
	(*ObjectCreated)(nil),   // 1: com.coralogix.archive.v2.ObjectCreated
	(*ObjectRemoved)(nil),   // 2: com.coralogix.archive.v2.ObjectRemoved
	(*ObjectsReplaced)(nil), // 3: com.coralogix.archive.v2.ObjectsReplaced
	(*CompactionEvent)(nil), // 4: com.coralogix.archive.v2.CompactionEvent
	(*Labels)(nil),          // 5: com.coralogix.archive.v2.Labels
	(*KeyVal)(nil),          // 6: com.coralogix.archive.v2.KeyVal
	(*v1.Dataset)(nil),      // 7: com.coralogix.archive.dataset.v1.Dataset
	(*ArchiveTarget)(nil),   // 8: com.coralogix.archive.v2.ArchiveTarget
}
var file_com_coralogix_archive_v2_event_proto_depIdxs = []int32{
	7,  // 0: com.coralogix.archive.v2.Object.dataset:type_name -> com.coralogix.archive.dataset.v1.Dataset
	5,  // 1: com.coralogix.archive.v2.Object.collected_event_labels:type_name -> com.coralogix.archive.v2.Labels
	5,  // 2: com.coralogix.archive.v2.Object.split_by_event_labels:type_name -> com.coralogix.archive.v2.Labels
	8,  // 3: com.coralogix.archive.v2.Object.archive_target:type_name -> com.coralogix.archive.v2.ArchiveTarget
	0,  // 4: com.coralogix.archive.v2.ObjectCreated.object:type_name -> com.coralogix.archive.v2.Object
	0,  // 5: com.coralogix.archive.v2.ObjectRemoved.object:type_name -> com.coralogix.archive.v2.Object
	1,  // 6: com.coralogix.archive.v2.ObjectsReplaced.created:type_name -> com.coralogix.archive.v2.ObjectCreated
	2,  // 7: com.coralogix.archive.v2.ObjectsReplaced.removed:type_name -> com.coralogix.archive.v2.ObjectRemoved
	1,  // 8: com.coralogix.archive.v2.CompactionEvent.created:type_name -> com.coralogix.archive.v2.ObjectCreated
	3,  // 9: com.coralogix.archive.v2.CompactionEvent.replaced:type_name -> com.coralogix.archive.v2.ObjectsReplaced
	6,  // 10: com.coralogix.archive.v2.Labels.kvs:type_name -> com.coralogix.archive.v2.KeyVal
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_com_coralogix_archive_v2_event_proto_init() }
func file_com_coralogix_archive_v2_event_proto_init() {
	if File_com_coralogix_archive_v2_event_proto != nil {
		return
	}
	file_com_coralogix_archive_v2_archive_target_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_archive_v2_event_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Object); i {
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
		file_com_coralogix_archive_v2_event_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ObjectCreated); i {
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
		file_com_coralogix_archive_v2_event_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ObjectRemoved); i {
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
		file_com_coralogix_archive_v2_event_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ObjectsReplaced); i {
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
		file_com_coralogix_archive_v2_event_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CompactionEvent); i {
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
		file_com_coralogix_archive_v2_event_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Labels); i {
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
		file_com_coralogix_archive_v2_event_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*KeyVal); i {
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
	file_com_coralogix_archive_v2_event_proto_msgTypes[4].OneofWrappers = []any{
		(*CompactionEvent_Created)(nil),
		(*CompactionEvent_Replaced)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_archive_v2_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_archive_v2_event_proto_goTypes,
		DependencyIndexes: file_com_coralogix_archive_v2_event_proto_depIdxs,
		MessageInfos:      file_com_coralogix_archive_v2_event_proto_msgTypes,
	}.Build()
	File_com_coralogix_archive_v2_event_proto = out.File
	file_com_coralogix_archive_v2_event_proto_rawDesc = nil
	file_com_coralogix_archive_v2_event_proto_goTypes = nil
	file_com_coralogix_archive_v2_event_proto_depIdxs = nil
}
