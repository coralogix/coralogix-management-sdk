// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: com/coralogixapis/apm/services/v1/apm_service_service.proto

package v1

import (
	_ "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/common/v1"
	v2 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/apm/common/v2"
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

type GetApmServiceRequest struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Id            *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetApmServiceRequest) Reset() {
	*x = GetApmServiceRequest{}
	mi := &file_com_coralogixapis_apm_services_v1_apm_service_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetApmServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApmServiceRequest) ProtoMessage() {}

func (x *GetApmServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_services_v1_apm_service_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApmServiceRequest.ProtoReflect.Descriptor instead.
func (*GetApmServiceRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_services_v1_apm_service_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetApmServiceRequest) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

type GetApmServiceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Service       *ApmService            `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetApmServiceResponse) Reset() {
	*x = GetApmServiceResponse{}
	mi := &file_com_coralogixapis_apm_services_v1_apm_service_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetApmServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApmServiceResponse) ProtoMessage() {}

func (x *GetApmServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_services_v1_apm_service_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApmServiceResponse.ProtoReflect.Descriptor instead.
func (*GetApmServiceResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_services_v1_apm_service_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetApmServiceResponse) GetService() *ApmService {
	if x != nil {
		return x.Service
	}
	return nil
}

type BatchGetApmServicesRequest struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Ids           []*wrapperspb.StringValue `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BatchGetApmServicesRequest) Reset() {
	*x = BatchGetApmServicesRequest{}
	mi := &file_com_coralogixapis_apm_services_v1_apm_service_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchGetApmServicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetApmServicesRequest) ProtoMessage() {}

func (x *BatchGetApmServicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_services_v1_apm_service_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetApmServicesRequest.ProtoReflect.Descriptor instead.
func (*BatchGetApmServicesRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_services_v1_apm_service_service_proto_rawDescGZIP(), []int{2}
}

func (x *BatchGetApmServicesRequest) GetIds() []*wrapperspb.StringValue {
	if x != nil {
		return x.Ids
	}
	return nil
}

type BatchGetApmServicesResponse struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Services      map[string]*ApmService    `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	NotFoundIds   []*wrapperspb.StringValue `protobuf:"bytes,2,rep,name=not_found_ids,json=notFoundIds,proto3" json:"not_found_ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BatchGetApmServicesResponse) Reset() {
	*x = BatchGetApmServicesResponse{}
	mi := &file_com_coralogixapis_apm_services_v1_apm_service_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchGetApmServicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetApmServicesResponse) ProtoMessage() {}

func (x *BatchGetApmServicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_services_v1_apm_service_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetApmServicesResponse.ProtoReflect.Descriptor instead.
func (*BatchGetApmServicesResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_services_v1_apm_service_service_proto_rawDescGZIP(), []int{3}
}

func (x *BatchGetApmServicesResponse) GetServices() map[string]*ApmService {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *BatchGetApmServicesResponse) GetNotFoundIds() []*wrapperspb.StringValue {
	if x != nil {
		return x.NotFoundIds
	}
	return nil
}

type DeleteApmServiceRequest struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Id            *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteApmServiceRequest) Reset() {
	*x = DeleteApmServiceRequest{}
	mi := &file_com_coralogixapis_apm_services_v1_apm_service_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteApmServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteApmServiceRequest) ProtoMessage() {}

func (x *DeleteApmServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_services_v1_apm_service_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteApmServiceRequest.ProtoReflect.Descriptor instead.
func (*DeleteApmServiceRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_services_v1_apm_service_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteApmServiceRequest) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

type DeleteApmServiceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteApmServiceResponse) Reset() {
	*x = DeleteApmServiceResponse{}
	mi := &file_com_coralogixapis_apm_services_v1_apm_service_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteApmServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteApmServiceResponse) ProtoMessage() {}

func (x *DeleteApmServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_services_v1_apm_service_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteApmServiceResponse.ProtoReflect.Descriptor instead.
func (*DeleteApmServiceResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_services_v1_apm_service_service_proto_rawDescGZIP(), []int{5}
}

type ListApmServicesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderBy       *v2.OrderBy            `protobuf:"bytes,1,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListApmServicesRequest) Reset() {
	*x = ListApmServicesRequest{}
	mi := &file_com_coralogixapis_apm_services_v1_apm_service_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListApmServicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApmServicesRequest) ProtoMessage() {}

func (x *ListApmServicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_services_v1_apm_service_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApmServicesRequest.ProtoReflect.Descriptor instead.
func (*ListApmServicesRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_services_v1_apm_service_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListApmServicesRequest) GetOrderBy() *v2.OrderBy {
	if x != nil {
		return x.OrderBy
	}
	return nil
}

type ListApmServicesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Services      []*ApmService          `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListApmServicesResponse) Reset() {
	*x = ListApmServicesResponse{}
	mi := &file_com_coralogixapis_apm_services_v1_apm_service_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListApmServicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApmServicesResponse) ProtoMessage() {}

func (x *ListApmServicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_services_v1_apm_service_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApmServicesResponse.ProtoReflect.Descriptor instead.
func (*ListApmServicesResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_services_v1_apm_service_service_proto_rawDescGZIP(), []int{7}
}

func (x *ListApmServicesResponse) GetServices() []*ApmService {
	if x != nil {
		return x.Services
	}
	return nil
}

var File_com_coralogixapis_apm_services_v1_apm_service_service_proto protoreflect.FileDescriptor

var file_com_coralogixapis_apm_services_v1_apm_service_service_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x33, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x70, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x60, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x70, 0x6d, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x70, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x4c, 0x0a, 0x1a, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47,
	0x65, 0x74, 0x41, 0x70, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x03, 0x69, 0x64, 0x73, 0x22, 0xb5, 0x02, 0x0a, 0x1b, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65,
	0x74, 0x41, 0x70, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x68, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x47, 0x65, 0x74, 0x41, 0x70, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x40,
	0x0a, 0x0d, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0b, 0x6e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x64, 0x73,
	0x1a, 0x6a, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x43, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x47, 0x0a, 0x17,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1a, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x70, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x5d, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x6d, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79,
	0x22, 0x64, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x70, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x32, 0xb0, 0x05, 0x0a, 0x11, 0x41, 0x70, 0x6d, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x99, 0x01, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x41, 0x70, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x70, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x15, 0xc2, 0xb8, 0x02, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x20, 0x41, 0x70, 0x6d,
	0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xb2, 0x01, 0x0a, 0x13, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x47, 0x65, 0x74, 0x41, 0x70, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x12, 0x3d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x41, 0x70, 0x6d,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x3e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x41, 0x70, 0x6d, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1c, 0xc2, 0xb8, 0x02, 0x18, 0x0a, 0x16, 0x42, 0x61, 0x74, 0x63, 0x68, 0x20, 0x47, 0x65, 0x74,
	0x20, 0x41, 0x70, 0x6d, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0xa5, 0x01,
	0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x6d,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x6d, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0xc2, 0xb8, 0x02,
	0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x41, 0x70, 0x6d, 0x20, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa1, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70,
	0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70,
	0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x70, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x6d,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x17, 0xc2, 0xb8, 0x02, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x41, 0x70, 0x6d,
	0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_com_coralogixapis_apm_services_v1_apm_service_service_proto_rawDescOnce sync.Once
	file_com_coralogixapis_apm_services_v1_apm_service_service_proto_rawDescData = file_com_coralogixapis_apm_services_v1_apm_service_service_proto_rawDesc
)

func file_com_coralogixapis_apm_services_v1_apm_service_service_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_apm_services_v1_apm_service_service_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_apm_services_v1_apm_service_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_apm_services_v1_apm_service_service_proto_rawDescData)
	})
	return file_com_coralogixapis_apm_services_v1_apm_service_service_proto_rawDescData
}

var file_com_coralogixapis_apm_services_v1_apm_service_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_com_coralogixapis_apm_services_v1_apm_service_service_proto_goTypes = []any{
	(*GetApmServiceRequest)(nil),        // 0: com.coralogixapis.apm.services.v1.GetApmServiceRequest
	(*GetApmServiceResponse)(nil),       // 1: com.coralogixapis.apm.services.v1.GetApmServiceResponse
	(*BatchGetApmServicesRequest)(nil),  // 2: com.coralogixapis.apm.services.v1.BatchGetApmServicesRequest
	(*BatchGetApmServicesResponse)(nil), // 3: com.coralogixapis.apm.services.v1.BatchGetApmServicesResponse
	(*DeleteApmServiceRequest)(nil),     // 4: com.coralogixapis.apm.services.v1.DeleteApmServiceRequest
	(*DeleteApmServiceResponse)(nil),    // 5: com.coralogixapis.apm.services.v1.DeleteApmServiceResponse
	(*ListApmServicesRequest)(nil),      // 6: com.coralogixapis.apm.services.v1.ListApmServicesRequest
	(*ListApmServicesResponse)(nil),     // 7: com.coralogixapis.apm.services.v1.ListApmServicesResponse
	nil,                                 // 8: com.coralogixapis.apm.services.v1.BatchGetApmServicesResponse.ServicesEntry
	(*wrapperspb.StringValue)(nil),      // 9: google.protobuf.StringValue
	(*ApmService)(nil),                  // 10: com.coralogixapis.apm.services.v1.ApmService
	(*v2.OrderBy)(nil),                  // 11: com.coralogixapis.apm.common.v2.OrderBy
}
var file_com_coralogixapis_apm_services_v1_apm_service_service_proto_depIdxs = []int32{
	9,  // 0: com.coralogixapis.apm.services.v1.GetApmServiceRequest.id:type_name -> google.protobuf.StringValue
	10, // 1: com.coralogixapis.apm.services.v1.GetApmServiceResponse.service:type_name -> com.coralogixapis.apm.services.v1.ApmService
	9,  // 2: com.coralogixapis.apm.services.v1.BatchGetApmServicesRequest.ids:type_name -> google.protobuf.StringValue
	8,  // 3: com.coralogixapis.apm.services.v1.BatchGetApmServicesResponse.services:type_name -> com.coralogixapis.apm.services.v1.BatchGetApmServicesResponse.ServicesEntry
	9,  // 4: com.coralogixapis.apm.services.v1.BatchGetApmServicesResponse.not_found_ids:type_name -> google.protobuf.StringValue
	9,  // 5: com.coralogixapis.apm.services.v1.DeleteApmServiceRequest.id:type_name -> google.protobuf.StringValue
	11, // 6: com.coralogixapis.apm.services.v1.ListApmServicesRequest.order_by:type_name -> com.coralogixapis.apm.common.v2.OrderBy
	10, // 7: com.coralogixapis.apm.services.v1.ListApmServicesResponse.services:type_name -> com.coralogixapis.apm.services.v1.ApmService
	10, // 8: com.coralogixapis.apm.services.v1.BatchGetApmServicesResponse.ServicesEntry.value:type_name -> com.coralogixapis.apm.services.v1.ApmService
	0,  // 9: com.coralogixapis.apm.services.v1.ApmServiceService.GetApmService:input_type -> com.coralogixapis.apm.services.v1.GetApmServiceRequest
	2,  // 10: com.coralogixapis.apm.services.v1.ApmServiceService.BatchGetApmServices:input_type -> com.coralogixapis.apm.services.v1.BatchGetApmServicesRequest
	4,  // 11: com.coralogixapis.apm.services.v1.ApmServiceService.DeleteApmService:input_type -> com.coralogixapis.apm.services.v1.DeleteApmServiceRequest
	6,  // 12: com.coralogixapis.apm.services.v1.ApmServiceService.ListApmServices:input_type -> com.coralogixapis.apm.services.v1.ListApmServicesRequest
	1,  // 13: com.coralogixapis.apm.services.v1.ApmServiceService.GetApmService:output_type -> com.coralogixapis.apm.services.v1.GetApmServiceResponse
	3,  // 14: com.coralogixapis.apm.services.v1.ApmServiceService.BatchGetApmServices:output_type -> com.coralogixapis.apm.services.v1.BatchGetApmServicesResponse
	5,  // 15: com.coralogixapis.apm.services.v1.ApmServiceService.DeleteApmService:output_type -> com.coralogixapis.apm.services.v1.DeleteApmServiceResponse
	7,  // 16: com.coralogixapis.apm.services.v1.ApmServiceService.ListApmServices:output_type -> com.coralogixapis.apm.services.v1.ListApmServicesResponse
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_apm_services_v1_apm_service_service_proto_init() }
func file_com_coralogixapis_apm_services_v1_apm_service_service_proto_init() {
	if File_com_coralogixapis_apm_services_v1_apm_service_service_proto != nil {
		return
	}
	file_com_coralogixapis_apm_services_v1_apm_service_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_apm_services_v1_apm_service_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogixapis_apm_services_v1_apm_service_service_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_apm_services_v1_apm_service_service_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_apm_services_v1_apm_service_service_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_apm_services_v1_apm_service_service_proto = out.File
	file_com_coralogixapis_apm_services_v1_apm_service_service_proto_rawDesc = nil
	file_com_coralogixapis_apm_services_v1_apm_service_service_proto_goTypes = nil
	file_com_coralogixapis_apm_services_v1_apm_service_service_proto_depIdxs = nil
}
