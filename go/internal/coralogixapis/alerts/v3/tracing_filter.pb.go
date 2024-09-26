// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.28.1
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/tracing/common/tracing_filter.proto

package v3

import (
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

type TracingFilterOperationType int32

const (
	TracingFilterOperationType_TRACING_FILTER_OPERATION_TYPE_IS_OR_UNSPECIFIED TracingFilterOperationType = 0
	TracingFilterOperationType_TRACING_FILTER_OPERATION_TYPE_INCLUDES          TracingFilterOperationType = 1
	TracingFilterOperationType_TRACING_FILTER_OPERATION_TYPE_ENDS_WITH         TracingFilterOperationType = 2
	TracingFilterOperationType_TRACING_FILTER_OPERATION_TYPE_STARTS_WITH       TracingFilterOperationType = 3
	TracingFilterOperationType_TRACING_FILTER_OPERATION_TYPE_IS_NOT            TracingFilterOperationType = 4
)

// Enum value maps for TracingFilterOperationType.
var (
	TracingFilterOperationType_name = map[int32]string{
		0: "TRACING_FILTER_OPERATION_TYPE_IS_OR_UNSPECIFIED",
		1: "TRACING_FILTER_OPERATION_TYPE_INCLUDES",
		2: "TRACING_FILTER_OPERATION_TYPE_ENDS_WITH",
		3: "TRACING_FILTER_OPERATION_TYPE_STARTS_WITH",
		4: "TRACING_FILTER_OPERATION_TYPE_IS_NOT",
	}
	TracingFilterOperationType_value = map[string]int32{
		"TRACING_FILTER_OPERATION_TYPE_IS_OR_UNSPECIFIED": 0,
		"TRACING_FILTER_OPERATION_TYPE_INCLUDES":          1,
		"TRACING_FILTER_OPERATION_TYPE_ENDS_WITH":         2,
		"TRACING_FILTER_OPERATION_TYPE_STARTS_WITH":       3,
		"TRACING_FILTER_OPERATION_TYPE_IS_NOT":            4,
	}
)

func (x TracingFilterOperationType) Enum() *TracingFilterOperationType {
	p := new(TracingFilterOperationType)
	*p = x
	return p
}

func (x TracingFilterOperationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TracingFilterOperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_enumTypes[0].Descriptor()
}

func (TracingFilterOperationType) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_enumTypes[0]
}

func (x TracingFilterOperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TracingFilterOperationType.Descriptor instead.
func (TracingFilterOperationType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_rawDescGZIP(), []int{0}
}

type TracingFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to FilterType:
	//
	//	*TracingFilter_SimpleFilter
	FilterType isTracingFilter_FilterType `protobuf_oneof:"filter_type"`
}

func (x *TracingFilter) Reset() {
	*x = TracingFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TracingFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TracingFilter) ProtoMessage() {}

func (x *TracingFilter) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TracingFilter.ProtoReflect.Descriptor instead.
func (*TracingFilter) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_rawDescGZIP(), []int{0}
}

func (m *TracingFilter) GetFilterType() isTracingFilter_FilterType {
	if m != nil {
		return m.FilterType
	}
	return nil
}

func (x *TracingFilter) GetSimpleFilter() *TracingSimpleFilter {
	if x, ok := x.GetFilterType().(*TracingFilter_SimpleFilter); ok {
		return x.SimpleFilter
	}
	return nil
}

type isTracingFilter_FilterType interface {
	isTracingFilter_FilterType()
}

type TracingFilter_SimpleFilter struct {
	SimpleFilter *TracingSimpleFilter `protobuf:"bytes,1,opt,name=simple_filter,json=simpleFilter,proto3,oneof"`
}

func (*TracingFilter_SimpleFilter) isTracingFilter_FilterType() {}

type TracingSimpleFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TracingLabelFilters *TracingLabelFilters    `protobuf:"bytes,1,opt,name=tracing_label_filters,json=tracingLabelFilters,proto3" json:"tracing_label_filters,omitempty"`
	LatencyThresholdMs  *wrapperspb.UInt32Value `protobuf:"bytes,2,opt,name=latency_threshold_ms,json=latencyThresholdMs,proto3" json:"latency_threshold_ms,omitempty"`
}

func (x *TracingSimpleFilter) Reset() {
	*x = TracingSimpleFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TracingSimpleFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TracingSimpleFilter) ProtoMessage() {}

func (x *TracingSimpleFilter) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TracingSimpleFilter.ProtoReflect.Descriptor instead.
func (*TracingSimpleFilter) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_rawDescGZIP(), []int{1}
}

func (x *TracingSimpleFilter) GetTracingLabelFilters() *TracingLabelFilters {
	if x != nil {
		return x.TracingLabelFilters
	}
	return nil
}

func (x *TracingSimpleFilter) GetLatencyThresholdMs() *wrapperspb.UInt32Value {
	if x != nil {
		return x.LatencyThresholdMs
	}
	return nil
}

type TracingLabelFilters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplicationName []*TracingFilterType           `protobuf:"bytes,1,rep,name=application_name,json=applicationName,proto3" json:"application_name,omitempty"`
	SubsystemName   []*TracingFilterType           `protobuf:"bytes,2,rep,name=subsystem_name,json=subsystemName,proto3" json:"subsystem_name,omitempty"`
	ServiceName     []*TracingFilterType           `protobuf:"bytes,3,rep,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	OperationName   []*TracingFilterType           `protobuf:"bytes,4,rep,name=operation_name,json=operationName,proto3" json:"operation_name,omitempty"`
	SpanFields      []*TracingSpanFieldsFilterType `protobuf:"bytes,5,rep,name=span_fields,json=spanFields,proto3" json:"span_fields,omitempty"`
}

func (x *TracingLabelFilters) Reset() {
	*x = TracingLabelFilters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TracingLabelFilters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TracingLabelFilters) ProtoMessage() {}

func (x *TracingLabelFilters) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TracingLabelFilters.ProtoReflect.Descriptor instead.
func (*TracingLabelFilters) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_rawDescGZIP(), []int{2}
}

func (x *TracingLabelFilters) GetApplicationName() []*TracingFilterType {
	if x != nil {
		return x.ApplicationName
	}
	return nil
}

func (x *TracingLabelFilters) GetSubsystemName() []*TracingFilterType {
	if x != nil {
		return x.SubsystemName
	}
	return nil
}

func (x *TracingLabelFilters) GetServiceName() []*TracingFilterType {
	if x != nil {
		return x.ServiceName
	}
	return nil
}

func (x *TracingLabelFilters) GetOperationName() []*TracingFilterType {
	if x != nil {
		return x.OperationName
	}
	return nil
}

func (x *TracingLabelFilters) GetSpanFields() []*TracingSpanFieldsFilterType {
	if x != nil {
		return x.SpanFields
	}
	return nil
}

type TracingSpanFieldsFilterType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key        *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	FilterType *TracingFilterType      `protobuf:"bytes,2,opt,name=filter_type,json=filterType,proto3" json:"filter_type,omitempty"`
}

func (x *TracingSpanFieldsFilterType) Reset() {
	*x = TracingSpanFieldsFilterType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TracingSpanFieldsFilterType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TracingSpanFieldsFilterType) ProtoMessage() {}

func (x *TracingSpanFieldsFilterType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TracingSpanFieldsFilterType.ProtoReflect.Descriptor instead.
func (*TracingSpanFieldsFilterType) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_rawDescGZIP(), []int{3}
}

func (x *TracingSpanFieldsFilterType) GetKey() *wrapperspb.StringValue {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *TracingSpanFieldsFilterType) GetFilterType() *TracingFilterType {
	if x != nil {
		return x.FilterType
	}
	return nil
}

type TracingFilterType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values    []*wrapperspb.StringValue  `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	Operation TracingFilterOperationType `protobuf:"varint,2,opt,name=operation,proto3,enum=com.coralogixapis.alerts.v3.TracingFilterOperationType" json:"operation,omitempty"`
}

func (x *TracingFilterType) Reset() {
	*x = TracingFilterType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TracingFilterType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TracingFilterType) ProtoMessage() {}

func (x *TracingFilterType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TracingFilterType.ProtoReflect.Descriptor instead.
func (*TracingFilterType) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_rawDescGZIP(), []int{4}
}

func (x *TracingFilterType) GetValues() []*wrapperspb.StringValue {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *TracingFilterType) GetOperation() TracingFilterOperationType {
	if x != nil {
		return x.Operation
	}
	return TracingFilterOperationType_TRACING_FILTER_OPERATION_TYPE_IS_OR_UNSPECIFIED
}

var File_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_rawDesc = []byte{
	0x0a, 0x59, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x77, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x63,
	0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x57, 0x0a, 0x0d, 0x73, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x54,
	0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x42, 0x0d, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x22, 0xcb, 0x01, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x64, 0x0a, 0x15, 0x74, 0x72, 0x61,
	0x63, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x13, 0x74, 0x72, 0x61, 0x63,
	0x69, 0x6e, 0x67, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x4e, 0x0a, 0x14, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73,
	0x68, 0x6f, 0x6c, 0x64, 0x5f, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x12, 0x6c, 0x61, 0x74,
	0x65, 0x6e, 0x63, 0x79, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x4d, 0x73, 0x22,
	0xcc, 0x03, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x59, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e,
	0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x55, 0x0a, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x51, 0x0a, 0x0c, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x72,
	0x61, 0x63, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x55, 0x0a, 0x0e,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e,
	0x76, 0x33, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x59, 0x0a, 0x0b, 0x73, 0x70, 0x61, 0x6e, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x53, 0x70,
	0x61, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0a, 0x73, 0x70, 0x61, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x9e,
	0x01, 0x0a, 0x1b, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x61, 0x6e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x4f,
	0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76,
	0x33, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22,
	0xa0, 0x01, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x55, 0x0a, 0x09, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x37,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x72, 0x61,
	0x63, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2a, 0x83, 0x02, 0x0a, 0x1a, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x33, 0x0a, 0x2f, 0x54, 0x52, 0x41, 0x43, 0x49, 0x4e, 0x47, 0x5f, 0x46, 0x49, 0x4c,
	0x54, 0x45, 0x52, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x49, 0x53, 0x5f, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x2a, 0x0a, 0x26, 0x54, 0x52, 0x41, 0x43, 0x49, 0x4e,
	0x47, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x43, 0x4c, 0x55, 0x44, 0x45, 0x53,
	0x10, 0x01, 0x12, 0x2b, 0x0a, 0x27, 0x54, 0x52, 0x41, 0x43, 0x49, 0x4e, 0x47, 0x5f, 0x46, 0x49,
	0x4c, 0x54, 0x45, 0x52, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x45, 0x4e, 0x44, 0x53, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x10, 0x02, 0x12,
	0x2d, 0x0a, 0x29, 0x54, 0x52, 0x41, 0x43, 0x49, 0x4e, 0x47, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45,
	0x52, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x53, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x10, 0x03, 0x12, 0x28,
	0x0a, 0x24, 0x54, 0x52, 0x41, 0x43, 0x49, 0x4e, 0x47, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52,
	0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x49, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x10, 0x04, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_rawDescData = file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_goTypes = []interface{}{
	(TracingFilterOperationType)(0),     // 0: com.coralogixapis.alerts.v3.TracingFilterOperationType
	(*TracingFilter)(nil),               // 1: com.coralogixapis.alerts.v3.TracingFilter
	(*TracingSimpleFilter)(nil),         // 2: com.coralogixapis.alerts.v3.TracingSimpleFilter
	(*TracingLabelFilters)(nil),         // 3: com.coralogixapis.alerts.v3.TracingLabelFilters
	(*TracingSpanFieldsFilterType)(nil), // 4: com.coralogixapis.alerts.v3.TracingSpanFieldsFilterType
	(*TracingFilterType)(nil),           // 5: com.coralogixapis.alerts.v3.TracingFilterType
	(*wrapperspb.UInt32Value)(nil),      // 6: google.protobuf.UInt32Value
	(*wrapperspb.StringValue)(nil),      // 7: google.protobuf.StringValue
}
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_depIdxs = []int32{
	2,  // 0: com.coralogixapis.alerts.v3.TracingFilter.simple_filter:type_name -> com.coralogixapis.alerts.v3.TracingSimpleFilter
	3,  // 1: com.coralogixapis.alerts.v3.TracingSimpleFilter.tracing_label_filters:type_name -> com.coralogixapis.alerts.v3.TracingLabelFilters
	6,  // 2: com.coralogixapis.alerts.v3.TracingSimpleFilter.latency_threshold_ms:type_name -> google.protobuf.UInt32Value
	5,  // 3: com.coralogixapis.alerts.v3.TracingLabelFilters.application_name:type_name -> com.coralogixapis.alerts.v3.TracingFilterType
	5,  // 4: com.coralogixapis.alerts.v3.TracingLabelFilters.subsystem_name:type_name -> com.coralogixapis.alerts.v3.TracingFilterType
	5,  // 5: com.coralogixapis.alerts.v3.TracingLabelFilters.service_name:type_name -> com.coralogixapis.alerts.v3.TracingFilterType
	5,  // 6: com.coralogixapis.alerts.v3.TracingLabelFilters.operation_name:type_name -> com.coralogixapis.alerts.v3.TracingFilterType
	4,  // 7: com.coralogixapis.alerts.v3.TracingLabelFilters.span_fields:type_name -> com.coralogixapis.alerts.v3.TracingSpanFieldsFilterType
	7,  // 8: com.coralogixapis.alerts.v3.TracingSpanFieldsFilterType.key:type_name -> google.protobuf.StringValue
	5,  // 9: com.coralogixapis.alerts.v3.TracingSpanFieldsFilterType.filter_type:type_name -> com.coralogixapis.alerts.v3.TracingFilterType
	7,  // 10: com.coralogixapis.alerts.v3.TracingFilterType.values:type_name -> google.protobuf.StringValue
	0,  // 11: com.coralogixapis.alerts.v3.TracingFilterType.operation:type_name -> com.coralogixapis.alerts.v3.TracingFilterOperationType
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_init()
}
func file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TracingFilter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TracingSimpleFilter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TracingLabelFilters); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TracingSpanFieldsFilterType); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TracingFilterType); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TracingFilter_SimpleFilter)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_depIdxs = nil
}
