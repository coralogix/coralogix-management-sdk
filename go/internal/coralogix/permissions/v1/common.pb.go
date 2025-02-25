// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.26.1
// source: com/coralogix/permissions/v1/common.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type UserId struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserId) Reset() {
	*x = UserId{}
	mi := &file_com_coralogix_permissions_v1_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserId) ProtoMessage() {}

func (x *UserId) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_permissions_v1_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserId.ProtoReflect.Descriptor instead.
func (*UserId) Descriptor() ([]byte, []int) {
	return file_com_coralogix_permissions_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *UserId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TeamId struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TeamId) Reset() {
	*x = TeamId{}
	mi := &file_com_coralogix_permissions_v1_common_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TeamId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamId) ProtoMessage() {}

func (x *TeamId) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_permissions_v1_common_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamId.ProtoReflect.Descriptor instead.
func (*TeamId) Descriptor() ([]byte, []int) {
	return file_com_coralogix_permissions_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *TeamId) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UserAccountId struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserAccountId) Reset() {
	*x = UserAccountId{}
	mi := &file_com_coralogix_permissions_v1_common_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserAccountId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAccountId) ProtoMessage() {}

func (x *UserAccountId) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_permissions_v1_common_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAccountId.ProtoReflect.Descriptor instead.
func (*UserAccountId) Descriptor() ([]byte, []int) {
	return file_com_coralogix_permissions_v1_common_proto_rawDescGZIP(), []int{2}
}

func (x *UserAccountId) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type OrganizationId struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrganizationId) Reset() {
	*x = OrganizationId{}
	mi := &file_com_coralogix_permissions_v1_common_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrganizationId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationId) ProtoMessage() {}

func (x *OrganizationId) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_permissions_v1_common_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationId.ProtoReflect.Descriptor instead.
func (*OrganizationId) Descriptor() ([]byte, []int) {
	return file_com_coralogix_permissions_v1_common_proto_rawDescGZIP(), []int{3}
}

func (x *OrganizationId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RoleId struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RoleId) Reset() {
	*x = RoleId{}
	mi := &file_com_coralogix_permissions_v1_common_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoleId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleId) ProtoMessage() {}

func (x *RoleId) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_permissions_v1_common_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleId.ProtoReflect.Descriptor instead.
func (*RoleId) Descriptor() ([]byte, []int) {
	return file_com_coralogix_permissions_v1_common_proto_rawDescGZIP(), []int{4}
}

func (x *RoleId) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type OrgGroupId struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrgGroupId) Reset() {
	*x = OrgGroupId{}
	mi := &file_com_coralogix_permissions_v1_common_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrgGroupId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrgGroupId) ProtoMessage() {}

func (x *OrgGroupId) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_permissions_v1_common_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrgGroupId.ProtoReflect.Descriptor instead.
func (*OrgGroupId) Descriptor() ([]byte, []int) {
	return file_com_coralogix_permissions_v1_common_proto_rawDescGZIP(), []int{5}
}

func (x *OrgGroupId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TeamGroupId struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TeamGroupId) Reset() {
	*x = TeamGroupId{}
	mi := &file_com_coralogix_permissions_v1_common_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TeamGroupId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamGroupId) ProtoMessage() {}

func (x *TeamGroupId) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_permissions_v1_common_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamGroupId.ProtoReflect.Descriptor instead.
func (*TeamGroupId) Descriptor() ([]byte, []int) {
	return file_com_coralogix_permissions_v1_common_proto_rawDescGZIP(), []int{6}
}

func (x *TeamGroupId) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ScopeId struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ScopeId) Reset() {
	*x = ScopeId{}
	mi := &file_com_coralogix_permissions_v1_common_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScopeId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScopeId) ProtoMessage() {}

func (x *ScopeId) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_permissions_v1_common_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScopeId.ProtoReflect.Descriptor instead.
func (*ScopeId) Descriptor() ([]byte, []int) {
	return file_com_coralogix_permissions_v1_common_proto_rawDescGZIP(), []int{7}
}

func (x *ScopeId) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_com_coralogix_permissions_v1_common_proto protoreflect.FileDescriptor

var file_com_coralogix_permissions_v1_common_proto_rawDesc = []byte{
	0x0a, 0x29, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x18, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x18, 0x0a, 0x06, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1f, 0x0a,
	0x0d, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x20,
	0x0a, 0x0e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x18, 0x0a, 0x06, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1c, 0x0a, 0x0a, 0x4f, 0x72,
	0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x94, 0x02, 0x0a, 0x0b, 0x54, 0x65, 0x61,
	0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x3a, 0xf4, 0x01, 0x92, 0x41, 0xf0, 0x01, 0x0a,
	0x5b, 0x2a, 0x0b, 0x54, 0x65, 0x61, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x32, 0x4c,
	0x54, 0x68, 0x69, 0x73, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x20, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x73, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x61,
	0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x61,
	0x20, 0x74, 0x65, 0x61, 0x6d, 0x20, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x2a, 0x90, 0x01, 0x0a,
	0x1a, 0x46, 0x69, 0x6e, 0x64, 0x20, 0x6f, 0x75, 0x74, 0x20, 0x6d, 0x6f, 0x72, 0x65, 0x20, 0x61,
	0x62, 0x6f, 0x75, 0x74, 0x20, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x72, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x63, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x67, 0x75, 0x69,
	0x64, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2d, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x2d, 0x75, 0x73,
	0x65, 0x72, 0x2d, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2d, 0x61, 0x6e, 0x64, 0x2d, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x73, 0x2d, 0x76, 0x69, 0x61, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x22,
	0x19, 0x0a, 0x07, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_com_coralogix_permissions_v1_common_proto_rawDescOnce sync.Once
	file_com_coralogix_permissions_v1_common_proto_rawDescData = file_com_coralogix_permissions_v1_common_proto_rawDesc
)

func file_com_coralogix_permissions_v1_common_proto_rawDescGZIP() []byte {
	file_com_coralogix_permissions_v1_common_proto_rawDescOnce.Do(func() {
		file_com_coralogix_permissions_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_permissions_v1_common_proto_rawDescData)
	})
	return file_com_coralogix_permissions_v1_common_proto_rawDescData
}

var file_com_coralogix_permissions_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_com_coralogix_permissions_v1_common_proto_goTypes = []any{
	(*UserId)(nil),         // 0: com.coralogix.permissions.v1.UserId
	(*TeamId)(nil),         // 1: com.coralogix.permissions.v1.TeamId
	(*UserAccountId)(nil),  // 2: com.coralogix.permissions.v1.UserAccountId
	(*OrganizationId)(nil), // 3: com.coralogix.permissions.v1.OrganizationId
	(*RoleId)(nil),         // 4: com.coralogix.permissions.v1.RoleId
	(*OrgGroupId)(nil),     // 5: com.coralogix.permissions.v1.OrgGroupId
	(*TeamGroupId)(nil),    // 6: com.coralogix.permissions.v1.TeamGroupId
	(*ScopeId)(nil),        // 7: com.coralogix.permissions.v1.ScopeId
}
var file_com_coralogix_permissions_v1_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogix_permissions_v1_common_proto_init() }
func file_com_coralogix_permissions_v1_common_proto_init() {
	if File_com_coralogix_permissions_v1_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_permissions_v1_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_permissions_v1_common_proto_goTypes,
		DependencyIndexes: file_com_coralogix_permissions_v1_common_proto_depIdxs,
		MessageInfos:      file_com_coralogix_permissions_v1_common_proto_msgTypes,
	}.Build()
	File_com_coralogix_permissions_v1_common_proto = out.File
	file_com_coralogix_permissions_v1_common_proto_rawDesc = nil
	file_com_coralogix_permissions_v1_common_proto_goTypes = nil
	file_com_coralogix_permissions_v1_common_proto_depIdxs = nil
}
