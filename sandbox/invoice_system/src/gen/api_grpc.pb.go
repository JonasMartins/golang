// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: src/proto/api.proto

package gen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FinancialServiceClient is the client API for FinancialService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FinancialServiceClient interface {
	AddInvoice(ctx context.Context, in *AddInvoiceRequest, opts ...grpc.CallOption) (*BasicResponse, error)
}

type financialServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFinancialServiceClient(cc grpc.ClientConnInterface) FinancialServiceClient {
	return &financialServiceClient{cc}
}

func (c *financialServiceClient) AddInvoice(ctx context.Context, in *AddInvoiceRequest, opts ...grpc.CallOption) (*BasicResponse, error) {
	out := new(BasicResponse)
	err := c.cc.Invoke(ctx, "/project.proto.FinancialService/AddInvoice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FinancialServiceServer is the server API for FinancialService service.
// All implementations should embed UnimplementedFinancialServiceServer
// for forward compatibility
type FinancialServiceServer interface {
	AddInvoice(context.Context, *AddInvoiceRequest) (*BasicResponse, error)
}

// UnimplementedFinancialServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFinancialServiceServer struct {
}

func (UnimplementedFinancialServiceServer) AddInvoice(context.Context, *AddInvoiceRequest) (*BasicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddInvoice not implemented")
}

// UnsafeFinancialServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FinancialServiceServer will
// result in compilation errors.
type UnsafeFinancialServiceServer interface {
	mustEmbedUnimplementedFinancialServiceServer()
}

func RegisterFinancialServiceServer(s grpc.ServiceRegistrar, srv FinancialServiceServer) {
	s.RegisterService(&FinancialService_ServiceDesc, srv)
}

func _FinancialService_AddInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialServiceServer).AddInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.proto.FinancialService/AddInvoice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialServiceServer).AddInvoice(ctx, req.(*AddInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FinancialService_ServiceDesc is the grpc.ServiceDesc for FinancialService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FinancialService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "project.proto.FinancialService",
	HandlerType: (*FinancialServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddInvoice",
			Handler:    _FinancialService_AddInvoice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/proto/api.proto",
}
