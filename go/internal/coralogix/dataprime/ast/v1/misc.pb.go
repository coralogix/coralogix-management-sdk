// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: com/coralogix/dataprime/ast/v1/misc.proto

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

type KeypathType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path     *Expression_Keypath `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Datatype *Datatype           `protobuf:"bytes,2,opt,name=datatype,proto3" json:"datatype,omitempty"`
}

func (x *KeypathType) Reset() {
	*x = KeypathType{}
	mi := &file_com_coralogix_dataprime_ast_v1_misc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KeypathType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeypathType) ProtoMessage() {}

func (x *KeypathType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_ast_v1_misc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeypathType.ProtoReflect.Descriptor instead.
func (*KeypathType) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_ast_v1_misc_proto_rawDescGZIP(), []int{0}
}

func (x *KeypathType) GetPath() *Expression_Keypath {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *KeypathType) GetDatatype() *Datatype {
	if x != nil {
		return x.Datatype
	}
	return nil
}

var File_com_coralogix_dataprime_ast_v1_misc_proto protoreflect.FileDescriptor

var file_com_coralogix_dataprime_ast_v1_misc_proto_rawDesc = []byte{
	0x0a, 0x29, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x61, 0x73, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x69, 0x73, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70,
	0x72, 0x69, 0x6d, 0x65, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x2f, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70,
	0x72, 0x69, 0x6d, 0x65, 0x2f, 0x61, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x61, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x01, 0x0a, 0x0b, 0x4b, 0x65, 0x79,
	0x70, 0x61, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x46, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65,
	0x2e, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x4b, 0x65, 0x79, 0x70, 0x61, 0x74, 0x68, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x12, 0x44, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x61, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x52, 0x08, 0x64, 0x61,
	0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_dataprime_ast_v1_misc_proto_rawDescOnce sync.Once
	file_com_coralogix_dataprime_ast_v1_misc_proto_rawDescData = file_com_coralogix_dataprime_ast_v1_misc_proto_rawDesc
)

func file_com_coralogix_dataprime_ast_v1_misc_proto_rawDescGZIP() []byte {
	file_com_coralogix_dataprime_ast_v1_misc_proto_rawDescOnce.Do(func() {
		file_com_coralogix_dataprime_ast_v1_misc_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_dataprime_ast_v1_misc_proto_rawDescData)
	})
	return file_com_coralogix_dataprime_ast_v1_misc_proto_rawDescData
}

var file_com_coralogix_dataprime_ast_v1_misc_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogix_dataprime_ast_v1_misc_proto_goTypes = []any{
	(*KeypathType)(nil),        // 0: com.coralogix.dataprime.ast.v1.KeypathType
	(*Expression_Keypath)(nil), // 1: com.coralogix.dataprime.ast.v1.Expression.Keypath
	(*Datatype)(nil),           // 2: com.coralogix.dataprime.ast.v1.Datatype
}
var file_com_coralogix_dataprime_ast_v1_misc_proto_depIdxs = []int32{
	1, // 0: com.coralogix.dataprime.ast.v1.KeypathType.path:type_name -> com.coralogix.dataprime.ast.v1.Expression.Keypath
	2, // 1: com.coralogix.dataprime.ast.v1.KeypathType.datatype:type_name -> com.coralogix.dataprime.ast.v1.Datatype
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_com_coralogix_dataprime_ast_v1_misc_proto_init() }
func file_com_coralogix_dataprime_ast_v1_misc_proto_init() {
	if File_com_coralogix_dataprime_ast_v1_misc_proto != nil {
		return
	}
	file_com_coralogix_dataprime_ast_v1_expression_proto_init()
	file_com_coralogix_dataprime_ast_v1_types_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_dataprime_ast_v1_misc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_dataprime_ast_v1_misc_proto_goTypes,
		DependencyIndexes: file_com_coralogix_dataprime_ast_v1_misc_proto_depIdxs,
		MessageInfos:      file_com_coralogix_dataprime_ast_v1_misc_proto_msgTypes,
	}.Build()
	File_com_coralogix_dataprime_ast_v1_misc_proto = out.File
	file_com_coralogix_dataprime_ast_v1_misc_proto_rawDesc = nil
	file_com_coralogix_dataprime_ast_v1_misc_proto_goTypes = nil
	file_com_coralogix_dataprime_ast_v1_misc_proto_depIdxs = nil
}
