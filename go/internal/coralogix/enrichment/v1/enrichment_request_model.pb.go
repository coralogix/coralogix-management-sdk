// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: com/coralogix/enrichment/v1/enrichment_request_model.proto

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

type EnrichmentRequestModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldName         *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	EnrichmentType    *EnrichmentType         `protobuf:"bytes,2,opt,name=enrichment_type,json=enrichmentType,proto3" json:"enrichment_type,omitempty"`
	EnrichedFieldName *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=enriched_field_name,json=enrichedFieldName,proto3,oneof" json:"enriched_field_name,omitempty"`
}

func (x *EnrichmentRequestModel) Reset() {
	*x = EnrichmentRequestModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_enrichment_v1_enrichment_request_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrichmentRequestModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrichmentRequestModel) ProtoMessage() {}

func (x *EnrichmentRequestModel) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_request_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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

type EnrichmentFieldDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the field name of the key targeted for enrichment
	FieldName *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	// the field of the enriched key after enrichment
	EnrichedFieldName *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=enriched_field_name,json=enrichedFieldName,proto3,oneof" json:"enriched_field_name,omitempty"`
}

func (x *EnrichmentFieldDefinition) Reset() {
	*x = EnrichmentFieldDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_enrichment_v1_enrichment_request_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrichmentFieldDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrichmentFieldDefinition) ProtoMessage() {}

func (x *EnrichmentFieldDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_request_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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

var File_com_coralogix_enrichment_v1_enrichment_request_model_proto protoreflect.FileDescriptor

var file_com_coralogix_enrichment_v1_enrichment_request_model_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e,
	0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x6e, 0x72, 0x69,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x02, 0x0a,
	0x16, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x3b, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x54, 0x0a, 0x0f, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x6e,
	0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x72, 0x69,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x65, 0x6e, 0x72, 0x69,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x51, 0x0a, 0x13, 0x65, 0x6e,
	0x72, 0x69, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x11, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x65,
	0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x16, 0x0a,
	0x14, 0x5f, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xc3, 0x01, 0x0a, 0x19, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x51, 0x0a, 0x13, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x11, 0x65,
	0x6e, 0x72, 0x69, 0x63, 0x68, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x65, 0x64,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_enrichment_v1_enrichment_request_model_proto_rawDescOnce sync.Once
	file_com_coralogix_enrichment_v1_enrichment_request_model_proto_rawDescData = file_com_coralogix_enrichment_v1_enrichment_request_model_proto_rawDesc
)

func file_com_coralogix_enrichment_v1_enrichment_request_model_proto_rawDescGZIP() []byte {
	file_com_coralogix_enrichment_v1_enrichment_request_model_proto_rawDescOnce.Do(func() {
		file_com_coralogix_enrichment_v1_enrichment_request_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_enrichment_v1_enrichment_request_model_proto_rawDescData)
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
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_enrichment_v1_enrichment_request_model_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*EnrichmentRequestModel); i {
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
		file_com_coralogix_enrichment_v1_enrichment_request_model_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*EnrichmentFieldDefinition); i {
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
	file_com_coralogix_enrichment_v1_enrichment_request_model_proto_msgTypes[0].OneofWrappers = []any{}
	file_com_coralogix_enrichment_v1_enrichment_request_model_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_enrichment_v1_enrichment_request_model_proto_rawDesc,
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
	file_com_coralogix_enrichment_v1_enrichment_request_model_proto_rawDesc = nil
	file_com_coralogix_enrichment_v1_enrichment_request_model_proto_goTypes = nil
	file_com_coralogix_enrichment_v1_enrichment_request_model_proto_depIdxs = nil
}
