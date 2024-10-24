// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: com/coralogixapis/views/v1/time_selection.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type TimeSelection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to SelectionType:
	//
	//	*TimeSelection_QuickSelection
	//	*TimeSelection_CustomSelection
	SelectionType isTimeSelection_SelectionType `protobuf_oneof:"selection_type"`
}

func (x *TimeSelection) Reset() {
	*x = TimeSelection{}
	mi := &file_com_coralogixapis_views_v1_time_selection_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimeSelection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeSelection) ProtoMessage() {}

func (x *TimeSelection) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_views_v1_time_selection_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeSelection.ProtoReflect.Descriptor instead.
func (*TimeSelection) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_views_v1_time_selection_proto_rawDescGZIP(), []int{0}
}

func (m *TimeSelection) GetSelectionType() isTimeSelection_SelectionType {
	if m != nil {
		return m.SelectionType
	}
	return nil
}

func (x *TimeSelection) GetQuickSelection() *QuickTimeSelection {
	if x, ok := x.GetSelectionType().(*TimeSelection_QuickSelection); ok {
		return x.QuickSelection
	}
	return nil
}

func (x *TimeSelection) GetCustomSelection() *CustomTimeSelection {
	if x, ok := x.GetSelectionType().(*TimeSelection_CustomSelection); ok {
		return x.CustomSelection
	}
	return nil
}

type isTimeSelection_SelectionType interface {
	isTimeSelection_SelectionType()
}

type TimeSelection_QuickSelection struct {
	QuickSelection *QuickTimeSelection `protobuf:"bytes,1,opt,name=quick_selection,json=quickSelection,proto3,oneof"`
}

type TimeSelection_CustomSelection struct {
	CustomSelection *CustomTimeSelection `protobuf:"bytes,2,opt,name=custom_selection,json=customSelection,proto3,oneof"`
}

func (*TimeSelection_QuickSelection) isTimeSelection_SelectionType() {}

func (*TimeSelection_CustomSelection) isTimeSelection_SelectionType() {}

type QuickTimeSelection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Caption *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=caption,proto3" json:"caption,omitempty"`
	Seconds uint32                  `protobuf:"varint,3,opt,name=seconds,proto3" json:"seconds,omitempty"`
}

func (x *QuickTimeSelection) Reset() {
	*x = QuickTimeSelection{}
	mi := &file_com_coralogixapis_views_v1_time_selection_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QuickTimeSelection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuickTimeSelection) ProtoMessage() {}

func (x *QuickTimeSelection) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_views_v1_time_selection_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuickTimeSelection.ProtoReflect.Descriptor instead.
func (*QuickTimeSelection) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_views_v1_time_selection_proto_rawDescGZIP(), []int{1}
}

func (x *QuickTimeSelection) GetCaption() *wrapperspb.StringValue {
	if x != nil {
		return x.Caption
	}
	return nil
}

func (x *QuickTimeSelection) GetSeconds() uint32 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

type CustomTimeSelection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=from_time,json=fromTime,proto3" json:"from_time,omitempty"`
	ToTime   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=to_time,json=toTime,proto3" json:"to_time,omitempty"`
}

func (x *CustomTimeSelection) Reset() {
	*x = CustomTimeSelection{}
	mi := &file_com_coralogixapis_views_v1_time_selection_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomTimeSelection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomTimeSelection) ProtoMessage() {}

func (x *CustomTimeSelection) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_views_v1_time_selection_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomTimeSelection.ProtoReflect.Descriptor instead.
func (*CustomTimeSelection) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_views_v1_time_selection_proto_rawDescGZIP(), []int{2}
}

func (x *CustomTimeSelection) GetFromTime() *timestamppb.Timestamp {
	if x != nil {
		return x.FromTime
	}
	return nil
}

func (x *CustomTimeSelection) GetToTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ToTime
	}
	return nil
}

var File_com_coralogixapis_views_v1_time_selection_proto protoreflect.FileDescriptor

var file_com_coralogixapis_views_v1_time_selection_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda,
	0x01, 0x0a, 0x0d, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x59, 0x0a, 0x0f, 0x71, 0x75, 0x69, 0x63, 0x6b, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x69,
	0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x69, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0e, 0x71, 0x75, 0x69,
	0x63, 0x6b, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5c, 0x0a, 0x10, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x10, 0x0a, 0x0e, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x66, 0x0a, 0x12, 0x51,
	0x75, 0x69, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x36, 0x0a, 0x07, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x07, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x22, 0x83, 0x01, 0x0a, 0x13, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x69,
	0x6d, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x09, 0x66,
	0x72, 0x6f, 0x6d, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x06, 0x74, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_com_coralogixapis_views_v1_time_selection_proto_rawDescOnce sync.Once
	file_com_coralogixapis_views_v1_time_selection_proto_rawDescData = file_com_coralogixapis_views_v1_time_selection_proto_rawDesc
)

func file_com_coralogixapis_views_v1_time_selection_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_views_v1_time_selection_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_views_v1_time_selection_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_views_v1_time_selection_proto_rawDescData)
	})
	return file_com_coralogixapis_views_v1_time_selection_proto_rawDescData
}

var file_com_coralogixapis_views_v1_time_selection_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogixapis_views_v1_time_selection_proto_goTypes = []any{
	(*TimeSelection)(nil),          // 0: com.coralogixapis.views.v1.TimeSelection
	(*QuickTimeSelection)(nil),     // 1: com.coralogixapis.views.v1.QuickTimeSelection
	(*CustomTimeSelection)(nil),    // 2: com.coralogixapis.views.v1.CustomTimeSelection
	(*wrapperspb.StringValue)(nil), // 3: google.protobuf.StringValue
	(*timestamppb.Timestamp)(nil),  // 4: google.protobuf.Timestamp
}
var file_com_coralogixapis_views_v1_time_selection_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.views.v1.TimeSelection.quick_selection:type_name -> com.coralogixapis.views.v1.QuickTimeSelection
	2, // 1: com.coralogixapis.views.v1.TimeSelection.custom_selection:type_name -> com.coralogixapis.views.v1.CustomTimeSelection
	3, // 2: com.coralogixapis.views.v1.QuickTimeSelection.caption:type_name -> google.protobuf.StringValue
	4, // 3: com.coralogixapis.views.v1.CustomTimeSelection.from_time:type_name -> google.protobuf.Timestamp
	4, // 4: com.coralogixapis.views.v1.CustomTimeSelection.to_time:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_views_v1_time_selection_proto_init() }
func file_com_coralogixapis_views_v1_time_selection_proto_init() {
	if File_com_coralogixapis_views_v1_time_selection_proto != nil {
		return
	}
	file_com_coralogixapis_views_v1_time_selection_proto_msgTypes[0].OneofWrappers = []any{
		(*TimeSelection_QuickSelection)(nil),
		(*TimeSelection_CustomSelection)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_views_v1_time_selection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_views_v1_time_selection_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_views_v1_time_selection_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_views_v1_time_selection_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_views_v1_time_selection_proto = out.File
	file_com_coralogixapis_views_v1_time_selection_proto_rawDesc = nil
	file_com_coralogixapis_views_v1_time_selection_proto_goTypes = nil
	file_com_coralogixapis_views_v1_time_selection_proto_depIdxs = nil
}
