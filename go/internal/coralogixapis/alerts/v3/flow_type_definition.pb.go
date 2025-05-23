// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/flow/flow_type_definition.proto

package v3

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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
	state              protoimpl.MessageState `protogen:"open.v1"`
	Stages             []*FlowStages          `protobuf:"bytes,1,rep,name=stages,proto3" json:"stages,omitempty"`
	EnforceSuppression *wrapperspb.BoolValue  `protobuf:"bytes,2,opt,name=enforce_suppression,json=enforceSuppression,proto3" json:"enforce_suppression,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to FlowStages:
	//
	//	*FlowStages_FlowStagesGroups
	FlowStages    isFlowStages_FlowStages `protobuf_oneof:"flow_stages"`
	TimeframeMs   *wrapperspb.Int64Value  `protobuf:"bytes,3,opt,name=timeframe_ms,json=timeframeMs,proto3" json:"timeframe_ms,omitempty"`
	TimeframeType TimeframeType           `protobuf:"varint,4,opt,name=timeframe_type,json=timeframeType,proto3,enum=com.coralogixapis.alerts.v3.TimeframeType" json:"timeframe_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

func (x *FlowStages) GetFlowStages() isFlowStages_FlowStages {
	if x != nil {
		return x.FlowStages
	}
	return nil
}

func (x *FlowStages) GetFlowStagesGroups() *FlowStagesGroups {
	if x != nil {
		if x, ok := x.FlowStages.(*FlowStages_FlowStagesGroups); ok {
			return x.FlowStagesGroups
		}
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Groups        []*FlowStagesGroup     `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state         protoimpl.MessageState       `protogen:"open.v1"`
	AlertDefs     []*FlowStagesGroupsAlertDefs `protobuf:"bytes,1,rep,name=alert_defs,json=alertDefs,proto3" json:"alert_defs,omitempty"`
	NextOp        NextOp                       `protobuf:"varint,2,opt,name=next_op,json=nextOp,proto3,enum=com.coralogixapis.alerts.v3.NextOp" json:"next_op,omitempty"`
	AlertsOp      AlertsOp                     `protobuf:"varint,3,opt,name=alerts_op,json=alertsOp,proto3,enum=com.coralogixapis.alerts.v3.AlertsOp" json:"alerts_op,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Id            *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Not           *wrapperspb.BoolValue   `protobuf:"bytes,2,opt,name=not,proto3" json:"not,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

const file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_rawDesc = "" +
	"\n" +
	"Ucom/coralogixapis/alerts/v3/alert_def_type_definition/flow/flow_type_definition.proto\x12\x1bcom.coralogixapis.alerts.v3\x1a\x1egoogle/protobuf/wrappers.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\xa0\x03\n" +
	"\bFlowType\x12f\n" +
	"\x06stages\x18\x01 \x03(\v2'.com.coralogixapis.alerts.v3.FlowStagesB%\x92A\"2\x1dThe stages of the flow alert.\xa8\x01\x01R\x06stages\x12K\n" +
	"\x13enforce_suppression\x18\x02 \x01(\v2\x1a.google.protobuf.BoolValueR\x12enforceSuppression:\xde\x01\x92A\xda\x01\n" +
	"T*\x0fFlow alert type28Configuration for flow-based alerts with multiple stages\xd2\x01\x06stages*\x81\x01\n" +
	"1Learn more about flow alerts in our documentation\x12Lhttps://coralogix.com/docs/user-guides/alerting/create-an-alert/flow-alerts/\"\xd6\x03\n" +
	"\n" +
	"FlowStages\x12{\n" +
	"\x12flow_stages_groups\x18\n" +
	" \x01(\v2-.com.coralogixapis.alerts.v3.FlowStagesGroupsB\x1c\x92A\x192\x17The flow stages groups.H\x00R\x10flowStagesGroups\x12>\n" +
	"\ftimeframe_ms\x18\x03 \x01(\v2\x1b.google.protobuf.Int64ValueR\vtimeframeMs\x12\x99\x01\n" +
	"\x0etimeframe_type\x18\x04 \x01(\x0e2*.com.coralogixapis.alerts.v3.TimeframeTypeBF\x92AC2)The type of timeframe for the flow alert.J\x16\"TIMEFRAME_TYPE_UP_TO\"R\rtimeframeType:`\x92A]\n" +
	"[*\vFlow stages2\x1eDefines stages in a flow alert\xd2\x01\vflow_stages\xd2\x01\ftimeframe_ms\xd2\x01\x0etimeframe_typeB\r\n" +
	"\vflow_stages\"\x9d\x01\n" +
	"\x10FlowStagesGroups\x12D\n" +
	"\x06groups\x18\x01 \x03(\v2,.com.coralogixapis.alerts.v3.FlowStagesGroupR\x06groups:C\x92A@\n" +
	">*\x11Flow stage groups2 Groups of stages in a flow alert\xd2\x01\x06groups\"\xa4\x04\n" +
	"\x0fFlowStagesGroup\x12\x8e\x01\n" +
	"\n" +
	"alert_defs\x18\x01 \x03(\v26.com.coralogixapis.alerts.v3.FlowStagesGroupsAlertDefsB7\x92A42/The alert definitions for the flow stage group.\xa8\x01\x01R\talertDefs\x12\x82\x01\n" +
	"\anext_op\x18\x02 \x01(\x0e2#.com.coralogixapis.alerts.v3.NextOpBD\x92AA21The logical operation to apply to the next stage.J\f\"NEXT_OP_OR\"R\x06nextOp\x12\x93\x01\n" +
	"\talerts_op\x18\x03 \x01(\x0e2%.com.coralogixapis.alerts.v3.AlertsOpBO\x92AL2:The logical operation to apply to the alerts in the group.J\x0e\"ALERTS_OP_OR\"R\balertsOp:e\x92Ab\n" +
	"`*\x10Flow stage group2)Defines a group of stages in a flow alert\xd2\x01\n" +
	"alert_defs\xd2\x01\anext_op\xd2\x01\talerts_op\"\xdb\x02\n" +
	"\x19FlowStagesGroupsAlertDefs\x12{\n" +
	"\x02id\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueBM\x92AJ2\x17The alert definition IDJ&\"123e4567-e89b-12d3-a456-426614174000\"\xa2\x02\x06UUIDv4R\x02id\x12g\n" +
	"\x03not\x18\x02 \x01(\v2\x1a.google.protobuf.BoolValueB9\x92A62.Whether to negate the alert definition or not.J\x04trueR\x03not:X\x92AU\n" +
	"S*\"Flow stage group alert definitions2(Alert definitions for a flow stage group\xd2\x01\x02id*8\n" +
	"\x06NextOp\x12\x1e\n" +
	"\x1aNEXT_OP_AND_OR_UNSPECIFIED\x10\x00\x12\x0e\n" +
	"\n" +
	"NEXT_OP_OR\x10\x01*>\n" +
	"\bAlertsOp\x12 \n" +
	"\x1cALERTS_OP_AND_OR_UNSPECIFIED\x10\x00\x12\x10\n" +
	"\fALERTS_OP_OR\x10\x01*I\n" +
	"\rTimeframeType\x12\x1e\n" +
	"\x1aTIMEFRAME_TYPE_UNSPECIFIED\x10\x00\x12\x18\n" +
	"\x14TIMEFRAME_TYPE_UP_TO\x10\x01b\x06proto3"

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_rawDesc)))
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_rawDesc)),
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
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_depIdxs = nil
}
