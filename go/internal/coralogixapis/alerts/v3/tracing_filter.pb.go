// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/tracing/common/tracing_filter.proto

package v3

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
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to FilterType:
	//
	//	*TracingFilter_SimpleFilter
	FilterType    isTracingFilter_FilterType `protobuf_oneof:"filter_type"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TracingFilter) Reset() {
	*x = TracingFilter{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TracingFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TracingFilter) ProtoMessage() {}

func (x *TracingFilter) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_msgTypes[0]
	if x != nil {
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

func (x *TracingFilter) GetFilterType() isTracingFilter_FilterType {
	if x != nil {
		return x.FilterType
	}
	return nil
}

func (x *TracingFilter) GetSimpleFilter() *TracingSimpleFilter {
	if x != nil {
		if x, ok := x.FilterType.(*TracingFilter_SimpleFilter); ok {
			return x.SimpleFilter
		}
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
	state               protoimpl.MessageState  `protogen:"open.v1"`
	TracingLabelFilters *TracingLabelFilters    `protobuf:"bytes,1,opt,name=tracing_label_filters,json=tracingLabelFilters,proto3" json:"tracing_label_filters,omitempty"`
	LatencyThresholdMs  *wrapperspb.UInt64Value `protobuf:"bytes,2,opt,name=latency_threshold_ms,json=latencyThresholdMs,proto3" json:"latency_threshold_ms,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *TracingSimpleFilter) Reset() {
	*x = TracingSimpleFilter{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TracingSimpleFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TracingSimpleFilter) ProtoMessage() {}

func (x *TracingSimpleFilter) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_msgTypes[1]
	if x != nil {
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

func (x *TracingSimpleFilter) GetLatencyThresholdMs() *wrapperspb.UInt64Value {
	if x != nil {
		return x.LatencyThresholdMs
	}
	return nil
}

type TracingLabelFilters struct {
	state           protoimpl.MessageState         `protogen:"open.v1"`
	ApplicationName []*TracingFilterType           `protobuf:"bytes,1,rep,name=application_name,json=applicationName,proto3" json:"application_name,omitempty"`
	SubsystemName   []*TracingFilterType           `protobuf:"bytes,2,rep,name=subsystem_name,json=subsystemName,proto3" json:"subsystem_name,omitempty"`
	ServiceName     []*TracingFilterType           `protobuf:"bytes,3,rep,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	OperationName   []*TracingFilterType           `protobuf:"bytes,4,rep,name=operation_name,json=operationName,proto3" json:"operation_name,omitempty"`
	SpanFields      []*TracingSpanFieldsFilterType `protobuf:"bytes,5,rep,name=span_fields,json=spanFields,proto3" json:"span_fields,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *TracingLabelFilters) Reset() {
	*x = TracingLabelFilters{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TracingLabelFilters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TracingLabelFilters) ProtoMessage() {}

func (x *TracingLabelFilters) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_msgTypes[2]
	if x != nil {
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
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Key           *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	FilterType    *TracingFilterType      `protobuf:"bytes,2,opt,name=filter_type,json=filterType,proto3" json:"filter_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TracingSpanFieldsFilterType) Reset() {
	*x = TracingSpanFieldsFilterType{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TracingSpanFieldsFilterType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TracingSpanFieldsFilterType) ProtoMessage() {}

func (x *TracingSpanFieldsFilterType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_msgTypes[3]
	if x != nil {
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
	state         protoimpl.MessageState     `protogen:"open.v1"`
	Values        []*wrapperspb.StringValue  `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	Operation     TracingFilterOperationType `protobuf:"varint,2,opt,name=operation,proto3,enum=com.coralogixapis.alerts.v3.TracingFilterOperationType" json:"operation,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TracingFilterType) Reset() {
	*x = TracingFilterType{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TracingFilterType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TracingFilterType) ProtoMessage() {}

func (x *TracingFilterType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_msgTypes[4]
	if x != nil {
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

const file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_rawDesc = "" +
	"\n" +
	"Ycom/coralogixapis/alerts/v3/alert_def_type_definition/tracing/common/tracing_filter.proto\x12\x1bcom.coralogixapis.alerts.v3\x1a\x1egoogle/protobuf/wrappers.proto\"w\n" +
	"\rTracingFilter\x12W\n" +
	"\rsimple_filter\x18\x01 \x01(\v20.com.coralogixapis.alerts.v3.TracingSimpleFilterH\x00R\fsimpleFilterB\r\n" +
	"\vfilter_type\"\xcb\x01\n" +
	"\x13TracingSimpleFilter\x12d\n" +
	"\x15tracing_label_filters\x18\x01 \x01(\v20.com.coralogixapis.alerts.v3.TracingLabelFiltersR\x13tracingLabelFilters\x12N\n" +
	"\x14latency_threshold_ms\x18\x02 \x01(\v2\x1c.google.protobuf.UInt64ValueR\x12latencyThresholdMs\"\xcc\x03\n" +
	"\x13TracingLabelFilters\x12Y\n" +
	"\x10application_name\x18\x01 \x03(\v2..com.coralogixapis.alerts.v3.TracingFilterTypeR\x0fapplicationName\x12U\n" +
	"\x0esubsystem_name\x18\x02 \x03(\v2..com.coralogixapis.alerts.v3.TracingFilterTypeR\rsubsystemName\x12Q\n" +
	"\fservice_name\x18\x03 \x03(\v2..com.coralogixapis.alerts.v3.TracingFilterTypeR\vserviceName\x12U\n" +
	"\x0eoperation_name\x18\x04 \x03(\v2..com.coralogixapis.alerts.v3.TracingFilterTypeR\roperationName\x12Y\n" +
	"\vspan_fields\x18\x05 \x03(\v28.com.coralogixapis.alerts.v3.TracingSpanFieldsFilterTypeR\n" +
	"spanFields\"\x9e\x01\n" +
	"\x1bTracingSpanFieldsFilterType\x12.\n" +
	"\x03key\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueR\x03key\x12O\n" +
	"\vfilter_type\x18\x02 \x01(\v2..com.coralogixapis.alerts.v3.TracingFilterTypeR\n" +
	"filterType\"\xa0\x01\n" +
	"\x11TracingFilterType\x124\n" +
	"\x06values\x18\x01 \x03(\v2\x1c.google.protobuf.StringValueR\x06values\x12U\n" +
	"\toperation\x18\x02 \x01(\x0e27.com.coralogixapis.alerts.v3.TracingFilterOperationTypeR\toperation*\x83\x02\n" +
	"\x1aTracingFilterOperationType\x123\n" +
	"/TRACING_FILTER_OPERATION_TYPE_IS_OR_UNSPECIFIED\x10\x00\x12*\n" +
	"&TRACING_FILTER_OPERATION_TYPE_INCLUDES\x10\x01\x12+\n" +
	"'TRACING_FILTER_OPERATION_TYPE_ENDS_WITH\x10\x02\x12-\n" +
	")TRACING_FILTER_OPERATION_TYPE_STARTS_WITH\x10\x03\x12(\n" +
	"$TRACING_FILTER_OPERATION_TYPE_IS_NOT\x10\x04b\x06proto3"

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_goTypes = []any{
	(TracingFilterOperationType)(0),     // 0: com.coralogixapis.alerts.v3.TracingFilterOperationType
	(*TracingFilter)(nil),               // 1: com.coralogixapis.alerts.v3.TracingFilter
	(*TracingSimpleFilter)(nil),         // 2: com.coralogixapis.alerts.v3.TracingSimpleFilter
	(*TracingLabelFilters)(nil),         // 3: com.coralogixapis.alerts.v3.TracingLabelFilters
	(*TracingSpanFieldsFilterType)(nil), // 4: com.coralogixapis.alerts.v3.TracingSpanFieldsFilterType
	(*TracingFilterType)(nil),           // 5: com.coralogixapis.alerts.v3.TracingFilterType
	(*wrapperspb.UInt64Value)(nil),      // 6: google.protobuf.UInt64Value
	(*wrapperspb.StringValue)(nil),      // 7: google.protobuf.StringValue
}
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_depIdxs = []int32{
	2,  // 0: com.coralogixapis.alerts.v3.TracingFilter.simple_filter:type_name -> com.coralogixapis.alerts.v3.TracingSimpleFilter
	3,  // 1: com.coralogixapis.alerts.v3.TracingSimpleFilter.tracing_label_filters:type_name -> com.coralogixapis.alerts.v3.TracingLabelFilters
	6,  // 2: com.coralogixapis.alerts.v3.TracingSimpleFilter.latency_threshold_ms:type_name -> google.protobuf.UInt64Value
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
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_msgTypes[0].OneofWrappers = []any{
		(*TracingFilter_SimpleFilter)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_rawDesc)),
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
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_common_tracing_filter_proto_depIdxs = nil
}
