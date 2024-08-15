// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: com/coralogix/datausage/v1/daily_usage_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DailyUsageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to RequestTime:
	//
	//	*DailyUsageRequest_Range
	//	*DailyUsageRequest_DateRange
	RequestTime isDailyUsageRequest_RequestTime `protobuf_oneof:"request_time"`
}

func (x *DailyUsageRequest) Reset() {
	*x = DailyUsageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_datausage_v1_daily_usage_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DailyUsageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DailyUsageRequest) ProtoMessage() {}

func (x *DailyUsageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_datausage_v1_daily_usage_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DailyUsageRequest.ProtoReflect.Descriptor instead.
func (*DailyUsageRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_datausage_v1_daily_usage_service_proto_rawDescGZIP(), []int{0}
}

func (m *DailyUsageRequest) GetRequestTime() isDailyUsageRequest_RequestTime {
	if m != nil {
		return m.RequestTime
	}
	return nil
}

func (x *DailyUsageRequest) GetRange() Range {
	if x, ok := x.GetRequestTime().(*DailyUsageRequest_Range); ok {
		return x.Range
	}
	return Range_RANGE_UNSPECIFIED
}

func (x *DailyUsageRequest) GetDateRange() *DateRange {
	if x, ok := x.GetRequestTime().(*DailyUsageRequest_DateRange); ok {
		return x.DateRange
	}
	return nil
}

type isDailyUsageRequest_RequestTime interface {
	isDailyUsageRequest_RequestTime()
}

type DailyUsageRequest_Range struct {
	Range Range `protobuf:"varint,1,opt,name=range,proto3,enum=com.coralogix.datausage.v1.Range,oneof"`
}

type DailyUsageRequest_DateRange struct {
	DateRange *DateRange `protobuf:"bytes,2,opt,name=date_range,json=dateRange,proto3,oneof"`
}

func (*DailyUsageRequest_Range) isDailyUsageRequest_RequestTime() {}

func (*DailyUsageRequest_DateRange) isDailyUsageRequest_RequestTime() {}

type DailyUsageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response []*DetailedDailyUsage `protobuf:"bytes,1,rep,name=response,proto3" json:"response,omitempty"`
}

func (x *DailyUsageResponse) Reset() {
	*x = DailyUsageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_datausage_v1_daily_usage_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DailyUsageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DailyUsageResponse) ProtoMessage() {}

func (x *DailyUsageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_datausage_v1_daily_usage_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DailyUsageResponse.ProtoReflect.Descriptor instead.
func (*DailyUsageResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_datausage_v1_daily_usage_service_proto_rawDescGZIP(), []int{1}
}

func (x *DailyUsageResponse) GetResponse() []*DetailedDailyUsage {
	if x != nil {
		return x.Response
	}
	return nil
}

type OverageDailyUsageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to RequestTime:
	//
	//	*OverageDailyUsageRequest_Range
	//	*OverageDailyUsageRequest_DateRange
	RequestTime isOverageDailyUsageRequest_RequestTime `protobuf_oneof:"request_time"`
}

func (x *OverageDailyUsageRequest) Reset() {
	*x = OverageDailyUsageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_datausage_v1_daily_usage_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OverageDailyUsageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverageDailyUsageRequest) ProtoMessage() {}

func (x *OverageDailyUsageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_datausage_v1_daily_usage_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverageDailyUsageRequest.ProtoReflect.Descriptor instead.
func (*OverageDailyUsageRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_datausage_v1_daily_usage_service_proto_rawDescGZIP(), []int{2}
}

func (m *OverageDailyUsageRequest) GetRequestTime() isOverageDailyUsageRequest_RequestTime {
	if m != nil {
		return m.RequestTime
	}
	return nil
}

func (x *OverageDailyUsageRequest) GetRange() Range {
	if x, ok := x.GetRequestTime().(*OverageDailyUsageRequest_Range); ok {
		return x.Range
	}
	return Range_RANGE_UNSPECIFIED
}

func (x *OverageDailyUsageRequest) GetDateRange() *DateRange {
	if x, ok := x.GetRequestTime().(*OverageDailyUsageRequest_DateRange); ok {
		return x.DateRange
	}
	return nil
}

type isOverageDailyUsageRequest_RequestTime interface {
	isOverageDailyUsageRequest_RequestTime()
}

type OverageDailyUsageRequest_Range struct {
	Range Range `protobuf:"varint,1,opt,name=range,proto3,enum=com.coralogix.datausage.v1.Range,oneof"`
}

type OverageDailyUsageRequest_DateRange struct {
	DateRange *DateRange `protobuf:"bytes,2,opt,name=date_range,json=dateRange,proto3,oneof"`
}

func (*OverageDailyUsageRequest_Range) isOverageDailyUsageRequest_RequestTime() {}

func (*OverageDailyUsageRequest_DateRange) isOverageDailyUsageRequest_RequestTime() {}

type OverageDailyUsageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OverageDetail []*OverageDailyReport `protobuf:"bytes,1,rep,name=overage_detail,json=overageDetail,proto3" json:"overage_detail,omitempty"`
}

func (x *OverageDailyUsageResponse) Reset() {
	*x = OverageDailyUsageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_datausage_v1_daily_usage_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OverageDailyUsageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverageDailyUsageResponse) ProtoMessage() {}

func (x *OverageDailyUsageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_datausage_v1_daily_usage_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverageDailyUsageResponse.ProtoReflect.Descriptor instead.
func (*OverageDailyUsageResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_datausage_v1_daily_usage_service_proto_rawDescGZIP(), []int{3}
}

func (x *OverageDailyUsageResponse) GetOverageDetail() []*OverageDailyReport {
	if x != nil {
		return x.OverageDetail
	}
	return nil
}

var File_com_coralogix_datausage_v1_daily_usage_service_proto protoreflect.FileDescriptor

var file_com_coralogix_datausage_v1_daily_usage_service_proto_rawDesc = []byte{
	0x0a, 0x34, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x69,
	0x6c, 0x79, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x2a, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x69, 0x6c, 0x79,
	0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x11, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x05, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x75,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x00, 0x52,
	0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x46, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x75,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x48, 0x00, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x0e,
	0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x60,
	0x0a, 0x12, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x75, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x44, 0x61, 0x69, 0x6c,
	0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0xad, 0x01, 0x0a, 0x18, 0x4f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x44, 0x61, 0x69, 0x6c,
	0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a,
	0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x48,
	0x00, 0x52, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x46, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x48, 0x00, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x42, 0x0e, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x22, 0x72, 0x0a, 0x19, 0x4f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x44, 0x61, 0x69, 0x6c, 0x79,
	0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a,
	0x0e, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x0d, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x32, 0xbd, 0x03, 0x0a, 0x11, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xc7, 0x01, 0x0a, 0x0a, 0x44,
	0x61, 0x69, 0x6c, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x75, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5a, 0xc2, 0xb8, 0x02, 0x27, 0x0a, 0x25,
	0x47, 0x65, 0x74, 0x20, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x75, 0x73, 0x61, 0x67, 0x65, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x65, 0x61, 0x63, 0x68, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x20, 0x64, 0x61, 0x79, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x3a, 0x01, 0x2a, 0x22, 0x24,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67,
	0x2f, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x12, 0xdd, 0x01, 0x0a, 0x11, 0x4f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65,
	0x44, 0x61, 0x69, 0x6c, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x34, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x75,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x44,
	0x61, 0x69, 0x6c, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x76,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5b, 0xc2, 0xb8, 0x02, 0x29, 0x0a, 0x27, 0x47,
	0x65, 0x74, 0x20, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x20,
	0x66, 0x6f, 0x72, 0x20, 0x65, 0x61, 0x63, 0x68, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x20, 0x64, 0x61, 0x79, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x3a, 0x01, 0x2a, 0x22,
	0x23, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e,
	0x67, 0x2f, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6f, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_datausage_v1_daily_usage_service_proto_rawDescOnce sync.Once
	file_com_coralogix_datausage_v1_daily_usage_service_proto_rawDescData = file_com_coralogix_datausage_v1_daily_usage_service_proto_rawDesc
)

func file_com_coralogix_datausage_v1_daily_usage_service_proto_rawDescGZIP() []byte {
	file_com_coralogix_datausage_v1_daily_usage_service_proto_rawDescOnce.Do(func() {
		file_com_coralogix_datausage_v1_daily_usage_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_datausage_v1_daily_usage_service_proto_rawDescData)
	})
	return file_com_coralogix_datausage_v1_daily_usage_service_proto_rawDescData
}

var file_com_coralogix_datausage_v1_daily_usage_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_com_coralogix_datausage_v1_daily_usage_service_proto_goTypes = []any{
	(*DailyUsageRequest)(nil),         // 0: com.coralogix.datausage.v1.DailyUsageRequest
	(*DailyUsageResponse)(nil),        // 1: com.coralogix.datausage.v1.DailyUsageResponse
	(*OverageDailyUsageRequest)(nil),  // 2: com.coralogix.datausage.v1.OverageDailyUsageRequest
	(*OverageDailyUsageResponse)(nil), // 3: com.coralogix.datausage.v1.OverageDailyUsageResponse
	(Range)(0),                        // 4: com.coralogix.datausage.v1.Range
	(*DateRange)(nil),                 // 5: com.coralogix.datausage.v1.DateRange
	(*DetailedDailyUsage)(nil),        // 6: com.coralogix.datausage.v1.DetailedDailyUsage
	(*OverageDailyReport)(nil),        // 7: com.coralogix.datausage.v1.OverageDailyReport
}
var file_com_coralogix_datausage_v1_daily_usage_service_proto_depIdxs = []int32{
	4, // 0: com.coralogix.datausage.v1.DailyUsageRequest.range:type_name -> com.coralogix.datausage.v1.Range
	5, // 1: com.coralogix.datausage.v1.DailyUsageRequest.date_range:type_name -> com.coralogix.datausage.v1.DateRange
	6, // 2: com.coralogix.datausage.v1.DailyUsageResponse.response:type_name -> com.coralogix.datausage.v1.DetailedDailyUsage
	4, // 3: com.coralogix.datausage.v1.OverageDailyUsageRequest.range:type_name -> com.coralogix.datausage.v1.Range
	5, // 4: com.coralogix.datausage.v1.OverageDailyUsageRequest.date_range:type_name -> com.coralogix.datausage.v1.DateRange
	7, // 5: com.coralogix.datausage.v1.OverageDailyUsageResponse.overage_detail:type_name -> com.coralogix.datausage.v1.OverageDailyReport
	0, // 6: com.coralogix.datausage.v1.DailyUsageService.DailyUsage:input_type -> com.coralogix.datausage.v1.DailyUsageRequest
	2, // 7: com.coralogix.datausage.v1.DailyUsageService.OverageDailyUsage:input_type -> com.coralogix.datausage.v1.OverageDailyUsageRequest
	1, // 8: com.coralogix.datausage.v1.DailyUsageService.DailyUsage:output_type -> com.coralogix.datausage.v1.DailyUsageResponse
	3, // 9: com.coralogix.datausage.v1.DailyUsageService.OverageDailyUsage:output_type -> com.coralogix.datausage.v1.OverageDailyUsageResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_com_coralogix_datausage_v1_daily_usage_service_proto_init() }
func file_com_coralogix_datausage_v1_daily_usage_service_proto_init() {
	if File_com_coralogix_datausage_v1_daily_usage_service_proto != nil {
		return
	}
	file_com_coralogix_datausage_v1_audit_log_proto_init()
	file_com_coralogix_datausage_v1_daily_usage_proto_init()
	file_com_coralogix_datausage_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_datausage_v1_daily_usage_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DailyUsageRequest); i {
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
		file_com_coralogix_datausage_v1_daily_usage_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DailyUsageResponse); i {
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
		file_com_coralogix_datausage_v1_daily_usage_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*OverageDailyUsageRequest); i {
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
		file_com_coralogix_datausage_v1_daily_usage_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*OverageDailyUsageResponse); i {
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
	file_com_coralogix_datausage_v1_daily_usage_service_proto_msgTypes[0].OneofWrappers = []any{
		(*DailyUsageRequest_Range)(nil),
		(*DailyUsageRequest_DateRange)(nil),
	}
	file_com_coralogix_datausage_v1_daily_usage_service_proto_msgTypes[2].OneofWrappers = []any{
		(*OverageDailyUsageRequest_Range)(nil),
		(*OverageDailyUsageRequest_DateRange)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_datausage_v1_daily_usage_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogix_datausage_v1_daily_usage_service_proto_goTypes,
		DependencyIndexes: file_com_coralogix_datausage_v1_daily_usage_service_proto_depIdxs,
		MessageInfos:      file_com_coralogix_datausage_v1_daily_usage_service_proto_msgTypes,
	}.Build()
	File_com_coralogix_datausage_v1_daily_usage_service_proto = out.File
	file_com_coralogix_datausage_v1_daily_usage_service_proto_rawDesc = nil
	file_com_coralogix_datausage_v1_daily_usage_service_proto_goTypes = nil
	file_com_coralogix_datausage_v1_daily_usage_service_proto_depIdxs = nil
}
