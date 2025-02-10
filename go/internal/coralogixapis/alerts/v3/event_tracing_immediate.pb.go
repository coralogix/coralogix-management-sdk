// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: com/coralogixapis/alerts/v3/event/payload/tracing/event_tracing_immediate.proto

package v3

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

var file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_rawDesc = []byte{
	0x0a, 0x4f, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x74, 0x72, 0x61, 0x63,
	0x69, 0x6e, 0x67, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e,
	0x67, 0x5f, 0x69, 0x6d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x47, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x69,
	0x6e, 0x67, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x01, 0x0a, 0x15, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x49, 0x6d, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x74, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x73, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x39, 0x0a,
	0x09, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x73, 0x12, 0x35, 0x0a, 0x04, 0x73, 0x70, 0x61, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x73, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x52, 0x04, 0x73, 0x70, 0x61, 0x6e, 0x22,
	0x7a, 0x0a, 0x04, 0x53, 0x70, 0x61, 0x6e, 0x12, 0x37, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_rawDescData = file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_rawDescData)
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
			RawDescriptor: file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_rawDesc,
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
	file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_immediate_proto_depIdxs = nil
}
