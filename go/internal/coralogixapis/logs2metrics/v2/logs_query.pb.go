// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: com/coralogixapis/logs2metrics/v2/logs_query.proto

package v2

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

// enum that represents severity types
type Severity int32

const (
	Severity_SEVERITY_UNSPECIFIED Severity = 0
	Severity_SEVERITY_DEBUG       Severity = 1
	Severity_SEVERITY_VERBOSE     Severity = 2
	Severity_SEVERITY_INFO        Severity = 3
	Severity_SEVERITY_WARNING     Severity = 4
	Severity_SEVERITY_ERROR       Severity = 5
	Severity_SEVERITY_CRITICAL    Severity = 6
)

// Enum value maps for Severity.
var (
	Severity_name = map[int32]string{
		0: "SEVERITY_UNSPECIFIED",
		1: "SEVERITY_DEBUG",
		2: "SEVERITY_VERBOSE",
		3: "SEVERITY_INFO",
		4: "SEVERITY_WARNING",
		5: "SEVERITY_ERROR",
		6: "SEVERITY_CRITICAL",
	}
	Severity_value = map[string]int32{
		"SEVERITY_UNSPECIFIED": 0,
		"SEVERITY_DEBUG":       1,
		"SEVERITY_VERBOSE":     2,
		"SEVERITY_INFO":        3,
		"SEVERITY_WARNING":     4,
		"SEVERITY_ERROR":       5,
		"SEVERITY_CRITICAL":    6,
	}
)

func (x Severity) Enum() *Severity {
	p := new(Severity)
	*p = x
	return p
}

func (x Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_logs2metrics_v2_logs_query_proto_enumTypes[0].Descriptor()
}

func (Severity) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_logs2metrics_v2_logs_query_proto_enumTypes[0]
}

func (x Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Severity.Descriptor instead.
func (Severity) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_logs2metrics_v2_logs_query_proto_rawDescGZIP(), []int{0}
}

type LogsQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// lucene query
	Lucene *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=lucene,proto3" json:"lucene,omitempty"`
	// alias
	Alias *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
	// application name filters
	ApplicationnameFilters []*wrapperspb.StringValue `protobuf:"bytes,3,rep,name=applicationname_filters,json=applicationnameFilters,proto3" json:"applicationname_filters,omitempty"`
	// subsystem names filters
	SubsystemnameFilters []*wrapperspb.StringValue `protobuf:"bytes,4,rep,name=subsystemname_filters,json=subsystemnameFilters,proto3" json:"subsystemname_filters,omitempty"`
	// severity type filters
	SeverityFilters []Severity `protobuf:"varint,5,rep,packed,name=severity_filters,json=severityFilters,proto3,enum=com.coralogixapis.logs2metrics.v2.Severity" json:"severity_filters,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *LogsQuery) Reset() {
	*x = LogsQuery{}
	mi := &file_com_coralogixapis_logs2metrics_v2_logs_query_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsQuery) ProtoMessage() {}

func (x *LogsQuery) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_logs2metrics_v2_logs_query_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsQuery.ProtoReflect.Descriptor instead.
func (*LogsQuery) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_logs2metrics_v2_logs_query_proto_rawDescGZIP(), []int{0}
}

func (x *LogsQuery) GetLucene() *wrapperspb.StringValue {
	if x != nil {
		return x.Lucene
	}
	return nil
}

func (x *LogsQuery) GetAlias() *wrapperspb.StringValue {
	if x != nil {
		return x.Alias
	}
	return nil
}

func (x *LogsQuery) GetApplicationnameFilters() []*wrapperspb.StringValue {
	if x != nil {
		return x.ApplicationnameFilters
	}
	return nil
}

func (x *LogsQuery) GetSubsystemnameFilters() []*wrapperspb.StringValue {
	if x != nil {
		return x.SubsystemnameFilters
	}
	return nil
}

func (x *LogsQuery) GetSeverityFilters() []Severity {
	if x != nil {
		return x.SeverityFilters
	}
	return nil
}

var File_com_coralogixapis_logs2metrics_v2_logs_query_proto protoreflect.FileDescriptor

const file_com_coralogixapis_logs2metrics_v2_logs_query_proto_rawDesc = "" +
	"\n" +
	"2com/coralogixapis/logs2metrics/v2/logs_query.proto\x12!com.coralogixapis.logs2metrics.v2\x1a\x1egoogle/protobuf/wrappers.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\x93\x05\n" +
	"\tLogsQuery\x12_\n" +
	"\x06lucene\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueB)\x92A&J$\"log_obj.numeric_field: [50 TO 100]\"R\x06lucene\x12D\n" +
	"\x05alias\x18\x02 \x01(\v2\x1c.google.protobuf.StringValueB\x10\x92A\rJ\v\"new_query\"R\x05alias\x12f\n" +
	"\x17applicationname_filters\x18\x03 \x03(\v2\x1c.google.protobuf.StringValueB\x0f\x92A\fJ\n" +
	"\"app_name\"R\x16applicationnameFilters\x12b\n" +
	"\x15subsystemname_filters\x18\x04 \x03(\v2\x1c.google.protobuf.StringValueB\x0f\x92A\fJ\n" +
	"\"sub_name\"R\x14subsystemnameFilters\x12V\n" +
	"\x10severity_filters\x18\x05 \x03(\x0e2+.com.coralogixapis.logs2metrics.v2.SeverityR\x0fseverityFilters:\xba\x01\x92A\xb6\x01\n" +
	">*\n" +
	"SpansQuery20This data structure represents a query for logs.*t\n" +
	"\"Find out more about events2metrics\x12Nhttps://coralogix.com/docs/user-guides/monitoring-and-insights/events2metrics/*\xa2\x01\n" +
	"\bSeverity\x12\x18\n" +
	"\x14SEVERITY_UNSPECIFIED\x10\x00\x12\x12\n" +
	"\x0eSEVERITY_DEBUG\x10\x01\x12\x14\n" +
	"\x10SEVERITY_VERBOSE\x10\x02\x12\x11\n" +
	"\rSEVERITY_INFO\x10\x03\x12\x14\n" +
	"\x10SEVERITY_WARNING\x10\x04\x12\x12\n" +
	"\x0eSEVERITY_ERROR\x10\x05\x12\x15\n" +
	"\x11SEVERITY_CRITICAL\x10\x06b\x06proto3"

var (
	file_com_coralogixapis_logs2metrics_v2_logs_query_proto_rawDescOnce sync.Once
	file_com_coralogixapis_logs2metrics_v2_logs_query_proto_rawDescData []byte
)

func file_com_coralogixapis_logs2metrics_v2_logs_query_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_logs2metrics_v2_logs_query_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_logs2metrics_v2_logs_query_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_logs2metrics_v2_logs_query_proto_rawDesc), len(file_com_coralogixapis_logs2metrics_v2_logs_query_proto_rawDesc)))
	})
	return file_com_coralogixapis_logs2metrics_v2_logs_query_proto_rawDescData
}

var file_com_coralogixapis_logs2metrics_v2_logs_query_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_logs2metrics_v2_logs_query_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_logs2metrics_v2_logs_query_proto_goTypes = []any{
	(Severity)(0),                  // 0: com.coralogixapis.logs2metrics.v2.Severity
	(*LogsQuery)(nil),              // 1: com.coralogixapis.logs2metrics.v2.LogsQuery
	(*wrapperspb.StringValue)(nil), // 2: google.protobuf.StringValue
}
var file_com_coralogixapis_logs2metrics_v2_logs_query_proto_depIdxs = []int32{
	2, // 0: com.coralogixapis.logs2metrics.v2.LogsQuery.lucene:type_name -> google.protobuf.StringValue
	2, // 1: com.coralogixapis.logs2metrics.v2.LogsQuery.alias:type_name -> google.protobuf.StringValue
	2, // 2: com.coralogixapis.logs2metrics.v2.LogsQuery.applicationname_filters:type_name -> google.protobuf.StringValue
	2, // 3: com.coralogixapis.logs2metrics.v2.LogsQuery.subsystemname_filters:type_name -> google.protobuf.StringValue
	0, // 4: com.coralogixapis.logs2metrics.v2.LogsQuery.severity_filters:type_name -> com.coralogixapis.logs2metrics.v2.Severity
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_logs2metrics_v2_logs_query_proto_init() }
func file_com_coralogixapis_logs2metrics_v2_logs_query_proto_init() {
	if File_com_coralogixapis_logs2metrics_v2_logs_query_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_logs2metrics_v2_logs_query_proto_rawDesc), len(file_com_coralogixapis_logs2metrics_v2_logs_query_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_logs2metrics_v2_logs_query_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_logs2metrics_v2_logs_query_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_logs2metrics_v2_logs_query_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_logs2metrics_v2_logs_query_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_logs2metrics_v2_logs_query_proto = out.File
	file_com_coralogixapis_logs2metrics_v2_logs_query_proto_goTypes = nil
	file_com_coralogixapis_logs2metrics_v2_logs_query_proto_depIdxs = nil
}
