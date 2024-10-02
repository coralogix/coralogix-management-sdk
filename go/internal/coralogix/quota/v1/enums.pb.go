// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: com/coralogix/quota/v1/enums.proto

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

type Priority int32

const (
	Priority_PRIORITY_TYPE_UNSPECIFIED Priority = 0
	Priority_PRIORITY_TYPE_BLOCK       Priority = 1
	Priority_PRIORITY_TYPE_LOW         Priority = 2
	Priority_PRIORITY_TYPE_MEDIUM      Priority = 3
	Priority_PRIORITY_TYPE_HIGH        Priority = 4
)

// Enum value maps for Priority.
var (
	Priority_name = map[int32]string{
		0: "PRIORITY_TYPE_UNSPECIFIED",
		1: "PRIORITY_TYPE_BLOCK",
		2: "PRIORITY_TYPE_LOW",
		3: "PRIORITY_TYPE_MEDIUM",
		4: "PRIORITY_TYPE_HIGH",
	}
	Priority_value = map[string]int32{
		"PRIORITY_TYPE_UNSPECIFIED": 0,
		"PRIORITY_TYPE_BLOCK":       1,
		"PRIORITY_TYPE_LOW":         2,
		"PRIORITY_TYPE_MEDIUM":      3,
		"PRIORITY_TYPE_HIGH":        4,
	}
)

func (x Priority) Enum() *Priority {
	p := new(Priority)
	*p = x
	return p
}

func (x Priority) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Priority) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_quota_v1_enums_proto_enumTypes[0].Descriptor()
}

func (Priority) Type() protoreflect.EnumType {
	return &file_com_coralogix_quota_v1_enums_proto_enumTypes[0]
}

func (x Priority) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Priority.Descriptor instead.
func (Priority) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_quota_v1_enums_proto_rawDescGZIP(), []int{0}
}

type Severity int32

const (
	Severity_SEVERITY_UNSPECIFIED Severity = 0
	Severity_SEVERITY_DEBUG       Severity = 1
	Severity_SEVERITY_VERBOSE     Severity = 2
	Severity_SEVERITY_INFO        Severity = 3
	Severity_SEVERITY_WARNING     Severity = 4
	Severity_SEVERITY_ERROR       Severity = 5
	Severity_SEVERITY_CRITICAL    Severity = 6
)

// Enum value maps for Severity.
var (
	Severity_name = map[int32]string{
		0: "SEVERITY_UNSPECIFIED",
		1: "SEVERITY_DEBUG",
		2: "SEVERITY_VERBOSE",
		3: "SEVERITY_INFO",
		4: "SEVERITY_WARNING",
		5: "SEVERITY_ERROR",
		6: "SEVERITY_CRITICAL",
	}
	Severity_value = map[string]int32{
		"SEVERITY_UNSPECIFIED": 0,
		"SEVERITY_DEBUG":       1,
		"SEVERITY_VERBOSE":     2,
		"SEVERITY_INFO":        3,
		"SEVERITY_WARNING":     4,
		"SEVERITY_ERROR":       5,
		"SEVERITY_CRITICAL":    6,
	}
)

func (x Severity) Enum() *Severity {
	p := new(Severity)
	*p = x
	return p
}

func (x Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_quota_v1_enums_proto_enumTypes[1].Descriptor()
}

func (Severity) Type() protoreflect.EnumType {
	return &file_com_coralogix_quota_v1_enums_proto_enumTypes[1]
}

func (x Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Severity.Descriptor instead.
func (Severity) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_quota_v1_enums_proto_rawDescGZIP(), []int{1}
}

type SourceType int32

const (
	SourceType_SOURCE_TYPE_UNSPECIFIED SourceType = 0
	SourceType_SOURCE_TYPE_LOGS        SourceType = 1
	SourceType_SOURCE_TYPE_SPANS       SourceType = 2
)

// Enum value maps for SourceType.
var (
	SourceType_name = map[int32]string{
		0: "SOURCE_TYPE_UNSPECIFIED",
		1: "SOURCE_TYPE_LOGS",
		2: "SOURCE_TYPE_SPANS",
	}
	SourceType_value = map[string]int32{
		"SOURCE_TYPE_UNSPECIFIED": 0,
		"SOURCE_TYPE_LOGS":        1,
		"SOURCE_TYPE_SPANS":       2,
	}
)

func (x SourceType) Enum() *SourceType {
	p := new(SourceType)
	*p = x
	return p
}

func (x SourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_quota_v1_enums_proto_enumTypes[2].Descriptor()
}

func (SourceType) Type() protoreflect.EnumType {
	return &file_com_coralogix_quota_v1_enums_proto_enumTypes[2]
}

func (x SourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SourceType.Descriptor instead.
func (SourceType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_quota_v1_enums_proto_rawDescGZIP(), []int{2}
}

type RuleTypeId int32

const (
	RuleTypeId_RULE_TYPE_ID_UNSPECIFIED RuleTypeId = 0
	RuleTypeId_RULE_TYPE_ID_IS          RuleTypeId = 2
	RuleTypeId_RULE_TYPE_ID_IS_NOT      RuleTypeId = 3
	RuleTypeId_RULE_TYPE_ID_START_WITH  RuleTypeId = 4
	RuleTypeId_RULE_TYPE_ID_INCLUDES    RuleTypeId = 6
)

// Enum value maps for RuleTypeId.
var (
	RuleTypeId_name = map[int32]string{
		0: "RULE_TYPE_ID_UNSPECIFIED",
		2: "RULE_TYPE_ID_IS",
		3: "RULE_TYPE_ID_IS_NOT",
		4: "RULE_TYPE_ID_START_WITH",
		6: "RULE_TYPE_ID_INCLUDES",
	}
	RuleTypeId_value = map[string]int32{
		"RULE_TYPE_ID_UNSPECIFIED": 0,
		"RULE_TYPE_ID_IS":          2,
		"RULE_TYPE_ID_IS_NOT":      3,
		"RULE_TYPE_ID_START_WITH":  4,
		"RULE_TYPE_ID_INCLUDES":    6,
	}
)

func (x RuleTypeId) Enum() *RuleTypeId {
	p := new(RuleTypeId)
	*p = x
	return p
}

func (x RuleTypeId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RuleTypeId) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_quota_v1_enums_proto_enumTypes[3].Descriptor()
}

func (RuleTypeId) Type() protoreflect.EnumType {
	return &file_com_coralogix_quota_v1_enums_proto_enumTypes[3]
}

func (x RuleTypeId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RuleTypeId.Descriptor instead.
func (RuleTypeId) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_quota_v1_enums_proto_rawDescGZIP(), []int{3}
}

var File_com_coralogix_quota_v1_enums_proto protoreflect.FileDescriptor

var file_com_coralogix_quota_v1_enums_proto_rawDesc = []byte{
	0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2a, 0x8b, 0x01, 0x0a,
	0x08, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x19, 0x50, 0x52, 0x49,
	0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x52, 0x49, 0x4f,
	0x52, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x10,
	0x01, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4c, 0x4f, 0x57, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x52, 0x49, 0x4f,
	0x52, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x55, 0x4d,
	0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x10, 0x04, 0x2a, 0xa2, 0x01, 0x0a, 0x08, 0x53,
	0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x45, 0x56, 0x45, 0x52,
	0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x44, 0x45,
	0x42, 0x55, 0x47, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54,
	0x59, 0x5f, 0x56, 0x45, 0x52, 0x42, 0x4f, 0x53, 0x45, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x53,
	0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x03, 0x12, 0x14,
	0x0a, 0x10, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x57, 0x41, 0x52, 0x4e, 0x49,
	0x4e, 0x47, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x05, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x45, 0x56, 0x45,
	0x52, 0x49, 0x54, 0x59, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x06, 0x2a,
	0x56, 0x0a, 0x0a, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a,
	0x17, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x4f,
	0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x4f, 0x47, 0x53, 0x10, 0x01,
	0x12, 0x15, 0x0a, 0x11, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x53, 0x50, 0x41, 0x4e, 0x53, 0x10, 0x02, 0x2a, 0xa2, 0x01, 0x0a, 0x0a, 0x52, 0x75, 0x6c, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x18, 0x52, 0x55, 0x4c, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x49, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x55, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x49, 0x44, 0x5f, 0x49, 0x53, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x55, 0x4c,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x44, 0x5f, 0x49, 0x53, 0x5f, 0x4e, 0x4f, 0x54,
	0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x52, 0x55, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x49, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x10, 0x04, 0x12,
	0x19, 0x0a, 0x15, 0x52, 0x55, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x44, 0x5f,
	0x49, 0x4e, 0x43, 0x4c, 0x55, 0x44, 0x45, 0x53, 0x10, 0x06, 0x22, 0x04, 0x08, 0x01, 0x10, 0x01,
	0x22, 0x04, 0x08, 0x05, 0x10, 0x05, 0x22, 0x04, 0x08, 0x07, 0x10, 0x07, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_quota_v1_enums_proto_rawDescOnce sync.Once
	file_com_coralogix_quota_v1_enums_proto_rawDescData = file_com_coralogix_quota_v1_enums_proto_rawDesc
)

func file_com_coralogix_quota_v1_enums_proto_rawDescGZIP() []byte {
	file_com_coralogix_quota_v1_enums_proto_rawDescOnce.Do(func() {
		file_com_coralogix_quota_v1_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_quota_v1_enums_proto_rawDescData)
	})
	return file_com_coralogix_quota_v1_enums_proto_rawDescData
}

var file_com_coralogix_quota_v1_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_com_coralogix_quota_v1_enums_proto_goTypes = []any{
	(Priority)(0),   // 0: com.coralogix.quota.v1.Priority
	(Severity)(0),   // 1: com.coralogix.quota.v1.Severity
	(SourceType)(0), // 2: com.coralogix.quota.v1.SourceType
	(RuleTypeId)(0), // 3: com.coralogix.quota.v1.RuleTypeId
}
var file_com_coralogix_quota_v1_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogix_quota_v1_enums_proto_init() }
func file_com_coralogix_quota_v1_enums_proto_init() {
	if File_com_coralogix_quota_v1_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_quota_v1_enums_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_quota_v1_enums_proto_goTypes,
		DependencyIndexes: file_com_coralogix_quota_v1_enums_proto_depIdxs,
		EnumInfos:         file_com_coralogix_quota_v1_enums_proto_enumTypes,
	}.Build()
	File_com_coralogix_quota_v1_enums_proto = out.File
	file_com_coralogix_quota_v1_enums_proto_rawDesc = nil
	file_com_coralogix_quota_v1_enums_proto_goTypes = nil
	file_com_coralogix_quota_v1_enums_proto_depIdxs = nil
}
