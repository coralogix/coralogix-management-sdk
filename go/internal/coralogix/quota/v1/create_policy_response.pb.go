// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogix/quota/v1/create_policy_response.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type CreatePolicyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Policy        *Policy                `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePolicyResponse) Reset() {
	*x = CreatePolicyResponse{}
	mi := &file_com_coralogix_quota_v1_create_policy_response_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePolicyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePolicyResponse) ProtoMessage() {}

func (x *CreatePolicyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_quota_v1_create_policy_response_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePolicyResponse.ProtoReflect.Descriptor instead.
func (*CreatePolicyResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_quota_v1_create_policy_response_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePolicyResponse) GetPolicy() *Policy {
	if x != nil {
		return x.Policy
	}
	return nil
}

var File_com_coralogix_quota_v1_create_policy_response_proto protoreflect.FileDescriptor

const file_com_coralogix_quota_v1_create_policy_response_proto_rawDesc = "" +
	"\n" +
	"3com/coralogix/quota/v1/create_policy_response.proto\x12\x16com.coralogix.quota.v1\x1a#com/coralogix/quota/v1/policy.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\xbb\x02\n" +
	"\x14CreatePolicyResponse\x126\n" +
	"\x06policy\x18\x01 \x01(\v2\x1e.com.coralogix.quota.v1.PolicyR\x06policy:\xea\x01\x92A\xe6\x01\n" +
	"Y*\x16Create Policy Response26This data structue is obtained when creating a policy.\xd2\x01\x06policy*\x88\x01\n" +
	"%Find out more about quota management.\x12_https://coralogix.com/docs/user-guides/account-management/payment-and-billing/quota-management/b\x06proto3"

var (
	file_com_coralogix_quota_v1_create_policy_response_proto_rawDescOnce sync.Once
	file_com_coralogix_quota_v1_create_policy_response_proto_rawDescData []byte
)

func file_com_coralogix_quota_v1_create_policy_response_proto_rawDescGZIP() []byte {
	file_com_coralogix_quota_v1_create_policy_response_proto_rawDescOnce.Do(func() {
		file_com_coralogix_quota_v1_create_policy_response_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogix_quota_v1_create_policy_response_proto_rawDesc), len(file_com_coralogix_quota_v1_create_policy_response_proto_rawDesc)))
	})
	return file_com_coralogix_quota_v1_create_policy_response_proto_rawDescData
}

var file_com_coralogix_quota_v1_create_policy_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogix_quota_v1_create_policy_response_proto_goTypes = []any{
	(*CreatePolicyResponse)(nil), // 0: com.coralogix.quota.v1.CreatePolicyResponse
	(*Policy)(nil),               // 1: com.coralogix.quota.v1.Policy
}
var file_com_coralogix_quota_v1_create_policy_response_proto_depIdxs = []int32{
	1, // 0: com.coralogix.quota.v1.CreatePolicyResponse.policy:type_name -> com.coralogix.quota.v1.Policy
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_com_coralogix_quota_v1_create_policy_response_proto_init() }
func file_com_coralogix_quota_v1_create_policy_response_proto_init() {
	if File_com_coralogix_quota_v1_create_policy_response_proto != nil {
		return
	}
	file_com_coralogix_quota_v1_policy_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogix_quota_v1_create_policy_response_proto_rawDesc), len(file_com_coralogix_quota_v1_create_policy_response_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_quota_v1_create_policy_response_proto_goTypes,
		DependencyIndexes: file_com_coralogix_quota_v1_create_policy_response_proto_depIdxs,
		MessageInfos:      file_com_coralogix_quota_v1_create_policy_response_proto_msgTypes,
	}.Build()
	File_com_coralogix_quota_v1_create_policy_response_proto = out.File
	file_com_coralogix_quota_v1_create_policy_response_proto_goTypes = nil
	file_com_coralogix_quota_v1_create_policy_response_proto_depIdxs = nil
}
