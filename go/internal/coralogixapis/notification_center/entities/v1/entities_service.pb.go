// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.26.1
// source: com/coralogixapis/notification_center/entities/v1/entities_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListEntityTypesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListEntityTypesRequest) Reset() {
	*x = ListEntityTypesRequest{}
	mi := &file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListEntityTypesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEntityTypesRequest) ProtoMessage() {}

func (x *ListEntityTypesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEntityTypesRequest.ProtoReflect.Descriptor instead.
func (*ListEntityTypesRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_rawDescGZIP(), []int{0}
}

type ListEntityTypesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EntityTypes   []string               `protobuf:"bytes,1,rep,name=entity_types,json=entityTypes,proto3" json:"entity_types,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListEntityTypesResponse) Reset() {
	*x = ListEntityTypesResponse{}
	mi := &file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListEntityTypesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEntityTypesResponse) ProtoMessage() {}

func (x *ListEntityTypesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEntityTypesResponse.ProtoReflect.Descriptor instead.
func (*ListEntityTypesResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListEntityTypesResponse) GetEntityTypes() []string {
	if x != nil {
		return x.EntityTypes
	}
	return nil
}

var File_com_coralogixapis_notification_center_entities_v1_entities_service_proto protoreflect.FileDescriptor

var file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_rawDesc = []byte{
	0x0a, 0x48, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x31, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x18, 0x0a,
	0x16, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3c, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x32, 0xf8, 0x01, 0x0a, 0x0f, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xe4, 0x01, 0x0a, 0x0f, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x49, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x4a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x34, 0x3a, 0x01, 0x2a, 0x12,
	0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x3a, 0x6c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_rawDescOnce sync.Once
	file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_rawDescData = file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_rawDesc
)

func file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_rawDescData)
	})
	return file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_rawDescData
}

var file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_goTypes = []any{
	(*ListEntityTypesRequest)(nil),  // 0: com.coralogixapis.notification_center.entities.v1.ListEntityTypesRequest
	(*ListEntityTypesResponse)(nil), // 1: com.coralogixapis.notification_center.entities.v1.ListEntityTypesResponse
}
var file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_depIdxs = []int32{
	0, // 0: com.coralogixapis.notification_center.entities.v1.EntitiesService.ListEntityTypes:input_type -> com.coralogixapis.notification_center.entities.v1.ListEntityTypesRequest
	1, // 1: com.coralogixapis.notification_center.entities.v1.EntitiesService.ListEntityTypes:output_type -> com.coralogixapis.notification_center.entities.v1.ListEntityTypesResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_init() }
func file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_init() {
	if File_com_coralogixapis_notification_center_entities_v1_entities_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_notification_center_entities_v1_entities_service_proto = out.File
	file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_rawDesc = nil
	file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_goTypes = nil
	file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_depIdxs = nil
}
