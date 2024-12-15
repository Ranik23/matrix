// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: proto/worker.proto

package proto

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
	WorkerService_GetStatus_FullMethodName                        = "/worker.WorkerService/GetStatus"
	WorkerService_GetLinearFormCalculation_FullMethodName         = "/worker.WorkerService/GetLinearFormCalculation"
	WorkerService_GetParallelLinearFormCalculation_FullMethodName = "/worker.WorkerService/GetParallelLinearFormCalculation"
	WorkerService_GetPolynomialCalculation_FullMethodName         = "/worker.WorkerService/GetPolynomialCalculation"
	WorkerService_GetParallelPolynomialCalculation_FullMethodName = "/worker.WorkerService/GetParallelPolynomialCalculation"
)

// WorkerServiceClient is the client API for WorkerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkerServiceClient interface {
	GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error)
	GetLinearFormCalculation(ctx context.Context, in *GetLinearFormCalculationRequest, opts ...grpc.CallOption) (*GetLinearFormCalculationResponse, error)
	GetParallelLinearFormCalculation(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[GetParallelLinearFormCalculationRequest, GetParallelLinearFormCalculationResponse], error)
	GetPolynomialCalculation(ctx context.Context, in *GetPolynomialCalculationRequest, opts ...grpc.CallOption) (*GetPolynomialCalculationResponse, error)
	GetParallelPolynomialCalculation(ctx context.Context, in *GetParallelPolynomialCalculationRequest, opts ...grpc.CallOption) (*GetParallelPolynomialCalculationResponse, error)
}

type workerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerServiceClient(cc grpc.ClientConnInterface) WorkerServiceClient {
	return &workerServiceClient{cc}
}

func (c *workerServiceClient) GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStatusResponse)
	err := c.cc.Invoke(ctx, WorkerService_GetStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) GetLinearFormCalculation(ctx context.Context, in *GetLinearFormCalculationRequest, opts ...grpc.CallOption) (*GetLinearFormCalculationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLinearFormCalculationResponse)
	err := c.cc.Invoke(ctx, WorkerService_GetLinearFormCalculation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) GetParallelLinearFormCalculation(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[GetParallelLinearFormCalculationRequest, GetParallelLinearFormCalculationResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &WorkerService_ServiceDesc.Streams[0], WorkerService_GetParallelLinearFormCalculation_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetParallelLinearFormCalculationRequest, GetParallelLinearFormCalculationResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type WorkerService_GetParallelLinearFormCalculationClient = grpc.BidiStreamingClient[GetParallelLinearFormCalculationRequest, GetParallelLinearFormCalculationResponse]

func (c *workerServiceClient) GetPolynomialCalculation(ctx context.Context, in *GetPolynomialCalculationRequest, opts ...grpc.CallOption) (*GetPolynomialCalculationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPolynomialCalculationResponse)
	err := c.cc.Invoke(ctx, WorkerService_GetPolynomialCalculation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) GetParallelPolynomialCalculation(ctx context.Context, in *GetParallelPolynomialCalculationRequest, opts ...grpc.CallOption) (*GetParallelPolynomialCalculationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetParallelPolynomialCalculationResponse)
	err := c.cc.Invoke(ctx, WorkerService_GetParallelPolynomialCalculation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServiceServer is the server API for WorkerService service.
// All implementations must embed UnimplementedWorkerServiceServer
// for forward compatibility.
type WorkerServiceServer interface {
	GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error)
	GetLinearFormCalculation(context.Context, *GetLinearFormCalculationRequest) (*GetLinearFormCalculationResponse, error)
	GetParallelLinearFormCalculation(grpc.BidiStreamingServer[GetParallelLinearFormCalculationRequest, GetParallelLinearFormCalculationResponse]) error
	GetPolynomialCalculation(context.Context, *GetPolynomialCalculationRequest) (*GetPolynomialCalculationResponse, error)
	GetParallelPolynomialCalculation(context.Context, *GetParallelPolynomialCalculationRequest) (*GetParallelPolynomialCalculationResponse, error)
	mustEmbedUnimplementedWorkerServiceServer()
}

// UnimplementedWorkerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWorkerServiceServer struct{}

func (UnimplementedWorkerServiceServer) GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedWorkerServiceServer) GetLinearFormCalculation(context.Context, *GetLinearFormCalculationRequest) (*GetLinearFormCalculationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLinearFormCalculation not implemented")
}
func (UnimplementedWorkerServiceServer) GetParallelLinearFormCalculation(grpc.BidiStreamingServer[GetParallelLinearFormCalculationRequest, GetParallelLinearFormCalculationResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetParallelLinearFormCalculation not implemented")
}
func (UnimplementedWorkerServiceServer) GetPolynomialCalculation(context.Context, *GetPolynomialCalculationRequest) (*GetPolynomialCalculationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPolynomialCalculation not implemented")
}
func (UnimplementedWorkerServiceServer) GetParallelPolynomialCalculation(context.Context, *GetParallelPolynomialCalculationRequest) (*GetParallelPolynomialCalculationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParallelPolynomialCalculation not implemented")
}
func (UnimplementedWorkerServiceServer) mustEmbedUnimplementedWorkerServiceServer() {}
func (UnimplementedWorkerServiceServer) testEmbeddedByValue()                       {}

// UnsafeWorkerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkerServiceServer will
// result in compilation errors.
type UnsafeWorkerServiceServer interface {
	mustEmbedUnimplementedWorkerServiceServer()
}

func RegisterWorkerServiceServer(s grpc.ServiceRegistrar, srv WorkerServiceServer) {
	// If the following call pancis, it indicates UnimplementedWorkerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WorkerService_ServiceDesc, srv)
}

func _WorkerService_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkerService_GetStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).GetStatus(ctx, req.(*GetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_GetLinearFormCalculation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLinearFormCalculationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).GetLinearFormCalculation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkerService_GetLinearFormCalculation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).GetLinearFormCalculation(ctx, req.(*GetLinearFormCalculationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_GetParallelLinearFormCalculation_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WorkerServiceServer).GetParallelLinearFormCalculation(&grpc.GenericServerStream[GetParallelLinearFormCalculationRequest, GetParallelLinearFormCalculationResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type WorkerService_GetParallelLinearFormCalculationServer = grpc.BidiStreamingServer[GetParallelLinearFormCalculationRequest, GetParallelLinearFormCalculationResponse]

func _WorkerService_GetPolynomialCalculation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPolynomialCalculationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).GetPolynomialCalculation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkerService_GetPolynomialCalculation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).GetPolynomialCalculation(ctx, req.(*GetPolynomialCalculationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_GetParallelPolynomialCalculation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParallelPolynomialCalculationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).GetParallelPolynomialCalculation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkerService_GetParallelPolynomialCalculation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).GetParallelPolynomialCalculation(ctx, req.(*GetParallelPolynomialCalculationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkerService_ServiceDesc is the grpc.ServiceDesc for WorkerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "worker.WorkerService",
	HandlerType: (*WorkerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _WorkerService_GetStatus_Handler,
		},
		{
			MethodName: "GetLinearFormCalculation",
			Handler:    _WorkerService_GetLinearFormCalculation_Handler,
		},
		{
			MethodName: "GetPolynomialCalculation",
			Handler:    _WorkerService_GetPolynomialCalculation_Handler,
		},
		{
			MethodName: "GetParallelPolynomialCalculation",
			Handler:    _WorkerService_GetParallelPolynomialCalculation_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetParallelLinearFormCalculation",
			Handler:       _WorkerService_GetParallelLinearFormCalculation_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/worker.proto",
}
