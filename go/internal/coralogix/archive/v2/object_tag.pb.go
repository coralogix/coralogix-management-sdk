// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: com/coralogix/archive/v2/object_tag.proto

package v2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ObjectTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ObjectTag) Reset() {
	*x = ObjectTag{}
	mi := &file_com_coralogix_archive_v2_object_tag_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ObjectTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectTag) ProtoMessage() {}

func (x *ObjectTag) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_object_tag_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectTag.ProtoReflect.Descriptor instead.
func (*ObjectTag) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_object_tag_proto_rawDescGZIP(), []int{0}
}

func (x *ObjectTag) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ObjectTag) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_com_coralogix_archive_v2_object_tag_proto protoreflect.FileDescriptor

var file_com_coralogix_archive_v2_object_tag_proto_rawDesc = []byte{
	0x0a, 0x29, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x2e, 0x76, 0x32, 0x22, 0x33, 0x0a, 0x09, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54,
	0x61, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_com_coralogix_archive_v2_object_tag_proto_rawDescOnce sync.Once
	file_com_coralogix_archive_v2_object_tag_proto_rawDescData = file_com_coralogix_archive_v2_object_tag_proto_rawDesc
)

func file_com_coralogix_archive_v2_object_tag_proto_rawDescGZIP() []byte {
	file_com_coralogix_archive_v2_object_tag_proto_rawDescOnce.Do(func() {
		file_com_coralogix_archive_v2_object_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_archive_v2_object_tag_proto_rawDescData)
	})
	return file_com_coralogix_archive_v2_object_tag_proto_rawDescData
}

var file_com_coralogix_archive_v2_object_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogix_archive_v2_object_tag_proto_goTypes = []any{
	(*ObjectTag)(nil), // 0: com.coralogix.archive.v2.ObjectTag
}
var file_com_coralogix_archive_v2_object_tag_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogix_archive_v2_object_tag_proto_init() }
func file_com_coralogix_archive_v2_object_tag_proto_init() {
	if File_com_coralogix_archive_v2_object_tag_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_archive_v2_object_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_archive_v2_object_tag_proto_goTypes,
		DependencyIndexes: file_com_coralogix_archive_v2_object_tag_proto_depIdxs,
		MessageInfos:      file_com_coralogix_archive_v2_object_tag_proto_msgTypes,
	}.Build()
	File_com_coralogix_archive_v2_object_tag_proto = out.File
	file_com_coralogix_archive_v2_object_tag_proto_rawDesc = nil
	file_com_coralogix_archive_v2_object_tag_proto_goTypes = nil
	file_com_coralogix_archive_v2_object_tag_proto_depIdxs = nil
}
