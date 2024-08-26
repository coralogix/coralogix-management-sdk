// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.3
// source: com/coralogix/enrichment/v1/enrichment_internal_service.proto

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

type GetInternalEnrichmentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetInternalEnrichmentsRequest) Reset() {
	*x = GetInternalEnrichmentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInternalEnrichmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInternalEnrichmentsRequest) ProtoMessage() {}

func (x *GetInternalEnrichmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInternalEnrichmentsRequest.ProtoReflect.Descriptor instead.
func (*GetInternalEnrichmentsRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDescGZIP(), []int{0}
}

type GetInternalEnrichmentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enrichments []*Enrichment `protobuf:"bytes,1,rep,name=enrichments,proto3" json:"enrichments,omitempty"`
}

func (x *GetInternalEnrichmentsResponse) Reset() {
	*x = GetInternalEnrichmentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInternalEnrichmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInternalEnrichmentsResponse) ProtoMessage() {}

func (x *GetInternalEnrichmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInternalEnrichmentsResponse.ProtoReflect.Descriptor instead.
func (*GetInternalEnrichmentsResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetInternalEnrichmentsResponse) GetEnrichments() []*Enrichment {
	if x != nil {
		return x.Enrichments
	}
	return nil
}

type AddInternalEnrichmentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestEnrichments []*EnrichmentRequestModel `protobuf:"bytes,1,rep,name=request_enrichments,json=requestEnrichments,proto3" json:"request_enrichments,omitempty"`
}

func (x *AddInternalEnrichmentsRequest) Reset() {
	*x = AddInternalEnrichmentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddInternalEnrichmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddInternalEnrichmentsRequest) ProtoMessage() {}

func (x *AddInternalEnrichmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddInternalEnrichmentsRequest.ProtoReflect.Descriptor instead.
func (*AddInternalEnrichmentsRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDescGZIP(), []int{2}
}

func (x *AddInternalEnrichmentsRequest) GetRequestEnrichments() []*EnrichmentRequestModel {
	if x != nil {
		return x.RequestEnrichments
	}
	return nil
}

type AddInternalEnrichmentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enrichments []*Enrichment `protobuf:"bytes,1,rep,name=enrichments,proto3" json:"enrichments,omitempty"`
}

func (x *AddInternalEnrichmentsResponse) Reset() {
	*x = AddInternalEnrichmentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddInternalEnrichmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddInternalEnrichmentsResponse) ProtoMessage() {}

func (x *AddInternalEnrichmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddInternalEnrichmentsResponse.ProtoReflect.Descriptor instead.
func (*AddInternalEnrichmentsResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDescGZIP(), []int{3}
}

func (x *AddInternalEnrichmentsResponse) GetEnrichments() []*Enrichment {
	if x != nil {
		return x.Enrichments
	}
	return nil
}

type RemoveInternalEnrichmentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnrichmentIds []*wrapperspb.UInt32Value `protobuf:"bytes,1,rep,name=enrichment_ids,json=enrichmentIds,proto3" json:"enrichment_ids,omitempty"`
}

func (x *RemoveInternalEnrichmentsRequest) Reset() {
	*x = RemoveInternalEnrichmentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveInternalEnrichmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveInternalEnrichmentsRequest) ProtoMessage() {}

func (x *RemoveInternalEnrichmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveInternalEnrichmentsRequest.ProtoReflect.Descriptor instead.
func (*RemoveInternalEnrichmentsRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveInternalEnrichmentsRequest) GetEnrichmentIds() []*wrapperspb.UInt32Value {
	if x != nil {
		return x.EnrichmentIds
	}
	return nil
}

type RemoveInternalEnrichmentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RemainingEnrichments []*Enrichment `protobuf:"bytes,1,rep,name=remaining_enrichments,json=remainingEnrichments,proto3" json:"remaining_enrichments,omitempty"`
}

func (x *RemoveInternalEnrichmentsResponse) Reset() {
	*x = RemoveInternalEnrichmentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveInternalEnrichmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveInternalEnrichmentsResponse) ProtoMessage() {}

func (x *RemoveInternalEnrichmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveInternalEnrichmentsResponse.ProtoReflect.Descriptor instead.
func (*RemoveInternalEnrichmentsResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDescGZIP(), []int{5}
}

func (x *RemoveInternalEnrichmentsResponse) GetRemainingEnrichments() []*Enrichment {
	if x != nil {
		return x.RemainingEnrichments
	}
	return nil
}

var File_com_coralogix_enrichment_v1_enrichment_internal_service_proto protoreflect.FileDescriptor

var file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e,
	0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65,
	0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x65, 0x6e, 0x72, 0x69,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x6b, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0b, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x6e, 0x72, 0x69,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x22, 0x85, 0x01, 0x0a, 0x1d, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x64, 0x0a, 0x13, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x65,
	0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x6e,
	0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x6b, 0x0a, 0x1e, 0x41, 0x64, 0x64,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0b, 0x65,
	0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x65, 0x6e, 0x72, 0x69, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x67, 0x0a, 0x20, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x0e, 0x65, 0x6e,
	0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0d, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x22,
	0x81, 0x01, 0x0a, 0x21, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x15, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x5f, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x14, 0x72,
	0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x32, 0xc3, 0x04, 0x0a, 0x19, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xb1, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x3a, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x6e, 0x72,
	0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0xc2, 0xb8, 0x02, 0x1a, 0x0a, 0x18, 0x47, 0x65, 0x74,
	0x20, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0xb1, 0x01, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x3a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x64, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x6e, 0x72,
	0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0xc2, 0xb8, 0x02, 0x1a, 0x0a,
	0x18, 0x41, 0x64, 0x64, 0x20, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x65, 0x6e,
	0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0xbd, 0x01, 0x0a, 0x19, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x6e, 0x72, 0x69,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x3d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0xc2, 0xb8, 0x02, 0x1d, 0x0a, 0x1b, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x20, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x65, 0x6e,
	0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDescOnce sync.Once
	file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDescData = file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDesc
)

func file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDescGZIP() []byte {
	file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDescOnce.Do(func() {
		file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDescData)
	})
	return file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDescData
}

var file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_goTypes = []interface{}{
	(*GetInternalEnrichmentsRequest)(nil),     // 0: com.coralogix.enrichment.v1.GetInternalEnrichmentsRequest
	(*GetInternalEnrichmentsResponse)(nil),    // 1: com.coralogix.enrichment.v1.GetInternalEnrichmentsResponse
	(*AddInternalEnrichmentsRequest)(nil),     // 2: com.coralogix.enrichment.v1.AddInternalEnrichmentsRequest
	(*AddInternalEnrichmentsResponse)(nil),    // 3: com.coralogix.enrichment.v1.AddInternalEnrichmentsResponse
	(*RemoveInternalEnrichmentsRequest)(nil),  // 4: com.coralogix.enrichment.v1.RemoveInternalEnrichmentsRequest
	(*RemoveInternalEnrichmentsResponse)(nil), // 5: com.coralogix.enrichment.v1.RemoveInternalEnrichmentsResponse
	(*Enrichment)(nil),                        // 6: com.coralogix.enrichment.v1.Enrichment
	(*EnrichmentRequestModel)(nil),            // 7: com.coralogix.enrichment.v1.EnrichmentRequestModel
	(*wrapperspb.UInt32Value)(nil),            // 8: google.protobuf.UInt32Value
}
var file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_depIdxs = []int32{
	6, // 0: com.coralogix.enrichment.v1.GetInternalEnrichmentsResponse.enrichments:type_name -> com.coralogix.enrichment.v1.Enrichment
	7, // 1: com.coralogix.enrichment.v1.AddInternalEnrichmentsRequest.request_enrichments:type_name -> com.coralogix.enrichment.v1.EnrichmentRequestModel
	6, // 2: com.coralogix.enrichment.v1.AddInternalEnrichmentsResponse.enrichments:type_name -> com.coralogix.enrichment.v1.Enrichment
	8, // 3: com.coralogix.enrichment.v1.RemoveInternalEnrichmentsRequest.enrichment_ids:type_name -> google.protobuf.UInt32Value
	6, // 4: com.coralogix.enrichment.v1.RemoveInternalEnrichmentsResponse.remaining_enrichments:type_name -> com.coralogix.enrichment.v1.Enrichment
	0, // 5: com.coralogix.enrichment.v1.EnrichmentInternalService.GetInternalEnrichments:input_type -> com.coralogix.enrichment.v1.GetInternalEnrichmentsRequest
	2, // 6: com.coralogix.enrichment.v1.EnrichmentInternalService.AddInternalEnrichments:input_type -> com.coralogix.enrichment.v1.AddInternalEnrichmentsRequest
	4, // 7: com.coralogix.enrichment.v1.EnrichmentInternalService.RemoveInternalEnrichments:input_type -> com.coralogix.enrichment.v1.RemoveInternalEnrichmentsRequest
	1, // 8: com.coralogix.enrichment.v1.EnrichmentInternalService.GetInternalEnrichments:output_type -> com.coralogix.enrichment.v1.GetInternalEnrichmentsResponse
	3, // 9: com.coralogix.enrichment.v1.EnrichmentInternalService.AddInternalEnrichments:output_type -> com.coralogix.enrichment.v1.AddInternalEnrichmentsResponse
	5, // 10: com.coralogix.enrichment.v1.EnrichmentInternalService.RemoveInternalEnrichments:output_type -> com.coralogix.enrichment.v1.RemoveInternalEnrichmentsResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_init() }
func file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_init() {
	if File_com_coralogix_enrichment_v1_enrichment_internal_service_proto != nil {
		return
	}
	file_com_coralogix_enrichment_v1_enrichment_proto_init()
	file_com_coralogix_enrichment_v1_enrichment_request_model_proto_init()
	file_com_coralogix_enrichment_v1_audit_log_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInternalEnrichmentsRequest); i {
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
		file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInternalEnrichmentsResponse); i {
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
		file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddInternalEnrichmentsRequest); i {
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
		file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddInternalEnrichmentsResponse); i {
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
		file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveInternalEnrichmentsRequest); i {
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
		file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveInternalEnrichmentsResponse); i {
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
			RawDescriptor: file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_goTypes,
		DependencyIndexes: file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_depIdxs,
		MessageInfos:      file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes,
	}.Build()
	File_com_coralogix_enrichment_v1_enrichment_internal_service_proto = out.File
	file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDesc = nil
	file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_goTypes = nil
	file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_depIdxs = nil
}
