// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogix/enrichment/v1/enrichment_internal_service.proto

package v1

import (
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

type GetInternalEnrichmentsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetInternalEnrichmentsRequest) Reset() {
	*x = GetInternalEnrichmentsRequest{}
	mi := &file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInternalEnrichmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInternalEnrichmentsRequest) ProtoMessage() {}

func (x *GetInternalEnrichmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInternalEnrichmentsRequest.ProtoReflect.Descriptor instead.
func (*GetInternalEnrichmentsRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDescGZIP(), []int{0}
}

type GetInternalEnrichmentsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Enrichments   []*Enrichment          `protobuf:"bytes,1,rep,name=enrichments,proto3" json:"enrichments,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetInternalEnrichmentsResponse) Reset() {
	*x = GetInternalEnrichmentsResponse{}
	mi := &file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInternalEnrichmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInternalEnrichmentsResponse) ProtoMessage() {}

func (x *GetInternalEnrichmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInternalEnrichmentsResponse.ProtoReflect.Descriptor instead.
func (*GetInternalEnrichmentsResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetInternalEnrichmentsResponse) GetEnrichments() []*Enrichment {
	if x != nil {
		return x.Enrichments
	}
	return nil
}

type AddInternalEnrichmentsRequest struct {
	state              protoimpl.MessageState    `protogen:"open.v1"`
	RequestEnrichments []*EnrichmentRequestModel `protobuf:"bytes,1,rep,name=request_enrichments,json=requestEnrichments,proto3" json:"request_enrichments,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *AddInternalEnrichmentsRequest) Reset() {
	*x = AddInternalEnrichmentsRequest{}
	mi := &file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddInternalEnrichmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddInternalEnrichmentsRequest) ProtoMessage() {}

func (x *AddInternalEnrichmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddInternalEnrichmentsRequest.ProtoReflect.Descriptor instead.
func (*AddInternalEnrichmentsRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDescGZIP(), []int{2}
}

func (x *AddInternalEnrichmentsRequest) GetRequestEnrichments() []*EnrichmentRequestModel {
	if x != nil {
		return x.RequestEnrichments
	}
	return nil
}

type AddInternalEnrichmentsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Enrichments   []*Enrichment          `protobuf:"bytes,1,rep,name=enrichments,proto3" json:"enrichments,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddInternalEnrichmentsResponse) Reset() {
	*x = AddInternalEnrichmentsResponse{}
	mi := &file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddInternalEnrichmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddInternalEnrichmentsResponse) ProtoMessage() {}

func (x *AddInternalEnrichmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddInternalEnrichmentsResponse.ProtoReflect.Descriptor instead.
func (*AddInternalEnrichmentsResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDescGZIP(), []int{3}
}

func (x *AddInternalEnrichmentsResponse) GetEnrichments() []*Enrichment {
	if x != nil {
		return x.Enrichments
	}
	return nil
}

type RemoveInternalEnrichmentsRequest struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	EnrichmentIds []*wrapperspb.UInt32Value `protobuf:"bytes,1,rep,name=enrichment_ids,json=enrichmentIds,proto3" json:"enrichment_ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveInternalEnrichmentsRequest) Reset() {
	*x = RemoveInternalEnrichmentsRequest{}
	mi := &file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveInternalEnrichmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveInternalEnrichmentsRequest) ProtoMessage() {}

func (x *RemoveInternalEnrichmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveInternalEnrichmentsRequest.ProtoReflect.Descriptor instead.
func (*RemoveInternalEnrichmentsRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveInternalEnrichmentsRequest) GetEnrichmentIds() []*wrapperspb.UInt32Value {
	if x != nil {
		return x.EnrichmentIds
	}
	return nil
}

type RemoveInternalEnrichmentsResponse struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	RemainingEnrichments []*Enrichment          `protobuf:"bytes,1,rep,name=remaining_enrichments,json=remainingEnrichments,proto3" json:"remaining_enrichments,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *RemoveInternalEnrichmentsResponse) Reset() {
	*x = RemoveInternalEnrichmentsResponse{}
	mi := &file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveInternalEnrichmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveInternalEnrichmentsResponse) ProtoMessage() {}

func (x *RemoveInternalEnrichmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveInternalEnrichmentsResponse.ProtoReflect.Descriptor instead.
func (*RemoveInternalEnrichmentsResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDescGZIP(), []int{5}
}

func (x *RemoveInternalEnrichmentsResponse) GetRemainingEnrichments() []*Enrichment {
	if x != nil {
		return x.RemainingEnrichments
	}
	return nil
}

var File_com_coralogix_enrichment_v1_enrichment_internal_service_proto protoreflect.FileDescriptor

const file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDesc = "" +
	"\n" +
	"=com/coralogix/enrichment/v1/enrichment_internal_service.proto\x12\x1bcom.coralogix.enrichment.v1\x1a\x1egoogle/protobuf/wrappers.proto\x1a,com/coralogix/enrichment/v1/enrichment.proto\x1a:com/coralogix/enrichment/v1/enrichment_request_model.proto\"\x1f\n" +
	"\x1dGetInternalEnrichmentsRequest\"k\n" +
	"\x1eGetInternalEnrichmentsResponse\x12I\n" +
	"\venrichments\x18\x01 \x03(\v2'.com.coralogix.enrichment.v1.EnrichmentR\venrichments\"\x85\x01\n" +
	"\x1dAddInternalEnrichmentsRequest\x12d\n" +
	"\x13request_enrichments\x18\x01 \x03(\v23.com.coralogix.enrichment.v1.EnrichmentRequestModelR\x12requestEnrichments\"k\n" +
	"\x1eAddInternalEnrichmentsResponse\x12I\n" +
	"\venrichments\x18\x01 \x03(\v2'.com.coralogix.enrichment.v1.EnrichmentR\venrichments\"g\n" +
	" RemoveInternalEnrichmentsRequest\x12C\n" +
	"\x0eenrichment_ids\x18\x01 \x03(\v2\x1c.google.protobuf.UInt32ValueR\renrichmentIds\"\x81\x01\n" +
	"!RemoveInternalEnrichmentsResponse\x12\\\n" +
	"\x15remaining_enrichments\x18\x01 \x03(\v2'.com.coralogix.enrichment.v1.EnrichmentR\x14remainingEnrichments2\xe0\x03\n" +
	"\x19EnrichmentInternalService\x12\x91\x01\n" +
	"\x16GetInternalEnrichments\x12:.com.coralogix.enrichment.v1.GetInternalEnrichmentsRequest\x1a;.com.coralogix.enrichment.v1.GetInternalEnrichmentsResponse\x12\x91\x01\n" +
	"\x16AddInternalEnrichments\x12:.com.coralogix.enrichment.v1.AddInternalEnrichmentsRequest\x1a;.com.coralogix.enrichment.v1.AddInternalEnrichmentsResponse\x12\x9a\x01\n" +
	"\x19RemoveInternalEnrichments\x12=.com.coralogix.enrichment.v1.RemoveInternalEnrichmentsRequest\x1a>.com.coralogix.enrichment.v1.RemoveInternalEnrichmentsResponseb\x06proto3"

var (
	file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDescOnce sync.Once
	file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDescData []byte
)

func file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDescGZIP() []byte {
	file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDescOnce.Do(func() {
		file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDesc), len(file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDesc)))
	})
	return file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDescData
}

var file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_goTypes = []any{
	(*GetInternalEnrichmentsRequest)(nil),     // 0: com.coralogix.enrichment.v1.GetInternalEnrichmentsRequest
	(*GetInternalEnrichmentsResponse)(nil),    // 1: com.coralogix.enrichment.v1.GetInternalEnrichmentsResponse
	(*AddInternalEnrichmentsRequest)(nil),     // 2: com.coralogix.enrichment.v1.AddInternalEnrichmentsRequest
	(*AddInternalEnrichmentsResponse)(nil),    // 3: com.coralogix.enrichment.v1.AddInternalEnrichmentsResponse
	(*RemoveInternalEnrichmentsRequest)(nil),  // 4: com.coralogix.enrichment.v1.RemoveInternalEnrichmentsRequest
	(*RemoveInternalEnrichmentsResponse)(nil), // 5: com.coralogix.enrichment.v1.RemoveInternalEnrichmentsResponse
	(*Enrichment)(nil),                        // 6: com.coralogix.enrichment.v1.Enrichment
	(*EnrichmentRequestModel)(nil),            // 7: com.coralogix.enrichment.v1.EnrichmentRequestModel
	(*wrapperspb.UInt32Value)(nil),            // 8: google.protobuf.UInt32Value
}
var file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_depIdxs = []int32{
	6, // 0: com.coralogix.enrichment.v1.GetInternalEnrichmentsResponse.enrichments:type_name -> com.coralogix.enrichment.v1.Enrichment
	7, // 1: com.coralogix.enrichment.v1.AddInternalEnrichmentsRequest.request_enrichments:type_name -> com.coralogix.enrichment.v1.EnrichmentRequestModel
	6, // 2: com.coralogix.enrichment.v1.AddInternalEnrichmentsResponse.enrichments:type_name -> com.coralogix.enrichment.v1.Enrichment
	8, // 3: com.coralogix.enrichment.v1.RemoveInternalEnrichmentsRequest.enrichment_ids:type_name -> google.protobuf.UInt32Value
	6, // 4: com.coralogix.enrichment.v1.RemoveInternalEnrichmentsResponse.remaining_enrichments:type_name -> com.coralogix.enrichment.v1.Enrichment
	0, // 5: com.coralogix.enrichment.v1.EnrichmentInternalService.GetInternalEnrichments:input_type -> com.coralogix.enrichment.v1.GetInternalEnrichmentsRequest
	2, // 6: com.coralogix.enrichment.v1.EnrichmentInternalService.AddInternalEnrichments:input_type -> com.coralogix.enrichment.v1.AddInternalEnrichmentsRequest
	4, // 7: com.coralogix.enrichment.v1.EnrichmentInternalService.RemoveInternalEnrichments:input_type -> com.coralogix.enrichment.v1.RemoveInternalEnrichmentsRequest
	1, // 8: com.coralogix.enrichment.v1.EnrichmentInternalService.GetInternalEnrichments:output_type -> com.coralogix.enrichment.v1.GetInternalEnrichmentsResponse
	3, // 9: com.coralogix.enrichment.v1.EnrichmentInternalService.AddInternalEnrichments:output_type -> com.coralogix.enrichment.v1.AddInternalEnrichmentsResponse
	5, // 10: com.coralogix.enrichment.v1.EnrichmentInternalService.RemoveInternalEnrichments:output_type -> com.coralogix.enrichment.v1.RemoveInternalEnrichmentsResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_init() }
func file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_init() {
	if File_com_coralogix_enrichment_v1_enrichment_internal_service_proto != nil {
		return
	}
	file_com_coralogix_enrichment_v1_enrichment_proto_init()
	file_com_coralogix_enrichment_v1_enrichment_request_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDesc), len(file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_goTypes,
		DependencyIndexes: file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_depIdxs,
		MessageInfos:      file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_msgTypes,
	}.Build()
	File_com_coralogix_enrichment_v1_enrichment_internal_service_proto = out.File
	file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_goTypes = nil
	file_com_coralogix_enrichment_v1_enrichment_internal_service_proto_depIdxs = nil
}
