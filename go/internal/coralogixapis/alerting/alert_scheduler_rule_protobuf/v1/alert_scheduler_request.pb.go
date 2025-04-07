// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.26.1
// source: com/coralogixapis/alerting/alert_scheduler_rule_protobuf/v1/alert_scheduler_request.proto

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

type GetAlertSchedulerRuleRequest struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	AlertSchedulerRuleId string                 `protobuf:"bytes,1,opt,name=alert_scheduler_rule_id,json=alertSchedulerRuleId,proto3" json:"alert_scheduler_rule_id,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *GetAlertSchedulerRuleRequest) Reset() {
	*x = GetAlertSchedulerRuleRequest{}
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAlertSchedulerRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlertSchedulerRuleRequest) ProtoMessage() {}

func (x *GetAlertSchedulerRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlertSchedulerRuleRequest.ProtoReflect.Descriptor instead.
func (*GetAlertSchedulerRuleRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_rawDescGZIP(), []int{0}
}

func (x *GetAlertSchedulerRuleRequest) GetAlertSchedulerRuleId() string {
	if x != nil {
		return x.AlertSchedulerRuleId
	}
	return ""
}

type CreateAlertSchedulerRuleRequest struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	AlertSchedulerRule *AlertSchedulerRule    `protobuf:"bytes,1,opt,name=alert_scheduler_rule,json=alertSchedulerRule,proto3" json:"alert_scheduler_rule,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *CreateAlertSchedulerRuleRequest) Reset() {
	*x = CreateAlertSchedulerRuleRequest{}
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAlertSchedulerRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAlertSchedulerRuleRequest) ProtoMessage() {}

func (x *CreateAlertSchedulerRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAlertSchedulerRuleRequest.ProtoReflect.Descriptor instead.
func (*CreateAlertSchedulerRuleRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAlertSchedulerRuleRequest) GetAlertSchedulerRule() *AlertSchedulerRule {
	if x != nil {
		return x.AlertSchedulerRule
	}
	return nil
}

type UpdateAlertSchedulerRuleRequest struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	AlertSchedulerRule *AlertSchedulerRule    `protobuf:"bytes,1,opt,name=alert_scheduler_rule,json=alertSchedulerRule,proto3" json:"alert_scheduler_rule,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *UpdateAlertSchedulerRuleRequest) Reset() {
	*x = UpdateAlertSchedulerRuleRequest{}
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAlertSchedulerRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAlertSchedulerRuleRequest) ProtoMessage() {}

func (x *UpdateAlertSchedulerRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAlertSchedulerRuleRequest.ProtoReflect.Descriptor instead.
func (*UpdateAlertSchedulerRuleRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateAlertSchedulerRuleRequest) GetAlertSchedulerRule() *AlertSchedulerRule {
	if x != nil {
		return x.AlertSchedulerRule
	}
	return nil
}

type DeleteAlertSchedulerRuleRequest struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	AlertSchedulerRuleId string                 `protobuf:"bytes,1,opt,name=alert_scheduler_rule_id,json=alertSchedulerRuleId,proto3" json:"alert_scheduler_rule_id,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *DeleteAlertSchedulerRuleRequest) Reset() {
	*x = DeleteAlertSchedulerRuleRequest{}
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAlertSchedulerRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAlertSchedulerRuleRequest) ProtoMessage() {}

func (x *DeleteAlertSchedulerRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAlertSchedulerRuleRequest.ProtoReflect.Descriptor instead.
func (*DeleteAlertSchedulerRuleRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteAlertSchedulerRuleRequest) GetAlertSchedulerRuleId() string {
	if x != nil {
		return x.AlertSchedulerRuleId
	}
	return ""
}

type AlertSchedulerRuleIds struct {
	state                 protoimpl.MessageState `protogen:"open.v1"`
	AlertSchedulerRuleIds []string               `protobuf:"bytes,1,rep,name=alert_scheduler_rule_ids,json=alertSchedulerRuleIds,proto3" json:"alert_scheduler_rule_ids,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *AlertSchedulerRuleIds) Reset() {
	*x = AlertSchedulerRuleIds{}
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlertSchedulerRuleIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertSchedulerRuleIds) ProtoMessage() {}

func (x *AlertSchedulerRuleIds) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertSchedulerRuleIds.ProtoReflect.Descriptor instead.
func (*AlertSchedulerRuleIds) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_rawDescGZIP(), []int{4}
}

func (x *AlertSchedulerRuleIds) GetAlertSchedulerRuleIds() []string {
	if x != nil {
		return x.AlertSchedulerRuleIds
	}
	return nil
}

type AlertSchedulerRuleVersionIds struct {
	state                        protoimpl.MessageState `protogen:"open.v1"`
	AlertSchedulerRuleVersionIds []string               `protobuf:"bytes,1,rep,name=alert_scheduler_rule_version_ids,json=alertSchedulerRuleVersionIds,proto3" json:"alert_scheduler_rule_version_ids,omitempty"`
	unknownFields                protoimpl.UnknownFields
	sizeCache                    protoimpl.SizeCache
}

func (x *AlertSchedulerRuleVersionIds) Reset() {
	*x = AlertSchedulerRuleVersionIds{}
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlertSchedulerRuleVersionIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertSchedulerRuleVersionIds) ProtoMessage() {}

func (x *AlertSchedulerRuleVersionIds) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertSchedulerRuleVersionIds.ProtoReflect.Descriptor instead.
func (*AlertSchedulerRuleVersionIds) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_rawDescGZIP(), []int{5}
}

func (x *AlertSchedulerRuleVersionIds) GetAlertSchedulerRuleVersionIds() []string {
	if x != nil {
		return x.AlertSchedulerRuleVersionIds
	}
	return nil
}

type FilterByAlertSchedulerRuleIds struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to AlertSchedulerRuleIds:
	//
	//	*FilterByAlertSchedulerRuleIds_AlertSchedulerIds
	//	*FilterByAlertSchedulerRuleIds_AlertSchedulerVersionIds
	AlertSchedulerRuleIds isFilterByAlertSchedulerRuleIds_AlertSchedulerRuleIds `protobuf_oneof:"alert_scheduler_rule_ids"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *FilterByAlertSchedulerRuleIds) Reset() {
	*x = FilterByAlertSchedulerRuleIds{}
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FilterByAlertSchedulerRuleIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterByAlertSchedulerRuleIds) ProtoMessage() {}

func (x *FilterByAlertSchedulerRuleIds) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterByAlertSchedulerRuleIds.ProtoReflect.Descriptor instead.
func (*FilterByAlertSchedulerRuleIds) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_rawDescGZIP(), []int{6}
}

func (x *FilterByAlertSchedulerRuleIds) GetAlertSchedulerRuleIds() isFilterByAlertSchedulerRuleIds_AlertSchedulerRuleIds {
	if x != nil {
		return x.AlertSchedulerRuleIds
	}
	return nil
}

func (x *FilterByAlertSchedulerRuleIds) GetAlertSchedulerIds() *AlertSchedulerRuleIds {
	if x != nil {
		if x, ok := x.AlertSchedulerRuleIds.(*FilterByAlertSchedulerRuleIds_AlertSchedulerIds); ok {
			return x.AlertSchedulerIds
		}
	}
	return nil
}

func (x *FilterByAlertSchedulerRuleIds) GetAlertSchedulerVersionIds() *AlertSchedulerRuleVersionIds {
	if x != nil {
		if x, ok := x.AlertSchedulerRuleIds.(*FilterByAlertSchedulerRuleIds_AlertSchedulerVersionIds); ok {
			return x.AlertSchedulerVersionIds
		}
	}
	return nil
}

type isFilterByAlertSchedulerRuleIds_AlertSchedulerRuleIds interface {
	isFilterByAlertSchedulerRuleIds_AlertSchedulerRuleIds()
}

type FilterByAlertSchedulerRuleIds_AlertSchedulerIds struct {
	AlertSchedulerIds *AlertSchedulerRuleIds `protobuf:"bytes,3,opt,name=alert_scheduler_ids,json=alertSchedulerIds,proto3,oneof"`
}

type FilterByAlertSchedulerRuleIds_AlertSchedulerVersionIds struct {
	AlertSchedulerVersionIds *AlertSchedulerRuleVersionIds `protobuf:"bytes,4,opt,name=alert_scheduler_version_ids,json=alertSchedulerVersionIds,proto3,oneof"`
}

func (*FilterByAlertSchedulerRuleIds_AlertSchedulerIds) isFilterByAlertSchedulerRuleIds_AlertSchedulerRuleIds() {
}

func (*FilterByAlertSchedulerRuleIds_AlertSchedulerVersionIds) isFilterByAlertSchedulerRuleIds_AlertSchedulerRuleIds() {
}

type GetBulkAlertSchedulerRuleRequest struct {
	state                  protoimpl.MessageState         `protogen:"open.v1"`
	ActiveTimeframe        *ActiveTimeframe               `protobuf:"bytes,1,opt,name=active_timeframe,json=activeTimeframe,proto3" json:"active_timeframe,omitempty"`
	Enabled                *bool                          `protobuf:"varint,2,opt,name=enabled,proto3,oneof" json:"enabled,omitempty"`
	AlertSchedulerRulesIds *FilterByAlertSchedulerRuleIds `protobuf:"bytes,3,opt,name=alert_scheduler_rules_ids,json=alertSchedulerRulesIds,proto3" json:"alert_scheduler_rules_ids,omitempty"`
	NextPageToken          *string                        `protobuf:"bytes,15,opt,name=next_page_token,json=nextPageToken,proto3,oneof" json:"next_page_token,omitempty"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *GetBulkAlertSchedulerRuleRequest) Reset() {
	*x = GetBulkAlertSchedulerRuleRequest{}
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBulkAlertSchedulerRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBulkAlertSchedulerRuleRequest) ProtoMessage() {}

func (x *GetBulkAlertSchedulerRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBulkAlertSchedulerRuleRequest.ProtoReflect.Descriptor instead.
func (*GetBulkAlertSchedulerRuleRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_rawDescGZIP(), []int{7}
}

func (x *GetBulkAlertSchedulerRuleRequest) GetActiveTimeframe() *ActiveTimeframe {
	if x != nil {
		return x.ActiveTimeframe
	}
	return nil
}

func (x *GetBulkAlertSchedulerRuleRequest) GetEnabled() bool {
	if x != nil && x.Enabled != nil {
		return *x.Enabled
	}
	return false
}

func (x *GetBulkAlertSchedulerRuleRequest) GetAlertSchedulerRulesIds() *FilterByAlertSchedulerRuleIds {
	if x != nil {
		return x.AlertSchedulerRulesIds
	}
	return nil
}

func (x *GetBulkAlertSchedulerRuleRequest) GetNextPageToken() string {
	if x != nil && x.NextPageToken != nil {
		return *x.NextPageToken
	}
	return ""
}

type CreateBulkAlertSchedulerRuleRequest struct {
	state                            protoimpl.MessageState             `protogen:"open.v1"`
	CreateAlertSchedulerRuleRequests []*CreateAlertSchedulerRuleRequest `protobuf:"bytes,1,rep,name=create_alert_scheduler_rule_requests,json=createAlertSchedulerRuleRequests,proto3" json:"create_alert_scheduler_rule_requests,omitempty"`
	unknownFields                    protoimpl.UnknownFields
	sizeCache                        protoimpl.SizeCache
}

func (x *CreateBulkAlertSchedulerRuleRequest) Reset() {
	*x = CreateBulkAlertSchedulerRuleRequest{}
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBulkAlertSchedulerRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBulkAlertSchedulerRuleRequest) ProtoMessage() {}

func (x *CreateBulkAlertSchedulerRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBulkAlertSchedulerRuleRequest.ProtoReflect.Descriptor instead.
func (*CreateBulkAlertSchedulerRuleRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_rawDescGZIP(), []int{8}
}

func (x *CreateBulkAlertSchedulerRuleRequest) GetCreateAlertSchedulerRuleRequests() []*CreateAlertSchedulerRuleRequest {
	if x != nil {
		return x.CreateAlertSchedulerRuleRequests
	}
	return nil
}

type UpdateBulkAlertSchedulerRuleRequest struct {
	state                            protoimpl.MessageState             `protogen:"open.v1"`
	UpdateAlertSchedulerRuleRequests []*UpdateAlertSchedulerRuleRequest `protobuf:"bytes,1,rep,name=update_alert_scheduler_rule_requests,json=updateAlertSchedulerRuleRequests,proto3" json:"update_alert_scheduler_rule_requests,omitempty"`
	unknownFields                    protoimpl.UnknownFields
	sizeCache                        protoimpl.SizeCache
}

func (x *UpdateBulkAlertSchedulerRuleRequest) Reset() {
	*x = UpdateBulkAlertSchedulerRuleRequest{}
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBulkAlertSchedulerRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBulkAlertSchedulerRuleRequest) ProtoMessage() {}

func (x *UpdateBulkAlertSchedulerRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBulkAlertSchedulerRuleRequest.ProtoReflect.Descriptor instead.
func (*UpdateBulkAlertSchedulerRuleRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateBulkAlertSchedulerRuleRequest) GetUpdateAlertSchedulerRuleRequests() []*UpdateAlertSchedulerRuleRequest {
	if x != nil {
		return x.UpdateAlertSchedulerRuleRequests
	}
	return nil
}

type DeleteBulkAlertSchedulerRuleRequest struct {
	state                            protoimpl.MessageState             `protogen:"open.v1"`
	DeleteAlertSchedulerRuleRequests []*DeleteAlertSchedulerRuleRequest `protobuf:"bytes,1,rep,name=delete_alert_scheduler_rule_requests,json=deleteAlertSchedulerRuleRequests,proto3" json:"delete_alert_scheduler_rule_requests,omitempty"`
	unknownFields                    protoimpl.UnknownFields
	sizeCache                        protoimpl.SizeCache
}

func (x *DeleteBulkAlertSchedulerRuleRequest) Reset() {
	*x = DeleteBulkAlertSchedulerRuleRequest{}
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteBulkAlertSchedulerRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBulkAlertSchedulerRuleRequest) ProtoMessage() {}

func (x *DeleteBulkAlertSchedulerRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBulkAlertSchedulerRuleRequest.ProtoReflect.Descriptor instead.
func (*DeleteBulkAlertSchedulerRuleRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteBulkAlertSchedulerRuleRequest) GetDeleteAlertSchedulerRuleRequests() []*DeleteAlertSchedulerRuleRequest {
	if x != nil {
		return x.DeleteAlertSchedulerRuleRequests
	}
	return nil
}

var File_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto protoreflect.FileDescriptor

const file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_rawDesc = "" +
	"\n" +
	"Ycom/coralogixapis/alerting/alert_scheduler_rule_protobuf/v1/alert_scheduler_request.proto\x12;com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1\x1aVcom/coralogixapis/alerting/alert_scheduler_rule_protobuf/v1/alert_scheduler_rule.proto\x1aRcom/coralogixapis/alerting/alert_scheduler_rule_protobuf/v1/active_timeframe.proto\"U\n" +
	"\x1cGetAlertSchedulerRuleRequest\x125\n" +
	"\x17alert_scheduler_rule_id\x18\x01 \x01(\tR\x14alertSchedulerRuleId\"\xa5\x01\n" +
	"\x1fCreateAlertSchedulerRuleRequest\x12\x81\x01\n" +
	"\x14alert_scheduler_rule\x18\x01 \x01(\v2O.com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleR\x12alertSchedulerRule\"\xa5\x01\n" +
	"\x1fUpdateAlertSchedulerRuleRequest\x12\x81\x01\n" +
	"\x14alert_scheduler_rule\x18\x01 \x01(\v2O.com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleR\x12alertSchedulerRule\"X\n" +
	"\x1fDeleteAlertSchedulerRuleRequest\x125\n" +
	"\x17alert_scheduler_rule_id\x18\x01 \x01(\tR\x14alertSchedulerRuleId\"P\n" +
	"\x15AlertSchedulerRuleIds\x127\n" +
	"\x18alert_scheduler_rule_ids\x18\x01 \x03(\tR\x15alertSchedulerRuleIds\"f\n" +
	"\x1cAlertSchedulerRuleVersionIds\x12F\n" +
	" alert_scheduler_rule_version_ids\x18\x01 \x03(\tR\x1calertSchedulerRuleVersionIds\"\xdf\x02\n" +
	"\x1dFilterByAlertSchedulerRuleIds\x12\x84\x01\n" +
	"\x13alert_scheduler_ids\x18\x03 \x01(\v2R.com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleIdsH\x00R\x11alertSchedulerIds\x12\x9a\x01\n" +
	"\x1balert_scheduler_version_ids\x18\x04 \x01(\v2Y.com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleVersionIdsH\x00R\x18alertSchedulerVersionIdsB\x1a\n" +
	"\x18alert_scheduler_rule_ids\"\x9f\x03\n" +
	" GetBulkAlertSchedulerRuleRequest\x12w\n" +
	"\x10active_timeframe\x18\x01 \x01(\v2L.com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.ActiveTimeframeR\x0factiveTimeframe\x12\x1d\n" +
	"\aenabled\x18\x02 \x01(\bH\x00R\aenabled\x88\x01\x01\x12\x95\x01\n" +
	"\x19alert_scheduler_rules_ids\x18\x03 \x01(\v2Z.com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.FilterByAlertSchedulerRuleIdsR\x16alertSchedulerRulesIds\x12+\n" +
	"\x0fnext_page_token\x18\x0f \x01(\tH\x01R\rnextPageToken\x88\x01\x01B\n" +
	"\n" +
	"\b_enabledB\x12\n" +
	"\x10_next_page_token\"\xd4\x01\n" +
	"#CreateBulkAlertSchedulerRuleRequest\x12\xac\x01\n" +
	"$create_alert_scheduler_rule_requests\x18\x01 \x03(\v2\\.com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.CreateAlertSchedulerRuleRequestR createAlertSchedulerRuleRequests\"\xd4\x01\n" +
	"#UpdateBulkAlertSchedulerRuleRequest\x12\xac\x01\n" +
	"$update_alert_scheduler_rule_requests\x18\x01 \x03(\v2\\.com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.UpdateAlertSchedulerRuleRequestR updateAlertSchedulerRuleRequests\"\xd4\x01\n" +
	"#DeleteBulkAlertSchedulerRuleRequest\x12\xac\x01\n" +
	"$delete_alert_scheduler_rule_requests\x18\x01 \x03(\v2\\.com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.DeleteAlertSchedulerRuleRequestR deleteAlertSchedulerRuleRequestsb\x06proto3"

var (
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_rawDescData []byte
)

func file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_rawDesc), len(file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_rawDescData
}

var file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_goTypes = []any{
	(*GetAlertSchedulerRuleRequest)(nil),        // 0: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.GetAlertSchedulerRuleRequest
	(*CreateAlertSchedulerRuleRequest)(nil),     // 1: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.CreateAlertSchedulerRuleRequest
	(*UpdateAlertSchedulerRuleRequest)(nil),     // 2: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.UpdateAlertSchedulerRuleRequest
	(*DeleteAlertSchedulerRuleRequest)(nil),     // 3: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.DeleteAlertSchedulerRuleRequest
	(*AlertSchedulerRuleIds)(nil),               // 4: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleIds
	(*AlertSchedulerRuleVersionIds)(nil),        // 5: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleVersionIds
	(*FilterByAlertSchedulerRuleIds)(nil),       // 6: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.FilterByAlertSchedulerRuleIds
	(*GetBulkAlertSchedulerRuleRequest)(nil),    // 7: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.GetBulkAlertSchedulerRuleRequest
	(*CreateBulkAlertSchedulerRuleRequest)(nil), // 8: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.CreateBulkAlertSchedulerRuleRequest
	(*UpdateBulkAlertSchedulerRuleRequest)(nil), // 9: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.UpdateBulkAlertSchedulerRuleRequest
	(*DeleteBulkAlertSchedulerRuleRequest)(nil), // 10: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.DeleteBulkAlertSchedulerRuleRequest
	(*AlertSchedulerRule)(nil),                  // 11: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRule
	(*ActiveTimeframe)(nil),                     // 12: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.ActiveTimeframe
}
var file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_depIdxs = []int32{
	11, // 0: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.CreateAlertSchedulerRuleRequest.alert_scheduler_rule:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRule
	11, // 1: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.UpdateAlertSchedulerRuleRequest.alert_scheduler_rule:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRule
	4,  // 2: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.FilterByAlertSchedulerRuleIds.alert_scheduler_ids:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleIds
	5,  // 3: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.FilterByAlertSchedulerRuleIds.alert_scheduler_version_ids:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleVersionIds
	12, // 4: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.GetBulkAlertSchedulerRuleRequest.active_timeframe:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.ActiveTimeframe
	6,  // 5: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.GetBulkAlertSchedulerRuleRequest.alert_scheduler_rules_ids:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.FilterByAlertSchedulerRuleIds
	1,  // 6: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.CreateBulkAlertSchedulerRuleRequest.create_alert_scheduler_rule_requests:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.CreateAlertSchedulerRuleRequest
	2,  // 7: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.UpdateBulkAlertSchedulerRuleRequest.update_alert_scheduler_rule_requests:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.UpdateAlertSchedulerRuleRequest
	3,  // 8: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.DeleteBulkAlertSchedulerRuleRequest.delete_alert_scheduler_rule_requests:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.DeleteAlertSchedulerRuleRequest
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_init()
}
func file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_init() {
	if File_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto != nil {
		return
	}
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_proto_init()
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_active_timeframe_proto_init()
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_msgTypes[6].OneofWrappers = []any{
		(*FilterByAlertSchedulerRuleIds_AlertSchedulerIds)(nil),
		(*FilterByAlertSchedulerRuleIds_AlertSchedulerVersionIds)(nil),
	}
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_msgTypes[7].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_rawDesc), len(file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto = out.File
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_goTypes = nil
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_request_proto_depIdxs = nil
}
