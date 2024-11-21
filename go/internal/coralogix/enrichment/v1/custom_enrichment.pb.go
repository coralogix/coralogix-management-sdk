// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: com/coralogix/enrichment/v1/custom_enrichment.proto

package v1

import (
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type CustomEnrichment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Version     uint32 `protobuf:"varint,5,opt,name=version,proto3" json:"version,omitempty"`
	IsQueryOnly bool   `protobuf:"varint,6,opt,name=is_query_only,json=isQueryOnly,proto3" json:"is_query_only,omitempty"`
}

func (x *CustomEnrichment) Reset() {
	*x = CustomEnrichment{}
	mi := &file_com_coralogix_enrichment_v1_custom_enrichment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomEnrichment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomEnrichment) ProtoMessage() {}

func (x *CustomEnrichment) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_custom_enrichment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomEnrichment.ProtoReflect.Descriptor instead.
func (*CustomEnrichment) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_custom_enrichment_proto_rawDescGZIP(), []int{0}
}

func (x *CustomEnrichment) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CustomEnrichment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CustomEnrichment) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CustomEnrichment) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *CustomEnrichment) GetIsQueryOnly() bool {
	if x != nil {
		return x.IsQueryOnly
	}
	return false
}

type CustomEnrichmentData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Definition *CustomEnrichment `protobuf:"bytes,1,opt,name=definition,proto3" json:"definition,omitempty"`
	// Types that are assignable to Content:
	//
	//	*CustomEnrichmentData_Textual
	//	*CustomEnrichmentData_Binary
	Content isCustomEnrichmentData_Content `protobuf_oneof:"content"`
}

func (x *CustomEnrichmentData) Reset() {
	*x = CustomEnrichmentData{}
	mi := &file_com_coralogix_enrichment_v1_custom_enrichment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomEnrichmentData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomEnrichmentData) ProtoMessage() {}

func (x *CustomEnrichmentData) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_custom_enrichment_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomEnrichmentData.ProtoReflect.Descriptor instead.
func (*CustomEnrichmentData) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_custom_enrichment_proto_rawDescGZIP(), []int{1}
}

func (x *CustomEnrichmentData) GetDefinition() *CustomEnrichment {
	if x != nil {
		return x.Definition
	}
	return nil
}

func (m *CustomEnrichmentData) GetContent() isCustomEnrichmentData_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *CustomEnrichmentData) GetTextual() *wrappers.StringValue {
	if x, ok := x.GetContent().(*CustomEnrichmentData_Textual); ok {
		return x.Textual
	}
	return nil
}

func (x *CustomEnrichmentData) GetBinary() *wrappers.BytesValue {
	if x, ok := x.GetContent().(*CustomEnrichmentData_Binary); ok {
		return x.Binary
	}
	return nil
}

type isCustomEnrichmentData_Content interface {
	isCustomEnrichmentData_Content()
}

type CustomEnrichmentData_Textual struct {
	Textual *wrappers.StringValue `protobuf:"bytes,3,opt,name=textual,proto3,oneof"`
}

type CustomEnrichmentData_Binary struct {
	Binary *wrappers.BytesValue `protobuf:"bytes,4,opt,name=binary,proto3,oneof"`
}

func (*CustomEnrichmentData_Textual) isCustomEnrichmentData_Content() {}

func (*CustomEnrichmentData_Binary) isCustomEnrichmentData_Content() {}

var File_com_coralogix_enrichment_v1_custom_enrichment_proto protoreflect.FileDescriptor

var file_com_coralogix_enrichment_v1_custom_enrichment_proto_rawDesc = []byte{
	0x0a, 0x33, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x96, 0x01, 0x0a, 0x10, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x72,
	0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x69, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0xe1, 0x01, 0x0a, 0x14,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x4d, 0x0a, 0x0a, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x72,
	0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x07, 0x74, 0x65, 0x78, 0x74, 0x75, 0x61, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x48, 0x00, 0x52, 0x07, 0x74, 0x65, 0x78, 0x74, 0x75, 0x61, 0x6c, 0x12, 0x35, 0x0a,
	0x06, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x06, 0x62, 0x69,
	0x6e, 0x61, 0x72, 0x79, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_enrichment_v1_custom_enrichment_proto_rawDescOnce sync.Once
	file_com_coralogix_enrichment_v1_custom_enrichment_proto_rawDescData = file_com_coralogix_enrichment_v1_custom_enrichment_proto_rawDesc
)

func file_com_coralogix_enrichment_v1_custom_enrichment_proto_rawDescGZIP() []byte {
	file_com_coralogix_enrichment_v1_custom_enrichment_proto_rawDescOnce.Do(func() {
		file_com_coralogix_enrichment_v1_custom_enrichment_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_enrichment_v1_custom_enrichment_proto_rawDescData)
	})
	return file_com_coralogix_enrichment_v1_custom_enrichment_proto_rawDescData
}

var file_com_coralogix_enrichment_v1_custom_enrichment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogix_enrichment_v1_custom_enrichment_proto_goTypes = []any{
	(*CustomEnrichment)(nil),     // 0: com.coralogix.enrichment.v1.CustomEnrichment
	(*CustomEnrichmentData)(nil), // 1: com.coralogix.enrichment.v1.CustomEnrichmentData
	(*wrappers.StringValue)(nil), // 2: google.protobuf.StringValue
	(*wrappers.BytesValue)(nil),  // 3: google.protobuf.BytesValue
}
var file_com_coralogix_enrichment_v1_custom_enrichment_proto_depIdxs = []int32{
	0, // 0: com.coralogix.enrichment.v1.CustomEnrichmentData.definition:type_name -> com.coralogix.enrichment.v1.CustomEnrichment
	2, // 1: com.coralogix.enrichment.v1.CustomEnrichmentData.textual:type_name -> google.protobuf.StringValue
	3, // 2: com.coralogix.enrichment.v1.CustomEnrichmentData.binary:type_name -> google.protobuf.BytesValue
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_com_coralogix_enrichment_v1_custom_enrichment_proto_init() }
func file_com_coralogix_enrichment_v1_custom_enrichment_proto_init() {
	if File_com_coralogix_enrichment_v1_custom_enrichment_proto != nil {
		return
	}
	file_com_coralogix_enrichment_v1_custom_enrichment_proto_msgTypes[1].OneofWrappers = []any{
		(*CustomEnrichmentData_Textual)(nil),
		(*CustomEnrichmentData_Binary)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_enrichment_v1_custom_enrichment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_enrichment_v1_custom_enrichment_proto_goTypes,
		DependencyIndexes: file_com_coralogix_enrichment_v1_custom_enrichment_proto_depIdxs,
		MessageInfos:      file_com_coralogix_enrichment_v1_custom_enrichment_proto_msgTypes,
	}.Build()
	File_com_coralogix_enrichment_v1_custom_enrichment_proto = out.File
	file_com_coralogix_enrichment_v1_custom_enrichment_proto_rawDesc = nil
	file_com_coralogix_enrichment_v1_custom_enrichment_proto_goTypes = nil
	file_com_coralogix_enrichment_v1_custom_enrichment_proto_depIdxs = nil
}
