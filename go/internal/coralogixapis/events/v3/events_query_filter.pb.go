// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/events/v3/events_query_filter.proto

package v3

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type FilterOperator int32

const (
	FilterOperator_FILTER_OPERATOR_AND_OR_UNSPECIFIED FilterOperator = 0
	FilterOperator_FILTER_OPERATOR_OR                 FilterOperator = 1
)

// Enum value maps for FilterOperator.
var (
	FilterOperator_name = map[int32]string{
		0: "FILTER_OPERATOR_AND_OR_UNSPECIFIED",
		1: "FILTER_OPERATOR_OR",
	}
	FilterOperator_value = map[string]int32{
		"FILTER_OPERATOR_AND_OR_UNSPECIFIED": 0,
		"FILTER_OPERATOR_OR":                 1,
	}
)

func (x FilterOperator) Enum() *FilterOperator {
	p := new(FilterOperator)
	*p = x
	return p
}

func (x FilterOperator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FilterOperator) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_events_v3_events_query_filter_proto_enumTypes[0].Descriptor()
}

func (FilterOperator) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_events_v3_events_query_filter_proto_enumTypes[0]
}

func (x FilterOperator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FilterOperator.Descriptor instead.
func (FilterOperator) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_events_v3_events_query_filter_proto_rawDescGZIP(), []int{0}
}

type FilterMatcher int32

const (
	FilterMatcher_FILTER_MATCHER_TEXT_OR_UNSPECIFIED FilterMatcher = 0
	FilterMatcher_FILTER_MATCHER_REGEXP              FilterMatcher = 1
)

// Enum value maps for FilterMatcher.
var (
	FilterMatcher_name = map[int32]string{
		0: "FILTER_MATCHER_TEXT_OR_UNSPECIFIED",
		1: "FILTER_MATCHER_REGEXP",
	}
	FilterMatcher_value = map[string]int32{
		"FILTER_MATCHER_TEXT_OR_UNSPECIFIED": 0,
		"FILTER_MATCHER_REGEXP":              1,
	}
)

func (x FilterMatcher) Enum() *FilterMatcher {
	p := new(FilterMatcher)
	*p = x
	return p
}

func (x FilterMatcher) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FilterMatcher) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_events_v3_events_query_filter_proto_enumTypes[1].Descriptor()
}

func (FilterMatcher) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_events_v3_events_query_filter_proto_enumTypes[1]
}

func (x FilterMatcher) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FilterMatcher.Descriptor instead.
func (FilterMatcher) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_events_v3_events_query_filter_proto_rawDescGZIP(), []int{1}
}

type EventsQueryFilter struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Timestamp     *TimestampRange        `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EventsQueryFilter) Reset() {
	*x = EventsQueryFilter{}
	mi := &file_com_coralogixapis_events_v3_events_query_filter_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventsQueryFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventsQueryFilter) ProtoMessage() {}

func (x *EventsQueryFilter) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_events_v3_events_query_filter_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventsQueryFilter.ProtoReflect.Descriptor instead.
func (*EventsQueryFilter) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_events_v3_events_query_filter_proto_rawDescGZIP(), []int{0}
}

func (x *EventsQueryFilter) GetTimestamp() *TimestampRange {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type EventsFilter struct {
	state                  protoimpl.MessageState    `protogen:"open.v1"`
	Timestamp              *TimestampRange           `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	CxEventTypes           []*wrapperspb.StringValue `protobuf:"bytes,2,rep,name=cx_event_types,json=cxEventTypes,proto3" json:"cx_event_types,omitempty"`
	CxEventKeys            []*wrapperspb.StringValue `protobuf:"bytes,4,rep,name=cx_event_keys,json=cxEventKeys,proto3" json:"cx_event_keys,omitempty"`
	CxEventMetadataFilters *Filters                  `protobuf:"bytes,5,opt,name=cx_event_metadata_filters,json=cxEventMetadataFilters,proto3" json:"cx_event_metadata_filters,omitempty"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *EventsFilter) Reset() {
	*x = EventsFilter{}
	mi := &file_com_coralogixapis_events_v3_events_query_filter_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventsFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventsFilter) ProtoMessage() {}

func (x *EventsFilter) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_events_v3_events_query_filter_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventsFilter.ProtoReflect.Descriptor instead.
func (*EventsFilter) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_events_v3_events_query_filter_proto_rawDescGZIP(), []int{1}
}

func (x *EventsFilter) GetTimestamp() *TimestampRange {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *EventsFilter) GetCxEventTypes() []*wrapperspb.StringValue {
	if x != nil {
		return x.CxEventTypes
	}
	return nil
}

func (x *EventsFilter) GetCxEventKeys() []*wrapperspb.StringValue {
	if x != nil {
		return x.CxEventKeys
	}
	return nil
}

func (x *EventsFilter) GetCxEventMetadataFilters() *Filters {
	if x != nil {
		return x.CxEventMetadataFilters
	}
	return nil
}

type Filters struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Operator      FilterOperator         `protobuf:"varint,1,opt,name=operator,proto3,enum=com.coralogixapis.events.v3.FilterOperator" json:"operator,omitempty"`
	PathAndValues []*FilterPathAndValues `protobuf:"bytes,2,rep,name=path_and_values,json=pathAndValues,proto3" json:"path_and_values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Filters) Reset() {
	*x = Filters{}
	mi := &file_com_coralogixapis_events_v3_events_query_filter_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Filters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filters) ProtoMessage() {}

func (x *Filters) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_events_v3_events_query_filter_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filters.ProtoReflect.Descriptor instead.
func (*Filters) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_events_v3_events_query_filter_proto_rawDescGZIP(), []int{2}
}

func (x *Filters) GetOperator() FilterOperator {
	if x != nil {
		return x.Operator
	}
	return FilterOperator_FILTER_OPERATOR_AND_OR_UNSPECIFIED
}

func (x *Filters) GetPathAndValues() []*FilterPathAndValues {
	if x != nil {
		return x.PathAndValues
	}
	return nil
}

type FilterPathAndValues struct {
	state protoimpl.MessageState  `protogen:"open.v1"`
	Path  *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Types that are valid to be assigned to Values:
	//
	//	*FilterPathAndValues_MultipleValues
	//	*FilterPathAndValues_Filters
	Values        isFilterPathAndValues_Values `protobuf_oneof:"values"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FilterPathAndValues) Reset() {
	*x = FilterPathAndValues{}
	mi := &file_com_coralogixapis_events_v3_events_query_filter_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FilterPathAndValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterPathAndValues) ProtoMessage() {}

func (x *FilterPathAndValues) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_events_v3_events_query_filter_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterPathAndValues.ProtoReflect.Descriptor instead.
func (*FilterPathAndValues) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_events_v3_events_query_filter_proto_rawDescGZIP(), []int{3}
}

func (x *FilterPathAndValues) GetPath() *wrapperspb.StringValue {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *FilterPathAndValues) GetValues() isFilterPathAndValues_Values {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *FilterPathAndValues) GetMultipleValues() *MultipleValues {
	if x != nil {
		if x, ok := x.Values.(*FilterPathAndValues_MultipleValues); ok {
			return x.MultipleValues
		}
	}
	return nil
}

func (x *FilterPathAndValues) GetFilters() *Filters {
	if x != nil {
		if x, ok := x.Values.(*FilterPathAndValues_Filters); ok {
			return x.Filters
		}
	}
	return nil
}

type isFilterPathAndValues_Values interface {
	isFilterPathAndValues_Values()
}

type FilterPathAndValues_MultipleValues struct {
	MultipleValues *MultipleValues `protobuf:"bytes,100,opt,name=multiple_values,json=multipleValues,proto3,oneof"`
}

type FilterPathAndValues_Filters struct {
	Filters *Filters `protobuf:"bytes,101,opt,name=filters,proto3,oneof"`
}

func (*FilterPathAndValues_MultipleValues) isFilterPathAndValues_Values() {}

func (*FilterPathAndValues_Filters) isFilterPathAndValues_Values() {}

type MultipleValues struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Values        []*wrapperspb.StringValue `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	Matcher       FilterMatcher             `protobuf:"varint,2,opt,name=matcher,proto3,enum=com.coralogixapis.events.v3.FilterMatcher" json:"matcher,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MultipleValues) Reset() {
	*x = MultipleValues{}
	mi := &file_com_coralogixapis_events_v3_events_query_filter_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MultipleValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultipleValues) ProtoMessage() {}

func (x *MultipleValues) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_events_v3_events_query_filter_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultipleValues.ProtoReflect.Descriptor instead.
func (*MultipleValues) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_events_v3_events_query_filter_proto_rawDescGZIP(), []int{4}
}

func (x *MultipleValues) GetValues() []*wrapperspb.StringValue {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *MultipleValues) GetMatcher() FilterMatcher {
	if x != nil {
		return x.Matcher
	}
	return FilterMatcher_FILTER_MATCHER_TEXT_OR_UNSPECIFIED
}

type TimestampRange struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	From          *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To            *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TimestampRange) Reset() {
	*x = TimestampRange{}
	mi := &file_com_coralogixapis_events_v3_events_query_filter_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimestampRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimestampRange) ProtoMessage() {}

func (x *TimestampRange) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_events_v3_events_query_filter_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimestampRange.ProtoReflect.Descriptor instead.
func (*TimestampRange) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_events_v3_events_query_filter_proto_rawDescGZIP(), []int{5}
}

func (x *TimestampRange) GetFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *TimestampRange) GetTo() *timestamppb.Timestamp {
	if x != nil {
		return x.To
	}
	return nil
}

var File_com_coralogixapis_events_v3_events_query_filter_proto protoreflect.FileDescriptor

const file_com_coralogixapis_events_v3_events_query_filter_proto_rawDesc = "" +
	"\n" +
	"5com/coralogixapis/events/v3/events_query_filter.proto\x12\x1bcom.coralogixapis.events.v3\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/wrappers.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"^\n" +
	"\x11EventsQueryFilter\x12I\n" +
	"\ttimestamp\x18\x01 \x01(\v2+.com.coralogixapis.events.v3.TimestampRangeR\ttimestamp\"\xda\x03\n" +
	"\fEventsFilter\x12I\n" +
	"\ttimestamp\x18\x01 \x01(\v2+.com.coralogixapis.events.v3.TimestampRangeR\ttimestamp\x12V\n" +
	"\x0ecx_event_types\x18\x02 \x03(\v2\x1c.google.protobuf.StringValueB\x12\x92A\x0fJ\r[\"test_type\"]R\fcxEventTypes\x12S\n" +
	"\rcx_event_keys\x18\x04 \x03(\v2\x1c.google.protobuf.StringValueB\x11\x92A\x0eJ\f[\"test_key\"]R\vcxEventKeys\x12_\n" +
	"\x19cx_event_metadata_filters\x18\x05 \x01(\v2$.com.coralogixapis.events.v3.FiltersR\x16cxEventMetadataFilters:q\x92An\n" +
	"l*\fEventsFilter2/This data structure represents an events filter\xd2\x01\ttimestamp\xd2\x01\x0ecx_event_types\xd2\x01\rcx_event_keys\"\x92\x02\n" +
	"\aFilters\x12b\n" +
	"\boperator\x18\x01 \x01(\x0e2+.com.coralogixapis.events.v3.FilterOperatorB\x19\x92A\x16J\x14\"FILTER_OPERATOR_OR\"R\boperator\x12X\n" +
	"\x0fpath_and_values\x18\x02 \x03(\v20.com.coralogixapis.events.v3.FilterPathAndValuesR\rpathAndValues:I\x92AF\n" +
	"D*\aFilters2'This data structure represents a filter\xd2\x01\x0fpath_and_values\"\xdd\x02\n" +
	"\x13FilterPathAndValues\x12=\n" +
	"\x04path\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueB\v\x92A\bJ\x06\"test\"R\x04path\x12V\n" +
	"\x0fmultiple_values\x18d \x01(\v2+.com.coralogixapis.events.v3.MultipleValuesH\x00R\x0emultipleValues\x12@\n" +
	"\afilters\x18e \x01(\v2$.com.coralogixapis.events.v3.FiltersH\x00R\afilters:c\x92A`\n" +
	"^*\x13FilterPathAndValues27This data structure represents a filter path and values\xd2\x01\x04path\xd2\x01\x06valuesB\b\n" +
	"\x06values\"\x8c\x01\n" +
	"\x0eMultipleValues\x124\n" +
	"\x06values\x18\x01 \x03(\v2\x1c.google.protobuf.StringValueR\x06values\x12D\n" +
	"\amatcher\x18\x02 \x01(\x0e2*.com.coralogixapis.events.v3.FilterMatcherR\amatcher\"\xc1\x01\n" +
	"\x0eTimestampRange\x12.\n" +
	"\x04from\x18\x01 \x01(\v2\x1a.google.protobuf.TimestampR\x04from\x12*\n" +
	"\x02to\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\x02to:S\x92AP\n" +
	"N*\x0eTimestampRange20This data structure represents a timestamp range\xd2\x01\x04from\xd2\x01\x02to*P\n" +
	"\x0eFilterOperator\x12&\n" +
	"\"FILTER_OPERATOR_AND_OR_UNSPECIFIED\x10\x00\x12\x16\n" +
	"\x12FILTER_OPERATOR_OR\x10\x01*R\n" +
	"\rFilterMatcher\x12&\n" +
	"\"FILTER_MATCHER_TEXT_OR_UNSPECIFIED\x10\x00\x12\x19\n" +
	"\x15FILTER_MATCHER_REGEXP\x10\x01b\x06proto3"

var (
	file_com_coralogixapis_events_v3_events_query_filter_proto_rawDescOnce sync.Once
	file_com_coralogixapis_events_v3_events_query_filter_proto_rawDescData []byte
)

func file_com_coralogixapis_events_v3_events_query_filter_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_events_v3_events_query_filter_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_events_v3_events_query_filter_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_events_v3_events_query_filter_proto_rawDesc), len(file_com_coralogixapis_events_v3_events_query_filter_proto_rawDesc)))
	})
	return file_com_coralogixapis_events_v3_events_query_filter_proto_rawDescData
}

var file_com_coralogixapis_events_v3_events_query_filter_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_com_coralogixapis_events_v3_events_query_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_com_coralogixapis_events_v3_events_query_filter_proto_goTypes = []any{
	(FilterOperator)(0),            // 0: com.coralogixapis.events.v3.FilterOperator
	(FilterMatcher)(0),             // 1: com.coralogixapis.events.v3.FilterMatcher
	(*EventsQueryFilter)(nil),      // 2: com.coralogixapis.events.v3.EventsQueryFilter
	(*EventsFilter)(nil),           // 3: com.coralogixapis.events.v3.EventsFilter
	(*Filters)(nil),                // 4: com.coralogixapis.events.v3.Filters
	(*FilterPathAndValues)(nil),    // 5: com.coralogixapis.events.v3.FilterPathAndValues
	(*MultipleValues)(nil),         // 6: com.coralogixapis.events.v3.MultipleValues
	(*TimestampRange)(nil),         // 7: com.coralogixapis.events.v3.TimestampRange
	(*wrapperspb.StringValue)(nil), // 8: google.protobuf.StringValue
	(*timestamppb.Timestamp)(nil),  // 9: google.protobuf.Timestamp
}
var file_com_coralogixapis_events_v3_events_query_filter_proto_depIdxs = []int32{
	7,  // 0: com.coralogixapis.events.v3.EventsQueryFilter.timestamp:type_name -> com.coralogixapis.events.v3.TimestampRange
	7,  // 1: com.coralogixapis.events.v3.EventsFilter.timestamp:type_name -> com.coralogixapis.events.v3.TimestampRange
	8,  // 2: com.coralogixapis.events.v3.EventsFilter.cx_event_types:type_name -> google.protobuf.StringValue
	8,  // 3: com.coralogixapis.events.v3.EventsFilter.cx_event_keys:type_name -> google.protobuf.StringValue
	4,  // 4: com.coralogixapis.events.v3.EventsFilter.cx_event_metadata_filters:type_name -> com.coralogixapis.events.v3.Filters
	0,  // 5: com.coralogixapis.events.v3.Filters.operator:type_name -> com.coralogixapis.events.v3.FilterOperator
	5,  // 6: com.coralogixapis.events.v3.Filters.path_and_values:type_name -> com.coralogixapis.events.v3.FilterPathAndValues
	8,  // 7: com.coralogixapis.events.v3.FilterPathAndValues.path:type_name -> google.protobuf.StringValue
	6,  // 8: com.coralogixapis.events.v3.FilterPathAndValues.multiple_values:type_name -> com.coralogixapis.events.v3.MultipleValues
	4,  // 9: com.coralogixapis.events.v3.FilterPathAndValues.filters:type_name -> com.coralogixapis.events.v3.Filters
	8,  // 10: com.coralogixapis.events.v3.MultipleValues.values:type_name -> google.protobuf.StringValue
	1,  // 11: com.coralogixapis.events.v3.MultipleValues.matcher:type_name -> com.coralogixapis.events.v3.FilterMatcher
	9,  // 12: com.coralogixapis.events.v3.TimestampRange.from:type_name -> google.protobuf.Timestamp
	9,  // 13: com.coralogixapis.events.v3.TimestampRange.to:type_name -> google.protobuf.Timestamp
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_events_v3_events_query_filter_proto_init() }
func file_com_coralogixapis_events_v3_events_query_filter_proto_init() {
	if File_com_coralogixapis_events_v3_events_query_filter_proto != nil {
		return
	}
	file_com_coralogixapis_events_v3_events_query_filter_proto_msgTypes[3].OneofWrappers = []any{
		(*FilterPathAndValues_MultipleValues)(nil),
		(*FilterPathAndValues_Filters)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_events_v3_events_query_filter_proto_rawDesc), len(file_com_coralogixapis_events_v3_events_query_filter_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_events_v3_events_query_filter_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_events_v3_events_query_filter_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_events_v3_events_query_filter_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_events_v3_events_query_filter_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_events_v3_events_query_filter_proto = out.File
	file_com_coralogixapis_events_v3_events_query_filter_proto_goTypes = nil
	file_com_coralogixapis_events_v3_events_query_filter_proto_depIdxs = nil
}
