// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: github.com/dogmatiq/interopspec/discoverspec/discoverapi.proto

package discoverspec

import (
	context "context"
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

type ListApplicationIdentitiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListApplicationIdentitiesRequest) Reset() {
	*x = ListApplicationIdentitiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListApplicationIdentitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApplicationIdentitiesRequest) ProtoMessage() {}

func (x *ListApplicationIdentitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_msgTypes[0]
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
	return file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_rawDescGZIP(), []int{0}
}

type ListApplicationIdentitiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identities []*Identity `protobuf:"bytes,1,rep,name=identities,proto3" json:"identities,omitempty"`
}

func (x *ListApplicationIdentitiesResponse) Reset() {
	*x = ListApplicationIdentitiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListApplicationIdentitiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApplicationIdentitiesResponse) ProtoMessage() {}

func (x *ListApplicationIdentitiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_msgTypes[1]
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
	return file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_rawDescGZIP(), []int{1}
}

func (x *ListApplicationIdentitiesResponse) GetIdentities() []*Identity {
	if x != nil {
		return x.Identities
	}
	return nil
}

type WatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WatchRequest) Reset() {
	*x = WatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchRequest) ProtoMessage() {}

func (x *WatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchRequest.ProtoReflect.Descriptor instead.
func (*WatchRequest) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_rawDescGZIP(), []int{2}
}

type WatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WatchResponse) Reset() {
	*x = WatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchResponse) ProtoMessage() {}

func (x *WatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchResponse.ProtoReflect.Descriptor instead.
func (*WatchResponse) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_rawDescGZIP(), []int{3}
}

// Identity represents the identity of an application or handler.
type Identity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is the entity's unique name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Key is the entity's immutable, unique key.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *Identity) Reset() {
	*x = Identity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identity) ProtoMessage() {}

func (x *Identity) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identity.ProtoReflect.Descriptor instead.
func (*Identity) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_rawDescGZIP(), []int{4}
}

func (x *Identity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Identity) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto protoreflect.FileDescriptor

var file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67,
	0x6d, 0x61, 0x74, 0x69, 0x71, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x73, 0x70, 0x65,
	0x63, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x19, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x22, 0x22, 0x0a, 0x20, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x68, 0x0a, 0x21, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x64, 0x6f, 0x67, 0x6d, 0x61,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0x0a, 0x08, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x32, 0x84, 0x02, 0x0a,
	0x0b, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x41, 0x50, 0x49, 0x12, 0x96, 0x01, 0x0a,
	0x19, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x3b, 0x2e, 0x64, 0x6f, 0x67,
	0x6d, 0x61, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x05, 0x57, 0x61, 0x74, 0x63, 0x68, 0x12, 0x27,
	0x2e, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x30, 0x01, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x74, 0x69, 0x71, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6f, 0x70, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x73,
	0x70, 0x65, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_rawDescOnce sync.Once
	file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_rawDescData = file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_rawDesc
)

func file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_rawDescGZIP() []byte {
	file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_rawDescOnce.Do(func() {
		file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_rawDescData)
	})
	return file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_rawDescData
}

var file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_goTypes = []interface{}{
	(*ListApplicationIdentitiesRequest)(nil),  // 0: dogma.interop.v1.discover.ListApplicationIdentitiesRequest
	(*ListApplicationIdentitiesResponse)(nil), // 1: dogma.interop.v1.discover.ListApplicationIdentitiesResponse
	(*WatchRequest)(nil),                      // 2: dogma.interop.v1.discover.WatchRequest
	(*WatchResponse)(nil),                     // 3: dogma.interop.v1.discover.WatchResponse
	(*Identity)(nil),                          // 4: dogma.interop.v1.discover.Identity
}
var file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_depIdxs = []int32{
	4, // 0: dogma.interop.v1.discover.ListApplicationIdentitiesResponse.identities:type_name -> dogma.interop.v1.discover.Identity
	0, // 1: dogma.interop.v1.discover.DiscoverAPI.ListApplicationIdentities:input_type -> dogma.interop.v1.discover.ListApplicationIdentitiesRequest
	2, // 2: dogma.interop.v1.discover.DiscoverAPI.Watch:input_type -> dogma.interop.v1.discover.WatchRequest
	1, // 3: dogma.interop.v1.discover.DiscoverAPI.ListApplicationIdentities:output_type -> dogma.interop.v1.discover.ListApplicationIdentitiesResponse
	3, // 4: dogma.interop.v1.discover.DiscoverAPI.Watch:output_type -> dogma.interop.v1.discover.WatchResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_init() }
func file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_init() {
	if File_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchRequest); i {
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
		file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchResponse); i {
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
		file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identity); i {
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
			RawDescriptor: file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_goTypes,
		DependencyIndexes: file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_depIdxs,
		MessageInfos:      file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_msgTypes,
	}.Build()
	File_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto = out.File
	file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_rawDesc = nil
	file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_goTypes = nil
	file_github_com_dogmatiq_interopspec_discoverspec_discoverapi_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DiscoverAPIClient is the client API for DiscoverAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DiscoverAPIClient interface {
	// ListApplicationIdentities returns the identity of all applications.
	ListApplicationIdentities(ctx context.Context, in *ListApplicationIdentitiesRequest, opts ...grpc.CallOption) (*ListApplicationIdentitiesResponse, error)
	// Watch starts watching the server until it goes offline.
	//
	// The server MUST NOT send any responses.
	//
	// The caller SHOULD await a response (which is never sent) until an error
	// occurs, indicating that the server has gone offline.
	Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (DiscoverAPI_WatchClient, error)
}

type discoverAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewDiscoverAPIClient(cc grpc.ClientConnInterface) DiscoverAPIClient {
	return &discoverAPIClient{cc}
}

func (c *discoverAPIClient) ListApplicationIdentities(ctx context.Context, in *ListApplicationIdentitiesRequest, opts ...grpc.CallOption) (*ListApplicationIdentitiesResponse, error) {
	out := new(ListApplicationIdentitiesResponse)
	err := c.cc.Invoke(ctx, "/dogma.interop.v1.discover.DiscoverAPI/ListApplicationIdentities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoverAPIClient) Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (DiscoverAPI_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DiscoverAPI_serviceDesc.Streams[0], "/dogma.interop.v1.discover.DiscoverAPI/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &discoverAPIWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DiscoverAPI_WatchClient interface {
	Recv() (*WatchResponse, error)
	grpc.ClientStream
}

type discoverAPIWatchClient struct {
	grpc.ClientStream
}

func (x *discoverAPIWatchClient) Recv() (*WatchResponse, error) {
	m := new(WatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DiscoverAPIServer is the server API for DiscoverAPI service.
type DiscoverAPIServer interface {
	// ListApplicationIdentities returns the identity of all applications.
	ListApplicationIdentities(context.Context, *ListApplicationIdentitiesRequest) (*ListApplicationIdentitiesResponse, error)
	// Watch starts watching the server until it goes offline.
	//
	// The server MUST NOT send any responses.
	//
	// The caller SHOULD await a response (which is never sent) until an error
	// occurs, indicating that the server has gone offline.
	Watch(*WatchRequest, DiscoverAPI_WatchServer) error
}

// UnimplementedDiscoverAPIServer can be embedded to have forward compatible implementations.
type UnimplementedDiscoverAPIServer struct {
}

func (*UnimplementedDiscoverAPIServer) ListApplicationIdentities(context.Context, *ListApplicationIdentitiesRequest) (*ListApplicationIdentitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListApplicationIdentities not implemented")
}
func (*UnimplementedDiscoverAPIServer) Watch(*WatchRequest, DiscoverAPI_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}

func RegisterDiscoverAPIServer(s *grpc.Server, srv DiscoverAPIServer) {
	s.RegisterService(&_DiscoverAPI_serviceDesc, srv)
}

func _DiscoverAPI_ListApplicationIdentities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApplicationIdentitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoverAPIServer).ListApplicationIdentities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dogma.interop.v1.discover.DiscoverAPI/ListApplicationIdentities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoverAPIServer).ListApplicationIdentities(ctx, req.(*ListApplicationIdentitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscoverAPI_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DiscoverAPIServer).Watch(m, &discoverAPIWatchServer{stream})
}

type DiscoverAPI_WatchServer interface {
	Send(*WatchResponse) error
	grpc.ServerStream
}

type discoverAPIWatchServer struct {
	grpc.ServerStream
}

func (x *discoverAPIWatchServer) Send(m *WatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _DiscoverAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dogma.interop.v1.discover.DiscoverAPI",
	HandlerType: (*DiscoverAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListApplicationIdentities",
			Handler:    _DiscoverAPI_ListApplicationIdentities_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _DiscoverAPI_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/dogmatiq/interopspec/discoverspec/discoverapi.proto",
}
