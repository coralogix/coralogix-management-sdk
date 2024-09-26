// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.28.1
// source: com/coralogix/enrichment/v1/enrichment_type.proto

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

type EnrichmentType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*EnrichmentType_GeoIp
	//	*EnrichmentType_SuspiciousIp
	//	*EnrichmentType_Aws
	//	*EnrichmentType_CustomEnrichment
	Type isEnrichmentType_Type `protobuf_oneof:"type"`
}

func (x *EnrichmentType) Reset() {
	*x = EnrichmentType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_enrichment_v1_enrichment_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrichmentType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrichmentType) ProtoMessage() {}

func (x *EnrichmentType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnrichmentType.ProtoReflect.Descriptor instead.
func (*EnrichmentType) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_type_proto_rawDescGZIP(), []int{0}
}

func (m *EnrichmentType) GetType() isEnrichmentType_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *EnrichmentType) GetGeoIp() *GeoIpType {
	if x, ok := x.GetType().(*EnrichmentType_GeoIp); ok {
		return x.GeoIp
	}
	return nil
}

func (x *EnrichmentType) GetSuspiciousIp() *SuspiciousIpType {
	if x, ok := x.GetType().(*EnrichmentType_SuspiciousIp); ok {
		return x.SuspiciousIp
	}
	return nil
}

func (x *EnrichmentType) GetAws() *AwsType {
	if x, ok := x.GetType().(*EnrichmentType_Aws); ok {
		return x.Aws
	}
	return nil
}

func (x *EnrichmentType) GetCustomEnrichment() *CustomEnrichmentType {
	if x, ok := x.GetType().(*EnrichmentType_CustomEnrichment); ok {
		return x.CustomEnrichment
	}
	return nil
}

type isEnrichmentType_Type interface {
	isEnrichmentType_Type()
}

type EnrichmentType_GeoIp struct {
	GeoIp *GeoIpType `protobuf:"bytes,1,opt,name=geo_ip,json=geoIp,proto3,oneof"`
}

type EnrichmentType_SuspiciousIp struct {
	SuspiciousIp *SuspiciousIpType `protobuf:"bytes,2,opt,name=suspicious_ip,json=suspiciousIp,proto3,oneof"`
}

type EnrichmentType_Aws struct {
	Aws *AwsType `protobuf:"bytes,3,opt,name=aws,proto3,oneof"`
}

type EnrichmentType_CustomEnrichment struct {
	CustomEnrichment *CustomEnrichmentType `protobuf:"bytes,4,opt,name=custom_enrichment,json=customEnrichment,proto3,oneof"`
}

func (*EnrichmentType_GeoIp) isEnrichmentType_Type() {}

func (*EnrichmentType_SuspiciousIp) isEnrichmentType_Type() {}

func (*EnrichmentType_Aws) isEnrichmentType_Type() {}

func (*EnrichmentType_CustomEnrichment) isEnrichmentType_Type() {}

type GeoIpType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GeoIpType) Reset() {
	*x = GeoIpType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_enrichment_v1_enrichment_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoIpType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoIpType) ProtoMessage() {}

func (x *GeoIpType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoIpType.ProtoReflect.Descriptor instead.
func (*GeoIpType) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_type_proto_rawDescGZIP(), []int{1}
}

type SuspiciousIpType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SuspiciousIpType) Reset() {
	*x = SuspiciousIpType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_enrichment_v1_enrichment_type_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuspiciousIpType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuspiciousIpType) ProtoMessage() {}

func (x *SuspiciousIpType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_type_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuspiciousIpType.ProtoReflect.Descriptor instead.
func (*SuspiciousIpType) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_type_proto_rawDescGZIP(), []int{2}
}

type AwsType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceType *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
}

func (x *AwsType) Reset() {
	*x = AwsType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_enrichment_v1_enrichment_type_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AwsType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AwsType) ProtoMessage() {}

func (x *AwsType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_type_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AwsType.ProtoReflect.Descriptor instead.
func (*AwsType) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_type_proto_rawDescGZIP(), []int{3}
}

func (x *AwsType) GetResourceType() *wrapperspb.StringValue {
	if x != nil {
		return x.ResourceType
	}
	return nil
}

type CustomEnrichmentType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CustomEnrichmentType) Reset() {
	*x = CustomEnrichmentType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_enrichment_v1_enrichment_type_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomEnrichmentType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomEnrichmentType) ProtoMessage() {}

func (x *CustomEnrichmentType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_type_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomEnrichmentType.ProtoReflect.Descriptor instead.
func (*CustomEnrichmentType) Descriptor() ([]byte, []int) {
	return file_com_coralogix_enrichment_v1_enrichment_type_proto_rawDescGZIP(), []int{4}
}

func (x *CustomEnrichmentType) GetId() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Id
	}
	return nil
}

var File_com_coralogix_enrichment_v1_enrichment_type_proto protoreflect.FileDescriptor

var file_com_coralogix_enrichment_v1_enrichment_type_proto_rawDesc = []byte{
	0x0a, 0x31, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e,
	0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xcb, 0x02, 0x0a, 0x0e, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x67, 0x65, 0x6f, 0x5f, 0x69, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x6f, 0x49, 0x70, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x05, 0x67,
	0x65, 0x6f, 0x49, 0x70, 0x12, 0x54, 0x0a, 0x0d, 0x73, 0x75, 0x73, 0x70, 0x69, 0x63, 0x69, 0x6f,
	0x75, 0x73, 0x5f, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x6e, 0x72, 0x69,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x73, 0x70, 0x69, 0x63,
	0x69, 0x6f, 0x75, 0x73, 0x49, 0x70, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x75,
	0x73, 0x70, 0x69, 0x63, 0x69, 0x6f, 0x75, 0x73, 0x49, 0x70, 0x12, 0x38, 0x0a, 0x03, 0x61, 0x77,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x77, 0x73, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52,
	0x03, 0x61, 0x77, 0x73, 0x12, 0x60, 0x0a, 0x11, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x65,
	0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x48, 0x00, 0x52, 0x10, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x72, 0x69,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x0b,
	0x0a, 0x09, 0x47, 0x65, 0x6f, 0x49, 0x70, 0x54, 0x79, 0x70, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x53,
	0x75, 0x73, 0x70, 0x69, 0x63, 0x69, 0x6f, 0x75, 0x73, 0x49, 0x70, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x4c, 0x0a, 0x07, 0x41, 0x77, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x41, 0x0a, 0x0d, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x44, 0x0a,
	0x14, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x02, 0x69, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_enrichment_v1_enrichment_type_proto_rawDescOnce sync.Once
	file_com_coralogix_enrichment_v1_enrichment_type_proto_rawDescData = file_com_coralogix_enrichment_v1_enrichment_type_proto_rawDesc
)

func file_com_coralogix_enrichment_v1_enrichment_type_proto_rawDescGZIP() []byte {
	file_com_coralogix_enrichment_v1_enrichment_type_proto_rawDescOnce.Do(func() {
		file_com_coralogix_enrichment_v1_enrichment_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_enrichment_v1_enrichment_type_proto_rawDescData)
	})
	return file_com_coralogix_enrichment_v1_enrichment_type_proto_rawDescData
}

var file_com_coralogix_enrichment_v1_enrichment_type_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_com_coralogix_enrichment_v1_enrichment_type_proto_goTypes = []interface{}{
	(*EnrichmentType)(nil),         // 0: com.coralogix.enrichment.v1.EnrichmentType
	(*GeoIpType)(nil),              // 1: com.coralogix.enrichment.v1.GeoIpType
	(*SuspiciousIpType)(nil),       // 2: com.coralogix.enrichment.v1.SuspiciousIpType
	(*AwsType)(nil),                // 3: com.coralogix.enrichment.v1.AwsType
	(*CustomEnrichmentType)(nil),   // 4: com.coralogix.enrichment.v1.CustomEnrichmentType
	(*wrapperspb.StringValue)(nil), // 5: google.protobuf.StringValue
	(*wrapperspb.UInt32Value)(nil), // 6: google.protobuf.UInt32Value
}
var file_com_coralogix_enrichment_v1_enrichment_type_proto_depIdxs = []int32{
	1, // 0: com.coralogix.enrichment.v1.EnrichmentType.geo_ip:type_name -> com.coralogix.enrichment.v1.GeoIpType
	2, // 1: com.coralogix.enrichment.v1.EnrichmentType.suspicious_ip:type_name -> com.coralogix.enrichment.v1.SuspiciousIpType
	3, // 2: com.coralogix.enrichment.v1.EnrichmentType.aws:type_name -> com.coralogix.enrichment.v1.AwsType
	4, // 3: com.coralogix.enrichment.v1.EnrichmentType.custom_enrichment:type_name -> com.coralogix.enrichment.v1.CustomEnrichmentType
	5, // 4: com.coralogix.enrichment.v1.AwsType.resource_type:type_name -> google.protobuf.StringValue
	6, // 5: com.coralogix.enrichment.v1.CustomEnrichmentType.id:type_name -> google.protobuf.UInt32Value
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_com_coralogix_enrichment_v1_enrichment_type_proto_init() }
func file_com_coralogix_enrichment_v1_enrichment_type_proto_init() {
	if File_com_coralogix_enrichment_v1_enrichment_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_enrichment_v1_enrichment_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnrichmentType); i {
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
		file_com_coralogix_enrichment_v1_enrichment_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoIpType); i {
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
		file_com_coralogix_enrichment_v1_enrichment_type_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuspiciousIpType); i {
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
		file_com_coralogix_enrichment_v1_enrichment_type_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AwsType); i {
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
		file_com_coralogix_enrichment_v1_enrichment_type_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomEnrichmentType); i {
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
	file_com_coralogix_enrichment_v1_enrichment_type_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*EnrichmentType_GeoIp)(nil),
		(*EnrichmentType_SuspiciousIp)(nil),
		(*EnrichmentType_Aws)(nil),
		(*EnrichmentType_CustomEnrichment)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_enrichment_v1_enrichment_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_enrichment_v1_enrichment_type_proto_goTypes,
		DependencyIndexes: file_com_coralogix_enrichment_v1_enrichment_type_proto_depIdxs,
		MessageInfos:      file_com_coralogix_enrichment_v1_enrichment_type_proto_msgTypes,
	}.Build()
	File_com_coralogix_enrichment_v1_enrichment_type_proto = out.File
	file_com_coralogix_enrichment_v1_enrichment_type_proto_rawDesc = nil
	file_com_coralogix_enrichment_v1_enrichment_type_proto_goTypes = nil
	file_com_coralogix_enrichment_v1_enrichment_type_proto_depIdxs = nil
}
