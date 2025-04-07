// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.26.1
// source: com/coralogixapis/incidents/v1/incident_query_filter.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

const file_com_coralogixapis_incidents_v1_incident_query_filter_proto_rawDesc = "" +
	"\n" +
	":com/coralogixapis/incidents/v1/incident_query_filter.proto\x12\x1ecom.coralogixapis.incidents.v1\x1a5com/coralogixapis/incidents/v1/filter_operators.proto\x1a6com/coralogixapis/incidents/v1/incident_severity.proto\x1a3com/coralogixapis/incidents/v1/incident_query.proto\x1a4com/coralogixapis/incidents/v1/incident_status.proto\x1a3com/coralogixapis/incidents/v1/incident_state.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/wrappers.proto\x1a/com/coralogixapis/incidents/v1/meta_label.proto\"\x8f\n" +
	"\n" +
	"\x13IncidentQueryFilter\x128\n" +
	"\bassignee\x18\x01 \x03(\v2\x1c.google.protobuf.StringValueR\bassignee\x12F\n" +
	"\x06status\x18\x02 \x03(\x0e2..com.coralogixapis.incidents.v1.IncidentStatusR\x06status\x12C\n" +
	"\x05state\x18\x03 \x03(\x0e2-.com.coralogixapis.incidents.v1.IncidentStateR\x05state\x12L\n" +
	"\bseverity\x18\x04 \x03(\x0e20.com.coralogixapis.incidents.v1.IncidentSeverityR\bseverity\x12v\n" +
	"\x11contextual_labels\x18\x05 \x03(\v2I.com.coralogixapis.incidents.v1.IncidentQueryFilter.ContextualLabelsEntryR\x10contextualLabels\x12=\n" +
	"\n" +
	"start_time\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampB\x02\x18\x01R\tstartTime\x129\n" +
	"\bend_time\x18\a \x01(\v2\x1a.google.protobuf.TimestampB\x02\x18\x01R\aendTime\x12V\n" +
	"\fsearch_query\x18\b \x01(\v23.com.coralogixapis.incidents.v1.IncidentSearchQueryR\vsearchQuery\x12G\n" +
	"\x10application_name\x18\t \x03(\v2\x1c.google.protobuf.StringValueR\x0fapplicationName\x12C\n" +
	"\x0esubsystem_name\x18\n" +
	" \x03(\v2\x1c.google.protobuf.StringValueR\rsubsystemName\x125\n" +
	"\bis_muted\x18\v \x01(\v2\x1a.google.protobuf.BoolValueR\aisMuted\x12S\n" +
	"\x10created_at_range\x18\f \x01(\v2).com.coralogixapis.incidents.v1.TimeRangeR\x0ecreatedAtRange\x12a\n" +
	"\x17incident_duration_range\x18\r \x01(\v2).com.coralogixapis.incidents.v1.TimeRangeR\x15incidentDurationRange\x12J\n" +
	"\vmeta_labels\x18\x0e \x03(\v2).com.coralogixapis.incidents.v1.MetaLabelR\n" +
	"metaLabels\x12T\n" +
	"\x0emeta_labels_op\x18\x0f \x01(\x0e2..com.coralogixapis.incidents.v1.FilterOperatorR\fmetaLabelsOp\x1az\n" +
	"\x15ContextualLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12K\n" +
	"\x05value\x18\x02 \x01(\v25.com.coralogixapis.incidents.v1.ContextualLabelValuesR\x05value:\x028\x01\"}\n" +
	"\tTimeRange\x129\n" +
	"\n" +
	"start_time\x18\x01 \x01(\v2\x1a.google.protobuf.TimestampR\tstartTime\x125\n" +
	"\bend_time\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\aendTime\"m\n" +
	"\x15ContextualLabelValues\x12T\n" +
	"\x17contextual_label_values\x18\x01 \x03(\v2\x1c.google.protobuf.StringValueR\x15contextualLabelValuesb\x06proto3"

var (
	file_com_coralogixapis_incidents_v1_incident_query_filter_proto_rawDescOnce sync.Once
	file_com_coralogixapis_incidents_v1_incident_query_filter_proto_rawDescData []byte
)

func file_com_coralogixapis_incidents_v1_incident_query_filter_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_incidents_v1_incident_query_filter_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_incidents_v1_incident_query_filter_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_incidents_v1_incident_query_filter_proto_rawDesc), len(file_com_coralogixapis_incidents_v1_incident_query_filter_proto_rawDesc)))
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_incidents_v1_incident_query_filter_proto_rawDesc), len(file_com_coralogixapis_incidents_v1_incident_query_filter_proto_rawDesc)),
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
	file_com_coralogixapis_incidents_v1_incident_query_filter_proto_goTypes = nil
	file_com_coralogixapis_incidents_v1_incident_query_filter_proto_depIdxs = nil
}
