// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: com/coralogix/dataprime/v1/feature_control.proto

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

type FeatureControl_Feature int32

const (
	FeatureControl_FEATURE_UNSPECIFIED     FeatureControl_Feature = 0
	FeatureControl_FEATURE_UNKNOWN_FIELDS  FeatureControl_Feature = 1
	FeatureControl_FEATURE_OPTIONAL_SOURCE FeatureControl_Feature = 2
	FeatureControl_FEATURE_INVALID_REGEX   FeatureControl_Feature = 3
	FeatureControl_FEATURE_WILDFIND        FeatureControl_Feature = 4
)

// Enum value maps for FeatureControl_Feature.
var (
	FeatureControl_Feature_name = map[int32]string{
		0: "FEATURE_UNSPECIFIED",
		1: "FEATURE_UNKNOWN_FIELDS",
		2: "FEATURE_OPTIONAL_SOURCE",
		3: "FEATURE_INVALID_REGEX",
		4: "FEATURE_WILDFIND",
	}
	FeatureControl_Feature_value = map[string]int32{
		"FEATURE_UNSPECIFIED":     0,
		"FEATURE_UNKNOWN_FIELDS":  1,
		"FEATURE_OPTIONAL_SOURCE": 2,
		"FEATURE_INVALID_REGEX":   3,
		"FEATURE_WILDFIND":        4,
	}
)

func (x FeatureControl_Feature) Enum() *FeatureControl_Feature {
	p := new(FeatureControl_Feature)
	*p = x
	return p
}

func (x FeatureControl_Feature) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FeatureControl_Feature) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_dataprime_v1_feature_control_proto_enumTypes[0].Descriptor()
}

func (FeatureControl_Feature) Type() protoreflect.EnumType {
	return &file_com_coralogix_dataprime_v1_feature_control_proto_enumTypes[0]
}

func (x FeatureControl_Feature) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FeatureControl_Feature.Descriptor instead.
func (FeatureControl_Feature) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_feature_control_proto_rawDescGZIP(), []int{0, 0}
}

type FeatureControl_Control int32

const (
	FeatureControl_CONTROL_UNSPECIFIED FeatureControl_Control = 0
	FeatureControl_CONTROL_WARNING     FeatureControl_Control = 1
	FeatureControl_CONTROL_ERROR       FeatureControl_Control = 2
)

// Enum value maps for FeatureControl_Control.
var (
	FeatureControl_Control_name = map[int32]string{
		0: "CONTROL_UNSPECIFIED",
		1: "CONTROL_WARNING",
		2: "CONTROL_ERROR",
	}
	FeatureControl_Control_value = map[string]int32{
		"CONTROL_UNSPECIFIED": 0,
		"CONTROL_WARNING":     1,
		"CONTROL_ERROR":       2,
	}
)

func (x FeatureControl_Control) Enum() *FeatureControl_Control {
	p := new(FeatureControl_Control)
	*p = x
	return p
}

func (x FeatureControl_Control) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FeatureControl_Control) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_dataprime_v1_feature_control_proto_enumTypes[1].Descriptor()
}

func (FeatureControl_Control) Type() protoreflect.EnumType {
	return &file_com_coralogix_dataprime_v1_feature_control_proto_enumTypes[1]
}

func (x FeatureControl_Control) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FeatureControl_Control.Descriptor instead.
func (FeatureControl_Control) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_feature_control_proto_rawDescGZIP(), []int{0, 1}
}

type FeatureControl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feature           FeatureControl_Feature  `protobuf:"varint,1,opt,name=feature,proto3,enum=com.coralogix.dataprime.v1.FeatureControl_Feature" json:"feature,omitempty"`
	OnQueryDurationGt *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=on_query_duration_gt,json=onQueryDurationGt,proto3,oneof" json:"on_query_duration_gt,omitempty"`
	Control           *FeatureControl_Control `protobuf:"varint,3,opt,name=control,proto3,enum=com.coralogix.dataprime.v1.FeatureControl_Control,oneof" json:"control,omitempty"`
}

func (x *FeatureControl) Reset() {
	*x = FeatureControl{}
	mi := &file_com_coralogix_dataprime_v1_feature_control_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FeatureControl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureControl) ProtoMessage() {}

func (x *FeatureControl) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_feature_control_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureControl.ProtoReflect.Descriptor instead.
func (*FeatureControl) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_feature_control_proto_rawDescGZIP(), []int{0}
}

func (x *FeatureControl) GetFeature() FeatureControl_Feature {
	if x != nil {
		return x.Feature
	}
	return FeatureControl_FEATURE_UNSPECIFIED
}

func (x *FeatureControl) GetOnQueryDurationGt() *wrapperspb.StringValue {
	if x != nil {
		return x.OnQueryDurationGt
	}
	return nil
}

func (x *FeatureControl) GetControl() FeatureControl_Control {
	if x != nil && x.Control != nil {
		return *x.Control
	}
	return FeatureControl_CONTROL_UNSPECIFIED
}

type ListFeatureControlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListFeatureControlRequest) Reset() {
	*x = ListFeatureControlRequest{}
	mi := &file_com_coralogix_dataprime_v1_feature_control_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFeatureControlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFeatureControlRequest) ProtoMessage() {}

func (x *ListFeatureControlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_feature_control_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFeatureControlRequest.ProtoReflect.Descriptor instead.
func (*ListFeatureControlRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_feature_control_proto_rawDescGZIP(), []int{1}
}

type ListFeatureControlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FeatureControls []*FeatureControl `protobuf:"bytes,1,rep,name=feature_controls,json=featureControls,proto3" json:"feature_controls,omitempty"`
}

func (x *ListFeatureControlResponse) Reset() {
	*x = ListFeatureControlResponse{}
	mi := &file_com_coralogix_dataprime_v1_feature_control_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFeatureControlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFeatureControlResponse) ProtoMessage() {}

func (x *ListFeatureControlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_feature_control_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFeatureControlResponse.ProtoReflect.Descriptor instead.
func (*ListFeatureControlResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_feature_control_proto_rawDescGZIP(), []int{2}
}

func (x *ListFeatureControlResponse) GetFeatureControls() []*FeatureControl {
	if x != nil {
		return x.FeatureControls
	}
	return nil
}

var File_com_coralogix_dataprime_v1_feature_control_proto protoreflect.FileDescriptor

var file_com_coralogix_dataprime_v1_feature_control_proto_rawDesc = []byte{
	0x0a, 0x30, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85,
	0x04, 0x0a, 0x0e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x12, 0x4c, 0x0a, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x52, 0x0a, 0x14, 0x6f, 0x6e, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x11, 0x6f,
	0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x74,
	0x88, 0x01, 0x01, 0x12, 0x51, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x48, 0x01, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x88, 0x01, 0x01, 0x22, 0x8c, 0x01, 0x0a, 0x07, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x46,
	0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x46,
	0x49, 0x45, 0x4c, 0x44, 0x53, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x46, 0x45, 0x41, 0x54, 0x55,
	0x52, 0x45, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x5f, 0x53, 0x4f, 0x55, 0x52,
	0x43, 0x45, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x52, 0x45, 0x47, 0x45, 0x58, 0x10, 0x03, 0x12,
	0x14, 0x0a, 0x10, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x57, 0x49, 0x4c, 0x44, 0x46,
	0x49, 0x4e, 0x44, 0x10, 0x04, 0x22, 0x4a, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4f, 0x4e,
	0x54, 0x52, 0x4f, 0x4c, 0x5f, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x11,
	0x0a, 0x0d, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x02, 0x42, 0x17, 0x0a, 0x15, 0x5f, 0x6f, 0x6e, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x22, 0x1b, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x73, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x55, 0x0a, 0x10, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x0f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_dataprime_v1_feature_control_proto_rawDescOnce sync.Once
	file_com_coralogix_dataprime_v1_feature_control_proto_rawDescData = file_com_coralogix_dataprime_v1_feature_control_proto_rawDesc
)

func file_com_coralogix_dataprime_v1_feature_control_proto_rawDescGZIP() []byte {
	file_com_coralogix_dataprime_v1_feature_control_proto_rawDescOnce.Do(func() {
		file_com_coralogix_dataprime_v1_feature_control_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_dataprime_v1_feature_control_proto_rawDescData)
	})
	return file_com_coralogix_dataprime_v1_feature_control_proto_rawDescData
}

var file_com_coralogix_dataprime_v1_feature_control_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_com_coralogix_dataprime_v1_feature_control_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogix_dataprime_v1_feature_control_proto_goTypes = []any{
	(FeatureControl_Feature)(0),        // 0: com.coralogix.dataprime.v1.FeatureControl.Feature
	(FeatureControl_Control)(0),        // 1: com.coralogix.dataprime.v1.FeatureControl.Control
	(*FeatureControl)(nil),             // 2: com.coralogix.dataprime.v1.FeatureControl
	(*ListFeatureControlRequest)(nil),  // 3: com.coralogix.dataprime.v1.ListFeatureControlRequest
	(*ListFeatureControlResponse)(nil), // 4: com.coralogix.dataprime.v1.ListFeatureControlResponse
	(*wrapperspb.StringValue)(nil),     // 5: google.protobuf.StringValue
}
var file_com_coralogix_dataprime_v1_feature_control_proto_depIdxs = []int32{
	0, // 0: com.coralogix.dataprime.v1.FeatureControl.feature:type_name -> com.coralogix.dataprime.v1.FeatureControl.Feature
	5, // 1: com.coralogix.dataprime.v1.FeatureControl.on_query_duration_gt:type_name -> google.protobuf.StringValue
	1, // 2: com.coralogix.dataprime.v1.FeatureControl.control:type_name -> com.coralogix.dataprime.v1.FeatureControl.Control
	2, // 3: com.coralogix.dataprime.v1.ListFeatureControlResponse.feature_controls:type_name -> com.coralogix.dataprime.v1.FeatureControl
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_com_coralogix_dataprime_v1_feature_control_proto_init() }
func file_com_coralogix_dataprime_v1_feature_control_proto_init() {
	if File_com_coralogix_dataprime_v1_feature_control_proto != nil {
		return
	}
	file_com_coralogix_dataprime_v1_feature_control_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_dataprime_v1_feature_control_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_dataprime_v1_feature_control_proto_goTypes,
		DependencyIndexes: file_com_coralogix_dataprime_v1_feature_control_proto_depIdxs,
		EnumInfos:         file_com_coralogix_dataprime_v1_feature_control_proto_enumTypes,
		MessageInfos:      file_com_coralogix_dataprime_v1_feature_control_proto_msgTypes,
	}.Build()
	File_com_coralogix_dataprime_v1_feature_control_proto = out.File
	file_com_coralogix_dataprime_v1_feature_control_proto_rawDesc = nil
	file_com_coralogix_dataprime_v1_feature_control_proto_goTypes = nil
	file_com_coralogix_dataprime_v1_feature_control_proto_depIdxs = nil
}
