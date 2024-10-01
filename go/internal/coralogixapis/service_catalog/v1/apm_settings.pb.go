// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: com/coralogixapis/service_catalog/v1/apm_settings.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type CatalogSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source                 ApmSource              `protobuf:"varint,1,opt,name=source,proto3,enum=com.coralogixapis.service_catalog.v1.ApmSource" json:"source,omitempty"`
	MigrationPeriodEndDate *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=migration_period_end_date,json=migrationPeriodEndDate,proto3" json:"migration_period_end_date,omitempty"`
	RetentionPolicy        *wrapperspb.Int32Value `protobuf:"bytes,3,opt,name=retention_policy,json=retentionPolicy,proto3" json:"retention_policy,omitempty"`
	Catalog                Catalog                `protobuf:"varint,4,opt,name=catalog,proto3,enum=com.coralogixapis.service_catalog.v1.Catalog" json:"catalog,omitempty"`
}

func (x *CatalogSettings) Reset() {
	*x = CatalogSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_service_catalog_v1_apm_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CatalogSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CatalogSettings) ProtoMessage() {}

func (x *CatalogSettings) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_service_catalog_v1_apm_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CatalogSettings.ProtoReflect.Descriptor instead.
func (*CatalogSettings) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_service_catalog_v1_apm_settings_proto_rawDescGZIP(), []int{0}
}

func (x *CatalogSettings) GetSource() ApmSource {
	if x != nil {
		return x.Source
	}
	return ApmSource_APM_SOURCE_UNSPECIFIED
}

func (x *CatalogSettings) GetMigrationPeriodEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.MigrationPeriodEndDate
	}
	return nil
}

func (x *CatalogSettings) GetRetentionPolicy() *wrapperspb.Int32Value {
	if x != nil {
		return x.RetentionPolicy
	}
	return nil
}

func (x *CatalogSettings) GetCatalog() Catalog {
	if x != nil {
		return x.Catalog
	}
	return Catalog_CATALOG_UNSPECIFIED
}

type ApmSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CatalogSettings []*CatalogSettings `protobuf:"bytes,1,rep,name=catalog_settings,json=catalogSettings,proto3" json:"catalog_settings,omitempty"`
}

func (x *ApmSettings) Reset() {
	*x = ApmSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_service_catalog_v1_apm_settings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApmSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApmSettings) ProtoMessage() {}

func (x *ApmSettings) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_service_catalog_v1_apm_settings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApmSettings.ProtoReflect.Descriptor instead.
func (*ApmSettings) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_service_catalog_v1_apm_settings_proto_rawDescGZIP(), []int{1}
}

func (x *ApmSettings) GetCatalogSettings() []*CatalogSettings {
	if x != nil {
		return x.CatalogSettings
	}
	return nil
}

var File_com_coralogixapis_service_catalog_v1_apm_settings_proto protoreflect.FileDescriptor

var file_com_coralogixapis_service_catalog_v1_apm_settings_proto_rawDesc = []byte{
	0x0a, 0x37, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x6d, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x35, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x6d, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x02, 0x0a, 0x0f,
	0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x47, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x6d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x19, 0x6d, 0x69, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x65, 0x6e, 0x64,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x16, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x46, 0x0a, 0x10, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x47, 0x0a, 0x07, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x07, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x22, 0x6f, 0x0a, 0x0b, 0x41, 0x70, 0x6d, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x60, 0x0a, 0x10, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x0f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_service_catalog_v1_apm_settings_proto_rawDescOnce sync.Once
	file_com_coralogixapis_service_catalog_v1_apm_settings_proto_rawDescData = file_com_coralogixapis_service_catalog_v1_apm_settings_proto_rawDesc
)

func file_com_coralogixapis_service_catalog_v1_apm_settings_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_service_catalog_v1_apm_settings_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_service_catalog_v1_apm_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_service_catalog_v1_apm_settings_proto_rawDescData)
	})
	return file_com_coralogixapis_service_catalog_v1_apm_settings_proto_rawDescData
}

var file_com_coralogixapis_service_catalog_v1_apm_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogixapis_service_catalog_v1_apm_settings_proto_goTypes = []any{
	(*CatalogSettings)(nil),       // 0: com.coralogixapis.service_catalog.v1.CatalogSettings
	(*ApmSettings)(nil),           // 1: com.coralogixapis.service_catalog.v1.ApmSettings
	(ApmSource)(0),                // 2: com.coralogixapis.service_catalog.v1.ApmSource
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*wrapperspb.Int32Value)(nil), // 4: google.protobuf.Int32Value
	(Catalog)(0),                  // 5: com.coralogixapis.service_catalog.v1.Catalog
}
var file_com_coralogixapis_service_catalog_v1_apm_settings_proto_depIdxs = []int32{
	2, // 0: com.coralogixapis.service_catalog.v1.CatalogSettings.source:type_name -> com.coralogixapis.service_catalog.v1.ApmSource
	3, // 1: com.coralogixapis.service_catalog.v1.CatalogSettings.migration_period_end_date:type_name -> google.protobuf.Timestamp
	4, // 2: com.coralogixapis.service_catalog.v1.CatalogSettings.retention_policy:type_name -> google.protobuf.Int32Value
	5, // 3: com.coralogixapis.service_catalog.v1.CatalogSettings.catalog:type_name -> com.coralogixapis.service_catalog.v1.Catalog
	0, // 4: com.coralogixapis.service_catalog.v1.ApmSettings.catalog_settings:type_name -> com.coralogixapis.service_catalog.v1.CatalogSettings
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_service_catalog_v1_apm_settings_proto_init() }
func file_com_coralogixapis_service_catalog_v1_apm_settings_proto_init() {
	if File_com_coralogixapis_service_catalog_v1_apm_settings_proto != nil {
		return
	}
	file_com_coralogixapis_service_catalog_v1_apm_source_proto_init()
	file_com_coralogixapis_service_catalog_v1_catalog_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_service_catalog_v1_apm_settings_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CatalogSettings); i {
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
		file_com_coralogixapis_service_catalog_v1_apm_settings_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ApmSettings); i {
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
			RawDescriptor: file_com_coralogixapis_service_catalog_v1_apm_settings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_service_catalog_v1_apm_settings_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_service_catalog_v1_apm_settings_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_service_catalog_v1_apm_settings_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_service_catalog_v1_apm_settings_proto = out.File
	file_com_coralogixapis_service_catalog_v1_apm_settings_proto_rawDesc = nil
	file_com_coralogixapis_service_catalog_v1_apm_settings_proto_goTypes = nil
	file_com_coralogixapis_service_catalog_v1_apm_settings_proto_depIdxs = nil
}
