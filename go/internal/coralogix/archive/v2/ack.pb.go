// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: com/coralogix/archive/v2/ack.proto

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

type MetastoreAck struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Created       []*ObjectCreatedStatus `protobuf:"bytes,1,rep,name=created,proto3" json:"created,omitempty"`
	Removed       []*ObjectRemovedStatus `protobuf:"bytes,2,rep,name=removed,proto3" json:"removed,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MetastoreAck) Reset() {
	*x = MetastoreAck{}
	mi := &file_com_coralogix_archive_v2_ack_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetastoreAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetastoreAck) ProtoMessage() {}

func (x *MetastoreAck) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_ack_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetastoreAck.ProtoReflect.Descriptor instead.
func (*MetastoreAck) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_ack_proto_rawDescGZIP(), []int{0}
}

func (x *MetastoreAck) GetCreated() []*ObjectCreatedStatus {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *MetastoreAck) GetRemoved() []*ObjectRemovedStatus {
	if x != nil {
		return x.Removed
	}
	return nil
}

type ObjectCreatedStatus struct {
	state   protoimpl.MessageState `protogen:"open.v1"`
	Command *ObjectCreated         `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	// Types that are valid to be assigned to Status:
	//
	//	*ObjectCreatedStatus_Success
	//	*ObjectCreatedStatus_AlreadyExists
	//	*ObjectCreatedStatus_Skipped
	Status        isObjectCreatedStatus_Status `protobuf_oneof:"status"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ObjectCreatedStatus) Reset() {
	*x = ObjectCreatedStatus{}
	mi := &file_com_coralogix_archive_v2_ack_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ObjectCreatedStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectCreatedStatus) ProtoMessage() {}

func (x *ObjectCreatedStatus) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_ack_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectCreatedStatus.ProtoReflect.Descriptor instead.
func (*ObjectCreatedStatus) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_ack_proto_rawDescGZIP(), []int{1}
}

func (x *ObjectCreatedStatus) GetCommand() *ObjectCreated {
	if x != nil {
		return x.Command
	}
	return nil
}

func (x *ObjectCreatedStatus) GetStatus() isObjectCreatedStatus_Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ObjectCreatedStatus) GetSuccess() *Success {
	if x != nil {
		if x, ok := x.Status.(*ObjectCreatedStatus_Success); ok {
			return x.Success
		}
	}
	return nil
}

func (x *ObjectCreatedStatus) GetAlreadyExists() *AlreadyExists {
	if x != nil {
		if x, ok := x.Status.(*ObjectCreatedStatus_AlreadyExists); ok {
			return x.AlreadyExists
		}
	}
	return nil
}

func (x *ObjectCreatedStatus) GetSkipped() *Skipped {
	if x != nil {
		if x, ok := x.Status.(*ObjectCreatedStatus_Skipped); ok {
			return x.Skipped
		}
	}
	return nil
}

type isObjectCreatedStatus_Status interface {
	isObjectCreatedStatus_Status()
}

type ObjectCreatedStatus_Success struct {
	Success *Success `protobuf:"bytes,2,opt,name=success,proto3,oneof"`
}

type ObjectCreatedStatus_AlreadyExists struct {
	AlreadyExists *AlreadyExists `protobuf:"bytes,3,opt,name=already_exists,json=alreadyExists,proto3,oneof"`
}

type ObjectCreatedStatus_Skipped struct {
	Skipped *Skipped `protobuf:"bytes,4,opt,name=skipped,proto3,oneof"`
}

func (*ObjectCreatedStatus_Success) isObjectCreatedStatus_Status() {}

func (*ObjectCreatedStatus_AlreadyExists) isObjectCreatedStatus_Status() {}

func (*ObjectCreatedStatus_Skipped) isObjectCreatedStatus_Status() {}

type ObjectRemovedStatus struct {
	state   protoimpl.MessageState `protogen:"open.v1"`
	Command *ObjectRemoved         `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	// Types that are valid to be assigned to Status:
	//
	//	*ObjectRemovedStatus_Success
	//	*ObjectRemovedStatus_AlreadyRemoved
	Status        isObjectRemovedStatus_Status `protobuf_oneof:"status"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ObjectRemovedStatus) Reset() {
	*x = ObjectRemovedStatus{}
	mi := &file_com_coralogix_archive_v2_ack_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ObjectRemovedStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectRemovedStatus) ProtoMessage() {}

func (x *ObjectRemovedStatus) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_ack_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectRemovedStatus.ProtoReflect.Descriptor instead.
func (*ObjectRemovedStatus) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_ack_proto_rawDescGZIP(), []int{2}
}

func (x *ObjectRemovedStatus) GetCommand() *ObjectRemoved {
	if x != nil {
		return x.Command
	}
	return nil
}

func (x *ObjectRemovedStatus) GetStatus() isObjectRemovedStatus_Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ObjectRemovedStatus) GetSuccess() *Success {
	if x != nil {
		if x, ok := x.Status.(*ObjectRemovedStatus_Success); ok {
			return x.Success
		}
	}
	return nil
}

func (x *ObjectRemovedStatus) GetAlreadyRemoved() *AlreadyRemoved {
	if x != nil {
		if x, ok := x.Status.(*ObjectRemovedStatus_AlreadyRemoved); ok {
			return x.AlreadyRemoved
		}
	}
	return nil
}

type isObjectRemovedStatus_Status interface {
	isObjectRemovedStatus_Status()
}

type ObjectRemovedStatus_Success struct {
	Success *Success `protobuf:"bytes,2,opt,name=success,proto3,oneof"`
}

type ObjectRemovedStatus_AlreadyRemoved struct {
	AlreadyRemoved *AlreadyRemoved `protobuf:"bytes,3,opt,name=already_removed,json=alreadyRemoved,proto3,oneof"`
}

func (*ObjectRemovedStatus_Success) isObjectRemovedStatus_Status() {}

func (*ObjectRemovedStatus_AlreadyRemoved) isObjectRemovedStatus_Status() {}

type AlreadyExists struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AlreadyExists) Reset() {
	*x = AlreadyExists{}
	mi := &file_com_coralogix_archive_v2_ack_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlreadyExists) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlreadyExists) ProtoMessage() {}

func (x *AlreadyExists) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_ack_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlreadyExists.ProtoReflect.Descriptor instead.
func (*AlreadyExists) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_ack_proto_rawDescGZIP(), []int{3}
}

type AlreadyRemoved struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AlreadyRemoved) Reset() {
	*x = AlreadyRemoved{}
	mi := &file_com_coralogix_archive_v2_ack_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlreadyRemoved) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlreadyRemoved) ProtoMessage() {}

func (x *AlreadyRemoved) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_ack_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlreadyRemoved.ProtoReflect.Descriptor instead.
func (*AlreadyRemoved) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_ack_proto_rawDescGZIP(), []int{4}
}

type Skipped struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Skipped) Reset() {
	*x = Skipped{}
	mi := &file_com_coralogix_archive_v2_ack_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Skipped) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Skipped) ProtoMessage() {}

func (x *Skipped) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_ack_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Skipped.ProtoReflect.Descriptor instead.
func (*Skipped) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_ack_proto_rawDescGZIP(), []int{5}
}

type Success struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Success) Reset() {
	*x = Success{}
	mi := &file_com_coralogix_archive_v2_ack_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Success) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Success) ProtoMessage() {}

func (x *Success) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_ack_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Success.ProtoReflect.Descriptor instead.
func (*Success) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_ack_proto_rawDescGZIP(), []int{6}
}

var File_com_coralogix_archive_v2_ack_proto protoreflect.FileDescriptor

var file_com_coralogix_archive_v2_ack_proto_rawDesc = []byte{
	0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x63, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x1a, 0x24,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x61, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x01, 0x0a, 0x0c, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x41, 0x63, 0x6b, 0x12, 0x47, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76,
	0x32, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x47,
	0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x07,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x22, 0xb2, 0x02, 0x0a, 0x13, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x41, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x3d, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x48, 0x00, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x50, 0x0a, 0x0e, 0x61, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x5f, 0x65, 0x78, 0x69,
	0x73, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x45, 0x78, 0x69, 0x73,
	0x74, 0x73, 0x48, 0x00, 0x52, 0x0d, 0x61, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x73, 0x12, 0x3d, 0x0a, 0x07, 0x73, 0x6b, 0x69, 0x70, 0x70, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e,
	0x53, 0x6b, 0x69, 0x70, 0x70, 0x65, 0x64, 0x48, 0x00, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x70, 0x70,
	0x65, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xf6, 0x01, 0x0a,
	0x13, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x41, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32,
	0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x52, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x3d, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x2e, 0x76, 0x32, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x48, 0x00, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x53, 0x0a, 0x0f, 0x61, 0x6c, 0x72, 0x65, 0x61, 0x64,
	0x79, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x6c, 0x72, 0x65, 0x61,
	0x64, 0x79, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0e, 0x61, 0x6c, 0x72,
	0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x0f, 0x0a, 0x0d, 0x41, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x22, 0x10, 0x0a, 0x0e, 0x41, 0x6c, 0x72, 0x65, 0x61, 0x64,
	0x79, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x22, 0x09, 0x0a, 0x07, 0x53, 0x6b, 0x69, 0x70,
	0x70, 0x65, 0x64, 0x22, 0x09, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_archive_v2_ack_proto_rawDescOnce sync.Once
	file_com_coralogix_archive_v2_ack_proto_rawDescData = file_com_coralogix_archive_v2_ack_proto_rawDesc
)

func file_com_coralogix_archive_v2_ack_proto_rawDescGZIP() []byte {
	file_com_coralogix_archive_v2_ack_proto_rawDescOnce.Do(func() {
		file_com_coralogix_archive_v2_ack_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_archive_v2_ack_proto_rawDescData)
	})
	return file_com_coralogix_archive_v2_ack_proto_rawDescData
}

var file_com_coralogix_archive_v2_ack_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_com_coralogix_archive_v2_ack_proto_goTypes = []any{
	(*MetastoreAck)(nil),        // 0: com.coralogix.archive.v2.MetastoreAck
	(*ObjectCreatedStatus)(nil), // 1: com.coralogix.archive.v2.ObjectCreatedStatus
	(*ObjectRemovedStatus)(nil), // 2: com.coralogix.archive.v2.ObjectRemovedStatus
	(*AlreadyExists)(nil),       // 3: com.coralogix.archive.v2.AlreadyExists
	(*AlreadyRemoved)(nil),      // 4: com.coralogix.archive.v2.AlreadyRemoved
	(*Skipped)(nil),             // 5: com.coralogix.archive.v2.Skipped
	(*Success)(nil),             // 6: com.coralogix.archive.v2.Success
	(*ObjectCreated)(nil),       // 7: com.coralogix.archive.v2.ObjectCreated
	(*ObjectRemoved)(nil),       // 8: com.coralogix.archive.v2.ObjectRemoved
}
var file_com_coralogix_archive_v2_ack_proto_depIdxs = []int32{
	1, // 0: com.coralogix.archive.v2.MetastoreAck.created:type_name -> com.coralogix.archive.v2.ObjectCreatedStatus
	2, // 1: com.coralogix.archive.v2.MetastoreAck.removed:type_name -> com.coralogix.archive.v2.ObjectRemovedStatus
	7, // 2: com.coralogix.archive.v2.ObjectCreatedStatus.command:type_name -> com.coralogix.archive.v2.ObjectCreated
	6, // 3: com.coralogix.archive.v2.ObjectCreatedStatus.success:type_name -> com.coralogix.archive.v2.Success
	3, // 4: com.coralogix.archive.v2.ObjectCreatedStatus.already_exists:type_name -> com.coralogix.archive.v2.AlreadyExists
	5, // 5: com.coralogix.archive.v2.ObjectCreatedStatus.skipped:type_name -> com.coralogix.archive.v2.Skipped
	8, // 6: com.coralogix.archive.v2.ObjectRemovedStatus.command:type_name -> com.coralogix.archive.v2.ObjectRemoved
	6, // 7: com.coralogix.archive.v2.ObjectRemovedStatus.success:type_name -> com.coralogix.archive.v2.Success
	4, // 8: com.coralogix.archive.v2.ObjectRemovedStatus.already_removed:type_name -> com.coralogix.archive.v2.AlreadyRemoved
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_com_coralogix_archive_v2_ack_proto_init() }
func file_com_coralogix_archive_v2_ack_proto_init() {
	if File_com_coralogix_archive_v2_ack_proto != nil {
		return
	}
	file_com_coralogix_archive_v2_event_proto_init()
	file_com_coralogix_archive_v2_ack_proto_msgTypes[1].OneofWrappers = []any{
		(*ObjectCreatedStatus_Success)(nil),
		(*ObjectCreatedStatus_AlreadyExists)(nil),
		(*ObjectCreatedStatus_Skipped)(nil),
	}
	file_com_coralogix_archive_v2_ack_proto_msgTypes[2].OneofWrappers = []any{
		(*ObjectRemovedStatus_Success)(nil),
		(*ObjectRemovedStatus_AlreadyRemoved)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_archive_v2_ack_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_archive_v2_ack_proto_goTypes,
		DependencyIndexes: file_com_coralogix_archive_v2_ack_proto_depIdxs,
		MessageInfos:      file_com_coralogix_archive_v2_ack_proto_msgTypes,
	}.Build()
	File_com_coralogix_archive_v2_ack_proto = out.File
	file_com_coralogix_archive_v2_ack_proto_rawDesc = nil
	file_com_coralogix_archive_v2_ack_proto_goTypes = nil
	file_com_coralogix_archive_v2_ack_proto_depIdxs = nil
}
