// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: com/coralogix/extensions/v1/extension_testing_service.proto

package v1

import (
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

type InitializeTestingRevisionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExtensionData *ExtensionData `protobuf:"bytes,1,opt,name=extension_data,json=extensionData,proto3" json:"extension_data,omitempty"`
}

func (x *InitializeTestingRevisionRequest) Reset() {
	*x = InitializeTestingRevisionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_extensions_v1_extension_testing_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeTestingRevisionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeTestingRevisionRequest) ProtoMessage() {}

func (x *InitializeTestingRevisionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_extensions_v1_extension_testing_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeTestingRevisionRequest.ProtoReflect.Descriptor instead.
func (*InitializeTestingRevisionRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_extensions_v1_extension_testing_service_proto_rawDescGZIP(), []int{0}
}

func (x *InitializeTestingRevisionRequest) GetExtensionData() *ExtensionData {
	if x != nil {
		return x.ExtensionData
	}
	return nil
}

type InitializeTestingRevisionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InitializeTestingRevisionResponse) Reset() {
	*x = InitializeTestingRevisionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_extensions_v1_extension_testing_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeTestingRevisionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeTestingRevisionResponse) ProtoMessage() {}

func (x *InitializeTestingRevisionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_extensions_v1_extension_testing_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeTestingRevisionResponse.ProtoReflect.Descriptor instead.
func (*InitializeTestingRevisionResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_extensions_v1_extension_testing_service_proto_rawDescGZIP(), []int{1}
}

type CleanupTestingRevisionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CleanupTestingRevisionRequest) Reset() {
	*x = CleanupTestingRevisionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_extensions_v1_extension_testing_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CleanupTestingRevisionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CleanupTestingRevisionRequest) ProtoMessage() {}

func (x *CleanupTestingRevisionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_extensions_v1_extension_testing_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CleanupTestingRevisionRequest.ProtoReflect.Descriptor instead.
func (*CleanupTestingRevisionRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_extensions_v1_extension_testing_service_proto_rawDescGZIP(), []int{2}
}

func (x *CleanupTestingRevisionRequest) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

type CleanupTestingRevisionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CleanupTestingRevisionResponse) Reset() {
	*x = CleanupTestingRevisionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_extensions_v1_extension_testing_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CleanupTestingRevisionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CleanupTestingRevisionResponse) ProtoMessage() {}

func (x *CleanupTestingRevisionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_extensions_v1_extension_testing_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CleanupTestingRevisionResponse.ProtoReflect.Descriptor instead.
func (*CleanupTestingRevisionResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_extensions_v1_extension_testing_service_proto_rawDescGZIP(), []int{3}
}

type TestExtensionRevisionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExtensionData    *ExtensionData        `protobuf:"bytes,1,opt,name=extension_data,json=extensionData,proto3" json:"extension_data,omitempty"`
	CleanupAfterTest *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=cleanup_after_test,json=cleanupAfterTest,proto3" json:"cleanup_after_test,omitempty"`
}

func (x *TestExtensionRevisionRequest) Reset() {
	*x = TestExtensionRevisionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_extensions_v1_extension_testing_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestExtensionRevisionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestExtensionRevisionRequest) ProtoMessage() {}

func (x *TestExtensionRevisionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_extensions_v1_extension_testing_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestExtensionRevisionRequest.ProtoReflect.Descriptor instead.
func (*TestExtensionRevisionRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_extensions_v1_extension_testing_service_proto_rawDescGZIP(), []int{4}
}

func (x *TestExtensionRevisionRequest) GetExtensionData() *ExtensionData {
	if x != nil {
		return x.ExtensionData
	}
	return nil
}

func (x *TestExtensionRevisionRequest) GetCleanupAfterTest() *wrapperspb.BoolValue {
	if x != nil {
		return x.CleanupAfterTest
	}
	return nil
}

type TestExtensionRevisionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TestExtensionRevisionResponse) Reset() {
	*x = TestExtensionRevisionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_extensions_v1_extension_testing_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestExtensionRevisionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestExtensionRevisionResponse) ProtoMessage() {}

func (x *TestExtensionRevisionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_extensions_v1_extension_testing_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestExtensionRevisionResponse.ProtoReflect.Descriptor instead.
func (*TestExtensionRevisionResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_extensions_v1_extension_testing_service_proto_rawDescGZIP(), []int{5}
}

var File_com_coralogix_extensions_v1_extension_testing_service_proto protoreflect.FileDescriptor

var file_com_coralogix_extensions_v1_extension_testing_service_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x2b, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x20, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x51, 0x0a, 0x0e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0d, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x23, 0x0a, 0x21, 0x49,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x4d, 0x0a, 0x1d, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x54, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x20, 0x0a, 0x1e, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0xbb, 0x01, 0x0a, 0x1c, 0x54, 0x65, 0x73, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x51, 0x0a, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0d, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x48, 0x0a, 0x12, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70,
	0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x63,
	0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x41, 0x66, 0x74, 0x65, 0x72, 0x54, 0x65, 0x73, 0x74, 0x22,
	0x1f, 0x0a, 0x1d, 0x54, 0x65, 0x73, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0xb0, 0x05, 0x0a, 0x17, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xe9, 0x01, 0x0a,
	0x19, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x54, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4d, 0xc2, 0xb8, 0x02, 0x1d, 0x0a,
	0x1b, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x20, 0x74, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x20, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x26, 0x3a, 0x01, 0x2a, 0x22, 0x21, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x12, 0xdb, 0x01, 0x0a, 0x16, 0x43, 0x6c, 0x65,
	0x61, 0x6e, 0x75, 0x70, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x3b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c,
	0x65, 0x61, 0x6e, 0x75, 0x70, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x48, 0xc2, 0xb8,
	0x02, 0x1b, 0x0a, 0x19, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x20, 0x74, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x20, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a, 0x22, 0x1e, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x12, 0xca, 0x01, 0x0a, 0x15, 0x54, 0x65, 0x73, 0x74, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3a, 0xc2, 0xb8, 0x02, 0x10, 0x0a, 0x0e, 0x54,
	0x65, 0x73, 0x74, 0x20, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_extensions_v1_extension_testing_service_proto_rawDescOnce sync.Once
	file_com_coralogix_extensions_v1_extension_testing_service_proto_rawDescData = file_com_coralogix_extensions_v1_extension_testing_service_proto_rawDesc
)

func file_com_coralogix_extensions_v1_extension_testing_service_proto_rawDescGZIP() []byte {
	file_com_coralogix_extensions_v1_extension_testing_service_proto_rawDescOnce.Do(func() {
		file_com_coralogix_extensions_v1_extension_testing_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_extensions_v1_extension_testing_service_proto_rawDescData)
	})
	return file_com_coralogix_extensions_v1_extension_testing_service_proto_rawDescData
}

var file_com_coralogix_extensions_v1_extension_testing_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_com_coralogix_extensions_v1_extension_testing_service_proto_goTypes = []any{
	(*InitializeTestingRevisionRequest)(nil),  // 0: com.coralogix.extensions.v1.InitializeTestingRevisionRequest
	(*InitializeTestingRevisionResponse)(nil), // 1: com.coralogix.extensions.v1.InitializeTestingRevisionResponse
	(*CleanupTestingRevisionRequest)(nil),     // 2: com.coralogix.extensions.v1.CleanupTestingRevisionRequest
	(*CleanupTestingRevisionResponse)(nil),    // 3: com.coralogix.extensions.v1.CleanupTestingRevisionResponse
	(*TestExtensionRevisionRequest)(nil),      // 4: com.coralogix.extensions.v1.TestExtensionRevisionRequest
	(*TestExtensionRevisionResponse)(nil),     // 5: com.coralogix.extensions.v1.TestExtensionRevisionResponse
	(*ExtensionData)(nil),                     // 6: com.coralogix.extensions.v1.ExtensionData
	(*wrapperspb.StringValue)(nil),            // 7: google.protobuf.StringValue
	(*wrapperspb.BoolValue)(nil),              // 8: google.protobuf.BoolValue
}
var file_com_coralogix_extensions_v1_extension_testing_service_proto_depIdxs = []int32{
	6, // 0: com.coralogix.extensions.v1.InitializeTestingRevisionRequest.extension_data:type_name -> com.coralogix.extensions.v1.ExtensionData
	7, // 1: com.coralogix.extensions.v1.CleanupTestingRevisionRequest.id:type_name -> google.protobuf.StringValue
	6, // 2: com.coralogix.extensions.v1.TestExtensionRevisionRequest.extension_data:type_name -> com.coralogix.extensions.v1.ExtensionData
	8, // 3: com.coralogix.extensions.v1.TestExtensionRevisionRequest.cleanup_after_test:type_name -> google.protobuf.BoolValue
	0, // 4: com.coralogix.extensions.v1.ExtensionTestingService.InitializeTestingRevision:input_type -> com.coralogix.extensions.v1.InitializeTestingRevisionRequest
	2, // 5: com.coralogix.extensions.v1.ExtensionTestingService.CleanupTestingRevision:input_type -> com.coralogix.extensions.v1.CleanupTestingRevisionRequest
	4, // 6: com.coralogix.extensions.v1.ExtensionTestingService.TestExtensionRevision:input_type -> com.coralogix.extensions.v1.TestExtensionRevisionRequest
	1, // 7: com.coralogix.extensions.v1.ExtensionTestingService.InitializeTestingRevision:output_type -> com.coralogix.extensions.v1.InitializeTestingRevisionResponse
	3, // 8: com.coralogix.extensions.v1.ExtensionTestingService.CleanupTestingRevision:output_type -> com.coralogix.extensions.v1.CleanupTestingRevisionResponse
	5, // 9: com.coralogix.extensions.v1.ExtensionTestingService.TestExtensionRevision:output_type -> com.coralogix.extensions.v1.TestExtensionRevisionResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_com_coralogix_extensions_v1_extension_testing_service_proto_init() }
func file_com_coralogix_extensions_v1_extension_testing_service_proto_init() {
	if File_com_coralogix_extensions_v1_extension_testing_service_proto != nil {
		return
	}
	file_com_coralogix_extensions_v1_audit_log_proto_init()
	file_com_coralogix_extensions_v1_extension_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_extensions_v1_extension_testing_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*InitializeTestingRevisionRequest); i {
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
		file_com_coralogix_extensions_v1_extension_testing_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*InitializeTestingRevisionResponse); i {
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
		file_com_coralogix_extensions_v1_extension_testing_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CleanupTestingRevisionRequest); i {
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
		file_com_coralogix_extensions_v1_extension_testing_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CleanupTestingRevisionResponse); i {
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
		file_com_coralogix_extensions_v1_extension_testing_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*TestExtensionRevisionRequest); i {
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
		file_com_coralogix_extensions_v1_extension_testing_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*TestExtensionRevisionResponse); i {
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
			RawDescriptor: file_com_coralogix_extensions_v1_extension_testing_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogix_extensions_v1_extension_testing_service_proto_goTypes,
		DependencyIndexes: file_com_coralogix_extensions_v1_extension_testing_service_proto_depIdxs,
		MessageInfos:      file_com_coralogix_extensions_v1_extension_testing_service_proto_msgTypes,
	}.Build()
	File_com_coralogix_extensions_v1_extension_testing_service_proto = out.File
	file_com_coralogix_extensions_v1_extension_testing_service_proto_rawDesc = nil
	file_com_coralogix_extensions_v1_extension_testing_service_proto_goTypes = nil
	file_com_coralogix_extensions_v1_extension_testing_service_proto_depIdxs = nil
}
