// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: com/coralogix/archive/kafka_out_targets/v1/kafka_out_targets_service.proto

package v1

import (
	_ "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/archive/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CompressionType int32

const (
	CompressionType_COMPRESSION_TYPE_UNSPECIFIED CompressionType = 0
	CompressionType_COMPRESSION_TYPE_GZIP        CompressionType = 1
)

// Enum value maps for CompressionType.
var (
	CompressionType_name = map[int32]string{
		0: "COMPRESSION_TYPE_UNSPECIFIED",
		1: "COMPRESSION_TYPE_GZIP",
	}
	CompressionType_value = map[string]int32{
		"COMPRESSION_TYPE_UNSPECIFIED": 0,
		"COMPRESSION_TYPE_GZIP":        1,
	}
)

func (x CompressionType) Enum() *CompressionType {
	p := new(CompressionType)
	*p = x
	return p
}

func (x CompressionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CompressionType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_enumTypes[0].Descriptor()
}

func (CompressionType) Type() protoreflect.EnumType {
	return &file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_enumTypes[0]
}

func (x CompressionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CompressionType.Descriptor instead.
func (CompressionType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_rawDescGZIP(), []int{0}
}

type GetKafkaOutTargetsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetKafkaOutTargetsRequest) Reset() {
	*x = GetKafkaOutTargetsRequest{}
	mi := &file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetKafkaOutTargetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKafkaOutTargetsRequest) ProtoMessage() {}

func (x *GetKafkaOutTargetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_msgTypes[0]
	if x != nil {
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
	return file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_rawDescGZIP(), []int{0}
}

type GetKafkaOutTargetsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Targets       []*Target              `protobuf:"bytes,1,rep,name=targets,proto3" json:"targets,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetKafkaOutTargetsResponse) Reset() {
	*x = GetKafkaOutTargetsResponse{}
	mi := &file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetKafkaOutTargetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKafkaOutTargetsResponse) ProtoMessage() {}

func (x *GetKafkaOutTargetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_msgTypes[1]
	if x != nil {
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
	return file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetKafkaOutTargetsResponse) GetTargets() []*Target {
	if x != nil {
		return x.Targets
	}
	return nil
}

type UpsertKafkaOutTargetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Target        *Target                `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpsertKafkaOutTargetRequest) Reset() {
	*x = UpsertKafkaOutTargetRequest{}
	mi := &file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertKafkaOutTargetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertKafkaOutTargetRequest) ProtoMessage() {}

func (x *UpsertKafkaOutTargetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_msgTypes[2]
	if x != nil {
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
	return file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpsertKafkaOutTargetRequest) GetTarget() *Target {
	if x != nil {
		return x.Target
	}
	return nil
}

type UpsertKafkaOutTargetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Target        *Target                `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpsertKafkaOutTargetResponse) Reset() {
	*x = UpsertKafkaOutTargetResponse{}
	mi := &file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertKafkaOutTargetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertKafkaOutTargetResponse) ProtoMessage() {}

func (x *UpsertKafkaOutTargetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_msgTypes[3]
	if x != nil {
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
	return file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpsertKafkaOutTargetResponse) GetTarget() *Target {
	if x != nil {
		return x.Target
	}
	return nil
}

type DeleteKafkaOutTargetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteKafkaOutTargetRequest) Reset() {
	*x = DeleteKafkaOutTargetRequest{}
	mi := &file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteKafkaOutTargetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteKafkaOutTargetRequest) ProtoMessage() {}

func (x *DeleteKafkaOutTargetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_msgTypes[4]
	if x != nil {
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
	return file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteKafkaOutTargetRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteKafkaOutTargetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteKafkaOutTargetResponse) Reset() {
	*x = DeleteKafkaOutTargetResponse{}
	mi := &file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteKafkaOutTargetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteKafkaOutTargetResponse) ProtoMessage() {}

func (x *DeleteKafkaOutTargetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_msgTypes[5]
	if x != nil {
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
	return file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_rawDescGZIP(), []int{5}
}

type SendKafkaOutTestMessageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Target        *Target                `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendKafkaOutTestMessageRequest) Reset() {
	*x = SendKafkaOutTestMessageRequest{}
	mi := &file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendKafkaOutTestMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendKafkaOutTestMessageRequest) ProtoMessage() {}

func (x *SendKafkaOutTestMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_msgTypes[6]
	if x != nil {
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
	return file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_rawDescGZIP(), []int{6}
}

func (x *SendKafkaOutTestMessageRequest) GetTarget() *Target {
	if x != nil {
		return x.Target
	}
	return nil
}

type SendKafkaOutTestMessageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendKafkaOutTestMessageResponse) Reset() {
	*x = SendKafkaOutTestMessageResponse{}
	mi := &file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendKafkaOutTestMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendKafkaOutTestMessageResponse) ProtoMessage() {}

func (x *SendKafkaOutTestMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_msgTypes[7]
	if x != nil {
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
	return file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_rawDescGZIP(), []int{7}
}

type Target struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Id              uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsActive        bool                   `protobuf:"varint,4,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	Topic           string                 `protobuf:"bytes,5,opt,name=topic,proto3" json:"topic,omitempty"`
	Brokers         string                 `protobuf:"bytes,6,opt,name=brokers,proto3" json:"brokers,omitempty"`
	DpxlExpression  string                 `protobuf:"bytes,7,opt,name=dpxl_expression,json=dpxlExpression,proto3" json:"dpxl_expression,omitempty"`
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	ApiKey          string                 `protobuf:"bytes,10,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	CompressionType *CompressionType       `protobuf:"varint,11,opt,name=compression_type,json=compressionType,proto3,enum=com.coralogix.archive.kafka_out_targets.v1.CompressionType,oneof" json:"compression_type,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Target) Reset() {
	*x = Target{}
	mi := &file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Target) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Target) ProtoMessage() {}

func (x *Target) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_msgTypes[8]
	if x != nil {
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
	return file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_rawDescGZIP(), []int{8}
}

func (x *Target) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Target) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Target) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *Target) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Target) GetBrokers() string {
	if x != nil {
		return x.Brokers
	}
	return ""
}

func (x *Target) GetDpxlExpression() string {
	if x != nil {
		return x.DpxlExpression
	}
	return ""
}

func (x *Target) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Target) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Target) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *Target) GetCompressionType() CompressionType {
	if x != nil && x.CompressionType != nil {
		return *x.CompressionType
	}
	return CompressionType_COMPRESSION_TYPE_UNSPECIFIED
}

var File_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto protoreflect.FileDescriptor

var file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_rawDesc = []byte{
	0x0a, 0x4a, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x6f, 0x75,
	0x74, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x61, 0x66,
	0x6b, 0x61, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x28, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x1b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4f,
	0x75, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x6a, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4f, 0x75, 0x74, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c,
	0x0a, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x6f, 0x75,
	0x74, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x22, 0x69, 0x0a, 0x1b,
	0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4f, 0x75, 0x74, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x06, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52,
	0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x6a, 0x0a, 0x1c, 0x55, 0x70, 0x73, 0x65, 0x72,
	0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4f, 0x75, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e,
	0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x22, 0x2d, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x61, 0x66,
	0x6b, 0x61, 0x4f, 0x75, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x1e, 0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x61, 0x66, 0x6b,
	0x61, 0x4f, 0x75, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x6c, 0x0a, 0x1e, 0x53, 0x65, 0x6e, 0x64, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4f,
	0x75, 0x74, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x6b, 0x61, 0x66,
	0x6b, 0x61, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x22, 0x21, 0x0a, 0x1f, 0x53, 0x65, 0x6e, 0x64, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4f, 0x75, 0x74,
	0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0xb3, 0x03, 0x0a, 0x06, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x73, 0x12,
	0x27, 0x0a, 0x0f, 0x64, 0x70, 0x78, 0x6c, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x70, 0x78, 0x6c, 0x45, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x6b, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x3b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f,
	0x6f, 0x75, 0x74, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00,
	0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2a, 0x4e, 0x0a, 0x0f, 0x43, 0x6f, 0x6d,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x1c,
	0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19,
	0x0a, 0x15, 0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x47, 0x5a, 0x49, 0x50, 0x10, 0x01, 0x32, 0xc8, 0x06, 0x0a, 0x15, 0x4b, 0x61,
	0x66, 0x6b, 0x61, 0x4f, 0x75, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0xc0, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61,
	0x4f, 0x75, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x45, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61,
	0x4f, 0x75, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x46, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f,
	0x6f, 0x75, 0x74, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4f, 0x75, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0xc2, 0xb8, 0x02, 0x17, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x20, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x20, 0x6f, 0x75, 0x74, 0x20, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0xc8, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x73, 0x65, 0x72,
	0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4f, 0x75, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12,
	0x47, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x6f, 0x75,
	0x74, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x73,
	0x65, 0x72, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4f, 0x75, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x48, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4b, 0x61, 0x66, 0x6b,
	0x61, 0x4f, 0x75, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1d, 0xc2, 0xb8, 0x02, 0x19, 0x0a, 0x17, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74,
	0x20, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x20, 0x6f, 0x75, 0x74, 0x20, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x12, 0xc8, 0x01, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x61, 0x66, 0x6b,
	0x61, 0x4f, 0x75, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x47, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x61,
	0x66, 0x6b, 0x61, 0x4f, 0x75, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x48, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x6b, 0x61, 0x66, 0x6b,
	0x61, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4f, 0x75, 0x74, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0xc2,
	0xb8, 0x02, 0x19, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x6b, 0x61, 0x66, 0x6b,
	0x61, 0x20, 0x6f, 0x75, 0x74, 0x20, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0xd5, 0x01, 0x0a,
	0x17, 0x53, 0x65, 0x6e, 0x64, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4f, 0x75, 0x74, 0x54, 0x65, 0x73,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4f,
	0x75, 0x74, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x4b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x6b, 0x61, 0x66,
	0x6b, 0x61, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4f, 0x75, 0x74, 0x54, 0x65,
	0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x21, 0xc2, 0xb8, 0x02, 0x1d, 0x0a, 0x1b, 0x53, 0x65, 0x6e, 0x64, 0x20, 0x6b, 0x61,
	0x66, 0x6b, 0x61, 0x20, 0x6f, 0x75, 0x74, 0x20, 0x74, 0x65, 0x73, 0x74, 0x20, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_rawDescOnce sync.Once
	file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_rawDescData = file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_rawDesc
)

func file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_rawDescGZIP() []byte {
	file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_rawDescOnce.Do(func() {
		file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_rawDescData)
	})
	return file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_rawDescData
}

var file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_goTypes = []any{
	(CompressionType)(0),                    // 0: com.coralogix.archive.kafka_out_targets.v1.CompressionType
	(*GetKafkaOutTargetsRequest)(nil),       // 1: com.coralogix.archive.kafka_out_targets.v1.GetKafkaOutTargetsRequest
	(*GetKafkaOutTargetsResponse)(nil),      // 2: com.coralogix.archive.kafka_out_targets.v1.GetKafkaOutTargetsResponse
	(*UpsertKafkaOutTargetRequest)(nil),     // 3: com.coralogix.archive.kafka_out_targets.v1.UpsertKafkaOutTargetRequest
	(*UpsertKafkaOutTargetResponse)(nil),    // 4: com.coralogix.archive.kafka_out_targets.v1.UpsertKafkaOutTargetResponse
	(*DeleteKafkaOutTargetRequest)(nil),     // 5: com.coralogix.archive.kafka_out_targets.v1.DeleteKafkaOutTargetRequest
	(*DeleteKafkaOutTargetResponse)(nil),    // 6: com.coralogix.archive.kafka_out_targets.v1.DeleteKafkaOutTargetResponse
	(*SendKafkaOutTestMessageRequest)(nil),  // 7: com.coralogix.archive.kafka_out_targets.v1.SendKafkaOutTestMessageRequest
	(*SendKafkaOutTestMessageResponse)(nil), // 8: com.coralogix.archive.kafka_out_targets.v1.SendKafkaOutTestMessageResponse
	(*Target)(nil),                          // 9: com.coralogix.archive.kafka_out_targets.v1.Target
	(*timestamppb.Timestamp)(nil),           // 10: google.protobuf.Timestamp
}
var file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_depIdxs = []int32{
	9,  // 0: com.coralogix.archive.kafka_out_targets.v1.GetKafkaOutTargetsResponse.targets:type_name -> com.coralogix.archive.kafka_out_targets.v1.Target
	9,  // 1: com.coralogix.archive.kafka_out_targets.v1.UpsertKafkaOutTargetRequest.target:type_name -> com.coralogix.archive.kafka_out_targets.v1.Target
	9,  // 2: com.coralogix.archive.kafka_out_targets.v1.UpsertKafkaOutTargetResponse.target:type_name -> com.coralogix.archive.kafka_out_targets.v1.Target
	9,  // 3: com.coralogix.archive.kafka_out_targets.v1.SendKafkaOutTestMessageRequest.target:type_name -> com.coralogix.archive.kafka_out_targets.v1.Target
	10, // 4: com.coralogix.archive.kafka_out_targets.v1.Target.created_at:type_name -> google.protobuf.Timestamp
	10, // 5: com.coralogix.archive.kafka_out_targets.v1.Target.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 6: com.coralogix.archive.kafka_out_targets.v1.Target.compression_type:type_name -> com.coralogix.archive.kafka_out_targets.v1.CompressionType
	1,  // 7: com.coralogix.archive.kafka_out_targets.v1.KafkaOutTargetService.GetKafkaOutTargets:input_type -> com.coralogix.archive.kafka_out_targets.v1.GetKafkaOutTargetsRequest
	3,  // 8: com.coralogix.archive.kafka_out_targets.v1.KafkaOutTargetService.UpsertKafkaOutTarget:input_type -> com.coralogix.archive.kafka_out_targets.v1.UpsertKafkaOutTargetRequest
	5,  // 9: com.coralogix.archive.kafka_out_targets.v1.KafkaOutTargetService.DeleteKafkaOutTarget:input_type -> com.coralogix.archive.kafka_out_targets.v1.DeleteKafkaOutTargetRequest
	7,  // 10: com.coralogix.archive.kafka_out_targets.v1.KafkaOutTargetService.SendKafkaOutTestMessage:input_type -> com.coralogix.archive.kafka_out_targets.v1.SendKafkaOutTestMessageRequest
	2,  // 11: com.coralogix.archive.kafka_out_targets.v1.KafkaOutTargetService.GetKafkaOutTargets:output_type -> com.coralogix.archive.kafka_out_targets.v1.GetKafkaOutTargetsResponse
	4,  // 12: com.coralogix.archive.kafka_out_targets.v1.KafkaOutTargetService.UpsertKafkaOutTarget:output_type -> com.coralogix.archive.kafka_out_targets.v1.UpsertKafkaOutTargetResponse
	6,  // 13: com.coralogix.archive.kafka_out_targets.v1.KafkaOutTargetService.DeleteKafkaOutTarget:output_type -> com.coralogix.archive.kafka_out_targets.v1.DeleteKafkaOutTargetResponse
	8,  // 14: com.coralogix.archive.kafka_out_targets.v1.KafkaOutTargetService.SendKafkaOutTestMessage:output_type -> com.coralogix.archive.kafka_out_targets.v1.SendKafkaOutTestMessageResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_init() }
func file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_init() {
	if File_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto != nil {
		return
	}
	file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_msgTypes[8].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_goTypes,
		DependencyIndexes: file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_depIdxs,
		EnumInfos:         file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_enumTypes,
		MessageInfos:      file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_msgTypes,
	}.Build()
	File_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto = out.File
	file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_rawDesc = nil
	file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_goTypes = nil
	file_com_coralogix_archive_kafka_out_targets_v1_kafka_out_targets_service_proto_depIdxs = nil
}
