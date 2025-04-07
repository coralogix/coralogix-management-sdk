// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.26.1
// source: com/coralogixapis/notification_center/routers/v1/router.proto

package v1

import (
	v1 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/notification_center/common/v1"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// The private router allows notification destinations and routing rules
// to be specified as part of the notification request.
type PrivateRouter struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RouterKey     string                 `protobuf:"bytes,1,opt,name=router_key,json=routerKey,proto3" json:"router_key,omitempty"`
	Rules         []*v1.RoutingRule      `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty"`
	Fallback      []*v1.RoutingTarget    `protobuf:"bytes,3,rep,name=fallback,proto3" json:"fallback,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrivateRouter) Reset() {
	*x = PrivateRouter{}
	mi := &file_com_coralogixapis_notification_center_routers_v1_router_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrivateRouter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateRouter) ProtoMessage() {}

func (x *PrivateRouter) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_routers_v1_router_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateRouter.ProtoReflect.Descriptor instead.
func (*PrivateRouter) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_routers_v1_router_proto_rawDescGZIP(), []int{0}
}

func (x *PrivateRouter) GetRouterKey() string {
	if x != nil {
		return x.RouterKey
	}
	return ""
}

func (x *PrivateRouter) GetRules() []*v1.RoutingRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *PrivateRouter) GetFallback() []*v1.RoutingTarget {
	if x != nil {
		return x.Fallback
	}
	return nil
}

// The global router contains a pre-configured list of routing rules
// and can be specified as a part of the notification request.
type GlobalRouter struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// A unique identifier automatically generated by the service (do not provide this from the client side)
	Id *string `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// A unique identifier provided by the user
	UserDefinedId *string `protobuf:"bytes,2,opt,name=user_defined_id,json=userDefinedId,proto3,oneof" json:"user_defined_id,omitempty"`
	// The source of the notification (e.g., "alerts")
	EntityType  string `protobuf:"bytes,3,opt,name=entity_type,json=entityType,proto3" json:"entity_type,omitempty"`
	Name        string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// The routing rules to be applied in order to determine the destinations of the notification, teh first rule that matches will be used
	Rules []*v1.RoutingRule `protobuf:"bytes,6,rep,name=rules,proto3" json:"rules,omitempty"`
	// The fallback destinations to be used if no rule matches
	Fallback []*v1.RoutingTarget `protobuf:"bytes,7,rep,name=fallback,proto3" json:"fallback,omitempty"`
	// System-generated timestamp for when the router was last updated
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=create_time,json=createTime,proto3,oneof" json:"create_time,omitempty"`
	// System-generated timestamp for when the router was last updated
	UpdateTime    *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=update_time,json=updateTime,proto3,oneof" json:"update_time,omitempty"`
	EntityLabels  map[string]string      `protobuf:"bytes,10,rep,name=entity_labels,json=entityLabels,proto3" json:"entity_labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GlobalRouter) Reset() {
	*x = GlobalRouter{}
	mi := &file_com_coralogixapis_notification_center_routers_v1_router_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GlobalRouter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobalRouter) ProtoMessage() {}

func (x *GlobalRouter) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_notification_center_routers_v1_router_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobalRouter.ProtoReflect.Descriptor instead.
func (*GlobalRouter) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_notification_center_routers_v1_router_proto_rawDescGZIP(), []int{1}
}

func (x *GlobalRouter) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *GlobalRouter) GetUserDefinedId() string {
	if x != nil && x.UserDefinedId != nil {
		return *x.UserDefinedId
	}
	return ""
}

func (x *GlobalRouter) GetEntityType() string {
	if x != nil {
		return x.EntityType
	}
	return ""
}

func (x *GlobalRouter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GlobalRouter) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GlobalRouter) GetRules() []*v1.RoutingRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *GlobalRouter) GetFallback() []*v1.RoutingTarget {
	if x != nil {
		return x.Fallback
	}
	return nil
}

func (x *GlobalRouter) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *GlobalRouter) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *GlobalRouter) GetEntityLabels() map[string]string {
	if x != nil {
		return x.EntityLabels
	}
	return nil
}

var File_com_coralogixapis_notification_center_routers_v1_router_proto protoreflect.FileDescriptor

const file_com_coralogixapis_notification_center_routers_v1_router_proto_rawDesc = "" +
	"\n" +
	"=com/coralogixapis/notification_center/routers/v1/router.proto\x120com.coralogixapis.notification_center.routers.v1\x1a=com/coralogixapis/notification_center/common/v1/routing.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\xda\x01\n" +
	"\rPrivateRouter\x12\x1d\n" +
	"\n" +
	"router_key\x18\x01 \x01(\tR\trouterKey\x12P\n" +
	"\x05rules\x18\x02 \x03(\v2:.com.coralogixapis.notification_center.routing.RoutingRuleR\x05rules\x12X\n" +
	"\bfallback\x18\x03 \x03(\v2<.com.coralogixapis.notification_center.routing.RoutingTargetR\bfallback\"\xa6\b\n" +
	"\fGlobalRouter\x12@\n" +
	"\x02id\x18\x01 \x01(\tB+\x92A(J&\"a16e24c8-4db2-4abf-ba3c-c9e1fc35a3b9\"H\x00R\x02id\x88\x01\x01\x12D\n" +
	"\x0fuser_defined_id\x18\x02 \x01(\tB\x17\x92A\x14J\x12\"user-provided-id\"H\x01R\ruserDefinedId\x88\x01\x01\x12.\n" +
	"\ventity_type\x18\x03 \x01(\tB\r\x92A\n" +
	"J\b\"alerts\"R\n" +
	"entityType\x12$\n" +
	"\x04name\x18\x04 \x01(\tB\x10\x92A\rJ\v\"My Router\"R\x04name\x12 \n" +
	"\vdescription\x18\x05 \x01(\tR\vdescription\x12P\n" +
	"\x05rules\x18\x06 \x03(\v2:.com.coralogixapis.notification_center.routing.RoutingRuleR\x05rules\x12X\n" +
	"\bfallback\x18\a \x03(\v2<.com.coralogixapis.notification_center.routing.RoutingTargetR\bfallback\x12@\n" +
	"\vcreate_time\x18\b \x01(\v2\x1a.google.protobuf.TimestampH\x02R\n" +
	"createTime\x88\x01\x01\x12@\n" +
	"\vupdate_time\x18\t \x01(\v2\x1a.google.protobuf.TimestampH\x03R\n" +
	"updateTime\x88\x01\x01\x12u\n" +
	"\rentity_labels\x18\n" +
	" \x03(\v2P.com.coralogixapis.notification_center.routers.v1.GlobalRouter.EntityLabelsEntryR\fentityLabels\x1a?\n" +
	"\x11EntityLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01:\xf2\x01\x92A\xee\x01\n" +
	"o*\rGlobal Router2IDefines a set of pre-configured routing rules for directing notifications\xd2\x01\ventity_type\xd2\x01\x04name*{\n" +
	"'Find out more about notification center\x12Phttps://coralogix.com/docs/user-guides/notification-center/introduction/welcome/B\x05\n" +
	"\x03_idB\x12\n" +
	"\x10_user_defined_idB\x0e\n" +
	"\f_create_timeB\x0e\n" +
	"\f_update_timeb\x06proto3"

var (
	file_com_coralogixapis_notification_center_routers_v1_router_proto_rawDescOnce sync.Once
	file_com_coralogixapis_notification_center_routers_v1_router_proto_rawDescData []byte
)

func file_com_coralogixapis_notification_center_routers_v1_router_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_notification_center_routers_v1_router_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_notification_center_routers_v1_router_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_notification_center_routers_v1_router_proto_rawDesc), len(file_com_coralogixapis_notification_center_routers_v1_router_proto_rawDesc)))
	})
	return file_com_coralogixapis_notification_center_routers_v1_router_proto_rawDescData
}

var file_com_coralogixapis_notification_center_routers_v1_router_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogixapis_notification_center_routers_v1_router_proto_goTypes = []any{
	(*PrivateRouter)(nil),         // 0: com.coralogixapis.notification_center.routers.v1.PrivateRouter
	(*GlobalRouter)(nil),          // 1: com.coralogixapis.notification_center.routers.v1.GlobalRouter
	nil,                           // 2: com.coralogixapis.notification_center.routers.v1.GlobalRouter.EntityLabelsEntry
	(*v1.RoutingRule)(nil),        // 3: com.coralogixapis.notification_center.routing.RoutingRule
	(*v1.RoutingTarget)(nil),      // 4: com.coralogixapis.notification_center.routing.RoutingTarget
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_com_coralogixapis_notification_center_routers_v1_router_proto_depIdxs = []int32{
	3, // 0: com.coralogixapis.notification_center.routers.v1.PrivateRouter.rules:type_name -> com.coralogixapis.notification_center.routing.RoutingRule
	4, // 1: com.coralogixapis.notification_center.routers.v1.PrivateRouter.fallback:type_name -> com.coralogixapis.notification_center.routing.RoutingTarget
	3, // 2: com.coralogixapis.notification_center.routers.v1.GlobalRouter.rules:type_name -> com.coralogixapis.notification_center.routing.RoutingRule
	4, // 3: com.coralogixapis.notification_center.routers.v1.GlobalRouter.fallback:type_name -> com.coralogixapis.notification_center.routing.RoutingTarget
	5, // 4: com.coralogixapis.notification_center.routers.v1.GlobalRouter.create_time:type_name -> google.protobuf.Timestamp
	5, // 5: com.coralogixapis.notification_center.routers.v1.GlobalRouter.update_time:type_name -> google.protobuf.Timestamp
	2, // 6: com.coralogixapis.notification_center.routers.v1.GlobalRouter.entity_labels:type_name -> com.coralogixapis.notification_center.routers.v1.GlobalRouter.EntityLabelsEntry
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_notification_center_routers_v1_router_proto_init() }
func file_com_coralogixapis_notification_center_routers_v1_router_proto_init() {
	if File_com_coralogixapis_notification_center_routers_v1_router_proto != nil {
		return
	}
	file_com_coralogixapis_notification_center_routers_v1_router_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_notification_center_routers_v1_router_proto_rawDesc), len(file_com_coralogixapis_notification_center_routers_v1_router_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_notification_center_routers_v1_router_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_notification_center_routers_v1_router_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_notification_center_routers_v1_router_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_notification_center_routers_v1_router_proto = out.File
	file_com_coralogixapis_notification_center_routers_v1_router_proto_goTypes = nil
	file_com_coralogixapis_notification_center_routers_v1_router_proto_depIdxs = nil
}
