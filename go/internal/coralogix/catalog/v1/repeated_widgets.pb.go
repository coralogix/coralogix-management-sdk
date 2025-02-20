// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.26.1
// source: com/coralogixapis/service_catalog/v1/repeated_widgets.proto

package v1

import (
	v1 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/apm/widgets/v1"
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

type RepeatedLineChart struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	LineCharts    []*v1.LineChart         `protobuf:"bytes,1,rep,name=line_charts,json=lineCharts,proto3" json:"line_charts,omitempty"`
	TotalAmount   *wrapperspb.UInt64Value `protobuf:"bytes,2,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RepeatedLineChart) Reset() {
	*x = RepeatedLineChart{}
	mi := &file_com_coralogixapis_service_catalog_v1_repeated_widgets_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RepeatedLineChart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepeatedLineChart) ProtoMessage() {}

func (x *RepeatedLineChart) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_service_catalog_v1_repeated_widgets_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepeatedLineChart.ProtoReflect.Descriptor instead.
func (*RepeatedLineChart) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_service_catalog_v1_repeated_widgets_proto_rawDescGZIP(), []int{0}
}

func (x *RepeatedLineChart) GetLineCharts() []*v1.LineChart {
	if x != nil {
		return x.LineCharts
	}
	return nil
}

func (x *RepeatedLineChart) GetTotalAmount() *wrapperspb.UInt64Value {
	if x != nil {
		return x.TotalAmount
	}
	return nil
}

type RepeatedStats struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Stats         []*v1.Stat             `protobuf:"bytes,1,rep,name=stats,proto3" json:"stats,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RepeatedStats) Reset() {
	*x = RepeatedStats{}
	mi := &file_com_coralogixapis_service_catalog_v1_repeated_widgets_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RepeatedStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepeatedStats) ProtoMessage() {}

func (x *RepeatedStats) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_service_catalog_v1_repeated_widgets_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepeatedStats.ProtoReflect.Descriptor instead.
func (*RepeatedStats) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_service_catalog_v1_repeated_widgets_proto_rawDescGZIP(), []int{1}
}

func (x *RepeatedStats) GetStats() []*v1.Stat {
	if x != nil {
		return x.Stats
	}
	return nil
}

var File_com_coralogixapis_service_catalog_v1_repeated_widgets_proto protoreflect.FileDescriptor

var file_com_coralogixapis_service_catalog_v1_repeated_widgets_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x2e, 0x76, 0x31, 0x1a, 0x31, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65,
	0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x77, 0x69,
	0x64, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x4c, 0x69, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x72, 0x74, 0x12, 0x4c, 0x0a, 0x0b, 0x6c, 0x69, 0x6e,
	0x65, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x0a, 0x6c, 0x69, 0x6e,
	0x65, 0x43, 0x68, 0x61, 0x72, 0x74, 0x73, 0x12, 0x3f, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4d, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x3c, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d,
	0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_service_catalog_v1_repeated_widgets_proto_rawDescOnce sync.Once
	file_com_coralogixapis_service_catalog_v1_repeated_widgets_proto_rawDescData = file_com_coralogixapis_service_catalog_v1_repeated_widgets_proto_rawDesc
)

func file_com_coralogixapis_service_catalog_v1_repeated_widgets_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_service_catalog_v1_repeated_widgets_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_service_catalog_v1_repeated_widgets_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_service_catalog_v1_repeated_widgets_proto_rawDescData)
	})
	return file_com_coralogixapis_service_catalog_v1_repeated_widgets_proto_rawDescData
}

var file_com_coralogixapis_service_catalog_v1_repeated_widgets_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogixapis_service_catalog_v1_repeated_widgets_proto_goTypes = []any{
	(*RepeatedLineChart)(nil),      // 0: com.coralogixapis.service_catalog.v1.RepeatedLineChart
	(*RepeatedStats)(nil),          // 1: com.coralogixapis.service_catalog.v1.RepeatedStats
	(*v1.LineChart)(nil),           // 2: com.coralogixapis.apm.widgets.v1.LineChart
	(*wrapperspb.UInt64Value)(nil), // 3: google.protobuf.UInt64Value
	(*v1.Stat)(nil),                // 4: com.coralogixapis.apm.widgets.v1.Stat
}
var file_com_coralogixapis_service_catalog_v1_repeated_widgets_proto_depIdxs = []int32{
	2, // 0: com.coralogixapis.service_catalog.v1.RepeatedLineChart.line_charts:type_name -> com.coralogixapis.apm.widgets.v1.LineChart
	3, // 1: com.coralogixapis.service_catalog.v1.RepeatedLineChart.total_amount:type_name -> google.protobuf.UInt64Value
	4, // 2: com.coralogixapis.service_catalog.v1.RepeatedStats.stats:type_name -> com.coralogixapis.apm.widgets.v1.Stat
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_service_catalog_v1_repeated_widgets_proto_init() }
func file_com_coralogixapis_service_catalog_v1_repeated_widgets_proto_init() {
	if File_com_coralogixapis_service_catalog_v1_repeated_widgets_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_service_catalog_v1_repeated_widgets_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_service_catalog_v1_repeated_widgets_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_service_catalog_v1_repeated_widgets_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_service_catalog_v1_repeated_widgets_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_service_catalog_v1_repeated_widgets_proto = out.File
	file_com_coralogixapis_service_catalog_v1_repeated_widgets_proto_rawDesc = nil
	file_com_coralogixapis_service_catalog_v1_repeated_widgets_proto_goTypes = nil
	file_com_coralogixapis_service_catalog_v1_repeated_widgets_proto_depIdxs = nil
}
