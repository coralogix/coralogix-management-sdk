// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: com/coralogixapis/incidents/v1/incident_event_extended/incident_event_extended.proto

package v1

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

type IncidentEventAlertType int32

const (
	IncidentEventAlertType_INCIDENT_EVENT_ALERT_TYPE_STANDARD_OR_UNSPECIFIED IncidentEventAlertType = 0
	IncidentEventAlertType_INCIDENT_EVENT_ALERT_TYPE_METRIC                  IncidentEventAlertType = 1
	IncidentEventAlertType_INCIDENT_EVENT_ALERT_TYPE_NEW_VALUE               IncidentEventAlertType = 2
	IncidentEventAlertType_INCIDENT_EVENT_ALERT_TYPE_RATIO                   IncidentEventAlertType = 3
	IncidentEventAlertType_INCIDENT_EVENT_ALERT_TYPE_TIME_RELATIVE           IncidentEventAlertType = 4
	IncidentEventAlertType_INCIDENT_EVENT_ALERT_TYPE_UNIQUE_COUNT            IncidentEventAlertType = 5
	IncidentEventAlertType_INCIDENT_EVENT_ALERT_TYPE_TRACING                 IncidentEventAlertType = 6
	IncidentEventAlertType_INCIDENT_EVENT_ALERT_TYPE_FLOW                    IncidentEventAlertType = 7
)

// Enum value maps for IncidentEventAlertType.
var (
	IncidentEventAlertType_name = map[int32]string{
		0: "INCIDENT_EVENT_ALERT_TYPE_STANDARD_OR_UNSPECIFIED",
		1: "INCIDENT_EVENT_ALERT_TYPE_METRIC",
		2: "INCIDENT_EVENT_ALERT_TYPE_NEW_VALUE",
		3: "INCIDENT_EVENT_ALERT_TYPE_RATIO",
		4: "INCIDENT_EVENT_ALERT_TYPE_TIME_RELATIVE",
		5: "INCIDENT_EVENT_ALERT_TYPE_UNIQUE_COUNT",
		6: "INCIDENT_EVENT_ALERT_TYPE_TRACING",
		7: "INCIDENT_EVENT_ALERT_TYPE_FLOW",
	}
	IncidentEventAlertType_value = map[string]int32{
		"INCIDENT_EVENT_ALERT_TYPE_STANDARD_OR_UNSPECIFIED": 0,
		"INCIDENT_EVENT_ALERT_TYPE_METRIC":                  1,
		"INCIDENT_EVENT_ALERT_TYPE_NEW_VALUE":               2,
		"INCIDENT_EVENT_ALERT_TYPE_RATIO":                   3,
		"INCIDENT_EVENT_ALERT_TYPE_TIME_RELATIVE":           4,
		"INCIDENT_EVENT_ALERT_TYPE_UNIQUE_COUNT":            5,
		"INCIDENT_EVENT_ALERT_TYPE_TRACING":                 6,
		"INCIDENT_EVENT_ALERT_TYPE_FLOW":                    7,
	}
)

func (x IncidentEventAlertType) Enum() *IncidentEventAlertType {
	p := new(IncidentEventAlertType)
	*p = x
	return p
}

func (x IncidentEventAlertType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IncidentEventAlertType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_enumTypes[0].Descriptor()
}

func (IncidentEventAlertType) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_enumTypes[0]
}

func (x IncidentEventAlertType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IncidentEventAlertType.Descriptor instead.
func (IncidentEventAlertType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_rawDescGZIP(), []int{0}
}

type IncidentEventExtended struct {
	state                         protoimpl.MessageState         `protogen:"open.v1"`
	CxEventKey                    *wrapperspb.StringValue        `protobuf:"bytes,1,opt,name=cx_event_key,json=cxEventKey,proto3" json:"cx_event_key,omitempty"`
	IncidentEvent                 *IncidentEvent                 `protobuf:"bytes,2,opt,name=incident_event,json=incidentEvent,proto3" json:"incident_event,omitempty"`
	CxEventTimestamp              *timestamppb.Timestamp         `protobuf:"bytes,4,opt,name=cx_event_timestamp,json=cxEventTimestamp,proto3" json:"cx_event_timestamp,omitempty"`
	IncidentEventExtendedMetadata *IncidentEventExtendedMetadata `protobuf:"bytes,5,opt,name=incident_event_extended_metadata,json=incidentEventExtendedMetadata,proto3" json:"incident_event_extended_metadata,omitempty"`
	unknownFields                 protoimpl.UnknownFields
	sizeCache                     protoimpl.SizeCache
}

func (x *IncidentEventExtended) Reset() {
	*x = IncidentEventExtended{}
	mi := &file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IncidentEventExtended) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncidentEventExtended) ProtoMessage() {}

func (x *IncidentEventExtended) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncidentEventExtended.ProtoReflect.Descriptor instead.
func (*IncidentEventExtended) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_rawDescGZIP(), []int{0}
}

func (x *IncidentEventExtended) GetCxEventKey() *wrapperspb.StringValue {
	if x != nil {
		return x.CxEventKey
	}
	return nil
}

func (x *IncidentEventExtended) GetIncidentEvent() *IncidentEvent {
	if x != nil {
		return x.IncidentEvent
	}
	return nil
}

func (x *IncidentEventExtended) GetCxEventTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.CxEventTimestamp
	}
	return nil
}

func (x *IncidentEventExtended) GetIncidentEventExtendedMetadata() *IncidentEventExtendedMetadata {
	if x != nil {
		return x.IncidentEventExtendedMetadata
	}
	return nil
}

type IncidentEventExtendedMetadata struct {
	state               protoimpl.MessageState    `protogen:"open.v1"`
	IncidentSeverity    IncidentSeverity          `protobuf:"varint,1,opt,name=incident_severity,json=incidentSeverity,proto3,enum=com.coralogixapis.incidents.v1.IncidentSeverity" json:"incident_severity,omitempty"`
	AlertId             *wrapperspb.StringValue   `protobuf:"bytes,2,opt,name=alert_id,json=alertId,proto3" json:"alert_id,omitempty"`
	IncidentState       IncidentState             `protobuf:"varint,3,opt,name=incident_state,json=incidentState,proto3,enum=com.coralogixapis.incidents.v1.IncidentState" json:"incident_state,omitempty"`
	AlertName           *wrapperspb.StringValue   `protobuf:"bytes,4,opt,name=alert_name,json=alertName,proto3" json:"alert_name,omitempty"`
	AlertType           IncidentEventAlertType    `protobuf:"varint,5,opt,name=alert_type,json=alertType,proto3,enum=com.coralogixapis.incidents.v1.IncidentEventAlertType" json:"alert_type,omitempty"`
	IsMuted             *wrapperspb.BoolValue     `protobuf:"bytes,6,opt,name=is_muted,json=isMuted,proto3" json:"is_muted,omitempty"`
	IncidentStatus      IncidentStatus            `protobuf:"varint,7,opt,name=incident_status,json=incidentStatus,proto3,enum=com.coralogixapis.incidents.v1.IncidentStatus" json:"incident_status,omitempty"`
	AlertGroupByFields  []*wrapperspb.StringValue `protobuf:"bytes,8,rep,name=alert_group_by_fields,json=alertGroupByFields,proto3" json:"alert_group_by_fields,omitempty"`
	AlertLabels         []*MetaLabel              `protobuf:"bytes,9,rep,name=alert_labels,json=alertLabels,proto3" json:"alert_labels,omitempty"`
	IncidentPermutation map[string]string         `protobuf:"bytes,10,rep,name=incident_permutation,json=incidentPermutation,proto3" json:"incident_permutation,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *IncidentEventExtendedMetadata) Reset() {
	*x = IncidentEventExtendedMetadata{}
	mi := &file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IncidentEventExtendedMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncidentEventExtendedMetadata) ProtoMessage() {}

func (x *IncidentEventExtendedMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncidentEventExtendedMetadata.ProtoReflect.Descriptor instead.
func (*IncidentEventExtendedMetadata) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_rawDescGZIP(), []int{1}
}

func (x *IncidentEventExtendedMetadata) GetIncidentSeverity() IncidentSeverity {
	if x != nil {
		return x.IncidentSeverity
	}
	return IncidentSeverity_INCIDENT_SEVERITY_UNSPECIFIED
}

func (x *IncidentEventExtendedMetadata) GetAlertId() *wrapperspb.StringValue {
	if x != nil {
		return x.AlertId
	}
	return nil
}

func (x *IncidentEventExtendedMetadata) GetIncidentState() IncidentState {
	if x != nil {
		return x.IncidentState
	}
	return IncidentState_INCIDENT_STATE_UNSPECIFIED
}

func (x *IncidentEventExtendedMetadata) GetAlertName() *wrapperspb.StringValue {
	if x != nil {
		return x.AlertName
	}
	return nil
}

func (x *IncidentEventExtendedMetadata) GetAlertType() IncidentEventAlertType {
	if x != nil {
		return x.AlertType
	}
	return IncidentEventAlertType_INCIDENT_EVENT_ALERT_TYPE_STANDARD_OR_UNSPECIFIED
}

func (x *IncidentEventExtendedMetadata) GetIsMuted() *wrapperspb.BoolValue {
	if x != nil {
		return x.IsMuted
	}
	return nil
}

func (x *IncidentEventExtendedMetadata) GetIncidentStatus() IncidentStatus {
	if x != nil {
		return x.IncidentStatus
	}
	return IncidentStatus_INCIDENT_STATUS_UNSPECIFIED
}

func (x *IncidentEventExtendedMetadata) GetAlertGroupByFields() []*wrapperspb.StringValue {
	if x != nil {
		return x.AlertGroupByFields
	}
	return nil
}

func (x *IncidentEventExtendedMetadata) GetAlertLabels() []*MetaLabel {
	if x != nil {
		return x.AlertLabels
	}
	return nil
}

func (x *IncidentEventExtendedMetadata) GetIncidentPermutation() map[string]string {
	if x != nil {
		return x.IncidentPermutation
	}
	return nil
}

var File_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto protoreflect.FileDescriptor

var file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_rawDesc = []byte{
	0x0a, 0x54, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x42, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e,
	0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x33, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x80, 0x03, 0x0a, 0x15, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x3e, 0x0a, 0x0c, 0x63, 0x78, 0x5f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x63,
	0x78, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x54, 0x0a, 0x0e, 0x69, 0x6e, 0x63,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x0d, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x48, 0x0a, 0x12, 0x63, 0x78, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x63, 0x78, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x86, 0x01, 0x0a, 0x20, 0x69, 0x6e,
	0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x1d, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x22, 0xa4, 0x07, 0x0a, 0x1d, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x5d, 0x0a, 0x11, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74,
	0x79, 0x52, 0x10, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x76, 0x65, 0x72,
	0x69, 0x74, 0x79, 0x12, 0x37, 0x0a, 0x08, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x07, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x49, 0x64, 0x12, 0x54, 0x0a, 0x0e,
	0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x0d, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x55, 0x0a, 0x0a, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x6d, 0x75, 0x74,
	0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x69, 0x73, 0x4d, 0x75, 0x74, 0x65, 0x64, 0x12, 0x57, 0x0a,
	0x0f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4f, 0x0a, 0x15, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x62, 0x79, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x12, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42,
	0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x4c, 0x0a, 0x0c, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x0b, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x89, 0x01, 0x0a, 0x14, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x56, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x6d,
	0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x13, 0x69, 0x6e,
	0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x46, 0x0a, 0x18, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x72,
	0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0xe7, 0x02, 0x0a, 0x16, 0x49, 0x6e,
	0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x6c, 0x65, 0x72, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x35, 0x0a, 0x31, 0x49, 0x4e, 0x43, 0x49, 0x44, 0x45, 0x4e, 0x54,
	0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x41, 0x52, 0x44, 0x5f, 0x4f, 0x52, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x24, 0x0a, 0x20, 0x49,
	0x4e, 0x43, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x4c,
	0x45, 0x52, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x10,
	0x01, 0x12, 0x27, 0x0a, 0x23, 0x49, 0x4e, 0x43, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x5f, 0x45, 0x56,
	0x45, 0x4e, 0x54, 0x5f, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e,
	0x45, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x02, 0x12, 0x23, 0x0a, 0x1f, 0x49, 0x4e,
	0x43, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x4c, 0x45,
	0x52, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x10, 0x03, 0x12,
	0x2b, 0x0a, 0x27, 0x49, 0x4e, 0x43, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x49, 0x4d,
	0x45, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x56, 0x45, 0x10, 0x04, 0x12, 0x2a, 0x0a, 0x26,
	0x49, 0x4e, 0x43, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x41,
	0x4c, 0x45, 0x52, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x49, 0x51, 0x55, 0x45,
	0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x05, 0x12, 0x25, 0x0a, 0x21, 0x49, 0x4e, 0x43, 0x49,
	0x44, 0x45, 0x4e, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x4c, 0x45, 0x52, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x49, 0x4e, 0x47, 0x10, 0x06, 0x12,
	0x22, 0x0a, 0x1e, 0x49, 0x4e, 0x43, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x4c, 0x4f,
	0x57, 0x10, 0x07, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_rawDescOnce sync.Once
	file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_rawDescData = file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_rawDesc
)

func file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_rawDescData)
	})
	return file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_rawDescData
}

var file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_goTypes = []any{
	(IncidentEventAlertType)(0),           // 0: com.coralogixapis.incidents.v1.IncidentEventAlertType
	(*IncidentEventExtended)(nil),         // 1: com.coralogixapis.incidents.v1.IncidentEventExtended
	(*IncidentEventExtendedMetadata)(nil), // 2: com.coralogixapis.incidents.v1.IncidentEventExtendedMetadata
	nil,                                   // 3: com.coralogixapis.incidents.v1.IncidentEventExtendedMetadata.IncidentPermutationEntry
	(*wrapperspb.StringValue)(nil),        // 4: google.protobuf.StringValue
	(*IncidentEvent)(nil),                 // 5: com.coralogixapis.incidents.v1.IncidentEvent
	(*timestamppb.Timestamp)(nil),         // 6: google.protobuf.Timestamp
	(IncidentSeverity)(0),                 // 7: com.coralogixapis.incidents.v1.IncidentSeverity
	(IncidentState)(0),                    // 8: com.coralogixapis.incidents.v1.IncidentState
	(*wrapperspb.BoolValue)(nil),          // 9: google.protobuf.BoolValue
	(IncidentStatus)(0),                   // 10: com.coralogixapis.incidents.v1.IncidentStatus
	(*MetaLabel)(nil),                     // 11: com.coralogixapis.incidents.v1.MetaLabel
}
var file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_depIdxs = []int32{
	4,  // 0: com.coralogixapis.incidents.v1.IncidentEventExtended.cx_event_key:type_name -> google.protobuf.StringValue
	5,  // 1: com.coralogixapis.incidents.v1.IncidentEventExtended.incident_event:type_name -> com.coralogixapis.incidents.v1.IncidentEvent
	6,  // 2: com.coralogixapis.incidents.v1.IncidentEventExtended.cx_event_timestamp:type_name -> google.protobuf.Timestamp
	2,  // 3: com.coralogixapis.incidents.v1.IncidentEventExtended.incident_event_extended_metadata:type_name -> com.coralogixapis.incidents.v1.IncidentEventExtendedMetadata
	7,  // 4: com.coralogixapis.incidents.v1.IncidentEventExtendedMetadata.incident_severity:type_name -> com.coralogixapis.incidents.v1.IncidentSeverity
	4,  // 5: com.coralogixapis.incidents.v1.IncidentEventExtendedMetadata.alert_id:type_name -> google.protobuf.StringValue
	8,  // 6: com.coralogixapis.incidents.v1.IncidentEventExtendedMetadata.incident_state:type_name -> com.coralogixapis.incidents.v1.IncidentState
	4,  // 7: com.coralogixapis.incidents.v1.IncidentEventExtendedMetadata.alert_name:type_name -> google.protobuf.StringValue
	0,  // 8: com.coralogixapis.incidents.v1.IncidentEventExtendedMetadata.alert_type:type_name -> com.coralogixapis.incidents.v1.IncidentEventAlertType
	9,  // 9: com.coralogixapis.incidents.v1.IncidentEventExtendedMetadata.is_muted:type_name -> google.protobuf.BoolValue
	10, // 10: com.coralogixapis.incidents.v1.IncidentEventExtendedMetadata.incident_status:type_name -> com.coralogixapis.incidents.v1.IncidentStatus
	4,  // 11: com.coralogixapis.incidents.v1.IncidentEventExtendedMetadata.alert_group_by_fields:type_name -> google.protobuf.StringValue
	11, // 12: com.coralogixapis.incidents.v1.IncidentEventExtendedMetadata.alert_labels:type_name -> com.coralogixapis.incidents.v1.MetaLabel
	3,  // 13: com.coralogixapis.incidents.v1.IncidentEventExtendedMetadata.incident_permutation:type_name -> com.coralogixapis.incidents.v1.IncidentEventExtendedMetadata.IncidentPermutationEntry
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_init()
}
func file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_init() {
	if File_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto != nil {
		return
	}
	file_com_coralogixapis_incidents_v1_incident_event_incident_event_proto_init()
	file_com_coralogixapis_incidents_v1_incident_severity_proto_init()
	file_com_coralogixapis_incidents_v1_incident_state_proto_init()
	file_com_coralogixapis_incidents_v1_incident_status_proto_init()
	file_com_coralogixapis_incidents_v1_meta_label_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto = out.File
	file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_rawDesc = nil
	file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_goTypes = nil
	file_com_coralogixapis_incidents_v1_incident_event_extended_incident_event_extended_proto_depIdxs = nil
}
