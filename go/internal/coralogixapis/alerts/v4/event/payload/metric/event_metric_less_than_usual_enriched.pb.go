// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.26.1
// source: com/coralogixapis/alerts/v4/event/payload/metric/event_metric_less_than_usual_enriched.proto

package metric

import (
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

type EventMetricLessThanUsualEnriched struct {
	state                    protoimpl.MessageState    `protogen:"open.v1"`
	EventMetricLessThanUsual *EventMetricLessThanUsual `protobuf:"bytes,1,opt,name=event_metric_less_than_usual,json=eventMetricLessThanUsual,proto3" json:"event_metric_less_than_usual,omitempty"`
	Prediction               *Prediction               `protobuf:"bytes,2,opt,name=prediction,proto3" json:"prediction,omitempty"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *EventMetricLessThanUsualEnriched) Reset() {
	*x = EventMetricLessThanUsualEnriched{}
	mi := &file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_enriched_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventMetricLessThanUsualEnriched) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventMetricLessThanUsualEnriched) ProtoMessage() {}

func (x *EventMetricLessThanUsualEnriched) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_enriched_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventMetricLessThanUsualEnriched.ProtoReflect.Descriptor instead.
func (*EventMetricLessThanUsualEnriched) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_enriched_proto_rawDescGZIP(), []int{0}
}

func (x *EventMetricLessThanUsualEnriched) GetEventMetricLessThanUsual() *EventMetricLessThanUsual {
	if x != nil {
		return x.EventMetricLessThanUsual
	}
	return nil
}

func (x *EventMetricLessThanUsualEnriched) GetPrediction() *Prediction {
	if x != nil {
		return x.Prediction
	}
	return nil
}

var File_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_enriched_proto protoreflect.FileDescriptor

const file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_enriched_proto_rawDesc = "" +
	"\n" +
	"\\com/coralogixapis/alerts/v4/event/payload/metric/event_metric_less_than_usual_enriched.proto\x12\x1bcom.coralogixapis.alerts.v4\x1aScom/coralogixapis/alerts/v4/event/payload/metric/event_metric_less_than_usual.proto\x1aVcom/coralogixapis/alerts/v4/event/payload/metric/event_metric_unusual_prediction.proto\"\xe2\x01\n" +
	" EventMetricLessThanUsualEnriched\x12u\n" +
	"\x1cevent_metric_less_than_usual\x18\x01 \x01(\v25.com.coralogixapis.alerts.v4.EventMetricLessThanUsualR\x18eventMetricLessThanUsual\x12G\n" +
	"\n" +
	"prediction\x18\x02 \x01(\v2'.com.coralogixapis.alerts.v4.PredictionR\n" +
	"predictionb\x06proto3"

var (
	file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_enriched_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_enriched_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_enriched_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_enriched_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_enriched_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_enriched_proto_rawDesc), len(file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_enriched_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_enriched_proto_rawDescData
}

var file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_enriched_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_enriched_proto_goTypes = []any{
	(*EventMetricLessThanUsualEnriched)(nil), // 0: com.coralogixapis.alerts.v4.EventMetricLessThanUsualEnriched
	(*EventMetricLessThanUsual)(nil),         // 1: com.coralogixapis.alerts.v4.EventMetricLessThanUsual
	(*Prediction)(nil),                       // 2: com.coralogixapis.alerts.v4.Prediction
}
var file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_enriched_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.alerts.v4.EventMetricLessThanUsualEnriched.event_metric_less_than_usual:type_name -> com.coralogixapis.alerts.v4.EventMetricLessThanUsual
	2, // 1: com.coralogixapis.alerts.v4.EventMetricLessThanUsualEnriched.prediction:type_name -> com.coralogixapis.alerts.v4.Prediction
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_enriched_proto_init()
}
func file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_enriched_proto_init() {
	if File_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_enriched_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_proto_init()
	file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_unusual_prediction_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_enriched_proto_rawDesc), len(file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_enriched_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_enriched_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_enriched_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_enriched_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_enriched_proto = out.File
	file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_enriched_proto_goTypes = nil
	file_com_coralogixapis_alerts_v4_event_payload_metric_event_metric_less_than_usual_enriched_proto_depIdxs = nil
}
