// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: com/coralogixapis/alerts/v5/event/payload/standard/event_standard_unusual_payload.proto

package standard

import (
	metric "github.com/coralogix/coralogix-management-sdk/internal/coralogixapis/alerts/v5/event/payload/metric"
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

type EventStandardUnusualPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromTimestamp     *timestamppb.Timestamp        `protobuf:"bytes,1,opt,name=from_timestamp,json=fromTimestamp,proto3" json:"from_timestamp,omitempty"`
	ToTimestamp       *timestamppb.Timestamp        `protobuf:"bytes,2,opt,name=to_timestamp,json=toTimestamp,proto3" json:"to_timestamp,omitempty"`
	ForecastTimestamp *wrapperspb.UInt64Value       `protobuf:"bytes,3,opt,name=forecast_timestamp,json=forecastTimestamp,proto3" json:"forecast_timestamp,omitempty"`
	SearchDetails     *metric.TimeRangeWithInterval `protobuf:"bytes,4,opt,name=search_details,json=searchDetails,proto3" json:"search_details,omitempty"`
	Distance          *wrapperspb.DoubleValue       `protobuf:"bytes,5,opt,name=distance,proto3" json:"distance,omitempty"`
	SumHitCount       *wrapperspb.Int64Value        `protobuf:"bytes,6,opt,name=sum_hit_count,json=sumHitCount,proto3" json:"sum_hit_count,omitempty"`
}

func (x *EventStandardUnusualPayload) Reset() {
	*x = EventStandardUnusualPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_unusual_payload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventStandardUnusualPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventStandardUnusualPayload) ProtoMessage() {}

func (x *EventStandardUnusualPayload) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_unusual_payload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventStandardUnusualPayload.ProtoReflect.Descriptor instead.
func (*EventStandardUnusualPayload) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_unusual_payload_proto_rawDescGZIP(), []int{0}
}

func (x *EventStandardUnusualPayload) GetFromTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.FromTimestamp
	}
	return nil
}

func (x *EventStandardUnusualPayload) GetToTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.ToTimestamp
	}
	return nil
}

func (x *EventStandardUnusualPayload) GetForecastTimestamp() *wrapperspb.UInt64Value {
	if x != nil {
		return x.ForecastTimestamp
	}
	return nil
}

func (x *EventStandardUnusualPayload) GetSearchDetails() *metric.TimeRangeWithInterval {
	if x != nil {
		return x.SearchDetails
	}
	return nil
}

func (x *EventStandardUnusualPayload) GetDistance() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Distance
	}
	return nil
}

func (x *EventStandardUnusualPayload) GetSumHitCount() *wrapperspb.Int64Value {
	if x != nil {
		return x.SumHitCount
	}
	return nil
}

var File_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_unusual_payload_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_unusual_payload_proto_rawDesc = []byte{
	0x0a, 0x57, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x35, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x73, 0x74, 0x61, 0x6e,
	0x64, 0x61, 0x72, 0x64, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x6e, 0x64,
	0x61, 0x72, 0x64, 0x5f, 0x75, 0x6e, 0x75, 0x73, 0x75, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x73, 0x2e, 0x76, 0x35, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x53, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x73, 0x2f, 0x76, 0x35, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x75, 0x6e, 0x75, 0x73, 0x75, 0x61, 0x6c, 0x5f, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x03, 0x0a,
	0x1b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x55, 0x6e,
	0x75, 0x73, 0x75, 0x61, 0x6c, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x41, 0x0a, 0x0e,
	0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x3d, 0x0a, 0x0c, 0x74, 0x6f, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0b, 0x74, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x4b,
	0x0a, 0x12, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x11, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x61,
	0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x59, 0x0a, 0x0e, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76,
	0x35, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x0d, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x38, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x3f, 0x0a, 0x0d, 0x73, 0x75, 0x6d, 0x5f, 0x68, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x73, 0x75, 0x6d, 0x48, 0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_unusual_payload_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_unusual_payload_proto_rawDescData = file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_unusual_payload_proto_rawDesc
)

func file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_unusual_payload_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_unusual_payload_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_unusual_payload_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_unusual_payload_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_unusual_payload_proto_rawDescData
}

var file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_unusual_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_unusual_payload_proto_goTypes = []any{
	(*EventStandardUnusualPayload)(nil),  // 0: com.coralogixapis.alerts.v5.EventStandardUnusualPayload
	(*timestamppb.Timestamp)(nil),        // 1: google.protobuf.Timestamp
	(*wrapperspb.UInt64Value)(nil),       // 2: google.protobuf.UInt64Value
	(*metric.TimeRangeWithInterval)(nil), // 3: com.coralogixapis.alerts.v5.TimeRangeWithInterval
	(*wrapperspb.DoubleValue)(nil),       // 4: google.protobuf.DoubleValue
	(*wrapperspb.Int64Value)(nil),        // 5: google.protobuf.Int64Value
}
var file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_unusual_payload_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.alerts.v5.EventStandardUnusualPayload.from_timestamp:type_name -> google.protobuf.Timestamp
	1, // 1: com.coralogixapis.alerts.v5.EventStandardUnusualPayload.to_timestamp:type_name -> google.protobuf.Timestamp
	2, // 2: com.coralogixapis.alerts.v5.EventStandardUnusualPayload.forecast_timestamp:type_name -> google.protobuf.UInt64Value
	3, // 3: com.coralogixapis.alerts.v5.EventStandardUnusualPayload.search_details:type_name -> com.coralogixapis.alerts.v5.TimeRangeWithInterval
	4, // 4: com.coralogixapis.alerts.v5.EventStandardUnusualPayload.distance:type_name -> google.protobuf.DoubleValue
	5, // 5: com.coralogixapis.alerts.v5.EventStandardUnusualPayload.sum_hit_count:type_name -> google.protobuf.Int64Value
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_unusual_payload_proto_init()
}
func file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_unusual_payload_proto_init() {
	if File_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_unusual_payload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_unusual_payload_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*EventStandardUnusualPayload); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_unusual_payload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_unusual_payload_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_unusual_payload_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_unusual_payload_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_unusual_payload_proto = out.File
	file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_unusual_payload_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_unusual_payload_proto_goTypes = nil
	file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_unusual_payload_proto_depIdxs = nil
}
