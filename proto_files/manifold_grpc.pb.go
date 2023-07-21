// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: manifold.proto

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

const (
	Manifold_Enqueue_FullMethodName       = "/manifold.Manifold/Enqueue"
	Manifold_Dequeue_FullMethodName       = "/manifold.Manifold/Dequeue"
	Manifold_StreamDequeue_FullMethodName = "/manifold.Manifold/StreamDequeue"
	Manifold_StreamEnqueue_FullMethodName = "/manifold.Manifold/StreamEnqueue"
	Manifold_ListTopics_FullMethodName    = "/manifold.Manifold/ListTopics"
	Manifold_DeleteTopic_FullMethodName   = "/manifold.Manifold/DeleteTopic"
	Manifold_TopicLength_FullMethodName   = "/manifold.Manifold/TopicLength"
)

// ManifoldClient is the client API for Manifold service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManifoldClient interface {
	Enqueue(ctx context.Context, in *EnqueueMsg, opts ...grpc.CallOption) (*EnqueueAck, error)
	Dequeue(ctx context.Context, in *DequeueMsg, opts ...grpc.CallOption) (*DequeueAck, error)
	StreamDequeue(ctx context.Context, in *DequeueMsg, opts ...grpc.CallOption) (Manifold_StreamDequeueClient, error)
	StreamEnqueue(ctx context.Context, in *EnqueueMsg, opts ...grpc.CallOption) (Manifold_StreamEnqueueClient, error)
	ListTopics(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StringList, error)
	DeleteTopic(ctx context.Context, in *DeleteTopicMsg, opts ...grpc.CallOption) (*Empty, error)
	TopicLength(ctx context.Context, in *DequeueMsg, opts ...grpc.CallOption) (*Length, error)
}

type manifoldClient struct {
	cc grpc.ClientConnInterface
}

func NewManifoldClient(cc grpc.ClientConnInterface) ManifoldClient {
	return &manifoldClient{cc}
}

func (c *manifoldClient) Enqueue(ctx context.Context, in *EnqueueMsg, opts ...grpc.CallOption) (*EnqueueAck, error) {
	out := new(EnqueueAck)
	err := c.cc.Invoke(ctx, Manifold_Enqueue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manifoldClient) Dequeue(ctx context.Context, in *DequeueMsg, opts ...grpc.CallOption) (*DequeueAck, error) {
	out := new(DequeueAck)
	err := c.cc.Invoke(ctx, Manifold_Dequeue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manifoldClient) StreamDequeue(ctx context.Context, in *DequeueMsg, opts ...grpc.CallOption) (Manifold_StreamDequeueClient, error) {
	stream, err := c.cc.NewStream(ctx, &Manifold_ServiceDesc.Streams[0], Manifold_StreamDequeue_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &manifoldStreamDequeueClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Manifold_StreamDequeueClient interface {
	Recv() (*DequeueAck, error)
	grpc.ClientStream
}

type manifoldStreamDequeueClient struct {
	grpc.ClientStream
}

func (x *manifoldStreamDequeueClient) Recv() (*DequeueAck, error) {
	m := new(DequeueAck)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *manifoldClient) StreamEnqueue(ctx context.Context, in *EnqueueMsg, opts ...grpc.CallOption) (Manifold_StreamEnqueueClient, error) {
	stream, err := c.cc.NewStream(ctx, &Manifold_ServiceDesc.Streams[1], Manifold_StreamEnqueue_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &manifoldStreamEnqueueClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Manifold_StreamEnqueueClient interface {
	Recv() (*EnqueueAck, error)
	grpc.ClientStream
}

type manifoldStreamEnqueueClient struct {
	grpc.ClientStream
}

func (x *manifoldStreamEnqueueClient) Recv() (*EnqueueAck, error) {
	m := new(EnqueueAck)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *manifoldClient) ListTopics(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StringList, error) {
	out := new(StringList)
	err := c.cc.Invoke(ctx, Manifold_ListTopics_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manifoldClient) DeleteTopic(ctx context.Context, in *DeleteTopicMsg, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Manifold_DeleteTopic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manifoldClient) TopicLength(ctx context.Context, in *DequeueMsg, opts ...grpc.CallOption) (*Length, error) {
	out := new(Length)
	err := c.cc.Invoke(ctx, Manifold_TopicLength_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManifoldServer is the server API for Manifold service.
// All implementations must embed UnimplementedManifoldServer
// for forward compatibility
type ManifoldServer interface {
	Enqueue(context.Context, *EnqueueMsg) (*EnqueueAck, error)
	Dequeue(context.Context, *DequeueMsg) (*DequeueAck, error)
	StreamDequeue(*DequeueMsg, Manifold_StreamDequeueServer) error
	StreamEnqueue(*EnqueueMsg, Manifold_StreamEnqueueServer) error
	ListTopics(context.Context, *Empty) (*StringList, error)
	DeleteTopic(context.Context, *DeleteTopicMsg) (*Empty, error)
	TopicLength(context.Context, *DequeueMsg) (*Length, error)
	mustEmbedUnimplementedManifoldServer()
}

// UnimplementedManifoldServer must be embedded to have forward compatible implementations.
type UnimplementedManifoldServer struct {
}

func (UnimplementedManifoldServer) Enqueue(context.Context, *EnqueueMsg) (*EnqueueAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enqueue not implemented")
}
func (UnimplementedManifoldServer) Dequeue(context.Context, *DequeueMsg) (*DequeueAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dequeue not implemented")
}
func (UnimplementedManifoldServer) StreamDequeue(*DequeueMsg, Manifold_StreamDequeueServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamDequeue not implemented")
}
func (UnimplementedManifoldServer) StreamEnqueue(*EnqueueMsg, Manifold_StreamEnqueueServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamEnqueue not implemented")
}
func (UnimplementedManifoldServer) ListTopics(context.Context, *Empty) (*StringList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTopics not implemented")
}
func (UnimplementedManifoldServer) DeleteTopic(context.Context, *DeleteTopicMsg) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTopic not implemented")
}
func (UnimplementedManifoldServer) TopicLength(context.Context, *DequeueMsg) (*Length, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TopicLength not implemented")
}
func (UnimplementedManifoldServer) mustEmbedUnimplementedManifoldServer() {}

// UnsafeManifoldServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManifoldServer will
// result in compilation errors.
type UnsafeManifoldServer interface {
	mustEmbedUnimplementedManifoldServer()
}

func RegisterManifoldServer(s grpc.ServiceRegistrar, srv ManifoldServer) {
	s.RegisterService(&Manifold_ServiceDesc, srv)
}

func _Manifold_Enqueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnqueueMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManifoldServer).Enqueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manifold_Enqueue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManifoldServer).Enqueue(ctx, req.(*EnqueueMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manifold_Dequeue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DequeueMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManifoldServer).Dequeue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manifold_Dequeue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManifoldServer).Dequeue(ctx, req.(*DequeueMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manifold_StreamDequeue_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DequeueMsg)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ManifoldServer).StreamDequeue(m, &manifoldStreamDequeueServer{stream})
}

type Manifold_StreamDequeueServer interface {
	Send(*DequeueAck) error
	grpc.ServerStream
}

type manifoldStreamDequeueServer struct {
	grpc.ServerStream
}

func (x *manifoldStreamDequeueServer) Send(m *DequeueAck) error {
	return x.ServerStream.SendMsg(m)
}

func _Manifold_StreamEnqueue_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EnqueueMsg)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ManifoldServer).StreamEnqueue(m, &manifoldStreamEnqueueServer{stream})
}

type Manifold_StreamEnqueueServer interface {
	Send(*EnqueueAck) error
	grpc.ServerStream
}

type manifoldStreamEnqueueServer struct {
	grpc.ServerStream
}

func (x *manifoldStreamEnqueueServer) Send(m *EnqueueAck) error {
	return x.ServerStream.SendMsg(m)
}

func _Manifold_ListTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManifoldServer).ListTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manifold_ListTopics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManifoldServer).ListTopics(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manifold_DeleteTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTopicMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManifoldServer).DeleteTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manifold_DeleteTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManifoldServer).DeleteTopic(ctx, req.(*DeleteTopicMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manifold_TopicLength_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DequeueMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManifoldServer).TopicLength(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manifold_TopicLength_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManifoldServer).TopicLength(ctx, req.(*DequeueMsg))
	}
	return interceptor(ctx, in, info, handler)
}

// Manifold_ServiceDesc is the grpc.ServiceDesc for Manifold service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manifold_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "manifold.Manifold",
	HandlerType: (*ManifoldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Enqueue",
			Handler:    _Manifold_Enqueue_Handler,
		},
		{
			MethodName: "Dequeue",
			Handler:    _Manifold_Dequeue_Handler,
		},
		{
			MethodName: "ListTopics",
			Handler:    _Manifold_ListTopics_Handler,
		},
		{
			MethodName: "DeleteTopic",
			Handler:    _Manifold_DeleteTopic_Handler,
		},
		{
			MethodName: "TopicLength",
			Handler:    _Manifold_TopicLength_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamDequeue",
			Handler:       _Manifold_StreamDequeue_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamEnqueue",
			Handler:       _Manifold_StreamEnqueue_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "manifold.proto",
}
