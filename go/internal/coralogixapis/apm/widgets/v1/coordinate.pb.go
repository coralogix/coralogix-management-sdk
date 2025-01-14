// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: com/coralogixapis/apm/widgets/v1/coordinate.proto

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

type Coordinate struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Lat           *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Long          *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=long,proto3" json:"long,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Coordinate) Reset() {
	*x = Coordinate{}
	mi := &file_com_coralogixapis_apm_widgets_v1_coordinate_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Coordinate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coordinate) ProtoMessage() {}

func (x *Coordinate) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_widgets_v1_coordinate_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coordinate.ProtoReflect.Descriptor instead.
func (*Coordinate) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_widgets_v1_coordinate_proto_rawDescGZIP(), []int{0}
}

func (x *Coordinate) GetLat() *wrapperspb.StringValue {
	if x != nil {
		return x.Lat
	}
	return nil
}

func (x *Coordinate) GetLong() *wrapperspb.StringValue {
	if x != nil {
		return x.Long
	}
	return nil
}

var File_com_coralogixapis_apm_widgets_v1_coordinate_proto protoreflect.FileDescriptor

var file_com_coralogixapis_apm_widgets_v1_coordinate_proto_rawDesc = []byte{
	0x0a, 0x31, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x0a, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03,
	0x6c, 0x61, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_apm_widgets_v1_coordinate_proto_rawDescOnce sync.Once
	file_com_coralogixapis_apm_widgets_v1_coordinate_proto_rawDescData = file_com_coralogixapis_apm_widgets_v1_coordinate_proto_rawDesc
)

func file_com_coralogixapis_apm_widgets_v1_coordinate_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_apm_widgets_v1_coordinate_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_apm_widgets_v1_coordinate_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_apm_widgets_v1_coordinate_proto_rawDescData)
	})
	return file_com_coralogixapis_apm_widgets_v1_coordinate_proto_rawDescData
}

var file_com_coralogixapis_apm_widgets_v1_coordinate_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_apm_widgets_v1_coordinate_proto_goTypes = []any{
	(*Coordinate)(nil),             // 0: com.coralogixapis.apm.widgets.v1.Coordinate
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
}
var file_com_coralogixapis_apm_widgets_v1_coordinate_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.apm.widgets.v1.Coordinate.lat:type_name -> google.protobuf.StringValue
	1, // 1: com.coralogixapis.apm.widgets.v1.Coordinate.long:type_name -> google.protobuf.StringValue
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_apm_widgets_v1_coordinate_proto_init() }
func file_com_coralogixapis_apm_widgets_v1_coordinate_proto_init() {
	if File_com_coralogixapis_apm_widgets_v1_coordinate_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_apm_widgets_v1_coordinate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_apm_widgets_v1_coordinate_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_apm_widgets_v1_coordinate_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_apm_widgets_v1_coordinate_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_apm_widgets_v1_coordinate_proto = out.File
	file_com_coralogixapis_apm_widgets_v1_coordinate_proto_rawDesc = nil
	file_com_coralogixapis_apm_widgets_v1_coordinate_proto_goTypes = nil
	file_com_coralogixapis_apm_widgets_v1_coordinate_proto_depIdxs = nil
}
