// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.26.1
// source: com/coralogix/global_mapping/v1/measurements_internal_service.proto

package v1

import (
	_ "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/common/v1"
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

type UpsertGlobalMeasurementsInternalRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Measurements  []*Measurement         `protobuf:"bytes,1,rep,name=measurements,proto3" json:"measurements,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpsertGlobalMeasurementsInternalRequest) Reset() {
	*x = UpsertGlobalMeasurementsInternalRequest{}
	mi := &file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertGlobalMeasurementsInternalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertGlobalMeasurementsInternalRequest) ProtoMessage() {}

func (x *UpsertGlobalMeasurementsInternalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertGlobalMeasurementsInternalRequest.ProtoReflect.Descriptor instead.
func (*UpsertGlobalMeasurementsInternalRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_rawDescGZIP(), []int{0}
}

func (x *UpsertGlobalMeasurementsInternalRequest) GetMeasurements() []*Measurement {
	if x != nil {
		return x.Measurements
	}
	return nil
}

type UpsertGlobalMeasurementsInternalResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Measurements  []*Measurement         `protobuf:"bytes,1,rep,name=measurements,proto3" json:"measurements,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpsertGlobalMeasurementsInternalResponse) Reset() {
	*x = UpsertGlobalMeasurementsInternalResponse{}
	mi := &file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertGlobalMeasurementsInternalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertGlobalMeasurementsInternalResponse) ProtoMessage() {}

func (x *UpsertGlobalMeasurementsInternalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertGlobalMeasurementsInternalResponse.ProtoReflect.Descriptor instead.
func (*UpsertGlobalMeasurementsInternalResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_rawDescGZIP(), []int{1}
}

func (x *UpsertGlobalMeasurementsInternalResponse) GetMeasurements() []*Measurement {
	if x != nil {
		return x.Measurements
	}
	return nil
}

type RemoveGlobalMeasurementsInternalRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	MeasurementIds []string               `protobuf:"bytes,1,rep,name=measurement_ids,json=measurementIds,proto3" json:"measurement_ids,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *RemoveGlobalMeasurementsInternalRequest) Reset() {
	*x = RemoveGlobalMeasurementsInternalRequest{}
	mi := &file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveGlobalMeasurementsInternalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveGlobalMeasurementsInternalRequest) ProtoMessage() {}

func (x *RemoveGlobalMeasurementsInternalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveGlobalMeasurementsInternalRequest.ProtoReflect.Descriptor instead.
func (*RemoveGlobalMeasurementsInternalRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveGlobalMeasurementsInternalRequest) GetMeasurementIds() []string {
	if x != nil {
		return x.MeasurementIds
	}
	return nil
}

type RemoveGlobalMeasurementsInternalResponse struct {
	state                       protoimpl.MessageState `protogen:"open.v1"`
	NumberOfDeletedMeasurements *wrapperspb.Int32Value `protobuf:"bytes,1,opt,name=number_of_deleted_measurements,json=numberOfDeletedMeasurements,proto3" json:"number_of_deleted_measurements,omitempty"`
	unknownFields               protoimpl.UnknownFields
	sizeCache                   protoimpl.SizeCache
}

func (x *RemoveGlobalMeasurementsInternalResponse) Reset() {
	*x = RemoveGlobalMeasurementsInternalResponse{}
	mi := &file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveGlobalMeasurementsInternalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveGlobalMeasurementsInternalResponse) ProtoMessage() {}

func (x *RemoveGlobalMeasurementsInternalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveGlobalMeasurementsInternalResponse.ProtoReflect.Descriptor instead.
func (*RemoveGlobalMeasurementsInternalResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveGlobalMeasurementsInternalResponse) GetNumberOfDeletedMeasurements() *wrapperspb.Int32Value {
	if x != nil {
		return x.NumberOfDeletedMeasurements
	}
	return nil
}

type GetMeasurementsInternalRequest struct {
	state            protoimpl.MessageState    `protogen:"open.v1"`
	MeasurementNames []*wrapperspb.StringValue `protobuf:"bytes,1,rep,name=measurement_names,json=measurementNames,proto3" json:"measurement_names,omitempty"`
	DataSourceTypes  []DataSourceType          `protobuf:"varint,2,rep,packed,name=data_source_types,json=dataSourceTypes,proto3,enum=com.coralogix.global_mapping.v1.DataSourceType" json:"data_source_types,omitempty"`
	Labels           []*wrapperspb.StringValue `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty"`
	DataSources      []*DataSource             `protobuf:"bytes,5,rep,name=data_sources,json=dataSources,proto3" json:"data_sources,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *GetMeasurementsInternalRequest) Reset() {
	*x = GetMeasurementsInternalRequest{}
	mi := &file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMeasurementsInternalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMeasurementsInternalRequest) ProtoMessage() {}

func (x *GetMeasurementsInternalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMeasurementsInternalRequest.ProtoReflect.Descriptor instead.
func (*GetMeasurementsInternalRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetMeasurementsInternalRequest) GetMeasurementNames() []*wrapperspb.StringValue {
	if x != nil {
		return x.MeasurementNames
	}
	return nil
}

func (x *GetMeasurementsInternalRequest) GetDataSourceTypes() []DataSourceType {
	if x != nil {
		return x.DataSourceTypes
	}
	return nil
}

func (x *GetMeasurementsInternalRequest) GetLabels() []*wrapperspb.StringValue {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *GetMeasurementsInternalRequest) GetDataSources() []*DataSource {
	if x != nil {
		return x.DataSources
	}
	return nil
}

type GetMeasurementsInternalResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Measurements  []*Measurement         `protobuf:"bytes,1,rep,name=measurements,proto3" json:"measurements,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetMeasurementsInternalResponse) Reset() {
	*x = GetMeasurementsInternalResponse{}
	mi := &file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMeasurementsInternalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMeasurementsInternalResponse) ProtoMessage() {}

func (x *GetMeasurementsInternalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMeasurementsInternalResponse.ProtoReflect.Descriptor instead.
func (*GetMeasurementsInternalResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetMeasurementsInternalResponse) GetMeasurements() []*Measurement {
	if x != nil {
		return x.Measurements
	}
	return nil
}

type UpsertMeasurementsInternalRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Measurements  []*Measurement         `protobuf:"bytes,1,rep,name=measurements,proto3" json:"measurements,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpsertMeasurementsInternalRequest) Reset() {
	*x = UpsertMeasurementsInternalRequest{}
	mi := &file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertMeasurementsInternalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertMeasurementsInternalRequest) ProtoMessage() {}

func (x *UpsertMeasurementsInternalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertMeasurementsInternalRequest.ProtoReflect.Descriptor instead.
func (*UpsertMeasurementsInternalRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpsertMeasurementsInternalRequest) GetMeasurements() []*Measurement {
	if x != nil {
		return x.Measurements
	}
	return nil
}

type UpsertCompanyProvidersInternalRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Deprecated: Marked as deprecated in com/coralogix/global_mapping/v1/measurements_internal_service.proto.
	Providers     []*wrapperspb.StringValue `protobuf:"bytes,1,rep,name=providers,proto3" json:"providers,omitempty"`
	DataSources   []*DataSource             `protobuf:"bytes,2,rep,name=data_sources,json=dataSources,proto3" json:"data_sources,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpsertCompanyProvidersInternalRequest) Reset() {
	*x = UpsertCompanyProvidersInternalRequest{}
	mi := &file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertCompanyProvidersInternalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertCompanyProvidersInternalRequest) ProtoMessage() {}

func (x *UpsertCompanyProvidersInternalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertCompanyProvidersInternalRequest.ProtoReflect.Descriptor instead.
func (*UpsertCompanyProvidersInternalRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_rawDescGZIP(), []int{7}
}

// Deprecated: Marked as deprecated in com/coralogix/global_mapping/v1/measurements_internal_service.proto.
func (x *UpsertCompanyProvidersInternalRequest) GetProviders() []*wrapperspb.StringValue {
	if x != nil {
		return x.Providers
	}
	return nil
}

func (x *UpsertCompanyProvidersInternalRequest) GetDataSources() []*DataSource {
	if x != nil {
		return x.DataSources
	}
	return nil
}

type UpsertCompanyProvidersInternalResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Deprecated: Marked as deprecated in com/coralogix/global_mapping/v1/measurements_internal_service.proto.
	Providers     []*wrapperspb.StringValue `protobuf:"bytes,1,rep,name=providers,proto3" json:"providers,omitempty"`
	DataSources   []*DataSource             `protobuf:"bytes,2,rep,name=data_sources,json=dataSources,proto3" json:"data_sources,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpsertCompanyProvidersInternalResponse) Reset() {
	*x = UpsertCompanyProvidersInternalResponse{}
	mi := &file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertCompanyProvidersInternalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertCompanyProvidersInternalResponse) ProtoMessage() {}

func (x *UpsertCompanyProvidersInternalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertCompanyProvidersInternalResponse.ProtoReflect.Descriptor instead.
func (*UpsertCompanyProvidersInternalResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_rawDescGZIP(), []int{8}
}

// Deprecated: Marked as deprecated in com/coralogix/global_mapping/v1/measurements_internal_service.proto.
func (x *UpsertCompanyProvidersInternalResponse) GetProviders() []*wrapperspb.StringValue {
	if x != nil {
		return x.Providers
	}
	return nil
}

func (x *UpsertCompanyProvidersInternalResponse) GetDataSources() []*DataSource {
	if x != nil {
		return x.DataSources
	}
	return nil
}

var File_com_coralogix_global_mapping_v1_measurements_internal_service_proto protoreflect.FileDescriptor

const file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_rawDesc = "" +
	"\n" +
	"Ccom/coralogix/global_mapping/v1/measurements_internal_service.proto\x12\x1fcom.coralogix.global_mapping.v1\x1a\x1egoogle/protobuf/wrappers.proto\x1a1com/coralogix/global_mapping/v1/measurement.proto\x1a'com/coralogix/common/v1/audit_log.proto\x1a6com/coralogix/global_mapping/v1/data_source_type.proto\x1a1com/coralogix/global_mapping/v1/data_source.proto\"{\n" +
	"'UpsertGlobalMeasurementsInternalRequest\x12P\n" +
	"\fmeasurements\x18\x01 \x03(\v2,.com.coralogix.global_mapping.v1.MeasurementR\fmeasurements\"|\n" +
	"(UpsertGlobalMeasurementsInternalResponse\x12P\n" +
	"\fmeasurements\x18\x01 \x03(\v2,.com.coralogix.global_mapping.v1.MeasurementR\fmeasurements\"R\n" +
	"'RemoveGlobalMeasurementsInternalRequest\x12'\n" +
	"\x0fmeasurement_ids\x18\x01 \x03(\tR\x0emeasurementIds\"\x8c\x01\n" +
	"(RemoveGlobalMeasurementsInternalResponse\x12`\n" +
	"\x1enumber_of_deleted_measurements\x18\x01 \x01(\v2\x1b.google.protobuf.Int32ValueR\x1bnumberOfDeletedMeasurements\"\xeb\x02\n" +
	"\x1eGetMeasurementsInternalRequest\x12I\n" +
	"\x11measurement_names\x18\x01 \x03(\v2\x1c.google.protobuf.StringValueR\x10measurementNames\x12[\n" +
	"\x11data_source_types\x18\x02 \x03(\x0e2/.com.coralogix.global_mapping.v1.DataSourceTypeR\x0fdataSourceTypes\x124\n" +
	"\x06labels\x18\x04 \x03(\v2\x1c.google.protobuf.StringValueR\x06labels\x12N\n" +
	"\fdata_sources\x18\x05 \x03(\v2+.com.coralogix.global_mapping.v1.DataSourceR\vdataSourcesJ\x04\b\x03\x10\x04R\x15data_source_providers\"s\n" +
	"\x1fGetMeasurementsInternalResponse\x12P\n" +
	"\fmeasurements\x18\x01 \x03(\v2,.com.coralogix.global_mapping.v1.MeasurementR\fmeasurements\"u\n" +
	"!UpsertMeasurementsInternalRequest\x12P\n" +
	"\fmeasurements\x18\x01 \x03(\v2,.com.coralogix.global_mapping.v1.MeasurementR\fmeasurements\"\xb7\x01\n" +
	"%UpsertCompanyProvidersInternalRequest\x12>\n" +
	"\tproviders\x18\x01 \x03(\v2\x1c.google.protobuf.StringValueB\x02\x18\x01R\tproviders\x12N\n" +
	"\fdata_sources\x18\x02 \x03(\v2+.com.coralogix.global_mapping.v1.DataSourceR\vdataSources\"\xb8\x01\n" +
	"&UpsertCompanyProvidersInternalResponse\x12>\n" +
	"\tproviders\x18\x01 \x03(\v2\x1c.google.protobuf.StringValueB\x02\x18\x01R\tproviders\x12N\n" +
	"\fdata_sources\x18\x02 \x03(\v2+.com.coralogix.global_mapping.v1.DataSourceR\vdataSources2\xfb\x06\n" +
	"\x1bMeasurementsInternalService\x12\xe2\x01\n" +
	" UpsertGlobalMeasurementsInternal\x12H.com.coralogix.global_mapping.v1.UpsertGlobalMeasurementsInternalRequest\x1aI.com.coralogix.global_mapping.v1.UpsertGlobalMeasurementsInternalResponse\")¸\x02%\n" +
	"#Upsert Global Measurements Internal\x12\xe2\x01\n" +
	" RemoveGlobalMeasurementsInternal\x12H.com.coralogix.global_mapping.v1.RemoveGlobalMeasurementsInternalRequest\x1aI.com.coralogix.global_mapping.v1.RemoveGlobalMeasurementsInternalResponse\")¸\x02%\n" +
	"#Remove Global Measurements Internal\x12\xda\x01\n" +
	"\x1eUpsertCompanyProvidersInternal\x12F.com.coralogix.global_mapping.v1.UpsertCompanyProvidersInternalRequest\x1aG.com.coralogix.global_mapping.v1.UpsertCompanyProvidersInternalResponse\"'¸\x02#\n" +
	"!Upsert Company Providers Internal\x12\xb4\x01\n" +
	"\x17GetMeasurementsInternal\x12?.com.coralogix.global_mapping.v1.GetMeasurementsInternalRequest\x1a@.com.coralogix.global_mapping.v1.GetMeasurementsInternalResponse\"\x16¸\x02\x12\n" +
	"\x10Get Measurementsb\x06proto3"

var (
	file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_rawDescOnce sync.Once
	file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_rawDescData []byte
)

func file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_rawDescGZIP() []byte {
	file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_rawDescOnce.Do(func() {
		file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_rawDesc), len(file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_rawDesc)))
	})
	return file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_rawDescData
}

var file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_goTypes = []any{
	(*UpsertGlobalMeasurementsInternalRequest)(nil),  // 0: com.coralogix.global_mapping.v1.UpsertGlobalMeasurementsInternalRequest
	(*UpsertGlobalMeasurementsInternalResponse)(nil), // 1: com.coralogix.global_mapping.v1.UpsertGlobalMeasurementsInternalResponse
	(*RemoveGlobalMeasurementsInternalRequest)(nil),  // 2: com.coralogix.global_mapping.v1.RemoveGlobalMeasurementsInternalRequest
	(*RemoveGlobalMeasurementsInternalResponse)(nil), // 3: com.coralogix.global_mapping.v1.RemoveGlobalMeasurementsInternalResponse
	(*GetMeasurementsInternalRequest)(nil),           // 4: com.coralogix.global_mapping.v1.GetMeasurementsInternalRequest
	(*GetMeasurementsInternalResponse)(nil),          // 5: com.coralogix.global_mapping.v1.GetMeasurementsInternalResponse
	(*UpsertMeasurementsInternalRequest)(nil),        // 6: com.coralogix.global_mapping.v1.UpsertMeasurementsInternalRequest
	(*UpsertCompanyProvidersInternalRequest)(nil),    // 7: com.coralogix.global_mapping.v1.UpsertCompanyProvidersInternalRequest
	(*UpsertCompanyProvidersInternalResponse)(nil),   // 8: com.coralogix.global_mapping.v1.UpsertCompanyProvidersInternalResponse
	(*Measurement)(nil),                              // 9: com.coralogix.global_mapping.v1.Measurement
	(*wrapperspb.Int32Value)(nil),                    // 10: google.protobuf.Int32Value
	(*wrapperspb.StringValue)(nil),                   // 11: google.protobuf.StringValue
	(DataSourceType)(0),                              // 12: com.coralogix.global_mapping.v1.DataSourceType
	(*DataSource)(nil),                               // 13: com.coralogix.global_mapping.v1.DataSource
}
var file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_depIdxs = []int32{
	9,  // 0: com.coralogix.global_mapping.v1.UpsertGlobalMeasurementsInternalRequest.measurements:type_name -> com.coralogix.global_mapping.v1.Measurement
	9,  // 1: com.coralogix.global_mapping.v1.UpsertGlobalMeasurementsInternalResponse.measurements:type_name -> com.coralogix.global_mapping.v1.Measurement
	10, // 2: com.coralogix.global_mapping.v1.RemoveGlobalMeasurementsInternalResponse.number_of_deleted_measurements:type_name -> google.protobuf.Int32Value
	11, // 3: com.coralogix.global_mapping.v1.GetMeasurementsInternalRequest.measurement_names:type_name -> google.protobuf.StringValue
	12, // 4: com.coralogix.global_mapping.v1.GetMeasurementsInternalRequest.data_source_types:type_name -> com.coralogix.global_mapping.v1.DataSourceType
	11, // 5: com.coralogix.global_mapping.v1.GetMeasurementsInternalRequest.labels:type_name -> google.protobuf.StringValue
	13, // 6: com.coralogix.global_mapping.v1.GetMeasurementsInternalRequest.data_sources:type_name -> com.coralogix.global_mapping.v1.DataSource
	9,  // 7: com.coralogix.global_mapping.v1.GetMeasurementsInternalResponse.measurements:type_name -> com.coralogix.global_mapping.v1.Measurement
	9,  // 8: com.coralogix.global_mapping.v1.UpsertMeasurementsInternalRequest.measurements:type_name -> com.coralogix.global_mapping.v1.Measurement
	11, // 9: com.coralogix.global_mapping.v1.UpsertCompanyProvidersInternalRequest.providers:type_name -> google.protobuf.StringValue
	13, // 10: com.coralogix.global_mapping.v1.UpsertCompanyProvidersInternalRequest.data_sources:type_name -> com.coralogix.global_mapping.v1.DataSource
	11, // 11: com.coralogix.global_mapping.v1.UpsertCompanyProvidersInternalResponse.providers:type_name -> google.protobuf.StringValue
	13, // 12: com.coralogix.global_mapping.v1.UpsertCompanyProvidersInternalResponse.data_sources:type_name -> com.coralogix.global_mapping.v1.DataSource
	0,  // 13: com.coralogix.global_mapping.v1.MeasurementsInternalService.UpsertGlobalMeasurementsInternal:input_type -> com.coralogix.global_mapping.v1.UpsertGlobalMeasurementsInternalRequest
	2,  // 14: com.coralogix.global_mapping.v1.MeasurementsInternalService.RemoveGlobalMeasurementsInternal:input_type -> com.coralogix.global_mapping.v1.RemoveGlobalMeasurementsInternalRequest
	7,  // 15: com.coralogix.global_mapping.v1.MeasurementsInternalService.UpsertCompanyProvidersInternal:input_type -> com.coralogix.global_mapping.v1.UpsertCompanyProvidersInternalRequest
	4,  // 16: com.coralogix.global_mapping.v1.MeasurementsInternalService.GetMeasurementsInternal:input_type -> com.coralogix.global_mapping.v1.GetMeasurementsInternalRequest
	1,  // 17: com.coralogix.global_mapping.v1.MeasurementsInternalService.UpsertGlobalMeasurementsInternal:output_type -> com.coralogix.global_mapping.v1.UpsertGlobalMeasurementsInternalResponse
	3,  // 18: com.coralogix.global_mapping.v1.MeasurementsInternalService.RemoveGlobalMeasurementsInternal:output_type -> com.coralogix.global_mapping.v1.RemoveGlobalMeasurementsInternalResponse
	8,  // 19: com.coralogix.global_mapping.v1.MeasurementsInternalService.UpsertCompanyProvidersInternal:output_type -> com.coralogix.global_mapping.v1.UpsertCompanyProvidersInternalResponse
	5,  // 20: com.coralogix.global_mapping.v1.MeasurementsInternalService.GetMeasurementsInternal:output_type -> com.coralogix.global_mapping.v1.GetMeasurementsInternalResponse
	17, // [17:21] is the sub-list for method output_type
	13, // [13:17] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_init() }
func file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_init() {
	if File_com_coralogix_global_mapping_v1_measurements_internal_service_proto != nil {
		return
	}
	file_com_coralogix_global_mapping_v1_measurement_proto_init()
	file_com_coralogix_global_mapping_v1_data_source_type_proto_init()
	file_com_coralogix_global_mapping_v1_data_source_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_rawDesc), len(file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_goTypes,
		DependencyIndexes: file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_depIdxs,
		MessageInfos:      file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_msgTypes,
	}.Build()
	File_com_coralogix_global_mapping_v1_measurements_internal_service_proto = out.File
	file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_goTypes = nil
	file_com_coralogix_global_mapping_v1_measurements_internal_service_proto_depIdxs = nil
}
