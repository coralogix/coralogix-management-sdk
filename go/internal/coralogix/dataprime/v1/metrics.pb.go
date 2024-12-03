// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: com/coralogix/dataprime/v1/metrics.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MetricType int32

const (
	MetricType_METRIC_TYPE_UNSPECIFIED MetricType = 0
	MetricType_METRIC_TYPE_TIMESTAMP   MetricType = 1
	MetricType_METRIC_TYPE_COUNT       MetricType = 2
	MetricType_METRIC_TYPE_GAUGE       MetricType = 3
	MetricType_METRIC_TYPE_TIME        MetricType = 4
)

// Enum value maps for MetricType.
var (
	MetricType_name = map[int32]string{
		0: "METRIC_TYPE_UNSPECIFIED",
		1: "METRIC_TYPE_TIMESTAMP",
		2: "METRIC_TYPE_COUNT",
		3: "METRIC_TYPE_GAUGE",
		4: "METRIC_TYPE_TIME",
	}
	MetricType_value = map[string]int32{
		"METRIC_TYPE_UNSPECIFIED": 0,
		"METRIC_TYPE_TIMESTAMP":   1,
		"METRIC_TYPE_COUNT":       2,
		"METRIC_TYPE_GAUGE":       3,
		"METRIC_TYPE_TIME":        4,
	}
)

func (x MetricType) Enum() *MetricType {
	p := new(MetricType)
	*p = x
	return p
}

func (x MetricType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetricType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_dataprime_v1_metrics_proto_enumTypes[0].Descriptor()
}

func (MetricType) Type() protoreflect.EnumType {
	return &file_com_coralogix_dataprime_v1_metrics_proto_enumTypes[0]
}

func (x MetricType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetricType.Descriptor instead.
func (MetricType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_metrics_proto_rawDescGZIP(), []int{0}
}

type QueryMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueryId           string          `protobuf:"bytes,1,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
	StageMetrics      []*StageMetrics `protobuf:"bytes,2,rep,name=stage_metrics,json=stageMetrics,proto3" json:"stage_metrics,omitempty"`
	AdditionalMetrics []*Metric       `protobuf:"bytes,3,rep,name=additional_metrics,json=additionalMetrics,proto3" json:"additional_metrics,omitempty"`
}

func (x *QueryMetrics) Reset() {
	*x = QueryMetrics{}
	mi := &file_com_coralogix_dataprime_v1_metrics_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMetrics) ProtoMessage() {}

func (x *QueryMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_metrics_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMetrics.ProtoReflect.Descriptor instead.
func (*QueryMetrics) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_metrics_proto_rawDescGZIP(), []int{0}
}

func (x *QueryMetrics) GetQueryId() string {
	if x != nil {
		return x.QueryId
	}
	return ""
}

func (x *QueryMetrics) GetStageMetrics() []*StageMetrics {
	if x != nil {
		return x.StageMetrics
	}
	return nil
}

func (x *QueryMetrics) GetAdditionalMetrics() []*Metric {
	if x != nil {
		return x.AdditionalMetrics
	}
	return nil
}

type StageMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StageId         int64              `protobuf:"varint,1,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
	PartitionId     int64              `protobuf:"varint,2,opt,name=partition_id,json=partitionId,proto3" json:"partition_id,omitempty"`
	OperatorMetrics []*OperatorMetrics `protobuf:"bytes,3,rep,name=operator_metrics,json=operatorMetrics,proto3" json:"operator_metrics,omitempty"`
}

func (x *StageMetrics) Reset() {
	*x = StageMetrics{}
	mi := &file_com_coralogix_dataprime_v1_metrics_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StageMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StageMetrics) ProtoMessage() {}

func (x *StageMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_metrics_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StageMetrics.ProtoReflect.Descriptor instead.
func (*StageMetrics) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_metrics_proto_rawDescGZIP(), []int{1}
}

func (x *StageMetrics) GetStageId() int64 {
	if x != nil {
		return x.StageId
	}
	return 0
}

func (x *StageMetrics) GetPartitionId() int64 {
	if x != nil {
		return x.PartitionId
	}
	return 0
}

func (x *StageMetrics) GetOperatorMetrics() []*OperatorMetrics {
	if x != nil {
		return x.OperatorMetrics
	}
	return nil
}

type OperatorMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperatorName string    `protobuf:"bytes,1,opt,name=operator_name,json=operatorName,proto3" json:"operator_name,omitempty"`
	Metrics      []*Metric `protobuf:"bytes,2,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *OperatorMetrics) Reset() {
	*x = OperatorMetrics{}
	mi := &file_com_coralogix_dataprime_v1_metrics_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OperatorMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorMetrics) ProtoMessage() {}

func (x *OperatorMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_metrics_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperatorMetrics.ProtoReflect.Descriptor instead.
func (*OperatorMetrics) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_metrics_proto_rawDescGZIP(), []int{2}
}

func (x *OperatorMetrics) GetOperatorName() string {
	if x != nil {
		return x.OperatorName
	}
	return ""
}

func (x *OperatorMetrics) GetMetrics() []*Metric {
	if x != nil {
		return x.Metrics
	}
	return nil
}

type MetricLabel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MetricLabel) Reset() {
	*x = MetricLabel{}
	mi := &file_com_coralogix_dataprime_v1_metrics_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetricLabel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricLabel) ProtoMessage() {}

func (x *MetricLabel) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_metrics_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricLabel.ProtoReflect.Descriptor instead.
func (*MetricLabel) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_metrics_proto_rawDescGZIP(), []int{3}
}

func (x *MetricLabel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MetricLabel) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Metric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value      int64      `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Name       string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MetricType MetricType `protobuf:"varint,3,opt,name=metric_type,json=metricType,proto3,enum=com.coralogix.dataprime.v1.MetricType" json:"metric_type,omitempty"`
	Labels     []string   `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *Metric) Reset() {
	*x = Metric{}
	mi := &file_com_coralogix_dataprime_v1_metrics_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Metric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metric) ProtoMessage() {}

func (x *Metric) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_metrics_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metric.ProtoReflect.Descriptor instead.
func (*Metric) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_metrics_proto_rawDescGZIP(), []int{4}
}

func (x *Metric) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Metric) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Metric) GetMetricType() MetricType {
	if x != nil {
		return x.MetricType
	}
	return MetricType_METRIC_TYPE_UNSPECIFIED
}

func (x *Metric) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type GetMetricsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueryId string `protobuf:"bytes,1,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
}

func (x *GetMetricsRequest) Reset() {
	*x = GetMetricsRequest{}
	mi := &file_com_coralogix_dataprime_v1_metrics_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMetricsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetricsRequest) ProtoMessage() {}

func (x *GetMetricsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_metrics_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetricsRequest.ProtoReflect.Descriptor instead.
func (*GetMetricsRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_metrics_proto_rawDescGZIP(), []int{5}
}

func (x *GetMetricsRequest) GetQueryId() string {
	if x != nil {
		return x.QueryId
	}
	return ""
}

type GetMetricsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics *QueryMetrics `protobuf:"bytes,1,opt,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *GetMetricsResponse) Reset() {
	*x = GetMetricsResponse{}
	mi := &file_com_coralogix_dataprime_v1_metrics_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMetricsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetricsResponse) ProtoMessage() {}

func (x *GetMetricsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_metrics_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetricsResponse.ProtoReflect.Descriptor instead.
func (*GetMetricsResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_metrics_proto_rawDescGZIP(), []int{6}
}

func (x *GetMetricsResponse) GetMetrics() *QueryMetrics {
	if x != nil {
		return x.Metrics
	}
	return nil
}

var File_com_coralogix_dataprime_v1_metrics_proto protoreflect.FileDescriptor

var file_com_coralogix_dataprime_v1_metrics_proto_rawDesc = []byte{
	0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x22, 0xcb, 0x01, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x49, 0x64, 0x12, 0x4d, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x12, 0x51, 0x0a, 0x12, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x52, 0x11, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x22, 0xa4, 0x01, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x67, 0x65, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x56, 0x0a, 0x10, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0x74, 0x0a, 0x0f, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x23,
	0x0a, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x22, 0x37, 0x0a, 0x0b, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x06, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x47, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x22, 0x2e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x71, 0x75, 0x65, 0x72, 0x79, 0x49, 0x64,
	0x22, 0x58, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2a, 0x88, 0x01, 0x0a, 0x0a, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x45, 0x54,
	0x52, 0x49, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x10,
	0x01, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x45, 0x54, 0x52,
	0x49, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x41, 0x55, 0x47, 0x45, 0x10, 0x03, 0x12,
	0x14, 0x0a, 0x10, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54,
	0x49, 0x4d, 0x45, 0x10, 0x04, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_dataprime_v1_metrics_proto_rawDescOnce sync.Once
	file_com_coralogix_dataprime_v1_metrics_proto_rawDescData = file_com_coralogix_dataprime_v1_metrics_proto_rawDesc
)

func file_com_coralogix_dataprime_v1_metrics_proto_rawDescGZIP() []byte {
	file_com_coralogix_dataprime_v1_metrics_proto_rawDescOnce.Do(func() {
		file_com_coralogix_dataprime_v1_metrics_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_dataprime_v1_metrics_proto_rawDescData)
	})
	return file_com_coralogix_dataprime_v1_metrics_proto_rawDescData
}

var file_com_coralogix_dataprime_v1_metrics_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogix_dataprime_v1_metrics_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_com_coralogix_dataprime_v1_metrics_proto_goTypes = []any{
	(MetricType)(0),            // 0: com.coralogix.dataprime.v1.MetricType
	(*QueryMetrics)(nil),       // 1: com.coralogix.dataprime.v1.QueryMetrics
	(*StageMetrics)(nil),       // 2: com.coralogix.dataprime.v1.StageMetrics
	(*OperatorMetrics)(nil),    // 3: com.coralogix.dataprime.v1.OperatorMetrics
	(*MetricLabel)(nil),        // 4: com.coralogix.dataprime.v1.MetricLabel
	(*Metric)(nil),             // 5: com.coralogix.dataprime.v1.Metric
	(*GetMetricsRequest)(nil),  // 6: com.coralogix.dataprime.v1.GetMetricsRequest
	(*GetMetricsResponse)(nil), // 7: com.coralogix.dataprime.v1.GetMetricsResponse
}
var file_com_coralogix_dataprime_v1_metrics_proto_depIdxs = []int32{
	2, // 0: com.coralogix.dataprime.v1.QueryMetrics.stage_metrics:type_name -> com.coralogix.dataprime.v1.StageMetrics
	5, // 1: com.coralogix.dataprime.v1.QueryMetrics.additional_metrics:type_name -> com.coralogix.dataprime.v1.Metric
	3, // 2: com.coralogix.dataprime.v1.StageMetrics.operator_metrics:type_name -> com.coralogix.dataprime.v1.OperatorMetrics
	5, // 3: com.coralogix.dataprime.v1.OperatorMetrics.metrics:type_name -> com.coralogix.dataprime.v1.Metric
	0, // 4: com.coralogix.dataprime.v1.Metric.metric_type:type_name -> com.coralogix.dataprime.v1.MetricType
	1, // 5: com.coralogix.dataprime.v1.GetMetricsResponse.metrics:type_name -> com.coralogix.dataprime.v1.QueryMetrics
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_com_coralogix_dataprime_v1_metrics_proto_init() }
func file_com_coralogix_dataprime_v1_metrics_proto_init() {
	if File_com_coralogix_dataprime_v1_metrics_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_dataprime_v1_metrics_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_dataprime_v1_metrics_proto_goTypes,
		DependencyIndexes: file_com_coralogix_dataprime_v1_metrics_proto_depIdxs,
		EnumInfos:         file_com_coralogix_dataprime_v1_metrics_proto_enumTypes,
		MessageInfos:      file_com_coralogix_dataprime_v1_metrics_proto_msgTypes,
	}.Build()
	File_com_coralogix_dataprime_v1_metrics_proto = out.File
	file_com_coralogix_dataprime_v1_metrics_proto_rawDesc = nil
	file_com_coralogix_dataprime_v1_metrics_proto_goTypes = nil
	file_com_coralogix_dataprime_v1_metrics_proto_depIdxs = nil
}
