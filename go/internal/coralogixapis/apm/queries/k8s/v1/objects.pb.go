// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: com/coralogixapis/apm/queries/k8s/v1/objects.proto

package v1

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

type K8SObject int32

const (
	K8SObject_K8S_OBJECT_UNSPECIFIED K8SObject = 0
	K8SObject_K8S_OBJECT_POD         K8SObject = 1
	K8SObject_K8S_OBJECT_NODE        K8SObject = 2
	K8SObject_K8S_OBJECT_APP         K8SObject = 3
)

// Enum value maps for K8SObject.
var (
	K8SObject_name = map[int32]string{
		0: "K8S_OBJECT_UNSPECIFIED",
		1: "K8S_OBJECT_POD",
		2: "K8S_OBJECT_NODE",
		3: "K8S_OBJECT_APP",
	}
	K8SObject_value = map[string]int32{
		"K8S_OBJECT_UNSPECIFIED": 0,
		"K8S_OBJECT_POD":         1,
		"K8S_OBJECT_NODE":        2,
		"K8S_OBJECT_APP":         3,
	}
)

func (x K8SObject) Enum() *K8SObject {
	p := new(K8SObject)
	*p = x
	return p
}

func (x K8SObject) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (K8SObject) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_apm_queries_k8s_v1_objects_proto_enumTypes[0].Descriptor()
}

func (K8SObject) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_apm_queries_k8s_v1_objects_proto_enumTypes[0]
}

func (x K8SObject) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use K8SObject.Descriptor instead.
func (K8SObject) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_queries_k8s_v1_objects_proto_rawDescGZIP(), []int{0}
}

var File_com_coralogixapis_apm_queries_k8s_v1_objects_proto protoreflect.FileDescriptor

var file_com_coralogixapis_apm_queries_k8s_v1_objects_proto_rawDesc = []byte{
	0x0a, 0x32, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2f,
	0x6b, 0x38, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x71, 0x75, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x2a, 0x64, 0x0a, 0x09, 0x4b, 0x38,
	0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x16, 0x4b, 0x38, 0x53, 0x5f, 0x4f,
	0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x4b, 0x38, 0x53, 0x5f, 0x4f, 0x42, 0x4a, 0x45, 0x43,
	0x54, 0x5f, 0x50, 0x4f, 0x44, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x4b, 0x38, 0x53, 0x5f, 0x4f,
	0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x4e, 0x4f, 0x44, 0x45, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e,
	0x4b, 0x38, 0x53, 0x5f, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x41, 0x50, 0x50, 0x10, 0x03,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_apm_queries_k8s_v1_objects_proto_rawDescOnce sync.Once
	file_com_coralogixapis_apm_queries_k8s_v1_objects_proto_rawDescData = file_com_coralogixapis_apm_queries_k8s_v1_objects_proto_rawDesc
)

func file_com_coralogixapis_apm_queries_k8s_v1_objects_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_apm_queries_k8s_v1_objects_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_apm_queries_k8s_v1_objects_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_apm_queries_k8s_v1_objects_proto_rawDescData)
	})
	return file_com_coralogixapis_apm_queries_k8s_v1_objects_proto_rawDescData
}

var file_com_coralogixapis_apm_queries_k8s_v1_objects_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_apm_queries_k8s_v1_objects_proto_goTypes = []any{
	(K8SObject)(0), // 0: com.coralogixapis.apm.queries.k8s.v1.K8sObject
}
var file_com_coralogixapis_apm_queries_k8s_v1_objects_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_apm_queries_k8s_v1_objects_proto_init() }
func file_com_coralogixapis_apm_queries_k8s_v1_objects_proto_init() {
	if File_com_coralogixapis_apm_queries_k8s_v1_objects_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_apm_queries_k8s_v1_objects_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_apm_queries_k8s_v1_objects_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_apm_queries_k8s_v1_objects_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_apm_queries_k8s_v1_objects_proto_enumTypes,
	}.Build()
	File_com_coralogixapis_apm_queries_k8s_v1_objects_proto = out.File
	file_com_coralogixapis_apm_queries_k8s_v1_objects_proto_rawDesc = nil
	file_com_coralogixapis_apm_queries_k8s_v1_objects_proto_goTypes = nil
	file_com_coralogixapis_apm_queries_k8s_v1_objects_proto_depIdxs = nil
}
