// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/alerts/v3/event/payload/tracing/event_tracing_immediate.proto

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

type EventTracingImmediate struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	TimeRange     *SpanTimeRange            `protobuf:"bytes,1,opt,name=time_range,json=timeRange,proto3" json:"time_range,omitempty"`
	TraceIds      []*wrapperspb.StringValue `protobuf:"bytes,2,rep,name=trace_ids,json=traceIds,proto3" json:"trace_ids,omitempty"`
	Span          *Span                     `protobuf:"bytes,3,opt,name=span,proto3" json:"span,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EventTracingImmediate) Reset() {
	*x = EventTracingImmediate{}
	mi := &file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventTracingImmediate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventTracingImmediate) ProtoMessage() {}

func (x *EventTracingImmediate) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventTracingImmediate.ProtoReflect.Descriptor instead.
func (*EventTracingImmediate) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_rawDescGZIP(), []int{0}
}

func (x *EventTracingImmediate) GetTimeRange() *SpanTimeRange {
	if x != nil {
		return x.TimeRange
	}
	return nil
}

func (x *EventTracingImmediate) GetTraceIds() []*wrapperspb.StringValue {
	if x != nil {
		return x.TraceIds
	}
	return nil
}

func (x *EventTracingImmediate) GetSpan() *Span {
	if x != nil {
		return x.Span
	}
	return nil
}

type Span struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	TraceId       *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	StartTime     *timestamppb.Timestamp  `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Span) Reset() {
	*x = Span{}
	mi := &file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Span) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Span) ProtoMessage() {}

func (x *Span) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Span.ProtoReflect.Descriptor instead.
func (*Span) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_rawDescGZIP(), []int{1}
}

func (x *Span) GetTraceId() *wrapperspb.StringValue {
	if x != nil {
		return x.TraceId
	}
	return nil
}

func (x *Span) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

var File_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto protoreflect.FileDescriptor

const file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_rawDesc = "" +
	"\n" +
	"Ocom/coralogixapis/alerts/v3/event/payload/tracing/event_tracing_immediate.proto\x12\x1bcom.coralogixapis.alerts.v3\x1a\x1egoogle/protobuf/wrappers.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1aGcom/coralogixapis/alerts/v3/event/payload/tracing/span_time_range.proto\"\xd4\x01\n" +
	"\x15EventTracingImmediate\x12I\n" +
	"\n" +
	"time_range\x18\x01 \x01(\v2*.com.coralogixapis.alerts.v3.SpanTimeRangeR\ttimeRange\x129\n" +
	"\ttrace_ids\x18\x02 \x03(\v2\x1c.google.protobuf.StringValueR\btraceIds\x125\n" +
	"\x04span\x18\x03 \x01(\v2!.com.coralogixapis.alerts.v3.SpanR\x04span\"z\n" +
	"\x04Span\x127\n" +
	"\btrace_id\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueR\atraceId\x129\n" +
	"\n" +
	"start_time\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\tstartTimeb\x06proto3"

var (
	file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_goTypes = []any{
	(*EventTracingImmediate)(nil),  // 0: com.coralogixapis.alerts.v3.EventTracingImmediate
	(*Span)(nil),                   // 1: com.coralogixapis.alerts.v3.Span
	(*SpanTimeRange)(nil),          // 2: com.coralogixapis.alerts.v3.SpanTimeRange
	(*wrapperspb.StringValue)(nil), // 3: google.protobuf.StringValue
	(*timestamppb.Timestamp)(nil),  // 4: google.protobuf.Timestamp
}
var file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_depIdxs = []int32{
	2, // 0: com.coralogixapis.alerts.v3.EventTracingImmediate.time_range:type_name -> com.coralogixapis.alerts.v3.SpanTimeRange
	3, // 1: com.coralogixapis.alerts.v3.EventTracingImmediate.trace_ids:type_name -> google.protobuf.StringValue
	1, // 2: com.coralogixapis.alerts.v3.EventTracingImmediate.span:type_name -> com.coralogixapis.alerts.v3.Span
	3, // 3: com.coralogixapis.alerts.v3.Span.trace_id:type_name -> google.protobuf.StringValue
	4, // 4: com.coralogixapis.alerts.v3.Span.start_time:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_init()
}
func file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_init() {
	if File_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_event_payload_tracing_span_time_range_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto = out.File
	file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_depIdxs = nil
}
