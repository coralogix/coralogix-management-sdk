// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.2
// source: com/coralogixapis/alerts/v3/commons.proto

package v3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type NotifyOn int32

const (
	NotifyOn_NOTIFY_ON_TRIGGERED_ONLY_UNSPECIFIED NotifyOn = 0
	NotifyOn_NOTIFY_ON_TRIGGERED_AND_RESOLVED     NotifyOn = 1
)

// Enum value maps for NotifyOn.
var (
	NotifyOn_name = map[int32]string{
		0: "NOTIFY_ON_TRIGGERED_ONLY_UNSPECIFIED",
		1: "NOTIFY_ON_TRIGGERED_AND_RESOLVED",
	}
	NotifyOn_value = map[string]int32{
		"NOTIFY_ON_TRIGGERED_ONLY_UNSPECIFIED": 0,
		"NOTIFY_ON_TRIGGERED_AND_RESOLVED":     1,
	}
)

func (x NotifyOn) Enum() *NotifyOn {
	p := new(NotifyOn)
	*p = x
	return p
}

func (x NotifyOn) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotifyOn) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_alerts_v3_commons_proto_enumTypes[0].Descriptor()
}

func (NotifyOn) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_alerts_v3_commons_proto_enumTypes[0]
}

func (x NotifyOn) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotifyOn.Descriptor instead.
func (NotifyOn) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_commons_proto_rawDescGZIP(), []int{0}
}

var File_com_coralogixapis_alerts_v3_commons_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_commons_proto_rawDesc = string([]byte{
	0x0a, 0x29, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2a, 0x5a, 0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x4f, 0x6e, 0x12, 0x28, 0x0a, 0x24, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x4f,
	0x4e, 0x5f, 0x54, 0x52, 0x49, 0x47, 0x47, 0x45, 0x52, 0x45, 0x44, 0x5f, 0x4f, 0x4e, 0x4c, 0x59,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x24,
	0x0a, 0x20, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x4f, 0x4e, 0x5f, 0x54, 0x52, 0x49, 0x47,
	0x47, 0x45, 0x52, 0x45, 0x44, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x56,
	0x45, 0x44, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_com_coralogixapis_alerts_v3_commons_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_commons_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v3_commons_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_commons_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_commons_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_commons_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_commons_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerts_v3_commons_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_commons_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_alerts_v3_commons_proto_goTypes = []any{
	(NotifyOn)(0), // 0: com.coralogixapis.alerts.v3.NotifyOn
}
var file_com_coralogixapis_alerts_v3_commons_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_alerts_v3_commons_proto_init() }
func file_com_coralogixapis_alerts_v3_commons_proto_init() {
	if File_com_coralogixapis_alerts_v3_commons_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_commons_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_commons_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_commons_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_commons_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_alerts_v3_commons_proto_enumTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_commons_proto = out.File
	file_com_coralogixapis_alerts_v3_commons_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_commons_proto_depIdxs = nil
}
