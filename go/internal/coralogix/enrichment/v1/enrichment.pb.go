// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.26.1
// source: com/coralogix/enrichment/v1/enrichment.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type Enrichment struct {
	state             protoimpl.MessageState  `protogen:"open.v1"`
	Id                uint32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FieldName         string                  `protobuf:"bytes,2,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	EnrichmentType    *EnrichmentType         `protobuf:"bytes,3,opt,name=enrichment_type,json=enrichmentType,proto3" json:"enrichment_type,omitempty"`
	EnrichedFieldName *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=enriched_field_name,json=enrichedFieldName,proto3,oneof" json:"enriched_field_name,omitempty"`
	SelectedColumns   []string                `protobuf:"bytes,5,rep,name=selected_columns,json=selectedColumns,proto3" json:"selected_columns,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *Enrichment) Reset() {
	*x = Enrichment{}
	mi := &file_com_coralogix_enrichment_v1_enrichment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Enrichment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Enrichment) ProtoMessage() {}

func (x *Enrichment) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Enrichment.ProtoReflect.Descriptor instead.
func (*Enrichment) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_proto_rawDescGZIP(), []int{0}
}

func (x *Enrichment) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Enrichment) GetFieldName() string {
	if x != nil {
		return x.FieldName
	}
	return ""
}

func (x *Enrichment) GetEnrichmentType() *EnrichmentType {
	if x != nil {
		return x.EnrichmentType
	}
	return nil
}

func (x *Enrichment) GetEnrichedFieldName() *wrapperspb.StringValue {
	if x != nil {
		return x.EnrichedFieldName
	}
	return nil
}

func (x *Enrichment) GetSelectedColumns() []string {
	if x != nil {
		return x.SelectedColumns
	}
	return nil
}

var File_com_coralogix_enrichment_v1_enrichment_proto protoreflect.FileDescriptor

var file_com_coralogix_enrichment_v1_enrichment_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e,
	0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x6e,
	0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x65, 0x6e, 0x72, 0x69, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1,
	0x04, 0x0a, 0x0a, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x06, 0x92, 0x41, 0x03, 0x4a, 0x01,
	0x31, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0x92, 0x41, 0x03, 0x4a, 0x01,
	0x31, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x54, 0x0a, 0x0f,
	0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x59, 0x0a, 0x13, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x06, 0x92,
	0x41, 0x03, 0x4a, 0x01, 0x31, 0x48, 0x00, 0x52, 0x11, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x65,
	0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x46, 0x0a,
	0x10, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x42, 0x1b, 0x92, 0x41, 0x18, 0x4a, 0x16, 0x5b, 0x22,
	0x63, 0x69, 0x74, 0x79, 0x22, 0x2c, 0x20, 0x22, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x5d, 0x52, 0x0f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x43, 0x6f,
	0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x3a, 0xe2, 0x01, 0x92, 0x41, 0xde, 0x01, 0x0a, 0x5e, 0x2a, 0x0a,
	0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x32, 0x2c, 0x54, 0x68, 0x69, 0x73,
	0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x20,
	0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x73, 0x20, 0x61, 0x6e, 0x20, 0x65, 0x6e,
	0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0xd2, 0x01, 0x02, 0x69, 0x64, 0xd2, 0x01, 0x0a,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0xd2, 0x01, 0x0f, 0x65, 0x6e, 0x72,
	0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2a, 0x7c, 0x0a, 0x1f,
	0x46, 0x69, 0x6e, 0x64, 0x20, 0x6f, 0x75, 0x74, 0x20, 0x6d, 0x6f, 0x72, 0x65, 0x20, 0x61, 0x62,
	0x6f, 0x75, 0x74, 0x20, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x59, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x63, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2d, 0x67, 0x75, 0x69, 0x64, 0x65, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x72, 0x69,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2d, 0x65,
	0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x65,
	0x6e, 0x72, 0x69, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_enrichment_v1_enrichment_proto_rawDescOnce sync.Once
	file_com_coralogix_enrichment_v1_enrichment_proto_rawDescData = file_com_coralogix_enrichment_v1_enrichment_proto_rawDesc
)

func file_com_coralogix_enrichment_v1_enrichment_proto_rawDescGZIP() []byte {
	file_com_coralogix_enrichment_v1_enrichment_proto_rawDescOnce.Do(func() {
		file_com_coralogix_enrichment_v1_enrichment_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_enrichment_v1_enrichment_proto_rawDescData)
	})
	return file_com_coralogix_enrichment_v1_enrichment_proto_rawDescData
}

var file_com_coralogix_enrichment_v1_enrichment_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogix_enrichment_v1_enrichment_proto_goTypes = []any{
	(*Enrichment)(nil),             // 0: com.coralogix.enrichment.v1.Enrichment
	(*EnrichmentType)(nil),         // 1: com.coralogix.enrichment.v1.EnrichmentType
	(*wrapperspb.StringValue)(nil), // 2: google.protobuf.StringValue
}
var file_com_coralogix_enrichment_v1_enrichment_proto_depIdxs = []int32{
	1, // 0: com.coralogix.enrichment.v1.Enrichment.enrichment_type:type_name -> com.coralogix.enrichment.v1.EnrichmentType
	2, // 1: com.coralogix.enrichment.v1.Enrichment.enriched_field_name:type_name -> google.protobuf.StringValue
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_com_coralogix_enrichment_v1_enrichment_proto_init() }
func file_com_coralogix_enrichment_v1_enrichment_proto_init() {
	if File_com_coralogix_enrichment_v1_enrichment_proto != nil {
		return
	}
	file_com_coralogix_enrichment_v1_enrichment_type_proto_init()
	file_com_coralogix_enrichment_v1_enrichment_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_enrichment_v1_enrichment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_enrichment_v1_enrichment_proto_goTypes,
		DependencyIndexes: file_com_coralogix_enrichment_v1_enrichment_proto_depIdxs,
		MessageInfos:      file_com_coralogix_enrichment_v1_enrichment_proto_msgTypes,
	}.Build()
	File_com_coralogix_enrichment_v1_enrichment_proto = out.File
	file_com_coralogix_enrichment_v1_enrichment_proto_rawDesc = nil
	file_com_coralogix_enrichment_v1_enrichment_proto_goTypes = nil
	file_com_coralogix_enrichment_v1_enrichment_proto_depIdxs = nil
}
