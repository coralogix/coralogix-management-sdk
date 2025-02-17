// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.26.1
// source: com/coralogixapis/views/v1/services/views_service.proto

package services

import (
	v1 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/views/v1"
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

type CreateViewRequest struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Name          *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SearchQuery   *v1.SearchQuery         `protobuf:"bytes,2,opt,name=search_query,json=searchQuery,proto3" json:"search_query,omitempty"`
	TimeSelection *v1.TimeSelection       `protobuf:"bytes,3,opt,name=time_selection,json=timeSelection,proto3" json:"time_selection,omitempty"`
	Filters       *v1.SelectedFilters     `protobuf:"bytes,4,opt,name=filters,proto3" json:"filters,omitempty"`
	FolderId      *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateViewRequest) Reset() {
	*x = CreateViewRequest{}
	mi := &file_com_coralogixapis_views_v1_services_views_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateViewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateViewRequest) ProtoMessage() {}

func (x *CreateViewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_views_v1_services_views_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateViewRequest.ProtoReflect.Descriptor instead.
func (*CreateViewRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_views_v1_services_views_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateViewRequest) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *CreateViewRequest) GetSearchQuery() *v1.SearchQuery {
	if x != nil {
		return x.SearchQuery
	}
	return nil
}

func (x *CreateViewRequest) GetTimeSelection() *v1.TimeSelection {
	if x != nil {
		return x.TimeSelection
	}
	return nil
}

func (x *CreateViewRequest) GetFilters() *v1.SelectedFilters {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *CreateViewRequest) GetFolderId() *wrapperspb.StringValue {
	if x != nil {
		return x.FolderId
	}
	return nil
}

type CreateViewResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	View          *View                  `protobuf:"bytes,1,opt,name=view,proto3" json:"view,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateViewResponse) Reset() {
	*x = CreateViewResponse{}
	mi := &file_com_coralogixapis_views_v1_services_views_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateViewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateViewResponse) ProtoMessage() {}

func (x *CreateViewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_views_v1_services_views_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateViewResponse.ProtoReflect.Descriptor instead.
func (*CreateViewResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_views_v1_services_views_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateViewResponse) GetView() *View {
	if x != nil {
		return x.View
	}
	return nil
}

type ReplaceViewRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	View          *View                  `protobuf:"bytes,1,opt,name=view,proto3" json:"view,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReplaceViewRequest) Reset() {
	*x = ReplaceViewRequest{}
	mi := &file_com_coralogixapis_views_v1_services_views_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReplaceViewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplaceViewRequest) ProtoMessage() {}

func (x *ReplaceViewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_views_v1_services_views_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplaceViewRequest.ProtoReflect.Descriptor instead.
func (*ReplaceViewRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_views_v1_services_views_service_proto_rawDescGZIP(), []int{2}
}

func (x *ReplaceViewRequest) GetView() *View {
	if x != nil {
		return x.View
	}
	return nil
}

type ReplaceViewResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	View          *View                  `protobuf:"bytes,1,opt,name=view,proto3" json:"view,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReplaceViewResponse) Reset() {
	*x = ReplaceViewResponse{}
	mi := &file_com_coralogixapis_views_v1_services_views_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReplaceViewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplaceViewResponse) ProtoMessage() {}

func (x *ReplaceViewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_views_v1_services_views_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplaceViewResponse.ProtoReflect.Descriptor instead.
func (*ReplaceViewResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_views_v1_services_views_service_proto_rawDescGZIP(), []int{3}
}

func (x *ReplaceViewResponse) GetView() *View {
	if x != nil {
		return x.View
	}
	return nil
}

type GetViewRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            *wrapperspb.Int32Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetViewRequest) Reset() {
	*x = GetViewRequest{}
	mi := &file_com_coralogixapis_views_v1_services_views_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetViewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetViewRequest) ProtoMessage() {}

func (x *GetViewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_views_v1_services_views_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetViewRequest.ProtoReflect.Descriptor instead.
func (*GetViewRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_views_v1_services_views_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetViewRequest) GetId() *wrapperspb.Int32Value {
	if x != nil {
		return x.Id
	}
	return nil
}

type GetViewResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	View          *View                  `protobuf:"bytes,1,opt,name=view,proto3" json:"view,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetViewResponse) Reset() {
	*x = GetViewResponse{}
	mi := &file_com_coralogixapis_views_v1_services_views_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetViewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetViewResponse) ProtoMessage() {}

func (x *GetViewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_views_v1_services_views_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetViewResponse.ProtoReflect.Descriptor instead.
func (*GetViewResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_views_v1_services_views_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetViewResponse) GetView() *View {
	if x != nil {
		return x.View
	}
	return nil
}

type DeleteViewRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            *wrapperspb.Int32Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteViewRequest) Reset() {
	*x = DeleteViewRequest{}
	mi := &file_com_coralogixapis_views_v1_services_views_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteViewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteViewRequest) ProtoMessage() {}

func (x *DeleteViewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_views_v1_services_views_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteViewRequest.ProtoReflect.Descriptor instead.
func (*DeleteViewRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_views_v1_services_views_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteViewRequest) GetId() *wrapperspb.Int32Value {
	if x != nil {
		return x.Id
	}
	return nil
}

type DeleteViewResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteViewResponse) Reset() {
	*x = DeleteViewResponse{}
	mi := &file_com_coralogixapis_views_v1_services_views_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteViewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteViewResponse) ProtoMessage() {}

func (x *DeleteViewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_views_v1_services_views_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteViewResponse.ProtoReflect.Descriptor instead.
func (*DeleteViewResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_views_v1_services_views_service_proto_rawDescGZIP(), []int{7}
}

type ListViewsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListViewsRequest) Reset() {
	*x = ListViewsRequest{}
	mi := &file_com_coralogixapis_views_v1_services_views_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListViewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListViewsRequest) ProtoMessage() {}

func (x *ListViewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_views_v1_services_views_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListViewsRequest.ProtoReflect.Descriptor instead.
func (*ListViewsRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_views_v1_services_views_service_proto_rawDescGZIP(), []int{8}
}

type ListViewsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Views         []*View                `protobuf:"bytes,1,rep,name=views,proto3" json:"views,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListViewsResponse) Reset() {
	*x = ListViewsResponse{}
	mi := &file_com_coralogixapis_views_v1_services_views_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListViewsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListViewsResponse) ProtoMessage() {}

func (x *ListViewsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_views_v1_services_views_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListViewsResponse.ProtoReflect.Descriptor instead.
func (*ListViewsResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_views_v1_services_views_service_proto_rawDescGZIP(), []int{9}
}

func (x *ListViewsResponse) GetViews() []*View {
	if x != nil {
		return x.Views
	}
	return nil
}

type View struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Id            *wrapperspb.Int32Value  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	SearchQuery   *v1.SearchQuery         `protobuf:"bytes,3,opt,name=search_query,json=searchQuery,proto3" json:"search_query,omitempty"`
	TimeSelection *v1.TimeSelection       `protobuf:"bytes,4,opt,name=time_selection,json=timeSelection,proto3" json:"time_selection,omitempty"`
	Filters       *v1.SelectedFilters     `protobuf:"bytes,5,opt,name=filters,proto3" json:"filters,omitempty"`
	FolderId      *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	IsCompactMode *wrapperspb.BoolValue   `protobuf:"bytes,7,opt,name=is_compact_mode,json=isCompactMode,proto3" json:"is_compact_mode,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *View) Reset() {
	*x = View{}
	mi := &file_com_coralogixapis_views_v1_services_views_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *View) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*View) ProtoMessage() {}

func (x *View) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_views_v1_services_views_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use View.ProtoReflect.Descriptor instead.
func (*View) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_views_v1_services_views_service_proto_rawDescGZIP(), []int{10}
}

func (x *View) GetId() *wrapperspb.Int32Value {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *View) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *View) GetSearchQuery() *v1.SearchQuery {
	if x != nil {
		return x.SearchQuery
	}
	return nil
}

func (x *View) GetTimeSelection() *v1.TimeSelection {
	if x != nil {
		return x.TimeSelection
	}
	return nil
}

func (x *View) GetFilters() *v1.SelectedFilters {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *View) GetFolderId() *wrapperspb.StringValue {
	if x != nil {
		return x.FolderId
	}
	return nil
}

func (x *View) GetIsCompactMode() *wrapperspb.BoolValue {
	if x != nil {
		return x.IsCompactMode
	}
	return nil
}

var File_com_coralogixapis_views_v1_services_views_service_proto protoreflect.FileDescriptor

var file_com_coralogixapis_views_v1_services_views_service_proto_rawDesc = []byte{
	0x0a, 0x37, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x69, 0x65, 0x77, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x69, 0x65,
	0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x2a,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x69,
	0x65, 0x77, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x02, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56,
	0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4a, 0x0a, 0x0c,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x0b, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x50, 0x0a, 0x0e, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x74, 0x69, 0x6d,
	0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x07, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x12, 0x39, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x53, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x76, 0x69, 0x65, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x52, 0x04, 0x76, 0x69, 0x65,
	0x77, 0x22, 0x53, 0x0a, 0x12, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x56, 0x69, 0x65, 0x77,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x04, 0x76, 0x69, 0x65, 0x77, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x65, 0x77,
	0x52, 0x04, 0x76, 0x69, 0x65, 0x77, 0x22, 0x54, 0x0a, 0x13, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a,
	0x04, 0x76, 0x69, 0x65, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x52, 0x04, 0x76, 0x69, 0x65, 0x77, 0x22, 0x3d, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x69, 0x64, 0x22, 0x50, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d,
	0x0a, 0x04, 0x76, 0x69, 0x65, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x52, 0x04, 0x76, 0x69, 0x65, 0x77, 0x22, 0x40, 0x0a,
	0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2b, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x65,
	0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x54, 0x0a, 0x11, 0x4c, 0x69, 0x73,
	0x74, 0x56, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f,
	0x0a, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x52, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x22,
	0xc9, 0x03, 0x0a, 0x04, 0x56, 0x69, 0x65, 0x77, 0x12, 0x2b, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4a, 0x0a, 0x0c, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x50, 0x0a, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x39, 0x0a, 0x09,
	0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x66,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x42, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x63, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x69, 0x73,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x32, 0xf1, 0x05, 0x0a, 0x0c,
	0x56, 0x69, 0x65, 0x77, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x90, 0x01, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x69, 0x65, 0x77, 0x12, 0x36, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0xba, 0xb8,
	0x02, 0x0d, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x76, 0x69, 0x65, 0x77, 0x12,
	0x94, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x56, 0x69, 0x65, 0x77, 0x12,
	0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x56, 0x69, 0x65,
	0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x69, 0x65,
	0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52,
	0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x12, 0xba, 0xb8, 0x02, 0x0e, 0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x20, 0x76, 0x69, 0x65, 0x77, 0x12, 0x84, 0x01, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x56, 0x69,
	0x65, 0x77, 0x12, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69, 0x65, 0x77,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x69, 0x65, 0x77,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0e, 0xba,
	0xb8, 0x02, 0x0a, 0x0a, 0x08, 0x67, 0x65, 0x74, 0x20, 0x76, 0x69, 0x65, 0x77, 0x12, 0x8d, 0x01,
	0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x69, 0x65, 0x77, 0x12, 0x36, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0e, 0xba,
	0xb8, 0x02, 0x0a, 0x0a, 0x08, 0x67, 0x65, 0x74, 0x20, 0x76, 0x69, 0x65, 0x77, 0x12, 0x9f, 0x01,
	0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x65, 0x77, 0x73, 0x12, 0x35, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x65,
	0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0xba, 0xb8, 0x02, 0x1f,
	0x0a, 0x1d, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x20, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x20, 0x76, 0x69, 0x65, 0x77, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_views_v1_services_views_service_proto_rawDescOnce sync.Once
	file_com_coralogixapis_views_v1_services_views_service_proto_rawDescData = file_com_coralogixapis_views_v1_services_views_service_proto_rawDesc
)

func file_com_coralogixapis_views_v1_services_views_service_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_views_v1_services_views_service_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_views_v1_services_views_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_views_v1_services_views_service_proto_rawDescData)
	})
	return file_com_coralogixapis_views_v1_services_views_service_proto_rawDescData
}

var file_com_coralogixapis_views_v1_services_views_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_com_coralogixapis_views_v1_services_views_service_proto_goTypes = []any{
	(*CreateViewRequest)(nil),      // 0: com.coralogixapis.views.v1.services.CreateViewRequest
	(*CreateViewResponse)(nil),     // 1: com.coralogixapis.views.v1.services.CreateViewResponse
	(*ReplaceViewRequest)(nil),     // 2: com.coralogixapis.views.v1.services.ReplaceViewRequest
	(*ReplaceViewResponse)(nil),    // 3: com.coralogixapis.views.v1.services.ReplaceViewResponse
	(*GetViewRequest)(nil),         // 4: com.coralogixapis.views.v1.services.GetViewRequest
	(*GetViewResponse)(nil),        // 5: com.coralogixapis.views.v1.services.GetViewResponse
	(*DeleteViewRequest)(nil),      // 6: com.coralogixapis.views.v1.services.DeleteViewRequest
	(*DeleteViewResponse)(nil),     // 7: com.coralogixapis.views.v1.services.DeleteViewResponse
	(*ListViewsRequest)(nil),       // 8: com.coralogixapis.views.v1.services.ListViewsRequest
	(*ListViewsResponse)(nil),      // 9: com.coralogixapis.views.v1.services.ListViewsResponse
	(*View)(nil),                   // 10: com.coralogixapis.views.v1.services.View
	(*wrapperspb.StringValue)(nil), // 11: google.protobuf.StringValue
	(*v1.SearchQuery)(nil),         // 12: com.coralogixapis.views.v1.SearchQuery
	(*v1.TimeSelection)(nil),       // 13: com.coralogixapis.views.v1.TimeSelection
	(*v1.SelectedFilters)(nil),     // 14: com.coralogixapis.views.v1.SelectedFilters
	(*wrapperspb.Int32Value)(nil),  // 15: google.protobuf.Int32Value
	(*wrapperspb.BoolValue)(nil),   // 16: google.protobuf.BoolValue
}
var file_com_coralogixapis_views_v1_services_views_service_proto_depIdxs = []int32{
	11, // 0: com.coralogixapis.views.v1.services.CreateViewRequest.name:type_name -> google.protobuf.StringValue
	12, // 1: com.coralogixapis.views.v1.services.CreateViewRequest.search_query:type_name -> com.coralogixapis.views.v1.SearchQuery
	13, // 2: com.coralogixapis.views.v1.services.CreateViewRequest.time_selection:type_name -> com.coralogixapis.views.v1.TimeSelection
	14, // 3: com.coralogixapis.views.v1.services.CreateViewRequest.filters:type_name -> com.coralogixapis.views.v1.SelectedFilters
	11, // 4: com.coralogixapis.views.v1.services.CreateViewRequest.folder_id:type_name -> google.protobuf.StringValue
	10, // 5: com.coralogixapis.views.v1.services.CreateViewResponse.view:type_name -> com.coralogixapis.views.v1.services.View
	10, // 6: com.coralogixapis.views.v1.services.ReplaceViewRequest.view:type_name -> com.coralogixapis.views.v1.services.View
	10, // 7: com.coralogixapis.views.v1.services.ReplaceViewResponse.view:type_name -> com.coralogixapis.views.v1.services.View
	15, // 8: com.coralogixapis.views.v1.services.GetViewRequest.id:type_name -> google.protobuf.Int32Value
	10, // 9: com.coralogixapis.views.v1.services.GetViewResponse.view:type_name -> com.coralogixapis.views.v1.services.View
	15, // 10: com.coralogixapis.views.v1.services.DeleteViewRequest.id:type_name -> google.protobuf.Int32Value
	10, // 11: com.coralogixapis.views.v1.services.ListViewsResponse.views:type_name -> com.coralogixapis.views.v1.services.View
	15, // 12: com.coralogixapis.views.v1.services.View.id:type_name -> google.protobuf.Int32Value
	11, // 13: com.coralogixapis.views.v1.services.View.name:type_name -> google.protobuf.StringValue
	12, // 14: com.coralogixapis.views.v1.services.View.search_query:type_name -> com.coralogixapis.views.v1.SearchQuery
	13, // 15: com.coralogixapis.views.v1.services.View.time_selection:type_name -> com.coralogixapis.views.v1.TimeSelection
	14, // 16: com.coralogixapis.views.v1.services.View.filters:type_name -> com.coralogixapis.views.v1.SelectedFilters
	11, // 17: com.coralogixapis.views.v1.services.View.folder_id:type_name -> google.protobuf.StringValue
	16, // 18: com.coralogixapis.views.v1.services.View.is_compact_mode:type_name -> google.protobuf.BoolValue
	0,  // 19: com.coralogixapis.views.v1.services.ViewsService.CreateView:input_type -> com.coralogixapis.views.v1.services.CreateViewRequest
	2,  // 20: com.coralogixapis.views.v1.services.ViewsService.ReplaceView:input_type -> com.coralogixapis.views.v1.services.ReplaceViewRequest
	4,  // 21: com.coralogixapis.views.v1.services.ViewsService.GetView:input_type -> com.coralogixapis.views.v1.services.GetViewRequest
	6,  // 22: com.coralogixapis.views.v1.services.ViewsService.DeleteView:input_type -> com.coralogixapis.views.v1.services.DeleteViewRequest
	8,  // 23: com.coralogixapis.views.v1.services.ViewsService.ListViews:input_type -> com.coralogixapis.views.v1.services.ListViewsRequest
	1,  // 24: com.coralogixapis.views.v1.services.ViewsService.CreateView:output_type -> com.coralogixapis.views.v1.services.CreateViewResponse
	3,  // 25: com.coralogixapis.views.v1.services.ViewsService.ReplaceView:output_type -> com.coralogixapis.views.v1.services.ReplaceViewResponse
	5,  // 26: com.coralogixapis.views.v1.services.ViewsService.GetView:output_type -> com.coralogixapis.views.v1.services.GetViewResponse
	7,  // 27: com.coralogixapis.views.v1.services.ViewsService.DeleteView:output_type -> com.coralogixapis.views.v1.services.DeleteViewResponse
	9,  // 28: com.coralogixapis.views.v1.services.ViewsService.ListViews:output_type -> com.coralogixapis.views.v1.services.ListViewsResponse
	24, // [24:29] is the sub-list for method output_type
	19, // [19:24] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_views_v1_services_views_service_proto_init() }
func file_com_coralogixapis_views_v1_services_views_service_proto_init() {
	if File_com_coralogixapis_views_v1_services_views_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_views_v1_services_views_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogixapis_views_v1_services_views_service_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_views_v1_services_views_service_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_views_v1_services_views_service_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_views_v1_services_views_service_proto = out.File
	file_com_coralogixapis_views_v1_services_views_service_proto_rawDesc = nil
	file_com_coralogixapis_views_v1_services_views_service_proto_goTypes = nil
	file_com_coralogixapis_views_v1_services_views_service_proto_depIdxs = nil
}
