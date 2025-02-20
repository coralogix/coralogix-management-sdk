// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.26.1
// source: com/coralogix/quota/v1/policy_order.proto

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

type PolicyOrder struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Order         *wrapperspb.Int32Value  `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	Id            *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PolicyOrder) Reset() {
	*x = PolicyOrder{}
	mi := &file_com_coralogix_quota_v1_policy_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PolicyOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyOrder) ProtoMessage() {}

func (x *PolicyOrder) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_quota_v1_policy_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyOrder.ProtoReflect.Descriptor instead.
func (*PolicyOrder) Descriptor() ([]byte, []int) {
	return file_com_coralogix_quota_v1_policy_order_proto_rawDescGZIP(), []int{0}
}

func (x *PolicyOrder) GetOrder() *wrapperspb.Int32Value {
	if x != nil {
		return x.Order
	}
	return nil
}

func (x *PolicyOrder) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

var File_com_coralogix_quota_v1_policy_order_proto protoreflect.FileDescriptor

var file_com_coralogix_quota_v1_policy_order_proto_rawDesc = []byte{
	0x0a, 0x29, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61,
	0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x0b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x31, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x02, 0x69, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_quota_v1_policy_order_proto_rawDescOnce sync.Once
	file_com_coralogix_quota_v1_policy_order_proto_rawDescData = file_com_coralogix_quota_v1_policy_order_proto_rawDesc
)

func file_com_coralogix_quota_v1_policy_order_proto_rawDescGZIP() []byte {
	file_com_coralogix_quota_v1_policy_order_proto_rawDescOnce.Do(func() {
		file_com_coralogix_quota_v1_policy_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_quota_v1_policy_order_proto_rawDescData)
	})
	return file_com_coralogix_quota_v1_policy_order_proto_rawDescData
}

var file_com_coralogix_quota_v1_policy_order_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogix_quota_v1_policy_order_proto_goTypes = []any{
	(*PolicyOrder)(nil),            // 0: com.coralogix.quota.v1.PolicyOrder
	(*wrapperspb.Int32Value)(nil),  // 1: google.protobuf.Int32Value
	(*wrapperspb.StringValue)(nil), // 2: google.protobuf.StringValue
}
var file_com_coralogix_quota_v1_policy_order_proto_depIdxs = []int32{
	1, // 0: com.coralogix.quota.v1.PolicyOrder.order:type_name -> google.protobuf.Int32Value
	2, // 1: com.coralogix.quota.v1.PolicyOrder.id:type_name -> google.protobuf.StringValue
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_com_coralogix_quota_v1_policy_order_proto_init() }
func file_com_coralogix_quota_v1_policy_order_proto_init() {
	if File_com_coralogix_quota_v1_policy_order_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_quota_v1_policy_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_quota_v1_policy_order_proto_goTypes,
		DependencyIndexes: file_com_coralogix_quota_v1_policy_order_proto_depIdxs,
		MessageInfos:      file_com_coralogix_quota_v1_policy_order_proto_msgTypes,
	}.Build()
	File_com_coralogix_quota_v1_policy_order_proto = out.File
	file_com_coralogix_quota_v1_policy_order_proto_rawDesc = nil
	file_com_coralogix_quota_v1_policy_order_proto_goTypes = nil
	file_com_coralogix_quota_v1_policy_order_proto_depIdxs = nil
}
