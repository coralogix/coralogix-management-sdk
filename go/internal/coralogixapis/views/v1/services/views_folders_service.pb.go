// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: com/coralogixapis/views/v1/services/views_folders_service.proto

package services

import (
	v1 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/views/v1"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

// Create view folder
type CreateViewFolderRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// View folder name
	Name          *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateViewFolderRequest) Reset() {
	*x = CreateViewFolderRequest{}
	mi := &file_com_coralogixapis_views_v1_services_views_folders_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateViewFolderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateViewFolderRequest) ProtoMessage() {}

func (x *CreateViewFolderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_views_v1_services_views_folders_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateViewFolderRequest.ProtoReflect.Descriptor instead.
func (*CreateViewFolderRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_views_v1_services_views_folders_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateViewFolderRequest) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

type CreateViewFolderResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// View folder
	Folder        *v1.ViewFolder `protobuf:"bytes,1,opt,name=folder,proto3" json:"folder,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateViewFolderResponse) Reset() {
	*x = CreateViewFolderResponse{}
	mi := &file_com_coralogixapis_views_v1_services_views_folders_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateViewFolderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateViewFolderResponse) ProtoMessage() {}

func (x *CreateViewFolderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_views_v1_services_views_folders_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateViewFolderResponse.ProtoReflect.Descriptor instead.
func (*CreateViewFolderResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_views_v1_services_views_folders_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateViewFolderResponse) GetFolder() *v1.ViewFolder {
	if x != nil {
		return x.Folder
	}
	return nil
}

type GetViewFolderRequest struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Id            *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetViewFolderRequest) Reset() {
	*x = GetViewFolderRequest{}
	mi := &file_com_coralogixapis_views_v1_services_views_folders_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetViewFolderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetViewFolderRequest) ProtoMessage() {}

func (x *GetViewFolderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_views_v1_services_views_folders_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetViewFolderRequest.ProtoReflect.Descriptor instead.
func (*GetViewFolderRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_views_v1_services_views_folders_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetViewFolderRequest) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

type GetViewFolderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Folder        *v1.ViewFolder         `protobuf:"bytes,1,opt,name=folder,proto3" json:"folder,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetViewFolderResponse) Reset() {
	*x = GetViewFolderResponse{}
	mi := &file_com_coralogixapis_views_v1_services_views_folders_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetViewFolderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetViewFolderResponse) ProtoMessage() {}

func (x *GetViewFolderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_views_v1_services_views_folders_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetViewFolderResponse.ProtoReflect.Descriptor instead.
func (*GetViewFolderResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_views_v1_services_views_folders_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetViewFolderResponse) GetFolder() *v1.ViewFolder {
	if x != nil {
		return x.Folder
	}
	return nil
}

type ListViewFoldersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListViewFoldersRequest) Reset() {
	*x = ListViewFoldersRequest{}
	mi := &file_com_coralogixapis_views_v1_services_views_folders_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListViewFoldersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListViewFoldersRequest) ProtoMessage() {}

func (x *ListViewFoldersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_views_v1_services_views_folders_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListViewFoldersRequest.ProtoReflect.Descriptor instead.
func (*ListViewFoldersRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_views_v1_services_views_folders_service_proto_rawDescGZIP(), []int{4}
}

type ListViewFoldersResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of view folders
	Folders       []*v1.ViewFolder `protobuf:"bytes,1,rep,name=folders,proto3" json:"folders,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListViewFoldersResponse) Reset() {
	*x = ListViewFoldersResponse{}
	mi := &file_com_coralogixapis_views_v1_services_views_folders_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListViewFoldersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListViewFoldersResponse) ProtoMessage() {}

func (x *ListViewFoldersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_views_v1_services_views_folders_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListViewFoldersResponse.ProtoReflect.Descriptor instead.
func (*ListViewFoldersResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_views_v1_services_views_folders_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListViewFoldersResponse) GetFolders() []*v1.ViewFolder {
	if x != nil {
		return x.Folders
	}
	return nil
}

type DeleteViewFolderRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// View folder id
	Id            *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteViewFolderRequest) Reset() {
	*x = DeleteViewFolderRequest{}
	mi := &file_com_coralogixapis_views_v1_services_views_folders_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteViewFolderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteViewFolderRequest) ProtoMessage() {}

func (x *DeleteViewFolderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_views_v1_services_views_folders_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteViewFolderRequest.ProtoReflect.Descriptor instead.
func (*DeleteViewFolderRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_views_v1_services_views_folders_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteViewFolderRequest) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

type DeleteViewFolderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteViewFolderResponse) Reset() {
	*x = DeleteViewFolderResponse{}
	mi := &file_com_coralogixapis_views_v1_services_views_folders_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteViewFolderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteViewFolderResponse) ProtoMessage() {}

func (x *DeleteViewFolderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_views_v1_services_views_folders_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteViewFolderResponse.ProtoReflect.Descriptor instead.
func (*DeleteViewFolderResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_views_v1_services_views_folders_service_proto_rawDescGZIP(), []int{7}
}

type ReplaceViewFolderRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// View folder
	Folder        *v1.ViewFolder `protobuf:"bytes,1,opt,name=folder,proto3" json:"folder,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReplaceViewFolderRequest) Reset() {
	*x = ReplaceViewFolderRequest{}
	mi := &file_com_coralogixapis_views_v1_services_views_folders_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReplaceViewFolderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplaceViewFolderRequest) ProtoMessage() {}

func (x *ReplaceViewFolderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_views_v1_services_views_folders_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplaceViewFolderRequest.ProtoReflect.Descriptor instead.
func (*ReplaceViewFolderRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_views_v1_services_views_folders_service_proto_rawDescGZIP(), []int{8}
}

func (x *ReplaceViewFolderRequest) GetFolder() *v1.ViewFolder {
	if x != nil {
		return x.Folder
	}
	return nil
}

type ReplaceViewFolderResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// View folder
	Folder        *v1.ViewFolder `protobuf:"bytes,1,opt,name=folder,proto3" json:"folder,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReplaceViewFolderResponse) Reset() {
	*x = ReplaceViewFolderResponse{}
	mi := &file_com_coralogixapis_views_v1_services_views_folders_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReplaceViewFolderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplaceViewFolderResponse) ProtoMessage() {}

func (x *ReplaceViewFolderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_views_v1_services_views_folders_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplaceViewFolderResponse.ProtoReflect.Descriptor instead.
func (*ReplaceViewFolderResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_views_v1_services_views_folders_service_proto_rawDescGZIP(), []int{9}
}

func (x *ReplaceViewFolderResponse) GetFolder() *v1.ViewFolder {
	if x != nil {
		return x.Folder
	}
	return nil
}

var File_com_coralogixapis_views_v1_services_views_folders_service_proto protoreflect.FileDescriptor

var file_com_coralogixapis_views_v1_services_views_folders_service_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x69, 0x65, 0x77, 0x73, 0x5f, 0x66, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x2c, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56,
	0x69, 0x65, 0x77, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x52, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x20, 0x92, 0x41,
	0x1d, 0x32, 0x0b, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x4a, 0x0b,
	0x22, 0x4d, 0x79, 0x20, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x22, 0x80, 0x01, 0x01, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x33, 0x92, 0x41, 0x30, 0x0a, 0x2e, 0x2a, 0x17, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x56, 0x69, 0x65, 0x77, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x32, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x76, 0x69, 0x65,
	0x77, 0x20, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x22, 0x9d, 0x01, 0x0a, 0x18, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x56, 0x69, 0x65, 0x77, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x06,
	0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x3a, 0x41, 0x92, 0x41, 0x3e, 0x0a, 0x3c, 0x2a, 0x17, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x69, 0x65, 0x77, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x21, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20,
	0x66, 0x6f, 0x72, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x76, 0x69, 0x65,
	0x77, 0x20, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x22, 0x9f, 0x02, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x56, 0x69, 0x65, 0x77, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0xc7, 0x01, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x98, 0x01,
	0x92, 0x41, 0x94, 0x01, 0x32, 0x1d, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x20, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x66, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x73, 0x4a, 0x26, 0x22, 0x33, 0x64, 0x63, 0x30, 0x32, 0x39, 0x39, 0x38, 0x2d, 0x30,
	0x62, 0x35, 0x30, 0x2d, 0x34, 0x65, 0x61, 0x38, 0x2d, 0x62, 0x36, 0x38, 0x61, 0x2d, 0x34, 0x37,
	0x37, 0x39, 0x64, 0x37, 0x31, 0x36, 0x66, 0x61, 0x31, 0x66, 0x22, 0x78, 0x24, 0x80, 0x01, 0x24,
	0x8a, 0x01, 0x3e, 0x5e, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x66, 0x5d, 0x7b, 0x38, 0x7d, 0x2d,
	0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x66, 0x5d, 0x7b, 0x34, 0x7d, 0x2d, 0x5b, 0x30, 0x2d, 0x39,
	0x61, 0x2d, 0x66, 0x5d, 0x7b, 0x34, 0x7d, 0x2d, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x66, 0x5d,
	0x7b, 0x34, 0x7d, 0x2d, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x66, 0x5d, 0x7b, 0x31, 0x32, 0x7d,
	0x24, 0xa2, 0x02, 0x04, 0x75, 0x75, 0x69, 0x64, 0x52, 0x02, 0x69, 0x64, 0x3a, 0x3d, 0x92, 0x41,
	0x3a, 0x0a, 0x38, 0x2a, 0x14, 0x47, 0x65, 0x74, 0x56, 0x69, 0x65, 0x77, 0x46, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x20, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x76,
	0x69, 0x65, 0x77, 0x20, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x22, 0x57, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x56, 0x69, 0x65, 0x77, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x06, 0x66, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x22, 0x18, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x65, 0x77,
	0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5b,
	0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x65, 0x77, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x66, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76,
	0x69, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x46, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x52, 0x07, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x22, 0xe3, 0x01, 0x0a, 0x17,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x69, 0x65, 0x77, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0xc7, 0x01, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x98, 0x01, 0x92, 0x41, 0x94, 0x01, 0x32, 0x1d, 0x55, 0x6e, 0x69, 0x71, 0x75,
	0x65, 0x20, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x20, 0x66, 0x6f, 0x72,
	0x20, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x4a, 0x26, 0x22, 0x33, 0x64, 0x63, 0x30, 0x32,
	0x39, 0x39, 0x38, 0x2d, 0x30, 0x62, 0x35, 0x30, 0x2d, 0x34, 0x65, 0x61, 0x38, 0x2d, 0x62, 0x36,
	0x38, 0x61, 0x2d, 0x34, 0x37, 0x37, 0x39, 0x64, 0x37, 0x31, 0x36, 0x66, 0x61, 0x31, 0x66, 0x22,
	0x78, 0x24, 0x80, 0x01, 0x24, 0x8a, 0x01, 0x3e, 0x5e, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x66,
	0x5d, 0x7b, 0x38, 0x7d, 0x2d, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x66, 0x5d, 0x7b, 0x34, 0x7d,
	0x2d, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x66, 0x5d, 0x7b, 0x34, 0x7d, 0x2d, 0x5b, 0x30, 0x2d,
	0x39, 0x61, 0x2d, 0x66, 0x5d, 0x7b, 0x34, 0x7d, 0x2d, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x66,
	0x5d, 0x7b, 0x31, 0x32, 0x7d, 0x24, 0xa2, 0x02, 0x04, 0x75, 0x75, 0x69, 0x64, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x1a, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x69, 0x65, 0x77, 0x46,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5a, 0x0a,
	0x18, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x56, 0x69, 0x65, 0x77, 0x46, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x06, 0x66, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x69,
	0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x46, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x52, 0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x22, 0x5b, 0x0a, 0x19, 0x52, 0x65, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x56, 0x69, 0x65, 0x77, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x06,
	0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x32, 0xfa, 0x0b, 0x0a, 0x13, 0x56, 0x69, 0x65, 0x77, 0x73,
	0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xb2,
	0x02, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x65, 0x77, 0x46, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x73, 0x12, 0x3b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x65,
	0x77, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x3c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x65, 0x77, 0x46, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa3, 0x01,
	0x92, 0x41, 0x87, 0x01, 0x1a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x76, 0x69, 0x65, 0x77, 0x27,
	0x73, 0x20, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x4a, 0x70, 0x0a, 0x03, 0x32, 0x30, 0x30,
	0x12, 0x69, 0x22, 0x67, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x53, 0x7b, 0x22, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x66,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x22, 0x3a, 0x5b, 0x7b, 0x22, 0x69, 0x64, 0x22, 0x3a, 0x22,
	0x30, 0x63, 0x64, 0x34, 0x65, 0x64, 0x35, 0x34, 0x2d, 0x39, 0x62, 0x35, 0x39, 0x2d, 0x34, 0x65,
	0x65, 0x66, 0x2d, 0x62, 0x33, 0x32, 0x39, 0x2d, 0x31, 0x35, 0x31, 0x61, 0x31, 0x36, 0x37, 0x65,
	0x32, 0x33, 0x30, 0x34, 0x22, 0x2c, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x22, 0x4d, 0x79,
	0x20, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x22, 0x7d, 0x5d, 0x7d, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x12, 0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x66, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x73, 0x12, 0xa4, 0x02, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x56, 0x69, 0x65, 0x77, 0x46,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x56,
	0x69, 0x65, 0x77, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x3a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69, 0x65, 0x77, 0x46, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x9b, 0x01, 0x92,
	0x41, 0x73, 0x1a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x76, 0x69, 0x65, 0x77, 0x20,
	0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x4a, 0x5d, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x56, 0x22,
	0x54, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x12, 0x40, 0x7b, 0x22, 0x69, 0x64, 0x22, 0x3a, 0x22, 0x30, 0x63, 0x64, 0x34,
	0x65, 0x64, 0x35, 0x34, 0x2d, 0x39, 0x62, 0x35, 0x39, 0x2d, 0x34, 0x65, 0x65, 0x66, 0x2d, 0x62,
	0x33, 0x32, 0x39, 0x2d, 0x31, 0x35, 0x31, 0x61, 0x31, 0x36, 0x37, 0x65, 0x32, 0x33, 0x30, 0x34,
	0x22, 0x2c, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x22, 0x4d, 0x79, 0x20, 0x46, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x22, 0x7d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x62, 0x06, 0x66, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x12, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x66, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0xab, 0x02, 0x0a, 0x10, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x69, 0x65, 0x77, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12,
	0x3c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x69, 0x65, 0x77,
	0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x69, 0x65, 0x77, 0x46, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x99, 0x01, 0x92,
	0x41, 0x73, 0x1a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x76, 0x69, 0x65, 0x77, 0x20,
	0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x4a, 0x5d, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x56, 0x22,
	0x54, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x12, 0x40, 0x7b, 0x22, 0x69, 0x64, 0x22, 0x3a, 0x22, 0x30, 0x63, 0x64, 0x34,
	0x65, 0x64, 0x35, 0x34, 0x2d, 0x39, 0x62, 0x35, 0x39, 0x2d, 0x34, 0x65, 0x65, 0x66, 0x2d, 0x62,
	0x33, 0x32, 0x39, 0x2d, 0x31, 0x35, 0x31, 0x61, 0x31, 0x36, 0x37, 0x65, 0x32, 0x33, 0x30, 0x34,
	0x22, 0x2c, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x22, 0x4d, 0x79, 0x20, 0x46, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x22, 0x7d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x62, 0x06,
	0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x69, 0x65, 0x77,
	0x5f, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x12, 0xce, 0x01, 0x0a, 0x10, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x56, 0x69, 0x65, 0x77, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x3c, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x69, 0x65, 0x77, 0x46, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x69, 0x65, 0x77, 0x46, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3d, 0x92, 0x41, 0x1d, 0x1a,
	0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x73, 0x20, 0x61, 0x20, 0x76, 0x69, 0x65, 0x77, 0x20,
	0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x20, 0x62, 0x79, 0x20, 0x49, 0x44, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x17, 0x2a, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x66, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0xce, 0x02, 0x0a, 0x11, 0x52, 0x65,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x56, 0x69, 0x65, 0x77, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12,
	0x3d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x56, 0x69, 0x65,
	0x77, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3e,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x56, 0x69, 0x65, 0x77,
	0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xb9,
	0x01, 0x92, 0x41, 0x81, 0x01, 0x1a, 0x20, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x20,
	0x61, 0x6e, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x76, 0x69, 0x65, 0x77,
	0x20, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x4a, 0x5d, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x56,
	0x22, 0x54, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x40, 0x7b, 0x22, 0x69, 0x64, 0x22, 0x3a, 0x22, 0x30, 0x63, 0x64,
	0x34, 0x65, 0x64, 0x35, 0x34, 0x2d, 0x39, 0x62, 0x35, 0x39, 0x2d, 0x34, 0x65, 0x65, 0x66, 0x2d,
	0x62, 0x33, 0x32, 0x39, 0x2d, 0x31, 0x35, 0x31, 0x61, 0x31, 0x36, 0x37, 0x65, 0x32, 0x33, 0x30,
	0x34, 0x22, 0x2c, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x22, 0x4d, 0x79, 0x20, 0x46, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x22, 0x7d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x3a, 0x06, 0x66, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x62, 0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x1a, 0x1c, 0x2f, 0x76,
	0x31, 0x2f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x7b,
	0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x69, 0x64, 0x7d, 0x1a, 0x37, 0x92, 0x41, 0x34, 0x0a,
	0x11, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x76, 0x69, 0x65,
	0x77, 0x73, 0x12, 0x1f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x20, 0x76, 0x69, 0x65, 0x77, 0x20, 0x66, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x73, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_views_v1_services_views_folders_service_proto_rawDescOnce sync.Once
	file_com_coralogixapis_views_v1_services_views_folders_service_proto_rawDescData = file_com_coralogixapis_views_v1_services_views_folders_service_proto_rawDesc
)

func file_com_coralogixapis_views_v1_services_views_folders_service_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_views_v1_services_views_folders_service_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_views_v1_services_views_folders_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_views_v1_services_views_folders_service_proto_rawDescData)
	})
	return file_com_coralogixapis_views_v1_services_views_folders_service_proto_rawDescData
}

var file_com_coralogixapis_views_v1_services_views_folders_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_com_coralogixapis_views_v1_services_views_folders_service_proto_goTypes = []any{
	(*CreateViewFolderRequest)(nil),   // 0: com.coralogixapis.views.v1.services.CreateViewFolderRequest
	(*CreateViewFolderResponse)(nil),  // 1: com.coralogixapis.views.v1.services.CreateViewFolderResponse
	(*GetViewFolderRequest)(nil),      // 2: com.coralogixapis.views.v1.services.GetViewFolderRequest
	(*GetViewFolderResponse)(nil),     // 3: com.coralogixapis.views.v1.services.GetViewFolderResponse
	(*ListViewFoldersRequest)(nil),    // 4: com.coralogixapis.views.v1.services.ListViewFoldersRequest
	(*ListViewFoldersResponse)(nil),   // 5: com.coralogixapis.views.v1.services.ListViewFoldersResponse
	(*DeleteViewFolderRequest)(nil),   // 6: com.coralogixapis.views.v1.services.DeleteViewFolderRequest
	(*DeleteViewFolderResponse)(nil),  // 7: com.coralogixapis.views.v1.services.DeleteViewFolderResponse
	(*ReplaceViewFolderRequest)(nil),  // 8: com.coralogixapis.views.v1.services.ReplaceViewFolderRequest
	(*ReplaceViewFolderResponse)(nil), // 9: com.coralogixapis.views.v1.services.ReplaceViewFolderResponse
	(*wrapperspb.StringValue)(nil),    // 10: google.protobuf.StringValue
	(*v1.ViewFolder)(nil),             // 11: com.coralogixapis.views.v1.ViewFolder
}
var file_com_coralogixapis_views_v1_services_views_folders_service_proto_depIdxs = []int32{
	10, // 0: com.coralogixapis.views.v1.services.CreateViewFolderRequest.name:type_name -> google.protobuf.StringValue
	11, // 1: com.coralogixapis.views.v1.services.CreateViewFolderResponse.folder:type_name -> com.coralogixapis.views.v1.ViewFolder
	10, // 2: com.coralogixapis.views.v1.services.GetViewFolderRequest.id:type_name -> google.protobuf.StringValue
	11, // 3: com.coralogixapis.views.v1.services.GetViewFolderResponse.folder:type_name -> com.coralogixapis.views.v1.ViewFolder
	11, // 4: com.coralogixapis.views.v1.services.ListViewFoldersResponse.folders:type_name -> com.coralogixapis.views.v1.ViewFolder
	10, // 5: com.coralogixapis.views.v1.services.DeleteViewFolderRequest.id:type_name -> google.protobuf.StringValue
	11, // 6: com.coralogixapis.views.v1.services.ReplaceViewFolderRequest.folder:type_name -> com.coralogixapis.views.v1.ViewFolder
	11, // 7: com.coralogixapis.views.v1.services.ReplaceViewFolderResponse.folder:type_name -> com.coralogixapis.views.v1.ViewFolder
	4,  // 8: com.coralogixapis.views.v1.services.ViewsFoldersService.ListViewFolders:input_type -> com.coralogixapis.views.v1.services.ListViewFoldersRequest
	2,  // 9: com.coralogixapis.views.v1.services.ViewsFoldersService.GetViewFolder:input_type -> com.coralogixapis.views.v1.services.GetViewFolderRequest
	0,  // 10: com.coralogixapis.views.v1.services.ViewsFoldersService.CreateViewFolder:input_type -> com.coralogixapis.views.v1.services.CreateViewFolderRequest
	6,  // 11: com.coralogixapis.views.v1.services.ViewsFoldersService.DeleteViewFolder:input_type -> com.coralogixapis.views.v1.services.DeleteViewFolderRequest
	8,  // 12: com.coralogixapis.views.v1.services.ViewsFoldersService.ReplaceViewFolder:input_type -> com.coralogixapis.views.v1.services.ReplaceViewFolderRequest
	5,  // 13: com.coralogixapis.views.v1.services.ViewsFoldersService.ListViewFolders:output_type -> com.coralogixapis.views.v1.services.ListViewFoldersResponse
	3,  // 14: com.coralogixapis.views.v1.services.ViewsFoldersService.GetViewFolder:output_type -> com.coralogixapis.views.v1.services.GetViewFolderResponse
	1,  // 15: com.coralogixapis.views.v1.services.ViewsFoldersService.CreateViewFolder:output_type -> com.coralogixapis.views.v1.services.CreateViewFolderResponse
	7,  // 16: com.coralogixapis.views.v1.services.ViewsFoldersService.DeleteViewFolder:output_type -> com.coralogixapis.views.v1.services.DeleteViewFolderResponse
	9,  // 17: com.coralogixapis.views.v1.services.ViewsFoldersService.ReplaceViewFolder:output_type -> com.coralogixapis.views.v1.services.ReplaceViewFolderResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_views_v1_services_views_folders_service_proto_init() }
func file_com_coralogixapis_views_v1_services_views_folders_service_proto_init() {
	if File_com_coralogixapis_views_v1_services_views_folders_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_views_v1_services_views_folders_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogixapis_views_v1_services_views_folders_service_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_views_v1_services_views_folders_service_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_views_v1_services_views_folders_service_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_views_v1_services_views_folders_service_proto = out.File
	file_com_coralogixapis_views_v1_services_views_folders_service_proto_rawDesc = nil
	file_com_coralogixapis_views_v1_services_views_folders_service_proto_goTypes = nil
	file_com_coralogixapis_views_v1_services_views_folders_service_proto_depIdxs = nil
}
