// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: com/coralogix/schemastore/v1/tracing_fields.proto

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

type TracingFieldsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TeamIds       []uint32               `protobuf:"varint,1,rep,packed,name=team_ids,json=teamIds,proto3" json:"team_ids,omitempty"`
	TimeRange     *TimeRange             `protobuf:"bytes,2,opt,name=time_range,json=timeRange,proto3" json:"time_range,omitempty"`
	ResultLimit   int32                  `protobuf:"varint,3,opt,name=result_limit,json=resultLimit,proto3" json:"result_limit,omitempty"`
	LabelFilters  []*LabelFilter         `protobuf:"bytes,4,rep,name=label_filters,json=labelFilters,proto3" json:"label_filters,omitempty"`
	SchemaFilters []*SchemaFilter        `protobuf:"bytes,5,rep,name=schema_filters,json=schemaFilters,proto3" json:"schema_filters,omitempty"`
	// Deprecated: please use paths instead
	//
	// Deprecated: Marked as deprecated in com/coralogix/schemastore/v1/tracing_fields.proto.
	Path           string         `protobuf:"bytes,6,opt,name=path,proto3" json:"path,omitempty"`
	Paths          []string       `protobuf:"bytes,10,rep,name=paths,proto3" json:"paths,omitempty"`
	TimeResolution TimeResolution `protobuf:"varint,7,opt,name=time_resolution,json=timeResolution,proto3,enum=com.coralogix.schemastore.v1.TimeResolution" json:"time_resolution,omitempty"`
	ExactPathMatch bool           `protobuf:"varint,8,opt,name=exact_path_match,json=exactPathMatch,proto3" json:"exact_path_match,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *TracingFieldsRequest) Reset() {
	*x = TracingFieldsRequest{}
	mi := &file_com_coralogix_schemastore_v1_tracing_fields_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TracingFieldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TracingFieldsRequest) ProtoMessage() {}

func (x *TracingFieldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_schemastore_v1_tracing_fields_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TracingFieldsRequest.ProtoReflect.Descriptor instead.
func (*TracingFieldsRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_schemastore_v1_tracing_fields_proto_rawDescGZIP(), []int{0}
}

func (x *TracingFieldsRequest) GetTeamIds() []uint32 {
	if x != nil {
		return x.TeamIds
	}
	return nil
}

func (x *TracingFieldsRequest) GetTimeRange() *TimeRange {
	if x != nil {
		return x.TimeRange
	}
	return nil
}

func (x *TracingFieldsRequest) GetResultLimit() int32 {
	if x != nil {
		return x.ResultLimit
	}
	return 0
}

func (x *TracingFieldsRequest) GetLabelFilters() []*LabelFilter {
	if x != nil {
		return x.LabelFilters
	}
	return nil
}

func (x *TracingFieldsRequest) GetSchemaFilters() []*SchemaFilter {
	if x != nil {
		return x.SchemaFilters
	}
	return nil
}

// Deprecated: Marked as deprecated in com/coralogix/schemastore/v1/tracing_fields.proto.
func (x *TracingFieldsRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *TracingFieldsRequest) GetPaths() []string {
	if x != nil {
		return x.Paths
	}
	return nil
}

func (x *TracingFieldsRequest) GetTimeResolution() TimeResolution {
	if x != nil {
		return x.TimeResolution
	}
	return TimeResolution_TIME_RESOLUTION_UNSPECIFIED
}

func (x *TracingFieldsRequest) GetExactPathMatch() bool {
	if x != nil {
		return x.ExactPathMatch
	}
	return false
}

type TracingFieldsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Fields        []*TracingField        `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TracingFieldsResponse) Reset() {
	*x = TracingFieldsResponse{}
	mi := &file_com_coralogix_schemastore_v1_tracing_fields_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TracingFieldsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TracingFieldsResponse) ProtoMessage() {}

func (x *TracingFieldsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_schemastore_v1_tracing_fields_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TracingFieldsResponse.ProtoReflect.Descriptor instead.
func (*TracingFieldsResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_schemastore_v1_tracing_fields_proto_rawDescGZIP(), []int{1}
}

func (x *TracingFieldsResponse) GetFields() []*TracingField {
	if x != nil {
		return x.Fields
	}
	return nil
}

type TracingField struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Path          []string                `protobuf:"bytes,1,rep,name=path,proto3" json:"path,omitempty"`
	Metadata      []*TracingFieldMetadata `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TracingField) Reset() {
	*x = TracingField{}
	mi := &file_com_coralogix_schemastore_v1_tracing_fields_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TracingField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TracingField) ProtoMessage() {}

func (x *TracingField) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_schemastore_v1_tracing_fields_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TracingField.ProtoReflect.Descriptor instead.
func (*TracingField) Descriptor() ([]byte, []int) {
	return file_com_coralogix_schemastore_v1_tracing_fields_proto_rawDescGZIP(), []int{2}
}

func (x *TracingField) GetPath() []string {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *TracingField) GetMetadata() []*TracingFieldMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type TracingFieldMetadata struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	DataType        DataType               `protobuf:"varint,1,opt,name=data_type,json=dataType,proto3,enum=com.coralogix.schemastore.v1.DataType" json:"data_type,omitempty"`
	LogicalDataType string                 `protobuf:"bytes,2,opt,name=logical_data_type,json=logicalDataType,proto3" json:"logical_data_type,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *TracingFieldMetadata) Reset() {
	*x = TracingFieldMetadata{}
	mi := &file_com_coralogix_schemastore_v1_tracing_fields_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TracingFieldMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TracingFieldMetadata) ProtoMessage() {}

func (x *TracingFieldMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_schemastore_v1_tracing_fields_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TracingFieldMetadata.ProtoReflect.Descriptor instead.
func (*TracingFieldMetadata) Descriptor() ([]byte, []int) {
	return file_com_coralogix_schemastore_v1_tracing_fields_proto_rawDescGZIP(), []int{3}
}

func (x *TracingFieldMetadata) GetDataType() DataType {
	if x != nil {
		return x.DataType
	}
	return DataType_DATA_TYPE_UNSPECIFIED
}

func (x *TracingFieldMetadata) GetLogicalDataType() string {
	if x != nil {
		return x.LogicalDataType
	}
	return ""
}

var File_com_coralogix_schemastore_v1_tracing_fields_proto protoreflect.FileDescriptor

var file_com_coralogix_schemastore_v1_tracing_fields_proto_rawDesc = []byte{
	0x0a, 0x31, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x29, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x03, 0x0a,
	0x14, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x73,
	0x12, 0x46, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x4e, 0x0a, 0x0d, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0c, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x51, 0x0a, 0x0e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52,
	0x0d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x16,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x12, 0x55, 0x0a, 0x0f,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x78, 0x61, 0x63, 0x74, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x65,
	0x78, 0x61, 0x63, 0x74, 0x50, 0x61, 0x74, 0x68, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x22, 0x5b, 0x0a,
	0x15, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x72, 0x0a, 0x0c, 0x54, 0x72,
	0x61, 0x63, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x4e,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x87,
	0x01, 0x0a, 0x14, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x43, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x11,
	0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c,
	0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_schemastore_v1_tracing_fields_proto_rawDescOnce sync.Once
	file_com_coralogix_schemastore_v1_tracing_fields_proto_rawDescData = file_com_coralogix_schemastore_v1_tracing_fields_proto_rawDesc
)

func file_com_coralogix_schemastore_v1_tracing_fields_proto_rawDescGZIP() []byte {
	file_com_coralogix_schemastore_v1_tracing_fields_proto_rawDescOnce.Do(func() {
		file_com_coralogix_schemastore_v1_tracing_fields_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_schemastore_v1_tracing_fields_proto_rawDescData)
	})
	return file_com_coralogix_schemastore_v1_tracing_fields_proto_rawDescData
}

var file_com_coralogix_schemastore_v1_tracing_fields_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_com_coralogix_schemastore_v1_tracing_fields_proto_goTypes = []any{
	(*TracingFieldsRequest)(nil),  // 0: com.coralogix.schemastore.v1.TracingFieldsRequest
	(*TracingFieldsResponse)(nil), // 1: com.coralogix.schemastore.v1.TracingFieldsResponse
	(*TracingField)(nil),          // 2: com.coralogix.schemastore.v1.TracingField
	(*TracingFieldMetadata)(nil),  // 3: com.coralogix.schemastore.v1.TracingFieldMetadata
	(*TimeRange)(nil),             // 4: com.coralogix.schemastore.v1.TimeRange
	(*LabelFilter)(nil),           // 5: com.coralogix.schemastore.v1.LabelFilter
	(*SchemaFilter)(nil),          // 6: com.coralogix.schemastore.v1.SchemaFilter
	(TimeResolution)(0),           // 7: com.coralogix.schemastore.v1.TimeResolution
	(DataType)(0),                 // 8: com.coralogix.schemastore.v1.DataType
}
var file_com_coralogix_schemastore_v1_tracing_fields_proto_depIdxs = []int32{
	4, // 0: com.coralogix.schemastore.v1.TracingFieldsRequest.time_range:type_name -> com.coralogix.schemastore.v1.TimeRange
	5, // 1: com.coralogix.schemastore.v1.TracingFieldsRequest.label_filters:type_name -> com.coralogix.schemastore.v1.LabelFilter
	6, // 2: com.coralogix.schemastore.v1.TracingFieldsRequest.schema_filters:type_name -> com.coralogix.schemastore.v1.SchemaFilter
	7, // 3: com.coralogix.schemastore.v1.TracingFieldsRequest.time_resolution:type_name -> com.coralogix.schemastore.v1.TimeResolution
	2, // 4: com.coralogix.schemastore.v1.TracingFieldsResponse.fields:type_name -> com.coralogix.schemastore.v1.TracingField
	3, // 5: com.coralogix.schemastore.v1.TracingField.metadata:type_name -> com.coralogix.schemastore.v1.TracingFieldMetadata
	8, // 6: com.coralogix.schemastore.v1.TracingFieldMetadata.data_type:type_name -> com.coralogix.schemastore.v1.DataType
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_com_coralogix_schemastore_v1_tracing_fields_proto_init() }
func file_com_coralogix_schemastore_v1_tracing_fields_proto_init() {
	if File_com_coralogix_schemastore_v1_tracing_fields_proto != nil {
		return
	}
	file_com_coralogix_schemastore_v1_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_schemastore_v1_tracing_fields_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_schemastore_v1_tracing_fields_proto_goTypes,
		DependencyIndexes: file_com_coralogix_schemastore_v1_tracing_fields_proto_depIdxs,
		MessageInfos:      file_com_coralogix_schemastore_v1_tracing_fields_proto_msgTypes,
	}.Build()
	File_com_coralogix_schemastore_v1_tracing_fields_proto = out.File
	file_com_coralogix_schemastore_v1_tracing_fields_proto_rawDesc = nil
	file_com_coralogix_schemastore_v1_tracing_fields_proto_goTypes = nil
	file_com_coralogix_schemastore_v1_tracing_fields_proto_depIdxs = nil
}
