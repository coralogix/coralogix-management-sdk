// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/notification_center/notifications/v1/testing.proto

package v1

import (
	v1 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/notification_center/common/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

type TestResult struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Result:
	//
	//	*TestResult_Success_
	//	*TestResult_Failure_
	Result        isTestResult_Result `protobuf_oneof:"result"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TestResult) Reset() {
	*x = TestResult{}
	mi := &file_com_coralogixapis_notification_center_notifications_v1_testing_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TestResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResult) ProtoMessage() {}

func (x *TestResult) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_notifications_v1_testing_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResult.ProtoReflect.Descriptor instead.
func (*TestResult) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_notifications_v1_testing_proto_rawDescGZIP(), []int{0}
}

func (x *TestResult) GetResult() isTestResult_Result {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *TestResult) GetSuccess() *TestResult_Success {
	if x != nil {
		if x, ok := x.Result.(*TestResult_Success_); ok {
			return x.Success
		}
	}
	return nil
}

func (x *TestResult) GetFailure() *TestResult_Failure {
	if x != nil {
		if x, ok := x.Result.(*TestResult_Failure_); ok {
			return x.Failure
		}
	}
	return nil
}

type isTestResult_Result interface {
	isTestResult_Result()
}

type TestResult_Success_ struct {
	Success *TestResult_Success `protobuf:"bytes,1,opt,name=success,proto3,oneof"`
}

type TestResult_Failure_ struct {
	Failure *TestResult_Failure `protobuf:"bytes,2,opt,name=failure,proto3,oneof"`
}

func (*TestResult_Success_) isTestResult_Result() {}

func (*TestResult_Failure_) isTestResult_Result() {}

type TestNotification struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TeamId         uint32                 `protobuf:"varint,2,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	ConnectorId    string                 `protobuf:"bytes,3,opt,name=connector_id,json=connectorId,proto3" json:"connector_id,omitempty"`
	OutputSchemaId string                 `protobuf:"bytes,4,opt,name=output_schema_id,json=outputSchemaId,proto3" json:"output_schema_id,omitempty"`
	// Deprecated: Marked as deprecated in com/coralogixapis/notification_center/notifications/v1/testing.proto.
	DeprecatedEntityData *structpb.Struct           `protobuf:"bytes,5,opt,name=deprecated_entity_data,json=deprecatedEntityData,proto3,oneof" json:"deprecated_entity_data,omitempty"`
	ConnectorConfig      []*v1.ConnectorConfigField `protobuf:"bytes,6,rep,name=connector_config,json=connectorConfig,proto3" json:"connector_config,omitempty"`
	MessageConfig        []*v1.MessageConfigField   `protobuf:"bytes,7,rep,name=message_config,json=messageConfig,proto3" json:"message_config,omitempty"`
	AvailableOperations  []*NotificationOperation   `protobuf:"bytes,8,rep,name=available_operations,json=availableOperations,proto3" json:"available_operations,omitempty"`
	Attachments          []*NotificationAttachment  `protobuf:"bytes,9,rep,name=attachments,proto3" json:"attachments,omitempty"`
	EntityData           string                     `protobuf:"bytes,10,opt,name=entity_data,json=entityData,proto3" json:"entity_data,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *TestNotification) Reset() {
	*x = TestNotification{}
	mi := &file_com_coralogixapis_notification_center_notifications_v1_testing_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TestNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestNotification) ProtoMessage() {}

func (x *TestNotification) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_notifications_v1_testing_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestNotification.ProtoReflect.Descriptor instead.
func (*TestNotification) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_notifications_v1_testing_proto_rawDescGZIP(), []int{1}
}

func (x *TestNotification) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TestNotification) GetTeamId() uint32 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *TestNotification) GetConnectorId() string {
	if x != nil {
		return x.ConnectorId
	}
	return ""
}

func (x *TestNotification) GetOutputSchemaId() string {
	if x != nil {
		return x.OutputSchemaId
	}
	return ""
}

// Deprecated: Marked as deprecated in com/coralogixapis/notification_center/notifications/v1/testing.proto.
func (x *TestNotification) GetDeprecatedEntityData() *structpb.Struct {
	if x != nil {
		return x.DeprecatedEntityData
	}
	return nil
}

func (x *TestNotification) GetConnectorConfig() []*v1.ConnectorConfigField {
	if x != nil {
		return x.ConnectorConfig
	}
	return nil
}

func (x *TestNotification) GetMessageConfig() []*v1.MessageConfigField {
	if x != nil {
		return x.MessageConfig
	}
	return nil
}

func (x *TestNotification) GetAvailableOperations() []*NotificationOperation {
	if x != nil {
		return x.AvailableOperations
	}
	return nil
}

func (x *TestNotification) GetAttachments() []*NotificationAttachment {
	if x != nil {
		return x.Attachments
	}
	return nil
}

func (x *TestNotification) GetEntityData() string {
	if x != nil {
		return x.EntityData
	}
	return ""
}

type TestResult_Success struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TestResult_Success) Reset() {
	*x = TestResult_Success{}
	mi := &file_com_coralogixapis_notification_center_notifications_v1_testing_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TestResult_Success) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResult_Success) ProtoMessage() {}

func (x *TestResult_Success) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_notifications_v1_testing_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResult_Success.ProtoReflect.Descriptor instead.
func (*TestResult_Success) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_notifications_v1_testing_proto_rawDescGZIP(), []int{0, 0}
}

type TestResult_Failure struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	StatusCode    *uint32                `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3,oneof" json:"status_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TestResult_Failure) Reset() {
	*x = TestResult_Failure{}
	mi := &file_com_coralogixapis_notification_center_notifications_v1_testing_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TestResult_Failure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResult_Failure) ProtoMessage() {}

func (x *TestResult_Failure) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_notifications_v1_testing_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResult_Failure.ProtoReflect.Descriptor instead.
func (*TestResult_Failure) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_notifications_v1_testing_proto_rawDescGZIP(), []int{0, 1}
}

func (x *TestResult_Failure) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *TestResult_Failure) GetStatusCode() uint32 {
	if x != nil && x.StatusCode != nil {
		return *x.StatusCode
	}
	return 0
}

var File_com_coralogixapis_notification_center_notifications_v1_testing_proto protoreflect.FileDescriptor

const file_com_coralogixapis_notification_center_notifications_v1_testing_proto_rawDesc = "" +
	"\n" +
	"Dcom/coralogixapis/notification_center/notifications/v1/testing.proto\x126com.coralogixapis.notification_center.notifications.v1\x1aIcom/coralogixapis/notification_center/notifications/v1/notification.proto\x1a\x1cgoogle/protobuf/struct.proto\x1aCcom/coralogixapis/notification_center/common/v1/config_fields.proto\"\xcc\x02\n" +
	"\n" +
	"TestResult\x12f\n" +
	"\asuccess\x18\x01 \x01(\v2J.com.coralogixapis.notification_center.notifications.v1.TestResult.SuccessH\x00R\asuccess\x12f\n" +
	"\afailure\x18\x02 \x01(\v2J.com.coralogixapis.notification_center.notifications.v1.TestResult.FailureH\x00R\afailure\x1a\t\n" +
	"\aSuccess\x1aY\n" +
	"\aFailure\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\x12$\n" +
	"\vstatus_code\x18\x02 \x01(\rH\x00R\n" +
	"statusCode\x88\x01\x01B\x0e\n" +
	"\f_status_codeB\b\n" +
	"\x06result\"\xdb\x05\n" +
	"\x10TestNotification\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x17\n" +
	"\ateam_id\x18\x02 \x01(\rR\x06teamId\x12!\n" +
	"\fconnector_id\x18\x03 \x01(\tR\vconnectorId\x12(\n" +
	"\x10output_schema_id\x18\x04 \x01(\tR\x0eoutputSchemaId\x12V\n" +
	"\x16deprecated_entity_data\x18\x05 \x01(\v2\x17.google.protobuf.StructB\x02\x18\x01H\x00R\x14deprecatedEntityData\x88\x01\x01\x12f\n" +
	"\x10connector_config\x18\x06 \x03(\v2;.com.coralogixapis.notification_center.ConnectorConfigFieldR\x0fconnectorConfig\x12`\n" +
	"\x0emessage_config\x18\a \x03(\v29.com.coralogixapis.notification_center.MessageConfigFieldR\rmessageConfig\x12\x80\x01\n" +
	"\x14available_operations\x18\b \x03(\v2M.com.coralogixapis.notification_center.notifications.v1.NotificationOperationR\x13availableOperations\x12p\n" +
	"\vattachments\x18\t \x03(\v2N.com.coralogixapis.notification_center.notifications.v1.NotificationAttachmentR\vattachments\x12\x1f\n" +
	"\ventity_data\x18\n" +
	" \x01(\tR\n" +
	"entityDataB\x19\n" +
	"\x17_deprecated_entity_datab\x06proto3"

var (
	file_com_coralogixapis_notification_center_notifications_v1_testing_proto_rawDescOnce sync.Once
	file_com_coralogixapis_notification_center_notifications_v1_testing_proto_rawDescData []byte
)

func file_com_coralogixapis_notification_center_notifications_v1_testing_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_notification_center_notifications_v1_testing_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_notification_center_notifications_v1_testing_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_notification_center_notifications_v1_testing_proto_rawDesc), len(file_com_coralogixapis_notification_center_notifications_v1_testing_proto_rawDesc)))
	})
	return file_com_coralogixapis_notification_center_notifications_v1_testing_proto_rawDescData
}

var file_com_coralogixapis_notification_center_notifications_v1_testing_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_com_coralogixapis_notification_center_notifications_v1_testing_proto_goTypes = []any{
	(*TestResult)(nil),              // 0: com.coralogixapis.notification_center.notifications.v1.TestResult
	(*TestNotification)(nil),        // 1: com.coralogixapis.notification_center.notifications.v1.TestNotification
	(*TestResult_Success)(nil),      // 2: com.coralogixapis.notification_center.notifications.v1.TestResult.Success
	(*TestResult_Failure)(nil),      // 3: com.coralogixapis.notification_center.notifications.v1.TestResult.Failure
	(*structpb.Struct)(nil),         // 4: google.protobuf.Struct
	(*v1.ConnectorConfigField)(nil), // 5: com.coralogixapis.notification_center.ConnectorConfigField
	(*v1.MessageConfigField)(nil),   // 6: com.coralogixapis.notification_center.MessageConfigField
	(*NotificationOperation)(nil),   // 7: com.coralogixapis.notification_center.notifications.v1.NotificationOperation
	(*NotificationAttachment)(nil),  // 8: com.coralogixapis.notification_center.notifications.v1.NotificationAttachment
}
var file_com_coralogixapis_notification_center_notifications_v1_testing_proto_depIdxs = []int32{
	2, // 0: com.coralogixapis.notification_center.notifications.v1.TestResult.success:type_name -> com.coralogixapis.notification_center.notifications.v1.TestResult.Success
	3, // 1: com.coralogixapis.notification_center.notifications.v1.TestResult.failure:type_name -> com.coralogixapis.notification_center.notifications.v1.TestResult.Failure
	4, // 2: com.coralogixapis.notification_center.notifications.v1.TestNotification.deprecated_entity_data:type_name -> google.protobuf.Struct
	5, // 3: com.coralogixapis.notification_center.notifications.v1.TestNotification.connector_config:type_name -> com.coralogixapis.notification_center.ConnectorConfigField
	6, // 4: com.coralogixapis.notification_center.notifications.v1.TestNotification.message_config:type_name -> com.coralogixapis.notification_center.MessageConfigField
	7, // 5: com.coralogixapis.notification_center.notifications.v1.TestNotification.available_operations:type_name -> com.coralogixapis.notification_center.notifications.v1.NotificationOperation
	8, // 6: com.coralogixapis.notification_center.notifications.v1.TestNotification.attachments:type_name -> com.coralogixapis.notification_center.notifications.v1.NotificationAttachment
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_notification_center_notifications_v1_testing_proto_init() }
func file_com_coralogixapis_notification_center_notifications_v1_testing_proto_init() {
	if File_com_coralogixapis_notification_center_notifications_v1_testing_proto != nil {
		return
	}
	file_com_coralogixapis_notification_center_notifications_v1_notification_proto_init()
	file_com_coralogixapis_notification_center_notifications_v1_testing_proto_msgTypes[0].OneofWrappers = []any{
		(*TestResult_Success_)(nil),
		(*TestResult_Failure_)(nil),
	}
	file_com_coralogixapis_notification_center_notifications_v1_testing_proto_msgTypes[1].OneofWrappers = []any{}
	file_com_coralogixapis_notification_center_notifications_v1_testing_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_notification_center_notifications_v1_testing_proto_rawDesc), len(file_com_coralogixapis_notification_center_notifications_v1_testing_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_notification_center_notifications_v1_testing_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_notification_center_notifications_v1_testing_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_notification_center_notifications_v1_testing_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_notification_center_notifications_v1_testing_proto = out.File
	file_com_coralogixapis_notification_center_notifications_v1_testing_proto_goTypes = nil
	file_com_coralogixapis_notification_center_notifications_v1_testing_proto_depIdxs = nil
}
