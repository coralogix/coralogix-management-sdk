// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.26.1
// source: com/coralogix/enrichment/v1/enrichment_request_model.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type EnrichmentRequestModel struct {
	state             protoimpl.MessageState  `protogen:"open.v1"`
	FieldName         *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	EnrichmentType    *EnrichmentType         `protobuf:"bytes,2,opt,name=enrichment_type,json=enrichmentType,proto3" json:"enrichment_type,omitempty"`
	EnrichedFieldName *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=enriched_field_name,json=enrichedFieldName,proto3,oneof" json:"enriched_field_name,omitempty"`
	SelectedColumns   []string                `protobuf:"bytes,4,rep,name=selected_columns,json=selectedColumns,proto3" json:"selected_columns,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *EnrichmentRequestModel) Reset() {
	*x = EnrichmentRequestModel{}
	mi := &file_com_coralogix_enrichment_v1_enrichment_request_model_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EnrichmentRequestModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrichmentRequestModel) ProtoMessage() {}

func (x *EnrichmentRequestModel) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_request_model_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnrichmentRequestModel.ProtoReflect.Descriptor instead.
func (*EnrichmentRequestModel) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_request_model_proto_rawDescGZIP(), []int{0}
}

func (x *EnrichmentRequestModel) GetFieldName() *wrapperspb.StringValue {
	if x != nil {
		return x.FieldName
	}
	return nil
}

func (x *EnrichmentRequestModel) GetEnrichmentType() *EnrichmentType {
	if x != nil {
		return x.EnrichmentType
	}
	return nil
}

func (x *EnrichmentRequestModel) GetEnrichedFieldName() *wrapperspb.StringValue {
	if x != nil {
		return x.EnrichedFieldName
	}
	return nil
}

func (x *EnrichmentRequestModel) GetSelectedColumns() []string {
	if x != nil {
		return x.SelectedColumns
	}
	return nil
}

type EnrichmentFieldDefinition struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the field name of the key targeted for enrichment
	FieldName *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	// the field of the enriched key after enrichment
	EnrichedFieldName *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=enriched_field_name,json=enrichedFieldName,proto3,oneof" json:"enriched_field_name,omitempty"`
	// the columns to be selected from the enrichment
	SelectedColumns []string `protobuf:"bytes,4,rep,name=selected_columns,json=selectedColumns,proto3" json:"selected_columns,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *EnrichmentFieldDefinition) Reset() {
	*x = EnrichmentFieldDefinition{}
	mi := &file_com_coralogix_enrichment_v1_enrichment_request_model_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EnrichmentFieldDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrichmentFieldDefinition) ProtoMessage() {}

func (x *EnrichmentFieldDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_request_model_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnrichmentFieldDefinition.ProtoReflect.Descriptor instead.
func (*EnrichmentFieldDefinition) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_request_model_proto_rawDescGZIP(), []int{1}
}

func (x *EnrichmentFieldDefinition) GetFieldName() *wrapperspb.StringValue {
	if x != nil {
		return x.FieldName
	}
	return nil
}

func (x *EnrichmentFieldDefinition) GetEnrichedFieldName() *wrapperspb.StringValue {
	if x != nil {
		return x.EnrichedFieldName
	}
	return nil
}

func (x *EnrichmentFieldDefinition) GetSelectedColumns() []string {
	if x != nil {
		return x.SelectedColumns
	}
	return nil
}

var File_com_coralogix_enrichment_v1_enrichment_request_model_proto protoreflect.FileDescriptor

const file_com_coralogix_enrichment_v1_enrichment_request_model_proto_rawDesc = "" +
	"\n" +
	":com/coralogix/enrichment/v1/enrichment_request_model.proto\x12\x1bcom.coralogix.enrichment.v1\x1a\x1egoogle/protobuf/wrappers.proto\x1a1com/coralogix/enrichment/v1/enrichment_type.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\xad\x04\n" +
	"\x16EnrichmentRequestModel\x12M\n" +
	"\n" +
	"field_name\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueB\x10\x92A\rJ\v\"sourceIPs\"R\tfieldName\x12T\n" +
	"\x0fenrichment_type\x18\x02 \x01(\v2+.com.coralogix.enrichment.v1.EnrichmentTypeR\x0eenrichmentType\x12Q\n" +
	"\x13enriched_field_name\x18\x03 \x01(\v2\x1c.google.protobuf.StringValueH\x00R\x11enrichedFieldName\x88\x01\x01\x12)\n" +
	"\x10selected_columns\x18\x04 \x03(\tR\x0fselectedColumns:\xd7\x01\x92A\xd3\x01\n" +
	"S*\x14Enrichment Prototype2\x1cThe enrichment request model\xd2\x01\n" +
	"field_name\xd2\x01\x0fenrichment_type*|\n" +
	"\x1fFind out more about enrichments\x12Yhttps://coralogix.com/docs/user-guides/data-transformation/enrichments/custom-enrichment/B\x16\n" +
	"\x14_enriched_field_name\"\xee\x01\n" +
	"\x19EnrichmentFieldDefinition\x12;\n" +
	"\n" +
	"field_name\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueR\tfieldName\x12Q\n" +
	"\x13enriched_field_name\x18\x03 \x01(\v2\x1c.google.protobuf.StringValueH\x00R\x11enrichedFieldName\x88\x01\x01\x12)\n" +
	"\x10selected_columns\x18\x04 \x03(\tR\x0fselectedColumnsB\x16\n" +
	"\x14_enriched_field_nameb\x06proto3"

var (
	file_com_coralogix_enrichment_v1_enrichment_request_model_proto_rawDescOnce sync.Once
	file_com_coralogix_enrichment_v1_enrichment_request_model_proto_rawDescData []byte
)

func file_com_coralogix_enrichment_v1_enrichment_request_model_proto_rawDescGZIP() []byte {
	file_com_coralogix_enrichment_v1_enrichment_request_model_proto_rawDescOnce.Do(func() {
		file_com_coralogix_enrichment_v1_enrichment_request_model_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogix_enrichment_v1_enrichment_request_model_proto_rawDesc), len(file_com_coralogix_enrichment_v1_enrichment_request_model_proto_rawDesc)))
	})
	return file_com_coralogix_enrichment_v1_enrichment_request_model_proto_rawDescData
}

var file_com_coralogix_enrichment_v1_enrichment_request_model_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogix_enrichment_v1_enrichment_request_model_proto_goTypes = []any{
	(*EnrichmentRequestModel)(nil),    // 0: com.coralogix.enrichment.v1.EnrichmentRequestModel
	(*EnrichmentFieldDefinition)(nil), // 1: com.coralogix.enrichment.v1.EnrichmentFieldDefinition
	(*wrapperspb.StringValue)(nil),    // 2: google.protobuf.StringValue
	(*EnrichmentType)(nil),            // 3: com.coralogix.enrichment.v1.EnrichmentType
}
var file_com_coralogix_enrichment_v1_enrichment_request_model_proto_depIdxs = []int32{
	2, // 0: com.coralogix.enrichment.v1.EnrichmentRequestModel.field_name:type_name -> google.protobuf.StringValue
	3, // 1: com.coralogix.enrichment.v1.EnrichmentRequestModel.enrichment_type:type_name -> com.coralogix.enrichment.v1.EnrichmentType
	2, // 2: com.coralogix.enrichment.v1.EnrichmentRequestModel.enriched_field_name:type_name -> google.protobuf.StringValue
	2, // 3: com.coralogix.enrichment.v1.EnrichmentFieldDefinition.field_name:type_name -> google.protobuf.StringValue
	2, // 4: com.coralogix.enrichment.v1.EnrichmentFieldDefinition.enriched_field_name:type_name -> google.protobuf.StringValue
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_com_coralogix_enrichment_v1_enrichment_request_model_proto_init() }
func file_com_coralogix_enrichment_v1_enrichment_request_model_proto_init() {
	if File_com_coralogix_enrichment_v1_enrichment_request_model_proto != nil {
		return
	}
	file_com_coralogix_enrichment_v1_enrichment_type_proto_init()
	file_com_coralogix_enrichment_v1_enrichment_request_model_proto_msgTypes[0].OneofWrappers = []any{}
	file_com_coralogix_enrichment_v1_enrichment_request_model_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogix_enrichment_v1_enrichment_request_model_proto_rawDesc), len(file_com_coralogix_enrichment_v1_enrichment_request_model_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_enrichment_v1_enrichment_request_model_proto_goTypes,
		DependencyIndexes: file_com_coralogix_enrichment_v1_enrichment_request_model_proto_depIdxs,
		MessageInfos:      file_com_coralogix_enrichment_v1_enrichment_request_model_proto_msgTypes,
	}.Build()
	File_com_coralogix_enrichment_v1_enrichment_request_model_proto = out.File
	file_com_coralogix_enrichment_v1_enrichment_request_model_proto_goTypes = nil
	file_com_coralogix_enrichment_v1_enrichment_request_model_proto_depIdxs = nil
}
