// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: com/coralogix/catalog/v1/service_action.proto

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

type ServiceAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActionName                *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=action_name,json=actionName,proto3" json:"action_name,omitempty"`
	ServiceName               *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	Method                    *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	TimeConsumingPercentage   *wrapperspb.Int64Value  `protobuf:"bytes,4,opt,name=time_consuming_percentage,json=timeConsumingPercentage,proto3" json:"time_consuming_percentage,omitempty"`
	TimeConsuming             *wrapperspb.Int64Value  `protobuf:"bytes,5,opt,name=time_consuming,json=timeConsuming,proto3" json:"time_consuming,omitempty"`
	ErrorsPercentage          *wrapperspb.Int64Value  `protobuf:"bytes,6,opt,name=errors_percentage,json=errorsPercentage,proto3" json:"errors_percentage,omitempty"`
	Errors                    *wrapperspb.Int64Value  `protobuf:"bytes,7,opt,name=errors,proto3" json:"errors,omitempty"`
	NumberOfRequest           *wrapperspb.Int64Value  `protobuf:"bytes,8,opt,name=number_of_request,json=numberOfRequest,proto3" json:"number_of_request,omitempty"`
	NumberOfRequestPercentage *wrapperspb.Int64Value  `protobuf:"bytes,9,opt,name=number_of_request_percentage,json=numberOfRequestPercentage,proto3" json:"number_of_request_percentage,omitempty"`
	P95Latency                *wrapperspb.Int64Value  `protobuf:"bytes,10,opt,name=p95_latency,json=p95Latency,proto3" json:"p95_latency,omitempty"`
	P99Latency                *wrapperspb.Int64Value  `protobuf:"bytes,11,opt,name=p99_latency,json=p99Latency,proto3" json:"p99_latency,omitempty"`
}

func (x *ServiceAction) Reset() {
	*x = ServiceAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_catalog_v1_service_action_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceAction) ProtoMessage() {}

func (x *ServiceAction) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_catalog_v1_service_action_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceAction.ProtoReflect.Descriptor instead.
func (*ServiceAction) Descriptor() ([]byte, []int) {
	return file_com_coralogix_catalog_v1_service_action_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceAction) GetActionName() *wrapperspb.StringValue {
	if x != nil {
		return x.ActionName
	}
	return nil
}

func (x *ServiceAction) GetServiceName() *wrapperspb.StringValue {
	if x != nil {
		return x.ServiceName
	}
	return nil
}

func (x *ServiceAction) GetMethod() *wrapperspb.StringValue {
	if x != nil {
		return x.Method
	}
	return nil
}

func (x *ServiceAction) GetTimeConsumingPercentage() *wrapperspb.Int64Value {
	if x != nil {
		return x.TimeConsumingPercentage
	}
	return nil
}

func (x *ServiceAction) GetTimeConsuming() *wrapperspb.Int64Value {
	if x != nil {
		return x.TimeConsuming
	}
	return nil
}

func (x *ServiceAction) GetErrorsPercentage() *wrapperspb.Int64Value {
	if x != nil {
		return x.ErrorsPercentage
	}
	return nil
}

func (x *ServiceAction) GetErrors() *wrapperspb.Int64Value {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *ServiceAction) GetNumberOfRequest() *wrapperspb.Int64Value {
	if x != nil {
		return x.NumberOfRequest
	}
	return nil
}

func (x *ServiceAction) GetNumberOfRequestPercentage() *wrapperspb.Int64Value {
	if x != nil {
		return x.NumberOfRequestPercentage
	}
	return nil
}

func (x *ServiceAction) GetP95Latency() *wrapperspb.Int64Value {
	if x != nil {
		return x.P95Latency
	}
	return nil
}

func (x *ServiceAction) GetP99Latency() *wrapperspb.Int64Value {
	if x != nil {
		return x.P99Latency
	}
	return nil
}

var File_com_coralogix_catalog_v1_service_action_proto protoreflect.FileDescriptor

var file_com_coralogix_catalog_v1_service_action_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x06, 0x0a, 0x0d, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x0b, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x57, 0x0a, 0x19, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x69, 0x6e, 0x67, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x17, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x69, 0x6e, 0x67,
	0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x42, 0x0a, 0x0e, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x0d, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x48,
	0x0a, 0x11, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36,
	0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x50, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x47, 0x0a,
	0x11, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5c, 0x0a, 0x1c, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x5f, 0x6f, 0x66, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49,
	0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x19, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x4f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x61, 0x67, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x70, 0x39, 0x35, 0x5f, 0x6c, 0x61, 0x74, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36,
	0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x70, 0x39, 0x35, 0x4c, 0x61, 0x74, 0x65, 0x6e,
	0x63, 0x79, 0x12, 0x3c, 0x0a, 0x0b, 0x70, 0x39, 0x39, 0x5f, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63,
	0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x70, 0x39, 0x39, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_catalog_v1_service_action_proto_rawDescOnce sync.Once
	file_com_coralogix_catalog_v1_service_action_proto_rawDescData = file_com_coralogix_catalog_v1_service_action_proto_rawDesc
)

func file_com_coralogix_catalog_v1_service_action_proto_rawDescGZIP() []byte {
	file_com_coralogix_catalog_v1_service_action_proto_rawDescOnce.Do(func() {
		file_com_coralogix_catalog_v1_service_action_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_catalog_v1_service_action_proto_rawDescData)
	})
	return file_com_coralogix_catalog_v1_service_action_proto_rawDescData
}

var file_com_coralogix_catalog_v1_service_action_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogix_catalog_v1_service_action_proto_goTypes = []any{
	(*ServiceAction)(nil),          // 0: com.coralogix.catalog.v1.ServiceAction
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
	(*wrapperspb.Int64Value)(nil),  // 2: google.protobuf.Int64Value
}
var file_com_coralogix_catalog_v1_service_action_proto_depIdxs = []int32{
	1,  // 0: com.coralogix.catalog.v1.ServiceAction.action_name:type_name -> google.protobuf.StringValue
	1,  // 1: com.coralogix.catalog.v1.ServiceAction.service_name:type_name -> google.protobuf.StringValue
	1,  // 2: com.coralogix.catalog.v1.ServiceAction.method:type_name -> google.protobuf.StringValue
	2,  // 3: com.coralogix.catalog.v1.ServiceAction.time_consuming_percentage:type_name -> google.protobuf.Int64Value
	2,  // 4: com.coralogix.catalog.v1.ServiceAction.time_consuming:type_name -> google.protobuf.Int64Value
	2,  // 5: com.coralogix.catalog.v1.ServiceAction.errors_percentage:type_name -> google.protobuf.Int64Value
	2,  // 6: com.coralogix.catalog.v1.ServiceAction.errors:type_name -> google.protobuf.Int64Value
	2,  // 7: com.coralogix.catalog.v1.ServiceAction.number_of_request:type_name -> google.protobuf.Int64Value
	2,  // 8: com.coralogix.catalog.v1.ServiceAction.number_of_request_percentage:type_name -> google.protobuf.Int64Value
	2,  // 9: com.coralogix.catalog.v1.ServiceAction.p95_latency:type_name -> google.protobuf.Int64Value
	2,  // 10: com.coralogix.catalog.v1.ServiceAction.p99_latency:type_name -> google.protobuf.Int64Value
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_com_coralogix_catalog_v1_service_action_proto_init() }
func file_com_coralogix_catalog_v1_service_action_proto_init() {
	if File_com_coralogix_catalog_v1_service_action_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_catalog_v1_service_action_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ServiceAction); i {
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
			RawDescriptor: file_com_coralogix_catalog_v1_service_action_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_catalog_v1_service_action_proto_goTypes,
		DependencyIndexes: file_com_coralogix_catalog_v1_service_action_proto_depIdxs,
		MessageInfos:      file_com_coralogix_catalog_v1_service_action_proto_msgTypes,
	}.Build()
	File_com_coralogix_catalog_v1_service_action_proto = out.File
	file_com_coralogix_catalog_v1_service_action_proto_rawDesc = nil
	file_com_coralogix_catalog_v1_service_action_proto_goTypes = nil
	file_com_coralogix_catalog_v1_service_action_proto_depIdxs = nil
}
