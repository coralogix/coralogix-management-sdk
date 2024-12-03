// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: com/coralogix/archive/v2/target_service.proto

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

type GetTargetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTargetRequest) Reset() {
	*x = GetTargetRequest{}
	mi := &file_com_coralogix_archive_v2_target_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTargetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTargetRequest) ProtoMessage() {}

func (x *GetTargetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_target_service_proto_msgTypes[0]
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
	return file_com_coralogix_archive_v2_target_service_proto_rawDescGZIP(), []int{0}
}

type GetTargetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target *Target `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *GetTargetResponse) Reset() {
	*x = GetTargetResponse{}
	mi := &file_com_coralogix_archive_v2_target_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTargetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTargetResponse) ProtoMessage() {}

func (x *GetTargetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_target_service_proto_msgTypes[1]
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
	return file_com_coralogix_archive_v2_target_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetTargetResponse) GetTarget() *Target {
	if x != nil {
		return x.Target
	}
	return nil
}

type SetTargetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsActive bool `protobuf:"varint,1,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	// Types that are assignable to TargetSpec:
	//
	//	*SetTargetRequest_S3
	//	*SetTargetRequest_IbmCos
	TargetSpec isSetTargetRequest_TargetSpec `protobuf_oneof:"target_spec"`
}

func (x *SetTargetRequest) Reset() {
	*x = SetTargetRequest{}
	mi := &file_com_coralogix_archive_v2_target_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetTargetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTargetRequest) ProtoMessage() {}

func (x *SetTargetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_target_service_proto_msgTypes[2]
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
	return file_com_coralogix_archive_v2_target_service_proto_rawDescGZIP(), []int{2}
}

func (x *SetTargetRequest) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (m *SetTargetRequest) GetTargetSpec() isSetTargetRequest_TargetSpec {
	if m != nil {
		return m.TargetSpec
	}
	return nil
}

func (x *SetTargetRequest) GetS3() *S3TargetSpec {
	if x, ok := x.GetTargetSpec().(*SetTargetRequest_S3); ok {
		return x.S3
	}
	return nil
}

func (x *SetTargetRequest) GetIbmCos() *IBMCosTargetSpec {
	if x, ok := x.GetTargetSpec().(*SetTargetRequest_IbmCos); ok {
		return x.IbmCos
	}
	return nil
}

type isSetTargetRequest_TargetSpec interface {
	isSetTargetRequest_TargetSpec()
}

type SetTargetRequest_S3 struct {
	S3 *S3TargetSpec `protobuf:"bytes,2,opt,name=s3,proto3,oneof"`
}

type SetTargetRequest_IbmCos struct {
	IbmCos *IBMCosTargetSpec `protobuf:"bytes,3,opt,name=ibm_cos,json=ibmCos,proto3,oneof"`
}

func (*SetTargetRequest_S3) isSetTargetRequest_TargetSpec() {}

func (*SetTargetRequest_IbmCos) isSetTargetRequest_TargetSpec() {}

type SetTargetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target *Target `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *SetTargetResponse) Reset() {
	*x = SetTargetResponse{}
	mi := &file_com_coralogix_archive_v2_target_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetTargetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTargetResponse) ProtoMessage() {}

func (x *SetTargetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_target_service_proto_msgTypes[3]
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
	return file_com_coralogix_archive_v2_target_service_proto_rawDescGZIP(), []int{3}
}

func (x *SetTargetResponse) GetTarget() *Target {
	if x != nil {
		return x.Target
	}
	return nil
}

type SetExternalTargetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsActive bool `protobuf:"varint,1,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	// Types that are assignable to TargetSpec:
	//
	//	*SetExternalTargetRequest_S3
	//	*SetExternalTargetRequest_IbmCos
	TargetSpec isSetExternalTargetRequest_TargetSpec `protobuf_oneof:"target_spec"`
	CompanyId  uint32                                `protobuf:"varint,4,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
}

func (x *SetExternalTargetRequest) Reset() {
	*x = SetExternalTargetRequest{}
	mi := &file_com_coralogix_archive_v2_target_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetExternalTargetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetExternalTargetRequest) ProtoMessage() {}

func (x *SetExternalTargetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_target_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetExternalTargetRequest.ProtoReflect.Descriptor instead.
func (*SetExternalTargetRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_target_service_proto_rawDescGZIP(), []int{4}
}

func (x *SetExternalTargetRequest) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (m *SetExternalTargetRequest) GetTargetSpec() isSetExternalTargetRequest_TargetSpec {
	if m != nil {
		return m.TargetSpec
	}
	return nil
}

func (x *SetExternalTargetRequest) GetS3() *S3TargetSpec {
	if x, ok := x.GetTargetSpec().(*SetExternalTargetRequest_S3); ok {
		return x.S3
	}
	return nil
}

func (x *SetExternalTargetRequest) GetIbmCos() *IBMCosTargetSpec {
	if x, ok := x.GetTargetSpec().(*SetExternalTargetRequest_IbmCos); ok {
		return x.IbmCos
	}
	return nil
}

func (x *SetExternalTargetRequest) GetCompanyId() uint32 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

type isSetExternalTargetRequest_TargetSpec interface {
	isSetExternalTargetRequest_TargetSpec()
}

type SetExternalTargetRequest_S3 struct {
	S3 *S3TargetSpec `protobuf:"bytes,2,opt,name=s3,proto3,oneof"`
}

type SetExternalTargetRequest_IbmCos struct {
	IbmCos *IBMCosTargetSpec `protobuf:"bytes,3,opt,name=ibm_cos,json=ibmCos,proto3,oneof"`
}

func (*SetExternalTargetRequest_S3) isSetExternalTargetRequest_TargetSpec() {}

func (*SetExternalTargetRequest_IbmCos) isSetExternalTargetRequest_TargetSpec() {}

type SetExternalTargetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target *Target `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *SetExternalTargetResponse) Reset() {
	*x = SetExternalTargetResponse{}
	mi := &file_com_coralogix_archive_v2_target_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetExternalTargetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetExternalTargetResponse) ProtoMessage() {}

func (x *SetExternalTargetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_target_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetExternalTargetResponse.ProtoReflect.Descriptor instead.
func (*SetExternalTargetResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_target_service_proto_rawDescGZIP(), []int{5}
}

func (x *SetExternalTargetResponse) GetTarget() *Target {
	if x != nil {
		return x.Target
	}
	return nil
}

type CompanyArchiveConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId         uint32 `protobuf:"varint,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	ArchiveConfigured bool   `protobuf:"varint,2,opt,name=archive_configured,json=archiveConfigured,proto3" json:"archive_configured,omitempty"`
}

func (x *CompanyArchiveConfig) Reset() {
	*x = CompanyArchiveConfig{}
	mi := &file_com_coralogix_archive_v2_target_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompanyArchiveConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyArchiveConfig) ProtoMessage() {}

func (x *CompanyArchiveConfig) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_target_service_proto_msgTypes[6]
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
	return file_com_coralogix_archive_v2_target_service_proto_rawDescGZIP(), []int{6}
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

type ValidateTargetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsActive bool `protobuf:"varint,1,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	// Types that are assignable to TargetSpec:
	//
	//	*ValidateTargetRequest_S3
	//	*ValidateTargetRequest_IbmCos
	TargetSpec isValidateTargetRequest_TargetSpec `protobuf_oneof:"target_spec"`
}

func (x *ValidateTargetRequest) Reset() {
	*x = ValidateTargetRequest{}
	mi := &file_com_coralogix_archive_v2_target_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateTargetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateTargetRequest) ProtoMessage() {}

func (x *ValidateTargetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_target_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateTargetRequest.ProtoReflect.Descriptor instead.
func (*ValidateTargetRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_target_service_proto_rawDescGZIP(), []int{7}
}

func (x *ValidateTargetRequest) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (m *ValidateTargetRequest) GetTargetSpec() isValidateTargetRequest_TargetSpec {
	if m != nil {
		return m.TargetSpec
	}
	return nil
}

func (x *ValidateTargetRequest) GetS3() *S3TargetSpec {
	if x, ok := x.GetTargetSpec().(*ValidateTargetRequest_S3); ok {
		return x.S3
	}
	return nil
}

func (x *ValidateTargetRequest) GetIbmCos() *IBMCosTargetSpec {
	if x, ok := x.GetTargetSpec().(*ValidateTargetRequest_IbmCos); ok {
		return x.IbmCos
	}
	return nil
}

type isValidateTargetRequest_TargetSpec interface {
	isValidateTargetRequest_TargetSpec()
}

type ValidateTargetRequest_S3 struct {
	S3 *S3TargetSpec `protobuf:"bytes,2,opt,name=s3,proto3,oneof"`
}

type ValidateTargetRequest_IbmCos struct {
	IbmCos *IBMCosTargetSpec `protobuf:"bytes,3,opt,name=ibm_cos,json=ibmCos,proto3,oneof"`
}

func (*ValidateTargetRequest_S3) isValidateTargetRequest_TargetSpec() {}

func (*ValidateTargetRequest_IbmCos) isValidateTargetRequest_TargetSpec() {}

type ValidateTargetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsValid bool `protobuf:"varint,1,opt,name=is_valid,json=isValid,proto3" json:"is_valid,omitempty"`
}

func (x *ValidateTargetResponse) Reset() {
	*x = ValidateTargetResponse{}
	mi := &file_com_coralogix_archive_v2_target_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateTargetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateTargetResponse) ProtoMessage() {}

func (x *ValidateTargetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_target_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateTargetResponse.ProtoReflect.Descriptor instead.
func (*ValidateTargetResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_target_service_proto_rawDescGZIP(), []int{8}
}

func (x *ValidateTargetResponse) GetIsValid() bool {
	if x != nil {
		return x.IsValid
	}
	return false
}

var File_com_coralogix_archive_v2_target_service_proto protoreflect.FileDescriptor

var file_com_coralogix_archive_v2_target_service_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x1a, 0x28, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x67, 0x65,
	0x74, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x4d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x2e, 0x76, 0x32, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x22, 0xbf, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x12, 0x38, 0x0a, 0x02, 0x73, 0x33, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x33, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x48, 0x00, 0x52, 0x02, 0x73, 0x33, 0x12, 0x45,
	0x0a, 0x07, 0x69, 0x62, 0x6d, 0x5f, 0x63, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x42, 0x4d, 0x43, 0x6f,
	0x73, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x48, 0x00, 0x52, 0x06, 0x69,
	0x62, 0x6d, 0x43, 0x6f, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x22, 0x4d, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x22, 0xe6, 0x01, 0x0a, 0x18, 0x53, 0x65, 0x74, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x38, 0x0a,
	0x02, 0x73, 0x33, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x33, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x70, 0x65,
	0x63, 0x48, 0x00, 0x52, 0x02, 0x73, 0x33, 0x12, 0x45, 0x0a, 0x07, 0x69, 0x62, 0x6d, 0x5f, 0x63,
	0x6f, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x2e, 0x76, 0x32, 0x2e, 0x49, 0x42, 0x4d, 0x43, 0x6f, 0x73, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x53, 0x70, 0x65, 0x63, 0x48, 0x00, 0x52, 0x06, 0x69, 0x62, 0x6d, 0x43, 0x6f, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x42, 0x0d, 0x0a,
	0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x22, 0x55, 0x0a, 0x19,
	0x53, 0x65, 0x74, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x22, 0x64, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x41, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x12, 0x61, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x22, 0xc4, 0x01, 0x0a, 0x15, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x38, 0x0a, 0x02, 0x73, 0x33, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x33, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x53, 0x70, 0x65, 0x63, 0x48, 0x00, 0x52, 0x02, 0x73, 0x33, 0x12, 0x45, 0x0a, 0x07, 0x69, 0x62,
	0x6d, 0x5f, 0x63, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x42, 0x4d, 0x43, 0x6f, 0x73, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x48, 0x00, 0x52, 0x06, 0x69, 0x62, 0x6d, 0x43, 0x6f,
	0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x63,
	0x22, 0x33, 0x0a, 0x16, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73,
	0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x32, 0x8c, 0x03, 0x0a, 0x0d, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x76, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x12, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0xc2,
	0xb8, 0x02, 0x0c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x20, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12,
	0x76, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x2a, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0xc2, 0xb8, 0x02, 0x0c, 0x0a, 0x0a, 0x53, 0x65, 0x74,
	0x20, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x8a, 0x01, 0x0a, 0x0e, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x2f, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0xc2,
	0xb8, 0x02, 0x11, 0x0a, 0x0f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x20, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x32, 0xf5, 0x02, 0x0a, 0x15, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa4,
	0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x40, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65,
	0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x41, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x11, 0xc2, 0xb8, 0x02, 0x0d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x20, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0xb4, 0x01, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x32, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x74, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x74, 0x45, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0xc2, 0xb8, 0x02, 0x32, 0x0a, 0x30, 0x53, 0x65, 0x74, 0x20,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x20, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x20, 0x66, 0x6f, 0x72, 0x20,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x75, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_archive_v2_target_service_proto_rawDescOnce sync.Once
	file_com_coralogix_archive_v2_target_service_proto_rawDescData = file_com_coralogix_archive_v2_target_service_proto_rawDesc
)

func file_com_coralogix_archive_v2_target_service_proto_rawDescGZIP() []byte {
	file_com_coralogix_archive_v2_target_service_proto_rawDescOnce.Do(func() {
		file_com_coralogix_archive_v2_target_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_archive_v2_target_service_proto_rawDescData)
	})
	return file_com_coralogix_archive_v2_target_service_proto_rawDescData
}

var file_com_coralogix_archive_v2_target_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_com_coralogix_archive_v2_target_service_proto_goTypes = []any{
	(*GetTargetRequest)(nil),                        // 0: com.coralogix.archive.v2.GetTargetRequest
	(*GetTargetResponse)(nil),                       // 1: com.coralogix.archive.v2.GetTargetResponse
	(*SetTargetRequest)(nil),                        // 2: com.coralogix.archive.v2.SetTargetRequest
	(*SetTargetResponse)(nil),                       // 3: com.coralogix.archive.v2.SetTargetResponse
	(*SetExternalTargetRequest)(nil),                // 4: com.coralogix.archive.v2.SetExternalTargetRequest
	(*SetExternalTargetResponse)(nil),               // 5: com.coralogix.archive.v2.SetExternalTargetResponse
	(*CompanyArchiveConfig)(nil),                    // 6: com.coralogix.archive.v2.CompanyArchiveConfig
	(*ValidateTargetRequest)(nil),                   // 7: com.coralogix.archive.v2.ValidateTargetRequest
	(*ValidateTargetResponse)(nil),                  // 8: com.coralogix.archive.v2.ValidateTargetResponse
	(*Target)(nil),                                  // 9: com.coralogix.archive.v2.Target
	(*S3TargetSpec)(nil),                            // 10: com.coralogix.archive.v2.S3TargetSpec
	(*IBMCosTargetSpec)(nil),                        // 11: com.coralogix.archive.v2.IBMCosTargetSpec
	(*InternalTargetServiceGetTargetsRequest)(nil),  // 12: com.coralogix.archive.v2.InternalTargetServiceGetTargetsRequest
	(*InternalTargetServiceGetTargetsResponse)(nil), // 13: com.coralogix.archive.v2.InternalTargetServiceGetTargetsResponse
}
var file_com_coralogix_archive_v2_target_service_proto_depIdxs = []int32{
	9,  // 0: com.coralogix.archive.v2.GetTargetResponse.target:type_name -> com.coralogix.archive.v2.Target
	10, // 1: com.coralogix.archive.v2.SetTargetRequest.s3:type_name -> com.coralogix.archive.v2.S3TargetSpec
	11, // 2: com.coralogix.archive.v2.SetTargetRequest.ibm_cos:type_name -> com.coralogix.archive.v2.IBMCosTargetSpec
	9,  // 3: com.coralogix.archive.v2.SetTargetResponse.target:type_name -> com.coralogix.archive.v2.Target
	10, // 4: com.coralogix.archive.v2.SetExternalTargetRequest.s3:type_name -> com.coralogix.archive.v2.S3TargetSpec
	11, // 5: com.coralogix.archive.v2.SetExternalTargetRequest.ibm_cos:type_name -> com.coralogix.archive.v2.IBMCosTargetSpec
	9,  // 6: com.coralogix.archive.v2.SetExternalTargetResponse.target:type_name -> com.coralogix.archive.v2.Target
	10, // 7: com.coralogix.archive.v2.ValidateTargetRequest.s3:type_name -> com.coralogix.archive.v2.S3TargetSpec
	11, // 8: com.coralogix.archive.v2.ValidateTargetRequest.ibm_cos:type_name -> com.coralogix.archive.v2.IBMCosTargetSpec
	0,  // 9: com.coralogix.archive.v2.TargetService.GetTarget:input_type -> com.coralogix.archive.v2.GetTargetRequest
	2,  // 10: com.coralogix.archive.v2.TargetService.SetTarget:input_type -> com.coralogix.archive.v2.SetTargetRequest
	7,  // 11: com.coralogix.archive.v2.TargetService.ValidateTarget:input_type -> com.coralogix.archive.v2.ValidateTargetRequest
	12, // 12: com.coralogix.archive.v2.InternalTargetService.GetTargets:input_type -> com.coralogix.archive.v2.InternalTargetServiceGetTargetsRequest
	4,  // 13: com.coralogix.archive.v2.InternalTargetService.SetExternalTarget:input_type -> com.coralogix.archive.v2.SetExternalTargetRequest
	1,  // 14: com.coralogix.archive.v2.TargetService.GetTarget:output_type -> com.coralogix.archive.v2.GetTargetResponse
	3,  // 15: com.coralogix.archive.v2.TargetService.SetTarget:output_type -> com.coralogix.archive.v2.SetTargetResponse
	8,  // 16: com.coralogix.archive.v2.TargetService.ValidateTarget:output_type -> com.coralogix.archive.v2.ValidateTargetResponse
	13, // 17: com.coralogix.archive.v2.InternalTargetService.GetTargets:output_type -> com.coralogix.archive.v2.InternalTargetServiceGetTargetsResponse
	5,  // 18: com.coralogix.archive.v2.InternalTargetService.SetExternalTarget:output_type -> com.coralogix.archive.v2.SetExternalTargetResponse
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_com_coralogix_archive_v2_target_service_proto_init() }
func file_com_coralogix_archive_v2_target_service_proto_init() {
	if File_com_coralogix_archive_v2_target_service_proto != nil {
		return
	}
	file_com_coralogix_archive_v2_target_proto_init()
	file_com_coralogix_archive_v2_internal_get_targets_proto_init()
	file_com_coralogix_archive_v2_target_service_proto_msgTypes[2].OneofWrappers = []any{
		(*SetTargetRequest_S3)(nil),
		(*SetTargetRequest_IbmCos)(nil),
	}
	file_com_coralogix_archive_v2_target_service_proto_msgTypes[4].OneofWrappers = []any{
		(*SetExternalTargetRequest_S3)(nil),
		(*SetExternalTargetRequest_IbmCos)(nil),
	}
	file_com_coralogix_archive_v2_target_service_proto_msgTypes[7].OneofWrappers = []any{
		(*ValidateTargetRequest_S3)(nil),
		(*ValidateTargetRequest_IbmCos)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_archive_v2_target_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_com_coralogix_archive_v2_target_service_proto_goTypes,
		DependencyIndexes: file_com_coralogix_archive_v2_target_service_proto_depIdxs,
		MessageInfos:      file_com_coralogix_archive_v2_target_service_proto_msgTypes,
	}.Build()
	File_com_coralogix_archive_v2_target_service_proto = out.File
	file_com_coralogix_archive_v2_target_service_proto_rawDesc = nil
	file_com_coralogix_archive_v2_target_service_proto_goTypes = nil
	file_com_coralogix_archive_v2_target_service_proto_depIdxs = nil
}
