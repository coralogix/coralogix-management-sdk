// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: com/coralogixapis/alerts/v3/event/payload/flow/event_flow.proto

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

type EventFlow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlowAlertsMatch []*EventFlowMatch `protobuf:"bytes,1,rep,name=flow_alerts_match,json=flowAlertsMatch,proto3" json:"flow_alerts_match,omitempty"`
}

func (x *EventFlow) Reset() {
	*x = EventFlow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventFlow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventFlow) ProtoMessage() {}

func (x *EventFlow) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventFlow.ProtoReflect.Descriptor instead.
func (*EventFlow) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_proto_rawDescGZIP(), []int{0}
}

func (x *EventFlow) GetFlowAlertsMatch() []*EventFlowMatch {
	if x != nil {
		return x.FlowAlertsMatch
	}
	return nil
}

var File_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x66, 0x6c, 0x6f, 0x77,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x45,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x46, 0x6c,
	0x6f, 0x77, 0x12, 0x57, 0x0a, 0x11, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x73, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x46, 0x6c, 0x6f, 0x77, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x0f, 0x66, 0x6c, 0x6f, 0x77,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_proto_rawDescData = file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_proto_goTypes = []any{
	(*EventFlow)(nil),      // 0: com.coralogixapis.alerts.v3.EventFlow
	(*EventFlowMatch)(nil), // 1: com.coralogixapis.alerts.v3.EventFlowMatch
}
var file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.alerts.v3.EventFlow.flow_alerts_match:type_name -> com.coralogixapis.alerts.v3.EventFlowMatch
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_proto_init() }
func file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_proto_init() {
	if File_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_match_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*EventFlow); i {
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
			RawDescriptor: file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_proto = out.File
	file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_event_payload_flow_event_flow_proto_depIdxs = nil
}
