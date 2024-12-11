// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.1
// source: com/coralogix/enrichment/v1/enrichment_service.proto

package v1

import (
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

type GetEnrichmentLimitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetEnrichmentLimitRequest) Reset() {
	*x = GetEnrichmentLimitRequest{}
	mi := &file_com_coralogix_enrichment_v1_enrichment_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEnrichmentLimitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEnrichmentLimitRequest) ProtoMessage() {}

func (x *GetEnrichmentLimitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEnrichmentLimitRequest.ProtoReflect.Descriptor instead.
func (*GetEnrichmentLimitRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_service_proto_rawDescGZIP(), []int{0}
}

type GetEnrichmentLimitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit uint32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Used  uint32 `protobuf:"varint,2,opt,name=used,proto3" json:"used,omitempty"`
}

func (x *GetEnrichmentLimitResponse) Reset() {
	*x = GetEnrichmentLimitResponse{}
	mi := &file_com_coralogix_enrichment_v1_enrichment_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEnrichmentLimitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEnrichmentLimitResponse) ProtoMessage() {}

func (x *GetEnrichmentLimitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEnrichmentLimitResponse.ProtoReflect.Descriptor instead.
func (*GetEnrichmentLimitResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetEnrichmentLimitResponse) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetEnrichmentLimitResponse) GetUsed() uint32 {
	if x != nil {
		return x.Used
	}
	return 0
}

type GetEnrichmentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetEnrichmentsRequest) Reset() {
	*x = GetEnrichmentsRequest{}
	mi := &file_com_coralogix_enrichment_v1_enrichment_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEnrichmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEnrichmentsRequest) ProtoMessage() {}

func (x *GetEnrichmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEnrichmentsRequest.ProtoReflect.Descriptor instead.
func (*GetEnrichmentsRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_service_proto_rawDescGZIP(), []int{2}
}

type GetEnrichmentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enrichments []*Enrichment `protobuf:"bytes,1,rep,name=enrichments,proto3" json:"enrichments,omitempty"`
}

func (x *GetEnrichmentsResponse) Reset() {
	*x = GetEnrichmentsResponse{}
	mi := &file_com_coralogix_enrichment_v1_enrichment_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEnrichmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEnrichmentsResponse) ProtoMessage() {}

func (x *GetEnrichmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEnrichmentsResponse.ProtoReflect.Descriptor instead.
func (*GetEnrichmentsResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetEnrichmentsResponse) GetEnrichments() []*Enrichment {
	if x != nil {
		return x.Enrichments
	}
	return nil
}

type AddEnrichmentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestEnrichments []*EnrichmentRequestModel `protobuf:"bytes,1,rep,name=request_enrichments,json=requestEnrichments,proto3" json:"request_enrichments,omitempty"`
}

func (x *AddEnrichmentsRequest) Reset() {
	*x = AddEnrichmentsRequest{}
	mi := &file_com_coralogix_enrichment_v1_enrichment_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddEnrichmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEnrichmentsRequest) ProtoMessage() {}

func (x *AddEnrichmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEnrichmentsRequest.ProtoReflect.Descriptor instead.
func (*AddEnrichmentsRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_service_proto_rawDescGZIP(), []int{4}
}

func (x *AddEnrichmentsRequest) GetRequestEnrichments() []*EnrichmentRequestModel {
	if x != nil {
		return x.RequestEnrichments
	}
	return nil
}

type AddEnrichmentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enrichments []*Enrichment `protobuf:"bytes,1,rep,name=enrichments,proto3" json:"enrichments,omitempty"`
}

func (x *AddEnrichmentsResponse) Reset() {
	*x = AddEnrichmentsResponse{}
	mi := &file_com_coralogix_enrichment_v1_enrichment_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddEnrichmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEnrichmentsResponse) ProtoMessage() {}

func (x *AddEnrichmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEnrichmentsResponse.ProtoReflect.Descriptor instead.
func (*AddEnrichmentsResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_service_proto_rawDescGZIP(), []int{5}
}

func (x *AddEnrichmentsResponse) GetEnrichments() []*Enrichment {
	if x != nil {
		return x.Enrichments
	}
	return nil
}

type RemoveEnrichmentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnrichmentIds []*wrapperspb.UInt32Value `protobuf:"bytes,1,rep,name=enrichment_ids,json=enrichmentIds,proto3" json:"enrichment_ids,omitempty"`
}

func (x *RemoveEnrichmentsRequest) Reset() {
	*x = RemoveEnrichmentsRequest{}
	mi := &file_com_coralogix_enrichment_v1_enrichment_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveEnrichmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveEnrichmentsRequest) ProtoMessage() {}

func (x *RemoveEnrichmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveEnrichmentsRequest.ProtoReflect.Descriptor instead.
func (*RemoveEnrichmentsRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_service_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveEnrichmentsRequest) GetEnrichmentIds() []*wrapperspb.UInt32Value {
	if x != nil {
		return x.EnrichmentIds
	}
	return nil
}

type RemoveEnrichmentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RemainingEnrichments []*Enrichment `protobuf:"bytes,1,rep,name=remaining_enrichments,json=remainingEnrichments,proto3" json:"remaining_enrichments,omitempty"`
}

func (x *RemoveEnrichmentsResponse) Reset() {
	*x = RemoveEnrichmentsResponse{}
	mi := &file_com_coralogix_enrichment_v1_enrichment_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveEnrichmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveEnrichmentsResponse) ProtoMessage() {}

func (x *RemoveEnrichmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveEnrichmentsResponse.ProtoReflect.Descriptor instead.
func (*RemoveEnrichmentsResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_service_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveEnrichmentsResponse) GetRemainingEnrichments() []*Enrichment {
	if x != nil {
		return x.RemainingEnrichments
	}
	return nil
}

// Deletes all enrichments of the given type and replaces them with the new ones
type AtomicOverwriteEnrichmentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnrichmentType   *EnrichmentType              `protobuf:"bytes,1,opt,name=enrichment_type,json=enrichmentType,proto3" json:"enrichment_type,omitempty"`
	EnrichmentFields []*EnrichmentFieldDefinition `protobuf:"bytes,2,rep,name=enrichment_fields,json=enrichmentFields,proto3" json:"enrichment_fields,omitempty"`
}

func (x *AtomicOverwriteEnrichmentsRequest) Reset() {
	*x = AtomicOverwriteEnrichmentsRequest{}
	mi := &file_com_coralogix_enrichment_v1_enrichment_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AtomicOverwriteEnrichmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AtomicOverwriteEnrichmentsRequest) ProtoMessage() {}

func (x *AtomicOverwriteEnrichmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AtomicOverwriteEnrichmentsRequest.ProtoReflect.Descriptor instead.
func (*AtomicOverwriteEnrichmentsRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_service_proto_rawDescGZIP(), []int{8}
}

func (x *AtomicOverwriteEnrichmentsRequest) GetEnrichmentType() *EnrichmentType {
	if x != nil {
		return x.EnrichmentType
	}
	return nil
}

func (x *AtomicOverwriteEnrichmentsRequest) GetEnrichmentFields() []*EnrichmentFieldDefinition {
	if x != nil {
		return x.EnrichmentFields
	}
	return nil
}

type AtomicOverwriteEnrichmentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enrichments []*Enrichment `protobuf:"bytes,1,rep,name=enrichments,proto3" json:"enrichments,omitempty"`
}

func (x *AtomicOverwriteEnrichmentsResponse) Reset() {
	*x = AtomicOverwriteEnrichmentsResponse{}
	mi := &file_com_coralogix_enrichment_v1_enrichment_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AtomicOverwriteEnrichmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AtomicOverwriteEnrichmentsResponse) ProtoMessage() {}

func (x *AtomicOverwriteEnrichmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AtomicOverwriteEnrichmentsResponse.ProtoReflect.Descriptor instead.
func (*AtomicOverwriteEnrichmentsResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_service_proto_rawDescGZIP(), []int{9}
}

func (x *AtomicOverwriteEnrichmentsResponse) GetEnrichments() []*Enrichment {
	if x != nil {
		return x.Enrichments
	}
	return nil
}

type GetCompanyEnrichmentSettingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCompanyEnrichmentSettingsRequest) Reset() {
	*x = GetCompanyEnrichmentSettingsRequest{}
	mi := &file_com_coralogix_enrichment_v1_enrichment_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCompanyEnrichmentSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompanyEnrichmentSettingsRequest) ProtoMessage() {}

func (x *GetCompanyEnrichmentSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCompanyEnrichmentSettingsRequest.ProtoReflect.Descriptor instead.
func (*GetCompanyEnrichmentSettingsRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_service_proto_rawDescGZIP(), []int{10}
}

type GetCompanyEnrichmentSettingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnrichmentSettings *CompanyEnrichmentSettings `protobuf:"bytes,1,opt,name=enrichment_settings,json=enrichmentSettings,proto3" json:"enrichment_settings,omitempty"`
}

func (x *GetCompanyEnrichmentSettingsResponse) Reset() {
	*x = GetCompanyEnrichmentSettingsResponse{}
	mi := &file_com_coralogix_enrichment_v1_enrichment_service_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCompanyEnrichmentSettingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompanyEnrichmentSettingsResponse) ProtoMessage() {}

func (x *GetCompanyEnrichmentSettingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_service_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCompanyEnrichmentSettingsResponse.ProtoReflect.Descriptor instead.
func (*GetCompanyEnrichmentSettingsResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_service_proto_rawDescGZIP(), []int{11}
}

func (x *GetCompanyEnrichmentSettingsResponse) GetEnrichmentSettings() *CompanyEnrichmentSettings {
	if x != nil {
		return x.EnrichmentSettings
	}
	return nil
}

var File_com_coralogix_enrichment_v1_enrichment_service_proto protoreflect.FileDescriptor

var file_com_coralogix_enrichment_v1_enrichment_service_proto_rawDesc = []byte{
	0x0a, 0x34, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e,
	0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2f, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x3a, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2f, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x65, 0x6e, 0x72,
	0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x72, 0x69, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x3d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x1b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x46, 0x0a, 0x1a,
	0x47, 0x65, 0x74, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x72, 0x69, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x63, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0b, 0x65, 0x6e, 0x72, 0x69, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x6e, 0x72,
	0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x72, 0x69, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x7d, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x64, 0x0a, 0x13, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x12, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x22, 0x63, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0b, 0x65,
	0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x65, 0x6e, 0x72, 0x69, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x5f, 0x0a, 0x18, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x43, 0x0a, 0x0e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x22, 0x79, 0x0a, 0x19, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x15, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x5f, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x14, 0x72, 0x65,
	0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0xde, 0x01, 0x0a, 0x21, 0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x4f, 0x76, 0x65,
	0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x54, 0x0a, 0x0f, 0x65, 0x6e, 0x72, 0x69,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e,
	0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x63,
	0x0a, 0x11, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x10, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x22, 0x6f, 0x0a, 0x22, 0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x4f, 0x76, 0x65,
	0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0b, 0x65, 0x6e, 0x72,
	0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65,
	0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x72,
	0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x22, 0x25, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x8f, 0x01, 0x0a, 0x24,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a, 0x13, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x12, 0x65, 0x6e, 0x72, 0x69, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x32, 0xdc, 0x06,
	0x0a, 0x11, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x79, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x72, 0x69, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x79,
	0x0a, 0x0e, 0x41, 0x64, 0x64, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x64, 0x64, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x11, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x45, 0x6e, 0x72, 0x69, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x85,
	0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x6e,
	0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45,
	0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x9d, 0x01, 0x0a, 0x1a, 0x41, 0x74, 0x6f, 0x6d, 0x69,
	0x63, 0x4f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x3e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x4f, 0x76, 0x65, 0x72, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x4f, 0x76, 0x65, 0x72, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0xa3, 0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x40, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x41, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_enrichment_v1_enrichment_service_proto_rawDescOnce sync.Once
	file_com_coralogix_enrichment_v1_enrichment_service_proto_rawDescData = file_com_coralogix_enrichment_v1_enrichment_service_proto_rawDesc
)

func file_com_coralogix_enrichment_v1_enrichment_service_proto_rawDescGZIP() []byte {
	file_com_coralogix_enrichment_v1_enrichment_service_proto_rawDescOnce.Do(func() {
		file_com_coralogix_enrichment_v1_enrichment_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_enrichment_v1_enrichment_service_proto_rawDescData)
	})
	return file_com_coralogix_enrichment_v1_enrichment_service_proto_rawDescData
}

var file_com_coralogix_enrichment_v1_enrichment_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_com_coralogix_enrichment_v1_enrichment_service_proto_goTypes = []any{
	(*GetEnrichmentLimitRequest)(nil),            // 0: com.coralogix.enrichment.v1.GetEnrichmentLimitRequest
	(*GetEnrichmentLimitResponse)(nil),           // 1: com.coralogix.enrichment.v1.GetEnrichmentLimitResponse
	(*GetEnrichmentsRequest)(nil),                // 2: com.coralogix.enrichment.v1.GetEnrichmentsRequest
	(*GetEnrichmentsResponse)(nil),               // 3: com.coralogix.enrichment.v1.GetEnrichmentsResponse
	(*AddEnrichmentsRequest)(nil),                // 4: com.coralogix.enrichment.v1.AddEnrichmentsRequest
	(*AddEnrichmentsResponse)(nil),               // 5: com.coralogix.enrichment.v1.AddEnrichmentsResponse
	(*RemoveEnrichmentsRequest)(nil),             // 6: com.coralogix.enrichment.v1.RemoveEnrichmentsRequest
	(*RemoveEnrichmentsResponse)(nil),            // 7: com.coralogix.enrichment.v1.RemoveEnrichmentsResponse
	(*AtomicOverwriteEnrichmentsRequest)(nil),    // 8: com.coralogix.enrichment.v1.AtomicOverwriteEnrichmentsRequest
	(*AtomicOverwriteEnrichmentsResponse)(nil),   // 9: com.coralogix.enrichment.v1.AtomicOverwriteEnrichmentsResponse
	(*GetCompanyEnrichmentSettingsRequest)(nil),  // 10: com.coralogix.enrichment.v1.GetCompanyEnrichmentSettingsRequest
	(*GetCompanyEnrichmentSettingsResponse)(nil), // 11: com.coralogix.enrichment.v1.GetCompanyEnrichmentSettingsResponse
	(*Enrichment)(nil),                           // 12: com.coralogix.enrichment.v1.Enrichment
	(*EnrichmentRequestModel)(nil),               // 13: com.coralogix.enrichment.v1.EnrichmentRequestModel
	(*wrapperspb.UInt32Value)(nil),               // 14: google.protobuf.UInt32Value
	(*EnrichmentType)(nil),                       // 15: com.coralogix.enrichment.v1.EnrichmentType
	(*EnrichmentFieldDefinition)(nil),            // 16: com.coralogix.enrichment.v1.EnrichmentFieldDefinition
	(*CompanyEnrichmentSettings)(nil),            // 17: com.coralogix.enrichment.v1.CompanyEnrichmentSettings
}
var file_com_coralogix_enrichment_v1_enrichment_service_proto_depIdxs = []int32{
	12, // 0: com.coralogix.enrichment.v1.GetEnrichmentsResponse.enrichments:type_name -> com.coralogix.enrichment.v1.Enrichment
	13, // 1: com.coralogix.enrichment.v1.AddEnrichmentsRequest.request_enrichments:type_name -> com.coralogix.enrichment.v1.EnrichmentRequestModel
	12, // 2: com.coralogix.enrichment.v1.AddEnrichmentsResponse.enrichments:type_name -> com.coralogix.enrichment.v1.Enrichment
	14, // 3: com.coralogix.enrichment.v1.RemoveEnrichmentsRequest.enrichment_ids:type_name -> google.protobuf.UInt32Value
	12, // 4: com.coralogix.enrichment.v1.RemoveEnrichmentsResponse.remaining_enrichments:type_name -> com.coralogix.enrichment.v1.Enrichment
	15, // 5: com.coralogix.enrichment.v1.AtomicOverwriteEnrichmentsRequest.enrichment_type:type_name -> com.coralogix.enrichment.v1.EnrichmentType
	16, // 6: com.coralogix.enrichment.v1.AtomicOverwriteEnrichmentsRequest.enrichment_fields:type_name -> com.coralogix.enrichment.v1.EnrichmentFieldDefinition
	12, // 7: com.coralogix.enrichment.v1.AtomicOverwriteEnrichmentsResponse.enrichments:type_name -> com.coralogix.enrichment.v1.Enrichment
	17, // 8: com.coralogix.enrichment.v1.GetCompanyEnrichmentSettingsResponse.enrichment_settings:type_name -> com.coralogix.enrichment.v1.CompanyEnrichmentSettings
	2,  // 9: com.coralogix.enrichment.v1.EnrichmentService.GetEnrichments:input_type -> com.coralogix.enrichment.v1.GetEnrichmentsRequest
	4,  // 10: com.coralogix.enrichment.v1.EnrichmentService.AddEnrichments:input_type -> com.coralogix.enrichment.v1.AddEnrichmentsRequest
	6,  // 11: com.coralogix.enrichment.v1.EnrichmentService.RemoveEnrichments:input_type -> com.coralogix.enrichment.v1.RemoveEnrichmentsRequest
	0,  // 12: com.coralogix.enrichment.v1.EnrichmentService.GetEnrichmentLimit:input_type -> com.coralogix.enrichment.v1.GetEnrichmentLimitRequest
	8,  // 13: com.coralogix.enrichment.v1.EnrichmentService.AtomicOverwriteEnrichments:input_type -> com.coralogix.enrichment.v1.AtomicOverwriteEnrichmentsRequest
	10, // 14: com.coralogix.enrichment.v1.EnrichmentService.GetCompanyEnrichmentSettings:input_type -> com.coralogix.enrichment.v1.GetCompanyEnrichmentSettingsRequest
	3,  // 15: com.coralogix.enrichment.v1.EnrichmentService.GetEnrichments:output_type -> com.coralogix.enrichment.v1.GetEnrichmentsResponse
	5,  // 16: com.coralogix.enrichment.v1.EnrichmentService.AddEnrichments:output_type -> com.coralogix.enrichment.v1.AddEnrichmentsResponse
	7,  // 17: com.coralogix.enrichment.v1.EnrichmentService.RemoveEnrichments:output_type -> com.coralogix.enrichment.v1.RemoveEnrichmentsResponse
	1,  // 18: com.coralogix.enrichment.v1.EnrichmentService.GetEnrichmentLimit:output_type -> com.coralogix.enrichment.v1.GetEnrichmentLimitResponse
	9,  // 19: com.coralogix.enrichment.v1.EnrichmentService.AtomicOverwriteEnrichments:output_type -> com.coralogix.enrichment.v1.AtomicOverwriteEnrichmentsResponse
	11, // 20: com.coralogix.enrichment.v1.EnrichmentService.GetCompanyEnrichmentSettings:output_type -> com.coralogix.enrichment.v1.GetCompanyEnrichmentSettingsResponse
	15, // [15:21] is the sub-list for method output_type
	9,  // [9:15] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_com_coralogix_enrichment_v1_enrichment_service_proto_init() }
func file_com_coralogix_enrichment_v1_enrichment_service_proto_init() {
	if File_com_coralogix_enrichment_v1_enrichment_service_proto != nil {
		return
	}
	file_com_coralogix_enrichment_v1_enrichment_proto_init()
	file_com_coralogix_enrichment_v1_enrichment_request_model_proto_init()
	file_com_coralogix_enrichment_v1_enrichment_type_proto_init()
	file_com_coralogix_enrichment_v1_company_enrichment_settings_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_enrichment_v1_enrichment_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogix_enrichment_v1_enrichment_service_proto_goTypes,
		DependencyIndexes: file_com_coralogix_enrichment_v1_enrichment_service_proto_depIdxs,
		MessageInfos:      file_com_coralogix_enrichment_v1_enrichment_service_proto_msgTypes,
	}.Build()
	File_com_coralogix_enrichment_v1_enrichment_service_proto = out.File
	file_com_coralogix_enrichment_v1_enrichment_service_proto_rawDesc = nil
	file_com_coralogix_enrichment_v1_enrichment_service_proto_goTypes = nil
	file_com_coralogix_enrichment_v1_enrichment_service_proto_depIdxs = nil
}
