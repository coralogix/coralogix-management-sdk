// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: com/coralogix/global_mapping/v1/query_pagination.proto

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

type SortOrder int32

const (
	SortOrder_SORT_ORDER_DESCENDING_OR_UNSPECIFIED SortOrder = 0
	SortOrder_SORT_ORDER_ASCENDING                 SortOrder = 1
)

// Enum value maps for SortOrder.
var (
	SortOrder_name = map[int32]string{
		0: "SORT_ORDER_DESCENDING_OR_UNSPECIFIED",
		1: "SORT_ORDER_ASCENDING",
	}
	SortOrder_value = map[string]int32{
		"SORT_ORDER_DESCENDING_OR_UNSPECIFIED": 0,
		"SORT_ORDER_ASCENDING":                 1,
	}
)

func (x SortOrder) Enum() *SortOrder {
	p := new(SortOrder)
	*p = x
	return p
}

func (x SortOrder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortOrder) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogix_global_mapping_v1_query_pagination_proto_enumTypes[0].Descriptor()
}

func (SortOrder) Type() protoreflect.EnumType {
	return &file_com_coralogix_global_mapping_v1_query_pagination_proto_enumTypes[0]
}

func (x SortOrder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortOrder.Descriptor instead.
func (SortOrder) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogix_global_mapping_v1_query_pagination_proto_rawDescGZIP(), []int{0}
}

type PaginationData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize           *wrapperspb.Int32Value  `protobuf:"bytes,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageIndex          *wrapperspb.Int32Value  `protobuf:"bytes,2,opt,name=page_index,json=pageIndex,proto3" json:"page_index,omitempty"`
	OrderBy            SortOrder               `protobuf:"varint,3,opt,name=order_by,json=orderBy,proto3,enum=com.coralogix.global_mapping.v1.SortOrder" json:"order_by,omitempty"`
	OrderByMeasurement *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=order_by_measurement,json=orderByMeasurement,proto3" json:"order_by_measurement,omitempty"`
}

func (x *PaginationData) Reset() {
	*x = PaginationData{}
	mi := &file_com_coralogix_global_mapping_v1_query_pagination_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaginationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationData) ProtoMessage() {}

func (x *PaginationData) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_global_mapping_v1_query_pagination_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationData.ProtoReflect.Descriptor instead.
func (*PaginationData) Descriptor() ([]byte, []int) {
	return file_com_coralogix_global_mapping_v1_query_pagination_proto_rawDescGZIP(), []int{0}
}

func (x *PaginationData) GetPageSize() *wrapperspb.Int32Value {
	if x != nil {
		return x.PageSize
	}
	return nil
}

func (x *PaginationData) GetPageIndex() *wrapperspb.Int32Value {
	if x != nil {
		return x.PageIndex
	}
	return nil
}

func (x *PaginationData) GetOrderBy() SortOrder {
	if x != nil {
		return x.OrderBy
	}
	return SortOrder_SORT_ORDER_DESCENDING_OR_UNSPECIFIED
}

func (x *PaginationData) GetOrderByMeasurement() *wrapperspb.StringValue {
	if x != nil {
		return x.OrderByMeasurement
	}
	return nil
}

type TableRow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Queries []*MeasurementQuery `protobuf:"bytes,1,rep,name=queries,proto3" json:"queries,omitempty"`
}

func (x *TableRow) Reset() {
	*x = TableRow{}
	mi := &file_com_coralogix_global_mapping_v1_query_pagination_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TableRow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableRow) ProtoMessage() {}

func (x *TableRow) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_global_mapping_v1_query_pagination_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableRow.ProtoReflect.Descriptor instead.
func (*TableRow) Descriptor() ([]byte, []int) {
	return file_com_coralogix_global_mapping_v1_query_pagination_proto_rawDescGZIP(), []int{1}
}

func (x *TableRow) GetQueries() []*MeasurementQuery {
	if x != nil {
		return x.Queries
	}
	return nil
}

var File_com_coralogix_global_mapping_v1_query_pagination_proto protoreflect.FileDescriptor

var file_com_coralogix_global_mapping_v1_query_pagination_proto_rawDesc = []byte{
	0x0a, 0x36, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x37, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f,
	0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x9d, 0x02, 0x0a, 0x0e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x3a, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x45, 0x0a, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x42, 0x79, 0x12, 0x4e, 0x0a, 0x14, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x5f, 0x6d,
	0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x12,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0x57, 0x0a, 0x08, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x6f, 0x77, 0x12, 0x4b,
	0x0a, 0x07, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x07, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2a, 0x4f, 0x0a, 0x09, 0x53,
	0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x24, 0x53, 0x4f, 0x52, 0x54,
	0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e,
	0x47, 0x5f, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x4f, 0x52, 0x54, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52,
	0x5f, 0x41, 0x53, 0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_global_mapping_v1_query_pagination_proto_rawDescOnce sync.Once
	file_com_coralogix_global_mapping_v1_query_pagination_proto_rawDescData = file_com_coralogix_global_mapping_v1_query_pagination_proto_rawDesc
)

func file_com_coralogix_global_mapping_v1_query_pagination_proto_rawDescGZIP() []byte {
	file_com_coralogix_global_mapping_v1_query_pagination_proto_rawDescOnce.Do(func() {
		file_com_coralogix_global_mapping_v1_query_pagination_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_global_mapping_v1_query_pagination_proto_rawDescData)
	})
	return file_com_coralogix_global_mapping_v1_query_pagination_proto_rawDescData
}

var file_com_coralogix_global_mapping_v1_query_pagination_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogix_global_mapping_v1_query_pagination_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogix_global_mapping_v1_query_pagination_proto_goTypes = []any{
	(SortOrder)(0),                 // 0: com.coralogix.global_mapping.v1.SortOrder
	(*PaginationData)(nil),         // 1: com.coralogix.global_mapping.v1.PaginationData
	(*TableRow)(nil),               // 2: com.coralogix.global_mapping.v1.TableRow
	(*wrapperspb.Int32Value)(nil),  // 3: google.protobuf.Int32Value
	(*wrapperspb.StringValue)(nil), // 4: google.protobuf.StringValue
	(*MeasurementQuery)(nil),       // 5: com.coralogix.global_mapping.v1.MeasurementQuery
}
var file_com_coralogix_global_mapping_v1_query_pagination_proto_depIdxs = []int32{
	3, // 0: com.coralogix.global_mapping.v1.PaginationData.page_size:type_name -> google.protobuf.Int32Value
	3, // 1: com.coralogix.global_mapping.v1.PaginationData.page_index:type_name -> google.protobuf.Int32Value
	0, // 2: com.coralogix.global_mapping.v1.PaginationData.order_by:type_name -> com.coralogix.global_mapping.v1.SortOrder
	4, // 3: com.coralogix.global_mapping.v1.PaginationData.order_by_measurement:type_name -> google.protobuf.StringValue
	5, // 4: com.coralogix.global_mapping.v1.TableRow.queries:type_name -> com.coralogix.global_mapping.v1.MeasurementQuery
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_com_coralogix_global_mapping_v1_query_pagination_proto_init() }
func file_com_coralogix_global_mapping_v1_query_pagination_proto_init() {
	if File_com_coralogix_global_mapping_v1_query_pagination_proto != nil {
		return
	}
	file_com_coralogix_global_mapping_v1_measurement_query_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_global_mapping_v1_query_pagination_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_global_mapping_v1_query_pagination_proto_goTypes,
		DependencyIndexes: file_com_coralogix_global_mapping_v1_query_pagination_proto_depIdxs,
		EnumInfos:         file_com_coralogix_global_mapping_v1_query_pagination_proto_enumTypes,
		MessageInfos:      file_com_coralogix_global_mapping_v1_query_pagination_proto_msgTypes,
	}.Build()
	File_com_coralogix_global_mapping_v1_query_pagination_proto = out.File
	file_com_coralogix_global_mapping_v1_query_pagination_proto_rawDesc = nil
	file_com_coralogix_global_mapping_v1_query_pagination_proto_goTypes = nil
	file_com_coralogix_global_mapping_v1_query_pagination_proto_depIdxs = nil
}
