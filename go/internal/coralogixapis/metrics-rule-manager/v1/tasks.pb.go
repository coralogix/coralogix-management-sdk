// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: com/coralogixapis/metrics-rule-manager/v1/tasks.proto

package golang

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

type TaskCompletedRequest_Failure_Kind int32

const (
	TaskCompletedRequest_Failure_Transient TaskCompletedRequest_Failure_Kind = 0
	TaskCompletedRequest_Failure_Final     TaskCompletedRequest_Failure_Kind = 1
)

// Enum value maps for TaskCompletedRequest_Failure_Kind.
var (
	TaskCompletedRequest_Failure_Kind_name = map[int32]string{
		0: "Transient",
		1: "Final",
	}
	TaskCompletedRequest_Failure_Kind_value = map[string]int32{
		"Transient": 0,
		"Final":     1,
	}
)

func (x TaskCompletedRequest_Failure_Kind) Enum() *TaskCompletedRequest_Failure_Kind {
	p := new(TaskCompletedRequest_Failure_Kind)
	*p = x
	return p
}

func (x TaskCompletedRequest_Failure_Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskCompletedRequest_Failure_Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_enumTypes[0].Descriptor()
}

func (TaskCompletedRequest_Failure_Kind) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_enumTypes[0]
}

func (x TaskCompletedRequest_Failure_Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskCompletedRequest_Failure_Kind.Descriptor instead.
func (TaskCompletedRequest_Failure_Kind) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_rawDescGZIP(), []int{5, 1, 0}
}

type GetTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTaskRequest) Reset() {
	*x = GetTaskRequest{}
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskRequest) ProtoMessage() {}

func (x *GetTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskRequest.ProtoReflect.Descriptor instead.
func (*GetTaskRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_rawDescGZIP(), []int{0}
}

type GetTaskResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Task:
	//
	//	*GetTaskResponse_Eval
	Task          isGetTaskResponse_Task `protobuf_oneof:"task"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTaskResponse) Reset() {
	*x = GetTaskResponse{}
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskResponse) ProtoMessage() {}

func (x *GetTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskResponse.ProtoReflect.Descriptor instead.
func (*GetTaskResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_rawDescGZIP(), []int{1}
}

func (x *GetTaskResponse) GetTask() isGetTaskResponse_Task {
	if x != nil {
		return x.Task
	}
	return nil
}

func (x *GetTaskResponse) GetEval() *RuleGroupEval {
	if x != nil {
		if x, ok := x.Task.(*GetTaskResponse_Eval); ok {
			return x.Eval
		}
	}
	return nil
}

type isGetTaskResponse_Task interface {
	isGetTaskResponse_Task()
}

type GetTaskResponse_Eval struct {
	Eval *RuleGroupEval `protobuf:"bytes,1,opt,name=eval,proto3,oneof"`
}

func (*GetTaskResponse_Eval) isGetTaskResponse_Task() {}

type RuleGroupEval struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RuleGroup     *RuleGroup             `protobuf:"bytes,2,opt,name=rule_group,json=ruleGroup,proto3" json:"rule_group,omitempty"`
	EvalAt        uint64                 `protobuf:"varint,3,opt,name=eval_at,json=evalAt,proto3" json:"eval_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RuleGroupEval) Reset() {
	*x = RuleGroupEval{}
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RuleGroupEval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuleGroupEval) ProtoMessage() {}

func (x *RuleGroupEval) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuleGroupEval.ProtoReflect.Descriptor instead.
func (*RuleGroupEval) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_rawDescGZIP(), []int{2}
}

func (x *RuleGroupEval) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RuleGroupEval) GetRuleGroup() *RuleGroup {
	if x != nil {
		return x.RuleGroup
	}
	return nil
}

func (x *RuleGroupEval) GetEvalAt() uint64 {
	if x != nil {
		return x.EvalAt
	}
	return 0
}

type RuleGroup struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TenantId      uint32                 `protobuf:"varint,2,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	Rules         []*Rule                `protobuf:"bytes,3,rep,name=rules,proto3" json:"rules,omitempty"`
	Limit         uint64                 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RuleGroup) Reset() {
	*x = RuleGroup{}
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RuleGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuleGroup) ProtoMessage() {}

func (x *RuleGroup) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuleGroup.ProtoReflect.Descriptor instead.
func (*RuleGroup) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_rawDescGZIP(), []int{3}
}

func (x *RuleGroup) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RuleGroup) GetTenantId() uint32 {
	if x != nil {
		return x.TenantId
	}
	return 0
}

func (x *RuleGroup) GetRules() []*Rule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *RuleGroup) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type Rule struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Record        string                 `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	Expr          string                 `protobuf:"bytes,2,opt,name=expr,proto3" json:"expr,omitempty"`
	Labels        map[string]string      `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Rule) Reset() {
	*x = Rule{}
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Rule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rule) ProtoMessage() {}

func (x *Rule) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rule.ProtoReflect.Descriptor instead.
func (*Rule) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_rawDescGZIP(), []int{4}
}

func (x *Rule) GetRecord() string {
	if x != nil {
		return x.Record
	}
	return ""
}

func (x *Rule) GetExpr() string {
	if x != nil {
		return x.Expr
	}
	return ""
}

func (x *Rule) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type TaskCompletedRequest struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	TaskId string                 `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	// Types that are valid to be assigned to Result:
	//
	//	*TaskCompletedRequest_Success_
	//	*TaskCompletedRequest_Failure_
	Result        isTaskCompletedRequest_Result `protobuf_oneof:"result"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskCompletedRequest) Reset() {
	*x = TaskCompletedRequest{}
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskCompletedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskCompletedRequest) ProtoMessage() {}

func (x *TaskCompletedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskCompletedRequest.ProtoReflect.Descriptor instead.
func (*TaskCompletedRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_rawDescGZIP(), []int{5}
}

func (x *TaskCompletedRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *TaskCompletedRequest) GetResult() isTaskCompletedRequest_Result {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *TaskCompletedRequest) GetSuccess() *TaskCompletedRequest_Success {
	if x != nil {
		if x, ok := x.Result.(*TaskCompletedRequest_Success_); ok {
			return x.Success
		}
	}
	return nil
}

func (x *TaskCompletedRequest) GetFailure() *TaskCompletedRequest_Failure {
	if x != nil {
		if x, ok := x.Result.(*TaskCompletedRequest_Failure_); ok {
			return x.Failure
		}
	}
	return nil
}

type isTaskCompletedRequest_Result interface {
	isTaskCompletedRequest_Result()
}

type TaskCompletedRequest_Success_ struct {
	Success *TaskCompletedRequest_Success `protobuf:"bytes,2,opt,name=success,proto3,oneof"`
}

type TaskCompletedRequest_Failure_ struct {
	Failure *TaskCompletedRequest_Failure `protobuf:"bytes,3,opt,name=failure,proto3,oneof"`
}

func (*TaskCompletedRequest_Success_) isTaskCompletedRequest_Result() {}

func (*TaskCompletedRequest_Failure_) isTaskCompletedRequest_Result() {}

type TaskCompletedResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskCompletedResponse) Reset() {
	*x = TaskCompletedResponse{}
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskCompletedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskCompletedResponse) ProtoMessage() {}

func (x *TaskCompletedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskCompletedResponse.ProtoReflect.Descriptor instead.
func (*TaskCompletedResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_rawDescGZIP(), []int{6}
}

type TaskCompletedRequest_Success struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	RuleEvalDurationsMs map[string]uint64      `protobuf:"bytes,1,rep,name=rule_eval_durations_ms,json=ruleEvalDurationsMs,proto3" json:"rule_eval_durations_ms,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *TaskCompletedRequest_Success) Reset() {
	*x = TaskCompletedRequest_Success{}
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskCompletedRequest_Success) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskCompletedRequest_Success) ProtoMessage() {}

func (x *TaskCompletedRequest_Success) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskCompletedRequest_Success.ProtoReflect.Descriptor instead.
func (*TaskCompletedRequest_Success) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_rawDescGZIP(), []int{5, 0}
}

func (x *TaskCompletedRequest_Success) GetRuleEvalDurationsMs() map[string]uint64 {
	if x != nil {
		return x.RuleEvalDurationsMs
	}
	return nil
}

type TaskCompletedRequest_Failure struct {
	state         protoimpl.MessageState            `protogen:"open.v1"`
	Message       string                            `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Kind          TaskCompletedRequest_Failure_Kind `protobuf:"varint,2,opt,name=kind,proto3,enum=rule_manager.tasks.TaskCompletedRequest_Failure_Kind" json:"kind,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskCompletedRequest_Failure) Reset() {
	*x = TaskCompletedRequest_Failure{}
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskCompletedRequest_Failure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskCompletedRequest_Failure) ProtoMessage() {}

func (x *TaskCompletedRequest_Failure) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskCompletedRequest_Failure.ProtoReflect.Descriptor instead.
func (*TaskCompletedRequest_Failure) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_rawDescGZIP(), []int{5, 1}
}

func (x *TaskCompletedRequest_Failure) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *TaskCompletedRequest_Failure) GetKind() TaskCompletedRequest_Failure_Kind {
	if x != nil {
		return x.Kind
	}
	return TaskCompletedRequest_Failure_Transient
}

var File_com_coralogixapis_metrics_rule_manager_v1_tasks_proto protoreflect.FileDescriptor

var file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_rawDesc = []byte{
	0x0a, 0x35, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2d, 0x72, 0x75, 0x6c, 0x65,
	0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x10, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x52, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x37, 0x0a, 0x04, 0x65, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x45, 0x76, 0x61,
	0x6c, 0x48, 0x00, 0x52, 0x04, 0x65, 0x76, 0x61, 0x6c, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x61, 0x73,
	0x6b, 0x22, 0x76, 0x0a, 0x0d, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x45, 0x76,
	0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x3c, 0x0a, 0x0a, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x52, 0x75, 0x6c, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x09, 0x72, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x17, 0x0a, 0x07, 0x65, 0x76, 0x61, 0x6c, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x65, 0x76, 0x61, 0x6c, 0x41, 0x74, 0x22, 0x7e, 0x0a, 0x09, 0x52, 0x75, 0x6c,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75,
	0x6c, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0xab, 0x01, 0x0a, 0x04, 0x52, 0x75,
	0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x78,
	0x70, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x78, 0x70, 0x72, 0x12, 0x3c,
	0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xbc, 0x04, 0x0a, 0x14, 0x54, 0x61, 0x73, 0x6b,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x4c, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x72, 0x75, 0x6c,
	0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x48, 0x00, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x4c, 0x0a, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75,
	0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x48, 0x00, 0x52, 0x07, 0x66, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x1a, 0xd1, 0x01, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x7e, 0x0a, 0x16, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x5f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x49, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x45, 0x76, 0x61, 0x6c, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x13, 0x72, 0x75,
	0x6c, 0x65, 0x45, 0x76, 0x61, 0x6c, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4d,
	0x73, 0x1a, 0x46, 0x0a, 0x18, 0x52, 0x75, 0x6c, 0x65, 0x45, 0x76, 0x61, 0x6c, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x90, 0x01, 0x0a, 0x07, 0x46, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x49, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e,
	0x72, 0x75, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x2e,
	0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x20, 0x0a, 0x04, 0x4b, 0x69,
	0x6e, 0x64, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x65, 0x6e, 0x74, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x10, 0x01, 0x42, 0x08, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xc5, 0x01, 0x0a, 0x05, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x54, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x54, 0x61, 0x73, 0x6b, 0x12, 0x22, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x66, 0x0a, 0x0d, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x12, 0x28, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x72, 0x75, 0x6c,
	0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_rawDescOnce sync.Once
	file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_rawDescData = file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_rawDesc
)

func file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_rawDescData)
	})
	return file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_rawDescData
}

var file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_goTypes = []any{
	(TaskCompletedRequest_Failure_Kind)(0), // 0: rule_manager.tasks.TaskCompletedRequest.Failure.Kind
	(*GetTaskRequest)(nil),                 // 1: rule_manager.tasks.GetTaskRequest
	(*GetTaskResponse)(nil),                // 2: rule_manager.tasks.GetTaskResponse
	(*RuleGroupEval)(nil),                  // 3: rule_manager.tasks.RuleGroupEval
	(*RuleGroup)(nil),                      // 4: rule_manager.tasks.RuleGroup
	(*Rule)(nil),                           // 5: rule_manager.tasks.Rule
	(*TaskCompletedRequest)(nil),           // 6: rule_manager.tasks.TaskCompletedRequest
	(*TaskCompletedResponse)(nil),          // 7: rule_manager.tasks.TaskCompletedResponse
	nil,                                    // 8: rule_manager.tasks.Rule.LabelsEntry
	(*TaskCompletedRequest_Success)(nil),   // 9: rule_manager.tasks.TaskCompletedRequest.Success
	(*TaskCompletedRequest_Failure)(nil),   // 10: rule_manager.tasks.TaskCompletedRequest.Failure
	nil,                                    // 11: rule_manager.tasks.TaskCompletedRequest.Success.RuleEvalDurationsMsEntry
}
var file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_depIdxs = []int32{
	3,  // 0: rule_manager.tasks.GetTaskResponse.eval:type_name -> rule_manager.tasks.RuleGroupEval
	4,  // 1: rule_manager.tasks.RuleGroupEval.rule_group:type_name -> rule_manager.tasks.RuleGroup
	5,  // 2: rule_manager.tasks.RuleGroup.rules:type_name -> rule_manager.tasks.Rule
	8,  // 3: rule_manager.tasks.Rule.labels:type_name -> rule_manager.tasks.Rule.LabelsEntry
	9,  // 4: rule_manager.tasks.TaskCompletedRequest.success:type_name -> rule_manager.tasks.TaskCompletedRequest.Success
	10, // 5: rule_manager.tasks.TaskCompletedRequest.failure:type_name -> rule_manager.tasks.TaskCompletedRequest.Failure
	11, // 6: rule_manager.tasks.TaskCompletedRequest.Success.rule_eval_durations_ms:type_name -> rule_manager.tasks.TaskCompletedRequest.Success.RuleEvalDurationsMsEntry
	0,  // 7: rule_manager.tasks.TaskCompletedRequest.Failure.kind:type_name -> rule_manager.tasks.TaskCompletedRequest.Failure.Kind
	1,  // 8: rule_manager.tasks.Tasks.GetTask:input_type -> rule_manager.tasks.GetTaskRequest
	6,  // 9: rule_manager.tasks.Tasks.TaskCompleted:input_type -> rule_manager.tasks.TaskCompletedRequest
	2,  // 10: rule_manager.tasks.Tasks.GetTask:output_type -> rule_manager.tasks.GetTaskResponse
	7,  // 11: rule_manager.tasks.Tasks.TaskCompleted:output_type -> rule_manager.tasks.TaskCompletedResponse
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_init() }
func file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_init() {
	if File_com_coralogixapis_metrics_rule_manager_v1_tasks_proto != nil {
		return
	}
	file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_msgTypes[1].OneofWrappers = []any{
		(*GetTaskResponse_Eval)(nil),
	}
	file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_msgTypes[5].OneofWrappers = []any{
		(*TaskCompletedRequest_Success_)(nil),
		(*TaskCompletedRequest_Failure_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_metrics_rule_manager_v1_tasks_proto = out.File
	file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_rawDesc = nil
	file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_goTypes = nil
	file_com_coralogixapis_metrics_rule_manager_v1_tasks_proto_depIdxs = nil
}
