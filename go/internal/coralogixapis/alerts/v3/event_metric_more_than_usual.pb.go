// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.26.1
// source: com/coralogixapis/alerts/v3/event/payload/metric/event_metric_more_than_usual.proto

package v3

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

type EventMetricMoreThanUsual struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	DynamicAlertMatch *DynamicAlertMatch     `protobuf:"bytes,1,opt,name=dynamic_alert_match,json=dynamicAlertMatch,proto3" json:"dynamic_alert_match,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *EventMetricMoreThanUsual) Reset() {
	*x = EventMetricMoreThanUsual{}
	mi := &file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventMetricMoreThanUsual) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventMetricMoreThanUsual) ProtoMessage() {}

func (x *EventMetricMoreThanUsual) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventMetricMoreThanUsual.ProtoReflect.Descriptor instead.
func (*EventMetricMoreThanUsual) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_rawDescGZIP(), []int{0}
}

func (x *EventMetricMoreThanUsual) GetDynamicAlertMatch() *DynamicAlertMatch {
	if x != nil {
		return x.DynamicAlertMatch
	}
	return nil
}

type DynamicAlertMatch struct {
	state             protoimpl.MessageState  `protogen:"open.v1"`
	ObservedDetails   *ObservedDetails        `protobuf:"bytes,1,opt,name=observed_details,json=observedDetails,proto3" json:"observed_details,omitempty"`
	SearchDetails     *SearchDetails          `protobuf:"bytes,2,opt,name=search_details,json=searchDetails,proto3" json:"search_details,omitempty"`
	TopBucket         *TopBucket              `protobuf:"bytes,3,opt,name=top_bucket,json=topBucket,proto3" json:"top_bucket,omitempty"`
	ForecastTimestamp *wrapperspb.UInt64Value `protobuf:"bytes,4,opt,name=forecast_timestamp,json=forecastTimestamp,proto3" json:"forecast_timestamp,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *DynamicAlertMatch) Reset() {
	*x = DynamicAlertMatch{}
	mi := &file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DynamicAlertMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DynamicAlertMatch) ProtoMessage() {}

func (x *DynamicAlertMatch) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DynamicAlertMatch.ProtoReflect.Descriptor instead.
func (*DynamicAlertMatch) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_rawDescGZIP(), []int{1}
}

func (x *DynamicAlertMatch) GetObservedDetails() *ObservedDetails {
	if x != nil {
		return x.ObservedDetails
	}
	return nil
}

func (x *DynamicAlertMatch) GetSearchDetails() *SearchDetails {
	if x != nil {
		return x.SearchDetails
	}
	return nil
}

func (x *DynamicAlertMatch) GetTopBucket() *TopBucket {
	if x != nil {
		return x.TopBucket
	}
	return nil
}

func (x *DynamicAlertMatch) GetForecastTimestamp() *wrapperspb.UInt64Value {
	if x != nil {
		return x.ForecastTimestamp
	}
	return nil
}

type ObservedDetails struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	FromTimestamp *timestamppb.Timestamp  `protobuf:"bytes,1,opt,name=from_timestamp,json=fromTimestamp,proto3" json:"from_timestamp,omitempty"`
	ToTimestamp   *timestamppb.Timestamp  `protobuf:"bytes,2,opt,name=to_timestamp,json=toTimestamp,proto3" json:"to_timestamp,omitempty"`
	Ratio         *wrapperspb.DoubleValue `protobuf:"bytes,3,opt,name=ratio,proto3" json:"ratio,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ObservedDetails) Reset() {
	*x = ObservedDetails{}
	mi := &file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ObservedDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObservedDetails) ProtoMessage() {}

func (x *ObservedDetails) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObservedDetails.ProtoReflect.Descriptor instead.
func (*ObservedDetails) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_rawDescGZIP(), []int{2}
}

func (x *ObservedDetails) GetFromTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.FromTimestamp
	}
	return nil
}

func (x *ObservedDetails) GetToTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.ToTimestamp
	}
	return nil
}

func (x *ObservedDetails) GetRatio() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Ratio
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
	mi := &file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchDetails) ProtoMessage() {}

func (x *SearchDetails) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_msgTypes[3]
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
	return file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_rawDescGZIP(), []int{3}
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
	mi := &file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchDetailsInterval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchDetailsInterval) ProtoMessage() {}

func (x *SearchDetailsInterval) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_msgTypes[4]
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
	return file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_rawDescGZIP(), []int{4}
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

type TopBucket struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Limits        []*wrapperspb.DoubleValue `protobuf:"bytes,1,rep,name=limits,proto3" json:"limits,omitempty"`
	Timestamp     *timestamppb.Timestamp    `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Value         *wrapperspb.DoubleValue   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TopBucket) Reset() {
	*x = TopBucket{}
	mi := &file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TopBucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopBucket) ProtoMessage() {}

func (x *TopBucket) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopBucket.ProtoReflect.Descriptor instead.
func (*TopBucket) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_rawDescGZIP(), []int{5}
}

func (x *TopBucket) GetLimits() []*wrapperspb.DoubleValue {
	if x != nil {
		return x.Limits
	}
	return nil
}

func (x *TopBucket) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *TopBucket) GetValue() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto protoreflect.FileDescriptor

const file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_rawDesc = "" +
	"\n" +
	"Scom/coralogixapis/alerts/v3/event/payload/metric/event_metric_more_than_usual.proto\x12\x1bcom.coralogixapis.alerts.v3\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/wrappers.proto\"z\n" +
	"\x18EventMetricMoreThanUsual\x12^\n" +
	"\x13dynamic_alert_match\x18\x01 \x01(\v2..com.coralogixapis.alerts.v3.DynamicAlertMatchR\x11dynamicAlertMatch\"\xd3\x02\n" +
	"\x11DynamicAlertMatch\x12W\n" +
	"\x10observed_details\x18\x01 \x01(\v2,.com.coralogixapis.alerts.v3.ObservedDetailsR\x0fobservedDetails\x12Q\n" +
	"\x0esearch_details\x18\x02 \x01(\v2*.com.coralogixapis.alerts.v3.SearchDetailsR\rsearchDetails\x12E\n" +
	"\n" +
	"top_bucket\x18\x03 \x01(\v2&.com.coralogixapis.alerts.v3.TopBucketR\ttopBucket\x12K\n" +
	"\x12forecast_timestamp\x18\x04 \x01(\v2\x1c.google.protobuf.UInt64ValueR\x11forecastTimestamp\"\xc7\x01\n" +
	"\x0fObservedDetails\x12A\n" +
	"\x0efrom_timestamp\x18\x01 \x01(\v2\x1a.google.protobuf.TimestampR\rfromTimestamp\x12=\n" +
	"\fto_timestamp\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\vtoTimestamp\x122\n" +
	"\x05ratio\x18\x03 \x01(\v2\x1c.google.protobuf.DoubleValueR\x05ratio\"\xe1\x01\n" +
	"\rSearchDetails\x12A\n" +
	"\x0efrom_timestamp\x18\x01 \x01(\v2\x1a.google.protobuf.TimestampR\rfromTimestamp\x12=\n" +
	"\fto_timestamp\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\vtoTimestamp\x12N\n" +
	"\binterval\x18\x03 \x01(\v22.com.coralogixapis.alerts.v3.SearchDetailsIntervalR\binterval\"\x87\x01\n" +
	"\x15SearchDetailsInterval\x120\n" +
	"\x04name\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueR\x04name\x12<\n" +
	"\vduration_ms\x18\x02 \x01(\v2\x1b.google.protobuf.Int64ValueR\n" +
	"durationMs\"\xaf\x01\n" +
	"\tTopBucket\x124\n" +
	"\x06limits\x18\x01 \x03(\v2\x1c.google.protobuf.DoubleValueR\x06limits\x128\n" +
	"\ttimestamp\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\ttimestamp\x122\n" +
	"\x05value\x18\x03 \x01(\v2\x1c.google.protobuf.DoubleValueR\x05valueb\x06proto3"

var (
	file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_goTypes = []any{
	(*EventMetricMoreThanUsual)(nil), // 0: com.coralogixapis.alerts.v3.EventMetricMoreThanUsual
	(*DynamicAlertMatch)(nil),        // 1: com.coralogixapis.alerts.v3.DynamicAlertMatch
	(*ObservedDetails)(nil),          // 2: com.coralogixapis.alerts.v3.ObservedDetails
	(*SearchDetails)(nil),            // 3: com.coralogixapis.alerts.v3.SearchDetails
	(*SearchDetailsInterval)(nil),    // 4: com.coralogixapis.alerts.v3.SearchDetailsInterval
	(*TopBucket)(nil),                // 5: com.coralogixapis.alerts.v3.TopBucket
	(*wrapperspb.UInt64Value)(nil),   // 6: google.protobuf.UInt64Value
	(*timestamppb.Timestamp)(nil),    // 7: google.protobuf.Timestamp
	(*wrapperspb.DoubleValue)(nil),   // 8: google.protobuf.DoubleValue
	(*wrapperspb.StringValue)(nil),   // 9: google.protobuf.StringValue
	(*wrapperspb.Int64Value)(nil),    // 10: google.protobuf.Int64Value
}
var file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_depIdxs = []int32{
	1,  // 0: com.coralogixapis.alerts.v3.EventMetricMoreThanUsual.dynamic_alert_match:type_name -> com.coralogixapis.alerts.v3.DynamicAlertMatch
	2,  // 1: com.coralogixapis.alerts.v3.DynamicAlertMatch.observed_details:type_name -> com.coralogixapis.alerts.v3.ObservedDetails
	3,  // 2: com.coralogixapis.alerts.v3.DynamicAlertMatch.search_details:type_name -> com.coralogixapis.alerts.v3.SearchDetails
	5,  // 3: com.coralogixapis.alerts.v3.DynamicAlertMatch.top_bucket:type_name -> com.coralogixapis.alerts.v3.TopBucket
	6,  // 4: com.coralogixapis.alerts.v3.DynamicAlertMatch.forecast_timestamp:type_name -> google.protobuf.UInt64Value
	7,  // 5: com.coralogixapis.alerts.v3.ObservedDetails.from_timestamp:type_name -> google.protobuf.Timestamp
	7,  // 6: com.coralogixapis.alerts.v3.ObservedDetails.to_timestamp:type_name -> google.protobuf.Timestamp
	8,  // 7: com.coralogixapis.alerts.v3.ObservedDetails.ratio:type_name -> google.protobuf.DoubleValue
	7,  // 8: com.coralogixapis.alerts.v3.SearchDetails.from_timestamp:type_name -> google.protobuf.Timestamp
	7,  // 9: com.coralogixapis.alerts.v3.SearchDetails.to_timestamp:type_name -> google.protobuf.Timestamp
	4,  // 10: com.coralogixapis.alerts.v3.SearchDetails.interval:type_name -> com.coralogixapis.alerts.v3.SearchDetailsInterval
	9,  // 11: com.coralogixapis.alerts.v3.SearchDetailsInterval.name:type_name -> google.protobuf.StringValue
	10, // 12: com.coralogixapis.alerts.v3.SearchDetailsInterval.duration_ms:type_name -> google.protobuf.Int64Value
	8,  // 13: com.coralogixapis.alerts.v3.TopBucket.limits:type_name -> google.protobuf.DoubleValue
	7,  // 14: com.coralogixapis.alerts.v3.TopBucket.timestamp:type_name -> google.protobuf.Timestamp
	8,  // 15: com.coralogixapis.alerts.v3.TopBucket.value:type_name -> google.protobuf.DoubleValue
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_init()
}
func file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_init() {
	if File_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto = out.File
	file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_depIdxs = nil
}
