// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/dashboards/v1/common/query.proto

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

type PromQLQueryType int32

const (
	PromQLQueryType_PROM_QL_QUERY_TYPE_UNSPECIFIED PromQLQueryType = 0
	PromQLQueryType_PROM_QL_QUERY_TYPE_RANGE       PromQLQueryType = 1
	PromQLQueryType_PROM_QL_QUERY_TYPE_INSTANT     PromQLQueryType = 2
)

// Enum value maps for PromQLQueryType.
var (
	PromQLQueryType_name = map[int32]string{
		0: "PROM_QL_QUERY_TYPE_UNSPECIFIED",
		1: "PROM_QL_QUERY_TYPE_RANGE",
		2: "PROM_QL_QUERY_TYPE_INSTANT",
	}
	PromQLQueryType_value = map[string]int32{
		"PROM_QL_QUERY_TYPE_UNSPECIFIED": 0,
		"PROM_QL_QUERY_TYPE_RANGE":       1,
		"PROM_QL_QUERY_TYPE_INSTANT":     2,
	}
)

func (x PromQLQueryType) Enum() *PromQLQueryType {
	p := new(PromQLQueryType)
	*p = x
	return p
}

func (x PromQLQueryType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PromQLQueryType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_dashboards_v1_common_query_proto_enumTypes[0].Descriptor()
}

func (PromQLQueryType) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_dashboards_v1_common_query_proto_enumTypes[0]
}

func (x PromQLQueryType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PromQLQueryType.Descriptor instead.
func (PromQLQueryType) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_query_proto_rawDescGZIP(), []int{0}
}

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

const file_com_coralogixapis_dashboards_v1_common_query_proto_rawDesc = "" +
	"\n" +
	"2com/coralogixapis/dashboards/v1/common/query.proto\x12&com.coralogixapis.dashboards.v1.common\x1a\x1egoogle/protobuf/wrappers.proto\"$\n" +
	"\x0eDataprimeQuery\x12\x12\n" +
	"\x04text\x18\x01 \x01(\tR\x04text\".\n" +
	"\x18SerializedDataprimeQuery\x12\x12\n" +
	"\x04data\x18\x01 \x01(\fR\x04data\"\xc0\x01\n" +
	"\x12FullDataprimeQuery\x12`\n" +
	"\n" +
	"serialized\x18\x01 \x01(\v2@.com.coralogixapis.dashboards.v1.common.SerializedDataprimeQueryR\n" +
	"serialized\x12H\n" +
	"\x03raw\x18\x02 \x01(\v26.com.coralogixapis.dashboards.v1.common.DataprimeQueryR\x03raw\"A\n" +
	"\vPromQlQuery\x122\n" +
	"\x05value\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueR\x05value\"A\n" +
	"\vLuceneQuery\x122\n" +
	"\x05value\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueR\x05value*s\n" +
	"\x0fPromQLQueryType\x12\"\n" +
	"\x1ePROM_QL_QUERY_TYPE_UNSPECIFIED\x10\x00\x12\x1c\n" +
	"\x18PROM_QL_QUERY_TYPE_RANGE\x10\x01\x12\x1e\n" +
	"\x1aPROM_QL_QUERY_TYPE_INSTANT\x10\x02b\x06proto3"

var (
	file_com_coralogixapis_dashboards_v1_common_query_proto_rawDescOnce sync.Once
	file_com_coralogixapis_dashboards_v1_common_query_proto_rawDescData []byte
)

func file_com_coralogixapis_dashboards_v1_common_query_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_dashboards_v1_common_query_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_dashboards_v1_common_query_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_dashboards_v1_common_query_proto_rawDesc), len(file_com_coralogixapis_dashboards_v1_common_query_proto_rawDesc)))
	})
	return file_com_coralogixapis_dashboards_v1_common_query_proto_rawDescData
}

var file_com_coralogixapis_dashboards_v1_common_query_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_dashboards_v1_common_query_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_com_coralogixapis_dashboards_v1_common_query_proto_goTypes = []any{
	(PromQLQueryType)(0),             // 0: com.coralogixapis.dashboards.v1.common.PromQLQueryType
	(*DataprimeQuery)(nil),           // 1: com.coralogixapis.dashboards.v1.common.DataprimeQuery
	(*SerializedDataprimeQuery)(nil), // 2: com.coralogixapis.dashboards.v1.common.SerializedDataprimeQuery
	(*FullDataprimeQuery)(nil),       // 3: com.coralogixapis.dashboards.v1.common.FullDataprimeQuery
	(*PromQlQuery)(nil),              // 4: com.coralogixapis.dashboards.v1.common.PromQlQuery
	(*LuceneQuery)(nil),              // 5: com.coralogixapis.dashboards.v1.common.LuceneQuery
	(*wrapperspb.StringValue)(nil),   // 6: google.protobuf.StringValue
}
var file_com_coralogixapis_dashboards_v1_common_query_proto_depIdxs = []int32{
	2, // 0: com.coralogixapis.dashboards.v1.common.FullDataprimeQuery.serialized:type_name -> com.coralogixapis.dashboards.v1.common.SerializedDataprimeQuery
	1, // 1: com.coralogixapis.dashboards.v1.common.FullDataprimeQuery.raw:type_name -> com.coralogixapis.dashboards.v1.common.DataprimeQuery
	6, // 2: com.coralogixapis.dashboards.v1.common.PromQlQuery.value:type_name -> google.protobuf.StringValue
	6, // 3: com.coralogixapis.dashboards.v1.common.LuceneQuery.value:type_name -> google.protobuf.StringValue
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_dashboards_v1_common_query_proto_rawDesc), len(file_com_coralogixapis_dashboards_v1_common_query_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_dashboards_v1_common_query_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_dashboards_v1_common_query_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_dashboards_v1_common_query_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_dashboards_v1_common_query_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_dashboards_v1_common_query_proto = out.File
	file_com_coralogixapis_dashboards_v1_common_query_proto_goTypes = nil
	file_com_coralogixapis_dashboards_v1_common_query_proto_depIdxs = nil
}
