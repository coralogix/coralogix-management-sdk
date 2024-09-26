// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.28.1
// source: com/coralogix/archive/v2/format.proto

package v2

import (
	v1 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/archive/format/generic/v1"
	v11 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/archive/format/wide_parquet/v1"
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

type LogsAvro struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LogsAvro) Reset() {
	*x = LogsAvro{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_archive_v2_format_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsAvro) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsAvro) ProtoMessage() {}

func (x *LogsAvro) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_format_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsAvro.ProtoReflect.Descriptor instead.
func (*LogsAvro) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_format_proto_rawDescGZIP(), []int{0}
}

type SpansAvro struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SpansAvro) Reset() {
	*x = SpansAvro{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_archive_v2_format_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpansAvro) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpansAvro) ProtoMessage() {}

func (x *SpansAvro) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_format_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpansAvro.ProtoReflect.Descriptor instead.
func (*SpansAvro) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_format_proto_rawDescGZIP(), []int{1}
}

type Format struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Format:
	//
	//	*Format_LogsAvro
	//	*Format_SpansAvro
	//	*Format_GenericEventAvro
	//	*Format_WideParquet
	Format isFormat_Format `protobuf_oneof:"format"`
}

func (x *Format) Reset() {
	*x = Format{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_archive_v2_format_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Format) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Format) ProtoMessage() {}

func (x *Format) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v2_format_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Format.ProtoReflect.Descriptor instead.
func (*Format) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v2_format_proto_rawDescGZIP(), []int{2}
}

func (m *Format) GetFormat() isFormat_Format {
	if m != nil {
		return m.Format
	}
	return nil
}

func (x *Format) GetLogsAvro() *LogsAvro {
	if x, ok := x.GetFormat().(*Format_LogsAvro); ok {
		return x.LogsAvro
	}
	return nil
}

func (x *Format) GetSpansAvro() *SpansAvro {
	if x, ok := x.GetFormat().(*Format_SpansAvro); ok {
		return x.SpansAvro
	}
	return nil
}

func (x *Format) GetGenericEventAvro() *v1.GenericEventAvro {
	if x, ok := x.GetFormat().(*Format_GenericEventAvro); ok {
		return x.GenericEventAvro
	}
	return nil
}

func (x *Format) GetWideParquet() *v11.WideParquet {
	if x, ok := x.GetFormat().(*Format_WideParquet); ok {
		return x.WideParquet
	}
	return nil
}

type isFormat_Format interface {
	isFormat_Format()
}

type Format_LogsAvro struct {
	LogsAvro *LogsAvro `protobuf:"bytes,1,opt,name=logs_avro,json=logsAvro,proto3,oneof"`
}

type Format_SpansAvro struct {
	SpansAvro *SpansAvro `protobuf:"bytes,2,opt,name=spans_avro,json=spansAvro,proto3,oneof"`
}

type Format_GenericEventAvro struct {
	GenericEventAvro *v1.GenericEventAvro `protobuf:"bytes,3,opt,name=generic_event_avro,json=genericEventAvro,proto3,oneof"`
}

type Format_WideParquet struct {
	WideParquet *v11.WideParquet `protobuf:"bytes,4,opt,name=wide_parquet,json=wideParquet,proto3,oneof"`
}

func (*Format_LogsAvro) isFormat_Format() {}

func (*Format_SpansAvro) isFormat_Format() {}

func (*Format_GenericEventAvro) isFormat_Format() {}

func (*Format_WideParquet) isFormat_Format() {}

var File_com_coralogix_archive_v2_format_proto protoreflect.FileDescriptor

var file_com_coralogix_archive_v2_format_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76,
	0x32, 0x1a, 0x35, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3f, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2f, 0x77, 0x69, 0x64, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x71,
	0x75, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x69, 0x64, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x71,
	0x75, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0a, 0x0a, 0x08, 0x4c, 0x6f, 0x67,
	0x73, 0x41, 0x76, 0x72, 0x6f, 0x22, 0x0b, 0x0a, 0x09, 0x53, 0x70, 0x61, 0x6e, 0x73, 0x41, 0x76,
	0x72, 0x6f, 0x22, 0xe6, 0x02, 0x0a, 0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x41, 0x0a,
	0x09, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x61, 0x76, 0x72, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x6f, 0x67, 0x73,
	0x41, 0x76, 0x72, 0x6f, 0x48, 0x00, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x73, 0x41, 0x76, 0x72, 0x6f,
	0x12, 0x44, 0x0a, 0x0a, 0x73, 0x70, 0x61, 0x6e, 0x73, 0x5f, 0x61, 0x76, 0x72, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e,
	0x53, 0x70, 0x61, 0x6e, 0x73, 0x41, 0x76, 0x72, 0x6f, 0x48, 0x00, 0x52, 0x09, 0x73, 0x70, 0x61,
	0x6e, 0x73, 0x41, 0x76, 0x72, 0x6f, 0x12, 0x69, 0x0a, 0x12, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69,
	0x63, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x76, 0x72, 0x6f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x69, 0x63, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x76, 0x72, 0x6f, 0x48, 0x00, 0x52,
	0x10, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x76, 0x72,
	0x6f, 0x12, 0x5e, 0x0a, 0x0c, 0x77, 0x69, 0x64, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x71, 0x75, 0x65,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2e, 0x77, 0x69, 0x64, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x71,
	0x75, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x64, 0x65, 0x50, 0x61, 0x72, 0x71, 0x75,
	0x65, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x77, 0x69, 0x64, 0x65, 0x50, 0x61, 0x72, 0x71, 0x75, 0x65,
	0x74, 0x42, 0x08, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_archive_v2_format_proto_rawDescOnce sync.Once
	file_com_coralogix_archive_v2_format_proto_rawDescData = file_com_coralogix_archive_v2_format_proto_rawDesc
)

func file_com_coralogix_archive_v2_format_proto_rawDescGZIP() []byte {
	file_com_coralogix_archive_v2_format_proto_rawDescOnce.Do(func() {
		file_com_coralogix_archive_v2_format_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_archive_v2_format_proto_rawDescData)
	})
	return file_com_coralogix_archive_v2_format_proto_rawDescData
}

var file_com_coralogix_archive_v2_format_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogix_archive_v2_format_proto_goTypes = []interface{}{
	(*LogsAvro)(nil),            // 0: com.coralogix.archive.v2.LogsAvro
	(*SpansAvro)(nil),           // 1: com.coralogix.archive.v2.SpansAvro
	(*Format)(nil),              // 2: com.coralogix.archive.v2.Format
	(*v1.GenericEventAvro)(nil), // 3: com.coralogix.archive.format.generic.v1.GenericEventAvro
	(*v11.WideParquet)(nil),     // 4: com.coralogix.archive.format.wide_parquet.v1.WideParquet
}
var file_com_coralogix_archive_v2_format_proto_depIdxs = []int32{
	0, // 0: com.coralogix.archive.v2.Format.logs_avro:type_name -> com.coralogix.archive.v2.LogsAvro
	1, // 1: com.coralogix.archive.v2.Format.spans_avro:type_name -> com.coralogix.archive.v2.SpansAvro
	3, // 2: com.coralogix.archive.v2.Format.generic_event_avro:type_name -> com.coralogix.archive.format.generic.v1.GenericEventAvro
	4, // 3: com.coralogix.archive.v2.Format.wide_parquet:type_name -> com.coralogix.archive.format.wide_parquet.v1.WideParquet
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_com_coralogix_archive_v2_format_proto_init() }
func file_com_coralogix_archive_v2_format_proto_init() {
	if File_com_coralogix_archive_v2_format_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_archive_v2_format_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogsAvro); i {
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
		file_com_coralogix_archive_v2_format_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpansAvro); i {
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
		file_com_coralogix_archive_v2_format_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Format); i {
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
	file_com_coralogix_archive_v2_format_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Format_LogsAvro)(nil),
		(*Format_SpansAvro)(nil),
		(*Format_GenericEventAvro)(nil),
		(*Format_WideParquet)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_archive_v2_format_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_archive_v2_format_proto_goTypes,
		DependencyIndexes: file_com_coralogix_archive_v2_format_proto_depIdxs,
		MessageInfos:      file_com_coralogix_archive_v2_format_proto_msgTypes,
	}.Build()
	File_com_coralogix_archive_v2_format_proto = out.File
	file_com_coralogix_archive_v2_format_proto_rawDesc = nil
	file_com_coralogix_archive_v2_format_proto_goTypes = nil
	file_com_coralogix_archive_v2_format_proto_depIdxs = nil
}
