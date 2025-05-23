// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/alerts/v5/event/payload/metric/event_metric_unusual_prediction.proto

package metric

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

type Prediction struct {
	state                  protoimpl.MessageState           `protogen:"open.v1"`
	PredictionTimestampMap map[string]*PredictionProperties `protobuf:"bytes,1,rep,name=prediction_timestamp_map,json=predictionTimestampMap,proto3" json:"prediction_timestamp_map,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *Prediction) Reset() {
	*x = Prediction{}
	mi := &file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_prediction_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Prediction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prediction) ProtoMessage() {}

func (x *Prediction) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_prediction_proto_msgTypes[0]
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
	return file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_prediction_proto_rawDescGZIP(), []int{0}
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
	mi := &file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_prediction_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PredictionProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictionProperties) ProtoMessage() {}

func (x *PredictionProperties) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_prediction_proto_msgTypes[1]
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
	return file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_prediction_proto_rawDescGZIP(), []int{1}
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

var File_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_prediction_proto protoreflect.FileDescriptor

const file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_prediction_proto_rawDesc = "" +
	"\n" +
	"Vcom/coralogixapis/alerts/v5/event/payload/metric/event_metric_unusual_prediction.proto\x12\x1bcom.coralogixapis.alerts.v5\x1a\x1egoogle/protobuf/wrappers.proto\"\x89\x02\n" +
	"\n" +
	"Prediction\x12}\n" +
	"\x18prediction_timestamp_map\x18\x01 \x03(\v2C.com.coralogixapis.alerts.v5.Prediction.PredictionTimestampMapEntryR\x16predictionTimestampMap\x1a|\n" +
	"\x1bPredictionTimestampMapEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12G\n" +
	"\x05value\x18\x02 \x01(\v21.com.coralogixapis.alerts.v5.PredictionPropertiesR\x05value:\x028\x01\"\x94\x01\n" +
	"\x14PredictionProperties\x12=\n" +
	"\vlower_limit\x18\x01 \x01(\v2\x1c.google.protobuf.DoubleValueR\n" +
	"lowerLimit\x12=\n" +
	"\vupper_limit\x18\x02 \x01(\v2\x1c.google.protobuf.DoubleValueR\n" +
	"upperLimitb\x06proto3"

var (
	file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_prediction_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_prediction_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_prediction_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_prediction_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_prediction_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_prediction_proto_rawDesc), len(file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_prediction_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_prediction_proto_rawDescData
}

var file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_prediction_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_prediction_proto_goTypes = []any{
	(*Prediction)(nil),             // 0: com.coralogixapis.alerts.v5.Prediction
	(*PredictionProperties)(nil),   // 1: com.coralogixapis.alerts.v5.PredictionProperties
	nil,                            // 2: com.coralogixapis.alerts.v5.Prediction.PredictionTimestampMapEntry
	(*wrapperspb.DoubleValue)(nil), // 3: google.protobuf.DoubleValue
}
var file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_prediction_proto_depIdxs = []int32{
	2, // 0: com.coralogixapis.alerts.v5.Prediction.prediction_timestamp_map:type_name -> com.coralogixapis.alerts.v5.Prediction.PredictionTimestampMapEntry
	3, // 1: com.coralogixapis.alerts.v5.PredictionProperties.lower_limit:type_name -> google.protobuf.DoubleValue
	3, // 2: com.coralogixapis.alerts.v5.PredictionProperties.upper_limit:type_name -> google.protobuf.DoubleValue
	1, // 3: com.coralogixapis.alerts.v5.Prediction.PredictionTimestampMapEntry.value:type_name -> com.coralogixapis.alerts.v5.PredictionProperties
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_prediction_proto_init()
}
func file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_prediction_proto_init() {
	if File_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_prediction_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_prediction_proto_rawDesc), len(file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_prediction_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_prediction_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_prediction_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_prediction_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_prediction_proto = out.File
	file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_prediction_proto_goTypes = nil
	file_com_coralogixapis_alerts_v5_event_payload_metric_event_metric_unusual_prediction_proto_depIdxs = nil
}
