// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: com/coralogix/dataprime/v1/company_config.proto

package v1

import (
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

type GetCompanyConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCompanyConfigRequest) Reset() {
	*x = GetCompanyConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_dataprime_v1_company_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCompanyConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompanyConfigRequest) ProtoMessage() {}

func (x *GetCompanyConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_company_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCompanyConfigRequest.ProtoReflect.Descriptor instead.
func (*GetCompanyConfigRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_company_config_proto_rawDescGZIP(), []int{0}
}

type GetCompanyConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *CompanyConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *GetCompanyConfigResponse) Reset() {
	*x = GetCompanyConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_dataprime_v1_company_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCompanyConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompanyConfigResponse) ProtoMessage() {}

func (x *GetCompanyConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_company_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCompanyConfigResponse.ProtoReflect.Descriptor instead.
func (*GetCompanyConfigResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_company_config_proto_rawDescGZIP(), []int{1}
}

func (x *GetCompanyConfigResponse) GetConfig() *CompanyConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type CompanyConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MigratedCompanies []*MigratedCompany `protobuf:"bytes,1,rep,name=migrated_companies,json=migratedCompanies,proto3" json:"migrated_companies,omitempty"`
	MaxTeamsPerQuery  int32              `protobuf:"varint,2,opt,name=max_teams_per_query,json=maxTeamsPerQuery,proto3" json:"max_teams_per_query,omitempty"`
}

func (x *CompanyConfig) Reset() {
	*x = CompanyConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_dataprime_v1_company_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanyConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyConfig) ProtoMessage() {}

func (x *CompanyConfig) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_company_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyConfig.ProtoReflect.Descriptor instead.
func (*CompanyConfig) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_company_config_proto_rawDescGZIP(), []int{2}
}

func (x *CompanyConfig) GetMigratedCompanies() []*MigratedCompany {
	if x != nil {
		return x.MigratedCompanies
	}
	return nil
}

func (x *CompanyConfig) GetMaxTeamsPerQuery() int32 {
	if x != nil {
		return x.MaxTeamsPerQuery
	}
	return 0
}

type MigratedCompany struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId int32                  `protobuf:"varint,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Cutoff    *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=cutoff,proto3" json:"cutoff,omitempty"`
}

func (x *MigratedCompany) Reset() {
	*x = MigratedCompany{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_dataprime_v1_company_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MigratedCompany) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MigratedCompany) ProtoMessage() {}

func (x *MigratedCompany) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_company_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MigratedCompany.ProtoReflect.Descriptor instead.
func (*MigratedCompany) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_company_config_proto_rawDescGZIP(), []int{3}
}

func (x *MigratedCompany) GetCompanyId() int32 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *MigratedCompany) GetCutoff() *timestamppb.Timestamp {
	if x != nil {
		return x.Cutoff
	}
	return nil
}

var File_com_coralogix_dataprime_v1_company_config_proto protoreflect.FileDescriptor

var file_com_coralogix_dataprime_v1_company_config_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x19,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5d, 0x0a, 0x18, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x9a, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x5a, 0x0a, 0x12, 0x6d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x52, 0x11, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x65,
	0x61, 0x6d, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x10, 0x6d, 0x61, 0x78, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x50, 0x65, 0x72,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x22, 0x64, 0x0a, 0x0f, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x06, 0x63, 0x75, 0x74, 0x6f, 0x66,
	0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x06, 0x63, 0x75, 0x74, 0x6f, 0x66, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_dataprime_v1_company_config_proto_rawDescOnce sync.Once
	file_com_coralogix_dataprime_v1_company_config_proto_rawDescData = file_com_coralogix_dataprime_v1_company_config_proto_rawDesc
)

func file_com_coralogix_dataprime_v1_company_config_proto_rawDescGZIP() []byte {
	file_com_coralogix_dataprime_v1_company_config_proto_rawDescOnce.Do(func() {
		file_com_coralogix_dataprime_v1_company_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_dataprime_v1_company_config_proto_rawDescData)
	})
	return file_com_coralogix_dataprime_v1_company_config_proto_rawDescData
}

var file_com_coralogix_dataprime_v1_company_config_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_com_coralogix_dataprime_v1_company_config_proto_goTypes = []any{
	(*GetCompanyConfigRequest)(nil),  // 0: com.coralogix.dataprime.v1.GetCompanyConfigRequest
	(*GetCompanyConfigResponse)(nil), // 1: com.coralogix.dataprime.v1.GetCompanyConfigResponse
	(*CompanyConfig)(nil),            // 2: com.coralogix.dataprime.v1.CompanyConfig
	(*MigratedCompany)(nil),          // 3: com.coralogix.dataprime.v1.MigratedCompany
	(*timestamppb.Timestamp)(nil),    // 4: google.protobuf.Timestamp
}
var file_com_coralogix_dataprime_v1_company_config_proto_depIdxs = []int32{
	2, // 0: com.coralogix.dataprime.v1.GetCompanyConfigResponse.config:type_name -> com.coralogix.dataprime.v1.CompanyConfig
	3, // 1: com.coralogix.dataprime.v1.CompanyConfig.migrated_companies:type_name -> com.coralogix.dataprime.v1.MigratedCompany
	4, // 2: com.coralogix.dataprime.v1.MigratedCompany.cutoff:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_com_coralogix_dataprime_v1_company_config_proto_init() }
func file_com_coralogix_dataprime_v1_company_config_proto_init() {
	if File_com_coralogix_dataprime_v1_company_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_dataprime_v1_company_config_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetCompanyConfigRequest); i {
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
		file_com_coralogix_dataprime_v1_company_config_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetCompanyConfigResponse); i {
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
		file_com_coralogix_dataprime_v1_company_config_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CompanyConfig); i {
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
		file_com_coralogix_dataprime_v1_company_config_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*MigratedCompany); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_dataprime_v1_company_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_dataprime_v1_company_config_proto_goTypes,
		DependencyIndexes: file_com_coralogix_dataprime_v1_company_config_proto_depIdxs,
		MessageInfos:      file_com_coralogix_dataprime_v1_company_config_proto_msgTypes,
	}.Build()
	File_com_coralogix_dataprime_v1_company_config_proto = out.File
	file_com_coralogix_dataprime_v1_company_config_proto_rawDesc = nil
	file_com_coralogix_dataprime_v1_company_config_proto_goTypes = nil
	file_com_coralogix_dataprime_v1_company_config_proto_depIdxs = nil
}
