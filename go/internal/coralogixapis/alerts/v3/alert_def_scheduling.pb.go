// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: com/coralogixapis/alerts/v3/alert_def_scheduling.proto

package v3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DayOfWeek int32

const (
	DayOfWeek_DAY_OF_WEEK_MONDAY_OR_UNSPECIFIED DayOfWeek = 0
	DayOfWeek_DAY_OF_WEEK_TUESDAY               DayOfWeek = 1
	DayOfWeek_DAY_OF_WEEK_WEDNESDAY             DayOfWeek = 2
	DayOfWeek_DAY_OF_WEEK_THURSDAY              DayOfWeek = 3
	DayOfWeek_DAY_OF_WEEK_FRIDAY                DayOfWeek = 4
	DayOfWeek_DAY_OF_WEEK_SATURDAY              DayOfWeek = 5
	DayOfWeek_DAY_OF_WEEK_SUNDAY                DayOfWeek = 6
)

// Enum value maps for DayOfWeek.
var (
	DayOfWeek_name = map[int32]string{
		0: "DAY_OF_WEEK_MONDAY_OR_UNSPECIFIED",
		1: "DAY_OF_WEEK_TUESDAY",
		2: "DAY_OF_WEEK_WEDNESDAY",
		3: "DAY_OF_WEEK_THURSDAY",
		4: "DAY_OF_WEEK_FRIDAY",
		5: "DAY_OF_WEEK_SATURDAY",
		6: "DAY_OF_WEEK_SUNDAY",
	}
	DayOfWeek_value = map[string]int32{
		"DAY_OF_WEEK_MONDAY_OR_UNSPECIFIED": 0,
		"DAY_OF_WEEK_TUESDAY":               1,
		"DAY_OF_WEEK_WEDNESDAY":             2,
		"DAY_OF_WEEK_THURSDAY":              3,
		"DAY_OF_WEEK_FRIDAY":                4,
		"DAY_OF_WEEK_SATURDAY":              5,
		"DAY_OF_WEEK_SUNDAY":                6,
	}
)

func (x DayOfWeek) Enum() *DayOfWeek {
	p := new(DayOfWeek)
	*p = x
	return p
}

func (x DayOfWeek) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DayOfWeek) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_enumTypes[0].Descriptor()
}

func (DayOfWeek) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_enumTypes[0]
}

func (x DayOfWeek) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DayOfWeek.Descriptor instead.
func (DayOfWeek) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_rawDescGZIP(), []int{0}
}

type ActivitySchedule struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DayOfWeek     []DayOfWeek            `protobuf:"varint,1,rep,packed,name=day_of_week,json=dayOfWeek,proto3,enum=com.coralogixapis.alerts.v3.DayOfWeek" json:"day_of_week,omitempty"`
	StartTime     *TimeOfDay             `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime       *TimeOfDay             `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ActivitySchedule) Reset() {
	*x = ActivitySchedule{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActivitySchedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivitySchedule) ProtoMessage() {}

func (x *ActivitySchedule) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivitySchedule.ProtoReflect.Descriptor instead.
func (*ActivitySchedule) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_rawDescGZIP(), []int{0}
}

func (x *ActivitySchedule) GetDayOfWeek() []DayOfWeek {
	if x != nil {
		return x.DayOfWeek
	}
	return nil
}

func (x *ActivitySchedule) GetStartTime() *TimeOfDay {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *ActivitySchedule) GetEndTime() *TimeOfDay {
	if x != nil {
		return x.EndTime
	}
	return nil
}

type TimeOfDay struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Hours of day in 24 hour format. Should be from 0 to 23.
	Hours int32 `protobuf:"varint,1,opt,name=hours,proto3" json:"hours,omitempty"`
	// Minutes of hour of day. Must be from 0 to 59.
	Minutes       int32 `protobuf:"varint,2,opt,name=minutes,proto3" json:"minutes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TimeOfDay) Reset() {
	*x = TimeOfDay{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimeOfDay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeOfDay) ProtoMessage() {}

func (x *TimeOfDay) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeOfDay.ProtoReflect.Descriptor instead.
func (*TimeOfDay) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_rawDescGZIP(), []int{1}
}

func (x *TimeOfDay) GetHours() int32 {
	if x != nil {
		return x.Hours
	}
	return 0
}

func (x *TimeOfDay) GetMinutes() int32 {
	if x != nil {
		return x.Minutes
	}
	return 0
}

var File_com_coralogixapis_alerts_v3_alert_def_scheduling_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_rawDesc = []byte{
	0x0a, 0x36, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x01, 0x0a, 0x10, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x64, 0x61,
	0x79, 0x5f, 0x6f, 0x66, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x61,
	0x79, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x52, 0x09, 0x64, 0x61, 0x79, 0x4f, 0x66, 0x57, 0x65,
	0x65, 0x6b, 0x12, 0x45, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x73, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x66, 0x44, 0x61, 0x79, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x66,
	0x44, 0x61, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x3b, 0x0a, 0x09,
	0x54, 0x69, 0x6d, 0x65, 0x4f, 0x66, 0x44, 0x61, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x6f, 0x75,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x2a, 0xca, 0x01, 0x0a, 0x09, 0x44, 0x61,
	0x79, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x12, 0x25, 0x0a, 0x21, 0x44, 0x41, 0x59, 0x5f, 0x4f,
	0x46, 0x5f, 0x57, 0x45, 0x45, 0x4b, 0x5f, 0x4d, 0x4f, 0x4e, 0x44, 0x41, 0x59, 0x5f, 0x4f, 0x52,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17,
	0x0a, 0x13, 0x44, 0x41, 0x59, 0x5f, 0x4f, 0x46, 0x5f, 0x57, 0x45, 0x45, 0x4b, 0x5f, 0x54, 0x55,
	0x45, 0x53, 0x44, 0x41, 0x59, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x41, 0x59, 0x5f, 0x4f,
	0x46, 0x5f, 0x57, 0x45, 0x45, 0x4b, 0x5f, 0x57, 0x45, 0x44, 0x4e, 0x45, 0x53, 0x44, 0x41, 0x59,
	0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x41, 0x59, 0x5f, 0x4f, 0x46, 0x5f, 0x57, 0x45, 0x45,
	0x4b, 0x5f, 0x54, 0x48, 0x55, 0x52, 0x53, 0x44, 0x41, 0x59, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12,
	0x44, 0x41, 0x59, 0x5f, 0x4f, 0x46, 0x5f, 0x57, 0x45, 0x45, 0x4b, 0x5f, 0x46, 0x52, 0x49, 0x44,
	0x41, 0x59, 0x10, 0x04, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x41, 0x59, 0x5f, 0x4f, 0x46, 0x5f, 0x57,
	0x45, 0x45, 0x4b, 0x5f, 0x53, 0x41, 0x54, 0x55, 0x52, 0x44, 0x41, 0x59, 0x10, 0x05, 0x12, 0x16,
	0x0a, 0x12, 0x44, 0x41, 0x59, 0x5f, 0x4f, 0x46, 0x5f, 0x57, 0x45, 0x45, 0x4b, 0x5f, 0x53, 0x55,
	0x4e, 0x44, 0x41, 0x59, 0x10, 0x06, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_rawDescData = file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_goTypes = []any{
	(DayOfWeek)(0),           // 0: com.coralogixapis.alerts.v3.DayOfWeek
	(*ActivitySchedule)(nil), // 1: com.coralogixapis.alerts.v3.ActivitySchedule
	(*TimeOfDay)(nil),        // 2: com.coralogixapis.alerts.v3.TimeOfDay
}
var file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_depIdxs = []int32{
	0, // 0: com.coralogixapis.alerts.v3.ActivitySchedule.day_of_week:type_name -> com.coralogixapis.alerts.v3.DayOfWeek
	2, // 1: com.coralogixapis.alerts.v3.ActivitySchedule.start_time:type_name -> com.coralogixapis.alerts.v3.TimeOfDay
	2, // 2: com.coralogixapis.alerts.v3.ActivitySchedule.end_time:type_name -> com.coralogixapis.alerts.v3.TimeOfDay
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_init() }
func file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_scheduling_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_scheduling_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_depIdxs = nil
}
