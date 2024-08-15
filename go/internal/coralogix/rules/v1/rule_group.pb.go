// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: com/coralogix/rules/v1/rule_group.proto

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

type RuleGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Creator       *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	Enabled       *wrapperspb.BoolValue   `protobuf:"bytes,7,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Hidden        *wrapperspb.BoolValue   `protobuf:"bytes,8,opt,name=hidden,proto3" json:"hidden,omitempty"`
	RuleMatchers  []*RuleMatcher          `protobuf:"bytes,9,rep,name=rule_matchers,json=ruleMatchers,proto3" json:"rule_matchers,omitempty"`
	RuleSubgroups []*RuleSubgroup         `protobuf:"bytes,10,rep,name=rule_subgroups,json=ruleSubgroups,proto3" json:"rule_subgroups,omitempty"`
	Order         *wrapperspb.UInt32Value `protobuf:"bytes,11,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *RuleGroup) Reset() {
	*x = RuleGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_rules_v1_rule_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuleGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuleGroup) ProtoMessage() {}

func (x *RuleGroup) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_rules_v1_rule_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TeamId) Reset() {
	*x = TeamId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_rules_v1_rule_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamId) ProtoMessage() {}

func (x *TeamId) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_rules_v1_rule_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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

var file_com_coralogix_rules_v1_rule_group_proto_rawDesc = []byte{
	0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x72, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x29, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x72, 0x75, 0x6c, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x75, 0x62, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x04, 0x0a, 0x09, 0x52, 0x75, 0x6c,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x34,
	0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x12, 0x32, 0x0a, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x12, 0x48, 0x0a, 0x0d, 0x72, 0x75, 0x6c, 0x65,
	0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x52, 0x0c, 0x72, 0x75, 0x6c, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x73, 0x12, 0x4b, 0x0a, 0x0e, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x75, 0x62, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x75, 0x62, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x0d, 0x72, 0x75, 0x6c, 0x65, 0x53, 0x75, 0x62, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12,
	0x32, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x22, 0x18, 0x0a, 0x06, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_rules_v1_rule_group_proto_rawDescOnce sync.Once
	file_com_coralogix_rules_v1_rule_group_proto_rawDescData = file_com_coralogix_rules_v1_rule_group_proto_rawDesc
)

func file_com_coralogix_rules_v1_rule_group_proto_rawDescGZIP() []byte {
	file_com_coralogix_rules_v1_rule_group_proto_rawDescOnce.Do(func() {
		file_com_coralogix_rules_v1_rule_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_rules_v1_rule_group_proto_rawDescData)
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
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_rules_v1_rule_group_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RuleGroup); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_com_coralogix_rules_v1_rule_group_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TeamId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_rules_v1_rule_group_proto_rawDesc,
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
	file_com_coralogix_rules_v1_rule_group_proto_rawDesc = nil
	file_com_coralogix_rules_v1_rule_group_proto_goTypes = nil
	file_com_coralogix_rules_v1_rule_group_proto_depIdxs = nil
}
