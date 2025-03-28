// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: com/coralogix/archive/v1/target_service.proto

package v1

import (
	_ "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/common/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IsArchiveConfiguredRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CompanyIds    []uint32               `protobuf:"varint,1,rep,packed,name=company_ids,json=companyIds,proto3" json:"company_ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IsArchiveConfiguredRequest) Reset() {
	*x = IsArchiveConfiguredRequest{}
	mi := &file_com_coralogix_archive_v1_target_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IsArchiveConfiguredRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsArchiveConfiguredRequest) ProtoMessage() {}

func (x *IsArchiveConfiguredRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v1_target_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsArchiveConfiguredRequest.ProtoReflect.Descriptor instead.
func (*IsArchiveConfiguredRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v1_target_service_proto_rawDescGZIP(), []int{0}
}

func (x *IsArchiveConfiguredRequest) GetCompanyIds() []uint32 {
	if x != nil {
		return x.CompanyIds
	}
	return nil
}

type IsArchiveConfiguredResponse struct {
	state                 protoimpl.MessageState  `protogen:"open.v1"`
	IsConfigured          bool                    `protobuf:"varint,1,opt,name=is_configured,json=isConfigured,proto3" json:"is_configured,omitempty"`
	CompanyArchiveConfigs []*CompanyArchiveConfig `protobuf:"bytes,2,rep,name=company_archive_configs,json=companyArchiveConfigs,proto3" json:"company_archive_configs,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *IsArchiveConfiguredResponse) Reset() {
	*x = IsArchiveConfiguredResponse{}
	mi := &file_com_coralogix_archive_v1_target_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IsArchiveConfiguredResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsArchiveConfiguredResponse) ProtoMessage() {}

func (x *IsArchiveConfiguredResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v1_target_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsArchiveConfiguredResponse.ProtoReflect.Descriptor instead.
func (*IsArchiveConfiguredResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v1_target_service_proto_rawDescGZIP(), []int{1}
}

func (x *IsArchiveConfiguredResponse) GetIsConfigured() bool {
	if x != nil {
		return x.IsConfigured
	}
	return false
}

func (x *IsArchiveConfiguredResponse) GetCompanyArchiveConfigs() []*CompanyArchiveConfig {
	if x != nil {
		return x.CompanyArchiveConfigs
	}
	return nil
}

type InternalTargetServiceIsArchiveConfiguredRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CompanyIds    []uint32               `protobuf:"varint,1,rep,packed,name=company_ids,json=companyIds,proto3" json:"company_ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InternalTargetServiceIsArchiveConfiguredRequest) Reset() {
	*x = InternalTargetServiceIsArchiveConfiguredRequest{}
	mi := &file_com_coralogix_archive_v1_target_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InternalTargetServiceIsArchiveConfiguredRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalTargetServiceIsArchiveConfiguredRequest) ProtoMessage() {}

func (x *InternalTargetServiceIsArchiveConfiguredRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v1_target_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalTargetServiceIsArchiveConfiguredRequest.ProtoReflect.Descriptor instead.
func (*InternalTargetServiceIsArchiveConfiguredRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v1_target_service_proto_rawDescGZIP(), []int{2}
}

func (x *InternalTargetServiceIsArchiveConfiguredRequest) GetCompanyIds() []uint32 {
	if x != nil {
		return x.CompanyIds
	}
	return nil
}

type InternalTargetServiceIsArchiveConfiguredResponse struct {
	state                 protoimpl.MessageState  `protogen:"open.v1"`
	CompanyArchiveConfigs []*CompanyArchiveConfig `protobuf:"bytes,1,rep,name=company_archive_configs,json=companyArchiveConfigs,proto3" json:"company_archive_configs,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *InternalTargetServiceIsArchiveConfiguredResponse) Reset() {
	*x = InternalTargetServiceIsArchiveConfiguredResponse{}
	mi := &file_com_coralogix_archive_v1_target_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InternalTargetServiceIsArchiveConfiguredResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalTargetServiceIsArchiveConfiguredResponse) ProtoMessage() {}

func (x *InternalTargetServiceIsArchiveConfiguredResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v1_target_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalTargetServiceIsArchiveConfiguredResponse.ProtoReflect.Descriptor instead.
func (*InternalTargetServiceIsArchiveConfiguredResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v1_target_service_proto_rawDescGZIP(), []int{3}
}

func (x *InternalTargetServiceIsArchiveConfiguredResponse) GetCompanyArchiveConfigs() []*CompanyArchiveConfig {
	if x != nil {
		return x.CompanyArchiveConfigs
	}
	return nil
}

type CompanyArchiveConfig struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	CompanyId         uint32                 `protobuf:"varint,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	ArchiveConfigured bool                   `protobuf:"varint,2,opt,name=archive_configured,json=archiveConfigured,proto3" json:"archive_configured,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *CompanyArchiveConfig) Reset() {
	*x = CompanyArchiveConfig{}
	mi := &file_com_coralogix_archive_v1_target_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompanyArchiveConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyArchiveConfig) ProtoMessage() {}

func (x *CompanyArchiveConfig) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v1_target_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyArchiveConfig.ProtoReflect.Descriptor instead.
func (*CompanyArchiveConfig) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v1_target_service_proto_rawDescGZIP(), []int{4}
}

func (x *CompanyArchiveConfig) GetCompanyId() uint32 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *CompanyArchiveConfig) GetArchiveConfigured() bool {
	if x != nil {
		return x.ArchiveConfigured
	}
	return false
}

type GetTargetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTargetRequest) Reset() {
	*x = GetTargetRequest{}
	mi := &file_com_coralogix_archive_v1_target_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTargetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTargetRequest) ProtoMessage() {}

func (x *GetTargetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v1_target_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTargetRequest.ProtoReflect.Descriptor instead.
func (*GetTargetRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v1_target_service_proto_rawDescGZIP(), []int{5}
}

type GetTargetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Target        *Target                `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTargetResponse) Reset() {
	*x = GetTargetResponse{}
	mi := &file_com_coralogix_archive_v1_target_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTargetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTargetResponse) ProtoMessage() {}

func (x *GetTargetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v1_target_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTargetResponse.ProtoReflect.Descriptor instead.
func (*GetTargetResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v1_target_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetTargetResponse) GetTarget() *Target {
	if x != nil {
		return x.Target
	}
	return nil
}

type SetTargetRequest struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Bucket        *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	IsActive      *wrapperspb.BoolValue   `protobuf:"bytes,2,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetTargetRequest) Reset() {
	*x = SetTargetRequest{}
	mi := &file_com_coralogix_archive_v1_target_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetTargetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTargetRequest) ProtoMessage() {}

func (x *SetTargetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v1_target_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTargetRequest.ProtoReflect.Descriptor instead.
func (*SetTargetRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v1_target_service_proto_rawDescGZIP(), []int{7}
}

func (x *SetTargetRequest) GetBucket() *wrapperspb.StringValue {
	if x != nil {
		return x.Bucket
	}
	return nil
}

func (x *SetTargetRequest) GetIsActive() *wrapperspb.BoolValue {
	if x != nil {
		return x.IsActive
	}
	return nil
}

type SetTargetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Target        *Target                `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetTargetResponse) Reset() {
	*x = SetTargetResponse{}
	mi := &file_com_coralogix_archive_v1_target_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetTargetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTargetResponse) ProtoMessage() {}

func (x *SetTargetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v1_target_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTargetResponse.ProtoReflect.Descriptor instead.
func (*SetTargetResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v1_target_service_proto_rawDescGZIP(), []int{8}
}

func (x *SetTargetResponse) GetTarget() *Target {
	if x != nil {
		return x.Target
	}
	return nil
}

var File_com_coralogix_archive_v1_target_service_proto protoreflect.FileDescriptor

var file_com_coralogix_archive_v1_target_service_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x25, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x1a, 0x49, 0x73, 0x41,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x73, 0x22, 0xaa, 0x01, 0x0a, 0x1b, 0x49, 0x73, 0x41,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x12, 0x66, 0x0a,
	0x17, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x15,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x73, 0x22, 0x52, 0x0a, 0x2f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x73,
	0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x73, 0x22, 0x9a, 0x01, 0x0a, 0x30, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x73, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66,
	0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x15, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x22, 0x64, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x2d, 0x0a,
	0x12, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x22, 0x12, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x4d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22,
	0x81, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x37, 0x0a, 0x09, 0x69, 0x73,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x22, 0x4d, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x32, 0xa1, 0x03, 0x0a, 0x0d, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x9f, 0x01, 0x0a, 0x13, 0x49, 0x73, 0x41, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x12, 0x34, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73,
	0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0xc2, 0xb8, 0x02, 0x17, 0x0a,
	0x15, 0x49, 0x73, 0x20, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x20, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x12, 0x76, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x12, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0xc2, 0xb8,
	0x02, 0x0c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x20, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x76,
	0x0a, 0x09, 0x53, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x2a, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0xc2, 0xb8, 0x02, 0x0c, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x20,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x32, 0xe3, 0x01, 0x0a, 0x15, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0xc9, 0x01, 0x0a, 0x13, 0x49, 0x73, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x12, 0x49, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x73, 0x41, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x4a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x73, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1b, 0xc2, 0xb8, 0x02, 0x17, 0x0a, 0x15, 0x49, 0x73, 0x20, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x20, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_archive_v1_target_service_proto_rawDescOnce sync.Once
	file_com_coralogix_archive_v1_target_service_proto_rawDescData = file_com_coralogix_archive_v1_target_service_proto_rawDesc
)

func file_com_coralogix_archive_v1_target_service_proto_rawDescGZIP() []byte {
	file_com_coralogix_archive_v1_target_service_proto_rawDescOnce.Do(func() {
		file_com_coralogix_archive_v1_target_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_archive_v1_target_service_proto_rawDescData)
	})
	return file_com_coralogix_archive_v1_target_service_proto_rawDescData
}

var file_com_coralogix_archive_v1_target_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_com_coralogix_archive_v1_target_service_proto_goTypes = []any{
	(*IsArchiveConfiguredRequest)(nil),                       // 0: com.coralogix.archive.v1.IsArchiveConfiguredRequest
	(*IsArchiveConfiguredResponse)(nil),                      // 1: com.coralogix.archive.v1.IsArchiveConfiguredResponse
	(*InternalTargetServiceIsArchiveConfiguredRequest)(nil),  // 2: com.coralogix.archive.v1.InternalTargetServiceIsArchiveConfiguredRequest
	(*InternalTargetServiceIsArchiveConfiguredResponse)(nil), // 3: com.coralogix.archive.v1.InternalTargetServiceIsArchiveConfiguredResponse
	(*CompanyArchiveConfig)(nil),                             // 4: com.coralogix.archive.v1.CompanyArchiveConfig
	(*GetTargetRequest)(nil),                                 // 5: com.coralogix.archive.v1.GetTargetRequest
	(*GetTargetResponse)(nil),                                // 6: com.coralogix.archive.v1.GetTargetResponse
	(*SetTargetRequest)(nil),                                 // 7: com.coralogix.archive.v1.SetTargetRequest
	(*SetTargetResponse)(nil),                                // 8: com.coralogix.archive.v1.SetTargetResponse
	(*Target)(nil),                                           // 9: com.coralogix.archive.v1.Target
	(*wrapperspb.StringValue)(nil),                           // 10: google.protobuf.StringValue
	(*wrapperspb.BoolValue)(nil),                             // 11: google.protobuf.BoolValue
}
var file_com_coralogix_archive_v1_target_service_proto_depIdxs = []int32{
	4,  // 0: com.coralogix.archive.v1.IsArchiveConfiguredResponse.company_archive_configs:type_name -> com.coralogix.archive.v1.CompanyArchiveConfig
	4,  // 1: com.coralogix.archive.v1.InternalTargetServiceIsArchiveConfiguredResponse.company_archive_configs:type_name -> com.coralogix.archive.v1.CompanyArchiveConfig
	9,  // 2: com.coralogix.archive.v1.GetTargetResponse.target:type_name -> com.coralogix.archive.v1.Target
	10, // 3: com.coralogix.archive.v1.SetTargetRequest.bucket:type_name -> google.protobuf.StringValue
	11, // 4: com.coralogix.archive.v1.SetTargetRequest.is_active:type_name -> google.protobuf.BoolValue
	9,  // 5: com.coralogix.archive.v1.SetTargetResponse.target:type_name -> com.coralogix.archive.v1.Target
	0,  // 6: com.coralogix.archive.v1.TargetService.IsArchiveConfigured:input_type -> com.coralogix.archive.v1.IsArchiveConfiguredRequest
	5,  // 7: com.coralogix.archive.v1.TargetService.GetTarget:input_type -> com.coralogix.archive.v1.GetTargetRequest
	7,  // 8: com.coralogix.archive.v1.TargetService.SetTarget:input_type -> com.coralogix.archive.v1.SetTargetRequest
	2,  // 9: com.coralogix.archive.v1.InternalTargetService.IsArchiveConfigured:input_type -> com.coralogix.archive.v1.InternalTargetServiceIsArchiveConfiguredRequest
	1,  // 10: com.coralogix.archive.v1.TargetService.IsArchiveConfigured:output_type -> com.coralogix.archive.v1.IsArchiveConfiguredResponse
	6,  // 11: com.coralogix.archive.v1.TargetService.GetTarget:output_type -> com.coralogix.archive.v1.GetTargetResponse
	8,  // 12: com.coralogix.archive.v1.TargetService.SetTarget:output_type -> com.coralogix.archive.v1.SetTargetResponse
	3,  // 13: com.coralogix.archive.v1.InternalTargetService.IsArchiveConfigured:output_type -> com.coralogix.archive.v1.InternalTargetServiceIsArchiveConfiguredResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_com_coralogix_archive_v1_target_service_proto_init() }
func file_com_coralogix_archive_v1_target_service_proto_init() {
	if File_com_coralogix_archive_v1_target_service_proto != nil {
		return
	}
	file_com_coralogix_archive_v1_target_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_archive_v1_target_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_com_coralogix_archive_v1_target_service_proto_goTypes,
		DependencyIndexes: file_com_coralogix_archive_v1_target_service_proto_depIdxs,
		MessageInfos:      file_com_coralogix_archive_v1_target_service_proto_msgTypes,
	}.Build()
	File_com_coralogix_archive_v1_target_service_proto = out.File
	file_com_coralogix_archive_v1_target_service_proto_rawDesc = nil
	file_com_coralogix_archive_v1_target_service_proto_goTypes = nil
	file_com_coralogix_archive_v1_target_service_proto_depIdxs = nil
}
