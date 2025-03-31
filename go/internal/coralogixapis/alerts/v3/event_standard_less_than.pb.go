// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.26.1
// source: com/coralogixapis/alerts/v3/event/payload/standard/event_standard_less_than.proto

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

type EventStandardLessThan struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	HitValue      *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=hit_value,json=hitValue,proto3" json:"hit_value,omitempty"`
	FromTimestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=from_timestamp,json=fromTimestamp,proto3" json:"from_timestamp,omitempty"`
	ToTimestamp   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=to_timestamp,json=toTimestamp,proto3" json:"to_timestamp,omitempty"`
	IsDeadman     *wrapperspb.BoolValue  `protobuf:"bytes,4,opt,name=is_deadman,json=isDeadman,proto3" json:"is_deadman,omitempty"`
	Severity      *AlertDefSeverity      `protobuf:"bytes,5,opt,name=severity,proto3" json:"severity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EventStandardLessThan) Reset() {
	*x = EventStandardLessThan{}
	mi := &file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_less_than_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventStandardLessThan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventStandardLessThan) ProtoMessage() {}

func (x *EventStandardLessThan) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_less_than_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventStandardLessThan.ProtoReflect.Descriptor instead.
func (*EventStandardLessThan) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_less_than_proto_rawDescGZIP(), []int{0}
}

func (x *EventStandardLessThan) GetHitValue() *wrapperspb.Int64Value {
	if x != nil {
		return x.HitValue
	}
	return nil
}

func (x *EventStandardLessThan) GetFromTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.FromTimestamp
	}
	return nil
}

func (x *EventStandardLessThan) GetToTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.ToTimestamp
	}
	return nil
}

func (x *EventStandardLessThan) GetIsDeadman() *wrapperspb.BoolValue {
	if x != nil {
		return x.IsDeadman
	}
	return nil
}

func (x *EventStandardLessThan) GetSeverity() *AlertDefSeverity {
	if x != nil {
		return x.Severity
	}
	return nil
}

var File_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_less_than_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_less_than_proto_rawDesc = []byte{
	0x0a, 0x51, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x73, 0x74, 0x61, 0x6e,
	0x64, 0x61, 0x72, 0x64, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x6e, 0x64,
	0x61, 0x72, 0x64, 0x5f, 0x6c, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x68, 0x61, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x34, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x02, 0x0a, 0x15, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x4c, 0x65, 0x73, 0x73, 0x54, 0x68, 0x61,
	0x6e, 0x12, 0x38, 0x0a, 0x09, 0x68, 0x69, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x08, 0x68, 0x69, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x66,
	0x72, 0x6f, 0x6d, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3d,
	0x0a, 0x0c, 0x74, 0x6f, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0b, 0x74, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x39, 0x0a,
	0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x61, 0x64, 0x6d, 0x61, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x69,
	0x73, 0x44, 0x65, 0x61, 0x64, 0x6d, 0x61, 0x6e, 0x12, 0x49, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x44, 0x65,
	0x66, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72,
	0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_less_than_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_less_than_proto_rawDescData = file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_less_than_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_less_than_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_less_than_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_less_than_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_less_than_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_less_than_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_less_than_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_less_than_proto_goTypes = []any{
	(*EventStandardLessThan)(nil), // 0: com.coralogixapis.alerts.v3.EventStandardLessThan
	(*wrapperspb.Int64Value)(nil), // 1: google.protobuf.Int64Value
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*wrapperspb.BoolValue)(nil),  // 3: google.protobuf.BoolValue
	(*AlertDefSeverity)(nil),      // 4: com.coralogixapis.alerts.v3.AlertDefSeverity
}
var file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_less_than_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.alerts.v3.EventStandardLessThan.hit_value:type_name -> google.protobuf.Int64Value
	2, // 1: com.coralogixapis.alerts.v3.EventStandardLessThan.from_timestamp:type_name -> google.protobuf.Timestamp
	2, // 2: com.coralogixapis.alerts.v3.EventStandardLessThan.to_timestamp:type_name -> google.protobuf.Timestamp
	3, // 3: com.coralogixapis.alerts.v3.EventStandardLessThan.is_deadman:type_name -> google.protobuf.BoolValue
	4, // 4: com.coralogixapis.alerts.v3.EventStandardLessThan.severity:type_name -> com.coralogixapis.alerts.v3.AlertDefSeverity
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_less_than_proto_init()
}
func file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_less_than_proto_init() {
	if File_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_less_than_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_alert_def_severity_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_less_than_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_less_than_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_less_than_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_less_than_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_less_than_proto = out.File
	file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_less_than_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_less_than_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_less_than_proto_depIdxs = nil
}
