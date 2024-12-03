// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: com/coralogix/dataprime/v1/object_store_warning.proto

package v1

import (
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

type ObjectStoreWarning struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to WarningType:
	//
	//	*ObjectStoreWarning_MetastoreWarning_
	//	*ObjectStoreWarning_BlockWarning_
	WarningType isObjectStoreWarning_WarningType `protobuf_oneof:"warning_type"`
}

func (x *ObjectStoreWarning) Reset() {
	*x = ObjectStoreWarning{}
	mi := &file_com_coralogix_dataprime_v1_object_store_warning_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ObjectStoreWarning) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectStoreWarning) ProtoMessage() {}

func (x *ObjectStoreWarning) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_object_store_warning_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectStoreWarning.ProtoReflect.Descriptor instead.
func (*ObjectStoreWarning) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_object_store_warning_proto_rawDescGZIP(), []int{0}
}

func (m *ObjectStoreWarning) GetWarningType() isObjectStoreWarning_WarningType {
	if m != nil {
		return m.WarningType
	}
	return nil
}

func (x *ObjectStoreWarning) GetMetastoreWarning() *ObjectStoreWarning_MetastoreWarning {
	if x, ok := x.GetWarningType().(*ObjectStoreWarning_MetastoreWarning_); ok {
		return x.MetastoreWarning
	}
	return nil
}

func (x *ObjectStoreWarning) GetBlockWarning() *ObjectStoreWarning_BlockWarning {
	if x, ok := x.GetWarningType().(*ObjectStoreWarning_BlockWarning_); ok {
		return x.BlockWarning
	}
	return nil
}

type isObjectStoreWarning_WarningType interface {
	isObjectStoreWarning_WarningType()
}

type ObjectStoreWarning_MetastoreWarning_ struct {
	MetastoreWarning *ObjectStoreWarning_MetastoreWarning `protobuf:"bytes,1,opt,name=metastore_warning,json=metastoreWarning,proto3,oneof"`
}

type ObjectStoreWarning_BlockWarning_ struct {
	BlockWarning *ObjectStoreWarning_BlockWarning `protobuf:"bytes,2,opt,name=block_warning,json=blockWarning,proto3,oneof"`
}

func (*ObjectStoreWarning_MetastoreWarning_) isObjectStoreWarning_WarningType() {}

func (*ObjectStoreWarning_BlockWarning_) isObjectStoreWarning_WarningType() {}

type ObjectStoreWarning_MetastoreWarning struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to WarningType:
	//
	//	*ObjectStoreWarning_MetastoreWarning_NoFilesFoundWarning
	WarningType isObjectStoreWarning_MetastoreWarning_WarningType `protobuf_oneof:"warning_type"`
}

func (x *ObjectStoreWarning_MetastoreWarning) Reset() {
	*x = ObjectStoreWarning_MetastoreWarning{}
	mi := &file_com_coralogix_dataprime_v1_object_store_warning_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ObjectStoreWarning_MetastoreWarning) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectStoreWarning_MetastoreWarning) ProtoMessage() {}

func (x *ObjectStoreWarning_MetastoreWarning) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_object_store_warning_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectStoreWarning_MetastoreWarning.ProtoReflect.Descriptor instead.
func (*ObjectStoreWarning_MetastoreWarning) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_object_store_warning_proto_rawDescGZIP(), []int{0, 0}
}

func (m *ObjectStoreWarning_MetastoreWarning) GetWarningType() isObjectStoreWarning_MetastoreWarning_WarningType {
	if m != nil {
		return m.WarningType
	}
	return nil
}

func (x *ObjectStoreWarning_MetastoreWarning) GetNoFilesFoundWarning() *ObjectStoreWarning_MetastoreWarning_NoFilesFound {
	if x, ok := x.GetWarningType().(*ObjectStoreWarning_MetastoreWarning_NoFilesFoundWarning); ok {
		return x.NoFilesFoundWarning
	}
	return nil
}

type isObjectStoreWarning_MetastoreWarning_WarningType interface {
	isObjectStoreWarning_MetastoreWarning_WarningType()
}

type ObjectStoreWarning_MetastoreWarning_NoFilesFoundWarning struct {
	NoFilesFoundWarning *ObjectStoreWarning_MetastoreWarning_NoFilesFound `protobuf:"bytes,1,opt,name=no_files_found_warning,json=noFilesFoundWarning,proto3,oneof"`
}

func (*ObjectStoreWarning_MetastoreWarning_NoFilesFoundWarning) isObjectStoreWarning_MetastoreWarning_WarningType() {
}

type ObjectStoreWarning_BlockWarning struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to WarningType:
	//
	//	*ObjectStoreWarning_BlockWarning_FileNotFoundWarning
	//	*ObjectStoreWarning_BlockWarning_AccessDeniedWarning
	//	*ObjectStoreWarning_BlockWarning_ReadFailedWarning
	WarningType isObjectStoreWarning_BlockWarning_WarningType `protobuf_oneof:"warning_type"`
}

func (x *ObjectStoreWarning_BlockWarning) Reset() {
	*x = ObjectStoreWarning_BlockWarning{}
	mi := &file_com_coralogix_dataprime_v1_object_store_warning_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ObjectStoreWarning_BlockWarning) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectStoreWarning_BlockWarning) ProtoMessage() {}

func (x *ObjectStoreWarning_BlockWarning) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_object_store_warning_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectStoreWarning_BlockWarning.ProtoReflect.Descriptor instead.
func (*ObjectStoreWarning_BlockWarning) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_object_store_warning_proto_rawDescGZIP(), []int{0, 1}
}

func (m *ObjectStoreWarning_BlockWarning) GetWarningType() isObjectStoreWarning_BlockWarning_WarningType {
	if m != nil {
		return m.WarningType
	}
	return nil
}

func (x *ObjectStoreWarning_BlockWarning) GetFileNotFoundWarning() *ObjectStoreWarning_BlockWarning_FileNotFound {
	if x, ok := x.GetWarningType().(*ObjectStoreWarning_BlockWarning_FileNotFoundWarning); ok {
		return x.FileNotFoundWarning
	}
	return nil
}

func (x *ObjectStoreWarning_BlockWarning) GetAccessDeniedWarning() *ObjectStoreWarning_BlockWarning_AccessDenied {
	if x, ok := x.GetWarningType().(*ObjectStoreWarning_BlockWarning_AccessDeniedWarning); ok {
		return x.AccessDeniedWarning
	}
	return nil
}

func (x *ObjectStoreWarning_BlockWarning) GetReadFailedWarning() *ObjectStoreWarning_BlockWarning_ReadFailed {
	if x, ok := x.GetWarningType().(*ObjectStoreWarning_BlockWarning_ReadFailedWarning); ok {
		return x.ReadFailedWarning
	}
	return nil
}

type isObjectStoreWarning_BlockWarning_WarningType interface {
	isObjectStoreWarning_BlockWarning_WarningType()
}

type ObjectStoreWarning_BlockWarning_FileNotFoundWarning struct {
	FileNotFoundWarning *ObjectStoreWarning_BlockWarning_FileNotFound `protobuf:"bytes,1,opt,name=file_not_found_warning,json=fileNotFoundWarning,proto3,oneof"`
}

type ObjectStoreWarning_BlockWarning_AccessDeniedWarning struct {
	AccessDeniedWarning *ObjectStoreWarning_BlockWarning_AccessDenied `protobuf:"bytes,2,opt,name=access_denied_warning,json=accessDeniedWarning,proto3,oneof"`
}

type ObjectStoreWarning_BlockWarning_ReadFailedWarning struct {
	ReadFailedWarning *ObjectStoreWarning_BlockWarning_ReadFailed `protobuf:"bytes,3,opt,name=read_failed_warning,json=readFailedWarning,proto3,oneof"`
}

func (*ObjectStoreWarning_BlockWarning_FileNotFoundWarning) isObjectStoreWarning_BlockWarning_WarningType() {
}

func (*ObjectStoreWarning_BlockWarning_AccessDeniedWarning) isObjectStoreWarning_BlockWarning_WarningType() {
}

func (*ObjectStoreWarning_BlockWarning_ReadFailedWarning) isObjectStoreWarning_BlockWarning_WarningType() {
}

type ObjectStoreWarning_MetastoreWarning_NoFilesFound struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ObjectStoreWarning_MetastoreWarning_NoFilesFound) Reset() {
	*x = ObjectStoreWarning_MetastoreWarning_NoFilesFound{}
	mi := &file_com_coralogix_dataprime_v1_object_store_warning_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ObjectStoreWarning_MetastoreWarning_NoFilesFound) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectStoreWarning_MetastoreWarning_NoFilesFound) ProtoMessage() {}

func (x *ObjectStoreWarning_MetastoreWarning_NoFilesFound) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_object_store_warning_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectStoreWarning_MetastoreWarning_NoFilesFound.ProtoReflect.Descriptor instead.
func (*ObjectStoreWarning_MetastoreWarning_NoFilesFound) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_object_store_warning_proto_rawDescGZIP(), []int{0, 0, 0}
}

type ObjectStoreWarning_BlockWarning_FileNotFound struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ObjectStoreWarning_BlockWarning_FileNotFound) Reset() {
	*x = ObjectStoreWarning_BlockWarning_FileNotFound{}
	mi := &file_com_coralogix_dataprime_v1_object_store_warning_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ObjectStoreWarning_BlockWarning_FileNotFound) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectStoreWarning_BlockWarning_FileNotFound) ProtoMessage() {}

func (x *ObjectStoreWarning_BlockWarning_FileNotFound) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_object_store_warning_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectStoreWarning_BlockWarning_FileNotFound.ProtoReflect.Descriptor instead.
func (*ObjectStoreWarning_BlockWarning_FileNotFound) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_object_store_warning_proto_rawDescGZIP(), []int{0, 1, 0}
}

type ObjectStoreWarning_BlockWarning_AccessDenied struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ObjectStoreWarning_BlockWarning_AccessDenied) Reset() {
	*x = ObjectStoreWarning_BlockWarning_AccessDenied{}
	mi := &file_com_coralogix_dataprime_v1_object_store_warning_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ObjectStoreWarning_BlockWarning_AccessDenied) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectStoreWarning_BlockWarning_AccessDenied) ProtoMessage() {}

func (x *ObjectStoreWarning_BlockWarning_AccessDenied) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_object_store_warning_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectStoreWarning_BlockWarning_AccessDenied.ProtoReflect.Descriptor instead.
func (*ObjectStoreWarning_BlockWarning_AccessDenied) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_object_store_warning_proto_rawDescGZIP(), []int{0, 1, 1}
}

type ObjectStoreWarning_BlockWarning_ReadFailed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ObjectStoreWarning_BlockWarning_ReadFailed) Reset() {
	*x = ObjectStoreWarning_BlockWarning_ReadFailed{}
	mi := &file_com_coralogix_dataprime_v1_object_store_warning_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ObjectStoreWarning_BlockWarning_ReadFailed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectStoreWarning_BlockWarning_ReadFailed) ProtoMessage() {}

func (x *ObjectStoreWarning_BlockWarning_ReadFailed) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_object_store_warning_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectStoreWarning_BlockWarning_ReadFailed.ProtoReflect.Descriptor instead.
func (*ObjectStoreWarning_BlockWarning_ReadFailed) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_object_store_warning_proto_rawDescGZIP(), []int{0, 1, 2}
}

var File_com_coralogix_dataprime_v1_object_store_warning_proto protoreflect.FileDescriptor

var file_com_coralogix_dataprime_v1_object_store_warning_proto_rawDesc = []byte{
	0x0a, 0x35, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65,
	0x2e, 0x76, 0x31, 0x22, 0xfd, 0x06, 0x0a, 0x12, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x6e, 0x0a, 0x11, 0x6d, 0x65,
	0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x57, 0x61,
	0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x57,
	0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x10, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x62, 0x0a, 0x0d, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x3b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e,
	0x67, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x48, 0x00,
	0x52, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x1a, 0xb8,
	0x01, 0x0a, 0x10, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x57, 0x61, 0x72, 0x6e,
	0x69, 0x6e, 0x67, 0x12, 0x83, 0x01, 0x0a, 0x16, 0x6e, 0x6f, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x4c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x57, 0x61, 0x72,
	0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x57, 0x61,
	0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4e, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x46, 0x6f, 0x75,
	0x6e, 0x64, 0x48, 0x00, 0x52, 0x13, 0x6e, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x46, 0x6f, 0x75,
	0x6e, 0x64, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x1a, 0x0e, 0x0a, 0x0c, 0x4e, 0x6f, 0x46,
	0x69, 0x6c, 0x65, 0x73, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x77, 0x61, 0x72,
	0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x1a, 0xc7, 0x03, 0x0a, 0x0c, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x7f, 0x0a, 0x16, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x77, 0x61, 0x72,
	0x6e, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70,
	0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x46,
	0x6f, 0x75, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x13, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x46,
	0x6f, 0x75, 0x6e, 0x64, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x7e, 0x0a, 0x15, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x64, 0x65, 0x6e, 0x69, 0x65, 0x64, 0x5f, 0x77, 0x61, 0x72,
	0x6e, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70,
	0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x44, 0x65,
	0x6e, 0x69, 0x65, 0x64, 0x48, 0x00, 0x52, 0x13, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x44, 0x65,
	0x6e, 0x69, 0x65, 0x64, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x78, 0x0a, 0x13, 0x72,
	0x65, 0x61, 0x64, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x77, 0x61, 0x72, 0x6e, 0x69,
	0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69,
	0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x57, 0x61,
	0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x48, 0x00, 0x52, 0x11, 0x72, 0x65, 0x61, 0x64, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x57, 0x61,
	0x72, 0x6e, 0x69, 0x6e, 0x67, 0x1a, 0x0e, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x6f, 0x74,
	0x46, 0x6f, 0x75, 0x6e, 0x64, 0x1a, 0x0e, 0x0a, 0x0c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x44,
	0x65, 0x6e, 0x69, 0x65, 0x64, 0x1a, 0x0c, 0x0a, 0x0a, 0x52, 0x65, 0x61, 0x64, 0x46, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_dataprime_v1_object_store_warning_proto_rawDescOnce sync.Once
	file_com_coralogix_dataprime_v1_object_store_warning_proto_rawDescData = file_com_coralogix_dataprime_v1_object_store_warning_proto_rawDesc
)

func file_com_coralogix_dataprime_v1_object_store_warning_proto_rawDescGZIP() []byte {
	file_com_coralogix_dataprime_v1_object_store_warning_proto_rawDescOnce.Do(func() {
		file_com_coralogix_dataprime_v1_object_store_warning_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_dataprime_v1_object_store_warning_proto_rawDescData)
	})
	return file_com_coralogix_dataprime_v1_object_store_warning_proto_rawDescData
}

var file_com_coralogix_dataprime_v1_object_store_warning_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_com_coralogix_dataprime_v1_object_store_warning_proto_goTypes = []any{
	(*ObjectStoreWarning)(nil),                               // 0: com.coralogix.dataprime.v1.ObjectStoreWarning
	(*ObjectStoreWarning_MetastoreWarning)(nil),              // 1: com.coralogix.dataprime.v1.ObjectStoreWarning.MetastoreWarning
	(*ObjectStoreWarning_BlockWarning)(nil),                  // 2: com.coralogix.dataprime.v1.ObjectStoreWarning.BlockWarning
	(*ObjectStoreWarning_MetastoreWarning_NoFilesFound)(nil), // 3: com.coralogix.dataprime.v1.ObjectStoreWarning.MetastoreWarning.NoFilesFound
	(*ObjectStoreWarning_BlockWarning_FileNotFound)(nil),     // 4: com.coralogix.dataprime.v1.ObjectStoreWarning.BlockWarning.FileNotFound
	(*ObjectStoreWarning_BlockWarning_AccessDenied)(nil),     // 5: com.coralogix.dataprime.v1.ObjectStoreWarning.BlockWarning.AccessDenied
	(*ObjectStoreWarning_BlockWarning_ReadFailed)(nil),       // 6: com.coralogix.dataprime.v1.ObjectStoreWarning.BlockWarning.ReadFailed
}
var file_com_coralogix_dataprime_v1_object_store_warning_proto_depIdxs = []int32{
	1, // 0: com.coralogix.dataprime.v1.ObjectStoreWarning.metastore_warning:type_name -> com.coralogix.dataprime.v1.ObjectStoreWarning.MetastoreWarning
	2, // 1: com.coralogix.dataprime.v1.ObjectStoreWarning.block_warning:type_name -> com.coralogix.dataprime.v1.ObjectStoreWarning.BlockWarning
	3, // 2: com.coralogix.dataprime.v1.ObjectStoreWarning.MetastoreWarning.no_files_found_warning:type_name -> com.coralogix.dataprime.v1.ObjectStoreWarning.MetastoreWarning.NoFilesFound
	4, // 3: com.coralogix.dataprime.v1.ObjectStoreWarning.BlockWarning.file_not_found_warning:type_name -> com.coralogix.dataprime.v1.ObjectStoreWarning.BlockWarning.FileNotFound
	5, // 4: com.coralogix.dataprime.v1.ObjectStoreWarning.BlockWarning.access_denied_warning:type_name -> com.coralogix.dataprime.v1.ObjectStoreWarning.BlockWarning.AccessDenied
	6, // 5: com.coralogix.dataprime.v1.ObjectStoreWarning.BlockWarning.read_failed_warning:type_name -> com.coralogix.dataprime.v1.ObjectStoreWarning.BlockWarning.ReadFailed
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_com_coralogix_dataprime_v1_object_store_warning_proto_init() }
func file_com_coralogix_dataprime_v1_object_store_warning_proto_init() {
	if File_com_coralogix_dataprime_v1_object_store_warning_proto != nil {
		return
	}
	file_com_coralogix_dataprime_v1_object_store_warning_proto_msgTypes[0].OneofWrappers = []any{
		(*ObjectStoreWarning_MetastoreWarning_)(nil),
		(*ObjectStoreWarning_BlockWarning_)(nil),
	}
	file_com_coralogix_dataprime_v1_object_store_warning_proto_msgTypes[1].OneofWrappers = []any{
		(*ObjectStoreWarning_MetastoreWarning_NoFilesFoundWarning)(nil),
	}
	file_com_coralogix_dataprime_v1_object_store_warning_proto_msgTypes[2].OneofWrappers = []any{
		(*ObjectStoreWarning_BlockWarning_FileNotFoundWarning)(nil),
		(*ObjectStoreWarning_BlockWarning_AccessDeniedWarning)(nil),
		(*ObjectStoreWarning_BlockWarning_ReadFailedWarning)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_dataprime_v1_object_store_warning_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_dataprime_v1_object_store_warning_proto_goTypes,
		DependencyIndexes: file_com_coralogix_dataprime_v1_object_store_warning_proto_depIdxs,
		MessageInfos:      file_com_coralogix_dataprime_v1_object_store_warning_proto_msgTypes,
	}.Build()
	File_com_coralogix_dataprime_v1_object_store_warning_proto = out.File
	file_com_coralogix_dataprime_v1_object_store_warning_proto_rawDesc = nil
	file_com_coralogix_dataprime_v1_object_store_warning_proto_goTypes = nil
	file_com_coralogix_dataprime_v1_object_store_warning_proto_depIdxs = nil
}
