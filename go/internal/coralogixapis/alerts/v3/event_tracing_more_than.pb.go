// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: com/coralogixapis/alerts/v3/event/payload/tracing/event_tracing_more_than.proto

package v3

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

type EventTracingMoreThan struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CountersMap   map[uint32]uint32      `protobuf:"bytes,1,rep,name=counters_map,json=countersMap,proto3" json:"counters_map,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	TimeRange     *SpanTimeRange         `protobuf:"bytes,2,opt,name=time_range,json=timeRange,proto3" json:"time_range,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EventTracingMoreThan) Reset() {
	*x = EventTracingMoreThan{}
	mi := &file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_more_than_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventTracingMoreThan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventTracingMoreThan) ProtoMessage() {}

func (x *EventTracingMoreThan) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_more_than_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventTracingMoreThan.ProtoReflect.Descriptor instead.
func (*EventTracingMoreThan) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_more_than_proto_rawDescGZIP(), []int{0}
}

func (x *EventTracingMoreThan) GetCountersMap() map[uint32]uint32 {
	if x != nil {
		return x.CountersMap
	}
	return nil
}

func (x *EventTracingMoreThan) GetTimeRange() *SpanTimeRange {
	if x != nil {
		return x.TimeRange
	}
	return nil
}

var File_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_more_than_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_more_than_proto_rawDesc = []byte{
	0x0a, 0x4f, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x74, 0x72, 0x61, 0x63,
	0x69, 0x6e, 0x67, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e,
	0x67, 0x5f, 0x6d, 0x6f, 0x72, 0x65, 0x5f, 0x74, 0x68, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x47,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e,
	0x67, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x02, 0x0a, 0x14, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x72, 0x65, 0x54, 0x68, 0x61, 0x6e,
	0x12, 0x65, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x5f, 0x6d, 0x61, 0x70,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x73, 0x2e, 0x76, 0x33, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e,
	0x67, 0x4d, 0x6f, 0x72, 0x65, 0x54, 0x68, 0x61, 0x6e, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x73, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x73, 0x4d, 0x61, 0x70, 0x12, 0x49, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x1a, 0x3e, 0x0a, 0x10, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x4d, 0x61,
	0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_more_than_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_more_than_proto_rawDescData = file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_more_than_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_more_than_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_more_than_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_more_than_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_more_than_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_more_than_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_more_than_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_more_than_proto_goTypes = []any{
	(*EventTracingMoreThan)(nil), // 0: com.coralogixapis.alerts.v3.EventTracingMoreThan
	nil,                          // 1: com.coralogixapis.alerts.v3.EventTracingMoreThan.CountersMapEntry
	(*SpanTimeRange)(nil),        // 2: com.coralogixapis.alerts.v3.SpanTimeRange
}
var file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_more_than_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.alerts.v3.EventTracingMoreThan.counters_map:type_name -> com.coralogixapis.alerts.v3.EventTracingMoreThan.CountersMapEntry
	2, // 1: com.coralogixapis.alerts.v3.EventTracingMoreThan.time_range:type_name -> com.coralogixapis.alerts.v3.SpanTimeRange
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_more_than_proto_init()
}
func file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_more_than_proto_init() {
	if File_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_more_than_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_event_payload_tracing_span_time_range_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_more_than_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_more_than_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_more_than_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_more_than_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_more_than_proto = out.File
	file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_more_than_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_more_than_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_event_payload_tracing_event_tracing_more_than_proto_depIdxs = nil
}
