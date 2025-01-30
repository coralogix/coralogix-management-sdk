// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.2
// source: com/coralogixapis/apm/queries/k8s/v1/infra_filters.proto

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

type K8SFilter int32

const (
	K8SFilter_K8S_FILTER_UNSPECIFIED K8SFilter = 0
	K8SFilter_K8S_FILTER_NAMESPACE   K8SFilter = 1
	K8SFilter_K8S_FILTER_NODE        K8SFilter = 2
	K8SFilter_K8S_FILTER_APP         K8SFilter = 3
	K8SFilter_K8S_FILTER_POD         K8SFilter = 4
	K8SFilter_K8S_FILTER_CLUSTER     K8SFilter = 5
)

// Enum value maps for K8SFilter.
var (
	K8SFilter_name = map[int32]string{
		0: "K8S_FILTER_UNSPECIFIED",
		1: "K8S_FILTER_NAMESPACE",
		2: "K8S_FILTER_NODE",
		3: "K8S_FILTER_APP",
		4: "K8S_FILTER_POD",
		5: "K8S_FILTER_CLUSTER",
	}
	K8SFilter_value = map[string]int32{
		"K8S_FILTER_UNSPECIFIED": 0,
		"K8S_FILTER_NAMESPACE":   1,
		"K8S_FILTER_NODE":        2,
		"K8S_FILTER_APP":         3,
		"K8S_FILTER_POD":         4,
		"K8S_FILTER_CLUSTER":     5,
	}
)

func (x K8SFilter) Enum() *K8SFilter {
	p := new(K8SFilter)
	*p = x
	return p
}

func (x K8SFilter) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (K8SFilter) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_enumTypes[0].Descriptor()
}

func (K8SFilter) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_enumTypes[0]
}

func (x K8SFilter) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use K8SFilter.Descriptor instead.
func (K8SFilter) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_rawDescGZIP(), []int{0}
}

type Filter struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Kind          K8SFilter              `protobuf:"varint,1,opt,name=kind,proto3,enum=com.coralogixapis.apm.queries.k8s.v1.K8SFilter" json:"kind,omitempty"`
	Values        map[string]bool        `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"` // String represents the object's name, value represents if it is selected or not
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Filter) Reset() {
	*x = Filter{}
	mi := &file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_rawDescGZIP(), []int{0}
}

func (x *Filter) GetKind() K8SFilter {
	if x != nil {
		return x.Kind
	}
	return K8SFilter_K8S_FILTER_UNSPECIFIED
}

func (x *Filter) GetValues() map[string]bool {
	if x != nil {
		return x.Values
	}
	return nil
}

type SelectedFilter struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Kind          K8SFilter                 `protobuf:"varint,1,opt,name=kind,proto3,enum=com.coralogixapis.apm.queries.k8s.v1.K8SFilter" json:"kind,omitempty"`
	Values        []*wrapperspb.StringValue `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SelectedFilter) Reset() {
	*x = SelectedFilter{}
	mi := &file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SelectedFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectedFilter) ProtoMessage() {}

func (x *SelectedFilter) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectedFilter.ProtoReflect.Descriptor instead.
func (*SelectedFilter) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_rawDescGZIP(), []int{1}
}

func (x *SelectedFilter) GetKind() K8SFilter {
	if x != nil {
		return x.Kind
	}
	return K8SFilter_K8S_FILTER_UNSPECIFIED
}

func (x *SelectedFilter) GetValues() []*wrapperspb.StringValue {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto protoreflect.FileDescriptor

var file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_rawDesc = string([]byte{
	0x0a, 0x38, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2f,
	0x6b, 0x38, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70,
	0x6d, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xda, 0x01, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70,
	0x6d, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x4b, 0x38, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x12, 0x50, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73,
	0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8b, 0x01,
	0x0a, 0x0e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x43, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x6b,
	0x38, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x38, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x34, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2a, 0x96, 0x01, 0x0a, 0x09,
	0x4b, 0x38, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x16, 0x4b, 0x38, 0x53,
	0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x4b, 0x38, 0x53, 0x5f, 0x46, 0x49, 0x4c,
	0x54, 0x45, 0x52, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x53, 0x50, 0x41, 0x43, 0x45, 0x10, 0x01, 0x12,
	0x13, 0x0a, 0x0f, 0x4b, 0x38, 0x53, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x4e, 0x4f,
	0x44, 0x45, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x4b, 0x38, 0x53, 0x5f, 0x46, 0x49, 0x4c, 0x54,
	0x45, 0x52, 0x5f, 0x41, 0x50, 0x50, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x4b, 0x38, 0x53, 0x5f,
	0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x50, 0x4f, 0x44, 0x10, 0x04, 0x12, 0x16, 0x0a, 0x12,
	0x4b, 0x38, 0x53, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x43, 0x4c, 0x55, 0x53, 0x54,
	0x45, 0x52, 0x10, 0x05, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_rawDescOnce sync.Once
	file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_rawDescData []byte
)

func file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_rawDesc), len(file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_rawDesc)))
	})
	return file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_rawDescData
}

var file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_goTypes = []any{
	(K8SFilter)(0),                 // 0: com.coralogixapis.apm.queries.k8s.v1.K8sFilter
	(*Filter)(nil),                 // 1: com.coralogixapis.apm.queries.k8s.v1.Filter
	(*SelectedFilter)(nil),         // 2: com.coralogixapis.apm.queries.k8s.v1.SelectedFilter
	nil,                            // 3: com.coralogixapis.apm.queries.k8s.v1.Filter.ValuesEntry
	(*wrapperspb.StringValue)(nil), // 4: google.protobuf.StringValue
}
var file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_depIdxs = []int32{
	0, // 0: com.coralogixapis.apm.queries.k8s.v1.Filter.kind:type_name -> com.coralogixapis.apm.queries.k8s.v1.K8sFilter
	3, // 1: com.coralogixapis.apm.queries.k8s.v1.Filter.values:type_name -> com.coralogixapis.apm.queries.k8s.v1.Filter.ValuesEntry
	0, // 2: com.coralogixapis.apm.queries.k8s.v1.SelectedFilter.kind:type_name -> com.coralogixapis.apm.queries.k8s.v1.K8sFilter
	4, // 3: com.coralogixapis.apm.queries.k8s.v1.SelectedFilter.values:type_name -> google.protobuf.StringValue
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_init() }
func file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_init() {
	if File_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_rawDesc), len(file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto = out.File
	file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_goTypes = nil
	file_com_coralogixapis_apm_queries_k8s_v1_infra_filters_proto_depIdxs = nil
}
