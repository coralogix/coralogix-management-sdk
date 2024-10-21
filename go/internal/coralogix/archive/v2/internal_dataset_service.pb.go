// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: com/coralogix/archive/dataset/v2/internal_dataset_service.proto

package v2

import (
	_ "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/archive/v1"
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

type GetAllDatasetRulesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastUpdateTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=last_update_time,json=lastUpdateTime,proto3,oneof" json:"last_update_time,omitempty"`
	ForceUpdate    bool                   `protobuf:"varint,2,opt,name=force_update,json=forceUpdate,proto3" json:"force_update,omitempty"`
}

func (x *GetAllDatasetRulesRequest) Reset() {
	*x = GetAllDatasetRulesRequest{}
	mi := &file_com_coralogix_archive_dataset_v2_internal_dataset_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllDatasetRulesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllDatasetRulesRequest) ProtoMessage() {}

func (x *GetAllDatasetRulesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_dataset_v2_internal_dataset_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllDatasetRulesRequest.ProtoReflect.Descriptor instead.
func (*GetAllDatasetRulesRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_internal_dataset_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllDatasetRulesRequest) GetLastUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdateTime
	}
	return nil
}

func (x *GetAllDatasetRulesRequest) GetForceUpdate() bool {
	if x != nil {
		return x.ForceUpdate
	}
	return false
}

type CompanyDatasetRules struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId uint32         `protobuf:"varint,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Rules     []*DatasetRule `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *CompanyDatasetRules) Reset() {
	*x = CompanyDatasetRules{}
	mi := &file_com_coralogix_archive_dataset_v2_internal_dataset_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompanyDatasetRules) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyDatasetRules) ProtoMessage() {}

func (x *CompanyDatasetRules) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_dataset_v2_internal_dataset_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyDatasetRules.ProtoReflect.Descriptor instead.
func (*CompanyDatasetRules) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_internal_dataset_service_proto_rawDescGZIP(), []int{1}
}

func (x *CompanyDatasetRules) GetCompanyId() uint32 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *CompanyDatasetRules) GetRules() []*DatasetRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

type GetAllDatasetRulesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rules          []*CompanyDatasetRules `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	LastUpdateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_update_time,json=lastUpdateTime,proto3" json:"last_update_time,omitempty"`
}

func (x *GetAllDatasetRulesResponse) Reset() {
	*x = GetAllDatasetRulesResponse{}
	mi := &file_com_coralogix_archive_dataset_v2_internal_dataset_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllDatasetRulesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllDatasetRulesResponse) ProtoMessage() {}

func (x *GetAllDatasetRulesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_dataset_v2_internal_dataset_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllDatasetRulesResponse.ProtoReflect.Descriptor instead.
func (*GetAllDatasetRulesResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_dataset_v2_internal_dataset_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllDatasetRulesResponse) GetRules() []*CompanyDatasetRules {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *GetAllDatasetRulesResponse) GetLastUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdateTime
	}
	return nil
}

var File_com_coralogix_archive_dataset_v2_internal_dataset_service_proto protoreflect.FileDescriptor

var file_com_coralogix_archive_dataset_v2_internal_dataset_service_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2f,
	0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x2e, 0x76, 0x32, 0x1a, 0x34, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x72, 0x75,
	0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x49, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a,
	0x0c, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x42, 0x13, 0x0a, 0x11, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x79, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x43, 0x0a, 0x05, 0x72,
	0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x22, 0xaf, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4b, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76,
	0x32, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x10,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x32, 0xef, 0x01, 0x0a, 0x20, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xca, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x3b,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76,
	0x32, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52,
	0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0xc2, 0xb8, 0x02, 0x35, 0x0a,
	0x33, 0x47, 0x65, 0x74, 0x20, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x20, 0x72, 0x75, 0x6c,
	0x65, 0x73, 0x20, 0x28, 0x62, 0x79, 0x20, 0x6c, 0x61, 0x73, 0x74, 0x20, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x20, 0x74, 0x69, 0x6d, 0x65, 0x2c, 0x20, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x6c, 0x79, 0x29, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_archive_dataset_v2_internal_dataset_service_proto_rawDescOnce sync.Once
	file_com_coralogix_archive_dataset_v2_internal_dataset_service_proto_rawDescData = file_com_coralogix_archive_dataset_v2_internal_dataset_service_proto_rawDesc
)

func file_com_coralogix_archive_dataset_v2_internal_dataset_service_proto_rawDescGZIP() []byte {
	file_com_coralogix_archive_dataset_v2_internal_dataset_service_proto_rawDescOnce.Do(func() {
		file_com_coralogix_archive_dataset_v2_internal_dataset_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_archive_dataset_v2_internal_dataset_service_proto_rawDescData)
	})
	return file_com_coralogix_archive_dataset_v2_internal_dataset_service_proto_rawDescData
}

var file_com_coralogix_archive_dataset_v2_internal_dataset_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogix_archive_dataset_v2_internal_dataset_service_proto_goTypes = []any{
	(*GetAllDatasetRulesRequest)(nil),  // 0: com.coralogix.archive.dataset.v2.GetAllDatasetRulesRequest
	(*CompanyDatasetRules)(nil),        // 1: com.coralogix.archive.dataset.v2.CompanyDatasetRules
	(*GetAllDatasetRulesResponse)(nil), // 2: com.coralogix.archive.dataset.v2.GetAllDatasetRulesResponse
	(*timestamppb.Timestamp)(nil),      // 3: google.protobuf.Timestamp
	(*DatasetRule)(nil),                // 4: com.coralogix.archive.dataset.v2.DatasetRule
}
var file_com_coralogix_archive_dataset_v2_internal_dataset_service_proto_depIdxs = []int32{
	3, // 0: com.coralogix.archive.dataset.v2.GetAllDatasetRulesRequest.last_update_time:type_name -> google.protobuf.Timestamp
	4, // 1: com.coralogix.archive.dataset.v2.CompanyDatasetRules.rules:type_name -> com.coralogix.archive.dataset.v2.DatasetRule
	1, // 2: com.coralogix.archive.dataset.v2.GetAllDatasetRulesResponse.rules:type_name -> com.coralogix.archive.dataset.v2.CompanyDatasetRules
	3, // 3: com.coralogix.archive.dataset.v2.GetAllDatasetRulesResponse.last_update_time:type_name -> google.protobuf.Timestamp
	0, // 4: com.coralogix.archive.dataset.v2.InternalDatasetManagementService.GetAllDatasetRules:input_type -> com.coralogix.archive.dataset.v2.GetAllDatasetRulesRequest
	2, // 5: com.coralogix.archive.dataset.v2.InternalDatasetManagementService.GetAllDatasetRules:output_type -> com.coralogix.archive.dataset.v2.GetAllDatasetRulesResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_com_coralogix_archive_dataset_v2_internal_dataset_service_proto_init() }
func file_com_coralogix_archive_dataset_v2_internal_dataset_service_proto_init() {
	if File_com_coralogix_archive_dataset_v2_internal_dataset_service_proto != nil {
		return
	}
	file_com_coralogix_archive_dataset_v2_dataset_rules_proto_init()
	file_com_coralogix_archive_dataset_v2_internal_dataset_service_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_archive_dataset_v2_internal_dataset_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogix_archive_dataset_v2_internal_dataset_service_proto_goTypes,
		DependencyIndexes: file_com_coralogix_archive_dataset_v2_internal_dataset_service_proto_depIdxs,
		MessageInfos:      file_com_coralogix_archive_dataset_v2_internal_dataset_service_proto_msgTypes,
	}.Build()
	File_com_coralogix_archive_dataset_v2_internal_dataset_service_proto = out.File
	file_com_coralogix_archive_dataset_v2_internal_dataset_service_proto_rawDesc = nil
	file_com_coralogix_archive_dataset_v2_internal_dataset_service_proto_goTypes = nil
	file_com_coralogix_archive_dataset_v2_internal_dataset_service_proto_depIdxs = nil
}
