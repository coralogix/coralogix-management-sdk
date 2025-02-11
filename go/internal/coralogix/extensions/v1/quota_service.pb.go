// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: com/coralogix/extensions/v1/quota_service.proto

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
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetQuotasRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetQuotasRequest) Reset() {
	*x = GetQuotasRequest{}
	mi := &file_com_coralogix_extensions_v1_quota_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetQuotasRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuotasRequest) ProtoMessage() {}

func (x *GetQuotasRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_extensions_v1_quota_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuotasRequest.ProtoReflect.Descriptor instead.
func (*GetQuotasRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_extensions_v1_quota_service_proto_rawDescGZIP(), []int{0}
}

type GetQuotasResponse struct {
	state            protoimpl.MessageState                 `protogen:"open.v1"`
	CompanyId        *wrapperspb.StringValue                `protobuf:"bytes,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Alert            *GetQuotasResponse_Usage               `protobuf:"bytes,2,opt,name=alert,proto3" json:"alert,omitempty"`
	Enrichment       *GetQuotasResponse_Usage               `protobuf:"bytes,3,opt,name=enrichment,proto3" json:"enrichment,omitempty"`
	ParsingRule      *GetQuotasResponse_Usage               `protobuf:"bytes,4,opt,name=parsing_rule,json=parsingRule,proto3" json:"parsing_rule,omitempty"`
	ParsingRuleGroup *GetQuotasResponse_Usage               `protobuf:"bytes,5,opt,name=parsing_rule_group,json=parsingRuleGroup,proto3" json:"parsing_rule_group,omitempty"`
	ParsingTheme     *GetQuotasResponse_Usage               `protobuf:"bytes,6,opt,name=parsing_theme,json=parsingTheme,proto3" json:"parsing_theme,omitempty"`
	DynamicAlert     *GetQuotasResponse_Usage               `protobuf:"bytes,7,opt,name=dynamic_alert,json=dynamicAlert,proto3" json:"dynamic_alert,omitempty"`
	Events_2Metrics  *GetQuotasResponse_Events2MetricsUsage `protobuf:"bytes,8,opt,name=events_2_metrics,json=events2Metrics,proto3" json:"events_2_metrics,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *GetQuotasResponse) Reset() {
	*x = GetQuotasResponse{}
	mi := &file_com_coralogix_extensions_v1_quota_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetQuotasResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuotasResponse) ProtoMessage() {}

func (x *GetQuotasResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_extensions_v1_quota_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuotasResponse.ProtoReflect.Descriptor instead.
func (*GetQuotasResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_extensions_v1_quota_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetQuotasResponse) GetCompanyId() *wrapperspb.StringValue {
	if x != nil {
		return x.CompanyId
	}
	return nil
}

func (x *GetQuotasResponse) GetAlert() *GetQuotasResponse_Usage {
	if x != nil {
		return x.Alert
	}
	return nil
}

func (x *GetQuotasResponse) GetEnrichment() *GetQuotasResponse_Usage {
	if x != nil {
		return x.Enrichment
	}
	return nil
}

func (x *GetQuotasResponse) GetParsingRule() *GetQuotasResponse_Usage {
	if x != nil {
		return x.ParsingRule
	}
	return nil
}

func (x *GetQuotasResponse) GetParsingRuleGroup() *GetQuotasResponse_Usage {
	if x != nil {
		return x.ParsingRuleGroup
	}
	return nil
}

func (x *GetQuotasResponse) GetParsingTheme() *GetQuotasResponse_Usage {
	if x != nil {
		return x.ParsingTheme
	}
	return nil
}

func (x *GetQuotasResponse) GetDynamicAlert() *GetQuotasResponse_Usage {
	if x != nil {
		return x.DynamicAlert
	}
	return nil
}

func (x *GetQuotasResponse) GetEvents_2Metrics() *GetQuotasResponse_Events2MetricsUsage {
	if x != nil {
		return x.Events_2Metrics
	}
	return nil
}

type GetQuotasResponse_Usage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Used          *wrapperspb.Int32Value `protobuf:"bytes,1,opt,name=used,proto3" json:"used,omitempty"`
	Limit         *wrapperspb.Int32Value `protobuf:"bytes,2,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetQuotasResponse_Usage) Reset() {
	*x = GetQuotasResponse_Usage{}
	mi := &file_com_coralogix_extensions_v1_quota_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetQuotasResponse_Usage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuotasResponse_Usage) ProtoMessage() {}

func (x *GetQuotasResponse_Usage) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_extensions_v1_quota_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuotasResponse_Usage.ProtoReflect.Descriptor instead.
func (*GetQuotasResponse_Usage) Descriptor() ([]byte, []int) {
	return file_com_coralogix_extensions_v1_quota_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetQuotasResponse_Usage) GetUsed() *wrapperspb.Int32Value {
	if x != nil {
		return x.Used
	}
	return nil
}

func (x *GetQuotasResponse_Usage) GetLimit() *wrapperspb.Int32Value {
	if x != nil {
		return x.Limit
	}
	return nil
}

type GetQuotasResponse_Events2MetricsUsage struct {
	state         protoimpl.MessageState   `protogen:"open.v1"`
	LabelsLimit   *wrapperspb.Int32Value   `protobuf:"bytes,1,opt,name=labels_limit,json=labelsLimit,proto3" json:"labels_limit,omitempty"`
	Permutations  *GetQuotasResponse_Usage `protobuf:"bytes,2,opt,name=permutations,proto3" json:"permutations,omitempty"`
	Metrics       *GetQuotasResponse_Usage `protobuf:"bytes,3,opt,name=metrics,proto3" json:"metrics,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetQuotasResponse_Events2MetricsUsage) Reset() {
	*x = GetQuotasResponse_Events2MetricsUsage{}
	mi := &file_com_coralogix_extensions_v1_quota_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetQuotasResponse_Events2MetricsUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuotasResponse_Events2MetricsUsage) ProtoMessage() {}

func (x *GetQuotasResponse_Events2MetricsUsage) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_extensions_v1_quota_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuotasResponse_Events2MetricsUsage.ProtoReflect.Descriptor instead.
func (*GetQuotasResponse_Events2MetricsUsage) Descriptor() ([]byte, []int) {
	return file_com_coralogix_extensions_v1_quota_service_proto_rawDescGZIP(), []int{1, 1}
}

func (x *GetQuotasResponse_Events2MetricsUsage) GetLabelsLimit() *wrapperspb.Int32Value {
	if x != nil {
		return x.LabelsLimit
	}
	return nil
}

func (x *GetQuotasResponse_Events2MetricsUsage) GetPermutations() *GetQuotasResponse_Usage {
	if x != nil {
		return x.Permutations
	}
	return nil
}

func (x *GetQuotasResponse_Events2MetricsUsage) GetMetrics() *GetQuotasResponse_Usage {
	if x != nil {
		return x.Metrics
	}
	return nil
}

var File_com_coralogix_extensions_v1_quota_service_proto protoreflect.FileDescriptor

var file_com_coralogix_extensions_v1_quota_service_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75,
	0x6f, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x27,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x51, 0x75,
	0x6f, 0x74, 0x61, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xc2, 0x08, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x4a,
	0x0a, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x51,
	0x75, 0x6f, 0x74, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x54, 0x0a, 0x0a, 0x65, 0x6e,
	0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x51, 0x75, 0x6f, 0x74, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x0a, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x57, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x75, 0x6c, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0b, 0x70, 0x61,
	0x72, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x62, 0x0a, 0x12, 0x70, 0x61, 0x72,
	0x73, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x10, 0x70, 0x61, 0x72,
	0x73, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x59, 0x0a,
	0x0d, 0x70, 0x61, 0x72, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x73,
	0x69, 0x6e, 0x67, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x59, 0x0a, 0x0d, 0x64, 0x79, 0x6e, 0x61,
	0x6d, 0x69, 0x63, 0x5f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0c, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x12, 0x6c, 0x0a, 0x10, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x32, 0x5f,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x51,
	0x75, 0x6f, 0x74, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x32, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x1a, 0x6b, 0x0a, 0x05, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x75, 0x73, 0x65, 0x64, 0x12, 0x31, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x1a, 0xff,
	0x01, 0x0a, 0x13, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x58, 0x0a, 0x0c, 0x70, 0x65, 0x72, 0x6d, 0x75, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75,
	0x6f, 0x74, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x0c, 0x70, 0x65, 0x72, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x4e, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x32, 0xc3, 0x01, 0x0a, 0x0c, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xb2, 0x01, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x73, 0x12,
	0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x51, 0x75, 0x6f, 0x74, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x46,
	0xc2, 0xb8, 0x02, 0x2a, 0x0a, 0x28, 0x47, 0x65, 0x74, 0x20, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x73,
	0x20, 0x66, 0x6f, 0x72, 0x20, 0x76, 0x61, 0x72, 0x69, 0x6f, 0x75, 0x73, 0x20, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_extensions_v1_quota_service_proto_rawDescOnce sync.Once
	file_com_coralogix_extensions_v1_quota_service_proto_rawDescData = file_com_coralogix_extensions_v1_quota_service_proto_rawDesc
)

func file_com_coralogix_extensions_v1_quota_service_proto_rawDescGZIP() []byte {
	file_com_coralogix_extensions_v1_quota_service_proto_rawDescOnce.Do(func() {
		file_com_coralogix_extensions_v1_quota_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_extensions_v1_quota_service_proto_rawDescData)
	})
	return file_com_coralogix_extensions_v1_quota_service_proto_rawDescData
}

var file_com_coralogix_extensions_v1_quota_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_com_coralogix_extensions_v1_quota_service_proto_goTypes = []any{
	(*GetQuotasRequest)(nil),                      // 0: com.coralogix.extensions.v1.GetQuotasRequest
	(*GetQuotasResponse)(nil),                     // 1: com.coralogix.extensions.v1.GetQuotasResponse
	(*GetQuotasResponse_Usage)(nil),               // 2: com.coralogix.extensions.v1.GetQuotasResponse.Usage
	(*GetQuotasResponse_Events2MetricsUsage)(nil), // 3: com.coralogix.extensions.v1.GetQuotasResponse.Events2MetricsUsage
	(*wrapperspb.StringValue)(nil),                // 4: google.protobuf.StringValue
	(*wrapperspb.Int32Value)(nil),                 // 5: google.protobuf.Int32Value
}
var file_com_coralogix_extensions_v1_quota_service_proto_depIdxs = []int32{
	4,  // 0: com.coralogix.extensions.v1.GetQuotasResponse.company_id:type_name -> google.protobuf.StringValue
	2,  // 1: com.coralogix.extensions.v1.GetQuotasResponse.alert:type_name -> com.coralogix.extensions.v1.GetQuotasResponse.Usage
	2,  // 2: com.coralogix.extensions.v1.GetQuotasResponse.enrichment:type_name -> com.coralogix.extensions.v1.GetQuotasResponse.Usage
	2,  // 3: com.coralogix.extensions.v1.GetQuotasResponse.parsing_rule:type_name -> com.coralogix.extensions.v1.GetQuotasResponse.Usage
	2,  // 4: com.coralogix.extensions.v1.GetQuotasResponse.parsing_rule_group:type_name -> com.coralogix.extensions.v1.GetQuotasResponse.Usage
	2,  // 5: com.coralogix.extensions.v1.GetQuotasResponse.parsing_theme:type_name -> com.coralogix.extensions.v1.GetQuotasResponse.Usage
	2,  // 6: com.coralogix.extensions.v1.GetQuotasResponse.dynamic_alert:type_name -> com.coralogix.extensions.v1.GetQuotasResponse.Usage
	3,  // 7: com.coralogix.extensions.v1.GetQuotasResponse.events_2_metrics:type_name -> com.coralogix.extensions.v1.GetQuotasResponse.Events2MetricsUsage
	5,  // 8: com.coralogix.extensions.v1.GetQuotasResponse.Usage.used:type_name -> google.protobuf.Int32Value
	5,  // 9: com.coralogix.extensions.v1.GetQuotasResponse.Usage.limit:type_name -> google.protobuf.Int32Value
	5,  // 10: com.coralogix.extensions.v1.GetQuotasResponse.Events2MetricsUsage.labels_limit:type_name -> google.protobuf.Int32Value
	2,  // 11: com.coralogix.extensions.v1.GetQuotasResponse.Events2MetricsUsage.permutations:type_name -> com.coralogix.extensions.v1.GetQuotasResponse.Usage
	2,  // 12: com.coralogix.extensions.v1.GetQuotasResponse.Events2MetricsUsage.metrics:type_name -> com.coralogix.extensions.v1.GetQuotasResponse.Usage
	0,  // 13: com.coralogix.extensions.v1.QuotaService.GetQuotas:input_type -> com.coralogix.extensions.v1.GetQuotasRequest
	1,  // 14: com.coralogix.extensions.v1.QuotaService.GetQuotas:output_type -> com.coralogix.extensions.v1.GetQuotasResponse
	14, // [14:15] is the sub-list for method output_type
	13, // [13:14] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_com_coralogix_extensions_v1_quota_service_proto_init() }
func file_com_coralogix_extensions_v1_quota_service_proto_init() {
	if File_com_coralogix_extensions_v1_quota_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_extensions_v1_quota_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogix_extensions_v1_quota_service_proto_goTypes,
		DependencyIndexes: file_com_coralogix_extensions_v1_quota_service_proto_depIdxs,
		MessageInfos:      file_com_coralogix_extensions_v1_quota_service_proto_msgTypes,
	}.Build()
	File_com_coralogix_extensions_v1_quota_service_proto = out.File
	file_com_coralogix_extensions_v1_quota_service_proto_rawDesc = nil
	file_com_coralogix_extensions_v1_quota_service_proto_goTypes = nil
	file_com_coralogix_extensions_v1_quota_service_proto_depIdxs = nil
}
