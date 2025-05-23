// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogix/quota/v1/policy_order.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

const file_com_coralogix_quota_v1_policy_order_proto_rawDesc = "" +
	"\n" +
	")com/coralogix/quota/v1/policy_order.proto\x12\x16com.coralogix.quota.v1\x1a\x1egoogle/protobuf/wrappers.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\xc4\x02\n" +
	"\vPolicyOrder\x129\n" +
	"\x05order\x18\x01 \x01(\v2\x1b.google.protobuf.Int32ValueB\x06\x92A\x03J\x011R\x05order\x127\n" +
	"\x02id\x18\x02 \x01(\v2\x1c.google.protobuf.StringValueB\t\x92A\x06J\x04\"id\"R\x02id:\xc0\x01\x92A\xbc\x01\n" +
	"/*\fPolicy Order2\x12Order of a policy.\xd2\x01\x05order\xd2\x01\x02id*\x88\x01\n" +
	"%Find out more about quota management.\x12_https://coralogix.com/docs/user-guides/account-management/payment-and-billing/quota-management/b\x06proto3"

var (
	file_com_coralogix_quota_v1_policy_order_proto_rawDescOnce sync.Once
	file_com_coralogix_quota_v1_policy_order_proto_rawDescData []byte
)

func file_com_coralogix_quota_v1_policy_order_proto_rawDescGZIP() []byte {
	file_com_coralogix_quota_v1_policy_order_proto_rawDescOnce.Do(func() {
		file_com_coralogix_quota_v1_policy_order_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogix_quota_v1_policy_order_proto_rawDesc), len(file_com_coralogix_quota_v1_policy_order_proto_rawDesc)))
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogix_quota_v1_policy_order_proto_rawDesc), len(file_com_coralogix_quota_v1_policy_order_proto_rawDesc)),
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
	file_com_coralogix_quota_v1_policy_order_proto_goTypes = nil
	file_com_coralogix_quota_v1_policy_order_proto_depIdxs = nil
}
