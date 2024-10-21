// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: com/coralogixapis/service_catalog/v1/report_service.proto

package v1

import (
	_ "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/apm/common/v2"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/structpb"
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

type GetReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
}

func (x *GetReportRequest) Reset() {
	*x = GetReportRequest{}
	mi := &file_com_coralogixapis_service_catalog_v1_report_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReportRequest) ProtoMessage() {}

func (x *GetReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_service_catalog_v1_report_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReportRequest.ProtoReflect.Descriptor instead.
func (*GetReportRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_service_catalog_v1_report_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetReportRequest) GetCompanyId() *wrapperspb.StringValue {
	if x != nil {
		return x.CompanyId
	}
	return nil
}

type GetReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyReport []*CompanyReport `protobuf:"bytes,1,rep,name=company_report,json=companyReport,proto3" json:"company_report,omitempty"`
}

func (x *GetReportResponse) Reset() {
	*x = GetReportResponse{}
	mi := &file_com_coralogixapis_service_catalog_v1_report_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReportResponse) ProtoMessage() {}

func (x *GetReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_service_catalog_v1_report_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReportResponse.ProtoReflect.Descriptor instead.
func (*GetReportResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_service_catalog_v1_report_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetReportResponse) GetCompanyReport() []*CompanyReport {
	if x != nil {
		return x.CompanyReport
	}
	return nil
}

type CompanyReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId                     *wrapperspb.Int64Value  `protobuf:"bytes,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	CompanyName                   *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=company_name,json=companyName,proto3" json:"company_name,omitempty"`
	Environment                   *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=environment,proto3" json:"environment,omitempty"`
	ServiceRetentionPeriod        *wrapperspb.Int64Value  `protobuf:"bytes,4,opt,name=service_retention_period,json=serviceRetentionPeriod,proto3" json:"service_retention_period,omitempty"`
	ApmSource                     ApmSource               `protobuf:"varint,5,opt,name=apm_source,json=apmSource,proto3,enum=com.coralogixapis.service_catalog.v1.ApmSource" json:"apm_source,omitempty"`
	MigrationPeriod               *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=migration_period,json=migrationPeriod,proto3" json:"migration_period,omitempty"`
	ServiceCount                  *wrapperspb.Int64Value  `protobuf:"bytes,7,opt,name=service_count,json=serviceCount,proto3" json:"service_count,omitempty"`
	ActiveServiceCount            *wrapperspb.Int64Value  `protobuf:"bytes,8,opt,name=active_service_count,json=activeServiceCount,proto3" json:"active_service_count,omitempty"`
	DatabaseCount                 *wrapperspb.Int64Value  `protobuf:"bytes,9,opt,name=database_count,json=databaseCount,proto3" json:"database_count,omitempty"`
	ActiveDatabaseCount           *wrapperspb.Int64Value  `protobuf:"bytes,10,opt,name=active_database_count,json=activeDatabaseCount,proto3" json:"active_database_count,omitempty"`
	ServerlessCount               *wrapperspb.Int64Value  `protobuf:"bytes,11,opt,name=serverless_count,json=serverlessCount,proto3" json:"serverless_count,omitempty"`
	ServiceFlowsCount             *wrapperspb.Int64Value  `protobuf:"bytes,12,opt,name=service_flows_count,json=serviceFlowsCount,proto3" json:"service_flows_count,omitempty"`
	ServiceFlowsInstrumentedCount *wrapperspb.Int64Value  `protobuf:"bytes,13,opt,name=service_flows_instrumented_count,json=serviceFlowsInstrumentedCount,proto3" json:"service_flows_instrumented_count,omitempty"`
	SloCount                      *wrapperspb.Int64Value  `protobuf:"bytes,14,opt,name=slo_count,json=sloCount,proto3" json:"slo_count,omitempty"`
	AlertCount                    *wrapperspb.Int64Value  `protobuf:"bytes,15,opt,name=alert_count,json=alertCount,proto3" json:"alert_count,omitempty"`
	DbRetentionPeriod             *wrapperspb.Int64Value  `protobuf:"bytes,16,opt,name=db_retention_period,json=dbRetentionPeriod,proto3" json:"db_retention_period,omitempty"`
	DbApmSource                   ApmSource               `protobuf:"varint,17,opt,name=db_apm_source,json=dbApmSource,proto3,enum=com.coralogixapis.service_catalog.v1.ApmSource" json:"db_apm_source,omitempty"`
	DbMigrationPeriod             *wrapperspb.StringValue `protobuf:"bytes,18,opt,name=db_migration_period,json=dbMigrationPeriod,proto3" json:"db_migration_period,omitempty"`
}

func (x *CompanyReport) Reset() {
	*x = CompanyReport{}
	mi := &file_com_coralogixapis_service_catalog_v1_report_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompanyReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyReport) ProtoMessage() {}

func (x *CompanyReport) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_service_catalog_v1_report_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyReport.ProtoReflect.Descriptor instead.
func (*CompanyReport) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_service_catalog_v1_report_service_proto_rawDescGZIP(), []int{2}
}

func (x *CompanyReport) GetCompanyId() *wrapperspb.Int64Value {
	if x != nil {
		return x.CompanyId
	}
	return nil
}

func (x *CompanyReport) GetCompanyName() *wrapperspb.StringValue {
	if x != nil {
		return x.CompanyName
	}
	return nil
}

func (x *CompanyReport) GetEnvironment() *wrapperspb.StringValue {
	if x != nil {
		return x.Environment
	}
	return nil
}

func (x *CompanyReport) GetServiceRetentionPeriod() *wrapperspb.Int64Value {
	if x != nil {
		return x.ServiceRetentionPeriod
	}
	return nil
}

func (x *CompanyReport) GetApmSource() ApmSource {
	if x != nil {
		return x.ApmSource
	}
	return ApmSource_APM_SOURCE_UNSPECIFIED
}

func (x *CompanyReport) GetMigrationPeriod() *wrapperspb.StringValue {
	if x != nil {
		return x.MigrationPeriod
	}
	return nil
}

func (x *CompanyReport) GetServiceCount() *wrapperspb.Int64Value {
	if x != nil {
		return x.ServiceCount
	}
	return nil
}

func (x *CompanyReport) GetActiveServiceCount() *wrapperspb.Int64Value {
	if x != nil {
		return x.ActiveServiceCount
	}
	return nil
}

func (x *CompanyReport) GetDatabaseCount() *wrapperspb.Int64Value {
	if x != nil {
		return x.DatabaseCount
	}
	return nil
}

func (x *CompanyReport) GetActiveDatabaseCount() *wrapperspb.Int64Value {
	if x != nil {
		return x.ActiveDatabaseCount
	}
	return nil
}

func (x *CompanyReport) GetServerlessCount() *wrapperspb.Int64Value {
	if x != nil {
		return x.ServerlessCount
	}
	return nil
}

func (x *CompanyReport) GetServiceFlowsCount() *wrapperspb.Int64Value {
	if x != nil {
		return x.ServiceFlowsCount
	}
	return nil
}

func (x *CompanyReport) GetServiceFlowsInstrumentedCount() *wrapperspb.Int64Value {
	if x != nil {
		return x.ServiceFlowsInstrumentedCount
	}
	return nil
}

func (x *CompanyReport) GetSloCount() *wrapperspb.Int64Value {
	if x != nil {
		return x.SloCount
	}
	return nil
}

func (x *CompanyReport) GetAlertCount() *wrapperspb.Int64Value {
	if x != nil {
		return x.AlertCount
	}
	return nil
}

func (x *CompanyReport) GetDbRetentionPeriod() *wrapperspb.Int64Value {
	if x != nil {
		return x.DbRetentionPeriod
	}
	return nil
}

func (x *CompanyReport) GetDbApmSource() ApmSource {
	if x != nil {
		return x.DbApmSource
	}
	return ApmSource_APM_SOURCE_UNSPECIFIED
}

func (x *CompanyReport) GetDbMigrationPeriod() *wrapperspb.StringValue {
	if x != nil {
		return x.DbMigrationPeriod
	}
	return nil
}

var File_com_coralogixapis_service_catalog_v1_report_service_proto protoreflect.FileDescriptor

var file_com_coralogixapis_service_catalog_v1_report_service_proto_rawDesc = []byte{
	0x0a, 0x39, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76,
	0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x34, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x6d, 0x5f,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x70, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3b, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x22, 0x6f,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22,
	0xc5, 0x0a, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x3f, 0x0a,
	0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3e,
	0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x55,
	0x0a, 0x18, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x16, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x4e, 0x0a, 0x0a, 0x61, 0x70, 0x6d, 0x5f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x70, 0x6d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x09, 0x61, 0x70, 0x6d, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x10, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x6d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x40,
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x4d, 0x0a, 0x14, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x12, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x42, 0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x4f, 0x0a, 0x15, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x13, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x46, 0x0a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65,
	0x73, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x4b, 0x0a, 0x13,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36,
	0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x11, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x46,
	0x6c, 0x6f, 0x77, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x64, 0x0a, 0x20, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x1d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x46, 0x6c, 0x6f, 0x77, 0x73, 0x49, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x38, 0x0a, 0x09, 0x73, 0x6c, 0x6f, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x08, 0x73, 0x6c, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x0b, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x4b, 0x0a, 0x13, 0x64, 0x62, 0x5f, 0x72, 0x65,
	0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x11, 0x64, 0x62, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x12, 0x53, 0x0a, 0x0d, 0x64, 0x62, 0x5f, 0x61, 0x70, 0x6d, 0x5f, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x70, 0x6d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0b, 0x64, 0x62,
	0x41, 0x70, 0x6d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x13, 0x64, 0x62, 0x5f,
	0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x11, 0x64, 0x62, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x32, 0xba, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa8, 0x01, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0xc2, 0xb8, 0x02, 0x26, 0x0a, 0x24,
	0x47, 0x65, 0x74, 0x20, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x20, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_service_catalog_v1_report_service_proto_rawDescOnce sync.Once
	file_com_coralogixapis_service_catalog_v1_report_service_proto_rawDescData = file_com_coralogixapis_service_catalog_v1_report_service_proto_rawDesc
)

func file_com_coralogixapis_service_catalog_v1_report_service_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_service_catalog_v1_report_service_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_service_catalog_v1_report_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_service_catalog_v1_report_service_proto_rawDescData)
	})
	return file_com_coralogixapis_service_catalog_v1_report_service_proto_rawDescData
}

var file_com_coralogixapis_service_catalog_v1_report_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogixapis_service_catalog_v1_report_service_proto_goTypes = []any{
	(*GetReportRequest)(nil),       // 0: com.coralogixapis.service_catalog.v1.GetReportRequest
	(*GetReportResponse)(nil),      // 1: com.coralogixapis.service_catalog.v1.GetReportResponse
	(*CompanyReport)(nil),          // 2: com.coralogixapis.service_catalog.v1.CompanyReport
	(*wrapperspb.StringValue)(nil), // 3: google.protobuf.StringValue
	(*wrapperspb.Int64Value)(nil),  // 4: google.protobuf.Int64Value
	(ApmSource)(0),                 // 5: com.coralogixapis.service_catalog.v1.ApmSource
}
var file_com_coralogixapis_service_catalog_v1_report_service_proto_depIdxs = []int32{
	3,  // 0: com.coralogixapis.service_catalog.v1.GetReportRequest.company_id:type_name -> google.protobuf.StringValue
	2,  // 1: com.coralogixapis.service_catalog.v1.GetReportResponse.company_report:type_name -> com.coralogixapis.service_catalog.v1.CompanyReport
	4,  // 2: com.coralogixapis.service_catalog.v1.CompanyReport.company_id:type_name -> google.protobuf.Int64Value
	3,  // 3: com.coralogixapis.service_catalog.v1.CompanyReport.company_name:type_name -> google.protobuf.StringValue
	3,  // 4: com.coralogixapis.service_catalog.v1.CompanyReport.environment:type_name -> google.protobuf.StringValue
	4,  // 5: com.coralogixapis.service_catalog.v1.CompanyReport.service_retention_period:type_name -> google.protobuf.Int64Value
	5,  // 6: com.coralogixapis.service_catalog.v1.CompanyReport.apm_source:type_name -> com.coralogixapis.service_catalog.v1.ApmSource
	3,  // 7: com.coralogixapis.service_catalog.v1.CompanyReport.migration_period:type_name -> google.protobuf.StringValue
	4,  // 8: com.coralogixapis.service_catalog.v1.CompanyReport.service_count:type_name -> google.protobuf.Int64Value
	4,  // 9: com.coralogixapis.service_catalog.v1.CompanyReport.active_service_count:type_name -> google.protobuf.Int64Value
	4,  // 10: com.coralogixapis.service_catalog.v1.CompanyReport.database_count:type_name -> google.protobuf.Int64Value
	4,  // 11: com.coralogixapis.service_catalog.v1.CompanyReport.active_database_count:type_name -> google.protobuf.Int64Value
	4,  // 12: com.coralogixapis.service_catalog.v1.CompanyReport.serverless_count:type_name -> google.protobuf.Int64Value
	4,  // 13: com.coralogixapis.service_catalog.v1.CompanyReport.service_flows_count:type_name -> google.protobuf.Int64Value
	4,  // 14: com.coralogixapis.service_catalog.v1.CompanyReport.service_flows_instrumented_count:type_name -> google.protobuf.Int64Value
	4,  // 15: com.coralogixapis.service_catalog.v1.CompanyReport.slo_count:type_name -> google.protobuf.Int64Value
	4,  // 16: com.coralogixapis.service_catalog.v1.CompanyReport.alert_count:type_name -> google.protobuf.Int64Value
	4,  // 17: com.coralogixapis.service_catalog.v1.CompanyReport.db_retention_period:type_name -> google.protobuf.Int64Value
	5,  // 18: com.coralogixapis.service_catalog.v1.CompanyReport.db_apm_source:type_name -> com.coralogixapis.service_catalog.v1.ApmSource
	3,  // 19: com.coralogixapis.service_catalog.v1.CompanyReport.db_migration_period:type_name -> google.protobuf.StringValue
	0,  // 20: com.coralogixapis.service_catalog.v1.ReportService.GetReport:input_type -> com.coralogixapis.service_catalog.v1.GetReportRequest
	1,  // 21: com.coralogixapis.service_catalog.v1.ReportService.GetReport:output_type -> com.coralogixapis.service_catalog.v1.GetReportResponse
	21, // [21:22] is the sub-list for method output_type
	20, // [20:21] is the sub-list for method input_type
	20, // [20:20] is the sub-list for extension type_name
	20, // [20:20] is the sub-list for extension extendee
	0,  // [0:20] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_service_catalog_v1_report_service_proto_init() }
func file_com_coralogixapis_service_catalog_v1_report_service_proto_init() {
	if File_com_coralogixapis_service_catalog_v1_report_service_proto != nil {
		return
	}
	file_com_coralogixapis_service_catalog_v1_operation_proto_init()
	file_com_coralogixapis_service_catalog_v1_apm_source_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_service_catalog_v1_report_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogixapis_service_catalog_v1_report_service_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_service_catalog_v1_report_service_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_service_catalog_v1_report_service_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_service_catalog_v1_report_service_proto = out.File
	file_com_coralogixapis_service_catalog_v1_report_service_proto_rawDesc = nil
	file_com_coralogixapis_service_catalog_v1_report_service_proto_goTypes = nil
	file_com_coralogixapis_service_catalog_v1_report_service_proto_depIdxs = nil
}
