// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: com/coralogix/quota/v1/test_policies_result.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type TestPoliciesResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetaFieldsValues *LogMetaFieldsValues  `protobuf:"bytes,1,opt,name=meta_fields_values,json=metaFieldsValues,proto3" json:"meta_fields_values,omitempty"`
	Matched          *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=matched,proto3" json:"matched,omitempty"`
	Policy           *Policy               `protobuf:"bytes,3,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (x *TestPoliciesResult) Reset() {
	*x = TestPoliciesResult{}
	mi := &file_com_coralogix_quota_v1_test_policies_result_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TestPoliciesResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestPoliciesResult) ProtoMessage() {}

func (x *TestPoliciesResult) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_quota_v1_test_policies_result_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestPoliciesResult.ProtoReflect.Descriptor instead.
func (*TestPoliciesResult) Descriptor() ([]byte, []int) {
	return file_com_coralogix_quota_v1_test_policies_result_proto_rawDescGZIP(), []int{0}
}

func (x *TestPoliciesResult) GetMetaFieldsValues() *LogMetaFieldsValues {
	if x != nil {
		return x.MetaFieldsValues
	}
	return nil
}

func (x *TestPoliciesResult) GetMatched() *wrapperspb.BoolValue {
	if x != nil {
		return x.Matched
	}
	return nil
}

func (x *TestPoliciesResult) GetPolicy() *Policy {
	if x != nil {
		return x.Policy
	}
	return nil
}

var File_com_coralogix_quota_v1_test_policies_result_proto protoreflect.FileDescriptor

var file_com_coralogix_quota_v1_test_policies_result_proto_rawDesc = []byte{
	0x0a, 0x31, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x32, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x6d, 0x65, 0x74,
	0x61, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x03, 0x0a, 0x12, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x59, 0x0a, 0x12, 0x6d,
	0x65, 0x74, 0x61, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x6f, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x52, 0x10, 0x6d, 0x65, 0x74, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x09, 0x92, 0x41, 0x06, 0x4a, 0x04, 0x74, 0x72, 0x75, 0x65, 0x52, 0x07,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x12, 0x36, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x3a,
	0xec, 0x01, 0x92, 0x41, 0xe8, 0x01, 0x0a, 0x5b, 0x2a, 0x14, 0x54, 0x65, 0x73, 0x74, 0x20, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x20, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x1b,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0xd2, 0x01, 0x12, 0x6d, 0x65,
	0x74, 0x61, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0xd2, 0x01, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0xd2, 0x01, 0x06, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2a, 0x88, 0x01, 0x0a, 0x25, 0x46, 0x69, 0x6e, 0x64, 0x20, 0x6f, 0x75, 0x74,
	0x20, 0x6d, 0x6f, 0x72, 0x65, 0x20, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x20, 0x71, 0x75, 0x6f, 0x74,
	0x61, 0x20, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x12, 0x5f, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x63, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x67,
	0x75, 0x69, 0x64, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2d, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2d, 0x61, 0x6e, 0x64, 0x2d, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x71, 0x75, 0x6f,
	0x74, 0x61, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_quota_v1_test_policies_result_proto_rawDescOnce sync.Once
	file_com_coralogix_quota_v1_test_policies_result_proto_rawDescData = file_com_coralogix_quota_v1_test_policies_result_proto_rawDesc
)

func file_com_coralogix_quota_v1_test_policies_result_proto_rawDescGZIP() []byte {
	file_com_coralogix_quota_v1_test_policies_result_proto_rawDescOnce.Do(func() {
		file_com_coralogix_quota_v1_test_policies_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_quota_v1_test_policies_result_proto_rawDescData)
	})
	return file_com_coralogix_quota_v1_test_policies_result_proto_rawDescData
}

var file_com_coralogix_quota_v1_test_policies_result_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogix_quota_v1_test_policies_result_proto_goTypes = []any{
	(*TestPoliciesResult)(nil),   // 0: com.coralogix.quota.v1.TestPoliciesResult
	(*LogMetaFieldsValues)(nil),  // 1: com.coralogix.quota.v1.LogMetaFieldsValues
	(*wrapperspb.BoolValue)(nil), // 2: google.protobuf.BoolValue
	(*Policy)(nil),               // 3: com.coralogix.quota.v1.Policy
}
var file_com_coralogix_quota_v1_test_policies_result_proto_depIdxs = []int32{
	1, // 0: com.coralogix.quota.v1.TestPoliciesResult.meta_fields_values:type_name -> com.coralogix.quota.v1.LogMetaFieldsValues
	2, // 1: com.coralogix.quota.v1.TestPoliciesResult.matched:type_name -> google.protobuf.BoolValue
	3, // 2: com.coralogix.quota.v1.TestPoliciesResult.policy:type_name -> com.coralogix.quota.v1.Policy
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_com_coralogix_quota_v1_test_policies_result_proto_init() }
func file_com_coralogix_quota_v1_test_policies_result_proto_init() {
	if File_com_coralogix_quota_v1_test_policies_result_proto != nil {
		return
	}
	file_com_coralogix_quota_v1_policy_proto_init()
	file_com_coralogix_quota_v1_log_meta_field_values_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_quota_v1_test_policies_result_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_quota_v1_test_policies_result_proto_goTypes,
		DependencyIndexes: file_com_coralogix_quota_v1_test_policies_result_proto_depIdxs,
		MessageInfos:      file_com_coralogix_quota_v1_test_policies_result_proto_msgTypes,
	}.Build()
	File_com_coralogix_quota_v1_test_policies_result_proto = out.File
	file_com_coralogix_quota_v1_test_policies_result_proto_rawDesc = nil
	file_com_coralogix_quota_v1_test_policies_result_proto_goTypes = nil
	file_com_coralogix_quota_v1_test_policies_result_proto_depIdxs = nil
}
