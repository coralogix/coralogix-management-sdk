// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.1
// source: com/coralogixapis/dashboards/v1/services/dataprime_data_source_service.proto

package v1

import (
	_ "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/common/v1"
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

type SearchDataprimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataprimeQuery    *SerializedDataprimeQuery `protobuf:"bytes,1,opt,name=dataprime_query,json=dataprimeQuery,proto3" json:"dataprime_query,omitempty"`
	DataprimeQueryRaw *DataprimeQuery           `protobuf:"bytes,2,opt,name=dataprime_query_raw,json=dataprimeQueryRaw,proto3" json:"dataprime_query_raw,omitempty"`
	TimeFrame         *TimeFrame                `protobuf:"bytes,3,opt,name=time_frame,json=timeFrame,proto3" json:"time_frame,omitempty"`
	Limit             *wrapperspb.Int32Value    `protobuf:"bytes,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *SearchDataprimeRequest) Reset() {
	*x = SearchDataprimeRequest{}
	mi := &file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchDataprimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchDataprimeRequest) ProtoMessage() {}

func (x *SearchDataprimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchDataprimeRequest.ProtoReflect.Descriptor instead.
func (*SearchDataprimeRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_rawDescGZIP(), []int{0}
}

func (x *SearchDataprimeRequest) GetDataprimeQuery() *SerializedDataprimeQuery {
	if x != nil {
		return x.DataprimeQuery
	}
	return nil
}

func (x *SearchDataprimeRequest) GetDataprimeQueryRaw() *DataprimeQuery {
	if x != nil {
		return x.DataprimeQueryRaw
	}
	return nil
}

func (x *SearchDataprimeRequest) GetTimeFrame() *TimeFrame {
	if x != nil {
		return x.TimeFrame
	}
	return nil
}

func (x *SearchDataprimeRequest) GetLimit() *wrapperspb.Int32Value {
	if x != nil {
		return x.Limit
	}
	return nil
}

type SearchDataprimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*DataprimeResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Meta    map[string]string  `protobuf:"bytes,2,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SearchDataprimeResponse) Reset() {
	*x = SearchDataprimeResponse{}
	mi := &file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchDataprimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchDataprimeResponse) ProtoMessage() {}

func (x *SearchDataprimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchDataprimeResponse.ProtoReflect.Descriptor instead.
func (*SearchDataprimeResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_rawDescGZIP(), []int{1}
}

func (x *SearchDataprimeResponse) GetResults() []*DataprimeResult {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *SearchDataprimeResponse) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

type SearchDataprimeArchiveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataprimeQuery    *SerializedDataprimeQuery `protobuf:"bytes,1,opt,name=dataprime_query,json=dataprimeQuery,proto3" json:"dataprime_query,omitempty"`
	DataprimeQueryRaw *DataprimeQuery           `protobuf:"bytes,2,opt,name=dataprime_query_raw,json=dataprimeQueryRaw,proto3" json:"dataprime_query_raw,omitempty"`
	TimeFrame         *TimeFrame                `protobuf:"bytes,3,opt,name=time_frame,json=timeFrame,proto3" json:"time_frame,omitempty"`
	Limit             *wrapperspb.Int32Value    `protobuf:"bytes,4,opt,name=limit,proto3" json:"limit,omitempty"`
	WidgetId          *wrapperspb.StringValue   `protobuf:"bytes,5,opt,name=widget_id,json=widgetId,proto3" json:"widget_id,omitempty"`
}

func (x *SearchDataprimeArchiveRequest) Reset() {
	*x = SearchDataprimeArchiveRequest{}
	mi := &file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchDataprimeArchiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchDataprimeArchiveRequest) ProtoMessage() {}

func (x *SearchDataprimeArchiveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchDataprimeArchiveRequest.ProtoReflect.Descriptor instead.
func (*SearchDataprimeArchiveRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_rawDescGZIP(), []int{2}
}

func (x *SearchDataprimeArchiveRequest) GetDataprimeQuery() *SerializedDataprimeQuery {
	if x != nil {
		return x.DataprimeQuery
	}
	return nil
}

func (x *SearchDataprimeArchiveRequest) GetDataprimeQueryRaw() *DataprimeQuery {
	if x != nil {
		return x.DataprimeQueryRaw
	}
	return nil
}

func (x *SearchDataprimeArchiveRequest) GetTimeFrame() *TimeFrame {
	if x != nil {
		return x.TimeFrame
	}
	return nil
}

func (x *SearchDataprimeArchiveRequest) GetLimit() *wrapperspb.Int32Value {
	if x != nil {
		return x.Limit
	}
	return nil
}

func (x *SearchDataprimeArchiveRequest) GetWidgetId() *wrapperspb.StringValue {
	if x != nil {
		return x.WidgetId
	}
	return nil
}

type SearchDataprimeArchiveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*DataprimeResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Meta    map[string]string  `protobuf:"bytes,2,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SearchDataprimeArchiveResponse) Reset() {
	*x = SearchDataprimeArchiveResponse{}
	mi := &file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchDataprimeArchiveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchDataprimeArchiveResponse) ProtoMessage() {}

func (x *SearchDataprimeArchiveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchDataprimeArchiveResponse.ProtoReflect.Descriptor instead.
func (*SearchDataprimeArchiveResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_rawDescGZIP(), []int{3}
}

func (x *SearchDataprimeArchiveResponse) GetResults() []*DataprimeResult {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *SearchDataprimeArchiveResponse) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

var File_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto protoreflect.FileDescriptor

var file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_rawDesc = []byte{
	0x0a, 0x4c, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70,
	0x72, 0x69, 0x6d, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x27, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x3d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x69, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x32, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x37, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x02,
	0x0a, 0x16, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x69, 0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61,
	0x70, 0x72, 0x69, 0x6d, 0x65, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x40, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x66, 0x0a, 0x13, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65,
	0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x72, 0x61, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x69, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x11, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x69, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x61, 0x77, 0x12, 0x50, 0x0a, 0x0a, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x61,
	0x6d, 0x65, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x22, 0x86, 0x02, 0x0a, 0x17, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x70,
	0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12,
	0x5f, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4b, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44,
	0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x1a, 0x37, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb2, 0x03, 0x0a, 0x1d, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x41, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x69, 0x0a, 0x0f, 0x64,
	0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d,
	0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d,
	0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x66, 0x0a, 0x13, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x69, 0x6d, 0x65, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x72, 0x61, 0x77, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x11, 0x64, 0x61, 0x74,
	0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x61, 0x77, 0x12, 0x50,
	0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x46, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x12, 0x31, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x39, 0x0a, 0x09, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x49, 0x64, 0x22, 0x94,
	0x02, 0x0a, 0x1e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69,
	0x6d, 0x65, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x51, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x70, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x12, 0x66, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x52, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x41, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09,
	0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x9b, 0x03, 0x0a, 0x1a, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x69, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0xae, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44,
	0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x41, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61,
	0x70, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0xc2,
	0xb8, 0x02, 0x12, 0x0a, 0x10, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x20, 0x44, 0x61, 0x74, 0x61,
	0x70, 0x72, 0x69, 0x6d, 0x65, 0x12, 0xcb, 0x01, 0x0a, 0x16, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x12, 0x47, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x41, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x48, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x70,
	0x72, 0x69, 0x6d, 0x65, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1e, 0xc2, 0xb8, 0x02, 0x1a, 0x0a, 0x18, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x20, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x20, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_rawDescOnce sync.Once
	file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_rawDescData = file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_rawDesc
)

func file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_rawDescData)
	})
	return file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_rawDescData
}

var file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_goTypes = []any{
	(*SearchDataprimeRequest)(nil),         // 0: com.coralogixapis.dashboards.v1.services.SearchDataprimeRequest
	(*SearchDataprimeResponse)(nil),        // 1: com.coralogixapis.dashboards.v1.services.SearchDataprimeResponse
	(*SearchDataprimeArchiveRequest)(nil),  // 2: com.coralogixapis.dashboards.v1.services.SearchDataprimeArchiveRequest
	(*SearchDataprimeArchiveResponse)(nil), // 3: com.coralogixapis.dashboards.v1.services.SearchDataprimeArchiveResponse
	nil,                                    // 4: com.coralogixapis.dashboards.v1.services.SearchDataprimeResponse.MetaEntry
	nil,                                    // 5: com.coralogixapis.dashboards.v1.services.SearchDataprimeArchiveResponse.MetaEntry
	(*SerializedDataprimeQuery)(nil),       // 6: com.coralogixapis.dashboards.v1.common.SerializedDataprimeQuery
	(*DataprimeQuery)(nil),                 // 7: com.coralogixapis.dashboards.v1.common.DataprimeQuery
	(*TimeFrame)(nil),                      // 8: com.coralogixapis.dashboards.v1.common.TimeFrame
	(*wrapperspb.Int32Value)(nil),          // 9: google.protobuf.Int32Value
	(*DataprimeResult)(nil),                // 10: com.coralogixapis.dashboards.v1.common.DataprimeResult
	(*wrapperspb.StringValue)(nil),         // 11: google.protobuf.StringValue
}
var file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_depIdxs = []int32{
	6,  // 0: com.coralogixapis.dashboards.v1.services.SearchDataprimeRequest.dataprime_query:type_name -> com.coralogixapis.dashboards.v1.common.SerializedDataprimeQuery
	7,  // 1: com.coralogixapis.dashboards.v1.services.SearchDataprimeRequest.dataprime_query_raw:type_name -> com.coralogixapis.dashboards.v1.common.DataprimeQuery
	8,  // 2: com.coralogixapis.dashboards.v1.services.SearchDataprimeRequest.time_frame:type_name -> com.coralogixapis.dashboards.v1.common.TimeFrame
	9,  // 3: com.coralogixapis.dashboards.v1.services.SearchDataprimeRequest.limit:type_name -> google.protobuf.Int32Value
	10, // 4: com.coralogixapis.dashboards.v1.services.SearchDataprimeResponse.results:type_name -> com.coralogixapis.dashboards.v1.common.DataprimeResult
	4,  // 5: com.coralogixapis.dashboards.v1.services.SearchDataprimeResponse.meta:type_name -> com.coralogixapis.dashboards.v1.services.SearchDataprimeResponse.MetaEntry
	6,  // 6: com.coralogixapis.dashboards.v1.services.SearchDataprimeArchiveRequest.dataprime_query:type_name -> com.coralogixapis.dashboards.v1.common.SerializedDataprimeQuery
	7,  // 7: com.coralogixapis.dashboards.v1.services.SearchDataprimeArchiveRequest.dataprime_query_raw:type_name -> com.coralogixapis.dashboards.v1.common.DataprimeQuery
	8,  // 8: com.coralogixapis.dashboards.v1.services.SearchDataprimeArchiveRequest.time_frame:type_name -> com.coralogixapis.dashboards.v1.common.TimeFrame
	9,  // 9: com.coralogixapis.dashboards.v1.services.SearchDataprimeArchiveRequest.limit:type_name -> google.protobuf.Int32Value
	11, // 10: com.coralogixapis.dashboards.v1.services.SearchDataprimeArchiveRequest.widget_id:type_name -> google.protobuf.StringValue
	10, // 11: com.coralogixapis.dashboards.v1.services.SearchDataprimeArchiveResponse.results:type_name -> com.coralogixapis.dashboards.v1.common.DataprimeResult
	5,  // 12: com.coralogixapis.dashboards.v1.services.SearchDataprimeArchiveResponse.meta:type_name -> com.coralogixapis.dashboards.v1.services.SearchDataprimeArchiveResponse.MetaEntry
	0,  // 13: com.coralogixapis.dashboards.v1.services.DataprimeDataSourceService.SearchDataprime:input_type -> com.coralogixapis.dashboards.v1.services.SearchDataprimeRequest
	2,  // 14: com.coralogixapis.dashboards.v1.services.DataprimeDataSourceService.SearchDataprimeArchive:input_type -> com.coralogixapis.dashboards.v1.services.SearchDataprimeArchiveRequest
	1,  // 15: com.coralogixapis.dashboards.v1.services.DataprimeDataSourceService.SearchDataprime:output_type -> com.coralogixapis.dashboards.v1.services.SearchDataprimeResponse
	3,  // 16: com.coralogixapis.dashboards.v1.services.DataprimeDataSourceService.SearchDataprimeArchive:output_type -> com.coralogixapis.dashboards.v1.services.SearchDataprimeArchiveResponse
	15, // [15:17] is the sub-list for method output_type
	13, // [13:15] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_init() }
func file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_init() {
	if File_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto != nil {
		return
	}
	file_com_coralogixapis_dashboards_v1_common_dataprime_result_proto_init()
	file_com_coralogixapis_dashboards_v1_common_query_proto_init()
	file_com_coralogixapis_dashboards_v1_common_time_frame_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto = out.File
	file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_rawDesc = nil
	file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_goTypes = nil
	file_com_coralogixapis_dashboards_v1_services_dataprime_data_source_service_proto_depIdxs = nil
}
