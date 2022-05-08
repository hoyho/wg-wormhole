// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: wgSvc.proto

package proto

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

// RegistryClient is the client API for Registry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegistryClient interface {
	GetEndpoint(ctx context.Context, in *GetEndpointRequest, opts ...grpc.CallOption) (*GetEndpointReply, error)
	RegisterEndpoint(ctx context.Context, in *RegRequest, opts ...grpc.CallOption) (*RegReply, error)
}

type registryClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistryClient(cc grpc.ClientConnInterface) RegistryClient {
	return &registryClient{cc}
}

func (c *registryClient) GetEndpoint(ctx context.Context, in *GetEndpointRequest, opts ...grpc.CallOption) (*GetEndpointReply, error) {
	out := new(GetEndpointReply)
	err := c.cc.Invoke(ctx, "/wgSvc.proto.Registry/GetEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) RegisterEndpoint(ctx context.Context, in *RegRequest, opts ...grpc.CallOption) (*RegReply, error) {
	out := new(RegReply)
	err := c.cc.Invoke(ctx, "/wgSvc.proto.Registry/RegisterEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistryServer is the server API for Registry service.
// All implementations must embed UnimplementedRegistryServer
// for forward compatibility
type RegistryServer interface {
	GetEndpoint(context.Context, *GetEndpointRequest) (*GetEndpointReply, error)
	RegisterEndpoint(context.Context, *RegRequest) (*RegReply, error)
	mustEmbedUnimplementedRegistryServer()
}

// UnimplementedRegistryServer must be embedded to have forward compatible implementations.
type UnimplementedRegistryServer struct {
}

func (UnimplementedRegistryServer) GetEndpoint(context.Context, *GetEndpointRequest) (*GetEndpointReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEndpoint not implemented")
}
func (UnimplementedRegistryServer) RegisterEndpoint(context.Context, *RegRequest) (*RegReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterEndpoint not implemented")
}
func (UnimplementedRegistryServer) mustEmbedUnimplementedRegistryServer() {}

// UnsafeRegistryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegistryServer will
// result in compilation errors.
type UnsafeRegistryServer interface {
	mustEmbedUnimplementedRegistryServer()
}

func RegisterRegistryServer(s grpc.ServiceRegistrar, srv RegistryServer) {
	s.RegisterService(&Registry_ServiceDesc, srv)
}

func _Registry_GetEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).GetEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wgSvc.proto.Registry/GetEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).GetEndpoint(ctx, req.(*GetEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_RegisterEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).RegisterEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wgSvc.proto.Registry/RegisterEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).RegisterEndpoint(ctx, req.(*RegRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Registry_ServiceDesc is the grpc.ServiceDesc for Registry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Registry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wgSvc.proto.Registry",
	HandlerType: (*RegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEndpoint",
			Handler:    _Registry_GetEndpoint_Handler,
		},
		{
			MethodName: "RegisterEndpoint",
			Handler:    _Registry_RegisterEndpoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wgSvc.proto",
}
