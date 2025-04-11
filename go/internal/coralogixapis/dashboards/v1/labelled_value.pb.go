// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/dashboards/v1/common/labelled_value.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

type LabelledValue struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Name          *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Labels        map[string]string       `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Value         *wrapperspb.DoubleValue `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LabelledValue) Reset() {
	*x = LabelledValue{}
	mi := &file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LabelledValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelledValue) ProtoMessage() {}

func (x *LabelledValue) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_msgTypes[0]
	if x != nil {
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

const file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_rawDesc = "" +
	"\n" +
	";com/coralogixapis/dashboards/v1/common/labelled_value.proto\x12&com.coralogixapis.dashboards.v1.common\x1a\x1egoogle/protobuf/wrappers.proto\"\x8b\x02\n" +
	"\rLabelledValue\x120\n" +
	"\x04name\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueR\x04name\x12Y\n" +
	"\x06labels\x18\x02 \x03(\v2A.com.coralogixapis.dashboards.v1.common.LabelledValue.LabelsEntryR\x06labels\x122\n" +
	"\x05value\x18\x03 \x01(\v2\x1c.google.protobuf.DoubleValueR\x05value\x1a9\n" +
	"\vLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01b\x06proto3"

var (
	file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_rawDescOnce sync.Once
	file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_rawDescData []byte
)

func file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_rawDesc), len(file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_rawDesc)))
	})
	return file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_rawDescData
}

var file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_goTypes = []any{
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_rawDesc), len(file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_rawDesc)),
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
	file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_goTypes = nil
	file_com_coralogixapis_dashboards_v1_common_labelled_value_proto_depIdxs = nil
}
