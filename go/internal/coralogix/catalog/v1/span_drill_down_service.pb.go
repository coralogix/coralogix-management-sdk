// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: com/coralogix/catalog/v1/span_drill_down_service.proto

package v1

import (
	_ "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/apm/common/v2"
	_ "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/apm/widgets/v1"
	v1 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/service_catalog/v1"
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

type GetServiceSideModelGraphsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName           *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	Operation             *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	TimeRange             *TimeRange              `protobuf:"bytes,3,opt,name=time_range,json=timeRange,proto3" json:"time_range,omitempty"`
	ServiceCatalogFilters []*ApmFilter            `protobuf:"bytes,4,rep,name=service_catalog_filters,json=serviceCatalogFilters,proto3" json:"service_catalog_filters,omitempty"`
	SpanKind              []SpanKind              `protobuf:"varint,5,rep,packed,name=span_kind,json=spanKind,proto3,enum=com.coralogix.catalog.v1.SpanKind" json:"span_kind,omitempty"`
	Method                *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *GetServiceSideModelGraphsRequest) Reset() {
	*x = GetServiceSideModelGraphsRequest{}
	mi := &file_com_coralogix_catalog_v1_span_drill_down_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetServiceSideModelGraphsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceSideModelGraphsRequest) ProtoMessage() {}

func (x *GetServiceSideModelGraphsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_catalog_v1_span_drill_down_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServiceSideModelGraphsRequest.ProtoReflect.Descriptor instead.
func (*GetServiceSideModelGraphsRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_catalog_v1_span_drill_down_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetServiceSideModelGraphsRequest) GetServiceName() *wrapperspb.StringValue {
	if x != nil {
		return x.ServiceName
	}
	return nil
}

func (x *GetServiceSideModelGraphsRequest) GetOperation() *wrapperspb.StringValue {
	if x != nil {
		return x.Operation
	}
	return nil
}

func (x *GetServiceSideModelGraphsRequest) GetTimeRange() *TimeRange {
	if x != nil {
		return x.TimeRange
	}
	return nil
}

func (x *GetServiceSideModelGraphsRequest) GetServiceCatalogFilters() []*ApmFilter {
	if x != nil {
		return x.ServiceCatalogFilters
	}
	return nil
}

func (x *GetServiceSideModelGraphsRequest) GetSpanKind() []SpanKind {
	if x != nil {
		return x.SpanKind
	}
	return nil
}

func (x *GetServiceSideModelGraphsRequest) GetMethod() *wrapperspb.StringValue {
	if x != nil {
		return x.Method
	}
	return nil
}

type GetServiceSideModelGraphsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Widget:
	//
	//	*GetServiceSideModelGraphsResponse_SpanDrillDownSideModalStats
	//	*GetServiceSideModelGraphsResponse_RequestsAndErrors
	//	*GetServiceSideModelGraphsResponse_LatencyPercentiles
	//	*GetServiceSideModelGraphsResponse_RequestsAndErrorsPerSecond
	//	*GetServiceSideModelGraphsResponse_LatencyPercentilesPerSecond
	Widget isGetServiceSideModelGraphsResponse_Widget `protobuf_oneof:"widget"`
}

func (x *GetServiceSideModelGraphsResponse) Reset() {
	*x = GetServiceSideModelGraphsResponse{}
	mi := &file_com_coralogix_catalog_v1_span_drill_down_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetServiceSideModelGraphsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceSideModelGraphsResponse) ProtoMessage() {}

func (x *GetServiceSideModelGraphsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_catalog_v1_span_drill_down_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServiceSideModelGraphsResponse.ProtoReflect.Descriptor instead.
func (*GetServiceSideModelGraphsResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_catalog_v1_span_drill_down_service_proto_rawDescGZIP(), []int{1}
}

func (m *GetServiceSideModelGraphsResponse) GetWidget() isGetServiceSideModelGraphsResponse_Widget {
	if m != nil {
		return m.Widget
	}
	return nil
}

func (x *GetServiceSideModelGraphsResponse) GetSpanDrillDownSideModalStats() *v1.RepeatedStats {
	if x, ok := x.GetWidget().(*GetServiceSideModelGraphsResponse_SpanDrillDownSideModalStats); ok {
		return x.SpanDrillDownSideModalStats
	}
	return nil
}

func (x *GetServiceSideModelGraphsResponse) GetRequestsAndErrors() *v1.RepeatedLineChart {
	if x, ok := x.GetWidget().(*GetServiceSideModelGraphsResponse_RequestsAndErrors); ok {
		return x.RequestsAndErrors
	}
	return nil
}

func (x *GetServiceSideModelGraphsResponse) GetLatencyPercentiles() *v1.RepeatedLineChart {
	if x, ok := x.GetWidget().(*GetServiceSideModelGraphsResponse_LatencyPercentiles); ok {
		return x.LatencyPercentiles
	}
	return nil
}

func (x *GetServiceSideModelGraphsResponse) GetRequestsAndErrorsPerSecond() *v1.RepeatedLineChart {
	if x, ok := x.GetWidget().(*GetServiceSideModelGraphsResponse_RequestsAndErrorsPerSecond); ok {
		return x.RequestsAndErrorsPerSecond
	}
	return nil
}

func (x *GetServiceSideModelGraphsResponse) GetLatencyPercentilesPerSecond() *v1.RepeatedLineChart {
	if x, ok := x.GetWidget().(*GetServiceSideModelGraphsResponse_LatencyPercentilesPerSecond); ok {
		return x.LatencyPercentilesPerSecond
	}
	return nil
}

type isGetServiceSideModelGraphsResponse_Widget interface {
	isGetServiceSideModelGraphsResponse_Widget()
}

type GetServiceSideModelGraphsResponse_SpanDrillDownSideModalStats struct {
	SpanDrillDownSideModalStats *v1.RepeatedStats `protobuf:"bytes,1,opt,name=span_drill_down_side_modal_stats,json=spanDrillDownSideModalStats,proto3,oneof"`
}

type GetServiceSideModelGraphsResponse_RequestsAndErrors struct {
	RequestsAndErrors *v1.RepeatedLineChart `protobuf:"bytes,2,opt,name=requests_and_errors,json=requestsAndErrors,proto3,oneof"`
}

type GetServiceSideModelGraphsResponse_LatencyPercentiles struct {
	LatencyPercentiles *v1.RepeatedLineChart `protobuf:"bytes,3,opt,name=latency_percentiles,json=latencyPercentiles,proto3,oneof"`
}

type GetServiceSideModelGraphsResponse_RequestsAndErrorsPerSecond struct {
	RequestsAndErrorsPerSecond *v1.RepeatedLineChart `protobuf:"bytes,4,opt,name=requests_and_errors_per_second,json=requestsAndErrorsPerSecond,proto3,oneof"`
}

type GetServiceSideModelGraphsResponse_LatencyPercentilesPerSecond struct {
	LatencyPercentilesPerSecond *v1.RepeatedLineChart `protobuf:"bytes,5,opt,name=latency_percentiles_per_second,json=latencyPercentilesPerSecond,proto3,oneof"`
}

func (*GetServiceSideModelGraphsResponse_SpanDrillDownSideModalStats) isGetServiceSideModelGraphsResponse_Widget() {
}

func (*GetServiceSideModelGraphsResponse_RequestsAndErrors) isGetServiceSideModelGraphsResponse_Widget() {
}

func (*GetServiceSideModelGraphsResponse_LatencyPercentiles) isGetServiceSideModelGraphsResponse_Widget() {
}

func (*GetServiceSideModelGraphsResponse_RequestsAndErrorsPerSecond) isGetServiceSideModelGraphsResponse_Widget() {
}

func (*GetServiceSideModelGraphsResponse_LatencyPercentilesPerSecond) isGetServiceSideModelGraphsResponse_Widget() {
}

var File_com_coralogix_catalog_v1_span_drill_down_service_proto protoreflect.FileDescriptor

var file_com_coralogix_catalog_v1_span_drill_down_service_proto_rawDesc = []byte{
	0x0a, 0x36, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x5f,
	0x64, 0x72, 0x69, 0x6c, 0x6c, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e,
	0x76, 0x31, 0x1a, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x32, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x70, 0x6d, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x3b, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x03, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x53, 0x69, 0x64, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0c, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x5b, 0x0a, 0x17, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x6d, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x52, 0x15, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x3f, 0x0a, 0x09, 0x73, 0x70, 0x61, 0x6e,
	0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x4b, 0x69, 0x6e, 0x64, 0x52,
	0x08, 0x73, 0x70, 0x61, 0x6e, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x34, 0x0a, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22,
	0x81, 0x05, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x69,
	0x64, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x47, 0x72, 0x61, 0x70, 0x68, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7c, 0x0a, 0x20, 0x73, 0x70, 0x61, 0x6e, 0x5f, 0x64, 0x72,
	0x69, 0x6c, 0x6c, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x5f, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x6d, 0x6f,
	0x64, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x48, 0x00, 0x52, 0x1b, 0x73, 0x70, 0x61, 0x6e, 0x44, 0x72, 0x69, 0x6c,
	0x6c, 0x44, 0x6f, 0x77, 0x6e, 0x53, 0x69, 0x64, 0x65, 0x4d, 0x6f, 0x64, 0x61, 0x6c, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x12, 0x69, 0x0a, 0x13, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x5f,
	0x61, 0x6e, 0x64, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x4c, 0x69, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x72, 0x74, 0x48, 0x00, 0x52, 0x11, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x41, 0x6e, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x6a,
	0x0a, 0x13, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x43,
	0x68, 0x61, 0x72, 0x74, 0x48, 0x00, 0x52, 0x12, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x50,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x7d, 0x0a, 0x1e, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x72, 0x74, 0x48, 0x00, 0x52, 0x1a, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x41, 0x6e, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x7e, 0x0a, 0x1e, 0x6c, 0x61, 0x74,
	0x65, 0x6e, 0x63, 0x79, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x6c, 0x65, 0x73,
	0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x4c, 0x69, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x72, 0x74, 0x48, 0x00, 0x52, 0x1b, 0x6c, 0x61,
	0x74, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x6c, 0x65, 0x73,
	0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x77, 0x69, 0x64,
	0x67, 0x65, 0x74, 0x32, 0xd4, 0x01, 0x0a, 0x14, 0x53, 0x70, 0x61, 0x6e, 0x44, 0x72, 0x69, 0x6c,
	0x6c, 0x44, 0x6f, 0x77, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xbb, 0x01, 0x0a,
	0x19, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x69, 0x64, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x47, 0x72, 0x61, 0x70, 0x68, 0x73, 0x12, 0x3a, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x53, 0x69, 0x64, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x47, 0x72, 0x61, 0x70, 0x68, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x69, 0x64, 0x65,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x47, 0x72, 0x61, 0x70, 0x68, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x23, 0xc2, 0xb8, 0x02, 0x1f, 0x0a, 0x1d, 0x67, 0x65, 0x74, 0x20, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x73, 0x69, 0x64, 0x65, 0x20, 0x6d, 0x6f, 0x64, 0x61,
	0x6c, 0x20, 0x67, 0x72, 0x61, 0x70, 0x68, 0x73, 0x30, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_com_coralogix_catalog_v1_span_drill_down_service_proto_rawDescOnce sync.Once
	file_com_coralogix_catalog_v1_span_drill_down_service_proto_rawDescData = file_com_coralogix_catalog_v1_span_drill_down_service_proto_rawDesc
)

func file_com_coralogix_catalog_v1_span_drill_down_service_proto_rawDescGZIP() []byte {
	file_com_coralogix_catalog_v1_span_drill_down_service_proto_rawDescOnce.Do(func() {
		file_com_coralogix_catalog_v1_span_drill_down_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_catalog_v1_span_drill_down_service_proto_rawDescData)
	})
	return file_com_coralogix_catalog_v1_span_drill_down_service_proto_rawDescData
}

var file_com_coralogix_catalog_v1_span_drill_down_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogix_catalog_v1_span_drill_down_service_proto_goTypes = []any{
	(*GetServiceSideModelGraphsRequest)(nil),  // 0: com.coralogix.catalog.v1.GetServiceSideModelGraphsRequest
	(*GetServiceSideModelGraphsResponse)(nil), // 1: com.coralogix.catalog.v1.GetServiceSideModelGraphsResponse
	(*wrapperspb.StringValue)(nil),            // 2: google.protobuf.StringValue
	(*TimeRange)(nil),                         // 3: com.coralogix.catalog.v1.TimeRange
	(*ApmFilter)(nil),                         // 4: com.coralogix.catalog.v1.ApmFilter
	(SpanKind)(0),                             // 5: com.coralogix.catalog.v1.SpanKind
	(*v1.RepeatedStats)(nil),                  // 6: com.coralogixapis.service_catalog.v1.RepeatedStats
	(*v1.RepeatedLineChart)(nil),              // 7: com.coralogixapis.service_catalog.v1.RepeatedLineChart
}
var file_com_coralogix_catalog_v1_span_drill_down_service_proto_depIdxs = []int32{
	2,  // 0: com.coralogix.catalog.v1.GetServiceSideModelGraphsRequest.service_name:type_name -> google.protobuf.StringValue
	2,  // 1: com.coralogix.catalog.v1.GetServiceSideModelGraphsRequest.operation:type_name -> google.protobuf.StringValue
	3,  // 2: com.coralogix.catalog.v1.GetServiceSideModelGraphsRequest.time_range:type_name -> com.coralogix.catalog.v1.TimeRange
	4,  // 3: com.coralogix.catalog.v1.GetServiceSideModelGraphsRequest.service_catalog_filters:type_name -> com.coralogix.catalog.v1.ApmFilter
	5,  // 4: com.coralogix.catalog.v1.GetServiceSideModelGraphsRequest.span_kind:type_name -> com.coralogix.catalog.v1.SpanKind
	2,  // 5: com.coralogix.catalog.v1.GetServiceSideModelGraphsRequest.method:type_name -> google.protobuf.StringValue
	6,  // 6: com.coralogix.catalog.v1.GetServiceSideModelGraphsResponse.span_drill_down_side_modal_stats:type_name -> com.coralogixapis.service_catalog.v1.RepeatedStats
	7,  // 7: com.coralogix.catalog.v1.GetServiceSideModelGraphsResponse.requests_and_errors:type_name -> com.coralogixapis.service_catalog.v1.RepeatedLineChart
	7,  // 8: com.coralogix.catalog.v1.GetServiceSideModelGraphsResponse.latency_percentiles:type_name -> com.coralogixapis.service_catalog.v1.RepeatedLineChart
	7,  // 9: com.coralogix.catalog.v1.GetServiceSideModelGraphsResponse.requests_and_errors_per_second:type_name -> com.coralogixapis.service_catalog.v1.RepeatedLineChart
	7,  // 10: com.coralogix.catalog.v1.GetServiceSideModelGraphsResponse.latency_percentiles_per_second:type_name -> com.coralogixapis.service_catalog.v1.RepeatedLineChart
	0,  // 11: com.coralogix.catalog.v1.SpanDrillDownService.GetServiceSideModelGraphs:input_type -> com.coralogix.catalog.v1.GetServiceSideModelGraphsRequest
	1,  // 12: com.coralogix.catalog.v1.SpanDrillDownService.GetServiceSideModelGraphs:output_type -> com.coralogix.catalog.v1.GetServiceSideModelGraphsResponse
	12, // [12:13] is the sub-list for method output_type
	11, // [11:12] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_com_coralogix_catalog_v1_span_drill_down_service_proto_init() }
func file_com_coralogix_catalog_v1_span_drill_down_service_proto_init() {
	if File_com_coralogix_catalog_v1_span_drill_down_service_proto != nil {
		return
	}
	file_com_coralogix_catalog_v1_common_proto_init()
	file_com_coralogix_catalog_v1_span_kind_proto_init()
	file_com_coralogix_catalog_v1_span_drill_down_service_proto_msgTypes[1].OneofWrappers = []any{
		(*GetServiceSideModelGraphsResponse_SpanDrillDownSideModalStats)(nil),
		(*GetServiceSideModelGraphsResponse_RequestsAndErrors)(nil),
		(*GetServiceSideModelGraphsResponse_LatencyPercentiles)(nil),
		(*GetServiceSideModelGraphsResponse_RequestsAndErrorsPerSecond)(nil),
		(*GetServiceSideModelGraphsResponse_LatencyPercentilesPerSecond)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_catalog_v1_span_drill_down_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogix_catalog_v1_span_drill_down_service_proto_goTypes,
		DependencyIndexes: file_com_coralogix_catalog_v1_span_drill_down_service_proto_depIdxs,
		MessageInfos:      file_com_coralogix_catalog_v1_span_drill_down_service_proto_msgTypes,
	}.Build()
	File_com_coralogix_catalog_v1_span_drill_down_service_proto = out.File
	file_com_coralogix_catalog_v1_span_drill_down_service_proto_rawDesc = nil
	file_com_coralogix_catalog_v1_span_drill_down_service_proto_goTypes = nil
	file_com_coralogix_catalog_v1_span_drill_down_service_proto_depIdxs = nil
}
