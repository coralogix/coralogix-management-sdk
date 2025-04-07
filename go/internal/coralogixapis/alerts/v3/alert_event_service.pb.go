// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.26.1
// source: com/coralogixapis/alerts/v3/event/alert_event_service.proto

package v3

import (
	_ "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/common/v1"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
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

type OrderByAlertEventFields int32

const (
	OrderByAlertEventFields_ORDER_BY_ALERT_EVENT_FIELDS_UNSPECIFIED OrderByAlertEventFields = 0
	OrderByAlertEventFields_ORDER_BY_ALERT_EVENT_FIELDS_TIMESTAMP   OrderByAlertEventFields = 1
)

// Enum value maps for OrderByAlertEventFields.
var (
	OrderByAlertEventFields_name = map[int32]string{
		0: "ORDER_BY_ALERT_EVENT_FIELDS_UNSPECIFIED",
		1: "ORDER_BY_ALERT_EVENT_FIELDS_TIMESTAMP",
	}
	OrderByAlertEventFields_value = map[string]int32{
		"ORDER_BY_ALERT_EVENT_FIELDS_UNSPECIFIED": 0,
		"ORDER_BY_ALERT_EVENT_FIELDS_TIMESTAMP":   1,
	}
)

func (x OrderByAlertEventFields) Enum() *OrderByAlertEventFields {
	p := new(OrderByAlertEventFields)
	*p = x
	return p
}

func (x OrderByAlertEventFields) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderByAlertEventFields) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_enumTypes[0].Descriptor()
}

func (OrderByAlertEventFields) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_enumTypes[0]
}

func (x OrderByAlertEventFields) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderByAlertEventFields.Descriptor instead.
func (OrderByAlertEventFields) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_rawDescGZIP(), []int{0}
}

type OrderByAlertEventDirection int32

const (
	OrderByAlertEventDirection_ORDER_BY_ALERT_EVENT_DIRECTION_UNSPECIFIED OrderByAlertEventDirection = 0
	OrderByAlertEventDirection_ORDER_BY_ALERT_EVENT_DIRECTION_ASC         OrderByAlertEventDirection = 1
	OrderByAlertEventDirection_ORDER_BY_ALERT_EVENT_DIRECTION_DESC        OrderByAlertEventDirection = 2
)

// Enum value maps for OrderByAlertEventDirection.
var (
	OrderByAlertEventDirection_name = map[int32]string{
		0: "ORDER_BY_ALERT_EVENT_DIRECTION_UNSPECIFIED",
		1: "ORDER_BY_ALERT_EVENT_DIRECTION_ASC",
		2: "ORDER_BY_ALERT_EVENT_DIRECTION_DESC",
	}
	OrderByAlertEventDirection_value = map[string]int32{
		"ORDER_BY_ALERT_EVENT_DIRECTION_UNSPECIFIED": 0,
		"ORDER_BY_ALERT_EVENT_DIRECTION_ASC":         1,
		"ORDER_BY_ALERT_EVENT_DIRECTION_DESC":        2,
	}
)

func (x OrderByAlertEventDirection) Enum() *OrderByAlertEventDirection {
	p := new(OrderByAlertEventDirection)
	*p = x
	return p
}

func (x OrderByAlertEventDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderByAlertEventDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_enumTypes[1].Descriptor()
}

func (OrderByAlertEventDirection) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_enumTypes[1]
}

func (x OrderByAlertEventDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderByAlertEventDirection.Descriptor instead.
func (OrderByAlertEventDirection) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_rawDescGZIP(), []int{1}
}

type GetAlertEventStatsRequest struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Ids           []*wrapperspb.StringValue `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	OrderBys      []*AlertEventOrderBy      `protobuf:"bytes,2,rep,name=order_bys,json=orderBys,proto3" json:"order_bys,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAlertEventStatsRequest) Reset() {
	*x = GetAlertEventStatsRequest{}
	mi := &file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAlertEventStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlertEventStatsRequest) ProtoMessage() {}

func (x *GetAlertEventStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlertEventStatsRequest.ProtoReflect.Descriptor instead.
func (*GetAlertEventStatsRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetAlertEventStatsRequest) GetIds() []*wrapperspb.StringValue {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *GetAlertEventStatsRequest) GetOrderBys() []*AlertEventOrderBy {
	if x != nil {
		return x.OrderBys
	}
	return nil
}

type GetAlertEventStatsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EventsStats   map[string]*EventStats `protobuf:"bytes,1,rep,name=events_stats,json=eventsStats,proto3" json:"events_stats,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAlertEventStatsResponse) Reset() {
	*x = GetAlertEventStatsResponse{}
	mi := &file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAlertEventStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlertEventStatsResponse) ProtoMessage() {}

func (x *GetAlertEventStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlertEventStatsResponse.ProtoReflect.Descriptor instead.
func (*GetAlertEventStatsResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetAlertEventStatsResponse) GetEventsStats() map[string]*EventStats {
	if x != nil {
		return x.EventsStats
	}
	return nil
}

type Permutation struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	PermutationLabels map[string]string      `protobuf:"bytes,1,rep,name=permutation_labels,json=permutationLabels,proto3" json:"permutation_labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *Permutation) Reset() {
	*x = Permutation{}
	mi := &file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Permutation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permutation) ProtoMessage() {}

func (x *Permutation) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permutation.ProtoReflect.Descriptor instead.
func (*Permutation) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_rawDescGZIP(), []int{2}
}

func (x *Permutation) GetPermutationLabels() map[string]string {
	if x != nil {
		return x.PermutationLabels
	}
	return nil
}

type EventStats struct {
	state                        protoimpl.MessageState  `protogen:"open.v1"`
	Count                        *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=count,proto3" json:"count,omitempty"`
	ResolvedCount                *wrapperspb.UInt32Value `protobuf:"bytes,2,opt,name=resolved_count,json=resolvedCount,proto3" json:"resolved_count,omitempty"`
	TriggeredCount               *wrapperspb.UInt32Value `protobuf:"bytes,3,opt,name=triggered_count,json=triggeredCount,proto3" json:"triggered_count,omitempty"`
	TriggeredPermutationsSamples []*Permutation          `protobuf:"bytes,4,rep,name=triggered_permutations_samples,json=triggeredPermutationsSamples,proto3" json:"triggered_permutations_samples,omitempty"`
	ResolvedPermutationsSamples  []*Permutation          `protobuf:"bytes,5,rep,name=resolved_permutations_samples,json=resolvedPermutationsSamples,proto3" json:"resolved_permutations_samples,omitempty"`
	ActivityAnalysisStats        *ActivityAnalysisStats  `protobuf:"bytes,6,opt,name=activity_analysis_stats,json=activityAnalysisStats,proto3" json:"activity_analysis_stats,omitempty"`
	unknownFields                protoimpl.UnknownFields
	sizeCache                    protoimpl.SizeCache
}

func (x *EventStats) Reset() {
	*x = EventStats{}
	mi := &file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventStats) ProtoMessage() {}

func (x *EventStats) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventStats.ProtoReflect.Descriptor instead.
func (*EventStats) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_rawDescGZIP(), []int{3}
}

func (x *EventStats) GetCount() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Count
	}
	return nil
}

func (x *EventStats) GetResolvedCount() *wrapperspb.UInt32Value {
	if x != nil {
		return x.ResolvedCount
	}
	return nil
}

func (x *EventStats) GetTriggeredCount() *wrapperspb.UInt32Value {
	if x != nil {
		return x.TriggeredCount
	}
	return nil
}

func (x *EventStats) GetTriggeredPermutationsSamples() []*Permutation {
	if x != nil {
		return x.TriggeredPermutationsSamples
	}
	return nil
}

func (x *EventStats) GetResolvedPermutationsSamples() []*Permutation {
	if x != nil {
		return x.ResolvedPermutationsSamples
	}
	return nil
}

func (x *EventStats) GetActivityAnalysisStats() *ActivityAnalysisStats {
	if x != nil {
		return x.ActivityAnalysisStats
	}
	return nil
}

type ActivityAnalysisStats struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	IsMutedCount  *wrapperspb.UInt32Value   `protobuf:"bytes,1,opt,name=is_muted_count,json=isMutedCount,proto3" json:"is_muted_count,omitempty"`
	Rules         []*wrapperspb.StringValue `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ActivityAnalysisStats) Reset() {
	*x = ActivityAnalysisStats{}
	mi := &file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActivityAnalysisStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityAnalysisStats) ProtoMessage() {}

func (x *ActivityAnalysisStats) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityAnalysisStats.ProtoReflect.Descriptor instead.
func (*ActivityAnalysisStats) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_rawDescGZIP(), []int{4}
}

func (x *ActivityAnalysisStats) GetIsMutedCount() *wrapperspb.UInt32Value {
	if x != nil {
		return x.IsMutedCount
	}
	return nil
}

func (x *ActivityAnalysisStats) GetRules() []*wrapperspb.StringValue {
	if x != nil {
		return x.Rules
	}
	return nil
}

type GetAlertEventRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Alert event ID
	Id            *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderBys      []*AlertEventOrderBy    `protobuf:"bytes,2,rep,name=order_bys,json=orderBys,proto3" json:"order_bys,omitempty"`
	Pagination    *PaginationRequest      `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAlertEventRequest) Reset() {
	*x = GetAlertEventRequest{}
	mi := &file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAlertEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlertEventRequest) ProtoMessage() {}

func (x *GetAlertEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlertEventRequest.ProtoReflect.Descriptor instead.
func (*GetAlertEventRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetAlertEventRequest) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *GetAlertEventRequest) GetOrderBys() []*AlertEventOrderBy {
	if x != nil {
		return x.OrderBys
	}
	return nil
}

func (x *GetAlertEventRequest) GetPagination() *PaginationRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type GetAlertEventResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Alert event ID
	Id *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to AlertEvent:
	//
	//	*GetAlertEventResponse_SinglePermutation
	//	*GetAlertEventResponse_MultiplePermutation
	AlertEvent    isGetAlertEventResponse_AlertEvent `protobuf_oneof:"alert_event"`
	Pagination    *PaginationResponse                `protobuf:"bytes,4,opt,name=pagination,proto3" json:"pagination,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAlertEventResponse) Reset() {
	*x = GetAlertEventResponse{}
	mi := &file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAlertEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlertEventResponse) ProtoMessage() {}

func (x *GetAlertEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlertEventResponse.ProtoReflect.Descriptor instead.
func (*GetAlertEventResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetAlertEventResponse) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *GetAlertEventResponse) GetAlertEvent() isGetAlertEventResponse_AlertEvent {
	if x != nil {
		return x.AlertEvent
	}
	return nil
}

func (x *GetAlertEventResponse) GetSinglePermutation() *AlertEvent {
	if x != nil {
		if x, ok := x.AlertEvent.(*GetAlertEventResponse_SinglePermutation); ok {
			return x.SinglePermutation
		}
	}
	return nil
}

func (x *GetAlertEventResponse) GetMultiplePermutation() *AlertEventMultiplePermutation {
	if x != nil {
		if x, ok := x.AlertEvent.(*GetAlertEventResponse_MultiplePermutation); ok {
			return x.MultiplePermutation
		}
	}
	return nil
}

func (x *GetAlertEventResponse) GetPagination() *PaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type isGetAlertEventResponse_AlertEvent interface {
	isGetAlertEventResponse_AlertEvent()
}

type GetAlertEventResponse_SinglePermutation struct {
	SinglePermutation *AlertEvent `protobuf:"bytes,2,opt,name=single_permutation,json=singlePermutation,proto3,oneof"`
}

type GetAlertEventResponse_MultiplePermutation struct {
	MultiplePermutation *AlertEventMultiplePermutation `protobuf:"bytes,3,opt,name=multiple_permutation,json=multiplePermutation,proto3,oneof"`
}

func (*GetAlertEventResponse_SinglePermutation) isGetAlertEventResponse_AlertEvent() {}

func (*GetAlertEventResponse_MultiplePermutation) isGetAlertEventResponse_AlertEvent() {}

type PaginationRequest struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	PageSize      *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken     *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaginationRequest) Reset() {
	*x = PaginationRequest{}
	mi := &file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaginationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationRequest) ProtoMessage() {}

func (x *PaginationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationRequest.ProtoReflect.Descriptor instead.
func (*PaginationRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_rawDescGZIP(), []int{7}
}

func (x *PaginationRequest) GetPageSize() *wrapperspb.UInt32Value {
	if x != nil {
		return x.PageSize
	}
	return nil
}

func (x *PaginationRequest) GetPageToken() *wrapperspb.StringValue {
	if x != nil {
		return x.PageToken
	}
	return nil
}

type PaginationResponse struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	TotalSize     *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	NextPageToken *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaginationResponse) Reset() {
	*x = PaginationResponse{}
	mi := &file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationResponse) ProtoMessage() {}

func (x *PaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationResponse.ProtoReflect.Descriptor instead.
func (*PaginationResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_rawDescGZIP(), []int{8}
}

func (x *PaginationResponse) GetTotalSize() *wrapperspb.UInt32Value {
	if x != nil {
		return x.TotalSize
	}
	return nil
}

func (x *PaginationResponse) GetNextPageToken() *wrapperspb.StringValue {
	if x != nil {
		return x.NextPageToken
	}
	return nil
}

type AlertEventMultiplePermutation struct {
	state                         protoimpl.MessageState `protogen:"open.v1"`
	AlertEventMultiplePermutation []*AlertEvent          `protobuf:"bytes,1,rep,name=alert_event_multiple_permutation,json=alertEventMultiplePermutation,proto3" json:"alert_event_multiple_permutation,omitempty"`
	unknownFields                 protoimpl.UnknownFields
	sizeCache                     protoimpl.SizeCache
}

func (x *AlertEventMultiplePermutation) Reset() {
	*x = AlertEventMultiplePermutation{}
	mi := &file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlertEventMultiplePermutation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertEventMultiplePermutation) ProtoMessage() {}

func (x *AlertEventMultiplePermutation) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertEventMultiplePermutation.ProtoReflect.Descriptor instead.
func (*AlertEventMultiplePermutation) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_rawDescGZIP(), []int{9}
}

func (x *AlertEventMultiplePermutation) GetAlertEventMultiplePermutation() []*AlertEvent {
	if x != nil {
		return x.AlertEventMultiplePermutation
	}
	return nil
}

type AlertEventOrderBy struct {
	state         protoimpl.MessageState     `protogen:"open.v1"`
	FieldName     OrderByAlertEventFields    `protobuf:"varint,1,opt,name=field_name,json=fieldName,proto3,enum=com.coralogixapis.alerts.v3.OrderByAlertEventFields" json:"field_name,omitempty"`
	Direction     OrderByAlertEventDirection `protobuf:"varint,2,opt,name=direction,proto3,enum=com.coralogixapis.alerts.v3.OrderByAlertEventDirection" json:"direction,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AlertEventOrderBy) Reset() {
	*x = AlertEventOrderBy{}
	mi := &file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlertEventOrderBy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertEventOrderBy) ProtoMessage() {}

func (x *AlertEventOrderBy) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertEventOrderBy.ProtoReflect.Descriptor instead.
func (*AlertEventOrderBy) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_rawDescGZIP(), []int{10}
}

func (x *AlertEventOrderBy) GetFieldName() OrderByAlertEventFields {
	if x != nil {
		return x.FieldName
	}
	return OrderByAlertEventFields_ORDER_BY_ALERT_EVENT_FIELDS_UNSPECIFIED
}

func (x *AlertEventOrderBy) GetDirection() OrderByAlertEventDirection {
	if x != nil {
		return x.Direction
	}
	return OrderByAlertEventDirection_ORDER_BY_ALERT_EVENT_DIRECTION_UNSPECIFIED
}

var File_com_coralogixapis_alerts_v3_event_alert_event_service_proto protoreflect.FileDescriptor

const file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_rawDesc = "" +
	"\n" +
	";com/coralogixapis/alerts/v3/event/alert_event_service.proto\x12\x1bcom.coralogixapis.alerts.v3\x1a google/protobuf/descriptor.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x1egoogle/protobuf/wrappers.proto\x1a3com/coralogixapis/alerts/v3/event/alert_event.proto\x1a'com/coralogix/common/v1/audit_log.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\xc3\x01\n" +
	"\x19GetAlertEventStatsRequest\x12.\n" +
	"\x03ids\x18\x01 \x03(\v2\x1c.google.protobuf.StringValueR\x03ids\x12K\n" +
	"\torder_bys\x18\x02 \x03(\v2..com.coralogixapis.alerts.v3.AlertEventOrderByR\borderBys:)\x92A&\n" +
	"$*\"Get alert event statistics request\"\x9e\x02\n" +
	"\x1aGetAlertEventStatsResponse\x12k\n" +
	"\fevents_stats\x18\x01 \x03(\v2H.com.coralogixapis.alerts.v3.GetAlertEventStatsResponse.EventsStatsEntryR\veventsStats\x1ag\n" +
	"\x10EventsStatsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12=\n" +
	"\x05value\x18\x02 \x01(\v2'.com.coralogixapis.alerts.v3.EventStatsR\x05value:\x028\x01:*\x92A'\n" +
	"%*#Get alert event statistics response\"\xe6\x01\n" +
	"\vPermutation\x12n\n" +
	"\x12permutation_labels\x18\x01 \x03(\v2?.com.coralogixapis.alerts.v3.Permutation.PermutationLabelsEntryR\x11permutationLabels\x1aD\n" +
	"\x16PermutationLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01:!\x92A\x1e\n" +
	"\x1c*\x1aPermutation data structure\"\xc4\x04\n" +
	"\n" +
	"EventStats\x122\n" +
	"\x05count\x18\x01 \x01(\v2\x1c.google.protobuf.UInt32ValueR\x05count\x12C\n" +
	"\x0eresolved_count\x18\x02 \x01(\v2\x1c.google.protobuf.UInt32ValueR\rresolvedCount\x12E\n" +
	"\x0ftriggered_count\x18\x03 \x01(\v2\x1c.google.protobuf.UInt32ValueR\x0etriggeredCount\x12n\n" +
	"\x1etriggered_permutations_samples\x18\x04 \x03(\v2(.com.coralogixapis.alerts.v3.PermutationR\x1ctriggeredPermutationsSamples\x12l\n" +
	"\x1dresolved_permutations_samples\x18\x05 \x03(\v2(.com.coralogixapis.alerts.v3.PermutationR\x1bresolvedPermutationsSamples\x12j\n" +
	"\x17activity_analysis_stats\x18\x06 \x01(\v22.com.coralogixapis.alerts.v3.ActivityAnalysisStatsR\x15activityAnalysisStats:,\x92A)\n" +
	"'*%Alert event statistics data structure\"\xba\x01\n" +
	"\x15ActivityAnalysisStats\x12B\n" +
	"\x0eis_muted_count\x18\x01 \x01(\v2\x1c.google.protobuf.UInt32ValueR\fisMutedCount\x122\n" +
	"\x05rules\x18\x02 \x03(\v2\x1c.google.protobuf.StringValueR\x05rules:)\x92A&\n" +
	"$*\"Event activity analysis statistics\"\x87\x02\n" +
	"\x14GetAlertEventRequest\x12,\n" +
	"\x02id\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueR\x02id\x12K\n" +
	"\torder_bys\x18\x02 \x03(\v2..com.coralogixapis.alerts.v3.AlertEventOrderByR\borderBys\x12N\n" +
	"\n" +
	"pagination\x18\x03 \x01(\v2..com.coralogixapis.alerts.v3.PaginationRequestR\n" +
	"pagination:$\x92A!\n" +
	"\x1f*\x1dGet alert event by ID request\"\x91\x03\n" +
	"\x15GetAlertEventResponse\x12,\n" +
	"\x02id\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueR\x02id\x12X\n" +
	"\x12single_permutation\x18\x02 \x01(\v2'.com.coralogixapis.alerts.v3.AlertEventH\x00R\x11singlePermutation\x12o\n" +
	"\x14multiple_permutation\x18\x03 \x01(\v2:.com.coralogixapis.alerts.v3.AlertEventMultiplePermutationH\x00R\x13multiplePermutation\x12O\n" +
	"\n" +
	"pagination\x18\x04 \x01(\v2/.com.coralogixapis.alerts.v3.PaginationResponseR\n" +
	"pagination:\x1f\x92A\x1c\n" +
	"\x1a*\x18Get alert event responseB\r\n" +
	"\valert_event\"\x8b\x01\n" +
	"\x11PaginationRequest\x129\n" +
	"\tpage_size\x18\x01 \x01(\v2\x1c.google.protobuf.UInt32ValueR\bpageSize\x12;\n" +
	"\n" +
	"page_token\x18\x02 \x01(\v2\x1c.google.protobuf.StringValueR\tpageToken\"\x97\x01\n" +
	"\x12PaginationResponse\x12;\n" +
	"\n" +
	"total_size\x18\x01 \x01(\v2\x1c.google.protobuf.UInt32ValueR\ttotalSize\x12D\n" +
	"\x0fnext_page_token\x18\x02 \x01(\v2\x1c.google.protobuf.StringValueR\rnextPageToken\"\x91\x01\n" +
	"\x1dAlertEventMultiplePermutation\x12p\n" +
	" alert_event_multiple_permutation\x18\x01 \x03(\v2'.com.coralogixapis.alerts.v3.AlertEventR\x1dalertEventMultiplePermutation\"\xbf\x01\n" +
	"\x11AlertEventOrderBy\x12S\n" +
	"\n" +
	"field_name\x18\x01 \x01(\x0e24.com.coralogixapis.alerts.v3.OrderByAlertEventFieldsR\tfieldName\x12U\n" +
	"\tdirection\x18\x02 \x01(\x0e27.com.coralogixapis.alerts.v3.OrderByAlertEventDirectionR\tdirection*q\n" +
	"\x17OrderByAlertEventFields\x12+\n" +
	"'ORDER_BY_ALERT_EVENT_FIELDS_UNSPECIFIED\x10\x00\x12)\n" +
	"%ORDER_BY_ALERT_EVENT_FIELDS_TIMESTAMP\x10\x01*\x9d\x01\n" +
	"\x1aOrderByAlertEventDirection\x12.\n" +
	"*ORDER_BY_ALERT_EVENT_DIRECTION_UNSPECIFIED\x10\x00\x12&\n" +
	"\"ORDER_BY_ALERT_EVENT_DIRECTION_ASC\x10\x01\x12'\n" +
	"#ORDER_BY_ALERT_EVENT_DIRECTION_DESC\x10\x022\xef\a\n" +
	"\x11AlertEventService\x12\xe3\x02\n" +
	"\rGetAlertEvent\x121.com.coralogixapis.alerts.v3.GetAlertEventRequest\x1a2.com.coralogixapis.alerts.v3.GetAlertEventResponse\"\xea\x01\x92A\xb5\x01\n" +
	"\x14Alert events service*\x15Get alert event by IDJ\x14\n" +
	"\x03400\x12\r\n" +
	"\vBad RequestJ\x1d\n" +
	"\x03401\x12\x16\n" +
	"\x14Unauthorized requestJ\x1e\n" +
	"\x03500\x12\x17\n" +
	"\x15Internal server errorj1\n" +
	"\x16x-coralogixPermissions\x12\x172\x15\n" +
	"\x13\x1a\x11alerts:ReadConfig¸\x02\x11\n" +
	"\x0fget alert event\x82\xd3\xe4\x93\x02\x16\x12\x14/v3/alert-event/{id}\x12\x81\x03\n" +
	"\x13GetAlertEventsStats\x126.com.coralogixapis.alerts.v3.GetAlertEventStatsRequest\x1a7.com.coralogixapis.alerts.v3.GetAlertEventStatsResponse\"\xf8\x01\x92A\xbb\x01\n" +
	"\x14Alert events service*\x1bGet alert events statisticsJ\x14\n" +
	"\x03400\x12\r\n" +
	"\vBad RequestJ\x1d\n" +
	"\x03401\x12\x16\n" +
	"\x14Unauthorized requestJ\x1e\n" +
	"\x03500\x12\x17\n" +
	"\x15Internal server errorj1\n" +
	"\x16x-coralogixPermissions\x12\x172\x15\n" +
	"\x13\x1a\x11alerts:ReadConfig¸\x02\x18\n" +
	"\x16get alert events stats\x82\xd3\xe4\x93\x02\x17\x12\x15/v3/alert-event-stats\x1a\xef\x01\x92A\xeb\x01\n" +
	"\x14Alert events service\x12RGet information regarding your alert events - instances of alerts being triggered.\x1a\x7f\n" +
	"ALearn more about alert events and incidents in our documentation.\x12:https://coralogix.com/docs/user-guides/alerting/incidents/b\x06proto3"

var (
	file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_rawDescData []byte
)

func file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_rawDesc)))
	})
	return file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_rawDescData
}

var file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_goTypes = []any{
	(OrderByAlertEventFields)(0),          // 0: com.coralogixapis.alerts.v3.OrderByAlertEventFields
	(OrderByAlertEventDirection)(0),       // 1: com.coralogixapis.alerts.v3.OrderByAlertEventDirection
	(*GetAlertEventStatsRequest)(nil),     // 2: com.coralogixapis.alerts.v3.GetAlertEventStatsRequest
	(*GetAlertEventStatsResponse)(nil),    // 3: com.coralogixapis.alerts.v3.GetAlertEventStatsResponse
	(*Permutation)(nil),                   // 4: com.coralogixapis.alerts.v3.Permutation
	(*EventStats)(nil),                    // 5: com.coralogixapis.alerts.v3.EventStats
	(*ActivityAnalysisStats)(nil),         // 6: com.coralogixapis.alerts.v3.ActivityAnalysisStats
	(*GetAlertEventRequest)(nil),          // 7: com.coralogixapis.alerts.v3.GetAlertEventRequest
	(*GetAlertEventResponse)(nil),         // 8: com.coralogixapis.alerts.v3.GetAlertEventResponse
	(*PaginationRequest)(nil),             // 9: com.coralogixapis.alerts.v3.PaginationRequest
	(*PaginationResponse)(nil),            // 10: com.coralogixapis.alerts.v3.PaginationResponse
	(*AlertEventMultiplePermutation)(nil), // 11: com.coralogixapis.alerts.v3.AlertEventMultiplePermutation
	(*AlertEventOrderBy)(nil),             // 12: com.coralogixapis.alerts.v3.AlertEventOrderBy
	nil,                                   // 13: com.coralogixapis.alerts.v3.GetAlertEventStatsResponse.EventsStatsEntry
	nil,                                   // 14: com.coralogixapis.alerts.v3.Permutation.PermutationLabelsEntry
	(*wrapperspb.StringValue)(nil),        // 15: google.protobuf.StringValue
	(*wrapperspb.UInt32Value)(nil),        // 16: google.protobuf.UInt32Value
	(*AlertEvent)(nil),                    // 17: com.coralogixapis.alerts.v3.AlertEvent
}
var file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_depIdxs = []int32{
	15, // 0: com.coralogixapis.alerts.v3.GetAlertEventStatsRequest.ids:type_name -> google.protobuf.StringValue
	12, // 1: com.coralogixapis.alerts.v3.GetAlertEventStatsRequest.order_bys:type_name -> com.coralogixapis.alerts.v3.AlertEventOrderBy
	13, // 2: com.coralogixapis.alerts.v3.GetAlertEventStatsResponse.events_stats:type_name -> com.coralogixapis.alerts.v3.GetAlertEventStatsResponse.EventsStatsEntry
	14, // 3: com.coralogixapis.alerts.v3.Permutation.permutation_labels:type_name -> com.coralogixapis.alerts.v3.Permutation.PermutationLabelsEntry
	16, // 4: com.coralogixapis.alerts.v3.EventStats.count:type_name -> google.protobuf.UInt32Value
	16, // 5: com.coralogixapis.alerts.v3.EventStats.resolved_count:type_name -> google.protobuf.UInt32Value
	16, // 6: com.coralogixapis.alerts.v3.EventStats.triggered_count:type_name -> google.protobuf.UInt32Value
	4,  // 7: com.coralogixapis.alerts.v3.EventStats.triggered_permutations_samples:type_name -> com.coralogixapis.alerts.v3.Permutation
	4,  // 8: com.coralogixapis.alerts.v3.EventStats.resolved_permutations_samples:type_name -> com.coralogixapis.alerts.v3.Permutation
	6,  // 9: com.coralogixapis.alerts.v3.EventStats.activity_analysis_stats:type_name -> com.coralogixapis.alerts.v3.ActivityAnalysisStats
	16, // 10: com.coralogixapis.alerts.v3.ActivityAnalysisStats.is_muted_count:type_name -> google.protobuf.UInt32Value
	15, // 11: com.coralogixapis.alerts.v3.ActivityAnalysisStats.rules:type_name -> google.protobuf.StringValue
	15, // 12: com.coralogixapis.alerts.v3.GetAlertEventRequest.id:type_name -> google.protobuf.StringValue
	12, // 13: com.coralogixapis.alerts.v3.GetAlertEventRequest.order_bys:type_name -> com.coralogixapis.alerts.v3.AlertEventOrderBy
	9,  // 14: com.coralogixapis.alerts.v3.GetAlertEventRequest.pagination:type_name -> com.coralogixapis.alerts.v3.PaginationRequest
	15, // 15: com.coralogixapis.alerts.v3.GetAlertEventResponse.id:type_name -> google.protobuf.StringValue
	17, // 16: com.coralogixapis.alerts.v3.GetAlertEventResponse.single_permutation:type_name -> com.coralogixapis.alerts.v3.AlertEvent
	11, // 17: com.coralogixapis.alerts.v3.GetAlertEventResponse.multiple_permutation:type_name -> com.coralogixapis.alerts.v3.AlertEventMultiplePermutation
	10, // 18: com.coralogixapis.alerts.v3.GetAlertEventResponse.pagination:type_name -> com.coralogixapis.alerts.v3.PaginationResponse
	16, // 19: com.coralogixapis.alerts.v3.PaginationRequest.page_size:type_name -> google.protobuf.UInt32Value
	15, // 20: com.coralogixapis.alerts.v3.PaginationRequest.page_token:type_name -> google.protobuf.StringValue
	16, // 21: com.coralogixapis.alerts.v3.PaginationResponse.total_size:type_name -> google.protobuf.UInt32Value
	15, // 22: com.coralogixapis.alerts.v3.PaginationResponse.next_page_token:type_name -> google.protobuf.StringValue
	17, // 23: com.coralogixapis.alerts.v3.AlertEventMultiplePermutation.alert_event_multiple_permutation:type_name -> com.coralogixapis.alerts.v3.AlertEvent
	0,  // 24: com.coralogixapis.alerts.v3.AlertEventOrderBy.field_name:type_name -> com.coralogixapis.alerts.v3.OrderByAlertEventFields
	1,  // 25: com.coralogixapis.alerts.v3.AlertEventOrderBy.direction:type_name -> com.coralogixapis.alerts.v3.OrderByAlertEventDirection
	5,  // 26: com.coralogixapis.alerts.v3.GetAlertEventStatsResponse.EventsStatsEntry.value:type_name -> com.coralogixapis.alerts.v3.EventStats
	7,  // 27: com.coralogixapis.alerts.v3.AlertEventService.GetAlertEvent:input_type -> com.coralogixapis.alerts.v3.GetAlertEventRequest
	2,  // 28: com.coralogixapis.alerts.v3.AlertEventService.GetAlertEventsStats:input_type -> com.coralogixapis.alerts.v3.GetAlertEventStatsRequest
	8,  // 29: com.coralogixapis.alerts.v3.AlertEventService.GetAlertEvent:output_type -> com.coralogixapis.alerts.v3.GetAlertEventResponse
	3,  // 30: com.coralogixapis.alerts.v3.AlertEventService.GetAlertEventsStats:output_type -> com.coralogixapis.alerts.v3.GetAlertEventStatsResponse
	29, // [29:31] is the sub-list for method output_type
	27, // [27:29] is the sub-list for method input_type
	27, // [27:27] is the sub-list for extension type_name
	27, // [27:27] is the sub-list for extension extendee
	0,  // [0:27] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_init() }
func file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_init() {
	if File_com_coralogixapis_alerts_v3_event_alert_event_service_proto != nil {
		return
	}
	file_com_coralogixapis_alerts_v3_event_alert_event_proto_init()
	file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_msgTypes[6].OneofWrappers = []any{
		(*GetAlertEventResponse_SinglePermutation)(nil),
		(*GetAlertEventResponse_MultiplePermutation)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_rawDesc), len(file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_enumTypes,
		MessageInfos:      file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerts_v3_event_alert_event_service_proto = out.File
	file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_goTypes = nil
	file_com_coralogixapis_alerts_v3_event_alert_event_service_proto_depIdxs = nil
}
