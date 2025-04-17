// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogix/rules/v1/rule_group.proto

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

type RuleGroup struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Id            *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Creator       *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	Enabled       *wrapperspb.BoolValue   `protobuf:"bytes,7,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Hidden        *wrapperspb.BoolValue   `protobuf:"bytes,8,opt,name=hidden,proto3" json:"hidden,omitempty"`
	RuleMatchers  []*RuleMatcher          `protobuf:"bytes,9,rep,name=rule_matchers,json=ruleMatchers,proto3" json:"rule_matchers,omitempty"`
	RuleSubgroups []*RuleSubgroup         `protobuf:"bytes,10,rep,name=rule_subgroups,json=ruleSubgroups,proto3" json:"rule_subgroups,omitempty"`
	Order         *wrapperspb.UInt32Value `protobuf:"bytes,11,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RuleGroup) Reset() {
	*x = RuleGroup{}
	mi := &file_com_coralogix_rules_v1_rule_group_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RuleGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuleGroup) ProtoMessage() {}

func (x *RuleGroup) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_rules_v1_rule_group_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuleGroup.ProtoReflect.Descriptor instead.
func (*RuleGroup) Descriptor() ([]byte, []int) {
	return file_com_coralogix_rules_v1_rule_group_proto_rawDescGZIP(), []int{0}
}

func (x *RuleGroup) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *RuleGroup) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *RuleGroup) GetDescription() *wrapperspb.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *RuleGroup) GetCreator() *wrapperspb.StringValue {
	if x != nil {
		return x.Creator
	}
	return nil
}

func (x *RuleGroup) GetEnabled() *wrapperspb.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

func (x *RuleGroup) GetHidden() *wrapperspb.BoolValue {
	if x != nil {
		return x.Hidden
	}
	return nil
}

func (x *RuleGroup) GetRuleMatchers() []*RuleMatcher {
	if x != nil {
		return x.RuleMatchers
	}
	return nil
}

func (x *RuleGroup) GetRuleSubgroups() []*RuleSubgroup {
	if x != nil {
		return x.RuleSubgroups
	}
	return nil
}

func (x *RuleGroup) GetOrder() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Order
	}
	return nil
}

type TeamId struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TeamId) Reset() {
	*x = TeamId{}
	mi := &file_com_coralogix_rules_v1_rule_group_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TeamId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamId) ProtoMessage() {}

func (x *TeamId) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_rules_v1_rule_group_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamId.ProtoReflect.Descriptor instead.
func (*TeamId) Descriptor() ([]byte, []int) {
	return file_com_coralogix_rules_v1_rule_group_proto_rawDescGZIP(), []int{1}
}

func (x *TeamId) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_com_coralogix_rules_v1_rule_group_proto protoreflect.FileDescriptor

const file_com_coralogix_rules_v1_rule_group_proto_rawDesc = "" +
	"\n" +
	"'com/coralogix/rules/v1/rule_group.proto\x12\x16com.coralogix.rules.v1\x1a\x1egoogle/protobuf/wrappers.proto\x1a)com/coralogix/rules/v1/rule_matcher.proto\x1a*com/coralogix/rules/v1/rule_subgroup.proto\"\x98\x04\n" +
	"\tRuleGroup\x12,\n" +
	"\x02id\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueR\x02id\x120\n" +
	"\x04name\x18\x02 \x01(\v2\x1c.google.protobuf.StringValueR\x04name\x12>\n" +
	"\vdescription\x18\x03 \x01(\v2\x1c.google.protobuf.StringValueR\vdescription\x126\n" +
	"\acreator\x18\x06 \x01(\v2\x1c.google.protobuf.StringValueR\acreator\x124\n" +
	"\aenabled\x18\a \x01(\v2\x1a.google.protobuf.BoolValueR\aenabled\x122\n" +
	"\x06hidden\x18\b \x01(\v2\x1a.google.protobuf.BoolValueR\x06hidden\x12H\n" +
	"\rrule_matchers\x18\t \x03(\v2#.com.coralogix.rules.v1.RuleMatcherR\fruleMatchers\x12K\n" +
	"\x0erule_subgroups\x18\n" +
	" \x03(\v2$.com.coralogix.rules.v1.RuleSubgroupR\rruleSubgroups\x122\n" +
	"\x05order\x18\v \x01(\v2\x1c.google.protobuf.UInt32ValueR\x05order\"\x18\n" +
	"\x06TeamId\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02idb\x06proto3"

var (
	file_com_coralogix_rules_v1_rule_group_proto_rawDescOnce sync.Once
	file_com_coralogix_rules_v1_rule_group_proto_rawDescData []byte
)

func file_com_coralogix_rules_v1_rule_group_proto_rawDescGZIP() []byte {
	file_com_coralogix_rules_v1_rule_group_proto_rawDescOnce.Do(func() {
		file_com_coralogix_rules_v1_rule_group_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogix_rules_v1_rule_group_proto_rawDesc), len(file_com_coralogix_rules_v1_rule_group_proto_rawDesc)))
	})
	return file_com_coralogix_rules_v1_rule_group_proto_rawDescData
}

var file_com_coralogix_rules_v1_rule_group_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogix_rules_v1_rule_group_proto_goTypes = []any{
	(*RuleGroup)(nil),              // 0: com.coralogix.rules.v1.RuleGroup
	(*TeamId)(nil),                 // 1: com.coralogix.rules.v1.TeamId
	(*wrapperspb.StringValue)(nil), // 2: google.protobuf.StringValue
	(*wrapperspb.BoolValue)(nil),   // 3: google.protobuf.BoolValue
	(*RuleMatcher)(nil),            // 4: com.coralogix.rules.v1.RuleMatcher
	(*RuleSubgroup)(nil),           // 5: com.coralogix.rules.v1.RuleSubgroup
	(*wrapperspb.UInt32Value)(nil), // 6: google.protobuf.UInt32Value
}
var file_com_coralogix_rules_v1_rule_group_proto_depIdxs = []int32{
	2, // 0: com.coralogix.rules.v1.RuleGroup.id:type_name -> google.protobuf.StringValue
	2, // 1: com.coralogix.rules.v1.RuleGroup.name:type_name -> google.protobuf.StringValue
	2, // 2: com.coralogix.rules.v1.RuleGroup.description:type_name -> google.protobuf.StringValue
	2, // 3: com.coralogix.rules.v1.RuleGroup.creator:type_name -> google.protobuf.StringValue
	3, // 4: com.coralogix.rules.v1.RuleGroup.enabled:type_name -> google.protobuf.BoolValue
	3, // 5: com.coralogix.rules.v1.RuleGroup.hidden:type_name -> google.protobuf.BoolValue
	4, // 6: com.coralogix.rules.v1.RuleGroup.rule_matchers:type_name -> com.coralogix.rules.v1.RuleMatcher
	5, // 7: com.coralogix.rules.v1.RuleGroup.rule_subgroups:type_name -> com.coralogix.rules.v1.RuleSubgroup
	6, // 8: com.coralogix.rules.v1.RuleGroup.order:type_name -> google.protobuf.UInt32Value
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_com_coralogix_rules_v1_rule_group_proto_init() }
func file_com_coralogix_rules_v1_rule_group_proto_init() {
	if File_com_coralogix_rules_v1_rule_group_proto != nil {
		return
	}
	file_com_coralogix_rules_v1_rule_matcher_proto_init()
	file_com_coralogix_rules_v1_rule_subgroup_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogix_rules_v1_rule_group_proto_rawDesc), len(file_com_coralogix_rules_v1_rule_group_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_rules_v1_rule_group_proto_goTypes,
		DependencyIndexes: file_com_coralogix_rules_v1_rule_group_proto_depIdxs,
		MessageInfos:      file_com_coralogix_rules_v1_rule_group_proto_msgTypes,
	}.Build()
	File_com_coralogix_rules_v1_rule_group_proto = out.File
	file_com_coralogix_rules_v1_rule_group_proto_goTypes = nil
	file_com_coralogix_rules_v1_rule_group_proto_depIdxs = nil
}
