// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: com/coralogix/archive/v1/target.proto

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

type Target struct {
	state             protoimpl.MessageState  `protogen:"open.v1"`
	ArchivingFormatId *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=archiving_format_id,json=archivingFormatId,proto3" json:"archiving_format_id,omitempty"`
	IsActive          *wrapperspb.BoolValue   `protobuf:"bytes,2,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	Region            *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	Bucket            *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=bucket,proto3" json:"bucket,omitempty"`
	EnableTags        *wrapperspb.BoolValue   `protobuf:"bytes,5,opt,name=enable_tags,json=enableTags,proto3" json:"enable_tags,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *Target) Reset() {
	*x = Target{}
	mi := &file_com_coralogix_archive_v1_target_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Target) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Target) ProtoMessage() {}

func (x *Target) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_v1_target_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Target.ProtoReflect.Descriptor instead.
func (*Target) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_v1_target_proto_rawDescGZIP(), []int{0}
}

func (x *Target) GetArchivingFormatId() *wrapperspb.StringValue {
	if x != nil {
		return x.ArchivingFormatId
	}
	return nil
}

func (x *Target) GetIsActive() *wrapperspb.BoolValue {
	if x != nil {
		return x.IsActive
	}
	return nil
}

func (x *Target) GetRegion() *wrapperspb.StringValue {
	if x != nil {
		return x.Region
	}
	return nil
}

func (x *Target) GetBucket() *wrapperspb.StringValue {
	if x != nil {
		return x.Bucket
	}
	return nil
}

func (x *Target) GetEnableTags() *wrapperspb.BoolValue {
	if x != nil {
		return x.EnableTags
	}
	return nil
}

var File_com_coralogix_archive_v1_target_proto protoreflect.FileDescriptor

var file_com_coralogix_archive_v1_target_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb8, 0x02, 0x0a, 0x06, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x4c, 0x0a, 0x13,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x11, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69,
	0x6e, 0x67, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x09, 0x69, 0x73,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x06, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12,
	0x3b, 0x0a, 0x0b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0a, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x61, 0x67, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_archive_v1_target_proto_rawDescOnce sync.Once
	file_com_coralogix_archive_v1_target_proto_rawDescData = file_com_coralogix_archive_v1_target_proto_rawDesc
)

func file_com_coralogix_archive_v1_target_proto_rawDescGZIP() []byte {
	file_com_coralogix_archive_v1_target_proto_rawDescOnce.Do(func() {
		file_com_coralogix_archive_v1_target_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_archive_v1_target_proto_rawDescData)
	})
	return file_com_coralogix_archive_v1_target_proto_rawDescData
}

var file_com_coralogix_archive_v1_target_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogix_archive_v1_target_proto_goTypes = []any{
	(*Target)(nil),                 // 0: com.coralogix.archive.v1.Target
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
	(*wrapperspb.BoolValue)(nil),   // 2: google.protobuf.BoolValue
}
var file_com_coralogix_archive_v1_target_proto_depIdxs = []int32{
	1, // 0: com.coralogix.archive.v1.Target.archiving_format_id:type_name -> google.protobuf.StringValue
	2, // 1: com.coralogix.archive.v1.Target.is_active:type_name -> google.protobuf.BoolValue
	1, // 2: com.coralogix.archive.v1.Target.region:type_name -> google.protobuf.StringValue
	1, // 3: com.coralogix.archive.v1.Target.bucket:type_name -> google.protobuf.StringValue
	2, // 4: com.coralogix.archive.v1.Target.enable_tags:type_name -> google.protobuf.BoolValue
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_com_coralogix_archive_v1_target_proto_init() }
func file_com_coralogix_archive_v1_target_proto_init() {
	if File_com_coralogix_archive_v1_target_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_archive_v1_target_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_archive_v1_target_proto_goTypes,
		DependencyIndexes: file_com_coralogix_archive_v1_target_proto_depIdxs,
		MessageInfos:      file_com_coralogix_archive_v1_target_proto_msgTypes,
	}.Build()
	File_com_coralogix_archive_v1_target_proto = out.File
	file_com_coralogix_archive_v1_target_proto_rawDesc = nil
	file_com_coralogix_archive_v1_target_proto_goTypes = nil
	file_com_coralogix_archive_v1_target_proto_depIdxs = nil
}
