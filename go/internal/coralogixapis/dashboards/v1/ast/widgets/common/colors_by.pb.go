// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/dashboards/v1/ast/widgets/common/colors_by.proto

package common

import (
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

type ColorsBy struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Value:
	//
	//	*ColorsBy_Stack
	//	*ColorsBy_GroupBy
	//	*ColorsBy_Aggregation
	Value         isColorsBy_Value `protobuf_oneof:"value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ColorsBy) Reset() {
	*x = ColorsBy{}
	mi := &file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ColorsBy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColorsBy) ProtoMessage() {}

func (x *ColorsBy) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColorsBy.ProtoReflect.Descriptor instead.
func (*ColorsBy) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_rawDescGZIP(), []int{0}
}

func (x *ColorsBy) GetValue() isColorsBy_Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *ColorsBy) GetStack() *ColorsBy_ColorsByStack {
	if x != nil {
		if x, ok := x.Value.(*ColorsBy_Stack); ok {
			return x.Stack
		}
	}
	return nil
}

func (x *ColorsBy) GetGroupBy() *ColorsBy_ColorsByGroupBy {
	if x != nil {
		if x, ok := x.Value.(*ColorsBy_GroupBy); ok {
			return x.GroupBy
		}
	}
	return nil
}

func (x *ColorsBy) GetAggregation() *ColorsBy_ColorsByAggregation {
	if x != nil {
		if x, ok := x.Value.(*ColorsBy_Aggregation); ok {
			return x.Aggregation
		}
	}
	return nil
}

type isColorsBy_Value interface {
	isColorsBy_Value()
}

type ColorsBy_Stack struct {
	Stack *ColorsBy_ColorsByStack `protobuf:"bytes,1,opt,name=stack,proto3,oneof"`
}

type ColorsBy_GroupBy struct {
	GroupBy *ColorsBy_ColorsByGroupBy `protobuf:"bytes,2,opt,name=group_by,json=groupBy,proto3,oneof"`
}

type ColorsBy_Aggregation struct {
	Aggregation *ColorsBy_ColorsByAggregation `protobuf:"bytes,3,opt,name=aggregation,proto3,oneof"`
}

func (*ColorsBy_Stack) isColorsBy_Value() {}

func (*ColorsBy_GroupBy) isColorsBy_Value() {}

func (*ColorsBy_Aggregation) isColorsBy_Value() {}

type ColorsBy_ColorsByStack struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ColorsBy_ColorsByStack) Reset() {
	*x = ColorsBy_ColorsByStack{}
	mi := &file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ColorsBy_ColorsByStack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColorsBy_ColorsByStack) ProtoMessage() {}

func (x *ColorsBy_ColorsByStack) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColorsBy_ColorsByStack.ProtoReflect.Descriptor instead.
func (*ColorsBy_ColorsByStack) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_rawDescGZIP(), []int{0, 0}
}

type ColorsBy_ColorsByGroupBy struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ColorsBy_ColorsByGroupBy) Reset() {
	*x = ColorsBy_ColorsByGroupBy{}
	mi := &file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ColorsBy_ColorsByGroupBy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColorsBy_ColorsByGroupBy) ProtoMessage() {}

func (x *ColorsBy_ColorsByGroupBy) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColorsBy_ColorsByGroupBy.ProtoReflect.Descriptor instead.
func (*ColorsBy_ColorsByGroupBy) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_rawDescGZIP(), []int{0, 1}
}

type ColorsBy_ColorsByAggregation struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ColorsBy_ColorsByAggregation) Reset() {
	*x = ColorsBy_ColorsByAggregation{}
	mi := &file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ColorsBy_ColorsByAggregation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColorsBy_ColorsByAggregation) ProtoMessage() {}

func (x *ColorsBy_ColorsByAggregation) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColorsBy_ColorsByAggregation.ProtoReflect.Descriptor instead.
func (*ColorsBy_ColorsByAggregation) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_rawDescGZIP(), []int{0, 2}
}

var File_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto protoreflect.FileDescriptor

const file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_rawDesc = "" +
	"\n" +
	"Bcom/coralogixapis/dashboards/v1/ast/widgets/common/colors_by.proto\x122com.coralogixapis.dashboards.v1.ast.widgets.common\"\x93\x03\n" +
	"\bColorsBy\x12b\n" +
	"\x05stack\x18\x01 \x01(\v2J.com.coralogixapis.dashboards.v1.ast.widgets.common.ColorsBy.ColorsByStackH\x00R\x05stack\x12i\n" +
	"\bgroup_by\x18\x02 \x01(\v2L.com.coralogixapis.dashboards.v1.ast.widgets.common.ColorsBy.ColorsByGroupByH\x00R\agroupBy\x12t\n" +
	"\vaggregation\x18\x03 \x01(\v2P.com.coralogixapis.dashboards.v1.ast.widgets.common.ColorsBy.ColorsByAggregationH\x00R\vaggregation\x1a\x0f\n" +
	"\rColorsByStack\x1a\x11\n" +
	"\x0fColorsByGroupBy\x1a\x15\n" +
	"\x13ColorsByAggregationB\a\n" +
	"\x05valueb\x06proto3"

var (
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_rawDescOnce sync.Once
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_rawDescData []byte
)

func file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_rawDesc), len(file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_rawDesc)))
	})
	return file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_rawDescData
}

var file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_goTypes = []any{
	(*ColorsBy)(nil),                     // 0: com.coralogixapis.dashboards.v1.ast.widgets.common.ColorsBy
	(*ColorsBy_ColorsByStack)(nil),       // 1: com.coralogixapis.dashboards.v1.ast.widgets.common.ColorsBy.ColorsByStack
	(*ColorsBy_ColorsByGroupBy)(nil),     // 2: com.coralogixapis.dashboards.v1.ast.widgets.common.ColorsBy.ColorsByGroupBy
	(*ColorsBy_ColorsByAggregation)(nil), // 3: com.coralogixapis.dashboards.v1.ast.widgets.common.ColorsBy.ColorsByAggregation
}
var file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.dashboards.v1.ast.widgets.common.ColorsBy.stack:type_name -> com.coralogixapis.dashboards.v1.ast.widgets.common.ColorsBy.ColorsByStack
	2, // 1: com.coralogixapis.dashboards.v1.ast.widgets.common.ColorsBy.group_by:type_name -> com.coralogixapis.dashboards.v1.ast.widgets.common.ColorsBy.ColorsByGroupBy
	3, // 2: com.coralogixapis.dashboards.v1.ast.widgets.common.ColorsBy.aggregation:type_name -> com.coralogixapis.dashboards.v1.ast.widgets.common.ColorsBy.ColorsByAggregation
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_init() }
func file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_init() {
	if File_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto != nil {
		return
	}
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_msgTypes[0].OneofWrappers = []any{
		(*ColorsBy_Stack)(nil),
		(*ColorsBy_GroupBy)(nil),
		(*ColorsBy_Aggregation)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_rawDesc), len(file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto = out.File
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_goTypes = nil
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_colors_by_proto_depIdxs = nil
}
