// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogix/enrichment/v1/enrichment_type.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

type EnrichmentType struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Type:
	//
	//	*EnrichmentType_GeoIp
	//	*EnrichmentType_SuspiciousIp
	//	*EnrichmentType_Aws
	//	*EnrichmentType_CustomEnrichment
	Type          isEnrichmentType_Type `protobuf_oneof:"type"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EnrichmentType) Reset() {
	*x = EnrichmentType{}
	mi := &file_com_coralogix_enrichment_v1_enrichment_type_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EnrichmentType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrichmentType) ProtoMessage() {}

func (x *EnrichmentType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_type_proto_msgTypes[0]
	if x != nil {
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

func (x *EnrichmentType) GetType() isEnrichmentType_Type {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *EnrichmentType) GetGeoIp() *GeoIpType {
	if x != nil {
		if x, ok := x.Type.(*EnrichmentType_GeoIp); ok {
			return x.GeoIp
		}
	}
	return nil
}

func (x *EnrichmentType) GetSuspiciousIp() *SuspiciousIpType {
	if x != nil {
		if x, ok := x.Type.(*EnrichmentType_SuspiciousIp); ok {
			return x.SuspiciousIp
		}
	}
	return nil
}

func (x *EnrichmentType) GetAws() *AwsType {
	if x != nil {
		if x, ok := x.Type.(*EnrichmentType_Aws); ok {
			return x.Aws
		}
	}
	return nil
}

func (x *EnrichmentType) GetCustomEnrichment() *CustomEnrichmentType {
	if x != nil {
		if x, ok := x.Type.(*EnrichmentType_CustomEnrichment); ok {
			return x.CustomEnrichment
		}
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	WithAsn       *bool                  `protobuf:"varint,4,opt,name=with_asn,json=withAsn,proto3,oneof" json:"with_asn,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GeoIpType) Reset() {
	*x = GeoIpType{}
	mi := &file_com_coralogix_enrichment_v1_enrichment_type_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GeoIpType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoIpType) ProtoMessage() {}

func (x *GeoIpType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_type_proto_msgTypes[1]
	if x != nil {
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

func (x *GeoIpType) GetWithAsn() bool {
	if x != nil && x.WithAsn != nil {
		return *x.WithAsn
	}
	return false
}

type SuspiciousIpType struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SuspiciousIpType) Reset() {
	*x = SuspiciousIpType{}
	mi := &file_com_coralogix_enrichment_v1_enrichment_type_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SuspiciousIpType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuspiciousIpType) ProtoMessage() {}

func (x *SuspiciousIpType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_type_proto_msgTypes[2]
	if x != nil {
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
	state         protoimpl.MessageState  `protogen:"open.v1"`
	ResourceType  *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AwsType) Reset() {
	*x = AwsType{}
	mi := &file_com_coralogix_enrichment_v1_enrichment_type_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AwsType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AwsType) ProtoMessage() {}

func (x *AwsType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_type_proto_msgTypes[3]
	if x != nil {
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
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Id            *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CustomEnrichmentType) Reset() {
	*x = CustomEnrichmentType{}
	mi := &file_com_coralogix_enrichment_v1_enrichment_type_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomEnrichmentType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomEnrichmentType) ProtoMessage() {}

func (x *CustomEnrichmentType) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_enrichment_v1_enrichment_type_proto_msgTypes[4]
	if x != nil {
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

const file_com_coralogix_enrichment_v1_enrichment_type_proto_rawDesc = "" +
	"\n" +
	"1com/coralogix/enrichment/v1/enrichment_type.proto\x12\x1bcom.coralogix.enrichment.v1\x1a\x1egoogle/protobuf/wrappers.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\xcb\x02\n" +
	"\x0eEnrichmentType\x12?\n" +
	"\x06geo_ip\x18\x01 \x01(\v2&.com.coralogix.enrichment.v1.GeoIpTypeH\x00R\x05geoIp\x12T\n" +
	"\rsuspicious_ip\x18\x02 \x01(\v2-.com.coralogix.enrichment.v1.SuspiciousIpTypeH\x00R\fsuspiciousIp\x128\n" +
	"\x03aws\x18\x03 \x01(\v2$.com.coralogix.enrichment.v1.AwsTypeH\x00R\x03aws\x12`\n" +
	"\x11custom_enrichment\x18\x04 \x01(\v21.com.coralogix.enrichment.v1.CustomEnrichmentTypeH\x00R\x10customEnrichmentB\x06\n" +
	"\x04type\"8\n" +
	"\tGeoIpType\x12\x1e\n" +
	"\bwith_asn\x18\x04 \x01(\bH\x00R\awithAsn\x88\x01\x01B\v\n" +
	"\t_with_asn\"\x12\n" +
	"\x10SuspiciousIpType\"X\n" +
	"\aAwsType\x12M\n" +
	"\rresource_type\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueB\n" +
	"\x92A\aJ\x05\"ec2\"R\fresourceType\"L\n" +
	"\x14CustomEnrichmentType\x124\n" +
	"\x02id\x18\x01 \x01(\v2\x1c.google.protobuf.UInt32ValueB\x06\x92A\x03J\x011R\x02idb\x06proto3"

var (
	file_com_coralogix_enrichment_v1_enrichment_type_proto_rawDescOnce sync.Once
	file_com_coralogix_enrichment_v1_enrichment_type_proto_rawDescData []byte
)

func file_com_coralogix_enrichment_v1_enrichment_type_proto_rawDescGZIP() []byte {
	file_com_coralogix_enrichment_v1_enrichment_type_proto_rawDescOnce.Do(func() {
		file_com_coralogix_enrichment_v1_enrichment_type_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogix_enrichment_v1_enrichment_type_proto_rawDesc), len(file_com_coralogix_enrichment_v1_enrichment_type_proto_rawDesc)))
	})
	return file_com_coralogix_enrichment_v1_enrichment_type_proto_rawDescData
}

var file_com_coralogix_enrichment_v1_enrichment_type_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_com_coralogix_enrichment_v1_enrichment_type_proto_goTypes = []any{
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
	file_com_coralogix_enrichment_v1_enrichment_type_proto_msgTypes[0].OneofWrappers = []any{
		(*EnrichmentType_GeoIp)(nil),
		(*EnrichmentType_SuspiciousIp)(nil),
		(*EnrichmentType_Aws)(nil),
		(*EnrichmentType_CustomEnrichment)(nil),
	}
	file_com_coralogix_enrichment_v1_enrichment_type_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogix_enrichment_v1_enrichment_type_proto_rawDesc), len(file_com_coralogix_enrichment_v1_enrichment_type_proto_rawDesc)),
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
	file_com_coralogix_enrichment_v1_enrichment_type_proto_goTypes = nil
	file_com_coralogix_enrichment_v1_enrichment_type_proto_depIdxs = nil
}
