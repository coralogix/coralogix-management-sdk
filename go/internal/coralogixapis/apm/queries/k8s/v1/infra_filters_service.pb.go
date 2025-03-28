// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: com/coralogixapis/apm/queries/k8s/v1/infra_filters_service.proto

package v1

import (
	v1 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/global_mapping/v1"
	v2 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/apm/common/v2"
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

type ListFiltersRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	IsInitialLoad   *wrapperspb.BoolValue  `protobuf:"bytes,1,opt,name=is_initial_load,json=isInitialLoad,proto3" json:"is_initial_load,omitempty"`
	Tab             K8SObject              `protobuf:"varint,2,opt,name=tab,proto3,enum=com.coralogixapis.apm.queries.k8s.v1.K8SObject" json:"tab,omitempty"`
	TimeRange       *v2.TimeRange          `protobuf:"bytes,3,opt,name=time_range,json=timeRange,proto3" json:"time_range,omitempty"`
	DataSource      *v1.DataSource         `protobuf:"bytes,4,opt,name=data_source,json=dataSource,proto3" json:"data_source,omitempty"`
	SelectedFilters []*SelectedFilter      `protobuf:"bytes,5,rep,name=selected_filters,json=selectedFilters,proto3" json:"selected_filters,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ListFiltersRequest) Reset() {
	*x = ListFiltersRequest{}
	mi := &file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFiltersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFiltersRequest) ProtoMessage() {}

func (x *ListFiltersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFiltersRequest.ProtoReflect.Descriptor instead.
func (*ListFiltersRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListFiltersRequest) GetIsInitialLoad() *wrapperspb.BoolValue {
	if x != nil {
		return x.IsInitialLoad
	}
	return nil
}

func (x *ListFiltersRequest) GetTab() K8SObject {
	if x != nil {
		return x.Tab
	}
	return K8SObject_K8S_OBJECT_UNSPECIFIED
}

func (x *ListFiltersRequest) GetTimeRange() *v2.TimeRange {
	if x != nil {
		return x.TimeRange
	}
	return nil
}

func (x *ListFiltersRequest) GetDataSource() *v1.DataSource {
	if x != nil {
		return x.DataSource
	}
	return nil
}

func (x *ListFiltersRequest) GetSelectedFilters() []*SelectedFilter {
	if x != nil {
		return x.SelectedFilters
	}
	return nil
}

type ListFiltersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tab           K8SObject              `protobuf:"varint,1,opt,name=tab,proto3,enum=com.coralogixapis.apm.queries.k8s.v1.K8SObject" json:"tab,omitempty"`
	Filters       []*Filter              `protobuf:"bytes,2,rep,name=filters,proto3" json:"filters,omitempty"`
	DataSource    *v1.DataSource         `protobuf:"bytes,3,opt,name=data_source,json=dataSource,proto3" json:"data_source,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListFiltersResponse) Reset() {
	*x = ListFiltersResponse{}
	mi := &file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFiltersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFiltersResponse) ProtoMessage() {}

func (x *ListFiltersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFiltersResponse.ProtoReflect.Descriptor instead.
func (*ListFiltersResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListFiltersResponse) GetTab() K8SObject {
	if x != nil {
		return x.Tab
	}
	return K8SObject_K8S_OBJECT_UNSPECIFIED
}

func (x *ListFiltersResponse) GetFilters() []*Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *ListFiltersResponse) GetDataSource() *v1.DataSource {
	if x != nil {
		return x.DataSource
	}
	return nil
}

var File_com_coralogixapis_apm_queries_k8s_v1_infra_filters_service_proto protoreflect.FileDescriptor

var file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_service_proto_rawDesc = []byte{
	0x0a, 0x40, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2f,
	0x6b, 0x38, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65,
	0x73, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70,
	0x6d, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x6b, 0x38, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x71, 0x75, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x2f, 0x6b, 0x38, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x03, 0x0a, 0x12,
	0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x42, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c,
	0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x69, 0x73, 0x49, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x41, 0x0a, 0x03, 0x74, 0x61, 0x62, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x71, 0x75, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x38, 0x73, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x03, 0x74, 0x61, 0x62, 0x12, 0x49, 0x0a, 0x0a, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x5f, 0x0a, 0x10, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x6b, 0x38, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x0f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x22, 0xee, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x03, 0x74,
	0x61, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d,
	0x2e, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4b, 0x38, 0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x03, 0x74, 0x61, 0x62, 0x12, 0x46,
	0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e,
	0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x4c, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x32, 0x9a, 0x01, 0x0a, 0x13, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x82, 0x01, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x38, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x6b, 0x38, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x71,
	0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_service_proto_rawDescOnce sync.Once
	file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_service_proto_rawDescData = file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_service_proto_rawDesc
)

func file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_service_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_service_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_service_proto_rawDescData)
	})
	return file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_service_proto_rawDescData
}

var file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_service_proto_goTypes = []any{
	(*ListFiltersRequest)(nil),   // 0: com.coralogixapis.apm.queries.k8s.v1.ListFiltersRequest
	(*ListFiltersResponse)(nil),  // 1: com.coralogixapis.apm.queries.k8s.v1.ListFiltersResponse
	(*wrapperspb.BoolValue)(nil), // 2: google.protobuf.BoolValue
	(K8SObject)(0),               // 3: com.coralogixapis.apm.queries.k8s.v1.K8sObject
	(*v2.TimeRange)(nil),         // 4: com.coralogixapis.apm.common.v2.TimeRange
	(*v1.DataSource)(nil),        // 5: com.coralogix.global_mapping.v1.DataSource
	(*SelectedFilter)(nil),       // 6: com.coralogixapis.apm.queries.k8s.v1.SelectedFilter
	(*Filter)(nil),               // 7: com.coralogixapis.apm.queries.k8s.v1.Filter
}
var file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_service_proto_depIdxs = []int32{
	2, // 0: com.coralogixapis.apm.queries.k8s.v1.ListFiltersRequest.is_initial_load:type_name -> google.protobuf.BoolValue
	3, // 1: com.coralogixapis.apm.queries.k8s.v1.ListFiltersRequest.tab:type_name -> com.coralogixapis.apm.queries.k8s.v1.K8sObject
	4, // 2: com.coralogixapis.apm.queries.k8s.v1.ListFiltersRequest.time_range:type_name -> com.coralogixapis.apm.common.v2.TimeRange
	5, // 3: com.coralogixapis.apm.queries.k8s.v1.ListFiltersRequest.data_source:type_name -> com.coralogix.global_mapping.v1.DataSource
	6, // 4: com.coralogixapis.apm.queries.k8s.v1.ListFiltersRequest.selected_filters:type_name -> com.coralogixapis.apm.queries.k8s.v1.SelectedFilter
	3, // 5: com.coralogixapis.apm.queries.k8s.v1.ListFiltersResponse.tab:type_name -> com.coralogixapis.apm.queries.k8s.v1.K8sObject
	7, // 6: com.coralogixapis.apm.queries.k8s.v1.ListFiltersResponse.filters:type_name -> com.coralogixapis.apm.queries.k8s.v1.Filter
	5, // 7: com.coralogixapis.apm.queries.k8s.v1.ListFiltersResponse.data_source:type_name -> com.coralogix.global_mapping.v1.DataSource
	0, // 8: com.coralogixapis.apm.queries.k8s.v1.InfraFiltersService.ListFilters:input_type -> com.coralogixapis.apm.queries.k8s.v1.ListFiltersRequest
	1, // 9: com.coralogixapis.apm.queries.k8s.v1.InfraFiltersService.ListFilters:output_type -> com.coralogixapis.apm.queries.k8s.v1.ListFiltersResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_service_proto_init() }
func file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_service_proto_init() {
	if File_com_coralogixapis_apm_queries_k8s_v1_infra_filters_service_proto != nil {
		return
	}
	file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_init()
	file_com_coralogixapis_apm_queries_k8s_v1_objects_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_service_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_service_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_service_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_apm_queries_k8s_v1_infra_filters_service_proto = out.File
	file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_service_proto_rawDesc = nil
	file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_service_proto_goTypes = nil
	file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_service_proto_depIdxs = nil
}
