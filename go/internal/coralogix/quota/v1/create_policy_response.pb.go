// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.1
// source: com/coralogix/quota/v1/create_policy_response.proto

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

type CreatePolicyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Policy *Policy `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
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

var file_com_coralogix_quota_v1_create_policy_response_proto_rawDesc = []byte{
	0x0a, 0x33, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x23, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x71, 0x75, 0x6f,
	0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_quota_v1_create_policy_response_proto_rawDescOnce sync.Once
	file_com_coralogix_quota_v1_create_policy_response_proto_rawDescData = file_com_coralogix_quota_v1_create_policy_response_proto_rawDesc
)

func file_com_coralogix_quota_v1_create_policy_response_proto_rawDescGZIP() []byte {
	file_com_coralogix_quota_v1_create_policy_response_proto_rawDescOnce.Do(func() {
		file_com_coralogix_quota_v1_create_policy_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_quota_v1_create_policy_response_proto_rawDescData)
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
			RawDescriptor: file_com_coralogix_quota_v1_create_policy_response_proto_rawDesc,
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
	file_com_coralogix_quota_v1_create_policy_response_proto_rawDesc = nil
	file_com_coralogix_quota_v1_create_policy_response_proto_goTypes = nil
	file_com_coralogix_quota_v1_create_policy_response_proto_depIdxs = nil
}
