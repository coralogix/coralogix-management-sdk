// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.26.1
// source: com/coralogix/global_mapping/v1/providers_service.proto

package v1

import (
	_ "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/common/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
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

const file_com_coralogix_global_mapping_v1_providers_service_proto_rawDesc = "" +
	"\n" +
	"7com/coralogix/global_mapping/v1/providers_service.proto\x12\x1fcom.coralogix.global_mapping.v1\x1a\x1egoogle/protobuf/wrappers.proto\x1a'com/coralogix/common/v1/audit_log.proto\"\x1c\n" +
	"\x1aGetCompanyProvidersRequest\"h\n" +
	"\x1bGetCompanyProvidersResponse\x12I\n" +
	"\x11company_providers\x18\x01 \x03(\v2\x1c.google.protobuf.StringValueR\x10companyProviders2\xc9\x01\n" +
	"\x17CompanyProvidersService\x12\xad\x01\n" +
	"\x13GetCompanyProviders\x12;.com.coralogix.global_mapping.v1.GetCompanyProvidersRequest\x1a<.com.coralogix.global_mapping.v1.GetCompanyProvidersResponse\"\x1b¸\x02\x17\n" +
	"\x15Get Company Providersb\x06proto3"

var (
	file_com_coralogix_global_mapping_v1_providers_service_proto_rawDescOnce sync.Once
	file_com_coralogix_global_mapping_v1_providers_service_proto_rawDescData []byte
)

func file_com_coralogix_global_mapping_v1_providers_service_proto_rawDescGZIP() []byte {
	file_com_coralogix_global_mapping_v1_providers_service_proto_rawDescOnce.Do(func() {
		file_com_coralogix_global_mapping_v1_providers_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogix_global_mapping_v1_providers_service_proto_rawDesc), len(file_com_coralogix_global_mapping_v1_providers_service_proto_rawDesc)))
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogix_global_mapping_v1_providers_service_proto_rawDesc), len(file_com_coralogix_global_mapping_v1_providers_service_proto_rawDesc)),
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
	file_com_coralogix_global_mapping_v1_providers_service_proto_goTypes = nil
	file_com_coralogix_global_mapping_v1_providers_service_proto_depIdxs = nil
}
