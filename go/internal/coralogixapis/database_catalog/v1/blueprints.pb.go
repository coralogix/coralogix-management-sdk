// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: com/coralogixapis/database_catalog/v1/blueprints.proto

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

type BlueprintData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are assignable to BlueprintData:
	//
	//	*BlueprintData_Blueprint
	//	*BlueprintData_TotalClientServices
	//	*BlueprintData_P95Latency
	//	*BlueprintData_TotalQueries
	//	*BlueprintData_TotalFailures
	//	*BlueprintData_FailuresPercentage
	//	*BlueprintData_TimeConsuming
	//	*BlueprintData_TimeConsumingPercentage
	BlueprintData isBlueprintData_BlueprintData `protobuf_oneof:"blueprint_data"`
}

func (x *BlueprintData) Reset() {
	*x = BlueprintData{}
	mi := &file_com_coralogixapis_database_catalog_v1_blueprints_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BlueprintData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlueprintData) ProtoMessage() {}

func (x *BlueprintData) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_database_catalog_v1_blueprints_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlueprintData.ProtoReflect.Descriptor instead.
func (*BlueprintData) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_database_catalog_v1_blueprints_proto_rawDescGZIP(), []int{0}
}

func (x *BlueprintData) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

func (m *BlueprintData) GetBlueprintData() isBlueprintData_BlueprintData {
	if m != nil {
		return m.BlueprintData
	}
	return nil
}

func (x *BlueprintData) GetBlueprint() *wrapperspb.StringValue {
	if x, ok := x.GetBlueprintData().(*BlueprintData_Blueprint); ok {
		return x.Blueprint
	}
	return nil
}

func (x *BlueprintData) GetTotalClientServices() *wrapperspb.Int64Value {
	if x, ok := x.GetBlueprintData().(*BlueprintData_TotalClientServices); ok {
		return x.TotalClientServices
	}
	return nil
}

func (x *BlueprintData) GetP95Latency() *wrapperspb.Int64Value {
	if x, ok := x.GetBlueprintData().(*BlueprintData_P95Latency); ok {
		return x.P95Latency
	}
	return nil
}

func (x *BlueprintData) GetTotalQueries() *wrapperspb.Int64Value {
	if x, ok := x.GetBlueprintData().(*BlueprintData_TotalQueries); ok {
		return x.TotalQueries
	}
	return nil
}

func (x *BlueprintData) GetTotalFailures() *wrapperspb.Int64Value {
	if x, ok := x.GetBlueprintData().(*BlueprintData_TotalFailures); ok {
		return x.TotalFailures
	}
	return nil
}

func (x *BlueprintData) GetFailuresPercentage() *wrapperspb.FloatValue {
	if x, ok := x.GetBlueprintData().(*BlueprintData_FailuresPercentage); ok {
		return x.FailuresPercentage
	}
	return nil
}

func (x *BlueprintData) GetTimeConsuming() *wrapperspb.Int64Value {
	if x, ok := x.GetBlueprintData().(*BlueprintData_TimeConsuming); ok {
		return x.TimeConsuming
	}
	return nil
}

func (x *BlueprintData) GetTimeConsumingPercentage() *wrapperspb.Int64Value {
	if x, ok := x.GetBlueprintData().(*BlueprintData_TimeConsumingPercentage); ok {
		return x.TimeConsumingPercentage
	}
	return nil
}

type isBlueprintData_BlueprintData interface {
	isBlueprintData_BlueprintData()
}

type BlueprintData_Blueprint struct {
	Blueprint *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=blueprint,proto3,oneof"`
}

type BlueprintData_TotalClientServices struct {
	TotalClientServices *wrapperspb.Int64Value `protobuf:"bytes,3,opt,name=total_client_services,json=totalClientServices,proto3,oneof"`
}

type BlueprintData_P95Latency struct {
	P95Latency *wrapperspb.Int64Value `protobuf:"bytes,4,opt,name=p95_latency,json=p95Latency,proto3,oneof"`
}

type BlueprintData_TotalQueries struct {
	TotalQueries *wrapperspb.Int64Value `protobuf:"bytes,5,opt,name=total_queries,json=totalQueries,proto3,oneof"`
}

type BlueprintData_TotalFailures struct {
	TotalFailures *wrapperspb.Int64Value `protobuf:"bytes,6,opt,name=total_failures,json=totalFailures,proto3,oneof"`
}

type BlueprintData_FailuresPercentage struct {
	FailuresPercentage *wrapperspb.FloatValue `protobuf:"bytes,7,opt,name=failures_percentage,json=failuresPercentage,proto3,oneof"`
}

type BlueprintData_TimeConsuming struct {
	TimeConsuming *wrapperspb.Int64Value `protobuf:"bytes,8,opt,name=time_consuming,json=timeConsuming,proto3,oneof"`
}

type BlueprintData_TimeConsumingPercentage struct {
	TimeConsumingPercentage *wrapperspb.Int64Value `protobuf:"bytes,9,opt,name=time_consuming_percentage,json=timeConsumingPercentage,proto3,oneof"`
}

func (*BlueprintData_Blueprint) isBlueprintData_BlueprintData() {}

func (*BlueprintData_TotalClientServices) isBlueprintData_BlueprintData() {}

func (*BlueprintData_P95Latency) isBlueprintData_BlueprintData() {}

func (*BlueprintData_TotalQueries) isBlueprintData_BlueprintData() {}

func (*BlueprintData_TotalFailures) isBlueprintData_BlueprintData() {}

func (*BlueprintData_FailuresPercentage) isBlueprintData_BlueprintData() {}

func (*BlueprintData_TimeConsuming) isBlueprintData_BlueprintData() {}

func (*BlueprintData_TimeConsumingPercentage) isBlueprintData_BlueprintData() {}

var File_com_coralogixapis_database_catalog_v1_blueprints_proto protoreflect.FileDescriptor

var file_com_coralogixapis_database_catalog_v1_blueprints_proto_rawDesc = []byte{
	0x0a, 0x36, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9b, 0x05, 0x0a, 0x0d, 0x42, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x3c, 0x0a, 0x09, 0x62, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x48, 0x00, 0x52, 0x09, 0x62, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x12, 0x51, 0x0a,
	0x15, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49,
	0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x13, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x12, 0x3e, 0x0a, 0x0b, 0x70, 0x39, 0x35, 0x5f, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x70, 0x39, 0x35, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x42, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x51, 0x75, 0x65,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x66, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49,
	0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x12, 0x4e, 0x0a, 0x13, 0x66, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x12, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73,
	0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x44, 0x0a, 0x0e, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48,
	0x00, 0x52, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x69, 0x6e, 0x67,
	0x12, 0x59, 0x0a, 0x19, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x69,
	0x6e, 0x67, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x48, 0x00, 0x52, 0x17, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x69, 0x6e,
	0x67, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x62,
	0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_database_catalog_v1_blueprints_proto_rawDescOnce sync.Once
	file_com_coralogixapis_database_catalog_v1_blueprints_proto_rawDescData = file_com_coralogixapis_database_catalog_v1_blueprints_proto_rawDesc
)

func file_com_coralogixapis_database_catalog_v1_blueprints_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_database_catalog_v1_blueprints_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_database_catalog_v1_blueprints_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_database_catalog_v1_blueprints_proto_rawDescData)
	})
	return file_com_coralogixapis_database_catalog_v1_blueprints_proto_rawDescData
}

var file_com_coralogixapis_database_catalog_v1_blueprints_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_database_catalog_v1_blueprints_proto_goTypes = []any{
	(*BlueprintData)(nil),          // 0: com.coralogixapis.database_catalog.v1.BlueprintData
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
	(*wrapperspb.Int64Value)(nil),  // 2: google.protobuf.Int64Value
	(*wrapperspb.FloatValue)(nil),  // 3: google.protobuf.FloatValue
}
var file_com_coralogixapis_database_catalog_v1_blueprints_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.database_catalog.v1.BlueprintData.id:type_name -> google.protobuf.StringValue
	1, // 1: com.coralogixapis.database_catalog.v1.BlueprintData.blueprint:type_name -> google.protobuf.StringValue
	2, // 2: com.coralogixapis.database_catalog.v1.BlueprintData.total_client_services:type_name -> google.protobuf.Int64Value
	2, // 3: com.coralogixapis.database_catalog.v1.BlueprintData.p95_latency:type_name -> google.protobuf.Int64Value
	2, // 4: com.coralogixapis.database_catalog.v1.BlueprintData.total_queries:type_name -> google.protobuf.Int64Value
	2, // 5: com.coralogixapis.database_catalog.v1.BlueprintData.total_failures:type_name -> google.protobuf.Int64Value
	3, // 6: com.coralogixapis.database_catalog.v1.BlueprintData.failures_percentage:type_name -> google.protobuf.FloatValue
	2, // 7: com.coralogixapis.database_catalog.v1.BlueprintData.time_consuming:type_name -> google.protobuf.Int64Value
	2, // 8: com.coralogixapis.database_catalog.v1.BlueprintData.time_consuming_percentage:type_name -> google.protobuf.Int64Value
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_database_catalog_v1_blueprints_proto_init() }
func file_com_coralogixapis_database_catalog_v1_blueprints_proto_init() {
	if File_com_coralogixapis_database_catalog_v1_blueprints_proto != nil {
		return
	}
	file_com_coralogixapis_database_catalog_v1_blueprints_proto_msgTypes[0].OneofWrappers = []any{
		(*BlueprintData_Blueprint)(nil),
		(*BlueprintData_TotalClientServices)(nil),
		(*BlueprintData_P95Latency)(nil),
		(*BlueprintData_TotalQueries)(nil),
		(*BlueprintData_TotalFailures)(nil),
		(*BlueprintData_FailuresPercentage)(nil),
		(*BlueprintData_TimeConsuming)(nil),
		(*BlueprintData_TimeConsumingPercentage)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_database_catalog_v1_blueprints_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_database_catalog_v1_blueprints_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_database_catalog_v1_blueprints_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_database_catalog_v1_blueprints_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_database_catalog_v1_blueprints_proto = out.File
	file_com_coralogixapis_database_catalog_v1_blueprints_proto_rawDesc = nil
	file_com_coralogixapis_database_catalog_v1_blueprints_proto_goTypes = nil
	file_com_coralogixapis_database_catalog_v1_blueprints_proto_depIdxs = nil
}
