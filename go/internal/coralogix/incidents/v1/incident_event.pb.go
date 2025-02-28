// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.26.1
// source: com/coralogixapis/incidents/v1/incident_event/incident_event.proto

package v1

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

type IncidentEvent struct {
	state             protoimpl.MessageState  `protogen:"open.v1"`
	Id                *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	IncidentEventType IncidentEventType       `protobuf:"varint,2,opt,name=incident_event_type,json=incidentEventType,proto3,enum=com.coralogixapis.incidents.v1.IncidentEventType" json:"incident_event_type,omitempty"`
	// Types that are valid to be assigned to IncidentEventPayload:
	//
	//	*IncidentEvent_SnoozeIndicator
	//	*IncidentEvent_Assignment
	//	*IncidentEvent_Unassign
	//	*IncidentEvent_UpsertState
	//	*IncidentEvent_Acknowledge
	//	*IncidentEvent_Close
	IncidentEventPayload isIncidentEvent_IncidentEventPayload `protobuf_oneof:"incident_event_payload"`
	OriginatorType       OriginatorType                       `protobuf:"varint,3,opt,name=originator_type,json=originatorType,proto3,enum=com.coralogixapis.incidents.v1.OriginatorType" json:"originator_type,omitempty"`
	// Types that are valid to be assigned to Originator:
	//
	//	*IncidentEvent_AdministrativeEvent
	//	*IncidentEvent_OperationalEvent
	Originator    isIncidentEvent_Originator `protobuf_oneof:"originator"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IncidentEvent) Reset() {
	*x = IncidentEvent{}
	mi := &file_com_coralogixapis_incidents_v1_incident_event_incident_event_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IncidentEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncidentEvent) ProtoMessage() {}

func (x *IncidentEvent) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_incidents_v1_incident_event_incident_event_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncidentEvent.ProtoReflect.Descriptor instead.
func (*IncidentEvent) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_incidents_v1_incident_event_incident_event_proto_rawDescGZIP(), []int{0}
}

func (x *IncidentEvent) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *IncidentEvent) GetIncidentEventType() IncidentEventType {
	if x != nil {
		return x.IncidentEventType
	}
	return IncidentEventType_INCIDENT_EVENT_TYPE_UNSPECIFIED
}

func (x *IncidentEvent) GetIncidentEventPayload() isIncidentEvent_IncidentEventPayload {
	if x != nil {
		return x.IncidentEventPayload
	}
	return nil
}

func (x *IncidentEvent) GetSnoozeIndicator() *IncidentEventSnoozeIndicator {
	if x != nil {
		if x, ok := x.IncidentEventPayload.(*IncidentEvent_SnoozeIndicator); ok {
			return x.SnoozeIndicator
		}
	}
	return nil
}

func (x *IncidentEvent) GetAssignment() *IncidentEventAssign {
	if x != nil {
		if x, ok := x.IncidentEventPayload.(*IncidentEvent_Assignment); ok {
			return x.Assignment
		}
	}
	return nil
}

func (x *IncidentEvent) GetUnassign() *IncidentEventUnassign {
	if x != nil {
		if x, ok := x.IncidentEventPayload.(*IncidentEvent_Unassign); ok {
			return x.Unassign
		}
	}
	return nil
}

func (x *IncidentEvent) GetUpsertState() *IncidentEventUpsertState {
	if x != nil {
		if x, ok := x.IncidentEventPayload.(*IncidentEvent_UpsertState); ok {
			return x.UpsertState
		}
	}
	return nil
}

func (x *IncidentEvent) GetAcknowledge() *IncidentEventAcknowledge {
	if x != nil {
		if x, ok := x.IncidentEventPayload.(*IncidentEvent_Acknowledge); ok {
			return x.Acknowledge
		}
	}
	return nil
}

func (x *IncidentEvent) GetClose() *IncidentEventClose {
	if x != nil {
		if x, ok := x.IncidentEventPayload.(*IncidentEvent_Close); ok {
			return x.Close
		}
	}
	return nil
}

func (x *IncidentEvent) GetOriginatorType() OriginatorType {
	if x != nil {
		return x.OriginatorType
	}
	return OriginatorType_ORIGINATOR_TYPE_UNSPECIFIED
}

func (x *IncidentEvent) GetOriginator() isIncidentEvent_Originator {
	if x != nil {
		return x.Originator
	}
	return nil
}

func (x *IncidentEvent) GetAdministrativeEvent() *IncidentEventOriginatorAdministrative {
	if x != nil {
		if x, ok := x.Originator.(*IncidentEvent_AdministrativeEvent); ok {
			return x.AdministrativeEvent
		}
	}
	return nil
}

func (x *IncidentEvent) GetOperationalEvent() *IncidentEventOriginatorOperational {
	if x != nil {
		if x, ok := x.Originator.(*IncidentEvent_OperationalEvent); ok {
			return x.OperationalEvent
		}
	}
	return nil
}

type isIncidentEvent_IncidentEventPayload interface {
	isIncidentEvent_IncidentEventPayload()
}

type IncidentEvent_SnoozeIndicator struct {
	SnoozeIndicator *IncidentEventSnoozeIndicator `protobuf:"bytes,200,opt,name=snooze_indicator,json=snoozeIndicator,proto3,oneof"`
}

type IncidentEvent_Assignment struct {
	Assignment *IncidentEventAssign `protobuf:"bytes,201,opt,name=assignment,proto3,oneof"`
}

type IncidentEvent_Unassign struct {
	Unassign *IncidentEventUnassign `protobuf:"bytes,205,opt,name=unassign,proto3,oneof"`
}

type IncidentEvent_UpsertState struct {
	UpsertState *IncidentEventUpsertState `protobuf:"bytes,202,opt,name=upsert_state,json=upsertState,proto3,oneof"`
}

type IncidentEvent_Acknowledge struct {
	Acknowledge *IncidentEventAcknowledge `protobuf:"bytes,203,opt,name=acknowledge,proto3,oneof"`
}

type IncidentEvent_Close struct {
	Close *IncidentEventClose `protobuf:"bytes,204,opt,name=close,proto3,oneof"`
}

func (*IncidentEvent_SnoozeIndicator) isIncidentEvent_IncidentEventPayload() {}

func (*IncidentEvent_Assignment) isIncidentEvent_IncidentEventPayload() {}

func (*IncidentEvent_Unassign) isIncidentEvent_IncidentEventPayload() {}

func (*IncidentEvent_UpsertState) isIncidentEvent_IncidentEventPayload() {}

func (*IncidentEvent_Acknowledge) isIncidentEvent_IncidentEventPayload() {}

func (*IncidentEvent_Close) isIncidentEvent_IncidentEventPayload() {}

type isIncidentEvent_Originator interface {
	isIncidentEvent_Originator()
}

type IncidentEvent_AdministrativeEvent struct {
	AdministrativeEvent *IncidentEventOriginatorAdministrative `protobuf:"bytes,100,opt,name=administrative_event,json=administrativeEvent,proto3,oneof"`
}

type IncidentEvent_OperationalEvent struct {
	OperationalEvent *IncidentEventOriginatorOperational `protobuf:"bytes,101,opt,name=operational_event,json=operationalEvent,proto3,oneof"`
}

func (*IncidentEvent_AdministrativeEvent) isIncidentEvent_Originator() {}

func (*IncidentEvent_OperationalEvent) isIncidentEvent_Originator() {}

var File_com_coralogixapis_incidents_v1_incident_event_incident_event_proto protoreflect.FileDescriptor

var file_com_coralogixapis_incidents_v1_incident_event_incident_event_proto_rawDesc = []byte{
	0x0a, 0x42, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f,
	0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x1a, 0x4e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x61, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x49, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x48, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x69,
	0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c,
	0x6f, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x5c, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x63,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x6f, 0x72, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x59, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72,
	0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x52, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x53, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x69,
	0x63, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x47, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69,
	0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x63, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4f, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x75, 0x70, 0x73, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4b, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x75, 0x6e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb6, 0x08, 0x0a, 0x0d, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x61, 0x0a, 0x13, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x11, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x6a, 0x0a, 0x10, 0x73, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x5f,
	0x69, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53,
	0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x48, 0x00,
	0x52, 0x0f, 0x73, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f,
	0x72, 0x12, 0x56, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0xc9, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x48, 0x00, 0x52, 0x0a, 0x61,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x54, 0x0a, 0x08, 0x75, 0x6e, 0x61,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x18, 0xcd, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e,
	0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x55, 0x6e, 0x61, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x48, 0x00, 0x52, 0x08, 0x75, 0x6e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x12,
	0x5e, 0x0a, 0x0c, 0x75, 0x70, 0x73, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0xca, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x48, 0x00, 0x52, 0x0b, 0x75, 0x70, 0x73, 0x65, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x5d, 0x0a, 0x0b, 0x61, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x18, 0xcb,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x48,
	0x00, 0x52, 0x0b, 0x61, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x12, 0x4b,
	0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0xcc, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x48, 0x00, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x0f, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x7a, 0x0a, 0x14, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x64, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x45, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x48, 0x01, 0x52, 0x13, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x71, 0x0a, 0x11, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x63,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x6f, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x48,
	0x01, 0x52, 0x10, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x42, 0x18, 0x0a, 0x16, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x0c, 0x0a,
	0x0a, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_incidents_v1_incident_event_incident_event_proto_rawDescOnce sync.Once
	file_com_coralogixapis_incidents_v1_incident_event_incident_event_proto_rawDescData = file_com_coralogixapis_incidents_v1_incident_event_incident_event_proto_rawDesc
)

func file_com_coralogixapis_incidents_v1_incident_event_incident_event_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_incidents_v1_incident_event_incident_event_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_incidents_v1_incident_event_incident_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_incidents_v1_incident_event_incident_event_proto_rawDescData)
	})
	return file_com_coralogixapis_incidents_v1_incident_event_incident_event_proto_rawDescData
}

var file_com_coralogixapis_incidents_v1_incident_event_incident_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_incidents_v1_incident_event_incident_event_proto_goTypes = []any{
	(*IncidentEvent)(nil),                         // 0: com.coralogixapis.incidents.v1.IncidentEvent
	(*wrapperspb.StringValue)(nil),                // 1: google.protobuf.StringValue
	(IncidentEventType)(0),                        // 2: com.coralogixapis.incidents.v1.IncidentEventType
	(*IncidentEventSnoozeIndicator)(nil),          // 3: com.coralogixapis.incidents.v1.IncidentEventSnoozeIndicator
	(*IncidentEventAssign)(nil),                   // 4: com.coralogixapis.incidents.v1.IncidentEventAssign
	(*IncidentEventUnassign)(nil),                 // 5: com.coralogixapis.incidents.v1.IncidentEventUnassign
	(*IncidentEventUpsertState)(nil),              // 6: com.coralogixapis.incidents.v1.IncidentEventUpsertState
	(*IncidentEventAcknowledge)(nil),              // 7: com.coralogixapis.incidents.v1.IncidentEventAcknowledge
	(*IncidentEventClose)(nil),                    // 8: com.coralogixapis.incidents.v1.IncidentEventClose
	(OriginatorType)(0),                           // 9: com.coralogixapis.incidents.v1.OriginatorType
	(*IncidentEventOriginatorAdministrative)(nil), // 10: com.coralogixapis.incidents.v1.IncidentEventOriginatorAdministrative
	(*IncidentEventOriginatorOperational)(nil),    // 11: com.coralogixapis.incidents.v1.IncidentEventOriginatorOperational
}
var file_com_coralogixapis_incidents_v1_incident_event_incident_event_proto_depIdxs = []int32{
	1,  // 0: com.coralogixapis.incidents.v1.IncidentEvent.id:type_name -> google.protobuf.StringValue
	2,  // 1: com.coralogixapis.incidents.v1.IncidentEvent.incident_event_type:type_name -> com.coralogixapis.incidents.v1.IncidentEventType
	3,  // 2: com.coralogixapis.incidents.v1.IncidentEvent.snooze_indicator:type_name -> com.coralogixapis.incidents.v1.IncidentEventSnoozeIndicator
	4,  // 3: com.coralogixapis.incidents.v1.IncidentEvent.assignment:type_name -> com.coralogixapis.incidents.v1.IncidentEventAssign
	5,  // 4: com.coralogixapis.incidents.v1.IncidentEvent.unassign:type_name -> com.coralogixapis.incidents.v1.IncidentEventUnassign
	6,  // 5: com.coralogixapis.incidents.v1.IncidentEvent.upsert_state:type_name -> com.coralogixapis.incidents.v1.IncidentEventUpsertState
	7,  // 6: com.coralogixapis.incidents.v1.IncidentEvent.acknowledge:type_name -> com.coralogixapis.incidents.v1.IncidentEventAcknowledge
	8,  // 7: com.coralogixapis.incidents.v1.IncidentEvent.close:type_name -> com.coralogixapis.incidents.v1.IncidentEventClose
	9,  // 8: com.coralogixapis.incidents.v1.IncidentEvent.originator_type:type_name -> com.coralogixapis.incidents.v1.OriginatorType
	10, // 9: com.coralogixapis.incidents.v1.IncidentEvent.administrative_event:type_name -> com.coralogixapis.incidents.v1.IncidentEventOriginatorAdministrative
	11, // 10: com.coralogixapis.incidents.v1.IncidentEvent.operational_event:type_name -> com.coralogixapis.incidents.v1.IncidentEventOriginatorOperational
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_incidents_v1_incident_event_incident_event_proto_init() }
func file_com_coralogixapis_incidents_v1_incident_event_incident_event_proto_init() {
	if File_com_coralogixapis_incidents_v1_incident_event_incident_event_proto != nil {
		return
	}
	file_com_coralogixapis_incidents_v1_incident_event_incident_event_acknowledge_proto_init()
	file_com_coralogixapis_incidents_v1_incident_event_incident_event_assign_proto_init()
	file_com_coralogixapis_incidents_v1_incident_event_incident_event_close_proto_init()
	file_com_coralogixapis_incidents_v1_incident_event_incident_event_originator_administrative_proto_init()
	file_com_coralogixapis_incidents_v1_incident_event_incident_event_originator_operational_proto_init()
	file_com_coralogixapis_incidents_v1_incident_event_incident_event_originator_type_proto_init()
	file_com_coralogixapis_incidents_v1_incident_event_incident_event_snooze_indicator_proto_init()
	file_com_coralogixapis_incidents_v1_incident_event_incident_event_type_proto_init()
	file_com_coralogixapis_incidents_v1_incident_event_incident_event_upsert_state_proto_init()
	file_com_coralogixapis_incidents_v1_incident_event_incident_event_unassign_proto_init()
	file_com_coralogixapis_incidents_v1_incident_event_incident_event_proto_msgTypes[0].OneofWrappers = []any{
		(*IncidentEvent_SnoozeIndicator)(nil),
		(*IncidentEvent_Assignment)(nil),
		(*IncidentEvent_Unassign)(nil),
		(*IncidentEvent_UpsertState)(nil),
		(*IncidentEvent_Acknowledge)(nil),
		(*IncidentEvent_Close)(nil),
		(*IncidentEvent_AdministrativeEvent)(nil),
		(*IncidentEvent_OperationalEvent)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_incidents_v1_incident_event_incident_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_incidents_v1_incident_event_incident_event_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_incidents_v1_incident_event_incident_event_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_incidents_v1_incident_event_incident_event_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_incidents_v1_incident_event_incident_event_proto = out.File
	file_com_coralogixapis_incidents_v1_incident_event_incident_event_proto_rawDesc = nil
	file_com_coralogixapis_incidents_v1_incident_event_incident_event_proto_goTypes = nil
	file_com_coralogixapis_incidents_v1_incident_event_incident_event_proto_depIdxs = nil
}
