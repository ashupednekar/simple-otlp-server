// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: otlp.proto

package otlp

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
	TraceService_Export_FullMethodName = "/otlp.TraceService/Export"
)

// TraceServiceClient is the client API for TraceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TraceServiceClient interface {
	Export(ctx context.Context, in *ExportTraceServiceRequest, opts ...grpc.CallOption) (*ExportTraceServiceResponse, error)
}

type traceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTraceServiceClient(cc grpc.ClientConnInterface) TraceServiceClient {
	return &traceServiceClient{cc}
}

func (c *traceServiceClient) Export(ctx context.Context, in *ExportTraceServiceRequest, opts ...grpc.CallOption) (*ExportTraceServiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExportTraceServiceResponse)
	err := c.cc.Invoke(ctx, TraceService_Export_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TraceServiceServer is the server API for TraceService service.
// All implementations must embed UnimplementedTraceServiceServer
// for forward compatibility.
type TraceServiceServer interface {
	Export(context.Context, *ExportTraceServiceRequest) (*ExportTraceServiceResponse, error)
	mustEmbedUnimplementedTraceServiceServer()
}

// UnimplementedTraceServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTraceServiceServer struct{}

func (UnimplementedTraceServiceServer) Export(context.Context, *ExportTraceServiceRequest) (*ExportTraceServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Export not implemented")
}
func (UnimplementedTraceServiceServer) mustEmbedUnimplementedTraceServiceServer() {}
func (UnimplementedTraceServiceServer) testEmbeddedByValue()                      {}

// UnsafeTraceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TraceServiceServer will
// result in compilation errors.
type UnsafeTraceServiceServer interface {
	mustEmbedUnimplementedTraceServiceServer()
}

func RegisterTraceServiceServer(s grpc.ServiceRegistrar, srv TraceServiceServer) {
	// If the following call pancis, it indicates UnimplementedTraceServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TraceService_ServiceDesc, srv)
}

func _TraceService_Export_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportTraceServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceServiceServer).Export(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TraceService_Export_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceServiceServer).Export(ctx, req.(*ExportTraceServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TraceService_ServiceDesc is the grpc.ServiceDesc for TraceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TraceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "otlp.TraceService",
	HandlerType: (*TraceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Export",
			Handler:    _TraceService_Export_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "otlp.proto",
}
