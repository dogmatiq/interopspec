// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: github.com/dogmatiq/interopspec/eventstreamspec/streamapi.proto

package eventstreamspec

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
	StreamAPI_Consume_FullMethodName    = "/dogma.interop.v1.eventstream.StreamAPI/Consume"
	StreamAPI_EventTypes_FullMethodName = "/dogma.interop.v1.eventstream.StreamAPI/EventTypes"
)

// StreamAPIClient is the client API for StreamAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamAPIClient interface {
	// Consume starts consuming from a particular "start point" within an
	// application's event stream.
	//
	// The start point can be specified using different mechanisms. If the server
	// does not support the requested mechanism it MUST return an UNIMPLEMENTED
	// error with an attached UnimplementedStartPointMechanism value.
	//
	// If the requested start point is beyond the end of the application's event
	// stream, the server SHOULD keep the stream open and send new events as they
	// are produced.
	//
	// If the start point is requested in a manner not supported by the server, it
	// MUST return an UNIMPLEMENTED error with an attached UnsupportedStartPoint
	// value.
	//
	// If the server does not host the application specified in the request, it
	// MUST return a NOT_FOUND error with an attached UnrecognizedApplication
	// value.
	//
	// If any of the requested event types are not produced by the specified
	// application the server MUST return an INVALID_ARGUMENT error with an
	// attached UnrecognizedEventType value for each unrecognized event type.
	//
	// If none of the requested media-types for a given event type are supported
	// the server MUST return an INVALID_ARGUMENT error with an attached
	// NoRecognizedMediaTypes value for each such event type.
	Consume(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (StreamAPI_ConsumeClient, error)
	// EventTypes queries the event types that the server supports for a specific
	// application.
	//
	// If the server does not host the application specified in the request, it
	// MUST return a NOT_FOUND error with an attached UnrecognizedApplication
	// value.
	EventTypes(ctx context.Context, in *EventTypesRequest, opts ...grpc.CallOption) (*EventTypesResponse, error)
}

type streamAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamAPIClient(cc grpc.ClientConnInterface) StreamAPIClient {
	return &streamAPIClient{cc}
}

func (c *streamAPIClient) Consume(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (StreamAPI_ConsumeClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamAPI_ServiceDesc.Streams[0], StreamAPI_Consume_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &streamAPIConsumeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamAPI_ConsumeClient interface {
	Recv() (*ConsumeResponse, error)
	grpc.ClientStream
}

type streamAPIConsumeClient struct {
	grpc.ClientStream
}

func (x *streamAPIConsumeClient) Recv() (*ConsumeResponse, error) {
	m := new(ConsumeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamAPIClient) EventTypes(ctx context.Context, in *EventTypesRequest, opts ...grpc.CallOption) (*EventTypesResponse, error) {
	out := new(EventTypesResponse)
	err := c.cc.Invoke(ctx, StreamAPI_EventTypes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreamAPIServer is the server API for StreamAPI service.
// All implementations should embed UnimplementedStreamAPIServer
// for forward compatibility
type StreamAPIServer interface {
	// Consume starts consuming from a particular "start point" within an
	// application's event stream.
	//
	// The start point can be specified using different mechanisms. If the server
	// does not support the requested mechanism it MUST return an UNIMPLEMENTED
	// error with an attached UnimplementedStartPointMechanism value.
	//
	// If the requested start point is beyond the end of the application's event
	// stream, the server SHOULD keep the stream open and send new events as they
	// are produced.
	//
	// If the start point is requested in a manner not supported by the server, it
	// MUST return an UNIMPLEMENTED error with an attached UnsupportedStartPoint
	// value.
	//
	// If the server does not host the application specified in the request, it
	// MUST return a NOT_FOUND error with an attached UnrecognizedApplication
	// value.
	//
	// If any of the requested event types are not produced by the specified
	// application the server MUST return an INVALID_ARGUMENT error with an
	// attached UnrecognizedEventType value for each unrecognized event type.
	//
	// If none of the requested media-types for a given event type are supported
	// the server MUST return an INVALID_ARGUMENT error with an attached
	// NoRecognizedMediaTypes value for each such event type.
	Consume(*ConsumeRequest, StreamAPI_ConsumeServer) error
	// EventTypes queries the event types that the server supports for a specific
	// application.
	//
	// If the server does not host the application specified in the request, it
	// MUST return a NOT_FOUND error with an attached UnrecognizedApplication
	// value.
	EventTypes(context.Context, *EventTypesRequest) (*EventTypesResponse, error)
}

// UnimplementedStreamAPIServer should be embedded to have forward compatible implementations.
type UnimplementedStreamAPIServer struct {
}

func (UnimplementedStreamAPIServer) Consume(*ConsumeRequest, StreamAPI_ConsumeServer) error {
	return status.Errorf(codes.Unimplemented, "method Consume not implemented")
}
func (UnimplementedStreamAPIServer) EventTypes(context.Context, *EventTypesRequest) (*EventTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EventTypes not implemented")
}

// UnsafeStreamAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamAPIServer will
// result in compilation errors.
type UnsafeStreamAPIServer interface {
	mustEmbedUnimplementedStreamAPIServer()
}

func RegisterStreamAPIServer(s grpc.ServiceRegistrar, srv StreamAPIServer) {
	s.RegisterService(&StreamAPI_ServiceDesc, srv)
}

func _StreamAPI_Consume_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConsumeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamAPIServer).Consume(m, &streamAPIConsumeServer{stream})
}

type StreamAPI_ConsumeServer interface {
	Send(*ConsumeResponse) error
	grpc.ServerStream
}

type streamAPIConsumeServer struct {
	grpc.ServerStream
}

func (x *streamAPIConsumeServer) Send(m *ConsumeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _StreamAPI_EventTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamAPIServer).EventTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamAPI_EventTypes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamAPIServer).EventTypes(ctx, req.(*EventTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StreamAPI_ServiceDesc is the grpc.ServiceDesc for StreamAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dogma.interop.v1.eventstream.StreamAPI",
	HandlerType: (*StreamAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EventTypes",
			Handler:    _StreamAPI_EventTypes_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Consume",
			Handler:       _StreamAPI_Consume_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/dogmatiq/interopspec/eventstreamspec/streamapi.proto",
}
