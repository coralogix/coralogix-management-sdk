// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.26.1
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
	state                protoimpl.MessageState  `protogen:"open.v1"`
	ObjectKey            string                  `protobuf:"bytes,2,opt,name=object_key,json=objectKey,proto3" json:"object_key,omitempty"`
	ObjectSize           uint64                  `protobuf:"varint,3,opt,name=object_size,json=objectSize,proto3" json:"object_size,omitempty"`
	CompactionLevel      uint32                  `protobuf:"varint,6,opt,name=compaction_level,json=compactionLevel,proto3" json:"compaction_level,omitempty"`
	Dataset              *v1.Dataset             `protobuf:"bytes,7,opt,name=dataset,proto3" json:"dataset,omitempty"`
	RowCount             uint64                  `protobuf:"varint,8,opt,name=row_count,json=rowCount,proto3" json:"row_count,omitempty"`
	CollectedEventLabels []*Labels               `protobuf:"bytes,9,rep,name=collected_event_labels,json=collectedEventLabels,proto3" json:"collected_event_labels,omitempty"`
	SplitByEventLabels   *Labels                 `protobuf:"bytes,10,opt,name=split_by_event_labels,json=splitByEventLabels,proto3" json:"split_by_event_labels,omitempty"`
	ArchiveTarget        *ArchiveTarget          `protobuf:"bytes,11,opt,name=archive_target,json=archiveTarget,proto3" json:"archive_target,omitempty"`
	ColumnsLimitReached  bool                    `protobuf:"varint,14,opt,name=columns_limit_reached,json=columnsLimitReached,proto3" json:"columns_limit_reached,omitempty"`
	EventTimeRange       *EventTimeRange         `protobuf:"bytes,15,opt,name=event_time_range,json=eventTimeRange,proto3" json:"event_time_range,omitempty"`
	PhysicalLocation     *IngestionFinalLocation `protobuf:"bytes,16,opt,name=physical_location,json=physicalLocation,proto3" json:"physical_location,omitempty"`
	Schema               *SchemaTree             `protobuf:"bytes,17,opt,name=schema,proto3" json:"schema,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *Object) Reset() {
	*x = Object{}
	mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[0]
	if x != nil {
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

func (x *Object) GetColumnsLimitReached() bool {
	if x != nil {
		return x.ColumnsLimitReached
	}
	return false
}

func (x *Object) GetEventTimeRange() *EventTimeRange {
	if x != nil {
		return x.EventTimeRange
	}
	return nil
}

func (x *Object) GetPhysicalLocation() *IngestionFinalLocation {
	if x != nil {
		return x.PhysicalLocation
	}
	return nil
}

func (x *Object) GetSchema() *SchemaTree {
	if x != nil {
		return x.Schema
	}
	return nil
}

type ObjectCreated struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Object        *Object                `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ObjectCreated) Reset() {
	*x = ObjectCreated{}
	mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ObjectCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectCreated) ProtoMessage() {}

func (x *ObjectCreated) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[1]
	if x != nil {
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Object        *Object                `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ObjectRemoved) Reset() {
	*x = ObjectRemoved{}
	mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ObjectRemoved) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectRemoved) ProtoMessage() {}

func (x *ObjectRemoved) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[2]
	if x != nil {
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Created       []*ObjectCreated       `protobuf:"bytes,1,rep,name=created,proto3" json:"created,omitempty"`
	Removed       []*ObjectRemoved       `protobuf:"bytes,2,rep,name=removed,proto3" json:"removed,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ObjectsReplaced) Reset() {
	*x = ObjectsReplaced{}
	mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ObjectsReplaced) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectsReplaced) ProtoMessage() {}

func (x *ObjectsReplaced) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[3]
	if x != nil {
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

type ObjectsReplacedAll struct {
	state                      protoimpl.MessageState `protogen:"open.v1"`
	Created                    []*ObjectCreated       `protobuf:"bytes,1,rep,name=created,proto3" json:"created,omitempty"`
	RemovedPhysicalLocationIds []string               `protobuf:"bytes,2,rep,name=removed_physical_location_ids,json=removedPhysicalLocationIds,proto3" json:"removed_physical_location_ids,omitempty"`
	// Added for compatibility reasons. Will be removed once metastore moved to PL stable id everywhere
	Locations     []*FinalLocation `protobuf:"bytes,3,rep,name=locations,proto3" json:"locations,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ObjectsReplacedAll) Reset() {
	*x = ObjectsReplacedAll{}
	mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ObjectsReplacedAll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectsReplacedAll) ProtoMessage() {}

func (x *ObjectsReplacedAll) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectsReplacedAll.ProtoReflect.Descriptor instead.
func (*ObjectsReplacedAll) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_event_proto_rawDescGZIP(), []int{4}
}

func (x *ObjectsReplacedAll) GetCreated() []*ObjectCreated {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *ObjectsReplacedAll) GetRemovedPhysicalLocationIds() []string {
	if x != nil {
		return x.RemovedPhysicalLocationIds
	}
	return nil
}

func (x *ObjectsReplacedAll) GetLocations() []*FinalLocation {
	if x != nil {
		return x.Locations
	}
	return nil
}

type CompactionEvent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Event:
	//
	//	*CompactionEvent_Created
	//	*CompactionEvent_Replaced
	//	*CompactionEvent_ReplacedAll
	Event         isCompactionEvent_Event `protobuf_oneof:"event"`
	ArchiveId     string                  `protobuf:"bytes,3,opt,name=archive_id,json=archiveId,proto3" json:"archive_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CompactionEvent) Reset() {
	*x = CompactionEvent{}
	mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompactionEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompactionEvent) ProtoMessage() {}

func (x *CompactionEvent) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[5]
	if x != nil {
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
	return file_com_coralogix_archive_v2_event_proto_rawDescGZIP(), []int{5}
}

func (x *CompactionEvent) GetEvent() isCompactionEvent_Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *CompactionEvent) GetCreated() *ObjectCreated {
	if x != nil {
		if x, ok := x.Event.(*CompactionEvent_Created); ok {
			return x.Created
		}
	}
	return nil
}

func (x *CompactionEvent) GetReplaced() *ObjectsReplaced {
	if x != nil {
		if x, ok := x.Event.(*CompactionEvent_Replaced); ok {
			return x.Replaced
		}
	}
	return nil
}

func (x *CompactionEvent) GetReplacedAll() *ObjectsReplacedAll {
	if x != nil {
		if x, ok := x.Event.(*CompactionEvent_ReplacedAll); ok {
			return x.ReplacedAll
		}
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

type CompactionEvent_ReplacedAll struct {
	ReplacedAll *ObjectsReplacedAll `protobuf:"bytes,4,opt,name=replaced_all,json=replacedAll,proto3,oneof"`
}

func (*CompactionEvent_Created) isCompactionEvent_Event() {}

func (*CompactionEvent_Replaced) isCompactionEvent_Event() {}

func (*CompactionEvent_ReplacedAll) isCompactionEvent_Event() {}

type Labels struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Kvs           []*KeyVal              `protobuf:"bytes,1,rep,name=kvs,proto3" json:"kvs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Labels) Reset() {
	*x = Labels{}
	mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Labels) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Labels) ProtoMessage() {}

func (x *Labels) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[6]
	if x != nil {
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
	return file_com_coralogix_archive_v2_event_proto_rawDescGZIP(), []int{6}
}

func (x *Labels) GetKvs() []*KeyVal {
	if x != nil {
		return x.Kvs
	}
	return nil
}

type KeyVal struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KeyVal) Reset() {
	*x = KeyVal{}
	mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KeyVal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyVal) ProtoMessage() {}

func (x *KeyVal) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[7]
	if x != nil {
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
	return file_com_coralogix_archive_v2_event_proto_rawDescGZIP(), []int{7}
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

type EventTimeRange struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MaxTimestamp  int64                  `protobuf:"varint,1,opt,name=max_timestamp,json=maxTimestamp,proto3" json:"max_timestamp,omitempty"`
	MinTimestamp  int64                  `protobuf:"varint,2,opt,name=min_timestamp,json=minTimestamp,proto3" json:"min_timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EventTimeRange) Reset() {
	*x = EventTimeRange{}
	mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventTimeRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventTimeRange) ProtoMessage() {}

func (x *EventTimeRange) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_event_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventTimeRange.ProtoReflect.Descriptor instead.
func (*EventTimeRange) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_event_proto_rawDescGZIP(), []int{8}
}

func (x *EventTimeRange) GetMaxTimestamp() int64 {
	if x != nil {
		return x.MaxTimestamp
	}
	return 0
}

func (x *EventTimeRange) GetMinTimestamp() int64 {
	if x != nil {
		return x.MinTimestamp
	}
	return 0
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
	0x76, 0x65, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x38, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2f, 0x76,
	0x32, 0x2f, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x2f, 0x76, 0x32, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x9d, 0x06, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x43, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72,
	0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x72, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x56, 0x0a, 0x16, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x14, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x12, 0x53, 0x0a, 0x15, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x52, 0x12, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x42, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x4e, 0x0a, 0x0e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x0d, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x63, 0x68, 0x65, 0x64, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x52, 0x65, 0x61, 0x63, 0x68, 0x65, 0x64, 0x12, 0x52, 0x0a, 0x10, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x65, 0x0a,
	0x11, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6e, 0x67, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x10, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x54, 0x72, 0x65, 0x65, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x4a, 0x04,
	0x08, 0x05, 0x10, 0x06, 0x4a, 0x04, 0x08, 0x0c, 0x10, 0x0d, 0x4a, 0x04, 0x08, 0x0d, 0x10, 0x0e,
	0x22, 0x49, 0x0a, 0x0d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x38, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x49, 0x0a, 0x0d, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x06,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x97, 0x01, 0x0a, 0x0f, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x64, 0x12, 0x41, 0x0a, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x41, 0x0a,
	0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64,
	0x22, 0xe9, 0x01, 0x0a, 0x12, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x41, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x41, 0x0a, 0x1d, 0x72, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x64, 0x5f, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x1a, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63,
	0x61, 0x6c, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x12, 0x4d, 0x0a,
	0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x2e, 0x76, 0x32, 0x2e, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x9a, 0x02, 0x0a,
	0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x43, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x47, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e,
	0x76, 0x32, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x64, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x64, 0x12, 0x51,
	0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x64, 0x5f, 0x61, 0x6c, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x64, 0x41,
	0x6c, 0x6c, 0x48, 0x00, 0x52, 0x0b, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x64, 0x41, 0x6c,
	0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x49, 0x64,
	0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x3c, 0x0a, 0x06, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x12, 0x32, 0x0a, 0x03, 0x6b, 0x76, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x4b, 0x65, 0x79, 0x56,
	0x61, 0x6c, 0x52, 0x03, 0x6b, 0x76, 0x73, 0x22, 0x30, 0x0a, 0x06, 0x4b, 0x65, 0x79, 0x56, 0x61,
	0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x5a, 0x0a, 0x0e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d,
	0x61, 0x78, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_com_coralogix_archive_v2_event_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_com_coralogix_archive_v2_event_proto_goTypes = []any{
	(*Object)(nil),                 // 0: com.coralogix.archive.v2.Object
	(*ObjectCreated)(nil),          // 1: com.coralogix.archive.v2.ObjectCreated
	(*ObjectRemoved)(nil),          // 2: com.coralogix.archive.v2.ObjectRemoved
	(*ObjectsReplaced)(nil),        // 3: com.coralogix.archive.v2.ObjectsReplaced
	(*ObjectsReplacedAll)(nil),     // 4: com.coralogix.archive.v2.ObjectsReplacedAll
	(*CompactionEvent)(nil),        // 5: com.coralogix.archive.v2.CompactionEvent
	(*Labels)(nil),                 // 6: com.coralogix.archive.v2.Labels
	(*KeyVal)(nil),                 // 7: com.coralogix.archive.v2.KeyVal
	(*EventTimeRange)(nil),         // 8: com.coralogix.archive.v2.EventTimeRange
	(*v1.Dataset)(nil),             // 9: com.coralogix.archive.dataset.v1.Dataset
	(*ArchiveTarget)(nil),          // 10: com.coralogix.archive.v2.ArchiveTarget
	(*IngestionFinalLocation)(nil), // 11: com.coralogix.archive.dataset.v2.IngestionFinalLocation
	(*SchemaTree)(nil),             // 12: com.coralogix.archive.v2.SchemaTree
	(*FinalLocation)(nil),          // 13: com.coralogix.archive.dataset.v2.FinalLocation
}
var file_com_coralogix_archive_v2_event_proto_depIdxs = []int32{
	9,  // 0: com.coralogix.archive.v2.Object.dataset:type_name -> com.coralogix.archive.dataset.v1.Dataset
	6,  // 1: com.coralogix.archive.v2.Object.collected_event_labels:type_name -> com.coralogix.archive.v2.Labels
	6,  // 2: com.coralogix.archive.v2.Object.split_by_event_labels:type_name -> com.coralogix.archive.v2.Labels
	10, // 3: com.coralogix.archive.v2.Object.archive_target:type_name -> com.coralogix.archive.v2.ArchiveTarget
	8,  // 4: com.coralogix.archive.v2.Object.event_time_range:type_name -> com.coralogix.archive.v2.EventTimeRange
	11, // 5: com.coralogix.archive.v2.Object.physical_location:type_name -> com.coralogix.archive.dataset.v2.IngestionFinalLocation
	12, // 6: com.coralogix.archive.v2.Object.schema:type_name -> com.coralogix.archive.v2.SchemaTree
	0,  // 7: com.coralogix.archive.v2.ObjectCreated.object:type_name -> com.coralogix.archive.v2.Object
	0,  // 8: com.coralogix.archive.v2.ObjectRemoved.object:type_name -> com.coralogix.archive.v2.Object
	1,  // 9: com.coralogix.archive.v2.ObjectsReplaced.created:type_name -> com.coralogix.archive.v2.ObjectCreated
	2,  // 10: com.coralogix.archive.v2.ObjectsReplaced.removed:type_name -> com.coralogix.archive.v2.ObjectRemoved
	1,  // 11: com.coralogix.archive.v2.ObjectsReplacedAll.created:type_name -> com.coralogix.archive.v2.ObjectCreated
	13, // 12: com.coralogix.archive.v2.ObjectsReplacedAll.locations:type_name -> com.coralogix.archive.dataset.v2.FinalLocation
	1,  // 13: com.coralogix.archive.v2.CompactionEvent.created:type_name -> com.coralogix.archive.v2.ObjectCreated
	3,  // 14: com.coralogix.archive.v2.CompactionEvent.replaced:type_name -> com.coralogix.archive.v2.ObjectsReplaced
	4,  // 15: com.coralogix.archive.v2.CompactionEvent.replaced_all:type_name -> com.coralogix.archive.v2.ObjectsReplacedAll
	7,  // 16: com.coralogix.archive.v2.Labels.kvs:type_name -> com.coralogix.archive.v2.KeyVal
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_com_coralogix_archive_v2_event_proto_init() }
func file_com_coralogix_archive_v2_event_proto_init() {
	if File_com_coralogix_archive_v2_event_proto != nil {
		return
	}
	file_com_coralogix_archive_v2_archive_target_proto_init()
	file_com_coralogix_archive_dataset_v2_physical_location_proto_init()
	file_com_coralogix_archive_v2_schema_proto_init()
	file_com_coralogix_archive_v2_event_proto_msgTypes[5].OneofWrappers = []any{
		(*CompactionEvent_Created)(nil),
		(*CompactionEvent_Replaced)(nil),
		(*CompactionEvent_ReplacedAll)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_archive_v2_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
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
