// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: github.com/dogmatiq/interopspec/eventstreamspec/streamapi.proto

package eventstreamspec

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
	// EventTypes is a list of event types to be consumed.
	EventTypes []*EventType `protobuf:"bytes,2,rep,name=event_types,json=eventTypes,proto3" json:"event_types,omitempty"`
	// StartPoint specifies the point within the stream to start consuming.
	//
	// Types that are assignable to StartPoint:
	//	*ConsumeRequest_Offset
	StartPoint isConsumeRequest_StartPoint `protobuf_oneof:"start_point"`
}

func (x *ConsumeRequest) Reset() {
	*x = ConsumeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumeRequest) ProtoMessage() {}

func (x *ConsumeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_msgTypes[0]
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
	return file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_rawDescGZIP(), []int{0}
}

func (x *ConsumeRequest) GetApplicationKey() string {
	if x != nil {
		return x.ApplicationKey
	}
	return ""
}

func (x *ConsumeRequest) GetEventTypes() []*EventType {
	if x != nil {
		return x.EventTypes
	}
	return nil
}

func (m *ConsumeRequest) GetStartPoint() isConsumeRequest_StartPoint {
	if m != nil {
		return m.StartPoint
	}
	return nil
}

func (x *ConsumeRequest) GetOffset() uint64 {
	if x, ok := x.GetStartPoint().(*ConsumeRequest_Offset); ok {
		return x.Offset
	}
	return 0
}

type isConsumeRequest_StartPoint interface {
	isConsumeRequest_StartPoint()
}

type ConsumeRequest_Offset struct {
	// Offset is the offset of the earliest message to be consumed.
	//
	// The offset of the message returned will be greater than this value if the
	// event at that offset is not one of the requested event types.
	Offset uint64 `protobuf:"varint,3,opt,name=offset,proto3,oneof"`
}

func (*ConsumeRequest_Offset) isConsumeRequest_StartPoint() {}

type ConsumeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Offset is the offset of the message.
	Offset uint64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	// Envelope is the envelope containing the event.
	Envelope *envelopespec.Envelope `protobuf:"bytes,2,opt,name=envelope,proto3" json:"envelope,omitempty"`
}

func (x *ConsumeResponse) Reset() {
	*x = ConsumeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumeResponse) ProtoMessage() {}

func (x *ConsumeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_msgTypes[1]
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
	return file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_rawDescGZIP(), []int{1}
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

type EventTypesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ApplicationKey is the identity key of the application to query.
	ApplicationKey string `protobuf:"bytes,1,opt,name=application_key,json=applicationKey,proto3" json:"application_key,omitempty"`
}

func (x *EventTypesRequest) Reset() {
	*x = EventTypesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventTypesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventTypesRequest) ProtoMessage() {}

func (x *EventTypesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventTypesRequest.ProtoReflect.Descriptor instead.
func (*EventTypesRequest) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_rawDescGZIP(), []int{2}
}

func (x *EventTypesRequest) GetApplicationKey() string {
	if x != nil {
		return x.ApplicationKey
	}
	return ""
}

type EventTypesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// EventTypes is a list of event types supported by the server.
	EventTypes []*EventType `protobuf:"bytes,1,rep,name=event_types,json=eventTypes,proto3" json:"event_types,omitempty"`
}

func (x *EventTypesResponse) Reset() {
	*x = EventTypesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventTypesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventTypesResponse) ProtoMessage() {}

func (x *EventTypesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventTypesResponse.ProtoReflect.Descriptor instead.
func (*EventTypesResponse) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_rawDescGZIP(), []int{3}
}

func (x *EventTypesResponse) GetEventTypes() []*EventType {
	if x != nil {
		return x.EventTypes
	}
	return nil
}

type EventType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PortableName is a name that uniquely identifies the event type across
	// process boundaries.
	PortableName string `protobuf:"bytes,1,opt,name=portable_name,json=portableName,proto3" json:"portable_name,omitempty"`
	// MediaTypes is the set of supported media-types that can be used to
	// represents events of this type, in order of preference.
	MediaTypes []string `protobuf:"bytes,2,rep,name=media_types,json=mediaTypes,proto3" json:"media_types,omitempty"`
}

func (x *EventType) Reset() {
	*x = EventType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventType) ProtoMessage() {}

func (x *EventType) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventType.ProtoReflect.Descriptor instead.
func (*EventType) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_rawDescGZIP(), []int{4}
}

func (x *EventType) GetPortableName() string {
	if x != nil {
		return x.PortableName
	}
	return ""
}

func (x *EventType) GetMediaTypes() []string {
	if x != nil {
		return x.MediaTypes
	}
	return nil
}

// UnrecognizedApplication is an error-details value for INVALID_ARGUMENT
// errors that occurred because a specific application key was not recognized by
// the server.
type UnrecognizedApplication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ApplicationKey is the identity of the application that produced the error.
	ApplicationKey string `protobuf:"bytes,1,opt,name=application_key,json=applicationKey,proto3" json:"application_key,omitempty"`
}

func (x *UnrecognizedApplication) Reset() {
	*x = UnrecognizedApplication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnrecognizedApplication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnrecognizedApplication) ProtoMessage() {}

func (x *UnrecognizedApplication) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnrecognizedApplication.ProtoReflect.Descriptor instead.
func (*UnrecognizedApplication) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_rawDescGZIP(), []int{5}
}

func (x *UnrecognizedApplication) GetApplicationKey() string {
	if x != nil {
		return x.ApplicationKey
	}
	return ""
}

// UnrecognizedEventType is an error-details value for INVALID_ARGUMENT errors
// that occurred because a specific event type was not recognized by the
// server.
type UnrecognizedEventType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PortableName is a name that uniquely identifies the event type across
	// process boundaries.
	PortableName string `protobuf:"bytes,2,opt,name=portable_name,json=portableName,proto3" json:"portable_name,omitempty"`
}

func (x *UnrecognizedEventType) Reset() {
	*x = UnrecognizedEventType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnrecognizedEventType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnrecognizedEventType) ProtoMessage() {}

func (x *UnrecognizedEventType) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnrecognizedEventType.ProtoReflect.Descriptor instead.
func (*UnrecognizedEventType) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_rawDescGZIP(), []int{6}
}

func (x *UnrecognizedEventType) GetPortableName() string {
	if x != nil {
		return x.PortableName
	}
	return ""
}

// NoRecognizedMediaTypes is an error-details value for INVALID_ARGUMENT errors
// that occurred because a client requested that events of a specific type be
// encoded using media-types that are not recognized by the server.
type NoRecognizedMediaTypes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PortableName is a name that uniquely identifies the event type across
	// process boundaries.
	PortableName string `protobuf:"bytes,2,opt,name=portable_name,json=portableName,proto3" json:"portable_name,omitempty"`
}

func (x *NoRecognizedMediaTypes) Reset() {
	*x = NoRecognizedMediaTypes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoRecognizedMediaTypes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoRecognizedMediaTypes) ProtoMessage() {}

func (x *NoRecognizedMediaTypes) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoRecognizedMediaTypes.ProtoReflect.Descriptor instead.
func (*NoRecognizedMediaTypes) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_rawDescGZIP(), []int{7}
}

func (x *NoRecognizedMediaTypes) GetPortableName() string {
	if x != nil {
		return x.PortableName
	}
	return ""
}

// UnsupportedStartPoint is an error-details value for UNIMPLEMENTED errors that
// occurred because a client started consuming using a start point specified in
// a way that is not supported by the server.
type UnsupportedStartPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnsupportedStartPoint) Reset() {
	*x = UnsupportedStartPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsupportedStartPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsupportedStartPoint) ProtoMessage() {}

func (x *UnsupportedStartPoint) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsupportedStartPoint.ProtoReflect.Descriptor instead.
func (*UnsupportedStartPoint) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_rawDescGZIP(), []int{8}
}

var File_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto protoreflect.FileDescriptor

var file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67,
	0x6d, 0x61, 0x74, 0x69, 0x71, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x73, 0x70, 0x65,
	0x63, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x70, 0x65,
	0x63, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1c, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x1a,
	0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67, 0x6d,
	0x61, 0x74, 0x69, 0x71, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x73, 0x70, 0x65, 0x63,
	0x2f, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x65, 0x6e,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x01, 0x0a,
	0x0e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x27, 0x0a, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x48, 0x0a, 0x0b, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x64, 0x6f, 0x67, 0x6d, 0x61, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x12, 0x18, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x42, 0x0d, 0x0a, 0x0b,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x6a, 0x0a, 0x0f, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x3f, 0x0a, 0x08, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x64, 0x6f, 0x67, 0x6d, 0x61,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x76, 0x65,
	0x6c, 0x6f, 0x70, 0x65, 0x2e, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x52, 0x08, 0x65,
	0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x22, 0x3c, 0x0a, 0x11, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x22, 0x5e, 0x0a, 0x12, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x22, 0x51, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x6f, 0x72, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x73, 0x22, 0x42, 0x0a, 0x17, 0x55, 0x6e, 0x72, 0x65,
	0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x64, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x22, 0x3c, 0x0a, 0x15,
	0x55, 0x6e, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x6f,
	0x72, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3d, 0x0a, 0x16, 0x4e, 0x6f,
	0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x64, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x6f, 0x72,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x55, 0x6e, 0x73,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x32, 0xe6, 0x01, 0x0a, 0x09, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x50, 0x49,
	0x12, 0x68, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x12, 0x2c, 0x2e, 0x64, 0x6f,
	0x67, 0x6d, 0x61, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x64, 0x6f, 0x67, 0x6d,
	0x61, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x6f, 0x0a, 0x0a, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x2f, 0x2e, 0x64, 0x6f, 0x67, 0x6d, 0x61,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x64, 0x6f, 0x67, 0x6d,
	0x61, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x74,
	0x69, 0x71, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x70, 0x65, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_rawDescOnce sync.Once
	file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_rawDescData = file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_rawDesc
)

func file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_rawDescGZIP() []byte {
	file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_rawDescOnce.Do(func() {
		file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_rawDescData)
	})
	return file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_rawDescData
}

var file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_goTypes = []interface{}{
	(*ConsumeRequest)(nil),          // 0: dogma.interop.v1.eventstream.ConsumeRequest
	(*ConsumeResponse)(nil),         // 1: dogma.interop.v1.eventstream.ConsumeResponse
	(*EventTypesRequest)(nil),       // 2: dogma.interop.v1.eventstream.EventTypesRequest
	(*EventTypesResponse)(nil),      // 3: dogma.interop.v1.eventstream.EventTypesResponse
	(*EventType)(nil),               // 4: dogma.interop.v1.eventstream.EventType
	(*UnrecognizedApplication)(nil), // 5: dogma.interop.v1.eventstream.UnrecognizedApplication
	(*UnrecognizedEventType)(nil),   // 6: dogma.interop.v1.eventstream.UnrecognizedEventType
	(*NoRecognizedMediaTypes)(nil),  // 7: dogma.interop.v1.eventstream.NoRecognizedMediaTypes
	(*UnsupportedStartPoint)(nil),   // 8: dogma.interop.v1.eventstream.UnsupportedStartPoint
	(*envelopespec.Envelope)(nil),   // 9: dogma.interop.v1.envelope.Envelope
}
var file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_depIdxs = []int32{
	4, // 0: dogma.interop.v1.eventstream.ConsumeRequest.event_types:type_name -> dogma.interop.v1.eventstream.EventType
	9, // 1: dogma.interop.v1.eventstream.ConsumeResponse.envelope:type_name -> dogma.interop.v1.envelope.Envelope
	4, // 2: dogma.interop.v1.eventstream.EventTypesResponse.event_types:type_name -> dogma.interop.v1.eventstream.EventType
	0, // 3: dogma.interop.v1.eventstream.StreamAPI.Consume:input_type -> dogma.interop.v1.eventstream.ConsumeRequest
	2, // 4: dogma.interop.v1.eventstream.StreamAPI.EventTypes:input_type -> dogma.interop.v1.eventstream.EventTypesRequest
	1, // 5: dogma.interop.v1.eventstream.StreamAPI.Consume:output_type -> dogma.interop.v1.eventstream.ConsumeResponse
	3, // 6: dogma.interop.v1.eventstream.StreamAPI.EventTypes:output_type -> dogma.interop.v1.eventstream.EventTypesResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_init() }
func file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_init() {
	if File_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventTypesRequest); i {
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
		file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventTypesResponse); i {
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
		file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventType); i {
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
		file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnrecognizedApplication); i {
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
		file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnrecognizedEventType); i {
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
		file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoRecognizedMediaTypes); i {
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
		file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsupportedStartPoint); i {
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
	file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ConsumeRequest_Offset)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_goTypes,
		DependencyIndexes: file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_depIdxs,
		MessageInfos:      file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_msgTypes,
	}.Build()
	File_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto = out.File
	file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_rawDesc = nil
	file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_goTypes = nil
	file_github_com_dogmatiq_interopspec_eventstreamspec_streamapi_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StreamAPIClient is the client API for StreamAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
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
	stream, err := c.cc.NewStream(ctx, &_StreamAPI_serviceDesc.Streams[0], "/dogma.interop.v1.eventstream.StreamAPI/Consume", opts...)
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
	err := c.cc.Invoke(ctx, "/dogma.interop.v1.eventstream.StreamAPI/EventTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreamAPIServer is the server API for StreamAPI service.
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

// UnimplementedStreamAPIServer can be embedded to have forward compatible implementations.
type UnimplementedStreamAPIServer struct {
}

func (*UnimplementedStreamAPIServer) Consume(*ConsumeRequest, StreamAPI_ConsumeServer) error {
	return status.Errorf(codes.Unimplemented, "method Consume not implemented")
}
func (*UnimplementedStreamAPIServer) EventTypes(context.Context, *EventTypesRequest) (*EventTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EventTypes not implemented")
}

func RegisterStreamAPIServer(s *grpc.Server, srv StreamAPIServer) {
	s.RegisterService(&_StreamAPI_serviceDesc, srv)
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
		FullMethod: "/dogma.interop.v1.eventstream.StreamAPI/EventTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamAPIServer).EventTypes(ctx, req.(*EventTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StreamAPI_serviceDesc = grpc.ServiceDesc{
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
