// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.26.1
// source: com/coralogixapis/apm/common/v2/order_by.proto

package v2

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

type OrderByDirection int32

const (
	OrderByDirection_ORDER_BY_DIRECTION_UNSPECIFIED OrderByDirection = 0
	OrderByDirection_ORDER_BY_DIRECTION_ASC         OrderByDirection = 1
	OrderByDirection_ORDER_BY_DIRECTION_DESC        OrderByDirection = 2
)

// Enum value maps for OrderByDirection.
var (
	OrderByDirection_name = map[int32]string{
		0: "ORDER_BY_DIRECTION_UNSPECIFIED",
		1: "ORDER_BY_DIRECTION_ASC",
		2: "ORDER_BY_DIRECTION_DESC",
	}
	OrderByDirection_value = map[string]int32{
		"ORDER_BY_DIRECTION_UNSPECIFIED": 0,
		"ORDER_BY_DIRECTION_ASC":         1,
		"ORDER_BY_DIRECTION_DESC":        2,
	}
)

func (x OrderByDirection) Enum() *OrderByDirection {
	p := new(OrderByDirection)
	*p = x
	return p
}

func (x OrderByDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderByDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_apm_common_v2_order_by_proto_enumTypes[0].Descriptor()
}

func (OrderByDirection) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_apm_common_v2_order_by_proto_enumTypes[0]
}

func (x OrderByDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderByDirection.Descriptor instead.
func (OrderByDirection) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_common_v2_order_by_proto_rawDescGZIP(), []int{0}
}

type OrderBy struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	FieldName     *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	Direction     OrderByDirection        `protobuf:"varint,2,opt,name=direction,proto3,enum=com.coralogixapis.apm.common.v2.OrderByDirection" json:"direction,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderBy) Reset() {
	*x = OrderBy{}
	mi := &file_com_coralogixapis_apm_common_v2_order_by_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderBy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderBy) ProtoMessage() {}

func (x *OrderBy) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_common_v2_order_by_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderBy.ProtoReflect.Descriptor instead.
func (*OrderBy) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_common_v2_order_by_proto_rawDescGZIP(), []int{0}
}

func (x *OrderBy) GetFieldName() *wrapperspb.StringValue {
	if x != nil {
		return x.FieldName
	}
	return nil
}

func (x *OrderBy) GetDirection() OrderByDirection {
	if x != nil {
		return x.Direction
	}
	return OrderByDirection_ORDER_BY_DIRECTION_UNSPECIFIED
}

var File_com_coralogixapis_apm_common_v2_order_by_proto protoreflect.FileDescriptor

var file_com_coralogixapis_apm_common_v2_order_by_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x32, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x32, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa3, 0x03, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x4e, 0x0a,
	0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x11, 0x92, 0x41, 0x0e, 0x4a, 0x0c, 0x22, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4f, 0x0a,
	0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x32, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0xf6,
	0x01, 0x92, 0x41, 0xf2, 0x01, 0x0a, 0x68, 0x2a, 0x08, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x20, 0x42,
	0x79, 0x32, 0x43, 0x54, 0x68, 0x69, 0x73, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x20, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74,
	0x73, 0x20, 0x61, 0x6e, 0x20, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x20, 0x62, 0x79, 0x20, 0x63, 0x6c,
	0x61, 0x75, 0x73, 0x65, 0x20, 0x69, 0x6e, 0x20, 0x43, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x20, 0x41, 0x50, 0x4d, 0x2e, 0xd2, 0x01, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0xd2, 0x01, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2a,
	0x85, 0x01, 0x0a, 0x29, 0x46, 0x69, 0x6e, 0x64, 0x20, 0x6f, 0x75, 0x74, 0x20, 0x6d, 0x6f, 0x72,
	0x65, 0x20, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x20, 0x53, 0x4c, 0x4f, 0x73, 0x20, 0x69, 0x6e, 0x20,
	0x43, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x20, 0x41, 0x50, 0x4d, 0x12, 0x58, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x63, 0x61, 0x64, 0x65, 0x6d, 0x79, 0x2f, 0x67, 0x65, 0x74,
	0x2d, 0x74, 0x6f, 0x2d, 0x6b, 0x6e, 0x6f, 0x77, 0x2d, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2f, 0x73, 0x6c, 0x6f, 0x2d, 0x73, 0x6c, 0x69, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x69, 0x6e, 0x2d, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2d, 0x61, 0x70, 0x6d, 0x2f, 0x2a, 0x6f, 0x0a, 0x10, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x42, 0x79, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x1e, 0x4f,
	0x52, 0x44, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x1a, 0x0a, 0x16, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x44, 0x49, 0x52, 0x45,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x53, 0x43, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x4f,
	0x52, 0x44, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_apm_common_v2_order_by_proto_rawDescOnce sync.Once
	file_com_coralogixapis_apm_common_v2_order_by_proto_rawDescData = file_com_coralogixapis_apm_common_v2_order_by_proto_rawDesc
)

func file_com_coralogixapis_apm_common_v2_order_by_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_apm_common_v2_order_by_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_apm_common_v2_order_by_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_apm_common_v2_order_by_proto_rawDescData)
	})
	return file_com_coralogixapis_apm_common_v2_order_by_proto_rawDescData
}

var file_com_coralogixapis_apm_common_v2_order_by_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_apm_common_v2_order_by_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_apm_common_v2_order_by_proto_goTypes = []any{
	(OrderByDirection)(0),          // 0: com.coralogixapis.apm.common.v2.OrderByDirection
	(*OrderBy)(nil),                // 1: com.coralogixapis.apm.common.v2.OrderBy
	(*wrapperspb.StringValue)(nil), // 2: google.protobuf.StringValue
}
var file_com_coralogixapis_apm_common_v2_order_by_proto_depIdxs = []int32{
	2, // 0: com.coralogixapis.apm.common.v2.OrderBy.field_name:type_name -> google.protobuf.StringValue
	0, // 1: com.coralogixapis.apm.common.v2.OrderBy.direction:type_name -> com.coralogixapis.apm.common.v2.OrderByDirection
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_apm_common_v2_order_by_proto_init() }
func file_com_coralogixapis_apm_common_v2_order_by_proto_init() {
	if File_com_coralogixapis_apm_common_v2_order_by_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_apm_common_v2_order_by_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_apm_common_v2_order_by_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_apm_common_v2_order_by_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_apm_common_v2_order_by_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_apm_common_v2_order_by_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_apm_common_v2_order_by_proto = out.File
	file_com_coralogixapis_apm_common_v2_order_by_proto_rawDesc = nil
	file_com_coralogixapis_apm_common_v2_order_by_proto_goTypes = nil
	file_com_coralogixapis_apm_common_v2_order_by_proto_depIdxs = nil
}
