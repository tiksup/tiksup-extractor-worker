// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: internal/proto/user_interacctions.proto

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
	InteractionsService_ProcessData_FullMethodName = "/interactions.InteractionsService/ProcessData"
)

// InteractionsServiceClient is the client API for InteractionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InteractionsServiceClient interface {
	ProcessData(ctx context.Context, in *PreprocessedDataRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
}

type interactionsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInteractionsServiceClient(cc grpc.ClientConnInterface) InteractionsServiceClient {
	return &interactionsServiceClient{cc}
}

func (c *interactionsServiceClient) ProcessData(ctx context.Context, in *PreprocessedDataRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, InteractionsService_ProcessData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InteractionsServiceServer is the server API for InteractionsService service.
// All implementations must embed UnimplementedInteractionsServiceServer
// for forward compatibility.
type InteractionsServiceServer interface {
	ProcessData(context.Context, *PreprocessedDataRequest) (*SuccessResponse, error)
	mustEmbedUnimplementedInteractionsServiceServer()
}

// UnimplementedInteractionsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedInteractionsServiceServer struct{}

func (UnimplementedInteractionsServiceServer) ProcessData(context.Context, *PreprocessedDataRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessData not implemented")
}
func (UnimplementedInteractionsServiceServer) mustEmbedUnimplementedInteractionsServiceServer() {}
func (UnimplementedInteractionsServiceServer) testEmbeddedByValue()                             {}

// UnsafeInteractionsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InteractionsServiceServer will
// result in compilation errors.
type UnsafeInteractionsServiceServer interface {
	mustEmbedUnimplementedInteractionsServiceServer()
}

func RegisterInteractionsServiceServer(s grpc.ServiceRegistrar, srv InteractionsServiceServer) {
	// If the following call pancis, it indicates UnimplementedInteractionsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&InteractionsService_ServiceDesc, srv)
}

func _InteractionsService_ProcessData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreprocessedDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionsServiceServer).ProcessData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractionsService_ProcessData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionsServiceServer).ProcessData(ctx, req.(*PreprocessedDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InteractionsService_ServiceDesc is the grpc.ServiceDesc for InteractionsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InteractionsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "interactions.InteractionsService",
	HandlerType: (*InteractionsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessData",
			Handler:    _InteractionsService_ProcessData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/proto/user_interacctions.proto",
}
