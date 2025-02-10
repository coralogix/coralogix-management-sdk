// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: com/coralogix/extensions/v1/extension_content_management_service.proto

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

type ImportAndReleaseExtensionsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImportAndReleaseExtensionsResponse) Reset() {
	*x = ImportAndReleaseExtensionsResponse{}
	mi := &file_com_coralogix_extensions_v1_extension_content_management_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImportAndReleaseExtensionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportAndReleaseExtensionsResponse) ProtoMessage() {}

func (x *ImportAndReleaseExtensionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_extensions_v1_extension_content_management_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportAndReleaseExtensionsResponse.ProtoReflect.Descriptor instead.
func (*ImportAndReleaseExtensionsResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_extensions_v1_extension_content_management_service_proto_rawDescGZIP(), []int{0}
}

type ValidateExtensionItemsResponse struct {
	state             protoimpl.MessageState                             `protogen:"open.v1"`
	InvalidExtensions []*ValidateExtensionItemsResponse_InvalidExtension `protobuf:"bytes,1,rep,name=invalid_extensions,json=invalidExtensions,proto3" json:"invalid_extensions,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *ValidateExtensionItemsResponse) Reset() {
	*x = ValidateExtensionItemsResponse{}
	mi := &file_com_coralogix_extensions_v1_extension_content_management_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateExtensionItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateExtensionItemsResponse) ProtoMessage() {}

func (x *ValidateExtensionItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_extensions_v1_extension_content_management_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateExtensionItemsResponse.ProtoReflect.Descriptor instead.
func (*ValidateExtensionItemsResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_extensions_v1_extension_content_management_service_proto_rawDescGZIP(), []int{1}
}

func (x *ValidateExtensionItemsResponse) GetInvalidExtensions() []*ValidateExtensionItemsResponse_InvalidExtension {
	if x != nil {
		return x.InvalidExtensions
	}
	return nil
}

type ValidateExtensionItemsResponse_InvalidExtension struct {
	state       protoimpl.MessageState  `protogen:"open.v1"`
	ExtensionId *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=extension_id,json=extensionId,proto3" json:"extension_id,omitempty"`
	// Includes extension item parsing errors.
	ErrorMessage  *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ValidateExtensionItemsResponse_InvalidExtension) Reset() {
	*x = ValidateExtensionItemsResponse_InvalidExtension{}
	mi := &file_com_coralogix_extensions_v1_extension_content_management_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateExtensionItemsResponse_InvalidExtension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateExtensionItemsResponse_InvalidExtension) ProtoMessage() {}

func (x *ValidateExtensionItemsResponse_InvalidExtension) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_extensions_v1_extension_content_management_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateExtensionItemsResponse_InvalidExtension.ProtoReflect.Descriptor instead.
func (*ValidateExtensionItemsResponse_InvalidExtension) Descriptor() ([]byte, []int) {
	return file_com_coralogix_extensions_v1_extension_content_management_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ValidateExtensionItemsResponse_InvalidExtension) GetExtensionId() *wrapperspb.StringValue {
	if x != nil {
		return x.ExtensionId
	}
	return nil
}

func (x *ValidateExtensionItemsResponse_InvalidExtension) GetErrorMessage() *wrapperspb.StringValue {
	if x != nil {
		return x.ErrorMessage
	}
	return nil
}

var File_com_coralogix_extensions_v1_extension_content_management_service_proto protoreflect.FileDescriptor

var file_com_coralogix_extensions_v1_extension_content_management_service_proto_rawDesc = []byte{
	0x0a, 0x46, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x27, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x22, 0x49,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0xb6, 0x02, 0x0a, 0x1e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7b, 0x0a, 0x12, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x4c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49, 0x6e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x11,
	0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x1a, 0x96, 0x01, 0x0a, 0x10, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x0c, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xdf, 0x03, 0x0a, 0x21, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0xe3, 0x01, 0x0a, 0x1a, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x6e, 0x64, 0x52, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x3f, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x41, 0x6e, 0x64, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x56, 0xca, 0xb8,
	0x02, 0x1f, 0x0a, 0x1d, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x20, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x3a, 0x01, 0x2a, 0x22, 0x28, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2d,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x28, 0x01, 0x12, 0xd3, 0x01, 0x0a, 0x16, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x12, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4e, 0xca, 0xb8, 0x02, 0x1a,
	0x0a, 0x18, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x20, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x20, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a,
	0x3a, 0x01, 0x2a, 0x22, 0x25, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2d, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2d, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x28, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_extensions_v1_extension_content_management_service_proto_rawDescOnce sync.Once
	file_com_coralogix_extensions_v1_extension_content_management_service_proto_rawDescData = file_com_coralogix_extensions_v1_extension_content_management_service_proto_rawDesc
)

func file_com_coralogix_extensions_v1_extension_content_management_service_proto_rawDescGZIP() []byte {
	file_com_coralogix_extensions_v1_extension_content_management_service_proto_rawDescOnce.Do(func() {
		file_com_coralogix_extensions_v1_extension_content_management_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_extensions_v1_extension_content_management_service_proto_rawDescData)
	})
	return file_com_coralogix_extensions_v1_extension_content_management_service_proto_rawDescData
}

var file_com_coralogix_extensions_v1_extension_content_management_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogix_extensions_v1_extension_content_management_service_proto_goTypes = []any{
	(*ImportAndReleaseExtensionsResponse)(nil),              // 0: com.coralogix.extensions.v1.ImportAndReleaseExtensionsResponse
	(*ValidateExtensionItemsResponse)(nil),                  // 1: com.coralogix.extensions.v1.ValidateExtensionItemsResponse
	(*ValidateExtensionItemsResponse_InvalidExtension)(nil), // 2: com.coralogix.extensions.v1.ValidateExtensionItemsResponse.InvalidExtension
	(*wrapperspb.StringValue)(nil),                          // 3: google.protobuf.StringValue
	(*ExtensionData)(nil),                                   // 4: com.coralogix.extensions.v1.ExtensionData
}
var file_com_coralogix_extensions_v1_extension_content_management_service_proto_depIdxs = []int32{
	2, // 0: com.coralogix.extensions.v1.ValidateExtensionItemsResponse.invalid_extensions:type_name -> com.coralogix.extensions.v1.ValidateExtensionItemsResponse.InvalidExtension
	3, // 1: com.coralogix.extensions.v1.ValidateExtensionItemsResponse.InvalidExtension.extension_id:type_name -> google.protobuf.StringValue
	3, // 2: com.coralogix.extensions.v1.ValidateExtensionItemsResponse.InvalidExtension.error_message:type_name -> google.protobuf.StringValue
	4, // 3: com.coralogix.extensions.v1.ExtensionContentManagementService.ImportAndReleaseExtensions:input_type -> com.coralogix.extensions.v1.ExtensionData
	4, // 4: com.coralogix.extensions.v1.ExtensionContentManagementService.ValidateExtensionItems:input_type -> com.coralogix.extensions.v1.ExtensionData
	0, // 5: com.coralogix.extensions.v1.ExtensionContentManagementService.ImportAndReleaseExtensions:output_type -> com.coralogix.extensions.v1.ImportAndReleaseExtensionsResponse
	1, // 6: com.coralogix.extensions.v1.ExtensionContentManagementService.ValidateExtensionItems:output_type -> com.coralogix.extensions.v1.ValidateExtensionItemsResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_com_coralogix_extensions_v1_extension_content_management_service_proto_init() }
func file_com_coralogix_extensions_v1_extension_content_management_service_proto_init() {
	if File_com_coralogix_extensions_v1_extension_content_management_service_proto != nil {
		return
	}
	file_com_coralogix_extensions_v1_extension_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_extensions_v1_extension_content_management_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogix_extensions_v1_extension_content_management_service_proto_goTypes,
		DependencyIndexes: file_com_coralogix_extensions_v1_extension_content_management_service_proto_depIdxs,
		MessageInfos:      file_com_coralogix_extensions_v1_extension_content_management_service_proto_msgTypes,
	}.Build()
	File_com_coralogix_extensions_v1_extension_content_management_service_proto = out.File
	file_com_coralogix_extensions_v1_extension_content_management_service_proto_rawDesc = nil
	file_com_coralogix_extensions_v1_extension_content_management_service_proto_goTypes = nil
	file_com_coralogix_extensions_v1_extension_content_management_service_proto_depIdxs = nil
}
