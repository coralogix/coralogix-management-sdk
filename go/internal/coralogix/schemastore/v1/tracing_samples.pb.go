// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: com/coralogix/schemastore/v1/tracing_samples.proto

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

type TracingSamplesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamIds      []uint32       `protobuf:"varint,1,rep,packed,name=team_ids,json=teamIds,proto3" json:"team_ids,omitempty"`
	Path         string         `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	TimeRange    *TimeRange     `protobuf:"bytes,3,opt,name=time_range,json=timeRange,proto3" json:"time_range,omitempty"`
	LabelFilters []*LabelFilter `protobuf:"bytes,4,rep,name=label_filters,json=labelFilters,proto3" json:"label_filters,omitempty"`
	ResultLimit  uint32         `protobuf:"varint,5,opt,name=result_limit,json=resultLimit,proto3" json:"result_limit,omitempty"`
	// deprecated
	CalculateEstimatedCardinality bool `protobuf:"varint,6,opt,name=calculate_estimated_cardinality,json=calculateEstimatedCardinality,proto3" json:"calculate_estimated_cardinality,omitempty"`
}

func (x *TracingSamplesRequest) Reset() {
	*x = TracingSamplesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_schemastore_v1_tracing_samples_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TracingSamplesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TracingSamplesRequest) ProtoMessage() {}

func (x *TracingSamplesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_schemastore_v1_tracing_samples_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TracingSamplesRequest.ProtoReflect.Descriptor instead.
func (*TracingSamplesRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_schemastore_v1_tracing_samples_proto_rawDescGZIP(), []int{0}
}

func (x *TracingSamplesRequest) GetTeamIds() []uint32 {
	if x != nil {
		return x.TeamIds
	}
	return nil
}

func (x *TracingSamplesRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *TracingSamplesRequest) GetTimeRange() *TimeRange {
	if x != nil {
		return x.TimeRange
	}
	return nil
}

func (x *TracingSamplesRequest) GetLabelFilters() []*LabelFilter {
	if x != nil {
		return x.LabelFilters
	}
	return nil
}

func (x *TracingSamplesRequest) GetResultLimit() uint32 {
	if x != nil {
		return x.ResultLimit
	}
	return 0
}

func (x *TracingSamplesRequest) GetCalculateEstimatedCardinality() bool {
	if x != nil {
		return x.CalculateEstimatedCardinality
	}
	return false
}

type TracingSamplesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Samples []string `protobuf:"bytes,1,rep,name=samples,proto3" json:"samples,omitempty"`
}

func (x *TracingSamplesResponse) Reset() {
	*x = TracingSamplesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_schemastore_v1_tracing_samples_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TracingSamplesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TracingSamplesResponse) ProtoMessage() {}

func (x *TracingSamplesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_schemastore_v1_tracing_samples_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TracingSamplesResponse.ProtoReflect.Descriptor instead.
func (*TracingSamplesResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_schemastore_v1_tracing_samples_proto_rawDescGZIP(), []int{1}
}

func (x *TracingSamplesResponse) GetSamples() []string {
	if x != nil {
		return x.Samples
	}
	return nil
}

type BatchTracingSamplesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamIds      []uint32       `protobuf:"varint,1,rep,packed,name=team_ids,json=teamIds,proto3" json:"team_ids,omitempty"`
	Paths        []string       `protobuf:"bytes,2,rep,name=paths,proto3" json:"paths,omitempty"`
	TimeRange    *TimeRange     `protobuf:"bytes,3,opt,name=time_range,json=timeRange,proto3" json:"time_range,omitempty"`
	LabelFilters []*LabelFilter `protobuf:"bytes,4,rep,name=label_filters,json=labelFilters,proto3" json:"label_filters,omitempty"`
	ResultLimit  uint32         `protobuf:"varint,5,opt,name=result_limit,json=resultLimit,proto3" json:"result_limit,omitempty"`
	GroupByLabel string         `protobuf:"bytes,6,opt,name=group_by_label,json=groupByLabel,proto3" json:"group_by_label,omitempty"`
}

func (x *BatchTracingSamplesRequest) Reset() {
	*x = BatchTracingSamplesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_schemastore_v1_tracing_samples_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchTracingSamplesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchTracingSamplesRequest) ProtoMessage() {}

func (x *BatchTracingSamplesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_schemastore_v1_tracing_samples_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchTracingSamplesRequest.ProtoReflect.Descriptor instead.
func (*BatchTracingSamplesRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_schemastore_v1_tracing_samples_proto_rawDescGZIP(), []int{2}
}

func (x *BatchTracingSamplesRequest) GetTeamIds() []uint32 {
	if x != nil {
		return x.TeamIds
	}
	return nil
}

func (x *BatchTracingSamplesRequest) GetPaths() []string {
	if x != nil {
		return x.Paths
	}
	return nil
}

func (x *BatchTracingSamplesRequest) GetTimeRange() *TimeRange {
	if x != nil {
		return x.TimeRange
	}
	return nil
}

func (x *BatchTracingSamplesRequest) GetLabelFilters() []*LabelFilter {
	if x != nil {
		return x.LabelFilters
	}
	return nil
}

func (x *BatchTracingSamplesRequest) GetResultLimit() uint32 {
	if x != nil {
		return x.ResultLimit
	}
	return 0
}

func (x *BatchTracingSamplesRequest) GetGroupByLabel() string {
	if x != nil {
		return x.GroupByLabel
	}
	return ""
}

type BatchTracingSamplesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupedBy string           `protobuf:"bytes,1,opt,name=grouped_by,json=groupedBy,proto3" json:"grouped_by,omitempty"`
	Result    []*SamplesResult `protobuf:"bytes,2,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *BatchTracingSamplesResponse) Reset() {
	*x = BatchTracingSamplesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_schemastore_v1_tracing_samples_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchTracingSamplesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchTracingSamplesResponse) ProtoMessage() {}

func (x *BatchTracingSamplesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_schemastore_v1_tracing_samples_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchTracingSamplesResponse.ProtoReflect.Descriptor instead.
func (*BatchTracingSamplesResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_schemastore_v1_tracing_samples_proto_rawDescGZIP(), []int{3}
}

func (x *BatchTracingSamplesResponse) GetGroupedBy() string {
	if x != nil {
		return x.GroupedBy
	}
	return ""
}

func (x *BatchTracingSamplesResponse) GetResult() []*SamplesResult {
	if x != nil {
		return x.Result
	}
	return nil
}

type SamplesResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path               string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	AdditionalGrouping string   `protobuf:"bytes,2,opt,name=additional_grouping,json=additionalGrouping,proto3" json:"additional_grouping,omitempty"`
	Samples            []string `protobuf:"bytes,3,rep,name=samples,proto3" json:"samples,omitempty"`
}

func (x *SamplesResult) Reset() {
	*x = SamplesResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_schemastore_v1_tracing_samples_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SamplesResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SamplesResult) ProtoMessage() {}

func (x *SamplesResult) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_schemastore_v1_tracing_samples_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SamplesResult.ProtoReflect.Descriptor instead.
func (*SamplesResult) Descriptor() ([]byte, []int) {
	return file_com_coralogix_schemastore_v1_tracing_samples_proto_rawDescGZIP(), []int{4}
}

func (x *SamplesResult) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *SamplesResult) GetAdditionalGrouping() string {
	if x != nil {
		return x.AdditionalGrouping
	}
	return ""
}

func (x *SamplesResult) GetSamples() []string {
	if x != nil {
		return x.Samples
	}
	return nil
}

var File_com_coralogix_schemastore_v1_tracing_samples_proto protoreflect.FileDescriptor

var file_com_coralogix_schemastore_v1_tracing_samples_proto_rawDesc = []byte{
	0x0a, 0x32, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x29, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x02,
	0x0a, 0x15, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x49,
	0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x46, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x4e,
	0x0a, 0x0d, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x52, 0x0c, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x21,
	0x0a, 0x0c, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x46, 0x0a, 0x1f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x65,
	0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1d, 0x63, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x43, 0x61,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x22, 0x38, 0x0a, 0x16, 0x54, 0x72, 0x61,
	0x63, 0x69, 0x6e, 0x67, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x4a, 0x04, 0x08,
	0x02, 0x10, 0x03, 0x22, 0xae, 0x02, 0x0a, 0x1a, 0x42, 0x61, 0x74, 0x63, 0x68, 0x54, 0x72, 0x61,
	0x63, 0x69, 0x6e, 0x67, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61,
	0x74, 0x68, 0x73, 0x12, 0x46, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0c, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x24,
	0x0a, 0x0e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x62, 0x79, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x79, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x22, 0x81, 0x01, 0x0a, 0x1b, 0x42, 0x61, 0x74, 0x63, 0x68, 0x54, 0x72,
	0x61, 0x63, 0x69, 0x6e, 0x67, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x65, 0x64, 0x5f,
	0x62, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x65,
	0x64, 0x42, 0x79, 0x12, 0x43, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x6e, 0x0a, 0x0d, 0x53, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x2f, 0x0a,
	0x13, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x61, 0x64, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_schemastore_v1_tracing_samples_proto_rawDescOnce sync.Once
	file_com_coralogix_schemastore_v1_tracing_samples_proto_rawDescData = file_com_coralogix_schemastore_v1_tracing_samples_proto_rawDesc
)

func file_com_coralogix_schemastore_v1_tracing_samples_proto_rawDescGZIP() []byte {
	file_com_coralogix_schemastore_v1_tracing_samples_proto_rawDescOnce.Do(func() {
		file_com_coralogix_schemastore_v1_tracing_samples_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_schemastore_v1_tracing_samples_proto_rawDescData)
	})
	return file_com_coralogix_schemastore_v1_tracing_samples_proto_rawDescData
}

var file_com_coralogix_schemastore_v1_tracing_samples_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_com_coralogix_schemastore_v1_tracing_samples_proto_goTypes = []any{
	(*TracingSamplesRequest)(nil),       // 0: com.coralogix.schemastore.v1.TracingSamplesRequest
	(*TracingSamplesResponse)(nil),      // 1: com.coralogix.schemastore.v1.TracingSamplesResponse
	(*BatchTracingSamplesRequest)(nil),  // 2: com.coralogix.schemastore.v1.BatchTracingSamplesRequest
	(*BatchTracingSamplesResponse)(nil), // 3: com.coralogix.schemastore.v1.BatchTracingSamplesResponse
	(*SamplesResult)(nil),               // 4: com.coralogix.schemastore.v1.SamplesResult
	(*TimeRange)(nil),                   // 5: com.coralogix.schemastore.v1.TimeRange
	(*LabelFilter)(nil),                 // 6: com.coralogix.schemastore.v1.LabelFilter
}
var file_com_coralogix_schemastore_v1_tracing_samples_proto_depIdxs = []int32{
	5, // 0: com.coralogix.schemastore.v1.TracingSamplesRequest.time_range:type_name -> com.coralogix.schemastore.v1.TimeRange
	6, // 1: com.coralogix.schemastore.v1.TracingSamplesRequest.label_filters:type_name -> com.coralogix.schemastore.v1.LabelFilter
	5, // 2: com.coralogix.schemastore.v1.BatchTracingSamplesRequest.time_range:type_name -> com.coralogix.schemastore.v1.TimeRange
	6, // 3: com.coralogix.schemastore.v1.BatchTracingSamplesRequest.label_filters:type_name -> com.coralogix.schemastore.v1.LabelFilter
	4, // 4: com.coralogix.schemastore.v1.BatchTracingSamplesResponse.result:type_name -> com.coralogix.schemastore.v1.SamplesResult
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_com_coralogix_schemastore_v1_tracing_samples_proto_init() }
func file_com_coralogix_schemastore_v1_tracing_samples_proto_init() {
	if File_com_coralogix_schemastore_v1_tracing_samples_proto != nil {
		return
	}
	file_com_coralogix_schemastore_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_schemastore_v1_tracing_samples_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TracingSamplesRequest); i {
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
		file_com_coralogix_schemastore_v1_tracing_samples_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TracingSamplesResponse); i {
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
		file_com_coralogix_schemastore_v1_tracing_samples_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*BatchTracingSamplesRequest); i {
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
		file_com_coralogix_schemastore_v1_tracing_samples_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*BatchTracingSamplesResponse); i {
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
		file_com_coralogix_schemastore_v1_tracing_samples_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*SamplesResult); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_schemastore_v1_tracing_samples_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_schemastore_v1_tracing_samples_proto_goTypes,
		DependencyIndexes: file_com_coralogix_schemastore_v1_tracing_samples_proto_depIdxs,
		MessageInfos:      file_com_coralogix_schemastore_v1_tracing_samples_proto_msgTypes,
	}.Build()
	File_com_coralogix_schemastore_v1_tracing_samples_proto = out.File
	file_com_coralogix_schemastore_v1_tracing_samples_proto_rawDesc = nil
	file_com_coralogix_schemastore_v1_tracing_samples_proto_goTypes = nil
	file_com_coralogix_schemastore_v1_tracing_samples_proto_depIdxs = nil
}
