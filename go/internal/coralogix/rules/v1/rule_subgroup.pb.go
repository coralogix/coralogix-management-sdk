// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: com/coralogix/rules/v1/rule_subgroup.proto

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

type RuleSubgroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Rules   []*Rule                 `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty"`
	Enabled *wrapperspb.BoolValue   `protobuf:"bytes,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Order   *wrapperspb.UInt32Value `protobuf:"bytes,4,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *RuleSubgroup) Reset() {
	*x = RuleSubgroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_rules_v1_rule_subgroup_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuleSubgroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuleSubgroup) ProtoMessage() {}

func (x *RuleSubgroup) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_rules_v1_rule_subgroup_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuleSubgroup.ProtoReflect.Descriptor instead.
func (*RuleSubgroup) Descriptor() ([]byte, []int) {
	return file_com_coralogix_rules_v1_rule_subgroup_proto_rawDescGZIP(), []int{0}
}

func (x *RuleSubgroup) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *RuleSubgroup) GetRules() []*Rule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *RuleSubgroup) GetEnabled() *wrapperspb.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

func (x *RuleSubgroup) GetOrder() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Order
	}
	return nil
}

var File_com_coralogix_rules_v1_rule_subgroup_proto protoreflect.FileDescriptor

var file_com_coralogix_rules_v1_rule_subgroup_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x72, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x75,
	0x62, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x72, 0x75, 0x6c, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x01, 0x0a, 0x0c, 0x52, 0x75, 0x6c, 0x65,
	0x53, 0x75, 0x62, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x02, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x12, 0x32, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_rules_v1_rule_subgroup_proto_rawDescOnce sync.Once
	file_com_coralogix_rules_v1_rule_subgroup_proto_rawDescData = file_com_coralogix_rules_v1_rule_subgroup_proto_rawDesc
)

func file_com_coralogix_rules_v1_rule_subgroup_proto_rawDescGZIP() []byte {
	file_com_coralogix_rules_v1_rule_subgroup_proto_rawDescOnce.Do(func() {
		file_com_coralogix_rules_v1_rule_subgroup_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_rules_v1_rule_subgroup_proto_rawDescData)
	})
	return file_com_coralogix_rules_v1_rule_subgroup_proto_rawDescData
}

var file_com_coralogix_rules_v1_rule_subgroup_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogix_rules_v1_rule_subgroup_proto_goTypes = []any{
	(*RuleSubgroup)(nil),           // 0: com.coralogix.rules.v1.RuleSubgroup
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
	(*Rule)(nil),                   // 2: com.coralogix.rules.v1.Rule
	(*wrapperspb.BoolValue)(nil),   // 3: google.protobuf.BoolValue
	(*wrapperspb.UInt32Value)(nil), // 4: google.protobuf.UInt32Value
}
var file_com_coralogix_rules_v1_rule_subgroup_proto_depIdxs = []int32{
	1, // 0: com.coralogix.rules.v1.RuleSubgroup.id:type_name -> google.protobuf.StringValue
	2, // 1: com.coralogix.rules.v1.RuleSubgroup.rules:type_name -> com.coralogix.rules.v1.Rule
	3, // 2: com.coralogix.rules.v1.RuleSubgroup.enabled:type_name -> google.protobuf.BoolValue
	4, // 3: com.coralogix.rules.v1.RuleSubgroup.order:type_name -> google.protobuf.UInt32Value
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_com_coralogix_rules_v1_rule_subgroup_proto_init() }
func file_com_coralogix_rules_v1_rule_subgroup_proto_init() {
	if File_com_coralogix_rules_v1_rule_subgroup_proto != nil {
		return
	}
	file_com_coralogix_rules_v1_rule_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_rules_v1_rule_subgroup_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RuleSubgroup); i {
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
			RawDescriptor: file_com_coralogix_rules_v1_rule_subgroup_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_rules_v1_rule_subgroup_proto_goTypes,
		DependencyIndexes: file_com_coralogix_rules_v1_rule_subgroup_proto_depIdxs,
		MessageInfos:      file_com_coralogix_rules_v1_rule_subgroup_proto_msgTypes,
	}.Build()
	File_com_coralogix_rules_v1_rule_subgroup_proto = out.File
	file_com_coralogix_rules_v1_rule_subgroup_proto_rawDesc = nil
	file_com_coralogix_rules_v1_rule_subgroup_proto_goTypes = nil
	file_com_coralogix_rules_v1_rule_subgroup_proto_depIdxs = nil
}
