// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: com/coralogixapis/dashboards/v1/common/logs_aggregation.proto

package v1

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

type LogsAggregation struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Value:
	//
	//	*LogsAggregation_Count_
	//	*LogsAggregation_CountDistinct_
	//	*LogsAggregation_Sum_
	//	*LogsAggregation_Average_
	//	*LogsAggregation_Min_
	//	*LogsAggregation_Max_
	//	*LogsAggregation_Percentile_
	Value         isLogsAggregation_Value `protobuf_oneof:"value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogsAggregation) Reset() {
	*x = LogsAggregation{}
	mi := &file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsAggregation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsAggregation) ProtoMessage() {}

func (x *LogsAggregation) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsAggregation.ProtoReflect.Descriptor instead.
func (*LogsAggregation) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_rawDescGZIP(), []int{0}
}

func (x *LogsAggregation) GetValue() isLogsAggregation_Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *LogsAggregation) GetCount() *LogsAggregation_Count {
	if x != nil {
		if x, ok := x.Value.(*LogsAggregation_Count_); ok {
			return x.Count
		}
	}
	return nil
}

func (x *LogsAggregation) GetCountDistinct() *LogsAggregation_CountDistinct {
	if x != nil {
		if x, ok := x.Value.(*LogsAggregation_CountDistinct_); ok {
			return x.CountDistinct
		}
	}
	return nil
}

func (x *LogsAggregation) GetSum() *LogsAggregation_Sum {
	if x != nil {
		if x, ok := x.Value.(*LogsAggregation_Sum_); ok {
			return x.Sum
		}
	}
	return nil
}

func (x *LogsAggregation) GetAverage() *LogsAggregation_Average {
	if x != nil {
		if x, ok := x.Value.(*LogsAggregation_Average_); ok {
			return x.Average
		}
	}
	return nil
}

func (x *LogsAggregation) GetMin() *LogsAggregation_Min {
	if x != nil {
		if x, ok := x.Value.(*LogsAggregation_Min_); ok {
			return x.Min
		}
	}
	return nil
}

func (x *LogsAggregation) GetMax() *LogsAggregation_Max {
	if x != nil {
		if x, ok := x.Value.(*LogsAggregation_Max_); ok {
			return x.Max
		}
	}
	return nil
}

func (x *LogsAggregation) GetPercentile() *LogsAggregation_Percentile {
	if x != nil {
		if x, ok := x.Value.(*LogsAggregation_Percentile_); ok {
			return x.Percentile
		}
	}
	return nil
}

type isLogsAggregation_Value interface {
	isLogsAggregation_Value()
}

type LogsAggregation_Count_ struct {
	Count *LogsAggregation_Count `protobuf:"bytes,1,opt,name=count,proto3,oneof"`
}

type LogsAggregation_CountDistinct_ struct {
	CountDistinct *LogsAggregation_CountDistinct `protobuf:"bytes,2,opt,name=count_distinct,json=countDistinct,proto3,oneof"`
}

type LogsAggregation_Sum_ struct {
	Sum *LogsAggregation_Sum `protobuf:"bytes,3,opt,name=sum,proto3,oneof"`
}

type LogsAggregation_Average_ struct {
	Average *LogsAggregation_Average `protobuf:"bytes,4,opt,name=average,proto3,oneof"`
}

type LogsAggregation_Min_ struct {
	Min *LogsAggregation_Min `protobuf:"bytes,5,opt,name=min,proto3,oneof"`
}

type LogsAggregation_Max_ struct {
	Max *LogsAggregation_Max `protobuf:"bytes,6,opt,name=max,proto3,oneof"`
}

type LogsAggregation_Percentile_ struct {
	Percentile *LogsAggregation_Percentile `protobuf:"bytes,7,opt,name=percentile,proto3,oneof"`
}

func (*LogsAggregation_Count_) isLogsAggregation_Value() {}

func (*LogsAggregation_CountDistinct_) isLogsAggregation_Value() {}

func (*LogsAggregation_Sum_) isLogsAggregation_Value() {}

func (*LogsAggregation_Average_) isLogsAggregation_Value() {}

func (*LogsAggregation_Min_) isLogsAggregation_Value() {}

func (*LogsAggregation_Max_) isLogsAggregation_Value() {}

func (*LogsAggregation_Percentile_) isLogsAggregation_Value() {}

type LogsAggregation_Count struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogsAggregation_Count) Reset() {
	*x = LogsAggregation_Count{}
	mi := &file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsAggregation_Count) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsAggregation_Count) ProtoMessage() {}

func (x *LogsAggregation_Count) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsAggregation_Count.ProtoReflect.Descriptor instead.
func (*LogsAggregation_Count) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_rawDescGZIP(), []int{0, 0}
}

type LogsAggregation_CountDistinct struct {
	state            protoimpl.MessageState  `protogen:"open.v1"`
	Field            *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	ObservationField *ObservationField       `protobuf:"bytes,2,opt,name=observation_field,json=observationField,proto3" json:"observation_field,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *LogsAggregation_CountDistinct) Reset() {
	*x = LogsAggregation_CountDistinct{}
	mi := &file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsAggregation_CountDistinct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsAggregation_CountDistinct) ProtoMessage() {}

func (x *LogsAggregation_CountDistinct) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsAggregation_CountDistinct.ProtoReflect.Descriptor instead.
func (*LogsAggregation_CountDistinct) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_rawDescGZIP(), []int{0, 1}
}

func (x *LogsAggregation_CountDistinct) GetField() *wrapperspb.StringValue {
	if x != nil {
		return x.Field
	}
	return nil
}

func (x *LogsAggregation_CountDistinct) GetObservationField() *ObservationField {
	if x != nil {
		return x.ObservationField
	}
	return nil
}

type LogsAggregation_Sum struct {
	state            protoimpl.MessageState  `protogen:"open.v1"`
	Field            *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	ObservationField *ObservationField       `protobuf:"bytes,2,opt,name=observation_field,json=observationField,proto3" json:"observation_field,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *LogsAggregation_Sum) Reset() {
	*x = LogsAggregation_Sum{}
	mi := &file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsAggregation_Sum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsAggregation_Sum) ProtoMessage() {}

func (x *LogsAggregation_Sum) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsAggregation_Sum.ProtoReflect.Descriptor instead.
func (*LogsAggregation_Sum) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_rawDescGZIP(), []int{0, 2}
}

func (x *LogsAggregation_Sum) GetField() *wrapperspb.StringValue {
	if x != nil {
		return x.Field
	}
	return nil
}

func (x *LogsAggregation_Sum) GetObservationField() *ObservationField {
	if x != nil {
		return x.ObservationField
	}
	return nil
}

type LogsAggregation_Average struct {
	state            protoimpl.MessageState  `protogen:"open.v1"`
	Field            *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	ObservationField *ObservationField       `protobuf:"bytes,2,opt,name=observation_field,json=observationField,proto3" json:"observation_field,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *LogsAggregation_Average) Reset() {
	*x = LogsAggregation_Average{}
	mi := &file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsAggregation_Average) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsAggregation_Average) ProtoMessage() {}

func (x *LogsAggregation_Average) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsAggregation_Average.ProtoReflect.Descriptor instead.
func (*LogsAggregation_Average) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_rawDescGZIP(), []int{0, 3}
}

func (x *LogsAggregation_Average) GetField() *wrapperspb.StringValue {
	if x != nil {
		return x.Field
	}
	return nil
}

func (x *LogsAggregation_Average) GetObservationField() *ObservationField {
	if x != nil {
		return x.ObservationField
	}
	return nil
}

type LogsAggregation_Min struct {
	state            protoimpl.MessageState  `protogen:"open.v1"`
	Field            *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	ObservationField *ObservationField       `protobuf:"bytes,2,opt,name=observation_field,json=observationField,proto3" json:"observation_field,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *LogsAggregation_Min) Reset() {
	*x = LogsAggregation_Min{}
	mi := &file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsAggregation_Min) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsAggregation_Min) ProtoMessage() {}

func (x *LogsAggregation_Min) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsAggregation_Min.ProtoReflect.Descriptor instead.
func (*LogsAggregation_Min) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_rawDescGZIP(), []int{0, 4}
}

func (x *LogsAggregation_Min) GetField() *wrapperspb.StringValue {
	if x != nil {
		return x.Field
	}
	return nil
}

func (x *LogsAggregation_Min) GetObservationField() *ObservationField {
	if x != nil {
		return x.ObservationField
	}
	return nil
}

type LogsAggregation_Max struct {
	state            protoimpl.MessageState  `protogen:"open.v1"`
	Field            *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	ObservationField *ObservationField       `protobuf:"bytes,2,opt,name=observation_field,json=observationField,proto3" json:"observation_field,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *LogsAggregation_Max) Reset() {
	*x = LogsAggregation_Max{}
	mi := &file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsAggregation_Max) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsAggregation_Max) ProtoMessage() {}

func (x *LogsAggregation_Max) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsAggregation_Max.ProtoReflect.Descriptor instead.
func (*LogsAggregation_Max) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_rawDescGZIP(), []int{0, 5}
}

func (x *LogsAggregation_Max) GetField() *wrapperspb.StringValue {
	if x != nil {
		return x.Field
	}
	return nil
}

func (x *LogsAggregation_Max) GetObservationField() *ObservationField {
	if x != nil {
		return x.ObservationField
	}
	return nil
}

type LogsAggregation_Percentile struct {
	state protoimpl.MessageState  `protogen:"open.v1"`
	Field *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	// Value in range (0, 100]
	Percent          *wrapperspb.DoubleValue `protobuf:"bytes,2,opt,name=percent,proto3" json:"percent,omitempty"`
	ObservationField *ObservationField       `protobuf:"bytes,3,opt,name=observation_field,json=observationField,proto3" json:"observation_field,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *LogsAggregation_Percentile) Reset() {
	*x = LogsAggregation_Percentile{}
	mi := &file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsAggregation_Percentile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsAggregation_Percentile) ProtoMessage() {}

func (x *LogsAggregation_Percentile) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsAggregation_Percentile.ProtoReflect.Descriptor instead.
func (*LogsAggregation_Percentile) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_rawDescGZIP(), []int{0, 6}
}

func (x *LogsAggregation_Percentile) GetField() *wrapperspb.StringValue {
	if x != nil {
		return x.Field
	}
	return nil
}

func (x *LogsAggregation_Percentile) GetPercent() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Percent
	}
	return nil
}

func (x *LogsAggregation_Percentile) GetObservationField() *ObservationField {
	if x != nil {
		return x.ObservationField
	}
	return nil
}

var File_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto protoreflect.FileDescriptor

var file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x61, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x3e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x0d, 0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x73,
	0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x55, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x6e, 0x0a, 0x0e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x74,
	0x69, 0x6e, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x63,
	0x74, 0x48, 0x00, 0x52, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x63, 0x74, 0x12, 0x4f, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x3b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x41, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x6d, 0x48, 0x00, 0x52, 0x03,
	0x73, 0x75, 0x6d, 0x12, 0x5b, 0x0a, 0x07, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f,
	0x67, 0x73, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x76,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x07, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65,
	0x12, 0x4f, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x41, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x69, 0x6e, 0x48, 0x00, 0x52, 0x03, 0x6d, 0x69,
	0x6e, 0x12, 0x4f, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x41, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x61, 0x78, 0x48, 0x00, 0x52, 0x03, 0x6d,
	0x61, 0x78, 0x12, 0x64, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x6c, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x4c, 0x6f, 0x67, 0x73, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x70, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x6c, 0x65, 0x1a, 0x07, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x1a, 0xaa, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x74, 0x69,
	0x6e, 0x63, 0x74, 0x12, 0x32, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x65, 0x0a, 0x11, 0x6f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x62, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x10, 0x6f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x1a, 0xa0,
	0x01, 0x0a, 0x03, 0x53, 0x75, 0x6d, 0x12, 0x32, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x65, 0x0a, 0x11, 0x6f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52,
	0x10, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x1a, 0xa4, 0x01, 0x0a, 0x07, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a,
	0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x65, 0x0a, 0x11, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x10, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x1a, 0xa0, 0x01, 0x0a, 0x03, 0x4d, 0x69, 0x6e,
	0x12, 0x32, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x65, 0x0a, 0x11, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x10, 0x6f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x1a, 0xa0, 0x01, 0x0a, 0x03,
	0x4d, 0x61, 0x78, 0x12, 0x32, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x65, 0x0a, 0x11, 0x6f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x62, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x10, 0x6f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x1a, 0xdf,
	0x01, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x6c, 0x65, 0x12, 0x32, 0x0a,
	0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x36, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x07, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x65, 0x0a, 0x11, 0x6f, 0x62, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x10,
	0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_rawDescOnce sync.Once
	file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_rawDescData = file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_rawDesc
)

func file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_rawDescData)
	})
	return file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_rawDescData
}

var file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_goTypes = []any{
	(*LogsAggregation)(nil),               // 0: com.coralogixapis.dashboards.v1.common.LogsAggregation
	(*LogsAggregation_Count)(nil),         // 1: com.coralogixapis.dashboards.v1.common.LogsAggregation.Count
	(*LogsAggregation_CountDistinct)(nil), // 2: com.coralogixapis.dashboards.v1.common.LogsAggregation.CountDistinct
	(*LogsAggregation_Sum)(nil),           // 3: com.coralogixapis.dashboards.v1.common.LogsAggregation.Sum
	(*LogsAggregation_Average)(nil),       // 4: com.coralogixapis.dashboards.v1.common.LogsAggregation.Average
	(*LogsAggregation_Min)(nil),           // 5: com.coralogixapis.dashboards.v1.common.LogsAggregation.Min
	(*LogsAggregation_Max)(nil),           // 6: com.coralogixapis.dashboards.v1.common.LogsAggregation.Max
	(*LogsAggregation_Percentile)(nil),    // 7: com.coralogixapis.dashboards.v1.common.LogsAggregation.Percentile
	(*wrapperspb.StringValue)(nil),        // 8: google.protobuf.StringValue
	(*ObservationField)(nil),              // 9: com.coralogixapis.dashboards.v1.common.ObservationField
	(*wrapperspb.DoubleValue)(nil),        // 10: google.protobuf.DoubleValue
}
var file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_depIdxs = []int32{
	1,  // 0: com.coralogixapis.dashboards.v1.common.LogsAggregation.count:type_name -> com.coralogixapis.dashboards.v1.common.LogsAggregation.Count
	2,  // 1: com.coralogixapis.dashboards.v1.common.LogsAggregation.count_distinct:type_name -> com.coralogixapis.dashboards.v1.common.LogsAggregation.CountDistinct
	3,  // 2: com.coralogixapis.dashboards.v1.common.LogsAggregation.sum:type_name -> com.coralogixapis.dashboards.v1.common.LogsAggregation.Sum
	4,  // 3: com.coralogixapis.dashboards.v1.common.LogsAggregation.average:type_name -> com.coralogixapis.dashboards.v1.common.LogsAggregation.Average
	5,  // 4: com.coralogixapis.dashboards.v1.common.LogsAggregation.min:type_name -> com.coralogixapis.dashboards.v1.common.LogsAggregation.Min
	6,  // 5: com.coralogixapis.dashboards.v1.common.LogsAggregation.max:type_name -> com.coralogixapis.dashboards.v1.common.LogsAggregation.Max
	7,  // 6: com.coralogixapis.dashboards.v1.common.LogsAggregation.percentile:type_name -> com.coralogixapis.dashboards.v1.common.LogsAggregation.Percentile
	8,  // 7: com.coralogixapis.dashboards.v1.common.LogsAggregation.CountDistinct.field:type_name -> google.protobuf.StringValue
	9,  // 8: com.coralogixapis.dashboards.v1.common.LogsAggregation.CountDistinct.observation_field:type_name -> com.coralogixapis.dashboards.v1.common.ObservationField
	8,  // 9: com.coralogixapis.dashboards.v1.common.LogsAggregation.Sum.field:type_name -> google.protobuf.StringValue
	9,  // 10: com.coralogixapis.dashboards.v1.common.LogsAggregation.Sum.observation_field:type_name -> com.coralogixapis.dashboards.v1.common.ObservationField
	8,  // 11: com.coralogixapis.dashboards.v1.common.LogsAggregation.Average.field:type_name -> google.protobuf.StringValue
	9,  // 12: com.coralogixapis.dashboards.v1.common.LogsAggregation.Average.observation_field:type_name -> com.coralogixapis.dashboards.v1.common.ObservationField
	8,  // 13: com.coralogixapis.dashboards.v1.common.LogsAggregation.Min.field:type_name -> google.protobuf.StringValue
	9,  // 14: com.coralogixapis.dashboards.v1.common.LogsAggregation.Min.observation_field:type_name -> com.coralogixapis.dashboards.v1.common.ObservationField
	8,  // 15: com.coralogixapis.dashboards.v1.common.LogsAggregation.Max.field:type_name -> google.protobuf.StringValue
	9,  // 16: com.coralogixapis.dashboards.v1.common.LogsAggregation.Max.observation_field:type_name -> com.coralogixapis.dashboards.v1.common.ObservationField
	8,  // 17: com.coralogixapis.dashboards.v1.common.LogsAggregation.Percentile.field:type_name -> google.protobuf.StringValue
	10, // 18: com.coralogixapis.dashboards.v1.common.LogsAggregation.Percentile.percent:type_name -> google.protobuf.DoubleValue
	9,  // 19: com.coralogixapis.dashboards.v1.common.LogsAggregation.Percentile.observation_field:type_name -> com.coralogixapis.dashboards.v1.common.ObservationField
	20, // [20:20] is the sub-list for method output_type
	20, // [20:20] is the sub-list for method input_type
	20, // [20:20] is the sub-list for extension type_name
	20, // [20:20] is the sub-list for extension extendee
	0,  // [0:20] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_init() }
func file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_init() {
	if File_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto != nil {
		return
	}
	file_com_coralogixapis_dashboards_v1_common_observation_field_proto_init()
	file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_msgTypes[0].OneofWrappers = []any{
		(*LogsAggregation_Count_)(nil),
		(*LogsAggregation_CountDistinct_)(nil),
		(*LogsAggregation_Sum_)(nil),
		(*LogsAggregation_Average_)(nil),
		(*LogsAggregation_Min_)(nil),
		(*LogsAggregation_Max_)(nil),
		(*LogsAggregation_Percentile_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto = out.File
	file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_rawDesc = nil
	file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_goTypes = nil
	file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_depIdxs = nil
}
