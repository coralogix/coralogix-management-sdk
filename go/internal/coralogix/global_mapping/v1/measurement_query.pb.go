// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: com/coralogix/global_mapping/v1/measurement_query.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

type AggregationType int32

const (
	AggregationType_AGGREGATION_TYPE_UNSPECIFIED AggregationType = 0
	AggregationType_AGGREGATION_TYPE_AVG         AggregationType = 1
	AggregationType_AGGREGATION_TYPE_MIN         AggregationType = 2
	AggregationType_AGGREGATION_TYPE_MAX         AggregationType = 3
	AggregationType_AGGREGATION_TYPE_SUM         AggregationType = 4
	AggregationType_AGGREGATION_TYPE_COUNT       AggregationType = 5
)

// Enum value maps for AggregationType.
var (
	AggregationType_name = map[int32]string{
		0: "AGGREGATION_TYPE_UNSPECIFIED",
		1: "AGGREGATION_TYPE_AVG",
		2: "AGGREGATION_TYPE_MIN",
		3: "AGGREGATION_TYPE_MAX",
		4: "AGGREGATION_TYPE_SUM",
		5: "AGGREGATION_TYPE_COUNT",
	}
	AggregationType_value = map[string]int32{
		"AGGREGATION_TYPE_UNSPECIFIED": 0,
		"AGGREGATION_TYPE_AVG":         1,
		"AGGREGATION_TYPE_MIN":         2,
		"AGGREGATION_TYPE_MAX":         3,
		"AGGREGATION_TYPE_SUM":         4,
		"AGGREGATION_TYPE_COUNT":       5,
	}
)

func (x AggregationType) Enum() *AggregationType {
	p := new(AggregationType)
	*p = x
	return p
}

func (x AggregationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AggregationType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_global_mapping_v1_measurement_query_proto_enumTypes[0].Descriptor()
}

func (AggregationType) Type() protoreflect.EnumType {
	return &file_com_coralogix_global_mapping_v1_measurement_query_proto_enumTypes[0]
}

func (x AggregationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AggregationType.Descriptor instead.
func (AggregationType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_global_mapping_v1_measurement_query_proto_rawDescGZIP(), []int{0}
}

type QueryType int32

const (
	QueryType_QUERY_TYPE_UNSPECIFIED QueryType = 0
	QueryType_QUERY_TYPE_RANGE       QueryType = 1
	QueryType_QUERY_TYPE_INSTANT     QueryType = 2
)

// Enum value maps for QueryType.
var (
	QueryType_name = map[int32]string{
		0: "QUERY_TYPE_UNSPECIFIED",
		1: "QUERY_TYPE_RANGE",
		2: "QUERY_TYPE_INSTANT",
	}
	QueryType_value = map[string]int32{
		"QUERY_TYPE_UNSPECIFIED": 0,
		"QUERY_TYPE_RANGE":       1,
		"QUERY_TYPE_INSTANT":     2,
	}
)

func (x QueryType) Enum() *QueryType {
	p := new(QueryType)
	*p = x
	return p
}

func (x QueryType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QueryType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_global_mapping_v1_measurement_query_proto_enumTypes[1].Descriptor()
}

func (QueryType) Type() protoreflect.EnumType {
	return &file_com_coralogix_global_mapping_v1_measurement_query_proto_enumTypes[1]
}

func (x QueryType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QueryType.Descriptor instead.
func (QueryType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_global_mapping_v1_measurement_query_proto_rawDescGZIP(), []int{1}
}

type TopkAggregationType int32

const (
	TopkAggregationType_TOPK_AGGREGATION_TYPE_UNSPECIFIED TopkAggregationType = 0
	TopkAggregationType_TOPK_AGGREGATION_TYPE_LAST        TopkAggregationType = 1
	TopkAggregationType_TOPK_AGGREGATION_TYPE_MAX         TopkAggregationType = 2
	TopkAggregationType_TOPK_AGGREGATION_TYPE_MIN         TopkAggregationType = 3
	TopkAggregationType_TOPK_AGGREGATION_TYPE_AVG         TopkAggregationType = 4
)

// Enum value maps for TopkAggregationType.
var (
	TopkAggregationType_name = map[int32]string{
		0: "TOPK_AGGREGATION_TYPE_UNSPECIFIED",
		1: "TOPK_AGGREGATION_TYPE_LAST",
		2: "TOPK_AGGREGATION_TYPE_MAX",
		3: "TOPK_AGGREGATION_TYPE_MIN",
		4: "TOPK_AGGREGATION_TYPE_AVG",
	}
	TopkAggregationType_value = map[string]int32{
		"TOPK_AGGREGATION_TYPE_UNSPECIFIED": 0,
		"TOPK_AGGREGATION_TYPE_LAST":        1,
		"TOPK_AGGREGATION_TYPE_MAX":         2,
		"TOPK_AGGREGATION_TYPE_MIN":         3,
		"TOPK_AGGREGATION_TYPE_AVG":         4,
	}
)

func (x TopkAggregationType) Enum() *TopkAggregationType {
	p := new(TopkAggregationType)
	*p = x
	return p
}

func (x TopkAggregationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TopkAggregationType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_global_mapping_v1_measurement_query_proto_enumTypes[2].Descriptor()
}

func (TopkAggregationType) Type() protoreflect.EnumType {
	return &file_com_coralogix_global_mapping_v1_measurement_query_proto_enumTypes[2]
}

func (x TopkAggregationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TopkAggregationType.Descriptor instead.
func (TopkAggregationType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_global_mapping_v1_measurement_query_proto_rawDescGZIP(), []int{2}
}

type QueryResult struct {
	state         protoimpl.MessageState             `protogen:"open.v1"`
	Metric        map[string]*wrapperspb.StringValue `protobuf:"bytes,1,rep,name=metric,proto3" json:"metric,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Values        []*structpb.ListValue              `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *QueryResult) Reset() {
	*x = QueryResult{}
	mi := &file_com_coralogix_global_mapping_v1_measurement_query_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResult) ProtoMessage() {}

func (x *QueryResult) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_global_mapping_v1_measurement_query_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResult.ProtoReflect.Descriptor instead.
func (*QueryResult) Descriptor() ([]byte, []int) {
	return file_com_coralogix_global_mapping_v1_measurement_query_proto_rawDescGZIP(), []int{0}
}

func (x *QueryResult) GetMetric() map[string]*wrapperspb.StringValue {
	if x != nil {
		return x.Metric
	}
	return nil
}

func (x *QueryResult) GetValues() []*structpb.ListValue {
	if x != nil {
		return x.Values
	}
	return nil
}

type TopkAggregation struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Aggregation   TopkAggregationType     `protobuf:"varint,1,opt,name=aggregation,proto3,enum=com.coralogix.global_mapping.v1.TopkAggregationType" json:"aggregation,omitempty"`
	Topk          *wrapperspb.UInt32Value `protobuf:"bytes,2,opt,name=topk,proto3" json:"topk,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TopkAggregation) Reset() {
	*x = TopkAggregation{}
	mi := &file_com_coralogix_global_mapping_v1_measurement_query_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TopkAggregation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopkAggregation) ProtoMessage() {}

func (x *TopkAggregation) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_global_mapping_v1_measurement_query_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopkAggregation.ProtoReflect.Descriptor instead.
func (*TopkAggregation) Descriptor() ([]byte, []int) {
	return file_com_coralogix_global_mapping_v1_measurement_query_proto_rawDescGZIP(), []int{1}
}

func (x *TopkAggregation) GetAggregation() TopkAggregationType {
	if x != nil {
		return x.Aggregation
	}
	return TopkAggregationType_TOPK_AGGREGATION_TYPE_UNSPECIFIED
}

func (x *TopkAggregation) GetTopk() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Topk
	}
	return nil
}

type MeasurementMetadata struct {
	state                    protoimpl.MessageState  `protogen:"open.v1"`
	MeasurementName          *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=measurement_name,json=measurementName,proto3" json:"measurement_name,omitempty"`
	AggregationType          AggregationType         `protobuf:"varint,2,opt,name=aggregation_type,json=aggregationType,proto3,enum=com.coralogix.global_mapping.v1.AggregationType" json:"aggregation_type,omitempty"`
	QueryType                QueryType               `protobuf:"varint,3,opt,name=query_type,json=queryType,proto3,enum=com.coralogix.global_mapping.v1.QueryType" json:"query_type,omitempty"`
	Step                     *wrapperspb.Int32Value  `protobuf:"bytes,4,opt,name=step,proto3" json:"step,omitempty"`
	AggregationWindowSeconds *wrapperspb.Int32Value  `protobuf:"bytes,5,opt,name=aggregation_window_seconds,json=aggregationWindowSeconds,proto3" json:"aggregation_window_seconds,omitempty"`
	TopkAggregation          *TopkAggregation        `protobuf:"bytes,6,opt,name=topk_aggregation,json=topkAggregation,proto3" json:"topk_aggregation,omitempty"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *MeasurementMetadata) Reset() {
	*x = MeasurementMetadata{}
	mi := &file_com_coralogix_global_mapping_v1_measurement_query_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MeasurementMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeasurementMetadata) ProtoMessage() {}

func (x *MeasurementMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_global_mapping_v1_measurement_query_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeasurementMetadata.ProtoReflect.Descriptor instead.
func (*MeasurementMetadata) Descriptor() ([]byte, []int) {
	return file_com_coralogix_global_mapping_v1_measurement_query_proto_rawDescGZIP(), []int{2}
}

func (x *MeasurementMetadata) GetMeasurementName() *wrapperspb.StringValue {
	if x != nil {
		return x.MeasurementName
	}
	return nil
}

func (x *MeasurementMetadata) GetAggregationType() AggregationType {
	if x != nil {
		return x.AggregationType
	}
	return AggregationType_AGGREGATION_TYPE_UNSPECIFIED
}

func (x *MeasurementMetadata) GetQueryType() QueryType {
	if x != nil {
		return x.QueryType
	}
	return QueryType_QUERY_TYPE_UNSPECIFIED
}

func (x *MeasurementMetadata) GetStep() *wrapperspb.Int32Value {
	if x != nil {
		return x.Step
	}
	return nil
}

func (x *MeasurementMetadata) GetAggregationWindowSeconds() *wrapperspb.Int32Value {
	if x != nil {
		return x.AggregationWindowSeconds
	}
	return nil
}

func (x *MeasurementMetadata) GetTopkAggregation() *TopkAggregation {
	if x != nil {
		return x.TopkAggregation
	}
	return nil
}

type MeasurementQuery struct {
	state                    protoimpl.MessageState  `protogen:"open.v1"`
	Query                    *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Name                     *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description              *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	DisplayName              *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	LogicalGroup             *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=logical_group,json=logicalGroup,proto3" json:"logical_group,omitempty"`
	QueryResults             []*QueryResult          `protobuf:"bytes,6,rep,name=query_results,json=queryResults,proto3" json:"query_results,omitempty"`
	OrderingQuery            *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=ordering_query,json=orderingQuery,proto3" json:"ordering_query,omitempty"`
	AggregationType          AggregationType         `protobuf:"varint,8,opt,name=aggregation_type,json=aggregationType,proto3,enum=com.coralogix.global_mapping.v1.AggregationType" json:"aggregation_type,omitempty"`
	QueryType                QueryType               `protobuf:"varint,9,opt,name=query_type,json=queryType,proto3,enum=com.coralogix.global_mapping.v1.QueryType" json:"query_type,omitempty"`
	LabelValues              *LabelValues            `protobuf:"bytes,10,opt,name=label_values,json=labelValues,proto3" json:"label_values,omitempty"`
	Step                     *wrapperspb.Int32Value  `protobuf:"bytes,11,opt,name=step,proto3" json:"step,omitempty"`
	AggregationWindowSeconds *wrapperspb.Int32Value  `protobuf:"bytes,12,opt,name=aggregation_window_seconds,json=aggregationWindowSeconds,proto3" json:"aggregation_window_seconds,omitempty"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *MeasurementQuery) Reset() {
	*x = MeasurementQuery{}
	mi := &file_com_coralogix_global_mapping_v1_measurement_query_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MeasurementQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeasurementQuery) ProtoMessage() {}

func (x *MeasurementQuery) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_global_mapping_v1_measurement_query_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeasurementQuery.ProtoReflect.Descriptor instead.
func (*MeasurementQuery) Descriptor() ([]byte, []int) {
	return file_com_coralogix_global_mapping_v1_measurement_query_proto_rawDescGZIP(), []int{3}
}

func (x *MeasurementQuery) GetQuery() *wrapperspb.StringValue {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *MeasurementQuery) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *MeasurementQuery) GetDescription() *wrapperspb.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *MeasurementQuery) GetDisplayName() *wrapperspb.StringValue {
	if x != nil {
		return x.DisplayName
	}
	return nil
}

func (x *MeasurementQuery) GetLogicalGroup() *wrapperspb.StringValue {
	if x != nil {
		return x.LogicalGroup
	}
	return nil
}

func (x *MeasurementQuery) GetQueryResults() []*QueryResult {
	if x != nil {
		return x.QueryResults
	}
	return nil
}

func (x *MeasurementQuery) GetOrderingQuery() *wrapperspb.StringValue {
	if x != nil {
		return x.OrderingQuery
	}
	return nil
}

func (x *MeasurementQuery) GetAggregationType() AggregationType {
	if x != nil {
		return x.AggregationType
	}
	return AggregationType_AGGREGATION_TYPE_UNSPECIFIED
}

func (x *MeasurementQuery) GetQueryType() QueryType {
	if x != nil {
		return x.QueryType
	}
	return QueryType_QUERY_TYPE_UNSPECIFIED
}

func (x *MeasurementQuery) GetLabelValues() *LabelValues {
	if x != nil {
		return x.LabelValues
	}
	return nil
}

func (x *MeasurementQuery) GetStep() *wrapperspb.Int32Value {
	if x != nil {
		return x.Step
	}
	return nil
}

func (x *MeasurementQuery) GetAggregationWindowSeconds() *wrapperspb.Int32Value {
	if x != nil {
		return x.AggregationWindowSeconds
	}
	return nil
}

var File_com_coralogix_global_mapping_v1_measurement_query_proto protoreflect.FileDescriptor

var file_com_coralogix_global_mapping_v1_measurement_query_proto_rawDesc = []byte{
	0x0a, 0x37, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f,
	0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x01, 0x0a,
	0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x50, 0x0a, 0x06,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x32,
	0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x1a, 0x57, 0x0a, 0x0b, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9b, 0x01, 0x0a, 0x0f,
	0x54, 0x6f, 0x70, 0x6b, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x56, 0x0a, 0x0b, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70, 0x6b, 0x41, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x61, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x6f, 0x70, 0x6b, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x04, 0x74, 0x6f, 0x70, 0x6b, 0x22, 0xef, 0x03, 0x0a, 0x13, 0x4d, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x47, 0x0a, 0x10, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x6d, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x5b, 0x0a, 0x10, 0x61, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x73,
	0x74, 0x65, 0x70, 0x12, 0x59, 0x0a, 0x1a, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x18, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x5b,
	0x0a, 0x10, 0x74, 0x6f, 0x70, 0x6b, 0x5f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f,
	0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70, 0x6b, 0x41,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x74, 0x6f, 0x70, 0x6b,
	0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xd9, 0x06, 0x0a, 0x10,
	0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x32, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x30, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0d, 0x6c, 0x6f, 0x67, 0x69, 0x63,
	0x61, 0x6c, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x6c, 0x6f,
	0x67, 0x69, 0x63, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x51, 0x0a, 0x0d, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x0c, 0x71, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x43, 0x0a,
	0x0e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x12, 0x5b, 0x0a, 0x10, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x49, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x0b,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x04, 0x73,
	0x74, 0x65, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x12, 0x59, 0x0a, 0x1a,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x77, 0x69, 0x6e, 0x64,
	0x6f, 0x77, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x18, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77,
	0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x2a, 0xb7, 0x01, 0x0a, 0x0f, 0x41, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x1c, 0x41,
	0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a,
	0x14, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x41, 0x56, 0x47, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x47, 0x47, 0x52, 0x45,
	0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x49, 0x4e, 0x10,
	0x02, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x41, 0x58, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x41,
	0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x53, 0x55, 0x4d, 0x10, 0x04, 0x12, 0x1a, 0x0a, 0x16, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10,
	0x05, 0x2a, 0x55, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a,
	0x0a, 0x16, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x51, 0x55,
	0x45, 0x52, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x01,
	0x12, 0x16, 0x0a, 0x12, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49,
	0x4e, 0x53, 0x54, 0x41, 0x4e, 0x54, 0x10, 0x02, 0x2a, 0xb9, 0x01, 0x0a, 0x13, 0x54, 0x6f, 0x70,
	0x6b, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x25, 0x0a, 0x21, 0x54, 0x4f, 0x50, 0x4b, 0x5f, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x54, 0x4f, 0x50, 0x4b, 0x5f,
	0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4c, 0x41, 0x53, 0x54, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x54, 0x4f, 0x50, 0x4b, 0x5f,
	0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4d, 0x41, 0x58, 0x10, 0x02, 0x12, 0x1d, 0x0a, 0x19, 0x54, 0x4f, 0x50, 0x4b, 0x5f, 0x41,
	0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4d, 0x49, 0x4e, 0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19, 0x54, 0x4f, 0x50, 0x4b, 0x5f, 0x41, 0x47,
	0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41,
	0x56, 0x47, 0x10, 0x04, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_global_mapping_v1_measurement_query_proto_rawDescOnce sync.Once
	file_com_coralogix_global_mapping_v1_measurement_query_proto_rawDescData = file_com_coralogix_global_mapping_v1_measurement_query_proto_rawDesc
)

func file_com_coralogix_global_mapping_v1_measurement_query_proto_rawDescGZIP() []byte {
	file_com_coralogix_global_mapping_v1_measurement_query_proto_rawDescOnce.Do(func() {
		file_com_coralogix_global_mapping_v1_measurement_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_global_mapping_v1_measurement_query_proto_rawDescData)
	})
	return file_com_coralogix_global_mapping_v1_measurement_query_proto_rawDescData
}

var file_com_coralogix_global_mapping_v1_measurement_query_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_com_coralogix_global_mapping_v1_measurement_query_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_com_coralogix_global_mapping_v1_measurement_query_proto_goTypes = []any{
	(AggregationType)(0),           // 0: com.coralogix.global_mapping.v1.AggregationType
	(QueryType)(0),                 // 1: com.coralogix.global_mapping.v1.QueryType
	(TopkAggregationType)(0),       // 2: com.coralogix.global_mapping.v1.TopkAggregationType
	(*QueryResult)(nil),            // 3: com.coralogix.global_mapping.v1.QueryResult
	(*TopkAggregation)(nil),        // 4: com.coralogix.global_mapping.v1.TopkAggregation
	(*MeasurementMetadata)(nil),    // 5: com.coralogix.global_mapping.v1.MeasurementMetadata
	(*MeasurementQuery)(nil),       // 6: com.coralogix.global_mapping.v1.MeasurementQuery
	nil,                            // 7: com.coralogix.global_mapping.v1.QueryResult.MetricEntry
	(*structpb.ListValue)(nil),     // 8: google.protobuf.ListValue
	(*wrapperspb.UInt32Value)(nil), // 9: google.protobuf.UInt32Value
	(*wrapperspb.StringValue)(nil), // 10: google.protobuf.StringValue
	(*wrapperspb.Int32Value)(nil),  // 11: google.protobuf.Int32Value
	(*LabelValues)(nil),            // 12: com.coralogix.global_mapping.v1.LabelValues
}
var file_com_coralogix_global_mapping_v1_measurement_query_proto_depIdxs = []int32{
	7,  // 0: com.coralogix.global_mapping.v1.QueryResult.metric:type_name -> com.coralogix.global_mapping.v1.QueryResult.MetricEntry
	8,  // 1: com.coralogix.global_mapping.v1.QueryResult.values:type_name -> google.protobuf.ListValue
	2,  // 2: com.coralogix.global_mapping.v1.TopkAggregation.aggregation:type_name -> com.coralogix.global_mapping.v1.TopkAggregationType
	9,  // 3: com.coralogix.global_mapping.v1.TopkAggregation.topk:type_name -> google.protobuf.UInt32Value
	10, // 4: com.coralogix.global_mapping.v1.MeasurementMetadata.measurement_name:type_name -> google.protobuf.StringValue
	0,  // 5: com.coralogix.global_mapping.v1.MeasurementMetadata.aggregation_type:type_name -> com.coralogix.global_mapping.v1.AggregationType
	1,  // 6: com.coralogix.global_mapping.v1.MeasurementMetadata.query_type:type_name -> com.coralogix.global_mapping.v1.QueryType
	11, // 7: com.coralogix.global_mapping.v1.MeasurementMetadata.step:type_name -> google.protobuf.Int32Value
	11, // 8: com.coralogix.global_mapping.v1.MeasurementMetadata.aggregation_window_seconds:type_name -> google.protobuf.Int32Value
	4,  // 9: com.coralogix.global_mapping.v1.MeasurementMetadata.topk_aggregation:type_name -> com.coralogix.global_mapping.v1.TopkAggregation
	10, // 10: com.coralogix.global_mapping.v1.MeasurementQuery.query:type_name -> google.protobuf.StringValue
	10, // 11: com.coralogix.global_mapping.v1.MeasurementQuery.name:type_name -> google.protobuf.StringValue
	10, // 12: com.coralogix.global_mapping.v1.MeasurementQuery.description:type_name -> google.protobuf.StringValue
	10, // 13: com.coralogix.global_mapping.v1.MeasurementQuery.display_name:type_name -> google.protobuf.StringValue
	10, // 14: com.coralogix.global_mapping.v1.MeasurementQuery.logical_group:type_name -> google.protobuf.StringValue
	3,  // 15: com.coralogix.global_mapping.v1.MeasurementQuery.query_results:type_name -> com.coralogix.global_mapping.v1.QueryResult
	10, // 16: com.coralogix.global_mapping.v1.MeasurementQuery.ordering_query:type_name -> google.protobuf.StringValue
	0,  // 17: com.coralogix.global_mapping.v1.MeasurementQuery.aggregation_type:type_name -> com.coralogix.global_mapping.v1.AggregationType
	1,  // 18: com.coralogix.global_mapping.v1.MeasurementQuery.query_type:type_name -> com.coralogix.global_mapping.v1.QueryType
	12, // 19: com.coralogix.global_mapping.v1.MeasurementQuery.label_values:type_name -> com.coralogix.global_mapping.v1.LabelValues
	11, // 20: com.coralogix.global_mapping.v1.MeasurementQuery.step:type_name -> google.protobuf.Int32Value
	11, // 21: com.coralogix.global_mapping.v1.MeasurementQuery.aggregation_window_seconds:type_name -> google.protobuf.Int32Value
	10, // 22: com.coralogix.global_mapping.v1.QueryResult.MetricEntry.value:type_name -> google.protobuf.StringValue
	23, // [23:23] is the sub-list for method output_type
	23, // [23:23] is the sub-list for method input_type
	23, // [23:23] is the sub-list for extension type_name
	23, // [23:23] is the sub-list for extension extendee
	0,  // [0:23] is the sub-list for field type_name
}

func init() { file_com_coralogix_global_mapping_v1_measurement_query_proto_init() }
func file_com_coralogix_global_mapping_v1_measurement_query_proto_init() {
	if File_com_coralogix_global_mapping_v1_measurement_query_proto != nil {
		return
	}
	file_com_coralogix_global_mapping_v1_label_values_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_global_mapping_v1_measurement_query_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_global_mapping_v1_measurement_query_proto_goTypes,
		DependencyIndexes: file_com_coralogix_global_mapping_v1_measurement_query_proto_depIdxs,
		EnumInfos:         file_com_coralogix_global_mapping_v1_measurement_query_proto_enumTypes,
		MessageInfos:      file_com_coralogix_global_mapping_v1_measurement_query_proto_msgTypes,
	}.Build()
	File_com_coralogix_global_mapping_v1_measurement_query_proto = out.File
	file_com_coralogix_global_mapping_v1_measurement_query_proto_rawDesc = nil
	file_com_coralogix_global_mapping_v1_measurement_query_proto_goTypes = nil
	file_com_coralogix_global_mapping_v1_measurement_query_proto_depIdxs = nil
}
