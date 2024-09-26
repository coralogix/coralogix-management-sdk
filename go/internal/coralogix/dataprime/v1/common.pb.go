// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.28.1
// source: com/coralogix/dataprime/v1/common.proto

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

type Backend int32

const (
	Backend_BACKEND_ARCHIVE_UNSPECIFIED Backend = 0
	Backend_BACKEND_OPENSEARCH          Backend = 1
)

// Enum value maps for Backend.
var (
	Backend_name = map[int32]string{
		0: "BACKEND_ARCHIVE_UNSPECIFIED",
		1: "BACKEND_OPENSEARCH",
	}
	Backend_value = map[string]int32{
		"BACKEND_ARCHIVE_UNSPECIFIED": 0,
		"BACKEND_OPENSEARCH":          1,
	}
)

func (x Backend) Enum() *Backend {
	p := new(Backend)
	*p = x
	return p
}

func (x Backend) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Backend) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_dataprime_v1_common_proto_enumTypes[0].Descriptor()
}

func (Backend) Type() protoreflect.EnumType {
	return &file_com_coralogix_dataprime_v1_common_proto_enumTypes[0]
}

func (x Backend) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Backend.Descriptor instead.
func (Backend) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_common_proto_rawDescGZIP(), []int{0}
}

type UntypedKeypath_Root int32

const (
	UntypedKeypath_ROOT_EVENT_LABELS_UNSPECIFIED UntypedKeypath_Root = 0
	UntypedKeypath_ROOT_EVENT_METADATA           UntypedKeypath_Root = 1
	UntypedKeypath_ROOT_USER_DATA                UntypedKeypath_Root = 2
)

// Enum value maps for UntypedKeypath_Root.
var (
	UntypedKeypath_Root_name = map[int32]string{
		0: "ROOT_EVENT_LABELS_UNSPECIFIED",
		1: "ROOT_EVENT_METADATA",
		2: "ROOT_USER_DATA",
	}
	UntypedKeypath_Root_value = map[string]int32{
		"ROOT_EVENT_LABELS_UNSPECIFIED": 0,
		"ROOT_EVENT_METADATA":           1,
		"ROOT_USER_DATA":                2,
	}
)

func (x UntypedKeypath_Root) Enum() *UntypedKeypath_Root {
	p := new(UntypedKeypath_Root)
	*p = x
	return p
}

func (x UntypedKeypath_Root) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UntypedKeypath_Root) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_dataprime_v1_common_proto_enumTypes[1].Descriptor()
}

func (UntypedKeypath_Root) Type() protoreflect.EnumType {
	return &file_com_coralogix_dataprime_v1_common_proto_enumTypes[1]
}

func (x UntypedKeypath_Root) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UntypedKeypath_Root.Descriptor instead.
func (UntypedKeypath_Root) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_common_proto_rawDescGZIP(), []int{0, 0}
}

type UntypedKeypath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Root         UntypedKeypath_Root       `protobuf:"varint,1,opt,name=root,proto3,enum=com.coralogix.dataprime.v1.UntypedKeypath_Root" json:"root,omitempty"`
	PathElements []*wrapperspb.StringValue `protobuf:"bytes,2,rep,name=path_elements,json=pathElements,proto3" json:"path_elements,omitempty"`
}

func (x *UntypedKeypath) Reset() {
	*x = UntypedKeypath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_dataprime_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UntypedKeypath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UntypedKeypath) ProtoMessage() {}

func (x *UntypedKeypath) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UntypedKeypath.ProtoReflect.Descriptor instead.
func (*UntypedKeypath) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *UntypedKeypath) GetRoot() UntypedKeypath_Root {
	if x != nil {
		return x.Root
	}
	return UntypedKeypath_ROOT_EVENT_LABELS_UNSPECIFIED
}

func (x *UntypedKeypath) GetPathElements() []*wrapperspb.StringValue {
	if x != nil {
		return x.PathElements
	}
	return nil
}

type KeyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KeyValue) Reset() {
	*x = KeyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_dataprime_v1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValue) ProtoMessage() {}

func (x *KeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValue.ProtoReflect.Descriptor instead.
func (*KeyValue) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *KeyValue) GetKey() *wrapperspb.StringValue {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *KeyValue) GetValue() *wrapperspb.StringValue {
	if x != nil {
		return x.Value
	}
	return nil
}

type Interval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To   *wrapperspb.Int64Value `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *Interval) Reset() {
	*x = Interval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_dataprime_v1_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Interval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Interval) ProtoMessage() {}

func (x *Interval) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Interval.ProtoReflect.Descriptor instead.
func (*Interval) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_common_proto_rawDescGZIP(), []int{2}
}

func (x *Interval) GetFrom() *wrapperspb.Int64Value {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *Interval) GetTo() *wrapperspb.Int64Value {
	if x != nil {
		return x.To
	}
	return nil
}

var File_com_coralogix_dataprime_v1_common_proto protoreflect.FileDescriptor

var file_com_coralogix_dataprime_v1_common_proto_rawDesc = []byte{
	0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69,
	0x6d, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x01, 0x0a, 0x0e, 0x55, 0x6e, 0x74, 0x79, 0x70, 0x65,
	0x64, 0x4b, 0x65, 0x79, 0x70, 0x61, 0x74, 0x68, 0x12, 0x43, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x74, 0x79, 0x70, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x70, 0x61,
	0x74, 0x68, 0x2e, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x12, 0x41, 0x0a,
	0x0d, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0c, 0x70, 0x61, 0x74, 0x68, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x22, 0x56, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x21, 0x0a, 0x1d, 0x52, 0x4f, 0x4f, 0x54,
	0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x53, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x52,
	0x4f, 0x4f, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x44, 0x41,
	0x54, 0x41, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x4f, 0x4f, 0x54, 0x5f, 0x55, 0x53, 0x45,
	0x52, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0x02, 0x22, 0x6e, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x2e, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x68, 0x0a, 0x08, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x12, 0x2f, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x2b, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02,
	0x74, 0x6f, 0x2a, 0x42, 0x0a, 0x07, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x12, 0x1f, 0x0a,
	0x1b, 0x42, 0x41, 0x43, 0x4b, 0x45, 0x4e, 0x44, 0x5f, 0x41, 0x52, 0x43, 0x48, 0x49, 0x56, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16,
	0x0a, 0x12, 0x42, 0x41, 0x43, 0x4b, 0x45, 0x4e, 0x44, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x53, 0x45,
	0x41, 0x52, 0x43, 0x48, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_dataprime_v1_common_proto_rawDescOnce sync.Once
	file_com_coralogix_dataprime_v1_common_proto_rawDescData = file_com_coralogix_dataprime_v1_common_proto_rawDesc
)

func file_com_coralogix_dataprime_v1_common_proto_rawDescGZIP() []byte {
	file_com_coralogix_dataprime_v1_common_proto_rawDescOnce.Do(func() {
		file_com_coralogix_dataprime_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_dataprime_v1_common_proto_rawDescData)
	})
	return file_com_coralogix_dataprime_v1_common_proto_rawDescData
}

var file_com_coralogix_dataprime_v1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_com_coralogix_dataprime_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogix_dataprime_v1_common_proto_goTypes = []interface{}{
	(Backend)(0),                   // 0: com.coralogix.dataprime.v1.Backend
	(UntypedKeypath_Root)(0),       // 1: com.coralogix.dataprime.v1.UntypedKeypath.Root
	(*UntypedKeypath)(nil),         // 2: com.coralogix.dataprime.v1.UntypedKeypath
	(*KeyValue)(nil),               // 3: com.coralogix.dataprime.v1.KeyValue
	(*Interval)(nil),               // 4: com.coralogix.dataprime.v1.Interval
	(*wrapperspb.StringValue)(nil), // 5: google.protobuf.StringValue
	(*wrapperspb.Int64Value)(nil),  // 6: google.protobuf.Int64Value
}
var file_com_coralogix_dataprime_v1_common_proto_depIdxs = []int32{
	1, // 0: com.coralogix.dataprime.v1.UntypedKeypath.root:type_name -> com.coralogix.dataprime.v1.UntypedKeypath.Root
	5, // 1: com.coralogix.dataprime.v1.UntypedKeypath.path_elements:type_name -> google.protobuf.StringValue
	5, // 2: com.coralogix.dataprime.v1.KeyValue.key:type_name -> google.protobuf.StringValue
	5, // 3: com.coralogix.dataprime.v1.KeyValue.value:type_name -> google.protobuf.StringValue
	6, // 4: com.coralogix.dataprime.v1.Interval.from:type_name -> google.protobuf.Int64Value
	6, // 5: com.coralogix.dataprime.v1.Interval.to:type_name -> google.protobuf.Int64Value
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_com_coralogix_dataprime_v1_common_proto_init() }
func file_com_coralogix_dataprime_v1_common_proto_init() {
	if File_com_coralogix_dataprime_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_dataprime_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UntypedKeypath); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_com_coralogix_dataprime_v1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_com_coralogix_dataprime_v1_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Interval); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_dataprime_v1_common_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_dataprime_v1_common_proto_goTypes,
		DependencyIndexes: file_com_coralogix_dataprime_v1_common_proto_depIdxs,
		EnumInfos:         file_com_coralogix_dataprime_v1_common_proto_enumTypes,
		MessageInfos:      file_com_coralogix_dataprime_v1_common_proto_msgTypes,
	}.Build()
	File_com_coralogix_dataprime_v1_common_proto = out.File
	file_com_coralogix_dataprime_v1_common_proto_rawDesc = nil
	file_com_coralogix_dataprime_v1_common_proto_goTypes = nil
	file_com_coralogix_dataprime_v1_common_proto_depIdxs = nil
}
