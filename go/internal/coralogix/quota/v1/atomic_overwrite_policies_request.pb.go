// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: com/coralogix/quota/v1/atomic_overwrite_policies_request.proto

package v1

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

// In an atomic operation delete all existing log policies and create the provided list by order
type AtomicOverwriteLogPoliciesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Policies []*CreateLogPolicyRequest `protobuf:"bytes,1,rep,name=policies,proto3" json:"policies,omitempty"`
}

func (x *AtomicOverwriteLogPoliciesRequest) Reset() {
	*x = AtomicOverwriteLogPoliciesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AtomicOverwriteLogPoliciesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AtomicOverwriteLogPoliciesRequest) ProtoMessage() {}

func (x *AtomicOverwriteLogPoliciesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AtomicOverwriteLogPoliciesRequest.ProtoReflect.Descriptor instead.
func (*AtomicOverwriteLogPoliciesRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_rawDescGZIP(), []int{0}
}

func (x *AtomicOverwriteLogPoliciesRequest) GetPolicies() []*CreateLogPolicyRequest {
	if x != nil {
		return x.Policies
	}
	return nil
}

// In an atomic operation delete all existing span policies and create the provided list by order
type AtomicOverwriteSpanPoliciesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Policies []*CreateSpanPolicyRequest `protobuf:"bytes,1,rep,name=policies,proto3" json:"policies,omitempty"`
}

func (x *AtomicOverwriteSpanPoliciesRequest) Reset() {
	*x = AtomicOverwriteSpanPoliciesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AtomicOverwriteSpanPoliciesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AtomicOverwriteSpanPoliciesRequest) ProtoMessage() {}

func (x *AtomicOverwriteSpanPoliciesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AtomicOverwriteSpanPoliciesRequest.ProtoReflect.Descriptor instead.
func (*AtomicOverwriteSpanPoliciesRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_rawDescGZIP(), []int{1}
}

func (x *AtomicOverwriteSpanPoliciesRequest) GetPolicies() []*CreateSpanPolicyRequest {
	if x != nil {
		return x.Policies
	}
	return nil
}

type CreateSpanPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Policy    *CreateGenericPolicyRequest `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
	SpanRules *SpanRules                  `protobuf:"bytes,2,opt,name=span_rules,json=spanRules,proto3" json:"span_rules,omitempty"`
}

func (x *CreateSpanPolicyRequest) Reset() {
	*x = CreateSpanPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSpanPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSpanPolicyRequest) ProtoMessage() {}

func (x *CreateSpanPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSpanPolicyRequest.ProtoReflect.Descriptor instead.
func (*CreateSpanPolicyRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSpanPolicyRequest) GetPolicy() *CreateGenericPolicyRequest {
	if x != nil {
		return x.Policy
	}
	return nil
}

func (x *CreateSpanPolicyRequest) GetSpanRules() *SpanRules {
	if x != nil {
		return x.SpanRules
	}
	return nil
}

type CreateLogPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Policy   *CreateGenericPolicyRequest `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
	LogRules *LogRules                   `protobuf:"bytes,2,opt,name=log_rules,json=logRules,proto3" json:"log_rules,omitempty"`
}

func (x *CreateLogPolicyRequest) Reset() {
	*x = CreateLogPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLogPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLogPolicyRequest) ProtoMessage() {}

func (x *CreateLogPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLogPolicyRequest.ProtoReflect.Descriptor instead.
func (*CreateLogPolicyRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_rawDescGZIP(), []int{3}
}

func (x *CreateLogPolicyRequest) GetPolicy() *CreateGenericPolicyRequest {
	if x != nil {
		return x.Policy
	}
	return nil
}

func (x *CreateLogPolicyRequest) GetLogRules() *LogRules {
	if x != nil {
		return x.LogRules
	}
	return nil
}

type CreateGenericPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description      *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Priority         Priority                `protobuf:"varint,3,opt,name=priority,proto3,enum=com.coralogix.quota.v1.Priority" json:"priority,omitempty"`
	ApplicationRule  *Rule                   `protobuf:"bytes,4,opt,name=application_rule,json=applicationRule,proto3,oneof" json:"application_rule,omitempty"`
	SubsystemRule    *Rule                   `protobuf:"bytes,5,opt,name=subsystem_rule,json=subsystemRule,proto3,oneof" json:"subsystem_rule,omitempty"`
	ArchiveRetention *ArchiveRetention       `protobuf:"bytes,6,opt,name=archive_retention,json=archiveRetention,proto3,oneof" json:"archive_retention,omitempty"`
}

func (x *CreateGenericPolicyRequest) Reset() {
	*x = CreateGenericPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGenericPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGenericPolicyRequest) ProtoMessage() {}

func (x *CreateGenericPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGenericPolicyRequest.ProtoReflect.Descriptor instead.
func (*CreateGenericPolicyRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_rawDescGZIP(), []int{4}
}

func (x *CreateGenericPolicyRequest) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *CreateGenericPolicyRequest) GetDescription() *wrapperspb.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *CreateGenericPolicyRequest) GetPriority() Priority {
	if x != nil {
		return x.Priority
	}
	return Priority_PRIORITY_TYPE_UNSPECIFIED
}

func (x *CreateGenericPolicyRequest) GetApplicationRule() *Rule {
	if x != nil {
		return x.ApplicationRule
	}
	return nil
}

func (x *CreateGenericPolicyRequest) GetSubsystemRule() *Rule {
	if x != nil {
		return x.SubsystemRule
	}
	return nil
}

func (x *CreateGenericPolicyRequest) GetArchiveRetention() *ArchiveRetention {
	if x != nil {
		return x.ArchiveRetention
	}
	return nil
}

var File_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto protoreflect.FileDescriptor

var file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x5f,
	0x6f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x71, 0x75, 0x6f, 0x74,
	0x61, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x71,
	0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x5f,
	0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x26, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x71,
	0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x72, 0x75, 0x6c, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x70, 0x61, 0x6e, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x6f, 0x0a, 0x21, 0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x4f, 0x76, 0x65, 0x72, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x22, 0x71, 0x0a, 0x22, 0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x4f, 0x76, 0x65, 0x72, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x53, 0x70, 0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x70, 0x61, 0x6e, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x65, 0x73, 0x22, 0xa7, 0x01, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x70, 0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x4a, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x40, 0x0a, 0x0a,
	0x73, 0x70, 0x61, 0x6e, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x52, 0x75,
	0x6c, 0x65, 0x73, 0x52, 0x09, 0x73, 0x70, 0x61, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x22, 0xa3,
	0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x06, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x06, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x3d, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x72, 0x75, 0x6c,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x52,
	0x75, 0x6c, 0x65, 0x73, 0x22, 0xfe, 0x03, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x12, 0x4c, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75,
	0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x48, 0x0a, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x72,
	0x75, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x48, 0x01, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x52, 0x75, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x5a, 0x0a, 0x11, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x02, 0x52, 0x10, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x52, 0x65, 0x74, 0x65, 0x6e,
	0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x42, 0x11, 0x0a, 0x0f,
	0x5f, 0x73, 0x75, 0x62, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x42,
	0x14, 0x0a, 0x12, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x74, 0x65,
	0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_rawDescOnce sync.Once
	file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_rawDescData = file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_rawDesc
)

func file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_rawDescGZIP() []byte {
	file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_rawDescOnce.Do(func() {
		file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_rawDescData)
	})
	return file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_rawDescData
}

var file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_goTypes = []any{
	(*AtomicOverwriteLogPoliciesRequest)(nil),  // 0: com.coralogix.quota.v1.AtomicOverwriteLogPoliciesRequest
	(*AtomicOverwriteSpanPoliciesRequest)(nil), // 1: com.coralogix.quota.v1.AtomicOverwriteSpanPoliciesRequest
	(*CreateSpanPolicyRequest)(nil),            // 2: com.coralogix.quota.v1.CreateSpanPolicyRequest
	(*CreateLogPolicyRequest)(nil),             // 3: com.coralogix.quota.v1.CreateLogPolicyRequest
	(*CreateGenericPolicyRequest)(nil),         // 4: com.coralogix.quota.v1.CreateGenericPolicyRequest
	(*SpanRules)(nil),                          // 5: com.coralogix.quota.v1.SpanRules
	(*LogRules)(nil),                           // 6: com.coralogix.quota.v1.LogRules
	(*wrapperspb.StringValue)(nil),             // 7: google.protobuf.StringValue
	(Priority)(0),                              // 8: com.coralogix.quota.v1.Priority
	(*Rule)(nil),                               // 9: com.coralogix.quota.v1.Rule
	(*ArchiveRetention)(nil),                   // 10: com.coralogix.quota.v1.ArchiveRetention
}
var file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_depIdxs = []int32{
	3,  // 0: com.coralogix.quota.v1.AtomicOverwriteLogPoliciesRequest.policies:type_name -> com.coralogix.quota.v1.CreateLogPolicyRequest
	2,  // 1: com.coralogix.quota.v1.AtomicOverwriteSpanPoliciesRequest.policies:type_name -> com.coralogix.quota.v1.CreateSpanPolicyRequest
	4,  // 2: com.coralogix.quota.v1.CreateSpanPolicyRequest.policy:type_name -> com.coralogix.quota.v1.CreateGenericPolicyRequest
	5,  // 3: com.coralogix.quota.v1.CreateSpanPolicyRequest.span_rules:type_name -> com.coralogix.quota.v1.SpanRules
	4,  // 4: com.coralogix.quota.v1.CreateLogPolicyRequest.policy:type_name -> com.coralogix.quota.v1.CreateGenericPolicyRequest
	6,  // 5: com.coralogix.quota.v1.CreateLogPolicyRequest.log_rules:type_name -> com.coralogix.quota.v1.LogRules
	7,  // 6: com.coralogix.quota.v1.CreateGenericPolicyRequest.name:type_name -> google.protobuf.StringValue
	7,  // 7: com.coralogix.quota.v1.CreateGenericPolicyRequest.description:type_name -> google.protobuf.StringValue
	8,  // 8: com.coralogix.quota.v1.CreateGenericPolicyRequest.priority:type_name -> com.coralogix.quota.v1.Priority
	9,  // 9: com.coralogix.quota.v1.CreateGenericPolicyRequest.application_rule:type_name -> com.coralogix.quota.v1.Rule
	9,  // 10: com.coralogix.quota.v1.CreateGenericPolicyRequest.subsystem_rule:type_name -> com.coralogix.quota.v1.Rule
	10, // 11: com.coralogix.quota.v1.CreateGenericPolicyRequest.archive_retention:type_name -> com.coralogix.quota.v1.ArchiveRetention
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_init() }
func file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_init() {
	if File_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto != nil {
		return
	}
	file_com_coralogix_quota_v1_enums_proto_init()
	file_com_coralogix_quota_v1_rule_proto_init()
	file_com_coralogix_quota_v1_archive_retention_proto_init()
	file_com_coralogix_quota_v1_log_rules_proto_init()
	file_com_coralogix_quota_v1_span_rules_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AtomicOverwriteLogPoliciesRequest); i {
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
		file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AtomicOverwriteSpanPoliciesRequest); i {
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
		file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateSpanPolicyRequest); i {
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
		file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreateLogPolicyRequest); i {
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
		file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CreateGenericPolicyRequest); i {
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
	file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_msgTypes[4].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_goTypes,
		DependencyIndexes: file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_depIdxs,
		MessageInfos:      file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_msgTypes,
	}.Build()
	File_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto = out.File
	file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_rawDesc = nil
	file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_goTypes = nil
	file_com_coralogix_quota_v1_atomic_overwrite_policies_request_proto_depIdxs = nil
}
