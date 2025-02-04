// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: com/coralogixapis/notification_center/routers/v1/router.proto

package v1

import (
	v1 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/notification_center/common/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// The private router allows notification destinations and routing rules
// to be specified as part of the notification request.
type PrivateRouter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouterKey string              `protobuf:"bytes,1,opt,name=router_key,json=routerKey,proto3" json:"router_key,omitempty"`
	Rules     []*v1.RoutingRule   `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty"`
	Fallback  []*v1.RoutingTarget `protobuf:"bytes,3,rep,name=fallback,proto3" json:"fallback,omitempty"`
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           *string                `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	UserFacingId *string                `protobuf:"bytes,2,opt,name=user_facing_id,json=userFacingId,proto3,oneof" json:"user_facing_id,omitempty"`
	EntityType   string                 `protobuf:"bytes,3,opt,name=entity_type,json=entityType,proto3" json:"entity_type,omitempty"`
	Name         string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description  string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Rules        []*v1.RoutingRule      `protobuf:"bytes,6,rep,name=rules,proto3" json:"rules,omitempty"`
	Fallback     []*v1.RoutingTarget    `protobuf:"bytes,7,rep,name=fallback,proto3" json:"fallback,omitempty"`
	CreateTime   *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=create_time,json=createTime,proto3,oneof" json:"create_time,omitempty"`
	UpdateTime   *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=update_time,json=updateTime,proto3,oneof" json:"update_time,omitempty"`
	EntityLabels map[string]string      `protobuf:"bytes,10,rep,name=entity_labels,json=entityLabels,proto3" json:"entity_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
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

func (x *GlobalRouter) GetUserFacingId() string {
	if x != nil && x.UserFacingId != nil {
		return *x.UserFacingId
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

var file_com_coralogixapis_notification_center_routers_v1_router_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x30, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x3d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xda, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x4b,
	0x65, 0x79, 0x12, 0x50, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x3a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72,
	0x75, 0x6c, 0x65, 0x73, 0x12, 0x58, 0x0a, 0x08, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x52, 0x08, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x22, 0xc7,
	0x05, 0x0a, 0x0c, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x12,
	0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x66, 0x61, 0x63,
	0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0c,
	0x75, 0x73, 0x65, 0x72, 0x46, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x50, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c,
	0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x58, 0x0a, 0x08, 0x66, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x67, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x08, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x48, 0x02, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x75, 0x0a, 0x0d, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x50, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x3f, 0x0a,
	0x11, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x05,
	0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x66,
	0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_notification_center_routers_v1_router_proto_rawDescOnce sync.Once
	file_com_coralogixapis_notification_center_routers_v1_router_proto_rawDescData = file_com_coralogixapis_notification_center_routers_v1_router_proto_rawDesc
)

func file_com_coralogixapis_notification_center_routers_v1_router_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_notification_center_routers_v1_router_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_notification_center_routers_v1_router_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_notification_center_routers_v1_router_proto_rawDescData)
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
			RawDescriptor: file_com_coralogixapis_notification_center_routers_v1_router_proto_rawDesc,
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
	file_com_coralogixapis_notification_center_routers_v1_router_proto_rawDesc = nil
	file_com_coralogixapis_notification_center_routers_v1_router_proto_goTypes = nil
	file_com_coralogixapis_notification_center_routers_v1_router_proto_depIdxs = nil
}
