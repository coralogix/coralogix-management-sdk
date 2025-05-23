// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/alerts/v3/event/payload/standard/event_standard_more_than.proto

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

type EventStandardMoreThan struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	HitValue      *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=hit_value,json=hitValue,proto3" json:"hit_value,omitempty"`
	FromTimestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=from_timestamp,json=fromTimestamp,proto3" json:"from_timestamp,omitempty"`
	ToTimestamp   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=to_timestamp,json=toTimestamp,proto3" json:"to_timestamp,omitempty"`
	Severity      *AlertDefSeverity      `protobuf:"bytes,4,opt,name=severity,proto3" json:"severity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EventStandardMoreThan) Reset() {
	*x = EventStandardMoreThan{}
	mi := &file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_more_than_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventStandardMoreThan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventStandardMoreThan) ProtoMessage() {}

func (x *EventStandardMoreThan) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_more_than_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventStandardMoreThan.ProtoReflect.Descriptor instead.
func (*EventStandardMoreThan) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_more_than_proto_rawDescGZIP(), []int{0}
}

func (x *EventStandardMoreThan) GetHitValue() *wrapperspb.Int64Value {
	if x != nil {
		return x.HitValue
	}
	return nil
}

func (x *EventStandardMoreThan) GetFromTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.FromTimestamp
	}
	return nil
}

func (x *EventStandardMoreThan) GetToTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.ToTimestamp
	}
	return nil
}

func (x *EventStandardMoreThan) GetSeverity() *AlertDefSeverity {
	if x != nil {
		return x.Severity
	}
	return nil
}

var File_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_more_than_proto protoreflect.FileDescriptor

const file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_more_than_proto_rawDesc = "" +
	"\n" +
	"Qcom/coralogixapis/alerts/v3/event/payload/standard/event_standard_more_than.proto\x12\x1bcom.coralogixapis.alerts.v3\x1a\x1egoogle/protobuf/wrappers.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a4com/coralogixapis/alerts/v3/alert_def_severity.proto\"\x9e\x02\n" +
	"\x15EventStandardMoreThan\x128\n" +
	"\thit_value\x18\x01 \x01(\v2\x1b.google.protobuf.Int64ValueR\bhitValue\x12A\n" +
	"\x0efrom_timestamp\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\rfromTimestamp\x12=\n" +
	"\fto_timestamp\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\vtoTimestamp\x12I\n" +
	"\bseverity\x18\x04 \x01(\v2-.com.coralogixapis.alerts.v3.AlertDefSeverityR\bseverityb\x06proto3"

var (
	file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_more_than_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_more_than_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_more_than_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_more_than_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_more_than_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_more_than_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_more_than_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_more_than_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_more_than_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_more_than_proto_goTypes = []any{
	(*EventStandardMoreThan)(nil), // 0: com.coralogixapis.alerts.v3.EventStandardMoreThan
	(*wrapperspb.Int64Value)(nil), // 1: google.protobuf.Int64Value
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*AlertDefSeverity)(nil),      // 3: com.coralogixapis.alerts.v3.AlertDefSeverity
}
var file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_more_than_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.alerts.v3.EventStandardMoreThan.hit_value:type_name -> google.protobuf.Int64Value
	2, // 1: com.coralogixapis.alerts.v3.EventStandardMoreThan.from_timestamp:type_name -> google.protobuf.Timestamp
	2, // 2: com.coralogixapis.alerts.v3.EventStandardMoreThan.to_timestamp:type_name -> google.protobuf.Timestamp
	3, // 3: com.coralogixapis.alerts.v3.EventStandardMoreThan.severity:type_name -> com.coralogixapis.alerts.v3.AlertDefSeverity
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_more_than_proto_init()
}
func file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_more_than_proto_init() {
	if File_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_more_than_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_alert_def_severity_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_more_than_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_more_than_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_more_than_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_more_than_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_more_than_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_more_than_proto = out.File
	file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_more_than_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_event_payload_standard_event_standard_more_than_proto_depIdxs = nil
}
