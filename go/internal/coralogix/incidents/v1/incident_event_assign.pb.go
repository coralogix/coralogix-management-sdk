// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: com/coralogixapis/incidents/v1/incident_event/incident_event_assign.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IncidentEventAssign struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Assignment *Assignment `protobuf:"bytes,1,opt,name=assignment,proto3" json:"assignment,omitempty"`
}

func (x *IncidentEventAssign) Reset() {
	*x = IncidentEventAssign{}
	mi := &file_com_coralogixapis_incidents_v1_incident_event_incident_event_assign_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IncidentEventAssign) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncidentEventAssign) ProtoMessage() {}

func (x *IncidentEventAssign) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_incidents_v1_incident_event_incident_event_assign_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncidentEventAssign.ProtoReflect.Descriptor instead.
func (*IncidentEventAssign) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_incidents_v1_incident_event_incident_event_assign_proto_rawDescGZIP(), []int{0}
}

func (x *IncidentEventAssign) GetAssignment() *Assignment {
	if x != nil {
		return x.Assignment
	}
	return nil
}

var File_com_coralogixapis_incidents_v1_incident_event_incident_event_assign_proto protoreflect.FileDescriptor

var file_com_coralogixapis_incidents_v1_incident_event_incident_event_assign_proto_rawDesc = []byte{
	0x0a, 0x49, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f,
	0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x61,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69,
	0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x2d, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69,
	0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x13, 0x49, 0x6e,
	0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x12, 0x4a, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_incidents_v1_incident_event_incident_event_assign_proto_rawDescOnce sync.Once
	file_com_coralogixapis_incidents_v1_incident_event_incident_event_assign_proto_rawDescData = file_com_coralogixapis_incidents_v1_incident_event_incident_event_assign_proto_rawDesc
)

func file_com_coralogixapis_incidents_v1_incident_event_incident_event_assign_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_incidents_v1_incident_event_incident_event_assign_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_incidents_v1_incident_event_incident_event_assign_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_incidents_v1_incident_event_incident_event_assign_proto_rawDescData)
	})
	return file_com_coralogixapis_incidents_v1_incident_event_incident_event_assign_proto_rawDescData
}

var file_com_coralogixapis_incidents_v1_incident_event_incident_event_assign_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_incidents_v1_incident_event_incident_event_assign_proto_goTypes = []any{
	(*IncidentEventAssign)(nil), // 0: com.coralogixapis.incidents.v1.IncidentEventAssign
	(*Assignment)(nil),          // 1: com.coralogixapis.incidents.v1.Assignment
}
var file_com_coralogixapis_incidents_v1_incident_event_incident_event_assign_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.incidents.v1.IncidentEventAssign.assignment:type_name -> com.coralogixapis.incidents.v1.Assignment
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_incidents_v1_incident_event_incident_event_assign_proto_init() }
func file_com_coralogixapis_incidents_v1_incident_event_incident_event_assign_proto_init() {
	if File_com_coralogixapis_incidents_v1_incident_event_incident_event_assign_proto != nil {
		return
	}
	file_com_coralogixapis_incidents_v1_assignee_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_incidents_v1_incident_event_incident_event_assign_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_incidents_v1_incident_event_incident_event_assign_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_incidents_v1_incident_event_incident_event_assign_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_incidents_v1_incident_event_incident_event_assign_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_incidents_v1_incident_event_incident_event_assign_proto = out.File
	file_com_coralogixapis_incidents_v1_incident_event_incident_event_assign_proto_rawDesc = nil
	file_com_coralogixapis_incidents_v1_incident_event_incident_event_assign_proto_goTypes = nil
	file_com_coralogixapis_incidents_v1_incident_event_incident_event_assign_proto_depIdxs = nil
}
