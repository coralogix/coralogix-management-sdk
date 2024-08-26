// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.3
// source: com/coralogix/global_mapping/v1/context_data.proto

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

type Span struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags        map[string]*wrapperspb.StringValue `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ProcessTags map[string]*wrapperspb.StringValue `protobuf:"bytes,2,rep,name=process_tags,json=processTags,proto3" json:"process_tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Span) Reset() {
	*x = Span{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_global_mapping_v1_context_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Span) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Span) ProtoMessage() {}

func (x *Span) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_global_mapping_v1_context_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Span.ProtoReflect.Descriptor instead.
func (*Span) Descriptor() ([]byte, []int) {
	return file_com_coralogix_global_mapping_v1_context_data_proto_rawDescGZIP(), []int{0}
}

func (x *Span) GetTags() map[string]*wrapperspb.StringValue {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Span) GetProcessTags() map[string]*wrapperspb.StringValue {
	if x != nil {
		return x.ProcessTags
	}
	return nil
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata map[string]*wrapperspb.StringValue `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Text     *wrapperspb.StringValue            `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_global_mapping_v1_context_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_global_mapping_v1_context_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_com_coralogix_global_mapping_v1_context_data_proto_rawDescGZIP(), []int{1}
}

func (x *Log) GetMetadata() map[string]*wrapperspb.StringValue {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Log) GetText() *wrapperspb.StringValue {
	if x != nil {
		return x.Text
	}
	return nil
}

var File_com_coralogix_global_mapping_v1_context_data_proto protoreflect.FileDescriptor

var file_com_coralogix_global_mapping_v1_context_data_proto_rawDesc = []byte{
	0x0a, 0x32, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x02, 0x0a, 0x04, 0x53, 0x70, 0x61, 0x6e, 0x12, 0x43,
	0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x70, 0x61, 0x6e, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x12, 0x59, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x61, 0x6e,
	0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x54, 0x61, 0x67, 0x73, 0x1a, 0x55,
	0x0a, 0x09, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x32, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x5c, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0xe2, 0x01, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x4e, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x6f, 0x67, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x30, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x1a, 0x59, 0x0a,
	0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_global_mapping_v1_context_data_proto_rawDescOnce sync.Once
	file_com_coralogix_global_mapping_v1_context_data_proto_rawDescData = file_com_coralogix_global_mapping_v1_context_data_proto_rawDesc
)

func file_com_coralogix_global_mapping_v1_context_data_proto_rawDescGZIP() []byte {
	file_com_coralogix_global_mapping_v1_context_data_proto_rawDescOnce.Do(func() {
		file_com_coralogix_global_mapping_v1_context_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_global_mapping_v1_context_data_proto_rawDescData)
	})
	return file_com_coralogix_global_mapping_v1_context_data_proto_rawDescData
}

var file_com_coralogix_global_mapping_v1_context_data_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_com_coralogix_global_mapping_v1_context_data_proto_goTypes = []interface{}{
	(*Span)(nil),                   // 0: com.coralogix.global_mapping.v1.Span
	(*Log)(nil),                    // 1: com.coralogix.global_mapping.v1.Log
	nil,                            // 2: com.coralogix.global_mapping.v1.Span.TagsEntry
	nil,                            // 3: com.coralogix.global_mapping.v1.Span.ProcessTagsEntry
	nil,                            // 4: com.coralogix.global_mapping.v1.Log.MetadataEntry
	(*wrapperspb.StringValue)(nil), // 5: google.protobuf.StringValue
}
var file_com_coralogix_global_mapping_v1_context_data_proto_depIdxs = []int32{
	2, // 0: com.coralogix.global_mapping.v1.Span.tags:type_name -> com.coralogix.global_mapping.v1.Span.TagsEntry
	3, // 1: com.coralogix.global_mapping.v1.Span.process_tags:type_name -> com.coralogix.global_mapping.v1.Span.ProcessTagsEntry
	4, // 2: com.coralogix.global_mapping.v1.Log.metadata:type_name -> com.coralogix.global_mapping.v1.Log.MetadataEntry
	5, // 3: com.coralogix.global_mapping.v1.Log.text:type_name -> google.protobuf.StringValue
	5, // 4: com.coralogix.global_mapping.v1.Span.TagsEntry.value:type_name -> google.protobuf.StringValue
	5, // 5: com.coralogix.global_mapping.v1.Span.ProcessTagsEntry.value:type_name -> google.protobuf.StringValue
	5, // 6: com.coralogix.global_mapping.v1.Log.MetadataEntry.value:type_name -> google.protobuf.StringValue
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_com_coralogix_global_mapping_v1_context_data_proto_init() }
func file_com_coralogix_global_mapping_v1_context_data_proto_init() {
	if File_com_coralogix_global_mapping_v1_context_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_global_mapping_v1_context_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Span); i {
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
		file_com_coralogix_global_mapping_v1_context_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
			RawDescriptor: file_com_coralogix_global_mapping_v1_context_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_global_mapping_v1_context_data_proto_goTypes,
		DependencyIndexes: file_com_coralogix_global_mapping_v1_context_data_proto_depIdxs,
		MessageInfos:      file_com_coralogix_global_mapping_v1_context_data_proto_msgTypes,
	}.Build()
	File_com_coralogix_global_mapping_v1_context_data_proto = out.File
	file_com_coralogix_global_mapping_v1_context_data_proto_rawDesc = nil
	file_com_coralogix_global_mapping_v1_context_data_proto_goTypes = nil
	file_com_coralogix_global_mapping_v1_context_data_proto_depIdxs = nil
}
