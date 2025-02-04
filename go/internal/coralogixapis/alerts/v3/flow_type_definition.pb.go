// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/flow/flow_type_definition.proto

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

type NextOp int32

const (
	NextOp_NEXT_OP_AND_OR_UNSPECIFIED NextOp = 0
	NextOp_NEXT_OP_OR                 NextOp = 1
)

// Enum value maps for NextOp.
var (
	NextOp_name = map[int32]string{
		0: "NEXT_OP_AND_OR_UNSPECIFIED",
		1: "NEXT_OP_OR",
	}
	NextOp_value = map[string]int32{
		"NEXT_OP_AND_OR_UNSPECIFIED": 0,
		"NEXT_OP_OR":                 1,
	}
)

func (x NextOp) Enum() *NextOp {
	p := new(NextOp)
	*p = x
	return p
}

func (x NextOp) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NextOp) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_enumTypes[0].Descriptor()
}

func (NextOp) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_enumTypes[0]
}

func (x NextOp) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NextOp.Descriptor instead.
func (NextOp) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_rawDescGZIP(), []int{0}
}

type AlertsOp int32

const (
	AlertsOp_ALERTS_OP_AND_OR_UNSPECIFIED AlertsOp = 0
	AlertsOp_ALERTS_OP_OR                 AlertsOp = 1
)

// Enum value maps for AlertsOp.
var (
	AlertsOp_name = map[int32]string{
		0: "ALERTS_OP_AND_OR_UNSPECIFIED",
		1: "ALERTS_OP_OR",
	}
	AlertsOp_value = map[string]int32{
		"ALERTS_OP_AND_OR_UNSPECIFIED": 0,
		"ALERTS_OP_OR":                 1,
	}
)

func (x AlertsOp) Enum() *AlertsOp {
	p := new(AlertsOp)
	*p = x
	return p
}

func (x AlertsOp) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AlertsOp) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_enumTypes[1].Descriptor()
}

func (AlertsOp) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_enumTypes[1]
}

func (x AlertsOp) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AlertsOp.Descriptor instead.
func (AlertsOp) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_rawDescGZIP(), []int{1}
}

type TimeframeType int32

const (
	TimeframeType_TIMEFRAME_TYPE_UNSPECIFIED TimeframeType = 0
	TimeframeType_TIMEFRAME_TYPE_UP_TO       TimeframeType = 1
)

// Enum value maps for TimeframeType.
var (
	TimeframeType_name = map[int32]string{
		0: "TIMEFRAME_TYPE_UNSPECIFIED",
		1: "TIMEFRAME_TYPE_UP_TO",
	}
	TimeframeType_value = map[string]int32{
		"TIMEFRAME_TYPE_UNSPECIFIED": 0,
		"TIMEFRAME_TYPE_UP_TO":       1,
	}
)

func (x TimeframeType) Enum() *TimeframeType {
	p := new(TimeframeType)
	*p = x
	return p
}

func (x TimeframeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TimeframeType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_enumTypes[2].Descriptor()
}

func (TimeframeType) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_enumTypes[2]
}

func (x TimeframeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TimeframeType.Descriptor instead.
func (TimeframeType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_rawDescGZIP(), []int{2}
}

type FlowType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stages             []*FlowStages         `protobuf:"bytes,1,rep,name=stages,proto3" json:"stages,omitempty"`
	EnforceSuppression *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=enforce_suppression,json=enforceSuppression,proto3" json:"enforce_suppression,omitempty"`
}

func (x *FlowType) Reset() {
	*x = FlowType{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FlowType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowType) ProtoMessage() {}

func (x *FlowType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowType.ProtoReflect.Descriptor instead.
func (*FlowType) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_rawDescGZIP(), []int{0}
}

func (x *FlowType) GetStages() []*FlowStages {
	if x != nil {
		return x.Stages
	}
	return nil
}

func (x *FlowType) GetEnforceSuppression() *wrapperspb.BoolValue {
	if x != nil {
		return x.EnforceSuppression
	}
	return nil
}

type FlowStages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to FlowStages:
	//
	//	*FlowStages_FlowStagesGroups
	FlowStages    isFlowStages_FlowStages `protobuf_oneof:"flow_stages"`
	TimeframeMs   *wrapperspb.Int64Value  `protobuf:"bytes,3,opt,name=timeframe_ms,json=timeframeMs,proto3" json:"timeframe_ms,omitempty"`
	TimeframeType TimeframeType           `protobuf:"varint,4,opt,name=timeframe_type,json=timeframeType,proto3,enum=com.coralogixapis.alerts.v3.TimeframeType" json:"timeframe_type,omitempty"`
}

func (x *FlowStages) Reset() {
	*x = FlowStages{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FlowStages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowStages) ProtoMessage() {}

func (x *FlowStages) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowStages.ProtoReflect.Descriptor instead.
func (*FlowStages) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_rawDescGZIP(), []int{1}
}

func (m *FlowStages) GetFlowStages() isFlowStages_FlowStages {
	if m != nil {
		return m.FlowStages
	}
	return nil
}

func (x *FlowStages) GetFlowStagesGroups() *FlowStagesGroups {
	if x, ok := x.GetFlowStages().(*FlowStages_FlowStagesGroups); ok {
		return x.FlowStagesGroups
	}
	return nil
}

func (x *FlowStages) GetTimeframeMs() *wrapperspb.Int64Value {
	if x != nil {
		return x.TimeframeMs
	}
	return nil
}

func (x *FlowStages) GetTimeframeType() TimeframeType {
	if x != nil {
		return x.TimeframeType
	}
	return TimeframeType_TIMEFRAME_TYPE_UNSPECIFIED
}

type isFlowStages_FlowStages interface {
	isFlowStages_FlowStages()
}

type FlowStages_FlowStagesGroups struct {
	FlowStagesGroups *FlowStagesGroups `protobuf:"bytes,10,opt,name=flow_stages_groups,json=flowStagesGroups,proto3,oneof"`
}

func (*FlowStages_FlowStagesGroups) isFlowStages_FlowStages() {}

type FlowStagesGroups struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groups []*FlowStagesGroup `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *FlowStagesGroups) Reset() {
	*x = FlowStagesGroups{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FlowStagesGroups) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowStagesGroups) ProtoMessage() {}

func (x *FlowStagesGroups) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowStagesGroups.ProtoReflect.Descriptor instead.
func (*FlowStagesGroups) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_rawDescGZIP(), []int{2}
}

func (x *FlowStagesGroups) GetGroups() []*FlowStagesGroup {
	if x != nil {
		return x.Groups
	}
	return nil
}

type FlowStagesGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlertDefs []*FlowStagesGroupsAlertDefs `protobuf:"bytes,1,rep,name=alert_defs,json=alertDefs,proto3" json:"alert_defs,omitempty"`
	NextOp    NextOp                       `protobuf:"varint,2,opt,name=next_op,json=nextOp,proto3,enum=com.coralogixapis.alerts.v3.NextOp" json:"next_op,omitempty"`
	AlertsOp  AlertsOp                     `protobuf:"varint,3,opt,name=alerts_op,json=alertsOp,proto3,enum=com.coralogixapis.alerts.v3.AlertsOp" json:"alerts_op,omitempty"`
}

func (x *FlowStagesGroup) Reset() {
	*x = FlowStagesGroup{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FlowStagesGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowStagesGroup) ProtoMessage() {}

func (x *FlowStagesGroup) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowStagesGroup.ProtoReflect.Descriptor instead.
func (*FlowStagesGroup) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_rawDescGZIP(), []int{3}
}

func (x *FlowStagesGroup) GetAlertDefs() []*FlowStagesGroupsAlertDefs {
	if x != nil {
		return x.AlertDefs
	}
	return nil
}

func (x *FlowStagesGroup) GetNextOp() NextOp {
	if x != nil {
		return x.NextOp
	}
	return NextOp_NEXT_OP_AND_OR_UNSPECIFIED
}

func (x *FlowStagesGroup) GetAlertsOp() AlertsOp {
	if x != nil {
		return x.AlertsOp
	}
	return AlertsOp_ALERTS_OP_AND_OR_UNSPECIFIED
}

type FlowStagesGroupsAlertDefs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Not *wrapperspb.BoolValue   `protobuf:"bytes,2,opt,name=not,proto3" json:"not,omitempty"`
}

func (x *FlowStagesGroupsAlertDefs) Reset() {
	*x = FlowStagesGroupsAlertDefs{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FlowStagesGroupsAlertDefs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowStagesGroupsAlertDefs) ProtoMessage() {}

func (x *FlowStagesGroupsAlertDefs) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowStagesGroupsAlertDefs.ProtoReflect.Descriptor instead.
func (*FlowStagesGroupsAlertDefs) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_rawDescGZIP(), []int{4}
}

func (x *FlowStagesGroupsAlertDefs) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *FlowStagesGroupsAlertDefs) GetNot() *wrapperspb.BoolValue {
	if x != nil {
		return x.Not
	}
	return nil
}

var File_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_rawDesc = []byte{
	0x0a, 0x55, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x66, 0x6c, 0x6f,
	0x77, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x73, 0x2e, 0x76, 0x33, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x08, 0x46, 0x6c, 0x6f, 0x77, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x3f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e,
	0x46, 0x6c, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x67, 0x65, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x4b, 0x0a, 0x13, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x75,
	0x70, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x12, 0x65, 0x6e, 0x66,
	0x6f, 0x72, 0x63, 0x65, 0x53, 0x75, 0x70, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0x8d, 0x02, 0x0a, 0x0a, 0x46, 0x6c, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x67, 0x65, 0x73, 0x12, 0x5d,
	0x0a, 0x12, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x53, 0x74, 0x61,
	0x67, 0x65, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x48, 0x00, 0x52, 0x10, 0x66, 0x6c, 0x6f,
	0x77, 0x53, 0x74, 0x61, 0x67, 0x65, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x3e, 0x0a,
	0x0c, 0x74, 0x69, 0x6d, 0x65, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x4d, 0x73, 0x12, 0x51, 0x0a,
	0x0e, 0x74, 0x69, 0x6d, 0x65, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x2e, 0x76, 0x33, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x42, 0x0d, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x73, 0x22,
	0x58, 0x0a, 0x10, 0x46, 0x6c, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x67, 0x65, 0x73, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x12, 0x44, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76,
	0x33, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x67, 0x65, 0x73, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0xea, 0x01, 0x0a, 0x0f, 0x46, 0x6c,
	0x6f, 0x77, 0x53, 0x74, 0x61, 0x67, 0x65, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x55, 0x0a,
	0x0a, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e,
	0x46, 0x6c, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x67, 0x65, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x44, 0x65, 0x66, 0x73, 0x52, 0x09, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x44, 0x65, 0x66, 0x73, 0x12, 0x3c, 0x0a, 0x07, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x6f, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x2e, 0x76, 0x33, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x4f, 0x70, 0x52, 0x06, 0x6e, 0x65, 0x78, 0x74,
	0x4f, 0x70, 0x12, 0x42, 0x0a, 0x09, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x5f, 0x6f, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x2e, 0x76, 0x33, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x4f, 0x70, 0x52, 0x08, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x4f, 0x70, 0x22, 0x77, 0x0a, 0x19, 0x46, 0x6c, 0x6f, 0x77, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x44,
	0x65, 0x66, 0x73, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x2c, 0x0a, 0x03, 0x6e, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x6e, 0x6f, 0x74, 0x2a,
	0x38, 0x0a, 0x06, 0x4e, 0x65, 0x78, 0x74, 0x4f, 0x70, 0x12, 0x1e, 0x0a, 0x1a, 0x4e, 0x45, 0x58,
	0x54, 0x5f, 0x4f, 0x50, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x45, 0x58,
	0x54, 0x5f, 0x4f, 0x50, 0x5f, 0x4f, 0x52, 0x10, 0x01, 0x2a, 0x3e, 0x0a, 0x08, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x73, 0x4f, 0x70, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x53, 0x5f,
	0x4f, 0x50, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x4c, 0x45, 0x52, 0x54,
	0x53, 0x5f, 0x4f, 0x50, 0x5f, 0x4f, 0x52, 0x10, 0x01, 0x2a, 0x49, 0x0a, 0x0d, 0x54, 0x69, 0x6d,
	0x65, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x54, 0x49,
	0x4d, 0x45, 0x46, 0x52, 0x41, 0x4d, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x49,
	0x4d, 0x45, 0x46, 0x52, 0x41, 0x4d, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x50, 0x5f,
	0x54, 0x4f, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_rawDescData = file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_goTypes = []any{
	(NextOp)(0),                       // 0: com.coralogixapis.alerts.v3.NextOp
	(AlertsOp)(0),                     // 1: com.coralogixapis.alerts.v3.AlertsOp
	(TimeframeType)(0),                // 2: com.coralogixapis.alerts.v3.TimeframeType
	(*FlowType)(nil),                  // 3: com.coralogixapis.alerts.v3.FlowType
	(*FlowStages)(nil),                // 4: com.coralogixapis.alerts.v3.FlowStages
	(*FlowStagesGroups)(nil),          // 5: com.coralogixapis.alerts.v3.FlowStagesGroups
	(*FlowStagesGroup)(nil),           // 6: com.coralogixapis.alerts.v3.FlowStagesGroup
	(*FlowStagesGroupsAlertDefs)(nil), // 7: com.coralogixapis.alerts.v3.FlowStagesGroupsAlertDefs
	(*wrapperspb.BoolValue)(nil),      // 8: google.protobuf.BoolValue
	(*wrapperspb.Int64Value)(nil),     // 9: google.protobuf.Int64Value
	(*wrapperspb.StringValue)(nil),    // 10: google.protobuf.StringValue
}
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_depIdxs = []int32{
	4,  // 0: com.coralogixapis.alerts.v3.FlowType.stages:type_name -> com.coralogixapis.alerts.v3.FlowStages
	8,  // 1: com.coralogixapis.alerts.v3.FlowType.enforce_suppression:type_name -> google.protobuf.BoolValue
	5,  // 2: com.coralogixapis.alerts.v3.FlowStages.flow_stages_groups:type_name -> com.coralogixapis.alerts.v3.FlowStagesGroups
	9,  // 3: com.coralogixapis.alerts.v3.FlowStages.timeframe_ms:type_name -> google.protobuf.Int64Value
	2,  // 4: com.coralogixapis.alerts.v3.FlowStages.timeframe_type:type_name -> com.coralogixapis.alerts.v3.TimeframeType
	6,  // 5: com.coralogixapis.alerts.v3.FlowStagesGroups.groups:type_name -> com.coralogixapis.alerts.v3.FlowStagesGroup
	7,  // 6: com.coralogixapis.alerts.v3.FlowStagesGroup.alert_defs:type_name -> com.coralogixapis.alerts.v3.FlowStagesGroupsAlertDefs
	0,  // 7: com.coralogixapis.alerts.v3.FlowStagesGroup.next_op:type_name -> com.coralogixapis.alerts.v3.NextOp
	1,  // 8: com.coralogixapis.alerts.v3.FlowStagesGroup.alerts_op:type_name -> com.coralogixapis.alerts.v3.AlertsOp
	10, // 9: com.coralogixapis.alerts.v3.FlowStagesGroupsAlertDefs.id:type_name -> google.protobuf.StringValue
	8,  // 10: com.coralogixapis.alerts.v3.FlowStagesGroupsAlertDefs.not:type_name -> google.protobuf.BoolValue
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_init()
}
func file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_msgTypes[1].OneofWrappers = []any{
		(*FlowStages_FlowStagesGroups)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_depIdxs = nil
}
