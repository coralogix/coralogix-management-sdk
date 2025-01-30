// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.2
// source: com/coralogixapis/apm/requirements/v2/features.proto

package v2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Feature int32

const (
	Feature_FEATURE_UNSPECIFIED   Feature = 0
	Feature_FEATURE_K8S_DASHBOARD Feature = 1
)

// Enum value maps for Feature.
var (
	Feature_name = map[int32]string{
		0: "FEATURE_UNSPECIFIED",
		1: "FEATURE_K8S_DASHBOARD",
	}
	Feature_value = map[string]int32{
		"FEATURE_UNSPECIFIED":   0,
		"FEATURE_K8S_DASHBOARD": 1,
	}
)

func (x Feature) Enum() *Feature {
	p := new(Feature)
	*p = x
	return p
}

func (x Feature) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Feature) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_apm_requirements_v2_features_proto_enumTypes[0].Descriptor()
}

func (Feature) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_apm_requirements_v2_features_proto_enumTypes[0]
}

func (x Feature) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Feature.Descriptor instead.
func (Feature) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_requirements_v2_features_proto_rawDescGZIP(), []int{0}
}

var File_com_coralogixapis_apm_requirements_v2_features_proto protoreflect.FileDescriptor

var file_com_coralogixapis_apm_requirements_v2_features_proto_rawDesc = string([]byte{
	0x0a, 0x34, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x32, 0x2a, 0x3d, 0x0a,
	0x07, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x45, 0x41, 0x54,
	0x55, 0x52, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x19, 0x0a, 0x15, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x4b, 0x38, 0x53,
	0x5f, 0x44, 0x41, 0x53, 0x48, 0x42, 0x4f, 0x41, 0x52, 0x44, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_com_coralogixapis_apm_requirements_v2_features_proto_rawDescOnce sync.Once
	file_com_coralogixapis_apm_requirements_v2_features_proto_rawDescData []byte
)

func file_com_coralogixapis_apm_requirements_v2_features_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_apm_requirements_v2_features_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_apm_requirements_v2_features_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_apm_requirements_v2_features_proto_rawDesc), len(file_com_coralogixapis_apm_requirements_v2_features_proto_rawDesc)))
	})
	return file_com_coralogixapis_apm_requirements_v2_features_proto_rawDescData
}

var file_com_coralogixapis_apm_requirements_v2_features_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_apm_requirements_v2_features_proto_goTypes = []any{
	(Feature)(0), // 0: com.coralogixapis.apm.requirements.v2.Feature
}
var file_com_coralogixapis_apm_requirements_v2_features_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_apm_requirements_v2_features_proto_init() }
func file_com_coralogixapis_apm_requirements_v2_features_proto_init() {
	if File_com_coralogixapis_apm_requirements_v2_features_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_apm_requirements_v2_features_proto_rawDesc), len(file_com_coralogixapis_apm_requirements_v2_features_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_apm_requirements_v2_features_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_apm_requirements_v2_features_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_apm_requirements_v2_features_proto_enumTypes,
	}.Build()
	File_com_coralogixapis_apm_requirements_v2_features_proto = out.File
	file_com_coralogixapis_apm_requirements_v2_features_proto_goTypes = nil
	file_com_coralogixapis_apm_requirements_v2_features_proto_depIdxs = nil
}
