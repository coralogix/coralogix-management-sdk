// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.26.1
// source: com/coralogixapis/dashboards/v1/common/query.proto

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

type DataprimeQuery struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Text          string                 `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DataprimeQuery) Reset() {
	*x = DataprimeQuery{}
	mi := &file_com_coralogixapis_dashboards_v1_common_query_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataprimeQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataprimeQuery) ProtoMessage() {}

func (x *DataprimeQuery) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_common_query_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataprimeQuery.ProtoReflect.Descriptor instead.
func (*DataprimeQuery) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_query_proto_rawDescGZIP(), []int{0}
}

func (x *DataprimeQuery) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type SerializedDataprimeQuery struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          []byte                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SerializedDataprimeQuery) Reset() {
	*x = SerializedDataprimeQuery{}
	mi := &file_com_coralogixapis_dashboards_v1_common_query_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SerializedDataprimeQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SerializedDataprimeQuery) ProtoMessage() {}

func (x *SerializedDataprimeQuery) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_common_query_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SerializedDataprimeQuery.ProtoReflect.Descriptor instead.
func (*SerializedDataprimeQuery) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_query_proto_rawDescGZIP(), []int{1}
}

func (x *SerializedDataprimeQuery) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type FullDataprimeQuery struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Serialized    *SerializedDataprimeQuery `protobuf:"bytes,1,opt,name=serialized,proto3" json:"serialized,omitempty"`
	Raw           *DataprimeQuery           `protobuf:"bytes,2,opt,name=raw,proto3" json:"raw,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FullDataprimeQuery) Reset() {
	*x = FullDataprimeQuery{}
	mi := &file_com_coralogixapis_dashboards_v1_common_query_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FullDataprimeQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FullDataprimeQuery) ProtoMessage() {}

func (x *FullDataprimeQuery) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_common_query_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FullDataprimeQuery.ProtoReflect.Descriptor instead.
func (*FullDataprimeQuery) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_query_proto_rawDescGZIP(), []int{2}
}

func (x *FullDataprimeQuery) GetSerialized() *SerializedDataprimeQuery {
	if x != nil {
		return x.Serialized
	}
	return nil
}

func (x *FullDataprimeQuery) GetRaw() *DataprimeQuery {
	if x != nil {
		return x.Raw
	}
	return nil
}

type PromQlQuery struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Value         *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PromQlQuery) Reset() {
	*x = PromQlQuery{}
	mi := &file_com_coralogixapis_dashboards_v1_common_query_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PromQlQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromQlQuery) ProtoMessage() {}

func (x *PromQlQuery) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_common_query_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromQlQuery.ProtoReflect.Descriptor instead.
func (*PromQlQuery) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_query_proto_rawDescGZIP(), []int{3}
}

func (x *PromQlQuery) GetValue() *wrapperspb.StringValue {
	if x != nil {
		return x.Value
	}
	return nil
}

type LuceneQuery struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Value         *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LuceneQuery) Reset() {
	*x = LuceneQuery{}
	mi := &file_com_coralogixapis_dashboards_v1_common_query_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LuceneQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LuceneQuery) ProtoMessage() {}

func (x *LuceneQuery) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_common_query_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LuceneQuery.ProtoReflect.Descriptor instead.
func (*LuceneQuery) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_query_proto_rawDescGZIP(), []int{4}
}

func (x *LuceneQuery) GetValue() *wrapperspb.StringValue {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_com_coralogixapis_dashboards_v1_common_query_proto protoreflect.FileDescriptor

var file_com_coralogixapis_dashboards_v1_common_query_proto_rawDesc = []byte{
	0x0a, 0x32, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x0e,
	0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x22, 0x2e, 0x0a, 0x18, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0xc0, 0x01, 0x0a, 0x12, 0x46, 0x75, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x70,
	0x72, 0x69, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x60, 0x0a, 0x0a, 0x73, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x64, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x0a, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x48, 0x0a, 0x03, 0x72,
	0x61, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x03, 0x72, 0x61, 0x77, 0x22, 0x41, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x6d, 0x51, 0x6c, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x41, 0x0a, 0x0b, 0x4c, 0x75, 0x63, 0x65,
	0x6e, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_dashboards_v1_common_query_proto_rawDescOnce sync.Once
	file_com_coralogixapis_dashboards_v1_common_query_proto_rawDescData = file_com_coralogixapis_dashboards_v1_common_query_proto_rawDesc
)

func file_com_coralogixapis_dashboards_v1_common_query_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_dashboards_v1_common_query_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_dashboards_v1_common_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_dashboards_v1_common_query_proto_rawDescData)
	})
	return file_com_coralogixapis_dashboards_v1_common_query_proto_rawDescData
}

var file_com_coralogixapis_dashboards_v1_common_query_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_com_coralogixapis_dashboards_v1_common_query_proto_goTypes = []any{
	(*DataprimeQuery)(nil),           // 0: com.coralogixapis.dashboards.v1.common.DataprimeQuery
	(*SerializedDataprimeQuery)(nil), // 1: com.coralogixapis.dashboards.v1.common.SerializedDataprimeQuery
	(*FullDataprimeQuery)(nil),       // 2: com.coralogixapis.dashboards.v1.common.FullDataprimeQuery
	(*PromQlQuery)(nil),              // 3: com.coralogixapis.dashboards.v1.common.PromQlQuery
	(*LuceneQuery)(nil),              // 4: com.coralogixapis.dashboards.v1.common.LuceneQuery
	(*wrapperspb.StringValue)(nil),   // 5: google.protobuf.StringValue
}
var file_com_coralogixapis_dashboards_v1_common_query_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.dashboards.v1.common.FullDataprimeQuery.serialized:type_name -> com.coralogixapis.dashboards.v1.common.SerializedDataprimeQuery
	0, // 1: com.coralogixapis.dashboards.v1.common.FullDataprimeQuery.raw:type_name -> com.coralogixapis.dashboards.v1.common.DataprimeQuery
	5, // 2: com.coralogixapis.dashboards.v1.common.PromQlQuery.value:type_name -> google.protobuf.StringValue
	5, // 3: com.coralogixapis.dashboards.v1.common.LuceneQuery.value:type_name -> google.protobuf.StringValue
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_dashboards_v1_common_query_proto_init() }
func file_com_coralogixapis_dashboards_v1_common_query_proto_init() {
	if File_com_coralogixapis_dashboards_v1_common_query_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_dashboards_v1_common_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_dashboards_v1_common_query_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_dashboards_v1_common_query_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_dashboards_v1_common_query_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_dashboards_v1_common_query_proto = out.File
	file_com_coralogixapis_dashboards_v1_common_query_proto_rawDesc = nil
	file_com_coralogixapis_dashboards_v1_common_query_proto_goTypes = nil
	file_com_coralogixapis_dashboards_v1_common_query_proto_depIdxs = nil
}
