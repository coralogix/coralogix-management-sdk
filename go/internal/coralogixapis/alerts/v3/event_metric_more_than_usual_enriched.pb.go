// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: com/coralogixapis/alerts/v3/event/payload/metric/event_metric_more_than_usual_enriched.proto

package v3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type EventMetricMoreThanUsualEnriched struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventMetricMoreThanUsual *EventMetricMoreThanUsual `protobuf:"bytes,1,opt,name=event_metric_more_than_usual,json=eventMetricMoreThanUsual,proto3" json:"event_metric_more_than_usual,omitempty"`
	Prediction               *Prediction               `protobuf:"bytes,2,opt,name=prediction,proto3" json:"prediction,omitempty"`
}

func (x *EventMetricMoreThanUsualEnriched) Reset() {
	*x = EventMetricMoreThanUsualEnriched{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventMetricMoreThanUsualEnriched) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventMetricMoreThanUsualEnriched) ProtoMessage() {}

func (x *EventMetricMoreThanUsualEnriched) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PredictionTimestampMap map[string]*PredictionProperties `protobuf:"bytes,1,rep,name=prediction_timestamp_map,json=predictionTimestampMap,proto3" json:"prediction_timestamp_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Prediction) Reset() {
	*x = Prediction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prediction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prediction) ProtoMessage() {}

func (x *Prediction) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LowerLimit *wrapperspb.DoubleValue `protobuf:"bytes,1,opt,name=lower_limit,json=lowerLimit,proto3" json:"lower_limit,omitempty"`
	UpperLimit *wrapperspb.DoubleValue `protobuf:"bytes,2,opt,name=upper_limit,json=upperLimit,proto3" json:"upper_limit,omitempty"`
}

func (x *PredictionProperties) Reset() {
	*x = PredictionProperties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictionProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictionProperties) ProtoMessage() {}

func (x *PredictionProperties) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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

var file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_rawDesc = []byte{
	0x0a, 0x5c, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f,
	0x6d, 0x6f, 0x72, 0x65, 0x5f, 0x74, 0x68, 0x61, 0x6e, 0x5f, 0x75, 0x73, 0x75, 0x61, 0x6c, 0x5f,
	0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x53, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x6d, 0x6f, 0x72, 0x65, 0x5f,
	0x74, 0x68, 0x61, 0x6e, 0x5f, 0x75, 0x73, 0x75, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xe2, 0x01, 0x0a, 0x20, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x4d, 0x6f, 0x72, 0x65, 0x54, 0x68, 0x61, 0x6e, 0x55, 0x73, 0x75, 0x61, 0x6c, 0x45, 0x6e, 0x72,
	0x69, 0x63, 0x68, 0x65, 0x64, 0x12, 0x75, 0x0a, 0x1c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x6d, 0x6f, 0x72, 0x65, 0x5f, 0x74, 0x68, 0x61, 0x6e, 0x5f,
	0x75, 0x73, 0x75, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x4d, 0x6f, 0x72, 0x65, 0x54, 0x68, 0x61, 0x6e, 0x55, 0x73, 0x75,
	0x61, 0x6c, 0x52, 0x18, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4d,
	0x6f, 0x72, 0x65, 0x54, 0x68, 0x61, 0x6e, 0x55, 0x73, 0x75, 0x61, 0x6c, 0x12, 0x47, 0x0a, 0x0a,
	0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x50,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x64, 0x69,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x89, 0x02, 0x0a, 0x0a, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x7d, 0x0a, 0x18, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x6d, 0x61, 0x70,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x73, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x16, 0x70, 0x72, 0x65,
	0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x4d, 0x61, 0x70, 0x1a, 0x7c, 0x0a, 0x1b, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x47, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76,
	0x33, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x94, 0x01, 0x0a, 0x14, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x0b, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x3d, 0x0a, 0x0b, 0x75, 0x70, 0x70,
	0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x75, 0x70,
	0x70, 0x65, 0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_rawDescData = file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_rawDescData)
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
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*EventMetricMoreThanUsualEnriched); i {
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
		file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Prediction); i {
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
		file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PredictionProperties); i {
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
			RawDescriptor: file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_rawDesc,
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
	file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_event_payload_metric_event_metric_more_than_usual_enriched_proto_depIdxs = nil
}
