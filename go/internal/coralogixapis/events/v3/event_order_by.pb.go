// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.26.1
// source: com/coralogixapis/events/v3/event_order_by.proto

package v3

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

type OrderByFields int32

const (
	OrderByFields_ORDER_BY_FIELDS_UNSPECIFIED OrderByFields = 0
	OrderByFields_ORDER_BY_FIELDS_TIMESTAMP   OrderByFields = 1
)

// Enum value maps for OrderByFields.
var (
	OrderByFields_name = map[int32]string{
		0: "ORDER_BY_FIELDS_UNSPECIFIED",
		1: "ORDER_BY_FIELDS_TIMESTAMP",
	}
	OrderByFields_value = map[string]int32{
		"ORDER_BY_FIELDS_UNSPECIFIED": 0,
		"ORDER_BY_FIELDS_TIMESTAMP":   1,
	}
)

func (x OrderByFields) Enum() *OrderByFields {
	p := new(OrderByFields)
	*p = x
	return p
}

func (x OrderByFields) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderByFields) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_events_v3_event_order_by_proto_enumTypes[0].Descriptor()
}

func (OrderByFields) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_events_v3_event_order_by_proto_enumTypes[0]
}

func (x OrderByFields) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderByFields.Descriptor instead.
func (OrderByFields) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_events_v3_event_order_by_proto_rawDescGZIP(), []int{0}
}

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
	return file_com_coralogixapis_events_v3_event_order_by_proto_enumTypes[1].Descriptor()
}

func (OrderByDirection) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_events_v3_event_order_by_proto_enumTypes[1]
}

func (x OrderByDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderByDirection.Descriptor instead.
func (OrderByDirection) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_events_v3_event_order_by_proto_rawDescGZIP(), []int{1}
}

type OrderBy struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FieldName     OrderByFields          `protobuf:"varint,1,opt,name=field_name,json=fieldName,proto3,enum=com.coralogixapis.events.v3.OrderByFields" json:"field_name,omitempty"`
	Direction     OrderByDirection       `protobuf:"varint,2,opt,name=direction,proto3,enum=com.coralogixapis.events.v3.OrderByDirection" json:"direction,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderBy) Reset() {
	*x = OrderBy{}
	mi := &file_com_coralogixapis_events_v3_event_order_by_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderBy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderBy) ProtoMessage() {}

func (x *OrderBy) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_events_v3_event_order_by_proto_msgTypes[0]
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
	return file_com_coralogixapis_events_v3_event_order_by_proto_rawDescGZIP(), []int{0}
}

func (x *OrderBy) GetFieldName() OrderByFields {
	if x != nil {
		return x.FieldName
	}
	return OrderByFields_ORDER_BY_FIELDS_UNSPECIFIED
}

func (x *OrderBy) GetDirection() OrderByDirection {
	if x != nil {
		return x.Direction
	}
	return OrderByDirection_ORDER_BY_DIRECTION_UNSPECIFIED
}

var File_com_coralogixapis_events_v3_event_order_by_proto protoreflect.FileDescriptor

const file_com_coralogixapis_events_v3_event_order_by_proto_rawDesc = "" +
	"\n" +
	"0com/coralogixapis/events/v3/event_order_by.proto\x12\x1bcom.coralogixapis.events.v3\x1a.protoc-gen-openapiv2/options/annotations.proto\"\xe2\x01\n" +
	"\aOrderBy\x12k\n" +
	"\n" +
	"field_name\x18\x01 \x01(\x0e2*.com.coralogixapis.events.v3.OrderByFieldsB \x92A\x1dJ\x1b\"ORDER_BY_FIELDS_TIMESTAMP\"R\tfieldName\x12j\n" +
	"\tdirection\x18\x02 \x01(\x0e2-.com.coralogixapis.events.v3.OrderByDirectionB\x1d\x92A\x1aJ\x18\"ORDER_BY_DIRECTION_ASC\"R\tdirection*O\n" +
	"\rOrderByFields\x12\x1f\n" +
	"\x1bORDER_BY_FIELDS_UNSPECIFIED\x10\x00\x12\x1d\n" +
	"\x19ORDER_BY_FIELDS_TIMESTAMP\x10\x01*o\n" +
	"\x10OrderByDirection\x12\"\n" +
	"\x1eORDER_BY_DIRECTION_UNSPECIFIED\x10\x00\x12\x1a\n" +
	"\x16ORDER_BY_DIRECTION_ASC\x10\x01\x12\x1b\n" +
	"\x17ORDER_BY_DIRECTION_DESC\x10\x02b\x06proto3"

var (
	file_com_coralogixapis_events_v3_event_order_by_proto_rawDescOnce sync.Once
	file_com_coralogixapis_events_v3_event_order_by_proto_rawDescData []byte
)

func file_com_coralogixapis_events_v3_event_order_by_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_events_v3_event_order_by_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_events_v3_event_order_by_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_events_v3_event_order_by_proto_rawDesc), len(file_com_coralogixapis_events_v3_event_order_by_proto_rawDesc)))
	})
	return file_com_coralogixapis_events_v3_event_order_by_proto_rawDescData
}

var file_com_coralogixapis_events_v3_event_order_by_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_com_coralogixapis_events_v3_event_order_by_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_events_v3_event_order_by_proto_goTypes = []any{
	(OrderByFields)(0),    // 0: com.coralogixapis.events.v3.OrderByFields
	(OrderByDirection)(0), // 1: com.coralogixapis.events.v3.OrderByDirection
	(*OrderBy)(nil),       // 2: com.coralogixapis.events.v3.OrderBy
}
var file_com_coralogixapis_events_v3_event_order_by_proto_depIdxs = []int32{
	0, // 0: com.coralogixapis.events.v3.OrderBy.field_name:type_name -> com.coralogixapis.events.v3.OrderByFields
	1, // 1: com.coralogixapis.events.v3.OrderBy.direction:type_name -> com.coralogixapis.events.v3.OrderByDirection
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_events_v3_event_order_by_proto_init() }
func file_com_coralogixapis_events_v3_event_order_by_proto_init() {
	if File_com_coralogixapis_events_v3_event_order_by_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_events_v3_event_order_by_proto_rawDesc), len(file_com_coralogixapis_events_v3_event_order_by_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_events_v3_event_order_by_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_events_v3_event_order_by_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_events_v3_event_order_by_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_events_v3_event_order_by_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_events_v3_event_order_by_proto = out.File
	file_com_coralogixapis_events_v3_event_order_by_proto_goTypes = nil
	file_com_coralogixapis_events_v3_event_order_by_proto_depIdxs = nil
}
