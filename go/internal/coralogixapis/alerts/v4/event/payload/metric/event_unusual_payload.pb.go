// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: com/coralogixapis/alerts/v4/event/payload/metric/event_unusual_payload.proto

package metric

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

type EventUnusualPayload struct {
	state               protoimpl.MessageState  `protogen:"open.v1"`
	FromTimestamp       *timestamppb.Timestamp  `protobuf:"bytes,1,opt,name=from_timestamp,json=fromTimestamp,proto3" json:"from_timestamp,omitempty"`
	ToTimestamp         *timestamppb.Timestamp  `protobuf:"bytes,2,opt,name=to_timestamp,json=toTimestamp,proto3" json:"to_timestamp,omitempty"`
	ForecastTimestamp   *wrapperspb.UInt64Value `protobuf:"bytes,3,opt,name=forecast_timestamp,json=forecastTimestamp,proto3" json:"forecast_timestamp,omitempty"`
	SearchDetails       *SearchDetails          `protobuf:"bytes,4,opt,name=search_details,json=searchDetails,proto3" json:"search_details,omitempty"`
	Distance            *wrapperspb.DoubleValue `protobuf:"bytes,5,opt,name=distance,proto3" json:"distance,omitempty"`
	ExtremeSample       *ExtremeSample          `protobuf:"bytes,6,opt,name=extreme_sample,json=extremeSample,proto3" json:"extreme_sample,omitempty"`
	RatioUnusualSamples *wrapperspb.DoubleValue `protobuf:"bytes,7,opt,name=ratio_unusual_samples,json=ratioUnusualSamples,proto3" json:"ratio_unusual_samples,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *EventUnusualPayload) Reset() {
	*x = EventUnusualPayload{}
	mi := &file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventUnusualPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventUnusualPayload) ProtoMessage() {}

func (x *EventUnusualPayload) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventUnusualPayload.ProtoReflect.Descriptor instead.
func (*EventUnusualPayload) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_rawDescGZIP(), []int{0}
}

func (x *EventUnusualPayload) GetFromTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.FromTimestamp
	}
	return nil
}

func (x *EventUnusualPayload) GetToTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.ToTimestamp
	}
	return nil
}

func (x *EventUnusualPayload) GetForecastTimestamp() *wrapperspb.UInt64Value {
	if x != nil {
		return x.ForecastTimestamp
	}
	return nil
}

func (x *EventUnusualPayload) GetSearchDetails() *SearchDetails {
	if x != nil {
		return x.SearchDetails
	}
	return nil
}

func (x *EventUnusualPayload) GetDistance() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Distance
	}
	return nil
}

func (x *EventUnusualPayload) GetExtremeSample() *ExtremeSample {
	if x != nil {
		return x.ExtremeSample
	}
	return nil
}

func (x *EventUnusualPayload) GetRatioUnusualSamples() *wrapperspb.DoubleValue {
	if x != nil {
		return x.RatioUnusualSamples
	}
	return nil
}

type SearchDetails struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FromTimestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=from_timestamp,json=fromTimestamp,proto3" json:"from_timestamp,omitempty"`
	ToTimestamp   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=to_timestamp,json=toTimestamp,proto3" json:"to_timestamp,omitempty"`
	Interval      *SearchDetailsInterval `protobuf:"bytes,3,opt,name=interval,proto3" json:"interval,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchDetails) Reset() {
	*x = SearchDetails{}
	mi := &file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchDetails) ProtoMessage() {}

func (x *SearchDetails) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchDetails.ProtoReflect.Descriptor instead.
func (*SearchDetails) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_rawDescGZIP(), []int{1}
}

func (x *SearchDetails) GetFromTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.FromTimestamp
	}
	return nil
}

func (x *SearchDetails) GetToTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.ToTimestamp
	}
	return nil
}

func (x *SearchDetails) GetInterval() *SearchDetailsInterval {
	if x != nil {
		return x.Interval
	}
	return nil
}

type SearchDetailsInterval struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Name          *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DurationMs    *wrapperspb.Int64Value  `protobuf:"bytes,2,opt,name=duration_ms,json=durationMs,proto3" json:"duration_ms,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchDetailsInterval) Reset() {
	*x = SearchDetailsInterval{}
	mi := &file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchDetailsInterval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchDetailsInterval) ProtoMessage() {}

func (x *SearchDetailsInterval) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchDetailsInterval.ProtoReflect.Descriptor instead.
func (*SearchDetailsInterval) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_rawDescGZIP(), []int{2}
}

func (x *SearchDetailsInterval) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *SearchDetailsInterval) GetDurationMs() *wrapperspb.Int64Value {
	if x != nil {
		return x.DurationMs
	}
	return nil
}

type ExtremeSample struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Timestamp     *timestamppb.Timestamp  `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	LowerLimit    *wrapperspb.DoubleValue `protobuf:"bytes,2,opt,name=lower_limit,json=lowerLimit,proto3" json:"lower_limit,omitempty"`
	UpperLimit    *wrapperspb.DoubleValue `protobuf:"bytes,3,opt,name=upper_limit,json=upperLimit,proto3" json:"upper_limit,omitempty"`
	Value         *wrapperspb.DoubleValue `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExtremeSample) Reset() {
	*x = ExtremeSample{}
	mi := &file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExtremeSample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtremeSample) ProtoMessage() {}

func (x *ExtremeSample) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtremeSample.ProtoReflect.Descriptor instead.
func (*ExtremeSample) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_rawDescGZIP(), []int{3}
}

func (x *ExtremeSample) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *ExtremeSample) GetLowerLimit() *wrapperspb.DoubleValue {
	if x != nil {
		return x.LowerLimit
	}
	return nil
}

func (x *ExtremeSample) GetUpperLimit() *wrapperspb.DoubleValue {
	if x != nil {
		return x.UpperLimit
	}
	return nil
}

func (x *ExtremeSample) GetValue() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_rawDesc = []byte{
	0x0a, 0x4c, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x34, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x6e, 0x75, 0x73, 0x75, 0x61, 0x6c,
	0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x34, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x04, 0x0a,
	0x13, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x55, 0x6e, 0x75, 0x73, 0x75, 0x61, 0x6c, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x41, 0x0a, 0x0e, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3d, 0x0a, 0x0c, 0x74, 0x6f, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x74, 0x6f, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x4b, 0x0a, 0x12, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x61,
	0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x11, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x51, 0x0a, 0x0e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x34, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0d, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x38, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x51, 0x0a, 0x0e, 0x65, 0x78, 0x74, 0x72, 0x65, 0x6d, 0x65, 0x5f, 0x73, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x73, 0x2e, 0x76, 0x34, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x65, 0x6d, 0x65, 0x53, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x52, 0x0d, 0x65, 0x78, 0x74, 0x72, 0x65, 0x6d, 0x65, 0x53, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x12, 0x50, 0x0a, 0x15, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x5f, 0x75, 0x6e, 0x75,
	0x73, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x13, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x55, 0x6e, 0x75, 0x73, 0x75, 0x61, 0x6c, 0x53, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x22, 0xe1, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x41, 0x0a, 0x0e, 0x66, 0x72, 0x6f, 0x6d, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x66, 0x72, 0x6f,
	0x6d, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3d, 0x0a, 0x0c, 0x74, 0x6f,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x74, 0x6f,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x4e, 0x0a, 0x08, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x34, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52,
	0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x22, 0x87, 0x01, 0x0a, 0x15, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x12, 0x30, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x73, 0x22, 0xfb, 0x01, 0x0a, 0x0d, 0x45, 0x78, 0x74, 0x72, 0x65, 0x6d, 0x65, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x3d, 0x0a, 0x0b, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0a, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x3d,
	0x0a, 0x0b, 0x75, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0a, 0x75, 0x70, 0x70, 0x65, 0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x32, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_rawDescData = file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_rawDesc
)

func file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_rawDescData
}

var file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_goTypes = []any{
	(*EventUnusualPayload)(nil),    // 0: com.coralogixapis.alerts.v4.EventUnusualPayload
	(*SearchDetails)(nil),          // 1: com.coralogixapis.alerts.v4.SearchDetails
	(*SearchDetailsInterval)(nil),  // 2: com.coralogixapis.alerts.v4.SearchDetailsInterval
	(*ExtremeSample)(nil),          // 3: com.coralogixapis.alerts.v4.ExtremeSample
	(*timestamppb.Timestamp)(nil),  // 4: google.protobuf.Timestamp
	(*wrapperspb.UInt64Value)(nil), // 5: google.protobuf.UInt64Value
	(*wrapperspb.DoubleValue)(nil), // 6: google.protobuf.DoubleValue
	(*wrapperspb.StringValue)(nil), // 7: google.protobuf.StringValue
	(*wrapperspb.Int64Value)(nil),  // 8: google.protobuf.Int64Value
}
var file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_depIdxs = []int32{
	4,  // 0: com.coralogixapis.alerts.v4.EventUnusualPayload.from_timestamp:type_name -> google.protobuf.Timestamp
	4,  // 1: com.coralogixapis.alerts.v4.EventUnusualPayload.to_timestamp:type_name -> google.protobuf.Timestamp
	5,  // 2: com.coralogixapis.alerts.v4.EventUnusualPayload.forecast_timestamp:type_name -> google.protobuf.UInt64Value
	1,  // 3: com.coralogixapis.alerts.v4.EventUnusualPayload.search_details:type_name -> com.coralogixapis.alerts.v4.SearchDetails
	6,  // 4: com.coralogixapis.alerts.v4.EventUnusualPayload.distance:type_name -> google.protobuf.DoubleValue
	3,  // 5: com.coralogixapis.alerts.v4.EventUnusualPayload.extreme_sample:type_name -> com.coralogixapis.alerts.v4.ExtremeSample
	6,  // 6: com.coralogixapis.alerts.v4.EventUnusualPayload.ratio_unusual_samples:type_name -> google.protobuf.DoubleValue
	4,  // 7: com.coralogixapis.alerts.v4.SearchDetails.from_timestamp:type_name -> google.protobuf.Timestamp
	4,  // 8: com.coralogixapis.alerts.v4.SearchDetails.to_timestamp:type_name -> google.protobuf.Timestamp
	2,  // 9: com.coralogixapis.alerts.v4.SearchDetails.interval:type_name -> com.coralogixapis.alerts.v4.SearchDetailsInterval
	7,  // 10: com.coralogixapis.alerts.v4.SearchDetailsInterval.name:type_name -> google.protobuf.StringValue
	8,  // 11: com.coralogixapis.alerts.v4.SearchDetailsInterval.duration_ms:type_name -> google.protobuf.Int64Value
	4,  // 12: com.coralogixapis.alerts.v4.ExtremeSample.timestamp:type_name -> google.protobuf.Timestamp
	6,  // 13: com.coralogixapis.alerts.v4.ExtremeSample.lower_limit:type_name -> google.protobuf.DoubleValue
	6,  // 14: com.coralogixapis.alerts.v4.ExtremeSample.upper_limit:type_name -> google.protobuf.DoubleValue
	6,  // 15: com.coralogixapis.alerts.v4.ExtremeSample.value:type_name -> google.protobuf.DoubleValue
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_init() }
func file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_init() {
	if File_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto = out.File
	file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_goTypes = nil
	file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_depIdxs = nil
}
