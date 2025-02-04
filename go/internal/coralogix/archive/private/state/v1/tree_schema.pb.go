// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: com/coralogix/archive/private/state/v1/tree_schema.proto

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

type TreeSchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*InnerNodeAndKey `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *TreeSchema) Reset() {
	*x = TreeSchema{}
	mi := &file_com_coralogix_archive_private_state_v1_tree_schema_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TreeSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TreeSchema) ProtoMessage() {}

func (x *TreeSchema) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_private_state_v1_tree_schema_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TreeSchema.ProtoReflect.Descriptor instead.
func (*TreeSchema) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_private_state_v1_tree_schema_proto_rawDescGZIP(), []int{0}
}

func (x *TreeSchema) GetNodes() []*InnerNodeAndKey {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type InnerNodeAndKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string     `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value *InnerNode `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *InnerNodeAndKey) Reset() {
	*x = InnerNodeAndKey{}
	mi := &file_com_coralogix_archive_private_state_v1_tree_schema_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InnerNodeAndKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InnerNodeAndKey) ProtoMessage() {}

func (x *InnerNodeAndKey) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_private_state_v1_tree_schema_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InnerNodeAndKey.ProtoReflect.Descriptor instead.
func (*InnerNodeAndKey) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_private_state_v1_tree_schema_proto_rawDescGZIP(), []int{1}
}

func (x *InnerNodeAndKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *InnerNodeAndKey) GetValue() *InnerNode {
	if x != nil {
		return x.Value
	}
	return nil
}

type InnerNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubtreeSize    int32              `protobuf:"varint,1,opt,name=subtree_size,json=subtreeSize,proto3" json:"subtree_size,omitempty"`
	MinOccurrences int32              `protobuf:"varint,2,opt,name=min_occurrences,json=minOccurrences,proto3" json:"min_occurrences,omitempty"`
	MaxOccurrences int32              `protobuf:"varint,3,opt,name=max_occurrences,json=maxOccurrences,proto3" json:"max_occurrences,omitempty"`
	InnerNodes     []*InnerNodeAndKey `protobuf:"bytes,4,rep,name=inner_nodes,json=innerNodes,proto3" json:"inner_nodes,omitempty"`
	Leafs          []*LeafAndKey      `protobuf:"bytes,5,rep,name=leafs,proto3" json:"leafs,omitempty"`
}

func (x *InnerNode) Reset() {
	*x = InnerNode{}
	mi := &file_com_coralogix_archive_private_state_v1_tree_schema_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InnerNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InnerNode) ProtoMessage() {}

func (x *InnerNode) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_private_state_v1_tree_schema_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InnerNode.ProtoReflect.Descriptor instead.
func (*InnerNode) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_private_state_v1_tree_schema_proto_rawDescGZIP(), []int{2}
}

func (x *InnerNode) GetSubtreeSize() int32 {
	if x != nil {
		return x.SubtreeSize
	}
	return 0
}

func (x *InnerNode) GetMinOccurrences() int32 {
	if x != nil {
		return x.MinOccurrences
	}
	return 0
}

func (x *InnerNode) GetMaxOccurrences() int32 {
	if x != nil {
		return x.MaxOccurrences
	}
	return 0
}

func (x *InnerNode) GetInnerNodes() []*InnerNodeAndKey {
	if x != nil {
		return x.InnerNodes
	}
	return nil
}

func (x *InnerNode) GetLeafs() []*LeafAndKey {
	if x != nil {
		return x.Leafs
	}
	return nil
}

type LeafAndKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value *Leaf  `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *LeafAndKey) Reset() {
	*x = LeafAndKey{}
	mi := &file_com_coralogix_archive_private_state_v1_tree_schema_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LeafAndKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeafAndKey) ProtoMessage() {}

func (x *LeafAndKey) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_private_state_v1_tree_schema_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeafAndKey.ProtoReflect.Descriptor instead.
func (*LeafAndKey) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_private_state_v1_tree_schema_proto_rawDescGZIP(), []int{3}
}

func (x *LeafAndKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *LeafAndKey) GetValue() *Leaf {
	if x != nil {
		return x.Value
	}
	return nil
}

type Leaf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldType   FieldType `protobuf:"varint,1,opt,name=field_type,json=fieldType,proto3,enum=com.coralogix.archive.private.state.v1.FieldType" json:"field_type,omitempty"`
	Occurrences int32     `protobuf:"varint,2,opt,name=occurrences,proto3" json:"occurrences,omitempty"`
	Required    bool      `protobuf:"varint,3,opt,name=required,proto3" json:"required,omitempty"`
}

func (x *Leaf) Reset() {
	*x = Leaf{}
	mi := &file_com_coralogix_archive_private_state_v1_tree_schema_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Leaf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Leaf) ProtoMessage() {}

func (x *Leaf) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_archive_private_state_v1_tree_schema_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Leaf.ProtoReflect.Descriptor instead.
func (*Leaf) Descriptor() ([]byte, []int) {
	return file_com_coralogix_archive_private_state_v1_tree_schema_proto_rawDescGZIP(), []int{4}
}

func (x *Leaf) GetFieldType() FieldType {
	if x != nil {
		return x.FieldType
	}
	return FieldType_FIELD_TYPE_UNSPECIFIED
}

func (x *Leaf) GetOccurrences() int32 {
	if x != nil {
		return x.Occurrences
	}
	return 0
}

func (x *Leaf) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

var File_com_coralogix_archive_private_state_v1_tree_schema_proto protoreflect.FileDescriptor

var file_com_coralogix_archive_private_state_v1_tree_schema_proto_rawDesc = []byte{
	0x0a, 0x38, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x65, 0x65, 0x5f, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x37, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x0a, 0x54,
	0x72, 0x65, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x4d, 0x0a, 0x05, 0x6e, 0x6f, 0x64,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x6e, 0x64, 0x4b, 0x65,
	0x79, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x6c, 0x0a, 0x0f, 0x49, 0x6e, 0x6e, 0x65,
	0x72, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x6e, 0x64, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x47, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xa4, 0x02, 0x0a, 0x09, 0x49, 0x6e, 0x6e, 0x65, 0x72,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x74, 0x72, 0x65, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x74,
	0x72, 0x65, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x69, 0x6e, 0x5f, 0x6f,
	0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0e, 0x6d, 0x69, 0x6e, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73,
	0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x5f, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6d, 0x61, 0x78, 0x4f, 0x63,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x58, 0x0a, 0x0b, 0x69, 0x6e, 0x6e,
	0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x4e, 0x6f, 0x64,
	0x65, 0x41, 0x6e, 0x64, 0x4b, 0x65, 0x79, 0x52, 0x0a, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x4e, 0x6f,
	0x64, 0x65, 0x73, 0x12, 0x48, 0x0a, 0x05, 0x6c, 0x65, 0x61, 0x66, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x65, 0x61, 0x66,
	0x41, 0x6e, 0x64, 0x4b, 0x65, 0x79, 0x52, 0x05, 0x6c, 0x65, 0x61, 0x66, 0x73, 0x22, 0x62, 0x0a,
	0x0a, 0x4c, 0x65, 0x61, 0x66, 0x41, 0x6e, 0x64, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x42, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x65, 0x61, 0x66, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x96, 0x01, 0x0a, 0x04, 0x4c, 0x65, 0x61, 0x66, 0x12, 0x50, 0x0a, 0x0a, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_com_coralogix_archive_private_state_v1_tree_schema_proto_rawDescOnce sync.Once
	file_com_coralogix_archive_private_state_v1_tree_schema_proto_rawDescData = file_com_coralogix_archive_private_state_v1_tree_schema_proto_rawDesc
)

func file_com_coralogix_archive_private_state_v1_tree_schema_proto_rawDescGZIP() []byte {
	file_com_coralogix_archive_private_state_v1_tree_schema_proto_rawDescOnce.Do(func() {
		file_com_coralogix_archive_private_state_v1_tree_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_archive_private_state_v1_tree_schema_proto_rawDescData)
	})
	return file_com_coralogix_archive_private_state_v1_tree_schema_proto_rawDescData
}

var file_com_coralogix_archive_private_state_v1_tree_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_com_coralogix_archive_private_state_v1_tree_schema_proto_goTypes = []any{
	(*TreeSchema)(nil),      // 0: com.coralogix.archive.private.state.v1.TreeSchema
	(*InnerNodeAndKey)(nil), // 1: com.coralogix.archive.private.state.v1.InnerNodeAndKey
	(*InnerNode)(nil),       // 2: com.coralogix.archive.private.state.v1.InnerNode
	(*LeafAndKey)(nil),      // 3: com.coralogix.archive.private.state.v1.LeafAndKey
	(*Leaf)(nil),            // 4: com.coralogix.archive.private.state.v1.Leaf
	(FieldType)(0),          // 5: com.coralogix.archive.private.state.v1.FieldType
}
var file_com_coralogix_archive_private_state_v1_tree_schema_proto_depIdxs = []int32{
	1, // 0: com.coralogix.archive.private.state.v1.TreeSchema.nodes:type_name -> com.coralogix.archive.private.state.v1.InnerNodeAndKey
	2, // 1: com.coralogix.archive.private.state.v1.InnerNodeAndKey.value:type_name -> com.coralogix.archive.private.state.v1.InnerNode
	1, // 2: com.coralogix.archive.private.state.v1.InnerNode.inner_nodes:type_name -> com.coralogix.archive.private.state.v1.InnerNodeAndKey
	3, // 3: com.coralogix.archive.private.state.v1.InnerNode.leafs:type_name -> com.coralogix.archive.private.state.v1.LeafAndKey
	4, // 4: com.coralogix.archive.private.state.v1.LeafAndKey.value:type_name -> com.coralogix.archive.private.state.v1.Leaf
	5, // 5: com.coralogix.archive.private.state.v1.Leaf.field_type:type_name -> com.coralogix.archive.private.state.v1.FieldType
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_com_coralogix_archive_private_state_v1_tree_schema_proto_init() }
func file_com_coralogix_archive_private_state_v1_tree_schema_proto_init() {
	if File_com_coralogix_archive_private_state_v1_tree_schema_proto != nil {
		return
	}
	file_com_coralogix_archive_private_state_v1_field_type_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogix_archive_private_state_v1_tree_schema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_archive_private_state_v1_tree_schema_proto_goTypes,
		DependencyIndexes: file_com_coralogix_archive_private_state_v1_tree_schema_proto_depIdxs,
		MessageInfos:      file_com_coralogix_archive_private_state_v1_tree_schema_proto_msgTypes,
	}.Build()
	File_com_coralogix_archive_private_state_v1_tree_schema_proto = out.File
	file_com_coralogix_archive_private_state_v1_tree_schema_proto_rawDesc = nil
	file_com_coralogix_archive_private_state_v1_tree_schema_proto_goTypes = nil
	file_com_coralogix_archive_private_state_v1_tree_schema_proto_depIdxs = nil
}
