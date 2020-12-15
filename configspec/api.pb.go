// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: github.com/dogmatiq/interopspec/configspec/api.proto

package configspec

import (
	context "context"
	interopspec "github.com/dogmatiq/interopspec"
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

// MessageRole is a protocol buffers representation of the
// configkit/message.Role enumeration.
type MessageRole int32

const (
	MessageRole_UNKNOWN_MESSAGE_ROLE MessageRole = 0
	MessageRole_COMMAND              MessageRole = 1
	MessageRole_EVENT                MessageRole = 2
	MessageRole_TIMEOUT              MessageRole = 3
)

// Enum value maps for MessageRole.
var (
	MessageRole_name = map[int32]string{
		0: "UNKNOWN_MESSAGE_ROLE",
		1: "COMMAND",
		2: "EVENT",
		3: "TIMEOUT",
	}
	MessageRole_value = map[string]int32{
		"UNKNOWN_MESSAGE_ROLE": 0,
		"COMMAND":              1,
		"EVENT":                2,
		"TIMEOUT":              3,
	}
)

func (x MessageRole) Enum() *MessageRole {
	p := new(MessageRole)
	*p = x
	return p
}

func (x MessageRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageRole) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_dogmatiq_interopspec_configspec_api_proto_enumTypes[0].Descriptor()
}

func (MessageRole) Type() protoreflect.EnumType {
	return &file_github_com_dogmatiq_interopspec_configspec_api_proto_enumTypes[0]
}

func (x MessageRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageRole.Descriptor instead.
func (MessageRole) EnumDescriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_interopspec_configspec_api_proto_rawDescGZIP(), []int{0}
}

// HandlerType is a protocol buffers representation of the configkit.HandlerType
// enumeration.
type HandlerType int32

const (
	HandlerType_UNKNOWN_HANDLER_TYPE HandlerType = 0
	HandlerType_AGGREGATE            HandlerType = 1
	HandlerType_PROCESS              HandlerType = 2
	HandlerType_INTEGRATION          HandlerType = 3
	HandlerType_PROJECTION           HandlerType = 4
)

// Enum value maps for HandlerType.
var (
	HandlerType_name = map[int32]string{
		0: "UNKNOWN_HANDLER_TYPE",
		1: "AGGREGATE",
		2: "PROCESS",
		3: "INTEGRATION",
		4: "PROJECTION",
	}
	HandlerType_value = map[string]int32{
		"UNKNOWN_HANDLER_TYPE": 0,
		"AGGREGATE":            1,
		"PROCESS":              2,
		"INTEGRATION":          3,
		"PROJECTION":           4,
	}
)

func (x HandlerType) Enum() *HandlerType {
	p := new(HandlerType)
	*p = x
	return p
}

func (x HandlerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HandlerType) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_dogmatiq_interopspec_configspec_api_proto_enumTypes[1].Descriptor()
}

func (HandlerType) Type() protoreflect.EnumType {
	return &file_github_com_dogmatiq_interopspec_configspec_api_proto_enumTypes[1]
}

func (x HandlerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HandlerType.Descriptor instead.
func (HandlerType) EnumDescriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_interopspec_configspec_api_proto_rawDescGZIP(), []int{1}
}

type ListApplicationIdentitiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListApplicationIdentitiesRequest) Reset() {
	*x = ListApplicationIdentitiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_interopspec_configspec_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListApplicationIdentitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApplicationIdentitiesRequest) ProtoMessage() {}

func (x *ListApplicationIdentitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_interopspec_configspec_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApplicationIdentitiesRequest.ProtoReflect.Descriptor instead.
func (*ListApplicationIdentitiesRequest) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_interopspec_configspec_api_proto_rawDescGZIP(), []int{0}
}

type ListApplicationIdentitiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identities []*interopspec.Identity `protobuf:"bytes,1,rep,name=identities,proto3" json:"identities,omitempty"`
}

func (x *ListApplicationIdentitiesResponse) Reset() {
	*x = ListApplicationIdentitiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_interopspec_configspec_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListApplicationIdentitiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApplicationIdentitiesResponse) ProtoMessage() {}

func (x *ListApplicationIdentitiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_interopspec_configspec_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApplicationIdentitiesResponse.ProtoReflect.Descriptor instead.
func (*ListApplicationIdentitiesResponse) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_interopspec_configspec_api_proto_rawDescGZIP(), []int{1}
}

func (x *ListApplicationIdentitiesResponse) GetIdentities() []*interopspec.Identity {
	if x != nil {
		return x.Identities
	}
	return nil
}

type ListApplicationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListApplicationsRequest) Reset() {
	*x = ListApplicationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_interopspec_configspec_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListApplicationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApplicationsRequest) ProtoMessage() {}

func (x *ListApplicationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_interopspec_configspec_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApplicationsRequest.ProtoReflect.Descriptor instead.
func (*ListApplicationsRequest) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_interopspec_configspec_api_proto_rawDescGZIP(), []int{2}
}

type ListApplicationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Applications []*Application `protobuf:"bytes,1,rep,name=applications,proto3" json:"applications,omitempty"`
}

func (x *ListApplicationsResponse) Reset() {
	*x = ListApplicationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_interopspec_configspec_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListApplicationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApplicationsResponse) ProtoMessage() {}

func (x *ListApplicationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_interopspec_configspec_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApplicationsResponse.ProtoReflect.Descriptor instead.
func (*ListApplicationsResponse) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_interopspec_configspec_api_proto_rawDescGZIP(), []int{3}
}

func (x *ListApplicationsResponse) GetApplications() []*Application {
	if x != nil {
		return x.Applications
	}
	return nil
}

// Application is a protocol buffers representation of the configkit.Application
// interface.
type Application struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identity is the application's identity.
	Identity *interopspec.Identity `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	// TypeName is the fully-qualified name of the Go type used to implement the
	// application.
	TypeName string `protobuf:"bytes,2,opt,name=type_name,json=typeName,proto3" json:"type_name,omitempty"`
	// Messages is an ordered-sequence of message name / role pairs.
	//
	// This directly correlates to the configkit.Application.MessageNames().Roles
	// value. The produced/consumed message names are not encoded directly in the
	// application, but rather rebuilt from the handlers when the application is
	// unmarshaled.
	Messages []*NameRole `protobuf:"bytes,3,rep,name=messages,proto3" json:"messages,omitempty"`
	// Handlers is the set of handlers within the application.
	Handlers []*Handler `protobuf:"bytes,4,rep,name=handlers,proto3" json:"handlers,omitempty"`
}

func (x *Application) Reset() {
	*x = Application{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_interopspec_configspec_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application) ProtoMessage() {}

func (x *Application) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_interopspec_configspec_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application.ProtoReflect.Descriptor instead.
func (*Application) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_interopspec_configspec_api_proto_rawDescGZIP(), []int{4}
}

func (x *Application) GetIdentity() *interopspec.Identity {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *Application) GetTypeName() string {
	if x != nil {
		return x.TypeName
	}
	return ""
}

func (x *Application) GetMessages() []*NameRole {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *Application) GetHandlers() []*Handler {
	if x != nil {
		return x.Handlers
	}
	return nil
}

// NameRole is a 2-tuple containing a message name and its role.
type NameRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is the fully-qualified message name.
	Name []byte `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Role is the role this message plays within the application.
	Role MessageRole `protobuf:"varint,2,opt,name=role,proto3,enum=dogma.interop.v1.config.MessageRole" json:"role,omitempty"`
}

func (x *NameRole) Reset() {
	*x = NameRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_interopspec_configspec_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameRole) ProtoMessage() {}

func (x *NameRole) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_interopspec_configspec_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameRole.ProtoReflect.Descriptor instead.
func (*NameRole) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_interopspec_configspec_api_proto_rawDescGZIP(), []int{5}
}

func (x *NameRole) GetName() []byte {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *NameRole) GetRole() MessageRole {
	if x != nil {
		return x.Role
	}
	return MessageRole_UNKNOWN_MESSAGE_ROLE
}

// Handler is a protocol buffers representation of the configkit.Handler
// interface.
type Handler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identity is the handler's identity.
	Identity *interopspec.Identity `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	// TypeName is the fully-qualified name of the Go type used to implement the
	// handler.
	TypeName string `protobuf:"bytes,2,opt,name=type_name,json=typeName,proto3" json:"type_name,omitempty"`
	// Type is the handler's type.
	Type HandlerType `protobuf:"varint,3,opt,name=type,proto3,enum=dogma.interop.v1.config.HandlerType" json:"type,omitempty"`
	// Produced is a list of the messages produced by this handler.
	// Each value is the index of a MessageRolePair within the application.
	Produced []uint32 `protobuf:"varint,4,rep,packed,name=produced,proto3" json:"produced,omitempty"`
	// Consumed is a list of the messages consumed by this handler.
	// Each value is the index of a MessageRolePair within the application.
	Consumed []uint32 `protobuf:"varint,5,rep,packed,name=consumed,proto3" json:"consumed,omitempty"`
}

func (x *Handler) Reset() {
	*x = Handler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_interopspec_configspec_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Handler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Handler) ProtoMessage() {}

func (x *Handler) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_interopspec_configspec_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Handler.ProtoReflect.Descriptor instead.
func (*Handler) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_interopspec_configspec_api_proto_rawDescGZIP(), []int{6}
}

func (x *Handler) GetIdentity() *interopspec.Identity {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *Handler) GetTypeName() string {
	if x != nil {
		return x.TypeName
	}
	return ""
}

func (x *Handler) GetType() HandlerType {
	if x != nil {
		return x.Type
	}
	return HandlerType_UNKNOWN_HANDLER_TYPE
}

func (x *Handler) GetProduced() []uint32 {
	if x != nil {
		return x.Produced
	}
	return nil
}

func (x *Handler) GetConsumed() []uint32 {
	if x != nil {
		return x.Consumed
	}
	return nil
}

var File_github_com_dogmatiq_interopspec_configspec_api_proto protoreflect.FileDescriptor

var file_github_com_dogmatiq_interopspec_configspec_api_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67,
	0x6d, 0x61, 0x74, 0x69, 0x71, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x73, 0x70, 0x65,
	0x63, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67, 0x6d,
	0x61, 0x74, 0x69, 0x71, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x73, 0x70, 0x65, 0x63,
	0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x22, 0x0a, 0x20, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x5f, 0x0a, 0x21, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64,
	0x6f, 0x67, 0x6d, 0x61, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x22, 0x19, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x64, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0c, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f,
	0x70, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xdf, 0x01, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a,
	0x09, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x64,
	0x6f, 0x67, 0x6d, 0x61, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x08, 0x68, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64, 0x6f,
	0x67, 0x6d, 0x61, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x08, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x22, 0x58, 0x0a, 0x08, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x6f, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x22, 0xd0, 0x01, 0x0a, 0x07, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x36, 0x0a,
	0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x79, 0x70, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x24, 0x2e, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x64, 0x2a, 0x4c, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x6f, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x4d,
	0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x56,
	0x45, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54,
	0x10, 0x03, 0x2a, 0x64, 0x0a, 0x0b, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x48, 0x41, 0x4e,
	0x44, 0x4c, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x41,
	0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52,
	0x4f, 0x43, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x54, 0x45, 0x47,
	0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x52, 0x4f, 0x4a,
	0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x32, 0x93, 0x02, 0x0a, 0x03, 0x41, 0x50, 0x49,
	0x12, 0x92, 0x01, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x39,
	0x2e, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x64, 0x6f, 0x67, 0x6d,
	0x61, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x77, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x30, 0x2e, 0x64, 0x6f, 0x67, 0x6d,
	0x61, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x64, 0x6f,
	0x67, 0x6d, 0x61, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2c,
	0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67,
	0x6d, 0x61, 0x74, 0x69, 0x71, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x73, 0x70, 0x65,
	0x63, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x70, 0x65, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_dogmatiq_interopspec_configspec_api_proto_rawDescOnce sync.Once
	file_github_com_dogmatiq_interopspec_configspec_api_proto_rawDescData = file_github_com_dogmatiq_interopspec_configspec_api_proto_rawDesc
)

func file_github_com_dogmatiq_interopspec_configspec_api_proto_rawDescGZIP() []byte {
	file_github_com_dogmatiq_interopspec_configspec_api_proto_rawDescOnce.Do(func() {
		file_github_com_dogmatiq_interopspec_configspec_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_dogmatiq_interopspec_configspec_api_proto_rawDescData)
	})
	return file_github_com_dogmatiq_interopspec_configspec_api_proto_rawDescData
}

var file_github_com_dogmatiq_interopspec_configspec_api_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_github_com_dogmatiq_interopspec_configspec_api_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_github_com_dogmatiq_interopspec_configspec_api_proto_goTypes = []interface{}{
	(MessageRole)(0),                          // 0: dogma.interop.v1.config.MessageRole
	(HandlerType)(0),                          // 1: dogma.interop.v1.config.HandlerType
	(*ListApplicationIdentitiesRequest)(nil),  // 2: dogma.interop.v1.config.ListApplicationIdentitiesRequest
	(*ListApplicationIdentitiesResponse)(nil), // 3: dogma.interop.v1.config.ListApplicationIdentitiesResponse
	(*ListApplicationsRequest)(nil),           // 4: dogma.interop.v1.config.ListApplicationsRequest
	(*ListApplicationsResponse)(nil),          // 5: dogma.interop.v1.config.ListApplicationsResponse
	(*Application)(nil),                       // 6: dogma.interop.v1.config.Application
	(*NameRole)(nil),                          // 7: dogma.interop.v1.config.NameRole
	(*Handler)(nil),                           // 8: dogma.interop.v1.config.Handler
	(*interopspec.Identity)(nil),              // 9: dogma.interop.v1.Identity
}
var file_github_com_dogmatiq_interopspec_configspec_api_proto_depIdxs = []int32{
	9,  // 0: dogma.interop.v1.config.ListApplicationIdentitiesResponse.identities:type_name -> dogma.interop.v1.Identity
	6,  // 1: dogma.interop.v1.config.ListApplicationsResponse.applications:type_name -> dogma.interop.v1.config.Application
	9,  // 2: dogma.interop.v1.config.Application.identity:type_name -> dogma.interop.v1.Identity
	7,  // 3: dogma.interop.v1.config.Application.messages:type_name -> dogma.interop.v1.config.NameRole
	8,  // 4: dogma.interop.v1.config.Application.handlers:type_name -> dogma.interop.v1.config.Handler
	0,  // 5: dogma.interop.v1.config.NameRole.role:type_name -> dogma.interop.v1.config.MessageRole
	9,  // 6: dogma.interop.v1.config.Handler.identity:type_name -> dogma.interop.v1.Identity
	1,  // 7: dogma.interop.v1.config.Handler.type:type_name -> dogma.interop.v1.config.HandlerType
	2,  // 8: dogma.interop.v1.config.API.ListApplicationIdentities:input_type -> dogma.interop.v1.config.ListApplicationIdentitiesRequest
	4,  // 9: dogma.interop.v1.config.API.ListApplications:input_type -> dogma.interop.v1.config.ListApplicationsRequest
	3,  // 10: dogma.interop.v1.config.API.ListApplicationIdentities:output_type -> dogma.interop.v1.config.ListApplicationIdentitiesResponse
	5,  // 11: dogma.interop.v1.config.API.ListApplications:output_type -> dogma.interop.v1.config.ListApplicationsResponse
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_github_com_dogmatiq_interopspec_configspec_api_proto_init() }
func file_github_com_dogmatiq_interopspec_configspec_api_proto_init() {
	if File_github_com_dogmatiq_interopspec_configspec_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_dogmatiq_interopspec_configspec_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListApplicationIdentitiesRequest); i {
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
		file_github_com_dogmatiq_interopspec_configspec_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListApplicationIdentitiesResponse); i {
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
		file_github_com_dogmatiq_interopspec_configspec_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListApplicationsRequest); i {
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
		file_github_com_dogmatiq_interopspec_configspec_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListApplicationsResponse); i {
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
		file_github_com_dogmatiq_interopspec_configspec_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application); i {
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
		file_github_com_dogmatiq_interopspec_configspec_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameRole); i {
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
		file_github_com_dogmatiq_interopspec_configspec_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Handler); i {
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
			RawDescriptor: file_github_com_dogmatiq_interopspec_configspec_api_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_dogmatiq_interopspec_configspec_api_proto_goTypes,
		DependencyIndexes: file_github_com_dogmatiq_interopspec_configspec_api_proto_depIdxs,
		EnumInfos:         file_github_com_dogmatiq_interopspec_configspec_api_proto_enumTypes,
		MessageInfos:      file_github_com_dogmatiq_interopspec_configspec_api_proto_msgTypes,
	}.Build()
	File_github_com_dogmatiq_interopspec_configspec_api_proto = out.File
	file_github_com_dogmatiq_interopspec_configspec_api_proto_rawDesc = nil
	file_github_com_dogmatiq_interopspec_configspec_api_proto_goTypes = nil
	file_github_com_dogmatiq_interopspec_configspec_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type APIClient interface {
	// ListApplicationIdentities returns the identity of all applications.
	ListApplicationIdentities(ctx context.Context, in *ListApplicationIdentitiesRequest, opts ...grpc.CallOption) (*ListApplicationIdentitiesResponse, error)
	// ListApplications returns the full configuration of all applications.
	ListApplications(ctx context.Context, in *ListApplicationsRequest, opts ...grpc.CallOption) (*ListApplicationsResponse, error)
}

type aPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAPIClient(cc grpc.ClientConnInterface) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) ListApplicationIdentities(ctx context.Context, in *ListApplicationIdentitiesRequest, opts ...grpc.CallOption) (*ListApplicationIdentitiesResponse, error) {
	out := new(ListApplicationIdentitiesResponse)
	err := c.cc.Invoke(ctx, "/dogma.interop.v1.config.API/ListApplicationIdentities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListApplications(ctx context.Context, in *ListApplicationsRequest, opts ...grpc.CallOption) (*ListApplicationsResponse, error) {
	out := new(ListApplicationsResponse)
	err := c.cc.Invoke(ctx, "/dogma.interop.v1.config.API/ListApplications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServer is the server API for API service.
type APIServer interface {
	// ListApplicationIdentities returns the identity of all applications.
	ListApplicationIdentities(context.Context, *ListApplicationIdentitiesRequest) (*ListApplicationIdentitiesResponse, error)
	// ListApplications returns the full configuration of all applications.
	ListApplications(context.Context, *ListApplicationsRequest) (*ListApplicationsResponse, error)
}

// UnimplementedAPIServer can be embedded to have forward compatible implementations.
type UnimplementedAPIServer struct {
}

func (*UnimplementedAPIServer) ListApplicationIdentities(context.Context, *ListApplicationIdentitiesRequest) (*ListApplicationIdentitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListApplicationIdentities not implemented")
}
func (*UnimplementedAPIServer) ListApplications(context.Context, *ListApplicationsRequest) (*ListApplicationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListApplications not implemented")
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_ListApplicationIdentities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApplicationIdentitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ListApplicationIdentities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dogma.interop.v1.config.API/ListApplicationIdentities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ListApplicationIdentities(ctx, req.(*ListApplicationIdentitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListApplications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApplicationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ListApplications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dogma.interop.v1.config.API/ListApplications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ListApplications(ctx, req.(*ListApplicationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dogma.interop.v1.config.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListApplicationIdentities",
			Handler:    _API_ListApplicationIdentities_Handler,
		},
		{
			MethodName: "ListApplications",
			Handler:    _API_ListApplications_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/dogmatiq/interopspec/configspec/api.proto",
}
