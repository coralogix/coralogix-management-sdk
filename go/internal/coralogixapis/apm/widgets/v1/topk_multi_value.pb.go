// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: com/coralogixapis/apm/widgets/v1/topk_multi_value.proto

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

type TopkMultiValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group       *WidgetGroup                       `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Name        *wrapperspb.StringValue            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName *wrapperspb.StringValue            `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Rows        []*RowWrapper                      `protobuf:"bytes,4,rep,name=rows,proto3" json:"rows,omitempty"`
	Query       map[string]*wrapperspb.StringValue `protobuf:"bytes,5,rep,name=query,proto3" json:"query,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ToolTip     *wrapperspb.StringValue            `protobuf:"bytes,6,opt,name=tool_tip,json=toolTip,proto3" json:"tool_tip,omitempty"`
}

func (x *TopkMultiValue) Reset() {
	*x = TopkMultiValue{}
	mi := &file_com_coralogixapis_apm_widgets_v1_topk_multi_value_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TopkMultiValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopkMultiValue) ProtoMessage() {}

func (x *TopkMultiValue) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_widgets_v1_topk_multi_value_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopkMultiValue.ProtoReflect.Descriptor instead.
func (*TopkMultiValue) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_widgets_v1_topk_multi_value_proto_rawDescGZIP(), []int{0}
}

func (x *TopkMultiValue) GetGroup() *WidgetGroup {
	if x != nil {
		return x.Group
	}
	return nil
}

func (x *TopkMultiValue) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *TopkMultiValue) GetDisplayName() *wrapperspb.StringValue {
	if x != nil {
		return x.DisplayName
	}
	return nil
}

func (x *TopkMultiValue) GetRows() []*RowWrapper {
	if x != nil {
		return x.Rows
	}
	return nil
}

func (x *TopkMultiValue) GetQuery() map[string]*wrapperspb.StringValue {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *TopkMultiValue) GetToolTip() *wrapperspb.StringValue {
	if x != nil {
		return x.ToolTip
	}
	return nil
}

var File_com_coralogixapis_apm_widgets_v1_topk_multi_value_proto protoreflect.FileDescriptor

var file_com_coralogixapis_apm_widgets_v1_topk_multi_value_proto_rawDesc = []byte{
	0x0a, 0x37, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x6f, 0x70, 0x6b, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d,
	0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x6d, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x69,
	0x64, 0x67, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x32, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x6f, 0x77, 0x5f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x03, 0x0a, 0x0e, 0x54, 0x6f, 0x70, 0x6b, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x43, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x77,
	0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x30, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3f,
	0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x40, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x6f, 0x77, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x52, 0x04, 0x72, 0x6f, 0x77,
	0x73, 0x12, 0x51, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x3b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70, 0x6b, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x37, 0x0a, 0x08, 0x74, 0x6f, 0x6f, 0x6c, 0x5f, 0x74, 0x69, 0x70,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x74, 0x6f, 0x6f, 0x6c, 0x54, 0x69, 0x70, 0x1a, 0x56, 0x0a,
	0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x32, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_apm_widgets_v1_topk_multi_value_proto_rawDescOnce sync.Once
	file_com_coralogixapis_apm_widgets_v1_topk_multi_value_proto_rawDescData = file_com_coralogixapis_apm_widgets_v1_topk_multi_value_proto_rawDesc
)

func file_com_coralogixapis_apm_widgets_v1_topk_multi_value_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_apm_widgets_v1_topk_multi_value_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_apm_widgets_v1_topk_multi_value_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_apm_widgets_v1_topk_multi_value_proto_rawDescData)
	})
	return file_com_coralogixapis_apm_widgets_v1_topk_multi_value_proto_rawDescData
}

var file_com_coralogixapis_apm_widgets_v1_topk_multi_value_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogixapis_apm_widgets_v1_topk_multi_value_proto_goTypes = []any{
	(*TopkMultiValue)(nil),         // 0: com.coralogixapis.apm.widgets.v1.TopkMultiValue
	nil,                            // 1: com.coralogixapis.apm.widgets.v1.TopkMultiValue.QueryEntry
	(*WidgetGroup)(nil),            // 2: com.coralogixapis.apm.widgets.v1.WidgetGroup
	(*wrapperspb.StringValue)(nil), // 3: google.protobuf.StringValue
	(*RowWrapper)(nil),             // 4: com.coralogixapis.apm.widgets.v1.RowWrapper
}
var file_com_coralogixapis_apm_widgets_v1_topk_multi_value_proto_depIdxs = []int32{
	2, // 0: com.coralogixapis.apm.widgets.v1.TopkMultiValue.group:type_name -> com.coralogixapis.apm.widgets.v1.WidgetGroup
	3, // 1: com.coralogixapis.apm.widgets.v1.TopkMultiValue.name:type_name -> google.protobuf.StringValue
	3, // 2: com.coralogixapis.apm.widgets.v1.TopkMultiValue.display_name:type_name -> google.protobuf.StringValue
	4, // 3: com.coralogixapis.apm.widgets.v1.TopkMultiValue.rows:type_name -> com.coralogixapis.apm.widgets.v1.RowWrapper
	1, // 4: com.coralogixapis.apm.widgets.v1.TopkMultiValue.query:type_name -> com.coralogixapis.apm.widgets.v1.TopkMultiValue.QueryEntry
	3, // 5: com.coralogixapis.apm.widgets.v1.TopkMultiValue.tool_tip:type_name -> google.protobuf.StringValue
	3, // 6: com.coralogixapis.apm.widgets.v1.TopkMultiValue.QueryEntry.value:type_name -> google.protobuf.StringValue
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_apm_widgets_v1_topk_multi_value_proto_init() }
func file_com_coralogixapis_apm_widgets_v1_topk_multi_value_proto_init() {
	if File_com_coralogixapis_apm_widgets_v1_topk_multi_value_proto != nil {
		return
	}
	file_com_coralogixapis_apm_widgets_v1_widget_group_proto_init()
	file_com_coralogixapis_apm_widgets_v1_row_wrapper_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_apm_widgets_v1_topk_multi_value_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_apm_widgets_v1_topk_multi_value_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_apm_widgets_v1_topk_multi_value_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_apm_widgets_v1_topk_multi_value_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_apm_widgets_v1_topk_multi_value_proto = out.File
	file_com_coralogixapis_apm_widgets_v1_topk_multi_value_proto_rawDesc = nil
	file_com_coralogixapis_apm_widgets_v1_topk_multi_value_proto_goTypes = nil
	file_com_coralogixapis_apm_widgets_v1_topk_multi_value_proto_depIdxs = nil
}
