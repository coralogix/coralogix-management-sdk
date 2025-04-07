// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.26.1
// source: com/coralogixapis/alerts/v5/event/payload/metric/event_metric_unusual_payload.proto

package metric

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type EventMetricUnusualPayload struct {
	state               protoimpl.MessageState  `protogen:"open.v1"`
	FromTimestamp       *timestamppb.Timestamp  `protobuf:"bytes,1,opt,name=from_timestamp,json=fromTimestamp,proto3" json:"from_timestamp,omitempty"`
	ToTimestamp         *timestamppb.Timestamp  `protobuf:"bytes,2,opt,name=to_timestamp,json=toTimestamp,proto3" json:"to_timestamp,omitempty"`
	ForecastTimestamp   *wrapperspb.UInt64Value `protobuf:"bytes,3,opt,name=forecast_timestamp,json=forecastTimestamp,proto3" json:"forecast_timestamp,omitempty"`
	SearchDetails       *TimeRangeWithInterval  `protobuf:"bytes,4,opt,name=search_details,json=searchDetails,proto3" json:"search_details,omitempty"`
	Distance            *wrapperspb.DoubleValue `protobuf:"bytes,5,opt,name=distance,proto3" json:"distance,omitempty"`
	ExtremeSample       *ExtremeSample          `protobuf:"bytes,6,opt,name=extreme_sample,json=extremeSample,proto3" json:"extreme_sample,omitempty"`
	RatioUnusualSamples *wrapperspb.DoubleValue `protobuf:"bytes,7,opt,name=ratio_unusual_samples,json=ratioUnusualSamples,proto3" json:"ratio_unusual_samples,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *EventMetricUnusualPayload) Reset() {
	*x = EventMetricUnusualPayload{}
	mi := &file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventMetricUnusualPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventMetricUnusualPayload) ProtoMessage() {}

func (x *EventMetricUnusualPayload) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventMetricUnusualPayload.ProtoReflect.Descriptor instead.
func (*EventMetricUnusualPayload) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_rawDescGZIP(), []int{0}
}

func (x *EventMetricUnusualPayload) GetFromTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.FromTimestamp
	}
	return nil
}

func (x *EventMetricUnusualPayload) GetToTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.ToTimestamp
	}
	return nil
}

func (x *EventMetricUnusualPayload) GetForecastTimestamp() *wrapperspb.UInt64Value {
	if x != nil {
		return x.ForecastTimestamp
	}
	return nil
}

func (x *EventMetricUnusualPayload) GetSearchDetails() *TimeRangeWithInterval {
	if x != nil {
		return x.SearchDetails
	}
	return nil
}

func (x *EventMetricUnusualPayload) GetDistance() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Distance
	}
	return nil
}

func (x *EventMetricUnusualPayload) GetExtremeSample() *ExtremeSample {
	if x != nil {
		return x.ExtremeSample
	}
	return nil
}

func (x *EventMetricUnusualPayload) GetRatioUnusualSamples() *wrapperspb.DoubleValue {
	if x != nil {
		return x.RatioUnusualSamples
	}
	return nil
}

type TimeRangeWithInterval struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FromTimestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=from_timestamp,json=fromTimestamp,proto3" json:"from_timestamp,omitempty"`
	ToTimestamp   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=to_timestamp,json=toTimestamp,proto3" json:"to_timestamp,omitempty"`
	Interval      *NamedInterval         `protobuf:"bytes,3,opt,name=interval,proto3" json:"interval,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TimeRangeWithInterval) Reset() {
	*x = TimeRangeWithInterval{}
	mi := &file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimeRangeWithInterval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeRangeWithInterval) ProtoMessage() {}

func (x *TimeRangeWithInterval) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeRangeWithInterval.ProtoReflect.Descriptor instead.
func (*TimeRangeWithInterval) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_rawDescGZIP(), []int{1}
}

func (x *TimeRangeWithInterval) GetFromTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.FromTimestamp
	}
	return nil
}

func (x *TimeRangeWithInterval) GetToTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.ToTimestamp
	}
	return nil
}

func (x *TimeRangeWithInterval) GetInterval() *NamedInterval {
	if x != nil {
		return x.Interval
	}
	return nil
}

type NamedInterval struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Name          *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DurationMs    *wrapperspb.Int64Value  `protobuf:"bytes,2,opt,name=duration_ms,json=durationMs,proto3" json:"duration_ms,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NamedInterval) Reset() {
	*x = NamedInterval{}
	mi := &file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NamedInterval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamedInterval) ProtoMessage() {}

func (x *NamedInterval) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamedInterval.ProtoReflect.Descriptor instead.
func (*NamedInterval) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_rawDescGZIP(), []int{2}
}

func (x *NamedInterval) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *NamedInterval) GetDurationMs() *wrapperspb.Int64Value {
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
	mi := &file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExtremeSample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtremeSample) ProtoMessage() {}

func (x *ExtremeSample) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_msgTypes[3]
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
	return file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_rawDescGZIP(), []int{3}
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

var File_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto protoreflect.FileDescriptor

const file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_rawDesc = "" +
	"\n" +
	"Scom/coralogixapis/alerts/v5/event/payload/metric/event_metric_unusual_payload.proto\x12\x1bcom.coralogixapis.alerts.v5\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/wrappers.proto\"\xa4\x04\n" +
	"\x19EventMetricUnusualPayload\x12A\n" +
	"\x0efrom_timestamp\x18\x01 \x01(\v2\x1a.google.protobuf.TimestampR\rfromTimestamp\x12=\n" +
	"\fto_timestamp\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\vtoTimestamp\x12K\n" +
	"\x12forecast_timestamp\x18\x03 \x01(\v2\x1c.google.protobuf.UInt64ValueR\x11forecastTimestamp\x12Y\n" +
	"\x0esearch_details\x18\x04 \x01(\v22.com.coralogixapis.alerts.v5.TimeRangeWithIntervalR\rsearchDetails\x128\n" +
	"\bdistance\x18\x05 \x01(\v2\x1c.google.protobuf.DoubleValueR\bdistance\x12Q\n" +
	"\x0eextreme_sample\x18\x06 \x01(\v2*.com.coralogixapis.alerts.v5.ExtremeSampleR\rextremeSample\x12P\n" +
	"\x15ratio_unusual_samples\x18\a \x01(\v2\x1c.google.protobuf.DoubleValueR\x13ratioUnusualSamples\"\xe1\x01\n" +
	"\x15TimeRangeWithInterval\x12A\n" +
	"\x0efrom_timestamp\x18\x01 \x01(\v2\x1a.google.protobuf.TimestampR\rfromTimestamp\x12=\n" +
	"\fto_timestamp\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\vtoTimestamp\x12F\n" +
	"\binterval\x18\x03 \x01(\v2*.com.coralogixapis.alerts.v5.NamedIntervalR\binterval\"\x7f\n" +
	"\rNamedInterval\x120\n" +
	"\x04name\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueR\x04name\x12<\n" +
	"\vduration_ms\x18\x02 \x01(\v2\x1b.google.protobuf.Int64ValueR\n" +
	"durationMs\"\xfb\x01\n" +
	"\rExtremeSample\x128\n" +
	"\ttimestamp\x18\x01 \x01(\v2\x1a.google.protobuf.TimestampR\ttimestamp\x12=\n" +
	"\vlower_limit\x18\x02 \x01(\v2\x1c.google.protobuf.DoubleValueR\n" +
	"lowerLimit\x12=\n" +
	"\vupper_limit\x18\x03 \x01(\v2\x1c.google.protobuf.DoubleValueR\n" +
	"upperLimit\x122\n" +
	"\x05value\x18\x04 \x01(\v2\x1c.google.protobuf.DoubleValueR\x05valueb\x06proto3"

var (
	file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_rawDesc), len(file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_rawDescData
}

var file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_goTypes = []any{
	(*EventMetricUnusualPayload)(nil), // 0: com.coralogixapis.alerts.v5.EventMetricUnusualPayload
	(*TimeRangeWithInterval)(nil),     // 1: com.coralogixapis.alerts.v5.TimeRangeWithInterval
	(*NamedInterval)(nil),             // 2: com.coralogixapis.alerts.v5.NamedInterval
	(*ExtremeSample)(nil),             // 3: com.coralogixapis.alerts.v5.ExtremeSample
	(*timestamppb.Timestamp)(nil),     // 4: google.protobuf.Timestamp
	(*wrapperspb.UInt64Value)(nil),    // 5: google.protobuf.UInt64Value
	(*wrapperspb.DoubleValue)(nil),    // 6: google.protobuf.DoubleValue
	(*wrapperspb.StringValue)(nil),    // 7: google.protobuf.StringValue
	(*wrapperspb.Int64Value)(nil),     // 8: google.protobuf.Int64Value
}
var file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_depIdxs = []int32{
	4,  // 0: com.coralogixapis.alerts.v5.EventMetricUnusualPayload.from_timestamp:type_name -> google.protobuf.Timestamp
	4,  // 1: com.coralogixapis.alerts.v5.EventMetricUnusualPayload.to_timestamp:type_name -> google.protobuf.Timestamp
	5,  // 2: com.coralogixapis.alerts.v5.EventMetricUnusualPayload.forecast_timestamp:type_name -> google.protobuf.UInt64Value
	1,  // 3: com.coralogixapis.alerts.v5.EventMetricUnusualPayload.search_details:type_name -> com.coralogixapis.alerts.v5.TimeRangeWithInterval
	6,  // 4: com.coralogixapis.alerts.v5.EventMetricUnusualPayload.distance:type_name -> google.protobuf.DoubleValue
	3,  // 5: com.coralogixapis.alerts.v5.EventMetricUnusualPayload.extreme_sample:type_name -> com.coralogixapis.alerts.v5.ExtremeSample
	6,  // 6: com.coralogixapis.alerts.v5.EventMetricUnusualPayload.ratio_unusual_samples:type_name -> google.protobuf.DoubleValue
	4,  // 7: com.coralogixapis.alerts.v5.TimeRangeWithInterval.from_timestamp:type_name -> google.protobuf.Timestamp
	4,  // 8: com.coralogixapis.alerts.v5.TimeRangeWithInterval.to_timestamp:type_name -> google.protobuf.Timestamp
	2,  // 9: com.coralogixapis.alerts.v5.TimeRangeWithInterval.interval:type_name -> com.coralogixapis.alerts.v5.NamedInterval
	7,  // 10: com.coralogixapis.alerts.v5.NamedInterval.name:type_name -> google.protobuf.StringValue
	8,  // 11: com.coralogixapis.alerts.v5.NamedInterval.duration_ms:type_name -> google.protobuf.Int64Value
	4,  // 12: com.coralogixapis.alerts.v5.ExtremeSample.timestamp:type_name -> google.protobuf.Timestamp
	6,  // 13: com.coralogixapis.alerts.v5.ExtremeSample.lower_limit:type_name -> google.protobuf.DoubleValue
	6,  // 14: com.coralogixapis.alerts.v5.ExtremeSample.upper_limit:type_name -> google.protobuf.DoubleValue
	6,  // 15: com.coralogixapis.alerts.v5.ExtremeSample.value:type_name -> google.protobuf.DoubleValue
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_init()
}
func file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_init() {
	if File_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_rawDesc), len(file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto = out.File
	file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_goTypes = nil
	file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_payload_proto_depIdxs = nil
}
