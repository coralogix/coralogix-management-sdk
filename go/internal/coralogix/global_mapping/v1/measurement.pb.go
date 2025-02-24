// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: com/coralogix/global_mapping/v1/measurement.proto

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

type Measurement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MeasurementName *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=measurement_name,json=measurementName,proto3" json:"measurement_name,omitempty"`
	DataSourceType  DataSourceType          `protobuf:"varint,2,opt,name=data_source_type,json=dataSourceType,proto3,enum=com.coralogix.global_mapping.v1.DataSourceType" json:"data_source_type,omitempty"`
	Query           *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty"`
	// Deprecated: Marked as deprecated in com/coralogix/global_mapping/v1/measurement.proto.
	Labels       []*wrapperspb.StringValue `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty"`
	Id           *wrapperspb.StringValue   `protobuf:"bytes,6,opt,name=id,proto3" json:"id,omitempty"`
	Description  *wrapperspb.StringValue   `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	DisplayName  *wrapperspb.StringValue   `protobuf:"bytes,8,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	LogicalGroup *wrapperspb.StringValue   `protobuf:"bytes,9,opt,name=logical_group,json=logicalGroup,proto3" json:"logical_group,omitempty"`
	// Deprecated: Marked as deprecated in com/coralogix/global_mapping/v1/measurement.proto.
	OrderingQuery     *wrapperspb.StringValue   `protobuf:"bytes,10,opt,name=ordering_query,json=orderingQuery,proto3" json:"ordering_query,omitempty"`
	AppendableMetrics []*wrapperspb.StringValue `protobuf:"bytes,11,rep,name=appendable_metrics,json=appendableMetrics,proto3" json:"appendable_metrics,omitempty"`
	DataSource        *DataSource               `protobuf:"bytes,12,opt,name=data_source,json=dataSource,proto3" json:"data_source,omitempty"`
	GroupLeft         *wrapperspb.StringValue   `protobuf:"bytes,13,opt,name=group_left,json=groupLeft,proto3" json:"group_left,omitempty"`
	JoinOn            *wrapperspb.StringValue   `protobuf:"bytes,14,opt,name=join_on,json=joinOn,proto3" json:"join_on,omitempty"`
	SubjectLabel      *wrapperspb.StringValue   `protobuf:"bytes,15,opt,name=subject_label,json=subjectLabel,proto3" json:"subject_label,omitempty"`
}

func (x *Measurement) Reset() {
	*x = Measurement{}
	mi := &file_com_coralogix_global_mapping_v1_measurement_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Measurement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Measurement) ProtoMessage() {}

func (x *Measurement) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_global_mapping_v1_measurement_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Measurement.ProtoReflect.Descriptor instead.
func (*Measurement) Descriptor() ([]byte, []int) {
	return file_com_coralogix_global_mapping_v1_measurement_proto_rawDescGZIP(), []int{0}
}

func (x *Measurement) GetMeasurementName() *wrapperspb.StringValue {
	if x != nil {
		return x.MeasurementName
	}
	return nil
}

func (x *Measurement) GetDataSourceType() DataSourceType {
	if x != nil {
		return x.DataSourceType
	}
	return DataSourceType_DATA_SOURCE_TYPE_UNSPECIFIED
}

func (x *Measurement) GetQuery() *wrapperspb.StringValue {
	if x != nil {
		return x.Query
	}
	return nil
}

// Deprecated: Marked as deprecated in com/coralogix/global_mapping/v1/measurement.proto.
func (x *Measurement) GetLabels() []*wrapperspb.StringValue {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Measurement) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Measurement) GetDescription() *wrapperspb.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *Measurement) GetDisplayName() *wrapperspb.StringValue {
	if x != nil {
		return x.DisplayName
	}
	return nil
}

func (x *Measurement) GetLogicalGroup() *wrapperspb.StringValue {
	if x != nil {
		return x.LogicalGroup
	}
	return nil
}

// Deprecated: Marked as deprecated in com/coralogix/global_mapping/v1/measurement.proto.
func (x *Measurement) GetOrderingQuery() *wrapperspb.StringValue {
	if x != nil {
		return x.OrderingQuery
	}
	return nil
}

func (x *Measurement) GetAppendableMetrics() []*wrapperspb.StringValue {
	if x != nil {
		return x.AppendableMetrics
	}
	return nil
}

func (x *Measurement) GetDataSource() *DataSource {
	if x != nil {
		return x.DataSource
	}
	return nil
}

func (x *Measurement) GetGroupLeft() *wrapperspb.StringValue {
	if x != nil {
		return x.GroupLeft
	}
	return nil
}

func (x *Measurement) GetJoinOn() *wrapperspb.StringValue {
	if x != nil {
		return x.JoinOn
	}
	return nil
}

func (x *Measurement) GetSubjectLabel() *wrapperspb.StringValue {
	if x != nil {
		return x.SubjectLabel
	}
	return nil
}

var File_com_coralogix_global_mapping_v1_measurement_proto protoreflect.FileDescriptor

var file_com_coralogix_global_mapping_v1_measurement_proto_rawDesc = []byte{
	0x0a, 0x31, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc8, 0x07, 0x0a, 0x0b, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x47, 0x0a, 0x10, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x59, 0x0a, 0x10, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x38, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x02, 0x18, 0x01, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x3e, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x3f, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x41, 0x0a, 0x0d, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x47, 0x0a, 0x0e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0d, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x4b, 0x0a, 0x12,
	0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x11, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x61, 0x62,
	0x6c, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x4c, 0x0a, 0x0b, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0a, 0x64, 0x61, 0x74,
	0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x4c, 0x65, 0x66, 0x74, 0x12, 0x35, 0x0a, 0x07, 0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x6f, 0x6e, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x06, 0x6a, 0x6f, 0x69, 0x6e, 0x4f, 0x6e, 0x12, 0x41, 0x0a, 0x0d, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0c, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4a, 0x04,
	0x08, 0x05, 0x10, 0x06, 0x52, 0x14, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_com_coralogix_global_mapping_v1_measurement_proto_rawDescOnce sync.Once
	file_com_coralogix_global_mapping_v1_measurement_proto_rawDescData = file_com_coralogix_global_mapping_v1_measurement_proto_rawDesc
)

func file_com_coralogix_global_mapping_v1_measurement_proto_rawDescGZIP() []byte {
	file_com_coralogix_global_mapping_v1_measurement_proto_rawDescOnce.Do(func() {
		file_com_coralogix_global_mapping_v1_measurement_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_global_mapping_v1_measurement_proto_rawDescData)
	})
	return file_com_coralogix_global_mapping_v1_measurement_proto_rawDescData
}

var file_com_coralogix_global_mapping_v1_measurement_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogix_global_mapping_v1_measurement_proto_goTypes = []any{
	(*Measurement)(nil),            // 0: com.coralogix.global_mapping.v1.Measurement
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
	(DataSourceType)(0),            // 2: com.coralogix.global_mapping.v1.DataSourceType
	(*DataSource)(nil),             // 3: com.coralogix.global_mapping.v1.DataSource
}
var file_com_coralogix_global_mapping_v1_measurement_proto_depIdxs = []int32{
	1,  // 0: com.coralogix.global_mapping.v1.Measurement.measurement_name:type_name -> google.protobuf.StringValue
	2,  // 1: com.coralogix.global_mapping.v1.Measurement.data_source_type:type_name -> com.coralogix.global_mapping.v1.DataSourceType
	1,  // 2: com.coralogix.global_mapping.v1.Measurement.query:type_name -> google.protobuf.StringValue
	1,  // 3: com.coralogix.global_mapping.v1.Measurement.labels:type_name -> google.protobuf.StringValue
	1,  // 4: com.coralogix.global_mapping.v1.Measurement.id:type_name -> google.protobuf.StringValue
	1,  // 5: com.coralogix.global_mapping.v1.Measurement.description:type_name -> google.protobuf.StringValue
	1,  // 6: com.coralogix.global_mapping.v1.Measurement.display_name:type_name -> google.protobuf.StringValue
	1,  // 7: com.coralogix.global_mapping.v1.Measurement.logical_group:type_name -> google.protobuf.StringValue
	1,  // 8: com.coralogix.global_mapping.v1.Measurement.ordering_query:type_name -> google.protobuf.StringValue
	1,  // 9: com.coralogix.global_mapping.v1.Measurement.appendable_metrics:type_name -> google.protobuf.StringValue
	3,  // 10: com.coralogix.global_mapping.v1.Measurement.data_source:type_name -> com.coralogix.global_mapping.v1.DataSource
	1,  // 11: com.coralogix.global_mapping.v1.Measurement.group_left:type_name -> google.protobuf.StringValue
	1,  // 12: com.coralogix.global_mapping.v1.Measurement.join_on:type_name -> google.protobuf.StringValue
	1,  // 13: com.coralogix.global_mapping.v1.Measurement.subject_label:type_name -> google.protobuf.StringValue
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_com_coralogix_global_mapping_v1_measurement_proto_init() }
func file_com_coralogix_global_mapping_v1_measurement_proto_init() {
	if File_com_coralogix_global_mapping_v1_measurement_proto != nil {
		return
	}
	file_com_coralogix_global_mapping_v1_data_source_type_proto_init()
	file_com_coralogix_global_mapping_v1_data_source_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_global_mapping_v1_measurement_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_global_mapping_v1_measurement_proto_goTypes,
		DependencyIndexes: file_com_coralogix_global_mapping_v1_measurement_proto_depIdxs,
		MessageInfos:      file_com_coralogix_global_mapping_v1_measurement_proto_msgTypes,
	}.Build()
	File_com_coralogix_global_mapping_v1_measurement_proto = out.File
	file_com_coralogix_global_mapping_v1_measurement_proto_rawDesc = nil
	file_com_coralogix_global_mapping_v1_measurement_proto_goTypes = nil
	file_com_coralogix_global_mapping_v1_measurement_proto_depIdxs = nil
}
