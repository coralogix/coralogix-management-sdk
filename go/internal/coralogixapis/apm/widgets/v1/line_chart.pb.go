// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: com/coralogixapis/apm/widgets/v1/line_chart.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

type InternalMetricSchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service         *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Operation       *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	Error           *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	Method          *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=method,proto3" json:"method,omitempty"`
	SpanKind        *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=spanKind,proto3" json:"spanKind,omitempty"`
	PeerService     *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=peerService,proto3" json:"peerService,omitempty"`
	Transaction     *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=transaction,proto3" json:"transaction,omitempty"`
	TransactionRoot *wrapperspb.StringValue `protobuf:"bytes,8,opt,name=transactionRoot,proto3" json:"transactionRoot,omitempty"`
	HttpStatusCode  *wrapperspb.StringValue `protobuf:"bytes,9,opt,name=httpStatusCode,proto3" json:"httpStatusCode,omitempty"`
	GrpcStatusCode  *wrapperspb.StringValue `protobuf:"bytes,10,opt,name=grpcStatusCode,proto3" json:"grpcStatusCode,omitempty"`
	DbOperation     *wrapperspb.StringValue `protobuf:"bytes,11,opt,name=dbOperation,proto3" json:"dbOperation,omitempty"`
	ServiceVersion  *wrapperspb.StringValue `protobuf:"bytes,12,opt,name=service_version,json=serviceVersion,proto3" json:"service_version,omitempty"`
}

func (x *InternalMetricSchema) Reset() {
	*x = InternalMetricSchema{}
	mi := &file_com_coralogixapis_apm_widgets_v1_line_chart_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InternalMetricSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalMetricSchema) ProtoMessage() {}

func (x *InternalMetricSchema) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_widgets_v1_line_chart_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalMetricSchema.ProtoReflect.Descriptor instead.
func (*InternalMetricSchema) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_widgets_v1_line_chart_proto_rawDescGZIP(), []int{0}
}

func (x *InternalMetricSchema) GetService() *wrapperspb.StringValue {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *InternalMetricSchema) GetOperation() *wrapperspb.StringValue {
	if x != nil {
		return x.Operation
	}
	return nil
}

func (x *InternalMetricSchema) GetError() *wrapperspb.StringValue {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *InternalMetricSchema) GetMethod() *wrapperspb.StringValue {
	if x != nil {
		return x.Method
	}
	return nil
}

func (x *InternalMetricSchema) GetSpanKind() *wrapperspb.StringValue {
	if x != nil {
		return x.SpanKind
	}
	return nil
}

func (x *InternalMetricSchema) GetPeerService() *wrapperspb.StringValue {
	if x != nil {
		return x.PeerService
	}
	return nil
}

func (x *InternalMetricSchema) GetTransaction() *wrapperspb.StringValue {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *InternalMetricSchema) GetTransactionRoot() *wrapperspb.StringValue {
	if x != nil {
		return x.TransactionRoot
	}
	return nil
}

func (x *InternalMetricSchema) GetHttpStatusCode() *wrapperspb.StringValue {
	if x != nil {
		return x.HttpStatusCode
	}
	return nil
}

func (x *InternalMetricSchema) GetGrpcStatusCode() *wrapperspb.StringValue {
	if x != nil {
		return x.GrpcStatusCode
	}
	return nil
}

func (x *InternalMetricSchema) GetDbOperation() *wrapperspb.StringValue {
	if x != nil {
		return x.DbOperation
	}
	return nil
}

func (x *InternalMetricSchema) GetServiceVersion() *wrapperspb.StringValue {
	if x != nil {
		return x.ServiceVersion
	}
	return nil
}

type LineChart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DisplayName          *wrapperspb.StringValue            `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Name                 *wrapperspb.StringValue            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Query                *wrapperspb.StringValue            `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty"`
	Points               []*structpb.ListValue              `protobuf:"bytes,4,rep,name=points,proto3" json:"points,omitempty"`
	ToolTip              *wrapperspb.StringValue            `protobuf:"bytes,5,opt,name=tool_tip,json=toolTip,proto3" json:"tool_tip,omitempty"`
	Unit                 Unit                               `protobuf:"varint,6,opt,name=unit,proto3,enum=com.coralogixapis.apm.widgets.v1.Unit" json:"unit,omitempty"`
	Metric               map[string]*wrapperspb.StringValue `protobuf:"bytes,7,rep,name=metric,proto3" json:"metric,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	InternalMetricSchema *InternalMetricSchema              `protobuf:"bytes,8,opt,name=internalMetricSchema,proto3" json:"internalMetricSchema,omitempty"`
}

func (x *LineChart) Reset() {
	*x = LineChart{}
	mi := &file_com_coralogixapis_apm_widgets_v1_line_chart_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LineChart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineChart) ProtoMessage() {}

func (x *LineChart) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_apm_widgets_v1_line_chart_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineChart.ProtoReflect.Descriptor instead.
func (*LineChart) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_apm_widgets_v1_line_chart_proto_rawDescGZIP(), []int{1}
}

func (x *LineChart) GetDisplayName() *wrapperspb.StringValue {
	if x != nil {
		return x.DisplayName
	}
	return nil
}

func (x *LineChart) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *LineChart) GetQuery() *wrapperspb.StringValue {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *LineChart) GetPoints() []*structpb.ListValue {
	if x != nil {
		return x.Points
	}
	return nil
}

func (x *LineChart) GetToolTip() *wrapperspb.StringValue {
	if x != nil {
		return x.ToolTip
	}
	return nil
}

func (x *LineChart) GetUnit() Unit {
	if x != nil {
		return x.Unit
	}
	return Unit_UNIT_UNSPECIFIED
}

func (x *LineChart) GetMetric() map[string]*wrapperspb.StringValue {
	if x != nil {
		return x.Metric
	}
	return nil
}

func (x *LineChart) GetInternalMetricSchema() *InternalMetricSchema {
	if x != nil {
		return x.InternalMetricSchema
	}
	return nil
}

var File_com_coralogixapis_apm_widgets_v1_line_chart_proto protoreflect.FileDescriptor

var file_com_coralogixapis_apm_widgets_v1_line_chart_proto_rawDesc = []byte{
	0x0a, 0x31, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x6d, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65,
	0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x89, 0x06, 0x0a, 0x14, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x36, 0x0a, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x34, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x38, 0x0a, 0x08, 0x73, 0x70, 0x61, 0x6e,
	0x4b, 0x69, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x73, 0x70, 0x61, 0x6e, 0x4b, 0x69,
	0x6e, 0x64, 0x12, 0x3e, 0x0a, 0x0b, 0x70, 0x65, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x70, 0x65, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x6f, 0x6f, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x44, 0x0a, 0x0e, 0x68, 0x74,
	0x74, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0e, 0x68, 0x74, 0x74, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x44, 0x0a, 0x0e, 0x67, 0x72, 0x70, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x67, 0x72, 0x70, 0x63, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x64, 0x62, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x62, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xf1, 0x04,
	0x0a, 0x09, 0x4c, 0x69, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x72, 0x74, 0x12, 0x3f, 0x0a, 0x0c, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x32,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x12, 0x32, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x37, 0x0a, 0x08, 0x74, 0x6f, 0x6f, 0x6c, 0x5f, 0x74,
	0x69, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x74, 0x6f, 0x6f, 0x6c, 0x54, 0x69, 0x70, 0x12,
	0x3a, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x4f, 0x0a, 0x06, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x61, 0x70, 0x6d, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x72, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x6a, 0x0a, 0x14,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61,
	0x70, 0x6d, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x52, 0x14, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x1a, 0x57, 0x0a, 0x0b, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_apm_widgets_v1_line_chart_proto_rawDescOnce sync.Once
	file_com_coralogixapis_apm_widgets_v1_line_chart_proto_rawDescData = file_com_coralogixapis_apm_widgets_v1_line_chart_proto_rawDesc
)

func file_com_coralogixapis_apm_widgets_v1_line_chart_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_apm_widgets_v1_line_chart_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_apm_widgets_v1_line_chart_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_apm_widgets_v1_line_chart_proto_rawDescData)
	})
	return file_com_coralogixapis_apm_widgets_v1_line_chart_proto_rawDescData
}

var file_com_coralogixapis_apm_widgets_v1_line_chart_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogixapis_apm_widgets_v1_line_chart_proto_goTypes = []any{
	(*InternalMetricSchema)(nil),   // 0: com.coralogixapis.apm.widgets.v1.InternalMetricSchema
	(*LineChart)(nil),              // 1: com.coralogixapis.apm.widgets.v1.LineChart
	nil,                            // 2: com.coralogixapis.apm.widgets.v1.LineChart.MetricEntry
	(*wrapperspb.StringValue)(nil), // 3: google.protobuf.StringValue
	(*structpb.ListValue)(nil),     // 4: google.protobuf.ListValue
	(Unit)(0),                      // 5: com.coralogixapis.apm.widgets.v1.Unit
}
var file_com_coralogixapis_apm_widgets_v1_line_chart_proto_depIdxs = []int32{
	3,  // 0: com.coralogixapis.apm.widgets.v1.InternalMetricSchema.service:type_name -> google.protobuf.StringValue
	3,  // 1: com.coralogixapis.apm.widgets.v1.InternalMetricSchema.operation:type_name -> google.protobuf.StringValue
	3,  // 2: com.coralogixapis.apm.widgets.v1.InternalMetricSchema.error:type_name -> google.protobuf.StringValue
	3,  // 3: com.coralogixapis.apm.widgets.v1.InternalMetricSchema.method:type_name -> google.protobuf.StringValue
	3,  // 4: com.coralogixapis.apm.widgets.v1.InternalMetricSchema.spanKind:type_name -> google.protobuf.StringValue
	3,  // 5: com.coralogixapis.apm.widgets.v1.InternalMetricSchema.peerService:type_name -> google.protobuf.StringValue
	3,  // 6: com.coralogixapis.apm.widgets.v1.InternalMetricSchema.transaction:type_name -> google.protobuf.StringValue
	3,  // 7: com.coralogixapis.apm.widgets.v1.InternalMetricSchema.transactionRoot:type_name -> google.protobuf.StringValue
	3,  // 8: com.coralogixapis.apm.widgets.v1.InternalMetricSchema.httpStatusCode:type_name -> google.protobuf.StringValue
	3,  // 9: com.coralogixapis.apm.widgets.v1.InternalMetricSchema.grpcStatusCode:type_name -> google.protobuf.StringValue
	3,  // 10: com.coralogixapis.apm.widgets.v1.InternalMetricSchema.dbOperation:type_name -> google.protobuf.StringValue
	3,  // 11: com.coralogixapis.apm.widgets.v1.InternalMetricSchema.service_version:type_name -> google.protobuf.StringValue
	3,  // 12: com.coralogixapis.apm.widgets.v1.LineChart.display_name:type_name -> google.protobuf.StringValue
	3,  // 13: com.coralogixapis.apm.widgets.v1.LineChart.name:type_name -> google.protobuf.StringValue
	3,  // 14: com.coralogixapis.apm.widgets.v1.LineChart.query:type_name -> google.protobuf.StringValue
	4,  // 15: com.coralogixapis.apm.widgets.v1.LineChart.points:type_name -> google.protobuf.ListValue
	3,  // 16: com.coralogixapis.apm.widgets.v1.LineChart.tool_tip:type_name -> google.protobuf.StringValue
	5,  // 17: com.coralogixapis.apm.widgets.v1.LineChart.unit:type_name -> com.coralogixapis.apm.widgets.v1.Unit
	2,  // 18: com.coralogixapis.apm.widgets.v1.LineChart.metric:type_name -> com.coralogixapis.apm.widgets.v1.LineChart.MetricEntry
	0,  // 19: com.coralogixapis.apm.widgets.v1.LineChart.internalMetricSchema:type_name -> com.coralogixapis.apm.widgets.v1.InternalMetricSchema
	3,  // 20: com.coralogixapis.apm.widgets.v1.LineChart.MetricEntry.value:type_name -> google.protobuf.StringValue
	21, // [21:21] is the sub-list for method output_type
	21, // [21:21] is the sub-list for method input_type
	21, // [21:21] is the sub-list for extension type_name
	21, // [21:21] is the sub-list for extension extendee
	0,  // [0:21] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_apm_widgets_v1_line_chart_proto_init() }
func file_com_coralogixapis_apm_widgets_v1_line_chart_proto_init() {
	if File_com_coralogixapis_apm_widgets_v1_line_chart_proto != nil {
		return
	}
	file_com_coralogixapis_apm_widgets_v1_units_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_apm_widgets_v1_line_chart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_apm_widgets_v1_line_chart_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_apm_widgets_v1_line_chart_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_apm_widgets_v1_line_chart_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_apm_widgets_v1_line_chart_proto = out.File
	file_com_coralogixapis_apm_widgets_v1_line_chart_proto_rawDesc = nil
	file_com_coralogixapis_apm_widgets_v1_line_chart_proto_goTypes = nil
	file_com_coralogixapis_apm_widgets_v1_line_chart_proto_depIdxs = nil
}
