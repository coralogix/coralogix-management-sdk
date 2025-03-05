// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: com/coralogixapis/alerts/v3/event/payload/ratio_special_values.proto

package v3

import (
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

type SpecialRatioValues int32

const (
	SpecialRatioValues_SPECIAL_RATIO_VALUES_INFINITY_OR_UNSPECIFIED SpecialRatioValues = 0
)

// Enum value maps for SpecialRatioValues.
var (
	SpecialRatioValues_name = map[int32]string{
		0: "SPECIAL_RATIO_VALUES_INFINITY_OR_UNSPECIFIED",
	}
	SpecialRatioValues_value = map[string]int32{
		"SPECIAL_RATIO_VALUES_INFINITY_OR_UNSPECIFIED": 0,
	}
)

func (x SpecialRatioValues) Enum() *SpecialRatioValues {
	p := new(SpecialRatioValues)
	*p = x
	return p
}

func (x SpecialRatioValues) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SpecialRatioValues) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_alerts_v3_event_payload_ratio_special_values_proto_enumTypes[0].Descriptor()
}

func (SpecialRatioValues) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_alerts_v3_event_payload_ratio_special_values_proto_enumTypes[0]
}

func (x SpecialRatioValues) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SpecialRatioValues.Descriptor instead.
func (SpecialRatioValues) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_payload_ratio_special_values_proto_rawDescGZIP(), []int{0}
}

var File_com_coralogixapis_alerts_v3_event_payload_ratio_special_values_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_event_payload_ratio_special_values_proto_rawDesc = []byte{
	0x0a, 0x44, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x2e, 0x76, 0x33, 0x2a, 0x46, 0x0a, 0x12, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x52, 0x61,
	0x74, 0x69, 0x6f, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x2c, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x41, 0x4c, 0x5f, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45,
	0x53, 0x5f, 0x49, 0x4e, 0x46, 0x49, 0x4e, 0x49, 0x54, 0x59, 0x5f, 0x4f, 0x52, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerts_v3_event_payload_ratio_special_values_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_event_payload_ratio_special_values_proto_rawDescData = file_com_coralogixapis_alerts_v3_event_payload_ratio_special_values_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_event_payload_ratio_special_values_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_event_payload_ratio_special_values_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_event_payload_ratio_special_values_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_event_payload_ratio_special_values_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_event_payload_ratio_special_values_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_event_payload_ratio_special_values_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_alerts_v3_event_payload_ratio_special_values_proto_goTypes = []any{
	(SpecialRatioValues)(0), // 0: com.coralogixapis.alerts.v3.SpecialRatioValues
}
var file_com_coralogixapis_alerts_v3_event_payload_ratio_special_values_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_alerts_v3_event_payload_ratio_special_values_proto_init() }
func file_com_coralogixapis_alerts_v3_event_payload_ratio_special_values_proto_init() {
	if File_com_coralogixapis_alerts_v3_event_payload_ratio_special_values_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_event_payload_ratio_special_values_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_event_payload_ratio_special_values_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_event_payload_ratio_special_values_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_alerts_v3_event_payload_ratio_special_values_proto_enumTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_event_payload_ratio_special_values_proto = out.File
	file_com_coralogixapis_alerts_v3_event_payload_ratio_special_values_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_event_payload_ratio_special_values_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_event_payload_ratio_special_values_proto_depIdxs = nil
}
