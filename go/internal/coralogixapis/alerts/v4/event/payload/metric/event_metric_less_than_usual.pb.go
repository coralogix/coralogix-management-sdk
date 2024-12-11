// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.1
// source: com/coralogixapis/alerts/v4/event/payload/metric/event_metric_less_than_usual.proto

package metric

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

type EventMetricLessThanUsual struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LessThanUsual *EventUnusualPayload `protobuf:"bytes,1,opt,name=less_than_usual,json=lessThanUsual,proto3" json:"less_than_usual,omitempty"`
}

func (x *EventMetricLessThanUsual) Reset() {
	*x = EventMetricLessThanUsual{}
	mi := &file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventMetricLessThanUsual) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventMetricLessThanUsual) ProtoMessage() {}

func (x *EventMetricLessThanUsual) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventMetricLessThanUsual.ProtoReflect.Descriptor instead.
func (*EventMetricLessThanUsual) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_proto_rawDescGZIP(), []int{0}
}

func (x *EventMetricLessThanUsual) GetLessThanUsual() *EventUnusualPayload {
	if x != nil {
		return x.LessThanUsual
	}
	return nil
}

var File_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_proto_rawDesc = []byte{
	0x0a, 0x53, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x34, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f,
	0x6c, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x68, 0x61, 0x6e, 0x5f, 0x75, 0x73, 0x75, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e,
	0x76, 0x34, 0x1a, 0x4c, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x34, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x6e, 0x75, 0x73, 0x75,
	0x61, 0x6c, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x74, 0x0a, 0x18, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4c,
	0x65, 0x73, 0x73, 0x54, 0x68, 0x61, 0x6e, 0x55, 0x73, 0x75, 0x61, 0x6c, 0x12, 0x58, 0x0a, 0x0f,
	0x6c, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x68, 0x61, 0x6e, 0x5f, 0x75, 0x73, 0x75, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x2e, 0x76, 0x34, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x55, 0x6e, 0x75, 0x73, 0x75, 0x61, 0x6c,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0d, 0x6c, 0x65, 0x73, 0x73, 0x54, 0x68, 0x61,
	0x6e, 0x55, 0x73, 0x75, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_proto_rawDescData = file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_proto_rawDesc
)

func file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_proto_rawDescData
}

var file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_proto_goTypes = []any{
	(*EventMetricLessThanUsual)(nil), // 0: com.coralogixapis.alerts.v4.EventMetricLessThanUsual
	(*EventUnusualPayload)(nil),      // 1: com.coralogixapis.alerts.v4.EventUnusualPayload
}
var file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.alerts.v4.EventMetricLessThanUsual.less_than_usual:type_name -> com.coralogixapis.alerts.v4.EventUnusualPayload
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_proto_init()
}
func file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_proto_init() {
	if File_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v4_event_payload_metric_event_unusual_payload_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_proto = out.File
	file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_proto_goTypes = nil
	file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_proto_depIdxs = nil
}
