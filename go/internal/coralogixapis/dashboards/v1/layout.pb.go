// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.3
// source: com/coralogixapis/dashboards/v1/ast/layout.proto

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

type SectionPredefinedColor int32

const (
	SectionPredefinedColor_SECTION_PREDEFINED_COLOR_UNSPECIFIED SectionPredefinedColor = 0
	SectionPredefinedColor_SECTION_PREDEFINED_COLOR_CYAN        SectionPredefinedColor = 1
	SectionPredefinedColor_SECTION_PREDEFINED_COLOR_GREEN       SectionPredefinedColor = 2
	SectionPredefinedColor_SECTION_PREDEFINED_COLOR_BLUE        SectionPredefinedColor = 3
	SectionPredefinedColor_SECTION_PREDEFINED_COLOR_PURPLE      SectionPredefinedColor = 4
	SectionPredefinedColor_SECTION_PREDEFINED_COLOR_MAGENTA     SectionPredefinedColor = 5
	SectionPredefinedColor_SECTION_PREDEFINED_COLOR_PINK        SectionPredefinedColor = 6
	SectionPredefinedColor_SECTION_PREDEFINED_COLOR_ORANGE      SectionPredefinedColor = 7
)

// Enum value maps for SectionPredefinedColor.
var (
	SectionPredefinedColor_name = map[int32]string{
		0: "SECTION_PREDEFINED_COLOR_UNSPECIFIED",
		1: "SECTION_PREDEFINED_COLOR_CYAN",
		2: "SECTION_PREDEFINED_COLOR_GREEN",
		3: "SECTION_PREDEFINED_COLOR_BLUE",
		4: "SECTION_PREDEFINED_COLOR_PURPLE",
		5: "SECTION_PREDEFINED_COLOR_MAGENTA",
		6: "SECTION_PREDEFINED_COLOR_PINK",
		7: "SECTION_PREDEFINED_COLOR_ORANGE",
	}
	SectionPredefinedColor_value = map[string]int32{
		"SECTION_PREDEFINED_COLOR_UNSPECIFIED": 0,
		"SECTION_PREDEFINED_COLOR_CYAN":        1,
		"SECTION_PREDEFINED_COLOR_GREEN":       2,
		"SECTION_PREDEFINED_COLOR_BLUE":        3,
		"SECTION_PREDEFINED_COLOR_PURPLE":      4,
		"SECTION_PREDEFINED_COLOR_MAGENTA":     5,
		"SECTION_PREDEFINED_COLOR_PINK":        6,
		"SECTION_PREDEFINED_COLOR_ORANGE":      7,
	}
)

func (x SectionPredefinedColor) Enum() *SectionPredefinedColor {
	p := new(SectionPredefinedColor)
	*p = x
	return p
}

func (x SectionPredefinedColor) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SectionPredefinedColor) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_dashboards_v1_ast_layout_proto_enumTypes[0].Descriptor()
}

func (SectionPredefinedColor) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_dashboards_v1_ast_layout_proto_enumTypes[0]
}

func (x SectionPredefinedColor) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SectionPredefinedColor.Descriptor instead.
func (SectionPredefinedColor) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_layout_proto_rawDescGZIP(), []int{0}
}

type Layout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sections []*Section `protobuf:"bytes,1,rep,name=sections,proto3" json:"sections,omitempty"`
}

func (x *Layout) Reset() {
	*x = Layout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_dashboards_v1_ast_layout_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Layout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Layout) ProtoMessage() {}

func (x *Layout) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_ast_layout_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Layout.ProtoReflect.Descriptor instead.
func (*Layout) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_layout_proto_rawDescGZIP(), []int{0}
}

func (x *Layout) GetSections() []*Section {
	if x != nil {
		return x.Sections
	}
	return nil
}

type Section struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      *UUID           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Rows    []*Row          `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows,omitempty"`
	Options *SectionOptions `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *Section) Reset() {
	*x = Section{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_dashboards_v1_ast_layout_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Section) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Section) ProtoMessage() {}

func (x *Section) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_ast_layout_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Section.ProtoReflect.Descriptor instead.
func (*Section) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_layout_proto_rawDescGZIP(), []int{1}
}

func (x *Section) GetId() *UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Section) GetRows() []*Row {
	if x != nil {
		return x.Rows
	}
	return nil
}

func (x *Section) GetOptions() *SectionOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

type Row struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         *UUID           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Appearance *Row_Appearance `protobuf:"bytes,2,opt,name=appearance,proto3" json:"appearance,omitempty"`
	Widgets    []*Widget       `protobuf:"bytes,3,rep,name=widgets,proto3" json:"widgets,omitempty"`
}

func (x *Row) Reset() {
	*x = Row{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_dashboards_v1_ast_layout_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Row) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Row) ProtoMessage() {}

func (x *Row) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_ast_layout_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Row.ProtoReflect.Descriptor instead.
func (*Row) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_layout_proto_rawDescGZIP(), []int{2}
}

func (x *Row) GetId() *UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Row) GetAppearance() *Row_Appearance {
	if x != nil {
		return x.Appearance
	}
	return nil
}

func (x *Row) GetWidgets() []*Widget {
	if x != nil {
		return x.Widgets
	}
	return nil
}

type SectionOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*SectionOptions_Internal
	//	*SectionOptions_Custom
	Value isSectionOptions_Value `protobuf_oneof:"value"`
}

func (x *SectionOptions) Reset() {
	*x = SectionOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_dashboards_v1_ast_layout_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SectionOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SectionOptions) ProtoMessage() {}

func (x *SectionOptions) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_ast_layout_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SectionOptions.ProtoReflect.Descriptor instead.
func (*SectionOptions) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_layout_proto_rawDescGZIP(), []int{3}
}

func (m *SectionOptions) GetValue() isSectionOptions_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *SectionOptions) GetInternal() *InternalSectionOptions {
	if x, ok := x.GetValue().(*SectionOptions_Internal); ok {
		return x.Internal
	}
	return nil
}

func (x *SectionOptions) GetCustom() *CustomSectionOptions {
	if x, ok := x.GetValue().(*SectionOptions_Custom); ok {
		return x.Custom
	}
	return nil
}

type isSectionOptions_Value interface {
	isSectionOptions_Value()
}

type SectionOptions_Internal struct {
	Internal *InternalSectionOptions `protobuf:"bytes,1,opt,name=internal,proto3,oneof"`
}

type SectionOptions_Custom struct {
	Custom *CustomSectionOptions `protobuf:"bytes,2,opt,name=custom,proto3,oneof"`
}

func (*SectionOptions_Internal) isSectionOptions_Value() {}

func (*SectionOptions_Custom) isSectionOptions_Value() {}

type InternalSectionOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InternalSectionOptions) Reset() {
	*x = InternalSectionOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_dashboards_v1_ast_layout_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalSectionOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalSectionOptions) ProtoMessage() {}

func (x *InternalSectionOptions) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_ast_layout_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalSectionOptions.ProtoReflect.Descriptor instead.
func (*InternalSectionOptions) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_layout_proto_rawDescGZIP(), []int{4}
}

type CustomSectionOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Collapsed   *wrapperspb.BoolValue   `protobuf:"bytes,3,opt,name=collapsed,proto3" json:"collapsed,omitempty"`
	Color       *SectionColor           `protobuf:"bytes,4,opt,name=color,proto3" json:"color,omitempty"`
}

func (x *CustomSectionOptions) Reset() {
	*x = CustomSectionOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_dashboards_v1_ast_layout_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomSectionOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomSectionOptions) ProtoMessage() {}

func (x *CustomSectionOptions) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_ast_layout_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomSectionOptions.ProtoReflect.Descriptor instead.
func (*CustomSectionOptions) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_layout_proto_rawDescGZIP(), []int{5}
}

func (x *CustomSectionOptions) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *CustomSectionOptions) GetDescription() *wrapperspb.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *CustomSectionOptions) GetCollapsed() *wrapperspb.BoolValue {
	if x != nil {
		return x.Collapsed
	}
	return nil
}

func (x *CustomSectionOptions) GetColor() *SectionColor {
	if x != nil {
		return x.Color
	}
	return nil
}

type SectionColor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*SectionColor_Predefined
	Value isSectionColor_Value `protobuf_oneof:"value"`
}

func (x *SectionColor) Reset() {
	*x = SectionColor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_dashboards_v1_ast_layout_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SectionColor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SectionColor) ProtoMessage() {}

func (x *SectionColor) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_ast_layout_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SectionColor.ProtoReflect.Descriptor instead.
func (*SectionColor) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_layout_proto_rawDescGZIP(), []int{6}
}

func (m *SectionColor) GetValue() isSectionColor_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *SectionColor) GetPredefined() SectionPredefinedColor {
	if x, ok := x.GetValue().(*SectionColor_Predefined); ok {
		return x.Predefined
	}
	return SectionPredefinedColor_SECTION_PREDEFINED_COLOR_UNSPECIFIED
}

type isSectionColor_Value interface {
	isSectionColor_Value()
}

type SectionColor_Predefined struct {
	Predefined SectionPredefinedColor `protobuf:"varint,1,opt,name=predefined,proto3,enum=com.coralogixapis.dashboards.v1.ast.SectionPredefinedColor,oneof"`
}

func (*SectionColor_Predefined) isSectionColor_Value() {}

type Row_Appearance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height *wrapperspb.Int32Value `protobuf:"bytes,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *Row_Appearance) Reset() {
	*x = Row_Appearance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_dashboards_v1_ast_layout_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Row_Appearance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Row_Appearance) ProtoMessage() {}

func (x *Row_Appearance) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_ast_layout_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Row_Appearance.ProtoReflect.Descriptor instead.
func (*Row_Appearance) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_layout_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Row_Appearance) GetHeight() *wrapperspb.Int32Value {
	if x != nil {
		return x.Height
	}
	return nil
}

var File_com_coralogixapis_dashboards_v1_ast_layout_proto protoreflect.FileDescriptor

var file_com_coralogixapis_dashboards_v1_ast_layout_proto_rawDesc = []byte{
	0x0a, 0x30, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x73, 0x74, 0x2f, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x74, 0x1a, 0x30, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x74, 0x2f, 0x77, 0x69, 0x64,
	0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x06, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74,
	0x12, 0x48, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x08, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xcd, 0x01, 0x0a, 0x07, 0x53,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3c, 0x0a,
	0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73,
	0x74, 0x2e, 0x52, 0x6f, 0x77, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x12, 0x4d, 0x0a, 0x07, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61,
	0x73, 0x74, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x9b, 0x02, 0x0a, 0x03, 0x52,
	0x6f, 0x77, 0x12, 0x35, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x53, 0x0a, 0x0a, 0x61, 0x70, 0x70,
	0x65, 0x61, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x61, 0x73, 0x74, 0x2e, 0x52, 0x6f, 0x77, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x61, 0x72, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x65, 0x61, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x45,
	0x0a, 0x07, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x52, 0x07, 0x77, 0x69,
	0x64, 0x67, 0x65, 0x74, 0x73, 0x1a, 0x41, 0x0a, 0x0a, 0x41, 0x70, 0x70, 0x65, 0x61, 0x72, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0xc9, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x59, 0x0a, 0x08, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x61, 0x73, 0x74, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x08, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x53, 0x0a, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x48, 0x00, 0x52, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x42, 0x07, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x8b,
	0x02, 0x0a, 0x14, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x30, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x6f, 0x6c,
	0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x70,
	0x73, 0x65, 0x64, 0x12, 0x47, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x22, 0x76, 0x0a, 0x0c,
	0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x5d, 0x0a, 0x0a,
	0x70, 0x72, 0x65, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x3b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x65, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x48, 0x00, 0x52,
	0x0a, 0x70, 0x72, 0x65, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x2a, 0xbf, 0x02, 0x0a, 0x16, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x72, 0x65, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12,
	0x28, 0x0a, 0x24, 0x53, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x52, 0x45, 0x44, 0x45,
	0x46, 0x49, 0x4e, 0x45, 0x44, 0x5f, 0x43, 0x4f, 0x4c, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x53, 0x45, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x52, 0x45, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x5f,
	0x43, 0x4f, 0x4c, 0x4f, 0x52, 0x5f, 0x43, 0x59, 0x41, 0x4e, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e,
	0x53, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x52, 0x45, 0x44, 0x45, 0x46, 0x49, 0x4e,
	0x45, 0x44, 0x5f, 0x43, 0x4f, 0x4c, 0x4f, 0x52, 0x5f, 0x47, 0x52, 0x45, 0x45, 0x4e, 0x10, 0x02,
	0x12, 0x21, 0x0a, 0x1d, 0x53, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x52, 0x45, 0x44,
	0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x5f, 0x43, 0x4f, 0x4c, 0x4f, 0x52, 0x5f, 0x42, 0x4c, 0x55,
	0x45, 0x10, 0x03, 0x12, 0x23, 0x0a, 0x1f, 0x53, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50,
	0x52, 0x45, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x5f, 0x43, 0x4f, 0x4c, 0x4f, 0x52, 0x5f,
	0x50, 0x55, 0x52, 0x50, 0x4c, 0x45, 0x10, 0x04, 0x12, 0x24, 0x0a, 0x20, 0x53, 0x45, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x52, 0x45, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x5f, 0x43,
	0x4f, 0x4c, 0x4f, 0x52, 0x5f, 0x4d, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x41, 0x10, 0x05, 0x12, 0x21,
	0x0a, 0x1d, 0x53, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x52, 0x45, 0x44, 0x45, 0x46,
	0x49, 0x4e, 0x45, 0x44, 0x5f, 0x43, 0x4f, 0x4c, 0x4f, 0x52, 0x5f, 0x50, 0x49, 0x4e, 0x4b, 0x10,
	0x06, 0x12, 0x23, 0x0a, 0x1f, 0x53, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x52, 0x45,
	0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x5f, 0x43, 0x4f, 0x4c, 0x4f, 0x52, 0x5f, 0x4f, 0x52,
	0x41, 0x4e, 0x47, 0x45, 0x10, 0x07, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_dashboards_v1_ast_layout_proto_rawDescOnce sync.Once
	file_com_coralogixapis_dashboards_v1_ast_layout_proto_rawDescData = file_com_coralogixapis_dashboards_v1_ast_layout_proto_rawDesc
)

func file_com_coralogixapis_dashboards_v1_ast_layout_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_dashboards_v1_ast_layout_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_dashboards_v1_ast_layout_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_dashboards_v1_ast_layout_proto_rawDescData)
	})
	return file_com_coralogixapis_dashboards_v1_ast_layout_proto_rawDescData
}

var file_com_coralogixapis_dashboards_v1_ast_layout_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_dashboards_v1_ast_layout_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_com_coralogixapis_dashboards_v1_ast_layout_proto_goTypes = []interface{}{
	(SectionPredefinedColor)(0),    // 0: com.coralogixapis.dashboards.v1.ast.SectionPredefinedColor
	(*Layout)(nil),                 // 1: com.coralogixapis.dashboards.v1.ast.Layout
	(*Section)(nil),                // 2: com.coralogixapis.dashboards.v1.ast.Section
	(*Row)(nil),                    // 3: com.coralogixapis.dashboards.v1.ast.Row
	(*SectionOptions)(nil),         // 4: com.coralogixapis.dashboards.v1.ast.SectionOptions
	(*InternalSectionOptions)(nil), // 5: com.coralogixapis.dashboards.v1.ast.InternalSectionOptions
	(*CustomSectionOptions)(nil),   // 6: com.coralogixapis.dashboards.v1.ast.CustomSectionOptions
	(*SectionColor)(nil),           // 7: com.coralogixapis.dashboards.v1.ast.SectionColor
	(*Row_Appearance)(nil),         // 8: com.coralogixapis.dashboards.v1.ast.Row.Appearance
	(*UUID)(nil),                   // 9: com.coralogixapis.dashboards.v1.UUID
	(*Widget)(nil),                 // 10: com.coralogixapis.dashboards.v1.ast.Widget
	(*wrapperspb.StringValue)(nil), // 11: google.protobuf.StringValue
	(*wrapperspb.BoolValue)(nil),   // 12: google.protobuf.BoolValue
	(*wrapperspb.Int32Value)(nil),  // 13: google.protobuf.Int32Value
}
var file_com_coralogixapis_dashboards_v1_ast_layout_proto_depIdxs = []int32{
	2,  // 0: com.coralogixapis.dashboards.v1.ast.Layout.sections:type_name -> com.coralogixapis.dashboards.v1.ast.Section
	9,  // 1: com.coralogixapis.dashboards.v1.ast.Section.id:type_name -> com.coralogixapis.dashboards.v1.UUID
	3,  // 2: com.coralogixapis.dashboards.v1.ast.Section.rows:type_name -> com.coralogixapis.dashboards.v1.ast.Row
	4,  // 3: com.coralogixapis.dashboards.v1.ast.Section.options:type_name -> com.coralogixapis.dashboards.v1.ast.SectionOptions
	9,  // 4: com.coralogixapis.dashboards.v1.ast.Row.id:type_name -> com.coralogixapis.dashboards.v1.UUID
	8,  // 5: com.coralogixapis.dashboards.v1.ast.Row.appearance:type_name -> com.coralogixapis.dashboards.v1.ast.Row.Appearance
	10, // 6: com.coralogixapis.dashboards.v1.ast.Row.widgets:type_name -> com.coralogixapis.dashboards.v1.ast.Widget
	5,  // 7: com.coralogixapis.dashboards.v1.ast.SectionOptions.internal:type_name -> com.coralogixapis.dashboards.v1.ast.InternalSectionOptions
	6,  // 8: com.coralogixapis.dashboards.v1.ast.SectionOptions.custom:type_name -> com.coralogixapis.dashboards.v1.ast.CustomSectionOptions
	11, // 9: com.coralogixapis.dashboards.v1.ast.CustomSectionOptions.name:type_name -> google.protobuf.StringValue
	11, // 10: com.coralogixapis.dashboards.v1.ast.CustomSectionOptions.description:type_name -> google.protobuf.StringValue
	12, // 11: com.coralogixapis.dashboards.v1.ast.CustomSectionOptions.collapsed:type_name -> google.protobuf.BoolValue
	7,  // 12: com.coralogixapis.dashboards.v1.ast.CustomSectionOptions.color:type_name -> com.coralogixapis.dashboards.v1.ast.SectionColor
	0,  // 13: com.coralogixapis.dashboards.v1.ast.SectionColor.predefined:type_name -> com.coralogixapis.dashboards.v1.ast.SectionPredefinedColor
	13, // 14: com.coralogixapis.dashboards.v1.ast.Row.Appearance.height:type_name -> google.protobuf.Int32Value
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_dashboards_v1_ast_layout_proto_init() }
func file_com_coralogixapis_dashboards_v1_ast_layout_proto_init() {
	if File_com_coralogixapis_dashboards_v1_ast_layout_proto != nil {
		return
	}
	file_com_coralogixapis_dashboards_v1_ast_widget_proto_init()
	file_com_coralogixapis_dashboards_v1_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_dashboards_v1_ast_layout_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Layout); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_com_coralogixapis_dashboards_v1_ast_layout_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Section); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_com_coralogixapis_dashboards_v1_ast_layout_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Row); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_com_coralogixapis_dashboards_v1_ast_layout_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SectionOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_com_coralogixapis_dashboards_v1_ast_layout_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalSectionOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_com_coralogixapis_dashboards_v1_ast_layout_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomSectionOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_com_coralogixapis_dashboards_v1_ast_layout_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SectionColor); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_com_coralogixapis_dashboards_v1_ast_layout_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Row_Appearance); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_com_coralogixapis_dashboards_v1_ast_layout_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*SectionOptions_Internal)(nil),
		(*SectionOptions_Custom)(nil),
	}
	file_com_coralogixapis_dashboards_v1_ast_layout_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*SectionColor_Predefined)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_dashboards_v1_ast_layout_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_dashboards_v1_ast_layout_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_dashboards_v1_ast_layout_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_dashboards_v1_ast_layout_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_dashboards_v1_ast_layout_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_dashboards_v1_ast_layout_proto = out.File
	file_com_coralogixapis_dashboards_v1_ast_layout_proto_rawDesc = nil
	file_com_coralogixapis_dashboards_v1_ast_layout_proto_goTypes = nil
	file_com_coralogixapis_dashboards_v1_ast_layout_proto_depIdxs = nil
}
