// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.26.1
// source: com/coralogix/integrations/v1/external_token_service.proto

package v1

import (
	_ "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/common/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
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

type GenerateNewTokenRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Platform      PushBasedPlatform      `protobuf:"varint,1,opt,name=platform,proto3,enum=com.coralogix.integrations.v1.PushBasedPlatform" json:"platform,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateNewTokenRequest) Reset() {
	*x = GenerateNewTokenRequest{}
	mi := &file_com_coralogix_integrations_v1_external_token_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateNewTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateNewTokenRequest) ProtoMessage() {}

func (x *GenerateNewTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_integrations_v1_external_token_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateNewTokenRequest.ProtoReflect.Descriptor instead.
func (*GenerateNewTokenRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_integrations_v1_external_token_service_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateNewTokenRequest) GetPlatform() PushBasedPlatform {
	if x != nil {
		return x.Platform
	}
	return PushBasedPlatform_UNDEFINED
}

type GenerateNewTokenResponse struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Token         *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Id            *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateNewTokenResponse) Reset() {
	*x = GenerateNewTokenResponse{}
	mi := &file_com_coralogix_integrations_v1_external_token_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateNewTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateNewTokenResponse) ProtoMessage() {}

func (x *GenerateNewTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_integrations_v1_external_token_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateNewTokenResponse.ProtoReflect.Descriptor instead.
func (*GenerateNewTokenResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_integrations_v1_external_token_service_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateNewTokenResponse) GetToken() *wrapperspb.StringValue {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *GenerateNewTokenResponse) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

type UpdateTokenRequest struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Id            *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTokenRequest) Reset() {
	*x = UpdateTokenRequest{}
	mi := &file_com_coralogix_integrations_v1_external_token_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTokenRequest) ProtoMessage() {}

func (x *UpdateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_integrations_v1_external_token_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTokenRequest.ProtoReflect.Descriptor instead.
func (*UpdateTokenRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_integrations_v1_external_token_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateTokenRequest) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

type UpdateTokenResponse struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Token         *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTokenResponse) Reset() {
	*x = UpdateTokenResponse{}
	mi := &file_com_coralogix_integrations_v1_external_token_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTokenResponse) ProtoMessage() {}

func (x *UpdateTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_integrations_v1_external_token_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTokenResponse.ProtoReflect.Descriptor instead.
func (*UpdateTokenResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_integrations_v1_external_token_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateTokenResponse) GetToken() *wrapperspb.StringValue {
	if x != nil {
		return x.Token
	}
	return nil
}

var File_com_coralogix_integrations_v1_external_token_service_proto protoreflect.FileDescriptor

var file_com_coralogix_integrations_v1_external_token_service_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x27, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x37, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x64, 0x5f, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a,
	0x17, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x42,
	0x61, 0x73, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x08, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x7c, 0x0a, 0x18, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x4e, 0x65, 0x77, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x42, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x32, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x32, 0x93, 0x03, 0x0a, 0x14, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xc3, 0x01, 0x0a,
	0x10, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x4e, 0x65, 0x77, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x3e, 0xc2, 0xb8, 0x02, 0x1f, 0x0a, 0x1d, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x13, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2d, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x12, 0xb4, 0x01, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3e, 0xc2, 0xb8, 0x02, 0x1a, 0x0a,
	0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x6e, 0x20, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x1a,
	0x18, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2d, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_com_coralogix_integrations_v1_external_token_service_proto_rawDescOnce sync.Once
	file_com_coralogix_integrations_v1_external_token_service_proto_rawDescData = file_com_coralogix_integrations_v1_external_token_service_proto_rawDesc
)

func file_com_coralogix_integrations_v1_external_token_service_proto_rawDescGZIP() []byte {
	file_com_coralogix_integrations_v1_external_token_service_proto_rawDescOnce.Do(func() {
		file_com_coralogix_integrations_v1_external_token_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_integrations_v1_external_token_service_proto_rawDescData)
	})
	return file_com_coralogix_integrations_v1_external_token_service_proto_rawDescData
}

var file_com_coralogix_integrations_v1_external_token_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_com_coralogix_integrations_v1_external_token_service_proto_goTypes = []any{
	(*GenerateNewTokenRequest)(nil),  // 0: com.coralogix.integrations.v1.GenerateNewTokenRequest
	(*GenerateNewTokenResponse)(nil), // 1: com.coralogix.integrations.v1.GenerateNewTokenResponse
	(*UpdateTokenRequest)(nil),       // 2: com.coralogix.integrations.v1.UpdateTokenRequest
	(*UpdateTokenResponse)(nil),      // 3: com.coralogix.integrations.v1.UpdateTokenResponse
	(PushBasedPlatform)(0),           // 4: com.coralogix.integrations.v1.PushBasedPlatform
	(*wrapperspb.StringValue)(nil),   // 5: google.protobuf.StringValue
}
var file_com_coralogix_integrations_v1_external_token_service_proto_depIdxs = []int32{
	4, // 0: com.coralogix.integrations.v1.GenerateNewTokenRequest.platform:type_name -> com.coralogix.integrations.v1.PushBasedPlatform
	5, // 1: com.coralogix.integrations.v1.GenerateNewTokenResponse.token:type_name -> google.protobuf.StringValue
	5, // 2: com.coralogix.integrations.v1.GenerateNewTokenResponse.id:type_name -> google.protobuf.StringValue
	5, // 3: com.coralogix.integrations.v1.UpdateTokenRequest.id:type_name -> google.protobuf.StringValue
	5, // 4: com.coralogix.integrations.v1.UpdateTokenResponse.token:type_name -> google.protobuf.StringValue
	0, // 5: com.coralogix.integrations.v1.ExternalTokenService.GenerateNewToken:input_type -> com.coralogix.integrations.v1.GenerateNewTokenRequest
	2, // 6: com.coralogix.integrations.v1.ExternalTokenService.UpdateToken:input_type -> com.coralogix.integrations.v1.UpdateTokenRequest
	1, // 7: com.coralogix.integrations.v1.ExternalTokenService.GenerateNewToken:output_type -> com.coralogix.integrations.v1.GenerateNewTokenResponse
	3, // 8: com.coralogix.integrations.v1.ExternalTokenService.UpdateToken:output_type -> com.coralogix.integrations.v1.UpdateTokenResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_com_coralogix_integrations_v1_external_token_service_proto_init() }
func file_com_coralogix_integrations_v1_external_token_service_proto_init() {
	if File_com_coralogix_integrations_v1_external_token_service_proto != nil {
		return
	}
	file_com_coralogix_integrations_v1_push_based_platform_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_integrations_v1_external_token_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogix_integrations_v1_external_token_service_proto_goTypes,
		DependencyIndexes: file_com_coralogix_integrations_v1_external_token_service_proto_depIdxs,
		MessageInfos:      file_com_coralogix_integrations_v1_external_token_service_proto_msgTypes,
	}.Build()
	File_com_coralogix_integrations_v1_external_token_service_proto = out.File
	file_com_coralogix_integrations_v1_external_token_service_proto_rawDesc = nil
	file_com_coralogix_integrations_v1_external_token_service_proto_goTypes = nil
	file_com_coralogix_integrations_v1_external_token_service_proto_depIdxs = nil
}
