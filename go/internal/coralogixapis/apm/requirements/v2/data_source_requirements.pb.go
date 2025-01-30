// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.2
// source: com/coralogixapis/apm/requirements/v2/data_source_requirements.proto

package v2

import (
	v2 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/apm/common/v2"
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

type LabelExistence struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Label         *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Exists        *wrapperspb.BoolValue   `protobuf:"bytes,2,opt,name=exists,proto3" json:"exists,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LabelExistence) Reset() {
	*x = LabelExistence{}
	mi := &file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LabelExistence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelExistence) ProtoMessage() {}

func (x *LabelExistence) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelExistence.ProtoReflect.Descriptor instead.
func (*LabelExistence) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_rawDescGZIP(), []int{0}
}

func (x *LabelExistence) GetLabel() *wrapperspb.StringValue {
	if x != nil {
		return x.Label
	}
	return nil
}

func (x *LabelExistence) GetExists() *wrapperspb.BoolValue {
	if x != nil {
		return x.Exists
	}
	return nil
}

type MetricExistence struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Metric        *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=metric,proto3" json:"metric,omitempty"`
	Exists        *wrapperspb.BoolValue   `protobuf:"bytes,2,opt,name=exists,proto3" json:"exists,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MetricExistence) Reset() {
	*x = MetricExistence{}
	mi := &file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetricExistence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricExistence) ProtoMessage() {}

func (x *MetricExistence) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricExistence.ProtoReflect.Descriptor instead.
func (*MetricExistence) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_rawDescGZIP(), []int{1}
}

func (x *MetricExistence) GetMetric() *wrapperspb.StringValue {
	if x != nil {
		return x.Metric
	}
	return nil
}

func (x *MetricExistence) GetExists() *wrapperspb.BoolValue {
	if x != nil {
		return x.Exists
	}
	return nil
}

type MetricLabelsExistence struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Metric        *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=metric,proto3" json:"metric,omitempty"`
	Labels        []*LabelExistence       `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MetricLabelsExistence) Reset() {
	*x = MetricLabelsExistence{}
	mi := &file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetricLabelsExistence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricLabelsExistence) ProtoMessage() {}

func (x *MetricLabelsExistence) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricLabelsExistence.ProtoReflect.Descriptor instead.
func (*MetricLabelsExistence) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_rawDescGZIP(), []int{2}
}

func (x *MetricLabelsExistence) GetMetric() *wrapperspb.StringValue {
	if x != nil {
		return x.Metric
	}
	return nil
}

func (x *MetricLabelsExistence) GetLabels() []*LabelExistence {
	if x != nil {
		return x.Labels
	}
	return nil
}

type Requirements struct {
	state                protoimpl.MessageState   `protogen:"open.v1"`
	MetricExistence      []*MetricExistence       `protobuf:"bytes,1,rep,name=metric_existence,json=metricExistence,proto3" json:"metric_existence,omitempty"`
	MetricLabelExistence []*MetricLabelsExistence `protobuf:"bytes,2,rep,name=metric_label_existence,json=metricLabelExistence,proto3" json:"metric_label_existence,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *Requirements) Reset() {
	*x = Requirements{}
	mi := &file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Requirements) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Requirements) ProtoMessage() {}

func (x *Requirements) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Requirements.ProtoReflect.Descriptor instead.
func (*Requirements) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_rawDescGZIP(), []int{3}
}

func (x *Requirements) GetMetricExistence() []*MetricExistence {
	if x != nil {
		return x.MetricExistence
	}
	return nil
}

func (x *Requirements) GetMetricLabelExistence() []*MetricLabelsExistence {
	if x != nil {
		return x.MetricLabelExistence
	}
	return nil
}

type MissingNothing struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MissingNothing) Reset() {
	*x = MissingNothing{}
	mi := &file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MissingNothing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MissingNothing) ProtoMessage() {}

func (x *MissingNothing) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MissingNothing.ProtoReflect.Descriptor instead.
func (*MissingNothing) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_rawDescGZIP(), []int{4}
}

type PageRequirement struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	Requirements       *Requirements          `protobuf:"bytes,1,opt,name=requirements,proto3" json:"requirements,omitempty"`
	PassedRequirements *wrapperspb.BoolValue  `protobuf:"bytes,2,opt,name=passed_requirements,json=passedRequirements,proto3" json:"passed_requirements,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *PageRequirement) Reset() {
	*x = PageRequirement{}
	mi := &file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PageRequirement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageRequirement) ProtoMessage() {}

func (x *PageRequirement) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageRequirement.ProtoReflect.Descriptor instead.
func (*PageRequirement) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_rawDescGZIP(), []int{5}
}

func (x *PageRequirement) GetRequirements() *Requirements {
	if x != nil {
		return x.Requirements
	}
	return nil
}

func (x *PageRequirement) GetPassedRequirements() *wrapperspb.BoolValue {
	if x != nil {
		return x.PassedRequirements
	}
	return nil
}

type DataSourceRequirements struct {
	state              protoimpl.MessageState      `protogen:"open.v1"`
	DataSource         *v2.DataSource              `protobuf:"bytes,1,opt,name=data_source,json=dataSource,proto3" json:"data_source,omitempty"`
	Requirements       *Requirements               `protobuf:"bytes,2,opt,name=requirements,proto3" json:"requirements,omitempty"`
	PassedRequirements *wrapperspb.BoolValue       `protobuf:"bytes,3,opt,name=passed_requirements,json=passedRequirements,proto3" json:"passed_requirements,omitempty"`
	PageRequirements   map[string]*PageRequirement `protobuf:"bytes,4,rep,name=page_requirements,json=pageRequirements,proto3" json:"page_requirements,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *DataSourceRequirements) Reset() {
	*x = DataSourceRequirements{}
	mi := &file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataSourceRequirements) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSourceRequirements) ProtoMessage() {}

func (x *DataSourceRequirements) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSourceRequirements.ProtoReflect.Descriptor instead.
func (*DataSourceRequirements) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_rawDescGZIP(), []int{6}
}

func (x *DataSourceRequirements) GetDataSource() *v2.DataSource {
	if x != nil {
		return x.DataSource
	}
	return nil
}

func (x *DataSourceRequirements) GetRequirements() *Requirements {
	if x != nil {
		return x.Requirements
	}
	return nil
}

func (x *DataSourceRequirements) GetPassedRequirements() *wrapperspb.BoolValue {
	if x != nil {
		return x.PassedRequirements
	}
	return nil
}

func (x *DataSourceRequirements) GetPageRequirements() map[string]*PageRequirement {
	if x != nil {
		return x.PageRequirements
	}
	return nil
}

var File_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto protoreflect.FileDescriptor

var file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_rawDesc = string([]byte{
	0x0a, 0x44, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x32, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x78, 0x0a, 0x0e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x45, 0x78, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x32, 0x0a, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x22, 0x7b, 0x0a, 0x0f, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x45, 0x78, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x34, 0x0a,
	0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x12, 0x32, 0x0a, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x22, 0x9c, 0x01, 0x0a, 0x15, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63,
	0x65, 0x12, 0x34, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x4d, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x32, 0x2e,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x45, 0x78, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0xe5, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x61, 0x0a, 0x10, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0f, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x45, 0x78, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x72, 0x0a, 0x16, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x65, 0x78, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61,
	0x70, 0x6d, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x76, 0x32, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45,
	0x78, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x14, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x45, 0x78, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x10,
	0x0a, 0x0e, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67,
	0x22, 0xb7, 0x01, 0x0a, 0x0f, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x57, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61,
	0x70, 0x6d, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x76, 0x32, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x4b, 0x0a,
	0x13, 0x70, 0x61, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f,
	0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x12, 0x70, 0x61, 0x73, 0x73, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x8c, 0x04, 0x0a, 0x16, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x4c, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61,
	0x70, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70,
	0x6d, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x0c,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x4b, 0x0a, 0x13,
	0x70, 0x61, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x12, 0x70, 0x61, 0x73, 0x73, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x80, 0x01, 0x0a, 0x11, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x53, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x70, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x7b, 0x0a, 0x15,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x4c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_rawDescOnce sync.Once
	file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_rawDescData []byte
)

func file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_rawDesc), len(file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_rawDesc)))
	})
	return file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_rawDescData
}

var file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_goTypes = []any{
	(*LabelExistence)(nil),         // 0: com.coralogixapis.apm.requirements.v2.LabelExistence
	(*MetricExistence)(nil),        // 1: com.coralogixapis.apm.requirements.v2.MetricExistence
	(*MetricLabelsExistence)(nil),  // 2: com.coralogixapis.apm.requirements.v2.MetricLabelsExistence
	(*Requirements)(nil),           // 3: com.coralogixapis.apm.requirements.v2.Requirements
	(*MissingNothing)(nil),         // 4: com.coralogixapis.apm.requirements.v2.MissingNothing
	(*PageRequirement)(nil),        // 5: com.coralogixapis.apm.requirements.v2.PageRequirement
	(*DataSourceRequirements)(nil), // 6: com.coralogixapis.apm.requirements.v2.DataSourceRequirements
	nil,                            // 7: com.coralogixapis.apm.requirements.v2.DataSourceRequirements.PageRequirementsEntry
	(*wrapperspb.StringValue)(nil), // 8: google.protobuf.StringValue
	(*wrapperspb.BoolValue)(nil),   // 9: google.protobuf.BoolValue
	(*v2.DataSource)(nil),          // 10: com.coralogixapis.apm.common.v2.DataSource
}
var file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_depIdxs = []int32{
	8,  // 0: com.coralogixapis.apm.requirements.v2.LabelExistence.label:type_name -> google.protobuf.StringValue
	9,  // 1: com.coralogixapis.apm.requirements.v2.LabelExistence.exists:type_name -> google.protobuf.BoolValue
	8,  // 2: com.coralogixapis.apm.requirements.v2.MetricExistence.metric:type_name -> google.protobuf.StringValue
	9,  // 3: com.coralogixapis.apm.requirements.v2.MetricExistence.exists:type_name -> google.protobuf.BoolValue
	8,  // 4: com.coralogixapis.apm.requirements.v2.MetricLabelsExistence.metric:type_name -> google.protobuf.StringValue
	0,  // 5: com.coralogixapis.apm.requirements.v2.MetricLabelsExistence.labels:type_name -> com.coralogixapis.apm.requirements.v2.LabelExistence
	1,  // 6: com.coralogixapis.apm.requirements.v2.Requirements.metric_existence:type_name -> com.coralogixapis.apm.requirements.v2.MetricExistence
	2,  // 7: com.coralogixapis.apm.requirements.v2.Requirements.metric_label_existence:type_name -> com.coralogixapis.apm.requirements.v2.MetricLabelsExistence
	3,  // 8: com.coralogixapis.apm.requirements.v2.PageRequirement.requirements:type_name -> com.coralogixapis.apm.requirements.v2.Requirements
	9,  // 9: com.coralogixapis.apm.requirements.v2.PageRequirement.passed_requirements:type_name -> google.protobuf.BoolValue
	10, // 10: com.coralogixapis.apm.requirements.v2.DataSourceRequirements.data_source:type_name -> com.coralogixapis.apm.common.v2.DataSource
	3,  // 11: com.coralogixapis.apm.requirements.v2.DataSourceRequirements.requirements:type_name -> com.coralogixapis.apm.requirements.v2.Requirements
	9,  // 12: com.coralogixapis.apm.requirements.v2.DataSourceRequirements.passed_requirements:type_name -> google.protobuf.BoolValue
	7,  // 13: com.coralogixapis.apm.requirements.v2.DataSourceRequirements.page_requirements:type_name -> com.coralogixapis.apm.requirements.v2.DataSourceRequirements.PageRequirementsEntry
	5,  // 14: com.coralogixapis.apm.requirements.v2.DataSourceRequirements.PageRequirementsEntry.value:type_name -> com.coralogixapis.apm.requirements.v2.PageRequirement
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_init() }
func file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_init() {
	if File_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_rawDesc), len(file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto = out.File
	file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_goTypes = nil
	file_com_coralogixapis_apm_requirements_v2_data_source_requirements_proto_depIdxs = nil
}
