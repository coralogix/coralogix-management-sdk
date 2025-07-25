// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/alerts/v3/pagination.proto

package v3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type PaginationRequest struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	PageSize      *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken     *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaginationRequest) Reset() {
	*x = PaginationRequest{}
	mi := &file_com_coralogixapis_alerts_v3_pagination_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaginationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationRequest) ProtoMessage() {}

func (x *PaginationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_pagination_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationRequest.ProtoReflect.Descriptor instead.
func (*PaginationRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_pagination_proto_rawDescGZIP(), []int{0}
}

func (x *PaginationRequest) GetPageSize() *wrapperspb.UInt32Value {
	if x != nil {
		return x.PageSize
	}
	return nil
}

func (x *PaginationRequest) GetPageToken() *wrapperspb.StringValue {
	if x != nil {
		return x.PageToken
	}
	return nil
}

type PaginationResponse struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	TotalSize     *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	NextPageToken *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaginationResponse) Reset() {
	*x = PaginationResponse{}
	mi := &file_com_coralogixapis_alerts_v3_pagination_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationResponse) ProtoMessage() {}

func (x *PaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_pagination_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationResponse.ProtoReflect.Descriptor instead.
func (*PaginationResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_pagination_proto_rawDescGZIP(), []int{1}
}

func (x *PaginationResponse) GetTotalSize() *wrapperspb.UInt32Value {
	if x != nil {
		return x.TotalSize
	}
	return nil
}

func (x *PaginationResponse) GetNextPageToken() *wrapperspb.StringValue {
	if x != nil {
		return x.NextPageToken
	}
	return nil
}

var File_com_coralogixapis_alerts_v3_pagination_proto protoreflect.FileDescriptor

const file_com_coralogixapis_alerts_v3_pagination_proto_rawDesc = "" +
	"\n" +
	",com/coralogixapis/alerts/v3/pagination.proto\x12\x1bcom.coralogixapis.alerts.v3\x1a\x1egoogle/protobuf/wrappers.proto\"\x8b\x01\n" +
	"\x11PaginationRequest\x129\n" +
	"\tpage_size\x18\x01 \x01(\v2\x1c.google.protobuf.UInt32ValueR\bpageSize\x12;\n" +
	"\n" +
	"page_token\x18\x02 \x01(\v2\x1c.google.protobuf.StringValueR\tpageToken\"\x97\x01\n" +
	"\x12PaginationResponse\x12;\n" +
	"\n" +
	"total_size\x18\x01 \x01(\v2\x1c.google.protobuf.UInt32ValueR\ttotalSize\x12D\n" +
	"\x0fnext_page_token\x18\x02 \x01(\v2\x1c.google.protobuf.StringValueR\rnextPageTokenb\x06proto3"

var (
	file_com_coralogixapis_alerts_v3_pagination_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_pagination_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v3_pagination_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_pagination_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_pagination_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_pagination_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_pagination_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerts_v3_pagination_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_pagination_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogixapis_alerts_v3_pagination_proto_goTypes = []any{
	(*PaginationRequest)(nil),      // 0: com.coralogixapis.alerts.v3.PaginationRequest
	(*PaginationResponse)(nil),     // 1: com.coralogixapis.alerts.v3.PaginationResponse
	(*wrapperspb.UInt32Value)(nil), // 2: google.protobuf.UInt32Value
	(*wrapperspb.StringValue)(nil), // 3: google.protobuf.StringValue
}
var file_com_coralogixapis_alerts_v3_pagination_proto_depIdxs = []int32{
	2, // 0: com.coralogixapis.alerts.v3.PaginationRequest.page_size:type_name -> google.protobuf.UInt32Value
	3, // 1: com.coralogixapis.alerts.v3.PaginationRequest.page_token:type_name -> google.protobuf.StringValue
	2, // 2: com.coralogixapis.alerts.v3.PaginationResponse.total_size:type_name -> google.protobuf.UInt32Value
	3, // 3: com.coralogixapis.alerts.v3.PaginationResponse.next_page_token:type_name -> google.protobuf.StringValue
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_alerts_v3_pagination_proto_init() }
func file_com_coralogixapis_alerts_v3_pagination_proto_init() {
	if File_com_coralogixapis_alerts_v3_pagination_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_pagination_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_pagination_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_pagination_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_pagination_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_pagination_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_pagination_proto = out.File
	file_com_coralogixapis_alerts_v3_pagination_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_pagination_proto_depIdxs = nil
}
