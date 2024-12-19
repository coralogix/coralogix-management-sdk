// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.26.1
// source: com/coralogix/catalog/v1/common.proto

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

type TimeRange struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Start         *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End           *wrapperspb.Int64Value `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TimeRange) Reset() {
	*x = TimeRange{}
	mi := &file_com_coralogix_catalog_v1_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimeRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeRange) ProtoMessage() {}

func (x *TimeRange) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_catalog_v1_common_proto_msgTypes[0]
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
	return file_com_coralogix_catalog_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *TimeRange) GetStart() *wrapperspb.Int64Value {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *TimeRange) GetEnd() *wrapperspb.Int64Value {
	if x != nil {
		return x.End
	}
	return nil
}

type ApmFilter struct {
	state           protoimpl.MessageState    `protogen:"open.v1"`
	LabelName       *wrapperspb.StringValue   `protobuf:"bytes,1,opt,name=label_name,json=labelName,proto3" json:"label_name,omitempty"`
	LabelKey        *wrapperspb.StringValue   `protobuf:"bytes,2,opt,name=label_key,json=labelKey,proto3" json:"label_key,omitempty"`
	LabelValues     []*wrapperspb.StringValue `protobuf:"bytes,3,rep,name=label_values,json=labelValues,proto3" json:"label_values,omitempty"`
	CreatedAt       *wrapperspb.Int64Value    `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       *wrapperspb.Int64Value    `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	MetricLabelName *wrapperspb.StringValue   `protobuf:"bytes,6,opt,name=metric_label_name,json=metricLabelName,proto3" json:"metric_label_name,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ApmFilter) Reset() {
	*x = ApmFilter{}
	mi := &file_com_coralogix_catalog_v1_common_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApmFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApmFilter) ProtoMessage() {}

func (x *ApmFilter) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_catalog_v1_common_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApmFilter.ProtoReflect.Descriptor instead.
func (*ApmFilter) Descriptor() ([]byte, []int) {
	return file_com_coralogix_catalog_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *ApmFilter) GetLabelName() *wrapperspb.StringValue {
	if x != nil {
		return x.LabelName
	}
	return nil
}

func (x *ApmFilter) GetLabelKey() *wrapperspb.StringValue {
	if x != nil {
		return x.LabelKey
	}
	return nil
}

func (x *ApmFilter) GetLabelValues() []*wrapperspb.StringValue {
	if x != nil {
		return x.LabelValues
	}
	return nil
}

func (x *ApmFilter) GetCreatedAt() *wrapperspb.Int64Value {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ApmFilter) GetUpdatedAt() *wrapperspb.Int64Value {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *ApmFilter) GetMetricLabelName() *wrapperspb.StringValue {
	if x != nil {
		return x.MetricLabelName
	}
	return nil
}

var File_com_coralogix_catalog_v1_common_proto protoreflect.FileDescriptor

var file_com_coralogix_catalog_v1_common_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76,
	0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x6d, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x31,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x2d, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x65, 0x6e, 0x64,
	0x22, 0x86, 0x03, 0x0a, 0x09, 0x41, 0x70, 0x6d, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x3b,
	0x0a, 0x0a, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x09, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x4b, 0x65, 0x79, 0x12, 0x3f, 0x0a, 0x0c, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x3a, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x48, 0x0a, 0x11, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_com_coralogix_catalog_v1_common_proto_rawDescOnce sync.Once
	file_com_coralogix_catalog_v1_common_proto_rawDescData = file_com_coralogix_catalog_v1_common_proto_rawDesc
)

func file_com_coralogix_catalog_v1_common_proto_rawDescGZIP() []byte {
	file_com_coralogix_catalog_v1_common_proto_rawDescOnce.Do(func() {
		file_com_coralogix_catalog_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_catalog_v1_common_proto_rawDescData)
	})
	return file_com_coralogix_catalog_v1_common_proto_rawDescData
}

var file_com_coralogix_catalog_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogix_catalog_v1_common_proto_goTypes = []any{
	(*TimeRange)(nil),              // 0: com.coralogix.catalog.v1.TimeRange
	(*ApmFilter)(nil),              // 1: com.coralogix.catalog.v1.ApmFilter
	(*wrapperspb.Int64Value)(nil),  // 2: google.protobuf.Int64Value
	(*wrapperspb.StringValue)(nil), // 3: google.protobuf.StringValue
}
var file_com_coralogix_catalog_v1_common_proto_depIdxs = []int32{
	2, // 0: com.coralogix.catalog.v1.TimeRange.start:type_name -> google.protobuf.Int64Value
	2, // 1: com.coralogix.catalog.v1.TimeRange.end:type_name -> google.protobuf.Int64Value
	3, // 2: com.coralogix.catalog.v1.ApmFilter.label_name:type_name -> google.protobuf.StringValue
	3, // 3: com.coralogix.catalog.v1.ApmFilter.label_key:type_name -> google.protobuf.StringValue
	3, // 4: com.coralogix.catalog.v1.ApmFilter.label_values:type_name -> google.protobuf.StringValue
	2, // 5: com.coralogix.catalog.v1.ApmFilter.created_at:type_name -> google.protobuf.Int64Value
	2, // 6: com.coralogix.catalog.v1.ApmFilter.updated_at:type_name -> google.protobuf.Int64Value
	3, // 7: com.coralogix.catalog.v1.ApmFilter.metric_label_name:type_name -> google.protobuf.StringValue
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_com_coralogix_catalog_v1_common_proto_init() }
func file_com_coralogix_catalog_v1_common_proto_init() {
	if File_com_coralogix_catalog_v1_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_catalog_v1_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_catalog_v1_common_proto_goTypes,
		DependencyIndexes: file_com_coralogix_catalog_v1_common_proto_depIdxs,
		MessageInfos:      file_com_coralogix_catalog_v1_common_proto_msgTypes,
	}.Build()
	File_com_coralogix_catalog_v1_common_proto = out.File
	file_com_coralogix_catalog_v1_common_proto_rawDesc = nil
	file_com_coralogix_catalog_v1_common_proto_goTypes = nil
	file_com_coralogix_catalog_v1_common_proto_depIdxs = nil
}
