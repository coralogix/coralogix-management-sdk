// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/notification_center/entities/v1/entities_service.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
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

const file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_rawDesc = "" +
	"\n" +
	"Hcom/coralogixapis/notification_center/entities/v1/entities_service.proto\x121com.coralogixapis.notification_center.entities.v1\x1a google/protobuf/descriptor.proto\x1a\x1cgoogle/api/annotations.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\xf0\x01\n" +
	"\x16ListEntityTypesRequest:\xd5\x01\x92A\xd1\x01\n" +
	"R*\x11List Entity Types2=Request to list entity types supported by Notification Center*{\n" +
	"'Find out more about notification center\x12Phttps://coralogix.com/docs/user-guides/notification-center/introduction/welcome/\"<\n" +
	"\x17ListEntityTypesResponse\x12!\n" +
	"\fentity_types\x18\x01 \x03(\tR\ventityTypes2\xf6\x01\n" +
	"\x0fEntitiesService\x12\xe2\x01\n" +
	"\x0fListEntityTypes\x12I.com.coralogixapis.notification_center.entities.v1.ListEntityTypesRequest\x1aJ.com.coralogixapis.notification_center.entities.v1.ListEntityTypesResponse\"8\x82\xd3\xe4\x93\x022\x120/v1/notification-center/entities:listEntityTypesb\x06proto3"

var (
	file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_rawDescOnce sync.Once
	file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_rawDescData []byte
)

func file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_rawDesc), len(file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_rawDesc)))
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_rawDesc), len(file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_rawDesc)),
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
	file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_goTypes = nil
	file_com_coralogixapis_notification_center_entities_v1_entities_service_proto_depIdxs = nil
}
