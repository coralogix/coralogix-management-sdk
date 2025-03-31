// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogix/archive/dataset/v2/dataset.proto

package v2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Dataspace struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Dataspace     string                 `protobuf:"bytes,1,opt,name=dataspace,proto3" json:"dataspace,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Dataspace) Reset() {
	*x = Dataspace{}
	mi := &file_com_coralogix_archive_dataset_v2_dataset_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Dataspace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dataspace) ProtoMessage() {}

func (x *Dataspace) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_dataset_v2_dataset_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dataspace.ProtoReflect.Descriptor instead.
func (*Dataspace) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_dataset_proto_rawDescGZIP(), []int{0}
}

func (x *Dataspace) GetDataspace() string {
	if x != nil {
		return x.Dataspace
	}
	return ""
}

type Dataset struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Dataspace     *Dataspace             `protobuf:"bytes,1,opt,name=dataspace,proto3" json:"dataspace,omitempty"`
	Dataset       string                 `protobuf:"bytes,2,opt,name=dataset,proto3" json:"dataset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Dataset) Reset() {
	*x = Dataset{}
	mi := &file_com_coralogix_archive_dataset_v2_dataset_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Dataset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dataset) ProtoMessage() {}

func (x *Dataset) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_dataset_v2_dataset_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dataset.ProtoReflect.Descriptor instead.
func (*Dataset) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_dataset_proto_rawDescGZIP(), []int{1}
}

func (x *Dataset) GetDataspace() *Dataspace {
	if x != nil {
		return x.Dataspace
	}
	return nil
}

func (x *Dataset) GetDataset() string {
	if x != nil {
		return x.Dataset
	}
	return ""
}

var File_com_coralogix_archive_dataset_v2_dataset_proto protoreflect.FileDescriptor

const file_com_coralogix_archive_dataset_v2_dataset_proto_rawDesc = "" +
	"\n" +
	".com/coralogix/archive/dataset/v2/dataset.proto\x12 com.coralogix.archive.dataset.v2\")\n" +
	"\tDataspace\x12\x1c\n" +
	"\tdataspace\x18\x01 \x01(\tR\tdataspace\"n\n" +
	"\aDataset\x12I\n" +
	"\tdataspace\x18\x01 \x01(\v2+.com.coralogix.archive.dataset.v2.DataspaceR\tdataspace\x12\x18\n" +
	"\adataset\x18\x02 \x01(\tR\adatasetb\x06proto3"

var (
	file_com_coralogix_archive_dataset_v2_dataset_proto_rawDescOnce sync.Once
	file_com_coralogix_archive_dataset_v2_dataset_proto_rawDescData []byte
)

func file_com_coralogix_archive_dataset_v2_dataset_proto_rawDescGZIP() []byte {
	file_com_coralogix_archive_dataset_v2_dataset_proto_rawDescOnce.Do(func() {
		file_com_coralogix_archive_dataset_v2_dataset_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogix_archive_dataset_v2_dataset_proto_rawDesc), len(file_com_coralogix_archive_dataset_v2_dataset_proto_rawDesc)))
	})
	return file_com_coralogix_archive_dataset_v2_dataset_proto_rawDescData
}

var file_com_coralogix_archive_dataset_v2_dataset_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogix_archive_dataset_v2_dataset_proto_goTypes = []any{
	(*Dataspace)(nil), // 0: com.coralogix.archive.dataset.v2.Dataspace
	(*Dataset)(nil),   // 1: com.coralogix.archive.dataset.v2.Dataset
}
var file_com_coralogix_archive_dataset_v2_dataset_proto_depIdxs = []int32{
	0, // 0: com.coralogix.archive.dataset.v2.Dataset.dataspace:type_name -> com.coralogix.archive.dataset.v2.Dataspace
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_com_coralogix_archive_dataset_v2_dataset_proto_init() }
func file_com_coralogix_archive_dataset_v2_dataset_proto_init() {
	if File_com_coralogix_archive_dataset_v2_dataset_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogix_archive_dataset_v2_dataset_proto_rawDesc), len(file_com_coralogix_archive_dataset_v2_dataset_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_archive_dataset_v2_dataset_proto_goTypes,
		DependencyIndexes: file_com_coralogix_archive_dataset_v2_dataset_proto_depIdxs,
		MessageInfos:      file_com_coralogix_archive_dataset_v2_dataset_proto_msgTypes,
	}.Build()
	File_com_coralogix_archive_dataset_v2_dataset_proto = out.File
	file_com_coralogix_archive_dataset_v2_dataset_proto_goTypes = nil
	file_com_coralogix_archive_dataset_v2_dataset_proto_depIdxs = nil
}
