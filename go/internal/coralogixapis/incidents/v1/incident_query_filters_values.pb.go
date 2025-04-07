// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.26.1
// source: com/coralogixapis/incidents/v1/incident_query_filters_values.proto

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

type AssigneeWithCount struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Assignee      *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=assignee,proto3" json:"assignee,omitempty"`
	Count         *wrapperspb.Int32Value  `protobuf:"bytes,2,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AssigneeWithCount) Reset() {
	*x = AssigneeWithCount{}
	mi := &file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssigneeWithCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssigneeWithCount) ProtoMessage() {}

func (x *AssigneeWithCount) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssigneeWithCount.ProtoReflect.Descriptor instead.
func (*AssigneeWithCount) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_rawDescGZIP(), []int{0}
}

func (x *AssigneeWithCount) GetAssignee() *wrapperspb.StringValue {
	if x != nil {
		return x.Assignee
	}
	return nil
}

func (x *AssigneeWithCount) GetCount() *wrapperspb.Int32Value {
	if x != nil {
		return x.Count
	}
	return nil
}

type IncidentStatusWithCount struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        IncidentStatus         `protobuf:"varint,1,opt,name=status,proto3,enum=com.coralogixapis.incidents.v1.IncidentStatus" json:"status,omitempty"`
	Count         *wrapperspb.Int32Value `protobuf:"bytes,2,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IncidentStatusWithCount) Reset() {
	*x = IncidentStatusWithCount{}
	mi := &file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IncidentStatusWithCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncidentStatusWithCount) ProtoMessage() {}

func (x *IncidentStatusWithCount) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncidentStatusWithCount.ProtoReflect.Descriptor instead.
func (*IncidentStatusWithCount) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_rawDescGZIP(), []int{1}
}

func (x *IncidentStatusWithCount) GetStatus() IncidentStatus {
	if x != nil {
		return x.Status
	}
	return IncidentStatus_INCIDENT_STATUS_UNSPECIFIED
}

func (x *IncidentStatusWithCount) GetCount() *wrapperspb.Int32Value {
	if x != nil {
		return x.Count
	}
	return nil
}

type IncidentStateWithCount struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	State         IncidentState          `protobuf:"varint,1,opt,name=state,proto3,enum=com.coralogixapis.incidents.v1.IncidentState" json:"state,omitempty"`
	Count         *wrapperspb.Int32Value `protobuf:"bytes,2,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IncidentStateWithCount) Reset() {
	*x = IncidentStateWithCount{}
	mi := &file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IncidentStateWithCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncidentStateWithCount) ProtoMessage() {}

func (x *IncidentStateWithCount) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncidentStateWithCount.ProtoReflect.Descriptor instead.
func (*IncidentStateWithCount) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_rawDescGZIP(), []int{2}
}

func (x *IncidentStateWithCount) GetState() IncidentState {
	if x != nil {
		return x.State
	}
	return IncidentState_INCIDENT_STATE_UNSPECIFIED
}

func (x *IncidentStateWithCount) GetCount() *wrapperspb.Int32Value {
	if x != nil {
		return x.Count
	}
	return nil
}

type IncidentSeverityWithCount struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Severity      IncidentSeverity       `protobuf:"varint,1,opt,name=severity,proto3,enum=com.coralogixapis.incidents.v1.IncidentSeverity" json:"severity,omitempty"`
	Count         *wrapperspb.Int32Value `protobuf:"bytes,2,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IncidentSeverityWithCount) Reset() {
	*x = IncidentSeverityWithCount{}
	mi := &file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IncidentSeverityWithCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncidentSeverityWithCount) ProtoMessage() {}

func (x *IncidentSeverityWithCount) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncidentSeverityWithCount.ProtoReflect.Descriptor instead.
func (*IncidentSeverityWithCount) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_rawDescGZIP(), []int{3}
}

func (x *IncidentSeverityWithCount) GetSeverity() IncidentSeverity {
	if x != nil {
		return x.Severity
	}
	return IncidentSeverity_INCIDENT_SEVERITY_UNSPECIFIED
}

func (x *IncidentSeverityWithCount) GetCount() *wrapperspb.Int32Value {
	if x != nil {
		return x.Count
	}
	return nil
}

type IncidentMetaLabelsWithCount struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MetaLabel     *MetaLabel             `protobuf:"bytes,1,opt,name=meta_label,json=metaLabel,proto3" json:"meta_label,omitempty"`
	Count         *wrapperspb.Int32Value `protobuf:"bytes,2,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IncidentMetaLabelsWithCount) Reset() {
	*x = IncidentMetaLabelsWithCount{}
	mi := &file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IncidentMetaLabelsWithCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncidentMetaLabelsWithCount) ProtoMessage() {}

func (x *IncidentMetaLabelsWithCount) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncidentMetaLabelsWithCount.ProtoReflect.Descriptor instead.
func (*IncidentMetaLabelsWithCount) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_rawDescGZIP(), []int{4}
}

func (x *IncidentMetaLabelsWithCount) GetMetaLabel() *MetaLabel {
	if x != nil {
		return x.MetaLabel
	}
	return nil
}

func (x *IncidentMetaLabelsWithCount) GetCount() *wrapperspb.Int32Value {
	if x != nil {
		return x.Count
	}
	return nil
}

type IncidentQueryFiltersValues struct {
	state               protoimpl.MessageState                     `protogen:"open.v1"`
	AssigneeWithCount   []*AssigneeWithCount                       `protobuf:"bytes,1,rep,name=assignee_with_count,json=assigneeWithCount,proto3" json:"assignee_with_count,omitempty"`
	StatusWithCount     []*IncidentStatusWithCount                 `protobuf:"bytes,2,rep,name=status_with_count,json=statusWithCount,proto3" json:"status_with_count,omitempty"`
	StateWithCount      []*IncidentStateWithCount                  `protobuf:"bytes,3,rep,name=state_with_count,json=stateWithCount,proto3" json:"state_with_count,omitempty"`
	SeverityWithCount   []*IncidentSeverityWithCount               `protobuf:"bytes,4,rep,name=severity_with_count,json=severityWithCount,proto3" json:"severity_with_count,omitempty"`
	ContextualLabels    map[string]*ContextualLabelValuesWithCount `protobuf:"bytes,5,rep,name=contextual_labels,json=contextualLabels,proto3" json:"contextual_labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	MetaLabelsWithCount []*IncidentMetaLabelsWithCount             `protobuf:"bytes,6,rep,name=meta_labels_with_count,json=metaLabelsWithCount,proto3" json:"meta_labels_with_count,omitempty"`
	MetaLabelsOp        FilterOperator                             `protobuf:"varint,7,opt,name=meta_labels_op,json=metaLabelsOp,proto3,enum=com.coralogixapis.incidents.v1.FilterOperator" json:"meta_labels_op,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *IncidentQueryFiltersValues) Reset() {
	*x = IncidentQueryFiltersValues{}
	mi := &file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IncidentQueryFiltersValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncidentQueryFiltersValues) ProtoMessage() {}

func (x *IncidentQueryFiltersValues) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncidentQueryFiltersValues.ProtoReflect.Descriptor instead.
func (*IncidentQueryFiltersValues) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_rawDescGZIP(), []int{5}
}

func (x *IncidentQueryFiltersValues) GetAssigneeWithCount() []*AssigneeWithCount {
	if x != nil {
		return x.AssigneeWithCount
	}
	return nil
}

func (x *IncidentQueryFiltersValues) GetStatusWithCount() []*IncidentStatusWithCount {
	if x != nil {
		return x.StatusWithCount
	}
	return nil
}

func (x *IncidentQueryFiltersValues) GetStateWithCount() []*IncidentStateWithCount {
	if x != nil {
		return x.StateWithCount
	}
	return nil
}

func (x *IncidentQueryFiltersValues) GetSeverityWithCount() []*IncidentSeverityWithCount {
	if x != nil {
		return x.SeverityWithCount
	}
	return nil
}

func (x *IncidentQueryFiltersValues) GetContextualLabels() map[string]*ContextualLabelValuesWithCount {
	if x != nil {
		return x.ContextualLabels
	}
	return nil
}

func (x *IncidentQueryFiltersValues) GetMetaLabelsWithCount() []*IncidentMetaLabelsWithCount {
	if x != nil {
		return x.MetaLabelsWithCount
	}
	return nil
}

func (x *IncidentQueryFiltersValues) GetMetaLabelsOp() FilterOperator {
	if x != nil {
		return x.MetaLabelsOp
	}
	return FilterOperator_FILTER_OPERATOR_OR_OR_UNSPECIFIED
}

type ContextualLabelValuesWithCount struct {
	state           protoimpl.MessageState           `protogen:"open.v1"`
	ValuesWithCount []*ContextualLabelValueWithCount `protobuf:"bytes,1,rep,name=values_with_count,json=valuesWithCount,proto3" json:"values_with_count,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ContextualLabelValuesWithCount) Reset() {
	*x = ContextualLabelValuesWithCount{}
	mi := &file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContextualLabelValuesWithCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextualLabelValuesWithCount) ProtoMessage() {}

func (x *ContextualLabelValuesWithCount) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextualLabelValuesWithCount.ProtoReflect.Descriptor instead.
func (*ContextualLabelValuesWithCount) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_rawDescGZIP(), []int{6}
}

func (x *ContextualLabelValuesWithCount) GetValuesWithCount() []*ContextualLabelValueWithCount {
	if x != nil {
		return x.ValuesWithCount
	}
	return nil
}

type ContextualLabelValueWithCount struct {
	state                protoimpl.MessageState  `protogen:"open.v1"`
	ContextualLabelValue *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=contextual_label_value,json=contextualLabelValue,proto3" json:"contextual_label_value,omitempty"`
	Count                *wrapperspb.Int32Value  `protobuf:"bytes,2,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *ContextualLabelValueWithCount) Reset() {
	*x = ContextualLabelValueWithCount{}
	mi := &file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContextualLabelValueWithCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextualLabelValueWithCount) ProtoMessage() {}

func (x *ContextualLabelValueWithCount) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextualLabelValueWithCount.ProtoReflect.Descriptor instead.
func (*ContextualLabelValueWithCount) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_rawDescGZIP(), []int{7}
}

func (x *ContextualLabelValueWithCount) GetContextualLabelValue() *wrapperspb.StringValue {
	if x != nil {
		return x.ContextualLabelValue
	}
	return nil
}

func (x *ContextualLabelValueWithCount) GetCount() *wrapperspb.Int32Value {
	if x != nil {
		return x.Count
	}
	return nil
}

var File_com_coralogixapis_incidents_v1_incident_query_filters_values_proto protoreflect.FileDescriptor

const file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_rawDesc = "" +
	"\n" +
	"Bcom/coralogixapis/incidents/v1/incident_query_filters_values.proto\x12\x1ecom.coralogixapis.incidents.v1\x1a5com/coralogixapis/incidents/v1/filter_operators.proto\x1a6com/coralogixapis/incidents/v1/incident_severity.proto\x1a4com/coralogixapis/incidents/v1/incident_status.proto\x1a3com/coralogixapis/incidents/v1/incident_state.proto\x1a/com/coralogixapis/incidents/v1/meta_label.proto\x1a\x1egoogle/protobuf/wrappers.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\xdf\x01\n" +
	"\x11AssigneeWithCount\x12I\n" +
	"\bassignee\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueB\x0f\x92A\fJ\n" +
	"\"assignee\"R\bassignee\x12:\n" +
	"\x05count\x18\x02 \x01(\v2\x1b.google.protobuf.Int32ValueB\a\x92A\x04J\x0210R\x05count:C\x92A@\n" +
	">*\x11AssigneeWithCount2\x16An assignee with count\xd2\x01\bassignee\xd2\x01\x05count\"\x85\x02\n" +
	"\x17IncidentStatusWithCount\x12^\n" +
	"\x06status\x18\x01 \x01(\x0e2..com.coralogixapis.incidents.v1.IncidentStatusB\x16\x92A\x13J\x11\"incident_status\"R\x06status\x12:\n" +
	"\x05count\x18\x02 \x01(\v2\x1b.google.protobuf.Int32ValueB\a\x92A\x04J\x0210R\x05count:N\x92AK\n" +
	"I*\x17IncidentStatusWithCount2\x1dAn incident status with count\xd2\x01\x06status\xd2\x01\x05count\"\xe6\x01\n" +
	"\x16IncidentStateWithCount\x12C\n" +
	"\x05state\x18\x01 \x01(\x0e2-.com.coralogixapis.incidents.v1.IncidentStateR\x05state\x12:\n" +
	"\x05count\x18\x02 \x01(\v2\x1b.google.protobuf.Int32ValueB\a\x92A\x04J\x0210R\x05count:K\x92AH\n" +
	"F*\x16IncidentStateWithCount2\x1cAn incident state with count\xd2\x01\x05state\xd2\x01\x05count\"\xfb\x01\n" +
	"\x19IncidentSeverityWithCount\x12L\n" +
	"\bseverity\x18\x01 \x01(\x0e20.com.coralogixapis.incidents.v1.IncidentSeverityR\bseverity\x12:\n" +
	"\x05count\x18\x02 \x01(\v2\x1b.google.protobuf.Int32ValueB\a\x92A\x04J\x0210R\x05count:T\x92AQ\n" +
	"O*\x19IncidentSeverityWithCount2\x1fAn incident severity with count\xd2\x01\bseverity\xd2\x01\x05count\"\x80\x02\n" +
	"\x1bIncidentMetaLabelsWithCount\x12H\n" +
	"\n" +
	"meta_label\x18\x01 \x01(\v2).com.coralogixapis.incidents.v1.MetaLabelR\tmetaLabel\x12:\n" +
	"\x05count\x18\x02 \x01(\v2\x1b.google.protobuf.Int32ValueB\a\x92A\x04J\x0210R\x05count:[\x92AX\n" +
	"V*\x1bIncidentMetaLabelsWithCount2\"An incident meta labels with count\xd2\x01\n" +
	"meta_label\xd2\x01\x05count\"\xd7\b\n" +
	"\x1aIncidentQueryFiltersValues\x12a\n" +
	"\x13assignee_with_count\x18\x01 \x03(\v21.com.coralogixapis.incidents.v1.AssigneeWithCountR\x11assigneeWithCount\x12c\n" +
	"\x11status_with_count\x18\x02 \x03(\v27.com.coralogixapis.incidents.v1.IncidentStatusWithCountR\x0fstatusWithCount\x12`\n" +
	"\x10state_with_count\x18\x03 \x03(\v26.com.coralogixapis.incidents.v1.IncidentStateWithCountR\x0estateWithCount\x12i\n" +
	"\x13severity_with_count\x18\x04 \x03(\v29.com.coralogixapis.incidents.v1.IncidentSeverityWithCountR\x11severityWithCount\x12}\n" +
	"\x11contextual_labels\x18\x05 \x03(\v2P.com.coralogixapis.incidents.v1.IncidentQueryFiltersValues.ContextualLabelsEntryR\x10contextualLabels\x12p\n" +
	"\x16meta_labels_with_count\x18\x06 \x03(\v2;.com.coralogixapis.incidents.v1.IncidentMetaLabelsWithCountR\x13metaLabelsWithCount\x12T\n" +
	"\x0emeta_labels_op\x18\a \x01(\x0e2..com.coralogixapis.incidents.v1.FilterOperatorR\fmetaLabelsOp\x1a\x83\x01\n" +
	"\x15ContextualLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12T\n" +
	"\x05value\x18\x02 \x01(\v2>.com.coralogixapis.incidents.v1.ContextualLabelValuesWithCountR\x05value:\x028\x01:\xd6\x01\x92A\xd2\x01\n" +
	"\xcf\x01*\x1aIncidentQueryFiltersValues2 An incident query filters values\xd2\x01\x13assignee_with_count\xd2\x01\x11status_with_count\xd2\x01\x10state_with_count\xd2\x01\x13severity_with_count\xd2\x01\x11contextual_labels\xd2\x01\x16meta_labels_with_count\xd2\x01\x0emeta_labels_op\"\xec\x01\n" +
	"\x1eContextualLabelValuesWithCount\x12i\n" +
	"\x11values_with_count\x18\x01 \x03(\v2=.com.coralogixapis.incidents.v1.ContextualLabelValueWithCountR\x0fvaluesWithCount:_\x92A\\\n" +
	"Z*\x1eContextualLabelValuesWithCount2$A contextual label values with count\xd2\x01\x11values_with_count\"\xba\x02\n" +
	"\x1dContextualLabelValueWithCount\x12q\n" +
	"\x16contextual_label_value\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueB\x1d\x92A\x1aJ\x18\"contextual_label_value\"R\x14contextualLabelValue\x12:\n" +
	"\x05count\x18\x02 \x01(\v2\x1b.google.protobuf.Int32ValueB\a\x92A\x04J\x0210R\x05count:j\x92Ag\n" +
	"e*\x1dContextualLabelValueWithCount2#A contextual label value with count\xd2\x01\x16contextual_label_value\xd2\x01\x05countb\x06proto3"

var (
	file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_rawDescOnce sync.Once
	file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_rawDescData []byte
)

func file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_rawDesc), len(file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_rawDesc)))
	})
	return file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_rawDescData
}

var file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_goTypes = []any{
	(*AssigneeWithCount)(nil),              // 0: com.coralogixapis.incidents.v1.AssigneeWithCount
	(*IncidentStatusWithCount)(nil),        // 1: com.coralogixapis.incidents.v1.IncidentStatusWithCount
	(*IncidentStateWithCount)(nil),         // 2: com.coralogixapis.incidents.v1.IncidentStateWithCount
	(*IncidentSeverityWithCount)(nil),      // 3: com.coralogixapis.incidents.v1.IncidentSeverityWithCount
	(*IncidentMetaLabelsWithCount)(nil),    // 4: com.coralogixapis.incidents.v1.IncidentMetaLabelsWithCount
	(*IncidentQueryFiltersValues)(nil),     // 5: com.coralogixapis.incidents.v1.IncidentQueryFiltersValues
	(*ContextualLabelValuesWithCount)(nil), // 6: com.coralogixapis.incidents.v1.ContextualLabelValuesWithCount
	(*ContextualLabelValueWithCount)(nil),  // 7: com.coralogixapis.incidents.v1.ContextualLabelValueWithCount
	nil,                                    // 8: com.coralogixapis.incidents.v1.IncidentQueryFiltersValues.ContextualLabelsEntry
	(*wrapperspb.StringValue)(nil),         // 9: google.protobuf.StringValue
	(*wrapperspb.Int32Value)(nil),          // 10: google.protobuf.Int32Value
	(IncidentStatus)(0),                    // 11: com.coralogixapis.incidents.v1.IncidentStatus
	(IncidentState)(0),                     // 12: com.coralogixapis.incidents.v1.IncidentState
	(IncidentSeverity)(0),                  // 13: com.coralogixapis.incidents.v1.IncidentSeverity
	(*MetaLabel)(nil),                      // 14: com.coralogixapis.incidents.v1.MetaLabel
	(FilterOperator)(0),                    // 15: com.coralogixapis.incidents.v1.FilterOperator
}
var file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_depIdxs = []int32{
	9,  // 0: com.coralogixapis.incidents.v1.AssigneeWithCount.assignee:type_name -> google.protobuf.StringValue
	10, // 1: com.coralogixapis.incidents.v1.AssigneeWithCount.count:type_name -> google.protobuf.Int32Value
	11, // 2: com.coralogixapis.incidents.v1.IncidentStatusWithCount.status:type_name -> com.coralogixapis.incidents.v1.IncidentStatus
	10, // 3: com.coralogixapis.incidents.v1.IncidentStatusWithCount.count:type_name -> google.protobuf.Int32Value
	12, // 4: com.coralogixapis.incidents.v1.IncidentStateWithCount.state:type_name -> com.coralogixapis.incidents.v1.IncidentState
	10, // 5: com.coralogixapis.incidents.v1.IncidentStateWithCount.count:type_name -> google.protobuf.Int32Value
	13, // 6: com.coralogixapis.incidents.v1.IncidentSeverityWithCount.severity:type_name -> com.coralogixapis.incidents.v1.IncidentSeverity
	10, // 7: com.coralogixapis.incidents.v1.IncidentSeverityWithCount.count:type_name -> google.protobuf.Int32Value
	14, // 8: com.coralogixapis.incidents.v1.IncidentMetaLabelsWithCount.meta_label:type_name -> com.coralogixapis.incidents.v1.MetaLabel
	10, // 9: com.coralogixapis.incidents.v1.IncidentMetaLabelsWithCount.count:type_name -> google.protobuf.Int32Value
	0,  // 10: com.coralogixapis.incidents.v1.IncidentQueryFiltersValues.assignee_with_count:type_name -> com.coralogixapis.incidents.v1.AssigneeWithCount
	1,  // 11: com.coralogixapis.incidents.v1.IncidentQueryFiltersValues.status_with_count:type_name -> com.coralogixapis.incidents.v1.IncidentStatusWithCount
	2,  // 12: com.coralogixapis.incidents.v1.IncidentQueryFiltersValues.state_with_count:type_name -> com.coralogixapis.incidents.v1.IncidentStateWithCount
	3,  // 13: com.coralogixapis.incidents.v1.IncidentQueryFiltersValues.severity_with_count:type_name -> com.coralogixapis.incidents.v1.IncidentSeverityWithCount
	8,  // 14: com.coralogixapis.incidents.v1.IncidentQueryFiltersValues.contextual_labels:type_name -> com.coralogixapis.incidents.v1.IncidentQueryFiltersValues.ContextualLabelsEntry
	4,  // 15: com.coralogixapis.incidents.v1.IncidentQueryFiltersValues.meta_labels_with_count:type_name -> com.coralogixapis.incidents.v1.IncidentMetaLabelsWithCount
	15, // 16: com.coralogixapis.incidents.v1.IncidentQueryFiltersValues.meta_labels_op:type_name -> com.coralogixapis.incidents.v1.FilterOperator
	7,  // 17: com.coralogixapis.incidents.v1.ContextualLabelValuesWithCount.values_with_count:type_name -> com.coralogixapis.incidents.v1.ContextualLabelValueWithCount
	9,  // 18: com.coralogixapis.incidents.v1.ContextualLabelValueWithCount.contextual_label_value:type_name -> google.protobuf.StringValue
	10, // 19: com.coralogixapis.incidents.v1.ContextualLabelValueWithCount.count:type_name -> google.protobuf.Int32Value
	6,  // 20: com.coralogixapis.incidents.v1.IncidentQueryFiltersValues.ContextualLabelsEntry.value:type_name -> com.coralogixapis.incidents.v1.ContextualLabelValuesWithCount
	21, // [21:21] is the sub-list for method output_type
	21, // [21:21] is the sub-list for method input_type
	21, // [21:21] is the sub-list for extension type_name
	21, // [21:21] is the sub-list for extension extendee
	0,  // [0:21] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_init() }
func file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_init() {
	if File_com_coralogixapis_incidents_v1_incident_query_filters_values_proto != nil {
		return
	}
	file_com_coralogixapis_incidents_v1_filter_operators_proto_init()
	file_com_coralogixapis_incidents_v1_incident_severity_proto_init()
	file_com_coralogixapis_incidents_v1_incident_status_proto_init()
	file_com_coralogixapis_incidents_v1_incident_state_proto_init()
	file_com_coralogixapis_incidents_v1_meta_label_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_rawDesc), len(file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_incidents_v1_incident_query_filters_values_proto = out.File
	file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_goTypes = nil
	file_com_coralogixapis_incidents_v1_incident_query_filters_values_proto_depIdxs = nil
}
