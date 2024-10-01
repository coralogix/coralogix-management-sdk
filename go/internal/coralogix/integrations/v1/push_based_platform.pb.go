// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: com/coralogix/integrations/v1/push_based_platform.proto

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

type PushBasedPlatform int32

const (
	PushBasedPlatform_UNDEFINED           PushBasedPlatform = 0
	PushBasedPlatform_PLATFORM_BITBUCKET  PushBasedPlatform = 1
	PushBasedPlatform_PLATFORM_GITHUB     PushBasedPlatform = 2
	PushBasedPlatform_PLATFORM_GITLAB     PushBasedPlatform = 3
	PushBasedPlatform_PLATFORM_AWS_SNS    PushBasedPlatform = 4
	PushBasedPlatform_PLATFORM_OPSGENIE   PushBasedPlatform = 5
	PushBasedPlatform_PLATFORM_PAGERDUTY  PushBasedPlatform = 6
	PushBasedPlatform_PLATFORM_PROMETHEUS PushBasedPlatform = 7
	PushBasedPlatform_PLATFORM_SLACK      PushBasedPlatform = 8
	PushBasedPlatform_PLATFORM_INTERCOM   PushBasedPlatform = 9
)

// Enum value maps for PushBasedPlatform.
var (
	PushBasedPlatform_name = map[int32]string{
		0: "UNDEFINED",
		1: "PLATFORM_BITBUCKET",
		2: "PLATFORM_GITHUB",
		3: "PLATFORM_GITLAB",
		4: "PLATFORM_AWS_SNS",
		5: "PLATFORM_OPSGENIE",
		6: "PLATFORM_PAGERDUTY",
		7: "PLATFORM_PROMETHEUS",
		8: "PLATFORM_SLACK",
		9: "PLATFORM_INTERCOM",
	}
	PushBasedPlatform_value = map[string]int32{
		"UNDEFINED":           0,
		"PLATFORM_BITBUCKET":  1,
		"PLATFORM_GITHUB":     2,
		"PLATFORM_GITLAB":     3,
		"PLATFORM_AWS_SNS":    4,
		"PLATFORM_OPSGENIE":   5,
		"PLATFORM_PAGERDUTY":  6,
		"PLATFORM_PROMETHEUS": 7,
		"PLATFORM_SLACK":      8,
		"PLATFORM_INTERCOM":   9,
	}
)

func (x PushBasedPlatform) Enum() *PushBasedPlatform {
	p := new(PushBasedPlatform)
	*p = x
	return p
}

func (x PushBasedPlatform) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PushBasedPlatform) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_integrations_v1_push_based_platform_proto_enumTypes[0].Descriptor()
}

func (PushBasedPlatform) Type() protoreflect.EnumType {
	return &file_com_coralogix_integrations_v1_push_based_platform_proto_enumTypes[0]
}

func (x PushBasedPlatform) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PushBasedPlatform.Descriptor instead.
func (PushBasedPlatform) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_integrations_v1_push_based_platform_proto_rawDescGZIP(), []int{0}
}

var File_com_coralogix_integrations_v1_push_based_platform_proto protoreflect.FileDescriptor

var file_com_coralogix_integrations_v1_push_based_platform_proto_rawDesc = []byte{
	0x0a, 0x37, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x75, 0x73, 0x68, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x64, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2a, 0xed, 0x01, 0x0a, 0x11, 0x50, 0x75, 0x73,
	0x68, 0x42, 0x61, 0x73, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x0d,
	0x0a, 0x09, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a,
	0x12, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x42, 0x49, 0x54, 0x42, 0x55, 0x43,
	0x4b, 0x45, 0x54, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52,
	0x4d, 0x5f, 0x47, 0x49, 0x54, 0x48, 0x55, 0x42, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x4c,
	0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x47, 0x49, 0x54, 0x4c, 0x41, 0x42, 0x10, 0x03, 0x12,
	0x14, 0x0a, 0x10, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x41, 0x57, 0x53, 0x5f,
	0x53, 0x4e, 0x53, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52,
	0x4d, 0x5f, 0x4f, 0x50, 0x53, 0x47, 0x45, 0x4e, 0x49, 0x45, 0x10, 0x05, 0x12, 0x16, 0x0a, 0x12,
	0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x50, 0x41, 0x47, 0x45, 0x52, 0x44, 0x55,
	0x54, 0x59, 0x10, 0x06, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d,
	0x5f, 0x50, 0x52, 0x4f, 0x4d, 0x45, 0x54, 0x48, 0x45, 0x55, 0x53, 0x10, 0x07, 0x12, 0x12, 0x0a,
	0x0e, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x53, 0x4c, 0x41, 0x43, 0x4b, 0x10,
	0x08, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x49, 0x4e,
	0x54, 0x45, 0x52, 0x43, 0x4f, 0x4d, 0x10, 0x09, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_integrations_v1_push_based_platform_proto_rawDescOnce sync.Once
	file_com_coralogix_integrations_v1_push_based_platform_proto_rawDescData = file_com_coralogix_integrations_v1_push_based_platform_proto_rawDesc
)

func file_com_coralogix_integrations_v1_push_based_platform_proto_rawDescGZIP() []byte {
	file_com_coralogix_integrations_v1_push_based_platform_proto_rawDescOnce.Do(func() {
		file_com_coralogix_integrations_v1_push_based_platform_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_integrations_v1_push_based_platform_proto_rawDescData)
	})
	return file_com_coralogix_integrations_v1_push_based_platform_proto_rawDescData
}

var file_com_coralogix_integrations_v1_push_based_platform_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogix_integrations_v1_push_based_platform_proto_goTypes = []any{
	(PushBasedPlatform)(0), // 0: com.coralogix.integrations.v1.PushBasedPlatform
}
var file_com_coralogix_integrations_v1_push_based_platform_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogix_integrations_v1_push_based_platform_proto_init() }
func file_com_coralogix_integrations_v1_push_based_platform_proto_init() {
	if File_com_coralogix_integrations_v1_push_based_platform_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_integrations_v1_push_based_platform_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_integrations_v1_push_based_platform_proto_goTypes,
		DependencyIndexes: file_com_coralogix_integrations_v1_push_based_platform_proto_depIdxs,
		EnumInfos:         file_com_coralogix_integrations_v1_push_based_platform_proto_enumTypes,
	}.Build()
	File_com_coralogix_integrations_v1_push_based_platform_proto = out.File
	file_com_coralogix_integrations_v1_push_based_platform_proto_rawDesc = nil
	file_com_coralogix_integrations_v1_push_based_platform_proto_goTypes = nil
	file_com_coralogix_integrations_v1_push_based_platform_proto_depIdxs = nil
}
