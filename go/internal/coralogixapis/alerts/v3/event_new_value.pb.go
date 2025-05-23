// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/alerts/v3/event/payload/new_value/event_new_value.proto

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

type EventNewValue struct {
	state                protoimpl.MessageState  `protogen:"open.v1"`
	LogRecord            *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=log_record,json=logRecord,proto3" json:"log_record,omitempty"`
	CoralogixLogMetadata *CoralogixLogMetadata   `protobuf:"bytes,2,opt,name=coralogix_log_metadata,json=coralogixLogMetadata,proto3" json:"coralogix_log_metadata,omitempty"`
	LogTimestamp         *timestamppb.Timestamp  `protobuf:"bytes,3,opt,name=log_timestamp,json=logTimestamp,proto3" json:"log_timestamp,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *EventNewValue) Reset() {
	*x = EventNewValue{}
	mi := &file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventNewValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventNewValue) ProtoMessage() {}

func (x *EventNewValue) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventNewValue.ProtoReflect.Descriptor instead.
func (*EventNewValue) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_rawDescGZIP(), []int{0}
}

func (x *EventNewValue) GetLogRecord() *wrapperspb.StringValue {
	if x != nil {
		return x.LogRecord
	}
	return nil
}

func (x *EventNewValue) GetCoralogixLogMetadata() *CoralogixLogMetadata {
	if x != nil {
		return x.CoralogixLogMetadata
	}
	return nil
}

func (x *EventNewValue) GetLogTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.LogTimestamp
	}
	return nil
}

var File_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto protoreflect.FileDescriptor

const file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_rawDesc = "" +
	"\n" +
	"Icom/coralogixapis/alerts/v3/event/payload/new_value/event_new_value.proto\x12\x1bcom.coralogixapis.alerts.v3\x1a\x1egoogle/protobuf/wrappers.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a6com/coralogixapis/alerts/v3/event/payload/common.proto\"\xf6\x01\n" +
	"\rEventNewValue\x12;\n" +
	"\n" +
	"log_record\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueR\tlogRecord\x12g\n" +
	"\x16coralogix_log_metadata\x18\x02 \x01(\v21.com.coralogixapis.alerts.v3.CoralogixLogMetadataR\x14coralogixLogMetadata\x12?\n" +
	"\rlog_timestamp\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\flogTimestampb\x06proto3"

var (
	file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_goTypes = []any{
	(*EventNewValue)(nil),          // 0: com.coralogixapis.alerts.v3.EventNewValue
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
	(*CoralogixLogMetadata)(nil),   // 2: com.coralogixapis.alerts.v3.CoralogixLogMetadata
	(*timestamppb.Timestamp)(nil),  // 3: google.protobuf.Timestamp
}
var file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.alerts.v3.EventNewValue.log_record:type_name -> google.protobuf.StringValue
	2, // 1: com.coralogixapis.alerts.v3.EventNewValue.coralogix_log_metadata:type_name -> com.coralogixapis.alerts.v3.CoralogixLogMetadata
	3, // 2: com.coralogixapis.alerts.v3.EventNewValue.log_timestamp:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_init() }
func file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_init() {
	if File_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_event_payload_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto = out.File
	file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_depIdxs = nil
}
