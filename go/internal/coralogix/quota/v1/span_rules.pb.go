// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.3
// source: com/coralogix/quota/v1/span_rules.proto

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

type SpanRules struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceRule *Rule      `protobuf:"bytes,1,opt,name=service_rule,json=serviceRule,proto3,oneof" json:"service_rule,omitempty"`
	ActionRule  *Rule      `protobuf:"bytes,2,opt,name=action_rule,json=actionRule,proto3,oneof" json:"action_rule,omitempty"`
	TagRules    []*TagRule `protobuf:"bytes,3,rep,name=tag_rules,json=tagRules,proto3" json:"tag_rules,omitempty"`
}

func (x *SpanRules) Reset() {
	*x = SpanRules{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_quota_v1_span_rules_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpanRules) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpanRules) ProtoMessage() {}

func (x *SpanRules) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_quota_v1_span_rules_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpanRules.ProtoReflect.Descriptor instead.
func (*SpanRules) Descriptor() ([]byte, []int) {
	return file_com_coralogix_quota_v1_span_rules_proto_rawDescGZIP(), []int{0}
}

func (x *SpanRules) GetServiceRule() *Rule {
	if x != nil {
		return x.ServiceRule
	}
	return nil
}

func (x *SpanRules) GetActionRule() *Rule {
	if x != nil {
		return x.ActionRule
	}
	return nil
}

func (x *SpanRules) GetTagRules() []*TagRule {
	if x != nil {
		return x.TagRules
	}
	return nil
}

var File_com_coralogix_quota_v1_span_rules_proto protoreflect.FileDescriptor

var file_com_coralogix_quota_v1_span_rules_proto_rawDesc = []byte{
	0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x5f, 0x72, 0x75,
	0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76,
	0x31, 0x1a, 0x21, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x67,
	0x5f, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4, 0x01, 0x0a, 0x09,
	0x53, 0x70, 0x61, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x0c, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x48, 0x00, 0x52,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x42, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75,
	0x6c, 0x65, 0x48, 0x01, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6c, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x3c, 0x0a, 0x09, 0x74, 0x61, 0x67, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x61, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x08, 0x74, 0x61, 0x67, 0x52, 0x75, 0x6c, 0x65,
	0x73, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x72, 0x75,
	0x6c, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x75,
	0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_quota_v1_span_rules_proto_rawDescOnce sync.Once
	file_com_coralogix_quota_v1_span_rules_proto_rawDescData = file_com_coralogix_quota_v1_span_rules_proto_rawDesc
)

func file_com_coralogix_quota_v1_span_rules_proto_rawDescGZIP() []byte {
	file_com_coralogix_quota_v1_span_rules_proto_rawDescOnce.Do(func() {
		file_com_coralogix_quota_v1_span_rules_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_quota_v1_span_rules_proto_rawDescData)
	})
	return file_com_coralogix_quota_v1_span_rules_proto_rawDescData
}

var file_com_coralogix_quota_v1_span_rules_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogix_quota_v1_span_rules_proto_goTypes = []interface{}{
	(*SpanRules)(nil), // 0: com.coralogix.quota.v1.SpanRules
	(*Rule)(nil),      // 1: com.coralogix.quota.v1.Rule
	(*TagRule)(nil),   // 2: com.coralogix.quota.v1.TagRule
}
var file_com_coralogix_quota_v1_span_rules_proto_depIdxs = []int32{
	1, // 0: com.coralogix.quota.v1.SpanRules.service_rule:type_name -> com.coralogix.quota.v1.Rule
	1, // 1: com.coralogix.quota.v1.SpanRules.action_rule:type_name -> com.coralogix.quota.v1.Rule
	2, // 2: com.coralogix.quota.v1.SpanRules.tag_rules:type_name -> com.coralogix.quota.v1.TagRule
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_com_coralogix_quota_v1_span_rules_proto_init() }
func file_com_coralogix_quota_v1_span_rules_proto_init() {
	if File_com_coralogix_quota_v1_span_rules_proto != nil {
		return
	}
	file_com_coralogix_quota_v1_rule_proto_init()
	file_com_coralogix_quota_v1_tag_rule_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_quota_v1_span_rules_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpanRules); i {
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
	file_com_coralogix_quota_v1_span_rules_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_quota_v1_span_rules_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_quota_v1_span_rules_proto_goTypes,
		DependencyIndexes: file_com_coralogix_quota_v1_span_rules_proto_depIdxs,
		MessageInfos:      file_com_coralogix_quota_v1_span_rules_proto_msgTypes,
	}.Build()
	File_com_coralogix_quota_v1_span_rules_proto = out.File
	file_com_coralogix_quota_v1_span_rules_proto_rawDesc = nil
	file_com_coralogix_quota_v1_span_rules_proto_goTypes = nil
	file_com_coralogix_quota_v1_span_rules_proto_depIdxs = nil
}
