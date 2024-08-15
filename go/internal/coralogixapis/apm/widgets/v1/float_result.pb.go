// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: com/coralogixapis/apm/widgets/v1/float_result.proto

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

// A wrapper for a floating point result
type FloatResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value       *wrapperspb.FloatValue             `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Name        *wrapperspb.StringValue            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // This is the key name
	DisplayName *wrapperspb.StringValue            `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Unit        *wrapperspb.StringValue            `protobuf:"bytes,4,opt,name=unit,proto3" json:"unit,omitempty"` // The unit of measurement
	Query       *wrapperspb.StringValue            `protobuf:"bytes,5,opt,name=query,proto3" json:"query,omitempty"`
	Metric      map[string]*wrapperspb.StringValue `protobuf:"bytes,6,rep,name=metric,proto3" json:"metric,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *FloatResult) Reset() {
	*x = FloatResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_apm_widgets_v1_float_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FloatResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FloatResult) ProtoMessage() {}

func (x *FloatResult) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_widgets_v1_float_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FloatResult.ProtoReflect.Descriptor instead.
func (*FloatResult) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_widgets_v1_float_result_proto_rawDescGZIP(), []int{0}
}

func (x *FloatResult) GetValue() *wrapperspb.FloatValue {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *FloatResult) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *FloatResult) GetDisplayName() *wrapperspb.StringValue {
	if x != nil {
		return x.DisplayName
	}
	return nil
}

func (x *FloatResult) GetUnit() *wrapperspb.StringValue {
	if x != nil {
		return x.Unit
	}
	return nil
}

func (x *FloatResult) GetQuery() *wrapperspb.StringValue {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *FloatResult) GetMetric() map[string]*wrapperspb.StringValue {
	if x != nil {
		return x.Metric
	}
	return nil
}

var File_com_coralogixapis_apm_widgets_v1_float_result_proto protoreflect.FileDescriptor

var file_com_coralogixapis_apm_widgets_v1_float_result_proto_rawDesc = []byte{
	0x0a, 0x33, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x77, 0x69, 0x64,
	0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x03, 0x0a, 0x0b, 0x46, 0x6c, 0x6f, 0x61,
	0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x0c,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a,
	0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12,
	0x32, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x51, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x77, 0x69, 0x64, 0x67,
	0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x1a, 0x57, 0x0a, 0x0b, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_apm_widgets_v1_float_result_proto_rawDescOnce sync.Once
	file_com_coralogixapis_apm_widgets_v1_float_result_proto_rawDescData = file_com_coralogixapis_apm_widgets_v1_float_result_proto_rawDesc
)

func file_com_coralogixapis_apm_widgets_v1_float_result_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_apm_widgets_v1_float_result_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_apm_widgets_v1_float_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_apm_widgets_v1_float_result_proto_rawDescData)
	})
	return file_com_coralogixapis_apm_widgets_v1_float_result_proto_rawDescData
}

var file_com_coralogixapis_apm_widgets_v1_float_result_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogixapis_apm_widgets_v1_float_result_proto_goTypes = []any{
	(*FloatResult)(nil),            // 0: com.coralogixapis.apm.widgets.v1.FloatResult
	nil,                            // 1: com.coralogixapis.apm.widgets.v1.FloatResult.MetricEntry
	(*wrapperspb.FloatValue)(nil),  // 2: google.protobuf.FloatValue
	(*wrapperspb.StringValue)(nil), // 3: google.protobuf.StringValue
}
var file_com_coralogixapis_apm_widgets_v1_float_result_proto_depIdxs = []int32{
	2, // 0: com.coralogixapis.apm.widgets.v1.FloatResult.value:type_name -> google.protobuf.FloatValue
	3, // 1: com.coralogixapis.apm.widgets.v1.FloatResult.name:type_name -> google.protobuf.StringValue
	3, // 2: com.coralogixapis.apm.widgets.v1.FloatResult.display_name:type_name -> google.protobuf.StringValue
	3, // 3: com.coralogixapis.apm.widgets.v1.FloatResult.unit:type_name -> google.protobuf.StringValue
	3, // 4: com.coralogixapis.apm.widgets.v1.FloatResult.query:type_name -> google.protobuf.StringValue
	1, // 5: com.coralogixapis.apm.widgets.v1.FloatResult.metric:type_name -> com.coralogixapis.apm.widgets.v1.FloatResult.MetricEntry
	3, // 6: com.coralogixapis.apm.widgets.v1.FloatResult.MetricEntry.value:type_name -> google.protobuf.StringValue
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_apm_widgets_v1_float_result_proto_init() }
func file_com_coralogixapis_apm_widgets_v1_float_result_proto_init() {
	if File_com_coralogixapis_apm_widgets_v1_float_result_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_apm_widgets_v1_float_result_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*FloatResult); i {
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
			RawDescriptor: file_com_coralogixapis_apm_widgets_v1_float_result_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_apm_widgets_v1_float_result_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_apm_widgets_v1_float_result_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_apm_widgets_v1_float_result_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_apm_widgets_v1_float_result_proto = out.File
	file_com_coralogixapis_apm_widgets_v1_float_result_proto_rawDesc = nil
	file_com_coralogixapis_apm_widgets_v1_float_result_proto_goTypes = nil
	file_com_coralogixapis_apm_widgets_v1_float_result_proto_depIdxs = nil
}
