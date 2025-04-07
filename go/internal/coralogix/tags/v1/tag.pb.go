// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.26.1
// source: com/coralogix/tags/v1/tag.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Status int32

const (
	Status_STATUS_UNSPECIFIED Status = 0
	Status_STATUS_SUCCESSFUL  Status = 1
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUCCESSFUL",
	}
	Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_SUCCESSFUL":  1,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_tags_v1_tag_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_com_coralogix_tags_v1_tag_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_tags_v1_tag_proto_rawDescGZIP(), []int{0}
}

type Tag struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Id            *wrapperspb.UInt64Value   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Key           *wrapperspb.StringValue   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Name          *wrapperspb.StringValue   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CompanyId     *wrapperspb.UInt32Value   `protobuf:"bytes,4,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Status        Status                    `protobuf:"varint,5,opt,name=status,proto3,enum=com.coralogix.tags.v1.Status" json:"status,omitempty"`
	IconUrl       *wrapperspb.StringValue   `protobuf:"bytes,6,opt,name=icon_url,json=iconUrl,proto3" json:"icon_url,omitempty"`
	Timestamp     *timestamppb.Timestamp    `protobuf:"bytes,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Application   []*wrapperspb.StringValue `protobuf:"bytes,8,rep,name=application,proto3" json:"application,omitempty"`
	Subsystem     []*wrapperspb.StringValue `protobuf:"bytes,9,rep,name=subsystem,proto3" json:"subsystem,omitempty"`
	UpdatedAt     *timestamppb.Timestamp    `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedAt     *timestamppb.Timestamp    `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Type          *TagType                  `protobuf:"bytes,13,opt,name=type,proto3" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Tag) Reset() {
	*x = Tag{}
	mi := &file_com_coralogix_tags_v1_tag_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_tags_v1_tag_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_com_coralogix_tags_v1_tag_proto_rawDescGZIP(), []int{0}
}

func (x *Tag) GetId() *wrapperspb.UInt64Value {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Tag) GetKey() *wrapperspb.StringValue {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *Tag) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *Tag) GetCompanyId() *wrapperspb.UInt32Value {
	if x != nil {
		return x.CompanyId
	}
	return nil
}

func (x *Tag) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_STATUS_UNSPECIFIED
}

func (x *Tag) GetIconUrl() *wrapperspb.StringValue {
	if x != nil {
		return x.IconUrl
	}
	return nil
}

func (x *Tag) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Tag) GetApplication() []*wrapperspb.StringValue {
	if x != nil {
		return x.Application
	}
	return nil
}

func (x *Tag) GetSubsystem() []*wrapperspb.StringValue {
	if x != nil {
		return x.Subsystem
	}
	return nil
}

func (x *Tag) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Tag) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Tag) GetType() *TagType {
	if x != nil {
		return x.Type
	}
	return nil
}

var File_com_coralogix_tags_v1_tag_proto protoreflect.FileDescriptor

const file_com_coralogix_tags_v1_tag_proto_rawDesc = "" +
	"\n" +
	"\x1fcom/coralogix/tags/v1/tag.proto\x12\x15com.coralogix.tags.v1\x1a\x1egoogle/protobuf/wrappers.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a$com/coralogix/tags/v1/tag_type.proto\"\xa2\x05\n" +
	"\x03Tag\x12,\n" +
	"\x02id\x18\x01 \x01(\v2\x1c.google.protobuf.UInt64ValueR\x02id\x12.\n" +
	"\x03key\x18\x02 \x01(\v2\x1c.google.protobuf.StringValueR\x03key\x120\n" +
	"\x04name\x18\x03 \x01(\v2\x1c.google.protobuf.StringValueR\x04name\x12;\n" +
	"\n" +
	"company_id\x18\x04 \x01(\v2\x1c.google.protobuf.UInt32ValueR\tcompanyId\x125\n" +
	"\x06status\x18\x05 \x01(\x0e2\x1d.com.coralogix.tags.v1.StatusR\x06status\x127\n" +
	"\bicon_url\x18\x06 \x01(\v2\x1c.google.protobuf.StringValueR\aiconUrl\x128\n" +
	"\ttimestamp\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\ttimestamp\x12>\n" +
	"\vapplication\x18\b \x03(\v2\x1c.google.protobuf.StringValueR\vapplication\x12:\n" +
	"\tsubsystem\x18\t \x03(\v2\x1c.google.protobuf.StringValueR\tsubsystem\x129\n" +
	"\n" +
	"updated_at\x18\n" +
	" \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\x129\n" +
	"\n" +
	"created_at\x18\v \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x122\n" +
	"\x04type\x18\r \x01(\v2\x1e.com.coralogix.tags.v1.TagTypeR\x04type*7\n" +
	"\x06Status\x12\x16\n" +
	"\x12STATUS_UNSPECIFIED\x10\x00\x12\x15\n" +
	"\x11STATUS_SUCCESSFUL\x10\x01b\x06proto3"

var (
	file_com_coralogix_tags_v1_tag_proto_rawDescOnce sync.Once
	file_com_coralogix_tags_v1_tag_proto_rawDescData []byte
)

func file_com_coralogix_tags_v1_tag_proto_rawDescGZIP() []byte {
	file_com_coralogix_tags_v1_tag_proto_rawDescOnce.Do(func() {
		file_com_coralogix_tags_v1_tag_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogix_tags_v1_tag_proto_rawDesc), len(file_com_coralogix_tags_v1_tag_proto_rawDesc)))
	})
	return file_com_coralogix_tags_v1_tag_proto_rawDescData
}

var file_com_coralogix_tags_v1_tag_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogix_tags_v1_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogix_tags_v1_tag_proto_goTypes = []any{
	(Status)(0),                    // 0: com.coralogix.tags.v1.Status
	(*Tag)(nil),                    // 1: com.coralogix.tags.v1.Tag
	(*wrapperspb.UInt64Value)(nil), // 2: google.protobuf.UInt64Value
	(*wrapperspb.StringValue)(nil), // 3: google.protobuf.StringValue
	(*wrapperspb.UInt32Value)(nil), // 4: google.protobuf.UInt32Value
	(*timestamppb.Timestamp)(nil),  // 5: google.protobuf.Timestamp
	(*TagType)(nil),                // 6: com.coralogix.tags.v1.TagType
}
var file_com_coralogix_tags_v1_tag_proto_depIdxs = []int32{
	2,  // 0: com.coralogix.tags.v1.Tag.id:type_name -> google.protobuf.UInt64Value
	3,  // 1: com.coralogix.tags.v1.Tag.key:type_name -> google.protobuf.StringValue
	3,  // 2: com.coralogix.tags.v1.Tag.name:type_name -> google.protobuf.StringValue
	4,  // 3: com.coralogix.tags.v1.Tag.company_id:type_name -> google.protobuf.UInt32Value
	0,  // 4: com.coralogix.tags.v1.Tag.status:type_name -> com.coralogix.tags.v1.Status
	3,  // 5: com.coralogix.tags.v1.Tag.icon_url:type_name -> google.protobuf.StringValue
	5,  // 6: com.coralogix.tags.v1.Tag.timestamp:type_name -> google.protobuf.Timestamp
	3,  // 7: com.coralogix.tags.v1.Tag.application:type_name -> google.protobuf.StringValue
	3,  // 8: com.coralogix.tags.v1.Tag.subsystem:type_name -> google.protobuf.StringValue
	5,  // 9: com.coralogix.tags.v1.Tag.updated_at:type_name -> google.protobuf.Timestamp
	5,  // 10: com.coralogix.tags.v1.Tag.created_at:type_name -> google.protobuf.Timestamp
	6,  // 11: com.coralogix.tags.v1.Tag.type:type_name -> com.coralogix.tags.v1.TagType
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_com_coralogix_tags_v1_tag_proto_init() }
func file_com_coralogix_tags_v1_tag_proto_init() {
	if File_com_coralogix_tags_v1_tag_proto != nil {
		return
	}
	file_com_coralogix_tags_v1_tag_type_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogix_tags_v1_tag_proto_rawDesc), len(file_com_coralogix_tags_v1_tag_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_tags_v1_tag_proto_goTypes,
		DependencyIndexes: file_com_coralogix_tags_v1_tag_proto_depIdxs,
		EnumInfos:         file_com_coralogix_tags_v1_tag_proto_enumTypes,
		MessageInfos:      file_com_coralogix_tags_v1_tag_proto_msgTypes,
	}.Build()
	File_com_coralogix_tags_v1_tag_proto = out.File
	file_com_coralogix_tags_v1_tag_proto_goTypes = nil
	file_com_coralogix_tags_v1_tag_proto_depIdxs = nil
}
