// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: com/coralogixapis/dashboards/v1/common/order_direction.proto

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

type OrderDirection int32

const (
	OrderDirection_ORDER_DIRECTION_UNSPECIFIED OrderDirection = 0
	OrderDirection_ORDER_DIRECTION_ASC         OrderDirection = 1
	OrderDirection_ORDER_DIRECTION_DESC        OrderDirection = 2
)

// Enum value maps for OrderDirection.
var (
	OrderDirection_name = map[int32]string{
		0: "ORDER_DIRECTION_UNSPECIFIED",
		1: "ORDER_DIRECTION_ASC",
		2: "ORDER_DIRECTION_DESC",
	}
	OrderDirection_value = map[string]int32{
		"ORDER_DIRECTION_UNSPECIFIED": 0,
		"ORDER_DIRECTION_ASC":         1,
		"ORDER_DIRECTION_DESC":        2,
	}
)

func (x OrderDirection) Enum() *OrderDirection {
	p := new(OrderDirection)
	*p = x
	return p
}

func (x OrderDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_dashboards_v1_common_order_direction_proto_enumTypes[0].Descriptor()
}

func (OrderDirection) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_dashboards_v1_common_order_direction_proto_enumTypes[0]
}

func (x OrderDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderDirection.Descriptor instead.
func (OrderDirection) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_order_direction_proto_rawDescGZIP(), []int{0}
}

var File_com_coralogixapis_dashboards_v1_common_order_direction_proto protoreflect.FileDescriptor

var file_com_coralogixapis_dashboards_v1_common_order_direction_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2a, 0x64, 0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x1b, 0x4f, 0x52, 0x44, 0x45,
	0x52, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x52, 0x44,
	0x45, 0x52, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x53, 0x43,
	0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x44, 0x49, 0x52, 0x45,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_dashboards_v1_common_order_direction_proto_rawDescOnce sync.Once
	file_com_coralogixapis_dashboards_v1_common_order_direction_proto_rawDescData = file_com_coralogixapis_dashboards_v1_common_order_direction_proto_rawDesc
)

func file_com_coralogixapis_dashboards_v1_common_order_direction_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_dashboards_v1_common_order_direction_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_dashboards_v1_common_order_direction_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_dashboards_v1_common_order_direction_proto_rawDescData)
	})
	return file_com_coralogixapis_dashboards_v1_common_order_direction_proto_rawDescData
}

var file_com_coralogixapis_dashboards_v1_common_order_direction_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_dashboards_v1_common_order_direction_proto_goTypes = []any{
	(OrderDirection)(0), // 0: com.coralogixapis.dashboards.v1.common.OrderDirection
}
var file_com_coralogixapis_dashboards_v1_common_order_direction_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_dashboards_v1_common_order_direction_proto_init() }
func file_com_coralogixapis_dashboards_v1_common_order_direction_proto_init() {
	if File_com_coralogixapis_dashboards_v1_common_order_direction_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_dashboards_v1_common_order_direction_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_dashboards_v1_common_order_direction_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_dashboards_v1_common_order_direction_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_dashboards_v1_common_order_direction_proto_enumTypes,
	}.Build()
	File_com_coralogixapis_dashboards_v1_common_order_direction_proto = out.File
	file_com_coralogixapis_dashboards_v1_common_order_direction_proto_rawDesc = nil
	file_com_coralogixapis_dashboards_v1_common_order_direction_proto_goTypes = nil
	file_com_coralogixapis_dashboards_v1_common_order_direction_proto_depIdxs = nil
}
