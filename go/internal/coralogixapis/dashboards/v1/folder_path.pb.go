// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: com/coralogixapis/dashboards/v1/ast/folder_path.proto

package v1

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

type FolderPath struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Segments      []string               `protobuf:"bytes,1,rep,name=segments,proto3" json:"segments,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FolderPath) Reset() {
	*x = FolderPath{}
	mi := &file_com_coralogixapis_dashboards_v1_ast_folder_path_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FolderPath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FolderPath) ProtoMessage() {}

func (x *FolderPath) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_ast_folder_path_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FolderPath.ProtoReflect.Descriptor instead.
func (*FolderPath) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_folder_path_proto_rawDescGZIP(), []int{0}
}

func (x *FolderPath) GetSegments() []string {
	if x != nil {
		return x.Segments
	}
	return nil
}

var File_com_coralogixapis_dashboards_v1_ast_folder_path_proto protoreflect.FileDescriptor

var file_com_coralogixapis_dashboards_v1_ast_folder_path_proto_rawDesc = []byte{
	0x0a, 0x35, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x73, 0x74, 0x2f, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x74, 0x22, 0x28, 0x0a, 0x0a,
	0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65,
	0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65,
	0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_dashboards_v1_ast_folder_path_proto_rawDescOnce sync.Once
	file_com_coralogixapis_dashboards_v1_ast_folder_path_proto_rawDescData = file_com_coralogixapis_dashboards_v1_ast_folder_path_proto_rawDesc
)

func file_com_coralogixapis_dashboards_v1_ast_folder_path_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_dashboards_v1_ast_folder_path_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_dashboards_v1_ast_folder_path_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_dashboards_v1_ast_folder_path_proto_rawDescData)
	})
	return file_com_coralogixapis_dashboards_v1_ast_folder_path_proto_rawDescData
}

var file_com_coralogixapis_dashboards_v1_ast_folder_path_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_dashboards_v1_ast_folder_path_proto_goTypes = []any{
	(*FolderPath)(nil), // 0: com.coralogixapis.dashboards.v1.ast.FolderPath
}
var file_com_coralogixapis_dashboards_v1_ast_folder_path_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_dashboards_v1_ast_folder_path_proto_init() }
func file_com_coralogixapis_dashboards_v1_ast_folder_path_proto_init() {
	if File_com_coralogixapis_dashboards_v1_ast_folder_path_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_dashboards_v1_ast_folder_path_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_dashboards_v1_ast_folder_path_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_dashboards_v1_ast_folder_path_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_dashboards_v1_ast_folder_path_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_dashboards_v1_ast_folder_path_proto = out.File
	file_com_coralogixapis_dashboards_v1_ast_folder_path_proto_rawDesc = nil
	file_com_coralogixapis_dashboards_v1_ast_folder_path_proto_goTypes = nil
	file_com_coralogixapis_dashboards_v1_ast_folder_path_proto_depIdxs = nil
}
