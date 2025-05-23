// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/incidents/v1/incident_query.proto

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

type OrderByFields int32

const (
	OrderByFields_ORDER_BY_FIELDS_UNSPECIFIED  OrderByFields = 0
	OrderByFields_ORDER_BY_FIELDS_ID           OrderByFields = 1
	OrderByFields_ORDER_BY_FIELDS_SEVERITY     OrderByFields = 2
	OrderByFields_ORDER_BY_FIELDS_NAME         OrderByFields = 3
	OrderByFields_ORDER_BY_FIELDS_CREATED_TIME OrderByFields = 4
	OrderByFields_ORDER_BY_FIELDS_CLOSED_TIME  OrderByFields = 5
)

// Enum value maps for OrderByFields.
var (
	OrderByFields_name = map[int32]string{
		0: "ORDER_BY_FIELDS_UNSPECIFIED",
		1: "ORDER_BY_FIELDS_ID",
		2: "ORDER_BY_FIELDS_SEVERITY",
		3: "ORDER_BY_FIELDS_NAME",
		4: "ORDER_BY_FIELDS_CREATED_TIME",
		5: "ORDER_BY_FIELDS_CLOSED_TIME",
	}
	OrderByFields_value = map[string]int32{
		"ORDER_BY_FIELDS_UNSPECIFIED":  0,
		"ORDER_BY_FIELDS_ID":           1,
		"ORDER_BY_FIELDS_SEVERITY":     2,
		"ORDER_BY_FIELDS_NAME":         3,
		"ORDER_BY_FIELDS_CREATED_TIME": 4,
		"ORDER_BY_FIELDS_CLOSED_TIME":  5,
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
	return file_com_coralogixapis_incidents_v1_incident_query_proto_enumTypes[0].Descriptor()
}

func (OrderByFields) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_incidents_v1_incident_query_proto_enumTypes[0]
}

func (x OrderByFields) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderByFields.Descriptor instead.
func (OrderByFields) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_incidents_v1_incident_query_proto_rawDescGZIP(), []int{0}
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
	return file_com_coralogixapis_incidents_v1_incident_query_proto_enumTypes[1].Descriptor()
}

func (OrderByDirection) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_incidents_v1_incident_query_proto_enumTypes[1]
}

func (x OrderByDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderByDirection.Descriptor instead.
func (OrderByDirection) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_incidents_v1_incident_query_proto_rawDescGZIP(), []int{1}
}

type OrderBy struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Field:
	//
	//	*OrderBy_IncidentField
	//	*OrderBy_ContextualLabel
	Field         isOrderBy_Field  `protobuf_oneof:"field"`
	Direction     OrderByDirection `protobuf:"varint,2,opt,name=direction,proto3,enum=com.coralogixapis.incidents.v1.OrderByDirection" json:"direction,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderBy) Reset() {
	*x = OrderBy{}
	mi := &file_com_coralogixapis_incidents_v1_incident_query_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderBy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderBy) ProtoMessage() {}

func (x *OrderBy) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_incidents_v1_incident_query_proto_msgTypes[0]
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
	return file_com_coralogixapis_incidents_v1_incident_query_proto_rawDescGZIP(), []int{0}
}

func (x *OrderBy) GetField() isOrderBy_Field {
	if x != nil {
		return x.Field
	}
	return nil
}

func (x *OrderBy) GetIncidentField() IncidentFields {
	if x != nil {
		if x, ok := x.Field.(*OrderBy_IncidentField); ok {
			return x.IncidentField
		}
	}
	return IncidentFields_INCIDENTS_FIELDS_UNSPECIFIED
}

func (x *OrderBy) GetContextualLabel() *wrapperspb.StringValue {
	if x != nil {
		if x, ok := x.Field.(*OrderBy_ContextualLabel); ok {
			return x.ContextualLabel
		}
	}
	return nil
}

func (x *OrderBy) GetDirection() OrderByDirection {
	if x != nil {
		return x.Direction
	}
	return OrderByDirection_ORDER_BY_DIRECTION_UNSPECIFIED
}

type isOrderBy_Field interface {
	isOrderBy_Field()
}

type OrderBy_IncidentField struct {
	IncidentField IncidentFields `protobuf:"varint,10,opt,name=incident_field,json=incidentField,proto3,enum=com.coralogixapis.incidents.v1.IncidentFields,oneof"`
}

type OrderBy_ContextualLabel struct {
	ContextualLabel *wrapperspb.StringValue `protobuf:"bytes,11,opt,name=contextual_label,json=contextualLabel,proto3,oneof"`
}

func (*OrderBy_IncidentField) isOrderBy_Field() {}

func (*OrderBy_ContextualLabel) isOrderBy_Field() {}

type GroupBy struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Field:
	//
	//	*GroupBy_IncidentField
	//	*GroupBy_ContextualLabel
	Field            isGroupBy_Field  `protobuf_oneof:"field"`
	OrderByDirection OrderByDirection `protobuf:"varint,1,opt,name=order_by_direction,json=orderByDirection,proto3,enum=com.coralogixapis.incidents.v1.OrderByDirection" json:"order_by_direction,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *GroupBy) Reset() {
	*x = GroupBy{}
	mi := &file_com_coralogixapis_incidents_v1_incident_query_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GroupBy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupBy) ProtoMessage() {}

func (x *GroupBy) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_incidents_v1_incident_query_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupBy.ProtoReflect.Descriptor instead.
func (*GroupBy) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_incidents_v1_incident_query_proto_rawDescGZIP(), []int{1}
}

func (x *GroupBy) GetField() isGroupBy_Field {
	if x != nil {
		return x.Field
	}
	return nil
}

func (x *GroupBy) GetIncidentField() IncidentFields {
	if x != nil {
		if x, ok := x.Field.(*GroupBy_IncidentField); ok {
			return x.IncidentField
		}
	}
	return IncidentFields_INCIDENTS_FIELDS_UNSPECIFIED
}

func (x *GroupBy) GetContextualLabel() *wrapperspb.StringValue {
	if x != nil {
		if x, ok := x.Field.(*GroupBy_ContextualLabel); ok {
			return x.ContextualLabel
		}
	}
	return nil
}

func (x *GroupBy) GetOrderByDirection() OrderByDirection {
	if x != nil {
		return x.OrderByDirection
	}
	return OrderByDirection_ORDER_BY_DIRECTION_UNSPECIFIED
}

type isGroupBy_Field interface {
	isGroupBy_Field()
}

type GroupBy_IncidentField struct {
	IncidentField IncidentFields `protobuf:"varint,10,opt,name=incident_field,json=incidentField,proto3,enum=com.coralogixapis.incidents.v1.IncidentFields,oneof"`
}

type GroupBy_ContextualLabel struct {
	ContextualLabel *wrapperspb.StringValue `protobuf:"bytes,11,opt,name=contextual_label,json=contextualLabel,proto3,oneof"`
}

func (*GroupBy_IncidentField) isGroupBy_Field() {}

func (*GroupBy_ContextualLabel) isGroupBy_Field() {}

type IncidentSearchQuery struct {
	state protoimpl.MessageState  `protogen:"open.v1"`
	Query *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// Types that are valid to be assigned to Field:
	//
	//	*IncidentSearchQuery_IncidentField
	//	*IncidentSearchQuery_ContextualLabel
	Field         isIncidentSearchQuery_Field `protobuf_oneof:"field"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IncidentSearchQuery) Reset() {
	*x = IncidentSearchQuery{}
	mi := &file_com_coralogixapis_incidents_v1_incident_query_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IncidentSearchQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncidentSearchQuery) ProtoMessage() {}

func (x *IncidentSearchQuery) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_incidents_v1_incident_query_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncidentSearchQuery.ProtoReflect.Descriptor instead.
func (*IncidentSearchQuery) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_incidents_v1_incident_query_proto_rawDescGZIP(), []int{2}
}

func (x *IncidentSearchQuery) GetQuery() *wrapperspb.StringValue {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *IncidentSearchQuery) GetField() isIncidentSearchQuery_Field {
	if x != nil {
		return x.Field
	}
	return nil
}

func (x *IncidentSearchQuery) GetIncidentField() IncidentFields {
	if x != nil {
		if x, ok := x.Field.(*IncidentSearchQuery_IncidentField); ok {
			return x.IncidentField
		}
	}
	return IncidentFields_INCIDENTS_FIELDS_UNSPECIFIED
}

func (x *IncidentSearchQuery) GetContextualLabel() *wrapperspb.StringValue {
	if x != nil {
		if x, ok := x.Field.(*IncidentSearchQuery_ContextualLabel); ok {
			return x.ContextualLabel
		}
	}
	return nil
}

type isIncidentSearchQuery_Field interface {
	isIncidentSearchQuery_Field()
}

type IncidentSearchQuery_IncidentField struct {
	IncidentField IncidentFields `protobuf:"varint,10,opt,name=incident_field,json=incidentField,proto3,enum=com.coralogixapis.incidents.v1.IncidentFields,oneof"`
}

type IncidentSearchQuery_ContextualLabel struct {
	ContextualLabel *wrapperspb.StringValue `protobuf:"bytes,11,opt,name=contextual_label,json=contextualLabel,proto3,oneof"`
}

func (*IncidentSearchQuery_IncidentField) isIncidentSearchQuery_Field() {}

func (*IncidentSearchQuery_ContextualLabel) isIncidentSearchQuery_Field() {}

var File_com_coralogixapis_incidents_v1_incident_query_proto protoreflect.FileDescriptor

const file_com_coralogixapis_incidents_v1_incident_query_proto_rawDesc = "" +
	"\n" +
	"3com/coralogixapis/incidents/v1/incident_query.proto\x12\x1ecom.coralogixapis.incidents.v1\x1a-com/coralogixapis/incidents/v1/incident.proto\x1a\x1egoogle/protobuf/wrappers.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\xb4\x02\n" +
	"\aOrderBy\x12W\n" +
	"\x0eincident_field\x18\n" +
	" \x01(\x0e2..com.coralogixapis.incidents.v1.IncidentFieldsH\x00R\rincidentField\x12I\n" +
	"\x10contextual_label\x18\v \x01(\v2\x1c.google.protobuf.StringValueH\x00R\x0fcontextualLabel\x12N\n" +
	"\tdirection\x18\x02 \x01(\x0e20.com.coralogixapis.incidents.v1.OrderByDirectionR\tdirection:,\x92A)\n" +
	"'*\x11Incident order by\xd2\x01\x05field\xd2\x01\tdirectionB\a\n" +
	"\x05field\"\xcd\x03\n" +
	"\aGroupBy\x12\x8a\x01\n" +
	"\x0eincident_field\x18\n" +
	" \x01(\x0e2..com.coralogixapis.incidents.v1.IncidentFieldsB1\x92A.2\x15The field to group byJ\x15\"INCIDENTS_FIELDS_ID\"H\x00R\rincidentField\x12q\n" +
	"\x10contextual_label\x18\v \x01(\v2\x1c.google.protobuf.StringValueB&\x92A#2!The contextual label to group by.H\x00R\x0fcontextualLabel\x12\x96\x01\n" +
	"\x12order_by_direction\x18\x01 \x01(\x0e20.com.coralogixapis.incidents.v1.OrderByDirectionB6\x92A32\x17The order by direction.J\x18\"ORDER_BY_DIRECTION_ASC\"R\x10orderByDirection: \x92A\x1d\n" +
	"\x1b*\x11Incident group by\xd2\x01\x05fieldB\a\n" +
	"\x05field\"\xa6\x03\n" +
	"\x13IncidentSearchQuery\x12R\n" +
	"\x05query\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueB\x1e\x92A\x1b2\x10The search queryJ\a\"error\"R\x05query\x12\x8f\x01\n" +
	"\x0eincident_field\x18\n" +
	" \x01(\x0e2..com.coralogixapis.incidents.v1.IncidentFieldsB6\x92A32\x16The field to search inJ\x19\"INCIDENTS_FIELDS_STATUS\"H\x00R\rincidentField\x12r\n" +
	"\x10contextual_label\x18\v \x01(\v2\x1c.google.protobuf.StringValueB'\x92A$2\"The contextual label to search in.H\x00R\x0fcontextualLabel:,\x92A)\n" +
	"'*\x15Incident search query\xd2\x01\x05query\xd2\x01\x05fieldB\a\n" +
	"\x05field*\xc3\x01\n" +
	"\rOrderByFields\x12\x1f\n" +
	"\x1bORDER_BY_FIELDS_UNSPECIFIED\x10\x00\x12\x16\n" +
	"\x12ORDER_BY_FIELDS_ID\x10\x01\x12\x1c\n" +
	"\x18ORDER_BY_FIELDS_SEVERITY\x10\x02\x12\x18\n" +
	"\x14ORDER_BY_FIELDS_NAME\x10\x03\x12 \n" +
	"\x1cORDER_BY_FIELDS_CREATED_TIME\x10\x04\x12\x1f\n" +
	"\x1bORDER_BY_FIELDS_CLOSED_TIME\x10\x05*o\n" +
	"\x10OrderByDirection\x12\"\n" +
	"\x1eORDER_BY_DIRECTION_UNSPECIFIED\x10\x00\x12\x1a\n" +
	"\x16ORDER_BY_DIRECTION_ASC\x10\x01\x12\x1b\n" +
	"\x17ORDER_BY_DIRECTION_DESC\x10\x02b\x06proto3"

var (
	file_com_coralogixapis_incidents_v1_incident_query_proto_rawDescOnce sync.Once
	file_com_coralogixapis_incidents_v1_incident_query_proto_rawDescData []byte
)

func file_com_coralogixapis_incidents_v1_incident_query_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_incidents_v1_incident_query_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_incidents_v1_incident_query_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_incidents_v1_incident_query_proto_rawDesc), len(file_com_coralogixapis_incidents_v1_incident_query_proto_rawDesc)))
	})
	return file_com_coralogixapis_incidents_v1_incident_query_proto_rawDescData
}

var file_com_coralogixapis_incidents_v1_incident_query_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_com_coralogixapis_incidents_v1_incident_query_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogixapis_incidents_v1_incident_query_proto_goTypes = []any{
	(OrderByFields)(0),             // 0: com.coralogixapis.incidents.v1.OrderByFields
	(OrderByDirection)(0),          // 1: com.coralogixapis.incidents.v1.OrderByDirection
	(*OrderBy)(nil),                // 2: com.coralogixapis.incidents.v1.OrderBy
	(*GroupBy)(nil),                // 3: com.coralogixapis.incidents.v1.GroupBy
	(*IncidentSearchQuery)(nil),    // 4: com.coralogixapis.incidents.v1.IncidentSearchQuery
	(IncidentFields)(0),            // 5: com.coralogixapis.incidents.v1.IncidentFields
	(*wrapperspb.StringValue)(nil), // 6: google.protobuf.StringValue
}
var file_com_coralogixapis_incidents_v1_incident_query_proto_depIdxs = []int32{
	5, // 0: com.coralogixapis.incidents.v1.OrderBy.incident_field:type_name -> com.coralogixapis.incidents.v1.IncidentFields
	6, // 1: com.coralogixapis.incidents.v1.OrderBy.contextual_label:type_name -> google.protobuf.StringValue
	1, // 2: com.coralogixapis.incidents.v1.OrderBy.direction:type_name -> com.coralogixapis.incidents.v1.OrderByDirection
	5, // 3: com.coralogixapis.incidents.v1.GroupBy.incident_field:type_name -> com.coralogixapis.incidents.v1.IncidentFields
	6, // 4: com.coralogixapis.incidents.v1.GroupBy.contextual_label:type_name -> google.protobuf.StringValue
	1, // 5: com.coralogixapis.incidents.v1.GroupBy.order_by_direction:type_name -> com.coralogixapis.incidents.v1.OrderByDirection
	6, // 6: com.coralogixapis.incidents.v1.IncidentSearchQuery.query:type_name -> google.protobuf.StringValue
	5, // 7: com.coralogixapis.incidents.v1.IncidentSearchQuery.incident_field:type_name -> com.coralogixapis.incidents.v1.IncidentFields
	6, // 8: com.coralogixapis.incidents.v1.IncidentSearchQuery.contextual_label:type_name -> google.protobuf.StringValue
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_incidents_v1_incident_query_proto_init() }
func file_com_coralogixapis_incidents_v1_incident_query_proto_init() {
	if File_com_coralogixapis_incidents_v1_incident_query_proto != nil {
		return
	}
	file_com_coralogixapis_incidents_v1_incident_proto_init()
	file_com_coralogixapis_incidents_v1_incident_query_proto_msgTypes[0].OneofWrappers = []any{
		(*OrderBy_IncidentField)(nil),
		(*OrderBy_ContextualLabel)(nil),
	}
	file_com_coralogixapis_incidents_v1_incident_query_proto_msgTypes[1].OneofWrappers = []any{
		(*GroupBy_IncidentField)(nil),
		(*GroupBy_ContextualLabel)(nil),
	}
	file_com_coralogixapis_incidents_v1_incident_query_proto_msgTypes[2].OneofWrappers = []any{
		(*IncidentSearchQuery_IncidentField)(nil),
		(*IncidentSearchQuery_ContextualLabel)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_incidents_v1_incident_query_proto_rawDesc), len(file_com_coralogixapis_incidents_v1_incident_query_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_incidents_v1_incident_query_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_incidents_v1_incident_query_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_incidents_v1_incident_query_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_incidents_v1_incident_query_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_incidents_v1_incident_query_proto = out.File
	file_com_coralogixapis_incidents_v1_incident_query_proto_goTypes = nil
	file_com_coralogixapis_incidents_v1_incident_query_proto_depIdxs = nil
}
