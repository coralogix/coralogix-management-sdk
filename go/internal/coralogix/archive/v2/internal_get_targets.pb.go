// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: com/coralogix/archive/v2/internal_get_targets.proto

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

type ArchivingFormatId int32

const (
	ArchivingFormatId_ARCHIVING_FORMAT_ID_UNSPECIFIED     ArchivingFormatId = 0
	ArchivingFormatId_ARCHIVING_FORMAT_ID_WIDE_PARQUET_V1 ArchivingFormatId = 1
)

// Enum value maps for ArchivingFormatId.
var (
	ArchivingFormatId_name = map[int32]string{
		0: "ARCHIVING_FORMAT_ID_UNSPECIFIED",
		1: "ARCHIVING_FORMAT_ID_WIDE_PARQUET_V1",
	}
	ArchivingFormatId_value = map[string]int32{
		"ARCHIVING_FORMAT_ID_UNSPECIFIED":     0,
		"ARCHIVING_FORMAT_ID_WIDE_PARQUET_V1": 1,
	}
)

func (x ArchivingFormatId) Enum() *ArchivingFormatId {
	p := new(ArchivingFormatId)
	*p = x
	return p
}

func (x ArchivingFormatId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ArchivingFormatId) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_archive_v2_internal_get_targets_proto_enumTypes[0].Descriptor()
}

func (ArchivingFormatId) Type() protoreflect.EnumType {
	return &file_com_coralogix_archive_v2_internal_get_targets_proto_enumTypes[0]
}

func (x ArchivingFormatId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ArchivingFormatId.Descriptor instead.
func (ArchivingFormatId) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_internal_get_targets_proto_rawDescGZIP(), []int{0}
}

type OrderBy int32

const (
	OrderBy_ORDER_BY_UNSPECIFIED OrderBy = 0
	OrderBy_ORDER_BY_UPDATED_AT  OrderBy = 1
	OrderBy_ORDER_BY_CREATED_AT  OrderBy = 2
	OrderBy_ORDER_BY_COMPANY_ID  OrderBy = 3
)

// Enum value maps for OrderBy.
var (
	OrderBy_name = map[int32]string{
		0: "ORDER_BY_UNSPECIFIED",
		1: "ORDER_BY_UPDATED_AT",
		2: "ORDER_BY_CREATED_AT",
		3: "ORDER_BY_COMPANY_ID",
	}
	OrderBy_value = map[string]int32{
		"ORDER_BY_UNSPECIFIED": 0,
		"ORDER_BY_UPDATED_AT":  1,
		"ORDER_BY_CREATED_AT":  2,
		"ORDER_BY_COMPANY_ID":  3,
	}
)

func (x OrderBy) Enum() *OrderBy {
	p := new(OrderBy)
	*p = x
	return p
}

func (x OrderBy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderBy) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_archive_v2_internal_get_targets_proto_enumTypes[1].Descriptor()
}

func (OrderBy) Type() protoreflect.EnumType {
	return &file_com_coralogix_archive_v2_internal_get_targets_proto_enumTypes[1]
}

func (x OrderBy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderBy.Descriptor instead.
func (OrderBy) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_internal_get_targets_proto_rawDescGZIP(), []int{1}
}

type OrderDirection int32

const (
	OrderDirection_ORDER_DIRECTION_UNSPECIFIED OrderDirection = 0
	OrderDirection_ORDER_DIRECTION_ASC         OrderDirection = 1
	OrderDirection_ORDER_DIRECTION_DESC        OrderDirection = 2
)

// Enum value maps for OrderDirection.
var (
	OrderDirection_name = map[int32]string{
		0: "ORDER_DIRECTION_UNSPECIFIED",
		1: "ORDER_DIRECTION_ASC",
		2: "ORDER_DIRECTION_DESC",
	}
	OrderDirection_value = map[string]int32{
		"ORDER_DIRECTION_UNSPECIFIED": 0,
		"ORDER_DIRECTION_ASC":         1,
		"ORDER_DIRECTION_DESC":        2,
	}
)

func (x OrderDirection) Enum() *OrderDirection {
	p := new(OrderDirection)
	*p = x
	return p
}

func (x OrderDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_archive_v2_internal_get_targets_proto_enumTypes[2].Descriptor()
}

func (OrderDirection) Type() protoreflect.EnumType {
	return &file_com_coralogix_archive_v2_internal_get_targets_proto_enumTypes[2]
}

func (x OrderDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderDirection.Descriptor instead.
func (OrderDirection) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_internal_get_targets_proto_rawDescGZIP(), []int{2}
}

type InternalTargetServiceGetTargetsRequest struct {
	state         protoimpl.MessageState           `protogen:"open.v1"`
	Filter        *InternalGetTargetsRequestFilter `protobuf:"bytes,1,opt,name=filter,proto3,oneof" json:"filter,omitempty"`
	Order         *TargetResponseOrder             `protobuf:"bytes,2,opt,name=order,proto3,oneof" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InternalTargetServiceGetTargetsRequest) Reset() {
	*x = InternalTargetServiceGetTargetsRequest{}
	mi := &file_com_coralogix_archive_v2_internal_get_targets_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InternalTargetServiceGetTargetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalTargetServiceGetTargetsRequest) ProtoMessage() {}

func (x *InternalTargetServiceGetTargetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_internal_get_targets_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalTargetServiceGetTargetsRequest.ProtoReflect.Descriptor instead.
func (*InternalTargetServiceGetTargetsRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_internal_get_targets_proto_rawDescGZIP(), []int{0}
}

func (x *InternalTargetServiceGetTargetsRequest) GetFilter() *InternalGetTargetsRequestFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *InternalTargetServiceGetTargetsRequest) GetOrder() *TargetResponseOrder {
	if x != nil {
		return x.Order
	}
	return nil
}

type InternalTargetServiceGetTargetsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Targets       []*TargetResponse      `protobuf:"bytes,1,rep,name=targets,proto3" json:"targets,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InternalTargetServiceGetTargetsResponse) Reset() {
	*x = InternalTargetServiceGetTargetsResponse{}
	mi := &file_com_coralogix_archive_v2_internal_get_targets_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InternalTargetServiceGetTargetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalTargetServiceGetTargetsResponse) ProtoMessage() {}

func (x *InternalTargetServiceGetTargetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_internal_get_targets_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalTargetServiceGetTargetsResponse.ProtoReflect.Descriptor instead.
func (*InternalTargetServiceGetTargetsResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_internal_get_targets_proto_rawDescGZIP(), []int{1}
}

func (x *InternalTargetServiceGetTargetsResponse) GetTargets() []*TargetResponse {
	if x != nil {
		return x.Targets
	}
	return nil
}

type InternalGetTargetsRequestFilter struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CompanyIds    []uint32               `protobuf:"varint,1,rep,packed,name=company_ids,json=companyIds,proto3" json:"company_ids,omitempty"`
	FormatId      []ArchivingFormatId    `protobuf:"varint,2,rep,packed,name=format_id,json=formatId,proto3,enum=com.coralogix.archive.v2.ArchivingFormatId" json:"format_id,omitempty"`
	IsActive      *bool                  `protobuf:"varint,3,opt,name=is_active,json=isActive,proto3,oneof" json:"is_active,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InternalGetTargetsRequestFilter) Reset() {
	*x = InternalGetTargetsRequestFilter{}
	mi := &file_com_coralogix_archive_v2_internal_get_targets_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InternalGetTargetsRequestFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalGetTargetsRequestFilter) ProtoMessage() {}

func (x *InternalGetTargetsRequestFilter) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_internal_get_targets_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalGetTargetsRequestFilter.ProtoReflect.Descriptor instead.
func (*InternalGetTargetsRequestFilter) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_internal_get_targets_proto_rawDescGZIP(), []int{2}
}

func (x *InternalGetTargetsRequestFilter) GetCompanyIds() []uint32 {
	if x != nil {
		return x.CompanyIds
	}
	return nil
}

func (x *InternalGetTargetsRequestFilter) GetFormatId() []ArchivingFormatId {
	if x != nil {
		return x.FormatId
	}
	return nil
}

func (x *InternalGetTargetsRequestFilter) GetIsActive() bool {
	if x != nil && x.IsActive != nil {
		return *x.IsActive
	}
	return false
}

type TargetResponseOrder struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	OrderBy        *OrderBy               `protobuf:"varint,1,opt,name=order_by,json=orderBy,proto3,enum=com.coralogix.archive.v2.OrderBy,oneof" json:"order_by,omitempty"`
	OrderDirection *OrderDirection        `protobuf:"varint,2,opt,name=order_direction,json=orderDirection,proto3,enum=com.coralogix.archive.v2.OrderDirection,oneof" json:"order_direction,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *TargetResponseOrder) Reset() {
	*x = TargetResponseOrder{}
	mi := &file_com_coralogix_archive_v2_internal_get_targets_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TargetResponseOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetResponseOrder) ProtoMessage() {}

func (x *TargetResponseOrder) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_internal_get_targets_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetResponseOrder.ProtoReflect.Descriptor instead.
func (*TargetResponseOrder) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_internal_get_targets_proto_rawDescGZIP(), []int{3}
}

func (x *TargetResponseOrder) GetOrderBy() OrderBy {
	if x != nil && x.OrderBy != nil {
		return *x.OrderBy
	}
	return OrderBy_ORDER_BY_UNSPECIFIED
}

func (x *TargetResponseOrder) GetOrderDirection() OrderDirection {
	if x != nil && x.OrderDirection != nil {
		return *x.OrderDirection
	}
	return OrderDirection_ORDER_DIRECTION_UNSPECIFIED
}

type TargetResponse struct {
	state                      protoimpl.MessageState              `protogen:"open.v1"`
	CompanyId                  uint32                              `protobuf:"varint,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Target                     *Target                             `protobuf:"bytes,2,opt,name=target,proto3,oneof" json:"target,omitempty"`
	AdditionalTargetProperties *InternalAdditionalTargetProperties `protobuf:"bytes,3,opt,name=additional_target_properties,json=additionalTargetProperties,proto3,oneof" json:"additional_target_properties,omitempty"`
	unknownFields              protoimpl.UnknownFields
	sizeCache                  protoimpl.SizeCache
}

func (x *TargetResponse) Reset() {
	*x = TargetResponse{}
	mi := &file_com_coralogix_archive_v2_internal_get_targets_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TargetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetResponse) ProtoMessage() {}

func (x *TargetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_internal_get_targets_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetResponse.ProtoReflect.Descriptor instead.
func (*TargetResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_internal_get_targets_proto_rawDescGZIP(), []int{4}
}

func (x *TargetResponse) GetCompanyId() uint32 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *TargetResponse) GetTarget() *Target {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *TargetResponse) GetAdditionalTargetProperties() *InternalAdditionalTargetProperties {
	if x != nil {
		return x.AdditionalTargetProperties
	}
	return nil
}

type InternalAdditionalTargetProperties struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BucketV2      *string                `protobuf:"bytes,1,opt,name=bucket_v2,json=bucketV2,proto3,oneof" json:"bucket_v2,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InternalAdditionalTargetProperties) Reset() {
	*x = InternalAdditionalTargetProperties{}
	mi := &file_com_coralogix_archive_v2_internal_get_targets_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InternalAdditionalTargetProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalAdditionalTargetProperties) ProtoMessage() {}

func (x *InternalAdditionalTargetProperties) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_internal_get_targets_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalAdditionalTargetProperties.ProtoReflect.Descriptor instead.
func (*InternalAdditionalTargetProperties) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_internal_get_targets_proto_rawDescGZIP(), []int{5}
}

func (x *InternalAdditionalTargetProperties) GetBucketV2() string {
	if x != nil && x.BucketV2 != nil {
		return *x.BucketV2
	}
	return ""
}

var File_com_coralogix_archive_v2_internal_get_targets_proto protoreflect.FileDescriptor

var file_com_coralogix_archive_v2_internal_get_targets_proto_rawDesc = []byte{
	0x0a, 0x33, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x1a,
	0x25, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x01, 0x0a, 0x26, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x56, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x48, 0x0a, 0x05, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x2e, 0x76, 0x32, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x01, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x6d, 0x0a, 0x27, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x22, 0xbc, 0x01, 0x0a, 0x1f, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x73, 0x12, 0x48, 0x0a, 0x09,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x69, 0x6e, 0x67, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x49, 0x64, 0x52, 0x08, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x08, 0x69, 0x73, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x69, 0x73, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0xd1, 0x01, 0x0a, 0x13, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x41,
	0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x42, 0x79, 0x48, 0x00, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x88, 0x01,
	0x01, 0x12, 0x56, 0x0a, 0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x01, 0x52, 0x0e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa0, 0x02, 0x0a, 0x0e, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x06,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x48, 0x00,
	0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x88, 0x01, 0x01, 0x12, 0x83, 0x01, 0x0a, 0x1c,
	0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73,
	0x48, 0x01, 0x52, 0x1a, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x88, 0x01,
	0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x42, 0x1f, 0x0a, 0x1d,
	0x5f, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x54, 0x0a,
	0x22, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x09, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x76, 0x32,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x56, 0x32, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x5f, 0x76, 0x32, 0x2a, 0x61, 0x0a, 0x11, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x6e, 0x67,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x1f, 0x41, 0x52, 0x43, 0x48,
	0x49, 0x56, 0x49, 0x4e, 0x47, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x49, 0x44, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x27, 0x0a,
	0x23, 0x41, 0x52, 0x43, 0x48, 0x49, 0x56, 0x49, 0x4e, 0x47, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41,
	0x54, 0x5f, 0x49, 0x44, 0x5f, 0x57, 0x49, 0x44, 0x45, 0x5f, 0x50, 0x41, 0x52, 0x51, 0x55, 0x45,
	0x54, 0x5f, 0x56, 0x31, 0x10, 0x01, 0x2a, 0x6e, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42,
	0x79, 0x12, 0x18, 0x0a, 0x14, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x4f,
	0x52, 0x44, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x5f,
	0x41, 0x54, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x42, 0x59,
	0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x41, 0x54, 0x10, 0x02, 0x12, 0x17, 0x0a,
	0x13, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x4e,
	0x59, 0x5f, 0x49, 0x44, 0x10, 0x03, 0x2a, 0x64, 0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x1b, 0x4f, 0x52, 0x44, 0x45,
	0x52, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x52, 0x44,
	0x45, 0x52, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x53, 0x43,
	0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x44, 0x49, 0x52, 0x45,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_archive_v2_internal_get_targets_proto_rawDescOnce sync.Once
	file_com_coralogix_archive_v2_internal_get_targets_proto_rawDescData = file_com_coralogix_archive_v2_internal_get_targets_proto_rawDesc
)

func file_com_coralogix_archive_v2_internal_get_targets_proto_rawDescGZIP() []byte {
	file_com_coralogix_archive_v2_internal_get_targets_proto_rawDescOnce.Do(func() {
		file_com_coralogix_archive_v2_internal_get_targets_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_archive_v2_internal_get_targets_proto_rawDescData)
	})
	return file_com_coralogix_archive_v2_internal_get_targets_proto_rawDescData
}

var file_com_coralogix_archive_v2_internal_get_targets_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_com_coralogix_archive_v2_internal_get_targets_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_com_coralogix_archive_v2_internal_get_targets_proto_goTypes = []any{
	(ArchivingFormatId)(0), // 0: com.coralogix.archive.v2.ArchivingFormatId
	(OrderBy)(0),           // 1: com.coralogix.archive.v2.OrderBy
	(OrderDirection)(0),    // 2: com.coralogix.archive.v2.OrderDirection
	(*InternalTargetServiceGetTargetsRequest)(nil),  // 3: com.coralogix.archive.v2.InternalTargetServiceGetTargetsRequest
	(*InternalTargetServiceGetTargetsResponse)(nil), // 4: com.coralogix.archive.v2.InternalTargetServiceGetTargetsResponse
	(*InternalGetTargetsRequestFilter)(nil),         // 5: com.coralogix.archive.v2.InternalGetTargetsRequestFilter
	(*TargetResponseOrder)(nil),                     // 6: com.coralogix.archive.v2.TargetResponseOrder
	(*TargetResponse)(nil),                          // 7: com.coralogix.archive.v2.TargetResponse
	(*InternalAdditionalTargetProperties)(nil),      // 8: com.coralogix.archive.v2.InternalAdditionalTargetProperties
	(*Target)(nil), // 9: com.coralogix.archive.v2.Target
}
var file_com_coralogix_archive_v2_internal_get_targets_proto_depIdxs = []int32{
	5, // 0: com.coralogix.archive.v2.InternalTargetServiceGetTargetsRequest.filter:type_name -> com.coralogix.archive.v2.InternalGetTargetsRequestFilter
	6, // 1: com.coralogix.archive.v2.InternalTargetServiceGetTargetsRequest.order:type_name -> com.coralogix.archive.v2.TargetResponseOrder
	7, // 2: com.coralogix.archive.v2.InternalTargetServiceGetTargetsResponse.targets:type_name -> com.coralogix.archive.v2.TargetResponse
	0, // 3: com.coralogix.archive.v2.InternalGetTargetsRequestFilter.format_id:type_name -> com.coralogix.archive.v2.ArchivingFormatId
	1, // 4: com.coralogix.archive.v2.TargetResponseOrder.order_by:type_name -> com.coralogix.archive.v2.OrderBy
	2, // 5: com.coralogix.archive.v2.TargetResponseOrder.order_direction:type_name -> com.coralogix.archive.v2.OrderDirection
	9, // 6: com.coralogix.archive.v2.TargetResponse.target:type_name -> com.coralogix.archive.v2.Target
	8, // 7: com.coralogix.archive.v2.TargetResponse.additional_target_properties:type_name -> com.coralogix.archive.v2.InternalAdditionalTargetProperties
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_com_coralogix_archive_v2_internal_get_targets_proto_init() }
func file_com_coralogix_archive_v2_internal_get_targets_proto_init() {
	if File_com_coralogix_archive_v2_internal_get_targets_proto != nil {
		return
	}
	file_com_coralogix_archive_v2_target_proto_init()
	file_com_coralogix_archive_v2_internal_get_targets_proto_msgTypes[0].OneofWrappers = []any{}
	file_com_coralogix_archive_v2_internal_get_targets_proto_msgTypes[2].OneofWrappers = []any{}
	file_com_coralogix_archive_v2_internal_get_targets_proto_msgTypes[3].OneofWrappers = []any{}
	file_com_coralogix_archive_v2_internal_get_targets_proto_msgTypes[4].OneofWrappers = []any{}
	file_com_coralogix_archive_v2_internal_get_targets_proto_msgTypes[5].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_archive_v2_internal_get_targets_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_archive_v2_internal_get_targets_proto_goTypes,
		DependencyIndexes: file_com_coralogix_archive_v2_internal_get_targets_proto_depIdxs,
		EnumInfos:         file_com_coralogix_archive_v2_internal_get_targets_proto_enumTypes,
		MessageInfos:      file_com_coralogix_archive_v2_internal_get_targets_proto_msgTypes,
	}.Build()
	File_com_coralogix_archive_v2_internal_get_targets_proto = out.File
	file_com_coralogix_archive_v2_internal_get_targets_proto_rawDesc = nil
	file_com_coralogix_archive_v2_internal_get_targets_proto_goTypes = nil
	file_com_coralogix_archive_v2_internal_get_targets_proto_depIdxs = nil
}
