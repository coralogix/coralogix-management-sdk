// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.26.1
// source: com/coralogix/quota/v1/update_policy_request.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type UpdatePolicyRequest struct {
	state            protoimpl.MessageState  `protogen:"open.v1"`
	Id               *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name             *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description      *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Priority         Priority                `protobuf:"varint,4,opt,name=priority,proto3,enum=com.coralogix.quota.v1.Priority" json:"priority,omitempty"`
	ApplicationRule  *Rule                   `protobuf:"bytes,5,opt,name=application_rule,json=applicationRule,proto3,oneof" json:"application_rule,omitempty"`
	SubsystemRule    *Rule                   `protobuf:"bytes,6,opt,name=subsystem_rule,json=subsystemRule,proto3,oneof" json:"subsystem_rule,omitempty"`
	ArchiveRetention *ArchiveRetention       `protobuf:"bytes,7,opt,name=archive_retention,json=archiveRetention,proto3,oneof" json:"archive_retention,omitempty"`
	// Types that are valid to be assigned to SourceTypeRules:
	//
	//	*UpdatePolicyRequest_LogRules
	//	*UpdatePolicyRequest_SpanRules
	SourceTypeRules isUpdatePolicyRequest_SourceTypeRules `protobuf_oneof:"source_type_rules"`
	Enabled         *wrapperspb.BoolValue                 `protobuf:"bytes,10,opt,name=enabled,proto3" json:"enabled,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *UpdatePolicyRequest) Reset() {
	*x = UpdatePolicyRequest{}
	mi := &file_com_coralogix_quota_v1_update_policy_request_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePolicyRequest) ProtoMessage() {}

func (x *UpdatePolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_quota_v1_update_policy_request_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePolicyRequest.ProtoReflect.Descriptor instead.
func (*UpdatePolicyRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_quota_v1_update_policy_request_proto_rawDescGZIP(), []int{0}
}

func (x *UpdatePolicyRequest) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *UpdatePolicyRequest) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *UpdatePolicyRequest) GetDescription() *wrapperspb.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *UpdatePolicyRequest) GetPriority() Priority {
	if x != nil {
		return x.Priority
	}
	return Priority_PRIORITY_TYPE_UNSPECIFIED
}

func (x *UpdatePolicyRequest) GetApplicationRule() *Rule {
	if x != nil {
		return x.ApplicationRule
	}
	return nil
}

func (x *UpdatePolicyRequest) GetSubsystemRule() *Rule {
	if x != nil {
		return x.SubsystemRule
	}
	return nil
}

func (x *UpdatePolicyRequest) GetArchiveRetention() *ArchiveRetention {
	if x != nil {
		return x.ArchiveRetention
	}
	return nil
}

func (x *UpdatePolicyRequest) GetSourceTypeRules() isUpdatePolicyRequest_SourceTypeRules {
	if x != nil {
		return x.SourceTypeRules
	}
	return nil
}

func (x *UpdatePolicyRequest) GetLogRules() *LogRules {
	if x != nil {
		if x, ok := x.SourceTypeRules.(*UpdatePolicyRequest_LogRules); ok {
			return x.LogRules
		}
	}
	return nil
}

func (x *UpdatePolicyRequest) GetSpanRules() *SpanRules {
	if x != nil {
		if x, ok := x.SourceTypeRules.(*UpdatePolicyRequest_SpanRules); ok {
			return x.SpanRules
		}
	}
	return nil
}

func (x *UpdatePolicyRequest) GetEnabled() *wrapperspb.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

type isUpdatePolicyRequest_SourceTypeRules interface {
	isUpdatePolicyRequest_SourceTypeRules()
}

type UpdatePolicyRequest_LogRules struct {
	LogRules *LogRules `protobuf:"bytes,8,opt,name=log_rules,json=logRules,proto3,oneof"`
}

type UpdatePolicyRequest_SpanRules struct {
	SpanRules *SpanRules `protobuf:"bytes,9,opt,name=span_rules,json=spanRules,proto3,oneof"`
}

func (*UpdatePolicyRequest_LogRules) isUpdatePolicyRequest_SourceTypeRules() {}

func (*UpdatePolicyRequest_SpanRules) isUpdatePolicyRequest_SourceTypeRules() {}

var File_com_coralogix_quota_v1_update_policy_request_proto protoreflect.FileDescriptor

var file_com_coralogix_quota_v1_update_policy_request_proto_rawDesc = []byte{
	0x0a, 0x32, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x71, 0x75, 0x6f, 0x74,
	0x61, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x21, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x5f,
	0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x08, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x10, 0x92, 0x41, 0x0d, 0x4a, 0x0b, 0x22, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x42, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x10, 0x92, 0x41, 0x0d, 0x4a, 0x0b, 0x22,
	0x4d, 0x79, 0x20, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x5c, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x1c, 0x92, 0x41, 0x19, 0x4a, 0x17, 0x22, 0x4d, 0x79, 0x20, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x20, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c,
	0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x4c, 0x0a, 0x10,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x75, 0x6c, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x75, 0x6c, 0x65, 0x48, 0x01, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x48, 0x0a, 0x0e, 0x73, 0x75,
	0x62, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6c, 0x65,
	0x48, 0x02, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x75, 0x6c,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x5a, 0x0a, 0x11, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x5f,
	0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x03, 0x52, 0x10, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01,
	0x12, 0x3f, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67,
	0x52, 0x75, 0x6c, 0x65, 0x73, 0x48, 0x00, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x52, 0x75, 0x6c, 0x65,
	0x73, 0x12, 0x42, 0x0a, 0x0a, 0x73, 0x70, 0x61, 0x6e, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x70, 0x61, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x48, 0x00, 0x52, 0x09, 0x73, 0x70, 0x61, 0x6e,
	0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x09, 0x92, 0x41, 0x06, 0x4a, 0x04, 0x74, 0x72, 0x75, 0x65, 0x52, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x3a, 0xe7, 0x01, 0x92, 0x41, 0xe3, 0x01, 0x0a, 0x56, 0x2a,
	0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x20, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x38, 0x54, 0x68, 0x69, 0x73, 0x20, 0x64, 0x61, 0x74,
	0x61, 0x20, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x65, 0x20, 0x69, 0x73, 0x20, 0x75, 0x73,
	0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x6e, 0x20,
	0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0xd2, 0x01, 0x02, 0x69, 0x64, 0x2a, 0x88, 0x01, 0x0a, 0x25, 0x46, 0x69, 0x6e, 0x64, 0x20, 0x6f,
	0x75, 0x74, 0x20, 0x6d, 0x6f, 0x72, 0x65, 0x20, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x20, 0x71, 0x75,
	0x6f, 0x74, 0x61, 0x20, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x12,
	0x5f, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x63, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2d, 0x67, 0x75, 0x69, 0x64, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2d,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2d, 0x61, 0x6e, 0x64, 0x2d, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x71,
	0x75, 0x6f, 0x74, 0x61, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x42, 0x13, 0x0a, 0x11, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x72, 0x75, 0x6c, 0x65, 0x73, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x73,
	0x75, 0x62, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x42, 0x14, 0x0a,
	0x12, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74,
	0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_quota_v1_update_policy_request_proto_rawDescOnce sync.Once
	file_com_coralogix_quota_v1_update_policy_request_proto_rawDescData = file_com_coralogix_quota_v1_update_policy_request_proto_rawDesc
)

func file_com_coralogix_quota_v1_update_policy_request_proto_rawDescGZIP() []byte {
	file_com_coralogix_quota_v1_update_policy_request_proto_rawDescOnce.Do(func() {
		file_com_coralogix_quota_v1_update_policy_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_quota_v1_update_policy_request_proto_rawDescData)
	})
	return file_com_coralogix_quota_v1_update_policy_request_proto_rawDescData
}

var file_com_coralogix_quota_v1_update_policy_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogix_quota_v1_update_policy_request_proto_goTypes = []any{
	(*UpdatePolicyRequest)(nil),    // 0: com.coralogix.quota.v1.UpdatePolicyRequest
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
	(Priority)(0),                  // 2: com.coralogix.quota.v1.Priority
	(*Rule)(nil),                   // 3: com.coralogix.quota.v1.Rule
	(*ArchiveRetention)(nil),       // 4: com.coralogix.quota.v1.ArchiveRetention
	(*LogRules)(nil),               // 5: com.coralogix.quota.v1.LogRules
	(*SpanRules)(nil),              // 6: com.coralogix.quota.v1.SpanRules
	(*wrapperspb.BoolValue)(nil),   // 7: google.protobuf.BoolValue
}
var file_com_coralogix_quota_v1_update_policy_request_proto_depIdxs = []int32{
	1,  // 0: com.coralogix.quota.v1.UpdatePolicyRequest.id:type_name -> google.protobuf.StringValue
	1,  // 1: com.coralogix.quota.v1.UpdatePolicyRequest.name:type_name -> google.protobuf.StringValue
	1,  // 2: com.coralogix.quota.v1.UpdatePolicyRequest.description:type_name -> google.protobuf.StringValue
	2,  // 3: com.coralogix.quota.v1.UpdatePolicyRequest.priority:type_name -> com.coralogix.quota.v1.Priority
	3,  // 4: com.coralogix.quota.v1.UpdatePolicyRequest.application_rule:type_name -> com.coralogix.quota.v1.Rule
	3,  // 5: com.coralogix.quota.v1.UpdatePolicyRequest.subsystem_rule:type_name -> com.coralogix.quota.v1.Rule
	4,  // 6: com.coralogix.quota.v1.UpdatePolicyRequest.archive_retention:type_name -> com.coralogix.quota.v1.ArchiveRetention
	5,  // 7: com.coralogix.quota.v1.UpdatePolicyRequest.log_rules:type_name -> com.coralogix.quota.v1.LogRules
	6,  // 8: com.coralogix.quota.v1.UpdatePolicyRequest.span_rules:type_name -> com.coralogix.quota.v1.SpanRules
	7,  // 9: com.coralogix.quota.v1.UpdatePolicyRequest.enabled:type_name -> google.protobuf.BoolValue
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_com_coralogix_quota_v1_update_policy_request_proto_init() }
func file_com_coralogix_quota_v1_update_policy_request_proto_init() {
	if File_com_coralogix_quota_v1_update_policy_request_proto != nil {
		return
	}
	file_com_coralogix_quota_v1_enums_proto_init()
	file_com_coralogix_quota_v1_rule_proto_init()
	file_com_coralogix_quota_v1_archive_retention_proto_init()
	file_com_coralogix_quota_v1_log_rules_proto_init()
	file_com_coralogix_quota_v1_span_rules_proto_init()
	file_com_coralogix_quota_v1_update_policy_request_proto_msgTypes[0].OneofWrappers = []any{
		(*UpdatePolicyRequest_LogRules)(nil),
		(*UpdatePolicyRequest_SpanRules)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_quota_v1_update_policy_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_quota_v1_update_policy_request_proto_goTypes,
		DependencyIndexes: file_com_coralogix_quota_v1_update_policy_request_proto_depIdxs,
		MessageInfos:      file_com_coralogix_quota_v1_update_policy_request_proto_msgTypes,
	}.Build()
	File_com_coralogix_quota_v1_update_policy_request_proto = out.File
	file_com_coralogix_quota_v1_update_policy_request_proto_rawDesc = nil
	file_com_coralogix_quota_v1_update_policy_request_proto_goTypes = nil
	file_com_coralogix_quota_v1_update_policy_request_proto_depIdxs = nil
}
