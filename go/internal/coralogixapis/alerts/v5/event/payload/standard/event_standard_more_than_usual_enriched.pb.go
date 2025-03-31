// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/alerts/v5/event/payload/standard/event_standard_more_than_usual_enriched.proto

package standard

import (
	metric "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/alerts/v5/event/payload/metric"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type EventStandardMoreThanUsualEnriched struct {
	state                      protoimpl.MessageState      `protogen:"open.v1"`
	EventStandardMoreThanUsual *EventStandardMoreThanUsual `protobuf:"bytes,1,opt,name=event_standard_more_than_usual,json=eventStandardMoreThanUsual,proto3" json:"event_standard_more_than_usual,omitempty"`
	Prediction                 *metric.Prediction          `protobuf:"bytes,2,opt,name=prediction,proto3" json:"prediction,omitempty"`
	unknownFields              protoimpl.UnknownFields
	sizeCache                  protoimpl.SizeCache
}

func (x *EventStandardMoreThanUsualEnriched) Reset() {
	*x = EventStandardMoreThanUsualEnriched{}
	mi := &file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_more_than_usual_enriched_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventStandardMoreThanUsualEnriched) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventStandardMoreThanUsualEnriched) ProtoMessage() {}

func (x *EventStandardMoreThanUsualEnriched) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_more_than_usual_enriched_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventStandardMoreThanUsualEnriched.ProtoReflect.Descriptor instead.
func (*EventStandardMoreThanUsualEnriched) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_more_than_usual_enriched_proto_rawDescGZIP(), []int{0}
}

func (x *EventStandardMoreThanUsualEnriched) GetEventStandardMoreThanUsual() *EventStandardMoreThanUsual {
	if x != nil {
		return x.EventStandardMoreThanUsual
	}
	return nil
}

func (x *EventStandardMoreThanUsualEnriched) GetPrediction() *metric.Prediction {
	if x != nil {
		return x.Prediction
	}
	return nil
}

var File_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_more_than_usual_enriched_proto protoreflect.FileDescriptor

const file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_more_than_usual_enriched_proto_rawDesc = "" +
	"\n" +
	"`com/coralogixapis/alerts/v5/event/payload/standard/event_standard_more_than_usual_enriched.proto\x12\x1bcom.coralogixapis.alerts.v5\x1aWcom/coralogixapis/alerts/v5/event/payload/standard/event_standard_more_than_usual.proto\x1aVcom/coralogixapis/alerts/v5/event/payload/metric/event_metric_unusual_prediction.proto\"\xea\x01\n" +
	"\"EventStandardMoreThanUsualEnriched\x12{\n" +
	"\x1eevent_standard_more_than_usual\x18\x01 \x01(\v27.com.coralogixapis.alerts.v5.EventStandardMoreThanUsualR\x1aeventStandardMoreThanUsual\x12G\n" +
	"\n" +
	"prediction\x18\x02 \x01(\v2'.com.coralogixapis.alerts.v5.PredictionR\n" +
	"predictionb\x06proto3"

var (
	file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_more_than_usual_enriched_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_more_than_usual_enriched_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_more_than_usual_enriched_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_more_than_usual_enriched_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_more_than_usual_enriched_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_more_than_usual_enriched_proto_rawDesc), len(file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_more_than_usual_enriched_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_more_than_usual_enriched_proto_rawDescData
}

var file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_more_than_usual_enriched_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_more_than_usual_enriched_proto_goTypes = []any{
	(*EventStandardMoreThanUsualEnriched)(nil), // 0: com.coralogixapis.alerts.v5.EventStandardMoreThanUsualEnriched
	(*EventStandardMoreThanUsual)(nil),         // 1: com.coralogixapis.alerts.v5.EventStandardMoreThanUsual
	(*metric.Prediction)(nil),                  // 2: com.coralogixapis.alerts.v5.Prediction
}
var file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_more_than_usual_enriched_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.alerts.v5.EventStandardMoreThanUsualEnriched.event_standard_more_than_usual:type_name -> com.coralogixapis.alerts.v5.EventStandardMoreThanUsual
	2, // 1: com.coralogixapis.alerts.v5.EventStandardMoreThanUsualEnriched.prediction:type_name -> com.coralogixapis.alerts.v5.Prediction
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_more_than_usual_enriched_proto_init()
}
func file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_more_than_usual_enriched_proto_init() {
	if File_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_more_than_usual_enriched_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_more_than_usual_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_more_than_usual_enriched_proto_rawDesc), len(file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_more_than_usual_enriched_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_more_than_usual_enriched_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_more_than_usual_enriched_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_more_than_usual_enriched_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_more_than_usual_enriched_proto = out.File
	file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_more_than_usual_enriched_proto_goTypes = nil
	file_com_coralogixapis_alerts_v5_event_payload_standard_event_standard_more_than_usual_enriched_proto_depIdxs = nil
}
