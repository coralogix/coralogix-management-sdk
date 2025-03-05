// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: com/coralogixapis/dashboards/v1/ast/widgets/common/legend.proto

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

type LegendBy int32

const (
	LegendBy_LEGEND_BY_UNSPECIFIED LegendBy = 0
	LegendBy_LEGEND_BY_THRESHOLDS  LegendBy = 1
	LegendBy_LEGEND_BY_GROUPS      LegendBy = 2
)

// Enum value maps for LegendBy.
var (
	LegendBy_name = map[int32]string{
		0: "LEGEND_BY_UNSPECIFIED",
		1: "LEGEND_BY_THRESHOLDS",
		2: "LEGEND_BY_GROUPS",
	}
	LegendBy_value = map[string]int32{
		"LEGEND_BY_UNSPECIFIED": 0,
		"LEGEND_BY_THRESHOLDS":  1,
		"LEGEND_BY_GROUPS":      2,
	}
)

func (x LegendBy) Enum() *LegendBy {
	p := new(LegendBy)
	*p = x
	return p
}

func (x LegendBy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LegendBy) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_enumTypes[0].Descriptor()
}

func (LegendBy) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_enumTypes[0]
}

func (x LegendBy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LegendBy.Descriptor instead.
func (LegendBy) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_rawDescGZIP(), []int{0}
}

type Legend_LegendColumn int32

const (
	Legend_LEGEND_COLUMN_UNSPECIFIED Legend_LegendColumn = 0
	Legend_LEGEND_COLUMN_MIN         Legend_LegendColumn = 1
	Legend_LEGEND_COLUMN_MAX         Legend_LegendColumn = 2
	Legend_LEGEND_COLUMN_SUM         Legend_LegendColumn = 3
	Legend_LEGEND_COLUMN_AVG         Legend_LegendColumn = 4
	Legend_LEGEND_COLUMN_LAST        Legend_LegendColumn = 5
	Legend_LEGEND_COLUMN_NAME        Legend_LegendColumn = 6
)

// Enum value maps for Legend_LegendColumn.
var (
	Legend_LegendColumn_name = map[int32]string{
		0: "LEGEND_COLUMN_UNSPECIFIED",
		1: "LEGEND_COLUMN_MIN",
		2: "LEGEND_COLUMN_MAX",
		3: "LEGEND_COLUMN_SUM",
		4: "LEGEND_COLUMN_AVG",
		5: "LEGEND_COLUMN_LAST",
		6: "LEGEND_COLUMN_NAME",
	}
	Legend_LegendColumn_value = map[string]int32{
		"LEGEND_COLUMN_UNSPECIFIED": 0,
		"LEGEND_COLUMN_MIN":         1,
		"LEGEND_COLUMN_MAX":         2,
		"LEGEND_COLUMN_SUM":         3,
		"LEGEND_COLUMN_AVG":         4,
		"LEGEND_COLUMN_LAST":        5,
		"LEGEND_COLUMN_NAME":        6,
	}
)

func (x Legend_LegendColumn) Enum() *Legend_LegendColumn {
	p := new(Legend_LegendColumn)
	*p = x
	return p
}

func (x Legend_LegendColumn) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Legend_LegendColumn) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_enumTypes[1].Descriptor()
}

func (Legend_LegendColumn) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_enumTypes[1]
}

func (x Legend_LegendColumn) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Legend_LegendColumn.Descriptor instead.
func (Legend_LegendColumn) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_rawDescGZIP(), []int{0, 0}
}

type Legend_LegendPlacement int32

const (
	Legend_LEGEND_PLACEMENT_UNSPECIFIED Legend_LegendPlacement = 0
	Legend_LEGEND_PLACEMENT_AUTO        Legend_LegendPlacement = 1
	Legend_LEGEND_PLACEMENT_BOTTOM      Legend_LegendPlacement = 2
	Legend_LEGEND_PLACEMENT_SIDE        Legend_LegendPlacement = 3
	Legend_LEGEND_PLACEMENT_HIDDEN      Legend_LegendPlacement = 4
)

// Enum value maps for Legend_LegendPlacement.
var (
	Legend_LegendPlacement_name = map[int32]string{
		0: "LEGEND_PLACEMENT_UNSPECIFIED",
		1: "LEGEND_PLACEMENT_AUTO",
		2: "LEGEND_PLACEMENT_BOTTOM",
		3: "LEGEND_PLACEMENT_SIDE",
		4: "LEGEND_PLACEMENT_HIDDEN",
	}
	Legend_LegendPlacement_value = map[string]int32{
		"LEGEND_PLACEMENT_UNSPECIFIED": 0,
		"LEGEND_PLACEMENT_AUTO":        1,
		"LEGEND_PLACEMENT_BOTTOM":      2,
		"LEGEND_PLACEMENT_SIDE":        3,
		"LEGEND_PLACEMENT_HIDDEN":      4,
	}
)

func (x Legend_LegendPlacement) Enum() *Legend_LegendPlacement {
	p := new(Legend_LegendPlacement)
	*p = x
	return p
}

func (x Legend_LegendPlacement) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Legend_LegendPlacement) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_enumTypes[2].Descriptor()
}

func (Legend_LegendPlacement) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_enumTypes[2]
}

func (x Legend_LegendPlacement) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Legend_LegendPlacement.Descriptor instead.
func (Legend_LegendPlacement) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_rawDescGZIP(), []int{0, 1}
}

type Legend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsVisible    *wrapperspb.BoolValue  `protobuf:"bytes,1,opt,name=is_visible,json=isVisible,proto3" json:"is_visible,omitempty"`
	Columns      []Legend_LegendColumn  `protobuf:"varint,2,rep,packed,name=columns,proto3,enum=com.coralogixapis.dashboards.v1.ast.widgets.common.Legend_LegendColumn" json:"columns,omitempty"`
	GroupByQuery *wrapperspb.BoolValue  `protobuf:"bytes,3,opt,name=group_by_query,json=groupByQuery,proto3" json:"group_by_query,omitempty"`
	Placement    Legend_LegendPlacement `protobuf:"varint,4,opt,name=placement,proto3,enum=com.coralogixapis.dashboards.v1.ast.widgets.common.Legend_LegendPlacement" json:"placement,omitempty"`
}

func (x *Legend) Reset() {
	*x = Legend{}
	mi := &file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Legend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Legend) ProtoMessage() {}

func (x *Legend) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Legend.ProtoReflect.Descriptor instead.
func (*Legend) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_rawDescGZIP(), []int{0}
}

func (x *Legend) GetIsVisible() *wrapperspb.BoolValue {
	if x != nil {
		return x.IsVisible
	}
	return nil
}

func (x *Legend) GetColumns() []Legend_LegendColumn {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *Legend) GetGroupByQuery() *wrapperspb.BoolValue {
	if x != nil {
		return x.GroupByQuery
	}
	return nil
}

func (x *Legend) GetPlacement() Legend_LegendPlacement {
	if x != nil {
		return x.Placement
	}
	return Legend_LEGEND_PLACEMENT_UNSPECIFIED
}

var File_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto protoreflect.FileDescriptor

var file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x73, 0x74, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6c, 0x65, 0x67, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x32, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x05, 0x0a, 0x06, 0x4c, 0x65, 0x67, 0x65, 0x6e, 0x64,
	0x12, 0x39, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x09, 0x69, 0x73, 0x56, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x61, 0x0a, 0x07, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x47, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61,
	0x73, 0x74, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x4c, 0x65, 0x67, 0x65, 0x6e, 0x64, 0x2e, 0x4c, 0x65, 0x67, 0x65, 0x6e, 0x64, 0x43,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x12, 0x40,
	0x0a, 0x0e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x62, 0x79, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x68, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x4a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x65, 0x67, 0x65, 0x6e, 0x64, 0x2e,
	0x4c, 0x65, 0x67, 0x65, 0x6e, 0x64, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x09, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xb9, 0x01, 0x0a, 0x0c, 0x4c,
	0x65, 0x67, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x1d, 0x0a, 0x19, 0x4c,
	0x45, 0x47, 0x45, 0x4e, 0x44, 0x5f, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x45,
	0x47, 0x45, 0x4e, 0x44, 0x5f, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e, 0x5f, 0x4d, 0x49, 0x4e, 0x10,
	0x01, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x45, 0x47, 0x45, 0x4e, 0x44, 0x5f, 0x43, 0x4f, 0x4c, 0x55,
	0x4d, 0x4e, 0x5f, 0x4d, 0x41, 0x58, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x45, 0x47, 0x45,
	0x4e, 0x44, 0x5f, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e, 0x5f, 0x53, 0x55, 0x4d, 0x10, 0x03, 0x12,
	0x15, 0x0a, 0x11, 0x4c, 0x45, 0x47, 0x45, 0x4e, 0x44, 0x5f, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e,
	0x5f, 0x41, 0x56, 0x47, 0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x4c, 0x45, 0x47, 0x45, 0x4e, 0x44,
	0x5f, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e, 0x5f, 0x4c, 0x41, 0x53, 0x54, 0x10, 0x05, 0x12, 0x16,
	0x0a, 0x12, 0x4c, 0x45, 0x47, 0x45, 0x4e, 0x44, 0x5f, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e, 0x5f,
	0x4e, 0x41, 0x4d, 0x45, 0x10, 0x06, 0x22, 0xa3, 0x01, 0x0a, 0x0f, 0x4c, 0x65, 0x67, 0x65, 0x6e,
	0x64, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x1c, 0x4c, 0x45,
	0x47, 0x45, 0x4e, 0x44, 0x5f, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15,
	0x4c, 0x45, 0x47, 0x45, 0x4e, 0x44, 0x5f, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x41, 0x55, 0x54, 0x4f, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x4c, 0x45, 0x47, 0x45, 0x4e,
	0x44, 0x5f, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x42, 0x4f, 0x54, 0x54,
	0x4f, 0x4d, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x4c, 0x45, 0x47, 0x45, 0x4e, 0x44, 0x5f, 0x50,
	0x4c, 0x41, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x49, 0x44, 0x45, 0x10, 0x03, 0x12,
	0x1b, 0x0a, 0x17, 0x4c, 0x45, 0x47, 0x45, 0x4e, 0x44, 0x5f, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x48, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x10, 0x04, 0x2a, 0x55, 0x0a, 0x08,
	0x4c, 0x65, 0x67, 0x65, 0x6e, 0x64, 0x42, 0x79, 0x12, 0x19, 0x0a, 0x15, 0x4c, 0x45, 0x47, 0x45,
	0x4e, 0x44, 0x5f, 0x42, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x4c, 0x45, 0x47, 0x45, 0x4e, 0x44, 0x5f, 0x42, 0x59,
	0x5f, 0x54, 0x48, 0x52, 0x45, 0x53, 0x48, 0x4f, 0x4c, 0x44, 0x53, 0x10, 0x01, 0x12, 0x14, 0x0a,
	0x10, 0x4c, 0x45, 0x47, 0x45, 0x4e, 0x44, 0x5f, 0x42, 0x59, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50,
	0x53, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_rawDescOnce sync.Once
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_rawDescData = file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_rawDesc
)

func file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_rawDescData)
	})
	return file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_rawDescData
}

var file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_goTypes = []any{
	(LegendBy)(0),                // 0: com.coralogixapis.dashboards.v1.ast.widgets.common.LegendBy
	(Legend_LegendColumn)(0),     // 1: com.coralogixapis.dashboards.v1.ast.widgets.common.Legend.LegendColumn
	(Legend_LegendPlacement)(0),  // 2: com.coralogixapis.dashboards.v1.ast.widgets.common.Legend.LegendPlacement
	(*Legend)(nil),               // 3: com.coralogixapis.dashboards.v1.ast.widgets.common.Legend
	(*wrapperspb.BoolValue)(nil), // 4: google.protobuf.BoolValue
}
var file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_depIdxs = []int32{
	4, // 0: com.coralogixapis.dashboards.v1.ast.widgets.common.Legend.is_visible:type_name -> google.protobuf.BoolValue
	1, // 1: com.coralogixapis.dashboards.v1.ast.widgets.common.Legend.columns:type_name -> com.coralogixapis.dashboards.v1.ast.widgets.common.Legend.LegendColumn
	4, // 2: com.coralogixapis.dashboards.v1.ast.widgets.common.Legend.group_by_query:type_name -> google.protobuf.BoolValue
	2, // 3: com.coralogixapis.dashboards.v1.ast.widgets.common.Legend.placement:type_name -> com.coralogixapis.dashboards.v1.ast.widgets.common.Legend.LegendPlacement
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_init() }
func file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_init() {
	if File_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto = out.File
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_rawDesc = nil
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_goTypes = nil
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_legend_proto_depIdxs = nil
}
