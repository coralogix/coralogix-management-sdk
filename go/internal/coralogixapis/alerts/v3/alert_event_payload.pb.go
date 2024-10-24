// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: com/coralogixapis/alerts/v3/event/alert_event_payload.proto

package v3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

type AlertEventPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload     *structpb.Struct        `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	PayloadType *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=payload_type,json=payloadType,proto3" json:"payload_type,omitempty"`
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

var file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x11, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x31,
	0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x3f, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_rawDescData = file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_rawDescData)
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
			RawDescriptor: file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_rawDesc,
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
	file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_event_alert_event_payload_proto_depIdxs = nil
}
