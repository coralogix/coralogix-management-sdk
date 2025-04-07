// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.26.1
// source: com/coralogixapis/alerts/v3/event/alert_event_payload.proto

package v3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

type AlertEventPayload struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Payload       *structpb.Struct        `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	PayloadType   *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=payload_type,json=payloadType,proto3" json:"payload_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AlertEventPayload) Reset() {
	*x = AlertEventPayload{}
	mi := &file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlertEventPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertEventPayload) ProtoMessage() {}

func (x *AlertEventPayload) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertEventPayload.ProtoReflect.Descriptor instead.
func (*AlertEventPayload) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_rawDescGZIP(), []int{0}
}

func (x *AlertEventPayload) GetPayload() *structpb.Struct {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *AlertEventPayload) GetPayloadType() *wrapperspb.StringValue {
	if x != nil {
		return x.PayloadType
	}
	return nil
}

var File_com_coralogixapis_alerts_v3_event_alert_event_payload_proto protoreflect.FileDescriptor

const file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_rawDesc = "" +
	"\n" +
	";com/coralogixapis/alerts/v3/event/alert_event_payload.proto\x12\x1bcom.coralogixapis.alerts.v3\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1egoogle/protobuf/wrappers.proto\"\x87\x01\n" +
	"\x11AlertEventPayload\x121\n" +
	"\apayload\x18\x01 \x01(\v2\x17.google.protobuf.StructR\apayload\x12?\n" +
	"\fpayload_type\x18\x02 \x01(\v2\x1c.google.protobuf.StringValueR\vpayloadTypeb\x06proto3"

var (
	file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_goTypes = []any{
	(*AlertEventPayload)(nil),      // 0: com.coralogixapis.alerts.v3.AlertEventPayload
	(*structpb.Struct)(nil),        // 1: google.protobuf.Struct
	(*wrapperspb.StringValue)(nil), // 2: google.protobuf.StringValue
}
var file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.alerts.v3.AlertEventPayload.payload:type_name -> google.protobuf.Struct
	2, // 1: com.coralogixapis.alerts.v3.AlertEventPayload.payload_type:type_name -> google.protobuf.StringValue
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_init() }
func file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_init() {
	if File_com_coralogixapis_alerts_v3_event_alert_event_payload_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_event_alert_event_payload_proto = out.File
	file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_depIdxs = nil
}
