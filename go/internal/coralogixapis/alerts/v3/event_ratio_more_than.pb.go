// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: com/coralogixapis/alerts/v3/event/payload/ratio/event_ratio_more_than.proto

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

type EventRatioMoreThan struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	NumeratorHitValue   *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=numerator_hit_value,json=numeratorHitValue,proto3" json:"numerator_hit_value,omitempty"`
	DenominatorHitValue *wrapperspb.Int64Value `protobuf:"bytes,2,opt,name=denominator_hit_value,json=denominatorHitValue,proto3" json:"denominator_hit_value,omitempty"`
	// Types that are valid to be assigned to ResultHitValue:
	//
	//	*EventRatioMoreThan_Numeric
	//	*EventRatioMoreThan_Special
	ResultHitValue isEventRatioMoreThan_ResultHitValue `protobuf_oneof:"result_hit_value"`
	FromTimestamp  *timestamppb.Timestamp              `protobuf:"bytes,4,opt,name=from_timestamp,json=fromTimestamp,proto3" json:"from_timestamp,omitempty"`
	ToTimestamp    *timestamppb.Timestamp              `protobuf:"bytes,5,opt,name=to_timestamp,json=toTimestamp,proto3" json:"to_timestamp,omitempty"`
	Severity       *AlertDefSeverity                   `protobuf:"bytes,6,opt,name=severity,proto3" json:"severity,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *EventRatioMoreThan) Reset() {
	*x = EventRatioMoreThan{}
	mi := &file_com_coralogixapis_alerts_v3_event_payload_ratio_event_ratio_more_than_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventRatioMoreThan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventRatioMoreThan) ProtoMessage() {}

func (x *EventRatioMoreThan) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_payload_ratio_event_ratio_more_than_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventRatioMoreThan.ProtoReflect.Descriptor instead.
func (*EventRatioMoreThan) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_payload_ratio_event_ratio_more_than_proto_rawDescGZIP(), []int{0}
}

func (x *EventRatioMoreThan) GetNumeratorHitValue() *wrapperspb.Int64Value {
	if x != nil {
		return x.NumeratorHitValue
	}
	return nil
}

func (x *EventRatioMoreThan) GetDenominatorHitValue() *wrapperspb.Int64Value {
	if x != nil {
		return x.DenominatorHitValue
	}
	return nil
}

func (x *EventRatioMoreThan) GetResultHitValue() isEventRatioMoreThan_ResultHitValue {
	if x != nil {
		return x.ResultHitValue
	}
	return nil
}

func (x *EventRatioMoreThan) GetNumeric() *wrapperspb.FloatValue {
	if x != nil {
		if x, ok := x.ResultHitValue.(*EventRatioMoreThan_Numeric); ok {
			return x.Numeric
		}
	}
	return nil
}

func (x *EventRatioMoreThan) GetSpecial() SpecialRatioValues {
	if x != nil {
		if x, ok := x.ResultHitValue.(*EventRatioMoreThan_Special); ok {
			return x.Special
		}
	}
	return SpecialRatioValues_SPECIAL_RATIO_VALUES_INFINITY_OR_UNSPECIFIED
}

func (x *EventRatioMoreThan) GetFromTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.FromTimestamp
	}
	return nil
}

func (x *EventRatioMoreThan) GetToTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.ToTimestamp
	}
	return nil
}

func (x *EventRatioMoreThan) GetSeverity() *AlertDefSeverity {
	if x != nil {
		return x.Severity
	}
	return nil
}

type isEventRatioMoreThan_ResultHitValue interface {
	isEventRatioMoreThan_ResultHitValue()
}

type EventRatioMoreThan_Numeric struct {
	Numeric *wrapperspb.FloatValue `protobuf:"bytes,100,opt,name=numeric,proto3,oneof"`
}

type EventRatioMoreThan_Special struct {
	Special SpecialRatioValues `protobuf:"varint,101,opt,name=special,proto3,enum=com.coralogixapis.alerts.v3.SpecialRatioValues,oneof"`
}

func (*EventRatioMoreThan_Numeric) isEventRatioMoreThan_ResultHitValue() {}

func (*EventRatioMoreThan_Special) isEventRatioMoreThan_ResultHitValue() {}

var File_com_coralogixapis_alerts_v3_event_payload_ratio_event_ratio_more_than_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_event_payload_ratio_event_ratio_more_than_proto_rawDesc = []byte{
	0x0a, 0x4b, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x5f, 0x6d, 0x6f,
	0x72, 0x65, 0x5f, 0x74, 0x68, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x44, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x5f, 0x73, 0x70, 0x65,
	0x63, 0x69, 0x61, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x34, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x04, 0x0a, 0x12, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x4d, 0x6f, 0x72, 0x65, 0x54, 0x68, 0x61, 0x6e, 0x12, 0x4b,
	0x0a, 0x13, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x68, 0x69, 0x74, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x11, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x48, 0x69, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x4f, 0x0a, 0x15, 0x64,
	0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x68, 0x69, 0x74, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x13, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x6e,
	0x61, 0x74, 0x6f, 0x72, 0x48, 0x69, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x37, 0x0a, 0x07,
	0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x07, 0x6e, 0x75,
	0x6d, 0x65, 0x72, 0x69, 0x63, 0x12, 0x4b, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c,
	0x18, 0x65, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x73, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x69,
	0x6f, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x48, 0x00, 0x52, 0x07, 0x73, 0x70, 0x65, 0x63, 0x69,
	0x61, 0x6c, 0x12, 0x41, 0x0a, 0x0e, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3d, 0x0a, 0x0c, 0x74, 0x6f, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x74, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x49, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x73, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x44, 0x65, 0x66, 0x53, 0x65, 0x76,
	0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x42,
	0x12, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x68, 0x69, 0x74, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_event_payload_ratio_event_ratio_more_than_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_event_payload_ratio_event_ratio_more_than_proto_rawDescData = file_com_coralogixapis_alerts_v3_event_payload_ratio_event_ratio_more_than_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_event_payload_ratio_event_ratio_more_than_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_event_payload_ratio_event_ratio_more_than_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_event_payload_ratio_event_ratio_more_than_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_event_payload_ratio_event_ratio_more_than_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_event_payload_ratio_event_ratio_more_than_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_event_payload_ratio_event_ratio_more_than_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_alerts_v3_event_payload_ratio_event_ratio_more_than_proto_goTypes = []any{
	(*EventRatioMoreThan)(nil),    // 0: com.coralogixapis.alerts.v3.EventRatioMoreThan
	(*wrapperspb.Int64Value)(nil), // 1: google.protobuf.Int64Value
	(*wrapperspb.FloatValue)(nil), // 2: google.protobuf.FloatValue
	(SpecialRatioValues)(0),       // 3: com.coralogixapis.alerts.v3.SpecialRatioValues
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*AlertDefSeverity)(nil),      // 5: com.coralogixapis.alerts.v3.AlertDefSeverity
}
var file_com_coralogixapis_alerts_v3_event_payload_ratio_event_ratio_more_than_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.alerts.v3.EventRatioMoreThan.numerator_hit_value:type_name -> google.protobuf.Int64Value
	1, // 1: com.coralogixapis.alerts.v3.EventRatioMoreThan.denominator_hit_value:type_name -> google.protobuf.Int64Value
	2, // 2: com.coralogixapis.alerts.v3.EventRatioMoreThan.numeric:type_name -> google.protobuf.FloatValue
	3, // 3: com.coralogixapis.alerts.v3.EventRatioMoreThan.special:type_name -> com.coralogixapis.alerts.v3.SpecialRatioValues
	4, // 4: com.coralogixapis.alerts.v3.EventRatioMoreThan.from_timestamp:type_name -> google.protobuf.Timestamp
	4, // 5: com.coralogixapis.alerts.v3.EventRatioMoreThan.to_timestamp:type_name -> google.protobuf.Timestamp
	5, // 6: com.coralogixapis.alerts.v3.EventRatioMoreThan.severity:type_name -> com.coralogixapis.alerts.v3.AlertDefSeverity
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_alerts_v3_event_payload_ratio_event_ratio_more_than_proto_init() }
func file_com_coralogixapis_alerts_v3_event_payload_ratio_event_ratio_more_than_proto_init() {
	if File_com_coralogixapis_alerts_v3_event_payload_ratio_event_ratio_more_than_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_event_payload_ratio_special_values_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_severity_proto_init()
	file_com_coralogixapis_alerts_v3_event_payload_ratio_event_ratio_more_than_proto_msgTypes[0].OneofWrappers = []any{
		(*EventRatioMoreThan_Numeric)(nil),
		(*EventRatioMoreThan_Special)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_event_payload_ratio_event_ratio_more_than_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_event_payload_ratio_event_ratio_more_than_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_event_payload_ratio_event_ratio_more_than_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_event_payload_ratio_event_ratio_more_than_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_event_payload_ratio_event_ratio_more_than_proto = out.File
	file_com_coralogixapis_alerts_v3_event_payload_ratio_event_ratio_more_than_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_event_payload_ratio_event_ratio_more_than_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_event_payload_ratio_event_ratio_more_than_proto_depIdxs = nil
}
