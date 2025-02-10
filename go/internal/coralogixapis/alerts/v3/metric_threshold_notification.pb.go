// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: com/coralogixapis/alerts/v3/notification/metric_threshold_notification.proto

package v3

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

type MetricThresholdNotification struct {
	state                 protoimpl.MessageState       `protogen:"open.v1"`
	PctOverThreshold      *wrapperspb.DoubleValue      `protobuf:"bytes,1,opt,name=pct_over_threshold,json=pctOverThreshold,proto3" json:"pct_over_threshold,omitempty"`
	AvgValueOverThreshold *wrapperspb.DoubleValue      `protobuf:"bytes,2,opt,name=avg_value_over_threshold,json=avgValueOverThreshold,proto3" json:"avg_value_over_threshold,omitempty"`
	MinValueOverThreshold *wrapperspb.DoubleValue      `protobuf:"bytes,3,opt,name=min_value_over_threshold,json=minValueOverThreshold,proto3" json:"min_value_over_threshold,omitempty"`
	MaxValueOverThreshold *wrapperspb.DoubleValue      `protobuf:"bytes,4,opt,name=max_value_over_threshold,json=maxValueOverThreshold,proto3" json:"max_value_over_threshold,omitempty"`
	IsUndetectedValue     *wrapperspb.BoolValue        `protobuf:"bytes,5,opt,name=is_undetected_value,json=isUndetectedValue,proto3" json:"is_undetected_value,omitempty"`
	FromTimestamp         *timestamppb.Timestamp       `protobuf:"bytes,6,opt,name=from_timestamp,json=fromTimestamp,proto3" json:"from_timestamp,omitempty"`
	ToTimestamp           *timestamppb.Timestamp       `protobuf:"bytes,7,opt,name=to_timestamp,json=toTimestamp,proto3" json:"to_timestamp,omitempty"`
	ConditionType         MetricThresholdConditionType `protobuf:"varint,8,opt,name=condition_type,json=conditionType,proto3,enum=com.coralogixapis.alerts.v3.MetricThresholdConditionType" json:"condition_type,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *MetricThresholdNotification) Reset() {
	*x = MetricThresholdNotification{}
	mi := &file_com_coralogixapis_alerts_v3_notification_metric_threshold_notification_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetricThresholdNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricThresholdNotification) ProtoMessage() {}

func (x *MetricThresholdNotification) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_notification_metric_threshold_notification_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricThresholdNotification.ProtoReflect.Descriptor instead.
func (*MetricThresholdNotification) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_notification_metric_threshold_notification_proto_rawDescGZIP(), []int{0}
}

func (x *MetricThresholdNotification) GetPctOverThreshold() *wrapperspb.DoubleValue {
	if x != nil {
		return x.PctOverThreshold
	}
	return nil
}

func (x *MetricThresholdNotification) GetAvgValueOverThreshold() *wrapperspb.DoubleValue {
	if x != nil {
		return x.AvgValueOverThreshold
	}
	return nil
}

func (x *MetricThresholdNotification) GetMinValueOverThreshold() *wrapperspb.DoubleValue {
	if x != nil {
		return x.MinValueOverThreshold
	}
	return nil
}

func (x *MetricThresholdNotification) GetMaxValueOverThreshold() *wrapperspb.DoubleValue {
	if x != nil {
		return x.MaxValueOverThreshold
	}
	return nil
}

func (x *MetricThresholdNotification) GetIsUndetectedValue() *wrapperspb.BoolValue {
	if x != nil {
		return x.IsUndetectedValue
	}
	return nil
}

func (x *MetricThresholdNotification) GetFromTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.FromTimestamp
	}
	return nil
}

func (x *MetricThresholdNotification) GetToTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.ToTimestamp
	}
	return nil
}

func (x *MetricThresholdNotification) GetConditionType() MetricThresholdConditionType {
	if x != nil {
		return x.ConditionType
	}
	return MetricThresholdConditionType_METRIC_THRESHOLD_CONDITION_TYPE_MORE_THAN_OR_UNSPECIFIED
}

var File_com_coralogixapis_alerts_v3_notification_metric_threshold_notification_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_notification_metric_threshold_notification_proto_rawDesc = []byte{
	0x0a, 0x4c, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x5f, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x6c, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f,
	0x64, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2f, 0x74, 0x68, 0x72, 0x65, 0x73,
	0x68, 0x6f, 0x6c, 0x64, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x74, 0x68, 0x72, 0x65,
	0x73, 0x68, 0x6f, 0x6c, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x05, 0x0a, 0x1b, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x12, 0x70, 0x63,
	0x74, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x70, 0x63, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x54, 0x68, 0x72,
	0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x55, 0x0a, 0x18, 0x61, 0x76, 0x67, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f,
	0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x15, 0x61, 0x76, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x4f, 0x76, 0x65, 0x72, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x55, 0x0a,
	0x18, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x5f,
	0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x15, 0x6d,
	0x69, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x54, 0x68, 0x72, 0x65, 0x73,
	0x68, 0x6f, 0x6c, 0x64, 0x12, 0x55, 0x0a, 0x18, 0x6d, 0x61, 0x78, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x15, 0x6d, 0x61, 0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x76,
	0x65, 0x72, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x4a, 0x0a, 0x13, 0x69,
	0x73, 0x5f, 0x75, 0x6e, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x11, 0x69, 0x73, 0x55, 0x6e, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x66, 0x72, 0x6f, 0x6d, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x66, 0x72, 0x6f,
	0x6d, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3d, 0x0a, 0x0c, 0x74, 0x6f,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x74, 0x6f,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x60, 0x0a, 0x0e, 0x63, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_notification_metric_threshold_notification_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_notification_metric_threshold_notification_proto_rawDescData = file_com_coralogixapis_alerts_v3_notification_metric_threshold_notification_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_notification_metric_threshold_notification_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_notification_metric_threshold_notification_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_notification_metric_threshold_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_notification_metric_threshold_notification_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_notification_metric_threshold_notification_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_notification_metric_threshold_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_alerts_v3_notification_metric_threshold_notification_proto_goTypes = []any{
	(*MetricThresholdNotification)(nil), // 0: com.coralogixapis.alerts.v3.MetricThresholdNotification
	(*wrapperspb.DoubleValue)(nil),      // 1: google.protobuf.DoubleValue
	(*wrapperspb.BoolValue)(nil),        // 2: google.protobuf.BoolValue
	(*timestamppb.Timestamp)(nil),       // 3: google.protobuf.Timestamp
	(MetricThresholdConditionType)(0),   // 4: com.coralogixapis.alerts.v3.MetricThresholdConditionType
}
var file_com_coralogixapis_alerts_v3_notification_metric_threshold_notification_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.alerts.v3.MetricThresholdNotification.pct_over_threshold:type_name -> google.protobuf.DoubleValue
	1, // 1: com.coralogixapis.alerts.v3.MetricThresholdNotification.avg_value_over_threshold:type_name -> google.protobuf.DoubleValue
	1, // 2: com.coralogixapis.alerts.v3.MetricThresholdNotification.min_value_over_threshold:type_name -> google.protobuf.DoubleValue
	1, // 3: com.coralogixapis.alerts.v3.MetricThresholdNotification.max_value_over_threshold:type_name -> google.protobuf.DoubleValue
	2, // 4: com.coralogixapis.alerts.v3.MetricThresholdNotification.is_undetected_value:type_name -> google.protobuf.BoolValue
	3, // 5: com.coralogixapis.alerts.v3.MetricThresholdNotification.from_timestamp:type_name -> google.protobuf.Timestamp
	3, // 6: com.coralogixapis.alerts.v3.MetricThresholdNotification.to_timestamp:type_name -> google.protobuf.Timestamp
	4, // 7: com.coralogixapis.alerts.v3.MetricThresholdNotification.condition_type:type_name -> com.coralogixapis.alerts.v3.MetricThresholdConditionType
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_alerts_v3_notification_metric_threshold_notification_proto_init() }
func file_com_coralogixapis_alerts_v3_notification_metric_threshold_notification_proto_init() {
	if File_com_coralogixapis_alerts_v3_notification_metric_threshold_notification_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_alert_def_type_definition_metric_threshold_metric_threshold_condition_type_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_notification_metric_threshold_notification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_notification_metric_threshold_notification_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_notification_metric_threshold_notification_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_notification_metric_threshold_notification_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_notification_metric_threshold_notification_proto = out.File
	file_com_coralogixapis_alerts_v3_notification_metric_threshold_notification_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_notification_metric_threshold_notification_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_notification_metric_threshold_notification_proto_depIdxs = nil
}
