// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: com/coralogixapis/alerts/v3/event/payload/flow/event_flow_match.proto

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

type EventFlowMatch struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	From          *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To            *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Groups        []*FlowGroupEvents     `protobuf:"bytes,3,rep,name=groups,proto3" json:"groups,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EventFlowMatch) Reset() {
	*x = EventFlowMatch{}
	mi := &file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventFlowMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventFlowMatch) ProtoMessage() {}

func (x *EventFlowMatch) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventFlowMatch.ProtoReflect.Descriptor instead.
func (*EventFlowMatch) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto_rawDescGZIP(), []int{0}
}

func (x *EventFlowMatch) GetFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *EventFlowMatch) GetTo() *timestamppb.Timestamp {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *EventFlowMatch) GetGroups() []*FlowGroupEvents {
	if x != nil {
		return x.Groups
	}
	return nil
}

type FlowGroupEvents struct {
	state          protoimpl.MessageState    `protogen:"open.v1"`
	Events         []*AlertEventOut          `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	ResolvedAlerts []*wrapperspb.StringValue `protobuf:"bytes,2,rep,name=resolved_alerts,json=resolvedAlerts,proto3" json:"resolved_alerts,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *FlowGroupEvents) Reset() {
	*x = FlowGroupEvents{}
	mi := &file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FlowGroupEvents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowGroupEvents) ProtoMessage() {}

func (x *FlowGroupEvents) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowGroupEvents.ProtoReflect.Descriptor instead.
func (*FlowGroupEvents) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto_rawDescGZIP(), []int{1}
}

func (x *FlowGroupEvents) GetEvents() []*AlertEventOut {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *FlowGroupEvents) GetResolvedAlerts() []*wrapperspb.StringValue {
	if x != nil {
		return x.ResolvedAlerts
	}
	return nil
}

type AlertEventOut struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Timestamp     *timestamppb.Timestamp  `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	AlertId       *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=alert_id,json=alertId,proto3" json:"alert_id,omitempty"`
	EventId       *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AlertEventOut) Reset() {
	*x = AlertEventOut{}
	mi := &file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlertEventOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertEventOut) ProtoMessage() {}

func (x *AlertEventOut) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertEventOut.ProtoReflect.Descriptor instead.
func (*AlertEventOut) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto_rawDescGZIP(), []int{2}
}

func (x *AlertEventOut) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *AlertEventOut) GetAlertId() *wrapperspb.StringValue {
	if x != nil {
		return x.AlertId
	}
	return nil
}

func (x *AlertEventOut) GetEventId() *wrapperspb.StringValue {
	if x != nil {
		return x.EventId
	}
	return nil
}

var File_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto_rawDesc = []byte{
	0x0a, 0x45, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x66, 0x6c, 0x6f, 0x77,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x73, 0x2e, 0x76, 0x33, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x01, 0x0a, 0x0e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x46,
	0x6c, 0x6f, 0x77, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x2e, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x02, 0x74, 0x6f, 0x12, 0x44, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e,
	0x76, 0x33, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x9c, 0x01, 0x0a, 0x0f, 0x46,
	0x6c, 0x6f, 0x77, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x42,
	0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x45, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x5f, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x64, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x22, 0xbb, 0x01, 0x0a, 0x0d, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x37, 0x0a, 0x08, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x49, 0x64, 0x12, 0x37,
	0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto_rawDescData = file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto_goTypes = []any{
	(*EventFlowMatch)(nil),         // 0: com.coralogixapis.alerts.v3.EventFlowMatch
	(*FlowGroupEvents)(nil),        // 1: com.coralogixapis.alerts.v3.FlowGroupEvents
	(*AlertEventOut)(nil),          // 2: com.coralogixapis.alerts.v3.AlertEventOut
	(*timestamppb.Timestamp)(nil),  // 3: google.protobuf.Timestamp
	(*wrapperspb.StringValue)(nil), // 4: google.protobuf.StringValue
}
var file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto_depIdxs = []int32{
	3, // 0: com.coralogixapis.alerts.v3.EventFlowMatch.from:type_name -> google.protobuf.Timestamp
	3, // 1: com.coralogixapis.alerts.v3.EventFlowMatch.to:type_name -> google.protobuf.Timestamp
	1, // 2: com.coralogixapis.alerts.v3.EventFlowMatch.groups:type_name -> com.coralogixapis.alerts.v3.FlowGroupEvents
	2, // 3: com.coralogixapis.alerts.v3.FlowGroupEvents.events:type_name -> com.coralogixapis.alerts.v3.AlertEventOut
	4, // 4: com.coralogixapis.alerts.v3.FlowGroupEvents.resolved_alerts:type_name -> google.protobuf.StringValue
	3, // 5: com.coralogixapis.alerts.v3.AlertEventOut.timestamp:type_name -> google.protobuf.Timestamp
	4, // 6: com.coralogixapis.alerts.v3.AlertEventOut.alert_id:type_name -> google.protobuf.StringValue
	4, // 7: com.coralogixapis.alerts.v3.AlertEventOut.event_id:type_name -> google.protobuf.StringValue
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto_init() }
func file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto_init() {
	if File_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto = out.File
	file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto_depIdxs = nil
}
