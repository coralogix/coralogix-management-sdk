// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/alerts/v3/alert_def.proto

package v3

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

// *
// Represents an existing or created alert definition
type AlertDef struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	AlertDefProperties *AlertDefProperties    `protobuf:"bytes,1,opt,name=alert_def_properties,json=alertDefProperties,proto3" json:"alert_def_properties,omitempty"`
	// This is the alert definition's persistent ID (does not change on replace), AKA UniqueIdentifier
	Id             *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	AlertVersionId *wrapperspb.StringValue `protobuf:"bytes,20,opt,name=alert_version_id,json=alertVersionId,proto3" json:"alert_version_id,omitempty"` // the old alertId
	CreatedTime    *timestamppb.Timestamp  `protobuf:"bytes,3,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	UpdatedTime    *timestamppb.Timestamp  `protobuf:"bytes,4,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *AlertDef) Reset() {
	*x = AlertDef{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlertDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertDef) ProtoMessage() {}

func (x *AlertDef) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertDef.ProtoReflect.Descriptor instead.
func (*AlertDef) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_proto_rawDescGZIP(), []int{0}
}

func (x *AlertDef) GetAlertDefProperties() *AlertDefProperties {
	if x != nil {
		return x.AlertDefProperties
	}
	return nil
}

func (x *AlertDef) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *AlertDef) GetAlertVersionId() *wrapperspb.StringValue {
	if x != nil {
		return x.AlertVersionId
	}
	return nil
}

func (x *AlertDef) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *AlertDef) GetUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

// *
// Represents the non generated alert definition properties (the ones that are set by the user)
type AlertDefProperties struct {
	state       protoimpl.MessageState  `protogen:"open.v1"`
	Name        *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Enabled     *wrapperspb.BoolValue   `protobuf:"bytes,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Priority    AlertDefPriority        `protobuf:"varint,4,opt,name=priority,proto3,enum=com.coralogixapis.alerts.v3.AlertDefPriority" json:"priority,omitempty"`
	// Types that are valid to be assigned to Schedule:
	//
	//	*AlertDefProperties_ActiveOn
	Schedule isAlertDefProperties_Schedule `protobuf_oneof:"schedule"`
	Type     AlertDefType                  `protobuf:"varint,6,opt,name=type,proto3,enum=com.coralogixapis.alerts.v3.AlertDefType" json:"type,omitempty"`
	// Types that are valid to be assigned to TypeDefinition:
	//
	//	*AlertDefProperties_LogsImmediate
	//	*AlertDefProperties_TracingImmediate
	//	*AlertDefProperties_LogsThreshold
	//	*AlertDefProperties_LogsRatioThreshold
	//	*AlertDefProperties_LogsTimeRelativeThreshold
	//	*AlertDefProperties_MetricThreshold
	//	*AlertDefProperties_TracingThreshold
	//	*AlertDefProperties_Flow
	//	*AlertDefProperties_LogsAnomaly
	//	*AlertDefProperties_MetricAnomaly
	//	*AlertDefProperties_LogsNewValue
	//	*AlertDefProperties_LogsUniqueCount
	TypeDefinition    isAlertDefProperties_TypeDefinition `protobuf_oneof:"type_definition"`
	GroupByKeys       []*wrapperspb.StringValue           `protobuf:"bytes,7,rep,name=group_by_keys,json=groupByKeys,proto3" json:"group_by_keys,omitempty"`
	IncidentsSettings *AlertDefIncidentSettings           `protobuf:"bytes,8,opt,name=incidents_settings,json=incidentsSettings,proto3" json:"incidents_settings,omitempty"`
	NotificationGroup *AlertDefNotificationGroup          `protobuf:"bytes,9,opt,name=notification_group,json=notificationGroup,proto3" json:"notification_group,omitempty"`
	// Deprecated: Marked as deprecated in com/coralogixapis/alerts/v3/alert_def.proto.
	NotificationGroupExcess []*AlertDefNotificationGroup `protobuf:"bytes,210,rep,name=notification_group_excess,json=notificationGroupExcess,proto3" json:"notification_group_excess,omitempty"`
	EntityLabels            map[string]string            `protobuf:"bytes,10,rep,name=entity_labels,json=entityLabels,proto3" json:"entity_labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	PhantomMode             *wrapperspb.BoolValue        `protobuf:"bytes,11,opt,name=phantom_mode,json=phantomMode,proto3" json:"phantom_mode,omitempty"`
	Deleted                 *wrapperspb.BoolValue        `protobuf:"bytes,12,opt,name=deleted,proto3" json:"deleted,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *AlertDefProperties) Reset() {
	*x = AlertDefProperties{}
	mi := &file_com_coralogixapis_alerts_v3_alert_def_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlertDefProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertDefProperties) ProtoMessage() {}

func (x *AlertDefProperties) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_alert_def_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertDefProperties.ProtoReflect.Descriptor instead.
func (*AlertDefProperties) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_alert_def_proto_rawDescGZIP(), []int{1}
}

func (x *AlertDefProperties) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *AlertDefProperties) GetDescription() *wrapperspb.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *AlertDefProperties) GetEnabled() *wrapperspb.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

func (x *AlertDefProperties) GetPriority() AlertDefPriority {
	if x != nil {
		return x.Priority
	}
	return AlertDefPriority_ALERT_DEF_PRIORITY_P5_OR_UNSPECIFIED
}

func (x *AlertDefProperties) GetSchedule() isAlertDefProperties_Schedule {
	if x != nil {
		return x.Schedule
	}
	return nil
}

func (x *AlertDefProperties) GetActiveOn() *ActivitySchedule {
	if x != nil {
		if x, ok := x.Schedule.(*AlertDefProperties_ActiveOn); ok {
			return x.ActiveOn
		}
	}
	return nil
}

func (x *AlertDefProperties) GetType() AlertDefType {
	if x != nil {
		return x.Type
	}
	return AlertDefType_ALERT_DEF_TYPE_LOGS_IMMEDIATE_OR_UNSPECIFIED
}

func (x *AlertDefProperties) GetTypeDefinition() isAlertDefProperties_TypeDefinition {
	if x != nil {
		return x.TypeDefinition
	}
	return nil
}

func (x *AlertDefProperties) GetLogsImmediate() *LogsImmediateType {
	if x != nil {
		if x, ok := x.TypeDefinition.(*AlertDefProperties_LogsImmediate); ok {
			return x.LogsImmediate
		}
	}
	return nil
}

func (x *AlertDefProperties) GetTracingImmediate() *TracingImmediateType {
	if x != nil {
		if x, ok := x.TypeDefinition.(*AlertDefProperties_TracingImmediate); ok {
			return x.TracingImmediate
		}
	}
	return nil
}

func (x *AlertDefProperties) GetLogsThreshold() *LogsThresholdType {
	if x != nil {
		if x, ok := x.TypeDefinition.(*AlertDefProperties_LogsThreshold); ok {
			return x.LogsThreshold
		}
	}
	return nil
}

func (x *AlertDefProperties) GetLogsRatioThreshold() *LogsRatioThresholdType {
	if x != nil {
		if x, ok := x.TypeDefinition.(*AlertDefProperties_LogsRatioThreshold); ok {
			return x.LogsRatioThreshold
		}
	}
	return nil
}

func (x *AlertDefProperties) GetLogsTimeRelativeThreshold() *LogsTimeRelativeThresholdType {
	if x != nil {
		if x, ok := x.TypeDefinition.(*AlertDefProperties_LogsTimeRelativeThreshold); ok {
			return x.LogsTimeRelativeThreshold
		}
	}
	return nil
}

func (x *AlertDefProperties) GetMetricThreshold() *MetricThresholdType {
	if x != nil {
		if x, ok := x.TypeDefinition.(*AlertDefProperties_MetricThreshold); ok {
			return x.MetricThreshold
		}
	}
	return nil
}

func (x *AlertDefProperties) GetTracingThreshold() *TracingThresholdType {
	if x != nil {
		if x, ok := x.TypeDefinition.(*AlertDefProperties_TracingThreshold); ok {
			return x.TracingThreshold
		}
	}
	return nil
}

func (x *AlertDefProperties) GetFlow() *FlowType {
	if x != nil {
		if x, ok := x.TypeDefinition.(*AlertDefProperties_Flow); ok {
			return x.Flow
		}
	}
	return nil
}

func (x *AlertDefProperties) GetLogsAnomaly() *LogsAnomalyType {
	if x != nil {
		if x, ok := x.TypeDefinition.(*AlertDefProperties_LogsAnomaly); ok {
			return x.LogsAnomaly
		}
	}
	return nil
}

func (x *AlertDefProperties) GetMetricAnomaly() *MetricAnomalyType {
	if x != nil {
		if x, ok := x.TypeDefinition.(*AlertDefProperties_MetricAnomaly); ok {
			return x.MetricAnomaly
		}
	}
	return nil
}

func (x *AlertDefProperties) GetLogsNewValue() *LogsNewValueType {
	if x != nil {
		if x, ok := x.TypeDefinition.(*AlertDefProperties_LogsNewValue); ok {
			return x.LogsNewValue
		}
	}
	return nil
}

func (x *AlertDefProperties) GetLogsUniqueCount() *LogsUniqueCountType {
	if x != nil {
		if x, ok := x.TypeDefinition.(*AlertDefProperties_LogsUniqueCount); ok {
			return x.LogsUniqueCount
		}
	}
	return nil
}

func (x *AlertDefProperties) GetGroupByKeys() []*wrapperspb.StringValue {
	if x != nil {
		return x.GroupByKeys
	}
	return nil
}

func (x *AlertDefProperties) GetIncidentsSettings() *AlertDefIncidentSettings {
	if x != nil {
		return x.IncidentsSettings
	}
	return nil
}

func (x *AlertDefProperties) GetNotificationGroup() *AlertDefNotificationGroup {
	if x != nil {
		return x.NotificationGroup
	}
	return nil
}

// Deprecated: Marked as deprecated in com/coralogixapis/alerts/v3/alert_def.proto.
func (x *AlertDefProperties) GetNotificationGroupExcess() []*AlertDefNotificationGroup {
	if x != nil {
		return x.NotificationGroupExcess
	}
	return nil
}

func (x *AlertDefProperties) GetEntityLabels() map[string]string {
	if x != nil {
		return x.EntityLabels
	}
	return nil
}

func (x *AlertDefProperties) GetPhantomMode() *wrapperspb.BoolValue {
	if x != nil {
		return x.PhantomMode
	}
	return nil
}

func (x *AlertDefProperties) GetDeleted() *wrapperspb.BoolValue {
	if x != nil {
		return x.Deleted
	}
	return nil
}

type isAlertDefProperties_Schedule interface {
	isAlertDefProperties_Schedule()
}

type AlertDefProperties_ActiveOn struct {
	ActiveOn *ActivitySchedule `protobuf:"bytes,5,opt,name=active_on,json=activeOn,proto3,oneof"`
}

func (*AlertDefProperties_ActiveOn) isAlertDefProperties_Schedule() {}

type isAlertDefProperties_TypeDefinition interface {
	isAlertDefProperties_TypeDefinition()
}

type AlertDefProperties_LogsImmediate struct {
	LogsImmediate *LogsImmediateType `protobuf:"bytes,101,opt,name=logs_immediate,json=logsImmediate,proto3,oneof"`
}

type AlertDefProperties_TracingImmediate struct {
	TracingImmediate *TracingImmediateType `protobuf:"bytes,102,opt,name=tracing_immediate,json=tracingImmediate,proto3,oneof"`
}

type AlertDefProperties_LogsThreshold struct {
	LogsThreshold *LogsThresholdType `protobuf:"bytes,103,opt,name=logs_threshold,json=logsThreshold,proto3,oneof"`
}

type AlertDefProperties_LogsRatioThreshold struct {
	LogsRatioThreshold *LogsRatioThresholdType `protobuf:"bytes,104,opt,name=logs_ratio_threshold,json=logsRatioThreshold,proto3,oneof"`
}

type AlertDefProperties_LogsTimeRelativeThreshold struct {
	LogsTimeRelativeThreshold *LogsTimeRelativeThresholdType `protobuf:"bytes,105,opt,name=logs_time_relative_threshold,json=logsTimeRelativeThreshold,proto3,oneof"`
}

type AlertDefProperties_MetricThreshold struct {
	MetricThreshold *MetricThresholdType `protobuf:"bytes,106,opt,name=metric_threshold,json=metricThreshold,proto3,oneof"`
}

type AlertDefProperties_TracingThreshold struct {
	TracingThreshold *TracingThresholdType `protobuf:"bytes,107,opt,name=tracing_threshold,json=tracingThreshold,proto3,oneof"`
}

type AlertDefProperties_Flow struct {
	Flow *FlowType `protobuf:"bytes,108,opt,name=flow,proto3,oneof"`
}

type AlertDefProperties_LogsAnomaly struct {
	LogsAnomaly *LogsAnomalyType `protobuf:"bytes,109,opt,name=logs_anomaly,json=logsAnomaly,proto3,oneof"`
}

type AlertDefProperties_MetricAnomaly struct {
	MetricAnomaly *MetricAnomalyType `protobuf:"bytes,110,opt,name=metric_anomaly,json=metricAnomaly,proto3,oneof"`
}

type AlertDefProperties_LogsNewValue struct {
	LogsNewValue *LogsNewValueType `protobuf:"bytes,111,opt,name=logs_new_value,json=logsNewValue,proto3,oneof"`
}

type AlertDefProperties_LogsUniqueCount struct {
	LogsUniqueCount *LogsUniqueCountType `protobuf:"bytes,112,opt,name=logs_unique_count,json=logsUniqueCount,proto3,oneof"`
}

func (*AlertDefProperties_LogsImmediate) isAlertDefProperties_TypeDefinition() {}

func (*AlertDefProperties_TracingImmediate) isAlertDefProperties_TypeDefinition() {}

func (*AlertDefProperties_LogsThreshold) isAlertDefProperties_TypeDefinition() {}

func (*AlertDefProperties_LogsRatioThreshold) isAlertDefProperties_TypeDefinition() {}

func (*AlertDefProperties_LogsTimeRelativeThreshold) isAlertDefProperties_TypeDefinition() {}

func (*AlertDefProperties_MetricThreshold) isAlertDefProperties_TypeDefinition() {}

func (*AlertDefProperties_TracingThreshold) isAlertDefProperties_TypeDefinition() {}

func (*AlertDefProperties_Flow) isAlertDefProperties_TypeDefinition() {}

func (*AlertDefProperties_LogsAnomaly) isAlertDefProperties_TypeDefinition() {}

func (*AlertDefProperties_MetricAnomaly) isAlertDefProperties_TypeDefinition() {}

func (*AlertDefProperties_LogsNewValue) isAlertDefProperties_TypeDefinition() {}

func (*AlertDefProperties_LogsUniqueCount) isAlertDefProperties_TypeDefinition() {}

var File_com_coralogixapis_alerts_v3_alert_def_proto protoreflect.FileDescriptor

const file_com_coralogixapis_alerts_v3_alert_def_proto_rawDesc = "" +
	"\n" +
	"+com/coralogixapis/alerts/v3/alert_def.proto\x12\x1bcom.coralogixapis.alerts.v3\x1a>com/coralogixapis/alerts/v3/alert_def_notification_group.proto\x1a4com/coralogixapis/alerts/v3/alert_def_priority.proto\x1a6com/coralogixapis/alerts/v3/alert_def_scheduling.proto\x1a0com/coralogixapis/alerts/v3/alert_def_type.proto\x1aUcom/coralogixapis/alerts/v3/alert_def_type_definition/flow/flow_type_definition.proto\x1accom/coralogixapis/alerts/v3/alert_def_type_definition/metric/metric_threshold_type_definition.proto\x1aacom/coralogixapis/alerts/v3/alert_def_type_definition/metric/metric_anomaly_type_definition.proto\x1a_com/coralogixapis/alerts/v3/alert_def_type_definition/logs/logs_new_value_type_definition.proto\x1aecom/coralogixapis/alerts/v3/alert_def_type_definition/logs/logs_ratio_threshold_type_definition.proto\x1a_com/coralogixapis/alerts/v3/alert_def_type_definition/logs/logs_immediate_type_definition.proto\x1a_com/coralogixapis/alerts/v3/alert_def_type_definition/logs/logs_threshold_type_definition.proto\x1a]com/coralogixapis/alerts/v3/alert_def_type_definition/logs/logs_anomaly_type_definition.proto\x1amcom/coralogixapis/alerts/v3/alert_def_type_definition/logs/logs_time_relative_threshold_type_definition.proto\x1aecom/coralogixapis/alerts/v3/alert_def_type_definition/tracing/tracing_immediate_type_definition.proto\x1aecom/coralogixapis/alerts/v3/alert_def_type_definition/tracing/tracing_threshold_type_definition.proto\x1abcom/coralogixapis/alerts/v3/alert_def_type_definition/logs/logs_unique_count_type_definition.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/wrappers.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\xe3\x04\n" +
	"\bAlertDef\x12a\n" +
	"\x14alert_def_properties\x18\x01 \x01(\v2/.com.coralogixapis.alerts.v3.AlertDefPropertiesR\x12alertDefProperties\x124\n" +
	"\x02id\x18\x02 \x01(\v2\x1c.google.protobuf.StringValueB\x06\x92A\x03J\x011R\x02id\x12F\n" +
	"\x10alert_version_id\x18\x14 \x01(\v2\x1c.google.protobuf.StringValueR\x0ealertVersionId\x12=\n" +
	"\fcreated_time\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\vcreatedTime\x12=\n" +
	"\fupdated_time\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\vupdatedTime:\xf7\x01\x92A\xf3\x01\n" +
	"u*\x10Alert definition22This data structure represents an alert definition\xd2\x01\x14alert_def_properties\xd2\x01\x02id\xd2\x01\x10alert_version_id*z\n" +
	"/Find out more about alerts in our documentation\x12Ghttps://coralogix.com/docs/user-guides/alerting/introduction-to-alerts/\"\xda\x13\n" +
	"\x12AlertDefProperties\x12A\n" +
	"\x04name\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueB\x0f\x92A\fJ\n" +
	"\"My Alert\"R\x04name\x12X\n" +
	"\vdescription\x18\x02 \x01(\v2\x1c.google.protobuf.StringValueB\x18\x92A\x15J\x13\"Alert description\"R\vdescription\x12?\n" +
	"\aenabled\x18\x03 \x01(\v2\x1a.google.protobuf.BoolValueB\t\x92A\x06J\x04trueR\aenabled\x12I\n" +
	"\bpriority\x18\x04 \x01(\x0e2-.com.coralogixapis.alerts.v3.AlertDefPriorityR\bpriority\x12L\n" +
	"\tactive_on\x18\x05 \x01(\v2-.com.coralogixapis.alerts.v3.ActivityScheduleH\x00R\bactiveOn\x12=\n" +
	"\x04type\x18\x06 \x01(\x0e2).com.coralogixapis.alerts.v3.AlertDefTypeR\x04type\x12W\n" +
	"\x0elogs_immediate\x18e \x01(\v2..com.coralogixapis.alerts.v3.LogsImmediateTypeH\x01R\rlogsImmediate\x12`\n" +
	"\x11tracing_immediate\x18f \x01(\v21.com.coralogixapis.alerts.v3.TracingImmediateTypeH\x01R\x10tracingImmediate\x12W\n" +
	"\x0elogs_threshold\x18g \x01(\v2..com.coralogixapis.alerts.v3.LogsThresholdTypeH\x01R\rlogsThreshold\x12g\n" +
	"\x14logs_ratio_threshold\x18h \x01(\v23.com.coralogixapis.alerts.v3.LogsRatioThresholdTypeH\x01R\x12logsRatioThreshold\x12}\n" +
	"\x1clogs_time_relative_threshold\x18i \x01(\v2:.com.coralogixapis.alerts.v3.LogsTimeRelativeThresholdTypeH\x01R\x19logsTimeRelativeThreshold\x12]\n" +
	"\x10metric_threshold\x18j \x01(\v20.com.coralogixapis.alerts.v3.MetricThresholdTypeH\x01R\x0fmetricThreshold\x12`\n" +
	"\x11tracing_threshold\x18k \x01(\v21.com.coralogixapis.alerts.v3.TracingThresholdTypeH\x01R\x10tracingThreshold\x12;\n" +
	"\x04flow\x18l \x01(\v2%.com.coralogixapis.alerts.v3.FlowTypeH\x01R\x04flow\x12Q\n" +
	"\flogs_anomaly\x18m \x01(\v2,.com.coralogixapis.alerts.v3.LogsAnomalyTypeH\x01R\vlogsAnomaly\x12W\n" +
	"\x0emetric_anomaly\x18n \x01(\v2..com.coralogixapis.alerts.v3.MetricAnomalyTypeH\x01R\rmetricAnomaly\x12U\n" +
	"\x0elogs_new_value\x18o \x01(\v2-.com.coralogixapis.alerts.v3.LogsNewValueTypeH\x01R\flogsNewValue\x12^\n" +
	"\x11logs_unique_count\x18p \x01(\v20.com.coralogixapis.alerts.v3.LogsUniqueCountTypeH\x01R\x0flogsUniqueCount\x12@\n" +
	"\rgroup_by_keys\x18\a \x03(\v2\x1c.google.protobuf.StringValueR\vgroupByKeys\x12d\n" +
	"\x12incidents_settings\x18\b \x01(\v25.com.coralogixapis.alerts.v3.AlertDefIncidentSettingsR\x11incidentsSettings\x12e\n" +
	"\x12notification_group\x18\t \x01(\v26.com.coralogixapis.alerts.v3.AlertDefNotificationGroupR\x11notificationGroup\x12w\n" +
	"\x19notification_group_excess\x18\xd2\x01 \x03(\v26.com.coralogixapis.alerts.v3.AlertDefNotificationGroupB\x02\x18\x01R\x17notificationGroupExcess\x12f\n" +
	"\rentity_labels\x18\n" +
	" \x03(\v2A.com.coralogixapis.alerts.v3.AlertDefProperties.EntityLabelsEntryR\fentityLabels\x12=\n" +
	"\fphantom_mode\x18\v \x01(\v2\x1a.google.protobuf.BoolValueR\vphantomMode\x124\n" +
	"\adeleted\x18\f \x01(\v2\x1a.google.protobuf.BoolValueR\adeleted\x1a?\n" +
	"\x11EntityLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01:\xeb\x01\x92A\xe7\x01\n" +
	"\xe4\x01*\x1bAlert definition properties23User-configurable properties of an alert definition\xd2\x01\x04name\xd2\x01\vdescription\xd2\x01\aenabled\xd2\x01\bpriority\xd2\x01\bschedule\xd2\x01\x04type\xd2\x01\x0ftype_definition\xd2\x01\rgroup_by_keys\xd2\x01\x12notification_group\xd2\x01\rentity_labels\xd2\x01\fphantom_modeB\n" +
	"\n" +
	"\bscheduleB\x11\n" +
	"\x0ftype_definitionb\x06proto3"

var (
	file_com_coralogixapis_alerts_v3_alert_def_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_alert_def_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v3_alert_def_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_alert_def_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_alert_def_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerts_v3_alert_def_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_alert_def_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogixapis_alerts_v3_alert_def_proto_goTypes = []any{
	(*AlertDef)(nil),                      // 0: com.coralogixapis.alerts.v3.AlertDef
	(*AlertDefProperties)(nil),            // 1: com.coralogixapis.alerts.v3.AlertDefProperties
	nil,                                   // 2: com.coralogixapis.alerts.v3.AlertDefProperties.EntityLabelsEntry
	(*wrapperspb.StringValue)(nil),        // 3: google.protobuf.StringValue
	(*timestamppb.Timestamp)(nil),         // 4: google.protobuf.Timestamp
	(*wrapperspb.BoolValue)(nil),          // 5: google.protobuf.BoolValue
	(AlertDefPriority)(0),                 // 6: com.coralogixapis.alerts.v3.AlertDefPriority
	(*ActivitySchedule)(nil),              // 7: com.coralogixapis.alerts.v3.ActivitySchedule
	(AlertDefType)(0),                     // 8: com.coralogixapis.alerts.v3.AlertDefType
	(*LogsImmediateType)(nil),             // 9: com.coralogixapis.alerts.v3.LogsImmediateType
	(*TracingImmediateType)(nil),          // 10: com.coralogixapis.alerts.v3.TracingImmediateType
	(*LogsThresholdType)(nil),             // 11: com.coralogixapis.alerts.v3.LogsThresholdType
	(*LogsRatioThresholdType)(nil),        // 12: com.coralogixapis.alerts.v3.LogsRatioThresholdType
	(*LogsTimeRelativeThresholdType)(nil), // 13: com.coralogixapis.alerts.v3.LogsTimeRelativeThresholdType
	(*MetricThresholdType)(nil),           // 14: com.coralogixapis.alerts.v3.MetricThresholdType
	(*TracingThresholdType)(nil),          // 15: com.coralogixapis.alerts.v3.TracingThresholdType
	(*FlowType)(nil),                      // 16: com.coralogixapis.alerts.v3.FlowType
	(*LogsAnomalyType)(nil),               // 17: com.coralogixapis.alerts.v3.LogsAnomalyType
	(*MetricAnomalyType)(nil),             // 18: com.coralogixapis.alerts.v3.MetricAnomalyType
	(*LogsNewValueType)(nil),              // 19: com.coralogixapis.alerts.v3.LogsNewValueType
	(*LogsUniqueCountType)(nil),           // 20: com.coralogixapis.alerts.v3.LogsUniqueCountType
	(*AlertDefIncidentSettings)(nil),      // 21: com.coralogixapis.alerts.v3.AlertDefIncidentSettings
	(*AlertDefNotificationGroup)(nil),     // 22: com.coralogixapis.alerts.v3.AlertDefNotificationGroup
}
var file_com_coralogixapis_alerts_v3_alert_def_proto_depIdxs = []int32{
	1,  // 0: com.coralogixapis.alerts.v3.AlertDef.alert_def_properties:type_name -> com.coralogixapis.alerts.v3.AlertDefProperties
	3,  // 1: com.coralogixapis.alerts.v3.AlertDef.id:type_name -> google.protobuf.StringValue
	3,  // 2: com.coralogixapis.alerts.v3.AlertDef.alert_version_id:type_name -> google.protobuf.StringValue
	4,  // 3: com.coralogixapis.alerts.v3.AlertDef.created_time:type_name -> google.protobuf.Timestamp
	4,  // 4: com.coralogixapis.alerts.v3.AlertDef.updated_time:type_name -> google.protobuf.Timestamp
	3,  // 5: com.coralogixapis.alerts.v3.AlertDefProperties.name:type_name -> google.protobuf.StringValue
	3,  // 6: com.coralogixapis.alerts.v3.AlertDefProperties.description:type_name -> google.protobuf.StringValue
	5,  // 7: com.coralogixapis.alerts.v3.AlertDefProperties.enabled:type_name -> google.protobuf.BoolValue
	6,  // 8: com.coralogixapis.alerts.v3.AlertDefProperties.priority:type_name -> com.coralogixapis.alerts.v3.AlertDefPriority
	7,  // 9: com.coralogixapis.alerts.v3.AlertDefProperties.active_on:type_name -> com.coralogixapis.alerts.v3.ActivitySchedule
	8,  // 10: com.coralogixapis.alerts.v3.AlertDefProperties.type:type_name -> com.coralogixapis.alerts.v3.AlertDefType
	9,  // 11: com.coralogixapis.alerts.v3.AlertDefProperties.logs_immediate:type_name -> com.coralogixapis.alerts.v3.LogsImmediateType
	10, // 12: com.coralogixapis.alerts.v3.AlertDefProperties.tracing_immediate:type_name -> com.coralogixapis.alerts.v3.TracingImmediateType
	11, // 13: com.coralogixapis.alerts.v3.AlertDefProperties.logs_threshold:type_name -> com.coralogixapis.alerts.v3.LogsThresholdType
	12, // 14: com.coralogixapis.alerts.v3.AlertDefProperties.logs_ratio_threshold:type_name -> com.coralogixapis.alerts.v3.LogsRatioThresholdType
	13, // 15: com.coralogixapis.alerts.v3.AlertDefProperties.logs_time_relative_threshold:type_name -> com.coralogixapis.alerts.v3.LogsTimeRelativeThresholdType
	14, // 16: com.coralogixapis.alerts.v3.AlertDefProperties.metric_threshold:type_name -> com.coralogixapis.alerts.v3.MetricThresholdType
	15, // 17: com.coralogixapis.alerts.v3.AlertDefProperties.tracing_threshold:type_name -> com.coralogixapis.alerts.v3.TracingThresholdType
	16, // 18: com.coralogixapis.alerts.v3.AlertDefProperties.flow:type_name -> com.coralogixapis.alerts.v3.FlowType
	17, // 19: com.coralogixapis.alerts.v3.AlertDefProperties.logs_anomaly:type_name -> com.coralogixapis.alerts.v3.LogsAnomalyType
	18, // 20: com.coralogixapis.alerts.v3.AlertDefProperties.metric_anomaly:type_name -> com.coralogixapis.alerts.v3.MetricAnomalyType
	19, // 21: com.coralogixapis.alerts.v3.AlertDefProperties.logs_new_value:type_name -> com.coralogixapis.alerts.v3.LogsNewValueType
	20, // 22: com.coralogixapis.alerts.v3.AlertDefProperties.logs_unique_count:type_name -> com.coralogixapis.alerts.v3.LogsUniqueCountType
	3,  // 23: com.coralogixapis.alerts.v3.AlertDefProperties.group_by_keys:type_name -> google.protobuf.StringValue
	21, // 24: com.coralogixapis.alerts.v3.AlertDefProperties.incidents_settings:type_name -> com.coralogixapis.alerts.v3.AlertDefIncidentSettings
	22, // 25: com.coralogixapis.alerts.v3.AlertDefProperties.notification_group:type_name -> com.coralogixapis.alerts.v3.AlertDefNotificationGroup
	22, // 26: com.coralogixapis.alerts.v3.AlertDefProperties.notification_group_excess:type_name -> com.coralogixapis.alerts.v3.AlertDefNotificationGroup
	2,  // 27: com.coralogixapis.alerts.v3.AlertDefProperties.entity_labels:type_name -> com.coralogixapis.alerts.v3.AlertDefProperties.EntityLabelsEntry
	5,  // 28: com.coralogixapis.alerts.v3.AlertDefProperties.phantom_mode:type_name -> google.protobuf.BoolValue
	5,  // 29: com.coralogixapis.alerts.v3.AlertDefProperties.deleted:type_name -> google.protobuf.BoolValue
	30, // [30:30] is the sub-list for method output_type
	30, // [30:30] is the sub-list for method input_type
	30, // [30:30] is the sub-list for extension type_name
	30, // [30:30] is the sub-list for extension extendee
	0,  // [0:30] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_alerts_v3_alert_def_proto_init() }
func file_com_coralogixapis_alerts_v3_alert_def_proto_init() {
	if File_com_coralogixapis_alerts_v3_alert_def_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_alert_def_notification_group_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_priority_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_scheduling_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_flow_flow_type_definition_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_threshold_type_definition_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_metric_anomaly_type_definition_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_new_value_type_definition_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_ratio_threshold_type_definition_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_immediate_type_definition_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_threshold_type_definition_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_anomaly_type_definition_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_time_relative_threshold_type_definition_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_tracing_immediate_type_definition_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_tracing_tracing_threshold_type_definition_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_logs_logs_unique_count_type_definition_proto_init()
	file_com_coralogixapis_alerts_v3_alert_def_proto_msgTypes[1].OneofWrappers = []any{
		(*AlertDefProperties_ActiveOn)(nil),
		(*AlertDefProperties_LogsImmediate)(nil),
		(*AlertDefProperties_TracingImmediate)(nil),
		(*AlertDefProperties_LogsThreshold)(nil),
		(*AlertDefProperties_LogsRatioThreshold)(nil),
		(*AlertDefProperties_LogsTimeRelativeThreshold)(nil),
		(*AlertDefProperties_MetricThreshold)(nil),
		(*AlertDefProperties_TracingThreshold)(nil),
		(*AlertDefProperties_Flow)(nil),
		(*AlertDefProperties_LogsAnomaly)(nil),
		(*AlertDefProperties_MetricAnomaly)(nil),
		(*AlertDefProperties_LogsNewValue)(nil),
		(*AlertDefProperties_LogsUniqueCount)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_alert_def_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_alert_def_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_alert_def_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_alert_def_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_alert_def_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_alert_def_proto = out.File
	file_com_coralogixapis_alerts_v3_alert_def_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_alert_def_proto_depIdxs = nil
}
