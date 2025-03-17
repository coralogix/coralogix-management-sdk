// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.26.1
// source: com/coralogixapis/apm/services/v1/apm_service.proto

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

type SloStatusCount struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ok            *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Breach        *wrapperspb.Int64Value `protobuf:"bytes,2,opt,name=breach,proto3" json:"breach,omitempty"`
	NotAvailable  *wrapperspb.Int64Value `protobuf:"bytes,3,opt,name=not_available,json=notAvailable,proto3" json:"not_available,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SloStatusCount) Reset() {
	*x = SloStatusCount{}
	mi := &file_com_coralogixapis_apm_services_v1_apm_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SloStatusCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SloStatusCount) ProtoMessage() {}

func (x *SloStatusCount) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_services_v1_apm_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SloStatusCount.ProtoReflect.Descriptor instead.
func (*SloStatusCount) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_services_v1_apm_service_proto_rawDescGZIP(), []int{0}
}

func (x *SloStatusCount) GetOk() *wrapperspb.Int64Value {
	if x != nil {
		return x.Ok
	}
	return nil
}

func (x *SloStatusCount) GetBreach() *wrapperspb.Int64Value {
	if x != nil {
		return x.Breach
	}
	return nil
}

func (x *SloStatusCount) GetNotAvailable() *wrapperspb.Int64Value {
	if x != nil {
		return x.NotAvailable
	}
	return nil
}

type ApmService struct {
	state          protoimpl.MessageState    `protogen:"open.v1"`
	Id             *wrapperspb.StringValue   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           *wrapperspb.StringValue   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type           *wrapperspb.StringValue   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Workloads      []*wrapperspb.StringValue `protobuf:"bytes,4,rep,name=workloads,proto3" json:"workloads,omitempty"`
	SloStatusCount *SloStatusCount           `protobuf:"bytes,5,opt,name=slo_status_count,json=sloStatusCount,proto3" json:"slo_status_count,omitempty"`
	Technology     *wrapperspb.StringValue   `protobuf:"bytes,6,opt,name=technology,proto3" json:"technology,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ApmService) Reset() {
	*x = ApmService{}
	mi := &file_com_coralogixapis_apm_services_v1_apm_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApmService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApmService) ProtoMessage() {}

func (x *ApmService) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_services_v1_apm_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApmService.ProtoReflect.Descriptor instead.
func (*ApmService) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_services_v1_apm_service_proto_rawDescGZIP(), []int{1}
}

func (x *ApmService) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ApmService) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *ApmService) GetType() *wrapperspb.StringValue {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *ApmService) GetWorkloads() []*wrapperspb.StringValue {
	if x != nil {
		return x.Workloads
	}
	return nil
}

func (x *ApmService) GetSloStatusCount() *SloStatusCount {
	if x != nil {
		return x.SloStatusCount
	}
	return nil
}

func (x *ApmService) GetTechnology() *wrapperspb.StringValue {
	if x != nil {
		return x.Technology
	}
	return nil
}

var File_com_coralogixapis_apm_services_v1_apm_service_proto protoreflect.FileDescriptor

var file_com_coralogixapis_apm_services_v1_apm_service_proto_rawDesc = []byte{
	0x0a, 0x33, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x01, 0x0a, 0x0e, 0x53, 0x6c, 0x6f,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x02, 0x6f,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x33, 0x0a, 0x06, 0x62, 0x72, 0x65, 0x61,
	0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x62, 0x72, 0x65, 0x61, 0x63, 0x68, 0x12, 0x40, 0x0a,
	0x0d, 0x6e, 0x6f, 0x74, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0c, 0x6e, 0x6f, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x22,
	0xf5, 0x02, 0x0a, 0x0a, 0x41, 0x70, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x3a, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x12, 0x5b, 0x0a, 0x10,
	0x73, 0x6c, 0x6f, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6c, 0x6f, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0e, 0x73, 0x6c, 0x6f, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x0a, 0x74, 0x65, 0x63,
	0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x74, 0x65, 0x63,
	0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_apm_services_v1_apm_service_proto_rawDescOnce sync.Once
	file_com_coralogixapis_apm_services_v1_apm_service_proto_rawDescData = file_com_coralogixapis_apm_services_v1_apm_service_proto_rawDesc
)

func file_com_coralogixapis_apm_services_v1_apm_service_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_apm_services_v1_apm_service_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_apm_services_v1_apm_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_apm_services_v1_apm_service_proto_rawDescData)
	})
	return file_com_coralogixapis_apm_services_v1_apm_service_proto_rawDescData
}

var file_com_coralogixapis_apm_services_v1_apm_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogixapis_apm_services_v1_apm_service_proto_goTypes = []any{
	(*SloStatusCount)(nil),         // 0: com.coralogixapis.apm.services.v1.SloStatusCount
	(*ApmService)(nil),             // 1: com.coralogixapis.apm.services.v1.ApmService
	(*wrapperspb.Int64Value)(nil),  // 2: google.protobuf.Int64Value
	(*wrapperspb.StringValue)(nil), // 3: google.protobuf.StringValue
}
var file_com_coralogixapis_apm_services_v1_apm_service_proto_depIdxs = []int32{
	2, // 0: com.coralogixapis.apm.services.v1.SloStatusCount.ok:type_name -> google.protobuf.Int64Value
	2, // 1: com.coralogixapis.apm.services.v1.SloStatusCount.breach:type_name -> google.protobuf.Int64Value
	2, // 2: com.coralogixapis.apm.services.v1.SloStatusCount.not_available:type_name -> google.protobuf.Int64Value
	3, // 3: com.coralogixapis.apm.services.v1.ApmService.id:type_name -> google.protobuf.StringValue
	3, // 4: com.coralogixapis.apm.services.v1.ApmService.name:type_name -> google.protobuf.StringValue
	3, // 5: com.coralogixapis.apm.services.v1.ApmService.type:type_name -> google.protobuf.StringValue
	3, // 6: com.coralogixapis.apm.services.v1.ApmService.workloads:type_name -> google.protobuf.StringValue
	0, // 7: com.coralogixapis.apm.services.v1.ApmService.slo_status_count:type_name -> com.coralogixapis.apm.services.v1.SloStatusCount
	3, // 8: com.coralogixapis.apm.services.v1.ApmService.technology:type_name -> google.protobuf.StringValue
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_apm_services_v1_apm_service_proto_init() }
func file_com_coralogixapis_apm_services_v1_apm_service_proto_init() {
	if File_com_coralogixapis_apm_services_v1_apm_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_apm_services_v1_apm_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_apm_services_v1_apm_service_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_apm_services_v1_apm_service_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_apm_services_v1_apm_service_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_apm_services_v1_apm_service_proto = out.File
	file_com_coralogixapis_apm_services_v1_apm_service_proto_rawDesc = nil
	file_com_coralogixapis_apm_services_v1_apm_service_proto_goTypes = nil
	file_com_coralogixapis_apm_services_v1_apm_service_proto_depIdxs = nil
}
