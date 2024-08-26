// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.3
// source: com/coralogix/archive/kafka_out_targets/v2/kafka_out_target_service.proto

package v2

import (
	_ "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/archive/v1"
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

type GetKafkaOutTargetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetKafkaOutTargetsRequest) Reset() {
	*x = GetKafkaOutTargetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKafkaOutTargetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKafkaOutTargetsRequest) ProtoMessage() {}

func (x *GetKafkaOutTargetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKafkaOutTargetsRequest.ProtoReflect.Descriptor instead.
func (*GetKafkaOutTargetsRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_rawDescGZIP(), []int{0}
}

type GetKafkaOutTargetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Targets []*Target `protobuf:"bytes,1,rep,name=targets,proto3" json:"targets,omitempty"`
}

func (x *GetKafkaOutTargetsResponse) Reset() {
	*x = GetKafkaOutTargetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKafkaOutTargetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKafkaOutTargetsResponse) ProtoMessage() {}

func (x *GetKafkaOutTargetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKafkaOutTargetsResponse.ProtoReflect.Descriptor instead.
func (*GetKafkaOutTargetsResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetKafkaOutTargetsResponse) GetTargets() []*Target {
	if x != nil {
		return x.Targets
	}
	return nil
}

type UpsertKafkaOutTargetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target *Target `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *UpsertKafkaOutTargetRequest) Reset() {
	*x = UpsertKafkaOutTargetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertKafkaOutTargetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertKafkaOutTargetRequest) ProtoMessage() {}

func (x *UpsertKafkaOutTargetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertKafkaOutTargetRequest.ProtoReflect.Descriptor instead.
func (*UpsertKafkaOutTargetRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpsertKafkaOutTargetRequest) GetTarget() *Target {
	if x != nil {
		return x.Target
	}
	return nil
}

type UpsertKafkaOutTargetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target *Target `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *UpsertKafkaOutTargetResponse) Reset() {
	*x = UpsertKafkaOutTargetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertKafkaOutTargetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertKafkaOutTargetResponse) ProtoMessage() {}

func (x *UpsertKafkaOutTargetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertKafkaOutTargetResponse.ProtoReflect.Descriptor instead.
func (*UpsertKafkaOutTargetResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpsertKafkaOutTargetResponse) GetTarget() *Target {
	if x != nil {
		return x.Target
	}
	return nil
}

type DeleteKafkaOutTargetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteKafkaOutTargetRequest) Reset() {
	*x = DeleteKafkaOutTargetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteKafkaOutTargetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteKafkaOutTargetRequest) ProtoMessage() {}

func (x *DeleteKafkaOutTargetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteKafkaOutTargetRequest.ProtoReflect.Descriptor instead.
func (*DeleteKafkaOutTargetRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteKafkaOutTargetRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteKafkaOutTargetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteKafkaOutTargetResponse) Reset() {
	*x = DeleteKafkaOutTargetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteKafkaOutTargetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteKafkaOutTargetResponse) ProtoMessage() {}

func (x *DeleteKafkaOutTargetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteKafkaOutTargetResponse.ProtoReflect.Descriptor instead.
func (*DeleteKafkaOutTargetResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_rawDescGZIP(), []int{5}
}

type SendKafkaOutTestMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target *Target `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *SendKafkaOutTestMessageRequest) Reset() {
	*x = SendKafkaOutTestMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendKafkaOutTestMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendKafkaOutTestMessageRequest) ProtoMessage() {}

func (x *SendKafkaOutTestMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendKafkaOutTestMessageRequest.ProtoReflect.Descriptor instead.
func (*SendKafkaOutTestMessageRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_rawDescGZIP(), []int{6}
}

func (x *SendKafkaOutTestMessageRequest) GetTarget() *Target {
	if x != nil {
		return x.Target
	}
	return nil
}

type SendKafkaOutTestMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendKafkaOutTestMessageResponse) Reset() {
	*x = SendKafkaOutTestMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendKafkaOutTestMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendKafkaOutTestMessageResponse) ProtoMessage() {}

func (x *SendKafkaOutTestMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendKafkaOutTestMessageResponse.ProtoReflect.Descriptor instead.
func (*SendKafkaOutTestMessageResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_rawDescGZIP(), []int{7}
}

var File_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto protoreflect.FileDescriptor

var file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_rawDesc = []byte{
	0x0a, 0x49, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x6f, 0x75,
	0x74, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x6b, 0x61, 0x66,
	0x6b, 0x61, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x32, 0x1a, 0x28, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x41, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x6f,
	0x75, 0x74, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x6b, 0x61,
	0x66, 0x6b, 0x61, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61,
	0x4f, 0x75, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x6a, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4f, 0x75, 0x74,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4c, 0x0a, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x6f,
	0x75, 0x74, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x22, 0x69, 0x0a,
	0x1b, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4f, 0x75, 0x74, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x06,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x6a, 0x0a, 0x1c, 0x55, 0x70, 0x73, 0x65,
	0x72, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4f, 0x75, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x22, 0x2d, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x61,
	0x66, 0x6b, 0x61, 0x4f, 0x75, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x1e, 0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x61, 0x66,
	0x6b, 0x61, 0x4f, 0x75, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x6c, 0x0a, 0x1e, 0x53, 0x65, 0x6e, 0x64, 0x4b, 0x61, 0x66, 0x6b, 0x61,
	0x4f, 0x75, 0x74, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x6b, 0x61,
	0x66, 0x6b, 0x61, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2e,
	0x76, 0x32, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x22, 0x21, 0x0a, 0x1f, 0x53, 0x65, 0x6e, 0x64, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4f, 0x75,
	0x74, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0xc8, 0x06, 0x0a, 0x15, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4f, 0x75,
	0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xc0,
	0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4f, 0x75, 0x74, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x45, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x6b, 0x61,
	0x66, 0x6b, 0x61, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2e,
	0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4f, 0x75, 0x74, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x46, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x66,
	0x6b, 0x61, 0x4f, 0x75, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0xc2, 0xb8, 0x02, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x20,
	0x6b, 0x61, 0x66, 0x6b, 0x61, 0x20, 0x6f, 0x75, 0x74, 0x20, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x73, 0x12, 0xc8, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4b, 0x61, 0x66, 0x6b,
	0x61, 0x4f, 0x75, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x47, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4b, 0x61,
	0x66, 0x6b, 0x61, 0x4f, 0x75, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x48, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x6b, 0x61, 0x66, 0x6b,
	0x61, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x32,
	0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4f, 0x75, 0x74, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0xc2,
	0xb8, 0x02, 0x19, 0x0a, 0x17, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x20, 0x6b, 0x61, 0x66, 0x6b,
	0x61, 0x20, 0x6f, 0x75, 0x74, 0x20, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0xc8, 0x01, 0x0a,
	0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4f, 0x75, 0x74, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x47, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x6b, 0x61,
	0x66, 0x6b, 0x61, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2e,
	0x76, 0x32, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4f, 0x75,
	0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x48,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x6f, 0x75, 0x74,
	0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4f, 0x75, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0xc2, 0xb8, 0x02, 0x19, 0x0a, 0x17,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x20, 0x6f, 0x75, 0x74,
	0x20, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0xd5, 0x01, 0x0a, 0x17, 0x53, 0x65, 0x6e, 0x64,
	0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4f, 0x75, 0x74, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x4a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x6b, 0x61, 0x66, 0x6b,
	0x61, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x32,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4f, 0x75, 0x74, 0x54, 0x65, 0x73,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x4b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x6f, 0x75,
	0x74, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4f, 0x75, 0x74, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0xc2, 0xb8,
	0x02, 0x1d, 0x0a, 0x1b, 0x53, 0x65, 0x6e, 0x64, 0x20, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x20, 0x6f,
	0x75, 0x74, 0x20, 0x74, 0x65, 0x73, 0x74, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_rawDescOnce sync.Once
	file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_rawDescData = file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_rawDesc
)

func file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_rawDescGZIP() []byte {
	file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_rawDescOnce.Do(func() {
		file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_rawDescData)
	})
	return file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_rawDescData
}

var file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_goTypes = []interface{}{
	(*GetKafkaOutTargetsRequest)(nil),       // 0: com.coralogix.archive.kafka_out_targets.v2.GetKafkaOutTargetsRequest
	(*GetKafkaOutTargetsResponse)(nil),      // 1: com.coralogix.archive.kafka_out_targets.v2.GetKafkaOutTargetsResponse
	(*UpsertKafkaOutTargetRequest)(nil),     // 2: com.coralogix.archive.kafka_out_targets.v2.UpsertKafkaOutTargetRequest
	(*UpsertKafkaOutTargetResponse)(nil),    // 3: com.coralogix.archive.kafka_out_targets.v2.UpsertKafkaOutTargetResponse
	(*DeleteKafkaOutTargetRequest)(nil),     // 4: com.coralogix.archive.kafka_out_targets.v2.DeleteKafkaOutTargetRequest
	(*DeleteKafkaOutTargetResponse)(nil),    // 5: com.coralogix.archive.kafka_out_targets.v2.DeleteKafkaOutTargetResponse
	(*SendKafkaOutTestMessageRequest)(nil),  // 6: com.coralogix.archive.kafka_out_targets.v2.SendKafkaOutTestMessageRequest
	(*SendKafkaOutTestMessageResponse)(nil), // 7: com.coralogix.archive.kafka_out_targets.v2.SendKafkaOutTestMessageResponse
	(*Target)(nil),                          // 8: com.coralogix.archive.kafka_out_targets.v2.Target
}
var file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_depIdxs = []int32{
	8, // 0: com.coralogix.archive.kafka_out_targets.v2.GetKafkaOutTargetsResponse.targets:type_name -> com.coralogix.archive.kafka_out_targets.v2.Target
	8, // 1: com.coralogix.archive.kafka_out_targets.v2.UpsertKafkaOutTargetRequest.target:type_name -> com.coralogix.archive.kafka_out_targets.v2.Target
	8, // 2: com.coralogix.archive.kafka_out_targets.v2.UpsertKafkaOutTargetResponse.target:type_name -> com.coralogix.archive.kafka_out_targets.v2.Target
	8, // 3: com.coralogix.archive.kafka_out_targets.v2.SendKafkaOutTestMessageRequest.target:type_name -> com.coralogix.archive.kafka_out_targets.v2.Target
	0, // 4: com.coralogix.archive.kafka_out_targets.v2.KafkaOutTargetService.GetKafkaOutTargets:input_type -> com.coralogix.archive.kafka_out_targets.v2.GetKafkaOutTargetsRequest
	2, // 5: com.coralogix.archive.kafka_out_targets.v2.KafkaOutTargetService.UpsertKafkaOutTarget:input_type -> com.coralogix.archive.kafka_out_targets.v2.UpsertKafkaOutTargetRequest
	4, // 6: com.coralogix.archive.kafka_out_targets.v2.KafkaOutTargetService.DeleteKafkaOutTarget:input_type -> com.coralogix.archive.kafka_out_targets.v2.DeleteKafkaOutTargetRequest
	6, // 7: com.coralogix.archive.kafka_out_targets.v2.KafkaOutTargetService.SendKafkaOutTestMessage:input_type -> com.coralogix.archive.kafka_out_targets.v2.SendKafkaOutTestMessageRequest
	1, // 8: com.coralogix.archive.kafka_out_targets.v2.KafkaOutTargetService.GetKafkaOutTargets:output_type -> com.coralogix.archive.kafka_out_targets.v2.GetKafkaOutTargetsResponse
	3, // 9: com.coralogix.archive.kafka_out_targets.v2.KafkaOutTargetService.UpsertKafkaOutTarget:output_type -> com.coralogix.archive.kafka_out_targets.v2.UpsertKafkaOutTargetResponse
	5, // 10: com.coralogix.archive.kafka_out_targets.v2.KafkaOutTargetService.DeleteKafkaOutTarget:output_type -> com.coralogix.archive.kafka_out_targets.v2.DeleteKafkaOutTargetResponse
	7, // 11: com.coralogix.archive.kafka_out_targets.v2.KafkaOutTargetService.SendKafkaOutTestMessage:output_type -> com.coralogix.archive.kafka_out_targets.v2.SendKafkaOutTestMessageResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_init() }
func file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_init() {
	if File_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto != nil {
		return
	}
	file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKafkaOutTargetsRequest); i {
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
		file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKafkaOutTargetsResponse); i {
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
		file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertKafkaOutTargetRequest); i {
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
		file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertKafkaOutTargetResponse); i {
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
		file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteKafkaOutTargetRequest); i {
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
		file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteKafkaOutTargetResponse); i {
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
		file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendKafkaOutTestMessageRequest); i {
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
		file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendKafkaOutTestMessageResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_goTypes,
		DependencyIndexes: file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_depIdxs,
		MessageInfos:      file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_msgTypes,
	}.Build()
	File_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto = out.File
	file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_rawDesc = nil
	file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_goTypes = nil
	file_com_coralogix_archive_kafka_out_targets_v2_kafka_out_target_service_proto_depIdxs = nil
}
