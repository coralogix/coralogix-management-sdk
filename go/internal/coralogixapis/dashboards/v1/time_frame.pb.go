// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.28.1
// source: com/coralogixapis/dashboards/v1/common/time_frame.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TimeFrameSelect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*TimeFrameSelect_AbsoluteTimeFrame
	//	*TimeFrameSelect_RelativeTimeFrame
	Value isTimeFrameSelect_Value `protobuf_oneof:"value"`
}

func (x *TimeFrameSelect) Reset() {
	*x = TimeFrameSelect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_dashboards_v1_common_time_frame_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeFrameSelect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeFrameSelect) ProtoMessage() {}

func (x *TimeFrameSelect) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_common_time_frame_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeFrameSelect.ProtoReflect.Descriptor instead.
func (*TimeFrameSelect) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_time_frame_proto_rawDescGZIP(), []int{0}
}

func (m *TimeFrameSelect) GetValue() isTimeFrameSelect_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *TimeFrameSelect) GetAbsoluteTimeFrame() *TimeFrame {
	if x, ok := x.GetValue().(*TimeFrameSelect_AbsoluteTimeFrame); ok {
		return x.AbsoluteTimeFrame
	}
	return nil
}

func (x *TimeFrameSelect) GetRelativeTimeFrame() *durationpb.Duration {
	if x, ok := x.GetValue().(*TimeFrameSelect_RelativeTimeFrame); ok {
		return x.RelativeTimeFrame
	}
	return nil
}

type isTimeFrameSelect_Value interface {
	isTimeFrameSelect_Value()
}

type TimeFrameSelect_AbsoluteTimeFrame struct {
	AbsoluteTimeFrame *TimeFrame `protobuf:"bytes,1,opt,name=absolute_time_frame,json=absoluteTimeFrame,proto3,oneof"`
}

type TimeFrameSelect_RelativeTimeFrame struct {
	RelativeTimeFrame *durationpb.Duration `protobuf:"bytes,2,opt,name=relative_time_frame,json=relativeTimeFrame,proto3,oneof"`
}

func (*TimeFrameSelect_AbsoluteTimeFrame) isTimeFrameSelect_Value() {}

func (*TimeFrameSelect_RelativeTimeFrame) isTimeFrameSelect_Value() {}

type TimeFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *TimeFrame) Reset() {
	*x = TimeFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_dashboards_v1_common_time_frame_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeFrame) ProtoMessage() {}

func (x *TimeFrame) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_common_time_frame_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeFrame.ProtoReflect.Descriptor instead.
func (*TimeFrame) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_time_frame_proto_rawDescGZIP(), []int{1}
}

func (x *TimeFrame) GetFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *TimeFrame) GetTo() *timestamppb.Timestamp {
	if x != nil {
		return x.To
	}
	return nil
}

var File_com_coralogixapis_dashboards_v1_common_time_frame_proto protoreflect.FileDescriptor

var file_com_coralogixapis_dashboards_v1_common_time_frame_proto_rawDesc = []byte{
	0x0a, 0x37, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x66, 0x72,
	0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xcc, 0x01, 0x0a, 0x0f, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x63, 0x0a, 0x13, 0x61, 0x62, 0x73, 0x6f, 0x6c, 0x75,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x48, 0x00, 0x52, 0x11, 0x61, 0x62, 0x73, 0x6f, 0x6c, 0x75,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x4b, 0x0a, 0x13, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x66, 0x72, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x11, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x67, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x2e,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x2a,
	0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_com_coralogixapis_dashboards_v1_common_time_frame_proto_rawDescOnce sync.Once
	file_com_coralogixapis_dashboards_v1_common_time_frame_proto_rawDescData = file_com_coralogixapis_dashboards_v1_common_time_frame_proto_rawDesc
)

func file_com_coralogixapis_dashboards_v1_common_time_frame_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_dashboards_v1_common_time_frame_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_dashboards_v1_common_time_frame_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_dashboards_v1_common_time_frame_proto_rawDescData)
	})
	return file_com_coralogixapis_dashboards_v1_common_time_frame_proto_rawDescData
}

var file_com_coralogixapis_dashboards_v1_common_time_frame_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogixapis_dashboards_v1_common_time_frame_proto_goTypes = []interface{}{
	(*TimeFrameSelect)(nil),       // 0: com.coralogixapis.dashboards.v1.common.TimeFrameSelect
	(*TimeFrame)(nil),             // 1: com.coralogixapis.dashboards.v1.common.TimeFrame
	(*durationpb.Duration)(nil),   // 2: google.protobuf.Duration
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_com_coralogixapis_dashboards_v1_common_time_frame_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.dashboards.v1.common.TimeFrameSelect.absolute_time_frame:type_name -> com.coralogixapis.dashboards.v1.common.TimeFrame
	2, // 1: com.coralogixapis.dashboards.v1.common.TimeFrameSelect.relative_time_frame:type_name -> google.protobuf.Duration
	3, // 2: com.coralogixapis.dashboards.v1.common.TimeFrame.from:type_name -> google.protobuf.Timestamp
	3, // 3: com.coralogixapis.dashboards.v1.common.TimeFrame.to:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_dashboards_v1_common_time_frame_proto_init() }
func file_com_coralogixapis_dashboards_v1_common_time_frame_proto_init() {
	if File_com_coralogixapis_dashboards_v1_common_time_frame_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_dashboards_v1_common_time_frame_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeFrameSelect); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_com_coralogixapis_dashboards_v1_common_time_frame_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeFrame); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_com_coralogixapis_dashboards_v1_common_time_frame_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TimeFrameSelect_AbsoluteTimeFrame)(nil),
		(*TimeFrameSelect_RelativeTimeFrame)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_dashboards_v1_common_time_frame_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_dashboards_v1_common_time_frame_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_dashboards_v1_common_time_frame_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_dashboards_v1_common_time_frame_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_dashboards_v1_common_time_frame_proto = out.File
	file_com_coralogixapis_dashboards_v1_common_time_frame_proto_rawDesc = nil
	file_com_coralogixapis_dashboards_v1_common_time_frame_proto_goTypes = nil
	file_com_coralogixapis_dashboards_v1_common_time_frame_proto_depIdxs = nil
}
