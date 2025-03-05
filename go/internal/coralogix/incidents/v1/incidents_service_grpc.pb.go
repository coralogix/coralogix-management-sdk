// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: com/coralogixapis/incidents/v1/incidents_service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	IncidentsService_GetIncident_FullMethodName                    = "/com.coralogixapis.incidents.v1.IncidentsService/GetIncident"
	IncidentsService_BatchGetIncident_FullMethodName               = "/com.coralogixapis.incidents.v1.IncidentsService/BatchGetIncident"
	IncidentsService_ListIncidents_FullMethodName                  = "/com.coralogixapis.incidents.v1.IncidentsService/ListIncidents"
	IncidentsService_ListIncidentAggregations_FullMethodName       = "/com.coralogixapis.incidents.v1.IncidentsService/ListIncidentAggregations"
	IncidentsService_GetFilterValues_FullMethodName                = "/com.coralogixapis.incidents.v1.IncidentsService/GetFilterValues"
	IncidentsService_AssignIncidents_FullMethodName                = "/com.coralogixapis.incidents.v1.IncidentsService/AssignIncidents"
	IncidentsService_UnassignIncidents_FullMethodName              = "/com.coralogixapis.incidents.v1.IncidentsService/UnassignIncidents"
	IncidentsService_AcknowledgeIncidents_FullMethodName           = "/com.coralogixapis.incidents.v1.IncidentsService/AcknowledgeIncidents"
	IncidentsService_CloseIncidents_FullMethodName                 = "/com.coralogixapis.incidents.v1.IncidentsService/CloseIncidents"
	IncidentsService_GetIncidentEvents_FullMethodName              = "/com.coralogixapis.incidents.v1.IncidentsService/GetIncidentEvents"
	IncidentsService_ResolveIncidents_FullMethodName               = "/com.coralogixapis.incidents.v1.IncidentsService/ResolveIncidents"
	IncidentsService_GetIncidentUsingCorrelationKey_FullMethodName = "/com.coralogixapis.incidents.v1.IncidentsService/GetIncidentUsingCorrelationKey"
	IncidentsService_ListIncidentEvents_FullMethodName             = "/com.coralogixapis.incidents.v1.IncidentsService/ListIncidentEvents"
	IncidentsService_ListIncidentEventsTotalCount_FullMethodName   = "/com.coralogixapis.incidents.v1.IncidentsService/ListIncidentEventsTotalCount"
	IncidentsService_ListIncidentEventsFilterValues_FullMethodName = "/com.coralogixapis.incidents.v1.IncidentsService/ListIncidentEventsFilterValues"
)

// IncidentsServiceClient is the client API for IncidentsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IncidentsServiceClient interface {
	GetIncident(ctx context.Context, in *GetIncidentRequest, opts ...grpc.CallOption) (*GetIncidentResponse, error)
	BatchGetIncident(ctx context.Context, in *BatchGetIncidentRequest, opts ...grpc.CallOption) (*BatchGetIncidentResponse, error)
	ListIncidents(ctx context.Context, in *ListIncidentsRequest, opts ...grpc.CallOption) (*ListIncidentsResponse, error)
	ListIncidentAggregations(ctx context.Context, in *ListIncidentAggregationsRequest, opts ...grpc.CallOption) (*ListIncidentAggregationsResponse, error)
	GetFilterValues(ctx context.Context, in *GetFilterValuesRequest, opts ...grpc.CallOption) (*GetFilterValuesResponse, error)
	AssignIncidents(ctx context.Context, in *AssignIncidentsRequest, opts ...grpc.CallOption) (*AssignIncidentsResponse, error)
	UnassignIncidents(ctx context.Context, in *UnassignIncidentsRequest, opts ...grpc.CallOption) (*UnassignIncidentsResponse, error)
	AcknowledgeIncidents(ctx context.Context, in *AcknowledgeIncidentsRequest, opts ...grpc.CallOption) (*AcknowledgeIncidentsResponse, error)
	CloseIncidents(ctx context.Context, in *CloseIncidentsRequest, opts ...grpc.CallOption) (*CloseIncidentsResponse, error)
	GetIncidentEvents(ctx context.Context, in *GetIncidentEventsRequest, opts ...grpc.CallOption) (*GetIncidentEventsResponse, error)
	ResolveIncidents(ctx context.Context, in *ResolveIncidentsRequest, opts ...grpc.CallOption) (*ResolveIncidentsResponse, error)
	// This shouldn't be exposed in the docs, it has no external usecase
	GetIncidentUsingCorrelationKey(ctx context.Context, in *GetIncidentUsingCorrelationKeyRequest, opts ...grpc.CallOption) (*GetIncidentUsingCorrelationKeyResponse, error)
	ListIncidentEvents(ctx context.Context, in *ListIncidentEventsRequest, opts ...grpc.CallOption) (*ListIncidentEventsResponse, error)
	ListIncidentEventsTotalCount(ctx context.Context, in *ListIncidentEventsTotalCountRequest, opts ...grpc.CallOption) (*ListIncidentEventsTotalCountResponse, error)
	ListIncidentEventsFilterValues(ctx context.Context, in *ListIncidentEventsFilterValuesRequest, opts ...grpc.CallOption) (*ListIncidentEventsFilterValuesResponse, error)
}

type incidentsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIncidentsServiceClient(cc grpc.ClientConnInterface) IncidentsServiceClient {
	return &incidentsServiceClient{cc}
}

func (c *incidentsServiceClient) GetIncident(ctx context.Context, in *GetIncidentRequest, opts ...grpc.CallOption) (*GetIncidentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetIncidentResponse)
	err := c.cc.Invoke(ctx, IncidentsService_GetIncident_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incidentsServiceClient) BatchGetIncident(ctx context.Context, in *BatchGetIncidentRequest, opts ...grpc.CallOption) (*BatchGetIncidentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BatchGetIncidentResponse)
	err := c.cc.Invoke(ctx, IncidentsService_BatchGetIncident_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incidentsServiceClient) ListIncidents(ctx context.Context, in *ListIncidentsRequest, opts ...grpc.CallOption) (*ListIncidentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListIncidentsResponse)
	err := c.cc.Invoke(ctx, IncidentsService_ListIncidents_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incidentsServiceClient) ListIncidentAggregations(ctx context.Context, in *ListIncidentAggregationsRequest, opts ...grpc.CallOption) (*ListIncidentAggregationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListIncidentAggregationsResponse)
	err := c.cc.Invoke(ctx, IncidentsService_ListIncidentAggregations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incidentsServiceClient) GetFilterValues(ctx context.Context, in *GetFilterValuesRequest, opts ...grpc.CallOption) (*GetFilterValuesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFilterValuesResponse)
	err := c.cc.Invoke(ctx, IncidentsService_GetFilterValues_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incidentsServiceClient) AssignIncidents(ctx context.Context, in *AssignIncidentsRequest, opts ...grpc.CallOption) (*AssignIncidentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AssignIncidentsResponse)
	err := c.cc.Invoke(ctx, IncidentsService_AssignIncidents_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incidentsServiceClient) UnassignIncidents(ctx context.Context, in *UnassignIncidentsRequest, opts ...grpc.CallOption) (*UnassignIncidentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnassignIncidentsResponse)
	err := c.cc.Invoke(ctx, IncidentsService_UnassignIncidents_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incidentsServiceClient) AcknowledgeIncidents(ctx context.Context, in *AcknowledgeIncidentsRequest, opts ...grpc.CallOption) (*AcknowledgeIncidentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AcknowledgeIncidentsResponse)
	err := c.cc.Invoke(ctx, IncidentsService_AcknowledgeIncidents_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incidentsServiceClient) CloseIncidents(ctx context.Context, in *CloseIncidentsRequest, opts ...grpc.CallOption) (*CloseIncidentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CloseIncidentsResponse)
	err := c.cc.Invoke(ctx, IncidentsService_CloseIncidents_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incidentsServiceClient) GetIncidentEvents(ctx context.Context, in *GetIncidentEventsRequest, opts ...grpc.CallOption) (*GetIncidentEventsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetIncidentEventsResponse)
	err := c.cc.Invoke(ctx, IncidentsService_GetIncidentEvents_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incidentsServiceClient) ResolveIncidents(ctx context.Context, in *ResolveIncidentsRequest, opts ...grpc.CallOption) (*ResolveIncidentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResolveIncidentsResponse)
	err := c.cc.Invoke(ctx, IncidentsService_ResolveIncidents_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incidentsServiceClient) GetIncidentUsingCorrelationKey(ctx context.Context, in *GetIncidentUsingCorrelationKeyRequest, opts ...grpc.CallOption) (*GetIncidentUsingCorrelationKeyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetIncidentUsingCorrelationKeyResponse)
	err := c.cc.Invoke(ctx, IncidentsService_GetIncidentUsingCorrelationKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incidentsServiceClient) ListIncidentEvents(ctx context.Context, in *ListIncidentEventsRequest, opts ...grpc.CallOption) (*ListIncidentEventsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListIncidentEventsResponse)
	err := c.cc.Invoke(ctx, IncidentsService_ListIncidentEvents_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incidentsServiceClient) ListIncidentEventsTotalCount(ctx context.Context, in *ListIncidentEventsTotalCountRequest, opts ...grpc.CallOption) (*ListIncidentEventsTotalCountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListIncidentEventsTotalCountResponse)
	err := c.cc.Invoke(ctx, IncidentsService_ListIncidentEventsTotalCount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incidentsServiceClient) ListIncidentEventsFilterValues(ctx context.Context, in *ListIncidentEventsFilterValuesRequest, opts ...grpc.CallOption) (*ListIncidentEventsFilterValuesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListIncidentEventsFilterValuesResponse)
	err := c.cc.Invoke(ctx, IncidentsService_ListIncidentEventsFilterValues_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IncidentsServiceServer is the server API for IncidentsService service.
// All implementations must embed UnimplementedIncidentsServiceServer
// for forward compatibility.
type IncidentsServiceServer interface {
	GetIncident(context.Context, *GetIncidentRequest) (*GetIncidentResponse, error)
	BatchGetIncident(context.Context, *BatchGetIncidentRequest) (*BatchGetIncidentResponse, error)
	ListIncidents(context.Context, *ListIncidentsRequest) (*ListIncidentsResponse, error)
	ListIncidentAggregations(context.Context, *ListIncidentAggregationsRequest) (*ListIncidentAggregationsResponse, error)
	GetFilterValues(context.Context, *GetFilterValuesRequest) (*GetFilterValuesResponse, error)
	AssignIncidents(context.Context, *AssignIncidentsRequest) (*AssignIncidentsResponse, error)
	UnassignIncidents(context.Context, *UnassignIncidentsRequest) (*UnassignIncidentsResponse, error)
	AcknowledgeIncidents(context.Context, *AcknowledgeIncidentsRequest) (*AcknowledgeIncidentsResponse, error)
	CloseIncidents(context.Context, *CloseIncidentsRequest) (*CloseIncidentsResponse, error)
	GetIncidentEvents(context.Context, *GetIncidentEventsRequest) (*GetIncidentEventsResponse, error)
	ResolveIncidents(context.Context, *ResolveIncidentsRequest) (*ResolveIncidentsResponse, error)
	// This shouldn't be exposed in the docs, it has no external usecase
	GetIncidentUsingCorrelationKey(context.Context, *GetIncidentUsingCorrelationKeyRequest) (*GetIncidentUsingCorrelationKeyResponse, error)
	ListIncidentEvents(context.Context, *ListIncidentEventsRequest) (*ListIncidentEventsResponse, error)
	ListIncidentEventsTotalCount(context.Context, *ListIncidentEventsTotalCountRequest) (*ListIncidentEventsTotalCountResponse, error)
	ListIncidentEventsFilterValues(context.Context, *ListIncidentEventsFilterValuesRequest) (*ListIncidentEventsFilterValuesResponse, error)
	mustEmbedUnimplementedIncidentsServiceServer()
}

// UnimplementedIncidentsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedIncidentsServiceServer struct{}

func (UnimplementedIncidentsServiceServer) GetIncident(context.Context, *GetIncidentRequest) (*GetIncidentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIncident not implemented")
}
func (UnimplementedIncidentsServiceServer) BatchGetIncident(context.Context, *BatchGetIncidentRequest) (*BatchGetIncidentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGetIncident not implemented")
}
func (UnimplementedIncidentsServiceServer) ListIncidents(context.Context, *ListIncidentsRequest) (*ListIncidentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIncidents not implemented")
}
func (UnimplementedIncidentsServiceServer) ListIncidentAggregations(context.Context, *ListIncidentAggregationsRequest) (*ListIncidentAggregationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIncidentAggregations not implemented")
}
func (UnimplementedIncidentsServiceServer) GetFilterValues(context.Context, *GetFilterValuesRequest) (*GetFilterValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFilterValues not implemented")
}
func (UnimplementedIncidentsServiceServer) AssignIncidents(context.Context, *AssignIncidentsRequest) (*AssignIncidentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignIncidents not implemented")
}
func (UnimplementedIncidentsServiceServer) UnassignIncidents(context.Context, *UnassignIncidentsRequest) (*UnassignIncidentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnassignIncidents not implemented")
}
func (UnimplementedIncidentsServiceServer) AcknowledgeIncidents(context.Context, *AcknowledgeIncidentsRequest) (*AcknowledgeIncidentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcknowledgeIncidents not implemented")
}
func (UnimplementedIncidentsServiceServer) CloseIncidents(context.Context, *CloseIncidentsRequest) (*CloseIncidentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseIncidents not implemented")
}
func (UnimplementedIncidentsServiceServer) GetIncidentEvents(context.Context, *GetIncidentEventsRequest) (*GetIncidentEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIncidentEvents not implemented")
}
func (UnimplementedIncidentsServiceServer) ResolveIncidents(context.Context, *ResolveIncidentsRequest) (*ResolveIncidentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolveIncidents not implemented")
}
func (UnimplementedIncidentsServiceServer) GetIncidentUsingCorrelationKey(context.Context, *GetIncidentUsingCorrelationKeyRequest) (*GetIncidentUsingCorrelationKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIncidentUsingCorrelationKey not implemented")
}
func (UnimplementedIncidentsServiceServer) ListIncidentEvents(context.Context, *ListIncidentEventsRequest) (*ListIncidentEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIncidentEvents not implemented")
}
func (UnimplementedIncidentsServiceServer) ListIncidentEventsTotalCount(context.Context, *ListIncidentEventsTotalCountRequest) (*ListIncidentEventsTotalCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIncidentEventsTotalCount not implemented")
}
func (UnimplementedIncidentsServiceServer) ListIncidentEventsFilterValues(context.Context, *ListIncidentEventsFilterValuesRequest) (*ListIncidentEventsFilterValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIncidentEventsFilterValues not implemented")
}
func (UnimplementedIncidentsServiceServer) mustEmbedUnimplementedIncidentsServiceServer() {}
func (UnimplementedIncidentsServiceServer) testEmbeddedByValue()                          {}

// UnsafeIncidentsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IncidentsServiceServer will
// result in compilation errors.
type UnsafeIncidentsServiceServer interface {
	mustEmbedUnimplementedIncidentsServiceServer()
}

func RegisterIncidentsServiceServer(s grpc.ServiceRegistrar, srv IncidentsServiceServer) {
	// If the following call pancis, it indicates UnimplementedIncidentsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&IncidentsService_ServiceDesc, srv)
}

func _IncidentsService_GetIncident_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIncidentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncidentsServiceServer).GetIncident(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IncidentsService_GetIncident_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncidentsServiceServer).GetIncident(ctx, req.(*GetIncidentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IncidentsService_BatchGetIncident_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetIncidentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncidentsServiceServer).BatchGetIncident(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IncidentsService_BatchGetIncident_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncidentsServiceServer).BatchGetIncident(ctx, req.(*BatchGetIncidentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IncidentsService_ListIncidents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIncidentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncidentsServiceServer).ListIncidents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IncidentsService_ListIncidents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncidentsServiceServer).ListIncidents(ctx, req.(*ListIncidentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IncidentsService_ListIncidentAggregations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIncidentAggregationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncidentsServiceServer).ListIncidentAggregations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IncidentsService_ListIncidentAggregations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncidentsServiceServer).ListIncidentAggregations(ctx, req.(*ListIncidentAggregationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IncidentsService_GetFilterValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFilterValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncidentsServiceServer).GetFilterValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IncidentsService_GetFilterValues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncidentsServiceServer).GetFilterValues(ctx, req.(*GetFilterValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IncidentsService_AssignIncidents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignIncidentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncidentsServiceServer).AssignIncidents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IncidentsService_AssignIncidents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncidentsServiceServer).AssignIncidents(ctx, req.(*AssignIncidentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IncidentsService_UnassignIncidents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnassignIncidentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncidentsServiceServer).UnassignIncidents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IncidentsService_UnassignIncidents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncidentsServiceServer).UnassignIncidents(ctx, req.(*UnassignIncidentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IncidentsService_AcknowledgeIncidents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcknowledgeIncidentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncidentsServiceServer).AcknowledgeIncidents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IncidentsService_AcknowledgeIncidents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncidentsServiceServer).AcknowledgeIncidents(ctx, req.(*AcknowledgeIncidentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IncidentsService_CloseIncidents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseIncidentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncidentsServiceServer).CloseIncidents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IncidentsService_CloseIncidents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncidentsServiceServer).CloseIncidents(ctx, req.(*CloseIncidentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IncidentsService_GetIncidentEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIncidentEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncidentsServiceServer).GetIncidentEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IncidentsService_GetIncidentEvents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncidentsServiceServer).GetIncidentEvents(ctx, req.(*GetIncidentEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IncidentsService_ResolveIncidents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveIncidentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncidentsServiceServer).ResolveIncidents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IncidentsService_ResolveIncidents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncidentsServiceServer).ResolveIncidents(ctx, req.(*ResolveIncidentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IncidentsService_GetIncidentUsingCorrelationKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIncidentUsingCorrelationKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncidentsServiceServer).GetIncidentUsingCorrelationKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IncidentsService_GetIncidentUsingCorrelationKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncidentsServiceServer).GetIncidentUsingCorrelationKey(ctx, req.(*GetIncidentUsingCorrelationKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IncidentsService_ListIncidentEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIncidentEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncidentsServiceServer).ListIncidentEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IncidentsService_ListIncidentEvents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncidentsServiceServer).ListIncidentEvents(ctx, req.(*ListIncidentEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IncidentsService_ListIncidentEventsTotalCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIncidentEventsTotalCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncidentsServiceServer).ListIncidentEventsTotalCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IncidentsService_ListIncidentEventsTotalCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncidentsServiceServer).ListIncidentEventsTotalCount(ctx, req.(*ListIncidentEventsTotalCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IncidentsService_ListIncidentEventsFilterValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIncidentEventsFilterValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncidentsServiceServer).ListIncidentEventsFilterValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IncidentsService_ListIncidentEventsFilterValues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncidentsServiceServer).ListIncidentEventsFilterValues(ctx, req.(*ListIncidentEventsFilterValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IncidentsService_ServiceDesc is the grpc.ServiceDesc for IncidentsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IncidentsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.incidents.v1.IncidentsService",
	HandlerType: (*IncidentsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetIncident",
			Handler:    _IncidentsService_GetIncident_Handler,
		},
		{
			MethodName: "BatchGetIncident",
			Handler:    _IncidentsService_BatchGetIncident_Handler,
		},
		{
			MethodName: "ListIncidents",
			Handler:    _IncidentsService_ListIncidents_Handler,
		},
		{
			MethodName: "ListIncidentAggregations",
			Handler:    _IncidentsService_ListIncidentAggregations_Handler,
		},
		{
			MethodName: "GetFilterValues",
			Handler:    _IncidentsService_GetFilterValues_Handler,
		},
		{
			MethodName: "AssignIncidents",
			Handler:    _IncidentsService_AssignIncidents_Handler,
		},
		{
			MethodName: "UnassignIncidents",
			Handler:    _IncidentsService_UnassignIncidents_Handler,
		},
		{
			MethodName: "AcknowledgeIncidents",
			Handler:    _IncidentsService_AcknowledgeIncidents_Handler,
		},
		{
			MethodName: "CloseIncidents",
			Handler:    _IncidentsService_CloseIncidents_Handler,
		},
		{
			MethodName: "GetIncidentEvents",
			Handler:    _IncidentsService_GetIncidentEvents_Handler,
		},
		{
			MethodName: "ResolveIncidents",
			Handler:    _IncidentsService_ResolveIncidents_Handler,
		},
		{
			MethodName: "GetIncidentUsingCorrelationKey",
			Handler:    _IncidentsService_GetIncidentUsingCorrelationKey_Handler,
		},
		{
			MethodName: "ListIncidentEvents",
			Handler:    _IncidentsService_ListIncidentEvents_Handler,
		},
		{
			MethodName: "ListIncidentEventsTotalCount",
			Handler:    _IncidentsService_ListIncidentEventsTotalCount_Handler,
		},
		{
			MethodName: "ListIncidentEventsFilterValues",
			Handler:    _IncidentsService_ListIncidentEventsFilterValues_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/incidents/v1/incidents_service.proto",
}
