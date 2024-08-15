// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: com/coralogix/quota/v1/bulk_create_policy_request.proto

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

type AtomicBatchCreatePolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolicyRequests []*CreatePolicyRequest `protobuf:"bytes,1,rep,name=policy_requests,json=policyRequests,proto3" json:"policy_requests,omitempty"`
}

func (x *AtomicBatchCreatePolicyRequest) Reset() {
	*x = AtomicBatchCreatePolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_quota_v1_bulk_create_policy_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AtomicBatchCreatePolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AtomicBatchCreatePolicyRequest) ProtoMessage() {}

func (x *AtomicBatchCreatePolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_quota_v1_bulk_create_policy_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AtomicBatchCreatePolicyRequest.ProtoReflect.Descriptor instead.
func (*AtomicBatchCreatePolicyRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_quota_v1_bulk_create_policy_request_proto_rawDescGZIP(), []int{0}
}

func (x *AtomicBatchCreatePolicyRequest) GetPolicyRequests() []*CreatePolicyRequest {
	if x != nil {
		return x.PolicyRequests
	}
	return nil
}

var File_com_coralogix_quota_v1_bulk_create_policy_request_proto protoreflect.FileDescriptor

var file_com_coralogix_quota_v1_bulk_create_policy_request_proto_rawDesc = []byte{
	0x0a, 0x37, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x6c, 0x6b, 0x5f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76,
	0x31, 0x1a, 0x32, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x76, 0x0a, 0x1e, 0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x54, 0x0a, 0x0f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0e, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_quota_v1_bulk_create_policy_request_proto_rawDescOnce sync.Once
	file_com_coralogix_quota_v1_bulk_create_policy_request_proto_rawDescData = file_com_coralogix_quota_v1_bulk_create_policy_request_proto_rawDesc
)

func file_com_coralogix_quota_v1_bulk_create_policy_request_proto_rawDescGZIP() []byte {
	file_com_coralogix_quota_v1_bulk_create_policy_request_proto_rawDescOnce.Do(func() {
		file_com_coralogix_quota_v1_bulk_create_policy_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_quota_v1_bulk_create_policy_request_proto_rawDescData)
	})
	return file_com_coralogix_quota_v1_bulk_create_policy_request_proto_rawDescData
}

var file_com_coralogix_quota_v1_bulk_create_policy_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogix_quota_v1_bulk_create_policy_request_proto_goTypes = []any{
	(*AtomicBatchCreatePolicyRequest)(nil), // 0: com.coralogix.quota.v1.AtomicBatchCreatePolicyRequest
	(*CreatePolicyRequest)(nil),            // 1: com.coralogix.quota.v1.CreatePolicyRequest
}
var file_com_coralogix_quota_v1_bulk_create_policy_request_proto_depIdxs = []int32{
	1, // 0: com.coralogix.quota.v1.AtomicBatchCreatePolicyRequest.policy_requests:type_name -> com.coralogix.quota.v1.CreatePolicyRequest
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_com_coralogix_quota_v1_bulk_create_policy_request_proto_init() }
func file_com_coralogix_quota_v1_bulk_create_policy_request_proto_init() {
	if File_com_coralogix_quota_v1_bulk_create_policy_request_proto != nil {
		return
	}
	file_com_coralogix_quota_v1_create_policy_request_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_quota_v1_bulk_create_policy_request_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AtomicBatchCreatePolicyRequest); i {
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
			RawDescriptor: file_com_coralogix_quota_v1_bulk_create_policy_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_quota_v1_bulk_create_policy_request_proto_goTypes,
		DependencyIndexes: file_com_coralogix_quota_v1_bulk_create_policy_request_proto_depIdxs,
		MessageInfos:      file_com_coralogix_quota_v1_bulk_create_policy_request_proto_msgTypes,
	}.Build()
	File_com_coralogix_quota_v1_bulk_create_policy_request_proto = out.File
	file_com_coralogix_quota_v1_bulk_create_policy_request_proto_rawDesc = nil
	file_com_coralogix_quota_v1_bulk_create_policy_request_proto_goTypes = nil
	file_com_coralogix_quota_v1_bulk_create_policy_request_proto_depIdxs = nil
}
