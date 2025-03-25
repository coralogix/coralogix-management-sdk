// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: com/coralogix/integrations/v1/permission_check_service.proto

package v1

import (
	_ "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/common/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ResourceId struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Id            *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResourceId) Reset() {
	*x = ResourceId{}
	mi := &file_com_coralogix_integrations_v1_permission_check_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceId) ProtoMessage() {}

func (x *ResourceId) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_integrations_v1_permission_check_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceId.ProtoReflect.Descriptor instead.
func (*ResourceId) Descriptor() ([]byte, []int) {
	return file_com_coralogix_integrations_v1_permission_check_service_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceId) GetId() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Id
	}
	return nil
}

type ActionId struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Id            *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ActionId) Reset() {
	*x = ActionId{}
	mi := &file_com_coralogix_integrations_v1_permission_check_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActionId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionId) ProtoMessage() {}

func (x *ActionId) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_integrations_v1_permission_check_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionId.ProtoReflect.Descriptor instead.
func (*ActionId) Descriptor() ([]byte, []int) {
	return file_com_coralogix_integrations_v1_permission_check_service_proto_rawDescGZIP(), []int{1}
}

func (x *ActionId) GetId() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Id
	}
	return nil
}

type ApiKeyValue struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Key           *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ApiKeyValue) Reset() {
	*x = ApiKeyValue{}
	mi := &file_com_coralogix_integrations_v1_permission_check_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApiKeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiKeyValue) ProtoMessage() {}

func (x *ApiKeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_integrations_v1_permission_check_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiKeyValue.ProtoReflect.Descriptor instead.
func (*ApiKeyValue) Descriptor() ([]byte, []int) {
	return file_com_coralogix_integrations_v1_permission_check_service_proto_rawDescGZIP(), []int{2}
}

func (x *ApiKeyValue) GetKey() *wrapperspb.StringValue {
	if x != nil {
		return x.Key
	}
	return nil
}

type ApiKeyId struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Id            *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ApiKeyId) Reset() {
	*x = ApiKeyId{}
	mi := &file_com_coralogix_integrations_v1_permission_check_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApiKeyId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiKeyId) ProtoMessage() {}

func (x *ApiKeyId) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_integrations_v1_permission_check_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiKeyId.ProtoReflect.Descriptor instead.
func (*ApiKeyId) Descriptor() ([]byte, []int) {
	return file_com_coralogix_integrations_v1_permission_check_service_proto_rawDescGZIP(), []int{3}
}

func (x *ApiKeyId) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

type Permission struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Resource      *ResourceId            `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	Action        *ActionId              `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Permission) Reset() {
	*x = Permission{}
	mi := &file_com_coralogix_integrations_v1_permission_check_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_integrations_v1_permission_check_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_com_coralogix_integrations_v1_permission_check_service_proto_rawDescGZIP(), []int{4}
}

func (x *Permission) GetResource() *ResourceId {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *Permission) GetAction() *ActionId {
	if x != nil {
		return x.Action
	}
	return nil
}

type PermissionId struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Id            *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PermissionId) Reset() {
	*x = PermissionId{}
	mi := &file_com_coralogix_integrations_v1_permission_check_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PermissionId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionId) ProtoMessage() {}

func (x *PermissionId) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_integrations_v1_permission_check_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionId.ProtoReflect.Descriptor instead.
func (*PermissionId) Descriptor() ([]byte, []int) {
	return file_com_coralogix_integrations_v1_permission_check_service_proto_rawDescGZIP(), []int{5}
}

func (x *PermissionId) GetId() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Id
	}
	return nil
}

type CheckApiKeyPermissionRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Deprecated: Marked as deprecated in com/coralogix/integrations/v1/permission_check_service.proto.
	Permission    *Permission   `protobuf:"bytes,1,opt,name=permission,proto3" json:"permission,omitempty"`
	ApiKey        *ApiKeyValue  `protobuf:"bytes,2,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	PermissionId  *PermissionId `protobuf:"bytes,3,opt,name=permission_id,json=permissionId,proto3" json:"permission_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckApiKeyPermissionRequest) Reset() {
	*x = CheckApiKeyPermissionRequest{}
	mi := &file_com_coralogix_integrations_v1_permission_check_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckApiKeyPermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckApiKeyPermissionRequest) ProtoMessage() {}

func (x *CheckApiKeyPermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_integrations_v1_permission_check_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckApiKeyPermissionRequest.ProtoReflect.Descriptor instead.
func (*CheckApiKeyPermissionRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_integrations_v1_permission_check_service_proto_rawDescGZIP(), []int{6}
}

// Deprecated: Marked as deprecated in com/coralogix/integrations/v1/permission_check_service.proto.
func (x *CheckApiKeyPermissionRequest) GetPermission() *Permission {
	if x != nil {
		return x.Permission
	}
	return nil
}

func (x *CheckApiKeyPermissionRequest) GetApiKey() *ApiKeyValue {
	if x != nil {
		return x.ApiKey
	}
	return nil
}

func (x *CheckApiKeyPermissionRequest) GetPermissionId() *PermissionId {
	if x != nil {
		return x.PermissionId
	}
	return nil
}

type CheckApiKeyPermissionResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Result:
	//
	//	*CheckApiKeyPermissionResponse_Unauthorized_
	//	*CheckApiKeyPermissionResponse_AuthorizationResult
	Result        isCheckApiKeyPermissionResponse_Result `protobuf_oneof:"result"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckApiKeyPermissionResponse) Reset() {
	*x = CheckApiKeyPermissionResponse{}
	mi := &file_com_coralogix_integrations_v1_permission_check_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckApiKeyPermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckApiKeyPermissionResponse) ProtoMessage() {}

func (x *CheckApiKeyPermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_integrations_v1_permission_check_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckApiKeyPermissionResponse.ProtoReflect.Descriptor instead.
func (*CheckApiKeyPermissionResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_integrations_v1_permission_check_service_proto_rawDescGZIP(), []int{7}
}

func (x *CheckApiKeyPermissionResponse) GetResult() isCheckApiKeyPermissionResponse_Result {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *CheckApiKeyPermissionResponse) GetUnauthorized() *CheckApiKeyPermissionResponse_Unauthorized {
	if x != nil {
		if x, ok := x.Result.(*CheckApiKeyPermissionResponse_Unauthorized_); ok {
			return x.Unauthorized
		}
	}
	return nil
}

func (x *CheckApiKeyPermissionResponse) GetAuthorizationResult() *CheckApiKeyPermissionResponse_KeyAuthorizationResult {
	if x != nil {
		if x, ok := x.Result.(*CheckApiKeyPermissionResponse_AuthorizationResult); ok {
			return x.AuthorizationResult
		}
	}
	return nil
}

type isCheckApiKeyPermissionResponse_Result interface {
	isCheckApiKeyPermissionResponse_Result()
}

type CheckApiKeyPermissionResponse_Unauthorized_ struct {
	Unauthorized *CheckApiKeyPermissionResponse_Unauthorized `protobuf:"bytes,1,opt,name=unauthorized,proto3,oneof"`
}

type CheckApiKeyPermissionResponse_AuthorizationResult struct {
	AuthorizationResult *CheckApiKeyPermissionResponse_KeyAuthorizationResult `protobuf:"bytes,2,opt,name=authorization_result,json=authorizationResult,proto3,oneof"`
}

func (*CheckApiKeyPermissionResponse_Unauthorized_) isCheckApiKeyPermissionResponse_Result() {}

func (*CheckApiKeyPermissionResponse_AuthorizationResult) isCheckApiKeyPermissionResponse_Result() {}

type CheckApiKeyPermissionResponse_KeyAuthorizationResult struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ApiKeyId      *ApiKeyId              `protobuf:"bytes,1,opt,name=api_key_id,json=apiKeyId,proto3" json:"api_key_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckApiKeyPermissionResponse_KeyAuthorizationResult) Reset() {
	*x = CheckApiKeyPermissionResponse_KeyAuthorizationResult{}
	mi := &file_com_coralogix_integrations_v1_permission_check_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckApiKeyPermissionResponse_KeyAuthorizationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckApiKeyPermissionResponse_KeyAuthorizationResult) ProtoMessage() {}

func (x *CheckApiKeyPermissionResponse_KeyAuthorizationResult) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_integrations_v1_permission_check_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckApiKeyPermissionResponse_KeyAuthorizationResult.ProtoReflect.Descriptor instead.
func (*CheckApiKeyPermissionResponse_KeyAuthorizationResult) Descriptor() ([]byte, []int) {
	return file_com_coralogix_integrations_v1_permission_check_service_proto_rawDescGZIP(), []int{7, 0}
}

func (x *CheckApiKeyPermissionResponse_KeyAuthorizationResult) GetApiKeyId() *ApiKeyId {
	if x != nil {
		return x.ApiKeyId
	}
	return nil
}

type CheckApiKeyPermissionResponse_Unauthorized struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckApiKeyPermissionResponse_Unauthorized) Reset() {
	*x = CheckApiKeyPermissionResponse_Unauthorized{}
	mi := &file_com_coralogix_integrations_v1_permission_check_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckApiKeyPermissionResponse_Unauthorized) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckApiKeyPermissionResponse_Unauthorized) ProtoMessage() {}

func (x *CheckApiKeyPermissionResponse_Unauthorized) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_integrations_v1_permission_check_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckApiKeyPermissionResponse_Unauthorized.ProtoReflect.Descriptor instead.
func (*CheckApiKeyPermissionResponse_Unauthorized) Descriptor() ([]byte, []int) {
	return file_com_coralogix_integrations_v1_permission_check_service_proto_rawDescGZIP(), []int{7, 1}
}

var File_com_coralogix_integrations_v1_permission_check_service_proto protoreflect.FileDescriptor

var file_com_coralogix_integrations_v1_permission_check_service_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x27, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x38, 0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3d, 0x0a, 0x0b, 0x41, 0x70,
	0x69, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2e, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x38, 0x0a, 0x08, 0x41, 0x70, 0x69,
	0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x94, 0x01, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3c, 0x0a, 0x0c, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x69, 0x64, 0x22, 0x84, 0x02, 0x0a, 0x1c, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4d, 0x0a, 0x0a, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0a, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x50, 0x0a,
	0x0d, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x52, 0x0c, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22,
	0x96, 0x03, 0x0a, 0x1d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x6f, 0x0a, 0x0c, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x49, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x70, 0x69,
	0x4b, 0x65, 0x79, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x64, 0x48, 0x00, 0x52, 0x0c, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x64, 0x12, 0x88, 0x01, 0x0a, 0x14, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x53, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x13, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x5f, 0x0a,
	0x16, 0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x45, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x5f, 0x6b,
	0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x69, 0x4b,
	0x65, 0x79, 0x49, 0x64, 0x52, 0x08, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x1a, 0x0e,
	0x0a, 0x0c, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x42, 0x08,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xfc, 0x01, 0x0a, 0x16, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0xe1, 0x01, 0x0a, 0x15, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x70, 0x69,
	0x4b, 0x65, 0x79, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4d, 0xc2, 0xb8, 0x02, 0x1d, 0x0a, 0x1b,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20,
	0x66, 0x6f, 0x72, 0x20, 0x61, 0x70, 0x69, 0x2d, 0x6b, 0x65, 0x79, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x26, 0x3a, 0x01, 0x2a, 0x22, 0x21, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2d, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_integrations_v1_permission_check_service_proto_rawDescOnce sync.Once
	file_com_coralogix_integrations_v1_permission_check_service_proto_rawDescData = file_com_coralogix_integrations_v1_permission_check_service_proto_rawDesc
)

func file_com_coralogix_integrations_v1_permission_check_service_proto_rawDescGZIP() []byte {
	file_com_coralogix_integrations_v1_permission_check_service_proto_rawDescOnce.Do(func() {
		file_com_coralogix_integrations_v1_permission_check_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_integrations_v1_permission_check_service_proto_rawDescData)
	})
	return file_com_coralogix_integrations_v1_permission_check_service_proto_rawDescData
}

var file_com_coralogix_integrations_v1_permission_check_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_com_coralogix_integrations_v1_permission_check_service_proto_goTypes = []any{
	(*ResourceId)(nil),                    // 0: com.coralogix.integrations.v1.ResourceId
	(*ActionId)(nil),                      // 1: com.coralogix.integrations.v1.ActionId
	(*ApiKeyValue)(nil),                   // 2: com.coralogix.integrations.v1.ApiKeyValue
	(*ApiKeyId)(nil),                      // 3: com.coralogix.integrations.v1.ApiKeyId
	(*Permission)(nil),                    // 4: com.coralogix.integrations.v1.Permission
	(*PermissionId)(nil),                  // 5: com.coralogix.integrations.v1.PermissionId
	(*CheckApiKeyPermissionRequest)(nil),  // 6: com.coralogix.integrations.v1.CheckApiKeyPermissionRequest
	(*CheckApiKeyPermissionResponse)(nil), // 7: com.coralogix.integrations.v1.CheckApiKeyPermissionResponse
	(*CheckApiKeyPermissionResponse_KeyAuthorizationResult)(nil), // 8: com.coralogix.integrations.v1.CheckApiKeyPermissionResponse.KeyAuthorizationResult
	(*CheckApiKeyPermissionResponse_Unauthorized)(nil),           // 9: com.coralogix.integrations.v1.CheckApiKeyPermissionResponse.Unauthorized
	(*wrapperspb.UInt32Value)(nil),                               // 10: google.protobuf.UInt32Value
	(*wrapperspb.StringValue)(nil),                               // 11: google.protobuf.StringValue
}
var file_com_coralogix_integrations_v1_permission_check_service_proto_depIdxs = []int32{
	10, // 0: com.coralogix.integrations.v1.ResourceId.id:type_name -> google.protobuf.UInt32Value
	10, // 1: com.coralogix.integrations.v1.ActionId.id:type_name -> google.protobuf.UInt32Value
	11, // 2: com.coralogix.integrations.v1.ApiKeyValue.key:type_name -> google.protobuf.StringValue
	11, // 3: com.coralogix.integrations.v1.ApiKeyId.id:type_name -> google.protobuf.StringValue
	0,  // 4: com.coralogix.integrations.v1.Permission.resource:type_name -> com.coralogix.integrations.v1.ResourceId
	1,  // 5: com.coralogix.integrations.v1.Permission.action:type_name -> com.coralogix.integrations.v1.ActionId
	10, // 6: com.coralogix.integrations.v1.PermissionId.id:type_name -> google.protobuf.UInt32Value
	4,  // 7: com.coralogix.integrations.v1.CheckApiKeyPermissionRequest.permission:type_name -> com.coralogix.integrations.v1.Permission
	2,  // 8: com.coralogix.integrations.v1.CheckApiKeyPermissionRequest.api_key:type_name -> com.coralogix.integrations.v1.ApiKeyValue
	5,  // 9: com.coralogix.integrations.v1.CheckApiKeyPermissionRequest.permission_id:type_name -> com.coralogix.integrations.v1.PermissionId
	9,  // 10: com.coralogix.integrations.v1.CheckApiKeyPermissionResponse.unauthorized:type_name -> com.coralogix.integrations.v1.CheckApiKeyPermissionResponse.Unauthorized
	8,  // 11: com.coralogix.integrations.v1.CheckApiKeyPermissionResponse.authorization_result:type_name -> com.coralogix.integrations.v1.CheckApiKeyPermissionResponse.KeyAuthorizationResult
	3,  // 12: com.coralogix.integrations.v1.CheckApiKeyPermissionResponse.KeyAuthorizationResult.api_key_id:type_name -> com.coralogix.integrations.v1.ApiKeyId
	6,  // 13: com.coralogix.integrations.v1.PermissionCheckService.CheckApiKeyPermission:input_type -> com.coralogix.integrations.v1.CheckApiKeyPermissionRequest
	7,  // 14: com.coralogix.integrations.v1.PermissionCheckService.CheckApiKeyPermission:output_type -> com.coralogix.integrations.v1.CheckApiKeyPermissionResponse
	14, // [14:15] is the sub-list for method output_type
	13, // [13:14] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_com_coralogix_integrations_v1_permission_check_service_proto_init() }
func file_com_coralogix_integrations_v1_permission_check_service_proto_init() {
	if File_com_coralogix_integrations_v1_permission_check_service_proto != nil {
		return
	}
	file_com_coralogix_integrations_v1_permission_check_service_proto_msgTypes[7].OneofWrappers = []any{
		(*CheckApiKeyPermissionResponse_Unauthorized_)(nil),
		(*CheckApiKeyPermissionResponse_AuthorizationResult)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_integrations_v1_permission_check_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogix_integrations_v1_permission_check_service_proto_goTypes,
		DependencyIndexes: file_com_coralogix_integrations_v1_permission_check_service_proto_depIdxs,
		MessageInfos:      file_com_coralogix_integrations_v1_permission_check_service_proto_msgTypes,
	}.Build()
	File_com_coralogix_integrations_v1_permission_check_service_proto = out.File
	file_com_coralogix_integrations_v1_permission_check_service_proto_rawDesc = nil
	file_com_coralogix_integrations_v1_permission_check_service_proto_goTypes = nil
	file_com_coralogix_integrations_v1_permission_check_service_proto_depIdxs = nil
}
