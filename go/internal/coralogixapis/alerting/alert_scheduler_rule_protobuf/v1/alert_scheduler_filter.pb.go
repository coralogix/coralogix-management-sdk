// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: com/coralogixapis/alerting/alert_scheduler_rule_protobuf/v1/alert_scheduler_filter.proto

package v1

import (
	v1 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/alerting/meta_labels_protobuf/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WhatExpression string `protobuf:"bytes,1,opt,name=what_expression,json=whatExpression,proto3" json:"what_expression,omitempty"` // Filter what_expression: dataprime expression that filter the alerts group by values.
	// Types that are assignable to WhichAlerts:
	//
	//	*Filter_AlertMetaLabels
	//	*Filter_AlertUniqueIds
	WhichAlerts isFilter_WhichAlerts `protobuf_oneof:"which_alerts"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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

func (m *Filter) GetWhichAlerts() isFilter_WhichAlerts {
	if m != nil {
		return m.WhichAlerts
	}
	return nil
}

func (x *Filter) GetAlertMetaLabels() *MetaLabels {
	if x, ok := x.GetWhichAlerts().(*Filter_AlertMetaLabels); ok {
		return x.AlertMetaLabels
	}
	return nil
}

func (x *Filter) GetAlertUniqueIds() *AlertUniqueIds {
	if x, ok := x.GetWhichAlerts().(*Filter_AlertUniqueIds); ok {
		return x.AlertUniqueIds
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []string `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *AlertUniqueIds) Reset() {
	*x = AlertUniqueIds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertUniqueIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertUniqueIds) ProtoMessage() {}

func (x *AlertUniqueIds) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []*v1.MetaLabel `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *MetaLabels) Reset() {
	*x = MetaLabels{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaLabels) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaLabels) ProtoMessage() {}

func (x *MetaLabels) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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

var file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_rawDesc = []byte{
	0x0a, 0x58, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3b, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x1a, 0x44, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61,
	0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x02,
	0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x77, 0x68, 0x61, 0x74,
	0x5f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x77, 0x68, 0x61, 0x74, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x75, 0x0a, 0x11, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x5f,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x48, 0x00, 0x52, 0x0f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x77, 0x0a, 0x10, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x5f, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f,
	0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x73, 0x48,
	0x00, 0x52, 0x0e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64,
	0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x77, 0x68, 0x69, 0x63, 0x68, 0x5f, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x73, 0x22, 0x26, 0x0a, 0x0e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65,
	0x49, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x61, 0x0a, 0x0a, 0x4d, 0x65, 0x74,
	0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x53, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_rawDescData = file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_rawDesc
)

func file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_rawDescData)
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
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Filter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AlertUniqueIds); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*MetaLabels); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_msgTypes[0].OneofWrappers = []any{
		(*Filter_AlertMetaLabels)(nil),
		(*Filter_AlertUniqueIds)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_rawDesc,
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
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_rawDesc = nil
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_goTypes = nil
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_depIdxs = nil
}
