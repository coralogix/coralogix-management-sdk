// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: com/coralogix/archive/private/state/v1/s3_object.proto

package v1

import (
	v2 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/archive/v2"
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

type S3Object struct {
	state     protoimpl.MessageState `protogen:"open.v1"`
	Bucket    string                 `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Key       string                 `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	RowCount  *int64                 `protobuf:"varint,3,opt,name=row_count,json=rowCount,proto3,oneof" json:"row_count,omitempty"`
	FileSplit *FileSplit             `protobuf:"bytes,4,opt,name=file_split,json=fileSplit,proto3,oneof" json:"file_split,omitempty"`
	// Types that are valid to be assigned to PhysicalTarget:
	//
	//	*S3Object_S3
	//	*S3Object_IbmCos
	PhysicalTarget isS3Object_PhysicalTarget `protobuf_oneof:"physical_target"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *S3Object) Reset() {
	*x = S3Object{}
	mi := &file_com_coralogix_archive_private_state_v1_s3_object_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *S3Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3Object) ProtoMessage() {}

func (x *S3Object) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_private_state_v1_s3_object_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3Object.ProtoReflect.Descriptor instead.
func (*S3Object) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_private_state_v1_s3_object_proto_rawDescGZIP(), []int{0}
}

func (x *S3Object) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *S3Object) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *S3Object) GetRowCount() int64 {
	if x != nil && x.RowCount != nil {
		return *x.RowCount
	}
	return 0
}

func (x *S3Object) GetFileSplit() *FileSplit {
	if x != nil {
		return x.FileSplit
	}
	return nil
}

func (x *S3Object) GetPhysicalTarget() isS3Object_PhysicalTarget {
	if x != nil {
		return x.PhysicalTarget
	}
	return nil
}

func (x *S3Object) GetS3() *v2.S3TargetSpec {
	if x != nil {
		if x, ok := x.PhysicalTarget.(*S3Object_S3); ok {
			return x.S3
		}
	}
	return nil
}

func (x *S3Object) GetIbmCos() *v2.IBMCosTargetSpec {
	if x != nil {
		if x, ok := x.PhysicalTarget.(*S3Object_IbmCos); ok {
			return x.IbmCos
		}
	}
	return nil
}

type isS3Object_PhysicalTarget interface {
	isS3Object_PhysicalTarget()
}

type S3Object_S3 struct {
	S3 *v2.S3TargetSpec `protobuf:"bytes,5,opt,name=s3,proto3,oneof"`
}

type S3Object_IbmCos struct {
	IbmCos *v2.IBMCosTargetSpec `protobuf:"bytes,6,opt,name=ibm_cos,json=ibmCos,proto3,oneof"`
}

func (*S3Object_S3) isS3Object_PhysicalTarget() {}

func (*S3Object_IbmCos) isS3Object_PhysicalTarget() {}

var File_com_coralogix_archive_private_state_v1_s3_object_proto protoreflect.FileDescriptor

var file_com_coralogix_archive_private_state_v1_s3_object_proto_rawDesc = []byte{
	0x0a, 0x36, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x33, 0x5f, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x37, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x70,
	0x6c, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x2f, 0x76, 0x32, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xde, 0x02, 0x0a, 0x08, 0x53, 0x33, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x09, 0x72, 0x6f, 0x77, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x08, 0x72, 0x6f,
	0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x55, 0x0a, 0x0a, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x70, 0x6c, 0x69, 0x74,
	0x48, 0x02, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x38, 0x0a, 0x02, 0x73, 0x33, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x33, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x53, 0x70, 0x65, 0x63, 0x48, 0x00, 0x52, 0x02, 0x73, 0x33, 0x12, 0x45, 0x0a, 0x07, 0x69, 0x62,
	0x6d, 0x5f, 0x63, 0x6f, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x42, 0x4d, 0x43, 0x6f, 0x73, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x48, 0x00, 0x52, 0x06, 0x69, 0x62, 0x6d, 0x43, 0x6f,
	0x73, 0x42, 0x11, 0x0a, 0x0f, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x72, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x70, 0x6c, 0x69,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_archive_private_state_v1_s3_object_proto_rawDescOnce sync.Once
	file_com_coralogix_archive_private_state_v1_s3_object_proto_rawDescData = file_com_coralogix_archive_private_state_v1_s3_object_proto_rawDesc
)

func file_com_coralogix_archive_private_state_v1_s3_object_proto_rawDescGZIP() []byte {
	file_com_coralogix_archive_private_state_v1_s3_object_proto_rawDescOnce.Do(func() {
		file_com_coralogix_archive_private_state_v1_s3_object_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_archive_private_state_v1_s3_object_proto_rawDescData)
	})
	return file_com_coralogix_archive_private_state_v1_s3_object_proto_rawDescData
}

var file_com_coralogix_archive_private_state_v1_s3_object_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogix_archive_private_state_v1_s3_object_proto_goTypes = []any{
	(*S3Object)(nil),            // 0: com.coralogix.archive.private.state.v1.S3Object
	(*FileSplit)(nil),           // 1: com.coralogix.archive.private.state.v1.FileSplit
	(*v2.S3TargetSpec)(nil),     // 2: com.coralogix.archive.v2.S3TargetSpec
	(*v2.IBMCosTargetSpec)(nil), // 3: com.coralogix.archive.v2.IBMCosTargetSpec
}
var file_com_coralogix_archive_private_state_v1_s3_object_proto_depIdxs = []int32{
	1, // 0: com.coralogix.archive.private.state.v1.S3Object.file_split:type_name -> com.coralogix.archive.private.state.v1.FileSplit
	2, // 1: com.coralogix.archive.private.state.v1.S3Object.s3:type_name -> com.coralogix.archive.v2.S3TargetSpec
	3, // 2: com.coralogix.archive.private.state.v1.S3Object.ibm_cos:type_name -> com.coralogix.archive.v2.IBMCosTargetSpec
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_com_coralogix_archive_private_state_v1_s3_object_proto_init() }
func file_com_coralogix_archive_private_state_v1_s3_object_proto_init() {
	if File_com_coralogix_archive_private_state_v1_s3_object_proto != nil {
		return
	}
	file_com_coralogix_archive_private_state_v1_file_split_proto_init()
	file_com_coralogix_archive_private_state_v1_s3_object_proto_msgTypes[0].OneofWrappers = []any{
		(*S3Object_S3)(nil),
		(*S3Object_IbmCos)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_archive_private_state_v1_s3_object_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_archive_private_state_v1_s3_object_proto_goTypes,
		DependencyIndexes: file_com_coralogix_archive_private_state_v1_s3_object_proto_depIdxs,
		MessageInfos:      file_com_coralogix_archive_private_state_v1_s3_object_proto_msgTypes,
	}.Build()
	File_com_coralogix_archive_private_state_v1_s3_object_proto = out.File
	file_com_coralogix_archive_private_state_v1_s3_object_proto_rawDesc = nil
	file_com_coralogix_archive_private_state_v1_s3_object_proto_goTypes = nil
	file_com_coralogix_archive_private_state_v1_s3_object_proto_depIdxs = nil
}
