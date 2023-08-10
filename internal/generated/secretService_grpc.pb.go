// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: secretService.proto

package generated

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

const (
	SecretService_Ping_FullMethodName         = "/com.github.MaxReX92.go_yandex_gophkeeper.SecretService/Ping"
	SecretService_AddSecret_FullMethodName    = "/com.github.MaxReX92.go_yandex_gophkeeper.SecretService/AddSecret"
	SecretService_ChangeSecret_FullMethodName = "/com.github.MaxReX92.go_yandex_gophkeeper.SecretService/ChangeSecret"
	SecretService_RemoveSecret_FullMethodName = "/com.github.MaxReX92.go_yandex_gophkeeper.SecretService/RemoveSecret"
	SecretService_SecretEvents_FullMethodName = "/com.github.MaxReX92.go_yandex_gophkeeper.SecretService/SecretEvents"
)

// SecretServiceClient is the client API for SecretService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecretServiceClient interface {
	Ping(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error)
	AddSecret(ctx context.Context, in *SecretRequest, opts ...grpc.CallOption) (*Void, error)
	ChangeSecret(ctx context.Context, in *SecretRequest, opts ...grpc.CallOption) (*Void, error)
	RemoveSecret(ctx context.Context, in *SecretRequest, opts ...grpc.CallOption) (*Void, error)
	SecretEvents(ctx context.Context, in *User, opts ...grpc.CallOption) (SecretService_SecretEventsClient, error)
}

type secretServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSecretServiceClient(cc grpc.ClientConnInterface) SecretServiceClient {
	return &secretServiceClient{cc}
}

func (c *secretServiceClient) Ping(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, SecretService_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretServiceClient) AddSecret(ctx context.Context, in *SecretRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, SecretService_AddSecret_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretServiceClient) ChangeSecret(ctx context.Context, in *SecretRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, SecretService_ChangeSecret_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretServiceClient) RemoveSecret(ctx context.Context, in *SecretRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, SecretService_RemoveSecret_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretServiceClient) SecretEvents(ctx context.Context, in *User, opts ...grpc.CallOption) (SecretService_SecretEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &SecretService_ServiceDesc.Streams[0], SecretService_SecretEvents_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &secretServiceSecretEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SecretService_SecretEventsClient interface {
	Recv() (*SecretEvent, error)
	grpc.ClientStream
}

type secretServiceSecretEventsClient struct {
	grpc.ClientStream
}

func (x *secretServiceSecretEventsClient) Recv() (*SecretEvent, error) {
	m := new(SecretEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SecretServiceServer is the server API for SecretService service.
// All implementations must embed UnimplementedSecretServiceServer
// for forward compatibility
type SecretServiceServer interface {
	Ping(context.Context, *Void) (*Void, error)
	AddSecret(context.Context, *SecretRequest) (*Void, error)
	ChangeSecret(context.Context, *SecretRequest) (*Void, error)
	RemoveSecret(context.Context, *SecretRequest) (*Void, error)
	SecretEvents(*User, SecretService_SecretEventsServer) error
	mustEmbedUnimplementedSecretServiceServer()
}

// UnimplementedSecretServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSecretServiceServer struct {
}

func (UnimplementedSecretServiceServer) Ping(context.Context, *Void) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedSecretServiceServer) AddSecret(context.Context, *SecretRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSecret not implemented")
}
func (UnimplementedSecretServiceServer) ChangeSecret(context.Context, *SecretRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeSecret not implemented")
}
func (UnimplementedSecretServiceServer) RemoveSecret(context.Context, *SecretRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSecret not implemented")
}
func (UnimplementedSecretServiceServer) SecretEvents(*User, SecretService_SecretEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method SecretEvents not implemented")
}
func (UnimplementedSecretServiceServer) mustEmbedUnimplementedSecretServiceServer() {}

// UnsafeSecretServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecretServiceServer will
// result in compilation errors.
type UnsafeSecretServiceServer interface {
	mustEmbedUnimplementedSecretServiceServer()
}

func RegisterSecretServiceServer(s grpc.ServiceRegistrar, srv SecretServiceServer) {
	s.RegisterService(&SecretService_ServiceDesc, srv)
}

func _SecretService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecretService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).Ping(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretService_AddSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).AddSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecretService_AddSecret_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).AddSecret(ctx, req.(*SecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretService_ChangeSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).ChangeSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecretService_ChangeSecret_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).ChangeSecret(ctx, req.(*SecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretService_RemoveSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).RemoveSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecretService_RemoveSecret_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).RemoveSecret(ctx, req.(*SecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretService_SecretEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(User)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SecretServiceServer).SecretEvents(m, &secretServiceSecretEventsServer{stream})
}

type SecretService_SecretEventsServer interface {
	Send(*SecretEvent) error
	grpc.ServerStream
}

type secretServiceSecretEventsServer struct {
	grpc.ServerStream
}

func (x *secretServiceSecretEventsServer) Send(m *SecretEvent) error {
	return x.ServerStream.SendMsg(m)
}

// SecretService_ServiceDesc is the grpc.ServiceDesc for SecretService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecretService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.github.MaxReX92.go_yandex_gophkeeper.SecretService",
	HandlerType: (*SecretServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _SecretService_Ping_Handler,
		},
		{
			MethodName: "AddSecret",
			Handler:    _SecretService_AddSecret_Handler,
		},
		{
			MethodName: "ChangeSecret",
			Handler:    _SecretService_ChangeSecret_Handler,
		},
		{
			MethodName: "RemoveSecret",
			Handler:    _SecretService_RemoveSecret_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SecretEvents",
			Handler:       _SecretService_SecretEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "secretService.proto",
}