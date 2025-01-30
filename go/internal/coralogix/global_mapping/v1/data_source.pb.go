// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.2
// source: com/coralogix/global_mapping/v1/data_source.proto

package v1

import (
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

type DataSource struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Provider      *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	Exporter      *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=exporter,proto3" json:"exporter,omitempty"`
	LabelsMetric  *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=labels_metric,json=labelsMetric,proto3" json:"labels_metric,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DataSource) Reset() {
	*x = DataSource{}
	mi := &file_com_coralogix_global_mapping_v1_data_source_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSource) ProtoMessage() {}

func (x *DataSource) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_global_mapping_v1_data_source_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSource.ProtoReflect.Descriptor instead.
func (*DataSource) Descriptor() ([]byte, []int) {
	return file_com_coralogix_global_mapping_v1_data_source_proto_rawDescGZIP(), []int{0}
}

func (x *DataSource) GetProvider() *wrapperspb.StringValue {
	if x != nil {
		return x.Provider
	}
	return nil
}

func (x *DataSource) GetExporter() *wrapperspb.StringValue {
	if x != nil {
		return x.Exporter
	}
	return nil
}

func (x *DataSource) GetLabelsMetric() *wrapperspb.StringValue {
	if x != nil {
		return x.LabelsMetric
	}
	return nil
}

var File_com_coralogix_global_mapping_v1_data_source_proto protoreflect.FileDescriptor

var file_com_coralogix_global_mapping_v1_data_source_proto_rawDesc = string([]byte{
	0x0a, 0x31, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x01, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x38, 0x0a,
	0x08, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x65,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x0d, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_com_coralogix_global_mapping_v1_data_source_proto_rawDescOnce sync.Once
	file_com_coralogix_global_mapping_v1_data_source_proto_rawDescData []byte
)

func file_com_coralogix_global_mapping_v1_data_source_proto_rawDescGZIP() []byte {
	file_com_coralogix_global_mapping_v1_data_source_proto_rawDescOnce.Do(func() {
		file_com_coralogix_global_mapping_v1_data_source_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogix_global_mapping_v1_data_source_proto_rawDesc), len(file_com_coralogix_global_mapping_v1_data_source_proto_rawDesc)))
	})
	return file_com_coralogix_global_mapping_v1_data_source_proto_rawDescData
}

var file_com_coralogix_global_mapping_v1_data_source_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogix_global_mapping_v1_data_source_proto_goTypes = []any{
	(*DataSource)(nil),             // 0: com.coralogix.global_mapping.v1.DataSource
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
}
var file_com_coralogix_global_mapping_v1_data_source_proto_depIdxs = []int32{
	1, // 0: com.coralogix.global_mapping.v1.DataSource.provider:type_name -> google.protobuf.StringValue
	1, // 1: com.coralogix.global_mapping.v1.DataSource.exporter:type_name -> google.protobuf.StringValue
	1, // 2: com.coralogix.global_mapping.v1.DataSource.labels_metric:type_name -> google.protobuf.StringValue
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_com_coralogix_global_mapping_v1_data_source_proto_init() }
func file_com_coralogix_global_mapping_v1_data_source_proto_init() {
	if File_com_coralogix_global_mapping_v1_data_source_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogix_global_mapping_v1_data_source_proto_rawDesc), len(file_com_coralogix_global_mapping_v1_data_source_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_global_mapping_v1_data_source_proto_goTypes,
		DependencyIndexes: file_com_coralogix_global_mapping_v1_data_source_proto_depIdxs,
		MessageInfos:      file_com_coralogix_global_mapping_v1_data_source_proto_msgTypes,
	}.Build()
	File_com_coralogix_global_mapping_v1_data_source_proto = out.File
	file_com_coralogix_global_mapping_v1_data_source_proto_goTypes = nil
	file_com_coralogix_global_mapping_v1_data_source_proto_depIdxs = nil
}
