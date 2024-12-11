// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.1
// source: com/coralogix/extensions/v1/extension_service.proto

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

type GetAllExtensionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IncludeHiddenExtensions *wrapperspb.BoolValue           `protobuf:"bytes,1,opt,name=include_hidden_extensions,json=includeHiddenExtensions,proto3" json:"include_hidden_extensions,omitempty"`
	Filter                  *GetAllExtensionsRequest_Filter `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *GetAllExtensionsRequest) Reset() {
	*x = GetAllExtensionsRequest{}
	mi := &file_com_coralogix_extensions_v1_extension_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllExtensionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllExtensionsRequest) ProtoMessage() {}

func (x *GetAllExtensionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_extensions_v1_extension_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllExtensionsRequest.ProtoReflect.Descriptor instead.
func (*GetAllExtensionsRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_extensions_v1_extension_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllExtensionsRequest) GetIncludeHiddenExtensions() *wrapperspb.BoolValue {
	if x != nil {
		return x.IncludeHiddenExtensions
	}
	return nil
}

func (x *GetAllExtensionsRequest) GetFilter() *GetAllExtensionsRequest_Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type GetAllExtensionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Extensions []*GetAllExtensionsResponse_Extension `protobuf:"bytes,1,rep,name=extensions,proto3" json:"extensions,omitempty"`
}

func (x *GetAllExtensionsResponse) Reset() {
	*x = GetAllExtensionsResponse{}
	mi := &file_com_coralogix_extensions_v1_extension_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllExtensionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllExtensionsResponse) ProtoMessage() {}

func (x *GetAllExtensionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_extensions_v1_extension_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllExtensionsResponse.ProtoReflect.Descriptor instead.
func (*GetAllExtensionsResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_extensions_v1_extension_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllExtensionsResponse) GetExtensions() []*GetAllExtensionsResponse_Extension {
	if x != nil {
		return x.Extensions
	}
	return nil
}

type GetExtensionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// On the frontend, we don't need the dashboard binaries.
	// But we need them for deployment in the extensions-service - so it's kind of temporary argument,
	// so as soon all the deployment logic is moved to the extensions-api, it can be removed.
	IncludeDashboardBinaries *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=include_dashboard_binaries,json=includeDashboardBinaries,proto3" json:"include_dashboard_binaries,omitempty"`
	IncludeTestingRevision   *wrapperspb.BoolValue `protobuf:"bytes,3,opt,name=include_testing_revision,json=includeTestingRevision,proto3" json:"include_testing_revision,omitempty"`
}

func (x *GetExtensionRequest) Reset() {
	*x = GetExtensionRequest{}
	mi := &file_com_coralogix_extensions_v1_extension_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetExtensionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExtensionRequest) ProtoMessage() {}

func (x *GetExtensionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_extensions_v1_extension_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExtensionRequest.ProtoReflect.Descriptor instead.
func (*GetExtensionRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_extensions_v1_extension_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetExtensionRequest) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *GetExtensionRequest) GetIncludeDashboardBinaries() *wrapperspb.BoolValue {
	if x != nil {
		return x.IncludeDashboardBinaries
	}
	return nil
}

func (x *GetExtensionRequest) GetIncludeTestingRevision() *wrapperspb.BoolValue {
	if x != nil {
		return x.IncludeTestingRevision
	}
	return nil
}

type GetExtensionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Extension *Extension `protobuf:"bytes,1,opt,name=extension,proto3" json:"extension,omitempty"`
}

func (x *GetExtensionResponse) Reset() {
	*x = GetExtensionResponse{}
	mi := &file_com_coralogix_extensions_v1_extension_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetExtensionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExtensionResponse) ProtoMessage() {}

func (x *GetExtensionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_extensions_v1_extension_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExtensionResponse.ProtoReflect.Descriptor instead.
func (*GetExtensionResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_extensions_v1_extension_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetExtensionResponse) GetExtension() *Extension {
	if x != nil {
		return x.Extension
	}
	return nil
}

type GetAllExtensionsRequest_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Integrations []string `protobuf:"bytes,1,rep,name=integrations,proto3" json:"integrations,omitempty"`
}

func (x *GetAllExtensionsRequest_Filter) Reset() {
	*x = GetAllExtensionsRequest_Filter{}
	mi := &file_com_coralogix_extensions_v1_extension_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllExtensionsRequest_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllExtensionsRequest_Filter) ProtoMessage() {}

func (x *GetAllExtensionsRequest_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_extensions_v1_extension_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllExtensionsRequest_Filter.ProtoReflect.Descriptor instead.
func (*GetAllExtensionsRequest_Filter) Descriptor() ([]byte, []int) {
	return file_com_coralogix_extensions_v1_extension_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *GetAllExtensionsRequest_Filter) GetIntegrations() []string {
	if x != nil {
		return x.Integrations
	}
	return nil
}

type GetAllExtensionsResponse_RevisionSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemCounts *ItemCounts           `protobuf:"bytes,1,opt,name=item_counts,json=itemCounts,proto3" json:"item_counts,omitempty"`
	IsNew      *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=is_new,json=isNew,proto3" json:"is_new,omitempty"`
}

func (x *GetAllExtensionsResponse_RevisionSummary) Reset() {
	*x = GetAllExtensionsResponse_RevisionSummary{}
	mi := &file_com_coralogix_extensions_v1_extension_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllExtensionsResponse_RevisionSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllExtensionsResponse_RevisionSummary) ProtoMessage() {}

func (x *GetAllExtensionsResponse_RevisionSummary) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_extensions_v1_extension_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllExtensionsResponse_RevisionSummary.ProtoReflect.Descriptor instead.
func (*GetAllExtensionsResponse_RevisionSummary) Descriptor() ([]byte, []int) {
	return file_com_coralogix_extensions_v1_extension_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetAllExtensionsResponse_RevisionSummary) GetItemCounts() *ItemCounts {
	if x != nil {
		return x.ItemCounts
	}
	return nil
}

func (x *GetAllExtensionsResponse_RevisionSummary) GetIsNew() *wrapperspb.BoolValue {
	if x != nil {
		return x.IsNew
	}
	return nil
}

type GetAllExtensionsResponse_Revision struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version            *wrapperspb.StringValue                   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Description        *wrapperspb.StringValue                   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Excerpt            *wrapperspb.StringValue                   `protobuf:"bytes,3,opt,name=excerpt,proto3" json:"excerpt,omitempty"`
	Labels             []*wrapperspb.StringValue                 `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty"`
	IntegrationDetails []*IntegrationDetail                      `protobuf:"bytes,5,rep,name=integration_details,json=integrationDetails,proto3" json:"integration_details,omitempty"`
	Summary            *GetAllExtensionsResponse_RevisionSummary `protobuf:"bytes,6,opt,name=summary,proto3" json:"summary,omitempty"`
}

func (x *GetAllExtensionsResponse_Revision) Reset() {
	*x = GetAllExtensionsResponse_Revision{}
	mi := &file_com_coralogix_extensions_v1_extension_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllExtensionsResponse_Revision) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllExtensionsResponse_Revision) ProtoMessage() {}

func (x *GetAllExtensionsResponse_Revision) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_extensions_v1_extension_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllExtensionsResponse_Revision.ProtoReflect.Descriptor instead.
func (*GetAllExtensionsResponse_Revision) Descriptor() ([]byte, []int) {
	return file_com_coralogix_extensions_v1_extension_service_proto_rawDescGZIP(), []int{1, 1}
}

func (x *GetAllExtensionsResponse_Revision) GetVersion() *wrapperspb.StringValue {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *GetAllExtensionsResponse_Revision) GetDescription() *wrapperspb.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *GetAllExtensionsResponse_Revision) GetExcerpt() *wrapperspb.StringValue {
	if x != nil {
		return x.Excerpt
	}
	return nil
}

func (x *GetAllExtensionsResponse_Revision) GetLabels() []*wrapperspb.StringValue {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *GetAllExtensionsResponse_Revision) GetIntegrationDetails() []*IntegrationDetail {
	if x != nil {
		return x.IntegrationDetails
	}
	return nil
}

func (x *GetAllExtensionsResponse_Revision) GetSummary() *GetAllExtensionsResponse_RevisionSummary {
	if x != nil {
		return x.Summary
	}
	return nil
}

type GetAllExtensionsResponse_Extension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            *wrapperspb.StringValue              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          *wrapperspb.StringValue              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Image         *wrapperspb.StringValue              `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	DarkModeImage *wrapperspb.StringValue              `protobuf:"bytes,4,opt,name=dark_mode_image,json=darkModeImage,proto3" json:"dark_mode_image,omitempty"`
	Revisions     []*GetAllExtensionsResponse_Revision `protobuf:"bytes,5,rep,name=revisions,proto3" json:"revisions,omitempty"`
	IsHidden      *wrapperspb.BoolValue                `protobuf:"bytes,6,opt,name=is_hidden,json=isHidden,proto3" json:"is_hidden,omitempty"`
	Integrations  []string                             `protobuf:"bytes,7,rep,name=integrations,proto3" json:"integrations,omitempty"`
	Keywords      []string                             `protobuf:"bytes,8,rep,name=keywords,proto3" json:"keywords,omitempty"`
	Deprecation   *Deprecation                         `protobuf:"bytes,9,opt,name=deprecation,proto3" json:"deprecation,omitempty"`
}

func (x *GetAllExtensionsResponse_Extension) Reset() {
	*x = GetAllExtensionsResponse_Extension{}
	mi := &file_com_coralogix_extensions_v1_extension_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllExtensionsResponse_Extension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllExtensionsResponse_Extension) ProtoMessage() {}

func (x *GetAllExtensionsResponse_Extension) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_extensions_v1_extension_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllExtensionsResponse_Extension.ProtoReflect.Descriptor instead.
func (*GetAllExtensionsResponse_Extension) Descriptor() ([]byte, []int) {
	return file_com_coralogix_extensions_v1_extension_service_proto_rawDescGZIP(), []int{1, 2}
}

func (x *GetAllExtensionsResponse_Extension) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *GetAllExtensionsResponse_Extension) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *GetAllExtensionsResponse_Extension) GetImage() *wrapperspb.StringValue {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *GetAllExtensionsResponse_Extension) GetDarkModeImage() *wrapperspb.StringValue {
	if x != nil {
		return x.DarkModeImage
	}
	return nil
}

func (x *GetAllExtensionsResponse_Extension) GetRevisions() []*GetAllExtensionsResponse_Revision {
	if x != nil {
		return x.Revisions
	}
	return nil
}

func (x *GetAllExtensionsResponse_Extension) GetIsHidden() *wrapperspb.BoolValue {
	if x != nil {
		return x.IsHidden
	}
	return nil
}

func (x *GetAllExtensionsResponse_Extension) GetIntegrations() []string {
	if x != nil {
		return x.Integrations
	}
	return nil
}

func (x *GetAllExtensionsResponse_Extension) GetKeywords() []string {
	if x != nil {
		return x.Keywords
	}
	return nil
}

func (x *GetAllExtensionsResponse_Extension) GetDeprecation() *Deprecation {
	if x != nil {
		return x.Deprecation
	}
	return nil
}

var File_com_coralogix_extensions_v1_extension_service_proto protoreflect.FileDescriptor

var file_com_coralogix_extensions_v1_extension_service_proto_rawDesc = []byte{
	0x0a, 0x33, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x1a, 0x2b, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2b, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4, 0x01, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x56, 0x0a, 0x19, 0x69, 0x6e, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x5f, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x17, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65,
	0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x53, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x2c, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0xcc, 0x09, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x5f, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x1a, 0x8e, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x48, 0x0a, 0x0b, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x52, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12,
	0x31, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x6e, 0x65, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x69, 0x73, 0x4e,
	0x65, 0x77, 0x1a, 0xb2, 0x03, 0x0a, 0x08, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x36, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x07, 0x65, 0x78, 0x63, 0x65, 0x72,
	0x70, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x65, 0x78, 0x63, 0x65, 0x72, 0x70, 0x74, 0x12,
	0x34, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x5f, 0x0a, 0x13, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x12, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x5f, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x07,
	0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x1a, 0x88, 0x04, 0x0a, 0x09, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x64, 0x61, 0x72,
	0x6b, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0d, 0x64, 0x61, 0x72, 0x6b, 0x4d, 0x6f, 0x64, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x5c, 0x0a, 0x09, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x09, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x37, 0x0a,
	0x09, 0x69, 0x73, 0x5f, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x69, 0x73,
	0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x4a, 0x0a, 0x0b, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0xf3, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x69, 0x64, 0x12, 0x58, 0x0a, 0x1a, 0x69, 0x6e, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x5f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x62, 0x69,
	0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x18, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x54, 0x0a, 0x18, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x74, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x16, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x5c, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x44, 0x0a, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0x82, 0x03, 0x0a, 0x10, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xb6, 0x01, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0xc2,
	0xb8, 0x02, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a,
	0x22, 0x12, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x6c, 0x6c, 0x12, 0xb4, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3f, 0xc2, 0xb8, 0x02, 0x15,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x20, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x20,
	0x62, 0x79, 0x20, 0x69, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x62, 0x09, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_extensions_v1_extension_service_proto_rawDescOnce sync.Once
	file_com_coralogix_extensions_v1_extension_service_proto_rawDescData = file_com_coralogix_extensions_v1_extension_service_proto_rawDesc
)

func file_com_coralogix_extensions_v1_extension_service_proto_rawDescGZIP() []byte {
	file_com_coralogix_extensions_v1_extension_service_proto_rawDescOnce.Do(func() {
		file_com_coralogix_extensions_v1_extension_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_extensions_v1_extension_service_proto_rawDescData)
	})
	return file_com_coralogix_extensions_v1_extension_service_proto_rawDescData
}

var file_com_coralogix_extensions_v1_extension_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_com_coralogix_extensions_v1_extension_service_proto_goTypes = []any{
	(*GetAllExtensionsRequest)(nil),                  // 0: com.coralogix.extensions.v1.GetAllExtensionsRequest
	(*GetAllExtensionsResponse)(nil),                 // 1: com.coralogix.extensions.v1.GetAllExtensionsResponse
	(*GetExtensionRequest)(nil),                      // 2: com.coralogix.extensions.v1.GetExtensionRequest
	(*GetExtensionResponse)(nil),                     // 3: com.coralogix.extensions.v1.GetExtensionResponse
	(*GetAllExtensionsRequest_Filter)(nil),           // 4: com.coralogix.extensions.v1.GetAllExtensionsRequest.Filter
	(*GetAllExtensionsResponse_RevisionSummary)(nil), // 5: com.coralogix.extensions.v1.GetAllExtensionsResponse.RevisionSummary
	(*GetAllExtensionsResponse_Revision)(nil),        // 6: com.coralogix.extensions.v1.GetAllExtensionsResponse.Revision
	(*GetAllExtensionsResponse_Extension)(nil),       // 7: com.coralogix.extensions.v1.GetAllExtensionsResponse.Extension
	(*wrapperspb.BoolValue)(nil),                     // 8: google.protobuf.BoolValue
	(*wrapperspb.StringValue)(nil),                   // 9: google.protobuf.StringValue
	(*Extension)(nil),                                // 10: com.coralogix.extensions.v1.Extension
	(*ItemCounts)(nil),                               // 11: com.coralogix.extensions.v1.ItemCounts
	(*IntegrationDetail)(nil),                        // 12: com.coralogix.extensions.v1.IntegrationDetail
	(*Deprecation)(nil),                              // 13: com.coralogix.extensions.v1.Deprecation
}
var file_com_coralogix_extensions_v1_extension_service_proto_depIdxs = []int32{
	8,  // 0: com.coralogix.extensions.v1.GetAllExtensionsRequest.include_hidden_extensions:type_name -> google.protobuf.BoolValue
	4,  // 1: com.coralogix.extensions.v1.GetAllExtensionsRequest.filter:type_name -> com.coralogix.extensions.v1.GetAllExtensionsRequest.Filter
	7,  // 2: com.coralogix.extensions.v1.GetAllExtensionsResponse.extensions:type_name -> com.coralogix.extensions.v1.GetAllExtensionsResponse.Extension
	9,  // 3: com.coralogix.extensions.v1.GetExtensionRequest.id:type_name -> google.protobuf.StringValue
	8,  // 4: com.coralogix.extensions.v1.GetExtensionRequest.include_dashboard_binaries:type_name -> google.protobuf.BoolValue
	8,  // 5: com.coralogix.extensions.v1.GetExtensionRequest.include_testing_revision:type_name -> google.protobuf.BoolValue
	10, // 6: com.coralogix.extensions.v1.GetExtensionResponse.extension:type_name -> com.coralogix.extensions.v1.Extension
	11, // 7: com.coralogix.extensions.v1.GetAllExtensionsResponse.RevisionSummary.item_counts:type_name -> com.coralogix.extensions.v1.ItemCounts
	8,  // 8: com.coralogix.extensions.v1.GetAllExtensionsResponse.RevisionSummary.is_new:type_name -> google.protobuf.BoolValue
	9,  // 9: com.coralogix.extensions.v1.GetAllExtensionsResponse.Revision.version:type_name -> google.protobuf.StringValue
	9,  // 10: com.coralogix.extensions.v1.GetAllExtensionsResponse.Revision.description:type_name -> google.protobuf.StringValue
	9,  // 11: com.coralogix.extensions.v1.GetAllExtensionsResponse.Revision.excerpt:type_name -> google.protobuf.StringValue
	9,  // 12: com.coralogix.extensions.v1.GetAllExtensionsResponse.Revision.labels:type_name -> google.protobuf.StringValue
	12, // 13: com.coralogix.extensions.v1.GetAllExtensionsResponse.Revision.integration_details:type_name -> com.coralogix.extensions.v1.IntegrationDetail
	5,  // 14: com.coralogix.extensions.v1.GetAllExtensionsResponse.Revision.summary:type_name -> com.coralogix.extensions.v1.GetAllExtensionsResponse.RevisionSummary
	9,  // 15: com.coralogix.extensions.v1.GetAllExtensionsResponse.Extension.id:type_name -> google.protobuf.StringValue
	9,  // 16: com.coralogix.extensions.v1.GetAllExtensionsResponse.Extension.name:type_name -> google.protobuf.StringValue
	9,  // 17: com.coralogix.extensions.v1.GetAllExtensionsResponse.Extension.image:type_name -> google.protobuf.StringValue
	9,  // 18: com.coralogix.extensions.v1.GetAllExtensionsResponse.Extension.dark_mode_image:type_name -> google.protobuf.StringValue
	6,  // 19: com.coralogix.extensions.v1.GetAllExtensionsResponse.Extension.revisions:type_name -> com.coralogix.extensions.v1.GetAllExtensionsResponse.Revision
	8,  // 20: com.coralogix.extensions.v1.GetAllExtensionsResponse.Extension.is_hidden:type_name -> google.protobuf.BoolValue
	13, // 21: com.coralogix.extensions.v1.GetAllExtensionsResponse.Extension.deprecation:type_name -> com.coralogix.extensions.v1.Deprecation
	0,  // 22: com.coralogix.extensions.v1.ExtensionService.GetAllExtensions:input_type -> com.coralogix.extensions.v1.GetAllExtensionsRequest
	2,  // 23: com.coralogix.extensions.v1.ExtensionService.GetExtension:input_type -> com.coralogix.extensions.v1.GetExtensionRequest
	1,  // 24: com.coralogix.extensions.v1.ExtensionService.GetAllExtensions:output_type -> com.coralogix.extensions.v1.GetAllExtensionsResponse
	3,  // 25: com.coralogix.extensions.v1.ExtensionService.GetExtension:output_type -> com.coralogix.extensions.v1.GetExtensionResponse
	24, // [24:26] is the sub-list for method output_type
	22, // [22:24] is the sub-list for method input_type
	22, // [22:22] is the sub-list for extension type_name
	22, // [22:22] is the sub-list for extension extendee
	0,  // [0:22] is the sub-list for field type_name
}

func init() { file_com_coralogix_extensions_v1_extension_service_proto_init() }
func file_com_coralogix_extensions_v1_extension_service_proto_init() {
	if File_com_coralogix_extensions_v1_extension_service_proto != nil {
		return
	}
	file_com_coralogix_extensions_v1_audit_log_proto_init()
	file_com_coralogix_extensions_v1_extension_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_extensions_v1_extension_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogix_extensions_v1_extension_service_proto_goTypes,
		DependencyIndexes: file_com_coralogix_extensions_v1_extension_service_proto_depIdxs,
		MessageInfos:      file_com_coralogix_extensions_v1_extension_service_proto_msgTypes,
	}.Build()
	File_com_coralogix_extensions_v1_extension_service_proto = out.File
	file_com_coralogix_extensions_v1_extension_service_proto_rawDesc = nil
	file_com_coralogix_extensions_v1_extension_service_proto_goTypes = nil
	file_com_coralogix_extensions_v1_extension_service_proto_depIdxs = nil
}
