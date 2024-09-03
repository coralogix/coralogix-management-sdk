// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.3
// source: com/coralogixapis/alerts/v3/event/payload/common.proto

package v3

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

type CoralogixLogMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplicationName *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=application_name,json=applicationName,proto3" json:"application_name,omitempty"`
	SubsystemName   *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=subsystem_name,json=subsystemName,proto3" json:"subsystem_name,omitempty"`
	Severity        *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=severity,proto3" json:"severity,omitempty"`
}

func (x *CoralogixLogMetadata) Reset() {
	*x = CoralogixLogMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerts_v3_event_payload_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoralogixLogMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoralogixLogMetadata) ProtoMessage() {}

func (x *CoralogixLogMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_payload_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoralogixLogMetadata.ProtoReflect.Descriptor instead.
func (*CoralogixLogMetadata) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_payload_common_proto_rawDescGZIP(), []int{0}
}

func (x *CoralogixLogMetadata) GetApplicationName() *wrapperspb.StringValue {
	if x != nil {
		return x.ApplicationName
	}
	return nil
}

func (x *CoralogixLogMetadata) GetSubsystemName() *wrapperspb.StringValue {
	if x != nil {
		return x.SubsystemName
	}
	return nil
}

func (x *CoralogixLogMetadata) GetSeverity() *wrapperspb.StringValue {
	if x != nil {
		return x.Severity
	}
	return nil
}

var File_com_coralogixapis_alerts_v3_event_payload_common_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerts_v3_event_payload_common_proto_rawDesc = []byte{
	0x0a, 0x36, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xde, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x4c, 0x6f, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x47, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x43, 0x0a, 0x0e, 0x73, 0x75,
	0x62, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x38, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_com_coralogixapis_alerts_v3_event_payload_common_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_event_payload_common_proto_rawDescData = file_com_coralogixapis_alerts_v3_event_payload_common_proto_rawDesc
)

func file_com_coralogixapis_alerts_v3_event_payload_common_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_event_payload_common_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_event_payload_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerts_v3_event_payload_common_proto_rawDescData)
	})
	return file_com_coralogixapis_alerts_v3_event_payload_common_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_event_payload_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_alerts_v3_event_payload_common_proto_goTypes = []interface{}{
	(*CoralogixLogMetadata)(nil),   // 0: CoralogixLogMetadata
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
}
var file_com_coralogixapis_alerts_v3_event_payload_common_proto_depIdxs = []int32{
	1, // 0: CoralogixLogMetadata.application_name:type_name -> google.protobuf.StringValue
	1, // 1: CoralogixLogMetadata.subsystem_name:type_name -> google.protobuf.StringValue
	1, // 2: CoralogixLogMetadata.severity:type_name -> google.protobuf.StringValue
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_alerts_v3_event_payload_common_proto_init() }
func file_com_coralogixapis_alerts_v3_event_payload_common_proto_init() {
	if File_com_coralogixapis_alerts_v3_event_payload_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_alerts_v3_event_payload_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoralogixLogMetadata); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerts_v3_event_payload_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_event_payload_common_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_event_payload_common_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_event_payload_common_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_event_payload_common_proto = out.File
	file_com_coralogixapis_alerts_v3_event_payload_common_proto_rawDesc = nil
	file_com_coralogixapis_alerts_v3_event_payload_common_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_event_payload_common_proto_depIdxs = nil
}