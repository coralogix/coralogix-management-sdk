// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.3
// source: com/coralogixapis/dashboards/v1/common/labelled_value.proto

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

type LabelledValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Labels map[string]string       `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Value  *wrapperspb.DoubleValue `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *LabelledValue) Reset() {
	*x = LabelledValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelledValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelledValue) ProtoMessage() {}

func (x *LabelledValue) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelledValue.ProtoReflect.Descriptor instead.
func (*LabelledValue) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_rawDescGZIP(), []int{0}
}

func (x *LabelledValue) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *LabelledValue) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *LabelledValue) GetValue() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_com_coralogixapis_dashboards_v1_common_labelled_value_proto protoreflect.FileDescriptor

var file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x6c, 0x65,
	0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x02, 0x0a, 0x0d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x6c,
	0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x59, 0x0a, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_rawDescOnce sync.Once
	file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_rawDescData = file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_rawDesc
)

func file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_rawDescData)
	})
	return file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_rawDescData
}

var file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_goTypes = []interface{}{
	(*LabelledValue)(nil),          // 0: com.coralogixapis.dashboards.v1.common.LabelledValue
	nil,                            // 1: com.coralogixapis.dashboards.v1.common.LabelledValue.LabelsEntry
	(*wrapperspb.StringValue)(nil), // 2: google.protobuf.StringValue
	(*wrapperspb.DoubleValue)(nil), // 3: google.protobuf.DoubleValue
}
var file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_depIdxs = []int32{
	2, // 0: com.coralogixapis.dashboards.v1.common.LabelledValue.name:type_name -> google.protobuf.StringValue
	1, // 1: com.coralogixapis.dashboards.v1.common.LabelledValue.labels:type_name -> com.coralogixapis.dashboards.v1.common.LabelledValue.LabelsEntry
	3, // 2: com.coralogixapis.dashboards.v1.common.LabelledValue.value:type_name -> google.protobuf.DoubleValue
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_init() }
func file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_init() {
	if File_com_coralogixapis_dashboards_v1_common_labelled_value_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelledValue); i {
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
			RawDescriptor: file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_dashboards_v1_common_labelled_value_proto = out.File
	file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_rawDesc = nil
	file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_goTypes = nil
	file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_depIdxs = nil
}
