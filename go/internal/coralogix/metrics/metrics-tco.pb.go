// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogix/metrics/metrics-tco.proto

package metrics

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type Rule struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Label         string                 `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Regex         string                 `protobuf:"bytes,3,opt,name=regex,proto3" json:"regex,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Rule) Reset() {
	*x = Rule{}
	mi := &file_com_coralogix_metrics_metrics_tco_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Rule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rule) ProtoMessage() {}

func (x *Rule) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_metrics_metrics_tco_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rule.ProtoReflect.Descriptor instead.
func (*Rule) Descriptor() ([]byte, []int) {
	return file_com_coralogix_metrics_metrics_tco_proto_rawDescGZIP(), []int{0}
}

func (x *Rule) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Rule) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Rule) GetRegex() string {
	if x != nil {
		return x.Regex
	}
	return ""
}

type AddRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Label         string                 `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Regex         string                 `protobuf:"bytes,2,opt,name=regex,proto3" json:"regex,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	mi := &file_com_coralogix_metrics_metrics_tco_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_metrics_metrics_tco_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_metrics_metrics_tco_proto_rawDescGZIP(), []int{1}
}

func (x *AddRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *AddRequest) GetRegex() string {
	if x != nil {
		return x.Regex
	}
	return ""
}

type DeleteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RuleId        uint64                 `protobuf:"varint,1,opt,name=rule_id,json=ruleId,proto3" json:"rule_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	mi := &file_com_coralogix_metrics_metrics_tco_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_metrics_metrics_tco_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_metrics_metrics_tco_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteRequest) GetRuleId() uint64 {
	if x != nil {
		return x.RuleId
	}
	return 0
}

type GetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Rules         []*Rule                `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	mi := &file_com_coralogix_metrics_metrics_tco_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_metrics_metrics_tco_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_metrics_metrics_tco_proto_rawDescGZIP(), []int{3}
}

func (x *GetResponse) GetRules() []*Rule {
	if x != nil {
		return x.Rules
	}
	return nil
}

var File_com_coralogix_metrics_metrics_tco_proto protoreflect.FileDescriptor

const file_com_coralogix_metrics_metrics_tco_proto_rawDesc = "" +
	"\n" +
	"'com/coralogix/metrics/metrics-tco.proto\x12!com.coralogix.metrics.metrics_tco\x1a\x1bgoogle/protobuf/empty.proto\"B\n" +
	"\x04Rule\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\x12\x14\n" +
	"\x05label\x18\x02 \x01(\tR\x05label\x12\x14\n" +
	"\x05regex\x18\x03 \x01(\tR\x05regex\"8\n" +
	"\n" +
	"AddRequest\x12\x14\n" +
	"\x05label\x18\x01 \x01(\tR\x05label\x12\x14\n" +
	"\x05regex\x18\x02 \x01(\tR\x05regex\"(\n" +
	"\rDeleteRequest\x12\x17\n" +
	"\arule_id\x18\x01 \x01(\x04R\x06ruleId\"L\n" +
	"\vGetResponse\x12=\n" +
	"\x05rules\x18\x01 \x03(\v2'.com.coralogix.metrics.metrics_tco.RuleR\x05rules2\x8a\x02\n" +
	"\x11MetricsTcoService\x12N\n" +
	"\x03Add\x12-.com.coralogix.metrics.metrics_tco.AddRequest\x1a\x16.google.protobuf.Empty\"\x00\x12T\n" +
	"\x06Delete\x120.com.coralogix.metrics.metrics_tco.DeleteRequest\x1a\x16.google.protobuf.Empty\"\x00\x12O\n" +
	"\x03Get\x12\x16.google.protobuf.Empty\x1a..com.coralogix.metrics.metrics_tco.GetResponse\"\x00b\x06proto3"

var (
	file_com_coralogix_metrics_metrics_tco_proto_rawDescOnce sync.Once
	file_com_coralogix_metrics_metrics_tco_proto_rawDescData []byte
)

func file_com_coralogix_metrics_metrics_tco_proto_rawDescGZIP() []byte {
	file_com_coralogix_metrics_metrics_tco_proto_rawDescOnce.Do(func() {
		file_com_coralogix_metrics_metrics_tco_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogix_metrics_metrics_tco_proto_rawDesc), len(file_com_coralogix_metrics_metrics_tco_proto_rawDesc)))
	})
	return file_com_coralogix_metrics_metrics_tco_proto_rawDescData
}

var file_com_coralogix_metrics_metrics_tco_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_com_coralogix_metrics_metrics_tco_proto_goTypes = []any{
	(*Rule)(nil),          // 0: com.coralogix.metrics.metrics_tco.Rule
	(*AddRequest)(nil),    // 1: com.coralogix.metrics.metrics_tco.AddRequest
	(*DeleteRequest)(nil), // 2: com.coralogix.metrics.metrics_tco.DeleteRequest
	(*GetResponse)(nil),   // 3: com.coralogix.metrics.metrics_tco.GetResponse
	(*emptypb.Empty)(nil), // 4: google.protobuf.Empty
}
var file_com_coralogix_metrics_metrics_tco_proto_depIdxs = []int32{
	0, // 0: com.coralogix.metrics.metrics_tco.GetResponse.rules:type_name -> com.coralogix.metrics.metrics_tco.Rule
	1, // 1: com.coralogix.metrics.metrics_tco.MetricsTcoService.Add:input_type -> com.coralogix.metrics.metrics_tco.AddRequest
	2, // 2: com.coralogix.metrics.metrics_tco.MetricsTcoService.Delete:input_type -> com.coralogix.metrics.metrics_tco.DeleteRequest
	4, // 3: com.coralogix.metrics.metrics_tco.MetricsTcoService.Get:input_type -> google.protobuf.Empty
	4, // 4: com.coralogix.metrics.metrics_tco.MetricsTcoService.Add:output_type -> google.protobuf.Empty
	4, // 5: com.coralogix.metrics.metrics_tco.MetricsTcoService.Delete:output_type -> google.protobuf.Empty
	3, // 6: com.coralogix.metrics.metrics_tco.MetricsTcoService.Get:output_type -> com.coralogix.metrics.metrics_tco.GetResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_com_coralogix_metrics_metrics_tco_proto_init() }
func file_com_coralogix_metrics_metrics_tco_proto_init() {
	if File_com_coralogix_metrics_metrics_tco_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogix_metrics_metrics_tco_proto_rawDesc), len(file_com_coralogix_metrics_metrics_tco_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogix_metrics_metrics_tco_proto_goTypes,
		DependencyIndexes: file_com_coralogix_metrics_metrics_tco_proto_depIdxs,
		MessageInfos:      file_com_coralogix_metrics_metrics_tco_proto_msgTypes,
	}.Build()
	File_com_coralogix_metrics_metrics_tco_proto = out.File
	file_com_coralogix_metrics_metrics_tco_proto_goTypes = nil
	file_com_coralogix_metrics_metrics_tco_proto_depIdxs = nil
}
