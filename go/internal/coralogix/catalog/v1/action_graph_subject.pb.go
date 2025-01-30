// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.2
// source: com/coralogix/catalog/v1/action_graph_subject.proto

package v1

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

type ActionGraphSubject int32

const (
	ActionGraphSubject_ACTION_GRAPH_SUBJECT_UNSPECIFIED    ActionGraphSubject = 0
	ActionGraphSubject_ACTION_GRAPH_SUBJECT_TIME_CONSUMING ActionGraphSubject = 1
	ActionGraphSubject_ACTION_GRAPH_SUBJECT_ERROR_RATE     ActionGraphSubject = 2
	ActionGraphSubject_ACTION_GRAPH_SUBJECT_LATENCY        ActionGraphSubject = 3
)

// Enum value maps for ActionGraphSubject.
var (
	ActionGraphSubject_name = map[int32]string{
		0: "ACTION_GRAPH_SUBJECT_UNSPECIFIED",
		1: "ACTION_GRAPH_SUBJECT_TIME_CONSUMING",
		2: "ACTION_GRAPH_SUBJECT_ERROR_RATE",
		3: "ACTION_GRAPH_SUBJECT_LATENCY",
	}
	ActionGraphSubject_value = map[string]int32{
		"ACTION_GRAPH_SUBJECT_UNSPECIFIED":    0,
		"ACTION_GRAPH_SUBJECT_TIME_CONSUMING": 1,
		"ACTION_GRAPH_SUBJECT_ERROR_RATE":     2,
		"ACTION_GRAPH_SUBJECT_LATENCY":        3,
	}
)

func (x ActionGraphSubject) Enum() *ActionGraphSubject {
	p := new(ActionGraphSubject)
	*p = x
	return p
}

func (x ActionGraphSubject) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionGraphSubject) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_catalog_v1_action_graph_subject_proto_enumTypes[0].Descriptor()
}

func (ActionGraphSubject) Type() protoreflect.EnumType {
	return &file_com_coralogix_catalog_v1_action_graph_subject_proto_enumTypes[0]
}

func (x ActionGraphSubject) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionGraphSubject.Descriptor instead.
func (ActionGraphSubject) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_catalog_v1_action_graph_subject_proto_rawDescGZIP(), []int{0}
}

var File_com_coralogix_catalog_v1_action_graph_subject_proto protoreflect.FileDescriptor

var file_com_coralogix_catalog_v1_action_graph_subject_proto_rawDesc = string([]byte{
	0x0a, 0x33, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2a,
	0xaa, 0x01, 0x0a, 0x12, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x53,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x24, 0x0a, 0x20, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x47, 0x52, 0x41, 0x50, 0x48, 0x5f, 0x53, 0x55, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x27, 0x0a, 0x23,
	0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x47, 0x52, 0x41, 0x50, 0x48, 0x5f, 0x53, 0x55, 0x42,
	0x4a, 0x45, 0x43, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x55, 0x4d,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x23, 0x0a, 0x1f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x47, 0x52, 0x41, 0x50, 0x48, 0x5f, 0x53, 0x55, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x5f, 0x52, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x47, 0x52, 0x41, 0x50, 0x48, 0x5f, 0x53, 0x55, 0x42, 0x4a, 0x45,
	0x43, 0x54, 0x5f, 0x4c, 0x41, 0x54, 0x45, 0x4e, 0x43, 0x59, 0x10, 0x03, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_com_coralogix_catalog_v1_action_graph_subject_proto_rawDescOnce sync.Once
	file_com_coralogix_catalog_v1_action_graph_subject_proto_rawDescData []byte
)

func file_com_coralogix_catalog_v1_action_graph_subject_proto_rawDescGZIP() []byte {
	file_com_coralogix_catalog_v1_action_graph_subject_proto_rawDescOnce.Do(func() {
		file_com_coralogix_catalog_v1_action_graph_subject_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogix_catalog_v1_action_graph_subject_proto_rawDesc), len(file_com_coralogix_catalog_v1_action_graph_subject_proto_rawDesc)))
	})
	return file_com_coralogix_catalog_v1_action_graph_subject_proto_rawDescData
}

var file_com_coralogix_catalog_v1_action_graph_subject_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogix_catalog_v1_action_graph_subject_proto_goTypes = []any{
	(ActionGraphSubject)(0), // 0: com.coralogix.catalog.v1.ActionGraphSubject
}
var file_com_coralogix_catalog_v1_action_graph_subject_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogix_catalog_v1_action_graph_subject_proto_init() }
func file_com_coralogix_catalog_v1_action_graph_subject_proto_init() {
	if File_com_coralogix_catalog_v1_action_graph_subject_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogix_catalog_v1_action_graph_subject_proto_rawDesc), len(file_com_coralogix_catalog_v1_action_graph_subject_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_catalog_v1_action_graph_subject_proto_goTypes,
		DependencyIndexes: file_com_coralogix_catalog_v1_action_graph_subject_proto_depIdxs,
		EnumInfos:         file_com_coralogix_catalog_v1_action_graph_subject_proto_enumTypes,
	}.Build()
	File_com_coralogix_catalog_v1_action_graph_subject_proto = out.File
	file_com_coralogix_catalog_v1_action_graph_subject_proto_goTypes = nil
	file_com_coralogix_catalog_v1_action_graph_subject_proto_depIdxs = nil
}
