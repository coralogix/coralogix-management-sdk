// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: com/coralogixapis/alerting/alert_scheduler_rule_protobuf/v1/alert_scheduler_rule_timeframe.proto

package v1

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

type DurationFrequency int32

const (
	DurationFrequency_DURATION_FREQUENCY_UNSPECIFIED DurationFrequency = 0
	DurationFrequency_DURATION_FREQUENCY_MINUTE      DurationFrequency = 1
	DurationFrequency_DURATION_FREQUENCY_HOUR        DurationFrequency = 2
	DurationFrequency_DURATION_FREQUENCY_DAY         DurationFrequency = 3
)

// Enum value maps for DurationFrequency.
var (
	DurationFrequency_name = map[int32]string{
		0: "DURATION_FREQUENCY_UNSPECIFIED",
		1: "DURATION_FREQUENCY_MINUTE",
		2: "DURATION_FREQUENCY_HOUR",
		3: "DURATION_FREQUENCY_DAY",
	}
	DurationFrequency_value = map[string]int32{
		"DURATION_FREQUENCY_UNSPECIFIED": 0,
		"DURATION_FREQUENCY_MINUTE":      1,
		"DURATION_FREQUENCY_HOUR":        2,
		"DURATION_FREQUENCY_DAY":         3,
	}
)

func (x DurationFrequency) Enum() *DurationFrequency {
	p := new(DurationFrequency)
	*p = x
	return p
}

func (x DurationFrequency) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DurationFrequency) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_enumTypes[0].Descriptor()
}

func (DurationFrequency) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_enumTypes[0]
}

func (x DurationFrequency) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DurationFrequency.Descriptor instead.
func (DurationFrequency) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_rawDescGZIP(), []int{0}
}

type Timeframe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime string `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"` // Timeframe start time: The point in the time(date-time) when the rule will start to be active.
	// Types that are assignable to Until:
	//
	//	*Timeframe_EndTime
	//	*Timeframe_Duration
	Until    isTimeframe_Until `protobuf_oneof:"until"`
	Timezone string            `protobuf:"bytes,4,opt,name=timezone,proto3" json:"timezone,omitempty"` // Timeframe timezone: The rule will be active according to a specific timezone.
}

func (x *Timeframe) Reset() {
	*x = Timeframe{}
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Timeframe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timeframe) ProtoMessage() {}

func (x *Timeframe) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timeframe.ProtoReflect.Descriptor instead.
func (*Timeframe) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_rawDescGZIP(), []int{0}
}

func (x *Timeframe) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (m *Timeframe) GetUntil() isTimeframe_Until {
	if m != nil {
		return m.Until
	}
	return nil
}

func (x *Timeframe) GetEndTime() string {
	if x, ok := x.GetUntil().(*Timeframe_EndTime); ok {
		return x.EndTime
	}
	return ""
}

func (x *Timeframe) GetDuration() *Duration {
	if x, ok := x.GetUntil().(*Timeframe_Duration); ok {
		return x.Duration
	}
	return nil
}

func (x *Timeframe) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

type isTimeframe_Until interface {
	isTimeframe_Until()
}

type Timeframe_EndTime struct {
	EndTime string `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3,oneof"` // Timeframe end time: The point in the time(date-time) when the rule will finish to be active.
}

type Timeframe_Duration struct {
	Duration *Duration `protobuf:"bytes,3,opt,name=duration,proto3,oneof"` // Timeframe duration: The duration interval of the rule activation.
}

func (*Timeframe_EndTime) isTimeframe_Until() {}

func (*Timeframe_Duration) isTimeframe_Until() {}

type Duration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ForOver   int32             `protobuf:"varint,1,opt,name=for_over,json=forOver,proto3" json:"for_over,omitempty"`                                                                         // Duration for_over: The duration interval.
	Frequency DurationFrequency `protobuf:"varint,2,opt,name=frequency,proto3,enum=com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.DurationFrequency" json:"frequency,omitempty"` // Duration frequency: The duration frequency types (minute hour or day).
}

func (x *Duration) Reset() {
	*x = Duration{}
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Duration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Duration) ProtoMessage() {}

func (x *Duration) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Duration.ProtoReflect.Descriptor instead.
func (*Duration) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_rawDescGZIP(), []int{1}
}

func (x *Duration) GetForOver() int32 {
	if x != nil {
		return x.ForOver
	}
	return 0
}

func (x *Duration) GetFrequency() DurationFrequency {
	if x != nil {
		return x.Frequency
	}
	return DurationFrequency_DURATION_FREQUENCY_UNSPECIFIED
}

var File_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_rawDesc = []byte{
	0x0a, 0x60, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75,
	0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x3b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72,
	0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x22,
	0xd1, 0x01, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x08,
	0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x63, 0x0a, 0x08, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x00, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x75, 0x6e,
	0x74, 0x69, 0x6c, 0x22, 0x93, 0x01, 0x0a, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x19, 0x0a, 0x08, 0x66, 0x6f, 0x72, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x66, 0x6f, 0x72, 0x4f, 0x76, 0x65, 0x72, 0x12, 0x6c, 0x0a, 0x09, 0x66,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4e,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x09,
	0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x2a, 0x8f, 0x01, 0x0a, 0x11, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x22, 0x0a, 0x1e, 0x44, 0x55, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x52, 0x45, 0x51,
	0x55, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x55, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x46, 0x52, 0x45, 0x51, 0x55, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x4d, 0x49, 0x4e, 0x55, 0x54, 0x45,
	0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x44, 0x55, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46,
	0x52, 0x45, 0x51, 0x55, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x48, 0x4f, 0x55, 0x52, 0x10, 0x02, 0x12,
	0x1a, 0x0a, 0x16, 0x44, 0x55, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x52, 0x45, 0x51,
	0x55, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x44, 0x41, 0x59, 0x10, 0x03, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_rawDescData = file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_rawDesc
)

func file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_rawDescData)
	})
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_rawDescData
}

var file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_goTypes = []any{
	(DurationFrequency)(0), // 0: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.DurationFrequency
	(*Timeframe)(nil),      // 1: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Timeframe
	(*Duration)(nil),       // 2: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Duration
}
var file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_depIdxs = []int32{
	2, // 0: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Timeframe.duration:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Duration
	0, // 1: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Duration.frequency:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.DurationFrequency
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_init()
}
func file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_init() {
	if File_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto != nil {
		return
	}
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_msgTypes[0].OneofWrappers = []any{
		(*Timeframe_EndTime)(nil),
		(*Timeframe_Duration)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto = out.File
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_rawDesc = nil
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_goTypes = nil
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_depIdxs = nil
}
