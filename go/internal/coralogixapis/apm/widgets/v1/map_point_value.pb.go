// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: com/coralogixapis/apm/widgets/v1/map_point_value.proto

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

// A wrapper for a map point result
type MapPointValue struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Value         *wrapperspb.FloatValue  `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Coordinate    *Coordinate             `protobuf:"bytes,2,opt,name=coordinate,proto3" json:"coordinate,omitempty"`
	Name          *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"` // This is the key name
	DisplayName   *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Unit          *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=unit,proto3" json:"unit,omitempty"` // The unit of measurement
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MapPointValue) Reset() {
	*x = MapPointValue{}
	mi := &file_com_coralogixapis_apm_widgets_v1_map_point_value_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MapPointValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapPointValue) ProtoMessage() {}

func (x *MapPointValue) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_widgets_v1_map_point_value_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapPointValue.ProtoReflect.Descriptor instead.
func (*MapPointValue) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_widgets_v1_map_point_value_proto_rawDescGZIP(), []int{0}
}

func (x *MapPointValue) GetValue() *wrapperspb.FloatValue {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *MapPointValue) GetCoordinate() *Coordinate {
	if x != nil {
		return x.Coordinate
	}
	return nil
}

func (x *MapPointValue) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *MapPointValue) GetDisplayName() *wrapperspb.StringValue {
	if x != nil {
		return x.DisplayName
	}
	return nil
}

func (x *MapPointValue) GetUnit() *wrapperspb.StringValue {
	if x != nil {
		return x.Unit
	}
	return nil
}

var File_com_coralogixapis_apm_widgets_v1_map_point_value_proto protoreflect.FileDescriptor

var file_com_coralogixapis_apm_widgets_v1_map_point_value_proto_rawDesc = []byte{
	0x0a, 0x36, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x61, 0x70, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e,
	0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70,
	0x6d, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x02,
	0x0a, 0x0d, 0x4d, 0x61, 0x70, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x4c, 0x0a, 0x0a, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x77,
	0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65,
	0x12, 0x30, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x04, 0x75, 0x6e, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_apm_widgets_v1_map_point_value_proto_rawDescOnce sync.Once
	file_com_coralogixapis_apm_widgets_v1_map_point_value_proto_rawDescData = file_com_coralogixapis_apm_widgets_v1_map_point_value_proto_rawDesc
)

func file_com_coralogixapis_apm_widgets_v1_map_point_value_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_apm_widgets_v1_map_point_value_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_apm_widgets_v1_map_point_value_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_apm_widgets_v1_map_point_value_proto_rawDescData)
	})
	return file_com_coralogixapis_apm_widgets_v1_map_point_value_proto_rawDescData
}

var file_com_coralogixapis_apm_widgets_v1_map_point_value_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_apm_widgets_v1_map_point_value_proto_goTypes = []any{
	(*MapPointValue)(nil),          // 0: com.coralogixapis.apm.widgets.v1.MapPointValue
	(*wrapperspb.FloatValue)(nil),  // 1: google.protobuf.FloatValue
	(*Coordinate)(nil),             // 2: com.coralogixapis.apm.widgets.v1.Coordinate
	(*wrapperspb.StringValue)(nil), // 3: google.protobuf.StringValue
}
var file_com_coralogixapis_apm_widgets_v1_map_point_value_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.apm.widgets.v1.MapPointValue.value:type_name -> google.protobuf.FloatValue
	2, // 1: com.coralogixapis.apm.widgets.v1.MapPointValue.coordinate:type_name -> com.coralogixapis.apm.widgets.v1.Coordinate
	3, // 2: com.coralogixapis.apm.widgets.v1.MapPointValue.name:type_name -> google.protobuf.StringValue
	3, // 3: com.coralogixapis.apm.widgets.v1.MapPointValue.display_name:type_name -> google.protobuf.StringValue
	3, // 4: com.coralogixapis.apm.widgets.v1.MapPointValue.unit:type_name -> google.protobuf.StringValue
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_apm_widgets_v1_map_point_value_proto_init() }
func file_com_coralogixapis_apm_widgets_v1_map_point_value_proto_init() {
	if File_com_coralogixapis_apm_widgets_v1_map_point_value_proto != nil {
		return
	}
	file_com_coralogixapis_apm_widgets_v1_coordinate_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_apm_widgets_v1_map_point_value_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_apm_widgets_v1_map_point_value_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_apm_widgets_v1_map_point_value_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_apm_widgets_v1_map_point_value_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_apm_widgets_v1_map_point_value_proto = out.File
	file_com_coralogixapis_apm_widgets_v1_map_point_value_proto_rawDesc = nil
	file_com_coralogixapis_apm_widgets_v1_map_point_value_proto_goTypes = nil
	file_com_coralogixapis_apm_widgets_v1_map_point_value_proto_depIdxs = nil
}
