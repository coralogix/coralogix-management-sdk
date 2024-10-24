// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: com/coralogix/dataprime/ast/v1/ast.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Ast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query *Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *Ast) Reset() {
	*x = Ast{}
	mi := &file_com_coralogix_dataprime_ast_v1_ast_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ast) ProtoMessage() {}

func (x *Ast) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_ast_v1_ast_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ast.ProtoReflect.Descriptor instead.
func (*Ast) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_ast_v1_ast_proto_rawDescGZIP(), []int{0}
}

func (x *Ast) GetQuery() *Query {
	if x != nil {
		return x.Query
	}
	return nil
}

type DdlAst struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ast         *Ast   `protobuf:"bytes,1,opt,name=ast,proto3" json:"ast,omitempty"`
	DatasetName string `protobuf:"bytes,2,opt,name=dataset_name,json=datasetName,proto3" json:"dataset_name,omitempty"`
}

func (x *DdlAst) Reset() {
	*x = DdlAst{}
	mi := &file_com_coralogix_dataprime_ast_v1_ast_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DdlAst) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DdlAst) ProtoMessage() {}

func (x *DdlAst) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_ast_v1_ast_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DdlAst.ProtoReflect.Descriptor instead.
func (*DdlAst) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_ast_v1_ast_proto_rawDescGZIP(), []int{1}
}

func (x *DdlAst) GetAst() *Ast {
	if x != nil {
		return x.Ast
	}
	return nil
}

func (x *DdlAst) GetDatasetName() string {
	if x != nil {
		return x.DatasetName
	}
	return ""
}

type TimeRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *TimeRange) Reset() {
	*x = TimeRange{}
	mi := &file_com_coralogix_dataprime_ast_v1_ast_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimeRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeRange) ProtoMessage() {}

func (x *TimeRange) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_ast_v1_ast_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeRange.ProtoReflect.Descriptor instead.
func (*TimeRange) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_ast_v1_ast_proto_rawDescGZIP(), []int{2}
}

func (x *TimeRange) GetFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *TimeRange) GetTo() *timestamppb.Timestamp {
	if x != nil {
		return x.To
	}
	return nil
}

var File_com_coralogix_dataprime_ast_v1_ast_proto protoreflect.FileDescriptor

var file_com_coralogix_dataprime_ast_v1_ast_proto_rawDesc = []byte{
	0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x61, 0x73, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x69, 0x6d, 0x65, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70,
	0x72, 0x69, 0x6d, 0x65, 0x2f, 0x61, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x03, 0x41, 0x73, 0x74, 0x12, 0x3b,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x4a, 0x04, 0x08, 0x02, 0x10,
	0x03, 0x22, 0x62, 0x0a, 0x06, 0x44, 0x64, 0x6c, 0x41, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x03, 0x61,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69,
	0x6d, 0x65, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x74, 0x52, 0x03, 0x61,
	0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x67, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_dataprime_ast_v1_ast_proto_rawDescOnce sync.Once
	file_com_coralogix_dataprime_ast_v1_ast_proto_rawDescData = file_com_coralogix_dataprime_ast_v1_ast_proto_rawDesc
)

func file_com_coralogix_dataprime_ast_v1_ast_proto_rawDescGZIP() []byte {
	file_com_coralogix_dataprime_ast_v1_ast_proto_rawDescOnce.Do(func() {
		file_com_coralogix_dataprime_ast_v1_ast_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_dataprime_ast_v1_ast_proto_rawDescData)
	})
	return file_com_coralogix_dataprime_ast_v1_ast_proto_rawDescData
}

var file_com_coralogix_dataprime_ast_v1_ast_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogix_dataprime_ast_v1_ast_proto_goTypes = []any{
	(*Ast)(nil),                   // 0: com.coralogix.dataprime.ast.v1.Ast
	(*DdlAst)(nil),                // 1: com.coralogix.dataprime.ast.v1.DdlAst
	(*TimeRange)(nil),             // 2: com.coralogix.dataprime.ast.v1.TimeRange
	(*Query)(nil),                 // 3: com.coralogix.dataprime.ast.v1.Query
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_com_coralogix_dataprime_ast_v1_ast_proto_depIdxs = []int32{
	3, // 0: com.coralogix.dataprime.ast.v1.Ast.query:type_name -> com.coralogix.dataprime.ast.v1.Query
	0, // 1: com.coralogix.dataprime.ast.v1.DdlAst.ast:type_name -> com.coralogix.dataprime.ast.v1.Ast
	4, // 2: com.coralogix.dataprime.ast.v1.TimeRange.from:type_name -> google.protobuf.Timestamp
	4, // 3: com.coralogix.dataprime.ast.v1.TimeRange.to:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_com_coralogix_dataprime_ast_v1_ast_proto_init() }
func file_com_coralogix_dataprime_ast_v1_ast_proto_init() {
	if File_com_coralogix_dataprime_ast_v1_ast_proto != nil {
		return
	}
	file_com_coralogix_dataprime_ast_v1_query_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_dataprime_ast_v1_ast_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_dataprime_ast_v1_ast_proto_goTypes,
		DependencyIndexes: file_com_coralogix_dataprime_ast_v1_ast_proto_depIdxs,
		MessageInfos:      file_com_coralogix_dataprime_ast_v1_ast_proto_msgTypes,
	}.Build()
	File_com_coralogix_dataprime_ast_v1_ast_proto = out.File
	file_com_coralogix_dataprime_ast_v1_ast_proto_rawDesc = nil
	file_com_coralogix_dataprime_ast_v1_ast_proto_goTypes = nil
	file_com_coralogix_dataprime_ast_v1_ast_proto_depIdxs = nil
}
