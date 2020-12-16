// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: github.com/dogmatiq/interopspec/transportspec/stream.proto

package transportspec

import (
	context "context"
	envelopespec "github.com/dogmatiq/interopspec/envelopespec"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ConsumeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ApplicationKey is the identity key of the application to consume from.
	ApplicationKey string `protobuf:"bytes,1,opt,name=application_key,json=applicationKey,proto3" json:"application_key,omitempty"`
	// Offset is the offset of the earliest message to be consumed.
	//
	// The offset of the message returned will be greater than this value if the
	// event at that offset is one of the requested message types.
	Offset uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// Types is the set of message types to include in the results.
	Types []string `protobuf:"bytes,3,rep,name=types,proto3" json:"types,omitempty"`
}

func (x *ConsumeRequest) Reset() {
	*x = ConsumeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_interopspec_transportspec_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumeRequest) ProtoMessage() {}

func (x *ConsumeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_interopspec_transportspec_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumeRequest.ProtoReflect.Descriptor instead.
func (*ConsumeRequest) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_interopspec_transportspec_stream_proto_rawDescGZIP(), []int{0}
}

func (x *ConsumeRequest) GetApplicationKey() string {
	if x != nil {
		return x.ApplicationKey
	}
	return ""
}

func (x *ConsumeRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ConsumeRequest) GetTypes() []string {
	if x != nil {
		return x.Types
	}
	return nil
}

type ConsumeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Offset is the offset of the message.
	//
	// It will be greater than the offset provided in the ConsumeRequest if the
	// event at the requested offset was one fo the requested message types.
	Offset uint64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	// Envelope is the envelope containing the event.
	Envelope *envelopespec.Envelope `protobuf:"bytes,2,opt,name=envelope,proto3" json:"envelope,omitempty"`
}

func (x *ConsumeResponse) Reset() {
	*x = ConsumeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_interopspec_transportspec_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumeResponse) ProtoMessage() {}

func (x *ConsumeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_interopspec_transportspec_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumeResponse.ProtoReflect.Descriptor instead.
func (*ConsumeResponse) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_interopspec_transportspec_stream_proto_rawDescGZIP(), []int{1}
}

func (x *ConsumeResponse) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ConsumeResponse) GetEnvelope() *envelopespec.Envelope {
	if x != nil {
		return x.Envelope
	}
	return nil
}

var File_github_com_dogmatiq_interopspec_transportspec_stream_proto protoreflect.FileDescriptor

var file_github_com_dogmatiq_interopspec_transportspec_stream_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67,
	0x6d, 0x61, 0x74, 0x69, 0x71, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x73, 0x70, 0x65,
	0x63, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x70, 0x65, 0x63, 0x2f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x64, 0x6f,
	0x67, 0x6d, 0x61, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x74, 0x69, 0x71, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6f, 0x70, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x65, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x74, 0x69, 0x71, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6f, 0x70, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x70, 0x65, 0x63, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b,
	0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x22, 0x6a, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x3f, 0x0a, 0x08, 0x65,
	0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x64, 0x6f, 0x67, 0x6d, 0x61, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31,
	0x2e, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x2e, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x65, 0x52, 0x08, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x32, 0xe4, 0x01, 0x0a,
	0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x64, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x12, 0x2a, 0x2e, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x30, 0x01, 0x12, 0x6f, 0x0a, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x12, 0x2f, 0x2e, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x30, 0x2e, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f,
	0x70, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x74, 0x69, 0x71, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6f, 0x70, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x70, 0x65, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_dogmatiq_interopspec_transportspec_stream_proto_rawDescOnce sync.Once
	file_github_com_dogmatiq_interopspec_transportspec_stream_proto_rawDescData = file_github_com_dogmatiq_interopspec_transportspec_stream_proto_rawDesc
)

func file_github_com_dogmatiq_interopspec_transportspec_stream_proto_rawDescGZIP() []byte {
	file_github_com_dogmatiq_interopspec_transportspec_stream_proto_rawDescOnce.Do(func() {
		file_github_com_dogmatiq_interopspec_transportspec_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_dogmatiq_interopspec_transportspec_stream_proto_rawDescData)
	})
	return file_github_com_dogmatiq_interopspec_transportspec_stream_proto_rawDescData
}

var file_github_com_dogmatiq_interopspec_transportspec_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_dogmatiq_interopspec_transportspec_stream_proto_goTypes = []interface{}{
	(*ConsumeRequest)(nil),        // 0: dogma.interop.v1.transport.ConsumeRequest
	(*ConsumeResponse)(nil),       // 1: dogma.interop.v1.transport.ConsumeResponse
	(*envelopespec.Envelope)(nil), // 2: dogma.interop.v1.envelope.Envelope
	(*MessageTypesRequest)(nil),   // 3: dogma.interop.v1.transport.MessageTypesRequest
	(*MessageTypesResponse)(nil),  // 4: dogma.interop.v1.transport.MessageTypesResponse
}
var file_github_com_dogmatiq_interopspec_transportspec_stream_proto_depIdxs = []int32{
	2, // 0: dogma.interop.v1.transport.ConsumeResponse.envelope:type_name -> dogma.interop.v1.envelope.Envelope
	0, // 1: dogma.interop.v1.transport.EventStream.Consume:input_type -> dogma.interop.v1.transport.ConsumeRequest
	3, // 2: dogma.interop.v1.transport.EventStream.EventTypes:input_type -> dogma.interop.v1.transport.MessageTypesRequest
	1, // 3: dogma.interop.v1.transport.EventStream.Consume:output_type -> dogma.interop.v1.transport.ConsumeResponse
	4, // 4: dogma.interop.v1.transport.EventStream.EventTypes:output_type -> dogma.interop.v1.transport.MessageTypesResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_github_com_dogmatiq_interopspec_transportspec_stream_proto_init() }
func file_github_com_dogmatiq_interopspec_transportspec_stream_proto_init() {
	if File_github_com_dogmatiq_interopspec_transportspec_stream_proto != nil {
		return
	}
	file_github_com_dogmatiq_interopspec_transportspec_messagetypes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_dogmatiq_interopspec_transportspec_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_dogmatiq_interopspec_transportspec_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_dogmatiq_interopspec_transportspec_stream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_dogmatiq_interopspec_transportspec_stream_proto_goTypes,
		DependencyIndexes: file_github_com_dogmatiq_interopspec_transportspec_stream_proto_depIdxs,
		MessageInfos:      file_github_com_dogmatiq_interopspec_transportspec_stream_proto_msgTypes,
	}.Build()
	File_github_com_dogmatiq_interopspec_transportspec_stream_proto = out.File
	file_github_com_dogmatiq_interopspec_transportspec_stream_proto_rawDesc = nil
	file_github_com_dogmatiq_interopspec_transportspec_stream_proto_goTypes = nil
	file_github_com_dogmatiq_interopspec_transportspec_stream_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EventStreamClient is the client API for EventStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventStreamClient interface {
	// Consume starts consuming from an application's event stream.
	//
	// If the server does not host the application specified in the request, it
	// MUST return a NOT_FOUND error with an attached UnrecognizedApplication
	// value.
	//
	// If the requested offset is beyond the end of the application's event
	// stream, the server SHOULD keep the stream open, sending the messages as
	// they become available.
	//
	// If any of the message types are not produced by the specified application
	// the server MUST return an INVALID_ARGUMENT error with an attached
	// UnrecognizedMessage value for each unrecognized message type.
	Consume(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (EventStream_ConsumeClient, error)
	// EventTypes queries the event types that the server supports for a specific
	// application.
	//
	// If the server does not host the application specified in the request, it
	// MUST return a NOT_FOUND error with an attached UnrecognizedApplication
	// value.
	EventTypes(ctx context.Context, in *MessageTypesRequest, opts ...grpc.CallOption) (*MessageTypesResponse, error)
}

type eventStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewEventStreamClient(cc grpc.ClientConnInterface) EventStreamClient {
	return &eventStreamClient{cc}
}

func (c *eventStreamClient) Consume(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (EventStream_ConsumeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EventStream_serviceDesc.Streams[0], "/dogma.interop.v1.transport.EventStream/Consume", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventStreamConsumeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventStream_ConsumeClient interface {
	Recv() (*ConsumeResponse, error)
	grpc.ClientStream
}

type eventStreamConsumeClient struct {
	grpc.ClientStream
}

func (x *eventStreamConsumeClient) Recv() (*ConsumeResponse, error) {
	m := new(ConsumeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventStreamClient) EventTypes(ctx context.Context, in *MessageTypesRequest, opts ...grpc.CallOption) (*MessageTypesResponse, error) {
	out := new(MessageTypesResponse)
	err := c.cc.Invoke(ctx, "/dogma.interop.v1.transport.EventStream/EventTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventStreamServer is the server API for EventStream service.
type EventStreamServer interface {
	// Consume starts consuming from an application's event stream.
	//
	// If the server does not host the application specified in the request, it
	// MUST return a NOT_FOUND error with an attached UnrecognizedApplication
	// value.
	//
	// If the requested offset is beyond the end of the application's event
	// stream, the server SHOULD keep the stream open, sending the messages as
	// they become available.
	//
	// If any of the message types are not produced by the specified application
	// the server MUST return an INVALID_ARGUMENT error with an attached
	// UnrecognizedMessage value for each unrecognized message type.
	Consume(*ConsumeRequest, EventStream_ConsumeServer) error
	// EventTypes queries the event types that the server supports for a specific
	// application.
	//
	// If the server does not host the application specified in the request, it
	// MUST return a NOT_FOUND error with an attached UnrecognizedApplication
	// value.
	EventTypes(context.Context, *MessageTypesRequest) (*MessageTypesResponse, error)
}

// UnimplementedEventStreamServer can be embedded to have forward compatible implementations.
type UnimplementedEventStreamServer struct {
}

func (*UnimplementedEventStreamServer) Consume(*ConsumeRequest, EventStream_ConsumeServer) error {
	return status.Errorf(codes.Unimplemented, "method Consume not implemented")
}
func (*UnimplementedEventStreamServer) EventTypes(context.Context, *MessageTypesRequest) (*MessageTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EventTypes not implemented")
}

func RegisterEventStreamServer(s *grpc.Server, srv EventStreamServer) {
	s.RegisterService(&_EventStream_serviceDesc, srv)
}

func _EventStream_Consume_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConsumeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventStreamServer).Consume(m, &eventStreamConsumeServer{stream})
}

type EventStream_ConsumeServer interface {
	Send(*ConsumeResponse) error
	grpc.ServerStream
}

type eventStreamConsumeServer struct {
	grpc.ServerStream
}

func (x *eventStreamConsumeServer) Send(m *ConsumeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _EventStream_EventTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventStreamServer).EventTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dogma.interop.v1.transport.EventStream/EventTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventStreamServer).EventTypes(ctx, req.(*MessageTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventStream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dogma.interop.v1.transport.EventStream",
	HandlerType: (*EventStreamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EventTypes",
			Handler:    _EventStream_EventTypes_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Consume",
			Handler:       _EventStream_Consume_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/dogmatiq/interopspec/transportspec/stream.proto",
}
