// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/incidents/v1/incident_event_query_filter.proto

package v1

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

type IncidentEventQueryFilter struct {
	state            protoimpl.MessageState            `protogen:"open.v1"`
	Status           []IncidentStatus                  `protobuf:"varint,1,rep,packed,name=status,proto3,enum=com.coralogixapis.incidents.v1.IncidentStatus" json:"status,omitempty"`
	Severity         []IncidentSeverity                `protobuf:"varint,2,rep,packed,name=severity,proto3,enum=com.coralogixapis.incidents.v1.IncidentSeverity" json:"severity,omitempty"`
	ContextualLabels map[string]*ContextualLabelValues `protobuf:"bytes,3,rep,name=contextual_labels,json=contextualLabels,proto3" json:"contextual_labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Name             *wrapperspb.StringValue           `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	IsMuted          *wrapperspb.BoolValue             `protobuf:"bytes,5,opt,name=is_muted,json=isMuted,proto3" json:"is_muted,omitempty"`
	Timestamp        *TimeRange                        `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Labels           *LabelsFilter                     `protobuf:"bytes,7,opt,name=labels,proto3" json:"labels,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *IncidentEventQueryFilter) Reset() {
	*x = IncidentEventQueryFilter{}
	mi := &file_com_coralogixapis_incidents_v1_incident_event_query_filter_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IncidentEventQueryFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncidentEventQueryFilter) ProtoMessage() {}

func (x *IncidentEventQueryFilter) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_incidents_v1_incident_event_query_filter_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncidentEventQueryFilter.ProtoReflect.Descriptor instead.
func (*IncidentEventQueryFilter) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_incidents_v1_incident_event_query_filter_proto_rawDescGZIP(), []int{0}
}

func (x *IncidentEventQueryFilter) GetStatus() []IncidentStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *IncidentEventQueryFilter) GetSeverity() []IncidentSeverity {
	if x != nil {
		return x.Severity
	}
	return nil
}

func (x *IncidentEventQueryFilter) GetContextualLabels() map[string]*ContextualLabelValues {
	if x != nil {
		return x.ContextualLabels
	}
	return nil
}

func (x *IncidentEventQueryFilter) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *IncidentEventQueryFilter) GetIsMuted() *wrapperspb.BoolValue {
	if x != nil {
		return x.IsMuted
	}
	return nil
}

func (x *IncidentEventQueryFilter) GetTimestamp() *TimeRange {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *IncidentEventQueryFilter) GetLabels() *LabelsFilter {
	if x != nil {
		return x.Labels
	}
	return nil
}

type LabelsFilter struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MetaLabels    []*MetaLabel           `protobuf:"bytes,1,rep,name=meta_labels,json=metaLabels,proto3" json:"meta_labels,omitempty"`
	Operator      FilterOperator         `protobuf:"varint,2,opt,name=operator,proto3,enum=com.coralogixapis.incidents.v1.FilterOperator" json:"operator,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LabelsFilter) Reset() {
	*x = LabelsFilter{}
	mi := &file_com_coralogixapis_incidents_v1_incident_event_query_filter_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LabelsFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelsFilter) ProtoMessage() {}

func (x *LabelsFilter) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_incidents_v1_incident_event_query_filter_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelsFilter.ProtoReflect.Descriptor instead.
func (*LabelsFilter) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_incidents_v1_incident_event_query_filter_proto_rawDescGZIP(), []int{1}
}

func (x *LabelsFilter) GetMetaLabels() []*MetaLabel {
	if x != nil {
		return x.MetaLabels
	}
	return nil
}

func (x *LabelsFilter) GetOperator() FilterOperator {
	if x != nil {
		return x.Operator
	}
	return FilterOperator_FILTER_OPERATOR_OR_OR_UNSPECIFIED
}

var File_com_coralogixapis_incidents_v1_incident_event_query_filter_proto protoreflect.FileDescriptor

const file_com_coralogixapis_incidents_v1_incident_event_query_filter_proto_rawDesc = "" +
	"\n" +
	"@com/coralogixapis/incidents/v1/incident_event_query_filter.proto\x12\x1ecom.coralogixapis.incidents.v1\x1a5com/coralogixapis/incidents/v1/filter_operators.proto\x1a:com/coralogixapis/incidents/v1/incident_query_filter.proto\x1a6com/coralogixapis/incidents/v1/incident_severity.proto\x1a4com/coralogixapis/incidents/v1/incident_status.proto\x1a/com/coralogixapis/incidents/v1/meta_label.proto\x1a\x1egoogle/protobuf/wrappers.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\xee\a\n" +
	"\x18IncidentEventQueryFilter\x12g\n" +
	"\x06status\x18\x01 \x03(\x0e2..com.coralogixapis.incidents.v1.IncidentStatusB\x1f\x92A\x1c2\x1aThe status of the incidentR\x06status\x12o\n" +
	"\bseverity\x18\x02 \x03(\x0e20.com.coralogixapis.incidents.v1.IncidentSeverityB!\x92A\x1e2\x1cThe severity of the incidentR\bseverity\x12\xa7\x01\n" +
	"\x11contextual_labels\x18\x03 \x03(\v2N.com.coralogixapis.incidents.v1.IncidentEventQueryFilter.ContextualLabelsEntryB*\x92A'2%The contextual labels of the incidentR\x10contextualLabels\x12O\n" +
	"\x04name\x18\x04 \x01(\v2\x1c.google.protobuf.StringValueB\x1d\x92A\x1a2\x18The name of the incidentR\x04name\x12^\n" +
	"\bis_muted\x18\x05 \x01(\v2\x1a.google.protobuf.BoolValueB'\x92A$2\"Indicates if the incident is mutedR\aisMuted\x12l\n" +
	"\ttimestamp\x18\x06 \x01(\v2).com.coralogixapis.incidents.v1.TimeRangeB#\x92A 2\x1eThe time range of the incidentR\ttimestamp\x12e\n" +
	"\x06labels\x18\a \x01(\v2,.com.coralogixapis.incidents.v1.LabelsFilterB\x1f\x92A\x1c2\x1aThe labels of the incidentR\x06labels\x1az\n" +
	"\x15ContextualLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12K\n" +
	"\x05value\x18\x02 \x01(\v25.com.coralogixapis.incidents.v1.ContextualLabelValuesR\x05value:\x028\x01:L\x92AI\n" +
	"G*\x1bIncident event query filter2(Filter configuration for incident events\"\xa6\x02\n" +
	"\fLabelsFilter\x12p\n" +
	"\vmeta_labels\x18\x01 \x03(\v2).com.coralogixapis.incidents.v1.MetaLabelB$\x92A!2\x1fThe meta labels of the incidentR\n" +
	"metaLabels\x12s\n" +
	"\boperator\x18\x02 \x01(\x0e2..com.coralogixapis.incidents.v1.FilterOperatorB'\x92A$2\"The operator for the labels filterR\boperator:/\x92A,\n" +
	"**\x1aLabel filter configuration\xd2\x01\vmeta_labelsb\x06proto3"

var (
	file_com_coralogixapis_incidents_v1_incident_event_query_filter_proto_rawDescOnce sync.Once
	file_com_coralogixapis_incidents_v1_incident_event_query_filter_proto_rawDescData []byte
)

func file_com_coralogixapis_incidents_v1_incident_event_query_filter_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_incidents_v1_incident_event_query_filter_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_incidents_v1_incident_event_query_filter_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_incidents_v1_incident_event_query_filter_proto_rawDesc), len(file_com_coralogixapis_incidents_v1_incident_event_query_filter_proto_rawDesc)))
	})
	return file_com_coralogixapis_incidents_v1_incident_event_query_filter_proto_rawDescData
}

var file_com_coralogixapis_incidents_v1_incident_event_query_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogixapis_incidents_v1_incident_event_query_filter_proto_goTypes = []any{
	(*IncidentEventQueryFilter)(nil), // 0: com.coralogixapis.incidents.v1.IncidentEventQueryFilter
	(*LabelsFilter)(nil),             // 1: com.coralogixapis.incidents.v1.LabelsFilter
	nil,                              // 2: com.coralogixapis.incidents.v1.IncidentEventQueryFilter.ContextualLabelsEntry
	(IncidentStatus)(0),              // 3: com.coralogixapis.incidents.v1.IncidentStatus
	(IncidentSeverity)(0),            // 4: com.coralogixapis.incidents.v1.IncidentSeverity
	(*wrapperspb.StringValue)(nil),   // 5: google.protobuf.StringValue
	(*wrapperspb.BoolValue)(nil),     // 6: google.protobuf.BoolValue
	(*TimeRange)(nil),                // 7: com.coralogixapis.incidents.v1.TimeRange
	(*MetaLabel)(nil),                // 8: com.coralogixapis.incidents.v1.MetaLabel
	(FilterOperator)(0),              // 9: com.coralogixapis.incidents.v1.FilterOperator
	(*ContextualLabelValues)(nil),    // 10: com.coralogixapis.incidents.v1.ContextualLabelValues
}
var file_com_coralogixapis_incidents_v1_incident_event_query_filter_proto_depIdxs = []int32{
	3,  // 0: com.coralogixapis.incidents.v1.IncidentEventQueryFilter.status:type_name -> com.coralogixapis.incidents.v1.IncidentStatus
	4,  // 1: com.coralogixapis.incidents.v1.IncidentEventQueryFilter.severity:type_name -> com.coralogixapis.incidents.v1.IncidentSeverity
	2,  // 2: com.coralogixapis.incidents.v1.IncidentEventQueryFilter.contextual_labels:type_name -> com.coralogixapis.incidents.v1.IncidentEventQueryFilter.ContextualLabelsEntry
	5,  // 3: com.coralogixapis.incidents.v1.IncidentEventQueryFilter.name:type_name -> google.protobuf.StringValue
	6,  // 4: com.coralogixapis.incidents.v1.IncidentEventQueryFilter.is_muted:type_name -> google.protobuf.BoolValue
	7,  // 5: com.coralogixapis.incidents.v1.IncidentEventQueryFilter.timestamp:type_name -> com.coralogixapis.incidents.v1.TimeRange
	1,  // 6: com.coralogixapis.incidents.v1.IncidentEventQueryFilter.labels:type_name -> com.coralogixapis.incidents.v1.LabelsFilter
	8,  // 7: com.coralogixapis.incidents.v1.LabelsFilter.meta_labels:type_name -> com.coralogixapis.incidents.v1.MetaLabel
	9,  // 8: com.coralogixapis.incidents.v1.LabelsFilter.operator:type_name -> com.coralogixapis.incidents.v1.FilterOperator
	10, // 9: com.coralogixapis.incidents.v1.IncidentEventQueryFilter.ContextualLabelsEntry.value:type_name -> com.coralogixapis.incidents.v1.ContextualLabelValues
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_incidents_v1_incident_event_query_filter_proto_init() }
func file_com_coralogixapis_incidents_v1_incident_event_query_filter_proto_init() {
	if File_com_coralogixapis_incidents_v1_incident_event_query_filter_proto != nil {
		return
	}
	file_com_coralogixapis_incidents_v1_filter_operators_proto_init()
	file_com_coralogixapis_incidents_v1_incident_query_filter_proto_init()
	file_com_coralogixapis_incidents_v1_incident_severity_proto_init()
	file_com_coralogixapis_incidents_v1_incident_status_proto_init()
	file_com_coralogixapis_incidents_v1_meta_label_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_incidents_v1_incident_event_query_filter_proto_rawDesc), len(file_com_coralogixapis_incidents_v1_incident_event_query_filter_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_incidents_v1_incident_event_query_filter_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_incidents_v1_incident_event_query_filter_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_incidents_v1_incident_event_query_filter_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_incidents_v1_incident_event_query_filter_proto = out.File
	file_com_coralogixapis_incidents_v1_incident_event_query_filter_proto_goTypes = nil
	file_com_coralogixapis_incidents_v1_incident_event_query_filter_proto_depIdxs = nil
}
