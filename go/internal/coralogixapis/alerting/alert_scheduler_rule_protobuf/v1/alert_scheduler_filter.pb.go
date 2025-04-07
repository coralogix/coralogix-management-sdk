// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.26.1
// source: com/coralogixapis/alerting/alert_scheduler_rule_protobuf/v1/alert_scheduler_filter.proto

package v1

import (
	v1 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/alerting/meta_labels_protobuf/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Filter struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	WhatExpression string                 `protobuf:"bytes,1,opt,name=what_expression,json=whatExpression,proto3" json:"what_expression,omitempty"` // Filter what_expression: dataprime expression that filter the alerts group by values.
	// Types that are valid to be assigned to WhichAlerts:
	//
	//	*Filter_AlertMetaLabels
	//	*Filter_AlertUniqueIds
	WhichAlerts   isFilter_WhichAlerts `protobuf_oneof:"which_alerts"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Filter) Reset() {
	*x = Filter{}
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_rawDescGZIP(), []int{0}
}

func (x *Filter) GetWhatExpression() string {
	if x != nil {
		return x.WhatExpression
	}
	return ""
}

func (x *Filter) GetWhichAlerts() isFilter_WhichAlerts {
	if x != nil {
		return x.WhichAlerts
	}
	return nil
}

func (x *Filter) GetAlertMetaLabels() *MetaLabels {
	if x != nil {
		if x, ok := x.WhichAlerts.(*Filter_AlertMetaLabels); ok {
			return x.AlertMetaLabels
		}
	}
	return nil
}

func (x *Filter) GetAlertUniqueIds() *AlertUniqueIds {
	if x != nil {
		if x, ok := x.WhichAlerts.(*Filter_AlertUniqueIds); ok {
			return x.AlertUniqueIds
		}
	}
	return nil
}

type isFilter_WhichAlerts interface {
	isFilter_WhichAlerts()
}

type Filter_AlertMetaLabels struct {
	AlertMetaLabels *MetaLabels `protobuf:"bytes,2,opt,name=alert_meta_labels,json=alertMetaLabels,proto3,oneof"` // Filter alert_meta_labels: filter alerts by meta labels tagging.
}

type Filter_AlertUniqueIds struct {
	AlertUniqueIds *AlertUniqueIds `protobuf:"bytes,3,opt,name=alert_unique_ids,json=alertUniqueIds,proto3,oneof"` // Filter alert_unique_ids: filter specific alerts (when alert_unique_ids is empty meaning it wil filter all alerts).
}

func (*Filter_AlertMetaLabels) isFilter_WhichAlerts() {}

func (*Filter_AlertUniqueIds) isFilter_WhichAlerts() {}

type AlertUniqueIds struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         []string               `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AlertUniqueIds) Reset() {
	*x = AlertUniqueIds{}
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlertUniqueIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertUniqueIds) ProtoMessage() {}

func (x *AlertUniqueIds) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertUniqueIds.ProtoReflect.Descriptor instead.
func (*AlertUniqueIds) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_rawDescGZIP(), []int{1}
}

func (x *AlertUniqueIds) GetValue() []string {
	if x != nil {
		return x.Value
	}
	return nil
}

type MetaLabels struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         []*v1.MetaLabel        `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MetaLabels) Reset() {
	*x = MetaLabels{}
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetaLabels) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaLabels) ProtoMessage() {}

func (x *MetaLabels) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaLabels.ProtoReflect.Descriptor instead.
func (*MetaLabels) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_rawDescGZIP(), []int{2}
}

func (x *MetaLabels) GetValue() []*v1.MetaLabel {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto protoreflect.FileDescriptor

const file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_rawDesc = "" +
	"\n" +
	"Xcom/coralogixapis/alerting/alert_scheduler_rule_protobuf/v1/alert_scheduler_filter.proto\x12;com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1\x1aDcom/coralogixapis/alerting/meta_labels_protobuf/v1/meta_labels.proto\"\xb1\x02\n" +
	"\x06Filter\x12'\n" +
	"\x0fwhat_expression\x18\x01 \x01(\tR\x0ewhatExpression\x12u\n" +
	"\x11alert_meta_labels\x18\x02 \x01(\v2G.com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.MetaLabelsH\x00R\x0falertMetaLabels\x12w\n" +
	"\x10alert_unique_ids\x18\x03 \x01(\v2K.com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertUniqueIdsH\x00R\x0ealertUniqueIdsB\x0e\n" +
	"\fwhich_alerts\"&\n" +
	"\x0eAlertUniqueIds\x12\x14\n" +
	"\x05value\x18\x01 \x03(\tR\x05value\"a\n" +
	"\n" +
	"MetaLabels\x12S\n" +
	"\x05value\x18\x01 \x03(\v2=.com.coralogixapis.alerting.meta_labels_protobuf.v1.MetaLabelR\x05valueb\x06proto3"

var (
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_rawDescData []byte
)

func file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_rawDesc), len(file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_rawDescData
}

var file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_goTypes = []any{
	(*Filter)(nil),         // 0: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Filter
	(*AlertUniqueIds)(nil), // 1: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertUniqueIds
	(*MetaLabels)(nil),     // 2: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.MetaLabels
	(*v1.MetaLabel)(nil),   // 3: com.coralogixapis.alerting.meta_labels_protobuf.v1.MetaLabel
}
var file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_depIdxs = []int32{
	2, // 0: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Filter.alert_meta_labels:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.MetaLabels
	1, // 1: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Filter.alert_unique_ids:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertUniqueIds
	3, // 2: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.MetaLabels.value:type_name -> com.coralogixapis.alerting.meta_labels_protobuf.v1.MetaLabel
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_init()
}
func file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_init() {
	if File_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto != nil {
		return
	}
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_msgTypes[0].OneofWrappers = []any{
		(*Filter_AlertMetaLabels)(nil),
		(*Filter_AlertUniqueIds)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_rawDesc), len(file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto = out.File
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_goTypes = nil
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_depIdxs = nil
}
