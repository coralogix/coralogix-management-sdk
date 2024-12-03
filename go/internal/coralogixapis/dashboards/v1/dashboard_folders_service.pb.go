// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: com/coralogixapis/dashboards/v1/services/dashboard_folders_service.proto

package v1

import (
	_ "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/common/v1"
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

type CreateDashboardFolderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Folder    *DashboardFolder        `protobuf:"bytes,2,opt,name=folder,proto3" json:"folder,omitempty"`
}

func (x *CreateDashboardFolderRequest) Reset() {
	*x = CreateDashboardFolderRequest{}
	mi := &file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateDashboardFolderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDashboardFolderRequest) ProtoMessage() {}

func (x *CreateDashboardFolderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDashboardFolderRequest.ProtoReflect.Descriptor instead.
func (*CreateDashboardFolderRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDashboardFolderRequest) GetRequestId() *wrapperspb.StringValue {
	if x != nil {
		return x.RequestId
	}
	return nil
}

func (x *CreateDashboardFolderRequest) GetFolder() *DashboardFolder {
	if x != nil {
		return x.Folder
	}
	return nil
}

type CreateDashboardFolderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateDashboardFolderResponse) Reset() {
	*x = CreateDashboardFolderResponse{}
	mi := &file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateDashboardFolderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDashboardFolderResponse) ProtoMessage() {}

func (x *CreateDashboardFolderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDashboardFolderResponse.ProtoReflect.Descriptor instead.
func (*CreateDashboardFolderResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_rawDescGZIP(), []int{1}
}

type ReplaceDashboardFolderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Folder    *DashboardFolder        `protobuf:"bytes,2,opt,name=folder,proto3" json:"folder,omitempty"`
}

func (x *ReplaceDashboardFolderRequest) Reset() {
	*x = ReplaceDashboardFolderRequest{}
	mi := &file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReplaceDashboardFolderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplaceDashboardFolderRequest) ProtoMessage() {}

func (x *ReplaceDashboardFolderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplaceDashboardFolderRequest.ProtoReflect.Descriptor instead.
func (*ReplaceDashboardFolderRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_rawDescGZIP(), []int{2}
}

func (x *ReplaceDashboardFolderRequest) GetRequestId() *wrapperspb.StringValue {
	if x != nil {
		return x.RequestId
	}
	return nil
}

func (x *ReplaceDashboardFolderRequest) GetFolder() *DashboardFolder {
	if x != nil {
		return x.Folder
	}
	return nil
}

type ReplaceDashboardFolderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReplaceDashboardFolderResponse) Reset() {
	*x = ReplaceDashboardFolderResponse{}
	mi := &file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReplaceDashboardFolderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplaceDashboardFolderResponse) ProtoMessage() {}

func (x *ReplaceDashboardFolderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplaceDashboardFolderResponse.ProtoReflect.Descriptor instead.
func (*ReplaceDashboardFolderResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_rawDescGZIP(), []int{3}
}

type DeleteDashboardFolderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	FolderId  *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
}

func (x *DeleteDashboardFolderRequest) Reset() {
	*x = DeleteDashboardFolderRequest{}
	mi := &file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteDashboardFolderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDashboardFolderRequest) ProtoMessage() {}

func (x *DeleteDashboardFolderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDashboardFolderRequest.ProtoReflect.Descriptor instead.
func (*DeleteDashboardFolderRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteDashboardFolderRequest) GetRequestId() *wrapperspb.StringValue {
	if x != nil {
		return x.RequestId
	}
	return nil
}

func (x *DeleteDashboardFolderRequest) GetFolderId() *wrapperspb.StringValue {
	if x != nil {
		return x.FolderId
	}
	return nil
}

type DeleteDashboardFolderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteDashboardFolderResponse) Reset() {
	*x = DeleteDashboardFolderResponse{}
	mi := &file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteDashboardFolderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDashboardFolderResponse) ProtoMessage() {}

func (x *DeleteDashboardFolderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDashboardFolderResponse.ProtoReflect.Descriptor instead.
func (*DeleteDashboardFolderResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_rawDescGZIP(), []int{5}
}

type ListDashboardFoldersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListDashboardFoldersRequest) Reset() {
	*x = ListDashboardFoldersRequest{}
	mi := &file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListDashboardFoldersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDashboardFoldersRequest) ProtoMessage() {}

func (x *ListDashboardFoldersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDashboardFoldersRequest.ProtoReflect.Descriptor instead.
func (*ListDashboardFoldersRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_rawDescGZIP(), []int{6}
}

type ListDashboardFoldersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Folder []*DashboardFolder `protobuf:"bytes,1,rep,name=folder,proto3" json:"folder,omitempty"`
}

func (x *ListDashboardFoldersResponse) Reset() {
	*x = ListDashboardFoldersResponse{}
	mi := &file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListDashboardFoldersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDashboardFoldersResponse) ProtoMessage() {}

func (x *ListDashboardFoldersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDashboardFoldersResponse.ProtoReflect.Descriptor instead.
func (*ListDashboardFoldersResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_rawDescGZIP(), []int{7}
}

func (x *ListDashboardFoldersResponse) GetFolder() []*DashboardFolder {
	if x != nil {
		return x.Folder
	}
	return nil
}

type GetDashboardFolderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	FolderId  *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
}

func (x *GetDashboardFolderRequest) Reset() {
	*x = GetDashboardFolderRequest{}
	mi := &file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDashboardFolderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDashboardFolderRequest) ProtoMessage() {}

func (x *GetDashboardFolderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDashboardFolderRequest.ProtoReflect.Descriptor instead.
func (*GetDashboardFolderRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_rawDescGZIP(), []int{8}
}

func (x *GetDashboardFolderRequest) GetRequestId() *wrapperspb.StringValue {
	if x != nil {
		return x.RequestId
	}
	return nil
}

func (x *GetDashboardFolderRequest) GetFolderId() *wrapperspb.StringValue {
	if x != nil {
		return x.FolderId
	}
	return nil
}

type GetDashboardFolderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Folder *DashboardFolder `protobuf:"bytes,1,opt,name=folder,proto3" json:"folder,omitempty"`
}

func (x *GetDashboardFolderResponse) Reset() {
	*x = GetDashboardFolderResponse{}
	mi := &file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDashboardFolderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDashboardFolderResponse) ProtoMessage() {}

func (x *GetDashboardFolderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDashboardFolderResponse.ProtoReflect.Descriptor instead.
func (*GetDashboardFolderResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_rawDescGZIP(), []int{9}
}

func (x *GetDashboardFolderResponse) GetFolder() *DashboardFolder {
	if x != nil {
		return x.Folder
	}
	return nil
}

var File_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto protoreflect.FileDescriptor

var file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_rawDesc = []byte{
	0x0a, 0x48, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x5f, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x1a, 0x27, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x01, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x12, 0x4f, 0x0a, 0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x73, 0x68,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x06, 0x66, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x22, 0x1f, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0xad, 0x01, 0x0a, 0x1d, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x4f, 0x0a, 0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x06, 0x66, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x22, 0x20, 0x0a, 0x1e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x44,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x96, 0x01, 0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x1f, 0x0a, 0x1d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1d, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x6f, 0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4f, 0x0a, 0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x22, 0x93, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b,
	0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x09, 0x66,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x66, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x6d, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x44, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x06, 0x66,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x32, 0xff, 0x07, 0x0a, 0x17, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xc3, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x12, 0x45, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x46, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0xc2, 0xb8, 0x02, 0x18, 0x0a,
	0x16, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x20,
	0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x12, 0xbb, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x44,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x43,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x44, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0xc2, 0xb8, 0x02, 0x16, 0x0a,
	0x14, 0x67, 0x65, 0x74, 0x20, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x20, 0x66,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0xc7, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12,
	0x46, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x47, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1d, 0xc2, 0xb8, 0x02, 0x19, 0x0a, 0x17, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x20, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12,
	0xcb, 0x01, 0x0a, 0x16, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x44, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x47, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x44, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x48, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52,
	0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x46,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0xc2,
	0xb8, 0x02, 0x1a, 0x0a, 0x18, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x20, 0x64, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x20, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0xc7, 0x01,
	0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x46, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x47, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0xc2, 0xb8, 0x02, 0x19, 0x0a, 0x17,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x20, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_rawDescOnce sync.Once
	file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_rawDescData = file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_rawDesc
)

func file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_rawDescData)
	})
	return file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_rawDescData
}

var file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_goTypes = []any{
	(*CreateDashboardFolderRequest)(nil),   // 0: com.coralogixapis.dashboards.v1.services.CreateDashboardFolderRequest
	(*CreateDashboardFolderResponse)(nil),  // 1: com.coralogixapis.dashboards.v1.services.CreateDashboardFolderResponse
	(*ReplaceDashboardFolderRequest)(nil),  // 2: com.coralogixapis.dashboards.v1.services.ReplaceDashboardFolderRequest
	(*ReplaceDashboardFolderResponse)(nil), // 3: com.coralogixapis.dashboards.v1.services.ReplaceDashboardFolderResponse
	(*DeleteDashboardFolderRequest)(nil),   // 4: com.coralogixapis.dashboards.v1.services.DeleteDashboardFolderRequest
	(*DeleteDashboardFolderResponse)(nil),  // 5: com.coralogixapis.dashboards.v1.services.DeleteDashboardFolderResponse
	(*ListDashboardFoldersRequest)(nil),    // 6: com.coralogixapis.dashboards.v1.services.ListDashboardFoldersRequest
	(*ListDashboardFoldersResponse)(nil),   // 7: com.coralogixapis.dashboards.v1.services.ListDashboardFoldersResponse
	(*GetDashboardFolderRequest)(nil),      // 8: com.coralogixapis.dashboards.v1.services.GetDashboardFolderRequest
	(*GetDashboardFolderResponse)(nil),     // 9: com.coralogixapis.dashboards.v1.services.GetDashboardFolderResponse
	(*wrapperspb.StringValue)(nil),         // 10: google.protobuf.StringValue
	(*DashboardFolder)(nil),                // 11: com.coralogixapis.dashboards.v1.common.DashboardFolder
}
var file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_depIdxs = []int32{
	10, // 0: com.coralogixapis.dashboards.v1.services.CreateDashboardFolderRequest.request_id:type_name -> google.protobuf.StringValue
	11, // 1: com.coralogixapis.dashboards.v1.services.CreateDashboardFolderRequest.folder:type_name -> com.coralogixapis.dashboards.v1.common.DashboardFolder
	10, // 2: com.coralogixapis.dashboards.v1.services.ReplaceDashboardFolderRequest.request_id:type_name -> google.protobuf.StringValue
	11, // 3: com.coralogixapis.dashboards.v1.services.ReplaceDashboardFolderRequest.folder:type_name -> com.coralogixapis.dashboards.v1.common.DashboardFolder
	10, // 4: com.coralogixapis.dashboards.v1.services.DeleteDashboardFolderRequest.request_id:type_name -> google.protobuf.StringValue
	10, // 5: com.coralogixapis.dashboards.v1.services.DeleteDashboardFolderRequest.folder_id:type_name -> google.protobuf.StringValue
	11, // 6: com.coralogixapis.dashboards.v1.services.ListDashboardFoldersResponse.folder:type_name -> com.coralogixapis.dashboards.v1.common.DashboardFolder
	10, // 7: com.coralogixapis.dashboards.v1.services.GetDashboardFolderRequest.request_id:type_name -> google.protobuf.StringValue
	10, // 8: com.coralogixapis.dashboards.v1.services.GetDashboardFolderRequest.folder_id:type_name -> google.protobuf.StringValue
	11, // 9: com.coralogixapis.dashboards.v1.services.GetDashboardFolderResponse.folder:type_name -> com.coralogixapis.dashboards.v1.common.DashboardFolder
	6,  // 10: com.coralogixapis.dashboards.v1.services.DashboardFoldersService.ListDashboardFolders:input_type -> com.coralogixapis.dashboards.v1.services.ListDashboardFoldersRequest
	8,  // 11: com.coralogixapis.dashboards.v1.services.DashboardFoldersService.GetDashboardFolder:input_type -> com.coralogixapis.dashboards.v1.services.GetDashboardFolderRequest
	0,  // 12: com.coralogixapis.dashboards.v1.services.DashboardFoldersService.CreateDashboardFolder:input_type -> com.coralogixapis.dashboards.v1.services.CreateDashboardFolderRequest
	2,  // 13: com.coralogixapis.dashboards.v1.services.DashboardFoldersService.ReplaceDashboardFolder:input_type -> com.coralogixapis.dashboards.v1.services.ReplaceDashboardFolderRequest
	4,  // 14: com.coralogixapis.dashboards.v1.services.DashboardFoldersService.DeleteDashboardFolder:input_type -> com.coralogixapis.dashboards.v1.services.DeleteDashboardFolderRequest
	7,  // 15: com.coralogixapis.dashboards.v1.services.DashboardFoldersService.ListDashboardFolders:output_type -> com.coralogixapis.dashboards.v1.services.ListDashboardFoldersResponse
	9,  // 16: com.coralogixapis.dashboards.v1.services.DashboardFoldersService.GetDashboardFolder:output_type -> com.coralogixapis.dashboards.v1.services.GetDashboardFolderResponse
	1,  // 17: com.coralogixapis.dashboards.v1.services.DashboardFoldersService.CreateDashboardFolder:output_type -> com.coralogixapis.dashboards.v1.services.CreateDashboardFolderResponse
	3,  // 18: com.coralogixapis.dashboards.v1.services.DashboardFoldersService.ReplaceDashboardFolder:output_type -> com.coralogixapis.dashboards.v1.services.ReplaceDashboardFolderResponse
	5,  // 19: com.coralogixapis.dashboards.v1.services.DashboardFoldersService.DeleteDashboardFolder:output_type -> com.coralogixapis.dashboards.v1.services.DeleteDashboardFolderResponse
	15, // [15:20] is the sub-list for method output_type
	10, // [10:15] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_init() }
func file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_init() {
	if File_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto != nil {
		return
	}
	file_com_coralogixapis_dashboards_v1_common_folder_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto = out.File
	file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_rawDesc = nil
	file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_goTypes = nil
	file_com_coralogixapis_dashboards_v1_services_dashboard_folders_service_proto_depIdxs = nil
}
