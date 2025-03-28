// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: com/coralogixapis/incidents/v1/incident_timeline_event.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type IncidentTimelineEvent struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Name          *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	StartTime     *timestamppb.Timestamp  `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime       *timestamppb.Timestamp  `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	State         IncidentState           `protobuf:"varint,4,opt,name=state,proto3,enum=com.coralogixapis.incidents.v1.IncidentState" json:"state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IncidentTimelineEvent) Reset() {
	*x = IncidentTimelineEvent{}
	mi := &file_com_coralogixapis_incidents_v1_incident_timeline_event_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IncidentTimelineEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncidentTimelineEvent) ProtoMessage() {}

func (x *IncidentTimelineEvent) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_incidents_v1_incident_timeline_event_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncidentTimelineEvent.ProtoReflect.Descriptor instead.
func (*IncidentTimelineEvent) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_incidents_v1_incident_timeline_event_proto_rawDescGZIP(), []int{0}
}

func (x *IncidentTimelineEvent) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *IncidentTimelineEvent) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *IncidentTimelineEvent) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *IncidentTimelineEvent) GetState() IncidentState {
	if x != nil {
		return x.State
	}
	return IncidentState_INCIDENT_STATE_UNSPECIFIED
}

var File_com_coralogixapis_incidents_v1_incident_timeline_event_proto protoreflect.FileDescriptor

var file_com_coralogixapis_incidents_v1_incident_timeline_event_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x33,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69,
	0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x03, 0x0a, 0x15, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x55,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x23, 0x92, 0x41, 0x20, 0x4a,
	0x1e, 0x22, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x56, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x1b, 0x92, 0x41, 0x18, 0x4a, 0x16, 0x22, 0x32, 0x30, 0x32,
	0x34, 0x2d, 0x30, 0x31, 0x2d, 0x30, 0x31, 0x54, 0x30, 0x30, 0x3a, 0x30, 0x30, 0x3a, 0x30, 0x30,
	0x5a, 0x22, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x52, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x1b, 0x92, 0x41, 0x18,
	0x4a, 0x16, 0x22, 0x32, 0x30, 0x32, 0x34, 0x2d, 0x30, 0x31, 0x2d, 0x30, 0x31, 0x54, 0x30, 0x30,
	0x3a, 0x30, 0x30, 0x3a, 0x30, 0x30, 0x5a, 0x22, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x43, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x3a, 0x5f, 0x92, 0x41, 0x5c, 0x0a, 0x5a, 0x2a, 0x15, 0x49,
	0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x32, 0x1a, 0x41, 0x6e, 0x20, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x20, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0xd2, 0x01, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xd2, 0x01, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0xd2, 0x01, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0xd2,
	0x01, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_incidents_v1_incident_timeline_event_proto_rawDescOnce sync.Once
	file_com_coralogixapis_incidents_v1_incident_timeline_event_proto_rawDescData = file_com_coralogixapis_incidents_v1_incident_timeline_event_proto_rawDesc
)

func file_com_coralogixapis_incidents_v1_incident_timeline_event_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_incidents_v1_incident_timeline_event_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_incidents_v1_incident_timeline_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_incidents_v1_incident_timeline_event_proto_rawDescData)
	})
	return file_com_coralogixapis_incidents_v1_incident_timeline_event_proto_rawDescData
}

var file_com_coralogixapis_incidents_v1_incident_timeline_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_incidents_v1_incident_timeline_event_proto_goTypes = []any{
	(*IncidentTimelineEvent)(nil),  // 0: com.coralogixapis.incidents.v1.IncidentTimelineEvent
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
	(*timestamppb.Timestamp)(nil),  // 2: google.protobuf.Timestamp
	(IncidentState)(0),             // 3: com.coralogixapis.incidents.v1.IncidentState
}
var file_com_coralogixapis_incidents_v1_incident_timeline_event_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.incidents.v1.IncidentTimelineEvent.name:type_name -> google.protobuf.StringValue
	2, // 1: com.coralogixapis.incidents.v1.IncidentTimelineEvent.start_time:type_name -> google.protobuf.Timestamp
	2, // 2: com.coralogixapis.incidents.v1.IncidentTimelineEvent.end_time:type_name -> google.protobuf.Timestamp
	3, // 3: com.coralogixapis.incidents.v1.IncidentTimelineEvent.state:type_name -> com.coralogixapis.incidents.v1.IncidentState
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_incidents_v1_incident_timeline_event_proto_init() }
func file_com_coralogixapis_incidents_v1_incident_timeline_event_proto_init() {
	if File_com_coralogixapis_incidents_v1_incident_timeline_event_proto != nil {
		return
	}
	file_com_coralogixapis_incidents_v1_incident_state_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_incidents_v1_incident_timeline_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_incidents_v1_incident_timeline_event_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_incidents_v1_incident_timeline_event_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_incidents_v1_incident_timeline_event_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_incidents_v1_incident_timeline_event_proto = out.File
	file_com_coralogixapis_incidents_v1_incident_timeline_event_proto_rawDesc = nil
	file_com_coralogixapis_incidents_v1_incident_timeline_event_proto_goTypes = nil
	file_com_coralogixapis_incidents_v1_incident_timeline_event_proto_depIdxs = nil
}
