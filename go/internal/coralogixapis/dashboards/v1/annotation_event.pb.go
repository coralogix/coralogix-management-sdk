// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: com/coralogixapis/dashboards/v1/common/annotation_event.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

type AnnotationEvent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Value:
	//
	//	*AnnotationEvent_Instant_
	//	*AnnotationEvent_Range_
	Value         isAnnotationEvent_Value `protobuf_oneof:"value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AnnotationEvent) Reset() {
	*x = AnnotationEvent{}
	mi := &file_com_coralogixapis_dashboards_v1_common_annotation_event_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnnotationEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnotationEvent) ProtoMessage() {}

func (x *AnnotationEvent) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_common_annotation_event_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnotationEvent.ProtoReflect.Descriptor instead.
func (*AnnotationEvent) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_annotation_event_proto_rawDescGZIP(), []int{0}
}

func (x *AnnotationEvent) GetValue() isAnnotationEvent_Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *AnnotationEvent) GetInstant() *AnnotationEvent_Instant {
	if x != nil {
		if x, ok := x.Value.(*AnnotationEvent_Instant_); ok {
			return x.Instant
		}
	}
	return nil
}

func (x *AnnotationEvent) GetRange() *AnnotationEvent_Range {
	if x != nil {
		if x, ok := x.Value.(*AnnotationEvent_Range_); ok {
			return x.Range
		}
	}
	return nil
}

type isAnnotationEvent_Value interface {
	isAnnotationEvent_Value()
}

type AnnotationEvent_Instant_ struct {
	Instant *AnnotationEvent_Instant `protobuf:"bytes,1,opt,name=instant,proto3,oneof"`
}

type AnnotationEvent_Range_ struct {
	Range *AnnotationEvent_Range `protobuf:"bytes,2,opt,name=range,proto3,oneof"`
}

func (*AnnotationEvent_Instant_) isAnnotationEvent_Value() {}

func (*AnnotationEvent_Range_) isAnnotationEvent_Value() {}

type AnnotationEvent_Instant struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Labels        map[string]string      `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Payload       *structpb.Struct       `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AnnotationEvent_Instant) Reset() {
	*x = AnnotationEvent_Instant{}
	mi := &file_com_coralogixapis_dashboards_v1_common_annotation_event_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnnotationEvent_Instant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnotationEvent_Instant) ProtoMessage() {}

func (x *AnnotationEvent_Instant) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_common_annotation_event_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnotationEvent_Instant.ProtoReflect.Descriptor instead.
func (*AnnotationEvent_Instant) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_annotation_event_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AnnotationEvent_Instant) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *AnnotationEvent_Instant) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *AnnotationEvent_Instant) GetPayload() *structpb.Struct {
	if x != nil {
		return x.Payload
	}
	return nil
}

type AnnotationEvent_Range struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Start         *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End           *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	Labels        map[string]string      `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Payload       *structpb.Struct       `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AnnotationEvent_Range) Reset() {
	*x = AnnotationEvent_Range{}
	mi := &file_com_coralogixapis_dashboards_v1_common_annotation_event_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnnotationEvent_Range) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnotationEvent_Range) ProtoMessage() {}

func (x *AnnotationEvent_Range) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_common_annotation_event_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnotationEvent_Range.ProtoReflect.Descriptor instead.
func (*AnnotationEvent_Range) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_annotation_event_proto_rawDescGZIP(), []int{0, 1}
}

func (x *AnnotationEvent_Range) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *AnnotationEvent_Range) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

func (x *AnnotationEvent_Range) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *AnnotationEvent_Range) GetPayload() *structpb.Struct {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_com_coralogixapis_dashboards_v1_common_annotation_event_proto protoreflect.FileDescriptor

var file_com_coralogixapis_dashboards_v1_common_annotation_event_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x06, 0x0a, 0x0f, 0x41, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x5b, 0x0a, 0x07, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x07,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x12, 0x55, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x00, 0x52, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x1a, 0x96,
	0x02, 0x0a, 0x07, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x63, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x31, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x39, 0x0a, 0x0b,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0xb8, 0x02, 0x0a, 0x05, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x2c, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x65, 0x6e,
	0x64, 0x12, 0x61, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x49, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x31, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_dashboards_v1_common_annotation_event_proto_rawDescOnce sync.Once
	file_com_coralogixapis_dashboards_v1_common_annotation_event_proto_rawDescData = file_com_coralogixapis_dashboards_v1_common_annotation_event_proto_rawDesc
)

func file_com_coralogixapis_dashboards_v1_common_annotation_event_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_dashboards_v1_common_annotation_event_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_dashboards_v1_common_annotation_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_dashboards_v1_common_annotation_event_proto_rawDescData)
	})
	return file_com_coralogixapis_dashboards_v1_common_annotation_event_proto_rawDescData
}

var file_com_coralogixapis_dashboards_v1_common_annotation_event_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_com_coralogixapis_dashboards_v1_common_annotation_event_proto_goTypes = []any{
	(*AnnotationEvent)(nil),         // 0: com.coralogixapis.dashboards.v1.common.AnnotationEvent
	(*AnnotationEvent_Instant)(nil), // 1: com.coralogixapis.dashboards.v1.common.AnnotationEvent.Instant
	(*AnnotationEvent_Range)(nil),   // 2: com.coralogixapis.dashboards.v1.common.AnnotationEvent.Range
	nil,                             // 3: com.coralogixapis.dashboards.v1.common.AnnotationEvent.Instant.LabelsEntry
	nil,                             // 4: com.coralogixapis.dashboards.v1.common.AnnotationEvent.Range.LabelsEntry
	(*timestamppb.Timestamp)(nil),   // 5: google.protobuf.Timestamp
	(*structpb.Struct)(nil),         // 6: google.protobuf.Struct
}
var file_com_coralogixapis_dashboards_v1_common_annotation_event_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.dashboards.v1.common.AnnotationEvent.instant:type_name -> com.coralogixapis.dashboards.v1.common.AnnotationEvent.Instant
	2, // 1: com.coralogixapis.dashboards.v1.common.AnnotationEvent.range:type_name -> com.coralogixapis.dashboards.v1.common.AnnotationEvent.Range
	5, // 2: com.coralogixapis.dashboards.v1.common.AnnotationEvent.Instant.timestamp:type_name -> google.protobuf.Timestamp
	3, // 3: com.coralogixapis.dashboards.v1.common.AnnotationEvent.Instant.labels:type_name -> com.coralogixapis.dashboards.v1.common.AnnotationEvent.Instant.LabelsEntry
	6, // 4: com.coralogixapis.dashboards.v1.common.AnnotationEvent.Instant.payload:type_name -> google.protobuf.Struct
	5, // 5: com.coralogixapis.dashboards.v1.common.AnnotationEvent.Range.start:type_name -> google.protobuf.Timestamp
	5, // 6: com.coralogixapis.dashboards.v1.common.AnnotationEvent.Range.end:type_name -> google.protobuf.Timestamp
	4, // 7: com.coralogixapis.dashboards.v1.common.AnnotationEvent.Range.labels:type_name -> com.coralogixapis.dashboards.v1.common.AnnotationEvent.Range.LabelsEntry
	6, // 8: com.coralogixapis.dashboards.v1.common.AnnotationEvent.Range.payload:type_name -> google.protobuf.Struct
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_dashboards_v1_common_annotation_event_proto_init() }
func file_com_coralogixapis_dashboards_v1_common_annotation_event_proto_init() {
	if File_com_coralogixapis_dashboards_v1_common_annotation_event_proto != nil {
		return
	}
	file_com_coralogixapis_dashboards_v1_common_annotation_event_proto_msgTypes[0].OneofWrappers = []any{
		(*AnnotationEvent_Instant_)(nil),
		(*AnnotationEvent_Range_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_dashboards_v1_common_annotation_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_dashboards_v1_common_annotation_event_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_dashboards_v1_common_annotation_event_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_dashboards_v1_common_annotation_event_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_dashboards_v1_common_annotation_event_proto = out.File
	file_com_coralogixapis_dashboards_v1_common_annotation_event_proto_rawDesc = nil
	file_com_coralogixapis_dashboards_v1_common_annotation_event_proto_goTypes = nil
	file_com_coralogixapis_dashboards_v1_common_annotation_event_proto_depIdxs = nil
}
