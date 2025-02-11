// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: com/coralogixapis/events2metrics/v2/events2metrics_internal_service.proto

package v2

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

type ListE2MRequestInternal struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListE2MRequestInternal) Reset() {
	*x = ListE2MRequestInternal{}
	mi := &file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListE2MRequestInternal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListE2MRequestInternal) ProtoMessage() {}

func (x *ListE2MRequestInternal) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListE2MRequestInternal.ProtoReflect.Descriptor instead.
func (*ListE2MRequestInternal) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_rawDescGZIP(), []int{0}
}

type ListE2MResponseInternal struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	E2M           []*E2M                 `protobuf:"bytes,1,rep,name=e2m,proto3" json:"e2m,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListE2MResponseInternal) Reset() {
	*x = ListE2MResponseInternal{}
	mi := &file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListE2MResponseInternal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListE2MResponseInternal) ProtoMessage() {}

func (x *ListE2MResponseInternal) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListE2MResponseInternal.ProtoReflect.Descriptor instead.
func (*ListE2MResponseInternal) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListE2MResponseInternal) GetE2M() []*E2M {
	if x != nil {
		return x.E2M
	}
	return nil
}

type CreateE2MRequestInternal struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	E2M           *E2MCreateParams       `protobuf:"bytes,1,opt,name=e2m,proto3" json:"e2m,omitempty"`
	IsInternal    *wrapperspb.BoolValue  `protobuf:"bytes,2,opt,name=is_internal,json=isInternal,proto3" json:"is_internal,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateE2MRequestInternal) Reset() {
	*x = CreateE2MRequestInternal{}
	mi := &file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateE2MRequestInternal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateE2MRequestInternal) ProtoMessage() {}

func (x *CreateE2MRequestInternal) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateE2MRequestInternal.ProtoReflect.Descriptor instead.
func (*CreateE2MRequestInternal) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateE2MRequestInternal) GetE2M() *E2MCreateParams {
	if x != nil {
		return x.E2M
	}
	return nil
}

func (x *CreateE2MRequestInternal) GetIsInternal() *wrapperspb.BoolValue {
	if x != nil {
		return x.IsInternal
	}
	return nil
}

type CreateE2MResponseInternal struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	E2M           *E2M                   `protobuf:"bytes,1,opt,name=e2m,proto3" json:"e2m,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateE2MResponseInternal) Reset() {
	*x = CreateE2MResponseInternal{}
	mi := &file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateE2MResponseInternal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateE2MResponseInternal) ProtoMessage() {}

func (x *CreateE2MResponseInternal) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateE2MResponseInternal.ProtoReflect.Descriptor instead.
func (*CreateE2MResponseInternal) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateE2MResponseInternal) GetE2M() *E2M {
	if x != nil {
		return x.E2M
	}
	return nil
}

type ReplaceE2MRequestInternal struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	E2M           *E2M                   `protobuf:"bytes,1,opt,name=e2m,proto3" json:"e2m,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReplaceE2MRequestInternal) Reset() {
	*x = ReplaceE2MRequestInternal{}
	mi := &file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReplaceE2MRequestInternal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplaceE2MRequestInternal) ProtoMessage() {}

func (x *ReplaceE2MRequestInternal) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplaceE2MRequestInternal.ProtoReflect.Descriptor instead.
func (*ReplaceE2MRequestInternal) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_rawDescGZIP(), []int{4}
}

func (x *ReplaceE2MRequestInternal) GetE2M() *E2M {
	if x != nil {
		return x.E2M
	}
	return nil
}

type ReplaceE2MResponseInternal struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	E2M           *E2M                   `protobuf:"bytes,1,opt,name=e2m,proto3" json:"e2m,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReplaceE2MResponseInternal) Reset() {
	*x = ReplaceE2MResponseInternal{}
	mi := &file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReplaceE2MResponseInternal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplaceE2MResponseInternal) ProtoMessage() {}

func (x *ReplaceE2MResponseInternal) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplaceE2MResponseInternal.ProtoReflect.Descriptor instead.
func (*ReplaceE2MResponseInternal) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_rawDescGZIP(), []int{5}
}

func (x *ReplaceE2MResponseInternal) GetE2M() *E2M {
	if x != nil {
		return x.E2M
	}
	return nil
}

type DeleteE2MRequestInternal struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Id            *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteE2MRequestInternal) Reset() {
	*x = DeleteE2MRequestInternal{}
	mi := &file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteE2MRequestInternal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteE2MRequestInternal) ProtoMessage() {}

func (x *DeleteE2MRequestInternal) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteE2MRequestInternal.ProtoReflect.Descriptor instead.
func (*DeleteE2MRequestInternal) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteE2MRequestInternal) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

type DeleteE2MResponseInternal struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Id            *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteE2MResponseInternal) Reset() {
	*x = DeleteE2MResponseInternal{}
	mi := &file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteE2MResponseInternal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteE2MResponseInternal) ProtoMessage() {}

func (x *DeleteE2MResponseInternal) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteE2MResponseInternal.ProtoReflect.Descriptor instead.
func (*DeleteE2MResponseInternal) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteE2MResponseInternal) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

var File_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto protoreflect.FileDescriptor

var file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_rawDesc = []byte{
	0x0a, 0x49, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32,
	0x1a, 0x43, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x18,
	0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x32, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x22, 0x55, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74,
	0x45, 0x32, 0x4d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x12, 0x3a, 0x0a, 0x03, 0x65, 0x32, 0x6d, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x32, 0x4d, 0x52, 0x03, 0x65, 0x32, 0x6d, 0x22,
	0x9f, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x32, 0x4d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x46, 0x0a, 0x03,
	0x65, 0x32, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x32, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e,
	0x45, 0x32, 0x4d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52,
	0x03, 0x65, 0x32, 0x6d, 0x12, 0x3b, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x69, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x22, 0x57, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x32, 0x4d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x3a,
	0x0a, 0x03, 0x65, 0x32, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x45, 0x32, 0x4d, 0x52, 0x03, 0x65, 0x32, 0x6d, 0x22, 0x57, 0x0a, 0x19, 0x52, 0x65,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x45, 0x32, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x3a, 0x0a, 0x03, 0x65, 0x32, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x32, 0x4d, 0x52, 0x03,
	0x65, 0x32, 0x6d, 0x22, 0x58, 0x0a, 0x1a, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x45, 0x32,
	0x4d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x12, 0x3a, 0x0a, 0x03, 0x65, 0x32, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x32, 0x4d, 0x52, 0x03, 0x65, 0x32, 0x6d, 0x22, 0x48, 0x0a,
	0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x32, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x02, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x45, 0x32, 0x4d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02,
	0x69, 0x64, 0x32, 0xd6, 0x05, 0x0a, 0x1c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0xa9, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x32, 0x4d, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x3b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x32, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x32, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x1a, 0x3c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45,
	0x32, 0x4d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x22, 0x1b, 0xc2, 0xb8, 0x02, 0x17, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x61,
	0x6c, 0x6c, 0x20, 0x45, 0x32, 0x4d, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12,
	0xa9, 0x01, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x45, 0x32, 0x4d, 0x12, 0x3e,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x45, 0x32, 0x4d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x1a, 0x3f,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x45, 0x32, 0x4d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x22,
	0x1a, 0xc2, 0xb8, 0x02, 0x16, 0x0a, 0x14, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x20, 0x45,
	0x32, 0x4d, 0x20, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0xad, 0x01, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x32, 0x4d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x12, 0x3d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x32,
	0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x1a, 0x3e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x32, 0x4d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x22, 0x19, 0xc2, 0xb8, 0x02, 0x15, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x45,
	0x32, 0x4d, 0x20, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0xad, 0x01, 0x0a, 0x11,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x32, 0x4d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x12, 0x3d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x32,
	0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x1a, 0x3e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x32, 0x4d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x22, 0x19, 0xc2, 0xb8, 0x02, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x45,
	0x32, 0x4d, 0x20, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_rawDescOnce sync.Once
	file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_rawDescData = file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_rawDesc
)

func file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_rawDescData)
	})
	return file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_rawDescData
}

var file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_goTypes = []any{
	(*ListE2MRequestInternal)(nil),     // 0: com.coralogixapis.events2metrics.v2.ListE2MRequestInternal
	(*ListE2MResponseInternal)(nil),    // 1: com.coralogixapis.events2metrics.v2.ListE2MResponseInternal
	(*CreateE2MRequestInternal)(nil),   // 2: com.coralogixapis.events2metrics.v2.CreateE2MRequestInternal
	(*CreateE2MResponseInternal)(nil),  // 3: com.coralogixapis.events2metrics.v2.CreateE2MResponseInternal
	(*ReplaceE2MRequestInternal)(nil),  // 4: com.coralogixapis.events2metrics.v2.ReplaceE2MRequestInternal
	(*ReplaceE2MResponseInternal)(nil), // 5: com.coralogixapis.events2metrics.v2.ReplaceE2MResponseInternal
	(*DeleteE2MRequestInternal)(nil),   // 6: com.coralogixapis.events2metrics.v2.DeleteE2MRequestInternal
	(*DeleteE2MResponseInternal)(nil),  // 7: com.coralogixapis.events2metrics.v2.DeleteE2MResponseInternal
	(*E2M)(nil),                        // 8: com.coralogixapis.events2metrics.v2.E2M
	(*E2MCreateParams)(nil),            // 9: com.coralogixapis.events2metrics.v2.E2MCreateParams
	(*wrapperspb.BoolValue)(nil),       // 10: google.protobuf.BoolValue
	(*wrapperspb.StringValue)(nil),     // 11: google.protobuf.StringValue
}
var file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_depIdxs = []int32{
	8,  // 0: com.coralogixapis.events2metrics.v2.ListE2MResponseInternal.e2m:type_name -> com.coralogixapis.events2metrics.v2.E2M
	9,  // 1: com.coralogixapis.events2metrics.v2.CreateE2MRequestInternal.e2m:type_name -> com.coralogixapis.events2metrics.v2.E2MCreateParams
	10, // 2: com.coralogixapis.events2metrics.v2.CreateE2MRequestInternal.is_internal:type_name -> google.protobuf.BoolValue
	8,  // 3: com.coralogixapis.events2metrics.v2.CreateE2MResponseInternal.e2m:type_name -> com.coralogixapis.events2metrics.v2.E2M
	8,  // 4: com.coralogixapis.events2metrics.v2.ReplaceE2MRequestInternal.e2m:type_name -> com.coralogixapis.events2metrics.v2.E2M
	8,  // 5: com.coralogixapis.events2metrics.v2.ReplaceE2MResponseInternal.e2m:type_name -> com.coralogixapis.events2metrics.v2.E2M
	11, // 6: com.coralogixapis.events2metrics.v2.DeleteE2MRequestInternal.id:type_name -> google.protobuf.StringValue
	11, // 7: com.coralogixapis.events2metrics.v2.DeleteE2MResponseInternal.id:type_name -> google.protobuf.StringValue
	0,  // 8: com.coralogixapis.events2metrics.v2.Events2MetricInternalService.ListE2MInternal:input_type -> com.coralogixapis.events2metrics.v2.ListE2MRequestInternal
	4,  // 9: com.coralogixapis.events2metrics.v2.Events2MetricInternalService.ReplaceE2M:input_type -> com.coralogixapis.events2metrics.v2.ReplaceE2MRequestInternal
	2,  // 10: com.coralogixapis.events2metrics.v2.Events2MetricInternalService.CreateE2MInternal:input_type -> com.coralogixapis.events2metrics.v2.CreateE2MRequestInternal
	6,  // 11: com.coralogixapis.events2metrics.v2.Events2MetricInternalService.DeleteE2MInternal:input_type -> com.coralogixapis.events2metrics.v2.DeleteE2MRequestInternal
	1,  // 12: com.coralogixapis.events2metrics.v2.Events2MetricInternalService.ListE2MInternal:output_type -> com.coralogixapis.events2metrics.v2.ListE2MResponseInternal
	5,  // 13: com.coralogixapis.events2metrics.v2.Events2MetricInternalService.ReplaceE2M:output_type -> com.coralogixapis.events2metrics.v2.ReplaceE2MResponseInternal
	3,  // 14: com.coralogixapis.events2metrics.v2.Events2MetricInternalService.CreateE2MInternal:output_type -> com.coralogixapis.events2metrics.v2.CreateE2MResponseInternal
	7,  // 15: com.coralogixapis.events2metrics.v2.Events2MetricInternalService.DeleteE2MInternal:output_type -> com.coralogixapis.events2metrics.v2.DeleteE2MResponseInternal
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_init() }
func file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_init() {
	if File_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto != nil {
		return
	}
	file_com_coralogixapis_events2metrics_v2_events2metrics_definition_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto = out.File
	file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_rawDesc = nil
	file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_goTypes = nil
	file_com_coralogixapis_events2metrics_v2_events2metrics_internal_service_proto_depIdxs = nil
}
