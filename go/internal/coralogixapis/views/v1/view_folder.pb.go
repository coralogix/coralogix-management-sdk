// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/views/v1/view_folder.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

// View folder
type ViewFolder struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Id            *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ViewFolder) Reset() {
	*x = ViewFolder{}
	mi := &file_com_coralogixapis_views_v1_view_folder_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ViewFolder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewFolder) ProtoMessage() {}

func (x *ViewFolder) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_views_v1_view_folder_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewFolder.ProtoReflect.Descriptor instead.
func (*ViewFolder) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_views_v1_view_folder_proto_rawDescGZIP(), []int{0}
}

func (x *ViewFolder) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ViewFolder) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

var File_com_coralogixapis_views_v1_view_folder_proto protoreflect.FileDescriptor

const file_com_coralogixapis_views_v1_view_folder_proto_rawDesc = "" +
	"\n" +
	",com/coralogixapis/views/v1/view_folder.proto\x12\x1acom.coralogixapis.views.v1\x1a\x1egoogle/protobuf/wrappers.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\xd2\x02\n" +
	"\n" +
	"ViewFolder\x12\xc7\x01\n" +
	"\x02id\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueB\x98\x01\x92A\x94\x012\x1dUnique identifier for foldersJ&\"3dc02998-0b50-4ea8-b68a-4779d716fa1f\"x$\x80\x01$\x8a\x01>^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$\xa2\x02\x04uuidR\x02id\x12R\n" +
	"\x04name\x18\x02 \x01(\v2\x1c.google.protobuf.StringValueB \x92A\x1d2\vFolder nameJ\v\"My Folder\"\x80\x01\x01R\x04name:&\x92A#\n" +
	"!*\n" +
	"ViewFolder2\fView folder.\xd2\x01\x04nameb\x06proto3"

var (
	file_com_coralogixapis_views_v1_view_folder_proto_rawDescOnce sync.Once
	file_com_coralogixapis_views_v1_view_folder_proto_rawDescData []byte
)

func file_com_coralogixapis_views_v1_view_folder_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_views_v1_view_folder_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_views_v1_view_folder_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_views_v1_view_folder_proto_rawDesc), len(file_com_coralogixapis_views_v1_view_folder_proto_rawDesc)))
	})
	return file_com_coralogixapis_views_v1_view_folder_proto_rawDescData
}

var file_com_coralogixapis_views_v1_view_folder_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_views_v1_view_folder_proto_goTypes = []any{
	(*ViewFolder)(nil),             // 0: com.coralogixapis.views.v1.ViewFolder
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
}
var file_com_coralogixapis_views_v1_view_folder_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.views.v1.ViewFolder.id:type_name -> google.protobuf.StringValue
	1, // 1: com.coralogixapis.views.v1.ViewFolder.name:type_name -> google.protobuf.StringValue
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_views_v1_view_folder_proto_init() }
func file_com_coralogixapis_views_v1_view_folder_proto_init() {
	if File_com_coralogixapis_views_v1_view_folder_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_views_v1_view_folder_proto_rawDesc), len(file_com_coralogixapis_views_v1_view_folder_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_views_v1_view_folder_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_views_v1_view_folder_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_views_v1_view_folder_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_views_v1_view_folder_proto = out.File
	file_com_coralogixapis_views_v1_view_folder_proto_goTypes = nil
	file_com_coralogixapis_views_v1_view_folder_proto_depIdxs = nil
}
