// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: com/coralogixapis/dashboards/v1/common/observation_field.proto

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

type DatasetScope int32

const (
	DatasetScope_DATASET_SCOPE_UNSPECIFIED DatasetScope = 0
	DatasetScope_DATASET_SCOPE_USER_DATA   DatasetScope = 1
	DatasetScope_DATASET_SCOPE_LABEL       DatasetScope = 2
	DatasetScope_DATASET_SCOPE_METADATA    DatasetScope = 3
)

// Enum value maps for DatasetScope.
var (
	DatasetScope_name = map[int32]string{
		0: "DATASET_SCOPE_UNSPECIFIED",
		1: "DATASET_SCOPE_USER_DATA",
		2: "DATASET_SCOPE_LABEL",
		3: "DATASET_SCOPE_METADATA",
	}
	DatasetScope_value = map[string]int32{
		"DATASET_SCOPE_UNSPECIFIED": 0,
		"DATASET_SCOPE_USER_DATA":   1,
		"DATASET_SCOPE_LABEL":       2,
		"DATASET_SCOPE_METADATA":    3,
	}
)

func (x DatasetScope) Enum() *DatasetScope {
	p := new(DatasetScope)
	*p = x
	return p
}

func (x DatasetScope) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DatasetScope) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_dashboards_v1_common_observation_field_proto_enumTypes[0].Descriptor()
}

func (DatasetScope) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_dashboards_v1_common_observation_field_proto_enumTypes[0]
}

func (x DatasetScope) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DatasetScope.Descriptor instead.
func (DatasetScope) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_observation_field_proto_rawDescGZIP(), []int{0}
}

type ObservationField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keypath []*wrapperspb.StringValue `protobuf:"bytes,1,rep,name=keypath,proto3" json:"keypath,omitempty"`
	Scope   DatasetScope              `protobuf:"varint,2,opt,name=scope,proto3,enum=com.coralogixapis.dashboards.v1.common.DatasetScope" json:"scope,omitempty"`
}

func (x *ObservationField) Reset() {
	*x = ObservationField{}
	mi := &file_com_coralogixapis_dashboards_v1_common_observation_field_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ObservationField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObservationField) ProtoMessage() {}

func (x *ObservationField) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_common_observation_field_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObservationField.ProtoReflect.Descriptor instead.
func (*ObservationField) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_observation_field_proto_rawDescGZIP(), []int{0}
}

func (x *ObservationField) GetKeypath() []*wrapperspb.StringValue {
	if x != nil {
		return x.Keypath
	}
	return nil
}

func (x *ObservationField) GetScope() DatasetScope {
	if x != nil {
		return x.Scope
	}
	return DatasetScope_DATASET_SCOPE_UNSPECIFIED
}

var File_com_coralogixapis_dashboards_v1_common_observation_field_proto protoreflect.FileDescriptor

var file_com_coralogixapis_dashboards_v1_common_observation_field_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x01, 0x0a, 0x10, 0x4f, 0x62, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x36, 0x0a,
	0x07, 0x6b, 0x65, 0x79, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x6b, 0x65,
	0x79, 0x70, 0x61, 0x74, 0x68, 0x12, 0x4a, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x2a, 0x7f, 0x0a, 0x0c, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x70,
	0x65, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x41, 0x54, 0x41, 0x53, 0x45, 0x54, 0x5f, 0x53, 0x43, 0x4f,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x1b, 0x0a, 0x17, 0x44, 0x41, 0x54, 0x41, 0x53, 0x45, 0x54, 0x5f, 0x53, 0x43, 0x4f, 0x50,
	0x45, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0x01, 0x12, 0x17, 0x0a,
	0x13, 0x44, 0x41, 0x54, 0x41, 0x53, 0x45, 0x54, 0x5f, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f, 0x4c,
	0x41, 0x42, 0x45, 0x4c, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x44, 0x41, 0x54, 0x41, 0x53, 0x45,
	0x54, 0x5f, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x44, 0x41, 0x54, 0x41,
	0x10, 0x03, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_dashboards_v1_common_observation_field_proto_rawDescOnce sync.Once
	file_com_coralogixapis_dashboards_v1_common_observation_field_proto_rawDescData = file_com_coralogixapis_dashboards_v1_common_observation_field_proto_rawDesc
)

func file_com_coralogixapis_dashboards_v1_common_observation_field_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_dashboards_v1_common_observation_field_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_dashboards_v1_common_observation_field_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_dashboards_v1_common_observation_field_proto_rawDescData)
	})
	return file_com_coralogixapis_dashboards_v1_common_observation_field_proto_rawDescData
}

var file_com_coralogixapis_dashboards_v1_common_observation_field_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_dashboards_v1_common_observation_field_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_dashboards_v1_common_observation_field_proto_goTypes = []any{
	(DatasetScope)(0),              // 0: com.coralogixapis.dashboards.v1.common.DatasetScope
	(*ObservationField)(nil),       // 1: com.coralogixapis.dashboards.v1.common.ObservationField
	(*wrapperspb.StringValue)(nil), // 2: google.protobuf.StringValue
}
var file_com_coralogixapis_dashboards_v1_common_observation_field_proto_depIdxs = []int32{
	2, // 0: com.coralogixapis.dashboards.v1.common.ObservationField.keypath:type_name -> google.protobuf.StringValue
	0, // 1: com.coralogixapis.dashboards.v1.common.ObservationField.scope:type_name -> com.coralogixapis.dashboards.v1.common.DatasetScope
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_dashboards_v1_common_observation_field_proto_init() }
func file_com_coralogixapis_dashboards_v1_common_observation_field_proto_init() {
	if File_com_coralogixapis_dashboards_v1_common_observation_field_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_dashboards_v1_common_observation_field_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_dashboards_v1_common_observation_field_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_dashboards_v1_common_observation_field_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_dashboards_v1_common_observation_field_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_dashboards_v1_common_observation_field_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_dashboards_v1_common_observation_field_proto = out.File
	file_com_coralogixapis_dashboards_v1_common_observation_field_proto_rawDesc = nil
	file_com_coralogixapis_dashboards_v1_common_observation_field_proto_goTypes = nil
	file_com_coralogixapis_dashboards_v1_common_observation_field_proto_depIdxs = nil
}
