// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: com/coralogixapis/apm/queries/v1/event.proto

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

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectKind   *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=object_kind,json=objectKind,proto3" json:"object_kind,omitempty"`
	ObjectName   *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=object_name,json=objectName,proto3" json:"object_name,omitempty"`
	Namespace    *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ClusterName  *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	Message      *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	Severity     *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=severity,proto3" json:"severity,omitempty"`
	Reason       *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=reason,proto3" json:"reason,omitempty"`
	CreationTime *wrapperspb.StringValue `protobuf:"bytes,8,opt,name=creation_time,json=creationTime,proto3" json:"creation_time,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	mi := &file_com_coralogixapis_apm_queries_v1_event_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_queries_v1_event_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_queries_v1_event_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetObjectKind() *wrapperspb.StringValue {
	if x != nil {
		return x.ObjectKind
	}
	return nil
}

func (x *Event) GetObjectName() *wrapperspb.StringValue {
	if x != nil {
		return x.ObjectName
	}
	return nil
}

func (x *Event) GetNamespace() *wrapperspb.StringValue {
	if x != nil {
		return x.Namespace
	}
	return nil
}

func (x *Event) GetClusterName() *wrapperspb.StringValue {
	if x != nil {
		return x.ClusterName
	}
	return nil
}

func (x *Event) GetMessage() *wrapperspb.StringValue {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *Event) GetSeverity() *wrapperspb.StringValue {
	if x != nil {
		return x.Severity
	}
	return nil
}

func (x *Event) GetReason() *wrapperspb.StringValue {
	if x != nil {
		return x.Reason
	}
	return nil
}

func (x *Event) GetCreationTime() *wrapperspb.StringValue {
	if x != nil {
		return x.CreationTime
	}
	return nil
}

var File_com_coralogixapis_apm_queries_v1_event_proto protoreflect.FileDescriptor

var file_com_coralogixapis_apm_queries_v1_event_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xed, 0x03, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x3d, 0x0a, 0x0b, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x3d, 0x0a, 0x0b, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a,
	0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x73,
	0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x34, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x41, 0x0a,
	0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_apm_queries_v1_event_proto_rawDescOnce sync.Once
	file_com_coralogixapis_apm_queries_v1_event_proto_rawDescData = file_com_coralogixapis_apm_queries_v1_event_proto_rawDesc
)

func file_com_coralogixapis_apm_queries_v1_event_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_apm_queries_v1_event_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_apm_queries_v1_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_apm_queries_v1_event_proto_rawDescData)
	})
	return file_com_coralogixapis_apm_queries_v1_event_proto_rawDescData
}

var file_com_coralogixapis_apm_queries_v1_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_apm_queries_v1_event_proto_goTypes = []any{
	(*Event)(nil),                  // 0: com.coralogixapis.apm.queries.v1.Event
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
}
var file_com_coralogixapis_apm_queries_v1_event_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.apm.queries.v1.Event.object_kind:type_name -> google.protobuf.StringValue
	1, // 1: com.coralogixapis.apm.queries.v1.Event.object_name:type_name -> google.protobuf.StringValue
	1, // 2: com.coralogixapis.apm.queries.v1.Event.namespace:type_name -> google.protobuf.StringValue
	1, // 3: com.coralogixapis.apm.queries.v1.Event.cluster_name:type_name -> google.protobuf.StringValue
	1, // 4: com.coralogixapis.apm.queries.v1.Event.message:type_name -> google.protobuf.StringValue
	1, // 5: com.coralogixapis.apm.queries.v1.Event.severity:type_name -> google.protobuf.StringValue
	1, // 6: com.coralogixapis.apm.queries.v1.Event.reason:type_name -> google.protobuf.StringValue
	1, // 7: com.coralogixapis.apm.queries.v1.Event.creation_time:type_name -> google.protobuf.StringValue
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_apm_queries_v1_event_proto_init() }
func file_com_coralogixapis_apm_queries_v1_event_proto_init() {
	if File_com_coralogixapis_apm_queries_v1_event_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_apm_queries_v1_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_apm_queries_v1_event_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_apm_queries_v1_event_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_apm_queries_v1_event_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_apm_queries_v1_event_proto = out.File
	file_com_coralogixapis_apm_queries_v1_event_proto_rawDesc = nil
	file_com_coralogixapis_apm_queries_v1_event_proto_goTypes = nil
	file_com_coralogixapis_apm_queries_v1_event_proto_depIdxs = nil
}
