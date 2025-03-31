// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogix/extensions/v1/internal_onboarding_service.proto

package v1

import (
	_ "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/common/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
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

type InternalDeployExtensionRequest struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	CompanyId     *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Id            *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InternalDeployExtensionRequest) Reset() {
	*x = InternalDeployExtensionRequest{}
	mi := &file_com_coralogix_extensions_v1_internal_onboarding_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InternalDeployExtensionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalDeployExtensionRequest) ProtoMessage() {}

func (x *InternalDeployExtensionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_extensions_v1_internal_onboarding_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalDeployExtensionRequest.ProtoReflect.Descriptor instead.
func (*InternalDeployExtensionRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_extensions_v1_internal_onboarding_service_proto_rawDescGZIP(), []int{0}
}

func (x *InternalDeployExtensionRequest) GetCompanyId() *wrapperspb.StringValue {
	if x != nil {
		return x.CompanyId
	}
	return nil
}

func (x *InternalDeployExtensionRequest) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

type InternalDeployExtensionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InternalDeployExtensionResponse) Reset() {
	*x = InternalDeployExtensionResponse{}
	mi := &file_com_coralogix_extensions_v1_internal_onboarding_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InternalDeployExtensionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalDeployExtensionResponse) ProtoMessage() {}

func (x *InternalDeployExtensionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_extensions_v1_internal_onboarding_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalDeployExtensionResponse.ProtoReflect.Descriptor instead.
func (*InternalDeployExtensionResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_extensions_v1_internal_onboarding_service_proto_rawDescGZIP(), []int{1}
}

var File_com_coralogix_extensions_v1_internal_onboarding_service_proto protoreflect.FileDescriptor

const file_com_coralogix_extensions_v1_internal_onboarding_service_proto_rawDesc = "" +
	"\n" +
	"=com/coralogix/extensions/v1/internal_onboarding_service.proto\x12\x1bcom.coralogix.extensions.v1\x1a'com/coralogix/common/v1/audit_log.proto\x1a\x1egoogle/protobuf/wrappers.proto\x1a google/protobuf/descriptor.proto\x1a\x1cgoogle/api/annotations.proto\"\x8b\x01\n" +
	"\x1eInternalDeployExtensionRequest\x12;\n" +
	"\n" +
	"company_id\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueR\tcompanyId\x12,\n" +
	"\x02id\x18\x02 \x01(\v2\x1c.google.protobuf.StringValueR\x02id\"!\n" +
	"\x1fInternalDeployExtensionResponse2\xfd\x01\n" +
	"\x19InternalOnboardingService\x12\xdf\x01\n" +
	"\x17InternalDeployExtension\x12;.com.coralogix.extensions.v1.InternalDeployExtensionRequest\x1a<.com.coralogix.extensions.v1.InternalDeployExtensionResponse\"I¸\x02\x12\n" +
	"\x10Deploy Extension\x82\xd3\xe4\x93\x02-:\x01*\"(/internal-onboarding/v1/deploy-extensionb\x06proto3"

var (
	file_com_coralogix_extensions_v1_internal_onboarding_service_proto_rawDescOnce sync.Once
	file_com_coralogix_extensions_v1_internal_onboarding_service_proto_rawDescData []byte
)

func file_com_coralogix_extensions_v1_internal_onboarding_service_proto_rawDescGZIP() []byte {
	file_com_coralogix_extensions_v1_internal_onboarding_service_proto_rawDescOnce.Do(func() {
		file_com_coralogix_extensions_v1_internal_onboarding_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogix_extensions_v1_internal_onboarding_service_proto_rawDesc), len(file_com_coralogix_extensions_v1_internal_onboarding_service_proto_rawDesc)))
	})
	return file_com_coralogix_extensions_v1_internal_onboarding_service_proto_rawDescData
}

var file_com_coralogix_extensions_v1_internal_onboarding_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogix_extensions_v1_internal_onboarding_service_proto_goTypes = []any{
	(*InternalDeployExtensionRequest)(nil),  // 0: com.coralogix.extensions.v1.InternalDeployExtensionRequest
	(*InternalDeployExtensionResponse)(nil), // 1: com.coralogix.extensions.v1.InternalDeployExtensionResponse
	(*wrapperspb.StringValue)(nil),          // 2: google.protobuf.StringValue
}
var file_com_coralogix_extensions_v1_internal_onboarding_service_proto_depIdxs = []int32{
	2, // 0: com.coralogix.extensions.v1.InternalDeployExtensionRequest.company_id:type_name -> google.protobuf.StringValue
	2, // 1: com.coralogix.extensions.v1.InternalDeployExtensionRequest.id:type_name -> google.protobuf.StringValue
	0, // 2: com.coralogix.extensions.v1.InternalOnboardingService.InternalDeployExtension:input_type -> com.coralogix.extensions.v1.InternalDeployExtensionRequest
	1, // 3: com.coralogix.extensions.v1.InternalOnboardingService.InternalDeployExtension:output_type -> com.coralogix.extensions.v1.InternalDeployExtensionResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_com_coralogix_extensions_v1_internal_onboarding_service_proto_init() }
func file_com_coralogix_extensions_v1_internal_onboarding_service_proto_init() {
	if File_com_coralogix_extensions_v1_internal_onboarding_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogix_extensions_v1_internal_onboarding_service_proto_rawDesc), len(file_com_coralogix_extensions_v1_internal_onboarding_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogix_extensions_v1_internal_onboarding_service_proto_goTypes,
		DependencyIndexes: file_com_coralogix_extensions_v1_internal_onboarding_service_proto_depIdxs,
		MessageInfos:      file_com_coralogix_extensions_v1_internal_onboarding_service_proto_msgTypes,
	}.Build()
	File_com_coralogix_extensions_v1_internal_onboarding_service_proto = out.File
	file_com_coralogix_extensions_v1_internal_onboarding_service_proto_goTypes = nil
	file_com_coralogix_extensions_v1_internal_onboarding_service_proto_depIdxs = nil
}
