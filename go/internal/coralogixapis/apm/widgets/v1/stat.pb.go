// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: com/coralogixapis/apm/widgets/v1/stat.proto

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

type Stat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DisplayName *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Name        *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Query       *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty"`
	Value       *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	ToolTip     *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=tool_tip,json=toolTip,proto3" json:"tool_tip,omitempty"`
	Color       *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=color,proto3" json:"color,omitempty"`
	Unit        Unit                    `protobuf:"varint,7,opt,name=unit,proto3,enum=com.coralogixapis.apm.widgets.v1.Unit" json:"unit,omitempty"`
}

func (x *Stat) Reset() {
	*x = Stat{}
	mi := &file_com_coralogixapis_apm_widgets_v1_stat_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Stat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stat) ProtoMessage() {}

func (x *Stat) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_widgets_v1_stat_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stat.ProtoReflect.Descriptor instead.
func (*Stat) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_widgets_v1_stat_proto_rawDescGZIP(), []int{0}
}

func (x *Stat) GetDisplayName() *wrapperspb.StringValue {
	if x != nil {
		return x.DisplayName
	}
	return nil
}

func (x *Stat) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *Stat) GetQuery() *wrapperspb.StringValue {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *Stat) GetValue() *wrapperspb.StringValue {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Stat) GetToolTip() *wrapperspb.StringValue {
	if x != nil {
		return x.ToolTip
	}
	return nil
}

func (x *Stat) GetColor() *wrapperspb.StringValue {
	if x != nil {
		return x.Color
	}
	return nil
}

func (x *Stat) GetUnit() Unit {
	if x != nil {
		return x.Unit
	}
	return Unit_UNIT_UNSPECIFIED
}

var File_com_coralogixapis_apm_widgets_v1_stat_proto protoreflect.FileDescriptor

var file_com_coralogixapis_apm_widgets_v1_stat_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2c, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x03,
	0x0a, 0x04, 0x53, 0x74, 0x61, 0x74, 0x12, 0x3f, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x32, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x37, 0x0a, 0x08, 0x74, 0x6f, 0x6f, 0x6c, 0x5f, 0x74, 0x69, 0x70, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x07, 0x74, 0x6f, 0x6f, 0x6c, 0x54, 0x69, 0x70, 0x12, 0x32, 0x0a, 0x05, 0x63, 0x6f,
	0x6c, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x3a,
	0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x6e, 0x69, 0x74, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_com_coralogixapis_apm_widgets_v1_stat_proto_rawDescOnce sync.Once
	file_com_coralogixapis_apm_widgets_v1_stat_proto_rawDescData = file_com_coralogixapis_apm_widgets_v1_stat_proto_rawDesc
)

func file_com_coralogixapis_apm_widgets_v1_stat_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_apm_widgets_v1_stat_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_apm_widgets_v1_stat_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_apm_widgets_v1_stat_proto_rawDescData)
	})
	return file_com_coralogixapis_apm_widgets_v1_stat_proto_rawDescData
}

var file_com_coralogixapis_apm_widgets_v1_stat_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_apm_widgets_v1_stat_proto_goTypes = []any{
	(*Stat)(nil),                   // 0: com.coralogixapis.apm.widgets.v1.Stat
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
	(Unit)(0),                      // 2: com.coralogixapis.apm.widgets.v1.Unit
}
var file_com_coralogixapis_apm_widgets_v1_stat_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.apm.widgets.v1.Stat.display_name:type_name -> google.protobuf.StringValue
	1, // 1: com.coralogixapis.apm.widgets.v1.Stat.name:type_name -> google.protobuf.StringValue
	1, // 2: com.coralogixapis.apm.widgets.v1.Stat.query:type_name -> google.protobuf.StringValue
	1, // 3: com.coralogixapis.apm.widgets.v1.Stat.value:type_name -> google.protobuf.StringValue
	1, // 4: com.coralogixapis.apm.widgets.v1.Stat.tool_tip:type_name -> google.protobuf.StringValue
	1, // 5: com.coralogixapis.apm.widgets.v1.Stat.color:type_name -> google.protobuf.StringValue
	2, // 6: com.coralogixapis.apm.widgets.v1.Stat.unit:type_name -> com.coralogixapis.apm.widgets.v1.Unit
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_apm_widgets_v1_stat_proto_init() }
func file_com_coralogixapis_apm_widgets_v1_stat_proto_init() {
	if File_com_coralogixapis_apm_widgets_v1_stat_proto != nil {
		return
	}
	file_com_coralogixapis_apm_widgets_v1_units_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_apm_widgets_v1_stat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_apm_widgets_v1_stat_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_apm_widgets_v1_stat_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_apm_widgets_v1_stat_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_apm_widgets_v1_stat_proto = out.File
	file_com_coralogixapis_apm_widgets_v1_stat_proto_rawDesc = nil
	file_com_coralogixapis_apm_widgets_v1_stat_proto_goTypes = nil
	file_com_coralogixapis_apm_widgets_v1_stat_proto_depIdxs = nil
}
