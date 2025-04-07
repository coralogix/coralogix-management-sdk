// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.26.1
// source: com/coralogixapis/alerts/v3/event/payload/metric/event_metric_more_than_usual_enriched.proto

package v3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type EventMetricMoreThanUsualEnriched struct {
	state                    protoimpl.MessageState    `protogen:"open.v1"`
	EventMetricMoreThanUsual *EventMetricMoreThanUsual `protobuf:"bytes,1,opt,name=event_metric_more_than_usual,json=eventMetricMoreThanUsual,proto3" json:"event_metric_more_than_usual,omitempty"`
	Prediction               *Prediction               `protobuf:"bytes,2,opt,name=prediction,proto3" json:"prediction,omitempty"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *EventMetricMoreThanUsualEnriched) Reset() {
	*x = EventMetricMoreThanUsualEnriched{}
	mi := &file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventMetricMoreThanUsualEnriched) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventMetricMoreThanUsualEnriched) ProtoMessage() {}

func (x *EventMetricMoreThanUsualEnriched) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventMetricMoreThanUsualEnriched.ProtoReflect.Descriptor instead.
func (*EventMetricMoreThanUsualEnriched) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_rawDescGZIP(), []int{0}
}

func (x *EventMetricMoreThanUsualEnriched) GetEventMetricMoreThanUsual() *EventMetricMoreThanUsual {
	if x != nil {
		return x.EventMetricMoreThanUsual
	}
	return nil
}

func (x *EventMetricMoreThanUsualEnriched) GetPrediction() *Prediction {
	if x != nil {
		return x.Prediction
	}
	return nil
}

type Prediction struct {
	state                  protoimpl.MessageState           `protogen:"open.v1"`
	PredictionTimestampMap map[string]*PredictionProperties `protobuf:"bytes,1,rep,name=prediction_timestamp_map,json=predictionTimestampMap,proto3" json:"prediction_timestamp_map,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *Prediction) Reset() {
	*x = Prediction{}
	mi := &file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Prediction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prediction) ProtoMessage() {}

func (x *Prediction) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prediction.ProtoReflect.Descriptor instead.
func (*Prediction) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_rawDescGZIP(), []int{1}
}

func (x *Prediction) GetPredictionTimestampMap() map[string]*PredictionProperties {
	if x != nil {
		return x.PredictionTimestampMap
	}
	return nil
}

type PredictionProperties struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	LowerLimit    *wrapperspb.DoubleValue `protobuf:"bytes,1,opt,name=lower_limit,json=lowerLimit,proto3" json:"lower_limit,omitempty"`
	UpperLimit    *wrapperspb.DoubleValue `protobuf:"bytes,2,opt,name=upper_limit,json=upperLimit,proto3" json:"upper_limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PredictionProperties) Reset() {
	*x = PredictionProperties{}
	mi := &file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PredictionProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictionProperties) ProtoMessage() {}

func (x *PredictionProperties) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictionProperties.ProtoReflect.Descriptor instead.
func (*PredictionProperties) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_rawDescGZIP(), []int{2}
}

func (x *PredictionProperties) GetLowerLimit() *wrapperspb.DoubleValue {
	if x != nil {
		return x.LowerLimit
	}
	return nil
}

func (x *PredictionProperties) GetUpperLimit() *wrapperspb.DoubleValue {
	if x != nil {
		return x.UpperLimit
	}
	return nil
}

var File_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto protoreflect.FileDescriptor

const file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_rawDesc = "" +
	"\n" +
	"\\com/coralogixapis/alerts/v3/event/payload/metric/event_metric_more_than_usual_enriched.proto\x12\x1bcom.coralogixapis.alerts.v3\x1aScom/coralogixapis/alerts/v3/event/payload/metric/event_metric_more_than_usual.proto\x1a\x1egoogle/protobuf/wrappers.proto\"\xe2\x01\n" +
	" EventMetricMoreThanUsualEnriched\x12u\n" +
	"\x1cevent_metric_more_than_usual\x18\x01 \x01(\v25.com.coralogixapis.alerts.v3.EventMetricMoreThanUsualR\x18eventMetricMoreThanUsual\x12G\n" +
	"\n" +
	"prediction\x18\x02 \x01(\v2'.com.coralogixapis.alerts.v3.PredictionR\n" +
	"prediction\"\x89\x02\n" +
	"\n" +
	"Prediction\x12}\n" +
	"\x18prediction_timestamp_map\x18\x01 \x03(\v2C.com.coralogixapis.alerts.v3.Prediction.PredictionTimestampMapEntryR\x16predictionTimestampMap\x1a|\n" +
	"\x1bPredictionTimestampMapEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12G\n" +
	"\x05value\x18\x02 \x01(\v21.com.coralogixapis.alerts.v3.PredictionPropertiesR\x05value:\x028\x01\"\x94\x01\n" +
	"\x14PredictionProperties\x12=\n" +
	"\vlower_limit\x18\x01 \x01(\v2\x1c.google.protobuf.DoubleValueR\n" +
	"lowerLimit\x12=\n" +
	"\vupper_limit\x18\x02 \x01(\v2\x1c.google.protobuf.DoubleValueR\n" +
	"upperLimitb\x06proto3"

var (
	file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_goTypes = []any{
	(*EventMetricMoreThanUsualEnriched)(nil), // 0: com.coralogixapis.alerts.v3.EventMetricMoreThanUsualEnriched
	(*Prediction)(nil),                       // 1: com.coralogixapis.alerts.v3.Prediction
	(*PredictionProperties)(nil),             // 2: com.coralogixapis.alerts.v3.PredictionProperties
	nil,                                      // 3: com.coralogixapis.alerts.v3.Prediction.PredictionTimestampMapEntry
	(*EventMetricMoreThanUsual)(nil),         // 4: com.coralogixapis.alerts.v3.EventMetricMoreThanUsual
	(*wrapperspb.DoubleValue)(nil),           // 5: google.protobuf.DoubleValue
}
var file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_depIdxs = []int32{
	4, // 0: com.coralogixapis.alerts.v3.EventMetricMoreThanUsualEnriched.event_metric_more_than_usual:type_name -> com.coralogixapis.alerts.v3.EventMetricMoreThanUsual
	1, // 1: com.coralogixapis.alerts.v3.EventMetricMoreThanUsualEnriched.prediction:type_name -> com.coralogixapis.alerts.v3.Prediction
	3, // 2: com.coralogixapis.alerts.v3.Prediction.prediction_timestamp_map:type_name -> com.coralogixapis.alerts.v3.Prediction.PredictionTimestampMapEntry
	5, // 3: com.coralogixapis.alerts.v3.PredictionProperties.lower_limit:type_name -> google.protobuf.DoubleValue
	5, // 4: com.coralogixapis.alerts.v3.PredictionProperties.upper_limit:type_name -> google.protobuf.DoubleValue
	2, // 5: com.coralogixapis.alerts.v3.Prediction.PredictionTimestampMapEntry.value:type_name -> com.coralogixapis.alerts.v3.PredictionProperties
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_init()
}
func file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_init() {
	if File_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto = out.File
	file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_depIdxs = nil
}
