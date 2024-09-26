// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.28.1
// source: com/coralogixapis/dashboards/v1/ast/dashboard.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

// Dashboard represents the structure and configuration of a Coralogix Custom Dashboard.
type Dashboard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier for the dashboard.
	Id *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Display name of the dashboard.
	Name *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Brief description or summary of the dashboard's purpose or content.
	Description *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Layout configuration for the dashboard's visual elements.
	Layout *Layout `protobuf:"bytes,4,opt,name=layout,proto3" json:"layout,omitempty"`
	// List of variables that can be used within the dashboard for dynamic content.
	Variables []*Variable `protobuf:"bytes,5,rep,name=variables,proto3" json:"variables,omitempty"`
	// List of filters that can be applied to the dashboard's data.
	Filters []*Filter `protobuf:"bytes,6,rep,name=filters,proto3" json:"filters,omitempty"`
	// Specifies the time frame for the dashboard's data. Can be either absolute or relative.
	//
	// Types that are assignable to TimeFrame:
	//
	//	*Dashboard_AbsoluteTimeFrame
	//	*Dashboard_RelativeTimeFrame
	TimeFrame isDashboard_TimeFrame `protobuf_oneof:"time_frame"`
	// polymorphic field for the dashboard's folder. We accept either a folder ID or a folder path
	//
	// Types that are assignable to Folder:
	//
	//	*Dashboard_FolderId
	//	*Dashboard_FolderPath
	Folder      isDashboard_Folder `protobuf_oneof:"folder"`
	Annotations []*Annotation      `protobuf:"bytes,11,rep,name=annotations,proto3" json:"annotations,omitempty"`
	// Specifies the auto refresh interval for the dashboard.
	//
	// Types that are assignable to AutoRefresh:
	//
	//	*Dashboard_Off
	//	*Dashboard_TwoMinutes
	//	*Dashboard_FiveMinutes
	AutoRefresh isDashboard_AutoRefresh `protobuf_oneof:"auto_refresh"`
}

func (x *Dashboard) Reset() {
	*x = Dashboard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dashboard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dashboard) ProtoMessage() {}

func (x *Dashboard) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dashboard.ProtoReflect.Descriptor instead.
func (*Dashboard) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_rawDescGZIP(), []int{0}
}

func (x *Dashboard) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Dashboard) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *Dashboard) GetDescription() *wrapperspb.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *Dashboard) GetLayout() *Layout {
	if x != nil {
		return x.Layout
	}
	return nil
}

func (x *Dashboard) GetVariables() []*Variable {
	if x != nil {
		return x.Variables
	}
	return nil
}

func (x *Dashboard) GetFilters() []*Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (m *Dashboard) GetTimeFrame() isDashboard_TimeFrame {
	if m != nil {
		return m.TimeFrame
	}
	return nil
}

func (x *Dashboard) GetAbsoluteTimeFrame() *TimeFrame {
	if x, ok := x.GetTimeFrame().(*Dashboard_AbsoluteTimeFrame); ok {
		return x.AbsoluteTimeFrame
	}
	return nil
}

func (x *Dashboard) GetRelativeTimeFrame() *durationpb.Duration {
	if x, ok := x.GetTimeFrame().(*Dashboard_RelativeTimeFrame); ok {
		return x.RelativeTimeFrame
	}
	return nil
}

func (m *Dashboard) GetFolder() isDashboard_Folder {
	if m != nil {
		return m.Folder
	}
	return nil
}

func (x *Dashboard) GetFolderId() *UUID {
	if x, ok := x.GetFolder().(*Dashboard_FolderId); ok {
		return x.FolderId
	}
	return nil
}

func (x *Dashboard) GetFolderPath() *FolderPath {
	if x, ok := x.GetFolder().(*Dashboard_FolderPath); ok {
		return x.FolderPath
	}
	return nil
}

func (x *Dashboard) GetAnnotations() []*Annotation {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (m *Dashboard) GetAutoRefresh() isDashboard_AutoRefresh {
	if m != nil {
		return m.AutoRefresh
	}
	return nil
}

func (x *Dashboard) GetOff() *Dashboard_AutoRefreshOff {
	if x, ok := x.GetAutoRefresh().(*Dashboard_Off); ok {
		return x.Off
	}
	return nil
}

func (x *Dashboard) GetTwoMinutes() *Dashboard_AutoRefreshTwoMinutes {
	if x, ok := x.GetAutoRefresh().(*Dashboard_TwoMinutes); ok {
		return x.TwoMinutes
	}
	return nil
}

func (x *Dashboard) GetFiveMinutes() *Dashboard_AutoRefreshFiveMinutes {
	if x, ok := x.GetAutoRefresh().(*Dashboard_FiveMinutes); ok {
		return x.FiveMinutes
	}
	return nil
}

type isDashboard_TimeFrame interface {
	isDashboard_TimeFrame()
}

type Dashboard_AbsoluteTimeFrame struct {
	// Absolute time frame specifying a fixed start and end time.
	AbsoluteTimeFrame *TimeFrame `protobuf:"bytes,7,opt,name=absolute_time_frame,json=absoluteTimeFrame,proto3,oneof"`
}

type Dashboard_RelativeTimeFrame struct {
	// Relative time frame specifying a duration from the current time.
	RelativeTimeFrame *durationpb.Duration `protobuf:"bytes,8,opt,name=relative_time_frame,json=relativeTimeFrame,proto3,oneof"`
}

func (*Dashboard_AbsoluteTimeFrame) isDashboard_TimeFrame() {}

func (*Dashboard_RelativeTimeFrame) isDashboard_TimeFrame() {}

type isDashboard_Folder interface {
	isDashboard_Folder()
}

type Dashboard_FolderId struct {
	FolderId *UUID `protobuf:"bytes,9,opt,name=folder_id,json=folderId,proto3,oneof"`
}

type Dashboard_FolderPath struct {
	FolderPath *FolderPath `protobuf:"bytes,10,opt,name=folder_path,json=folderPath,proto3,oneof"`
}

func (*Dashboard_FolderId) isDashboard_Folder() {}

func (*Dashboard_FolderPath) isDashboard_Folder() {}

type isDashboard_AutoRefresh interface {
	isDashboard_AutoRefresh()
}

type Dashboard_Off struct {
	Off *Dashboard_AutoRefreshOff `protobuf:"bytes,12,opt,name=off,proto3,oneof"`
}

type Dashboard_TwoMinutes struct {
	TwoMinutes *Dashboard_AutoRefreshTwoMinutes `protobuf:"bytes,13,opt,name=two_minutes,json=twoMinutes,proto3,oneof"`
}

type Dashboard_FiveMinutes struct {
	FiveMinutes *Dashboard_AutoRefreshFiveMinutes `protobuf:"bytes,14,opt,name=five_minutes,json=fiveMinutes,proto3,oneof"`
}

func (*Dashboard_Off) isDashboard_AutoRefresh() {}

func (*Dashboard_TwoMinutes) isDashboard_AutoRefresh() {}

func (*Dashboard_FiveMinutes) isDashboard_AutoRefresh() {}

type Dashboard_AutoRefreshOff struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Dashboard_AutoRefreshOff) Reset() {
	*x = Dashboard_AutoRefreshOff{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dashboard_AutoRefreshOff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dashboard_AutoRefreshOff) ProtoMessage() {}

func (x *Dashboard_AutoRefreshOff) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dashboard_AutoRefreshOff.ProtoReflect.Descriptor instead.
func (*Dashboard_AutoRefreshOff) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_rawDescGZIP(), []int{0, 0}
}

type Dashboard_AutoRefreshTwoMinutes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Dashboard_AutoRefreshTwoMinutes) Reset() {
	*x = Dashboard_AutoRefreshTwoMinutes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dashboard_AutoRefreshTwoMinutes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dashboard_AutoRefreshTwoMinutes) ProtoMessage() {}

func (x *Dashboard_AutoRefreshTwoMinutes) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dashboard_AutoRefreshTwoMinutes.ProtoReflect.Descriptor instead.
func (*Dashboard_AutoRefreshTwoMinutes) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_rawDescGZIP(), []int{0, 1}
}

type Dashboard_AutoRefreshFiveMinutes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Dashboard_AutoRefreshFiveMinutes) Reset() {
	*x = Dashboard_AutoRefreshFiveMinutes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dashboard_AutoRefreshFiveMinutes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dashboard_AutoRefreshFiveMinutes) ProtoMessage() {}

func (x *Dashboard_AutoRefreshFiveMinutes) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dashboard_AutoRefreshFiveMinutes.ProtoReflect.Descriptor instead.
func (*Dashboard_AutoRefreshFiveMinutes) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_rawDescGZIP(), []int{0, 2}
}

var File_com_coralogixapis_dashboards_v1_ast_dashboard_proto protoreflect.FileDescriptor

var file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_rawDesc = []byte{
	0x0a, 0x33, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x73, 0x74, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x74, 0x1a, 0x34, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x74, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x30, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x73, 0x74, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x35, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x74, 0x2f, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x74, 0x2f, 0x6c,
	0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x74,
	0x2f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x37, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x66, 0x72, 0x61,
	0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x09, 0x0a, 0x09, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x30, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x06, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74,
	0x52, 0x06, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x12, 0x4b, 0x0a, 0x09, 0x76, 0x61, 0x72, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73,
	0x74, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x09, 0x76, 0x61, 0x72, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x45, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x63, 0x0a, 0x13,
	0x61, 0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x66, 0x72,
	0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x48, 0x00, 0x52, 0x11,
	0x61, 0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x61, 0x6d,
	0x65, 0x12, 0x4b, 0x0a, 0x13, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x11, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x44,
	0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x48, 0x01, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x52, 0x0a, 0x0b, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x74, 0x2e,
	0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x48, 0x01, 0x52, 0x0a, 0x66, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x12, 0x51, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x61, 0x73, 0x74, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x51, 0x0a, 0x03, 0x6f,
	0x66, 0x66, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x44,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x4f, 0x66, 0x66, 0x48, 0x02, 0x52, 0x03, 0x6f, 0x66, 0x66, 0x12, 0x67,
	0x0a, 0x0b, 0x74, 0x77, 0x6f, 0x5f, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x77, 0x6f, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x48, 0x02, 0x52, 0x0a, 0x74, 0x77, 0x6f,
	0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x12, 0x6a, 0x0a, 0x0c, 0x66, 0x69, 0x76, 0x65, 0x5f,
	0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x45, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x61, 0x73, 0x74, 0x2e, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x41, 0x75,
	0x74, 0x6f, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x46, 0x69, 0x76, 0x65, 0x4d, 0x69, 0x6e,
	0x75, 0x74, 0x65, 0x73, 0x48, 0x02, 0x52, 0x0b, 0x66, 0x69, 0x76, 0x65, 0x4d, 0x69, 0x6e, 0x75,
	0x74, 0x65, 0x73, 0x1a, 0x10, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x4f, 0x66, 0x66, 0x1a, 0x17, 0x0a, 0x15, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x54, 0x77, 0x6f, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x1a, 0x18,
	0x0a, 0x16, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x46, 0x69, 0x76,
	0x65, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x42, 0x0c, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x42, 0x0e, 0x0a, 0x0c, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_rawDescOnce sync.Once
	file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_rawDescData = file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_rawDesc
)

func file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_rawDescData)
	})
	return file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_rawDescData
}

var file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_goTypes = []interface{}{
	(*Dashboard)(nil),                        // 0: com.coralogixapis.dashboards.v1.ast.Dashboard
	(*Dashboard_AutoRefreshOff)(nil),         // 1: com.coralogixapis.dashboards.v1.ast.Dashboard.AutoRefreshOff
	(*Dashboard_AutoRefreshTwoMinutes)(nil),  // 2: com.coralogixapis.dashboards.v1.ast.Dashboard.AutoRefreshTwoMinutes
	(*Dashboard_AutoRefreshFiveMinutes)(nil), // 3: com.coralogixapis.dashboards.v1.ast.Dashboard.AutoRefreshFiveMinutes
	(*wrapperspb.StringValue)(nil),           // 4: google.protobuf.StringValue
	(*Layout)(nil),                           // 5: com.coralogixapis.dashboards.v1.ast.Layout
	(*Variable)(nil),                         // 6: com.coralogixapis.dashboards.v1.ast.Variable
	(*Filter)(nil),                           // 7: com.coralogixapis.dashboards.v1.ast.Filter
	(*TimeFrame)(nil),                        // 8: com.coralogixapis.dashboards.v1.common.TimeFrame
	(*durationpb.Duration)(nil),              // 9: google.protobuf.Duration
	(*UUID)(nil),                             // 10: com.coralogixapis.dashboards.v1.UUID
	(*FolderPath)(nil),                       // 11: com.coralogixapis.dashboards.v1.ast.FolderPath
	(*Annotation)(nil),                       // 12: com.coralogixapis.dashboards.v1.ast.Annotation
}
var file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_depIdxs = []int32{
	4,  // 0: com.coralogixapis.dashboards.v1.ast.Dashboard.id:type_name -> google.protobuf.StringValue
	4,  // 1: com.coralogixapis.dashboards.v1.ast.Dashboard.name:type_name -> google.protobuf.StringValue
	4,  // 2: com.coralogixapis.dashboards.v1.ast.Dashboard.description:type_name -> google.protobuf.StringValue
	5,  // 3: com.coralogixapis.dashboards.v1.ast.Dashboard.layout:type_name -> com.coralogixapis.dashboards.v1.ast.Layout
	6,  // 4: com.coralogixapis.dashboards.v1.ast.Dashboard.variables:type_name -> com.coralogixapis.dashboards.v1.ast.Variable
	7,  // 5: com.coralogixapis.dashboards.v1.ast.Dashboard.filters:type_name -> com.coralogixapis.dashboards.v1.ast.Filter
	8,  // 6: com.coralogixapis.dashboards.v1.ast.Dashboard.absolute_time_frame:type_name -> com.coralogixapis.dashboards.v1.common.TimeFrame
	9,  // 7: com.coralogixapis.dashboards.v1.ast.Dashboard.relative_time_frame:type_name -> google.protobuf.Duration
	10, // 8: com.coralogixapis.dashboards.v1.ast.Dashboard.folder_id:type_name -> com.coralogixapis.dashboards.v1.UUID
	11, // 9: com.coralogixapis.dashboards.v1.ast.Dashboard.folder_path:type_name -> com.coralogixapis.dashboards.v1.ast.FolderPath
	12, // 10: com.coralogixapis.dashboards.v1.ast.Dashboard.annotations:type_name -> com.coralogixapis.dashboards.v1.ast.Annotation
	1,  // 11: com.coralogixapis.dashboards.v1.ast.Dashboard.off:type_name -> com.coralogixapis.dashboards.v1.ast.Dashboard.AutoRefreshOff
	2,  // 12: com.coralogixapis.dashboards.v1.ast.Dashboard.two_minutes:type_name -> com.coralogixapis.dashboards.v1.ast.Dashboard.AutoRefreshTwoMinutes
	3,  // 13: com.coralogixapis.dashboards.v1.ast.Dashboard.five_minutes:type_name -> com.coralogixapis.dashboards.v1.ast.Dashboard.AutoRefreshFiveMinutes
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_init() }
func file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_init() {
	if File_com_coralogixapis_dashboards_v1_ast_dashboard_proto != nil {
		return
	}
	file_com_coralogixapis_dashboards_v1_ast_annotation_proto_init()
	file_com_coralogixapis_dashboards_v1_ast_filter_proto_init()
	file_com_coralogixapis_dashboards_v1_ast_folder_path_proto_init()
	file_com_coralogixapis_dashboards_v1_ast_layout_proto_init()
	file_com_coralogixapis_dashboards_v1_ast_variable_proto_init()
	file_com_coralogixapis_dashboards_v1_common_time_frame_proto_init()
	file_com_coralogixapis_dashboards_v1_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dashboard); i {
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
		file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dashboard_AutoRefreshOff); i {
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
		file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dashboard_AutoRefreshTwoMinutes); i {
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
		file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dashboard_AutoRefreshFiveMinutes); i {
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
	file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Dashboard_AbsoluteTimeFrame)(nil),
		(*Dashboard_RelativeTimeFrame)(nil),
		(*Dashboard_FolderId)(nil),
		(*Dashboard_FolderPath)(nil),
		(*Dashboard_Off)(nil),
		(*Dashboard_TwoMinutes)(nil),
		(*Dashboard_FiveMinutes)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_dashboards_v1_ast_dashboard_proto = out.File
	file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_rawDesc = nil
	file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_goTypes = nil
	file_com_coralogixapis_dashboards_v1_ast_dashboard_proto_depIdxs = nil
}
