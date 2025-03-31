// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/metrics-rule-manager/v1/groups.proto

package golang

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

// *
// A group of recording rules. Rules within a group are run sequentially at a regular interval,
// with the same evaluation time.
type InRuleGroup struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Name of the group.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// How often rules in the group are evaluated.
	Interval *uint32 `protobuf:"varint,2,opt,name=interval,proto3,oneof" json:"interval,omitempty"` // optional, default = 60 secs
	// Limits the number of series a rule can produce.
	Limit *uint64 `protobuf:"varint,3,opt,name=limit,proto3,oneof" json:"limit,omitempty"` // optional, 0 is no limit, default = 0
	// Rules of the group.
	Rules         []*InRule `protobuf:"bytes,4,rep,name=rules,proto3" json:"rules,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InRuleGroup) Reset() {
	*x = InRuleGroup{}
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InRuleGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InRuleGroup) ProtoMessage() {}

func (x *InRuleGroup) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InRuleGroup.ProtoReflect.Descriptor instead.
func (*InRuleGroup) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_rawDescGZIP(), []int{0}
}

func (x *InRuleGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InRuleGroup) GetInterval() uint32 {
	if x != nil && x.Interval != nil {
		return *x.Interval
	}
	return 0
}

func (x *InRuleGroup) GetLimit() uint64 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *InRuleGroup) GetRules() []*InRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

type InRule struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the time series to output to. Must be a valid metric name.
	Record string `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	// *
	// The PromQL expression to evaluate. Every evaluation cycle this is
	// evaluated at the current time, and the result recorded as a new set of
	// time series with the metric name as given by 'record'.
	Expr string `protobuf:"bytes,2,opt,name=expr,proto3" json:"expr,omitempty"`
	// Labels to add or overwrite before storing the result.
	Labels        map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InRule) Reset() {
	*x = InRule{}
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InRule) ProtoMessage() {}

func (x *InRule) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InRule.ProtoReflect.Descriptor instead.
func (*InRule) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_rawDescGZIP(), []int{1}
}

func (x *InRule) GetRecord() string {
	if x != nil {
		return x.Record
	}
	return ""
}

func (x *InRule) GetExpr() string {
	if x != nil {
		return x.Expr
	}
	return ""
}

func (x *InRule) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type OutRuleGroup struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Interval      *uint32                `protobuf:"varint,2,opt,name=interval,proto3,oneof" json:"interval,omitempty"`
	Limit         *uint64                `protobuf:"varint,3,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	Rules         []*OutRule             `protobuf:"bytes,4,rep,name=rules,proto3" json:"rules,omitempty"`
	LastEvalAt    *uint64                `protobuf:"varint,5,opt,name=last_eval_at,json=lastEvalAt,proto3,oneof" json:"last_eval_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OutRuleGroup) Reset() {
	*x = OutRuleGroup{}
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OutRuleGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutRuleGroup) ProtoMessage() {}

func (x *OutRuleGroup) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutRuleGroup.ProtoReflect.Descriptor instead.
func (*OutRuleGroup) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_rawDescGZIP(), []int{2}
}

func (x *OutRuleGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OutRuleGroup) GetInterval() uint32 {
	if x != nil && x.Interval != nil {
		return *x.Interval
	}
	return 0
}

func (x *OutRuleGroup) GetLimit() uint64 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *OutRuleGroup) GetRules() []*OutRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *OutRuleGroup) GetLastEvalAt() uint64 {
	if x != nil && x.LastEvalAt != nil {
		return *x.LastEvalAt
	}
	return 0
}

type OutRule struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	Record             string                 `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	Expr               string                 `protobuf:"bytes,2,opt,name=expr,proto3" json:"expr,omitempty"`
	Labels             map[string]string      `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	LastEvalDurationMs *uint64                `protobuf:"varint,4,opt,name=last_eval_duration_ms,json=lastEvalDurationMs,proto3,oneof" json:"last_eval_duration_ms,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *OutRule) Reset() {
	*x = OutRule{}
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OutRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutRule) ProtoMessage() {}

func (x *OutRule) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutRule.ProtoReflect.Descriptor instead.
func (*OutRule) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_rawDescGZIP(), []int{3}
}

func (x *OutRule) GetRecord() string {
	if x != nil {
		return x.Record
	}
	return ""
}

func (x *OutRule) GetExpr() string {
	if x != nil {
		return x.Expr
	}
	return ""
}

func (x *OutRule) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *OutRule) GetLastEvalDurationMs() uint64 {
	if x != nil && x.LastEvalDurationMs != nil {
		return *x.LastEvalDurationMs
	}
	return 0
}

// * A matcher specifying a group to delete.
type DeleteRuleGroup struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Name of the group.
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRuleGroup) Reset() {
	*x = DeleteRuleGroup{}
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRuleGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRuleGroup) ProtoMessage() {}

func (x *DeleteRuleGroup) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRuleGroup.ProtoReflect.Descriptor instead.
func (*DeleteRuleGroup) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRuleGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// * A listing of rule groups.
type RuleGroupListing struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of groups.
	RuleGroups    []*OutRuleGroup `protobuf:"bytes,1,rep,name=rule_groups,json=ruleGroups,proto3" json:"rule_groups,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RuleGroupListing) Reset() {
	*x = RuleGroupListing{}
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RuleGroupListing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuleGroupListing) ProtoMessage() {}

func (x *RuleGroupListing) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuleGroupListing.ProtoReflect.Descriptor instead.
func (*RuleGroupListing) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_rawDescGZIP(), []int{5}
}

func (x *RuleGroupListing) GetRuleGroups() []*OutRuleGroup {
	if x != nil {
		return x.RuleGroups
	}
	return nil
}

// * A matcher specifying a group to fetch.
type FetchRuleGroup struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Name of the group.
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FetchRuleGroup) Reset() {
	*x = FetchRuleGroup{}
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchRuleGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRuleGroup) ProtoMessage() {}

func (x *FetchRuleGroup) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRuleGroup.ProtoReflect.Descriptor instead.
func (*FetchRuleGroup) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_rawDescGZIP(), []int{6}
}

func (x *FetchRuleGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// * Result of a rule group fetch operation.
type FetchRuleGroupResult struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// A fetched rule group.
	RuleGroup     *OutRuleGroup `protobuf:"bytes,1,opt,name=rule_group,json=ruleGroup,proto3" json:"rule_group,omitempty"` // optional, empty if not found
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FetchRuleGroupResult) Reset() {
	*x = FetchRuleGroupResult{}
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchRuleGroupResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRuleGroupResult) ProtoMessage() {}

func (x *FetchRuleGroupResult) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRuleGroupResult.ProtoReflect.Descriptor instead.
func (*FetchRuleGroupResult) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_rawDescGZIP(), []int{7}
}

func (x *FetchRuleGroupResult) GetRuleGroup() *OutRuleGroup {
	if x != nil {
		return x.RuleGroup
	}
	return nil
}

type CreateRuleGroupSet struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Groups        []*InRuleGroup         `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
	Name          *string                `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRuleGroupSet) Reset() {
	*x = CreateRuleGroupSet{}
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRuleGroupSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRuleGroupSet) ProtoMessage() {}

func (x *CreateRuleGroupSet) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRuleGroupSet.ProtoReflect.Descriptor instead.
func (*CreateRuleGroupSet) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_rawDescGZIP(), []int{8}
}

func (x *CreateRuleGroupSet) GetGroups() []*InRuleGroup {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *CreateRuleGroupSet) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type CreateRuleGroupSetResult struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRuleGroupSetResult) Reset() {
	*x = CreateRuleGroupSetResult{}
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRuleGroupSetResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRuleGroupSetResult) ProtoMessage() {}

func (x *CreateRuleGroupSetResult) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRuleGroupSetResult.ProtoReflect.Descriptor instead.
func (*CreateRuleGroupSetResult) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_rawDescGZIP(), []int{9}
}

func (x *CreateRuleGroupSetResult) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateRuleGroupSet struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Groups        []*InRuleGroup         `protobuf:"bytes,2,rep,name=groups,proto3" json:"groups,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateRuleGroupSet) Reset() {
	*x = UpdateRuleGroupSet{}
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRuleGroupSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRuleGroupSet) ProtoMessage() {}

func (x *UpdateRuleGroupSet) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRuleGroupSet.ProtoReflect.Descriptor instead.
func (*UpdateRuleGroupSet) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateRuleGroupSet) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRuleGroupSet) GetGroups() []*InRuleGroup {
	if x != nil {
		return x.Groups
	}
	return nil
}

type RuleGroupSetListing struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Sets          []*OutRuleGroupSet     `protobuf:"bytes,1,rep,name=sets,proto3" json:"sets,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RuleGroupSetListing) Reset() {
	*x = RuleGroupSetListing{}
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RuleGroupSetListing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuleGroupSetListing) ProtoMessage() {}

func (x *RuleGroupSetListing) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuleGroupSetListing.ProtoReflect.Descriptor instead.
func (*RuleGroupSetListing) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_rawDescGZIP(), []int{11}
}

func (x *RuleGroupSetListing) GetSets() []*OutRuleGroupSet {
	if x != nil {
		return x.Sets
	}
	return nil
}

type FetchRuleGroupSet struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FetchRuleGroupSet) Reset() {
	*x = FetchRuleGroupSet{}
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchRuleGroupSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRuleGroupSet) ProtoMessage() {}

func (x *FetchRuleGroupSet) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRuleGroupSet.ProtoReflect.Descriptor instead.
func (*FetchRuleGroupSet) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_rawDescGZIP(), []int{12}
}

func (x *FetchRuleGroupSet) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteRuleGroupSet struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRuleGroupSet) Reset() {
	*x = DeleteRuleGroupSet{}
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRuleGroupSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRuleGroupSet) ProtoMessage() {}

func (x *DeleteRuleGroupSet) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRuleGroupSet.ProtoReflect.Descriptor instead.
func (*DeleteRuleGroupSet) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_rawDescGZIP(), []int{13}
}

func (x *DeleteRuleGroupSet) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type OutRuleGroupSet struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Groups        []*OutRuleGroup        `protobuf:"bytes,2,rep,name=groups,proto3" json:"groups,omitempty"`
	Name          *string                `protobuf:"bytes,3,opt,name=name,proto3,oneof" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OutRuleGroupSet) Reset() {
	*x = OutRuleGroupSet{}
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OutRuleGroupSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutRuleGroupSet) ProtoMessage() {}

func (x *OutRuleGroupSet) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutRuleGroupSet.ProtoReflect.Descriptor instead.
func (*OutRuleGroupSet) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_rawDescGZIP(), []int{14}
}

func (x *OutRuleGroupSet) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OutRuleGroupSet) GetGroups() []*OutRuleGroup {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *OutRuleGroupSet) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

var File_com_coralogixapis_metrics_rule_manager_v1_groups_proto protoreflect.FileDescriptor

const file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_rawDesc = "" +
	"\n" +
	"6com/coralogixapis/metrics-rule-manager/v1/groups.proto\x12\x13rule_manager.groups\x1a\x1bgoogle/protobuf/empty.proto\"\xa7\x01\n" +
	"\vInRuleGroup\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x1f\n" +
	"\binterval\x18\x02 \x01(\rH\x00R\binterval\x88\x01\x01\x12\x19\n" +
	"\x05limit\x18\x03 \x01(\x04H\x01R\x05limit\x88\x01\x01\x121\n" +
	"\x05rules\x18\x04 \x03(\v2\x1b.rule_manager.groups.InRuleR\x05rulesB\v\n" +
	"\t_intervalB\b\n" +
	"\x06_limit\"\xb0\x01\n" +
	"\x06InRule\x12\x16\n" +
	"\x06record\x18\x01 \x01(\tR\x06record\x12\x12\n" +
	"\x04expr\x18\x02 \x01(\tR\x04expr\x12?\n" +
	"\x06labels\x18\x03 \x03(\v2'.rule_manager.groups.InRule.LabelsEntryR\x06labels\x1a9\n" +
	"\vLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"\xe1\x01\n" +
	"\fOutRuleGroup\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x1f\n" +
	"\binterval\x18\x02 \x01(\rH\x00R\binterval\x88\x01\x01\x12\x19\n" +
	"\x05limit\x18\x03 \x01(\x04H\x01R\x05limit\x88\x01\x01\x122\n" +
	"\x05rules\x18\x04 \x03(\v2\x1c.rule_manager.groups.OutRuleR\x05rules\x12%\n" +
	"\flast_eval_at\x18\x05 \x01(\x04H\x02R\n" +
	"lastEvalAt\x88\x01\x01B\v\n" +
	"\t_intervalB\b\n" +
	"\x06_limitB\x0f\n" +
	"\r_last_eval_at\"\x84\x02\n" +
	"\aOutRule\x12\x16\n" +
	"\x06record\x18\x01 \x01(\tR\x06record\x12\x12\n" +
	"\x04expr\x18\x02 \x01(\tR\x04expr\x12@\n" +
	"\x06labels\x18\x03 \x03(\v2(.rule_manager.groups.OutRule.LabelsEntryR\x06labels\x126\n" +
	"\x15last_eval_duration_ms\x18\x04 \x01(\x04H\x00R\x12lastEvalDurationMs\x88\x01\x01\x1a9\n" +
	"\vLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01B\x18\n" +
	"\x16_last_eval_duration_ms\"%\n" +
	"\x0fDeleteRuleGroup\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"V\n" +
	"\x10RuleGroupListing\x12B\n" +
	"\vrule_groups\x18\x01 \x03(\v2!.rule_manager.groups.OutRuleGroupR\n" +
	"ruleGroups\"$\n" +
	"\x0eFetchRuleGroup\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"X\n" +
	"\x14FetchRuleGroupResult\x12@\n" +
	"\n" +
	"rule_group\x18\x01 \x01(\v2!.rule_manager.groups.OutRuleGroupR\truleGroup\"p\n" +
	"\x12CreateRuleGroupSet\x128\n" +
	"\x06groups\x18\x01 \x03(\v2 .rule_manager.groups.InRuleGroupR\x06groups\x12\x17\n" +
	"\x04name\x18\x02 \x01(\tH\x00R\x04name\x88\x01\x01B\a\n" +
	"\x05_name\"*\n" +
	"\x18CreateRuleGroupSetResult\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"^\n" +
	"\x12UpdateRuleGroupSet\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x128\n" +
	"\x06groups\x18\x02 \x03(\v2 .rule_manager.groups.InRuleGroupR\x06groups\"O\n" +
	"\x13RuleGroupSetListing\x128\n" +
	"\x04sets\x18\x01 \x03(\v2$.rule_manager.groups.OutRuleGroupSetR\x04sets\"#\n" +
	"\x11FetchRuleGroupSet\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"$\n" +
	"\x12DeleteRuleGroupSet\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"~\n" +
	"\x0fOutRuleGroupSet\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x129\n" +
	"\x06groups\x18\x02 \x03(\v2!.rule_manager.groups.OutRuleGroupR\x06groups\x12\x17\n" +
	"\x04name\x18\x03 \x01(\tH\x00R\x04name\x88\x01\x01B\a\n" +
	"\x05_name2\xbe\x02\n" +
	"\n" +
	"RuleGroups\x12B\n" +
	"\x04Save\x12 .rule_manager.groups.InRuleGroup\x1a\x16.google.protobuf.Empty\"\x00\x12H\n" +
	"\x06Delete\x12$.rule_manager.groups.DeleteRuleGroup\x1a\x16.google.protobuf.Empty\"\x00\x12G\n" +
	"\x04List\x12\x16.google.protobuf.Empty\x1a%.rule_manager.groups.RuleGroupListing\"\x00\x12Y\n" +
	"\x05Fetch\x12#.rule_manager.groups.FetchRuleGroup\x1a).rule_manager.groups.FetchRuleGroupResult\"\x002\xb2\x03\n" +
	"\rRuleGroupSets\x12b\n" +
	"\x06Create\x12'.rule_manager.groups.CreateRuleGroupSet\x1a-.rule_manager.groups.CreateRuleGroupSetResult\"\x00\x12K\n" +
	"\x06Update\x12'.rule_manager.groups.UpdateRuleGroupSet\x1a\x16.google.protobuf.Empty\"\x00\x12J\n" +
	"\x04List\x12\x16.google.protobuf.Empty\x1a(.rule_manager.groups.RuleGroupSetListing\"\x00\x12W\n" +
	"\x05Fetch\x12&.rule_manager.groups.FetchRuleGroupSet\x1a$.rule_manager.groups.OutRuleGroupSet\"\x00\x12K\n" +
	"\x06Delete\x12'.rule_manager.groups.DeleteRuleGroupSet\x1a\x16.google.protobuf.Empty\"\x00B\tZ\a/golangb\x06proto3"

var (
	file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_rawDescOnce sync.Once
	file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_rawDescData []byte
)

func file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_rawDesc), len(file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_rawDesc)))
	})
	return file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_rawDescData
}

var file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes = make([]protoimpl.MessageInfo, 17)
var file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_goTypes = []any{
	(*InRuleGroup)(nil),              // 0: rule_manager.groups.InRuleGroup
	(*InRule)(nil),                   // 1: rule_manager.groups.InRule
	(*OutRuleGroup)(nil),             // 2: rule_manager.groups.OutRuleGroup
	(*OutRule)(nil),                  // 3: rule_manager.groups.OutRule
	(*DeleteRuleGroup)(nil),          // 4: rule_manager.groups.DeleteRuleGroup
	(*RuleGroupListing)(nil),         // 5: rule_manager.groups.RuleGroupListing
	(*FetchRuleGroup)(nil),           // 6: rule_manager.groups.FetchRuleGroup
	(*FetchRuleGroupResult)(nil),     // 7: rule_manager.groups.FetchRuleGroupResult
	(*CreateRuleGroupSet)(nil),       // 8: rule_manager.groups.CreateRuleGroupSet
	(*CreateRuleGroupSetResult)(nil), // 9: rule_manager.groups.CreateRuleGroupSetResult
	(*UpdateRuleGroupSet)(nil),       // 10: rule_manager.groups.UpdateRuleGroupSet
	(*RuleGroupSetListing)(nil),      // 11: rule_manager.groups.RuleGroupSetListing
	(*FetchRuleGroupSet)(nil),        // 12: rule_manager.groups.FetchRuleGroupSet
	(*DeleteRuleGroupSet)(nil),       // 13: rule_manager.groups.DeleteRuleGroupSet
	(*OutRuleGroupSet)(nil),          // 14: rule_manager.groups.OutRuleGroupSet
	nil,                              // 15: rule_manager.groups.InRule.LabelsEntry
	nil,                              // 16: rule_manager.groups.OutRule.LabelsEntry
	(*emptypb.Empty)(nil),            // 17: google.protobuf.Empty
}
var file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_depIdxs = []int32{
	1,  // 0: rule_manager.groups.InRuleGroup.rules:type_name -> rule_manager.groups.InRule
	15, // 1: rule_manager.groups.InRule.labels:type_name -> rule_manager.groups.InRule.LabelsEntry
	3,  // 2: rule_manager.groups.OutRuleGroup.rules:type_name -> rule_manager.groups.OutRule
	16, // 3: rule_manager.groups.OutRule.labels:type_name -> rule_manager.groups.OutRule.LabelsEntry
	2,  // 4: rule_manager.groups.RuleGroupListing.rule_groups:type_name -> rule_manager.groups.OutRuleGroup
	2,  // 5: rule_manager.groups.FetchRuleGroupResult.rule_group:type_name -> rule_manager.groups.OutRuleGroup
	0,  // 6: rule_manager.groups.CreateRuleGroupSet.groups:type_name -> rule_manager.groups.InRuleGroup
	0,  // 7: rule_manager.groups.UpdateRuleGroupSet.groups:type_name -> rule_manager.groups.InRuleGroup
	14, // 8: rule_manager.groups.RuleGroupSetListing.sets:type_name -> rule_manager.groups.OutRuleGroupSet
	2,  // 9: rule_manager.groups.OutRuleGroupSet.groups:type_name -> rule_manager.groups.OutRuleGroup
	0,  // 10: rule_manager.groups.RuleGroups.Save:input_type -> rule_manager.groups.InRuleGroup
	4,  // 11: rule_manager.groups.RuleGroups.Delete:input_type -> rule_manager.groups.DeleteRuleGroup
	17, // 12: rule_manager.groups.RuleGroups.List:input_type -> google.protobuf.Empty
	6,  // 13: rule_manager.groups.RuleGroups.Fetch:input_type -> rule_manager.groups.FetchRuleGroup
	8,  // 14: rule_manager.groups.RuleGroupSets.Create:input_type -> rule_manager.groups.CreateRuleGroupSet
	10, // 15: rule_manager.groups.RuleGroupSets.Update:input_type -> rule_manager.groups.UpdateRuleGroupSet
	17, // 16: rule_manager.groups.RuleGroupSets.List:input_type -> google.protobuf.Empty
	12, // 17: rule_manager.groups.RuleGroupSets.Fetch:input_type -> rule_manager.groups.FetchRuleGroupSet
	13, // 18: rule_manager.groups.RuleGroupSets.Delete:input_type -> rule_manager.groups.DeleteRuleGroupSet
	17, // 19: rule_manager.groups.RuleGroups.Save:output_type -> google.protobuf.Empty
	17, // 20: rule_manager.groups.RuleGroups.Delete:output_type -> google.protobuf.Empty
	5,  // 21: rule_manager.groups.RuleGroups.List:output_type -> rule_manager.groups.RuleGroupListing
	7,  // 22: rule_manager.groups.RuleGroups.Fetch:output_type -> rule_manager.groups.FetchRuleGroupResult
	9,  // 23: rule_manager.groups.RuleGroupSets.Create:output_type -> rule_manager.groups.CreateRuleGroupSetResult
	17, // 24: rule_manager.groups.RuleGroupSets.Update:output_type -> google.protobuf.Empty
	11, // 25: rule_manager.groups.RuleGroupSets.List:output_type -> rule_manager.groups.RuleGroupSetListing
	14, // 26: rule_manager.groups.RuleGroupSets.Fetch:output_type -> rule_manager.groups.OutRuleGroupSet
	17, // 27: rule_manager.groups.RuleGroupSets.Delete:output_type -> google.protobuf.Empty
	19, // [19:28] is the sub-list for method output_type
	10, // [10:19] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_init() }
func file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_init() {
	if File_com_coralogixapis_metrics_rule_manager_v1_groups_proto != nil {
		return
	}
	file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[0].OneofWrappers = []any{}
	file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[2].OneofWrappers = []any{}
	file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[3].OneofWrappers = []any{}
	file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[8].OneofWrappers = []any{}
	file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes[14].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_rawDesc), len(file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   17,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_metrics_rule_manager_v1_groups_proto = out.File
	file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_goTypes = nil
	file_com_coralogixapis_metrics_rule_manager_v1_groups_proto_depIdxs = nil
}
