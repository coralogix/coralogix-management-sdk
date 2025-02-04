// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: com/coralogix/archive/dataset/v2/dataset_rules.proto

package v2

import (
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

type PredicateType int32

const (
	PredicateType_PREDICATE_TYPE_UNSPECIFIED PredicateType = 0
	PredicateType_PREDICATE_TYPE_REGEX       PredicateType = 1
	PredicateType_PREDICATE_TYPE_DPXL        PredicateType = 2
	PredicateType_PREDICATE_TYPE_EXACT       PredicateType = 3
)

// Enum value maps for PredicateType.
var (
	PredicateType_name = map[int32]string{
		0: "PREDICATE_TYPE_UNSPECIFIED",
		1: "PREDICATE_TYPE_REGEX",
		2: "PREDICATE_TYPE_DPXL",
		3: "PREDICATE_TYPE_EXACT",
	}
	PredicateType_value = map[string]int32{
		"PREDICATE_TYPE_UNSPECIFIED": 0,
		"PREDICATE_TYPE_REGEX":       1,
		"PREDICATE_TYPE_DPXL":        2,
		"PREDICATE_TYPE_EXACT":       3,
	}
)

func (x PredicateType) Enum() *PredicateType {
	p := new(PredicateType)
	*p = x
	return p
}

func (x PredicateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PredicateType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_archive_dataset_v2_dataset_rules_proto_enumTypes[0].Descriptor()
}

func (PredicateType) Type() protoreflect.EnumType {
	return &file_com_coralogix_archive_dataset_v2_dataset_rules_proto_enumTypes[0]
}

func (x PredicateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PredicateType.Descriptor instead.
func (PredicateType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_dataset_rules_proto_rawDescGZIP(), []int{0}
}

type DatasetRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dataspace         *Dataspace               `protobuf:"bytes,1,opt,name=dataspace,proto3" json:"dataspace,omitempty"`
	Predicate         *Predicate               `protobuf:"bytes,2,opt,name=predicate,proto3" json:"predicate,omitempty"`
	Locations         []*IngestionLocationSpec `protobuf:"bytes,3,rep,name=locations,proto3" json:"locations,omitempty"`
	SchemaTemplateIds []string                 `protobuf:"bytes,4,rep,name=schema_template_ids,json=schemaTemplateIds,proto3" json:"schema_template_ids,omitempty"`
}

func (x *DatasetRule) Reset() {
	*x = DatasetRule{}
	mi := &file_com_coralogix_archive_dataset_v2_dataset_rules_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DatasetRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatasetRule) ProtoMessage() {}

func (x *DatasetRule) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_dataset_v2_dataset_rules_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatasetRule.ProtoReflect.Descriptor instead.
func (*DatasetRule) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_dataset_rules_proto_rawDescGZIP(), []int{0}
}

func (x *DatasetRule) GetDataspace() *Dataspace {
	if x != nil {
		return x.Dataspace
	}
	return nil
}

func (x *DatasetRule) GetPredicate() *Predicate {
	if x != nil {
		return x.Predicate
	}
	return nil
}

func (x *DatasetRule) GetLocations() []*IngestionLocationSpec {
	if x != nil {
		return x.Locations
	}
	return nil
}

func (x *DatasetRule) GetSchemaTemplateIds() []string {
	if x != nil {
		return x.SchemaTemplateIds
	}
	return nil
}

type Predicate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type PredicateType `protobuf:"varint,1,opt,name=type,proto3,enum=com.coralogix.archive.dataset.v2.PredicateType" json:"type,omitempty"`
	Text string        `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Predicate) Reset() {
	*x = Predicate{}
	mi := &file_com_coralogix_archive_dataset_v2_dataset_rules_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Predicate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Predicate) ProtoMessage() {}

func (x *Predicate) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_dataset_v2_dataset_rules_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Predicate.ProtoReflect.Descriptor instead.
func (*Predicate) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_dataset_rules_proto_rawDescGZIP(), []int{1}
}

func (x *Predicate) GetType() PredicateType {
	if x != nil {
		return x.Type
	}
	return PredicateType_PREDICATE_TYPE_UNSPECIFIED
}

func (x *Predicate) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_com_coralogix_archive_dataset_v2_dataset_rules_proto protoreflect.FileDescriptor

var file_com_coralogix_archive_dataset_v2_dataset_rules_proto_rawDesc = []byte{
	0x0a, 0x34, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2f,
	0x76, 0x32, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x1a, 0x38, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x68, 0x79, 0x73, 0x69,
	0x63, 0x61, 0x6c, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x2f, 0x76, 0x32, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xaa, 0x02, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x75,
	0x6c, 0x65, 0x12, 0x49, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x49, 0x0a,
	0x09, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x09, 0x70,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x55, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x49,
	0x6e, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x70, 0x65, 0x63, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x2e, 0x0a, 0x13, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x73, 0x22,
	0x64, 0x0a, 0x09, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x43, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x2a, 0x7c, 0x0a, 0x0d, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x52, 0x45, 0x44, 0x49, 0x43,
	0x41, 0x54, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x52, 0x45, 0x44, 0x49, 0x43,
	0x41, 0x54, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x47, 0x45, 0x58, 0x10, 0x01,
	0x12, 0x17, 0x0a, 0x13, 0x50, 0x52, 0x45, 0x44, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x44, 0x50, 0x58, 0x4c, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x52, 0x45,
	0x44, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x58, 0x41, 0x43,
	0x54, 0x10, 0x03, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_archive_dataset_v2_dataset_rules_proto_rawDescOnce sync.Once
	file_com_coralogix_archive_dataset_v2_dataset_rules_proto_rawDescData = file_com_coralogix_archive_dataset_v2_dataset_rules_proto_rawDesc
)

func file_com_coralogix_archive_dataset_v2_dataset_rules_proto_rawDescGZIP() []byte {
	file_com_coralogix_archive_dataset_v2_dataset_rules_proto_rawDescOnce.Do(func() {
		file_com_coralogix_archive_dataset_v2_dataset_rules_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_archive_dataset_v2_dataset_rules_proto_rawDescData)
	})
	return file_com_coralogix_archive_dataset_v2_dataset_rules_proto_rawDescData
}

var file_com_coralogix_archive_dataset_v2_dataset_rules_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogix_archive_dataset_v2_dataset_rules_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogix_archive_dataset_v2_dataset_rules_proto_goTypes = []any{
	(PredicateType)(0),            // 0: com.coralogix.archive.dataset.v2.PredicateType
	(*DatasetRule)(nil),           // 1: com.coralogix.archive.dataset.v2.DatasetRule
	(*Predicate)(nil),             // 2: com.coralogix.archive.dataset.v2.Predicate
	(*Dataspace)(nil),             // 3: com.coralogix.archive.dataset.v2.Dataspace
	(*IngestionLocationSpec)(nil), // 4: com.coralogix.archive.dataset.v2.IngestionLocationSpec
}
var file_com_coralogix_archive_dataset_v2_dataset_rules_proto_depIdxs = []int32{
	3, // 0: com.coralogix.archive.dataset.v2.DatasetRule.dataspace:type_name -> com.coralogix.archive.dataset.v2.Dataspace
	2, // 1: com.coralogix.archive.dataset.v2.DatasetRule.predicate:type_name -> com.coralogix.archive.dataset.v2.Predicate
	4, // 2: com.coralogix.archive.dataset.v2.DatasetRule.locations:type_name -> com.coralogix.archive.dataset.v2.IngestionLocationSpec
	0, // 3: com.coralogix.archive.dataset.v2.Predicate.type:type_name -> com.coralogix.archive.dataset.v2.PredicateType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_com_coralogix_archive_dataset_v2_dataset_rules_proto_init() }
func file_com_coralogix_archive_dataset_v2_dataset_rules_proto_init() {
	if File_com_coralogix_archive_dataset_v2_dataset_rules_proto != nil {
		return
	}
	file_com_coralogix_archive_dataset_v2_physical_location_proto_init()
	file_com_coralogix_archive_dataset_v2_dataset_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_archive_dataset_v2_dataset_rules_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_archive_dataset_v2_dataset_rules_proto_goTypes,
		DependencyIndexes: file_com_coralogix_archive_dataset_v2_dataset_rules_proto_depIdxs,
		EnumInfos:         file_com_coralogix_archive_dataset_v2_dataset_rules_proto_enumTypes,
		MessageInfos:      file_com_coralogix_archive_dataset_v2_dataset_rules_proto_msgTypes,
	}.Build()
	File_com_coralogix_archive_dataset_v2_dataset_rules_proto = out.File
	file_com_coralogix_archive_dataset_v2_dataset_rules_proto_rawDesc = nil
	file_com_coralogix_archive_dataset_v2_dataset_rules_proto_goTypes = nil
	file_com_coralogix_archive_dataset_v2_dataset_rules_proto_depIdxs = nil
}
