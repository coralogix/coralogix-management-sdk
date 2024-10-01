// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: com/coralogixapis/apm/widgets/v1/map_chart.proto

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

type MapChart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DisplayName *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Name        *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Query       *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty"`
	// Deprecated: Marked as deprecated in com/coralogixapis/apm/widgets/v1/map_chart.proto.
	Results   []*FloatResult   `protobuf:"bytes,4,rep,name=results,proto3" json:"results,omitempty"`
	MapPoints []*MapPointValue `protobuf:"bytes,5,rep,name=map_points,json=mapPoints,proto3" json:"map_points,omitempty"`
}

func (x *MapChart) Reset() {
	*x = MapChart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_apm_widgets_v1_map_chart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapChart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapChart) ProtoMessage() {}

func (x *MapChart) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_widgets_v1_map_chart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapChart.ProtoReflect.Descriptor instead.
func (*MapChart) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_widgets_v1_map_chart_proto_rawDescGZIP(), []int{0}
}

func (x *MapChart) GetDisplayName() *wrapperspb.StringValue {
	if x != nil {
		return x.DisplayName
	}
	return nil
}

func (x *MapChart) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *MapChart) GetQuery() *wrapperspb.StringValue {
	if x != nil {
		return x.Query
	}
	return nil
}

// Deprecated: Marked as deprecated in com/coralogixapis/apm/widgets/v1/map_chart.proto.
func (x *MapChart) GetResults() []*FloatResult {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *MapChart) GetMapPoints() []*MapPointValue {
	if x != nil {
		return x.MapPoints
	}
	return nil
}

var File_com_coralogixapis_apm_widgets_v1_map_chart_proto protoreflect.FileDescriptor

var file_com_coralogixapis_apm_widgets_v1_map_chart_proto_rawDesc = []byte{
	0x0a, 0x30, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x61, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x77, 0x69, 0x64, 0x67,
	0x65, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d,
	0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x70, 0x5f,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xce, 0x02, 0x0a, 0x08, 0x4d, 0x61, 0x70, 0x43, 0x68, 0x61, 0x72, 0x74, 0x12, 0x3f,
	0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x30, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x32, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x4b, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x77,
	0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x02, 0x18, 0x01, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x12, 0x4e, 0x0a, 0x0a, 0x6d, 0x61, 0x70, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x77,
	0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x70, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x6d, 0x61, 0x70, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_apm_widgets_v1_map_chart_proto_rawDescOnce sync.Once
	file_com_coralogixapis_apm_widgets_v1_map_chart_proto_rawDescData = file_com_coralogixapis_apm_widgets_v1_map_chart_proto_rawDesc
)

func file_com_coralogixapis_apm_widgets_v1_map_chart_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_apm_widgets_v1_map_chart_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_apm_widgets_v1_map_chart_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_apm_widgets_v1_map_chart_proto_rawDescData)
	})
	return file_com_coralogixapis_apm_widgets_v1_map_chart_proto_rawDescData
}

var file_com_coralogixapis_apm_widgets_v1_map_chart_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_apm_widgets_v1_map_chart_proto_goTypes = []any{
	(*MapChart)(nil),               // 0: com.coralogixapis.apm.widgets.v1.MapChart
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
	(*FloatResult)(nil),            // 2: com.coralogixapis.apm.widgets.v1.FloatResult
	(*MapPointValue)(nil),          // 3: com.coralogixapis.apm.widgets.v1.MapPointValue
}
var file_com_coralogixapis_apm_widgets_v1_map_chart_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.apm.widgets.v1.MapChart.display_name:type_name -> google.protobuf.StringValue
	1, // 1: com.coralogixapis.apm.widgets.v1.MapChart.name:type_name -> google.protobuf.StringValue
	1, // 2: com.coralogixapis.apm.widgets.v1.MapChart.query:type_name -> google.protobuf.StringValue
	2, // 3: com.coralogixapis.apm.widgets.v1.MapChart.results:type_name -> com.coralogixapis.apm.widgets.v1.FloatResult
	3, // 4: com.coralogixapis.apm.widgets.v1.MapChart.map_points:type_name -> com.coralogixapis.apm.widgets.v1.MapPointValue
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_apm_widgets_v1_map_chart_proto_init() }
func file_com_coralogixapis_apm_widgets_v1_map_chart_proto_init() {
	if File_com_coralogixapis_apm_widgets_v1_map_chart_proto != nil {
		return
	}
	file_com_coralogixapis_apm_widgets_v1_float_result_proto_init()
	file_com_coralogixapis_apm_widgets_v1_map_point_value_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_apm_widgets_v1_map_chart_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*MapChart); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_apm_widgets_v1_map_chart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_apm_widgets_v1_map_chart_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_apm_widgets_v1_map_chart_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_apm_widgets_v1_map_chart_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_apm_widgets_v1_map_chart_proto = out.File
	file_com_coralogixapis_apm_widgets_v1_map_chart_proto_rawDesc = nil
	file_com_coralogixapis_apm_widgets_v1_map_chart_proto_goTypes = nil
	file_com_coralogixapis_apm_widgets_v1_map_chart_proto_depIdxs = nil
}
