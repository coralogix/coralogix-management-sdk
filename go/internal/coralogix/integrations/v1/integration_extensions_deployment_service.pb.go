// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: com/coralogix/integrations/v1/integration_extensions_deployment_service.proto

package v1

import (
	_ "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/common/v1"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type DeployIntegrationExtensionsRequest struct {
	state          protoimpl.MessageState  `protogen:"open.v1"`
	IntegrationKey *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=integration_key,json=integrationKey,proto3" json:"integration_key,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *DeployIntegrationExtensionsRequest) Reset() {
	*x = DeployIntegrationExtensionsRequest{}
	mi := &file_com_coralogix_integrations_v1_integration_extensions_deployment_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeployIntegrationExtensionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployIntegrationExtensionsRequest) ProtoMessage() {}

func (x *DeployIntegrationExtensionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_integrations_v1_integration_extensions_deployment_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployIntegrationExtensionsRequest.ProtoReflect.Descriptor instead.
func (*DeployIntegrationExtensionsRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_integrations_v1_integration_extensions_deployment_service_proto_rawDescGZIP(), []int{0}
}

func (x *DeployIntegrationExtensionsRequest) GetIntegrationKey() *wrapperspb.StringValue {
	if x != nil {
		return x.IntegrationKey
	}
	return nil
}

type DeployIntegrationExtensionsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeployIntegrationExtensionsResponse) Reset() {
	*x = DeployIntegrationExtensionsResponse{}
	mi := &file_com_coralogix_integrations_v1_integration_extensions_deployment_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeployIntegrationExtensionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployIntegrationExtensionsResponse) ProtoMessage() {}

func (x *DeployIntegrationExtensionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_integrations_v1_integration_extensions_deployment_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployIntegrationExtensionsResponse.ProtoReflect.Descriptor instead.
func (*DeployIntegrationExtensionsResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_integrations_v1_integration_extensions_deployment_service_proto_rawDescGZIP(), []int{1}
}

var File_com_coralogix_integrations_v1_integration_extensions_deployment_service_proto protoreflect.FileDescriptor

var file_com_coralogix_integrations_v1_integration_extensions_deployment_service_proto_rawDesc = []byte{
	0x0a, 0x4d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x27,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf2, 0x01, 0x0a, 0x22, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x45, 0x0a, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x3a, 0x84, 0x01, 0x92, 0x41, 0x80, 0x01, 0x0a, 0x7e,
	0x2a, 0x22, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x32, 0x46, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x74, 0x6f,
	0x20, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x20, 0x61, 0x75, 0x74, 0x6f, 0x2d, 0x61, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x65, 0x64, 0x20, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63,
	0x20, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xd2, 0x01, 0x0f, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x22, 0x96,
	0x01, 0x0a, 0x23, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x3a, 0x6f, 0x92, 0x41, 0x6c, 0x0a, 0x6a, 0x2a, 0x26, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x20, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x20, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x40, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x20, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x20, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x20, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0xf4, 0x01, 0x0a, 0x26, 0x49, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0xc9, 0x01, 0x0a, 0x1b, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x49, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x41, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x42, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x49, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0xc2, 0xb8, 0x02, 0x1f, 0x0a,
	0x1d, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x20, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_integrations_v1_integration_extensions_deployment_service_proto_rawDescOnce sync.Once
	file_com_coralogix_integrations_v1_integration_extensions_deployment_service_proto_rawDescData = file_com_coralogix_integrations_v1_integration_extensions_deployment_service_proto_rawDesc
)

func file_com_coralogix_integrations_v1_integration_extensions_deployment_service_proto_rawDescGZIP() []byte {
	file_com_coralogix_integrations_v1_integration_extensions_deployment_service_proto_rawDescOnce.Do(func() {
		file_com_coralogix_integrations_v1_integration_extensions_deployment_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_integrations_v1_integration_extensions_deployment_service_proto_rawDescData)
	})
	return file_com_coralogix_integrations_v1_integration_extensions_deployment_service_proto_rawDescData
}

var file_com_coralogix_integrations_v1_integration_extensions_deployment_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogix_integrations_v1_integration_extensions_deployment_service_proto_goTypes = []any{
	(*DeployIntegrationExtensionsRequest)(nil),  // 0: com.coralogix.integrations.v1.DeployIntegrationExtensionsRequest
	(*DeployIntegrationExtensionsResponse)(nil), // 1: com.coralogix.integrations.v1.DeployIntegrationExtensionsResponse
	(*wrapperspb.StringValue)(nil),              // 2: google.protobuf.StringValue
}
var file_com_coralogix_integrations_v1_integration_extensions_deployment_service_proto_depIdxs = []int32{
	2, // 0: com.coralogix.integrations.v1.DeployIntegrationExtensionsRequest.integration_key:type_name -> google.protobuf.StringValue
	0, // 1: com.coralogix.integrations.v1.IntegrationExtensionsDeploymentService.DeployIntegrationExtensions:input_type -> com.coralogix.integrations.v1.DeployIntegrationExtensionsRequest
	1, // 2: com.coralogix.integrations.v1.IntegrationExtensionsDeploymentService.DeployIntegrationExtensions:output_type -> com.coralogix.integrations.v1.DeployIntegrationExtensionsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_com_coralogix_integrations_v1_integration_extensions_deployment_service_proto_init()
}
func file_com_coralogix_integrations_v1_integration_extensions_deployment_service_proto_init() {
	if File_com_coralogix_integrations_v1_integration_extensions_deployment_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_integrations_v1_integration_extensions_deployment_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogix_integrations_v1_integration_extensions_deployment_service_proto_goTypes,
		DependencyIndexes: file_com_coralogix_integrations_v1_integration_extensions_deployment_service_proto_depIdxs,
		MessageInfos:      file_com_coralogix_integrations_v1_integration_extensions_deployment_service_proto_msgTypes,
	}.Build()
	File_com_coralogix_integrations_v1_integration_extensions_deployment_service_proto = out.File
	file_com_coralogix_integrations_v1_integration_extensions_deployment_service_proto_rawDesc = nil
	file_com_coralogix_integrations_v1_integration_extensions_deployment_service_proto_goTypes = nil
	file_com_coralogix_integrations_v1_integration_extensions_deployment_service_proto_depIdxs = nil
}
