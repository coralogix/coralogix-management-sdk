// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.26.1
// source: com/coralogix/global_mapping/v1/data_source_internal_service.proto

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

type GetCompanyDataSourcesInternalRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCompanyDataSourcesInternalRequest) Reset() {
	*x = GetCompanyDataSourcesInternalRequest{}
	mi := &file_com_coralogix_global_mapping_v1_data_source_internal_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCompanyDataSourcesInternalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompanyDataSourcesInternalRequest) ProtoMessage() {}

func (x *GetCompanyDataSourcesInternalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_global_mapping_v1_data_source_internal_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCompanyDataSourcesInternalRequest.ProtoReflect.Descriptor instead.
func (*GetCompanyDataSourcesInternalRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_global_mapping_v1_data_source_internal_service_proto_rawDescGZIP(), []int{0}
}

type GetCompanyDataSourcesInternalResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DataSources   []*DataSource          `protobuf:"bytes,1,rep,name=data_sources,json=dataSources,proto3" json:"data_sources,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCompanyDataSourcesInternalResponse) Reset() {
	*x = GetCompanyDataSourcesInternalResponse{}
	mi := &file_com_coralogix_global_mapping_v1_data_source_internal_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCompanyDataSourcesInternalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompanyDataSourcesInternalResponse) ProtoMessage() {}

func (x *GetCompanyDataSourcesInternalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_global_mapping_v1_data_source_internal_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCompanyDataSourcesInternalResponse.ProtoReflect.Descriptor instead.
func (*GetCompanyDataSourcesInternalResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_global_mapping_v1_data_source_internal_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetCompanyDataSourcesInternalResponse) GetDataSources() []*DataSource {
	if x != nil {
		return x.DataSources
	}
	return nil
}

var File_com_coralogix_global_mapping_v1_data_source_internal_service_proto protoreflect.FileDescriptor

var file_com_coralogix_global_mapping_v1_data_source_internal_service_proto_rawDesc = []byte{
	0x0a, 0x42, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x31, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f,
	0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x24, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x77, 0x0a, 0x25, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0c, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0b, 0x64,
	0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x32, 0xf4, 0x01, 0x0a, 0x21, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0xce, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x12, 0x45, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x46, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1e, 0xc2, 0xb8, 0x02, 0x1a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x20, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x20, 0x44, 0x61, 0x74, 0x61, 0x20, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_global_mapping_v1_data_source_internal_service_proto_rawDescOnce sync.Once
	file_com_coralogix_global_mapping_v1_data_source_internal_service_proto_rawDescData = file_com_coralogix_global_mapping_v1_data_source_internal_service_proto_rawDesc
)

func file_com_coralogix_global_mapping_v1_data_source_internal_service_proto_rawDescGZIP() []byte {
	file_com_coralogix_global_mapping_v1_data_source_internal_service_proto_rawDescOnce.Do(func() {
		file_com_coralogix_global_mapping_v1_data_source_internal_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_global_mapping_v1_data_source_internal_service_proto_rawDescData)
	})
	return file_com_coralogix_global_mapping_v1_data_source_internal_service_proto_rawDescData
}

var file_com_coralogix_global_mapping_v1_data_source_internal_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogix_global_mapping_v1_data_source_internal_service_proto_goTypes = []any{
	(*GetCompanyDataSourcesInternalRequest)(nil),  // 0: com.coralogix.global_mapping.v1.GetCompanyDataSourcesInternalRequest
	(*GetCompanyDataSourcesInternalResponse)(nil), // 1: com.coralogix.global_mapping.v1.GetCompanyDataSourcesInternalResponse
	(*DataSource)(nil),                            // 2: com.coralogix.global_mapping.v1.DataSource
}
var file_com_coralogix_global_mapping_v1_data_source_internal_service_proto_depIdxs = []int32{
	2, // 0: com.coralogix.global_mapping.v1.GetCompanyDataSourcesInternalResponse.data_sources:type_name -> com.coralogix.global_mapping.v1.DataSource
	0, // 1: com.coralogix.global_mapping.v1.CompanyDataSourcesInternalService.GetCompanyDataSourcesInternal:input_type -> com.coralogix.global_mapping.v1.GetCompanyDataSourcesInternalRequest
	1, // 2: com.coralogix.global_mapping.v1.CompanyDataSourcesInternalService.GetCompanyDataSourcesInternal:output_type -> com.coralogix.global_mapping.v1.GetCompanyDataSourcesInternalResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_com_coralogix_global_mapping_v1_data_source_internal_service_proto_init() }
func file_com_coralogix_global_mapping_v1_data_source_internal_service_proto_init() {
	if File_com_coralogix_global_mapping_v1_data_source_internal_service_proto != nil {
		return
	}
	file_com_coralogix_global_mapping_v1_data_source_proto_init()
	file_com_coralogix_global_mapping_v1_audit_log_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_global_mapping_v1_data_source_internal_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogix_global_mapping_v1_data_source_internal_service_proto_goTypes,
		DependencyIndexes: file_com_coralogix_global_mapping_v1_data_source_internal_service_proto_depIdxs,
		MessageInfos:      file_com_coralogix_global_mapping_v1_data_source_internal_service_proto_msgTypes,
	}.Build()
	File_com_coralogix_global_mapping_v1_data_source_internal_service_proto = out.File
	file_com_coralogix_global_mapping_v1_data_source_internal_service_proto_rawDesc = nil
	file_com_coralogix_global_mapping_v1_data_source_internal_service_proto_goTypes = nil
	file_com_coralogix_global_mapping_v1_data_source_internal_service_proto_depIdxs = nil
}
