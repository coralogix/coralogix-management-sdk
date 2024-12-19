// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.26.1
// source: com/coralogixapis/alerts/v3/event/payload/new_value/event_new_value.proto

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

var file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_rawDesc = []byte{
	0x0a, 0x49, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x6e, 0x65, 0x77, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x65, 0x77, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xda, 0x01, 0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x65, 0x77, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x12, 0x4b, 0x0a, 0x16, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x5f, 0x6c, 0x6f,
	0x67, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x43, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x4c, 0x6f, 0x67, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x14, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x4c, 0x6f, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3f, 0x0a,
	0x0d, 0x6c, 0x6f, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0c, 0x6c, 0x6f, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_rawDescData = file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_goTypes = []any{
	(*EventNewValue)(nil),          // 0: com.coralogixapis.alerts.v3.EventNewValue
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
	(*CoralogixLogMetadata)(nil),   // 2: CoralogixLogMetadata
	(*timestamppb.Timestamp)(nil),  // 3: google.protobuf.Timestamp
}
var file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.alerts.v3.EventNewValue.log_record:type_name -> google.protobuf.StringValue
	2, // 1: com.coralogixapis.alerts.v3.EventNewValue.coralogix_log_metadata:type_name -> CoralogixLogMetadata
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
			RawDescriptor: file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_rawDesc,
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
	file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_event_payload_new_value_event_new_value_proto_depIdxs = nil
}
