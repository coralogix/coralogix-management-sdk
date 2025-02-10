// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: com/coralogix/global_mapping/v1/label_mapping.proto

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

type LabelMapping struct {
	state          protoimpl.MessageState  `protogen:"open.v1"`
	Label          *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	DataSourceType DataSourceType          `protobuf:"varint,2,opt,name=data_source_type,json=dataSourceType,proto3,enum=com.coralogix.global_mapping.v1.DataSourceType" json:"data_source_type,omitempty"`
	// Deprecated: Marked as deprecated in com/coralogix/global_mapping/v1/label_mapping.proto.
	ExtractionTemplate  *wrapperspb.StringValue   `protobuf:"bytes,3,opt,name=extraction_template,json=extractionTemplate,proto3" json:"extraction_template,omitempty"`
	Id                  *wrapperspb.StringValue   `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
	Description         *wrapperspb.StringValue   `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	DisplayName         *wrapperspb.StringValue   `protobuf:"bytes,7,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	ExtractionTemplates []*wrapperspb.StringValue `protobuf:"bytes,8,rep,name=extraction_templates,json=extractionTemplates,proto3" json:"extraction_templates,omitempty"`
	IsCustomLabel       *wrapperspb.BoolValue     `protobuf:"bytes,9,opt,name=is_custom_label,json=isCustomLabel,proto3" json:"is_custom_label,omitempty"`
	DataSource          *DataSource               `protobuf:"bytes,10,opt,name=data_source,json=dataSource,proto3" json:"data_source,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *LabelMapping) Reset() {
	*x = LabelMapping{}
	mi := &file_com_coralogix_global_mapping_v1_label_mapping_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LabelMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelMapping) ProtoMessage() {}

func (x *LabelMapping) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_global_mapping_v1_label_mapping_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelMapping.ProtoReflect.Descriptor instead.
func (*LabelMapping) Descriptor() ([]byte, []int) {
	return file_com_coralogix_global_mapping_v1_label_mapping_proto_rawDescGZIP(), []int{0}
}

func (x *LabelMapping) GetLabel() *wrapperspb.StringValue {
	if x != nil {
		return x.Label
	}
	return nil
}

func (x *LabelMapping) GetDataSourceType() DataSourceType {
	if x != nil {
		return x.DataSourceType
	}
	return DataSourceType_DATA_SOURCE_TYPE_UNSPECIFIED
}

// Deprecated: Marked as deprecated in com/coralogix/global_mapping/v1/label_mapping.proto.
func (x *LabelMapping) GetExtractionTemplate() *wrapperspb.StringValue {
	if x != nil {
		return x.ExtractionTemplate
	}
	return nil
}

func (x *LabelMapping) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *LabelMapping) GetDescription() *wrapperspb.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *LabelMapping) GetDisplayName() *wrapperspb.StringValue {
	if x != nil {
		return x.DisplayName
	}
	return nil
}

func (x *LabelMapping) GetExtractionTemplates() []*wrapperspb.StringValue {
	if x != nil {
		return x.ExtractionTemplates
	}
	return nil
}

func (x *LabelMapping) GetIsCustomLabel() *wrapperspb.BoolValue {
	if x != nil {
		return x.IsCustomLabel
	}
	return nil
}

func (x *LabelMapping) GetDataSource() *DataSource {
	if x != nil {
		return x.DataSource
	}
	return nil
}

var File_com_coralogix_global_mapping_v1_label_mapping_proto protoreflect.FileDescriptor

var file_com_coralogix_global_mapping_v1_label_mapping_proto_rawDesc = []byte{
	0x0a, 0x33, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x31, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x9e, 0x05, 0x0a, 0x0c, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x12, 0x32, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x59, 0x0a, 0x10, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x51, 0x0a, 0x13, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x02, 0x18, 0x01,
	0x52, 0x12, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x3e, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x4f, 0x0a, 0x14, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x13, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x12, 0x42, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x69, 0x73, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x4c, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x52, 0x14, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_global_mapping_v1_label_mapping_proto_rawDescOnce sync.Once
	file_com_coralogix_global_mapping_v1_label_mapping_proto_rawDescData = file_com_coralogix_global_mapping_v1_label_mapping_proto_rawDesc
)

func file_com_coralogix_global_mapping_v1_label_mapping_proto_rawDescGZIP() []byte {
	file_com_coralogix_global_mapping_v1_label_mapping_proto_rawDescOnce.Do(func() {
		file_com_coralogix_global_mapping_v1_label_mapping_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_global_mapping_v1_label_mapping_proto_rawDescData)
	})
	return file_com_coralogix_global_mapping_v1_label_mapping_proto_rawDescData
}

var file_com_coralogix_global_mapping_v1_label_mapping_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogix_global_mapping_v1_label_mapping_proto_goTypes = []any{
	(*LabelMapping)(nil),           // 0: com.coralogix.global_mapping.v1.LabelMapping
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
	(DataSourceType)(0),            // 2: com.coralogix.global_mapping.v1.DataSourceType
	(*wrapperspb.BoolValue)(nil),   // 3: google.protobuf.BoolValue
	(*DataSource)(nil),             // 4: com.coralogix.global_mapping.v1.DataSource
}
var file_com_coralogix_global_mapping_v1_label_mapping_proto_depIdxs = []int32{
	1, // 0: com.coralogix.global_mapping.v1.LabelMapping.label:type_name -> google.protobuf.StringValue
	2, // 1: com.coralogix.global_mapping.v1.LabelMapping.data_source_type:type_name -> com.coralogix.global_mapping.v1.DataSourceType
	1, // 2: com.coralogix.global_mapping.v1.LabelMapping.extraction_template:type_name -> google.protobuf.StringValue
	1, // 3: com.coralogix.global_mapping.v1.LabelMapping.id:type_name -> google.protobuf.StringValue
	1, // 4: com.coralogix.global_mapping.v1.LabelMapping.description:type_name -> google.protobuf.StringValue
	1, // 5: com.coralogix.global_mapping.v1.LabelMapping.display_name:type_name -> google.protobuf.StringValue
	1, // 6: com.coralogix.global_mapping.v1.LabelMapping.extraction_templates:type_name -> google.protobuf.StringValue
	3, // 7: com.coralogix.global_mapping.v1.LabelMapping.is_custom_label:type_name -> google.protobuf.BoolValue
	4, // 8: com.coralogix.global_mapping.v1.LabelMapping.data_source:type_name -> com.coralogix.global_mapping.v1.DataSource
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_com_coralogix_global_mapping_v1_label_mapping_proto_init() }
func file_com_coralogix_global_mapping_v1_label_mapping_proto_init() {
	if File_com_coralogix_global_mapping_v1_label_mapping_proto != nil {
		return
	}
	file_com_coralogix_global_mapping_v1_data_source_type_proto_init()
	file_com_coralogix_global_mapping_v1_data_source_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_global_mapping_v1_label_mapping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_global_mapping_v1_label_mapping_proto_goTypes,
		DependencyIndexes: file_com_coralogix_global_mapping_v1_label_mapping_proto_depIdxs,
		MessageInfos:      file_com_coralogix_global_mapping_v1_label_mapping_proto_msgTypes,
	}.Build()
	File_com_coralogix_global_mapping_v1_label_mapping_proto = out.File
	file_com_coralogix_global_mapping_v1_label_mapping_proto_rawDesc = nil
	file_com_coralogix_global_mapping_v1_label_mapping_proto_goTypes = nil
	file_com_coralogix_global_mapping_v1_label_mapping_proto_depIdxs = nil
}
