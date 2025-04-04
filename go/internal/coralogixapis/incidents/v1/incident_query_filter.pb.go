// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: com/coralogixapis/incidents/v1/incident_query_filter.proto

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

type IncidentQueryFilter struct {
	state            protoimpl.MessageState            `protogen:"open.v1"`
	Assignee         []*wrapperspb.StringValue         `protobuf:"bytes,1,rep,name=assignee,proto3" json:"assignee,omitempty"`
	Status           []IncidentStatus                  `protobuf:"varint,2,rep,packed,name=status,proto3,enum=com.coralogixapis.incidents.v1.IncidentStatus" json:"status,omitempty"`
	State            []IncidentState                   `protobuf:"varint,3,rep,packed,name=state,proto3,enum=com.coralogixapis.incidents.v1.IncidentState" json:"state,omitempty"`
	Severity         []IncidentSeverity                `protobuf:"varint,4,rep,packed,name=severity,proto3,enum=com.coralogixapis.incidents.v1.IncidentSeverity" json:"severity,omitempty"`
	ContextualLabels map[string]*ContextualLabelValues `protobuf:"bytes,5,rep,name=contextual_labels,json=contextualLabels,proto3" json:"contextual_labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Deprecated: Marked as deprecated in com/coralogixapis/incidents/v1/incident_query_filter.proto.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"` //filters all incidents that were open in the given timeframe start time (deprecated, use incident_open_range instead)
	// Deprecated: Marked as deprecated in com/coralogixapis/incidents/v1/incident_query_filter.proto.
	EndTime               *timestamppb.Timestamp    `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"` //filters all incidents that were open in the given timeframe end time (deprecated, use incident_open_range instead)
	SearchQuery           *IncidentSearchQuery      `protobuf:"bytes,8,opt,name=search_query,json=searchQuery,proto3" json:"search_query,omitempty"`
	ApplicationName       []*wrapperspb.StringValue `protobuf:"bytes,9,rep,name=application_name,json=applicationName,proto3" json:"application_name,omitempty"`
	SubsystemName         []*wrapperspb.StringValue `protobuf:"bytes,10,rep,name=subsystem_name,json=subsystemName,proto3" json:"subsystem_name,omitempty"`
	IsMuted               *wrapperspb.BoolValue     `protobuf:"bytes,11,opt,name=is_muted,json=isMuted,proto3" json:"is_muted,omitempty"`
	CreatedAtRange        *TimeRange                `protobuf:"bytes,12,opt,name=created_at_range,json=createdAtRange,proto3" json:"created_at_range,omitempty"`                                               // filters all incidents created at the given time range
	IncidentDurationRange *TimeRange                `protobuf:"bytes,13,opt,name=incident_duration_range,json=incidentDurationRange,proto3" json:"incident_duration_range,omitempty"`                          // filters all incidents open (alive) at the given time range
	MetaLabels            []*MetaLabel              `protobuf:"bytes,14,rep,name=meta_labels,json=metaLabels,proto3" json:"meta_labels,omitempty"`                                                             // filters all incidents with the given meta labels
	MetaLabelsOp          FilterOperator            `protobuf:"varint,15,opt,name=meta_labels_op,json=metaLabelsOp,proto3,enum=com.coralogixapis.incidents.v1.FilterOperator" json:"meta_labels_op,omitempty"` // filter operations for meta labels filter
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *IncidentQueryFilter) Reset() {
	*x = IncidentQueryFilter{}
	mi := &file_com_coralogixapis_incidents_v1_incident_query_filter_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IncidentQueryFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncidentQueryFilter) ProtoMessage() {}

func (x *IncidentQueryFilter) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_incidents_v1_incident_query_filter_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncidentQueryFilter.ProtoReflect.Descriptor instead.
func (*IncidentQueryFilter) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_incidents_v1_incident_query_filter_proto_rawDescGZIP(), []int{0}
}

func (x *IncidentQueryFilter) GetAssignee() []*wrapperspb.StringValue {
	if x != nil {
		return x.Assignee
	}
	return nil
}

func (x *IncidentQueryFilter) GetStatus() []IncidentStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *IncidentQueryFilter) GetState() []IncidentState {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *IncidentQueryFilter) GetSeverity() []IncidentSeverity {
	if x != nil {
		return x.Severity
	}
	return nil
}

func (x *IncidentQueryFilter) GetContextualLabels() map[string]*ContextualLabelValues {
	if x != nil {
		return x.ContextualLabels
	}
	return nil
}

// Deprecated: Marked as deprecated in com/coralogixapis/incidents/v1/incident_query_filter.proto.
func (x *IncidentQueryFilter) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

// Deprecated: Marked as deprecated in com/coralogixapis/incidents/v1/incident_query_filter.proto.
func (x *IncidentQueryFilter) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *IncidentQueryFilter) GetSearchQuery() *IncidentSearchQuery {
	if x != nil {
		return x.SearchQuery
	}
	return nil
}

func (x *IncidentQueryFilter) GetApplicationName() []*wrapperspb.StringValue {
	if x != nil {
		return x.ApplicationName
	}
	return nil
}

func (x *IncidentQueryFilter) GetSubsystemName() []*wrapperspb.StringValue {
	if x != nil {
		return x.SubsystemName
	}
	return nil
}

func (x *IncidentQueryFilter) GetIsMuted() *wrapperspb.BoolValue {
	if x != nil {
		return x.IsMuted
	}
	return nil
}

func (x *IncidentQueryFilter) GetCreatedAtRange() *TimeRange {
	if x != nil {
		return x.CreatedAtRange
	}
	return nil
}

func (x *IncidentQueryFilter) GetIncidentDurationRange() *TimeRange {
	if x != nil {
		return x.IncidentDurationRange
	}
	return nil
}

func (x *IncidentQueryFilter) GetMetaLabels() []*MetaLabel {
	if x != nil {
		return x.MetaLabels
	}
	return nil
}

func (x *IncidentQueryFilter) GetMetaLabelsOp() FilterOperator {
	if x != nil {
		return x.MetaLabelsOp
	}
	return FilterOperator_FILTER_OPERATOR_OR_OR_UNSPECIFIED
}

type TimeRange struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StartTime     *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime       *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TimeRange) Reset() {
	*x = TimeRange{}
	mi := &file_com_coralogixapis_incidents_v1_incident_query_filter_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimeRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeRange) ProtoMessage() {}

func (x *TimeRange) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_incidents_v1_incident_query_filter_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeRange.ProtoReflect.Descriptor instead.
func (*TimeRange) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_incidents_v1_incident_query_filter_proto_rawDescGZIP(), []int{1}
}

func (x *TimeRange) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *TimeRange) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

type ContextualLabelValues struct {
	state                 protoimpl.MessageState    `protogen:"open.v1"`
	ContextualLabelValues []*wrapperspb.StringValue `protobuf:"bytes,1,rep,name=contextual_label_values,json=contextualLabelValues,proto3" json:"contextual_label_values,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *ContextualLabelValues) Reset() {
	*x = ContextualLabelValues{}
	mi := &file_com_coralogixapis_incidents_v1_incident_query_filter_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContextualLabelValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextualLabelValues) ProtoMessage() {}

func (x *ContextualLabelValues) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_incidents_v1_incident_query_filter_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextualLabelValues.ProtoReflect.Descriptor instead.
func (*ContextualLabelValues) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_incidents_v1_incident_query_filter_proto_rawDescGZIP(), []int{2}
}

func (x *ContextualLabelValues) GetContextualLabelValues() []*wrapperspb.StringValue {
	if x != nil {
		return x.ContextualLabelValues
	}
	return nil
}

var File_com_coralogixapis_incidents_v1_incident_query_filter_proto protoreflect.FileDescriptor

var file_com_coralogixapis_incidents_v1_incident_query_filter_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x35, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x76,
	0x65, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69,
	0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x63, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x34, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x0a,
	0x0a, 0x13, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x12,
	0x46, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x43, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x4c, 0x0a, 0x08,
	0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x30,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79,
	0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x76, 0x0a, 0x11, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x49, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x75, 0x61, 0x6c, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x75, 0x61, 0x6c, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x12, 0x3d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x02, 0x18, 0x01, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x39, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x02, 0x18, 0x01, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x56, 0x0a, 0x0c,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x47, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x43, 0x0a,
	0x0e, 0x73, 0x75, 0x62, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x6d, 0x75, 0x74, 0x65, 0x64, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x07, 0x69, 0x73, 0x4d, 0x75, 0x74, 0x65, 0x64, 0x12, 0x53, 0x0a, 0x10, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0e,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x61,
	0x0a, 0x17, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x15, 0x69, 0x6e, 0x63, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x54, 0x0a,
	0x0e, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x5f, 0x6f, 0x70, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x0c, 0x6d, 0x65, 0x74, 0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x4f, 0x70, 0x1a, 0x7a, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x75, 0x61,
	0x6c, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x4b,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x75, 0x61, 0x6c, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x7d, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x6d,
	0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x75, 0x61, 0x6c, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x54, 0x0a, 0x17, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x15, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x75,
	0x61, 0x6c, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_incidents_v1_incident_query_filter_proto_rawDescOnce sync.Once
	file_com_coralogixapis_incidents_v1_incident_query_filter_proto_rawDescData = file_com_coralogixapis_incidents_v1_incident_query_filter_proto_rawDesc
)

func file_com_coralogixapis_incidents_v1_incident_query_filter_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_incidents_v1_incident_query_filter_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_incidents_v1_incident_query_filter_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_incidents_v1_incident_query_filter_proto_rawDescData)
	})
	return file_com_coralogixapis_incidents_v1_incident_query_filter_proto_rawDescData
}

var file_com_coralogixapis_incidents_v1_incident_query_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_com_coralogixapis_incidents_v1_incident_query_filter_proto_goTypes = []any{
	(*IncidentQueryFilter)(nil),    // 0: com.coralogixapis.incidents.v1.IncidentQueryFilter
	(*TimeRange)(nil),              // 1: com.coralogixapis.incidents.v1.TimeRange
	(*ContextualLabelValues)(nil),  // 2: com.coralogixapis.incidents.v1.ContextualLabelValues
	nil,                            // 3: com.coralogixapis.incidents.v1.IncidentQueryFilter.ContextualLabelsEntry
	(*wrapperspb.StringValue)(nil), // 4: google.protobuf.StringValue
	(IncidentStatus)(0),            // 5: com.coralogixapis.incidents.v1.IncidentStatus
	(IncidentState)(0),             // 6: com.coralogixapis.incidents.v1.IncidentState
	(IncidentSeverity)(0),          // 7: com.coralogixapis.incidents.v1.IncidentSeverity
	(*timestamppb.Timestamp)(nil),  // 8: google.protobuf.Timestamp
	(*IncidentSearchQuery)(nil),    // 9: com.coralogixapis.incidents.v1.IncidentSearchQuery
	(*wrapperspb.BoolValue)(nil),   // 10: google.protobuf.BoolValue
	(*MetaLabel)(nil),              // 11: com.coralogixapis.incidents.v1.MetaLabel
	(FilterOperator)(0),            // 12: com.coralogixapis.incidents.v1.FilterOperator
}
var file_com_coralogixapis_incidents_v1_incident_query_filter_proto_depIdxs = []int32{
	4,  // 0: com.coralogixapis.incidents.v1.IncidentQueryFilter.assignee:type_name -> google.protobuf.StringValue
	5,  // 1: com.coralogixapis.incidents.v1.IncidentQueryFilter.status:type_name -> com.coralogixapis.incidents.v1.IncidentStatus
	6,  // 2: com.coralogixapis.incidents.v1.IncidentQueryFilter.state:type_name -> com.coralogixapis.incidents.v1.IncidentState
	7,  // 3: com.coralogixapis.incidents.v1.IncidentQueryFilter.severity:type_name -> com.coralogixapis.incidents.v1.IncidentSeverity
	3,  // 4: com.coralogixapis.incidents.v1.IncidentQueryFilter.contextual_labels:type_name -> com.coralogixapis.incidents.v1.IncidentQueryFilter.ContextualLabelsEntry
	8,  // 5: com.coralogixapis.incidents.v1.IncidentQueryFilter.start_time:type_name -> google.protobuf.Timestamp
	8,  // 6: com.coralogixapis.incidents.v1.IncidentQueryFilter.end_time:type_name -> google.protobuf.Timestamp
	9,  // 7: com.coralogixapis.incidents.v1.IncidentQueryFilter.search_query:type_name -> com.coralogixapis.incidents.v1.IncidentSearchQuery
	4,  // 8: com.coralogixapis.incidents.v1.IncidentQueryFilter.application_name:type_name -> google.protobuf.StringValue
	4,  // 9: com.coralogixapis.incidents.v1.IncidentQueryFilter.subsystem_name:type_name -> google.protobuf.StringValue
	10, // 10: com.coralogixapis.incidents.v1.IncidentQueryFilter.is_muted:type_name -> google.protobuf.BoolValue
	1,  // 11: com.coralogixapis.incidents.v1.IncidentQueryFilter.created_at_range:type_name -> com.coralogixapis.incidents.v1.TimeRange
	1,  // 12: com.coralogixapis.incidents.v1.IncidentQueryFilter.incident_duration_range:type_name -> com.coralogixapis.incidents.v1.TimeRange
	11, // 13: com.coralogixapis.incidents.v1.IncidentQueryFilter.meta_labels:type_name -> com.coralogixapis.incidents.v1.MetaLabel
	12, // 14: com.coralogixapis.incidents.v1.IncidentQueryFilter.meta_labels_op:type_name -> com.coralogixapis.incidents.v1.FilterOperator
	8,  // 15: com.coralogixapis.incidents.v1.TimeRange.start_time:type_name -> google.protobuf.Timestamp
	8,  // 16: com.coralogixapis.incidents.v1.TimeRange.end_time:type_name -> google.protobuf.Timestamp
	4,  // 17: com.coralogixapis.incidents.v1.ContextualLabelValues.contextual_label_values:type_name -> google.protobuf.StringValue
	2,  // 18: com.coralogixapis.incidents.v1.IncidentQueryFilter.ContextualLabelsEntry.value:type_name -> com.coralogixapis.incidents.v1.ContextualLabelValues
	19, // [19:19] is the sub-list for method output_type
	19, // [19:19] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_incidents_v1_incident_query_filter_proto_init() }
func file_com_coralogixapis_incidents_v1_incident_query_filter_proto_init() {
	if File_com_coralogixapis_incidents_v1_incident_query_filter_proto != nil {
		return
	}
	file_com_coralogixapis_incidents_v1_filter_operators_proto_init()
	file_com_coralogixapis_incidents_v1_incident_severity_proto_init()
	file_com_coralogixapis_incidents_v1_incident_query_proto_init()
	file_com_coralogixapis_incidents_v1_incident_status_proto_init()
	file_com_coralogixapis_incidents_v1_incident_state_proto_init()
	file_com_coralogixapis_incidents_v1_meta_label_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_incidents_v1_incident_query_filter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_incidents_v1_incident_query_filter_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_incidents_v1_incident_query_filter_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_incidents_v1_incident_query_filter_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_incidents_v1_incident_query_filter_proto = out.File
	file_com_coralogixapis_incidents_v1_incident_query_filter_proto_rawDesc = nil
	file_com_coralogixapis_incidents_v1_incident_query_filter_proto_goTypes = nil
	file_com_coralogixapis_incidents_v1_incident_query_filter_proto_depIdxs = nil
}
