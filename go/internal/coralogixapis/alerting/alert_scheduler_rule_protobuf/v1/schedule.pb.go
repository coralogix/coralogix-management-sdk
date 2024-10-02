// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: com/coralogixapis/alerting/alert_scheduler_rule_protobuf/v1/schedule.proto

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

type ScheduleOperation int32

const (
	ScheduleOperation_SCHEDULE_OPERATION_UNSPECIFIED ScheduleOperation = 0
	ScheduleOperation_SCHEDULE_OPERATION_MUTE        ScheduleOperation = 1
	ScheduleOperation_SCHEDULE_OPERATION_ACTIVATE    ScheduleOperation = 2
)

// Enum value maps for ScheduleOperation.
var (
	ScheduleOperation_name = map[int32]string{
		0: "SCHEDULE_OPERATION_UNSPECIFIED",
		1: "SCHEDULE_OPERATION_MUTE",
		2: "SCHEDULE_OPERATION_ACTIVATE",
	}
	ScheduleOperation_value = map[string]int32{
		"SCHEDULE_OPERATION_UNSPECIFIED": 0,
		"SCHEDULE_OPERATION_MUTE":        1,
		"SCHEDULE_OPERATION_ACTIVATE":    2,
	}
)

func (x ScheduleOperation) Enum() *ScheduleOperation {
	p := new(ScheduleOperation)
	*p = x
	return p
}

func (x ScheduleOperation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ScheduleOperation) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_enumTypes[0].Descriptor()
}

func (ScheduleOperation) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_enumTypes[0]
}

func (x ScheduleOperation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ScheduleOperation.Descriptor instead.
func (ScheduleOperation) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_rawDescGZIP(), []int{0}
}

type Schedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScheduleOperation ScheduleOperation `protobuf:"varint,1,opt,name=schedule_operation,json=scheduleOperation,proto3,enum=com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.ScheduleOperation" json:"schedule_operation,omitempty"` // Rule schedule_operation: The rule operation mode (mute/active).
	// Types that are assignable to Scheduler:
	//
	//	*Schedule_OneTime
	//	*Schedule_Recurring
	Scheduler isSchedule_Scheduler `protobuf_oneof:"scheduler"`
}

func (x *Schedule) Reset() {
	*x = Schedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedule) ProtoMessage() {}

func (x *Schedule) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedule.ProtoReflect.Descriptor instead.
func (*Schedule) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_rawDescGZIP(), []int{0}
}

func (x *Schedule) GetScheduleOperation() ScheduleOperation {
	if x != nil {
		return x.ScheduleOperation
	}
	return ScheduleOperation_SCHEDULE_OPERATION_UNSPECIFIED
}

func (m *Schedule) GetScheduler() isSchedule_Scheduler {
	if m != nil {
		return m.Scheduler
	}
	return nil
}

func (x *Schedule) GetOneTime() *OneTime {
	if x, ok := x.GetScheduler().(*Schedule_OneTime); ok {
		return x.OneTime
	}
	return nil
}

func (x *Schedule) GetRecurring() *Recurring {
	if x, ok := x.GetScheduler().(*Schedule_Recurring); ok {
		return x.Recurring
	}
	return nil
}

type isSchedule_Scheduler interface {
	isSchedule_Scheduler()
}

type Schedule_OneTime struct {
	OneTime *OneTime `protobuf:"bytes,2,opt,name=one_time,json=oneTime,proto3,oneof"` // Schedule one_time: The scheduling definition will activate the rule for a specific timeframe.
}

type Schedule_Recurring struct {
	Recurring *Recurring `protobuf:"bytes,3,opt,name=recurring,proto3,oneof"` // Schedule recurring: The scheduling definition will activate the rule for a recurring timeframe.
}

func (*Schedule_OneTime) isSchedule_Scheduler() {}

func (*Schedule_Recurring) isSchedule_Scheduler() {}

type OneTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timeframe *Timeframe `protobuf:"bytes,1,opt,name=timeframe,proto3" json:"timeframe,omitempty"`
}

func (x *OneTime) Reset() {
	*x = OneTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneTime) ProtoMessage() {}

func (x *OneTime) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneTime.ProtoReflect.Descriptor instead.
func (*OneTime) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_rawDescGZIP(), []int{1}
}

func (x *OneTime) GetTimeframe() *Timeframe {
	if x != nil {
		return x.Timeframe
	}
	return nil
}

type Recurring struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Condition:
	//
	//	*Recurring_Always_
	//	*Recurring_Dynamic_
	Condition isRecurring_Condition `protobuf_oneof:"condition"`
}

func (x *Recurring) Reset() {
	*x = Recurring{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Recurring) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recurring) ProtoMessage() {}

func (x *Recurring) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recurring.ProtoReflect.Descriptor instead.
func (*Recurring) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_rawDescGZIP(), []int{2}
}

func (m *Recurring) GetCondition() isRecurring_Condition {
	if m != nil {
		return m.Condition
	}
	return nil
}

func (x *Recurring) GetAlways() *Recurring_Always {
	if x, ok := x.GetCondition().(*Recurring_Always_); ok {
		return x.Always
	}
	return nil
}

func (x *Recurring) GetDynamic() *Recurring_Dynamic {
	if x, ok := x.GetCondition().(*Recurring_Dynamic_); ok {
		return x.Dynamic
	}
	return nil
}

type isRecurring_Condition interface {
	isRecurring_Condition()
}

type Recurring_Always_ struct {
	Always *Recurring_Always `protobuf:"bytes,1,opt,name=always,proto3,oneof"`
}

type Recurring_Dynamic_ struct {
	Dynamic *Recurring_Dynamic `protobuf:"bytes,2,opt,name=dynamic,proto3,oneof"`
}

func (*Recurring_Always_) isRecurring_Condition() {}

func (*Recurring_Dynamic_) isRecurring_Condition() {}

type Daily struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Daily) Reset() {
	*x = Daily{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Daily) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Daily) ProtoMessage() {}

func (x *Daily) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Daily.ProtoReflect.Descriptor instead.
func (*Daily) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_rawDescGZIP(), []int{3}
}

type Weekly struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DaysOfWeek []int32 `protobuf:"varint,1,rep,packed,name=days_of_week,json=daysOfWeek,proto3" json:"days_of_week,omitempty"` // Dynamic Weekly days_of_week: The rule will be activated at weekly intervals on specific days in a week.
}

func (x *Weekly) Reset() {
	*x = Weekly{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Weekly) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Weekly) ProtoMessage() {}

func (x *Weekly) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Weekly.ProtoReflect.Descriptor instead.
func (*Weekly) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_rawDescGZIP(), []int{4}
}

func (x *Weekly) GetDaysOfWeek() []int32 {
	if x != nil {
		return x.DaysOfWeek
	}
	return nil
}

type Monthly struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DaysOfMonth []int32 `protobuf:"varint,1,rep,packed,name=days_of_month,json=daysOfMonth,proto3" json:"days_of_month,omitempty"` // Dynamic Monthly days_of_month: The rule will be activated at monthly intervals on specific days in a month.
}

func (x *Monthly) Reset() {
	*x = Monthly{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Monthly) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Monthly) ProtoMessage() {}

func (x *Monthly) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Monthly.ProtoReflect.Descriptor instead.
func (*Monthly) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_rawDescGZIP(), []int{5}
}

func (x *Monthly) GetDaysOfMonth() []int32 {
	if x != nil {
		return x.DaysOfMonth
	}
	return nil
}

type Recurring_Always struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Recurring_Always) Reset() {
	*x = Recurring_Always{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Recurring_Always) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recurring_Always) ProtoMessage() {}

func (x *Recurring_Always) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recurring_Always.ProtoReflect.Descriptor instead.
func (*Recurring_Always) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_rawDescGZIP(), []int{2, 0}
}

type Recurring_Dynamic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepeatEvery int32 `protobuf:"varint,1,opt,name=repeat_every,json=repeatEvery,proto3" json:"repeat_every,omitempty"` // Recurring Dynamic repeat_every: The rule will be activated in a recurring mode according to the interval.
	// Types that are assignable to Frequency:
	//
	//	*Recurring_Dynamic_Daily
	//	*Recurring_Dynamic_Weekly
	//	*Recurring_Dynamic_Monthly
	Frequency       isRecurring_Dynamic_Frequency `protobuf_oneof:"frequency"`
	Timeframe       *Timeframe                    `protobuf:"bytes,5,opt,name=timeframe,proto3" json:"timeframe,omitempty"`                                          // Recurring Dynamic timeframe: The rule will be activated in a recurring mode according to the specific timeframe.
	TerminationDate *string                       `protobuf:"bytes,6,opt,name=termination_date,json=terminationDate,proto3,oneof" json:"termination_date,omitempty"` // Recurring Dynamic termination_date: The rule will be terminated according to termination_date.
}

func (x *Recurring_Dynamic) Reset() {
	*x = Recurring_Dynamic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Recurring_Dynamic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recurring_Dynamic) ProtoMessage() {}

func (x *Recurring_Dynamic) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recurring_Dynamic.ProtoReflect.Descriptor instead.
func (*Recurring_Dynamic) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_rawDescGZIP(), []int{2, 1}
}

func (x *Recurring_Dynamic) GetRepeatEvery() int32 {
	if x != nil {
		return x.RepeatEvery
	}
	return 0
}

func (m *Recurring_Dynamic) GetFrequency() isRecurring_Dynamic_Frequency {
	if m != nil {
		return m.Frequency
	}
	return nil
}

func (x *Recurring_Dynamic) GetDaily() *Daily {
	if x, ok := x.GetFrequency().(*Recurring_Dynamic_Daily); ok {
		return x.Daily
	}
	return nil
}

func (x *Recurring_Dynamic) GetWeekly() *Weekly {
	if x, ok := x.GetFrequency().(*Recurring_Dynamic_Weekly); ok {
		return x.Weekly
	}
	return nil
}

func (x *Recurring_Dynamic) GetMonthly() *Monthly {
	if x, ok := x.GetFrequency().(*Recurring_Dynamic_Monthly); ok {
		return x.Monthly
	}
	return nil
}

func (x *Recurring_Dynamic) GetTimeframe() *Timeframe {
	if x != nil {
		return x.Timeframe
	}
	return nil
}

func (x *Recurring_Dynamic) GetTerminationDate() string {
	if x != nil && x.TerminationDate != nil {
		return *x.TerminationDate
	}
	return ""
}

type isRecurring_Dynamic_Frequency interface {
	isRecurring_Dynamic_Frequency()
}

type Recurring_Dynamic_Daily struct {
	Daily *Daily `protobuf:"bytes,2,opt,name=daily,proto3,oneof"`
}

type Recurring_Dynamic_Weekly struct {
	Weekly *Weekly `protobuf:"bytes,3,opt,name=weekly,proto3,oneof"`
}

type Recurring_Dynamic_Monthly struct {
	Monthly *Monthly `protobuf:"bytes,4,opt,name=monthly,proto3,oneof"`
}

func (*Recurring_Dynamic_Daily) isRecurring_Dynamic_Frequency() {}

func (*Recurring_Dynamic_Weekly) isRecurring_Dynamic_Frequency() {}

func (*Recurring_Dynamic_Monthly) isRecurring_Dynamic_Frequency() {}

var File_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_rawDesc = []byte{
	0x0a, 0x4a, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3b, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x1a, 0x60, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x66, 0x72, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x02, 0x0a, 0x08,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x7d, 0x0a, 0x12, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x4e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x61, 0x0a, 0x08, 0x6f, 0x6e, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x48,
	0x00, 0x52, 0x07, 0x6f, 0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x66, 0x0a, 0x09, 0x72, 0x65,
	0x63, 0x75, 0x72, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x46, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x75,
	0x72, 0x72, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x09, 0x72, 0x65, 0x63, 0x75, 0x72, 0x72, 0x69,
	0x6e, 0x67, 0x42, 0x0b, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x22,
	0x6f, 0x0a, 0x07, 0x4f, 0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x64, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x46, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x66, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x66, 0x72, 0x61, 0x6d, 0x65,
	0x22, 0xfb, 0x05, 0x0a, 0x09, 0x52, 0x65, 0x63, 0x75, 0x72, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x67,
	0x0a, 0x06, 0x61, 0x6c, 0x77, 0x61, 0x79, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4d,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63,
	0x75, 0x72, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x6c, 0x77, 0x61, 0x79, 0x73, 0x48, 0x00, 0x52,
	0x06, 0x61, 0x6c, 0x77, 0x61, 0x79, 0x73, 0x12, 0x6a, 0x0a, 0x07, 0x64, 0x79, 0x6e, 0x61, 0x6d,
	0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x75, 0x72, 0x72, 0x69, 0x6e, 0x67,
	0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x48, 0x00, 0x52, 0x07, 0x64, 0x79, 0x6e, 0x61,
	0x6d, 0x69, 0x63, 0x1a, 0x08, 0x0a, 0x06, 0x41, 0x6c, 0x77, 0x61, 0x79, 0x73, 0x1a, 0x81, 0x04,
	0x0a, 0x07, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x45, 0x76, 0x65, 0x72, 0x79, 0x12, 0x5a, 0x0a, 0x05,
	0x64, 0x61, 0x69, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x48,
	0x00, 0x52, 0x05, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x12, 0x5d, 0x0a, 0x06, 0x77, 0x65, 0x65, 0x6b,
	0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x48, 0x00, 0x52,
	0x06, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x12, 0x60, 0x0a, 0x07, 0x6d, 0x6f, 0x6e, 0x74, 0x68,
	0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x48, 0x00,
	0x52, 0x07, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x12, 0x64, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x12,
	0x2e, 0x0a, 0x10, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0f, 0x74, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x42,
	0x0b, 0x0a, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x13, 0x0a, 0x11,
	0x5f, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x0b, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x07,
	0x0a, 0x05, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x22, 0x2a, 0x0a, 0x06, 0x57, 0x65, 0x65, 0x6b, 0x6c,
	0x79, 0x12, 0x20, 0x0a, 0x0c, 0x64, 0x61, 0x79, 0x73, 0x5f, 0x6f, 0x66, 0x5f, 0x77, 0x65, 0x65,
	0x6b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0a, 0x64, 0x61, 0x79, 0x73, 0x4f, 0x66, 0x57,
	0x65, 0x65, 0x6b, 0x22, 0x2d, 0x0a, 0x07, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x12, 0x22,
	0x0a, 0x0d, 0x64, 0x61, 0x79, 0x73, 0x5f, 0x6f, 0x66, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x64, 0x61, 0x79, 0x73, 0x4f, 0x66, 0x4d, 0x6f, 0x6e,
	0x74, 0x68, 0x2a, 0x75, 0x0a, 0x11, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x1e, 0x53, 0x43, 0x48, 0x45, 0x44,
	0x55, 0x4c, 0x45, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x53,
	0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x4d, 0x55, 0x54, 0x45, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x53, 0x43, 0x48, 0x45,
	0x44, 0x55, 0x4c, 0x45, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41,
	0x43, 0x54, 0x49, 0x56, 0x41, 0x54, 0x45, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_rawDescData = file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_rawDesc
)

func file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_rawDescData)
	})
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_rawDescData
}

var file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_goTypes = []any{
	(ScheduleOperation)(0),    // 0: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.ScheduleOperation
	(*Schedule)(nil),          // 1: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Schedule
	(*OneTime)(nil),           // 2: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.OneTime
	(*Recurring)(nil),         // 3: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Recurring
	(*Daily)(nil),             // 4: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Daily
	(*Weekly)(nil),            // 5: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Weekly
	(*Monthly)(nil),           // 6: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Monthly
	(*Recurring_Always)(nil),  // 7: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Recurring.Always
	(*Recurring_Dynamic)(nil), // 8: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Recurring.Dynamic
	(*Timeframe)(nil),         // 9: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Timeframe
}
var file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_depIdxs = []int32{
	0,  // 0: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Schedule.schedule_operation:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.ScheduleOperation
	2,  // 1: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Schedule.one_time:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.OneTime
	3,  // 2: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Schedule.recurring:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Recurring
	9,  // 3: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.OneTime.timeframe:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Timeframe
	7,  // 4: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Recurring.always:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Recurring.Always
	8,  // 5: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Recurring.dynamic:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Recurring.Dynamic
	4,  // 6: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Recurring.Dynamic.daily:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Daily
	5,  // 7: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Recurring.Dynamic.weekly:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Weekly
	6,  // 8: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Recurring.Dynamic.monthly:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Monthly
	9,  // 9: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Recurring.Dynamic.timeframe:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Timeframe
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_init() }
func file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_init() {
	if File_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto != nil {
		return
	}
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_timeframe_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Schedule); i {
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
		file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*OneTime); i {
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
		file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Recurring); i {
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
		file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Daily); i {
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
		file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Weekly); i {
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
		file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Monthly); i {
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
		file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*Recurring_Always); i {
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
		file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*Recurring_Dynamic); i {
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
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_msgTypes[0].OneofWrappers = []any{
		(*Schedule_OneTime)(nil),
		(*Schedule_Recurring)(nil),
	}
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_msgTypes[2].OneofWrappers = []any{
		(*Recurring_Always_)(nil),
		(*Recurring_Dynamic_)(nil),
	}
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_msgTypes[7].OneofWrappers = []any{
		(*Recurring_Dynamic_Daily)(nil),
		(*Recurring_Dynamic_Weekly)(nil),
		(*Recurring_Dynamic_Monthly)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto = out.File
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_rawDesc = nil
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_goTypes = nil
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_depIdxs = nil
}
