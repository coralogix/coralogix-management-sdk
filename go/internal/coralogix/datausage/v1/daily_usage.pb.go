// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: src/com/coralogix/datausage/v1/daily_usage.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// / Daily usage message
type DetailedDailyUsage struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Day of the daily usage report
	StatsDate *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=stats_date,json=statsDate,proto3" json:"stats_date,omitempty"`
	// blocked data in GB
	Block *wrapperspb.FloatValue `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
	// blocked quota in units
	BlockCountQuota *wrapperspb.FloatValue `protobuf:"bytes,3,opt,name=block_count_quota,json=blockCountQuota,proto3" json:"block_count_quota,omitempty"`
	// low TCO traffic data in GB
	Low *wrapperspb.FloatValue `protobuf:"bytes,4,opt,name=low,proto3" json:"low,omitempty"`
	// low TCO traffic data in units
	LowCountQuota *wrapperspb.FloatValue `protobuf:"bytes,5,opt,name=low_count_quota,json=lowCountQuota,proto3" json:"low_count_quota,omitempty"`
	// medium TCO traffic data in GB
	Medium *wrapperspb.FloatValue `protobuf:"bytes,6,opt,name=medium,proto3" json:"medium,omitempty"`
	// medium TCO traffic data in units
	MediumCountQuota *wrapperspb.FloatValue `protobuf:"bytes,7,opt,name=medium_count_quota,json=mediumCountQuota,proto3" json:"medium_count_quota,omitempty"`
	// normal traffic data in GB
	High *wrapperspb.FloatValue `protobuf:"bytes,8,opt,name=high,proto3" json:"high,omitempty"`
	// normal traffic data in units
	HighCountQuota *wrapperspb.FloatValue `protobuf:"bytes,9,opt,name=high_count_quota,json=highCountQuota,proto3" json:"high_count_quota,omitempty"`
	// total consumed bytes in GB
	Total *wrapperspb.FloatValue `protobuf:"bytes,10,opt,name=total,proto3" json:"total,omitempty"`
	// total consumed quota in units
	TotalCountQuota *wrapperspb.FloatValue `protobuf:"bytes,11,opt,name=total_count_quota,json=totalCountQuota,proto3" json:"total_count_quota,omitempty"`
	// total consumed quota in units for metrics
	HighMetricsCountQuota *wrapperspb.FloatValue `protobuf:"bytes,12,opt,name=high_metrics_count_quota,json=highMetricsCountQuota,proto3" json:"high_metrics_count_quota,omitempty"`
	// total consumed quota in GB for metrics
	HighMetricsQuota *wrapperspb.FloatValue `protobuf:"bytes,13,opt,name=high_metrics_quota,json=highMetricsQuota,proto3" json:"high_metrics_quota,omitempty"`
	// total consumed quota in units for high tracing
	HighTracingCountQuota *wrapperspb.FloatValue `protobuf:"bytes,14,opt,name=high_tracing_count_quota,json=highTracingCountQuota,proto3" json:"high_tracing_count_quota,omitempty"`
	// total consumed quota in GB for high tracing
	HighTracingQuota *wrapperspb.FloatValue `protobuf:"bytes,15,opt,name=high_tracing_quota,json=highTracingQuota,proto3" json:"high_tracing_quota,omitempty"`
	// total consumed quota in units for medium tracing
	MediumTracingCountQuota *wrapperspb.FloatValue `protobuf:"bytes,16,opt,name=medium_tracing_count_quota,json=mediumTracingCountQuota,proto3" json:"medium_tracing_count_quota,omitempty"`
	// total consumed quota in GB for medium tracing
	MediumTracingQuota *wrapperspb.FloatValue `protobuf:"bytes,17,opt,name=medium_tracing_quota,json=mediumTracingQuota,proto3" json:"medium_tracing_quota,omitempty"`
	// total consumed quota in units for low tracing
	LowTracingCountQuota *wrapperspb.FloatValue `protobuf:"bytes,18,opt,name=low_tracing_count_quota,json=lowTracingCountQuota,proto3" json:"low_tracing_count_quota,omitempty"`
	// total consumed quota in GB for low tracing
	LowTracingQuota *wrapperspb.FloatValue `protobuf:"bytes,19,opt,name=low_tracing_quota,json=lowTracingQuota,proto3" json:"low_tracing_quota,omitempty"`
	// total consumed quota in units for low session recording
	LowSessionRecordingCountQuota *wrapperspb.FloatValue `protobuf:"bytes,20,opt,name=low_session_recording_count_quota,json=lowSessionRecordingCountQuota,proto3" json:"low_session_recording_count_quota,omitempty"`
	// total consumed quota in GB for low session recording
	LowSessionRecordingQuota *wrapperspb.FloatValue `protobuf:"bytes,21,opt,name=low_session_recording_quota,json=lowSessionRecordingQuota,proto3" json:"low_session_recording_quota,omitempty"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *DetailedDailyUsage) Reset() {
	*x = DetailedDailyUsage{}
	mi := &file_src_com_coralogix_datausage_v1_daily_usage_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DetailedDailyUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailedDailyUsage) ProtoMessage() {}

func (x *DetailedDailyUsage) ProtoReflect() protoreflect.Message {
	mi := &file_src_com_coralogix_datausage_v1_daily_usage_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailedDailyUsage.ProtoReflect.Descriptor instead.
func (*DetailedDailyUsage) Descriptor() ([]byte, []int) {
	return file_src_com_coralogix_datausage_v1_daily_usage_proto_rawDescGZIP(), []int{0}
}

func (x *DetailedDailyUsage) GetStatsDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StatsDate
	}
	return nil
}

func (x *DetailedDailyUsage) GetBlock() *wrapperspb.FloatValue {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *DetailedDailyUsage) GetBlockCountQuota() *wrapperspb.FloatValue {
	if x != nil {
		return x.BlockCountQuota
	}
	return nil
}

func (x *DetailedDailyUsage) GetLow() *wrapperspb.FloatValue {
	if x != nil {
		return x.Low
	}
	return nil
}

func (x *DetailedDailyUsage) GetLowCountQuota() *wrapperspb.FloatValue {
	if x != nil {
		return x.LowCountQuota
	}
	return nil
}

func (x *DetailedDailyUsage) GetMedium() *wrapperspb.FloatValue {
	if x != nil {
		return x.Medium
	}
	return nil
}

func (x *DetailedDailyUsage) GetMediumCountQuota() *wrapperspb.FloatValue {
	if x != nil {
		return x.MediumCountQuota
	}
	return nil
}

func (x *DetailedDailyUsage) GetHigh() *wrapperspb.FloatValue {
	if x != nil {
		return x.High
	}
	return nil
}

func (x *DetailedDailyUsage) GetHighCountQuota() *wrapperspb.FloatValue {
	if x != nil {
		return x.HighCountQuota
	}
	return nil
}

func (x *DetailedDailyUsage) GetTotal() *wrapperspb.FloatValue {
	if x != nil {
		return x.Total
	}
	return nil
}

func (x *DetailedDailyUsage) GetTotalCountQuota() *wrapperspb.FloatValue {
	if x != nil {
		return x.TotalCountQuota
	}
	return nil
}

func (x *DetailedDailyUsage) GetHighMetricsCountQuota() *wrapperspb.FloatValue {
	if x != nil {
		return x.HighMetricsCountQuota
	}
	return nil
}

func (x *DetailedDailyUsage) GetHighMetricsQuota() *wrapperspb.FloatValue {
	if x != nil {
		return x.HighMetricsQuota
	}
	return nil
}

func (x *DetailedDailyUsage) GetHighTracingCountQuota() *wrapperspb.FloatValue {
	if x != nil {
		return x.HighTracingCountQuota
	}
	return nil
}

func (x *DetailedDailyUsage) GetHighTracingQuota() *wrapperspb.FloatValue {
	if x != nil {
		return x.HighTracingQuota
	}
	return nil
}

func (x *DetailedDailyUsage) GetMediumTracingCountQuota() *wrapperspb.FloatValue {
	if x != nil {
		return x.MediumTracingCountQuota
	}
	return nil
}

func (x *DetailedDailyUsage) GetMediumTracingQuota() *wrapperspb.FloatValue {
	if x != nil {
		return x.MediumTracingQuota
	}
	return nil
}

func (x *DetailedDailyUsage) GetLowTracingCountQuota() *wrapperspb.FloatValue {
	if x != nil {
		return x.LowTracingCountQuota
	}
	return nil
}

func (x *DetailedDailyUsage) GetLowTracingQuota() *wrapperspb.FloatValue {
	if x != nil {
		return x.LowTracingQuota
	}
	return nil
}

func (x *DetailedDailyUsage) GetLowSessionRecordingCountQuota() *wrapperspb.FloatValue {
	if x != nil {
		return x.LowSessionRecordingCountQuota
	}
	return nil
}

func (x *DetailedDailyUsage) GetLowSessionRecordingQuota() *wrapperspb.FloatValue {
	if x != nil {
		return x.LowSessionRecordingQuota
	}
	return nil
}

// / Data consumed that exceeds the planned quota (only for pay as you go)
type OverageDailyReport struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Date          *timestamppb.Timestamp  `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	OverageUnit   *wrapperspb.UInt64Value `protobuf:"bytes,2,opt,name=overage_unit,json=overageUnit,proto3" json:"overage_unit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OverageDailyReport) Reset() {
	*x = OverageDailyReport{}
	mi := &file_src_com_coralogix_datausage_v1_daily_usage_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OverageDailyReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverageDailyReport) ProtoMessage() {}

func (x *OverageDailyReport) ProtoReflect() protoreflect.Message {
	mi := &file_src_com_coralogix_datausage_v1_daily_usage_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverageDailyReport.ProtoReflect.Descriptor instead.
func (*OverageDailyReport) Descriptor() ([]byte, []int) {
	return file_src_com_coralogix_datausage_v1_daily_usage_proto_rawDescGZIP(), []int{1}
}

func (x *OverageDailyReport) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *OverageDailyReport) GetOverageUnit() *wrapperspb.UInt64Value {
	if x != nil {
		return x.OverageUnit
	}
	return nil
}

var File_src_com_coralogix_datausage_v1_daily_usage_proto protoreflect.FileDescriptor

var file_src_com_coralogix_datausage_v1_daily_usage_proto_rawDesc = []byte{
	0x0a, 0x30, 0x73, 0x72, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xfe, 0x0b, 0x0a, 0x12, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x44, 0x61, 0x69, 0x6c,
	0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x73, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x31, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x47, 0x0a, 0x11, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x2d, 0x0a,
	0x03, 0x6c, 0x6f, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f,
	0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x43, 0x0a, 0x0f,
	0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0d, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x51, 0x75, 0x6f, 0x74,
	0x61, 0x12, 0x33, 0x0a, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06,
	0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x12, 0x49, 0x0a, 0x12, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x10, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x51, 0x75, 0x6f, 0x74,
	0x61, 0x12, 0x2f, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x68, 0x69,
	0x67, 0x68, 0x12, 0x45, 0x0a, 0x10, 0x68, 0x69, 0x67, 0x68, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x68, 0x69, 0x67, 0x68, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x31, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x47, 0x0a, 0x11,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x71, 0x75, 0x6f, 0x74,
	0x61, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x54, 0x0a, 0x18, 0x68, 0x69, 0x67, 0x68, 0x5f, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x71, 0x75, 0x6f, 0x74,
	0x61, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x15, 0x68, 0x69, 0x67, 0x68, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x49, 0x0a, 0x12, 0x68,
	0x69, 0x67, 0x68, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x71, 0x75, 0x6f, 0x74,
	0x61, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x68, 0x69, 0x67, 0x68, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x54, 0x0a, 0x18, 0x68, 0x69, 0x67, 0x68, 0x5f, 0x74,
	0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x71, 0x75, 0x6f,
	0x74, 0x61, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x15, 0x68, 0x69, 0x67, 0x68, 0x54, 0x72, 0x61, 0x63, 0x69,
	0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x49, 0x0a, 0x12,
	0x68, 0x69, 0x67, 0x68, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x71, 0x75, 0x6f,
	0x74, 0x61, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x68, 0x69, 0x67, 0x68, 0x54, 0x72, 0x61, 0x63, 0x69,
	0x6e, 0x67, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x58, 0x0a, 0x1a, 0x6d, 0x65, 0x64, 0x69, 0x75,
	0x6d, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c,
	0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x17, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d,
	0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x51, 0x75, 0x6f, 0x74,
	0x61, 0x12, 0x4d, 0x0a, 0x14, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x5f, 0x74, 0x72, 0x61, 0x63,
	0x69, 0x6e, 0x67, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x12, 0x6d, 0x65,
	0x64, 0x69, 0x75, 0x6d, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x51, 0x75, 0x6f, 0x74, 0x61,
	0x12, 0x52, 0x0a, 0x17, 0x6c, 0x6f, 0x77, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x14,
	0x6c, 0x6f, 0x77, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x51,
	0x75, 0x6f, 0x74, 0x61, 0x12, 0x47, 0x0a, 0x11, 0x6c, 0x6f, 0x77, 0x5f, 0x74, 0x72, 0x61, 0x63,
	0x69, 0x6e, 0x67, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x6c, 0x6f,
	0x77, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x65, 0x0a,
	0x21, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x71, 0x75, 0x6f,
	0x74, 0x61, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x1d, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x51,
	0x75, 0x6f, 0x74, 0x61, 0x12, 0x5a, 0x0a, 0x1b, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x71, 0x75,
	0x6f, 0x74, 0x61, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x18, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x51, 0x75, 0x6f, 0x74, 0x61,
	0x22, 0x85, 0x01, 0x0a, 0x12, 0x4f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x44, 0x61, 0x69, 0x6c,
	0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x6f, 0x76, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x6f, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_com_coralogix_datausage_v1_daily_usage_proto_rawDescOnce sync.Once
	file_src_com_coralogix_datausage_v1_daily_usage_proto_rawDescData = file_src_com_coralogix_datausage_v1_daily_usage_proto_rawDesc
)

func file_src_com_coralogix_datausage_v1_daily_usage_proto_rawDescGZIP() []byte {
	file_src_com_coralogix_datausage_v1_daily_usage_proto_rawDescOnce.Do(func() {
		file_src_com_coralogix_datausage_v1_daily_usage_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_com_coralogix_datausage_v1_daily_usage_proto_rawDescData)
	})
	return file_src_com_coralogix_datausage_v1_daily_usage_proto_rawDescData
}

var file_src_com_coralogix_datausage_v1_daily_usage_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_src_com_coralogix_datausage_v1_daily_usage_proto_goTypes = []any{
	(*DetailedDailyUsage)(nil),     // 0: com.coralogix.datausage.v1.DetailedDailyUsage
	(*OverageDailyReport)(nil),     // 1: com.coralogix.datausage.v1.OverageDailyReport
	(*timestamppb.Timestamp)(nil),  // 2: google.protobuf.Timestamp
	(*wrapperspb.FloatValue)(nil),  // 3: google.protobuf.FloatValue
	(*wrapperspb.UInt64Value)(nil), // 4: google.protobuf.UInt64Value
}
var file_src_com_coralogix_datausage_v1_daily_usage_proto_depIdxs = []int32{
	2,  // 0: com.coralogix.datausage.v1.DetailedDailyUsage.stats_date:type_name -> google.protobuf.Timestamp
	3,  // 1: com.coralogix.datausage.v1.DetailedDailyUsage.block:type_name -> google.protobuf.FloatValue
	3,  // 2: com.coralogix.datausage.v1.DetailedDailyUsage.block_count_quota:type_name -> google.protobuf.FloatValue
	3,  // 3: com.coralogix.datausage.v1.DetailedDailyUsage.low:type_name -> google.protobuf.FloatValue
	3,  // 4: com.coralogix.datausage.v1.DetailedDailyUsage.low_count_quota:type_name -> google.protobuf.FloatValue
	3,  // 5: com.coralogix.datausage.v1.DetailedDailyUsage.medium:type_name -> google.protobuf.FloatValue
	3,  // 6: com.coralogix.datausage.v1.DetailedDailyUsage.medium_count_quota:type_name -> google.protobuf.FloatValue
	3,  // 7: com.coralogix.datausage.v1.DetailedDailyUsage.high:type_name -> google.protobuf.FloatValue
	3,  // 8: com.coralogix.datausage.v1.DetailedDailyUsage.high_count_quota:type_name -> google.protobuf.FloatValue
	3,  // 9: com.coralogix.datausage.v1.DetailedDailyUsage.total:type_name -> google.protobuf.FloatValue
	3,  // 10: com.coralogix.datausage.v1.DetailedDailyUsage.total_count_quota:type_name -> google.protobuf.FloatValue
	3,  // 11: com.coralogix.datausage.v1.DetailedDailyUsage.high_metrics_count_quota:type_name -> google.protobuf.FloatValue
	3,  // 12: com.coralogix.datausage.v1.DetailedDailyUsage.high_metrics_quota:type_name -> google.protobuf.FloatValue
	3,  // 13: com.coralogix.datausage.v1.DetailedDailyUsage.high_tracing_count_quota:type_name -> google.protobuf.FloatValue
	3,  // 14: com.coralogix.datausage.v1.DetailedDailyUsage.high_tracing_quota:type_name -> google.protobuf.FloatValue
	3,  // 15: com.coralogix.datausage.v1.DetailedDailyUsage.medium_tracing_count_quota:type_name -> google.protobuf.FloatValue
	3,  // 16: com.coralogix.datausage.v1.DetailedDailyUsage.medium_tracing_quota:type_name -> google.protobuf.FloatValue
	3,  // 17: com.coralogix.datausage.v1.DetailedDailyUsage.low_tracing_count_quota:type_name -> google.protobuf.FloatValue
	3,  // 18: com.coralogix.datausage.v1.DetailedDailyUsage.low_tracing_quota:type_name -> google.protobuf.FloatValue
	3,  // 19: com.coralogix.datausage.v1.DetailedDailyUsage.low_session_recording_count_quota:type_name -> google.protobuf.FloatValue
	3,  // 20: com.coralogix.datausage.v1.DetailedDailyUsage.low_session_recording_quota:type_name -> google.protobuf.FloatValue
	2,  // 21: com.coralogix.datausage.v1.OverageDailyReport.date:type_name -> google.protobuf.Timestamp
	4,  // 22: com.coralogix.datausage.v1.OverageDailyReport.overage_unit:type_name -> google.protobuf.UInt64Value
	23, // [23:23] is the sub-list for method output_type
	23, // [23:23] is the sub-list for method input_type
	23, // [23:23] is the sub-list for extension type_name
	23, // [23:23] is the sub-list for extension extendee
	0,  // [0:23] is the sub-list for field type_name
}

func init() { file_src_com_coralogix_datausage_v1_daily_usage_proto_init() }
func file_src_com_coralogix_datausage_v1_daily_usage_proto_init() {
	if File_src_com_coralogix_datausage_v1_daily_usage_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_src_com_coralogix_datausage_v1_daily_usage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_com_coralogix_datausage_v1_daily_usage_proto_goTypes,
		DependencyIndexes: file_src_com_coralogix_datausage_v1_daily_usage_proto_depIdxs,
		MessageInfos:      file_src_com_coralogix_datausage_v1_daily_usage_proto_msgTypes,
	}.Build()
	File_src_com_coralogix_datausage_v1_daily_usage_proto = out.File
	file_src_com_coralogix_datausage_v1_daily_usage_proto_rawDesc = nil
	file_src_com_coralogix_datausage_v1_daily_usage_proto_goTypes = nil
	file_src_com_coralogix_datausage_v1_daily_usage_proto_depIdxs = nil
}
