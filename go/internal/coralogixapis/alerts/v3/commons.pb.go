// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/alerts/v3/commons.proto

package v3

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
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

type NotifyOn int32

const (
	NotifyOn_NOTIFY_ON_TRIGGERED_ONLY_UNSPECIFIED NotifyOn = 0
	NotifyOn_NOTIFY_ON_TRIGGERED_AND_RESOLVED     NotifyOn = 1
)

// Enum value maps for NotifyOn.
var (
	NotifyOn_name = map[int32]string{
		0: "NOTIFY_ON_TRIGGERED_ONLY_UNSPECIFIED",
		1: "NOTIFY_ON_TRIGGERED_AND_RESOLVED",
	}
	NotifyOn_value = map[string]int32{
		"NOTIFY_ON_TRIGGERED_ONLY_UNSPECIFIED": 0,
		"NOTIFY_ON_TRIGGERED_AND_RESOLVED":     1,
	}
)

func (x NotifyOn) Enum() *NotifyOn {
	p := new(NotifyOn)
	*p = x
	return p
}

func (x NotifyOn) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotifyOn) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_alerts_v3_commons_proto_enumTypes[0].Descriptor()
}

func (NotifyOn) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_alerts_v3_commons_proto_enumTypes[0]
}

func (x NotifyOn) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotifyOn.Descriptor instead.
func (NotifyOn) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_commons_proto_rawDescGZIP(), []int{0}
}

type DurationUnit int32

const (
	DurationUnit_DURATION_UNIT_UNSPECIFIED DurationUnit = 0
	DurationUnit_DURATION_UNIT_HOURS       DurationUnit = 1
)

// Enum value maps for DurationUnit.
var (
	DurationUnit_name = map[int32]string{
		0: "DURATION_UNIT_UNSPECIFIED",
		1: "DURATION_UNIT_HOURS",
	}
	DurationUnit_value = map[string]int32{
		"DURATION_UNIT_UNSPECIFIED": 0,
		"DURATION_UNIT_HOURS":       1,
	}
)

func (x DurationUnit) Enum() *DurationUnit {
	p := new(DurationUnit)
	*p = x
	return p
}

func (x DurationUnit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DurationUnit) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_alerts_v3_commons_proto_enumTypes[1].Descriptor()
}

func (DurationUnit) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_alerts_v3_commons_proto_enumTypes[1]
}

func (x DurationUnit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DurationUnit.Descriptor instead.
func (DurationUnit) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_commons_proto_rawDescGZIP(), []int{1}
}

type BooleanOperator int32

const (
	BooleanOperator_BOOLEAN_OPERATOR_AND_UNSPECIFIED BooleanOperator = 0
	BooleanOperator_BOOLEAN_OPERATOR_OR              BooleanOperator = 1
)

// Enum value maps for BooleanOperator.
var (
	BooleanOperator_name = map[int32]string{
		0: "BOOLEAN_OPERATOR_AND_UNSPECIFIED",
		1: "BOOLEAN_OPERATOR_OR",
	}
	BooleanOperator_value = map[string]int32{
		"BOOLEAN_OPERATOR_AND_UNSPECIFIED": 0,
		"BOOLEAN_OPERATOR_OR":              1,
	}
)

func (x BooleanOperator) Enum() *BooleanOperator {
	p := new(BooleanOperator)
	*p = x
	return p
}

func (x BooleanOperator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BooleanOperator) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_alerts_v3_commons_proto_enumTypes[2].Descriptor()
}

func (BooleanOperator) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_alerts_v3_commons_proto_enumTypes[2]
}

func (x BooleanOperator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BooleanOperator.Descriptor instead.
func (BooleanOperator) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_commons_proto_rawDescGZIP(), []int{2}
}

type OrderByDirection int32

const (
	OrderByDirection_ORDER_BY_DIRECTION_ASC_OR_UNSPECIFIED OrderByDirection = 0
	OrderByDirection_ORDER_BY_DIRECTION_DESC               OrderByDirection = 1
)

// Enum value maps for OrderByDirection.
var (
	OrderByDirection_name = map[int32]string{
		0: "ORDER_BY_DIRECTION_ASC_OR_UNSPECIFIED",
		1: "ORDER_BY_DIRECTION_DESC",
	}
	OrderByDirection_value = map[string]int32{
		"ORDER_BY_DIRECTION_ASC_OR_UNSPECIFIED": 0,
		"ORDER_BY_DIRECTION_DESC":               1,
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
	return file_com_coralogixapis_alerts_v3_commons_proto_enumTypes[3].Descriptor()
}

func (OrderByDirection) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_alerts_v3_commons_proto_enumTypes[3]
}

func (x OrderByDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderByDirection.Descriptor instead.
func (OrderByDirection) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_commons_proto_rawDescGZIP(), []int{3}
}

type OrderByFields int32

const (
	OrderByFields_ORDER_BY_FIELDS_UNSPECIFIED    OrderByFields = 0
	OrderByFields_ORDER_BY_FIELDS_NAME           OrderByFields = 1
	OrderByFields_ORDER_BY_FIELDS_ID             OrderByFields = 2
	OrderByFields_ORDER_BY_FIELDS_SEVERITY       OrderByFields = 3
	OrderByFields_ORDER_BY_FIELDS_CREATED_TIME   OrderByFields = 4
	OrderByFields_ORDER_BY_FIELDS_UPDATED_TIME   OrderByFields = 5
	OrderByFields_ORDER_BY_FIELDS_LAST_TRIGGERED OrderByFields = 6
)

// Enum value maps for OrderByFields.
var (
	OrderByFields_name = map[int32]string{
		0: "ORDER_BY_FIELDS_UNSPECIFIED",
		1: "ORDER_BY_FIELDS_NAME",
		2: "ORDER_BY_FIELDS_ID",
		3: "ORDER_BY_FIELDS_SEVERITY",
		4: "ORDER_BY_FIELDS_CREATED_TIME",
		5: "ORDER_BY_FIELDS_UPDATED_TIME",
		6: "ORDER_BY_FIELDS_LAST_TRIGGERED",
	}
	OrderByFields_value = map[string]int32{
		"ORDER_BY_FIELDS_UNSPECIFIED":    0,
		"ORDER_BY_FIELDS_NAME":           1,
		"ORDER_BY_FIELDS_ID":             2,
		"ORDER_BY_FIELDS_SEVERITY":       3,
		"ORDER_BY_FIELDS_CREATED_TIME":   4,
		"ORDER_BY_FIELDS_UPDATED_TIME":   5,
		"ORDER_BY_FIELDS_LAST_TRIGGERED": 6,
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
	return file_com_coralogixapis_alerts_v3_commons_proto_enumTypes[4].Descriptor()
}

func (OrderByFields) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_alerts_v3_commons_proto_enumTypes[4]
}

func (x OrderByFields) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderByFields.Descriptor instead.
func (OrderByFields) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_commons_proto_rawDescGZIP(), []int{4}
}

type TimeDuration struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Duration      *wrapperspb.UInt64Value `protobuf:"bytes,1,opt,name=duration,proto3" json:"duration,omitempty"`
	Unit          DurationUnit            `protobuf:"varint,2,opt,name=unit,proto3,enum=com.coralogixapis.alerts.v3.DurationUnit" json:"unit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TimeDuration) Reset() {
	*x = TimeDuration{}
	mi := &file_com_coralogixapis_alerts_v3_commons_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimeDuration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeDuration) ProtoMessage() {}

func (x *TimeDuration) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_commons_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeDuration.ProtoReflect.Descriptor instead.
func (*TimeDuration) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_commons_proto_rawDescGZIP(), []int{0}
}

func (x *TimeDuration) GetDuration() *wrapperspb.UInt64Value {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *TimeDuration) GetUnit() DurationUnit {
	if x != nil {
		return x.Unit
	}
	return DurationUnit_DURATION_UNIT_UNSPECIFIED
}

type OrderByList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderBys      []*OrderBy             `protobuf:"bytes,1,rep,name=order_bys,json=orderBys,proto3" json:"order_bys,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderByList) Reset() {
	*x = OrderByList{}
	mi := &file_com_coralogixapis_alerts_v3_commons_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderByList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderByList) ProtoMessage() {}

func (x *OrderByList) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_commons_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderByList.ProtoReflect.Descriptor instead.
func (*OrderByList) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_commons_proto_rawDescGZIP(), []int{1}
}

func (x *OrderByList) GetOrderBys() []*OrderBy {
	if x != nil {
		return x.OrderBys
	}
	return nil
}

type OrderBy struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FieldName     OrderByFields          `protobuf:"varint,1,opt,name=field_name,json=fieldName,proto3,enum=com.coralogixapis.alerts.v3.OrderByFields" json:"field_name,omitempty"`
	Direction     OrderByDirection       `protobuf:"varint,2,opt,name=direction,proto3,enum=com.coralogixapis.alerts.v3.OrderByDirection" json:"direction,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderBy) Reset() {
	*x = OrderBy{}
	mi := &file_com_coralogixapis_alerts_v3_commons_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderBy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderBy) ProtoMessage() {}

func (x *OrderBy) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_commons_proto_msgTypes[2]
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
	return file_com_coralogixapis_alerts_v3_commons_proto_rawDescGZIP(), []int{2}
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
	return OrderByDirection_ORDER_BY_DIRECTION_ASC_OR_UNSPECIFIED
}

var File_com_coralogixapis_alerts_v3_commons_proto protoreflect.FileDescriptor

const file_com_coralogixapis_alerts_v3_commons_proto_rawDesc = "" +
	"\n" +
	")com/coralogixapis/alerts/v3/commons.proto\x12\x1bcom.coralogixapis.alerts.v3\x1a google/protobuf/descriptor.proto\x1a\x1egoogle/protobuf/wrappers.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\xa3\x02\n" +
	"\fTimeDuration\x12U\n" +
	"\bduration\x18\x01 \x01(\v2\x1c.google.protobuf.UInt64ValueB\x1b\x92A\x182\x12The duration valueJ\x0260R\bduration\x12s\n" +
	"\x04unit\x18\x02 \x01(\x0e2).com.coralogixapis.alerts.v3.DurationUnitB4\x92A12\x18The unit of the durationJ\x15\"DURATION_UNIT_HOURS\"R\x04unit:G\x92AD\n" +
	"B*\rTime duration2\x1fConfiguration for time duration\xd2\x01\bduration\xd2\x01\x04unit\"\xa4\x01\n" +
	"\vOrderByList\x12W\n" +
	"\torder_bys\x18\x01 \x03(\v2$.com.coralogixapis.alerts.v3.OrderByB\x14\x92A\x112\x0fOrder by fieldsR\borderBys:<\x92A9\n" +
	"7*\rOrder by list2\x1aList of fields to order by\xd2\x01\torder_bys\"\xa4\x03\n" +
	"\aOrderBy\x12\x86\x01\n" +
	"\n" +
	"field_name\x18\x01 \x01(\x0e2*.com.coralogixapis.alerts.v3.OrderByFieldsB;\x92A82\x16Field name to order byJ\x1e\"ORDER_BY_FIELDS_UPDATED_TIME\"R\tfieldName\x12\x91\x01\n" +
	"\tdirection\x18\x02 \x01(\x0e2-.com.coralogixapis.alerts.v3.OrderByDirectionBD\x92AA2\x16Direction for orderingJ'\"ORDER_BY_DIRECTION_ASC_OR_UNSPECIFIED\"R\tdirection:|\x92Ay\n" +
	"w*\bOrder by2RA data structure that specifies the field and direction for ordering query results\xd2\x01\n" +
	"field_name\xd2\x01\tdirection*Z\n" +
	"\bNotifyOn\x12(\n" +
	"$NOTIFY_ON_TRIGGERED_ONLY_UNSPECIFIED\x10\x00\x12$\n" +
	" NOTIFY_ON_TRIGGERED_AND_RESOLVED\x10\x01*F\n" +
	"\fDurationUnit\x12\x1d\n" +
	"\x19DURATION_UNIT_UNSPECIFIED\x10\x00\x12\x17\n" +
	"\x13DURATION_UNIT_HOURS\x10\x01*P\n" +
	"\x0fBooleanOperator\x12$\n" +
	" BOOLEAN_OPERATOR_AND_UNSPECIFIED\x10\x00\x12\x17\n" +
	"\x13BOOLEAN_OPERATOR_OR\x10\x01*Z\n" +
	"\x10OrderByDirection\x12)\n" +
	"%ORDER_BY_DIRECTION_ASC_OR_UNSPECIFIED\x10\x00\x12\x1b\n" +
	"\x17ORDER_BY_DIRECTION_DESC\x10\x01*\xe8\x01\n" +
	"\rOrderByFields\x12\x1f\n" +
	"\x1bORDER_BY_FIELDS_UNSPECIFIED\x10\x00\x12\x18\n" +
	"\x14ORDER_BY_FIELDS_NAME\x10\x01\x12\x16\n" +
	"\x12ORDER_BY_FIELDS_ID\x10\x02\x12\x1c\n" +
	"\x18ORDER_BY_FIELDS_SEVERITY\x10\x03\x12 \n" +
	"\x1cORDER_BY_FIELDS_CREATED_TIME\x10\x04\x12 \n" +
	"\x1cORDER_BY_FIELDS_UPDATED_TIME\x10\x05\x12\"\n" +
	"\x1eORDER_BY_FIELDS_LAST_TRIGGERED\x10\x06b\x06proto3"

var (
	file_com_coralogixapis_alerts_v3_commons_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_commons_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v3_commons_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_commons_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_commons_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_commons_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_commons_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerts_v3_commons_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_commons_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_com_coralogixapis_alerts_v3_commons_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogixapis_alerts_v3_commons_proto_goTypes = []any{
	(NotifyOn)(0),                  // 0: com.coralogixapis.alerts.v3.NotifyOn
	(DurationUnit)(0),              // 1: com.coralogixapis.alerts.v3.DurationUnit
	(BooleanOperator)(0),           // 2: com.coralogixapis.alerts.v3.BooleanOperator
	(OrderByDirection)(0),          // 3: com.coralogixapis.alerts.v3.OrderByDirection
	(OrderByFields)(0),             // 4: com.coralogixapis.alerts.v3.OrderByFields
	(*TimeDuration)(nil),           // 5: com.coralogixapis.alerts.v3.TimeDuration
	(*OrderByList)(nil),            // 6: com.coralogixapis.alerts.v3.OrderByList
	(*OrderBy)(nil),                // 7: com.coralogixapis.alerts.v3.OrderBy
	(*wrapperspb.UInt64Value)(nil), // 8: google.protobuf.UInt64Value
}
var file_com_coralogixapis_alerts_v3_commons_proto_depIdxs = []int32{
	8, // 0: com.coralogixapis.alerts.v3.TimeDuration.duration:type_name -> google.protobuf.UInt64Value
	1, // 1: com.coralogixapis.alerts.v3.TimeDuration.unit:type_name -> com.coralogixapis.alerts.v3.DurationUnit
	7, // 2: com.coralogixapis.alerts.v3.OrderByList.order_bys:type_name -> com.coralogixapis.alerts.v3.OrderBy
	4, // 3: com.coralogixapis.alerts.v3.OrderBy.field_name:type_name -> com.coralogixapis.alerts.v3.OrderByFields
	3, // 4: com.coralogixapis.alerts.v3.OrderBy.direction:type_name -> com.coralogixapis.alerts.v3.OrderByDirection
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_alerts_v3_commons_proto_init() }
func file_com_coralogixapis_alerts_v3_commons_proto_init() {
	if File_com_coralogixapis_alerts_v3_commons_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_commons_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_commons_proto_rawDesc)),
			NumEnums:      5,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_commons_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_commons_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_alerts_v3_commons_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_alerts_v3_commons_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_commons_proto = out.File
	file_com_coralogixapis_alerts_v3_commons_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_commons_proto_depIdxs = nil
}
