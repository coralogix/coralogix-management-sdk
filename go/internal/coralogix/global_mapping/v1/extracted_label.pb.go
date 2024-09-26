// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.28.1
// source: com/coralogix/global_mapping/v1/extracted_label.proto

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

type DataSourceTypeValues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataSourceType            DataSourceType            `protobuf:"varint,1,opt,name=data_source_type,json=dataSourceType,proto3,enum=com.coralogix.global_mapping.v1.DataSourceType" json:"data_source_type,omitempty"`
	DestinationExtractionKeys []*wrapperspb.StringValue `protobuf:"bytes,2,rep,name=destination_extraction_keys,json=destinationExtractionKeys,proto3" json:"destination_extraction_keys,omitempty"`
}

func (x *DataSourceTypeValues) Reset() {
	*x = DataSourceTypeValues{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_global_mapping_v1_extracted_label_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataSourceTypeValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSourceTypeValues) ProtoMessage() {}

func (x *DataSourceTypeValues) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_global_mapping_v1_extracted_label_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSourceTypeValues.ProtoReflect.Descriptor instead.
func (*DataSourceTypeValues) Descriptor() ([]byte, []int) {
	return file_com_coralogix_global_mapping_v1_extracted_label_proto_rawDescGZIP(), []int{0}
}

func (x *DataSourceTypeValues) GetDataSourceType() DataSourceType {
	if x != nil {
		return x.DataSourceType
	}
	return DataSourceType_DATA_SOURCE_TYPE_UNSPECIFIED
}

func (x *DataSourceTypeValues) GetDestinationExtractionKeys() []*wrapperspb.StringValue {
	if x != nil {
		return x.DestinationExtractionKeys
	}
	return nil
}

type ExtractedLabel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label  *wrapperspb.StringValue   `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Values []*wrapperspb.StringValue `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	// Deprecated: Marked as deprecated in com/coralogix/global_mapping/v1/extracted_label.proto.
	DestinationExtractionKey *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=destination_extraction_key,json=destinationExtractionKey,proto3" json:"destination_extraction_key,omitempty"`
	Description              *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	DisplayName              *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Deprecated: Marked as deprecated in com/coralogix/global_mapping/v1/extracted_label.proto.
	DestinationExtractionKeys     []*wrapperspb.StringValue `protobuf:"bytes,6,rep,name=destination_extraction_keys,json=destinationExtractionKeys,proto3" json:"destination_extraction_keys,omitempty"`
	DestinationTypeExtractionKeys []*DataSourceTypeValues   `protobuf:"bytes,7,rep,name=destination_type_extraction_keys,json=destinationTypeExtractionKeys,proto3" json:"destination_type_extraction_keys,omitempty"`
	DataSource                    *DataSource               `protobuf:"bytes,9,opt,name=data_source,json=dataSource,proto3" json:"data_source,omitempty"`
	IsCustomLabel                 *wrapperspb.BoolValue     `protobuf:"bytes,10,opt,name=is_custom_label,json=isCustomLabel,proto3" json:"is_custom_label,omitempty"`
}

func (x *ExtractedLabel) Reset() {
	*x = ExtractedLabel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_global_mapping_v1_extracted_label_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtractedLabel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtractedLabel) ProtoMessage() {}

func (x *ExtractedLabel) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_global_mapping_v1_extracted_label_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtractedLabel.ProtoReflect.Descriptor instead.
func (*ExtractedLabel) Descriptor() ([]byte, []int) {
	return file_com_coralogix_global_mapping_v1_extracted_label_proto_rawDescGZIP(), []int{1}
}

func (x *ExtractedLabel) GetLabel() *wrapperspb.StringValue {
	if x != nil {
		return x.Label
	}
	return nil
}

func (x *ExtractedLabel) GetValues() []*wrapperspb.StringValue {
	if x != nil {
		return x.Values
	}
	return nil
}

// Deprecated: Marked as deprecated in com/coralogix/global_mapping/v1/extracted_label.proto.
func (x *ExtractedLabel) GetDestinationExtractionKey() *wrapperspb.StringValue {
	if x != nil {
		return x.DestinationExtractionKey
	}
	return nil
}

func (x *ExtractedLabel) GetDescription() *wrapperspb.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *ExtractedLabel) GetDisplayName() *wrapperspb.StringValue {
	if x != nil {
		return x.DisplayName
	}
	return nil
}

// Deprecated: Marked as deprecated in com/coralogix/global_mapping/v1/extracted_label.proto.
func (x *ExtractedLabel) GetDestinationExtractionKeys() []*wrapperspb.StringValue {
	if x != nil {
		return x.DestinationExtractionKeys
	}
	return nil
}

func (x *ExtractedLabel) GetDestinationTypeExtractionKeys() []*DataSourceTypeValues {
	if x != nil {
		return x.DestinationTypeExtractionKeys
	}
	return nil
}

func (x *ExtractedLabel) GetDataSource() *DataSource {
	if x != nil {
		return x.DataSource
	}
	return nil
}

func (x *ExtractedLabel) GetIsCustomLabel() *wrapperspb.BoolValue {
	if x != nil {
		return x.IsCustomLabel
	}
	return nil
}

var File_com_coralogix_global_mapping_v1_extracted_label_proto protoreflect.FileDescriptor

var file_com_coralogix_global_mapping_v1_extracted_label_proto_rawDesc = []byte{
	0x0a, 0x35, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x31, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x01, 0x0a, 0x14, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x59, 0x0a, 0x10,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x5c, 0x0a, 0x1b, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x19, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4b, 0x65, 0x79, 0x73, 0x22, 0xef, 0x05, 0x0a, 0x0e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x65, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x32, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x38, 0x0a, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x02, 0x10, 0x00, 0x52, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x5e, 0x0a, 0x1a, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x02, 0x18, 0x01, 0x52, 0x18, 0x64, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x3e, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x60, 0x0a, 0x1b, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x02, 0x18, 0x01, 0x52, 0x19,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x7e, 0x0a, 0x20, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x1d, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x4c, 0x0a, 0x0b, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0a, 0x64, 0x61, 0x74,
	0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x69, 0x73,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4a, 0x04, 0x08, 0x08, 0x10,
	0x09, 0x52, 0x14, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_global_mapping_v1_extracted_label_proto_rawDescOnce sync.Once
	file_com_coralogix_global_mapping_v1_extracted_label_proto_rawDescData = file_com_coralogix_global_mapping_v1_extracted_label_proto_rawDesc
)

func file_com_coralogix_global_mapping_v1_extracted_label_proto_rawDescGZIP() []byte {
	file_com_coralogix_global_mapping_v1_extracted_label_proto_rawDescOnce.Do(func() {
		file_com_coralogix_global_mapping_v1_extracted_label_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_global_mapping_v1_extracted_label_proto_rawDescData)
	})
	return file_com_coralogix_global_mapping_v1_extracted_label_proto_rawDescData
}

var file_com_coralogix_global_mapping_v1_extracted_label_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_coralogix_global_mapping_v1_extracted_label_proto_goTypes = []interface{}{
	(*DataSourceTypeValues)(nil),   // 0: com.coralogix.global_mapping.v1.DataSourceTypeValues
	(*ExtractedLabel)(nil),         // 1: com.coralogix.global_mapping.v1.ExtractedLabel
	(DataSourceType)(0),            // 2: com.coralogix.global_mapping.v1.DataSourceType
	(*wrapperspb.StringValue)(nil), // 3: google.protobuf.StringValue
	(*DataSource)(nil),             // 4: com.coralogix.global_mapping.v1.DataSource
	(*wrapperspb.BoolValue)(nil),   // 5: google.protobuf.BoolValue
}
var file_com_coralogix_global_mapping_v1_extracted_label_proto_depIdxs = []int32{
	2,  // 0: com.coralogix.global_mapping.v1.DataSourceTypeValues.data_source_type:type_name -> com.coralogix.global_mapping.v1.DataSourceType
	3,  // 1: com.coralogix.global_mapping.v1.DataSourceTypeValues.destination_extraction_keys:type_name -> google.protobuf.StringValue
	3,  // 2: com.coralogix.global_mapping.v1.ExtractedLabel.label:type_name -> google.protobuf.StringValue
	3,  // 3: com.coralogix.global_mapping.v1.ExtractedLabel.values:type_name -> google.protobuf.StringValue
	3,  // 4: com.coralogix.global_mapping.v1.ExtractedLabel.destination_extraction_key:type_name -> google.protobuf.StringValue
	3,  // 5: com.coralogix.global_mapping.v1.ExtractedLabel.description:type_name -> google.protobuf.StringValue
	3,  // 6: com.coralogix.global_mapping.v1.ExtractedLabel.display_name:type_name -> google.protobuf.StringValue
	3,  // 7: com.coralogix.global_mapping.v1.ExtractedLabel.destination_extraction_keys:type_name -> google.protobuf.StringValue
	0,  // 8: com.coralogix.global_mapping.v1.ExtractedLabel.destination_type_extraction_keys:type_name -> com.coralogix.global_mapping.v1.DataSourceTypeValues
	4,  // 9: com.coralogix.global_mapping.v1.ExtractedLabel.data_source:type_name -> com.coralogix.global_mapping.v1.DataSource
	5,  // 10: com.coralogix.global_mapping.v1.ExtractedLabel.is_custom_label:type_name -> google.protobuf.BoolValue
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_com_coralogix_global_mapping_v1_extracted_label_proto_init() }
func file_com_coralogix_global_mapping_v1_extracted_label_proto_init() {
	if File_com_coralogix_global_mapping_v1_extracted_label_proto != nil {
		return
	}
	file_com_coralogix_global_mapping_v1_data_source_type_proto_init()
	file_com_coralogix_global_mapping_v1_data_source_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_global_mapping_v1_extracted_label_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataSourceTypeValues); i {
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
		file_com_coralogix_global_mapping_v1_extracted_label_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtractedLabel); i {
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
			RawDescriptor: file_com_coralogix_global_mapping_v1_extracted_label_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_global_mapping_v1_extracted_label_proto_goTypes,
		DependencyIndexes: file_com_coralogix_global_mapping_v1_extracted_label_proto_depIdxs,
		MessageInfos:      file_com_coralogix_global_mapping_v1_extracted_label_proto_msgTypes,
	}.Build()
	File_com_coralogix_global_mapping_v1_extracted_label_proto = out.File
	file_com_coralogix_global_mapping_v1_extracted_label_proto_rawDesc = nil
	file_com_coralogix_global_mapping_v1_extracted_label_proto_goTypes = nil
	file_com_coralogix_global_mapping_v1_extracted_label_proto_depIdxs = nil
}
