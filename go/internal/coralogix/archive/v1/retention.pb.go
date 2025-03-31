// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogix/archive/v1/retention.proto

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

type Retention struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Id            *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Order         *wrapperspb.Int32Value  `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
	Name          *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Editable      *wrapperspb.BoolValue   `protobuf:"bytes,4,opt,name=editable,proto3" json:"editable,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Retention) Reset() {
	*x = Retention{}
	mi := &file_com_coralogix_archive_v1_retention_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Retention) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Retention) ProtoMessage() {}

func (x *Retention) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v1_retention_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Retention.ProtoReflect.Descriptor instead.
func (*Retention) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v1_retention_proto_rawDescGZIP(), []int{0}
}

func (x *Retention) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Retention) GetOrder() *wrapperspb.Int32Value {
	if x != nil {
		return x.Order
	}
	return nil
}

func (x *Retention) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *Retention) GetEditable() *wrapperspb.BoolValue {
	if x != nil {
		return x.Editable
	}
	return nil
}

var File_com_coralogix_archive_v1_retention_proto protoreflect.FileDescriptor

const file_com_coralogix_archive_v1_retention_proto_rawDesc = "" +
	"\n" +
	"(com/coralogix/archive/v1/retention.proto\x12\x18com.coralogix.archive.v1\x1a\x1egoogle/protobuf/wrappers.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\x87\x03\n" +
	"\tRetention\x12,\n" +
	"\x02id\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueR\x02id\x121\n" +
	"\x05order\x18\x02 \x01(\v2\x1b.google.protobuf.Int32ValueR\x05order\x120\n" +
	"\x04name\x18\x03 \x01(\v2\x1c.google.protobuf.StringValueR\x04name\x126\n" +
	"\beditable\x18\x04 \x01(\v2\x1a.google.protobuf.BoolValueR\beditable:\xae\x01\x92A\xaa\x01\n" +
	"7*\tRetention2*This data structure represents a retention*o\n" +
	"\x1cFind out more about archives\x12Ohttps://coralogix.com/docs/user-guides/data-flow/s3-archive/connect-s3-archive/b\x06proto3"

var (
	file_com_coralogix_archive_v1_retention_proto_rawDescOnce sync.Once
	file_com_coralogix_archive_v1_retention_proto_rawDescData []byte
)

func file_com_coralogix_archive_v1_retention_proto_rawDescGZIP() []byte {
	file_com_coralogix_archive_v1_retention_proto_rawDescOnce.Do(func() {
		file_com_coralogix_archive_v1_retention_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogix_archive_v1_retention_proto_rawDesc), len(file_com_coralogix_archive_v1_retention_proto_rawDesc)))
	})
	return file_com_coralogix_archive_v1_retention_proto_rawDescData
}

var file_com_coralogix_archive_v1_retention_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogix_archive_v1_retention_proto_goTypes = []any{
	(*Retention)(nil),              // 0: com.coralogix.archive.v1.Retention
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
	(*wrapperspb.Int32Value)(nil),  // 2: google.protobuf.Int32Value
	(*wrapperspb.BoolValue)(nil),   // 3: google.protobuf.BoolValue
}
var file_com_coralogix_archive_v1_retention_proto_depIdxs = []int32{
	1, // 0: com.coralogix.archive.v1.Retention.id:type_name -> google.protobuf.StringValue
	2, // 1: com.coralogix.archive.v1.Retention.order:type_name -> google.protobuf.Int32Value
	1, // 2: com.coralogix.archive.v1.Retention.name:type_name -> google.protobuf.StringValue
	3, // 3: com.coralogix.archive.v1.Retention.editable:type_name -> google.protobuf.BoolValue
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_com_coralogix_archive_v1_retention_proto_init() }
func file_com_coralogix_archive_v1_retention_proto_init() {
	if File_com_coralogix_archive_v1_retention_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogix_archive_v1_retention_proto_rawDesc), len(file_com_coralogix_archive_v1_retention_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_archive_v1_retention_proto_goTypes,
		DependencyIndexes: file_com_coralogix_archive_v1_retention_proto_depIdxs,
		MessageInfos:      file_com_coralogix_archive_v1_retention_proto_msgTypes,
	}.Build()
	File_com_coralogix_archive_v1_retention_proto = out.File
	file_com_coralogix_archive_v1_retention_proto_goTypes = nil
	file_com_coralogix_archive_v1_retention_proto_depIdxs = nil
}
