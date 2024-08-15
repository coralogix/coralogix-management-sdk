// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: com/coralogix/dataprime/v1/list_queries.proto

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

type ListQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListQueryRequest) Reset() {
	*x = ListQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_dataprime_v1_list_queries_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQueryRequest) ProtoMessage() {}

func (x *ListQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_list_queries_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQueryRequest.ProtoReflect.Descriptor instead.
func (*ListQueryRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_list_queries_proto_rawDescGZIP(), []int{0}
}

type RunningQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// how long the query has been running (in wall clock time)
	DurationsMs uint64 `protobuf:"varint,2,opt,name=durations_ms,json=durationsMs,proto3" json:"durations_ms,omitempty"`
	// total aggregated duration of all completed tasks
	TotalTaskDurationMs uint64 `protobuf:"varint,3,opt,name=total_task_duration_ms,json=totalTaskDurationMs,proto3" json:"total_task_duration_ms,omitempty"`
	// How many tasks total in the query
	TotalTasks uint32 `protobuf:"varint,4,opt,name=total_tasks,json=totalTasks,proto3" json:"total_tasks,omitempty"`
	// How many tasks are completed
	CompletedTasks uint32 `protobuf:"varint,5,opt,name=completed_tasks,json=completedTasks,proto3" json:"completed_tasks,omitempty"`
}

func (x *RunningQuery) Reset() {
	*x = RunningQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_dataprime_v1_list_queries_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunningQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunningQuery) ProtoMessage() {}

func (x *RunningQuery) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_list_queries_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunningQuery.ProtoReflect.Descriptor instead.
func (*RunningQuery) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_list_queries_proto_rawDescGZIP(), []int{1}
}

func (x *RunningQuery) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RunningQuery) GetDurationsMs() uint64 {
	if x != nil {
		return x.DurationsMs
	}
	return 0
}

func (x *RunningQuery) GetTotalTaskDurationMs() uint64 {
	if x != nil {
		return x.TotalTaskDurationMs
	}
	return 0
}

func (x *RunningQuery) GetTotalTasks() uint32 {
	if x != nil {
		return x.TotalTasks
	}
	return 0
}

func (x *RunningQuery) GetCompletedTasks() uint32 {
	if x != nil {
		return x.CompletedTasks
	}
	return 0
}

type ListQueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RunningQueries []*RunningQuery `protobuf:"bytes,1,rep,name=running_queries,json=runningQueries,proto3" json:"running_queries,omitempty"`
}

func (x *ListQueryResponse) Reset() {
	*x = ListQueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_dataprime_v1_list_queries_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQueryResponse) ProtoMessage() {}

func (x *ListQueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_dataprime_v1_list_queries_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQueryResponse.ProtoReflect.Descriptor instead.
func (*ListQueryResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogix_dataprime_v1_list_queries_proto_rawDescGZIP(), []int{2}
}

func (x *ListQueryResponse) GetRunningQueries() []*RunningQuery {
	if x != nil {
		return x.RunningQueries
	}
	return nil
}

var File_com_coralogix_dataprime_v1_list_queries_proto protoreflect.FileDescriptor

var file_com_coralogix_dataprime_v1_list_queries_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73,
	0x74, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x12, 0x0a, 0x10, 0x4c,
	0x69, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0xc0, 0x01, 0x0a, 0x0c, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x6d, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x4d, 0x73, 0x12, 0x33, 0x0a, 0x16, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x74, 0x61, 0x73,
	0x6b, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x13, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x54, 0x61, 0x73,
	0x6b, 0x73, 0x22, 0x66, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0f, 0x72, 0x75, 0x6e, 0x6e, 0x69,
	0x6e, 0x67, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75,
	0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x0e, 0x72, 0x75, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_com_coralogix_dataprime_v1_list_queries_proto_rawDescOnce sync.Once
	file_com_coralogix_dataprime_v1_list_queries_proto_rawDescData = file_com_coralogix_dataprime_v1_list_queries_proto_rawDesc
)

func file_com_coralogix_dataprime_v1_list_queries_proto_rawDescGZIP() []byte {
	file_com_coralogix_dataprime_v1_list_queries_proto_rawDescOnce.Do(func() {
		file_com_coralogix_dataprime_v1_list_queries_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_dataprime_v1_list_queries_proto_rawDescData)
	})
	return file_com_coralogix_dataprime_v1_list_queries_proto_rawDescData
}

var file_com_coralogix_dataprime_v1_list_queries_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogix_dataprime_v1_list_queries_proto_goTypes = []any{
	(*ListQueryRequest)(nil),  // 0: com.coralogix.dataprime.v1.ListQueryRequest
	(*RunningQuery)(nil),      // 1: com.coralogix.dataprime.v1.RunningQuery
	(*ListQueryResponse)(nil), // 2: com.coralogix.dataprime.v1.ListQueryResponse
}
var file_com_coralogix_dataprime_v1_list_queries_proto_depIdxs = []int32{
	1, // 0: com.coralogix.dataprime.v1.ListQueryResponse.running_queries:type_name -> com.coralogix.dataprime.v1.RunningQuery
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_com_coralogix_dataprime_v1_list_queries_proto_init() }
func file_com_coralogix_dataprime_v1_list_queries_proto_init() {
	if File_com_coralogix_dataprime_v1_list_queries_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_dataprime_v1_list_queries_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ListQueryRequest); i {
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
		file_com_coralogix_dataprime_v1_list_queries_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RunningQuery); i {
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
		file_com_coralogix_dataprime_v1_list_queries_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListQueryResponse); i {
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
			RawDescriptor: file_com_coralogix_dataprime_v1_list_queries_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_dataprime_v1_list_queries_proto_goTypes,
		DependencyIndexes: file_com_coralogix_dataprime_v1_list_queries_proto_depIdxs,
		MessageInfos:      file_com_coralogix_dataprime_v1_list_queries_proto_msgTypes,
	}.Build()
	File_com_coralogix_dataprime_v1_list_queries_proto = out.File
	file_com_coralogix_dataprime_v1_list_queries_proto_rawDesc = nil
	file_com_coralogix_dataprime_v1_list_queries_proto_goTypes = nil
	file_com_coralogix_dataprime_v1_list_queries_proto_depIdxs = nil
}
