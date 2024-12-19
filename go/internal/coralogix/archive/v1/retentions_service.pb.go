// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.26.1
// source: com/coralogix/archive/v1/retentions_service.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
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

type GetRetentionsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRetentionsRequest) Reset() {
	*x = GetRetentionsRequest{}
	mi := &file_com_coralogix_archive_v1_retentions_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRetentionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRetentionsRequest) ProtoMessage() {}

func (x *GetRetentionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v1_retentions_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRetentionsRequest.ProtoReflect.Descriptor instead.
func (*GetRetentionsRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v1_retentions_service_proto_rawDescGZIP(), []int{0}
}

type GetRetentionsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Retentions    []*Retention           `protobuf:"bytes,1,rep,name=retentions,proto3" json:"retentions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRetentionsResponse) Reset() {
	*x = GetRetentionsResponse{}
	mi := &file_com_coralogix_archive_v1_retentions_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRetentionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRetentionsResponse) ProtoMessage() {}

func (x *GetRetentionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v1_retentions_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRetentionsResponse.ProtoReflect.Descriptor instead.
func (*GetRetentionsResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v1_retentions_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetRetentionsResponse) GetRetentions() []*Retention {
	if x != nil {
		return x.Retentions
	}
	return nil
}

type RetentionUpdateElement struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Id            *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RetentionUpdateElement) Reset() {
	*x = RetentionUpdateElement{}
	mi := &file_com_coralogix_archive_v1_retentions_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RetentionUpdateElement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetentionUpdateElement) ProtoMessage() {}

func (x *RetentionUpdateElement) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v1_retentions_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetentionUpdateElement.ProtoReflect.Descriptor instead.
func (*RetentionUpdateElement) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v1_retentions_service_proto_rawDescGZIP(), []int{2}
}

func (x *RetentionUpdateElement) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *RetentionUpdateElement) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

type UpdateRetentionsRequest struct {
	state                   protoimpl.MessageState    `protogen:"open.v1"`
	RetentionUpdateElements []*RetentionUpdateElement `protobuf:"bytes,1,rep,name=retention_update_elements,json=retentionUpdateElements,proto3" json:"retention_update_elements,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *UpdateRetentionsRequest) Reset() {
	*x = UpdateRetentionsRequest{}
	mi := &file_com_coralogix_archive_v1_retentions_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRetentionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRetentionsRequest) ProtoMessage() {}

func (x *UpdateRetentionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v1_retentions_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRetentionsRequest.ProtoReflect.Descriptor instead.
func (*UpdateRetentionsRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v1_retentions_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateRetentionsRequest) GetRetentionUpdateElements() []*RetentionUpdateElement {
	if x != nil {
		return x.RetentionUpdateElements
	}
	return nil
}

type UpdateRetentionsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Retentions    []*Retention           `protobuf:"bytes,1,rep,name=retentions,proto3" json:"retentions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateRetentionsResponse) Reset() {
	*x = UpdateRetentionsResponse{}
	mi := &file_com_coralogix_archive_v1_retentions_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRetentionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRetentionsResponse) ProtoMessage() {}

func (x *UpdateRetentionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v1_retentions_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRetentionsResponse.ProtoReflect.Descriptor instead.
func (*UpdateRetentionsResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v1_retentions_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateRetentionsResponse) GetRetentions() []*Retention {
	if x != nil {
		return x.Retentions
	}
	return nil
}

type ActivateRetentionsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ActivateRetentionsRequest) Reset() {
	*x = ActivateRetentionsRequest{}
	mi := &file_com_coralogix_archive_v1_retentions_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActivateRetentionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateRetentionsRequest) ProtoMessage() {}

func (x *ActivateRetentionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v1_retentions_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateRetentionsRequest.ProtoReflect.Descriptor instead.
func (*ActivateRetentionsRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v1_retentions_service_proto_rawDescGZIP(), []int{5}
}

type ActivateRetentionsResponse struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	ActivateRetentions *wrapperspb.BoolValue  `protobuf:"bytes,1,opt,name=activate_retentions,json=activateRetentions,proto3" json:"activate_retentions,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *ActivateRetentionsResponse) Reset() {
	*x = ActivateRetentionsResponse{}
	mi := &file_com_coralogix_archive_v1_retentions_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActivateRetentionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateRetentionsResponse) ProtoMessage() {}

func (x *ActivateRetentionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v1_retentions_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateRetentionsResponse.ProtoReflect.Descriptor instead.
func (*ActivateRetentionsResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v1_retentions_service_proto_rawDescGZIP(), []int{6}
}

func (x *ActivateRetentionsResponse) GetActivateRetentions() *wrapperspb.BoolValue {
	if x != nil {
		return x.ActivateRetentions
	}
	return nil
}

type GetRetentionsEnabledRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRetentionsEnabledRequest) Reset() {
	*x = GetRetentionsEnabledRequest{}
	mi := &file_com_coralogix_archive_v1_retentions_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRetentionsEnabledRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRetentionsEnabledRequest) ProtoMessage() {}

func (x *GetRetentionsEnabledRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v1_retentions_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRetentionsEnabledRequest.ProtoReflect.Descriptor instead.
func (*GetRetentionsEnabledRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v1_retentions_service_proto_rawDescGZIP(), []int{7}
}

type GetRetentionsEnabledResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EnableTags    *wrapperspb.BoolValue  `protobuf:"bytes,1,opt,name=enable_tags,json=enableTags,proto3" json:"enable_tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRetentionsEnabledResponse) Reset() {
	*x = GetRetentionsEnabledResponse{}
	mi := &file_com_coralogix_archive_v1_retentions_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRetentionsEnabledResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRetentionsEnabledResponse) ProtoMessage() {}

func (x *GetRetentionsEnabledResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v1_retentions_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRetentionsEnabledResponse.ProtoReflect.Descriptor instead.
func (*GetRetentionsEnabledResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v1_retentions_service_proto_rawDescGZIP(), []int{8}
}

func (x *GetRetentionsEnabledResponse) GetEnableTags() *wrapperspb.BoolValue {
	if x != nil {
		return x.EnableTags
	}
	return nil
}

var file_com_coralogix_archive_v1_retentions_service_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         65001,
		Name:          "com.coralogix.archive.v1.validation_pattern",
		Tag:           "bytes,65001,opt,name=validation_pattern",
		Filename:      "com/coralogix/archive/v1/retentions_service.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         65002,
		Name:          "com.coralogix.archive.v1.validation_max_length",
		Tag:           "varint,65002,opt,name=validation_max_length",
		Filename:      "com/coralogix/archive/v1/retentions_service.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         65003,
		Name:          "com.coralogix.archive.v1.validation_encoding",
		Tag:           "bytes,65003,opt,name=validation_encoding",
		Filename:      "com/coralogix/archive/v1/retentions_service.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         65004,
		Name:          "com.coralogix.archive.v1.validation_max_elements",
		Tag:           "varint,65004,opt,name=validation_max_elements",
		Filename:      "com/coralogix/archive/v1/retentions_service.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional string validation_pattern = 65001;
	E_ValidationPattern = &file_com_coralogix_archive_v1_retentions_service_proto_extTypes[0]
	// optional int32 validation_max_length = 65002;
	E_ValidationMaxLength = &file_com_coralogix_archive_v1_retentions_service_proto_extTypes[1]
	// optional string validation_encoding = 65003;
	E_ValidationEncoding = &file_com_coralogix_archive_v1_retentions_service_proto_extTypes[2]
	// optional int32 validation_max_elements = 65004;
	E_ValidationMaxElements = &file_com_coralogix_archive_v1_retentions_service_proto_extTypes[3]
)

var File_com_coralogix_archive_v1_retentions_service_proto protoreflect.FileDescriptor

var file_com_coralogix_archive_v1_retentions_service_proto_rawDesc = []byte{
	0x0a, 0x31, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x74, 0x65, 0x6e,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x28, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f,
	0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5c, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x72,
	0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xfb, 0x01, 0x0a, 0x16, 0x52, 0x65,
	0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x89, 0x01, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x5b, 0xca, 0xde, 0x1f, 0x57, 0x5e, 0x28, 0x5b, 0x61, 0x2d, 0x66, 0x41, 0x2d, 0x46, 0x30, 0x2d,
	0x39, 0x5d, 0x7b, 0x38, 0x7d, 0x5b, 0x2d, 0x5d, 0x5b, 0x61, 0x2d, 0x66, 0x41, 0x2d, 0x46, 0x30,
	0x2d, 0x39, 0x5d, 0x7b, 0x34, 0x7d, 0x5b, 0x2d, 0x5d, 0x5b, 0x61, 0x2d, 0x66, 0x41, 0x2d, 0x46,
	0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x34, 0x7d, 0x5b, 0x2d, 0x5d, 0x5b, 0x61, 0x2d, 0x66, 0x41, 0x2d,
	0x46, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x34, 0x7d, 0x5b, 0x2d, 0x5d, 0x5b, 0x61, 0x2d, 0x66, 0x41,
	0x2d, 0x46, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x31, 0x32, 0x7d, 0x29, 0x24, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x55, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x23, 0xca, 0xde,
	0x1f, 0x10, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5f, 0x2d, 0x5d,
	0x2b, 0x24, 0xd0, 0xde, 0x1f, 0xff, 0x01, 0xda, 0xde, 0x1f, 0x06, 0x6c, 0x61, 0x74, 0x69, 0x6e,
	0x31, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x8d, 0x01, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x72, 0x0a, 0x19, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x04, 0xe0, 0xde, 0x1f, 0x04, 0x52, 0x17,
	0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x5f, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x72, 0x65,
	0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x1b, 0x0a, 0x19, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x69, 0x0a, 0x1a, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x13, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f,
	0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x12, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x1d, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x5b, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3b, 0x0a, 0x0b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0a, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x61, 0x67, 0x73, 0x32, 0x88, 0x05, 0x0a,
	0x11, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x86, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0xc2, 0xb8, 0x02, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x20, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x97, 0x01, 0x0a, 0x10,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0xc2, 0xb8, 0x02, 0x18, 0x0a, 0x16, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x20,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x9e, 0x01, 0x0a, 0x12, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x33, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0xc2, 0xb8, 0x02, 0x19, 0x0a, 0x17, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x20, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f,
	0x6e, 0x20, 0x74, 0x61, 0x67, 0x73, 0x12, 0xae, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12,
	0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27,
	0xc2, 0xb8, 0x02, 0x23, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x20, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74,
	0x69, 0x6f, 0x6e, 0x20, 0x74, 0x61, 0x67, 0x73, 0x20, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x20, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x4e, 0x0a, 0x12, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x1d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe9, 0xfb, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x3a, 0x53, 0x0a, 0x15, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xea, 0xfb, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x78, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x3a, 0x50, 0x0a, 0x13,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xeb, 0xfb, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x57,
	0x0a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x78,
	0x5f, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xec, 0xfb, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x15, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x78, 0x45,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_archive_v1_retentions_service_proto_rawDescOnce sync.Once
	file_com_coralogix_archive_v1_retentions_service_proto_rawDescData = file_com_coralogix_archive_v1_retentions_service_proto_rawDesc
)

func file_com_coralogix_archive_v1_retentions_service_proto_rawDescGZIP() []byte {
	file_com_coralogix_archive_v1_retentions_service_proto_rawDescOnce.Do(func() {
		file_com_coralogix_archive_v1_retentions_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_archive_v1_retentions_service_proto_rawDescData)
	})
	return file_com_coralogix_archive_v1_retentions_service_proto_rawDescData
}

var file_com_coralogix_archive_v1_retentions_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_com_coralogix_archive_v1_retentions_service_proto_goTypes = []any{
	(*GetRetentionsRequest)(nil),         // 0: com.coralogix.archive.v1.GetRetentionsRequest
	(*GetRetentionsResponse)(nil),        // 1: com.coralogix.archive.v1.GetRetentionsResponse
	(*RetentionUpdateElement)(nil),       // 2: com.coralogix.archive.v1.RetentionUpdateElement
	(*UpdateRetentionsRequest)(nil),      // 3: com.coralogix.archive.v1.UpdateRetentionsRequest
	(*UpdateRetentionsResponse)(nil),     // 4: com.coralogix.archive.v1.UpdateRetentionsResponse
	(*ActivateRetentionsRequest)(nil),    // 5: com.coralogix.archive.v1.ActivateRetentionsRequest
	(*ActivateRetentionsResponse)(nil),   // 6: com.coralogix.archive.v1.ActivateRetentionsResponse
	(*GetRetentionsEnabledRequest)(nil),  // 7: com.coralogix.archive.v1.GetRetentionsEnabledRequest
	(*GetRetentionsEnabledResponse)(nil), // 8: com.coralogix.archive.v1.GetRetentionsEnabledResponse
	(*Retention)(nil),                    // 9: com.coralogix.archive.v1.Retention
	(*wrapperspb.StringValue)(nil),       // 10: google.protobuf.StringValue
	(*wrapperspb.BoolValue)(nil),         // 11: google.protobuf.BoolValue
	(*descriptorpb.FieldOptions)(nil),    // 12: google.protobuf.FieldOptions
}
var file_com_coralogix_archive_v1_retentions_service_proto_depIdxs = []int32{
	9,  // 0: com.coralogix.archive.v1.GetRetentionsResponse.retentions:type_name -> com.coralogix.archive.v1.Retention
	10, // 1: com.coralogix.archive.v1.RetentionUpdateElement.id:type_name -> google.protobuf.StringValue
	10, // 2: com.coralogix.archive.v1.RetentionUpdateElement.name:type_name -> google.protobuf.StringValue
	2,  // 3: com.coralogix.archive.v1.UpdateRetentionsRequest.retention_update_elements:type_name -> com.coralogix.archive.v1.RetentionUpdateElement
	9,  // 4: com.coralogix.archive.v1.UpdateRetentionsResponse.retentions:type_name -> com.coralogix.archive.v1.Retention
	11, // 5: com.coralogix.archive.v1.ActivateRetentionsResponse.activate_retentions:type_name -> google.protobuf.BoolValue
	11, // 6: com.coralogix.archive.v1.GetRetentionsEnabledResponse.enable_tags:type_name -> google.protobuf.BoolValue
	12, // 7: com.coralogix.archive.v1.validation_pattern:extendee -> google.protobuf.FieldOptions
	12, // 8: com.coralogix.archive.v1.validation_max_length:extendee -> google.protobuf.FieldOptions
	12, // 9: com.coralogix.archive.v1.validation_encoding:extendee -> google.protobuf.FieldOptions
	12, // 10: com.coralogix.archive.v1.validation_max_elements:extendee -> google.protobuf.FieldOptions
	0,  // 11: com.coralogix.archive.v1.RetentionsService.GetRetentions:input_type -> com.coralogix.archive.v1.GetRetentionsRequest
	3,  // 12: com.coralogix.archive.v1.RetentionsService.UpdateRetentions:input_type -> com.coralogix.archive.v1.UpdateRetentionsRequest
	5,  // 13: com.coralogix.archive.v1.RetentionsService.ActivateRetentions:input_type -> com.coralogix.archive.v1.ActivateRetentionsRequest
	7,  // 14: com.coralogix.archive.v1.RetentionsService.GetRetentionsEnabled:input_type -> com.coralogix.archive.v1.GetRetentionsEnabledRequest
	1,  // 15: com.coralogix.archive.v1.RetentionsService.GetRetentions:output_type -> com.coralogix.archive.v1.GetRetentionsResponse
	4,  // 16: com.coralogix.archive.v1.RetentionsService.UpdateRetentions:output_type -> com.coralogix.archive.v1.UpdateRetentionsResponse
	6,  // 17: com.coralogix.archive.v1.RetentionsService.ActivateRetentions:output_type -> com.coralogix.archive.v1.ActivateRetentionsResponse
	8,  // 18: com.coralogix.archive.v1.RetentionsService.GetRetentionsEnabled:output_type -> com.coralogix.archive.v1.GetRetentionsEnabledResponse
	15, // [15:19] is the sub-list for method output_type
	11, // [11:15] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	7,  // [7:11] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_com_coralogix_archive_v1_retentions_service_proto_init() }
func file_com_coralogix_archive_v1_retentions_service_proto_init() {
	if File_com_coralogix_archive_v1_retentions_service_proto != nil {
		return
	}
	file_com_coralogix_archive_v1_audit_log_proto_init()
	file_com_coralogix_archive_v1_retention_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_archive_v1_retentions_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 4,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogix_archive_v1_retentions_service_proto_goTypes,
		DependencyIndexes: file_com_coralogix_archive_v1_retentions_service_proto_depIdxs,
		MessageInfos:      file_com_coralogix_archive_v1_retentions_service_proto_msgTypes,
		ExtensionInfos:    file_com_coralogix_archive_v1_retentions_service_proto_extTypes,
	}.Build()
	File_com_coralogix_archive_v1_retentions_service_proto = out.File
	file_com_coralogix_archive_v1_retentions_service_proto_rawDesc = nil
	file_com_coralogix_archive_v1_retentions_service_proto_goTypes = nil
	file_com_coralogix_archive_v1_retentions_service_proto_depIdxs = nil
}
