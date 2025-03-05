// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: com/coralogixapis/incidents/v1/incident_action/incident_action.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AlertIncidentActionType int32

const (
	AlertIncidentActionType_ALERT_INCIDENT_ACTION_TYPE_UNSPECIFIED AlertIncidentActionType = 0
	AlertIncidentActionType_ALERT_INCIDENT_ACTION_TYPE_CLOSE       AlertIncidentActionType = 1
	AlertIncidentActionType_ALERT_INCIDENT_ACTION_TYPE_ACKNOWLEDGE AlertIncidentActionType = 2
	AlertIncidentActionType_ALERT_INCIDENT_ACTION_TYPE_ASSIGN      AlertIncidentActionType = 3
)

// Enum value maps for AlertIncidentActionType.
var (
	AlertIncidentActionType_name = map[int32]string{
		0: "ALERT_INCIDENT_ACTION_TYPE_UNSPECIFIED",
		1: "ALERT_INCIDENT_ACTION_TYPE_CLOSE",
		2: "ALERT_INCIDENT_ACTION_TYPE_ACKNOWLEDGE",
		3: "ALERT_INCIDENT_ACTION_TYPE_ASSIGN",
	}
	AlertIncidentActionType_value = map[string]int32{
		"ALERT_INCIDENT_ACTION_TYPE_UNSPECIFIED": 0,
		"ALERT_INCIDENT_ACTION_TYPE_CLOSE":       1,
		"ALERT_INCIDENT_ACTION_TYPE_ACKNOWLEDGE": 2,
		"ALERT_INCIDENT_ACTION_TYPE_ASSIGN":      3,
	}
)

func (x AlertIncidentActionType) Enum() *AlertIncidentActionType {
	p := new(AlertIncidentActionType)
	*p = x
	return p
}

func (x AlertIncidentActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AlertIncidentActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_enumTypes[0].Descriptor()
}

func (AlertIncidentActionType) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_enumTypes[0]
}

func (x AlertIncidentActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AlertIncidentActionType.Descriptor instead.
func (AlertIncidentActionType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_rawDescGZIP(), []int{0}
}

type IncidentActionType int32

const (
	IncidentActionType_INCIDENT_ACTION_TYPE_UNSPECIFIED  IncidentActionType = 0
	IncidentActionType_INCIDENT_ACTION_TYPE_CLOSE        IncidentActionType = 1
	IncidentActionType_INCIDENT_ACTION_TYPE_ACKNOWLEDGE  IncidentActionType = 2
	IncidentActionType_INCIDENT_ACTION_TYPE_ASSIGN       IncidentActionType = 3
	IncidentActionType_INCIDENT_ACTION_TYPE_UPSERT_STATE IncidentActionType = 4
)

// Enum value maps for IncidentActionType.
var (
	IncidentActionType_name = map[int32]string{
		0: "INCIDENT_ACTION_TYPE_UNSPECIFIED",
		1: "INCIDENT_ACTION_TYPE_CLOSE",
		2: "INCIDENT_ACTION_TYPE_ACKNOWLEDGE",
		3: "INCIDENT_ACTION_TYPE_ASSIGN",
		4: "INCIDENT_ACTION_TYPE_UPSERT_STATE",
	}
	IncidentActionType_value = map[string]int32{
		"INCIDENT_ACTION_TYPE_UNSPECIFIED":  0,
		"INCIDENT_ACTION_TYPE_CLOSE":        1,
		"INCIDENT_ACTION_TYPE_ACKNOWLEDGE":  2,
		"INCIDENT_ACTION_TYPE_ASSIGN":       3,
		"INCIDENT_ACTION_TYPE_UPSERT_STATE": 4,
	}
)

func (x IncidentActionType) Enum() *IncidentActionType {
	p := new(IncidentActionType)
	*p = x
	return p
}

func (x IncidentActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IncidentActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_enumTypes[1].Descriptor()
}

func (IncidentActionType) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_enumTypes[1]
}

func (x IncidentActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IncidentActionType.Descriptor instead.
func (IncidentActionType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_rawDescGZIP(), []int{1}
}

type IncidentAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActionType IncidentActionType `protobuf:"varint,1,opt,name=action_type,json=actionType,proto3,enum=com.coralogixapis.incidents.v1.IncidentActionType" json:"action_type,omitempty"`
	// Types that are assignable to Action:
	//
	//	*IncidentAction_UpsertState
	Action          isIncidentAction_Action `protobuf_oneof:"action"`
	ActionTimestamp *timestamppb.Timestamp  `protobuf:"bytes,2,opt,name=action_timestamp,json=actionTimestamp,proto3" json:"action_timestamp,omitempty"`
}

func (x *IncidentAction) Reset() {
	*x = IncidentAction{}
	mi := &file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IncidentAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncidentAction) ProtoMessage() {}

func (x *IncidentAction) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncidentAction.ProtoReflect.Descriptor instead.
func (*IncidentAction) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_rawDescGZIP(), []int{0}
}

func (x *IncidentAction) GetActionType() IncidentActionType {
	if x != nil {
		return x.ActionType
	}
	return IncidentActionType_INCIDENT_ACTION_TYPE_UNSPECIFIED
}

func (m *IncidentAction) GetAction() isIncidentAction_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *IncidentAction) GetUpsertState() *UpsertIncidentState {
	if x, ok := x.GetAction().(*IncidentAction_UpsertState); ok {
		return x.UpsertState
	}
	return nil
}

func (x *IncidentAction) GetActionTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.ActionTimestamp
	}
	return nil
}

type isIncidentAction_Action interface {
	isIncidentAction_Action()
}

type IncidentAction_UpsertState struct {
	UpsertState *UpsertIncidentState `protobuf:"bytes,104,opt,name=upsert_state,json=upsertState,proto3,oneof"`
}

func (*IncidentAction_UpsertState) isIncidentAction_Action() {}

var File_com_coralogixapis_incidents_v1_incident_action_incident_action_proto protoreflect.FileDescriptor

var file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_rawDesc = []byte{
	0x0a, 0x44, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x60, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x75, 0x70, 0x73, 0x65, 0x72, 0x74, 0x5f, 0x69, 0x6e,
	0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x75, 0x70, 0x73,
	0x65, 0x72, 0x74, 0x5f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x02, 0x0a, 0x0e, 0x49, 0x6e,
	0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x53, 0x0a, 0x0b,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x58, 0x0a, 0x0c, 0x75, 0x70, 0x73, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x68, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x49,
	0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0b,
	0x75, 0x70, 0x73, 0x65, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x45, 0x0a, 0x10, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0xbe, 0x01, 0x0a,
	0x17, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x26, 0x41, 0x4c, 0x45, 0x52,
	0x54, 0x5f, 0x49, 0x4e, 0x43, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x24, 0x0a, 0x20, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x5f, 0x49, 0x4e,
	0x43, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x10, 0x01, 0x12, 0x2a, 0x0a, 0x26, 0x41, 0x4c,
	0x45, 0x52, 0x54, 0x5f, 0x49, 0x4e, 0x43, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x43, 0x4b, 0x4e, 0x4f, 0x57, 0x4c,
	0x45, 0x44, 0x47, 0x45, 0x10, 0x02, 0x12, 0x25, 0x0a, 0x21, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x5f,
	0x49, 0x4e, 0x43, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x53, 0x53, 0x49, 0x47, 0x4e, 0x10, 0x03, 0x2a, 0xc8, 0x01,
	0x0a, 0x12, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x20, 0x49, 0x4e, 0x43, 0x49, 0x44, 0x45, 0x4e, 0x54,
	0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x49, 0x4e,
	0x43, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x10, 0x01, 0x12, 0x24, 0x0a, 0x20, 0x49, 0x4e,
	0x43, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x41, 0x43, 0x4b, 0x4e, 0x4f, 0x57, 0x4c, 0x45, 0x44, 0x47, 0x45, 0x10, 0x02,
	0x12, 0x1f, 0x0a, 0x1b, 0x49, 0x4e, 0x43, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x53, 0x53, 0x49, 0x47, 0x4e, 0x10,
	0x03, 0x12, 0x25, 0x0a, 0x21, 0x49, 0x4e, 0x43, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x50, 0x53, 0x45, 0x52, 0x54,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x10, 0x04, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_rawDescOnce sync.Once
	file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_rawDescData = file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_rawDesc
)

func file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_rawDescData)
	})
	return file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_rawDescData
}

var file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_goTypes = []any{
	(AlertIncidentActionType)(0),  // 0: com.coralogixapis.incidents.v1.AlertIncidentActionType
	(IncidentActionType)(0),       // 1: com.coralogixapis.incidents.v1.IncidentActionType
	(*IncidentAction)(nil),        // 2: com.coralogixapis.incidents.v1.IncidentAction
	(*UpsertIncidentState)(nil),   // 3: com.coralogixapis.incidents.v1.UpsertIncidentState
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.incidents.v1.IncidentAction.action_type:type_name -> com.coralogixapis.incidents.v1.IncidentActionType
	3, // 1: com.coralogixapis.incidents.v1.IncidentAction.upsert_state:type_name -> com.coralogixapis.incidents.v1.UpsertIncidentState
	4, // 2: com.coralogixapis.incidents.v1.IncidentAction.action_timestamp:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_init() }
func file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_init() {
	if File_com_coralogixapis_incidents_v1_incident_action_incident_action_proto != nil {
		return
	}
	file_com_coralogixapis_incidents_v1_incident_action_upsert_incident_state_upsert_incident_state_proto_init()
	file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_msgTypes[0].OneofWrappers = []any{
		(*IncidentAction_UpsertState)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_incidents_v1_incident_action_incident_action_proto = out.File
	file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_rawDesc = nil
	file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_goTypes = nil
	file_com_coralogixapis_incidents_v1_incident_action_incident_action_proto_depIdxs = nil
}
