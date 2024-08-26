// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.3
// source: com/coralogixapis/apm/requirements/v2/feature_requirements_internal_service.proto

package v2

import (
	v2 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/apm/common/v2"
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

type ValidateFeatureRequirementsInternalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feature     Feature          `protobuf:"varint,1,opt,name=feature,proto3,enum=com.coralogixapis.apm.requirements.v2.Feature" json:"feature,omitempty"`
	DataSources []*v2.DataSource `protobuf:"bytes,2,rep,name=data_sources,json=dataSources,proto3" json:"data_sources,omitempty"`
}

func (x *ValidateFeatureRequirementsInternalRequest) Reset() {
	*x = ValidateFeatureRequirementsInternalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateFeatureRequirementsInternalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateFeatureRequirementsInternalRequest) ProtoMessage() {}

func (x *ValidateFeatureRequirementsInternalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateFeatureRequirementsInternalRequest.ProtoReflect.Descriptor instead.
func (*ValidateFeatureRequirementsInternalRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_rawDescGZIP(), []int{0}
}

func (x *ValidateFeatureRequirementsInternalRequest) GetFeature() Feature {
	if x != nil {
		return x.Feature
	}
	return Feature_FEATURE_UNSPECIFIED
}

func (x *ValidateFeatureRequirementsInternalRequest) GetDataSources() []*v2.DataSource {
	if x != nil {
		return x.DataSources
	}
	return nil
}

type ValidateFeatureRequirementsInternalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Requirements []*DataSourceRequirements `protobuf:"bytes,1,rep,name=requirements,proto3" json:"requirements,omitempty"`
}

func (x *ValidateFeatureRequirementsInternalResponse) Reset() {
	*x = ValidateFeatureRequirementsInternalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateFeatureRequirementsInternalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateFeatureRequirementsInternalResponse) ProtoMessage() {}

func (x *ValidateFeatureRequirementsInternalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateFeatureRequirementsInternalResponse.ProtoReflect.Descriptor instead.
func (*ValidateFeatureRequirementsInternalResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_rawDescGZIP(), []int{1}
}

func (x *ValidateFeatureRequirementsInternalResponse) GetRequirements() []*DataSourceRequirements {
	if x != nil {
		return x.Requirements
	}
	return nil
}

type ListFeatureRequirementsInternalDataSourcesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feature Feature `protobuf:"varint,1,opt,name=feature,proto3,enum=com.coralogixapis.apm.requirements.v2.Feature" json:"feature,omitempty"`
}

func (x *ListFeatureRequirementsInternalDataSourcesRequest) Reset() {
	*x = ListFeatureRequirementsInternalDataSourcesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFeatureRequirementsInternalDataSourcesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFeatureRequirementsInternalDataSourcesRequest) ProtoMessage() {}

func (x *ListFeatureRequirementsInternalDataSourcesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFeatureRequirementsInternalDataSourcesRequest.ProtoReflect.Descriptor instead.
func (*ListFeatureRequirementsInternalDataSourcesRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListFeatureRequirementsInternalDataSourcesRequest) GetFeature() Feature {
	if x != nil {
		return x.Feature
	}
	return Feature_FEATURE_UNSPECIFIED
}

type ListFeatureRequirementsInternalDataSourcesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataSources []*v2.DataSource `protobuf:"bytes,1,rep,name=data_sources,json=dataSources,proto3" json:"data_sources,omitempty"`
}

func (x *ListFeatureRequirementsInternalDataSourcesResponse) Reset() {
	*x = ListFeatureRequirementsInternalDataSourcesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFeatureRequirementsInternalDataSourcesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFeatureRequirementsInternalDataSourcesResponse) ProtoMessage() {}

func (x *ListFeatureRequirementsInternalDataSourcesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFeatureRequirementsInternalDataSourcesResponse.ProtoReflect.Descriptor instead.
func (*ListFeatureRequirementsInternalDataSourcesResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListFeatureRequirementsInternalDataSourcesResponse) GetDataSources() []*v2.DataSource {
	if x != nil {
		return x.DataSources
	}
	return nil
}

var File_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto protoreflect.FileDescriptor

var file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_rawDesc = []byte{
	0x0a, 0x51, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x32, 0x1a, 0x2f, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70,
	0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x44,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x01, 0x0a, 0x2a, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x07, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61,
	0x70, 0x6d, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x76, 0x32, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x07, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x4e, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70,
	0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x22, 0x90, 0x01, 0x0a, 0x2b, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70,
	0x6d, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x7d, 0x0a, 0x31, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x07, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x07, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x84, 0x01, 0x0a, 0x32, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0c,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x0b, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x32, 0xaa, 0x04, 0x0a,
	0x22, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0xf1, 0x01, 0x0a, 0x23, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x51, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x61, 0x70, 0x6d, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x76, 0x32, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x52,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x23, 0xc2, 0xb8, 0x02, 0x1f, 0x0a, 0x1d, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x20, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x20, 0x52, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x8f, 0x02, 0x0a, 0x2a, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x58, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x59, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0xc2, 0xb8, 0x02,
	0x28, 0x0a, 0x26, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x20,
	0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x20, 0x44, 0x61, 0x74,
	0x61, 0x20, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_rawDescOnce sync.Once
	file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_rawDescData = file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_rawDesc
)

func file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_rawDescData)
	})
	return file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_rawDescData
}

var file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_goTypes = []interface{}{
	(*ValidateFeatureRequirementsInternalRequest)(nil),         // 0: com.coralogixapis.apm.requirements.v2.ValidateFeatureRequirementsInternalRequest
	(*ValidateFeatureRequirementsInternalResponse)(nil),        // 1: com.coralogixapis.apm.requirements.v2.ValidateFeatureRequirementsInternalResponse
	(*ListFeatureRequirementsInternalDataSourcesRequest)(nil),  // 2: com.coralogixapis.apm.requirements.v2.ListFeatureRequirementsInternalDataSourcesRequest
	(*ListFeatureRequirementsInternalDataSourcesResponse)(nil), // 3: com.coralogixapis.apm.requirements.v2.ListFeatureRequirementsInternalDataSourcesResponse
	(Feature)(0),                   // 4: com.coralogixapis.apm.requirements.v2.Feature
	(*v2.DataSource)(nil),          // 5: com.coralogixapis.apm.common.v2.DataSource
	(*DataSourceRequirements)(nil), // 6: com.coralogixapis.apm.requirements.v2.DataSourceRequirements
}
var file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_depIdxs = []int32{
	4, // 0: com.coralogixapis.apm.requirements.v2.ValidateFeatureRequirementsInternalRequest.feature:type_name -> com.coralogixapis.apm.requirements.v2.Feature
	5, // 1: com.coralogixapis.apm.requirements.v2.ValidateFeatureRequirementsInternalRequest.data_sources:type_name -> com.coralogixapis.apm.common.v2.DataSource
	6, // 2: com.coralogixapis.apm.requirements.v2.ValidateFeatureRequirementsInternalResponse.requirements:type_name -> com.coralogixapis.apm.requirements.v2.DataSourceRequirements
	4, // 3: com.coralogixapis.apm.requirements.v2.ListFeatureRequirementsInternalDataSourcesRequest.feature:type_name -> com.coralogixapis.apm.requirements.v2.Feature
	5, // 4: com.coralogixapis.apm.requirements.v2.ListFeatureRequirementsInternalDataSourcesResponse.data_sources:type_name -> com.coralogixapis.apm.common.v2.DataSource
	0, // 5: com.coralogixapis.apm.requirements.v2.FeatureRequirementsInternalService.ValidateFeatureRequirementsInternal:input_type -> com.coralogixapis.apm.requirements.v2.ValidateFeatureRequirementsInternalRequest
	2, // 6: com.coralogixapis.apm.requirements.v2.FeatureRequirementsInternalService.ListFeatureRequirementsInternalDataSources:input_type -> com.coralogixapis.apm.requirements.v2.ListFeatureRequirementsInternalDataSourcesRequest
	1, // 7: com.coralogixapis.apm.requirements.v2.FeatureRequirementsInternalService.ValidateFeatureRequirementsInternal:output_type -> com.coralogixapis.apm.requirements.v2.ValidateFeatureRequirementsInternalResponse
	3, // 8: com.coralogixapis.apm.requirements.v2.FeatureRequirementsInternalService.ListFeatureRequirementsInternalDataSources:output_type -> com.coralogixapis.apm.requirements.v2.ListFeatureRequirementsInternalDataSourcesResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_init()
}
func file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_init() {
	if File_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto != nil {
		return
	}
	file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_init()
	file_com_coralogixapis_apm_requirements_v2_features_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateFeatureRequirementsInternalRequest); i {
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
		file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateFeatureRequirementsInternalResponse); i {
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
		file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFeatureRequirementsInternalDataSourcesRequest); i {
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
		file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFeatureRequirementsInternalDataSourcesResponse); i {
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
			RawDescriptor: file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto = out.File
	file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_rawDesc = nil
	file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_goTypes = nil
	file_com_coralogixapis_apm_requirements_v2_feature_requirements_internal_service_proto_depIdxs = nil
}
