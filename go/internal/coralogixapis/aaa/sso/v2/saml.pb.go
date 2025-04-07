// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.26.1
// source: com/coralogixapis/aaa/sso/v2/saml.proto

package v2

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type GetSPParametersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TeamId        uint32                 `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSPParametersRequest) Reset() {
	*x = GetSPParametersRequest{}
	mi := &file_com_coralogixapis_aaa_sso_v2_saml_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSPParametersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSPParametersRequest) ProtoMessage() {}

func (x *GetSPParametersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_aaa_sso_v2_saml_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSPParametersRequest.ProtoReflect.Descriptor instead.
func (*GetSPParametersRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_aaa_sso_v2_saml_proto_rawDescGZIP(), []int{0}
}

func (x *GetSPParametersRequest) GetTeamId() uint32 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

type SPParameters struct {
	state                       protoimpl.MessageState `protogen:"open.v1"`
	MetadataUrl                 string                 `protobuf:"bytes,1,opt,name=metadata_url,json=metadataUrl,proto3" json:"metadata_url,omitempty"`
	SigningCertPem              string                 `protobuf:"bytes,2,opt,name=signing_cert_pem,json=signingCertPem,proto3" json:"signing_cert_pem,omitempty"`
	NameIdFormat                string                 `protobuf:"bytes,3,opt,name=name_id_format,json=nameIdFormat,proto3" json:"name_id_format,omitempty"`
	AssertionConsumerServiceUrl string                 `protobuf:"bytes,4,opt,name=assertion_consumer_service_url,json=assertionConsumerServiceUrl,proto3" json:"assertion_consumer_service_url,omitempty"`
	Binding                     string                 `protobuf:"bytes,5,opt,name=binding,proto3" json:"binding,omitempty"`
	unknownFields               protoimpl.UnknownFields
	sizeCache                   protoimpl.SizeCache
}

func (x *SPParameters) Reset() {
	*x = SPParameters{}
	mi := &file_com_coralogixapis_aaa_sso_v2_saml_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SPParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SPParameters) ProtoMessage() {}

func (x *SPParameters) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_aaa_sso_v2_saml_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SPParameters.ProtoReflect.Descriptor instead.
func (*SPParameters) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_aaa_sso_v2_saml_proto_rawDescGZIP(), []int{1}
}

func (x *SPParameters) GetMetadataUrl() string {
	if x != nil {
		return x.MetadataUrl
	}
	return ""
}

func (x *SPParameters) GetSigningCertPem() string {
	if x != nil {
		return x.SigningCertPem
	}
	return ""
}

func (x *SPParameters) GetNameIdFormat() string {
	if x != nil {
		return x.NameIdFormat
	}
	return ""
}

func (x *SPParameters) GetAssertionConsumerServiceUrl() string {
	if x != nil {
		return x.AssertionConsumerServiceUrl
	}
	return ""
}

func (x *SPParameters) GetBinding() string {
	if x != nil {
		return x.Binding
	}
	return ""
}

type GetSPParametersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Params        *SPParameters          `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSPParametersResponse) Reset() {
	*x = GetSPParametersResponse{}
	mi := &file_com_coralogixapis_aaa_sso_v2_saml_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSPParametersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSPParametersResponse) ProtoMessage() {}

func (x *GetSPParametersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_aaa_sso_v2_saml_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSPParametersResponse.ProtoReflect.Descriptor instead.
func (*GetSPParametersResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_aaa_sso_v2_saml_proto_rawDescGZIP(), []int{2}
}

func (x *GetSPParametersResponse) GetParams() *SPParameters {
	if x != nil {
		return x.Params
	}
	return nil
}

type IDPParameters struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	Active bool                   `protobuf:"varint,2,opt,name=active,proto3" json:"active,omitempty"`
	// Types that are valid to be assigned to Metadata:
	//
	//	*IDPParameters_MetadataUrl
	//	*IDPParameters_MetadataContent
	Metadata      isIDPParameters_Metadata `protobuf_oneof:"metadata"`
	TeamEntityId  *uint32                  `protobuf:"varint,5,opt,name=team_entity_id,json=teamEntityId,proto3,oneof" json:"team_entity_id,omitempty"`
	GroupNames    []string                 `protobuf:"bytes,6,rep,name=group_names,json=groupNames,proto3" json:"group_names,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IDPParameters) Reset() {
	*x = IDPParameters{}
	mi := &file_com_coralogixapis_aaa_sso_v2_saml_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IDPParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDPParameters) ProtoMessage() {}

func (x *IDPParameters) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_aaa_sso_v2_saml_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDPParameters.ProtoReflect.Descriptor instead.
func (*IDPParameters) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_aaa_sso_v2_saml_proto_rawDescGZIP(), []int{3}
}

func (x *IDPParameters) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *IDPParameters) GetMetadata() isIDPParameters_Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *IDPParameters) GetMetadataUrl() string {
	if x != nil {
		if x, ok := x.Metadata.(*IDPParameters_MetadataUrl); ok {
			return x.MetadataUrl
		}
	}
	return ""
}

func (x *IDPParameters) GetMetadataContent() string {
	if x != nil {
		if x, ok := x.Metadata.(*IDPParameters_MetadataContent); ok {
			return x.MetadataContent
		}
	}
	return ""
}

func (x *IDPParameters) GetTeamEntityId() uint32 {
	if x != nil && x.TeamEntityId != nil {
		return *x.TeamEntityId
	}
	return 0
}

func (x *IDPParameters) GetGroupNames() []string {
	if x != nil {
		return x.GroupNames
	}
	return nil
}

type isIDPParameters_Metadata interface {
	isIDPParameters_Metadata()
}

type IDPParameters_MetadataUrl struct {
	MetadataUrl string `protobuf:"bytes,3,opt,name=metadata_url,json=metadataUrl,proto3,oneof"`
}

type IDPParameters_MetadataContent struct {
	MetadataContent string `protobuf:"bytes,4,opt,name=metadata_content,json=metadataContent,proto3,oneof"`
}

func (*IDPParameters_MetadataUrl) isIDPParameters_Metadata() {}

func (*IDPParameters_MetadataContent) isIDPParameters_Metadata() {}

type SetIDPParametersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TeamId        uint32                 `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	Params        *IDPParameters         `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetIDPParametersRequest) Reset() {
	*x = SetIDPParametersRequest{}
	mi := &file_com_coralogixapis_aaa_sso_v2_saml_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetIDPParametersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetIDPParametersRequest) ProtoMessage() {}

func (x *SetIDPParametersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_aaa_sso_v2_saml_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetIDPParametersRequest.ProtoReflect.Descriptor instead.
func (*SetIDPParametersRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_aaa_sso_v2_saml_proto_rawDescGZIP(), []int{4}
}

func (x *SetIDPParametersRequest) GetTeamId() uint32 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *SetIDPParametersRequest) GetParams() *IDPParameters {
	if x != nil {
		return x.Params
	}
	return nil
}

type SetIDPParametersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetIDPParametersResponse) Reset() {
	*x = SetIDPParametersResponse{}
	mi := &file_com_coralogixapis_aaa_sso_v2_saml_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetIDPParametersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetIDPParametersResponse) ProtoMessage() {}

func (x *SetIDPParametersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_aaa_sso_v2_saml_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetIDPParametersResponse.ProtoReflect.Descriptor instead.
func (*SetIDPParametersResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_aaa_sso_v2_saml_proto_rawDescGZIP(), []int{5}
}

type SetActiveRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TeamId        uint32                 `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	IsActive      bool                   `protobuf:"varint,2,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetActiveRequest) Reset() {
	*x = SetActiveRequest{}
	mi := &file_com_coralogixapis_aaa_sso_v2_saml_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetActiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetActiveRequest) ProtoMessage() {}

func (x *SetActiveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_aaa_sso_v2_saml_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetActiveRequest.ProtoReflect.Descriptor instead.
func (*SetActiveRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_aaa_sso_v2_saml_proto_rawDescGZIP(), []int{6}
}

func (x *SetActiveRequest) GetTeamId() uint32 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *SetActiveRequest) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

type SetActiveResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetActiveResponse) Reset() {
	*x = SetActiveResponse{}
	mi := &file_com_coralogixapis_aaa_sso_v2_saml_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetActiveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetActiveResponse) ProtoMessage() {}

func (x *SetActiveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_aaa_sso_v2_saml_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetActiveResponse.ProtoReflect.Descriptor instead.
func (*SetActiveResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_aaa_sso_v2_saml_proto_rawDescGZIP(), []int{7}
}

type GetConfigurationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TeamId        uint32                 `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetConfigurationRequest) Reset() {
	*x = GetConfigurationRequest{}
	mi := &file_com_coralogixapis_aaa_sso_v2_saml_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetConfigurationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigurationRequest) ProtoMessage() {}

func (x *GetConfigurationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_aaa_sso_v2_saml_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigurationRequest.ProtoReflect.Descriptor instead.
func (*GetConfigurationRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_aaa_sso_v2_saml_proto_rawDescGZIP(), []int{8}
}

func (x *GetConfigurationRequest) GetTeamId() uint32 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

type GetConfigurationResponse struct {
	state         protoimpl.MessageState               `protogen:"open.v1"`
	TeamId        uint32                               `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	SpParameters  *SPParameters                        `protobuf:"bytes,2,opt,name=sp_parameters,json=spParameters,proto3" json:"sp_parameters,omitempty"`
	IdpParameters *IDPParameters                       `protobuf:"bytes,3,opt,name=idp_parameters,json=idpParameters,proto3" json:"idp_parameters,omitempty"`
	IdpDetails    *GetConfigurationResponse_IDPDetails `protobuf:"bytes,4,opt,name=idp_details,json=idpDetails,proto3,oneof" json:"idp_details,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetConfigurationResponse) Reset() {
	*x = GetConfigurationResponse{}
	mi := &file_com_coralogixapis_aaa_sso_v2_saml_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetConfigurationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigurationResponse) ProtoMessage() {}

func (x *GetConfigurationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_aaa_sso_v2_saml_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigurationResponse.ProtoReflect.Descriptor instead.
func (*GetConfigurationResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_aaa_sso_v2_saml_proto_rawDescGZIP(), []int{9}
}

func (x *GetConfigurationResponse) GetTeamId() uint32 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *GetConfigurationResponse) GetSpParameters() *SPParameters {
	if x != nil {
		return x.SpParameters
	}
	return nil
}

func (x *GetConfigurationResponse) GetIdpParameters() *IDPParameters {
	if x != nil {
		return x.IdpParameters
	}
	return nil
}

func (x *GetConfigurationResponse) GetIdpDetails() *GetConfigurationResponse_IDPDetails {
	if x != nil {
		return x.IdpDetails
	}
	return nil
}

type GetConfigurationResponse_IDPDetails struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Icon          string                 `protobuf:"bytes,2,opt,name=icon,proto3" json:"icon,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetConfigurationResponse_IDPDetails) Reset() {
	*x = GetConfigurationResponse_IDPDetails{}
	mi := &file_com_coralogixapis_aaa_sso_v2_saml_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetConfigurationResponse_IDPDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigurationResponse_IDPDetails) ProtoMessage() {}

func (x *GetConfigurationResponse_IDPDetails) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_aaa_sso_v2_saml_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigurationResponse_IDPDetails.ProtoReflect.Descriptor instead.
func (*GetConfigurationResponse_IDPDetails) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_aaa_sso_v2_saml_proto_rawDescGZIP(), []int{9, 0}
}

func (x *GetConfigurationResponse_IDPDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetConfigurationResponse_IDPDetails) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

var File_com_coralogixapis_aaa_sso_v2_saml_proto protoreflect.FileDescriptor

const file_com_coralogixapis_aaa_sso_v2_saml_proto_rawDesc = "" +
	"\n" +
	"'com/coralogixapis/aaa/sso/v2/saml.proto\x12\x1ccom.coralogixapis.aaa.sso.v2\x1a.protoc-gen-openapiv2/options/annotations.proto\x1a\x1cgoogle/api/annotations.proto\"\xbe\x02\n" +
	"\x16GetSPParametersRequest\x12%\n" +
	"\ateam_id\x18\x01 \x01(\rB\f\x92A\tJ\a1234567R\x06teamId:\xfc\x01\x92A\xf8\x01\n" +
	"x*\x19Get SP Parameters Request2QThis data structure is used to retrieve the parameters of a SAML service provider\xd2\x01\ateam_id*|\n" +
	"\x1fFind out more about enrichments\x12Yhttps://coralogix.com/docs/user-guides/data-transformation/enrichments/custom-enrichment/\"\xae\x05\n" +
	"\fSPParameters\x12\\\n" +
	"\fmetadata_url\x18\x01 \x01(\tB9\x92A6J4\"https://<...>.okta.com/app/<...>/sso/saml/metadata\"R\vmetadataUrl\x12<\n" +
	"\x10signing_cert_pem\x18\x02 \x01(\tB\x12\x92A\x0fJ\r\"certificate\"R\x0esigningCertPem\x124\n" +
	"\x0ename_id_format\x18\x03 \x01(\tB\x0e\x92A\vJ\t\"name_id\"R\fnameIdFormat\x12U\n" +
	"\x1eassertion_consumer_service_url\x18\x04 \x01(\tB\x10\x92A\rJ\v\"assertion\"R\x1bassertionConsumerServiceUrl\x12(\n" +
	"\abinding\x18\x05 \x01(\tB\x0e\x92A\vJ\t\"binding\"R\abinding:\xca\x02\x92A\xc6\x02\n" +
	"\xc5\x01*\x1bService Provider Parameters2HThis data structure represents the parameters of a SAML service provider\xd2\x01\fmetadata_url\xd2\x01\x10signing_cert_pem\xd2\x01\x0ename_id_format\xd2\x01\x1eassertion_consumer_service_url\xd2\x01\abinding*|\n" +
	"\x1fFind out more about enrichments\x12Yhttps://coralogix.com/docs/user-guides/data-transformation/enrichments/custom-enrichment/\"\xfc\x02\n" +
	"\x17GetSPParametersResponse\x12B\n" +
	"\x06params\x18\x01 \x01(\v2*.com.coralogixapis.aaa.sso.v2.SPParametersR\x06params:\x9c\x02\x92A\x98\x02\n" +
	"\x97\x01*\x1aGet SP Parameters Response2pThis data structure is obtained as a response to a request to retrieve the parameters of a SAML service provider\xd2\x01\x06params*|\n" +
	"\x1fFind out more about enrichments\x12Yhttps://coralogix.com/docs/user-guides/data-transformation/enrichments/custom-enrichment/\"\xfa\x04\n" +
	"\rIDPParameters\x12!\n" +
	"\x06active\x18\x02 \x01(\bB\t\x92A\x06J\x04trueR\x06active\x12^\n" +
	"\fmetadata_url\x18\x03 \x01(\tB9\x92A6J4\"https://<...>.okta.com/app/<...>/sso/saml/metadata\"H\x00R\vmetadataUrl\x12F\n" +
	"\x10metadata_content\x18\x04 \x01(\tB\x19\x92A\x16J\x14\"<?xml version= ...\"H\x00R\x0fmetadataContent\x127\n" +
	"\x0eteam_entity_id\x18\x05 \x01(\rB\f\x92A\tJ\a1234567H\x01R\fteamEntityId\x88\x01\x01\x120\n" +
	"\vgroup_names\x18\x06 \x03(\tB\x0f\x92A\fJ\n" +
	"[\"group1\"]R\n" +
	"groupNames:\x93\x02\x92A\x8f\x02\n" +
	"\x8e\x01*\x0eIDP Parameters2IThis data structure represents a set of SAML identity provider parameters\xd2\x01\x06active\xd2\x01\bmetadata\xd2\x01\x0eteam_entity_id\xd2\x01\vgroup_names*|\n" +
	"\x1fFind out more about enrichments\x12Yhttps://coralogix.com/docs/user-guides/data-transformation/enrichments/custom-enrichment/B\n" +
	"\n" +
	"\bmetadataB\x11\n" +
	"\x0f_team_entity_id\"\xfc\x02\n" +
	"\x17SetIDPParametersRequest\x12\x17\n" +
	"\ateam_id\x18\x01 \x01(\rR\x06teamId\x12C\n" +
	"\x06params\x18\x02 \x01(\v2+.com.coralogixapis.aaa.sso.v2.IDPParametersR\x06params:\x82\x02\x92A\xfe\x01\n" +
	"~*\x1aSet IDP Parameters Request2MThis data structure is used to set the parameters of a SAML identity provider\xd2\x01\ateam_id\xd2\x01\x06params*|\n" +
	"\x1fFind out more about enrichments\x12Yhttps://coralogix.com/docs/user-guides/data-transformation/enrichments/custom-enrichment/\"\x1a\n" +
	"\x18SetIDPParametersResponse\"\xc9\x02\n" +
	"\x10SetActiveRequest\x12\x17\n" +
	"\ateam_id\x18\x01 \x01(\rR\x06teamId\x12\x1b\n" +
	"\tis_active\x18\x02 \x01(\bR\bisActive:\xfe\x01\x92A\xfa\x01\n" +
	"z*\x12Set Active Request2NThis data structure is used to activate or deactivate a SAML identity provider\xd2\x01\ateam_id\xd2\x01\tis_active*|\n" +
	"\x1fFind out more about enrichments\x12Yhttps://coralogix.com/docs/user-guides/data-transformation/enrichments/custom-enrichment/\"\x13\n" +
	"\x11SetActiveResponse\"\xcb\x02\n" +
	"\x17GetConfigurationRequest\x12\x17\n" +
	"\ateam_id\x18\x01 \x01(\rR\x06teamId:\x96\x02\x92A\x92\x02\n" +
	"\x91\x01*\x19Get Configuration Request2jThis data structure is used to retrieve the configuration of a SAML service provider and identity provider\xd2\x01\ateam_id*|\n" +
	"\x1fFind out more about enrichments\x12Yhttps://coralogix.com/docs/user-guides/data-transformation/enrichments/custom-enrichment/\"\xe2\x05\n" +
	"\x18GetConfigurationResponse\x12\x17\n" +
	"\ateam_id\x18\x01 \x01(\rR\x06teamId\x12O\n" +
	"\rsp_parameters\x18\x02 \x01(\v2*.com.coralogixapis.aaa.sso.v2.SPParametersR\fspParameters\x12R\n" +
	"\x0eidp_parameters\x18\x03 \x01(\v2+.com.coralogixapis.aaa.sso.v2.IDPParametersR\ridpParameters\x12g\n" +
	"\vidp_details\x18\x04 \x01(\v2A.com.coralogixapis.aaa.sso.v2.GetConfigurationResponse.IDPDetailsH\x00R\n" +
	"idpDetails\x88\x01\x01\x1a4\n" +
	"\n" +
	"IDPDetails\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x12\n" +
	"\x04icon\x18\x02 \x01(\tR\x04icon:\xd8\x02\x92A\xd4\x02\n" +
	"\xd3\x01*\x1aGet Configuration Response2\x89\x01This data structure is obtained as a response to a request to retrieve the configuration of a SAML service provider and identity provider\xd2\x01\ateam_id\xd2\x01\rsp_parameters\xd2\x01\x0eidp_parameters*|\n" +
	"\x1fFind out more about enrichments\x12Yhttps://coralogix.com/docs/user-guides/data-transformation/enrichments/custom-enrichment/B\x0e\n" +
	"\f_idp_details2\xa2\x05\n" +
	"\x18SamlConfigurationService\x12\xa2\x01\n" +
	"\x0fGetSPParameters\x124.com.coralogixapis.aaa.sso.v2.GetSPParametersRequest\x1a5.com.coralogixapis.aaa.sso.v2.GetSPParametersResponse\"\"\x82\xd3\xe4\x93\x02\x1c\x12\x1a/v1/sso/saml/sp_parameters\x12\xa9\x01\n" +
	"\x10SetIDPParameters\x125.com.coralogixapis.aaa.sso.v2.SetIDPParametersRequest\x1a6.com.coralogixapis.aaa.sso.v2.SetIDPParametersResponse\"&\x82\xd3\xe4\x93\x02 :\x01*\"\x1b/v1/sso/saml/idp_parameters\x12\x8c\x01\n" +
	"\tSetActive\x12..com.coralogixapis.aaa.sso.v2.SetActiveRequest\x1a/.com.coralogixapis.aaa.sso.v2.SetActiveResponse\"\x1e\x82\xd3\xe4\x93\x02\x18:\x01*\"\x13/v1/sso/saml/active\x12\xa5\x01\n" +
	"\x10GetConfiguration\x125.com.coralogixapis.aaa.sso.v2.GetConfigurationRequest\x1a6.com.coralogixapis.aaa.sso.v2.GetConfigurationResponse\"\"\x82\xd3\xe4\x93\x02\x1c\x12\x1a/v1/sso/saml/configurationb\x06proto3"

var (
	file_com_coralogixapis_aaa_sso_v2_saml_proto_rawDescOnce sync.Once
	file_com_coralogixapis_aaa_sso_v2_saml_proto_rawDescData []byte
)

func file_com_coralogixapis_aaa_sso_v2_saml_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_aaa_sso_v2_saml_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_aaa_sso_v2_saml_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_aaa_sso_v2_saml_proto_rawDesc), len(file_com_coralogixapis_aaa_sso_v2_saml_proto_rawDesc)))
	})
	return file_com_coralogixapis_aaa_sso_v2_saml_proto_rawDescData
}

var file_com_coralogixapis_aaa_sso_v2_saml_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_com_coralogixapis_aaa_sso_v2_saml_proto_goTypes = []any{
	(*GetSPParametersRequest)(nil),              // 0: com.coralogixapis.aaa.sso.v2.GetSPParametersRequest
	(*SPParameters)(nil),                        // 1: com.coralogixapis.aaa.sso.v2.SPParameters
	(*GetSPParametersResponse)(nil),             // 2: com.coralogixapis.aaa.sso.v2.GetSPParametersResponse
	(*IDPParameters)(nil),                       // 3: com.coralogixapis.aaa.sso.v2.IDPParameters
	(*SetIDPParametersRequest)(nil),             // 4: com.coralogixapis.aaa.sso.v2.SetIDPParametersRequest
	(*SetIDPParametersResponse)(nil),            // 5: com.coralogixapis.aaa.sso.v2.SetIDPParametersResponse
	(*SetActiveRequest)(nil),                    // 6: com.coralogixapis.aaa.sso.v2.SetActiveRequest
	(*SetActiveResponse)(nil),                   // 7: com.coralogixapis.aaa.sso.v2.SetActiveResponse
	(*GetConfigurationRequest)(nil),             // 8: com.coralogixapis.aaa.sso.v2.GetConfigurationRequest
	(*GetConfigurationResponse)(nil),            // 9: com.coralogixapis.aaa.sso.v2.GetConfigurationResponse
	(*GetConfigurationResponse_IDPDetails)(nil), // 10: com.coralogixapis.aaa.sso.v2.GetConfigurationResponse.IDPDetails
}
var file_com_coralogixapis_aaa_sso_v2_saml_proto_depIdxs = []int32{
	1,  // 0: com.coralogixapis.aaa.sso.v2.GetSPParametersResponse.params:type_name -> com.coralogixapis.aaa.sso.v2.SPParameters
	3,  // 1: com.coralogixapis.aaa.sso.v2.SetIDPParametersRequest.params:type_name -> com.coralogixapis.aaa.sso.v2.IDPParameters
	1,  // 2: com.coralogixapis.aaa.sso.v2.GetConfigurationResponse.sp_parameters:type_name -> com.coralogixapis.aaa.sso.v2.SPParameters
	3,  // 3: com.coralogixapis.aaa.sso.v2.GetConfigurationResponse.idp_parameters:type_name -> com.coralogixapis.aaa.sso.v2.IDPParameters
	10, // 4: com.coralogixapis.aaa.sso.v2.GetConfigurationResponse.idp_details:type_name -> com.coralogixapis.aaa.sso.v2.GetConfigurationResponse.IDPDetails
	0,  // 5: com.coralogixapis.aaa.sso.v2.SamlConfigurationService.GetSPParameters:input_type -> com.coralogixapis.aaa.sso.v2.GetSPParametersRequest
	4,  // 6: com.coralogixapis.aaa.sso.v2.SamlConfigurationService.SetIDPParameters:input_type -> com.coralogixapis.aaa.sso.v2.SetIDPParametersRequest
	6,  // 7: com.coralogixapis.aaa.sso.v2.SamlConfigurationService.SetActive:input_type -> com.coralogixapis.aaa.sso.v2.SetActiveRequest
	8,  // 8: com.coralogixapis.aaa.sso.v2.SamlConfigurationService.GetConfiguration:input_type -> com.coralogixapis.aaa.sso.v2.GetConfigurationRequest
	2,  // 9: com.coralogixapis.aaa.sso.v2.SamlConfigurationService.GetSPParameters:output_type -> com.coralogixapis.aaa.sso.v2.GetSPParametersResponse
	5,  // 10: com.coralogixapis.aaa.sso.v2.SamlConfigurationService.SetIDPParameters:output_type -> com.coralogixapis.aaa.sso.v2.SetIDPParametersResponse
	7,  // 11: com.coralogixapis.aaa.sso.v2.SamlConfigurationService.SetActive:output_type -> com.coralogixapis.aaa.sso.v2.SetActiveResponse
	9,  // 12: com.coralogixapis.aaa.sso.v2.SamlConfigurationService.GetConfiguration:output_type -> com.coralogixapis.aaa.sso.v2.GetConfigurationResponse
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_aaa_sso_v2_saml_proto_init() }
func file_com_coralogixapis_aaa_sso_v2_saml_proto_init() {
	if File_com_coralogixapis_aaa_sso_v2_saml_proto != nil {
		return
	}
	file_com_coralogixapis_aaa_sso_v2_saml_proto_msgTypes[3].OneofWrappers = []any{
		(*IDPParameters_MetadataUrl)(nil),
		(*IDPParameters_MetadataContent)(nil),
	}
	file_com_coralogixapis_aaa_sso_v2_saml_proto_msgTypes[9].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_aaa_sso_v2_saml_proto_rawDesc), len(file_com_coralogixapis_aaa_sso_v2_saml_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogixapis_aaa_sso_v2_saml_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_aaa_sso_v2_saml_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_aaa_sso_v2_saml_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_aaa_sso_v2_saml_proto = out.File
	file_com_coralogixapis_aaa_sso_v2_saml_proto_goTypes = nil
	file_com_coralogixapis_aaa_sso_v2_saml_proto_depIdxs = nil
}
