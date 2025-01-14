// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: com/coralogix/integrations/v1/webhook.proto

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

type TimestampSourceGenerate struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TimestampSourceGenerate) Reset() {
	*x = TimestampSourceGenerate{}
	mi := &file_com_coralogix_integrations_v1_webhook_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimestampSourceGenerate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimestampSourceGenerate) ProtoMessage() {}

func (x *TimestampSourceGenerate) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_integrations_v1_webhook_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimestampSourceGenerate.ProtoReflect.Descriptor instead.
func (*TimestampSourceGenerate) Descriptor() ([]byte, []int) {
	return file_com_coralogix_integrations_v1_webhook_proto_rawDescGZIP(), []int{0}
}

type TimestampFormatAutomatic struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TimestampFormatAutomatic) Reset() {
	*x = TimestampFormatAutomatic{}
	mi := &file_com_coralogix_integrations_v1_webhook_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimestampFormatAutomatic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimestampFormatAutomatic) ProtoMessage() {}

func (x *TimestampFormatAutomatic) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_integrations_v1_webhook_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimestampFormatAutomatic.ProtoReflect.Descriptor instead.
func (*TimestampFormatAutomatic) Descriptor() ([]byte, []int) {
	return file_com_coralogix_integrations_v1_webhook_proto_rawDescGZIP(), []int{1}
}

type TimestampFormatCustom struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Format        *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=format,proto3" json:"format,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TimestampFormatCustom) Reset() {
	*x = TimestampFormatCustom{}
	mi := &file_com_coralogix_integrations_v1_webhook_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimestampFormatCustom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimestampFormatCustom) ProtoMessage() {}

func (x *TimestampFormatCustom) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_integrations_v1_webhook_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimestampFormatCustom.ProtoReflect.Descriptor instead.
func (*TimestampFormatCustom) Descriptor() ([]byte, []int) {
	return file_com_coralogix_integrations_v1_webhook_proto_rawDescGZIP(), []int{2}
}

func (x *TimestampFormatCustom) GetFormat() *wrapperspb.StringValue {
	if x != nil {
		return x.Format
	}
	return nil
}

type TimestampSourceJsonPath struct {
	state    protoimpl.MessageState  `protogen:"open.v1"`
	JsonPath *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=json_path,json=jsonPath,proto3" json:"json_path,omitempty"`
	// Types that are valid to be assigned to Format:
	//
	//	*TimestampSourceJsonPath_Automatic
	//	*TimestampSourceJsonPath_Custom
	Format        isTimestampSourceJsonPath_Format `protobuf_oneof:"format"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TimestampSourceJsonPath) Reset() {
	*x = TimestampSourceJsonPath{}
	mi := &file_com_coralogix_integrations_v1_webhook_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimestampSourceJsonPath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimestampSourceJsonPath) ProtoMessage() {}

func (x *TimestampSourceJsonPath) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_integrations_v1_webhook_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimestampSourceJsonPath.ProtoReflect.Descriptor instead.
func (*TimestampSourceJsonPath) Descriptor() ([]byte, []int) {
	return file_com_coralogix_integrations_v1_webhook_proto_rawDescGZIP(), []int{3}
}

func (x *TimestampSourceJsonPath) GetJsonPath() *wrapperspb.StringValue {
	if x != nil {
		return x.JsonPath
	}
	return nil
}

func (x *TimestampSourceJsonPath) GetFormat() isTimestampSourceJsonPath_Format {
	if x != nil {
		return x.Format
	}
	return nil
}

func (x *TimestampSourceJsonPath) GetAutomatic() *TimestampFormatAutomatic {
	if x != nil {
		if x, ok := x.Format.(*TimestampSourceJsonPath_Automatic); ok {
			return x.Automatic
		}
	}
	return nil
}

func (x *TimestampSourceJsonPath) GetCustom() *TimestampFormatCustom {
	if x != nil {
		if x, ok := x.Format.(*TimestampSourceJsonPath_Custom); ok {
			return x.Custom
		}
	}
	return nil
}

type isTimestampSourceJsonPath_Format interface {
	isTimestampSourceJsonPath_Format()
}

type TimestampSourceJsonPath_Automatic struct {
	Automatic *TimestampFormatAutomatic `protobuf:"bytes,101,opt,name=automatic,proto3,oneof"`
}

type TimestampSourceJsonPath_Custom struct {
	Custom *TimestampFormatCustom `protobuf:"bytes,102,opt,name=custom,proto3,oneof"`
}

func (*TimestampSourceJsonPath_Automatic) isTimestampSourceJsonPath_Format() {}

func (*TimestampSourceJsonPath_Custom) isTimestampSourceJsonPath_Format() {}

type TimestampSourceRegex struct {
	state protoimpl.MessageState  `protogen:"open.v1"`
	Regex *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=regex,proto3" json:"regex,omitempty"`
	// Types that are valid to be assigned to Format:
	//
	//	*TimestampSourceRegex_Automatic
	//	*TimestampSourceRegex_Custom
	Format        isTimestampSourceRegex_Format `protobuf_oneof:"format"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TimestampSourceRegex) Reset() {
	*x = TimestampSourceRegex{}
	mi := &file_com_coralogix_integrations_v1_webhook_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimestampSourceRegex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimestampSourceRegex) ProtoMessage() {}

func (x *TimestampSourceRegex) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_integrations_v1_webhook_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimestampSourceRegex.ProtoReflect.Descriptor instead.
func (*TimestampSourceRegex) Descriptor() ([]byte, []int) {
	return file_com_coralogix_integrations_v1_webhook_proto_rawDescGZIP(), []int{4}
}

func (x *TimestampSourceRegex) GetRegex() *wrapperspb.StringValue {
	if x != nil {
		return x.Regex
	}
	return nil
}

func (x *TimestampSourceRegex) GetFormat() isTimestampSourceRegex_Format {
	if x != nil {
		return x.Format
	}
	return nil
}

func (x *TimestampSourceRegex) GetAutomatic() *TimestampFormatAutomatic {
	if x != nil {
		if x, ok := x.Format.(*TimestampSourceRegex_Automatic); ok {
			return x.Automatic
		}
	}
	return nil
}

func (x *TimestampSourceRegex) GetCustom() *TimestampFormatCustom {
	if x != nil {
		if x, ok := x.Format.(*TimestampSourceRegex_Custom); ok {
			return x.Custom
		}
	}
	return nil
}

type isTimestampSourceRegex_Format interface {
	isTimestampSourceRegex_Format()
}

type TimestampSourceRegex_Automatic struct {
	Automatic *TimestampFormatAutomatic `protobuf:"bytes,101,opt,name=automatic,proto3,oneof"`
}

type TimestampSourceRegex_Custom struct {
	Custom *TimestampFormatCustom `protobuf:"bytes,102,opt,name=custom,proto3,oneof"`
}

func (*TimestampSourceRegex_Automatic) isTimestampSourceRegex_Format() {}

func (*TimestampSourceRegex_Custom) isTimestampSourceRegex_Format() {}

type JsonContentType struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to TimestampSource:
	//
	//	*JsonContentType_Generate
	//	*JsonContentType_JsonPath
	TimestampSource isJsonContentType_TimestampSource `protobuf_oneof:"timestamp_source"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *JsonContentType) Reset() {
	*x = JsonContentType{}
	mi := &file_com_coralogix_integrations_v1_webhook_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JsonContentType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JsonContentType) ProtoMessage() {}

func (x *JsonContentType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_integrations_v1_webhook_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JsonContentType.ProtoReflect.Descriptor instead.
func (*JsonContentType) Descriptor() ([]byte, []int) {
	return file_com_coralogix_integrations_v1_webhook_proto_rawDescGZIP(), []int{5}
}

func (x *JsonContentType) GetTimestampSource() isJsonContentType_TimestampSource {
	if x != nil {
		return x.TimestampSource
	}
	return nil
}

func (x *JsonContentType) GetGenerate() *TimestampSourceGenerate {
	if x != nil {
		if x, ok := x.TimestampSource.(*JsonContentType_Generate); ok {
			return x.Generate
		}
	}
	return nil
}

func (x *JsonContentType) GetJsonPath() *TimestampSourceJsonPath {
	if x != nil {
		if x, ok := x.TimestampSource.(*JsonContentType_JsonPath); ok {
			return x.JsonPath
		}
	}
	return nil
}

type isJsonContentType_TimestampSource interface {
	isJsonContentType_TimestampSource()
}

type JsonContentType_Generate struct {
	Generate *TimestampSourceGenerate `protobuf:"bytes,1,opt,name=generate,proto3,oneof"`
}

type JsonContentType_JsonPath struct {
	JsonPath *TimestampSourceJsonPath `protobuf:"bytes,2,opt,name=json_path,json=jsonPath,proto3,oneof"`
}

func (*JsonContentType_Generate) isJsonContentType_TimestampSource() {}

func (*JsonContentType_JsonPath) isJsonContentType_TimestampSource() {}

type TextContentType struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to TimestampSource:
	//
	//	*TextContentType_Generate
	//	*TextContentType_Regex
	TimestampSource isTextContentType_TimestampSource `protobuf_oneof:"timestamp_source"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *TextContentType) Reset() {
	*x = TextContentType{}
	mi := &file_com_coralogix_integrations_v1_webhook_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TextContentType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextContentType) ProtoMessage() {}

func (x *TextContentType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_integrations_v1_webhook_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextContentType.ProtoReflect.Descriptor instead.
func (*TextContentType) Descriptor() ([]byte, []int) {
	return file_com_coralogix_integrations_v1_webhook_proto_rawDescGZIP(), []int{6}
}

func (x *TextContentType) GetTimestampSource() isTextContentType_TimestampSource {
	if x != nil {
		return x.TimestampSource
	}
	return nil
}

func (x *TextContentType) GetGenerate() *TimestampSourceGenerate {
	if x != nil {
		if x, ok := x.TimestampSource.(*TextContentType_Generate); ok {
			return x.Generate
		}
	}
	return nil
}

func (x *TextContentType) GetRegex() *TimestampSourceRegex {
	if x != nil {
		if x, ok := x.TimestampSource.(*TextContentType_Regex); ok {
			return x.Regex
		}
	}
	return nil
}

type isTextContentType_TimestampSource interface {
	isTextContentType_TimestampSource()
}

type TextContentType_Generate struct {
	Generate *TimestampSourceGenerate `protobuf:"bytes,1,opt,name=generate,proto3,oneof"`
}

type TextContentType_Regex struct {
	Regex *TimestampSourceRegex `protobuf:"bytes,2,opt,name=regex,proto3,oneof"`
}

func (*TextContentType_Generate) isTextContentType_TimestampSource() {}

func (*TextContentType_Regex) isTextContentType_TimestampSource() {}

var File_com_coralogix_integrations_v1_webhook_proto protoreflect.FileDescriptor

var file_com_coralogix_integrations_v1_webhook_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x19, 0x0a, 0x17,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x22, 0x1a, 0x0a, 0x18, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x61,
	0x74, 0x69, 0x63, 0x22, 0x4d, 0x0a, 0x15, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x12, 0x34, 0x0a, 0x06,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x22, 0x87, 0x02, 0x0a, 0x17, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4a, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x12, 0x39,
	0x0a, 0x09, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x12, 0x57, 0x0a, 0x09, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x41, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x69, 0x63, 0x48, 0x00, 0x52, 0x09, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x69, 0x63, 0x12, 0x4e, 0x0a, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x18, 0x66, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x46, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x48, 0x00, 0x52, 0x06, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x42, 0x08, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0xfd, 0x01, 0x0a,
	0x14, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x65, 0x67, 0x65, 0x78, 0x12, 0x32, 0x0a, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x12, 0x57, 0x0a, 0x09, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x41, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x69, 0x63, 0x48, 0x00, 0x52, 0x09, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x69, 0x63, 0x12, 0x4e, 0x0a, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x18, 0x66, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x46, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x48, 0x00, 0x52, 0x06, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x42, 0x08, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0xd2, 0x01, 0x0a,
	0x0f, 0x4a, 0x73, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x54, 0x0a, 0x08, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x08, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x55, 0x0a, 0x09, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4a, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x74,
	0x68, 0x48, 0x00, 0x52, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x42, 0x12, 0x0a,
	0x10, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x22, 0xc8, 0x01, 0x0a, 0x0f, 0x54, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x54, 0x0a, 0x08, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x48,
	0x00, 0x52, 0x08, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x4b, 0x0a, 0x05, 0x72,
	0x65, 0x67, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x67, 0x65, 0x78, 0x48,
	0x00, 0x52, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x42, 0x12, 0x0a, 0x10, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_integrations_v1_webhook_proto_rawDescOnce sync.Once
	file_com_coralogix_integrations_v1_webhook_proto_rawDescData = file_com_coralogix_integrations_v1_webhook_proto_rawDesc
)

func file_com_coralogix_integrations_v1_webhook_proto_rawDescGZIP() []byte {
	file_com_coralogix_integrations_v1_webhook_proto_rawDescOnce.Do(func() {
		file_com_coralogix_integrations_v1_webhook_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_integrations_v1_webhook_proto_rawDescData)
	})
	return file_com_coralogix_integrations_v1_webhook_proto_rawDescData
}

var file_com_coralogix_integrations_v1_webhook_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_com_coralogix_integrations_v1_webhook_proto_goTypes = []any{
	(*TimestampSourceGenerate)(nil),  // 0: com.coralogix.integrations.v1.TimestampSourceGenerate
	(*TimestampFormatAutomatic)(nil), // 1: com.coralogix.integrations.v1.TimestampFormatAutomatic
	(*TimestampFormatCustom)(nil),    // 2: com.coralogix.integrations.v1.TimestampFormatCustom
	(*TimestampSourceJsonPath)(nil),  // 3: com.coralogix.integrations.v1.TimestampSourceJsonPath
	(*TimestampSourceRegex)(nil),     // 4: com.coralogix.integrations.v1.TimestampSourceRegex
	(*JsonContentType)(nil),          // 5: com.coralogix.integrations.v1.JsonContentType
	(*TextContentType)(nil),          // 6: com.coralogix.integrations.v1.TextContentType
	(*wrapperspb.StringValue)(nil),   // 7: google.protobuf.StringValue
}
var file_com_coralogix_integrations_v1_webhook_proto_depIdxs = []int32{
	7,  // 0: com.coralogix.integrations.v1.TimestampFormatCustom.format:type_name -> google.protobuf.StringValue
	7,  // 1: com.coralogix.integrations.v1.TimestampSourceJsonPath.json_path:type_name -> google.protobuf.StringValue
	1,  // 2: com.coralogix.integrations.v1.TimestampSourceJsonPath.automatic:type_name -> com.coralogix.integrations.v1.TimestampFormatAutomatic
	2,  // 3: com.coralogix.integrations.v1.TimestampSourceJsonPath.custom:type_name -> com.coralogix.integrations.v1.TimestampFormatCustom
	7,  // 4: com.coralogix.integrations.v1.TimestampSourceRegex.regex:type_name -> google.protobuf.StringValue
	1,  // 5: com.coralogix.integrations.v1.TimestampSourceRegex.automatic:type_name -> com.coralogix.integrations.v1.TimestampFormatAutomatic
	2,  // 6: com.coralogix.integrations.v1.TimestampSourceRegex.custom:type_name -> com.coralogix.integrations.v1.TimestampFormatCustom
	0,  // 7: com.coralogix.integrations.v1.JsonContentType.generate:type_name -> com.coralogix.integrations.v1.TimestampSourceGenerate
	3,  // 8: com.coralogix.integrations.v1.JsonContentType.json_path:type_name -> com.coralogix.integrations.v1.TimestampSourceJsonPath
	0,  // 9: com.coralogix.integrations.v1.TextContentType.generate:type_name -> com.coralogix.integrations.v1.TimestampSourceGenerate
	4,  // 10: com.coralogix.integrations.v1.TextContentType.regex:type_name -> com.coralogix.integrations.v1.TimestampSourceRegex
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_com_coralogix_integrations_v1_webhook_proto_init() }
func file_com_coralogix_integrations_v1_webhook_proto_init() {
	if File_com_coralogix_integrations_v1_webhook_proto != nil {
		return
	}
	file_com_coralogix_integrations_v1_webhook_proto_msgTypes[3].OneofWrappers = []any{
		(*TimestampSourceJsonPath_Automatic)(nil),
		(*TimestampSourceJsonPath_Custom)(nil),
	}
	file_com_coralogix_integrations_v1_webhook_proto_msgTypes[4].OneofWrappers = []any{
		(*TimestampSourceRegex_Automatic)(nil),
		(*TimestampSourceRegex_Custom)(nil),
	}
	file_com_coralogix_integrations_v1_webhook_proto_msgTypes[5].OneofWrappers = []any{
		(*JsonContentType_Generate)(nil),
		(*JsonContentType_JsonPath)(nil),
	}
	file_com_coralogix_integrations_v1_webhook_proto_msgTypes[6].OneofWrappers = []any{
		(*TextContentType_Generate)(nil),
		(*TextContentType_Regex)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_integrations_v1_webhook_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_integrations_v1_webhook_proto_goTypes,
		DependencyIndexes: file_com_coralogix_integrations_v1_webhook_proto_depIdxs,
		MessageInfos:      file_com_coralogix_integrations_v1_webhook_proto_msgTypes,
	}.Build()
	File_com_coralogix_integrations_v1_webhook_proto = out.File
	file_com_coralogix_integrations_v1_webhook_proto_rawDesc = nil
	file_com_coralogix_integrations_v1_webhook_proto_goTypes = nil
	file_com_coralogix_integrations_v1_webhook_proto_depIdxs = nil
}
