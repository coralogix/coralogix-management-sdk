// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: com/coralogixapis/alerts/v3/alert_def_type_definition/logs/common/logs_filter.proto

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

type LogFilterOperationType int32

const (
	LogFilterOperationType_LOG_FILTER_OPERATION_TYPE_IS_OR_UNSPECIFIED LogFilterOperationType = 0
	LogFilterOperationType_LOG_FILTER_OPERATION_TYPE_INCLUDES          LogFilterOperationType = 1
	LogFilterOperationType_LOG_FILTER_OPERATION_TYPE_ENDS_WITH         LogFilterOperationType = 2
	LogFilterOperationType_LOG_FILTER_OPERATION_TYPE_STARTS_WITH       LogFilterOperationType = 3
)

// Enum value maps for LogFilterOperationType.
var (
	LogFilterOperationType_name = map[int32]string{
		0: "LOG_FILTER_OPERATION_TYPE_IS_OR_UNSPECIFIED",
		1: "LOG_FILTER_OPERATION_TYPE_INCLUDES",
		2: "LOG_FILTER_OPERATION_TYPE_ENDS_WITH",
		3: "LOG_FILTER_OPERATION_TYPE_STARTS_WITH",
	}
	LogFilterOperationType_value = map[string]int32{
		"LOG_FILTER_OPERATION_TYPE_IS_OR_UNSPECIFIED": 0,
		"LOG_FILTER_OPERATION_TYPE_INCLUDES":          1,
		"LOG_FILTER_OPERATION_TYPE_ENDS_WITH":         2,
		"LOG_FILTER_OPERATION_TYPE_STARTS_WITH":       3,
	}
)

func (x LogFilterOperationType) Enum() *LogFilterOperationType {
	p := new(LogFilterOperationType)
	*p = x
	return p
}

func (x LogFilterOperationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogFilterOperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_enumTypes[0].Descriptor()
}

func (LogFilterOperationType) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_enumTypes[0]
}

func (x LogFilterOperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogFilterOperationType.Descriptor instead.
func (LogFilterOperationType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_rawDescGZIP(), []int{0}
}

type LogSeverity int32

const (
	LogSeverity_LOG_SEVERITY_VERBOSE_UNSPECIFIED LogSeverity = 0
	LogSeverity_LOG_SEVERITY_DEBUG               LogSeverity = 1
	LogSeverity_LOG_SEVERITY_INFO                LogSeverity = 2
	LogSeverity_LOG_SEVERITY_WARNING             LogSeverity = 3
	LogSeverity_LOG_SEVERITY_ERROR               LogSeverity = 4
	LogSeverity_LOG_SEVERITY_CRITICAL            LogSeverity = 5
)

// Enum value maps for LogSeverity.
var (
	LogSeverity_name = map[int32]string{
		0: "LOG_SEVERITY_VERBOSE_UNSPECIFIED",
		1: "LOG_SEVERITY_DEBUG",
		2: "LOG_SEVERITY_INFO",
		3: "LOG_SEVERITY_WARNING",
		4: "LOG_SEVERITY_ERROR",
		5: "LOG_SEVERITY_CRITICAL",
	}
	LogSeverity_value = map[string]int32{
		"LOG_SEVERITY_VERBOSE_UNSPECIFIED": 0,
		"LOG_SEVERITY_DEBUG":               1,
		"LOG_SEVERITY_INFO":                2,
		"LOG_SEVERITY_WARNING":             3,
		"LOG_SEVERITY_ERROR":               4,
		"LOG_SEVERITY_CRITICAL":            5,
	}
)

func (x LogSeverity) Enum() *LogSeverity {
	p := new(LogSeverity)
	*p = x
	return p
}

func (x LogSeverity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogSeverity) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_enumTypes[1].Descriptor()
}

func (LogSeverity) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_enumTypes[1]
}

func (x LogSeverity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogSeverity.Descriptor instead.
func (LogSeverity) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_rawDescGZIP(), []int{1}
}

type LogsFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to FilterType:
	//
	//	*LogsFilter_SimpleFilter
	FilterType isLogsFilter_FilterType `protobuf_oneof:"filter_type"`
}

func (x *LogsFilter) Reset() {
	*x = LogsFilter{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsFilter) ProtoMessage() {}

func (x *LogsFilter) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsFilter.ProtoReflect.Descriptor instead.
func (*LogsFilter) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_rawDescGZIP(), []int{0}
}

func (m *LogsFilter) GetFilterType() isLogsFilter_FilterType {
	if m != nil {
		return m.FilterType
	}
	return nil
}

func (x *LogsFilter) GetSimpleFilter() *LogsSimpleFilter {
	if x, ok := x.GetFilterType().(*LogsFilter_SimpleFilter); ok {
		return x.SimpleFilter
	}
	return nil
}

type isLogsFilter_FilterType interface {
	isLogsFilter_FilterType()
}

type LogsFilter_SimpleFilter struct {
	SimpleFilter *LogsSimpleFilter `protobuf:"bytes,1,opt,name=simple_filter,json=simpleFilter,proto3,oneof"`
}

func (*LogsFilter_SimpleFilter) isLogsFilter_FilterType() {}

type LogsSimpleFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LuceneQuery  *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=lucene_query,json=luceneQuery,proto3" json:"lucene_query,omitempty"`
	LabelFilters *LabelFilters           `protobuf:"bytes,2,opt,name=label_filters,json=labelFilters,proto3" json:"label_filters,omitempty"`
}

func (x *LogsSimpleFilter) Reset() {
	*x = LogsSimpleFilter{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsSimpleFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsSimpleFilter) ProtoMessage() {}

func (x *LogsSimpleFilter) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsSimpleFilter.ProtoReflect.Descriptor instead.
func (*LogsSimpleFilter) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_rawDescGZIP(), []int{1}
}

func (x *LogsSimpleFilter) GetLuceneQuery() *wrapperspb.StringValue {
	if x != nil {
		return x.LuceneQuery
	}
	return nil
}

func (x *LogsSimpleFilter) GetLabelFilters() *LabelFilters {
	if x != nil {
		return x.LabelFilters
	}
	return nil
}

type LabelFilters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplicationName []*LabelFilterType `protobuf:"bytes,1,rep,name=application_name,json=applicationName,proto3" json:"application_name,omitempty"`
	SubsystemName   []*LabelFilterType `protobuf:"bytes,2,rep,name=subsystem_name,json=subsystemName,proto3" json:"subsystem_name,omitempty"`
	Severities      []LogSeverity      `protobuf:"varint,3,rep,packed,name=severities,proto3,enum=com.coralogixapis.alerts.v3.LogSeverity" json:"severities,omitempty"`
}

func (x *LabelFilters) Reset() {
	*x = LabelFilters{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LabelFilters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelFilters) ProtoMessage() {}

func (x *LabelFilters) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelFilters.ProtoReflect.Descriptor instead.
func (*LabelFilters) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_rawDescGZIP(), []int{2}
}

func (x *LabelFilters) GetApplicationName() []*LabelFilterType {
	if x != nil {
		return x.ApplicationName
	}
	return nil
}

func (x *LabelFilters) GetSubsystemName() []*LabelFilterType {
	if x != nil {
		return x.SubsystemName
	}
	return nil
}

func (x *LabelFilters) GetSeverities() []LogSeverity {
	if x != nil {
		return x.Severities
	}
	return nil
}

type LabelFilterType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value     *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Operation LogFilterOperationType  `protobuf:"varint,2,opt,name=operation,proto3,enum=com.coralogixapis.alerts.v3.LogFilterOperationType" json:"operation,omitempty"`
}

func (x *LabelFilterType) Reset() {
	*x = LabelFilterType{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LabelFilterType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelFilterType) ProtoMessage() {}

func (x *LabelFilterType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelFilterType.ProtoReflect.Descriptor instead.
func (*LabelFilterType) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_rawDescGZIP(), []int{3}
}

func (x *LabelFilterType) GetValue() *wrapperspb.StringValue {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *LabelFilterType) GetOperation() LogFilterOperationType {
	if x != nil {
		return x.Operation
	}
	return LogFilterOperationType_LOG_FILTER_OPERATION_TYPE_IS_OR_UNSPECIFIED
}

var File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_rawDesc = []byte{
	0x0a, 0x53, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e,
	0x76, 0x33, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x71, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x54, 0x0a, 0x0d, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x0d, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0xa3, 0x01, 0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x73, 0x53, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x0c, 0x6c, 0x75,
	0x63, 0x65, 0x6e, 0x65, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b,
	0x6c, 0x75, 0x63, 0x65, 0x6e, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x4e, 0x0a, 0x0d, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x0c, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x22, 0x86, 0x02, 0x0a, 0x0c,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x57, 0x0a, 0x10,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x73, 0x75, 0x62,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x48, 0x0a, 0x0a, 0x73, 0x65,
	0x76, 0x65, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x28,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x67,
	0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x22, 0x98, 0x01, 0x0a, 0x0f, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x51, 0x0a, 0x09,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f,
	0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2a,
	0xc5, 0x01, 0x0a, 0x16, 0x4c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x2b, 0x4c, 0x4f,
	0x47, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x53, 0x5f, 0x4f, 0x52, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x26, 0x0a, 0x22, 0x4c,
	0x4f, 0x47, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x43, 0x4c, 0x55, 0x44, 0x45,
	0x53, 0x10, 0x01, 0x12, 0x27, 0x0a, 0x23, 0x4c, 0x4f, 0x47, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45,
	0x52, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x45, 0x4e, 0x44, 0x53, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x10, 0x02, 0x12, 0x29, 0x0a, 0x25,
	0x4c, 0x4f, 0x47, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x53,
	0x5f, 0x57, 0x49, 0x54, 0x48, 0x10, 0x03, 0x2a, 0xaf, 0x01, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x53,
	0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x24, 0x0a, 0x20, 0x4c, 0x4f, 0x47, 0x5f, 0x53,
	0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x56, 0x45, 0x52, 0x42, 0x4f, 0x53, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a,
	0x12, 0x4c, 0x4f, 0x47, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x44, 0x45,
	0x42, 0x55, 0x47, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x4f, 0x47, 0x5f, 0x53, 0x45, 0x56,
	0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14,
	0x4c, 0x4f, 0x47, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x57, 0x41, 0x52,
	0x4e, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x4c, 0x4f, 0x47, 0x5f, 0x53, 0x45,
	0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x12, 0x19,
	0x0a, 0x15, 0x4c, 0x4f, 0x47, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x43,
	0x52, 0x49, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x05, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_rawDescData = file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_goTypes = []any{
	(LogFilterOperationType)(0),    // 0: com.coralogixapis.alerts.v3.LogFilterOperationType
	(LogSeverity)(0),               // 1: com.coralogixapis.alerts.v3.LogSeverity
	(*LogsFilter)(nil),             // 2: com.coralogixapis.alerts.v3.LogsFilter
	(*LogsSimpleFilter)(nil),       // 3: com.coralogixapis.alerts.v3.LogsSimpleFilter
	(*LabelFilters)(nil),           // 4: com.coralogixapis.alerts.v3.LabelFilters
	(*LabelFilterType)(nil),        // 5: com.coralogixapis.alerts.v3.LabelFilterType
	(*wrapperspb.StringValue)(nil), // 6: google.protobuf.StringValue
}
var file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_depIdxs = []int32{
	3, // 0: com.coralogixapis.alerts.v3.LogsFilter.simple_filter:type_name -> com.coralogixapis.alerts.v3.LogsSimpleFilter
	6, // 1: com.coralogixapis.alerts.v3.LogsSimpleFilter.lucene_query:type_name -> google.protobuf.StringValue
	4, // 2: com.coralogixapis.alerts.v3.LogsSimpleFilter.label_filters:type_name -> com.coralogixapis.alerts.v3.LabelFilters
	5, // 3: com.coralogixapis.alerts.v3.LabelFilters.application_name:type_name -> com.coralogixapis.alerts.v3.LabelFilterType
	5, // 4: com.coralogixapis.alerts.v3.LabelFilters.subsystem_name:type_name -> com.coralogixapis.alerts.v3.LabelFilterType
	1, // 5: com.coralogixapis.alerts.v3.LabelFilters.severities:type_name -> com.coralogixapis.alerts.v3.LogSeverity
	6, // 6: com.coralogixapis.alerts.v3.LabelFilterType.value:type_name -> google.protobuf.StringValue
	0, // 7: com.coralogixapis.alerts.v3.LabelFilterType.operation:type_name -> com.coralogixapis.alerts.v3.LogFilterOperationType
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_init()
}
func file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_msgTypes[0].OneofWrappers = []any{
		(*LogsFilter_SimpleFilter)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_common_logs_filter_proto_depIdxs = nil
}
