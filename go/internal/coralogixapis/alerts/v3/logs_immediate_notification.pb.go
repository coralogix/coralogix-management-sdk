// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.26.1
// source: com/coralogixapis/alerts/v3/notification/logs_immediate_notification.proto

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

type LogsImmediateNotification struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogsImmediateNotification) Reset() {
	*x = LogsImmediateNotification{}
	mi := &file_com_coralogixapis_alerts_v3_notification_logs_immediate_notification_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsImmediateNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsImmediateNotification) ProtoMessage() {}

func (x *LogsImmediateNotification) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_notification_logs_immediate_notification_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsImmediateNotification.ProtoReflect.Descriptor instead.
func (*LogsImmediateNotification) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_notification_logs_immediate_notification_proto_rawDescGZIP(), []int{0}
}

var File_com_coralogixapis_alerts_v3_notification_logs_immediate_notification_proto protoreflect.FileDescriptor

const file_com_coralogixapis_alerts_v3_notification_logs_immediate_notification_proto_rawDesc = "" +
	"\n" +
	"Jcom/coralogixapis/alerts/v3/notification/logs_immediate_notification.proto\x12\x1bcom.coralogixapis.alerts.v3\"\x1b\n" +
	"\x19LogsImmediateNotificationb\x06proto3"

var (
	file_com_coralogixapis_alerts_v3_notification_logs_immediate_notification_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_notification_logs_immediate_notification_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v3_notification_logs_immediate_notification_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_notification_logs_immediate_notification_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_notification_logs_immediate_notification_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_notification_logs_immediate_notification_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_notification_logs_immediate_notification_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerts_v3_notification_logs_immediate_notification_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_notification_logs_immediate_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_alerts_v3_notification_logs_immediate_notification_proto_goTypes = []any{
	(*LogsImmediateNotification)(nil), // 0: com.coralogixapis.alerts.v3.LogsImmediateNotification
}
var file_com_coralogixapis_alerts_v3_notification_logs_immediate_notification_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_alerts_v3_notification_logs_immediate_notification_proto_init() }
func file_com_coralogixapis_alerts_v3_notification_logs_immediate_notification_proto_init() {
	if File_com_coralogixapis_alerts_v3_notification_logs_immediate_notification_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_notification_logs_immediate_notification_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_notification_logs_immediate_notification_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_notification_logs_immediate_notification_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_notification_logs_immediate_notification_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerts_v3_notification_logs_immediate_notification_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_notification_logs_immediate_notification_proto = out.File
	file_com_coralogixapis_alerts_v3_notification_logs_immediate_notification_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_notification_logs_immediate_notification_proto_depIdxs = nil
}
