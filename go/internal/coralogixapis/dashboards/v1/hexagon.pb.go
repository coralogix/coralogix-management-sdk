// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/dashboards/v1/ast/widgets/hexagon.proto

package v1

import (
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

type Hexagon struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Query         *Hexagon_Query          `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Min           *wrapperspb.DoubleValue `protobuf:"bytes,2,opt,name=min,proto3" json:"min,omitempty"`
	Max           *wrapperspb.DoubleValue `protobuf:"bytes,3,opt,name=max,proto3" json:"max,omitempty"`
	Unit          Unit                    `protobuf:"varint,6,opt,name=unit,proto3,enum=com.coralogixapis.dashboards.v1.ast.widgets.common.Unit" json:"unit,omitempty"`
	Thresholds    []*Threshold            `protobuf:"bytes,7,rep,name=thresholds,proto3" json:"thresholds,omitempty"`
	DataModeType  DataModeType            `protobuf:"varint,8,opt,name=data_mode_type,json=dataModeType,proto3,enum=com.coralogixapis.dashboards.v1.ast.widgets.common.DataModeType" json:"data_mode_type,omitempty"`
	CustomUnit    *wrapperspb.StringValue `protobuf:"bytes,10,opt,name=custom_unit,json=customUnit,proto3" json:"custom_unit,omitempty"`
	Decimal       *wrapperspb.Int32Value  `protobuf:"bytes,11,opt,name=decimal,proto3" json:"decimal,omitempty"`
	ThresholdType ThresholdType           `protobuf:"varint,12,opt,name=threshold_type,json=thresholdType,proto3,enum=com.coralogixapis.dashboards.v1.ast.widgets.common.ThresholdType" json:"threshold_type,omitempty"`
	Legend        *Legend                 `protobuf:"bytes,13,opt,name=legend,proto3" json:"legend,omitempty"`
	LegendBy      LegendBy                `protobuf:"varint,14,opt,name=legend_by,json=legendBy,proto3,enum=com.coralogixapis.dashboards.v1.ast.widgets.common.LegendBy" json:"legend_by,omitempty"`
	CustomLinks   []*CustomLink           `protobuf:"bytes,15,rep,name=custom_links,json=customLinks,proto3" json:"custom_links,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Hexagon) Reset() {
	*x = Hexagon{}
	mi := &file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Hexagon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hexagon) ProtoMessage() {}

func (x *Hexagon) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hexagon.ProtoReflect.Descriptor instead.
func (*Hexagon) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_rawDescGZIP(), []int{0}
}

func (x *Hexagon) GetQuery() *Hexagon_Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *Hexagon) GetMin() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Min
	}
	return nil
}

func (x *Hexagon) GetMax() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Max
	}
	return nil
}

func (x *Hexagon) GetUnit() Unit {
	if x != nil {
		return x.Unit
	}
	return Unit_UNIT_UNSPECIFIED
}

func (x *Hexagon) GetThresholds() []*Threshold {
	if x != nil {
		return x.Thresholds
	}
	return nil
}

func (x *Hexagon) GetDataModeType() DataModeType {
	if x != nil {
		return x.DataModeType
	}
	return DataModeType_DATA_MODE_TYPE_HIGH_UNSPECIFIED
}

func (x *Hexagon) GetCustomUnit() *wrapperspb.StringValue {
	if x != nil {
		return x.CustomUnit
	}
	return nil
}

func (x *Hexagon) GetDecimal() *wrapperspb.Int32Value {
	if x != nil {
		return x.Decimal
	}
	return nil
}

func (x *Hexagon) GetThresholdType() ThresholdType {
	if x != nil {
		return x.ThresholdType
	}
	return ThresholdType_THRESHOLD_TYPE_UNSPECIFIED
}

func (x *Hexagon) GetLegend() *Legend {
	if x != nil {
		return x.Legend
	}
	return nil
}

func (x *Hexagon) GetLegendBy() LegendBy {
	if x != nil {
		return x.LegendBy
	}
	return LegendBy_LEGEND_BY_UNSPECIFIED
}

func (x *Hexagon) GetCustomLinks() []*CustomLink {
	if x != nil {
		return x.CustomLinks
	}
	return nil
}

type Hexagon_Query struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Value:
	//
	//	*Hexagon_Query_Metrics
	//	*Hexagon_Query_Logs
	//	*Hexagon_Query_Spans
	//	*Hexagon_Query_Dataprime
	Value         isHexagon_Query_Value `protobuf_oneof:"value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Hexagon_Query) Reset() {
	*x = Hexagon_Query{}
	mi := &file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Hexagon_Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hexagon_Query) ProtoMessage() {}

func (x *Hexagon_Query) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hexagon_Query.ProtoReflect.Descriptor instead.
func (*Hexagon_Query) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Hexagon_Query) GetValue() isHexagon_Query_Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Hexagon_Query) GetMetrics() *Hexagon_MetricsQuery {
	if x != nil {
		if x, ok := x.Value.(*Hexagon_Query_Metrics); ok {
			return x.Metrics
		}
	}
	return nil
}

func (x *Hexagon_Query) GetLogs() *Hexagon_LogsQuery {
	if x != nil {
		if x, ok := x.Value.(*Hexagon_Query_Logs); ok {
			return x.Logs
		}
	}
	return nil
}

func (x *Hexagon_Query) GetSpans() *Hexagon_SpansQuery {
	if x != nil {
		if x, ok := x.Value.(*Hexagon_Query_Spans); ok {
			return x.Spans
		}
	}
	return nil
}

func (x *Hexagon_Query) GetDataprime() *Hexagon_DataprimeQuery {
	if x != nil {
		if x, ok := x.Value.(*Hexagon_Query_Dataprime); ok {
			return x.Dataprime
		}
	}
	return nil
}

type isHexagon_Query_Value interface {
	isHexagon_Query_Value()
}

type Hexagon_Query_Metrics struct {
	Metrics *Hexagon_MetricsQuery `protobuf:"bytes,1,opt,name=metrics,proto3,oneof"`
}

type Hexagon_Query_Logs struct {
	Logs *Hexagon_LogsQuery `protobuf:"bytes,2,opt,name=logs,proto3,oneof"`
}

type Hexagon_Query_Spans struct {
	Spans *Hexagon_SpansQuery `protobuf:"bytes,3,opt,name=spans,proto3,oneof"`
}

type Hexagon_Query_Dataprime struct {
	Dataprime *Hexagon_DataprimeQuery `protobuf:"bytes,4,opt,name=dataprime,proto3,oneof"`
}

func (*Hexagon_Query_Metrics) isHexagon_Query_Value() {}

func (*Hexagon_Query_Logs) isHexagon_Query_Value() {}

func (*Hexagon_Query_Spans) isHexagon_Query_Value() {}

func (*Hexagon_Query_Dataprime) isHexagon_Query_Value() {}

type Hexagon_MetricsQuery struct {
	state           protoimpl.MessageState  `protogen:"open.v1"`
	PromqlQuery     *PromQlQuery            `protobuf:"bytes,1,opt,name=promql_query,json=promqlQuery,proto3" json:"promql_query,omitempty"`
	Filters         []*Filter_MetricsFilter `protobuf:"bytes,2,rep,name=filters,proto3" json:"filters,omitempty"`
	EditorMode      MetricsQueryEditorMode  `protobuf:"varint,3,opt,name=editor_mode,json=editorMode,proto3,enum=com.coralogixapis.dashboards.v1.ast.widgets.common.MetricsQueryEditorMode" json:"editor_mode,omitempty"`
	TimeFrame       *TimeFrameSelect        `protobuf:"bytes,4,opt,name=time_frame,json=timeFrame,proto3" json:"time_frame,omitempty"`
	Aggregation     Aggregation             `protobuf:"varint,5,opt,name=aggregation,proto3,enum=com.coralogixapis.dashboards.v1.ast.widgets.common.Aggregation" json:"aggregation,omitempty"`
	PromqlQueryType PromQLQueryType         `protobuf:"varint,6,opt,name=promql_query_type,json=promqlQueryType,proto3,enum=com.coralogixapis.dashboards.v1.common.PromQLQueryType" json:"promql_query_type,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Hexagon_MetricsQuery) Reset() {
	*x = Hexagon_MetricsQuery{}
	mi := &file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Hexagon_MetricsQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hexagon_MetricsQuery) ProtoMessage() {}

func (x *Hexagon_MetricsQuery) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hexagon_MetricsQuery.ProtoReflect.Descriptor instead.
func (*Hexagon_MetricsQuery) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Hexagon_MetricsQuery) GetPromqlQuery() *PromQlQuery {
	if x != nil {
		return x.PromqlQuery
	}
	return nil
}

func (x *Hexagon_MetricsQuery) GetFilters() []*Filter_MetricsFilter {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *Hexagon_MetricsQuery) GetEditorMode() MetricsQueryEditorMode {
	if x != nil {
		return x.EditorMode
	}
	return MetricsQueryEditorMode_METRICS_QUERY_EDITOR_MODE_UNSPECIFIED
}

func (x *Hexagon_MetricsQuery) GetTimeFrame() *TimeFrameSelect {
	if x != nil {
		return x.TimeFrame
	}
	return nil
}

func (x *Hexagon_MetricsQuery) GetAggregation() Aggregation {
	if x != nil {
		return x.Aggregation
	}
	return Aggregation_AGGREGATION_UNSPECIFIED
}

func (x *Hexagon_MetricsQuery) GetPromqlQueryType() PromQLQueryType {
	if x != nil {
		return x.PromqlQueryType
	}
	return PromQLQueryType_PROM_QL_QUERY_TYPE_UNSPECIFIED
}

type Hexagon_LogsQuery struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	LuceneQuery     *LuceneQuery           `protobuf:"bytes,1,opt,name=lucene_query,json=luceneQuery,proto3" json:"lucene_query,omitempty"`
	LogsAggregation *LogsAggregation       `protobuf:"bytes,2,opt,name=logs_aggregation,json=logsAggregation,proto3" json:"logs_aggregation,omitempty"`
	Filters         []*Filter_LogsFilter   `protobuf:"bytes,3,rep,name=filters,proto3" json:"filters,omitempty"`
	GroupBy         []*ObservationField    `protobuf:"bytes,4,rep,name=group_by,json=groupBy,proto3" json:"group_by,omitempty"`
	TimeFrame       *TimeFrameSelect       `protobuf:"bytes,5,opt,name=time_frame,json=timeFrame,proto3" json:"time_frame,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Hexagon_LogsQuery) Reset() {
	*x = Hexagon_LogsQuery{}
	mi := &file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Hexagon_LogsQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hexagon_LogsQuery) ProtoMessage() {}

func (x *Hexagon_LogsQuery) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hexagon_LogsQuery.ProtoReflect.Descriptor instead.
func (*Hexagon_LogsQuery) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Hexagon_LogsQuery) GetLuceneQuery() *LuceneQuery {
	if x != nil {
		return x.LuceneQuery
	}
	return nil
}

func (x *Hexagon_LogsQuery) GetLogsAggregation() *LogsAggregation {
	if x != nil {
		return x.LogsAggregation
	}
	return nil
}

func (x *Hexagon_LogsQuery) GetFilters() []*Filter_LogsFilter {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *Hexagon_LogsQuery) GetGroupBy() []*ObservationField {
	if x != nil {
		return x.GroupBy
	}
	return nil
}

func (x *Hexagon_LogsQuery) GetTimeFrame() *TimeFrameSelect {
	if x != nil {
		return x.TimeFrame
	}
	return nil
}

type Hexagon_SpansQuery struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	LuceneQuery      *LuceneQuery           `protobuf:"bytes,1,opt,name=lucene_query,json=luceneQuery,proto3" json:"lucene_query,omitempty"`
	SpansAggregation *SpansAggregation      `protobuf:"bytes,2,opt,name=spans_aggregation,json=spansAggregation,proto3" json:"spans_aggregation,omitempty"`
	Filters          []*Filter_SpansFilter  `protobuf:"bytes,3,rep,name=filters,proto3" json:"filters,omitempty"`
	GroupBy          []*SpanField           `protobuf:"bytes,4,rep,name=group_by,json=groupBy,proto3" json:"group_by,omitempty"`
	TimeFrame        *TimeFrameSelect       `protobuf:"bytes,5,opt,name=time_frame,json=timeFrame,proto3" json:"time_frame,omitempty"`
	GroupBys         []*ObservationField    `protobuf:"bytes,6,rep,name=group_bys,json=groupBys,proto3" json:"group_bys,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Hexagon_SpansQuery) Reset() {
	*x = Hexagon_SpansQuery{}
	mi := &file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Hexagon_SpansQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hexagon_SpansQuery) ProtoMessage() {}

func (x *Hexagon_SpansQuery) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hexagon_SpansQuery.ProtoReflect.Descriptor instead.
func (*Hexagon_SpansQuery) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Hexagon_SpansQuery) GetLuceneQuery() *LuceneQuery {
	if x != nil {
		return x.LuceneQuery
	}
	return nil
}

func (x *Hexagon_SpansQuery) GetSpansAggregation() *SpansAggregation {
	if x != nil {
		return x.SpansAggregation
	}
	return nil
}

func (x *Hexagon_SpansQuery) GetFilters() []*Filter_SpansFilter {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *Hexagon_SpansQuery) GetGroupBy() []*SpanField {
	if x != nil {
		return x.GroupBy
	}
	return nil
}

func (x *Hexagon_SpansQuery) GetTimeFrame() *TimeFrameSelect {
	if x != nil {
		return x.TimeFrame
	}
	return nil
}

func (x *Hexagon_SpansQuery) GetGroupBys() []*ObservationField {
	if x != nil {
		return x.GroupBys
	}
	return nil
}

type Hexagon_DataprimeQuery struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	DataprimeQuery *DataprimeQuery        `protobuf:"bytes,1,opt,name=dataprime_query,json=dataprimeQuery,proto3" json:"dataprime_query,omitempty"`
	Filters        []*Filter_Source       `protobuf:"bytes,2,rep,name=filters,proto3" json:"filters,omitempty"`
	TimeFrame      *TimeFrameSelect       `protobuf:"bytes,3,opt,name=time_frame,json=timeFrame,proto3" json:"time_frame,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Hexagon_DataprimeQuery) Reset() {
	*x = Hexagon_DataprimeQuery{}
	mi := &file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Hexagon_DataprimeQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hexagon_DataprimeQuery) ProtoMessage() {}

func (x *Hexagon_DataprimeQuery) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hexagon_DataprimeQuery.ProtoReflect.Descriptor instead.
func (*Hexagon_DataprimeQuery) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_rawDescGZIP(), []int{0, 4}
}

func (x *Hexagon_DataprimeQuery) GetDataprimeQuery() *DataprimeQuery {
	if x != nil {
		return x.DataprimeQuery
	}
	return nil
}

func (x *Hexagon_DataprimeQuery) GetFilters() []*Filter_Source {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *Hexagon_DataprimeQuery) GetTimeFrame() *TimeFrameSelect {
	if x != nil {
		return x.TimeFrame
	}
	return nil
}

var File_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto protoreflect.FileDescriptor

const file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_rawDesc = "" +
	"\n" +
	"9com/coralogixapis/dashboards/v1/ast/widgets/hexagon.proto\x12+com.coralogixapis.dashboards.v1.ast.widgets\x1a0com/coralogixapis/dashboards/v1/ast/filter.proto\x1aEcom/coralogixapis/dashboards/v1/ast/widgets/common/custom_links.proto\x1aGcom/coralogixapis/dashboards/v1/ast/widgets/common/data_mode_type.proto\x1a?com/coralogixapis/dashboards/v1/ast/widgets/common/legend.proto\x1aQcom/coralogixapis/dashboards/v1/ast/widgets/common/metrics_aggregation_type.proto\x1aRcom/coralogixapis/dashboards/v1/ast/widgets/common/metrics_query_editor_mode.proto\x1a@com/coralogixapis/dashboards/v1/ast/widgets/common/queries.proto\x1aCcom/coralogixapis/dashboards/v1/ast/widgets/common/thresholds.proto\x1a>com/coralogixapis/dashboards/v1/ast/widgets/common/units.proto\x1a=com/coralogixapis/dashboards/v1/common/logs_aggregation.proto\x1a>com/coralogixapis/dashboards/v1/common/observation_field.proto\x1a2com/coralogixapis/dashboards/v1/common/query.proto\x1a7com/coralogixapis/dashboards/v1/common/span_field.proto\x1a>com/coralogixapis/dashboards/v1/common/spans_aggregation.proto\x1a7com/coralogixapis/dashboards/v1/common/time_frame.proto\x1a\x1egoogle/protobuf/wrappers.proto\"\xb8\x19\n" +
	"\aHexagon\x12P\n" +
	"\x05query\x18\x01 \x01(\v2:.com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.QueryR\x05query\x12.\n" +
	"\x03min\x18\x02 \x01(\v2\x1c.google.protobuf.DoubleValueR\x03min\x12.\n" +
	"\x03max\x18\x03 \x01(\v2\x1c.google.protobuf.DoubleValueR\x03max\x12L\n" +
	"\x04unit\x18\x06 \x01(\x0e28.com.coralogixapis.dashboards.v1.ast.widgets.common.UnitR\x04unit\x12]\n" +
	"\n" +
	"thresholds\x18\a \x03(\v2=.com.coralogixapis.dashboards.v1.ast.widgets.common.ThresholdR\n" +
	"thresholds\x12f\n" +
	"\x0edata_mode_type\x18\b \x01(\x0e2@.com.coralogixapis.dashboards.v1.ast.widgets.common.DataModeTypeR\fdataModeType\x12=\n" +
	"\vcustom_unit\x18\n" +
	" \x01(\v2\x1c.google.protobuf.StringValueR\n" +
	"customUnit\x125\n" +
	"\adecimal\x18\v \x01(\v2\x1b.google.protobuf.Int32ValueR\adecimal\x12h\n" +
	"\x0ethreshold_type\x18\f \x01(\x0e2A.com.coralogixapis.dashboards.v1.ast.widgets.common.ThresholdTypeR\rthresholdType\x12R\n" +
	"\x06legend\x18\r \x01(\v2:.com.coralogixapis.dashboards.v1.ast.widgets.common.LegendR\x06legend\x12Y\n" +
	"\tlegend_by\x18\x0e \x01(\x0e2<.com.coralogixapis.dashboards.v1.ast.widgets.common.LegendByR\blegendBy\x12a\n" +
	"\fcustom_links\x18\x0f \x03(\v2>.com.coralogixapis.dashboards.v1.ast.widgets.common.CustomLinkR\vcustomLinks\x1a\x83\x03\n" +
	"\x05Query\x12]\n" +
	"\ametrics\x18\x01 \x01(\v2A.com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.MetricsQueryH\x00R\ametrics\x12T\n" +
	"\x04logs\x18\x02 \x01(\v2>.com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.LogsQueryH\x00R\x04logs\x12W\n" +
	"\x05spans\x18\x03 \x01(\v2?.com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.SpansQueryH\x00R\x05spans\x12c\n" +
	"\tdataprime\x18\x04 \x01(\v2C.com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.DataprimeQueryH\x00R\tdataprimeB\a\n" +
	"\x05value\x1a\xd4\x04\n" +
	"\fMetricsQuery\x12b\n" +
	"\fpromql_query\x18\x01 \x01(\v2?.com.coralogixapis.dashboards.v1.ast.widgets.common.PromQlQueryR\vpromqlQuery\x12S\n" +
	"\afilters\x18\x02 \x03(\v29.com.coralogixapis.dashboards.v1.ast.Filter.MetricsFilterR\afilters\x12k\n" +
	"\veditor_mode\x18\x03 \x01(\x0e2J.com.coralogixapis.dashboards.v1.ast.widgets.common.MetricsQueryEditorModeR\n" +
	"editorMode\x12V\n" +
	"\n" +
	"time_frame\x18\x04 \x01(\v27.com.coralogixapis.dashboards.v1.common.TimeFrameSelectR\ttimeFrame\x12a\n" +
	"\vaggregation\x18\x05 \x01(\x0e2?.com.coralogixapis.dashboards.v1.ast.widgets.common.AggregationR\vaggregation\x12c\n" +
	"\x11promql_query_type\x18\x06 \x01(\x0e27.com.coralogixapis.dashboards.v1.common.PromQLQueryTypeR\x0fpromqlQueryType\x1a\xd2\x03\n" +
	"\tLogsQuery\x12b\n" +
	"\flucene_query\x18\x01 \x01(\v2?.com.coralogixapis.dashboards.v1.ast.widgets.common.LuceneQueryR\vluceneQuery\x12b\n" +
	"\x10logs_aggregation\x18\x02 \x01(\v27.com.coralogixapis.dashboards.v1.common.LogsAggregationR\x0flogsAggregation\x12P\n" +
	"\afilters\x18\x03 \x03(\v26.com.coralogixapis.dashboards.v1.ast.Filter.LogsFilterR\afilters\x12S\n" +
	"\bgroup_by\x18\x04 \x03(\v28.com.coralogixapis.dashboards.v1.common.ObservationFieldR\agroupBy\x12V\n" +
	"\n" +
	"time_frame\x18\x05 \x01(\v27.com.coralogixapis.dashboards.v1.common.TimeFrameSelectR\ttimeFrame\x1a\xa7\x04\n" +
	"\n" +
	"SpansQuery\x12b\n" +
	"\flucene_query\x18\x01 \x01(\v2?.com.coralogixapis.dashboards.v1.ast.widgets.common.LuceneQueryR\vluceneQuery\x12e\n" +
	"\x11spans_aggregation\x18\x02 \x01(\v28.com.coralogixapis.dashboards.v1.common.SpansAggregationR\x10spansAggregation\x12Q\n" +
	"\afilters\x18\x03 \x03(\v27.com.coralogixapis.dashboards.v1.ast.Filter.SpansFilterR\afilters\x12L\n" +
	"\bgroup_by\x18\x04 \x03(\v21.com.coralogixapis.dashboards.v1.common.SpanFieldR\agroupBy\x12V\n" +
	"\n" +
	"time_frame\x18\x05 \x01(\v27.com.coralogixapis.dashboards.v1.common.TimeFrameSelectR\ttimeFrame\x12U\n" +
	"\tgroup_bys\x18\x06 \x03(\v28.com.coralogixapis.dashboards.v1.common.ObservationFieldR\bgroupBys\x1a\x97\x02\n" +
	"\x0eDataprimeQuery\x12_\n" +
	"\x0fdataprime_query\x18\x01 \x01(\v26.com.coralogixapis.dashboards.v1.common.DataprimeQueryR\x0edataprimeQuery\x12L\n" +
	"\afilters\x18\x02 \x03(\v22.com.coralogixapis.dashboards.v1.ast.Filter.SourceR\afilters\x12V\n" +
	"\n" +
	"time_frame\x18\x03 \x01(\v27.com.coralogixapis.dashboards.v1.common.TimeFrameSelectR\ttimeFrameb\x06proto3"

var (
	file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_rawDescOnce sync.Once
	file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_rawDescData []byte
)

func file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_rawDesc), len(file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_rawDesc)))
	})
	return file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_rawDescData
}

var file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_goTypes = []any{
	(*Hexagon)(nil),                // 0: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon
	(*Hexagon_Query)(nil),          // 1: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.Query
	(*Hexagon_MetricsQuery)(nil),   // 2: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.MetricsQuery
	(*Hexagon_LogsQuery)(nil),      // 3: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.LogsQuery
	(*Hexagon_SpansQuery)(nil),     // 4: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.SpansQuery
	(*Hexagon_DataprimeQuery)(nil), // 5: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.DataprimeQuery
	(*wrapperspb.DoubleValue)(nil), // 6: google.protobuf.DoubleValue
	(Unit)(0),                      // 7: com.coralogixapis.dashboards.v1.ast.widgets.common.Unit
	(*Threshold)(nil),              // 8: com.coralogixapis.dashboards.v1.ast.widgets.common.Threshold
	(DataModeType)(0),              // 9: com.coralogixapis.dashboards.v1.ast.widgets.common.DataModeType
	(*wrapperspb.StringValue)(nil), // 10: google.protobuf.StringValue
	(*wrapperspb.Int32Value)(nil),  // 11: google.protobuf.Int32Value
	(ThresholdType)(0),             // 12: com.coralogixapis.dashboards.v1.ast.widgets.common.ThresholdType
	(*Legend)(nil),                 // 13: com.coralogixapis.dashboards.v1.ast.widgets.common.Legend
	(LegendBy)(0),                  // 14: com.coralogixapis.dashboards.v1.ast.widgets.common.LegendBy
	(*CustomLink)(nil),             // 15: com.coralogixapis.dashboards.v1.ast.widgets.common.CustomLink
	(*PromQlQuery)(nil),            // 16: com.coralogixapis.dashboards.v1.ast.widgets.common.PromQlQuery
	(*Filter_MetricsFilter)(nil),   // 17: com.coralogixapis.dashboards.v1.ast.Filter.MetricsFilter
	(MetricsQueryEditorMode)(0),    // 18: com.coralogixapis.dashboards.v1.ast.widgets.common.MetricsQueryEditorMode
	(*TimeFrameSelect)(nil),        // 19: com.coralogixapis.dashboards.v1.common.TimeFrameSelect
	(Aggregation)(0),               // 20: com.coralogixapis.dashboards.v1.ast.widgets.common.Aggregation
	(PromQLQueryType)(0),           // 21: com.coralogixapis.dashboards.v1.common.PromQLQueryType
	(*LuceneQuery)(nil),            // 22: com.coralogixapis.dashboards.v1.ast.widgets.common.LuceneQuery
	(*LogsAggregation)(nil),        // 23: com.coralogixapis.dashboards.v1.common.LogsAggregation
	(*Filter_LogsFilter)(nil),      // 24: com.coralogixapis.dashboards.v1.ast.Filter.LogsFilter
	(*ObservationField)(nil),       // 25: com.coralogixapis.dashboards.v1.common.ObservationField
	(*SpansAggregation)(nil),       // 26: com.coralogixapis.dashboards.v1.common.SpansAggregation
	(*Filter_SpansFilter)(nil),     // 27: com.coralogixapis.dashboards.v1.ast.Filter.SpansFilter
	(*SpanField)(nil),              // 28: com.coralogixapis.dashboards.v1.common.SpanField
	(*DataprimeQuery)(nil),         // 29: com.coralogixapis.dashboards.v1.common.DataprimeQuery
	(*Filter_Source)(nil),          // 30: com.coralogixapis.dashboards.v1.ast.Filter.Source
}
var file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_depIdxs = []int32{
	1,  // 0: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.query:type_name -> com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.Query
	6,  // 1: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.min:type_name -> google.protobuf.DoubleValue
	6,  // 2: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.max:type_name -> google.protobuf.DoubleValue
	7,  // 3: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.unit:type_name -> com.coralogixapis.dashboards.v1.ast.widgets.common.Unit
	8,  // 4: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.thresholds:type_name -> com.coralogixapis.dashboards.v1.ast.widgets.common.Threshold
	9,  // 5: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.data_mode_type:type_name -> com.coralogixapis.dashboards.v1.ast.widgets.common.DataModeType
	10, // 6: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.custom_unit:type_name -> google.protobuf.StringValue
	11, // 7: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.decimal:type_name -> google.protobuf.Int32Value
	12, // 8: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.threshold_type:type_name -> com.coralogixapis.dashboards.v1.ast.widgets.common.ThresholdType
	13, // 9: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.legend:type_name -> com.coralogixapis.dashboards.v1.ast.widgets.common.Legend
	14, // 10: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.legend_by:type_name -> com.coralogixapis.dashboards.v1.ast.widgets.common.LegendBy
	15, // 11: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.custom_links:type_name -> com.coralogixapis.dashboards.v1.ast.widgets.common.CustomLink
	2,  // 12: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.Query.metrics:type_name -> com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.MetricsQuery
	3,  // 13: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.Query.logs:type_name -> com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.LogsQuery
	4,  // 14: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.Query.spans:type_name -> com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.SpansQuery
	5,  // 15: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.Query.dataprime:type_name -> com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.DataprimeQuery
	16, // 16: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.MetricsQuery.promql_query:type_name -> com.coralogixapis.dashboards.v1.ast.widgets.common.PromQlQuery
	17, // 17: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.MetricsQuery.filters:type_name -> com.coralogixapis.dashboards.v1.ast.Filter.MetricsFilter
	18, // 18: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.MetricsQuery.editor_mode:type_name -> com.coralogixapis.dashboards.v1.ast.widgets.common.MetricsQueryEditorMode
	19, // 19: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.MetricsQuery.time_frame:type_name -> com.coralogixapis.dashboards.v1.common.TimeFrameSelect
	20, // 20: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.MetricsQuery.aggregation:type_name -> com.coralogixapis.dashboards.v1.ast.widgets.common.Aggregation
	21, // 21: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.MetricsQuery.promql_query_type:type_name -> com.coralogixapis.dashboards.v1.common.PromQLQueryType
	22, // 22: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.LogsQuery.lucene_query:type_name -> com.coralogixapis.dashboards.v1.ast.widgets.common.LuceneQuery
	23, // 23: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.LogsQuery.logs_aggregation:type_name -> com.coralogixapis.dashboards.v1.common.LogsAggregation
	24, // 24: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.LogsQuery.filters:type_name -> com.coralogixapis.dashboards.v1.ast.Filter.LogsFilter
	25, // 25: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.LogsQuery.group_by:type_name -> com.coralogixapis.dashboards.v1.common.ObservationField
	19, // 26: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.LogsQuery.time_frame:type_name -> com.coralogixapis.dashboards.v1.common.TimeFrameSelect
	22, // 27: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.SpansQuery.lucene_query:type_name -> com.coralogixapis.dashboards.v1.ast.widgets.common.LuceneQuery
	26, // 28: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.SpansQuery.spans_aggregation:type_name -> com.coralogixapis.dashboards.v1.common.SpansAggregation
	27, // 29: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.SpansQuery.filters:type_name -> com.coralogixapis.dashboards.v1.ast.Filter.SpansFilter
	28, // 30: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.SpansQuery.group_by:type_name -> com.coralogixapis.dashboards.v1.common.SpanField
	19, // 31: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.SpansQuery.time_frame:type_name -> com.coralogixapis.dashboards.v1.common.TimeFrameSelect
	25, // 32: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.SpansQuery.group_bys:type_name -> com.coralogixapis.dashboards.v1.common.ObservationField
	29, // 33: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.DataprimeQuery.dataprime_query:type_name -> com.coralogixapis.dashboards.v1.common.DataprimeQuery
	30, // 34: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.DataprimeQuery.filters:type_name -> com.coralogixapis.dashboards.v1.ast.Filter.Source
	19, // 35: com.coralogixapis.dashboards.v1.ast.widgets.Hexagon.DataprimeQuery.time_frame:type_name -> com.coralogixapis.dashboards.v1.common.TimeFrameSelect
	36, // [36:36] is the sub-list for method output_type
	36, // [36:36] is the sub-list for method input_type
	36, // [36:36] is the sub-list for extension type_name
	36, // [36:36] is the sub-list for extension extendee
	0,  // [0:36] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_init() }
func file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_init() {
	if File_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto != nil {
		return
	}
	file_com_coralogixapis_dashboards_v1_ast_filter_proto_init()
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_custom_links_proto_init()
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_data_mode_type_proto_init()
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_init()
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_aggregation_type_proto_init()
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_query_editor_mode_proto_init()
	//file_com_coralogixapis_dashboards_v1_ast_widgets_common_queries_proto_init()
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_thresholds_proto_init()
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_units_proto_init()
	file_com_coralogixapis_dashboards_v1_common_logs_aggregation_proto_init()
	file_com_coralogixapis_dashboards_v1_common_observation_field_proto_init()
	file_com_coralogixapis_dashboards_v1_common_query_proto_init()
	file_com_coralogixapis_dashboards_v1_common_span_field_proto_init()
	file_com_coralogixapis_dashboards_v1_common_spans_aggregation_proto_init()
	file_com_coralogixapis_dashboards_v1_common_time_frame_proto_init()
	file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_msgTypes[1].OneofWrappers = []any{
		(*Hexagon_Query_Metrics)(nil),
		(*Hexagon_Query_Logs)(nil),
		(*Hexagon_Query_Spans)(nil),
		(*Hexagon_Query_Dataprime)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_rawDesc), len(file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto = out.File
	file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_goTypes = nil
	file_com_coralogixapis_dashboards_v1_ast_widgets_hexagon_proto_depIdxs = nil
}
