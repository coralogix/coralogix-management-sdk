// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.26.1
// source: com/coralogix/rules/v1/rule_matcher.proto

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

type SeverityConstraint_Value int32

const (
	SeverityConstraint_VALUE_DEBUG_OR_UNSPECIFIED SeverityConstraint_Value = 0
	SeverityConstraint_VALUE_VERBOSE              SeverityConstraint_Value = 1
	SeverityConstraint_VALUE_INFO                 SeverityConstraint_Value = 2
	SeverityConstraint_VALUE_WARNING              SeverityConstraint_Value = 3
	SeverityConstraint_VALUE_ERROR                SeverityConstraint_Value = 4
	SeverityConstraint_VALUE_CRITICAL             SeverityConstraint_Value = 5
)

// Enum value maps for SeverityConstraint_Value.
var (
	SeverityConstraint_Value_name = map[int32]string{
		0: "VALUE_DEBUG_OR_UNSPECIFIED",
		1: "VALUE_VERBOSE",
		2: "VALUE_INFO",
		3: "VALUE_WARNING",
		4: "VALUE_ERROR",
		5: "VALUE_CRITICAL",
	}
	SeverityConstraint_Value_value = map[string]int32{
		"VALUE_DEBUG_OR_UNSPECIFIED": 0,
		"VALUE_VERBOSE":              1,
		"VALUE_INFO":                 2,
		"VALUE_WARNING":              3,
		"VALUE_ERROR":                4,
		"VALUE_CRITICAL":             5,
	}
)

func (x SeverityConstraint_Value) Enum() *SeverityConstraint_Value {
	p := new(SeverityConstraint_Value)
	*p = x
	return p
}

func (x SeverityConstraint_Value) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SeverityConstraint_Value) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_rules_v1_rule_matcher_proto_enumTypes[0].Descriptor()
}

func (SeverityConstraint_Value) Type() protoreflect.EnumType {
	return &file_com_coralogix_rules_v1_rule_matcher_proto_enumTypes[0]
}

func (x SeverityConstraint_Value) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SeverityConstraint_Value.Descriptor instead.
func (SeverityConstraint_Value) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_rules_v1_rule_matcher_proto_rawDescGZIP(), []int{3, 0}
}

type RuleMatcher struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Constraint:
	//
	//	*RuleMatcher_ApplicationName
	//	*RuleMatcher_SubsystemName
	//	*RuleMatcher_Severity
	Constraint    isRuleMatcher_Constraint `protobuf_oneof:"constraint"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RuleMatcher) Reset() {
	*x = RuleMatcher{}
	mi := &file_com_coralogix_rules_v1_rule_matcher_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RuleMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuleMatcher) ProtoMessage() {}

func (x *RuleMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_rules_v1_rule_matcher_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuleMatcher.ProtoReflect.Descriptor instead.
func (*RuleMatcher) Descriptor() ([]byte, []int) {
	return file_com_coralogix_rules_v1_rule_matcher_proto_rawDescGZIP(), []int{0}
}

func (x *RuleMatcher) GetConstraint() isRuleMatcher_Constraint {
	if x != nil {
		return x.Constraint
	}
	return nil
}

func (x *RuleMatcher) GetApplicationName() *ApplicationNameConstraint {
	if x != nil {
		if x, ok := x.Constraint.(*RuleMatcher_ApplicationName); ok {
			return x.ApplicationName
		}
	}
	return nil
}

func (x *RuleMatcher) GetSubsystemName() *SubsystemNameConstraint {
	if x != nil {
		if x, ok := x.Constraint.(*RuleMatcher_SubsystemName); ok {
			return x.SubsystemName
		}
	}
	return nil
}

func (x *RuleMatcher) GetSeverity() *SeverityConstraint {
	if x != nil {
		if x, ok := x.Constraint.(*RuleMatcher_Severity); ok {
			return x.Severity
		}
	}
	return nil
}

type isRuleMatcher_Constraint interface {
	isRuleMatcher_Constraint()
}

type RuleMatcher_ApplicationName struct {
	ApplicationName *ApplicationNameConstraint `protobuf:"bytes,1,opt,name=application_name,json=applicationName,proto3,oneof"`
}

type RuleMatcher_SubsystemName struct {
	SubsystemName *SubsystemNameConstraint `protobuf:"bytes,2,opt,name=subsystem_name,json=subsystemName,proto3,oneof"`
}

type RuleMatcher_Severity struct {
	Severity *SeverityConstraint `protobuf:"bytes,9,opt,name=severity,proto3,oneof"`
}

func (*RuleMatcher_ApplicationName) isRuleMatcher_Constraint() {}

func (*RuleMatcher_SubsystemName) isRuleMatcher_Constraint() {}

func (*RuleMatcher_Severity) isRuleMatcher_Constraint() {}

type ApplicationNameConstraint struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Value         *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ApplicationNameConstraint) Reset() {
	*x = ApplicationNameConstraint{}
	mi := &file_com_coralogix_rules_v1_rule_matcher_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApplicationNameConstraint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationNameConstraint) ProtoMessage() {}

func (x *ApplicationNameConstraint) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_rules_v1_rule_matcher_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationNameConstraint.ProtoReflect.Descriptor instead.
func (*ApplicationNameConstraint) Descriptor() ([]byte, []int) {
	return file_com_coralogix_rules_v1_rule_matcher_proto_rawDescGZIP(), []int{1}
}

func (x *ApplicationNameConstraint) GetValue() *wrapperspb.StringValue {
	if x != nil {
		return x.Value
	}
	return nil
}

type SubsystemNameConstraint struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Value         *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubsystemNameConstraint) Reset() {
	*x = SubsystemNameConstraint{}
	mi := &file_com_coralogix_rules_v1_rule_matcher_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubsystemNameConstraint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubsystemNameConstraint) ProtoMessage() {}

func (x *SubsystemNameConstraint) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_rules_v1_rule_matcher_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubsystemNameConstraint.ProtoReflect.Descriptor instead.
func (*SubsystemNameConstraint) Descriptor() ([]byte, []int) {
	return file_com_coralogix_rules_v1_rule_matcher_proto_rawDescGZIP(), []int{2}
}

func (x *SubsystemNameConstraint) GetValue() *wrapperspb.StringValue {
	if x != nil {
		return x.Value
	}
	return nil
}

type SeverityConstraint struct {
	state         protoimpl.MessageState   `protogen:"open.v1"`
	Value         SeverityConstraint_Value `protobuf:"varint,1,opt,name=value,proto3,enum=com.coralogix.rules.v1.SeverityConstraint_Value" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SeverityConstraint) Reset() {
	*x = SeverityConstraint{}
	mi := &file_com_coralogix_rules_v1_rule_matcher_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SeverityConstraint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeverityConstraint) ProtoMessage() {}

func (x *SeverityConstraint) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_rules_v1_rule_matcher_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeverityConstraint.ProtoReflect.Descriptor instead.
func (*SeverityConstraint) Descriptor() ([]byte, []int) {
	return file_com_coralogix_rules_v1_rule_matcher_proto_rawDescGZIP(), []int{3}
}

func (x *SeverityConstraint) GetValue() SeverityConstraint_Value {
	if x != nil {
		return x.Value
	}
	return SeverityConstraint_VALUE_DEBUG_OR_UNSPECIFIED
}

var File_com_coralogix_rules_v1_rule_matcher_proto protoreflect.FileDescriptor

const file_com_coralogix_rules_v1_rule_matcher_proto_rawDesc = "" +
	"\n" +
	")com/coralogix/rules/v1/rule_matcher.proto\x12\x16com.coralogix.rules.v1\x1a\x1egoogle/protobuf/wrappers.proto\"\x9f\x02\n" +
	"\vRuleMatcher\x12^\n" +
	"\x10application_name\x18\x01 \x01(\v21.com.coralogix.rules.v1.ApplicationNameConstraintH\x00R\x0fapplicationName\x12X\n" +
	"\x0esubsystem_name\x18\x02 \x01(\v2/.com.coralogix.rules.v1.SubsystemNameConstraintH\x00R\rsubsystemName\x12H\n" +
	"\bseverity\x18\t \x01(\v2*.com.coralogix.rules.v1.SeverityConstraintH\x00R\bseverityB\f\n" +
	"\n" +
	"constraint\"O\n" +
	"\x19ApplicationNameConstraint\x122\n" +
	"\x05value\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueR\x05value\"M\n" +
	"\x17SubsystemNameConstraint\x122\n" +
	"\x05value\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueR\x05value\"\xe1\x01\n" +
	"\x12SeverityConstraint\x12F\n" +
	"\x05value\x18\x01 \x01(\x0e20.com.coralogix.rules.v1.SeverityConstraint.ValueR\x05value\"\x82\x01\n" +
	"\x05Value\x12\x1e\n" +
	"\x1aVALUE_DEBUG_OR_UNSPECIFIED\x10\x00\x12\x11\n" +
	"\rVALUE_VERBOSE\x10\x01\x12\x0e\n" +
	"\n" +
	"VALUE_INFO\x10\x02\x12\x11\n" +
	"\rVALUE_WARNING\x10\x03\x12\x0f\n" +
	"\vVALUE_ERROR\x10\x04\x12\x12\n" +
	"\x0eVALUE_CRITICAL\x10\x05b\x06proto3"

var (
	file_com_coralogix_rules_v1_rule_matcher_proto_rawDescOnce sync.Once
	file_com_coralogix_rules_v1_rule_matcher_proto_rawDescData []byte
)

func file_com_coralogix_rules_v1_rule_matcher_proto_rawDescGZIP() []byte {
	file_com_coralogix_rules_v1_rule_matcher_proto_rawDescOnce.Do(func() {
		file_com_coralogix_rules_v1_rule_matcher_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogix_rules_v1_rule_matcher_proto_rawDesc), len(file_com_coralogix_rules_v1_rule_matcher_proto_rawDesc)))
	})
	return file_com_coralogix_rules_v1_rule_matcher_proto_rawDescData
}

var file_com_coralogix_rules_v1_rule_matcher_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogix_rules_v1_rule_matcher_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_com_coralogix_rules_v1_rule_matcher_proto_goTypes = []any{
	(SeverityConstraint_Value)(0),     // 0: com.coralogix.rules.v1.SeverityConstraint.Value
	(*RuleMatcher)(nil),               // 1: com.coralogix.rules.v1.RuleMatcher
	(*ApplicationNameConstraint)(nil), // 2: com.coralogix.rules.v1.ApplicationNameConstraint
	(*SubsystemNameConstraint)(nil),   // 3: com.coralogix.rules.v1.SubsystemNameConstraint
	(*SeverityConstraint)(nil),        // 4: com.coralogix.rules.v1.SeverityConstraint
	(*wrapperspb.StringValue)(nil),    // 5: google.protobuf.StringValue
}
var file_com_coralogix_rules_v1_rule_matcher_proto_depIdxs = []int32{
	2, // 0: com.coralogix.rules.v1.RuleMatcher.application_name:type_name -> com.coralogix.rules.v1.ApplicationNameConstraint
	3, // 1: com.coralogix.rules.v1.RuleMatcher.subsystem_name:type_name -> com.coralogix.rules.v1.SubsystemNameConstraint
	4, // 2: com.coralogix.rules.v1.RuleMatcher.severity:type_name -> com.coralogix.rules.v1.SeverityConstraint
	5, // 3: com.coralogix.rules.v1.ApplicationNameConstraint.value:type_name -> google.protobuf.StringValue
	5, // 4: com.coralogix.rules.v1.SubsystemNameConstraint.value:type_name -> google.protobuf.StringValue
	0, // 5: com.coralogix.rules.v1.SeverityConstraint.value:type_name -> com.coralogix.rules.v1.SeverityConstraint.Value
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_com_coralogix_rules_v1_rule_matcher_proto_init() }
func file_com_coralogix_rules_v1_rule_matcher_proto_init() {
	if File_com_coralogix_rules_v1_rule_matcher_proto != nil {
		return
	}
	file_com_coralogix_rules_v1_rule_matcher_proto_msgTypes[0].OneofWrappers = []any{
		(*RuleMatcher_ApplicationName)(nil),
		(*RuleMatcher_SubsystemName)(nil),
		(*RuleMatcher_Severity)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogix_rules_v1_rule_matcher_proto_rawDesc), len(file_com_coralogix_rules_v1_rule_matcher_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_rules_v1_rule_matcher_proto_goTypes,
		DependencyIndexes: file_com_coralogix_rules_v1_rule_matcher_proto_depIdxs,
		EnumInfos:         file_com_coralogix_rules_v1_rule_matcher_proto_enumTypes,
		MessageInfos:      file_com_coralogix_rules_v1_rule_matcher_proto_msgTypes,
	}.Build()
	File_com_coralogix_rules_v1_rule_matcher_proto = out.File
	file_com_coralogix_rules_v1_rule_matcher_proto_goTypes = nil
	file_com_coralogix_rules_v1_rule_matcher_proto_depIdxs = nil
}
