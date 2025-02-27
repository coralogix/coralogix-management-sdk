// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.26.1
// source: com/coralogixapis/alerts/v3/event/payload/metric/event_metric_more_than.proto

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

type EventMetricMoreThan struct {
	state                   protoimpl.MessageState `protogen:"open.v1"`
	PercentageOverThreshold *wrapperspb.FloatValue `protobuf:"bytes,1,opt,name=percentage_over_threshold,json=percentageOverThreshold,proto3" json:"percentage_over_threshold,omitempty"`
	AvgValueAcrossThreshold *wrapperspb.FloatValue `protobuf:"bytes,2,opt,name=avg_value_across_threshold,json=avgValueAcrossThreshold,proto3,oneof" json:"avg_value_across_threshold,omitempty"`
	MinValueAcrossThreshold *wrapperspb.FloatValue `protobuf:"bytes,3,opt,name=min_value_across_threshold,json=minValueAcrossThreshold,proto3,oneof" json:"min_value_across_threshold,omitempty"`
	MaxValueAcrossThreshold *wrapperspb.FloatValue `protobuf:"bytes,4,opt,name=max_value_across_threshold,json=maxValueAcrossThreshold,proto3,oneof" json:"max_value_across_threshold,omitempty"`
	FromTimestamp           *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=from_timestamp,json=fromTimestamp,proto3" json:"from_timestamp,omitempty"`
	ToTimestamp             *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=to_timestamp,json=toTimestamp,proto3" json:"to_timestamp,omitempty"`
	Severity                *AlertDefSeverity      `protobuf:"bytes,7,opt,name=severity,proto3" json:"severity,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *EventMetricMoreThan) Reset() {
	*x = EventMetricMoreThan{}
	mi := &file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventMetricMoreThan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventMetricMoreThan) ProtoMessage() {}

func (x *EventMetricMoreThan) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventMetricMoreThan.ProtoReflect.Descriptor instead.
func (*EventMetricMoreThan) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_proto_rawDescGZIP(), []int{0}
}

func (x *EventMetricMoreThan) GetPercentageOverThreshold() *wrapperspb.FloatValue {
	if x != nil {
		return x.PercentageOverThreshold
	}
	return nil
}

func (x *EventMetricMoreThan) GetAvgValueAcrossThreshold() *wrapperspb.FloatValue {
	if x != nil {
		return x.AvgValueAcrossThreshold
	}
	return nil
}

func (x *EventMetricMoreThan) GetMinValueAcrossThreshold() *wrapperspb.FloatValue {
	if x != nil {
		return x.MinValueAcrossThreshold
	}
	return nil
}

func (x *EventMetricMoreThan) GetMaxValueAcrossThreshold() *wrapperspb.FloatValue {
	if x != nil {
		return x.MaxValueAcrossThreshold
	}
	return nil
}

func (x *EventMetricMoreThan) GetFromTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.FromTimestamp
	}
	return nil
}

func (x *EventMetricMoreThan) GetToTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.ToTimestamp
	}
	return nil
}

func (x *EventMetricMoreThan) GetSeverity() *AlertDefSeverity {
	if x != nil {
		return x.Severity
	}
	return nil
}

var File_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_proto_rawDesc = []byte{
	0x0a, 0x4d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f,
	0x6d, 0x6f, 0x72, 0x65, 0x5f, 0x74, 0x68, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x5f, 0x64, 0x65, 0x66, 0x5f, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x05, 0x0a, 0x13, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x4d, 0x6f, 0x72, 0x65, 0x54, 0x68, 0x61, 0x6e, 0x12, 0x57, 0x0a, 0x19, 0x70,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x74,
	0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x17, 0x70, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x54, 0x68, 0x72, 0x65, 0x73,
	0x68, 0x6f, 0x6c, 0x64, 0x12, 0x5d, 0x0a, 0x1a, 0x61, 0x76, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x5f, 0x61, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f,
	0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x17, 0x61, 0x76, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x41, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x5d, 0x0a, 0x1a, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x5f, 0x61, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x48, 0x01, 0x52, 0x17, 0x6d, 0x69, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x41, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x5d, 0x0a, 0x1a, 0x6d, 0x61, 0x78, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f,
	0x61, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x48, 0x02, 0x52, 0x17, 0x6d, 0x61, 0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x41,
	0x63, 0x72, 0x6f, 0x73, 0x73, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x41, 0x0a, 0x0e, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x3d, 0x0a, 0x0c, 0x74, 0x6f, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x74, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x49, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x2e, 0x76, 0x33, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x44, 0x65, 0x66, 0x53, 0x65, 0x76, 0x65,
	0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x42, 0x1d,
	0x0a, 0x1b, 0x5f, 0x61, 0x76, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x61, 0x63, 0x72,
	0x6f, 0x73, 0x73, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x42, 0x1d, 0x0a,
	0x1b, 0x5f, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x61, 0x63, 0x72, 0x6f,
	0x73, 0x73, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x42, 0x1d, 0x0a, 0x1b,
	0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x61, 0x63, 0x72, 0x6f, 0x73,
	0x73, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_proto_rawDescData = file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_proto_goTypes = []any{
	(*EventMetricMoreThan)(nil),   // 0: com.coralogixapis.alerts.v3.EventMetricMoreThan
	(*wrapperspb.FloatValue)(nil), // 1: google.protobuf.FloatValue
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*AlertDefSeverity)(nil),      // 3: com.coralogixapis.alerts.v3.AlertDefSeverity
}
var file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.alerts.v3.EventMetricMoreThan.percentage_over_threshold:type_name -> google.protobuf.FloatValue
	1, // 1: com.coralogixapis.alerts.v3.EventMetricMoreThan.avg_value_across_threshold:type_name -> google.protobuf.FloatValue
	1, // 2: com.coralogixapis.alerts.v3.EventMetricMoreThan.min_value_across_threshold:type_name -> google.protobuf.FloatValue
	1, // 3: com.coralogixapis.alerts.v3.EventMetricMoreThan.max_value_across_threshold:type_name -> google.protobuf.FloatValue
	2, // 4: com.coralogixapis.alerts.v3.EventMetricMoreThan.from_timestamp:type_name -> google.protobuf.Timestamp
	2, // 5: com.coralogixapis.alerts.v3.EventMetricMoreThan.to_timestamp:type_name -> google.protobuf.Timestamp
	3, // 6: com.coralogixapis.alerts.v3.EventMetricMoreThan.severity:type_name -> com.coralogixapis.alerts.v3.AlertDefSeverity
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_proto_init()
}
func file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_proto_init() {
	if File_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_alert_def_severity_proto_init()
	file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_proto = out.File
	file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_proto_depIdxs = nil
}
