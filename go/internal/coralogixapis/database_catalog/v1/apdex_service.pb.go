// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.26.1
// source: com/coralogixapis/database_catalog/v1/apdex_service.proto

package v1

import (
	_ "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/apm/common/v2"
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

type GetApdexThresholdRequest struct {
	state              protoimpl.MessageState  `protogen:"open.v1"`
	DatabaseName       *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	DatabaseNameSource *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=database_name_source,json=databaseNameSource,proto3" json:"database_name_source,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *GetApdexThresholdRequest) Reset() {
	*x = GetApdexThresholdRequest{}
	mi := &file_com_coralogixapis_database_catalog_v1_apdex_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetApdexThresholdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApdexThresholdRequest) ProtoMessage() {}

func (x *GetApdexThresholdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_database_catalog_v1_apdex_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApdexThresholdRequest.ProtoReflect.Descriptor instead.
func (*GetApdexThresholdRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_database_catalog_v1_apdex_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetApdexThresholdRequest) GetDatabaseName() *wrapperspb.StringValue {
	if x != nil {
		return x.DatabaseName
	}
	return nil
}

func (x *GetApdexThresholdRequest) GetDatabaseNameSource() *wrapperspb.StringValue {
	if x != nil {
		return x.DatabaseNameSource
	}
	return nil
}

type GetApdexThresholdResponse struct {
	state              protoimpl.MessageState  `protogen:"open.v1"`
	DatabaseName       *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	DatabaseNameSource *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=database_name_source,json=databaseNameSource,proto3" json:"database_name_source,omitempty"`
	ApdexThresholdUs   *wrapperspb.UInt64Value `protobuf:"bytes,3,opt,name=apdex_threshold_us,json=apdexThresholdUs,proto3" json:"apdex_threshold_us,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *GetApdexThresholdResponse) Reset() {
	*x = GetApdexThresholdResponse{}
	mi := &file_com_coralogixapis_database_catalog_v1_apdex_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetApdexThresholdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApdexThresholdResponse) ProtoMessage() {}

func (x *GetApdexThresholdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_database_catalog_v1_apdex_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApdexThresholdResponse.ProtoReflect.Descriptor instead.
func (*GetApdexThresholdResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_database_catalog_v1_apdex_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetApdexThresholdResponse) GetDatabaseName() *wrapperspb.StringValue {
	if x != nil {
		return x.DatabaseName
	}
	return nil
}

func (x *GetApdexThresholdResponse) GetDatabaseNameSource() *wrapperspb.StringValue {
	if x != nil {
		return x.DatabaseNameSource
	}
	return nil
}

func (x *GetApdexThresholdResponse) GetApdexThresholdUs() *wrapperspb.UInt64Value {
	if x != nil {
		return x.ApdexThresholdUs
	}
	return nil
}

type ReplaceApdexThresholdRequest struct {
	state              protoimpl.MessageState  `protogen:"open.v1"`
	DatabaseName       *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	DatabaseNameSource *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=database_name_source,json=databaseNameSource,proto3" json:"database_name_source,omitempty"`
	ApdexThresholdUs   *wrapperspb.UInt64Value `protobuf:"bytes,3,opt,name=apdex_threshold_us,json=apdexThresholdUs,proto3" json:"apdex_threshold_us,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *ReplaceApdexThresholdRequest) Reset() {
	*x = ReplaceApdexThresholdRequest{}
	mi := &file_com_coralogixapis_database_catalog_v1_apdex_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReplaceApdexThresholdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplaceApdexThresholdRequest) ProtoMessage() {}

func (x *ReplaceApdexThresholdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_database_catalog_v1_apdex_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplaceApdexThresholdRequest.ProtoReflect.Descriptor instead.
func (*ReplaceApdexThresholdRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_database_catalog_v1_apdex_service_proto_rawDescGZIP(), []int{2}
}

func (x *ReplaceApdexThresholdRequest) GetDatabaseName() *wrapperspb.StringValue {
	if x != nil {
		return x.DatabaseName
	}
	return nil
}

func (x *ReplaceApdexThresholdRequest) GetDatabaseNameSource() *wrapperspb.StringValue {
	if x != nil {
		return x.DatabaseNameSource
	}
	return nil
}

func (x *ReplaceApdexThresholdRequest) GetApdexThresholdUs() *wrapperspb.UInt64Value {
	if x != nil {
		return x.ApdexThresholdUs
	}
	return nil
}

type ReplaceApdexThresholdResponse struct {
	state              protoimpl.MessageState  `protogen:"open.v1"`
	DatabaseName       *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	DatabaseNameSource *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=database_name_source,json=databaseNameSource,proto3" json:"database_name_source,omitempty"`
	ApdexThresholdUs   *wrapperspb.UInt64Value `protobuf:"bytes,3,opt,name=apdex_threshold_us,json=apdexThresholdUs,proto3" json:"apdex_threshold_us,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *ReplaceApdexThresholdResponse) Reset() {
	*x = ReplaceApdexThresholdResponse{}
	mi := &file_com_coralogixapis_database_catalog_v1_apdex_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReplaceApdexThresholdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplaceApdexThresholdResponse) ProtoMessage() {}

func (x *ReplaceApdexThresholdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_database_catalog_v1_apdex_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplaceApdexThresholdResponse.ProtoReflect.Descriptor instead.
func (*ReplaceApdexThresholdResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_database_catalog_v1_apdex_service_proto_rawDescGZIP(), []int{3}
}

func (x *ReplaceApdexThresholdResponse) GetDatabaseName() *wrapperspb.StringValue {
	if x != nil {
		return x.DatabaseName
	}
	return nil
}

func (x *ReplaceApdexThresholdResponse) GetDatabaseNameSource() *wrapperspb.StringValue {
	if x != nil {
		return x.DatabaseNameSource
	}
	return nil
}

func (x *ReplaceApdexThresholdResponse) GetApdexThresholdUs() *wrapperspb.UInt64Value {
	if x != nil {
		return x.ApdexThresholdUs
	}
	return nil
}

type ApdexThresholdExecutionRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Request:
	//
	//	*ApdexThresholdExecutionRequest_Replace
	Request       isApdexThresholdExecutionRequest_Request `protobuf_oneof:"request"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ApdexThresholdExecutionRequest) Reset() {
	*x = ApdexThresholdExecutionRequest{}
	mi := &file_com_coralogixapis_database_catalog_v1_apdex_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApdexThresholdExecutionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApdexThresholdExecutionRequest) ProtoMessage() {}

func (x *ApdexThresholdExecutionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_database_catalog_v1_apdex_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApdexThresholdExecutionRequest.ProtoReflect.Descriptor instead.
func (*ApdexThresholdExecutionRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_database_catalog_v1_apdex_service_proto_rawDescGZIP(), []int{4}
}

func (x *ApdexThresholdExecutionRequest) GetRequest() isApdexThresholdExecutionRequest_Request {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *ApdexThresholdExecutionRequest) GetReplace() *ReplaceApdexThresholdRequest {
	if x != nil {
		if x, ok := x.Request.(*ApdexThresholdExecutionRequest_Replace); ok {
			return x.Replace
		}
	}
	return nil
}

type isApdexThresholdExecutionRequest_Request interface {
	isApdexThresholdExecutionRequest_Request()
}

type ApdexThresholdExecutionRequest_Replace struct {
	Replace *ReplaceApdexThresholdRequest `protobuf:"bytes,1,opt,name=replace,proto3,oneof"`
}

func (*ApdexThresholdExecutionRequest_Replace) isApdexThresholdExecutionRequest_Request() {}

type ApdexThresholdExecutionResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Response:
	//
	//	*ApdexThresholdExecutionResponse_Replace
	Response      isApdexThresholdExecutionResponse_Response `protobuf_oneof:"response"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ApdexThresholdExecutionResponse) Reset() {
	*x = ApdexThresholdExecutionResponse{}
	mi := &file_com_coralogixapis_database_catalog_v1_apdex_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApdexThresholdExecutionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApdexThresholdExecutionResponse) ProtoMessage() {}

func (x *ApdexThresholdExecutionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_database_catalog_v1_apdex_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApdexThresholdExecutionResponse.ProtoReflect.Descriptor instead.
func (*ApdexThresholdExecutionResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_database_catalog_v1_apdex_service_proto_rawDescGZIP(), []int{5}
}

func (x *ApdexThresholdExecutionResponse) GetResponse() isApdexThresholdExecutionResponse_Response {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *ApdexThresholdExecutionResponse) GetReplace() *ReplaceApdexThresholdResponse {
	if x != nil {
		if x, ok := x.Response.(*ApdexThresholdExecutionResponse_Replace); ok {
			return x.Replace
		}
	}
	return nil
}

type isApdexThresholdExecutionResponse_Response interface {
	isApdexThresholdExecutionResponse_Response()
}

type ApdexThresholdExecutionResponse_Replace struct {
	Replace *ReplaceApdexThresholdResponse `protobuf:"bytes,1,opt,name=replace,proto3,oneof"`
}

func (*ApdexThresholdExecutionResponse_Replace) isApdexThresholdExecutionResponse_Response() {}

type BatchExecuteApdexThresholdRequest struct {
	state         protoimpl.MessageState            `protogen:"open.v1"`
	Requests      []*ApdexThresholdExecutionRequest `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BatchExecuteApdexThresholdRequest) Reset() {
	*x = BatchExecuteApdexThresholdRequest{}
	mi := &file_com_coralogixapis_database_catalog_v1_apdex_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchExecuteApdexThresholdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchExecuteApdexThresholdRequest) ProtoMessage() {}

func (x *BatchExecuteApdexThresholdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_database_catalog_v1_apdex_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchExecuteApdexThresholdRequest.ProtoReflect.Descriptor instead.
func (*BatchExecuteApdexThresholdRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_database_catalog_v1_apdex_service_proto_rawDescGZIP(), []int{6}
}

func (x *BatchExecuteApdexThresholdRequest) GetRequests() []*ApdexThresholdExecutionRequest {
	if x != nil {
		return x.Requests
	}
	return nil
}

type BatchExecuteApdexThresholdResponse struct {
	state         protoimpl.MessageState             `protogen:"open.v1"`
	Responses     []*ApdexThresholdExecutionResponse `protobuf:"bytes,1,rep,name=responses,proto3" json:"responses,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BatchExecuteApdexThresholdResponse) Reset() {
	*x = BatchExecuteApdexThresholdResponse{}
	mi := &file_com_coralogixapis_database_catalog_v1_apdex_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchExecuteApdexThresholdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchExecuteApdexThresholdResponse) ProtoMessage() {}

func (x *BatchExecuteApdexThresholdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_database_catalog_v1_apdex_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchExecuteApdexThresholdResponse.ProtoReflect.Descriptor instead.
func (*BatchExecuteApdexThresholdResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_database_catalog_v1_apdex_service_proto_rawDescGZIP(), []int{7}
}

func (x *BatchExecuteApdexThresholdResponse) GetResponses() []*ApdexThresholdExecutionResponse {
	if x != nil {
		return x.Responses
	}
	return nil
}

var File_com_coralogixapis_database_catalog_v1_apdex_service_proto protoreflect.FileDescriptor

var file_com_coralogixapis_database_catalog_v1_apdex_service_proto_rawDesc = []byte{
	0x0a, 0x39, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x64, 0x65, 0x78, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e,
	0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x32, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xad, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x70, 0x64, 0x65, 0x78,
	0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x41, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x4e, 0x0a, 0x14, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x12, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x22, 0xfa, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x70, 0x64, 0x65, 0x78,
	0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x41, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4e, 0x0a, 0x14, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x12, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x12, 0x61, 0x70, 0x64, 0x65, 0x78, 0x5f, 0x74, 0x68,
	0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x5f, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10,
	0x61, 0x70, 0x64, 0x65, 0x78, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x55, 0x73,
	0x22, 0xfd, 0x01, 0x0a, 0x1c, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x41, 0x70, 0x64, 0x65,
	0x78, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x41, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4e, 0x0a, 0x14, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x12, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x12, 0x61, 0x70, 0x64, 0x65, 0x78, 0x5f, 0x74, 0x68,
	0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x5f, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10,
	0x61, 0x70, 0x64, 0x65, 0x78, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x55, 0x73,
	0x22, 0xfe, 0x01, 0x0a, 0x1d, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x41, 0x70, 0x64, 0x65,
	0x78, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x41, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4e, 0x0a, 0x14, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x12, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x12, 0x61, 0x70, 0x64, 0x65, 0x78, 0x5f, 0x74,
	0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x5f, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x10, 0x61, 0x70, 0x64, 0x65, 0x78, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x55,
	0x73, 0x22, 0x8c, 0x01, 0x0a, 0x1e, 0x41, 0x70, 0x64, 0x65, 0x78, 0x54, 0x68, 0x72, 0x65, 0x73,
	0x68, 0x6f, 0x6c, 0x64, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x5f, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x41, 0x70, 0x64, 0x65, 0x78, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68,
	0x6f, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x8f, 0x01, 0x0a, 0x1f, 0x41, 0x70, 0x64, 0x65, 0x78, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68,
	0x6f, 0x6c, 0x64, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x41, 0x70, 0x64, 0x65, 0x78, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68,
	0x6f, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x07, 0x72,
	0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x21, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x41, 0x70, 0x64, 0x65, 0x78, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x61, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x70, 0x64, 0x65, 0x78, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0x8a, 0x01, 0x0a, 0x22,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x41, 0x70, 0x64, 0x65,
	0x78, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x64, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70,
	0x64, 0x65, 0x78, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x09, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x32, 0xe0, 0x04, 0x0a, 0x0c, 0x41, 0x70, 0x64,
	0x65, 0x78, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xb1, 0x01, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x41, 0x70, 0x64, 0x65, 0x78, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12,
	0x3f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x64, 0x65, 0x78,
	0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x40, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x64, 0x65,
	0x78, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x19, 0xc2, 0xb8, 0x02, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x20, 0x41, 0x70,
	0x64, 0x65, 0x78, 0x20, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0xc1, 0x01,
	0x0a, 0x15, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x41, 0x70, 0x64, 0x65, 0x78, 0x54, 0x68,
	0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x43, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x41, 0x70, 0x64, 0x65, 0x78, 0x54, 0x68, 0x72, 0x65,
	0x73, 0x68, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x44, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x41, 0x70, 0x64, 0x65,
	0x78, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1d, 0xc2, 0xb8, 0x02, 0x19, 0x0a, 0x17, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x20, 0x41, 0x70, 0x64, 0x65, 0x78, 0x20, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x12, 0xd7, 0x01, 0x0a, 0x1a, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x41, 0x70, 0x64, 0x65, 0x78, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x12, 0x48, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x65, 0x41, 0x70, 0x64, 0x65, 0x78, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68,
	0x6f, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x49, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x41,
	0x70, 0x64, 0x65, 0x78, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0xc2, 0xb8, 0x02, 0x20, 0x0a, 0x1e, 0x20, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x20, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x20, 0x41, 0x70, 0x64, 0x65,
	0x78, 0x20, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_database_catalog_v1_apdex_service_proto_rawDescOnce sync.Once
	file_com_coralogixapis_database_catalog_v1_apdex_service_proto_rawDescData = file_com_coralogixapis_database_catalog_v1_apdex_service_proto_rawDesc
)

func file_com_coralogixapis_database_catalog_v1_apdex_service_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_database_catalog_v1_apdex_service_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_database_catalog_v1_apdex_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_database_catalog_v1_apdex_service_proto_rawDescData)
	})
	return file_com_coralogixapis_database_catalog_v1_apdex_service_proto_rawDescData
}

var file_com_coralogixapis_database_catalog_v1_apdex_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_com_coralogixapis_database_catalog_v1_apdex_service_proto_goTypes = []any{
	(*GetApdexThresholdRequest)(nil),           // 0: com.coralogixapis.database_catalog.v1.GetApdexThresholdRequest
	(*GetApdexThresholdResponse)(nil),          // 1: com.coralogixapis.database_catalog.v1.GetApdexThresholdResponse
	(*ReplaceApdexThresholdRequest)(nil),       // 2: com.coralogixapis.database_catalog.v1.ReplaceApdexThresholdRequest
	(*ReplaceApdexThresholdResponse)(nil),      // 3: com.coralogixapis.database_catalog.v1.ReplaceApdexThresholdResponse
	(*ApdexThresholdExecutionRequest)(nil),     // 4: com.coralogixapis.database_catalog.v1.ApdexThresholdExecutionRequest
	(*ApdexThresholdExecutionResponse)(nil),    // 5: com.coralogixapis.database_catalog.v1.ApdexThresholdExecutionResponse
	(*BatchExecuteApdexThresholdRequest)(nil),  // 6: com.coralogixapis.database_catalog.v1.BatchExecuteApdexThresholdRequest
	(*BatchExecuteApdexThresholdResponse)(nil), // 7: com.coralogixapis.database_catalog.v1.BatchExecuteApdexThresholdResponse
	(*wrapperspb.StringValue)(nil),             // 8: google.protobuf.StringValue
	(*wrapperspb.UInt64Value)(nil),             // 9: google.protobuf.UInt64Value
}
var file_com_coralogixapis_database_catalog_v1_apdex_service_proto_depIdxs = []int32{
	8,  // 0: com.coralogixapis.database_catalog.v1.GetApdexThresholdRequest.database_name:type_name -> google.protobuf.StringValue
	8,  // 1: com.coralogixapis.database_catalog.v1.GetApdexThresholdRequest.database_name_source:type_name -> google.protobuf.StringValue
	8,  // 2: com.coralogixapis.database_catalog.v1.GetApdexThresholdResponse.database_name:type_name -> google.protobuf.StringValue
	8,  // 3: com.coralogixapis.database_catalog.v1.GetApdexThresholdResponse.database_name_source:type_name -> google.protobuf.StringValue
	9,  // 4: com.coralogixapis.database_catalog.v1.GetApdexThresholdResponse.apdex_threshold_us:type_name -> google.protobuf.UInt64Value
	8,  // 5: com.coralogixapis.database_catalog.v1.ReplaceApdexThresholdRequest.database_name:type_name -> google.protobuf.StringValue
	8,  // 6: com.coralogixapis.database_catalog.v1.ReplaceApdexThresholdRequest.database_name_source:type_name -> google.protobuf.StringValue
	9,  // 7: com.coralogixapis.database_catalog.v1.ReplaceApdexThresholdRequest.apdex_threshold_us:type_name -> google.protobuf.UInt64Value
	8,  // 8: com.coralogixapis.database_catalog.v1.ReplaceApdexThresholdResponse.database_name:type_name -> google.protobuf.StringValue
	8,  // 9: com.coralogixapis.database_catalog.v1.ReplaceApdexThresholdResponse.database_name_source:type_name -> google.protobuf.StringValue
	9,  // 10: com.coralogixapis.database_catalog.v1.ReplaceApdexThresholdResponse.apdex_threshold_us:type_name -> google.protobuf.UInt64Value
	2,  // 11: com.coralogixapis.database_catalog.v1.ApdexThresholdExecutionRequest.replace:type_name -> com.coralogixapis.database_catalog.v1.ReplaceApdexThresholdRequest
	3,  // 12: com.coralogixapis.database_catalog.v1.ApdexThresholdExecutionResponse.replace:type_name -> com.coralogixapis.database_catalog.v1.ReplaceApdexThresholdResponse
	4,  // 13: com.coralogixapis.database_catalog.v1.BatchExecuteApdexThresholdRequest.requests:type_name -> com.coralogixapis.database_catalog.v1.ApdexThresholdExecutionRequest
	5,  // 14: com.coralogixapis.database_catalog.v1.BatchExecuteApdexThresholdResponse.responses:type_name -> com.coralogixapis.database_catalog.v1.ApdexThresholdExecutionResponse
	0,  // 15: com.coralogixapis.database_catalog.v1.ApdexService.GetApdexThreshold:input_type -> com.coralogixapis.database_catalog.v1.GetApdexThresholdRequest
	2,  // 16: com.coralogixapis.database_catalog.v1.ApdexService.ReplaceApdexThreshold:input_type -> com.coralogixapis.database_catalog.v1.ReplaceApdexThresholdRequest
	6,  // 17: com.coralogixapis.database_catalog.v1.ApdexService.BatchExecuteApdexThreshold:input_type -> com.coralogixapis.database_catalog.v1.BatchExecuteApdexThresholdRequest
	1,  // 18: com.coralogixapis.database_catalog.v1.ApdexService.GetApdexThreshold:output_type -> com.coralogixapis.database_catalog.v1.GetApdexThresholdResponse
	3,  // 19: com.coralogixapis.database_catalog.v1.ApdexService.ReplaceApdexThreshold:output_type -> com.coralogixapis.database_catalog.v1.ReplaceApdexThresholdResponse
	7,  // 20: com.coralogixapis.database_catalog.v1.ApdexService.BatchExecuteApdexThreshold:output_type -> com.coralogixapis.database_catalog.v1.BatchExecuteApdexThresholdResponse
	18, // [18:21] is the sub-list for method output_type
	15, // [15:18] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_database_catalog_v1_apdex_service_proto_init() }
func file_com_coralogixapis_database_catalog_v1_apdex_service_proto_init() {
	if File_com_coralogixapis_database_catalog_v1_apdex_service_proto != nil {
		return
	}
	file_com_coralogixapis_database_catalog_v1_apdex_service_proto_msgTypes[4].OneofWrappers = []any{
		(*ApdexThresholdExecutionRequest_Replace)(nil),
	}
	file_com_coralogixapis_database_catalog_v1_apdex_service_proto_msgTypes[5].OneofWrappers = []any{
		(*ApdexThresholdExecutionResponse_Replace)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_database_catalog_v1_apdex_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogixapis_database_catalog_v1_apdex_service_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_database_catalog_v1_apdex_service_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_database_catalog_v1_apdex_service_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_database_catalog_v1_apdex_service_proto = out.File
	file_com_coralogixapis_database_catalog_v1_apdex_service_proto_rawDesc = nil
	file_com_coralogixapis_database_catalog_v1_apdex_service_proto_goTypes = nil
	file_com_coralogixapis_database_catalog_v1_apdex_service_proto_depIdxs = nil
}
