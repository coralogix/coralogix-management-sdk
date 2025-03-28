// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: com/coralogix/global_mapping/v1/providers_service.proto

package v1

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

type GetCompanyProvidersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCompanyProvidersRequest) Reset() {
	*x = GetCompanyProvidersRequest{}
	mi := &file_com_coralogix_global_mapping_v1_providers_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCompanyProvidersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompanyProvidersRequest) ProtoMessage() {}

func (x *GetCompanyProvidersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_global_mapping_v1_providers_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCompanyProvidersRequest.ProtoReflect.Descriptor instead.
func (*GetCompanyProvidersRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_global_mapping_v1_providers_service_proto_rawDescGZIP(), []int{0}
}

type GetCompanyProvidersResponse struct {
	state            protoimpl.MessageState    `protogen:"open.v1"`
	CompanyProviders []*wrapperspb.StringValue `protobuf:"bytes,1,rep,name=company_providers,json=companyProviders,proto3" json:"company_providers,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *GetCompanyProvidersResponse) Reset() {
	*x = GetCompanyProvidersResponse{}
	mi := &file_com_coralogix_global_mapping_v1_providers_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCompanyProvidersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompanyProvidersResponse) ProtoMessage() {}

func (x *GetCompanyProvidersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_global_mapping_v1_providers_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCompanyProvidersResponse.ProtoReflect.Descriptor instead.
func (*GetCompanyProvidersResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_global_mapping_v1_providers_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetCompanyProvidersResponse) GetCompanyProviders() []*wrapperspb.StringValue {
	if x != nil {
		return x.CompanyProviders
	}
	return nil
}

var File_com_coralogix_global_mapping_v1_providers_service_proto protoreflect.FileDescriptor

var file_com_coralogix_global_mapping_v1_providers_service_proto_rawDesc = []byte{
	0x0a, 0x37, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f,
	0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x1c, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x68, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x49, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x32, 0xc9, 0x01, 0x0a, 0x17,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xad, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x12,
	0x3b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0xc2, 0xb8, 0x02, 0x17,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x20, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x20, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_global_mapping_v1_providers_service_proto_rawDescOnce sync.Once
	file_com_coralogix_global_mapping_v1_providers_service_proto_rawDescData = file_com_coralogix_global_mapping_v1_providers_service_proto_rawDesc
)

func file_com_coralogix_global_mapping_v1_providers_service_proto_rawDescGZIP() []byte {
	file_com_coralogix_global_mapping_v1_providers_service_proto_rawDescOnce.Do(func() {
		file_com_coralogix_global_mapping_v1_providers_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_global_mapping_v1_providers_service_proto_rawDescData)
	})
	return file_com_coralogix_global_mapping_v1_providers_service_proto_rawDescData
}

var file_com_coralogix_global_mapping_v1_providers_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogix_global_mapping_v1_providers_service_proto_goTypes = []any{
	(*GetCompanyProvidersRequest)(nil),  // 0: com.coralogix.global_mapping.v1.GetCompanyProvidersRequest
	(*GetCompanyProvidersResponse)(nil), // 1: com.coralogix.global_mapping.v1.GetCompanyProvidersResponse
	(*wrapperspb.StringValue)(nil),      // 2: google.protobuf.StringValue
}
var file_com_coralogix_global_mapping_v1_providers_service_proto_depIdxs = []int32{
	2, // 0: com.coralogix.global_mapping.v1.GetCompanyProvidersResponse.company_providers:type_name -> google.protobuf.StringValue
	0, // 1: com.coralogix.global_mapping.v1.CompanyProvidersService.GetCompanyProviders:input_type -> com.coralogix.global_mapping.v1.GetCompanyProvidersRequest
	1, // 2: com.coralogix.global_mapping.v1.CompanyProvidersService.GetCompanyProviders:output_type -> com.coralogix.global_mapping.v1.GetCompanyProvidersResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_com_coralogix_global_mapping_v1_providers_service_proto_init() }
func file_com_coralogix_global_mapping_v1_providers_service_proto_init() {
	if File_com_coralogix_global_mapping_v1_providers_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_global_mapping_v1_providers_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogix_global_mapping_v1_providers_service_proto_goTypes,
		DependencyIndexes: file_com_coralogix_global_mapping_v1_providers_service_proto_depIdxs,
		MessageInfos:      file_com_coralogix_global_mapping_v1_providers_service_proto_msgTypes,
	}.Build()
	File_com_coralogix_global_mapping_v1_providers_service_proto = out.File
	file_com_coralogix_global_mapping_v1_providers_service_proto_rawDesc = nil
	file_com_coralogix_global_mapping_v1_providers_service_proto_goTypes = nil
	file_com_coralogix_global_mapping_v1_providers_service_proto_depIdxs = nil
}
